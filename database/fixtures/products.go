package fixtures

import (
	"github.com/juju/errors"
)

const productModelQuery = `INSERT INTO "ProductModel"("createdByID","nameRepresentationID",
	"descriptionRepresentationID","imageRepresentationID","taxGroupID","active")
	VALUES($1,$2,$3,$4,$5,$6) RETURNING "id"`
const productModelDimensionQuery = `INSERT INTO "ProductModelDimension"("createdByID","productModelID","propertyID")
	VALUES($1,$2,$3)`
const productVariantQuery = `INSERT INTO "ProductVariant"("createdByID","productModelID","taxGroupID","active")
	VALUES($1,$2,$3,$4)`

func (f *Fixtures) products() error {
	// ProductModel
	row := f.inserterDatabase.QueryRow(productModelQuery, f.mainUserID, f.propertyNameID, f.propertyDescriptionID,
		f.propertyImageID, f.taxGroupID, true)
	if err := row.Scan(&f.productModelID); err != nil {
		return errors.Trace(err)
	}

	// ProductModelDimension
	if _, err := f.inserterDatabase.Exec(productModelDimensionQuery, f.mainUserID, f.productModelID,
		f.propertyColorID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(productModelDimensionQuery, f.mainUserID, f.productModelID,
		f.propertyLengthSizeID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(productModelDimensionQuery, f.mainUserID, f.productModelID,
		f.propertyLengthWidthSizeID); err != nil {
		return errors.Trace(err)
	}

	// ProductVariant
	if _, err := f.inserterDatabase.Exec(productVariantQuery, f.mainUserID, f.productModelID, f.taxGroupID,
		true); err != nil {
		return errors.Trace(err)
	}

	return nil
}
