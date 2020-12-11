package tables

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/csv"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/urllib"
	_ "github.com/pygolin/stdlib/pkg/urllib/error"
	_ "github.com/pygolin/stdlib/pkg/urllib/request"
	_ "github.com/pygolin/stdlib/pkg/urllib2"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßCSVTable := πg.InternStr("CSVTable")
		ßDialect := πg.InternStr("Dialect")
		ßDirective := πg.InternStr("Directive")
		ßDocutilsDialect := πg.InternStr("DocutilsDialect")
		ßElement := πg.InternStr("Element")
		ßError := πg.InternStr("Error")
		ßFalse := πg.InternStr("False")
		ßFileInput := πg.InternStr("FileInput")
		ßHeaderDialect := πg.InternStr("HeaderDialect")
		ßIOError := πg.InternStr("IOError")
		ßListTable := πg.InternStr("ListTable")
		ßNone := πg.InternStr("None")
		ßOSError := πg.InternStr("OSError")
		ßQUOTE_MINIMAL := πg.InternStr("QUOTE_MINIMAL")
		ßRSTTable := πg.InternStr("RSTTable")
		ßSafeString := πg.InternStr("SafeString")
		ßStringInput := πg.InternStr("StringInput")
		ßStringList := πg.InternStr("StringList")
		ßSystemMessagePropagation := πg.InternStr("SystemMessagePropagation")
		ßTable := πg.InternStr("Table")
		ßTrue := πg.InternStr("True")
		ßURLError := πg.InternStr("URLError")
		ßValueError := πg.InternStr("ValueError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßabspath := πg.InternStr("abspath")
		ßadd := πg.InternStr("add")
		ßadd_name := πg.InternStr("add_name")
		ßalign := πg.InternStr("align")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßarguments := πg.InternStr("arguments")
		ßattributes := πg.InternStr("attributes")
		ßauto := πg.InternStr("auto")
		ßblock_text := πg.InternStr("block_text")
		ßbuild_table := πg.InternStr("build_table")
		ßbuild_table_from_list := πg.InternStr("build_table_from_list")
		ßbullet_list := πg.InternStr("bullet_list")
		ßcenter := πg.InternStr("center")
		ßcheck_list_content := πg.InternStr("check_list_content")
		ßcheck_requirements := πg.InternStr("check_requirements")
		ßcheck_table_dimensions := πg.InternStr("check_table_dimensions")
		ßchildren := πg.InternStr("children")
		ßchoice := πg.InternStr("choice")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcolspec := πg.InternStr("colspec")
		ßcolwidth := πg.InternStr("colwidth")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßcsv := πg.InternStr("csv")
		ßcurrent_source := πg.InternStr("current_source")
		ßdecode := πg.InternStr("decode")
		ßdecode_from_csv := πg.InternStr("decode_from_csv")
		ßdelim := πg.InternStr("delim")
		ßdelimiter := πg.InternStr("delimiter")
		ßdirectives := πg.InternStr("directives")
		ßdirname := πg.InternStr("dirname")
		ßdocument := πg.InternStr("document")
		ßdoublequote := πg.InternStr("doublequote")
		ßencode := πg.InternStr("encode")
		ßencode_for_csv := πg.InternStr("encode_for_csv")
		ßencoding := πg.InternStr("encoding")
		ßentry := πg.InternStr("entry")
		ßerror := πg.InternStr("error")
		ßescape := πg.InternStr("escape")
		ßescapechar := πg.InternStr("escapechar")
		ßextend := πg.InternStr("extend")
		ßextend_short_rows_with_empty_cells := πg.InternStr("extend_short_rows_with_empty_cells")
		ßfile := πg.InternStr("file")
		ßfile_insertion_enabled := πg.InternStr("file_insertion_enabled")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßflag := πg.InternStr("flag")
		ßget := πg.InternStr("get")
		ßget_column_widths := πg.InternStr("get_column_widths")
		ßget_csv_data := πg.InternStr("get_csv_data")
		ßget_source := πg.InternStr("get_source")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßgrid := πg.InternStr("grid")
		ßhas_content := πg.InternStr("has_content")
		ßheader := πg.InternStr("header")
		ßinline_text := πg.InternStr("inline_text")
		ßinput_encoding := πg.InternStr("input_encoding")
		ßinput_encoding_error_handler := πg.InternStr("input_encoding_error_handler")
		ßinsert := πg.InternStr("insert")
		ßio := πg.InternStr("io")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßkeepspace := πg.InternStr("keepspace")
		ßleft := πg.InternStr("left")
		ßlen := πg.InternStr("len")
		ßlength_or_percentage_or_unitless := πg.InternStr("length_or_percentage_or_unitless")
		ßline := πg.InternStr("line")
		ßlineno := πg.InternStr("lineno")
		ßlineterminator := πg.InternStr("lineterminator")
		ßlist := πg.InternStr("list")
		ßliteral_block := πg.InternStr("literal_block")
		ßmake_title := πg.InternStr("make_title")
		ßmax := πg.InternStr("max")
		ßname := πg.InternStr("name")
		ßnested_parse := πg.InternStr("nested_parse")
		ßnodes := πg.InternStr("nodes")
		ßnonnegative_int := πg.InternStr("nonnegative_int")
		ßnormpath := πg.InternStr("normpath")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptional_arguments := πg.InternStr("optional_arguments")
		ßoptions := πg.InternStr("options")
		ßos := πg.InternStr("os")
		ßparse_csv_data_into_rows := πg.InternStr("parse_csv_data_into_rows")
		ßpath := πg.InternStr("path")
		ßpositive_int_list := πg.InternStr("positive_int_list")
		ßprocess_header_option := πg.InternStr("process_header_option")
		ßproperty := πg.InternStr("property")
		ßquote := πg.InternStr("quote")
		ßquotechar := πg.InternStr("quotechar")
		ßquoting := πg.InternStr("quoting")
		ßrange := πg.InternStr("range")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßreader := πg.InternStr("reader")
		ßrecord_dependencies := πg.InternStr("record_dependencies")
		ßrelative_path := πg.InternStr("relative_path")
		ßreporter := πg.InternStr("reporter")
		ßright := πg.InternStr("right")
		ßrow := πg.InternStr("row")
		ßrun := πg.InternStr("run")
		ßset_table_width := πg.InternStr("set_table_width")
		ßsettings := πg.InternStr("settings")
		ßsevere := πg.InternStr("severe")
		ßsingle_char_or_unicode := πg.InternStr("single_char_or_unicode")
		ßsingle_char_or_whitespace_or_unicode := πg.InternStr("single_char_or_whitespace_or_unicode")
		ßskipinitialspace := πg.InternStr("skipinitialspace")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßsplitlines := πg.InternStr("splitlines")
		ßstate := πg.InternStr("state")
		ßstate_machine := πg.InternStr("state_machine")
		ßstatemachine := πg.InternStr("statemachine")
		ßstaticmethod := πg.InternStr("staticmethod")
		ßstr := πg.InternStr("str")
		ßstrict := πg.InternStr("strict")
		ßstub := πg.InternStr("stub")
		ßsys := πg.InternStr("sys")
		ßtable := πg.InternStr("table")
		ßtagname := πg.InternStr("tagname")
		ßtbody := πg.InternStr("tbody")
		ßtgroup := πg.InternStr("tgroup")
		ßthead := πg.InternStr("thead")
		ßtitle := πg.InternStr("title")
		ßunchanged := πg.InternStr("unchanged")
		ßuri := πg.InternStr("uri")
		ßurl := πg.InternStr("url")
		ßurlopen := πg.InternStr("urlopen")
		ßutils := πg.InternStr("utils")
		ßvalue_or := πg.InternStr("value_or")
		ßversion_info := πg.InternStr("version_info")
		ßwarning := πg.InternStr("warning")
		ßwidth := πg.InternStr("width")
		ßwidths := πg.InternStr("widths")
		ßzip := πg.InternStr("zip")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Dict
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
			// line 5: """
			πF.SetLineno(5)
			// line 5: """
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDirectives for table elements.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 12: import sys
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: import os.path
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import csv
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "csv"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcsv.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: from docutils import io, nodes, statemachine, utils
			πF.SetLineno(16)
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
			// line 17: from docutils.utils.error_reporting import SafeString
			πF.SetLineno(17)
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
			// line 18: from docutils.utils import SystemMessagePropagation
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSystemMessagePropagation); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSystemMessagePropagation.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from docutils.parsers.rst import Directive
			πF.SetLineno(19)
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
			// line 20: from docutils.parsers.rst import directives
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: def align(argument):
			πF.SetLineno(23)
			πTemp004 = make([]πg.Param, 1)
			πTemp004[0] = πg.Param{Name: "argument", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("align", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargument *πg.Object = πArgs[0]
				_ = µargument
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
					// line 24: return directives.choice(argument, ('left', 'center', 'right'))
					πF.SetLineno(24)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
						continue
					}
					πTemp001[0] = µargument
					πTemp002 = πg.NewTuple3(ßleft.ToObject(), ßcenter.ToObject(), ßright.ToObject()).ToObject()
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchoice, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßalign.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: class Table(Directive):
			πF.SetLineno(27)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Table", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 29: """
					πF.SetLineno(29)
					// line 29: """
					πF.SetLineno(29)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Generic table base class.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 33: optional_arguments = 1
					πF.SetLineno(33)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 34: final_argument_whitespace = True
					πF.SetLineno(34)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 35: option_spec = {'class': directives.class_option,
					πF.SetLineno(35)
					πTemp002 = πg.NewDict()
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
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßalign); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßalign.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlength_or_percentage_or_unitless, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßwidth.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(2)
					πTemp001 = πg.NewTuple2(ßauto.ToObject(), ßgrid.ToObject()).ToObject()
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpositive_int_list, nil); πE != nil {
						continue
					}
					πTemp004[1] = πTemp003
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßvalue_or, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πTemp002.SetItem(πF, ßwidths.ToObject(), πTemp001); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 41: has_content = True
					πF.SetLineno(41)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 43: def make_title(self):
					πF.SetLineno(43)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("make_title", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtitle_text *πg.Object = πg.UnboundLocal
						_ = µtitle_text
						var µtext_nodes *πg.Object = πg.UnboundLocal
						_ = µtext_nodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 44: if self.arguments:
							πF.SetLineno(44)
						Label1:
							// line 45: title_text = self.arguments[0]
							πF.SetLineno(45)
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
							µtitle_text = πTemp003
							// line 46: text_nodes, messages = self.state.inline_text(title_text,
							πF.SetLineno(46)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp005[0] = µtitle_text
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinline_text, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µtext_nodes = πTemp003
							µmessages = πTemp004
							// line 48: title = nodes.title(title_text, '', *text_nodes)
							πF.SetLineno(48)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp005[0] = µtitle_text
							πTemp005[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtext_nodes, "text_nodes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp003, πTemp005, µtext_nodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtitle = πTemp001
							// line 49: (title.source,
							πF.SetLineno(49)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_source_and_line, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtitle, ßsource, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtitle, ßline, πTemp004); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 52: title = None
							πF.SetLineno(52)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtitle = πTemp001
							// line 53: messages = []
							πF.SetLineno(53)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µmessages = πTemp001
							goto Label3
						Label3:
							// line 54: return title, messages
							πF.SetLineno(54)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µtitle, µmessages).ToObject()
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
					if πE = πClass.SetItem(πF, ßmake_title.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 56: def process_header_option(self):
					πF.SetLineno(56)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("process_header_option", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µtable_head *πg.Object = πg.UnboundLocal
						_ = µtable_head
						var µmax_header_cols *πg.Object = πg.UnboundLocal
						_ = µmax_header_cols
						var µrows *πg.Object = πg.UnboundLocal
						_ = µrows
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
							// line 57: source = self.state_machine.get_source(self.lineno - 1)
							πF.SetLineno(57)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_source, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsource = πTemp002
							// line 58: table_head = []
							πF.SetLineno(58)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µtable_head = πTemp002
							// line 59: max_header_cols = 0
							πF.SetLineno(59)
							µmax_header_cols = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßheader.ToObject()); πE != nil {
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
							// line 60: if 'header' in self.options:   # separate table header in option
							πF.SetLineno(60)
						Label1:
							// line 61: rows, max_header_cols = self.parse_csv_data_into_rows(
							πF.SetLineno(61)
							πTemp001 = πF.MakeArgs(3)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							πTemp002 = ßheader.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßHeaderDialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[2] = µsource
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_csv_data_into_rows, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µrows = πTemp002
							µmax_header_cols = πTemp006
							// line 64: table_head.extend(rows)
							πF.SetLineno(64)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp001[0] = µrows
							if πE = πg.CheckLocal(πF, µtable_head, "table_head"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtable_head, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 65: return table_head, max_header_cols
							πF.SetLineno(65)
							if πE = πg.CheckLocal(πF, µtable_head, "table_head"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmax_header_cols, "max_header_cols"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µtable_head, µmax_header_cols).ToObject()
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
					if πE = πClass.SetItem(πF, ßprocess_header_option.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 67: def check_table_dimensions(self, rows, header_rows, stub_columns):
					πF.SetLineno(67)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "rows", Def: nil}
					πTemp005[2] = πg.Param{Name: "header_rows", Def: nil}
					πTemp005[3] = πg.Param{Name: "stub_columns", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("check_table_dimensions", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrows *πg.Object = πArgs[1]
						_ = µrows
						var µheader_rows *πg.Object = πArgs[2]
						_ = µheader_rows
						var µstub_columns *πg.Object = πArgs[3]
						_ = µstub_columns
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µrow *πg.Object = πg.UnboundLocal
						_ = µrow
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp002[0] = µrows
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp004, µheader_rows); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 68: if len(rows) < header_rows:
							πF.SetLineno(68)
						Label1:
							// line 69: error = self.state_machine.reporter.error(
							πF.SetLineno(69)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp006[0] = µrows
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(µheader_rows, πTemp007, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s header row(s) specified but only %s row(s) of data supplied (\"%s\" directive).").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
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
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µerror = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp002[0] = µerror
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 74: raise SystemMessagePropagation(error)
							πF.SetLineno(74)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp002[0] = µrows
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, µheader_rows); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							if πTemp001, πE = πg.GT(πF, µheader_rows, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 75: if len(rows) == header_rows > 0:
							πF.SetLineno(75)
						Label4:
							// line 76: error = self.state_machine.reporter.error(
							πF.SetLineno(76)
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp006[0] = µrows
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Insufficient data supplied (%s row(s)); no data remaining for table body, required by \"%s\" directive.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
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
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µerror = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp002[0] = µerror
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 81: raise SystemMessagePropagation(error)
							πF.SetLineno(81)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µrows); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µrow = πTemp003
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp002[0] = µrow
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp007, µstub_columns); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 83: if len(row) < stub_columns:
							πF.SetLineno(83)
						Label9:
							// line 84: error = self.state_machine.reporter.error(
							πF.SetLineno(84)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp006[0] = µrow
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(µstub_columns, πTemp010, πTemp007).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s stub column(s) specified but only %s columns(s) of data supplied (\"%s\" directive).").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
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
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µerror = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp002[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 89: raise SystemMessagePropagation(error)
							πF.SetLineno(89)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label10
						Label10:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp002[0] = µrow
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp007, µstub_columns); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label11
							}
							if πTemp003, πE = πg.GT(πF, µstub_columns, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
						Label11:
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label12
							}
							goto Label13
							// line 90: if len(row) == stub_columns > 0:
							πF.SetLineno(90)
						Label12:
							// line 91: error = self.state_machine.reporter.error(
							πF.SetLineno(91)
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp006[0] = µrow
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πTemp010, πTemp007).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Insufficient data supplied (%s columns(s)); no data remaining for table body, required by \"%s\" directive.").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
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
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µerror = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp002[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 96: raise SystemMessagePropagation(error)
							πF.SetLineno(96)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label13
						Label13:
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
					if πE = πClass.SetItem(πF, ßcheck_table_dimensions.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 98: def set_table_width(self, table_node):
					πF.SetLineno(98)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "table_node", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("set_table_width", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtable_node *πg.Object = πArgs[1]
						_ = µtable_node
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
						var πR *πg.Object
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßwidth.ToObject()); πE != nil {
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
							// line 99: if 'width' in self.options:
							πF.SetLineno(99)
						Label1:
							// line 100: table_node['width'] = self.options.get('width')
							πF.SetLineno(100)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßwidth.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp005 = ßwidth.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp005, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_table_width.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 103: def widths(self):
					πF.SetLineno(103)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("widths", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 104: return self.options.get('widths', '')
							πF.SetLineno(104)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßwidths.ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwidths.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 102: @property
					πF.SetLineno(102)
					πTemp004 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßwidths); πE != nil {
						continue
					}
					πTemp004[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßwidths.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 106: def get_column_widths(self, max_cols):
					πF.SetLineno(106)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "max_cols", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("get_column_widths", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmax_cols *πg.Object = πArgs[1]
						_ = µmax_cols
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µcol_widths *πg.Object = πg.UnboundLocal
						_ = µcol_widths
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmax_cols); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 107: if isinstance(self.widths, list):
							πF.SetLineno(107)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
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
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp005, µmax_cols); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 108: if len(self.widths) != max_cols:
							πF.SetLineno(108)
						Label5:
							// line 109: error = self.state_machine.reporter.error(
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, µmax_cols).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\"%s\" widths do not match the number of columns in table (%s).").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 113: raise SystemMessagePropagation(error)
							πF.SetLineno(113)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label6
						Label6:
							// line 114: col_widths = self.widths
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							µcol_widths = πTemp002
							goto Label4
							// line 115: elif max_cols:
							πF.SetLineno(115)
						Label2:
							// line 116: col_widths = [100 // max_cols] * max_cols
							πF.SetLineno(116)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, πg.NewInt(100).ToObject(), µmax_cols); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp003, µmax_cols); πE != nil {
								continue
							}
							µcol_widths = πTemp002
							goto Label4
						Label3:
							// line 118: error = self.state_machine.reporter.error(
							πF.SetLineno(118)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("No table data detected in CSV file.").ToObject()
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 121: raise SystemMessagePropagation(error)
							πF.SetLineno(121)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 122: return col_widths
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µcol_widths, "col_widths"); πE != nil {
								continue
							}
							πR = µcol_widths
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_column_widths.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 124: def extend_short_rows_with_empty_cells(self, columns, parts):
					πF.SetLineno(124)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "columns", Def: nil}
					πTemp005[2] = πg.Param{Name: "parts", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("extend_short_rows_with_empty_cells", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcolumns *πg.Object = πArgs[1]
						_ = µcolumns
						var µparts *πg.Object = πArgs[2]
						_ = µparts
						var µpart *πg.Object = πg.UnboundLocal
						_ = µpart
						var µrow *πg.Object = πg.UnboundLocal
						_ = µrow
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µpart, "part"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µpart); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp003 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µrow = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp007[0] = µrow
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LT(πF, πTemp009, µcolumns); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 127: if len(row) < columns:
							πF.SetLineno(127)
						Label7:
							// line 128: row.extend([(0, 0, 0, [])] * (columns - len(row)))
							πF.SetLineno(128)
							πTemp007 = πF.MakeArgs(1)
							πTemp010 = make([]*πg.Object, 1)
							πTemp011 = make([]*πg.Object, 0)
							πTemp009 = πg.NewList(πTemp011...).ToObject()
							πTemp008 = πg.NewTuple4(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πTemp009).ToObject()
							πTemp010[0] = πTemp008
							πTemp008 = πg.NewList(πTemp010...).ToObject()
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp010[0] = µrow
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp009, πE = πg.Sub(πF, µcolumns, πTemp013); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πTemp007[0] = πTemp006
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µrow, ßextend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label8
						Label8:
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßextend_short_rows_with_empty_cells.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Table").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 131: class RSTTable(Table):
			πF.SetLineno(131)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßTable); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("RSTTable", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
					// line 133: def run(self):
					πF.SetLineno(133)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µwarning *πg.Object = πg.UnboundLocal
						_ = µwarning
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µtable_node *πg.Object = πg.UnboundLocal
						_ = µtable_node
						var µtgroup *πg.Object = πg.UnboundLocal
						_ = µtgroup
						var µcolspecs *πg.Object = πg.UnboundLocal
						_ = µcolspecs
						var µcolspec *πg.Object = πg.UnboundLocal
						_ = µcolspec
						var µcol_width *πg.Object = πg.UnboundLocal
						_ = µcol_width
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []πg.Param
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
							case 10:
								goto Label10
							case 11:
								goto Label11
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
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
							// line 134: if not self.content:
							πF.SetLineno(134)
						Label1:
							// line 135: warning = self.state_machine.reporter.warning(
							πF.SetLineno(135)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Content block expected for the \"%s\" directive; none found.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µwarning = πTemp002
							// line 139: return [warning]
							πF.SetLineno(139)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µwarning, "warning"); πE != nil {
								continue
							}
							πTemp004[0] = µwarning
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 140: title, messages = self.make_title()
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmake_title, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µtitle = πTemp001
							µmessages = πTemp007
							// line 141: node = nodes.Element()          # anonymous container for parsing
							πF.SetLineno(141)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßElement, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnode = πTemp001
							// line 142: self.state.nested_parse(self.content, self.content_offset, node)
							πF.SetLineno(142)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[2] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.NE(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp007 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp007); πE != nil {
								continue
							}
							πTemp004[0] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßtable, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp009).ToObject()
							πTemp001 = πTemp002
						Label3:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 143: if len(node) != 1 or not isinstance(node[0], nodes.table):
							πF.SetLineno(143)
						Label4:
							// line 144: error = self.state_machine.reporter.error(
							πF.SetLineno(144)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error parsing content block for the \"%s\" directive: exactly one table expected.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerror = πTemp002
							// line 148: return [error]
							πF.SetLineno(148)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp004[0] = µerror
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πR = πTemp001
							continue
							goto Label5
						Label5:
							// line 149: table_node = node[0]
							πF.SetLineno(149)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							µtable_node = πTemp002
							// line 150: table_node['classes'] += self.options.get('class', [])
							πF.SetLineno(150)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtable_node, πTemp001); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßclass.ToObject()
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp008 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp008, πTemp007); πE != nil {
								continue
							}
							// line 151: self.set_table_width(table_node)
							πF.SetLineno(151)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_table_width, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßalign.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 152: if 'align' in self.options:
							πF.SetLineno(152)
						Label6:
							// line 153: table_node['align'] = self.options.get('align')
							πF.SetLineno(153)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp007 = ßalign.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 154: tgroup = table_node[0]
							πF.SetLineno(154)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtable_node, πTemp001); πE != nil {
								continue
							}
							µtgroup = πTemp002
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 155: if isinstance(self.widths, list):
							πF.SetLineno(155)
						Label8:
							// line 156: colspecs = [child for child in tgroup.children
							πF.SetLineno(156)
							πTemp010 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										case 6:
											goto Label6
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µtgroup, "tgroup"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µtgroup, ßchildren, nil); πE != nil {
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
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetAttr(πF, µchild, ßtagname, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πg.Eq(πF, πTemp005, ßcolspec.ToObject()); πE != nil {
											continue
										}
										if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
											continue
										}
										if πTemp004 {
											goto Label4
										}
										goto Label5
										// line 156: colspecs = [child for child in tgroup.children
										πF.SetLineno(156)
									Label4:
										// line 156: colspecs = [child for child in tgroup.children
										πF.SetLineno(156)
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
							if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							µcolspecs = πTemp001
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcolspecs, "colspecs"); πE != nil {
								continue
							}
							πTemp004[0] = µcolspecs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp003 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp007, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp011}}}, πTemp007); πE != nil {
									continue
								}
								µcolspec = πTemp008
								µcol_width = πTemp011
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(10)
							// line 159: colspec['colwidth'] = col_width
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µcol_width, "col_width"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µcol_width); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolspec, "colspec"); πE != nil {
								continue
							}
							πTemp008 = ßcolwidth.ToObject()
							if πE = πg.SetItem(πF, µcolspec, πTemp008, πTemp007); πE != nil {
								continue
							}
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp007, ßauto.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label13
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							goto Label15
							// line 163: if self.widths == 'auto':
							πF.SetLineno(163)
						Label13:
							// line 164: table_node['classes'] += ['colwidths-auto']
							πF.SetLineno(164)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtable_node, πTemp001); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("colwidths-auto").ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πTemp008, πE = πg.IAdd(πF, πTemp007, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp011 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp011, πTemp008); πE != nil {
								continue
							}
							goto Label15
							// line 165: elif self.widths: # "grid" or list of integers
							πF.SetLineno(165)
						Label14:
							// line 166: table_node['classes'] += ['colwidths-given']
							πF.SetLineno(166)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtable_node, πTemp001); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("colwidths-given").ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πTemp008, πE = πg.IAdd(πF, πTemp007, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp011 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp011, πTemp008); πE != nil {
								continue
							}
							goto Label15
						Label15:
							// line 167: self.add_name(table_node)
							πF.SetLineno(167)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µtitle); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							goto Label17
							// line 168: if title:
							πF.SetLineno(168)
						Label16:
							// line 169: table_node.insert(0, title)
							πF.SetLineno(169)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp004[1] = µtitle
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtable_node, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label17
						Label17:
							// line 170: return [table_node] + messages
							πF.SetLineno(170)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_node
							πTemp007 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp007, µmessages); πE != nil {
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
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("RSTTable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRSTTable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 173: class CSVTable(Table):
			πF.SetLineno(173)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßTable); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("CSVTable", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				var πTemp007 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 175: option_spec = {'header-rows': directives.nonnegative_int,
					πF.SetLineno(175)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("header-rows").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("stub-columns").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßheader.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlength_or_percentage_or_unitless, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßwidth.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(2)
					πTemp002 = πg.NewTuple1(ßauto.ToObject()).ToObject()
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpositive_int_list, nil); πE != nil {
						continue
					}
					πTemp004[1] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßvalue_or, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πTemp001.SetItem(πF, ßwidths.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßfile.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßuri, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßurl.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßencoding, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßencoding.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßalign); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßalign.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsingle_char_or_whitespace_or_unicode, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdelim.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßflag, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßkeepspace.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsingle_char_or_unicode, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßquote.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsingle_char_or_unicode, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßescape.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 196: class DocutilsDialect(csv.Dialect):
					πF.SetLineno(196)
					πTemp004 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßcsv); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßDialect, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("DocutilsDialect", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 198: """CSV dialect for `csv_table` directive."""
							πF.SetLineno(198)
							// line 198: """CSV dialect for `csv_table` directive."""
							πF.SetLineno(198)
							if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV dialect for `csv_table` directive.").ToObject()); πE != nil {
								continue
							}
							// line 200: delimiter = ','
							πF.SetLineno(200)
							if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr(",").ToObject()); πE != nil {
								continue
							}
							// line 201: quotechar = '"'
							πF.SetLineno(201)
							if πE = πClass.SetItem(πF, ßquotechar.ToObject(), πg.NewStr("\"").ToObject()); πE != nil {
								continue
							}
							// line 202: doublequote = True
							πF.SetLineno(202)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßdoublequote.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 203: skipinitialspace = True
							πF.SetLineno(203)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßskipinitialspace.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 204: strict = True
							πF.SetLineno(204)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßstrict.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 205: lineterminator = '\n'
							πF.SetLineno(205)
							if πE = πClass.SetItem(πF, ßlineterminator.ToObject(), πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							// line 206: quoting = csv.QUOTE_MINIMAL
							πF.SetLineno(206)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßcsv); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßQUOTE_MINIMAL, nil); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßquoting.ToObject(), πTemp002); πE != nil {
								continue
							}
							// line 208: def __init__(self, options):
							πF.SetLineno(208)
							πTemp003 = make([]πg.Param, 2)
							πTemp003[0] = πg.Param{Name: "self", Def: nil}
							πTemp003[1] = πg.Param{Name: "options", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]
								_ = µself
								var µoptions *πg.Object = πArgs[1]
								_ = µoptions
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
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Contains(πF, µoptions, ßdelim.ToObject()); πE != nil {
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
									// line 209: if 'delim' in options:
									πF.SetLineno(209)
								Label1:
									// line 210: self.delimiter = CSVTable.encode_for_csv(options['delim'])
									πF.SetLineno(210)
									πTemp003 = πF.MakeArgs(1)
									πTemp001 = ßdelim.ToObject()
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetItem(πF, µoptions, πTemp001); πE != nil {
										continue
									}
									πTemp003[0] = πTemp004
									if πTemp001, πE = πg.ResolveGlobal(πF, ßCSVTable); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßencode_for_csv, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßdelimiter, πTemp004); πE != nil {
										continue
									}
									goto Label2
								Label2:
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Contains(πF, µoptions, ßkeepspace.ToObject()); πE != nil {
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
									// line 211: if 'keepspace' in options:
									πF.SetLineno(211)
								Label3:
									// line 212: self.skipinitialspace = False
									πF.SetLineno(212)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßskipinitialspace, πTemp004); πE != nil {
										continue
									}
									goto Label4
								Label4:
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Contains(πF, µoptions, ßquote.ToObject()); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp002).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label5
									}
									goto Label6
									// line 213: if 'quote' in options:
									πF.SetLineno(213)
								Label5:
									// line 214: self.quotechar = CSVTable.encode_for_csv(options['quote'])
									πF.SetLineno(214)
									πTemp003 = πF.MakeArgs(1)
									πTemp001 = ßquote.ToObject()
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetItem(πF, µoptions, πTemp001); πE != nil {
										continue
									}
									πTemp003[0] = πTemp004
									if πTemp001, πE = πg.ResolveGlobal(πF, ßCSVTable); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßencode_for_csv, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßquotechar, πTemp004); πE != nil {
										continue
									}
									goto Label6
								Label6:
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Contains(πF, µoptions, ßescape.ToObject()); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp002).ToObject()
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label7
									}
									goto Label8
									// line 215: if 'escape' in options:
									πF.SetLineno(215)
								Label7:
									// line 216: self.doublequote = False
									πF.SetLineno(216)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßdoublequote, πTemp004); πE != nil {
										continue
									}
									// line 217: self.escapechar = CSVTable.encode_for_csv(options['escape'])
									πF.SetLineno(217)
									πTemp003 = πF.MakeArgs(1)
									πTemp001 = ßescape.ToObject()
									if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetItem(πF, µoptions, πTemp001); πE != nil {
										continue
									}
									πTemp003[0] = πTemp004
									if πTemp001, πE = πg.ResolveGlobal(πF, ßCSVTable); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßencode_for_csv, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßescapechar, πTemp004); πE != nil {
										continue
									}
									goto Label8
								Label8:
									// line 218: csv.Dialect.__init__(self)
									πF.SetLineno(218)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp003[0] = µself
									if πTemp001, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßDialect, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, πTemp004, ß__init__, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
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
							if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp003 == nil {
						πTemp003 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DocutilsDialect").ToObject(), πg.NewTuple(πTemp004...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßDocutilsDialect.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 221: class HeaderDialect(csv.Dialect):
					πF.SetLineno(221)
					πTemp004 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßcsv); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßDialect, nil); πE != nil {
						continue
					}
					πTemp004[0] = πTemp006
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("HeaderDialect", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 223: """CSV dialect to use for the "header" option data."""
							πF.SetLineno(223)
							// line 223: """CSV dialect to use for the "header" option data."""
							πF.SetLineno(223)
							if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV dialect to use for the \"header\" option data.").ToObject()); πE != nil {
								continue
							}
							// line 225: delimiter = ','
							πF.SetLineno(225)
							if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr(",").ToObject()); πE != nil {
								continue
							}
							// line 226: quotechar = '"'
							πF.SetLineno(226)
							if πE = πClass.SetItem(πF, ßquotechar.ToObject(), πg.NewStr("\"").ToObject()); πE != nil {
								continue
							}
							// line 227: escapechar = '\\'
							πF.SetLineno(227)
							if πE = πClass.SetItem(πF, ßescapechar.ToObject(), πg.NewStr("\\").ToObject()); πE != nil {
								continue
							}
							// line 228: doublequote = False
							πF.SetLineno(228)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßdoublequote.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 229: skipinitialspace = True
							πF.SetLineno(229)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßskipinitialspace.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 230: strict = True
							πF.SetLineno(230)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßstrict.ToObject(), πTemp001); πE != nil {
								continue
							}
							// line 231: lineterminator = '\n'
							πF.SetLineno(231)
							if πE = πClass.SetItem(πF, ßlineterminator.ToObject(), πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							// line 232: quoting = csv.QUOTE_MINIMAL
							πF.SetLineno(232)
							if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßcsv); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßQUOTE_MINIMAL, nil); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßquoting.ToObject(), πTemp002); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp003 == nil {
						πTemp003 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("HeaderDialect").ToObject(), πg.NewTuple(πTemp004...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßHeaderDialect.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 234: def check_requirements(self):
					πF.SetLineno(234)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("check_requirements", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 235: pass
							πF.SetLineno(235)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_requirements.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 237: def run(self):
					πF.SetLineno(237)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µwarning *πg.Object = πg.UnboundLocal
						_ = µwarning
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µcsv_data *πg.Object = πg.UnboundLocal
						_ = µcsv_data
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µtable_head *πg.Object = πg.UnboundLocal
						_ = µtable_head
						var µmax_header_cols *πg.Object = πg.UnboundLocal
						_ = µmax_header_cols
						var µrows *πg.Object = πg.UnboundLocal
						_ = µrows
						var µmax_cols *πg.Object = πg.UnboundLocal
						_ = µmax_cols
						var µheader_rows *πg.Object = πg.UnboundLocal
						_ = µheader_rows
						var µstub_columns *πg.Object = πg.UnboundLocal
						_ = µstub_columns
						var µtable_body *πg.Object = πg.UnboundLocal
						_ = µtable_body
						var µcol_widths *πg.Object = πg.UnboundLocal
						_ = µcol_widths
						var µdetail *πg.Object = πg.UnboundLocal
						_ = µdetail
						var µmessage *πg.Object = πg.UnboundLocal
						_ = µmessage
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µtable *πg.Object = πg.UnboundLocal
						_ = µtable
						var µtable_node *πg.Object = πg.UnboundLocal
						_ = µtable_node
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
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
							// line 238: try:
							πF.SetLineno(238)
							πF.PushCheckpoint(2)
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
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfile_insertion_enabled, nil); πE != nil {
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
							if !πTemp002 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp005, ßfile.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp007).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp005, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp007).ToObject()
							πTemp003 = πTemp004
						Label4:
							πTemp001 = πTemp003
						Label3:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 239: if (not self.state.document.settings.file_insertion_enabled
							πF.SetLineno(239)
						Label5:
							// line 242: warning = self.state_machine.reporter.warning(
							πF.SetLineno(242)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("File and URL access deactivated; ignoring \"%s\" directive.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
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
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µwarning = πTemp003
							// line 246: return [warning]
							πF.SetLineno(246)
							πTemp008 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µwarning, "warning"); πE != nil {
								continue
							}
							πTemp008[0] = µwarning
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πR = πTemp001
							continue
							goto Label6
						Label6:
							// line 247: self.check_requirements()
							πF.SetLineno(247)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_requirements, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 248: title, messages = self.make_title()
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmake_title, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µtitle = πTemp001
							µmessages = πTemp004
							// line 249: csv_data, source = self.get_csv_data()
							πF.SetLineno(249)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_csv_data, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µcsv_data = πTemp001
							µsource = πTemp004
							// line 250: table_head, max_header_cols = self.process_header_option()
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßprocess_header_option, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µtable_head = πTemp001
							µmax_header_cols = πTemp004
							// line 251: rows, max_cols = self.parse_csv_data_into_rows(
							πF.SetLineno(251)
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µcsv_data, "csv_data"); πE != nil {
								continue
							}
							πTemp008[0] = µcsv_data
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßDocutilsDialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp008[1] = πTemp003
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp008[2] = µsource
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparse_csv_data_into_rows, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µrows = πTemp001
							µmax_cols = πTemp004
							// line 253: max_cols = max(max_cols, max_header_cols)
							πF.SetLineno(253)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							πTemp008[0] = µmax_cols
							if πE = πg.CheckLocal(πF, µmax_header_cols, "max_header_cols"); πE != nil {
								continue
							}
							πTemp008[1] = µmax_header_cols
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µmax_cols = πTemp003
							// line 254: header_rows = self.options.get('header-rows', 0)
							πF.SetLineno(254)
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = πg.NewStr("header-rows").ToObject()
							πTemp008[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µheader_rows = πTemp001
							// line 255: stub_columns = self.options.get('stub-columns', 0)
							πF.SetLineno(255)
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = πg.NewStr("stub-columns").ToObject()
							πTemp008[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µstub_columns = πTemp001
							// line 256: self.check_table_dimensions(rows, header_rows, stub_columns)
							πF.SetLineno(256)
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp008[0] = µrows
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							πTemp008[1] = µheader_rows
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							πTemp008[2] = µstub_columns
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_table_dimensions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 257: table_head.extend(rows[:header_rows])
							πF.SetLineno(257)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µheader_rows, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µrows, πTemp001); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtable_head, "table_head"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtable_head, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 258: table_body = rows[header_rows:]
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µheader_rows, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µrows, πTemp001); πE != nil {
								continue
							}
							µtable_body = πTemp003
							// line 259: col_widths = self.get_column_widths(max_cols)
							πF.SetLineno(259)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							πTemp008[0] = µmax_cols
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_column_widths, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µcol_widths = πTemp003
							// line 260: self.extend_short_rows_with_empty_cells(max_cols,
							πF.SetLineno(260)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							πTemp008[0] = µmax_cols
							if πE = πg.CheckLocal(πF, µtable_head, "table_head"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_body, "table_body"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µtable_head, µtable_body).ToObject()
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßextend_short_rows_with_empty_cells, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 262: except SystemMessagePropagation as detail:
							πF.SetLineno(262)
						Label7:
							µdetail = πTemp011.ToObject()
							// line 263: return [detail.args[0]]
							πF.SetLineno(263)
							πTemp008 = make([]*πg.Object, 1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µdetail, ßargs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
							// line 264: except csv.Error as detail:
							πF.SetLineno(264)
						Label8:
							µdetail = πTemp011.ToObject()
							// line 265: message = str(detail)
							πF.SetLineno(265)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
								continue
							}
							πTemp008[0] = µdetail
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µmessage = πTemp003
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
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µmessage, πg.NewStr("1-character string").ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp003
						Label9:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							goto Label11
							// line 266: if sys.version_info < (3, 0) and '1-character string' in message:
							πF.SetLineno(266)
						Label10:
							// line 267: message += '\nwith Python 2.x this must be an ASCII character.'
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µmessage, πg.NewStr("\nwith Python 2.x this must be an ASCII character.").ToObject()); πE != nil {
								continue
							}
							µmessage = πTemp001
							goto Label11
						Label11:
							// line 268: error = self.state_machine.reporter.error(
							πF.SetLineno(268)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, µmessage).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error with CSV data in \"%s\" directive:\n%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
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
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µerror = πTemp003
							// line 272: return [error]
							πF.SetLineno(272)
							πTemp008 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp008[0] = µerror
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 273: table = (col_widths, table_head, table_body)
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µcol_widths, "col_widths"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_head, "table_head"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_body, "table_body"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µcol_widths, µtable_head, µtable_body).ToObject()
							µtable = πTemp001
							// line 274: table_node = self.state.build_table(table, self.content_offset,
							πF.SetLineno(274)
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							πTemp008[0] = µtable
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							πTemp008[2] = µstub_columns
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"widths", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßbuild_table, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µtable_node = πTemp001
							// line 276: table_node['classes'] += self.options.get('class', [])
							πF.SetLineno(276)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtable_node, πTemp001); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = ßclass.ToObject()
							πTemp009 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp009...).ToObject()
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp004, πE = πg.IAdd(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp005, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp003, ßalign.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label12
							}
							goto Label13
							// line 277: if 'align' in self.options:
							πF.SetLineno(277)
						Label12:
							// line 278: table_node['align'] = self.options.get('align')
							πF.SetLineno(278)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004 = ßalign.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label13
						Label13:
							// line 279: self.set_table_width(table_node)
							πF.SetLineno(279)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp008[0] = µtable_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_table_width, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 280: self.add_name(table_node)
							πF.SetLineno(280)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp008[0] = µtable_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µtitle); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label14
							}
							goto Label15
							// line 281: if title:
							πF.SetLineno(281)
						Label14:
							// line 282: table_node.insert(0, title)
							πF.SetLineno(282)
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp008[1] = µtitle
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtable_node, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label15
						Label15:
							// line 283: return [table_node] + messages
							πF.SetLineno(283)
							πTemp008 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp008[0] = µtable_node
							πTemp003 = πg.NewList(πTemp008...).ToObject()
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µmessages); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 285: def get_csv_data(self):
					πF.SetLineno(285)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("get_csv_data", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µencoding *πg.Object = πg.UnboundLocal
						_ = µencoding
						var µerror_handler *πg.Object = πg.UnboundLocal
						_ = µerror_handler
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µcsv_data *πg.Object = πg.UnboundLocal
						_ = µcsv_data
						var µsource_dir *πg.Object = πg.UnboundLocal
						_ = µsource_dir
						var µcsv_file *πg.Object = πg.UnboundLocal
						_ = µcsv_file
						var µsevere *πg.Object = πg.UnboundLocal
						_ = µsevere
						var µurlopen *πg.Object = πg.UnboundLocal
						_ = µurlopen
						var µURLError *πg.Object = πg.UnboundLocal
						_ = µURLError
						var µcsv_text *πg.Object = πg.UnboundLocal
						_ = µcsv_text
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
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
							case 18:
								goto Label18
							case 12:
								goto Label12
							default:
								panic("unexpected function state")
							}
							// line 286: """
							πF.SetLineno(286)
							// line 290: encoding = self.options.get(
							πF.SetLineno(290)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßencoding.ToObject()
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinput_encoding, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µencoding = πTemp002
							// line 292: error_handler = self.state.document.settings.input_encoding_error_handler
							πF.SetLineno(292)
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinput_encoding_error_handler, nil); πE != nil {
								continue
							}
							µerror_handler = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßfile.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 293: if self.content:
							πF.SetLineno(293)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, ßfile.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp002 = πTemp003
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 295: if 'file' in self.options or 'url' in self.options:
							πF.SetLineno(295)
						Label7:
							// line 296: error = self.state_machine.reporter.error(
							πF.SetLineno(296)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\"%s\" directive may not both specify an external file and have content.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 300: raise SystemMessagePropagation(error)
							πF.SetLineno(300)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label8
						Label8:
							// line 301: source = self.content.source(0)
							πF.SetLineno(301)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsource, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsource = πTemp002
							// line 302: csv_data = self.content
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							µcsv_data = πTemp002
							goto Label5
							// line 303: elif 'file' in self.options:
							πF.SetLineno(303)
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßurl.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 305: if 'url' in self.options:
							πF.SetLineno(305)
						Label9:
							// line 306: error = self.state_machine.reporter.error(
							πF.SetLineno(306)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("The \"file\" and \"url\" options may not be simultaneously specified for the \"%s\" directive.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 311: raise SystemMessagePropagation(error)
							πF.SetLineno(311)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label10
						Label10:
							// line 312: source_dir = os.path.dirname(
							πF.SetLineno(312)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßcurrent_source, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsource_dir = πTemp003
							// line 314: source = os.path.normpath(os.path.join(source_dir,
							πF.SetLineno(314)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsource_dir, "source_dir"); πE != nil {
								continue
							}
							πTemp007[0] = µsource_dir
							πTemp002 = ßfile.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp007[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßnormpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsource = πTemp003
							// line 316: source = utils.relative_path(None, source)
							πF.SetLineno(316)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[1] = µsource
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrelative_path, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsource = πTemp002
							// line 317: try:
							πF.SetLineno(317)
							πF.PushCheckpoint(12)
							// line 318: self.state.document.settings.record_dependencies.add(source)
							πF.SetLineno(318)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[0] = µsource
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 319: csv_file = io.FileInput(source_path=source,
							πF.SetLineno(319)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"source_path", µsource},
								{"encoding", µencoding},
								{"error_handler", µerror_handler},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßFileInput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							µcsv_file = πTemp002
							// line 322: csv_data = csv_file.read().splitlines()
							πF.SetLineno(322)
							if πE = πg.CheckLocal(πF, µcsv_file, "csv_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcsv_file, ßread, nil); πE != nil {
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
							µcsv_data = πTemp003
							πF.PopCheckpoint()
							goto Label11
						Label12:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 323: except IOError as error:
							πF.SetLineno(323)
						Label13:
							µerror = πTemp009.ToObject()
							// line 324: severe = self.state_machine.reporter.severe(
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp007[0] = µerror
							if πTemp011, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003 = πg.NewTuple2(πTemp005, πTemp012).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("Problems with \"%s\" directive path:\n%s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsevere = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsevere, "severe"); πE != nil {
								continue
							}
							πTemp001[0] = µsevere
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 329: raise SystemMessagePropagation(severe)
							πF.SetLineno(329)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label11
						Label11:
							goto Label5
							// line 330: elif 'url' in self.options:
							πF.SetLineno(330)
						Label3:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp002, πE = πg.GE(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 335: if sys.version_info >= (3, 0):
							πF.SetLineno(335)
						Label14:
							// line 336: from urllib.request import urlopen
							πF.SetLineno(336)
							if πTemp001, πE = πg.ImportModule(πF, "urllib.request"); πE != nil {
								continue
							}
							πTemp002 = πTemp001[1]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßurlopen); πE != nil {
								continue
							}
							µurlopen = πTemp003
							// line 337: from urllib.error import URLError
							πF.SetLineno(337)
							if πTemp001, πE = πg.ImportModule(πF, "urllib.error"); πE != nil {
								continue
							}
							πTemp002 = πTemp001[1]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßURLError); πE != nil {
								continue
							}
							µURLError = πTemp003
							goto Label16
						Label15:
							// line 339: from urllib2 import urlopen, URLError
							πF.SetLineno(339)
							if πTemp001, πE = πg.ImportModule(πF, "urllib2"); πE != nil {
								continue
							}
							πTemp002 = πTemp001[0]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßurlopen); πE != nil {
								continue
							}
							µurlopen = πTemp003
							πTemp002 = πTemp001[0]
							if πTemp003, πE = πg.GetAttrImport(πF, πTemp002, ßURLError); πE != nil {
								continue
							}
							µURLError = πTemp003
							goto Label16
						Label16:
							// line 341: source = self.options['url']
							πF.SetLineno(341)
							πTemp002 = ßurl.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							µsource = πTemp003
							// line 342: try:
							πF.SetLineno(342)
							πF.PushCheckpoint(18)
							// line 343: csv_text = urlopen(source).read()
							πF.SetLineno(343)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[0] = µsource
							if πE = πg.CheckLocal(πF, µurlopen, "urlopen"); πE != nil {
								continue
							}
							if πTemp002, πE = µurlopen.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcsv_text = πTemp002
							πF.PopCheckpoint()
							goto Label17
						Label18:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πE = πg.CheckLocal(πF, µURLError, "URLError"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßOSError); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(µURLError, πTemp003, πTemp005, πTemp011).ToObject()
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label19
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 344: except (URLError, IOError, OSError, ValueError) as error:
							πF.SetLineno(344)
						Label19:
							µerror = πTemp009.ToObject()
							// line 345: severe = self.state_machine.reporter.severe(
							πF.SetLineno(345)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp011 = ßurl.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp013, πTemp011); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp007[0] = µerror
							if πTemp011, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp011.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003 = πg.NewTuple3(πTemp005, πTemp012, πTemp013).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Problems with \"%s\" directive URL \"%s\":\n%s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsevere, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsevere = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsevere, "severe"); πE != nil {
								continue
							}
							πTemp001[0] = µsevere
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 350: raise SystemMessagePropagation(severe)
							πF.SetLineno(350)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label17
						Label17:
							// line 351: csv_file = io.StringInput(
							πF.SetLineno(351)
							if πE = πg.CheckLocal(πF, µcsv_text, "csv_text"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinput_encoding_error_handler, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"source", µcsv_text},
								{"source_path", µsource},
								{"encoding", µencoding},
								{"error_handler", πTemp003},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringInput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							µcsv_file = πTemp002
							// line 355: csv_data = csv_file.read().splitlines()
							πF.SetLineno(355)
							if πE = πg.CheckLocal(πF, µcsv_file, "csv_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcsv_file, ßread, nil); πE != nil {
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
							µcsv_data = πTemp003
							goto Label5
						Label4:
							// line 357: error = self.state_machine.reporter.warning(
							πF.SetLineno(357)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("The \"%s\" directive requires content; none supplied.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 361: raise SystemMessagePropagation(error)
							πF.SetLineno(361)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 362: return csv_data, source
							πF.SetLineno(362)
							if πE = πg.CheckLocal(πF, µcsv_data, "csv_data"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µcsv_data, µsource).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_csv_data.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 286: """
					πF.SetLineno(286)
					// line 286: """
					πF.SetLineno(286)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n        Get CSV data from the directive content, from an external\n        file, or from a URL reference.\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Get CSV data from the directive content, from an external\n        file, or from a URL reference.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßget_csv_data); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp006); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp008 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp006, πE = πg.LT(πF, πTemp009, πTemp008); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label1
					}
					goto Label2
					// line 364: if sys.version_info < (3, 0):
					πF.SetLineno(364)
				Label1:
					// line 366: def decode_from_csv(s):
					πF.SetLineno(366)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("decode_from_csv", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]
						_ = µs
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
							// line 367: return s.decode('utf-8')
							πF.SetLineno(367)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("utf-8").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßdecode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdecode_from_csv.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 368: def encode_for_csv(s):
					πF.SetLineno(368)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("encode_for_csv", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]
						_ = µs
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
							// line 369: return s.encode('utf-8')
							πF.SetLineno(369)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("utf-8").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßencode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßencode_for_csv.ToObject(), πTemp008); πE != nil {
						continue
					}
					goto Label3
				Label2:
					// line 371: def decode_from_csv(s):
					πF.SetLineno(371)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("decode_from_csv", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]
						_ = µs
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 372: return s
							πF.SetLineno(372)
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
					if πE = πClass.SetItem(πF, ßdecode_from_csv.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 373: def encode_for_csv(s):
					πF.SetLineno(373)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "s", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("encode_for_csv", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πArgs[0]
						_ = µs
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 374: return s
							πF.SetLineno(374)
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
					if πE = πClass.SetItem(πF, ßencode_for_csv.ToObject(), πTemp011); πE != nil {
						continue
					}
					goto Label3
				Label3:
					// line 375: decode_from_csv = staticmethod(decode_from_csv)
					πF.SetLineno(375)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdecode_from_csv); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßdecode_from_csv.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 376: encode_for_csv = staticmethod(encode_for_csv)
					πF.SetLineno(376)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßencode_for_csv); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßencode_for_csv.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 378: def parse_csv_data_into_rows(self, csv_data, dialect, source):
					πF.SetLineno(378)
					πTemp007 = make([]πg.Param, 4)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp007[1] = πg.Param{Name: "csv_data", Def: nil}
					πTemp007[2] = πg.Param{Name: "dialect", Def: nil}
					πTemp007[3] = πg.Param{Name: "source", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("parse_csv_data_into_rows", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcsv_data *πg.Object = πArgs[1]
						_ = µcsv_data
						var µdialect *πg.Object = πArgs[2]
						_ = µdialect
						var µsource *πg.Object = πArgs[3]
						_ = µsource
						var µcsv_reader *πg.Object = πg.UnboundLocal
						_ = µcsv_reader
						var µrows *πg.Object = πg.UnboundLocal
						_ = µrows
						var µmax_cols *πg.Object = πg.UnboundLocal
						_ = µmax_cols
						var µrow *πg.Object = πg.UnboundLocal
						_ = µrow
						var µrow_data *πg.Object = πg.UnboundLocal
						_ = µrow_data
						var µcell *πg.Object = πg.UnboundLocal
						_ = µcell
						var µcell_text *πg.Object = πg.UnboundLocal
						_ = µcell_text
						var µcell_data *πg.Object = πg.UnboundLocal
						_ = µcell_data
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 []*πg.Object
						_ = πTemp013
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
							// line 380: csv_reader = csv.reader([self.encode_for_csv(line + '\n')
							πF.SetLineno(380)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µcsv_data, "csv_data"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µcsv_data); πE != nil {
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
										// line 380: csv_reader = csv.reader([self.encode_for_csv(line + '\n')
										πF.SetLineno(380)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Add(πF, µline, πg.NewStr("\n").ToObject()); πE != nil {
											continue
										}
										πTemp005[0] = πTemp004
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µself, ßencode_for_csv, nil); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"dialect", µdialect},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcsv); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcsv_reader = πTemp002
							// line 383: rows = []
							πF.SetLineno(383)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µrows = πTemp002
							// line 384: max_cols = 0
							πF.SetLineno(384)
							µmax_cols = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcsv_reader, "csv_reader"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µcsv_reader); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µrow = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 386: row_data = []
							πF.SetLineno(386)
							πTemp001 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp001...).ToObject()
							µrow_data = πTemp005
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Iter(πF, µrow); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp008 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp010, πE = πg.Next(πF, πTemp005); πE != nil {
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
								µcell = πTemp010
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 389: cell_text = self.decode_from_csv(cell)
							πF.SetLineno(389)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcell, "cell"); πE != nil {
								continue
							}
							πTemp001[0] = µcell
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßdecode_from_csv, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcell_text = πTemp011
							// line 390: cell_data = (0, 0, 0, statemachine.StringList(
							πF.SetLineno(390)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcell_text, "cell_text"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µcell_text, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp012
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"source", µsource},
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßstatemachine); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßStringList, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp010 = πg.NewTuple4(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πTemp011).ToObject()
							µcell_data = πTemp010
							// line 392: row_data.append(cell_data)
							πF.SetLineno(392)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcell_data, "cell_data"); πE != nil {
								continue
							}
							πTemp001[0] = µcell_data
							if πE = πg.CheckLocal(πF, µrow_data, "row_data"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µrow_data, ßappend, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 393: rows.append(row_data)
							πF.SetLineno(393)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow_data, "row_data"); πE != nil {
								continue
							}
							πTemp001[0] = µrow_data
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µrows, ßappend, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 394: max_cols = max(max_cols, len(row))
							πF.SetLineno(394)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							πTemp001[0] = µmax_cols
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp013[0] = µrow
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[1] = πTemp010
							if πTemp005, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmax_cols = πTemp010
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 395: return rows, max_cols
							πF.SetLineno(395)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmax_cols, "max_cols"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µrows, µmax_cols).ToObject()
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
					if πE = πClass.SetItem(πF, ßparse_csv_data_into_rows.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("CSVTable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCSVTable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 398: class ListTable(Table):
			πF.SetLineno(398)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßTable); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ListTable", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 400: """
					πF.SetLineno(400)
					// line 400: """
					πF.SetLineno(400)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Implement tables whose data is encoded as a uniform two-level bullet list.\n    For further ideas, see\n    http://docutils.sf.net/docs/dev/rst/alternatives.html#list-driven-tables\n    ").ToObject()); πE != nil {
						continue
					}
					// line 406: option_spec = {'header-rows': directives.nonnegative_int,
					πF.SetLineno(406)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("header-rows").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("stub-columns").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlength_or_percentage_or_unitless, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßwidth.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(2)
					πTemp002 = πg.NewTuple1(ßauto.ToObject()).ToObject()
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpositive_int_list, nil); πE != nil {
						continue
					}
					πTemp004[1] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßvalue_or, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πTemp001.SetItem(πF, ßwidths.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßalign); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßalign.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 415: def run(self):
					πF.SetLineno(415)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µnum_cols *πg.Object = πg.UnboundLocal
						_ = µnum_cols
						var µcol_widths *πg.Object = πg.UnboundLocal
						_ = µcol_widths
						var µtable_data *πg.Object = πg.UnboundLocal
						_ = µtable_data
						var µheader_rows *πg.Object = πg.UnboundLocal
						_ = µheader_rows
						var µstub_columns *πg.Object = πg.UnboundLocal
						_ = µstub_columns
						var µdetail *πg.Object = πg.UnboundLocal
						_ = µdetail
						var µtable_node *πg.Object = πg.UnboundLocal
						_ = µtable_node
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
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
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
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
							// line 416: if not self.content:
							πF.SetLineno(416)
						Label1:
							// line 417: error = self.state_machine.reporter.error(
							πF.SetLineno(417)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("The \"%s\" directive is empty; content required.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerror = πTemp002
							// line 421: return [error]
							πF.SetLineno(421)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp004[0] = µerror
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 422: title, messages = self.make_title()
							πF.SetLineno(422)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmake_title, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µtitle = πTemp001
							µmessages = πTemp007
							// line 423: node = nodes.Element()          # anonymous container for parsing
							πF.SetLineno(423)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßElement, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnode = πTemp001
							// line 424: self.state.nested_parse(self.content, self.content_offset, node)
							πF.SetLineno(424)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[2] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 425: try:
							πF.SetLineno(425)
							πF.PushCheckpoint(4)
							// line 426: num_cols, col_widths = self.check_list_content(node)
							πF.SetLineno(426)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_list_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µnum_cols = πTemp001
							µcol_widths = πTemp007
							// line 427: table_data = [[item.children for item in row_list[0]]
							πF.SetLineno(427)
							πTemp008 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µrow_list *πg.Object = πg.UnboundLocal
								_ = µrow_list
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
								var πTemp006 []πg.Param
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
										πTemp002 = πg.NewInt(0).ToObject()
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
											µrow_list = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 427: table_data = [[item.children for item in row_list[0]]
										πF.SetLineno(427)
										πTemp006 = make([]πg.Param, 0)
										πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
											var µitem *πg.Object = πg.UnboundLocal
											_ = µitem
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
													πTemp002 = πg.NewInt(0).ToObject()
													if πE = πg.CheckLocal(πF, µrow_list, "row_list"); πE != nil {
														continue
													}
													if πTemp003, πE = πg.GetItem(πF, µrow_list, πTemp002); πE != nil {
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
														µitem = πTemp002
													}
													if πE != nil || !πTemp005 {
														continue
													}
													πF.PushCheckpoint(1)
													// line 427: table_data = [[item.children for item in row_list[0]]
													πF.SetLineno(427)
													if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
														continue
													}
													if πTemp002, πE = πg.GetAttr(πF, µitem, ßchildren, nil); πE != nil {
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
										if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp002, nil
									Label4:
										πTemp007 = πSent
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
							if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							µtable_data = πTemp001
							// line 429: header_rows = self.options.get('header-rows', 0)
							πF.SetLineno(429)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("header-rows").ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µheader_rows = πTemp001
							// line 430: stub_columns = self.options.get('stub-columns', 0)
							πF.SetLineno(430)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("stub-columns").ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstub_columns = πTemp001
							// line 431: self.check_table_dimensions(table_data, header_rows, stub_columns)
							πF.SetLineno(431)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtable_data, "table_data"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_data
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							πTemp004[1] = µheader_rows
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							πTemp004[2] = µstub_columns
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_table_dimensions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 432: except SystemMessagePropagation as detail:
							πF.SetLineno(432)
						Label5:
							µdetail = πTemp009.ToObject()
							// line 433: return [detail.args[0]]
							πF.SetLineno(433)
							πTemp004 = make([]*πg.Object, 1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µdetail, ßargs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp011, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp007
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 434: table_node = self.build_table_from_list(table_data, col_widths,
							πF.SetLineno(434)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtable_data, "table_data"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_data
							if πE = πg.CheckLocal(πF, µcol_widths, "col_widths"); πE != nil {
								continue
							}
							πTemp004[1] = µcol_widths
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							πTemp004[2] = µheader_rows
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							πTemp004[3] = µstub_columns
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbuild_table_from_list, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µtable_node = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp007, ßalign.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 436: if 'align' in self.options:
							πF.SetLineno(436)
						Label6:
							// line 437: table_node['align'] = self.options.get('align')
							πF.SetLineno(437)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp011 = ßalign.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp011, πTemp007); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 438: table_node['classes'] += self.options.get('class', [])
							πF.SetLineno(438)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtable_node, πTemp001); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßclass.ToObject()
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp011.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp011, πE = πg.IAdd(πF, πTemp007, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp012 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable_node, πTemp012, πTemp011); πE != nil {
								continue
							}
							// line 439: self.set_table_width(table_node)
							πF.SetLineno(439)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_table_width, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 440: self.add_name(table_node)
							πF.SetLineno(440)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µtitle); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 441: if title:
							πF.SetLineno(441)
						Label8:
							// line 442: table_node.insert(0, title)
							πF.SetLineno(442)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp004[1] = µtitle
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtable_node, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label9
						Label9:
							// line 443: return [table_node] + messages
							πF.SetLineno(443)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtable_node, "table_node"); πE != nil {
								continue
							}
							πTemp004[0] = µtable_node
							πTemp007 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp007, µmessages); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 445: def check_list_content(self, node):
					πF.SetLineno(445)
					πTemp005 = make([]πg.Param, 2)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("check_list_content", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µlist_node *πg.Object = πg.UnboundLocal
						_ = µlist_node
						var µnum_cols *πg.Object = πg.UnboundLocal
						_ = µnum_cols
						var µitem_index *πg.Object = πg.UnboundLocal
						_ = µitem_index
						var µitem *πg.Object = πg.UnboundLocal
						_ = µitem
						var µcol_widths *πg.Object = πg.UnboundLocal
						_ = µcol_widths
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßbullet_list, nil); πE != nil {
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
							// line 446: if len(node) != 1 or not isinstance(node[0], nodes.bullet_list):
							πF.SetLineno(446)
						Label2:
							// line 447: error = self.state_machine.reporter.error(
							πF.SetLineno(447)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error parsing content block for the \"%s\" directive: exactly one bullet list expected.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
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
							πTemp004[1] = πTemp001
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
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerror = πTemp003
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp004[0] = µerror
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 452: raise SystemMessagePropagation(error)
							πF.SetLineno(452)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 453: list_node = node[0]
							πF.SetLineno(453)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							µlist_node = πTemp003
							// line 454: num_cols = 0
							πF.SetLineno(454)
							µnum_cols = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlist_node, "list_node"); πE != nil {
								continue
							}
							πTemp008[0] = µlist_node
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
								µitem_index = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 457: item = list_node[item_index]
							πF.SetLineno(457)
							if πE = πg.CheckLocal(πF, µitem_index, "item_index"); πE != nil {
								continue
							}
							πTemp003 = µitem_index
							if πE = πg.CheckLocal(πF, µlist_node, "list_node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µlist_node, πTemp003); πE != nil {
								continue
							}
							µitem = πTemp005
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[0] = µitem
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.NE(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µitem, πTemp006); πE != nil {
								continue
							}
							πTemp004[0] = πTemp010
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp006, ßbullet_list, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp010
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp011, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp011).ToObject()
							πTemp003 = πTemp005
						Label7:
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µitem_index, "item_index"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µitem_index); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 458: if len(item) != 1 or not isinstance(item[0], nodes.bullet_list):
							πF.SetLineno(458)
						Label8:
							// line 459: error = self.state_machine.reporter.error(
							πF.SetLineno(459)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitem_index, "item_index"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, µitem_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πTemp010).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Error parsing content block for the \"%s\" directive: two-level bullet list expected, but row %s does not contain a second-level bullet list.").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßerror, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerror = πTemp005
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp004[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 465: raise SystemMessagePropagation(error)
							πF.SetLineno(465)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label11
							// line 466: elif item_index:
							πF.SetLineno(466)
						Label9:
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µitem, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µnum_cols, "num_cols"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp006, µnum_cols); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							goto Label13
							// line 467: if len(item[0]) != num_cols:
							πF.SetLineno(467)
						Label12:
							// line 468: error = self.state_machine.reporter.error(
							πF.SetLineno(468)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitem_index, "item_index"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, µitem_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							πTemp012 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µitem, πTemp012); πE != nil {
								continue
							}
							πTemp008[0] = πTemp013
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µnum_cols, "num_cols"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple4(πTemp006, πTemp010, πTemp013, µnum_cols).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Error parsing content block for the \"%s\" directive: uniform two-level bullet list expected, but row %s does not contain the same number of items as row 1 (%s vs %s).").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßerror, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerror = πTemp005
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp004[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSystemMessagePropagation); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 476: raise SystemMessagePropagation(error)
							πF.SetLineno(476)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label13
						Label13:
							goto Label11
						Label10:
							// line 478: num_cols = len(item[0])
							πF.SetLineno(478)
							πTemp004 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µitem, πTemp003); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µnum_cols = πTemp005
							goto Label11
						Label11:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 479: col_widths = self.get_column_widths(num_cols)
							πF.SetLineno(479)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnum_cols, "num_cols"); πE != nil {
								continue
							}
							πTemp004[0] = µnum_cols
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_column_widths, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µcol_widths = πTemp003
							// line 480: return num_cols, col_widths
							πF.SetLineno(480)
							if πE = πg.CheckLocal(πF, µnum_cols, "num_cols"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcol_widths, "col_widths"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µnum_cols, µcol_widths).ToObject()
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
					if πE = πClass.SetItem(πF, ßcheck_list_content.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 482: def build_table_from_list(self, table_data, col_widths, header_rows, stub_columns):
					πF.SetLineno(482)
					πTemp005 = make([]πg.Param, 5)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "table_data", Def: nil}
					πTemp005[2] = πg.Param{Name: "col_widths", Def: nil}
					πTemp005[3] = πg.Param{Name: "header_rows", Def: nil}
					πTemp005[4] = πg.Param{Name: "stub_columns", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("build_table_from_list", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/tables.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtable_data *πg.Object = πArgs[1]
						_ = µtable_data
						var µcol_widths *πg.Object = πArgs[2]
						_ = µcol_widths
						var µheader_rows *πg.Object = πArgs[3]
						_ = µheader_rows
						var µstub_columns *πg.Object = πArgs[4]
						_ = µstub_columns
						var µtable *πg.Object = πg.UnboundLocal
						_ = µtable
						var µtgroup *πg.Object = πg.UnboundLocal
						_ = µtgroup
						var µcol_width *πg.Object = πg.UnboundLocal
						_ = µcol_width
						var µcolspec *πg.Object = πg.UnboundLocal
						_ = µcolspec
						var µrows *πg.Object = πg.UnboundLocal
						_ = µrows
						var µrow *πg.Object = πg.UnboundLocal
						_ = µrow
						var µrow_node *πg.Object = πg.UnboundLocal
						_ = µrow_node
						var µcell *πg.Object = πg.UnboundLocal
						_ = µcell
						var µentry *πg.Object = πg.UnboundLocal
						_ = µentry
						var µthead *πg.Object = πg.UnboundLocal
						_ = µthead
						var µtbody *πg.Object = πg.UnboundLocal
						_ = µtbody
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 bool
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							case 11:
								goto Label11
							case 12:
								goto Label12
							case 14:
								goto Label14
							case 15:
								goto Label15
							default:
								panic("unexpected function state")
							}
							// line 483: table = nodes.table()
							πF.SetLineno(483)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtable, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtable = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßauto.ToObject()); πE != nil {
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidths, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 484: if self.widths == 'auto':
							πF.SetLineno(484)
						Label1:
							// line 485: table['classes'] += ['colwidths-auto']
							πF.SetLineno(485)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtable, πTemp001); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("colwidths-auto").ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							πTemp006 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable, πTemp006, πTemp005); πE != nil {
								continue
							}
							goto Label3
							// line 486: elif self.widths: # "grid" or list of integers
							πF.SetLineno(486)
						Label2:
							// line 487: table['classes'] += ['colwidths-given']
							πF.SetLineno(487)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtable, πTemp001); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("colwidths-given").ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							πTemp006 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtable, πTemp006, πTemp005); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 488: tgroup = nodes.tgroup(cols=len(col_widths))
							πF.SetLineno(488)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcol_widths, "col_widths"); πE != nil {
								continue
							}
							πTemp004[0] = µcol_widths
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp007 = πg.KWArgs{
								{"cols", πTemp002},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtgroup, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							µtgroup = πTemp001
							// line 489: table += tgroup
							πF.SetLineno(489)
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtgroup, "tgroup"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µtable, µtgroup); πE != nil {
								continue
							}
							µtable = πTemp001
							if πE = πg.CheckLocal(πF, µcol_widths, "col_widths"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µcol_widths); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp003 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µcol_width = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 491: colspec = nodes.colspec()
							πF.SetLineno(491)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßcolspec, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcolspec = πTemp002
							if πE = πg.CheckLocal(πF, µcol_width, "col_width"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µcol_width != πTemp005).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 492: if col_width is not None:
							πF.SetLineno(492)
						Label7:
							// line 493: colspec.attributes['colwidth'] = col_width
							πF.SetLineno(493)
							if πE = πg.CheckLocal(πF, µcol_width, "col_width"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µcol_width); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolspec, "colspec"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcolspec, ßattributes, nil); πE != nil {
								continue
							}
							πTemp006 = ßcolwidth.ToObject()
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µstub_columns); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 494: if stub_columns:
							πF.SetLineno(494)
						Label9:
							// line 495: colspec.attributes['stub'] = 1
							πF.SetLineno(495)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolspec, "colspec"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcolspec, ßattributes, nil); πE != nil {
								continue
							}
							πTemp006 = ßstub.ToObject()
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							// line 496: stub_columns -= 1
							πF.SetLineno(496)
							if πE = πg.CheckLocal(πF, µstub_columns, "stub_columns"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, µstub_columns, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstub_columns = πTemp002
							goto Label10
						Label10:
							// line 497: tgroup += colspec
							πF.SetLineno(497)
							if πE = πg.CheckLocal(πF, µtgroup, "tgroup"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolspec, "colspec"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µtgroup, µcolspec); πE != nil {
								continue
							}
							µtgroup = πTemp002
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 498: rows = []
							πF.SetLineno(498)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µrows = πTemp001
							if πE = πg.CheckLocal(πF, µtable_data, "table_data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtable_data); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp003 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µrow = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(11)
							// line 500: row_node = nodes.row()
							πF.SetLineno(500)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßrow, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrow_node = πTemp002
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µrow); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp008 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label16
							}
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µcell = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(14)
							// line 502: entry = nodes.entry()
							πF.SetLineno(502)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßentry, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							µentry = πTemp005
							// line 503: entry += cell
							πF.SetLineno(503)
							if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcell, "cell"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µentry, µcell); πE != nil {
								continue
							}
							µentry = πTemp005
							// line 504: row_node += entry
							πF.SetLineno(504)
							if πE = πg.CheckLocal(πF, µrow_node, "row_node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µrow_node, µentry); πE != nil {
								continue
							}
							µrow_node = πTemp005
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							// line 505: rows.append(row_node)
							πF.SetLineno(505)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow_node, "row_node"); πE != nil {
								continue
							}
							πTemp004[0] = µrow_node
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µrows, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µheader_rows); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label17
							}
							goto Label18
							// line 506: if header_rows:
							πF.SetLineno(506)
						Label17:
							// line 507: thead = nodes.thead()
							πF.SetLineno(507)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßthead, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µthead = πTemp001
							// line 508: thead.extend(rows[:header_rows])
							πF.SetLineno(508)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µheader_rows, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrows, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µthead, "thead"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µthead, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 509: tgroup += thead
							πF.SetLineno(509)
							if πE = πg.CheckLocal(πF, µtgroup, "tgroup"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthead, "thead"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µtgroup, µthead); πE != nil {
								continue
							}
							µtgroup = πTemp001
							goto Label18
						Label18:
							// line 510: tbody = nodes.tbody()
							πF.SetLineno(510)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtbody, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtbody = πTemp001
							// line 511: tbody.extend(rows[header_rows:])
							πF.SetLineno(511)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader_rows, "header_rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µheader_rows, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrows, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µtbody, "tbody"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtbody, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 512: tgroup += tbody
							πF.SetLineno(512)
							if πE = πg.CheckLocal(πF, µtgroup, "tgroup"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtbody, "tbody"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µtgroup, µtbody); πE != nil {
								continue
							}
							µtgroup = πTemp001
							// line 513: return table
							πF.SetLineno(513)
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							πR = µtable
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßbuild_table_from_list.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("ListTable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßListTable.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("tables", Code)
}
