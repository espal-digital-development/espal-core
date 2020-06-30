package config_test

var configYmlBlueprint = []byte(`
general:
  development: true
  logging: true
  pprof: true
  languages: [en]
  defaultLanguage: en
server:
  host: localhost
  port: 8443
  httpRedirectPort: 8080
database:
  host: localhost
  port: 36257
  name: espal
  users:
    selecter: espal_selecter
    creator: espal_creator
    inserter: espal_inserter
    updater: espal_updater
    deletor: espal_deletor
    migrator: espal_migrator
email:
  host: smtp.domain.dev
  port: 2525
  username: espal
  password: fakePassword
  noReplyAddress: noreply@domain.dev
security:
  globalAuthentication: true
  bcryptRounds: 12
  formTokenLifespan: 8m
  formTokenCleanupInterval: 10s
session:
  cookieName: s
  expiration: 45m
  rememberMeExpiration: 720h
urls:
  admin: _adminPath
  pprof: _pprofPath
assets:
  gzip: true
  brotli: true
  gzipFiles: true
  brotliFiles: true
  cacheMaxAge: 60
paths:
  server:
    sslCertificateFile: ./app/localhost.crt
    sslKeyFile: ./app/localhost.key
  database:
    sslRootCertificateFile: ./app/database/ca.crt
    selecter:
      sslCertificateFile: ./app/database/client.espal_selecter.crt
      sslKeyFile: ./app/database/client.espal_selecter.key
    creator:
      sslCertificateFile: ./app/database/client.espal_creator.crt
      sslKeyFile: ./app/database/client.espal_creator.key
    inserter:
      sslCertificateFile: ./app/database/client.espal_inserter.crt
      sslKeyFile: ./app/database/client.espal_inserter.key
    updater:
      sslCertificateFile: ./app/database/client.espal_updater.crt
      sslKeyFile: ./app/database/client.espal_updater.key
    deletor:
      sslCertificateFile: ./app/database/client.espal_deletor.crt
      sslKeyFile: ./app/database/client.espal_deletor.key
    migrator:
      sslCertificateFile: ./app/database/client.espal_migrator.crt
      sslKeyFile: ./app/database/client.espal_migrator.key
  assets:
    stylesheets: ./app/assets/css
    javascript: ./app/assets/js
    images: ./app/assets/images
    publicRootFiles: ./app/assets/files/root
    publicFiles: ./app/assets/files/public
    privateFiles: ./app/assets/files/private
  translations: ./app/translations
`)
