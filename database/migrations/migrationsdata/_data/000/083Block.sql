CREATE TABLE "Block" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	INDEX "Block_idx_createdByID" ("createdByID"),
	INDEX "Block_idx_updatedByID" ("updatedByID"),
	INDEX "Block_idx_domainID" ("domainID"),
	CONSTRAINT "Block_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Block_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Block_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	FAMILY "Primary" ("id", "createdByID", "domainID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active")
);
