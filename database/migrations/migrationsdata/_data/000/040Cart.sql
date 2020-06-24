CREATE TABLE "Cart" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
    "domainID" UUID NOT NULL,
    "userID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	INDEX "Cart_idx_createdByID" ("createdByID"),
	INDEX "Cart_idx_updatedByID" ("updatedByID"),
	INDEX "Cart_idx_domainID" ("domainID"),
	INDEX "Cart_idx_userID" ("userID"),
	CONSTRAINT "Cart_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Cart_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Cart_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	CONSTRAINT "Cart_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt", "domainID", "userID"),
    FAMILY "Seconday" ("updatedByID", "updatedAt")
);
