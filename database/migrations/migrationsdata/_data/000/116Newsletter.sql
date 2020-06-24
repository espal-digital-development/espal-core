CREATE TABLE "Newsletter" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sendAt" TIMESTAMP,
	INDEX "Newsletter_idx_createdByID" ("createdByID"),
	INDEX "Newsletter_idx_updatedByID" ("updatedByID"),
	INDEX "Newsletter_idx_domainID" ("domainID"),
	CONSTRAINT "Newsletter_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Newsletter_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Newsletter_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	FAMILY "Primary" ("id", "createdByID", "domainID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sendAt")
);
