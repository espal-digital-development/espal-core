CREATE TABLE "DeliveryMethod" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"price" DECIMAL(9,2),
	INDEX "DeliveryMethod_idx_createdByID" ("createdByID"),
	INDEX "DeliveryMethod_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "DeliveryMethod_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "DeliveryMethod_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "price")
);
