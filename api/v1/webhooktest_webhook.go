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

package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var webhooktestlog = logf.Log.WithName("webhooktest-resource")
var cronjoblog = logf.Log.WithName("cronjob-resource")

func (r *WebhookTest) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-webhook-operators-coreos-io-v1-webhooktest,mutating=true,failurePolicy=fail,groups=webhook.operators.coreos.io,resources=webhooktests,verbs=create;update,versions=v1,name=mwebhooktest.kb.io

var _ webhook.Defaulter = &WebhookTest{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *WebhookTest) Default() {
	webhooktestlog.Info("default", "name", r.Name)

	if r.Spec.Mutate != true {
		r.Spec.Mutate = true
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-webhook-operators-coreos-io-v1-webhooktest,mutating=false,failurePolicy=fail,groups=webhook.operators.coreos.io,resources=webhooktests,versions=v1,name=vwebhooktest.kb.io

var _ webhook.Validator = &WebhookTest{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *WebhookTest) ValidateCreate() error {
	webhooktestlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return r.validateWebhookTest()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *WebhookTest) ValidateUpdate(old runtime.Object) error {
	webhooktestlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return r.validateWebhookTest()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *WebhookTest) ValidateDelete() error {
	webhooktestlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *WebhookTest) validateWebhookTest() error {
	var allErrs field.ErrorList
	if r.Spec.Valid != true {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("schedule"), r.Spec.Valid, "Spec.Valid must be true"))
	}

	if len(allErrs) != 0 {
		return apierrors.NewInvalid(
			schema.GroupKind{Group: "test.operators.coreos.com", Kind: "WebhookTest"},
			r.Name, allErrs)
	}

	return nil
}
