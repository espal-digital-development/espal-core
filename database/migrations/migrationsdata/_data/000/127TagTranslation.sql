CREATE TABLE "TagTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"tagID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "TagTranslation_uidx_tagID_language_field" ("tagID", "language", "field"),
	INDEX "TagTranslation_idx_createdByID" ("createdByID"),
	INDEX "TagTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "TagTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "TagTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "TagTranslation_tagID" FOREIGN KEY ("tagID") REFERENCES "Tag" ("id"),
	FAMILY "Primary" ("id", "createdByID", "tagID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
