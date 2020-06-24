CREATE TABLE "PropertyUnit" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,	
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "display" STRING(45) NOT NULL,
    INDEX "PropertyUnit_idx_createdByID" ("createdByID"),
	INDEX "PropertyUnit_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "PropertyUnit_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyUnit_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "display")
);
