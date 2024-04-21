package models

type Checkbox struct {
	Id             string
	Name           string
	Checked        bool
	Disabled       bool
	ReadOnly       bool
	Required       bool
	InputCSSClass  string
	For            string
	LabelCSSClass  string
	LabelText      string
	HXPostEndpoint string
	HXSwap         string
}

func NewCheckbox(name string) *Checkbox {
	return &Checkbox{
		Id:             name,
		Name:           name,
		Checked:        false,
		Disabled:       false,
		ReadOnly:       false,
		Required:       false,
		InputCSSClass:  "",
		For:            name,
		LabelCSSClass:  "",
		LabelText:      "",
		HXPostEndpoint: "",
		HXSwap:         "",
	}
}
func (c *Checkbox) SetId(id string) *Checkbox {
	c.Id = id
	return c
}
func (c *Checkbox) SetName(name string) *Checkbox {
	c.Name = name
	return c
}
func (c *Checkbox) SetChecked(checked bool) *Checkbox {
	c.Checked = checked
	return c
}
func (c *Checkbox) SetDisabled(disabled bool) *Checkbox {
	c.Disabled = disabled
	return c
}
func (c *Checkbox) SetReadOnly(readOnly bool) *Checkbox {
	c.ReadOnly = readOnly
	return c
}
func (c *Checkbox) SetRequired(required bool) *Checkbox {
	c.Required = required
	return c
}

func (c *Checkbox) SetInputCSSClassByStatus() *Checkbox {

	if c.Checked {
		c.InputCSSClass = "h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
		return c
	}
	c.InputCSSClass = "h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
	return c

}
func (c *Checkbox) SetLabelCSSClassByDefault() *Checkbox {
	c.LabelCSSClass = "block text-sm font-medium leading-6 text-gray-900 ml-1"
	return c
}
func (c *Checkbox) SetFor(forValue string) *Checkbox {
	c.For = forValue
	return c
}
func (c *Checkbox) SetLabelCSSClass(labelCSSClass string) *Checkbox {
	c.LabelCSSClass = labelCSSClass
	return c
}
func (c *Checkbox) SetLabelText(labelText string) *Checkbox {
	c.LabelText = labelText
	return c
}
func (c *Checkbox) SetHXPostEndpoint(hxPostEndpoint string) *Checkbox {
	c.HXPostEndpoint = hxPostEndpoint
	return c
}
func (c *Checkbox) SetHXSwap(hxSwap string) *Checkbox {
	c.HXSwap = hxSwap
	return c
}
