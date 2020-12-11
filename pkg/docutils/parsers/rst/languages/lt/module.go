package lt

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/lt.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßa := πg.InternStr("a")
		ßabbreviation := πg.InternStr("abbreviation")
		ßacronym := πg.InternStr("acronym")
		ßadmonition := πg.InternStr("admonition")
		ßak := πg.InternStr("ak")
		ßakronimas := πg.InternStr("akronimas")
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
		ßi := πg.InternStr("i")
		ßimage := πg.InternStr("image")
		ßimportant := πg.InternStr("important")
		ßinclude := πg.InternStr("include")
		ßindeksas := πg.InternStr("indeksas")
		ßindex := πg.InternStr("index")
		ßliteral := πg.InternStr("literal")
		ßmath := πg.InternStr("math")
		ßmeta := πg.InternStr("meta")
		ßnote := πg.InternStr("note")
		ßpep := πg.InternStr("pep")
		ßraw := πg.InternStr("raw")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreplace := πg.InternStr("replace")
		ßrfc := πg.InternStr("rfc")
		ßrole := πg.InternStr("role")
		ßroles := πg.InternStr("roles")
		ßrubric := πg.InternStr("rubric")
		ßsa := πg.InternStr("sa")
		ßsantrumpa := πg.InternStr("santrumpa")
		ßsectnum := πg.InternStr("sectnum")
		ßsidebar := πg.InternStr("sidebar")
		ßstrong := πg.InternStr("strong")
		ßsub := πg.InternStr("sub")
		ßsubscript := πg.InternStr("subscript")
		ßsup := πg.InternStr("sup")
		ßsuperscript := πg.InternStr("superscript")
		ßtable := πg.InternStr("table")
		ßtarget := πg.InternStr("target")
		ßtip := πg.InternStr("tip")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßunicode := πg.InternStr("unicode")
		ßuri := πg.InternStr("uri")
		ßurl := πg.InternStr("url")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nLithuanian-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("d\xc4\x97mesio").ToObject(), ßattention.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("atsargiai").ToObject(), ßcaution.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code (translation required)").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pavojinga").ToObject(), ßdanger.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("klaida").ToObject(), ßerror.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("u\xc5\xbeuomina").ToObject(), ßhint.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("svarbu").ToObject(), ßimportant.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pastaba").ToObject(), ßnote.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("patarimas").ToObject(), ßtip.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xc4\xafsp\xc4\x97jimas").ToObject(), ßwarning.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("persp\xc4\x97jimas").ToObject(), ßadmonition.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xc5\xa1onin\xc4\x97-juosta").ToObject(), ßsidebar.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tema").ToObject(), ßtopic.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("linijinis-blokas").ToObject(), πg.NewStr("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("i\xc5\xa1analizuotas-literalas").ToObject(), πg.NewStr("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rubrika").ToObject(), ßrubric.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("epigrafas").ToObject(), ßepigraph.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pagridiniai-momentai").ToObject(), ßhighlights.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("atitraukta-citata").ToObject(), πg.NewStr("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sud\xc4\x97tinis-darinys").ToObject(), ßcompound.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("konteineris").ToObject(), ßcontainer.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("lentel\xc4\x97").ToObject(), ßtable.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("csv-lentel\xc4\x97").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("s\xc4\x85ra\xc5\xa1o-lentel\xc4\x97").ToObject(), πg.NewStr("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("meta").ToObject(), ßmeta.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("matematika").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("paveiksliukas").ToObject(), ßimage.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("iliustracija").ToObject(), ßfigure.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("prid\xc4\x97ti").ToObject(), ßinclude.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xc5\xbealia").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pakeisti").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("unikodas").ToObject(), ßunicode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("data").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("klas\xc4\x97").ToObject(), ßclass.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rol\xc4\x97").ToObject(), ßrole.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("numatytoji-rol\xc4\x97").ToObject(), πg.NewStr("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titulas").ToObject(), ßtitle.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("turinys").ToObject(), ßcontents.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("seknum").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sekcijos-numeravimas").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("antra\xc5\xa1t\xc4\x97").ToObject(), ßheader.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pora\xc5\xa1t\xc4\x97").ToObject(), ßfooter.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("nutaikytos-pastaba").ToObject(), πg.NewStr("target-notes").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("restructuredtext-testin\xc4\x97-direktyva").ToObject(), πg.NewStr("restructuredtext-test-directive").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 71: """Lithuanian name to registered (in directives/__init__.py) directive name
			πF.SetLineno(71)
			// line 74: roles = {
			πF.SetLineno(74)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßsantrumpa.ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsa.ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßakronimas.ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßak.ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code (translation required)").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßindeksas.ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßi.ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("apatinis-indeksas").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsub.ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vir\xc5\xa1utinis-indeksas").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsup.ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("antra\xc5\xa1\xc4\x97s-nuoroda").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("antra\xc5\xa1t\xc4\x97").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("pep-nuoroda").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpep.ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("rfc-nuoroda").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrfc.ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pary\xc5\xa1kinimas").ToObject(), ßemphasis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sustiprintas").ToObject(), ßstrong.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("literalas").ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("matematika").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("vardin\xc4\x97-nuoroda").ToObject(), πg.NewStr("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anonimin\xc4\x97-nuoroda").ToObject(), πg.NewStr("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("i\xc5\xa1na\xc5\xa1os-nuoroda").ToObject(), πg.NewStr("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("citatos-nuoroda").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pakeitimo-nuoroda").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("taikinys").ToObject(), ßtarget.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri-nuoroda").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßuri.ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßurl.ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("\xc5\xbealia").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 108: """Mapping of English role names to canonical role names for interpreted text.
			πF.SetLineno(108)
		}
		return nil, πE
	})
	πg.RegisterModule("lt", Code)
}
