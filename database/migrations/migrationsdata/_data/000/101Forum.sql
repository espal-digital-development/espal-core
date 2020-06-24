CREATE TABLE "Forum" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"parentID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
	INDEX "Forum_idx_createdByID" ("createdByID"),
	INDEX "Forum_idx_updatedByID" ("updatedByID"),
	INDEX "Forum_idx_parentID" ("parentID"),
	CONSTRAINT "Forum_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Forum_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Forum_parentID" FOREIGN KEY ("parentID") REFERENCES "Forum" ("id"),
	FAMILY "Primary" ("id", "createdByID", "parentID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting")
);
