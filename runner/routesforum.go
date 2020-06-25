package runner

import (
	forumpage "github.com/espal-digital-development/espal-core/pages/forum"
	overviewpage "github.com/espal-digital-development/espal-core/pages/forum/overview"
	postpage "github.com/espal-digital-development/espal-core/pages/forum/post"
	postcreatepage "github.com/espal-digital-development/espal-core/pages/forum/post/create"
	posteditpage "github.com/espal-digital-development/espal-core/pages/forum/post/edit"
	"github.com/espal-digital-development/espal-core/routing/routes/forum"
	"github.com/espal-digital-development/espal-core/routing/routes/forum/overview"
	"github.com/espal-digital-development/espal-core/routing/routes/forum/post"
	"github.com/espal-digital-development/espal-core/routing/routes/forum/post/create"
	"github.com/espal-digital-development/espal-core/routing/routes/forum/post/delete"
	"github.com/espal-digital-development/espal-core/routing/routes/forum/post/edit"
	"github.com/juju/errors"
)

func (r *Runner) routesForum() error {
	if err := r.services.router.RegisterRoute("/Forums", overview.New(r.stores.forum, overviewpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/Forum", forum.New(r.repositories.regularExpressions, r.stores.forum, forumpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/ForumPost", post.New(r.repositories.regularExpressions, r.stores.forum, postpage.New(r.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/ForumPostCreate", create.New(r.repositories.regularExpressions, r.stores.forum, postcreatepage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/ForumPostEdit", edit.New(r.repositories.regularExpressions, r.stores.forum, posteditpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := r.services.router.RegisterRoute("/ForumPostDelete", delete.New(r.stores.forum)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
