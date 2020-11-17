/*
Copyright 2020 Devtron Labs Pvt Ltd.

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

package language

import "testing"

func TestKlangStatusUpdater_updateStepStatus(t *testing.T) {
	type fields struct {
		kubectlCommands int
		stepStatuses    []*stepStatus
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		out    []*stepStatus
	}{
		{
			name: "status update test",
			fields: fields{
				kubectlCommands: 0,
				stepStatuses:    []*stepStatus{&stepStatus{name: "1"}, &stepStatus{name: "2"}, &stepStatus{name: "3"}},
			},
			args: args{name: "2"},
			out: []*stepStatus{&stepStatus{name: "1", status: StepCompleted}, &stepStatus{name: "2", status: StepInProgress}, &stepStatus{name: "3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KlangStatusUpdater{
				kubectlCommands: tt.fields.kubectlCommands,
				stepStatuses:    tt.fields.stepStatuses,
			}
			l.ReceiveStep(tt.args.name)
			for i, s := range tt.fields.stepStatuses {
				if s.status != tt.out[i].status || s.name != tt.out[i].name {
					t.Errorf("stepStatus for index %d is name = %s, status = %s, want name=%s, status = %s", i, s.name, s.status, tt.out[i].name, tt.out[i].status)
				}
			}
		})
	}
}
