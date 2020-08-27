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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Interval",type=string,JSONPath=`.spec.interval`
// +kubebuilder:printcolumn:name="InternalPrefixes",type=string,JSONPath=`.spec.internalprefixes`
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirewallSpec   `json:"spec,omitempty"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

type FirewallSpec struct {
	Interval         string      `json:"interval,omitempty"`
	DryRun           bool        `json:"dryrun,omitempty"`
	Ipv4RuleFile     string      `json:"ipv4rulefile,omitempty"`
	RateLimits       []RateLimit `json:"ratelimits,omitempty"`
	InternalPrefixes []string    `json:"internalprefixes,omitempty"`
}

type FirewallStatus struct {
	Message       string        `json:"message,omitempty"`
	FirewallStats FirewallStats `json:"stats"`
	Updated       metav1.Time   `json:"lastRun,omitempty"`
}

type FirewallStats struct {
	RuleStats   RuleStatsByAction   `json:"rules"`
	DeviceStats DeviceStatsByDevice `json:"devices"`
	IDSStats    IDSStatsByDevice    `json:"idsstats"`
}

type RuleStatsByAction map[string]RuleStats

type RuleStats map[string]RuleStat

type RuleStat struct {
	Counter Counter `json:"counter"`
}

type Counter struct {
	Bytes   uint64 `json:"bytes"`
	Packets uint64 `json:"packets"`
}

type RateLimit struct {
	Interface string `json:"interface,omitempty"`
	Rate      uint32 `json:"rate,omitempty"`
}

type DeviceStatsByDevice map[string]DeviceStat

type DeviceStat struct {
	InBytes    uint64 `json:"in"`
	OutBytes   uint64 `json:"out"`
	TotalBytes uint64 `json:"total"`
}

type IDSStatsByDevice map[string]InterfaceStat

type InterfaceStat struct {
	Drop             int `json:"drop"`
	InvalidChecksums int `json:"invalidchecksums"`
	Packets          int `json:"packets"`
}

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
