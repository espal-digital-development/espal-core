CREATE TABLE "UserGroupTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"userGroupID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
    "value" STRING NOT NULL,
    UNIQUE INDEX "UserGroupTranslation_uidx_userID" ("userGroupID", "language", "field"),
	INDEX "UserGroupTranslation_idx_createdByID" ("createdByID"),
	INDEX "UserGroupTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "UserGroupTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserGroupTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserGroupTranslation_userGroupID" FOREIGN KEY ("userGroupID") REFERENCES "UserGroup" ("id"),
	FAMILY "Primary" ("id", "createdByID", "userGroupID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
