# validate
Usage:
```go
type Person struct {
    Age  int       `@validate:"@MinInt(0);@MaxInt(200)"`    // 0 <= age <= 200
    Name string    `@validate:"@NotEmpty;@MaxLength(20)"` // name should not be empty, and max length is 20
    Address string `@validate:"@NotEmpty;@Regexp(^[a-zA-Z_][0-9a-zA-Z_]*$)"` // address should not be empty, and should match regular expression '^[a-zA-Z_][0-9a-zA-Z_]*$'
}

// err is nil
err := validate.ValidateStructByTags(&Person{
    Age:     30,
    Name:    "name",
    Address: "China",
})

// err: min is 0
err := validate.ValidateStructByTags(&Person{
    Age:     -1,
    Name:    "name",
    Address: "China",
})

// err: name too long
err := validate.ValidateStructByTags(&Person{
    Age:     30,
    Name:    "name_looooooooooooooooooooooooong",
    Address: "China",
})

// err: name is empty
err := validate.ValidateStructByTags(&Person{
    Age:     30,
    Name:    "",
    Address: "China",
})

// err: name too long
err := validate.ValidateStructByTags(&Person{
    Age:     30,
    Name:    "name",
    Address: "China",
})
```

## Supported Validator

|type|validator|err message|remark|
|:-:|:-|:-|:-|
|int/int8/int16/int32/int64|@NotZero||should not be zero (!=0)|
|int/int8/int16/int32/int64|@Zero||should be zero (=0)|
|int/int8/int16/int32/int64|@One||should be one (=1)|
|int/int8/int16/int32/int64|@PositiveInt||should be positive int (>0)|
|int/int8/int16/int32/int64|@NonNegativeInt||should be non-negative int (>=0)|
|int/int8/int16/int32/int64|@NonPositiveInt||should be non-positive int (<=0)|
|int/int8/int16/int32/int64|@NegativeInt||should be negative int (<0)|
|int/int8/int16/int32/int64|@MaxInt(3)||should less or equal than 3 (<=3)|
|int/int8/int16/int32/int64|@MinInt(3)||should greater or equal than 3 (>=3)|
|int/int8/int16/int32/int64|@IntIn(1,3,5)||shoule be 1, 3 or 5|
|int/int8/int16/int32/int64|@IntNotIn(1,3,5)||should not be 1, 3 or 5|
|string|@NotEmpty||should not be empty string|
|string|@Empty||should be empty string|
|string|@NotBlank||should not be blank|
|string|@Regexp(**expr**)||should match regular expression **expr** |
|string|@MaxLength(10)||length should less than 10|
|string|@MinLength(10)||length should longer than 10|
|string|@StringNotIn(a,bbb,ccc)||string should in **a**, **bbb** or **ccc**|
|string|@StringIn(aa,bb,cc)||string not should be **a**, **bbb** or **ccc**|


	
	
