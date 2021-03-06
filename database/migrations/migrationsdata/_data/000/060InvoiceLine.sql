CREATE TABLE "InvoiceLine" (
	"id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	"createdByID" UUID NOT NULL,
	"updatedByID" UUID,
	"invoiceID" UUID NOT NULL,
	"saleOrderLineID" UUID,
	"productVariantID" UUID,
	"bundledProductID" UUID,
	"createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updatedAt" TIMESTAMP,
	"sorting" INT NOT NULL DEFAULT 0,
	"quantity" INT NOT NULL,
	"price" DECIMAL(9,2) NOT NULL,
	"vatPercentage" DECIMAL(9,6) NOT NULL,
	"comments" STRING,
	INDEX "InvoiceLine_idx_createdByID" ("createdByID"),
	INDEX "InvoiceLine_idx_updatedByID" ("updatedByID"),
	INDEX "InvoiceLine_idx_invoiceID" ("invoiceID"),
	INDEX "InvoiceLine_idx_saleOrderLineID" ("saleOrderLineID"),
	INDEX "InvoiceLine_idx_productVariantID" ("productVariantID"),
	INDEX "InvoiceLine_idx_bundledProductID" ("bundledProductID"),
	CONSTRAINT "InvoiceLine_createdByID" FOREIGN KEY ("createdByID") REFERENCES "User" ("id"),
	CONSTRAINT "InvoiceLine_updatedByID" FOREIGN KEY ("updatedByID") REFERENCES "User" ("id"),
	CONSTRAINT "InvoiceLine_invoiceID" FOREIGN KEY ("invoiceID") REFERENCES "Invoice" ("id"),
	CONSTRAINT "InvoiceLine_saleOrderLineID" FOREIGN KEY ("saleOrderLineID") REFERENCES "SaleOrderLine" ("id"),
	CONSTRAINT "InvoiceLine_productVariantID" FOREIGN KEY ("productVariantID") REFERENCES "ProductVariant" ("id"),
	CONSTRAINT "InvoiceLine_bundledProductID" FOREIGN KEY ("bundledProductID") REFERENCES "BundledProduct" ("id"),
	FAMILY "Primary" ("id", "createdByID", "invoiceID", "saleOrderLineID", "productVariantID", "bundledProductID", "createdAt"),
	FAMILY "Secondary" ("updatedByID", "updatedAt"),
	FAMILY "Tertiary" ("sorting", "comments"),
	FAMILY "QuantityPrice" ("quantity", "price", "vatPercentage")
);
