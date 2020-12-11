package fr

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/fr.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßabbreviation := πg.InternStr("abbreviation")
		ßacronym := πg.InternStr("acronym")
		ßadmonition := πg.InternStr("admonition")
		ßattention := πg.InternStr("attention")
		ßcaution := πg.InternStr("caution")
		ßclass := πg.InternStr("class")
		ßcode := πg.InternStr("code")
		ßcompound := πg.InternStr("compound")
		ßcontainer := πg.InternStr("container")
		ßcontents := πg.InternStr("contents")
		ßdanger := πg.InternStr("danger")
		ßdate := πg.InternStr("date")
		ßdirectives := πg.InternStr("directives")
		ßemphasis := πg.InternStr("emphasis")
		ßepigraph := πg.InternStr("epigraph")
		ßerror := πg.InternStr("error")
		ßfigure := πg.InternStr("figure")
		ßfooter := πg.InternStr("footer")
		ßheader := πg.InternStr("header")
		ßhighlights := πg.InternStr("highlights")
		ßhint := πg.InternStr("hint")
		ßimage := πg.InternStr("image")
		ßimportant := πg.InternStr("important")
		ßinclude := πg.InternStr("include")
		ßindex := πg.InternStr("index")
		ßliteral := πg.InternStr("literal")
		ßmath := πg.InternStr("math")
		ßmeta := πg.InternStr("meta")
		ßnote := πg.InternStr("note")
		ßraw := πg.InternStr("raw")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreplace := πg.InternStr("replace")
		ßrole := πg.InternStr("role")
		ßroles := πg.InternStr("roles")
		ßrubric := πg.InternStr("rubric")
		ßsectnum := πg.InternStr("sectnum")
		ßsidebar := πg.InternStr("sidebar")
		ßstrong := πg.InternStr("strong")
		ßsubscript := πg.InternStr("subscript")
		ßsuperscript := πg.InternStr("superscript")
		ßtable := πg.InternStr("table")
		ßtarget := πg.InternStr("target")
		ßtip := πg.InternStr("tip")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßunicode := πg.InternStr("unicode")
		ßwarning := πg.InternStr("warning")
		var πTemp001 *πg.Dict
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nFrench-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
				continue
			}
			// line 15: __docformat__ = 'reStructuredText'
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 18: directives = {
			πF.SetLineno(18)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("attention").ToObject(), ßattention.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pr\xc3\xa9caution").ToObject(), ßcaution.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("danger").ToObject(), ßdanger.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("erreur").ToObject(), ßerror.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("conseil").ToObject(), ßhint.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("important").ToObject(), ßimportant.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("note").ToObject(), ßnote.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("astuce").ToObject(), ßtip.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("avertissement").ToObject(), ßwarning.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("admonition").ToObject(), ßadmonition.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("encadr\xc3\xa9").ToObject(), ßsidebar.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sujet").ToObject(), ßtopic.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("bloc-textuel").ToObject(), πg.NewStr("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("bloc-interpr\xc3\xa9t\xc3\xa9").ToObject(), πg.NewStr("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code-interpr\xc3\xa9t\xc3\xa9").ToObject(), πg.NewStr("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("intertitre").ToObject(), ßrubric.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("exergue").ToObject(), ßepigraph.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xc3\xa9pigraphe").ToObject(), ßepigraph.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("chapeau").ToObject(), ßhighlights.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("accroche").ToObject(), πg.NewStr("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("compound (translation required)").ToObject(), ßcompound.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("container (translation required)").ToObject(), ßcontainer.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tableau").ToObject(), ßtable.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("csv-table (translation required)").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("list-table (translation required)").ToObject(), πg.NewStr("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("m\xc3\xa9ta").ToObject(), ßmeta.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("image").ToObject(), ßimage.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("figure").ToObject(), ßfigure.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("inclure").ToObject(), ßinclude.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("brut").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("remplacer").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("remplace").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("unicode").ToObject(), ßunicode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("date").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("classe").ToObject(), ßclass.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("role (translation required)").ToObject(), ßrole.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("default-role (translation required)").ToObject(), πg.NewStr("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titre (translation required)").ToObject(), ßtitle.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sommaire").ToObject(), ßcontents.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("table-des-mati\xc3\xa8res").ToObject(), ßcontents.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sectnum").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("section-num\xc3\xa9rot\xc3\xa9e").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("liens").ToObject(), πg.NewStr("target-notes").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("header (translation required)").ToObject(), ßheader.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("footer (translation required)").ToObject(), ßfooter.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 73: """French name to registered (in directives/__init__.py) directive name
			πF.SetLineno(73)
			// line 76: roles = {
			πF.SetLineno(76)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("abr\xc3\xa9viation").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("acronyme").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sigle").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("index").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("indice").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("ind").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("exposant").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("exp").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titre-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titre").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pep-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("emphase").ToObject(), ßemphasis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("fort").ToObject(), ßstrong.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("litt\xc3\xa9ral").ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("nomm\xc3\xa9e-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anonyme-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("note-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("citation-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("substitution-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("lien").ToObject(), ßtarget.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri-r\xc3\xa9f\xc3\xa9rence").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("brut").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 102: """Mapping of French role names to canonical role names for interpreted text.
			πF.SetLineno(102)
		}
		return nil, πE
	})
	πg.RegisterModule("fr", Code)
}
