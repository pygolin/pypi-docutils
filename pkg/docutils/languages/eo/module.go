package eo

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/eo.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßabstract := πg.InternStr("abstract")
		ßaddress := πg.InternStr("address")
		ßadreso := πg.InternStr("adreso")
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
		ßdato := πg.InternStr("dato")
		ßdedication := πg.InternStr("dedication")
		ßerror := πg.InternStr("error")
		ßhint := πg.InternStr("hint")
		ßimportant := πg.InternStr("important")
		ßkontakto := πg.InternStr("kontakto")
		ßlabels := πg.InternStr("labels")
		ßnote := πg.InternStr("note")
		ßorganization := πg.InternStr("organization")
		ßorganizo := πg.InternStr("organizo")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßresumo := πg.InternStr("resumo")
		ßrevido := πg.InternStr("revido")
		ßrevision := πg.InternStr("revision")
		ßstato := πg.InternStr("stato")
		ßstatus := πg.InternStr("status")
		ßtip := πg.InternStr("tip")
		ßversio := πg.InternStr("versio")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nEsperanto-language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, ßauthor.ToObject(), πg.NewUnicode("A\xc5\xadtoro").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßauthors.ToObject(), πg.NewUnicode("A\xc5\xadtoroj").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßorganization.ToObject(), πg.NewUnicode("Organizo").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaddress.ToObject(), πg.NewUnicode("Adreso").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontact.ToObject(), πg.NewUnicode("Kontakto").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßversion.ToObject(), πg.NewUnicode("Versio").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrevision.ToObject(), πg.NewUnicode("Revido").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstatus.ToObject(), πg.NewUnicode("Stato").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdate.ToObject(), πg.NewUnicode("Dato").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcopyright.ToObject(), πg.NewUnicode("A\xc5\xadtorrajto").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdedication.ToObject(), πg.NewUnicode("Dedi\xc4\x89o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßabstract.ToObject(), πg.NewUnicode("Resumo").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßattention.ToObject(), πg.NewUnicode("Atentu!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcaution.ToObject(), πg.NewUnicode("Zorgu!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdanger.ToObject(), πg.NewUnicode("DAN\xc4\x9cERO!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßerror.ToObject(), πg.NewUnicode("Eraro").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhint.ToObject(), πg.NewUnicode("Spuro").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimportant.ToObject(), πg.NewUnicode("Grava").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnote.ToObject(), πg.NewUnicode("Noto").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtip.ToObject(), πg.NewUnicode("Helpeto").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwarning.ToObject(), πg.NewUnicode("Averto").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontents.ToObject(), πg.NewUnicode("Enhavo").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewStr("a\\u016dtoro").ToObject(), ßauthor.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("a\\u016dtoroj").ToObject(), ßauthors.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßorganizo.ToObject(), ßorganization.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßadreso.ToObject(), ßaddress.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßkontakto.ToObject(), ßcontact.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßversio.ToObject(), ßversion.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrevido.ToObject(), ßrevision.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstato.ToObject(), ßstatus.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdato.ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("a\\u016dtorrajto").ToObject(), ßcopyright.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("dedi\\u0109o").ToObject(), ßdedication.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßresumo.ToObject(), ßabstract.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 57: """Esperanto (lowcased) to canonical name mapping for bibliographic fields."""
			πF.SetLineno(57)
			// line 59: author_separators = [';', ',']
			πF.SetLineno(59)
			πTemp003 = make([]*πg.Object, 2)
			πTemp003[0] = πg.NewStr(";").ToObject()
			πTemp003[1] = πg.NewStr(",").ToObject()
			πTemp002 = πg.NewList(πTemp003...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßauthor_separators.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 60: """List of separator strings for the 'Authors' bibliographic field. Tried in
			πF.SetLineno(60)
		}
		return nil, πE
	})
	πg.RegisterModule("eo", Code)
}
