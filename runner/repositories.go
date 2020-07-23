package runner

import (
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
)

type repositories struct {
	regularExpressions regularexpressions.Repository
	userRights         userrights.Repository
	languages          languages.Repository
	translations       translations.Repository
	countries          countries.Repository
	currencies         currencies.Repository
	themes             themes.Repository
}
