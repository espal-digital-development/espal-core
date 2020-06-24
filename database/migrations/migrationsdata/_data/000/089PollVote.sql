CREATE TABLE "PollVote" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"pollOptionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	UNIQUE INDEX "PollVote_uidx_createdByID_pollOptionID" ("createdByID", "pollOptionID"),
	INDEX "PollVote_idx_updatedByID" ("updatedByID"),
	INDEX "PollVote_idx_pollOptionID" ("pollOptionID"),
	CONSTRAINT "PollVote_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollVote_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollVote_pollOptionID" FOREIGN KEY ("pollOptionID") REFERENCES "PollOption" ("id"),
	FAMILY "Primary" ("id", "createdByID", "pollOptionID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt")
);
