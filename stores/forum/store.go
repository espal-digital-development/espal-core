package forum

import (
	"database/sql"
	errorsNative "errors"

	"github.com/espal-digital-development/espal-core/database"
	"github.com/juju/errors"
)

const (
	getForParentQuery = `SELECT
			f."id", f."createdByID", f."updatedByID", u."firstName", u."surname", uu."firstName" AS fn,
			uu."surname" AS sn, ft."value",
			(SELECT COUNT(*) FROM "ForumPost" WHERE "forumID" = f."id" AND "responseToID" IS NULL) AS TopicCount,
			(SELECT COUNT(*) FROM "ForumPost" WHERE "forumID" = f."id" AND "responseToID" IS NOT NULL) AS PostCount
		FROM "Forum" f
		LEFT JOIN "ForumTranslation" ft ON ft."forumID" = f."id" AND ft."language" = $1 AND ft."field" = $2
		LEFT JOIN "User" u ON u."id" = f."createdByID" LEFT JOIN "User" uu ON uu."id" = f."updatedByID"
		WHERE f."parentID" = $3 AND f."active" = true ORDER BY f."sorting" LIMIT 10`
	getForParentTopLevelQuery = `SELECT
			f."id", f."createdByID", f."updatedByID", u."firstName", u."surname", uu."firstName" AS fn,
			uu."surname" AS sn, ft."value",
			(SELECT COUNT(*) FROM "ForumPost" WHERE "forumID" = f."id" AND "responseToID" IS NULL) AS TopicCount,
			(SELECT COUNT(*) FROM "ForumPost" WHERE "forumID" = f."id" AND "responseToID" IS NOT NULL) AS PostCount
		FROM "Forum" f
		LEFT JOIN "ForumTranslation" ft ON ft."forumID" = f."id" AND ft."language" = $1 AND ft."field" = $2
		LEFT JOIN "User" u ON u."id" = f."createdByID" LEFT JOIN "User" uu ON uu."id" = f."updatedByID"
		WHERE f."parentID" IS NULL AND f."active" = true ORDER BY f."sorting" LIMIT 10`
)

// Language type interface.
type Language interface {
	ID() uint16
	Code() string
	Translate(uint16) string
}

// ForumsStore data store.
type ForumsStore struct {
	selecterDatabase database.Database
	deletorDatabase  database.Database
}

// GetOneByID fetches by ID.
func (s *ForumsStore) GetOneByID(forumID string, language Language) (*Forum, bool, error) {
	forum := newForum()
	var nameTranslation *string
	err := s.selecterDatabase.QueryRow(`
		SELECT
			f."id", f."createdByID", u."updatedByID", u."firstName",
			u."surname", uu."firstName" AS fn, uu."surname" AS sn,
			ft."value",
			(SELECT COUNT(*) FROM "ForumPost" WHERE "forumID" = f."id"
				AND "responseToID" IS NULL) AS topicCount,
			(SELECT COUNT(*) FROM "ForumPost" WHERE "forumID" = f."id"
				AND "responseToID" IS NULL) AS postCount
		FROM "Forum" f
		LEFT JOIN "ForumTranslation" ft ON ft."forumID" = f."id"
			AND ft."language" = $1 AND ft."field" = $2
		LEFT JOIN "User" u ON u."id" = f."createdByID"
		LEFT JOIN "User" uu ON uu."id" = f."updatedByID"
		WHERE f."id" = $3 AND f."active" = true
		LIMIT 1
	`, language.ID(), database.DBTranslationFieldName, forumID).
		Scan(&forum.id, &forum.createdByID, &forum.updatedByID, &forum.createdByFirstName, &forum.createdBySurname,
			&forum.updatedByFirstName, &forum.updatedBySurname, &nameTranslation, &forum.topicsCount,
			&forum.postsCount)
	if errorsNative.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, errors.Trace(err)
	}
	if nameTranslation != nil {
		forum.SetName(*nameTranslation)
	}
	return forum, true, nil
}

// GetOnePostByID fetches by ID.
func (s *ForumsStore) GetOnePostByID(postID string) (*Post, bool, error) {
	post := newPost()
	err := s.selecterDatabase.QueryRow(`
		SELECT
			f."id", f."createdByID", f."updatedByID",
			u."firstName", u."surname",
			uu."firstName" AS fn, uu."surname" AS sn,
			f."sticky", f."title", f."message"
		FROM "ForumPost" f
		LEFT JOIN "User" u ON u."id" = f."createdByID"
		LEFT JOIN "User" uu ON uu."id" = f."updatedByID"
		WHERE f."id" = $1 AND f."responseToID" IS NULL
		ORDER BY f."sticky" ASC, f."updatedAt" DESC,
			f."createdAt" DESC
		LIMIT 10
	`, postID).Scan(&post.id, &post.createdByID, &post.updatedByID, &post.createdByFirstName, &post.createdBySurname,
		&post.updatedByFirstName, &post.updatedBySurname, &post.sticky, &post.title, &post.message)
	if err != nil && !errorsNative.Is(err, sql.ErrNoRows) {
		return post, false, errors.Trace(err)
	}
	return post, true, errors.Trace(err)
}

// GetTopLevel returns all top-level Forum entries.
func (s *ForumsStore) GetTopLevel(language Language) ([]*Forum, bool, error) {
	return s.GetForParent("", language)
}

// GetForParent returns all Forum entries for the given parent.
// nolint:nakedret
func (s *ForumsStore) GetForParent(parentID string, language Language) (result []*Forum, ok bool, err error) {
	var rows database.Rows
	if parentID == "" {
		rows, err = s.selecterDatabase.Query(getForParentTopLevelQuery, language.ID(), database.DBTranslationFieldName)
	} else {
		var parentIDPointer *string
		if parentID != "" {
			parentIDPointer = &parentID
		}
		rows, err = s.selecterDatabase.Query(getForParentQuery, language.ID(), database.DBTranslationFieldName,
			parentIDPointer)
	}
	if errorsNative.Is(err, sql.ErrNoRows) {
		err = nil
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	defer func(dbRows database.Rows) {
		closeErr := dbRows.Close()
		if err != nil && closeErr != nil {
			err = errors.Wrap(err, closeErr)
		} else if closeErr != nil {
			err = errors.Trace(closeErr)
		}
	}(rows)
	var nameTranslation *string
	result = make([]*Forum, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		forum := newForum()
		nameTranslation, err = s.scanDefaultFieldsIntoForum(rows, forum)
		if err != nil {
			err = errors.Trace(err)
			return
		}
		if nameTranslation != nil {
			forum.SetName(*nameTranslation)
		}
		result = append(result, forum)
	}
	ok = len(result) > 0
	return
}

// GetPosts returns the posts for the given Forum ID.
// nolint:nakedret
func (s *ForumsStore) GetPosts(forumID string) (posts []*Post, ok bool, err error) {
	rows, err := s.selecterDatabase.Query(`SELECT f."id", f."createdByID", f."updatedByID", u."firstName", u."surname",
			uu."firstName" AS fn, uu."surname" AS sn, f."sticky", f."title", f."message"
		FROM "ForumPost" f
		LEFT JOIN "User" u ON u."id" = f."createdByID"
		LEFT JOIN "User" uu ON uu."id" = f."updatedByID"
		WHERE f."forumID" = $1 AND f."responseToID" IS NULL ORDER BY f."sticky" ASC, f."updatedAt" DESC,
			f."createdAt" DESC LIMIT 10
	`, forumID)
	if errorsNative.Is(err, sql.ErrNoRows) {
		err = nil
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	defer func(dbRows database.Rows) {
		closeErr := dbRows.Close()
		if err != nil && closeErr != nil {
			err = errors.Wrap(err, closeErr)
		} else if closeErr != nil {
			err = errors.Trace(closeErr)
		}
	}(rows)
	posts = make([]*Post, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		post := newPost()
		err = rows.Scan(&post.id, &post.createdByID, &post.updatedByID, &post.createdByFirstName,
			&post.createdBySurname, &post.updatedByFirstName, &post.updatedBySurname, &post.sticky, &post.title,
			&post.message)
		if err != nil {
			err = errors.Trace(err)
			return
		}
		posts = append(posts, post)
	}
	ok = len(posts) > 0
	return
}

// GetPostReplies returns the posted replies for the given Post ID.
func (s *ForumsStore) GetPostReplies(postID string) (replies []*Post, ok bool, err error) {
	rows, err := s.selecterDatabase.Query(`SELECT f."id", f."createdByID", f."updatedByID", u."firstName",
			u."surname", uu."firstName" AS fn, uu."surname" AS sn, f."sticky", f."title", f."message"
		FROM "ForumPost" f
		LEFT JOIN "User" u ON u."id" = f."createdByID"
		LEFT JOIN "User" uu ON uu."id" = f."updatedByID"
		WHERE f."responseToID" = $1 ORDER BY f."sticky" ASC, f."updatedAt" DESC, f."createdAt" DESC LIMIT 10`, postID)
	if errorsNative.Is(err, sql.ErrNoRows) {
		err = nil
		return
	}
	if err != nil {
		err = errors.Trace(err)
		return
	}
	defer func(dbRows database.Rows) {
		closeErr := dbRows.Close()
		if err != nil && closeErr != nil {
			err = errors.Wrap(err, closeErr)
		} else if closeErr != nil {
			err = errors.Trace(closeErr)
		}
	}(rows)
	replies = make([]*Post, 0)
	for rows.Next() {
		if err = rows.Err(); err != nil {
			err = errors.Trace(err)
			return
		}
		reply := newPost()
		if err = s.scanDefaultFieldsIntoPost(rows, reply); err != nil {
			err = errors.Trace(err)
			return
		}
		replies = append(replies, reply)
	}
	ok = len(replies) > 0
	return
}

// GetForumIDForPostID returns the ForumID that belongs to the given ForumPost ID.
func (s *ForumsStore) GetForumIDForPostID(postID string) (string, bool, error) {
	var id string
	err := s.selecterDatabase.QueryRow(`SELECT "forumID" FROM "ForumPost" WHERE "id" = $1 LIMIT 1`, postID).Scan(&id)
	if err != nil && !errorsNative.Is(err, sql.ErrNoRows) {
		return id, false, errors.Trace(err)
	}
	return id, !errorsNative.Is(err, sql.ErrNoRows), nil
}

// DeleteOneForumPostByID deletes one ForumPost entry based on the given id.
func (s *ForumsStore) DeleteOneForumPostByID(id string) error {
	r, err := s.deletorDatabase.Exec(`DELETE FROM "ForumPost" WHERE "id" = $1`, id)
	if err != nil {
		return errors.Trace(err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return errors.Trace(err)
	}
	if rowsAffected == 0 {
		return errors.Errorf("nothing matched to be deleted")
	}
	return nil
}

func (s *ForumsStore) scanDefaultFieldsIntoForum(row database.Row, forum *Forum) (*string, error) {
	var nameTranslation *string
	err := row.Scan(&forum.id, &forum.createdByID, &forum.updatedByID, &forum.createdByFirstName,
		&forum.createdBySurname, &forum.updatedByFirstName, &forum.updatedBySurname, &nameTranslation,
		&forum.topicsCount, &forum.postsCount)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return nameTranslation, nil
}

func (s *ForumsStore) scanDefaultFieldsIntoPost(row database.Row, post *Post) error {
	err := row.Scan(&post.id, &post.createdByID, &post.updatedByID, &post.createdByFirstName, &post.createdBySurname,
		&post.updatedByFirstName, &post.updatedBySurname, &post.sticky, &post.title, &post.message)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
