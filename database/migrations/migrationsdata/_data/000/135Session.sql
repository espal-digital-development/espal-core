CREATE TABLE "Session" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"timeout" INT NOT NULL,
	"hash" STRING(255),
	"data" JSON NOT NULL,
	UNIQUE INDEX "Session_uidx_userID" ("hash"),
	INDEX "Session_idx_createdByID" ("createdByID"),
	INDEX "Session_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "Session_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Session_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt", "timeout", "hash"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "data")
);
