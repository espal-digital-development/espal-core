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
  name: app
  users:
    selecter: selecter
    creator: creator
    inserter: inserter
    updater: updater
    deletor: deletor
    migrator: migrator
email:
  host: smtp.domain.dev
  port: 2525
  username: fakeUsername
  password: fakePassword
  noReplyAddress: noreply@domain.dev
security:
  globalAuthentication: true
  bcryptRounds: 12
  formTokenLifespan: 8m
  formTokenCleanupInterval: 10s
  jwtSigningMethod: HS512
  jwtPassword: e86074797a09ccb62688c0fdf149ab18
  httpReferrerPolicy: same-origin
  httpContentSecurityPolicy: default-src 'self'; frame-ancestors 'self'
session:
  cookieName: s
  expiration: 45m
  rememberMeExpiration: 720h
urls:
  admin: _adminPath
  pprof: _pprofPath
assets:
  brotli: true
  gzip: true
  brotliFiles: true
  gzipFiles: true
  optimizePngs: true
  optimizeJpegs: true
  optimizeGifs: true
  optimizeSvgs: true
  cacheMaxAge: 60
paths:
  server:
    sslCertificateFile: ./app/localhost.crt
    sslKeyFile: ./app/localhost.key
  database:
    sslRootCertificateFile: ./app/database/ca.crt
    selecter:
      sslCertificateFile: ./app/database/client.selecter.crt
      sslKeyFile: ./app/database/client.selecter.key
    creator:
      sslCertificateFile: ./app/database/client.creator.crt
      sslKeyFile: ./app/database/client.creator.key
    inserter:
      sslCertificateFile: ./app/database/client.inserter.crt
      sslKeyFile: ./app/database/client.inserter.key
    updater:
      sslCertificateFile: ./app/database/client.updater.crt
      sslKeyFile: ./app/database/client.updater.key
    deletor:
      sslCertificateFile: ./app/database/client.deletor.crt
      sslKeyFile: ./app/database/client.deletor.key
    migrator:
      sslCertificateFile: ./app/database/client.migrator.crt
      sslKeyFile: ./app/database/client.migrator.key
  assets:
    stylesheets: ./app/assets/css
    javascript: ./app/assets/js
    images: ./app/assets/images
    publicRootFiles: ./app/assets/files/root
    publicFiles: ./app/assets/files/public
    privateFiles: ./app/assets/files/private
  translations: ./app/translations
`)
