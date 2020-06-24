CREATE TABLE "NewsArticleSectionTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"newsArticleSectionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "NewsArticleSectionTranslation_uidx_newsArticleSectionID_language_field" ("newsArticleSectionID", "language", "field"),
	INDEX "NewsArticleSectionTranslation_idx_createdByID" ("createdByID"),
	INDEX "NewsArticleSectionTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "NewsArticleSectionTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleSectionTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsArticleSectionTranslation_newsArticleSectionID" FOREIGN KEY ("newsArticleSectionID") REFERENCES "NewsArticleSection" ("id"),
	FAMILY "Primary" ("id", "createdByID", "newsArticleSectionID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
