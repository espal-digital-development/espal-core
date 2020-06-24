package domain

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/juju/errors"
)

// Filter filters results based on the given context.
func (domainsStore *DomainsStore) Filter(context filters.QueryReader) (result []*Domain, filter filters.Filter, err error) {
	alias := (&Domain{}).TableAlias()
	filter = domainsStore.databaseFiltersFactory.NewFilter(context, newDomain())
	filter.AddSelectField(filter.NewSelectField("id")).
		AddSelectField(filter.NewSelectField("createdByID")).
		AddSelectField(filter.NewSelectField("updatedByID")).
		AddSelectField(filter.NewSelectField("siteID")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("cu").SetMapTo("createdByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("cu").SetMapTo("createdBySurname")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("uu").SetMapTo("updatedByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("uu").SetMapTo("updatedBySurname")).
		AddSelectField(filter.NewSelectField("active")).
		AddSelectField(filter.NewSelectField("createdAt")).
		AddSelectField(filter.NewSelectField("updatedAt")).
		AddSelectField(filter.NewSelectField("host")).
		AddSelectField(filter.NewSelectField("language")).
		AddSelectField(filter.NewSelectField("currencies")).
		AddColumn(filter.NewColumn("id")).
		AddColumn(filter.NewColumn("active")).
		AddColumn(filter.NewColumn("createdBy")).
		AddColumn(filter.NewColumn("createdAt")).
		AddColumn(filter.NewColumn("updatedBy")).
		AddColumn(filter.NewColumn("updatedAt")).
		AddColumn(filter.NewColumn("host")).
		AddColumn(filter.NewColumn("language")).
		AddColumn(filter.NewColumn("currency").SetPlural(true)).
		AddSearchField(filter.NewSearchField("host").SetTableAlias(alias)).
		AddJoinStatement(filter.NewJoin("cu", `LEFT JOIN "User" cu ON cu."id" = `+alias+`."createdByID"`)).
		AddJoinStatement(filter.NewJoin("uu", `LEFT JOIN "User" uu ON cu."id" = `+alias+`."updatedByID"`)).
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
		result = make([]*Domain, 0)
		for filter.Rows().Next() {
			if err = filter.Rows().Err(); err != nil {
				err = errors.Trace(err)
				return
			}
			domain := newDomain()
			if err = filter.Rows().Scan(&domain.id, &domain.createdByID, &domain.updatedByID, &domain.siteID, &domain.createdByFirstName, &domain.createdBySurname, &domain.updatedByFirstName, &domain.updatedBySurname, &domain.active, &domain.createdAt, &domain.updatedAt, &domain.host, &domain.language, &domain.currencies); err != nil {
				err = errors.Trace(err)
				return
			}
			result = append(result, domain)
		}
	}

	return
}
