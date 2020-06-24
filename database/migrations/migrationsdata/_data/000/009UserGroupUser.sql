CREATE TABLE "UserGroupUser" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"userGroupID" UUID NOT NULL,
	"userID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    UNIQUE INDEX "UserGroupUser_userGroupID_userID" ("userGroupID", "userID"),
	INDEX "UserGroupUser_idx_createdByID" ("createdByID"),
	INDEX "UserGroupUser_idx_updatedByID" ("updatedByID"),
	INDEX "UserGroupUser_idx_userID" ("userID"),
	CONSTRAINT "UserGroupUser_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserGroupUser_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserGroupUser_userGroupID" FOREIGN KEY ("userGroupID") REFERENCES "UserGroup" ("id"),
	CONSTRAINT "UserGroupUser_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "updatedByID", "userGroupID", "userID",
        "createdAt", "updatedAt")
);
