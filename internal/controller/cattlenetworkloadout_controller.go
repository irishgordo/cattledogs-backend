package controller

import (
	"context"
	cattledogsv1 "github.com/irishgordo/cattledogs-backend/api/v1"
	"github.com/unpoller/unifi"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	otherlogger "log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strconv"
	"time"
)

// CattleNetworkLoadoutReconciler reconciles a CattleNetworkLoadout object
type CattleNetworkLoadoutReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cattledogs.cattledogs-backend.gordoughnet.com,resources=cattlenetworkloadouts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cattledogs.cattledogs-backend.gordoughnet.com,resources=cattlenetworkloadouts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cattledogs.cattledogs-backend.gordoughnet.com,resources=cattlenetworkloadouts/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CattleNetworkLoadout object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *CattleNetworkLoadoutReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// TODO(user): your logic here

	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	reqLogger := logger.WithValues("cattlenetworkloadout", req.NamespacedName)
	reqLogger.Info("=== Reconciling CattleNetworkLoadout")
	instance := &cattledogsv1.CattleNetworkLoadout{}
	err := r.Client.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	reqLogger.Info("phase check...")
	reqLogger.WithValues("Current-Phase", instance.Status.Phase)
	if instance.Status.Phase == "" || instance.Status.Phase == cattledogsv1.PhaseRefresh {
		instance.Status.Phase = cattledogsv1.PhasePending
		r.Status().Update(ctx, instance)
	}
	if instance.Status.Phase == cattledogsv1.PhaseDone {
		reqLogger.Info("req should be done...")
		reqLogger.Info("done with cattle-network-loadout")
		return ctrl.Result{
			Requeue: false,
		}, nil
	}

	if instance.Spec.Unifi.Pass != "" && instance.Status.Phase != "DONE" {
		reqLogger.Info("=== unifi password provided... leveraging with other details")
		var validatedUser string = "admin"
		var validatedPass string = "pass"
		var validatedURL string = "https://127.0.0.1"
		var validateVerifySSL bool = false
		if instance.Spec.Unifi.User != "" {
			validatedUser = instance.Spec.Unifi.User
		}
		if instance.Spec.Unifi.Pass != "" {
			validatedPass = instance.Spec.Unifi.Pass
		}
		if instance.Spec.Unifi.URL != "" {
			validatedURL = instance.Spec.Unifi.URL
		}
		if instance.Spec.Unifi.VerifySSL != validateVerifySSL {
			validateVerifySSL = instance.Spec.Unifi.VerifySSL
		}

		unifiConfig := &unifi.Config{
			User:      validatedUser,
			Pass:      validatedPass,
			URL:       validatedURL,
			VerifySSL: validateVerifySSL,
			ErrorLog:  otherlogger.Printf,
			DebugLog:  otherlogger.Printf,
		}

		uni, err := unifi.NewUnifi(unifiConfig)

		if err != nil {
			return ctrl.Result{}, err
		} else {
			instance.Status.Phase = cattledogsv1.PhaseAuthUnifi
			r.Status().Update(ctx, instance)
		}

		reqLogger.Info("=== able to acquire unifi client...")
		time.Sleep(10 * time.Second)
		sites, err := uni.GetSites()
		time.Sleep(10 * time.Second)
		if err != nil {
			return ctrl.Result{}, err
		} else {
			instance.Status.Phase = cattledogsv1.PhaseGetUnifiSites
			r.Status().Update(ctx, instance)
			var newUnifiSites []cattledogsv1.CattleNetworkUnifiSite
			for _, site := range sites {
				var newSite cattledogsv1.CattleNetworkUnifiSite
				newSite.AttrHiddenID = site.AttrHiddenID
				newSite.Desc = site.Desc
				newSite.ID = site.ID
				newSite.Name = site.Name
				newSite.SiteName = site.SiteName
				newSite.SourceName = site.SourceName
				newUnifiSites = append(newUnifiSites, newSite)
			}
			instance.Spec.UnifiResults.Sites = newUnifiSites
			r.Update(ctx, instance)
		}

		reqLogger.Info("=== able to acquire unifi sites...")
		time.Sleep(10 * time.Second)
		clients, err := uni.GetClients(sites)
		if err != nil {
			return ctrl.Result{}, err
		} else {
			instance.Status.Phase = cattledogsv1.PhaseGetUnifiClients
			r.Status().Update(ctx, instance)
		}

		time.Sleep(10 * time.Second)
		var newClientsToUse []cattledogsv1.CattleNetworkUnifiClientWrapper
		for i, clientNew := range clients {
			newClient := cattledogsv1.CattleNetworkUnifiClientWrapper{}
			newClient.ApMac = clientNew.ApMac
			newClient.Blocked = clientNew.Blocked
			newClient.Bssid = clientNew.Bssid
			newClient.Essid = clientNew.Essid
			newClient.FixedIP = clientNew.FixedIP
			newClient.GwMac = clientNew.GwMac
			newClient.Hostname = clientNew.Hostname
			newClient.ID = clientNew.ID
			newClient.IP = clientNew.IP
			newClient.Mac = clientNew.Mac
			newClient.Name = clientNew.Name
			newClient.Network = clientNew.Network
			newClient.NetworkID = clientNew.NetworkID
			newClient.Note = clientNew.Note
			newClient.Radio = clientNew.Radio
			newClient.RadioName = clientNew.RadioName
			newClient.RadioProto = clientNew.RadioProto
			newClient.SiteID = clientNew.SiteID
			newClient.UserGroupID = clientNew.UserGroupID
			newClient.UserID = clientNew.UserID
			newClientsToUse = append(newClientsToUse, newClient)
			i++
		}
		reqLogger.Info("=== finished setting clients...")
		if len(newClientsToUse) > 0 {
			instance.Spec.UnifiResults.Clients = newClientsToUse
			r.Update(ctx, instance)
		}

		reqLogger.Info("=== able to acquire unifi clients on the network...")
		time.Sleep(20 * time.Second)
		instance.Status.Phase = cattledogsv1.PhaseGetUnifiDevices
		r.Status().Update(ctx, instance)
		devices, err := uni.GetDevices(sites)
		time.Sleep(20 * time.Second)
		if err != nil {
			return ctrl.Result{}, err
		} else {
			instance.Status.Phase = cattledogsv1.PhaseGetUnifiUDM
			r.Status().Update(ctx, instance)
			var udmList []cattledogsv1.CattleNetworkUnifiUDMWrapper
			for _, udm := range devices.UDMs {
				var newUDM cattledogsv1.CattleNetworkUnifiUDMWrapper
				newUDM.AdoptURL = udm.AdoptURL
				newUDM.AdoptIP = udm.AdoptIP
				newUDM.Architecture = udm.Architecture
				newUDM.BandsteeringMode = udm.BandsteeringMode
				newUDM.Cfgversion = udm.Cfgversion
				newUDM.ConnectRequestIP = udm.ConnectRequestIP
				newUDM.ConnectRequestPort = udm.ConnectRequestPort
				newUDM.ConnectionNetworkName = udm.ConnectionNetworkName
				newUDM.DeviceDomain = udm.DeviceDomain
				newUDM.DeviceID = udm.DeviceID
				newUDM.DiscoveredVia = udm.DiscoveredVia
				newUDM.DisplayableVersion = udm.DisplayableVersion
				var newUDMDownlinks []cattledogsv1.CattleUnifiDownlinkTable
				for _, item := range udm.DownlinkTable {
					var newDownlink cattledogsv1.CattleUnifiDownlinkTable
					newDownlink.Mac = item.Mac
					newUDMDownlinks = append(newUDMDownlinks, newDownlink)
				}
				newUDM.DownlinkTable = newUDMDownlinks
				newUDM.GuestToken = udm.GuestToken
				newUDM.ID = udm.ID
				newUDM.IP = udm.IP
				newUDM.InformIP = udm.InformIP
				newUDM.InformURL = udm.InformURL
				newUDM.KernelVersion = udm.KernelVersion
				newUDM.KnownCfgversion = udm.KnownCfgversion
				newUDM.LanIP = udm.LanIP
				newUDM.LastWlanIP = udm.LastWlanIP
				newUDM.LcmNightModeBegins = udm.LcmNightModeBegins
				newUDM.LcmNightModeEnds = udm.LcmNightModeEnds
				newUDM.LcmTrackerSeed = udm.LcmTrackerSeed
				newUDM.LicenseState = udm.LicenseState
				newUDM.Mac = udm.Mac
				newUDM.Model = udm.Model
				newUDM.Name = udm.Name
				newUDM.RequiredVersion = udm.RequiredVersion
				newUDM.Serial = udm.Serial
				newUDM.SiteID = udm.SiteID
				newUDM.SiteName = udm.SiteName
				newUDM.SourceName = udm.SourceName
				newUDM.StpVersion = udm.StpVersion
				newUDM.SyslogKey = udm.SyslogKey
				var newUDMTemps []cattledogsv1.CattleUnifiTemperature
				for _, itrTemp := range udm.Temperatures {
					var newTemp cattledogsv1.CattleUnifiTemperature
					newTemp.Name = itrTemp.Name
					newTemp.Value = strconv.FormatFloat(itrTemp.Value, 'f', -1, 64)
					newTemp.Type = itrTemp.Type
					newUDMTemps = append(newUDMTemps, newTemp)
				}
				newUDM.Temperatures = newUDMTemps
				newUDM.Type = udm.Type
				newUDM.Version = udm.Version
				newUDM.WlangroupIDNg = udm.WlangroupIDNg
				newUDM.WlangroupIDNa = udm.WlangroupIDNa
				newUDM.XInformAuthkey = udm.XInformAuthkey
				udmList = append(udmList, newUDM)
			}
			instance.Spec.UnifiResults.UDMS = udmList
			r.Update(ctx, instance)
			instance.Status.Phase = cattledogsv1.PhaseGetUnifiUSW
			r.Status().Update(ctx, instance)
			var uswFoundList []cattledogsv1.CattleNetworkUnifiUSWWrapper
			for _, uswFound := range devices.USWs {
				var newUSWToAdd cattledogsv1.CattleNetworkUnifiUSWWrapper
				newUSWToAdd.Architecture = uswFound.Architecture
				newUSWToAdd.Cfgversion = uswFound.Cfgversion
				newUSWToAdd.ConnectRequestIP = uswFound.ConnectRequestIP
				newUSWToAdd.ConnectRequestPort = uswFound.ConnectRequestPort
				newUSWToAdd.ConnectionNetworkName = uswFound.ConnectionNetworkName
				newUSWToAdd.DeviceID = uswFound.DeviceID
				newUSWToAdd.DisplayableVersion = uswFound.DisplayableVersion
				newUSWToAdd.GatewayMac = uswFound.GatewayMac
				newUSWToAdd.ID = uswFound.ID
				newUSWToAdd.IP = uswFound.IP
				newUSWToAdd.InformIP = uswFound.InformIP
				newUSWToAdd.InformURL = uswFound.InformURL
				newUSWToAdd.KernelVersion = uswFound.KernelVersion
				newUSWToAdd.KnownCfgversion = uswFound.KnownCfgversion
				newUSWToAdd.LedOverride = uswFound.LedOverride
				newUSWToAdd.LicenseState = uswFound.LicenseState
				newUSWToAdd.Mac = uswFound.Mac
				newUSWToAdd.Model = uswFound.Model
				newUSWToAdd.Name = uswFound.Name
				newUSWToAdd.OutdoorModeOverride = uswFound.OutdoorModeOverride
				newUSWToAdd.PowerSource = uswFound.PowerSource
				newUSWToAdd.PowerSourceVoltage = uswFound.PowerSourceVoltage
				newUSWToAdd.RequiredVersion = uswFound.RequiredVersion
				newUSWToAdd.Serial = uswFound.Serial
				newUSWToAdd.SetupID = uswFound.SetupID
				newUSWToAdd.SiteID = uswFound.SiteID
				newUSWToAdd.SiteName = uswFound.SiteName
				newUSWToAdd.SourceName = uswFound.SourceName
				newUSWToAdd.StpVersion = uswFound.StpVersion
				newUSWToAdd.Type = uswFound.Type
				newUSWToAdd.Version = uswFound.Version
				uswFoundList = append(uswFoundList, newUSWToAdd)
			}
			instance.Spec.UnifiResults.USWS = uswFoundList
			r.Update(ctx, instance)
			instance.Status.Phase = cattledogsv1.PhaseGetUnifiUSG
			r.Status().Update(ctx, instance)
			var newUSGList []cattledogsv1.CattleNetworkUnifiUSGWrapper
			for _, usgFound := range devices.USGs {
				var newUSG cattledogsv1.CattleNetworkUnifiUSGWrapper
				newUSG.Cfgversion = usgFound.Cfgversion
				newUSG.ConnectRequestIP = usgFound.ConnectRequestIP
				newUSG.ConnectRequestPort = usgFound.ConnectRequestPort
				newUSG.DeviceID = usgFound.DeviceID
				newUSG.GuestToken = usgFound.GuestToken
				newUSG.ID = usgFound.ID
				newUSG.InformIP = usgFound.InformIP
				newUSG.InformURL = usgFound.InformURL
				newUSG.IP = usgFound.IP
				newUSG.KnownCfgversion = usgFound.KnownCfgversion
				newUSG.LedOverride = usgFound.LedOverride
				newUSG.LicenseState = usgFound.LicenseState
				newUSG.Mac = usgFound.Mac
				newUSG.Model = usgFound.Model
				newUSG.Name = usgFound.Name
				newUSG.OutdoorModeOverride = usgFound.OutdoorModeOverride
				newUSG.RequiredVersion = usgFound.RequiredVersion
				newUSG.Serial = usgFound.Serial
				newUSG.SiteID = usgFound.SiteID
				newUSG.SourceName = usgFound.SourceName
				var newTemperatureList []cattledogsv1.CattleUnifiTemperature
				for _, temperatureFound := range usgFound.Temperatures {
					var newTemperature cattledogsv1.CattleUnifiTemperature
					newTemperature.Name = temperatureFound.Name
					newTemperature.Type = temperatureFound.Type
					newTemperature.Value = strconv.FormatFloat(temperatureFound.Value, 'f', -1, 64)
					newTemperatureList = append(newTemperatureList, newTemperature)
				}
				newUSG.Temperatures = newTemperatureList
				newUSG.Type = usgFound.Type
				newUSG.Version = usgFound.Version
				var newUSGWan1 cattledogsv1.CattleNetworkUnifiWanWrapper
				var wan1Found = usgFound.Wan1
				newUSGWan1.DNS = wan1Found.DNS
				newUSGWan1.Gateway = wan1Found.Gateway
				newUSGWan1.IP = wan1Found.IP
				newUSGWan1.Ifname = wan1Found.Ifname
				newUSGWan1.Mac = wan1Found.Mac
				newUSGWan1.Media = wan1Found.Media
				newUSGWan1.Name = wan1Found.Name
				newUSGWan1.Netmask = wan1Found.Netmask
				newUSGWan1.NumPort = wan1Found.NumPort
				newUSGWan1.PortIdx = wan1Found.PortIdx
				newUSGWan1.Type = wan1Found.Type
				newUSG.Wan1 = newUSGWan1
				var newUSGWan2 cattledogsv1.CattleNetworkUnifiWanWrapper
				var wan2Found = usgFound.Wan2
				newUSGWan2.DNS = wan2Found.DNS
				newUSGWan2.Gateway = wan2Found.Gateway
				newUSGWan2.IP = wan2Found.IP
				newUSGWan2.Ifname = wan2Found.Ifname
				newUSGWan2.Mac = wan2Found.Mac
				newUSGWan2.Media = wan2Found.Media
				newUSGWan2.Name = wan2Found.Name
				newUSGWan2.Netmask = wan2Found.Netmask
				newUSGWan2.NumPort = wan2Found.NumPort
				newUSGWan2.PortIdx = wan2Found.PortIdx
				newUSGWan2.Type = wan2Found.Type
				newUSG.Wan2 = newUSGWan2
				newUSGList = append(newUSGList, newUSG)
			}
			instance.Spec.UnifiResults.USGS = newUSGList
			r.Update(ctx, instance)
		}
		//
		reqLogger.Info("=== able to acquire unifi devices...")
		instance.Status.Phase = cattledogsv1.PhaseGetUnifiNetworks
		r.Status().Update(ctx, instance)
		time.Sleep(20 * time.Second)
		networksFound, err := uni.GetNetworks(sites)
		time.Sleep(20 * time.Second)
		reqLogger.Info("=== able to get networks")
		var networksFoundToAdd []cattledogsv1.CattleNetworkUnifiNetworkWrapper
		for _, netFound := range networksFound {
			var newNet cattledogsv1.CattleNetworkUnifiNetworkWrapper
			newNet.DhcpGuardEnabled.Txt = netFound.DhcpGuardEnabled.Txt
			newNet.DhcpGuardEnabled.Val = netFound.DhcpGuardEnabled.Val
			newNet.DhcpdDNSEnabled.Txt = netFound.DhcpdDNSEnabled.Txt
			newNet.DhcpdDNSEnabled.Val = netFound.DhcpdDNSEnabled.Val
			newNet.DhcpRelayEnabled.Txt = netFound.DhcpRelayEnabled.Txt
			newNet.DhcpRelayEnabled.Val = netFound.DhcpRelayEnabled.Val
			newNet.DhcpdDNSEnabled.Txt = netFound.DhcpdDNSEnabled.Txt
			newNet.DhcpdDNSEnabled.Val = netFound.DhcpdDNSEnabled.Val
			newNet.DhcpdIP1 = netFound.DhcpdIP1
			newNet.DhcpdLeasetime.Txt = netFound.DhcpdLeasetime.Txt
			newNet.DhcpdLeasetime.Val = strconv.FormatFloat(netFound.DhcpdLeasetime.Val, 'f', -1, 64)
			newNet.DhcpdTimeOffsetEnabled.Val = netFound.DhcpdTimeOffsetEnabled.Val
			newNet.DhcpdTimeOffsetEnabled.Txt = netFound.DhcpdTimeOffsetEnabled.Txt
			newNet.DomainName = netFound.DomainName
			newNet.Enabled.Txt = netFound.Enabled.Txt
			newNet.Enabled.Val = netFound.Enabled.Val
			newNet.ID = netFound.ID
			newNet.IPSubnet = netFound.IPSubnet
			newNet.IsNat.Val = netFound.IsNat.Val
			newNet.IsNat.Val = netFound.IsNat.Val
			newNet.Name = netFound.Name
			newNet.Networkgroup = netFound.Networkgroup
			newNet.Purpose = netFound.Purpose
			newNet.SiteID = netFound.SiteID
			newNet.Vlan.Txt = netFound.Vlan.Txt
			newNet.Vlan.Val = strconv.FormatFloat(netFound.Vlan.Val, 'f', -1, 64)
			newNet.Vlan.Txt = netFound.Vlan.Txt
			newNet.VlanEnabled.Txt = netFound.VlanEnabled.Txt
			newNet.VlanEnabled.Val = netFound.VlanEnabled.Val
			networksFoundToAdd = append(networksFoundToAdd, newNet)
		}
		instance.Spec.UnifiResults.UnifiNetworks = networksFoundToAdd
		r.Update(ctx, instance)
		instance.Status.Phase = cattledogsv1.PhaseDone
		r.Status().Update(ctx, instance)
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CattleNetworkLoadoutReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cattledogsv1.CattleNetworkLoadout{}).
		Complete(r)
}
