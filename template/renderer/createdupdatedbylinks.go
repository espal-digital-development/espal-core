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

	r.perror(out.WriteString(`<p>`))
	r.perror(out.WriteString(`<span>`))

	if entity.IsUpdated() {
		r.perror(out.WriteString(ctx.Translate("updatedBy")))
	} else {
		r.perror(out.WriteString(ctx.Translate("createdBy")))
	}

	r.perror(out.WriteString(`: </span><a href="`))
	r.perror(out.WriteString(ctx.AdminURL()))
	r.perror(out.WriteString(`/User/View?id=`))

	if entity.IsUpdated() {
		r.perror(out.WriteString(*entity.UpdatedByID()))
		r.perror(out.WriteString(`">`))
		r.perror(out.WriteString(r.UpdatedBy(entity, languageID)))
	} else {
		r.perror(out.WriteString(entity.CreatedByID()))
		r.perror(out.WriteString(`">`))
		r.perror(out.WriteString(r.CreatedBy(entity, languageID)))
	}

	r.perror(out.WriteString(`</a>`))
	r.perror(out.WriteString(`</p>`))

	return out.String()
}
