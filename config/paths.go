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
func (c *Configuration) ServerSSLCertificateFilePath() string {
	return c.paths.Server.SSLCertificateFile
}

// ServerSSLKeyFilePath returns the SSL key file path.
func (c *Configuration) ServerSSLKeyFilePath() string {
	return c.paths.Server.SSLKeyFile
}

// DatabaseSSLRootCertificateFile returns the path for the database root's SSL Certificate File.
func (c *Configuration) DatabaseSSLRootCertificateFile() string {
	return c.paths.Database.SSLRootCertificateFile
}

// DatabaseSelecterSSLCertificateFile returns the path for the selecter's SSL Certificate File.
func (c *Configuration) DatabaseSelecterSSLCertificateFile() string {
	return c.paths.Database.Selecter.SSLCertificateFile
}

// DatabaseSelecterSSLKeyFile returns the path for the selecter's SSL Key File.
func (c *Configuration) DatabaseSelecterSSLKeyFile() string {
	return c.paths.Database.Selecter.SSLKeyFile
}

// DatabaseCreatorSSLCertificateFile returns the path for the creator's SSL Certificate File.
func (c *Configuration) DatabaseCreatorSSLCertificateFile() string {
	return c.paths.Database.Creator.SSLCertificateFile
}

// DatabaseCreatorSSLKeyFile returns the path for the creator's SSL Key File.
func (c *Configuration) DatabaseCreatorSSLKeyFile() string {
	return c.paths.Database.Creator.SSLKeyFile
}

// DatabaseInserterSSLCertificateFile returns the path for the inserter's SSL Certificate File.
func (c *Configuration) DatabaseInserterSSLCertificateFile() string {
	return c.paths.Database.Inserter.SSLCertificateFile
}

// DatabaseInserterSSLKeyFile returns the path for the inserter's SSL Key File.
func (c *Configuration) DatabaseInserterSSLKeyFile() string {
	return c.paths.Database.Inserter.SSLKeyFile
}

// DatabaseUpdaterSSLCertificateFile returns the path for the updater's SSL Certificate File.
func (c *Configuration) DatabaseUpdaterSSLCertificateFile() string {
	return c.paths.Database.Updater.SSLCertificateFile
}

// DatabaseUpdaterSSLKeyFile returns the path for the updater's SSL Key File.
func (c *Configuration) DatabaseUpdaterSSLKeyFile() string {
	return c.paths.Database.Updater.SSLKeyFile
}

// DatabaseDeletorSSLCertificateFile returns the path for the deletor's SSL Certificate File.
func (c *Configuration) DatabaseDeletorSSLCertificateFile() string {
	return c.paths.Database.Deletor.SSLCertificateFile
}

// DatabaseDeletorSSLKeyFile returns the path for the deletor's SSL Key File.
func (c *Configuration) DatabaseDeletorSSLKeyFile() string {
	return c.paths.Database.Deletor.SSLKeyFile
}

// DatabaseMigratorSSLCertificateFile returns the path for the migrator's SSL Certificate File.
func (c *Configuration) DatabaseMigratorSSLCertificateFile() string {
	return c.paths.Database.Migrator.SSLCertificateFile
}

// DatabaseMigratorSSLKeyFile returns the path for the migrator's SSL Key File.
func (c *Configuration) DatabaseMigratorSSLKeyFile() string {
	return c.paths.Database.Migrator.SSLKeyFile
}

// StylesheetsAssetsPath returns the stylesheets path.
func (c *Configuration) StylesheetsAssetsPath() string {
	return c.paths.Assets.Stylesheets
}

// JavaScriptAssetsPath returns the javaScript path.
func (c *Configuration) JavaScriptAssetsPath() string {
	return c.paths.Assets.JavaScript
}

// ImagesAssetsPath returns the images path.
func (c *Configuration) ImagesAssetsPath() string {
	return c.paths.Assets.Images
}

// PublicRootFilesAssetsPath returns the public root-files path.
// The difference from the normal public files is that these will be forced on the root path of the serve.
func (c *Configuration) PublicRootFilesAssetsPath() string {
	return c.paths.Assets.PublicRootFiles
}

// PublicFilesAssetsPath returns the public files path.
func (c *Configuration) PublicFilesAssetsPath() string {
	return c.paths.Assets.PublicFiles
}

// PrivateFilesAssetsPath returns the private files path.
func (c *Configuration) PrivateFilesAssetsPath() string {
	return c.paths.Assets.PrivateFiles
}

// TranslationsPath returns the translations path.
func (c *Configuration) TranslationsPath() string {
	return c.paths.Translations
}
