package sk

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/sk.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßabstract := πg.InternStr("abstract")
		ßaddress := πg.InternStr("address")
		ßattention := πg.InternStr("attention")
		ßauthor := πg.InternStr("author")
		ßauthor_separators := πg.InternStr("author_separators")
		ßauthors := πg.InternStr("authors")
		ßbibliographic_fields := πg.InternStr("bibliographic_fields")
		ßcaution := πg.InternStr("caution")
		ßcontact := πg.InternStr("contact")
		ßcontents := πg.InternStr("contents")
		ßcopyright := πg.InternStr("copyright")
		ßdanger := πg.InternStr("danger")
		ßdate := πg.InternStr("date")
		ßdedication := πg.InternStr("dedication")
		ßerror := πg.InternStr("error")
		ßhint := πg.InternStr("hint")
		ßimportant := πg.InternStr("important")
		ßlabels := πg.InternStr("labels")
		ßnote := πg.InternStr("note")
		ßorganization := πg.InternStr("organization")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrevision := πg.InternStr("revision")
		ßstatus := πg.InternStr("status")
		ßtip := πg.InternStr("tip")
		ßversion := πg.InternStr("version")
		ßwarning := πg.InternStr("warning")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nSlovak-language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, ßauthor.ToObject(), πg.NewUnicode("Autor").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßauthors.ToObject(), πg.NewUnicode("Autori").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßorganization.ToObject(), πg.NewUnicode("Organiz\xc3\xa1cia").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaddress.ToObject(), πg.NewUnicode("Adresa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontact.ToObject(), πg.NewUnicode("Kontakt").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßversion.ToObject(), πg.NewUnicode("Verzia").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrevision.ToObject(), πg.NewUnicode("Rev\xc3\xadzia").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstatus.ToObject(), πg.NewUnicode("Stav").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdate.ToObject(), πg.NewUnicode("D\xc3\xa1tum").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcopyright.ToObject(), πg.NewUnicode("Copyright").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdedication.ToObject(), πg.NewUnicode("Venovanie").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßabstract.ToObject(), πg.NewUnicode("Abstraktne").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßattention.ToObject(), πg.NewUnicode("Pozor!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcaution.ToObject(), πg.NewUnicode("Opatrne!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdanger.ToObject(), πg.NewUnicode("!NEBEZPE\xc4\x8cENSTVO!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßerror.ToObject(), πg.NewUnicode("Chyba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhint.ToObject(), πg.NewUnicode("Rada").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimportant.ToObject(), πg.NewUnicode("D\xc3\xb4le\xc5\xbeit\xc3\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnote.ToObject(), πg.NewUnicode("Pozn\xc3\xa1mka").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtip.ToObject(), πg.NewUnicode("Tip").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwarning.ToObject(), πg.NewUnicode("Varovanie").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontents.ToObject(), πg.NewUnicode("Obsah").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßlabels.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 39: """Mapping of node class name to label text."""
			πF.SetLineno(39)
			// line 41: bibliographic_fields = {
			πF.SetLineno(41)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("autor").ToObject(), ßauthor.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("autori").ToObject(), ßauthors.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("organiz\xc3\xa1cia").ToObject(), ßorganization.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("adresa").ToObject(), ßaddress.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kontakt").ToObject(), ßcontact.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("verzia").ToObject(), ßversion.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rev\xc3\xadzia").ToObject(), ßrevision.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("stav").ToObject(), ßstatus.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("d\xc3\xa1tum").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("copyright").ToObject(), ßcopyright.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("venovanie").ToObject(), ßdedication.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("abstraktne").ToObject(), ßabstract.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 54: """Slovak (lowcased) to canonical name mapping for bibliographic fields."""
			πF.SetLineno(54)
			// line 56: author_separators = [';', ',']
			πF.SetLineno(56)
			πTemp003 = make([]*πg.Object, 2)
			πTemp003[0] = πg.NewStr(";").ToObject()
			πTemp003[1] = πg.NewStr(",").ToObject()
			πTemp002 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßauthor_separators.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 57: """List of separator strings for the 'Authors' bibliographic field. Tried in
			πF.SetLineno(57)
		}
		return nil, πE
	})
	πg.RegisterModule("sk", Code)
}
