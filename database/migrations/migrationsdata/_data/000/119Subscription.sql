CREATE TABLE "Subscription" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	INDEX "Subscription_idx_createdByID" ("createdByID"),
	INDEX "Subscription_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "Subscription_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Subscription_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt")
);
