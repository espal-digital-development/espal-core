package validators

import (
	"strings"
)

func (f *Form) renderTextAreaField(field FormField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)
	f.perror(out.WriteString(`<textarea name="`))
	f.perror(out.WriteString(field.Name()))
	f.perror(out.WriteString(`"`))
	if field.Placeholder() != "" {
		f.perror(out.WriteString(` placeholder="`))
		f.perror(out.WriteString(field.Placeholder()))
		f.perror(out.WriteString(`"`))
	}
	if !field.Optional() {
		f.perror(out.WriteString(` required`))
	}
	f.perror(out.WriteString(`>`))
	if field.Value() != "" {
		f.perror(out.WriteString(field.Value()))
	}
	f.perror(out.WriteString(`</textarea>`))
	return out.String()
}
