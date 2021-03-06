CREATE TABLE "SaleOffer" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"domainID" UUID NOT NULL,
	"userID" UUID NOT NULL,
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
	"shippingAddressBusiness" BOOL NOT NULL,
	"shippingAddressBusinessCOCNumber" STRING(255),
	"shippingAddressFirstName" STRING(50) NOT NULL,
	"shippingAddressSurname" STRING(50) NOT NULL,
	"shippingAddressStreet" STRING(72) NOT NULL,
	"shippingAddressStreetLine2" STRING(72),
	"shippingAddressNumber" STRING(12),
	"shippingAddressNumberAddition" STRING(12),
	"shippingAddressZipCode" STRING(12),
	"shippingAddressCity" STRING(255) NOT NULL,
	"shippingAddressState" INT,
	"shippingAddressCountry" INT2,
	"shippingAddressPhoneNumber" STRING(32),
	"shippingAddressEmail" STRING(255),
	"comments" STRING,
	"sellingPartyAutograph" BYTES,
	"buyingPartyAutograph" BYTES,
	INDEX "SaleOffer_idx_createdByID" ("createdByID"),
	INDEX "SaleOffer_idx_updatedByID" ("updatedByID"),
	INDEX "SaleOffer_idx_domainID" ("domainID"),
	INDEX "SaleOffer_idx_userID" ("userID"),
	CONSTRAINT "SaleOffer_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "SaleOffer_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "SaleOffer_domainID" FOREIGN KEY ("domainID") REFERENCES "Domain" ("id"),
	CONSTRAINT "SaleOffer_userID" FOREIGN KEY ("userID") REFERENCES "User" ("id"),
    FAMILY "Primary" ("id", "createdByID", "domainID", "userID", "createdAt", "currency", "code"),
	FAMILY "Secondary" ("updatedByID", "updatedAt"),
	FAMILY "Tertiary" ("comments"),
	FAMILY "Autopgrahs" ("sellingPartyAutograph", "buyingPartyAutograph"),
	FAMILY "Info" ("userInfoBusiness", "userInfoBusinessCOCNumber", "userInfoFirstName",
		"userInfoSurname", "userInfoStreet", "userInfoStreetLine2", "userInfoNumber",
        "userInfoNumberAddition", "userInfoZipCode", "userInfoCity", "userInfoCountry",
        "userInfoPhoneNumber", "userInfoEmail"),
	FAMILY "ShippingInfo" ("shippingAddressBusiness", "shippingAddressBusinessCOCNumber",
		"shippingAddressFirstName", "shippingAddressSurname", "shippingAddressStreet",
        "shippingAddressStreetLine2", "shippingAddressNumber", "shippingAddressNumberAddition",
        "shippingAddressZipCode", "shippingAddressCity", "shippingAddressState", "shippingAddressCountry",
        "shippingAddressPhoneNumber", "shippingAddressEmail")
);
