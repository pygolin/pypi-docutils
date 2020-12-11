package eo

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/eo.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nEsperanto-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
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
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("atentu").ToObject(), ßattention.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("zorgu").ToObject(), ßcaution.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code (translation required)").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dangxero").ToObject(), ßdanger.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dan\xc4\x9dero").ToObject(), ßdanger.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("eraro").ToObject(), ßerror.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("spuro").ToObject(), ßhint.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("grava").ToObject(), ßimportant.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("noto").ToObject(), ßnote.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("helpeto").ToObject(), ßtip.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("averto").ToObject(), ßwarning.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("admono").ToObject(), ßadmonition.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("flankteksto").ToObject(), ßsidebar.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("temo").ToObject(), ßtopic.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("linea-bloko").ToObject(), πg.NewStr("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("analizota-literalo").ToObject(), πg.NewStr("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rubriko").ToObject(), ßrubric.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("epigrafo").ToObject(), ßepigraph.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("elstarajxoj").ToObject(), ßhighlights.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("elstara\xc4\xb5oj").ToObject(), ßhighlights.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("ekstera-citajxo").ToObject(), πg.NewStr("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("ekstera-cita\xc4\xb5o").ToObject(), πg.NewStr("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kombinajxo").ToObject(), ßcompound.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kombina\xc4\xb5o").ToObject(), ßcompound.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tekstingo").ToObject(), ßcontainer.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("enhavilo").ToObject(), ßcontainer.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tabelo").ToObject(), ßtable.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tabelo-vdk").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tabelo-csv").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("tabelo-lista").ToObject(), πg.NewStr("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("meta").ToObject(), ßmeta.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("bildo").ToObject(), ßimage.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("figuro").ToObject(), ßfigure.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("inkludi").ToObject(), ßinclude.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("senanaliza").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anstatauxi").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anstata\xc5\xadi").ToObject(), ßreplace.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("unicode").ToObject(), ßunicode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("dato").ToObject(), ßdate.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("klaso").ToObject(), ßclass.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rolo").ToObject(), ßrole.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("preterlasita-rolo").ToObject(), πg.NewStr("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titolo").ToObject(), ßtitle.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("enhavo").ToObject(), ßcontents.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("seknum").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sekcia-numerado").ToObject(), ßsectnum.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kapsekcio").ToObject(), ßheader.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("piedsekcio").ToObject(), ßfooter.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("celaj-notoj").ToObject(), πg.NewStr("target-notes").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("restructuredtext-test-directive").ToObject(), πg.NewStr("restructuredtext-test-directive").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 77: """Esperanto name to registered (in directives/__init__.py) directive name
			πF.SetLineno(77)
			// line 80: roles = {
			πF.SetLineno(80)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("mallongigo").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("mall").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("komenclitero").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("kl").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("code (translation required)").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("indekso").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("i").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("subskribo").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sub").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("supraskribo").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("sup").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titola-referenco").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("titolo").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("t").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pep-referenco").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("pep").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc-referenco").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("emfazo").ToObject(), ßemphasis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("forta").ToObject(), ßstrong.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("litera").ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("math (translation required)").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("nomita-referenco").ToObject(), πg.NewStr("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("nenomita-referenco").ToObject(), πg.NewStr("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("piednota-referenco").ToObject(), πg.NewStr("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("citajxo-referenco").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("cita\xc4\xb5o-referenco").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anstatauxa-referenco").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("anstata\xc5\xada-referenco").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("celo").ToObject(), ßtarget.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri-referenco").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("url").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("senanaliza").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 117: """Mapping of Esperanto role names to canonical role names for interpreted text.
			πF.SetLineno(117)
		}
		return nil, πE
	})
	πg.RegisterModule("eo", Code)
}
