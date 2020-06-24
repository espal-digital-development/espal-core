CREATE TABLE "ReturnOrderReceiving" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"comments" STRING,
	INDEX "ReturnOrderReceiving_idx_createdByID" ("createdByID"),
	INDEX "ReturnOrderReceiving_idx_updatedByID" ("updatedByID"),
	INDEX "ReturnOrderReceiving_idx_domainID" ("domainID"),
	CONSTRAINT "ReturnOrderReceiving_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ReturnOrderReceiving_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ReturnOrderReceiving_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	FAMILY "Primary" ("id", "createdByID", "domainID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "comments")
);
