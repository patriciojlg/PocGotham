package models

type HandlerValidator struct {
	Validators []func(string) error
	ValueModel string
}

func NewHandlerValidators(value *string) *HandlerValidator {
	return &HandlerValidator{
		ValueModel: *value,
	}
}

func (c *HandlerValidator) SetValue(value string) *HandlerValidator {
	c.ValueModel = value
	return c
}
func (c *HandlerValidator) AddFuncValidator(validator func(string) error) *HandlerValidator {
	c.Validators = append(c.Validators, validator)
	return c
}

func (c *HandlerValidator) AddArrayFuncsValidators(validators []func(string) error) *HandlerValidator {
	c.Validators = append(c.Validators, validators...)
	return c
}
func (c *HandlerValidator) ValidateValue() error {
	for _, validator := range c.Validators {
		if err := validator(c.ValueModel); err != nil {
			return err
		}
	}
	return nil
}
