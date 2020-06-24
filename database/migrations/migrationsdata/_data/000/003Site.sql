CREATE TABLE "Site" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"online" BOOL NOT NULL DEFAULT false,
    "country" INT2,
    "language" INT2,
	"currencies" STRING NOT NULL DEFAULT '',
	INDEX "Site_idx_createdByID" ("createdByID"),
	INDEX "Site_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "Site_created_by_id" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Site_updated_by_id" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt"),
	FAMILY "Tertiary" ("online", "country", "language", "currencies")
);
