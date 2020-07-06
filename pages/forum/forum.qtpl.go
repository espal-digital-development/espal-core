// Code generated by qtc from "forum.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pages/forum/forum.qtpl:1
package forum

//line pages/forum/forum.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pages/forum/forum.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pages/forum/forum.qtpl:1
func (p *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line pages/forum/forum.qtpl:1
	qw422016.E().S(p.forum.Name())
//line pages/forum/forum.qtpl:1
}

//line pages/forum/forum.qtpl:1
func (p *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line pages/forum/forum.qtpl:1
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/forum/forum.qtpl:1
	p.StreamTitle(qw422016)
//line pages/forum/forum.qtpl:1
	qt422016.ReleaseWriter(qw422016)
//line pages/forum/forum.qtpl:1
}

//line pages/forum/forum.qtpl:1
func (p *Page) Title() string {
//line pages/forum/forum.qtpl:1
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/forum/forum.qtpl:1
	p.WriteTitle(qb422016)
//line pages/forum/forum.qtpl:1
	qs422016 := string(qb422016.B)
//line pages/forum/forum.qtpl:1
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/forum/forum.qtpl:1
	return qs422016
//line pages/forum/forum.qtpl:1
}

//line pages/forum/forum.qtpl:3
func (p *Page) StreamContent(qw422016 *qt422016.Writer) {
//line pages/forum/forum.qtpl:3
	qw422016.N().S(`<div class="content">`)
//line pages/forum/forum.qtpl:5
	if len(p.posts) > 0 || len(p.forums) > 0 {
//line pages/forum/forum.qtpl:6
		if len(p.posts) > 0 {
//line pages/forum/forum.qtpl:6
			qw422016.N().S(`<p>`)
//line pages/forum/forum.qtpl:7
			qw422016.E().S(p.TranslatePlural("post"))
//line pages/forum/forum.qtpl:7
			qw422016.N().S(`</p><table>`)
//line pages/forum/forum.qtpl:9
			for k := range p.posts {
//line pages/forum/forum.qtpl:9
				qw422016.N().S(`<tr><td><table><tr><td>`)
//line pages/forum/forum.qtpl:14
				qw422016.E().S(p.posts[k].ID())
//line pages/forum/forum.qtpl:14
				qw422016.N().S(`</td><td><a href="/ForumPost?id=`)
//line pages/forum/forum.qtpl:16
				qw422016.E().S(p.posts[k].ID())
//line pages/forum/forum.qtpl:16
				qw422016.N().S(`">`)
//line pages/forum/forum.qtpl:17
				if p.posts[k].Title() != nil {
//line pages/forum/forum.qtpl:18
					qw422016.E().S(*p.posts[k].Title())
//line pages/forum/forum.qtpl:19
				}
//line pages/forum/forum.qtpl:19
				qw422016.N().S(`</a></td><td>`)
//line pages/forum/forum.qtpl:23
				qw422016.E().S(p.rendererService.CreatedBy(p.posts[k], p.language.ID()))
//line pages/forum/forum.qtpl:23
				qw422016.N().S(`</td></tr></table></td></tr>`)
//line pages/forum/forum.qtpl:29
			}
//line pages/forum/forum.qtpl:29
			qw422016.N().S(`</table>`)
//line pages/forum/forum.qtpl:31
		}
//line pages/forum/forum.qtpl:33
		if len(p.forums) > 0 {
//line pages/forum/forum.qtpl:33
			qw422016.N().S(`<p>`)
//line pages/forum/forum.qtpl:34
			qw422016.E().S(p.TranslatePlural("forum"))
//line pages/forum/forum.qtpl:34
			qw422016.N().S(`</p><table>`)
//line pages/forum/forum.qtpl:36
			for k := range p.forums {
//line pages/forum/forum.qtpl:36
				qw422016.N().S(`<tr><td><table><tr><th></th><th>`)
//line pages/forum/forum.qtpl:42
				qw422016.E().S(p.TranslatePlural("topic"))
//line pages/forum/forum.qtpl:42
				qw422016.N().S(`</th><th>`)
//line pages/forum/forum.qtpl:43
				qw422016.E().S(p.TranslatePlural("post"))
//line pages/forum/forum.qtpl:43
				qw422016.N().S(`</th><th>`)
//line pages/forum/forum.qtpl:44
				qw422016.E().S(p.Translate("createdBy"))
//line pages/forum/forum.qtpl:44
				qw422016.N().S(`</th></tr><tr><td><a href="/Forum?id=`)
//line pages/forum/forum.qtpl:48
				qw422016.E().S(p.forums[k].ID())
//line pages/forum/forum.qtpl:48
				qw422016.N().S(`">`)
//line pages/forum/forum.qtpl:48
				qw422016.E().S(p.forums[k].Name())
//line pages/forum/forum.qtpl:48
				qw422016.N().S(`</a></td><td>`)
//line pages/forum/forum.qtpl:51
				qw422016.E().S(p.forums[k].TopicsCountAsString())
//line pages/forum/forum.qtpl:51
				qw422016.N().S(`</td><td>`)
//line pages/forum/forum.qtpl:54
				qw422016.E().S(p.forums[k].PostsCountAsString())
//line pages/forum/forum.qtpl:54
				qw422016.N().S(`</td><td>`)
//line pages/forum/forum.qtpl:57
				qw422016.E().S(p.rendererService.CreatedBy(p.forums[k], p.language.ID()))
//line pages/forum/forum.qtpl:57
				qw422016.N().S(`</td></tr></table></td></tr>`)
//line pages/forum/forum.qtpl:63
			}
//line pages/forum/forum.qtpl:63
			qw422016.N().S(`</table>`)
//line pages/forum/forum.qtpl:65
		}
//line pages/forum/forum.qtpl:66
	} else {
//line pages/forum/forum.qtpl:66
		qw422016.N().S(`<p>`)
//line pages/forum/forum.qtpl:67
		qw422016.E().S(p.TranslatePlural("thereIsNoPostOrForumToDisplay"))
//line pages/forum/forum.qtpl:67
		qw422016.N().S(`</p>`)
//line pages/forum/forum.qtpl:68
	}
//line pages/forum/forum.qtpl:68
	qw422016.N().S(`</div>`)
//line pages/forum/forum.qtpl:70
}

//line pages/forum/forum.qtpl:70
func (p *Page) WriteContent(qq422016 qtio422016.Writer) {
//line pages/forum/forum.qtpl:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/forum/forum.qtpl:70
	p.StreamContent(qw422016)
//line pages/forum/forum.qtpl:70
	qt422016.ReleaseWriter(qw422016)
//line pages/forum/forum.qtpl:70
}

//line pages/forum/forum.qtpl:70
func (p *Page) Content() string {
//line pages/forum/forum.qtpl:70
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/forum/forum.qtpl:70
	p.WriteContent(qb422016)
//line pages/forum/forum.qtpl:70
	qs422016 := string(qb422016.B)
//line pages/forum/forum.qtpl:70
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/forum/forum.qtpl:70
	return qs422016
//line pages/forum/forum.qtpl:70
}
