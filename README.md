# Validate

Current package contains tools to validate arbitrary data in Go. 
This main difference from other competitors - usage of type casting instead of functions, like others.
Core is `validate.Interface`:

```go
type Interface interface {
	Validate() error
}
```

Any object (even yours) that implements this interface is correct validator. 
To use bundled validators just cast your value and check error presence:

```go
func foo(name string) error {
        if err := validate.StringNotWhitespace(name); err != nil {
                return err
        }
        
        // Your BL
}
```

To perform more than one check at once, use `validate.All` function:
```go
func foo(number int) error {
        if err := validate.All(validate.IntPositive(number), validate.IntOdd(number)); err != nil {
                // If both validation fails, resulting error will contain combined message
                return err
        }
        
        // Your BL
}
```

# Validator types

| Cast from | Name | Valid values | Invalid values |  |
| --------- | ---- | ------------ | -------------- | --- |
| `int` | `validate.IntPositive` | `1` | `0`, `-1` | |
| `int` | `validate.IntOdd` | `1`, `27`, `-3` | `0`, `2` | |
| `int` | `validate.IntEven` | `0`, `-2`, `1024` | `15`, `-1` | |
| `string` | `validate.StringAlpha` | `"abc"`, `"foo"`, `"русский"` | `""`, `"123abc"` | Only UTF-8 letters |
| `string` | `validate.StringAlpha` | `"abc"`, `"foo"` | `""`, `"123abc"`, `"русский"` | Only latin-1 letters |
| `string` | `validate.StringNotEmpty` | `"abc"`, `"foo"`, `" "` | `""` | |
| `string` | `validate.StringNotWhitespace` | `"abc"`, `"foo"` | `""`, `" "`, `" \n "` | |
