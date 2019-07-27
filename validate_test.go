package validate

import (
	"testing"
)

func TestValidateStringNotIn(t *testing.T) {
	validElems := []interface{}{
		struct {
			validIn string `@validate:"@StringNotIn(aaa, bbb, ccc)"`
		}{"a"},
		struct {
			validIn string `@validate:"@StringNotIn(aaa, bbb, ccc)"`
		}{"e"},
		struct {
			validIn string `@validate:"@StringNotIn(aaa, bbb, ccc)"`
		}{"f"},
	}

	for _, elem := range validElems {
		if err := ValidateStructByTags(elem); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	invalidElems := []interface{}{
		struct {
			invalidIn string `@validate:"@StringNotIn(aaa, bbb, ccc)"`
		}{"aaa"},
		struct {
			invalidIn string `@validate:"@StringNotIn(aaa, bbb, ccc)"`
		}{"bbb"},
		struct {
			invalidIn string `@validate:"@StringNotIn(aaa, bbb, ccc)"`
		}{"ccc"},
	}

	for _, elem := range invalidElems {
		if err := ValidateStructByTags(elem); err == nil {
			t.Errorf("invalid value, should throw exception, value:%d", elem)
		}
	}
}

func TestValidateStringIn(t *testing.T) {
	validElems := []interface{}{
		struct {
			validIn string `@validate:"@StringIn(aaa, bbb, ccc)"`
		}{"aaa"},
		struct {
			validIn string `@validate:"@StringIn(aaa, bbb, ccc)"`
		}{"bbb"},
		struct {
			validIn string `@validate:"@StringIn(aaa, bbb, ccc)"`
		}{"ccc"},
	}

	for _, elem := range validElems {
		if err := ValidateStructByTags(elem); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	invalidElems := []interface{}{
		struct {
			invalidIn string `@validate:"@StringIn(aaa, bbb, ccc)"`
		}{"a"},
		struct {
			invalidIn string `@validate:"@StringIn(aaa, bbb, ccc)"`
		}{"e"},
		struct {
			invalidIn string `@validate:"@StringIn(aaa, bbb, ccc)"`
		}{"f"},
	}

	for _, elem := range invalidElems {
		if err := ValidateStructByTags(elem); err == nil {
			t.Errorf("invalid value, should throw exception, value:%d", elem)
		}
	}
}

func TestValidateIntNotIn(t *testing.T) {
	validElems := []interface{}{
		struct {
			validIn int `@validate:"@IntNotIn(1, 4, 8)"`
		}{2},
		struct {
			validIn int `@validate:"@IntNotIn(1, 4, 8)"`
		}{3},
		struct {
			validIn int `@validate:"@IntNotIn(1, 4, 8)"`
		}{5},
	}

	for _, elem := range validElems {
		if err := ValidateStructByTags(elem); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	invalidElems := []interface{}{
		struct {
			invalidIn int `@validate:"@IntNotIn(1, 4, 8)"`
		}{1},
		struct {
			invalidIn int `@validate:"@IntNotIn(1, 4, 8)"`
		}{4},
		struct {
			invalidIn int `@validate:"@IntNotIn(1, 4, 8)"`
		}{8},
	}

	for _, elem := range invalidElems {
		if err := ValidateStructByTags(elem); err == nil {
			t.Errorf("invalid value, should throw exception, value:%d", elem)
		}
	}
}
func TestValidateIntIn(t *testing.T) {
	validElems := []interface{}{
		struct {
			validIn int `@validate:"@IntIn(1, 4, 8)"`
		}{1},
		struct {
			validIn int `@validate:"@IntIn(1, 4, 8)"`
		}{4},
		struct {
			validIn int `@validate:"@IntIn(1, 4, 8)"`
		}{8},
	}

	for _, elem := range validElems {
		if err := ValidateStructByTags(elem); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	invalidElems := []interface{}{
		struct {
			invalidIn int `@validate:"@IntIn(1, 4, 8)"`
		}{2},
		struct {
			invalidIn int `@validate:"@IntIn(1, 4, 8)"`
		}{3},
		struct {
			invalidIn int `@validate:"@IntIn(1, 4, 8)"`
		}{5},
	}

	for _, elem := range invalidElems {
		if err := ValidateStructByTags(elem); err == nil {
			t.Errorf("invalid value, should throw exception, value:%d", elem)
		}
	}
}

func TestValidateStructInt(t *testing.T) {
	validElems := []interface{}{
		struct {
			maxInt int `@validate:"@MaxInt(10)"`
		}{2},
		struct {
			maxInt int `@validate:"@MaxInt(10)"`
		}{10},
		struct {
			minInt int `@validate:"@MinInt(10)"`
		}{12},
		struct {
			minInt int `@validate:"@MinInt(10)"`
		}{10},
		struct {
			maxMinInt int `@validate:"@MaxInt(10);@MinInt(5)"`
		}{8},
		struct {
			notZero int `@validate:"@NotZero"`
		}{6},
		struct {
			zero int `@validate:"@Zero"`
		}{0},
		struct {
			one int `@validate:"@One"`
		}{1},
	}

	for _, elem := range validElems {
		if err := ValidateStructByTags(elem); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	invalidElems := []interface{}{
		struct {
			maxInt int `@validate:"@MaxInt(10)"`
		}{12},
		struct {
			minInt int `@validate:"@MinInt(10)"`
		}{2},
		struct {
			maxMinInt int `@validate:"@MaxInt(10);@MinInt(5)"`
		}{3},
		struct {
			maxMinInt int `@validate:"@MaxInt(10);@MinInt(5)"`
		}{18},
		struct {
			notZero int `@validate:"@NotZero"`
		}{0},
		struct {
			zero int `@validate:"@Zero"`
		}{7},
		struct {
			one int `@validate:"@One"`
		}{16},
	}

	for _, elem := range invalidElems {
		if err := ValidateStructByTags(elem); err == nil {
			t.Errorf("invalid value, should throw exception, value:%d", elem)
		}
	}

}

func TestValidateStructString(t *testing.T) {
	validElems := []interface{}{
		struct {
			fieldMaxLen1 string `@validate:"@MaxLength(3)"`
		}{"aa"},
		struct {
			fieldMaxLen2 string `@validate:"@MaxLength(6)"`
		}{"bbbbbb"},
		struct {
			fieldMinLen1 string `@validate:"@MinLength(3)"`
		}{"cccc"},
		struct {
			filedMinLen2 string `@validate:"@MinLength(6)"`
		}{"dddddd"},
		struct {
			fieldMaxMinLen string `@validate:"@MinLength(3);@MaxLength(6)"`
		}{"eeeee"},
		struct {
			fieldNotEmpty1 string `@validate:"@NotEmpty"`
		}{"ffff"},
		struct {
			fieldNotEmpty2 string `@validate:"@NotEmpty"`
		}{" "},
		struct {
			fieldEmpty string `@validate:"@Empty"`
		}{""},
		struct {
			fieldNotEmpty1 string `@validate:"@NotBlank"`
		}{"gggg"},
		struct {
			fieldNotEmpty2 string `@validate:"@NotBlank"`
		}{" h "},
		struct {
			fieldNotEmpty2 string `@validate:"@NotBlank"`
		}{"i i"},
	}

	for _, elem := range validElems {
		if err := ValidateStructByTags(elem); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	validRegexp := []string{"_aB", "_P", "O_", "N", "_m", "l_", "kk", "_"}

	for _, elem := range validRegexp {
		if err := ValidateStructByTags(struct {
			regexpo string `@validate:"@Regexp(^[a-zA-Z_][0-9a-zA-Z_]*$)"`
		}{elem}); err != nil {
			t.Errorf("valid value, should not throw exception, actual:%+v", err)
		}
	}

	invalidElems := []interface{}{
		struct {
			fieldMaxLen1 string `@validate:"@MaxLength(3)"`
		}{"aaaa"},
		struct {
			fieldMaxLen2 string `@validate:"@MaxLength(6)"`
		}{"bbbbbbbbb"},
		struct {
			fieldMinLen1 string `@validate:"@MinLength(3)"`
		}{"cc"},
		struct {
			filedMinLen2 string `@validate:"@MinLength(6)"`
		}{"ddddd"},
		struct {
			fieldMaxMinLen string `@validate:"@MinLength(3);@MaxLength(6)"`
		}{"ee"},
		struct {
			fieldMaxMinLen2 string `@validate:"@MinLength(3);@MaxLength(6)"`
		}{"fffffff"},
		struct {
			fieldNotEmpty1 string `@validate:"@NotEmpty"`
		}{""},
		struct {
			fieldEmpty string `@validate:"@Empty"`
		}{" "},
		struct {
			fieldNotEmpty1 string `@validate:"@NotBlank"`
		}{""},
		struct {
			fieldNotEmpty2 string `@validate:"@NotBlank"`
		}{" "},
	}

	for _, elem := range invalidElems {
		if err := ValidateStructByTags(elem); err == nil {
			t.Errorf("invalid value, should throw exception, element:%s", elem)
		}
	}

	invalidRegexp := []string{"0a", "1A", "2_", "_$", "a%", "b*", "A)", "B(", "C#", "D~"}

	for _, elem := range invalidRegexp {
		if err := ValidateStructByTags(struct {
			regexpo string `@validate:"@Regexp(^[a-zA-Z_][0-9a-zA-Z_]*$)"`
		}{elem}); err == nil {
			t.Errorf("invalid regexp value, should throw exception, element:%s", elem)
		}
	}
}

func TestValidateInt(t *testing.T) {
	if err := ValidateStructByTags(3); err != nil {
		t.Errorf("not support validate int now, should not throw exception, error: %+v", err)
	}
}

func TestValidateString(t *testing.T) {
	if err := ValidateStructByTags("ssss"); err != nil {
		t.Errorf("not support validate string now, should not throw exception, error: %+v", err)
	}
}

func TestValidateBool(t *testing.T) {
	if err := ValidateStructByTags(1.3333); err != nil {
		t.Errorf("not support validate bool now, should not throw exception, error: %+v", err)
	}
}

func TestValidateFloat(t *testing.T) {
	if err := ValidateStructByTags(1.3333); err != nil {
		t.Errorf("not support validate float now, should not throw exception, error: %+v", err)
	}
}

func TestValidateFloatPtr(t *testing.T) {
	var value float32 = 1.333
	if err := ValidateStructByTags(&value); err != nil {
		t.Errorf("not support validate float ptr now, should not throw exception, error: %+v", err)
	}
}
