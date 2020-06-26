package runner

import (
	"github.com/espal-digital-development/espal-core/stores/cachenotify"
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/stores/session"
	"github.com/espal-digital-development/espal-core/stores/setting"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/stores/slug"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/stores/user/address"
	"github.com/espal-digital-development/espal-core/stores/user/contact"
	"github.com/espal-digital-development/espal-core/stores/user/group"
	"github.com/juju/errors"
)

type stores struct {
	setting     setting.Store
	cacheNotify cachenotify.Store
	session     session.Store
	domain      domain.Store
	site        site.Store
	slug        slug.Store
	userGroup   group.Store
	user        user.Store
	userAddress address.Store
	userContact contact.Store
	forum       forum.Store
}

func (r *Runner) dataStores() error {
	var err error
	if r.stores.setting, err = setting.New(r.databases.selecter); err != nil {
		return errors.Trace(err)
	}
	if r.stores.cacheNotify, err = cachenotify.New(r.databases.selecter, r.databases.updater); err != nil {
		return errors.Trace(err)
	}
	if r.stores.session, err = session.New(r.databases.selecter, r.databases.inserter, r.databases.updater); err != nil {
		return errors.Trace(err)
	}
	if r.stores.domain, err = domain.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseFilters); err != nil {
		return errors.Trace(err)
	}
	if r.stores.site, err = site.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseFilters, r.repositories.translations, r.services.logger); err != nil {
		return errors.Trace(err)
	}
	if r.stores.slug, err = slug.New(r.databases.selecter); err != nil {
		return errors.Trace(err)
	}
	if r.stores.userGroup, err = group.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.services.databaseFilters, r.repositories.translations, r.services.logger); err != nil {
		return errors.Trace(err)
	}
	if r.stores.user, err = user.New(r.databases.selecter, r.databases.inserter, r.databases.updater,
		r.databases.deletor, r.services.databaseFilters, r.repositories.translations,
		r.repositories.userRights); err != nil {
		return errors.Trace(err)
	}
	if r.stores.userAddress, err = address.New(r.databases.selecter, r.databases.updater, r.databases.deletor,
		r.repositories.translations, r.repositories.countries, r.services.logger); err != nil {
		return errors.Trace(err)
	}
	if r.stores.userContact, err = contact.New(r.databases.selecter, r.databases.deletor,
		r.repositories.translations); err != nil {
		return errors.Trace(err)
	}
	if r.stores.forum, err = forum.New(r.databases.selecter, r.databases.deletor); err != nil {
		return errors.Trace(err)
	}
	return nil
}
