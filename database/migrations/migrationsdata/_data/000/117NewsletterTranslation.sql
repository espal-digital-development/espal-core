CREATE TABLE "NewsletterTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"newsletterID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "NewsletterTranslation_uidx_newsletterID_language_field" ("newsletterID", "language", "field"),
	INDEX "NewsletterTranslation_idx_createdByID" ("createdByID"),
	INDEX "NewsletterTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "NewsletterTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsletterTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "NewsletterTranslation_newsletterID" FOREIGN KEY ("newsletterID") REFERENCES "Newsletter" ("id"),
	FAMILY "Primary" ("id", "createdByID", "newsletterID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
