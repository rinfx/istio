// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

// ID defines underlying platform supporting service registry
type ID string

const (
	// Mock is a service registry that contains 2 hard-coded test services
	Mock ID = "Mock"
	// Kubernetes is a service registry backed by k8s API server
	Kubernetes ID = "Kubernetes"
	// External is a service registry for externally provided ServiceEntries
	External ID = "External"

	// Added by gateway
	// LocalConfig means the service info will get from local cache store.
	LocalConfig ID = "LocalConfig"
	// RemoteKubernetes means that istiod only watch remote cluster services not
	// the cluster where istiod resides.
	RemoteKubernetes ID = "RemoteKubernetes"
	// End added by gateway
)

func (id ID) String() string {
	return string(id)
}
