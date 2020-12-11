package sk

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/sk.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nSlovak-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pozor").ToObject(), ßattention.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("opatrne").ToObject(), ßcaution.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code (translation required)").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("nebezpe\xc3\xa8enstvo").ToObject(), ßdanger.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("chyba").ToObject(), ßerror.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rada").ToObject(), ßhint.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("d\xc3\xb4le\xc2\x9eit\xc3\xa9").ToObject(), ßimportant.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pozn\xc3\xa1mka").ToObject(), ßnote.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tip (translation required)").ToObject(), ßtip.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("varovanie").ToObject(), ßwarning.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("admonition (translation required)").ToObject(), ßadmonition.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sidebar (translation required)").ToObject(), ßsidebar.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("t\xc3\xa9ma").ToObject(), ßtopic.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("blok-riadkov").ToObject(), πg.NewStr("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("parsed-literal").ToObject(), πg.NewStr("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rubric (translation required)").ToObject(), ßrubric.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("epigraph (translation required)").ToObject(), ßepigraph.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("highlights (translation required)").ToObject(), ßhighlights.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pull-quote (translation required)").ToObject(), πg.NewStr("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("compound (translation required)").ToObject(), ßcompound.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("container (translation required)").ToObject(), ßcontainer.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("table (translation required)").ToObject(), ßtable.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("csv-table (translation required)").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("list-table (translation required)").ToObject(), πg.NewStr("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("meta").ToObject(), ßmeta.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("obr\xc3\xa1zok").ToObject(), ßimage.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tvar").ToObject(), ßfigure.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vlo\xc2\x9ei\xc2\x9d").ToObject(), ßinclude.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("raw (translation required)").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("nahradi\xc2\x9d").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("unicode").ToObject(), ßunicode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("d\xc3\xa1tum").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("class (translation required)").ToObject(), ßclass.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("role (translation required)").ToObject(), ßrole.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("default-role (translation required)").ToObject(), πg.NewStr("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("title (translation required)").ToObject(), ßtitle.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("obsah").ToObject(), ßcontents.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xc3\xa8as\xc2\x9d").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xc3\xa8as\xc2\x9d-\xc3\xa8\xc3\xadslovanie").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("cie\xc2\xbeov\xc3\xa9-pozn\xc3\xa1mky").ToObject(), πg.NewStr("target-notes").ToObject()); πE != nil {
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
			// line 69: """Slovak name to registered (in directives/__init__.py) directive name
			πF.SetLineno(69)
			// line 72: roles = {
			πF.SetLineno(72)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("abbreviation (translation required)").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("acronym (translation required)").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code (translation required)").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("index (translation required)").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("subscript (translation required)").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("superscript (translation required)").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("title-reference (translation required)").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pep-reference (translation required)").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc-reference (translation required)").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("emphasis (translation required)").ToObject(), ßemphasis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("strong (translation required)").ToObject(), ßstrong.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("literal (translation required)").ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("named-reference (translation required)").ToObject(), πg.NewStr("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anonymous-reference (translation required)").ToObject(), πg.NewStr("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("footnote-reference (translation required)").ToObject(), πg.NewStr("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("citation-reference (translation required)").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("substitution-reference (translation required)").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("target (translation required)").ToObject(), ßtarget.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri-reference (translation required)").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("raw (translation required)").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 94: """Mapping of Slovak role names to canonical role names for interpreted text.
			πF.SetLineno(94)
		}
		return nil, πE
	})
	πg.RegisterModule("sk", Code)
}
