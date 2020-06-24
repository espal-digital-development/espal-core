CREATE TABLE "PollOptionTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"pollOptionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "PollOptionTranslation_uidx_pollOptionID_language_field" ("pollOptionID", "language", "field"),
	INDEX "PollOptionTranslation_idx_createdByID" ("createdByID"),
	INDEX "PollOptionTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PollOptionTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollOptionTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollOptionTranslation_pollOptionID" FOREIGN KEY ("pollOptionID") REFERENCES "PollOption" ("id"),
	FAMILY "Primary" ("id", "createdByID", "pollOptionID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
