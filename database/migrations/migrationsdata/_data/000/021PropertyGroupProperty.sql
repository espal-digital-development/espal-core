CREATE TABLE "PropertyGroupProperty" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"propertyGroupID" UUID NOT NULL,
	"propertyID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    UNIQUE INDEX "PropertyGroupProperty_uidx_propertyGroupID_propertyID" ("propertyGroupID", "propertyID"),
	INDEX "PropertyGroupProperty_idx_createdByID" ("createdByID"),
	INDEX "PropertyGroupProperty_idx_updatedByID" ("updatedByID"),
	INDEX "PropertyGroupProperty_idx_propertyID" ("propertyID"),
	CONSTRAINT "PropertyGroupProperty_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyGroupProperty_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PropertyGroupProperty_propertyGroupID" FOREIGN KEY ("propertyGroupID") REFERENCES "PropertyGroup" ("id"),
	CONSTRAINT "PropertyGroupProperty_propertyID" FOREIGN KEY ("propertyID") REFERENCES "Property" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt", "propertyGroupID", "propertyID", "updatedByID", "updatedAt")
);
