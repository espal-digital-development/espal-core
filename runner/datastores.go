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

func (runner *Runner) dataStores() error {
	var err error
	if runner.stores.setting, err = setting.New(runner.databases.selecter); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.cacheNotify, err = cachenotify.New(runner.databases.selecter, runner.databases.updater); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.session, err = session.New(runner.databases.selecter, runner.databases.inserter, runner.databases.updater); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.domain, err = domain.New(runner.databases.selecter, runner.databases.updater, runner.databases.deletor, runner.services.databaseFilters); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.site, err = site.New(runner.databases.selecter, runner.databases.updater, runner.databases.deletor, runner.services.databaseFilters, runner.repositories.translations, runner.services.logger); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.slug, err = slug.New(runner.databases.selecter); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.userGroup, err = group.New(runner.databases.selecter, runner.databases.updater, runner.databases.deletor, runner.services.databaseFilters, runner.repositories.translations, runner.services.logger); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.user, err = user.New(runner.databases.selecter, runner.databases.inserter, runner.databases.updater, runner.databases.deletor, runner.services.databaseFilters, runner.repositories.translations, runner.repositories.userRights); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.userAddress, err = address.New(runner.databases.selecter, runner.databases.updater, runner.databases.deletor, runner.repositories.translations, runner.repositories.countries, runner.services.logger); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.userContact, err = contact.New(runner.databases.selecter, runner.databases.deletor, runner.repositories.translations); err != nil {
		return errors.Trace(err)
	}
	if runner.stores.forum, err = forum.New(runner.databases.selecter, runner.databases.deletor); err != nil {
		return errors.Trace(err)
	}
	return nil
}
