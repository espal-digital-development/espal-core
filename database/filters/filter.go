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
func (f *filter) AddSelectField(selectField SelectField) Filter {
	f.selectFields = append(f.selectFields, selectField)
	return f
}

// AddColumn adds a displayColumn.
func (f *filter) AddColumn(displayColumn DisplayColumn) Filter {
	f.columns = append(f.columns, displayColumn)
	return f
}

// AddSearchField adds a searchField.
func (f *filter) AddSearchField(searchField SearchField) Filter {
	f.searchFields = append(f.searchFields, searchField)
	return f
}

// AddJoinStatement adds a joinStatement.
func (f *filter) AddJoinStatement(joinStatement Join) Filter {
	f.joinStatements = append(f.joinStatements, joinStatement)
	return f
}

// AddSortField adds a sortField.
func (f *filter) AddSortField(sortField SortField) Filter {
	f.sortFields = append(f.sortFields, sortField)
	return f
}

// TotalResults returns the total results in the current filtered result.
func (f *filter) TotalResults() uint {
	return f.totalResults
}

// TotalPages returns the total pages in the current filtered result.
func (f *filter) TotalPages() uint {
	return f.totalPages
}

// CurrentPage returns the current page in the current filtered result.
func (f *filter) CurrentPage() uint {
	return f.currentPage
}

// HasPreviousPage returns if there's a previous page in the current filtered
// result based on the current page.
func (f *filter) HasPreviousPage() bool {
	return f.hasPreviousPage
}

// HasNextPage returns if there's a next page in the current filtered
// result based on the current page.
func (f *filter) HasNextPage() bool {
	return f.hasNextPage
}

// PaginationBlocks returns an easy index list how to display the pagination blocks.
func (f *filter) PaginationBlocks() []uint {
	var page uint
	var blocks []uint
	if f.TotalPages() <= maxPaginationBarPages {
		// [ 1 2 3 4 5 6 7 ] (up till 7)
		for page = 1; page <= f.totalPages; page++ {
			blocks = append(blocks, page)
		}
		return blocks
	}
	if f.CurrentPage() <= paginationLeftTreshHold {
		// [ 1 2 3 4 5 . 9 ] (. (or ...) will be indicated as 0)
		return []uint{1, 2, 3, 4, 5, 0, f.totalPages}
	}
	if f.CurrentPage() >= (f.TotalPages() - 3) {
		// [ 1 . 5 6 7 8 9 ]
		blocks = append(blocks, 1)
		blocks = append(blocks, 0)
		for page = f.totalPages - paginationLeftTreshHold; page <= f.totalPages; page++ {
			blocks = append(blocks, page)
		}
		return blocks
	}
	// [ 1 . 4 5 6 . 9 ]
	return []uint{1, 0, f.currentPage - 1, f.currentPage, f.currentPage + 1, 0, f.totalPages}
}

// incrementParameterCount increments the parameter count and then returns the new value.
func (f *filter) incrementParameterCount() uint16 {
	f.parameterCount++
	return f.parameterCount
}

func (f *filter) sort() error {
	if len(f.sortFields) == 0 {
		return nil
	}
	if _, err := f.query.WriteString(" ORDER BY"); err != nil {
		return errors.Trace(err)
	}
	var firstHad bool
	for k := range f.sortFields {
		if firstHad {
			if _, err := f.query.WriteString(", "); err != nil {
				return errors.Trace(err)
			}
		} else {
			if _, err := f.query.WriteString(" "); err != nil {
				return errors.Trace(err)
			}
			firstHad = true
		}

		if f.sortFields[k].TableAlias() != "" {
			if _, err := f.query.WriteString(f.sortFields[k].TableAlias()); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.query.WriteString(`.`); err != nil {
				return errors.Trace(err)
			}
		}

		if _, err := f.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString(f.sortFields[k].Name()); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}

		if f.sortFields[k].Descending() {
			if _, err := f.query.WriteString(" DESC"); err != nil {
				return errors.Trace(err)
			}
		}
	}
	return nil
}

func (f *filter) pagination() error {
	if f.limit > 0 {
		if _, err := f.query.WriteString(" LIMIT "); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString(strconv.FormatUint(uint64(f.limit), 10)); err != nil {
			return errors.Trace(err)
		}
		if f.offset > 0 {
			if _, err := f.query.WriteString(" OFFSET "); err != nil {
				return errors.Trace(err)
			}
			if _, err := f.query.WriteString(strconv.FormatUint(uint64(f.offset), 10)); err != nil {
				return errors.Trace(err)
			}
		}
	}
	if f.limit == 0 && f.offset > 0 {
		if _, err := f.query.WriteString(" OFFSET "); err != nil {
			return errors.Trace(err)
		}
		if _, err := f.query.WriteString(strconv.FormatUint(uint64(f.offset), 10)); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// TableAlias returns the alias for this filter's model.
func (f *filter) TableAlias() string {
	return f.tableAlias
}

// HasResults returns if there were any results at all in the f.
func (f *filter) HasResults() bool {
	return f.totalResults > 0
}

// HasError returns if there was an error while running the query.
func (f *filter) HasError() bool {
	return f.rowsError != nil && sql.ErrNoRows != f.rowsError
}

// Limit returns the set result limit.
func (f *filter) Limit() uint {
	return f.limit
}

// Rows returns the internal rows cursor.
func (f *filter) Rows() database.Rows {
	return f.rows
}

// CloseRows closes the internal rows cursor.
func (f *filter) CloseRows() error {
	return f.rows.Close()
}

// ShouldShowSearch returns if the search field should be shown.
func (f *filter) ShouldShowSearch() bool {
	return (f.TotalResults() > f.limit || f.search != "") && len(f.searchFields) > 0
}

// ColumnsInOrder returns the filter columns in their required order.
func (f *filter) ColumnsInOrder() []DisplayColumn {
	if length := len(f.columnDisplayOrder); length > 0 {
		columns := make([]DisplayColumn, 0, length)
		for k := range f.columnDisplayOrder {
			columns = append(columns, f.columns[f.columnDisplayOrder[k]])
		}
		return columns
	}
	return f.columns
}

func (f *filter) perror(i int, err error) {
	if err != nil {
		// TODO :: 777 Can't log easily here because it's instanced. Need some solution
		panic(errors.ErrorStack(err))
	}
}
