CREATE TABLE "PaymentAccount" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"paymentMethodID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"name" STRING(255) NOT NULL,
	"username" STRING(255),
	"passphrase" STRING(255),
	"secretKey" STRING(255),
	"publicPey" STRING(255),
	"certificate" STRING(255),
	INDEX "PaymentAccount_idx_createdByID" ("createdByID"),
	INDEX "PaymentAccount_idx_updatedByID" ("updatedByID"),
	INDEX "PaymentAccount_idx_paymentMethodID" ("paymentMethodID"),
	CONSTRAINT "PaymentAccount_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PaymentAccount_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PaymentAccount_paymentMethodID" FOREIGN KEY ("paymentMethodID") REFERENCES "PaymentMethod" ("id"),
	FAMILY "Primary" ("id", "createdByID", "paymentMethodID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "name", "username", "passphrase", "secretKey",
		"publicPey", "certificate")
);
