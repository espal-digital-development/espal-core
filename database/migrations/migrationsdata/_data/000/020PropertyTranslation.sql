CREATE TABLE "PropertyTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"propertyID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "PropertyTranslation_uidx_propertyID_language_field" ("propertyID", "language", "field"),
	INDEX "PropertyTranslation_idx_createdByID" ("createdByID"),
	INDEX "PropertyTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PropertyTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "PropertyTranslation_propertyID" FOREIGN KEY ("propertyID") REFERENCES "Property" ("id"),
	FAMILY "Primary" ("id", "createdByID", "propertyID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
