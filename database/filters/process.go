package filters

import (
	"bytes"
	"math"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

// Process the (submitted) request information for the filter.
// Params legend:
// p = Page (Offset)
// r = Results Per Page (Limit)
// o = Ordering (Order By)
// c = Columns (Select ~)
// s = Search (LIKE %Word%)
func (filter *filter) Process() error {
	// TODO :: Handle posted data when send (still to do; `o`, `c` & `s`)
	resultsPerPage, err := strconv.ParseUint(filter.queryReader.QueryValue("r"), 10, 64)
	if err != nil {
		resultsPerPage = 0
	}
	if resultsPerPage > 0 {
		filter.limit = uint(resultsPerPage)
	}
	page, err := strconv.ParseUint(filter.queryReader.QueryValue("p"), 10, 64)
	if err != nil {
		page = 0
	}
	if page > 0 {
		filter.offset = (uint(page) - 1) * filter.limit
	}
	if search := filter.queryReader.QueryValue("s"); len(search) > 0 {
		filter.search = search
		filter.makeSearchChunks()
	}

	filter.query = bytes.NewBufferString("SELECT ")
	filter.countQuery = bytes.NewBufferString("SELECT COUNT(*)")

	if err := filter.handleSelectFields(); err != nil {
		return errors.Trace(err)
	}

	if _, err := filter.query.WriteString(` FROM "`); err != nil {
		return errors.Trace(err)
	}
	if _, err := filter.query.WriteString(filter.table); err != nil {
		return errors.Trace(err)
	}
	if _, err := filter.query.WriteString(`"`); err != nil {
		return errors.Trace(err)
	}
	if _, err := filter.countQuery.WriteString(` FROM "`); err != nil {
		return errors.Trace(err)
	}
	if _, err := filter.countQuery.WriteString(filter.table); err != nil {
		return errors.Trace(err)
	}
	if _, err := filter.countQuery.WriteString(`"`); err != nil {
		return errors.Trace(err)
	}
	if filter.tableAlias != "" {
		if _, err := filter.query.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(filter.tableAlias); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(filter.tableAlias); err != nil {
			return errors.Trace(err)
		}
	}

	if err := filter.handleJoinStatements(); err != nil {
		return errors.Trace(err)
	}

	if len(filter.fields) > 0 || filter.search != "" {
		if _, err := filter.query.WriteString(" WHERE"); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" WHERE"); err != nil {
			return errors.Trace(err)
		}

		for k := range filter.fields {
			if err := filter.handleWhereStatementForField(k); err != nil {
				return errors.Trace(err)
			}
		}

		if filter.search != "" {
			if err := filter.addSearchToQuery(); err != nil {
				return errors.Trace(err)
			}
		}
	}

	if err := filter.sort(); err != nil {
		return errors.Trace(err)
	}
	if err := filter.pagination(); err != nil {
		return errors.Trace(err)
	}

	filter.handleParams()

	// TODO :: If there is no WHERE and all joins are LEFT, all joins could be removed. Yet there is only one exception; if a JOIN results in an expansion of results it could result in unexpected results.

	// TODO :: 7 :: COUNT() on is not very efficient. Need a better way (UPDATE: Better way in Cockroach?)
	if err := filter.selecterDatabase.QueryRow(filter.countQuery.String(), filter.params...).Scan(&filter.totalResults); err != nil {
		return errors.Trace(err)
	}

	// TODO :: If someone requests a too high page it can still count +0 results.
	// Need to reverse the totalResults for the pagination/results still to be correct.
	if filter.totalResults > 0 {
		filter.rows, filter.rowsError = filter.selecterDatabase.Query(filter.query.String(), filter.params...)
	}

	filter.totalPages = uint(math.Ceil(float64(filter.totalResults) / float64(filter.limit)))
	filter.currentPage = filter.offset/filter.limit + 1
	filter.hasPreviousPage = filter.currentPage > 1
	filter.hasNextPage = filter.currentPage < filter.totalPages

	if filter.HasError() {
		return filter.rowsError
	}

	return nil
}

func (filter *filter) makeSearchChunks() {
	pieces := strings.Split(strings.Trim(filter.search, " "), " ")
	for k := range pieces {
		pieces[k] = strings.Trim(pieces[k], " ")
		if len(pieces[k]) == 0 {
			continue
		}
		filter.searchChunks = append(filter.searchChunks, pieces[k])
	}
}

func (filter *filter) handleSelectFields() error {
	var firstHad bool
	for k := range filter.selectFields {
		if firstHad {
			if _, err := filter.query.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		} else {
			firstHad = true
		}

		if filter.selectFields[k].Alias() == "" {
			if _, err := filter.query.WriteString(filter.tableAlias); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString("."); err != nil {
				return errors.Trace(err)
			}
		} else if filter.selectFields[k].Alias() != "" {
			if _, err := filter.query.WriteString(filter.selectFields[k].Alias()); err != nil {
				return errors.Trace(err)
			}
			if _, err := filter.query.WriteString("."); err != nil {
				return errors.Trace(err)
			}
		}

		if _, err := filter.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(filter.selectFields[k].Name()); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(`"`); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (filter *filter) handleJoinStatements() error {
	// TODO :: 77 Maybe JOIN can be done too in a way with sub-WHERE structs to fully build it instead of concatenating a custom written piece of SQL
	for k := range filter.joinStatements {
		var found bool
		for k2 := range filter.selectFields {
			if filter.selectFields[k2].Alias() == filter.joinStatements[k].Alias() {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		if _, err := filter.query.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.query.WriteString(filter.joinStatements[k].Statement()); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(" "); err != nil {
			return errors.Trace(err)
		}
		if _, err := filter.countQuery.WriteString(filter.joinStatements[k].Statement()); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func (filter *filter) handleParams() {
	filter.params = make([]interface{}, uint(len(filter.fields))+filter.amountOfBetweens-filter.amountOfNotNulls)

	for k := range filter.fields {
		if filterFieldActionIsNull == filter.fields[k].action {
			continue
		}
		if filterFieldActionLike == filter.fields[k].action {
			filter.params = append(filter.params, "%"+filter.fields[k].value+"%")
		} else {
			filter.params = append(filter.params, filter.fields[k].value)
		}
		if filterFieldActionBetween == filter.fields[k].action {
			filter.params = append(filter.params, filter.fields[k].value2)
		}
	}

	if filter.search != "" {
		if sfLen := len(filter.searchFields); sfLen > 0 {
			for k := range filter.searchChunks {
				for i := 0; i < sfLen; i++ {
					filter.params = append(filter.params, "%"+filter.searchChunks[k]+"%")
				}
			}
		}
	}
}
