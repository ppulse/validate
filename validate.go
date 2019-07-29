package validate

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ValidateStructByTags(data interface{}) error {
	switch reflect.ValueOf(data).Kind() {
	case reflect.Ptr:
		return validatePtr(data)
	case reflect.Struct:
		return validateStruct(reflect.ValueOf(data))
	}

	return nil
}

func validatePtr(data interface{}) error {
	v := reflect.ValueOf(data).Elem()
	switch v.Kind() {
	case reflect.Struct:
		return validateStruct(v)
	}
	return nil

}

func validateStruct(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		if e := validateTags(tag.Get("@validate"), v.Field(i)); e != nil {
			return fmt.Errorf("%s:%s", fieldInfo.Name, e.Error())
		}
	}

	return nil
}

func validateTags(tagsStr string, value reflect.Value) error {
	tags := strings.Split(tagsStr, ";")
	for _, tag := range tags {
		if err := validateEachTag(tag, value); err != nil {
			return err
		}
	}

	return nil
}

func validateEachTag(tag string, value reflect.Value) error {
	trimedTag := strings.TrimSpace(tag)
	switch {
	case trimedTag == "@NotZero":
		return intNotEquals(value, 0)
	case trimedTag == "@Zero":
		return intZero(value)
	case trimedTag == "@One":
		return intOne(value)
	case trimedTag == "@NotEmpty":
		return stringNotEmpty(value)
	case trimedTag == "@Empty":
		return stringEmpty(value)
	case trimedTag == "@NotBlank":
		return stringNotBlank(value)
	case trimedTag == "@PositiveInt":
		return intMin(value, 1)
	case trimedTag == "@NonNegativeInt":
		return intMin(value, 0)
	case trimedTag == "@NonPositiveInt":
		return intMax(value, 0)
	case trimedTag == "@NegativeInt":
		return intMax(value, -1)
	case strings.HasPrefix(trimedTag, "@Regexp(") && strings.HasSuffix(trimedTag, ")"):
		re := trimedTag[len("@Regexp(") : len(trimedTag)-1]
		return stringRegexp(value, re)
	case strings.HasPrefix(trimedTag, "@MaxLength(") && strings.HasSuffix(trimedTag, ")"):
		maxLen, err := strconv.Atoi(trimedTag[len("@MaxLength(") : len(trimedTag)-1])
		if err != nil {
			return fmt.Errorf("tag error: %+v", err)
		}
		return stringMaxLength(value, maxLen)
	case strings.HasPrefix(trimedTag, "@MaxInt(") && strings.HasSuffix(trimedTag, ")"):
		max, err := strconv.Atoi(trimedTag[len("@MaxInt(") : len(trimedTag)-1])
		if err != nil {
			return fmt.Errorf("tag error: %+v", err)
		}
		return intMax(value, max)
	case strings.HasPrefix(trimedTag, "@MinLength(") && strings.HasSuffix(trimedTag, ")"):
		minLen, err := strconv.Atoi(trimedTag[len("@MinLength(") : len(trimedTag)-1])
		if err != nil {
			return fmt.Errorf("tag error: %+v", err)
		}
		return stringMinLength(value, minLen)
	case strings.HasPrefix(trimedTag, "@MinInt(") && strings.HasSuffix(trimedTag, ")"):
		min, err := strconv.Atoi(trimedTag[len("@MinInt(") : len(trimedTag)-1])
		if err != nil {
			return fmt.Errorf("tag error: %+v", err)
		}
		return intMin(value, min)
	case strings.HasPrefix(trimedTag, "@IntIn(") && strings.HasSuffix(trimedTag, ")"):
		elems := strings.Split(trimedTag[len("@IntIn("):len(trimedTag)-1], ",")
		elemsInt := make([]int64, 0)
		for _, elem := range elems {
			elemInt, err := strconv.Atoi(strings.TrimSpace(elem))
			if err != nil {
				return fmt.Errorf("tag error: %+v", err)
			}

			elemsInt = append(elemsInt, int64(elemInt))
		}
		return intIn(value, elemsInt)
	case strings.HasPrefix(trimedTag, "@IntNotIn(") && strings.HasSuffix(trimedTag, ")"):
		elems := strings.Split(trimedTag[len("@IntNotIn("):len(trimedTag)-1], ",")
		elemsInt := make([]int64, 0)
		for _, elem := range elems {
			elemInt, err := strconv.Atoi(strings.TrimSpace(elem))
			if err != nil {
				return fmt.Errorf("tag error: %+v", err)
			}

			elemsInt = append(elemsInt, int64(elemInt))
		}
		return intNotIn(value, elemsInt)
	case strings.HasPrefix(trimedTag, "@StringNotIn(") && strings.HasSuffix(trimedTag, ")"):
		elems := strings.Split(trimedTag[len("@StringNotIn("):len(trimedTag)-1], ",")
		trimedElems := make([]string, 0)
		for _, elem := range elems {
			trimedElems = append(trimedElems, strings.TrimSpace(elem))
		}
		return stringNotIn(value, trimedElems)
	case strings.HasPrefix(trimedTag, "@StringIn(") && strings.HasSuffix(trimedTag, ")"):
		elems := strings.Split(trimedTag[len("@StringIn("):len(trimedTag)-1], ",")
		trimedElems := make([]string, 0)
		for _, elem := range elems {
			trimedElems = append(trimedElems, strings.TrimSpace(elem))
		}
		return stringIn(value, trimedElems)
	default:
	}

	return nil
}

func isTagHasOneParameterMatch(tag string, prefix string, suffix string) bool {
	return strings.HasPrefix(tag, prefix) && strings.HasSuffix(tag, suffix)
}

func intIn(v reflect.Value, elems []int64) error {
	return validateInt(v, func(v reflect.Value) bool {
		for _, elem := range elems {
			if elem == v.Int() {
				return false
			}
		}
		return true
	}, fmt.Errorf("expected %+v, actual is %d", elems, v.Int()))
}

func intNotIn(v reflect.Value, elems []int64) error {
	return validateInt(v, func(v reflect.Value) bool {
		for _, elem := range elems {
			if elem == v.Int() {
				return true
			}
		}
		return false
	}, fmt.Errorf("expected %+v, actual is %d", elems, v.Int()))
}

func stringIn(v reflect.Value, elems []string) error {
	return validateString(v, func(v reflect.Value) bool {
		for _, elem := range elems {
			if elem == v.String() {
				return false
			}
		}
		return true
	}, fmt.Errorf("expected %+v, actual is %s", elems, v.String()))
}

func stringNotIn(v reflect.Value, elems []string) error {
	return validateString(v, func(v reflect.Value) bool {
		for _, elem := range elems {
			if elem == v.String() {
				return true
			}
		}
		return false
	}, fmt.Errorf("expected %+v, actual is %s", elems, v.String()))
}

func intZero(v reflect.Value) error {
	return intEqual(v, 0)
}

func intOne(v reflect.Value) error {
	return intEqual(v, 1)
}

func intMinusOne(v reflect.Value) error {
	return intEqual(v, -1)
}

func intNotEquals(v reflect.Value, unexpected int64) error {
	return validateInt(v, func(v reflect.Value) bool {
		return v.Int() == unexpected
	}, fmt.Errorf("expected not equals, unexpected is %d, actual is %d", unexpected, v.Int()))
}

func intEqual(v reflect.Value, expected int64) error {
	return validateInt(v, func(v reflect.Value) bool {
		return v.Int() != expected
	}, fmt.Errorf("expected equals, expected is %d, actual is %d", expected, v.Int()))
}

func intMax(v reflect.Value, max int) error {
	return validateInt(v, func(v reflect.Value) bool {
		return v.Int() > int64(max)
	}, fmt.Errorf("too large, max is %d", max))
}

func intMin(v reflect.Value, min int) error {
	return validateInt(v, func(v reflect.Value) bool {
		return v.Int() < int64(min)
	}, fmt.Errorf("too small, min is %d", min))
}

func validateInt(v reflect.Value, cmp func(reflect.Value) bool, err error) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if cmp(v) {
			return err
		}
	default:
		return errors.New("not int type")
	}
	return nil
}

func stringRegexp(v reflect.Value, re string) error {
	return validateString(v, func(v reflect.Value) bool {
		match, err := regexp.MatchString(re, v.String())
		if err != nil {
			return true
		}
		if !match {
			return true
		}
		return false
	}, fmt.Errorf("not match, format:%s", re))
}

func stringMaxLength(v reflect.Value, maxLen int) error {
	return validateString(v, func(v reflect.Value) bool {
		return len(v.String()) > maxLen
	}, fmt.Errorf("too long, max is %d", maxLen))
}

func stringMinLength(v reflect.Value, minLen int) error {
	return validateString(v, func(v reflect.Value) bool {
		return len(v.String()) < minLen
	}, fmt.Errorf("too short, min is %d", minLen))
}

func stringNotEmpty(v reflect.Value) error {
	return validateString(v, func(v reflect.Value) bool {
		return len(v.String()) == 0
	}, fmt.Errorf("is empty"))
}

func stringEmpty(v reflect.Value) error {
	return validateString(v, func(v reflect.Value) bool {
		return len(v.String()) != 0
	}, fmt.Errorf("is not empty"))
}

func stringBlank(v reflect.Value) error {
	return validateString(v, func(v reflect.Value) bool {
		return len(strings.TrimSpace(v.String())) != 0
	}, fmt.Errorf("is not blank"))
}

func stringNotBlank(v reflect.Value) error {
	return validateString(v, func(v reflect.Value) bool {
		return len(strings.TrimSpace(v.String())) == 0
	}, fmt.Errorf("is blank"))
}

func validateString(v reflect.Value, cmp func(reflect.Value) bool, err error) error {
	switch v.Kind() {
	case reflect.String:
		if cmp(v) {
			return err
		}
	default:
		return errors.New("not string type")
	}

	return nil
}
