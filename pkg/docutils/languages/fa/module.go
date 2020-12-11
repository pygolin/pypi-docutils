package fa

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/fa.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßauthor_separators := πg.InternStr("author_separators")
		ßbibliographic_fields := πg.InternStr("bibliographic_fields")
		ßlabels := πg.InternStr("labels")
		ßreStructuredText := πg.InternStr("reStructuredText")
		var πTemp001 *πg.Dict
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		var πTemp003 []*πg.Object
		_ = πTemp003
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 11: """
			πF.SetLineno(11)
			// line 11: """
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nPersian-language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
				continue
			}
			// line 15: __docformat__ = 'reStructuredText'
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 17: labels = {
			πF.SetLineno(17)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("author").ToObject(), πg.NewUnicode("\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3\xd9\x86\xd8\xaf\xd9\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("authors").ToObject(), πg.NewUnicode("\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3\xd9\x86\xd8\xaf\xda\xaf\xd8\xa7\xd9\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("organization").ToObject(), πg.NewUnicode("\xd8\xb3\xd8\xa7\xd8\xb2\xd9\x85\xd8\xa7\xd9\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("address").ToObject(), πg.NewUnicode("\xd8\xa2\xd8\xaf\xd8\xb1\xd8\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contact").ToObject(), πg.NewUnicode("\xd8\xaa\xd9\x85\xd8\xa7\xd8\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("version").ToObject(), πg.NewUnicode("\xd9\x86\xd8\xb3\xd8\xae\xd9\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("revision").ToObject(), πg.NewUnicode("\xd8\xa8\xd8\xa7\xd8\xb2\xd8\xa8\xdb\x8c\xd9\x86\xdb\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("status").ToObject(), πg.NewUnicode("\xd9\x88\xd8\xb6\xd8\xb9\xdb\x8c\xd8\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("date").ToObject(), πg.NewUnicode("\xd8\xaa\xd8\xa7\xd8\xb1\xdb\x8c\xd8\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("copyright").ToObject(), πg.NewUnicode("\xda\xa9\xd9\xbe\xdb\x8c\xe2\x80\x8c\xd8\xb1\xd8\xa7\xdb\x8c\xd8\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dedication").ToObject(), πg.NewUnicode("\xd8\xaa\xd8\xae\xd8\xb5\xdb\x8c\xd8\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("abstract").ToObject(), πg.NewUnicode("\xda\x86\xda\xa9\xdb\x8c\xd8\xaf\xd9\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("attention").ToObject(), πg.NewUnicode("\xd8\xaa\xd9\x88\xd8\xac\xd9\x87!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("caution").ToObject(), πg.NewUnicode("\xd8\xa7\xd8\xad\xd8\xaa\xdb\x8c\xd8\xa7\xd8\xb7!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("danger").ToObject(), πg.NewUnicode("\xd8\xae\xd8\xb7\xd8\xb1!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("error").ToObject(), πg.NewUnicode("\xd8\xae\xd8\xb7\xd8\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("hint").ToObject(), πg.NewUnicode("\xd8\xb1\xd8\xa7\xd9\x87\xd9\x86\xd9\x85\xd8\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("important").ToObject(), πg.NewUnicode("\xd9\x85\xd9\x87\xd9\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("note").ToObject(), πg.NewUnicode("\xdb\x8c\xd8\xa7\xd8\xaf\xd8\xaf\xd8\xa7\xd8\xb4\xd8\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tip").ToObject(), πg.NewUnicode("\xd9\x86\xda\xa9\xd8\xaa\xd9\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("warning").ToObject(), πg.NewUnicode("\xd8\xa7\xd8\xae\xd8\xb7\xd8\xa7\xd8\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contents").ToObject(), πg.NewUnicode("\xd9\x85\xd8\xad\xd8\xaa\xd9\x88\xd8\xa7").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßlabels.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 41: """Mapping of node class name to label text."""
			πF.SetLineno(41)
			// line 43: bibliographic_fields = {
			πF.SetLineno(43)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3\xd9\x86\xd8\xaf\xd9\x87").ToObject(), πg.NewUnicode("author").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3\xd9\x86\xd8\xaf\xda\xaf\xd8\xa7\xd9\x86").ToObject(), πg.NewUnicode("authors").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb3\xd8\xa7\xd8\xb2\xd9\x85\xd8\xa7\xd9\x86").ToObject(), πg.NewUnicode("organization").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa2\xd8\xaf\xd8\xb1\xd8\xb3").ToObject(), πg.NewUnicode("address").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd9\x85\xd8\xa7\xd8\xb3").ToObject(), πg.NewUnicode("contact").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xd8\xb3\xd8\xae\xd9\x87").ToObject(), πg.NewUnicode("version").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa8\xd8\xa7\xd8\xb2\xd8\xa8\xdb\x8c\xd9\x86\xdb\x8c").ToObject(), πg.NewUnicode("revision").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x88\xd8\xb6\xd8\xb9\xdb\x8c\xd8\xaa").ToObject(), πg.NewUnicode("status").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xa7\xd8\xb1\xdb\x8c\xd8\xae").ToObject(), πg.NewUnicode("date").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\xa9\xd9\xbe\xdb\x8c\xe2\x80\x8c\xd8\xb1\xd8\xa7\xdb\x8c\xd8\xaa").ToObject(), πg.NewUnicode("copyright").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xae\xd8\xb5\xdb\x8c\xd8\xb5").ToObject(), πg.NewUnicode("dedication").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\x86\xda\xa9\xdb\x8c\xd8\xaf\xd9\x87").ToObject(), πg.NewUnicode("abstract").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 57: """Persian (lowcased) to canonical name mapping for bibliographic fields."""
			πF.SetLineno(57)
			// line 59: author_separators = [u'؛', u'،']
			πF.SetLineno(59)
			πTemp003 = make([]*πg.Object, 2)
			πTemp003[0] = πg.NewUnicode("\xd8\x9b").ToObject()
			πTemp003[1] = πg.NewUnicode("\xd8\x8c").ToObject()
			πTemp002 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßauthor_separators.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 60: """List of separator strings for the 'Authors' bibliographic field. Tried in
			πF.SetLineno(60)
		}
		return nil, πE
	})
	πg.RegisterModule("fa", Code)
}
