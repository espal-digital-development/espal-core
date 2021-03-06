CREATE TABLE "Invoice" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID NOT NULL,
	"userID" UUID NOT NULL,
	"saleOrderID" UUID NOT NULL,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"currency" INT2 NOT NULL,
	"code" STRING(255),
	"userInfoBusiness" BOOL NOT NULL,
	"userInfoBusinessCOCNumber" STRING(255),
	"userInfoFirstName" STRING(50) NOT NULL,
	"userInfoSurname" STRING(50) NOT NULL,
	"userInfoStreet" STRING(72) NOT NULL,
	"userInfoStreetLine2" STRING(72),
	"userInfoNumber" STRING(12),
	"userInfoNumberAddition" STRING(12),
	"userInfoZipCode" STRING(12),
	"userInfoCity" STRING(255) NOT NULL,
	"userInfoState" INT,
	"userInfoCountry" INT2,
	"userInfoPhoneNumber" STRING(32),
	"userInfoEmail" STRING(255),
	"comments" STRING,
	"sellingPartyAutograph" BYTES,
	"buyingPartyAutograph" BYTES,
	INDEX "Invoice_idx_createdByID" ("createdByID"),
	INDEX "Invoice_idx_updatedByID" ("updatedByID"),
	INDEX "Invoice_idx_domainID" ("domainID"),
	INDEX "Invoice_idx_userID" ("userID"),
	INDEX "Invoice_idx_saleOrderID" ("saleOrderID"),
	CONSTRAINT "Invoice_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "Invoice_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "Invoice_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	CONSTRAINT "Invoice_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
	CONSTRAINT "Invoice_saleOrderID" FOREIGN KEY ("saleOrderID") REFERENCES "SaleOrder" ("id"),
	FAMILY "Primary" ("id", "createdByID", "domainID", "userID", "saleOrderID", "createdAt", "currency", "code"),
	FAMILY "Secondary" ("updatedByID", "updatedAt"),
	FAMILY "Tertiary" ("comments"),
	FAMILY "Autopgrahs" ("sellingPartyAutograph", "buyingPartyAutograph"),
	FAMILY "Info" ("userInfoBusiness", "userInfoBusinessCOCNumber", "userInfoFirstName",
		"userInfoSurname", "userInfoStreet", "userInfoStreetLine2", "userInfoNumber",
        "userInfoNumberAddition", "userInfoZipCode", "userInfoCity", "userInfoCountry",
        "userInfoPhoneNumber", "userInfoEmail")
);
