CREATE TABLE "PurchaseOrderReceiving" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"comments" STRING,
	INDEX "PurchaseOrderReceiving_idx_createdByID" ("createdByID"),
	INDEX "PurchaseOrderReceiving_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PurchaseOrderReceiving_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PurchaseOrderReceiving_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "comments")
);
