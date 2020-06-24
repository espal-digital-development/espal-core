CREATE TABLE "BlogPostComment" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"blogPostID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"title" STRING(255),
	"message" STRING(255) NOT NULL,
	UNIQUE INDEX "BlogPostComment_uidx_createdByID_blogPostID" ("createdByID", "blogPostID"),
	INDEX "BlogPostComment_idx_updatedByID" ("updatedByID"),
	INDEX "BlogPostComment_idx_blogPostID" ("blogPostID"),
	CONSTRAINT "BlogPostComment_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostComment_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostComment_blogPostID" FOREIGN KEY ("blogPostID") REFERENCES "BlogPost" ("id"),
	FAMILY "Primary" ("id", "createdByID", "blogPostID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "title", "message")
);
