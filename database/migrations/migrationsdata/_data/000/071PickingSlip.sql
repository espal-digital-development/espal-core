CREATE TABLE "PickingSlip" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"comments" STRING,
	INDEX "PickingSlip_idx_createdByID" ("createdByID"),
	INDEX "PickingSlip_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PickingSlip_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PickingSlip_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "comments")
);
