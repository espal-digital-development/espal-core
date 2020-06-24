CREATE TABLE "StoreTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"storeID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "StoreTranslation_uidx_storeID_language_field" ("storeID", "language", "field"),
	INDEX "StoreTranslation_idx_createdByID" ("createdByID"),
	INDEX "StoreTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "StoreTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "StoreTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "StoreTranslation_storeID" FOREIGN KEY ("storeID") REFERENCES "Store" ("id"),
	FAMILY "Primary" ("id", "createdByID", "storeID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
