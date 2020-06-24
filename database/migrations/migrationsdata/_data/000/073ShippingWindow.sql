CREATE TABLE "ShippingWindow" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"userGroupID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"startDate" TIMESTAMP,
	"endDate" TIMESTAMP,
	INDEX "ShippingWindow_idx_createdByID" ("createdByID"),
	INDEX "ShippingWindow_idx_updatedByID" ("updatedByID"),
	INDEX "ShippingWindow_idx_userGroupID" ("userGroupID"),
	CONSTRAINT "ShippingWindow_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ShippingWindow_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ShippingWindow_userGroupID" FOREIGN KEY ("userGroupID") REFERENCES "UserGroup" ("id"),
	FAMILY "Primary" ("id", "createdByID", "userGroupID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "startDate", "endDate")
);
