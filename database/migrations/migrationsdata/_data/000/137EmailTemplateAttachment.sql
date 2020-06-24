CREATE TABLE "EmailTemplateAttachment" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"emailTemplateID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"filePath" STRING(255) NOT NULL,
    "language" INT2 NOT NULL,
	INDEX "EmailTemplateAttachment_idx_createdByID" ("createdByID"),
	INDEX "EmailTemplateAttachment_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "EmailTemplateAttachment_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "EmailTemplateAttachment_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "EmailTemplateAttachment_emailTemplateID" FOREIGN KEY ("emailTemplateID") REFERENCES "EmailTemplate" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt", "emailTemplateID"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "filePath", "language")
);
