package manpage

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/roman"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßA := πg.InternStr("A")
		ßBLOCKQOUTE_INDENT := πg.InternStr("BLOCKQOUTE_INDENT")
		ßDEFINITION_LIST_INDENT := πg.InternStr("DEFINITION_LIST_INDENT")
		ßFIELD_LIST_INDENT := πg.InternStr("FIELD_LIST_INDENT")
		ßFalse := πg.InternStr("False")
		ßImportError := πg.InternStr("ImportError")
		ßInvisible := πg.InternStr("Invisible")
		ßLITERAL_BLOCK_INDENT := πg.InternStr("LITERAL_BLOCK_INDENT")
		ßMACRO_DEF := πg.InternStr("MACRO_DEF")
		ßNodeVisitor := πg.InternStr("NodeVisitor")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOPTION_LIST_INDENT := πg.InternStr("OPTION_LIST_INDENT")
		ßSkipNode := πg.InternStr("SkipNode")
		ßTable := πg.InternStr("Table")
		ßTranslator := πg.InternStr("Translator")
		ßTrue := πg.InternStr("True")
		ßWriter := πg.InternStr("Writer")
		ß_ := πg.InternStr("_")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__next__ := πg.InternStr("__next__")
		ß__repr__ := πg.InternStr("__repr__")
		ß_active_table := πg.InternStr("_active_table")
		ß_cnt := πg.InternStr("_cnt")
		ß_coldefs := πg.InternStr("_coldefs")
		ß_docinfo := πg.InternStr("_docinfo")
		ß_docinfo_keys := πg.InternStr("_docinfo_keys")
		ß_docinfo_names := πg.InternStr("_docinfo_names")
		ß_field_name := πg.InternStr("_field_name")
		ß_in_docinfo := πg.InternStr("_in_docinfo")
		ß_in_literal := πg.InternStr("_in_literal")
		ß_indent := πg.InternStr("_indent")
		ß_line_block := πg.InternStr("_line_block")
		ß_list_char := πg.InternStr("_list_char")
		ß_minimize_cell := πg.InternStr("_minimize_cell")
		ß_options := πg.InternStr("_options")
		ß_rows := πg.InternStr("_rows")
		ß_style := πg.InternStr("_style")
		ß_tab_char := πg.InternStr("_tab_char")
		ßa := πg.InternStr("a")
		ßaddress := πg.InternStr("address")
		ßadmonition := πg.InternStr("admonition")
		ßalt := πg.InternStr("alt")
		ßappend := πg.InternStr("append")
		ßappend_cell := πg.InternStr("append_cell")
		ßappend_header := πg.InternStr("append_header")
		ßappend_separator := πg.InternStr("append_separator")
		ßarabic := πg.InternStr("arabic")
		ßas_list := πg.InternStr("as_list")
		ßastext := πg.InternStr("astext")
		ßattention := πg.InternStr("attention")
		ßattributes := πg.InternStr("attributes")
		ßauthor := πg.InternStr("author")
		ßauthors := πg.InternStr("authors")
		ßbody := πg.InternStr("body")
		ßbullet := πg.InternStr("bullet")
		ßcaution := πg.InternStr("caution")
		ßcenter := πg.InternStr("center")
		ßchildren := πg.InternStr("children")
		ßcitation := πg.InternStr("citation")
		ßcolspecs := πg.InternStr("colspecs")
		ßcomment := πg.InternStr("comment")
		ßcomment_begin := πg.InternStr("comment_begin")
		ßcompact_p := πg.InternStr("compact_p")
		ßcompact_simple := πg.InternStr("compact_simple")
		ßcompile := πg.InternStr("compile")
		ßcontact := πg.InternStr("contact")
		ßcontext := πg.InternStr("context")
		ßcopyright := πg.InternStr("copyright")
		ßdanger := πg.InternStr("danger")
		ßdate := πg.InternStr("date")
		ßdedent := πg.InternStr("dedent")
		ßdefinition_list_item := πg.InternStr("definition_list_item")
		ßdefs := πg.InternStr("defs")
		ßdelimiter := πg.InternStr("delimiter")
		ßdepart_Text := πg.InternStr("depart_Text")
		ßdepart_address := πg.InternStr("depart_address")
		ßdepart_admonition := πg.InternStr("depart_admonition")
		ßdepart_attention := πg.InternStr("depart_attention")
		ßdepart_attribution := πg.InternStr("depart_attribution")
		ßdepart_author := πg.InternStr("depart_author")
		ßdepart_authors := πg.InternStr("depart_authors")
		ßdepart_block_quote := πg.InternStr("depart_block_quote")
		ßdepart_bullet_list := πg.InternStr("depart_bullet_list")
		ßdepart_caption := πg.InternStr("depart_caption")
		ßdepart_caution := πg.InternStr("depart_caution")
		ßdepart_citation := πg.InternStr("depart_citation")
		ßdepart_classifier := πg.InternStr("depart_classifier")
		ßdepart_colspec := πg.InternStr("depart_colspec")
		ßdepart_compound := πg.InternStr("depart_compound")
		ßdepart_contact := πg.InternStr("depart_contact")
		ßdepart_container := πg.InternStr("depart_container")
		ßdepart_danger := πg.InternStr("depart_danger")
		ßdepart_decoration := πg.InternStr("depart_decoration")
		ßdepart_definition := πg.InternStr("depart_definition")
		ßdepart_definition_list := πg.InternStr("depart_definition_list")
		ßdepart_definition_list_item := πg.InternStr("depart_definition_list_item")
		ßdepart_description := πg.InternStr("depart_description")
		ßdepart_docinfo := πg.InternStr("depart_docinfo")
		ßdepart_docinfo_item := πg.InternStr("depart_docinfo_item")
		ßdepart_doctest_block := πg.InternStr("depart_doctest_block")
		ßdepart_document := πg.InternStr("depart_document")
		ßdepart_emphasis := πg.InternStr("depart_emphasis")
		ßdepart_entry := πg.InternStr("depart_entry")
		ßdepart_enumerated_list := πg.InternStr("depart_enumerated_list")
		ßdepart_error := πg.InternStr("depart_error")
		ßdepart_field := πg.InternStr("depart_field")
		ßdepart_field_body := πg.InternStr("depart_field_body")
		ßdepart_field_list := πg.InternStr("depart_field_list")
		ßdepart_field_name := πg.InternStr("depart_field_name")
		ßdepart_figure := πg.InternStr("depart_figure")
		ßdepart_footer := πg.InternStr("depart_footer")
		ßdepart_footnote := πg.InternStr("depart_footnote")
		ßdepart_footnote_reference := πg.InternStr("depart_footnote_reference")
		ßdepart_generated := πg.InternStr("depart_generated")
		ßdepart_header := πg.InternStr("depart_header")
		ßdepart_hint := πg.InternStr("depart_hint")
		ßdepart_important := πg.InternStr("depart_important")
		ßdepart_inline := πg.InternStr("depart_inline")
		ßdepart_label := πg.InternStr("depart_label")
		ßdepart_legend := πg.InternStr("depart_legend")
		ßdepart_line := πg.InternStr("depart_line")
		ßdepart_line_block := πg.InternStr("depart_line_block")
		ßdepart_list_item := πg.InternStr("depart_list_item")
		ßdepart_literal := πg.InternStr("depart_literal")
		ßdepart_literal_block := πg.InternStr("depart_literal_block")
		ßdepart_math := πg.InternStr("depart_math")
		ßdepart_math_block := πg.InternStr("depart_math_block")
		ßdepart_meta := πg.InternStr("depart_meta")
		ßdepart_note := πg.InternStr("depart_note")
		ßdepart_option := πg.InternStr("depart_option")
		ßdepart_option_argument := πg.InternStr("depart_option_argument")
		ßdepart_option_group := πg.InternStr("depart_option_group")
		ßdepart_option_list := πg.InternStr("depart_option_list")
		ßdepart_option_list_item := πg.InternStr("depart_option_list_item")
		ßdepart_option_string := πg.InternStr("depart_option_string")
		ßdepart_organization := πg.InternStr("depart_organization")
		ßdepart_paragraph := πg.InternStr("depart_paragraph")
		ßdepart_problematic := πg.InternStr("depart_problematic")
		ßdepart_reference := πg.InternStr("depart_reference")
		ßdepart_revision := πg.InternStr("depart_revision")
		ßdepart_row := πg.InternStr("depart_row")
		ßdepart_rubric := πg.InternStr("depart_rubric")
		ßdepart_section := πg.InternStr("depart_section")
		ßdepart_sidebar := πg.InternStr("depart_sidebar")
		ßdepart_status := πg.InternStr("depart_status")
		ßdepart_strong := πg.InternStr("depart_strong")
		ßdepart_subscript := πg.InternStr("depart_subscript")
		ßdepart_subtitle := πg.InternStr("depart_subtitle")
		ßdepart_superscript := πg.InternStr("depart_superscript")
		ßdepart_system_message := πg.InternStr("depart_system_message")
		ßdepart_table := πg.InternStr("depart_table")
		ßdepart_tbody := πg.InternStr("depart_tbody")
		ßdepart_term := πg.InternStr("depart_term")
		ßdepart_tgroup := πg.InternStr("depart_tgroup")
		ßdepart_thead := πg.InternStr("depart_thead")
		ßdepart_tip := πg.InternStr("depart_tip")
		ßdepart_title := πg.InternStr("depart_title")
		ßdepart_title_reference := πg.InternStr("depart_title_reference")
		ßdepart_topic := πg.InternStr("depart_topic")
		ßdepart_transition := πg.InternStr("depart_transition")
		ßdepart_warning := πg.InternStr("depart_warning")
		ßdeunicode := πg.InternStr("deunicode")
		ßdocument := πg.InternStr("document")
		ßdocument_start := πg.InternStr("document_start")
		ßdocutils := πg.InternStr("docutils")
		ßemdash := πg.InternStr("emdash")
		ßemphasis := πg.InternStr("emphasis")
		ßendswith := πg.InternStr("endswith")
		ßensure_eol := πg.InternStr("ensure_eol")
		ßenum_style := πg.InternStr("enum_style")
		ßenumtype := πg.InternStr("enumtype")
		ßerror := πg.InternStr("error")
		ßextend := πg.InternStr("extend")
		ßfield_name := πg.InternStr("field_name")
		ßfirst_child := πg.InternStr("first_child")
		ßfoot := πg.InternStr("foot")
		ßfootnote := πg.InternStr("footnote")
		ßfootnote_backrefs := πg.InternStr("footnote_backrefs")
		ßformat := πg.InternStr("format")
		ßget := πg.InternStr("get")
		ßget_language := πg.InternStr("get_language")
		ßget_width := πg.InternStr("get_width")
		ßhasattr := πg.InternStr("hasattr")
		ßhead := πg.InternStr("head")
		ßheader := πg.InternStr("header")
		ßheader_written := πg.InternStr("header_written")
		ßhint := πg.InternStr("hint")
		ßid := πg.InternStr("id")
		ßimportant := πg.InternStr("important")
		ßindent := πg.InternStr("indent")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßl := πg.InternStr("l")
		ßlabel := πg.InternStr("label")
		ßlabels := πg.InternStr("labels")
		ßlanguage := πg.InternStr("language")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßline := πg.InternStr("line")
		ßlist := πg.InternStr("list")
		ßlist_end := πg.InternStr("list_end")
		ßlist_start := πg.InternStr("list_start")
		ßliteral := πg.InternStr("literal")
		ßliteral_block := πg.InternStr("literal_block")
		ßlower := πg.InternStr("lower")
		ßloweralpha := πg.InternStr("loweralpha")
		ßmanpage := πg.InternStr("manpage")
		ßmanual_group := πg.InternStr("manual_group")
		ßmanual_section := πg.InternStr("manual_section")
		ßmatch := πg.InternStr("match")
		ßmorecols := πg.InternStr("morecols")
		ßmorerows := πg.InternStr("morerows")
		ßname := πg.InternStr("name")
		ßnew_row := πg.InternStr("new_row")
		ßnext := πg.InternStr("next")
		ßnodes := πg.InternStr("nodes")
		ßnote := πg.InternStr("note")
		ßobject := πg.InternStr("object")
		ßoption_list_item := πg.InternStr("option_list_item")
		ßord := πg.InternStr("ord")
		ßorganization := πg.InternStr("organization")
		ßoutput := πg.InternStr("output")
		ßparent := πg.InternStr("parent")
		ßpop := πg.InternStr("pop")
		ßpossibly_a_roff_command := πg.InternStr("possibly_a_roff_command")
		ßproblematic := πg.InternStr("problematic")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßreplace := πg.InternStr("replace")
		ßreporter := πg.InternStr("reporter")
		ßrevision := πg.InternStr("revision")
		ßroman := πg.InternStr("roman")
		ßrstrip := πg.InternStr("rstrip")
		ßsection := πg.InternStr("section")
		ßsection_level := πg.InternStr("section_level")
		ßsettings := πg.InternStr("settings")
		ßsidebar := πg.InternStr("sidebar")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßstart := πg.InternStr("start")
		ßstartswith := πg.InternStr("startswith")
		ßstatus := πg.InternStr("status")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßstrong := πg.InternStr("strong")
		ßsub := πg.InternStr("sub")
		ßsubtitle := πg.InternStr("subtitle")
		ßsupported := πg.InternStr("supported")
		ßsys := πg.InternStr("sys")
		ßterm := πg.InternStr("term")
		ßtip := πg.InternStr("tip")
		ßtitle := πg.InternStr("title")
		ßtitle_reference := πg.InternStr("title_reference")
		ßtitle_upper := πg.InternStr("title_upper")
		ßtoRoman := πg.InternStr("toRoman")
		ßtopic := πg.InternStr("topic")
		ßtopic_class := πg.InternStr("topic_class")
		ßtranslate := πg.InternStr("translate")
		ßtranslator_class := πg.InternStr("translator_class")
		ßtype := πg.InternStr("type")
		ßunimplemented_visit := πg.InternStr("unimplemented_visit")
		ßupper := πg.InternStr("upper")
		ßupperalpha := πg.InternStr("upperalpha")
		ßuri := πg.InternStr("uri")
		ßversion := πg.InternStr("version")
		ßversion_info := πg.InternStr("version_info")
		ßvisit_Text := πg.InternStr("visit_Text")
		ßvisit_address := πg.InternStr("visit_address")
		ßvisit_admonition := πg.InternStr("visit_admonition")
		ßvisit_attention := πg.InternStr("visit_attention")
		ßvisit_attribution := πg.InternStr("visit_attribution")
		ßvisit_author := πg.InternStr("visit_author")
		ßvisit_authors := πg.InternStr("visit_authors")
		ßvisit_block_quote := πg.InternStr("visit_block_quote")
		ßvisit_bullet_list := πg.InternStr("visit_bullet_list")
		ßvisit_caption := πg.InternStr("visit_caption")
		ßvisit_caution := πg.InternStr("visit_caution")
		ßvisit_citation := πg.InternStr("visit_citation")
		ßvisit_citation_reference := πg.InternStr("visit_citation_reference")
		ßvisit_classifier := πg.InternStr("visit_classifier")
		ßvisit_colspec := πg.InternStr("visit_colspec")
		ßvisit_comment := πg.InternStr("visit_comment")
		ßvisit_compound := πg.InternStr("visit_compound")
		ßvisit_contact := πg.InternStr("visit_contact")
		ßvisit_container := πg.InternStr("visit_container")
		ßvisit_copyright := πg.InternStr("visit_copyright")
		ßvisit_danger := πg.InternStr("visit_danger")
		ßvisit_date := πg.InternStr("visit_date")
		ßvisit_decoration := πg.InternStr("visit_decoration")
		ßvisit_definition := πg.InternStr("visit_definition")
		ßvisit_definition_list := πg.InternStr("visit_definition_list")
		ßvisit_definition_list_item := πg.InternStr("visit_definition_list_item")
		ßvisit_description := πg.InternStr("visit_description")
		ßvisit_docinfo := πg.InternStr("visit_docinfo")
		ßvisit_docinfo_item := πg.InternStr("visit_docinfo_item")
		ßvisit_doctest_block := πg.InternStr("visit_doctest_block")
		ßvisit_document := πg.InternStr("visit_document")
		ßvisit_emphasis := πg.InternStr("visit_emphasis")
		ßvisit_entry := πg.InternStr("visit_entry")
		ßvisit_enumerated_list := πg.InternStr("visit_enumerated_list")
		ßvisit_error := πg.InternStr("visit_error")
		ßvisit_field := πg.InternStr("visit_field")
		ßvisit_field_body := πg.InternStr("visit_field_body")
		ßvisit_field_list := πg.InternStr("visit_field_list")
		ßvisit_field_name := πg.InternStr("visit_field_name")
		ßvisit_figure := πg.InternStr("visit_figure")
		ßvisit_footer := πg.InternStr("visit_footer")
		ßvisit_footnote := πg.InternStr("visit_footnote")
		ßvisit_footnote_reference := πg.InternStr("visit_footnote_reference")
		ßvisit_generated := πg.InternStr("visit_generated")
		ßvisit_header := πg.InternStr("visit_header")
		ßvisit_hint := πg.InternStr("visit_hint")
		ßvisit_image := πg.InternStr("visit_image")
		ßvisit_important := πg.InternStr("visit_important")
		ßvisit_inline := πg.InternStr("visit_inline")
		ßvisit_label := πg.InternStr("visit_label")
		ßvisit_legend := πg.InternStr("visit_legend")
		ßvisit_line := πg.InternStr("visit_line")
		ßvisit_line_block := πg.InternStr("visit_line_block")
		ßvisit_list_item := πg.InternStr("visit_list_item")
		ßvisit_literal := πg.InternStr("visit_literal")
		ßvisit_literal_block := πg.InternStr("visit_literal_block")
		ßvisit_math := πg.InternStr("visit_math")
		ßvisit_math_block := πg.InternStr("visit_math_block")
		ßvisit_meta := πg.InternStr("visit_meta")
		ßvisit_note := πg.InternStr("visit_note")
		ßvisit_option := πg.InternStr("visit_option")
		ßvisit_option_argument := πg.InternStr("visit_option_argument")
		ßvisit_option_group := πg.InternStr("visit_option_group")
		ßvisit_option_list := πg.InternStr("visit_option_list")
		ßvisit_option_list_item := πg.InternStr("visit_option_list_item")
		ßvisit_option_string := πg.InternStr("visit_option_string")
		ßvisit_organization := πg.InternStr("visit_organization")
		ßvisit_paragraph := πg.InternStr("visit_paragraph")
		ßvisit_problematic := πg.InternStr("visit_problematic")
		ßvisit_raw := πg.InternStr("visit_raw")
		ßvisit_reference := πg.InternStr("visit_reference")
		ßvisit_revision := πg.InternStr("visit_revision")
		ßvisit_row := πg.InternStr("visit_row")
		ßvisit_rubric := πg.InternStr("visit_rubric")
		ßvisit_section := πg.InternStr("visit_section")
		ßvisit_sidebar := πg.InternStr("visit_sidebar")
		ßvisit_status := πg.InternStr("visit_status")
		ßvisit_strong := πg.InternStr("visit_strong")
		ßvisit_subscript := πg.InternStr("visit_subscript")
		ßvisit_substitution_definition := πg.InternStr("visit_substitution_definition")
		ßvisit_substitution_reference := πg.InternStr("visit_substitution_reference")
		ßvisit_subtitle := πg.InternStr("visit_subtitle")
		ßvisit_superscript := πg.InternStr("visit_superscript")
		ßvisit_system_message := πg.InternStr("visit_system_message")
		ßvisit_table := πg.InternStr("visit_table")
		ßvisit_target := πg.InternStr("visit_target")
		ßvisit_tbody := πg.InternStr("visit_tbody")
		ßvisit_term := πg.InternStr("visit_term")
		ßvisit_tgroup := πg.InternStr("visit_tgroup")
		ßvisit_thead := πg.InternStr("visit_thead")
		ßvisit_tip := πg.InternStr("visit_tip")
		ßvisit_title := πg.InternStr("visit_title")
		ßvisit_title_reference := πg.InternStr("visit_title_reference")
		ßvisit_topic := πg.InternStr("visit_topic")
		ßvisit_transition := πg.InternStr("visit_transition")
		ßvisit_version := πg.InternStr("visit_version")
		ßvisit_warning := πg.InternStr("visit_warning")
		ßwalkabout := πg.InternStr("walkabout")
		ßwarning := πg.InternStr("warning")
		ßwords_and_spaces := πg.InternStr("words_and_spaces")
		ßwrite_colspecs := πg.InternStr("write_colspecs")
		ßwriters := πg.InternStr("writers")
		ßxrange := πg.InternStr("xrange")
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
		var πTemp006 *πg.BaseException
		_ = πTemp006
		var πTemp007 *πg.Traceback
		_ = πTemp007
		var πTemp008 *πg.Dict
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 4:
				goto Label4
			default:
				panic("unexpected function state")
			}
			// line 6: """
			πF.SetLineno(6)
			// line 6: """
			πF.SetLineno(6)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nSimple man page writer for reStructuredText.\n\nMan pages (short for \"manual pages\") contain system documentation on unix-like\nsystems. The pages are grouped in numbered sections:\n\n 1 executable programs and shell commands\n 2 system calls\n 3 library functions\n 4 special files\n 5 file formats\n 6 games\n 7 miscellaneous\n 8 system administration\n\nMan pages are written *troff*, a text file formatting system.\n\nSee http://www.tldp.org/HOWTO/Man-Page for a start.\n\nMan pages have no subsection only parts.\nStandard parts\n\n  NAME ,\n  SYNOPSIS ,\n  DESCRIPTION ,\n  OPTIONS ,\n  FILES ,\n  SEE ALSO ,\n  BUGS ,\n\nand\n\n  AUTHOR .\n\nA unix-like system keeps an index of the DESCRIPTIONs, which is accessible\nby the command whatis or apropos.\n\n").ToObject()); πE != nil {
				continue
			}
			// line 45: __docformat__ = 'reStructuredText'
			πF.SetLineno(45)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 47: import re
			πF.SetLineno(47)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 48: import sys
			πF.SetLineno(48)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.LT(πF, πTemp004, πTemp003); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 50: if sys.version_info < (3, 0):
			πF.SetLineno(50)
		Label1:
			// line 51: range = xrange
			πF.SetLineno(51)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrange.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 53: import docutils
			πF.SetLineno(53)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 54: from docutils import nodes, writers, languages
			πF.SetLineno(54)
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 55: try:
			πF.SetLineno(55)
			πF.PushCheckpoint(4)
			// line 56: import roman
			πF.SetLineno(56)
			if πTemp002, πE = πg.ImportModule(πF, "roman"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßroman.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label3
		Label4:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp006, πTemp007 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label5
			}
			πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
			continue
			// line 57: except ImportError:
			πF.SetLineno(57)
		Label5:
			// line 58: import docutils.utils.roman as roman
			πF.SetLineno(58)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.roman"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßroman.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label3
		Label3:
			// line 60: FIELD_LIST_INDENT = 7
			πF.SetLineno(60)
			if πE = πF.Globals().SetItem(πF, ßFIELD_LIST_INDENT.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
				continue
			}
			// line 61: DEFINITION_LIST_INDENT = 7
			πF.SetLineno(61)
			if πE = πF.Globals().SetItem(πF, ßDEFINITION_LIST_INDENT.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
				continue
			}
			// line 62: OPTION_LIST_INDENT = 7
			πF.SetLineno(62)
			if πE = πF.Globals().SetItem(πF, ßOPTION_LIST_INDENT.ToObject(), πg.NewInt(7).ToObject()); πE != nil {
				continue
			}
			// line 63: BLOCKQOUTE_INDENT = 3.5
			πF.SetLineno(63)
			if πE = πF.Globals().SetItem(πF, ßBLOCKQOUTE_INDENT.ToObject(), πg.NewFloat(3.5).ToObject()); πE != nil {
				continue
			}
			// line 64: LITERAL_BLOCK_INDENT = 3.5
			πF.SetLineno(64)
			if πE = πF.Globals().SetItem(πF, ßLITERAL_BLOCK_INDENT.ToObject(), πg.NewFloat(3.5).ToObject()); πE != nil {
				continue
			}
			// line 68: MACRO_DEF = (r""".
			πF.SetLineno(68)
			if πE = πF.Globals().SetItem(πF, ßMACRO_DEF.ToObject(), πg.NewStr(".\n.nr rst2man-indent-level 0\n.\n.de1 rstReportMargin\n\\\\$1 \\\\n[an-margin]\nlevel \\\\n[rst2man-indent-level]\nlevel margin: \\\\n[rst2man-indent\\\\n[rst2man-indent-level]]\n-\n\\\\n[rst2man-indent0]\n\\\\n[rst2man-indent1]\n\\\\n[rst2man-indent2]\n..\n.de1 INDENT\n.\\\" .rstReportMargin pre:\n. RS \\\\$1\n. nr rst2man-indent\\\\n[rst2man-indent-level] \\\\n[an-margin]\n. nr rst2man-indent-level +1\n.\\\" .rstReportMargin post:\n..\n.de UNINDENT\n. RE\n.\\\" indent \\\\n[an-margin]\n.\\\" old: \\\\n[rst2man-indent\\\\n[rst2man-indent-level]]\n.nr rst2man-indent-level -1\n.\\\" new: \\\\n[rst2man-indent\\\\n[rst2man-indent-level]]\n.in \\\\n[rst2man-indent\\\\n[rst2man-indent-level]]u\n..\n").ToObject()); πE != nil {
				continue
			}
			// line 97: class Writer(writers.Writer):
			πF.SetLineno(97)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp004, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 99: supported = ('manpage',)
					πF.SetLineno(99)
					πTemp001 = πg.NewTuple1(ßmanpage.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 100: """Formats this writer supports."""
					πF.SetLineno(100)
					// line 100: """Formats this writer supports."""
					πF.SetLineno(100)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Formats this writer supports.").ToObject()); πE != nil {
						continue
					}
					// line 102: output = None
					πF.SetLineno(102)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßoutput.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 103: """Final translated form of `document`."""
					πF.SetLineno(103)
					// line 105: def __init__(self):
					πF.SetLineno(105)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 106: writers.Writer.__init__(self)
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßWriter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 107: self.translator_class = Translator
							πF.SetLineno(107)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTranslator); πE != nil {
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
					// line 109: def translate(self):
					πF.SetLineno(109)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("translate", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
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
							// line 110: visitor = self.translator_class(self.document)
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtranslator_class, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvisitor = πTemp003
							// line 111: self.document.walkabout(visitor)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp001[0] = µvisitor
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwalkabout, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 112: self.output = visitor.astext()
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µvisitor, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoutput, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtranslate.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 115: class Table(object):
			πF.SetLineno(115)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Table", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 116: def __init__(self):
					πF.SetLineno(116)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 117: self._rows = []
							πF.SetLineno(117)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_rows, πTemp003); πE != nil {
								continue
							}
							// line 118: self._options = ['center']
							πF.SetLineno(118)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = ßcenter.ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_options, πTemp003); πE != nil {
								continue
							}
							// line 119: self._tab_char = '\t'
							πF.SetLineno(119)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("\t").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tab_char, πTemp002); πE != nil {
								continue
							}
							// line 120: self._coldefs = []
							πF.SetLineno(120)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_coldefs, πTemp003); πE != nil {
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
					// line 121: def new_row(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("new_row", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 122: self._rows.append([])
							πF.SetLineno(122)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_rows, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnew_row.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 123: def append_separator(self, separator):
					πF.SetLineno(123)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "separator", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("append_separator", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µseparator *πg.Object = πArgs[1]
						_ = µseparator
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
							// line 124: """Append the separator for table head."""
							πF.SetLineno(124)
							// line 125: self._rows.append([separator])
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µseparator, "separator"); πE != nil {
								continue
							}
							πTemp002[0] = µseparator
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_rows, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßappend_separator.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 124: """Append the separator for table head."""
					πF.SetLineno(124)
					// line 124: """Append the separator for table head."""
					πF.SetLineno(124)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Append the separator for table head.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Append the separator for table head.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßappend_separator); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 126: def append_cell(self, cell_lines):
					πF.SetLineno(126)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "cell_lines", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("append_cell", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcell_lines *πg.Object = πArgs[1]
						_ = µcell_lines
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
							// line 127: """cell_lines is an array of lines"""
							πF.SetLineno(127)
							// line 128: start = 0
							πF.SetLineno(128)
							µstart = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							πTemp004[0] = µcell_lines
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µcell_lines, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewStr(".sp\n").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 129: if len(cell_lines) > 0 and cell_lines[0] == '.sp\n':
							πF.SetLineno(129)
						Label2:
							// line 130: start = 1
							πF.SetLineno(130)
							µstart = πg.NewInt(1).ToObject()
							goto Label3
						Label3:
							// line 131: self._rows[-1].append(cell_lines[start:])
							πF.SetLineno(131)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µcell_lines, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_rows, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_coldefs, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_rows, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.LT(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 132: if len(self._coldefs) < len(self._rows[-1]):
							πF.SetLineno(132)
						Label4:
							// line 133: self._coldefs.append('l')
							πF.SetLineno(133)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßl.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_coldefs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßappend_cell.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 127: """cell_lines is an array of lines"""
					πF.SetLineno(127)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("cell_lines is an array of lines").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßappend_cell); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 134: def _minimize_cell(self, cell_lines):
					πF.SetLineno(134)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "cell_lines", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_minimize_cell", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcell_lines *πg.Object = πArgs[1]
						_ = µcell_lines
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
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
							case 5:
								goto Label5
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 135: """Remove leading and trailing blank and ``.sp`` lines"""
							πF.SetLineno(135)
							// line 136: while (cell_lines and cell_lines[0] in ('\n', '.sp\n')):
							πF.SetLineno(136)
							πF.PushCheckpoint(2)
							πTemp001 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							πTemp003 = µcell_lines
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µcell_lines, πTemp006); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πg.NewStr("\n").ToObject(), πg.NewStr(".sp\n").ToObject()).ToObject()
							if πTemp008, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp008).ToObject()
							πTemp003 = πTemp005
						Label4:
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 137: del cell_lines[0]
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.DelItem(πF, µcell_lines, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 138: while (cell_lines and cell_lines[-1] in ('\n', '.sp\n')):
							πF.SetLineno(138)
							πF.PushCheckpoint(6)
							πTemp001 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							πTemp003 = µcell_lines
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label8
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µcell_lines, πTemp006); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πg.NewStr("\n").ToObject(), πg.NewStr(".sp\n").ToObject()).ToObject()
							if πTemp008, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp008).ToObject()
							πTemp003 = πTemp005
						Label8:
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(5)
							// line 139: del cell_lines[-1]
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µcell_lines, "cell_lines"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.DelItem(πF, µcell_lines, πTemp003); πE != nil {
								continue
							}
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_minimize_cell.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 135: """Remove leading and trailing blank and ``.sp`` lines"""
					πF.SetLineno(135)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Remove leading and trailing blank and ``.sp`` lines").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_minimize_cell); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 140: def as_list(self):
					πF.SetLineno(140)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("as_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µrow *πg.Object = πg.UnboundLocal
						_ = µrow
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µcell *πg.Object = πg.UnboundLocal
						_ = µcell
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 141: text = ['.TS\n']
							πF.SetLineno(141)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = πg.NewStr(".TS\n").ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µtext = πTemp002
							// line 142: text.append(' '.join(self._options) + ';\n')
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_options, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr(";\n").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 143: text.append('|%s|.\n' % ('|'.join(self._coldefs)))
							πF.SetLineno(143)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_coldefs, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("|").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("|%s|.\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_rows, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µrow = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 146: text.append('_\n')       # line above
							πF.SetLineno(146)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("_\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 147: text.append('T{\n')
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("T{\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp003[0] = µrow
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[0] = πTemp008
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp007 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µi = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 149: cell = row[i]
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µrow, πTemp005); πE != nil {
								continue
							}
							µcell = πTemp008
							// line 150: self._minimize_cell(cell)
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcell, "cell"); πE != nil {
								continue
							}
							πTemp001[0] = µcell
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_minimize_cell, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 151: text.extend(cell)
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcell, "cell"); πE != nil {
								continue
							}
							πTemp001[0] = µcell
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtext, ßextend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n").ToObject()
							if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µtext, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp010, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp009, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label7
							}
							goto Label8
							// line 152: if not text[-1].endswith('\n'):
							πF.SetLineno(152)
						Label7:
							// line 153: text[-1] += '\n'
							πF.SetLineno(153)
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp008
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µtext, πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp008, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp011
							if πE = πg.SetItem(πF, µtext, πTemp010, πTemp005); πE != nil {
								continue
							}
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp001[0] = µrow
							if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.Sub(πF, πTemp011, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LT(πF, µi, πTemp008); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 154: if i < len(row)-1:
							πF.SetLineno(154)
						Label9:
							// line 155: text.append('T}'+self._tab_char+'T{\n')
							πF.SetLineno(155)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ß_tab_char, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πg.NewStr("T}").ToObject(), πTemp010); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp008, πg.NewStr("T{\n").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label11
						Label10:
							// line 157: text.append('T}\n')
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("T}\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label11
						Label11:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 158: text.append('_\n')
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("_\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 159: text.append('.TE\n')
							πF.SetLineno(159)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".TE\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 160: return text
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πR = µtext
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßas_list.ToObject(), πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Table").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTable.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 162: class Translator(nodes.NodeVisitor):
			πF.SetLineno(162)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp004, ßNodeVisitor, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Translator", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
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
				var πTemp072 *πg.Object
				_ = πTemp072
				var πTemp073 *πg.Object
				_ = πTemp073
				var πTemp074 *πg.Object
				_ = πTemp074
				var πTemp075 *πg.Object
				_ = πTemp075
				var πTemp076 *πg.Object
				_ = πTemp076
				var πTemp077 *πg.Object
				_ = πTemp077
				var πTemp078 *πg.Object
				_ = πTemp078
				var πTemp079 *πg.Object
				_ = πTemp079
				var πTemp080 *πg.Object
				_ = πTemp080
				var πTemp081 *πg.Object
				_ = πTemp081
				var πTemp082 *πg.Object
				_ = πTemp082
				var πTemp083 *πg.Object
				_ = πTemp083
				var πTemp084 *πg.Object
				_ = πTemp084
				var πTemp085 *πg.Object
				_ = πTemp085
				var πTemp086 *πg.Object
				_ = πTemp086
				var πTemp087 *πg.Object
				_ = πTemp087
				var πTemp088 *πg.Object
				_ = πTemp088
				var πTemp089 *πg.Object
				_ = πTemp089
				var πTemp090 *πg.Object
				_ = πTemp090
				var πTemp091 *πg.Object
				_ = πTemp091
				var πTemp092 *πg.Object
				_ = πTemp092
				var πTemp093 *πg.Object
				_ = πTemp093
				var πTemp094 *πg.Object
				_ = πTemp094
				var πTemp095 *πg.Object
				_ = πTemp095
				var πTemp096 *πg.Object
				_ = πTemp096
				var πTemp097 *πg.Object
				_ = πTemp097
				var πTemp098 *πg.Object
				_ = πTemp098
				var πTemp099 *πg.Object
				_ = πTemp099
				var πTemp100 *πg.Object
				_ = πTemp100
				var πTemp101 *πg.Object
				_ = πTemp101
				var πTemp102 *πg.Object
				_ = πTemp102
				var πTemp103 *πg.Object
				_ = πTemp103
				var πTemp104 *πg.Object
				_ = πTemp104
				var πTemp105 *πg.Object
				_ = πTemp105
				var πTemp106 *πg.Object
				_ = πTemp106
				var πTemp107 *πg.Object
				_ = πTemp107
				var πTemp108 *πg.Object
				_ = πTemp108
				var πTemp109 *πg.Object
				_ = πTemp109
				var πTemp110 *πg.Object
				_ = πTemp110
				var πTemp111 *πg.Object
				_ = πTemp111
				var πTemp112 *πg.Object
				_ = πTemp112
				var πTemp113 *πg.Object
				_ = πTemp113
				var πTemp114 *πg.Object
				_ = πTemp114
				var πTemp115 *πg.Object
				_ = πTemp115
				var πTemp116 *πg.Object
				_ = πTemp116
				var πTemp117 *πg.Object
				_ = πTemp117
				var πTemp118 *πg.Object
				_ = πTemp118
				var πTemp119 *πg.Object
				_ = πTemp119
				var πTemp120 *πg.Object
				_ = πTemp120
				var πTemp121 *πg.Object
				_ = πTemp121
				var πTemp122 *πg.Object
				_ = πTemp122
				var πTemp123 *πg.Object
				_ = πTemp123
				var πTemp124 *πg.Object
				_ = πTemp124
				var πTemp125 *πg.Object
				_ = πTemp125
				var πTemp126 *πg.Object
				_ = πTemp126
				var πTemp127 *πg.Object
				_ = πTemp127
				var πTemp128 *πg.Object
				_ = πTemp128
				var πTemp129 *πg.Object
				_ = πTemp129
				var πTemp130 *πg.Object
				_ = πTemp130
				var πTemp131 *πg.Object
				_ = πTemp131
				var πTemp132 *πg.Object
				_ = πTemp132
				var πTemp133 *πg.Object
				_ = πTemp133
				var πTemp134 *πg.Object
				_ = πTemp134
				var πTemp135 *πg.Object
				_ = πTemp135
				var πTemp136 *πg.Object
				_ = πTemp136
				var πTemp137 *πg.Object
				_ = πTemp137
				var πTemp138 *πg.Object
				_ = πTemp138
				var πTemp139 *πg.Object
				_ = πTemp139
				var πTemp140 *πg.Object
				_ = πTemp140
				var πTemp141 *πg.Object
				_ = πTemp141
				var πTemp142 *πg.Object
				_ = πTemp142
				var πTemp143 *πg.Object
				_ = πTemp143
				var πTemp144 *πg.Object
				_ = πTemp144
				var πTemp145 *πg.Object
				_ = πTemp145
				var πTemp146 *πg.Object
				_ = πTemp146
				var πTemp147 *πg.Object
				_ = πTemp147
				var πTemp148 *πg.Object
				_ = πTemp148
				var πTemp149 *πg.Object
				_ = πTemp149
				var πTemp150 *πg.Object
				_ = πTemp150
				var πTemp151 *πg.Object
				_ = πTemp151
				var πTemp152 *πg.Object
				_ = πTemp152
				var πTemp153 *πg.Object
				_ = πTemp153
				var πTemp154 *πg.Object
				_ = πTemp154
				var πTemp155 *πg.Object
				_ = πTemp155
				var πTemp156 *πg.Object
				_ = πTemp156
				var πTemp157 *πg.Object
				_ = πTemp157
				var πTemp158 *πg.Object
				_ = πTemp158
				var πTemp159 *πg.Object
				_ = πTemp159
				var πTemp160 *πg.Object
				_ = πTemp160
				var πTemp161 *πg.Object
				_ = πTemp161
				var πTemp162 *πg.Object
				_ = πTemp162
				var πTemp163 *πg.Object
				_ = πTemp163
				var πTemp164 *πg.Object
				_ = πTemp164
				var πTemp165 *πg.Object
				_ = πTemp165
				var πTemp166 *πg.Object
				_ = πTemp166
				var πTemp167 *πg.Object
				_ = πTemp167
				var πTemp168 *πg.Object
				_ = πTemp168
				var πTemp169 *πg.Object
				_ = πTemp169
				var πTemp170 *πg.Object
				_ = πTemp170
				var πTemp171 *πg.Object
				_ = πTemp171
				var πTemp172 *πg.Object
				_ = πTemp172
				var πTemp173 *πg.Object
				_ = πTemp173
				var πTemp174 *πg.Object
				_ = πTemp174
				var πTemp175 *πg.Object
				_ = πTemp175
				var πTemp176 *πg.Object
				_ = πTemp176
				var πTemp177 *πg.Object
				_ = πTemp177
				var πTemp178 *πg.Object
				_ = πTemp178
				var πTemp179 *πg.Object
				_ = πTemp179
				var πTemp180 *πg.Object
				_ = πTemp180
				var πTemp181 *πg.Object
				_ = πTemp181
				var πTemp182 *πg.Object
				_ = πTemp182
				var πTemp183 *πg.Object
				_ = πTemp183
				var πTemp184 *πg.Object
				_ = πTemp184
				var πTemp185 *πg.Object
				_ = πTemp185
				var πTemp186 *πg.Object
				_ = πTemp186
				var πTemp187 *πg.Object
				_ = πTemp187
				var πTemp188 *πg.Object
				_ = πTemp188
				var πTemp189 *πg.Object
				_ = πTemp189
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 163: """"""
					πF.SetLineno(163)
					// line 163: """"""
					πF.SetLineno(163)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					// line 165: words_and_spaces = re.compile(r'\S+| +|\n')
					πF.SetLineno(165)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\\S+| +|\\n").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßwords_and_spaces.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 166: possibly_a_roff_command = re.compile(r'\.\w')
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\\.\\w").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßpossibly_a_roff_command.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 167: document_start = """Man page generated from reStructuredText."""
					πF.SetLineno(167)
					if πE = πClass.SetItem(πF, ßdocument_start.ToObject(), πg.NewStr("Man page generated from reStructuredText.").ToObject()); πE != nil {
						continue
					}
					// line 169: def __init__(self, document):
					πF.SetLineno(169)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "document", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µlcode *πg.Object = πg.UnboundLocal
						_ = µlcode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 170: nodes.NodeVisitor.__init__(self, document)
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßNodeVisitor, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 171: self.settings = settings = document.settings
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocument, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsettings, πTemp003); πE != nil {
								continue
							}
							µsettings = πTemp002
							// line 172: lcode = settings.language_code
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßlanguage_code, nil); πE != nil {
								continue
							}
							µlcode = πTemp002
							// line 173: self.language = languages.get_language(lcode, document.reporter)
							πF.SetLineno(173)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlcode, "lcode"); πE != nil {
								continue
							}
							πTemp001[0] = µlcode
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlanguages); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage, πTemp003); πE != nil {
								continue
							}
							// line 174: self.head = []
							πF.SetLineno(174)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhead, πTemp003); πE != nil {
								continue
							}
							// line 175: self.body = []
							πF.SetLineno(175)
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
							// line 176: self.foot = []
							πF.SetLineno(176)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfoot, πTemp003); πE != nil {
								continue
							}
							// line 177: self.section_level = 0
							πF.SetLineno(177)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_level, πTemp002); πE != nil {
								continue
							}
							// line 178: self.context = []
							πF.SetLineno(178)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontext, πTemp003); πE != nil {
								continue
							}
							// line 179: self.topic_class = ''
							πF.SetLineno(179)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtopic_class, πTemp002); πE != nil {
								continue
							}
							// line 180: self.colspecs = []
							πF.SetLineno(180)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcolspecs, πTemp003); πE != nil {
								continue
							}
							// line 181: self.compact_p = 1
							πF.SetLineno(181)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_p, πTemp002); πE != nil {
								continue
							}
							// line 182: self.compact_simple = None
							πF.SetLineno(182)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_simple, πTemp003); πE != nil {
								continue
							}
							// line 184: self._list_char = []
							πF.SetLineno(184)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_list_char, πTemp003); πE != nil {
								continue
							}
							// line 187: self._docinfo = {
							πF.SetLineno(187)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßtitle.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßtitle_upper.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßsubtitle.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßmanual_section.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßmanual_group.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πTemp004.SetItem(πF, ßauthor.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßdate.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßcopyright.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßversion.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_docinfo, πTemp003); πE != nil {
								continue
							}
							// line 196: self._docinfo_keys = []     # a list to keep the sequence as in source.
							πF.SetLineno(196)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_docinfo_keys, πTemp003); πE != nil {
								continue
							}
							// line 197: self._docinfo_names = {}    # to get name from text not normalized.
							πF.SetLineno(197)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_docinfo_names, πTemp003); πE != nil {
								continue
							}
							// line 198: self._in_docinfo = None
							πF.SetLineno(198)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_docinfo, πTemp003); πE != nil {
								continue
							}
							// line 199: self._active_table = None
							πF.SetLineno(199)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_active_table, πTemp003); πE != nil {
								continue
							}
							// line 200: self._in_literal = False
							πF.SetLineno(200)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp003); πE != nil {
								continue
							}
							// line 201: self.header_written = 0
							πF.SetLineno(201)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßheader_written, πTemp002); πE != nil {
								continue
							}
							// line 202: self._line_block = 0
							πF.SetLineno(202)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_line_block, πTemp002); πE != nil {
								continue
							}
							// line 203: self.authors = []
							πF.SetLineno(203)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßauthors, πTemp003); πE != nil {
								continue
							}
							// line 204: self.section_level = 0
							πF.SetLineno(204)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_level, πTemp002); πE != nil {
								continue
							}
							// line 205: self._indent = [0]
							πF.SetLineno(205)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_indent, πTemp003); πE != nil {
								continue
							}
							// line 216: self.defs = {
							πF.SetLineno(216)
							πTemp004 = πg.NewDict()
							πTemp002 = πg.NewTuple2(πg.NewStr(".INDENT %.1f\n").ToObject(), πg.NewStr(".UNINDENT\n").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßindent.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr(".TP").ToObject(), ß.ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßdefinition_list_item.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr(".TP\n.B ").ToObject(), πg.NewStr("\n").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßfield_name.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\\fB").ToObject(), πg.NewStr("\\fP").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßliteral.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr(".sp\n.nf\n.ft C\n").ToObject(), πg.NewStr("\n.ft P\n.fi\n").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßliteral_block.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr(".TP\n").ToObject(), ß.ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßoption_list_item.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\\fI\\%").ToObject(), πg.NewStr("\\fP").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßreference.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\\fI").ToObject(), πg.NewStr("\\fP").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßemphasis.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\\fB").ToObject(), πg.NewStr("\\fP").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßstrong.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\n.B ").ToObject(), πg.NewStr("\n").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßterm.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\\fI").ToObject(), πg.NewStr("\\fP").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßtitle_reference.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(πg.NewStr(".SS ").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, πg.NewStr("topic-title").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(πg.NewStr(".SS ").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, πg.NewStr("sidebar-title").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("\n.nf\n").ToObject(), πg.NewStr("\n.fi\n").ToObject()).ToObject()
							if πE = πTemp004.SetItem(πF, ßproblematic.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdefs, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 239: def comment_begin(self, text):
					πF.SetLineno(239)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "text", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("comment_begin", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πArgs[1]
						_ = µtext
						var µprefix *πg.Object = πg.UnboundLocal
						_ = µprefix
						var µout_text *πg.Object = πg.UnboundLocal
						_ = µout_text
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
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
							// line 240: """Return commented version of the passed text WITHOUT end of
							πF.SetLineno(240)
							// line 242: prefix = '.\\" '
							πF.SetLineno(242)
							µprefix = πg.NewStr(".\\\" ").ToObject()
							// line 243: out_text = ''.join(
							πF.SetLineno(243)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µin_line *πg.Object = πg.UnboundLocal
								_ = µin_line
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
								var πTemp006 bool
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
										πTemp002 = πF.MakeArgs(1)
										πTemp002[0] = πg.NewStr("\n").ToObject()
										if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp005 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp005 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
											µin_line = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 244: [(prefix + in_line + '\n')
										πF.SetLineno(244)
										if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µin_line, "in_line"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Add(πF, µprefix, µin_line); πE != nil {
											continue
										}
										if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp003, nil
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µout_text = πTemp005
							// line 246: return out_text
							πF.SetLineno(246)
							if πE = πg.CheckLocal(πF, µout_text, "out_text"); πE != nil {
								continue
							}
							πR = µout_text
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcomment_begin.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 240: """Return commented version of the passed text WITHOUT end of
					πF.SetLineno(240)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return commented version of the passed text WITHOUT end of\n        line/comment.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßcomment_begin); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 248: def comment(self, text):
					πF.SetLineno(248)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "text", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("comment", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πArgs[1]
						_ = µtext
						var πTemp001 *πg.Object
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
							// line 249: """Return commented version of the passed text."""
							πF.SetLineno(249)
							// line 250: return self.comment_begin(text)+'.\n'
							πF.SetLineno(250)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp002[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcomment_begin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr(".\n").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcomment.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 249: """Return commented version of the passed text."""
					πF.SetLineno(249)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Return commented version of the passed text.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßcomment); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 252: def ensure_eol(self):
					πF.SetLineno(252)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("ensure_eol", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 253: """Ensure the last line in body is terminated by new line."""
							πF.SetLineno(253)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 254: if len(self.body) > 0 and self.body[-1][-1] != '\n':
							πF.SetLineno(254)
						Label2:
							// line 255: self.body.append('\n')
							πF.SetLineno(255)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßensure_eol.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 253: """Ensure the last line in body is terminated by new line."""
					πF.SetLineno(253)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Ensure the last line in body is terminated by new line.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßensure_eol); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 257: def astext(self):
					πF.SetLineno(257)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
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
							// line 258: """Return the final formatted document as a string."""
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßheader_written, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 259: if not self.header_written:
							πF.SetLineno(259)
						Label1:
							// line 261: self.append_header()
							πF.SetLineno(261)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend_header, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							πTemp004 = πF.MakeArgs(3)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewStr(".sp\n").ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 265: if self.body[i] == '.sp\n':
							πF.SetLineno(265)
						Label6:
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(4).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp009 = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πg.NewStr(".BI ").ToObject(), πg.NewStr(".IP ").ToObject()).ToObject()
							if πTemp008, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label8
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp011
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp010); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Eq(πF, πTemp009, πg.NewStr(".B ").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label9
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(4).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Sub(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp011
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp010); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Eq(πF, πTemp009, πg.NewStr(".TP\n").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
						Label9:
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Eq(πF, πTemp009, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label11
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Sub(πF, µi, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp007); πE != nil {
								continue
							}
							πTemp004[0] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßpossibly_a_roff_command, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp013, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(!πTemp013).ToObject()
							πTemp002 = πTemp006
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label11
							}
							if πTemp009, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(7).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Sub(πF, µi, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp011 = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp012, πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Eq(πF, πTemp010, πg.NewStr(".TP\n.B ").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πTemp013, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp013 {
								goto Label12
							}
							if πTemp009, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(4).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Sub(πF, µi, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp011 = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp012, πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Eq(πF, πTemp010, πg.NewStr("\n.B ").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label12:
							πTemp002 = πTemp006
						Label11:
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label13
							}
							goto Label14
							// line 266: if self.body[i - 1][:4] in ('.BI ', '.IP '):
							πF.SetLineno(266)
						Label8:
							// line 267: self.body[i] = '.\n'
							πF.SetLineno(267)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr(".\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label14
							// line 268: elif (self.body[i - 1][:3] == '.B ' and
							πF.SetLineno(268)
						Label10:
							// line 270: self.body[i] = '.\n'
							πF.SetLineno(270)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr(".\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label14
							// line 271: elif (self.body[i - 1] == '\n' and
							πF.SetLineno(271)
						Label13:
							// line 276: self.body[i] = '.\n'
							πF.SetLineno(276)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr(".\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = µi
							if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label14
						Label14:
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 277: return ''.join(self.head + self.body + self.foot)
							πF.SetLineno(277)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßfoot, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 258: """Return the final formatted document as a string."""
					πF.SetLineno(258)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return the final formatted document as a string.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßastext); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 279: def deunicode(self, text):
					πF.SetLineno(279)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "text", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("deunicode", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πArgs[1]
						_ = µtext
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
							// line 280: text = text.replace(u'\xa0', '\\ ')
							πF.SetLineno(280)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewUnicode("\xc2\xa0").ToObject()
							πTemp001[1] = πg.NewStr("\\ ").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp003
							// line 281: text = text.replace(u'\u2020', '\\(dg')
							πF.SetLineno(281)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewUnicode("\xe2\x80\xa0").ToObject()
							πTemp001[1] = πg.NewStr("\\(dg").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp003
							// line 282: return text
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πR = µtext
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdeunicode.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 284: def visit_Text(self, node):
					πF.SetLineno(284)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("visit_Text", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µreplace_pairs *πg.Object = πg.UnboundLocal
						_ = µreplace_pairs
						var µin_char *πg.Object = πg.UnboundLocal
						_ = µin_char
						var µout_markup *πg.Object = πg.UnboundLocal
						_ = µout_markup
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 285: text = node.astext()
							πF.SetLineno(285)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp002
							// line 286: text = text.replace('\\', '\\e')
							πF.SetLineno(286)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("\\").ToObject()
							πTemp003[1] = πg.NewStr("\\e").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp002
							// line 287: replace_pairs = [
							πF.SetLineno(287)
							πTemp003 = make([]*πg.Object, 4)
							πTemp001 = πg.NewTuple2(πg.NewUnicode("-").ToObject(), πg.NewUnicode("\\-").ToObject()).ToObject()
							πTemp003[0] = πTemp001
							πTemp001 = πg.NewTuple2(πg.NewUnicode("'").ToObject(), πg.NewUnicode("\\(aq").ToObject()).ToObject()
							πTemp003[1] = πTemp001
							πTemp001 = πg.NewTuple2(πg.NewUnicode("\xc2\xb4").ToObject(), πg.NewUnicode("\\'").ToObject()).ToObject()
							πTemp003[2] = πTemp001
							πTemp001 = πg.NewTuple2(πg.NewUnicode("`").ToObject(), πg.NewUnicode("\\(ga").ToObject()).ToObject()
							πTemp003[3] = πTemp001
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µreplace_pairs = πTemp001
							if πE = πg.CheckLocal(πF, µreplace_pairs, "replace_pairs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µreplace_pairs); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
									continue
								}
								µin_char = πTemp006
								µout_markup = πTemp007
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 294: text = text.replace(in_char, out_markup)
							πF.SetLineno(294)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µin_char, "in_char"); πE != nil {
								continue
							}
							πTemp003[0] = µin_char
							if πE = πg.CheckLocal(πF, µout_markup, "out_markup"); πE != nil {
								continue
							}
							πTemp003[1] = µout_markup
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp006
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 296: text = self.deunicode(text)
							πF.SetLineno(296)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdeunicode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 298: if text.startswith('.'):
							πF.SetLineno(298)
						Label4:
							// line 299: text = '\\&' + text
							πF.SetLineno(299)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πg.NewStr("\\&").ToObject(), µtext); πE != nil {
								continue
							}
							µtext = πTemp001
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_in_literal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 300: if self._in_literal:
							πF.SetLineno(300)
						Label6:
							// line 301: text = text.replace('\n.', '\n\\&.')
							πF.SetLineno(301)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("\n.").ToObject()
							πTemp003[1] = πg.NewStr("\n\\&.").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp002
							goto Label7
						Label7:
							// line 302: self.body.append(text)
							πF.SetLineno(302)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
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
					if πE = πClass.SetItem(πF, ßvisit_Text.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 304: def depart_Text(self, node):
					πF.SetLineno(304)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("depart_Text", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 305: pass
							πF.SetLineno(305)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_Text.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 307: def list_start(self, node):
					πF.SetLineno(307)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("list_start", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µenum_char *πg.Object = πg.UnboundLocal
						_ = µenum_char
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 308: class enum_char(object):
							πF.SetLineno(308)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("enum_char", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								var πTemp001 *πg.Dict
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []πg.Param
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 *πg.Object
								_ = πTemp007
								var πTemp008 bool
								_ = πTemp008
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default:
										panic("unexpected function state")
									}
									// line 309: enum_style = {
									πF.SetLineno(309)
									πTemp001 = πg.NewDict()
									if πE = πTemp001.SetItem(πF, ßbullet.ToObject(), πg.NewStr("\\(bu").ToObject()); πE != nil {
										continue
									}
									if πE = πTemp001.SetItem(πF, ßemdash.ToObject(), πg.NewStr("\\(em").ToObject()); πE != nil {
										continue
									}
									πTemp002 = πTemp001.ToObject()
									if πE = πClass.SetItem(πF, ßenum_style.ToObject(), πTemp002); πE != nil {
										continue
									}
									// line 314: def __init__(self, style):
									πF.SetLineno(314)
									πTemp003 = make([]πg.Param, 2)
									πTemp003[0] = πg.Param{Name: "self", Def: nil}
									πTemp003[1] = πg.Param{Name: "style", Def: nil}
									πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]
										_ = µself
										var µstyle *πg.Object = πArgs[1]
										_ = µstyle
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
										var πTemp006 []*πg.Object
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
											// line 315: self._style = style
											πF.SetLineno(315)
											if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstyle); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_style, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.Contains(πF, µnode, ßstart.ToObject()); πE != nil {
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
											// line 316: if 'start' in node:
											πF.SetLineno(316)
										Label1:
											// line 317: self._cnt = node['start'] - 1
											πF.SetLineno(317)
											πTemp003 = ßstart.ToObject()
											if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_cnt, πTemp003); πE != nil {
												continue
											}
											goto Label3
										Label2:
											// line 319: self._cnt = 0
											πF.SetLineno(319)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_cnt, πTemp001); πE != nil {
												continue
											}
											goto Label3
										Label3:
											// line 320: self._indent = 2
											πF.SetLineno(320)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(2).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_indent, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µstyle, ßarabic.ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label4
											}
											if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µstyle, ßloweralpha.ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label5
											}
											if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µstyle, ßupperalpha.ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label6
											}
											πTemp005 = πF.MakeArgs(1)
											πTemp005[0] = ßroman.ToObject()
											if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µstyle, ßendswith, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label7
											}
											goto Label8
											// line 321: if style == 'arabic':
											πF.SetLineno(321)
										Label4:
											// line 324: self._indent = len(str(len(node.children)))
											πF.SetLineno(324)
											πTemp005 = πF.MakeArgs(1)
											πTemp006 = πF.MakeArgs(1)
											πTemp007 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
												continue
											}
											πTemp007[0] = πTemp001
											if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp007)
											πTemp006[0] = πTemp003
											if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp006)
											πTemp005[0] = πTemp003
											if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_indent, πTemp001); πE != nil {
												continue
											}
											// line 325: self._indent += len(str(self._cnt)) + 1
											πF.SetLineno(325)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_indent, nil); πE != nil {
												continue
											}
											πTemp005 = πF.MakeArgs(1)
											πTemp006 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											πTemp006[0] = πTemp004
											if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
												continue
											}
											if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp006)
											πTemp005[0] = πTemp008
											if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
												continue
											}
											if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp003, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πTemp004, πE = πg.IAdd(πF, πTemp001, πTemp003); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_indent, πTemp004); πE != nil {
												continue
											}
											goto Label8
											// line 326: elif style == 'loweralpha':
											πF.SetLineno(326)
										Label5:
											// line 327: self._cnt += ord('a') - 1
											πF.SetLineno(327)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											πTemp005 = πF.MakeArgs(1)
											πTemp005[0] = ßa.ToObject()
											if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
												continue
											}
											if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp003, πE = πg.Sub(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πTemp004, πE = πg.IAdd(πF, πTemp001, πTemp003); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_cnt, πTemp004); πE != nil {
												continue
											}
											// line 328: self._indent = 3
											πF.SetLineno(328)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(3).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_indent, πTemp001); πE != nil {
												continue
											}
											goto Label8
											// line 329: elif style == 'upperalpha':
											πF.SetLineno(329)
										Label6:
											// line 330: self._cnt += ord('A') - 1
											πF.SetLineno(330)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											πTemp005 = πF.MakeArgs(1)
											πTemp005[0] = ßA.ToObject()
											if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
												continue
											}
											if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp003, πE = πg.Sub(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πTemp004, πE = πg.IAdd(πF, πTemp001, πTemp003); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_cnt, πTemp004); πE != nil {
												continue
											}
											// line 331: self._indent = 3
											πF.SetLineno(331)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(3).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_indent, πTemp001); πE != nil {
												continue
											}
											goto Label8
											// line 332: elif style.endswith('roman'):
											πF.SetLineno(332)
										Label7:
											// line 333: self._indent = 5
											πF.SetLineno(333)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(5).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_indent, πTemp001); πE != nil {
												continue
											}
											goto Label8
										Label8:
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
										continue
									}
									// line 335: def __next__(self):
									πF.SetLineno(335)
									πTemp003 = make([]πg.Param, 1)
									πTemp003[0] = πg.Param{Name: "self", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__next__", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]
										_ = µself
										var µres *πg.Object = πg.UnboundLocal
										_ = µres
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 *πg.Object
										_ = πTemp002
										var πTemp003 bool
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
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, πTemp002, ßbullet.ToObject()); πE != nil {
												continue
											}
											if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label1
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, πTemp002, ßemdash.ToObject()); πE != nil {
												continue
											}
											if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label2
											}
											goto Label3
											// line 336: if self._style == 'bullet':
											πF.SetLineno(336)
										Label1:
											// line 337: return self.enum_style[self._style]
											πF.SetLineno(337)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											πTemp001 = πTemp002
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetAttr(πF, µself, ßenum_style, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
												continue
											}
											πR = πTemp002
											continue
											goto Label3
											// line 338: elif self._style == 'emdash':
											πF.SetLineno(338)
										Label2:
											// line 339: return self.enum_style[self._style]
											πF.SetLineno(339)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											πTemp001 = πTemp002
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetAttr(πF, µself, ßenum_style, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
												continue
											}
											πR = πTemp002
											continue
											goto Label3
										Label3:
											// line 340: self._cnt += 1
											πF.SetLineno(340)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß_cnt, πTemp002); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, πTemp002, ßarabic.ToObject()); πE != nil {
												continue
											}
											if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label4
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											πTemp004 = πg.NewTuple2(ßloweralpha.ToObject(), ßupperalpha.ToObject()).ToObject()
											if πTemp003, πE = πg.Contains(πF, πTemp004, πTemp002); πE != nil {
												continue
											}
											πTemp001 = πg.GetBool(πTemp003).ToObject()
											if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label5
											}
											πTemp005 = πF.MakeArgs(1)
											πTemp005[0] = ßroman.ToObject()
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßendswith, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label6
											}
											goto Label7
											// line 342: if self._style == 'arabic':
											πF.SetLineno(342)
										Label4:
											// line 343: return "%d." % self._cnt
											πF.SetLineno(343)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Mod(πF, πg.NewStr("%d.").ToObject(), πTemp002); πE != nil {
												continue
											}
											πR = πTemp001
											continue
											goto Label8
											// line 344: elif self._style in ('loweralpha', 'upperalpha'):
											πF.SetLineno(344)
										Label5:
											// line 345: return "%c." % self._cnt
											πF.SetLineno(345)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Mod(πF, πg.NewStr("%c.").ToObject(), πTemp002); πE != nil {
												continue
											}
											πR = πTemp001
											continue
											goto Label8
											// line 346: elif self._style.endswith('roman'):
											πF.SetLineno(346)
										Label6:
											// line 347: res = roman.toRoman(self._cnt) + '.'
											πF.SetLineno(347)
											πTemp005 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											πTemp005[0] = πTemp002
											if πTemp002, πE = πg.ResolveGlobal(πF, ßroman); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtoRoman, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr(".").ToObject()); πE != nil {
												continue
											}
											µres = πTemp001
											πTemp005 = πF.MakeArgs(1)
											πTemp005[0] = ßupper.ToObject()
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstartswith, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label9
											}
											goto Label10
											// line 348: if self._style.startswith('upper'):
											πF.SetLineno(348)
										Label9:
											// line 349: return res.upper()
											πF.SetLineno(349)
											if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µres, ßupper, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
												continue
											}
											πR = πTemp002
											continue
											goto Label10
										Label10:
											// line 350: return res.lower()
											πF.SetLineno(350)
											if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µres, ßlower, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
												continue
											}
											πR = πTemp002
											continue
											goto Label8
										Label7:
											// line 352: return "%d." % self._cnt
											πF.SetLineno(352)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß_cnt, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Mod(πF, πg.NewStr("%d.").ToObject(), πTemp002); πE != nil {
												continue
											}
											πR = πTemp001
											continue
											goto Label8
										Label8:
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__next__.ToObject(), πTemp004); πE != nil {
										continue
									}
									if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
										continue
									}
									if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßversion_info, nil); πE != nil {
										continue
									}
									πTemp006 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
									if πTemp005, πE = πg.LT(πF, πTemp007, πTemp006); πE != nil {
										continue
									}
									if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
										continue
									}
									if πTemp008 {
										goto Label1
									}
									goto Label2
									// line 354: if sys.version_info < (3, 0):
									πF.SetLineno(354)
								Label1:
									// line 355: next = __next__
									πF.SetLineno(355)
									if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__next__); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp005); πE != nil {
										continue
									}
									goto Label2
								Label2:
									// line 357: def get_width(self):
									πF.SetLineno(357)
									πTemp003 = make([]πg.Param, 1)
									πTemp003[0] = πg.Param{Name: "self", Def: nil}
									πTemp005 = πg.NewFunction(πg.NewCode("get_width", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]
										_ = µself
										var πTemp001 *πg.Object
										_ = πTemp001
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
											// line 358: return self._indent
											πF.SetLineno(358)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_indent, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßget_width.ToObject(), πTemp005); πE != nil {
										continue
									}
									// line 359: def __repr__(self):
									πF.SetLineno(359)
									πTemp003 = make([]πg.Param, 1)
									πTemp003[0] = πg.Param{Name: "self", Def: nil}
									πTemp006 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]
										_ = µself
										var πTemp001 *πg.Object
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
											// line 360: return 'enum_style-%s' % list(self._style)
											πF.SetLineno(360)
											πTemp002 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, µself, ß_style, nil); πE != nil {
												continue
											}
											πTemp002[0] = πTemp003
											if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
												continue
											}
											if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp002)
											if πTemp001, πE = πg.Mod(πF, πg.NewStr("enum_style-%s").ToObject(), πTemp004); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp006); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("enum_char").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µenum_char = πTemp005
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µnode, ßenumtype.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 362: if 'enumtype' in node:
							πF.SetLineno(362)
						Label1:
							// line 363: self._list_char.append(enum_char(node['enumtype']))
							πF.SetLineno(363)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp002 = ßenumtype.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µenum_char, "enum_char"); πE != nil {
								continue
							}
							if πTemp002, πE = µenum_char.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 365: self._list_char.append(enum_char('bullet'))
							πF.SetLineno(365)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = ßbullet.ToObject()
							if πE = πg.CheckLocal(πF, µenum_char, "enum_char"); πE != nil {
								continue
							}
							if πTemp002, πE = µenum_char.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label3:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 366: if len(self._list_char) > 1:
							πF.SetLineno(366)
						Label4:
							// line 368: self.indent(self._list_char[-2].get_width())
							πF.SetLineno(368)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßget_width, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label6
						Label5:
							// line 370: self.indent(self._list_char[-1].get_width())
							πF.SetLineno(370)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßget_width, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßlist_start.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 372: def list_end(self):
					πF.SetLineno(372)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("list_end", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 373: self.dedent()
							πF.SetLineno(373)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 374: self._list_char.pop()
							πF.SetLineno(374)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßlist_end.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 376: def header(self):
					πF.SetLineno(376)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("header", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtmpl *πg.Object = πg.UnboundLocal
						_ = µtmpl
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
							// line 377: tmpl = (".TH %(title_upper)s %(manual_section)s"
							πF.SetLineno(377)
							µtmpl = πg.NewStr(".TH %(title_upper)s %(manual_section)s \"%(date)s\" \"%(version)s\" \"%(manual_group)s\"\n.SH NAME\n%(title)s \\- %(subtitle)s\n").ToObject()
							// line 381: return tmpl % self._docinfo
							πF.SetLineno(381)
							if πE = πg.CheckLocal(πF, µtmpl, "tmpl"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, µtmpl, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßheader.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 383: def append_header(self):
					πF.SetLineno(383)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("append_header", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 384: """append header with .TH and .SH NAME"""
							πF.SetLineno(384)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßheader_written, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 387: if self.header_written:
							πF.SetLineno(387)
						Label1:
							// line 388: return
							πF.SetLineno(388)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 389: self.head.append(self.header())
							πF.SetLineno(389)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßheader, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 390: self.head.append(MACRO_DEF)
							πF.SetLineno(390)
							πTemp003 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMACRO_DEF); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 391: self.header_written = 1
							πF.SetLineno(391)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßheader_written, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßappend_header.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 384: """append header with .TH and .SH NAME"""
					πF.SetLineno(384)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("append header with .TH and .SH NAME").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßappend_header); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 393: def visit_address(self, node):
					πF.SetLineno(393)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("visit_address", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 394: self.visit_docinfo_item(node, 'address')
							πF.SetLineno(394)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßaddress.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_address.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 396: def depart_address(self, node):
					πF.SetLineno(396)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("depart_address", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 397: pass
							πF.SetLineno(397)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_address.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 399: def visit_admonition(self, node, name=None):
					πF.SetLineno(399)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "name", Def: πTemp018}
					πTemp017 = πg.NewFunction(πg.NewCode("visit_admonition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µname *πg.Object = πArgs[2]
						_ = µname
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
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
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µname); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 408: if name:
							πF.SetLineno(408)
						Label1:
							// line 410: self.body.append('.sp\n')
							πF.SetLineno(410)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(".sp\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 411: name = '%s%s:%s\n' % (
							πF.SetLineno(411)
							πTemp005 = πg.NewInt(0).ToObject()
							πTemp007 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[0] = µname
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßget, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßupper, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							πTemp009 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s%s:%s\n").ToObject(), πTemp004); πE != nil {
								continue
							}
							µname = πTemp003
							// line 416: self.body.append(name)
							πF.SetLineno(416)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							// line 417: self.visit_block_quote(node)
							πF.SetLineno(417)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisit_block_quote, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_admonition.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 419: def depart_admonition(self, node):
					πF.SetLineno(419)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("depart_admonition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 420: self.depart_block_quote(node)
							πF.SetLineno(420)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdepart_block_quote, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_admonition.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 422: def visit_attention(self, node):
					πF.SetLineno(422)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("visit_attention", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 423: self.visit_admonition(node, 'attention')
							πF.SetLineno(423)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßattention.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_attention.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 425: depart_attention = depart_admonition
					πF.SetLineno(425)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_attention.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 427: def visit_docinfo_item(self, node, name):
					πF.SetLineno(427)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp004[2] = πg.Param{Name: "name", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("visit_docinfo_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µname *πg.Object = πArgs[2]
						_ = µname
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
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µname, ßauthor.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 428: if name == 'author':
							πF.SetLineno(428)
						Label1:
							// line 429: self._docinfo[name].append(node.astext())
							πF.SetLineno(429)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 431: self._docinfo[name] = node.astext()
							πF.SetLineno(431)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006 = µname
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 432: self._docinfo_keys.append(name)
							πF.SetLineno(432)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_docinfo_keys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 433: raise nodes.SkipNode
							πF.SetLineno(433)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_docinfo_item.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 435: def depart_docinfo_item(self, node):
					πF.SetLineno(435)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("depart_docinfo_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 436: pass
							πF.SetLineno(436)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_docinfo_item.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 438: def visit_author(self, node):
					πF.SetLineno(438)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("visit_author", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 439: self.visit_docinfo_item(node, 'author')
							πF.SetLineno(439)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_author.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 441: depart_author = depart_docinfo_item
					πF.SetLineno(441)
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_docinfo_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_author.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 443: def visit_authors(self, node):
					πF.SetLineno(443)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("visit_authors", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 445: pass
							πF.SetLineno(445)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_authors.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 447: def depart_authors(self, node):
					πF.SetLineno(447)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("depart_authors", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 448: pass
							πF.SetLineno(448)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_authors.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 450: def visit_block_quote(self, node):
					πF.SetLineno(450)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("visit_block_quote", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 453: self.indent(BLOCKQOUTE_INDENT)
							πF.SetLineno(453)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBLOCKQOUTE_INDENT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 454: self.indent(0)
							πF.SetLineno(454)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_block_quote.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 456: def depart_block_quote(self, node):
					πF.SetLineno(456)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("depart_block_quote", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 457: self.dedent()
							πF.SetLineno(457)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 458: self.dedent()
							πF.SetLineno(458)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_block_quote.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 460: def visit_bullet_list(self, node):
					πF.SetLineno(460)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("visit_bullet_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 461: self.list_start(node)
							πF.SetLineno(461)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlist_start, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_bullet_list.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 463: def depart_bullet_list(self, node):
					πF.SetLineno(463)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("depart_bullet_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 464: self.list_end()
							πF.SetLineno(464)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlist_end, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_bullet_list.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 466: def visit_caption(self, node):
					πF.SetLineno(466)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("visit_caption", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 467: pass
							πF.SetLineno(467)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_caption.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 469: def depart_caption(self, node):
					πF.SetLineno(469)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("depart_caption", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 470: pass
							πF.SetLineno(470)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_caption.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 472: def visit_caution(self, node):
					πF.SetLineno(472)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("visit_caution", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 473: self.visit_admonition(node, 'caution')
							πF.SetLineno(473)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßcaution.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_caution.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 475: depart_caution = depart_admonition
					πF.SetLineno(475)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_caution.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 477: def visit_citation(self, node):
					πF.SetLineno(477)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("visit_citation", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µnum *πg.Object = πg.UnboundLocal
						_ = µnum
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
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
							// line 478: num, text = node.astext().split(None, 1)
							πF.SetLineno(478)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µnum = πTemp002
							µtext = πTemp004
							// line 479: num = num.strip()
							πF.SetLineno(479)
							if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnum, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnum = πTemp003
							// line 480: self.body.append('.IP [%s] 5\n' % num)
							πF.SetLineno(480)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(".IP [%s] 5\n").ToObject(), µnum); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_citation.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 482: def depart_citation(self, node):
					πF.SetLineno(482)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("depart_citation", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 483: pass
							πF.SetLineno(483)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_citation.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 485: def visit_citation_reference(self, node):
					πF.SetLineno(485)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("visit_citation_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 486: self.body.append('['+node.astext()+']')
							πF.SetLineno(486)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πg.NewStr("[").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("]").ToObject()); πE != nil {
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 487: raise nodes.SkipNode
							πF.SetLineno(487)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_citation_reference.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 489: def visit_classifier(self, node):
					πF.SetLineno(489)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("visit_classifier", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 490: pass
							πF.SetLineno(490)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_classifier.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 492: def depart_classifier(self, node):
					πF.SetLineno(492)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("depart_classifier", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 493: pass
							πF.SetLineno(493)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_classifier.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 495: def visit_colspec(self, node):
					πF.SetLineno(495)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("visit_colspec", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 496: self.colspecs.append(node)
							πF.SetLineno(496)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_colspec.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 498: def depart_colspec(self, node):
					πF.SetLineno(498)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("depart_colspec", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 499: pass
							πF.SetLineno(499)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_colspec.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 501: def write_colspecs(self):
					πF.SetLineno(501)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("write_colspecs", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 502: self.body.append("%s.\n" % ('L '*len(self.colspecs)))
							πF.SetLineno(502)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcolspecs, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Mul(πF, πg.NewStr("L ").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s.\n").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwrite_colspecs.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 504: def visit_comment(self, node,
					πF.SetLineno(504)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("-(?=-)").ToObject()
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp042, πE = πg.GetAttr(πF, πTemp041, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp041, πE = πTemp042.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp042, πE = πg.GetAttr(πF, πTemp041, ßsub, nil); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "sub", Def: πTemp042}
					πTemp040 = πg.NewFunction(πg.NewCode("visit_comment", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µsub *πg.Object = πArgs[2]
						_ = µsub
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
							// line 506: self.body.append(self.comment(node.astext()))
							πF.SetLineno(506)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcomment, nil); πE != nil {
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
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 507: raise nodes.SkipNode
							πF.SetLineno(507)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_comment.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 509: def visit_contact(self, node):
					πF.SetLineno(509)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("visit_contact", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 510: self.visit_docinfo_item(node, 'contact')
							πF.SetLineno(510)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßcontact.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_contact.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 512: depart_contact = depart_docinfo_item
					πF.SetLineno(512)
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_docinfo_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_contact.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 514: def visit_container(self, node):
					πF.SetLineno(514)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp042 = πg.NewFunction(πg.NewCode("visit_container", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 515: pass
							πF.SetLineno(515)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_container.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 517: def depart_container(self, node):
					πF.SetLineno(517)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp043 = πg.NewFunction(πg.NewCode("depart_container", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 518: pass
							πF.SetLineno(518)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_container.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 520: def visit_compound(self, node):
					πF.SetLineno(520)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp044 = πg.NewFunction(πg.NewCode("visit_compound", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 521: pass
							πF.SetLineno(521)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_compound.ToObject(), πTemp044); πE != nil {
						continue
					}
					// line 523: def depart_compound(self, node):
					πF.SetLineno(523)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp045 = πg.NewFunction(πg.NewCode("depart_compound", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 524: pass
							πF.SetLineno(524)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_compound.ToObject(), πTemp045); πE != nil {
						continue
					}
					// line 526: def visit_copyright(self, node):
					πF.SetLineno(526)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp046 = πg.NewFunction(πg.NewCode("visit_copyright", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 527: self.visit_docinfo_item(node, 'copyright')
							πF.SetLineno(527)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßcopyright.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_copyright.ToObject(), πTemp046); πE != nil {
						continue
					}
					// line 529: def visit_danger(self, node):
					πF.SetLineno(529)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp047 = πg.NewFunction(πg.NewCode("visit_danger", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 530: self.visit_admonition(node, 'danger')
							πF.SetLineno(530)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßdanger.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_danger.ToObject(), πTemp047); πE != nil {
						continue
					}
					// line 532: depart_danger = depart_admonition
					πF.SetLineno(532)
					if πTemp048, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_danger.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 534: def visit_date(self, node):
					πF.SetLineno(534)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp048 = πg.NewFunction(πg.NewCode("visit_date", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 535: self.visit_docinfo_item(node, 'date')
							πF.SetLineno(535)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßdate.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_date.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 537: def visit_decoration(self, node):
					πF.SetLineno(537)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp049 = πg.NewFunction(πg.NewCode("visit_decoration", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 538: pass
							πF.SetLineno(538)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_decoration.ToObject(), πTemp049); πE != nil {
						continue
					}
					// line 540: def depart_decoration(self, node):
					πF.SetLineno(540)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp050 = πg.NewFunction(πg.NewCode("depart_decoration", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 541: pass
							πF.SetLineno(541)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_decoration.ToObject(), πTemp050); πE != nil {
						continue
					}
					// line 543: def visit_definition(self, node):
					πF.SetLineno(543)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp051 = πg.NewFunction(πg.NewCode("visit_definition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 544: pass
							πF.SetLineno(544)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_definition.ToObject(), πTemp051); πE != nil {
						continue
					}
					// line 546: def depart_definition(self, node):
					πF.SetLineno(546)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp052 = πg.NewFunction(πg.NewCode("depart_definition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 547: pass
							πF.SetLineno(547)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_definition.ToObject(), πTemp052); πE != nil {
						continue
					}
					// line 549: def visit_definition_list(self, node):
					πF.SetLineno(549)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp053 = πg.NewFunction(πg.NewCode("visit_definition_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 550: self.indent(DEFINITION_LIST_INDENT)
							πF.SetLineno(550)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßDEFINITION_LIST_INDENT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_definition_list.ToObject(), πTemp053); πE != nil {
						continue
					}
					// line 552: def depart_definition_list(self, node):
					πF.SetLineno(552)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp054 = πg.NewFunction(πg.NewCode("depart_definition_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 553: self.dedent()
							πF.SetLineno(553)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_definition_list.ToObject(), πTemp054); πE != nil {
						continue
					}
					// line 555: def visit_definition_list_item(self, node):
					πF.SetLineno(555)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp055 = πg.NewFunction(πg.NewCode("visit_definition_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 556: self.body.append(self.defs['definition_list_item'][0])
							πF.SetLineno(556)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßdefinition_list_item.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_definition_list_item.ToObject(), πTemp055); πE != nil {
						continue
					}
					// line 558: def depart_definition_list_item(self, node):
					πF.SetLineno(558)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp056 = πg.NewFunction(πg.NewCode("depart_definition_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 559: self.body.append(self.defs['definition_list_item'][1])
							πF.SetLineno(559)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßdefinition_list_item.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_definition_list_item.ToObject(), πTemp056); πE != nil {
						continue
					}
					// line 561: def visit_description(self, node):
					πF.SetLineno(561)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp057 = πg.NewFunction(πg.NewCode("visit_description", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 562: pass
							πF.SetLineno(562)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_description.ToObject(), πTemp057); πE != nil {
						continue
					}
					// line 564: def depart_description(self, node):
					πF.SetLineno(564)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp058 = πg.NewFunction(πg.NewCode("depart_description", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 565: pass
							πF.SetLineno(565)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_description.ToObject(), πTemp058); πE != nil {
						continue
					}
					// line 567: def visit_docinfo(self, node):
					πF.SetLineno(567)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp059 = πg.NewFunction(πg.NewCode("visit_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 568: self._in_docinfo = 1
							πF.SetLineno(568)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_docinfo, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_docinfo.ToObject(), πTemp059); πE != nil {
						continue
					}
					// line 570: def depart_docinfo(self, node):
					πF.SetLineno(570)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp060 = πg.NewFunction(πg.NewCode("depart_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 571: self._in_docinfo = None
							πF.SetLineno(571)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_docinfo, πTemp002); πE != nil {
								continue
							}
							// line 573: self.append_header()
							πF.SetLineno(573)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend_header, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_docinfo.ToObject(), πTemp060); πE != nil {
						continue
					}
					// line 575: def visit_doctest_block(self, node):
					πF.SetLineno(575)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp061 = πg.NewFunction(πg.NewCode("visit_doctest_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 576: self.body.append(self.defs['literal_block'][0])
							πF.SetLineno(576)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßliteral_block.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
							// line 577: self._in_literal = True
							πF.SetLineno(577)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_doctest_block.ToObject(), πTemp061); πE != nil {
						continue
					}
					// line 579: def depart_doctest_block(self, node):
					πF.SetLineno(579)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp062 = πg.NewFunction(πg.NewCode("depart_doctest_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 580: self._in_literal = False
							πF.SetLineno(580)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp002); πE != nil {
								continue
							}
							// line 581: self.body.append(self.defs['literal_block'][1])
							πF.SetLineno(581)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp004 = ßliteral_block.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßdepart_doctest_block.ToObject(), πTemp062); πE != nil {
						continue
					}
					// line 583: def visit_document(self, node):
					πF.SetLineno(583)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp063 = πg.NewFunction(πg.NewCode("visit_document", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 585: self.head.append(self.comment(self.document_start).rstrip()+'\n')
							πF.SetLineno(585)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument_start, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcomment, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 587: self.header_written = 0
							πF.SetLineno(587)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßheader_written, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_document.ToObject(), πTemp063); πE != nil {
						continue
					}
					// line 589: def depart_document(self, node):
					πF.SetLineno(589)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp064 = πg.NewFunction(πg.NewCode("depart_document", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µskip *πg.Object = πg.UnboundLocal
						_ = µskip
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µlabel *πg.Object = πg.UnboundLocal
						_ = µlabel
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
						var πTemp015 *πg.Object
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
						var πTemp017 *πg.Object
						_ = πTemp017
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
							πTemp001 = ßauthor.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 590: if self._docinfo['author']:
							πF.SetLineno(590)
						Label1:
							// line 591: self.body.append('.SH AUTHOR\n%s\n'
							πF.SetLineno(591)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp002 = ßauthor.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(".SH AUTHOR\n%s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
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
							goto Label2
						Label2:
							// line 593: skip = ('author', 'copyright', 'date',
							πF.SetLineno(593)
							πTemp005 = make([]*πg.Object, 9)
							πTemp005[0] = ßauthor.ToObject()
							πTemp005[1] = ßcopyright.ToObject()
							πTemp005[2] = ßdate.ToObject()
							πTemp005[3] = ßmanual_group.ToObject()
							πTemp005[4] = ßmanual_section.ToObject()
							πTemp005[5] = ßsubtitle.ToObject()
							πTemp005[6] = ßtitle.ToObject()
							πTemp005[7] = ßtitle_upper.ToObject()
							πTemp005[8] = ßversion.ToObject()
							πTemp001 = πg.NewTuple(πTemp005...).ToObject()
							µskip = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_docinfo_keys, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp004 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µname = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µname, ßaddress.ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µskip, µname); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 598: if name == 'address':
							πF.SetLineno(598)
						Label6:
							// line 599: self.body.append("\n%s:\n%s%s.nf\n%s\n.fi\n%s%s" % (
							πF.SetLineno(599)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp009, ßget, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp010 = πg.NewInt(0).ToObject()
							πTemp012 = ßindent.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, πTemp014, πTemp012); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp010); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mod(πF, πTemp011, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp011 = πg.NewInt(0).ToObject()
							πTemp013 = ßindent.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, πTemp015, πTemp013); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßBLOCKQOUTE_INDENT); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Mod(πF, πTemp012, πTemp011); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp011 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp013, πTemp011); πE != nil {
								continue
							}
							πTemp011 = πg.NewInt(1).ToObject()
							πTemp014 = ßindent.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetItem(πF, πTemp016, πTemp014); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, πTemp015, πTemp011); πE != nil {
								continue
							}
							πTemp011 = πg.NewInt(1).ToObject()
							πTemp015 = ßindent.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetItem(πF, πTemp017, πTemp015); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp011); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple6(πTemp009, πTemp007, πTemp010, πTemp012, πTemp013, πTemp014).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n%s:\n%s%s.nf\n%s\n.fi\n%s%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label8
							// line 606: elif not name in skip:
							πF.SetLineno(606)
						Label7:
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_docinfo_names, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 607: if name in self._docinfo_names:
							πF.SetLineno(607)
						Label9:
							// line 608: label = self._docinfo_names[name]
							πF.SetLineno(608)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_docinfo_names, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µlabel = πTemp003
							goto Label11
						Label10:
							// line 610: label = self.language.labels.get(name, name)
							πF.SetLineno(610)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[0] = µname
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µlabel = πTemp003
							goto Label11
						Label11:
							// line 611: self.body.append("\n%s: %s\n" % (label, self._docinfo[name]))
							πF.SetLineno(611)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µlabel, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n%s: %s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label8
						Label8:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							πTemp001 = ßcopyright.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 612: if self._docinfo['copyright']:
							πF.SetLineno(612)
						Label12:
							// line 613: self.body.append('.SH COPYRIGHT\n%s\n'
							πF.SetLineno(613)
							πTemp005 = πF.MakeArgs(1)
							πTemp002 = ßcopyright.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(".SH COPYRIGHT\n%s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
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
							goto Label13
						Label13:
							// line 615: self.body.append(self.comment(
							πF.SetLineno(615)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("Generated by docutils manpage writer.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcomment, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_document.ToObject(), πTemp064); πE != nil {
						continue
					}
					// line 618: def visit_emphasis(self, node):
					πF.SetLineno(618)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp065 = πg.NewFunction(πg.NewCode("visit_emphasis", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 619: self.body.append(self.defs['emphasis'][0])
							πF.SetLineno(619)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßemphasis.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_emphasis.ToObject(), πTemp065); πE != nil {
						continue
					}
					// line 621: def depart_emphasis(self, node):
					πF.SetLineno(621)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp066 = πg.NewFunction(πg.NewCode("depart_emphasis", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 622: self.body.append(self.defs['emphasis'][1])
							πF.SetLineno(622)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßemphasis.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_emphasis.ToObject(), πTemp066); πE != nil {
						continue
					}
					// line 624: def visit_entry(self, node):
					πF.SetLineno(624)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp067 = πg.NewFunction(πg.NewCode("visit_entry", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnode, ßmorerows.ToObject()); πE != nil {
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
							// line 626: if 'morerows' in node:
							πF.SetLineno(626)
						Label1:
							// line 627: self.document.reporter.warning('"table row spanning" not supported',
							πF.SetLineno(627)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\"table row spanning\" not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnode, ßmorecols.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 629: if 'morecols' in node:
							πF.SetLineno(629)
						Label3:
							// line 630: self.document.reporter.warning(
							πF.SetLineno(630)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\"table cell spanning\" not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
						Label4:
							// line 632: self.context.append(len(self.body))
							πF.SetLineno(632)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_entry.ToObject(), πTemp067); πE != nil {
						continue
					}
					// line 634: def depart_entry(self, node):
					πF.SetLineno(634)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp068 = πg.NewFunction(πg.NewCode("depart_entry", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 635: start = self.context.pop()
							πF.SetLineno(635)
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
							µstart = πTemp001
							// line 636: self._active_table.append_cell(self.body[start:])
							πF.SetLineno(636)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_active_table, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend_cell, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 637: del self.body[start:]
							πF.SetLineno(637)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_entry.ToObject(), πTemp068); πE != nil {
						continue
					}
					// line 639: def visit_enumerated_list(self, node):
					πF.SetLineno(639)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp069 = πg.NewFunction(πg.NewCode("visit_enumerated_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 640: self.list_start(node)
							πF.SetLineno(640)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlist_start, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_enumerated_list.ToObject(), πTemp069); πE != nil {
						continue
					}
					// line 642: def depart_enumerated_list(self, node):
					πF.SetLineno(642)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp070 = πg.NewFunction(πg.NewCode("depart_enumerated_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 643: self.list_end()
							πF.SetLineno(643)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlist_end, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_enumerated_list.ToObject(), πTemp070); πE != nil {
						continue
					}
					// line 645: def visit_error(self, node):
					πF.SetLineno(645)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp071 = πg.NewFunction(πg.NewCode("visit_error", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 646: self.visit_admonition(node, 'error')
							πF.SetLineno(646)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßerror.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_error.ToObject(), πTemp071); πE != nil {
						continue
					}
					// line 648: depart_error = depart_admonition
					πF.SetLineno(648)
					if πTemp072, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_error.ToObject(), πTemp072); πE != nil {
						continue
					}
					// line 650: def visit_field(self, node):
					πF.SetLineno(650)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp072 = πg.NewFunction(πg.NewCode("visit_field", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 651: pass
							πF.SetLineno(651)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field.ToObject(), πTemp072); πE != nil {
						continue
					}
					// line 653: def depart_field(self, node):
					πF.SetLineno(653)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp073 = πg.NewFunction(πg.NewCode("depart_field", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 654: pass
							πF.SetLineno(654)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field.ToObject(), πTemp073); πE != nil {
						continue
					}
					// line 656: def visit_field_body(self, node):
					πF.SetLineno(656)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp074 = πg.NewFunction(πg.NewCode("visit_field_body", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µname_normalized *πg.Object = πg.UnboundLocal
						_ = µname_normalized
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_in_docinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 657: if self._in_docinfo:
							πF.SetLineno(657)
						Label1:
							// line 658: name_normalized = self._field_name.lower().replace(" ", "_")
							πF.SetLineno(658)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr(" ").ToObject()
							πTemp003[1] = ß_.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_field_name, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßlower, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µname_normalized = πTemp001
							// line 659: self._docinfo_names[name_normalized] = self._field_name
							πF.SetLineno(659)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_field_name, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_docinfo_names, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname_normalized, "name_normalized"); πE != nil {
								continue
							}
							πTemp006 = µname_normalized
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp004); πE != nil {
								continue
							}
							// line 660: self.visit_docinfo_item(node, name_normalized)
							πF.SetLineno(660)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µname_normalized, "name_normalized"); πE != nil {
								continue
							}
							πTemp003[1] = µname_normalized
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 661: raise nodes.SkipNode
							πF.SetLineno(661)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßvisit_field_body.ToObject(), πTemp074); πE != nil {
						continue
					}
					// line 663: def depart_field_body(self, node):
					πF.SetLineno(663)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp075 = πg.NewFunction(πg.NewCode("depart_field_body", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 664: pass
							πF.SetLineno(664)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field_body.ToObject(), πTemp075); πE != nil {
						continue
					}
					// line 666: def visit_field_list(self, node):
					πF.SetLineno(666)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp076 = πg.NewFunction(πg.NewCode("visit_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 667: self.indent(FIELD_LIST_INDENT)
							πF.SetLineno(667)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFIELD_LIST_INDENT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_field_list.ToObject(), πTemp076); πE != nil {
						continue
					}
					// line 669: def depart_field_list(self, node):
					πF.SetLineno(669)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp077 = πg.NewFunction(πg.NewCode("depart_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 670: self.dedent()
							πF.SetLineno(670)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_field_list.ToObject(), πTemp077); πE != nil {
						continue
					}
					// line 672: def visit_field_name(self, node):
					πF.SetLineno(672)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp078 = πg.NewFunction(πg.NewCode("visit_field_name", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_in_docinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 673: if self._in_docinfo:
							πF.SetLineno(673)
						Label1:
							// line 674: self._field_name = node.astext()
							πF.SetLineno(674)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_field_name, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 675: raise nodes.SkipNode
							πF.SetLineno(675)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label2:
							// line 677: self.body.append(self.defs['field_name'][0])
							πF.SetLineno(677)
							πTemp004 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp005 = ßfield_name.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßvisit_field_name.ToObject(), πTemp078); πE != nil {
						continue
					}
					// line 679: def depart_field_name(self, node):
					πF.SetLineno(679)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp079 = πg.NewFunction(πg.NewCode("depart_field_name", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 680: self.body.append(self.defs['field_name'][1])
							πF.SetLineno(680)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßfield_name.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_field_name.ToObject(), πTemp079); πE != nil {
						continue
					}
					// line 682: def visit_figure(self, node):
					πF.SetLineno(682)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp080 = πg.NewFunction(πg.NewCode("visit_figure", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 683: self.indent(2.5)
							πF.SetLineno(683)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewFloat(2.5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 684: self.indent(0)
							πF.SetLineno(684)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_figure.ToObject(), πTemp080); πE != nil {
						continue
					}
					// line 686: def depart_figure(self, node):
					πF.SetLineno(686)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp081 = πg.NewFunction(πg.NewCode("depart_figure", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 687: self.dedent()
							πF.SetLineno(687)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 688: self.dedent()
							πF.SetLineno(688)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_figure.ToObject(), πTemp081); πE != nil {
						continue
					}
					// line 690: def visit_footer(self, node):
					πF.SetLineno(690)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp082 = πg.NewFunction(πg.NewCode("visit_footer", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 691: self.document.reporter.warning('"footer" not supported',
							πF.SetLineno(691)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"footer\" not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_footer.ToObject(), πTemp082); πE != nil {
						continue
					}
					// line 694: def depart_footer(self, node):
					πF.SetLineno(694)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp083 = πg.NewFunction(πg.NewCode("depart_footer", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 695: pass
							πF.SetLineno(695)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_footer.ToObject(), πTemp083); πE != nil {
						continue
					}
					// line 697: def visit_footnote(self, node):
					πF.SetLineno(697)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp084 = πg.NewFunction(πg.NewCode("visit_footnote", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µnum *πg.Object = πg.UnboundLocal
						_ = µnum
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 698: num, text = node.astext().split(None, 1)
							πF.SetLineno(698)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µnum = πTemp002
							µtext = πTemp004
							// line 699: num = num.strip()
							πF.SetLineno(699)
							if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnum, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnum = πTemp003
							// line 700: self.body.append('.IP [%s] 5\n' % self.deunicode(num))
							πF.SetLineno(700)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnum, "num"); πE != nil {
								continue
							}
							πTemp005[0] = µnum
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdeunicode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(".IP [%s] 5\n").ToObject(), πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_footnote.ToObject(), πTemp084); πE != nil {
						continue
					}
					// line 702: def depart_footnote(self, node):
					πF.SetLineno(702)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp085 = πg.NewFunction(πg.NewCode("depart_footnote", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 703: pass
							πF.SetLineno(703)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_footnote.ToObject(), πTemp085); πE != nil {
						continue
					}
					// line 705: def footnote_backrefs(self, node):
					πF.SetLineno(705)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp086 = πg.NewFunction(πg.NewCode("footnote_backrefs", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 706: self.document.reporter.warning('"footnote_backrefs" not supported',
							πF.SetLineno(706)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"footnote_backrefs\" not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfootnote_backrefs.ToObject(), πTemp086); πE != nil {
						continue
					}
					// line 709: def visit_footnote_reference(self, node):
					πF.SetLineno(709)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp087 = πg.NewFunction(πg.NewCode("visit_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 710: self.body.append('['+self.deunicode(node.astext())+']')
							πF.SetLineno(710)
							πTemp001 = πF.MakeArgs(1)
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
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdeunicode, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, πg.NewStr("[").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("]").ToObject()); πE != nil {
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 711: raise nodes.SkipNode
							πF.SetLineno(711)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_footnote_reference.ToObject(), πTemp087); πE != nil {
						continue
					}
					// line 713: def depart_footnote_reference(self, node):
					πF.SetLineno(713)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp088 = πg.NewFunction(πg.NewCode("depart_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 714: pass
							πF.SetLineno(714)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_footnote_reference.ToObject(), πTemp088); πE != nil {
						continue
					}
					// line 716: def visit_generated(self, node):
					πF.SetLineno(716)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp089 = πg.NewFunction(πg.NewCode("visit_generated", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 717: pass
							πF.SetLineno(717)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_generated.ToObject(), πTemp089); πE != nil {
						continue
					}
					// line 719: def depart_generated(self, node):
					πF.SetLineno(719)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp090 = πg.NewFunction(πg.NewCode("depart_generated", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 720: pass
							πF.SetLineno(720)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_generated.ToObject(), πTemp090); πE != nil {
						continue
					}
					// line 722: def visit_header(self, node):
					πF.SetLineno(722)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp091 = πg.NewFunction(πg.NewCode("visit_header", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 723: raise NotImplementedError(node.astext())
							πF.SetLineno(723)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_header.ToObject(), πTemp091); πE != nil {
						continue
					}
					// line 725: def depart_header(self, node):
					πF.SetLineno(725)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp092 = πg.NewFunction(πg.NewCode("depart_header", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 726: pass
							πF.SetLineno(726)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_header.ToObject(), πTemp092); πE != nil {
						continue
					}
					// line 728: def visit_hint(self, node):
					πF.SetLineno(728)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp093 = πg.NewFunction(πg.NewCode("visit_hint", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 729: self.visit_admonition(node, 'hint')
							πF.SetLineno(729)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßhint.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_hint.ToObject(), πTemp093); πE != nil {
						continue
					}
					// line 731: depart_hint = depart_admonition
					πF.SetLineno(731)
					if πTemp094, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_hint.ToObject(), πTemp094); πE != nil {
						continue
					}
					// line 733: def visit_subscript(self, node):
					πF.SetLineno(733)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp094 = πg.NewFunction(πg.NewCode("visit_subscript", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 734: self.body.append('\\s-2\\d')
							πF.SetLineno(734)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\\s-2\\d").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_subscript.ToObject(), πTemp094); πE != nil {
						continue
					}
					// line 736: def depart_subscript(self, node):
					πF.SetLineno(736)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp095 = πg.NewFunction(πg.NewCode("depart_subscript", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 737: self.body.append('\\u\\s0')
							πF.SetLineno(737)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\\u\\s0").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_subscript.ToObject(), πTemp095); πE != nil {
						continue
					}
					// line 739: def visit_superscript(self, node):
					πF.SetLineno(739)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp096 = πg.NewFunction(πg.NewCode("visit_superscript", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 740: self.body.append('\\s-2\\u')
							πF.SetLineno(740)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\\s-2\\u").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_superscript.ToObject(), πTemp096); πE != nil {
						continue
					}
					// line 742: def depart_superscript(self, node):
					πF.SetLineno(742)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp097 = πg.NewFunction(πg.NewCode("depart_superscript", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 743: self.body.append('\\d\\s0')
							πF.SetLineno(743)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\\d\\s0").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_superscript.ToObject(), πTemp097); πE != nil {
						continue
					}
					// line 745: def visit_attribution(self, node):
					πF.SetLineno(745)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp098 = πg.NewFunction(πg.NewCode("visit_attribution", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 746: self.body.append('\\(em ')
							πF.SetLineno(746)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\\(em ").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_attribution.ToObject(), πTemp098); πE != nil {
						continue
					}
					// line 748: def depart_attribution(self, node):
					πF.SetLineno(748)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp099 = πg.NewFunction(πg.NewCode("depart_attribution", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 749: self.body.append('\n')
							πF.SetLineno(749)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_attribution.ToObject(), πTemp099); πE != nil {
						continue
					}
					// line 751: def visit_image(self, node):
					πF.SetLineno(751)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp100 = πg.NewFunction(πg.NewCode("visit_image", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 752: self.document.reporter.warning('"image" not supported',
							πF.SetLineno(752)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"image\" not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 754: text = []
							πF.SetLineno(754)
							πTemp001 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							µtext = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, ßalt.ToObject()); πE != nil {
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
							// line 755: if 'alt' in node.attributes:
							πF.SetLineno(755)
						Label1:
							// line 756: text.append(node.attributes['alt'])
							πF.SetLineno(756)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = ßalt.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, ßuri.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 757: if 'uri' in node.attributes:
							πF.SetLineno(757)
						Label3:
							// line 758: text.append(node.attributes['uri'])
							πF.SetLineno(758)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = ßuri.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							// line 759: self.body.append('[image: %s]\n' % ('/'.join(text)))
							πF.SetLineno(759)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp007[0] = µtext
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("[image: %s]\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
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
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 760: raise nodes.SkipNode
							πF.SetLineno(760)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_image.ToObject(), πTemp100); πE != nil {
						continue
					}
					// line 762: def visit_important(self, node):
					πF.SetLineno(762)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp101 = πg.NewFunction(πg.NewCode("visit_important", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 763: self.visit_admonition(node, 'important')
							πF.SetLineno(763)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßimportant.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_important.ToObject(), πTemp101); πE != nil {
						continue
					}
					// line 765: depart_important = depart_admonition
					πF.SetLineno(765)
					if πTemp102, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_important.ToObject(), πTemp102); πE != nil {
						continue
					}
					// line 767: def visit_inline(self, node):
					πF.SetLineno(767)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp102 = πg.NewFunction(πg.NewCode("visit_inline", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 768: pass
							πF.SetLineno(768)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_inline.ToObject(), πTemp102); πE != nil {
						continue
					}
					// line 770: def depart_inline(self, node):
					πF.SetLineno(770)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp103 = πg.NewFunction(πg.NewCode("depart_inline", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 771: pass
							πF.SetLineno(771)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_inline.ToObject(), πTemp103); πE != nil {
						continue
					}
					// line 773: def visit_label(self, node):
					πF.SetLineno(773)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp104 = πg.NewFunction(πg.NewCode("visit_label", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfootnote, nil); πE != nil {
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
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcitation, nil); πE != nil {
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
							// line 775: if (isinstance(node.parent, nodes.footnote)
							πF.SetLineno(775)
						Label2:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 777: raise nodes.SkipNode
							πF.SetLineno(777)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label3
						Label3:
							// line 778: self.document.reporter.warning('"unsupported "label"',
							πF.SetLineno(778)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\"unsupported \"label\"").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 780: self.body.append('[')
							πF.SetLineno(780)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("[").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_label.ToObject(), πTemp104); πE != nil {
						continue
					}
					// line 782: def depart_label(self, node):
					πF.SetLineno(782)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp105 = πg.NewFunction(πg.NewCode("depart_label", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 783: self.body.append(']\n')
							πF.SetLineno(783)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("]\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_label.ToObject(), πTemp105); πE != nil {
						continue
					}
					// line 785: def visit_legend(self, node):
					πF.SetLineno(785)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp106 = πg.NewFunction(πg.NewCode("visit_legend", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 786: pass
							πF.SetLineno(786)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_legend.ToObject(), πTemp106); πE != nil {
						continue
					}
					// line 788: def depart_legend(self, node):
					πF.SetLineno(788)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp107 = πg.NewFunction(πg.NewCode("depart_legend", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 789: pass
							πF.SetLineno(789)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_legend.ToObject(), πTemp107); πE != nil {
						continue
					}
					// line 792: def visit_line_block(self, node):
					πF.SetLineno(792)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp108 = πg.NewFunction(πg.NewCode("visit_line_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 793: self._line_block += 1
							πF.SetLineno(793)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_line_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_line_block, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_line_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 794: if self._line_block == 1:
							πF.SetLineno(794)
						Label1:
							// line 799: self.body.append('.nf\n')
							πF.SetLineno(799)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(".nf\n").ToObject()
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
							goto Label3
						Label2:
							// line 801: self.body.append('.in +2\n')
							πF.SetLineno(801)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(".in +2\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_line_block.ToObject(), πTemp108); πE != nil {
						continue
					}
					// line 803: def depart_line_block(self, node):
					πF.SetLineno(803)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp109 = πg.NewFunction(πg.NewCode("depart_line_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 804: self._line_block -= 1
							πF.SetLineno(804)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_line_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_line_block, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_line_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 805: if self._line_block == 0:
							πF.SetLineno(805)
						Label1:
							// line 806: self.body.append('.fi\n')
							πF.SetLineno(806)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(".fi\n").ToObject()
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
							// line 807: self.body.append('.sp\n')
							πF.SetLineno(807)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(".sp\n").ToObject()
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
							goto Label3
						Label2:
							// line 809: self.body.append('.in -2\n')
							πF.SetLineno(809)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(".in -2\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_line_block.ToObject(), πTemp109); πE != nil {
						continue
					}
					// line 811: def visit_line(self, node):
					πF.SetLineno(811)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp110 = πg.NewFunction(πg.NewCode("visit_line", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 812: pass
							πF.SetLineno(812)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_line.ToObject(), πTemp110); πE != nil {
						continue
					}
					// line 814: def depart_line(self, node):
					πF.SetLineno(814)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp111 = πg.NewFunction(πg.NewCode("depart_line", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 815: self.body.append('\n')
							πF.SetLineno(815)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_line.ToObject(), πTemp111); πE != nil {
						continue
					}
					// line 817: def visit_list_item(self, node):
					πF.SetLineno(817)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp112 = πg.NewFunction(πg.NewCode("visit_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
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
							// line 819: self.body.append('.IP %s %d\n' % (
							πF.SetLineno(819)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_list_char, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßget_width, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(".IP %s %d\n").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_list_item.ToObject(), πTemp112); πE != nil {
						continue
					}
					// line 823: def depart_list_item(self, node):
					πF.SetLineno(823)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp113 = πg.NewFunction(πg.NewCode("depart_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 824: pass
							πF.SetLineno(824)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_list_item.ToObject(), πTemp113); πE != nil {
						continue
					}
					// line 826: def visit_literal(self, node):
					πF.SetLineno(826)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp114 = πg.NewFunction(πg.NewCode("visit_literal", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 827: self.body.append(self.defs['literal'][0])
							πF.SetLineno(827)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßliteral.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_literal.ToObject(), πTemp114); πE != nil {
						continue
					}
					// line 829: def depart_literal(self, node):
					πF.SetLineno(829)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp115 = πg.NewFunction(πg.NewCode("depart_literal", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 830: self.body.append(self.defs['literal'][1])
							πF.SetLineno(830)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßliteral.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_literal.ToObject(), πTemp115); πE != nil {
						continue
					}
					// line 832: def visit_literal_block(self, node):
					πF.SetLineno(832)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp116 = πg.NewFunction(πg.NewCode("visit_literal_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 835: self.indent(LITERAL_BLOCK_INDENT)
							πF.SetLineno(835)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßLITERAL_BLOCK_INDENT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 836: self.indent(0)
							πF.SetLineno(836)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 837: self.body.append(self.defs['literal_block'][0])
							πF.SetLineno(837)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßliteral_block.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
							// line 838: self._in_literal = True
							πF.SetLineno(838)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_literal_block.ToObject(), πTemp116); πE != nil {
						continue
					}
					// line 840: def depart_literal_block(self, node):
					πF.SetLineno(840)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp117 = πg.NewFunction(πg.NewCode("depart_literal_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 841: self._in_literal = False
							πF.SetLineno(841)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp002); πE != nil {
								continue
							}
							// line 842: self.body.append(self.defs['literal_block'][1])
							πF.SetLineno(842)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp004 = ßliteral_block.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
							// line 843: self.dedent()
							πF.SetLineno(843)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 844: self.dedent()
							πF.SetLineno(844)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_literal_block.ToObject(), πTemp117); πE != nil {
						continue
					}
					// line 846: def visit_math(self, node):
					πF.SetLineno(846)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp118 = πg.NewFunction(πg.NewCode("visit_math", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 847: self.document.reporter.warning('"math" role not supported',
							πF.SetLineno(847)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"math\" role not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 849: self.visit_literal(node)
							πF.SetLineno(849)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisit_literal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_math.ToObject(), πTemp118); πE != nil {
						continue
					}
					// line 851: def depart_math(self, node):
					πF.SetLineno(851)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp119 = πg.NewFunction(πg.NewCode("depart_math", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 852: self.depart_literal(node)
							πF.SetLineno(852)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdepart_literal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_math.ToObject(), πTemp119); πE != nil {
						continue
					}
					// line 854: def visit_math_block(self, node):
					πF.SetLineno(854)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp120 = πg.NewFunction(πg.NewCode("visit_math_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 855: self.document.reporter.warning('"math" directive not supported',
							πF.SetLineno(855)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"math\" directive not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 857: self.visit_literal_block(node)
							πF.SetLineno(857)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisit_literal_block, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_math_block.ToObject(), πTemp120); πE != nil {
						continue
					}
					// line 859: def depart_math_block(self, node):
					πF.SetLineno(859)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp121 = πg.NewFunction(πg.NewCode("depart_math_block", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 860: self.depart_literal_block(node)
							πF.SetLineno(860)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdepart_literal_block, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_math_block.ToObject(), πTemp121); πE != nil {
						continue
					}
					// line 862: def visit_meta(self, node):
					πF.SetLineno(862)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp122 = πg.NewFunction(πg.NewCode("visit_meta", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 863: raise NotImplementedError(node.astext())
							πF.SetLineno(863)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_meta.ToObject(), πTemp122); πE != nil {
						continue
					}
					// line 865: def depart_meta(self, node):
					πF.SetLineno(865)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp123 = πg.NewFunction(πg.NewCode("depart_meta", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 866: pass
							πF.SetLineno(866)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_meta.ToObject(), πTemp123); πE != nil {
						continue
					}
					// line 868: def visit_note(self, node):
					πF.SetLineno(868)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp124 = πg.NewFunction(πg.NewCode("visit_note", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 869: self.visit_admonition(node, 'note')
							πF.SetLineno(869)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßnote.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_note.ToObject(), πTemp124); πE != nil {
						continue
					}
					// line 871: depart_note = depart_admonition
					πF.SetLineno(871)
					if πTemp125, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_note.ToObject(), πTemp125); πE != nil {
						continue
					}
					// line 873: def indent(self, by=0.5):
					πF.SetLineno(873)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "by", Def: πg.NewFloat(0.5).ToObject()}
					πTemp125 = πg.NewFunction(πg.NewCode("indent", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µby *πg.Object = πArgs[1]
						_ = µby
						var µstep *πg.Object = πg.UnboundLocal
						_ = µstep
						var πTemp001 *πg.Object
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
							// line 875: step = self._indent[-1]
							πF.SetLineno(875)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_indent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µstep = πTemp002
							// line 876: self._indent.append(by)
							πF.SetLineno(876)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µby, "by"); πE != nil {
								continue
							}
							πTemp004[0] = µby
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_indent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 877: self.body.append(self.defs['indent'][0] % step)
							πF.SetLineno(877)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp005 = ßindent.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp003, µstep); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
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
					if πE = πClass.SetItem(πF, ßindent.ToObject(), πTemp125); πE != nil {
						continue
					}
					// line 879: def dedent(self):
					πF.SetLineno(879)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp126 = πg.NewFunction(πg.NewCode("dedent", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 880: self._indent.pop()
							πF.SetLineno(880)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_indent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 881: self.body.append(self.defs['indent'][1])
							πF.SetLineno(881)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp004 = ßindent.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßdedent.ToObject(), πTemp126); πE != nil {
						continue
					}
					// line 883: def visit_option_list(self, node):
					πF.SetLineno(883)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp127 = πg.NewFunction(πg.NewCode("visit_option_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 884: self.indent(OPTION_LIST_INDENT)
							πF.SetLineno(884)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOPTION_LIST_INDENT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_option_list.ToObject(), πTemp127); πE != nil {
						continue
					}
					// line 886: def depart_option_list(self, node):
					πF.SetLineno(886)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp128 = πg.NewFunction(πg.NewCode("depart_option_list", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 887: self.dedent()
							πF.SetLineno(887)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_option_list.ToObject(), πTemp128); πE != nil {
						continue
					}
					// line 889: def visit_option_list_item(self, node):
					πF.SetLineno(889)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp129 = πg.NewFunction(πg.NewCode("visit_option_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 891: self.body.append(self.defs['option_list_item'][0])
							πF.SetLineno(891)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßoption_list_item.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_option_list_item.ToObject(), πTemp129); πE != nil {
						continue
					}
					// line 893: def depart_option_list_item(self, node):
					πF.SetLineno(893)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp130 = πg.NewFunction(πg.NewCode("depart_option_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 894: self.body.append(self.defs['option_list_item'][1])
							πF.SetLineno(894)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßoption_list_item.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_option_list_item.ToObject(), πTemp130); πE != nil {
						continue
					}
					// line 896: def visit_option_group(self, node):
					πF.SetLineno(896)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp131 = πg.NewFunction(πg.NewCode("visit_option_group", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 902: self.context.append('.B')           # blind guess
							πF.SetLineno(902)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".B").ToObject()
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
							// line 903: self.context.append(len(self.body)) # to be able to insert later
							πF.SetLineno(903)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
							// line 904: self.context.append(0)              # option counter
							πF.SetLineno(904)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_group.ToObject(), πTemp131); πE != nil {
						continue
					}
					// line 906: def depart_option_group(self, node):
					πF.SetLineno(906)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp132 = πg.NewFunction(πg.NewCode("depart_option_group", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart_position *πg.Object = πg.UnboundLocal
						_ = µstart_position
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							// line 907: self.context.pop()  # the counter
							πF.SetLineno(907)
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
							// line 908: start_position = self.context.pop()
							πF.SetLineno(908)
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
							µstart_position = πTemp001
							// line 909: text = self.body[start_position:]
							πF.SetLineno(909)
							if πE = πg.CheckLocal(πF, µstart_position, "start_position"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstart_position, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µtext = πTemp002
							// line 910: del self.body[start_position:]
							πF.SetLineno(910)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart_position, "start_position"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart_position, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 911: self.body.append('%s%s\n' % (self.context.pop(), ''.join(text)))
							πF.SetLineno(911)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp006[0] = µtext
							if πTemp005, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s%s\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
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
					if πE = πClass.SetItem(πF, ßdepart_option_group.ToObject(), πTemp132); πE != nil {
						continue
					}
					// line 913: def visit_option(self, node):
					πF.SetLineno(913)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp133 = πg.NewFunction(πg.NewCode("visit_option", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 915: if self.context[-1] > 0:
							πF.SetLineno(915)
						Label1:
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr(".BI").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 916: if self.context[-3] == '.BI':
							πF.SetLineno(916)
						Label3:
							// line 917: self.body.append('\\fR,\\fB ')
							πF.SetLineno(917)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\\fR,\\fB ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label5
						Label4:
							// line 919: self.body.append('\\fP,\\fB ')
							πF.SetLineno(919)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\\fP,\\fB ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label5
						Label5:
							goto Label2
						Label2:
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr(".BI").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 920: if self.context[-3] == '.BI':
							πF.SetLineno(920)
						Label6:
							// line 921: self.body.append('\\')
							πF.SetLineno(921)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\\").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label7
						Label7:
							// line 922: self.body.append(' ')
							πF.SetLineno(922)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_option.ToObject(), πTemp133); πE != nil {
						continue
					}
					// line 924: def depart_option(self, node):
					πF.SetLineno(924)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp134 = πg.NewFunction(πg.NewCode("depart_option", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 925: self.context[-1] += 1
							πF.SetLineno(925)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_option.ToObject(), πTemp134); πE != nil {
						continue
					}
					// line 927: def visit_option_string(self, node):
					πF.SetLineno(927)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp135 = πg.NewFunction(πg.NewCode("visit_option_string", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 929: pass
							πF.SetLineno(929)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_string.ToObject(), πTemp135); πE != nil {
						continue
					}
					// line 931: def depart_option_string(self, node):
					πF.SetLineno(931)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp136 = πg.NewFunction(πg.NewCode("depart_option_string", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 932: pass
							πF.SetLineno(932)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_string.ToObject(), πTemp136); πE != nil {
						continue
					}
					// line 934: def visit_option_argument(self, node):
					πF.SetLineno(934)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp137 = πg.NewFunction(πg.NewCode("visit_option_argument", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 935: self.context[-3] = '.BI' # bold/italic alternate
							πF.SetLineno(935)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr(".BI").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp002 = ßdelimiter.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp003, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("=").ToObject()
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 936: if node['delimiter'] != ' ':
							πF.SetLineno(936)
						Label1:
							// line 937: self.body.append('\\fB%s ' % node['delimiter'])
							πF.SetLineno(937)
							πTemp006 = πF.MakeArgs(1)
							πTemp002 = ßdelimiter.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\\fB%s ").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label4
							// line 938: elif self.body[len(self.body)-1].endswith('='):
							πF.SetLineno(938)
						Label2:
							// line 940: self.body.append(' ')
							πF.SetLineno(940)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label4
						Label3:
							// line 943: self.body.append(' \\ ')
							πF.SetLineno(943)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr(" \\ ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßvisit_option_argument.ToObject(), πTemp137); πE != nil {
						continue
					}
					// line 945: def depart_option_argument(self, node):
					πF.SetLineno(945)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp138 = πg.NewFunction(πg.NewCode("depart_option_argument", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 946: pass
							πF.SetLineno(946)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_argument.ToObject(), πTemp138); πE != nil {
						continue
					}
					// line 948: def visit_organization(self, node):
					πF.SetLineno(948)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp139 = πg.NewFunction(πg.NewCode("visit_organization", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 949: self.visit_docinfo_item(node, 'organization')
							πF.SetLineno(949)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßorganization.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_organization.ToObject(), πTemp139); πE != nil {
						continue
					}
					// line 951: def depart_organization(self, node):
					πF.SetLineno(951)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp140 = πg.NewFunction(πg.NewCode("depart_organization", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 952: pass
							πF.SetLineno(952)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_organization.ToObject(), πTemp140); πE != nil {
						continue
					}
					// line 954: def first_child(self, node):
					πF.SetLineno(954)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp141 = πg.NewFunction(πg.NewCode("first_child", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µfirst *πg.Object = πg.UnboundLocal
						_ = µfirst
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
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
							// line 955: first = isinstance(node.parent[0], nodes.label) # skip label
							πF.SetLineno(955)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlabel, nil); πE != nil {
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
							µfirst = πTemp003
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µfirst, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µchild = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp001[0] = µchild
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßInvisible, nil); πE != nil {
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
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label4
							}
							goto Label5
							// line 957: if isinstance(child, nodes.Invisible):
							πF.SetLineno(957)
						Label4:
							// line 958: continue
							πF.SetLineno(958)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µchild == µnode).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 959: if child is node:
							πF.SetLineno(959)
						Label6:
							// line 960: return 1
							πF.SetLineno(960)
							πR = πg.NewInt(1).ToObject()
							continue
							goto Label7
						Label7:
							// line 961: break
							πF.SetLineno(961)
							πTemp007 = true
							continue
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 962: return 0
							πF.SetLineno(962)
							πR = πg.NewInt(0).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfirst_child.ToObject(), πTemp141); πE != nil {
						continue
					}
					// line 964: def visit_paragraph(self, node):
					πF.SetLineno(964)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp142 = πg.NewFunction(πg.NewCode("visit_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 970: self.ensure_eol()
							πF.SetLineno(970)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßensure_eol, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfirst_child, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 971: if not self.first_child(node):
							πF.SetLineno(971)
						Label1:
							// line 972: self.body.append('.sp\n')
							πF.SetLineno(972)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr(".sp\n").ToObject()
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
							goto Label2
						Label2:
							// line 974: self._in_literal = True
							πF.SetLineno(974)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_paragraph.ToObject(), πTemp142); πE != nil {
						continue
					}
					// line 976: def depart_paragraph(self, node):
					πF.SetLineno(976)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp143 = πg.NewFunction(πg.NewCode("depart_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 977: self._in_literal = False
							πF.SetLineno(977)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_in_literal, πTemp002); πE != nil {
								continue
							}
							// line 978: self.body.append('\n')
							πF.SetLineno(978)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_paragraph.ToObject(), πTemp143); πE != nil {
						continue
					}
					// line 980: def visit_problematic(self, node):
					πF.SetLineno(980)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp144 = πg.NewFunction(πg.NewCode("visit_problematic", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 981: self.body.append(self.defs['problematic'][0])
							πF.SetLineno(981)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßproblematic.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_problematic.ToObject(), πTemp144); πE != nil {
						continue
					}
					// line 983: def depart_problematic(self, node):
					πF.SetLineno(983)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp145 = πg.NewFunction(πg.NewCode("depart_problematic", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 984: self.body.append(self.defs['problematic'][1])
							πF.SetLineno(984)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßproblematic.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_problematic.ToObject(), πTemp145); πE != nil {
						continue
					}
					// line 986: def visit_raw(self, node):
					πF.SetLineno(986)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp146 = πg.NewFunction(πg.NewCode("visit_raw", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßformat.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, ßmanpage.ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 987: if node.get('format') == 'manpage':
							πF.SetLineno(987)
						Label1:
							// line 988: self.body.append(node.astext() + "\n")
							πF.SetLineno(988)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 990: raise nodes.SkipNode
							πF.SetLineno(990)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_raw.ToObject(), πTemp146); πE != nil {
						continue
					}
					// line 992: def visit_reference(self, node):
					πF.SetLineno(992)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp147 = πg.NewFunction(πg.NewCode("visit_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 993: """E.g. link or email address."""
							πF.SetLineno(993)
							// line 994: self.body.append(self.defs['reference'][0])
							πF.SetLineno(994)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßreference.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_reference.ToObject(), πTemp147); πE != nil {
						continue
					}
					// line 993: """E.g. link or email address."""
					πF.SetLineno(993)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp148}, πg.NewStr("E.g. link or email address.").ToObject()); πE != nil {
						continue
					}
					if πTemp149, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_reference); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp149, ß__doc__, πTemp148); πE != nil {
						continue
					}
					// line 996: def depart_reference(self, node):
					πF.SetLineno(996)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp148 = πg.NewFunction(πg.NewCode("depart_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 997: self.body.append(self.defs['reference'][1])
							πF.SetLineno(997)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßreference.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_reference.ToObject(), πTemp148); πE != nil {
						continue
					}
					// line 999: def visit_revision(self, node):
					πF.SetLineno(999)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp149 = πg.NewFunction(πg.NewCode("visit_revision", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1000: self.visit_docinfo_item(node, 'revision')
							πF.SetLineno(1000)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßrevision.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_revision.ToObject(), πTemp149); πE != nil {
						continue
					}
					// line 1002: depart_revision = depart_docinfo_item
					πF.SetLineno(1002)
					if πTemp150, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_docinfo_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_revision.ToObject(), πTemp150); πE != nil {
						continue
					}
					// line 1004: def visit_row(self, node):
					πF.SetLineno(1004)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp150 = πg.NewFunction(πg.NewCode("visit_row", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1005: self._active_table.new_row()
							πF.SetLineno(1005)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_active_table, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnew_row, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_row.ToObject(), πTemp150); πE != nil {
						continue
					}
					// line 1007: def depart_row(self, node):
					πF.SetLineno(1007)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp151 = πg.NewFunction(πg.NewCode("depart_row", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1008: pass
							πF.SetLineno(1008)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_row.ToObject(), πTemp151); πE != nil {
						continue
					}
					// line 1010: def visit_section(self, node):
					πF.SetLineno(1010)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp152 = πg.NewFunction(πg.NewCode("visit_section", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1011: self.section_level += 1
							πF.SetLineno(1011)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_level, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_section.ToObject(), πTemp152); πE != nil {
						continue
					}
					// line 1013: def depart_section(self, node):
					πF.SetLineno(1013)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp153 = πg.NewFunction(πg.NewCode("depart_section", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1014: self.section_level -= 1
							πF.SetLineno(1014)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_level, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_section.ToObject(), πTemp153); πE != nil {
						continue
					}
					// line 1016: def visit_status(self, node):
					πF.SetLineno(1016)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp154 = πg.NewFunction(πg.NewCode("visit_status", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1017: self.visit_docinfo_item(node, 'status')
							πF.SetLineno(1017)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßstatus.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_status.ToObject(), πTemp154); πE != nil {
						continue
					}
					// line 1019: depart_status = depart_docinfo_item
					πF.SetLineno(1019)
					if πTemp155, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_docinfo_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_status.ToObject(), πTemp155); πE != nil {
						continue
					}
					// line 1021: def visit_strong(self, node):
					πF.SetLineno(1021)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp155 = πg.NewFunction(πg.NewCode("visit_strong", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1022: self.body.append(self.defs['strong'][0])
							πF.SetLineno(1022)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_strong.ToObject(), πTemp155); πE != nil {
						continue
					}
					// line 1024: def depart_strong(self, node):
					πF.SetLineno(1024)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp156 = πg.NewFunction(πg.NewCode("depart_strong", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1025: self.body.append(self.defs['strong'][1])
							πF.SetLineno(1025)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_strong.ToObject(), πTemp156); πE != nil {
						continue
					}
					// line 1027: def visit_substitution_definition(self, node):
					πF.SetLineno(1027)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp157 = πg.NewFunction(πg.NewCode("visit_substitution_definition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1028: """Internal only."""
							πF.SetLineno(1028)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1029: raise nodes.SkipNode
							πF.SetLineno(1029)
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
					if πE = πClass.SetItem(πF, ßvisit_substitution_definition.ToObject(), πTemp157); πE != nil {
						continue
					}
					// line 1028: """Internal only."""
					πF.SetLineno(1028)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp158}, πg.NewStr("Internal only.").ToObject()); πE != nil {
						continue
					}
					if πTemp159, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_substitution_definition); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp159, ß__doc__, πTemp158); πE != nil {
						continue
					}
					// line 1031: def visit_substitution_reference(self, node):
					πF.SetLineno(1031)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp158 = πg.NewFunction(πg.NewCode("visit_substitution_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 1032: self.document.reporter.warning('"substitution_reference" not supported',
							πF.SetLineno(1032)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"substitution_reference\" not supported").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_substitution_reference.ToObject(), πTemp158); πE != nil {
						continue
					}
					// line 1035: def visit_subtitle(self, node):
					πF.SetLineno(1035)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp159 = πg.NewFunction(πg.NewCode("visit_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1036: if isinstance(node.parent, nodes.sidebar):
							πF.SetLineno(1036)
						Label1:
							// line 1037: self.body.append(self.defs['strong'][0])
							πF.SetLineno(1037)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp005 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
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
							goto Label4
							// line 1038: elif isinstance(node.parent, nodes.document):
							πF.SetLineno(1038)
						Label2:
							// line 1039: self.visit_docinfo_item(node, 'subtitle')
							πF.SetLineno(1039)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßsubtitle.ToObject()
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
							goto Label4
							// line 1040: elif isinstance(node.parent, nodes.section):
							πF.SetLineno(1040)
						Label3:
							// line 1041: self.body.append(self.defs['strong'][0])
							πF.SetLineno(1041)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp005 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_subtitle.ToObject(), πTemp159); πE != nil {
						continue
					}
					// line 1043: def depart_subtitle(self, node):
					πF.SetLineno(1043)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp160 = πg.NewFunction(πg.NewCode("depart_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1045: self.body.append(self.defs['strong'][1]+'\n.PP\n')
							πF.SetLineno(1045)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							πTemp005 = ßstrong.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewStr("\n.PP\n").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_subtitle.ToObject(), πTemp160); πE != nil {
						continue
					}
					// line 1047: def visit_system_message(self, node):
					πF.SetLineno(1047)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp161 = πg.NewFunction(πg.NewCode("visit_system_message", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µattr *πg.Object = πg.UnboundLocal
						_ = µattr
						var µbackref_text *πg.Object = πg.UnboundLocal
						_ = µbackref_text
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 1052: attr = {}
							πF.SetLineno(1052)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µattr = πTemp002
							// line 1053: backref_text = ''
							πF.SetLineno(1053)
							µbackref_text = ß.ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1054: if node.hasattr('id'):
							πF.SetLineno(1054)
						Label1:
							// line 1055: attr['name'] = node['id']
							πF.SetLineno(1055)
							πTemp002 = ßid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp006 = ßname.ToObject()
							if πE = πg.SetItem(πF, µattr, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßline.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1056: if node.hasattr('line'):
							πF.SetLineno(1056)
						Label3:
							// line 1057: line = ', line %s' % node['line']
							πF.SetLineno(1057)
							πTemp004 = ßline.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(", line %s").ToObject(), πTemp006); πE != nil {
								continue
							}
							µline = πTemp002
							goto Label5
						Label4:
							// line 1059: line = ''
							πF.SetLineno(1059)
							µline = ß.ToObject()
							goto Label5
						Label5:
							// line 1060: self.body.append('.IP "System Message: %s/%s (%s:%s)"\n'
							πF.SetLineno(1060)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = ßtype.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							πTemp006 = ßlevel.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							πTemp006 = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple4(πTemp007, πTemp008, πTemp009, µline).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(".IP \"System Message: %s/%s (%s:%s)\"\n").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_system_message.ToObject(), πTemp161); πE != nil {
						continue
					}
					// line 1063: def depart_system_message(self, node):
					πF.SetLineno(1063)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp162 = πg.NewFunction(πg.NewCode("depart_system_message", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1064: pass
							πF.SetLineno(1064)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_system_message.ToObject(), πTemp162); πE != nil {
						continue
					}
					// line 1066: def visit_table(self, node):
					πF.SetLineno(1066)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp163 = πg.NewFunction(πg.NewCode("visit_table", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1067: self._active_table = Table()
							πF.SetLineno(1067)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTable); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_active_table, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_table.ToObject(), πTemp163); πE != nil {
						continue
					}
					// line 1069: def depart_table(self, node):
					πF.SetLineno(1069)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp164 = πg.NewFunction(πg.NewCode("depart_table", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1070: self.ensure_eol()
							πF.SetLineno(1070)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßensure_eol, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1071: self.body.extend(self._active_table.as_list())
							πF.SetLineno(1071)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_active_table, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßas_list, nil); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1072: self._active_table = None
							πF.SetLineno(1072)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_active_table, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_table.ToObject(), πTemp164); πE != nil {
						continue
					}
					// line 1074: def visit_target(self, node):
					πF.SetLineno(1074)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp165 = πg.NewFunction(πg.NewCode("visit_target", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1076: raise nodes.SkipNode
							πF.SetLineno(1076)
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
					if πE = πClass.SetItem(πF, ßvisit_target.ToObject(), πTemp165); πE != nil {
						continue
					}
					// line 1078: def visit_tbody(self, node):
					πF.SetLineno(1078)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp166 = πg.NewFunction(πg.NewCode("visit_tbody", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1079: pass
							πF.SetLineno(1079)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_tbody.ToObject(), πTemp166); πE != nil {
						continue
					}
					// line 1081: def depart_tbody(self, node):
					πF.SetLineno(1081)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp167 = πg.NewFunction(πg.NewCode("depart_tbody", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1082: pass
							πF.SetLineno(1082)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_tbody.ToObject(), πTemp167); πE != nil {
						continue
					}
					// line 1084: def visit_term(self, node):
					πF.SetLineno(1084)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp168 = πg.NewFunction(πg.NewCode("visit_term", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1085: self.body.append(self.defs['term'][0])
							πF.SetLineno(1085)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßterm.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_term.ToObject(), πTemp168); πE != nil {
						continue
					}
					// line 1087: def depart_term(self, node):
					πF.SetLineno(1087)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp169 = πg.NewFunction(πg.NewCode("depart_term", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1088: self.body.append(self.defs['term'][1])
							πF.SetLineno(1088)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßterm.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_term.ToObject(), πTemp169); πE != nil {
						continue
					}
					// line 1090: def visit_tgroup(self, node):
					πF.SetLineno(1090)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp170 = πg.NewFunction(πg.NewCode("visit_tgroup", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1091: pass
							πF.SetLineno(1091)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_tgroup.ToObject(), πTemp170); πE != nil {
						continue
					}
					// line 1093: def depart_tgroup(self, node):
					πF.SetLineno(1093)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp171 = πg.NewFunction(πg.NewCode("depart_tgroup", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1094: pass
							πF.SetLineno(1094)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_tgroup.ToObject(), πTemp171); πE != nil {
						continue
					}
					// line 1096: def visit_thead(self, node):
					πF.SetLineno(1096)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp172 = πg.NewFunction(πg.NewCode("visit_thead", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1098: pass
							πF.SetLineno(1098)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_thead.ToObject(), πTemp172); πE != nil {
						continue
					}
					// line 1100: def depart_thead(self, node):
					πF.SetLineno(1100)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp173 = πg.NewFunction(πg.NewCode("depart_thead", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1102: pass
							πF.SetLineno(1102)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_thead.ToObject(), πTemp173); πE != nil {
						continue
					}
					// line 1104: def visit_tip(self, node):
					πF.SetLineno(1104)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp174 = πg.NewFunction(πg.NewCode("visit_tip", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1105: self.visit_admonition(node, 'tip')
							πF.SetLineno(1105)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßtip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_tip.ToObject(), πTemp174); πE != nil {
						continue
					}
					// line 1107: depart_tip = depart_admonition
					πF.SetLineno(1107)
					if πTemp175, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_tip.ToObject(), πTemp175); πE != nil {
						continue
					}
					// line 1109: def visit_title(self, node):
					πF.SetLineno(1109)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp175 = πg.NewFunction(πg.NewCode("visit_title", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 *πg.Object
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtopic, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadmonition, nil); πE != nil {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 1110: if isinstance(node.parent, nodes.topic):
							πF.SetLineno(1110)
						Label1:
							// line 1111: self.body.append(self.defs['topic-title'][0])
							πF.SetLineno(1111)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp005 = πg.NewStr("topic-title").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
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
							goto Label7
							// line 1112: elif isinstance(node.parent, nodes.sidebar):
							πF.SetLineno(1112)
						Label2:
							// line 1113: self.body.append(self.defs['sidebar-title'][0])
							πF.SetLineno(1113)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp005 = πg.NewStr("sidebar-title").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
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
							goto Label7
							// line 1114: elif isinstance(node.parent, nodes.admonition):
							πF.SetLineno(1114)
						Label3:
							// line 1115: self.body.append('.IP "')
							πF.SetLineno(1115)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".IP \"").ToObject()
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
							goto Label7
							// line 1116: elif self.section_level == 0:
							πF.SetLineno(1116)
						Label4:
							// line 1117: self._docinfo['title'] = node.astext()
							πF.SetLineno(1117)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							πTemp006 = ßtitle.ToObject()
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							// line 1119: self._docinfo['title_upper'] = node.astext().upper()
							πF.SetLineno(1119)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßupper, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_docinfo, nil); πE != nil {
								continue
							}
							πTemp006 = ßtitle_upper.ToObject()
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1120: raise nodes.SkipNode
							πF.SetLineno(1120)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label7
							// line 1121: elif self.section_level == 1:
							πF.SetLineno(1121)
						Label5:
							// line 1122: self.body.append('.SH %s\n' % self.deunicode(node.astext().upper()))
							πF.SetLineno(1122)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßupper, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdeunicode, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr(".SH %s\n").ToObject(), πTemp005); πE != nil {
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1123: raise nodes.SkipNode
							πF.SetLineno(1123)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label7
						Label6:
							// line 1125: self.body.append('.SS ')
							πF.SetLineno(1125)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(".SS ").ToObject()
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
							goto Label7
						Label7:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_title.ToObject(), πTemp175); πE != nil {
						continue
					}
					// line 1127: def depart_title(self, node):
					πF.SetLineno(1127)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp176 = πg.NewFunction(πg.NewCode("depart_title", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadmonition, nil); πE != nil {
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
							// line 1128: if isinstance(node.parent, nodes.admonition):
							πF.SetLineno(1128)
						Label1:
							// line 1129: self.body.append('"')
							πF.SetLineno(1129)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\"").ToObject()
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
							// line 1130: self.body.append('\n')
							πF.SetLineno(1130)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_title.ToObject(), πTemp176); πE != nil {
						continue
					}
					// line 1132: def visit_title_reference(self, node):
					πF.SetLineno(1132)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp177 = πg.NewFunction(πg.NewCode("visit_title_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1133: """inline citation reference"""
							πF.SetLineno(1133)
							// line 1134: self.body.append(self.defs['title_reference'][0])
							πF.SetLineno(1134)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = ßtitle_reference.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_title_reference.ToObject(), πTemp177); πE != nil {
						continue
					}
					// line 1133: """inline citation reference"""
					πF.SetLineno(1133)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp178}, πg.NewStr("inline citation reference").ToObject()); πE != nil {
						continue
					}
					if πTemp179, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_title_reference); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp179, ß__doc__, πTemp178); πE != nil {
						continue
					}
					// line 1136: def depart_title_reference(self, node):
					πF.SetLineno(1136)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp178 = πg.NewFunction(πg.NewCode("depart_title_reference", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1137: self.body.append(self.defs['title_reference'][1])
							πF.SetLineno(1137)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = ßtitle_reference.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdefs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßdepart_title_reference.ToObject(), πTemp178); πE != nil {
						continue
					}
					// line 1139: def visit_topic(self, node):
					πF.SetLineno(1139)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp179 = πg.NewFunction(πg.NewCode("visit_topic", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1140: pass
							πF.SetLineno(1140)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_topic.ToObject(), πTemp179); πE != nil {
						continue
					}
					// line 1142: def depart_topic(self, node):
					πF.SetLineno(1142)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp180 = πg.NewFunction(πg.NewCode("depart_topic", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1143: pass
							πF.SetLineno(1143)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_topic.ToObject(), πTemp180); πE != nil {
						continue
					}
					// line 1145: def visit_sidebar(self, node):
					πF.SetLineno(1145)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp181 = πg.NewFunction(πg.NewCode("visit_sidebar", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1146: pass
							πF.SetLineno(1146)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_sidebar.ToObject(), πTemp181); πE != nil {
						continue
					}
					// line 1148: def depart_sidebar(self, node):
					πF.SetLineno(1148)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp182 = πg.NewFunction(πg.NewCode("depart_sidebar", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1149: pass
							πF.SetLineno(1149)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_sidebar.ToObject(), πTemp182); πE != nil {
						continue
					}
					// line 1151: def visit_rubric(self, node):
					πF.SetLineno(1151)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp183 = πg.NewFunction(πg.NewCode("visit_rubric", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1152: pass
							πF.SetLineno(1152)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_rubric.ToObject(), πTemp183); πE != nil {
						continue
					}
					// line 1154: def depart_rubric(self, node):
					πF.SetLineno(1154)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp184 = πg.NewFunction(πg.NewCode("depart_rubric", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1155: pass
							πF.SetLineno(1155)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_rubric.ToObject(), πTemp184); πE != nil {
						continue
					}
					// line 1157: def visit_transition(self, node):
					πF.SetLineno(1157)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp185 = πg.NewFunction(πg.NewCode("visit_transition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1161: self.body.append('\n.sp\n.ce\n----\n')
							πF.SetLineno(1161)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n.sp\n.ce\n----\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_transition.ToObject(), πTemp185); πE != nil {
						continue
					}
					// line 1163: def depart_transition(self, node):
					πF.SetLineno(1163)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp186 = πg.NewFunction(πg.NewCode("depart_transition", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1164: self.body.append('\n.ce 0\n.sp\n')
							πF.SetLineno(1164)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n.ce 0\n.sp\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_transition.ToObject(), πTemp186); πE != nil {
						continue
					}
					// line 1166: def visit_version(self, node):
					πF.SetLineno(1166)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp187 = πg.NewFunction(πg.NewCode("visit_version", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1167: self.visit_docinfo_item(node, 'version')
							πF.SetLineno(1167)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßversion.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_version.ToObject(), πTemp187); πE != nil {
						continue
					}
					// line 1169: def visit_warning(self, node):
					πF.SetLineno(1169)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp188 = πg.NewFunction(πg.NewCode("visit_warning", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1170: self.visit_admonition(node, 'warning')
							πF.SetLineno(1170)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßwarning.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_admonition, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_warning.ToObject(), πTemp188); πE != nil {
						continue
					}
					// line 1172: depart_warning = depart_admonition
					πF.SetLineno(1172)
					if πTemp189, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_admonition); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdepart_warning.ToObject(), πTemp189); πE != nil {
						continue
					}
					// line 1174: def unimplemented_visit(self, node):
					πF.SetLineno(1174)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp189 = πg.NewFunction(πg.NewCode("unimplemented_visit", "/usr/lib/python2.7/site-packages/docutils/writers/manpage.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("visiting unimplemented node type: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1175: raise NotImplementedError('visiting unimplemented node type: %s'
							πF.SetLineno(1175)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßunimplemented_visit.ToObject(), πTemp189); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Translator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTranslator.ToObject(), πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("manpage", Code)
}
