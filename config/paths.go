package config

// Paths config section.
type Paths interface {
	ServerSSLCertificateFilePath() string
	ServerSSLKeyFilePath() string

	DatabaseSSLRootCertificateFile() string
	DatabaseSelecterSSLCertificateFile() string
	DatabaseSelecterSSLKeyFile() string
	DatabaseCreatorSSLCertificateFile() string
	DatabaseCreatorSSLKeyFile() string
	DatabaseInserterSSLCertificateFile() string
	DatabaseInserterSSLKeyFile() string
	DatabaseUpdaterSSLCertificateFile() string
	DatabaseUpdaterSSLKeyFile() string
	DatabaseDeletorSSLCertificateFile() string
	DatabaseDeletorSSLKeyFile() string
	DatabaseMigratorSSLCertificateFile() string
	DatabaseMigratorSSLKeyFile() string

	StylesheetsAssetsPath() string
	JavaScriptAssetsPath() string
	ImagesAssetsPath() string
	PublicRootFilesAssetsPath() string
	PublicFilesAssetsPath() string
	PrivateFilesAssetsPath() string

	TranslationsPath() string
}

type pathsDatabaseUser struct {
	SSLCertificateFile string `yaml:"sslCertificateFile"`
	SSLKeyFile         string `yaml:"sslKeyFile"`
}

type paths struct {
	Server struct {
		SSLCertificateFile string `yaml:"sslCertificateFile"`
		SSLKeyFile         string `yaml:"sslKeyFile"`
	}
	Database struct {
		SSLRootCertificateFile string `yaml:"sslRootCertificateFile"`
		Selecter               pathsDatabaseUser
		Creator                pathsDatabaseUser
		Inserter               pathsDatabaseUser
		Updater                pathsDatabaseUser
		Deletor                pathsDatabaseUser
		Migrator               pathsDatabaseUser
	}
	Assets struct {
		Stylesheets     string
		JavaScript      string
		Images          string
		PublicRootFiles string `yaml:"publicRootFiles"`
		PublicFiles     string `yaml:"publicFiles"`
		PrivateFiles    string `yaml:"privateFiles"`
	}
	Translations string
}

// ServerSSLCertificateFilePath returns the SSL certificate file path.
func (configuration *Configuration) ServerSSLCertificateFilePath() string {
	return configuration.paths.Server.SSLCertificateFile
}

// ServerSSLKeyFilePath returns the SSL key file path.
func (configuration *Configuration) ServerSSLKeyFilePath() string {
	return configuration.paths.Server.SSLKeyFile
}

// DatabaseSSLRootCertificateFile returns the path for the database root's SSL Certificate File.
func (configuration *Configuration) DatabaseSSLRootCertificateFile() string {
	return configuration.paths.Database.SSLRootCertificateFile
}

// DatabaseSelecterSSLCertificateFile returns the path for the selecter's SSL Certificate File.
func (configuration *Configuration) DatabaseSelecterSSLCertificateFile() string {
	return configuration.paths.Database.Selecter.SSLCertificateFile
}

// DatabaseSelecterSSLKeyFile returns the path for the selecter's SSL Key File.
func (configuration *Configuration) DatabaseSelecterSSLKeyFile() string {
	return configuration.paths.Database.Selecter.SSLKeyFile
}

// DatabaseCreatorSSLCertificateFile returns the path for the creator's SSL Certificate File.
func (configuration *Configuration) DatabaseCreatorSSLCertificateFile() string {
	return configuration.paths.Database.Creator.SSLCertificateFile
}

// DatabaseCreatorSSLKeyFile returns the path for the creator's SSL Key File.
func (configuration *Configuration) DatabaseCreatorSSLKeyFile() string {
	return configuration.paths.Database.Creator.SSLKeyFile
}

// DatabaseInserterSSLCertificateFile returns the path for the inserter's SSL Certificate File.
func (configuration *Configuration) DatabaseInserterSSLCertificateFile() string {
	return configuration.paths.Database.Inserter.SSLCertificateFile
}

// DatabaseInserterSSLKeyFile returns the path for the inserter's SSL Key File.
func (configuration *Configuration) DatabaseInserterSSLKeyFile() string {
	return configuration.paths.Database.Inserter.SSLKeyFile
}

// DatabaseUpdaterSSLCertificateFile returns the path for the updater's SSL Certificate File.
func (configuration *Configuration) DatabaseUpdaterSSLCertificateFile() string {
	return configuration.paths.Database.Updater.SSLCertificateFile
}

// DatabaseUpdaterSSLKeyFile returns the path for the updater's SSL Key File.
func (configuration *Configuration) DatabaseUpdaterSSLKeyFile() string {
	return configuration.paths.Database.Updater.SSLKeyFile
}

// DatabaseDeletorSSLCertificateFile returns the path for the deletor's SSL Certificate File.
func (configuration *Configuration) DatabaseDeletorSSLCertificateFile() string {
	return configuration.paths.Database.Deletor.SSLCertificateFile
}

// DatabaseDeletorSSLKeyFile returns the path for the deletor's SSL Key File.
func (configuration *Configuration) DatabaseDeletorSSLKeyFile() string {
	return configuration.paths.Database.Deletor.SSLKeyFile
}

// DatabaseMigratorSSLCertificateFile returns the path for the migrator's SSL Certificate File.
func (configuration *Configuration) DatabaseMigratorSSLCertificateFile() string {
	return configuration.paths.Database.Migrator.SSLCertificateFile
}

// DatabaseMigratorSSLKeyFile returns the path for the migrator's SSL Key File.
func (configuration *Configuration) DatabaseMigratorSSLKeyFile() string {
	return configuration.paths.Database.Migrator.SSLKeyFile
}

// StylesheetsAssetsPath returns the stylesheets path.
func (configuration *Configuration) StylesheetsAssetsPath() string {
	return configuration.paths.Assets.Stylesheets
}

// JavaScriptAssetsPath returns the javaScript path.
func (configuration *Configuration) JavaScriptAssetsPath() string {
	return configuration.paths.Assets.JavaScript
}

// ImagesAssetsPath returns the images path.
func (configuration *Configuration) ImagesAssetsPath() string {
	return configuration.paths.Assets.Images
}

// PublicRootFilesAssetsPath returns the public root-files path.
// The difference from the normal public files is that
// these will be forced on the root path of the serve.
func (configuration *Configuration) PublicRootFilesAssetsPath() string {
	return configuration.paths.Assets.PublicRootFiles
}

// PublicFilesAssetsPath returns the public files path.
func (configuration *Configuration) PublicFilesAssetsPath() string {
	return configuration.paths.Assets.PublicFiles
}

// PrivateFilesAssetsPath returns the private files path.
func (configuration *Configuration) PrivateFilesAssetsPath() string {
	return configuration.paths.Assets.PrivateFiles
}

// TranslationsPath returns the translations path.
func (configuration *Configuration) TranslationsPath() string {
	return configuration.paths.Translations
}
