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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	workloadsv1alpha1 "github.com/ciberkleid/k8s-workload-operator/api/v1alpha1"
)

// WebSvcReconciler reconciles a WebSvc object
type WebSvcReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=workloads.k8s.coraiberkleid.xyz,resources=websvcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=workloads.k8s.coraiberkleid.xyz,resources=websvcs/status,verbs=get;update;patch

func (r *WebSvcReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("websvc", req.NamespacedName)

	// your logic here
	log.Info("in Reconcile func")
	var resource workloadsv1alpha1.WebSvc
	r.Get(ctx, req.NamespacedName, &resource)

	var deployReplicas int32
	switch resource.Spec.DeploymentTier {
	case "dev":
		deployReplicas = int32(2)
	}

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resource.Name,
			Namespace: resource.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deployReplicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": resource.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": resource.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  resource.Name,
							Image: resource.Spec.Image,
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(&resource, deploy, r.Scheme); err != nil {
		log.Error(err, "Unable to set owner reference on deployment")
		return ctrl.Result{}, err
	}

	if err := r.Create(ctx, deploy); err != nil {
		log.Error(err, "Unable to create deployment")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *WebSvcReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&workloadsv1alpha1.WebSvc{}).
		Complete(r)
}
