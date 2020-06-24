CREATE TABLE "FrequentlyAskedQuestionSectionTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"frequentlyAskedQuestionSectionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "FrequentlyAskedQuestionSectionTranslation_uidx_frequentlyAskedQuestionSectionID_language_field" ("frequentlyAskedQuestionSectionID", "language", "field"),
	INDEX "FrequentlyAskedQuestionSectionTranslation_idx_createdByID" ("createdByID"),
	INDEX "FrequentlyAskedQuestionSectionTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "FrequentlyAskedQuestionSectionTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "FrequentlyAskedQuestionSectionTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "FrequentlyAskedQuestionSectionTranslation_frequentlyAskedQuestionSectionID" FOREIGN KEY ("frequentlyAskedQuestionSectionID") REFERENCES "FrequentlyAskedQuestionSection" ("id"),
	FAMILY "Primary" ("id", "createdByID", "frequentlyAskedQuestionSectionID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
