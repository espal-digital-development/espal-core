CREATE TABLE "Property" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"propertyUnitID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
    "sorting" INT NOT NULL DEFAULT 0,
    "key" STRING(255),
	"_type" INT2,
    "multiLingual" BOOL NOT NULL DEFAULT false,
    UNIQUE INDEX "Property_uidx_key" ("key"),
	INDEX "Property_idx_createdByID" ("createdByID"),
	INDEX "Property_idx_updatedByID" ("updatedByID"),
	INDEX "Property_idx_propertyUnitID" ("propertyUnitID"),
	CONSTRAINT "Property_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Property_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Property_propertyUnitID" FOREIGN KEY ("propertyUnitID") REFERENCES "PropertyUnit" ("id"),
    FAMILY "Primary" ("id", "createdByID", "createdAt", "propertyUnitID"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting", "key", "multiLingual")
);
