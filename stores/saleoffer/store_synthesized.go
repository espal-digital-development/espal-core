// Code generated by espal-store-synthesizer. DO NOT EDIT.
package saleoffer

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &SaleOffersStore{}

// Store represents a data interaction object.
type Store interface {
}

func (s *SaleOffersStore) fetch(query string, withCreators bool, params ...interface{}) (result []*SaleOffer, ok bool, err error) {
	rows, err := s.selecterDatabase.Query(query, params...)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	defer func(dbRows database.Rows) {
		closeErr := dbRows.Close()
		if err != nil && closeErr != nil {
			err = errors.Wrap(err, closeErr)
		} else if closeErr != nil {
			err = errors.Trace(closeErr)
		}
	}(rows)
	result = make([]*SaleOffer, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		s := newSaleOffer()
		fields := []interface{}{&s.id, &s.createdByID, &s.updatedByID, &s.createdAt, &s.updatedAt, &s.userID, &s.domainID, &s.currency, &s.code, &s.userInfoBusiness, &s.userInfoBusinessCocNumber, &s.userInfoFirstName, &s.userInfoSurname, &s.userInfoStreet, &s.userInfoStreetLine2, &s.userInfoNumber, &s.userInfoNumberAddition, &s.userInfoZipCode, &s.userInfoCity, &s.userInfoState, &s.userInfoCountry, &s.userInfoPhoneNumber, &s.userInfoEmail, &s.shippingAddressBusiness, &s.shippingAddressBusinessCocNumber, &s.shippingAddressFirstName, &s.shippingAddressSurname, &s.shippingAddressStreet, &s.shippingAddressStreetLine2, &s.shippingAddressNumber, &s.shippingAddressNumberAddition, &s.shippingAddressZipCode, &s.shippingAddressCity, &s.shippingAddressState, &s.shippingAddressCountry, &s.shippingAddressPhoneNumber, &s.shippingAddressEmail, &s.comments, &s.sellingPartyAutograph, &s.buyingPartyAutograph}
		if withCreators {
			fields = append(fields, &s.createdByFirstName, &s.createdBySurname, &s.updatedByFirstName, &s.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, s)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of SaleOffersStore.
func New(selecterDatabase database.Database) (*SaleOffersStore, error) {
	s := &SaleOffersStore{
		selecterDatabase: selecterDatabase,
	}
	return s, nil
}
