package contexts

import (
	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/juju/errors"
)

// AdminContext for admin page rendering.
type AdminContext interface {
	AdminMainMenu() []*adminmenu.Block
	GetAdminCreateUpdateTitle(string, string) string
}

// AdminMainMenu returns a localized version of the AdminMenu for the current user.
func (httpContext *HTTPContext) AdminMainMenu() []*adminmenu.Block {
	user, ok, err := httpContext.GetUser()
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return nil
	}
	if !ok {
		httpContext.loggerService.Errorf("no user found!")
		return nil
	}
	adminMenu, err := httpContext.adminMenuService.GenerateAdminMenuStructure(user.ID(), httpContext.language.ID())
	if err != nil {
		httpContext.loggerService.Error(errors.ErrorStack(err))
		return nil
	}
	return adminMenu
}

// GetAdminCreateUpdateTitle returns the title for the create/update
// page for the id and subject given.
func (httpContext *HTTPContext) GetAdminCreateUpdateTitle(id string, subject string) string {
	var displayTitle string
	if id == "" {
		displayTitle += httpContext.Translate("create")
	} else {
		displayTitle += httpContext.Translate("update")
	}
	displayTitle += " " + httpContext.Translate(subject)
	return displayTitle
}
