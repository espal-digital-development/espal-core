CREATE TABLE "PriceGroupTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"priceGroupID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "PriceGroupTranslation_uidx_priceGroupID_language_field" ("priceGroupID", "language", "field"),
	INDEX "PriceGroupTranslation_idx_createdByID" ("createdByID"),
	INDEX "PriceGroupTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PriceGroupTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PriceGroupTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "PriceGroupTranslation_priceGroupID" FOREIGN KEY ("priceGroupID") REFERENCES "PriceGroup" ("id"),
	FAMILY "Primary" ("id", "createdByID", "priceGroupID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
