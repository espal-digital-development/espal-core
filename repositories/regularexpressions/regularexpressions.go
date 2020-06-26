package regularexpressions

import (
	"regexp"

	"github.com/juju/errors"
)

var _ Repository = &RegularExpressions{}

// Repository represents a regular expressions repository.
type Repository interface {
	GetEmail() *regexp.Regexp
	GetPasswordRecoveryhash() *regexp.Regexp
	GetActivateAccounthash() *regexp.Regexp
	GetRouteIDs() *regexp.Regexp
}

// RegularExpressions holds common regular expressions.
type RegularExpressions struct {
	email                *regexp.Regexp
	passwordRecoveryhash *regexp.Regexp
	activateAccounthash  *regexp.Regexp
	routeIDs             *regexp.Regexp
}

// GetEmail returns the similar-named regular expression.
func (r *RegularExpressions) GetEmail() *regexp.Regexp {
	return r.email
}

// GetPasswordRecoveryhash returns the similar-named regular expression.
func (r *RegularExpressions) GetPasswordRecoveryhash() *regexp.Regexp {
	return r.passwordRecoveryhash
}

// GetActivateAccounthash returns the similar-named regular expression.
func (r *RegularExpressions) GetActivateAccounthash() *regexp.Regexp {
	return r.activateAccounthash
}

// GetRouteIDs returns the similar-named regular expression.
func (r *RegularExpressions) GetRouteIDs() *regexp.Regexp {
	return r.routeIDs
}

// New returns a new instance of RegularExpressions.
func New() (*RegularExpressions, error) {
	var err error
	r := &RegularExpressions{}
	if r.email, err = regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,9}\.?$`); err != nil {
		return nil, errors.Trace(err)
	}
	if r.passwordRecoveryhash, err = regexp.Compile(`^[a-zA-Z0-9]{72}$`); err != nil {
		return nil, errors.Trace(err)
	}
	if r.activateAccounthash, err = regexp.Compile(`^[a-zA-Z0-9]{72}$`); err != nil {
		return nil, errors.Trace(err)
	}
	if r.routeIDs, err = regexp.Compile(
		`^(?:\w{8}-\w{4}-\w{4}-\w{4}-\w{12})(?:,\w{8}-\w{4}-\w{4}-\w{4}-\w{12})*$`); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
