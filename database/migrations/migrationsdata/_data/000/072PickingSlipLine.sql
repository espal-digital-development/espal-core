CREATE TABLE "PickingSlipLine" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"shipmentLineID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"quantity" INT NOT NULL,
	"comments" STRING,
	INDEX "PickingSlipLine_idx_createdByID" ("createdByID"),
	INDEX "PickingSlipLine_idx_updatedByID" ("updatedByID"),
	INDEX "PickingSlipLine_idx_shipmentLineID" ("shipmentLineID"),
	CONSTRAINT "PickingSlipLine_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PickingSlipLine_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PickingSlipLine_shipmentLineID" FOREIGN KEY ("shipmentLineID") REFERENCES "ShipmentLine" ("id"),
	FAMILY "Primary" ("id", "createdByID", "shipmentLineID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "quantity", "comments")
);
