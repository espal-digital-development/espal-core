package fixtures

import (
	"github.com/juju/errors"
)

const productModelQuery = `INSERT INTO "ProductModel"("createdByID","nameRepresentationID","descriptionRepresentationID","imageRepresentationID","taxGroupID","active") VALUES($1,$2,$3,$4,$5,$6) RETURNING "id"`
const productModelDimensionQuery = `INSERT INTO "ProductModelDimension"("createdByID","productModelID","propertyID") VALUES($1,$2,$3)`
const productVariantQuery = `INSERT INTO "ProductVariant"("createdByID","productModelID","taxGroupID","active") VALUES($1,$2,$3,$4)`

func (fixtures *Fixtures) products() error {
	// ProductModel
	row := fixtures.inserterDatabase.QueryRow(productModelQuery, fixtures.mainUserID, fixtures.propertyNameID, fixtures.propertyDescriptionID, fixtures.propertyImageID, fixtures.taxGroupID, true)
	if err := row.Scan(&fixtures.productModelID); err != nil {
		return errors.Trace(err)
	}

	// ProductModelDimension
	if _, err := fixtures.inserterDatabase.Exec(productModelDimensionQuery, fixtures.mainUserID, fixtures.productModelID, fixtures.propertyColorID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(productModelDimensionQuery, fixtures.mainUserID, fixtures.productModelID, fixtures.propertyLengthSizeID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.inserterDatabase.Exec(productModelDimensionQuery, fixtures.mainUserID, fixtures.productModelID, fixtures.propertyLengthWidthSizeID); err != nil {
		return errors.Trace(err)
	}

	// ProductVariant
	if _, err := fixtures.inserterDatabase.Exec(productVariantQuery, fixtures.mainUserID, fixtures.productModelID, fixtures.taxGroupID, true); err != nil {
		return errors.Trace(err)
	}

	return nil
}
