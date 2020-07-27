package configmock

//go:generate moq -pkg configmock -out config.go .. Config

// DefaultConfigMock returns a quick-to-use basic instance of ConfigMock.
func DefaultConfigMock() *ConfigMock {
	return &ConfigMock{
		// DatabaseCreatorFunc: func() string {
		// 	return "creator"
		// },
		// DatabaseDeletorFunc: func() string {
		// 	return "deletor"
		// },
		// DatabaseInserterFunc: func() string {
		// 	return "inserter"
		// },
		// DatabaseMigratorFunc: func() string {
		// 	return "migrator"
		// },
		// DatabaseSelecterFunc: func() string {
		// 	return "selecter"
		// },
		// DatabaseUpdaterFunc: func() string {
		// 	return "updater"
		// },

		// DatabaseHostFunc: func() string {
		// 	return "localhost"
		// },
		// DatabasePortFunc: func() int {
		// 	return 36257
		// },
		// DatabaseNameFunc: func() string {
		// 	return "app"
		// },

		// DatabaseSSLRootCertificateFileFunc: func() string {
		// 	return "ca.crt"
		// },

		// DatabaseCreatorSSLCertificateFileFunc: func() string {
		// 	return "client.creator.crt"
		// },
		// DatabaseCreatorSSLKeyFileFunc: func() string {
		// 	return "client.creator.key"
		// },
		// DatabaseDeletorSSLCertificateFileFunc: func() string {
		// 	return "client.deletor.crt"
		// },
		// DatabaseDeletorSSLKeyFileFunc: func() string {
		// 	return "client.deletor.key"
		// },
		// DatabaseInserterSSLCertificateFileFunc: func() string {
		// 	return "client.inserter.crt"
		// },
		// DatabaseInserterSSLKeyFileFunc: func() string {
		// 	return "client.inserter.key"
		// },
		// DatabaseMigratorSSLCertificateFileFunc: func() string {
		// 	return "client.migrator.crt"
		// },
		// DatabaseMigratorSSLKeyFileFunc: func() string {
		// 	return "client.migrator.key"
		// },
		// DatabaseSelecterSSLCertificateFileFunc: func() string {
		// 	return "client.selecter.crt"
		// },
		// DatabaseSelecterSSLKeyFileFunc: func() string {
		// 	return "client.selecter.key"
		// },
		// DatabaseUpdaterSSLCertificateFileFunc: func() string {
		// 	return "client.updater.crt"
		// },
		// DatabaseUpdaterSSLKeyFileFunc: func() string {
		// 	return "client.updater.key"
		// },
	}
}
