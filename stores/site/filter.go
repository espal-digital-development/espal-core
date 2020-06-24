package site

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/juju/errors"
)

// Filter filters results based on the given context.
func (sitesStore *SitesStore) Filter(context filters.QueryReader, language language) (sites []*Site, filter filters.Filter, err error) {
	alias := (&Site{}).TableAlias()
	filter = sitesStore.databaseFiltersFactory.NewFilter(context, newSite())
	filter.AddSelectField(filter.NewSelectField("id")).
		AddSelectField(filter.NewSelectField("createdByID")).
		AddSelectField(filter.NewSelectField("updatedByID")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("cu").SetMapTo("createdByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("cu").SetMapTo("createdBySurname")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("uu").SetMapTo("updatedByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("uu").SetMapTo("updatedBySurname")).
		AddSelectField(filter.NewSelectField("value").SetAlias("t").SetMapTo("localizedName")).
		AddSelectField(filter.NewSelectField("online")).
		AddSelectField(filter.NewSelectField("createdAt")).
		AddSelectField(filter.NewSelectField("updatedAt")).
		AddSelectField(filter.NewSelectField("language")).
		AddSelectField(filter.NewSelectField("country")).
		AddSelectField(filter.NewSelectField("currencies")).
		AddColumn(filter.NewColumn("id")).
		AddColumn(filter.NewColumn("online")).
		AddColumn(filter.NewColumn("createdBy")).
		AddColumn(filter.NewColumn("createdAt")).
		AddColumn(filter.NewColumn("updatedBy")).
		AddColumn(filter.NewColumn("updatedAt")).
		AddColumn(filter.NewColumn("name")).
		AddColumn(filter.NewColumn("language")).
		AddColumn(filter.NewColumn("country")).
		AddColumn(filter.NewColumn("currency").SetPlural(true)).
		AddSearchField(filter.NewSearchField("value").SetTableAlias("t")).
		AddJoinStatement(filter.NewJoin("cu", `LEFT JOIN "User" cu ON cu."id" = `+alias+`."createdByID"`)).
		AddJoinStatement(filter.NewJoin("uu", `LEFT JOIN "User" uu ON cu."id" = `+alias+`."updatedByID"`)).
		AddJoinStatement(filter.NewJoin("t", fmt.Sprintf(`LEFT JOIN "SiteTranslation" t ON t."siteID" = %s."id" AND t."language" = %d AND t."field" = %d`, alias, language.ID(), database.DBTranslationFieldName))).
		AddSortField(filter.NewSortField("id", true).SetTableAlias(alias))

	if err = filter.Process(); err != nil {
		err = errors.Trace(err)
		return
	}
	if filter.HasResults() {
		defer func(dbRows database.Rows) {
			closeErr := dbRows.Close()
			if err != nil && closeErr != nil {
				err = errors.Wrap(err, closeErr)
			} else if closeErr != nil {
				err = errors.Trace(closeErr)
			}
		}(filter.Rows())
		sites = make([]*Site, 0)
		for filter.Rows().Next() {
			if err = filter.Rows().Err(); err != nil {
				err = errors.Trace(err)
				return
			}
			site := newSite()
			if err = filter.Rows().Scan(&site.id, &site.createdByID, &site.updatedByID, &site.createdByFirstName, &site.createdBySurname, &site.updatedByFirstName, &site.updatedBySurname, &site.localizedName, &site.online, &site.createdAt, &site.updatedAt, &site.language, &site.country, &site.currencies); err != nil {
				err = errors.Trace(err)
				return
			}
			sites = append(sites, site)
		}
	}

	return
}
