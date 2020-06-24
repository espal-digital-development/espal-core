package renderer

import (
	"strings"
)

// DefaultOverviewTranslations will render the shared JavaScript translations.
func (templateRenderer *TemplateRenderer) DefaultOverviewTranslations(ctx context) string {
	out := strings.Builder{}

	templateRenderer.perror(out.WriteString(`<script>const trans = {`))

	templateRenderer.perror(out.WriteString(`'nothingSelected': '`))
	templateRenderer.perror(out.WriteString(ctx.Translate("nothingSelected")))
	templateRenderer.perror(out.WriteString(`',`))

	templateRenderer.perror(out.WriteString(`'areYouSure': '`))
	templateRenderer.perror(out.WriteString(ctx.Translate("areYouSure")))
	templateRenderer.perror(out.WriteString(`',`))

	templateRenderer.perror(out.WriteString(`};</script>`))

	return out.String()
}
