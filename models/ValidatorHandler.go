package models

type HandlerErrors struct {
	Validators []func(string) error
	ValueModel string
}

func NewHandlerErrors(value string) *HandlerErrors {
	return &HandlerErrors{
		ValueModel: value,
	}
}

func (c *HandlerErrors) SetValue(value string) *HandlerErrors {
	c.ValueModel = value
	return c
}
func (c *HandlerErrors) AddValueValidator(validator func(string) error) *HandlerErrors {
	c.Validators = append(c.Validators, validator)
	return c
}

func (c *HandlerErrors) AddArrayValueValidators(validators []func(string) error) *HandlerErrors {
	c.Validators = append(c.Validators, validators...)
	return c
}
func (c *HandlerErrors) ValidateValue() error {
	for _, validator := range c.Validators {
		if err := validator(c.ValueModel); err != nil {
			return err
		}
	}
	return nil
}
