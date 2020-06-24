package user

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
	"github.com/juju/errors"
)

// Search searches results based on the given context through the filter mechanics.
func (usersStore *UsersStore) Search(context filters.QueryReader) (result []*User, filter filters.Filter, err error) {
	alias := (&User{}).TableAlias()
	filter = usersStore.databaseFiltersFactory.NewFilter(context, newUser())
	filter.AddSelectField(filter.NewSelectField("id")).
		AddSelectField(filter.NewSelectField("firstName")).
		AddSelectField(filter.NewSelectField("surname")).
		AddSelectField(filter.NewSelectField("email")).
		AddSearchField(filter.NewSearchField("firstName").SetTableAlias(alias)).
		AddSearchField(filter.NewSearchField("surname").SetTableAlias(alias)).
		AddSearchField(filter.NewSearchField("email").SetTableAlias(alias))

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
			if err = filter.Rows().Scan(&user.id, &user.createdByID, &user.updatedByID, &user.createdByFirstName, &user.createdBySurname, &user.updatedByFirstName, &user.updatedBySurname, &user.active, &user.createdAt, &user.updatedAt, &user.firstName, &user.surname, &user.email, &user.dateOfBirth, &user.language, &user.country, &user.currencies); err != nil {
				err = errors.Trace(err)
				return
			}
			result = append(result, user)
		}
	}

	return
}
