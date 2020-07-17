package validators

import (
	"strings"

	"github.com/juju/errors"
)

// TODO :: 777 Escape HTML on all field.Value calls

func (f *Form) renderInputTypeField(field *formField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)
	out.WriteString(`<input`)
	if TextFormField != field.Type() {
		out.WriteString(` type="`)
		// TODO :: 77 Do this through read-only mapping?
		switch field.Type() {
		case HiddenFormField, TokenFormField, HoneypotFormField:
			out.WriteString(`hidden`)
		case NumberFormField:
			out.WriteString(`number`)
		case DateTimeFormField:
			out.WriteString(`date`)
		case EmailFormField:
			out.WriteString(`email`)
		case SearchFormField:
			out.WriteString(`search`)
		case PasswordFormField:
			out.WriteString(`password`)
		case FileFormField:
			out.WriteString(`file`)
		default:
			err := errors.Errorf("unknown type `%v` to render input for", field.Type())
			f.loggerService.Error(err.Error())
			panic(err)
		}
		out.WriteString(`"`)
	}
	out.WriteString(` name="`)
	out.WriteString(field.Name())
	if FileFormField == field.Type() && field.Multiple() {
		// TODO :: 777 Need to expand this to count an extra entry every time
		//         a new File field is rendered.
		out.WriteString(`[0]`)
	}
	out.WriteString(`"`)
	if field.Placeholder() != "" {
		out.WriteString(` placeholder="`)
		out.WriteString(field.Placeholder())
		if DateTimeFormField == field.Type() {
			out.WriteString(` (`)
			if !field.ExcludeYear() {
				out.WriteString(`YYYY`)
			}
			if !field.ExcludeYear() && !field.ExcludeMonth() {
				out.WriteString(`-`)
			}
			if field.ExcludeMonth() {
				out.WriteString(`MM`)
			}
			if !field.ExcludeYear() && !field.ExcludeMonth() && !field.ExcludeDay() {
				out.WriteString(`-`)
			}
			if field.ExcludeDay() {
				out.WriteString(`DD`)
			}
			out.WriteString(`)`)
		}
		out.WriteString(`"`)
	}
	if field.Value() != "" {
		out.WriteString(` value="`)
		out.WriteString(field.Value())
		out.WriteString(`"`)
	}
	if !field.Optional() && HiddenFormField != field.Type() && HoneypotFormField != field.Type() &&
		TokenFormField != field.Type() {
		out.WriteString(` required`)
	}
	out.WriteString(`>`)
	return out.String()
}
