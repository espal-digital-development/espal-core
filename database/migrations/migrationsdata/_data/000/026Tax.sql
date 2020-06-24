CREATE TABLE "Tax" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"taxGroupID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "country" INT,
    "rate" DECIMAL(9,6),
    UNIQUE INDEX "Tax_uidx_taxGroupID_country" ("taxGroupID", "country"),
	INDEX "Tax_idx_createdByID" ("createdByID"),
	INDEX "Tax_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "Tax_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Tax_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Tax_taxGroupID" FOREIGN KEY ("taxGroupID") REFERENCES "TaxGroup" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt", "taxGroupID"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "country", "rate")
);
