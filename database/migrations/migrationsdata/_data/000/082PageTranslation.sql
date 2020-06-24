CREATE TABLE "PageTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"pageID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "PageTranslation_uidx_pageID_language_field" ("pageID", "language", "field"),
	INDEX "PageTranslation_idx_createdByID" ("createdByID"),
	INDEX "PageTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PageTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PageTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PageTranslation_pageID" FOREIGN KEY ("pageID") REFERENCES "Page" ("id"),
	FAMILY "Primary" ("id", "createdByID", "pageID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
