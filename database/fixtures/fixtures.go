package fixtures

import (
	"bytes"
	"strconv"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/juju/errors"
)

var _ Set = &Fixtures{}

// Set represents a collection of database fixtures.
type Set interface {
	Run() error
}

// Fixtures contains all database fixtures logic.
type Fixtures struct {
	inserterDatabase     database.Database
	updaterDatabase      database.Database
	languagesRepository  languages.Repository
	countriesRepository  countries.Repository
	currenciesRepository currencies.Repository
	userRightsRepository userrights.Repository

	userRightsBuffer bytes.Buffer

	englishLanguage      languages.Data
	dutchLanguage        languages.Data
	unitedKingdomCountry countries.Data
	euroCurrency         currencies.Data

	// UUID Strings
	mainUserID                string
	propertyGroupID           string
	propertyColorID           string
	propertyLengthSizeID      string
	propertyLengthWidthSizeID string
	propertyNameID            string
	propertyDescriptionID     string
	propertyImageID           string
	propertyPriceID           string
	taxGroupID                string
	productModelID            string
}

// Run runs all basic fixtures for the system.
func (fixtures *Fixtures) Run() error {
	if err := fixtures.users(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.usersAndUserGroups(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.forums(); err != nil {
		return errors.Trace(err)
	}

	// Site
	siteQuery := `INSERT INTO "Site"("createdByID","online","language","currencies") VALUES($1,$2,$3,$4) RETURNING "id"`
	var site1ID string
	row := fixtures.inserterDatabase.QueryRow(siteQuery, fixtures.mainUserID, true, strconv.Itoa(int(fixtures.dutchLanguage.ID())), strconv.Itoa(int(fixtures.euroCurrency.ID())))
	if err := row.Scan(&site1ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(`INSERT INTO "SiteTranslation"("createdByID","siteID","language","field","value") VALUES($1,$2,$3,$4,$5)`, fixtures.mainUserID, site1ID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "Localhost Website on port 8443"); err != nil {
		return errors.Trace(err)
	}

	// Setting
	if _, err := fixtures.inserterDatabase.Exec(`INSERT INTO "Setting"("createdByID","siteID","key","value") VALUES($1,$2,$3,$4)`, fixtures.mainUserID, site1ID, 1, "2"); err != nil {
		return errors.Trace(err)
	}

	// SiteUser
	siteUserQuery := `INSERT INTO "SiteUser"("createdByID","siteID","userID") VALUES($1,$2,$3)`
	if _, err := fixtures.inserterDatabase.Exec(siteUserQuery, fixtures.mainUserID, site1ID, fixtures.mainUserID); err != nil {
		return errors.Trace(err)
	}

	// Domain
	domainQuery := `INSERT INTO "Domain"("createdByID","siteID","active","language","host","currencies") VALUES($1,$2,$3,$4,$5,$6) RETURNING "id"`
	var domain1ID string
	row = fixtures.inserterDatabase.QueryRow(domainQuery, fixtures.mainUserID, site1ID, true, fixtures.englishLanguage.ID(), "localhost:8443", "")
	if err := row.Scan(&domain1ID); err != nil {
		return errors.Trace(err)
	}
	var domain2ID string
	row = fixtures.inserterDatabase.QueryRow(domainQuery, fixtures.mainUserID, site1ID, true, fixtures.englishLanguage.ID(), "domain.test:8443", "")
	if err := row.Scan(&domain2ID); err != nil {
		return errors.Trace(err)
	}

	// Slug
	slugQuery := `INSERT INTO "Slug"("createdByID","domainID","language","path","rerouteTo") VALUES($1,$2,$3,$4,$5)`
	if _, err := fixtures.inserterDatabase.Exec(slugQuery, fixtures.mainUserID, domain1ID, fixtures.englishLanguage.ID(), "fake-slug", "Login"); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(slugQuery, fixtures.mainUserID, domain2ID, fixtures.dutchLanguage.ID(), "inloggen", "Login"); err != nil {
		return errors.Trace(err)
	}

	if err := fixtures.taxes(); err != nil {
		return errors.Trace(err)
	}

	// Propertygroup
	propertyGroupTranslationQuery := `INSERT INTO "PropertyGroupTranslation"("createdByID","propertyGroupID","language","field","value") VALUES($1,$2,$3,$4,$5) RETURNING "id"`
	row = fixtures.inserterDatabase.QueryRow(`INSERT INTO "PropertyGroup"("createdByID","active") VALUES($1,$2) RETURNING "id"`, fixtures.mainUserID, true)
	if err := row.Scan(&fixtures.propertyGroupID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyGroupTranslationQuery, fixtures.mainUserID, fixtures.propertyGroupID, fixtures.englishLanguage.ID(), database.DBTranslationFieldName, "General"); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(propertyGroupTranslationQuery, fixtures.mainUserID, fixtures.propertyGroupID, fixtures.englishLanguage.ID(), database.DBTranslationFieldDescription, "The most generic collection of properties"); err != nil {
		return errors.Trace(err)
	}

	if err := fixtures.properties(); err != nil {
		return errors.Trace(err)
	}
	if err := fixtures.products(); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// New returns a new instance of Fixtures.
func New(inserterDatabase database.Database, updaterDatabase database.Database, languagesRepository languages.Repository, countriesRepository countries.Repository, currenciesRepository currencies.Repository, userRightsRepository userrights.Repository) (*Fixtures, error) {
	fixtures := &Fixtures{
		inserterDatabase:     inserterDatabase,
		updaterDatabase:      updaterDatabase,
		languagesRepository:  languagesRepository,
		countriesRepository:  countriesRepository,
		currenciesRepository: currenciesRepository,
		userRightsRepository: userRightsRepository,
	}

	var err error
	fixtures.dutchLanguage, err = languagesRepository.ByCode("nl")
	if err != nil {
		return nil, errors.Trace(err)
	}
	fixtures.englishLanguage, err = languagesRepository.ByCode("en")
	if err != nil {
		return nil, errors.Trace(err)
	}
	fixtures.unitedKingdomCountry, err = countriesRepository.ByCode("GB")
	if err != nil {
		return nil, errors.Trace(err)
	}
	fixtures.euroCurrency, err = currenciesRepository.ByCode("EUR")
	if err != nil {
		return nil, errors.Trace(err)
	}

	return fixtures, nil
}
