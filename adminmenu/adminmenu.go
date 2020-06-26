package adminmenu

import (
	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/database"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
)

var _ Menu = &AdminMenu{}

// Menu represents an object that can generate a menu.
type Menu interface {
	GenerateAdminMenuStructure(userID string, localeID uint16) ([]*Block, error)
}

// AdminMenu service object.
type AdminMenu struct {
	configService          config.Config
	selecterDatabase       database.Database
	translationsRepository translations.Repository
	userRightsRepository   userrights.Repository

	itemToGenerate []itemToGenerate
}

// New returns a new instance of AdminMenu.
func New(configService config.Config, selecterDatabase database.Database,
	translationsRepository translations.Repository,
	userRightsRepository userrights.Repository) *AdminMenu {
	return &AdminMenu{
		configService:          configService,
		selecterDatabase:       selecterDatabase,
		translationsRepository: translationsRepository,
		userRightsRepository:   userRightsRepository,

		itemToGenerate: []itemToGenerate{{
			title: "catalog",
			items: []string{"productModel", "productVariant", "bundledProduct", "subscription", "propertyGroup",
				"property", "productReview", "filter", "cart", "wishlist", "discountCode", "couponCode", "credit"},
		}, {
			title: "financial",
			items: []string{"saleOrder", "invoice", "purchaseOrder", "paymentTransaction", "paymentAccount",
				"paymentProvider", "priceGroup", "priceList", "shipmentCost", "tax"},
		}, {
			title:       "customerRelation",
			titlePlural: true,
			items: []string{"lead", "opportunity", "account", "contact", "report", "group", "person", "offer",
				"prospect", "campaign", "project"},
		}, {
			title:       "logistic",
			titlePlural: true,
			items: []string{"shipment", "returnOrder", "receiving", "pickingSlip", "shippingWindow",
				"deliveryMethod", "stock", "stockLocation", "supplier"},
		}, {
			title: "content",
			items: []string{"menu", "page", "block", "poll", "newsArticle", "blogPost", "forum", "office", "reseller",
				"frequentlyAskedQuestion", "emailTemplate", "newsletter", "gift", "giftWrapping", "media", "download"},
		}, {
			title: "productLifecycle",
			items: []string{"phase", "prototype", "sizescreen", "billOfMaterial"},
		}, {
			title: "technicalService",
			items: []string{"importProfile", "exportProfile", "swimmingLane", "dataMapping", "webservice"},
		}, {
			title:       "reportAndLog",
			titlePlural: true,
			items:       []string{"accessLog", "errorLog", "searchHistory", "emailLog"},
		}, {
			title: "core",
			items: []string{"user", "userGroup", "userTask", "shop", "slug", "tag", "site", "domain", "setting",
				"technicalStatistic"},
		}},
	}
}
