package config

import "time"

const (
	defaultHTTPPort                                = 80
	defaultServerHost                              = "localhost"
	defaultServerPort                              = 8443
	defaultServerHTTPRedirectPort                  = 8080
	defaultDatabaseHost                            = "localhost"
	defaultDatabasePort                            = 36257
	defaultDatabaseName                            = "app"
	defaultDatabaseUsersSelecter                   = "selecter"
	defaultDatabaseUsersCreator                    = "creator"
	defaultDatabaseUsersInserter                   = "inserter"
	defaultDatabaseUsersUpdater                    = "updater"
	defaultDatabaseUsersDeletor                    = "deletor"
	defaultDatabaseUsersMigrator                   = "migrator"
	defaultSecurityBcryptRounds                    = 12
	defaultSecurityFormTokenLifespan               = time.Minute * 8
	defaultSecurityFormTokenCleanupInterval        = time.Second * 10
	defaultSessionCookieName                       = "s"
	defaultSessionExpiration                       = time.Minute * 45
	defaultSessionRememberMeExpiration             = time.Hour * 720
	defaultAssetsGzip                              = true
	defaultAssetsGzipFiles                         = true
	defaultAssetsCacheMaxAge                       = "60"
	defaultPathsServerSSLCertificateFile           = "./app/server/localhost.crt"
	defaultPathsServerSSLKeyFile                   = "./app/server/localhost.key"
	defaultPathsDatabaseSSLRootCertificateFile     = "./app/database/certs/ca.crt"
	defaultPathsDatabaseSelecterSSLCertificateFile = "./app/database/certs/client.selecter.crt"
	defaultPathsDatabaseSelecterSSLKeyFile         = "./app/database/certs/client.selecter.key"
	defaultPathsDatabaseCreatorSSLCertificateFile  = "./app/database/certs/client.creator.crt"
	defaultPathsDatabaseCreatorSSLKeyFile          = "./app/database/certs/client.creator.key"
	defaultPathsDatabaseInserterSSLCertificateFile = "./app/database/certs/client.inserter.crt"
	defaultPathsDatabaseInserterSSLKeyFile         = "./app/database/certs/client.inserter.key"
	defaultPathsDatabaseUpdaterSSLCertificateFile  = "./app/database/certs/client.updater.crt"
	defaultPathsDatabaseUpdaterSSLKeyFile          = "./app/database/certs/client.updater.key"
	defaultPathsDatabaseDeletorSSLCertificateFile  = "./app/database/certs/client.deletor.crt"
	defaultPathsDatabaseDeletorSSLKeyFile          = "./app/database/certs/client.deletor.key"
	defaultPathsDatabaseMigratorSSLCertificateFile = "./app/database/certs/client.migrator.crt"
	defaultPathsDatabaseMigratorSSLKeyFile         = "./app/database/certs/client.migrator.key"
)
