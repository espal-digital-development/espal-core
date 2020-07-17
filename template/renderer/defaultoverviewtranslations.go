package renderer

import (
	"strings"
)

// DefaultOverviewTranslations will render the shared JavaScript translations.
func (r *TemplateRenderer) DefaultOverviewTranslations(ctx context) string {
	out := strings.Builder{}

	out.WriteString(`<script>const trans = {`)

	out.WriteString(`'nothingSelected': '`)
	out.WriteString(ctx.Translate("nothingSelected"))
	out.WriteString(`',`)

	out.WriteString(`'areYouSure': '`)
	out.WriteString(ctx.Translate("areYouSure"))
	out.WriteString(`',`)

	out.WriteString(`};</script>`)

	return out.String()
}
