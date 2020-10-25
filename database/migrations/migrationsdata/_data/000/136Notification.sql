CREATE TABLE "Notification" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"target" STRING(50) NOT NULL,
	"key" STRING(125) NOT NULL,
	"value" STRING,
	INDEX "Notification_idx_createdByID" ("createdByID"),
	INDEX "Notification_idx_updatedByID" ("updatedByID"),
	INDEX "Notification_idx_createdAt" ("createdAt"),
	CONSTRAINT "Notification_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Notification_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "target", "key")
);
