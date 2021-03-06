CREATE TABLE "PersonalMessage" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"userID" UUID NOT NULL,
	"recipientID" UUID NOT NULL,
	"responseToID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
    "title" STRING(125),
    "message" STRING,
    INDEX "PersonalMessage_idx_createdByID" ("createdByID"),
	INDEX "PersonalMessage_idx_updatedByID" ("updatedByID"),
	INDEX "PersonalMessage_idx_userID" ("userID"),
	INDEX "PersonalMessage_idx_recipientID" ("recipientID"),
	INDEX "PersonalMessage_idx_responseToID" ("responseToID"),
    CONSTRAINT "PersonalMessage_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "PersonalMessage_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "PersonalMessage_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
	CONSTRAINT "PersonalMessage_recipientID" FOREIGN KEY ("recipientID") REFERENCES "User" ("id"),
	CONSTRAINT "PersonalMessage_responseToID" FOREIGN KEY ("responseToID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "userID", "recipientID", "responseToID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt"),
	FAMILY "Tertiary" ("title", "message")
);
