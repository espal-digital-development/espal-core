// Code generated by qtc from "create.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pages/forum/post/create/create.qtpl:1
package create

//line pages/forum/post/create/create.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pages/forum/post/create/create.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pages/forum/post/create/create.qtpl:1
func (p *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line pages/forum/post/create/create.qtpl:1
	qw422016.E().S(p.forum.Name())
//line pages/forum/post/create/create.qtpl:1
}

//line pages/forum/post/create/create.qtpl:1
func (p *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line pages/forum/post/create/create.qtpl:1
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/forum/post/create/create.qtpl:1
	p.StreamTitle(qw422016)
//line pages/forum/post/create/create.qtpl:1
	qt422016.ReleaseWriter(qw422016)
//line pages/forum/post/create/create.qtpl:1
}

//line pages/forum/post/create/create.qtpl:1
func (p *Page) Title() string {
//line pages/forum/post/create/create.qtpl:1
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/forum/post/create/create.qtpl:1
	p.WriteTitle(qb422016)
//line pages/forum/post/create/create.qtpl:1
	qs422016 := string(qb422016.B)
//line pages/forum/post/create/create.qtpl:1
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/forum/post/create/create.qtpl:1
	return qs422016
//line pages/forum/post/create/create.qtpl:1
}

//line pages/forum/post/create/create.qtpl:3
func (p *Page) StreamContent(qw422016 *qt422016.Writer) {
//line pages/forum/post/create/create.qtpl:3
	qw422016.N().S(`<div class="content">`)
//line pages/forum/post/create/create.qtpl:5
	qw422016.N().S(`CREATE</div>`)
//line pages/forum/post/create/create.qtpl:9
}

//line pages/forum/post/create/create.qtpl:9
func (p *Page) WriteContent(qq422016 qtio422016.Writer) {
//line pages/forum/post/create/create.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/forum/post/create/create.qtpl:9
	p.StreamContent(qw422016)
//line pages/forum/post/create/create.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line pages/forum/post/create/create.qtpl:9
}

//line pages/forum/post/create/create.qtpl:9
func (p *Page) Content() string {
//line pages/forum/post/create/create.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/forum/post/create/create.qtpl:9
	p.WriteContent(qb422016)
//line pages/forum/post/create/create.qtpl:9
	qs422016 := string(qb422016.B)
//line pages/forum/post/create/create.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/forum/post/create/create.qtpl:9
	return qs422016
//line pages/forum/post/create/create.qtpl:9
}
