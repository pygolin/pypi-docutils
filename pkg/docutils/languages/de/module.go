package de

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/de.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßAdresse := πg.InternStr("Adresse")
		ßAutor := πg.InternStr("Autor")
		ßAutoren := πg.InternStr("Autoren")
		ßBemerkung := πg.InternStr("Bemerkung")
		ßCopyright := πg.InternStr("Copyright")
		ßDatum := πg.InternStr("Datum")
		ßFehler := πg.InternStr("Fehler")
		ßHinweis := πg.InternStr("Hinweis")
		ßInhalt := πg.InternStr("Inhalt")
		ßKontakt := πg.InternStr("Kontakt")
		ßOrganisation := πg.InternStr("Organisation")
		ßRevision := πg.InternStr("Revision")
		ßStatus := πg.InternStr("Status")
		ßTipp := πg.InternStr("Tipp")
		ßVersion := πg.InternStr("Version")
		ßWarnung := πg.InternStr("Warnung")
		ßWichtig := πg.InternStr("Wichtig")
		ßWidmung := πg.InternStr("Widmung")
		ßZusammenfassung := πg.InternStr("Zusammenfassung")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßabstract := πg.InternStr("abstract")
		ßaddress := πg.InternStr("address")
		ßadresse := πg.InternStr("adresse")
		ßattention := πg.InternStr("attention")
		ßauthor := πg.InternStr("author")
		ßauthor_separators := πg.InternStr("author_separators")
		ßauthors := πg.InternStr("authors")
		ßautor := πg.InternStr("autor")
		ßautoren := πg.InternStr("autoren")
		ßbibliographic_fields := πg.InternStr("bibliographic_fields")
		ßcaution := πg.InternStr("caution")
		ßcontact := πg.InternStr("contact")
		ßcontents := πg.InternStr("contents")
		ßcopyright := πg.InternStr("copyright")
		ßdanger := πg.InternStr("danger")
		ßdate := πg.InternStr("date")
		ßdatum := πg.InternStr("datum")
		ßdedication := πg.InternStr("dedication")
		ßerror := πg.InternStr("error")
		ßhint := πg.InternStr("hint")
		ßimportant := πg.InternStr("important")
		ßkontakt := πg.InternStr("kontakt")
		ßlabels := πg.InternStr("labels")
		ßnote := πg.InternStr("note")
		ßorganisation := πg.InternStr("organisation")
		ßorganization := πg.InternStr("organization")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrevision := πg.InternStr("revision")
		ßstatus := πg.InternStr("status")
		ßtip := πg.InternStr("tip")
		ßversion := πg.InternStr("version")
		ßwarning := πg.InternStr("warning")
		ßwidmung := πg.InternStr("widmung")
		ßzusammenfassung := πg.InternStr("zusammenfassung")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nGerman language mappings for language-dependent features of Docutils.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, ßauthor.ToObject(), ßAutor.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßauthors.ToObject(), ßAutoren.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßorganization.ToObject(), ßOrganisation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaddress.ToObject(), ßAdresse.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontact.ToObject(), ßKontakt.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßversion.ToObject(), ßVersion.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrevision.ToObject(), ßRevision.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstatus.ToObject(), ßStatus.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdate.ToObject(), ßDatum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdedication.ToObject(), ßWidmung.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcopyright.ToObject(), ßCopyright.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßabstract.ToObject(), ßZusammenfassung.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßattention.ToObject(), πg.NewStr("Achtung!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcaution.ToObject(), πg.NewStr("Vorsicht!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdanger.ToObject(), πg.NewStr("!GEFAHR!").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßerror.ToObject(), ßFehler.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhint.ToObject(), ßHinweis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimportant.ToObject(), ßWichtig.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnote.ToObject(), ßBemerkung.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtip.ToObject(), ßTipp.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwarning.ToObject(), ßWarnung.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcontents.ToObject(), ßInhalt.ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, ßautor.ToObject(), ßauthor.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßautoren.ToObject(), ßauthors.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßorganisation.ToObject(), ßorganization.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßadresse.ToObject(), ßaddress.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßkontakt.ToObject(), ßcontact.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßversion.ToObject(), ßversion.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrevision.ToObject(), ßrevision.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstatus.ToObject(), ßstatus.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdatum.ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcopyright.ToObject(), ßcopyright.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwidmung.ToObject(), ßdedication.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßzusammenfassung.ToObject(), ßabstract.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßbibliographic_fields.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 54: """German (lowcased) to canonical name mapping for bibliographic fields."""
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
	πg.RegisterModule("de", Code)
}
