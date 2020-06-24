CREATE TABLE "SiteUser" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"siteID" UUID NOT NULL,
    "userID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    UNIQUE INDEX "SiteUser_uidx_siteID_userID" ("siteID", "userID"),
	INDEX "SiteUser_idx_createdByID" ("createdByID"),
	INDEX "SiteUser_idx_updatedByID" ("updatedByID"),
	INDEX "SiteUser_idx_userID" ("userID"),
    CONSTRAINT "SiteUser_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "SiteUser_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "SiteUser_siteID" FOREIGN KEY ("siteID") REFERENCES "Site" ("id"),
    CONSTRAINT "SiteUser_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "updatedByID", "siteID", "userID", "createdAt", "updatedAt")
);
