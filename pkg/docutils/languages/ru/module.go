package ru

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/ru.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nRussian-language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("abstract").ToObject(), πg.NewUnicode("\xd0\x90\xd0\xbd\xd0\xbd\xd0\xbe\xd1\x82\xd0\xb0\xd1\x86\xd0\xb8\xd1\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("address").ToObject(), πg.NewUnicode("\xd0\x90\xd0\xb4\xd1\x80\xd0\xb5\xd1\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("attention").ToObject(), πg.NewUnicode("\xd0\x92\xd0\xbd\xd0\xb8\xd0\xbc\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("author").ToObject(), πg.NewUnicode("\xd0\x90\xd0\xb2\xd1\x82\xd0\xbe\xd1\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("authors").ToObject(), πg.NewUnicode("\xd0\x90\xd0\xb2\xd1\x82\xd0\xbe\xd1\x80\xd1\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("caution").ToObject(), πg.NewUnicode("\xd0\x9e\xd1\x81\xd1\x82\xd0\xbe\xd1\x80\xd0\xbe\xd0\xb6\xd0\xbd\xd0\xbe!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contact").ToObject(), πg.NewUnicode("\xd0\x9a\xd0\xbe\xd0\xbd\xd1\x82\xd0\xb0\xd0\xba\xd1\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contents").ToObject(), πg.NewUnicode("\xd0\xa1\xd0\xbe\xd0\xb4\xd0\xb5\xd1\x80\xd0\xb6\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("copyright").ToObject(), πg.NewUnicode("\xd0\x9f\xd1\x80\xd0\xb0\xd0\xb2\xd0\xb0 \xd0\xba\xd0\xbe\xd0\xbf\xd0\xb8\xd1\x80\xd0\xbe\xd0\xb2\xd0\xb0\xd0\xbd\xd0\xb8\xd1\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("danger").ToObject(), πg.NewUnicode("\xd0\x9e\xd0\x9f\xd0\x90\xd0\xa1\xd0\x9d\xd0\x9e!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("date").ToObject(), πg.NewUnicode("\xd0\x94\xd0\xb0\xd1\x82\xd0\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dedication").ToObject(), πg.NewUnicode("\xd0\x9f\xd0\xbe\xd1\x81\xd0\xb2\xd1\x8f\xd1\x89\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("error").ToObject(), πg.NewUnicode("\xd0\x9e\xd1\x88\xd0\xb8\xd0\xb1\xd0\xba\xd0\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("hint").ToObject(), πg.NewUnicode("\xd0\xa1\xd0\xbe\xd0\xb2\xd0\xb5\xd1\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("important").ToObject(), πg.NewUnicode("\xd0\x92\xd0\xb0\xd0\xb6\xd0\xbd\xd0\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("note").ToObject(), πg.NewUnicode("\xd0\x9f\xd1\x80\xd0\xb8\xd0\xbc\xd0\xb5\xd1\x87\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("organization").ToObject(), πg.NewUnicode("\xd0\x9e\xd1\x80\xd0\xb3\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb7\xd0\xb0\xd1\x86\xd0\xb8\xd1\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("revision").ToObject(), πg.NewUnicode("\xd0\xa0\xd0\xb5\xd0\xb4\xd0\xb0\xd0\xba\xd1\x86\xd0\xb8\xd1\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("status").ToObject(), πg.NewUnicode("\xd0\xa1\xd1\x82\xd0\xb0\xd1\x82\xd1\x83\xd1\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tip").ToObject(), πg.NewUnicode("\xd0\x9f\xd0\xbe\xd0\xb4\xd1\x81\xd0\xba\xd0\xb0\xd0\xb7\xd0\xba\xd0\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("version").ToObject(), πg.NewUnicode("\xd0\x92\xd0\xb5\xd1\x80\xd1\x81\xd0\xb8\xd1\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("warning").ToObject(), πg.NewUnicode("\xd0\x9f\xd1\x80\xd0\xb5\xd0\xb4\xd1\x83\xd0\xbf\xd1\x80\xd0\xb5\xd0\xb6\xd0\xb4\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßlabels.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 40: """Mapping of node class name to label text."""
			πF.SetLineno(40)
			// line 42: bibliographic_fields = {
			πF.SetLineno(42)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb0\xd0\xbd\xd0\xbd\xd0\xbe\xd1\x82\xd0\xb0\xd1\x86\xd0\xb8\xd1\x8f").ToObject(), πg.NewUnicode("abstract").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb0\xd0\xb4\xd1\x80\xd0\xb5\xd1\x81").ToObject(), πg.NewUnicode("address").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb0\xd0\xb2\xd1\x82\xd0\xbe\xd1\x80").ToObject(), πg.NewUnicode("author").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb0\xd0\xb2\xd1\x82\xd0\xbe\xd1\x80\xd1\x8b").ToObject(), πg.NewUnicode("authors").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xba\xd0\xbe\xd0\xbd\xd1\x82\xd0\xb0\xd0\xba\xd1\x82").ToObject(), πg.NewUnicode("contact").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbf\xd1\x80\xd0\xb0\xd0\xb2\xd0\xb0 \xd0\xba\xd0\xbe\xd0\xbf\xd0\xb8\xd1\x80\xd0\xbe\xd0\xb2\xd0\xb0\xd0\xbd\xd0\xb8\xd1\x8f").ToObject(), πg.NewUnicode("copyright").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb4\xd0\xb0\xd1\x82\xd0\xb0").ToObject(), πg.NewUnicode("date").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbf\xd0\xbe\xd1\x81\xd0\xb2\xd1\x8f\xd1\x89\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("dedication").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbe\xd1\x80\xd0\xb3\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb7\xd0\xb0\xd1\x86\xd0\xb8\xd1\x8f").ToObject(), πg.NewUnicode("organization").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x80\xd0\xb5\xd0\xb4\xd0\xb0\xd0\xba\xd1\x86\xd0\xb8\xd1\x8f").ToObject(), πg.NewUnicode("revision").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x82\xd0\xb0\xd1\x82\xd1\x83\xd1\x81").ToObject(), πg.NewUnicode("status").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd0\xb5\xd1\x80\xd1\x81\xd0\xb8\xd1\x8f").ToObject(), πg.NewUnicode("version").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 55: """Russian (lowcased) to canonical name mapping for bibliographic fields."""
			πF.SetLineno(55)
			// line 57: author_separators =  [';', ',']
			πF.SetLineno(57)
			πTemp003 = make([]*πg.Object, 2)
			πTemp003[0] = πg.NewStr(";").ToObject()
			πTemp003[1] = πg.NewStr(",").ToObject()
			πTemp002 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßauthor_separators.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 58: """List of separator strings for the 'Authors' bibliographic field. Tried in
			πF.SetLineno(58)
		}
		return nil, πE
	})
	πg.RegisterModule("ru", Code)
}
