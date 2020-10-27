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
func (f *Fixtures) Run() error {
	if err := f.users(); err != nil {
		return errors.Trace(err)
	}
	if err := f.usersAndUserGroups(); err != nil {
		return errors.Trace(err)
	}
	if err := f.forums(); err != nil {
		return errors.Trace(err)
	}

	// Site
	siteQuery := `INSERT INTO "Site"("createdByID","online","language","currencies") VALUES($1,$2,$3,$4) RETURNING "id"`
	var site1ID string
	row := f.inserterDatabase.QueryRow(siteQuery, f.mainUserID, true, strconv.Itoa(int(f.dutchLanguage.ID())),
		strconv.Itoa(int(f.euroCurrency.ID())))
	if err := row.Scan(&site1ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(
		`INSERT INTO "SiteTranslation"("createdByID","siteID","language","field","value") VALUES($1,$2,$3,$4,$5)`,
		f.mainUserID, site1ID, f.englishLanguage.ID(), database.DBTranslationFieldName,
		"Localhost Website on port 8443"); err != nil {
		return errors.Trace(err)
	}

	// Setting
	if _, err := f.inserterDatabase.Exec(
		`INSERT INTO "Setting"("createdByID","siteID","key","value") VALUES($1,$2,$3,$4)`, f.mainUserID, site1ID, 1,
		"2"); err != nil {
		return errors.Trace(err)
	}

	// SiteUser
	siteUserQuery := `INSERT INTO "SiteUser"("createdByID","siteID","userID") VALUES($1,$2,$3)`
	if _, err := f.inserterDatabase.Exec(siteUserQuery, f.mainUserID, site1ID, f.mainUserID); err != nil {
		return errors.Trace(err)
	}

	// Domain
	domainQuery := `INSERT INTO "Domain"("createdByID","siteID","active","language","host","currencies")
		VALUES($1,$2,$3,$4,$5,$6) RETURNING "id"`
	domains := map[string]string{
		"espal.loc":          "",
		"www.espal.loc":      "",
		"espal.loc:8443":     "",
		"www.espal.loc:8443": "",
		"localhost:8443":     "",
	}
	for domainName := range domains {
		row = f.inserterDatabase.QueryRow(domainQuery, f.mainUserID, site1ID, true, f.englishLanguage.ID(),
			domainName, "")
		var domainID string
		if err := row.Scan(&domainID); err != nil {
			return errors.Trace(err)
		}
		domains[domainName] = domainID
	}

	// Slug
	slugQuery := `INSERT INTO "Slug"("createdByID","domainID","language","path","rerouteTo") VALUES($1,$2,$3,$4,$5)`
	if _, err := f.inserterDatabase.Exec(slugQuery, f.mainUserID, domains["espal.loc"], f.englishLanguage.ID(),
		"fake-slug", "Login"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(slugQuery, f.mainUserID, domains["www.espal.loc"], f.dutchLanguage.ID(),
		"inloggen", "Login"); err != nil {
		return errors.Trace(err)
	}

	if err := f.taxes(); err != nil {
		return errors.Trace(err)
	}

	// Propertygroup
	propertyGroupTranslationQuery := `INSERT INTO "PropertyGroupTranslation"("createdByID","propertyGroupID",
		"language","field","value") VALUES($1,$2,$3,$4,$5) RETURNING "id"`
	row = f.inserterDatabase.QueryRow(`INSERT INTO "PropertyGroup"("createdByID","active") VALUES($1,$2)
		RETURNING "id"`, f.mainUserID, true)
	if err := row.Scan(&f.propertyGroupID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyGroupTranslationQuery, f.mainUserID, f.propertyGroupID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "General"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyGroupTranslationQuery, f.mainUserID, f.propertyGroupID,
		f.englishLanguage.ID(), database.DBTranslationFieldDescription,
		"The most generic collection of properties"); err != nil {
		return errors.Trace(err)
	}

	if err := f.properties(); err != nil {
		return errors.Trace(err)
	}
	if err := f.products(); err != nil {
		return errors.Trace(err)
	}

	return nil
}

// New returns a new instance of Fixtures.
func New(inserterDatabase database.Database, updaterDatabase database.Database,
	languagesRepository languages.Repository, countriesRepository countries.Repository,
	currenciesRepository currencies.Repository,
	userRightsRepository userrights.Repository) (*Fixtures, error) {
	f := &Fixtures{
		inserterDatabase:     inserterDatabase,
		updaterDatabase:      updaterDatabase,
		languagesRepository:  languagesRepository,
		countriesRepository:  countriesRepository,
		currenciesRepository: currenciesRepository,
		userRightsRepository: userRightsRepository,
	}

	var err error
	f.dutchLanguage, err = languagesRepository.ByCode("nl")
	if err != nil {
		return nil, errors.Trace(err)
	}
	f.englishLanguage, err = languagesRepository.ByCode("en")
	if err != nil {
		return nil, errors.Trace(err)
	}
	f.unitedKingdomCountry, err = countriesRepository.ByCode("GB")
	if err != nil {
		return nil, errors.Trace(err)
	}
	f.euroCurrency, err = currenciesRepository.ByCode("EUR")
	if err != nil {
		return nil, errors.Trace(err)
	}

	return f, nil
}
