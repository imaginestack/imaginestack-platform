/*

 Copyright 2023 The ImagineKube Authors.

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

package v1alpha3

import "testing"

func TestLabelMatch(t *testing.T) {
	tests := []struct {
		labels       map[string]string
		filter       string
		expectResult bool
	}{
		{
			labels: map[string]string{
				"imaginekube.com/workspace": "imaginekube-system",
			},
			filter:       "imaginekube.com/workspace",
			expectResult: true,
		},
		{
			labels: map[string]string{
				"imaginekube.com/creator": "system",
			},
			filter:       "imaginekube.com/workspace",
			expectResult: false,
		},
		{
			labels: map[string]string{
				"imaginekube.com/workspace": "imaginekube-system",
			},
			filter:       "imaginekube.com/workspace=",
			expectResult: false,
		},
		{
			labels: map[string]string{
				"imaginekube.com/workspace": "imaginekube-system",
			},
			filter:       "imaginekube.com/workspace!=",
			expectResult: true,
		},
		{
			labels: map[string]string{
				"imaginekube.com/workspace": "imaginekube-system",
			},
			filter:       "imaginekube.com/workspace!=imaginekube-system",
			expectResult: false,
		},
		{
			labels: map[string]string{
				"imaginekube.com/workspace": "imaginekube-system",
			},
			filter:       "imaginekube.com/workspace=imaginekube-system",
			expectResult: true,
		},
		{
			labels: map[string]string{
				"imaginekube.com/workspace": "imaginekube-system",
			},
			filter:       "imaginekube.com/workspace=system",
			expectResult: false,
		},
	}
	for i, test := range tests {
		result := labelMatch(test.labels, test.filter)
		if result != test.expectResult {
			t.Errorf("case %d, got %#v, expected %#v", i, result, test.expectResult)
		}
	}
}
