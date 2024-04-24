package models

/*
<div>

		<label for="email" class="block text-sm font-medium leading-6 text-gray-900">Email address</label>
		<div class="mt-2">
			<input id="email" name="email" type="email" autocomplete="email" required class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"/>
		</div>
	</div>
*/

func isInList(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

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
	Validator    *HandlerValidator
}

func NewInputTextModel(name string) *InputTextModel {
	thisModel := &InputTextModel{
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
		Value:        "",
	}
	thisModel.Validator = NewHandlerValidators(&thisModel.Value)
	return thisModel

}

func (i *InputTextModel) SetHXTrigger(hxTrigger string) *InputTextModel {
	i.HxTrigger = hxTrigger
	return i

}
func (i *InputTextModel) SetValue(value string) *InputTextModel {
	i.Validator.ValueModel = value
	i.Value = value
	return i
}
func (i *InputTextModel) SetHXSwap(hxSwap string) *InputTextModel {
	i.HXSwap = hxSwap
	return i
}

func (i *InputTextModel) SetId(id string) *InputTextModel {
	i.Id = id
	return i
}
func (i *InputTextModel) SetName(name string) *InputTextModel {
	i.Name = name
	return i
}
func (i *InputTextModel) SetFor(forName string) *InputTextModel {
	i.For = forName
	return i
}
func (i *InputTextModel) SetHXEndpoint(hxEndpoint string) *InputTextModel {
	i.HXEndpoint = hxEndpoint
	return i
}

func (i *InputTextModel) SetLabelText(labelText string) *InputTextModel {
	i.LabelText = labelText
	return i
}

func (i *InputTextModel) SetType(inputType string) *InputTextModel {
	i.Type = inputType
	return i
}

func (i *InputTextModel) SetAutocomplete(autocomplete string) *InputTextModel {
	i.Autocomplete = autocomplete
	return i
}

func (i *InputTextModel) SetRequired(required bool) *InputTextModel {
	i.Required = required
	return i
}

func (i *InputTextModel) SetClassInput(classInput string) *InputTextModel {
	i.ClassInput = classInput
	return i
}
func (i *InputTextModel) SetDisabled(statusDisable bool) *InputTextModel {
	i.Disabled = statusDisable
	return i
}
func (i *InputTextModel) SetErrorMessage(errorMessage string) *InputTextModel {
	i.ErrorMessage = errorMessage
	return i
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

func (i *InputTextModel) SetArrayValidators(arrayValidatorError []func(string) error) *InputTextModel {
	i.Validator.AddArrayFuncsValidators(arrayValidatorError)
	return i
}

func (i *InputTextModel) SetValidatorError(validatorError func(string) error) *InputTextModel {
	i.Validator.AddFuncValidator(validatorError)
	return i
}

func (i *InputTextModel) SetStatusByString(status, errorMessage string) *InputTextModel {
	isValidStatus := isInList([]string{"success", "error", "default"}, status)
	if !isValidStatus {
		panic("Invalid status")
	}
	i.SetErrorMessage(errorMessage).SetClassLabelByStatus(status).SetClassInputTextByStatus(status)
	return i
}

func (i *InputTextModel) ValidateNow() *InputTextModel {
	err := i.Validator.ValidateValue()
	if err != nil {
		i.SetStatusByString("error", err.Error())
		return i
	}
	i.SetStatusByString("success", "")
	return i

}
