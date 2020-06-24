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

func (runner *Runner) routesForum() error {
	if err := runner.services.router.RegisterRoute("/Forums", overview.New(runner.stores.forum, overviewpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/Forum", forum.New(runner.repositories.regularExpressions, runner.stores.forum, forumpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ForumPost", post.New(runner.repositories.regularExpressions, runner.stores.forum, postpage.New(runner.services.renderer))); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ForumPostCreate", create.New(runner.repositories.regularExpressions, runner.stores.forum, postcreatepage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ForumPostEdit", edit.New(runner.repositories.regularExpressions, runner.stores.forum, posteditpage.New())); err != nil {
		return errors.Trace(err)
	}
	if err := runner.services.router.RegisterRoute("/ForumPostDelete", delete.New(runner.stores.forum)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
