package runner

import (
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/stores/notification"
	"github.com/espal-digital-development/espal-core/stores/session"
	"github.com/espal-digital-development/espal-core/stores/setting"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/stores/slug"
	"github.com/espal-digital-development/espal-core/stores/user"
	useraddress "github.com/espal-digital-development/espal-core/stores/user/address"
	usercontact "github.com/espal-digital-development/espal-core/stores/user/contact"
	usergroup "github.com/espal-digital-development/espal-core/stores/user/group"
	"github.com/juju/errors"
)

type stores struct {
	setting      setting.Store
	notification notification.Store
	session      session.Store
	domain       domain.Store
	site         site.Store
	slug         slug.Store
	userGroup    usergroup.Store
	user         user.Store
	userAddress  useraddress.Store
	userContact  usercontact.Store
	forum        forum.Store
}

// Setting returns the Setting Store.
func (s *stores) Setting() setting.Store {
	return s.setting
}

// Notification returns the Notification Store.
func (s *stores) Notification() notification.Store {
	return s.notification
}

// Session returns the Session Store.
func (s *stores) Session() session.Store {
	return s.session
}

// Domain returns the Domain Store.
func (s *stores) Domain() domain.Store {
	return s.domain
}

// Site returns the Site Store.
func (s *stores) Site() site.Store {
	return s.site
}

// Slug returns the Slug Store.
func (s *stores) Slug() slug.Store {
	return s.slug
}

// UserGroup returns the UserGroup Store.
func (s *stores) UserGroup() usergroup.Store {
	return s.userGroup
}

// User returns the User Store.
func (s *stores) User() user.Store {
	return s.user
}

// UserAddress returns the UserAddress Store.
func (s *stores) UserAddress() useraddress.Store {
	return s.userAddress
}

// UserContact returns the UserContact Store.
func (s *stores) UserContact() usercontact.Store {
	return s.userContact
}

// Forum returns the Forum Store.
func (s *stores) Forum() forum.Store {
	return s.forum
}

func (r *Runner) dataStores() error {
	var err error
	if r.stores.setting, err = setting.New(r.databases.selecter); err != nil {
		return errors.Trace(err)
	}
	if r.stores.notification, err = notification.New(r.databases.selecter, r.databases.inserter,
		r.databases.deletor); err != nil {
		return errors.Trace(err)
	}
	if r.stores.session, err = session.New(r.databases.selecter, r.databases.inserter, r.databases.updater); err != nil {
		return errors.Trace(err)
	}
	if r.stores.domain, err = domain.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseQueryHelper, r.services.databaseFilters); err != nil {
		return errors.Trace(err)
	}
	if r.stores.site, err = site.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseQueryHelper, r.services.databaseFilters, r.repositories.translations,
		r.services.logger); err != nil {
		return errors.Trace(err)
	}
	if r.stores.slug, err = slug.New(r.databases.selecter); err != nil {
		return errors.Trace(err)
	}
	if r.stores.userGroup, err = usergroup.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseQueryHelper, r.services.databaseFilters, r.repositories.translations,
		r.services.logger); err != nil {
		return errors.Trace(err)
	}
	if r.stores.user, err = user.New(r.databases.selecter, r.databases.inserter, r.databases.updater,
		r.databases.deletor, r.services.databaseQueryHelper, r.services.databaseFilters, r.repositories.translations,
		r.repositories.userRights); err != nil {
		return errors.Trace(err)
	}
	if r.stores.userAddress, err = useraddress.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseQueryHelper, r.repositories.translations, r.repositories.countries,
		r.services.logger); err != nil {
		return errors.Trace(err)
	}
	if r.stores.userContact, err = usercontact.New(r.databases.selecter, r.databases.deletor,
		r.services.databaseQueryHelper, r.repositories.translations); err != nil {
		return errors.Trace(err)
	}
	if r.stores.forum, err = forum.New(r.databases.selecter, r.databases.deletor); err != nil {
		return errors.Trace(err)
	}

	for k := range r.modulesRegistry {
		r.modulesRegistry[k].RegisterCoreStores(r.stores)
	}

	return nil
}
