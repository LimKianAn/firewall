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
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
type ClusterwideNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PolicySpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
type ClusterwideNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterwideNetworkPolicy `json:"items"`
}

type PolicySpec struct {
	// +optional
	Description string `json:"description,omitempty"`

	// +optional
	Ingress []IngressRule `json:"ingress,omitempty"`

	// +optional
	Egress []EgressRule `json:"egress,omitempty"`
}

type IngressRule struct {
	// +optional
	Ports []networking.NetworkPolicyPort `json:"ports,omitempty"`

	// +optional
	From []networking.IPBlock `json:"from,omitempty"`
}

type EgressRule struct {
	// +optional
	Ports []networking.NetworkPolicyPort `json:"ports,omitempty"`

	// +optional
	To []networking.IPBlock `json:"to,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ClusterwideNetworkPolicy{}, &ClusterwideNetworkPolicyList{})
}
