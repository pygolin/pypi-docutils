package nodes

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/collections"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/unicodedata"
	_ "github.com/pygolin/stdlib/pkg/warnings"
	_ "github.com/pygolin/stdlib/pkg/xml"
	_ "github.com/pygolin/stdlib/pkg/xml/dom"
	_ "github.com/pygolin/stdlib/pkg/xml/dom/minidom"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAdmonition := πg.InternStr("Admonition")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBackLinkable := πg.InternStr("BackLinkable")
		ßBibliographic := πg.InternStr("Bibliographic")
		ßBody := πg.InternStr("Body")
		ßCounter := πg.InternStr("Counter")
		ßDecorative := πg.InternStr("Decorative")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßDocument := πg.InternStr("Document")
		ßElement := πg.InternStr("Element")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßFixedTextElement := πg.InternStr("FixedTextElement")
		ßFutureWarning := πg.InternStr("FutureWarning")
		ßGeneral := πg.InternStr("General")
		ßGenericNodeVisitor := πg.InternStr("GenericNodeVisitor")
		ßIndexError := πg.InternStr("IndexError")
		ßInline := πg.InternStr("Inline")
		ßInvisible := πg.InternStr("Invisible")
		ßLabeled := πg.InternStr("Labeled")
		ßNFKD := πg.InternStr("NFKD")
		ßNode := πg.InternStr("Node")
		ßNodeFound := πg.InternStr("NodeFound")
		ßNodeVisitor := πg.InternStr("NodeVisitor")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßPart := πg.InternStr("Part")
		ßPreBibliographic := πg.InternStr("PreBibliographic")
		ßReferential := πg.InternStr("Referential")
		ßResolvable := πg.InternStr("Resolvable")
		ßRoot := πg.InternStr("Root")
		ßSequential := πg.InternStr("Sequential")
		ßSkipChildren := πg.InternStr("SkipChildren")
		ßSkipDeparture := πg.InternStr("SkipDeparture")
		ßSkipNode := πg.InternStr("SkipNode")
		ßSkipSiblings := πg.InternStr("SkipSiblings")
		ßSparseNodeVisitor := πg.InternStr("SparseNodeVisitor")
		ßSpecial := πg.InternStr("Special")
		ßStopIteration := πg.InternStr("StopIteration")
		ßStopTraversal := πg.InternStr("StopTraversal")
		ßStructural := πg.InternStr("Structural")
		ßTargetable := πg.InternStr("Targetable")
		ßText := πg.InternStr("Text")
		ßTextElement := πg.InternStr("TextElement")
		ßTitular := πg.InternStr("Titular")
		ßTransformer := πg.InternStr("Transformer")
		ßTreeCopyVisitor := πg.InternStr("TreeCopyVisitor")
		ßTreePruningException := πg.InternStr("TreePruningException")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ß__add__ := πg.InternStr("__add__")
		ß__bool__ := πg.InternStr("__bool__")
		ß__class__ := πg.InternStr("__class__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__init__ := πg.InternStr("__init__")
		ß__len__ := πg.InternStr("__len__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__nonzero__ := πg.InternStr("__nonzero__")
		ß__radd__ := πg.InternStr("__radd__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__reversed__ := πg.InternStr("__reversed__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__str__ := πg.InternStr("__str__")
		ß__unicode__ := πg.InternStr("__unicode__")
		ß_add_node_class_names := πg.InternStr("_add_node_class_names")
		ß_all_traverse := πg.InternStr("_all_traverse")
		ß_call_default_departure := πg.InternStr("_call_default_departure")
		ß_call_default_visit := πg.InternStr("_call_default_visit")
		ß_dom_node := πg.InternStr("_dom_node")
		ß_fast_traverse := πg.InternStr("_fast_traverse")
		ß_non_id_at_ends := πg.InternStr("_non_id_at_ends")
		ß_non_id_chars := πg.InternStr("_non_id_chars")
		ß_non_id_translate := πg.InternStr("_non_id_translate")
		ß_non_id_translate_digraphs := πg.InternStr("_non_id_translate_digraphs")
		ß_nop := πg.InternStr("_nop")
		ß_traversal_list := πg.InternStr("_traversal_list")
		ß_traverse := πg.InternStr("_traverse")
		ß_warning_decorator := πg.InternStr("_warning_decorator")
		ßabbreviation := πg.InternStr("abbreviation")
		ßacronym := πg.InternStr("acronym")
		ßadd_backref := πg.InternStr("add_backref")
		ßadd_pending := πg.InternStr("add_pending")
		ßaddress := πg.InternStr("address")
		ßadmonition := πg.InternStr("admonition")
		ßalt := πg.InternStr("alt")
		ßappend := πg.InternStr("append")
		ßappendChild := πg.InternStr("appendChild")
		ßappend_attr_list := πg.InternStr("append_attr_list")
		ßascii := πg.InternStr("ascii")
		ßasdom := πg.InternStr("asdom")
		ßastext := πg.InternStr("astext")
		ßattention := πg.InternStr("attention")
		ßattlist := πg.InternStr("attlist")
		ßattributes := πg.InternStr("attributes")
		ßattribution := πg.InternStr("attribution")
		ßauthor := πg.InternStr("author")
		ßauthors := πg.InternStr("authors")
		ßauto_id_prefix := πg.InternStr("auto_id_prefix")
		ßautofootnote_refs := πg.InternStr("autofootnote_refs")
		ßautofootnote_start := πg.InternStr("autofootnote_start")
		ßautofootnotes := πg.InternStr("autofootnotes")
		ßbackrefs := πg.InternStr("backrefs")
		ßbackslashreplace := πg.InternStr("backslashreplace")
		ßbasestring := πg.InternStr("basestring")
		ßbasic_attributes := πg.InternStr("basic_attributes")
		ßblock_quote := πg.InternStr("block_quote")
		ßbullet_list := πg.InternStr("bullet_list")
		ßbytes := πg.InternStr("bytes")
		ßcaption := πg.InternStr("caption")
		ßcaution := πg.InternStr("caution")
		ßchild_text_separator := πg.InternStr("child_text_separator")
		ßchildren := πg.InternStr("children")
		ßcitation := πg.InternStr("citation")
		ßcitation_reference := πg.InternStr("citation_reference")
		ßcitation_refs := πg.InternStr("citation_refs")
		ßcitations := πg.InternStr("citations")
		ßclasses := πg.InternStr("classes")
		ßclassifier := πg.InternStr("classifier")
		ßclassmethod := πg.InternStr("classmethod")
		ßclear := πg.InternStr("clear")
		ßcoerce_append_attr_list := πg.InternStr("coerce_append_attr_list")
		ßcolspec := πg.InternStr("colspec")
		ßcomment := πg.InternStr("comment")
		ßcompile := πg.InternStr("compile")
		ßcompound := πg.InternStr("compound")
		ßcontact := πg.InternStr("contact")
		ßcontainer := πg.InternStr("container")
		ßcopy := πg.InternStr("copy")
		ßcopy_attr_coerce := πg.InternStr("copy_attr_coerce")
		ßcopy_attr_concatenate := πg.InternStr("copy_attr_concatenate")
		ßcopy_attr_consistent := πg.InternStr("copy_attr_consistent")
		ßcopy_attr_convert := πg.InternStr("copy_attr_convert")
		ßcopyright := πg.InternStr("copyright")
		ßcount := πg.InternStr("count")
		ßcreateElement := πg.InternStr("createElement")
		ßcreateTextNode := πg.InternStr("createTextNode")
		ßcurrent_line := πg.InternStr("current_line")
		ßcurrent_source := πg.InternStr("current_source")
		ßdanger := πg.InternStr("danger")
		ßdate := πg.InternStr("date")
		ßdebug := πg.InternStr("debug")
		ßdecode := πg.InternStr("decode")
		ßdecoration := πg.InternStr("decoration")
		ßdeepcopy := πg.InternStr("deepcopy")
		ßdefault_departure := πg.InternStr("default_departure")
		ßdefault_visit := πg.InternStr("default_visit")
		ßdefinition := πg.InternStr("definition")
		ßdefinition_list := πg.InternStr("definition_list")
		ßdefinition_list_item := πg.InternStr("definition_list_item")
		ßdelattr := πg.InternStr("delattr")
		ßdelimiter := πg.InternStr("delimiter")
		ßdepart_ := πg.InternStr("depart_")
		ßdescription := πg.InternStr("description")
		ßdetails := πg.InternStr("details")
		ßdispatch_departure := πg.InternStr("dispatch_departure")
		ßdispatch_visit := πg.InternStr("dispatch_visit")
		ßdocinfo := πg.InternStr("docinfo")
		ßdoctest_block := πg.InternStr("doctest_block")
		ßdocument := πg.InternStr("document")
		ßdone := πg.InternStr("done")
		ßdupname := πg.InternStr("dupname")
		ßdupnames := πg.InternStr("dupnames")
		ßemphasis := πg.InternStr("emphasis")
		ßemptytag := πg.InternStr("emptytag")
		ßencode := πg.InternStr("encode")
		ßendswith := πg.InternStr("endswith")
		ßendtag := πg.InternStr("endtag")
		ßensure_str := πg.InternStr("ensure_str")
		ßentry := πg.InternStr("entry")
		ßenumerated_list := πg.InternStr("enumerated_list")
		ßerror := πg.InternStr("error")
		ßexpect_referenced_by_id := πg.InternStr("expect_referenced_by_id")
		ßexpect_referenced_by_name := πg.InternStr("expect_referenced_by_name")
		ßextend := πg.InternStr("extend")
		ßfield := πg.InternStr("field")
		ßfield_body := πg.InternStr("field_body")
		ßfield_list := πg.InternStr("field_list")
		ßfield_name := πg.InternStr("field_name")
		ßfigure := πg.InternStr("figure")
		ßfilter := πg.InternStr("filter")
		ßfirst_child_matching_class := πg.InternStr("first_child_matching_class")
		ßfirst_child_not_matching_class := πg.InternStr("first_child_not_matching_class")
		ßfooter := πg.InternStr("footer")
		ßfootnote := πg.InternStr("footnote")
		ßfootnote_reference := πg.InternStr("footnote_reference")
		ßfootnote_refs := πg.InternStr("footnote_refs")
		ßfootnotes := πg.InternStr("footnotes")
		ßfully_normalize_name := πg.InternStr("fully_normalize_name")
		ßgenerated := πg.InternStr("generated")
		ßget := πg.InternStr("get")
		ßget_decoration := πg.InternStr("get_decoration")
		ßget_footer := πg.InternStr("get_footer")
		ßget_header := πg.InternStr("get_header")
		ßget_language := πg.InternStr("get_language")
		ßget_language_code := πg.InternStr("get_language_code")
		ßget_tree_copy := πg.InternStr("get_tree_copy")
		ßgetattr := πg.InternStr("getattr")
		ßhas_key := πg.InternStr("has_key")
		ßhas_name := πg.InternStr("has_name")
		ßhasattr := πg.InternStr("hasattr")
		ßheader := πg.InternStr("header")
		ßhint := πg.InternStr("hint")
		ßid_counter := πg.InternStr("id_counter")
		ßid_prefix := πg.InternStr("id_prefix")
		ßids := πg.InternStr("ids")
		ßignore := πg.InternStr("ignore")
		ßimage := πg.InternStr("image")
		ßimportant := πg.InternStr("important")
		ßindent := πg.InternStr("indent")
		ßindex := πg.InternStr("index")
		ßindirect_reference_name := πg.InternStr("indirect_reference_name")
		ßindirect_targets := πg.InternStr("indirect_targets")
		ßinfo := πg.InternStr("info")
		ßinline := πg.InternStr("inline")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßis_not_default := πg.InternStr("is_not_default")
		ßis_not_known_attribute := πg.InternStr("is_not_known_attribute")
		ßis_not_list_attribute := πg.InternStr("is_not_list_attribute")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßknown_attributes := πg.InternStr("known_attributes")
		ßlabel := πg.InternStr("label")
		ßlegend := πg.InternStr("legend")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßline := πg.InternStr("line")
		ßline_block := πg.InternStr("line_block")
		ßlist := πg.InternStr("list")
		ßlist_attributes := πg.InternStr("list_attributes")
		ßlist_item := πg.InternStr("list_item")
		ßliteral := πg.InternStr("literal")
		ßliteral_block := πg.InternStr("literal_block")
		ßlocal_attributes := πg.InternStr("local_attributes")
		ßlower := πg.InternStr("lower")
		ßlstrip := πg.InternStr("lstrip")
		ßmake_id := πg.InternStr("make_id")
		ßmath := πg.InternStr("math")
		ßmath_block := πg.InternStr("math_block")
		ßmaxsize := πg.InternStr("maxsize")
		ßmin := πg.InternStr("min")
		ßnameids := πg.InternStr("nameids")
		ßnames := πg.InternStr("names")
		ßnametypes := πg.InternStr("nametypes")
		ßnext := πg.InternStr("next")
		ßnext_node := πg.InternStr("next_node")
		ßnode_class_names := πg.InternStr("node_class_names")
		ßnon_default_attributes := πg.InternStr("non_default_attributes")
		ßnormalize := πg.InternStr("normalize")
		ßnote := πg.InternStr("note")
		ßnote_anonymous_target := πg.InternStr("note_anonymous_target")
		ßnote_autofootnote := πg.InternStr("note_autofootnote")
		ßnote_autofootnote_ref := πg.InternStr("note_autofootnote_ref")
		ßnote_citation := πg.InternStr("note_citation")
		ßnote_citation_ref := πg.InternStr("note_citation_ref")
		ßnote_explicit_target := πg.InternStr("note_explicit_target")
		ßnote_footnote := πg.InternStr("note_footnote")
		ßnote_footnote_ref := πg.InternStr("note_footnote_ref")
		ßnote_implicit_target := πg.InternStr("note_implicit_target")
		ßnote_indirect_target := πg.InternStr("note_indirect_target")
		ßnote_parse_message := πg.InternStr("note_parse_message")
		ßnote_pending := πg.InternStr("note_pending")
		ßnote_referenced_by := πg.InternStr("note_referenced_by")
		ßnote_refid := πg.InternStr("note_refid")
		ßnote_refname := πg.InternStr("note_refname")
		ßnote_source := πg.InternStr("note_source")
		ßnote_substitution_def := πg.InternStr("note_substitution_def")
		ßnote_substitution_ref := πg.InternStr("note_substitution_ref")
		ßnote_symbol_footnote := πg.InternStr("note_symbol_footnote")
		ßnote_symbol_footnote_ref := πg.InternStr("note_symbol_footnote_ref")
		ßnote_transform_message := πg.InternStr("note_transform_message")
		ßobject := πg.InternStr("object")
		ßoption := πg.InternStr("option")
		ßoption_argument := πg.InternStr("option_argument")
		ßoption_group := πg.InternStr("option_group")
		ßoption_list := πg.InternStr("option_list")
		ßoption_list_item := πg.InternStr("option_list_item")
		ßoption_string := πg.InternStr("option_string")
		ßoptional := πg.InternStr("optional")
		ßorganization := πg.InternStr("organization")
		ßos := πg.InternStr("os")
		ßparagraph := πg.InternStr("paragraph")
		ßparent := πg.InternStr("parent")
		ßparent_stack := πg.InternStr("parent_stack")
		ßparse_messages := πg.InternStr("parse_messages")
		ßpending := πg.InternStr("pending")
		ßpformat := πg.InternStr("pformat")
		ßpop := πg.InternStr("pop")
		ßpreserve := πg.InternStr("preserve")
		ßprint := πg.InternStr("print")
		ßproblematic := πg.InternStr("problematic")
		ßpseudo_quoteattr := πg.InternStr("pseudo_quoteattr")
		ßrange := πg.InternStr("range")
		ßraw := πg.InternStr("raw")
		ßraw_unicode_escape := πg.InternStr("raw_unicode_escape")
		ßrawsource := πg.InternStr("rawsource")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßreferenced := πg.InternStr("referenced")
		ßrefid := πg.InternStr("refid")
		ßrefids := πg.InternStr("refids")
		ßrefname := πg.InternStr("refname")
		ßrefnames := πg.InternStr("refnames")
		ßrefuri := πg.InternStr("refuri")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßreplace_attr := πg.InternStr("replace_attr")
		ßreplace_self := πg.InternStr("replace_self")
		ßreporter := πg.InternStr("reporter")
		ßreprunicode := πg.InternStr("reprunicode")
		ßresolved := πg.InternStr("resolved")
		ßreverse := πg.InternStr("reverse")
		ßrevision := πg.InternStr("revision")
		ßrow := πg.InternStr("row")
		ßrstrip := πg.InternStr("rstrip")
		ßrubric := πg.InternStr("rubric")
		ßsection := πg.InternStr("section")
		ßserial_escape := πg.InternStr("serial_escape")
		ßsetAttribute := πg.InternStr("setAttribute")
		ßset_class := πg.InternStr("set_class")
		ßset_duplicate_name_id := πg.InternStr("set_duplicate_name_id")
		ßset_id := πg.InternStr("set_id")
		ßset_name_id_map := πg.InternStr("set_name_id_map")
		ßsetattr := πg.InternStr("setattr")
		ßsetdefault := πg.InternStr("setdefault")
		ßsettings := πg.InternStr("settings")
		ßsetup_child := πg.InternStr("setup_child")
		ßsevere := πg.InternStr("severe")
		ßshortrepr := πg.InternStr("shortrepr")
		ßsidebar := πg.InternStr("sidebar")
		ßslice := πg.InternStr("slice")
		ßsorted := πg.InternStr("sorted")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßsplitlines := πg.InternStr("splitlines")
		ßstart := πg.InternStr("start")
		ßstartswith := πg.InternStr("startswith")
		ßstarttag := πg.InternStr("starttag")
		ßstatus := πg.InternStr("status")
		ßstep := πg.InternStr("step")
		ßstop := πg.InternStr("stop")
		ßstr := πg.InternStr("str")
		ßstrict_visitor := πg.InternStr("strict_visitor")
		ßstrong := πg.InternStr("strong")
		ßsub := πg.InternStr("sub")
		ßsubscript := πg.InternStr("subscript")
		ßsubstitution_definition := πg.InternStr("substitution_definition")
		ßsubstitution_defs := πg.InternStr("substitution_defs")
		ßsubstitution_names := πg.InternStr("substitution_names")
		ßsubstitution_reference := πg.InternStr("substitution_reference")
		ßsubtitle := πg.InternStr("subtitle")
		ßsuperscript := πg.InternStr("superscript")
		ßsymbol_footnote_refs := πg.InternStr("symbol_footnote_refs")
		ßsymbol_footnote_start := πg.InternStr("symbol_footnote_start")
		ßsymbol_footnotes := πg.InternStr("symbol_footnotes")
		ßsys := πg.InternStr("sys")
		ßsystem_message := πg.InternStr("system_message")
		ßtable := πg.InternStr("table")
		ßtagname := πg.InternStr("tagname")
		ßtarget := πg.InternStr("target")
		ßtbody := πg.InternStr("tbody")
		ßterm := πg.InternStr("term")
		ßtgroup := πg.InternStr("tgroup")
		ßthead := πg.InternStr("thead")
		ßtip := πg.InternStr("tip")
		ßtitle := πg.InternStr("title")
		ßtitle_reference := πg.InternStr("title_reference")
		ßtopic := πg.InternStr("topic")
		ßtransform := πg.InternStr("transform")
		ßtransform_messages := πg.InternStr("transform_messages")
		ßtransformer := πg.InternStr("transformer")
		ßtransforms := πg.InternStr("transforms")
		ßtransition := πg.InternStr("transition")
		ßtranslate := πg.InternStr("translate")
		ßtraverse := πg.InternStr("traverse")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßunescape := πg.InternStr("unescape")
		ßunicode := πg.InternStr("unicode")
		ßunicodedata := πg.InternStr("unicodedata")
		ßunknown_departure := πg.InternStr("unknown_departure")
		ßunknown_visit := πg.InternStr("unknown_visit")
		ßupdate_all_atts := πg.InternStr("update_all_atts")
		ßupdate_all_atts_coercion := πg.InternStr("update_all_atts_coercion")
		ßupdate_all_atts_concatenating := πg.InternStr("update_all_atts_concatenating")
		ßupdate_all_atts_consistantly := πg.InternStr("update_all_atts_consistantly")
		ßupdate_all_atts_convert := πg.InternStr("update_all_atts_convert")
		ßupdate_basic_atts := πg.InternStr("update_basic_atts")
		ßversion := πg.InternStr("version")
		ßversion_info := πg.InternStr("version_info")
		ßvisit_ := πg.InternStr("visit_")
		ßwalk := πg.InternStr("walk")
		ßwalkabout := πg.InternStr("walkabout")
		ßwarn := πg.InternStr("warn")
		ßwarning := πg.InternStr("warning")
		ßwarnings := πg.InternStr("warnings")
		ßwhitespace_normalize_name := πg.InternStr("whitespace_normalize_name")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDocutils document tree element class library.\n\nClasses in CamelCase are abstract base classes or auxiliary classes. The one\nexception is `Text`, for a text (PCDATA) node; uppercase is used to\ndifferentiate from element classes.  Classes in lower_case_with_underscores\nare element classes, matching the XML element generic identifiers in the DTD_.\n\nThe position of each node (the level at which it can occur) is significant and\nis represented by abstract base classes (`Root`, `Structural`, `Body`,\n`Inline`, etc.).  Certain transformations will be easier because we can use\n``isinstance(node, base_class)`` to determine the position of the node in the\nhierarchy.\n\n.. _DTD: http://docutils.sourceforge.net/docs/ref/docutils.dtd\n").ToObject()); πE != nil {
				continue
			}
			// line 22: from __future__ import print_function
			πF.SetLineno(22)
			// line 23: from collections import Counter
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßCounter); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCounter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 25: __docformat__ = 'reStructuredText'
			πF.SetLineno(25)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 27: import sys
			πF.SetLineno(27)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: import os
			πF.SetLineno(28)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: import re
			πF.SetLineno(29)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: import warnings
			πF.SetLineno(30)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: import unicodedata
			πF.SetLineno(31)
			if πTemp002, πE = πg.ImportModule(πF, "unicodedata"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunicodedata.ToObject(), πTemp001); πE != nil {
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
			// line 33: if sys.version_info >= (3, 0):
			πF.SetLineno(33)
		Label1:
			// line 34: unicode = str  # noqa
			πF.SetLineno(34)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: basestring = str  # noqa
			πF.SetLineno(35)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßbasestring.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 37: class _traversal_list(list):
			πF.SetLineno(37)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_traversal_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 39: done = False
					πF.SetLineno(39)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdone.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 40: def _warning_decorator(fun):
					πF.SetLineno(40)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "fun", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_warning_decorator", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µfun *πg.Object = πArgs[0]
						_ = µfun
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µwrapper *πg.Object = πg.UnboundLocal
						_ = µwrapper
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
							// line 41: msg = ("\n   The iterable returned by Node.traverse()"
							πF.SetLineno(41)
							µmsg = πg.NewStr("\n   The iterable returned by Node.traverse()\n   will become an iterator instead of a list in Docutils > 0.16.").ToObject()
							// line 44: def wrapper(self, *args, **kwargs):
							πF.SetLineno(44)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("wrapper", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]
								_ = µself
								var µargs *πg.Object = πArgs[1]
								_ = µargs
								var µkwargs *πg.Object = πArgs[2]
								_ = µkwargs
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
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
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßdone, nil); πE != nil {
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
									// line 45: if not self.done:
									πF.SetLineno(45)
								Label1:
									// line 46: warnings.warn(msg, FutureWarning, stacklevel=2)
									πF.SetLineno(46)
									πTemp004 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
										continue
									}
									πTemp004[0] = µmsg
									if πTemp001, πE = πg.ResolveGlobal(πF, ßFutureWarning); πE != nil {
										continue
									}
									πTemp004[1] = πTemp001
									πTemp005 = πg.KWArgs{
										{"stacklevel", πg.NewInt(2).ToObject()},
									}
									if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwarn, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp004, πTemp005); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									goto Label2
								Label2:
									// line 47: self.done = True
									πF.SetLineno(47)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßdone, πTemp002); πE != nil {
										continue
									}
									// line 48: return fun(self, *args, **kwargs)
									πF.SetLineno(48)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp004[0] = µself
									if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µfun, "fun"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Invoke(πF, µfun, πTemp004, µargs, nil, µkwargs); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
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
							µwrapper = πTemp001
							// line 49: return wrapper
							πF.SetLineno(49)
							if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
								continue
							}
							πR = µwrapper
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_warning_decorator.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 51: __add__ = _warning_decorator(list.__add__)
					πF.SetLineno(51)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__add__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 52: __contains__ = _warning_decorator(list.__contains__)
					πF.SetLineno(52)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__contains__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 53: __getitem__ = _warning_decorator(list.__getitem__)
					πF.SetLineno(53)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__getitem__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 54: __reversed__ = _warning_decorator(list.__reversed__)
					πF.SetLineno(54)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__reversed__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__reversed__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 55: __setitem__ = _warning_decorator(list.__setitem__)
					πF.SetLineno(55)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__setitem__, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 56: append = _warning_decorator(list.append)
					πF.SetLineno(56)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 57: count = _warning_decorator(list.count)
					πF.SetLineno(57)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcount, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßcount.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 58: extend = _warning_decorator(list.extend)
					πF.SetLineno(58)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßextend, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßextend.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 59: index = _warning_decorator(list.index)
					πF.SetLineno(59)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßindex, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßindex.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 60: insert = _warning_decorator(list.insert)
					πF.SetLineno(60)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßinsert, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 61: pop = _warning_decorator(list.pop)
					πF.SetLineno(61)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßpop, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 62: reverse = _warning_decorator(list.reverse)
					πF.SetLineno(62)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßreverse, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_warning_decorator); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßreverse.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_traversal_list").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_traversal_list.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 69: class Node(object):
			πF.SetLineno(69)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Node", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 71: """Abstract base class of nodes in a document tree."""
					πF.SetLineno(71)
					// line 71: """Abstract base class of nodes in a document tree."""
					πF.SetLineno(71)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Abstract base class of nodes in a document tree.").ToObject()); πE != nil {
						continue
					}
					// line 73: parent = None
					πF.SetLineno(73)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßparent.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 74: """Back-reference to the Node immediately containing this Node."""
					πF.SetLineno(74)
					// line 76: document = None
					πF.SetLineno(76)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdocument.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 77: """The `document` node at the root of the tree containing this Node."""
					πF.SetLineno(77)
					// line 79: source = None
					πF.SetLineno(79)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsource.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 80: """Path or description of the input source which generated this Node."""
					πF.SetLineno(80)
					// line 82: line = None
					πF.SetLineno(82)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßline.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 83: """The line number (1-based) of the beginning of this Node in `source`."""
					πF.SetLineno(83)
					// line 85: def __bool__(self):
					πF.SetLineno(85)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__bool__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 86: """
							πF.SetLineno(86)
							// line 94: return True
							πF.SetLineno(94)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__bool__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 86: """
					πF.SetLineno(86)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Node instances are always true, even if they're empty.  A node is more\n        than a simple container.  Its boolean \"truth\" does not depend on\n        having one or more subnodes in the doctree.\n\n        Use `len()` to check node length.  Use `None` to represent a boolean\n        false value.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__bool__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp003, πE = πg.LT(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 96: if sys.version_info < (3, 0):
					πF.SetLineno(96)
				Label1:
					// line 97: __nonzero__ = __bool__
					πF.SetLineno(97)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß__bool__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__nonzero__.ToObject(), πTemp003); πE != nil {
						continue
					}
					goto Label2
				Label2:
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp003, πE = πg.LT(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 99: if sys.version_info < (3, 0):
					πF.SetLineno(99)
				Label3:
					// line 102: def __str__(self):
					πF.SetLineno(102)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 103: return unicode(self).encode('raw_unicode_escape')
							πF.SetLineno(103)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßraw_unicode_escape.ToObject()
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßencode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
					goto Label4
				Label4:
					// line 105: def asdom(self, dom=None):
					πF.SetLineno(105)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "dom", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("asdom", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdom *πg.Object = πArgs[1]
						_ = µdom
						var µdomroot *πg.Object = πg.UnboundLocal
						_ = µdomroot
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
							// line 106: """Return a DOM **fragment** representation of this Node."""
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µdom, "dom"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdom == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 107: if dom is None:
							πF.SetLineno(107)
						Label1:
							// line 108: import xml.dom.minidom as dom
							πF.SetLineno(108)
							if πTemp004, πE = πg.ImportModule(πF, "xml.dom.minidom"); πE != nil {
								continue
							}
							πTemp001 = πTemp004[2]
							µdom = πTemp001
							goto Label2
						Label2:
							// line 109: domroot = dom.Document()
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µdom, "dom"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdom, ßDocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdomroot = πTemp002
							// line 110: return self._dom_node(domroot)
							πF.SetLineno(110)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							πTemp004[0] = µdomroot
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_dom_node, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßasdom.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 106: """Return a DOM **fragment** representation of this Node."""
					πF.SetLineno(106)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a DOM **fragment** representation of this Node.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßasdom); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 112: def pformat(self, indent='    ', level=0):
					πF.SetLineno(112)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent", Def: πg.NewStr("    ").ToObject()}
					πTemp002[2] = πg.Param{Name: "level", Def: πg.NewInt(0).ToObject()}
					πTemp005 = πg.NewFunction(πg.NewCode("pformat", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindent *πg.Object = πArgs[1]
						_ = µindent
						var µlevel *πg.Object = πArgs[2]
						_ = µlevel
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
							// line 113: """
							πF.SetLineno(113)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 118: raise NotImplementedError
							πF.SetLineno(118)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpformat.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 113: """
					πF.SetLineno(113)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Return an indented pseudo-XML representation, for test purposes.\n\n        Override in subclasses.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßpformat); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 120: def copy(self):
					πF.SetLineno(120)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 121: """Return a copy of self."""
							πF.SetLineno(121)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 122: raise NotImplementedError
							πF.SetLineno(122)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 121: """Return a copy of self."""
					πF.SetLineno(121)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return a copy of self.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßcopy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 124: def deepcopy(self):
					πF.SetLineno(124)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("deepcopy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 125: """Return a deep copy of self (also copying children)."""
							πF.SetLineno(125)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 126: raise NotImplementedError
							πF.SetLineno(126)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdeepcopy.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 125: """Return a deep copy of self (also copying children)."""
					πF.SetLineno(125)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return a deep copy of self (also copying children).").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdeepcopy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 128: def astext(self):
					πF.SetLineno(128)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 129: """Return a string representation of this Node."""
							πF.SetLineno(129)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 130: raise NotImplementedError
							πF.SetLineno(130)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 129: """Return a string representation of this Node."""
					πF.SetLineno(129)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Return a string representation of this Node.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßastext); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 132: def setup_child(self, child):
					πF.SetLineno(132)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "child", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("setup_child", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchild *πg.Object = πArgs[1]
						_ = µchild
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 133: child.parent = self
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µself); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µchild, ßparent, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 134: if self.document:
							πF.SetLineno(134)
						Label1:
							// line 135: child.document = self.document
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µchild, ßdocument, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchild, ßsource, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 136: if child.source is None:
							πF.SetLineno(136)
						Label3:
							// line 137: child.source = self.document.current_source
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcurrent_source, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µchild, ßsource, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchild, ßline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 138: if child.line is None:
							πF.SetLineno(138)
						Label5:
							// line 139: child.line = self.document.current_line
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcurrent_line, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µchild, ßline, πTemp001); πE != nil {
								continue
							}
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ßsetup_child.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 141: def walk(self, visitor):
					πF.SetLineno(141)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "visitor", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("walk", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πArgs[1]
						_ = µvisitor
						var µstop *πg.Object = πg.UnboundLocal
						_ = µstop
						var µchildren *πg.Object = πg.UnboundLocal
						_ = µchildren
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
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
							case 8:
								goto Label8
							case 9:
								goto Label9
							case 2:
								goto Label2
							case 4:
								goto Label4
							case 10:
								goto Label10
							default:
								panic("unexpected function state")
							}
							// line 142: """
							πF.SetLineno(142)
							// line 164: stop = False
							πF.SetLineno(164)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µstop = πTemp001
							// line 165: visitor.document.reporter.debug(
							πF.SetLineno(165)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("docutils.nodes.Node.walk calling dispatch_visit for %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvisitor, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 168: try:
							πF.SetLineno(168)
							πF.PushCheckpoint(2)
							// line 169: try:
							πF.SetLineno(169)
							πF.PushCheckpoint(4)
							// line 170: visitor.dispatch_visit(self)
							πF.SetLineno(170)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvisitor, ßdispatch_visit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSkipChildren); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßSkipNode); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipDeparture); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 171: except (SkipChildren, SkipNode):
							πF.SetLineno(171)
						Label5:
							// line 172: return stop
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							πR = µstop
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
							// line 173: except SkipDeparture:           # not applicable; ignore
							πF.SetLineno(173)
						Label6:
							// line 174: pass
							πF.SetLineno(174)
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 175: children = self.children
							πF.SetLineno(175)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							µchildren = πTemp001
							// line 176: try:
							πF.SetLineno(176)
							πF.PushCheckpoint(8)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µchildren, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp007 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label11
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
							πF.PushCheckpoint(9)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp002[0] = µvisitor
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchild, ßwalk, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label12
							}
							goto Label13
							// line 178: if child.walk(visitor):
							πF.SetLineno(178)
						Label12:
							// line 179: stop = True
							πF.SetLineno(179)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µstop = πTemp003
							// line 180: break
							πF.SetLineno(180)
							πTemp007 = true
							continue
							goto Label13
						Label13:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipSiblings); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label14
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 181: except SkipSiblings:
							πF.SetLineno(181)
						Label14:
							// line 182: pass
							πF.SetLineno(182)
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopTraversal); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label15
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 183: except StopTraversal:
							πF.SetLineno(183)
						Label15:
							// line 184: stop = True
							πF.SetLineno(184)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µstop = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 185: return stop
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							πR = µstop
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwalk.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 142: """
					πF.SetLineno(142)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n        Traverse a tree of `Node` objects, calling the\n        `dispatch_visit()` method of `visitor` when entering each\n        node.  (The `walkabout()` method is similar, except it also\n        calls the `dispatch_departure()` method before exiting each\n        node.)\n\n        This tree traversal supports limited in-place tree\n        modifications.  Replacing one node with one or more nodes is\n        OK, as is removing an element.  However, if the node removed\n        or replaced occurs after the current node, the old node will\n        still be traversed, and any new nodes will not.\n\n        Within ``visit`` methods (and ``depart`` methods for\n        `walkabout()`), `TreePruningException` subclasses may be raised\n        (`SkipChildren`, `SkipSiblings`, `SkipNode`, `SkipDeparture`).\n\n        Parameter `visitor`: A `NodeVisitor` object, containing a\n        ``visit`` implementation for each `Node` subclass encountered.\n\n        Return true if we should stop the traversal.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßwalk); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 187: def walkabout(self, visitor):
					πF.SetLineno(187)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "visitor", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("walkabout", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πArgs[1]
						_ = µvisitor
						var µcall_depart *πg.Object = πg.UnboundLocal
						_ = µcall_depart
						var µstop *πg.Object = πg.UnboundLocal
						_ = µstop
						var µchildren *πg.Object = πg.UnboundLocal
						_ = µchildren
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
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
							case 8:
								goto Label8
							case 9:
								goto Label9
							case 2:
								goto Label2
							case 4:
								goto Label4
							case 10:
								goto Label10
							default:
								panic("unexpected function state")
							}
							// line 188: """
							πF.SetLineno(188)
							// line 199: call_depart = True
							πF.SetLineno(199)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µcall_depart = πTemp001
							// line 200: stop = False
							πF.SetLineno(200)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µstop = πTemp001
							// line 201: visitor.document.reporter.debug(
							πF.SetLineno(201)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("docutils.nodes.Node.walkabout calling dispatch_visit for %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvisitor, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 204: try:
							πF.SetLineno(204)
							πF.PushCheckpoint(2)
							// line 205: try:
							πF.SetLineno(205)
							πF.PushCheckpoint(4)
							// line 206: visitor.dispatch_visit(self)
							πF.SetLineno(206)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvisitor, ßdispatch_visit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipNode); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipDeparture); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 207: except SkipNode:
							πF.SetLineno(207)
						Label5:
							// line 208: return stop
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							πR = µstop
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
							// line 209: except SkipDeparture:
							πF.SetLineno(209)
						Label6:
							// line 210: call_depart = False
							πF.SetLineno(210)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µcall_depart = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 211: children = self.children
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							µchildren = πTemp001
							// line 212: try:
							πF.SetLineno(212)
							πF.PushCheckpoint(8)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µchildren, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp007 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label11
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
							πF.PushCheckpoint(9)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp002[0] = µvisitor
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchild, ßwalkabout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label12
							}
							goto Label13
							// line 214: if child.walkabout(visitor):
							πF.SetLineno(214)
						Label12:
							// line 215: stop = True
							πF.SetLineno(215)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µstop = πTemp003
							// line 216: break
							πF.SetLineno(216)
							πTemp007 = true
							continue
							goto Label13
						Label13:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipSiblings); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label14
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 217: except SkipSiblings:
							πF.SetLineno(217)
						Label14:
							// line 218: pass
							πF.SetLineno(218)
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipChildren); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label15
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopTraversal); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label16
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 219: except SkipChildren:
							πF.SetLineno(219)
						Label15:
							// line 220: pass
							πF.SetLineno(220)
							πF.RestoreExc(nil, nil)
							goto Label1
							// line 221: except StopTraversal:
							πF.SetLineno(221)
						Label16:
							// line 222: stop = True
							πF.SetLineno(222)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µstop = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							if πE = πg.CheckLocal(πF, µcall_depart, "call_depart"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µcall_depart); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label17
							}
							goto Label18
							// line 223: if call_depart:
							πF.SetLineno(223)
						Label17:
							// line 224: visitor.document.reporter.debug(
							πF.SetLineno(224)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("docutils.nodes.Node.walkabout calling dispatch_departure for %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvisitor, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 227: visitor.dispatch_departure(self)
							πF.SetLineno(227)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvisitor, ßdispatch_departure, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label18
						Label18:
							// line 228: return stop
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							πR = µstop
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwalkabout.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 188: """
					πF.SetLineno(188)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n        Perform a tree traversal similarly to `Node.walk()` (which\n        see), except also call the `dispatch_departure()` method\n        before exiting each node.\n\n        Parameter `visitor`: A `NodeVisitor` object, containing a\n        ``visit`` and ``depart`` implementation for each `Node`\n        subclass encountered.\n\n        Return true if we should stop the traversal.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßwalkabout); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 230: def _fast_traverse(self, cls):
					πF.SetLineno(230)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "cls", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("_fast_traverse", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcls *πg.Object = πArgs[1]
						_ = µcls
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var µsubnode *πg.Object = πg.UnboundLocal
						_ = µsubnode
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
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 3:
									goto Label3
								case 4:
									goto Label4
								case 5:
									goto Label5
								case 7:
									goto Label7
								case 8:
									goto Label8
								case 10:
									goto Label10
								default:
									panic("unexpected function state")
								}
								// line 231: """Return iterator that only supports instance checks."""
								πF.SetLineno(231)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
									continue
								}
								πTemp001[1] = µcls
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
								// line 232: if isinstance(self, cls):
								πF.SetLineno(232)
							Label1:
								// line 233: yield self
								πF.SetLineno(233)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								return µself, nil
							Label3:
								πTemp002 = πSent
								goto Label2
							Label2:
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								πTemp004 = false
							Label4:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label6
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
								πF.PushCheckpoint(4)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
									continue
								}
								πTemp001[0] = µcls
								if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µchild, ß_fast_traverse, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
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
									πTemp008 = !isStop
								} else {
									πTemp008 = true
									µsubnode = πTemp006
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(7)
								// line 236: yield subnode
								πF.SetLineno(236)
								if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
									continue
								}
								πF.PushCheckpoint(10)
								return µsubnode, nil
							Label10:
								πTemp006 = πSent
								continue
							Label8:
								if πE != nil || πR != nil {
									continue
								}
							Label9:
								continue
							Label5:
								if πE != nil || πR != nil {
									continue
								}
							Label6:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_fast_traverse.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 231: """Return iterator that only supports instance checks."""
					πF.SetLineno(231)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Return iterator that only supports instance checks.").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_fast_traverse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 238: def _all_traverse(self):
					πF.SetLineno(238)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("_all_traverse", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var µsubnode *πg.Object = πg.UnboundLocal
						_ = µsubnode
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
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
								case 3:
									goto Label3
								case 5:
									goto Label5
								case 6:
									goto Label6
								case 8:
									goto Label8
								default:
									panic("unexpected function state")
								}
								// line 239: """Return iterator that doesn't check for a condition."""
								πF.SetLineno(239)
								// line 240: yield self
								πF.SetLineno(240)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								return µself, nil
							Label1:
								πTemp001 = πSent
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								πTemp003 = false
							Label2:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp003 {
									πF.PopCheckpoint()
									goto Label4
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
								πF.PushCheckpoint(2)
								if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µchild, ß_all_traverse, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								πTemp004 = false
							Label5:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label7
								}
								if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
									µsubnode = πTemp005
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(5)
								// line 243: yield subnode
								πF.SetLineno(243)
								if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								return µsubnode, nil
							Label8:
								πTemp005 = πSent
								continue
							Label6:
								if πE != nil || πR != nil {
									continue
								}
							Label7:
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_all_traverse.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 239: """Return iterator that doesn't check for a condition."""
					πF.SetLineno(239)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Return iterator that doesn't check for a condition.").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ß_all_traverse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 245: def traverse(self, condition=None, include_self=True, descend=True,
					πF.SetLineno(245)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "condition", Def: πTemp016}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "include_self", Def: πTemp016}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "descend", Def: πTemp016}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "siblings", Def: πTemp016}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "ascend", Def: πTemp016}
					πTemp015 = πg.NewFunction(πg.NewCode("traverse", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcondition *πg.Object = πArgs[1]
						_ = µcondition
						var µinclude_self *πg.Object = πArgs[2]
						_ = µinclude_self
						var µdescend *πg.Object = πArgs[3]
						_ = µdescend
						var µsiblings *πg.Object = πArgs[4]
						_ = µsiblings
						var µascend *πg.Object = πArgs[5]
						_ = µascend
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
							// line 247: """
							πF.SetLineno(247)
							// line 286: return _traversal_list(self._traverse(condition, include_self,
							πF.SetLineno(286)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
								continue
							}
							πTemp002[0] = µcondition
							if πE = πg.CheckLocal(πF, µinclude_self, "include_self"); πE != nil {
								continue
							}
							πTemp002[1] = µinclude_self
							if πE = πg.CheckLocal(πF, µdescend, "descend"); πE != nil {
								continue
							}
							πTemp002[2] = µdescend
							if πE = πg.CheckLocal(πF, µsiblings, "siblings"); πE != nil {
								continue
							}
							πTemp002[3] = µsiblings
							if πE = πg.CheckLocal(πF, µascend, "ascend"); πE != nil {
								continue
							}
							πTemp002[4] = µascend
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_traverse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_traversal_list); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtraverse.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 247: """
					πF.SetLineno(247)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("\n        Return an iterable containing\n\n        * self (if include_self is true)\n        * all descendants in tree traversal order (if descend is true)\n        * all siblings (if siblings is true) and their descendants (if\n          also descend is true)\n        * the siblings of the parent (if ascend is true) and their\n          descendants (if also descend is true), and so on\n\n        If `condition` is not None, the iterable contains only nodes\n        for which ``condition(node)`` is true.  If `condition` is a\n        node class ``cls``, it is equivalent to a function consisting\n        of ``return isinstance(node, cls)``.\n\n        If ascend is true, assume siblings to be true as well.\n\n        For example, given the following tree::\n\n            <paragraph>\n                <emphasis>      <--- emphasis.traverse() and\n                    <strong>    <--- strong.traverse() are called.\n                        Foo\n                    Bar\n                <reference name=\"Baz\" refid=\"baz\">\n                    Baz\n\n        Then list(emphasis.traverse()) equals ::\n\n            [<emphasis>, <strong>, <#text: Foo>, <#text: Bar>]\n\n        and list(strong.traverse(ascend=True)) equals ::\n\n            [<strong>, <#text: Foo>, <#text: Bar>, <reference>, <#text: Baz>]\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtraverse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 289: def _traverse(self, condition=None, include_self=True, descend=True,
					πF.SetLineno(289)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "condition", Def: πTemp017}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "include_self", Def: πTemp017}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "descend", Def: πTemp017}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "siblings", Def: πTemp017}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "ascend", Def: πTemp017}
					πTemp016 = πg.NewFunction(πg.NewCode("_traverse", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcondition *πg.Object = πArgs[1]
						_ = µcondition
						var µinclude_self *πg.Object = πArgs[2]
						_ = µinclude_self
						var µdescend *πg.Object = πArgs[3]
						_ = µdescend
						var µsiblings *πg.Object = πArgs[4]
						_ = µsiblings
						var µascend *πg.Object = πArgs[5]
						_ = µascend
						var µsubnode *πg.Object = πg.UnboundLocal
						_ = µsubnode
						var µnode_class *πg.Object = πg.UnboundLocal
						_ = µnode_class
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µsibling *πg.Object = πg.UnboundLocal
						_ = µsibling
						var πTemp001 bool
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
						var πTemp012 πg.KWArgs
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 33:
									goto Label33
								case 37:
									goto Label37
								case 38:
									goto Label38
								case 40:
									goto Label40
								case 9:
									goto Label9
								case 10:
									goto Label10
								case 43:
									goto Label43
								case 12:
									goto Label12
								case 13:
									goto Label13
								case 14:
									goto Label14
								case 44:
									goto Label44
								case 16:
									goto Label16
								case 41:
									goto Label41
								case 23:
									goto Label23
								case 46:
									goto Label46
								case 27:
									goto Label27
								case 28:
									goto Label28
								case 30:
									goto Label30
								case 31:
									goto Label31
								default:
									panic("unexpected function state")
								}
								// line 291: """Return iterator over nodes following `self`. See `traverse()`."""
								πF.SetLineno(291)
								if πE = πg.CheckLocal(πF, µascend, "ascend"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.IsTrue(πF, µascend); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label1
								}
								goto Label2
								// line 292: if ascend:
								πF.SetLineno(292)
							Label1:
								// line 293: siblings=True
								πF.SetLineno(293)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								µsiblings = πTemp002
								goto Label2
							Label2:
								if πE = πg.CheckLocal(πF, µinclude_self, "include_self"); πE != nil {
									continue
								}
								πTemp002 = µinclude_self
								if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if !πTemp001 {
									goto Label3
								}
								if πE = πg.CheckLocal(πF, µdescend, "descend"); πE != nil {
									continue
								}
								πTemp002 = µdescend
								if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if !πTemp001 {
									goto Label3
								}
								if πE = πg.CheckLocal(πF, µsiblings, "siblings"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, µsiblings); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp004).ToObject()
								πTemp002 = πTemp003
							Label3:
								if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label4
								}
								goto Label5
								// line 296: if include_self and descend and not siblings:
								πF.SetLineno(296)
							Label4:
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(µcondition == πTemp003).ToObject()
								if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label6
								}
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								πTemp005[0] = µcondition
								if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
									continue
								}
								πTemp005[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label7
								}
								goto Label8
								// line 297: if condition is None:
								πF.SetLineno(297)
							Label6:
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_all_traverse, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
									continue
								}
								πF.PushCheckpoint(10)
								πTemp001 = false
							Label9:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label11
								}
								if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
									µsubnode = πTemp003
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(9)
								// line 299: yield subnode
								πF.SetLineno(299)
								if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								return µsubnode, nil
							Label12:
								πTemp003 = πSent
								continue
							Label10:
								if πE != nil || πR != nil {
									continue
								}
							Label11:
								// line 300: return
								πF.SetLineno(300)
								πR = πg.None
								continue
								goto Label8
								// line 301: elif isinstance(condition, type):
								πF.SetLineno(301)
							Label7:
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								πTemp005[0] = µcondition
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ß_fast_traverse, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
									continue
								}
								πF.PushCheckpoint(14)
								πTemp001 = false
							Label13:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label15
								}
								if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
									µsubnode = πTemp003
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(13)
								// line 303: yield subnode
								πF.SetLineno(303)
								if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
									continue
								}
								πF.PushCheckpoint(16)
								return µsubnode, nil
							Label16:
								πTemp003 = πSent
								continue
							Label14:
								if πE != nil || πR != nil {
									continue
								}
							Label15:
								// line 304: return
								πF.SetLineno(304)
								πR = πg.None
								continue
								goto Label8
							Label8:
								goto Label5
							Label5:
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								πTemp005[0] = µcondition
								if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
									continue
								}
								πTemp005[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label17
								}
								goto Label18
								// line 307: if isinstance(condition, type):
								πF.SetLineno(307)
							Label17:
								// line 308: node_class = condition
								πF.SetLineno(308)
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								µnode_class = µcondition
								// line 309: def condition(node, node_class=node_class):
								πF.SetLineno(309)
								πTemp007 = make([]πg.Param, 2)
								πTemp007[0] = πg.Param{Name: "node", Def: nil}
								if πE = πg.CheckLocal(πF, µnode_class, "node_class"); πE != nil {
									continue
								}
								πTemp007[1] = πg.Param{Name: "node_class", Def: µnode_class}
								πTemp002 = πg.NewFunction(πg.NewCode("condition", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
									var µnode *πg.Object = πArgs[0]
									_ = µnode
									var µnode_class *πg.Object = πArgs[1]
									_ = µnode_class
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
										// line 310: return isinstance(node, node_class)
										πF.SetLineno(310)
										πTemp001 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										πTemp001[0] = µnode
										if πE = πg.CheckLocal(πF, µnode_class, "node_class"); πE != nil {
											continue
										}
										πTemp001[1] = µnode_class
										if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								µcondition = πTemp002
								goto Label18
							Label18:
								if πE = πg.CheckLocal(πF, µinclude_self, "include_self"); πE != nil {
									continue
								}
								πTemp003 = µinclude_self
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if !πTemp001 {
									goto Label19
								}
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp008 = πg.GetBool(µcondition == πTemp009).ToObject()
								πTemp006 = πTemp008
								if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label20
								}
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp005[0] = µself
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								if πTemp008, πE = µcondition.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp006 = πTemp008
							Label20:
								πTemp003 = πTemp006
							Label19:
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label21
								}
								goto Label22
								// line 313: if include_self and (condition is None or condition(self)):
								πF.SetLineno(313)
							Label21:
								// line 314: yield self
								πF.SetLineno(314)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πF.PushCheckpoint(23)
								return µself, nil
							Label23:
								πTemp003 = πSent
								goto Label22
							Label22:
								if πE = πg.CheckLocal(πF, µdescend, "descend"); πE != nil {
									continue
								}
								πTemp003 = µdescend
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if !πTemp001 {
									goto Label24
								}
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp006
								if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp003 = πTemp008
							Label24:
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label25
								}
								goto Label26
								// line 315: if descend and len(self.children):
								πF.SetLineno(315)
							Label25:
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Iter(πF, µself); πE != nil {
									continue
								}
								πF.PushCheckpoint(28)
								πTemp001 = false
							Label27:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label29
								}
								if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
									µchild = πTemp006
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(27)
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp012 = πg.KWArgs{
									{"condition", µcondition},
									{"include_self", πTemp008},
									{"descend", πTemp009},
									{"siblings", πTemp010},
									{"ascend", πTemp011},
								}
								if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µchild, ß_traverse, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, nil, πTemp012); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Iter(πF, πTemp009); πE != nil {
									continue
								}
								πF.PushCheckpoint(31)
								πTemp004 = false
							Label30:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label32
								}
								if πTemp008, πE = πg.Next(πF, πTemp006); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp013 = !isStop
								} else {
									πTemp013 = true
									µsubnode = πTemp008
								}
								if πE != nil || !πTemp013 {
									continue
								}
								πF.PushCheckpoint(30)
								// line 320: yield subnode
								πF.SetLineno(320)
								if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
									continue
								}
								πF.PushCheckpoint(33)
								return µsubnode, nil
							Label33:
								πTemp008 = πSent
								continue
							Label31:
								if πE != nil || πR != nil {
									continue
								}
							Label32:
								continue
							Label28:
								if πE != nil || πR != nil {
									continue
								}
							Label29:
								goto Label26
							Label26:
								if πE = πg.CheckLocal(πF, µsiblings, "siblings"); πE != nil {
									continue
								}
								πTemp003 = µsiblings
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label34
								}
								if πE = πg.CheckLocal(πF, µascend, "ascend"); πE != nil {
									continue
								}
								πTemp003 = µascend
							Label34:
								if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label35
								}
								goto Label36
								// line 321: if siblings or ascend:
								πF.SetLineno(321)
							Label35:
								// line 322: node = self
								πF.SetLineno(322)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								µnode = µself
								// line 323: while node.parent:
								πF.SetLineno(323)
								πF.PushCheckpoint(38)
								πTemp001 = false
							Label37:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label39
								}
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(37)
								// line 324: index = node.parent.index(node)
								πF.SetLineno(324)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								πTemp005[0] = µnode
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßindex, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µindex = πTemp003
								if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πTemp008, πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
									continue
								}
								πF.PushCheckpoint(41)
								πTemp004 = false
							Label40:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label42
								}
								if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp013 = !isStop
								} else {
									πTemp013 = true
									µsibling = πTemp006
								}
								if πE != nil || !πTemp013 {
									continue
								}
								πF.PushCheckpoint(40)
								if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µdescend, "descend"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πTemp012 = πg.KWArgs{
									{"condition", µcondition},
									{"include_self", πTemp008},
									{"descend", µdescend},
									{"siblings", πTemp009},
									{"ascend", πTemp010},
								}
								if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µsibling, ß_traverse, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, nil, πTemp012); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Iter(πF, πTemp009); πE != nil {
									continue
								}
								πF.PushCheckpoint(44)
								πTemp013 = false
							Label43:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp013 {
									πF.PopCheckpoint()
									goto Label45
								}
								if πTemp008, πE = πg.Next(πF, πTemp006); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp014 = !isStop
								} else {
									πTemp014 = true
									µsubnode = πTemp008
								}
								if πE != nil || !πTemp014 {
									continue
								}
								πF.PushCheckpoint(43)
								// line 329: yield subnode
								πF.SetLineno(329)
								if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
									continue
								}
								πF.PushCheckpoint(46)
								return µsubnode, nil
							Label46:
								πTemp008 = πSent
								continue
							Label44:
								if πE != nil || πR != nil {
									continue
								}
							Label45:
								continue
							Label41:
								if πE != nil || πR != nil {
									continue
								}
							Label42:
								if πE = πg.CheckLocal(πF, µascend, "ascend"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, µascend); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(!πTemp004).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label47
								}
								goto Label48
								// line 330: if not ascend:
								πF.SetLineno(330)
							Label47:
								// line 331: break
								πF.SetLineno(331)
								πTemp001 = true
								continue
								goto Label49
							Label48:
								// line 333: node = node.parent
								πF.SetLineno(333)
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
									continue
								}
								µnode = πTemp003
								goto Label49
							Label49:
								continue
							Label38:
								if πE != nil || πR != nil {
									continue
								}
							Label39:
								goto Label36
							Label36:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_traverse.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 291: """Return iterator over nodes following `self`. See `traverse()`."""
					πF.SetLineno(291)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Return iterator over nodes following `self`. See `traverse()`.").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß_traverse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 335: def next_node(self, condition=None, include_self=False, descend=True,
					πF.SetLineno(335)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "condition", Def: πTemp018}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "include_self", Def: πTemp018}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "descend", Def: πTemp018}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "siblings", Def: πTemp018}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "ascend", Def: πTemp018}
					πTemp017 = πg.NewFunction(πg.NewCode("next_node", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcondition *πg.Object = πArgs[1]
						_ = µcondition
						var µinclude_self *πg.Object = πArgs[2]
						_ = µinclude_self
						var µdescend *πg.Object = πArgs[3]
						_ = µdescend
						var µsiblings *πg.Object = πArgs[4]
						_ = µsiblings
						var µascend *πg.Object = πArgs[5]
						_ = µascend
						var µnode_iterator *πg.Object = πg.UnboundLocal
						_ = µnode_iterator
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
							// line 337: """
							πF.SetLineno(337)
							// line 344: node_iterator = self._traverse(condition, include_self,
							πF.SetLineno(344)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
								continue
							}
							πTemp001[0] = µcondition
							if πE = πg.CheckLocal(πF, µinclude_self, "include_self"); πE != nil {
								continue
							}
							πTemp001[1] = µinclude_self
							if πE = πg.CheckLocal(πF, µdescend, "descend"); πE != nil {
								continue
							}
							πTemp001[2] = µdescend
							if πE = πg.CheckLocal(πF, µsiblings, "siblings"); πE != nil {
								continue
							}
							πTemp001[3] = µsiblings
							if πE = πg.CheckLocal(πF, µascend, "ascend"); πE != nil {
								continue
							}
							πTemp001[4] = µascend
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_traverse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnode_iterator = πTemp003
							// line 346: try:
							πF.SetLineno(346)
							πF.PushCheckpoint(2)
							// line 347: return next(node_iterator)
							πF.SetLineno(347)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode_iterator, "node_iterator"); πE != nil {
								continue
							}
							πTemp001[0] = µnode_iterator
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
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
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
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
							// line 348: except StopIteration:
							πF.SetLineno(348)
						Label3:
							// line 349: return None
							πF.SetLineno(349)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnext_node.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 337: """
					πF.SetLineno(337)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("\n        Return the first node in the iterable returned by traverse(),\n        or None if the iterable is empty.\n\n        Parameter list is the same as of traverse.  Note that\n        include_self defaults to False, though.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßnext_node); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Node").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNode.ToObject(), πTemp004); πE != nil {
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
				goto Label3
			}
			goto Label4
			// line 351: if sys.version_info < (3, 0):
			πF.SetLineno(351)
		Label3:
			// line 352: class reprunicode(unicode):
			πF.SetLineno(352)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("reprunicode", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 353: """
					πF.SetLineno(353)
					// line 353: """
					πF.SetLineno(353)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n        A unicode sub-class that removes the initial u from unicode's repr.\n        ").ToObject()); πE != nil {
						continue
					}
					// line 357: def __repr__(self):
					πF.SetLineno(357)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 358: return unicode.__repr__(self)[1:]
							πF.SetLineno(358)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__repr__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("reprunicode").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreprunicode.ToObject(), πTemp004); πE != nil {
				continue
			}
			goto Label5
		Label4:
			// line 360: reprunicode = unicode
			πF.SetLineno(360)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreprunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label5
		Label5:
			// line 363: def ensure_str(s):
			πF.SetLineno(363)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("ensure_str", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]
				_ = µs
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
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 364: """
					πF.SetLineno(364)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp003, πE = πg.LT(πF, πTemp005, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[0] = µs
					if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp006[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001 = πTemp004
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 367: if sys.version_info < (3, 0) and isinstance(s, unicode):
					πF.SetLineno(367)
				Label2:
					// line 368: return s.encode('ascii', 'backslashreplace')
					πF.SetLineno(368)
					πTemp006 = πF.MakeArgs(2)
					πTemp006[0] = ßascii.ToObject()
					πTemp006[1] = ßbackslashreplace.ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µs, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πR = πTemp003
					continue
					goto Label3
				Label3:
					// line 369: return s
					πF.SetLineno(369)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßensure_str.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 364: """
			πF.SetLineno(364)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Failsave conversion of `unicode` to `str`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßensure_str); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 372: def unescape(text, restore_backslashes=False, respect_whitespace=False):
			πF.SetLineno(372)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp007[1] = πg.Param{Name: "restore_backslashes", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp007[2] = πg.Param{Name: "respect_whitespace", Def: πTemp004}
			πTemp003 = πg.NewFunction(πg.NewCode("unescape", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µrestore_backslashes *πg.Object = πArgs[1]
				_ = µrestore_backslashes
				var µrespect_whitespace *πg.Object = πArgs[2]
				_ = µrespect_whitespace
				var µsep *πg.Object = πg.UnboundLocal
				_ = µsep
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					case 4:
						goto Label4
					case 5:
						goto Label5
					default:
						panic("unexpected function state")
					}
					// line 373: """
					πF.SetLineno(373)
					if πE = πg.CheckLocal(πF, µrestore_backslashes, "restore_backslashes"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µrestore_backslashes); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 378: if restore_backslashes:
					πF.SetLineno(378)
				Label1:
					// line 379: return text.replace('\x00', '\\')
					πF.SetLineno(379)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("\x00").ToObject()
					πTemp002[1] = πg.NewStr("\\").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp004
					continue
					goto Label3
				Label2:
					πTemp002 = make([]*πg.Object, 3)
					πTemp002[0] = πg.NewStr("\x00 ").ToObject()
					πTemp002[1] = πg.NewStr("\x00\n").ToObject()
					πTemp002[2] = πg.NewStr("\x00").ToObject()
					πTemp004 = πg.NewList(πTemp002...).ToObject()
					if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp001 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp001 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µsep = πTemp004
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 382: text = ''.join(text.split(sep))
					πF.SetLineno(382)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					πTemp006[0] = µsep
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002[0] = πTemp007
					if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µtext = πTemp007
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 383: return text
					πF.SetLineno(383)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
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
			if πE = πF.Globals().SetItem(πF, ßunescape.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 373: """
			πF.SetLineno(373)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Return a string with nulls removed or restored to backslashes.\n    Backslash-escaped spaces are also removed.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßunescape); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 386: class Text(Node, reprunicode):
			πF.SetLineno(386)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Text", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 388: """
					πF.SetLineno(388)
					// line 388: """
					πF.SetLineno(388)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Instances are terminal nodes (leaves) containing text only; no child\n    nodes or attributes.  Initialize by passing a string to the constructor.\n    Access the text itself with the `astext` method.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 394: tagname = '#text'
					πF.SetLineno(394)
					if πE = πClass.SetItem(πF, ßtagname.ToObject(), πg.NewStr("#text").ToObject()); πE != nil {
						continue
					}
					// line 396: children = ()
					πF.SetLineno(396)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßchildren.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 397: """Text nodes have no children, and cannot have children."""
					πF.SetLineno(397)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp001, πE = πg.GT(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 399: if sys.version_info > (3, 0):
					πF.SetLineno(399)
				Label1:
					// line 400: def __new__(cls, data, rawsource=None):
					πF.SetLineno(400)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "data", Def: nil}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "rawsource", Def: πTemp002}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]
						_ = µcls
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µrawsource *πg.Object = πArgs[2]
						_ = µrawsource
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
							// line 401: """Prevent the rawsource argument from propagating to str."""
							πF.SetLineno(401)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 402: if isinstance(data, bytes):
							πF.SetLineno(402)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("expecting str data, not bytes").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 403: raise TypeError('expecting str data, not bytes')
							πF.SetLineno(403)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 404: return reprunicode.__new__(cls, data)
							πF.SetLineno(404)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[1] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 401: """Prevent the rawsource argument from propagating to str."""
					πF.SetLineno(401)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("Prevent the rawsource argument from propagating to str.").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp003, ß__doc__, πTemp002); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 406: def __new__(cls, data, rawsource=None):
					πF.SetLineno(406)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "cls", Def: nil}
					πTemp005[1] = πg.Param{Name: "data", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "rawsource", Def: πTemp003}
					πTemp002 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]
						_ = µcls
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µrawsource *πg.Object = πArgs[2]
						_ = µrawsource
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
							// line 407: """Prevent the rawsource argument from propagating to str."""
							πF.SetLineno(407)
							// line 408: return reprunicode.__new__(cls, data)
							πF.SetLineno(408)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[1] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 407: """Prevent the rawsource argument from propagating to str."""
					πF.SetLineno(407)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Prevent the rawsource argument from propagating to str.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp003); πE != nil {
						continue
					}
					goto Label3
				Label3:
					// line 410: def __init__(self, data, rawsource=''):
					πF.SetLineno(410)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "data", Def: nil}
					πTemp005[2] = πg.Param{Name: "rawsource", Def: ß.ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µrawsource *πg.Object = πArgs[2]
						_ = µrawsource
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
							// line 411: self.rawsource = rawsource
							πF.SetLineno(411)
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrawsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrawsource, πTemp001); πE != nil {
								continue
							}
							// line 412: """The raw text from which this element was constructed."""
							πF.SetLineno(412)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 414: def shortrepr(self, maxlen=18):
					πF.SetLineno(414)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "maxlen", Def: πg.NewInt(18).ToObject()}
					πTemp006 = πg.NewFunction(πg.NewCode("shortrepr", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmaxlen *πg.Object = πArgs[1]
						_ = µmaxlen
						var µdata *πg.Object = πg.UnboundLocal
						_ = µdata
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
							// line 415: data = self
							πF.SetLineno(415)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							µdata = µself
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µmaxlen, "maxlen"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp004, µmaxlen); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 416: if len(data) > maxlen:
							πF.SetLineno(416)
						Label1:
							// line 417: data = data[:maxlen-4] + ' ...'
							πF.SetLineno(417)
							if πE = πg.CheckLocal(πF, µmaxlen, "maxlen"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µmaxlen, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µdata, πTemp003); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr(" ...").ToObject()); πE != nil {
								continue
							}
							µdata = πTemp001
							goto Label2
						Label2:
							// line 418: return '<%s: %r>' % (self.tagname, reprunicode(data))
							πF.SetLineno(418)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πTemp006, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s: %r>").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßshortrepr.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 420: def __repr__(self):
					πF.SetLineno(420)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 πg.KWArgs
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
							// line 421: return self.shortrepr(maxlen=68)
							πF.SetLineno(421)
							πTemp001 = πg.KWArgs{
								{"maxlen", πg.NewInt(68).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßshortrepr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 423: def _dom_node(self, domroot):
					πF.SetLineno(423)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "domroot", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_dom_node", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdomroot *πg.Object = πArgs[1]
						_ = µdomroot
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
							// line 424: return domroot.createTextNode(unicode(self))
							πF.SetLineno(424)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdomroot, ßcreateTextNode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_dom_node.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 426: def astext(self):
					πF.SetLineno(426)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 427: return reprunicode(unescape(self))
							πF.SetLineno(427)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunescape); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 438: def copy(self):
					πF.SetLineno(438)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 439: return self.__class__(reprunicode(self), rawsource=self.rawsource)
							πF.SetLineno(439)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"rawsource", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 441: def deepcopy(self):
					πF.SetLineno(441)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("deepcopy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 442: return self.copy()
							πF.SetLineno(442)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdeepcopy.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 444: def pformat(self, indent='    ', level=0):
					πF.SetLineno(444)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "indent", Def: πg.NewStr("    ").ToObject()}
					πTemp005[2] = πg.Param{Name: "level", Def: πg.NewInt(0).ToObject()}
					πTemp012 = πg.NewFunction(πg.NewCode("pformat", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindent *πg.Object = πArgs[1]
						_ = µindent
						var µlevel *πg.Object = πArgs[2]
						_ = µlevel
						var µlines *πg.Object = πg.UnboundLocal
						_ = µlines
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 445: indent = indent * level
							πF.SetLineno(445)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µindent, µlevel); πE != nil {
								continue
							}
							µindent = πTemp001
							// line 446: lines = [indent+line for line in self.astext().splitlines()]
							πF.SetLineno(446)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßastext, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
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
											µline = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 446: lines = [indent+line for line in self.astext().splitlines()]
										πF.SetLineno(446)
										if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.Add(πF, µindent, µline); πE != nil {
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µlines = πTemp001
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µlines); πE != nil {
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
							// line 447: if not lines:
							πF.SetLineno(447)
						Label1:
							// line 448: return ''
							πF.SetLineno(448)
							πR = ß.ToObject()
							continue
							goto Label2
						Label2:
							// line 449: return '\n'.join(lines) + '\n'
							πF.SetLineno(449)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp006[0] = µlines
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Add(πF, πTemp007, πg.NewStr("\n").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpformat.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 455: def rstrip(self, chars=None):
					πF.SetLineno(455)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πg.Param{Name: "chars", Def: πTemp014}
					πTemp013 = πg.NewFunction(πg.NewCode("rstrip", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchars *πg.Object = πArgs[1]
						_ = µchars
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
							// line 456: return self.__class__(reprunicode.rstrip(self, chars), self.rawsource)
							πF.SetLineno(456)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µchars, "chars"); πE != nil {
								continue
							}
							πTemp002[1] = µchars
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrstrip.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 458: def lstrip(self, chars=None):
					πF.SetLineno(458)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πg.Param{Name: "chars", Def: πTemp015}
					πTemp014 = πg.NewFunction(πg.NewCode("lstrip", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchars *πg.Object = πArgs[1]
						_ = µchars
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
							// line 459: return self.__class__(reprunicode.lstrip(self, chars), self.rawsource)
							πF.SetLineno(459)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µchars, "chars"); πE != nil {
								continue
							}
							πTemp002[1] = µchars
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreprunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßlstrip.ToObject(), πTemp014); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Text").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßText.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 461: class Element(Node):
			πF.SetLineno(461)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Element", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
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
				var πTemp010 bool
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
				var πTemp060 []*πg.Object
				_ = πTemp060
				var πTemp061 *πg.Object
				_ = πTemp061
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 463: """
					πF.SetLineno(463)
					// line 463: """
					πF.SetLineno(463)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    `Element` is the superclass to all specific elements.\n\n    Elements contain attributes and child nodes.  Elements emulate\n    dictionaries for attributes, indexing by attribute name (a string).  To\n    set the attribute 'att' to 'value', do::\n\n        element['att'] = 'value'\n\n    There are two special attributes: 'ids' and 'names'.  Both are\n    lists of unique identifiers, and names serve as human interfaces\n    to IDs.  Names are case- and whitespace-normalized (see the\n    fully_normalize_name() function), and IDs conform to the regular\n    expression ``[a-z](-?[a-z0-9]+)*`` (see the make_id() function).\n\n    Elements also emulate lists for child nodes (element nodes and/or text\n    nodes), indexing by integer.  To get the first child node, use::\n\n        element[0]\n\n    Elements may be constructed using the ``+=`` operator.  To add one new\n    child node to element, do::\n\n        element += node\n\n    This is equivalent to ``element.append(node)``.\n\n    To add a list of multiple child nodes at once, use the same ``+=``\n    operator::\n\n        element += [node1, node2]\n\n    This is equivalent to ``element.extend([node1, node2])``.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 498: basic_attributes = ('ids', 'classes', 'names', 'dupnames')
					πF.SetLineno(498)
					πTemp001 = πg.NewTuple4(ßids.ToObject(), ßclasses.ToObject(), ßnames.ToObject(), ßdupnames.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßbasic_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 499: """List attributes which are defined for every Element-derived class
					πF.SetLineno(499)
					// line 502: local_attributes = ('backrefs',)
					πF.SetLineno(502)
					πTemp001 = πg.NewTuple1(ßbackrefs.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßlocal_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 503: """A list of class-specific attributes that should not be copied with the
					πF.SetLineno(503)
					// line 509: list_attributes = basic_attributes + local_attributes
					πF.SetLineno(509)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßbasic_attributes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßlocal_attributes); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßlist_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 510: """List attributes, automatically initialized to empty lists for
					πF.SetLineno(510)
					// line 513: known_attributes = list_attributes + ('source', 'rawsource')
					πF.SetLineno(513)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßlist_attributes); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(ßsource.ToObject(), ßrawsource.ToObject()).ToObject()
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßknown_attributes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 514: """List attributes that are known to the Element base class."""
					πF.SetLineno(514)
					// line 516: tagname = None
					πF.SetLineno(516)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtagname.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 517: """The element generic identifier. If None, it is set as an instance
					πF.SetLineno(517)
					// line 520: child_text_separator = '\n\n'
					πF.SetLineno(520)
					if πE = πClass.SetItem(πF, ßchild_text_separator.ToObject(), πg.NewStr("\n\n").ToObject()); πE != nil {
						continue
					}
					// line 521: """Separator for child nodes, used by `astext()` method."""
					πF.SetLineno(521)
					// line 523: def __init__(self, rawsource='', *children, **attributes):
					πF.SetLineno(523)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "rawsource", Def: ß.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrawsource *πg.Object = πArgs[1]
						_ = µrawsource
						var µchildren *πg.Object = πArgs[2]
						_ = µchildren
						var µattributes *πg.Object = πArgs[3]
						_ = µattributes
						var µatt *πg.Object = πg.UnboundLocal
						_ = µatt
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 524: self.rawsource = rawsource
							πF.SetLineno(524)
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrawsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrawsource, πTemp001); πE != nil {
								continue
							}
							// line 525: """The raw text from which this element was constructed.
							πF.SetLineno(525)
							// line 530: self.children = []
							πF.SetLineno(530)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßchildren, πTemp003); πE != nil {
								continue
							}
							// line 531: """List of child nodes (elements and/or `Text`)."""
							πF.SetLineno(531)
							// line 533: self.extend(children)           # maintain parent info
							πF.SetLineno(533)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp002[0] = µchildren
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 535: self.attributes = {}
							πF.SetLineno(535)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßattributes, πTemp003); πE != nil {
								continue
							}
							// line 536: """Dictionary of attribute {name: value}."""
							πF.SetLineno(536)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlist_attributes, nil); πE != nil {
								continue
							}
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
								µatt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 540: self.attributes[att] = []
							πF.SetLineno(540)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp009 = µatt
							if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp007); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µattributes, ßitems, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label6
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
									continue
								}
								µatt = πTemp007
								µvalue = πTemp008
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 543: att = att.lower()
							πF.SetLineno(543)
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µatt, ßlower, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µatt = πTemp007
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßlist_attributes, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp007, µatt); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 544: if att in self.list_attributes:
							πF.SetLineno(544)
						Label7:
							// line 546: self.attributes[att] = value[:]
							πF.SetLineno(546)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µvalue, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp009 = µatt
							if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp003); πE != nil {
								continue
							}
							goto Label9
						Label8:
							// line 548: self.attributes[att] = value
							πF.SetLineno(548)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp008 = µatt
							if πE = πg.SetItem(πF, πTemp007, πTemp008, πTemp003); πE != nil {
								continue
							}
							goto Label9
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 == πTemp007).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 550: if self.tagname is None:
							πF.SetLineno(550)
						Label10:
							// line 551: self.tagname = self.__class__.__name__
							πF.SetLineno(551)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtagname, πTemp001); πE != nil {
								continue
							}
							goto Label11
						Label11:
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
					// line 553: def _dom_node(self, domroot):
					πF.SetLineno(553)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "domroot", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("_dom_node", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdomroot *πg.Object = πArgs[1]
						_ = µdomroot
						var µelement *πg.Object = πg.UnboundLocal
						_ = µelement
						var µattribute *πg.Object = πg.UnboundLocal
						_ = µattribute
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πTemp009 []*πg.Object
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 554: element = domroot.createElement(self.tagname)
							πF.SetLineno(554)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdomroot, ßcreateElement, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µelement = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßattlist, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µattribute = πTemp004
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 556: if isinstance(value, list):
							πF.SetLineno(556)
						Label4:
							// line 557: value = ' '.join([serial_escape('%s' % (v,)) for v in value])
							πF.SetLineno(557)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 557: value = ' '.join([serial_escape('%s' % (v,)) for v in value])
										πF.SetLineno(557)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple1(µv).ToObject()
										if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), πTemp006); πE != nil {
											continue
										}
										πTemp005[0] = πTemp004
										if πTemp004, πE = πg.ResolveGlobal(πF, ßserial_escape); πE != nil {
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
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp007
							goto Label5
						Label5:
							// line 558: element.setAttribute(attribute, '%s' % value)
							πF.SetLineno(558)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattribute, "attribute"); πE != nil {
								continue
							}
							πTemp001[0] = µattribute
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), µvalue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µelement, ßsetAttribute, nil); πE != nil {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp005 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label8
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
								µchild = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 560: element.appendChild(child._dom_node(domroot))
							πF.SetLineno(560)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							πTemp009[0] = µdomroot
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchild, ß_dom_node, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µelement, ßappendChild, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 561: return element
							πF.SetLineno(561)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πR = µelement
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_dom_node.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 563: def __repr__(self):
					πF.SetLineno(563)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πg.UnboundLocal
						_ = µdata
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []πg.Param
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
							// line 564: data = ''
							πF.SetLineno(564)
							µdata = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
								µc = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 566: data += c.shortrepr()
							πF.SetLineno(566)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µc, ßshortrepr, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdata, πTemp005); πE != nil {
								continue
							}
							µdata = πTemp002
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.GT(πF, πTemp007, πg.NewInt(60).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 567: if len(data) > 60:
							πF.SetLineno(567)
						Label4:
							// line 568: data = data[:56] + ' ...'
							πF.SetLineno(568)
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(56).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µdata, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp007, πg.NewStr(" ...").ToObject()); πE != nil {
								continue
							}
							µdata = πTemp002
							// line 569: break
							πF.SetLineno(569)
							πTemp003 = true
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp001 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 570: if self['names']:
							πF.SetLineno(570)
						Label6:
							// line 571: return '<%s "%s": %s>' % (self.__class__.__name__,
							πF.SetLineno(571)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µn *πg.Object = πg.UnboundLocal
								_ = µn
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
								var πTemp006 []*πg.Object
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
										πTemp002 = ßnames.ToObject()
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
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
											µn = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 572: '; '.join([ensure_str(n) for n in self['names']]), data)
										πF.SetLineno(572)
										πTemp006 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πTemp006[0] = µn
										if πTemp002, πE = πg.ResolveGlobal(πF, ßensure_str); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
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
							if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πTemp005, πE = πg.GetAttr(πF, πg.NewStr("; ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp007, πTemp010, µdata).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s \"%s\": %s>").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label8
						Label7:
							// line 574: return '<%s: %s>' % (self.__class__.__name__, data)
							πF.SetLineno(574)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp007, µdata).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s: %s>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 576: def shortrepr(self):
					πF.SetLineno(576)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("shortrepr", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
						var πTemp008 []πg.Param
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
							πTemp001 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 577: if self['names']:
							πF.SetLineno(577)
						Label1:
							// line 578: return '<%s "%s"...>' % (self.__class__.__name__,
							πF.SetLineno(578)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = make([]πg.Param, 0)
							πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µn *πg.Object = πg.UnboundLocal
								_ = µn
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
								var πTemp006 []*πg.Object
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
										πTemp002 = ßnames.ToObject()
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
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
											µn = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 579: '; '.join([ensure_str(n) for n in self['names']]))
										πF.SetLineno(579)
										πTemp006 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πTemp006[0] = µn
										if πTemp002, πE = πg.ResolveGlobal(πF, ßensure_str); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
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
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("; ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πg.NewTuple2(πTemp005, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s \"%s\"...>").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 581: return '<%s...>' % self.tagname
							πF.SetLineno(581)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s...>").ToObject(), πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßshortrepr.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 583: def __unicode__(self):
					πF.SetLineno(583)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__unicode__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 584: if self.children:
							πF.SetLineno(584)
						Label1:
							// line 585: return u'%s%s%s' % (self.starttag(),
							πF.SetLineno(585)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = make([]πg.Param, 0)
							πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µc *πg.Object = πg.UnboundLocal
								_ = µc
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
											µc = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 586: ''.join([unicode(c) for c in self.children]),
										πF.SetLineno(586)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
											continue
										}
										πTemp005[0] = µc
										if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
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
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßendtag, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(πTemp005, πTemp009, πTemp010).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("%s%s%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 589: return self.emptytag()
							πF.SetLineno(589)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßemptytag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp003
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
					if πE = πClass.SetItem(πF, ß__unicode__.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp008 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp007, πE = πg.GE(πF, πTemp009, πTemp008); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label1
					}
					goto Label2
					// line 591: if sys.version_info >= (3, 0):
					πF.SetLineno(591)
				Label1:
					// line 592: __str__ = __unicode__
					πF.SetLineno(592)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß__unicode__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp007); πE != nil {
						continue
					}
					goto Label2
				Label2:
					// line 594: def starttag(self, quoteattr=None):
					πF.SetLineno(594)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[1] = πg.Param{Name: "quoteattr", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("starttag", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µquoteattr *πg.Object = πArgs[1]
						_ = µquoteattr
						var µparts *πg.Object = πg.UnboundLocal
						_ = µparts
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µvalues *πg.Object = πg.UnboundLocal
						_ = µvalues
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
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
							if πE = πg.CheckLocal(πF, µquoteattr, "quoteattr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µquoteattr == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 596: if quoteattr is None:
							πF.SetLineno(596)
						Label1:
							// line 597: quoteattr = pseudo_quoteattr
							πF.SetLineno(597)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßpseudo_quoteattr); πE != nil {
								continue
							}
							µquoteattr = πTemp001
							goto Label2
						Label2:
							// line 598: parts = [self.tagname]
							πF.SetLineno(598)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µparts = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
									continue
								}
								µname = πTemp005
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µvalue == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 600: if value is None:           # boolean attribute
							πF.SetLineno(600)
						Label6:
							// line 601: parts.append('%s="True"' % name)
							πF.SetLineno(601)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s=\"True\"").ToObject(), µname); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 602: continue
							πF.SetLineno(602)
							continue
							goto Label7
						Label7:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 603: if isinstance(value, list):
							πF.SetLineno(603)
						Label8:
							// line 604: values = [serial_escape('%s' % (v,)) for v in value]
							πF.SetLineno(604)
							πTemp008 = make([]πg.Param, 0)
							πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 604: values = [serial_escape('%s' % (v,)) for v in value]
										πF.SetLineno(604)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple1(µv).ToObject()
										if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s").ToObject(), πTemp006); πE != nil {
											continue
										}
										πTemp005[0] = πTemp004
										if πTemp004, πE = πg.ResolveGlobal(πF, ßserial_escape); πE != nil {
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
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							µvalues = πTemp002
							// line 605: value = ' '.join(values)
							πF.SetLineno(605)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp004[0] = µvalues
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µvalue = πTemp007
							goto Label10
						Label9:
							// line 607: value = unicode(value)
							πF.SetLineno(607)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µvalue = πTemp007
							goto Label10
						Label10:
							// line 608: value = quoteattr(value)
							πF.SetLineno(608)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πE = πg.CheckLocal(πF, µquoteattr, "quoteattr"); πE != nil {
								continue
							}
							if πTemp002, πE = µquoteattr.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µvalue = πTemp002
							// line 609: parts.append(u'%s=%s' % (name, value))
							πF.SetLineno(609)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(µname, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("%s=%s").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 610: return u'<%s>' % u' '.join(parts)
							πF.SetLineno(610)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							πTemp004[0] = µparts
							if πTemp002, πE = πg.GetAttr(πF, πg.NewUnicode(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("<%s>").ToObject(), πTemp007); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstarttag.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 612: def endtag(self):
					πF.SetLineno(612)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("endtag", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 613: return '</%s>' % self.tagname
							πF.SetLineno(613)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("</%s>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßendtag.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 615: def emptytag(self):
					πF.SetLineno(615)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("emptytag", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πTemp008 []πg.Param
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
							// line 616: return u'<%s/>' % u' '.join([self.tagname] +
							πF.SetLineno(616)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtagname, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							πTemp008 = make([]πg.Param, 0)
							πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µn *πg.Object = πg.UnboundLocal
								_ = µn
								var µv *πg.Object = πg.UnboundLocal
								_ = µv
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßattlist, nil); πE != nil {
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
											if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
												continue
											}
											µn = πTemp003
											µv = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 617: ['%s="%s"' % (n, v)
										πF.SetLineno(617)
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(µn, µv).ToObject()
										if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s=\"%s\"").ToObject(), πTemp003); πE != nil {
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
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewUnicode(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("<%s/>").ToObject(), πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßemptytag.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 620: def __len__(self):
					πF.SetLineno(620)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 621: return len(self.children)
							πF.SetLineno(621)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 623: def __getitem__(self, key):
					πF.SetLineno(623)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
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
						var πR *πg.Object
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 624: if isinstance(key, basestring):
							πF.SetLineno(624)
						Label1:
							// line 625: return self.attributes[key]
							πF.SetLineno(625)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label5
							// line 626: elif isinstance(key, int):
							πF.SetLineno(626)
						Label2:
							// line 627: return self.children[key]
							πF.SetLineno(627)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label5
							// line 628: elif isinstance(key, slice):
							πF.SetLineno(628)
						Label3:
							// line 629: assert key.step in (None, 1), 'cannot handle slice with stride'
							πF.SetLineno(629)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkey, ßstep, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πg.NewInt(1).ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("cannot handle slice with stride").ToObject()); πE != nil {
								continue
							}
							// line 630: return self.children[key.start:key.stop]
							πF.SetLineno(630)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkey, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µkey, ßstop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label5
						Label4:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("element index must be an integer, a slice, or an attribute name string").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 632: raise TypeError('element index must be an integer, a slice, or '
							πF.SetLineno(632)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 635: def __setitem__(self, key, item):
					πF.SetLineno(635)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					πTemp004[2] = πg.Param{Name: "item", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
						var µitem *πg.Object = πArgs[2]
						_ = µitem
						var µnode *πg.Object = πg.UnboundLocal
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
						var πTemp008 bool
						_ = πTemp008
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 636: if isinstance(key, basestring):
							πF.SetLineno(636)
						Label1:
							// line 637: self.attributes[str(key)] = item
							πF.SetLineno(637)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label5
							// line 638: elif isinstance(key, int):
							πF.SetLineno(638)
						Label2:
							// line 639: self.setup_child(item)
							πF.SetLineno(639)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_child, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 640: self.children[key] = item
							πF.SetLineno(640)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label5
							// line 641: elif isinstance(key, slice):
							πF.SetLineno(641)
						Label3:
							// line 642: assert key.step in (None, 1), 'cannot handle slice with stride'
							πF.SetLineno(642)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkey, ßstep, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πg.NewInt(1).ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("cannot handle slice with stride").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µitem); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label8
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
								µnode = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 644: self.setup_child(node)
							πF.SetLineno(644)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsetup_child, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 645: self.children[key.start:key.stop] = item
							πF.SetLineno(645)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µkey, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µkey, ßstop, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label4:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("element index must be an integer, a slice, or an attribute name string").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 647: raise TypeError('element index must be an integer, a slice, or '
							πF.SetLineno(647)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 650: def __delitem__(self, key):
					πF.SetLineno(650)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
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
						var πR *πg.Object
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 651: if isinstance(key, basestring):
							πF.SetLineno(651)
						Label1:
							// line 652: del self.attributes[key]
							πF.SetLineno(652)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							goto Label5
							// line 653: elif isinstance(key, int):
							πF.SetLineno(653)
						Label2:
							// line 654: del self.children[key]
							πF.SetLineno(654)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							goto Label5
							// line 655: elif isinstance(key, slice):
							πF.SetLineno(655)
						Label3:
							// line 656: assert key.step in (None, 1), 'cannot handle slice with stride'
							πF.SetLineno(656)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkey, ßstep, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πg.NewInt(1).ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("cannot handle slice with stride").ToObject()); πE != nil {
								continue
							}
							// line 657: del self.children[key.start:key.stop]
							πF.SetLineno(657)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µkey, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µkey, ßstop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							goto Label5
						Label4:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("element index must be an integer, a simple slice, or an attribute name string").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 659: raise TypeError('element index must be an integer, a simple '
							πF.SetLineno(659)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 662: def __add__(self, other):
					πF.SetLineno(662)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "other", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 663: return self.children + other
							πF.SetLineno(663)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µother); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 665: def __radd__(self, other):
					πF.SetLineno(665)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "other", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("__radd__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 666: return other + self.children
							πF.SetLineno(666)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µother, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__radd__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 668: def __iadd__(self, other):
					πF.SetLineno(668)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "other", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__iadd__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 669: """Append a node or a list of nodes to `self.children`."""
							πF.SetLineno(669)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µother != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 670: if isinstance(other, Node):
							πF.SetLineno(670)
						Label1:
							// line 671: self.append(other)
							πF.SetLineno(671)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 672: elif other is not None:
							πF.SetLineno(672)
						Label2:
							// line 673: self.extend(other)
							πF.SetLineno(673)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label3:
							// line 674: return self
							πF.SetLineno(674)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iadd__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 669: """Append a node or a list of nodes to `self.children`."""
					πF.SetLineno(669)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("Append a node or a list of nodes to `self.children`.").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ß__iadd__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 676: def astext(self):
					πF.SetLineno(676)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 677: return self.child_text_separator.join(
							πF.SetLineno(677)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
										// line 678: [child.astext() for child in self.children])
										πF.SetLineno(678)
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µchild, ßastext, nil); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchild_text_separator, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 680: def non_default_attributes(self):
					πF.SetLineno(680)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("non_default_attributes", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µatts *πg.Object = πg.UnboundLocal
						_ = µatts
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 681: atts = {}
							πF.SetLineno(681)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µatts = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp008[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßis_not_default, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 683: if self.is_not_default(key):
							πF.SetLineno(683)
						Label4:
							// line 684: atts[key] = value
							πF.SetLineno(684)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = µkey
							if πE = πg.SetItem(πF, µatts, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 685: return atts
							πF.SetLineno(685)
							if πE = πg.CheckLocal(πF, µatts, "atts"); πE != nil {
								continue
							}
							πR = µatts
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnon_default_attributes.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 687: def attlist(self):
					πF.SetLineno(687)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("attlist", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattlist *πg.Object = πg.UnboundLocal
						_ = µattlist
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
							// line 688: attlist = sorted(self.non_default_attributes().items())
							πF.SetLineno(688)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnon_default_attributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µattlist = πTemp003
							// line 689: return attlist
							πF.SetLineno(689)
							if πE = πg.CheckLocal(πF, µattlist, "attlist"); πE != nil {
								continue
							}
							πR = µattlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßattlist.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 691: def get(self, key, failobj=None):
					πF.SetLineno(691)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "failobj", Def: πTemp022}
					πTemp021 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
						var µfailobj *πg.Object = πArgs[2]
						_ = µfailobj
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
							// line 692: return self.attributes.get(key, failobj)
							πF.SetLineno(692)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πE = πg.CheckLocal(πF, µfailobj, "failobj"); πE != nil {
								continue
							}
							πTemp001[1] = µfailobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 694: def hasattr(self, attr):
					πF.SetLineno(694)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("hasattr", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
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
							// line 695: return attr in self.attributes
							πF.SetLineno(695)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µattr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßhasattr.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 697: def delattr(self, attr):
					πF.SetLineno(697)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("delattr", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
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
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µattr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 698: if attr in self.attributes:
							πF.SetLineno(698)
						Label1:
							// line 699: del self.attributes[attr]
							πF.SetLineno(699)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002 = µattr
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdelattr.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 701: def setdefault(self, key, failobj=None):
					πF.SetLineno(701)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "failobj", Def: πTemp025}
					πTemp024 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
						var µfailobj *πg.Object = πArgs[2]
						_ = µfailobj
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
							// line 702: return self.attributes.setdefault(key, failobj)
							πF.SetLineno(702)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πE = πg.CheckLocal(πF, µfailobj, "failobj"); πE != nil {
								continue
							}
							πTemp001[1] = µfailobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßsetdefault.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 704: has_key = hasattr
					πF.SetLineno(704)
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_key.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 707: def __contains__(self, key):
					πF.SetLineno(707)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 710: if isinstance(key, basestring):
							πF.SetLineno(710)
						Label1:
							// line 711: return key in self.attributes
							πF.SetLineno(711)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, µkey); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 712: return key in self.children
							πF.SetLineno(712)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, µkey); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 714: def get_language_code(self, fallback=''):
					πF.SetLineno(714)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "fallback", Def: ß.ToObject()}
					πTemp026 = πg.NewFunction(πg.NewCode("get_language_code", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfallback *πg.Object = πArgs[1]
						_ = µfallback
						var µcls *πg.Object = πg.UnboundLocal
						_ = µcls
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
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
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 715: """Return node's language tag.
							πF.SetLineno(715)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßclasses.ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µcls = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("language-").ToObject()
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcls, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 722: if cls.startswith('language-'):
							πF.SetLineno(722)
						Label4:
							// line 723: return cls[9:]
							πF.SetLineno(723)
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(9).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µcls, πTemp004); πE != nil {
								continue
							}
							πR = πTemp005
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 724: try:
							πF.SetLineno(724)
							πF.PushCheckpoint(7)
							// line 725: return self.parent.get_language(fallback)
							πF.SetLineno(725)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfallback, "fallback"); πE != nil {
								continue
							}
							πTemp002[0] = µfallback
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 726: except AttributeError:
							πF.SetLineno(726)
						Label8:
							// line 727: return fallback
							πF.SetLineno(727)
							if πE = πg.CheckLocal(πF, µfallback, "fallback"); πE != nil {
								continue
							}
							πR = µfallback
							continue
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ßget_language_code.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 715: """Return node's language tag.
					πF.SetLineno(715)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Return node's language tag.\n\n        Look iteratively in self and parents for a class argument\n        starting with ``language-`` and return the remainder of it\n        (which should be a `BCP49` language tag) or the `fallback`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßget_language_code); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 729: def append(self, item):
					πF.SetLineno(729)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "item", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("append", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
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
							// line 730: self.setup_child(item)
							πF.SetLineno(730)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_child, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 731: self.children.append(item)
							πF.SetLineno(731)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 733: def extend(self, item):
					πF.SetLineno(733)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "item", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("extend", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
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
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µitem); πE != nil {
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
								µnode = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 735: self.append(node)
							πF.SetLineno(735)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßextend.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 737: def insert(self, index, item):
					πF.SetLineno(737)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "index", Def: nil}
					πTemp004[2] = πg.Param{Name: "item", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("insert", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindex *πg.Object = πArgs[1]
						_ = µindex
						var µitem *πg.Object = πArgs[2]
						_ = µitem
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
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µitem != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 738: if isinstance(item, Node):
							πF.SetLineno(738)
						Label1:
							// line 739: self.setup_child(item)
							πF.SetLineno(739)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_child, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 740: self.children.insert(index, item)
							πF.SetLineno(740)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[1] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
							// line 741: elif item is not None:
							πF.SetLineno(741)
						Label2:
							// line 742: self[index:index] = item
							πF.SetLineno(742)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µindex, µindex, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µself, πTemp003, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 744: def pop(self, i=-1):
					πF.SetLineno(744)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					if πTemp031, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004[1] = πg.Param{Name: "i", Def: πTemp031}
					πTemp030 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
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
							// line 745: return self.children.pop(i)
							πF.SetLineno(745)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 747: def remove(self, item):
					πF.SetLineno(747)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "item", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
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
							// line 748: self.children.remove(item)
							πF.SetLineno(748)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremove.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 750: def index(self, item):
					πF.SetLineno(750)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "item", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("index", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
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
							// line 751: return self.children.index(item)
							πF.SetLineno(751)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßindex.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 753: def is_not_default(self, key):
					πF.SetLineno(753)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "key", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("is_not_default", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkey *πg.Object = πArgs[1]
						_ = µkey
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							πTemp006 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp006...).ToObject()
							if πTemp003, πE = πg.Eq(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßlist_attributes, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp004, µkey); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 754: if self[key] == [] and key in self.list_attributes:
							πF.SetLineno(754)
						Label2:
							// line 755: return 0
							πF.SetLineno(755)
							πR = πg.NewInt(0).ToObject()
							continue
							goto Label4
						Label3:
							// line 757: return 1
							πF.SetLineno(757)
							πR = πg.NewInt(1).ToObject()
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
					if πE = πClass.SetItem(πF, ßis_not_default.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 759: def update_basic_atts(self, dict_):
					πF.SetLineno(759)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "dict_", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("update_basic_atts", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdict_ *πg.Object = πArgs[1]
						_ = µdict_
						var µatt *πg.Object = πg.UnboundLocal
						_ = µatt
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 760: """
							πF.SetLineno(760)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 764: if isinstance(dict_, Node):
							πF.SetLineno(764)
						Label1:
							// line 765: dict_ = dict_.attributes
							πF.SetLineno(765)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdict_, ßattributes, nil); πE != nil {
								continue
							}
							µdict_ = πTemp002
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbasic_attributes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µatt = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 767: self.append_attr_list(att, dict_.get(att, []))
							πF.SetLineno(767)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp001[0] = µatt
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp006[0] = µatt
							πTemp007 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp006[1] = πTemp003
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdict_, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[1] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend_attr_list, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate_basic_atts.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 760: """
					πF.SetLineno(760)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp035}, πg.NewStr("\n        Update basic attributes ('ids', 'names', 'classes',\n        'dupnames', but not 'source') from node or dictionary `dict_`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßupdate_basic_atts); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp036, ß__doc__, πTemp035); πE != nil {
						continue
					}
					// line 769: def append_attr_list(self, attr, values):
					πF.SetLineno(769)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "values", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("append_attr_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalues *πg.Object = πArgs[2]
						_ = µvalues
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 770: """
							πF.SetLineno(770)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µvalues); πE != nil {
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
								µvalue = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp006 = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µself, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp007, µvalue); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
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
							// line 779: if not value in self[attr]:
							πF.SetLineno(779)
						Label4:
							// line 780: self[attr].append(value)
							πF.SetLineno(780)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp008[0] = µvalue
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004 = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label5
						Label5:
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
					if πE = πClass.SetItem(πF, ßappend_attr_list.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 770: """
					πF.SetLineno(770)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πg.NewStr("\n        For each element in values, if it does not exist in self[attr], append\n        it.\n\n        NOTE: Requires self[attr] and values to be sequence type and the\n        former should specifically be a list.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßappend_attr_list); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp037, ß__doc__, πTemp036); πE != nil {
						continue
					}
					// line 782: def coerce_append_attr_list(self, attr, value):
					πF.SetLineno(782)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("coerce_append_attr_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
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
							// line 783: """
							πF.SetLineno(783)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp003[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 791: if not isinstance(self.get(attr), list):
							πF.SetLineno(791)
						Label1:
							// line 792: self[attr] = [self[attr]]
							πF.SetLineno(792)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp001 = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005 = µattr
							if πE = πg.SetItem(πF, µself, πTemp005, πTemp004); πE != nil {
								continue
							}
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 793: if not isinstance(value, list):
							πF.SetLineno(793)
						Label3:
							// line 794: value = [value]
							πF.SetLineno(794)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µvalue = πTemp001
							goto Label4
						Label4:
							// line 795: self.append_attr_list(attr, value)
							πF.SetLineno(795)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend_attr_list, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcoerce_append_attr_list.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 783: """
					πF.SetLineno(783)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp037}, πg.NewStr("\n        First, convert both self[attr] and value to a non-string sequence\n        type; if either is not already a sequence, convert it to a list of one\n        element.  Then call append_attr_list.\n\n        NOTE: self[attr] and value both must not be None.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßcoerce_append_attr_list); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp038, ß__doc__, πTemp037); πE != nil {
						continue
					}
					// line 797: def replace_attr(self, attr, value, force = True):
					πF.SetLineno(797)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "force", Def: πTemp038}
					πTemp037 = πg.NewFunction(πg.NewCode("replace_attr", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
						var µforce *πg.Object = πArgs[3]
						_ = µforce
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 798: """
							πF.SetLineno(798)
							if πE = πg.CheckLocal(πF, µforce, "force"); πE != nil {
								continue
							}
							πTemp001 = µforce
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006 == πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 803: if force or self.get(attr) is None:
							πF.SetLineno(803)
						Label2:
							// line 804: self[attr] = value
							πF.SetLineno(804)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp003 = µattr
							if πE = πg.SetItem(πF, µself, πTemp003, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreplace_attr.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 798: """
					πF.SetLineno(798)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp038}, πg.NewStr("\n        If self[attr] does not exist or force is True or omitted, set\n        self[attr] to value, otherwise do nothing.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßreplace_attr); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp039, ß__doc__, πTemp038); πE != nil {
						continue
					}
					// line 806: def copy_attr_convert(self, attr, value, replace = True):
					πF.SetLineno(806)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "replace", Def: πTemp039}
					πTemp038 = πg.NewFunction(πg.NewCode("copy_attr_convert", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
						var µreplace *πg.Object = πArgs[3]
						_ = µreplace
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
							// line 807: """
							πF.SetLineno(807)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != µvalue).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 814: if self.get(attr) is not value:
							πF.SetLineno(814)
						Label1:
							// line 815: self.coerce_append_attr_list(attr, value)
							πF.SetLineno(815)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcoerce_append_attr_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßcopy_attr_convert.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 807: """
					πF.SetLineno(807)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp039}, πg.NewStr("\n        If attr is an attribute of self, set self[attr] to\n        [self[attr], value], otherwise set self[attr] to value.\n\n        NOTE: replace is not used by this function and is kept only for\n              compatibility with the other copy functions.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_attr_convert); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp040, ß__doc__, πTemp039); πE != nil {
						continue
					}
					// line 817: def copy_attr_coerce(self, attr, value, replace):
					πF.SetLineno(817)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp004[3] = πg.Param{Name: "replace", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("copy_attr_coerce", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
						var µreplace *πg.Object = πArgs[3]
						_ = µreplace
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
							// line 818: """
							πF.SetLineno(818)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != µvalue).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 826: if self.get(attr) is not value:
							πF.SetLineno(826)
						Label1:
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp006[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 827: if isinstance(self.get(attr), list) or \
							πF.SetLineno(827)
						Label4:
							// line 829: self.coerce_append_attr_list(attr, value)
							πF.SetLineno(829)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcoerce_append_attr_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label6
						Label5:
							// line 831: self.replace_attr(attr, value, replace)
							πF.SetLineno(831)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp002[2] = µreplace
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreplace_attr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ßcopy_attr_coerce.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 818: """
					πF.SetLineno(818)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp040}, πg.NewStr("\n        If attr is an attribute of self and either self[attr] or value is a\n        list, convert all non-sequence values to a sequence of 1 element and\n        then concatenate the two sequence, setting the result to self[attr].\n        If both self[attr] and value are non-sequences and replace is True or\n        self[attr] is None, replace self[attr] with value. Otherwise, do\n        nothing.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_attr_coerce); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp041, ß__doc__, πTemp040); πE != nil {
						continue
					}
					// line 833: def copy_attr_concatenate(self, attr, value, replace):
					πF.SetLineno(833)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp004[3] = πg.Param{Name: "replace", Def: nil}
					πTemp040 = πg.NewFunction(πg.NewCode("copy_attr_concatenate", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
						var µreplace *πg.Object = πArgs[3]
						_ = µreplace
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
							// line 834: """
							πF.SetLineno(834)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != µvalue).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 841: if self.get(attr) is not value:
							πF.SetLineno(841)
						Label1:
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp006[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 842: if isinstance(self.get(attr), list) and \
							πF.SetLineno(842)
						Label4:
							// line 844: self.append_attr_list(attr, value)
							πF.SetLineno(844)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend_attr_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label6
						Label5:
							// line 846: self.replace_attr(attr, value, replace)
							πF.SetLineno(846)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp002[2] = µreplace
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreplace_attr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ßcopy_attr_concatenate.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 834: """
					πF.SetLineno(834)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp041}, πg.NewStr("\n        If attr is an attribute of self and both self[attr] and value are\n        lists, concatenate the two sequences, setting the result to\n        self[attr].  If either self[attr] or value are non-sequences and\n        replace is True or self[attr] is None, replace self[attr] with value.\n        Otherwise, do nothing.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_attr_concatenate); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp042, ß__doc__, πTemp041); πE != nil {
						continue
					}
					// line 848: def copy_attr_consistent(self, attr, value, replace):
					πF.SetLineno(848)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp004[3] = πg.Param{Name: "replace", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("copy_attr_consistent", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattr *πg.Object = πArgs[1]
						_ = µattr
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
						var µreplace *πg.Object = πArgs[3]
						_ = µreplace
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
							// line 849: """
							πF.SetLineno(849)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != µvalue).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 853: if self.get(attr) is not value:
							πF.SetLineno(853)
						Label1:
							// line 854: self.replace_attr(attr, value, replace)
							πF.SetLineno(854)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[0] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[1] = µvalue
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp002[2] = µreplace
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreplace_attr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßcopy_attr_consistent.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 849: """
					πF.SetLineno(849)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp042}, πg.NewStr("\n        If replace is True or self[attr] is None, replace self[attr] with\n        value.  Otherwise, do nothing.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_attr_consistent); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp043, ß__doc__, πTemp042); πE != nil {
						continue
					}
					// line 856: def update_all_atts(self, dict_, update_fun = copy_attr_consistent,
					πF.SetLineno(856)
					πTemp004 = make([]πg.Param, 5)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "dict_", Def: nil}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_attr_consistent); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "update_fun", Def: πTemp043}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "replace", Def: πTemp043}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp004[4] = πg.Param{Name: "and_source", Def: πTemp043}
					πTemp042 = πg.NewFunction(πg.NewCode("update_all_atts", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdict_ *πg.Object = πArgs[1]
						_ = µdict_
						var µupdate_fun *πg.Object = πArgs[2]
						_ = µupdate_fun
						var µreplace *πg.Object = πArgs[3]
						_ = µreplace
						var µand_source *πg.Object = πArgs[4]
						_ = µand_source
						var µfilter_fun *πg.Object = πg.UnboundLocal
						_ = µfilter_fun
						var µatt *πg.Object = πg.UnboundLocal
						_ = µatt
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
							// line 858: """
							πF.SetLineno(858)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 877: if isinstance(dict_, Node):
							πF.SetLineno(877)
						Label1:
							// line 878: dict_ = dict_.attributes
							πF.SetLineno(878)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdict_, ßattributes, nil); πE != nil {
								continue
							}
							µdict_ = πTemp002
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µand_source, "and_source"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µand_source); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 881: if and_source:
							πF.SetLineno(881)
						Label3:
							// line 882: filter_fun = self.is_not_list_attribute
							πF.SetLineno(882)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßis_not_list_attribute, nil); πE != nil {
								continue
							}
							µfilter_fun = πTemp002
							goto Label5
						Label4:
							// line 884: filter_fun = self.is_not_known_attribute
							πF.SetLineno(884)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßis_not_known_attribute, nil); πE != nil {
								continue
							}
							µfilter_fun = πTemp002
							goto Label5
						Label5:
							// line 887: self.update_basic_atts(dict_)
							πF.SetLineno(887)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßupdate_basic_atts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfilter_fun, "filter_fun"); πE != nil {
								continue
							}
							πTemp001[0] = µfilter_fun
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[1] = µdict_
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfilter); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label8
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
								µatt = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 892: update_fun(self, att, dict_[att], replace)
							πF.SetLineno(892)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp001[1] = µatt
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp003 = µatt
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µdict_, πTemp003); πE != nil {
								continue
							}
							πTemp001[2] = πTemp005
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp001[3] = µreplace
							if πE = πg.CheckLocal(πF, µupdate_fun, "update_fun"); πE != nil {
								continue
							}
							if πTemp003, πE = µupdate_fun.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate_all_atts.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 858: """
					πF.SetLineno(858)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp043}, πg.NewStr("\n        Updates all attributes from node or dictionary `dict_`.\n\n        Appends the basic attributes ('ids', 'names', 'classes',\n        'dupnames', but not 'source') and then, for all other attributes in\n        dict_, updates the same attribute in self.  When attributes with the\n        same identifier appear in both self and dict_, the two values are\n        merged based on the value of update_fun.  Generally, when replace is\n        True, the values in self are replaced or merged with the values in\n        dict_; otherwise, the values in self may be preserved or merged.  When\n        and_source is True, the 'source' attribute is included in the copy.\n\n        NOTE: When replace is False, and self contains a 'source' attribute,\n              'source' is not replaced even when dict_ has a 'source'\n              attribute, though it may still be merged into a list depending\n              on the value of update_fun.\n        NOTE: It is easier to call the update-specific methods then to pass\n              the update_fun method to this function.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßupdate_all_atts); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp044, ß__doc__, πTemp043); πE != nil {
						continue
					}
					// line 894: def update_all_atts_consistantly(self, dict_, replace = True,
					πF.SetLineno(894)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "dict_", Def: nil}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "replace", Def: πTemp044}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "and_source", Def: πTemp044}
					πTemp043 = πg.NewFunction(πg.NewCode("update_all_atts_consistantly", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdict_ *πg.Object = πArgs[1]
						_ = µdict_
						var µreplace *πg.Object = πArgs[2]
						_ = µreplace
						var µand_source *πg.Object = πArgs[3]
						_ = µand_source
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
							// line 896: """
							πF.SetLineno(896)
							// line 912: self.update_all_atts(dict_, Element.copy_attr_consistent, replace,
							πF.SetLineno(912)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopy_attr_consistent, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp001[2] = µreplace
							if πE = πg.CheckLocal(πF, µand_source, "and_source"); πE != nil {
								continue
							}
							πTemp001[3] = µand_source
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßupdate_all_atts, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßupdate_all_atts_consistantly.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 896: """
					πF.SetLineno(896)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp044}, πg.NewStr("\n        Updates all attributes from node or dictionary `dict_`.\n\n        Appends the basic attributes ('ids', 'names', 'classes',\n        'dupnames', but not 'source') and then, for all other attributes in\n        dict_, updates the same attribute in self.  When attributes with the\n        same identifier appear in both self and dict_ and replace is True, the\n        values in self are replaced with the values in dict_; otherwise, the\n        values in self are preserved.  When and_source is True, the 'source'\n        attribute is included in the copy.\n\n        NOTE: When replace is False, and self contains a 'source' attribute,\n              'source' is not replaced even when dict_ has a 'source'\n              attribute, though it may still be merged into a list depending\n              on the value of update_fun.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßupdate_all_atts_consistantly); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp045, ß__doc__, πTemp044); πE != nil {
						continue
					}
					// line 915: def update_all_atts_concatenating(self, dict_, replace = True,
					πF.SetLineno(915)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "dict_", Def: nil}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "replace", Def: πTemp045}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "and_source", Def: πTemp045}
					πTemp044 = πg.NewFunction(πg.NewCode("update_all_atts_concatenating", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdict_ *πg.Object = πArgs[1]
						_ = µdict_
						var µreplace *πg.Object = πArgs[2]
						_ = µreplace
						var µand_source *πg.Object = πArgs[3]
						_ = µand_source
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
							// line 917: """
							πF.SetLineno(917)
							// line 936: self.update_all_atts(dict_, Element.copy_attr_concatenate, replace,
							πF.SetLineno(936)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopy_attr_concatenate, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp001[2] = µreplace
							if πE = πg.CheckLocal(πF, µand_source, "and_source"); πE != nil {
								continue
							}
							πTemp001[3] = µand_source
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßupdate_all_atts, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßupdate_all_atts_concatenating.ToObject(), πTemp044); πE != nil {
						continue
					}
					// line 917: """
					πF.SetLineno(917)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp045}, πg.NewStr("\n        Updates all attributes from node or dictionary `dict_`.\n\n        Appends the basic attributes ('ids', 'names', 'classes',\n        'dupnames', but not 'source') and then, for all other attributes in\n        dict_, updates the same attribute in self.  When attributes with the\n        same identifier appear in both self and dict_ whose values aren't each\n        lists and replace is True, the values in self are replaced with the\n        values in dict_; if the values from self and dict_ for the given\n        identifier are both of list type, then the two lists are concatenated\n        and the result stored in self; otherwise, the values in self are\n        preserved.  When and_source is True, the 'source' attribute is\n        included in the copy.\n\n        NOTE: When replace is False, and self contains a 'source' attribute,\n              'source' is not replaced even when dict_ has a 'source'\n              attribute, though it may still be merged into a list depending\n              on the value of update_fun.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßupdate_all_atts_concatenating); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp046, ß__doc__, πTemp045); πE != nil {
						continue
					}
					// line 939: def update_all_atts_coercion(self, dict_, replace = True,
					πF.SetLineno(939)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "dict_", Def: nil}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "replace", Def: πTemp046}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "and_source", Def: πTemp046}
					πTemp045 = πg.NewFunction(πg.NewCode("update_all_atts_coercion", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdict_ *πg.Object = πArgs[1]
						_ = µdict_
						var µreplace *πg.Object = πArgs[2]
						_ = µreplace
						var µand_source *πg.Object = πArgs[3]
						_ = µand_source
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
							// line 941: """
							πF.SetLineno(941)
							// line 961: self.update_all_atts(dict_, Element.copy_attr_coerce, replace,
							πF.SetLineno(961)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopy_attr_coerce, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
								continue
							}
							πTemp001[2] = µreplace
							if πE = πg.CheckLocal(πF, µand_source, "and_source"); πE != nil {
								continue
							}
							πTemp001[3] = µand_source
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßupdate_all_atts, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßupdate_all_atts_coercion.ToObject(), πTemp045); πE != nil {
						continue
					}
					// line 941: """
					πF.SetLineno(941)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp046}, πg.NewStr("\n        Updates all attributes from node or dictionary `dict_`.\n\n        Appends the basic attributes ('ids', 'names', 'classes',\n        'dupnames', but not 'source') and then, for all other attributes in\n        dict_, updates the same attribute in self.  When attributes with the\n        same identifier appear in both self and dict_ whose values are both\n        not lists and replace is True, the values in self are replaced with\n        the values in dict_; if either of the values from self and dict_ for\n        the given identifier are of list type, then first any non-lists are\n        converted to 1-element lists and then the two lists are concatenated\n        and the result stored in self; otherwise, the values in self are\n        preserved.  When and_source is True, the 'source' attribute is\n        included in the copy.\n\n        NOTE: When replace is False, and self contains a 'source' attribute,\n              'source' is not replaced even when dict_ has a 'source'\n              attribute, though it may still be merged into a list depending\n              on the value of update_fun.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßupdate_all_atts_coercion); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp047, ß__doc__, πTemp046); πE != nil {
						continue
					}
					// line 964: def update_all_atts_convert(self, dict_, and_source = False):
					πF.SetLineno(964)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "dict_", Def: nil}
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "and_source", Def: πTemp047}
					πTemp046 = πg.NewFunction(πg.NewCode("update_all_atts_convert", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdict_ *πg.Object = πArgs[1]
						_ = µdict_
						var µand_source *πg.Object = πArgs[2]
						_ = µand_source
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
							// line 965: """
							πF.SetLineno(965)
							// line 982: self.update_all_atts(dict_, Element.copy_attr_convert,
							πF.SetLineno(982)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdict_, "dict_"); πE != nil {
								continue
							}
							πTemp001[0] = µdict_
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopy_attr_convert, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µand_source, "and_source"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"and_source", µand_source},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßupdate_all_atts, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßupdate_all_atts_convert.ToObject(), πTemp046); πE != nil {
						continue
					}
					// line 965: """
					πF.SetLineno(965)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp047}, πg.NewStr("\n        Updates all attributes from node or dictionary `dict_`.\n\n        Appends the basic attributes ('ids', 'names', 'classes',\n        'dupnames', but not 'source') and then, for all other attributes in\n        dict_, updates the same attribute in self.  When attributes with the\n        same identifier appear in both self and dict_ then first any non-lists\n        are converted to 1-element lists and then the two lists are\n        concatenated and the result stored in self; otherwise, the values in\n        self are preserved.  When and_source is True, the 'source' attribute\n        is included in the copy.\n\n        NOTE: When replace is False, and self contains a 'source' attribute,\n              'source' is not replaced even when dict_ has a 'source'\n              attribute, though it may still be merged into a list depending\n              on the value of update_fun.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp048, πE = πg.ResolveClass(πF, πClass, nil, ßupdate_all_atts_convert); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp048, ß__doc__, πTemp047); πE != nil {
						continue
					}
					// line 985: def clear(self):
					πF.SetLineno(985)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp047 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 986: self.children = []
							πF.SetLineno(986)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßchildren, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclear.ToObject(), πTemp047); πE != nil {
						continue
					}
					// line 988: def replace(self, old, new):
					πF.SetLineno(988)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "old", Def: nil}
					πTemp004[2] = πg.Param{Name: "new", Def: nil}
					πTemp048 = πg.NewFunction(πg.NewCode("replace", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µold *πg.Object = πArgs[1]
						_ = µold
						var µnew *πg.Object = πArgs[2]
						_ = µnew
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
							// line 989: """Replace one child `Node` with another child or children."""
							πF.SetLineno(989)
							// line 990: index = self.index(old)
							πF.SetLineno(990)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µold, "old"); πE != nil {
								continue
							}
							πTemp001[0] = µold
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp001[0] = µnew
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µnew != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 991: if isinstance(new, Node):
							πF.SetLineno(991)
						Label1:
							// line 992: self.setup_child(new)
							πF.SetLineno(992)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp001[0] = µnew
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_child, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 993: self[index] = new
							πF.SetLineno(993)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnew); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp003 = µindex
							if πE = πg.SetItem(πF, µself, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label3
							// line 994: elif new is not None:
							πF.SetLineno(994)
						Label2:
							// line 995: self[index:index+1] = new
							πF.SetLineno(995)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnew); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µindex, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µself, πTemp003, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreplace.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 989: """Replace one child `Node` with another child or children."""
					πF.SetLineno(989)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp049}, πg.NewStr("Replace one child `Node` with another child or children.").ToObject()); πE != nil {
						continue
					}
					if πTemp050, πE = πg.ResolveClass(πF, πClass, nil, ßreplace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp050, ß__doc__, πTemp049); πE != nil {
						continue
					}
					// line 997: def replace_self(self, new):
					πF.SetLineno(997)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "new", Def: nil}
					πTemp049 = πg.NewFunction(πg.NewCode("replace_self", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnew *πg.Object = πArgs[1]
						_ = µnew
						var µupdate *πg.Object = πg.UnboundLocal
						_ = µupdate
						var µatt *πg.Object = πg.UnboundLocal
						_ = µatt
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
						var πTemp008 bool
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
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 998: """
							πF.SetLineno(998)
							// line 1002: update = new
							πF.SetLineno(1002)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							µupdate = µnew
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp002[0] = µnew
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
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
							// line 1003: if not isinstance(new, Node):
							πF.SetLineno(1003)
						Label1:
							// line 1005: try:
							πF.SetLineno(1005)
							πF.PushCheckpoint(4)
							// line 1006: update = new[0]
							πF.SetLineno(1006)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnew, πTemp001); πE != nil {
								continue
							}
							µupdate = πTemp003
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
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
							// line 1007: except IndexError:
							πF.SetLineno(1007)
						Label5:
							// line 1008: update = None
							πF.SetLineno(1008)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µupdate = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µupdate, "update"); πE != nil {
								continue
							}
							πTemp002[0] = µupdate
							if πTemp001, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 1009: if isinstance(update, Element):
							πF.SetLineno(1009)
						Label6:
							// line 1010: update.update_basic_atts(self)
							πF.SetLineno(1010)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µupdate, "update"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µupdate, ßupdate_basic_atts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label8
						Label7:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbasic_attributes, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp005 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label11
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µatt = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(9)
							// line 1015: assert not self[att], \
							πF.SetLineno(1015)
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp009 = µatt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µself, πTemp009); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µatt, πTemp010).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Losing \"%s\" attribute: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp009 = µatt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µself, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							if πE = πg.Assert(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							goto Label8
						Label8:
							// line 1017: self.parent.replace(self, new)
							πF.SetLineno(1017)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp002[1] = µnew
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreplace, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreplace_self.ToObject(), πTemp049); πE != nil {
						continue
					}
					// line 998: """
					πF.SetLineno(998)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp050}, πg.NewStr("\n        Replace `self` node with `new`, where `new` is a node or a\n        list of nodes.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßreplace_self); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp051, ß__doc__, πTemp050); πE != nil {
						continue
					}
					// line 1019: def first_child_matching_class(self, childclass, start=0, end=sys.maxsize):
					πF.SetLineno(1019)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "childclass", Def: nil}
					πTemp004[2] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp052, πE = πg.GetAttr(πF, πTemp051, ßmaxsize, nil); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "end", Def: πTemp052}
					πTemp050 = πg.NewFunction(πg.NewCode("first_child_matching_class", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildclass *πg.Object = πArgs[1]
						_ = µchildclass
						var µstart *πg.Object = πArgs[2]
						_ = µstart
						var µend *πg.Object = πArgs[3]
						_ = µend
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 1020: """
							πF.SetLineno(1020)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchildclass, "childclass"); πE != nil {
								continue
							}
							πTemp002[0] = µchildclass
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
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
							// line 1030: if not isinstance(childclass, tuple):
							πF.SetLineno(1030)
						Label1:
							// line 1031: childclass = (childclass,)
							πF.SetLineno(1031)
							if πE = πg.CheckLocal(πF, µchildclass, "childclass"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(µchildclass).ToObject()
							µchildclass = πTemp001
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp002[0] = µstart
							πTemp006 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp007[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp006[1] = µend
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µindex = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µchildclass, "childclass"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µchildclass); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp008 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µc = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp004 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp010
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp002[1] = µc
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp009, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 1034: if isinstance(self[index], c):
							πF.SetLineno(1034)
						Label9:
							// line 1035: return index
							πF.SetLineno(1035)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πR = µindex
							continue
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 1036: return None
							πF.SetLineno(1036)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfirst_child_matching_class.ToObject(), πTemp050); πE != nil {
						continue
					}
					// line 1020: """
					πF.SetLineno(1020)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp051}, πg.NewStr("\n        Return the index of the first child whose class exactly matches.\n\n        Parameters:\n\n        - `childclass`: A `Node` subclass to search for, or a tuple of `Node`\n          classes. If a tuple, any of the classes may match.\n        - `start`: Initial index to check.\n        - `end`: Initial index to *not* check.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßfirst_child_matching_class); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp052, ß__doc__, πTemp051); πE != nil {
						continue
					}
					// line 1038: def first_child_not_matching_class(self, childclass, start=0,
					πF.SetLineno(1038)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "childclass", Def: nil}
					πTemp004[2] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp053, πE = πg.GetAttr(πF, πTemp052, ßmaxsize, nil); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "end", Def: πTemp053}
					πTemp051 = πg.NewFunction(πg.NewCode("first_child_not_matching_class", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildclass *πg.Object = πArgs[1]
						_ = µchildclass
						var µstart *πg.Object = πArgs[2]
						_ = µstart
						var µend *πg.Object = πArgs[3]
						_ = µend
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 1040: """
							πF.SetLineno(1040)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchildclass, "childclass"); πE != nil {
								continue
							}
							πTemp002[0] = µchildclass
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
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
							// line 1050: if not isinstance(childclass, tuple):
							πF.SetLineno(1050)
						Label1:
							// line 1051: childclass = (childclass,)
							πF.SetLineno(1051)
							if πE = πg.CheckLocal(πF, µchildclass, "childclass"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(µchildclass).ToObject()
							µchildclass = πTemp001
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp002[0] = µstart
							πTemp006 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp007[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp006[1] = µend
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µindex = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µchildclass, "childclass"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µchildclass); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp008 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µc = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp004 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp010
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp002[1] = µc
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp009, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 1054: if isinstance(self.children[index], c):
							πF.SetLineno(1054)
						Label9:
							// line 1055: break
							πF.SetLineno(1055)
							πTemp008 = true
							continue
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							// line 1057: return index
							πF.SetLineno(1057)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πR = µindex
							continue
						Label8:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 1058: return None
							πF.SetLineno(1058)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfirst_child_not_matching_class.ToObject(), πTemp051); πE != nil {
						continue
					}
					// line 1040: """
					πF.SetLineno(1040)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp052}, πg.NewStr("\n        Return the index of the first child whose class does *not* match.\n\n        Parameters:\n\n        - `childclass`: A `Node` subclass to skip, or a tuple of `Node`\n          classes. If a tuple, none of the classes may match.\n        - `start`: Initial index to check.\n        - `end`: Initial index to *not* check.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp053, πE = πg.ResolveClass(πF, πClass, nil, ßfirst_child_not_matching_class); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp053, ß__doc__, πTemp052); πE != nil {
						continue
					}
					// line 1060: def pformat(self, indent='    ', level=0):
					πF.SetLineno(1060)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "indent", Def: πg.NewStr("    ").ToObject()}
					πTemp004[2] = πg.Param{Name: "level", Def: πg.NewInt(0).ToObject()}
					πTemp052 = πg.NewFunction(πg.NewCode("pformat", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindent *πg.Object = πArgs[1]
						_ = µindent
						var µlevel *πg.Object = πArgs[2]
						_ = µlevel
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []πg.Param
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
							// line 1061: return ''.join(['%s%s\n' % (indent * level, self.starttag())] +
							πF.SetLineno(1061)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, µindent, µlevel); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πTemp008).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s%s\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp009 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
										// line 1062: [child.pformat(indent, level+1)
										πF.SetLineno(1062)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
											continue
										}
										πTemp005[0] = µindent
										if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
											continue
										}
										πTemp005[1] = πTemp002
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µchild, ßpformat, nil); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
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
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpformat.ToObject(), πTemp052); πE != nil {
						continue
					}
					// line 1065: def copy(self):
					πF.SetLineno(1065)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp053 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πg.UnboundLocal
						_ = µobj
						var πTemp001 *πg.Object
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
							// line 1066: obj = self.__class__(rawsource=self.rawsource, **self.attributes)
							πF.SetLineno(1066)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"rawsource", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, nil, nil, πTemp002, πTemp001); πE != nil {
								continue
							}
							µobj = πTemp004
							// line 1067: obj.document = self.document
							πF.SetLineno(1067)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßdocument, πTemp003); πE != nil {
								continue
							}
							// line 1068: obj.source = self.source
							πF.SetLineno(1068)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßsource, πTemp003); πE != nil {
								continue
							}
							// line 1069: obj.line = self.line
							πF.SetLineno(1069)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßline, πTemp003); πE != nil {
								continue
							}
							// line 1070: return obj
							πF.SetLineno(1070)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp053); πE != nil {
						continue
					}
					// line 1072: def deepcopy(self):
					πF.SetLineno(1072)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp054 = πg.NewFunction(πg.NewCode("deepcopy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcopy *πg.Object = πg.UnboundLocal
						_ = µcopy
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 1073: copy = self.copy()
							πF.SetLineno(1073)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcopy = πTemp002
							// line 1074: copy.extend([child.deepcopy() for child in self.children])
							πF.SetLineno(1074)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
										// line 1074: copy.extend([child.deepcopy() for child in self.children])
										πF.SetLineno(1074)
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µchild, ßdeepcopy, nil); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcopy, "copy"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcopy, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1075: return copy
							πF.SetLineno(1075)
							if πE = πg.CheckLocal(πF, µcopy, "copy"); πE != nil {
								continue
							}
							πR = µcopy
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdeepcopy.ToObject(), πTemp054); πE != nil {
						continue
					}
					// line 1077: def set_class(self, name):
					πF.SetLineno(1077)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "name", Def: nil}
					πTemp055 = πg.NewFunction(πg.NewCode("set_class", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							// line 1078: """Add a new class to the "classes" attribute."""
							πF.SetLineno(1078)
							// line 1079: warnings.warn('docutils.nodes.Element.set_class deprecated; '
							πF.SetLineno(1079)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("docutils.nodes.Element.set_class deprecated; append to Element['classes'] list attribute directly").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp003 = πg.KWArgs{
								{"stacklevel", πg.NewInt(2).ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßwarn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1082: assert ' ' not in name
							πF.SetLineno(1082)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µname, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							// line 1083: self['classes'].append(name.lower())
							πF.SetLineno(1083)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_class.ToObject(), πTemp055); πE != nil {
						continue
					}
					// line 1078: """Add a new class to the "classes" attribute."""
					πF.SetLineno(1078)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp056}, πg.NewStr("Add a new class to the \"classes\" attribute.").ToObject()); πE != nil {
						continue
					}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßset_class); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp057, ß__doc__, πTemp056); πE != nil {
						continue
					}
					// line 1085: def note_referenced_by(self, name=None, id=None):
					πF.SetLineno(1085)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[1] = πg.Param{Name: "name", Def: πTemp057}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "id", Def: πTemp057}
					πTemp056 = πg.NewFunction(πg.NewCode("note_referenced_by", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var µid *πg.Object = πArgs[2]
						_ = µid
						var µby_name *πg.Object = πg.UnboundLocal
						_ = µby_name
						var µby_id *πg.Object = πg.UnboundLocal
						_ = µby_id
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 1086: """Note that this Element has been referenced by its name
							πF.SetLineno(1086)
							// line 1088: self.referenced = 1
							πF.SetLineno(1088)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreferenced, πTemp001); πE != nil {
								continue
							}
							// line 1093: by_name = getattr(self, 'expect_referenced_by_name', {}).get(name)
							πF.SetLineno(1093)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[0] = µname
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							πTemp003[1] = ßexpect_referenced_by_name.ToObject()
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µby_name = πTemp005
							// line 1094: by_id = getattr(self, 'expect_referenced_by_id', {}).get(id)
							πF.SetLineno(1094)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp002[0] = µid
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							πTemp003[1] = ßexpect_referenced_by_id.ToObject()
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µby_id = πTemp005
							if πE = πg.CheckLocal(πF, µby_name, "by_name"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µby_name); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 1095: if by_name:
							πF.SetLineno(1095)
						Label1:
							// line 1096: assert name is not None
							πF.SetLineno(1096)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µname != πTemp005).ToObject()
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 1097: by_name.referenced = 1
							πF.SetLineno(1097)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µby_name, "by_name"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µby_name, ßreferenced, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µby_id, "by_id"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µby_id); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 1098: if by_id:
							πF.SetLineno(1098)
						Label3:
							// line 1099: assert id is not None
							πF.SetLineno(1099)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µid != πTemp005).ToObject()
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 1100: by_id.referenced = 1
							πF.SetLineno(1100)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µby_id, "by_id"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µby_id, ßreferenced, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßnote_referenced_by.ToObject(), πTemp056); πE != nil {
						continue
					}
					// line 1086: """Note that this Element has been referenced by its name
					πF.SetLineno(1086)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp057}, πg.NewStr("Note that this Element has been referenced by its name\n        `name` or id `id`.").ToObject()); πE != nil {
						continue
					}
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßnote_referenced_by); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp058, ß__doc__, πTemp057); πE != nil {
						continue
					}
					// line 1103: def is_not_list_attribute(cls, attr):
					πF.SetLineno(1103)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "cls", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp057 = πg.NewFunction(πg.NewCode("is_not_list_attribute", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]
						_ = µcls
						var µattr *πg.Object = πArgs[1]
						_ = µattr
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
							// line 1104: """
							πF.SetLineno(1104)
							// line 1108: return attr not in cls.list_attributes
							πF.SetLineno(1108)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ßlist_attributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µattr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßis_not_list_attribute.ToObject(), πTemp057); πE != nil {
						continue
					}
					// line 1104: """
					πF.SetLineno(1104)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp058}, πg.NewStr("\n        Returns True if and only if the given attribute is NOT one of the\n        basic list attributes defined for all Elements.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßis_not_list_attribute); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp059, ß__doc__, πTemp058); πE != nil {
						continue
					}
					// line 1102: @classmethod
					πF.SetLineno(1102)
					πTemp060 = πF.MakeArgs(1)
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßis_not_list_attribute); πE != nil {
						continue
					}
					πTemp060[0] = πTemp058
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp059, πE = πTemp058.Call(πF, πTemp060, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp060)
					if πE = πClass.SetItem(πF, ßis_not_list_attribute.ToObject(), πTemp059); πE != nil {
						continue
					}
					// line 1111: def is_not_known_attribute(cls, attr):
					πF.SetLineno(1111)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "cls", Def: nil}
					πTemp004[1] = πg.Param{Name: "attr", Def: nil}
					πTemp058 = πg.NewFunction(πg.NewCode("is_not_known_attribute", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]
						_ = µcls
						var µattr *πg.Object = πArgs[1]
						_ = µattr
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
							// line 1112: """
							πF.SetLineno(1112)
							// line 1116: return attr not in cls.known_attributes
							πF.SetLineno(1116)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ßknown_attributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µattr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßis_not_known_attribute.ToObject(), πTemp058); πE != nil {
						continue
					}
					// line 1112: """
					πF.SetLineno(1112)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp059}, πg.NewStr("\n        Returns True if and only if the given attribute is NOT recognized by\n        this class.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp061, πE = πg.ResolveClass(πF, πClass, nil, ßis_not_known_attribute); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp061, ß__doc__, πTemp059); πE != nil {
						continue
					}
					// line 1110: @classmethod
					πF.SetLineno(1110)
					πTemp060 = πF.MakeArgs(1)
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßis_not_known_attribute); πE != nil {
						continue
					}
					πTemp060[0] = πTemp059
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp061, πE = πTemp059.Call(πF, πTemp060, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp060)
					if πE = πClass.SetItem(πF, ßis_not_known_attribute.ToObject(), πTemp061); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Element").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßElement.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1119: class TextElement(Element):
			πF.SetLineno(1119)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TextElement", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1121: """
					πF.SetLineno(1121)
					// line 1121: """
					πF.SetLineno(1121)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    An element which directly contains text.\n\n    Its children are all `Text` or `Inline` subclass nodes.  You can\n    check whether an element's context is inline simply by checking whether\n    its immediate parent is a `TextElement` instance (including subclasses).\n    This is handy for nodes like `image` that can appear both inline and as\n    standalone body elements.\n\n    If passing children to `__init__()`, make sure to set `text` to\n    ``''`` or some other suitable value.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1134: child_text_separator = ''
					πF.SetLineno(1134)
					if πE = πClass.SetItem(πF, ßchild_text_separator.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					// line 1135: """Separator for child nodes, used by `astext()` method."""
					πF.SetLineno(1135)
					// line 1137: def __init__(self, rawsource='', text='', *children, **attributes):
					πF.SetLineno(1137)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rawsource", Def: ß.ToObject()}
					πTemp002[2] = πg.Param{Name: "text", Def: ß.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrawsource *πg.Object = πArgs[1]
						_ = µrawsource
						var µtext *πg.Object = πArgs[2]
						_ = µtext
						var µchildren *πg.Object = πArgs[3]
						_ = µchildren
						var µattributes *πg.Object = πArgs[4]
						_ = µattributes
						var µtextnode *πg.Object = πg.UnboundLocal
						_ = µtextnode
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
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µtext, ß.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1138: if text != '':
							πF.SetLineno(1138)
						Label1:
							// line 1139: textnode = Text(text)
							πF.SetLineno(1139)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßText); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtextnode = πTemp004
							// line 1140: Element.__init__(self, rawsource, textnode, *children,
							πF.SetLineno(1140)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							πTemp003[1] = µrawsource
							if πE = πg.CheckLocal(πF, µtextnode, "textnode"); πE != nil {
								continue
							}
							πTemp003[2] = µtextnode
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp004, πTemp003, µchildren, nil, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 1143: Element.__init__(self, rawsource, *children, **attributes)
							πF.SetLineno(1143)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							πTemp003[1] = µrawsource
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp004, πTemp003, µchildren, nil, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("TextElement").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTextElement.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1146: class FixedTextElement(TextElement):
			πF.SetLineno(1146)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FixedTextElement", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1148: """An element which directly contains preformatted text."""
					πF.SetLineno(1148)
					// line 1148: """An element which directly contains preformatted text."""
					πF.SetLineno(1148)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("An element which directly contains preformatted text.").ToObject()); πE != nil {
						continue
					}
					// line 1150: def __init__(self, rawsource='', text='', *children, **attributes):
					πF.SetLineno(1150)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rawsource", Def: ß.ToObject()}
					πTemp002[2] = πg.Param{Name: "text", Def: ß.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrawsource *πg.Object = πArgs[1]
						_ = µrawsource
						var µtext *πg.Object = πArgs[2]
						_ = µtext
						var µchildren *πg.Object = πArgs[3]
						_ = µchildren
						var µattributes *πg.Object = πArgs[4]
						_ = µattributes
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
							// line 1151: TextElement.__init__(self, rawsource, text, *children, **attributes)
							πF.SetLineno(1151)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							πTemp001[1] = µrawsource
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[2] = µtext
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µchildren, nil, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1152: self.attributes['xml:space'] = 'preserve'
							πF.SetLineno(1152)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßpreserve.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewStr("xml:space").ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("FixedTextElement").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFixedTextElement.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1159: class Resolvable(object):
			πF.SetLineno(1159)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Resolvable", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1161: resolved = 0
					πF.SetLineno(1161)
					if πE = πClass.SetItem(πF, ßresolved.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Resolvable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßResolvable.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1164: class BackLinkable(object):
			πF.SetLineno(1164)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BackLinkable", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1166: def add_backref(self, refid):
					πF.SetLineno(1166)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "refid", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("add_backref", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrefid *πg.Object = πArgs[1]
						_ = µrefid
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
							// line 1167: self['backrefs'].append(refid)
							πF.SetLineno(1167)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefid, "refid"); πE != nil {
								continue
							}
							πTemp001[0] = µrefid
							πTemp002 = ßbackrefs.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_backref.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("BackLinkable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBackLinkable.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1174: class Root(object):
			πF.SetLineno(1174)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Root", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1175: pass
					πF.SetLineno(1175)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Root").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRoot.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1178: class Titular(object):
			πF.SetLineno(1178)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Titular", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1179: pass
					πF.SetLineno(1179)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Titular").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTitular.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1182: class PreBibliographic(object):
			πF.SetLineno(1182)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PreBibliographic", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1183: """Category of Node which may occur before Bibliographic Nodes."""
					πF.SetLineno(1183)
					// line 1183: """Category of Node which may occur before Bibliographic Nodes."""
					πF.SetLineno(1183)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Category of Node which may occur before Bibliographic Nodes.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("PreBibliographic").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPreBibliographic.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1186: class Bibliographic(object):
			πF.SetLineno(1186)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Bibliographic", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1187: pass
					πF.SetLineno(1187)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Bibliographic").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBibliographic.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1190: class Decorative(PreBibliographic):
			πF.SetLineno(1190)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPreBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Decorative", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1191: pass
					πF.SetLineno(1191)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Decorative").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDecorative.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1194: class Structural(object):
			πF.SetLineno(1194)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Structural", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1195: pass
					πF.SetLineno(1195)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Structural").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStructural.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1198: class Body(object):
			πF.SetLineno(1198)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Body", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1199: pass
					πF.SetLineno(1199)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Body").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBody.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1202: class General(Body):
			πF.SetLineno(1202)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBody); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("General", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1203: pass
					πF.SetLineno(1203)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("General").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGeneral.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1206: class Sequential(Body):
			πF.SetLineno(1206)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBody); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Sequential", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1207: """List-like elements."""
					πF.SetLineno(1207)
					// line 1207: """List-like elements."""
					πF.SetLineno(1207)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("List-like elements.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Sequential").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSequential.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1210: class Admonition(Body): pass
			πF.SetLineno(1210)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBody); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Admonition", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1210: class Admonition(Body): pass
					πF.SetLineno(1210)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Admonition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAdmonition.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1213: class Special(Body):
			πF.SetLineno(1213)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBody); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Special", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1214: """Special internal body elements."""
					πF.SetLineno(1214)
					// line 1214: """Special internal body elements."""
					πF.SetLineno(1214)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Special internal body elements.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Special").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSpecial.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1217: class Invisible(PreBibliographic):
			πF.SetLineno(1217)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPreBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Invisible", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1218: """Internal elements that don't appear in output."""
					πF.SetLineno(1218)
					// line 1218: """Internal elements that don't appear in output."""
					πF.SetLineno(1218)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Internal elements that don't appear in output.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Invisible").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInvisible.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1221: class Part(object):
			πF.SetLineno(1221)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Part", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1222: pass
					πF.SetLineno(1222)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Part").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPart.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1225: class Inline(object):
			πF.SetLineno(1225)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Inline", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1226: pass
					πF.SetLineno(1226)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Inline").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInline.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1229: class Referential(Resolvable):
			πF.SetLineno(1229)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßResolvable); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Referential", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1230: pass
					πF.SetLineno(1230)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Referential").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReferential.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1233: class Targetable(Resolvable):
			πF.SetLineno(1233)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßResolvable); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Targetable", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1235: referenced = 0
					πF.SetLineno(1235)
					if πE = πClass.SetItem(πF, ßreferenced.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 1237: indirect_reference_name = None
					πF.SetLineno(1237)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßindirect_reference_name.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1238: """Holds the whitespace_normalized_name (contains mixed case) of a target.
					πF.SetLineno(1238)
					// line 1238: """Holds the whitespace_normalized_name (contains mixed case) of a target.
					πF.SetLineno(1238)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Holds the whitespace_normalized_name (contains mixed case) of a target.\n    Required for MoinMoin/reST compatibility.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Targetable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTargetable.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1242: class Labeled(object):
			πF.SetLineno(1242)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Labeled", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1243: """Contains a `label` as its first element."""
					πF.SetLineno(1243)
					// line 1243: """Contains a `label` as its first element."""
					πF.SetLineno(1243)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Contains a `label` as its first element.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Labeled").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLabeled.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1250: class document(Root, Structural, Element):
			πF.SetLineno(1250)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßRoot); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßStructural); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("document", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1252: """
					πF.SetLineno(1252)
					// line 1252: """
					πF.SetLineno(1252)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    The document root element.\n\n    Do not instantiate this class directly; use\n    `docutils.utils.new_document()` instead.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1259: def __init__(self, settings, reporter, *args, **kwargs):
					πF.SetLineno(1259)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "settings", Def: nil}
					πTemp002[2] = πg.Param{Name: "reporter", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsettings *πg.Object = πArgs[1]
						_ = µsettings
						var µreporter *πg.Object = πArgs[2]
						_ = µreporter
						var µargs *πg.Object = πArgs[3]
						_ = µargs
						var µkwargs *πg.Object = πArgs[4]
						_ = µkwargs
						var µdocutils *πg.Object = πg.UnboundLocal
						_ = µdocutils
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
							// line 1260: Element.__init__(self, *args, **kwargs)
							πF.SetLineno(1260)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1262: self.current_source = None
							πF.SetLineno(1262)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_source, πTemp003); πE != nil {
								continue
							}
							// line 1263: """Path to or description of the input source being processed."""
							πF.SetLineno(1263)
							// line 1265: self.current_line = None
							πF.SetLineno(1265)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_line, πTemp003); πE != nil {
								continue
							}
							// line 1266: """Line number (1-based) of `current_source`."""
							πF.SetLineno(1266)
							// line 1268: self.settings = settings
							πF.SetLineno(1268)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsettings); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsettings, πTemp002); πE != nil {
								continue
							}
							// line 1269: """Runtime settings data record."""
							πF.SetLineno(1269)
							// line 1271: self.reporter = reporter
							πF.SetLineno(1271)
							if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µreporter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreporter, πTemp002); πE != nil {
								continue
							}
							// line 1272: """System message generator."""
							πF.SetLineno(1272)
							// line 1274: self.indirect_targets = []
							πF.SetLineno(1274)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindirect_targets, πTemp003); πE != nil {
								continue
							}
							// line 1275: """List of indirect target nodes."""
							πF.SetLineno(1275)
							// line 1277: self.substitution_defs = {}
							πF.SetLineno(1277)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsubstitution_defs, πTemp003); πE != nil {
								continue
							}
							// line 1278: """Mapping of substitution names to substitution_definition nodes."""
							πF.SetLineno(1278)
							// line 1280: self.substitution_names = {}
							πF.SetLineno(1280)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsubstitution_names, πTemp003); πE != nil {
								continue
							}
							// line 1281: """Mapping of case-normalized substitution names to case-sensitive
							πF.SetLineno(1281)
							// line 1284: self.refnames = {}
							πF.SetLineno(1284)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrefnames, πTemp003); πE != nil {
								continue
							}
							// line 1285: """Mapping of names to lists of referencing nodes."""
							πF.SetLineno(1285)
							// line 1287: self.refids = {}
							πF.SetLineno(1287)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrefids, πTemp003); πE != nil {
								continue
							}
							// line 1288: """Mapping of ids to lists of referencing nodes."""
							πF.SetLineno(1288)
							// line 1290: self.nameids = {}
							πF.SetLineno(1290)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnameids, πTemp003); πE != nil {
								continue
							}
							// line 1291: """Mapping of names to unique id's."""
							πF.SetLineno(1291)
							// line 1293: self.nametypes = {}
							πF.SetLineno(1293)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnametypes, πTemp003); πE != nil {
								continue
							}
							// line 1294: """Mapping of names to hyperlink type (boolean: True => explicit,
							πF.SetLineno(1294)
							// line 1297: self.ids = {}
							πF.SetLineno(1297)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßids, πTemp003); πE != nil {
								continue
							}
							// line 1298: """Mapping of ids to nodes."""
							πF.SetLineno(1298)
							// line 1300: self.footnote_refs = {}
							πF.SetLineno(1300)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfootnote_refs, πTemp003); πE != nil {
								continue
							}
							// line 1301: """Mapping of footnote labels to lists of footnote_reference nodes."""
							πF.SetLineno(1301)
							// line 1303: self.citation_refs = {}
							πF.SetLineno(1303)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcitation_refs, πTemp003); πE != nil {
								continue
							}
							// line 1304: """Mapping of citation labels to lists of citation_reference nodes."""
							πF.SetLineno(1304)
							// line 1306: self.autofootnotes = []
							πF.SetLineno(1306)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautofootnotes, πTemp003); πE != nil {
								continue
							}
							// line 1307: """List of auto-numbered footnote nodes."""
							πF.SetLineno(1307)
							// line 1309: self.autofootnote_refs = []
							πF.SetLineno(1309)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautofootnote_refs, πTemp003); πE != nil {
								continue
							}
							// line 1310: """List of auto-numbered footnote_reference nodes."""
							πF.SetLineno(1310)
							// line 1312: self.symbol_footnotes = []
							πF.SetLineno(1312)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsymbol_footnotes, πTemp003); πE != nil {
								continue
							}
							// line 1313: """List of symbol footnote nodes."""
							πF.SetLineno(1313)
							// line 1315: self.symbol_footnote_refs = []
							πF.SetLineno(1315)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsymbol_footnote_refs, πTemp003); πE != nil {
								continue
							}
							// line 1316: """List of symbol footnote_reference nodes."""
							πF.SetLineno(1316)
							// line 1318: self.footnotes = []
							πF.SetLineno(1318)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfootnotes, πTemp003); πE != nil {
								continue
							}
							// line 1319: """List of manually-numbered footnote nodes."""
							πF.SetLineno(1319)
							// line 1321: self.citations = []
							πF.SetLineno(1321)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcitations, πTemp003); πE != nil {
								continue
							}
							// line 1322: """List of citation nodes."""
							πF.SetLineno(1322)
							// line 1324: self.autofootnote_start = 1
							πF.SetLineno(1324)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautofootnote_start, πTemp002); πE != nil {
								continue
							}
							// line 1325: """Initial auto-numbered footnote number."""
							πF.SetLineno(1325)
							// line 1327: self.symbol_footnote_start = 0
							πF.SetLineno(1327)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsymbol_footnote_start, πTemp002); πE != nil {
								continue
							}
							// line 1328: """Initial symbol footnote symbol index."""
							πF.SetLineno(1328)
							// line 1330: self.id_counter = Counter()
							πF.SetLineno(1330)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßid_counter, πTemp002); πE != nil {
								continue
							}
							// line 1331: """Numbers added to otherwise identical IDs."""
							πF.SetLineno(1331)
							// line 1333: self.parse_messages = []
							πF.SetLineno(1333)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_messages, πTemp003); πE != nil {
								continue
							}
							// line 1334: """System messages generated while parsing."""
							πF.SetLineno(1334)
							// line 1336: self.transform_messages = []
							πF.SetLineno(1336)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtransform_messages, πTemp003); πE != nil {
								continue
							}
							// line 1337: """System messages generated while applying transforms."""
							πF.SetLineno(1337)
							// line 1339: import docutils.transforms
							πF.SetLineno(1339)
							if πTemp001, πE = πg.ImportModule(πF, "docutils.transforms"); πE != nil {
								continue
							}
							πTemp002 = πTemp001[0]
							µdocutils = πTemp002
							// line 1340: self.transformer = docutils.transforms.Transformer(self)
							πF.SetLineno(1340)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdocutils, "docutils"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocutils, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTransformer, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßtransformer, πTemp003); πE != nil {
								continue
							}
							// line 1341: """Storage for transforms to be applied to this document."""
							πF.SetLineno(1341)
							// line 1343: self.decoration = None
							πF.SetLineno(1343)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdecoration, πTemp003); πE != nil {
								continue
							}
							// line 1344: """Document's `decoration` node."""
							πF.SetLineno(1344)
							// line 1346: self.document = self
							πF.SetLineno(1346)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µself); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp002); πE != nil {
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
					// line 1348: def __getstate__(self):
					πF.SetLineno(1348)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getstate__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate *πg.Object = πg.UnboundLocal
						_ = µstate
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
							// line 1349: """
							πF.SetLineno(1349)
							// line 1352: state = self.__dict__.copy()
							πF.SetLineno(1352)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstate = πTemp001
							// line 1353: state['reporter'] = None
							πF.SetLineno(1353)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003 = ßreporter.ToObject()
							if πE = πg.SetItem(πF, µstate, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 1354: state['transformer'] = None
							πF.SetLineno(1354)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003 = ßtransformer.ToObject()
							if πE = πg.SetItem(πF, µstate, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 1355: return state
							πF.SetLineno(1355)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πR = µstate
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getstate__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1349: """
					πF.SetLineno(1349)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Return dict with unpicklable references removed.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__getstate__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1357: def asdom(self, dom=None):
					πF.SetLineno(1357)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "dom", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("asdom", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdom *πg.Object = πArgs[1]
						_ = µdom
						var µdomroot *πg.Object = πg.UnboundLocal
						_ = µdomroot
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 1358: """Return a DOM representation of this document."""
							πF.SetLineno(1358)
							if πE = πg.CheckLocal(πF, µdom, "dom"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdom == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1359: if dom is None:
							πF.SetLineno(1359)
						Label1:
							// line 1360: import xml.dom.minidom as dom
							πF.SetLineno(1360)
							if πTemp004, πE = πg.ImportModule(πF, "xml.dom.minidom"); πE != nil {
								continue
							}
							πTemp001 = πTemp004[2]
							µdom = πTemp001
							goto Label2
						Label2:
							// line 1361: domroot = dom.Document()
							πF.SetLineno(1361)
							if πE = πg.CheckLocal(πF, µdom, "dom"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdom, ßDocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdomroot = πTemp002
							// line 1362: domroot.appendChild(self._dom_node(domroot))
							πF.SetLineno(1362)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							πTemp005[0] = µdomroot
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_dom_node, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdomroot, ßappendChild, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1363: return domroot
							πF.SetLineno(1363)
							if πE = πg.CheckLocal(πF, µdomroot, "domroot"); πE != nil {
								continue
							}
							πR = µdomroot
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßasdom.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1358: """Return a DOM representation of this document."""
					πF.SetLineno(1358)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a DOM representation of this document.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßasdom); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 1365: def set_id(self, node, msgnode=None, suggested_prefix=''):
					πF.SetLineno(1365)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msgnode", Def: πTemp006}
					πTemp002[3] = πg.Param{Name: "suggested_prefix", Def: ß.ToObject()}
					πTemp005 = πg.NewFunction(πg.NewCode("set_id", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µmsgnode *πg.Object = πArgs[2]
						_ = µmsgnode
						var µsuggested_prefix *πg.Object = πArgs[3]
						_ = µsuggested_prefix
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µid_prefix *πg.Object = πg.UnboundLocal
						_ = µid_prefix
						var µauto_id_prefix *πg.Object = πg.UnboundLocal
						_ = µauto_id_prefix
						var µbase_id *πg.Object = πg.UnboundLocal
						_ = µbase_id
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µprefix *πg.Object = πg.UnboundLocal
						_ = µprefix
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
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
							case 11:
								goto Label11
							case 12:
								goto Label12
							case 24:
								goto Label24
							case 25:
								goto Label25
							default:
								panic("unexpected function state")
							}
							πTemp002 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
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
								µid = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp006, µid); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp006 = µid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008 != µnode).ToObject()
							πTemp002 = πTemp003
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 1367: if id in self.ids and self.ids[id] is not node:
							πF.SetLineno(1367)
						Label5:
							// line 1368: msg = self.reporter.severe('Duplicate ID: "%s".' % id)
							πF.SetLineno(1368)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Duplicate ID: \"%s\".").ToObject(), µid); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µmsg = πTemp002
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µmsgnode, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1369: if msgnode != None:
							πF.SetLineno(1369)
						Label7:
							// line 1370: msgnode += msg
							πF.SetLineno(1370)
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µmsgnode, µmsg); πE != nil {
								continue
							}
							µmsgnode = πTemp002
							goto Label8
						Label8:
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp002 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 1371: if not node['ids']:
							πF.SetLineno(1371)
						Label9:
							// line 1372: id_prefix = self.settings.id_prefix
							πF.SetLineno(1372)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßid_prefix, nil); πE != nil {
								continue
							}
							µid_prefix = πTemp002
							// line 1373: auto_id_prefix = self.settings.auto_id_prefix
							πF.SetLineno(1373)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßauto_id_prefix, nil); πE != nil {
								continue
							}
							µauto_id_prefix = πTemp002
							// line 1374: base_id = ''
							πF.SetLineno(1374)
							µbase_id = ß.ToObject()
							// line 1375: id = ''
							πF.SetLineno(1375)
							µid = ß.ToObject()
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp004 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label13
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
							πF.PushCheckpoint(11)
							// line 1377: base_id = make_id(name)
							πF.SetLineno(1377)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp010[0] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmake_id); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µbase_id = πTemp003
							// line 1378: id = id_prefix + base_id
							πF.SetLineno(1378)
							if πE = πg.CheckLocal(πF, µid_prefix, "id_prefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbase_id, "base_id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µid_prefix, µbase_id); πE != nil {
								continue
							}
							µid = πTemp002
							if πE = πg.CheckLocal(πF, µbase_id, "base_id"); πE != nil {
								continue
							}
							πTemp002 = µbase_id
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp006, µid); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label14:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							goto Label16
							// line 1381: if base_id and id not in self.ids:
							πF.SetLineno(1381)
						Label15:
							// line 1382: break
							πF.SetLineno(1382)
							πTemp004 = true
							continue
							goto Label16
						Label16:
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbase_id, "base_id"); πE != nil {
								continue
							}
							πTemp002 = µbase_id
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label17
							}
							πTemp010 = πF.MakeArgs(1)
							πTemp010[0] = πg.NewStr("%").ToObject()
							if πE = πg.CheckLocal(πF, µauto_id_prefix, "auto_id_prefix"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µauto_id_prefix, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp002 = πTemp006
						Label17:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							goto Label19
							// line 1384: if base_id and auto_id_prefix.endswith('%'):
							πF.SetLineno(1384)
						Label18:
							// line 1387: prefix = id + '-'
							πF.SetLineno(1387)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µid, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							µprefix = πTemp002
							goto Label20
						Label19:
							// line 1389: prefix = id_prefix + auto_id_prefix
							πF.SetLineno(1389)
							if πE = πg.CheckLocal(πF, µid_prefix, "id_prefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µauto_id_prefix, "auto_id_prefix"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µid_prefix, µauto_id_prefix); πE != nil {
								continue
							}
							µprefix = πTemp002
							πTemp010 = πF.MakeArgs(1)
							πTemp010[0] = πg.NewStr("%").ToObject()
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µprefix, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label21
							}
							goto Label22
							// line 1390: if  prefix.endswith('%'):
							πF.SetLineno(1390)
						Label21:
							// line 1391: prefix = '%s%s-' % (prefix[:-1], suggested_prefix
							πF.SetLineno(1391)
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µprefix, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsuggested_prefix, "suggested_prefix"); πE != nil {
								continue
							}
							πTemp006 = µsuggested_prefix
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label23
							}
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µnode, ßtagname, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßmake_id); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp006 = πTemp011
						Label23:
							πTemp003 = πg.NewTuple2(πTemp008, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s-").ToObject(), πTemp003); πE != nil {
								continue
							}
							µprefix = πTemp002
							goto Label22
						Label22:
							goto Label20
						Label20:
							// line 1393: while True:
							πF.SetLineno(1393)
							πF.PushCheckpoint(25)
							πTemp005 = false
						Label24:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label26
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(24)
							// line 1394: self.id_counter[prefix] += 1
							πF.SetLineno(1394)
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							πTemp002 = µprefix
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßid_counter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßid_counter, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							πTemp008 = µprefix
							if πE = πg.SetItem(πF, πTemp006, πTemp008, πTemp002); πE != nil {
								continue
							}
							// line 1395: id = '%s%d' % (prefix, self.id_counter[prefix])
							πF.SetLineno(1395)
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							πTemp006 = µprefix
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßid_counter, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µprefix, πTemp008).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%d").ToObject(), πTemp003); πE != nil {
								continue
							}
							µid = πTemp002
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp003, µid); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label27
							}
							goto Label28
							// line 1396: if id not in self.ids:
							πF.SetLineno(1396)
						Label27:
							// line 1397: break
							πF.SetLineno(1397)
							πTemp005 = true
							continue
							goto Label28
						Label28:
							continue
						Label25:
							if πE != nil || πR != nil {
								continue
							}
						Label26:
						Label13:
							// line 1398: node['ids'].append(id)
							πF.SetLineno(1398)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp010[0] = µid
							πTemp001 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							goto Label10
						Label10:
							// line 1399: self.ids[id] = node
							πF.SetLineno(1399)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp003 = µid
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 1400: return id
							πF.SetLineno(1400)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πR = µid
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßset_id.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1402: def set_name_id_map(self, node, id, msgnode=None, explicit=None):
					πF.SetLineno(1402)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp002[2] = πg.Param{Name: "id", Def: nil}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msgnode", Def: πTemp007}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "explicit", Def: πTemp007}
					πTemp006 = πg.NewFunction(πg.NewCode("set_name_id_map", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µid *πg.Object = πArgs[2]
						_ = µid
						var µmsgnode *πg.Object = πArgs[3]
						_ = µmsgnode
						var µexplicit *πg.Object = πArgs[4]
						_ = µexplicit
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
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
							// line 1403: """
							πF.SetLineno(1403)
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
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
								µname = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 1435: if name in self.nameids:
							πF.SetLineno(1435)
						Label4:
							// line 1436: self.set_duplicate_name_id(node, id, name, msgnode, explicit)
							πF.SetLineno(1436)
							πTemp006 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp006[1] = µid
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[2] = µname
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							πTemp006[3] = µmsgnode
							if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
								continue
							}
							πTemp006[4] = µexplicit
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_duplicate_name_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label6
						Label5:
							// line 1438: self.nameids[name] = id
							πF.SetLineno(1438)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007 = µname
							if πE = πg.SetItem(πF, πTemp003, πTemp007, πTemp002); πE != nil {
								continue
							}
							// line 1439: self.nametypes[name] = explicit
							πF.SetLineno(1439)
							if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µexplicit); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnametypes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007 = µname
							if πE = πg.SetItem(πF, πTemp003, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ßset_name_id_map.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1403: """
					πF.SetLineno(1403)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        `self.nameids` maps names to IDs, while `self.nametypes` maps names to\n        booleans representing hyperlink type (True==explicit,\n        False==implicit).  This method updates the mappings.\n\n        The following state transition table shows how `self.nameids` (\"ids\")\n        and `self.nametypes` (\"types\") change with new input (a call to this\n        method), and what actions are performed (\"implicit\"-type system\n        messages are INFO/1, and \"explicit\"-type system messages are ERROR/3):\n\n        ====  =====  ========  ========  =======  ====  =====  =====\n         Old State    Input          Action        New State   Notes\n        -----------  --------  -----------------  -----------  -----\n        ids   types  new type  sys.msg.  dupname  ids   types\n        ====  =====  ========  ========  =======  ====  =====  =====\n        -     -      explicit  -         -        new   True\n        -     -      implicit  -         -        new   False\n        None  False  explicit  -         -        new   True\n        old   False  explicit  implicit  old      new   True\n        None  True   explicit  explicit  new      None  True\n        old   True   explicit  explicit  new,old  None  True   [#]_\n        None  False  implicit  implicit  new      None  False\n        old   False  implicit  implicit  new,old  None  False\n        None  True   implicit  implicit  new      None  True\n        old   True   implicit  implicit  new      old   True\n        ====  =====  ========  ========  =======  ====  =====  =====\n\n        .. [#] Do not clear the name-to-id map or invalidate the old target if\n           both old and new targets are external and refer to identical URIs.\n           The new target is invalidated regardless.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßset_name_id_map); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 1441: def set_duplicate_name_id(self, node, id, name, msgnode, explicit):
					πF.SetLineno(1441)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp002[2] = πg.Param{Name: "id", Def: nil}
					πTemp002[3] = πg.Param{Name: "name", Def: nil}
					πTemp002[4] = πg.Param{Name: "msgnode", Def: nil}
					πTemp002[5] = πg.Param{Name: "explicit", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("set_duplicate_name_id", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µid *πg.Object = πArgs[2]
						_ = µid
						var µname *πg.Object = πArgs[3]
						_ = µname
						var µmsgnode *πg.Object = πArgs[4]
						_ = µmsgnode
						var µexplicit *πg.Object = πArgs[5]
						_ = µexplicit
						var µold_id *πg.Object = πg.UnboundLocal
						_ = µold_id
						var µold_explicit *πg.Object = πg.UnboundLocal
						_ = µold_explicit
						var µlevel *πg.Object = πg.UnboundLocal
						_ = µlevel
						var µold_node *πg.Object = πg.UnboundLocal
						_ = µold_node
						var µrefuri *πg.Object = πg.UnboundLocal
						_ = µrefuri
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var πTemp001 *πg.Object
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 bool
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
							// line 1442: old_id = self.nameids[name]
							πF.SetLineno(1442)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µold_id = πTemp002
							// line 1443: old_explicit = self.nametypes[name]
							πF.SetLineno(1443)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnametypes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µold_explicit = πTemp002
							// line 1444: self.nametypes[name] = old_explicit or explicit
							πF.SetLineno(1444)
							if πE = πg.CheckLocal(πF, µold_explicit, "old_explicit"); πE != nil {
								continue
							}
							πTemp001 = µold_explicit
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
								continue
							}
							πTemp001 = µexplicit
						Label1:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnametypes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = µname
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µexplicit); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 1445: if explicit:
							πF.SetLineno(1445)
						Label2:
							if πE = πg.CheckLocal(πF, µold_explicit, "old_explicit"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µold_explicit); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 1446: if old_explicit:
							πF.SetLineno(1446)
						Label5:
							// line 1447: level = 2
							πF.SetLineno(1447)
							µlevel = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µold_id != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 1448: if old_id is not None:
							πF.SetLineno(1448)
						Label8:
							// line 1449: old_node = self.ids[old_id]
							πF.SetLineno(1449)
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							πTemp001 = µold_id
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µold_node = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µnode, ßrefuri.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 1450: if 'refuri' in node:
							πF.SetLineno(1450)
						Label10:
							// line 1451: refuri = node['refuri']
							πF.SetLineno(1451)
							πTemp001 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							µrefuri = πTemp002
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µold_node, "old_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µold_node, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µold_node, "old_node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µold_node, ßrefuri.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label12
							}
							πTemp003 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µold_node, "old_node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µold_node, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrefuri, "refuri"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, µrefuri); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label12:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 1452: if old_node['names'] \
							πF.SetLineno(1452)
						Label13:
							// line 1455: level = 1   # just inform if refuri's identical
							πF.SetLineno(1455)
							µlevel = πg.NewInt(1).ToObject()
							goto Label14
						Label14:
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							goto Label16
							// line 1456: if level > 1:
							πF.SetLineno(1456)
						Label15:
							// line 1457: dupname(old_node, name)
							πF.SetLineno(1457)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µold_node, "old_node"); πE != nil {
								continue
							}
							πTemp007[0] = µold_node
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdupname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 1458: self.nameids[name] = None
							πF.SetLineno(1458)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = µname
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label16
						Label16:
							goto Label9
						Label9:
							// line 1459: msg = self.reporter.system_message(
							πF.SetLineno(1459)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp007[0] = µlevel
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Duplicate explicit target name: \"%s\".").ToObject(), µname); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							πTemp008 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp008[0] = µid
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"backrefs", πTemp001},
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmsg = πTemp001
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µmsgnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							goto Label18
							// line 1462: if msgnode != None:
							πF.SetLineno(1462)
						Label17:
							// line 1463: msgnode += msg
							πF.SetLineno(1463)
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µmsgnode, µmsg); πE != nil {
								continue
							}
							µmsgnode = πTemp001
							goto Label18
						Label18:
							// line 1464: dupname(node, name)
							πF.SetLineno(1464)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdupname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label7
						Label6:
							// line 1466: self.nameids[name] = id
							πF.SetLineno(1466)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003 = µname
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µold_id != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label19
							}
							goto Label20
							// line 1467: if old_id is not None:
							πF.SetLineno(1467)
						Label19:
							// line 1468: old_node = self.ids[old_id]
							πF.SetLineno(1468)
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							πTemp001 = µold_id
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µold_node = πTemp002
							// line 1469: dupname(old_node, name)
							πF.SetLineno(1469)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µold_node, "old_node"); πE != nil {
								continue
							}
							πTemp007[0] = µold_node
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdupname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label20
						Label20:
							goto Label7
						Label7:
							goto Label4
						Label3:
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µold_id != πTemp003).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label21
							}
							if πE = πg.CheckLocal(πF, µold_explicit, "old_explicit"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µold_explicit); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp002
						Label21:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label22
							}
							goto Label23
							// line 1471: if old_id is not None and not old_explicit:
							πF.SetLineno(1471)
						Label22:
							// line 1472: self.nameids[name] = None
							πF.SetLineno(1472)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = µname
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 1473: old_node = self.ids[old_id]
							πF.SetLineno(1473)
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							πTemp001 = µold_id
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßids, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µold_node = πTemp002
							// line 1474: dupname(old_node, name)
							πF.SetLineno(1474)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µold_node, "old_node"); πE != nil {
								continue
							}
							πTemp007[0] = µold_node
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdupname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label23
						Label23:
							// line 1475: dupname(node, name)
							πF.SetLineno(1475)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdupname); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µexplicit, "explicit"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µexplicit); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label24
							}
							if πE = πg.CheckLocal(πF, µold_explicit, "old_explicit"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, µold_explicit); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp010).ToObject()
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label25
							}
							if πE = πg.CheckLocal(πF, µold_id, "old_id"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µold_id != πTemp005).ToObject()
							πTemp002 = πTemp003
						Label25:
							πTemp001 = πTemp002
						Label24:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label26
							}
							goto Label27
							// line 1476: if not explicit or (not old_explicit and old_id is not None):
							πF.SetLineno(1476)
						Label26:
							// line 1477: msg = self.reporter.info(
							πF.SetLineno(1477)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Duplicate implicit target name: \"%s\".").ToObject(), µname); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							πTemp008 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp008[0] = µid
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"backrefs", πTemp001},
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmsg = πTemp001
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µmsgnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label28
							}
							goto Label29
							// line 1480: if msgnode != None:
							πF.SetLineno(1480)
						Label28:
							// line 1481: msgnode += msg
							πF.SetLineno(1481)
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µmsgnode, µmsg); πE != nil {
								continue
							}
							µmsgnode = πTemp001
							goto Label29
						Label29:
							goto Label27
						Label27:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßset_duplicate_name_id.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1483: def has_name(self, name):
					πF.SetLineno(1483)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("has_name", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
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
							// line 1484: return name in self.nameids
							πF.SetLineno(1484)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µname); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßhas_name.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1487: def note_implicit_target(self, target, msgnode=None):
					πF.SetLineno(1487)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msgnode", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("note_implicit_target", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var µmsgnode *πg.Object = πArgs[2]
						_ = µmsgnode
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
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
							// line 1488: id = self.set_id(target, msgnode)
							πF.SetLineno(1488)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							πTemp001[1] = µmsgnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µid = πTemp003
							// line 1489: self.set_name_id_map(target, id, msgnode, explicit=None)
							πF.SetLineno(1489)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp001[1] = µid
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							πTemp001[2] = µmsgnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"explicit", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_name_id_map, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_implicit_target.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1491: def note_explicit_target(self, target, msgnode=None):
					πF.SetLineno(1491)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msgnode", Def: πTemp011}
					πTemp010 = πg.NewFunction(πg.NewCode("note_explicit_target", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var µmsgnode *πg.Object = πArgs[2]
						_ = µmsgnode
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
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
							// line 1492: id = self.set_id(target, msgnode)
							πF.SetLineno(1492)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							πTemp001[1] = µmsgnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µid = πTemp003
							// line 1493: self.set_name_id_map(target, id, msgnode, explicit=True)
							πF.SetLineno(1493)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp001[1] = µid
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							πTemp001[2] = µmsgnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"explicit", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_name_id_map, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_explicit_target.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1495: def note_refname(self, node):
					πF.SetLineno(1495)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("note_refname", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1496: self.refnames.setdefault(node['refname'], []).append(node)
							πF.SetLineno(1496)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrefnames, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßnote_refname.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1498: def note_refid(self, node):
					πF.SetLineno(1498)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("note_refid", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1499: self.refids.setdefault(node['refid'], []).append(node)
							πF.SetLineno(1499)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrefids, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßnote_refid.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1501: def note_indirect_target(self, target):
					πF.SetLineno(1501)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("note_indirect_target", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
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
							// line 1502: self.indirect_targets.append(target)
							πF.SetLineno(1502)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindirect_targets, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1503: if target['names']:
							πF.SetLineno(1503)
						Label1:
							// line 1504: self.note_refname(target)
							πF.SetLineno(1504)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnote_refname, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_indirect_target.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1506: def note_anonymous_target(self, target):
					πF.SetLineno(1506)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("note_anonymous_target", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
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
							// line 1507: self.set_id(target)
							πF.SetLineno(1507)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_anonymous_target.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1509: def note_autofootnote(self, footnote):
					πF.SetLineno(1509)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "footnote", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("note_autofootnote", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfootnote *πg.Object = πArgs[1]
						_ = µfootnote
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
							// line 1510: self.set_id(footnote)
							πF.SetLineno(1510)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1511: self.autofootnotes.append(footnote)
							πF.SetLineno(1511)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßautofootnotes, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_autofootnote.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1513: def note_autofootnote_ref(self, ref):
					πF.SetLineno(1513)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ref", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("note_autofootnote_ref", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µref *πg.Object = πArgs[1]
						_ = µref
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
							// line 1514: self.set_id(ref)
							πF.SetLineno(1514)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1515: self.autofootnote_refs.append(ref)
							πF.SetLineno(1515)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßautofootnote_refs, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_autofootnote_ref.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1517: def note_symbol_footnote(self, footnote):
					πF.SetLineno(1517)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "footnote", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("note_symbol_footnote", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfootnote *πg.Object = πArgs[1]
						_ = µfootnote
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
							// line 1518: self.set_id(footnote)
							πF.SetLineno(1518)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1519: self.symbol_footnotes.append(footnote)
							πF.SetLineno(1519)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsymbol_footnotes, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_symbol_footnote.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 1521: def note_symbol_footnote_ref(self, ref):
					πF.SetLineno(1521)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ref", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("note_symbol_footnote_ref", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µref *πg.Object = πArgs[1]
						_ = µref
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
							// line 1522: self.set_id(ref)
							πF.SetLineno(1522)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1523: self.symbol_footnote_refs.append(ref)
							πF.SetLineno(1523)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsymbol_footnote_refs, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_symbol_footnote_ref.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 1525: def note_footnote(self, footnote):
					πF.SetLineno(1525)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "footnote", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("note_footnote", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfootnote *πg.Object = πArgs[1]
						_ = µfootnote
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
							// line 1526: self.set_id(footnote)
							πF.SetLineno(1526)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1527: self.footnotes.append(footnote)
							πF.SetLineno(1527)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfootnotes, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_footnote.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1529: def note_footnote_ref(self, ref):
					πF.SetLineno(1529)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ref", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("note_footnote_ref", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µref *πg.Object = πArgs[1]
						_ = µref
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 1530: self.set_id(ref)
							πF.SetLineno(1530)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1531: self.footnote_refs.setdefault(ref['refname'], []).append(ref)
							πF.SetLineno(1531)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							πTemp004 = πF.MakeArgs(2)
							πTemp002 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µref, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfootnote_refs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1532: self.note_refname(ref)
							πF.SetLineno(1532)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnote_refname, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_footnote_ref.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1534: def note_citation(self, citation):
					πF.SetLineno(1534)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "citation", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("note_citation", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcitation *πg.Object = πArgs[1]
						_ = µcitation
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
							// line 1535: self.citations.append(citation)
							πF.SetLineno(1535)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcitation, "citation"); πE != nil {
								continue
							}
							πTemp001[0] = µcitation
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcitations, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_citation.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1537: def note_citation_ref(self, ref):
					πF.SetLineno(1537)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ref", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("note_citation_ref", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µref *πg.Object = πArgs[1]
						_ = µref
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 1538: self.set_id(ref)
							πF.SetLineno(1538)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1539: self.citation_refs.setdefault(ref['refname'], []).append(ref)
							πF.SetLineno(1539)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							πTemp004 = πF.MakeArgs(2)
							πTemp002 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µref, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcitation_refs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1540: self.note_refname(ref)
							πF.SetLineno(1540)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnote_refname, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_citation_ref.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1542: def note_substitution_def(self, subdef, def_name, msgnode=None):
					πF.SetLineno(1542)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "subdef", Def: nil}
					πTemp002[2] = πg.Param{Name: "def_name", Def: nil}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msgnode", Def: πTemp024}
					πTemp023 = πg.NewFunction(πg.NewCode("note_substitution_def", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsubdef *πg.Object = πArgs[1]
						_ = µsubdef
						var µdef_name *πg.Object = πArgs[2]
						_ = µdef_name
						var µmsgnode *πg.Object = πArgs[3]
						_ = µmsgnode
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µoldnode *πg.Object = πg.UnboundLocal
						_ = µoldnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 1543: name = whitespace_normalize_name(def_name)
							πF.SetLineno(1543)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdef_name, "def_name"); πE != nil {
								continue
							}
							πTemp001[0] = µdef_name
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwhitespace_normalize_name); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µname = πTemp003
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsubstitution_defs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1544: if name in self.substitution_defs:
							πF.SetLineno(1544)
						Label1:
							// line 1545: msg = self.reporter.error(
							πF.SetLineno(1545)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Duplicate substitution definition name: \"%s\".").ToObject(), µname); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"base_node", µsubdef},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp002
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µmsgnode, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1548: if msgnode != None:
							πF.SetLineno(1548)
						Label3:
							// line 1549: msgnode += msg
							πF.SetLineno(1549)
							if πE = πg.CheckLocal(πF, µmsgnode, "msgnode"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µmsgnode, µmsg); πE != nil {
								continue
							}
							µmsgnode = πTemp002
							goto Label4
						Label4:
							// line 1550: oldnode = self.substitution_defs[name]
							πF.SetLineno(1550)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsubstitution_defs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							µoldnode = πTemp003
							// line 1551: dupname(oldnode, name)
							πF.SetLineno(1551)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µoldnode, "oldnode"); πE != nil {
								continue
							}
							πTemp001[0] = µoldnode
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdupname); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 1553: self.substitution_defs[name] = subdef
							πF.SetLineno(1553)
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsubdef); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsubstitution_defs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006 = µname
							if πE = πg.SetItem(πF, πTemp003, πTemp006, πTemp002); πE != nil {
								continue
							}
							// line 1555: self.substitution_names[fully_normalize_name(name)] = name
							πF.SetLineno(1555)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µname); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsubstitution_names, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πTemp007, πE = πg.ResolveGlobal(πF, ßfully_normalize_name); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp006 = πTemp008
							if πE = πg.SetItem(πF, πTemp003, πTemp006, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_substitution_def.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1557: def note_substitution_ref(self, subref, refname):
					πF.SetLineno(1557)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "subref", Def: nil}
					πTemp002[2] = πg.Param{Name: "refname", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("note_substitution_ref", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsubref *πg.Object = πArgs[1]
						_ = µsubref
						var µrefname *πg.Object = πArgs[2]
						_ = µrefname
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
							// line 1558: subref['refname'] = whitespace_normalize_name(refname)
							πF.SetLineno(1558)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							πTemp001[0] = µrefname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwhitespace_normalize_name); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubref, "subref"); πE != nil {
								continue
							}
							πTemp004 = ßrefname.ToObject()
							if πE = πg.SetItem(πF, µsubref, πTemp004, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_substitution_ref.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1560: def note_pending(self, pending, priority=None):
					πF.SetLineno(1560)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pending", Def: nil}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "priority", Def: πTemp026}
					πTemp025 = πg.NewFunction(πg.NewCode("note_pending", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πArgs[1]
						_ = µpending
						var µpriority *πg.Object = πArgs[2]
						_ = µpriority
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
							// line 1561: self.transformer.add_pending(pending, priority)
							πF.SetLineno(1561)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp001[1] = µpriority
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransformer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadd_pending, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_pending.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 1563: def note_parse_message(self, message):
					πF.SetLineno(1563)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("note_parse_message", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 1564: self.parse_messages.append(message)
							πF.SetLineno(1564)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[0] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_messages, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_parse_message.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 1566: def note_transform_message(self, message):
					πF.SetLineno(1566)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("note_transform_message", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 1567: self.transform_messages.append(message)
							πF.SetLineno(1567)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[0] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransform_messages, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_transform_message.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 1569: def note_source(self, source, offset):
					πF.SetLineno(1569)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "source", Def: nil}
					πTemp002[2] = πg.Param{Name: "offset", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("note_source", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πArgs[1]
						_ = µsource
						var µoffset *πg.Object = πArgs[2]
						_ = µoffset
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
							// line 1570: self.current_source = source
							πF.SetLineno(1570)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_source, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µoffset == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1571: if offset is None:
							πF.SetLineno(1571)
						Label1:
							// line 1572: self.current_line = offset
							πF.SetLineno(1572)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µoffset); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_line, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 1574: self.current_line = offset + 1
							πF.SetLineno(1574)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_line, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnote_source.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1576: def copy(self):
					πF.SetLineno(1576)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πg.UnboundLocal
						_ = µobj
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
							// line 1577: obj = self.__class__(self.settings, self.reporter,
							πF.SetLineno(1577)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp001, nil, nil, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobj = πTemp004
							// line 1579: obj.source = self.source
							πF.SetLineno(1579)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßsource, πTemp003); πE != nil {
								continue
							}
							// line 1580: obj.line = self.line
							πF.SetLineno(1580)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßline, πTemp003); πE != nil {
								continue
							}
							// line 1581: return obj
							πF.SetLineno(1581)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1583: def get_decoration(self):
					πF.SetLineno(1583)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("get_decoration", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdecoration, nil); πE != nil {
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
							// line 1584: if not self.decoration:
							πF.SetLineno(1584)
						Label1:
							// line 1585: self.decoration = decoration()
							πF.SetLineno(1585)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdecoration); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßdecoration, πTemp001); πE != nil {
								continue
							}
							// line 1586: index = self.first_child_not_matching_class(Titular)
							πF.SetLineno(1586)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTitular); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfirst_child_not_matching_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µindex = πTemp002
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µindex == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1587: if index is None:
							πF.SetLineno(1587)
						Label3:
							// line 1588: self.append(self.decoration)
							πF.SetLineno(1588)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdecoration, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label4:
							// line 1590: self.insert(index, self.decoration)
							πF.SetLineno(1590)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp004[0] = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdecoration, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 1591: return self.decoration
							πF.SetLineno(1591)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdecoration, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_decoration.ToObject(), πTemp030); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("document").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdocument.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1598: class title(Titular, PreBibliographic, TextElement): pass
			πF.SetLineno(1598)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTitular); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPreBibliographic); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("title", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1598: class title(Titular, PreBibliographic, TextElement): pass
					πF.SetLineno(1598)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("title").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtitle.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1599: class subtitle(Titular, PreBibliographic, TextElement): pass
			πF.SetLineno(1599)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTitular); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPreBibliographic); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("subtitle", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1599: class subtitle(Titular, PreBibliographic, TextElement): pass
					πF.SetLineno(1599)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("subtitle").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsubtitle.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1600: class rubric(Titular, TextElement): pass
			πF.SetLineno(1600)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTitular); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("rubric", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1600: class rubric(Titular, TextElement): pass
					πF.SetLineno(1600)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("rubric").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrubric.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1607: class docinfo(Bibliographic, Element): pass
			πF.SetLineno(1607)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("docinfo", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1607: class docinfo(Bibliographic, Element): pass
					πF.SetLineno(1607)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("docinfo").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdocinfo.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1608: class author(Bibliographic, TextElement): pass
			πF.SetLineno(1608)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("author", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1608: class author(Bibliographic, TextElement): pass
					πF.SetLineno(1608)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("author").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßauthor.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1609: class authors(Bibliographic, Element): pass
			πF.SetLineno(1609)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("authors", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1609: class authors(Bibliographic, Element): pass
					πF.SetLineno(1609)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("authors").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßauthors.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1610: class organization(Bibliographic, TextElement): pass
			πF.SetLineno(1610)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("organization", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1610: class organization(Bibliographic, TextElement): pass
					πF.SetLineno(1610)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("organization").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßorganization.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1611: class address(Bibliographic, FixedTextElement): pass
			πF.SetLineno(1611)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFixedTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("address", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1611: class address(Bibliographic, FixedTextElement): pass
					πF.SetLineno(1611)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("address").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßaddress.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1612: class contact(Bibliographic, TextElement): pass
			πF.SetLineno(1612)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("contact", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1612: class contact(Bibliographic, TextElement): pass
					πF.SetLineno(1612)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("contact").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcontact.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1613: class version(Bibliographic, TextElement): pass
			πF.SetLineno(1613)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("version", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1613: class version(Bibliographic, TextElement): pass
					πF.SetLineno(1613)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("version").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßversion.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1614: class revision(Bibliographic, TextElement): pass
			πF.SetLineno(1614)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("revision", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1614: class revision(Bibliographic, TextElement): pass
					πF.SetLineno(1614)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("revision").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrevision.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1615: class status(Bibliographic, TextElement): pass
			πF.SetLineno(1615)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("status", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1615: class status(Bibliographic, TextElement): pass
					πF.SetLineno(1615)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("status").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstatus.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1616: class date(Bibliographic, TextElement): pass
			πF.SetLineno(1616)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("date", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1616: class date(Bibliographic, TextElement): pass
					πF.SetLineno(1616)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("date").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdate.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1617: class copyright(Bibliographic, TextElement): pass
			πF.SetLineno(1617)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBibliographic); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("copyright", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1617: class copyright(Bibliographic, TextElement): pass
					πF.SetLineno(1617)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("copyright").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcopyright.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1624: class decoration(Decorative, Element):
			πF.SetLineno(1624)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßDecorative); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("decoration", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 1626: def get_header(self):
					πF.SetLineno(1626)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_header", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp007 bool
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
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßheader); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 1627: if not len(self.children) or not isinstance(self.children[0], header):
							πF.SetLineno(1627)
						Label2:
							// line 1628: self.insert(0, header())
							πF.SetLineno(1628)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßheader); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label3:
							// line 1629: return self.children[0]
							πF.SetLineno(1629)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßget_header.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1631: def get_footer(self):
					πF.SetLineno(1631)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("get_footer", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp007 bool
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
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
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
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(2)
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßfooter); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 1632: if not len(self.children) or not isinstance(self.children[-1], footer):
							πF.SetLineno(1632)
						Label2:
							// line 1633: self.append(footer())
							πF.SetLineno(1633)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßfooter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label3:
							// line 1634: return self.children[-1]
							πF.SetLineno(1634)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßget_footer.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("decoration").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdecoration.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1637: class header(Decorative, Element): pass
			πF.SetLineno(1637)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßDecorative); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("header", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1637: class header(Decorative, Element): pass
					πF.SetLineno(1637)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("header").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßheader.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1638: class footer(Decorative, Element): pass
			πF.SetLineno(1638)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßDecorative); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("footer", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1638: class footer(Decorative, Element): pass
					πF.SetLineno(1638)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("footer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfooter.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1645: class section(Structural, Element): pass
			πF.SetLineno(1645)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßStructural); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("section", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1645: class section(Structural, Element): pass
					πF.SetLineno(1645)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("section").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsection.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1648: class topic(Structural, Element):
			πF.SetLineno(1648)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßStructural); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("topic", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1650: """
					πF.SetLineno(1650)
					// line 1650: """
					πF.SetLineno(1650)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Topics are terminal, \"leaf\" mini-sections, like block quotes with titles,\n    or textual figures.  A topic is just like a section, except that it has no\n    subsections, and it doesn't have to conform to section placement rules.\n\n    Topics are allowed wherever body elements (list, table, etc.) are allowed,\n    but only at the top level of a section or document.  Topics cannot nest\n    inside topics, sidebars, or body elements; you can't have a topic inside a\n    table, list, block quote, etc.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("topic").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtopic.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1662: class sidebar(Structural, Element):
			πF.SetLineno(1662)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßStructural); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("sidebar", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1664: """
					πF.SetLineno(1664)
					// line 1664: """
					πF.SetLineno(1664)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Sidebars are like miniature, parallel documents that occur inside other\n    documents, providing related or reference material.  A sidebar is\n    typically offset by a border and \"floats\" to the side of the page; the\n    document's main text may flow around it.  Sidebars can also be likened to\n    super-footnotes; their content is outside of the flow of the document's\n    main text.\n\n    Sidebars are allowed wherever body elements (list, table, etc.) are\n    allowed, but only at the top level of a section or document.  Sidebars\n    cannot nest inside sidebars, topics, or body elements; you can't have a\n    sidebar inside a table, list, block quote, etc.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("sidebar").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsidebar.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1679: class transition(Structural, Element): pass
			πF.SetLineno(1679)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßStructural); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("transition", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1679: class transition(Structural, Element): pass
					πF.SetLineno(1679)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("transition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtransition.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1686: class paragraph(General, TextElement): pass
			πF.SetLineno(1686)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("paragraph", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1686: class paragraph(General, TextElement): pass
					πF.SetLineno(1686)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("paragraph").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßparagraph.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1687: class compound(General, Element): pass
			πF.SetLineno(1687)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("compound", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1687: class compound(General, Element): pass
					πF.SetLineno(1687)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("compound").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcompound.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1688: class container(General, Element): pass
			πF.SetLineno(1688)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("container", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1688: class container(General, Element): pass
					πF.SetLineno(1688)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("container").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcontainer.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1689: class bullet_list(Sequential, Element): pass
			πF.SetLineno(1689)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSequential); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("bullet_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1689: class bullet_list(Sequential, Element): pass
					πF.SetLineno(1689)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("bullet_list").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßbullet_list.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1690: class enumerated_list(Sequential, Element): pass
			πF.SetLineno(1690)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSequential); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("enumerated_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1690: class enumerated_list(Sequential, Element): pass
					πF.SetLineno(1690)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("enumerated_list").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßenumerated_list.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1691: class list_item(Part, Element): pass
			πF.SetLineno(1691)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("list_item", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1691: class list_item(Part, Element): pass
					πF.SetLineno(1691)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("list_item").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlist_item.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1692: class definition_list(Sequential, Element): pass
			πF.SetLineno(1692)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSequential); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("definition_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1692: class definition_list(Sequential, Element): pass
					πF.SetLineno(1692)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("definition_list").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefinition_list.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1693: class definition_list_item(Part, Element): pass
			πF.SetLineno(1693)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("definition_list_item", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1693: class definition_list_item(Part, Element): pass
					πF.SetLineno(1693)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("definition_list_item").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefinition_list_item.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1694: class term(Part, TextElement): pass
			πF.SetLineno(1694)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("term", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1694: class term(Part, TextElement): pass
					πF.SetLineno(1694)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("term").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßterm.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1695: class classifier(Part, TextElement): pass
			πF.SetLineno(1695)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("classifier", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1695: class classifier(Part, TextElement): pass
					πF.SetLineno(1695)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("classifier").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßclassifier.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1696: class definition(Part, Element): pass
			πF.SetLineno(1696)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("definition", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1696: class definition(Part, Element): pass
					πF.SetLineno(1696)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("definition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefinition.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1697: class field_list(Sequential, Element): pass
			πF.SetLineno(1697)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSequential); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("field_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1697: class field_list(Sequential, Element): pass
					πF.SetLineno(1697)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("field_list").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfield_list.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1698: class field(Part, Element): pass
			πF.SetLineno(1698)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("field", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1698: class field(Part, Element): pass
					πF.SetLineno(1698)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("field").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfield.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1699: class field_name(Part, TextElement): pass
			πF.SetLineno(1699)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("field_name", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1699: class field_name(Part, TextElement): pass
					πF.SetLineno(1699)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("field_name").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfield_name.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1700: class field_body(Part, Element): pass
			πF.SetLineno(1700)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("field_body", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1700: class field_body(Part, Element): pass
					πF.SetLineno(1700)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("field_body").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfield_body.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1703: class option(Part, Element):
			πF.SetLineno(1703)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("option", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1705: child_text_separator = ''
					πF.SetLineno(1705)
					// line 1705: child_text_separator = ''
					πF.SetLineno(1705)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßchild_text_separator.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("option").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoption.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1708: class option_argument(Part, TextElement):
			πF.SetLineno(1708)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("option_argument", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1710: def astext(self):
					πF.SetLineno(1710)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1711: return self.get('delimiter', ' ') + TextElement.astext(self)
							πF.SetLineno(1711)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßdelimiter.ToObject()
							πTemp002[1] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("option_argument").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoption_argument.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1714: class option_group(Part, Element):
			πF.SetLineno(1714)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("option_group", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1716: child_text_separator = ', '
					πF.SetLineno(1716)
					// line 1716: child_text_separator = ', '
					πF.SetLineno(1716)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr(", ").ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßchild_text_separator.ToObject(), πg.NewStr(", ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("option_group").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoption_group.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1719: class option_list(Sequential, Element): pass
			πF.SetLineno(1719)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSequential); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("option_list", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1719: class option_list(Sequential, Element): pass
					πF.SetLineno(1719)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("option_list").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoption_list.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1722: class option_list_item(Part, Element):
			πF.SetLineno(1722)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("option_list_item", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1724: child_text_separator = '  '
					πF.SetLineno(1724)
					// line 1724: child_text_separator = '  '
					πF.SetLineno(1724)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("  ").ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßchild_text_separator.ToObject(), πg.NewStr("  ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("option_list_item").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoption_list_item.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1727: class option_string(Part, TextElement): pass
			πF.SetLineno(1727)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("option_string", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1727: class option_string(Part, TextElement): pass
					πF.SetLineno(1727)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("option_string").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoption_string.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1728: class description(Part, Element): pass
			πF.SetLineno(1728)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("description", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1728: class description(Part, Element): pass
					πF.SetLineno(1728)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("description").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdescription.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1729: class literal_block(General, FixedTextElement): pass
			πF.SetLineno(1729)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFixedTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("literal_block", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1729: class literal_block(General, FixedTextElement): pass
					πF.SetLineno(1729)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("literal_block").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßliteral_block.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1730: class doctest_block(General, FixedTextElement): pass
			πF.SetLineno(1730)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFixedTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("doctest_block", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1730: class doctest_block(General, FixedTextElement): pass
					πF.SetLineno(1730)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("doctest_block").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdoctest_block.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1731: class math_block(General, FixedTextElement): pass
			πF.SetLineno(1731)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFixedTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("math_block", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1731: class math_block(General, FixedTextElement): pass
					πF.SetLineno(1731)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("math_block").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmath_block.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1732: class line_block(General, Element): pass
			πF.SetLineno(1732)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("line_block", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1732: class line_block(General, Element): pass
					πF.SetLineno(1732)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("line_block").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßline_block.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1735: class line(Part, TextElement):
			πF.SetLineno(1735)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("line", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1737: indent = None
					πF.SetLineno(1737)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßindent.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("line").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßline.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1740: class block_quote(General, Element): pass
			πF.SetLineno(1740)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("block_quote", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1740: class block_quote(General, Element): pass
					πF.SetLineno(1740)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("block_quote").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßblock_quote.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1741: class attribution(Part, TextElement): pass
			πF.SetLineno(1741)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("attribution", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1741: class attribution(Part, TextElement): pass
					πF.SetLineno(1741)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("attribution").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßattribution.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1742: class attention(Admonition, Element): pass
			πF.SetLineno(1742)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("attention", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1742: class attention(Admonition, Element): pass
					πF.SetLineno(1742)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("attention").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßattention.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1743: class caution(Admonition, Element): pass
			πF.SetLineno(1743)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("caution", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1743: class caution(Admonition, Element): pass
					πF.SetLineno(1743)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("caution").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcaution.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1744: class danger(Admonition, Element): pass
			πF.SetLineno(1744)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("danger", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1744: class danger(Admonition, Element): pass
					πF.SetLineno(1744)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("danger").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdanger.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1745: class error(Admonition, Element): pass
			πF.SetLineno(1745)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("error", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1745: class error(Admonition, Element): pass
					πF.SetLineno(1745)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1746: class important(Admonition, Element): pass
			πF.SetLineno(1746)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("important", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1746: class important(Admonition, Element): pass
					πF.SetLineno(1746)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("important").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßimportant.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1747: class note(Admonition, Element): pass
			πF.SetLineno(1747)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("note", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1747: class note(Admonition, Element): pass
					πF.SetLineno(1747)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("note").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnote.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1748: class tip(Admonition, Element): pass
			πF.SetLineno(1748)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("tip", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1748: class tip(Admonition, Element): pass
					πF.SetLineno(1748)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("tip").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtip.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1749: class hint(Admonition, Element): pass
			πF.SetLineno(1749)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("hint", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1749: class hint(Admonition, Element): pass
					πF.SetLineno(1749)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("hint").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßhint.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1750: class warning(Admonition, Element): pass
			πF.SetLineno(1750)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("warning", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1750: class warning(Admonition, Element): pass
					πF.SetLineno(1750)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("warning").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwarning.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1751: class admonition(Admonition, Element): pass
			πF.SetLineno(1751)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßAdmonition); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("admonition", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1751: class admonition(Admonition, Element): pass
					πF.SetLineno(1751)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("admonition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßadmonition.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1752: class comment(Special, Invisible, FixedTextElement): pass
			πF.SetLineno(1752)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSpecial); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInvisible); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFixedTextElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("comment", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1752: class comment(Special, Invisible, FixedTextElement): pass
					πF.SetLineno(1752)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("comment").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcomment.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1753: class substitution_definition(Special, Invisible, TextElement): pass
			πF.SetLineno(1753)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSpecial); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInvisible); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("substitution_definition", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1753: class substitution_definition(Special, Invisible, TextElement): pass
					πF.SetLineno(1753)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("substitution_definition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsubstitution_definition.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1754: class target(Special, Invisible, Inline, TextElement, Targetable): pass
			πF.SetLineno(1754)
			πTemp002 = make([]*πg.Object, 5)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSpecial); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInvisible); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTargetable); πE != nil {
				continue
			}
			πTemp002[4] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("target", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1754: class target(Special, Invisible, Inline, TextElement, Targetable): pass
					πF.SetLineno(1754)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("target").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtarget.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1755: class footnote(General, BackLinkable, Element, Labeled, Targetable): pass
			πF.SetLineno(1755)
			πTemp002 = make([]*πg.Object, 5)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBackLinkable); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßLabeled); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTargetable); πE != nil {
				continue
			}
			πTemp002[4] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("footnote", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1755: class footnote(General, BackLinkable, Element, Labeled, Targetable): pass
					πF.SetLineno(1755)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("footnote").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfootnote.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1756: class citation(General, BackLinkable, Element, Labeled, Targetable): pass
			πF.SetLineno(1756)
			πTemp002 = make([]*πg.Object, 5)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBackLinkable); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßLabeled); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTargetable); πE != nil {
				continue
			}
			πTemp002[4] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("citation", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1756: class citation(General, BackLinkable, Element, Labeled, Targetable): pass
					πF.SetLineno(1756)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("citation").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcitation.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1757: class label(Part, TextElement): pass
			πF.SetLineno(1757)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("label", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1757: class label(Part, TextElement): pass
					πF.SetLineno(1757)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("label").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlabel.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1758: class figure(General, Element): pass
			πF.SetLineno(1758)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("figure", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1758: class figure(General, Element): pass
					πF.SetLineno(1758)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("figure").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfigure.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1759: class caption(Part, TextElement): pass
			πF.SetLineno(1759)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("caption", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1759: class caption(Part, TextElement): pass
					πF.SetLineno(1759)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("caption").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcaption.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1760: class legend(Part, Element): pass
			πF.SetLineno(1760)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("legend", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1760: class legend(Part, Element): pass
					πF.SetLineno(1760)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("legend").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlegend.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1761: class table(General, Element): pass
			πF.SetLineno(1761)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("table", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1761: class table(General, Element): pass
					πF.SetLineno(1761)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("table").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtable.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1762: class tgroup(Part, Element): pass
			πF.SetLineno(1762)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("tgroup", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1762: class tgroup(Part, Element): pass
					πF.SetLineno(1762)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("tgroup").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtgroup.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1763: class colspec(Part, Element): pass
			πF.SetLineno(1763)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("colspec", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1763: class colspec(Part, Element): pass
					πF.SetLineno(1763)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("colspec").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcolspec.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1764: class thead(Part, Element): pass
			πF.SetLineno(1764)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("thead", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1764: class thead(Part, Element): pass
					πF.SetLineno(1764)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("thead").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßthead.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1765: class tbody(Part, Element): pass
			πF.SetLineno(1765)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("tbody", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1765: class tbody(Part, Element): pass
					πF.SetLineno(1765)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("tbody").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtbody.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1766: class row(Part, Element): pass
			πF.SetLineno(1766)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("row", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1766: class row(Part, Element): pass
					πF.SetLineno(1766)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("row").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßrow.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1767: class entry(Part, Element): pass
			πF.SetLineno(1767)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPart); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("entry", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1767: class entry(Part, Element): pass
					πF.SetLineno(1767)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("entry").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßentry.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1770: class system_message(Special, BackLinkable, PreBibliographic, Element):
			πF.SetLineno(1770)
			πTemp002 = make([]*πg.Object, 4)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSpecial); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBackLinkable); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPreBibliographic); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("system_message", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 1772: """
					πF.SetLineno(1772)
					// line 1772: """
					πF.SetLineno(1772)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    System message element.\n\n    Do not instantiate this class directly; use\n    ``document.reporter.info/warning/error/severe()`` instead.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1779: def __init__(self, message=None, *children, **attributes):
					πF.SetLineno(1779)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "message", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
						var µchildren *πg.Object = πArgs[2]
						_ = µchildren
						var µattributes *πg.Object = πArgs[3]
						_ = µattributes
						var µrawsource *πg.Object = πg.UnboundLocal
						_ = µrawsource
						var µp *πg.Object = πg.UnboundLocal
						_ = µp
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
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
							// line 1780: rawsource = attributes.get('rawsource', '')
							πF.SetLineno(1780)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßrawsource.ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattributes, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrawsource = πTemp003
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmessage); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1781: if message:
							πF.SetLineno(1781)
						Label1:
							// line 1782: p = paragraph('', message)
							πF.SetLineno(1782)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πTemp002, πE = πg.ResolveGlobal(πF, ßparagraph); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp003
							// line 1783: children = (p,) + children
							πF.SetLineno(1783)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(µp).ToObject()
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µchildren); πE != nil {
								continue
							}
							µchildren = πTemp002
							goto Label2
						Label2:
							// line 1784: try:
							πF.SetLineno(1784)
							πF.PushCheckpoint(4)
							// line 1785: Element.__init__(self, rawsource, *children, **attributes)
							πF.SetLineno(1785)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							πTemp001[1] = µrawsource
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µchildren, nil, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							goto Label5
							// line 1786: except:
							πF.SetLineno(1786)
						Label5:
							// line 1787: print('system_message: children=%r' % (children,))
							πF.SetLineno(1787)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(µchildren).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("system_message: children=%r").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1788: raise
							πF.SetLineno(1788)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1790: def astext(self):
					πF.SetLineno(1790)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 1791: line = self.get('line', '')
							πF.SetLineno(1791)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßline.ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µline = πTemp003
							// line 1792: return u'%s:%s: (%s/%s) %s' % (self['source'], line, self['type'],
							πF.SetLineno(1792)
							πTemp004 = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004 = ßtype.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							πTemp004 = ßlevel.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple5(πTemp005, µline, πTemp006, πTemp007, πTemp004).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("%s:%s: (%s/%s) %s").ToObject(), πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("system_message").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsystem_message.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1796: class pending(Special, Invisible, Element):
			πF.SetLineno(1796)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSpecial); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInvisible); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("pending", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1798: """
					πF.SetLineno(1798)
					// line 1798: """
					πF.SetLineno(1798)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    The \"pending\" element is used to encapsulate a pending operation: the\n    operation (transform), the point at which to apply it, and any data it\n    requires.  Only the pending operation's location within the document is\n    stored in the public document tree (by the \"pending\" object itself); the\n    operation and its data are stored in the \"pending\" object's internal\n    instance attributes.\n\n    For example, say you want a table of contents in your reStructuredText\n    document.  The easiest way to specify where to put it is from within the\n    document, with a directive::\n\n        .. contents::\n\n    But the \"contents\" directive can't do its work until the entire document\n    has been parsed and possibly transformed to some extent.  So the directive\n    code leaves a placeholder behind that will trigger the second phase of its\n    processing, something like this::\n\n        <pending ...public attributes...> + internal attributes\n\n    Use `document.note_pending()` so that the\n    `docutils.transforms.Transformer` stage of processing can run all pending\n    transforms.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1824: def __init__(self, transform, details=None,
					πF.SetLineno(1824)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "transform", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "details", Def: πTemp003}
					πTemp002[3] = πg.Param{Name: "rawsource", Def: ß.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtransform *πg.Object = πArgs[1]
						_ = µtransform
						var µdetails *πg.Object = πArgs[2]
						_ = µdetails
						var µrawsource *πg.Object = πArgs[3]
						_ = µrawsource
						var µchildren *πg.Object = πArgs[4]
						_ = µchildren
						var µattributes *πg.Object = πArgs[5]
						_ = µattributes
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Dict
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
							// line 1826: Element.__init__(self, rawsource, *children, **attributes)
							πF.SetLineno(1826)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µrawsource, "rawsource"); πE != nil {
								continue
							}
							πTemp001[1] = µrawsource
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µchildren, nil, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1828: self.transform = transform
							πF.SetLineno(1828)
							if πE = πg.CheckLocal(πF, µtransform, "transform"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µtransform); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtransform, πTemp002); πE != nil {
								continue
							}
							// line 1829: """The `docutils.transforms.Transform` class implementing the pending
							πF.SetLineno(1829)
							// line 1832: self.details = details or {}
							πF.SetLineno(1832)
							if πE = πg.CheckLocal(πF, µdetails, "details"); πE != nil {
								continue
							}
							πTemp002 = µdetails
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							πTemp005 = πg.NewDict()
							πTemp003 = πTemp005.ToObject()
							πTemp002 = πTemp003
						Label1:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdetails, πTemp003); πE != nil {
								continue
							}
							// line 1833: """Detail data (dictionary) required by the pending operation."""
							πF.SetLineno(1833)
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
					// line 1835: def pformat(self, indent='    ', level=0):
					πF.SetLineno(1835)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent", Def: πg.NewStr("    ").ToObject()}
					πTemp002[2] = πg.Param{Name: "level", Def: πg.NewInt(0).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("pformat", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindent *πg.Object = πArgs[1]
						_ = µindent
						var µlevel *πg.Object = πArgs[2]
						_ = µlevel
						var µinternals *πg.Object = πg.UnboundLocal
						_ = µinternals
						var µdetails *πg.Object = πg.UnboundLocal
						_ = µdetails
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µv *πg.Object = πg.UnboundLocal
						_ = µv
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
						var πTemp009 []πg.Param
						_ = πTemp009
						var πTemp010 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 10:
								goto Label10
							case 9:
								goto Label9
							default:
								panic("unexpected function state")
							}
							// line 1836: internals = [
							πF.SetLineno(1836)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = πg.NewStr(".. internal attributes:").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtransform, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__module__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtransform, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("     .transform: %s.%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("     .details:").ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µinternals = πTemp002
							// line 1841: details = sorted(self.details.items())
							πF.SetLineno(1841)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdetails = πTemp003
							if πE = πg.CheckLocal(πF, µdetails, "details"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µdetails); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µvalue = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003 = µvalue
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label5
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp005
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label5
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µvalue, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp005
						Label5:
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 1843: if isinstance(value, Node):
							πF.SetLineno(1843)
						Label4:
							// line 1844: internals.append('%7s%s:' % ('', key))
							πF.SetLineno(1844)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(ß.ToObject(), µkey).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%7s%s:").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µinternals, "internals"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µinternals, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1845: internals.extend(['%9s%s' % ('', line)
							πF.SetLineno(1845)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
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
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µvalue, ßpformat, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
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
											µline = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1845: internals.extend(['%9s%s' % ('', line)
										πF.SetLineno(1845)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(ß.ToObject(), µline).ToObject()
										if πTemp002, πE = πg.Mod(πF, πg.NewStr("%9s%s").ToObject(), πTemp003); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µinternals, "internals"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µinternals, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
							// line 1847: elif value and isinstance(value, list) \
							πF.SetLineno(1847)
						Label6:
							// line 1849: internals.append('%7s%s:' % ('', key))
							πF.SetLineno(1849)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(ß.ToObject(), µkey).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%7s%s:").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µinternals, "internals"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µinternals, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µvalue); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp008 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label11
							}
							if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µv = πTemp005
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(9)
							// line 1851: internals.extend(['%9s%s' % ('', line)
							πF.SetLineno(1851)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
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
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µv, ßpformat, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
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
											µline = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1851: internals.extend(['%9s%s' % ('', line)
										πF.SetLineno(1851)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(ß.ToObject(), µline).ToObject()
										if πTemp002, πE = πg.Mod(πF, πg.NewStr("%9s%s").ToObject(), πTemp003); πE != nil {
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
							if πTemp011, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µinternals, "internals"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µinternals, ßextend, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							goto Label8
						Label7:
							// line 1854: internals.append('%7s%s: %r' % ('', key, value))
							πF.SetLineno(1854)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple3(ß.ToObject(), µkey, µvalue).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%7s%s: %r").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µinternals, "internals"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µinternals, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1855: return (Element.pformat(self, indent, level)
							πF.SetLineno(1855)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp001[1] = µindent
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[2] = µlevel
							if πTemp003, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp011 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
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
										if πE = πg.CheckLocal(πF, µinternals, "internals"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µinternals); πE != nil {
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
											µline = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1856: + ''.join([('    %s%s\n' % (indent * level, line))
										πF.SetLineno(1856)
										if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
											continue
										}
										if πTemp006, πE = πg.Mul(πF, µindent, µlevel); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp005 = πg.NewTuple2(πTemp006, µline).ToObject()
										if πTemp004, πE = πg.Mod(πF, πg.NewStr("    %s%s\n").ToObject(), πTemp005); πE != nil {
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
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp012}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp012); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpformat.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1859: def copy(self):
					πF.SetLineno(1859)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πg.UnboundLocal
						_ = µobj
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
							// line 1860: obj = self.__class__(self.transform, self.details, self.rawsource,
							πF.SetLineno(1860)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransform, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdetails, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßattributes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp001, nil, nil, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobj = πTemp004
							// line 1862: obj.document = self.document
							πF.SetLineno(1862)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßdocument, πTemp003); πE != nil {
								continue
							}
							// line 1863: obj.source = self.source
							πF.SetLineno(1863)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßsource, πTemp003); πE != nil {
								continue
							}
							// line 1864: obj.line = self.line
							πF.SetLineno(1864)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µobj, ßline, πTemp003); πE != nil {
								continue
							}
							// line 1865: return obj
							πF.SetLineno(1865)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("pending").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpending.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1868: class raw(Special, Inline, PreBibliographic, FixedTextElement):
			πF.SetLineno(1868)
			πTemp002 = make([]*πg.Object, 4)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSpecial); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßPreBibliographic); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFixedTextElement); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("raw", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1870: """
					πF.SetLineno(1870)
					// line 1870: """
					πF.SetLineno(1870)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raw data that is to be passed untouched to the Writer.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1874: pass
					πF.SetLineno(1874)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("raw").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßraw.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1881: class emphasis(Inline, TextElement): pass
			πF.SetLineno(1881)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("emphasis", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1881: class emphasis(Inline, TextElement): pass
					πF.SetLineno(1881)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("emphasis").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßemphasis.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1882: class strong(Inline, TextElement): pass
			πF.SetLineno(1882)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("strong", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1882: class strong(Inline, TextElement): pass
					πF.SetLineno(1882)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("strong").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstrong.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1883: class literal(Inline, TextElement): pass
			πF.SetLineno(1883)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("literal", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1883: class literal(Inline, TextElement): pass
					πF.SetLineno(1883)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("literal").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßliteral.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1884: class reference(General, Inline, Referential, TextElement): pass
			πF.SetLineno(1884)
			πTemp002 = make([]*πg.Object, 4)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßReferential); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[3] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("reference", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1884: class reference(General, Inline, Referential, TextElement): pass
					πF.SetLineno(1884)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("reference").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreference.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1885: class footnote_reference(Inline, Referential, TextElement): pass
			πF.SetLineno(1885)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßReferential); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("footnote_reference", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1885: class footnote_reference(Inline, Referential, TextElement): pass
					πF.SetLineno(1885)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("footnote_reference").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfootnote_reference.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1886: class citation_reference(Inline, Referential, TextElement): pass
			πF.SetLineno(1886)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßReferential); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("citation_reference", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1886: class citation_reference(Inline, Referential, TextElement): pass
					πF.SetLineno(1886)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("citation_reference").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcitation_reference.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1887: class substitution_reference(Inline, TextElement): pass
			πF.SetLineno(1887)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("substitution_reference", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1887: class substitution_reference(Inline, TextElement): pass
					πF.SetLineno(1887)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("substitution_reference").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsubstitution_reference.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1888: class title_reference(Inline, TextElement): pass
			πF.SetLineno(1888)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("title_reference", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1888: class title_reference(Inline, TextElement): pass
					πF.SetLineno(1888)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("title_reference").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßtitle_reference.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1889: class abbreviation(Inline, TextElement): pass
			πF.SetLineno(1889)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("abbreviation", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1889: class abbreviation(Inline, TextElement): pass
					πF.SetLineno(1889)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("abbreviation").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßabbreviation.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1890: class acronym(Inline, TextElement): pass
			πF.SetLineno(1890)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("acronym", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1890: class acronym(Inline, TextElement): pass
					πF.SetLineno(1890)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("acronym").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßacronym.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1891: class superscript(Inline, TextElement): pass
			πF.SetLineno(1891)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("superscript", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1891: class superscript(Inline, TextElement): pass
					πF.SetLineno(1891)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("superscript").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsuperscript.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1892: class subscript(Inline, TextElement): pass
			πF.SetLineno(1892)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("subscript", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1892: class subscript(Inline, TextElement): pass
					πF.SetLineno(1892)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("subscript").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsubscript.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1893: class math(Inline, TextElement): pass
			πF.SetLineno(1893)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("math", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1893: class math(Inline, TextElement): pass
					πF.SetLineno(1893)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("math").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmath.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1896: class image(General, Inline, Element):
			πF.SetLineno(1896)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßGeneral); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßElement); πE != nil {
				continue
			}
			πTemp002[2] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("image", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1898: def astext(self):
					πF.SetLineno(1898)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("astext", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1899: return self.get('alt', '')
							πF.SetLineno(1899)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßalt.ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßastext.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("image").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßimage.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1902: class inline(Inline, TextElement): pass
			πF.SetLineno(1902)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("inline", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1902: class inline(Inline, TextElement): pass
					πF.SetLineno(1902)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("inline").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßinline.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1903: class problematic(Inline, TextElement): pass
			πF.SetLineno(1903)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("problematic", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1903: class problematic(Inline, TextElement): pass
					πF.SetLineno(1903)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("problematic").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßproblematic.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1904: class generated(Inline, TextElement): pass
			πF.SetLineno(1904)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßInline); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTextElement); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("generated", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1904: class generated(Inline, TextElement): pass
					πF.SetLineno(1904)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("generated").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgenerated.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1911: node_class_names = """
			πF.SetLineno(1911)
			if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("\n    Text\n    abbreviation acronym address admonition attention attribution author\n        authors\n    block_quote bullet_list\n    caption caution citation citation_reference classifier colspec comment\n        compound contact container copyright\n    danger date decoration definition definition_list definition_list_item\n        description docinfo doctest_block document\n    emphasis entry enumerated_list error\n    field field_body field_list field_name figure footer\n        footnote footnote_reference\n    generated\n    header hint\n    image important inline\n    label legend line line_block list_item literal literal_block\n    math math_block\n    note\n    option option_argument option_group option_list option_list_item\n        option_string organization\n    paragraph pending problematic\n    raw reference revision row rubric\n    section sidebar status strong subscript substitution_definition\n        substitution_reference subtitle superscript system_message\n    table target tbody term tgroup thead tip title title_reference topic\n        transition\n    version\n    warning").ToObject(), ßsplit, nil); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnode_class_names.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 1939: """A list of names of all concrete Node subclasses."""
			πF.SetLineno(1939)
			// line 1942: class NodeVisitor(object):
			πF.SetLineno(1942)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NodeVisitor", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 1944: """
					πF.SetLineno(1944)
					// line 1944: """
					πF.SetLineno(1944)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    \"Visitor\" pattern [GoF95]_ abstract superclass implementation for\n    document tree traversals.\n\n    Each node class has corresponding methods, doing nothing by\n    default; override individual methods for specific and useful\n    behaviour.  The `dispatch_visit()` method is called by\n    `Node.walk()` upon entering a node.  `Node.walkabout()` also calls\n    the `dispatch_departure()` method before exiting a node.\n\n    The dispatch methods call \"``visit_`` + node class name\" or\n    \"``depart_`` + node class name\", resp.\n\n    This is a base class for visitors whose ``visit_...`` & ``depart_...``\n    methods should be implemented for *all* node types encountered (such as\n    for `docutils.writers.Writer` subclasses).  Unimplemented methods will\n    raise exceptions.\n\n    For sparse traversals, where only certain node types are of interest,\n    subclass `SparseNodeVisitor` instead.  When (mostly or entirely) uniform\n    processing is desired, subclass `GenericNodeVisitor`.\n\n    .. [GoF95] Gamma, Helm, Johnson, Vlissides. *Design Patterns: Elements of\n       Reusable Object-Oriented Software*. Addison-Wesley, Reading, MA, USA,\n       1995.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1971: optional = ()
					πF.SetLineno(1971)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßoptional.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1972: """
					πF.SetLineno(1972)
					// line 1981: def __init__(self, document):
					πF.SetLineno(1981)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
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
							// line 1982: self.document = document
							πF.SetLineno(1982)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdocument); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp001); πE != nil {
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
					// line 1984: def dispatch_visit(self, node):
					πF.SetLineno(1984)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("dispatch_visit", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µnode_name *πg.Object = πg.UnboundLocal
						_ = µnode_name
						var µmethod *πg.Object = πg.UnboundLocal
						_ = µmethod
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
							// line 1985: """
							πF.SetLineno(1985)
							// line 1990: node_name = node.__class__.__name__
							πF.SetLineno(1990)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
								continue
							}
							µnode_name = πTemp002
							// line 1991: method = getattr(self, 'visit_' + node_name, self.unknown_visit)
							πF.SetLineno(1991)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µnode_name, "node_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, ßvisit_.ToObject(), µnode_name); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßunknown_visit, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmethod = πTemp002
							// line 1992: self.document.reporter.debug(
							πF.SetLineno(1992)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µmethod, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode_name, "node_name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µnode_name).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("docutils.nodes.NodeVisitor.dispatch_visit calling %s for %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1995: return method(node)
							πF.SetLineno(1995)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp001, πE = µmethod.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßdispatch_visit.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1985: """
					πF.SetLineno(1985)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Call self.\"``visit_`` + node class name\" with `node` as\n        parameter.  If the ``visit_...`` method does not exist, call\n        self.unknown_visit.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch_visit); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1997: def dispatch_departure(self, node):
					πF.SetLineno(1997)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("dispatch_departure", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µnode_name *πg.Object = πg.UnboundLocal
						_ = µnode_name
						var µmethod *πg.Object = πg.UnboundLocal
						_ = µmethod
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
							// line 1998: """
							πF.SetLineno(1998)
							// line 2003: node_name = node.__class__.__name__
							πF.SetLineno(2003)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
								continue
							}
							µnode_name = πTemp002
							// line 2004: method = getattr(self, 'depart_' + node_name, self.unknown_departure)
							πF.SetLineno(2004)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µnode_name, "node_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, ßdepart_.ToObject(), µnode_name); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßunknown_departure, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmethod = πTemp002
							// line 2005: self.document.reporter.debug(
							πF.SetLineno(2005)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µmethod, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode_name, "node_name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µnode_name).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("docutils.nodes.NodeVisitor.dispatch_departure calling %s for %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 2008: return method(node)
							πF.SetLineno(2008)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp001, πE = µmethod.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßdispatch_departure.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1998: """
					πF.SetLineno(1998)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Call self.\"``depart_`` + node class name\" with `node` as\n        parameter.  If the ``depart_...`` method does not exist, call\n        self.unknown_departure.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch_departure); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 2010: def unknown_visit(self, node):
					πF.SetLineno(2010)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("unknown_visit", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 2011: """
							πF.SetLineno(2011)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstrict_visitor, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptional, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 2016: if  (self.document.settings.strict_visitor
							πF.SetLineno(2016)
						Label2:
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s visiting unknown node type: %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 2018: raise NotImplementedError(
							πF.SetLineno(2018)
							πE = πF.Raise(πTemp003, nil, nil)
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
					if πE = πClass.SetItem(πF, ßunknown_visit.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 2011: """
					πF.SetLineno(2011)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Called when entering unknown `Node` types.\n\n        Raise an exception unless overridden.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunknown_visit); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 2022: def unknown_departure(self, node):
					πF.SetLineno(2022)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("unknown_departure", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 2023: """
							πF.SetLineno(2023)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßstrict_visitor, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptional, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 2028: if  (self.document.settings.strict_visitor
							πF.SetLineno(2028)
						Label2:
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s departing unknown node type: %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 2030: raise NotImplementedError(
							πF.SetLineno(2030)
							πE = πF.Raise(πTemp003, nil, nil)
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
					if πE = πClass.SetItem(πF, ßunknown_departure.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 2023: """
					πF.SetLineno(2023)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Called before exiting unknown `Node` types.\n\n        Raise exception unless overridden.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunknown_departure); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("NodeVisitor").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNodeVisitor.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 2035: class SparseNodeVisitor(NodeVisitor):
			πF.SetLineno(2035)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNodeVisitor); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SparseNodeVisitor", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2037: """
					πF.SetLineno(2037)
					// line 2037: """
					πF.SetLineno(2037)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Base class for sparse traversals, where only certain node types are of\n    interest.  When ``visit_...`` & ``depart_...`` methods should be\n    implemented for *all* node types (such as for `docutils.writers.Writer`\n    subclasses), subclass `NodeVisitor` instead.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("SparseNodeVisitor").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSparseNodeVisitor.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 2045: class GenericNodeVisitor(NodeVisitor):
			πF.SetLineno(2045)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNodeVisitor); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GenericNodeVisitor", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2047: """
					πF.SetLineno(2047)
					// line 2047: """
					πF.SetLineno(2047)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Generic \"Visitor\" abstract superclass, for simple traversals.\n\n    Unless overridden, each ``visit_...`` method calls `default_visit()`, and\n    each ``depart_...`` method (when using `Node.walkabout()`) calls\n    `default_departure()`. `default_visit()` (and `default_departure()`) must\n    be overridden in subclasses.\n\n    Define fully generic visitors by overriding `default_visit()` (and\n    `default_departure()`) only. Define semi-generic visitors by overriding\n    individual ``visit_...()`` (and ``depart_...()``) methods also.\n\n    `NodeVisitor.unknown_visit()` (`NodeVisitor.unknown_departure()`) should\n    be overridden for default behavior.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2063: def default_visit(self, node):
					πF.SetLineno(2063)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("default_visit", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 2064: """Override for generic, uniform traversals."""
							πF.SetLineno(2064)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 2065: raise NotImplementedError
							πF.SetLineno(2065)
							πE = πF.Raise(πTemp001, nil, nil)
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
					// line 2064: """Override for generic, uniform traversals."""
					πF.SetLineno(2064)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Override for generic, uniform traversals.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_visit); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 2067: def default_departure(self, node):
					πF.SetLineno(2067)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("default_departure", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 2068: """Override for generic, uniform traversals."""
							πF.SetLineno(2068)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 2069: raise NotImplementedError
							πF.SetLineno(2069)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdefault_departure.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 2068: """Override for generic, uniform traversals."""
					πF.SetLineno(2068)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Override for generic, uniform traversals.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_departure); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("GenericNodeVisitor").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGenericNodeVisitor.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 2071: def _call_default_visit(self, node):
			πF.SetLineno(2071)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp007[1] = πg.Param{Name: "node", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("_call_default_visit", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 2072: self.default_visit(node)
					πF.SetLineno(2072)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001[0] = µnode
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ßdefault_visit, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_call_default_visit.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 2074: def _call_default_departure(self, node):
			πF.SetLineno(2074)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp007[1] = πg.Param{Name: "node", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_call_default_departure", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 2075: self.default_departure(node)
					πF.SetLineno(2075)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001[0] = µnode
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ßdefault_departure, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_call_default_departure.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 2077: def _nop(self, node):
			πF.SetLineno(2077)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp007[1] = πg.Param{Name: "node", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_nop", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 2078: pass
					πF.SetLineno(2078)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_nop.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 2080: def _add_node_class_names(names):
			πF.SetLineno(2080)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "names", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_add_node_class_names", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnames *πg.Object = πArgs[0]
				_ = µnames
				var µ_name *πg.Object = πg.UnboundLocal
				_ = µ_name
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
					// line 2081: """Save typing with dynamic assignments:"""
					πF.SetLineno(2081)
					if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
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
						µ_name = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 2083: setattr(GenericNodeVisitor, "visit_" + _name, _call_default_visit)
					πF.SetLineno(2083)
					πTemp005 = πF.MakeArgs(3)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßGenericNodeVisitor); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µ_name, "_name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, ßvisit_.ToObject(), µ_name); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_call_default_visit); πE != nil {
						continue
					}
					πTemp005[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 2084: setattr(GenericNodeVisitor, "depart_" + _name, _call_default_departure)
					πF.SetLineno(2084)
					πTemp005 = πF.MakeArgs(3)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßGenericNodeVisitor); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µ_name, "_name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, ßdepart_.ToObject(), µ_name); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_call_default_departure); πE != nil {
						continue
					}
					πTemp005[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 2085: setattr(SparseNodeVisitor, 'visit_' + _name, _nop)
					πF.SetLineno(2085)
					πTemp005 = πF.MakeArgs(3)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSparseNodeVisitor); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µ_name, "_name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, ßvisit_.ToObject(), µ_name); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_nop); πE != nil {
						continue
					}
					πTemp005[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 2086: setattr(SparseNodeVisitor, 'depart_' + _name, _nop)
					πF.SetLineno(2086)
					πTemp005 = πF.MakeArgs(3)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSparseNodeVisitor); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µ_name, "_name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, ßdepart_.ToObject(), µ_name); πE != nil {
						continue
					}
					πTemp005[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_nop); πE != nil {
						continue
					}
					πTemp005[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ß_add_node_class_names.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 2081: """Save typing with dynamic assignments:"""
			πF.SetLineno(2081)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Save typing with dynamic assignments:").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_add_node_class_names); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 2088: _add_node_class_names(node_class_names)
			πF.SetLineno(2088)
			πTemp002 = πF.MakeArgs(1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßnode_class_names); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			if πTemp011, πE = πg.ResolveGlobal(πF, ß_add_node_class_names); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 2091: class TreeCopyVisitor(GenericNodeVisitor):
			πF.SetLineno(2091)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßGenericNodeVisitor); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TreeCopyVisitor", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp007 *πg.Object
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2093: """
					πF.SetLineno(2093)
					// line 2093: """
					πF.SetLineno(2093)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Make a complete copy of a tree or branch, including element attributes.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2097: def __init__(self, document):
					πF.SetLineno(2097)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
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
							// line 2098: GenericNodeVisitor.__init__(self, document)
							πF.SetLineno(2098)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πTemp002, πE = πg.ResolveGlobal(πF, ßGenericNodeVisitor); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 2099: self.parent_stack = []
							πF.SetLineno(2099)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent_stack, πTemp003); πE != nil {
								continue
							}
							// line 2100: self.parent = []
							πF.SetLineno(2100)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp003); πE != nil {
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
					// line 2102: def get_tree_copy(self):
					πF.SetLineno(2102)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("get_tree_copy", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 2103: return self.parent[0]
							πF.SetLineno(2103)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_tree_copy.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 2105: def default_visit(self, node):
					πF.SetLineno(2105)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("default_visit", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µnewnode *πg.Object = πg.UnboundLocal
						_ = µnewnode
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
							// line 2106: """Copy the current node, and make it the new acting parent."""
							πF.SetLineno(2106)
							// line 2107: newnode = node.copy()
							πF.SetLineno(2107)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnewnode = πTemp002
							// line 2108: self.parent.append(newnode)
							πF.SetLineno(2108)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnewnode, "newnode"); πE != nil {
								continue
							}
							πTemp003[0] = µnewnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 2109: self.parent_stack.append(self.parent)
							πF.SetLineno(2109)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent_stack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 2110: self.parent = newnode
							πF.SetLineno(2110)
							if πE = πg.CheckLocal(πF, µnewnode, "newnode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewnode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdefault_visit.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 2106: """Copy the current node, and make it the new acting parent."""
					πF.SetLineno(2106)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Copy the current node, and make it the new acting parent.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_visit); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 2112: def default_departure(self, node):
					πF.SetLineno(2112)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("default_departure", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 2113: """Restore the previous acting parent."""
							πF.SetLineno(2113)
							// line 2114: self.parent = self.parent_stack.pop()
							πF.SetLineno(2114)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent_stack, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdefault_departure.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 2113: """Restore the previous acting parent."""
					πF.SetLineno(2113)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Restore the previous acting parent.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_departure); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("TreeCopyVisitor").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTreeCopyVisitor.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2117: class TreePruningException(Exception):
			πF.SetLineno(2117)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TreePruningException", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2119: """
					πF.SetLineno(2119)
					// line 2119: """
					πF.SetLineno(2119)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Base class for `NodeVisitor`-related tree pruning exceptions.\n\n    Raise subclasses from within ``visit_...`` or ``depart_...`` methods\n    called from `Node.walk()` and `Node.walkabout()` tree traversals to prune\n    the tree traversed.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2127: pass
					πF.SetLineno(2127)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("TreePruningException").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTreePruningException.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2130: class SkipChildren(TreePruningException):
			πF.SetLineno(2130)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTreePruningException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SkipChildren", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2132: """
					πF.SetLineno(2132)
					// line 2132: """
					πF.SetLineno(2132)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Do not visit any children of the current node.  The current node's\n    siblings and ``depart_...`` method are not affected.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2137: pass
					πF.SetLineno(2137)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("SkipChildren").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSkipChildren.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2140: class SkipSiblings(TreePruningException):
			πF.SetLineno(2140)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTreePruningException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SkipSiblings", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2142: """
					πF.SetLineno(2142)
					// line 2142: """
					πF.SetLineno(2142)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Do not visit any more siblings (to the right) of the current node.  The\n    current node's children and its ``depart_...`` method are not affected.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2147: pass
					πF.SetLineno(2147)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("SkipSiblings").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSkipSiblings.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2150: class SkipNode(TreePruningException):
			πF.SetLineno(2150)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTreePruningException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SkipNode", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2152: """
					πF.SetLineno(2152)
					// line 2152: """
					πF.SetLineno(2152)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Do not visit the current node's children, and do not call the current\n    node's ``depart_...`` method.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2157: pass
					πF.SetLineno(2157)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("SkipNode").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSkipNode.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2160: class SkipDeparture(TreePruningException):
			πF.SetLineno(2160)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTreePruningException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SkipDeparture", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2162: """
					πF.SetLineno(2162)
					// line 2162: """
					πF.SetLineno(2162)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Do not call the current node's ``depart_...`` method.  The current node's\n    children and siblings are not affected.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2167: pass
					πF.SetLineno(2167)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("SkipDeparture").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSkipDeparture.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2170: class NodeFound(TreePruningException):
			πF.SetLineno(2170)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTreePruningException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NodeFound", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2172: """
					πF.SetLineno(2172)
					// line 2172: """
					πF.SetLineno(2172)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise to indicate that the target of a search has been found.  This\n    exception must be caught by the client; it is not caught by the traversal\n    code.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2178: pass
					πF.SetLineno(2178)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("NodeFound").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNodeFound.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2181: class StopTraversal(TreePruningException):
			πF.SetLineno(2181)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßTreePruningException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp006 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StopTraversal", "/usr/lib/python2.7/site-packages/docutils/nodes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 2183: """
					πF.SetLineno(2183)
					// line 2183: """
					πF.SetLineno(2183)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Stop the traversal alltogether.  The current node's ``depart_...`` method\n    is not affected.  The parent nodes ``depart_...`` methods are also called\n    as usual.  No other nodes are visited.  This is an alternative to\n    NodeFound that does not cause exception handling to trickle up to the\n    caller.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 2191: pass
					πF.SetLineno(2191)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("StopTraversal").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStopTraversal.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2194: def make_id(string):
			πF.SetLineno(2194)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "string", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("make_id", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstring *πg.Object = πArgs[0]
				_ = µstring
				var µid *πg.Object = πg.UnboundLocal
				_ = µid
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
					// line 2195: """
					πF.SetLineno(2195)
					// line 2229: id = string.lower()
					πF.SetLineno(2229)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µstring, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µid = πTemp002
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					πTemp003[0] = µid
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp003[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
					// line 2230: if not isinstance(id, unicode):
					πF.SetLineno(2230)
				Label1:
					// line 2231: id = id.decode()
					πF.SetLineno(2231)
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µid, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µid = πTemp002
					goto Label2
				Label2:
					// line 2232: id = id.translate(_non_id_translate_digraphs)
					πF.SetLineno(2232)
					πTemp003 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_non_id_translate_digraphs); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µid, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µid = πTemp002
					// line 2233: id = id.translate(_non_id_translate)
					πF.SetLineno(2233)
					πTemp003 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_non_id_translate); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µid, ßtranslate, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µid = πTemp002
					// line 2236: id = unicodedata.normalize('NFKD', id).\
					πF.SetLineno(2236)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßascii.ToObject()
					πTemp006 = πF.MakeArgs(2)
					πTemp006[0] = ßascii.ToObject()
					πTemp006[1] = ßignore.ToObject()
					πTemp007 = πF.MakeArgs(2)
					πTemp007[0] = ßNFKD.ToObject()
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					πTemp007[1] = µid
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicodedata); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnormalize, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßencode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µid = πTemp001
					// line 2239: id = _non_id_chars.sub('-', ' '.join(id.split()))
					πF.SetLineno(2239)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr("-").ToObject()
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µid, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006[0] = πTemp002
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003[1] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_non_id_chars); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsub, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µid = πTemp001
					// line 2240: id = _non_id_at_ends.sub('', id)
					πF.SetLineno(2240)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					πTemp003[1] = µid
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_non_id_at_ends); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsub, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µid = πTemp001
					// line 2241: return str(id)
					πF.SetLineno(2241)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
						continue
					}
					πTemp003[0] = µid
					if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßmake_id.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 2195: """
			πF.SetLineno(2195)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n    Convert `string` into an identifier and return it.\n\n    Docutils identifiers will conform to the regular expression\n    ``[a-z](-?[a-z0-9]+)*``.  For CSS compatibility, identifiers (the \"class\"\n    and \"id\" attributes) should have no underscores, colons, or periods.\n    Hyphens may be used.\n\n    - The `HTML 4.01 spec`_ defines identifiers based on SGML tokens:\n\n          ID and NAME tokens must begin with a letter ([A-Za-z]) and may be\n          followed by any number of letters, digits ([0-9]), hyphens (\"-\"),\n          underscores (\"_\"), colons (\":\"), and periods (\".\").\n\n    - However the `CSS1 spec`_ defines identifiers based on the \"name\" token,\n      a tighter interpretation (\"flex\" tokenizer notation; \"latin1\" and\n      \"escape\" 8-bit characters have been replaced with entities)::\n\n          unicode     \\[0-9a-f]{1,4}\n          latin1      [&iexcl;-&yuml;]\n          escape      {unicode}|\\[ -~&iexcl;-&yuml;]\n          nmchar      [-a-z0-9]|{latin1}|{escape}\n          name        {nmchar}+\n\n    The CSS1 \"nmchar\" rule does not include underscores (\"_\"), colons (\":\"),\n    or periods (\".\"), therefore \"class\" and \"id\" attributes should not contain\n    these characters. They should be replaced with hyphens (\"-\"). Combined\n    with HTML's requirements (the first character must be a letter; no\n    \"unicode\", \"latin1\", or \"escape\" characters), this results in the\n    ``[a-z](-?[a-z0-9]+)*`` pattern.\n\n    .. _HTML 4.01 spec: http://www.w3.org/TR/html401\n    .. _CSS1 spec: http://www.w3.org/TR/REC-CSS1\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßmake_id); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 2243: _non_id_chars = re.compile('[^a-z0-9]+')
			πF.SetLineno(2243)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("[^a-z0-9]+").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_non_id_chars.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 2244: _non_id_at_ends = re.compile('^[-0-9]+|-+$')
			πF.SetLineno(2244)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("^[-0-9]+|-+$").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_non_id_at_ends.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 2245: _non_id_translate = {
			πF.SetLineno(2245)
			πTemp006 = πg.NewDict()
			if πE = πTemp006.SetItem(πF, πg.NewInt(248).ToObject(), πg.NewUnicode("o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(273).ToObject(), πg.NewUnicode("d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(295).ToObject(), πg.NewUnicode("h").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(305).ToObject(), πg.NewUnicode("i").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(322).ToObject(), πg.NewUnicode("l").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(359).ToObject(), πg.NewUnicode("t").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(384).ToObject(), πg.NewUnicode("b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(387).ToObject(), πg.NewUnicode("b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(392).ToObject(), πg.NewUnicode("c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(396).ToObject(), πg.NewUnicode("d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(402).ToObject(), πg.NewUnicode("f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(409).ToObject(), πg.NewUnicode("k").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(410).ToObject(), πg.NewUnicode("l").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(414).ToObject(), πg.NewUnicode("n").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(421).ToObject(), πg.NewUnicode("p").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(427).ToObject(), πg.NewUnicode("t").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(429).ToObject(), πg.NewUnicode("t").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(436).ToObject(), πg.NewUnicode("y").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(438).ToObject(), πg.NewUnicode("z").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(485).ToObject(), πg.NewUnicode("g").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(549).ToObject(), πg.NewUnicode("z").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(564).ToObject(), πg.NewUnicode("l").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(565).ToObject(), πg.NewUnicode("n").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(566).ToObject(), πg.NewUnicode("t").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(567).ToObject(), πg.NewUnicode("j").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(572).ToObject(), πg.NewUnicode("c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(575).ToObject(), πg.NewUnicode("s").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(576).ToObject(), πg.NewUnicode("z").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(583).ToObject(), πg.NewUnicode("e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(585).ToObject(), πg.NewUnicode("j").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(587).ToObject(), πg.NewUnicode("q").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(589).ToObject(), πg.NewUnicode("r").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(591).ToObject(), πg.NewUnicode("y").ToObject()); πE != nil {
				continue
			}
			πTemp012 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_non_id_translate.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 2280: _non_id_translate_digraphs = {
			πF.SetLineno(2280)
			πTemp006 = πg.NewDict()
			if πE = πTemp006.SetItem(πF, πg.NewInt(223).ToObject(), πg.NewUnicode("sz").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(230).ToObject(), πg.NewUnicode("ae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(339).ToObject(), πg.NewUnicode("oe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(568).ToObject(), πg.NewUnicode("db").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewInt(569).ToObject(), πg.NewUnicode("qp").ToObject()); πE != nil {
				continue
			}
			πTemp012 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_non_id_translate_digraphs.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 2288: def dupname(node, name):
			πF.SetLineno(2288)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "node", Def: nil}
			πTemp007[1] = πg.Param{Name: "name", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("dupname", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnode *πg.Object = πArgs[0]
				_ = µnode
				var µname *πg.Object = πArgs[1]
				_ = µname
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
					// line 2289: node['dupnames'].append(name)
					πF.SetLineno(2289)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp002 = ßdupnames.ToObject()
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 2290: node['names'].remove(name)
					πF.SetLineno(2290)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					πTemp002 = ßnames.ToObject()
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 2293: node.referenced = 1
					πF.SetLineno(2293)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µnode, ßreferenced, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdupname.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 2295: def fully_normalize_name(name):
			πF.SetLineno(2295)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "name", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("fully_normalize_name", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
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
					// line 2296: """Return a case- and whitespace-normalized name."""
					πF.SetLineno(2296)
					// line 2297: return ' '.join(name.lower().split())
					πF.SetLineno(2297)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µname, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfully_normalize_name.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 2296: """Return a case- and whitespace-normalized name."""
			πF.SetLineno(2296)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Return a case- and whitespace-normalized name.").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßfully_normalize_name); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 2299: def whitespace_normalize_name(name):
			πF.SetLineno(2299)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "name", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("whitespace_normalize_name", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
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
					// line 2300: """Return a whitespace-normalized name."""
					πF.SetLineno(2300)
					// line 2301: return ' '.join(name.split())
					πF.SetLineno(2301)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µname, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßwhitespace_normalize_name.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 2300: """Return a whitespace-normalized name."""
			πF.SetLineno(2300)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Return a whitespace-normalized name.").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßwhitespace_normalize_name); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 2303: def serial_escape(value):
			πF.SetLineno(2303)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "value", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("serial_escape", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalue *πg.Object = πArgs[0]
				_ = µvalue
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
					// line 2304: """Escape string values that are elements of a list, for serialization."""
					πF.SetLineno(2304)
					// line 2305: return value.replace('\\', r'\\').replace(' ', r'\ ')
					πF.SetLineno(2305)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(" ").ToObject()
					πTemp001[1] = πg.NewStr("\\ ").ToObject()
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("\\").ToObject()
					πTemp002[1] = πg.NewStr("\\\\").ToObject()
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µvalue, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßserial_escape.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 2304: """Escape string values that are elements of a list, for serialization."""
			πF.SetLineno(2304)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Escape string values that are elements of a list, for serialization.").ToObject()); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßserial_escape); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
				continue
			}
			// line 2307: def pseudo_quoteattr(value):
			πF.SetLineno(2307)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "value", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("pseudo_quoteattr", "/usr/lib/python2.7/site-packages/docutils/nodes.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µvalue *πg.Object = πArgs[0]
				_ = µvalue
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
					// line 2308: """Quote attributes for pseudo-xml"""
					πF.SetLineno(2308)
					// line 2309: return '"%s"' % value
					πF.SetLineno(2309)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\"").ToObject(), µvalue); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpseudo_quoteattr.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 2308: """Quote attributes for pseudo-xml"""
			πF.SetLineno(2308)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Quote attributes for pseudo-xml").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßpseudo_quoteattr); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("nodes", Code)
}
