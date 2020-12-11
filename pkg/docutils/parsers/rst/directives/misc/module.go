package misc

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/time"
	_ "github.com/pygolin/stdlib/pkg/urllib"
	_ "github.com/pygolin/stdlib/pkg/urllib/error"
	_ "github.com/pygolin/stdlib/pkg/urllib/request"
	_ "github.com/pygolin/stdlib/pkg/urllib2"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßClass := πg.InternStr("Class")
		ßClassAttribute := πg.InternStr("ClassAttribute")
		ßCodeBlock := πg.InternStr("CodeBlock")
		ßCustomRole := πg.InternStr("CustomRole")
		ßDate := πg.InternStr("Date")
		ßDefaultRole := πg.InternStr("DefaultRole")
		ßDirective := πg.InternStr("Directive")
		ßElement := πg.InternStr("Element")
		ßErrorString := πg.InternStr("ErrorString")
		ßFalse := πg.InternStr("False")
		ßFileInput := πg.InternStr("FileInput")
		ßIOError := πg.InternStr("IOError")
		ßInclude := πg.InternStr("Include")
		ßInliner := πg.InternStr("Inliner")
		ßMarkupError := πg.InternStr("MarkupError")
		ßNone := πg.InternStr("None")
		ßNumberLines := πg.InternStr("NumberLines")
		ßOSError := πg.InternStr("OSError")
		ßRaw := πg.InternStr("Raw")
		ßReplace := πg.InternStr("Replace")
		ßRole := πg.InternStr("Role")
		ßSafeString := πg.InternStr("SafeString")
		ßStringInput := πg.InternStr("StringInput")
		ßSubstitutionDef := πg.InternStr("SubstitutionDef")
		ßTestDirective := πg.InternStr("TestDirective")
		ßText := πg.InternStr("Text")
		ßTitle := πg.InternStr("Title")
		ßTrue := πg.InternStr("True")
		ßURLError := πg.InternStr("URLError")
		ßUnicode := πg.InternStr("Unicode")
		ßUnicodeDecodeError := πg.InternStr("UnicodeDecodeError")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ßUnicodeError := πg.InternStr("UnicodeError")
		ßValueError := πg.InternStr("ValueError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__file__ := πg.InternStr("__file__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_roles := πg.InternStr("_roles")
		ßabspath := πg.InternStr("abspath")
		ßadd := πg.InternStr("add")
		ßadd_name := πg.InternStr("add_name")
		ßappend := πg.InternStr("append")
		ßargument_pattern := πg.InternStr("argument_pattern")
		ßarguments := πg.InternStr("arguments")
		ßassert_has_content := πg.InternStr("assert_has_content")
		ßattributes := πg.InternStr("attributes")
		ßbackrefs := πg.InternStr("backrefs")
		ßblock_text := πg.InternStr("block_text")
		ßchildren := πg.InternStr("children")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcode := πg.InternStr("code")
		ßcomment_pattern := πg.InternStr("comment_pattern")
		ßcompile := πg.InternStr("compile")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßconvert_directive_function := πg.InternStr("convert_directive_function")
		ßcurrent_source := πg.InternStr("current_source")
		ßdecode := πg.InternStr("decode")
		ßdirective := πg.InternStr("directive")
		ßdirectives := πg.InternStr("directives")
		ßdirname := πg.InternStr("dirname")
		ßdocument := πg.InternStr("document")
		ßencode := πg.InternStr("encode")
		ßencoding := πg.InternStr("encoding")
		ßendswith := πg.InternStr("endswith")
		ßerror := πg.InternStr("error")
		ßexpandtabs := πg.InternStr("expandtabs")
		ßextend := πg.InternStr("extend")
		ßfile := πg.InternStr("file")
		ßfile_insertion_enabled := πg.InternStr("file_insertion_enabled")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßfind := πg.InternStr("find")
		ßflag := πg.InternStr("flag")
		ßformat := πg.InternStr("format")
		ßgeneric_custom_role := πg.InternStr("generic_custom_role")
		ßget := πg.InternStr("get")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßgroup := πg.InternStr("group")
		ßhas_content := πg.InternStr("has_content")
		ßhasattr := πg.InternStr("hasattr")
		ßinclude := πg.InternStr("include")
		ßinfo := πg.InternStr("info")
		ßinline := πg.InternStr("inline")
		ßinput_encoding := πg.InternStr("input_encoding")
		ßinput_encoding_error_handler := πg.InternStr("input_encoding_error_handler")
		ßinput_lines := πg.InternStr("input_lines")
		ßinput_offset := πg.InternStr("input_offset")
		ßinsert_input := πg.InternStr("insert_input")
		ßint := πg.InternStr("int")
		ßio := πg.InternStr("io")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlanguage := πg.InternStr("language")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßlineno := πg.InternStr("lineno")
		ßliteral := πg.InternStr("literal")
		ßliteral_block := πg.InternStr("literal_block")
		ßlocale_encoding := πg.InternStr("locale_encoding")
		ßlower := πg.InternStr("lower")
		ßltrim := πg.InternStr("ltrim")
		ßmatch := πg.InternStr("match")
		ßmisc := πg.InternStr("misc")
		ßname := πg.InternStr("name")
		ßnested_parse := πg.InternStr("nested_parse")
		ßnode := πg.InternStr("node")
		ßnodes := πg.InternStr("nodes")
		ßnormpath := πg.InternStr("normpath")
		ßnote_pending := πg.InternStr("note_pending")
		ßoption := πg.InternStr("option")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptional_arguments := πg.InternStr("optional_arguments")
		ßoptions := πg.InternStr("options")
		ßos := πg.InternStr("os")
		ßparagraph := πg.InternStr("paragraph")
		ßparse_directive_block := πg.InternStr("parse_directive_block")
		ßpath := πg.InternStr("path")
		ßpending := πg.InternStr("pending")
		ßpop := πg.InternStr("pop")
		ßraw := πg.InternStr("raw")
		ßraw_enabled := πg.InternStr("raw_enabled")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßreadlines := πg.InternStr("readlines")
		ßrecord_dependencies := πg.InternStr("record_dependencies")
		ßregister_local_role := πg.InternStr("register_local_role")
		ßrelative_path := πg.InternStr("relative_path")
		ßreplace := πg.InternStr("replace")
		ßreporter := πg.InternStr("reporter")
		ßreprunicode := πg.InternStr("reprunicode")
		ßrequired_arguments := πg.InternStr("required_arguments")
		ßrole := πg.InternStr("role")
		ßroles := πg.InternStr("roles")
		ßrtrim := πg.InternStr("rtrim")
		ßrun := πg.InternStr("run")
		ßset_classes := πg.InternStr("set_classes")
		ßsettings := πg.InternStr("settings")
		ßsevere := πg.InternStr("severe")
		ßsimplename := πg.InternStr("simplename")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßsplitlines := πg.InternStr("splitlines")
		ßstandard_include_path := πg.InternStr("standard_include_path")
		ßstartswith := πg.InternStr("startswith")
		ßstate := πg.InternStr("state")
		ßstate_machine := πg.InternStr("state_machine")
		ßstatemachine := πg.InternStr("statemachine")
		ßstates := πg.InternStr("states")
		ßstrftime := πg.InternStr("strftime")
		ßstring2lines := πg.InternStr("string2lines")
		ßsys := πg.InternStr("sys")
		ßsystem_message := πg.InternStr("system_message")
		ßtab_width := πg.InternStr("tab_width")
		ßtime := πg.InternStr("time")
		ßtitle := πg.InternStr("title")
		ßtrim := πg.InternStr("trim")
		ßunchanged := πg.InternStr("unchanged")
		ßunchanged_required := πg.InternStr("unchanged_required")
		ßunicode_code := πg.InternStr("unicode_code")
		ßuri := πg.InternStr("uri")
		ßurl := πg.InternStr("url")
		ßurlopen := πg.InternStr("urlopen")
		ßutils := πg.InternStr("utils")
		ßversion_info := πg.InternStr("version_info")
		ßwarning := πg.InternStr("warning")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 5: """Miscellaneous directives."""
			πF.SetLineno(5)
			// line 5: """Miscellaneous directives."""
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Miscellaneous directives.").ToObject()); πE != nil {
				continue
			}
			// line 7: __docformat__ = 'reStructuredText'
			πF.SetLineno(7)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 9: import sys
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import os.path
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
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
			// line 12: import time
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from docutils import io, nodes, statemachine, utils
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.statemachine"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßstatemachine.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: from docutils.utils.error_reporting import SafeString, ErrorString
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSafeString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSafeString.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorString.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: from docutils.utils.error_reporting import locale_encoding
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßlocale_encoding); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 16: from docutils.parsers.rst import Directive, convert_directive_function
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDirective); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDirective.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßconvert_directive_function); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßconvert_directive_function.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: from docutils.parsers.rst import directives, roles, states
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.roles"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.states"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßstates.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: from docutils.parsers.rst.directives.body import CodeBlock, NumberLines
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives.body"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[4]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßCodeBlock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCodeBlock.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[4]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßNumberLines); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumberLines.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from docutils.parsers.rst.roles import set_classes
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.roles"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßset_classes); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßset_classes.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: from docutils.transforms import misc
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.misc"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßmisc.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: class Include(Directive):
			πF.SetLineno(22)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Include", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 []πg.Param
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 24: """
					πF.SetLineno(24)
					// line 24: """
					πF.SetLineno(24)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Include content read from a separate source file.\n\n    Content may be parsed by the parser, or included as a literal\n    block.  The encoding of the included file can be specified.  Only\n    a part of the given file argument may be included by specifying\n    start and end line or text to match before and/or after the text\n    to be used.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 34: required_arguments = 1
					πF.SetLineno(34)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 35: optional_arguments = 0
					πF.SetLineno(35)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 36: final_argument_whitespace = True
					πF.SetLineno(36)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 37: option_spec = {'literal': directives.flag,
					πF.SetLineno(37)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßflag, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßliteral.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßcode.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßencoding, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßencoding.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, πg.NewStr("tab-width").ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, πg.NewStr("start-line").ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, πg.NewStr("end-line").ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, πg.NewStr("start-after").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, πg.NewStr("end-before").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, πg.NewStr("number-lines").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 50: standard_include_path = os.path.join(os.path.dirname(states.__file__),
					πF.SetLineno(50)
					πTemp004 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßstates); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__file__, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp003
					πTemp004[1] = ßinclude.ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßstandard_include_path.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 53: def run(self):
					πF.SetLineno(53)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µsource_dir *πg.Object = πg.UnboundLocal
						_ = µsource_dir
						var µpath *πg.Object = πg.UnboundLocal
						_ = µpath
						var µencoding *πg.Object = πg.UnboundLocal
						_ = µencoding
						var µe_handler *πg.Object = πg.UnboundLocal
						_ = µe_handler
						var µtab_width *πg.Object = πg.UnboundLocal
						_ = µtab_width
						var µinclude_file *πg.Object = πg.UnboundLocal
						_ = µinclude_file
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µstartline *πg.Object = πg.UnboundLocal
						_ = µstartline
						var µendline *πg.Object = πg.UnboundLocal
						_ = µendline
						var µlines *πg.Object = πg.UnboundLocal
						_ = µlines
						var µrawtext *πg.Object = πg.UnboundLocal
						_ = µrawtext
						var µafter_text *πg.Object = πg.UnboundLocal
						_ = µafter_text
						var µafter_index *πg.Object = πg.UnboundLocal
						_ = µafter_index
						var µbefore_text *πg.Object = πg.UnboundLocal
						_ = µbefore_text
						var µbefore_index *πg.Object = πg.UnboundLocal
						_ = µbefore_index
						var µinclude_lines *πg.Object = πg.UnboundLocal
						_ = µinclude_lines
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µliteral_block *πg.Object = πg.UnboundLocal
						_ = µliteral_block
						var µtokens *πg.Object = πg.UnboundLocal
						_ = µtokens
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µcodeblock *πg.Object = πg.UnboundLocal
						_ = µcodeblock
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 40:
								goto Label40
							case 39:
								goto Label39
							case 34:
								goto Label34
							case 11:
								goto Label11
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 54: """Include a file as part of the content of this reST file."""
							πF.SetLineno(54)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfile_insertion_enabled, nil); πE != nil {
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
								goto Label1
							}
							goto Label2
							// line 55: if not self.state.document.settings.file_insertion_enabled:
							πF.SetLineno(55)
						Label1:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" directive disabled.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 56: raise self.warning('"%s" directive disabled.' % self.name)
							πF.SetLineno(56)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 57: source = self.state_machine.input_lines.source(
							πF.SetLineno(57)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsource, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µsource = πTemp002
							// line 59: source_dir = os.path.dirname(os.path.abspath(source))
							πF.SetLineno(59)
							πTemp005 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp008[0] = µsource
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µsource_dir = πTemp002
							// line 60: path = directives.path(self.arguments[0])
							πF.SetLineno(60)
							πTemp005 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpath = πTemp001
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("<").ToObject()
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpath, ßstartswith, nil); πE != nil {
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
								goto Label3
							}
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr(">").ToObject()
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpath, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp003
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 61: if path.startswith('<') and path.endswith('>'):
							πF.SetLineno(61)
						Label4:
							// line 62: path = os.path.join(self.standard_include_path, path[1:-1])
							πF.SetLineno(62)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstandard_include_path, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpath, πTemp001); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpath = πTemp002
							goto Label5
						Label5:
							// line 63: path = os.path.normpath(os.path.join(source_dir, path))
							πF.SetLineno(63)
							πTemp005 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsource_dir, "source_dir"); πE != nil {
								continue
							}
							πTemp008[0] = µsource_dir
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp008[1] = µpath
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßnormpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpath = πTemp002
							// line 64: path = utils.relative_path(None, path)
							πF.SetLineno(64)
							πTemp005 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
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
							// line 65: path = nodes.reprunicode(path)
							πF.SetLineno(65)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005[0] = µpath
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreprunicode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpath = πTemp001
							// line 66: encoding = self.options.get(
							πF.SetLineno(66)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßencoding.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinput_encoding, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µencoding = πTemp001
							// line 68: e_handler=self.state.document.settings.input_encoding_error_handler
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinput_encoding_error_handler, nil); πE != nil {
								continue
							}
							µe_handler = πTemp002
							// line 69: tab_width = self.options.get(
							πF.SetLineno(69)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("tab-width").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtab_width, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtab_width = πTemp001
							// line 71: try:
							πF.SetLineno(71)
							πF.PushCheckpoint(7)
							// line 72: self.state.document.settings.record_dependencies.add(path)
							πF.SetLineno(72)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
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
							// line 73: include_file = io.FileInput(source_path=path,
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µe_handler, "e_handler"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"source_path", µpath},
								{"encoding", µencoding},
								{"error_handler", µe_handler},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFileInput, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp009); πE != nil {
								continue
							}
							µinclude_file = πTemp001
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 76: except UnicodeEncodeError as error:
							πF.SetLineno(76)
						Label8:
							µerror = πTemp010.ToObject()
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp008[0] = µpath
							if πTemp006, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problems with \"%s\" directive path:\nCannot encode input file path \"%s\" (wrong locale?).").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 77: raise self.severe(u'Problems with "%s" directive path:\n'
							πF.SetLineno(77)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label6
							// line 81: except IOError as error:
							πF.SetLineno(81)
						Label9:
							µerror = πTemp010.ToObject()
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp008[0] = µerror
							if πTemp006, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problems with \"%s\" directive path:\n%s.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 82: raise self.severe(u'Problems with "%s" directive path:\n%s.' %
							πF.SetLineno(82)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							// line 84: startline = self.options.get('start-line', None)
							πF.SetLineno(84)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("start-line").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µstartline = πTemp001
							// line 85: endline = self.options.get('end-line', None)
							πF.SetLineno(85)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("end-line").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µendline = πTemp001
							// line 86: try:
							πF.SetLineno(86)
							πF.PushCheckpoint(11)
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							πTemp001 = µstartline
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µendline, "endline"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µendline != πTemp003).ToObject()
							πTemp001 = πTemp002
						Label12:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 87: if startline or (endline is not None):
							πF.SetLineno(87)
						Label13:
							// line 88: lines = include_file.readlines()
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µinclude_file, "include_file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinclude_file, ßreadlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlines = πTemp002
							// line 89: rawtext = ''.join(lines[startline:endline])
							πF.SetLineno(89)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µendline, "endline"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstartline, µendline, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µlines, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µrawtext = πTemp002
							goto Label15
						Label14:
							// line 91: rawtext = include_file.read()
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µinclude_file, "include_file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinclude_file, ßread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrawtext = πTemp002
							goto Label15
						Label15:
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 92: except UnicodeError as error:
							πF.SetLineno(92)
						Label16:
							µerror = πTemp010.ToObject()
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp008[0] = µerror
							if πTemp006, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problem with \"%s\" directive:\n%s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 93: raise self.severe(u'Problem with "%s" directive:\n%s' %
							πF.SetLineno(93)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							// line 97: after_text = self.options.get('start-after', None)
							πF.SetLineno(97)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("start-after").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µafter_text = πTemp001
							if πE = πg.CheckLocal(πF, µafter_text, "after_text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µafter_text); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							goto Label18
							// line 98: if after_text:
							πF.SetLineno(98)
						Label17:
							// line 100: after_index = rawtext.find(after_text)
							πF.SetLineno(100)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µafter_text, "after_text"); πE != nil {
								continue
							}
							πTemp005[0] = µafter_text
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrawtext, ßfind, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µafter_index = πTemp002
							if πE = πg.CheckLocal(πF, µafter_index, "after_index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µafter_index, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label19
							}
							goto Label20
							// line 101: if after_index < 0:
							πF.SetLineno(101)
						Label19:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Problem with \"start-after\" option of \"%s\" directive:\nText not found.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 102: raise self.severe('Problem with "start-after" option of "%s" '
							πF.SetLineno(102)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label20
						Label20:
							// line 104: rawtext = rawtext[after_index + len(after_text):]
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µafter_index, "after_index"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µafter_text, "after_text"); πE != nil {
								continue
							}
							πTemp005[0] = µafter_text
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Add(πF, µafter_index, πTemp006); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrawtext, πTemp001); πE != nil {
								continue
							}
							µrawtext = πTemp002
							goto Label18
						Label18:
							// line 105: before_text = self.options.get('end-before', None)
							πF.SetLineno(105)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("end-before").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µbefore_text = πTemp001
							if πE = πg.CheckLocal(πF, µbefore_text, "before_text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µbefore_text); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 106: if before_text:
							πF.SetLineno(106)
						Label21:
							// line 108: before_index = rawtext.find(before_text)
							πF.SetLineno(108)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbefore_text, "before_text"); πE != nil {
								continue
							}
							πTemp005[0] = µbefore_text
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrawtext, ßfind, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µbefore_index = πTemp002
							if πE = πg.CheckLocal(πF, µbefore_index, "before_index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µbefore_index, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							goto Label24
							// line 109: if before_index < 0:
							πF.SetLineno(109)
						Label23:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Problem with \"end-before\" option of \"%s\" directive:\nText not found.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 110: raise self.severe('Problem with "end-before" option of "%s" '
							πF.SetLineno(110)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label24
						Label24:
							// line 112: rawtext = rawtext[:before_index]
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µbefore_index, "before_index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µbefore_index, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrawtext, πTemp001); πE != nil {
								continue
							}
							µrawtext = πTemp002
							goto Label22
						Label22:
							// line 114: include_lines = statemachine.string2lines(rawtext, tab_width,
							πF.SetLineno(114)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							πTemp005[0] = µrawtext
							if πE = πg.CheckLocal(πF, µtab_width, "tab_width"); πE != nil {
								continue
							}
							πTemp005[1] = µtab_width
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"convert_whitespace", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstatemachine); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstring2lines, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µinclude_lines = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp002, ßliteral.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label25
							}
							goto Label26
							// line 116: if 'literal' in self.options:
							πF.SetLineno(116)
						Label25:
							if πE = πg.CheckLocal(πF, µtab_width, "tab_width"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µtab_width, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label27
							}
							goto Label28
							// line 118: if tab_width >= 0:
							πF.SetLineno(118)
						Label27:
							// line 119: text = rawtext.expandtabs(tab_width)
							πF.SetLineno(119)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtab_width, "tab_width"); πE != nil {
								continue
							}
							πTemp005[0] = µtab_width
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrawtext, ßexpandtabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtext = πTemp002
							goto Label29
						Label28:
							// line 121: text = rawtext
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							µtext = µrawtext
							goto Label29
						Label29:
							// line 122: literal_block = nodes.literal_block(rawtext, source=path,
							πF.SetLineno(122)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							πTemp005[0] = µrawtext
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = ßclass.ToObject()
							πTemp012 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp012...).ToObject()
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp009 = πg.KWArgs{
								{"source", µpath},
								{"classes", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µliteral_block = πTemp001
							// line 124: literal_block.line = 1
							πF.SetLineno(124)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µliteral_block, "literal_block"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µliteral_block, ßline, πTemp001); πE != nil {
								continue
							}
							// line 125: self.add_name(literal_block)
							πF.SetLineno(125)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µliteral_block, "literal_block"); πE != nil {
								continue
							}
							πTemp005[0] = µliteral_block
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp002, πg.NewStr("number-lines").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label30
							}
							goto Label31
							// line 126: if 'number-lines' in self.options:
							πF.SetLineno(126)
						Label30:
							// line 127: try:
							πF.SetLineno(127)
							πF.PushCheckpoint(34)
							// line 128: startline = int(self.options['number-lines'] or 1)
							πF.SetLineno(128)
							πTemp005 = πF.MakeArgs(1)
							πTemp002 = πg.NewStr("number-lines").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label35
							}
							πTemp001 = πg.NewInt(1).ToObject()
						Label35:
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µstartline = πTemp002
							πF.PopCheckpoint()
							goto Label33
						Label34:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label36
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 129: except ValueError:
							πF.SetLineno(129)
						Label36:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr(":number-lines: with non-integer start value").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 130: raise self.error(':number-lines: with non-integer '
							πF.SetLineno(130)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label33
						Label33:
							// line 132: endline = startline + len(include_lines)
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinclude_lines, "include_lines"); πE != nil {
								continue
							}
							πTemp005[0] = µinclude_lines
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, µstartline, πTemp003); πE != nil {
								continue
							}
							µendline = πTemp001
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label37
							}
							goto Label38
							// line 133: if text.endswith('\n'):
							πF.SetLineno(133)
						Label37:
							// line 134: text = text[:-1]
							πF.SetLineno(134)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtext, πTemp001); πE != nil {
								continue
							}
							µtext = πTemp002
							goto Label38
						Label38:
							// line 135: tokens = NumberLines([([], text)], startline, endline)
							πF.SetLineno(135)
							πTemp005 = πF.MakeArgs(3)
							πTemp008 = make([]*πg.Object, 1)
							πTemp012 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp012...).ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, µtext).ToObject()
							πTemp008[0] = πTemp001
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							πTemp005[1] = µstartline
							if πE = πg.CheckLocal(πF, µendline, "endline"); πE != nil {
								continue
							}
							πTemp005[2] = µendline
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNumberLines); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtokens = πTemp002
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtokens); πE != nil {
								continue
							}
							πF.PushCheckpoint(40)
							πTemp004 = false
						Label39:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label41
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µclasses = πTemp003
								µvalue = πTemp006
							}
							if πE != nil || !πTemp013 {
								continue
							}
							πF.PushCheckpoint(39)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.IsTrue(πF, µclasses); πE != nil {
								continue
							}
							if πTemp013 {
								goto Label42
							}
							goto Label43
							// line 137: if classes:
							πF.SetLineno(137)
						Label42:
							// line 138: literal_block += nodes.inline(value, value,
							πF.SetLineno(138)
							if πE = πg.CheckLocal(πF, µliteral_block, "literal_block"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[1] = µvalue
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"classes", µclasses},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp005, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.IAdd(πF, µliteral_block, πTemp002); πE != nil {
								continue
							}
							µliteral_block = πTemp003
							goto Label44
						Label43:
							// line 141: literal_block += nodes.Text(value)
							πF.SetLineno(141)
							if πE = πg.CheckLocal(πF, µliteral_block, "literal_block"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßText, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.IAdd(πF, µliteral_block, πTemp002); πE != nil {
								continue
							}
							µliteral_block = πTemp003
							goto Label44
						Label44:
							continue
						Label40:
							if πE != nil || πR != nil {
								continue
							}
						Label41:
							goto Label32
						Label31:
							// line 143: literal_block += nodes.Text(text)
							πF.SetLineno(143)
							if πE = πg.CheckLocal(πF, µliteral_block, "literal_block"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp005[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.IAdd(πF, µliteral_block, πTemp001); πE != nil {
								continue
							}
							µliteral_block = πTemp002
							goto Label32
						Label32:
							// line 144: return [literal_block]
							πF.SetLineno(144)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µliteral_block, "literal_block"); πE != nil {
								continue
							}
							πTemp005[0] = µliteral_block
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πR = πTemp001
							continue
							goto Label26
						Label26:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp002, ßcode.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label45
							}
							goto Label46
							// line 145: if 'code' in self.options:
							πF.SetLineno(145)
						Label45:
							// line 146: self.options['source'] = path
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpath); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp003 = ßsource.ToObject()
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtab_width, "tab_width"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µtab_width, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label47
							}
							goto Label48
							// line 148: if tab_width < 0:
							πF.SetLineno(148)
						Label47:
							// line 149: include_lines = rawtext.splitlines()
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrawtext, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µinclude_lines = πTemp002
							goto Label48
						Label48:
							// line 150: codeblock = CodeBlock(self.name,
							πF.SetLineno(150)
							πTemp005 = πF.MakeArgs(9)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp008 = make([]*πg.Object, 1)
							πTemp012 = πF.MakeArgs(1)
							πTemp012[0] = ßcode.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp008[0] = πTemp001
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							if πE = πg.CheckLocal(πF, µinclude_lines, "include_lines"); πE != nil {
								continue
							}
							πTemp005[3] = µinclude_lines
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp005[4] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp005[5] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[6] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp005[7] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							πTemp005[8] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCodeBlock); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µcodeblock = πTemp002
							// line 159: return codeblock.run()
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µcodeblock, "codeblock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcodeblock, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label46
						Label46:
							// line 160: self.state_machine.insert_input(include_lines, path)
							πF.SetLineno(160)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinclude_lines, "include_lines"); πE != nil {
								continue
							}
							πTemp005[0] = µinclude_lines
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp005[1] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert_input, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 161: return []
							πF.SetLineno(161)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 54: """Include a file as part of the content of this reST file."""
					πF.SetLineno(54)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Include a file as part of the content of this reST file.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßrun); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Include").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInclude.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 164: class Raw(Directive):
			πF.SetLineno(164)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Raw", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 166: """
					πF.SetLineno(166)
					// line 166: """
					πF.SetLineno(166)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Pass through content unchanged\n\n    Content is included in output based on type argument\n\n    Content may be included inline (content section of directive) or\n    imported from a file or url.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 175: required_arguments = 1
					πF.SetLineno(175)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 176: optional_arguments = 0
					πF.SetLineno(176)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 177: final_argument_whitespace = True
					πF.SetLineno(177)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 178: option_spec = {'file': directives.path,
					πF.SetLineno(178)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßfile.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßuri, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßurl.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßencoding, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßencoding.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 181: has_content = True
					πF.SetLineno(181)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 183: def run(self):
					πF.SetLineno(183)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µattributes *πg.Object = πg.UnboundLocal
						_ = µattributes
						var µencoding *πg.Object = πg.UnboundLocal
						_ = µencoding
						var µe_handler *πg.Object = πg.UnboundLocal
						_ = µe_handler
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µsource_dir *πg.Object = πg.UnboundLocal
						_ = µsource_dir
						var µpath *πg.Object = πg.UnboundLocal
						_ = µpath
						var µraw_file *πg.Object = πg.UnboundLocal
						_ = µraw_file
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µurlopen *πg.Object = πg.UnboundLocal
						_ = µurlopen
						var µURLError *πg.Object = πg.UnboundLocal
						_ = µURLError
						var µraw_text *πg.Object = πg.UnboundLocal
						_ = µraw_text
						var µraw_node *πg.Object = πg.UnboundLocal
						_ = µraw_node
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
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.Dict
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 πg.KWArgs
						_ = πTemp013
						var πTemp014 *πg.BaseException
						_ = πTemp014
						var πTemp015 *πg.Traceback
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
							case 17:
								goto Label17
							case 26:
								goto Label26
							case 20:
								goto Label20
							case 29:
								goto Label29
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßraw_enabled, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
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
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßfile_insertion_enabled, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
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
							if πTemp007, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, πTemp007, ßfile.ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp009).ToObject()
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
							if πTemp007, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, πTemp007, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp009).ToObject()
							πTemp004 = πTemp005
						Label3:
							πTemp003 = πTemp004
						Label2:
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 184: if (not self.state.document.settings.raw_enabled
							πF.SetLineno(184)
						Label4:
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" directive disabled.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 188: raise self.warning('"%s" directive disabled.' % self.name)
							πF.SetLineno(188)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 189: attributes = {'format': ' '.join(self.arguments[0].lower().split())}
							πF.SetLineno(189)
							πTemp011 = πg.NewDict()
							πTemp010 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp003
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πE = πTemp011.SetItem(πF, ßformat.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp011.ToObject()
							µattributes = πTemp001
							// line 190: encoding = self.options.get(
							πF.SetLineno(190)
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = ßencoding.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinput_encoding, nil); πE != nil {
								continue
							}
							πTemp010[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µencoding = πTemp001
							// line 192: e_handler=self.state.document.settings.input_encoding_error_handler
							πF.SetLineno(192)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinput_encoding_error_handler, nil); πE != nil {
								continue
							}
							µe_handler = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp003, ßfile.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp003, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							goto Label9
							// line 193: if self.content:
							πF.SetLineno(193)
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, ßfile.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp003
						Label11:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label12
							}
							goto Label13
							// line 194: if 'file' in self.options or 'url' in self.options:
							πF.SetLineno(194)
						Label12:
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" directive may not both specify an external file and have content.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 195: raise self.error(
							πF.SetLineno(195)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label13
						Label13:
							// line 198: text = '\n'.join(self.content)
							πF.SetLineno(198)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µtext = πTemp003
							goto Label10
							// line 199: elif 'file' in self.options:
							πF.SetLineno(199)
						Label7:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp003, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label14
							}
							goto Label15
							// line 200: if 'url' in self.options:
							πF.SetLineno(200)
						Label14:
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("The \"file\" and \"url\" options may not be simultaneously specified for the \"%s\" directive.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 201: raise self.error(
							πF.SetLineno(201)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label15
						Label15:
							// line 204: source_dir = os.path.dirname(
							πF.SetLineno(204)
							πTemp010 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßcurrent_source, nil); πE != nil {
								continue
							}
							πTemp012[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp010[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µsource_dir = πTemp003
							// line 206: path = os.path.normpath(os.path.join(source_dir,
							πF.SetLineno(206)
							πTemp010 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsource_dir, "source_dir"); πE != nil {
								continue
							}
							πTemp012[0] = µsource_dir
							πTemp001 = ßfile.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp012[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp010[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µpath = πTemp003
							// line 208: path = utils.relative_path(None, path)
							πF.SetLineno(208)
							πTemp010 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp010[1] = µpath
							if πTemp001, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrelative_path, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µpath = πTemp001
							// line 209: try:
							πF.SetLineno(209)
							πF.PushCheckpoint(17)
							// line 210: raw_file = io.FileInput(source_path=path,
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µe_handler, "e_handler"); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"source_path", µpath},
								{"encoding", µencoding},
								{"error_handler", µe_handler},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFileInput, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, πTemp013); πE != nil {
								continue
							}
							µraw_file = πTemp001
							// line 215: self.state.document.settings.record_dependencies.add(path)
							πF.SetLineno(215)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp010[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πF.PopCheckpoint()
							goto Label16
						Label17:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp014, πTemp015 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp014.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label18
							}
							πE = πF.Raise(πTemp014.ToObject(), nil, πTemp015.ToObject())
							continue
							// line 216: except IOError as error:
							πF.SetLineno(216)
						Label18:
							µerror = πTemp014.ToObject()
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp012[0] = µerror
							if πTemp005, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problems with \"%s\" directive path:\n%s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 217: raise self.severe(u'Problems with "%s" directive path:\n%s.'
							πF.SetLineno(217)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label16
						Label16:
							// line 219: try:
							πF.SetLineno(219)
							πF.PushCheckpoint(20)
							// line 220: text = raw_file.read()
							πF.SetLineno(220)
							if πE = πg.CheckLocal(πF, µraw_file, "raw_file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µraw_file, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp003
							πF.PopCheckpoint()
							goto Label19
						Label20:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp014, πTemp015 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp014.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label21
							}
							πE = πF.Raise(πTemp014.ToObject(), nil, πTemp015.ToObject())
							continue
							// line 221: except UnicodeError as error:
							πF.SetLineno(221)
						Label21:
							µerror = πTemp014.ToObject()
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp012[0] = µerror
							if πTemp005, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problem with \"%s\" directive:\n%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 222: raise self.severe(u'Problem with "%s" directive:\n%s'
							πF.SetLineno(222)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label19
						Label19:
							// line 224: attributes['source'] = path
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpath); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							πTemp003 = ßsource.ToObject()
							if πE = πg.SetItem(πF, µattributes, πTemp003, πTemp001); πE != nil {
								continue
							}
							goto Label10
							// line 225: elif 'url' in self.options:
							πF.SetLineno(225)
						Label8:
							// line 226: source = self.options['url']
							πF.SetLineno(226)
							πTemp001 = ßurl.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µsource = πTemp003
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
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label22
							}
							goto Label23
							// line 230: if sys.version_info >= (3, 0):
							πF.SetLineno(230)
						Label22:
							// line 231: from urllib.request import urlopen
							πF.SetLineno(231)
							if πTemp010, πE = πg.ImportModule(πF, "urllib.request"); πE != nil {
								continue
							}
							πTemp001 = πTemp010[1]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßurlopen); πE != nil {
								continue
							}
							µurlopen = πTemp003
							// line 232: from urllib.error import URLError
							πF.SetLineno(232)
							if πTemp010, πE = πg.ImportModule(πF, "urllib.error"); πE != nil {
								continue
							}
							πTemp001 = πTemp010[1]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßURLError); πE != nil {
								continue
							}
							µURLError = πTemp003
							goto Label24
						Label23:
							// line 234: from urllib2 import urlopen, URLError
							πF.SetLineno(234)
							if πTemp010, πE = πg.ImportModule(πF, "urllib2"); πE != nil {
								continue
							}
							πTemp001 = πTemp010[0]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßurlopen); πE != nil {
								continue
							}
							µurlopen = πTemp003
							πTemp001 = πTemp010[0]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßURLError); πE != nil {
								continue
							}
							µURLError = πTemp003
							goto Label24
						Label24:
							// line 235: try:
							πF.SetLineno(235)
							πF.PushCheckpoint(26)
							// line 236: raw_text = urlopen(source).read()
							πF.SetLineno(236)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp010[0] = µsource
							if πE = πg.CheckLocal(πF, µurlopen, "urlopen"); πE != nil {
								continue
							}
							if πTemp001, πE = µurlopen.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µraw_text = πTemp001
							πF.PopCheckpoint()
							goto Label25
						Label26:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp014, πTemp015 = πF.ExcInfo()
							if πE = πg.CheckLocal(πF, µURLError, "URLError"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µURLError, πTemp003, πTemp004).ToObject()
							if πTemp002, πE = πg.IsInstance(πF, πTemp014.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label27
							}
							πE = πF.Raise(πTemp014.ToObject(), nil, πTemp015.ToObject())
							continue
							// line 237: except (URLError, IOError, OSError) as error:
							πF.SetLineno(237)
						Label27:
							µerror = πTemp014.ToObject()
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp005 = ßurl.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp016, πTemp005); πE != nil {
								continue
							}
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp012[0] = µerror
							if πTemp005, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp003 = πg.NewTuple3(πTemp004, πTemp007, πTemp016).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problems with \"%s\" directive URL \"%s\":\n%s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 238: raise self.severe(u'Problems with "%s" directive URL "%s":\n%s.'
							πF.SetLineno(238)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label25
						Label25:
							// line 240: raw_file = io.StringInput(source=raw_text, source_path=source,
							πF.SetLineno(240)
							if πE = πg.CheckLocal(πF, µraw_text, "raw_text"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µe_handler, "e_handler"); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"source", µraw_text},
								{"source_path", µsource},
								{"encoding", µencoding},
								{"error_handler", µe_handler},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringInput, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, πTemp013); πE != nil {
								continue
							}
							µraw_file = πTemp001
							// line 243: try:
							πF.SetLineno(243)
							πF.PushCheckpoint(29)
							// line 244: text = raw_file.read()
							πF.SetLineno(244)
							if πE = πg.CheckLocal(πF, µraw_file, "raw_file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µraw_file, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp003
							πF.PopCheckpoint()
							goto Label28
						Label29:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp014, πTemp015 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp014.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label30
							}
							πE = πF.Raise(πTemp014.ToObject(), nil, πTemp015.ToObject())
							continue
							// line 245: except UnicodeError as error:
							πF.SetLineno(245)
						Label30:
							µerror = πTemp014.ToObject()
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp012[0] = µerror
							if πTemp005, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Problem with \"%s\" directive:\n%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							// line 246: raise self.severe(u'Problem with "%s" directive:\n%s'
							πF.SetLineno(246)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label28
						Label28:
							// line 248: attributes['source'] = source
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							πTemp003 = ßsource.ToObject()
							if πE = πg.SetItem(πF, µattributes, πTemp003, πTemp001); πE != nil {
								continue
							}
							goto Label10
						Label9:
							// line 251: self.assert_has_content()
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label10
						Label10:
							// line 252: raw_node = nodes.raw('', text, **attributes)
							πF.SetLineno(252)
							πTemp010 = πF.MakeArgs(2)
							πTemp010[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp010[1] = µtext
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßraw, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp003, πTemp010, nil, nil, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							µraw_node = πTemp001
							// line 253: (raw_node.source,
							πF.SetLineno(253)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_source_and_line, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µraw_node, "raw_node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µraw_node, ßsource, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µraw_node, "raw_node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µraw_node, ßline, πTemp004); πE != nil {
								continue
							}
							// line 255: return [raw_node]
							πF.SetLineno(255)
							πTemp010 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µraw_node, "raw_node"); πE != nil {
								continue
							}
							πTemp010[0] = µraw_node
							πTemp001 = πg.NewList(πTemp010...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Raw").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRaw.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 258: class Replace(Directive):
			πF.SetLineno(258)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Replace", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 260: has_content = True
					πF.SetLineno(260)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 262: def run(self):
					πF.SetLineno(262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µelement *πg.Object = πg.UnboundLocal
						_ = µelement
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µelem *πg.Object = πg.UnboundLocal
						_ = µelem
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
						var πTemp007 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSubstitutionDef, nil); πE != nil {
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
							// line 263: if not isinstance(self.state, states.SubstitutionDef):
							πF.SetLineno(263)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid context: the \"%s\" directive can only be used within a substitution definition.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 264: raise self.error(
							πF.SetLineno(264)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 267: self.assert_has_content()
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 268: text = '\n'.join(self.content)
							πF.SetLineno(268)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µtext = πTemp003
							// line 269: element = nodes.Element(text)
							πF.SetLineno(269)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp002[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßElement, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µelement = πTemp001
							// line 270: self.state.nested_parse(self.content, self.content_offset,
							πF.SetLineno(270)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp002[2] = µelement
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 273: node = None
							πF.SetLineno(273)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µnode = πTemp001
							// line 274: messages = []
							πF.SetLineno(274)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µmessages = πTemp001
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µelement); πE != nil {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µelem = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µnode); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label6
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp002[0] = µelem
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp008
						Label6:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp002[0] = µelem
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsystem_message, nil); πE != nil {
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 276: if not node and isinstance(elem, nodes.paragraph):
							πF.SetLineno(276)
						Label7:
							// line 277: node = elem
							πF.SetLineno(277)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							µnode = µelem
							goto Label10
							// line 278: elif isinstance(elem, nodes.system_message):
							πF.SetLineno(278)
						Label8:
							// line 279: elem['backrefs'] = []
							πF.SetLineno(279)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp008 = ßbackrefs.ToObject()
							if πE = πg.SetItem(πF, µelem, πTemp008, πTemp004); πE != nil {
								continue
							}
							// line 280: messages.append(elem)
							πF.SetLineno(280)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp002[0] = µelem
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label10
						Label9:
							// line 282: return [
							πF.SetLineno(282)
							πTemp002 = make([]*πg.Object, 1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Error in \"%s\" directive: may contain a single paragraph only.").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp009[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"line", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßerror, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp002[0] = πTemp004
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πR = πTemp003
							continue
							goto Label10
						Label10:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µnode); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 286: if node:
							πF.SetLineno(286)
						Label11:
							// line 287: return messages + node.children
							πF.SetLineno(287)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label12
						Label12:
							// line 288: return messages
							πF.SetLineno(288)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πR = µmessages
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Replace").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReplace.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 290: class Unicode(Directive):
			πF.SetLineno(290)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Unicode", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 292: r"""
					πF.SetLineno(292)
					// line 292: r"""
					πF.SetLineno(292)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Convert Unicode character codes (numbers) to characters.  Codes may be\n    decimal numbers, hexadecimal numbers (prefixed by ``0x``, ``x``, ``\\x``,\n    ``U+``, ``u``, or ``\\u``; e.g. ``U+262E``), or XML-style numeric character\n    entities (e.g. ``&#x262E;``).  Text following \"..\" is a comment and is\n    ignored.  Spaces are ignored, and any other text remains as-is.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 300: required_arguments = 1
					πF.SetLineno(300)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 301: optional_arguments = 0
					πF.SetLineno(301)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 302: final_argument_whitespace = True
					πF.SetLineno(302)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 303: option_spec = {'trim': directives.flag,
					πF.SetLineno(303)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßflag, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßtrim.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßflag, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßltrim.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßflag, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßrtrim.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 307: comment_pattern = re.compile(r'( |\n|^)\.\. ')
					πF.SetLineno(307)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("( |\\n|^)\\.\\. ").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßcomment_pattern.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 309: def run(self):
					πF.SetLineno(309)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsubstitution_definition *πg.Object = πg.UnboundLocal
						_ = µsubstitution_definition
						var µcodes *πg.Object = πg.UnboundLocal
						_ = µcodes
						var µelement *πg.Object = πg.UnboundLocal
						_ = µelement
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
						var µdecoded *πg.Object = πg.UnboundLocal
						_ = µdecoded
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
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
							case 13:
								goto Label13
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSubstitutionDef, nil); πE != nil {
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
							// line 310: if not isinstance(self.state, states.SubstitutionDef):
							πF.SetLineno(310)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid context: the \"%s\" directive can only be used within a substitution definition.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 311: raise self.error(
							πF.SetLineno(311)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 314: substitution_definition = self.state_machine.node
							πF.SetLineno(314)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnode, nil); πE != nil {
								continue
							}
							µsubstitution_definition = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, ßtrim.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 315: if 'trim' in self.options:
							πF.SetLineno(315)
						Label3:
							// line 316: substitution_definition.attributes['ltrim'] = 1
							πF.SetLineno(316)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubstitution_definition, "substitution_definition"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsubstitution_definition, ßattributes, nil); πE != nil {
								continue
							}
							πTemp004 = ßltrim.ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
								continue
							}
							// line 317: substitution_definition.attributes['rtrim'] = 1
							πF.SetLineno(317)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubstitution_definition, "substitution_definition"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsubstitution_definition, ßattributes, nil); πE != nil {
								continue
							}
							πTemp004 = ßrtrim.ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, ßltrim.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 318: if 'ltrim' in self.options:
							πF.SetLineno(318)
						Label5:
							// line 319: substitution_definition.attributes['ltrim'] = 1
							πF.SetLineno(319)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubstitution_definition, "substitution_definition"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsubstitution_definition, ßattributes, nil); πE != nil {
								continue
							}
							πTemp004 = ßltrim.ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
								continue
							}
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, ßrtrim.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 320: if 'rtrim' in self.options:
							πF.SetLineno(320)
						Label7:
							// line 321: substitution_definition.attributes['rtrim'] = 1
							πF.SetLineno(321)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubstitution_definition, "substitution_definition"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsubstitution_definition, ßattributes, nil); πE != nil {
								continue
							}
							πTemp004 = ßrtrim.ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
								continue
							}
							goto Label8
						Label8:
							// line 322: codes = self.comment_pattern.split(self.arguments[0])[0].split()
							πF.SetLineno(322)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcomment_pattern, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcodes = πTemp003
							// line 323: element = nodes.Element()
							πF.SetLineno(323)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßElement, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µelement = πTemp001
							if πE = πg.CheckLocal(πF, µcodes, "codes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µcodes); πE != nil {
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
								µcode = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(9)
							// line 325: try:
							πF.SetLineno(325)
							πF.PushCheckpoint(13)
							// line 326: decoded = directives.unicode_code(code)
							πF.SetLineno(326)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp002[0] = µcode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunicode_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdecoded = πTemp003
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label14
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 327: except ValueError as error:
							πF.SetLineno(327)
						Label14:
							µerror = πTemp009.ToObject()
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp011[0] = µerror
							if πTemp006, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp004 = πg.NewTuple2(µcode, πTemp007).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewUnicode("Invalid character code: %s\n%s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 328: raise self.error(u'Invalid character code: %s\n%s'
							πF.SetLineno(328)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							// line 330: element += nodes.Text(decoded)
							πF.SetLineno(330)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdecoded, "decoded"); πE != nil {
								continue
							}
							πTemp002[0] = µdecoded
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IAdd(πF, µelement, πTemp003); πE != nil {
								continue
							}
							µelement = πTemp004
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							// line 331: return element.children
							πF.SetLineno(331)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µelement, ßchildren, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Unicode").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnicode.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 334: class Class(Directive):
			πF.SetLineno(334)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Class", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 336: """
					πF.SetLineno(336)
					// line 336: """
					πF.SetLineno(336)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Set a \"class\" attribute on the directive content or the next element.\n    When applied to the next element, a \"pending\" element is inserted, and a\n    transform does the work later.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 342: required_arguments = 1
					πF.SetLineno(342)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 343: optional_arguments = 0
					πF.SetLineno(343)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 344: final_argument_whitespace = True
					πF.SetLineno(344)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 345: has_content = True
					πF.SetLineno(345)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 347: def run(self):
					πF.SetLineno(347)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µclass_value *πg.Object = πg.UnboundLocal
						_ = µclass_value
						var µnode_list *πg.Object = πg.UnboundLocal
						_ = µnode_list
						var µcontainer *πg.Object = πg.UnboundLocal
						_ = µcontainer
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Dict
						_ = πTemp012
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 2:
								goto Label2
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 348: try:
							πF.SetLineno(348)
							πF.PushCheckpoint(2)
							// line 349: class_value = directives.class_option(self.arguments[0])
							πF.SetLineno(349)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µclass_value = πTemp002
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 350: except ValueError:
							πF.SetLineno(350)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Invalid class attribute value for \"%s\" directive: \"%s\".").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 351: raise self.error(
							πF.SetLineno(351)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 354: node_list = []
							πF.SetLineno(354)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µnode_list = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 355: if self.content:
							πF.SetLineno(355)
						Label4:
							// line 356: container = nodes.Element()
							πF.SetLineno(356)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßElement, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcontainer = πTemp002
							// line 357: self.state.nested_parse(self.content, self.content_offset,
							πF.SetLineno(357)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							πTemp001[2] = µcontainer
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µcontainer); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp007 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								µnode = πTemp003
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 360: node['classes'].extend(class_value)
							πF.SetLineno(360)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclass_value, "class_value"); πE != nil {
								continue
							}
							πTemp001[0] = µclass_value
							πTemp003 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßextend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 361: node_list.extend(container.children)
							πF.SetLineno(361)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcontainer, ßchildren, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnode_list, "node_list"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode_list, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label5:
							// line 363: pending = nodes.pending(
							πF.SetLineno(363)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmisc); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßClassAttribute, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp012 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µclass_value, "class_value"); πE != nil {
								continue
							}
							if πE = πTemp012.SetItem(πF, ßclass.ToObject(), µclass_value); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πTemp012.SetItem(πF, ßdirective.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp012.ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpending, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µpending = πTemp002
							// line 367: self.state_machine.document.note_pending(pending)
							πF.SetLineno(367)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 368: node_list.append(pending)
							πF.SetLineno(368)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							if πE = πg.CheckLocal(πF, µnode_list, "node_list"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode_list, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							// line 369: return node_list
							πF.SetLineno(369)
							if πE = πg.CheckLocal(πF, µnode_list, "node_list"); πE != nil {
								continue
							}
							πR = µnode_list
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Class").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClass.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 372: class Role(Directive):
			πF.SetLineno(372)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []πg.Param
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 374: has_content = True
					πF.SetLineno(374)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 376: argument_pattern = re.compile(r'(%s)\s*(\(\s*(%s)\s*\)\s*)?$'
					πF.SetLineno(376)
					πTemp002 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßstates); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßInliner, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßsimplename, nil); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple1(πTemp005).ToObject()
					if πTemp003, πE = πg.Mul(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("(%s)\\s*(\\(\\s*(%s)\\s*\\)\\s*)?$").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
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
					if πE = πClass.SetItem(πF, ßargument_pattern.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 379: def run(self):
					πF.SetLineno(379)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var µmatch *πg.Object = πg.UnboundLocal
						_ = µmatch
						var µnew_role_name *πg.Object = πg.UnboundLocal
						_ = µnew_role_name
						var µbase_role_name *πg.Object = πg.UnboundLocal
						_ = µbase_role_name
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µbase_role *πg.Object = πg.UnboundLocal
						_ = µbase_role
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µconverted_role *πg.Object = πg.UnboundLocal
						_ = µconverted_role
						var µarguments *πg.Object = πg.UnboundLocal
						_ = µarguments
						var µoptions *πg.Object = πg.UnboundLocal
						_ = µoptions
						var µcontent *πg.Object = πg.UnboundLocal
						_ = µcontent
						var µcontent_offset *πg.Object = πg.UnboundLocal
						_ = µcontent_offset
						var µdetail *πg.Object = πg.UnboundLocal
						_ = µdetail
						var µrole *πg.Object = πg.UnboundLocal
						_ = µrole
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 *πg.Dict
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.BaseException
						_ = πTemp012
						var πTemp013 *πg.Traceback
						_ = πTemp013
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 17:
								goto Label17
							case 12:
								goto Label12
							default:
								panic("unexpected function state")
							}
							// line 380: """Dynamically create and register a custom interpreted text role."""
							πF.SetLineno(380)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
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
							// line 381: if self.content_offset > self.lineno or not self.content:
							πF.SetLineno(381)
						Label2:
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" directive requires arguments on the first line.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 382: raise self.error('"%s" directive requires arguments on the first '
							πF.SetLineno(382)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 384: args = self.content[0]
							πF.SetLineno(384)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp003
							// line 385: match = self.argument_pattern.match(args)
							πF.SetLineno(385)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp007[0] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßargument_pattern, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmatch = πTemp001
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µmatch); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 386: if not match:
							πF.SetLineno(386)
						Label4:
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, µargs).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" directive arguments not valid role names: \"%s\".").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 387: raise self.error('"%s" directive arguments not valid role names: '
							πF.SetLineno(387)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 389: new_role_name = match.group(1)
							πF.SetLineno(389)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µnew_role_name = πTemp003
							// line 390: base_role_name = match.group(3)
							πF.SetLineno(390)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µbase_role_name = πTemp003
							// line 391: messages = []
							πF.SetLineno(391)
							πTemp007 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp007...).ToObject()
							µmessages = πTemp001
							if πE = πg.CheckLocal(πF, µbase_role_name, "base_role_name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µbase_role_name); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 392: if base_role_name:
							πF.SetLineno(392)
						Label6:
							// line 393: base_role, messages = roles.role(
							πF.SetLineno(393)
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µbase_role_name, "base_role_name"); πE != nil {
								continue
							}
							πTemp007[0] = µbase_role_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlanguage, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp007[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							πTemp007[3] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrole, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µbase_role = πTemp003
							µmessages = πTemp004
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µbase_role == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 396: if base_role is None:
							πF.SetLineno(396)
						Label9:
							// line 397: error = self.state.reporter.error(
							πF.SetLineno(397)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbase_role_name, "base_role_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unknown interpreted text role \"%s\".").ToObject(), µbase_role_name); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µerror = πTemp003
							// line 401: return messages + [error]
							πF.SetLineno(401)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp007[0] = µerror
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label10
						Label10:
							goto Label8
						Label7:
							// line 403: base_role = roles.generic_custom_role
							πF.SetLineno(403)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßgeneric_custom_role, nil); πE != nil {
								continue
							}
							µbase_role = πTemp003
							goto Label8
						Label8:
							// line 404: assert not hasattr(base_role, 'arguments'), (
							πF.SetLineno(404)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, µbase_role).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Supplemental directive arguments for \"%s\" directive not supported (specified by \"%r\" role).").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							πTemp007[0] = µbase_role
							πTemp007[1] = ßarguments.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp002).ToObject()
							if πE = πg.Assert(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 407: try:
							πF.SetLineno(407)
							πF.PushCheckpoint(12)
							// line 408: converted_role = convert_directive_function(base_role)
							πF.SetLineno(408)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							πTemp007[0] = µbase_role
							if πTemp001, πE = πg.ResolveGlobal(πF, ßconvert_directive_function); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µconverted_role = πTemp003
							// line 409: (arguments, options, content, content_offset) = (
							πF.SetLineno(409)
							πTemp007 = πF.MakeArgs(3)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µconverted_role, "converted_role"); πE != nil {
								continue
							}
							πTemp007[2] = µconverted_role
							πTemp010 = πg.NewDict()
							πTemp001 = πTemp010.ToObject()
							πTemp009 = πg.KWArgs{
								{"option_presets", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßparse_directive_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp011}}}, πTemp001); πE != nil {
								continue
							}
							µarguments = πTemp003
							µoptions = πTemp004
							µcontent = πTemp005
							µcontent_offset = πTemp011
							πF.PopCheckpoint()
							goto Label11
						Label12:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp012, πTemp013 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßMarkupError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label13
							}
							πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
							continue
							// line 413: except states.MarkupError as detail:
							πF.SetLineno(413)
						Label13:
							µdetail = πTemp012.ToObject()
							// line 414: error = self.state_machine.reporter.error(
							πF.SetLineno(414)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, µdetail).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error in \"%s\" directive:\n%s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µerror = πTemp003
							// line 418: return messages + [error]
							πF.SetLineno(418)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp007[0] = µerror
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µoptions, ßclass.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label14
							}
							goto Label15
							// line 419: if 'class' not in options:
							πF.SetLineno(419)
						Label14:
							// line 420: try:
							πF.SetLineno(420)
							πF.PushCheckpoint(17)
							// line 421: options['class'] = directives.class_option(new_role_name)
							πF.SetLineno(421)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnew_role_name, "new_role_name"); πE != nil {
								continue
							}
							πTemp007[0] = µnew_role_name
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							πTemp004 = ßclass.ToObject()
							if πE = πg.SetItem(πF, µoptions, πTemp004, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label16
						Label17:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp012, πTemp013 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label18
							}
							πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
							continue
							// line 422: except ValueError as detail:
							πF.SetLineno(422)
						Label18:
							µdetail = πTemp012.ToObject()
							// line 423: error = self.state_machine.reporter.error(
							πF.SetLineno(423)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
								continue
							}
							πTemp008[0] = µdetail
							if πTemp005, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp003 = πg.NewTuple2(πTemp004, πTemp011).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Invalid argument for \"%s\" directive:\n%s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µerror = πTemp003
							// line 427: return messages + [error]
							πF.SetLineno(427)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp007[0] = µerror
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label16
						Label16:
							goto Label15
						Label15:
							// line 428: role = roles.CustomRole(new_role_name, base_role, options, content)
							πF.SetLineno(428)
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µnew_role_name, "new_role_name"); πE != nil {
								continue
							}
							πTemp007[0] = µnew_role_name
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							πTemp007[1] = µbase_role
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							πTemp007[2] = µoptions
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							πTemp007[3] = µcontent
							if πTemp001, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßCustomRole, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µrole = πTemp001
							// line 429: roles.register_local_role(new_role_name, role)
							πF.SetLineno(429)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnew_role_name, "new_role_name"); πE != nil {
								continue
							}
							πTemp007[0] = µnew_role_name
							if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
								continue
							}
							πTemp007[1] = µrole
							if πTemp001, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßregister_local_role, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 430: return messages
							πF.SetLineno(430)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πR = µmessages
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 380: """Dynamically create and register a custom interpreted text role."""
					πF.SetLineno(380)
					// line 380: """Dynamically create and register a custom interpreted text role."""
					πF.SetLineno(380)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Dynamically create and register a custom interpreted text role.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Dynamically create and register a custom interpreted text role.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßrun); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Role").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRole.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 433: class DefaultRole(Directive):
			πF.SetLineno(433)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DefaultRole", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 435: """Set the default interpreted text role."""
					πF.SetLineno(435)
					// line 435: """Set the default interpreted text role."""
					πF.SetLineno(435)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Set the default interpreted text role.").ToObject()); πE != nil {
						continue
					}
					// line 437: optional_arguments = 1
					πF.SetLineno(437)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 438: final_argument_whitespace = False
					πF.SetLineno(438)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 440: def run(self):
					πF.SetLineno(440)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrole_name *πg.Object = πg.UnboundLocal
						_ = µrole_name
						var µrole *πg.Object = πg.UnboundLocal
						_ = µrole
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
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
							// line 441: if not self.arguments:
							πF.SetLineno(441)
						Label1:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ß_roles, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp004, ß.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 442: if '' in roles._roles:
							πF.SetLineno(442)
						Label3:
							// line 444: del roles._roles['']
							πF.SetLineno(444)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_roles, nil); πE != nil {
								continue
							}
							πTemp001 = ß.ToObject()
							if πE = πg.DelItem(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 445: return []
							πF.SetLineno(445)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 446: role_name = self.arguments[0]
							πF.SetLineno(446)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µrole_name = πTemp002
							// line 447: role, messages = roles.role(role_name, self.state_machine.language,
							πF.SetLineno(447)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
								continue
							}
							πTemp005[0] = µrole_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlanguage, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							πTemp005[3] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrole, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µrole = πTemp002
							µmessages = πTemp004
							if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µrole == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 449: if role is None:
							πF.SetLineno(449)
						Label5:
							// line 450: error = self.state.reporter.error(
							πF.SetLineno(450)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unknown interpreted text role \"%s\".").ToObject(), µrole_name); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µerror = πTemp002
							// line 454: return messages + [error]
							πF.SetLineno(454)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp005[0] = µerror
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label6
						Label6:
							// line 455: roles._roles[''] = role
							πF.SetLineno(455)
							if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrole); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ß_roles, nil); πE != nil {
								continue
							}
							πTemp002 = ß.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 456: return messages
							πF.SetLineno(456)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πR = µmessages
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DefaultRole").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDefaultRole.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 459: class Title(Directive):
			πF.SetLineno(459)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Title", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 461: required_arguments = 1
					πF.SetLineno(461)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 462: optional_arguments = 0
					πF.SetLineno(462)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 463: final_argument_whitespace = True
					πF.SetLineno(463)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 465: def run(self):
					πF.SetLineno(465)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 466: self.state_machine.document['title'] = self.arguments[0]
							πF.SetLineno(466)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdocument, nil); πE != nil {
								continue
							}
							πTemp003 = ßtitle.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 467: return []
							πF.SetLineno(467)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Title").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTitle.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 470: class Date(Directive):
			πF.SetLineno(470)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Date", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 472: has_content = True
					πF.SetLineno(472)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 474: def run(self):
					πF.SetLineno(474)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µformat_str *πg.Object = πg.UnboundLocal
						_ = µformat_str
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 13:
								goto Label13
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSubstitutionDef, nil); πE != nil {
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
							// line 475: if not isinstance(self.state, states.SubstitutionDef):
							πF.SetLineno(475)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid context: the \"%s\" directive can only be used within a substitution definition.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 476: raise self.error(
							πF.SetLineno(476)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 479: format_str = '\n'.join(self.content) or '%Y-%m-%d'
							πF.SetLineno(479)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
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
							πTemp001 = πg.NewStr("%Y-%m-%d").ToObject()
						Label3:
							µformat_str = πTemp001
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
								goto Label4
							}
							goto Label5
							// line 480: if sys.version_info< (3, 0):
							πF.SetLineno(480)
						Label4:
							// line 481: try:
							πF.SetLineno(481)
							πF.PushCheckpoint(7)
							// line 482: format_str = format_str.encode(locale_encoding or 'utf-8')
							πF.SetLineno(482)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							πTemp001 = πg.NewStr("utf-8").ToObject()
						Label8:
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µformat_str, "format_str"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µformat_str, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µformat_str = πTemp003
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 483: except UnicodeEncodeError:
							πF.SetLineno(483)
						Label9:
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Cannot encode date format string with locale encoding \"%s\".").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 484: raise self.warning(u'Cannot encode date format string '
							πF.SetLineno(484)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							goto Label5
						Label5:
							// line 502: text = time.strftime(format_str)
							πF.SetLineno(502)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformat_str, "format_str"); πE != nil {
								continue
							}
							πTemp002[0] = µformat_str
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßstrftime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µtext = πTemp001
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
								goto Label10
							}
							goto Label11
							// line 503: if sys.version_info< (3, 0):
							πF.SetLineno(503)
						Label10:
							// line 505: try:
							πF.SetLineno(505)
							πF.PushCheckpoint(13)
							// line 506: text = text.decode(locale_encoding or 'utf-8')
							πF.SetLineno(506)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							πTemp001 = πg.NewStr("utf-8").ToObject()
						Label14:
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µtext = πTemp003
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 507: except UnicodeDecodeError:
							πF.SetLineno(507)
						Label15:
							// line 508: text = text.decode(locale_encoding or 'utf-8', 'replace')
							πF.SetLineno(508)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label16
							}
							πTemp001 = πg.NewStr("utf-8").ToObject()
						Label16:
							πTemp002[0] = πTemp001
							πTemp002[1] = ßreplace.ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µtext = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µtext, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("Error decoding \"%s\"with locale encoding \"%s\".").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 509: raise self.warning(u'Error decoding "%s"'
							πF.SetLineno(509)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							goto Label11
						Label11:
							// line 511: return [nodes.Text(text)]
							πF.SetLineno(511)
							πTemp002 = make([]*πg.Object, 1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp008[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewList(πTemp002...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Date").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDate.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 514: class TestDirective(Directive):
			πF.SetLineno(514)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestDirective", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 516: """This directive is useful only for testing purposes."""
					πF.SetLineno(516)
					// line 516: """This directive is useful only for testing purposes."""
					πF.SetLineno(516)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This directive is useful only for testing purposes.").ToObject()); πE != nil {
						continue
					}
					// line 518: optional_arguments = 1
					πF.SetLineno(518)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 519: final_argument_whitespace = True
					πF.SetLineno(519)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 520: option_spec = {'option': directives.unchanged_required}
					πF.SetLineno(520)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßoption.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 521: has_content = True
					πF.SetLineno(521)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 523: def run(self):
					πF.SetLineno(523)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/misc.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µinfo *πg.Object = πg.UnboundLocal
						_ = µinfo
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 524: if self.content:
							πF.SetLineno(524)
						Label1:
							// line 525: text = '\n'.join(self.content)
							πF.SetLineno(525)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp004
							// line 526: info = self.state_machine.reporter.info(
							πF.SetLineno(526)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Directive processed. Type=\"%s\", arguments=%r, options=%r, content:").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp008[0] = µtext
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp008[1] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µinfo = πTemp004
							goto Label3
						Label2:
							// line 531: info = self.state_machine.reporter.info(
							πF.SetLineno(531)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Directive processed. Type=\"%s\", arguments=%r, options=%r, content: None").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µinfo = πTemp004
							goto Label3
						Label3:
							// line 535: return [info]
							πF.SetLineno(535)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
								continue
							}
							πTemp003[0] = µinfo
							πTemp001 = πg.NewList(πTemp003...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestDirective").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestDirective.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("misc", Code)
}
