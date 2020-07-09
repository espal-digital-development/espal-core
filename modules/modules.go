package modules

import (
	"github.com/espal-digital-development/espal-core/modules/assets"
	"github.com/espal-digital-development/espal-core/modules/config"
	"github.com/espal-digital-development/espal-core/modules/databasemigrations"
	"github.com/espal-digital-development/espal-core/modules/meta"
	"github.com/espal-digital-development/espal-core/modules/pages"
	"github.com/espal-digital-development/espal-core/modules/repositories"
	"github.com/espal-digital-development/espal-core/modules/routes"
	"github.com/espal-digital-development/espal-core/modules/translations"
)

var _ Modular = &Module{}

// Modular represents an object that provides informational objects
// that define the details and behavior of a module.
type Modular interface {
	GetMeta() *meta.Meta
	GetConfig() *config.Config
	GetDatabaseMigrations() *databasemigrations.DatabaseMigrations
	GetAssets() *assets.Assets
	GetPages() *pages.Pages
	GetRoutes() *routes.Routes
	GetTranslations() *translations.Translations
	GetRepositories() *repositories.Repositories
}

// Module object.
type Module struct {
	metaDefinition             *meta.Meta
	configProvider             *config.Config
	databaseMigrationsProvider *databasemigrations.DatabaseMigrations
	assetsProvider             *assets.Assets
	pagesProvider              *pages.Pages
	routesProvider             *routes.Routes
	translationsProvider       *translations.Translations
	repositoriesProvider       *repositories.Repositories
}

// Config Module configuration object.
type Config struct {
	MetaDefinition             *meta.Meta
	ConfigProvider             *config.Config
	DatabaseMigrationsProvider *databasemigrations.DatabaseMigrations
	AssetsProvider             *assets.Assets
	PagesProvider              *pages.Pages
	RoutesProvider             *routes.Routes
	TranslationsProvider       *translations.Translations
	RepositoriesProvider       *repositories.Repositories
}

// GetMeta gets the meta definition.
func (m *Module) GetMeta() *meta.Meta {
	return m.metaDefinition
}

// GetConfig gets the config provider.
func (m *Module) GetConfig() *config.Config {
	return m.configProvider
}

// GetDatabaseMigrations gets the database migration provider.
func (m *Module) GetDatabaseMigrations() *databasemigrations.DatabaseMigrations {
	return m.databaseMigrationsProvider
}

// GetAssets gets the assets provider.
func (m *Module) GetAssets() *assets.Assets {
	return m.assetsProvider
}

// GetPages gets the pages provider.
func (m *Module) GetPages() *pages.Pages {
	return m.pagesProvider
}

// GetRoutes gets the routes provider.
func (m *Module) GetRoutes() *routes.Routes {
	return m.routesProvider
}

// GetTranslations gets the translations provider.
func (m *Module) GetTranslations() *translations.Translations {
	return m.translationsProvider
}

// GetRepositories gets the repositories provider.
func (m *Module) GetRepositories() *repositories.Repositories {
	return m.repositoriesProvider
}

// New returns a new instance of Module.
func New(config *Config) (*Module, error) {
	m := &Module{
		metaDefinition:             config.MetaDefinition,
		configProvider:             config.ConfigProvider,
		databaseMigrationsProvider: config.DatabaseMigrationsProvider,
		assetsProvider:             config.AssetsProvider,
		pagesProvider:              config.PagesProvider,
		routesProvider:             config.RoutesProvider,
		translationsProvider:       config.TranslationsProvider,
		repositoriesProvider:       config.RepositoriesProvider,
	}
	return m, nil
}
