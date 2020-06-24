CREATE TABLE "NewsArticle" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"sectionID" UUID,
	"approvedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
	"approvedDate" TIMESTAMP,
	"publishDate" TIMESTAMP,
	"expirationDate" TIMESTAMP,
	"comments" STRING,
	INDEX "NewsArticle_idx_createdByID" ("createdByID"),
	INDEX "NewsArticle_idx_updatedByID" ("updatedByID"),
	INDEX "NewsArticle_idx_sectionID" ("sectionID"),
	INDEX "NewsArticle_idx_approvedByID" ("approvedByID"),
	CONSTRAINT "NewsArticle_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticle_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticle_sectionID" FOREIGN KEY ("sectionID") REFERENCES "NewsArticleSection" ("id"),
	CONSTRAINT "NewsArticle_approvedByID" FOREIGN KEY ("approvedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "sectionID", "approvedByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting", "approvedDate", "publishDate",
		"expirationDate", "comments")
);
