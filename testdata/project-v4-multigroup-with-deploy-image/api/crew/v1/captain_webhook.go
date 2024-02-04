/*
Copyright 2024 The Kubernetes authors.

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
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var captainlog = logf.Log.WithName("captain-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Captain) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		WithValidator(r).
		WithDefaulter(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-crew-testproject-org-v1-captain,mutating=true,failurePolicy=fail,sideEffects=None,groups=crew.testproject.org,resources=captains,verbs=create;update,versions=v1,name=mcaptain.kb.io,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Captain{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type
func (r *Captain) Default(ctx context.Context, obj runtime.Object) error {
	captainlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-crew-testproject-org-v1-captain,mutating=false,failurePolicy=fail,sideEffects=None,groups=crew.testproject.org,resources=captains,verbs=create;update,versions=v1,name=vcaptain.kb.io,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Captain{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type
func (r *Captain) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	captainlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type
func (r *Captain) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	captainlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (r *Captain) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	captainlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
