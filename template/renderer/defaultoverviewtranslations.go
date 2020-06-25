package renderer

import (
	"strings"
)

// DefaultOverviewTranslations will render the shared JavaScript translations.
func (r *TemplateRenderer) DefaultOverviewTranslations(ctx context) string {
	out := strings.Builder{}

	r.perror(out.WriteString(`<script>const trans = {`))

	r.perror(out.WriteString(`'nothingSelected': '`))
	r.perror(out.WriteString(ctx.Translate("nothingSelected")))
	r.perror(out.WriteString(`',`))

	r.perror(out.WriteString(`'areYouSure': '`))
	r.perror(out.WriteString(ctx.Translate("areYouSure")))
	r.perror(out.WriteString(`',`))

	r.perror(out.WriteString(`};</script>`))

	return out.String()
}
