package user

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/juju/errors"
)

// Filter filters results based on the given context.
// nolint:nakedret
func (s *UsersStore) Filter(context filters.QueryReader) (result []*User, filter filters.Filter, err error) {
	alias := (&User{}).TableAlias()
	filter = s.databaseFiltersFactory.NewFilter(context, newUser())
	filter.AddSelectField(filter.NewSelectField("id")).
		AddSelectField(filter.NewSelectField("createdByID")).
		AddSelectField(filter.NewSelectField("updatedByID")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("cu").SetMapTo("createdByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("cu").SetMapTo("createdBySurname")).
		AddSelectField(filter.NewSelectField("firstName").SetAlias("uu").SetMapTo("updatedByFirstName")).
		AddSelectField(filter.NewSelectField("surname").SetAlias("uu").SetMapTo("updatedBySurname")).
		AddSelectField(filter.NewSelectField("active")).
		AddSelectField(filter.NewSelectField("createdAt")).
		AddSelectField(filter.NewSelectField("updatedAt")).
		AddSelectField(filter.NewSelectField("firstName")).
		AddSelectField(filter.NewSelectField("surname")).
		AddSelectField(filter.NewSelectField("email")).
		AddSelectField(filter.NewSelectField("dateOfBirth")).
		AddSelectField(filter.NewSelectField("language")).
		AddSelectField(filter.NewSelectField("country")).
		AddSelectField(filter.NewSelectField("currencies")).
		AddColumn(filter.NewColumn("id")).
		AddColumn(filter.NewColumn("active")).
		AddColumn(filter.NewColumn("createdBy")).
		AddColumn(filter.NewColumn("createdAt")).
		AddColumn(filter.NewColumn("updatedBy")).
		AddColumn(filter.NewColumn("updatedAt")).
		AddColumn(filter.NewColumn("firstName")).
		AddColumn(filter.NewColumn("surname")).
		AddColumn(filter.NewColumn("email")).
		AddColumn(filter.NewColumn("language")).
		AddColumn(filter.NewColumn("country")).
		AddColumn(filter.NewColumn("dateOfBirth")).
		AddColumn(filter.NewColumn("currency").SetPlural(true)).
		// TODO :: What if one would want to search Language/Country/Currency?
		AddSearchField(filter.NewSearchField("email").SetTableAlias(alias)).
		AddSearchField(filter.NewSearchField("firstName").SetTableAlias(alias)).
		AddSearchField(filter.NewSearchField("surname").SetTableAlias(alias)).
		AddJoinStatement(filter.NewJoin("cu", `LEFT JOIN "User" cu ON cu."id" = `+alias+`."createdByID"`)).
		AddJoinStatement(filter.NewJoin("uu", `LEFT JOIN "User" uu ON uu."id" = `+alias+`."updatedByID"`)).
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
		result = make([]*User, 0)
		for filter.Rows().Next() {
			if err = filter.Rows().Err(); err != nil {
				err = errors.Trace(err)
				return
			}
			user := newUser()
			if err = filter.Rows().Scan(&user.id, &user.createdByID, &user.updatedByID, &user.createdByFirstName,
				&user.createdBySurname, &user.updatedByFirstName, &user.updatedBySurname, &user.active,
				&user.createdAt, &user.updatedAt, &user.firstName, &user.surname, &user.email, &user.dateOfBirth,
				&user.language, &user.country, &user.currencies); err != nil {
				err = errors.Trace(err)
				return
			}
			result = append(result, user)
		}
	}

	return
}
