package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	PhasePending          = "PENDING"
	PhaseAuthUnifi        = "AUTH_UNIFI"
	PhaseGetUnifiSites    = "GET_UNIFI_SITES"
	PhaseGetUnifiClients  = "GET_UNIFI_CLIENTS"
	PhaseGetUnifiDevices  = "GET_UNIFI_DEVICES"
	PhaseGetUnifiUDM      = "GET_UNIFI_UDM"
	PhaseGetUnifiUSG      = "GET_UNIFI_USG"
	PhaseGetUnifiUSW      = "GET_UNIFI_USW"
	PhaseGetUnifiNetworks = "GET_UNIFI_NETWORKS"
	PhaseDone             = "DONE"
	PhaseRefresh          = "REFRESH"
)

type CattleNetworkUnifiInfo struct {
	User      string `json:"user"`
	Pass      string `json:"pass"`
	URL       string `json:"url"`
	VerifySSL bool   `json:"verifySSL"`
}

type CattleUnifiTemperature struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CattleUnifiDownlinkTable struct {
	Mac string `json:"mac"`
}

type CattleNetworkUnifiWanWrapper struct {
	DNS     []string `json:"dns"`     // may be deprecated
	Gateway string   `json:"gateway"` // may be deprecated
	IP      string   `json:"ip"`
	Ifname  string   `json:"ifname"`
	Mac     string   `json:"mac"`
	Media   string   `json:"media"`
	Name    string   `json:"name"`
	Netmask string   `json:"netmask"` // may be deprecated
	NumPort int      `json:"num_port"`
	PortIdx int      `json:"port_idx"`
	Type    string   `json:"type"`
}

type CattleNetworkUnifiSite struct {
	AttrHiddenID string `json:"attr_hidden_id"`
	Desc         string `json:"desc"`
	ID           string `json:"_id"`
	Name         string `json:"name"`
	SiteName     string `json:"-"`
	SourceName   string `json:"-"`
}

type CattleNetworkUnifiUDMWrapper struct {
	AdoptIP               string                     `json:"adopt_ip"`
	AdoptURL              string                     `json:"adopt_url"`
	Architecture          string                     `json:"architecture"`
	BandsteeringMode      string                     `json:"bandsteering_mode"`
	Cfgversion            string                     `json:"cfgversion"`
	ConnectRequestIP      string                     `json:"connect_request_ip"`
	ConnectRequestPort    string                     `json:"connect_request_port"`
	ConnectionNetworkName string                     `json:"connection_network_name"`
	DeviceDomain          string                     `json:"device_domain"`
	DeviceID              string                     `json:"device_id"`
	DiscoveredVia         string                     `json:"discovered_via"`
	DisplayableVersion    string                     `json:"displayable_version"`
	DownlinkTable         []CattleUnifiDownlinkTable `json:"downlink_table"`
	GuestToken            string                     `json:"guest_token"`
	ID                    string                     `json:"_id"`
	IP                    string                     `json:"ip"`
	InformIP              string                     `json:"inform_ip"`
	InformURL             string                     `json:"inform_url"`
	KernelVersion         string                     `json:"kernel_version"`
	KnownCfgversion       string                     `json:"known_cfgversion"`
	LanIP                 string                     `json:"lan_ip"`
	LastWlanIP            string                     `json:"last_wan_ip"`
	LcmNightModeBegins    string                     `json:"lcm_night_mode_begins"`
	LcmNightModeEnds      string                     `json:"lcm_night_mode_ends"`
	LcmTrackerSeed        string                     `json:"lcm_tracker_seed"`
	LicenseState          string                     `json:"license_state"`
	Mac                   string                     `json:"mac"`
	Model                 string                     `json:"model"`
	Name                  string                     `json:"name"`
	RequiredVersion       string                     `json:"required_version"`
	Serial                string                     `json:"serial"`
	SiteID                string                     `json:"site_id"`
	SiteName              string                     `json:"-"`
	SourceName            string                     `json:"-"`
	StpVersion            string                     `json:"stp_version"`
	SyslogKey             string                     `json:"syslog_key"`
	Temperatures          []CattleUnifiTemperature   `json:"temperatures,omitempty"`
	Type                  string                     `json:"type"`
	Version               string                     `json:"version"`
	WlangroupIDNa         string                     `json:"wlangroup_id_na"`
	WlangroupIDNg         string                     `json:"wlangroup_id_ng"`
	XInformAuthkey        string                     `json:"x_inform_authkey"`
}

type CattleNetworkUnifiClientWrapper struct {
	ApMac       string `json:"ap_mac,omitempty"`
	Blocked     bool   `json:"blocked,omitempty"`
	Bssid       string `json:"bssid,omitempty"`
	Essid       string `json:"essid,omitempty"`
	FixedIP     string `json:"fixed_ip,omitempty"`
	GwMac       string `json:"gw_mac,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	ID          string `json:"_id,omitempty"`
	IP          string `json:"ip,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Name        string `json:"name,omitempty"`
	Network     string `json:"network,omitempty"`
	NetworkID   string `json:"network_id,omitempty"`
	Note        string `json:"note,omitempty"`
	Radio       string `json:"radio,omitempty"`
	RadioName   string `json:"radio_name,omitempty"`
	RadioProto  string `json:"radio_proto,omitempty"`
	SiteID      string `json:"site_id,omitempty"`
	UserGroupID string `json:"usergroup_id,omitempty"`
	UserID      string `json:"user_id,omitempty"`
}

type CattleNetworkUnifiUSWWrapper struct {
	Architecture          string `json:"architecture"`
	Cfgversion            string `json:"cfgversion"`
	ConnectRequestIP      string `json:"connect_request_ip"`
	ConnectRequestPort    string `json:"connect_request_port"`
	ConnectionNetworkName string `json:"connection_network_name"`
	DeviceID              string `json:"device_id"`
	DisplayableVersion    string `json:"displayable_version"`
	GatewayMac            string `json:"gateway_mac"`
	ID                    string `json:"_id"`
	IP                    string `json:"ip"`
	InformIP              string `json:"inform_ip"`
	InformURL             string `json:"inform_url"`
	KernelVersion         string `json:"kernel_version"`
	KnownCfgversion       string `json:"known_cfgversion"`
	LedOverride           string `json:"led_override"`
	LicenseState          string `json:"license_state"`
	Mac                   string `json:"mac"`
	Model                 string `json:"model"`
	Name                  string `json:"name"`
	OutdoorModeOverride   string `json:"outdoor_mode_override"`
	PowerSource           string `json:"power_source"`
	PowerSourceVoltage    string `json:"power_source_voltage"`
	RequiredVersion       string `json:"required_version"`
	Serial                string `json:"serial"`
	SetupID               string `json:"setup_id"`
	SiteID                string `json:"site_id"`
	SiteName              string `json:"-"`
	SourceName            string `json:"-"`
	StpVersion            string `json:"stp_version"`
	Type                  string `json:"type"`
	Version               string `json:"version"`
}

type CattleNetworkUnifiUSGWrapper struct {
	Cfgversion          string                       `json:"cfgversion"`
	ConnectRequestIP    string                       `json:"connect_request_ip"`
	ConnectRequestPort  string                       `json:"connect_request_port"`
	DeviceID            string                       `json:"device_id"`
	GuestToken          string                       `json:"guest_token"`
	ID                  string                       `json:"_id"`
	InformIP            string                       `json:"inform_ip"`
	InformURL           string                       `json:"inform_url"`
	IP                  string                       `json:"ip"`
	KnownCfgversion     string                       `json:"known_cfgversion"`
	LedOverride         string                       `json:"led_override"`
	LicenseState        string                       `json:"license_state"`
	Mac                 string                       `json:"mac"`
	Model               string                       `json:"model"`
	Name                string                       `json:"name"`
	OutdoorModeOverride string                       `json:"outdoor_mode_override"`
	RequiredVersion     string                       `json:"required_version"`
	Serial              string                       `json:"serial"`
	SiteID              string                       `json:"site_id"`
	SiteName            string                       `json:"-"`
	SourceName          string                       `json:"-"`
	Temperatures        []CattleUnifiTemperature     `json:"temperatures,omitempty"`
	Type                string                       `json:"type"`
	Version             string                       `json:"version"`
	Wan1                CattleNetworkUnifiWanWrapper `json:"wan1"`
	Wan2                CattleNetworkUnifiWanWrapper `json:"wan2"`
}

type CattleNetworkUnifiFlexBool struct {
	Val bool   `json:"val"`
	Txt string `json:"txt"`
}

type CattleNetworkUnifiFlexInt struct {
	Val string `json:"val"`
	Txt string `json:"txt"`
}

type CattleNetworkUnifiNetworkWrapper struct {
	DhcpGuardEnabled       CattleNetworkUnifiFlexBool `json:"dhcpguard_enabled"`
	DhcpRelayEnabled       CattleNetworkUnifiFlexBool `json:"dhcp_relay_enabled"`
	DhcpdDNSEnabled        CattleNetworkUnifiFlexBool `json:"dhcpd_dns_enabled"`
	DhcpdEnabled           CattleNetworkUnifiFlexBool `json:"dhcpd_enabled"`
	DhcpdGatewayEnabled    CattleNetworkUnifiFlexBool `json:"dhcpd_gateway_enabled"`
	DhcpdIP1               string                     `json:"dhcpd_ip_1"`
	DhcpdLeasetime         CattleNetworkUnifiFlexInt  `json:"dhcpd_leasetime"`
	DhcpdTimeOffsetEnabled CattleNetworkUnifiFlexBool `json:"dhcpd_time_offset_enabled"`
	DomainName             string                     `json:"domain_name"`
	Enabled                CattleNetworkUnifiFlexBool `json:"enabled"`
	ID                     string                     `fake:"{uuid}"                    json:"_id"`
	IPSubnet               string                     `json:"ip_subnet"`
	IsNat                  CattleNetworkUnifiFlexBool `json:"is_nat"`
	Name                   string                     `json:"name"`
	Networkgroup           string                     `json:"networkgroup"`
	Purpose                string                     `json:"purpose"`
	SiteID                 string                     `fake:"{uuid}"                    json:"site_id"`
	Vlan                   CattleNetworkUnifiFlexInt  `json:"vlan"`
	VlanEnabled            CattleNetworkUnifiFlexBool `json:"vlan_enabled"`
}

type CattleNetworkUnifiResults struct {
	USWS          []CattleNetworkUnifiUSWWrapper     `json:"usws,omitempty"`
	UnifiNetworks []CattleNetworkUnifiNetworkWrapper `json:"unifinetworks,omitempty"`
	USGS          []CattleNetworkUnifiUSGWrapper     `json:"usgs,omitempty"`
	UDMS          []CattleNetworkUnifiUDMWrapper     `json:"udms,omitempty"`
	Clients       []CattleNetworkUnifiClientWrapper  `json:"clients,omitempty"`
	Sites         []CattleNetworkUnifiSite           `json:"sites,omitempty"`
}

// CattleNetworkLoadoutSpec defines the desired state of CattleNetworkLoadout
type CattleNetworkLoadoutSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CattleNetworkLoadout. Edit cattlenetworkloadout_types.go to remove/update
	Unifi        CattleNetworkUnifiInfo    `json:"unifi,omitempty"`
	UnifiResults CattleNetworkUnifiResults `json:"unifiResults,omitempty"`
}

// CattleNetworkLoadoutStatus defines the observed state of CattleNetworkLoadout
type CattleNetworkLoadoutStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase string `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CattleNetworkLoadout is the Schema for the cattlenetworkloadouts API
type CattleNetworkLoadout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CattleNetworkLoadoutSpec   `json:"spec,omitempty"`
	Status CattleNetworkLoadoutStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CattleNetworkLoadoutList contains a list of CattleNetworkLoadout
type CattleNetworkLoadoutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CattleNetworkLoadout `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CattleNetworkLoadout{}, &CattleNetworkLoadoutList{})
}
