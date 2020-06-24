CREATE TABLE "ForumPostRating" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"forumPostID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"score" DECIMAL(9,6),
	INDEX "ForumPostRating_idx_createdByID" ("createdByID"),
	INDEX "ForumPostRating_idx_updatedByID" ("updatedByID"),
	INDEX "ForumPostRating_idx_forumPostID" ("forumPostID"),
	CONSTRAINT "ForumPostRating_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ForumPostRating_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ForumPostRating_forumPostID" FOREIGN KEY ("forumPostID") REFERENCES "ForumPost" ("id"),
	FAMILY "Primary" ("id", "createdByID", "forumPostID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "score")
);
