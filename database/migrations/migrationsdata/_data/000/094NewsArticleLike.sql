CREATE TABLE "NewsArticleLike" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"newsArticleID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	UNIQUE INDEX "NewsArticleLike_uidx_createdByID_newsArticleID" ("createdByID", "newsArticleID"),
	INDEX "NewsArticleLike_idx_updatedByID" ("updatedByID"),
	INDEX "NewsArticleLike_idx_newsArticleID" ("newsArticleID"),
	CONSTRAINT "NewsArticleLike_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleLike_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleLike_newsArticleID" FOREIGN KEY ("newsArticleID") REFERENCES "NewsArticle" ("id"),
	FAMILY "Primary" ("id", "createdByID", "newsArticleID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt")
);
