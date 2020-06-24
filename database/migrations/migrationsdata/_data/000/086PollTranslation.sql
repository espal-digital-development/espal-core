CREATE TABLE "PollTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"pollID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "PollTranslation_uidx_pollID_language_field" ("pollID", "language", "field"),
	INDEX "PollTranslation_idx_createdByID" ("createdByID"),
	INDEX "PollTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PollTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollTranslation_pollID" FOREIGN KEY ("pollID") REFERENCES "Poll" ("id"),
	FAMILY "Primary" ("id", "createdByID", "pollID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
