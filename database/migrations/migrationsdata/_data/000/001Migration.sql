CREATE TABLE "Migration" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"revision" STRING(3),
	UNIQUE INDEX "Migration_uidx_revision" ("revision"),
	FAMILY "Primary" ("id", "createdAt", "revision")
);
