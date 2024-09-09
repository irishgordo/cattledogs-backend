package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cattledogsv1 "github.com/irishgordo/cattledogs-backend/api/v1"
)

// CattleLocationReconciler reconciles a CattleLocation object
type CattleLocationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cattledogs.cattledogs-backend.gordoughnet.com,resources=cattlelocations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cattledogs.cattledogs-backend.gordoughnet.com,resources=cattlelocations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cattledogs.cattledogs-backend.gordoughnet.com,resources=cattlelocations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CattleLocation object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *CattleLocationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	reqLogger := logger.WithValues("cattlelocation", req.NamespacedName)
	reqLogger.Info("=== Reconciling CattleLocation")
	instance := &cattledogsv1.CattleLocation{}
	err := r.Client.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if instance.Spec.CattleLocationActivity.String() == "" {
		instance.Spec.CattleLocationActivity = cattledogsv1.CattleLocationActivity(0)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CattleLocationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cattledogsv1.CattleLocation{}).
		Complete(r)
}
