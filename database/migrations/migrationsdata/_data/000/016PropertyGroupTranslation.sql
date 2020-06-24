CREATE TABLE "PropertyGroupTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
    "propertyGroupID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	value STRING NOT NULL,
	UNIQUE INDEX "PropertyGroupTranslation_uidx_propertyGroupID_language_field" ("propertyGroupID", "language", "field"),
	INDEX "PropertyGroupTranslation_idx_createdByID" ("createdByID"),
	INDEX "PropertyGroupTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PropertyGroupTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyGroupTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "PropertyGroupTranslation_propertyGroupID" FOREIGN KEY ("propertyGroupID") REFERENCES "PropertyGroup" ("id"),
	FAMILY "Primary" ("id", "createdByID", "propertyGroupID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
