CREATE TABLE "ReturnOrderReceivingLine" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"returnOrderReceivingID" UUID NOT NULL,
	"returnOrderLineID" UUID,
	"productVariantID" UUID,
	"bundledProductID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"quantity" INT NOT NULL,
	"comments" STRING,
	INDEX "ReturnOrderReceivingLine_idx_createdByID" ("createdByID"),
	INDEX "ReturnOrderReceivingLine_idx_updatedByID" ("updatedByID"),
	INDEX "ReturnOrderReceivingLine_idx_returnOrderReceivingID" ("returnOrderReceivingID"),
	INDEX "ReturnOrderReceivingLine_idx_returnOrderLineID" ("returnOrderLineID"),
	INDEX "ReturnOrderReceivingLine_idx_productVariantID" ("productVariantID"),
	INDEX "ReturnOrderReceivingLine_idx_bundledProductID" ("bundledProductID"),
	CONSTRAINT "ReturnOrderReceivingLine_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ReturnOrderReceivingLine_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ReturnOrderReceivingLine_returnOrderReceivingID" FOREIGN KEY ("returnOrderReceivingID") REFERENCES "ReturnOrderReceiving" ("id"),
	CONSTRAINT "ReturnOrderReceivingLine_returnOrderLineID" FOREIGN KEY ("returnOrderLineID") REFERENCES "ReturnOrderLine" ("id"),
	CONSTRAINT "ReturnOrderReceivingLine_productVariantID" FOREIGN KEY ("productVariantID") REFERENCES "ProductVariant" ("id"),
	CONSTRAINT "ReturnOrderReceivingLine_bundledProductID" FOREIGN KEY ("bundledProductID") REFERENCES "BundledProduct" ("id"),
	FAMILY "Primary" ("id", "createdByID", "returnOrderReceivingID", "returnOrderLineID", "productVariantID", "bundledProductID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "quantity", "comments")
);
