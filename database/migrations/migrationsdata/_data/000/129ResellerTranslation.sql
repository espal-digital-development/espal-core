CREATE TABLE "ResellerTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"resellerID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "ResellerTranslation_uidx_resellerID_language_field" ("resellerID", "language", "field"),
	INDEX "ResellerTranslation_idx_createdByID" ("createdByID"),
	INDEX "ResellerTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "ResellerTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ResellerTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ResellerTranslation_resellerID" FOREIGN KEY ("resellerID") REFERENCES "Reseller" ("id"),
	FAMILY "Primary" ("id", "createdByID", "resellerID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
