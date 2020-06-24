CREATE TABLE "PropertyOptionTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"propertyOptionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "PropertyOptionTranslation_uidx_propertyOptionID_language_field" ("propertyOptionID", "language", "field"),
	INDEX "PropertyOptionTranslation_idx_createdByID" ("createdByID"),
	INDEX "PropertyOptionTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PropertyOptionTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyOptionTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "PropertyOptionTranslation_propertyOptionID" FOREIGN KEY ("propertyOptionID") REFERENCES "PropertyOption" ("id"),
	FAMILY "Primary" ("id", "createdByID", "propertyOptionID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
