package site

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/juju/errors"
)

// Search searches results based on the given context through the filter mechanics.
func (sitesStore *SitesStore) Search(context filters.QueryReader, language language) (sites []*Site, filter filters.Filter, err error) {
	alias := (&Site{}).TableAlias()
	filter = sitesStore.databaseFiltersFactory.NewFilter(context, newSite())
	filter.AddSelectField(filter.NewSelectField("id")).
		AddSelectField(filter.NewSelectField("value").SetAlias("t").SetMapTo("localizedName")).
		AddSearchField(filter.NewSearchField("value").SetTableAlias("t")).
		AddJoinStatement(filter.NewJoin("t", fmt.Sprintf(`LEFT JOIN "SiteTranslation" t ON t."siteID" = %s."id" AND t."language" = %d AND t."field" = %d`, alias, language.ID(), database.DBTranslationFieldName)))

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
