package modules

import (
	"github.com/espal-digital-development/espal-core/stores/cachenotify"
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/stores/session"
	"github.com/espal-digital-development/espal-core/stores/setting"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/stores/slug"
	"github.com/espal-digital-development/espal-core/stores/user"
	useraddress "github.com/espal-digital-development/espal-core/stores/user/address"
	usercontact "github.com/espal-digital-development/espal-core/stores/user/contact"
	usergroup "github.com/espal-digital-development/espal-core/stores/user/group"
)

// Stores represents an object that provides any known core stores of the system.
type Stores interface {
	Setting() setting.Store
	CacheNotify() cachenotify.Store
	Session() session.Store
	Domain() domain.Store
	Site() site.Store
	Slug() slug.Store
	UserGroup() usergroup.Store
	User() user.Store
	UserAddress() useraddress.Store
	UserContact() usercontact.Store
	Forum() forum.Store
}
