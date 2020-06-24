CREATE TABLE "StockMutation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"sourceID" UUID,
	"targetID" UUID NOT NULL,
	"productVariantID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"modifier" INT NOT NULL,
	"ballanceAfterModifier" INT NOT NULL,
	"comments" STRING,
	INDEX "StockMutation_idx_createdByID" ("createdByID"),
	INDEX "StockMutation_idx_updatedByID" ("updatedByID"),
	INDEX "StockMutation_idx_sourceID" ("sourceID"),
	INDEX "StockMutation_idx_targetID" ("targetID"),
	INDEX "StockMutation_idx_productVariantID" ("productVariantID"),
	CONSTRAINT "StockMutation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "StockMutation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "StockMutation_sourceID" FOREIGN KEY ("sourceID") REFERENCES "StockLocation" ("id"),
	CONSTRAINT "StockMutation_targetID" FOREIGN KEY ("targetID") REFERENCES "StockLocation" ("id"),
	CONSTRAINT "StockMutation_productVariantID" FOREIGN KEY ("productVariantID") REFERENCES "ProductVariant" ("id"),
	FAMILY "Primary" ("id", "createdByID", "sourceID", "targetID", "productVariantID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "modifier", "ballanceAfterModifier", "comments")
);
