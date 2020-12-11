package _html_base

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/Image"
	_ "github.com/pygolin/stdlib/pkg/PIL"
	_ "github.com/pygolin/stdlib/pkg/PIL/Image"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/urllib"
	_ "github.com/pygolin/stdlib/pkg/urllib/request"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAdmonition := πg.InternStr("Admonition")
		ßAdmonitions := πg.InternStr("Admonitions")
		ßCLASS := πg.InternStr("CLASS")
		ßDocumentParameters := πg.InternStr("DocumentParameters")
		ßFalse := πg.InternStr("False")
		ßFileInput := πg.InternStr("FileInput")
		ßGenericNodeVisitor := πg.InternStr("GenericNodeVisitor")
		ßHTMLTranslator := πg.InternStr("HTMLTranslator")
		ßIOError := πg.InternStr("IOError")
		ßImage := πg.InternStr("Image")
		ßImportError := πg.InternStr("ImportError")
		ßIndexError := πg.InternStr("IndexError")
		ßInvisible := πg.InternStr("Invisible")
		ßNodeFound := πg.InternStr("NodeFound")
		ßNodeVisitor := πg.InternStr("NodeVisitor")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOSError := πg.InternStr("OSError")
		ßPIL := πg.InternStr("PIL")
		ßSafeString := πg.InternStr("SafeString")
		ßSimpleListChecker := πg.InternStr("SimpleListChecker")
		ßSkipNode := πg.InternStr("SkipNode")
		ßSyntaxError := πg.InternStr("SyntaxError")
		ßTextElement := πg.InternStr("TextElement")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßU := πg.InternStr("U")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ßWriter := πg.InternStr("Writer")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß_destination := πg.InternStr("_destination")
		ßa := πg.InternStr("a")
		ßabbr := πg.InternStr("abbr")
		ßacronym := πg.InternStr("acronym")
		ßadd := πg.InternStr("add")
		ßadd_meta := πg.InternStr("add_meta")
		ßaddress := πg.InternStr("address")
		ßadmonition := πg.InternStr("admonition")
		ßalign := πg.InternStr("align")
		ßalt := πg.InternStr("alt")
		ßappend := πg.InternStr("append")
		ßapply_template := πg.InternStr("apply_template")
		ßargs := πg.InternStr("args")
		ßassemble_parts := πg.InternStr("assemble_parts")
		ßastext := πg.InternStr("astext")
		ßattributes := πg.InternStr("attributes")
		ßattribution := πg.InternStr("attribution")
		ßattribution_formats := πg.InternStr("attribution_formats")
		ßattval := πg.InternStr("attval")
		ßauthor := πg.InternStr("author")
		ßauthor_in_authors := πg.InternStr("author_in_authors")
		ßauthors := πg.InternStr("authors")
		ßbackrefs := πg.InternStr("backrefs")
		ßbasename := πg.InternStr("basename")
		ßblahtexml := πg.InternStr("blahtexml")
		ßblockquote := πg.InternStr("blockquote")
		ßbody := πg.InternStr("body")
		ßbody_pre_docinfo := πg.InternStr("body_pre_docinfo")
		ßbody_prefix := πg.InternStr("body_prefix")
		ßbody_suffix := πg.InternStr("body_suffix")
		ßbool := πg.InternStr("bool")
		ßbrackets := πg.InternStr("brackets")
		ßbullet_list := πg.InternStr("bullet_list")
		ßcaption := πg.InternStr("caption")
		ßcheck_simple_list := πg.InternStr("check_simple_list")
		ßchildren := πg.InternStr("children")
		ßcitation := πg.InternStr("citation")
		ßcite := πg.InternStr("cite")
		ßclass := πg.InternStr("class")
		ßclasses := πg.InternStr("classes")
		ßclassifier := πg.InternStr("classifier")
		ßcloak_email := πg.InternStr("cloak_email")
		ßcloak_email_addresses := πg.InternStr("cloak_email_addresses")
		ßcloak_mailto := πg.InternStr("cloak_mailto")
		ßclose := πg.InternStr("close")
		ßcode := πg.InternStr("code")
		ßcol := πg.InternStr("col")
		ßcolgroup := πg.InternStr("colgroup")
		ßcolspan := πg.InternStr("colspan")
		ßcolspec := πg.InternStr("colspec")
		ßcolspecs := πg.InternStr("colspecs")
		ßcolumn := πg.InternStr("column")
		ßcolwidth := πg.InternStr("colwidth")
		ßcompact := πg.InternStr("compact")
		ßcompact_field_list := πg.InternStr("compact_field_list")
		ßcompact_field_lists := πg.InternStr("compact_field_lists")
		ßcompact_lists := πg.InternStr("compact_lists")
		ßcompact_p := πg.InternStr("compact_p")
		ßcompact_simple := πg.InternStr("compact_simple")
		ßcompile := πg.InternStr("compile")
		ßcompound := πg.InternStr("compound")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßcontact := πg.InternStr("contact")
		ßcontent_type := πg.InternStr("content_type")
		ßcontent_type_mathml := πg.InternStr("content_type_mathml")
		ßcontents := πg.InternStr("contents")
		ßcontext := πg.InternStr("context")
		ßcopyright := πg.InternStr("copyright")
		ßdash := πg.InternStr("dash")
		ßdata := πg.InternStr("data")
		ßdate := πg.InternStr("date")
		ßdd := πg.InternStr("dd")
		ßdefault_template := πg.InternStr("default_template")
		ßdefault_visit := πg.InternStr("default_visit")
		ßdefinition_list := πg.InternStr("definition_list")
		ßdelimiter := πg.InternStr("delimiter")
		ßdepart_Text := πg.InternStr("depart_Text")
		ßdepart_abbreviation := πg.InternStr("depart_abbreviation")
		ßdepart_acronym := πg.InternStr("depart_acronym")
		ßdepart_address := πg.InternStr("depart_address")
		ßdepart_admonition := πg.InternStr("depart_admonition")
		ßdepart_attribution := πg.InternStr("depart_attribution")
		ßdepart_author := πg.InternStr("depart_author")
		ßdepart_authors := πg.InternStr("depart_authors")
		ßdepart_block_quote := πg.InternStr("depart_block_quote")
		ßdepart_bullet_list := πg.InternStr("depart_bullet_list")
		ßdepart_caption := πg.InternStr("depart_caption")
		ßdepart_citation := πg.InternStr("depart_citation")
		ßdepart_citation_reference := πg.InternStr("depart_citation_reference")
		ßdepart_classifier := πg.InternStr("depart_classifier")
		ßdepart_colspec := πg.InternStr("depart_colspec")
		ßdepart_compound := πg.InternStr("depart_compound")
		ßdepart_contact := πg.InternStr("depart_contact")
		ßdepart_container := πg.InternStr("depart_container")
		ßdepart_copyright := πg.InternStr("depart_copyright")
		ßdepart_date := πg.InternStr("depart_date")
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
		ßdepart_image := πg.InternStr("depart_image")
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
		ßdepart_target := πg.InternStr("depart_target")
		ßdepart_tbody := πg.InternStr("depart_tbody")
		ßdepart_term := πg.InternStr("depart_term")
		ßdepart_tgroup := πg.InternStr("depart_tgroup")
		ßdepart_thead := πg.InternStr("depart_thead")
		ßdepart_title := πg.InternStr("depart_title")
		ßdepart_title_reference := πg.InternStr("depart_title_reference")
		ßdepart_topic := πg.InternStr("depart_topic")
		ßdepart_transition := πg.InternStr("depart_transition")
		ßdepart_version := πg.InternStr("depart_version")
		ßdisplaymode := πg.InternStr("displaymode")
		ßdiv := πg.InternStr("div")
		ßdl := πg.InternStr("dl")
		ßdocinfo := πg.InternStr("docinfo")
		ßdoctype := πg.InternStr("doctype")
		ßdoctype_mathml := πg.InternStr("doctype_mathml")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßdt := πg.InternStr("dt")
		ßem := πg.InternStr("em")
		ßembed_stylesheet := πg.InternStr("embed_stylesheet")
		ßembedded_stylesheet := πg.InternStr("embedded_stylesheet")
		ßemptytag := πg.InternStr("emptytag")
		ßencode := πg.InternStr("encode")
		ßencoding := πg.InternStr("encoding")
		ßentry := πg.InternStr("entry")
		ßenumerate := πg.InternStr("enumerate")
		ßenumerated_list := πg.InternStr("enumerated_list")
		ßenumtype := πg.InternStr("enumtype")
		ßerror := πg.InternStr("error")
		ßextend := πg.InternStr("extend")
		ßfield_list := πg.InternStr("field_list")
		ßfigure := πg.InternStr("figure")
		ßfile_insertion_enabled := πg.InternStr("file_insertion_enabled")
		ßfind_file_in_dirs := πg.InternStr("find_file_in_dirs")
		ßfindall := πg.InternStr("findall")
		ßfloat := πg.InternStr("float")
		ßfooter := πg.InternStr("footer")
		ßfootnote := πg.InternStr("footnote")
		ßfootnote_backlinks := πg.InternStr("footnote_backlinks")
		ßfootnote_references := πg.InternStr("footnote_references")
		ßformat := πg.InternStr("format")
		ßformula := πg.InternStr("formula")
		ßfragment := πg.InternStr("fragment")
		ßgenerator := πg.InternStr("generator")
		ßget := πg.InternStr("get")
		ßget_language := πg.InternStr("get_language")
		ßget_stylesheet_list := πg.InternStr("get_stylesheet_list")
		ßget_transforms := πg.InternStr("get_transforms")
		ßgetattr := πg.InternStr("getattr")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßgroup := πg.InternStr("group")
		ßh1 := πg.InternStr("h1")
		ßhasattr := πg.InternStr("hasattr")
		ßhead := πg.InternStr("head")
		ßhead_prefix := πg.InternStr("head_prefix")
		ßhead_prefix_template := πg.InternStr("head_prefix_template")
		ßheader := πg.InternStr("header")
		ßheight := πg.InternStr("height")
		ßhr := πg.InternStr("hr")
		ßhref := πg.InternStr("href")
		ßhtml := πg.InternStr("html")
		ßhtml_body := πg.InternStr("html_body")
		ßhtml_head := πg.InternStr("html_head")
		ßhtml_prolog := πg.InternStr("html_prolog")
		ßhtml_subtitle := πg.InternStr("html_subtitle")
		ßhtml_title := πg.InternStr("html_title")
		ßid := πg.InternStr("id")
		ßids := πg.InternStr("ids")
		ßignore_node := πg.InternStr("ignore_node")
		ßimage := πg.InternStr("image")
		ßimg := πg.InternStr("img")
		ßin_docinfo := πg.InternStr("in_docinfo")
		ßin_document_title := πg.InternStr("in_document_title")
		ßin_footnote_list := πg.InternStr("in_footnote_list")
		ßin_mailto := πg.InternStr("in_mailto")
		ßin_sidebar := πg.InternStr("in_sidebar")
		ßin_word_wrap_point := πg.InternStr("in_word_wrap_point")
		ßinitial_header_level := πg.InternStr("initial_header_level")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßinterpolation_dict := πg.InternStr("interpolation_dict")
		ßio := πg.InternStr("io")
		ßis_compactable := πg.InternStr("is_compactable")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßlabel := πg.InternStr("label")
		ßlabels := πg.InternStr("labels")
		ßlang := πg.InternStr("lang")
		ßlang_attribute := πg.InternStr("lang_attribute")
		ßlanguage := πg.InternStr("language")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßlatex := πg.InternStr("latex")
		ßlatex2mathml := πg.InternStr("latex2mathml")
		ßlatexml := πg.InternStr("latexml")
		ßlegend := πg.InternStr("legend")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßli := πg.InternStr("li")
		ßline := πg.InternStr("line")
		ßlist := πg.InternStr("list")
		ßlist_item := πg.InternStr("list_item")
		ßliteral_block := πg.InternStr("literal_block")
		ßlower := πg.InternStr("lower")
		ßmatch := πg.InternStr("match")
		ßmath := πg.InternStr("math")
		ßmath2html := πg.InternStr("math2html")
		ßmath_header := πg.InternStr("math_header")
		ßmath_output := πg.InternStr("math_output")
		ßmath_output_options := πg.InternStr("math_output_options")
		ßmath_tags := πg.InternStr("math_tags")
		ßmathjax := πg.InternStr("mathjax")
		ßmathjax_script := πg.InternStr("mathjax_script")
		ßmathjax_url := πg.InternStr("mathjax_url")
		ßmathml := πg.InternStr("mathml")
		ßmeta := πg.InternStr("meta")
		ßmorecols := πg.InternStr("morecols")
		ßmorerows := πg.InternStr("morerows")
		ßnameids := πg.InternStr("nameids")
		ßnames := πg.InternStr("names")
		ßnext_node := πg.InternStr("next_node")
		ßnodes := πg.InternStr("nodes")
		ßnon_default_attributes := πg.InternStr("non_default_attributes")
		ßnone := πg.InternStr("none")
		ßobject := πg.InternStr("object")
		ßobject_image_types := πg.InternStr("object_image_types")
		ßol := πg.InternStr("ol")
		ßopen := πg.InternStr("open")
		ßoption := πg.InternStr("option")
		ßoption_list := πg.InternStr("option_list")
		ßord := πg.InternStr("ord")
		ßorganization := πg.InternStr("organization")
		ßos := πg.InternStr("os")
		ßoutput := πg.InternStr("output")
		ßoutput_encoding := πg.InternStr("output_encoding")
		ßoutput_encoding_error_handler := πg.InternStr("output_encoding_error_handler")
		ßp := πg.InternStr("p")
		ßparagraph := πg.InternStr("paragraph")
		ßparens := πg.InternStr("parens")
		ßparent := πg.InternStr("parent")
		ßparentheses := πg.InternStr("parentheses")
		ßparts := πg.InternStr("parts")
		ßpass_node := πg.InternStr("pass_node")
		ßpath := πg.InternStr("path")
		ßpick_math_environment := πg.InternStr("pick_math_environment")
		ßpop := πg.InternStr("pop")
		ßpre := πg.InternStr("pre")
		ßproblematic := πg.InternStr("problematic")
		ßpx := πg.InternStr("px")
		ßrb := πg.InternStr("rb")
		ßre := πg.InternStr("re")
		ßread := πg.InternStr("read")
		ßrecord_dependencies := πg.InternStr("record_dependencies")
		ßreference := πg.InternStr("reference")
		ßrefid := πg.InternStr("refid")
		ßrefname := πg.InternStr("refname")
		ßrefuri := πg.InternStr("refuri")
		ßrelative_path := πg.InternStr("relative_path")
		ßreplace := πg.InternStr("replace")
		ßreporter := πg.InternStr("reporter")
		ßrevision := πg.InternStr("revision")
		ßrowspan := πg.InternStr("rowspan")
		ßrstrip := πg.InternStr("rstrip")
		ßrubric := πg.InternStr("rubric")
		ßscale := πg.InternStr("scale")
		ßsearch := πg.InternStr("search")
		ßsection := πg.InternStr("section")
		ßsection_level := πg.InternStr("section_level")
		ßsectnum := πg.InternStr("sectnum")
		ßset_class_on_child := πg.InternStr("set_class_on_child")
		ßsetattr := πg.InternStr("setattr")
		ßsetdefault := πg.InternStr("setdefault")
		ßsettings := πg.InternStr("settings")
		ßsettings_defaults := πg.InternStr("settings_defaults")
		ßsidebar := πg.InternStr("sidebar")
		ßsimple := πg.InternStr("simple")
		ßsize := πg.InternStr("size")
		ßsorted := πg.InternStr("sorted")
		ßsource := πg.InternStr("source")
		ßspan := πg.InternStr("span")
		ßspecial_characters := πg.InternStr("special_characters")
		ßsplit := πg.InternStr("split")
		ßsplitext := πg.InternStr("splitext")
		ßsrc := πg.InternStr("src")
		ßstart := πg.InternStr("start")
		ßstartswith := πg.InternStr("startswith")
		ßstarttag := πg.InternStr("starttag")
		ßstatus := πg.InternStr("status")
		ßstr := πg.InternStr("str")
		ßstrerror := πg.InternStr("strerror")
		ßstrip := πg.InternStr("strip")
		ßstrong := πg.InternStr("strong")
		ßstub := πg.InternStr("stub")
		ßstubs := πg.InternStr("stubs")
		ßstyle := πg.InternStr("style")
		ßstylesheet := πg.InternStr("stylesheet")
		ßstylesheet_call := πg.InternStr("stylesheet_call")
		ßstylesheet_dirs := πg.InternStr("stylesheet_dirs")
		ßstylesheet_link := πg.InternStr("stylesheet_link")
		ßstylesheet_path := πg.InternStr("stylesheet_path")
		ßsub := πg.InternStr("sub")
		ßsubtitle := πg.InternStr("subtitle")
		ßsum := πg.InternStr("sum")
		ßsup := πg.InternStr("sup")
		ßsupported := πg.InternStr("supported")
		ßsys := πg.InternStr("sys")
		ßtable := πg.InternStr("table")
		ßtable_style := πg.InternStr("table_style")
		ßtarget := πg.InternStr("target")
		ßtbody := πg.InternStr("tbody")
		ßtd := πg.InternStr("td")
		ßtemplate := πg.InternStr("template")
		ßtex2mathml := πg.InternStr("tex2mathml")
		ßtex2mathml_extern := πg.InternStr("tex2mathml_extern")
		ßth := πg.InternStr("th")
		ßthead := πg.InternStr("thead")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßtopic_classes := πg.InternStr("topic_classes")
		ßtr := πg.InternStr("tr")
		ßtranslate := πg.InternStr("translate")
		ßtranslator_class := πg.InternStr("translator_class")
		ßtt := πg.InternStr("tt")
		ßttm := πg.InternStr("ttm")
		ßtype := πg.InternStr("type")
		ßul := πg.InternStr("ul")
		ßuni2tex_table := πg.InternStr("uni2tex_table")
		ßunichar2tex := πg.InternStr("unichar2tex")
		ßunicode := πg.InternStr("unicode")
		ßunimplemented_visit := πg.InternStr("unimplemented_visit")
		ßuri := πg.InternStr("uri")
		ßurl2pathname := πg.InternStr("url2pathname")
		ßutils := πg.InternStr("utils")
		ßvar := πg.InternStr("var")
		ßversion := πg.InternStr("version")
		ßversion_info := πg.InternStr("version_info")
		ßvisit_Text := πg.InternStr("visit_Text")
		ßvisit_abbreviation := πg.InternStr("visit_abbreviation")
		ßvisit_acronym := πg.InternStr("visit_acronym")
		ßvisit_address := πg.InternStr("visit_address")
		ßvisit_admonition := πg.InternStr("visit_admonition")
		ßvisit_attribution := πg.InternStr("visit_attribution")
		ßvisit_author := πg.InternStr("visit_author")
		ßvisit_authors := πg.InternStr("visit_authors")
		ßvisit_block_quote := πg.InternStr("visit_block_quote")
		ßvisit_bullet_list := πg.InternStr("visit_bullet_list")
		ßvisit_caption := πg.InternStr("visit_caption")
		ßvisit_citation := πg.InternStr("visit_citation")
		ßvisit_citation_reference := πg.InternStr("visit_citation_reference")
		ßvisit_classifier := πg.InternStr("visit_classifier")
		ßvisit_colspec := πg.InternStr("visit_colspec")
		ßvisit_comment := πg.InternStr("visit_comment")
		ßvisit_compound := πg.InternStr("visit_compound")
		ßvisit_contact := πg.InternStr("visit_contact")
		ßvisit_container := πg.InternStr("visit_container")
		ßvisit_copyright := πg.InternStr("visit_copyright")
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
		ßvisit_image := πg.InternStr("visit_image")
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
		ßvisit_option := πg.InternStr("visit_option")
		ßvisit_option_argument := πg.InternStr("visit_option_argument")
		ßvisit_option_group := πg.InternStr("visit_option_group")
		ßvisit_option_list := πg.InternStr("visit_option_list")
		ßvisit_option_list_item := πg.InternStr("visit_option_list_item")
		ßvisit_option_string := πg.InternStr("visit_option_string")
		ßvisit_organization := πg.InternStr("visit_organization")
		ßvisit_paragraph := πg.InternStr("visit_paragraph")
		ßvisit_pending := πg.InternStr("visit_pending")
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
		ßvisit_title := πg.InternStr("visit_title")
		ßvisit_title_reference := πg.InternStr("visit_title_reference")
		ßvisit_topic := πg.InternStr("visit_topic")
		ßvisit_transition := πg.InternStr("visit_transition")
		ßvisit_version := πg.InternStr("visit_version")
		ßvisitor := πg.InternStr("visitor")
		ßvisitor_attributes := πg.InternStr("visitor_attributes")
		ßwalk := πg.InternStr("walk")
		ßwalkabout := πg.InternStr("walkabout")
		ßwarning := πg.InternStr("warning")
		ßwidth := πg.InternStr("width")
		ßwords_and_spaces := πg.InternStr("words_and_spaces")
		ßwriter_aux := πg.InternStr("writer_aux")
		ßwriters := πg.InternStr("writers")
		ßxhtml := πg.InternStr("xhtml")
		ßxml_declaration := πg.InternStr("xml_declaration")
		ßxmlcharrefreplace := πg.InternStr("xmlcharrefreplace")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.BaseException
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
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
			// line 18: """common definitions for Docutils HTML writers"""
			πF.SetLineno(18)
			// line 18: """common definitions for Docutils HTML writers"""
			πF.SetLineno(18)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("common definitions for Docutils HTML writers").ToObject()); πE != nil {
				continue
			}
			// line 20: import sys
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: import os.path
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: import re
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 24: try: # check for the Python Imaging Library
			πF.SetLineno(24)
			πF.PushCheckpoint(2)
			// line 25: import PIL.Image
			πF.SetLineno(25)
			if πTemp002, πE = πg.ImportModule(πF, "PIL.Image"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßPIL.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
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
			// line 26: except ImportError:
			πF.SetLineno(26)
		Label3:
			// line 27: try:  # sometimes PIL modules are put in PYTHONPATH's root
			πF.SetLineno(27)
			πF.PushCheckpoint(5)
			// line 28: import Image
			πF.SetLineno(28)
			if πTemp002, πE = πg.ImportModule(πF, "Image"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßImage.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: class PIL(object): pass  # dummy wrapper
			πF.SetLineno(29)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PIL", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 29: class PIL(object): pass  # dummy wrapper
					πF.SetLineno(29)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("PIL").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPIL.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 30: PIL.Image = Image
			πF.SetLineno(30)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImage); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp001); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßPIL); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ßImage, πTemp007); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp009, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label6
			}
			πE = πF.Raise(πTemp009.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 31: except ImportError:
			πF.SetLineno(31)
		Label6:
			// line 32: PIL = None
			πF.SetLineno(32)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPIL.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 34: import docutils
			πF.SetLineno(34)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: from docutils import nodes, utils, writers, languages, io
			πF.SetLineno(35)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 36: from docutils.utils.error_reporting import SafeString
			πF.SetLineno(36)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp007, πE = πg.GetAttrImport(πF, πTemp001, ßSafeString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSafeString.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 37: from docutils.transforms import writer_aux
			πF.SetLineno(37)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.writer_aux"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßwriter_aux.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 38: from docutils.utils.math import (unichar2tex, pick_math_environment,
			πF.SetLineno(38)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.math.unichar2tex"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßunichar2tex.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.math"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp007, πE = πg.GetAttrImport(πF, πTemp001, ßpick_math_environment); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpick_math_environment.ToObject(), πTemp007); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.math.math2html"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßmath2html.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.math.latex2mathml"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßlatex2mathml.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.math.tex2mathml_extern"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßtex2mathml_extern.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp007 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp008, πTemp007); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label7
			}
			goto Label8
			// line 41: if sys.version_info >= (3, 0):
			πF.SetLineno(41)
		Label7:
			// line 42: from urllib.request import url2pathname
			πF.SetLineno(42)
			if πTemp002, πE = πg.ImportModule(πF, "urllib.request"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp007, πE = πg.GetAttrImport(πF, πTemp001, ßurl2pathname); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßurl2pathname.ToObject(), πTemp007); πE != nil {
				continue
			}
			goto Label9
		Label8:
			// line 44: from urllib import url2pathname
			πF.SetLineno(44)
			if πTemp002, πE = πg.ImportModule(πF, "urllib"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp007, πE = πg.GetAttrImport(πF, πTemp001, ßurl2pathname); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßurl2pathname.ToObject(), πTemp007); πE != nil {
				continue
			}
			goto Label9
		Label9:
			if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp007 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp008, πTemp007); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label10
			}
			goto Label11
			// line 46: if sys.version_info >= (3, 0):
			πF.SetLineno(46)
		Label10:
			// line 47: unicode = str  # noqa
			πF.SetLineno(47)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label11
		Label11:
			// line 50: class Writer(writers.Writer):
			πF.SetLineno(50)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
				_ = πTemp002
				var πTemp003 []*πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 52: supported = ('html', 'xhtml') # update in subclass
					πF.SetLineno(52)
					πTemp001 = πg.NewTuple2(ßhtml.ToObject(), ßxhtml.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 53: """Formats this writer supports."""
					πF.SetLineno(53)
					// line 53: """Formats this writer supports."""
					πF.SetLineno(53)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Formats this writer supports.").ToObject()); πE != nil {
						continue
					}
					// line 57: default_template = 'template.txt'
					πF.SetLineno(57)
					if πE = πClass.SetItem(πF, ßdefault_template.ToObject(), πg.NewStr("template.txt").ToObject()); πE != nil {
						continue
					}
					// line 61: settings_defaults = {'output_encoding_error_handler': 'xmlcharrefreplace'}
					πF.SetLineno(61)
					πTemp002 = πg.NewDict()
					if πE = πTemp002.SetItem(πF, ßoutput_encoding_error_handler.ToObject(), ßxmlcharrefreplace.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßsettings_defaults.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 64: config_section_dependencies = ('writers', 'html writers')
					πF.SetLineno(64)
					πTemp001 = πg.NewTuple2(ßwriters.ToObject(), πg.NewStr("html writers").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 66: visitor_attributes = (
					πF.SetLineno(66)
					πTemp003 = make([]*πg.Object, 19)
					πTemp003[0] = ßhead_prefix.ToObject()
					πTemp003[1] = ßhead.ToObject()
					πTemp003[2] = ßstylesheet.ToObject()
					πTemp003[3] = ßbody_prefix.ToObject()
					πTemp003[4] = ßbody_pre_docinfo.ToObject()
					πTemp003[5] = ßdocinfo.ToObject()
					πTemp003[6] = ßbody.ToObject()
					πTemp003[7] = ßbody_suffix.ToObject()
					πTemp003[8] = ßtitle.ToObject()
					πTemp003[9] = ßsubtitle.ToObject()
					πTemp003[10] = ßheader.ToObject()
					πTemp003[11] = ßfooter.ToObject()
					πTemp003[12] = ßmeta.ToObject()
					πTemp003[13] = ßfragment.ToObject()
					πTemp003[14] = ßhtml_prolog.ToObject()
					πTemp003[15] = ßhtml_head.ToObject()
					πTemp003[16] = ßhtml_title.ToObject()
					πTemp003[17] = ßhtml_subtitle.ToObject()
					πTemp003[18] = ßhtml_body.ToObject()
					πTemp001 = πg.NewTuple(πTemp003...).ToObject()
					if πE = πClass.SetItem(πF, ßvisitor_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 73: def get_transforms(self):
					πF.SetLineno(73)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 74: return writers.Writer.get_transforms(self) + [writer_aux.Admonitions]
							πF.SetLineno(74)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßWriter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßget_transforms, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßwriter_aux); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßAdmonitions, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_transforms.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 76: def translate(self):
					πF.SetLineno(76)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("translate", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
						var µattr *πg.Object = πg.UnboundLocal
						_ = µattr
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 77: self.visitor = visitor = self.translator_class(self.document)
							πF.SetLineno(77)
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
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßvisitor, πTemp002); πE != nil {
								continue
							}
							µvisitor = πTemp003
							// line 78: self.document.walkabout(visitor)
							πF.SetLineno(78)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisitor_attributes, nil); πE != nil {
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
								µattr = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 80: setattr(self, attr, getattr(visitor, attr))
							πF.SetLineno(80)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp001[1] = µattr
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp006[0] = µvisitor
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp006[1] = µattr
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[2] = πTemp007
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 81: self.output = self.apply_template()
							πF.SetLineno(81)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßapply_template, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtranslate.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 83: def apply_template(self):
					πF.SetLineno(83)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("apply_template", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtemplate_file *πg.Object = πg.UnboundLocal
						_ = µtemplate_file
						var µtemplate *πg.Object = πg.UnboundLocal
						_ = µtemplate
						var µsubs *πg.Object = πg.UnboundLocal
						_ = µsubs
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
							// line 84: template_file = open(self.document.settings.template, 'rb')
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßtemplate, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßrb.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtemplate_file = πTemp003
							// line 85: template = unicode(template_file.read(), 'utf-8')
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtemplate_file, "template_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtemplate_file, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("utf-8").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtemplate = πTemp003
							// line 86: template_file.close()
							πF.SetLineno(86)
							if πE = πg.CheckLocal(πF, µtemplate_file, "template_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtemplate_file, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 87: subs = self.interpolation_dict()
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinterpolation_dict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsubs = πTemp003
							// line 88: return template % subs
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, µtemplate, µsubs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply_template.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 90: def interpolation_dict(self):
					πF.SetLineno(90)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("interpolation_dict", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsubs *πg.Object = πg.UnboundLocal
						_ = µsubs
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µattr *πg.Object = πg.UnboundLocal
						_ = µattr
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 91: subs = {}
							πF.SetLineno(91)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µsubs = πTemp002
							// line 92: settings = self.document.settings
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							µsettings = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisitor_attributes, nil); πE != nil {
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
								µattr = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 94: subs[attr] = ''.join(getattr(self, attr)).rstrip('\n')
							πF.SetLineno(94)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("\n").ToObject()
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp008[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp008[1] = µattr
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp009
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.GetAttr(πF, πTemp009, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp009); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp010 = µattr
							if πE = πg.SetItem(πF, µsubs, πTemp010, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 95: subs['encoding'] = settings.output_encoding
							πF.SetLineno(95)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp009 = ßencoding.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp009, πTemp003); πE != nil {
								continue
							}
							// line 96: subs['version'] = docutils.__version__
							πF.SetLineno(96)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__version__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp009 = ßversion.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp009, πTemp002); πE != nil {
								continue
							}
							// line 97: return subs
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πR = µsubs
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßinterpolation_dict.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 99: def assemble_parts(self):
					πF.SetLineno(99)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("assemble_parts", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpart *πg.Object = πg.UnboundLocal
						_ = µpart
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
						var πTemp006 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 100: writers.Writer.assemble_parts(self)
							πF.SetLineno(100)
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßassemble_parts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisitor_attributes, nil); πE != nil {
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
								µpart = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 102: self.parts[part] = ''.join(getattr(self, part))
							πF.SetLineno(102)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[0] = µself
							if πE = πg.CheckLocal(πF, µpart, "part"); πE != nil {
								continue
							}
							πTemp006[1] = µpart
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp007
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßparts, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpart, "part"); πE != nil {
								continue
							}
							πTemp009 = µpart
							if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßassemble_parts.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 105: class HTMLTranslator(nodes.NodeVisitor):
			πF.SetLineno(105)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßNodeVisitor, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("HTMLTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 []πg.Param
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
				var πTemp190 *πg.Object
				_ = πTemp190
				var πTemp191 *πg.Object
				_ = πTemp191
				var πTemp192 *πg.Object
				_ = πTemp192
				var πTemp193 *πg.Object
				_ = πTemp193
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 107: """
					πF.SetLineno(107)
					// line 107: """
					πF.SetLineno(107)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Generic Docutils to HTML translator.\n\n    See the `html4css1` and `html5_polyglot` writers for full featured\n    HTML writers.\n\n    .. IMPORTANT::\n      The `visit_*` and `depart_*` methods use a\n      heterogeneous stack, `self.context`.\n      When subclassing, make sure to be consistent in its use!\n\n      Examples for robust coding:\n\n      a) Override both `visit_*` and `depart_*` methods, don't call the\n         parent functions.\n\n      b) Extend both and unconditionally call the parent functions::\n\n           def visit_example(self, node):\n               if foo:\n                   self.body.append('<div class=\"foo\">')\n               html4css1.HTMLTranslator.visit_example(self, node)\n\n           def depart_example(self, node):\n               html4css1.HTMLTranslator.depart_example(self, node)\n               if foo:\n                   self.body.append('</div>')\n\n      c) Extend both, calling the parent functions under the same\n         conditions::\n\n           def visit_example(self, node):\n               if foo:\n                   self.body.append('<div class=\"foo\">\n')\n               else: # call the parent method\n                   _html_base.HTMLTranslator.visit_example(self, node)\n\n           def depart_example(self, node):\n               if foo:\n                   self.body.append('</div>\n')\n               else: # call the parent method\n                   _html_base.HTMLTranslator.depart_example(self, node)\n\n      d) Extend one method (call the parent), but don't otherwise use the\n         `self.context` stack::\n\n           def depart_example(self, node):\n               _html_base.HTMLTranslator.depart_example(self, node)\n               if foo:\n                   # implementation-specific code\n                   # that does not use `self.context`\n                   self.body.append('</div>\n')\n\n      This way, changes in stack use will not bite you.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 163: xml_declaration = '<?xml version="1.0" encoding="%s" ?>\n'
					πF.SetLineno(163)
					if πE = πClass.SetItem(πF, ßxml_declaration.ToObject(), πg.NewStr("<?xml version=\"1.0\" encoding=\"%s\" ?>\n").ToObject()); πE != nil {
						continue
					}
					// line 164: doctype = '<!DOCTYPE html>\n'
					πF.SetLineno(164)
					if πE = πClass.SetItem(πF, ßdoctype.ToObject(), πg.NewStr("<!DOCTYPE html>\n").ToObject()); πE != nil {
						continue
					}
					// line 165: doctype_mathml = doctype
					πF.SetLineno(165)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdoctype); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdoctype_mathml.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 167: head_prefix_template = ('<html xmlns="http://www.w3.org/1999/xhtml"'
					πF.SetLineno(167)
					if πE = πClass.SetItem(πF, ßhead_prefix_template.ToObject(), πg.NewStr("<html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"%(lang)s\" lang=\"%(lang)s\">\n<head>\n").ToObject()); πE != nil {
						continue
					}
					// line 169: content_type = ('<meta charset="%s"/>\n')
					πF.SetLineno(169)
					if πE = πClass.SetItem(πF, ßcontent_type.ToObject(), πg.NewStr("<meta charset=\"%s\"/>\n").ToObject()); πE != nil {
						continue
					}
					// line 170: generator = ('<meta name="generator" content="Docutils %s: '
					πF.SetLineno(170)
					if πE = πClass.SetItem(πF, ßgenerator.ToObject(), πg.NewStr("<meta name=\"generator\" content=\"Docutils %s: http://docutils.sourceforge.net/\" />\n").ToObject()); πE != nil {
						continue
					}
					// line 174: mathjax_script = '<script type="text/javascript" src="%s"></script>\n'
					πF.SetLineno(174)
					if πE = πClass.SetItem(πF, ßmathjax_script.ToObject(), πg.NewStr("<script type=\"text/javascript\" src=\"%s\"></script>\n").ToObject()); πE != nil {
						continue
					}
					// line 176: mathjax_url = 'file:/usr/share/javascript/mathjax/MathJax.js'
					πF.SetLineno(176)
					if πE = πClass.SetItem(πF, ßmathjax_url.ToObject(), πg.NewStr("file:/usr/share/javascript/mathjax/MathJax.js").ToObject()); πE != nil {
						continue
					}
					// line 177: """
					πF.SetLineno(177)
					// line 191: stylesheet_link = '<link rel="stylesheet" href="%s" type="text/css" />\n'
					πF.SetLineno(191)
					if πE = πClass.SetItem(πF, ßstylesheet_link.ToObject(), πg.NewStr("<link rel=\"stylesheet\" href=\"%s\" type=\"text/css\" />\n").ToObject()); πE != nil {
						continue
					}
					// line 192: embedded_stylesheet = '<style type="text/css">\n\n%s\n</style>\n'
					πF.SetLineno(192)
					if πE = πClass.SetItem(πF, ßembedded_stylesheet.ToObject(), πg.NewStr("<style type=\"text/css\">\n\n%s\n</style>\n").ToObject()); πE != nil {
						continue
					}
					// line 193: words_and_spaces = re.compile(r'[^ \n]+| +|\n')
					πF.SetLineno(193)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("[^ \\n]+| +|\\n").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßwords_and_spaces.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 195: in_word_wrap_point = re.compile(r'.+\W\W.+|[-?].+', re.U)
					πF.SetLineno(195)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr(".+\\W\\W.+|[-?].+").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßU, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßin_word_wrap_point.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 196: lang_attribute = 'lang' # name changes to 'xml:lang' in XHTML 1.1
					πF.SetLineno(196)
					if πE = πClass.SetItem(πF, ßlang_attribute.ToObject(), ßlang.ToObject()); πE != nil {
						continue
					}
					// line 198: special_characters = {ord('&'): u'&amp;',
					πF.SetLineno(198)
					πTemp004 = πg.NewDict()
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("&").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πTemp004.SetItem(πF, πTemp003, πg.NewUnicode("&amp;").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("<").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πTemp004.SetItem(πF, πTemp003, πg.NewUnicode("&lt;").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("\"").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πTemp004.SetItem(πF, πTemp003, πg.NewUnicode("&quot;").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr(">").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πTemp004.SetItem(πF, πTemp003, πg.NewUnicode("&gt;").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("@").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πTemp004.SetItem(πF, πTemp003, πg.NewUnicode("&#64;").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßspecial_characters.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 204: """Character references for characters with a special meaning in HTML."""
					πF.SetLineno(204)
					// line 207: def __init__(self, document):
					πF.SetLineno(207)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []πg.Param
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
							// line 208: nodes.NodeVisitor.__init__(self, document)
							πF.SetLineno(208)
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
							// line 209: self.settings = settings = document.settings
							πF.SetLineno(209)
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
							// line 210: lcode = settings.language_code
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßlanguage_code, nil); πE != nil {
								continue
							}
							µlcode = πTemp002
							// line 211: self.language = languages.get_language(lcode, document.reporter)
							πF.SetLineno(211)
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
							// line 212: self.meta = [self.generator % docutils.__version__]
							πF.SetLineno(212)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgenerator, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__version__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmeta, πTemp003); πE != nil {
								continue
							}
							// line 213: self.head_prefix = []
							πF.SetLineno(213)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhead_prefix, πTemp003); πE != nil {
								continue
							}
							// line 214: self.html_prolog = []
							πF.SetLineno(214)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhtml_prolog, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßxml_declaration, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 215: if settings.xml_declaration:
							πF.SetLineno(215)
						Label1:
							// line 216: self.head_prefix.append(self.xml_declaration
							πF.SetLineno(216)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßxml_declaration, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µsettings, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhead_prefix, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 220: self.html_prolog.append(self.xml_declaration)
							πF.SetLineno(220)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßxml_declaration, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhtml_prolog, nil); πE != nil {
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
							// line 221: self.head = self.meta[:]
							πF.SetLineno(221)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmeta, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßhead, πTemp002); πE != nil {
								continue
							}
							// line 222: self.stylesheet = [self.stylesheet_call(path)
							πF.SetLineno(222)
							πTemp007 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µpath *πg.Object = πg.UnboundLocal
								_ = µpath
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
										if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
											continue
										}
										πTemp002[0] = µsettings
										if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget_stylesheet_list, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
											µpath = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 222: self.stylesheet = [self.stylesheet_call(path)
										πF.SetLineno(222)
										πTemp002 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
											continue
										}
										πTemp002[0] = µpath
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßstylesheet_call, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstylesheet, πTemp004); πE != nil {
								continue
							}
							// line 224: self.body_prefix = ['</head>\n<body>\n']
							πF.SetLineno(224)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = πg.NewStr("</head>\n<body>\n").ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbody_prefix, πTemp004); πE != nil {
								continue
							}
							// line 226: self.body_pre_docinfo = []
							πF.SetLineno(226)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbody_pre_docinfo, πTemp004); πE != nil {
								continue
							}
							// line 228: self.docinfo = []
							πF.SetLineno(228)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocinfo, πTemp004); πE != nil {
								continue
							}
							// line 229: self.body = []
							πF.SetLineno(229)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbody, πTemp004); πE != nil {
								continue
							}
							// line 230: self.fragment = []
							πF.SetLineno(230)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfragment, πTemp004); πE != nil {
								continue
							}
							// line 231: self.body_suffix = ['</body>\n</html>\n']
							πF.SetLineno(231)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = πg.NewStr("</body>\n</html>\n").ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbody_suffix, πTemp004); πE != nil {
								continue
							}
							// line 232: self.section_level = 0
							πF.SetLineno(232)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_level, πTemp002); πE != nil {
								continue
							}
							// line 233: self.initial_header_level = int(settings.initial_header_level)
							πF.SetLineno(233)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßinitial_header_level, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinitial_header_level, πTemp002); πE != nil {
								continue
							}
							// line 235: self.math_output = settings.math_output.split()
							πF.SetLineno(235)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_output, πTemp004); πE != nil {
								continue
							}
							// line 236: self.math_output_options = self.math_output[1:]
							πF.SetLineno(236)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_output_options, πTemp002); πE != nil {
								continue
							}
							// line 237: self.math_output = self.math_output[0].lower()
							πF.SetLineno(237)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßlower, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_output, πTemp002); πE != nil {
								continue
							}
							// line 239: self.context = []
							πF.SetLineno(239)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontext, πTemp004); πE != nil {
								continue
							}
							// line 240: """Heterogeneous stack.
							πF.SetLineno(240)
							// line 245: self.topic_classes = [] # TODO: replace with self_in_contents
							πF.SetLineno(245)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtopic_classes, πTemp004); πE != nil {
								continue
							}
							// line 246: self.colspecs = []
							πF.SetLineno(246)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcolspecs, πTemp004); πE != nil {
								continue
							}
							// line 247: self.compact_p = True
							πF.SetLineno(247)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
							// line 248: self.compact_simple = False
							πF.SetLineno(248)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_simple, πTemp004); πE != nil {
								continue
							}
							// line 249: self.compact_field_list = False
							πF.SetLineno(249)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcompact_field_list, πTemp004); πE != nil {
								continue
							}
							// line 250: self.in_docinfo = False
							πF.SetLineno(250)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_docinfo, πTemp004); πE != nil {
								continue
							}
							// line 251: self.in_sidebar = False
							πF.SetLineno(251)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_sidebar, πTemp004); πE != nil {
								continue
							}
							// line 252: self.in_footnote_list = False
							πF.SetLineno(252)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_footnote_list, πTemp004); πE != nil {
								continue
							}
							// line 253: self.title = []
							πF.SetLineno(253)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtitle, πTemp004); πE != nil {
								continue
							}
							// line 254: self.subtitle = []
							πF.SetLineno(254)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsubtitle, πTemp004); πE != nil {
								continue
							}
							// line 255: self.header = []
							πF.SetLineno(255)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßheader, πTemp004); πE != nil {
								continue
							}
							// line 256: self.footer = []
							πF.SetLineno(256)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfooter, πTemp004); πE != nil {
								continue
							}
							// line 257: self.html_head = [self.content_type] # charset not interpolated
							πF.SetLineno(257)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent_type, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhtml_head, πTemp004); πE != nil {
								continue
							}
							// line 258: self.html_title = []
							πF.SetLineno(258)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhtml_title, πTemp004); πE != nil {
								continue
							}
							// line 259: self.html_subtitle = []
							πF.SetLineno(259)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhtml_subtitle, πTemp004); πE != nil {
								continue
							}
							// line 260: self.html_body = []
							πF.SetLineno(260)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhtml_body, πTemp004); πE != nil {
								continue
							}
							// line 261: self.in_document_title = 0   # len(self.body) or 0
							πF.SetLineno(261)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_document_title, πTemp002); πE != nil {
								continue
							}
							// line 262: self.in_mailto = False
							πF.SetLineno(262)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_mailto, πTemp004); πE != nil {
								continue
							}
							// line 263: self.author_in_authors = False # for html4css1
							πF.SetLineno(263)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßauthor_in_authors, πTemp004); πE != nil {
								continue
							}
							// line 264: self.math_header = []
							πF.SetLineno(264)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_header, πTemp004); πE != nil {
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
					// line 266: def astext(self):
					πF.SetLineno(266)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 267: return ''.join(self.head_prefix + self.head
							πF.SetLineno(267)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßhead_prefix, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßstylesheet, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody_pre_docinfo, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody_suffix, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 272: def encode(self, text):
					πF.SetLineno(272)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "text", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("encode", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 273: """Encode special characters in `text` & return."""
							πF.SetLineno(273)
							// line 277: text = unicode(text)
							πF.SetLineno(277)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp003
							// line 278: return text.translate(self.special_characters)
							πF.SetLineno(278)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßspecial_characters, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßtranslate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßencode.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 273: """Encode special characters in `text` & return."""
					πF.SetLineno(273)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Encode special characters in `text` & return.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßencode); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 280: def cloak_mailto(self, uri):
					πF.SetLineno(280)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "uri", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("cloak_mailto", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µuri *πg.Object = πArgs[1]
						_ = µuri
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
							// line 281: """Try to hide a mailto: URL from harvesters."""
							πF.SetLineno(281)
							// line 285: return uri.replace('@', '%40')
							πF.SetLineno(285)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("@").ToObject()
							πTemp001[1] = πg.NewStr("%40").ToObject()
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µuri, ßreplace, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcloak_mailto.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 281: """Try to hide a mailto: URL from harvesters."""
					πF.SetLineno(281)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Try to hide a mailto: URL from harvesters.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßcloak_mailto); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 287: def cloak_email(self, addr):
					πF.SetLineno(287)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "addr", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("cloak_email", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µaddr *πg.Object = πArgs[1]
						_ = µaddr
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
							// line 288: """Try to hide the link text of a email link from harversters."""
							πF.SetLineno(288)
							// line 291: addr = addr.replace('&#64;', '<span>&#64;</span>')
							πF.SetLineno(291)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("&#64;").ToObject()
							πTemp001[1] = πg.NewStr("<span>&#64;</span>").ToObject()
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µaddr, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µaddr = πTemp003
							// line 292: addr = addr.replace('.', '<span>&#46;</span>')
							πF.SetLineno(292)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr(".").ToObject()
							πTemp001[1] = πg.NewStr("<span>&#46;</span>").ToObject()
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µaddr, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µaddr = πTemp003
							// line 293: return addr
							πF.SetLineno(293)
							if πE = πg.CheckLocal(πF, µaddr, "addr"); πE != nil {
								continue
							}
							πR = µaddr
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcloak_email.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 288: """Try to hide the link text of a email link from harversters."""
					πF.SetLineno(288)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Try to hide the link text of a email link from harversters.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßcloak_email); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 295: def attval(self, text,
					πF.SetLineno(295)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "text", Def: nil}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("[\n\r\t\x0b\x0c]").ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp005[2] = πg.Param{Name: "whitespace", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("attval", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πArgs[1]
						_ = µtext
						var µwhitespace *πg.Object = πArgs[2]
						_ = µwhitespace
						var µencoded *πg.Object = πg.UnboundLocal
						_ = µencoded
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 297: """Cleanse, HTML encode, and return attribute value text."""
							πF.SetLineno(297)
							// line 298: encoded = self.encode(whitespace.sub(' ', text))
							πF.SetLineno(298)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp002[1] = µtext
							if πE = πg.CheckLocal(πF, µwhitespace, "whitespace"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µwhitespace, ßsub, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µencoded = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßin_mailto, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßcloak_email_addresses, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
						Label1:
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 299: if self.in_mailto and self.settings.cloak_email_addresses:
							πF.SetLineno(299)
						Label2:
							// line 301: encoded = encoded.replace('%40', '&#37;&#52;&#48;')
							πF.SetLineno(301)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("%40").ToObject()
							πTemp001[1] = πg.NewStr("&#37;&#52;&#48;").ToObject()
							if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µencoded, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µencoded = πTemp004
							// line 302: encoded = encoded.replace('.', '&#46;')
							πF.SetLineno(302)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr(".").ToObject()
							πTemp001[1] = πg.NewStr("&#46;").ToObject()
							if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µencoded, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µencoded = πTemp004
							goto Label3
						Label3:
							// line 303: return encoded
							πF.SetLineno(303)
							if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
								continue
							}
							πR = µencoded
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßattval.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 297: """Cleanse, HTML encode, and return attribute value text."""
					πF.SetLineno(297)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Cleanse, HTML encode, and return attribute value text.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßattval); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 305: def stylesheet_call(self, path):
					πF.SetLineno(305)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "path", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("stylesheet_call", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpath *πg.Object = πArgs[1]
						_ = µpath
						var µcontent *πg.Object = πg.UnboundLocal
						_ = µcontent
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
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
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 306: """Return code to reference or embed stylesheet file `path`"""
							πF.SetLineno(306)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßembed_stylesheet, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 307: if self.settings.embed_stylesheet:
							πF.SetLineno(307)
						Label1:
							// line 308: try:
							πF.SetLineno(308)
							πF.PushCheckpoint(4)
							// line 309: content = io.FileInput(source_path=path,
							πF.SetLineno(309)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"source_path", µpath},
								{"encoding", πg.NewStr("utf-8").ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFileInput, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcontent = πTemp001
							// line 311: self.settings.record_dependencies.add(path)
							πF.SetLineno(311)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 312: except IOError as err:
							πF.SetLineno(312)
						Label5:
							µerr = πTemp006.ToObject()
							// line 313: msg = u"Cannot embed stylesheet '%s': %s." % (
							πF.SetLineno(313)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µerr, ßstrerror, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							if πTemp008, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002 = πg.NewTuple2(µpath, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Cannot embed stylesheet '%s': %s.").ToObject(), πTemp002); πE != nil {
								continue
							}
							µmsg = πTemp001
							// line 315: self.document.reporter.error(msg)
							πF.SetLineno(315)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp005[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 316: return '<--- %s --->\n' % msg
							πF.SetLineno(316)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<--- %s --->\n").ToObject(), µmsg); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 317: return self.embedded_stylesheet % content
							πF.SetLineno(317)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßembedded_stylesheet, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, µcontent); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstylesheet_path, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 319: if self.settings.stylesheet_path:
							πF.SetLineno(319)
						Label6:
							// line 321: path = utils.relative_path(self.settings._destination, path)
							πF.SetLineno(321)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_destination, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005[1] = µpath
							if πTemp001, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelative_path, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpath = πTemp001
							goto Label7
						Label7:
							// line 322: return self.stylesheet_link % self.encode(path)
							πF.SetLineno(322)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstylesheet_link, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp009); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstylesheet_call.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 306: """Return code to reference or embed stylesheet file `path`"""
					πF.SetLineno(306)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Return code to reference or embed stylesheet file `path`").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßstylesheet_call); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 324: def starttag(self, node, tagname, suffix='\n', empty=False, **attributes):
					πF.SetLineno(324)
					πTemp005 = make([]πg.Param, 5)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp005[2] = πg.Param{Name: "tagname", Def: nil}
					πTemp005[3] = πg.Param{Name: "suffix", Def: πg.NewStr("\n").ToObject()}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp005[4] = πg.Param{Name: "empty", Def: πTemp012}
					πTemp011 = πg.NewFunction(πg.NewCode("starttag", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtagname *πg.Object = πArgs[2]
						_ = µtagname
						var µsuffix *πg.Object = πArgs[3]
						_ = µsuffix
						var µempty *πg.Object = πArgs[4]
						_ = µempty
						var µattributes *πg.Object = πArgs[5]
						_ = µattributes
						var µprefix *πg.Object = πg.UnboundLocal
						_ = µprefix
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µids *πg.Object = πg.UnboundLocal
						_ = µids
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µlanguages *πg.Object = πg.UnboundLocal
						_ = µlanguages
						var µcls *πg.Object = πg.UnboundLocal
						_ = µcls
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µattlist *πg.Object = πg.UnboundLocal
						_ = µattlist
						var µparts *πg.Object = πg.UnboundLocal
						_ = µparts
						var µvalues *πg.Object = πg.UnboundLocal
						_ = µvalues
						var µinfix *πg.Object = πg.UnboundLocal
						_ = µinfix
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 []πg.Param
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 *πg.Object
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
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
							case 19:
								goto Label19
							case 20:
								goto Label20
							case 26:
								goto Label26
							case 27:
								goto Label27
							default:
								panic("unexpected function state")
							}
							// line 325: """
							πF.SetLineno(325)
							// line 329: tagname = tagname.lower()
							πF.SetLineno(329)
							if πE = πg.CheckLocal(πF, µtagname, "tagname"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtagname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtagname = πTemp002
							// line 330: prefix = []
							πF.SetLineno(330)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µprefix = πTemp001
							// line 331: atts = {}
							πF.SetLineno(331)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							µatts = πTemp001
							// line 332: ids = []
							πF.SetLineno(332)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µids = πTemp001
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattributes, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µname = πTemp005
								µvalue = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 334: atts[name.lower()] = value
							πF.SetLineno(334)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp009
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 335: classes = []
							πF.SetLineno(335)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µclasses = πTemp001
							// line 336: languages = []
							πF.SetLineno(336)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µlanguages = πTemp001
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßclasses.ToObject()
							πTemp010 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp010...).ToObject()
							πTemp003[1] = πTemp005
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßclass.ToObject()
							πTemp003[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µatts, ßpop, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.GetAttr(πF, πTemp009, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µcls = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("language-").ToObject()
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp008
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, µclasses, µcls); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp011).ToObject()
							πTemp002 = πTemp005
						Label8:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 339: if cls.startswith('language-'):
							πF.SetLineno(339)
						Label7:
							// line 340: languages.append(cls[9:])
							πF.SetLineno(340)
							πTemp003 = πF.MakeArgs(1)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(9).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µcls, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µlanguages, "languages"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlanguages, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label10
							// line 341: elif cls.strip() and cls not in classes:
							πF.SetLineno(341)
						Label9:
							// line 342: classes.append(cls)
							πF.SetLineno(342)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp003[0] = µcls
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µclasses, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label10
						Label10:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							if πE = πg.CheckLocal(πF, µlanguages, "languages"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µlanguages); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							goto Label12
							// line 343: if languages:
							πF.SetLineno(343)
						Label11:
							// line 345: atts[self.lang_attribute] = languages[0]
							πF.SetLineno(345)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µlanguages, "languages"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µlanguages, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßlang_attribute, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp008
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp001); πE != nil {
								continue
							}
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µclasses); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label13
							}
							goto Label14
							// line 346: if classes:
							πF.SetLineno(346)
						Label13:
							// line 347: atts['class'] = ' '.join(classes)
							πF.SetLineno(347)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp003[0] = µclasses
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp001); πE != nil {
								continue
							}
							goto Label14
						Label14:
							// line 348: assert 'id' not in atts
							πF.SetLineno(348)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µatts, ßid.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 349: ids.extend(node.get('ids', []))
							πF.SetLineno(349)
							πTemp003 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = ßids.ToObject()
							πTemp012 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp012...).ToObject()
							πTemp010[1] = πTemp001
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µids, "ids"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µids, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µatts, ßids.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label15
							}
							goto Label16
							// line 350: if 'ids' in atts:
							πF.SetLineno(350)
						Label15:
							// line 351: ids.extend(atts['ids'])
							πF.SetLineno(351)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µatts, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µids, "ids"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µids, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 352: del atts['ids']
							πF.SetLineno(352)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp001 = ßids.ToObject()
							if πE = πg.DelItem(πF, µatts, πTemp001); πE != nil {
								continue
							}
							goto Label16
						Label16:
							if πE = πg.CheckLocal(πF, µids, "ids"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µids); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label17
							}
							goto Label18
							// line 353: if ids:
							πF.SetLineno(353)
						Label17:
							// line 354: atts['id'] = ids[0]
							πF.SetLineno(354)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µids, "ids"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µids, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßid.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µids, "ids"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µids, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(20)
							πTemp006 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label21
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µid = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(19)
							if πE = πg.CheckLocal(πF, µempty, "empty"); πE != nil {
								continue
							}
							πTemp002 = µempty
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label22
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							πTemp010 = make([]*πg.Object, 7)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßbullet_list, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßdocinfo, nil); πE != nil {
								continue
							}
							πTemp010[1] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßdefinition_list, nil); πE != nil {
								continue
							}
							πTemp010[2] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßenumerated_list, nil); πE != nil {
								continue
							}
							πTemp010[3] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßfield_list, nil); πE != nil {
								continue
							}
							πTemp010[4] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßoption_list, nil); πE != nil {
								continue
							}
							πTemp010[5] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßtable, nil); πE != nil {
								continue
							}
							πTemp010[6] = πTemp009
							πTemp005 = πg.NewTuple(πTemp010...).ToObject()
							πTemp003[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πTemp008
						Label22:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label23
							}
							goto Label24
							// line 361: if empty or isinstance(node,
							πF.SetLineno(361)
						Label23:
							// line 367: prefix.append('<span id="%s"></span>' % id)
							πF.SetLineno(367)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<span id=\"%s\"></span>").ToObject(), µid); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µprefix, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label25
						Label24:
							// line 371: suffix += '<span id="%s"></span>' % id
							πF.SetLineno(371)
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<span id=\"%s\"></span>").ToObject(), µid); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µsuffix, πTemp002); πE != nil {
								continue
							}
							µsuffix = πTemp005
							goto Label25
						Label25:
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
							goto Label18
						Label18:
							// line 372: attlist = sorted(atts.items())
							πF.SetLineno(372)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µatts, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µattlist = πTemp002
							// line 373: parts = [tagname]
							πF.SetLineno(373)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtagname, "tagname"); πE != nil {
								continue
							}
							πTemp003[0] = µtagname
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µparts = πTemp001
							if πE = πg.CheckLocal(πF, µattlist, "attlist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µattlist); πE != nil {
								continue
							}
							πF.PushCheckpoint(27)
							πTemp006 = false
						Label26:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label28
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µname = πTemp005
								µvalue = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(26)
							// line 377: assert value is not None
							πF.SetLineno(377)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µvalue != πTemp005).ToObject()
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[0] = µvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label29
							}
							goto Label30
							// line 378: if isinstance(value, list):
							πF.SetLineno(378)
						Label29:
							// line 379: values = [unicode(v) for v in value]
							πF.SetLineno(379)
							πTemp013 = make([]πg.Param, 0)
							πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µv *πg.Object = πg.UnboundLocal
								_ = µv
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
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µvalue); πE != nil {
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
											µv = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 379: values = [unicode(v) for v in value]
										πF.SetLineno(379)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp005[0] = µv
										if πTemp004, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
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
							if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							µvalues = πTemp002
							// line 380: parts.append('%s="%s"' % (name.lower(),
							πF.SetLineno(380)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp010 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp012[0] = µvalues
							if πTemp009, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp009.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp010[0] = πTemp015
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßattval, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp008 = πg.NewTuple2(πTemp014, πTemp015).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s=\"%s\"").ToObject(), πTemp008); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label31
						Label30:
							// line 383: parts.append('%s="%s"' % (name.lower(),
							πF.SetLineno(383)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp010 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp012[0] = µvalue
							if πTemp009, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp009.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp010[0] = πTemp015
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßattval, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp008 = πg.NewTuple2(πTemp014, πTemp015).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s=\"%s\"").ToObject(), πTemp008); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label31
						Label31:
							continue
						Label27:
							if πE != nil || πR != nil {
								continue
							}
						Label28:
							if πE = πg.CheckLocal(πF, µempty, "empty"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µempty); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label32
							}
							goto Label33
							// line 385: if empty:
							πF.SetLineno(385)
						Label32:
							// line 386: infix = ' /'
							πF.SetLineno(386)
							µinfix = πg.NewStr(" /").ToObject()
							goto Label34
						Label33:
							// line 388: infix = ''
							πF.SetLineno(388)
							µinfix = ß.ToObject()
							goto Label34
						Label34:
							// line 389: return ''.join(prefix) + '<%s%s>' % (' '.join(parts), infix) + suffix
							πF.SetLineno(389)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							πTemp003[0] = µprefix
							if πTemp008, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							πTemp003[0] = µparts
							if πTemp015, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp015.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µinfix, "infix"); πE != nil {
								continue
							}
							πTemp014 = πg.NewTuple2(πTemp016, µinfix).ToObject()
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("<%s%s>").ToObject(), πTemp014); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp009, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µsuffix); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstarttag.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 325: """
					πF.SetLineno(325)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n        Construct and return a start tag given a node (id & class attributes\n        are extracted), tag name, and optional attributes.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßstarttag); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 391: def emptytag(self, node, tagname, suffix='\n', **attributes):
					πF.SetLineno(391)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp005[2] = πg.Param{Name: "tagname", Def: nil}
					πTemp005[3] = πg.Param{Name: "suffix", Def: πg.NewStr("\n").ToObject()}
					πTemp012 = πg.NewFunction(πg.NewCode("emptytag", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtagname *πg.Object = πArgs[2]
						_ = µtagname
						var µsuffix *πg.Object = πArgs[3]
						_ = µsuffix
						var µattributes *πg.Object = πArgs[4]
						_ = µattributes
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							// line 392: """Construct and return an XML-compatible empty tag."""
							πF.SetLineno(392)
							// line 393: return self.starttag(node, tagname, suffix, empty=True, **attributes)
							πF.SetLineno(393)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µtagname, "tagname"); πE != nil {
								continue
							}
							πTemp001[1] = µtagname
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							πTemp001[2] = µsuffix
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"empty", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, πTemp003, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp004
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßemptytag.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 392: """Construct and return an XML-compatible empty tag."""
					πF.SetLineno(392)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Construct and return an XML-compatible empty tag.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßemptytag); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 395: def set_class_on_child(self, node, class_, index=0):
					πF.SetLineno(395)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp005[2] = πg.Param{Name: "class_", Def: nil}
					πTemp005[3] = πg.Param{Name: "index", Def: πg.NewInt(0).ToObject()}
					πTemp013 = πg.NewFunction(πg.NewCode("set_class_on_child", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclass_ *πg.Object = πArgs[2]
						_ = µclass_
						var µindex *πg.Object = πArgs[3]
						_ = µindex
						var µchildren *πg.Object = πg.UnboundLocal
						_ = µchildren
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
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
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 396: """
							πF.SetLineno(396)
							// line 400: children = [n for n in node if not isinstance(n, nodes.Invisible)]
							πF.SetLineno(400)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µnode); πE != nil {
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
										// line 400: children = [n for n in node if not isinstance(n, nodes.Invisible)]
										πF.SetLineno(400)
									Label4:
										// line 400: children = [n for n in node if not isinstance(n, nodes.Invisible)]
										πF.SetLineno(400)
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µchildren = πTemp001
							// line 401: try:
							πF.SetLineno(401)
							πF.PushCheckpoint(2)
							// line 402: child = children[index]
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = µindex
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µchildren, πTemp001); πE != nil {
								continue
							}
							µchild = πTemp004
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 403: except IndexError:
							πF.SetLineno(403)
						Label3:
							// line 404: return
							πF.SetLineno(404)
							πR = πg.None
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 405: child['classes'].append(class_)
							πF.SetLineno(405)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclass_, "class_"); πE != nil {
								continue
							}
							πTemp008[0] = µclass_
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µchild, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßset_class_on_child.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 396: """
					πF.SetLineno(396)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n        Set class `class_` on the visible child no. index of `node`.\n        Do nothing if node has fewer children than `index`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßset_class_on_child); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 407: def visit_Text(self, node):
					πF.SetLineno(407)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("visit_Text", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µencoded *πg.Object = πg.UnboundLocal
						_ = µencoded
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 408: text = node.astext()
							πF.SetLineno(408)
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
							// line 409: encoded = self.encode(text)
							πF.SetLineno(409)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µencoded = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_mailto, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßcloak_email_addresses, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp005
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 410: if self.in_mailto and self.settings.cloak_email_addresses:
							πF.SetLineno(410)
						Label2:
							// line 411: encoded = self.cloak_email(encoded)
							πF.SetLineno(411)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
								continue
							}
							πTemp003[0] = µencoded
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcloak_email, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µencoded = πTemp002
							goto Label3
						Label3:
							// line 412: self.body.append(encoded)
							πF.SetLineno(412)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µencoded, "encoded"); πE != nil {
								continue
							}
							πTemp003[0] = µencoded
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
					if πE = πClass.SetItem(πF, ßvisit_Text.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 414: def depart_Text(self, node):
					πF.SetLineno(414)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("depart_Text", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 415: pass
							πF.SetLineno(415)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_Text.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 417: def visit_abbreviation(self, node):
					πF.SetLineno(417)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("visit_abbreviation", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 419: self.body.append(self.starttag(node, 'abbr', ''))
							πF.SetLineno(419)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßabbr.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_abbreviation.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 421: def depart_abbreviation(self, node):
					πF.SetLineno(421)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("depart_abbreviation", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 422: self.body.append('</abbr>')
							πF.SetLineno(422)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</abbr>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_abbreviation.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 424: def visit_acronym(self, node):
					πF.SetLineno(424)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("visit_acronym", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 426: self.body.append(self.starttag(node, 'acronym', ''))
							πF.SetLineno(426)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßacronym.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_acronym.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 428: def depart_acronym(self, node):
					πF.SetLineno(428)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("depart_acronym", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 429: self.body.append('</acronym>')
							πF.SetLineno(429)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</acronym>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_acronym.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 431: def visit_address(self, node):
					πF.SetLineno(431)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("visit_address", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 432: self.visit_docinfo_item(node, 'address', meta=False)
							πF.SetLineno(432)
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
							// line 433: self.body.append(self.starttag(node, 'pre',
							πF.SetLineno(433)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßpre.ToObject()
							πTemp003 = πg.KWArgs{
								{"suffix", ß.ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_address.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 436: def depart_address(self, node):
					πF.SetLineno(436)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("depart_address", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 437: self.body.append('\n</pre>\n')
							πF.SetLineno(437)
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
							// line 438: self.depart_docinfo_item()
							πF.SetLineno(438)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_address.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 440: def visit_admonition(self, node):
					πF.SetLineno(440)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("visit_admonition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 441: node['classes'].insert(0, 'admonition')
							πF.SetLineno(441)
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
							// line 442: self.body.append(self.starttag(node, 'div'))
							πF.SetLineno(442)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_admonition.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 444: def depart_admonition(self, node=None):
					πF.SetLineno(444)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πg.Param{Name: "node", Def: πTemp024}
					πTemp023 = πg.NewFunction(πg.NewCode("depart_admonition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 445: self.body.append('</div>\n')
							πF.SetLineno(445)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_admonition.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 447: attribution_formats = {'dash': (u'\u2014', ''),
					πF.SetLineno(447)
					πTemp004 = πg.NewDict()
					πTemp024 = πg.NewTuple2(πg.NewUnicode("\xe2\x80\x94").ToObject(), ß.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßdash.ToObject(), πTemp024); πE != nil {
						continue
					}
					πTemp024 = πg.NewTuple2(πg.NewStr("(").ToObject(), πg.NewStr(")").ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßparentheses.ToObject(), πTemp024); πE != nil {
						continue
					}
					πTemp024 = πg.NewTuple2(πg.NewStr("(").ToObject(), πg.NewStr(")").ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßparens.ToObject(), πTemp024); πE != nil {
						continue
					}
					πTemp024 = πg.NewTuple2(ß.ToObject(), ß.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßnone.ToObject(), πTemp024); πE != nil {
						continue
					}
					πTemp024 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßattribution_formats.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 452: def visit_attribution(self, node):
					πF.SetLineno(452)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("visit_attribution", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µprefix *πg.Object = πg.UnboundLocal
						_ = µprefix
						var µsuffix *πg.Object = πg.UnboundLocal
						_ = µsuffix
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 453: prefix, suffix = self.attribution_formats[self.settings.attribution]
							πF.SetLineno(453)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßattribution, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßattribution_formats, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µprefix = πTemp001
							µsuffix = πTemp003
							// line 454: self.context.append(suffix)
							πF.SetLineno(454)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							πTemp004[0] = µsuffix
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 455: self.body.append(
							πF.SetLineno(455)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßp.ToObject()
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							πTemp005[2] = µprefix
							πTemp006 = πg.KWArgs{
								{"CLASS", ßattribution.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßvisit_attribution.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 458: def depart_attribution(self, node):
					πF.SetLineno(458)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("depart_attribution", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 459: self.body.append(self.context.pop() + '</p>\n')
							πF.SetLineno(459)
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
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("</p>\n").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_attribution.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 461: def visit_author(self, node):
					πF.SetLineno(461)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("visit_author", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßauthors, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
							// line 462: if not(isinstance(node.parent, nodes.authors)):
							πF.SetLineno(462)
						Label1:
							// line 463: self.visit_docinfo_item(node, 'author')
							πF.SetLineno(463)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßauthor.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							// line 464: self.body.append('<p>')
							πF.SetLineno(464)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("<p>").ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_author.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 466: def depart_author(self, node):
					πF.SetLineno(466)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("depart_author", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 467: self.body.append('</p>')
							πF.SetLineno(467)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>").ToObject()
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
							// line 468: if isinstance(node.parent, nodes.authors):
							πF.SetLineno(468)
						Label1:
							// line 469: self.body.append('\n')
							πF.SetLineno(469)
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
							goto Label3
						Label2:
							// line 471: self.depart_docinfo_item()
							πF.SetLineno(471)
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
					if πE = πClass.SetItem(πF, ßdepart_author.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 473: def visit_authors(self, node):
					πF.SetLineno(473)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("visit_authors", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 474: self.visit_docinfo_item(node, 'authors')
							πF.SetLineno(474)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_authors.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 476: def depart_authors(self, node):
					πF.SetLineno(476)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("depart_authors", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 477: self.depart_docinfo_item()
							πF.SetLineno(477)
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
					if πE = πClass.SetItem(πF, ßdepart_authors.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 479: def visit_block_quote(self, node):
					πF.SetLineno(479)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("visit_block_quote", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 480: self.body.append(self.starttag(node, 'blockquote'))
							πF.SetLineno(480)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßblockquote.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_block_quote.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 482: def depart_block_quote(self, node):
					πF.SetLineno(482)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("depart_block_quote", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 483: self.body.append('</blockquote>\n')
							πF.SetLineno(483)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</blockquote>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_block_quote.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 485: def check_simple_list(self, node):
					πF.SetLineno(485)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("check_simple_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
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
							// line 486: """Check for a simple list that can be rendered compactly."""
							πF.SetLineno(486)
							// line 487: visitor = SimpleListChecker(self.document)
							πF.SetLineno(487)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSimpleListChecker); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvisitor = πTemp003
							// line 488: try:
							πF.SetLineno(488)
							πF.PushCheckpoint(2)
							// line 489: node.walk(visitor)
							πF.SetLineno(489)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp001[0] = µvisitor
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßwalk, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							// line 493: return True
							πF.SetLineno(493)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßNodeFound, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 490: except nodes.NodeFound:
							πF.SetLineno(490)
						Label3:
							// line 491: return False
							πF.SetLineno(491)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ßcheck_simple_list.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 486: """Check for a simple list that can be rendered compactly."""
					πF.SetLineno(486)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πg.NewStr("Check for a simple list that can be rendered compactly.").ToObject()); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_simple_list); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp034, ß__doc__, πTemp033); πE != nil {
						continue
					}
					// line 502: def is_compactable(self, node):
					πF.SetLineno(502)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("is_compactable", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							default:
								panic("unexpected function state")
							}
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßcompact.ToObject()); πE != nil {
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
							// line 504: if 'compact' in node['classes']:
							πF.SetLineno(504)
						Label1:
							// line 505: return True
							πF.SetLineno(505)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßopen.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 506: if 'open' in node['classes']:
							πF.SetLineno(506)
						Label3:
							// line 507: return False
							πF.SetLineno(507)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßfield_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßdefinition_list, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							πTemp005[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcompact_field_lists, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							πTemp001 = πTemp002
						Label5:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 509: if (isinstance(node, (nodes.field_list, nodes.definition_list))
							πF.SetLineno(509)
						Label6:
							// line 511: return False
							πF.SetLineno(511)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label7
						Label7:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßenumerated_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßbullet_list, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							πTemp005[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcompact_lists, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							πTemp001 = πTemp002
						Label8:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 512: if (isinstance(node, (nodes.enumerated_list, nodes.bullet_list))
							πF.SetLineno(512)
						Label9:
							// line 514: return False
							πF.SetLineno(514)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtopic_classes, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = ßcontents.ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 516: if (self.topic_classes == ['contents']): # TODO: self.in_contents
							πF.SetLineno(516)
						Label11:
							// line 517: return True
							πF.SetLineno(517)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label12
						Label12:
							// line 519: return self.check_simple_list(node)
							πF.SetLineno(519)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_simple_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßis_compactable.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 521: def visit_bullet_list(self, node):
					πF.SetLineno(521)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("visit_bullet_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 522: atts = {}
							πF.SetLineno(522)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							// line 523: old_compact_simple = self.compact_simple
							πF.SetLineno(523)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcompact_simple, nil); πE != nil {
								continue
							}
							µold_compact_simple = πTemp002
							// line 524: self.context.append((self.compact_simple, self.compact_p))
							πF.SetLineno(524)
							πTemp003 = πF.MakeArgs(1)
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
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 525: self.compact_p = None
							πF.SetLineno(525)
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
							// line 526: self.compact_simple = self.is_compactable(node)
							πF.SetLineno(526)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßis_compactable, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µold_compact_simple, "old_compact_simple"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µold_compact_simple); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp004
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 527: if self.compact_simple and not old_compact_simple:
							πF.SetLineno(527)
						Label2:
							// line 528: atts['class'] = 'simple'
							πF.SetLineno(528)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßsimple.ToObject()); πE != nil {
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
							// line 529: self.body.append(self.starttag(node, 'ul', **atts))
							πF.SetLineno(529)
							πTemp003 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008[0] = µnode
							πTemp008[1] = ßul.ToObject()
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
							πTemp003[0] = πTemp004
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
					if πE = πClass.SetItem(πF, ßvisit_bullet_list.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 531: def depart_bullet_list(self, node):
					πF.SetLineno(531)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("depart_bullet_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 532: self.compact_simple, self.compact_p = self.context.pop()
							πF.SetLineno(532)
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
							// line 533: self.body.append('</ul>\n')
							πF.SetLineno(533)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("</ul>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_bullet_list.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 535: def visit_caption(self, node):
					πF.SetLineno(535)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("visit_caption", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 536: self.body.append(self.starttag(node, 'p', '', CLASS='caption'))
							πF.SetLineno(536)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßp.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßcaption.ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_caption.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 538: def depart_caption(self, node):
					πF.SetLineno(538)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("depart_caption", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 539: self.body.append('</p>\n')
							πF.SetLineno(539)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_caption.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 546: def visit_citation(self, node):
					πF.SetLineno(546)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("visit_citation", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_footnote_list, nil); πE != nil {
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
							// line 547: if not self.in_footnote_list:
							πF.SetLineno(547)
						Label1:
							// line 548: self.body.append('<dl class="citation">\n')
							πF.SetLineno(548)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("<dl class=\"citation\">\n").ToObject()
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
							// line 549: self.in_footnote_list = True
							πF.SetLineno(549)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_footnote_list, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_citation.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 551: def depart_citation(self, node):
					πF.SetLineno(551)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("depart_citation", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 552: self.body.append('</dd>\n')
							πF.SetLineno(552)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dd>\n").ToObject()
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
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"descend", πTemp003},
								{"siblings", πTemp004},
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßnext_node, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcitation, nil); πE != nil {
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 553: if not isinstance(node.next_node(descend=False, siblings=True),
							πF.SetLineno(553)
						Label1:
							// line 555: self.body.append('</dl>\n')
							πF.SetLineno(555)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dl>\n").ToObject()
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
							// line 556: self.in_footnote_list = False
							πF.SetLineno(556)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_footnote_list, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_citation.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 558: def visit_citation_reference(self, node):
					πF.SetLineno(558)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp040 = πg.NewFunction(πg.NewCode("visit_citation_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µhref *πg.Object = πg.UnboundLocal
						_ = µhref
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
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
							// line 559: href = '#'
							πF.SetLineno(559)
							µhref = πg.NewStr("#").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnode, ßrefid.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnode, ßrefname.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 560: if 'refid' in node:
							πF.SetLineno(560)
						Label1:
							// line 561: href += node['refid']
							πF.SetLineno(561)
							if πE = πg.CheckLocal(πF, µhref, "href"); πE != nil {
								continue
							}
							πTemp001 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µhref, πTemp003); πE != nil {
								continue
							}
							µhref = πTemp001
							goto Label3
							// line 562: elif 'refname' in node:
							πF.SetLineno(562)
						Label2:
							// line 563: href += self.document.nameids[node['refname']]
							πF.SetLineno(563)
							if πE = πg.CheckLocal(πF, µhref, "href"); πE != nil {
								continue
							}
							πTemp003 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µhref, πTemp003); πE != nil {
								continue
							}
							µhref = πTemp001
							goto Label3
						Label3:
							// line 566: self.body.append(self.starttag(
							πF.SetLineno(566)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							πTemp007[1] = ßa.ToObject()
							πTemp007[2] = πg.NewStr("[").ToObject()
							if πE = πg.CheckLocal(πF, µhref, "href"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"CLASS", πg.NewStr("citation-reference").ToObject()},
								{"href", µhref},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_citation_reference.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 569: def depart_citation_reference(self, node):
					πF.SetLineno(569)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("depart_citation_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 570: self.body.append(']</a>')
							πF.SetLineno(570)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("]</a>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_citation_reference.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 576: def visit_classifier(self, node):
					πF.SetLineno(576)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp042 = πg.NewFunction(πg.NewCode("visit_classifier", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 577: self.body.append(self.starttag(node, 'span', '', CLASS='classifier'))
							πF.SetLineno(577)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßspan.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßclassifier.ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_classifier.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 579: def depart_classifier(self, node):
					πF.SetLineno(579)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp043 = πg.NewFunction(πg.NewCode("depart_classifier", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 580: self.body.append('</span>')
							πF.SetLineno(580)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_classifier.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 582: def visit_colspec(self, node):
					πF.SetLineno(582)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp044 = πg.NewFunction(πg.NewCode("visit_colspec", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 583: self.colspecs.append(node)
							πF.SetLineno(583)
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
							// line 585: node.parent.stubs.append(node.attributes.get('stub'))
							πF.SetLineno(585)
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
					if πE = πClass.SetItem(πF, ßvisit_colspec.ToObject(), πTemp044); πE != nil {
						continue
					}
					// line 587: def depart_colspec(self, node):
					πF.SetLineno(587)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp045 = πg.NewFunction(πg.NewCode("depart_colspec", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 589: if isinstance(node.next_node(descend=False, siblings=True),
							πF.SetLineno(589)
						Label1:
							// line 591: return
							πF.SetLineno(591)
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
							// line 592: if 'colwidths-auto' in node.parent.parent['classes'] or (
							πF.SetLineno(592)
						Label5:
							// line 595: return
							πF.SetLineno(595)
							πR = πg.None
							continue
							goto Label6
						Label6:
							// line 596: total_width = sum(node['colwidth'] for node in self.colspecs)
							πF.SetLineno(596)
							πTemp001 = πF.MakeArgs(1)
							πTemp013 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 596: total_width = sum(node['colwidth'] for node in self.colspecs)
										πF.SetLineno(596)
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
							// line 597: self.body.append(self.starttag(node, 'colgroup'))
							πF.SetLineno(597)
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
							// line 599: colwidth = int(node['colwidth'] * 100.0 / total_width + 0.5)
							πF.SetLineno(599)
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
							// line 600: self.body.append(self.emptytag(node, 'col',
							πF.SetLineno(600)
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
							if πTemp006, πE = πg.Mod(πF, πg.NewStr("width: %i%%").ToObject(), µcolwidth); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"style", πTemp006},
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
							// line 602: self.body.append('</colgroup>\n')
							πF.SetLineno(602)
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
					if πE = πClass.SetItem(πF, ßdepart_colspec.ToObject(), πTemp045); πE != nil {
						continue
					}
					// line 604: def visit_comment(self, node,
					πF.SetLineno(604)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("-(?=-)").ToObject()
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp048, πE = πg.GetAttr(πF, πTemp047, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp047, πE = πTemp048.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp048, πE = πg.GetAttr(πF, πTemp047, ßsub, nil); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "sub", Def: πTemp048}
					πTemp046 = πg.NewFunction(πg.NewCode("visit_comment", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µsub *πg.Object = πArgs[2]
						_ = µsub
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
							// line 606: """Escape double-dashes in comment text."""
							πF.SetLineno(606)
							// line 607: self.body.append('<!-- %s -->\n' % sub('- ', node.astext()))
							πF.SetLineno(607)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("- ").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp005
							if πE = πg.CheckLocal(πF, µsub, "sub"); πE != nil {
								continue
							}
							if πTemp004, πE = µsub.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<!-- %s -->\n").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 609: raise nodes.SkipNode
							πF.SetLineno(609)
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
					if πE = πClass.SetItem(πF, ßvisit_comment.ToObject(), πTemp046); πE != nil {
						continue
					}
					// line 606: """Escape double-dashes in comment text."""
					πF.SetLineno(606)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp047}, πg.NewStr("Escape double-dashes in comment text.").ToObject()); πE != nil {
						continue
					}
					if πTemp048, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_comment); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp048, ß__doc__, πTemp047); πE != nil {
						continue
					}
					// line 611: def visit_compound(self, node):
					πF.SetLineno(611)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp047 = πg.NewFunction(πg.NewCode("visit_compound", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
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
							// line 612: self.body.append(self.starttag(node, 'div', CLASS='compound'))
							πF.SetLineno(612)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßcompound.ToObject()},
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 613: if len(node) > 1:
							πF.SetLineno(613)
						Label1:
							// line 614: node[0]['classes'].append('compound-first')
							πF.SetLineno(614)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("compound-first").ToObject()
							πTemp004 = ßclasses.ToObject()
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 615: node[-1]['classes'].append('compound-last')
							πF.SetLineno(615)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("compound-last").ToObject()
							πTemp004 = ßclasses.ToObject()
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp008
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp007 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label5
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
								µchild = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 617: child['classes'].append('compound-middle')
							πF.SetLineno(617)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("compound-middle").ToObject()
							πTemp005 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µchild, πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
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
					if πE = πClass.SetItem(πF, ßvisit_compound.ToObject(), πTemp047); πE != nil {
						continue
					}
					// line 619: def depart_compound(self, node):
					πF.SetLineno(619)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp048 = πg.NewFunction(πg.NewCode("depart_compound", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 620: self.body.append('</div>\n')
							πF.SetLineno(620)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_compound.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 622: def visit_container(self, node):
					πF.SetLineno(622)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp049 = πg.NewFunction(πg.NewCode("visit_container", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 623: self.body.append(self.starttag(node, 'div', CLASS='docutils container'))
							πF.SetLineno(623)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("docutils container").ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_container.ToObject(), πTemp049); πE != nil {
						continue
					}
					// line 625: def depart_container(self, node):
					πF.SetLineno(625)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp050 = πg.NewFunction(πg.NewCode("depart_container", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 626: self.body.append('</div>\n')
							πF.SetLineno(626)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_container.ToObject(), πTemp050); πE != nil {
						continue
					}
					// line 628: def visit_contact(self, node):
					πF.SetLineno(628)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp051 = πg.NewFunction(πg.NewCode("visit_contact", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 629: self.visit_docinfo_item(node, 'contact', meta=False)
							πF.SetLineno(629)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßcontact.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_contact.ToObject(), πTemp051); πE != nil {
						continue
					}
					// line 631: def depart_contact(self, node):
					πF.SetLineno(631)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp052 = πg.NewFunction(πg.NewCode("depart_contact", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 632: self.depart_docinfo_item()
							πF.SetLineno(632)
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
					if πE = πClass.SetItem(πF, ßdepart_contact.ToObject(), πTemp052); πE != nil {
						continue
					}
					// line 634: def visit_copyright(self, node):
					πF.SetLineno(634)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp053 = πg.NewFunction(πg.NewCode("visit_copyright", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 635: self.visit_docinfo_item(node, 'copyright')
							πF.SetLineno(635)
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
					if πE = πClass.SetItem(πF, ßvisit_copyright.ToObject(), πTemp053); πE != nil {
						continue
					}
					// line 637: def depart_copyright(self, node):
					πF.SetLineno(637)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp054 = πg.NewFunction(πg.NewCode("depart_copyright", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 638: self.depart_docinfo_item()
							πF.SetLineno(638)
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
					if πE = πClass.SetItem(πF, ßdepart_copyright.ToObject(), πTemp054); πE != nil {
						continue
					}
					// line 640: def visit_date(self, node):
					πF.SetLineno(640)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp055 = πg.NewFunction(πg.NewCode("visit_date", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 641: self.visit_docinfo_item(node, 'date')
							πF.SetLineno(641)
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
					if πE = πClass.SetItem(πF, ßvisit_date.ToObject(), πTemp055); πE != nil {
						continue
					}
					// line 643: def depart_date(self, node):
					πF.SetLineno(643)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp056 = πg.NewFunction(πg.NewCode("depart_date", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 644: self.depart_docinfo_item()
							πF.SetLineno(644)
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
					if πE = πClass.SetItem(πF, ßdepart_date.ToObject(), πTemp056); πE != nil {
						continue
					}
					// line 646: def visit_decoration(self, node):
					πF.SetLineno(646)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp057 = πg.NewFunction(πg.NewCode("visit_decoration", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 647: pass
							πF.SetLineno(647)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_decoration.ToObject(), πTemp057); πE != nil {
						continue
					}
					// line 649: def depart_decoration(self, node):
					πF.SetLineno(649)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp058 = πg.NewFunction(πg.NewCode("depart_decoration", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 650: pass
							πF.SetLineno(650)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_decoration.ToObject(), πTemp058); πE != nil {
						continue
					}
					// line 652: def visit_definition(self, node):
					πF.SetLineno(652)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp059 = πg.NewFunction(πg.NewCode("visit_definition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 653: self.body.append('</dt>\n')
							πF.SetLineno(653)
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
							// line 654: self.body.append(self.starttag(node, 'dd', ''))
							πF.SetLineno(654)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_definition.ToObject(), πTemp059); πE != nil {
						continue
					}
					// line 656: def depart_definition(self, node):
					πF.SetLineno(656)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp060 = πg.NewFunction(πg.NewCode("depart_definition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 657: self.body.append('</dd>\n')
							πF.SetLineno(657)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dd>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_definition.ToObject(), πTemp060); πE != nil {
						continue
					}
					// line 659: def visit_definition_list(self, node):
					πF.SetLineno(659)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp061 = πg.NewFunction(πg.NewCode("visit_definition_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 660: classes = node.setdefault('classes', [])
							πF.SetLineno(660)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßclasses.ToObject()
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µclasses = πTemp004
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßis_compactable, nil); πE != nil {
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
							// line 661: if self.is_compactable(node):
							πF.SetLineno(661)
						Label1:
							// line 662: classes.append('simple')
							πF.SetLineno(662)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßsimple.ToObject()
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µclasses, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 663: self.body.append(self.starttag(node, 'dl'))
							πF.SetLineno(663)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdl.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_definition_list.ToObject(), πTemp061); πE != nil {
						continue
					}
					// line 665: def depart_definition_list(self, node):
					πF.SetLineno(665)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp062 = πg.NewFunction(πg.NewCode("depart_definition_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 666: self.body.append('</dl>\n')
							πF.SetLineno(666)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dl>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_definition_list.ToObject(), πTemp062); πE != nil {
						continue
					}
					// line 668: def visit_definition_list_item(self, node):
					πF.SetLineno(668)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp063 = πg.NewFunction(πg.NewCode("visit_definition_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 670: node.children[0]['classes'] = (
							πF.SetLineno(670)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßclasses.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßclasses.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßget, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, πTemp006, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 672: node.children[0]['ids'] = (
							πF.SetLineno(672)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßids.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßids.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßget, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							πTemp005 = ßids.ToObject()
							if πE = πg.SetItem(πF, πTemp006, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 674: node.children[0]['names'] = (
							πF.SetLineno(674)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßnames.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßnames.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßget, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							πTemp005 = ßnames.ToObject()
							if πE = πg.SetItem(πF, πTemp006, πTemp005, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_definition_list_item.ToObject(), πTemp063); πE != nil {
						continue
					}
					// line 677: def depart_definition_list_item(self, node):
					πF.SetLineno(677)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp064 = πg.NewFunction(πg.NewCode("depart_definition_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 678: pass
							πF.SetLineno(678)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_definition_list_item.ToObject(), πTemp064); πE != nil {
						continue
					}
					// line 680: def visit_description(self, node):
					πF.SetLineno(680)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp065 = πg.NewFunction(πg.NewCode("visit_description", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 681: self.body.append(self.starttag(node, 'dd', ''))
							πF.SetLineno(681)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdd.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_description.ToObject(), πTemp065); πE != nil {
						continue
					}
					// line 683: def depart_description(self, node):
					πF.SetLineno(683)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp066 = πg.NewFunction(πg.NewCode("depart_description", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 684: self.body.append('</dd>\n')
							πF.SetLineno(684)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dd>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_description.ToObject(), πTemp066); πE != nil {
						continue
					}
					// line 686: def visit_docinfo(self, node):
					πF.SetLineno(686)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp067 = πg.NewFunction(πg.NewCode("visit_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
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
							// line 687: self.context.append(len(self.body))
							πF.SetLineno(687)
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
							// line 688: classes = 'docinfo'
							πF.SetLineno(688)
							µclasses = ßdocinfo.ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßis_compactable, nil); πE != nil {
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
							// line 689: if (self.is_compactable(node)):
							πF.SetLineno(689)
						Label1:
							// line 690: classes += ' simple'
							πF.SetLineno(690)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µclasses, πg.NewStr(" simple").ToObject()); πE != nil {
								continue
							}
							µclasses = πTemp003
							goto Label2
						Label2:
							// line 691: self.body.append(self.starttag(node, 'dl', CLASS=classes))
							πF.SetLineno(691)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdl.ToObject()
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"CLASS", µclasses},
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
					if πE = πClass.SetItem(πF, ßvisit_docinfo.ToObject(), πTemp067); πE != nil {
						continue
					}
					// line 693: def depart_docinfo(self, node):
					πF.SetLineno(693)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp068 = πg.NewFunction(πg.NewCode("depart_docinfo", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 694: self.body.append('</dl>\n')
							πF.SetLineno(694)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dl>\n").ToObject()
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
							// line 695: start = self.context.pop()
							πF.SetLineno(695)
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
							// line 696: self.docinfo = self.body[start:]
							πF.SetLineno(696)
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
							// line 697: self.body = []
							πF.SetLineno(697)
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
					if πE = πClass.SetItem(πF, ßdepart_docinfo.ToObject(), πTemp068); πE != nil {
						continue
					}
					// line 699: def visit_docinfo_item(self, node, name, meta=True):
					πF.SetLineno(699)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp005[2] = πg.Param{Name: "name", Def: nil}
					if πTemp070, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp005[3] = πg.Param{Name: "meta", Def: πTemp070}
					πTemp069 = πg.NewFunction(πg.NewCode("visit_docinfo_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 700: if meta:
							πF.SetLineno(700)
						Label1:
							// line 701: meta_tag = '<meta name="%s" content="%s" />\n' \
							πF.SetLineno(701)
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
							// line 703: self.add_meta(meta_tag)
							πF.SetLineno(703)
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
							// line 704: self.body.append('<dt class="%s">%s</dt>\n'
							πF.SetLineno(704)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µname, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<dt class=\"%s\">%s</dt>\n").ToObject(), πTemp003); πE != nil {
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
							// line 706: self.body.append(self.starttag(node, 'dd', '', CLASS=name))
							πF.SetLineno(706)
							πTemp004 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp009[0] = µnode
							πTemp009[1] = ßdd.ToObject()
							πTemp009[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"CLASS", µname},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_docinfo_item.ToObject(), πTemp069); πE != nil {
						continue
					}
					// line 708: def depart_docinfo_item(self):
					πF.SetLineno(708)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp070 = πg.NewFunction(πg.NewCode("depart_docinfo_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 709: self.body.append('</dd>\n')
							πF.SetLineno(709)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dd>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_docinfo_item.ToObject(), πTemp070); πE != nil {
						continue
					}
					// line 711: def visit_doctest_block(self, node):
					πF.SetLineno(711)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp071 = πg.NewFunction(πg.NewCode("visit_doctest_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 712: self.body.append(self.starttag(node, 'pre', suffix='',
							πF.SetLineno(712)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßpre.ToObject()
							πTemp003 = πg.KWArgs{
								{"suffix", ß.ToObject()},
								{"CLASS", πg.NewStr("code python doctest").ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_doctest_block.ToObject(), πTemp071); πE != nil {
						continue
					}
					// line 715: def depart_doctest_block(self, node):
					πF.SetLineno(715)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp072 = πg.NewFunction(πg.NewCode("depart_doctest_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 716: self.body.append('\n</pre>\n')
							πF.SetLineno(716)
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
					if πE = πClass.SetItem(πF, ßdepart_doctest_block.ToObject(), πTemp072); πE != nil {
						continue
					}
					// line 718: def visit_document(self, node):
					πF.SetLineno(718)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp073 = πg.NewFunction(πg.NewCode("visit_document", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
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
							// line 719: title = (node.get('title', '') or os.path.basename(node['source'])
							πF.SetLineno(719)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßtitle.ToObject()
							πTemp003[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
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
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpath, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßbasename, nil); πE != nil {
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
							πTemp001 = πg.NewStr("docutils document without title").ToObject()
						Label1:
							µtitle = πTemp001
							// line 721: self.head.append('<title>%s</title>\n' % self.encode(title))
							πF.SetLineno(721)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp006[0] = µtitle
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<title>%s</title>\n").ToObject(), πTemp005); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_document.ToObject(), πTemp073); πE != nil {
						continue
					}
					// line 723: def depart_document(self, node):
					πF.SetLineno(723)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp074 = πg.NewFunction(πg.NewCode("depart_document", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 724: self.head_prefix.extend([self.doctype,
							πF.SetLineno(724)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdoctype, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhead_prefix_template, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßlanguage_code, nil); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßlang.ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp006 = πTemp005.ToObject()
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 727: self.html_prolog.append(self.doctype)
							πF.SetLineno(727)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdoctype, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_prolog, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 728: self.meta.insert(0, self.content_type % self.settings.output_encoding)
							πF.SetLineno(728)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent_type, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmeta, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 729: self.head.insert(0, self.content_type % self.settings.output_encoding)
							πF.SetLineno(729)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent_type, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmeta, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.Contains(πF, πTemp006, πg.NewStr("name=\"dcterms.").ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 730: if 'name="dcterms.' in ''.join(self.meta):
							πF.SetLineno(730)
						Label1:
							// line 731: self.head.append(
							πF.SetLineno(731)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<link rel=\"schema.dcterms\" href=\"http://purl.org/dc/terms/\">").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							goto Label4
							// line 733: if self.math_header:
							πF.SetLineno(733)
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, ßmathjax.ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							goto Label6
							// line 734: if self.math_output == 'mathjax':
							πF.SetLineno(734)
						Label5:
							// line 735: self.head.extend(self.math_header)
							πF.SetLineno(735)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label7
						Label6:
							// line 737: self.stylesheet.extend(self.math_header)
							πF.SetLineno(737)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstylesheet, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label7
						Label7:
							goto Label4
						Label4:
							// line 739: self.html_head.extend(self.head[1:])
							πF.SetLineno(739)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_head, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 740: self.body_prefix.append(self.starttag(node, 'div', CLASS='document'))
							πF.SetLineno(740)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp009 = πg.KWArgs{
								{"CLASS", ßdocument.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 741: self.body_suffix.insert(0, '</div>\n')
							πF.SetLineno(741)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = πg.NewStr("</div>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody_suffix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 742: self.fragment.extend(self.body) # self.fragment is the "naked" body
							πF.SetLineno(742)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfragment, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 743: self.html_body.extend(self.body_prefix[1:] + self.body_pre_docinfo
							πF.SetLineno(743)
							πTemp001 = πF.MakeArgs(1)
							if πTemp010, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßbody_pre_docinfo, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßdocinfo, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßbody_suffix, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_body, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 746: assert not self.context, 'len(context) = %s' % len(self.context)
							πF.SetLineno(746)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("len(context) = %s").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_document.ToObject(), πTemp074); πE != nil {
						continue
					}
					// line 748: def visit_emphasis(self, node):
					πF.SetLineno(748)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp075 = πg.NewFunction(πg.NewCode("visit_emphasis", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 749: self.body.append(self.starttag(node, 'em', ''))
							πF.SetLineno(749)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßem.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_emphasis.ToObject(), πTemp075); πE != nil {
						continue
					}
					// line 751: def depart_emphasis(self, node):
					πF.SetLineno(751)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp076 = πg.NewFunction(πg.NewCode("depart_emphasis", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 752: self.body.append('</em>')
							πF.SetLineno(752)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</em>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_emphasis.ToObject(), πTemp076); πE != nil {
						continue
					}
					// line 754: def visit_entry(self, node):
					πF.SetLineno(754)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp077 = πg.NewFunction(πg.NewCode("visit_entry", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µtagname *πg.Object = πg.UnboundLocal
						_ = µtagname
						var πTemp001 *πg.Dict
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
							// line 755: atts = {'class': []}
							πF.SetLineno(755)
							πTemp001 = πg.NewDict()
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp003 = πTemp001.ToObject()
							µatts = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßparent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßthead, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 756: if isinstance(node.parent.parent, nodes.thead):
							πF.SetLineno(756)
						Label1:
							// line 757: atts['class'].append('head')
							πF.SetLineno(757)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßhead.ToObject()
							πTemp003 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßcolumn, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßparent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßstubs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 758: if node.parent.parent.parent.stubs[node.parent.column]:
							πF.SetLineno(758)
						Label3:
							// line 760: atts['class'].append('stub')
							πF.SetLineno(760)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßstub.ToObject()
							πTemp003 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label4
						Label4:
							πTemp003 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 761: if atts['class']:
							πF.SetLineno(761)
						Label5:
							// line 762: tagname = 'th'
							πF.SetLineno(762)
							µtagname = ßth.ToObject()
							// line 763: atts['class'] = ' '.join(atts['class'])
							πF.SetLineno(763)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp006 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp003); πE != nil {
								continue
							}
							goto Label7
						Label6:
							// line 765: tagname = 'td'
							πF.SetLineno(765)
							µtagname = ßtd.ToObject()
							// line 766: del atts['class']
							πF.SetLineno(766)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp003 = ßclass.ToObject()
							if πE = πg.DelItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 767: node.parent.column += 1
							πF.SetLineno(767)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcolumn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßcolumn, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßmorerows.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							goto Label9
							// line 768: if 'morerows' in node:
							πF.SetLineno(768)
						Label8:
							// line 769: atts['rowspan'] = node['morerows'] + 1
							πF.SetLineno(769)
							πTemp004 = ßmorerows.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp006 = ßrowspan.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp004); πE != nil {
								continue
							}
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßmorecols.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 770: if 'morecols' in node:
							πF.SetLineno(770)
						Label10:
							// line 771: atts['colspan'] = node['morecols'] + 1
							πF.SetLineno(771)
							πTemp004 = ßmorecols.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp006 = ßcolspan.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp004); πE != nil {
								continue
							}
							// line 772: node.parent.column += node['morecols']
							πF.SetLineno(772)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcolumn, nil); πE != nil {
								continue
							}
							πTemp003 = ßmorecols.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp007, ßcolumn, πTemp003); πE != nil {
								continue
							}
							goto Label11
						Label11:
							// line 773: self.body.append(self.starttag(node, tagname, '', **atts))
							πF.SetLineno(773)
							πTemp002 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008[0] = µnode
							if πE = πg.CheckLocal(πF, µtagname, "tagname"); πE != nil {
								continue
							}
							πTemp008[1] = µtagname
							πTemp008[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp008, nil, nil, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002[0] = πTemp004
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
							// line 774: self.context.append('</%s>\n' % tagname.lower())
							πF.SetLineno(774)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtagname, "tagname"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtagname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("</%s>\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_entry.ToObject(), πTemp077); πE != nil {
						continue
					}
					// line 779: def depart_entry(self, node):
					πF.SetLineno(779)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp078 = πg.NewFunction(πg.NewCode("depart_entry", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 780: self.body.append(self.context.pop())
							πF.SetLineno(780)
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
					if πE = πClass.SetItem(πF, ßdepart_entry.ToObject(), πTemp078); πE != nil {
						continue
					}
					// line 782: def visit_enumerated_list(self, node):
					πF.SetLineno(782)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp079 = πg.NewFunction(πg.NewCode("visit_enumerated_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 783: atts = {}
							πF.SetLineno(783)
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
							// line 784: if 'start' in node:
							πF.SetLineno(784)
						Label1:
							// line 785: atts['start'] = node['start']
							πF.SetLineno(785)
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
							// line 786: if 'enumtype' in node:
							πF.SetLineno(786)
						Label3:
							// line 787: atts['class'] = node['enumtype']
							πF.SetLineno(787)
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
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 788: if self.is_compactable(node):
							πF.SetLineno(788)
						Label5:
							// line 789: atts['class'] = (atts.get('class', '') + ' simple').strip()
							πF.SetLineno(789)
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
							goto Label6
						Label6:
							// line 790: self.body.append(self.starttag(node, 'ol', **atts))
							πF.SetLineno(790)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							πTemp007[1] = ßol.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp007, nil, nil, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßvisit_enumerated_list.ToObject(), πTemp079); πE != nil {
						continue
					}
					// line 792: def depart_enumerated_list(self, node):
					πF.SetLineno(792)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp080 = πg.NewFunction(πg.NewCode("depart_enumerated_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 793: self.body.append('</ol>\n')
							πF.SetLineno(793)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</ol>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_enumerated_list.ToObject(), πTemp080); πE != nil {
						continue
					}
					// line 795: def visit_field_list(self, node):
					πF.SetLineno(795)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp081 = πg.NewFunction(πg.NewCode("visit_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
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
							// line 798: classes = 'field-list'
							πF.SetLineno(798)
							µclasses = πg.NewStr("field-list").ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßis_compactable, nil); πE != nil {
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
							// line 799: if (self.is_compactable(node)):
							πF.SetLineno(799)
						Label1:
							// line 800: classes += ' simple'
							πF.SetLineno(800)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µclasses, πg.NewStr(" simple").ToObject()); πE != nil {
								continue
							}
							µclasses = πTemp002
							goto Label2
						Label2:
							// line 801: self.body.append(self.starttag(node, 'dl', CLASS=classes))
							πF.SetLineno(801)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßdl.ToObject()
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"CLASS", µclasses},
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field_list.ToObject(), πTemp081); πE != nil {
						continue
					}
					// line 803: def depart_field_list(self, node):
					πF.SetLineno(803)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp082 = πg.NewFunction(πg.NewCode("depart_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 804: self.body.append('</dl>\n')
							πF.SetLineno(804)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dl>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_field_list.ToObject(), πTemp082); πE != nil {
						continue
					}
					// line 806: def visit_field(self, node):
					πF.SetLineno(806)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp083 = πg.NewFunction(πg.NewCode("visit_field", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 807: pass
							πF.SetLineno(807)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field.ToObject(), πTemp083); πE != nil {
						continue
					}
					// line 809: def depart_field(self, node):
					πF.SetLineno(809)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp084 = πg.NewFunction(πg.NewCode("depart_field", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 810: pass
							πF.SetLineno(810)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field.ToObject(), πTemp084); πE != nil {
						continue
					}
					// line 814: def visit_field_name(self, node):
					πF.SetLineno(814)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp085 = πg.NewFunction(πg.NewCode("visit_field_name", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
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
							// line 815: self.body.append(self.starttag(node, 'dt', '',
							πF.SetLineno(815)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdt.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp007 = πg.KWArgs{
								{"CLASS", πTemp005},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp007); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_field_name.ToObject(), πTemp085); πE != nil {
						continue
					}
					// line 818: def depart_field_name(self, node):
					πF.SetLineno(818)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp086 = πg.NewFunction(πg.NewCode("depart_field_name", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 819: self.body.append('</dt>\n')
							πF.SetLineno(819)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_field_name.ToObject(), πTemp086); πE != nil {
						continue
					}
					// line 821: def visit_field_body(self, node):
					πF.SetLineno(821)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp087 = πg.NewFunction(πg.NewCode("visit_field_body", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 822: self.body.append(self.starttag(node, 'dd', '',
							πF.SetLineno(822)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdd.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp007 = πg.KWArgs{
								{"CLASS", πTemp005},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp007); πE != nil {
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 825: if not node.children:
							πF.SetLineno(825)
						Label1:
							// line 826: self.body.append('<p></p>')
							πF.SetLineno(826)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<p></p>").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_field_body.ToObject(), πTemp087); πE != nil {
						continue
					}
					// line 828: def depart_field_body(self, node):
					πF.SetLineno(828)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp088 = πg.NewFunction(πg.NewCode("depart_field_body", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 829: self.body.append('</dd>\n')
							πF.SetLineno(829)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dd>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_field_body.ToObject(), πTemp088); πE != nil {
						continue
					}
					// line 831: def visit_figure(self, node):
					πF.SetLineno(831)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp089 = πg.NewFunction(πg.NewCode("visit_figure", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 832: atts = {'class': 'figure'}
							πF.SetLineno(832)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßclass.ToObject(), ßfigure.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßwidth.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
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
							// line 833: if node.get('width'):
							πF.SetLineno(833)
						Label1:
							// line 834: atts['style'] = 'width: %s' % node['width']
							πF.SetLineno(834)
							πTemp004 = ßwidth.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("width: %s").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp006 = ßstyle.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp004); πE != nil {
								continue
							}
							goto Label2
						Label2:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
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
							// line 835: if node.get('align'):
							πF.SetLineno(835)
						Label3:
							// line 836: atts['class'] += " align-" + node['align']
							πF.SetLineno(836)
							πTemp002 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp002); πE != nil {
								continue
							}
							πTemp006 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr(" align-").ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IAdd(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp007 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp007, πTemp006); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 837: self.body.append(self.starttag(node, 'div', **atts))
							πF.SetLineno(837)
							πTemp003 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008[0] = µnode
							πTemp008[1] = ßdiv.ToObject()
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
							πTemp003[0] = πTemp004
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
					if πE = πClass.SetItem(πF, ßvisit_figure.ToObject(), πTemp089); πE != nil {
						continue
					}
					// line 839: def depart_figure(self, node):
					πF.SetLineno(839)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp090 = πg.NewFunction(πg.NewCode("depart_figure", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 840: self.body.append('</div>\n')
							πF.SetLineno(840)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_figure.ToObject(), πTemp090); πE != nil {
						continue
					}
					// line 843: def visit_footer(self, node):
					πF.SetLineno(843)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp091 = πg.NewFunction(πg.NewCode("visit_footer", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 844: self.context.append(len(self.body))
							πF.SetLineno(844)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_footer.ToObject(), πTemp091); πE != nil {
						continue
					}
					// line 846: def depart_footer(self, node):
					πF.SetLineno(846)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp092 = πg.NewFunction(πg.NewCode("depart_footer", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var µfooter *πg.Object = πg.UnboundLocal
						_ = µfooter
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 847: start = self.context.pop()
							πF.SetLineno(847)
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
							// line 848: footer = [self.starttag(node, 'div', CLASS='footer'),
							πF.SetLineno(848)
							πTemp003 = make([]*πg.Object, 2)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp004[1] = ßdiv.ToObject()
							πTemp005 = πg.KWArgs{
								{"CLASS", ßfooter.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("<hr class=\"footer\" />\n").ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µfooter = πTemp001
							// line 850: footer.extend(self.body[start:])
							πF.SetLineno(850)
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
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfooter, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 851: footer.append('\n</div>\n')
							πF.SetLineno(851)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\n</div>\n").ToObject()
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfooter, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 852: self.footer.extend(footer)
							πF.SetLineno(852)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							πTemp003[0] = µfooter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfooter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 853: self.body_suffix[:0] = footer
							πF.SetLineno(853)
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfooter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody_suffix, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp002, πTemp006, πTemp001); πE != nil {
								continue
							}
							// line 854: del self.body[start:]
							πF.SetLineno(854)
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
					if πE = πClass.SetItem(πF, ßdepart_footer.ToObject(), πTemp092); πE != nil {
						continue
					}
					// line 861: def visit_footnote(self, node):
					πF.SetLineno(861)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp093 = πg.NewFunction(πg.NewCode("visit_footnote", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_footnote_list, nil); πE != nil {
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
							// line 862: if not self.in_footnote_list:
							πF.SetLineno(862)
						Label1:
							// line 863: classes = 'footnote ' + self.settings.footnote_references
							πF.SetLineno(863)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßfootnote_references, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πg.NewStr("footnote ").ToObject(), πTemp004); πE != nil {
								continue
							}
							µclasses = πTemp001
							// line 864: self.body.append('<dl class="%s">\n'%classes)
							πF.SetLineno(864)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<dl class=\"%s\">\n").ToObject(), µclasses); πE != nil {
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
							// line 865: self.in_footnote_list = True
							πF.SetLineno(865)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_footnote_list, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_footnote.ToObject(), πTemp093); πE != nil {
						continue
					}
					// line 867: def depart_footnote(self, node):
					πF.SetLineno(867)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp094 = πg.NewFunction(πg.NewCode("depart_footnote", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 868: self.body.append('</dd>\n')
							πF.SetLineno(868)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dd>\n").ToObject()
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
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"descend", πTemp003},
								{"siblings", πTemp004},
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßnext_node, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfootnote, nil); πE != nil {
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 869: if not isinstance(node.next_node(descend=False, siblings=True),
							πF.SetLineno(869)
						Label1:
							// line 871: self.body.append('</dl>\n')
							πF.SetLineno(871)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dl>\n").ToObject()
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
							// line 872: self.in_footnote_list = False
							πF.SetLineno(872)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_footnote_list, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_footnote.ToObject(), πTemp094); πE != nil {
						continue
					}
					// line 874: def visit_footnote_reference(self, node):
					πF.SetLineno(874)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp095 = πg.NewFunction(πg.NewCode("visit_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µhref *πg.Object = πg.UnboundLocal
						_ = µhref
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 875: href = '#' + node['refid']
							πF.SetLineno(875)
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
							// line 876: classes = 'footnote-reference ' + self.settings.footnote_references
							πF.SetLineno(876)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfootnote_references, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πg.NewStr("footnote-reference ").ToObject(), πTemp003); πE != nil {
								continue
							}
							µclasses = πTemp001
							// line 877: self.body.append(self.starttag(node, 'a', '', #suffix,
							πF.SetLineno(877)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßa.ToObject()
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhref, "href"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"CLASS", µclasses},
								{"href", µhref},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßvisit_footnote_reference.ToObject(), πTemp095); πE != nil {
						continue
					}
					// line 880: def depart_footnote_reference(self, node):
					πF.SetLineno(880)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp096 = πg.NewFunction(πg.NewCode("depart_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 881: self.body.append('</a>')
							πF.SetLineno(881)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</a>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_footnote_reference.ToObject(), πTemp096); πE != nil {
						continue
					}
					// line 884: def visit_generated(self, node):
					πF.SetLineno(884)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp097 = πg.NewFunction(πg.NewCode("visit_generated", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µsectnum *πg.Object = πg.UnboundLocal
						_ = µsectnum
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßsectnum.ToObject()); πE != nil {
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
							// line 885: if 'sectnum' in node['classes']:
							πF.SetLineno(885)
						Label1:
							// line 887: sectnum = node.astext().rstrip(u' ')
							πF.SetLineno(887)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewUnicode("\xc2\xa0").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µsectnum = πTemp002
							// line 888: self.body.append('<span class="sectnum">%s</span> '
							πF.SetLineno(888)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsectnum, "sectnum"); πE != nil {
								continue
							}
							πTemp006[0] = µsectnum
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<span class=\"sectnum\">%s</span> ").ToObject(), πTemp003); πE != nil {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 891: raise nodes.SkipNode
							πF.SetLineno(891)
							πE = πF.Raise(πTemp002, nil, nil)
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
					if πE = πClass.SetItem(πF, ßvisit_generated.ToObject(), πTemp097); πE != nil {
						continue
					}
					// line 893: def depart_generated(self, node):
					πF.SetLineno(893)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp098 = πg.NewFunction(πg.NewCode("depart_generated", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 894: pass
							πF.SetLineno(894)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_generated.ToObject(), πTemp098); πE != nil {
						continue
					}
					// line 896: def visit_header(self, node):
					πF.SetLineno(896)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp099 = πg.NewFunction(πg.NewCode("visit_header", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 897: self.context.append(len(self.body))
							πF.SetLineno(897)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_header.ToObject(), πTemp099); πE != nil {
						continue
					}
					// line 899: def depart_header(self, node):
					πF.SetLineno(899)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp100 = πg.NewFunction(πg.NewCode("depart_header", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 900: start = self.context.pop()
							πF.SetLineno(900)
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
							// line 901: header = [self.starttag(node, 'div', CLASS='header')]
							πF.SetLineno(901)
							πTemp003 = make([]*πg.Object, 1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp004[1] = ßdiv.ToObject()
							πTemp005 = πg.KWArgs{
								{"CLASS", ßheader.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µheader = πTemp001
							// line 902: header.extend(self.body[start:])
							πF.SetLineno(902)
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
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µheader, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 903: header.append('\n<hr class="header"/>\n</div>\n')
							πF.SetLineno(903)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\n<hr class=\"header\"/>\n</div>\n").ToObject()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µheader, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 904: self.body_prefix.extend(header)
							πF.SetLineno(904)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp003[0] = µheader
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 905: self.header.extend(header)
							πF.SetLineno(905)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp003[0] = µheader
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßheader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 906: del self.body[start:]
							πF.SetLineno(906)
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
					if πE = πClass.SetItem(πF, ßdepart_header.ToObject(), πTemp100); πE != nil {
						continue
					}
					// line 909: object_image_types = {'.swf': 'application/x-shockwave-flash'}
					πF.SetLineno(909)
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, πg.NewStr(".swf").ToObject(), πg.NewStr("application/x-shockwave-flash").ToObject()); πE != nil {
						continue
					}
					πTemp101 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßobject_image_types.ToObject(), πTemp101); πE != nil {
						continue
					}
					// line 911: def visit_image(self, node):
					πF.SetLineno(911)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp101 = πg.NewFunction(πg.NewCode("visit_image", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µuri *πg.Object = πg.UnboundLocal
						_ = µuri
						var µext *πg.Object = πg.UnboundLocal
						_ = µext
						var µimagepath *πg.Object = πg.UnboundLocal
						_ = µimagepath
						var µimg *πg.Object = πg.UnboundLocal
						_ = µimg
						var µatt_name *πg.Object = πg.UnboundLocal
						_ = µatt_name
						var µmatch *πg.Object = πg.UnboundLocal
						_ = µmatch
						var µstyle *πg.Object = πg.UnboundLocal
						_ = µstyle
						var µsuffix *πg.Object = πg.UnboundLocal
						_ = µsuffix
						var πTemp001 *πg.Dict
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 *πg.Object
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 26:
								goto Label26
							case 27:
								goto Label27
							case 21:
								goto Label21
							case 22:
								goto Label22
							case 15:
								goto Label15
							default:
								panic("unexpected function state")
							}
							// line 912: atts = {}
							πF.SetLineno(912)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							// line 913: uri = node['uri']
							πF.SetLineno(913)
							πTemp002 = ßuri.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							µuri = πTemp003
							// line 914: ext = os.path.splitext(uri)[1].lower()
							πF.SetLineno(914)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							πTemp004[0] = µuri
							if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpath, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsplitext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µext = πTemp003
							if πE = πg.CheckLocal(πF, µext, "ext"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßobject_image_types, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp003, µext); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 915: if ext in self.object_image_types:
							πF.SetLineno(915)
						Label1:
							// line 916: atts['data'] = uri
							πF.SetLineno(916)
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µuri); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp003 = ßdata.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 917: atts['type'] = self.object_image_types[ext]
							πF.SetLineno(917)
							if πE = πg.CheckLocal(πF, µext, "ext"); πE != nil {
								continue
							}
							πTemp002 = µext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßobject_image_types, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßtype.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 919: atts['src'] = uri
							πF.SetLineno(919)
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µuri); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp003 = ßsrc.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 920: atts['alt'] = node.get('alt', uri)
							πF.SetLineno(920)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßalt.ToObject()
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							πTemp004[1] = µuri
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßalt.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnode, ßwidth.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 922: if 'width' in node:
							πF.SetLineno(922)
						Label4:
							// line 923: atts['width'] = node['width']
							πF.SetLineno(923)
							πTemp002 = ßwidth.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßwidth.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnode, ßheight.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							goto Label7
							// line 924: if 'height' in node:
							πF.SetLineno(924)
						Label6:
							// line 925: atts['height'] = node['height']
							πF.SetLineno(925)
							πTemp002 = ßheight.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßheight.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnode, ßscale.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 926: if 'scale' in node:
							πF.SetLineno(926)
						Label8:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßPIL); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µnode, ßwidth.ToObject()); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp009).ToObject()
							πTemp005 = πTemp006
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µnode, ßheight.ToObject()); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp009).ToObject()
							πTemp005 = πTemp006
						Label11:
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							πTemp002 = πTemp003
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßfile_insertion_enabled, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
						Label10:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							goto Label13
							// line 927: if (PIL and not ('width' in node and 'height' in node)
							πF.SetLineno(927)
						Label12:
							// line 929: imagepath = url2pathname(uri)
							πF.SetLineno(929)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							πTemp004[0] = µuri
							if πTemp002, πE = πg.ResolveGlobal(πF, ßurl2pathname); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µimagepath = πTemp003
							// line 930: try:
							πF.SetLineno(930)
							πF.PushCheckpoint(15)
							// line 931: img = PIL.Image.open(
							πF.SetLineno(931)
							πTemp004 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetfilesystemencoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							if πE = πg.CheckLocal(πF, µimagepath, "imagepath"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µimagepath, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPIL); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßImage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µimg = πTemp003
							πF.PopCheckpoint()
							// line 936: self.settings.record_dependencies.add(
							πF.SetLineno(936)
							πTemp004 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = πg.NewStr("\\").ToObject()
							πTemp010[1] = πg.NewStr("/").ToObject()
							if πE = πg.CheckLocal(πF, µimagepath, "imagepath"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µimagepath, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µatts, ßwidth.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label16
							}
							goto Label17
							// line 938: if 'width' not in atts:
							πF.SetLineno(938)
						Label16:
							// line 939: atts['width'] = '%dpx' % img.size[0]
							πF.SetLineno(939)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µimg, "img"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µimg, ßsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%dpx").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßwidth.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label17
						Label17:
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µatts, ßheight.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label18
							}
							goto Label19
							// line 940: if 'height' not in atts:
							πF.SetLineno(940)
						Label18:
							// line 941: atts['height'] = '%dpx' % img.size[1]
							πF.SetLineno(941)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µimg, "img"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µimg, ßsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%dpx").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßheight.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label19
						Label19:
							// line 942: del img
							πF.SetLineno(942)
							if πE = πg.CheckLocal(πF, µimg, "img"); πE != nil {
								continue
							}
							µimg = πg.UnboundLocal
							goto Label14
						Label15:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
							if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label20
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 933: except (IOError, UnicodeEncodeError):
							πF.SetLineno(933)
						Label20:
							// line 934: pass # TODO: warn?
							πF.SetLineno(934)
							πF.RestoreExc(nil, nil)
							goto Label14
						Label14:
							goto Label13
						Label13:
							πTemp003 = πg.NewTuple2(ßwidth.ToObject(), ßheight.ToObject()).ToObject()
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(22)
							πTemp007 = false
						Label21:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label23
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
								µatt_name = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(21)
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µatts, µatt_name); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label24
							}
							goto Label25
							// line 944: if att_name in atts:
							πF.SetLineno(944)
						Label24:
							// line 945: match = re.match(r'([0-9.]+)(\S*)$', atts[att_name])
							πF.SetLineno(945)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("([0-9.]+)(\\S*)$").ToObject()
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp003 = µatt_name
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmatch = πTemp003
							// line 946: assert match
							πF.SetLineno(946)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, µmatch, nil); πE != nil {
								continue
							}
							// line 947: atts[att_name] = '%s%s' % (
							πF.SetLineno(947)
							πTemp004 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							πTemp010[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp004[0] = πTemp014
							if πTemp013, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							πTemp015 = ßscale.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetItem(πF, µnode, πTemp015); πE != nil {
								continue
							}
							πTemp004[0] = πTemp016
							if πTemp015, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp015.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp013, πE = πg.Div(πF, πTemp016, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, πTemp014, πTemp013); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp005 = πg.NewTuple2(πTemp006, πTemp014).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s%s").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp006 = µatt_name
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp005); πE != nil {
								continue
							}
							goto Label25
						Label25:
							continue
						Label22:
							if πE != nil || πR != nil {
								continue
							}
						Label23:
							goto Label9
						Label9:
							// line 950: style = []
							πF.SetLineno(950)
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							µstyle = πTemp002
							πTemp003 = πg.NewTuple2(ßwidth.ToObject(), ßheight.ToObject()).ToObject()
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(27)
							πTemp007 = false
						Label26:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label28
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
								µatt_name = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(26)
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µatts, µatt_name); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label29
							}
							goto Label30
							// line 952: if att_name in atts:
							πF.SetLineno(952)
						Label29:
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("^[0-9.]+$").ToObject()
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp003 = µatt_name
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label31
							}
							goto Label32
							// line 953: if re.match(r'^[0-9.]+$', atts[att_name]):
							πF.SetLineno(953)
						Label31:
							// line 955: atts[att_name] += 'px'
							πF.SetLineno(955)
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp003 = µatt_name
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp005, ßpx.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp006 = µatt_name
							if πE = πg.SetItem(πF, µatts, πTemp006, πTemp003); πE != nil {
								continue
							}
							goto Label32
						Label32:
							// line 956: style.append('%s: %s;' % (att_name, atts[att_name]))
							πF.SetLineno(956)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp006 = µatt_name
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µatts, πTemp006); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µatt_name, πTemp013).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s: %s;").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstyle, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 957: del atts[att_name]
							πF.SetLineno(957)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt_name, "att_name"); πE != nil {
								continue
							}
							πTemp003 = µatt_name
							if πE = πg.DelItem(πF, µatts, πTemp003); πE != nil {
								continue
							}
							goto Label30
						Label30:
							continue
						Label27:
							if πE != nil || πR != nil {
								continue
							}
						Label28:
							if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µstyle); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label33
							}
							goto Label34
							// line 958: if style:
							πF.SetLineno(958)
						Label33:
							// line 959: atts['style'] = ' '.join(style)
							πF.SetLineno(959)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
								continue
							}
							πTemp004[0] = µstyle
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßstyle.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label34
						Label34:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πTemp005
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label35
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßreference, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πTemp006
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label36
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp006, ßparent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp013
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp006, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp013
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp009, πE = πg.IsTrue(πF, πTemp013); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp009).ToObject()
							πTemp003 = πTemp005
						Label36:
							πTemp002 = πTemp003
						Label35:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label37
							}
							goto Label38
							// line 960: if (isinstance(node.parent, nodes.TextElement) or
							πF.SetLineno(960)
						Label37:
							// line 964: suffix = ''
							πF.SetLineno(964)
							µsuffix = ß.ToObject()
							goto Label39
						Label38:
							// line 966: suffix = '\n'
							πF.SetLineno(966)
							µsuffix = πg.NewStr("\n").ToObject()
							goto Label39
						Label39:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnode, ßalign.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label40
							}
							goto Label41
							// line 967: if 'align' in node:
							πF.SetLineno(967)
						Label40:
							// line 968: atts['class'] = 'align-%s' % node['align']
							πF.SetLineno(968)
							πTemp003 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("align-%s").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label41
						Label41:
							if πE = πg.CheckLocal(πF, µext, "ext"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßobject_image_types, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp003, µext); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label42
							}
							goto Label43
							// line 969: if ext in self.object_image_types:
							πF.SetLineno(969)
						Label42:
							// line 971: self.body.append(self.starttag(node, 'object', suffix, **atts) +
							πF.SetLineno(971)
							πTemp004 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp010[0] = µnode
							πTemp010[1] = ßobject.ToObject()
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							πTemp010[2] = µsuffix
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Invoke(πF, πTemp006, πTemp010, nil, nil, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = ßalt.ToObject()
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							πTemp010[1] = µuri
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp005, πE = πg.Add(πF, πTemp013, πTemp014); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πg.NewStr("</object>").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µsuffix); πE != nil {
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
							goto Label44
						Label43:
							// line 974: self.body.append(self.emptytag(node, 'img', suffix, **atts))
							πF.SetLineno(974)
							πTemp004 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp010[0] = µnode
							πTemp010[1] = ßimg.ToObject()
							if πE = πg.CheckLocal(πF, µsuffix, "suffix"); πE != nil {
								continue
							}
							πTemp010[2] = µsuffix
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßemptytag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp010, nil, nil, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
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
							goto Label44
						Label44:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_image.ToObject(), πTemp101); πE != nil {
						continue
					}
					// line 976: def depart_image(self, node):
					πF.SetLineno(976)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp102 = πg.NewFunction(πg.NewCode("depart_image", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 977: pass
							πF.SetLineno(977)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_image.ToObject(), πTemp102); πE != nil {
						continue
					}
					// line 979: def visit_inline(self, node):
					πF.SetLineno(979)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp103 = πg.NewFunction(πg.NewCode("visit_inline", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 980: self.body.append(self.starttag(node, 'span', ''))
							πF.SetLineno(980)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßspan.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_inline.ToObject(), πTemp103); πE != nil {
						continue
					}
					// line 982: def depart_inline(self, node):
					πF.SetLineno(982)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp104 = πg.NewFunction(πg.NewCode("depart_inline", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 983: self.body.append('</span>')
							πF.SetLineno(983)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_inline.ToObject(), πTemp104); πE != nil {
						continue
					}
					// line 986: def visit_label(self, node):
					πF.SetLineno(986)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp105 = πg.NewFunction(πg.NewCode("visit_label", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µbackrefs *πg.Object = πg.UnboundLocal
						_ = µbackrefs
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
						var πR *πg.Object
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfootnote, nil); πE != nil {
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
							// line 987: if (isinstance(node.parent, nodes.footnote)):
							πF.SetLineno(987)
						Label1:
							// line 988: classes = self.settings.footnote_references
							πF.SetLineno(988)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfootnote_references, nil); πE != nil {
								continue
							}
							µclasses = πTemp003
							goto Label3
						Label2:
							// line 990: classes = 'brackets'
							πF.SetLineno(990)
							µclasses = ßbrackets.ToObject()
							goto Label3
						Label3:
							// line 992: self.body.append(self.starttag(node.parent, 'dt', '', CLASS='label'))
							πF.SetLineno(992)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = ßdt.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßlabel.ToObject()},
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
							// line 993: self.body.append(self.starttag(node, 'span', '', CLASS=classes))
							πF.SetLineno(993)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßspan.ToObject()
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"CLASS", µclasses},
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfootnote_backlinks, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 995: if self.settings.footnote_backlinks:
							πF.SetLineno(995)
						Label4:
							// line 996: backrefs = node.parent['backrefs']
							πF.SetLineno(996)
							πTemp002 = ßbackrefs.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µbackrefs = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp001[0] = µbackrefs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 997: if len(backrefs) == 1:
							πF.SetLineno(997)
						Label6:
							// line 998: self.body.append('<a class="fn-backref" href="#%s">'
							πF.SetLineno(998)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µbackrefs, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<a class=\"fn-backref\" href=\"#%s\">").ToObject(), πTemp007); πE != nil {
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
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ßvisit_label.ToObject(), πTemp105); πE != nil {
						continue
					}
					// line 1001: def depart_label(self, node):
					πF.SetLineno(1001)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp106 = πg.NewFunction(πg.NewCode("depart_label", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µbackrefs *πg.Object = πg.UnboundLocal
						_ = µbackrefs
						var µbacklinks *πg.Object = πg.UnboundLocal
						_ = µbacklinks
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []πg.Param
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfootnote_backlinks, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1002: if self.settings.footnote_backlinks:
							πF.SetLineno(1002)
						Label1:
							// line 1003: backrefs = node.parent['backrefs']
							πF.SetLineno(1003)
							πTemp001 = ßbackrefs.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µbackrefs = πTemp002
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp005[0] = µbackrefs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1004: if len(backrefs) == 1:
							πF.SetLineno(1004)
						Label3:
							// line 1005: self.body.append('</a>')
							πF.SetLineno(1005)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("</a>").ToObject()
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
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 1006: self.body.append('</span>')
							πF.SetLineno(1006)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("</span>").ToObject()
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßfootnote_backlinks, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label5
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
								continue
							}
							πTemp005[0] = µbackrefs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label5:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 1007: if self.settings.footnote_backlinks and len(backrefs) > 1:
							πF.SetLineno(1007)
						Label6:
							// line 1008: backlinks = ['<a href="#%s">%s</a>' % (ref, i)
							πF.SetLineno(1008)
							πTemp007 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
								var µref *πg.Object = πg.UnboundLocal
								_ = µref
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
										πTemp002 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µbackrefs, "backrefs"); πE != nil {
											continue
										}
										πTemp002[0] = µbackrefs
										πTemp002[1] = πg.NewInt(1).ToObject()
										if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
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
											if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
												continue
											}
											µi = πTemp004
											µref = πTemp007
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1008: backlinks = ['<a href="#%s">%s</a>' % (ref, i)
										πF.SetLineno(1008)
										if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp004 = πg.NewTuple2(µref, µi).ToObject()
										if πTemp003, πE = πg.Mod(πF, πg.NewStr("<a href=\"#%s\">%s</a>").ToObject(), πTemp004); πE != nil {
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µbacklinks = πTemp001
							// line 1010: self.body.append('<span class="fn-backref">(%s)</span>'
							πF.SetLineno(1010)
							πTemp005 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbacklinks, "backlinks"); πE != nil {
								continue
							}
							πTemp008[0] = µbacklinks
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<span class=\"fn-backref\">(%s)</span>").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label7:
							// line 1012: self.body.append('</dt>\n<dd>')
							πF.SetLineno(1012)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("</dt>\n<dd>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_label.ToObject(), πTemp106); πE != nil {
						continue
					}
					// line 1014: def visit_legend(self, node):
					πF.SetLineno(1014)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp107 = πg.NewFunction(πg.NewCode("visit_legend", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1015: self.body.append(self.starttag(node, 'div', CLASS='legend'))
							πF.SetLineno(1015)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßlegend.ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_legend.ToObject(), πTemp107); πE != nil {
						continue
					}
					// line 1017: def depart_legend(self, node):
					πF.SetLineno(1017)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp108 = πg.NewFunction(πg.NewCode("depart_legend", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1018: self.body.append('</div>\n')
							πF.SetLineno(1018)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_legend.ToObject(), πTemp108); πE != nil {
						continue
					}
					// line 1020: def visit_line(self, node):
					πF.SetLineno(1020)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp109 = πg.NewFunction(πg.NewCode("visit_line", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
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
							// line 1021: self.body.append(self.starttag(node, 'div', suffix='', CLASS='line'))
							πF.SetLineno(1021)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"suffix", ß.ToObject()},
								{"CLASS", ßline.ToObject()},
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 1022: if not len(node):
							πF.SetLineno(1022)
						Label1:
							// line 1023: self.body.append('<br />')
							πF.SetLineno(1023)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<br />").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_line.ToObject(), πTemp109); πE != nil {
						continue
					}
					// line 1025: def depart_line(self, node):
					πF.SetLineno(1025)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp110 = πg.NewFunction(πg.NewCode("depart_line", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1026: self.body.append('</div>\n')
							πF.SetLineno(1026)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_line.ToObject(), πTemp110); πE != nil {
						continue
					}
					// line 1028: def visit_line_block(self, node):
					πF.SetLineno(1028)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp111 = πg.NewFunction(πg.NewCode("visit_line_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1029: self.body.append(self.starttag(node, 'div', CLASS='line-block'))
							πF.SetLineno(1029)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("line-block").ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_line_block.ToObject(), πTemp111); πE != nil {
						continue
					}
					// line 1031: def depart_line_block(self, node):
					πF.SetLineno(1031)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp112 = πg.NewFunction(πg.NewCode("depart_line_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1032: self.body.append('</div>\n')
							πF.SetLineno(1032)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_line_block.ToObject(), πTemp112); πE != nil {
						continue
					}
					// line 1034: def visit_list_item(self, node):
					πF.SetLineno(1034)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp113 = πg.NewFunction(πg.NewCode("visit_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1035: self.body.append(self.starttag(node, 'li', ''))
							πF.SetLineno(1035)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_list_item.ToObject(), πTemp113); πE != nil {
						continue
					}
					// line 1037: def depart_list_item(self, node):
					πF.SetLineno(1037)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp114 = πg.NewFunction(πg.NewCode("depart_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1038: self.body.append('</li>\n')
							πF.SetLineno(1038)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</li>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_list_item.ToObject(), πTemp114); πE != nil {
						continue
					}
					// line 1041: def visit_literal(self, node):
					πF.SetLineno(1041)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp115 = πg.NewFunction(πg.NewCode("visit_literal", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5:
								goto Label5
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 1043: classes = node.get('classes', [])
							πF.SetLineno(1043)
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
							// line 1044: if 'code' in classes:
							πF.SetLineno(1044)
						Label1:
							// line 1046: node['classes'] = [cls for cls in classes if cls != 'code']
							πF.SetLineno(1046)
							πTemp006 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 1046: node['classes'] = [cls for cls in classes if cls != 'code']
										πF.SetLineno(1046)
									Label4:
										// line 1046: node['classes'] = [cls for cls in classes if cls != 'code']
										πF.SetLineno(1046)
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
							// line 1047: self.body.append(self.starttag(node, 'code', ''))
							πF.SetLineno(1047)
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
							// line 1048: return
							πF.SetLineno(1048)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 1049: self.body.append(
							πF.SetLineno(1049)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßspan.ToObject()
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
							// line 1051: text = node.astext()
							πF.SetLineno(1051)
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßliteral_block, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1053: if not isinstance(node.parent, nodes.literal_block):
							πF.SetLineno(1053)
						Label3:
							// line 1054: text = text.replace('\n', ' ')
							πF.SetLineno(1054)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("\n").ToObject()
							πTemp001[1] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp007
							goto Label4
						Label4:
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
							πF.PushCheckpoint(6)
							πTemp005 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label7
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
							πF.PushCheckpoint(5)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µtoken, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp011
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label8
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp001[0] = µtoken
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßin_word_wrap_point, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp008, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp007 = πTemp008
						Label8:
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label9
							}
							goto Label10
							// line 1058: if token.strip() and self.in_word_wrap_point.search(token):
							πF.SetLineno(1058)
						Label9:
							// line 1059: self.body.append('<span class="pre">%s</span>'
							πF.SetLineno(1059)
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
							goto Label11
						Label10:
							// line 1062: self.body.append(self.encode(token))
							πF.SetLineno(1062)
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
							goto Label11
						Label11:
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							// line 1063: self.body.append('</span>')
							πF.SetLineno(1063)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</span>").ToObject()
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
							// line 1065: raise nodes.SkipNode
							πF.SetLineno(1065)
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
					if πE = πClass.SetItem(πF, ßvisit_literal.ToObject(), πTemp115); πE != nil {
						continue
					}
					// line 1067: def depart_literal(self, node):
					πF.SetLineno(1067)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp116 = πg.NewFunction(πg.NewCode("depart_literal", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1069: self.body.append('</code>')
							πF.SetLineno(1069)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</code>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_literal.ToObject(), πTemp116); πE != nil {
						continue
					}
					// line 1071: def visit_literal_block(self, node):
					πF.SetLineno(1071)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp117 = πg.NewFunction(πg.NewCode("visit_literal_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
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
							// line 1072: self.body.append(self.starttag(node, 'pre', '', CLASS='literal-block'))
							πF.SetLineno(1072)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßpre.ToObject()
							πTemp002[2] = ß.ToObject()
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
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßclasses.ToObject()
							πTemp002 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp005
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.Contains(πF, πTemp006, ßcode.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 1073: if 'code' in node.get('classes', []):
							πF.SetLineno(1073)
						Label1:
							// line 1074: self.body.append('<code>')
							πF.SetLineno(1074)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<code>").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_literal_block.ToObject(), πTemp117); πE != nil {
						continue
					}
					// line 1076: def depart_literal_block(self, node):
					πF.SetLineno(1076)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp118 = πg.NewFunction(πg.NewCode("depart_literal_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßclasses.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.Contains(πF, πTemp005, ßcode.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 1077: if 'code' in node.get('classes', []):
							πF.SetLineno(1077)
						Label1:
							// line 1078: self.body.append('</code>')
							πF.SetLineno(1078)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("</code>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							// line 1079: self.body.append('</pre>\n')
							πF.SetLineno(1079)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("</pre>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_literal_block.ToObject(), πTemp118); πE != nil {
						continue
					}
					// line 1086: math_tags = {# math_output: (block, inline, class-arguments)
					πF.SetLineno(1086)
					πTemp004 = πg.NewDict()
					πTemp119 = πg.NewTuple3(ßdiv.ToObject(), ß.ToObject(), ß.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßmathml.ToObject(), πTemp119); πE != nil {
						continue
					}
					πTemp119 = πg.NewTuple3(ßdiv.ToObject(), ßspan.ToObject(), ßformula.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßhtml.ToObject(), πTemp119); πE != nil {
						continue
					}
					πTemp119 = πg.NewTuple3(ßdiv.ToObject(), ßspan.ToObject(), ßmath.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßmathjax.ToObject(), πTemp119); πE != nil {
						continue
					}
					πTemp119 = πg.NewTuple3(ßpre.ToObject(), ßtt.ToObject(), ßmath.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßlatex.ToObject(), πTemp119); πE != nil {
						continue
					}
					πTemp119 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßmath_tags.ToObject(), πTemp119); πE != nil {
						continue
					}
					// line 1093: def visit_math(self, node, math_env=''):
					πF.SetLineno(1093)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp005[2] = πg.Param{Name: "math_env", Def: ß.ToObject()}
					πTemp119 = πg.NewFunction(πg.NewCode("visit_math", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µmath_env *πg.Object = πArgs[2]
						_ = µmath_env
						var µtag *πg.Object = πg.UnboundLocal
						_ = µtag
						var µclsarg *πg.Object = πg.UnboundLocal
						_ = µclsarg
						var µwrappers *πg.Object = πg.UnboundLocal
						_ = µwrappers
						var µwrapper *πg.Object = πg.UnboundLocal
						_ = µwrapper
						var µmath_code *πg.Object = πg.UnboundLocal
						_ = µmath_code
						var µconverter *πg.Object = πg.UnboundLocal
						_ = µconverter
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
						var µerr_node *πg.Object = πg.UnboundLocal
						_ = µerr_node
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Dict
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.BaseException
						_ = πTemp012
						var πTemp013 *πg.Traceback
						_ = πTemp013
						var πTemp014 []πg.Param
						_ = πTemp014
						var πTemp015 πg.KWArgs
						_ = πTemp015
						var πTemp016 []*πg.Object
						_ = πTemp016
						var πTemp017 []*πg.Object
						_ = πTemp017
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10:
								goto Label10
							case 20:
								goto Label20
							case 30:
								goto Label30
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_tags, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1096: if self.math_output not in self.math_tags:
							πF.SetLineno(1096)
						Label1:
							// line 1097: self.document.reporter.error(
							πF.SetLineno(1097)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("math-output format \"%s\" not supported falling back to \"latex\"").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1100: self.math_output = 'latex'
							πF.SetLineno(1100)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ßlatex.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_output, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 1101: tag = self.math_tags[self.math_output][math_env == '']
							πF.SetLineno(1101)
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µmath_env, ß.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßmath_tags, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µtag = πTemp002
							// line 1102: clsarg = self.math_tags[self.math_output][2]
							πF.SetLineno(1102)
							πTemp001 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßmath_tags, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µclsarg = πTemp002
							// line 1104: wrappers = {# math_mode: (inline, block)
							πF.SetLineno(1104)
							πTemp008 = πg.NewDict()
							πTemp001 = πg.NewTuple2(πg.NewStr("$%s$").ToObject(), πg.NewUnicode("\\begin{%s}\n%s\n\\end{%s}").ToObject()).ToObject()
							if πE = πTemp008.SetItem(πF, ßmathml.ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πg.NewStr("$%s$").ToObject(), πg.NewUnicode("\\begin{%s}\n%s\n\\end{%s}").ToObject()).ToObject()
							if πE = πTemp008.SetItem(πF, ßhtml.ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πg.NewStr("\\(%s\\)").ToObject(), πg.NewUnicode("\\begin{%s}\n%s\n\\end{%s}").ToObject()).ToObject()
							if πE = πTemp008.SetItem(πF, ßmathjax.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
							if πE = πTemp008.SetItem(πF, ßlatex.ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp008.ToObject()
							µwrappers = πTemp001
							// line 1110: wrapper = wrappers[self.math_output][math_env != '']
							πF.SetLineno(1110)
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µmath_env, ß.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µwrappers, "wrappers"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µwrappers, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µwrapper = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp003, ßmathml.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmath_output_options, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp010).ToObject()
							πTemp002 = πTemp003
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label4
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßmath_output_options, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp011, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp007, ßblahtexml.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label4:
							πTemp001 = πTemp002
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 1111: if self.math_output == 'mathml' and (not self.math_output_options or
							πF.SetLineno(1111)
						Label5:
							// line 1113: wrapper = None
							πF.SetLineno(1113)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µwrapper = πTemp001
							goto Label6
						Label6:
							// line 1115: math_code = node.astext().translate(unichar2tex.uni2tex_table)
							πF.SetLineno(1115)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunichar2tex); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßuni2tex_table, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtranslate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp002
							if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µwrapper); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 1116: if wrapper:
							πF.SetLineno(1116)
						Label7:
							// line 1117: try: # wrapper with three "%s"
							πF.SetLineno(1117)
							πF.PushCheckpoint(10)
							// line 1118: math_code = wrapper % (math_env, math_code, math_env)
							πF.SetLineno(1118)
							if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µmath_env, µmath_code, µmath_env).ToObject()
							if πTemp001, πE = πg.Mod(πF, µwrapper, πTemp002); πE != nil {
								continue
							}
							µmath_code = πTemp001
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp012, πTemp013 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
							continue
							// line 1119: except TypeError: # wrapper with one "%s"
							πF.SetLineno(1119)
						Label11:
							// line 1120: math_code = wrapper % math_code
							πF.SetLineno(1120)
							if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, µwrapper, µmath_code); πE != nil {
								continue
							}
							µmath_code = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(ßlatex.ToObject(), ßmathjax.ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 1122: if self.math_output in ('latex', 'mathjax'):
							πF.SetLineno(1122)
						Label12:
							// line 1123: math_code = self.encode(math_code)
							πF.SetLineno(1123)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp002
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp003, ßmathjax.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp009).ToObject()
							πTemp001 = πTemp002
						Label14:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßhtml.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßmathml.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							goto Label18
							// line 1124: if self.math_output == 'mathjax' and not self.math_header:
							πF.SetLineno(1124)
						Label15:
							// line 1125: try:
							πF.SetLineno(1125)
							πF.PushCheckpoint(20)
							// line 1126: self.mathjax_url = self.math_output_options[0]
							πF.SetLineno(1126)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_output_options, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmathjax_url, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label19
						Label20:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp012, πTemp013 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
							continue
							// line 1127: except IndexError:
							πF.SetLineno(1127)
						Label21:
							// line 1128: self.document.reporter.warning('No MathJax URL specified, '
							πF.SetLineno(1128)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("No MathJax URL specified, using local fallback (see config.html)").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.RestoreExc(nil, nil)
							goto Label19
						Label19:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmathjax_url, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp002, πg.NewStr("?").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label22
							}
							goto Label23
							// line 1132: if '?' not in self.mathjax_url:
							πF.SetLineno(1132)
						Label22:
							// line 1133: self.mathjax_url += '?config=TeX-AMS_CHTML'
							πF.SetLineno(1133)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmathjax_url, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewStr("?config=TeX-AMS_CHTML").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmathjax_url, πTemp002); πE != nil {
								continue
							}
							goto Label23
						Label23:
							// line 1134: self.math_header = [self.mathjax_script % self.mathjax_url]
							πF.SetLineno(1134)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmathjax_script, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmathjax_url, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_header, πTemp002); πE != nil {
								continue
							}
							goto Label18
							// line 1135: elif self.math_output == 'html':
							πF.SetLineno(1135)
						Label16:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmath_output_options, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label24
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp009).ToObject()
							πTemp001 = πTemp002
						Label24:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label25
							}
							goto Label26
							// line 1136: if self.math_output_options and not self.math_header:
							πF.SetLineno(1136)
						Label25:
							// line 1137: self.math_header = [self.stylesheet_call(
							πF.SetLineno(1137)
							πTemp014 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µs *πg.Object = πg.UnboundLocal
								_ = µs
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 bool
								_ = πTemp006
								var πTemp007 bool
								_ = πTemp007
								var πTemp008 []*πg.Object
								_ = πTemp008
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
										πTemp002[0] = πg.NewStr(",").ToObject()
										πTemp003 = πg.NewInt(0).ToObject()
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetAttr(πF, µself, ßmath_output_options, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
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
										πTemp006 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp006 {
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
											πTemp007 = !isStop
										} else {
											πTemp007 = true
											µs = πTemp003
										}
										if πE != nil || !πTemp007 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1137: self.math_header = [self.stylesheet_call(
										πF.SetLineno(1137)
										πTemp002 = πF.MakeArgs(1)
										πTemp008 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
											continue
										}
										πTemp008[0] = µs
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstylesheet_dirs, nil); πE != nil {
											continue
										}
										πTemp008[1] = πTemp004
										if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfind_file_in_dirs, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp008)
										πTemp002[0] = πTemp003
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßstylesheet_call, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmath_header, πTemp003); πE != nil {
								continue
							}
							goto Label26
						Label26:
							// line 1141: math2html.DocumentParameters.displaymode = (math_env != '')
							πF.SetLineno(1141)
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µmath_env, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßmath2html); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßDocumentParameters, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp007, ßdisplaymode, πTemp003); πE != nil {
								continue
							}
							// line 1142: math_code = math2html.math2html(math_code)
							πF.SetLineno(1142)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath2html); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmath2html, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp001
							goto Label18
							// line 1143: elif self.math_output == 'mathml':
							πF.SetLineno(1143)
						Label17:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdoctype, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πg.NewStr("XHTML 1").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label27
							}
							goto Label28
							// line 1144: if  'XHTML 1' in self.doctype:
							πF.SetLineno(1144)
						Label27:
							// line 1145: self.doctype = self.doctype_mathml
							πF.SetLineno(1145)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdoctype_mathml, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdoctype, πTemp003); πE != nil {
								continue
							}
							// line 1146: self.content_type = self.content_type_mathml
							πF.SetLineno(1146)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_type_mathml, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontent_type, πTemp003); πE != nil {
								continue
							}
							goto Label28
						Label28:
							// line 1147: converter = ' '.join(self.math_output_options).lower()
							πF.SetLineno(1147)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmath_output_options, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µconverter = πTemp003
							// line 1148: try:
							πF.SetLineno(1148)
							πF.PushCheckpoint(30)
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µconverter, ßlatexml.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label31
							}
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µconverter, ßttm.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label32
							}
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µconverter, ßblahtexml.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label33
							}
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µconverter); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label34
							}
							goto Label35
							// line 1149: if converter == 'latexml':
							πF.SetLineno(1149)
						Label31:
							// line 1150: math_code = tex2mathml_extern.latexml(math_code,
							πF.SetLineno(1150)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2mathml_extern); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlatexml, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp001
							goto Label36
							// line 1152: elif converter == 'ttm':
							πF.SetLineno(1152)
						Label32:
							// line 1153: math_code = tex2mathml_extern.ttm(math_code,
							πF.SetLineno(1153)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2mathml_extern); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßttm, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp001
							goto Label36
							// line 1155: elif converter == 'blahtexml':
							πF.SetLineno(1155)
						Label33:
							// line 1156: math_code = tex2mathml_extern.blahtexml(math_code,
							πF.SetLineno(1156)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmath_env); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"inline", πTemp001},
								{"reporter", πTemp006},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2mathml_extern); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßblahtexml, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp001
							goto Label36
							// line 1159: elif not converter:
							πF.SetLineno(1159)
						Label34:
							// line 1160: math_code = latex2mathml.tex2mathml(math_code,
							πF.SetLineno(1160)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmath_env); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp015 = πg.KWArgs{
								{"inline", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlatex2mathml); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtex2mathml, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmath_code = πTemp001
							goto Label36
						Label35:
							// line 1163: self.document.reporter.error('option "%s" not supported '
							πF.SetLineno(1163)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("option \"%s\" not supported with math-output \"MathML\"").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label36
						Label36:
							πF.PopCheckpoint()
							goto Label29
						Label30:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp012, πTemp013 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label37
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label38
							}
							πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
							continue
							// line 1165: except OSError:
							πF.SetLineno(1165)
						Label37:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("is \"latexmlmath\" in your PATH?").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1166: raise OSError('is "latexmlmath" in your PATH?')
							πF.SetLineno(1166)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label29
							// line 1167: except SyntaxError as err:
							πF.SetLineno(1167)
						Label38:
							µerr = πTemp012.ToObject()
							// line 1168: err_node = self.document.reporter.error(err, base_node=node)
							πF.SetLineno(1168)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp005[0] = µerr
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µerr_node = πTemp003
							// line 1169: self.visit_system_message(err_node)
							πF.SetLineno(1169)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr_node, "err_node"); πE != nil {
								continue
							}
							πTemp005[0] = µerr_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvisit_system_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1170: self.body.append(self.starttag(node, 'p'))
							πF.SetLineno(1170)
							πTemp005 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp016[0] = µnode
							πTemp016[1] = ßp.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp016, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1171: self.body.append(u','.join(err.args))
							πF.SetLineno(1171)
							πTemp005 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerr, ßargs, nil); πE != nil {
								continue
							}
							πTemp016[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewUnicode(",").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp016, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1172: self.body.append('</p>\n')
							πF.SetLineno(1172)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("</p>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1173: self.body.append(self.starttag(node, 'pre',
							πF.SetLineno(1173)
							πTemp005 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp016[0] = µnode
							πTemp016[1] = ßpre.ToObject()
							πTemp015 = πg.KWArgs{
								{"CLASS", πg.NewStr("literal-block").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp016, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1175: self.body.append(self.encode(math_code))
							πF.SetLineno(1175)
							πTemp005 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp016[0] = µmath_code
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp016, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1176: self.body.append('\n</pre>\n')
							πF.SetLineno(1176)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n</pre>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1177: self.depart_system_message(err_node)
							πF.SetLineno(1177)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr_node, "err_node"); πE != nil {
								continue
							}
							πTemp005[0] = µerr_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdepart_system_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1178: raise nodes.SkipNode
							πF.SetLineno(1178)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label29
						Label29:
							goto Label18
						Label18:
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µtag); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label39
							}
							goto Label40
							// line 1180: if tag:
							πF.SetLineno(1180)
						Label39:
							// line 1181: self.body.append(self.starttag(node, tag,
							πF.SetLineno(1181)
							πTemp005 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp016[0] = µnode
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							πTemp016[1] = µtag
							πTemp017 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							πTemp017[0] = µmath_env
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp017, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp017)
							if πTemp001, πE = πg.Mul(πF, πg.NewStr("\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µclsarg, "clsarg"); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"suffix", πTemp001},
								{"CLASS", µclsarg},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp016, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label40
						Label40:
							// line 1184: self.body.append(math_code)
							πF.SetLineno(1184)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
								continue
							}
							πTemp005[0] = µmath_code
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmath_env); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label41
							}
							goto Label42
							// line 1185: if math_env: # block mode (equation, display)
							πF.SetLineno(1185)
						Label41:
							// line 1186: self.body.append('\n')
							πF.SetLineno(1186)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label42
						Label42:
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µtag); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label43
							}
							goto Label44
							// line 1187: if tag:
							πF.SetLineno(1187)
						Label43:
							// line 1188: self.body.append('</%s>' % tag)
							πF.SetLineno(1188)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("</%s>").ToObject(), µtag); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label44
						Label44:
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmath_env); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label45
							}
							goto Label46
							// line 1189: if math_env:
							πF.SetLineno(1189)
						Label45:
							// line 1190: self.body.append('\n')
							πF.SetLineno(1190)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label46
						Label46:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1192: raise nodes.SkipNode
							πF.SetLineno(1192)
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
					if πE = πClass.SetItem(πF, ßvisit_math.ToObject(), πTemp119); πE != nil {
						continue
					}
					// line 1194: def depart_math(self, node):
					πF.SetLineno(1194)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp120 = πg.NewFunction(πg.NewCode("depart_math", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1195: pass # never reached
							πF.SetLineno(1195)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_math.ToObject(), πTemp120); πE != nil {
						continue
					}
					// line 1197: def visit_math_block(self, node):
					πF.SetLineno(1197)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp121 = πg.NewFunction(πg.NewCode("visit_math_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µmath_env *πg.Object = πg.UnboundLocal
						_ = µmath_env
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 1198: math_env = pick_math_environment(node.astext())
							πF.SetLineno(1198)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpick_math_environment); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmath_env = πTemp003
							// line 1199: self.visit_math(node, math_env=math_env)
							πF.SetLineno(1199)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µmath_env, "math_env"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"math_env", µmath_env},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_math, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_math_block.ToObject(), πTemp121); πE != nil {
						continue
					}
					// line 1201: def depart_math_block(self, node):
					πF.SetLineno(1201)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp122 = πg.NewFunction(πg.NewCode("depart_math_block", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1202: pass # never reached
							πF.SetLineno(1202)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_math_block.ToObject(), πTemp122); πE != nil {
						continue
					}
					// line 1206: def visit_meta(self, node):
					πF.SetLineno(1206)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp123 = πg.NewFunction(πg.NewCode("visit_meta", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µmeta *πg.Object = πg.UnboundLocal
						_ = µmeta
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
							// line 1207: meta = self.emptytag(node, 'meta', **node.non_default_attributes())
							πF.SetLineno(1207)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßmeta.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßnon_default_attributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßemptytag, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmeta = πTemp004
							// line 1208: self.add_meta(meta)
							πF.SetLineno(1208)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmeta, "meta"); πE != nil {
								continue
							}
							πTemp001[0] = µmeta
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_meta, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_meta.ToObject(), πTemp123); πE != nil {
						continue
					}
					// line 1210: def depart_meta(self, node):
					πF.SetLineno(1210)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp124 = πg.NewFunction(πg.NewCode("depart_meta", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1211: pass
							πF.SetLineno(1211)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_meta.ToObject(), πTemp124); πE != nil {
						continue
					}
					// line 1213: def add_meta(self, tag):
					πF.SetLineno(1213)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "tag", Def: nil}
					πTemp125 = πg.NewFunction(πg.NewCode("add_meta", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtag *πg.Object = πArgs[1]
						_ = µtag
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
							// line 1214: self.meta.append(tag)
							πF.SetLineno(1214)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							πTemp001[0] = µtag
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmeta, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1215: self.head.append(tag)
							πF.SetLineno(1215)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							πTemp001[0] = µtag
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_meta.ToObject(), πTemp125); πE != nil {
						continue
					}
					// line 1217: def visit_option(self, node):
					πF.SetLineno(1217)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp126 = πg.NewFunction(πg.NewCode("visit_option", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1218: self.body.append(self.starttag(node, 'span', '', CLASS='option'))
							πF.SetLineno(1218)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßspan.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßoption.ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_option.ToObject(), πTemp126); πE != nil {
						continue
					}
					// line 1220: def depart_option(self, node):
					πF.SetLineno(1220)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp127 = πg.NewFunction(πg.NewCode("depart_option", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 πg.KWArgs
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
							// line 1221: self.body.append('</span>')
							πF.SetLineno(1221)
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßoption, nil); πE != nil {
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
							// line 1222: if isinstance(node.next_node(descend=False, siblings=True),
							πF.SetLineno(1222)
						Label1:
							// line 1224: self.body.append(', ')
							πF.SetLineno(1224)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr(", ").ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option.ToObject(), πTemp127); πE != nil {
						continue
					}
					// line 1226: def visit_option_argument(self, node):
					πF.SetLineno(1226)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp128 = πg.NewFunction(πg.NewCode("visit_option_argument", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1227: self.body.append(node.get('delimiter', ' '))
							πF.SetLineno(1227)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßdelimiter.ToObject()
							πTemp002[1] = πg.NewStr(" ").ToObject()
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
							// line 1228: self.body.append(self.starttag(node, 'var', ''))
							πF.SetLineno(1228)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßvar.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_option_argument.ToObject(), πTemp128); πE != nil {
						continue
					}
					// line 1230: def depart_option_argument(self, node):
					πF.SetLineno(1230)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp129 = πg.NewFunction(πg.NewCode("depart_option_argument", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1231: self.body.append('</var>')
							πF.SetLineno(1231)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</var>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_option_argument.ToObject(), πTemp129); πE != nil {
						continue
					}
					// line 1233: def visit_option_group(self, node):
					πF.SetLineno(1233)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp130 = πg.NewFunction(πg.NewCode("visit_option_group", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1234: self.body.append(self.starttag(node, 'dt', ''))
							πF.SetLineno(1234)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdt.ToObject()
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
							// line 1235: self.body.append('<kbd>')
							πF.SetLineno(1235)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<kbd>").ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_option_group.ToObject(), πTemp130); πE != nil {
						continue
					}
					// line 1237: def depart_option_group(self, node):
					πF.SetLineno(1237)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp131 = πg.NewFunction(πg.NewCode("depart_option_group", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1238: self.body.append('</kbd></dt>\n')
							πF.SetLineno(1238)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</kbd></dt>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_option_group.ToObject(), πTemp131); πE != nil {
						continue
					}
					// line 1240: def visit_option_list(self, node):
					πF.SetLineno(1240)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp132 = πg.NewFunction(πg.NewCode("visit_option_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1241: self.body.append(
							πF.SetLineno(1241)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdl.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", πg.NewStr("option-list").ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_option_list.ToObject(), πTemp132); πE != nil {
						continue
					}
					// line 1244: def depart_option_list(self, node):
					πF.SetLineno(1244)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp133 = πg.NewFunction(πg.NewCode("depart_option_list", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1245: self.body.append('</dl>\n')
							πF.SetLineno(1245)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</dl>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_option_list.ToObject(), πTemp133); πE != nil {
						continue
					}
					// line 1247: def visit_option_list_item(self, node):
					πF.SetLineno(1247)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp134 = πg.NewFunction(πg.NewCode("visit_option_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1248: pass
							πF.SetLineno(1248)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_list_item.ToObject(), πTemp134); πE != nil {
						continue
					}
					// line 1250: def depart_option_list_item(self, node):
					πF.SetLineno(1250)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp135 = πg.NewFunction(πg.NewCode("depart_option_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1251: pass
							πF.SetLineno(1251)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_list_item.ToObject(), πTemp135); πE != nil {
						continue
					}
					// line 1253: def visit_option_string(self, node):
					πF.SetLineno(1253)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp136 = πg.NewFunction(πg.NewCode("visit_option_string", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1254: pass
							πF.SetLineno(1254)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_option_string.ToObject(), πTemp136); πE != nil {
						continue
					}
					// line 1256: def depart_option_string(self, node):
					πF.SetLineno(1256)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp137 = πg.NewFunction(πg.NewCode("depart_option_string", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1257: pass
							πF.SetLineno(1257)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_option_string.ToObject(), πTemp137); πE != nil {
						continue
					}
					// line 1259: def visit_organization(self, node):
					πF.SetLineno(1259)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp138 = πg.NewFunction(πg.NewCode("visit_organization", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1260: self.visit_docinfo_item(node, 'organization')
							πF.SetLineno(1260)
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
					if πE = πClass.SetItem(πF, ßvisit_organization.ToObject(), πTemp138); πE != nil {
						continue
					}
					// line 1262: def depart_organization(self, node):
					πF.SetLineno(1262)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp139 = πg.NewFunction(πg.NewCode("depart_organization", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1263: self.depart_docinfo_item()
							πF.SetLineno(1263)
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
					if πE = πClass.SetItem(πF, ßdepart_organization.ToObject(), πTemp139); πE != nil {
						continue
					}
					// line 1280: def visit_paragraph(self, node):
					πF.SetLineno(1280)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp140 = πg.NewFunction(πg.NewCode("visit_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1281: self.body.append(self.starttag(node, 'p', ''))
							πF.SetLineno(1281)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßp.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_paragraph.ToObject(), πTemp140); πE != nil {
						continue
					}
					// line 1283: def depart_paragraph(self, node):
					πF.SetLineno(1283)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp141 = πg.NewFunction(πg.NewCode("depart_paragraph", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1284: self.body.append('</p>')
							πF.SetLineno(1284)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>").ToObject()
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßlist_item, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßentry, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 1285: if not (isinstance(node.parent, (nodes.list_item, nodes.entry)) and
							πF.SetLineno(1285)
						Label2:
							// line 1287: self.body.append('\n')
							πF.SetLineno(1287)
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
					if πE = πClass.SetItem(πF, ßdepart_paragraph.ToObject(), πTemp141); πE != nil {
						continue
					}
					// line 1289: def visit_problematic(self, node):
					πF.SetLineno(1289)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp142 = πg.NewFunction(πg.NewCode("visit_problematic", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
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
							// line 1290: if node.hasattr('refid'):
							πF.SetLineno(1290)
						Label1:
							// line 1291: self.body.append('<a href="#%s">' % node['refid'])
							πF.SetLineno(1291)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<a href=\"#%s\">").ToObject(), πTemp005); πE != nil {
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
							// line 1292: self.context.append('</a>')
							πF.SetLineno(1292)
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
							goto Label3
						Label2:
							// line 1294: self.context.append('')
							πF.SetLineno(1294)
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
						Label3:
							// line 1295: self.body.append(self.starttag(node, 'span', '', CLASS='problematic'))
							πF.SetLineno(1295)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							πTemp006[1] = ßspan.ToObject()
							πTemp006[2] = ß.ToObject()
							πTemp007 = πg.KWArgs{
								{"CLASS", ßproblematic.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßvisit_problematic.ToObject(), πTemp142); πE != nil {
						continue
					}
					// line 1297: def depart_problematic(self, node):
					πF.SetLineno(1297)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp143 = πg.NewFunction(πg.NewCode("depart_problematic", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1298: self.body.append('</span>')
							πF.SetLineno(1298)
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
							// line 1299: self.body.append(self.context.pop())
							πF.SetLineno(1299)
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
					if πE = πClass.SetItem(πF, ßdepart_problematic.ToObject(), πTemp143); πE != nil {
						continue
					}
					// line 1301: def visit_raw(self, node):
					πF.SetLineno(1301)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp144 = πg.NewFunction(πg.NewCode("visit_raw", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µt *πg.Object = πg.UnboundLocal
						_ = µt
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
						var πTemp007 *πg.Object
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
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßformat.ToObject()
							πTemp002[1] = ß.ToObject()
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, ßhtml.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1302: if 'html' in node.get('format', '').split():
							πF.SetLineno(1302)
						Label1:
							// line 1303: t = isinstance(node.parent, nodes.TextElement) and 'span' or 'div'
							πF.SetLineno(1303)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp007
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							πTemp003 = ßspan.ToObject()
						Label4:
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πTemp001 = ßdiv.ToObject()
						Label3:
							µt = πTemp001
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 1304: if node['classes']:
							πF.SetLineno(1304)
						Label5:
							// line 1305: self.body.append(self.starttag(node, t, suffix=''))
							πF.SetLineno(1305)
							πTemp002 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008[0] = µnode
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp008[1] = µt
							πTemp009 = πg.KWArgs{
								{"suffix", ß.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002[0] = πTemp003
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
							goto Label6
						Label6:
							// line 1306: self.body.append(node.astext())
							πF.SetLineno(1306)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
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
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1307: if node['classes']:
							πF.SetLineno(1307)
						Label7:
							// line 1308: self.body.append('</%s>' % t)
							πF.SetLineno(1308)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("</%s>").ToObject(), µt); πE != nil {
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
							goto Label8
						Label8:
							goto Label2
						Label2:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1310: raise nodes.SkipNode
							πF.SetLineno(1310)
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
					if πE = πClass.SetItem(πF, ßvisit_raw.ToObject(), πTemp144); πE != nil {
						continue
					}
					// line 1312: def visit_reference(self, node):
					πF.SetLineno(1312)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp145 = πg.NewFunction(πg.NewCode("visit_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1313: atts = {'class': 'reference'}
							πF.SetLineno(1313)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßclass.ToObject(), ßreference.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µnode, ßrefuri.ToObject()); πE != nil {
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
							// line 1314: if 'refuri' in node:
							πF.SetLineno(1314)
						Label1:
							// line 1315: atts['href'] = node['refuri']
							πF.SetLineno(1315)
							πTemp002 = ßrefuri.ToObject()
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
							πTemp005 = ßhref.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcloak_email_addresses, nil); πE != nil {
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
							πTemp006[0] = πg.NewStr("mailto:").ToObject()
							πTemp004 = ßhref.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µatts, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πTemp005
						Label4:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 1316: if ( self.settings.cloak_email_addresses
							πF.SetLineno(1316)
						Label5:
							// line 1318: atts['href'] = self.cloak_mailto(atts['href'])
							πF.SetLineno(1318)
							πTemp006 = πF.MakeArgs(1)
							πTemp002 = ßhref.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcloak_mailto, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßhref.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 1319: self.in_mailto = True
							πF.SetLineno(1319)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_mailto, πTemp004); πE != nil {
								continue
							}
							goto Label6
						Label6:
							// line 1320: atts['class'] += ' external'
							πF.SetLineno(1320)
							πTemp002 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp004, πg.NewStr(" external").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 1322: assert 'refid' in node, \
							πF.SetLineno(1322)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µnode, ßrefid.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003).ToObject()
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("References must have \"refuri\" or \"refid\" attribute.").ToObject()); πE != nil {
								continue
							}
							// line 1324: atts['href'] = '#' + node['refid']
							πF.SetLineno(1324)
							πTemp004 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("#").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßhref.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 1325: atts['class'] += ' internal'
							πF.SetLineno(1325)
							πTemp002 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp004, πg.NewStr(" internal").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label3:
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 1326: if not isinstance(node.parent, nodes.TextElement):
							πF.SetLineno(1326)
						Label7:
							// line 1327: assert len(node) == 1 and isinstance(node[0], nodes.image)
							πF.SetLineno(1327)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label9
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßimage, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πTemp005
						Label9:
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							// line 1328: atts['class'] += ' image-reference'
							πF.SetLineno(1328)
							πTemp002 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µatts, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp004, πg.NewStr(" image-reference").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label8
						Label8:
							// line 1329: self.body.append(self.starttag(node, 'a', '', **atts))
							πF.SetLineno(1329)
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008[0] = µnode
							πTemp008[1] = ßa.ToObject()
							πTemp008[2] = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_reference.ToObject(), πTemp145); πE != nil {
						continue
					}
					// line 1331: def depart_reference(self, node):
					πF.SetLineno(1331)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp146 = πg.NewFunction(πg.NewCode("depart_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1332: self.body.append('</a>')
							πF.SetLineno(1332)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</a>").ToObject()
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTextElement, nil); πE != nil {
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
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1333: if not isinstance(node.parent, nodes.TextElement):
							πF.SetLineno(1333)
						Label1:
							// line 1334: self.body.append('\n')
							πF.SetLineno(1334)
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
							goto Label2
						Label2:
							// line 1335: self.in_mailto = False
							πF.SetLineno(1335)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_mailto, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_reference.ToObject(), πTemp146); πE != nil {
						continue
					}
					// line 1337: def visit_revision(self, node):
					πF.SetLineno(1337)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp147 = πg.NewFunction(πg.NewCode("visit_revision", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 1338: self.visit_docinfo_item(node, 'revision', meta=False)
							πF.SetLineno(1338)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßrevision.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_revision.ToObject(), πTemp147); πE != nil {
						continue
					}
					// line 1340: def depart_revision(self, node):
					πF.SetLineno(1340)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp148 = πg.NewFunction(πg.NewCode("depart_revision", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1341: self.depart_docinfo_item()
							πF.SetLineno(1341)
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
					if πE = πClass.SetItem(πF, ßdepart_revision.ToObject(), πTemp148); πE != nil {
						continue
					}
					// line 1343: def visit_row(self, node):
					πF.SetLineno(1343)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp149 = πg.NewFunction(πg.NewCode("visit_row", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1344: self.body.append(self.starttag(node, 'tr', ''))
							πF.SetLineno(1344)
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
							// line 1345: node.column = 0
							πF.SetLineno(1345)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnode, ßcolumn, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_row.ToObject(), πTemp149); πE != nil {
						continue
					}
					// line 1347: def depart_row(self, node):
					πF.SetLineno(1347)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp150 = πg.NewFunction(πg.NewCode("depart_row", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1348: self.body.append('</tr>\n')
							πF.SetLineno(1348)
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
					if πE = πClass.SetItem(πF, ßdepart_row.ToObject(), πTemp150); πE != nil {
						continue
					}
					// line 1350: def visit_rubric(self, node):
					πF.SetLineno(1350)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp151 = πg.NewFunction(πg.NewCode("visit_rubric", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1351: self.body.append(self.starttag(node, 'p', '', CLASS='rubric'))
							πF.SetLineno(1351)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßp.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßrubric.ToObject()},
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
					if πE = πClass.SetItem(πF, ßvisit_rubric.ToObject(), πTemp151); πE != nil {
						continue
					}
					// line 1353: def depart_rubric(self, node):
					πF.SetLineno(1353)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp152 = πg.NewFunction(πg.NewCode("depart_rubric", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1354: self.body.append('</p>\n')
							πF.SetLineno(1354)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_rubric.ToObject(), πTemp152); πE != nil {
						continue
					}
					// line 1357: def visit_section(self, node):
					πF.SetLineno(1357)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp153 = πg.NewFunction(πg.NewCode("visit_section", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1358: self.section_level += 1
							πF.SetLineno(1358)
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
							// line 1359: self.body.append(
							πF.SetLineno(1359)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp004[1] = ßdiv.ToObject()
							πTemp005 = πg.KWArgs{
								{"CLASS", ßsection.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßvisit_section.ToObject(), πTemp153); πE != nil {
						continue
					}
					// line 1362: def depart_section(self, node):
					πF.SetLineno(1362)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp154 = πg.NewFunction(πg.NewCode("depart_section", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1363: self.section_level -= 1
							πF.SetLineno(1363)
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
							// line 1364: self.body.append('</div>\n')
							πF.SetLineno(1364)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_section.ToObject(), πTemp154); πE != nil {
						continue
					}
					// line 1367: def visit_sidebar(self, node):
					πF.SetLineno(1367)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp155 = πg.NewFunction(πg.NewCode("visit_sidebar", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1368: self.body.append(
							πF.SetLineno(1368)
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
							// line 1370: self.in_sidebar = True
							πF.SetLineno(1370)
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
					if πE = πClass.SetItem(πF, ßvisit_sidebar.ToObject(), πTemp155); πE != nil {
						continue
					}
					// line 1372: def depart_sidebar(self, node):
					πF.SetLineno(1372)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp156 = πg.NewFunction(πg.NewCode("depart_sidebar", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1373: self.body.append('</div>\n')
							πF.SetLineno(1373)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
							// line 1374: self.in_sidebar = False
							πF.SetLineno(1374)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_sidebar, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_sidebar.ToObject(), πTemp156); πE != nil {
						continue
					}
					// line 1376: def visit_status(self, node):
					πF.SetLineno(1376)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp157 = πg.NewFunction(πg.NewCode("visit_status", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 1377: self.visit_docinfo_item(node, 'status', meta=False)
							πF.SetLineno(1377)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßstatus.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_status.ToObject(), πTemp157); πE != nil {
						continue
					}
					// line 1379: def depart_status(self, node):
					πF.SetLineno(1379)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp158 = πg.NewFunction(πg.NewCode("depart_status", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1380: self.depart_docinfo_item()
							πF.SetLineno(1380)
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
					if πE = πClass.SetItem(πF, ßdepart_status.ToObject(), πTemp158); πE != nil {
						continue
					}
					// line 1382: def visit_strong(self, node):
					πF.SetLineno(1382)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp159 = πg.NewFunction(πg.NewCode("visit_strong", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1383: self.body.append(self.starttag(node, 'strong', ''))
							πF.SetLineno(1383)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßstrong.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_strong.ToObject(), πTemp159); πE != nil {
						continue
					}
					// line 1385: def depart_strong(self, node):
					πF.SetLineno(1385)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp160 = πg.NewFunction(πg.NewCode("depart_strong", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1386: self.body.append('</strong>')
							πF.SetLineno(1386)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</strong>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_strong.ToObject(), πTemp160); πE != nil {
						continue
					}
					// line 1388: def visit_subscript(self, node):
					πF.SetLineno(1388)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp161 = πg.NewFunction(πg.NewCode("visit_subscript", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1389: self.body.append(self.starttag(node, 'sub', ''))
							πF.SetLineno(1389)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßsub.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_subscript.ToObject(), πTemp161); πE != nil {
						continue
					}
					// line 1391: def depart_subscript(self, node):
					πF.SetLineno(1391)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp162 = πg.NewFunction(πg.NewCode("depart_subscript", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1392: self.body.append('</sub>')
							πF.SetLineno(1392)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_subscript.ToObject(), πTemp162); πE != nil {
						continue
					}
					// line 1394: def visit_substitution_definition(self, node):
					πF.SetLineno(1394)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp163 = πg.NewFunction(πg.NewCode("visit_substitution_definition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1395: """Internal only."""
							πF.SetLineno(1395)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 1396: raise nodes.SkipNode
							πF.SetLineno(1396)
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
					if πE = πClass.SetItem(πF, ßvisit_substitution_definition.ToObject(), πTemp163); πE != nil {
						continue
					}
					// line 1395: """Internal only."""
					πF.SetLineno(1395)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp164}, πg.NewStr("Internal only.").ToObject()); πE != nil {
						continue
					}
					if πTemp165, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_substitution_definition); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp165, ß__doc__, πTemp164); πE != nil {
						continue
					}
					// line 1398: def visit_substitution_reference(self, node):
					πF.SetLineno(1398)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp164 = πg.NewFunction(πg.NewCode("visit_substitution_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1399: self.unimplemented_visit(node)
							πF.SetLineno(1399)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßunimplemented_visit, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_substitution_reference.ToObject(), πTemp164); πE != nil {
						continue
					}
					// line 1405: def visit_subtitle(self, node):
					πF.SetLineno(1405)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp165 = πg.NewFunction(πg.NewCode("visit_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
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
							// line 1406: if isinstance(node.parent, nodes.sidebar):
							πF.SetLineno(1406)
						Label1:
							// line 1407: classes = 'sidebar-subtitle'
							πF.SetLineno(1407)
							µclasses = πg.NewStr("sidebar-subtitle").ToObject()
							goto Label4
							// line 1408: elif isinstance(node.parent, nodes.document):
							πF.SetLineno(1408)
						Label2:
							// line 1409: classes = 'subtitle'
							πF.SetLineno(1409)
							µclasses = ßsubtitle.ToObject()
							// line 1410: self.in_document_title = len(self.body)+1
							πF.SetLineno(1410)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_document_title, πTemp003); πE != nil {
								continue
							}
							goto Label4
							// line 1411: elif isinstance(node.parent, nodes.section):
							πF.SetLineno(1411)
						Label3:
							// line 1412: classes = 'section-subtitle'
							πF.SetLineno(1412)
							µclasses = πg.NewStr("section-subtitle").ToObject()
							goto Label4
						Label4:
							// line 1413: self.body.append(self.starttag(node, 'p', '', CLASS=classes))
							πF.SetLineno(1413)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							πTemp006[1] = ßp.ToObject()
							πTemp006[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"CLASS", µclasses},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßvisit_subtitle.ToObject(), πTemp165); πE != nil {
						continue
					}
					// line 1415: def depart_subtitle(self, node):
					πF.SetLineno(1415)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp166 = πg.NewFunction(πg.NewCode("depart_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1416: self.body.append('</p>\n')
							πF.SetLineno(1416)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</p>\n").ToObject()
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
								goto Label1
							}
							goto Label2
							// line 1417: if isinstance(node.parent, nodes.document):
							πF.SetLineno(1417)
						Label1:
							// line 1418: self.subtitle = self.body[self.in_document_title:-1]
							πF.SetLineno(1418)
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
							// line 1419: self.in_document_title = 0
							πF.SetLineno(1419)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_document_title, πTemp002); πE != nil {
								continue
							}
							// line 1420: self.body_pre_docinfo.extend(self.body)
							πF.SetLineno(1420)
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
							// line 1421: self.html_subtitle.extend(self.body)
							πF.SetLineno(1421)
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
							// line 1422: del self.body[:]
							πF.SetLineno(1422)
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
					if πE = πClass.SetItem(πF, ßdepart_subtitle.ToObject(), πTemp166); πE != nil {
						continue
					}
					// line 1424: def visit_superscript(self, node):
					πF.SetLineno(1424)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp167 = πg.NewFunction(πg.NewCode("visit_superscript", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1425: self.body.append(self.starttag(node, 'sup', ''))
							πF.SetLineno(1425)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßsup.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_superscript.ToObject(), πTemp167); πE != nil {
						continue
					}
					// line 1427: def depart_superscript(self, node):
					πF.SetLineno(1427)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp168 = πg.NewFunction(πg.NewCode("depart_superscript", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1428: self.body.append('</sup>')
							πF.SetLineno(1428)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_superscript.ToObject(), πTemp168); πE != nil {
						continue
					}
					// line 1430: def visit_system_message(self, node):
					πF.SetLineno(1430)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp169 = πg.NewFunction(πg.NewCode("visit_system_message", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1431: self.body.append(self.starttag(node, 'div', CLASS='system-message'))
							πF.SetLineno(1431)
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
							// line 1432: self.body.append('<p class="system-message-title">')
							πF.SetLineno(1432)
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
							// line 1433: backref_text = ''
							πF.SetLineno(1433)
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
							// line 1434: if len(node['backrefs']):
							πF.SetLineno(1434)
						Label1:
							// line 1435: backrefs = node['backrefs']
							πF.SetLineno(1435)
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
							// line 1436: if len(backrefs) == 1:
							πF.SetLineno(1436)
						Label3:
							// line 1437: backref_text = ('; <em><a href="#%s">backlink</a></em>'
							πF.SetLineno(1437)
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
							// line 1440: i = 1
							πF.SetLineno(1440)
							µi = πg.NewInt(1).ToObject()
							// line 1441: backlinks = []
							πF.SetLineno(1441)
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
							// line 1443: backlinks.append('<a href="#%s">%s</a>' % (backref, i))
							πF.SetLineno(1443)
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
							// line 1444: i += 1
							πF.SetLineno(1444)
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
							// line 1445: backref_text = ('; <em>backlinks: %s</em>'
							πF.SetLineno(1445)
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
							// line 1447: if node.hasattr('line'):
							πF.SetLineno(1447)
						Label9:
							// line 1448: line = ', line %s' % node['line']
							πF.SetLineno(1448)
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
							// line 1450: line = ''
							πF.SetLineno(1450)
							µline = ß.ToObject()
							goto Label11
						Label11:
							// line 1451: self.body.append('System Message: %s/%s '
							πF.SetLineno(1451)
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
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("System Message: %s/%s (<span class=\"docutils literal\">%s</span>%s)%s</p>\n").ToObject(), πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_system_message.ToObject(), πTemp169); πE != nil {
						continue
					}
					// line 1456: def depart_system_message(self, node):
					πF.SetLineno(1456)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp170 = πg.NewFunction(πg.NewCode("depart_system_message", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1457: self.body.append('</div>\n')
							πF.SetLineno(1457)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_system_message.ToObject(), πTemp170); πE != nil {
						continue
					}
					// line 1459: def visit_table(self, node):
					πF.SetLineno(1459)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp171 = πg.NewFunction(πg.NewCode("visit_table", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µtag *πg.Object = πg.UnboundLocal
						_ = µtag
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 1460: atts = {}
							πF.SetLineno(1460)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							// line 1461: classes = [cls.strip(u' \t\n')
							πF.SetLineno(1461)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µcls *πg.Object = πg.UnboundLocal
								_ = µcls
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
										πTemp002[0] = πg.NewStr(",").ToObject()
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtable_style, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
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
											µcls = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1461: classes = [cls.strip(u' \t\n')
										πF.SetLineno(1461)
										πTemp002 = πF.MakeArgs(1)
										πTemp002[0] = πg.NewUnicode(" \t\n").ToObject()
										if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µcls, ßstrip, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µclasses = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µnode, ßalign.ToObject()); πE != nil {
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
							// line 1463: if 'align' in node:
							πF.SetLineno(1463)
						Label1:
							// line 1464: classes.append('align-%s' % node['align'])
							πF.SetLineno(1464)
							πTemp007 = πF.MakeArgs(1)
							πTemp005 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("align-%s").ToObject(), πTemp008); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µclasses, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µnode, ßwidth.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 1465: if 'width' in node:
							πF.SetLineno(1465)
						Label3:
							// line 1466: atts['style'] = 'width: %s' % node['width']
							πF.SetLineno(1466)
							πTemp005 = ßwidth.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("width: %s").ToObject(), πTemp008); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp008 = ßstyle.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp008, πTemp005); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 1467: tag = self.starttag(node, 'table', CLASS=' '.join(classes), **atts)
							πF.SetLineno(1467)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							πTemp007[1] = ßtable.ToObject()
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp009[0] = µclasses
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp010 = πg.KWArgs{
								{"CLASS", πTemp005},
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
							if πTemp005, πE = πg.Invoke(πF, πTemp002, πTemp007, nil, πTemp010, µatts); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µtag = πTemp005
							// line 1468: self.body.append(tag)
							πF.SetLineno(1468)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							πTemp007[0] = µtag
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_table.ToObject(), πTemp171); πE != nil {
						continue
					}
					// line 1470: def depart_table(self, node):
					πF.SetLineno(1470)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp172 = πg.NewFunction(πg.NewCode("depart_table", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1471: self.body.append('</table>\n')
							πF.SetLineno(1471)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</table>\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_table.ToObject(), πTemp172); πE != nil {
						continue
					}
					// line 1473: def visit_target(self, node):
					πF.SetLineno(1473)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp173 = πg.NewFunction(πg.NewCode("visit_target", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßrefuri.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp005).ToObject()
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßrefid.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp005).ToObject()
							πTemp002 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µnode, ßrefname.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp005).ToObject()
							πTemp002 = πTemp004
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 1474: if not ('refuri' in node or 'refid' in node
							πF.SetLineno(1474)
						Label2:
							// line 1476: self.body.append(self.starttag(node, 'span', '', CLASS='target'))
							πF.SetLineno(1476)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							πTemp007[1] = ßspan.ToObject()
							πTemp007[2] = ß.ToObject()
							πTemp008 = πg.KWArgs{
								{"CLASS", ßtarget.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp002
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
							// line 1477: self.context.append('</span>')
							πF.SetLineno(1477)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("</span>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
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
							// line 1479: self.context.append('')
							πF.SetLineno(1479)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_target.ToObject(), πTemp173); πE != nil {
						continue
					}
					// line 1481: def depart_target(self, node):
					πF.SetLineno(1481)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp174 = πg.NewFunction(πg.NewCode("depart_target", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1482: self.body.append(self.context.pop())
							πF.SetLineno(1482)
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
					if πE = πClass.SetItem(πF, ßdepart_target.ToObject(), πTemp174); πE != nil {
						continue
					}
					// line 1485: def visit_tbody(self, node):
					πF.SetLineno(1485)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp175 = πg.NewFunction(πg.NewCode("visit_tbody", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1486: self.body.append(self.starttag(node, 'tbody'))
							πF.SetLineno(1486)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßtbody.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_tbody.ToObject(), πTemp175); πE != nil {
						continue
					}
					// line 1488: def depart_tbody(self, node):
					πF.SetLineno(1488)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp176 = πg.NewFunction(πg.NewCode("depart_tbody", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1489: self.body.append('</tbody>\n')
							πF.SetLineno(1489)
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
					if πE = πClass.SetItem(πF, ßdepart_tbody.ToObject(), πTemp176); πE != nil {
						continue
					}
					// line 1491: def visit_term(self, node):
					πF.SetLineno(1491)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp177 = πg.NewFunction(πg.NewCode("visit_term", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1492: self.body.append(self.starttag(node, 'dt', ''))
							πF.SetLineno(1492)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdt.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_term.ToObject(), πTemp177); πE != nil {
						continue
					}
					// line 1494: def depart_term(self, node):
					πF.SetLineno(1494)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp178 = πg.NewFunction(πg.NewCode("depart_term", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1495: """
							πF.SetLineno(1495)
							// line 1499: pass
							πF.SetLineno(1499)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_term.ToObject(), πTemp178); πE != nil {
						continue
					}
					// line 1495: """
					πF.SetLineno(1495)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp179}, πg.NewStr("\n        Leave the end tag to `self.visit_definition()`, in case there's a\n        classifier.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp180, πE = πg.ResolveClass(πF, πClass, nil, ßdepart_term); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp180, ß__doc__, πTemp179); πE != nil {
						continue
					}
					// line 1501: def visit_tgroup(self, node):
					πF.SetLineno(1501)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp179 = πg.NewFunction(πg.NewCode("visit_tgroup", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1502: self.colspecs = []
							πF.SetLineno(1502)
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
							// line 1503: node.stubs = []
							πF.SetLineno(1503)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnode, ßstubs, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_tgroup.ToObject(), πTemp179); πE != nil {
						continue
					}
					// line 1505: def depart_tgroup(self, node):
					πF.SetLineno(1505)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp180 = πg.NewFunction(πg.NewCode("depart_tgroup", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1506: pass
							πF.SetLineno(1506)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_tgroup.ToObject(), πTemp180); πE != nil {
						continue
					}
					// line 1508: def visit_thead(self, node):
					πF.SetLineno(1508)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp181 = πg.NewFunction(πg.NewCode("visit_thead", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1509: self.body.append(self.starttag(node, 'thead'))
							πF.SetLineno(1509)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßthead.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_thead.ToObject(), πTemp181); πE != nil {
						continue
					}
					// line 1511: def depart_thead(self, node):
					πF.SetLineno(1511)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp182 = πg.NewFunction(πg.NewCode("depart_thead", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1512: self.body.append('</thead>\n')
							πF.SetLineno(1512)
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
					if πE = πClass.SetItem(πF, ßdepart_thead.ToObject(), πTemp182); πE != nil {
						continue
					}
					// line 1514: def visit_title(self, node):
					πF.SetLineno(1514)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp183 = πg.NewFunction(πg.NewCode("visit_title", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclose_tag *πg.Object = πg.UnboundLocal
						_ = µclose_tag
						var µh_level *πg.Object = πg.UnboundLocal
						_ = µh_level
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
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
						var πTemp009 *πg.Dict
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
							// line 1515: """Only 6 section levels are supported by HTML."""
							πF.SetLineno(1515)
							// line 1516: close_tag = '</p>\n'
							πF.SetLineno(1516)
							µclose_tag = πg.NewStr("</p>\n").ToObject()
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßAdmonition, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtable, nil); πE != nil {
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
								goto Label4
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
								goto Label5
							}
							goto Label6
							// line 1517: if isinstance(node.parent, nodes.topic):
							πF.SetLineno(1517)
						Label1:
							// line 1518: self.body.append(
							πF.SetLineno(1518)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßp.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", πg.NewStr("topic-title").ToObject()},
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
							goto Label7
							// line 1520: elif isinstance(node.parent, nodes.sidebar):
							πF.SetLineno(1520)
						Label2:
							// line 1521: self.body.append(
							πF.SetLineno(1521)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßp.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", πg.NewStr("sidebar-title").ToObject()},
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
							goto Label7
							// line 1523: elif isinstance(node.parent, nodes.Admonition):
							πF.SetLineno(1523)
						Label3:
							// line 1524: self.body.append(
							πF.SetLineno(1524)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßp.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", πg.NewStr("admonition-title").ToObject()},
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
							goto Label7
							// line 1526: elif isinstance(node.parent, nodes.table):
							πF.SetLineno(1526)
						Label4:
							// line 1527: self.body.append(
							πF.SetLineno(1527)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßcaption.ToObject()
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
							// line 1529: close_tag = '</caption>\n'
							πF.SetLineno(1529)
							µclose_tag = πg.NewStr("</caption>\n").ToObject()
							goto Label7
							// line 1530: elif isinstance(node.parent, nodes.document):
							πF.SetLineno(1530)
						Label5:
							// line 1531: self.body.append(self.starttag(node, 'h1', '', CLASS='title'))
							πF.SetLineno(1531)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßh1.ToObject()
							πTemp005[2] = ß.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßtitle.ToObject()},
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
							// line 1532: close_tag = '</h1>\n'
							πF.SetLineno(1532)
							µclose_tag = πg.NewStr("</h1>\n").ToObject()
							// line 1533: self.in_document_title = len(self.body)
							πF.SetLineno(1533)
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
							goto Label7
						Label6:
							// line 1535: assert isinstance(node.parent, nodes.section)
							πF.SetLineno(1535)
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
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 1536: h_level = self.section_level + self.initial_header_level - 1
							πF.SetLineno(1536)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßinitial_header_level, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µh_level = πTemp002
							// line 1537: atts = {}
							πF.SetLineno(1537)
							πTemp009 = πg.NewDict()
							πTemp002 = πTemp009.ToObject()
							µatts = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GE(πF, πTemp008, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label8
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßsubtitle, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp007
						Label8:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 1538: if (len(node.parent) >= 2 and
							πF.SetLineno(1538)
						Label9:
							// line 1540: atts['CLASS'] = 'with-subtitle'
							πF.SetLineno(1540)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("with-subtitle").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp003 = ßCLASS.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label10
						Label10:
							// line 1541: self.body.append(
							πF.SetLineno(1541)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πE = πg.CheckLocal(πF, µh_level, "h_level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("h%s").ToObject(), µh_level); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp005, nil, nil, µatts); πE != nil {
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
							// line 1543: atts = {}
							πF.SetLineno(1543)
							πTemp009 = πg.NewDict()
							πTemp002 = πTemp009.ToObject()
							µatts = πTemp002
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
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
								goto Label11
							}
							goto Label12
							// line 1544: if node.hasattr('refid'):
							πF.SetLineno(1544)
						Label11:
							// line 1545: atts['class'] = 'toc-backref'
							πF.SetLineno(1545)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("toc-backref").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp003 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 1546: atts['href'] = '#' + node['refid']
							πF.SetLineno(1546)
							πTemp003 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("#").ToObject(), πTemp007); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πTemp007 = ßhref.ToObject()
							if πE = πg.SetItem(πF, µatts, πTemp007, πTemp003); πE != nil {
								continue
							}
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µatts); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 1547: if atts:
							πF.SetLineno(1547)
						Label13:
							// line 1548: self.body.append(self.starttag({}, 'a', '', **atts))
							πF.SetLineno(1548)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(3)
							πTemp009 = πg.NewDict()
							πTemp002 = πTemp009.ToObject()
							πTemp005[0] = πTemp002
							πTemp005[1] = ßa.ToObject()
							πTemp005[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp005, nil, nil, µatts); πE != nil {
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
							// line 1549: close_tag = '</a></h%s>\n' % (h_level)
							πF.SetLineno(1549)
							if πE = πg.CheckLocal(πF, µh_level, "h_level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("</a></h%s>\n").ToObject(), µh_level); πE != nil {
								continue
							}
							µclose_tag = πTemp002
							goto Label15
						Label14:
							// line 1551: close_tag = '</h%s>\n' % (h_level)
							πF.SetLineno(1551)
							if πE = πg.CheckLocal(πF, µh_level, "h_level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("</h%s>\n").ToObject(), µh_level); πE != nil {
								continue
							}
							µclose_tag = πTemp002
							goto Label15
						Label15:
							goto Label7
						Label7:
							// line 1552: self.context.append(close_tag)
							πF.SetLineno(1552)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclose_tag, "close_tag"); πE != nil {
								continue
							}
							πTemp001[0] = µclose_tag
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
					if πE = πClass.SetItem(πF, ßvisit_title.ToObject(), πTemp183); πE != nil {
						continue
					}
					// line 1515: """Only 6 section levels are supported by HTML."""
					πF.SetLineno(1515)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp184}, πg.NewStr("Only 6 section levels are supported by HTML.").ToObject()); πE != nil {
						continue
					}
					if πTemp185, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_title); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp185, ß__doc__, πTemp184); πE != nil {
						continue
					}
					// line 1554: def depart_title(self, node):
					πF.SetLineno(1554)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp184 = πg.NewFunction(πg.NewCode("depart_title", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1555: self.body.append(self.context.pop())
							πF.SetLineno(1555)
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
							// line 1556: if self.in_document_title:
							πF.SetLineno(1556)
						Label1:
							// line 1557: self.title = self.body[self.in_document_title:-1]
							πF.SetLineno(1557)
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
							if πE = πg.SetAttr(πF, µself, ßtitle, πTemp002); πE != nil {
								continue
							}
							// line 1558: self.in_document_title = 0
							πF.SetLineno(1558)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_document_title, πTemp002); πE != nil {
								continue
							}
							// line 1559: self.body_pre_docinfo.extend(self.body)
							πF.SetLineno(1559)
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
							// line 1560: self.html_title.extend(self.body)
							πF.SetLineno(1560)
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhtml_title, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1561: del self.body[:]
							πF.SetLineno(1561)
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
					if πE = πClass.SetItem(πF, ßdepart_title.ToObject(), πTemp184); πE != nil {
						continue
					}
					// line 1563: def visit_title_reference(self, node):
					πF.SetLineno(1563)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp185 = πg.NewFunction(πg.NewCode("visit_title_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1564: self.body.append(self.starttag(node, 'cite', ''))
							πF.SetLineno(1564)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßcite.ToObject()
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
					if πE = πClass.SetItem(πF, ßvisit_title_reference.ToObject(), πTemp185); πE != nil {
						continue
					}
					// line 1566: def depart_title_reference(self, node):
					πF.SetLineno(1566)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp186 = πg.NewFunction(πg.NewCode("depart_title_reference", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1567: self.body.append('</cite>')
							πF.SetLineno(1567)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</cite>").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_title_reference.ToObject(), πTemp186); πE != nil {
						continue
					}
					// line 1570: def visit_topic(self, node):
					πF.SetLineno(1570)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp187 = πg.NewFunction(πg.NewCode("visit_topic", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1571: self.body.append(self.starttag(node, 'div', CLASS='topic'))
							πF.SetLineno(1571)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßdiv.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßtopic.ToObject()},
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
							// line 1572: self.topic_classes = node['classes']
							πF.SetLineno(1572)
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtopic_classes, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_topic.ToObject(), πTemp187); πE != nil {
						continue
					}
					// line 1576: def depart_topic(self, node):
					πF.SetLineno(1576)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp188 = πg.NewFunction(πg.NewCode("depart_topic", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1577: self.body.append('</div>\n')
							πF.SetLineno(1577)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
							// line 1578: self.topic_classes = []
							πF.SetLineno(1578)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtopic_classes, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_topic.ToObject(), πTemp188); πE != nil {
						continue
					}
					// line 1581: def visit_transition(self, node):
					πF.SetLineno(1581)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp189 = πg.NewFunction(πg.NewCode("visit_transition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1582: self.body.append(self.emptytag(node, 'hr', CLASS='docutils'))
							πF.SetLineno(1582)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßhr.ToObject()
							πTemp003 = πg.KWArgs{
								{"CLASS", ßdocutils.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßemptytag, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_transition.ToObject(), πTemp189); πE != nil {
						continue
					}
					// line 1584: def depart_transition(self, node):
					πF.SetLineno(1584)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp190 = πg.NewFunction(πg.NewCode("depart_transition", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1585: pass
							πF.SetLineno(1585)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_transition.ToObject(), πTemp190); πE != nil {
						continue
					}
					// line 1587: def visit_version(self, node):
					πF.SetLineno(1587)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp191 = πg.NewFunction(πg.NewCode("visit_version", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 1588: self.visit_docinfo_item(node, 'version', meta=False)
							πF.SetLineno(1588)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßversion.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_version.ToObject(), πTemp191); πE != nil {
						continue
					}
					// line 1590: def depart_version(self, node):
					πF.SetLineno(1590)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp192 = πg.NewFunction(πg.NewCode("depart_version", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1591: self.depart_docinfo_item()
							πF.SetLineno(1591)
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
					if πE = πClass.SetItem(πF, ßdepart_version.ToObject(), πTemp192); πE != nil {
						continue
					}
					// line 1593: def unimplemented_visit(self, node):
					πF.SetLineno(1593)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp193 = πg.NewFunction(πg.NewCode("unimplemented_visit", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1594: raise NotImplementedError('visiting unimplemented node type: %s'
							πF.SetLineno(1594)
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
					if πE = πClass.SetItem(πF, ßunimplemented_visit.ToObject(), πTemp193); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("HTMLTranslator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHTMLTranslator.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 1598: class SimpleListChecker(nodes.GenericNodeVisitor):
			πF.SetLineno(1598)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßGenericNodeVisitor, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SimpleListChecker", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1600: """
					πF.SetLineno(1600)
					// line 1600: """
					πF.SetLineno(1600)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise `nodes.NodeFound` if non-simple list item is encountered.\n\n    Here \"simple\" means a list item containing nothing other than a single\n    paragraph, a simple list, or a paragraph followed by a simple list.\n\n    This version also checks for simple field lists and docinfo.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1609: def default_visit(self, node):
					πF.SetLineno(1609)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("default_visit", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1610: raise nodes.NodeFound
							πF.SetLineno(1610)
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
					if πE = πClass.SetItem(πF, ßdefault_visit.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1612: def visit_list_item(self, node):
					πF.SetLineno(1612)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("visit_list_item", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µchildren *πg.Object = πg.UnboundLocal
						_ = µchildren
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
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
							// line 1613: children = [child for child in node.children
							πF.SetLineno(1613)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µchild *πg.Object = πg.UnboundLocal
								_ = µchild
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
										if πTemp002, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
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
											µchild = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										πTemp005[0] = µchild
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
										// line 1613: children = [child for child in node.children
										πF.SetLineno(1613)
									Label4:
										// line 1613: children = [child for child in node.children
										πF.SetLineno(1613)
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µchild, nil
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µchildren = πTemp001
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001 = µchildren
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µchildren, πTemp004); πE != nil {
								continue
							}
							πTemp006[0] = πTemp007
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp007
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001 = πTemp007
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp006 = πF.MakeArgs(2)
							if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µchildren, πTemp007); πE != nil {
								continue
							}
							πTemp006[0] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßbullet_list, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004 = πTemp009
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label2
							}
							πTemp006 = πF.MakeArgs(2)
							if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µchildren, πTemp007); πE != nil {
								continue
							}
							πTemp006[0] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßenumerated_list, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004 = πTemp009
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label2
							}
							πTemp006 = πF.MakeArgs(2)
							if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µchildren, πTemp007); πE != nil {
								continue
							}
							πTemp006[0] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßfield_list, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004 = πTemp009
						Label2:
							πTemp001 = πTemp004
						Label1:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1615: if (children and isinstance(children[0], nodes.paragraph)
							πF.SetLineno(1615)
						Label3:
							// line 1619: children.pop()
							πF.SetLineno(1619)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µchildren, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label4
						Label4:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp006[0] = µchildren
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.LE(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 1620: if len(children) <= 1:
							πF.SetLineno(1620)
						Label5:
							// line 1621: return
							πF.SetLineno(1621)
							πR = πg.None
							continue
							goto Label7
						Label6:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßNodeFound, nil); πE != nil {
								continue
							}
							// line 1623: raise nodes.NodeFound
							πF.SetLineno(1623)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßvisit_list_item.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1625: def pass_node(self, node):
					πF.SetLineno(1625)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("pass_node", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1626: pass
							πF.SetLineno(1626)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpass_node.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1628: def ignore_node(self, node):
					πF.SetLineno(1628)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("ignore_node", "/usr/lib/python2.7/site-packages/docutils/writers/_html_base.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1630: raise nodes.SkipNode
							πF.SetLineno(1630)
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
					if πE = πClass.SetItem(πF, ßignore_node.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1633: visit_Text = ignore_node
					πF.SetLineno(1633)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_Text.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1634: visit_paragraph = ignore_node
					πF.SetLineno(1634)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_paragraph.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1637: visit_bullet_list = pass_node
					πF.SetLineno(1637)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_bullet_list.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1638: visit_enumerated_list = pass_node
					πF.SetLineno(1638)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_enumerated_list.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1639: visit_docinfo = pass_node
					πF.SetLineno(1639)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_docinfo.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1642: visit_author = ignore_node
					πF.SetLineno(1642)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_author.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1643: visit_authors = visit_list_item
					πF.SetLineno(1643)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_list_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_authors.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1644: visit_address = visit_list_item
					πF.SetLineno(1644)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_list_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_address.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1645: visit_contact = pass_node
					πF.SetLineno(1645)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_contact.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1646: visit_copyright = ignore_node
					πF.SetLineno(1646)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_copyright.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1647: visit_date = ignore_node
					πF.SetLineno(1647)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_date.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1648: visit_organization = ignore_node
					πF.SetLineno(1648)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_organization.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1649: visit_status = ignore_node
					πF.SetLineno(1649)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_status.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1650: visit_version = visit_list_item
					πF.SetLineno(1650)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_list_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_version.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1653: visit_definition_list = pass_node
					πF.SetLineno(1653)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_definition_list.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1654: visit_definition_list_item = pass_node
					πF.SetLineno(1654)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_definition_list_item.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1655: visit_term = ignore_node
					πF.SetLineno(1655)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_term.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1656: visit_classifier = pass_node
					πF.SetLineno(1656)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_classifier.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1657: visit_definition = visit_list_item
					πF.SetLineno(1657)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_list_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_definition.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1660: visit_field_list = pass_node
					πF.SetLineno(1660)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_field_list.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1661: visit_field = pass_node
					πF.SetLineno(1661)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßpass_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_field.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1663: visit_field_body = visit_list_item
					πF.SetLineno(1663)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_list_item); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_field_body.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1664: visit_field_name = ignore_node
					πF.SetLineno(1664)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_field_name.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1667: visit_comment = ignore_node
					πF.SetLineno(1667)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_comment.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1668: visit_substitution_definition = ignore_node
					πF.SetLineno(1668)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_substitution_definition.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1669: visit_target = ignore_node
					πF.SetLineno(1669)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_target.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1670: visit_pending = ignore_node
					πF.SetLineno(1670)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_pending.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp007 == nil {
				πTemp007 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp007.Call(πF, []*πg.Object{πg.NewStr("SimpleListChecker").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSimpleListChecker.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_html_base", Code)
}
