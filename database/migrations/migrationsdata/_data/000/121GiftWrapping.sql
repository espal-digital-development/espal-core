CREATE TABLE "GiftWrapping" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	INDEX "GiftWrapping_idx_createdByID" ("createdByID"),
	INDEX "GiftWrapping_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "GiftWrapping_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "GiftWrapping_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt")
);
