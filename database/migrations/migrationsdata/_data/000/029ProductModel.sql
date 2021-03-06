CREATE TABLE "ProductModel" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"taxGroupID" UUID NOT NULL,
	"nameRepresentationID" UUID,
	"descriptionRepresentationID" UUID,
	"imageRepresentationID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
    "key" STRING(255),
    UNIQUE INDEX "ProductModel_uidx_key" ("key"),
	INDEX "ProductModel_idx_createdByID" ("createdByID"),
	INDEX "ProductModel_idx_updatedByID" ("updatedByID"),
	INDEX "ProductModel_idx_taxGroupID" ("taxGroupID"),
	INDEX "ProductModel_idx_nameRepresentationID" ("nameRepresentationID"),
	INDEX "ProductModel_idx_descriptionRepresentationID" ("descriptionRepresentationID"),
	INDEX "ProductModel_idx_imageRepresentationID" ("imageRepresentationID"),
	CONSTRAINT "ProductModel_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ProductModel_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ProductModel_taxGroupID" FOREIGN KEY ("taxGroupID") REFERENCES "TaxGroup" ("id"),
	CONSTRAINT "ProductModel_nameRepresentationID" FOREIGN KEY ("nameRepresentationID") REFERENCES "Property" ("id"),
	CONSTRAINT "ProductModel_descriptionRepresentationID" FOREIGN KEY ("descriptionRepresentationID") REFERENCES "Property" ("id"),
	CONSTRAINT "ProductModel_imageRepresentationID" FOREIGN KEY ("imageRepresentationID") REFERENCES "Property" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt"),
    FAMILY "Tertiary" ("taxGroupID", "active", "sorting", "key"),
	FAMILY "Properties" ("nameRepresentationID", "descriptionRepresentationID", "imageRepresentationID")
);