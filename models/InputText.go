package models

/*

<div>

		<label for="email" class="block text-sm font-medium leading-6 text-gray-900">Email address</label>
		<div class="mt-2">
			<input id="email" name="email" type="email" autocomplete="email" required class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
		</div>
	</div>
*/

type InputTextModel struct {
	IdWrap       string
	HXTarget     string
	ErrorMessage string
	Value        string
	HxTrigger    string
	HXEndpoint   string
	For          string
	ClassLabel   string
	LabelText    string
	Id           string
	Name         string
	Type         string
	Autocomplete string
	Required     bool
	Disabled     bool
	ClassInput   string
	HXSwap       string
}

func NewInputTextModel(name string) *InputTextModel {
	return &InputTextModel{
		IdWrap:       "wrap-" + name,
		HXTarget:     "#wrap-" + name,
		HXEndpoint:   "",
		HxTrigger:    "load",
		HXSwap:       "outerHTML",
		For:          name,
		ClassLabel:   "block text-sm font-medium leading-6 text-gray-900",
		LabelText:    name,
		Id:           name,
		Name:         name,
		Type:         "text",
		Autocomplete: "",
		Required:     false,
		Disabled:     false,
		ClassInput:   "block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
	}
}

func (i *InputTextModel) SetHXTrigger(hxTrigger string) {
	i.HxTrigger = hxTrigger
}
func (i *InputTextModel) SetValue(value string) {
	i.Value = value
}
func (i *InputTextModel) SetHXSwap(hxSwap string) {
	i.HXSwap = hxSwap
}

func (i *InputTextModel) SetId(id string) {
	i.Id = id
}
func (i *InputTextModel) SetName(name string) {
	i.Name = name
}
func (i *InputTextModel) SetFor(forName string) {
	i.For = forName
}
func (i *InputTextModel) SetHXEndpoint(hxEndpoint string) {
	i.HXEndpoint = hxEndpoint
}

func (i *InputTextModel) SetLabelText(labelText string) {
	i.LabelText = labelText
}

func (i *InputTextModel) SetType(inputType string) {
	i.Type = inputType
}

func (i *InputTextModel) SetAutocomplete(autocomplete string) {
	i.Autocomplete = autocomplete
}

func (i *InputTextModel) SetRequired(required bool) {
	i.Required = required
}

func (i *InputTextModel) SetClassInput(classInput string) {
	i.ClassInput = classInput
}
func (i *InputTextModel) SetDisabled(statusDisable bool) {
	i.Disabled = statusDisable
}
func (i *InputTextModel) SetErrorMessage(errorMessage string) {
	i.ErrorMessage = errorMessage
}
func (i *InputTextModel) SetClassInputTextByStatus(status string) *InputTextModel {
	switch status {
	case "success":
		i.ClassInput = "block w-full rounded-md border-0 py-1.5 px-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-green-500 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-green-500 sm:text-sm sm:leading-6"
	case "error":
		i.ClassInput = "block w-full rounded-md border-0 py-1.5 px-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-red-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
	default:
		i.ClassInput = "block w-full rounded-md border-0 py-1.5 px-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
	}
	return i
}
func (i *InputTextModel) SetClassLabelByStatus(status string) *InputTextModel {
	switch status {
	case "success":
		i.ClassLabel = "block text-sm font-medium leading-6 text-green-900"

	case "error":
		i.ClassLabel = "block text-sm font-medium leading-6 text-red-900"

	default:
		i.ClassLabel = "block text-sm font-medium leading-6 text-gray-900"

	}
	return i
}
