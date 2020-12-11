package utils

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/itertools"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/unicodedata"
	_ "github.com/pygolin/stdlib/pkg/warnings"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßA := πg.InternStr("A")
		ßApplicationError := πg.InternStr("ApplicationError")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBadOptionDataError := πg.InternStr("BadOptionDataError")
		ßBadOptionError := πg.InternStr("BadOptionError")
		ßDEBUG_LEVEL := πg.InternStr("DEBUG_LEVEL")
		ßDataError := πg.InternStr("DataError")
		ßDependencyList := πg.InternStr("DependencyList")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßDuplicateOptionError := πg.InternStr("DuplicateOptionError")
		ßERROR_LEVEL := πg.InternStr("ERROR_LEVEL")
		ßErrorOutput := πg.InternStr("ErrorOutput")
		ßException := πg.InternStr("Exception")
		ßExtensionOptionError := πg.InternStr("ExtensionOptionError")
		ßF := πg.InternStr("F")
		ßFalse := πg.InternStr("False")
		ßFileOutput := πg.InternStr("FileOutput")
		ßH := πg.InternStr("H")
		ßINFO_LEVEL := πg.InternStr("INFO_LEVEL")
		ßKeyError := πg.InternStr("KeyError")
		ßN := πg.InternStr("N")
		ßNa := πg.InternStr("Na")
		ßNameValueError := πg.InternStr("NameValueError")
		ßNone := πg.InternStr("None")
		ßOptionParser := πg.InternStr("OptionParser")
		ßReporter := πg.InternStr("Reporter")
		ßSEVERE_LEVEL := πg.InternStr("SEVERE_LEVEL")
		ßSafeString := πg.InternStr("SafeString")
		ßSystemMessage := πg.InternStr("SystemMessage")
		ßSystemMessagePropagation := πg.InternStr("SystemMessagePropagation")
		ßText := πg.InternStr("Text")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeDecodeError := πg.InternStr("UnicodeDecodeError")
		ßValueError := πg.InternStr("ValueError")
		ßW := πg.InternStr("W")
		ßWARNING_LEVEL := πg.InternStr("WARNING_LEVEL")
		ß_ := πg.InternStr("_")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__version_info__ := πg.InternStr("__version_info__")
		ß_destination := πg.InternStr("_destination")
		ßa := πg.InternStr("a")
		ßabspath := πg.InternStr("abspath")
		ßadd := πg.InternStr("add")
		ßalpha := πg.InternStr("alpha")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßascii := πg.InternStr("ascii")
		ßassemble_option_dict := πg.InternStr("assemble_option_dict")
		ßastext := πg.InternStr("astext")
		ßattach_observer := πg.InternStr("attach_observer")
		ßb := πg.InternStr("b")
		ßbackslashreplace := πg.InternStr("backslashreplace")
		ßbase_node := πg.InternStr("base_node")
		ßbeta := πg.InternStr("beta")
		ßcandidate := πg.InternStr("candidate")
		ßchain := πg.InternStr("chain")
		ßclean_rcs_keywords := πg.InternStr("clean_rcs_keywords")
		ßclose := πg.InternStr("close")
		ßcolumn_indices := πg.InternStr("column_indices")
		ßcolumn_width := πg.InternStr("column_width")
		ßcombinations := πg.InternStr("combinations")
		ßcombining := πg.InternStr("combining")
		ßcopy := πg.InternStr("copy")
		ßdebug := πg.InternStr("debug")
		ßdebug_flag := πg.InternStr("debug_flag")
		ßdecode := πg.InternStr("decode")
		ßdecode_path := πg.InternStr("decode_path")
		ßdetach_observer := πg.InternStr("detach_observer")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßdummy_file := πg.InternStr("dummy_file")
		ßeast_asian_width := πg.InternStr("east_asian_width")
		ßeast_asian_widths := πg.InternStr("east_asian_widths")
		ßencoding := πg.InternStr("encoding")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßerror_encoding := πg.InternStr("error_encoding")
		ßerror_encoding_error_handler := πg.InternStr("error_encoding_error_handler")
		ßerror_handler := πg.InternStr("error_handler")
		ßescape2null := πg.InternStr("escape2null")
		ßexists := πg.InternStr("exists")
		ßexpanduser := πg.InternStr("expanduser")
		ßextract_extension_options := πg.InternStr("extract_extension_options")
		ßextract_name_value := πg.InternStr("extract_name_value")
		ßextract_options := πg.InternStr("extract_options")
		ßfile := πg.InternStr("file")
		ßfinal := πg.InternStr("final")
		ßfind := πg.InternStr("find")
		ßfind_combining_chars := πg.InternStr("find_combining_chars")
		ßfind_file_in_dirs := πg.InternStr("find_file_in_dirs")
		ßfootnote_references := πg.InternStr("footnote_references")
		ßget := πg.InternStr("get")
		ßget_default_values := πg.InternStr("get_default_values")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßget_source_line := πg.InternStr("get_source_line")
		ßget_stylesheet_list := πg.InternStr("get_stylesheet_list")
		ßget_stylesheet_reference := πg.InternStr("get_stylesheet_reference")
		ßget_trim_footnote_ref_space := πg.InternStr("get_trim_footnote_ref_space")
		ßgetattr := πg.InternStr("getattr")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßhalt_level := πg.InternStr("halt_level")
		ßhasattr := πg.InternStr("hasattr")
		ßinfo := πg.InternStr("info")
		ßio := πg.InternStr("io")
		ßisabs := πg.InternStr("isabs")
		ßisinstance := πg.InternStr("isinstance")
		ßitertools := πg.InternStr("itertools")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßlevels := πg.InternStr("levels")
		ßline := πg.InternStr("line")
		ßlist := πg.InternStr("list")
		ßlower := πg.InternStr("lower")
		ßlstrip := πg.InternStr("lstrip")
		ßmajor := πg.InternStr("major")
		ßmax := πg.InternStr("max")
		ßmax_level := πg.InternStr("max_level")
		ßmicro := πg.InternStr("micro")
		ßminor := πg.InternStr("minor")
		ßname := πg.InternStr("name")
		ßnew_document := πg.InternStr("new_document")
		ßnew_reporter := πg.InternStr("new_reporter")
		ßnodes := πg.InternStr("nodes")
		ßnormalize_language_tag := πg.InternStr("normalize_language_tag")
		ßnote_source := πg.InternStr("note_source")
		ßnotify_observers := πg.InternStr("notify_observers")
		ßobject := πg.InternStr("object")
		ßobservers := πg.InternStr("observers")
		ßos := πg.InternStr("os")
		ßparagraph := πg.InternStr("paragraph")
		ßparent := πg.InternStr("parent")
		ßpath := πg.InternStr("path")
		ßpop := πg.InternStr("pop")
		ßrange := πg.InternStr("range")
		ßrc := πg.InternStr("rc")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrelative_path := πg.InternStr("relative_path")
		ßrelease := πg.InternStr("release")
		ßrelease_level_abbreviations := πg.InternStr("release_level_abbreviations")
		ßreleaselevel := πg.InternStr("releaselevel")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßreport_level := πg.InternStr("report_level")
		ßreprunicode := πg.InternStr("reprunicode")
		ßreverse := πg.InternStr("reverse")
		ßsearch := πg.InternStr("search")
		ßsep := πg.InternStr("sep")
		ßserial := πg.InternStr("serial")
		ßset_conditions := πg.InternStr("set_conditions")
		ßset_output := πg.InternStr("set_output")
		ßsetdefault := πg.InternStr("setdefault")
		ßsevere := πg.InternStr("severe")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßsplit_escaped_whitespace := πg.InternStr("split_escaped_whitespace")
		ßstr := πg.InternStr("str")
		ßstream := πg.InternStr("stream")
		ßstrict := πg.InternStr("strict")
		ßstrip := πg.InternStr("strip")
		ßstrip_combining_chars := πg.InternStr("strip_combining_chars")
		ßstylesheet := πg.InternStr("stylesheet")
		ßstylesheet_dirs := πg.InternStr("stylesheet_dirs")
		ßstylesheet_path := πg.InternStr("stylesheet_path")
		ßsub := πg.InternStr("sub")
		ßsum := πg.InternStr("sum")
		ßsuperscript := πg.InternStr("superscript")
		ßsys := πg.InternStr("sys")
		ßsystem_message := πg.InternStr("system_message")
		ßtrim_footnote_reference_space := πg.InternStr("trim_footnote_reference_space")
		ßtype := πg.InternStr("type")
		ßunescape := πg.InternStr("unescape")
		ßunicode := πg.InternStr("unicode")
		ßunicodedata := πg.InternStr("unicodedata")
		ßuniq := πg.InternStr("uniq")
		ßunique_combinations := πg.InternStr("unique_combinations")
		ßutf8 := πg.InternStr("utf8")
		ßversion_identifier := πg.InternStr("version_identifier")
		ßversion_info := πg.InternStr("version_info")
		ßwarn := πg.InternStr("warn")
		ßwarning := πg.InternStr("warning")
		ßwarning_stream := πg.InternStr("warning_stream")
		ßwarnings := πg.InternStr("warnings")
		ßwrite := πg.InternStr("write")
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
			// line 6: """
			πF.SetLineno(6)
			// line 6: """
			πF.SetLineno(6)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nMiscellaneous utilities for the documentation utilities.\n").ToObject()); πE != nil {
				continue
			}
			// line 10: __docformat__ = 'reStructuredText'
			πF.SetLineno(10)
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
			// line 13: import os
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import os.path
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: import re
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import itertools
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import warnings
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import unicodedata
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "unicodedata"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunicodedata.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: from docutils import ApplicationError, DataError, __version_info__
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßApplicationError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßApplicationError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDataError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDataError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß__version_info__); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__version_info__.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: from docutils import nodes
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: from docutils.nodes import unescape
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßunescape); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunescape.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: import docutils.io
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: from docutils.utils.error_reporting import ErrorOutput, SafeString
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorOutput); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorOutput.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSafeString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSafeString.ToObject(), πTemp003); πE != nil {
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
			// line 25: if sys.version_info >= (3, 0):
			πF.SetLineno(25)
		Label1:
			// line 26: unicode = str
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 29: class SystemMessage(ApplicationError):
			πF.SetLineno(29)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßApplicationError); πE != nil {
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
			_, πE = πg.NewCode("SystemMessage", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 31: def __init__(self, system_message, level):
					πF.SetLineno(31)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "system_message", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsystem_message *πg.Object = πArgs[1]
						_ = µsystem_message
						var µlevel *πg.Object = πArgs[2]
						_ = µlevel
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
							// line 32: Exception.__init__(self, system_message.astext())
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µsystem_message, "system_message"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsystem_message, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 33: self.level = level
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µlevel); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
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
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SystemMessage").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSystemMessage.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 36: class SystemMessagePropagation(ApplicationError): pass
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßApplicationError); πE != nil {
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
			_, πE = πg.NewCode("SystemMessagePropagation", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 36: class SystemMessagePropagation(ApplicationError): pass
					πF.SetLineno(36)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SystemMessagePropagation").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSystemMessagePropagation.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 39: class Reporter(object):
			πF.SetLineno(39)
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
			_, πE = πg.NewCode("Reporter", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []πg.Param
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
					// line 41: """
					πF.SetLineno(41)
					// line 41: """
					πF.SetLineno(41)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Info/warning/error reporter and ``system_message`` element generator.\n\n    Five levels of system messages are defined, along with corresponding\n    methods: `debug()`, `info()`, `warning()`, `error()`, and `severe()`.\n\n    There is typically one Reporter object per process.  A Reporter object is\n    instantiated with thresholds for reporting (generating warnings) and\n    halting processing (raising exceptions), a switch to turn debug output on\n    or off, and an I/O stream for warnings.  These are stored as instance\n    attributes.\n\n    When a system message is generated, its level is compared to the stored\n    thresholds, and a warning or error is generated as appropriate.  Debug\n    messages are produced if the stored debug switch is on, independently of\n    other thresholds.  Message output is sent to the stored warning stream if\n    not set to ''.\n\n    The Reporter class also employs a modified form of the \"Observer\" pattern\n    [GoF95]_ to track system messages generated.  The `attach_observer` method\n    should be called before parsing, with a bound method or function which\n    accepts system messages.  The observer can be removed with\n    `detach_observer`, and another added in its place.\n\n    .. [GoF95] Gamma, Helm, Johnson, Vlissides. *Design Patterns: Elements of\n       Reusable Object-Oriented Software*. Addison-Wesley, Reading, MA, USA,\n       1995.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 70: levels = 'DEBUG INFO WARNING ERROR SEVERE'.split()
					πF.SetLineno(70)
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("DEBUG INFO WARNING ERROR SEVERE").ToObject(), ßsplit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßlevels.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 71: """List of names for system message levels, indexed by level."""
					πF.SetLineno(71)
					// line 74: (DEBUG_LEVEL,
					πF.SetLineno(74)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewInt(5).ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßrange); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßDEBUG_LEVEL.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßINFO_LEVEL.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßWARNING_LEVEL.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßERROR_LEVEL.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßSEVERE_LEVEL.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 80: def __init__(self, source, report_level, halt_level, stream=None,
					πF.SetLineno(80)
					πTemp008 = make([]πg.Param, 8)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp008[1] = πg.Param{Name: "source", Def: nil}
					πTemp008[2] = πg.Param{Name: "report_level", Def: nil}
					πTemp008[3] = πg.Param{Name: "halt_level", Def: nil}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp008[4] = πg.Param{Name: "stream", Def: πTemp002}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp008[5] = πg.Param{Name: "debug", Def: πTemp002}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp008[6] = πg.Param{Name: "encoding", Def: πTemp002}
					πTemp008[7] = πg.Param{Name: "error_handler", Def: ßbackslashreplace.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πArgs[1]
						_ = µsource
						var µreport_level *πg.Object = πArgs[2]
						_ = µreport_level
						var µhalt_level *πg.Object = πArgs[3]
						_ = µhalt_level
						var µstream *πg.Object = πArgs[4]
						_ = µstream
						var µdebug *πg.Object = πArgs[5]
						_ = µdebug
						var µencoding *πg.Object = πArgs[6]
						_ = µencoding
						var µerror_handler *πg.Object = πArgs[7]
						_ = µerror_handler
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
							// line 82: """
							πF.SetLineno(82)
							// line 98: self.source = source
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp001); πE != nil {
								continue
							}
							// line 99: """The path to or description of the source data."""
							πF.SetLineno(99)
							// line 101: self.error_handler = error_handler
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µerror_handler); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßerror_handler, πTemp001); πE != nil {
								continue
							}
							// line 102: """The character encoding error handler."""
							πF.SetLineno(102)
							// line 104: self.debug_flag = debug
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdebug); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdebug_flag, πTemp001); πE != nil {
								continue
							}
							// line 105: """Show debug (level=0) system messages?"""
							πF.SetLineno(105)
							// line 107: self.report_level = report_level
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µreport_level, "report_level"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreport_level); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreport_level, πTemp001); πE != nil {
								continue
							}
							// line 108: """The level at or above which warning output will be sent
							πF.SetLineno(108)
							// line 111: self.halt_level = halt_level
							πF.SetLineno(111)
							if πE = πg.CheckLocal(πF, µhalt_level, "halt_level"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µhalt_level); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhalt_level, πTemp001); πE != nil {
								continue
							}
							// line 112: """The level at or above which `SystemMessage` exceptions
							πF.SetLineno(112)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[0] = µstream
							if πTemp003, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
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
							// line 115: if not isinstance(stream, ErrorOutput):
							πF.SetLineno(115)
						Label1:
							// line 116: stream = ErrorOutput(stream, encoding, error_handler)
							πF.SetLineno(116)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[0] = µstream
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp002[1] = µencoding
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							πTemp002[2] = µerror_handler
							if πTemp001, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µstream = πTemp003
							goto Label2
						Label2:
							// line 118: self.stream = stream
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstream); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstream, πTemp001); πE != nil {
								continue
							}
							// line 119: """Where warning output is sent."""
							πF.SetLineno(119)
							// line 121: self.encoding = encoding or getattr(stream, 'encoding', 'ascii')
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001 = µencoding
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp002[0] = µstream
							πTemp002[1] = ßencoding.ToObject()
							πTemp002[2] = ßascii.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp003); πE != nil {
								continue
							}
							// line 122: """The output character encoding."""
							πF.SetLineno(122)
							// line 124: self.observers = []
							πF.SetLineno(124)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßobservers, πTemp003); πE != nil {
								continue
							}
							// line 125: """List of bound methods or functions to call with each system_message
							πF.SetLineno(125)
							// line 128: self.max_level = -1
							πF.SetLineno(128)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmax_level, πTemp003); πE != nil {
								continue
							}
							// line 129: """The highest level system message generated so far."""
							πF.SetLineno(129)
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
					// line 82: """
					πF.SetLineno(82)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("\n        :Parameters:\n            - `source`: The path to or description of the source data.\n            - `report_level`: The level at or above which warning output will\n              be sent to `stream`.\n            - `halt_level`: The level at or above which `SystemMessage`\n              exceptions will be raised, halting execution.\n            - `debug`: Show debug (level=0) system messages?\n            - `stream`: Where warning output is sent.  Can be file-like (has a\n              ``.write`` method), a string (file name, opened for writing),\n              '' (empty string) or `False` (for discarding all stream messages)\n              or `None` (implies `sys.stderr`; default).\n            - `encoding`: The output encoding.\n            - `error_handler`: The error handler for stderr output encoding.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp002); πE != nil {
						continue
					}
					// line 131: def set_conditions(self, category, report_level, halt_level,
					πF.SetLineno(131)
					πTemp008 = make([]πg.Param, 6)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp008[1] = πg.Param{Name: "category", Def: nil}
					πTemp008[2] = πg.Param{Name: "report_level", Def: nil}
					πTemp008[3] = πg.Param{Name: "halt_level", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp008[4] = πg.Param{Name: "stream", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp008[5] = πg.Param{Name: "debug", Def: πTemp004}
					πTemp002 = πg.NewFunction(πg.NewCode("set_conditions", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcategory *πg.Object = πArgs[1]
						_ = µcategory
						var µreport_level *πg.Object = πArgs[2]
						_ = µreport_level
						var µhalt_level *πg.Object = πArgs[3]
						_ = µhalt_level
						var µstream *πg.Object = πArgs[4]
						_ = µstream
						var µdebug *πg.Object = πArgs[5]
						_ = µdebug
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							// line 133: warnings.warn('docutils.utils.Reporter.set_conditions deprecated; '
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("docutils.utils.Reporter.set_conditions deprecated; set attributes via configuration settings or directly").ToObject()
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
							// line 136: self.report_level = report_level
							πF.SetLineno(136)
							if πE = πg.CheckLocal(πF, µreport_level, "report_level"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µreport_level); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreport_level, πTemp002); πE != nil {
								continue
							}
							// line 137: self.halt_level = halt_level
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µhalt_level, "halt_level"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µhalt_level); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhalt_level, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp001[0] = µstream
							if πTemp004, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
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
							// line 138: if not isinstance(stream, ErrorOutput):
							πF.SetLineno(138)
						Label1:
							// line 139: stream = ErrorOutput(stream, self.encoding, self.error_handler)
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp001[0] = µstream
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror_handler, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstream = πTemp004
							goto Label2
						Label2:
							// line 140: self.stream = stream
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µstream); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstream, πTemp002); πE != nil {
								continue
							}
							// line 141: self.debug_flag = debug
							πF.SetLineno(141)
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdebug); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdebug_flag, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_conditions.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 143: def attach_observer(self, observer):
					πF.SetLineno(143)
					πTemp008 = make([]πg.Param, 2)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp008[1] = πg.Param{Name: "observer", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("attach_observer", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobserver *πg.Object = πArgs[1]
						_ = µobserver
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
							// line 144: """
							πF.SetLineno(144)
							// line 148: self.observers.append(observer)
							πF.SetLineno(148)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobserver, "observer"); πE != nil {
								continue
							}
							πTemp001[0] = µobserver
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßobservers, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßattach_observer.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 144: """
					πF.SetLineno(144)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        The `observer` parameter is a function or bound method which takes one\n        argument, a `nodes.system_message` instance.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßattach_observer); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 150: def detach_observer(self, observer):
					πF.SetLineno(150)
					πTemp008 = make([]πg.Param, 2)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp008[1] = πg.Param{Name: "observer", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("detach_observer", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobserver *πg.Object = πArgs[1]
						_ = µobserver
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
							// line 151: self.observers.remove(observer)
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobserver, "observer"); πE != nil {
								continue
							}
							πTemp001[0] = µobserver
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßobservers, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdetach_observer.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 153: def notify_observers(self, message):
					πF.SetLineno(153)
					πTemp008 = make([]πg.Param, 2)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp008[1] = πg.Param{Name: "message", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("notify_observers", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
						var µobserver *πg.Object = πg.UnboundLocal
						_ = µobserver
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßobservers, nil); πE != nil {
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
								µobserver = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 155: observer(message)
							πF.SetLineno(155)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp005[0] = µmessage
							if πE = πg.CheckLocal(πF, µobserver, "observer"); πE != nil {
								continue
							}
							if πTemp002, πE = µobserver.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnotify_observers.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 157: def system_message(self, level, message, *children, **kwargs):
					πF.SetLineno(157)
					πTemp008 = make([]πg.Param, 3)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp008[1] = πg.Param{Name: "level", Def: nil}
					πTemp008[2] = πg.Param{Name: "message", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("system_message", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlevel *πg.Object = πArgs[1]
						_ = µlevel
						var µmessage *πg.Object = πArgs[2]
						_ = µmessage
						var µchildren *πg.Object = πArgs[3]
						_ = µchildren
						var µkwargs *πg.Object = πArgs[4]
						_ = µkwargs
						var µattributes *πg.Object = πg.UnboundLocal
						_ = µattributes
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
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
							case 12:
								goto Label12
							default:
								panic("unexpected function state")
							}
							// line 158: """
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[0] = µmessage
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
							// line 164: if isinstance(message, Exception):
							πF.SetLineno(164)
						Label1:
							// line 165: message = SafeString(message)
							πF.SetLineno(165)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[0] = µmessage
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmessage = πTemp003
							goto Label2
						Label2:
							// line 167: attributes = kwargs.copy()
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µkwargs, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µattributes = πTemp003
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µkwargs, ßbase_node.ToObject()); πE != nil {
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
							// line 168: if 'base_node' in kwargs:
							πF.SetLineno(168)
						Label3:
							// line 169: source, line = get_source_line(kwargs['base_node'])
							πF.SetLineno(169)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = ßbase_node.ToObject()
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µkwargs, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßget_source_line); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µsource = πTemp002
							µline = πTemp005
							// line 170: del attributes['base_node']
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							πTemp002 = ßbase_node.ToObject()
							if πE = πg.DelItem(πF, µattributes, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µsource != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 171: if source is not None:
							πF.SetLineno(171)
						Label5:
							// line 172: attributes.setdefault('source', source)
							πF.SetLineno(172)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[1] = µsource
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattributes, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µline != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 173: if line is not None:
							πF.SetLineno(173)
						Label7:
							// line 174: attributes.setdefault('line', line)
							πF.SetLineno(174)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßline.ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[1] = µline
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattributes, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µattributes, ßsource.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 176: if not 'source' in attributes: # 'line' is absolute line number
							πF.SetLineno(176)
						Label9:
							// line 177: try: # look up (source, line-in-source)
							πF.SetLineno(177)
							πF.PushCheckpoint(12)
							// line 178: source, line = self.get_source_and_line(attributes.get('line'))
							πF.SetLineno(178)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßline.ToObject()
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattributes, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_source_and_line, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µsource = πTemp002
							µline = πTemp005
							πF.PopCheckpoint()
							goto Label11
						Label12:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 179: except AttributeError:
							πF.SetLineno(179)
						Label13:
							// line 180: source, line = None, None
							πF.SetLineno(180)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µsource = πTemp003
							µline = πTemp005
							πF.RestoreExc(nil, nil)
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µsource != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 181: if source is not None:
							πF.SetLineno(181)
						Label14:
							// line 182: attributes['source'] = source
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							πTemp003 = ßsource.ToObject()
							if πE = πg.SetItem(πF, µattributes, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label15
						Label15:
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µline != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							goto Label17
							// line 183: if line is not None:
							πF.SetLineno(183)
						Label16:
							// line 184: attributes['line'] = line
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µline); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							πTemp003 = ßline.ToObject()
							if πE = πg.SetItem(πF, µattributes, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label17
						Label17:
							goto Label10
						Label10:
							// line 187: attributes.setdefault('source', self.source)
							πF.SetLineno(187)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µattributes, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 189: msg = nodes.system_message(message, level=level,
							πF.SetLineno(189)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[0] = µmessage
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp002 = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlevels, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"level", µlevel},
								{"type", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µchildren, πTemp009, µattributes); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label18
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßreport_level, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GE(πF, µlevel, πTemp011); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label19
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßdebug_flag, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp011
							if πTemp012, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µself, ßDEBUG_LEVEL, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Eq(πF, µlevel, πTemp013); πE != nil {
								continue
							}
							πTemp005 = πTemp011
						Label20:
							πTemp003 = πTemp005
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label19
							}
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßhalt_level, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GE(πF, µlevel, πTemp011); πE != nil {
								continue
							}
							πTemp003 = πTemp005
						Label19:
							πTemp002 = πTemp003
						Label18:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 192: if self.stream and (level >= self.report_level
							πF.SetLineno(192)
						Label21:
							// line 195: self.stream.write(msg.astext() + '\n')
							πF.SetLineno(195)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßastext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label22
						Label22:
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhalt_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µlevel, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							goto Label24
							// line 196: if level >= self.halt_level:
							πF.SetLineno(196)
						Label23:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[1] = µlevel
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSystemMessage); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 197: raise SystemMessage(msg, level)
							πF.SetLineno(197)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label24
						Label24:
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßDEBUG_LEVEL, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µlevel, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label25
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdebug_flag, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label25:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label26
							}
							goto Label27
							// line 198: if level > self.DEBUG_LEVEL or self.debug_flag:
							πF.SetLineno(198)
						Label26:
							// line 199: self.notify_observers(msg)
							πF.SetLineno(199)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnotify_observers, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label27
						Label27:
							// line 200: self.max_level = max(level, self.max_level)
							πF.SetLineno(200)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[0] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmax_level, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßmax_level, πTemp002); πE != nil {
								continue
							}
							// line 201: return msg
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πR = µmsg
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsystem_message.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 158: """
					πF.SetLineno(158)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Return a system_message object.\n\n        Raise an exception or generate a warning if appropriate.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßsystem_message); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 203: def debug(self, *args, **kwargs):
					πF.SetLineno(203)
					πTemp008 = make([]πg.Param, 1)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("debug", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 204: """
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug_flag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 209: if self.debug_flag:
							πF.SetLineno(209)
						Label1:
							// line 210: return self.system_message(self.DEBUG_LEVEL, *args, **kwargs)
							πF.SetLineno(210)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßDEBUG_LEVEL, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp001, πTemp003, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp004
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
					if πE = πClass.SetItem(πF, ßdebug.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 204: """
					πF.SetLineno(204)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n        Level-0, \"DEBUG\": an internal reporting issue. Typically, there is no\n        effect on the processing. Level-0 system messages are handled\n        separately from the others.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdebug); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 212: def info(self, *args, **kwargs):
					πF.SetLineno(212)
					πTemp008 = make([]πg.Param, 1)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("info", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 213: """
							πF.SetLineno(213)
							// line 217: return self.system_message(self.INFO_LEVEL, *args, **kwargs)
							πF.SetLineno(217)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßINFO_LEVEL, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkwargs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßinfo.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 213: """
					πF.SetLineno(213)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n        Level-1, \"INFO\": a minor issue that can be ignored. Typically there is\n        no effect on processing, and level-1 system messages are not reported.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßinfo); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 219: def warning(self, *args, **kwargs):
					πF.SetLineno(219)
					πTemp008 = make([]πg.Param, 1)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("warning", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 220: """
							πF.SetLineno(220)
							// line 224: return self.system_message(self.WARNING_LEVEL, *args, **kwargs)
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßWARNING_LEVEL, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkwargs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwarning.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 220: """
					πF.SetLineno(220)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n        Level-2, \"WARNING\": an issue that should be addressed. If ignored,\n        there may be unpredictable problems with the output.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßwarning); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 226: def error(self, *args, **kwargs):
					πF.SetLineno(226)
					πTemp008 = make([]πg.Param, 1)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("error", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 227: """
							πF.SetLineno(227)
							// line 231: return self.system_message(self.ERROR_LEVEL, *args, **kwargs)
							πF.SetLineno(231)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßERROR_LEVEL, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkwargs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßerror.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 227: """
					πF.SetLineno(227)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n        Level-3, \"ERROR\": an error that should be addressed. If ignored, the\n        output will contain errors.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßerror); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 233: def severe(self, *args, **kwargs):
					πF.SetLineno(233)
					πTemp008 = make([]πg.Param, 1)
					πTemp008[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("severe", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp008, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 234: """
							πF.SetLineno(234)
							// line 239: return self.system_message(self.SEVERE_LEVEL, *args, **kwargs)
							πF.SetLineno(239)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSEVERE_LEVEL, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsystem_message, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkwargs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsevere.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 234: """
					πF.SetLineno(234)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n        Level-4, \"SEVERE\": a severe error that must be addressed. If ignored,\n        the output will contain severe errors. Typically level-4 system\n        messages are turned into exceptions which halt processing.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßsevere); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Reporter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReporter.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 242: class ExtensionOptionError(DataError): pass
			πF.SetLineno(242)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
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
			_, πE = πg.NewCode("ExtensionOptionError", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 242: class ExtensionOptionError(DataError): pass
					πF.SetLineno(242)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ExtensionOptionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExtensionOptionError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 243: class BadOptionError(ExtensionOptionError): pass
			πF.SetLineno(243)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßExtensionOptionError); πE != nil {
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
			_, πE = πg.NewCode("BadOptionError", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 243: class BadOptionError(ExtensionOptionError): pass
					πF.SetLineno(243)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadOptionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBadOptionError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 244: class BadOptionDataError(ExtensionOptionError): pass
			πF.SetLineno(244)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßExtensionOptionError); πE != nil {
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
			_, πE = πg.NewCode("BadOptionDataError", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 244: class BadOptionDataError(ExtensionOptionError): pass
					πF.SetLineno(244)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadOptionDataError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBadOptionDataError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 245: class DuplicateOptionError(ExtensionOptionError): pass
			πF.SetLineno(245)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßExtensionOptionError); πE != nil {
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
			_, πE = πg.NewCode("DuplicateOptionError", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 245: class DuplicateOptionError(ExtensionOptionError): pass
					πF.SetLineno(245)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DuplicateOptionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDuplicateOptionError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 248: def extract_extension_options(field_list, options_spec):
			πF.SetLineno(248)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "field_list", Def: nil}
			πTemp007[1] = πg.Param{Name: "options_spec", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("extract_extension_options", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfield_list *πg.Object = πArgs[0]
				_ = µfield_list
				var µoptions_spec *πg.Object = πArgs[1]
				_ = µoptions_spec
				var µoption_list *πg.Object = πg.UnboundLocal
				_ = µoption_list
				var µoption_dict *πg.Object = πg.UnboundLocal
				_ = µoption_dict
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
					// line 249: """
					πF.SetLineno(249)
					// line 269: option_list = extract_options(field_list)
					πF.SetLineno(269)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfield_list, "field_list"); πE != nil {
						continue
					}
					πTemp001[0] = µfield_list
					if πTemp002, πE = πg.ResolveGlobal(πF, ßextract_options); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µoption_list = πTemp003
					// line 270: option_dict = assemble_option_dict(option_list, options_spec)
					πF.SetLineno(270)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
						continue
					}
					πTemp001[0] = µoption_list
					if πE = πg.CheckLocal(πF, µoptions_spec, "options_spec"); πE != nil {
						continue
					}
					πTemp001[1] = µoptions_spec
					if πTemp002, πE = πg.ResolveGlobal(πF, ßassemble_option_dict); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µoption_dict = πTemp003
					// line 271: return option_dict
					πF.SetLineno(271)
					if πE = πg.CheckLocal(πF, µoption_dict, "option_dict"); πE != nil {
						continue
					}
					πR = µoption_dict
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßextract_extension_options.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 249: """
			πF.SetLineno(249)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Return a dictionary mapping extension option names to converted values.\n\n    :Parameters:\n        - `field_list`: A flat field list without field arguments, where each\n          field body consists of a single paragraph only.\n        - `options_spec`: Dictionary mapping known option names to a\n          conversion function such as `int` or `float`.\n\n    :Exceptions:\n        - `KeyError` for unknown option names.\n        - `ValueError` for invalid option values (raised by the conversion\n           function).\n        - `TypeError` for invalid option value types (raised by conversion\n           function).\n        - `DuplicateOptionError` for duplicate options.\n        - `BadOptionError` for invalid fields.\n        - `BadOptionDataError` for invalid option data (missing name,\n          missing data, bad quotes, etc.).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßextract_extension_options); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 273: def extract_options(field_list):
			πF.SetLineno(273)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "field_list", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("extract_options", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfield_list *πg.Object = πArgs[0]
				_ = µfield_list
				var µoption_list *πg.Object = πg.UnboundLocal
				_ = µoption_list
				var µfield *πg.Object = πg.UnboundLocal
				_ = µfield
				var µname *πg.Object = πg.UnboundLocal
				_ = µname
				var µbody *πg.Object = πg.UnboundLocal
				_ = µbody
				var µdata *πg.Object = πg.UnboundLocal
				_ = µdata
				var πTemp001 []*πg.Object
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
				var πTemp007 *πg.Object
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
					default:
						panic("unexpected function state")
					}
					// line 274: """
					πF.SetLineno(274)
					// line 286: option_list = []
					πF.SetLineno(286)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µoption_list = πTemp002
					if πE = πg.CheckLocal(πF, µfield_list, "field_list"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µfield_list); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µfield = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µfield, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßastext, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.NE(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 288: if len(field[0].astext().split()) != 1:
					πF.SetLineno(288)
				Label4:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("extension option field name may not contain multiple words").ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßBadOptionError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 289: raise BadOptionError(
					πF.SetLineno(289)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label5
				Label5:
					// line 291: name = str(field[0].astext().lower())
					πF.SetLineno(291)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfield, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßastext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßlower, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µname = πTemp006
					// line 292: body = field[1]
					πF.SetLineno(292)
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µfield, πTemp005); πE != nil {
						continue
					}
					µbody = πTemp006
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
						continue
					}
					πTemp001[0] = µbody
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
						continue
					}
					πTemp001[0] = µbody
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.GT(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					πTemp001 = πF.MakeArgs(2)
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µbody, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßparagraph, nil); πE != nil {
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
					if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(!πTemp009).ToObject()
					πTemp005 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µbody, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.NE(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					πTemp001 = πF.MakeArgs(2)
					πTemp007 = πg.NewInt(0).ToObject()
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µbody, πTemp010); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßText, nil); πE != nil {
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
					if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(!πTemp009).ToObject()
					πTemp005 = πTemp006
				Label7:
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 293: if len(body) == 0:
					πF.SetLineno(293)
				Label6:
					// line 294: data = None
					πF.SetLineno(294)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µdata = πTemp005
					goto Label10
					// line 295: elif len(body) > 1 or not isinstance(body[0], nodes.paragraph) \
					πF.SetLineno(295)
				Label8:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("extension option field body may contain\na single paragraph only (option \"%s\")").ToObject(), µname); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßBadOptionDataError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 297: raise BadOptionDataError(
					πF.SetLineno(297)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label10
				Label9:
					// line 301: data = body[0][0].astext()
					πF.SetLineno(301)
					πTemp005 = πg.NewInt(0).ToObject()
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µbody, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßastext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					µdata = πTemp006
					goto Label10
				Label10:
					// line 302: option_list.append((name, data))
					πF.SetLineno(302)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µname, µdata).ToObject()
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µoption_list, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 303: return option_list
					πF.SetLineno(303)
					if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
						continue
					}
					πR = µoption_list
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßextract_options.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 274: """
			πF.SetLineno(274)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Return a list of option (name, value) pairs from field names & bodies.\n\n    :Parameter:\n        `field_list`: A flat field list, where each field name is a single\n        word and each field body consists of a single paragraph only.\n\n    :Exceptions:\n        - `BadOptionError` for invalid fields.\n        - `BadOptionDataError` for invalid option data (missing name,\n          missing data, bad quotes, etc.).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßextract_options); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 305: def assemble_option_dict(option_list, options_spec):
			πF.SetLineno(305)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "option_list", Def: nil}
			πTemp007[1] = πg.Param{Name: "options_spec", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("assemble_option_dict", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoption_list *πg.Object = πArgs[0]
				_ = µoption_list
				var µoptions_spec *πg.Object = πArgs[1]
				_ = µoptions_spec
				var µoptions *πg.Object = πg.UnboundLocal
				_ = µoptions
				var µname *πg.Object = πg.UnboundLocal
				_ = µname
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var µconvertor *πg.Object = πg.UnboundLocal
				_ = µconvertor
				var µdetail *πg.Object = πg.UnboundLocal
				_ = µdetail
				var πTemp001 *πg.Dict
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πTemp011 []*πg.Object
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
					case 9:
						goto Label9
					default:
						panic("unexpected function state")
					}
					// line 306: """
					πF.SetLineno(306)
					// line 323: options = {}
					πF.SetLineno(323)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µoptions = πTemp002
					if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µoption_list); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
							continue
						}
						µname = πTemp006
						µvalue = πTemp007
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 325: convertor = options_spec[name]  # raises KeyError if unknown
					πF.SetLineno(325)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp005 = µname
					if πE = πg.CheckLocal(πF, µoptions_spec, "options_spec"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µoptions_spec, πTemp005); πE != nil {
						continue
					}
					µconvertor = πTemp006
					if πE = πg.CheckLocal(πF, µconvertor, "convertor"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(µconvertor == πTemp006).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 326: if convertor is None:
					πF.SetLineno(326)
				Label4:
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp008[0] = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					// line 327: raise KeyError(name)        # or if explicitly disabled
					πF.SetLineno(327)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µoptions, µname); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 328: if name in options:
					πF.SetLineno(328)
				Label6:
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("duplicate option \"%s\"").ToObject(), µname); πE != nil {
						continue
					}
					πTemp008[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßDuplicateOptionError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					// line 329: raise DuplicateOptionError('duplicate option "%s"' % name)
					πF.SetLineno(329)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label7
				Label7:
					// line 330: try:
					πF.SetLineno(330)
					πF.PushCheckpoint(9)
					// line 331: options[name] = convertor(value)
					πF.SetLineno(331)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp008[0] = µvalue
					if πE = πg.CheckLocal(πF, µconvertor, "convertor"); πE != nil {
						continue
					}
					if πTemp005, πE = µconvertor.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp007 = µname
					if πE = πg.SetItem(πF, µoptions, πTemp007, πTemp006); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label8
				Label9:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 332: except (ValueError, TypeError) as detail:
					πF.SetLineno(332)
				Label10:
					µdetail = πTemp009.ToObject()
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp011 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µdetail, ßargs, nil); πE != nil {
						continue
					}
					πTemp011[0] = πTemp007
					if πTemp007, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp007.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006 = πg.NewTuple3(µname, µvalue, πTemp012).ToObject()
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("(option: \"%s\"; value: %r)\n%s").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp008[0] = πTemp005
					if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µdetail, ß__class__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					// line 333: raise detail.__class__('(option: "%s"; value: %r)\n%s'
					πF.SetLineno(333)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label8
				Label8:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 335: return options
					πF.SetLineno(335)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πR = µoptions
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßassemble_option_dict.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 306: """
			πF.SetLineno(306)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n    Return a mapping of option names to values.\n\n    :Parameters:\n        - `option_list`: A list of (name, value) pairs (the output of\n          `extract_options()`).\n        - `options_spec`: Dictionary mapping known option names to a\n          conversion function such as `int` or `float`.\n\n    :Exceptions:\n        - `KeyError` for unknown option names.\n        - `DuplicateOptionError` for duplicate options.\n        - `ValueError` for invalid option values (raised by conversion\n           function).\n        - `TypeError` for invalid option value types (raised by conversion\n           function).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßassemble_option_dict); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 338: class NameValueError(DataError): pass
			πF.SetLineno(338)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp006 = πg.NewDict()
			if πTemp008, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp008); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NameValueError", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 338: class NameValueError(DataError): pass
					πF.SetLineno(338)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp009, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp009 == nil {
				πTemp009 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp009.Call(πF, []*πg.Object{πg.NewStr("NameValueError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNameValueError.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 341: def decode_path(path):
			πF.SetLineno(341)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "path", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("decode_path", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]
				_ = µpath
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
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8:
						goto Label8
					case 4:
						goto Label4
					default:
						panic("unexpected function state")
					}
					// line 342: """
					πF.SetLineno(342)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
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
					// line 348: if isinstance(path, unicode):
					πF.SetLineno(348)
				Label1:
					// line 349: return path
					πF.SetLineno(349)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πR = µpath
					continue
					goto Label2
				Label2:
					// line 350: try:
					πF.SetLineno(350)
					πF.PushCheckpoint(4)
					// line 351: path = path.decode(sys.getfilesystemencoding(), 'strict')
					πF.SetLineno(351)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetfilesystemencoding, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = ßstrict.ToObject()
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpath, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpath = πTemp003
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 352: except AttributeError: # default value None has no decode method
					πF.SetLineno(352)
				Label5:
					// line 353: return nodes.reprunicode(path)
					πF.SetLineno(353)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreprunicode, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label3
					// line 354: except UnicodeDecodeError:
					πF.SetLineno(354)
				Label6:
					// line 355: try:
					πF.SetLineno(355)
					πF.PushCheckpoint(8)
					// line 356: path = path.decode('utf-8', 'strict')
					πF.SetLineno(356)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("utf-8").ToObject()
					πTemp001[1] = ßstrict.ToObject()
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpath, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpath = πTemp003
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp006 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 357: except UnicodeDecodeError:
					πF.SetLineno(357)
				Label9:
					// line 358: path = path.decode('ascii', 'replace')
					πF.SetLineno(358)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßascii.ToObject()
					πTemp001[1] = ßreplace.ToObject()
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpath, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpath = πTemp003
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 359: return nodes.reprunicode(path)
					πF.SetLineno(359)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreprunicode, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdecode_path.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 342: """
			πF.SetLineno(342)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Ensure `path` is Unicode. Return `nodes.reprunicode` object.\n\n    Decode file/path string in a failsave manner if not already done.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdecode_path); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 362: def extract_name_value(line):
			πF.SetLineno(362)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "line", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("extract_name_value", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µline *πg.Object = πArgs[0]
				_ = µline
				var µattlist *πg.Object = πg.UnboundLocal
				_ = µattlist
				var µequals *πg.Object = πg.UnboundLocal
				_ = µequals
				var µattname *πg.Object = πg.UnboundLocal
				_ = µattname
				var µendquote *πg.Object = πg.UnboundLocal
				_ = µendquote
				var µdata *πg.Object = πg.UnboundLocal
				_ = µdata
				var µspace *πg.Object = πg.UnboundLocal
				_ = µspace
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
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
					case 1:
						goto Label1
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 363: """
					πF.SetLineno(363)
					// line 370: attlist = []
					πF.SetLineno(370)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µattlist = πTemp002
					// line 371: while line:
					πF.SetLineno(371)
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
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 372: equals = line.find('=')
					πF.SetLineno(372)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µline, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µequals = πTemp005
					if πE = πg.CheckLocal(πF, µequals, "equals"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µequals, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 373: if equals == -1:
					πF.SetLineno(373)
				Label4:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("missing \"=\"").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNameValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 374: raise NameValueError('missing "="')
					πF.SetLineno(374)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label5
				Label5:
					// line 375: attname = line[:equals].strip()
					πF.SetLineno(375)
					if πE = πg.CheckLocal(πF, µequals, "equals"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µequals, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µattname = πTemp005
					if πE = πg.CheckLocal(πF, µequals, "equals"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µequals, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µattname); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp006).ToObject()
					πTemp002 = πTemp005
				Label6:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 376: if equals == 0 or not attname:
					πF.SetLineno(376)
				Label7:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("missing attribute name before \"=\"").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNameValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 377: raise NameValueError(
					πF.SetLineno(377)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label8
				Label8:
					// line 379: line = line[equals+1:].lstrip()
					πF.SetLineno(379)
					if πE = πg.CheckLocal(πF, µequals, "equals"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µequals, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp005
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µline); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 380: if not line:
					πF.SetLineno(380)
				Label9:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("missing value after \"%s=\"").ToObject(), µattname); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNameValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 381: raise NameValueError(
					πF.SetLineno(381)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label10
				Label10:
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µline, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πg.NewStr("'\"").ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label11
					}
					goto Label12
					// line 383: if line[0] in '\'"':
					πF.SetLineno(383)
				Label11:
					// line 384: endquote = line.find(line[0], 1)
					πF.SetLineno(384)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µline, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µendquote = πTemp005
					if πE = πg.CheckLocal(πF, µendquote, "endquote"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µendquote, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label14
					}
					goto Label15
					// line 385: if endquote == -1:
					πF.SetLineno(385)
				Label14:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
						continue
					}
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µline, πTemp007); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µattname, πTemp008).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("attribute \"%s\" missing end quote (%s)").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNameValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 386: raise NameValueError(
					πF.SetLineno(386)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label15
				Label15:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp001[0] = µline
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µendquote, "endquote"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µendquote, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, πTemp008, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µendquote, "endquote"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µendquote, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp007
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µline, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp007
				Label16:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label17
					}
					goto Label18
					// line 389: if len(line) > endquote + 1 and line[endquote + 1].strip():
					πF.SetLineno(389)
				Label17:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
						continue
					}
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µline, πTemp007); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µattname, πTemp008).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("attribute \"%s\" end quote (%s) not followed by whitespace").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNameValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 390: raise NameValueError(
					πF.SetLineno(390)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label18
				Label18:
					// line 393: data = line[1:endquote]
					πF.SetLineno(393)
					if πE = πg.CheckLocal(πF, µendquote, "endquote"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), µendquote, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					µdata = πTemp005
					// line 394: line = line[endquote+1:].lstrip()
					πF.SetLineno(394)
					if πE = πg.CheckLocal(πF, µendquote, "endquote"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µendquote, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp005
					goto Label13
				Label12:
					// line 396: space = line.find(' ')
					πF.SetLineno(396)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(" ").ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µline, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µspace = πTemp005
					if πE = πg.CheckLocal(πF, µspace, "space"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µspace, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label19
					}
					goto Label20
					// line 397: if space == -1:
					πF.SetLineno(397)
				Label19:
					// line 398: data = line
					πF.SetLineno(398)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					µdata = µline
					// line 399: line = ''
					πF.SetLineno(399)
					µline = ß.ToObject()
					goto Label21
				Label20:
					// line 401: data = line[:space]
					πF.SetLineno(401)
					if πE = πg.CheckLocal(πF, µspace, "space"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µspace, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					µdata = πTemp005
					// line 402: line = line[space+1:].lstrip()
					πF.SetLineno(402)
					if πE = πg.CheckLocal(πF, µspace, "space"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µspace, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßlstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp005
					goto Label21
				Label21:
					goto Label13
				Label13:
					// line 403: attlist.append((attname.lower(), data))
					πF.SetLineno(403)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µattname, ßlower, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp007, µdata).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µattlist, "attlist"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µattlist, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 404: return attlist
					πF.SetLineno(404)
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
			if πE = πF.Globals().SetItem(πF, ßextract_name_value.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 363: """
			πF.SetLineno(363)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    Return a list of (name, value) from a line of the form \"name=value ...\".\n\n    :Exception:\n        `NameValueError` for invalid input (missing name, missing data, bad\n        quotes, etc.).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßextract_name_value); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 406: def new_reporter(source_path, settings):
			πF.SetLineno(406)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "source_path", Def: nil}
			πTemp007[1] = πg.Param{Name: "settings", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("new_reporter", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource_path *πg.Object = πArgs[0]
				_ = µsource_path
				var µsettings *πg.Object = πArgs[1]
				_ = µsettings
				var µreporter *πg.Object = πg.UnboundLocal
				_ = µreporter
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
					// line 407: """
					πF.SetLineno(407)
					// line 416: reporter = Reporter(
					πF.SetLineno(416)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp001[0] = µsource_path
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings, ßreport_level, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings, ßhalt_level, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings, ßwarning_stream, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsettings, ßdebug, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsettings, ßerror_encoding, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsettings, ßerror_encoding_error_handler, nil); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"stream", πTemp002},
						{"debug", πTemp003},
						{"encoding", πTemp004},
						{"error_handler", πTemp005},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßReporter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µreporter = πTemp003
					// line 421: return reporter
					πF.SetLineno(421)
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					πR = µreporter
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnew_reporter.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 407: """
			πF.SetLineno(407)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n    Return a new Reporter object.\n\n    :Parameters:\n        `source` : string\n            The path to or description of the source text of the document.\n        `settings` : optparse.Values object\n            Runtime settings.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßnew_reporter); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 423: def new_document(source_path, settings=None):
			πF.SetLineno(423)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "source_path", Def: nil}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[1] = πg.Param{Name: "settings", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("new_document", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource_path *πg.Object = πArgs[0]
				_ = µsource_path
				var µsettings *πg.Object = πArgs[1]
				_ = µsettings
				var µfrontend *πg.Object = πg.UnboundLocal
				_ = µfrontend
				var µreporter *πg.Object = πg.UnboundLocal
				_ = µreporter
				var µdocument *πg.Object = πg.UnboundLocal
				_ = µdocument
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 424: """
					πF.SetLineno(424)
					// line 441: from docutils import frontend
					πF.SetLineno(441)
					if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[1]
					µfrontend = πTemp001
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µsettings == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 442: if settings is None:
					πF.SetLineno(442)
				Label1:
					// line 443: settings = frontend.OptionParser().get_default_values()
					πF.SetLineno(443)
					if πE = πg.CheckLocal(πF, µfrontend, "frontend"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfrontend, ßOptionParser, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßget_default_values, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsettings = πTemp003
					goto Label2
				Label2:
					// line 444: source_path = decode_path(source_path)
					πF.SetLineno(444)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp002[0] = µsource_path
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdecode_path); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µsource_path = πTemp003
					// line 445: reporter = new_reporter(source_path, settings)
					πF.SetLineno(445)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp002[0] = µsource_path
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp002[1] = µsettings
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnew_reporter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µreporter = πTemp003
					// line 446: document = nodes.document(settings, reporter, source=source_path)
					πF.SetLineno(446)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp002[0] = µsettings
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					πTemp002[1] = µreporter
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"source", µsource_path},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp002, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdocument = πTemp001
					// line 447: document.note_source(source_path, -1)
					πF.SetLineno(447)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp002[0] = µsource_path
					if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdocument, ßnote_source, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 448: return document
					πF.SetLineno(448)
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					πR = µdocument
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnew_document.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 424: """
			πF.SetLineno(424)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n    Return a new empty document object.\n\n    :Parameters:\n        `source_path` : string\n            The path to or description of the source text of the document.\n        `settings` : optparse.Values object\n            Runtime settings.  If none are provided, a default core set will\n            be used.  If you will use the document object with any Docutils\n            components, you must provide their default settings as well.  For\n            example, if parsing, at least provide the parser settings,\n            obtainable as follows::\n\n                settings = docutils.frontend.OptionParser(\n                    components=(docutils.parsers.rst.Parser,)\n                    ).get_default_values()\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßnew_document); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 450: def clean_rcs_keywords(paragraph, keyword_substitutions):
			πF.SetLineno(450)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "paragraph", Def: nil}
			πTemp007[1] = πg.Param{Name: "keyword_substitutions", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("clean_rcs_keywords", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µparagraph *πg.Object = πArgs[0]
				_ = µparagraph
				var µkeyword_substitutions *πg.Object = πArgs[1]
				_ = µkeyword_substitutions
				var µtextnode *πg.Object = πg.UnboundLocal
				_ = µtextnode
				var µpattern *πg.Object = πg.UnboundLocal
				_ = µpattern
				var µsubstitution *πg.Object = πg.UnboundLocal
				_ = µsubstitution
				var µmatch *πg.Object = πg.UnboundLocal
				_ = µmatch
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
					if πE = πg.CheckLocal(πF, µparagraph, "paragraph"); πE != nil {
						continue
					}
					πTemp004[0] = µparagraph
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp004 = πF.MakeArgs(2)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µparagraph, "paragraph"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µparagraph, πTemp003); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
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
					πTemp001 = πTemp005
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 451: if len(paragraph) == 1 and isinstance(paragraph[0], nodes.Text):
					πF.SetLineno(451)
				Label2:
					// line 452: textnode = paragraph[0]
					πF.SetLineno(452)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µparagraph, "paragraph"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µparagraph, πTemp001); πE != nil {
						continue
					}
					µtextnode = πTemp003
					if πE = πg.CheckLocal(πF, µkeyword_substitutions, "keyword_substitutions"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µkeyword_substitutions); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µpattern = πTemp005
						µsubstitution = πTemp006
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 454: match = pattern.search(textnode)
					πF.SetLineno(454)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtextnode, "textnode"); πE != nil {
						continue
					}
					πTemp004[0] = µtextnode
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpattern, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µmatch = πTemp005
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µmatch); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label7
					}
					goto Label8
					// line 455: if match:
					πF.SetLineno(455)
				Label7:
					// line 456: paragraph[0] = nodes.Text(pattern.sub(substitution, textnode))
					πF.SetLineno(456)
					πTemp004 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsubstitution, "substitution"); πE != nil {
						continue
					}
					πTemp008[0] = µsubstitution
					if πE = πg.CheckLocal(πF, µtextnode, "textnode"); πE != nil {
						continue
					}
					πTemp008[1] = µtextnode
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpattern, ßsub, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp004[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparagraph, "paragraph"); πE != nil {
						continue
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.SetItem(πF, µparagraph, πTemp006, πTemp005); πE != nil {
						continue
					}
					// line 457: return
					πF.SetLineno(457)
					πR = πg.None
					continue
					goto Label8
				Label8:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
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
			if πE = πF.Globals().SetItem(πF, ßclean_rcs_keywords.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 459: def relative_path(source, target):
			πF.SetLineno(459)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "source", Def: nil}
			πTemp007[1] = πg.Param{Name: "target", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("relative_path", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]
				_ = µsource
				var µtarget *πg.Object = πArgs[1]
				_ = µtarget
				var µsource_parts *πg.Object = πg.UnboundLocal
				_ = µsource_parts
				var µtarget_parts *πg.Object = πg.UnboundLocal
				_ = µtarget_parts
				var µparts *πg.Object = πg.UnboundLocal
				_ = µparts
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
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
				var πTemp010 bool
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
					// line 460: """
					πF.SetLineno(460)
					// line 465: source_parts = os.path.abspath(source or type(target)('dummy_file')
					πF.SetLineno(465)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsep, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp002 = µsource
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßdummy_file.ToObject()
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp007[0] = µtarget
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp003, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002 = πTemp003
				Label1:
					πTemp004[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsource_parts = πTemp003
					// line 467: target_parts = os.path.abspath(target).split(os.sep)
					πF.SetLineno(467)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsep, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
						continue
					}
					πTemp004[0] = µtarget
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtarget_parts = πTemp003
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_parts, "source_parts"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µsource_parts, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µtarget_parts, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					goto Label3
					// line 469: if source_parts[:2] != target_parts[:2]:
					πF.SetLineno(469)
				Label2:
					// line 472: return '/'.join(target_parts)
					πF.SetLineno(472)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					πTemp001[0] = µtarget_parts
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label3
				Label3:
					// line 473: source_parts.reverse()
					πF.SetLineno(473)
					if πE = πg.CheckLocal(πF, µsource_parts, "source_parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource_parts, ßreverse, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 474: target_parts.reverse()
					πF.SetLineno(474)
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtarget_parts, ßreverse, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 475: while (source_parts and target_parts
					πF.SetLineno(475)
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
					if πE = πg.CheckLocal(πF, µsource_parts, "source_parts"); πE != nil {
						continue
					}
					πTemp002 = µsource_parts
					if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					πTemp002 = µtarget_parts
					if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label7
					}
					if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp008 = πTemp009
					if πE = πg.CheckLocal(πF, µsource_parts, "source_parts"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µsource_parts, πTemp008); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp008 = πTemp012
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µtarget_parts, πTemp008); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp009, πTemp012); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label7:
					if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 478: source_parts.pop()
					πF.SetLineno(478)
					if πE = πg.CheckLocal(πF, µsource_parts, "source_parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource_parts, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 479: target_parts.pop()
					πF.SetLineno(479)
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtarget_parts, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 480: target_parts.reverse()
					πF.SetLineno(480)
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtarget_parts, ßreverse, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 481: parts = ['..'] * (len(source_parts) - 1) + target_parts
					πF.SetLineno(481)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("..").ToObject()
					πTemp008 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource_parts, "source_parts"); πE != nil {
						continue
					}
					πTemp001[0] = µsource_parts
					if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.Sub(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtarget_parts, "target_parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, µtarget_parts); πE != nil {
						continue
					}
					µparts = πTemp002
					// line 482: return '/'.join(parts)
					πF.SetLineno(482)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πTemp001[0] = µparts
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßrelative_path.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 460: """
			πF.SetLineno(460)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n    Build and return a path to `target`, relative to `source` (both files).\n\n    If there is no common prefix, return the absolute path to `target`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßrelative_path); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 484: def get_stylesheet_reference(settings, relative_to=None):
			πF.SetLineno(484)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "settings", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[1] = πg.Param{Name: "relative_to", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("get_stylesheet_reference", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsettings *πg.Object = πArgs[0]
				_ = µsettings
				var µrelative_to *πg.Object = πArgs[1]
				_ = µrelative_to
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 485: """
					πF.SetLineno(485)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsettings, ßstylesheet_path, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 492: if settings.stylesheet_path:
					πF.SetLineno(492)
				Label1:
					// line 493: assert not settings.stylesheet, (
					πF.SetLineno(493)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsettings, ßstylesheet, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πE = πg.Assert(πF, πTemp001, πg.NewStr("stylesheet and stylesheet_path are mutually exclusive.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrelative_to, "relative_to"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µrelative_to, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 495: if relative_to == None:
					πF.SetLineno(495)
				Label4:
					// line 496: relative_to = settings._destination
					πF.SetLineno(496)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsettings, ß_destination, nil); πE != nil {
						continue
					}
					µrelative_to = πTemp001
					goto Label5
				Label5:
					// line 497: return relative_path(relative_to, settings.stylesheet_path)
					πF.SetLineno(497)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrelative_to, "relative_to"); πE != nil {
						continue
					}
					πTemp004[0] = µrelative_to
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsettings, ßstylesheet_path, nil); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrelative_path); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp003
					continue
					goto Label3
				Label2:
					// line 499: return settings.stylesheet
					πF.SetLineno(499)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsettings, ßstylesheet, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_stylesheet_reference.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 485: """
			πF.SetLineno(485)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n    Retrieve a stylesheet reference from the settings object.\n\n    Deprecated. Use get_stylesheet_list() instead to\n    enable specification of multiple stylesheets as a comma-separated\n    list.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßget_stylesheet_reference); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 511: def get_stylesheet_list(settings):
			πF.SetLineno(511)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "settings", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("get_stylesheet_list", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsettings *πg.Object = πArgs[0]
				_ = µsettings
				var µstylesheets *πg.Object = πg.UnboundLocal
				_ = µstylesheets
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
					// line 512: """
					πF.SetLineno(512)
					// line 515: assert not (settings.stylesheet and settings.stylesheet_path), (
					πF.SetLineno(515)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsettings, ßstylesheet, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp004
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsettings, ßstylesheet_path, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp004
				Label1:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πE = πg.Assert(πF, πTemp001, πg.NewStr("stylesheet and stylesheet_path are mutually exclusive.").ToObject()); πE != nil {
						continue
					}
					// line 517: stylesheets = settings.stylesheet_path or settings.stylesheet or []
					πF.SetLineno(517)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings, ßstylesheet_path, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings, ßstylesheet, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					πTemp005 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					πTemp001 = πTemp002
				Label2:
					µstylesheets = πTemp001
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstylesheets, "stylesheets"); πE != nil {
						continue
					}
					πTemp005[0] = µstylesheets
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					πTemp005[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 519: if not isinstance(stylesheets, list):
					πF.SetLineno(519)
				Label3:
					// line 520: stylesheets = [path.strip() for path in stylesheets.split(',')]
					πF.SetLineno(520)
					πTemp006 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								πTemp002[0] = πg.NewStr(",").ToObject()
								if πE = πg.CheckLocal(πF, µstylesheets, "stylesheets"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µstylesheets, ßsplit, nil); πE != nil {
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
									µpath = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 520: stylesheets = [path.strip() for path in stylesheets.split(',')]
								πF.SetLineno(520)
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µpath, ßstrip, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					µstylesheets = πTemp001
					goto Label4
				Label4:
					// line 522: return [find_file_in_dirs(path, settings.stylesheet_dirs)
					πF.SetLineno(522)
					πTemp006 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µpath *πg.Object = πg.UnboundLocal
						_ = µpath
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
								if πE = πg.CheckLocal(πF, µstylesheets, "stylesheets"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µstylesheets); πE != nil {
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
									µpath = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 522: return [find_file_in_dirs(path, settings.stylesheet_dirs)
								πF.SetLineno(522)
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								πTemp005[0] = µpath
								if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µsettings, ßstylesheet_dirs, nil); πE != nil {
									continue
								}
								πTemp005[1] = πTemp004
								if πTemp004, πE = πg.ResolveGlobal(πF, ßfind_file_in_dirs); πE != nil {
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
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_stylesheet_list.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 512: """
			πF.SetLineno(512)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("\n    Retrieve list of stylesheet references from the settings object.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßget_stylesheet_list); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
				continue
			}
			// line 525: def find_file_in_dirs(path, dirs):
			πF.SetLineno(525)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "path", Def: nil}
			πTemp007[1] = πg.Param{Name: "dirs", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("find_file_in_dirs", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpath *πg.Object = πArgs[0]
				_ = µpath
				var µdirs *πg.Object = πArgs[1]
				_ = µdirs
				var µd *πg.Object = πg.UnboundLocal
				_ = µd
				var µf *πg.Object = πg.UnboundLocal
				_ = µf
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
					// line 526: """
					πF.SetLineno(526)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisabs, nil); πE != nil {
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
					// line 531: if os.path.isabs(path):
					πF.SetLineno(531)
				Label1:
					// line 532: return path
					πF.SetLineno(532)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πR = µpath
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µdirs, "dirs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µdirs); πE != nil {
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
						µd = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(3)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µd, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 534: if d == '.':
					πF.SetLineno(534)
				Label6:
					// line 535: f = path
					πF.SetLineno(535)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					µf = µpath
					goto Label8
				Label7:
					// line 537: d = os.path.expanduser(d)
					πF.SetLineno(537)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[0] = µd
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßexpanduser, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µd = πTemp006
					// line 538: f = os.path.join(d, path)
					πF.SetLineno(538)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[0] = µd
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[1] = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp006
					goto Label8
				Label8:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πTemp001[0] = µf
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßexists, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label9
					}
					goto Label10
					// line 539: if os.path.exists(f):
					πF.SetLineno(539)
				Label9:
					// line 540: return f
					πF.SetLineno(540)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					πR = µf
					continue
					goto Label10
				Label10:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 541: return path
					πF.SetLineno(541)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πR = µpath
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfind_file_in_dirs.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 526: """
			πF.SetLineno(526)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("\n    Search for `path` in the list of directories `dirs`.\n\n    Return the first expansion that matches an existing file.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßfind_file_in_dirs); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
			// line 543: def get_trim_footnote_ref_space(settings):
			πF.SetLineno(543)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "settings", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("get_trim_footnote_ref_space", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsettings *πg.Object = πArgs[0]
				_ = µsettings
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
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 544: """
					πF.SetLineno(544)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings, ßtrim_footnote_reference_space, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 552: if settings.trim_footnote_reference_space is None:
					πF.SetLineno(552)
				Label1:
					// line 553: return hasattr(settings, 'footnote_references') and \
					πF.SetLineno(553)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp005[0] = µsettings
					πTemp005[1] = ßfootnote_references.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsettings, ßfootnote_references, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp003, ßsuperscript.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label4:
					πR = πTemp001
					continue
					goto Label3
				Label2:
					// line 556: return settings.trim_footnote_reference_space
					πF.SetLineno(556)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsettings, ßtrim_footnote_reference_space, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_trim_footnote_ref_space.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 544: """
			πF.SetLineno(544)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("\n    Return whether or not to trim footnote space.\n\n    If trim_footnote_reference_space is not None, return it.\n\n    If trim_footnote_reference_space is None, return False unless the\n    footnote reference style is 'superscript'.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßget_trim_footnote_ref_space); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
				continue
			}
			// line 558: def get_source_line(node):
			πF.SetLineno(558)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "node", Def: nil}
			πTemp018 = πg.NewFunction(πg.NewCode("get_source_line", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnode *πg.Object = πArgs[0]
				_ = µnode
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 bool
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
					case 1:
						goto Label1
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 559: """
					πF.SetLineno(559)
					// line 563: while node:
					πF.SetLineno(563)
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
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µnode); πE != nil {
						continue
					}
					if πE != nil || !πTemp002 {
						continue
					}
					πF.PushCheckpoint(1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µnode, ßsource, nil); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µnode, ßline, nil); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label4:
					if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 564: if node.source or node.line:
					πF.SetLineno(564)
				Label5:
					// line 565: return node.source, node.line
					πF.SetLineno(565)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µnode, ßsource, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µnode, ßline, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πR = πTemp003
					continue
					goto Label6
				Label6:
					// line 566: node = node.parent
					πF.SetLineno(566)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 567: return None, None
					πF.SetLineno(567)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßget_source_line.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 559: """
			πF.SetLineno(559)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("\n    Return the \"source\" and \"line\" attributes from the `node` given or from\n    its closest ancestor.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßget_source_line); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
				continue
			}
			// line 569: def escape2null(text):
			πF.SetLineno(569)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("escape2null", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µparts *πg.Object = πg.UnboundLocal
				_ = µparts
				var µstart *πg.Object = πg.UnboundLocal
				_ = µstart
				var µfound *πg.Object = πg.UnboundLocal
				_ = µfound
				var πTemp001 []*πg.Object
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
					// line 570: """Return a string with escape-backslashes converted to nulls."""
					πF.SetLineno(570)
					// line 571: parts = []
					πF.SetLineno(571)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µparts = πTemp002
					// line 572: start = 0
					πF.SetLineno(572)
					µstart = πg.NewInt(0).ToObject()
					// line 573: while True:
					πF.SetLineno(573)
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
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 574: found = text.find('\\', start)
					πF.SetLineno(574)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("\\").ToObject()
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					πTemp001[1] = µstart
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtext, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfound = πTemp005
					if πE = πg.CheckLocal(πF, µfound, "found"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µfound, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 575: if found == -1:
					πF.SetLineno(575)
				Label4:
					// line 576: parts.append(text[start:])
					πF.SetLineno(576)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µtext, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 577: return ''.join(parts)
					πF.SetLineno(577)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πTemp001[0] = µparts
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
					goto Label5
				Label5:
					// line 578: parts.append(text[start:found])
					πF.SetLineno(578)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfound, "found"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, µfound, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µtext, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 579: parts.append('\x00' + text[found+1:found+2])
					πF.SetLineno(579)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfound, "found"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µfound, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfound, "found"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µfound, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πTemp007, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µtext, πTemp005); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πg.NewStr("\x00").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 580: start = found + 2               # skip character after escape
					πF.SetLineno(580)
					if πE = πg.CheckLocal(πF, µfound, "found"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µfound, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					µstart = πTemp002
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
			if πE = πF.Globals().SetItem(πF, ßescape2null.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 570: """Return a string with escape-backslashes converted to nulls."""
			πF.SetLineno(570)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Return a string with escape-backslashes converted to nulls.").ToObject()); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßescape2null); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
				continue
			}
			// line 584: def split_escaped_whitespace(text):
			πF.SetLineno(584)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("split_escaped_whitespace", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µstrings *πg.Object = πg.UnboundLocal
				_ = µstrings
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
					// line 585: """
					πF.SetLineno(585)
					// line 589: strings = text.split('\x00 ')
					πF.SetLineno(589)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\x00 ").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtext, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µstrings = πTemp003
					// line 590: strings = [string.split('\x00\n') for string in strings]
					πF.SetLineno(590)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µstring *πg.Object = πg.UnboundLocal
						_ = µstring
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
								if πE = πg.CheckLocal(πF, µstrings, "strings"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µstrings); πE != nil {
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
									µstring = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 590: strings = [string.split('\x00\n') for string in strings]
								πF.SetLineno(590)
								πTemp005 = πF.MakeArgs(1)
								πTemp005[0] = πg.NewStr("\x00\n").ToObject()
								if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µstring, ßsplit, nil); πE != nil {
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
					µstrings = πTemp002
					// line 592: return list(itertools.chain(*strings))
					πF.SetLineno(592)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstrings, "strings"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßchain, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp005, nil, µstrings, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsplit_escaped_whitespace.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 585: """
			πF.SetLineno(585)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("\n    Split `text` on escaped whitespace (null+space or null+newline).\n    Return a list of strings.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßsplit_escaped_whitespace); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
				continue
			}
			// line 594: def strip_combining_chars(text):
			πF.SetLineno(594)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("strip_combining_chars", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
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
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp003[0] = µtext
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
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
					if !πTemp002 {
						goto Label1
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp004, πE = πg.LT(πF, πTemp006, πTemp005); πE != nil {
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
					// line 595: if isinstance(text, str) and sys.version_info < (3, 0):
					πF.SetLineno(595)
				Label2:
					// line 596: return text
					πF.SetLineno(596)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
					goto Label3
				Label3:
					// line 597: return u''.join([c for c in text if not unicodedata.combining(c)])
					πF.SetLineno(597)
					πTemp003 = πF.MakeArgs(1)
					πTemp007 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µtext); πE != nil {
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
									µc = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp005[0] = µc
								if πTemp006, πE = πg.ResolveGlobal(πF, ßunicodedata); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßcombining, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
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
								// line 597: return u''.join([c for c in text if not unicodedata.combining(c)])
								πF.SetLineno(597)
							Label4:
								// line 597: return u''.join([c for c in text if not unicodedata.combining(c)])
								πF.SetLineno(597)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µc, nil
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
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.GetAttr(πF, πg.NewUnicode("").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp005
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstrip_combining_chars.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 599: def find_combining_chars(text):
			πF.SetLineno(599)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("find_combining_chars", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
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
					// line 600: """Return indices of all combining chars in  Unicode string `text`.
					πF.SetLineno(600)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp003[0] = µtext
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
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
					if !πTemp002 {
						goto Label1
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp004, πE = πg.LT(πF, πTemp006, πTemp005); πE != nil {
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
					// line 607: if isinstance(text, str) and sys.version_info < (3, 0):
					πF.SetLineno(607)
				Label2:
					// line 608: return []
					πF.SetLineno(608)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label3:
					// line 609: return [i for i,c in enumerate(text) if unicodedata.combining(c)]
					πF.SetLineno(609)
					πTemp007 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
								case 6:
									goto Label6
								default:
									panic("unexpected function state")
								}
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								πTemp002[0] = µtext
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
									µc = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)
								πTemp002 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp002[0] = µc
								if πTemp003, πE = πg.ResolveGlobal(πF, ßunicodedata); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcombining, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 609: return [i for i,c in enumerate(text) if unicodedata.combining(c)]
								πF.SetLineno(609)
							Label4:
								// line 609: return [i for i,c in enumerate(text) if unicodedata.combining(c)]
								πF.SetLineno(609)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µi, nil
							Label6:
								πTemp003 = πSent
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
			if πE = πF.Globals().SetItem(πF, ßfind_combining_chars.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 600: """Return indices of all combining chars in  Unicode string `text`.
			πF.SetLineno(600)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("Return indices of all combining chars in  Unicode string `text`.\n\n    >>> from docutils.utils import find_combining_chars\n    >>> find_combining_chars(u'A t\xcc\x86ab\xcc\x86le\xcc\x86')\n    [3, 6, 9]\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp024, πE = πg.ResolveGlobal(πF, ßfind_combining_chars); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
				continue
			}
			// line 611: def column_indices(text):
			πF.SetLineno(611)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			πTemp023 = πg.NewFunction(πg.NewCode("column_indices", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µstring_indices *πg.Object = πg.UnboundLocal
				_ = µstring_indices
				var µindex *πg.Object = πg.UnboundLocal
				_ = µindex
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []πg.Param
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
					// line 612: """Indices of Unicode string `text` when skipping combining characters.
					πF.SetLineno(612)
					// line 621: string_indices = list(range(len(text)))
					πF.SetLineno(621)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp003[0] = µtext
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µstring_indices = πTemp005
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[0] = µtext
					if πTemp005, πE = πg.ResolveGlobal(πF, ßfind_combining_chars); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Iter(πF, πTemp006); πE != nil {
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
						µindex = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 623: string_indices[index] = None
					πF.SetLineno(623)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstring_indices, "string_indices"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					πTemp009 = µindex
					if πE = πg.SetItem(πF, µstring_indices, πTemp009, πTemp006); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 624: return [i for i in string_indices if i is not None]
					πF.SetLineno(624)
					πTemp010 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								case 6:
									goto Label6
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µstring_indices, "string_indices"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µstring_indices); πE != nil {
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
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(µi != πTemp005).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								goto Label5
								// line 624: return [i for i in string_indices if i is not None]
								πF.SetLineno(624)
							Label4:
								// line 624: return [i for i in string_indices if i is not None]
								πF.SetLineno(624)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µi, nil
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
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßcolumn_indices.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 612: """Indices of Unicode string `text` when skipping combining characters.
			πF.SetLineno(612)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("Indices of Unicode string `text` when skipping combining characters.\n\n    >>> from docutils.utils import column_indices\n    >>> column_indices(u'A t\xcc\x86ab\xcc\x86le\xcc\x86')\n    [0, 1, 2, 4, 5, 7, 8]\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ßcolumn_indices); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
				continue
			}
			// line 626: east_asian_widths = {'W': 2,   # Wide
			πF.SetLineno(626)
			πTemp006 = πg.NewDict()
			if πE = πTemp006.SetItem(πF, ßW.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßF.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßNa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßH.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßN.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßA.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp024 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ßeast_asian_widths.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 633: """Mapping of result codes from `unicodedata.east_asian_widt()` to character
			πF.SetLineno(633)
			// line 636: def column_width(text):
			πF.SetLineno(636)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "text", Def: nil}
			πTemp024 = πg.NewFunction(πg.NewCode("column_width", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µwidth *πg.Object = πg.UnboundLocal
				_ = µwidth
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
					// line 637: """Return the column width of text.
					πF.SetLineno(637)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp003[0] = µtext
					if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
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
					if !πTemp002 {
						goto Label1
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp004, πE = πg.LT(πF, πTemp006, πTemp005); πE != nil {
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
					// line 641: if isinstance(text, str) and sys.version_info < (3, 0):
					πF.SetLineno(641)
				Label2:
					// line 642: return len(text)
					πF.SetLineno(642)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp003[0] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label3
				Label3:
					// line 643: width = sum([east_asian_widths[unicodedata.east_asian_width(c)]
					πF.SetLineno(643)
					πTemp003 = πF.MakeArgs(1)
					πTemp007 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µtext); πE != nil {
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
									µc = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 643: width = sum([east_asian_widths[unicodedata.east_asian_width(c)]
								πF.SetLineno(643)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								πTemp005[0] = µc
								if πTemp006, πE = πg.ResolveGlobal(πF, ßunicodedata); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßeast_asian_width, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp004 = πTemp006
								if πTemp007, πE = πg.ResolveGlobal(πF, ßeast_asian_widths); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
									continue
								}
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
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µwidth = πTemp005
					// line 646: width -= len(find_combining_chars(text))
					πF.SetLineno(646)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp008[0] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ßfind_combining_chars); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.ISub(πF, µwidth, πTemp005); πE != nil {
						continue
					}
					µwidth = πTemp001
					// line 647: return width
					πF.SetLineno(647)
					if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
						continue
					}
					πR = µwidth
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcolumn_width.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 637: """Return the column width of text.
			πF.SetLineno(637)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("Return the column width of text.\n\n    Correct ``len(text)`` for wide East Asian and combining Unicode chars.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp026, πE = πg.ResolveGlobal(πF, ßcolumn_width); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
				continue
			}
			// line 649: def uniq(L):
			πF.SetLineno(649)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "L", Def: nil}
			πTemp025 = πg.NewFunction(πg.NewCode("uniq", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µL *πg.Object = πArgs[0]
				_ = µL
				var µr *πg.Object = πg.UnboundLocal
				_ = µr
				var µitem *πg.Object = πg.UnboundLocal
				_ = µitem
				var πTemp001 []*πg.Object
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
					// line 650: r = []
					πF.SetLineno(650)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µr = πTemp002
					if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µL); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µitem = πTemp005
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µr, µitem); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 652: if not item in r:
					πF.SetLineno(652)
				Label4:
					// line 653: r.append(item)
					πF.SetLineno(653)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µr, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
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
					// line 654: return r
					πF.SetLineno(654)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πR = µr
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuniq.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 656: def unique_combinations(items, n):
			πF.SetLineno(656)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "items", Def: nil}
			πTemp007[1] = πg.Param{Name: "n", Def: nil}
			πTemp026 = πg.NewFunction(πg.NewCode("unique_combinations", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µitems *πg.Object = πArgs[0]
				_ = µitems
				var µn *πg.Object = πArgs[1]
				_ = µn
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
					// line 657: """Return `itertools.combinations`."""
					πF.SetLineno(657)
					// line 658: warnings.warn('docutils.utils.unique_combinations is deprecated; '
					πF.SetLineno(658)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("docutils.utils.unique_combinations is deprecated; use itertools.combinations directly.").ToObject()
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
					// line 661: return itertools.combinations(items, n)
					πF.SetLineno(661)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp001[0] = µitems
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp001[1] = µn
					if πTemp002, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcombinations, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßunique_combinations.ToObject(), πTemp026); πE != nil {
				continue
			}
			// line 657: """Return `itertools.combinations`."""
			πF.SetLineno(657)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Return `itertools.combinations`.").ToObject()); πE != nil {
				continue
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßunique_combinations); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
				continue
			}
			// line 663: def normalize_language_tag(tag):
			πF.SetLineno(663)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "tag", Def: nil}
			πTemp027 = πg.NewFunction(πg.NewCode("normalize_language_tag", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtag *πg.Object = πArgs[0]
				_ = µtag
				var µsubtags *πg.Object = πg.UnboundLocal
				_ = µsubtags
				var µbase_tag *πg.Object = πg.UnboundLocal
				_ = µbase_tag
				var µtaglist *πg.Object = πg.UnboundLocal
				_ = µtaglist
				var µn *πg.Object = πg.UnboundLocal
				_ = µn
				var µtags *πg.Object = πg.UnboundLocal
				_ = µtags
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
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
					// line 664: """Return a list of normalized combinations for a `BCP 47` language tag.
					πF.SetLineno(664)
					// line 676: tag = tag.lower().replace('-', '_')
					πF.SetLineno(676)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("-").ToObject()
					πTemp001[1] = ß_.ToObject()
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtag, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtag = πTemp003
					// line 678: tag = re.sub(r'_([a-zA-Z0-9])_', r'_\1-', tag)
					πF.SetLineno(678)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("_([a-zA-Z0-9])_").ToObject()
					πTemp001[1] = πg.NewStr("_\\1-").ToObject()
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp001[2] = µtag
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtag = πTemp002
					// line 679: subtags = [subtag for subtag in tag.split('_')]
					πF.SetLineno(679)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µsubtag *πg.Object = πg.UnboundLocal
						_ = µsubtag
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
								πTemp002[0] = ß_.ToObject()
								if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µtag, ßsplit, nil); πE != nil {
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
									µsubtag = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 679: subtags = [subtag for subtag in tag.split('_')]
								πF.SetLineno(679)
								if πE = πg.CheckLocal(πF, µsubtag, "subtag"); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return µsubtag, nil
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
					µsubtags = πTemp002
					// line 680: base_tag = (subtags.pop(0),)
					πF.SetLineno(680)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsubtags, "subtags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsubtags, ßpop, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple1(πTemp006).ToObject()
					µbase_tag = πTemp002
					// line 682: taglist = []
					πF.SetLineno(682)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µtaglist = πTemp002
					πTemp001 = πF.MakeArgs(3)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsubtags, "subtags"); πE != nil {
						continue
					}
					πTemp007[0] = µsubtags
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp001[0] = πTemp006
					πTemp001[1] = πg.NewInt(0).ToObject()
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[2] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp008 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						µn = πTemp005
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(1)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsubtags, "subtags"); πE != nil {
						continue
					}
					πTemp001[0] = µsubtags
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp001[1] = µn
					if πTemp006, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp006, ßcombinations, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp009 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp006, πE = πg.Next(πF, πTemp005); πE != nil {
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
						µtags = πTemp006
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 686: taglist.append('-'.join(base_tag+tags))
					πF.SetLineno(686)
					πTemp001 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbase_tag, "base_tag"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtags, "tags"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µbase_tag, µtags); πE != nil {
						continue
					}
					πTemp007[0] = πTemp006
					if πTemp006, πE = πg.GetAttr(πF, πg.NewStr("-").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp001[0] = πTemp010
					if πE = πg.CheckLocal(πF, µtaglist, "taglist"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µtaglist, ßappend, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
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
					// line 687: taglist += base_tag
					πF.SetLineno(687)
					if πE = πg.CheckLocal(πF, µtaglist, "taglist"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbase_tag, "base_tag"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µtaglist, µbase_tag); πE != nil {
						continue
					}
					µtaglist = πTemp002
					// line 688: return taglist
					πF.SetLineno(688)
					if πE = πg.CheckLocal(πF, µtaglist, "taglist"); πE != nil {
						continue
					}
					πR = µtaglist
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnormalize_language_tag.ToObject(), πTemp027); πE != nil {
				continue
			}
			// line 664: """Return a list of normalized combinations for a `BCP 47` language tag.
			πF.SetLineno(664)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("Return a list of normalized combinations for a `BCP 47` language tag.\n\n    Example:\n\n    >>> from docutils.utils import normalize_language_tag\n    >>> normalize_language_tag('de_AT-1901')\n    ['de-at-1901', 'de-at', 'de-1901', 'de']\n    >>> normalize_language_tag('de-CH-x_altquot')\n    ['de-ch-x-altquot', 'de-ch', 'de-x-altquot', 'de']\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp029, πE = πg.ResolveGlobal(πF, ßnormalize_language_tag); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
				continue
			}
			// line 691: class DependencyList(object):
			πF.SetLineno(691)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp030, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp030
			πTemp006 = πg.NewDict()
			if πTemp028, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp028); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DependencyList", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 693: """
					πF.SetLineno(693)
					// line 693: """
					πF.SetLineno(693)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    List of dependencies, with file recording support.\n\n    Note that the output file is not automatically closed.  You have\n    to explicitly call the close() method.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 700: def __init__(self, output_file=None, dependencies=[]):
					πF.SetLineno(700)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "output_file", Def: πTemp003}
					πTemp004 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					πTemp002[2] = πg.Param{Name: "dependencies", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µoutput_file *πg.Object = πArgs[1]
						_ = µoutput_file
						var µdependencies *πg.Object = πArgs[2]
						_ = µdependencies
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
							// line 701: """
							πF.SetLineno(701)
							// line 706: self.set_output(output_file)
							πF.SetLineno(706)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoutput_file, "output_file"); πE != nil {
								continue
							}
							πTemp001[0] = µoutput_file
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_output, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µdependencies, "dependencies"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µdependencies); πE != nil {
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
								µi = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 708: self.add(i)
							πF.SetLineno(708)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßadd, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 701: """
					πF.SetLineno(701)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Initialize the dependency list, automatically setting the\n        output file to `output_file` (see `set_output()`) and adding\n        all supplied dependencies.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 710: def set_output(self, output_file):
					πF.SetLineno(710)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "output_file", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("set_output", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µoutput_file *πg.Object = πArgs[1]
						_ = µoutput_file
						var µof *πg.Object = πg.UnboundLocal
						_ = µof
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 711: """
							πF.SetLineno(711)
							// line 721: self.list = []
							πF.SetLineno(721)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlist, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoutput_file, "output_file"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µoutput_file); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 722: if output_file:
							πF.SetLineno(722)
						Label1:
							if πE = πg.CheckLocal(πF, µoutput_file, "output_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µoutput_file, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 723: if output_file == '-':
							πF.SetLineno(723)
						Label4:
							// line 724: of = None
							πF.SetLineno(724)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µof = πTemp002
							goto Label6
						Label5:
							// line 726: of = output_file
							πF.SetLineno(726)
							if πE = πg.CheckLocal(πF, µoutput_file, "output_file"); πE != nil {
								continue
							}
							µof = µoutput_file
							goto Label6
						Label6:
							// line 727: self.file = docutils.io.FileOutput(destination_path=of,
							πF.SetLineno(727)
							if πE = πg.CheckLocal(πF, µof, "of"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"destination_path", µof},
								{"encoding", ßutf8.ToObject()},
								{"autoclose", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßio, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßFileOutput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfile, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 730: self.file = None
							πF.SetLineno(730)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfile, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_output.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 711: """
					πF.SetLineno(711)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Set the output file and clear the list of already added\n        dependencies.\n\n        `output_file` must be a string.  The specified file is\n        immediately overwritten.\n\n        If output_file is '-', the output will be written to stdout.\n        If it is None, no file output is done when calling add().\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßset_output); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 732: def add(self, *filenames):
					πF.SetLineno(732)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("add", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfilenames *πg.Object = πArgs[1]
						_ = µfilenames
						var µfilename *πg.Object = πg.UnboundLocal
						_ = µfilename
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
						var πTemp007 []*πg.Object
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
							// line 733: """
							πF.SetLineno(733)
							if πE = πg.CheckLocal(πF, µfilenames, "filenames"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µfilenames); πE != nil {
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
								µfilename = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßlist, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp006, µfilename); πE != nil {
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
							// line 739: if not filename in self.list:
							πF.SetLineno(739)
						Label4:
							// line 740: self.list.append(filename)
							πF.SetLineno(740)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp007[0] = µfilename
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßlist, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfile, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp005 != πTemp006).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 741: if self.file is not None:
							πF.SetLineno(741)
						Label6:
							// line 742: self.file.write(filename+'\n')
							πF.SetLineno(742)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µfilename, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfile, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ßadd.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 733: """
					πF.SetLineno(733)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        If the dependency `filename` has not already been added,\n        append it to self.list and print it to self.file if self.file\n        is not None.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßadd); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 744: def close(self):
					πF.SetLineno(744)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 745: """
							πF.SetLineno(745)
							// line 748: self.file.close()
							πF.SetLineno(748)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfile, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 749: self.file = None
							πF.SetLineno(749)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfile, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 745: """
					πF.SetLineno(745)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Close the output file.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßclose); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 751: def __repr__(self):
					πF.SetLineno(751)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µoutput_file *πg.Object = πg.UnboundLocal
						_ = µoutput_file
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
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 752: try:
							πF.SetLineno(752)
							πF.PushCheckpoint(2)
							// line 753: output_file = self.file.name
							πF.SetLineno(753)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfile, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßname, nil); πE != nil {
								continue
							}
							µoutput_file = πTemp002
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
							// line 754: except AttributeError:
							πF.SetLineno(754)
						Label3:
							// line 755: output_file = None
							πF.SetLineno(755)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µoutput_file = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 756: return '%s(%r, %s)' % (self.__class__.__name__, output_file, self.list)
							πF.SetLineno(756)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoutput_file, "output_file"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßlist, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp007, µoutput_file, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%r, %s)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp029, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp029 == nil {
				πTemp029 = πg.TypeType.ToObject()
			}
			if πTemp030, πE = πTemp029.Call(πF, []*πg.Object{πg.NewStr("DependencyList").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDependencyList.ToObject(), πTemp030); πE != nil {
				continue
			}
			// line 759: release_level_abbreviations = {
			πF.SetLineno(759)
			πTemp006 = πg.NewDict()
			if πE = πTemp006.SetItem(πF, ßalpha.ToObject(), ßa.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßbeta.ToObject(), ßb.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßcandidate.ToObject(), ßrc.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßfinal.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			πTemp028 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ßrelease_level_abbreviations.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 765: def version_identifier(version_info=None):
			πF.SetLineno(765)
			πTemp007 = make([]πg.Param, 1)
			if πTemp029, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[0] = πg.Param{Name: "version_info", Def: πTemp029}
			πTemp028 = πg.NewFunction(πg.NewCode("version_identifier", "/usr/lib/python2.7/site-packages/docutils/utils/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µversion_info *πg.Object = πArgs[0]
				_ = µversion_info
				var µmicro *πg.Object = πg.UnboundLocal
				_ = µmicro
				var µreleaselevel *πg.Object = πg.UnboundLocal
				_ = µreleaselevel
				var µserial *πg.Object = πg.UnboundLocal
				_ = µserial
				var µdev *πg.Object = πg.UnboundLocal
				_ = µdev
				var µversion *πg.Object = πg.UnboundLocal
				_ = µversion
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
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 766: """
					πF.SetLineno(766)
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µversion_info == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 773: if version_info is None:
					πF.SetLineno(773)
				Label1:
					// line 774: version_info = __version_info__
					πF.SetLineno(774)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__version_info__); πE != nil {
						continue
					}
					µversion_info = πTemp001
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µversion_info, ßmicro, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label3
					}
					goto Label4
					// line 775: if version_info.micro:
					πF.SetLineno(775)
				Label3:
					// line 776: micro = '.%s' % version_info.micro
					πF.SetLineno(776)
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µversion_info, ßmicro, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr(".%s").ToObject(), πTemp002); πE != nil {
						continue
					}
					µmicro = πTemp001
					goto Label5
				Label4:
					// line 779: micro = ''
					πF.SetLineno(779)
					µmicro = ß.ToObject()
					goto Label5
				Label5:
					// line 780: releaselevel = release_level_abbreviations[version_info.releaselevel]
					πF.SetLineno(780)
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µversion_info, ßreleaselevel, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrelease_level_abbreviations); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µreleaselevel = πTemp002
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µversion_info, ßserial, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label6
					}
					goto Label7
					// line 781: if version_info.serial:
					πF.SetLineno(781)
				Label6:
					// line 782: serial = version_info.serial
					πF.SetLineno(782)
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µversion_info, ßserial, nil); πE != nil {
						continue
					}
					µserial = πTemp001
					goto Label8
				Label7:
					// line 785: serial = ''
					πF.SetLineno(785)
					µserial = ß.ToObject()
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µversion_info, ßrelease, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label9
					}
					goto Label10
					// line 786: if version_info.release:
					πF.SetLineno(786)
				Label9:
					// line 787: dev = ''
					πF.SetLineno(787)
					µdev = ß.ToObject()
					goto Label11
				Label10:
					// line 789: dev = '.dev'
					πF.SetLineno(789)
					µdev = πg.NewStr(".dev").ToObject()
					goto Label11
				Label11:
					// line 790: version = '%s.%s%s%s%s%s' % (
					πF.SetLineno(790)
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µversion_info, ßmajor, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µversion_info, "version_info"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µversion_info, ßminor, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmicro, "micro"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreleaselevel, "releaselevel"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µserial, "serial"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdev, "dev"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple6(πTemp004, πTemp005, µmicro, µreleaselevel, µserial, µdev).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s.%s%s%s%s%s").ToObject(), πTemp002); πE != nil {
						continue
					}
					µversion = πTemp001
					// line 797: return version
					πF.SetLineno(797)
					if πE = πg.CheckLocal(πF, µversion, "version"); πE != nil {
						continue
					}
					πR = µversion
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßversion_identifier.ToObject(), πTemp028); πE != nil {
				continue
			}
			// line 766: """
			πF.SetLineno(766)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πg.NewStr("\n    Return a version identifier string built from `version_info`, a\n    `docutils.VersionInfo` namedtuple instance or compatible tuple. If\n    `version_info` is not provided, by default return a version identifier\n    string based on `docutils.__version_info__` (i.e. the current Docutils\n    version).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp030, πE = πg.ResolveGlobal(πF, ßversion_identifier); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp030, ß__doc__, πTemp029); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("utils", Code)
}
