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

package pkg

func StripQuotes(text string) string {
	startIndex := 0
	endIndex := len(text)
	firstChar := text[startIndex]
	lastChar := text[endIndex-1]
	if firstChar == '"' || firstChar == '\'' || firstChar == '`' {
		startIndex = 1
	}
	if lastChar == '"' || lastChar == '\'' || lastChar == '`' {
		endIndex--
	}
	return text[startIndex:endIndex]
}
