CREATE TABLE "BlogPostSectionTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"blogPostSectionID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "BlogPostSectionTranslation_uidx_blogPostSectionID_language_field" ("blogPostSectionID", "language", "field"),
	INDEX "BlogPostSectionTranslation_idx_createdByID" ("createdByID"),
	INDEX "BlogPostSectionTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "BlogPostSectionTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostSectionTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlogPostSectionTranslation_blogPostSectionID" FOREIGN KEY ("blogPostSectionID") REFERENCES "BlogPostSection" ("id"),
	FAMILY "Primary" ("id", "createdByID", "blogPostSectionID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
