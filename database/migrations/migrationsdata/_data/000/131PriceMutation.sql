CREATE TABLE "PriceMutation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"productVariantID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"price" DECIMAL(9,2) NOT NULL,
	INDEX "PriceMutation_idx_createdByID" ("createdByID"),
	INDEX "PriceMutation_idx_updatedByID" ("updatedByID"),
	INDEX "PriceMutation_idx_productVariantID" ("productVariantID"),
	CONSTRAINT "PriceMutation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PriceMutation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PriceMutation_productVariantID" FOREIGN KEY ("productVariantID") REFERENCES "ProductVariant" ("id"),
	FAMILY "Primary" ("id", "createdByID", "productVariantID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "price")
);
