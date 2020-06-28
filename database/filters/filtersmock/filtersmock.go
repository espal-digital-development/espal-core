package filtersmock

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/database/filters"
)

//go:generate moq -pkg filtersmock -out filters.go .. Factory
//go:generate moq -pkg filtersmock -out filter.go .. Filter
//go:generate moq -pkg filtersmock -out selectfield.go .. SelectField
//go:generate moq -pkg filtersmock -out searchfield.go .. SearchField
//go:generate moq -pkg filtersmock -out sortfield.go .. SortField
//go:generate moq -pkg filtersmock -out displaycolumn.go .. DisplayColumn
//go:generate moq -pkg filtersmock -out join.go .. Join

const defaultLimit = 10

// DefaultMocks returns a quick-to-use set of Mock instances.
func DefaultMocks() (*FactoryMock, *FilterMock, *SelectFieldMock, *SearchFieldMock, *JoinMock, *SortFieldMock,
	*DisplayColumnMock) {
	selectField := &SelectFieldMock{}
	selectField.SetMapToFunc = func(mapTo string) filters.SelectField {
		return selectField
	}
	selectField.SetAliasFunc = func(alias string) filters.SelectField {
		return selectField
	}
	searchField := &SearchFieldMock{}
	searchField.SetTableAliasFunc = func(tableAlias string) filters.SearchField {
		return searchField
	}
	sortField := &SortFieldMock{}
	sortField.SetTableAliasFunc = func(tableAlias string) filters.SortField {
		return sortField
	}
	displayColumn := &DisplayColumnMock{}
	displayColumn.SetPluralFunc = func(plural bool) filters.DisplayColumn {
		return displayColumn
	}
	join := &JoinMock{}
	filter := &FilterMock{
		NewSelectFieldFunc: func(name string) filters.SelectField {
			return selectField
		},
		NewSearchFieldFunc: func(name string) filters.SearchField {
			return searchField
		},
		NewJoinFunc: func(alias string, statement string) filters.Join {
			return join
		},
		NewSortFieldFunc: func(name string, descending bool) filters.SortField {
			return sortField
		},
		NewColumnFunc: func(name string) filters.DisplayColumn {
			return displayColumn
		},
		ProcessFunc: func() error {
			return nil
		},
		HasResultsFunc: func() bool {
			return false
		},
		LimitFunc: func() uint {
			return defaultLimit
		},
		RowsFunc: func() database.Rows {
			return nil
		},
	}
	filter.AddSelectFieldFunc = func(selectField filters.SelectField) filters.Filter {
		return filter
	}
	filter.AddSearchFieldFunc = func(searchField filters.SearchField) filters.Filter {
		return filter
	}
	filter.AddJoinStatementFunc = func(join filters.Join) filters.Filter {
		return filter
	}
	filter.AddSortFieldFunc = func(sortField filters.SortField) filters.Filter {
		return filter
	}
	filter.AddColumnFunc = func(displayColumn filters.DisplayColumn) filters.Filter {
		return filter
	}
	return &FactoryMock{
		NewFilterFunc: func(queryReader filters.QueryReader, m filters.Model) filters.Filter {
			return filter
		},
	}, filter, selectField, searchField, join, sortField, displayColumn
}
