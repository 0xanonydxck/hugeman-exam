package validator

var validator *Validator

func Init() {
	validator = NewValidator()
}

func Struct(s interface{}) error {
	return validator.Struct(s)
}

func Var(s interface{}, tag string) error {
	return validator.Var(s, tag)
}
