CREATE TABLE "PropertyUnitTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"propertyUnitID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	value STRING NOT NULL,
	UNIQUE INDEX "PropertyUnitTranslation_uidx_propertyUnitID_language_field" ("propertyUnitID", "language", "field"),
	INDEX "PropertyUnitTranslation_idx_createdByID" ("createdByID"),
	INDEX "PropertyUnitTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PropertyUnitTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyUnitTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "PropertyUnitTranslation_propertyUnitID" FOREIGN KEY ("propertyUnitID") REFERENCES "PropertyUnit" ("id"),
	FAMILY "Primary" ("id", "createdByID", "propertyUnitID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
