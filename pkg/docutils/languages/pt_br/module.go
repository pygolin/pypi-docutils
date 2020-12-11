package pt_br

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/pt_br.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nBrazilian Portuguese-language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, ßauthors.ToObject(), πg.NewUnicode("Autores").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßorganization.ToObject(), πg.NewUnicode("Organiza\xc3\xa7\xc3\xa3o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaddress.ToObject(), πg.NewUnicode("Endere\xc3\xa7o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontact.ToObject(), πg.NewUnicode("Contato").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßversion.ToObject(), πg.NewUnicode("Vers\xc3\xa3o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrevision.ToObject(), πg.NewUnicode("Revis\xc3\xa3o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstatus.ToObject(), πg.NewUnicode("Estado").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdate.ToObject(), πg.NewUnicode("Data").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcopyright.ToObject(), πg.NewUnicode("Copyright").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdedication.ToObject(), πg.NewUnicode("Dedicat\xc3\xb3ria").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßabstract.ToObject(), πg.NewUnicode("Resumo").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßattention.ToObject(), πg.NewUnicode("Aten\xc3\xa7\xc3\xa3o!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcaution.ToObject(), πg.NewUnicode("Cuidado!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdanger.ToObject(), πg.NewUnicode("PERIGO!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßerror.ToObject(), πg.NewUnicode("Erro").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhint.ToObject(), πg.NewUnicode("Sugest\xc3\xa3o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimportant.ToObject(), πg.NewUnicode("Importante").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnote.ToObject(), πg.NewUnicode("Nota").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtip.ToObject(), πg.NewUnicode("Dica").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwarning.ToObject(), πg.NewUnicode("Aviso").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontents.ToObject(), πg.NewUnicode("Sum\xc3\xa1rio").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("autor").ToObject(), ßauthor.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("autores").ToObject(), ßauthors.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("organiza\xc3\xa7\xc3\xa3o").ToObject(), ßorganization.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("endere\xc3\xa7o").ToObject(), ßaddress.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("contato").ToObject(), ßcontact.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vers\xc3\xa3o").ToObject(), ßversion.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("revis\xc3\xa3o").ToObject(), ßrevision.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("estado").ToObject(), ßstatus.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("data").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("copyright").ToObject(), ßcopyright.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dedicat\xc3\xb3ria").ToObject(), ßdedication.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("resumo").ToObject(), ßabstract.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 56: """Brazilian Portuguese (lowcased) to canonical name mapping for bibliographic fields."""
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
	πg.RegisterModule("pt_br", Code)
}
