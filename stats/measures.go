// Copyright 2020 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stats

import "go.opencensus.io/stats"

var (
	// CPURuntimeInSec measures the CPU runtime in seconds.
	CPURuntimeInSec = stats.Int64("CPURuntimeInSec", "Measures the CPU runtime in seconds", stats.UnitSeconds)
	// HTTPRequests measures the count of HTTP requests.
	HTTPRequests = stats.Int64("HTTPRequests", "Measures the count of HTTP requests", stats.UnitDimensionless)
)
