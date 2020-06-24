package group

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/juju/errors"
)

// Filter filters results based on the given context.
func (groupsStore *GroupsStore) Filter(context filters.QueryReader, language language) (result []*Group, filter filters.Filter, err error) {
	alias := (&Group{}).TableAlias()
	filter = groupsStore.databaseFiltersFactory.NewFilter(context, newGroup())
	filter.AddSelectField(filter.NewSelectField("id")).
		AddSelectField(filter.NewSelectField("createdByID")).
		AddSelectField(filter.NewSelectField("updatedByID")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("cu").SetMapTo("createdByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("cu").SetMapTo("createdBySurname")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("uu").SetMapTo("updatedByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("uu").SetMapTo("updatedBySurname")).
		AddSelectField(filter.NewSelectField("value").SetAlias("t").SetMapTo("localizedName")).
		AddSelectField(filter.NewSelectField("createdAt")).
		AddSelectField(filter.NewSelectField("updatedAt")).
		AddSelectField(filter.NewSelectField("active")).
		AddSelectField(filter.NewSelectField("userRights")).
		AddSelectField(filter.NewSelectField("currencies")).
		AddColumn(filter.NewColumn("id")).
		AddColumn(filter.NewColumn("active")).
		AddColumn(filter.NewColumn("createdBy")).
		AddColumn(filter.NewColumn("createdAt")).
		AddColumn(filter.NewColumn("updatedBy")).
		AddColumn(filter.NewColumn("updatedAt")).
		AddColumn(filter.NewColumn("name")).
		AddColumn(filter.NewColumn("userRight").SetPlural(true)).
		AddColumn(filter.NewColumn("currency").SetPlural(true)).
		AddSearchField(filter.NewSearchField("value").SetTableAlias("t")).
		AddJoinStatement(filter.NewJoin("cu", `LEFT JOIN "User" cu ON cu."id" = `+alias+`."createdByID"`)).
		AddJoinStatement(filter.NewJoin("uu", `LEFT JOIN "User" uu ON cu."id" = `+alias+`."updatedByID"`)).
		AddJoinStatement(filter.NewJoin("t", fmt.Sprintf(
			`LEFT JOIN "UserGroupTranslation" t ON t."userGroupID" = %s."id" AND t."language" = %d AND t."field" = %d`, alias, language.ID(), database.DBTranslationFieldName))).
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
		result = make([]*Group, 0)
		for filter.Rows().Next() {
			if err = filter.Rows().Err(); err != nil {
				err = errors.Trace(err)
				return
			}
			group := newGroup()
			if err = filter.Rows().Scan(&group.id, &group.createdByID, &group.updatedByID,
				&group.createdByFirstName, &group.createdBySurname, &group.updatedByFirstName,
				&group.updatedBySurname, &group.localizedName, &group.createdAt, &group.updatedAt,
				&group.active, &group.userRights, &group.currencies); err != nil {
				err = errors.Trace(err)
				return
			}
			result = append(result, group)
		}
	}

	return
}
