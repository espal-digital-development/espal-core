package configmock

//go:generate moq -pkg configmock -out config.go .. Config

// DefaultConfigMock returns a quick-to-use basic instance of ConfigMock.
func DefaultConfigMock() *ConfigMock {
	return &ConfigMock{
		// DatabaseCreatorFunc: func() string {
		// 	return "espal_creator"
		// },
		// DatabaseDeletorFunc: func() string {
		// 	return "espal_deletor"
		// },
		// DatabaseInserterFunc: func() string {
		// 	return "espal_inserter"
		// },
		// DatabaseMigratorFunc: func() string {
		// 	return "espal_migrator"
		// },
		// DatabaseSelecterFunc: func() string {
		// 	return "espal_selecter"
		// },
		// DatabaseUpdaterFunc: func() string {
		// 	return "espal_updater"
		// },

		// DatabaseHostFunc: func() string {
		// 	return "localhost"
		// },
		// DatabasePortFunc: func() int {
		// 	return 26257
		// },
		// DatabaseNameFunc: func() string {
		// 	return "espal"
		// },

		// DatabaseSSLRootCertificateFileFunc: func() string {
		// 	return "ca.crt"
		// },

		// DatabaseCreatorSSLCertificateFileFunc: func() string {
		// 	return "client.espal_creator.crt"
		// },
		// DatabaseCreatorSSLKeyFileFunc: func() string {
		// 	return "client.espal_creator.key"
		// },
		// DatabaseDeletorSSLCertificateFileFunc: func() string {
		// 	return "client.espal_deletor.crt"
		// },
		// DatabaseDeletorSSLKeyFileFunc: func() string {
		// 	return "client.espal_deletor.key"
		// },
		// DatabaseInserterSSLCertificateFileFunc: func() string {
		// 	return "client.espal_inserter.crt"
		// },
		// DatabaseInserterSSLKeyFileFunc: func() string {
		// 	return "client.espal_inserter.key"
		// },
		// DatabaseMigratorSSLCertificateFileFunc: func() string {
		// 	return "client.espal_migrator.crt"
		// },
		// DatabaseMigratorSSLKeyFileFunc: func() string {
		// 	return "client.espal_migrator.key"
		// },
		// DatabaseSelecterSSLCertificateFileFunc: func() string {
		// 	return "client.espal_selecter.crt"
		// },
		// DatabaseSelecterSSLKeyFileFunc: func() string {
		// 	return "client.espal_selecter.key"
		// },
		// DatabaseUpdaterSSLCertificateFileFunc: func() string {
		// 	return "client.espal_updater.crt"
		// },
		// DatabaseUpdaterSSLKeyFileFunc: func() string {
		// 	return "client.espal_updater.key"
		// },
	}
}
