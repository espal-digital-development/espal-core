CREATE TABLE "BlockTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"blockID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "BlockTranslation_uidx_blockID_language_field" ("blockID", "language", "field"),
	INDEX "BlockTranslation_idx_createdByID" ("createdByID"),
	INDEX "BlockTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "BlockTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlockTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "BlockTranslation_blockID" FOREIGN KEY ("blockID") REFERENCES "Block" ("id"),
	FAMILY "Primary" ("id", "createdByID", "blockID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
