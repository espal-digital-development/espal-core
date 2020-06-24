// Code generated by qtc from "forgot.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line account/password/forgot/forgot.qtpl:1
package forgot

//line account/password/forgot/forgot.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line account/password/forgot/forgot.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line account/password/forgot/forgot.qtpl:1
func (page *Page) StreamStylesheets(qw422016 *qt422016.Writer) {
//line account/password/forgot/forgot.qtpl:1
	qw422016.N().S(`<link rel="stylesheet" href="/c/simpleBox.css">`)
//line account/password/forgot/forgot.qtpl:1
}

//line account/password/forgot/forgot.qtpl:1
func (page *Page) WriteStylesheets(qq422016 qtio422016.Writer) {
//line account/password/forgot/forgot.qtpl:1
	qw422016 := qt422016.AcquireWriter(qq422016)
//line account/password/forgot/forgot.qtpl:1
	page.StreamStylesheets(qw422016)
//line account/password/forgot/forgot.qtpl:1
	qt422016.ReleaseWriter(qw422016)
//line account/password/forgot/forgot.qtpl:1
}

//line account/password/forgot/forgot.qtpl:1
func (page *Page) Stylesheets() string {
//line account/password/forgot/forgot.qtpl:1
	qb422016 := qt422016.AcquireByteBuffer()
//line account/password/forgot/forgot.qtpl:1
	page.WriteStylesheets(qb422016)
//line account/password/forgot/forgot.qtpl:1
	qs422016 := string(qb422016.B)
//line account/password/forgot/forgot.qtpl:1
	qt422016.ReleaseByteBuffer(qb422016)
//line account/password/forgot/forgot.qtpl:1
	return qs422016
//line account/password/forgot/forgot.qtpl:1
}

//line account/password/forgot/forgot.qtpl:3
func (page *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line account/password/forgot/forgot.qtpl:3
	qw422016.E().S(page.Translate("forgotPassword"))
//line account/password/forgot/forgot.qtpl:3
}

//line account/password/forgot/forgot.qtpl:3
func (page *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line account/password/forgot/forgot.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line account/password/forgot/forgot.qtpl:3
	page.StreamTitle(qw422016)
//line account/password/forgot/forgot.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line account/password/forgot/forgot.qtpl:3
}

//line account/password/forgot/forgot.qtpl:3
func (page *Page) Title() string {
//line account/password/forgot/forgot.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line account/password/forgot/forgot.qtpl:3
	page.WriteTitle(qb422016)
//line account/password/forgot/forgot.qtpl:3
	qs422016 := string(qb422016.B)
//line account/password/forgot/forgot.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line account/password/forgot/forgot.qtpl:3
	return qs422016
//line account/password/forgot/forgot.qtpl:3
}

//line account/password/forgot/forgot.qtpl:5
func (page *Page) StreamContent(qw422016 *qt422016.Writer) {
//line account/password/forgot/forgot.qtpl:5
	qw422016.N().S(`<div class="simpleBox">`)
//line account/password/forgot/forgot.qtpl:7
	qw422016.N().S(page.form.Errors())
//line account/password/forgot/forgot.qtpl:7
	qw422016.N().S(`<h1>`)
//line account/password/forgot/forgot.qtpl:8
	qw422016.E().S(page.Translate("forgotPassword"))
//line account/password/forgot/forgot.qtpl:8
	qw422016.N().S(`</h1>`)
//line account/password/forgot/forgot.qtpl:9
	qw422016.N().S(page.form.Open())
//line account/password/forgot/forgot.qtpl:10
	qw422016.N().S(page.form.Field("_uname"))
//line account/password/forgot/forgot.qtpl:11
	qw422016.N().S(page.form.Field("_t"))
//line account/password/forgot/forgot.qtpl:12
	qw422016.N().S(page.form.Field("email"))
//line account/password/forgot/forgot.qtpl:12
	qw422016.N().S(`<br>`)
//line account/password/forgot/forgot.qtpl:13
	qw422016.N().S(page.form.Field("repeatEmail"))
//line account/password/forgot/forgot.qtpl:13
	qw422016.N().S(`<br><input type="submit" value="`)
//line account/password/forgot/forgot.qtpl:14
	qw422016.E().S(page.Translate("mailMeTheLink"))
//line account/password/forgot/forgot.qtpl:14
	qw422016.N().S(`"></form></div>`)
//line account/password/forgot/forgot.qtpl:17
}

//line account/password/forgot/forgot.qtpl:17
func (page *Page) WriteContent(qq422016 qtio422016.Writer) {
//line account/password/forgot/forgot.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line account/password/forgot/forgot.qtpl:17
	page.StreamContent(qw422016)
//line account/password/forgot/forgot.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line account/password/forgot/forgot.qtpl:17
}

//line account/password/forgot/forgot.qtpl:17
func (page *Page) Content() string {
//line account/password/forgot/forgot.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line account/password/forgot/forgot.qtpl:17
	page.WriteContent(qb422016)
//line account/password/forgot/forgot.qtpl:17
	qs422016 := string(qb422016.B)
//line account/password/forgot/forgot.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line account/password/forgot/forgot.qtpl:17
	return qs422016
//line account/password/forgot/forgot.qtpl:17
}

//line account/password/forgot/forgot.qtpl:19
func (page *Page) StreamJavascripts(qw422016 *qt422016.Writer) {
//line account/password/forgot/forgot.qtpl:19
	qw422016.N().S(`<script src="/j/form.js"></script>`)
//line account/password/forgot/forgot.qtpl:19
}

//line account/password/forgot/forgot.qtpl:19
func (page *Page) WriteJavascripts(qq422016 qtio422016.Writer) {
//line account/password/forgot/forgot.qtpl:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line account/password/forgot/forgot.qtpl:19
	page.StreamJavascripts(qw422016)
//line account/password/forgot/forgot.qtpl:19
	qt422016.ReleaseWriter(qw422016)
//line account/password/forgot/forgot.qtpl:19
}

//line account/password/forgot/forgot.qtpl:19
func (page *Page) Javascripts() string {
//line account/password/forgot/forgot.qtpl:19
	qb422016 := qt422016.AcquireByteBuffer()
//line account/password/forgot/forgot.qtpl:19
	page.WriteJavascripts(qb422016)
//line account/password/forgot/forgot.qtpl:19
	qs422016 := string(qb422016.B)
//line account/password/forgot/forgot.qtpl:19
	qt422016.ReleaseByteBuffer(qb422016)
//line account/password/forgot/forgot.qtpl:19
	return qs422016
//line account/password/forgot/forgot.qtpl:19
}
