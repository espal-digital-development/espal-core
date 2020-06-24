package config

import (
	"strconv"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
	yaml "gopkg.in/yaml.v2"
)

const (
	DefaultHTTPPort = 80
)

var _ Config = &Configuration{}

var (
	// ErrorIncorrectDefaultLanguage is returned when no default language was specified in the configuration.
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

// AvailableLanguages returns all languages that are available in the current build based on the configuration.
func (configuration *Configuration) AvailableLanguages() []string {
	availableLanguages := make([]string, 0)
	availableLanguages = append(availableLanguages, configuration.general.Languages...)
	return availableLanguages
}

// LanguageIsAvailable indicates if the language is available in the current build based on the configuration.
func (configuration *Configuration) LanguageIsAvailable(code string) bool {
	for _, language := range configuration.general.Languages {
		if code == language {
			return true
		}
	}
	return false
}

func (configuration *Configuration) validate() error {
	if configuration.assets.CacheMaxAge != "" {
		assetCacheMaxAge, err := strconv.Atoi(configuration.assets.CacheMaxAge)
		if err != nil || assetCacheMaxAge < 1 {
			return errors.Trace(ErrorAssetCacheMaxAgeIncorrect)
		}
	}
	availableLanguages := map[string]bool{}
	for _, language := range configuration.general.Languages {
		availableLanguages[language] = true
	}
	if ok := availableLanguages[configuration.general.DefaultLanguage]; !ok {
		return errors.Trace(ErrorIncorrectDefaultLanguage)
	}
	if configuration.server.Port == DefaultHTTPPort {
		return errors.Trace(ErrorDontUsePort80ForTLS)
	}
	if []byte(configuration.urls.Admin)[0] != '/' {
		configuration.urls.Admin = "/" + configuration.urls.Admin
	}
	if []byte(configuration.urls.Pprof)[0] != '/' {
		configuration.urls.Pprof = "/" + configuration.urls.Pprof
	}
	return nil
}

// New returns a new instance of Configuration loaded from the target file.
func New(coreStorage storage.Storage) (*Configuration, error) {
	configuration := &Configuration{
		coreStorage: coreStorage,
	}

	configBytes, ok, err := coreStorage.Get("config.yml")
	if err != nil {
		return nil, errors.Trace(err)
	}
	if !ok {
		return nil, errors.Errorf("no config.yml file found. Please create one. The espal-run command can help automate this process")
	}

	var values configYaml
	if err := yaml.Unmarshal(configBytes, &values); err != nil {
		return nil, errors.Trace(err)
	}

	configuration.assets = values.Assets
	configuration.database = values.Database
	configuration.email = values.Email
	configuration.general = values.General
	configuration.paths = values.Paths
	configuration.security = values.Security
	configuration.server = values.Server
	configuration.session = values.Session
	configuration.urls = values.Urls

	if err := configuration.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	return configuration, nil
}
