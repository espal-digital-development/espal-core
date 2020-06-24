CREATE TABLE "ForumTranslation" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"forumID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"language" INT2 NOT NULL,
	"field" INT2 NOT NULL,
	"value" STRING NOT NULL,
	UNIQUE INDEX "ForumTranslation_uidx_forumID_language_field" ("forumID", "language", "field"),
	INDEX "ForumTranslation_idx_createdByID" ("createdByID"),
	INDEX "ForumTranslation_idx_updatedByID" ("updatedByID"),
	CONSTRAINT "ForumTranslation_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "ForumTranslation_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "ForumTranslation_forumID" FOREIGN KEY ("forumID") REFERENCES "Forum" ("id"),
	FAMILY "Primary" ("id", "createdByID", "forumID", "createdAt", "language", "field"),
	FAMILY "Secondary" ("updatedByID", "updatedAt", "value")
);
