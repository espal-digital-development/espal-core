CREATE TABLE "PollOption" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"pollID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
	INDEX "PollOption_idx_createdByID" ("createdByID"),
	INDEX "PollOption_idx_updatedByID" ("updatedByID"),
	INDEX "PollOption_idx_pollID" ("pollID"),
	CONSTRAINT "PollOption_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollOption_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PollOption_pollID" FOREIGN KEY ("pollID") REFERENCES "Poll" ("id"),
	FAMILY "Primary" ("id", "createdByID", "pollID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting")
);
