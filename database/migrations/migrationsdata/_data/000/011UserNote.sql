CREATE TABLE "UserNote" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"userID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "title" STRING(125),
    "contents" STRING,
	INDEX "UserNote_idx_createdByID" ("createdByID"),
	INDEX "UserNote_idx_updatedByID" ("updatedByID"),
	INDEX "UserNote_idx_userID" ("userID"),
	CONSTRAINT "UserNote_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserNote_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "UserNote_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "updatedByID", "userID", "createdAt",
        "updatedAt", "title", "contents")
);
