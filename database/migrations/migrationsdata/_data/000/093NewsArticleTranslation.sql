CREATE TABLE "NewsArticleTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"newsArticleID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "NewsArticleTranslation_uidx_newsArticleID_language_field" ("newsArticleID", "language", "field"),
	INDEX "NewsArticleTranslation_idx_createdByID" ("createdByID"),
	INDEX "NewsArticleTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "NewsArticleTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleTranslation_newsArticleID" FOREIGN KEY ("newsArticleID") REFERENCES "NewsArticle" ("id"),
	FAMILY "Primary" ("id", "createdByID", "newsArticleID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
