package validators

import (
	"io"
	"strings"

	"github.com/juju/errors"
)

// RenderErrors will render the list of form errors.
// Returns an empty string when no errors are present.
func (f *Form) RenderErrors() string {
	if f.isValid {
		return ""
	}
	errorList := strings.Builder{}
	errorList.WriteString(`<ul class="errors">`)
	for k := range f.submitErrors {
		errorList.WriteString("<li>")
		errorList.WriteString(f.submitErrors[k])
		errorList.WriteString("</li>")
	}
	for k := range f.fields {
		if f.fields[k].HasErrors() {
			for _, errorMessage := range f.fields[k].Errors() {
				errorList.WriteString("<li>")
				errorList.WriteString(errorMessage)
				errorList.WriteString("</li>")
			}
		}
	}
	errorList.WriteString("</ul>")
	return errorList.String()
}

// RenderOpen will render the form open tag.
func (f *Form) RenderOpen() string {
	if f.isMultipart() {
		return `<form method="post" enctype="multipart/form-data">`
	}
	return `<form method="post">`
}

// RenderField will render the form field and resolve it's rules and presets.
func (f *Form) RenderField(name string) string {
	field := f.field(name)
	switch field.Type() {
	case HiddenFormField, TokenFormField, HoneypotFormField, TextFormField, NumberFormField, DateTimeFormField,
		EmailFormField, SearchFormField, PasswordFormField, FileFormField:
		return f.renderInputTypeField(field)
	case TextAreaFormField:
		return f.renderTextAreaField(field)
	case CheckboxFormField:
		return f.renderCheckboxField(field)
	case ChoiceFormField:
		return f.renderChoiceField(field)
	default:
		err := errors.Errorf("invalid field type `%v`", field.Type())
		f.loggerService.Error(err.Error())
		panic(err)
	}
}

func (f *Form) addLabel(field FormField, stringBuilder io.StringWriter) {
	if field.HideLabel() {
		return
	}
	stringBuilder.WriteString(`<label for="`)
	stringBuilder.WriteString(field.Name())
	stringBuilder.WriteString(`">`)
	stringBuilder.WriteString(field.Placeholder())
	stringBuilder.WriteString(`</label>`)
}
