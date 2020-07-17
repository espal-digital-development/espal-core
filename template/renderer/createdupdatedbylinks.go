package renderer

import (
	"strings"
)

// CreatedUpdatedByLinks will render the shared createdBy/updatedBy links.
func (r *TemplateRenderer) CreatedUpdatedByLinks(ctx context, languageID uint16, entity entity) string {
	if entity.ID() == "" {
		return ""
	}

	out := strings.Builder{}

	out.WriteString(`<p>`)
	out.WriteString(`<span>`)

	if entity.IsUpdated() {
		out.WriteString(ctx.Translate("updatedBy"))
	} else {
		out.WriteString(ctx.Translate("createdBy"))
	}

	out.WriteString(`: </span><a href="`)
	out.WriteString(ctx.AdminURL())
	out.WriteString(`/User/View?id=`)

	if entity.IsUpdated() {
		out.WriteString(*entity.UpdatedByID())
		out.WriteString(`">`)
		out.WriteString(r.UpdatedBy(entity, languageID))
	} else {
		out.WriteString(entity.CreatedByID())
		out.WriteString(`">`)
		out.WriteString(r.CreatedBy(entity, languageID))
	}

	out.WriteString(`</a>`)
	out.WriteString(`</p>`)

	return out.String()
}
