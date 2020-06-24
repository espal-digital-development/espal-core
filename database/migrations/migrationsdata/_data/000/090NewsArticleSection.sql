CREATE TABLE "NewsArticleSection" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"parentID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
	INDEX "NewsArticleSection_idx_createdByID" ("createdByID"),
	INDEX "NewsArticleSection_idx_updatedByID" ("updatedByID"),
	INDEX "NewsArticleSection_idx_parentID" ("parentID"),
	CONSTRAINT "NewsArticleSection_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleSection_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleSection_parentID" FOREIGN KEY ("parentID") REFERENCES "NewsArticleSection" ("id"),
	FAMILY "Primary" ("id", "createdByID", "parentID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting")
);
