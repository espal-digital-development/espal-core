// Code generated by qtc from "createupdate.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line admin/user/address/createupdate/createupdate.qtpl:1
package createupdate

//line admin/user/address/createupdate/createupdate.qtpl:1
import "fmt"

//line admin/user/address/createupdate/createupdate.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin/user/address/createupdate/createupdate.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin/user/address/createupdate/createupdate.qtpl:3
func (page *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:3
	qw422016.E().S(page.displayTitle)
//line admin/user/address/createupdate/createupdate.qtpl:3
}

//line admin/user/address/createupdate/createupdate.qtpl:3
func (page *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/address/createupdate/createupdate.qtpl:3
	page.StreamTitle(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:3
}

//line admin/user/address/createupdate/createupdate.qtpl:3
func (page *Page) Title() string {
//line admin/user/address/createupdate/createupdate.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/address/createupdate/createupdate.qtpl:3
	page.WriteTitle(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:3
	qs422016 := string(qb422016.B)
//line admin/user/address/createupdate/createupdate.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:3
	return qs422016
//line admin/user/address/createupdate/createupdate.qtpl:3
}

//line admin/user/address/createupdate/createupdate.qtpl:5
func (page *Page) StreamContent(qw422016 *qt422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:5
	qw422016.N().S(`<main class="content"><h1>`)
//line admin/user/address/createupdate/createupdate.qtpl:7
	qw422016.E().S(page.displayTitle)
//line admin/user/address/createupdate/createupdate.qtpl:7
	qw422016.N().S(`</h1>`)
//line admin/user/address/createupdate/createupdate.qtpl:8
	qw422016.N().S(page.form.Errors())
//line admin/user/address/createupdate/createupdate.qtpl:9
	qw422016.N().S(page.rendererService.CreatedUpdatedByLinks(page.GetCoreContext(), page.language.ID(), page.userAddress))
//line admin/user/address/createupdate/createupdate.qtpl:10
	qw422016.N().S(page.form.Open())
//line admin/user/address/createupdate/createupdate.qtpl:11
	qw422016.N().S(page.form.Field("_uname"))
//line admin/user/address/createupdate/createupdate.qtpl:12
	qw422016.N().S(page.form.Field("_t"))
//line admin/user/address/createupdate/createupdate.qtpl:13
	qw422016.N().S(page.form.Field("active"))
//line admin/user/address/createupdate/createupdate.qtpl:13
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:14
	qw422016.N().S(page.form.Field("firstName"))
//line admin/user/address/createupdate/createupdate.qtpl:14
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:15
	qw422016.N().S(page.form.Field("surname"))
//line admin/user/address/createupdate/createupdate.qtpl:15
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:16
	qw422016.N().S(page.form.Field("street"))
//line admin/user/address/createupdate/createupdate.qtpl:16
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:17
	qw422016.N().S(page.form.Field("streetLine2"))
//line admin/user/address/createupdate/createupdate.qtpl:17
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:18
	qw422016.N().S(page.form.Field("number"))
//line admin/user/address/createupdate/createupdate.qtpl:18
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:19
	qw422016.N().S(page.form.Field("numberAddition"))
//line admin/user/address/createupdate/createupdate.qtpl:19
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:20
	qw422016.N().S(page.form.Field("zipCode"))
//line admin/user/address/createupdate/createupdate.qtpl:20
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:21
	qw422016.N().S(page.form.Field("city"))
//line admin/user/address/createupdate/createupdate.qtpl:21
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:22
	qw422016.N().S(page.form.Field("country"))
//line admin/user/address/createupdate/createupdate.qtpl:22
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:23
	qw422016.N().S(page.form.Field("phoneNumber"))
//line admin/user/address/createupdate/createupdate.qtpl:23
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:24
	qw422016.N().S(page.form.Field("email"))
//line admin/user/address/createupdate/createupdate.qtpl:24
	qw422016.N().S(`<br>`)
//line admin/user/address/createupdate/createupdate.qtpl:25
	qw422016.N().S(page.form.CreateUpdateActions("action", fmt.Sprintf("User/View?id=%s", page.user.ID())))
//line admin/user/address/createupdate/createupdate.qtpl:25
	qw422016.N().S(`</form></main>`)
//line admin/user/address/createupdate/createupdate.qtpl:28
}

//line admin/user/address/createupdate/createupdate.qtpl:28
func (page *Page) WriteContent(qq422016 qtio422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/address/createupdate/createupdate.qtpl:28
	page.StreamContent(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:28
}

//line admin/user/address/createupdate/createupdate.qtpl:28
func (page *Page) Content() string {
//line admin/user/address/createupdate/createupdate.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/address/createupdate/createupdate.qtpl:28
	page.WriteContent(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:28
	qs422016 := string(qb422016.B)
//line admin/user/address/createupdate/createupdate.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:28
	return qs422016
//line admin/user/address/createupdate/createupdate.qtpl:28
}

//line admin/user/address/createupdate/createupdate.qtpl:30
func (page *Page) StreamStylesheets(qw422016 *qt422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:30
	qw422016.N().S(`<link rel="stylesheet" href="/c/a/selectSearch.css"></link>`)
//line admin/user/address/createupdate/createupdate.qtpl:32
}

//line admin/user/address/createupdate/createupdate.qtpl:32
func (page *Page) WriteStylesheets(qq422016 qtio422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/address/createupdate/createupdate.qtpl:32
	page.StreamStylesheets(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:32
}

//line admin/user/address/createupdate/createupdate.qtpl:32
func (page *Page) Stylesheets() string {
//line admin/user/address/createupdate/createupdate.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/address/createupdate/createupdate.qtpl:32
	page.WriteStylesheets(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:32
	qs422016 := string(qb422016.B)
//line admin/user/address/createupdate/createupdate.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:32
	return qs422016
//line admin/user/address/createupdate/createupdate.qtpl:32
}

//line admin/user/address/createupdate/createupdate.qtpl:34
func (page *Page) StreamJavascripts(qw422016 *qt422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:34
	qw422016.N().S(`<script src="/j/a/cu.js"></script><script src="/j/a/selectSearch.js"></script>`)
//line admin/user/address/createupdate/createupdate.qtpl:37
}

//line admin/user/address/createupdate/createupdate.qtpl:37
func (page *Page) WriteJavascripts(qq422016 qtio422016.Writer) {
//line admin/user/address/createupdate/createupdate.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/address/createupdate/createupdate.qtpl:37
	page.StreamJavascripts(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line admin/user/address/createupdate/createupdate.qtpl:37
}

//line admin/user/address/createupdate/createupdate.qtpl:37
func (page *Page) Javascripts() string {
//line admin/user/address/createupdate/createupdate.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/address/createupdate/createupdate.qtpl:37
	page.WriteJavascripts(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:37
	qs422016 := string(qb422016.B)
//line admin/user/address/createupdate/createupdate.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/address/createupdate/createupdate.qtpl:37
	return qs422016
//line admin/user/address/createupdate/createupdate.qtpl:37
}
