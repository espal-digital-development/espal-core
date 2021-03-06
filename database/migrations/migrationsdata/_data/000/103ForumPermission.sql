CREATE TABLE "ForumPermission" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"forumID" UUID,
	"userID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"permission" INT2,
	UNIQUE INDEX "ForumPermission_uidx_forumID_userID" ("forumID", "userID"),
	INDEX "ForumPermission_idx_createdByID" ("createdByID"),
	INDEX "ForumPermission_idx_updatedByID" ("updatedByID"),
	INDEX "ForumPermission_idx_userID" ("userID"),
	CONSTRAINT "ForumPermission_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ForumPermission_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ForumPermission_forumID" FOREIGN KEY ("forumID") REFERENCES "Forum" ("id"),
	CONSTRAINT "ForumPermission_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "forumID", "userID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "permission")
);
