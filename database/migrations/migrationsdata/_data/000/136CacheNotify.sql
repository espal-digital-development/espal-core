CREATE TABLE "CacheNotify" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID,
	"updatedByID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"target" INT NOT NULL,
	"key" STRING(50) NOT NULL,
	INDEX "CacheNotify_idx_createdByID" ("createdByID"),
	INDEX "CacheNotify_idx_updatedByID" ("updatedByID"),
	INDEX "CacheNotify_idx_createdAt" ("createdAt"),
	CONSTRAINT "CacheNotify_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "CacheNotify_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	FAMILY "Primary" ("id", "createdByID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "target", "key")
);
