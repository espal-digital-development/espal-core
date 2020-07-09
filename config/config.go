package config

import (
	"strconv"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
	yaml "gopkg.in/yaml.v2"
)

var _ Config = &Configuration{}

var (
	// ErrorIncorrectDefaultLanguage is returned when no default language was specified in the c.
	ErrorIncorrectDefaultLanguage = errors.New("default Language must be available in the available languages file")
	// ErrorDontUsePort80ForTLS is returned when the default non-TLS port is used for TLS connections.
	ErrorDontUsePort80ForTLS = errors.New("do not use port 80 for TLS connections")
	// ErrorAssetCacheMaxAgeIncorrect is returned when the assetCacheMaxAge is malformed.
	ErrorAssetCacheMaxAgeIncorrect = errors.New("configuration value assetCacheMaxAge should be empty or higher than 0")
)

// Config object to represent the Config interaction.
type Config interface {
	General
	Server
	Database
	Email
	Security
	Session
	URLs
	Assets
	Paths
	AvailableLanguages() []string
	LanguageIsAvailable(code string) bool
}

type configYaml struct {
	General  general
	Server   server
	Database database
	Email    email
	Security security
	Session  session
	Urls     urls
	Assets   assets
	Paths    paths
}

// Configuration service object.
type Configuration struct {
	general     general
	server      server
	database    database
	email       email
	security    security
	session     session
	urls        urls
	assets      assets
	paths       paths
	coreStorage storage.Storage
}

// AvailableLanguages returns all languages that are available in the current build based on the c.
func (c *Configuration) AvailableLanguages() []string {
	availableLanguages := make([]string, 0)
	availableLanguages = append(availableLanguages, c.general.Languages...)
	return availableLanguages
}

// LanguageIsAvailable indicates if the language is available in the current build based on the c.
func (c *Configuration) LanguageIsAvailable(code string) bool {
	for _, language := range c.general.Languages {
		if code == language {
			return true
		}
	}
	return false
}

func (c *Configuration) validate() error {
	if c.assets.CacheMaxAge != "" {
		assetCacheMaxAge, err := strconv.Atoi(c.assets.CacheMaxAge)
		if err != nil || assetCacheMaxAge < 1 {
			return errors.Trace(ErrorAssetCacheMaxAgeIncorrect)
		}
	}
	availableLanguages := map[string]bool{}
	for _, language := range c.general.Languages {
		availableLanguages[language] = true
	}
	if ok := availableLanguages[c.general.DefaultLanguage]; !ok {
		return errors.Trace(ErrorIncorrectDefaultLanguage)
	}
	if c.server.Port == defaultHTTPPort {
		return errors.Trace(ErrorDontUsePort80ForTLS)
	}
	if []byte(c.urls.Admin)[0] != '/' {
		c.urls.Admin = "/" + c.urls.Admin
	}
	if []byte(c.urls.Pprof)[0] != '/' {
		c.urls.Pprof = "/" + c.urls.Pprof
	}
	return nil
}

// New returns a new instance of Configuration loaded from the target file.
func New(coreStorage storage.Storage) (*Configuration, error) {
	c := &Configuration{
		coreStorage: coreStorage,
	}

	configBytes, ok, err := coreStorage.Get("config.yml")
	if err != nil {
		return nil, errors.Trace(err)
	}
	if !ok {
		return nil, errors.Errorf(
			"no config.yml file found. Please create one. The espal-run command can help automate this process")
	}

	values := c.getDefaultYaml()
	if err := yaml.Unmarshal(configBytes, &values); err != nil {
		return nil, errors.Trace(err)
	}

	c.assets = values.Assets
	c.database = values.Database
	c.email = values.Email
	c.general = values.General
	c.paths = values.Paths
	c.security = values.Security
	c.server = values.Server
	c.session = values.Session
	c.urls = values.Urls

	if err := c.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	return c, nil
}
