package filters

import (
	"bytes"
	"database/sql"
	"strconv"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const (
	maxPaginationBarPages   = 7
	paginationLeftTreshHold = 4
)

// Model represents a database meta object.
type Model interface {
	TableName() string
	TableAlias() string
}

// QueryReader represents a request query object.
type QueryReader interface {
	QueryValue(string) string
}

// Filter holds filterdata that associate with the given resultset.
type Filter interface {
	AddSelectField(selectField SelectField) Filter
	AddColumn(displayColumn DisplayColumn) Filter
	AddSearchField(searchField SearchField) Filter
	AddJoinStatement(joinStatement Join) Filter
	AddSortField(sortField SortField) Filter

	TotalResults() uint
	TotalPages() uint
	CurrentPage() uint
	HasPreviousPage() bool
	HasNextPage() bool
	PaginationBlocks() []uint
	TableAlias() string
	HasResults() bool
	HasError() bool
	Limit() uint
	Rows() database.Rows
	CloseRows() error
	ShouldShowSearch() bool
	ColumnsInOrder() []DisplayColumn

	RenderOverviewColumnTableHeaders(ctx Context) string
	RenderOverviewFilter(ctx Context) string

	NewColumn(name string) DisplayColumn
	NewSelectField(name string) SelectField
	NewJoin(alias string, statement string) Join
	NewSortField(name string, descending bool) SortField
	NewSearchField(name string) SearchField
	Process() error
}

type filter struct {
	// TODO :: Try to learn what previously made the fastest exclusion order.
	// So if int-checks left the least results, try to use that again.
	queryReader      QueryReader
	selecterDatabase database.Database
	search           string
	table            string
	tableAlias       string
	limit            uint
	offset           uint
	// TODO :: Sorting could be limited to prevent excessive precise sorting
	// maxSortingCount    uint
	selectFields       []SelectField
	columns            []DisplayColumn
	columnDisplayOrder []int
	fields             []field
	searchFields       []SearchField
	joinStatements     []Join
	sortFields         []SortField

	rows      database.Rows
	rowsError error

	// parameterCount counts all the incremental $1, $2 etc. parameters that are used
	parameterCount         uint16
	amountOfBetweens       uint
	amountOfNotNulls       uint
	firstWhereStatementHad bool
	params                 []interface{}
	query                  *bytes.Buffer
	countQuery             *bytes.Buffer

	searchChunks []string

	totalResults    uint
	totalPages      uint
	currentPage     uint
	hasPreviousPage bool
	hasNextPage     bool
}

// AddSelectField adds a selectField.
func (filter *filter) AddSelectField(selectField SelectField) Filter {
	filter.selectFields = append(filter.selectFields, selectField)
	return filter
}

// AddColumn adds a displayColumn.
func (filter *filter) AddColumn(displayColumn DisplayColumn) Filter {
	filter.columns = append(filter.columns, displayColumn)
	return filter
}

// AddSearchField adds a searchField.
func (filter *filter) AddSearchField(searchField SearchField) Filter {
	filter.searchFields = append(filter.searchFields, searchField)
	return filter
}

// AddJoinStatement adds a joinStatement.
func (filter *filter) AddJoinStatement(joinStatement Join) Filter {
	filter.joinStatements = append(filter.joinStatements, joinStatement)
	return filter
}

// AddSortField adds a sortField.
func (filter *filter) AddSortField(sortField SortField) Filter {
	filter.sortFields = append(filter.sortFields, sortField)
	return filter
}

// TotalResults returns the total results in the current filtered result.
func (filter *filter) TotalResults() uint {
	return filter.totalResults
}

// TotalPages returns the total pages in the current filtered result.
func (filter *filter) TotalPages() uint {
	return filter.totalPages
}

// CurrentPage returns the current page in the current filtered result.
func (filter *filter) CurrentPage() uint {
	return filter.currentPage
}

// HasPreviousPage returns if there's a previous page in the current filtered
// result based on the current page.
func (filter *filter) HasPreviousPage() bool {
	return filter.hasPreviousPage
}

// HasNextPage returns if there's a next page in the current filtered
// result based on the current page.
func (filter *filter) HasNextPage() bool {
	return filter.hasNextPage
}

// PaginationBlocks returns an easy index list how to display the pagination blocks.
func (filter *filter) PaginationBlocks() []uint {
	var page uint
	var blocks []uint
	if filter.TotalPages() <= maxPaginationBarPages {
		// [ 1 2 3 4 5 6 7 ] (up till 7)
		for page = 1; page <= filter.totalPages; page++ {
			blocks = append(blocks, page)
		}
		return blocks
	}
	if filter.CurrentPage() <= paginationLeftTreshHold {
		// [ 1 2 3 4 5 . 9 ] (. (or ...) will be indicated as 0)
		return []uint{1, 2, 3, 4, 5, 0, filter.totalPages}
	}
	if filter.CurrentPage() >= (filter.TotalPages() - 3) {
		// [ 1 . 5 6 7 8 9 ]
		blocks = append(blocks, 1)
		blocks = append(blocks, 0)
		for page = filter.totalPages - paginationLeftTreshHold; page <= filter.totalPages; page++ {
			blocks = append(blocks, page)
		}
		return blocks
	}
	// [ 1 . 4 5 6 . 9 ]
	return []uint{1, 0, filter.currentPage - 1, filter.currentPage, filter.currentPage + 1, 0, filter.totalPages}
}

// incrementParameterCount increments the parameter count and then returns the new value.
func (filter *filter) incrementParameterCount() uint16 {
	filter.parameterCount++
	return filter.parameterCount
}

func (filter *filter) sort() error {
	if len(filter.sortFields) == 0 {
		return nil
	}
	if _, err := filter.query.WriteString(" ORDER BY"); err != nil {
		return errors.Trace(err)
	}
	var firstHad bool
	for k := range filter.sortFields {
		if firstHad {
			if _, err := filter.query.WriteString(", "); err != nil {
				return errors.Trace(err)
			}
		} else {
			if _, err := filter.query.WriteString(" "); err != nil {
				return errors.Trace(err)
			}
			firstHad = true
		}

		if filter.sortFields[k].TableAlias() != "" {
			if _, err := filter.query.WriteString(filter.sortFields[k].TableAlias()); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString(`.`); err != nil {
				return errors.Trace(err)
			}
		}

		if _, err := filter.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(filter.sortFields[k].Name()); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}

		if filter.sortFields[k].Descending() {
			if _, err := filter.query.WriteString(" DESC"); err != nil {
				return errors.Trace(err)
			}
		}
	}
	return nil
}

func (filter *filter) pagination() error {
	if filter.limit > 0 {
		if _, err := filter.query.WriteString(" LIMIT "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(strconv.FormatUint(uint64(filter.limit), 10)); err != nil {
			return errors.Trace(err)
		}
		if filter.offset > 0 {
			if _, err := filter.query.WriteString(" OFFSET "); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString(strconv.FormatUint(uint64(filter.offset), 10)); err != nil {
				return errors.Trace(err)
			}
		}
	}
	if filter.limit == 0 && filter.offset > 0 {
		if _, err := filter.query.WriteString(" OFFSET "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(strconv.FormatUint(uint64(filter.offset), 10)); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// TableAlias returns the alias for this filter's model.
func (filter *filter) TableAlias() string {
	return filter.tableAlias
}

// HasResults returns if there were any results at all in the filter.
func (filter *filter) HasResults() bool {
	return filter.totalResults > 0
}

// HasError returns if there was an error while running the query.
func (filter *filter) HasError() bool {
	return filter.rowsError != nil && sql.ErrNoRows != filter.rowsError
}

// Limit returns the set result limit.
func (filter *filter) Limit() uint {
	return filter.limit
}

// Rows returns the internal rows cursor.
func (filter *filter) Rows() database.Rows {
	return filter.rows
}

// CloseRows closes the internal rows cursor.
func (filter *filter) CloseRows() error {
	return filter.rows.Close()
}

// ShouldShowSearch returns if the search field should be shown.
func (filter *filter) ShouldShowSearch() bool {
	return (filter.TotalResults() > filter.limit || filter.search != "") && len(filter.searchFields) > 0
}

// ColumnsInOrder returns the filter columns in their required order.
func (filter *filter) ColumnsInOrder() []DisplayColumn {
	if length := len(filter.columnDisplayOrder); length > 0 {
		columns := make([]DisplayColumn, 0, length)
		for k := range filter.columnDisplayOrder {
			columns = append(columns, filter.columns[filter.columnDisplayOrder[k]])
		}
		return columns
	}
	return filter.columns
}

func (filter *filter) perror(i int, err error) {
	if err != nil {
		// TODO :: 777 Can't log easily here because it's instanced. Need some solution
		panic(errors.ErrorStack(err))
	}
}
