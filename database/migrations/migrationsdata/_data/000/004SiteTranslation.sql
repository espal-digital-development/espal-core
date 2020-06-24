CREATE TABLE "SiteTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"siteID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "SiteTranslation_uidx_siteID_language_field" ("siteID", "language", "field"),
	INDEX "SiteTranslation_idx_createdByID" ("createdByID"),
	INDEX "SiteTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "SiteTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "SiteTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "SiteTranslation_siteID" FOREIGN KEY ("siteID") REFERENCES "Site" ("id"),
	FAMILY "Primary" ("id", "createdByID", "siteID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
