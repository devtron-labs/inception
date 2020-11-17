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
	"fmt"
	"strconv"
)

func mathematicalIntIntOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := lhs.value.(int64)
	rv := rhs.value.(int64)
	switch operator {
	case PLUS:
		return newIntValHolder(lv + rv)
	case MINUS:
		return newIntValHolder(lv - rv)
	case MULT:
		return newIntValHolder(lv * rv)
	case DIV:
		return newFloatValHolder(float64(lv) / float64(rv))
	case MOD:
		return newIntValHolder(lv % rv)
	default:
		return valHolder{}
	}
}

func mathematicalFloatFloatOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := lhs.value.(float64)
	rv := rhs.value.(float64)
	switch operator {
	case PLUS:
		return newFloatValHolder(lv + rv)
	case MINUS:
		return newFloatValHolder(lv - rv)
	case MULT:
		return newFloatValHolder(lv * rv)
	case DIV:
		return newFloatValHolder(lv / rv)
	default:
		return valHolder{}
	}
}

func mathematicalIntFloatOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := float64(lhs.value.(int64))
	rv := rhs.value.(float64)
	switch operator {
	case PLUS:
		return newFloatValHolder(lv + rv)
	case MINUS:
		return newFloatValHolder(lv - rv)
	case MULT:
		return newFloatValHolder(lv * rv)
	case DIV:
		return newFloatValHolder(lv / rv)
	default:
		return valHolder{}
	}
}

func mathematicalFloatIntOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := lhs.value.(float64)
	rv := float64(rhs.value.(int64))
	switch operator {
	case PLUS:
		return newFloatValHolder(lv + rv)
	case MINUS:
		return newFloatValHolder(lv - rv)
	case MULT:
		return newFloatValHolder(lv * rv)
	case DIV:
		return newFloatValHolder(lv / rv)
	default:
		return valHolder{}
	}
}

func mathematicalStringStringOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := lhs.value.(string)
	rv := rhs.value.(string)
	switch operator {
	case PLUS:
		return newStringValHolder(lv + rv)
	default:
		return valHolder{}
	}
}

func mathematicalStringIntOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := lhs.value.(string)
	rv := rhs.value.(int)
	rvs := strconv.Itoa(rv)
	switch operator {
	case PLUS:
		return newStringValHolder(lv + rvs)
	default:
		return valHolder{}
	}
}

func mathematicalStringFloatOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	lv := lhs.value.(string)
	rv := rhs.value.(float64)
	rvs := fmt.Sprintf("%f", rv)
	if float64(int64(rv)) == rv {
		rvs = fmt.Sprintf("%d", int64(rv))
	}
	switch operator {
	case PLUS:
		return newStringValHolder(lv + rvs)
	default:
		return valHolder{}
	}
}

func mathematicalIntStringOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	rv := rhs.value.(string)
	lv := lhs.value.(int)
	lvs := strconv.Itoa(lv)
	switch operator {
	case PLUS:
		return newStringValHolder(lvs + rv)
	default:
		return valHolder{}
	}
}

func mathematicalFloatStringOperation(lhs, rhs valHolder, operator MathematicalOperator) valHolder {
	rv := rhs.value.(string)
	lv := lhs.value.(float64)
	lvs := fmt.Sprintf("%f", lv)
	if float64(int64(lv)) == lv {
		lvs = fmt.Sprintf("%d", int64(lv))
	}
	switch operator {
	case PLUS:
		return newStringValHolder(lvs + rv)
	default:
		return valHolder{}
	}
}
