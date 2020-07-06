// Code generated by qtc from "overview.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pages/pprof/overview/overview.qtpl:1
package overview

//line pages/pprof/overview/overview.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pages/pprof/overview/overview.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pages/pprof/overview/overview.qtpl:1
func (p *Page) StreamStylesheets(qw422016 *qt422016.Writer) {
//line pages/pprof/overview/overview.qtpl:1
	qw422016.N().S(`<link rel="stylesheet" href="/c/simpleBox.css">`)
//line pages/pprof/overview/overview.qtpl:1
}

//line pages/pprof/overview/overview.qtpl:1
func (p *Page) WriteStylesheets(qq422016 qtio422016.Writer) {
//line pages/pprof/overview/overview.qtpl:1
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/pprof/overview/overview.qtpl:1
	p.StreamStylesheets(qw422016)
//line pages/pprof/overview/overview.qtpl:1
	qt422016.ReleaseWriter(qw422016)
//line pages/pprof/overview/overview.qtpl:1
}

//line pages/pprof/overview/overview.qtpl:1
func (p *Page) Stylesheets() string {
//line pages/pprof/overview/overview.qtpl:1
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/pprof/overview/overview.qtpl:1
	p.WriteStylesheets(qb422016)
//line pages/pprof/overview/overview.qtpl:1
	qs422016 := string(qb422016.B)
//line pages/pprof/overview/overview.qtpl:1
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/pprof/overview/overview.qtpl:1
	return qs422016
//line pages/pprof/overview/overview.qtpl:1
}

//line pages/pprof/overview/overview.qtpl:3
func (p *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line pages/pprof/overview/overview.qtpl:3
	qw422016.N().S(`Pprof`)
//line pages/pprof/overview/overview.qtpl:3
}

//line pages/pprof/overview/overview.qtpl:3
func (p *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line pages/pprof/overview/overview.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/pprof/overview/overview.qtpl:3
	p.StreamTitle(qw422016)
//line pages/pprof/overview/overview.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line pages/pprof/overview/overview.qtpl:3
}

//line pages/pprof/overview/overview.qtpl:3
func (p *Page) Title() string {
//line pages/pprof/overview/overview.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/pprof/overview/overview.qtpl:3
	p.WriteTitle(qb422016)
//line pages/pprof/overview/overview.qtpl:3
	qs422016 := string(qb422016.B)
//line pages/pprof/overview/overview.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/pprof/overview/overview.qtpl:3
	return qs422016
//line pages/pprof/overview/overview.qtpl:3
}

//line pages/pprof/overview/overview.qtpl:5
func (p *Page) StreamContent(qw422016 *qt422016.Writer) {
//line pages/pprof/overview/overview.qtpl:5
	qw422016.N().S(`<div class="simpleBox"><h1>Pprof</h1><ul style="text-align: left; text-transform: capitalize;">`)
//line pages/pprof/overview/overview.qtpl:10
	for k := range p.profiles {
//line pages/pprof/overview/overview.qtpl:10
		qw422016.N().S(`<li><a href="`)
//line pages/pprof/overview/overview.qtpl:11
		qw422016.E().S(p.PprofURL())
//line pages/pprof/overview/overview.qtpl:11
		qw422016.N().S(`/`)
//line pages/pprof/overview/overview.qtpl:11
		qw422016.E().S(p.profiles[k].Name())
//line pages/pprof/overview/overview.qtpl:11
		qw422016.N().S(`?debug=1">`)
//line pages/pprof/overview/overview.qtpl:11
		qw422016.E().S(p.profiles[k].Name())
//line pages/pprof/overview/overview.qtpl:11
		qw422016.N().S(` `)
//line pages/pprof/overview/overview.qtpl:11
		qw422016.N().S(`(`)
//line pages/pprof/overview/overview.qtpl:11
		qw422016.N().D(p.profiles[k].Count())
//line pages/pprof/overview/overview.qtpl:11
		qw422016.N().S(`)</a></li>`)
//line pages/pprof/overview/overview.qtpl:12
	}
//line pages/pprof/overview/overview.qtpl:12
	qw422016.N().S(`<li><a href="`)
//line pages/pprof/overview/overview.qtpl:13
	qw422016.E().S(p.PprofURL())
//line pages/pprof/overview/overview.qtpl:13
	qw422016.N().S(`/goroutine?debug=2">Full Goroutine Stack Dump</a></li></ul></div>`)
//line pages/pprof/overview/overview.qtpl:16
}

//line pages/pprof/overview/overview.qtpl:16
func (p *Page) WriteContent(qq422016 qtio422016.Writer) {
//line pages/pprof/overview/overview.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/pprof/overview/overview.qtpl:16
	p.StreamContent(qw422016)
//line pages/pprof/overview/overview.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line pages/pprof/overview/overview.qtpl:16
}

//line pages/pprof/overview/overview.qtpl:16
func (p *Page) Content() string {
//line pages/pprof/overview/overview.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/pprof/overview/overview.qtpl:16
	p.WriteContent(qb422016)
//line pages/pprof/overview/overview.qtpl:16
	qs422016 := string(qb422016.B)
//line pages/pprof/overview/overview.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/pprof/overview/overview.qtpl:16
	return qs422016
//line pages/pprof/overview/overview.qtpl:16
}
