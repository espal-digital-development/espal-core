// Code generated by qtc from "view.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pages/admin/user/view/view.qtpl:1
package view

//line pages/admin/user/view/view.qtpl:1
import "time"

//line pages/admin/user/view/view.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pages/admin/user/view/view.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pages/admin/user/view/view.qtpl:3
func (p *Page) StreamTitle(qw422016 *qt422016.Writer) {
//line pages/admin/user/view/view.qtpl:3
	qw422016.E().S(p.TranslatePlural("user"))
//line pages/admin/user/view/view.qtpl:3
	qw422016.N().S(` > `)
//line pages/admin/user/view/view.qtpl:3
	qw422016.E().S(p.userStore.Name(p.user, p.language.ID()))
//line pages/admin/user/view/view.qtpl:3
}

//line pages/admin/user/view/view.qtpl:3
func (p *Page) WriteTitle(qq422016 qtio422016.Writer) {
//line pages/admin/user/view/view.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/user/view/view.qtpl:3
	p.StreamTitle(qw422016)
//line pages/admin/user/view/view.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/user/view/view.qtpl:3
}

//line pages/admin/user/view/view.qtpl:3
func (p *Page) Title() string {
//line pages/admin/user/view/view.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/user/view/view.qtpl:3
	p.WriteTitle(qb422016)
//line pages/admin/user/view/view.qtpl:3
	qs422016 := string(qb422016.B)
//line pages/admin/user/view/view.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/user/view/view.qtpl:3
	return qs422016
//line pages/admin/user/view/view.qtpl:3
}

//line pages/admin/user/view/view.qtpl:5
func (p *Page) StreamContent(qw422016 *qt422016.Writer) {
//line pages/admin/user/view/view.qtpl:5
	qw422016.N().S(`<main class="content"><h1>`)
//line pages/admin/user/view/view.qtpl:7
	qw422016.E().S(p.TranslatePlural("user"))
//line pages/admin/user/view/view.qtpl:7
	qw422016.E().S(" > ")
//line pages/admin/user/view/view.qtpl:7
	qw422016.E().S(p.userStore.Name(p.user, p.language.ID()))
//line pages/admin/user/view/view.qtpl:7
	qw422016.N().S(`</h1><p><a href="`)
//line pages/admin/user/view/view.qtpl:10
	qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:10
	qw422016.N().S(`/User">`)
//line pages/admin/user/view/view.qtpl:10
	qw422016.E().S(p.Translate("backToOverview"))
//line pages/admin/user/view/view.qtpl:10
	qw422016.N().S(`</a><br><a href="`)
//line pages/admin/user/view/view.qtpl:11
	qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:11
	qw422016.N().S(`/User/Update?id=`)
//line pages/admin/user/view/view.qtpl:11
	qw422016.E().S(p.user.ID())
//line pages/admin/user/view/view.qtpl:11
	qw422016.N().S(`">`)
//line pages/admin/user/view/view.qtpl:11
	qw422016.E().S("✎ ")
//line pages/admin/user/view/view.qtpl:11
	qw422016.E().S(p.Translate("update"))
//line pages/admin/user/view/view.qtpl:11
	qw422016.N().S(`</a></p><p>`)
//line pages/admin/user/view/view.qtpl:13
	qw422016.E().S(p.Translate("active"))
//line pages/admin/user/view/view.qtpl:13
	qw422016.E().S(": ")
//line pages/admin/user/view/view.qtpl:13
	if p.user.Active() {
//line pages/admin/user/view/view.qtpl:13
		qw422016.E().S(p.Translate("yes_"))
//line pages/admin/user/view/view.qtpl:13
	} else {
//line pages/admin/user/view/view.qtpl:13
		qw422016.E().S(p.Translate("no_"))
//line pages/admin/user/view/view.qtpl:13
	}
//line pages/admin/user/view/view.qtpl:13
	qw422016.N().S(`</p><p>`)
//line pages/admin/user/view/view.qtpl:14
	qw422016.E().S(p.user.Email())
//line pages/admin/user/view/view.qtpl:14
	qw422016.N().S(`</p><hr><h2>`)
//line pages/admin/user/view/view.qtpl:18
	qw422016.E().S(p.TranslatePlural("address"))
//line pages/admin/user/view/view.qtpl:18
	qw422016.N().S(`</h2>`)
//line pages/admin/user/view/view.qtpl:19
	qw422016.N().S(p.addressesActions.RenderOverviewActions())
//line pages/admin/user/view/view.qtpl:20
	if len(p.addresses) > 0 {
//line pages/admin/user/view/view.qtpl:20
		qw422016.N().S(`<table><tr>`)
//line pages/admin/user/view/view.qtpl:23
		if p.HasUserRight("UpdateUserAddress") || p.HasUserRight("DeleteUserAddress") {
//line pages/admin/user/view/view.qtpl:23
			qw422016.N().S(`<th><input type="checkbox" name="check"></th><th></th>`)
//line pages/admin/user/view/view.qtpl:26
		}
//line pages/admin/user/view/view.qtpl:27
		qw422016.N().S(`<th>`)
//line pages/admin/user/view/view.qtpl:28
		qw422016.E().S(p.Translate("id"))
//line pages/admin/user/view/view.qtpl:28
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:29
		qw422016.E().S(p.Translate("active"))
//line pages/admin/user/view/view.qtpl:29
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:30
		qw422016.E().S(p.Translate("createdBy"))
//line pages/admin/user/view/view.qtpl:30
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:31
		qw422016.E().S(p.Translate("createdAt"))
//line pages/admin/user/view/view.qtpl:31
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:32
		qw422016.E().S(p.Translate("updatedBy"))
//line pages/admin/user/view/view.qtpl:32
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:33
		qw422016.E().S(p.Translate("updatedAt"))
//line pages/admin/user/view/view.qtpl:33
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:34
		qw422016.E().S(p.Translate("firstName"))
//line pages/admin/user/view/view.qtpl:34
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:35
		qw422016.E().S(p.Translate("surname"))
//line pages/admin/user/view/view.qtpl:35
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:36
		qw422016.E().S(p.Translate("street"))
//line pages/admin/user/view/view.qtpl:36
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:37
		qw422016.E().S(p.Translate("streetLine2"))
//line pages/admin/user/view/view.qtpl:37
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:38
		qw422016.E().S(p.Translate("number"))
//line pages/admin/user/view/view.qtpl:38
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:39
		qw422016.E().S(p.Translate("numberAddition"))
//line pages/admin/user/view/view.qtpl:39
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:40
		qw422016.E().S(p.Translate("zipCode"))
//line pages/admin/user/view/view.qtpl:40
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:41
		qw422016.E().S(p.Translate("city"))
//line pages/admin/user/view/view.qtpl:41
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:42
		qw422016.E().S(p.Translate("state"))
//line pages/admin/user/view/view.qtpl:42
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:43
		qw422016.E().S(p.Translate("country"))
//line pages/admin/user/view/view.qtpl:43
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:44
		qw422016.E().S(p.Translate("phoneNumber"))
//line pages/admin/user/view/view.qtpl:44
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:45
		qw422016.E().S(p.Translate("email"))
//line pages/admin/user/view/view.qtpl:45
		qw422016.N().S(`</th></tr>`)
//line pages/admin/user/view/view.qtpl:47
		for k := range p.addresses {
//line pages/admin/user/view/view.qtpl:47
			qw422016.N().S(`<tr>`)
//line pages/admin/user/view/view.qtpl:49
			if p.HasUserRight("UpdateUserAddress") || p.HasUserRight("DeleteUserAddress") {
//line pages/admin/user/view/view.qtpl:49
				qw422016.N().S(`<td><input type="checkbox" name="check" data-id="`)
//line pages/admin/user/view/view.qtpl:50
				qw422016.E().S(p.addresses[k].ID())
//line pages/admin/user/view/view.qtpl:50
				qw422016.N().S(`"></td><td><a href="`)
//line pages/admin/user/view/view.qtpl:51
				qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:51
				qw422016.N().S(`/User/Address/Update?userid=`)
//line pages/admin/user/view/view.qtpl:51
				qw422016.E().S(p.addresses[k].ID())
//line pages/admin/user/view/view.qtpl:51
				qw422016.N().S(`&id=`)
//line pages/admin/user/view/view.qtpl:51
				qw422016.E().S(p.addresses[k].ID())
//line pages/admin/user/view/view.qtpl:51
				qw422016.N().S(`">✎</a></td>`)
//line pages/admin/user/view/view.qtpl:52
			}
//line pages/admin/user/view/view.qtpl:52
			qw422016.N().S(`<td>`)
//line pages/admin/user/view/view.qtpl:53
			qw422016.E().S(p.addresses[k].ID())
//line pages/admin/user/view/view.qtpl:53
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:54
			if p.addresses[k].Active() {
//line pages/admin/user/view/view.qtpl:54
				qw422016.E().S(p.Translate("yes_"))
//line pages/admin/user/view/view.qtpl:54
			} else {
//line pages/admin/user/view/view.qtpl:54
				qw422016.E().S(p.Translate("no_"))
//line pages/admin/user/view/view.qtpl:54
			}
//line pages/admin/user/view/view.qtpl:54
			qw422016.N().S(`</td><td><a href="`)
//line pages/admin/user/view/view.qtpl:55
			qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:55
			qw422016.N().S(`/User/View?id=`)
//line pages/admin/user/view/view.qtpl:55
			qw422016.E().S(p.addresses[k].CreatedByID())
//line pages/admin/user/view/view.qtpl:55
			qw422016.N().S(`">`)
//line pages/admin/user/view/view.qtpl:55
			qw422016.E().S(p.rendererService.CreatedBy(p.addresses[k], p.language.ID()))
//line pages/admin/user/view/view.qtpl:55
			qw422016.N().S(`</a></td><td>`)
//line pages/admin/user/view/view.qtpl:56
			qw422016.E().S(p.addresses[k].CreatedAt().Format(time.RFC3339[0:9]))
//line pages/admin/user/view/view.qtpl:56
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:57
			if p.addresses[k].IsUpdated() {
//line pages/admin/user/view/view.qtpl:57
				qw422016.N().S(`<a href="`)
//line pages/admin/user/view/view.qtpl:57
				qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:57
				qw422016.N().S(`/User/View?id=`)
//line pages/admin/user/view/view.qtpl:57
				qw422016.E().S(*p.addresses[k].UpdatedByID())
//line pages/admin/user/view/view.qtpl:57
				qw422016.N().S(`">`)
//line pages/admin/user/view/view.qtpl:57
				qw422016.E().S(p.rendererService.UpdatedBy(p.addresses[k], p.language.ID()))
//line pages/admin/user/view/view.qtpl:57
				qw422016.N().S(`</a>`)
//line pages/admin/user/view/view.qtpl:57
			}
//line pages/admin/user/view/view.qtpl:57
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:58
			if p.addresses[k].UpdatedAt() != nil {
//line pages/admin/user/view/view.qtpl:58
				qw422016.E().S(p.addresses[k].UpdatedAt().Format(time.RFC3339[0:9]))
//line pages/admin/user/view/view.qtpl:58
			}
//line pages/admin/user/view/view.qtpl:58
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:59
			if p.addresses[k].FirstName() != nil {
//line pages/admin/user/view/view.qtpl:59
				qw422016.E().S(*p.addresses[k].FirstName())
//line pages/admin/user/view/view.qtpl:59
			}
//line pages/admin/user/view/view.qtpl:59
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:60
			if p.addresses[k].Surname() != nil {
//line pages/admin/user/view/view.qtpl:60
				qw422016.E().S(*p.addresses[k].Surname())
//line pages/admin/user/view/view.qtpl:60
			}
//line pages/admin/user/view/view.qtpl:60
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:61
			qw422016.E().S(p.addresses[k].Street())
//line pages/admin/user/view/view.qtpl:61
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:62
			if p.addresses[k].StreetLine2() != nil {
//line pages/admin/user/view/view.qtpl:62
				qw422016.E().S(*p.addresses[k].StreetLine2())
//line pages/admin/user/view/view.qtpl:62
			}
//line pages/admin/user/view/view.qtpl:62
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:63
			qw422016.E().S(p.addresses[k].Number())
//line pages/admin/user/view/view.qtpl:63
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:64
			if p.addresses[k].NumberAddition() != nil {
//line pages/admin/user/view/view.qtpl:64
				qw422016.E().S(*p.addresses[k].NumberAddition())
//line pages/admin/user/view/view.qtpl:64
			}
//line pages/admin/user/view/view.qtpl:64
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:65
			qw422016.E().S(p.addresses[k].ZipCode())
//line pages/admin/user/view/view.qtpl:65
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:66
			qw422016.E().S(p.addresses[k].City())
//line pages/admin/user/view/view.qtpl:66
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:67
			if p.addresses[k].State() != nil {
//line pages/admin/user/view/view.qtpl:67
				qw422016.E().S(*p.addresses[k].State())
//line pages/admin/user/view/view.qtpl:67
			}
//line pages/admin/user/view/view.qtpl:67
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:69
			if p.addresses[k].Country() != nil {
//line pages/admin/user/view/view.qtpl:69
				qw422016.E().S(p.rendererService.CountryName(*p.addresses[k].Country(), p.language.ID()))
//line pages/admin/user/view/view.qtpl:69
			}
//line pages/admin/user/view/view.qtpl:69
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:71
			if p.addresses[k].PhoneNumber() != nil {
//line pages/admin/user/view/view.qtpl:71
				qw422016.E().S(*p.addresses[k].PhoneNumber())
//line pages/admin/user/view/view.qtpl:71
			}
//line pages/admin/user/view/view.qtpl:71
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:72
			if p.addresses[k].Email() != nil {
//line pages/admin/user/view/view.qtpl:72
				qw422016.E().S(*p.addresses[k].Email())
//line pages/admin/user/view/view.qtpl:72
			}
//line pages/admin/user/view/view.qtpl:72
			qw422016.N().S(`</td></tr>`)
//line pages/admin/user/view/view.qtpl:74
		}
//line pages/admin/user/view/view.qtpl:74
		qw422016.N().S(`</table>`)
//line pages/admin/user/view/view.qtpl:76
	} else {
//line pages/admin/user/view/view.qtpl:77
		qw422016.E().S(p.Translate("noResultsFound"))
//line pages/admin/user/view/view.qtpl:78
	}
//line pages/admin/user/view/view.qtpl:78
	qw422016.N().S(`<hr><h2>`)
//line pages/admin/user/view/view.qtpl:82
	qw422016.E().S(p.TranslatePlural("contact"))
//line pages/admin/user/view/view.qtpl:82
	qw422016.N().S(`</h2>`)
//line pages/admin/user/view/view.qtpl:83
	qw422016.N().S(p.contactsActions.RenderOverviewActions())
//line pages/admin/user/view/view.qtpl:84
	if len(p.contacts) > 0 {
//line pages/admin/user/view/view.qtpl:84
		qw422016.N().S(`<table><tr>`)
//line pages/admin/user/view/view.qtpl:87
		if p.HasUserRight("UpdateUserContact") || p.HasUserRight("DeleteUserContact") {
//line pages/admin/user/view/view.qtpl:87
			qw422016.N().S(`<th><input type="checkbox" name="check"></th><th></th>`)
//line pages/admin/user/view/view.qtpl:90
		}
//line pages/admin/user/view/view.qtpl:90
		qw422016.N().S(`<th>`)
//line pages/admin/user/view/view.qtpl:91
		qw422016.E().S(p.Translate("id"))
//line pages/admin/user/view/view.qtpl:91
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:92
		qw422016.E().S(p.Translate("createdBy"))
//line pages/admin/user/view/view.qtpl:92
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:93
		qw422016.E().S(p.Translate("createdAt"))
//line pages/admin/user/view/view.qtpl:93
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:94
		qw422016.E().S(p.Translate("updatedBy"))
//line pages/admin/user/view/view.qtpl:94
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:95
		qw422016.E().S(p.Translate("updatedAt"))
//line pages/admin/user/view/view.qtpl:95
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:96
		qw422016.E().S(p.Translate("name"))
//line pages/admin/user/view/view.qtpl:96
		qw422016.N().S(`</th><th>`)
//line pages/admin/user/view/view.qtpl:97
		qw422016.E().S(p.TranslatePlural("comment"))
//line pages/admin/user/view/view.qtpl:97
		qw422016.N().S(`</th></tr>`)
//line pages/admin/user/view/view.qtpl:99
		for k := range p.contacts {
//line pages/admin/user/view/view.qtpl:99
			qw422016.N().S(`<tr>`)
//line pages/admin/user/view/view.qtpl:101
			if p.HasUserRight("UpdateUserContact") || p.HasUserRight("DeleteUserContact") {
//line pages/admin/user/view/view.qtpl:101
				qw422016.N().S(`<td><input type="checkbox" name="check" data-id="`)
//line pages/admin/user/view/view.qtpl:102
				qw422016.E().S(p.contacts[k].ID())
//line pages/admin/user/view/view.qtpl:102
				qw422016.N().S(`"></td><td><a href="`)
//line pages/admin/user/view/view.qtpl:103
				qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:103
				qw422016.N().S(`/User/Contact/Update?userid=`)
//line pages/admin/user/view/view.qtpl:103
				qw422016.E().S(p.contacts[k].ID())
//line pages/admin/user/view/view.qtpl:103
				qw422016.N().S(`&id=`)
//line pages/admin/user/view/view.qtpl:103
				qw422016.E().S(p.contacts[k].ID())
//line pages/admin/user/view/view.qtpl:103
				qw422016.N().S(`">✎</a></td>`)
//line pages/admin/user/view/view.qtpl:104
			}
//line pages/admin/user/view/view.qtpl:104
			qw422016.N().S(`<td>`)
//line pages/admin/user/view/view.qtpl:105
			qw422016.E().S(p.contacts[k].ID())
//line pages/admin/user/view/view.qtpl:105
			qw422016.N().S(`</td><td><a href="`)
//line pages/admin/user/view/view.qtpl:106
			qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:106
			qw422016.N().S(`/User/View?id=`)
//line pages/admin/user/view/view.qtpl:106
			qw422016.E().S(p.contacts[k].CreatedByID())
//line pages/admin/user/view/view.qtpl:106
			qw422016.N().S(`">`)
//line pages/admin/user/view/view.qtpl:106
			qw422016.E().S(p.rendererService.CreatedBy(p.contacts[k], p.language.ID()))
//line pages/admin/user/view/view.qtpl:106
			qw422016.N().S(`</a></td><td>`)
//line pages/admin/user/view/view.qtpl:107
			qw422016.E().S(p.contacts[k].CreatedAt().Format(time.RFC3339[0:9]))
//line pages/admin/user/view/view.qtpl:107
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:108
			if p.contacts[k].IsUpdated() {
//line pages/admin/user/view/view.qtpl:108
				qw422016.N().S(`<a href="`)
//line pages/admin/user/view/view.qtpl:108
				qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:108
				qw422016.N().S(`/User/View/`)
//line pages/admin/user/view/view.qtpl:108
				qw422016.E().S(*p.contacts[k].UpdatedByID())
//line pages/admin/user/view/view.qtpl:108
				qw422016.N().S(`">`)
//line pages/admin/user/view/view.qtpl:108
				qw422016.E().S(p.rendererService.UpdatedBy(p.contacts[k], p.language.ID()))
//line pages/admin/user/view/view.qtpl:108
				qw422016.N().S(`</a>`)
//line pages/admin/user/view/view.qtpl:108
			}
//line pages/admin/user/view/view.qtpl:108
			qw422016.N().S(`</td><td>`)
//line pages/admin/user/view/view.qtpl:109
			if p.contacts[k].IsUpdated() {
//line pages/admin/user/view/view.qtpl:109
				qw422016.E().S(p.contacts[k].UpdatedAt().Format(time.RFC3339[0:9]))
//line pages/admin/user/view/view.qtpl:109
			}
//line pages/admin/user/view/view.qtpl:109
			qw422016.N().S(`</td><td><a href="`)
//line pages/admin/user/view/view.qtpl:111
			qw422016.E().S(p.AdminURL())
//line pages/admin/user/view/view.qtpl:111
			qw422016.N().S(`/User/View?id=`)
//line pages/admin/user/view/view.qtpl:111
			qw422016.E().S(p.contacts[k].ContactID())
//line pages/admin/user/view/view.qtpl:111
			qw422016.N().S(`">`)
//line pages/admin/user/view/view.qtpl:111
			qw422016.E().S(p.userContactStore.Name(p.contacts[k], p.language.ID()))
//line pages/admin/user/view/view.qtpl:111
			qw422016.N().S(`</a></td><td>`)
//line pages/admin/user/view/view.qtpl:113
			if p.contacts[k].Comments() != nil {
//line pages/admin/user/view/view.qtpl:113
				qw422016.E().S(*p.contacts[k].Comments())
//line pages/admin/user/view/view.qtpl:113
			}
//line pages/admin/user/view/view.qtpl:113
			qw422016.N().S(`</td></tr>`)
//line pages/admin/user/view/view.qtpl:115
		}
//line pages/admin/user/view/view.qtpl:115
		qw422016.N().S(`</table>`)
//line pages/admin/user/view/view.qtpl:117
	} else {
//line pages/admin/user/view/view.qtpl:118
		qw422016.E().S(p.Translate("noResultsFound"))
//line pages/admin/user/view/view.qtpl:119
	}
//line pages/admin/user/view/view.qtpl:119
	qw422016.N().S(`</main>`)
//line pages/admin/user/view/view.qtpl:121
}

//line pages/admin/user/view/view.qtpl:121
func (p *Page) WriteContent(qq422016 qtio422016.Writer) {
//line pages/admin/user/view/view.qtpl:121
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/user/view/view.qtpl:121
	p.StreamContent(qw422016)
//line pages/admin/user/view/view.qtpl:121
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/user/view/view.qtpl:121
}

//line pages/admin/user/view/view.qtpl:121
func (p *Page) Content() string {
//line pages/admin/user/view/view.qtpl:121
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/user/view/view.qtpl:121
	p.WriteContent(qb422016)
//line pages/admin/user/view/view.qtpl:121
	qs422016 := string(qb422016.B)
//line pages/admin/user/view/view.qtpl:121
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/user/view/view.qtpl:121
	return qs422016
//line pages/admin/user/view/view.qtpl:121
}

//line pages/admin/user/view/view.qtpl:123
func (p *Page) StreamJavascripts(qw422016 *qt422016.Writer) {
//line pages/admin/user/view/view.qtpl:124
	if len(p.addresses) > 0 || len(p.contacts) > 0 {
//line pages/admin/user/view/view.qtpl:124
		qw422016.N().S(`<script src="/j/a/o.js"></script>`)
//line pages/admin/user/view/view.qtpl:126
		qw422016.N().S(p.rendererService.DefaultOverviewTranslations(p.GetCoreContext()))
//line pages/admin/user/view/view.qtpl:127
	}
//line pages/admin/user/view/view.qtpl:128
}

//line pages/admin/user/view/view.qtpl:128
func (p *Page) WriteJavascripts(qq422016 qtio422016.Writer) {
//line pages/admin/user/view/view.qtpl:128
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pages/admin/user/view/view.qtpl:128
	p.StreamJavascripts(qw422016)
//line pages/admin/user/view/view.qtpl:128
	qt422016.ReleaseWriter(qw422016)
//line pages/admin/user/view/view.qtpl:128
}

//line pages/admin/user/view/view.qtpl:128
func (p *Page) Javascripts() string {
//line pages/admin/user/view/view.qtpl:128
	qb422016 := qt422016.AcquireByteBuffer()
//line pages/admin/user/view/view.qtpl:128
	p.WriteJavascripts(qb422016)
//line pages/admin/user/view/view.qtpl:128
	qs422016 := string(qb422016.B)
//line pages/admin/user/view/view.qtpl:128
	qt422016.ReleaseByteBuffer(qb422016)
//line pages/admin/user/view/view.qtpl:128
	return qs422016
//line pages/admin/user/view/view.qtpl:128
}
