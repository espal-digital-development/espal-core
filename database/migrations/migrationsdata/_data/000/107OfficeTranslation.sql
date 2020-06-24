CREATE TABLE "OfficeTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"officeID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "OfficeTranslation_uidx_officeID_language_field" ("officeID", "language", "field"),
	INDEX "OfficeTranslation_idx_createdByID" ("createdByID"),
	INDEX "OfficeTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "OfficeTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "OfficeTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "OfficeTranslation_officeID" FOREIGN KEY ("officeID") REFERENCES "Office" ("id"),
	FAMILY "Primary" ("id", "createdByID", "officeID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
