/*
Copyright 2023 The Kubernetes Authors.

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

package benchmark

import (
	"flag"
	"strings"
	"testing"

	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/kubernetes/pkg/features"
)

var testSchedulingLabelFilter = flag.String("test-scheduling-label-filter", "integration-test,-performance", "comma-separated list of labels which a testcase must have (no prefix or +) or must not have (-), used by TestScheduling")

func TestScheduling(t *testing.T) {
	testCases, err := getTestCases(configFile)
	if err != nil {
		t.Fatal(err)
	}
	if err = validateTestCases(testCases); err != nil {
		t.Fatal(err)
	}

	if testing.Short() {
		*testSchedulingLabelFilter += ",+short"
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			for _, w := range tc.Workloads {
				t.Run(w.Name, func(t *testing.T) {
					if !enabled(strings.Split(*testSchedulingLabelFilter, ","), append(tc.Labels, w.Labels...)...) {
						t.Skipf("disabled by label filter %q", *testSchedulingLabelFilter)
					}
					featureGates := featureGatesMerge(tc.FeatureGates, w.FeatureGates)
					informerFactory, tCtx := setupTestCase(t, tc, featureGates, nil, nil)

					runWorkload(tCtx, tc, w, informerFactory)

					if featureGates[features.SchedulerQueueingHints] {
						// In any case, we should make sure InFlightEvents is empty after running the scenario.
						if err = checkEmptyInFlightEvents(); err != nil {
							tCtx.Errorf("%s: %s", w.Name, err)
						}
					}

					// Reset metrics to prevent metrics generated in current workload gets
					// carried over to the next workload.
					legacyregistry.Reset()
				})
			}
		})
	}
}
