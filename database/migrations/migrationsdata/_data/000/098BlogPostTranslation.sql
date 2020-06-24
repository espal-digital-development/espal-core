CREATE TABLE "BlogPostTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"blogPostID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "BlogPostTranslation_uidx_blogPostID_language_field" ("blogPostID", "language", "field"),
	INDEX "BlogPostTranslation_idx_createdByID" ("createdByID"),
	INDEX "BlogPostTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "BlogPostTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostTranslation_blogPostID" FOREIGN KEY ("blogPostID") REFERENCES "BlogPost" ("id"),
	FAMILY "Primary" ("id", "createdByID", "blogPostID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
