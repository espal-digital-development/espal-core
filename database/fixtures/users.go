package fixtures

import (
	"strconv"

	"github.com/brianvoe/gofakeit"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

const userQuery = `INSERT INTO "User"("createdByID","active","language","firstName","surname","email","password",
	"currencies") VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`

func (f *Fixtures) users() error {
	// TODO :: 77777 Generate and print new passwords/emails on-the-fly for security
	// TODO :: 77777 The choice of making stub data should be a config option? Or a module to prevent dependencies?
	password, err := bcrypt.GenerateFromPassword([]byte("abc123"), 12)
	if err != nil {
		return errors.Trace(err)
	}

	// User
	row := f.inserterDatabase.QueryRow(userQuery, nil, true, f.dutchLanguage.ID(), "No", "One", "no@one.com",
		string(password), "")
	if err := row.Scan(&f.mainUserID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.updaterDatabase.Exec(`UPDATE "User" SET "createdByID" = "id" WHERE "id" = $1`,
		f.mainUserID); err != nil {
		return errors.Trace(err)
	}

	// Extra dummy users for pagination testing
	tx, err := f.inserterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	for i := 1; i <= 1e3; i++ {
		if _, err := tx.Exec(userQuery, f.mainUserID, true, f.englishLanguage.ID(), gofakeit.FirstName(),
			gofakeit.LastName(), gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 32),
			""); err != nil {
			return errors.Trace(err)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Trace(err)
	}

	var c uint
	for _, userRight := range f.userRightsRepository.AllByName() {
		if c > 0 {
			if _, err := f.userRightsBuffer.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		}
		if _, err := f.userRightsBuffer.WriteString(strconv.FormatUint(uint64(userRight), 10)); err != nil {
			return errors.Trace(err)
		}
		c++
	}

	// UserAddress
	if _, err := f.inserterDatabase.Exec(`INSERT INTO "UserAddress"("createdByID","userID","active","firstName",
		"surname","street","number","zipCode","city") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`, f.mainUserID,
		f.mainUserID, true, "No", "One", "Memory Lane", "7", "123456HI", "Villeville"); err != nil {
		return errors.Trace(err)
	}

	return nil
}
