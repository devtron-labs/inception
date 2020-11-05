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

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func JsonSelect(json, pattern string) valHolder {
	val := gjson.Get(json, pattern)
	v := toValHolder(val)
	return v
}

func JsonEdit(json, pattern string, value interface{}) valHolder {
	val, err := sjson.Set(json, pattern, value)
	if err != nil {
		return newErrHolder(err)
	}
	return newStringValHolder(val)
}

func JsonDelete(json, pattern string) valHolder {
	val, err := sjson.Delete(json, pattern)
	if err != nil {
		return newErrHolder(err)
	}
	return newStringValHolder(val)
}

func equalResult(v1, v2 gjson.Result) bool {
	if v1.Type != v2.Type {
		return false
	}
	if v1.Type == gjson.Null {
		return true
	}
	if v1.Type == gjson.False {
		return true
	}
	if v1.Type == gjson.True {
		return true
	}
	if v1.Type == gjson.Number {
		return v1.Num == v1.Num
	}
	if v1.Type == gjson.String {
		return v1.Str == v2.Str
	}
	if v1.Type == gjson.JSON {
		//TODO: use better algorithm to compare json
		return v1.Raw == v2.Raw
	}
	return false
}

// String returns a string representation of the type.
func toValHolder(result gjson.Result) valHolder {
	switch result.Type {
	default:
		return valHolder{}
	case gjson.Null:
		return valHolder{
			value:    "nil",
			dataType: "NIL",
		}
	case gjson.False:
		return valHolder{
			value:    false,
			dataType: "BOOLEAN",
		}
	case gjson.Number:
		return valHolder{
			value:    result.Num,
			dataType: "FLOAT",
		}
	case gjson.String:
		return valHolder{
			value:    result.Str,
			dataType: "STRING",
		}
	case gjson.True:
		return valHolder{
			value:    true,
			dataType: "BOOLEAN",
		}
	case gjson.JSON:
		return valHolder{
			value:    result.Raw,
			dataType: "STRING",
		}
	}
}
