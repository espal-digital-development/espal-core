CREATE TABLE "Poll" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"startDate" TIMESTAMP,
	"endDate" TIMESTAMP,
	"allowAnonymousVoting" BOOL NOT NULL,
	"comments" STRING,
	INDEX "Poll_idx_createdByID" ("createdByID"),
	INDEX "Poll_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "Poll_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Poll_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "startDate", "endDate",
		"allowAnonymousVoting", "comments")
);
