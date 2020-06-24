// Code generated by qtc from "view.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line admin/user/group/view/view.qtpl:1
package view

//line admin/user/group/view/view.qtpl:1
import "github.com/espal-digital-development/espal-core/text"

//line admin/user/group/view/view.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line admin/user/group/view/view.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line admin/user/group/view/view.qtpl:3
func (page *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line admin/user/group/view/view.qtpl:3
	qw422016.E().S(page.TranslatePlural("userGroup"))
//line admin/user/group/view/view.qtpl:3
	qw422016.N().S(` > `)
//line admin/user/group/view/view.qtpl:3
	qw422016.E().S(page.userGroup.ID())
//line admin/user/group/view/view.qtpl:3
}

//line admin/user/group/view/view.qtpl:3
func (page *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line admin/user/group/view/view.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/group/view/view.qtpl:3
	page.StreamTitle(qw422016)
//line admin/user/group/view/view.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line admin/user/group/view/view.qtpl:3
}

//line admin/user/group/view/view.qtpl:3
func (page *Page) Title() string {
//line admin/user/group/view/view.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/group/view/view.qtpl:3
	page.WriteTitle(qb422016)
//line admin/user/group/view/view.qtpl:3
	qs422016 := string(qb422016.B)
//line admin/user/group/view/view.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/group/view/view.qtpl:3
	return qs422016
//line admin/user/group/view/view.qtpl:3
}

//line admin/user/group/view/view.qtpl:5
func (page *Page) StreamContent(qw422016 *qt422016.Writer) {
//line admin/user/group/view/view.qtpl:5
	qw422016.N().S(`<main class="content"><h1>`)
//line admin/user/group/view/view.qtpl:7
	qw422016.E().S(page.TranslatePlural("userGroup"))
//line admin/user/group/view/view.qtpl:7
	qw422016.E().S(" > ")
//line admin/user/group/view/view.qtpl:7
	qw422016.E().S(page.userGroup.ID())
//line admin/user/group/view/view.qtpl:7
	qw422016.N().S(`</h1><p><a href="`)
//line admin/user/group/view/view.qtpl:10
	qw422016.E().S(page.AdminURL())
//line admin/user/group/view/view.qtpl:10
	qw422016.N().S(`/UserGroup">`)
//line admin/user/group/view/view.qtpl:10
	qw422016.E().S(page.Translate("backToOverview"))
//line admin/user/group/view/view.qtpl:10
	qw422016.N().S(`</a><br><a href="`)
//line admin/user/group/view/view.qtpl:11
	qw422016.E().S(page.AdminURL())
//line admin/user/group/view/view.qtpl:11
	qw422016.N().S(`/UserGroup/Update?id=`)
//line admin/user/group/view/view.qtpl:11
	qw422016.E().S(page.userGroup.ID())
//line admin/user/group/view/view.qtpl:11
	qw422016.N().S(`">`)
//line admin/user/group/view/view.qtpl:11
	qw422016.E().S("✎ ")
//line admin/user/group/view/view.qtpl:11
	qw422016.E().S(page.Translate("update"))
//line admin/user/group/view/view.qtpl:11
	qw422016.N().S(`</a></p><p>`)
//line admin/user/group/view/view.qtpl:13
	qw422016.E().S(page.Translate("active"))
//line admin/user/group/view/view.qtpl:13
	qw422016.E().S(": ")
//line admin/user/group/view/view.qtpl:13
	if page.userGroup.Active() {
//line admin/user/group/view/view.qtpl:13
		qw422016.E().S(page.Translate("yes_"))
//line admin/user/group/view/view.qtpl:13
	} else {
//line admin/user/group/view/view.qtpl:13
		qw422016.E().S(page.Translate("no_"))
//line admin/user/group/view/view.qtpl:13
	}
//line admin/user/group/view/view.qtpl:13
	qw422016.N().S(`</p><hr><h2>`)
//line admin/user/group/view/view.qtpl:17
	qw422016.E().S(page.TranslatePlural("userRight"))
//line admin/user/group/view/view.qtpl:17
	qw422016.N().S(`</h2>`)
//line admin/user/group/view/view.qtpl:18
	qw422016.N().S(page.userRightsActions.RenderOverviewActions())
//line admin/user/group/view/view.qtpl:19
	qw422016.N().S(`<div style="height: 270px; overflow-y: scroll;"><table><tr><th>`)
//line admin/user/group/view/view.qtpl:23
	qw422016.E().S(page.Translate("category"))
//line admin/user/group/view/view.qtpl:23
	qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:24
	qw422016.E().S(page.Translate("access"))
//line admin/user/group/view/view.qtpl:24
	qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:25
	qw422016.E().S(page.Translate("read"))
//line admin/user/group/view/view.qtpl:25
	qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:26
	qw422016.E().S(page.Translate("create"))
//line admin/user/group/view/view.qtpl:26
	qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:27
	qw422016.E().S(page.Translate("update"))
//line admin/user/group/view/view.qtpl:27
	qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:28
	qw422016.E().S(page.Translate("delete"))
//line admin/user/group/view/view.qtpl:28
	qw422016.N().S(`</th></tr>`)
//line admin/user/group/view/view.qtpl:30
	for k := range page.userRightsOrder {
//line admin/user/group/view/view.qtpl:31
		if page.userRightsOrder[k] >= 3e4 {
//line admin/user/group/view/view.qtpl:32
			if page.userRightsOrder[k] == 30000 || page.userRightsOrder[k] == 30001 {
//line admin/user/group/view/view.qtpl:32
				qw422016.N().S(`<tr><td>`)
//line admin/user/group/view/view.qtpl:34
				qw422016.E().S(page.TranslatePlural(text.LowerFirst(page.userRights[page.userRightsOrder[k]][6:])))
//line admin/user/group/view/view.qtpl:34
				qw422016.N().S(`</td><td colspan="5"><input type="checkbox" name="check" data-id="`)
//line admin/user/group/view/view.qtpl:35
				qw422016.E().V(page.userRightsOrder[k])
//line admin/user/group/view/view.qtpl:35
				qw422016.N().S(`"`)
//line admin/user/group/view/view.qtpl:35
				if page.userGroupUserRights[page.userRightsOrder[k]] {
//line admin/user/group/view/view.qtpl:35
					qw422016.N().S(`checked`)
//line admin/user/group/view/view.qtpl:35
				}
//line admin/user/group/view/view.qtpl:35
				qw422016.N().S(`></td></tr>`)
//line admin/user/group/view/view.qtpl:37
			}
//line admin/user/group/view/view.qtpl:38
		} else {
//line admin/user/group/view/view.qtpl:39
			if page.userRightsOrder[k]%5 == 1 {
//line admin/user/group/view/view.qtpl:39
				qw422016.N().S(`<tr><td>`)
//line admin/user/group/view/view.qtpl:41
				qw422016.E().S(page.TranslatePlural(text.LowerFirst(page.userRights[page.userRightsOrder[k]][6:])))
//line admin/user/group/view/view.qtpl:41
				qw422016.N().S(`</td>`)
//line admin/user/group/view/view.qtpl:42
			}
//line admin/user/group/view/view.qtpl:42
			qw422016.N().S(`<td><input type="checkbox" name="check" data-id="`)
//line admin/user/group/view/view.qtpl:43
			qw422016.E().V(page.userRightsOrder[k])
//line admin/user/group/view/view.qtpl:43
			qw422016.N().S(`"`)
//line admin/user/group/view/view.qtpl:43
			if page.userGroupUserRights[page.userRightsOrder[k]] {
//line admin/user/group/view/view.qtpl:43
				qw422016.N().S(`checked`)
//line admin/user/group/view/view.qtpl:43
			}
//line admin/user/group/view/view.qtpl:43
			qw422016.N().S(`></td>`)
//line admin/user/group/view/view.qtpl:44
			if page.userRightsOrder[k]%5 == 0 {
//line admin/user/group/view/view.qtpl:44
				qw422016.N().S(`<tr>`)
//line admin/user/group/view/view.qtpl:46
			}
//line admin/user/group/view/view.qtpl:47
		}
//line admin/user/group/view/view.qtpl:48
	}
//line admin/user/group/view/view.qtpl:48
	qw422016.N().S(`</table></div><hr><h2>`)
//line admin/user/group/view/view.qtpl:53
	qw422016.E().S(page.TranslatePlural("translation"))
//line admin/user/group/view/view.qtpl:53
	qw422016.N().S(`</h2>`)
//line admin/user/group/view/view.qtpl:54
	qw422016.N().S(page.translationsActions.RenderOverviewActions())
//line admin/user/group/view/view.qtpl:55
	if len(page.translations) > 0 {
//line admin/user/group/view/view.qtpl:55
		qw422016.N().S(`<table><tr>`)
//line admin/user/group/view/view.qtpl:58
		if page.canUpdate || page.canDelete {
//line admin/user/group/view/view.qtpl:58
			qw422016.N().S(`<th><input type="checkbox" name="check"></th><th></th>`)
//line admin/user/group/view/view.qtpl:61
		}
//line admin/user/group/view/view.qtpl:61
		qw422016.N().S(`<th>`)
//line admin/user/group/view/view.qtpl:62
		qw422016.E().S(page.Translate("language"))
//line admin/user/group/view/view.qtpl:62
		qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:63
		qw422016.E().S(page.Translate("field"))
//line admin/user/group/view/view.qtpl:63
		qw422016.N().S(`</th><th>`)
//line admin/user/group/view/view.qtpl:64
		qw422016.E().S(page.Translate("translation"))
//line admin/user/group/view/view.qtpl:64
		qw422016.N().S(`</th></tr>`)
//line admin/user/group/view/view.qtpl:66
		for k := range page.translations {
//line admin/user/group/view/view.qtpl:66
			qw422016.N().S(`<tr>`)
//line admin/user/group/view/view.qtpl:68
			if page.canUpdate || page.canDelete {
//line admin/user/group/view/view.qtpl:68
				qw422016.N().S(`<td><input type="checkbox" name="check" data-id="`)
//line admin/user/group/view/view.qtpl:69
				qw422016.E().S(page.translations[k].ID())
//line admin/user/group/view/view.qtpl:69
				qw422016.N().S(`"></td><td><a href="`)
//line admin/user/group/view/view.qtpl:70
				qw422016.E().S(page.AdminURL())
//line admin/user/group/view/view.qtpl:70
				qw422016.N().S(`/UserGroup/Translation/Update?id=`)
//line admin/user/group/view/view.qtpl:70
				qw422016.E().S(page.translations[k].GroupID())
//line admin/user/group/view/view.qtpl:70
				qw422016.N().S(`&transid=`)
//line admin/user/group/view/view.qtpl:70
				qw422016.E().S(page.translations[k].ID())
//line admin/user/group/view/view.qtpl:70
				qw422016.N().S(`">✎</a></td>`)
//line admin/user/group/view/view.qtpl:71
			}
//line admin/user/group/view/view.qtpl:71
			qw422016.N().S(`<td>`)
//line admin/user/group/view/view.qtpl:73
			qw422016.E().S(page.rendererService.LanguageName(page.translations[k].Language(), page.language.ID()))
//line admin/user/group/view/view.qtpl:73
			qw422016.N().S(`</td><td>`)
//line admin/user/group/view/view.qtpl:75
			if page.translations[k].Field() == 1 {
//line admin/user/group/view/view.qtpl:75
				qw422016.E().S(page.Translate("name"))
//line admin/user/group/view/view.qtpl:75
			} else {
//line admin/user/group/view/view.qtpl:75
				qw422016.E().S(page.Translate("description"))
//line admin/user/group/view/view.qtpl:75
			}
//line admin/user/group/view/view.qtpl:75
			qw422016.N().S(`</td><td>`)
//line admin/user/group/view/view.qtpl:76
			qw422016.E().S(page.translations[k].Value())
//line admin/user/group/view/view.qtpl:76
			qw422016.N().S(`</td></tr>`)
//line admin/user/group/view/view.qtpl:78
		}
//line admin/user/group/view/view.qtpl:78
		qw422016.N().S(`</table>`)
//line admin/user/group/view/view.qtpl:80
	} else {
//line admin/user/group/view/view.qtpl:81
		qw422016.E().S(page.Translate("noResultsFound"))
//line admin/user/group/view/view.qtpl:82
	}
//line admin/user/group/view/view.qtpl:82
	qw422016.N().S(`</main>`)
//line admin/user/group/view/view.qtpl:84
}

//line admin/user/group/view/view.qtpl:84
func (page *Page) WriteContent(qq422016 qtio422016.Writer) {
//line admin/user/group/view/view.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/group/view/view.qtpl:84
	page.StreamContent(qw422016)
//line admin/user/group/view/view.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line admin/user/group/view/view.qtpl:84
}

//line admin/user/group/view/view.qtpl:84
func (page *Page) Content() string {
//line admin/user/group/view/view.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/group/view/view.qtpl:84
	page.WriteContent(qb422016)
//line admin/user/group/view/view.qtpl:84
	qs422016 := string(qb422016.B)
//line admin/user/group/view/view.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/group/view/view.qtpl:84
	return qs422016
//line admin/user/group/view/view.qtpl:84
}

//line admin/user/group/view/view.qtpl:86
func (page *Page) StreamJavascripts(qw422016 *qt422016.Writer) {
//line admin/user/group/view/view.qtpl:87
	if len(page.translations) > 0 {
//line admin/user/group/view/view.qtpl:87
		qw422016.N().S(`<script src="/j/a/o.js"></script>`)
//line admin/user/group/view/view.qtpl:89
		qw422016.N().S(page.rendererService.DefaultOverviewTranslations(page.GetCoreContext()))
//line admin/user/group/view/view.qtpl:90
	}
//line admin/user/group/view/view.qtpl:91
}

//line admin/user/group/view/view.qtpl:91
func (page *Page) WriteJavascripts(qq422016 qtio422016.Writer) {
//line admin/user/group/view/view.qtpl:91
	qw422016 := qt422016.AcquireWriter(qq422016)
//line admin/user/group/view/view.qtpl:91
	page.StreamJavascripts(qw422016)
//line admin/user/group/view/view.qtpl:91
	qt422016.ReleaseWriter(qw422016)
//line admin/user/group/view/view.qtpl:91
}

//line admin/user/group/view/view.qtpl:91
func (page *Page) Javascripts() string {
//line admin/user/group/view/view.qtpl:91
	qb422016 := qt422016.AcquireByteBuffer()
//line admin/user/group/view/view.qtpl:91
	page.WriteJavascripts(qb422016)
//line admin/user/group/view/view.qtpl:91
	qs422016 := string(qb422016.B)
//line admin/user/group/view/view.qtpl:91
	qt422016.ReleaseByteBuffer(qb422016)
//line admin/user/group/view/view.qtpl:91
	return qs422016
//line admin/user/group/view/view.qtpl:91
}
