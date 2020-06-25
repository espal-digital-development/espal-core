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
func (c *HTTPContext) AdminMainMenu() []*adminmenu.Block {
	user, ok, err := c.GetUser()
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return nil
	}
	if !ok {
		c.loggerService.Errorf("no user found!")
		return nil
	}
	adminMenu, err := c.adminMenuService.GenerateAdminMenuStructure(user.ID(), c.language.ID())
	if err != nil {
		c.loggerService.Error(errors.ErrorStack(err))
		return nil
	}
	return adminMenu
}

// GetAdminCreateUpdateTitle returns the title for the create/update
// page for the id and subject given.
func (c *HTTPContext) GetAdminCreateUpdateTitle(id string, subject string) string {
	var displayTitle string
	if id == "" {
		displayTitle += c.Translate("create")
	} else {
		displayTitle += c.Translate("update")
	}
	displayTitle += " " + c.Translate(subject)
	return displayTitle
}
