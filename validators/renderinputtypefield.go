package validators

import (
	"strings"

	"github.com/juju/errors"
)

// TODO :: 777 Escape HTML on all field.Value calls

func (f *Form) renderInputTypeField(field *formField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)
	f.perror(out.WriteString(`<input`))
	if TextFormField != field.Type() {
		f.perror(out.WriteString(` type="`))
		// TODO :: 77 Do this through read-only mapping?
		switch field.Type() {
		case HiddenFormField, TokenFormField, HoneypotFormField:
			f.perror(out.WriteString(`hidden`))
		case NumberFormField:
			f.perror(out.WriteString(`number`))
		case DateTimeFormField:
			f.perror(out.WriteString(`date`))
		case EmailFormField:
			f.perror(out.WriteString(`email`))
		case SearchFormField:
			f.perror(out.WriteString(`search`))
		case PasswordFormField:
			f.perror(out.WriteString(`password`))
		case FileFormField:
			f.perror(out.WriteString(`file`))
		default:
			err := errors.Errorf("unknown type `%v` to render input for", field.Type())
			f.loggerService.Error(err.Error())
			panic(err)
		}
		f.perror(out.WriteString(`"`))
	}
	f.perror(out.WriteString(` name="`))
	f.perror(out.WriteString(field.Name()))
	if FileFormField == field.Type() && field.Multiple() {
		// TODO :: 777 Need to expand this to count an extra entry every time
		//         a new File field is rendered.
		f.perror(out.WriteString(`[0]`))
	}
	f.perror(out.WriteString(`"`))
	if field.Placeholder() != "" {
		f.perror(out.WriteString(` placeholder="`))
		f.perror(out.WriteString(field.Placeholder()))
		if DateTimeFormField == field.Type() {
			f.perror(out.WriteString(` (`))
			if !field.ExcludeYear() {
				f.perror(out.WriteString(`YYYY`))
			}
			if !field.ExcludeYear() && !field.ExcludeMonth() {
				f.perror(out.WriteString(`-`))
			}
			if field.ExcludeMonth() {
				f.perror(out.WriteString(`MM`))
			}
			if !field.ExcludeYear() && !field.ExcludeMonth() && !field.ExcludeDay() {
				f.perror(out.WriteString(`-`))
			}
			if field.ExcludeDay() {
				f.perror(out.WriteString(`DD`))
			}
			f.perror(out.WriteString(`)`))
		}
		f.perror(out.WriteString(`"`))
	}
	if field.Value() != "" {
		f.perror(out.WriteString(` value="`))
		f.perror(out.WriteString(field.Value()))
		f.perror(out.WriteString(`"`))
	}
	if !field.Optional() && HiddenFormField != field.Type() && HoneypotFormField != field.Type() &&
		TokenFormField != field.Type() {
		f.perror(out.WriteString(` required`))
	}
	f.perror(out.WriteString(`>`))
	return out.String()
}
