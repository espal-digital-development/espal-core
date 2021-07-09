package filters

import (
	"bytes"
	"math"
	"strconv"
	"strings"

	"github.com/espal-digital-development/system/units"
	"github.com/juju/errors"
)

// Process the (submitted) request information for the filter.
// Params legend:
// p = Page (Offset)
// r = Results Per Page (Limit)
// o = Ordering (Order By)
// c = Columns (Select ~)
// s = Search (LIKE %Word%).
func (f *filter) Process() error {
	// TODO :: Handle posted data when send (still to do; `o`, `c` & `s`)
	resultsPerPage, err := strconv.ParseUint(f.queryReader.QueryValue("r"), units.Base10, units.BitWidth64Bit)
	if err != nil {
		resultsPerPage = 0
	}
	if resultsPerPage > 0 {
		f.limit = uint(resultsPerPage)
	}
	page, err := strconv.ParseUint(f.queryReader.QueryValue("p"), units.Base10, units.BitWidth64Bit)
	if err != nil {
		page = 0
	}
	if page > 0 {
		f.offset = (uint(page) - 1) * f.limit
	}
	if search := f.queryReader.QueryValue("s"); len(search) > 0 {
		f.search = search
		f.makeSearchChunks()
	}

	f.query = bytes.NewBufferString("SELECT ")
	f.countQuery = bytes.NewBufferString("SELECT COUNT(*)")

	f.handleSelectFields()

	f.query.WriteString(` FROM "`)
	f.query.WriteString(f.table)
	f.query.WriteString(`"`)
	f.countQuery.WriteString(` FROM "`)
	f.countQuery.WriteString(f.table)
	f.countQuery.WriteString(`"`)
	if f.tableAlias != "" {
		f.query.WriteString(" ")
		f.query.WriteString(f.tableAlias)
		f.countQuery.WriteString(" ")
		f.countQuery.WriteString(f.tableAlias)
	}

	f.handleJoinStatements()

	if len(f.fields) > 0 || f.search != "" {
		f.query.WriteString(" WHERE")
		f.countQuery.WriteString(" WHERE")

		for k := range f.fields {
			f.handleWhereStatementForField(k)
		}

		if f.search != "" {
			f.addSearchToQuery()
		}
	}

	f.sort()
	f.pagination()

	f.handleParams()

	// TODO :: If there is no WHERE and all joins are LEFT, all joins could be removed. Yet there is only one
	// exception; if a JOIN results in an expansion of results it could result in unexpected results.

	// TODO :: 7 COUNT() on is not very efficient. Need a better way (UPDATE: Better way in Cockroach?)
	if err := f.selecterDatabase.QueryRow(f.countQuery.String(), f.params...).Scan(&f.totalResults); err != nil {
		return errors.Trace(err)
	}

	// TODO :: If someone requests a too high page it can still count +0 results.
	// Need to reverse the totalResults for the pagination/results still to be correct.
	if f.totalResults > 0 {
		f.rows, f.rowsError = f.selecterDatabase.Query(f.query.String(), f.params...)
	}

	f.totalPages = uint(math.Ceil(float64(f.totalResults) / float64(f.limit)))
	f.currentPage = f.offset/f.limit + 1
	f.hasPreviousPage = f.currentPage > 1
	f.hasNextPage = f.currentPage < f.totalPages

	if f.HasError() {
		return f.rowsError
	}

	return nil
}

func (f *filter) makeSearchChunks() {
	pieces := strings.Split(strings.TrimSpace(f.search), " ")
	for k := range pieces {
		pieces[k] = strings.TrimSpace(pieces[k])
		if len(pieces[k]) == 0 {
			continue
		}
		f.searchChunks = append(f.searchChunks, pieces[k])
	}
}

func (f *filter) handleSelectFields() {
	var firstHad bool
	for k := range f.selectFields {
		if firstHad {
			f.query.WriteString(",")
		} else {
			firstHad = true
		}

		if f.selectFields[k].Alias() == "" {
			f.query.WriteString(f.tableAlias)
			f.query.WriteString(".")
		} else if f.selectFields[k].Alias() != "" {
			f.query.WriteString(f.selectFields[k].Alias())
			f.query.WriteString(".")
		}

		f.query.WriteString(`"`)
		f.query.WriteString(f.selectFields[k].Name())
		f.query.WriteString(`"`)
	}
}

func (f *filter) handleJoinStatements() {
	// TODO :: 77 Maybe JOIN can be done too in a way with sub-WHERE structs to fully build it instead of
	// concatenating a custom written piece of SQL.
	for k := range f.joinStatements {
		var found bool
		for k2 := range f.selectFields {
			if f.selectFields[k2].Alias() == f.joinStatements[k].Alias() {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		f.query.WriteString(" ")
		f.query.WriteString(f.joinStatements[k].Statement())
		f.countQuery.WriteString(" ")
		f.countQuery.WriteString(f.joinStatements[k].Statement())
	}
}

func (f *filter) handleParams() {
	f.params = make([]interface{}, uint(len(f.fields))+f.amountOfBetweens-f.amountOfNotNulls)

	for k := range f.fields {
		if filterFieldActionIsNull == f.fields[k].action {
			continue
		}
		if filterFieldActionLike == f.fields[k].action {
			f.params = append(f.params, "%"+f.fields[k].value+"%")
		} else {
			f.params = append(f.params, f.fields[k].value)
		}
		if filterFieldActionBetween == f.fields[k].action {
			f.params = append(f.params, f.fields[k].value2)
		}
	}

	if f.search != "" {
		if sfLen := len(f.searchFields); sfLen > 0 {
			for k := range f.searchChunks {
				for i := 0; i < sfLen; i++ {
					f.params = append(f.params, "%"+f.searchChunks[k]+"%")
				}
			}
		}
	}
}
