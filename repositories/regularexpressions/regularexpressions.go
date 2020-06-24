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
func (regularExpressions *RegularExpressions) GetEmail() *regexp.Regexp {
	return regularExpressions.email
}

// GetPasswordRecoveryhash returns the similar-named regular expression.
func (regularExpressions *RegularExpressions) GetPasswordRecoveryhash() *regexp.Regexp {
	return regularExpressions.passwordRecoveryhash
}

// GetActivateAccounthash returns the similar-named regular expression.
func (regularExpressions *RegularExpressions) GetActivateAccounthash() *regexp.Regexp {
	return regularExpressions.activateAccounthash
}

// GetRouteIDs returns the similar-named regular expression.
func (regularExpressions *RegularExpressions) GetRouteIDs() *regexp.Regexp {
	return regularExpressions.routeIDs
}

// New returns a new instance of RegularExpressions.
func New() (*RegularExpressions, error) {
	var err error
	regularExpressions := &RegularExpressions{}
	if regularExpressions.email, err = regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,9}\.?$`); err != nil {
		return nil, errors.Trace(err)
	}
	if regularExpressions.passwordRecoveryhash, err = regexp.Compile(`^[a-zA-Z0-9]{72}$`); err != nil {
		return nil, errors.Trace(err)
	}
	if regularExpressions.activateAccounthash, err = regexp.Compile(`^[a-zA-Z0-9]{72}$`); err != nil {
		return nil, errors.Trace(err)
	}
	if regularExpressions.routeIDs, err = regexp.Compile(`^(?:\w{8}-\w{4}-\w{4}-\w{4}-\w{12})(?:,\w{8}-\w{4}-\w{4}-\w{4}-\w{12})*$`); err != nil { // `^[\w-]+(?:,[\w-]+)*$`
		return nil, errors.Trace(err)
	}
	return regularExpressions, nil
}
