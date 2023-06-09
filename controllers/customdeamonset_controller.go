/*
Copyright 2023 Jun Luo.

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

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1beta1 "github.com/luojun96/k8s-operator-opt/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomDeamonsetReconciler reconciles a CustomDeamonset object
type CustomDeamonsetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.luojun96.io,resources=customdeamonsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.luojun96.io,resources=customdeamonsets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.luojun96.io,resources=customdeamonsets/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CustomDeamonset object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *CustomDeamonsetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	cds := &appsv1beta1.CustomDeamonset{}
	if err := r.Get(ctx, req.NamespacedName, cds); err != nil {
		fmt.Println(err)
	}

	if cds.Spec.Image != "" {
		nodeList := &v1.NodeList{}
		if err := r.List(ctx, nodeList); err != nil {
			fmt.Println(err)
		}
		for _, node := range nodeList.Items {
			p := v1.Pod{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "v1",
					Kind:       "Pod",
				},
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: fmt.Sprintf("%s-", node.Name),
					Namespace:    cds.Namespace,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Image: cds.Spec.Image,
							Name:  "custom-container",
						},
					},
					NodeName: node.Name,
				},
			}
			if err := r.Create(ctx, &p); err != nil {
				fmt.Println(err)
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomDeamonsetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta1.CustomDeamonset{}).
		Complete(r)
}
