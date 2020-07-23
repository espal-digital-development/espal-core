package modules

import (
	"github.com/espal-digital-development/espal-core/modules/assets"
	"github.com/espal-digital-development/espal-core/modules/config"
	"github.com/espal-digital-development/espal-core/modules/databasemigrations"
	"github.com/espal-digital-development/espal-core/modules/meta"
	"github.com/espal-digital-development/espal-core/modules/repositories"
	"github.com/espal-digital-development/espal-core/modules/routes"
	"github.com/espal-digital-development/espal-core/modules/translations"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/juju/errors"
)

var _ Modular = &Module{}

// Modular represents an object that provides informational objects that define the details and behavior of a module.
type Modular interface {
	GetMeta() *meta.Meta

	SetConfig(config *config.Config)
	GetConfig() (*config.Config, error)
	SetDatabaseMigrations(databasemigrations *databasemigrations.DatabaseMigrations)
	GetDatabaseMigrations() (*databasemigrations.DatabaseMigrations, error)
	SetAssets(assets *assets.Assets)
	GetAssets() (*assets.Assets, error)
	SetRoutes(routes *routes.Routes)
	GetRoutes() (*routes.Routes, error)
	SetTranslations(translations *translations.Translations)
	GetTranslations() (*translations.Translations, error)
	SetRepositories(repositories *repositories.Repositories)
	GetRepositories() (*repositories.Repositories, error)

	RegisterValidatorsFactory(validatorsFactory validators.Factory)
	GetValidatorsFactory() validators.Factory
	RegisterCoreStores(stores Stores)
	GetStores() Stores
}

// Module object.
type Module struct {
	metaDefinition             *meta.Meta
	configProvider             *config.Config
	databaseMigrationsProvider *databasemigrations.DatabaseMigrations
	assetsProvider             *assets.Assets
	routesProvider             *routes.Routes
	translationsProvider       *translations.Translations
	repositoriesProvider       *repositories.Repositories

	preGetConfigCallback             func(m Modular) error
	preGetDatabaseMigrationsCallback func(m Modular) error
	preGetAssetsCallback             func(m Modular) error
	preGetRoutesCallback             func(m Modular) error
	preGetTranslationsCallback       func(m Modular) error
	preGetRepositoriesCallback       func(m Modular) error

	validatorsFactory validators.Factory
	stores            Stores
}

// Config Module configuration object.
type Config struct {
	MetaDefinition *meta.Meta

	PreGetConfigCallback             func(m Modular) error
	PreGetDatabaseMigrationsCallback func(m Modular) error
	PreGetAssetsCallback             func(m Modular) error
	PreGetRoutesCallback             func(m Modular) error
	PreGetTranslationsCallback       func(m Modular) error
	PreGetRepositoriesCallback       func(m Modular) error
}

// GetMeta gets the meta definition.
func (m *Module) GetMeta() *meta.Meta {
	return m.metaDefinition
}

// SetConfig sets the config provider.
func (m *Module) SetConfig(config *config.Config) {
	m.configProvider = config
}

// GetConfig gets the config provider.
func (m *Module) GetConfig() (*config.Config, error) {
	if err := m.preGetConfigCallback(m); err != nil {
		return nil, errors.Trace(err)
	}
	return m.configProvider, nil
}

// SetDatabaseMigrations sets the database migration provider.
func (m *Module) SetDatabaseMigrations(databasemigrations *databasemigrations.DatabaseMigrations) {
	m.databaseMigrationsProvider = databasemigrations
}

// GetDatabaseMigrations gets the database migration provider.
func (m *Module) GetDatabaseMigrations() (*databasemigrations.DatabaseMigrations, error) {
	if err := m.preGetDatabaseMigrationsCallback(m); err != nil {
		return nil, errors.Trace(err)
	}
	return m.databaseMigrationsProvider, nil
}

// SetAssets sets the assets provider.
func (m *Module) SetAssets(assets *assets.Assets) {
	m.assetsProvider = assets
}

// GetAssets gets the assets provider.
func (m *Module) GetAssets() (*assets.Assets, error) {
	if err := m.preGetAssetsCallback(m); err != nil {
		return nil, errors.Trace(err)
	}
	return m.assetsProvider, nil
}

// SetRoutes sets the routes provider.
func (m *Module) SetRoutes(routes *routes.Routes) {
	m.routesProvider = routes
}

// GetRoutes gets the routes provider.
func (m *Module) GetRoutes() (*routes.Routes, error) {
	if err := m.preGetRoutesCallback(m); err != nil {
		return nil, errors.Trace(err)
	}
	return m.routesProvider, nil
}

// SetTranslations sets the translations provider.
func (m *Module) SetTranslations(translations *translations.Translations) {
	m.translationsProvider = translations
}

// GetTranslations gets the translations provider.
func (m *Module) GetTranslations() (*translations.Translations, error) {
	if err := m.preGetTranslationsCallback(m); err != nil {
		return nil, errors.Trace(err)
	}
	return m.translationsProvider, nil
}

// SetRepositories sets the repositories provider.
func (m *Module) SetRepositories(repositories *repositories.Repositories) {
	m.repositoriesProvider = repositories
}

// GetRepositories gets the repositories provider.
func (m *Module) GetRepositories() (*repositories.Repositories, error) {
	if err := m.preGetRepositoriesCallback(m); err != nil {
		return nil, errors.Trace(err)
	}
	return m.repositoriesProvider, nil
}

// RegisterValidator registers the that is provided by the core to this module.
func (m *Module) RegisterValidatorsFactory(validatorsFactory validators.Factory) {
	m.validatorsFactory = validatorsFactory
}

// GetValidatorsFactory fetches the internally registered validators factory.
func (m *Module) GetValidatorsFactory() validators.Factory {
	return m.validatorsFactory
}

// RegisterCoreStores registers all stores that are provided by the core to this module.
func (m *Module) RegisterCoreStores(stores Stores) {
	m.stores = stores
}

// GetStores fetches the internally registered stores.
func (m *Module) GetStores() Stores {
	return m.stores
}

// New returns a new instance of Module.
func New(config *Config) (*Module, error) {
	m := &Module{
		metaDefinition:                   config.MetaDefinition,
		preGetConfigCallback:             config.PreGetConfigCallback,
		preGetDatabaseMigrationsCallback: config.PreGetDatabaseMigrationsCallback,
		preGetAssetsCallback:             config.PreGetAssetsCallback,
		preGetRoutesCallback:             config.PreGetRoutesCallback,
		preGetTranslationsCallback:       config.PreGetTranslationsCallback,
		preGetRepositoriesCallback:       config.PreGetRepositoriesCallback,
	}
	return m, nil
}
