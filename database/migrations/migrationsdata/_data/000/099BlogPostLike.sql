CREATE TABLE "BlogPostLike" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"blogPostID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	INDEX "BlogPostLike_idx_createdByID" ("createdByID"),
	INDEX "BlogPostLike_idx_updatedByID" ("updatedByID"),
	INDEX "BlogPostLike_idx_blogPostID" ("blogPostID"),
	CONSTRAINT "BlogPostLike_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostLike_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostLike_blogPostID" FOREIGN KEY ("blogPostID") REFERENCES "BlogPost" ("id"),
	FAMILY "Primary" ("id", "createdByID", "blogPostID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt")
);
