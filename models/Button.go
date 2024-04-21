package models

type Button struct {
	Id             string
	LabelText      string
	Disabled       bool
	HXPostEndpoint string
	HXTrigger      string
	HXSwap         string
	CSSClassButton string
}

func NewButton(id, labelText string) *Button {
	return &Button{
		Id:        id,
		LabelText: labelText,
		Disabled:  false,
	}
}
func (b *Button) SetId(id string) *Button {
	b.Id = id
	return b
}

// dev all set methods?
func (b *Button) SetHXSwap(swap string) *Button {
	b.HXSwap = swap
	return b
}

func (b *Button) SetHXTrigger(trigger string) *Button {
	b.HXTrigger = trigger
	return b
}

func (b *Button) SetCSSClassButton(cssClass string) *Button {
	b.CSSClassButton = cssClass
	return b
}

func (b *Button) SetHXPostEndpoint(endpoint string) *Button {
	b.HXPostEndpoint = endpoint
	return b
}

func (b *Button) SetDisabled(status bool) *Button {
	b.Disabled = status
	return b
}

func (b *Button) SetLabelText(text string) *Button {
	b.LabelText = text
	return b
}
