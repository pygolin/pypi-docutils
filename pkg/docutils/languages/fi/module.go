package fi

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/fi.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
			// line 10: """
			πF.SetLineno(10)
			// line 10: """
			πF.SetLineno(10)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nFinnish-language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
				continue
			}
			// line 14: __docformat__ = 'reStructuredText'
			πF.SetLineno(14)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 16: labels = {
			πF.SetLineno(16)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("author").ToObject(), πg.NewUnicode("Tekij\xc3\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("authors").ToObject(), πg.NewUnicode("Tekij\xc3\xa4t").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("organization").ToObject(), πg.NewUnicode("Yhteis\xc3\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("address").ToObject(), πg.NewUnicode("Osoite").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contact").ToObject(), πg.NewUnicode("Yhteystiedot").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("version").ToObject(), πg.NewUnicode("Versio").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("revision").ToObject(), πg.NewUnicode("Vedos").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("status").ToObject(), πg.NewUnicode("Tila").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("date").ToObject(), πg.NewUnicode("P\xc3\xa4iv\xc3\xa4ys").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("copyright").ToObject(), πg.NewUnicode("Tekij\xc3\xa4noikeudet").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dedication").ToObject(), πg.NewUnicode("Omistuskirjoitus").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("abstract").ToObject(), πg.NewUnicode("Tiivistelm\xc3\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("attention").ToObject(), πg.NewUnicode("Huomio!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("caution").ToObject(), πg.NewUnicode("Varo!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("danger").ToObject(), πg.NewUnicode("!VAARA!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("error").ToObject(), πg.NewUnicode("Virhe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("hint").ToObject(), πg.NewUnicode("Vihje").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("important").ToObject(), πg.NewUnicode("T\xc3\xa4rke\xc3\xa4\xc3\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("note").ToObject(), πg.NewUnicode("Huomautus").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tip").ToObject(), πg.NewUnicode("Neuvo").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("warning").ToObject(), πg.NewUnicode("Varoitus").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contents").ToObject(), πg.NewUnicode("Sis\xc3\xa4llys").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tekij\xc3\xa4").ToObject(), πg.NewUnicode("author").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tekij\xc3\xa4t").ToObject(), πg.NewUnicode("authors").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("yhteis\xc3\xb6").ToObject(), πg.NewUnicode("organization").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("osoite").ToObject(), πg.NewUnicode("address").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("yhteystiedot").ToObject(), πg.NewUnicode("contact").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("versio").ToObject(), πg.NewUnicode("version").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vedos").ToObject(), πg.NewUnicode("revision").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tila").ToObject(), πg.NewUnicode("status").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("p\xc3\xa4iv\xc3\xa4ys").ToObject(), πg.NewUnicode("date").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tekij\xc3\xa4noikeudet").ToObject(), πg.NewUnicode("copyright").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("omistuskirjoitus").ToObject(), πg.NewUnicode("dedication").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tiivistelm\xc3\xa4").ToObject(), πg.NewUnicode("abstract").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 56: """Finnish (lowcased) to canonical name mapping for bibliographic fields."""
			πF.SetLineno(56)
			// line 58: author_separators = [';', ',']
			πF.SetLineno(58)
			πTemp003 = make([]*πg.Object, 2)
			πTemp003[0] = πg.NewStr(";").ToObject()
			πTemp003[1] = πg.NewStr(",").ToObject()
			πTemp002 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßauthor_separators.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 59: """List of separator strings for the 'Authors' bibliographic field. Tried in
			πF.SetLineno(59)
		}
		return nil, πE
	})
	πg.RegisterModule("fi", Code)
}
