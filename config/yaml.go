package config

func (c *Configuration) getDefaultYaml() configYaml {
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
		},
		Session: session{
			CookieName:           defaultSessionCookieName,
			Expiration:           defaultSessionExpiration,
			RememberMeExpiration: defaultSessionRememberMeExpiration,
		},
		Assets: assets{
			Gzip:        defaultAssetsGzip,
			Brotli:      defaultAssetsBrotli,
			GzipFiles:   defaultAssetsGzipFiles,
			BrotliFiles: defaultAssetsBrotliFiles,
			CacheMaxAge: defaultAssetsCacheMaxAge,
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
