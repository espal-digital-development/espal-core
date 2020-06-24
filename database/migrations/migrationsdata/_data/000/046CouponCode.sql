CREATE TABLE "CouponCode" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"claimedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"key" STRING(255) NOT NULL,
    UNIQUE INDEX "CouponCode_uidx_key" ("key"),
	INDEX "CouponCode_idx_createdByID" ("createdByID"),
	INDEX "CouponCode_idx_updatedByID" ("updatedByID"),
	INDEX "CouponCode_idx_claimedByID" ("claimedByID"),
	CONSTRAINT "CouponCode_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "CouponCode_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "CouponCode_userID" FOREIGN KEY ("claimedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "claimedByID", "key")
);
