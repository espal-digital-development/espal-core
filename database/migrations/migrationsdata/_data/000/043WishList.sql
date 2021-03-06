CREATE TABLE "WishList" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID NOT NULL,
	"userID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "sorting" INT NOT NULL DEFAULT 0,
	INDEX "WishList_idx_createdByID" ("createdByID"),
	INDEX "WishList_idx_updatedByID" ("updatedByID"),
	INDEX "WishList_idx_domainID" ("domainID"),
	INDEX "WishList_idx_userID" ("userID"),
	CONSTRAINT "WishList_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "WishList_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "WishList_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	CONSTRAINT "WishList_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt", "domainID", "userID"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "sorting")
);
