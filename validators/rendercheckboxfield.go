package validators

import (
	"strings"
)

func (f *Form) renderCheckboxField(field FormField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)
	f.perror(out.WriteString(`<p><input type="checkbox" name="`))
	f.perror(out.WriteString(field.Name()))
	f.perror(out.WriteString(`"`))
	if field.Value() == "1" {
		f.perror(out.WriteString(` checked`))
	}
	f.perror(out.WriteString(`>`))
	if field.Placeholder() != "" {
		f.perror(out.WriteString(`<span> `))
		f.perror(out.WriteString(field.Placeholder()))
		f.perror(out.WriteString(`</span>`))
	}
	f.perror(out.WriteString(`</p>`))
	return out.String()
}
