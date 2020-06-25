package config_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/storage/storagemock"
	"github.com/espal-digital-development/espal-core/testtools"
	"github.com/juju/errors"
)

var (
	configYml   []byte
	coreStorage *storagemock.StorageMock
)

func initMocks() {
	configYml = configYmlBlueprint
	coreStorage = &storagemock.StorageMock{
		GetFunc: func(key string) ([]byte, bool, error) {
			return configYml, true, nil
		},
	}
}

func TestNew(t *testing.T) {
	initMocks()
	config, err := config.New(coreStorage)
	if err != nil {
		t.Fatal(err)
	}
	if config == nil {
		t.Fatal("config should not be nil")
	}
}

func TestYamlFileError(t *testing.T) {
	initMocks()
	yamlFileError := errors.New("yamlFileError")
	coreStorage.GetFunc = func(key string) ([]byte, bool, error) {
		return nil, false, yamlFileError
	}
	config, err := config.New(coreStorage)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "yamlFileError", err, yamlFileError)
	if config != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestYamlFileNotExist(t *testing.T) {
	initMocks()
	yamlFileNotExistError := errors.New("no config.yml file found. Please create one. The espal-run command can help automate this process")
	coreStorage.GetFunc = func(key string) ([]byte, bool, error) {
		return nil, false, nil
	}
	config, err := config.New(coreStorage)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "yamlFileNotExistError", err, yamlFileNotExistError)
	if config != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestYamlUnmarshalError(t *testing.T) {
	initMocks()
	brokenYaml := bytes.Replace(configYmlBlueprint, []byte(":"), []byte("%"), 1)
	coreStorage.GetFunc = func(key string) ([]byte, bool, error) {
		return brokenYaml, true, nil
	}
	config, err := config.New(coreStorage)
	if err == nil {
		t.Fatal("Should give an error")
	}
	// A bit literal, but it's ok until it fails on adjustment
	testtools.EqError(t, "yamlFileError", err, errors.New("yaml: line 3: mapping values are not allowed in this context"))
	if config != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestValidateCacheMaxAgeFailure(t *testing.T) {
	initMocks()
	incorrectYaml := bytes.Replace(configYmlBlueprint, []byte("cacheMaxAge: 60"), []byte("cacheMaxAge: 0"), 1)
	coreStorage.GetFunc = func(key string) ([]byte, bool, error) {
		return incorrectYaml, true, nil
	}
	conf, err := config.New(coreStorage)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "cacheMaxAgeFailure", err, config.ErrorAssetCacheMaxAgeIncorrect)
	if conf != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestValidateDontUsePort80ForTLSFailure(t *testing.T) {
	initMocks()
	incorrectYaml := bytes.Replace(configYmlBlueprint, []byte("  port: 8443"), []byte("  port: 80"), 1)
	coreStorage.GetFunc = func(key string) ([]byte, bool, error) {
		return incorrectYaml, true, nil
	}
	conf, err := config.New(coreStorage)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "dontUsePort80ForTLS", err, config.ErrorDontUsePort80ForTLS)
	if conf != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestValidateIncorrectDefaultLanguageFailure(t *testing.T) {
	initMocks()
	incorrectYaml := bytes.Replace(configYmlBlueprint, []byte("  defaultLanguage: en"), []byte("  defaultLanguage: xx"), 1)
	coreStorage.GetFunc = func(key string) ([]byte, bool, error) {
		return incorrectYaml, true, nil
	}
	conf, err := config.New(coreStorage)
	if err == nil {
		t.Fatal("Should give an error")
	}
	testtools.EqError(t, "incorrectDefaultLanguage", err, config.ErrorIncorrectDefaultLanguage)
	if conf != nil {
		t.Fatal("config should be nil when an error is thrown")
	}
}

func TestAvailableLanguages(t *testing.T) {
	initMocks()
	config, err := config.New(coreStorage)
	if err != nil {
		t.Fatal(err)
	}

	langs := config.AvailableLanguages()
	if len(langs) == 0 {
		t.Fatal("Should not be empty")
	}

	for k := range langs {
		if len(strings.TrimSpace(langs[k])) == 0 {
			t.Fatal("Language code keys should not be empty")
		}
	}
}

func TestLanguageIsAvailable(t *testing.T) {
	initMocks()
	config, err := config.New(coreStorage)
	if err != nil {
		t.Fatal(err)
	}

	langs := config.AvailableLanguages()
	for k := range langs {
		if !config.LanguageIsAvailable(langs[k]) {
			t.Errorf("should be able to find `%s`", langs[k])
		}
	}
}

func TestLanguageIsNotAvailable(t *testing.T) {
	initMocks()
	config, err := config.New(coreStorage)
	if err != nil {
		t.Fatal(err)
	}

	if config.LanguageIsAvailable("icantexist") {
		t.Fatal("impossible language shouldn't give a result")
	}
}

func TestConfigCallers(t *testing.T) {
	initMocks()
	config, err := config.New(coreStorage)
	if err != nil {
		t.Fatal(err)
	}

	// Match all the blueprinted yaml values
	testtools.EqString(t, "general.defautLangauge", config.DefaultLanguage(), "en")
	testtools.EqBool(t, "general.development", config.Development(), true)
	languages := config.Languages()
	if len(languages) == 0 {
		t.Fatalf("Expected general.languages to not be empty")
	}
	testtools.EqBool(t, "general.logging", config.Logging(), true)
	testtools.EqBool(t, "general.pprof", config.Pprof(), true)

	testtools.EqBool(t, "assets.brotli", config.AssetsBrotli(), true)
	testtools.EqBool(t, "assets.brotliFiles", config.AssetsBrotliFiles(), true)
	testtools.EqBool(t, "assets.gzip", config.AssetsGZip(), true)
	testtools.EqBool(t, "assets.gzipFiles", config.AssetsGZipFiles(), true)
	testtools.EqString(t, "assets.cacheMaxAge", config.AssetsCacheMaxAge(), "60")

	testtools.EqString(t, "database.host", config.DatabaseHost(), "localhost")
	testtools.EqString(t, "database.name", config.DatabaseName(), "espal")
	testtools.EqInt(t, "database.port", config.DatabasePort(), 26257)
	testtools.EqString(t, "database.creator.username", config.DatabaseCreator(), "espal_creator")
	testtools.EqString(t, "database.deletor.username", config.DatabaseDeletor(), "espal_deletor")
	testtools.EqString(t, "database.inserter.username", config.DatabaseInserter(), "espal_inserter")
	testtools.EqString(t, "database.migrator.username", config.DatabaseMigrator(), "espal_migrator")
	testtools.EqString(t, "database.selecter.username", config.DatabaseSelecter(), "espal_selecter")
	testtools.EqString(t, "database.updater.username", config.DatabaseUpdater(), "espal_updater")

	testtools.EqString(t, "email.host", config.EmailHost(), "smtp.domain.dev")
	testtools.EqString(t, "email.noReplyAddress", config.EmailNoReplyAddress(), "noreply@domain.dev")
	testtools.EqString(t, "email.password", config.EmailPassword(), "fakePassword")
	testtools.EqInt(t, "email.port", config.EmailPort(), 2525)
	testtools.EqString(t, "email.username", config.EmailUsername(), "espal")

	testtools.EqString(t, "paths.assets.images", config.ImagesAssetsPath(), "./app/assets/images")
	testtools.EqString(t, "paths.assets.javaScript", config.JavaScriptAssetsPath(), "./app/assets/js")
	testtools.EqString(t, "paths.assets.privateFiles", config.PrivateFilesAssetsPath(), "./app/assets/files/private")
	testtools.EqString(t, "paths.assets.publicFiles", config.PublicFilesAssetsPath(), "./app/assets/files/public")
	testtools.EqString(t, "paths.assets.publicRootFiles", config.PublicRootFilesAssetsPath(), "./app/assets/files/root")
	testtools.EqString(t, "paths.assets.stylesheets", config.StylesheetsAssetsPath(), "./app/assets/css")
	testtools.EqString(t, "paths.database.sslRootCertificateFile", config.DatabaseSSLRootCertificateFile(), "./app/database/ca.crt")
	testtools.EqString(t, "paths.database.selecter.sslCertificateFile", config.DatabaseCreatorSSLCertificateFile(), "./app/database/client.espal_creator.crt")
	testtools.EqString(t, "paths.database.selecter.sslKeyFile", config.DatabaseCreatorSSLKeyFile(), "./app/database/client.espal_creator.key")
	testtools.EqString(t, "paths.database.deletor.sslCertificateFile", config.DatabaseDeletorSSLCertificateFile(), "./app/database/client.espal_deletor.crt")
	testtools.EqString(t, "paths.database.deletor.sslKeyFile", config.DatabaseDeletorSSLKeyFile(), "./app/database/client.espal_deletor.key")
	testtools.EqString(t, "paths.database.inserter.sslCertificateFile", config.DatabaseInserterSSLCertificateFile(), "./app/database/client.espal_inserter.crt")
	testtools.EqString(t, "paths.database.inserter.sslKeyFile", config.DatabaseInserterSSLKeyFile(), "./app/database/client.espal_inserter.key")
	testtools.EqString(t, "paths.database.migrator.sslCertificateFile", config.DatabaseMigratorSSLCertificateFile(), "./app/database/client.espal_migrator.crt")
	testtools.EqString(t, "paths.database.migrator.sslKeyFile", config.DatabaseMigratorSSLKeyFile(), "./app/database/client.espal_migrator.key")
	testtools.EqString(t, "paths.database.selecter.sslCertificateFile", config.DatabaseSelecterSSLCertificateFile(), "./app/database/client.espal_selecter.crt")
	testtools.EqString(t, "paths.database.selecter.sslKeyFile", config.DatabaseSelecterSSLKeyFile(), "./app/database/client.espal_selecter.key")
	testtools.EqString(t, "paths.database.updater.sslCertificateFile", config.DatabaseUpdaterSSLCertificateFile(), "./app/database/client.espal_updater.crt")
	testtools.EqString(t, "paths.database.updater.sslKeyFile", config.DatabaseUpdaterSSLKeyFile(), "./app/database/client.espal_updater.key")
	testtools.EqString(t, "paths.server.sslCertificateFile", config.ServerSSLCertificateFilePath(), "./app/localhost.crt")
	testtools.EqString(t, "paths.server.sslKeyFile", config.ServerSSLKeyFilePath(), "./app/localhost.key")
	testtools.EqString(t, "paths.translations", config.TranslationsPath(), "./app/translations")

	testtools.EqString(t, "urls.admin", config.AdminURL(), "/_adminPath")
	testtools.EqString(t, "urls.pprof", config.PprofURL(), "/_pprofPath")

	testtools.EqInt(t, "security.bcryptRounds", config.SecurityBcryptRounds(), 12)
	testtools.EqDuration(t, "security.formTokenCleanupInterval", config.SecurityFormTokenCleanupInterval(), time.Second*10)
	testtools.EqDuration(t, "security.formTokenLifespan", config.SecurityFormTokenLifespan(), time.Minute*8)
	testtools.EqBool(t, "security.globalAuthentication", config.SecurityGlobalAuthentication(), true)

	testtools.EqString(t, "session.cookieName", config.SessionCookieName(), "s")
	testtools.EqDuration(t, "session.expiration", config.SessionExpiration(), time.Minute*45)
	testtools.EqDuration(t, "session.rememberMeExpiration", config.SessionRememberMeExpiration(), time.Hour*720)

	testtools.EqInt(t, "server.httpRedirectPort", config.ServerHTTPRedirectPort(), 8080)
	testtools.EqString(t, "server.host", config.ServerHost(), "localhost")
	testtools.EqInt(t, "server.port", config.ServerPort(), 8443)
}
