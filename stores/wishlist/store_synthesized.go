// Code generated by espal-store-synthesizer. DO NOT EDIT.
package wishlist

import (
	"database/sql"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

var _ Store = &WishlistsStore{}

// Store represents a data interaction object.
type Store interface {
}

func (wishlistsStore *WishlistsStore) fetch(query string, withCreators bool, params ...interface{}) (result []*Wishlist, ok bool, err error) {
	rows, err := wishlistsStore.selecterDatabase.Query(query, params...)
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
	result = make([]*Wishlist, 0)
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, false, errors.Trace(err)
		}
		wishlist := newWishlist()
		fields := []interface{}{&wishlist.id, &wishlist.createdByID, &wishlist.updatedByID, &wishlist.createdAt, &wishlist.updatedAt, &wishlist.domainID, &wishlist.userID, &wishlist.sorting}
		if withCreators {
			fields = append(fields, &wishlist.createdByFirstName, &wishlist.createdBySurname, &wishlist.updatedByFirstName, &wishlist.updatedBySurname)
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, false, errors.Trace(err)
		}
		result = append(result, wishlist)
	}
	ok = len(result) > 0
	return
}

// New returns a new instance of WishlistsStore.
func New(selecterDatabase database.Database) (*WishlistsStore, error) {
	wishlistsStore := &WishlistsStore{
		selecterDatabase: selecterDatabase,
	}
	return wishlistsStore, nil
}
