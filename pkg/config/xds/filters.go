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

package xds

import (
	resource "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
)

const (
	WasmHTTPFilterType = resource.APITypePrefix + wellknown.HTTPWasm
	RBACHTTPFilterType = resource.APITypePrefix + "envoy.extensions.filters.http.rbac.v3.RBAC"
	// Added by Ingress
	CompositeHTTPFilterType = resource.APITypePrefix + "envoy.extensions.filters.http.composite.v3.Composite"
	// Added by Ingress
	TypedStructType = resource.APITypePrefix + "udpa.type.v1.TypedStruct"

	StatsFilterName       = "istio.stats"
	StackdriverFilterName = "istio.stackdriver"
)
