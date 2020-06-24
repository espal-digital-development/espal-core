CREATE TABLE "Menu" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"parentID" UUID,
	"slugID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"active" BOOL NOT NULL DEFAULT false,
	"sorting" INT NOT NULL DEFAULT 0,
	"externalLink" STRING(255),
	"internalLink" STRING(255),
	INDEX "Menu_idx_createdByID" ("createdByID"),
	INDEX "Menu_idx_updatedByID" ("updatedByID"),
	INDEX "Menu_idx_parentID" ("parentID"),
	INDEX "Menu_idx_slugID" ("slugID"),
	CONSTRAINT "Menu_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Menu_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Menu_parentID" FOREIGN KEY ("parentID") REFERENCES "Menu" ("id"),
	CONSTRAINT "Menu_slugID" FOREIGN KEY ("slugID") REFERENCES "Slug" ("id"),
	FAMILY "Primary" ("id", "createdByID", "parentID", "slugID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "active", "sorting", "externalLink", "internalLink")
);
