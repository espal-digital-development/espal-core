package filters

// Join holds database a filter join statement definition.
type Join interface {
	Alias() string
	Statement() string
}

type join struct {
	alias     string
	statement string
}

// Alias returns the join table alias.
func (join *join) Alias() string {
	return join.alias
}

// Alias returns the join statement.
func (join *join) Statement() string {
	return join.statement
}

// NewJoin returns a new instance of a Join.
func (filter *filter) NewJoin(alias string, statement string) Join {
	return &join{
		alias:     alias,
		statement: statement,
	}
}
