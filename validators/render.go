package validators

import (
	"io"
	"strings"

	"github.com/juju/errors"
)

// RenderErrors will render the list of form errors.
// Returns an empty string when no errors are present.
func (form *Form) RenderErrors() string {
	if form.isValid {
		return ""
	}
	errorList := strings.Builder{}
	form.perror(errorList.WriteString(`<ul class="errors">`))
	for k := range form.submitErrors {
		form.perror(errorList.WriteString("<li>"))
		form.perror(errorList.WriteString(form.submitErrors[k]))
		form.perror(errorList.WriteString("</li>"))
	}
	for k := range form.fields {
		if form.fields[k].HasErrors() {
			for _, errorMessage := range form.fields[k].Errors() {
				form.perror(errorList.WriteString("<li>"))
				form.perror(errorList.WriteString(errorMessage))
				form.perror(errorList.WriteString("</li>"))
			}
		}
	}
	form.perror(errorList.WriteString("</ul>"))
	return errorList.String()
}

// RenderOpen will render the form open tag.
func (form *Form) RenderOpen() string {
	if form.isMultipart() {
		return `<form method="post" enctype="multipart/form-data">`
	}
	return `<form method="post">`
}

// RenderField will render the form field and resolve it's rules and presets.
func (form *Form) RenderField(name string) string {
	field := form.field(name)
	switch field.Type() {
	case HiddenFormField, TokenFormField, HoneypotFormField, TextFormField, NumberFormField, DateTimeFormField, EmailFormField, SearchFormField, PasswordFormField, FileFormField:
		return form.renderInputTypeField(field)
	case TextAreaFormField:
		return form.renderTextAreaField(field)
	case CheckboxFormField:
		return form.renderCheckboxField(field)
	case ChoiceFormField:
		return form.renderChoiceField(field)
	default:
		err := errors.Errorf("invalid field type `%v`", field.Type())
		form.loggerService.Error(err.Error())
		panic(err)
	}
}

func (form *Form) addLabel(field FormField, stringBuilder io.StringWriter) {
	if field.HideLabel() {
		return
	}
	form.perror(stringBuilder.WriteString(`<label for="`))
	form.perror(stringBuilder.WriteString(field.Name()))
	form.perror(stringBuilder.WriteString(`">`))
	form.perror(stringBuilder.WriteString(field.Placeholder()))
	form.perror(stringBuilder.WriteString(`</label>`))
}
