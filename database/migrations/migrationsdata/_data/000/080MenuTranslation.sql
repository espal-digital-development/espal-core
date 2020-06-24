CREATE TABLE "MenuTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"menuID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "MenuTranslation_uidx_menuID_language_field" ("menuID", "language", "field"),
	INDEX "MenuTranslation_idx_createdByID" ("createdByID"),
	INDEX "MenuTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "MenuTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "MenuTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "MenuTranslation_menuID" FOREIGN KEY ("menuID") REFERENCES "Menu" ("id"),
	FAMILY "Primary" ("id", "createdByID", "menuID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
