CREATE TABLE "StockLocationTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"stockLocationID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "StockLocationTranslation_uidx_stockLocationID_language_field" ("stockLocationID", "language", "field"),
	INDEX "StockLocationTranslation_idx_createdByID" ("createdByID"),
	INDEX "StockLocationTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "StockLocationTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "StockLocationTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "StockLocationTranslation_stockLocationID" FOREIGN KEY ("stockLocationID") REFERENCES "StockLocation" ("id"),
	FAMILY "Primary" ("id", "createdByID", "stockLocationID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
