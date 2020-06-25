package fixtures

import (
	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const forumQuery = `INSERT INTO "Forum"("createdByID","active") VALUES($1,$2) RETURNING "id"`
const forumWithParentQuery = `INSERT INTO "Forum"("createdByID","parentID","active") VALUES($1,$2,$3) RETURNING "id"`
const forumTranslationQuery = `INSERT INTO "ForumTranslation"("createdByID","forumID","language","field","value") VALUES($1,$2,$3,$4,$5)`
const forumPostQuery = `INSERT INTO "ForumPost"("createdByID","forumID","title","message") VALUES($1,$2,$3,$4)`

func (f *Fixtures) forums() error {
	// Forum 1
	var forum1ID string
	row := f.inserterDatabase.QueryRow(forumQuery, f.mainUserID, true)
	if err := row.Scan(&forum1ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumTranslationQuery, f.mainUserID, forum1ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Player versus Environment"); err != nil {
		return errors.Trace(err)
	}

	// Forum 2
	var forum2ID string
	row = f.inserterDatabase.QueryRow(forumWithParentQuery, f.mainUserID, forum1ID, true)
	if err := row.Scan(&forum2ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumTranslationQuery, f.mainUserID, forum2ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Hardcore Mode"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumPostQuery, f.mainUserID, forum2ID, "Factions Endgame builds", "Will post more soon!"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumPostQuery, f.mainUserID, forum2ID, "Nightfall Endgame builds", "Will post more soon!"); err != nil {
		return errors.Trace(err)
	}

	// Forum 3
	var forum3ID string
	row = f.inserterDatabase.QueryRow(forumQuery, f.mainUserID, true)
	if err := row.Scan(&forum3ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumTranslationQuery, f.mainUserID, forum3ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Player versus Player"); err != nil {
		return errors.Trace(err)
	}

	// Forum 4
	var forum4ID string
	row = f.inserterDatabase.QueryRow(forumWithParentQuery, f.mainUserID, forum3ID, true)
	if err := row.Scan(&forum4ID); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumTranslationQuery, f.mainUserID, forum4ID, f.englishLanguage.ID(), database.DBTranslationFieldName, "Heroes Ascent"); err != nil {
		return errors.Trace(err)
	}
	if _, err := f.inserterDatabase.Exec(forumPostQuery, f.mainUserID, forum4ID, `"Incoming!" Hall of Heroes Build`, "Will post more soon!"); err != nil {
		return errors.Trace(err)
	}

	return nil
}
