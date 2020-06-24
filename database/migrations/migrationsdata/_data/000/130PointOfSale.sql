CREATE TABLE "PointOfSale" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"shopID" UUID NOT NULL,
	INDEX "PointOfSale_idx_createdByID" ("createdByID"),
	INDEX "PointOfSale_idx_updatedByID" ("updatedByID"),
	INDEX "PointOfSale_idx_shopID" ("shopID"),
	CONSTRAINT "PointOfSale_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PointOfSale_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PointOfSale_shopID" FOREIGN KEY ("shopID") REFERENCES "Shop" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "shopID")
);
