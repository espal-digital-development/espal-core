CREATE TABLE "BlogPostSection" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"parentID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
	INDEX "BlogPostSection_idx_createdByID" ("createdByID"),
	INDEX "BlogPostSection_idx_updatedByID" ("updatedByID"),
	INDEX "BlogPostSection_idx_parentID" ("parentID"),
	CONSTRAINT "BlogPostSection_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostSection_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostSection_parentID" FOREIGN KEY ("parentID") REFERENCES "BlogPostSection" ("id"),
	FAMILY "Primary" ("id", "createdByID", "parentID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting")
);
