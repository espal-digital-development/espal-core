CREATE TABLE "DiscountCode" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"key" STRING(255),
    "maxUses" INT,
    "usesCounter" INT NOT NULL,
    "availableFrom" TIMESTAMP,
    "availableUntil" TIMESTAMP,
    "discountPercentage" DECIMAL(9,6),
    "discountAmount" DECIMAL(9,2),
    UNIQUE INDEX "DiscountCode_uidx_key" ("key"),
	INDEX "DiscountCode_idx_createdByID" ("createdByID"),
	INDEX "DiscountCode_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "DiscountCode_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "DiscountCode_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "usesCounter"),
	FAMILY "Tertiary" ("key", "maxUses", "discountPercentage", "discountAmount"),
    FAMILY "Availability" ("availableFrom", "availableUntil")
);
