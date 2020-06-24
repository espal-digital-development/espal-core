// Code generated by qtc from "view.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line admin/domain/view/view.qtpl:1
package view

//line admin/domain/view/view.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin/domain/view/view.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin/domain/view/view.qtpl:1
func (page *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line admin/domain/view/view.qtpl:1
	qw422016.E().S(page.TranslatePlural("domain"))
//line admin/domain/view/view.qtpl:1
	qw422016.N().S(` > `)
//line admin/domain/view/view.qtpl:1
	qw422016.E().S(page.domain.Host())
//line admin/domain/view/view.qtpl:1
}

//line admin/domain/view/view.qtpl:1
func (page *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line admin/domain/view/view.qtpl:1
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/domain/view/view.qtpl:1
	page.StreamTitle(qw422016)
//line admin/domain/view/view.qtpl:1
	qt422016.ReleaseWriter(qw422016)
//line admin/domain/view/view.qtpl:1
}

//line admin/domain/view/view.qtpl:1
func (page *Page) Title() string {
//line admin/domain/view/view.qtpl:1
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/domain/view/view.qtpl:1
	page.WriteTitle(qb422016)
//line admin/domain/view/view.qtpl:1
	qs422016 := string(qb422016.B)
//line admin/domain/view/view.qtpl:1
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/domain/view/view.qtpl:1
	return qs422016
//line admin/domain/view/view.qtpl:1
}

//line admin/domain/view/view.qtpl:3
func (page *Page) StreamContent(qw422016 *qt422016.Writer) {
//line admin/domain/view/view.qtpl:3
	qw422016.N().S(`<main class="content"><h1>`)
//line admin/domain/view/view.qtpl:5
	qw422016.E().S(page.TranslatePlural("domain"))
//line admin/domain/view/view.qtpl:5
	qw422016.E().S(" > ")
//line admin/domain/view/view.qtpl:5
	qw422016.E().S(page.domain.Host())
//line admin/domain/view/view.qtpl:5
	qw422016.N().S(`</h1><p><a href="`)
//line admin/domain/view/view.qtpl:8
	qw422016.E().S(page.AdminURL())
//line admin/domain/view/view.qtpl:8
	qw422016.N().S(`/Domain">`)
//line admin/domain/view/view.qtpl:8
	qw422016.E().S(page.Translate("backToOverview"))
//line admin/domain/view/view.qtpl:8
	qw422016.N().S(`</a><br><a href="`)
//line admin/domain/view/view.qtpl:9
	qw422016.E().S(page.AdminURL())
//line admin/domain/view/view.qtpl:9
	qw422016.N().S(`/Domain/Update?id=`)
//line admin/domain/view/view.qtpl:9
	qw422016.E().S(page.domain.ID())
//line admin/domain/view/view.qtpl:9
	qw422016.N().S(`">`)
//line admin/domain/view/view.qtpl:9
	qw422016.E().S("✎ ")
//line admin/domain/view/view.qtpl:9
	qw422016.E().S(page.Translate("update"))
//line admin/domain/view/view.qtpl:9
	qw422016.N().S(`</a></p><p>`)
//line admin/domain/view/view.qtpl:11
	qw422016.E().S(page.Translate("active"))
//line admin/domain/view/view.qtpl:11
	qw422016.N().S(`:`)
//line admin/domain/view/view.qtpl:11
	qw422016.N().S(` `)
//line admin/domain/view/view.qtpl:11
	if page.domain.Active() {
//line admin/domain/view/view.qtpl:11
		qw422016.E().S(page.Translate("yes_"))
//line admin/domain/view/view.qtpl:11
	} else {
//line admin/domain/view/view.qtpl:11
		qw422016.E().S(page.Translate("no_"))
//line admin/domain/view/view.qtpl:11
	}
//line admin/domain/view/view.qtpl:11
	qw422016.N().S(`</p><p>`)
//line admin/domain/view/view.qtpl:12
	qw422016.E().S(page.Translate("language"))
//line admin/domain/view/view.qtpl:12
	qw422016.N().S(`:`)
//line admin/domain/view/view.qtpl:12
	qw422016.N().S(` `)
//line admin/domain/view/view.qtpl:12
	if page.domainLanguage != nil {
//line admin/domain/view/view.qtpl:12
		qw422016.E().S(page.domainLanguage.Translate(page.language.ID()))
//line admin/domain/view/view.qtpl:12
	}
//line admin/domain/view/view.qtpl:12
	qw422016.N().S(`</p></main>`)
//line admin/domain/view/view.qtpl:14
}

//line admin/domain/view/view.qtpl:14
func (page *Page) WriteContent(qq422016 qtio422016.Writer) {
//line admin/domain/view/view.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/domain/view/view.qtpl:14
	page.StreamContent(qw422016)
//line admin/domain/view/view.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line admin/domain/view/view.qtpl:14
}

//line admin/domain/view/view.qtpl:14
func (page *Page) Content() string {
//line admin/domain/view/view.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/domain/view/view.qtpl:14
	page.WriteContent(qb422016)
//line admin/domain/view/view.qtpl:14
	qs422016 := string(qb422016.B)
//line admin/domain/view/view.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/domain/view/view.qtpl:14
	return qs422016
//line admin/domain/view/view.qtpl:14
}
