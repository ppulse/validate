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
	tags := strings.Split(tagsStr, ",")
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
	case trimedTag == "":
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
	case strings.HasPrefix(trimedTag, "@Regexp(") && strings.HasSuffix(trimedTag, ")"):
		re := trimedTag[len("@Regexp(") : len(trimedTag)-1]
		return stringRegexp(value, re)
	case strings.HasPrefix(trimedTag, "@MaxLength(") && strings.HasSuffix(trimedTag, ")"):
		maxLen, err := strconv.Atoi(trimedTag[len("@MaxLength(") : len(trimedTag)-1])
		if err != nil {
			return err
		}
		return stringMaxLength(value, maxLen)
	case strings.HasPrefix(trimedTag, "@MaxInt(") && strings.HasSuffix(trimedTag, ")"):
		max, err := strconv.Atoi(trimedTag[len("@MaxInt(") : len(trimedTag)-1])
		if err != nil {
			return err
		}
		return intMax(value, max)
	case strings.HasPrefix(trimedTag, "@MinLength(") && strings.HasSuffix(trimedTag, ")"):
		minLen, err := strconv.Atoi(trimedTag[len("@MinLength(") : len(trimedTag)-1])
		if err != nil {
			return err
		}
		return stringMinLength(value, minLen)
	case strings.HasPrefix(trimedTag, "@MinInt(") && strings.HasSuffix(trimedTag, ")"):
		min, err := strconv.Atoi(trimedTag[len("@MinInt(") : len(trimedTag)-1])
		if err != nil {
			return err
		}
		return intMin(value, min)
	default:
	}

	return nil
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
