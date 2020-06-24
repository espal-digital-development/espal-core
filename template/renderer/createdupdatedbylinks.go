package renderer

import (
	"strings"
)

// CreatedUpdatedByLinks will render the shared createdBy/updatedBy links.
func (templateRenderer *TemplateRenderer) CreatedUpdatedByLinks(ctx context, languageID uint16, entity entity) string {
	if entity.ID() == "" {
		return ""
	}

	out := strings.Builder{}

	templateRenderer.perror(out.WriteString(`<p>`))
	templateRenderer.perror(out.WriteString(`<span>`))

	if entity.IsUpdated() {
		templateRenderer.perror(out.WriteString(ctx.Translate("updatedBy")))
	} else {
		templateRenderer.perror(out.WriteString(ctx.Translate("createdBy")))
	}

	templateRenderer.perror(out.WriteString(`: </span><a href="`))
	templateRenderer.perror(out.WriteString(ctx.AdminURL()))
	templateRenderer.perror(out.WriteString(`/User/View?id=`))

	if entity.IsUpdated() {
		templateRenderer.perror(out.WriteString(*entity.UpdatedByID()))
		templateRenderer.perror(out.WriteString(`">`))
		templateRenderer.perror(out.WriteString(templateRenderer.UpdatedBy(entity, languageID)))
	} else {
		templateRenderer.perror(out.WriteString(entity.CreatedByID()))
		templateRenderer.perror(out.WriteString(`">`))
		templateRenderer.perror(out.WriteString(templateRenderer.CreatedBy(entity, languageID)))
	}

	templateRenderer.perror(out.WriteString(`</a>`))
	templateRenderer.perror(out.WriteString(`</p>`))

	return out.String()
}
