package fixtures

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const propertyQuery = `INSERT INTO "Property"("createdByID","active","_type") VALUES($1,$2,$3) RETURNING "id"`
const propertyMultilingualQuery = `INSERT INTO "Property"("createdByID","active","_type","multiLingual")
	VALUES($1,$2,$3,$4) RETURNING "id"`
const propertyTranslationQuery = `INSERT INTO "PropertyTranslation"("createdByID","propertyID","language","field",
	"value") VALUES($1,$2,$3,$4,$5)`
const propertyOptionQuery = `INSERT INTO "PropertyOption"("createdByID","propertyID","active") VALUES($1,$2,$3)
	RETURNING "id"`
const propertyOptionTranslationQuery = `INSERT INTO "PropertyOptionTranslation"("createdByID","propertyOptionID",
	"language","field","value") VALUES($1,$2,$3,$4,$5)`

func (f *Fixtures) properties() error {
	if err := f.propertyName(); err != nil {
		return errors.Trace(err)
	}
	if err := f.propertyDescription(); err != nil {
		return errors.Trace(err)
	}
	if err := f.propertyImage(); err != nil {
		return errors.Trace(err)
	}
	if err := f.propertyColor(); err != nil {
		return errors.Trace(err)
	}
	if err := f.propertyLengthSize(); err != nil {
		return errors.Trace(err)
	}
	if err := f.propertyWidthSize(); err != nil {
		return errors.Trace(err)
	}
	if err := f.propertyPrice(); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyName() error {
	row := f.inserterDatabase.QueryRow(propertyMultilingualQuery, f.mainUserID, true, database.PropertytypeText, true)
	if err := row.Scan(&f.propertyNameID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyNameID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Name"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyNameID,
		f.englishLanguage.ID(), database.DBTranslationFieldDescription,
		"Generic naming definition of a product"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyDescription() error {
	row := f.inserterDatabase.QueryRow(propertyMultilingualQuery, f.mainUserID, true, database.PropertytypeText, true)
	if err := row.Scan(&f.propertyDescriptionID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyDescriptionID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Description"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyImage() error {
	row := f.inserterDatabase.QueryRow(propertyQuery, f.mainUserID, true, database.PropertytypeSinglefile)
	if err := row.Scan(&f.propertyImageID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyImageID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Image"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyColor() error {
	row := f.inserterDatabase.QueryRow(propertyQuery, f.mainUserID, true, database.PropertytypeSingleselect)
	if err := row.Scan(&f.propertyColorID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyColorID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Color"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionRedID string
	row = f.inserterDatabase.QueryRow(propertyOptionQuery, f.mainUserID, f.propertyColorID, true)
	if err := row.Scan(&propertyOptionRedID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyOptionTranslationQuery, f.mainUserID, propertyOptionRedID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Red"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionBlueID string
	row = f.inserterDatabase.QueryRow(propertyOptionQuery, f.mainUserID, f.propertyColorID, true)
	if err := row.Scan(&propertyOptionBlueID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyOptionTranslationQuery, f.mainUserID, propertyOptionBlueID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Blue"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyLengthSize() error {
	row := f.inserterDatabase.QueryRow(propertyQuery, f.mainUserID, true, database.PropertytypeSingleselect)
	if err := row.Scan(&f.propertyLengthSizeID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyLengthSizeID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Length Size"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionSize34ID string
	row = f.inserterDatabase.QueryRow(propertyOptionQuery, f.mainUserID, f.propertyLengthSizeID, true)
	if err := row.Scan(&propertyOptionSize34ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyOptionTranslationQuery, f.mainUserID, propertyOptionSize34ID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "34"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionSize36ID string
	row = f.inserterDatabase.QueryRow(propertyOptionQuery, f.mainUserID, f.propertyLengthSizeID, true)
	if err := row.Scan(&propertyOptionSize36ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyOptionTranslationQuery, f.mainUserID, propertyOptionSize36ID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "36"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyWidthSize() error {
	row := f.inserterDatabase.QueryRow(propertyQuery, f.mainUserID, true, database.PropertytypeSingleselect)
	if err := row.Scan(&f.propertyLengthWidthSizeID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyLengthWidthSizeID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Width Size"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionWidthSize36ID string
	row = f.inserterDatabase.QueryRow(propertyOptionQuery, f.mainUserID, f.propertyLengthWidthSizeID, true)
	if err := row.Scan(&propertyOptionWidthSize36ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyOptionTranslationQuery, f.mainUserID, propertyOptionWidthSize36ID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "36"); err != nil {
		return errors.Trace(err)
	}

	var propertyOptionWidthSize38ID string
	row = f.inserterDatabase.QueryRow(propertyOptionQuery, f.mainUserID, f.propertyLengthWidthSizeID, true)
	if err := row.Scan(&propertyOptionWidthSize38ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyOptionTranslationQuery, f.mainUserID, propertyOptionWidthSize38ID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "38"); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (f *Fixtures) propertyPrice() error {
	row := f.inserterDatabase.QueryRow(propertyQuery, f.mainUserID, true, database.PropertytypeCurrency)
	if err := row.Scan(&f.propertyPriceID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(propertyTranslationQuery, f.mainUserID, f.propertyPriceID,
		f.englishLanguage.ID(), database.DBTranslationFieldName, "Price"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
