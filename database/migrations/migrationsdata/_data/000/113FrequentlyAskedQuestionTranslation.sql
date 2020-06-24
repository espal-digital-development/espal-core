CREATE TABLE "FrequentlyAskedQuestionTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"frequentlyAskedQuestionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "FrequentlyAskedQuestionTranslation_uidx_frequentlyAskedQuestionID_language_field" ("frequentlyAskedQuestionID", "language", "field"),
	INDEX "FrequentlyAskedQuestionTranslation_idx_createdByID" ("createdByID"),
	INDEX "FrequentlyAskedQuestionTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "FrequentlyAskedQuestionTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "FrequentlyAskedQuestionTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "FrequentlyAskedQuestionTranslation_frequentlyAskedQuestionID" FOREIGN KEY ("frequentlyAskedQuestionID") REFERENCES "FrequentlyAskedQuestion" ("id"),
	FAMILY "Primary" ("id", "createdByID", "frequentlyAskedQuestionID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
