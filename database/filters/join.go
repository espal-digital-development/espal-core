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
func (j *join) Alias() string {
	return j.alias
}

// Alias returns the join statement.
func (j *join) Statement() string {
	return j.statement
}

// NewJoin returns a new instance of a Join.
func (f *filter) NewJoin(alias string, statement string) Join {
	return &join{
		alias:     alias,
		statement: statement,
	}
}
