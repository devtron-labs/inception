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

import (
	parser2 "github.com/devtron-labs/inception/pkg/language/parser"
)

type KlangStatusUpdater struct {
	kubectlCommands int
	stepStatuses []*stepStatus
}

type stepStatus struct {
	name string
	status string
}

const StepInProgress = "InProgress"
const StepCompleted = "Completed"

// ExitStepInfo is called when production stepInfo is exited.
func (l *KlangStatusUpdater) ExitStepInfo(ctx *parser2.StepInfoContext) {
	step := ""
	if ctx.STRING() != nil {
		step = StripQuotes(ctx.STRING().GetText())
	} else if ctx.RAW_STRING_LIT() != nil {
		step = StripQuotes(ctx.RAW_STRING_LIT().GetText())
	}
	l.stepStatuses = append(l.stepStatuses, &stepStatus{name: step})
}

func (l *KlangStatusUpdater) ReceiveStep(name string) {
	found := false
	for _, s := range l.stepStatuses {
		if s.name == name {
			found = true
		}
	}
	if !found {
		return
	}
	for i := 0; i < len(l.stepStatuses); i++ {
		if l.stepStatuses[i].name != name {
			l.stepStatuses[i].status = StepCompleted
		} else if l.stepStatuses[i].name == name {
			l.stepStatuses[i].status = StepInProgress
			return
		}
	}
	return
}