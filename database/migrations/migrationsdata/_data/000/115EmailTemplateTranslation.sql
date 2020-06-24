CREATE TABLE "EmailTemplateTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"emailTemplateID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "EmailTemplateTranslation_uidx_emailTemplateID_language_field" ("emailTemplateID", "language", "field"),
	INDEX "EmailTemplateTranslation_idx_createdByID" ("createdByID"),
	INDEX "EmailTemplateTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "EmailTemplateTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "EmailTemplateTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "EmailTemplateTranslation_emailTemplateID" FOREIGN KEY ("emailTemplateID") REFERENCES "EmailTemplate" ("id"),
	FAMILY "Primary" ("id", "createdByID", "emailTemplateID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
