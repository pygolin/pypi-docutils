package da

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/da.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
			// line 11: """
			πF.SetLineno(11)
			// line 11: """
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDanish-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
				continue
			}
			// line 16: __docformat__ = 'reStructuredText'
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 19: directives = {
			πF.SetLineno(19)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("giv agt").ToObject(), ßattention.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pas p\xc3\xa5").ToObject(), ßcaution.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kode").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kode-blok").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kildekode").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("fare").ToObject(), ßdanger.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("fejl").ToObject(), ßerror.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vink").ToObject(), ßhint.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vigtigt").ToObject(), ßimportant.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("bem\xc3\xa6rk").ToObject(), ßnote.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tips").ToObject(), ßtip.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("advarsel").ToObject(), ßwarning.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("formaning").ToObject(), ßadmonition.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sidebj\xc3\xa6lke").ToObject(), ßsidebar.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("emne").ToObject(), ßtopic.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("linje-blok").ToObject(), πg.NewStr("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("linie-blok").ToObject(), πg.NewStr("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("parset-literal").ToObject(), πg.NewStr("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rubrik").ToObject(), ßrubric.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("epigraf").ToObject(), ßepigraph.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("fremh\xc3\xa6vninger").ToObject(), ßhighlights.ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tabel").ToObject(), ßtable.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("csv-tabel").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("liste-tabel").ToObject(), πg.NewStr("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("meta").ToObject(), ßmeta.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("billede").ToObject(), ßimage.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("figur").ToObject(), ßfigure.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("inklud\xc3\xa9r").ToObject(), ßinclude.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("inkluder").ToObject(), ßinclude.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("r\xc3\xa5").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("erstat").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("unicode").ToObject(), ßunicode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dato").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("klasse").ToObject(), ßclass.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rolle").ToObject(), ßrole.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("forvalgt-rolle").ToObject(), πg.NewStr("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titel").ToObject(), ßtitle.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("indhold").ToObject(), ßcontents.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sektnum").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sektions-nummerering").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sidehovede").ToObject(), ßheader.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sidefod").ToObject(), ßfooter.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("target-notes (translation required)").ToObject(), πg.NewStr("target-notes").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("restructuredtext-test-direktiv").ToObject(), πg.NewStr("restructuredtext-test-directive").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 75: """Danish name to registered (in directives/__init__.py) directive name
			πF.SetLineno(75)
			// line 78: roles = {
			πF.SetLineno(78)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("forkortelse").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("fork").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("akronym").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("ac (translation required)").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kode").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("indeks").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("i").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("subscript (translation required)").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sub (translation required)").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("superscript (translation required)").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sup (translation required)").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titel-reference").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titel").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("t").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pep-reference").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pep").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc-reference").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("emfase").ToObject(), ßemphasis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kraftig").ToObject(), ßstrong.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("literal").ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("navngivet-reference").ToObject(), πg.NewStr("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anonym-reference").ToObject(), πg.NewStr("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("fodnote-reference").ToObject(), πg.NewStr("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("citation-reference (translation required)").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("substitutions-reference").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("target (translation required)").ToObject(), ßtarget.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri-reference").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("url").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("r\xc3\xa5").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 112: """Mapping of Danish role names to canonical role names for interpreted text.
			πF.SetLineno(112)
		}
		return nil, πE
	})
	πg.RegisterModule("da", Code)
}
