package validators

import (
	"strings"

	"github.com/juju/errors"
)

// TODO :: 777 Escape HTML on all field.Value calls

func (form *Form) renderInputTypeField(field *formField) string {
	out := strings.Builder{}
	form.addLabel(field, &out)
	form.perror(out.WriteString(`<input`))
	if TextFormField != field.Type() {
		form.perror(out.WriteString(` type="`))
		// TODO :: 77 Do this through read-only mapping?
		switch field.Type() {
		case HiddenFormField, TokenFormField, HoneypotFormField:
			form.perror(out.WriteString(`hidden`))
		case NumberFormField:
			form.perror(out.WriteString(`number`))
		case DateTimeFormField:
			form.perror(out.WriteString(`date`))
		case EmailFormField:
			form.perror(out.WriteString(`email`))
		case SearchFormField:
			form.perror(out.WriteString(`search`))
		case PasswordFormField:
			form.perror(out.WriteString(`password`))
		case FileFormField:
			form.perror(out.WriteString(`file`))
		default:
			err := errors.Errorf("unknown type `%v` to render input for", field.Type())
			form.loggerService.Error(err.Error())
			panic(err)
		}
		form.perror(out.WriteString(`"`))
	}
	form.perror(out.WriteString(` name="`))
	form.perror(out.WriteString(field.Name()))
	if FileFormField == field.Type() && field.Multiple() {
		// TODO :: 777 Need to expand this to count an extra entry every time
		//         a new File field is rendered.
		form.perror(out.WriteString(`[0]`))
	}
	form.perror(out.WriteString(`"`))
	if field.Placeholder() != "" {
		form.perror(out.WriteString(` placeholder="`))
		form.perror(out.WriteString(field.Placeholder()))
		if DateTimeFormField == field.Type() {
			form.perror(out.WriteString(` (`))
			if !field.ExcludeYear() {
				form.perror(out.WriteString(`YYYY`))
			}
			if !field.ExcludeYear() && !field.ExcludeMonth() {
				form.perror(out.WriteString(`-`))
			}
			if field.ExcludeMonth() {
				form.perror(out.WriteString(`MM`))
			}
			if !field.ExcludeYear() && !field.ExcludeMonth() && !field.ExcludeDay() {
				form.perror(out.WriteString(`-`))
			}
			if field.ExcludeDay() {
				form.perror(out.WriteString(`DD`))
			}
			form.perror(out.WriteString(`)`))
		}
		form.perror(out.WriteString(`"`))
	}
	if field.Value() != "" {
		form.perror(out.WriteString(` value="`))
		form.perror(out.WriteString(field.Value()))
		form.perror(out.WriteString(`"`))
	}
	if !field.Optional() && HiddenFormField != field.Type() && HoneypotFormField != field.Type() && TokenFormField != field.Type() {
		form.perror(out.WriteString(` required`))
	}
	form.perror(out.WriteString(`>`))
	return out.String()
}
