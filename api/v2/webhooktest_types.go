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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebhookTestSpec defines the desired state of WebhookTest
type WebhookTestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conversion is an example field of WebhookTest. Edit WebhookTest_types.go to remove/update
	Conversion Conversion `json:"conversion"`
}

type Conversion struct {
	// Valid must be set to true or the validation webhook will reject the resource.
	Valid bool `json:"valid"`

	// Mutate is a field that will be set to true by the mutating webhook.
	// +optional
	Mutate bool `json:"mutate,omitempty"`
}

// WebhookTestStatus defines the observed state of WebhookTest
type WebhookTestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// WebhookTest is the Schema for the webhooktests API
type WebhookTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebhookTestSpec   `json:"spec,omitempty"`
	Status WebhookTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookTestList contains a list of WebhookTest
type WebhookTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebhookTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebhookTest{}, &WebhookTestList{})
}
