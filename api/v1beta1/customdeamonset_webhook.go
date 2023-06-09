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

package v1beta1

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var customdeamonsetlog = logf.Log.WithName("customdeamonset-resource")

func (r *CustomDeamonset) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-apps-luojun96-io-v1beta1-customdeamonset,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.luojun96.io,resources=customdeamonsets,verbs=create;update,versions=v1beta1,name=mcustomdeamonset.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &CustomDeamonset{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CustomDeamonset) Default() {
	customdeamonsetlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-apps-luojun96-io-v1beta1-customdeamonset,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.luojun96.io,resources=customdeamonsets,verbs=create;update,versions=v1beta1,name=vcustomdeamonset.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &CustomDeamonset{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CustomDeamonset) ValidateCreate() error {
	customdeamonsetlog.Info("validate create", "name", r.Name)

	// validate the image field is not empty
	if r.Spec.Image == "" {
		return errors.New("image field is empty")
	}
	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CustomDeamonset) ValidateUpdate(old runtime.Object) error {
	customdeamonsetlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CustomDeamonset) ValidateDelete() error {
	customdeamonsetlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
