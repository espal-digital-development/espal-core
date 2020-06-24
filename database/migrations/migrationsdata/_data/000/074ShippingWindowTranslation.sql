CREATE TABLE "ShippingWindowTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"shippingWindowID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "ShippingWindowTranslation_uidx_shippingWindowID_language_field" ("shippingWindowID", "language", "field"),
	INDEX "ShippingWindowTranslation_idx_createdByID" ("createdByID"),
	INDEX "ShippingWindowTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "ShippingWindowTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ShippingWindowTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ShippingWindowTranslation_shippingWindowID" FOREIGN KEY ("shippingWindowID") REFERENCES "ShippingWindow" ("id"),
	FAMILY "Primary" ("id", "createdByID", "shippingWindowID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
