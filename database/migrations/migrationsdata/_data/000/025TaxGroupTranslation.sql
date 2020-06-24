CREATE TABLE "TaxGroupTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"taxGroupID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "TaxGroupTranslation_uidx_taxGroupID_language_field" ("taxGroupID", "language", "field"),
	INDEX "TaxGroupTranslation_idx_createdByID" ("createdByID"),
	INDEX "TaxGroupTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "TaxGroupTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "TaxGroupTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "TaxGroupTranslation_taxGroupID" FOREIGN KEY ("taxGroupID") REFERENCES "TaxGroup" ("id"),
	FAMILY "Primary" ("id", "createdByID", "taxGroupID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
