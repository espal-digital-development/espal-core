CREATE TABLE "DeliveryMethodTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"deliveryMethodID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "language" INT2 NOT NULL,
    "field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "DeliveryMethodTranslation_uidx_deliveryMethodID_language_field" ("deliveryMethodID", "language", "field"),
	INDEX "DeliveryMethodTranslation_idx_createdByID" ("createdByID"),
	INDEX "DeliveryMethodTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "DeliveryMethodTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "DeliveryMethodTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
    CONSTRAINT "DeliveryMethodTranslation_deliveryMethodID" FOREIGN KEY ("deliveryMethodID") REFERENCES "DeliveryMethod" ("id"),
	FAMILY "Primary" ("id", "createdByID", "deliveryMethodID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
