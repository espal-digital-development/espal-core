CREATE TABLE "PriceGroup" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "code" STRING(255) NOT NULL,
    "priority" INT NOT NULL DEFAULT 0,
	INDEX "PriceGroup_idx_createdByID" ("createdByID"),
	INDEX "PriceGroup_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PriceGroup_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PriceGroup_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "code", "priority")
);
