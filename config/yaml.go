package config

import (
	"bytes"
	"fmt"
	"os/exec"
)

const notInstalledDisableErrBlueprint = "%s is enabled, but `%s` wasn't found. Disabling..\n"

func (c *Configuration) isInstalled(name string) bool {
	out, _ := exec.Command("which", name).CombinedOutput()
	return bytes.Contains(out, []byte("/"+name))
}

// nolint:funlen
// TODO :: Use Logger here for the fmt.Printf's.
func (c *Configuration) getDefaultYaml() configYaml {
	assetsOptimizePngs := defaultAssetsOptimizePngs
	assetsOptimizeJpegs := defaultAssetsOptimizeJpegs
	assetsOptimizeGifs := defaultAssetsOptimizeGifs
	assetsOptimizeSvgs := defaultAssetsOptimizeSvgs
	if assetsOptimizePngs && !c.isInstalled("pngquant") {
		fmt.Printf(notInstalledDisableErrBlueprint, "assets.optimizePngs", "pngquant")
		assetsOptimizePngs = false
	}
	if assetsOptimizeJpegs && !c.isInstalled("jpegoptim") {
		fmt.Printf(notInstalledDisableErrBlueprint, "assets.optimizeJpegs", "jpegoptim")
		assetsOptimizeJpegs = false
	}
	if assetsOptimizeGifs && !c.isInstalled("gifsicle") {
		fmt.Printf(notInstalledDisableErrBlueprint, "assets.optimizeGifs", "gifsicle")
		assetsOptimizeGifs = false
	}
	if assetsOptimizeSvgs && !c.isInstalled("svgo") {
		fmt.Printf(notInstalledDisableErrBlueprint, "assets.optimizeSvgs", "svgo")
		assetsOptimizeSvgs = false
	}

	return configYaml{
		General: general{
			Development:     false,
			Logging:         false,
			Pprof:           false,
			Languages:       []string{"en"},
			DefaultLanguage: "en",
		},
		Server: server{
			Host:             defaultServerHost,
			Port:             defaultServerPort,
			HTTPRedirectPort: defaultServerHTTPRedirectPort,
		},
		Database: database{
			Host: defaultDatabaseHost,
			Port: defaultDatabasePort,
			Name: defaultDatabaseName,
			Users: struct {
				Selecter string
				Creator  string
				Inserter string
				Updater  string
				Deletor  string
				Migrator string
			}{
				Selecter: defaultDatabaseUsersSelecter,
				Creator:  defaultDatabaseUsersCreator,
				Inserter: defaultDatabaseUsersInserter,
				Updater:  defaultDatabaseUsersUpdater,
				Deletor:  defaultDatabaseUsersDeletor,
				Migrator: defaultDatabaseUsersMigrator,
			},
		},
		Security: security{
			BcryptRounds:             defaultSecurityBcryptRounds,
			FormTokenLifespan:        defaultSecurityFormTokenLifespan,
			FormTokenCleanupInterval: defaultSecurityFormTokenCleanupInterval,
			JWTPassword:              defaultJWTPassword,
		},
		Session: session{
			CookieName:           defaultSessionCookieName,
			Expiration:           defaultSessionExpiration,
			RememberMeExpiration: defaultSessionRememberMeExpiration,
		},
		Assets: assets{
			Brotli:        defaultAssetsBrotli,
			Gzip:          defaultAssetsGzip,
			BrotliFiles:   defaultAssetsBrotliFiles,
			OptimizePngs:  assetsOptimizePngs,
			OptimizeJpegs: assetsOptimizeJpegs,
			OptimizeGifs:  assetsOptimizeGifs,
			OptimizeSvgs:  assetsOptimizeSvgs,
			GzipFiles:     defaultAssetsGzipFiles,
			CacheMaxAge:   defaultAssetsCacheMaxAge,
		},
		Paths: paths{
			Server: struct {
				SSLCertificateFile string `yaml:"sslCertificateFile"`
				SSLKeyFile         string `yaml:"sslKeyFile"`
			}{
				SSLCertificateFile: defaultPathsServerSSLCertificateFile,
				SSLKeyFile:         defaultPathsServerSSLKeyFile,
			},
			Database: struct {
				SSLRootCertificateFile string `yaml:"sslRootCertificateFile"`
				Selecter               pathsDatabaseUser
				Creator                pathsDatabaseUser
				Inserter               pathsDatabaseUser
				Updater                pathsDatabaseUser
				Deletor                pathsDatabaseUser
				Migrator               pathsDatabaseUser
			}{
				SSLRootCertificateFile: defaultPathsDatabaseSSLRootCertificateFile,
				Selecter: pathsDatabaseUser{
					SSLCertificateFile: defaultPathsDatabaseSelecterSSLCertificateFile,
					SSLKeyFile:         defaultPathsDatabaseSelecterSSLKeyFile,
				},
				Creator: pathsDatabaseUser{
					SSLCertificateFile: defaultPathsDatabaseCreatorSSLCertificateFile,
					SSLKeyFile:         defaultPathsDatabaseCreatorSSLKeyFile,
				},
				Inserter: pathsDatabaseUser{
					SSLCertificateFile: defaultPathsDatabaseInserterSSLCertificateFile,
					SSLKeyFile:         defaultPathsDatabaseInserterSSLKeyFile,
				},
				Updater: pathsDatabaseUser{
					SSLCertificateFile: defaultPathsDatabaseUpdaterSSLCertificateFile,
					SSLKeyFile:         defaultPathsDatabaseUpdaterSSLKeyFile,
				},
				Deletor: pathsDatabaseUser{
					SSLCertificateFile: defaultPathsDatabaseDeletorSSLCertificateFile,
					SSLKeyFile:         defaultPathsDatabaseDeletorSSLKeyFile,
				},
				Migrator: pathsDatabaseUser{
					SSLCertificateFile: defaultPathsDatabaseMigratorSSLCertificateFile,
					SSLKeyFile:         defaultPathsDatabaseMigratorSSLKeyFile,
				},
			},
		},
	}
}
