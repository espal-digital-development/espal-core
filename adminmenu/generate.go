package adminmenu

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

var _ MenuBlock = &Block{}
var _ MenuItem = &Item{}

// MenuBlock represents a menu block.
type MenuBlock interface {
	Title() string
	Items() []*Item
	// AccessRight() uint16
}

// Block menu block.
type Block struct {
	title string
	items []*Item
	// accessRight uint16
}

// Title returns the block title.
func (block *Block) Title() string {
	return block.title
}

// Items returns the block items.
func (block *Block) Items() []*Item {
	return block.items
}

// // AccessRight returns the block accessRight.
// func (block *block) AccessRight() uint16 {
// 	return block.accessRight
// }

// MenuItem represents a menu item.
type MenuItem interface {
	Title() string
	URL() string
	AccessRight() uint16
}

// Item holds a block entry for the main menu block.
type Item struct {
	title       string
	url         string
	accessRight uint16
}

// Title returns the menu item title.
func (item *Item) Title() string {
	return item.title
}

// URL returns the menu item url.
func (item *Item) URL() string {
	return item.url
}

// AccessRight returns the menu item accessRight.
func (item *Item) AccessRight() uint16 {
	return item.accessRight
}

type itemToGenerate struct {
	title       string
	titlePlural bool
	items       []string
}

// GenerateAdminMenuStructure generates and returns a rendered admin menu
// for the given user (based on it's userrights) and locale.
func (adminMenu *AdminMenu) GenerateAdminMenuStructure(userID string, localeID uint16) ([]*Block, error) {
	rows, err := adminMenu.selecterDatabase.Query(`SELECT ug."userRights" FROM "UserGroup" ug
		JOIN "UserGroupUser" uu ON uu."userGroupID" = ug."id" AND uu."userID" = $1
		WHERE ug."userRights" != ''`, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Trace(err)
	}

	rightsUserHas := map[uint16]bool{}
	var userRights string

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, errors.Trace(err)
		}
		if err := rows.Scan(&userRights); err != nil {
			return nil, errors.Trace(err)
		}
		splitRights := strings.Split(userRights, ",")

		for k := range splitRights {
			rightNumber, err := strconv.ParseUint(splitRights[k], 10, 16)
			if err != nil {
				return nil, errors.Trace(err)
			}
			rightsUserHas[uint16(rightNumber)] = true
		}
	}
	if err := rows.Close(); err != nil {
		return nil, errors.Trace(err)
	}

	// TODO :: 7 The blueprint will bleed memory lock on the Items. Needs a better version later
	var menu []*Block
	blueprint, err := adminMenu.generateAdminMenuForLocale(localeID)
	if err != nil {
		return nil, errors.Trace(err)
	}

	for k := range blueprint {
		menuBlock := &Block{
			title: blueprint[k].Title(),
			// accessRight: blueprint[k].AccessRight(),
			items: make([]*Item, 0),
		}

		for _, item := range blueprint[k].Items() {
			if ok := rightsUserHas[item.AccessRight()]; ok {
				menuBlock.items = append(menuBlock.items, item)
			}
		}

		if len(menuBlock.items) > 0 {
			menu = append(menu, menuBlock)
		}
	}

	return menu, nil
}

func (adminMenu *AdminMenu) generateAdminMenuForLocale(localeID uint16) ([]*Block, error) {
	adminURL := adminMenu.configService.AdminURL() + "/"
	blocks := make([]*Block, 0, len(adminMenu.itemToGenerate))

	for k := range adminMenu.itemToGenerate {
		items, err := adminMenu.generateItems(localeID, adminURL, adminMenu.itemToGenerate[k].items)
		if err != nil {
			return nil, errors.Trace(err)
		}

		block := &Block{
			items: items,
		}

		if adminMenu.itemToGenerate[k].titlePlural {
			block.title = adminMenu.translationsRepository.Plural(localeID, adminMenu.itemToGenerate[k].title)
		} else {
			block.title = adminMenu.translationsRepository.Singular(localeID, adminMenu.itemToGenerate[k].title)
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}

func (adminMenu *AdminMenu) generateItems(localeID uint16, adminURL string, list []string) ([]*Item, error) {
	items := make([]*Item, len(list))

	for k := range list {
		// Need to uppercase the first letter to make it match-up with the style of the UserRight name
		code, err := adminMenu.userRightsRepository.GetCode("Access" + strings.Title(list[k]))
		if err != nil {
			return nil, errors.Trace(err)
		}
		items[k] = &Item{
			title: adminMenu.translationsRepository.Plural(localeID, list[k]),
			// Need to uppercase the first letter to make it match-up with the style of the URLs
			url:         adminURL + strings.Title(list[k]),
			accessRight: code,
		}
	}

	return items, nil
}
