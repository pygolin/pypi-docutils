package directives

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/codecs"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAdmonition := πg.InternStr("Admonition")
		ßAttention := πg.InternStr("Attention")
		ßAttributeError := πg.InternStr("AttributeError")
		ßCSVTable := πg.InternStr("CSVTable")
		ßCaution := πg.InternStr("Caution")
		ßClass := πg.InternStr("Class")
		ßCodeBlock := πg.InternStr("CodeBlock")
		ßCompound := πg.InternStr("Compound")
		ßContainer := πg.InternStr("Container")
		ßContents := πg.InternStr("Contents")
		ßDanger := πg.InternStr("Danger")
		ßDate := πg.InternStr("Date")
		ßDefaultRole := πg.InternStr("DefaultRole")
		ßEpigraph := πg.InternStr("Epigraph")
		ßError := πg.InternStr("Error")
		ßFigure := πg.InternStr("Figure")
		ßFooter := πg.InternStr("Footer")
		ßHeader := πg.InternStr("Header")
		ßHighlights := πg.InternStr("Highlights")
		ßHint := πg.InternStr("Hint")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßImage := πg.InternStr("Image")
		ßImportError := πg.InternStr("ImportError")
		ßImportant := πg.InternStr("Important")
		ßInclude := πg.InternStr("Include")
		ßKeyError := πg.InternStr("KeyError")
		ßLineBlock := πg.InternStr("LineBlock")
		ßListTable := πg.InternStr("ListTable")
		ßLookupError := πg.InternStr("LookupError")
		ßMathBlock := πg.InternStr("MathBlock")
		ßMeta := πg.InternStr("Meta")
		ßNone := πg.InternStr("None")
		ßNote := πg.InternStr("Note")
		ßOverflowError := πg.InternStr("OverflowError")
		ßParsedLiteral := πg.InternStr("ParsedLiteral")
		ßPullQuote := πg.InternStr("PullQuote")
		ßRSTTable := πg.InternStr("RSTTable")
		ßRaw := πg.InternStr("Raw")
		ßReplace := πg.InternStr("Replace")
		ßRole := πg.InternStr("Role")
		ßRubric := πg.InternStr("Rubric")
		ßSectnum := πg.InternStr("Sectnum")
		ßSidebar := πg.InternStr("Sidebar")
		ßTargetNotes := πg.InternStr("TargetNotes")
		ßTestDirective := πg.InternStr("TestDirective")
		ßTip := πg.InternStr("Tip")
		ßTitle := πg.InternStr("Title")
		ßTopic := πg.InternStr("Topic")
		ßUnicode := πg.InternStr("Unicode")
		ßValueError := πg.InternStr("ValueError")
		ßWarning := πg.InternStr("Warning")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__import__ := πg.InternStr("__import__")
		ß__name__ := πg.InternStr("__name__")
		ß_directive_registry := πg.InternStr("_directive_registry")
		ß_directives := πg.InternStr("_directives")
		ß_fallback_language_module := πg.InternStr("_fallback_language_module")
		ßadmonition := πg.InternStr("admonition")
		ßadmonitions := πg.InternStr("admonitions")
		ßappend := πg.InternStr("append")
		ßattention := πg.InternStr("attention")
		ßbody := πg.InternStr("body")
		ßcaution := πg.InternStr("caution")
		ßchoice := πg.InternStr("choice")
		ßchr := πg.InternStr("chr")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßcm := πg.InternStr("cm")
		ßcode := πg.InternStr("code")
		ßcodecs := πg.InternStr("codecs")
		ßcompile := πg.InternStr("compile")
		ßcompound := πg.InternStr("compound")
		ßcontainer := πg.InternStr("container")
		ßcontents := πg.InternStr("contents")
		ßcurrent_line := πg.InternStr("current_line")
		ßdanger := πg.InternStr("danger")
		ßdate := πg.InternStr("date")
		ßdirective := πg.InternStr("directive")
		ßdirectives := πg.InternStr("directives")
		ßem := πg.InternStr("em")
		ßencoding := πg.InternStr("encoding")
		ßepigraph := πg.InternStr("epigraph")
		ßerror := πg.InternStr("error")
		ßescape2null := πg.InternStr("escape2null")
		ßex := πg.InternStr("ex")
		ßfigure := πg.InternStr("figure")
		ßflag := πg.InternStr("flag")
		ßfloat := πg.InternStr("float")
		ßfooter := πg.InternStr("footer")
		ßformat_values := πg.InternStr("format_values")
		ßget_measure := πg.InternStr("get_measure")
		ßgetattr := πg.InternStr("getattr")
		ßglobals := πg.InternStr("globals")
		ßgroup := πg.InternStr("group")
		ßheader := πg.InternStr("header")
		ßhighlights := πg.InternStr("highlights")
		ßhint := πg.InternStr("hint")
		ßhtml := πg.InternStr("html")
		ßimage := πg.InternStr("image")
		ßimages := πg.InternStr("images")
		ßimportant := πg.InternStr("important")
		ßin := πg.InternStr("in")
		ßinclude := πg.InternStr("include")
		ßinfo := πg.InternStr("info")
		ßint := πg.InternStr("int")
		ßisdigit := πg.InternStr("isdigit")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlength_or_percentage_or_unitless := πg.InternStr("length_or_percentage_or_unitless")
		ßlength_or_unitless := πg.InternStr("length_or_unitless")
		ßlength_units := πg.InternStr("length_units")
		ßlocals := πg.InternStr("locals")
		ßlookup := πg.InternStr("lookup")
		ßlower := πg.InternStr("lower")
		ßmake_id := πg.InternStr("make_id")
		ßmatch := πg.InternStr("match")
		ßmath := πg.InternStr("math")
		ßmeta := πg.InternStr("meta")
		ßmisc := πg.InternStr("misc")
		ßmm := πg.InternStr("mm")
		ßnodes := πg.InternStr("nodes")
		ßnonnegative_int := πg.InternStr("nonnegative_int")
		ßnote := πg.InternStr("note")
		ßparts := πg.InternStr("parts")
		ßpath := πg.InternStr("path")
		ßpc := πg.InternStr("pc")
		ßpercentage := πg.InternStr("percentage")
		ßpositive_int := πg.InternStr("positive_int")
		ßpositive_int_list := πg.InternStr("positive_int_list")
		ßpt := πg.InternStr("pt")
		ßpx := πg.InternStr("px")
		ßraw := πg.InternStr("raw")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreferences := πg.InternStr("references")
		ßregister_directive := πg.InternStr("register_directive")
		ßreplace := πg.InternStr("replace")
		ßreporter := πg.InternStr("reporter")
		ßrole := πg.InternStr("role")
		ßrstrip := πg.InternStr("rstrip")
		ßrubric := πg.InternStr("rubric")
		ßsectnum := πg.InternStr("sectnum")
		ßsidebar := πg.InternStr("sidebar")
		ßsingle_char_or_unicode := πg.InternStr("single_char_or_unicode")
		ßsingle_char_or_whitespace_or_unicode := πg.InternStr("single_char_or_whitespace_or_unicode")
		ßspace := πg.InternStr("space")
		ßsplit := πg.InternStr("split")
		ßsplit_escaped_whitespace := πg.InternStr("split_escaped_whitespace")
		ßsplitlines := πg.InternStr("splitlines")
		ßstrip := πg.InternStr("strip")
		ßsys := πg.InternStr("sys")
		ßtab := πg.InternStr("tab")
		ßtable := πg.InternStr("table")
		ßtables := πg.InternStr("tables")
		ßtip := πg.InternStr("tip")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßunchanged := πg.InternStr("unchanged")
		ßunchanged_required := πg.InternStr("unchanged_required")
		ßunescape := πg.InternStr("unescape")
		ßunichr := πg.InternStr("unichr")
		ßunicode := πg.InternStr("unicode")
		ßunicode_code := πg.InternStr("unicode_code")
		ßunicode_pattern := πg.InternStr("unicode_pattern")
		ßuri := πg.InternStr("uri")
		ßvalue_or := πg.InternStr("value_or")
		ßversion_info := πg.InternStr("version_info")
		ßwarning := πg.InternStr("warning")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		var πTemp022 *πg.Object
		_ = πTemp022
		var πTemp023 *πg.Object
		_ = πTemp023
		var πTemp024 *πg.Object
		_ = πTemp024
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 *πg.Object
		_ = πTemp026
		var πTemp027 *πg.Object
		_ = πTemp027
		var πTemp028 *πg.Object
		_ = πTemp028
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 5: """
			πF.SetLineno(5)
			// line 5: """
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis package contains directive implementation modules.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 11: import re
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: import codecs
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "codecs"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcodecs.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: import sys
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: from docutils import nodes
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: from docutils.utils import split_escaped_whitespace, escape2null, unescape
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßsplit_escaped_whitespace); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsplit_escaped_whitespace.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßescape2null); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßescape2null.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßunescape); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunescape.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: from docutils.parsers.rst.languages import en as _fallback_language_module
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.languages.en"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[4]
			if πE = πF.Globals().SetItem(πF, ß_fallback_language_module.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 19: if sys.version_info >= (3, 0):
			πF.SetLineno(19)
		Label1:
			// line 20: unichr = chr  # noqa
			πF.SetLineno(20)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunichr.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 23: _directive_registry = {
			πF.SetLineno(23)
			πTemp006 = πg.NewDict()
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßAttention.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßattention.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßCaution.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßcaution.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßCodeBlock.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßcode.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßDanger.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßdanger.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßError.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßerror.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßImportant.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßimportant.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßNote.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßnote.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßTip.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßtip.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßHint.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßhint.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßWarning.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßwarning.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßadmonitions.ToObject(), ßAdmonition.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßadmonition.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßSidebar.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßsidebar.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßTopic.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßtopic.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßLineBlock.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("line-block").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßParsedLiteral.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("parsed-literal").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßMathBlock.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßmath.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßRubric.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßrubric.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßEpigraph.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßepigraph.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßHighlights.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßhighlights.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßPullQuote.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("pull-quote").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßCompound.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßcompound.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßbody.ToObject(), ßContainer.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßcontainer.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßtables.ToObject(), ßRSTTable.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßtable.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßtables.ToObject(), ßCSVTable.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("csv-table").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßtables.ToObject(), ßListTable.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("list-table").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßimages.ToObject(), ßImage.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßimage.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßimages.ToObject(), ßFigure.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßfigure.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßparts.ToObject(), ßContents.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßcontents.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßparts.ToObject(), ßSectnum.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßsectnum.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßparts.ToObject(), ßHeader.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßheader.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßparts.ToObject(), ßFooter.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßfooter.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßreferences.ToObject(), ßTargetNotes.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("target-notes").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßhtml.ToObject(), ßMeta.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßmeta.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßRaw.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßraw.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßInclude.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßinclude.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßReplace.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßreplace.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßUnicode.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßClass.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßclass.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßRole.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßrole.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßDefaultRole.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("default-role").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßTitle.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßtitle.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßDate.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, ßdate.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(ßmisc.ToObject(), ßTestDirective.ToObject()).ToObject()
			if πE = πTemp006.SetItem(πF, πg.NewStr("restructuredtext-test-directive").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_directive_registry.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 71: """Mapping of directive name to (module name, class name).  The
			πF.SetLineno(71)
			// line 75: _directives = {}
			πF.SetLineno(75)
			πTemp006 = πg.NewDict()
			πTemp001 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_directives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 76: """Cache of imported directives."""
			πF.SetLineno(76)
			// line 78: def directive(directive_name, language_module, document):
			πF.SetLineno(78)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "directive_name", Def: nil}
			πTemp007[1] = πg.Param{Name: "language_module", Def: nil}
			πTemp007[2] = πg.Param{Name: "document", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("directive", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdirective_name *πg.Object = πArgs[0]
				_ = µdirective_name
				var µlanguage_module *πg.Object = πArgs[1]
				_ = µlanguage_module
				var µdocument *πg.Object = πArgs[2]
				_ = µdocument
				var µnormname *πg.Object = πg.UnboundLocal
				_ = µnormname
				var µmessages *πg.Object = πg.UnboundLocal
				_ = µmessages
				var µmsg_text *πg.Object = πg.UnboundLocal
				_ = µmsg_text
				var µcanonicalname *πg.Object = πg.UnboundLocal
				_ = µcanonicalname
				var µerror *πg.Object = πg.UnboundLocal
				_ = µerror
				var µmessage *πg.Object = πg.UnboundLocal
				_ = µmessage
				var µmodulename *πg.Object = πg.UnboundLocal
				_ = µmodulename
				var µclassname *πg.Object = πg.UnboundLocal
				_ = µclassname
				var µmodule *πg.Object = πg.UnboundLocal
				_ = µmodule
				var µdetail *πg.Object = πg.UnboundLocal
				_ = µdetail
				var µdirective *πg.Object = πg.UnboundLocal
				_ = µdirective
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 πg.KWArgs
				_ = πTemp010
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 10:
						goto Label10
					case 21:
						goto Label21
					case 4:
						goto Label4
					case 18:
						goto Label18
					case 15:
						goto Label15
					default:
						panic("unexpected function state")
					}
					// line 79: """
					πF.SetLineno(79)
					// line 84: normname = directive_name.lower()
					πF.SetLineno(84)
					if πE = πg.CheckLocal(πF, µdirective_name, "directive_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdirective_name, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnormname = πTemp002
					// line 85: messages = []
					πF.SetLineno(85)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µmessages = πTemp001
					// line 86: msg_text = []
					πF.SetLineno(86)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µmsg_text = πTemp001
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_directives); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp002, µnormname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 87: if normname in _directives:
					πF.SetLineno(87)
				Label1:
					// line 88: return _directives[normname], messages
					πF.SetLineno(88)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp002 = µnormname
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_directives); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp005, µmessages).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					// line 89: canonicalname = None
					πF.SetLineno(89)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µcanonicalname = πTemp001
					// line 90: try:
					πF.SetLineno(90)
					πF.PushCheckpoint(4)
					// line 91: canonicalname = language_module.directives[normname]
					πF.SetLineno(91)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp001 = µnormname
					if πE = πg.CheckLocal(πF, µlanguage_module, "language_module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlanguage_module, ßdirectives, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					µcanonicalname = πTemp002
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 92: except AttributeError as error:
					πF.SetLineno(92)
				Label5:
					µerror = πTemp007.ToObject()
					// line 93: msg_text.append('Problem retrieving directive entry from language '
					πF.SetLineno(93)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage_module, "language_module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µlanguage_module, µerror).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Problem retrieving directive entry from language module %r: %s.").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.RestoreExc(nil, nil)
					goto Label3
					// line 95: except KeyError:
					πF.SetLineno(95)
				Label6:
					// line 96: msg_text.append('No directive entry for "%s" in module "%s".'
					πF.SetLineno(96)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirective_name, "directive_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlanguage_module, "language_module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlanguage_module, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µdirective_name, πTemp005).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("No directive entry for \"%s\" in module \"%s\".").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µcanonicalname, "canonicalname"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcanonicalname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 98: if not canonicalname:
					πF.SetLineno(98)
				Label7:
					// line 99: try:
					πF.SetLineno(99)
					πF.PushCheckpoint(10)
					// line 100: canonicalname = _fallback_language_module.directives[normname]
					πF.SetLineno(100)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp001 = µnormname
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_fallback_language_module); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßdirectives, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					µcanonicalname = πTemp002
					// line 101: msg_text.append('Using English fallback for directive "%s".'
					πF.SetLineno(101)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirective_name, "directive_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Using English fallback for directive \"%s\".").ToObject(), µdirective_name); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 103: except KeyError:
					πF.SetLineno(103)
				Label11:
					// line 104: msg_text.append('Trying "%s" as canonical directive name.'
					πF.SetLineno(104)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdirective_name, "directive_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Trying \"%s\" as canonical directive name.").ToObject(), µdirective_name); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 107: canonicalname = normname
					πF.SetLineno(107)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					µcanonicalname = µnormname
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µmsg_text); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label12
					}
					goto Label13
					// line 108: if msg_text:
					πF.SetLineno(108)
				Label12:
					// line 109: message = document.reporter.info(
					πF.SetLineno(109)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					πTemp009[0] = µmsg_text
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßcurrent_line, nil); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"line", πTemp001},
					}
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µmessage = πTemp001
					// line 111: messages.append(message)
					πF.SetLineno(111)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp003[0] = µmessage
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label13
				Label13:
					// line 112: try:
					πF.SetLineno(112)
					πF.PushCheckpoint(15)
					// line 113: modulename, classname = _directive_registry[canonicalname]
					πF.SetLineno(113)
					if πE = πg.CheckLocal(πF, µcanonicalname, "canonicalname"); πE != nil {
						continue
					}
					πTemp001 = µcanonicalname
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_directive_registry); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µmodulename = πTemp001
					µclassname = πTemp005
					πF.PopCheckpoint()
					goto Label14
				Label15:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label16
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 114: except KeyError:
					πF.SetLineno(114)
				Label16:
					// line 116: return None, messages
					πF.SetLineno(116)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, µmessages).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label14
				Label14:
					// line 117: try:
					πF.SetLineno(117)
					πF.PushCheckpoint(18)
					// line 118: module = __import__(modulename, globals(), locals(), level=1)
					πF.SetLineno(118)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmodulename, "modulename"); πE != nil {
						continue
					}
					πTemp003[0] = µmodulename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[1] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003[2] = πTemp002
					πTemp010 = πg.KWArgs{
						{"level", πg.NewInt(1).ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µmodule = πTemp002
					πF.PopCheckpoint()
					goto Label17
				Label18:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label19
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 119: except ImportError as detail:
					πF.SetLineno(119)
				Label19:
					µdetail = πTemp007.ToObject()
					// line 120: messages.append(document.reporter.error(
					πF.SetLineno(120)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmodulename, "modulename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdirective_name, "directive_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µmodulename, µdirective_name, µdetail).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error importing directive module \"%s\" (directive \"%s\"):\n%s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp009[0] = πTemp001
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßcurrent_line, nil); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"line", πTemp001},
					}
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 124: return None, messages
					πF.SetLineno(124)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, µmessages).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label17
				Label17:
					// line 125: try:
					πF.SetLineno(125)
					πF.PushCheckpoint(21)
					// line 126: directive = getattr(module, classname)
					πF.SetLineno(126)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp003[0] = µmodule
					if πE = πg.CheckLocal(πF, µclassname, "classname"); πE != nil {
						continue
					}
					πTemp003[1] = µclassname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µdirective = πTemp002
					// line 127: _directives[normname] = directive
					πF.SetLineno(127)
					if πE = πg.CheckLocal(πF, µdirective, "directive"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdirective); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_directives); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp005 = µnormname
					if πE = πg.SetItem(πF, πTemp002, πTemp005, πTemp001); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label20
				Label21:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label22
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 128: except AttributeError:
					πF.SetLineno(128)
				Label22:
					// line 129: messages.append(document.reporter.error(
					πF.SetLineno(129)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µclassname, "classname"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmodulename, "modulename"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdirective_name, "directive_name"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µclassname, µmodulename, µdirective_name).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("No directive class \"%s\" in module \"%s\" (directive \"%s\").").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp009[0] = πTemp001
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßcurrent_line, nil); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"line", πTemp001},
					}
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp009, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 133: return None, messages
					πF.SetLineno(133)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, µmessages).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label20
				Label20:
					// line 134: return directive, messages
					πF.SetLineno(134)
					if πE = πg.CheckLocal(πF, µdirective, "directive"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µdirective, µmessages).ToObject()
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirective.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 79: """
			πF.SetLineno(79)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Locate and return a directive function from its language-dependent name.\n    If not found in the current language, check English.  Return None if the\n    named directive cannot be found.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdirective); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 136: def register_directive(name, directive):
			πF.SetLineno(136)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "name", Def: nil}
			πTemp007[1] = πg.Param{Name: "directive", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("register_directive", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
				var µdirective *πg.Object = πArgs[1]
				_ = µdirective
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 137: """
					πF.SetLineno(137)
					// line 141: _directives[name] = directive
					πF.SetLineno(141)
					if πE = πg.CheckLocal(πF, µdirective, "directive"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdirective); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_directives); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003 = µname
					if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
						continue
					}
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßregister_directive.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 137: """
			πF.SetLineno(137)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Register a nonstandard application-defined directive function.\n    Language lookups are not needed for such functions.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßregister_directive); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 143: def flag(argument):
			πF.SetLineno(143)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("flag", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 144: """
					πF.SetLineno(144)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001 = µargument
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µargument, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp004
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 150: if argument and argument.strip():
					πF.SetLineno(150)
				Label2:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("no argument is allowed; \"%s\" supplied").ToObject(), µargument); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 151: raise ValueError('no argument is allowed; "%s" supplied' % argument)
					πF.SetLineno(151)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label4
				Label3:
					// line 153: return None
					πF.SetLineno(153)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label4
				Label4:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßflag.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 144: """
			πF.SetLineno(144)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n    Check for a valid flag option (no argument) and return ``None``.\n    (Directive option conversion function.)\n\n    Raise ``ValueError`` if an argument is found.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßflag); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 155: def unchanged_required(argument):
			πF.SetLineno(155)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("unchanged_required", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 156: """
					πF.SetLineno(156)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µargument == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 162: if argument is None:
					πF.SetLineno(162)
				Label1:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("argument required but none supplied").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 163: raise ValueError('argument required but none supplied')
					πF.SetLineno(163)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label3
				Label2:
					// line 165: return argument  # unchanged!
					πF.SetLineno(165)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πR = µargument
					continue
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßunchanged_required.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 156: """
			πF.SetLineno(156)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Return the argument text, unchanged.\n    (Directive option conversion function.)\n\n    Raise ``ValueError`` if no argument is found.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßunchanged_required); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 167: def unchanged(argument):
			πF.SetLineno(167)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("unchanged", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 168: """
					πF.SetLineno(168)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µargument == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 174: if argument is None:
					πF.SetLineno(174)
				Label1:
					// line 175: return u''
					πF.SetLineno(175)
					πR = πg.NewUnicode("").ToObject()
					continue
					goto Label3
				Label2:
					// line 177: return argument  # unchanged!
					πF.SetLineno(177)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πR = µargument
					continue
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßunchanged.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 168: """
			πF.SetLineno(168)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    Return the argument text, unchanged.\n    (Directive option conversion function.)\n\n    No argument implies empty string (\"\").\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßunchanged); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 179: def path(argument):
			πF.SetLineno(179)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("path", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µpath *πg.Object = πg.UnboundLocal
				_ = µpath
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 180: """
					πF.SetLineno(180)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µargument == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 186: if argument is None:
					πF.SetLineno(186)
				Label1:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("argument required but none supplied").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 187: raise ValueError('argument required but none supplied')
					πF.SetLineno(187)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label3
				Label2:
					// line 189: path = ''.join([s.strip() for s in argument.splitlines()])
					πF.SetLineno(189)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µargument, ßsplitlines, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp004 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp005 = !isStop
								} else {
									πTemp005 = true
									µs = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 189: path = ''.join([s.strip() for s in argument.splitlines()])
								πF.SetLineno(189)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp002 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µpath = πTemp006
					// line 190: return path
					πF.SetLineno(190)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πR = µpath
					continue
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpath.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 180: """
			πF.SetLineno(180)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n    Return the path argument unwrapped (with newlines removed).\n    (Directive option conversion function.)\n\n    Raise ``ValueError`` if no argument is found.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßpath); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 192: def uri(argument):
			πF.SetLineno(192)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("uri", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µparts *πg.Object = πg.UnboundLocal
				_ = µparts
				var µuri *πg.Object = πg.UnboundLocal
				_ = µuri
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 193: """
					πF.SetLineno(193)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µargument == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 199: if argument is None:
					πF.SetLineno(199)
				Label1:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("argument required but none supplied").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 200: raise ValueError('argument required but none supplied')
					πF.SetLineno(200)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label3
				Label2:
					// line 202: parts = split_escaped_whitespace(escape2null(argument))
					πF.SetLineno(202)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp005[0] = µargument
					if πTemp001, πE = πg.ResolveGlobal(πF, ßescape2null); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsplit_escaped_whitespace); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µparts = πTemp002
					// line 203: uri = ' '.join(''.join(unescape(part).split()) for part in parts)
					πF.SetLineno(203)
					πTemp004 = πF.MakeArgs(1)
					πTemp006 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µpart *πg.Object = πg.UnboundLocal
						_ = µpart
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µparts); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µpart = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 203: uri = ' '.join(''.join(unescape(part).split()) for part in parts)
								πF.SetLineno(203)
								πTemp005 = πF.MakeArgs(1)
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µpart, "part"); πE != nil {
									continue
								}
								πTemp006[0] = µpart
								if πTemp004, πE = πg.ResolveGlobal(πF, ßunescape); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp004, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp007
								if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp007, nil
							Label4:
								πTemp004 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µuri = πTemp007
					// line 204: return uri
					πF.SetLineno(204)
					if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
						continue
					}
					πR = µuri
					continue
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuri.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 193: """
			πF.SetLineno(193)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n    Return the URI argument with unescaped whitespace removed.\n    (Directive option conversion function.)\n\n    Raise ``ValueError`` if no argument is found.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßuri); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 206: def nonnegative_int(argument):
			πF.SetLineno(206)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("nonnegative_int", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 207: """
					πF.SetLineno(207)
					// line 211: value = int(argument)
					πF.SetLineno(211)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp003
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µvalue, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 212: if value < 0:
					πF.SetLineno(212)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("negative value; must be positive or zero").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 213: raise ValueError('negative value; must be positive or zero')
					πF.SetLineno(213)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 214: return value
					πF.SetLineno(214)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnonnegative_int.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 207: """
			πF.SetLineno(207)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n    Check for a nonnegative integer argument; raise ``ValueError`` if not.\n    (Directive option conversion function.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßnonnegative_int); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 216: def percentage(argument):
			πF.SetLineno(216)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("percentage", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 217: """
					πF.SetLineno(217)
					// line 220: try:
					πF.SetLineno(220)
					πF.PushCheckpoint(2)
					// line 221: argument = argument.rstrip(' %')
					πF.SetLineno(221)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(" %").ToObject()
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µargument, ßrstrip, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µargument = πTemp003
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 222: except AttributeError:
					πF.SetLineno(222)
				Label3:
					// line 223: pass
					πF.SetLineno(223)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 224: return nonnegative_int(argument)
					πF.SetLineno(224)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnonnegative_int); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpercentage.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 217: """
			πF.SetLineno(217)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n    Check for an integer percentage value with optional percent sign.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßpercentage); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 226: length_units = ['em', 'ex', 'px', 'in', 'cm', 'mm', 'pt', 'pc']
			πF.SetLineno(226)
			πTemp002 = make([]*πg.Object, 8)
			πTemp002[0] = ßem.ToObject()
			πTemp002[1] = ßex.ToObject()
			πTemp002[2] = ßpx.ToObject()
			πTemp002[3] = ßin.ToObject()
			πTemp002[4] = ßcm.ToObject()
			πTemp002[5] = ßmm.ToObject()
			πTemp002[6] = ßpt.ToObject()
			πTemp002[7] = ßpc.ToObject()
			πTemp014 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlength_units.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 228: def get_measure(argument, units):
			πF.SetLineno(228)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp007[1] = πg.Param{Name: "units", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("get_measure", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µunits *πg.Object = πArgs[1]
				_ = µunits
				var µmatch *πg.Object = πg.UnboundLocal
				_ = µmatch
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 229: """
					πF.SetLineno(229)
					// line 236: match = re.match(r'^([0-9.]+) *(%s)$' % '|'.join(units), argument)
					πF.SetLineno(236)
					πTemp001 = πF.MakeArgs(2)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µunits, "units"); πE != nil {
						continue
					}
					πTemp003[0] = µunits
					if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("|").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("^([0-9.]+) *(%s)$").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[1] = µargument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmatch = πTemp002
					// line 237: try:
					πF.SetLineno(237)
					πF.PushCheckpoint(2)
					// line 238: float(match.group(1))
					πF.SetLineno(238)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001[0] = πTemp004
					if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 239: except (AttributeError, ValueError):
					πF.SetLineno(239)
				Label3:
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µunits, "units"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µunits); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µi = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 242: % ' '.join(['"%s"' % i for i in units]))
								πF.SetLineno(242)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Mod(πF, πg.NewStr("\"%s\"").ToObject(), µi); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp005 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp010, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("not a positive measure of one of the following units:\n%s").ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 240: raise ValueError(
					πF.SetLineno(240)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 243: return match.group(1) + match.group(2)
					πF.SetLineno(243)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Add(πF, πTemp010, πTemp011); πE != nil {
						continue
					}
					πR = πTemp002
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßget_measure.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 229: """
			πF.SetLineno(229)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n    Check for a positive argument of one of the units and return a\n    normalized string of the form \"<value><unit>\" (without space in\n    between).\n\n    To be called from directive option conversion functions.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_measure); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 245: def length_or_unitless(argument):
			πF.SetLineno(245)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("length_or_unitless", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 246: return get_measure(argument, length_units + [''])
					πF.SetLineno(246)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlength_units); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = ß.ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßget_measure); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlength_or_unitless.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 248: def length_or_percentage_or_unitless(argument, default=''):
			πF.SetLineno(248)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp007[1] = πg.Param{Name: "default", Def: ß.ToObject()}
			πTemp016 = πg.NewFunction(πg.NewCode("length_or_percentage_or_unitless", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µdefault *πg.Object = πArgs[1]
				_ = µdefault
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					case 5:
						goto Label5
					default:
						panic("unexpected function state")
					}
					// line 249: """
					πF.SetLineno(249)
					// line 264: try:
					πF.SetLineno(264)
					πF.PushCheckpoint(2)
					// line 265: return get_measure(argument, length_units + ['%'])
					πF.SetLineno(265)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlength_units); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = πg.NewStr("%").ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßget_measure); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 266: except ValueError:
					πF.SetLineno(266)
				Label3:
					// line 267: try:
					πF.SetLineno(267)
					πF.PushCheckpoint(5)
					// line 268: return get_measure(argument, ['']) + default
					πF.SetLineno(268)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = ß.ToObject()
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßget_measure); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp005, µdefault); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp009, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label6
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 269: except ValueError:
					πF.SetLineno(269)
				Label6:
					// line 271: return get_measure(argument, length_units + ['%'])
					πF.SetLineno(271)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlength_units); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 1)
					πTemp004[0] = πg.NewStr("%").ToObject()
					πTemp005 = πg.NewList(πTemp004...).ToObject()
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßget_measure); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßlength_or_percentage_or_unitless.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 249: """
			πF.SetLineno(249)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("\n    Return normalized string of a length or percentage unit.\n\n    Add <default> if there is no unit. Raise ValueError if the argument is not\n    a positive measure of one of the valid CSS units (or without unit).\n\n    >>> length_or_percentage_or_unitless('3 pt')\n    '3pt'\n    >>> length_or_percentage_or_unitless('3%', 'em')\n    '3%'\n    >>> length_or_percentage_or_unitless('3')\n    '3'\n    >>> length_or_percentage_or_unitless('3', 'px')\n    '3px'\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßlength_or_percentage_or_unitless); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
			// line 273: def class_option(argument):
			πF.SetLineno(273)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("class_option", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µnames *πg.Object = πg.UnboundLocal
				_ = µnames
				var µclass_names *πg.Object = πg.UnboundLocal
				_ = µclass_names
				var µname *πg.Object = πg.UnboundLocal
				_ = µname
				var µclass_name *πg.Object = πg.UnboundLocal
				_ = µclass_name
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3:
						goto Label3
					case 4:
						goto Label4
					default:
						panic("unexpected function state")
					}
					// line 274: """
					πF.SetLineno(274)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µargument == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 280: if argument is None:
					πF.SetLineno(280)
				Label1:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("argument required but none supplied").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 281: raise ValueError('argument required but none supplied')
					πF.SetLineno(281)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label2
				Label2:
					// line 282: names = argument.split()
					πF.SetLineno(282)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µargument, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnames = πTemp002
					// line 283: class_names = []
					πF.SetLineno(283)
					πTemp004 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µclass_names = πTemp001
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp003 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label5
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp005 = !isStop
					} else {
						πTemp005 = true
						µname = πTemp002
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(3)
					// line 285: class_name = nodes.make_id(name)
					πF.SetLineno(285)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004[0] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßmake_id, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µclass_name = πTemp002
					if πE = πg.CheckLocal(πF, µclass_name, "class_name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µclass_name); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 286: if not class_name:
					πF.SetLineno(286)
				Label6:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("cannot make \"%s\" into a class name").ToObject(), µname); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 287: raise ValueError('cannot make "%s" into a class name' % name)
					πF.SetLineno(287)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label7
				Label7:
					// line 288: class_names.append(class_name)
					πF.SetLineno(288)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µclass_name, "class_name"); πE != nil {
						continue
					}
					πTemp004[0] = µclass_name
					if πE = πg.CheckLocal(πF, µclass_names, "class_names"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µclass_names, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 289: return class_names
					πF.SetLineno(289)
					if πE = πg.CheckLocal(πF, µclass_names, "class_names"); πE != nil {
						continue
					}
					πR = µclass_names
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßclass_option.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 274: """
			πF.SetLineno(274)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("\n    Convert the argument into a list of ID-compatible strings and return it.\n    (Directive option conversion function.)\n\n    Raise ``ValueError`` if no argument is found.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßclass_option); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
				continue
			}
			// line 291: unicode_pattern = re.compile(
			πF.SetLineno(291)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("(?:0x|x|\\\\x|U\\+?|\\\\u)([0-9a-f]+)$|&#x([0-9a-f]+);$").ToObject()
			if πTemp018, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßIGNORECASE, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp019
			if πTemp018, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßunicode_pattern.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 294: def unicode_code(code):
			πF.SetLineno(294)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "code", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("unicode_code", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]
				_ = µcode
				var µmatch *πg.Object = πg.UnboundLocal
				_ = µmatch
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var µdetail *πg.Object = πg.UnboundLocal
				_ = µdetail
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 295: r"""
					πF.SetLineno(295)
					// line 305: try:
					πF.SetLineno(305)
					πF.PushCheckpoint(2)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcode, ßisdigit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 306: if code.isdigit():                  # decimal number
					πF.SetLineno(306)
				Label3:
					// line 307: return unichr(int(code))
					πF.SetLineno(307)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005[0] = µcode
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunichr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					goto Label5
				Label4:
					// line 309: match = unicode_pattern.match(code)
					πF.SetLineno(309)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp004[0] = µcode
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode_pattern); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µmatch = πTemp001
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µmatch); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					goto Label7
					// line 310: if match:                       # hex number
					πF.SetLineno(310)
				Label6:
					// line 311: value = match.group(1) or match.group(2)
					πF.SetLineno(311)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp006
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001 = πTemp006
				Label9:
					µvalue = πTemp001
					// line 312: return unichr(int(value, 16))
					πF.SetLineno(312)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp005[0] = µvalue
					πTemp005[1] = πg.NewInt(16).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunichr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp002
					continue
					goto Label8
				Label7:
					// line 314: return code
					πF.SetLineno(314)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
					goto Label8
				Label8:
					goto Label5
				Label5:
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label10
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 315: except OverflowError as detail:
					πF.SetLineno(315)
				Label10:
					µdetail = πTemp007.ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("code too large (%s)").ToObject(), µdetail); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 316: raise ValueError('code too large (%s)' % detail)
					πF.SetLineno(316)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßunicode_code.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 295: r"""
			πF.SetLineno(295)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("\n    Convert a Unicode character code to a Unicode character.\n    (Directive option conversion function.)\n\n    Codes may be decimal numbers, hexadecimal numbers (prefixed by ``0x``,\n    ``x``, ``\\x``, ``U+``, ``u``, or ``\\u``; e.g. ``U+262E``), or XML-style\n    numeric character entities (e.g. ``&#x262E;``).  Other text remains as-is.\n\n    Raise ValueError for illegal Unicode code values.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßunicode_code); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
				continue
			}
			// line 318: def single_char_or_unicode(argument):
			πF.SetLineno(318)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("single_char_or_unicode", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µchar *πg.Object = πg.UnboundLocal
				_ = µchar
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 319: """
					πF.SetLineno(319)
					// line 323: char = unicode_code(argument)
					πF.SetLineno(323)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode_code); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µchar = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 324: if len(char) > 1:
					πF.SetLineno(324)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%r invalid; must be a single character or a Unicode code").ToObject(), µchar); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 325: raise ValueError('%r invalid; must be a single character or '
					πF.SetLineno(325)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 327: return char
					πF.SetLineno(327)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πR = µchar
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsingle_char_or_unicode.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 319: """
			πF.SetLineno(319)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("\n    A single character is returned as-is.  Unicode characters codes are\n    converted as in `unicode_code`.  (Directive option conversion function.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßsingle_char_or_unicode); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
				continue
			}
			// line 329: def single_char_or_whitespace_or_unicode(argument):
			πF.SetLineno(329)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("single_char_or_whitespace_or_unicode", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µchar *πg.Object = πg.UnboundLocal
				_ = µchar
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 330: """
					πF.SetLineno(330)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µargument, ßtab.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µargument, ßspace.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 334: if argument == 'tab':
					πF.SetLineno(334)
				Label1:
					// line 335: char = '\t'
					πF.SetLineno(335)
					µchar = πg.NewStr("\t").ToObject()
					goto Label4
					// line 336: elif argument == 'space':
					πF.SetLineno(336)
				Label2:
					// line 337: char = ' '
					πF.SetLineno(337)
					µchar = πg.NewStr(" ").ToObject()
					goto Label4
				Label3:
					// line 339: char = single_char_or_unicode(argument)
					πF.SetLineno(339)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp003[0] = µargument
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsingle_char_or_unicode); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µchar = πTemp004
					goto Label4
				Label4:
					// line 340: return char
					πF.SetLineno(340)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πR = µchar
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsingle_char_or_whitespace_or_unicode.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 330: """
			πF.SetLineno(330)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("\n    As with `single_char_or_unicode`, but \"tab\" and \"space\" are also supported.\n    (Directive option conversion function.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßsingle_char_or_whitespace_or_unicode); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
				continue
			}
			// line 342: def positive_int(argument):
			πF.SetLineno(342)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("positive_int", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 343: """
					πF.SetLineno(343)
					// line 347: value = int(argument)
					πF.SetLineno(347)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp003
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µvalue, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 348: if value < 1:
					πF.SetLineno(348)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("negative or zero value; must be positive").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 349: raise ValueError('negative or zero value; must be positive')
					πF.SetLineno(349)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 350: return value
					πF.SetLineno(350)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpositive_int.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 343: """
			πF.SetLineno(343)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("\n    Converts the argument into an integer.  Raises ValueError for negative,\n    zero, or non-integer values.  (Directive option conversion function.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßpositive_int); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
				continue
			}
			// line 352: def positive_int_list(argument):
			πF.SetLineno(352)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("positive_int_list", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µentries *πg.Object = πg.UnboundLocal
				_ = µentries
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 353: """
					πF.SetLineno(353)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µargument, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 360: if ',' in argument:
					πF.SetLineno(360)
				Label1:
					// line 361: entries = argument.split(',')
					πF.SetLineno(361)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(",").ToObject()
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µargument, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µentries = πTemp004
					goto Label3
				Label2:
					// line 363: entries = argument.split()
					πF.SetLineno(363)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µargument, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µentries = πTemp004
					goto Label3
				Label3:
					// line 364: return [positive_int(entry) for entry in entries]
					πF.SetLineno(364)
					πTemp005 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µentry *πg.Object = πg.UnboundLocal
						_ = µentry
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µentries); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µentry = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 364: return [positive_int(entry) for entry in entries]
								πF.SetLineno(364)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
									continue
								}
								πTemp005[0] = µentry
								if πTemp004, πE = πg.ResolveGlobal(πF, ßpositive_int); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp006, nil
							Label4:
								πTemp004 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpositive_int_list.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 353: """
			πF.SetLineno(353)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("\n    Converts a space- or comma-separated list of values into a Python list\n    of integers.\n    (Directive option conversion function.)\n\n    Raises ValueError for non-positive-integer values.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßpositive_int_list); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
				continue
			}
			// line 366: def encoding(argument):
			πF.SetLineno(366)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp023 = πg.NewFunction(πg.NewCode("encoding", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 367: """
					πF.SetLineno(367)
					// line 373: try:
					πF.SetLineno(373)
					πF.PushCheckpoint(2)
					// line 374: codecs.lookup(argument)
					πF.SetLineno(374)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlookup, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 375: except LookupError:
					πF.SetLineno(375)
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown encoding: \"%s\"").ToObject(), µargument); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 376: raise ValueError('unknown encoding: "%s"' % argument)
					πF.SetLineno(376)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 377: return argument
					πF.SetLineno(377)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πR = µargument
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßencoding.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 367: """
			πF.SetLineno(367)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("\n    Verfies the encoding argument by lookup.\n    (Directive option conversion function.)\n\n    Raises ValueError for unknown encodings.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßencoding); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
				continue
			}
			// line 379: def choice(argument, values):
			πF.SetLineno(379)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "argument", Def: nil}
			πTemp007[1] = πg.Param{Name: "values", Def: nil}
			πTemp024 = πg.NewFunction(πg.NewCode("choice", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
				var µvalues *πg.Object = πArgs[1]
				_ = µvalues
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.BaseException
				_ = πTemp003
				var πTemp004 *πg.Traceback
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 380: """
					πF.SetLineno(380)
					// line 394: try:
					πF.SetLineno(394)
					πF.PushCheckpoint(2)
					// line 395: value = argument.lower().strip()
					πF.SetLineno(395)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µargument, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µvalue = πTemp002
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp003, πTemp004 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
					continue
					// line 396: except AttributeError:
					πF.SetLineno(396)
				Label3:
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					πTemp007[0] = µvalues
					if πTemp002, πE = πg.ResolveGlobal(πF, ßformat_values); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("must supply an argument; choose from %s").ToObject(), πTemp008); πE != nil {
						continue
					}
					πTemp006[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 397: raise ValueError('must supply an argument; choose from %s'
					πF.SetLineno(397)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µvalues, µvalue); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 399: if value in values:
					πF.SetLineno(399)
				Label4:
					// line 400: return value
					πF.SetLineno(400)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
					goto Label6
				Label5:
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					πTemp007[0] = µvalues
					if πTemp008, πE = πg.ResolveGlobal(πF, ßformat_values); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp002 = πg.NewTuple2(µargument, πTemp009).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" unknown; choose from %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp006[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 402: raise ValueError('"%s" unknown; choose from %s'
					πF.SetLineno(402)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label6
				Label6:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßchoice.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 380: """
			πF.SetLineno(380)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("\n    Directive option utility function, supplied to enable options whose\n    argument must be a member of a finite set of possible values (must be\n    lower case).  A custom conversion function must be written to use it.  For\n    example::\n\n        from docutils.parsers.rst import directives\n\n        def yesno(argument):\n            return directives.choice(argument, ('yes', 'no'))\n\n    Raise ``ValueError`` if no argument is found or if the argument's value is\n    not valid (not an entry in the supplied list).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp026, πE = πg.ResolveGlobal(πF, ßchoice); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
				continue
			}
			// line 405: def format_values(values):
			πF.SetLineno(405)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "values", Def: nil}
			πTemp025 = πg.NewFunction(πg.NewCode("format_values", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalues *πg.Object = πArgs[0]
				_ = µvalues
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 406: return '%s, or "%s"' % (', '.join(['"%s"' % s for s in values[:-1]]),
					πF.SetLineno(406)
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µvalues, πTemp002); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp004 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp005 = !isStop
								} else {
									πTemp005 = true
									µs = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 406: return '%s, or "%s"' % (', '.join(['"%s"' % s for s in values[:-1]]),
								πF.SetLineno(406)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Mod(πF, πg.NewStr("\"%s\"").ToObject(), µs); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp003 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp008
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µvalues, πTemp004); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s, or \"%s\"").ToObject(), πTemp002); πE != nil {
						continue
					}
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßformat_values.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 409: def value_or(values, other):
			πF.SetLineno(409)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "values", Def: nil}
			πTemp007[1] = πg.Param{Name: "other", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("value_or", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalues *πg.Object = πArgs[0]
				_ = µvalues
				var µother *πg.Object = πArgs[1]
				_ = µother
				var µauto_or_other *πg.Object = πg.UnboundLocal
				_ = µauto_or_other
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 410: """
					πF.SetLineno(410)
					// line 413: def auto_or_other(argument):
					πF.SetLineno(413)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "argument", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("auto_or_other", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargument *πg.Object = πArgs[0]
						_ = µargument
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µvalues, µargument); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 414: if argument in values:
							πF.SetLineno(414)
						Label1:
							// line 415: return argument
							πF.SetLineno(415)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πR = µargument
							continue
							goto Label3
						Label2:
							// line 417: return other(argument)
							πF.SetLineno(417)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp003[0] = µargument
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = µother.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp001
							continue
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µauto_or_other = πTemp001
					// line 418: return auto_or_other
					πF.SetLineno(418)
					if πE = πg.CheckLocal(πF, µauto_or_other, "auto_or_other"); πE != nil {
						continue
					}
					πR = µauto_or_other
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalue_or.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 410: """
			πF.SetLineno(410)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("\n    The argument can be any of `values` or `argument_type`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßvalue_or); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("directives", Code)
}
