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

import (
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	// CheckRuntime tracks CPU runtime stats.
	CheckRuntime = view.View{
		Name:        "CheckRuntime",
		Description: "CPU runtime stats per check",
		Measure:     CPURuntimeInSec,
		TagKeys:     []tag.Key{CheckName},
		//nolint:gomnd
		Aggregation: view.Distribution(
			0,
			1<<2,
			1<<3,
			1<<4,
			1<<5,
			1<<6,
			1<<7,
			1<<8,
			1<<9,
			1<<10,
			1<<11,
			1<<12,
			1<<13,
			1<<14,
			1<<15),
	}

	// OutgoingHTTPRequests tracks HTTPRequests made.
	OutgoingHTTPRequests = view.View{
		Name:        "OutgoingHTTPRequests",
		Description: "HTTPRequests made per check",
		Measure:     HTTPRequests,
		TagKeys:     []tag.Key{CheckName, RequestTag},
		Aggregation: view.Count(),
	}
)
