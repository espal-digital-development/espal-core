CREATE TABLE "Page" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	INDEX "Page_idx_createdByID" ("createdByID"),
	INDEX "Page_idx_updatedByID" ("updatedByID"),
	INDEX "Page_idx_domainID" ("domainID"),
	CONSTRAINT "Page_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Page_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Page_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	FAMILY "Primary" ("id", "createdByID", "domainID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active")
);
