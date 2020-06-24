package fixtures

import (
	"strconv"

	"github.com/brianvoe/gofakeit"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

const userQuery = `INSERT INTO "User"("createdByID","active","language","firstName","surname","email","password","currencies") VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`

func (fixtures *Fixtures) users() error {
	// TODO :: 77777 Generate and print new passwords/emails on-the-fly for security
	password, err := bcrypt.GenerateFromPassword([]byte("abc123"), 12)
	if err != nil {
		return errors.Trace(err)
	}

	// User
	row := fixtures.inserterDatabase.QueryRow(userQuery, nil, true, fixtures.dutchLanguage.ID(), "No", "One", "no@one.com", string(password), "")
	if err := row.Scan(&fixtures.mainUserID); err != nil {
		return errors.Trace(err)
	}
	if _, err := fixtures.updaterDatabase.Exec(`UPDATE "User" SET "createdByID" = "id" WHERE "id" = $1`, fixtures.mainUserID); err != nil {
		return errors.Trace(err)
	}

	// Extra dummy users for pagination testing
	tx, err := fixtures.inserterDatabase.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	for i := 1; i <= 1e3; i++ {
		if _, err := tx.Exec(userQuery, fixtures.mainUserID, true, fixtures.englishLanguage.ID(), gofakeit.FirstName(), gofakeit.LastName(), gofakeit.Email(), gofakeit.Password(true, true, true, true, true, 32), ""); err != nil {
			return errors.Trace(err)
		}
	}
	if err := tx.Commit(); err != nil {
		return errors.Trace(err)
	}

	var c uint
	for _, userRight := range fixtures.userRightsRepository.AllByName() {
		if c > 0 {
			if _, err := fixtures.userRightsBuffer.WriteString(","); err != nil {
				return errors.Trace(err)
			}
		}
		if _, err := fixtures.userRightsBuffer.WriteString(strconv.FormatUint(uint64(userRight), 10)); err != nil {
			return errors.Trace(err)
		}
		c++
	}

	// UserAddress
	if _, err := fixtures.inserterDatabase.Exec(`INSERT INTO "UserAddress"("createdByID","userID","active","firstName","surname","street","number","zipCode","city") VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`, fixtures.mainUserID, fixtures.mainUserID, true, "No", "One", "Memory Lane", "7", "123456HI", "Villeville"); err != nil {
		return errors.Trace(err)
	}

	return nil
}
