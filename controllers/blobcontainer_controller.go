/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	storagesresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
)

const blobContainerFinalizerName = "blobcontainer.finalizers.com"

// BlobContainerReconciler reconciles a BlobContainer object
type BlobContainerReconciler struct {
	client.Client
	Log                  logr.Logger
	Recorder             record.EventRecorder
	BlobContainerManager storagesresourcemanager.BlobContainerManager
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=blobcontainers/status,verbs=get;update;patch

func (r *BlobContainerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx = context.Background()
	log = r.Log.WithValues("blobcontainer", req.NamespacedName)

	var instance azurev1alpha1.BlobContainer
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve blobcontainer resource", "err", err.Error())
		// TODO: What is the requeue logic here?  What exactly is client.IgnoreNotFound() doing
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if instance.IsBeingDeleted() {
		if helpers.HasFinalizer(&instance, blobContainerFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Error", "Delete blob container failed with ", err)
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, blobContainerFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !instance.HasFinalizer(blobContainerFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when adding finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.IsSubmitted() {
		err := r.reconcileExternal(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *BlobContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.BlobContainer{}).
		Complete(r)
}

// TODO ------ Everything below this line - pick up here tomorrow ------ // TODO

func (r *BlobContainerReconciler) reconcileExternal(instance *azurev1alpha1.BlobContainer) error {

}

func (r *AzureSqlFirewallRuleReconciler) deleteExternal(instance *azurev1alpha1.AzureSqlFirewallRule) error {
	// ctx := context.Background()
	// groupName := instance.Spec.ResourceGroup
	// server := instance.Spec.Server
	// ruleName := instance.ObjectMeta.Name

	// // create the Go SDK client with relevant info
	// sdk := sql.GoSDKClient{
	// 	Ctx:               ctx,
	// 	ResourceGroupName: groupName,
	// 	ServerName:        server,
	// }

	// r.Log.Info(fmt.Sprintf("deleting external resource: group/%s/server/%s/firewallrule/%s"+groupName, server, ruleName))
	// err := sdk.DeleteSQLFirewallRule(ruleName)
	// if err != nil {
	// 	if errhelp.IsStatusCode204(err) {
	// 		r.Recorder.Event(instance, v1.EventTypeWarning, "DoesNotExist", "Resource to delete does not exist")
	// 		return nil
	// 	}
	// 	msg := "Couldn't delete resouce in azure"
	// 	r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", msg)
	// 	instance.Status.Message = msg

	// 	return err
	// }
	// msg := fmt.Sprintf("Deleted %s", ruleName)
	// r.Recorder.Event(instance, v1.EventTypeNormal, "Deleted", msg)
	// instance.Status.Message = msg

	// return nil
}

func (r *AzureSqlFirewallRuleReconciler) addFinalizer(instance *azurev1alpha1.AzureSqlFirewallRule) error {
	// helpers.AddFinalizer(instance, azureSQLFirewallRuleFinalizerName)
	// err := r.Update(context.Background(), instance)
	// if err != nil {
	// 	msg := fmt.Sprintf("Failed to update finalizer: %v", err)
	// 	instance.Status.Message = msg

	// 	return fmt.Errorf("failed to update finalizer: %v", err)
	// }
	// r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", azureSQLFirewallRuleFinalizerName))
	// return nil
}
