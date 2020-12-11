package html4css1

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß1 := πg.InternStr("1")
		ßElement := πg.InternStr("Element")
		ßFalse := πg.InternStr("False")
		ßHTMLTranslator := πg.InternStr("HTMLTranslator")
		ßInvisible := πg.InternStr("Invisible")
		ßNodeFound := πg.InternStr("NodeFound")
		ßNone := πg.InternStr("None")
		ßSimpleListChecker := πg.InternStr("SimpleListChecker")
		ßSkipNode := πg.InternStr("SkipNode")
		ßTrue := πg.InternStr("True")
		ßWriter := πg.InternStr("Writer")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__file__ := πg.InternStr("__file__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_html_base := πg.InternStr("_html_base")
		ßa := πg.InternStr("a")
		ßabspath := πg.InternStr("abspath")
		ßaction := πg.InternStr("action")
		ßadd_meta := πg.InternStr("add_meta")
		ßaddress := πg.InternStr("address")
		ßadmonition := πg.InternStr("admonition")
		ßalign := πg.InternStr("align")
		ßappend := πg.InternStr("append")
		ßastext := πg.InternStr("astext")
		ßattlist := πg.InternStr("attlist")
		ßattributes := πg.InternStr("attributes")
		ßattribution_formats := πg.InternStr("attribution_formats")
		ßattval := πg.InternStr("attval")
		ßauthor := πg.InternStr("author")
		ßauthor_in_authors := πg.InternStr("author_in_authors")
		ßauthors := πg.InternStr("authors")
		ßbackrefs := πg.InternStr("backrefs")
		ßbody := πg.InternStr("body")
		ßbody_pre_docinfo := πg.InternStr("body_pre_docinfo")
		ßborder := πg.InternStr("border")
		ßbottom := πg.InternStr("bottom")
		ßbrackets := πg.InternStr("brackets")
		ßbullet_list := πg.InternStr("bullet_list")
		ßcheck_simple_list := πg.InternStr("check_simple_list")
		ßchildren := πg.InternStr("children")
		ßchoices := πg.InternStr("choices")
		ßclass := πg.InternStr("class")
		ßclasses := πg.InternStr("classes")
		ßclassifier := πg.InternStr("classifier")
		ßcode := πg.InternStr("code")
		ßcol := πg.InternStr("col")
		ßcolgroup := πg.InternStr("colgroup")
		ßcolspan := πg.InternStr("colspan")
		ßcolspec := πg.InternStr("colspec")
		ßcolspecs := πg.InternStr("colspecs")
		ßcolwidth := πg.InternStr("colwidth")
		ßcompact := πg.InternStr("compact")
		ßcompact_field_list := πg.InternStr("compact_field_list")
		ßcompact_field_lists := πg.InternStr("compact_field_lists")
		ßcompact_lists := πg.InternStr("compact_lists")
		ßcompact_p := πg.InternStr("compact_p")
		ßcompact_simple := πg.InternStr("compact_simple")
		ßcompound := πg.InternStr("compound")
		ßconfig_section := πg.InternStr("config_section")
		ßcontent_type := πg.InternStr("content_type")
		ßcontent_type_mathml := πg.InternStr("content_type_mathml")
		ßcontents := πg.InternStr("contents")
		ßcontext := πg.InternStr("context")
		ßdash := πg.InternStr("dash")
		ßdd := πg.InternStr("dd")
		ßdefault := πg.InternStr("default")
		ßdefault_stylesheet_dirs := πg.InternStr("default_stylesheet_dirs")
		ßdefault_stylesheets := πg.InternStr("default_stylesheets")
		ßdefault_template := πg.InternStr("default_template")
		ßdefault_template_path := πg.InternStr("default_template_path")
		ßdepart_author := πg.InternStr("depart_author")
		ßdepart_authors := πg.InternStr("depart_authors")
		ßdepart_citation := πg.InternStr("depart_citation")
		ßdepart_colspec := πg.InternStr("depart_colspec")
		ßdepart_description := πg.InternStr("depart_description")
		ßdepart_docinfo := πg.InternStr("depart_docinfo")
		ßdepart_docinfo_item := πg.InternStr("depart_docinfo_item")
		ßdepart_enumerated_list := πg.InternStr("depart_enumerated_list")
		ßdepart_field := πg.InternStr("depart_field")
		ßdepart_field_body := πg.InternStr("depart_field_body")
		ßdepart_field_list := πg.InternStr("depart_field_list")
		ßdepart_field_name := πg.InternStr("depart_field_name")
		ßdepart_footnote := πg.InternStr("depart_footnote")
		ßdepart_footnote_reference := πg.InternStr("depart_footnote_reference")
		ßdepart_label := πg.InternStr("depart_label")
		ßdepart_literal_block := πg.InternStr("depart_literal_block")
		ßdepart_option_group := πg.InternStr("depart_option_group")
		ßdepart_option_list := πg.InternStr("depart_option_list")
		ßdepart_option_list_item := πg.InternStr("depart_option_list_item")
		ßdepart_paragraph := πg.InternStr("depart_paragraph")
		ßdepart_subscript := πg.InternStr("depart_subscript")
		ßdepart_subtitle := πg.InternStr("depart_subtitle")
		ßdepart_superscript := πg.InternStr("depart_superscript")
		ßdepart_table := πg.InternStr("depart_table")
		ßdepart_tbody := πg.InternStr("depart_tbody")
		ßdepart_thead := πg.InternStr("depart_thead")
		ßdest := πg.InternStr("dest")
		ßdict := πg.InternStr("dict")
		ßdirname := πg.InternStr("dirname")
		ßdiv := πg.InternStr("div")
		ßdl := πg.InternStr("dl")
		ßdocinfo := πg.InternStr("docinfo")
		ßdoctype := πg.InternStr("doctype")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßembed_stylesheet := πg.InternStr("embed_stylesheet")
		ßemptytag := πg.InternStr("emptytag")
		ßencode := πg.InternStr("encode")
		ßenumerate := πg.InternStr("enumerate")
		ßenumerated_list := πg.InternStr("enumerated_list")
		ßenumtype := πg.InternStr("enumtype")
		ßextend := πg.InternStr("extend")
		ßfield := πg.InternStr("field")
		ßfield_body := πg.InternStr("field_body")
		ßfield_name_limit := πg.InternStr("field_name_limit")
		ßfindall := πg.InternStr("findall")
		ßfirst := πg.InternStr("first")
		ßfootnote_backlinks := πg.InternStr("footnote_backlinks")
		ßfootnote_backrefs := πg.InternStr("footnote_backrefs")
		ßfootnote_references := πg.InternStr("footnote_references")
		ßfrontend := πg.InternStr("frontend")
		ßget := πg.InternStr("get")
		ßh2 := πg.InternStr("h2")
		ßhasattr := πg.InternStr("hasattr")
		ßhtml := πg.InternStr("html")
		ßhtml4 := πg.InternStr("html4")
		ßhtml4css1 := πg.InternStr("html4css1")
		ßhtml5_polyglot := πg.InternStr("html5_polyglot")
		ßhtml_subtitle := πg.InternStr("html_subtitle")
		ßin_docinfo := πg.InternStr("in_docinfo")
		ßin_document_title := πg.InternStr("in_document_title")
		ßin_sidebar := πg.InternStr("in_sidebar")
		ßin_word_wrap_point := πg.InternStr("in_word_wrap_point")
		ßindex := πg.InternStr("index")
		ßinitial_header_level := πg.InternStr("initial_header_level")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßio := πg.InternStr("io")
		ßis_compactable := πg.InternStr("is_compactable")
		ßis_not_default := πg.InternStr("is_not_default")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlabel := πg.InternStr("label")
		ßlabels := πg.InternStr("labels")
		ßlanguage := πg.InternStr("language")
		ßlast := πg.InternStr("last")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßli := πg.InternStr("li")
		ßline := πg.InternStr("line")
		ßline_block := πg.InternStr("line_block")
		ßliteral_block := πg.InternStr("literal_block")
		ßmetavar := πg.InternStr("metavar")
		ßnext_node := πg.InternStr("next_node")
		ßnodes := πg.InternStr("nodes")
		ßnone := πg.InternStr("none")
		ßobject_image_types := πg.InternStr("object_image_types")
		ßol := πg.InternStr("ol")
		ßopen := πg.InternStr("open")
		ßoption_limit := πg.InternStr("option_limit")
		ßos := πg.InternStr("os")
		ßoverrides := πg.InternStr("overrides")
		ßp := πg.InternStr("p")
		ßparagraph := πg.InternStr("paragraph")
		ßparens := πg.InternStr("parens")
		ßparent := πg.InternStr("parent")
		ßparentheses := πg.InternStr("parentheses")
		ßparts := πg.InternStr("parts")
		ßpath := πg.InternStr("path")
		ßpop := πg.InternStr("pop")
		ßpre := πg.InternStr("pre")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrefid := πg.InternStr("refid")
		ßsearch := πg.InternStr("search")
		ßsection := πg.InternStr("section")
		ßsection_level := πg.InternStr("section_level")
		ßset_class_on_child := πg.InternStr("set_class_on_child")
		ßset_first_last := πg.InternStr("set_first_last")
		ßsettings := πg.InternStr("settings")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßshould_be_compact_paragraph := πg.InternStr("should_be_compact_paragraph")
		ßsidebar := πg.InternStr("sidebar")
		ßsource := πg.InternStr("source")
		ßspan := πg.InternStr("span")
		ßspecial_characters := πg.InternStr("special_characters")
		ßsplit := πg.InternStr("split")
		ßstart := πg.InternStr("start")
		ßstarttag := πg.InternStr("starttag")
		ßstore_false := πg.InternStr("store_false")
		ßstore_true := πg.InternStr("store_true")
		ßstrip := πg.InternStr("strip")
		ßstub := πg.InternStr("stub")
		ßstubs := πg.InternStr("stubs")
		ßstyle := πg.InternStr("style")
		ßstylesheet := πg.InternStr("stylesheet")
		ßstylesheet_path := πg.InternStr("stylesheet_path")
		ßsub := πg.InternStr("sub")
		ßsubscript := πg.InternStr("subscript")
		ßsubtitle := πg.InternStr("subtitle")
		ßsum := πg.InternStr("sum")
		ßsup := πg.InternStr("sup")
		ßsuperscript := πg.InternStr("superscript")
		ßsupported := πg.InternStr("supported")
		ßtable := πg.InternStr("table")
		ßtable_style := πg.InternStr("table_style")
		ßtbody := πg.InternStr("tbody")
		ßtd := πg.InternStr("td")
		ßth := πg.InternStr("th")
		ßthead := πg.InternStr("thead")
		ßtop := πg.InternStr("top")
		ßtopic_classes := πg.InternStr("topic_classes")
		ßtr := πg.InternStr("tr")
		ßtranslator_class := πg.InternStr("translator_class")
		ßtrim_footnote_reference_space := πg.InternStr("trim_footnote_reference_space")
		ßtt := πg.InternStr("tt")
		ßtype := πg.InternStr("type")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidate_comma_separated_list := πg.InternStr("validate_comma_separated_list")
		ßvalidate_nonnegative_int := πg.InternStr("validate_nonnegative_int")
		ßvalidator := πg.InternStr("validator")
		ßvisit_address := πg.InternStr("visit_address")
		ßvisit_admonition := πg.InternStr("visit_admonition")
		ßvisit_author := πg.InternStr("visit_author")
		ßvisit_authors := πg.InternStr("visit_authors")
		ßvisit_citation := πg.InternStr("visit_citation")
		ßvisit_classifier := πg.InternStr("visit_classifier")
		ßvisit_colspec := πg.InternStr("visit_colspec")
		ßvisit_definition := πg.InternStr("visit_definition")
		ßvisit_definition_list := πg.InternStr("visit_definition_list")
		ßvisit_description := πg.InternStr("visit_description")
		ßvisit_docinfo := πg.InternStr("visit_docinfo")
		ßvisit_docinfo_item := πg.InternStr("visit_docinfo_item")
		ßvisit_doctest_block := πg.InternStr("visit_doctest_block")
		ßvisit_entry := πg.InternStr("visit_entry")
		ßvisit_enumerated_list := πg.InternStr("visit_enumerated_list")
		ßvisit_field := πg.InternStr("visit_field")
		ßvisit_field_body := πg.InternStr("visit_field_body")
		ßvisit_field_list := πg.InternStr("visit_field_list")
		ßvisit_field_name := πg.InternStr("visit_field_name")
		ßvisit_footnote := πg.InternStr("visit_footnote")
		ßvisit_footnote_reference := πg.InternStr("visit_footnote_reference")
		ßvisit_generated := πg.InternStr("visit_generated")
		ßvisit_label := πg.InternStr("visit_label")
		ßvisit_list_item := πg.InternStr("visit_list_item")
		ßvisit_literal := πg.InternStr("visit_literal")
		ßvisit_literal_block := πg.InternStr("visit_literal_block")
		ßvisit_option_group := πg.InternStr("visit_option_group")
		ßvisit_option_list := πg.InternStr("visit_option_list")
		ßvisit_option_list_item := πg.InternStr("visit_option_list_item")
		ßvisit_paragraph := πg.InternStr("visit_paragraph")
		ßvisit_sidebar := πg.InternStr("visit_sidebar")
		ßvisit_subscript := πg.InternStr("visit_subscript")
		ßvisit_subtitle := πg.InternStr("visit_subtitle")
		ßvisit_superscript := πg.InternStr("visit_superscript")
		ßvisit_system_message := πg.InternStr("visit_system_message")
		ßvisit_table := πg.InternStr("visit_table")
		ßvisit_tbody := πg.InternStr("visit_tbody")
		ßvisit_thead := πg.InternStr("visit_thead")
		ßvoid := πg.InternStr("void")
		ßwidth := πg.InternStr("width")
		ßwords_and_spaces := πg.InternStr("words_and_spaces")
		ßwriter_aux := πg.InternStr("writer_aux")
		ßwriters := πg.InternStr("writers")
		ßxhtml := πg.InternStr("xhtml")
		ßxhtml10 := πg.InternStr("xhtml10")
		ßxml_declaration := πg.InternStr("xml_declaration")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 6: """
			πF.SetLineno(6)
			// line 6: """
			πF.SetLineno(6)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nSimple HyperText Markup Language document tree Writer.\n\nThe output conforms to the XHTML version 1.0 Transitional DTD\n(*almost* strict).  The output contains a minimum of formatting\ninformation.  The cascading style sheet \"html4css1.css\" is required\nfor proper viewing with a modern graphical browser.\n").ToObject()); πE != nil {
				continue
			}
			// line 15: __docformat__ = 'reStructuredText'
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 17: import os.path
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import docutils
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: from docutils import frontend, nodes, writers, io
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßfrontend.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßwriters.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: from docutils.transforms import writer_aux
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.writer_aux"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßwriter_aux.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: from docutils.writers import _html_base
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers._html_base"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ß_html_base.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: class Writer(writers._html_base.Writer):
			πF.SetLineno(23)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß_html_base, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Dict
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 []πg.Param
				_ = πTemp015
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 25: supported = ('html', 'html4', 'html4css1', 'xhtml', 'xhtml10')
					πF.SetLineno(25)
					πTemp001 = πg.NewTuple5(ßhtml.ToObject(), ßhtml4.ToObject(), ßhtml4css1.ToObject(), ßxhtml.ToObject(), ßxhtml10.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 26: """Formats this writer supports."""
					πF.SetLineno(26)
					// line 26: """Formats this writer supports."""
					πF.SetLineno(26)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Formats this writer supports.").ToObject()); πE != nil {
						continue
					}
					// line 28: default_stylesheets = ['html4css1.css']
					πF.SetLineno(28)
					πTemp002 = make([]*πg.Object, 1)
					πTemp002[0] = πg.NewStr("html4css1.css").ToObject()
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßdefault_stylesheets.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 29: default_stylesheet_dirs = ['.',
					πF.SetLineno(29)
					πTemp002 = make([]*πg.Object, 3)
					πTemp002[0] = πg.NewStr(".").ToObject()
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[1] = πTemp005
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(2)
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp007[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp006[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004[0] = πTemp005
					πTemp004[1] = ßhtml5_polyglot.ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[2] = πTemp005
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßdefault_stylesheet_dirs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 36: default_template = 'template.txt'
					πF.SetLineno(36)
					if πE = πClass.SetItem(πF, ßdefault_template.ToObject(), πg.NewStr("template.txt").ToObject()); πE != nil {
						continue
					}
					// line 37: default_template_path = os.path.join(
					πF.SetLineno(37)
					πTemp002 = πF.MakeArgs(2)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßdefault_template_path.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 40: settings_spec = (
					πF.SetLineno(40)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002 = make([]*πg.Object, 19)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template_path); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Mod(πF, πg.NewStr("Specify the template file (UTF-8 encoded).  Default is \"%s\".").ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--template").ToObject()
					πTemp011 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template_path); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πTemp013); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file>").ToObject()); πE != nil {
						continue
					}
					πTemp013 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πTemp010, πTemp011, πTemp013).ToObject()
					πTemp002[0] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--stylesheet").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL[,URL,...]>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßoverrides.ToObject(), ßstylesheet_path.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_comma_separated_list, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Comma separated list of stylesheet URLs. Overrides previous --stylesheet and --stylesheet-path settings.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[1] = πTemp009
					πTemp003 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheets); πE != nil {
						continue
					}
					πTemp003[0] = πTemp011
					if πTemp011, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp011.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp010, πE = πg.Mod(πF, πg.NewStr("Comma separated list of stylesheet paths. Relative paths are expanded if a matching file is found in the --stylesheet-dirs. With --link-stylesheet, the path is rewritten relative to the output HTML file. Default: \"%s\"").ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--stylesheet-path").ToObject()
					πTemp011 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file[,file,...]>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßoverrides.ToObject(), ßstylesheet.ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßvalidate_comma_separated_list, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp014); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheets); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp013 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πTemp010, πTemp011, πTemp013).ToObject()
					πTemp002[2] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--embed-stylesheet").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Embed the stylesheet(s) in the output HTML file.  The stylesheet files must be accessible during processing. This is the default.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[3] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--link-stylesheet").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdest.ToObject(), ßembed_stylesheet.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Link to the stylesheet(s) in the output HTML file. Default: embed stylesheets.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[4] = πTemp009
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet_dirs); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Mod(πF, πg.NewStr("Comma-separated list of directories where stylesheets are found. Used by --stylesheet-path when expanding relative path arguments. Default: \"%s\"").ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--stylesheet-dirs").ToObject()
					πTemp011 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<dir[,dir,...]>").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßvalidate_comma_separated_list, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp014); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet_dirs); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp013 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πTemp010, πTemp011, πTemp013).ToObject()
					πTemp002[5] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--initial-header-level").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πTemp011, πE = πg.GetAttr(πF, πg.NewStr("1 2 3 4 5 6").ToObject(), ßsplit, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßchoices.ToObject(), πTemp013); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), ß1.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Specify the initial header level.  Default is 1 for \"<h1>\".  Does not affect document title & subtitle (see --no-doc-title).").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[6] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--field-name-limit").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewInt(14).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_nonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Specify the maximum width (in characters) for one-column field names.  Longer field names will span an entire row of the table used to render the field list.  Default is 14 characters.  Use 0 for \"no limit\".").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[7] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--option-limit").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewInt(14).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_nonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Specify the maximum width (in characters) for options in option lists.  Longer options will span an entire row of the table used to render the option list.  Default is 14 characters.  Use 0 for \"no limit\".").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[8] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--footnote-references").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ßsuperscript.ToObject()
					πTemp003[1] = ßbrackets.ToObject()
					πTemp011 = πg.NewList(πTemp003...).ToObject()
					if πE = πTemp012.SetItem(πF, ßchoices.ToObject(), πTemp011); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), ßbrackets.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<format>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßoverrides.ToObject(), ßtrim_footnote_reference_space.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Format for footnote references: one of \"superscript\" or \"brackets\".  Default is \"brackets\".").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[9] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--attribution").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = ßdash.ToObject()
					πTemp003[1] = ßparentheses.ToObject()
					πTemp003[2] = ßparens.ToObject()
					πTemp003[3] = ßnone.ToObject()
					πTemp011 = πg.NewList(πTemp003...).ToObject()
					if πE = πTemp012.SetItem(πF, ßchoices.ToObject(), πTemp011); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), ßdash.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<format>").ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Format for block quote attributions: one of \"dash\" (em-dash prefix), \"parentheses\"/\"parens\", or \"none\".  Default is \"dash\".").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[10] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--compact-lists").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Remove extra vertical whitespace between items of \"simple\" bullet lists and enumerated lists.  Default: enabled.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[11] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--no-compact-lists").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdest.ToObject(), ßcompact_lists.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Disable compact simple bullet and enumerated lists.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[12] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--compact-field-lists").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Remove extra vertical whitespace between items of simple field lists.  Default: enabled.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[13] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--no-compact-field-lists").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdest.ToObject(), ßcompact_field_lists.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Disable compact simple field lists.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[14] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--table-style").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Added to standard table classes. Defined styles: \"borderless\". Default: \"\"").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[15] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--math-output").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewStr("HTML math.css").ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Math output format, one of \"MathML\", \"HTML\", \"MathJax\" or \"LaTeX\". Default: \"HTML math.css\"").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[16] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--no-xml-declaration").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßdest.ToObject(), ßxml_declaration.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Omit the XML declaration.  Use with caution.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[17] = πTemp009
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--cloak-email-addresses").ToObject()
					πTemp010 = πg.NewList(πTemp003...).ToObject()
					πTemp012 = πg.NewDict()
					if πE = πTemp012.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp012.SetItem(πF, ßvalidator.ToObject(), πTemp013); πE != nil {
						continue
					}
					πTemp011 = πTemp012.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Obfuscate email addresses to confuse harvesters while still keeping email links usable with standards-compliant browsers.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp002[18] = πTemp009
					πTemp008 = πg.NewTuple(πTemp002...).ToObject()
					πTemp001 = πg.NewTuple3(πg.NewStr("HTML-Specific Options").ToObject(), πTemp005, πTemp008).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 140: config_section = 'html4css1 writer'
					πF.SetLineno(140)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("html4css1 writer").ToObject()); πE != nil {
						continue
					}
					// line 142: def __init__(self):
					πF.SetLineno(142)
					πTemp015 = make([]πg.Param, 1)
					πTemp015[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp015, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Dict
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
							// line 143: self.parts = {}
							πF.SetLineno(143)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparts, πTemp003); πE != nil {
								continue
							}
							// line 144: self.translator_class = HTMLTranslator
							πF.SetLineno(144)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHTMLTranslator); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtranslator_class, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 147: class HTMLTranslator(writers._html_base.HTMLTranslator):
			πF.SetLineno(147)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß_html_base, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßHTMLTranslator, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("HTMLTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Dict
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
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
				var πTemp029 *πg.Object
				_ = πTemp029
				var πTemp030 *πg.Object
				_ = πTemp030
				var πTemp031 *πg.Object
				_ = πTemp031
				var πTemp032 *πg.Object
				_ = πTemp032
				var πTemp033 *πg.Object
				_ = πTemp033
				var πTemp034 *πg.Object
				_ = πTemp034
				var πTemp035 *πg.Object
				_ = πTemp035
				var πTemp036 *πg.Object
				_ = πTemp036
				var πTemp037 *πg.Object
				_ = πTemp037
				var πTemp038 *πg.Object
				_ = πTemp038
				var πTemp039 *πg.Object
				_ = πTemp039
				var πTemp040 *πg.Object
				_ = πTemp040
				var πTemp041 *πg.Object
				_ = πTemp041
				var πTemp042 *πg.Object
				_ = πTemp042
				var πTemp043 *πg.Object
				_ = πTemp043
				var πTemp044 *πg.Object
				_ = πTemp044
				var πTemp045 *πg.Object
				_ = πTemp045
				var πTemp046 *πg.Object
				_ = πTemp046
				var πTemp047 *πg.Object
				_ = πTemp047
				var πTemp048 *πg.Object
				_ = πTemp048
				var πTemp049 *πg.Object
				_ = πTemp049
				var πTemp050 *πg.Object
				_ = πTemp050
				var πTemp051 *πg.Object
				_ = πTemp051
				var πTemp052 *πg.Object
				_ = πTemp052
				var πTemp053 *πg.Object
				_ = πTemp053
				var πTemp054 *πg.Object
				_ = πTemp054
				var πTemp055 *πg.Object
				_ = πTemp055
				var πTemp056 *πg.Object
				_ = πTemp056
				var πTemp057 *πg.Object
				_ = πTemp057
				var πTemp058 *πg.Object
				_ = πTemp058
				var πTemp059 *πg.Object
				_ = πTemp059
				var πTemp060 *πg.Object
				_ = πTemp060
				var πTemp061 *πg.Object
				_ = πTemp061
				var πTemp062 *πg.Object
				_ = πTemp062
				var πTemp063 *πg.Object
				_ = πTemp063
				var πTemp064 *πg.Object
				_ = πTemp064
				var πTemp065 *πg.Object
				_ = πTemp065
				var πTemp066 *πg.Object
				_ = πTemp066
				var πTemp067 *πg.Object
				_ = πTemp067
				var πTemp068 *πg.Object
				_ = πTemp068
				var πTemp069 *πg.Object
				_ = πTemp069
				var πTemp070 *πg.Object
				_ = πTemp070
				var πTemp071 *πg.Object
				_ = πTemp071
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 148: """
					πF.SetLineno(148)
					// line 148: """
					πF.SetLineno(148)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    The html4css1 writer has been optimized to produce visually compact\n    lists (less vertical whitespace).  HTML's mixed content models\n    allow list items to contain \"<li><p>body elements</p></li>\" or\n    \"<li>just text</li>\" or even \"<li>text<p>and body\n    elements</p>combined</li>\", each with different effects.  It would\n    be best to stick with strict body elements in list items, but they\n    affect vertical spacing in older browsers (although they really\n    shouldn't).\n    The html5_polyglot writer solves this using CSS2.\n\n    Here is an outline of the optimization:\n\n    - Check for and omit <p> tags in \"simple\" lists: list items\n      contain either a single paragraph, a nested simple list, or a\n      paragraph followed by a nested simple list.  This means that\n      this list can be compact:\n\n          - Item 1.\n          - Item 2.\n\n      But this list cannot be compact:\n\n          - Item 1.\n\n            This second paragraph forces space between list items.\n\n          - Item 2.\n\n    - In non-list contexts, omit <p> tags on a paragraph if that\n      paragraph is the only child of its parent (footnotes & citations\n      are allowed a label first).\n\n    - Regardless of the above, in definitions, table cells, field bodies,\n      option descriptions, and list items, mark the first child with\n      'class=\"first\"' and the last child with 'class=\"last\"'.  The stylesheet\n      sets the margins (top & bottom respectively) to 0 for these elements.\n\n    The ``no_compact_lists`` setting (``--no-compact-lists`` command-line\n    option) disables list whitespace optimization.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 193: doctype = (
					πF.SetLineno(193)
					if πE = πClass.SetItem(πF, ßdoctype.ToObject(), πg.NewStr("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n").ToObject()); πE != nil {
						continue
					}
					// line 197: content_type = ('<meta http-equiv="Content-Type"'
					πF.SetLineno(197)
					if πE = πClass.SetItem(πF, ßcontent_type.ToObject(), πg.NewStr("<meta http-equiv=\"Content-Type\" content=\"text/html; charset=%s\" />\n").ToObject()); πE != nil {
						continue
					}
					// line 199: content_type_mathml = ('<meta http-equiv="Content-Type"'
					πF.SetLineno(199)
					if πE = πClass.SetItem(πF, ßcontent_type_mathml.ToObject(), πg.NewStr("<meta http-equiv=\"Content-Type\" content=\"application/xhtml+xml; charset=%s\" />\n").ToObject()); πE != nil {
						continue
					}
					// line 203: special_characters = dict(_html_base.HTMLTranslator.special_characters)
					πF.SetLineno(203)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ß_html_base); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßHTMLTranslator, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßspecial_characters, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdict); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßspecial_characters.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 204: special_characters[0xa0] = u'&nbsp;'
					πF.SetLineno(204)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewUnicode("&nbsp;").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßspecial_characters); πE != nil {
						continue
					}
					πTemp004 = πg.NewInt(160).ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
						continue
					}
					// line 207: attribution_formats = {'dash': ('&mdash;', ''),
					πF.SetLineno(207)
					πTemp005 = πg.NewDict()
					πTemp002 = πg.NewTuple2(πg.NewStr("&mdash;").ToObject(), ß.ToObject()).ToObject()
					if πE = πTemp005.SetItem(πF, ßdash.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewStr("(").ToObject(), πg.NewStr(")").ToObject()).ToObject()
					if πE = πTemp005.SetItem(πF, ßparentheses.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewStr("(").ToObject(), πg.NewStr(")").ToObject()).ToObject()
					if πE = πTemp005.SetItem(πF, ßparens.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(ß.ToObject(), ß.ToObject()).ToObject()
					if πE = πTemp005.SetItem(πF, ßnone.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp005.ToObject()
					if πE = πClass.SetItem(πF, ßattribution_formats.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 213: def set_first_last(self, node):
					πF.SetLineno(213)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("set_first_last", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 214: self.set_class_on_child(node, 'first', 0)
							πF.SetLineno(214)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßfirst.ToObject()
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_class_on_child, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 215: self.set_class_on_child(node, 'last', -1)
							πF.SetLineno(215)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßlast.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_class_on_child, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßset_first_last.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 218: def visit_address(self, node):
					πF.SetLineno(218)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("visit_address", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							// line 219: self.visit_docinfo_item(node, 'address', meta=False)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßaddress.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"meta", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 220: self.body.append(self.starttag(node, 'pre', CLASS='address'))
							πF.SetLineno(220)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßpre.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßaddress.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_address.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 224: def visit_admonition(self, node):
					πF.SetLineno(224)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("visit_admonition", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 225: node['classes'].insert(0, 'admonition')
							πF.SetLineno(225)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = ßadmonition.ToObject()
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 226: self.body.append(self.starttag(node, 'div'))
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp004[1] = ßdiv.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 227: self.set_first_last(node)
							πF.SetLineno(227)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_first_last, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_admonition.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 230: def visit_author(self, node):
					πF.SetLineno(230)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("visit_author", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßauthors, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 231: if isinstance(node.parent, nodes.authors):
							πF.SetLineno(231)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßauthor_in_authors, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 232: if self.author_in_authors:
							πF.SetLineno(232)
						Label4:
							// line 233: self.body.append('\n<br />')
							πF.SetLineno(233)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n<br />").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							goto Label3
						Label2:
							// line 235: self.visit_docinfo_item(node, 'author')
							πF.SetLineno(235)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßauthor.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_author.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 237: def depart_author(self, node):
					πF.SetLineno(237)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("depart_author", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßauthors, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 238: if isinstance(node.parent, nodes.authors):
							πF.SetLineno(238)
						Label1:
							// line 239: self.author_in_authors = True
							πF.SetLineno(239)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßauthor_in_authors, πTemp003); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 241: self.depart_docinfo_item()
							πF.SetLineno(241)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_author.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 243: def visit_authors(self, node):
					πF.SetLineno(243)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("visit_authors", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 244: self.visit_docinfo_item(node, 'authors')
							πF.SetLineno(244)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßauthors.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 245: self.author_in_authors = False  # initialize
							πF.SetLineno(245)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßauthor_in_authors, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_authors.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 247: def depart_authors(self, node):
					πF.SetLineno(247)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("depart_authors", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 248: self.depart_docinfo_item()
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_authors.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 251: def visit_colspec(self, node):
					πF.SetLineno(251)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("visit_colspec", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 252: self.colspecs.append(node)
							πF.SetLineno(252)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcolspecs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 254: node.parent.stubs.append(node.attributes.get('stub'))
							πF.SetLineno(254)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßstub.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstubs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_colspec.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 256: def depart_colspec(self, node):
					πF.SetLineno(256)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("depart_colspec", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtotal_width *πg.Object = πg.UnboundLocal
						_ = µtotal_width
						var µcolwidth *πg.Object = πg.UnboundLocal
						_ = µcolwidth
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 []πg.Param
						_ = πTemp013
						var πTemp014 []*πg.Object
						_ = πTemp014
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"descend", πTemp002},
								{"siblings", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßnext_node, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcolspec, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 258: if isinstance(node.next_node(descend=False, siblings=True),
							πF.SetLineno(258)
						Label1:
							// line 260: return
							πF.SetLineno(260)
							πR = πg.None
							continue
							goto Label2
						Label2:
							πTemp006 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßparent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Contains(πF, πTemp007, πg.NewStr("colwidths-auto").ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp010).ToObject()
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßtable_style, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp008, πg.NewStr("colwidths-auto").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp011).ToObject()
							πTemp003 = πTemp006
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label4
							}
							πTemp007 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp009, ßparent, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp012, πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp008, πg.NewStr("colwidths-given").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(!πTemp011).ToObject()
							πTemp003 = πTemp006
						Label4:
							πTemp002 = πTemp003
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 261: if 'colwidths-auto' in node.parent.parent['classes'] or (
							πF.SetLineno(261)
						Label5:
							// line 264: return
							πF.SetLineno(264)
							πR = πg.None
							continue
							goto Label6
						Label6:
							// line 265: total_width = sum(node['colwidth'] for node in self.colspecs)
							πF.SetLineno(265)
							πTemp001 = πF.MakeArgs(1)
							πTemp013 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µnode *πg.Object = πg.UnboundLocal
								_ = µnode
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 bool
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßcolspecs, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp003 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp003 {
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
											πTemp004 = !isStop
										} else {
											πTemp004 = true
											µnode = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 265: total_width = sum(node['colwidth'] for node in self.colspecs)
										πF.SetLineno(265)
										πTemp002 = ßcolwidth.ToObject()
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
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
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtotal_width = πTemp006
							// line 266: self.body.append(self.starttag(node, 'colgroup'))
							πF.SetLineno(266)
							πTemp001 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp014[0] = µnode
							πTemp014[1] = ßcolgroup.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcolspecs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp005 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µnode = πTemp006
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 268: colwidth = int(node['colwidth'] * 100.0 / total_width + 0.5)
							πF.SetLineno(268)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = ßcolwidth.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µnode, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mul(πF, πTemp012, πg.NewFloat(100.0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtotal_width, "total_width"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Div(πF, πTemp008, µtotal_width); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πg.NewFloat(0.5).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcolwidth = πTemp007
							// line 269: self.body.append(self.emptytag(node, 'col',
							πF.SetLineno(269)
							πTemp001 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp014[0] = µnode
							πTemp014[1] = ßcol.ToObject()
							if πE = πg.CheckLocal(πF, µcolwidth, "colwidth"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mod(πF, πg.NewStr("%i%%").ToObject(), µcolwidth); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"width", πTemp006},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßemptytag, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp014, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 271: self.body.append('</colgroup>\n')
							πF.SetLineno(271)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</colgroup>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_colspec.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 276: def is_compactable(self, node):
					πF.SetLineno(276)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("is_compactable", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
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
							// line 277: return ('compact' in node['classes']
							πF.SetLineno(277)
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, ßcompact.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompact_lists, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label2
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, πTemp007, ßopen.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcompact_simple, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßtopic_classes, nil); πE != nil {
								continue
							}
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = ßcontents.ToObject()
							πTemp010 = πg.NewList(πTemp009...).ToObject()
							if πTemp005, πE = πg.Eq(πF, πTemp007, πTemp010); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp009[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcheck_simple_list, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp004 = πTemp007
						Label3:
							πTemp003 = πTemp004
						Label2:
							πTemp001 = πTemp003
						Label1:
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
					if πE = πClass.SetItem(πF, ßis_compactable.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 286: def visit_citation(self, node):
					πF.SetLineno(286)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("visit_citation", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 287: self.body.append(self.starttag(node, 'table',
							πF.SetLineno(287)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtable.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("docutils citation").ToObject()},
								{"frame", ßvoid.ToObject()},
								{"rules", ßnone.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 290: self.body.append('<colgroup><col class="label" /><col /></colgroup>\n'
							πF.SetLineno(290)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<colgroup><col class=\"label\" /><col /></colgroup>\n<tbody valign=\"top\">\n<tr>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 293: self.footnote_backrefs(node)
							πF.SetLineno(293)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfootnote_backrefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_citation.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 295: def depart_citation(self, node):
					πF.SetLineno(295)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("depart_citation", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 296: self.body.append('</td></tr>\n'
							πF.SetLineno(296)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</td></tr>\n</tbody>\n</table>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_citation.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 300: def visit_classifier(self, node):
					πF.SetLineno(300)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("visit_classifier", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 301: self.body.append(' <span class="classifier-delimiter">:</span> ')
							πF.SetLineno(301)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(" <span class=\"classifier-delimiter\">:</span> ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 302: self.body.append(self.starttag(node, 'span', '', CLASS='classifier'))
							πF.SetLineno(302)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp004[1] = ßspan.ToObject()
							πTemp004[2] = ß.ToObject()
							πTemp005 = πg.KWArgs{
								{"CLASS", ßclassifier.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_classifier.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 305: def visit_definition(self, node):
					πF.SetLineno(305)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("visit_definition", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 306: self.body.append('</dt>\n')
							πF.SetLineno(306)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dt>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 307: self.body.append(self.starttag(node, 'dd', ''))
							πF.SetLineno(307)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp004[1] = ßdd.ToObject()
							πTemp004[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 308: self.set_first_last(node)
							πF.SetLineno(308)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_first_last, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_definition.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 311: def visit_definition_list(self, node):
					πF.SetLineno(311)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("visit_definition_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 312: self.body.append(self.starttag(node, 'dl', CLASS='docutils'))
							πF.SetLineno(312)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdl.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßdocutils.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_definition_list.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 315: def visit_description(self, node):
					πF.SetLineno(315)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("visit_description", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 316: self.body.append(self.starttag(node, 'td', ''))
							πF.SetLineno(316)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtd.ToObject()
							πTemp002[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 317: self.set_first_last(node)
							πF.SetLineno(317)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßset_first_last, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_description.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 319: def depart_description(self, node):
					πF.SetLineno(319)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("depart_description", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 320: self.body.append('</td>')
							πF.SetLineno(320)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</td>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_description.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 323: def visit_docinfo(self, node):
					πF.SetLineno(323)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("visit_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 324: self.context.append(len(self.body))
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 325: self.body.append(self.starttag(node, 'table',
							πF.SetLineno(325)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtable.ToObject()
							πTemp005 = πg.KWArgs{
								{"CLASS", ßdocinfo.ToObject()},
								{"frame", ßvoid.ToObject()},
								{"rules", ßnone.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 328: self.body.append('<col class="docinfo-name" />\n'
							πF.SetLineno(328)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<col class=\"docinfo-name\" />\n<col class=\"docinfo-content\" />\n<tbody valign=\"top\">\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 331: self.in_docinfo = True
							πF.SetLineno(331)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_docinfo, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_docinfo.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 333: def depart_docinfo(self, node):
					πF.SetLineno(333)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("depart_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 334: self.body.append('</tbody>\n</table>\n')
							πF.SetLineno(334)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tbody>\n</table>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 335: self.in_docinfo = False
							πF.SetLineno(335)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_docinfo, πTemp003); πE != nil {
								continue
							}
							// line 336: start = self.context.pop()
							πF.SetLineno(336)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstart = πTemp002
							// line 337: self.docinfo = self.body[start:]
							πF.SetLineno(337)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocinfo, πTemp002); πE != nil {
								continue
							}
							// line 338: self.body = []
							πF.SetLineno(338)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbody, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_docinfo.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 340: def visit_docinfo_item(self, node, name, meta=True):
					πF.SetLineno(340)
					πTemp006 = make([]πg.Param, 4)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp006[2] = πg.Param{Name: "name", Def: nil}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp006[3] = πg.Param{Name: "meta", Def: πTemp024}
					πTemp023 = πg.NewFunction(πg.NewCode("visit_docinfo_item", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µname *πg.Object = πArgs[2]
						_ = µname
						var µmeta *πg.Object = πArgs[3]
						_ = µmeta
						var µmeta_tag *πg.Object = πg.UnboundLocal
						_ = µmeta_tag
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µmeta, "meta"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µmeta); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 341: if meta:
							πF.SetLineno(341)
						Label1:
							// line 342: meta_tag = '<meta name="%s" content="%s" />\n' \
							πF.SetLineno(342)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßattval, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πg.NewTuple2(µname, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<meta name=\"%s\" content=\"%s\" />\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							µmeta_tag = πTemp002
							// line 344: self.add_meta(meta_tag)
							πF.SetLineno(344)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmeta_tag, "meta_tag"); πE != nil {
								continue
							}
							πTemp004[0] = µmeta_tag
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_meta, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							// line 345: self.body.append(self.starttag(node, 'tr', ''))
							πF.SetLineno(345)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							πTemp007[1] = ßtr.ToObject()
							πTemp007[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 346: self.body.append('<th class="docinfo-name">%s:</th>\n<td>'
							πF.SetLineno(346)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<th class=\"docinfo-name\">%s:</th>\n<td>").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label3
							}
							goto Label4
							// line 348: if len(node):
							πF.SetLineno(348)
						Label3:
							πTemp004 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßElement, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label5
							}
							goto Label6
							// line 349: if isinstance(node[0], nodes.Element):
							πF.SetLineno(349)
						Label5:
							// line 350: node[0]['classes'].append('first')
							πF.SetLineno(350)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßfirst.ToObject()
							πTemp002 = ßclasses.ToObject()
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label6
						Label6:
							πTemp004 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßElement, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label7
							}
							goto Label8
							// line 351: if isinstance(node[-1], nodes.Element):
							πF.SetLineno(351)
						Label7:
							// line 352: node[-1]['classes'].append('last')
							πF.SetLineno(352)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßlast.ToObject()
							πTemp002 = ßclasses.ToObject()
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label8
						Label8:
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
					if πE = πClass.SetItem(πF, ßvisit_docinfo_item.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 354: def depart_docinfo_item(self):
					πF.SetLineno(354)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("depart_docinfo_item", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
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
							// line 355: self.body.append('</td></tr>\n')
							πF.SetLineno(355)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</td></tr>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_docinfo_item.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 358: def visit_doctest_block(self, node):
					πF.SetLineno(358)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("visit_doctest_block", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 359: self.body.append(self.starttag(node, 'pre', CLASS='doctest-block'))
							πF.SetLineno(359)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßpre.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("doctest-block").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_doctest_block.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 362: def visit_entry(self, node):
					πF.SetLineno(362)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("visit_entry", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 363: writers._html_base.HTMLTranslator.visit_entry(self, node)
							πF.SetLineno(363)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[1] = µnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_html_base, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßHTMLTranslator, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßvisit_entry, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 364: if len(node) == 0:              # empty cell
							πF.SetLineno(364)
						Label1:
							// line 365: self.body.append('&nbsp;')
							πF.SetLineno(365)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("&nbsp;").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 366: self.set_first_last(node)
							πF.SetLineno(366)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_first_last, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_entry.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 369: def visit_enumerated_list(self, node):
					πF.SetLineno(369)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("visit_enumerated_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µold_compact_simple *πg.Object = πg.UnboundLocal
						_ = µold_compact_simple
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							// line 370: """
							πF.SetLineno(370)
							// line 374: atts = {}
							πF.SetLineno(374)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µnode, ßstart.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 375: if 'start' in node:
							πF.SetLineno(375)
						Label1:
							// line 376: atts['start'] = node['start']
							πF.SetLineno(376)
							πTemp002 = ßstart.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßstart.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µnode, ßenumtype.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 377: if 'enumtype' in node:
							πF.SetLineno(377)
						Label3:
							// line 378: atts['class'] = node['enumtype']
							πF.SetLineno(378)
							πTemp002 = ßenumtype.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 381: old_compact_simple = self.compact_simple
							πF.SetLineno(381)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcompact_simple, nil); πE != nil {
								continue
							}
							µold_compact_simple = πTemp002
							// line 382: self.context.append((self.compact_simple, self.compact_p))
							πF.SetLineno(382)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcompact_simple, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcompact_p, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 383: self.compact_p = None
							πF.SetLineno(383)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp004); πE != nil {
								continue
							}
							// line 384: self.compact_simple = self.is_compactable(node)
							πF.SetLineno(384)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßis_compactable, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_simple, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcompact_simple, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µold_compact_simple, "old_compact_simple"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µold_compact_simple); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp004
						Label5:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 385: if self.compact_simple and not old_compact_simple:
							πF.SetLineno(385)
						Label6:
							// line 386: atts['class'] = (atts.get('class', '') + ' simple').strip()
							πF.SetLineno(386)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßclass.ToObject()
							πTemp006[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µatts, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr(" simple").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp004); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 387: self.body.append(self.starttag(node, 'ol', **atts))
							πF.SetLineno(387)
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008[0] = µnode
							πTemp008[1] = ßol.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp008, nil, nil, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_enumerated_list.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 370: """
					πF.SetLineno(370)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("\n        The 'start' attribute does not conform to HTML 4.01's strict.dtd, but\n        cannot be emulated in CSS1 (HTML 5 reincludes it).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_enumerated_list); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
						continue
					}
					// line 389: def depart_enumerated_list(self, node):
					πF.SetLineno(389)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("depart_enumerated_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 390: self.compact_simple, self.compact_p = self.context.pop()
							πF.SetLineno(390)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_simple, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp003); πE != nil {
								continue
							}
							// line 391: self.body.append('</ol>\n')
							πF.SetLineno(391)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("</ol>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_enumerated_list.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 394: def visit_field(self, node):
					πF.SetLineno(394)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("visit_field", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 395: self.body.append(self.starttag(node, 'tr', '', CLASS='field'))
							πF.SetLineno(395)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtr.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßfield.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 397: def depart_field(self, node):
					πF.SetLineno(397)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("depart_field", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 398: self.body.append('</tr>\n')
							πF.SetLineno(398)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tr>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 400: def visit_field_body(self, node):
					πF.SetLineno(400)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("visit_field_body", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µfield *πg.Object = πg.UnboundLocal
						_ = µfield
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
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
							// line 401: self.body.append(self.starttag(node, 'td', '', CLASS='field-body'))
							πF.SetLineno(401)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtd.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("field-body").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 402: self.set_class_on_child(node, 'first', 0)
							πF.SetLineno(402)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßfirst.ToObject()
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßset_class_on_child, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 403: field = node.parent
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							µfield = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcompact_field_list, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µfield, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßdocinfo, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp004 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp001[0] = µfield
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µfield, ßparent, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßindex, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µfield, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.Sub(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πTemp005
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 404: if (self.compact_field_list or
							πF.SetLineno(404)
						Label2:
							// line 410: self.set_class_on_child(node, 'last', -1)
							πF.SetLineno(410)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßlast.ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßset_class_on_child, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_field_body.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 412: def depart_field_body(self, node):
					πF.SetLineno(412)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("depart_field_body", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 413: self.body.append('</td>\n')
							πF.SetLineno(413)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</td>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field_body.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 415: def visit_field_list(self, node):
					πF.SetLineno(415)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("visit_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µfield *πg.Object = πg.UnboundLocal
						_ = µfield
						var µfield_body *πg.Object = πg.UnboundLocal
						_ = µfield_body
						var µchildren *πg.Object = πg.UnboundLocal
						_ = µchildren
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 []*πg.Object
						_ = πTemp015
						var πTemp016 πg.KWArgs
						_ = πTemp016
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 416: self.context.append((self.compact_field_list, self.compact_p))
							πF.SetLineno(416)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcompact_field_list, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcompact_p, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 417: self.compact_p = None
							πF.SetLineno(417)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp003); πE != nil {
								continue
							}
							πTemp003 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, ßcompact.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcompact_field_lists, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label2
							}
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp006, ßopen.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 418: if 'compact' in node['classes']:
							πF.SetLineno(418)
						Label1:
							// line 419: self.compact_field_list = True
							πF.SetLineno(419)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_field_list, πTemp003); πE != nil {
								continue
							}
							goto Label4
							// line 420: elif (self.settings.compact_field_lists
							πF.SetLineno(420)
						Label3:
							// line 422: self.compact_field_list = True
							πF.SetLineno(422)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_field_list, πTemp003); πE != nil {
								continue
							}
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcompact_field_list, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 423: if self.compact_field_list:
							πF.SetLineno(423)
						Label5:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µnode); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp005 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µfield = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 425: field_body = field[-1]
							πF.SetLineno(425)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
								continue
							}
							µfield_body = πTemp004
							// line 426: assert isinstance(field_body, nodes.field_body)
							πF.SetLineno(426)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfield_body, "field_body"); πE != nil {
								continue
							}
							πTemp001[0] = µfield_body
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfield_body, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
								continue
							}
							// line 427: children = [n for n in field_body
							πF.SetLineno(427)
							πTemp008 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µn *πg.Object = πg.UnboundLocal
								_ = µn
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
										case 6:
											goto Label6
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µfield_body, "field_body"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µfield_body); πE != nil {
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
											µn = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πTemp005[0] = µn
										if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßInvisible, nil); πE != nil {
											continue
										}
										πTemp005[1] = πTemp007
										if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
											continue
										}
										if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
											continue
										}
										πTemp004 = πg.GetBool(!πTemp003).ToObject()
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 427: children = [n for n in field_body
										πF.SetLineno(427)
									Label4:
										// line 427: children = [n for n in field_body
										πF.SetLineno(427)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µn, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							µchildren = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001[0] = µchildren
							if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp009, πE = πg.Eq(πF, πTemp011, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp009
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001[0] = µchildren
							if πTemp011, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp010, πE = πg.Eq(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp009 = πTemp010
							if πTemp012, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label11
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp010 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µchildren, πTemp010); πE != nil {
								continue
							}
							πTemp001[0] = πTemp011
							if πTemp011, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp011, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, πTemp011, ßline_block, nil); πE != nil {
								continue
							}
							πTemp010 = πg.NewTuple2(πTemp013, πTemp014).ToObject()
							πTemp001[1] = πTemp010
							if πTemp010, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp009 = πTemp011
						Label11:
							πTemp006 = πTemp009
						Label10:
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							goto Label13
							// line 429: if not (len(children) == 0 or
							πF.SetLineno(429)
						Label12:
							// line 433: self.compact_field_list = False
							πF.SetLineno(433)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_field_list, πTemp006); πE != nil {
								continue
							}
							// line 434: break
							πF.SetLineno(434)
							πTemp005 = true
							continue
							goto Label13
						Label13:
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							goto Label6
						Label6:
							// line 435: self.body.append(self.starttag(node, 'table', frame='void',
							πF.SetLineno(435)
							πTemp001 = πF.MakeArgs(1)
							πTemp015 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp015[0] = µnode
							πTemp015[1] = ßtable.ToObject()
							πTemp016 = πg.KWArgs{
								{"frame", ßvoid.ToObject()},
								{"rules", ßnone.ToObject()},
								{"CLASS", πg.NewStr("docutils field-list").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp015, πTemp016); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp015)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 438: self.body.append('<col class="field-name" />\n'
							πF.SetLineno(438)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<col class=\"field-name\" />\n<col class=\"field-body\" />\n<tbody valign=\"top\">\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field_list.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 442: def depart_field_list(self, node):
					πF.SetLineno(442)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("depart_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 443: self.body.append('</tbody>\n</table>\n')
							πF.SetLineno(443)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tbody>\n</table>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 444: self.compact_field_list, self.compact_p = self.context.pop()
							πF.SetLineno(444)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_field_list, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_field_list.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 446: def visit_field_name(self, node):
					πF.SetLineno(446)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("visit_field_name", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 447: atts = {}
							πF.SetLineno(447)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_docinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 448: if self.in_docinfo:
							πF.SetLineno(448)
						Label1:
							// line 449: atts['class'] = 'docinfo-name'
							πF.SetLineno(449)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("docinfo-name").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp004 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp004, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 451: atts['class'] = 'field-name'
							πF.SetLineno(451)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("field-name").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp004 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp004, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfield_name_limit, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label4
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßfield_name_limit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label4:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 452: if ( self.settings.field_name_limit
							πF.SetLineno(452)
						Label5:
							// line 454: atts['colspan'] = 2
							πF.SetLineno(454)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp004 = ßcolspan.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 455: self.context.append('</tr>\n'
							πF.SetLineno(455)
							πTemp006 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp005
							πTemp009[1] = ßtr.ToObject()
							πTemp009[2] = ß.ToObject()
							πTemp010 = πg.KWArgs{
								{"CLASS", ßfield.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp004, πE = πg.Add(πF, πg.NewStr("</tr>\n").ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewStr("<td>&nbsp;</td>").ToObject()); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label7
						Label6:
							// line 460: self.context.append('')
							πF.SetLineno(460)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label7
						Label7:
							// line 461: self.body.append(self.starttag(node, 'th', '', **atts))
							πF.SetLineno(461)
							πTemp006 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp009[0] = µnode
							πTemp009[1] = ßth.ToObject()
							πTemp009[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp009, nil, nil, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field_name.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 463: def depart_field_name(self, node):
					πF.SetLineno(463)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("depart_field_name", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 464: self.body.append(':</th>')
							πF.SetLineno(464)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(":</th>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 465: self.body.append(self.context.pop())
							πF.SetLineno(465)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field_name.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 468: def visit_footnote(self, node):
					πF.SetLineno(468)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("visit_footnote", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 469: self.body.append(self.starttag(node, 'table',
							πF.SetLineno(469)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtable.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("docutils footnote").ToObject()},
								{"frame", ßvoid.ToObject()},
								{"rules", ßnone.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 472: self.body.append('<colgroup><col class="label" /><col /></colgroup>\n'
							πF.SetLineno(472)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<colgroup><col class=\"label\" /><col /></colgroup>\n<tbody valign=\"top\">\n<tr>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 475: self.footnote_backrefs(node)
							πF.SetLineno(475)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfootnote_backrefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_footnote.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 477: def footnote_backrefs(self, node):
					πF.SetLineno(477)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("footnote_backrefs", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µbacklinks *πg.Object = πg.UnboundLocal
						_ = µbacklinks
						var µbackrefs *πg.Object = πg.UnboundLocal
						_ = µbackrefs
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µbackref *πg.Object = πg.UnboundLocal
						_ = µbackref
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 9:
								goto Label9
							default:
								panic("unexpected function state")
							}
							// line 478: backlinks = []
							πF.SetLineno(478)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µbacklinks = πTemp002
							// line 479: backrefs = node['backrefs']
							πF.SetLineno(479)
							πTemp002 = ßbackrefs.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							µbackrefs = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßfootnote_backlinks, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp002 = µbackrefs
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 480: if self.settings.footnote_backlinks and backrefs:
							πF.SetLineno(480)
						Label2:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp001[0] = µbackrefs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 481: if len(backrefs) == 1:
							πF.SetLineno(481)
						Label5:
							// line 482: self.context.append('')
							πF.SetLineno(482)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 483: self.context.append('</a>')
							πF.SetLineno(483)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</a>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 484: self.context.append('<a class="fn-backref" href="#%s">'
							πF.SetLineno(484)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µbackrefs, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<a class=\"fn-backref\" href=\"#%s\">").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label7
						Label6:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp001[0] = µbackrefs
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp004 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label10
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µi = πTemp005
								µbackref = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(8)
							// line 488: backlinks.append('<a class="fn-backref" href="#%s">%s</a>'
							πF.SetLineno(488)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackref, "backref"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µbackref, µi).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("<a class=\"fn-backref\" href=\"#%s\">%s</a>").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µbacklinks, "backlinks"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µbacklinks, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 490: self.context.append('<em>(%s)</em> ' % ', '.join(backlinks))
							πF.SetLineno(490)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbacklinks, "backlinks"); πE != nil {
								continue
							}
							πTemp008[0] = µbacklinks
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<em>(%s)</em> ").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 491: self.context += ['', '']
							πF.SetLineno(491)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ß.ToObject()
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontext, πTemp005); πE != nil {
								continue
							}
							goto Label7
						Label7:
							goto Label4
						Label3:
							// line 493: self.context.append('')
							πF.SetLineno(493)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 494: self.context += ['', '']
							πF.SetLineno(494)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ß.ToObject()
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontext, πTemp005); πE != nil {
								continue
							}
							goto Label4
						Label4:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 496: if len(node) > 1:
							πF.SetLineno(496)
						Label11:
							if πE = πg.CheckLocal(πF, µbacklinks, "backlinks"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µbacklinks); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 499: if not backlinks:
							πF.SetLineno(499)
						Label13:
							// line 500: node[1]['classes'].append('first')
							πF.SetLineno(500)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfirst.ToObject()
							πTemp002 = ßclasses.ToObject()
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label14
						Label14:
							// line 501: node[-1]['classes'].append('last')
							πF.SetLineno(501)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßlast.ToObject()
							πTemp002 = ßclasses.ToObject()
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label12
						Label12:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfootnote_backrefs.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 503: def depart_footnote(self, node):
					πF.SetLineno(503)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("depart_footnote", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 504: self.body.append('</td></tr>\n'
							πF.SetLineno(504)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</td></tr>\n</tbody>\n</table>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_footnote.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 508: def visit_footnote_reference(self, node):
					πF.SetLineno(508)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp040 = πg.NewFunction(πg.NewCode("visit_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µhref *πg.Object = πg.UnboundLocal
						_ = µhref
						var µformat *πg.Object = πg.UnboundLocal
						_ = µformat
						var µsuffix *πg.Object = πg.UnboundLocal
						_ = µsuffix
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							// line 509: href = '#' + node['refid']
							πF.SetLineno(509)
							πTemp002 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πg.NewStr("#").ToObject(), πTemp003); πE != nil {
								continue
							}
							µhref = πTemp001
							// line 510: format = self.settings.footnote_references
							πF.SetLineno(510)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfootnote_references, nil); πE != nil {
								continue
							}
							µformat = πTemp002
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µformat, ßbrackets.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 511: if format == 'brackets':
							πF.SetLineno(511)
						Label1:
							// line 512: suffix = '['
							πF.SetLineno(512)
							µsuffix = πg.NewStr("[").ToObject()
							// line 513: self.context.append(']')
							πF.SetLineno(513)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("]").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label3
						Label2:
							// line 515: assert format == 'superscript'
							πF.SetLineno(515)
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µformat, ßsuperscript.ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 516: suffix = '<sup>'
							πF.SetLineno(516)
							µsuffix = πg.NewStr("<sup>").ToObject()
							// line 517: self.context.append('</sup>')
							πF.SetLineno(517)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("</sup>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label3
						Label3:
							// line 518: self.body.append(self.starttag(node, 'a', suffix,
							πF.SetLineno(518)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							πTemp006[1] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							πTemp006[2] = µsuffix
							if πE = πg.CheckLocal(πF, µhref, "href"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"CLASS", πg.NewStr("footnote-reference").ToObject()},
								{"href", µhref},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_footnote_reference.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 521: def depart_footnote_reference(self, node):
					πF.SetLineno(521)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("depart_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 522: self.body.append(self.context.pop() + '</a>')
							πF.SetLineno(522)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("</a>").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_footnote_reference.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 525: def visit_generated(self, node):
					πF.SetLineno(525)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp042 = πg.NewFunction(πg.NewCode("visit_generated", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 526: pass
							πF.SetLineno(526)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_generated.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 531: object_image_types = {'.svg': 'image/svg+xml',
					πF.SetLineno(531)
					πTemp005 = πg.NewDict()
					if πE = πTemp005.SetItem(πF, πg.NewStr(".svg").ToObject(), πg.NewStr("image/svg+xml").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, πg.NewStr(".swf").ToObject(), πg.NewStr("application/x-shockwave-flash").ToObject()); πE != nil {
						continue
					}
					πTemp043 = πTemp005.ToObject()
					if πE = πClass.SetItem(πF, ßobject_image_types.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 536: def visit_label(self, node):
					πF.SetLineno(536)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp043 = πg.NewFunction(πg.NewCode("visit_label", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
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
							// line 537: self.body.append(self.starttag(node, 'td', '%s[' % self.context.pop(),
							πF.SetLineno(537)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtd.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s[").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							πTemp006 = πg.KWArgs{
								{"CLASS", ßlabel.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_label.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 540: def depart_label(self, node):
					πF.SetLineno(540)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp044 = πg.NewFunction(πg.NewCode("depart_label", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
							// line 541: self.body.append(']%s</td><td>%s' % (self.context.pop(), self.context.pop()))
							πF.SetLineno(541)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpop, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("]%s</td><td>%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_label.ToObject(), πTemp044); πE != nil {
						continue
					}
					// line 545: def visit_list_item(self, node):
					πF.SetLineno(545)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp045 = πg.NewFunction(πg.NewCode("visit_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
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
							// line 546: self.body.append(self.starttag(node, 'li', ''))
							πF.SetLineno(546)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßli.ToObject()
							πTemp002[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 547: if len(node):
							πF.SetLineno(547)
						Label1:
							// line 548: node[0]['classes'].append('first')
							πF.SetLineno(548)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfirst.ToObject()
							πTemp003 = ßclasses.ToObject()
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_list_item.ToObject(), πTemp045); πE != nil {
						continue
					}
					// line 552: def visit_literal(self, node):
					πF.SetLineno(552)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp046 = πg.NewFunction(πg.NewCode("visit_literal", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µtoken *πg.Object = πg.UnboundLocal
						_ = µtoken
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
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
							// line 554: classes = node.get('classes', [])
							πF.SetLineno(554)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßclasses.ToObject()
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µclasses = πTemp004
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µclasses, ßcode.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 555: if 'code' in classes:
							πF.SetLineno(555)
						Label1:
							// line 557: node['classes'] = [cls for cls in classes if cls != 'code']
							πF.SetLineno(557)
							πTemp006 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µcls *πg.Object = πg.UnboundLocal
								_ = µcls
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
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
										case 6:
											goto Label6
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µclasses); πE != nil {
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
											µcls = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.NE(πF, µcls, ßcode.ToObject()); πE != nil {
											continue
										}
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 557: node['classes'] = [cls for cls in classes if cls != 'code']
										πF.SetLineno(557)
									Label4:
										// line 557: node['classes'] = [cls for cls in classes if cls != 'code']
										πF.SetLineno(557)
										if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µcls, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp008, πTemp007); πE != nil {
								continue
							}
							// line 558: self.body.append(self.starttag(node, 'code', ''))
							πF.SetLineno(558)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßcode.ToObject()
							πTemp002[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 559: return
							πF.SetLineno(559)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 560: self.body.append(
							πF.SetLineno(560)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtt.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp009 = πg.KWArgs{
								{"CLASS", πg.NewStr("docutils literal").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 562: text = node.astext()
							πF.SetLineno(562)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp007
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßwords_and_spaces, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßfindall, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp007, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µtoken = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtoken, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp008 = πg.NewTuple2(πg.NewStr("\n").ToObject(), πg.NewStr(" ").ToObject()).ToObject()
							if πTemp010, πE = πg.Contains(πF, πTemp008, µtoken); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(πTemp010).ToObject()
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label7
							}
							goto Label8
							// line 564: if token.strip():
							πF.SetLineno(564)
						Label6:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp001[0] = µtoken
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßin_word_wrap_point, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label10
							}
							goto Label11
							// line 567: if self.in_word_wrap_point.search(token):
							πF.SetLineno(567)
						Label10:
							// line 568: self.body.append('<span class="pre">%s</span>'
							πF.SetLineno(568)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp002[0] = µtoken
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.Mod(πF, πg.NewStr("<span class=\"pre\">%s</span>").ToObject(), πTemp011); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label12
						Label11:
							// line 571: self.body.append(self.encode(token))
							πF.SetLineno(571)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp002[0] = µtoken
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label12
						Label12:
							goto Label9
							// line 572: elif token in ('\n', ' '):
							πF.SetLineno(572)
						Label7:
							// line 574: self.body.append(token)
							πF.SetLineno(574)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp001[0] = µtoken
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
						Label8:
							// line 577: self.body.append('&nbsp;' * (len(token) - 1) + ' ')
							πF.SetLineno(577)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp002[0] = µtoken
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp011, πE = πg.Sub(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mul(πF, πg.NewStr("&nbsp;").ToObject(), πTemp011); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp008, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
						Label9:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 578: self.body.append('</tt>')
							πF.SetLineno(578)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tt>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 580: raise nodes.SkipNode
							πF.SetLineno(580)
							πE = πF.Raise(πTemp007, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_literal.ToObject(), πTemp046); πE != nil {
						continue
					}
					// line 583: def visit_literal_block(self, node):
					πF.SetLineno(583)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp047 = πg.NewFunction(πg.NewCode("visit_literal_block", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 584: self.body.append(self.starttag(node, 'pre', CLASS='literal-block'))
							πF.SetLineno(584)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßpre.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("literal-block").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_literal_block.ToObject(), πTemp047); πE != nil {
						continue
					}
					// line 587: def depart_literal_block(self, node):
					πF.SetLineno(587)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp048 = πg.NewFunction(πg.NewCode("depart_literal_block", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 588: self.body.append('\n</pre>\n')
							πF.SetLineno(588)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n</pre>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_literal_block.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 591: def visit_option_group(self, node):
					πF.SetLineno(591)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp049 = πg.NewFunction(πg.NewCode("visit_option_group", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 592: atts = {}
							πF.SetLineno(592)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßoption_limit, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßoption_limit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πTemp004
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 593: if ( self.settings.option_limit
							πF.SetLineno(593)
						Label2:
							// line 595: atts['colspan'] = 2
							πF.SetLineno(595)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp004 = ßcolspan.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 596: self.context.append('</tr>\n<tr><td>&nbsp;</td>')
							πF.SetLineno(596)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("</tr>\n<tr><td>&nbsp;</td>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label4
						Label3:
							// line 598: self.context.append('')
							πF.SetLineno(598)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label4
						Label4:
							// line 599: self.body.append(
							πF.SetLineno(599)
							πTemp006 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp009[0] = µnode
							πTemp009[1] = ßtd.ToObject()
							πTemp010 = πg.KWArgs{
								{"CLASS", πg.NewStr("option-group").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp009, nil, πTemp010, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 601: self.body.append('<kbd>')
							πF.SetLineno(601)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("<kbd>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 602: self.context.append(0)          # count number of options
							πF.SetLineno(602)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_group.ToObject(), πTemp049); πE != nil {
						continue
					}
					// line 604: def depart_option_group(self, node):
					πF.SetLineno(604)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp050 = πg.NewFunction(πg.NewCode("depart_option_group", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 605: self.context.pop()
							πF.SetLineno(605)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 606: self.body.append('</kbd></td>\n')
							πF.SetLineno(606)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("</kbd></td>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 607: self.body.append(self.context.pop())
							πF.SetLineno(607)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_group.ToObject(), πTemp050); πE != nil {
						continue
					}
					// line 609: def visit_option_list(self, node):
					πF.SetLineno(609)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp051 = πg.NewFunction(πg.NewCode("visit_option_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 610: self.body.append(
							πF.SetLineno(610)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtable.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("docutils option-list").ToObject()},
								{"frame", ßvoid.ToObject()},
								{"rules", ßnone.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 613: self.body.append('<col class="option" />\n'
							πF.SetLineno(613)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<col class=\"option\" />\n<col class=\"description\" />\n<tbody valign=\"top\">\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_list.ToObject(), πTemp051); πE != nil {
						continue
					}
					// line 617: def depart_option_list(self, node):
					πF.SetLineno(617)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp052 = πg.NewFunction(πg.NewCode("depart_option_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 618: self.body.append('</tbody>\n</table>\n')
							πF.SetLineno(618)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tbody>\n</table>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_list.ToObject(), πTemp052); πE != nil {
						continue
					}
					// line 620: def visit_option_list_item(self, node):
					πF.SetLineno(620)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp053 = πg.NewFunction(πg.NewCode("visit_option_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 621: self.body.append(self.starttag(node, 'tr', ''))
							πF.SetLineno(621)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtr.ToObject()
							πTemp002[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_list_item.ToObject(), πTemp053); πE != nil {
						continue
					}
					// line 623: def depart_option_list_item(self, node):
					πF.SetLineno(623)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp054 = πg.NewFunction(πg.NewCode("depart_option_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 624: self.body.append('</tr>\n')
							πF.SetLineno(624)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tr>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_list_item.ToObject(), πTemp054); πE != nil {
						continue
					}
					// line 628: def should_be_compact_paragraph(self, node):
					πF.SetLineno(628)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp055 = πg.NewFunction(πg.NewCode("should_be_compact_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µfirst *πg.Object = πg.UnboundLocal
						_ = µfirst
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var µparent_length *πg.Object = πg.UnboundLocal
						_ = µparent_length
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
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
						var πTemp015 bool
						_ = πTemp015
						var πTemp016 []πg.Param
						_ = πTemp016
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 12:
								goto Label12
							case 11:
								goto Label11
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 629: """
							πF.SetLineno(629)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßdocument, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp005
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompound, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp005
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 632: if (isinstance(node.parent, nodes.document) or
							πF.SetLineno(632)
						Label2:
							// line 635: return False
							πF.SetLineno(635)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßattlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp002 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
									continue
								}
								µkey = πTemp005
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003[0] = µkey
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßis_not_default, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp004 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Eq(πF, µkey, ßclasses.ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003 = make([]*πg.Object, 0)
							πTemp011 = πg.NewList(πTemp003...).ToObject()
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = ßfirst.ToObject()
							πTemp012 = πg.NewList(πTemp003...).ToObject()
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = ßlast.ToObject()
							πTemp013 = πg.NewList(πTemp003...).ToObject()
							πTemp003 = make([]*πg.Object, 2)
							πTemp003[0] = ßfirst.ToObject()
							πTemp003[1] = ßlast.ToObject()
							πTemp014 = πg.NewList(πTemp003...).ToObject()
							πTemp010 = πg.NewTuple4(πTemp011, πTemp012, πTemp013, πTemp014).ToObject()
							if πTemp015, πE = πg.Contains(πF, πTemp010, µvalue); πE != nil {
								continue
							}
							πTemp009 = πg.GetBool(πTemp015).ToObject()
							πTemp007 = πTemp009
						Label8:
							if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp008).ToObject()
							πTemp004 = πTemp005
						Label7:
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 637: if (node.is_not_default(key) and
							πF.SetLineno(637)
						Label9:
							// line 641: return False
							πF.SetLineno(641)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							goto Label10
						Label10:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 642: first = isinstance(node.parent[0], nodes.label) # skip label
							πF.SetLineno(642)
							πTemp003 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßlabel, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µfirst = πTemp004
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{µfirst, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp002 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µchild = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp003[0] = µchild
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßInvisible, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							goto Label15
							// line 645: if isinstance(child, nodes.Invisible):
							πF.SetLineno(645)
						Label14:
							// line 646: continue
							πF.SetLineno(646)
							continue
							goto Label15
						Label15:
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µchild == µnode).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label16
							}
							goto Label17
							// line 647: if child is node:
							πF.SetLineno(647)
						Label16:
							// line 648: break
							πF.SetLineno(648)
							πTemp002 = true
							continue
							goto Label17
						Label17:
							// line 649: return False
							πF.SetLineno(649)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							// line 650: parent_length = len([n for n in node.parent if not isinstance(
							πF.SetLineno(650)
							πTemp003 = πF.MakeArgs(1)
							πTemp016 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp016, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µn *πg.Object = πg.UnboundLocal
								_ = µn
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 []*πg.Object
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 *πg.Object
								_ = πTemp007
								var πTemp008 *πg.Object
								_ = πTemp008
								var πTemp009 *πg.Object
								_ = πTemp009
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
										case 6:
											goto Label6
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp003 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp003 {
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
											πTemp004 = !isStop
										} else {
											πTemp004 = true
											µn = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πTemp005[0] = µn
										if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßInvisible, nil); πE != nil {
											continue
										}
										if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßlabel, nil); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(πTemp008, πTemp009).ToObject()
										πTemp005[1] = πTemp006
										if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
											continue
										}
										if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
											continue
										}
										πTemp002 = πg.GetBool(!πTemp004).ToObject()
										if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
											continue
										}
										if πTemp004 {
											goto Label4
										}
										goto Label5
										// line 650: parent_length = len([n for n in node.parent if not isinstance(
										πF.SetLineno(650)
									Label4:
										// line 650: parent_length = len([n for n in node.parent if not isinstance(
										πF.SetLineno(650)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µn, nil
									Label6:
										πTemp002 = πSent
										goto Label5
									Label5:
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
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µparent_length = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcompact_simple, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label18
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcompact_field_list, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label18
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcompact_p, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label19
							}
							if πE = πg.CheckLocal(πF, µparent_length, "parent_length"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Eq(πF, µparent_length, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp007
						Label19:
							πTemp001 = πTemp005
						Label18:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label20
							}
							goto Label21
							// line 652: if ( self.compact_simple
							πF.SetLineno(652)
						Label20:
							// line 655: return True
							πF.SetLineno(655)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label21
						Label21:
							// line 656: return False
							πF.SetLineno(656)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					if πE = πClass.SetItem(πF, ßshould_be_compact_paragraph.ToObject(), πTemp055); πE != nil {
						continue
					}
					// line 629: """
					πF.SetLineno(629)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp056}, πg.NewStr("\n        Determine if the <p> tags around paragraph ``node`` can be omitted.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßshould_be_compact_paragraph); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp057, ß__doc__, πTemp056); πE != nil {
						continue
					}
					// line 658: def visit_paragraph(self, node):
					πF.SetLineno(658)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp056 = πg.NewFunction(πg.NewCode("visit_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßshould_be_compact_paragraph, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 659: if self.should_be_compact_paragraph(node):
							πF.SetLineno(659)
						Label1:
							// line 660: self.context.append('')
							πF.SetLineno(660)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 662: self.body.append(self.starttag(node, 'p', ''))
							πF.SetLineno(662)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßp.ToObject()
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 663: self.context.append('</p>\n')
							πF.SetLineno(663)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_paragraph.ToObject(), πTemp056); πE != nil {
						continue
					}
					// line 665: def depart_paragraph(self, node):
					πF.SetLineno(665)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp057 = πg.NewFunction(πg.NewCode("depart_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 666: self.body.append(self.context.pop())
							πF.SetLineno(666)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_paragraph.ToObject(), πTemp057); πE != nil {
						continue
					}
					// line 669: def visit_sidebar(self, node):
					πF.SetLineno(669)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp058 = πg.NewFunction(πg.NewCode("visit_sidebar", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 670: self.body.append(
							πF.SetLineno(670)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßsidebar.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 672: self.set_first_last(node)
							πF.SetLineno(672)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßset_first_last, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 673: self.in_sidebar = True
							πF.SetLineno(673)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_sidebar, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_sidebar.ToObject(), πTemp058); πE != nil {
						continue
					}
					// line 676: def visit_subscript(self, node):
					πF.SetLineno(676)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp059 = πg.NewFunction(πg.NewCode("visit_subscript", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 677: if isinstance(node.parent, nodes.literal_block):
							πF.SetLineno(677)
						Label1:
							// line 678: self.body.append(self.starttag(node, 'span', '',
							πF.SetLineno(678)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßspan.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßsubscript.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 681: self.body.append(self.starttag(node, 'sub', ''))
							πF.SetLineno(681)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßsub.ToObject()
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_subscript.ToObject(), πTemp059); πE != nil {
						continue
					}
					// line 683: def depart_subscript(self, node):
					πF.SetLineno(683)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp060 = πg.NewFunction(πg.NewCode("depart_subscript", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 684: if isinstance(node.parent, nodes.literal_block):
							πF.SetLineno(684)
						Label1:
							// line 685: self.body.append('</span>')
							πF.SetLineno(685)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</span>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 687: self.body.append('</sub>')
							πF.SetLineno(687)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</sub>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßdepart_subscript.ToObject(), πTemp060); πE != nil {
						continue
					}
					// line 690: def visit_subtitle(self, node):
					πF.SetLineno(690)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp061 = πg.NewFunction(πg.NewCode("visit_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtag *πg.Object = πg.UnboundLocal
						_ = µtag
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Dict
						_ = πTemp010
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsidebar, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsection, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 691: if isinstance(node.parent, nodes.sidebar):
							πF.SetLineno(691)
						Label1:
							// line 692: self.body.append(self.starttag(node, 'p', '',
							πF.SetLineno(692)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßp.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", πg.NewStr("sidebar-subtitle").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 694: self.context.append('</p>\n')
							πF.SetLineno(694)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
							// line 695: elif isinstance(node.parent, nodes.document):
							πF.SetLineno(695)
						Label2:
							// line 696: self.body.append(self.starttag(node, 'h2', '', CLASS='subtitle'))
							πF.SetLineno(696)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßh2.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßsubtitle.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 697: self.context.append('</h2>\n')
							πF.SetLineno(697)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</h2>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 698: self.in_document_title = len(self.body)
							πF.SetLineno(698)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_document_title, πTemp002); πE != nil {
								continue
							}
							goto Label4
							// line 699: elif isinstance(node.parent, nodes.section):
							πF.SetLineno(699)
						Label3:
							// line 700: tag = 'h%s' % (self.section_level + self.initial_header_level - 1)
							πF.SetLineno(700)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßinitial_header_level, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("h%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µtag = πTemp002
							// line 701: self.body.append(
							πF.SetLineno(701)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							πTemp005[1] = µtag
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", πg.NewStr("section-subtitle").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(3)
							πTemp010 = πg.NewDict()
							πTemp003 = πTemp010.ToObject()
							πTemp005[0] = πTemp003
							πTemp005[1] = ßspan.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", πg.NewStr("section-subtitle").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 704: self.context.append('</span></%s>\n' % tag)
							πF.SetLineno(704)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("</span></%s>\n").ToObject(), µtag); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_subtitle.ToObject(), πTemp061); πE != nil {
						continue
					}
					// line 706: def depart_subtitle(self, node):
					πF.SetLineno(706)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp062 = πg.NewFunction(πg.NewCode("depart_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							// line 707: self.body.append(self.context.pop())
							πF.SetLineno(707)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_document_title, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 708: if self.in_document_title:
							πF.SetLineno(708)
						Label1:
							// line 709: self.subtitle = self.body[self.in_document_title:-1]
							πF.SetLineno(709)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßin_document_title, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsubtitle, πTemp002); πE != nil {
								continue
							}
							// line 710: self.in_document_title = 0
							πF.SetLineno(710)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_document_title, πTemp002); πE != nil {
								continue
							}
							// line 711: self.body_pre_docinfo.extend(self.body)
							πF.SetLineno(711)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody_pre_docinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 712: self.html_subtitle.extend(self.body)
							πF.SetLineno(712)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhtml_subtitle, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 713: del self.body[:]
							πF.SetLineno(713)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_subtitle.ToObject(), πTemp062); πE != nil {
						continue
					}
					// line 716: def visit_superscript(self, node):
					πF.SetLineno(716)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp063 = πg.NewFunction(πg.NewCode("visit_superscript", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 717: if isinstance(node.parent, nodes.literal_block):
							πF.SetLineno(717)
						Label1:
							// line 718: self.body.append(self.starttag(node, 'span', '',
							πF.SetLineno(718)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßspan.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßsuperscript.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 721: self.body.append(self.starttag(node, 'sup', ''))
							πF.SetLineno(721)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßsup.ToObject()
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_superscript.ToObject(), πTemp063); πE != nil {
						continue
					}
					// line 723: def depart_superscript(self, node):
					πF.SetLineno(723)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp064 = πg.NewFunction(πg.NewCode("depart_superscript", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 724: if isinstance(node.parent, nodes.literal_block):
							πF.SetLineno(724)
						Label1:
							// line 725: self.body.append('</span>')
							πF.SetLineno(725)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</span>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 727: self.body.append('</sup>')
							πF.SetLineno(727)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</sup>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßdepart_superscript.ToObject(), πTemp064); πE != nil {
						continue
					}
					// line 730: def visit_system_message(self, node):
					πF.SetLineno(730)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp065 = πg.NewFunction(πg.NewCode("visit_system_message", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µbackref_text *πg.Object = πg.UnboundLocal
						_ = µbackref_text
						var µbackrefs *πg.Object = πg.UnboundLocal
						_ = µbackrefs
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µbacklinks *πg.Object = πg.UnboundLocal
						_ = µbacklinks
						var µbackref *πg.Object = πg.UnboundLocal
						_ = µbackref
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 731: self.body.append(self.starttag(node, 'div', CLASS='system-message'))
							πF.SetLineno(731)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("system-message").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 732: self.body.append('<p class="system-message-title">')
							πF.SetLineno(732)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<p class=\"system-message-title\">").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 733: backref_text = ''
							πF.SetLineno(733)
							µbackref_text = ß.ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = ßbackrefs.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 734: if len(node['backrefs']):
							πF.SetLineno(734)
						Label1:
							// line 735: backrefs = node['backrefs']
							πF.SetLineno(735)
							πTemp004 = ßbackrefs.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							µbackrefs = πTemp005
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp001[0] = µbackrefs
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 736: if len(backrefs) == 1:
							πF.SetLineno(736)
						Label3:
							// line 737: backref_text = ('; <em><a href="#%s">backlink</a></em>'
							πF.SetLineno(737)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µbackrefs, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("; <em><a href=\"#%s\">backlink</a></em>").ToObject(), πTemp007); πE != nil {
								continue
							}
							µbackref_text = πTemp004
							goto Label5
						Label4:
							// line 740: i = 1
							πF.SetLineno(740)
							µi = πg.NewInt(1).ToObject()
							// line 741: backlinks = []
							πF.SetLineno(741)
							πTemp001 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							µbacklinks = πTemp004
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µbackrefs); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp006 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µbackref = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 743: backlinks.append('<a href="#%s">%s</a>' % (backref, i))
							πF.SetLineno(743)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackref, "backref"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(µbackref, µi).ToObject()
							if πTemp005, πE = πg.Mod(πF, πg.NewStr("<a href=\"#%s\">%s</a>").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µbacklinks, "backlinks"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µbacklinks, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 744: i += 1
							πF.SetLineno(744)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp005
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 745: backref_text = ('; <em>backlinks: %s</em>'
							πF.SetLineno(745)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbacklinks, "backlinks"); πE != nil {
								continue
							}
							πTemp001[0] = µbacklinks
							if πTemp005, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("; <em>backlinks: %s</em>").ToObject(), πTemp007); πE != nil {
								continue
							}
							µbackref_text = πTemp004
							goto Label5
						Label5:
							goto Label2
						Label2:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßline.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 747: if node.hasattr('line'):
							πF.SetLineno(747)
						Label9:
							// line 748: line = ', line %s' % node['line']
							πF.SetLineno(748)
							πTemp005 = ßline.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr(", line %s").ToObject(), πTemp007); πE != nil {
								continue
							}
							µline = πTemp004
							goto Label11
						Label10:
							// line 750: line = ''
							πF.SetLineno(750)
							µline = ß.ToObject()
							goto Label11
						Label11:
							// line 751: self.body.append('System Message: %s/%s '
							πF.SetLineno(751)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = ßtype.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µnode, πTemp007); πE != nil {
								continue
							}
							πTemp007 = ßlevel.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µnode, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µnode, πTemp007); πE != nil {
								continue
							}
							πTemp002[0] = πTemp011
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbackref_text, "backref_text"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple5(πTemp009, πTemp010, πTemp011, µline, µbackref_text).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("System Message: %s/%s (<tt class=\"docutils\">%s</tt>%s)%s</p>\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_system_message.ToObject(), πTemp065); πE != nil {
						continue
					}
					// line 757: def visit_table(self, node):
					πF.SetLineno(757)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp066 = πg.NewFunction(πg.NewCode("visit_table", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
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
							// line 758: self.context.append(self.compact_p)
							πF.SetLineno(758)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcompact_p, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 759: self.compact_p = True
							πF.SetLineno(759)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp003); πE != nil {
								continue
							}
							// line 760: atts = {'border': 1}
							πF.SetLineno(760)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßborder.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							µatts = πTemp002
							// line 761: classes = ['docutils', self.settings.table_style]
							πF.SetLineno(761)
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = ßdocutils.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtable_style, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µclasses = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßalign.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 762: if 'align' in node:
							πF.SetLineno(762)
						Label1:
							// line 763: classes.append('align-%s' % node['align'])
							πF.SetLineno(763)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("align-%s").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µclasses, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßwidth.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 764: if 'width' in node:
							πF.SetLineno(764)
						Label3:
							// line 765: atts['style'] = 'width: %s' % node['width']
							πF.SetLineno(765)
							πTemp003 = ßwidth.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("width: %s").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp006 = ßstyle.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp003); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 766: self.body.append(
							πF.SetLineno(766)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							πTemp007[1] = ßtable.ToObject()
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp008[0] = µclasses
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp009 = πg.KWArgs{
								{"CLASS", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp007, nil, πTemp009, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_table.ToObject(), πTemp066); πE != nil {
						continue
					}
					// line 769: def depart_table(self, node):
					πF.SetLineno(769)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp067 = πg.NewFunction(πg.NewCode("depart_table", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 770: self.compact_p = self.context.pop()
							πF.SetLineno(770)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp002); πE != nil {
								continue
							}
							// line 771: self.body.append('</table>\n')
							πF.SetLineno(771)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("</table>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_table.ToObject(), πTemp067); πE != nil {
						continue
					}
					// line 774: def visit_tbody(self, node):
					πF.SetLineno(774)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp068 = πg.NewFunction(πg.NewCode("visit_tbody", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 775: self.body.append(self.starttag(node, 'tbody', valign='top'))
							πF.SetLineno(775)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtbody.ToObject()
							πTemp003 = πg.KWArgs{
								{"valign", ßtop.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_tbody.ToObject(), πTemp068); πE != nil {
						continue
					}
					// line 777: def depart_tbody(self, node):
					πF.SetLineno(777)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp069 = πg.NewFunction(πg.NewCode("depart_tbody", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 778: self.body.append('</tbody>\n')
							πF.SetLineno(778)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</tbody>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_tbody.ToObject(), πTemp069); πE != nil {
						continue
					}
					// line 781: def visit_thead(self, node):
					πF.SetLineno(781)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp070 = πg.NewFunction(πg.NewCode("visit_thead", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 782: self.body.append(self.starttag(node, 'thead', valign='bottom'))
							πF.SetLineno(782)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßthead.ToObject()
							πTemp003 = πg.KWArgs{
								{"valign", ßbottom.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_thead.ToObject(), πTemp070); πE != nil {
						continue
					}
					// line 784: def depart_thead(self, node):
					πF.SetLineno(784)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "node", Def: nil}
					πTemp071 = πg.NewFunction(πg.NewCode("depart_thead", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 785: self.body.append('</thead>\n')
							πF.SetLineno(785)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</thead>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_thead.ToObject(), πTemp071); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("HTMLTranslator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHTMLTranslator.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 788: class SimpleListChecker(writers._html_base.SimpleListChecker):
			πF.SetLineno(788)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß_html_base, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßSimpleListChecker, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SimpleListChecker", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 790: """
					πF.SetLineno(790)
					// line 790: """
					πF.SetLineno(790)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise `nodes.NodeFound` if non-simple list item is encountered.\n\n    Here \"simple\" means a list item containing nothing other than a single\n    paragraph, a simple list, or a paragraph followed by a simple list.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 797: def visit_list_item(self, node):
					πF.SetLineno(797)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("visit_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µchildren *πg.Object = πg.UnboundLocal
						_ = µchildren
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 798: children = []
							πF.SetLineno(798)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µchildren = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µchild = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp001[0] = µchild
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßInvisible, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 800: if not isinstance(child, nodes.Invisible):
							πF.SetLineno(800)
						Label4:
							// line 801: children.append(child)
							πF.SetLineno(801)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp001[0] = µchild
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchildren, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp002 = µchildren
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µchildren, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							πTemp001 = πF.MakeArgs(2)
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µchildren, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßbullet_list, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							πTemp001 = πF.MakeArgs(2)
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µchildren, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßenumerated_list, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
						Label7:
							πTemp002 = πTemp003
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 802: if (children and isinstance(children[0], nodes.paragraph)
							πF.SetLineno(802)
						Label8:
							// line 805: children.pop()
							πF.SetLineno(805)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µchildren, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label9
						Label9:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001[0] = µchildren
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LE(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 806: if len(children) <= 1:
							πF.SetLineno(806)
						Label10:
							// line 807: return
							πF.SetLineno(807)
							πR = πg.None
							continue
							goto Label12
						Label11:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßNodeFound, nil); πE != nil {
								continue
							}
							// line 809: raise nodes.NodeFound
							πF.SetLineno(809)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label12
						Label12:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_list_item.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 817: def visit_paragraph(self, node):
					πF.SetLineno(817)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("visit_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 818: raise nodes.SkipNode
							πF.SetLineno(818)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_paragraph.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 820: def visit_definition_list(self, node):
					πF.SetLineno(820)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("visit_definition_list", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßNodeFound, nil); πE != nil {
								continue
							}
							// line 821: raise nodes.NodeFound
							πF.SetLineno(821)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_definition_list.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 823: def visit_docinfo(self, node):
					πF.SetLineno(823)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("visit_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/html4css1/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßNodeFound, nil); πE != nil {
								continue
							}
							// line 824: raise nodes.NodeFound
							πF.SetLineno(824)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_docinfo.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("SimpleListChecker").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSimpleListChecker.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("html4css1", Code)
}
