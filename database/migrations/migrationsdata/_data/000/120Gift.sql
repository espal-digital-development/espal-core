CREATE TABLE "Gift" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	INDEX "Gift_idx_createdByID" ("createdByID"),
	INDEX "Gift_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "Gift_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Gift_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt")
);
