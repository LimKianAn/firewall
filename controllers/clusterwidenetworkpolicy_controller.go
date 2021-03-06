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

	firewallv1 "github.com/LimKianAn/firewall/api/v1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:rbac:groups=firewall.metal-stack.io,resources=events,verbs=create;patch
type ClusterwideNetworkPolicyReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

const clusterwideNPNamespace = "firewall"

// +kubebuilder:rbac:groups=firewall.metal-stack.io,resources=clusterwidenetworkpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=firewall.metal-stack.io,resources=clusterwidenetworkpolicies/status,verbs=get;update;patch
func (r *ClusterwideNetworkPolicyReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	print("Inside function Reconcile\n")
	ctx := context.Background()

	if req.Namespace != clusterwideNPNamespace {
		print("Namespace not right!\n")
		var clusterNP firewallv1.ClusterwideNetworkPolicy
		if err := r.Get(ctx, req.NamespacedName, &clusterNP); err != nil {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}

		r.Recorder.Event(&clusterNP, "Warning", "Unapplicable", fmt.Sprintf("cluster wide network policies must be defined in namespace %s otherwise they won't take effect", clusterwideNPNamespace))
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *ClusterwideNetworkPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&firewallv1.ClusterwideNetworkPolicy{}).
		Complete(r)
}
