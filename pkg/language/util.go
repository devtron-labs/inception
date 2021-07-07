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
	"math/rand"
	"strings"
)

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

var chars = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func Generate(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}