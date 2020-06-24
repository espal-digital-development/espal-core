CREATE TABLE "PaymentMethod" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"name" STRING(255) NOT NULL,
	"description" STRING(255),
	INDEX "PaymentMethod_idx_createdByID" ("createdByID"),
	INDEX "PaymentMethod_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PaymentMethod_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PaymentMethod_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "name", "description")
);
