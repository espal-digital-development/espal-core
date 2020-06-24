CREATE TABLE "UserGroup" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "active" BOOL NOT NULL DEFAULT false,
    "userRights" STRING NOT NULL DEFAULT '',
    "currencies" STRING NOT NULL DEFAULT '',
	INDEX "UserGroup_idx_createdByID" ("createdByID"),
	INDEX "UserGroup_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "UserGroup_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserGroup_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "updatedByID", "createdAt", "updatedAt",
		"active", "userRights", "currencies")
);
