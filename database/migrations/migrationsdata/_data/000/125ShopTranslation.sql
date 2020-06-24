CREATE TABLE "ShopTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"shopID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "ShopTranslation_uidx_shopID_language_field" ("shopID", "language", "field"),
	INDEX "ShopTranslation_idx_createdByID" ("createdByID"),
	INDEX "ShopTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "ShopTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ShopTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ShopTranslation_shopID" FOREIGN KEY ("shopID") REFERENCES "Shop" ("id"),
	FAMILY "Primary" ("id", "createdByID", "shopID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
