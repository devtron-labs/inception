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

func logicalIntIntOperation(lhs, rhs valHolder, operator LogicalOperator) bool {
	lv := lhs.value.(int64)
	rv := rhs.value.(int64)
	switch operator {
	case EQ:
		return lv == rv
	case NEQ:
		return lv != rv
	case LTEQ:
		return lv <= rv
	case LT:
		return lv < rv
	case GTEQ:
		return lv >= rv
	case GT:
		return lv > rv
	default:
		return false
	}
}

func logicalFloatFloatOperation(lhs, rhs valHolder, operator LogicalOperator) bool {
	lv := lhs.value.(float64)
	rv := rhs.value.(float64)
	switch operator {
	case EQ:
		return lv == rv
	case NEQ:
		return lv != rv
	case LTEQ:
		return lv <= rv
	case LT:
		return lv < rv
	case GTEQ:
		return lv >= rv
	case GT:
		return lv > rv
	default:
		return false
	}
}

func logicalIntFloatOperation(lhs, rhs valHolder, operator LogicalOperator) bool {
	lv := float64(lhs.value.(int64))
	rv := rhs.value.(float64)
	switch operator {
	case EQ:
		return lv == rv
	case NEQ:
		return lv != rv
	case LTEQ:
		return lv <= rv
	case LT:
		return lv < rv
	case GTEQ:
		return lv >= rv
	case GT:
		return lv > rv
	default:
		return false
	}
}

func logicalFloatIntOperation(lhs, rhs valHolder, operator LogicalOperator) bool {
	lv := lhs.value.(float64)
	rv := float64(rhs.value.(int64))
	switch operator {
	case EQ:
		return lv == rv
	case NEQ:
		return lv != rv
	case LTEQ:
		return lv <= rv
	case LT:
		return lv < rv
	case GTEQ:
		return lv >= rv
	case GT:
		return lv > rv
	default:
		return false
	}
}

func logicalStringStringOperation(lhs, rhs valHolder, operator LogicalOperator) bool {
	lv := lhs.value.(string)
	rv := rhs.value.(string)
	switch operator {
	case EQ:
		return lv == rv
	case NEQ:
		return lv != rv
	case LTEQ:
		return lv <= rv
	case LT:
		return lv < rv
	case GTEQ:
		return lv >= rv
	case GT:
		return lv > rv
	default:
		return false
	}
}

