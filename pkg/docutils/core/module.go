package core

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/pprint"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/core.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßBinaryFileOutput := πg.InternStr("BinaryFileOutput")
		ßDocTreeInput := πg.InternStr("DocTreeInput")
		ßErrorOutput := πg.InternStr("ErrorOutput")
		ßErrorString := πg.InternStr("ErrorString")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßFileInput := πg.InternStr("FileInput")
		ßFileOutput := πg.InternStr("FileOutput")
		ßInputError := πg.InternStr("InputError")
		ßNone := πg.InternStr("None")
		ßNullOutput := πg.InternStr("NullOutput")
		ßOptionParser := πg.InternStr("OptionParser")
		ßOutputError := πg.InternStr("OutputError")
		ßPublisher := πg.InternStr("Publisher")
		ßReader := πg.InternStr("Reader")
		ßReporter := πg.InternStr("Reporter")
		ßSettingsSpec := πg.InternStr("SettingsSpec")
		ßStringInput := πg.InternStr("StringInput")
		ßStringOutput := πg.InternStr("StringOutput")
		ßSystemExit := πg.InternStr("SystemExit")
		ßSystemMessage := πg.InternStr("SystemMessage")
		ßTransformer := πg.InternStr("Transformer")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß__version_details__ := πg.InternStr("__version_details__")
		ß_destination := πg.InternStr("_destination")
		ß_source := πg.InternStr("_source")
		ß_stderr := πg.InternStr("_stderr")
		ßapplication := πg.InternStr("application")
		ßapplications := πg.InternStr("applications")
		ßapplied := πg.InternStr("applied")
		ßapply_transforms := πg.InternStr("apply_transforms")
		ßargv := πg.InternStr("argv")
		ßascii := πg.InternStr("ascii")
		ßassemble_parts := πg.InternStr("assemble_parts")
		ßbackslashreplace := πg.InternStr("backslashreplace")
		ßcode := πg.InternStr("code")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßcopy := πg.InternStr("copy")
		ßdebugging_dumps := πg.InternStr("debugging_dumps")
		ßdecode := πg.InternStr("decode")
		ßdefault_description := πg.InternStr("default_description")
		ßdefault_usage := πg.InternStr("default_usage")
		ßdestination := πg.InternStr("destination")
		ßdestination_class := πg.InternStr("destination_class")
		ßdetails := πg.InternStr("details")
		ßdoctree := πg.InternStr("doctree")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßdump_internals := πg.InternStr("dump_internals")
		ßdump_pseudo_xml := πg.InternStr("dump_pseudo_xml")
		ßdump_settings := πg.InternStr("dump_settings")
		ßdump_transforms := πg.InternStr("dump_transforms")
		ßencode := πg.InternStr("encode")
		ßend := πg.InternStr("end")
		ßexit := πg.InternStr("exit")
		ßexit_status_level := πg.InternStr("exit_status_level")
		ßfrontend := πg.InternStr("frontend")
		ßget_default_values := πg.InternStr("get_default_values")
		ßget_reader_class := πg.InternStr("get_reader_class")
		ßget_settings := πg.InternStr("get_settings")
		ßget_writer_class := πg.InternStr("get_writer_class")
		ßgetattr := πg.InternStr("getattr")
		ßinput_encoding := πg.InternStr("input_encoding")
		ßio := πg.InternStr("io")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßlevels := πg.InternStr("levels")
		ßlocale_encoding := πg.InternStr("locale_encoding")
		ßmax_level := πg.InternStr("max_level")
		ßnull := πg.InternStr("null")
		ßobject := πg.InternStr("object")
		ßoutput_encoding := πg.InternStr("output_encoding")
		ßoutput_encoding_error_handler := πg.InternStr("output_encoding_error_handler")
		ßparse_args := πg.InternStr("parse_args")
		ßparser := πg.InternStr("parser")
		ßparts := πg.InternStr("parts")
		ßpformat := πg.InternStr("pformat")
		ßpopulate_from_components := πg.InternStr("populate_from_components")
		ßpprint := πg.InternStr("pprint")
		ßprint := πg.InternStr("print")
		ßprocess_command_line := πg.InternStr("process_command_line")
		ßprocess_programmatic_settings := πg.InternStr("process_programmatic_settings")
		ßpseudoxml := πg.InternStr("pseudoxml")
		ßpublish := πg.InternStr("publish")
		ßpublish_cmdline := πg.InternStr("publish_cmdline")
		ßpublish_cmdline_to_binary := πg.InternStr("publish_cmdline_to_binary")
		ßpublish_doctree := πg.InternStr("publish_doctree")
		ßpublish_file := πg.InternStr("publish_file")
		ßpublish_from_doctree := πg.InternStr("publish_from_doctree")
		ßpublish_parts := πg.InternStr("publish_parts")
		ßpublish_programmatically := πg.InternStr("publish_programmatically")
		ßpublish_string := πg.InternStr("publish_string")
		ßraw_unicode_escape := πg.InternStr("raw_unicode_escape")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßreader := πg.InternStr("reader")
		ßreaders := πg.InternStr("readers")
		ßreport_Exception := πg.InternStr("report_Exception")
		ßreport_SystemMessage := πg.InternStr("report_SystemMessage")
		ßreport_UnicodeError := πg.InternStr("report_UnicodeError")
		ßreporter := πg.InternStr("reporter")
		ßrestructuredtext := πg.InternStr("restructuredtext")
		ßset_components := πg.InternStr("set_components")
		ßset_destination := πg.InternStr("set_destination")
		ßset_io := πg.InternStr("set_io")
		ßset_parser := πg.InternStr("set_parser")
		ßset_reader := πg.InternStr("set_reader")
		ßset_source := πg.InternStr("set_source")
		ßset_writer := πg.InternStr("set_writer")
		ßsetdefault := πg.InternStr("setdefault")
		ßsettings := πg.InternStr("settings")
		ßsetup_option_parser := πg.InternStr("setup_option_parser")
		ßsource := πg.InternStr("source")
		ßsource_class := πg.InternStr("source_class")
		ßsplit := πg.InternStr("split")
		ßstandalone := πg.InternStr("standalone")
		ßstart := πg.InternStr("start")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtraceback := πg.InternStr("traceback")
		ßtransformer := πg.InternStr("transformer")
		ßutils := πg.InternStr("utils")
		ßversion := πg.InternStr("version")
		ßversion_info := πg.InternStr("version_info")
		ßwrite := πg.InternStr("write")
		ßwriter := πg.InternStr("writer")
		ßwriters := πg.InternStr("writers")
		ßxmlcharrefreplace := πg.InternStr("xmlcharrefreplace")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nCalling the ``publish_*`` convenience functions (or instantiating a\n`Publisher` object) with component names will result in default\nbehavior.  For custom behavior (setting component options), create\ncustom component objects first, and pass *them* to\n``publish_*``/`Publisher`.  See `The Docutils Publisher`_.\n\n.. _The Docutils Publisher: http://docutils.sf.net/docs/api/publisher.html\n").ToObject()); πE != nil {
				continue
			}
			// line 14: from __future__ import print_function
			πF.SetLineno(14)
			// line 16: __docformat__ = 'reStructuredText'
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 18: import sys
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: import pprint
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "pprint"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßpprint.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: from docutils import __version__, __version_details__, SettingsSpec
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß__version__); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß__version_details__); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__version_details__.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSettingsSpec); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSettingsSpec.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: from docutils import frontend, io, utils, readers, writers
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßfrontend.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.readers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßreaders.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßwriters.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: from docutils.frontend import OptionParser
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßOptionParser); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionParser.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 23: from docutils.transforms import Transformer
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformer); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformer.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: from docutils.utils.error_reporting import ErrorOutput, ErrorString
			πF.SetLineno(24)
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
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorString.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 25: import docutils.readers.doctree
			πF.SetLineno(25)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.readers.doctree"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: class Publisher(object):
			πF.SetLineno(27)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("Publisher", "/usr/lib/python2.7/site-packages/docutils/core.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A facade encapsulating the high-level logic of a Docutils system.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 33: def __init__(self, reader=None, parser=None, writer=None,
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 9)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "reader", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "parser", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "writer", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "source", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßio); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßFileInput, nil); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "source_class", Def: πTemp004}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "destination", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßio); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßFileOutput, nil); πE != nil {
						continue
					}
					πTemp002[7] = πg.Param{Name: "destination_class", Def: πTemp004}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[8] = πg.Param{Name: "settings", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µreader *πg.Object = πArgs[1]
						_ = µreader
						var µparser *πg.Object = πArgs[2]
						_ = µparser
						var µwriter *πg.Object = πArgs[3]
						_ = µwriter
						var µsource *πg.Object = πArgs[4]
						_ = µsource
						var µsource_class *πg.Object = πArgs[5]
						_ = µsource_class
						var µdestination *πg.Object = πArgs[6]
						_ = µdestination
						var µdestination_class *πg.Object = πArgs[7]
						_ = µdestination_class
						var µsettings *πg.Object = πArgs[8]
						_ = µsettings
						var µcomponent *πg.Object = πg.UnboundLocal
						_ = µcomponent
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
							default:
								panic("unexpected function state")
							}
							// line 37: """
							πF.SetLineno(37)
							// line 43: self.document = None
							πF.SetLineno(43)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp002); πE != nil {
								continue
							}
							// line 44: """The document tree (`docutils.nodes` objects)."""
							πF.SetLineno(44)
							// line 46: self.reader = reader
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreader); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreader, πTemp001); πE != nil {
								continue
							}
							// line 47: """A `docutils.readers.Reader` instance."""
							πF.SetLineno(47)
							// line 49: self.parser = parser
							πF.SetLineno(49)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparser); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp001); πE != nil {
								continue
							}
							// line 50: """A `docutils.parsers.Parser` instance."""
							πF.SetLineno(50)
							// line 52: self.writer = writer
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µwriter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwriter, πTemp001); πE != nil {
								continue
							}
							// line 53: """A `docutils.writers.Writer` instance."""
							πF.SetLineno(53)
							πTemp002 = πg.NewTuple3(ßreader.ToObject(), ßparser.ToObject(), ßwriter.ToObject()).ToObject()
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
								µcomponent = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 56: assert not isinstance(getattr(self, component), str), (
							πF.SetLineno(56)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[0] = µself
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							πTemp006[1] = µcomponent
							if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple3(πTemp008, µcomponent, µcomponent).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("passed string \"%s\" as \"%s\" parameter; pass an instance, or use the \"%s_name\" parameter instead (in docutils.core.publish_* convenience functions).").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp009[0] = µself
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							πTemp009[1] = µcomponent
							if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp006[0] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp006[1] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 62: self.source = source
							πF.SetLineno(62)
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
							// line 63: """The source of input data, a `docutils.io.Input` instance."""
							πF.SetLineno(63)
							// line 65: self.source_class = source_class
							πF.SetLineno(65)
							if πE = πg.CheckLocal(πF, µsource_class, "source_class"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource_class); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource_class, πTemp001); πE != nil {
								continue
							}
							// line 66: """The class for dynamically created source objects."""
							πF.SetLineno(66)
							// line 68: self.destination = destination
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdestination); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination, πTemp001); πE != nil {
								continue
							}
							// line 69: """The destination for docutils output, a `docutils.io.Output`
							πF.SetLineno(69)
							// line 72: self.destination_class = destination_class
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µdestination_class, "destination_class"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdestination_class); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination_class, πTemp001); πE != nil {
								continue
							}
							// line 73: """The class for dynamically created destination objects."""
							πF.SetLineno(73)
							// line 75: self.settings = settings
							πF.SetLineno(75)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsettings); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsettings, πTemp001); πE != nil {
								continue
							}
							// line 76: """An object containing Docutils settings as instance attributes.
							πF.SetLineno(76)
							// line 79: self._stderr = ErrorOutput()
							πF.SetLineno(79)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß_stderr, πTemp001); πE != nil {
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
					// line 37: """
					πF.SetLineno(37)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Initial setup.  If any of `reader`, `parser`, or `writer` are not\n        specified, the corresponding ``set_...`` method should be called with\n        a component name (`set_reader` sets the parser as well).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 81: def set_reader(self, reader_name, parser, parser_name):
					πF.SetLineno(81)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "reader_name", Def: nil}
					πTemp002[2] = πg.Param{Name: "parser", Def: nil}
					πTemp002[3] = πg.Param{Name: "parser_name", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("set_reader", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µreader_name *πg.Object = πArgs[1]
						_ = µreader_name
						var µparser *πg.Object = πArgs[2]
						_ = µparser
						var µparser_name *πg.Object = πArgs[3]
						_ = µparser_name
						var µreader_class *πg.Object = πg.UnboundLocal
						_ = µreader_class
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
							// line 82: """Set `self.reader` by name."""
							πF.SetLineno(82)
							// line 83: reader_class = readers.get_reader_class(reader_name)
							πF.SetLineno(83)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
								continue
							}
							πTemp001[0] = µreader_name
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreaders); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_reader_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreader_class = πTemp002
							// line 84: self.reader = reader_class(parser, parser_name)
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							πTemp001[0] = µparser
							if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
								continue
							}
							πTemp001[1] = µparser_name
							if πE = πg.CheckLocal(πF, µreader_class, "reader_class"); πE != nil {
								continue
							}
							if πTemp002, πE = µreader_class.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreader, πTemp003); πE != nil {
								continue
							}
							// line 85: self.parser = self.reader.parser
							πF.SetLineno(85)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparser, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_reader.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 82: """Set `self.reader` by name."""
					πF.SetLineno(82)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Set `self.reader` by name.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßset_reader); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 87: def set_writer(self, writer_name):
					πF.SetLineno(87)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "writer_name", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("set_writer", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µwriter_name *πg.Object = πArgs[1]
						_ = µwriter_name
						var µwriter_class *πg.Object = πg.UnboundLocal
						_ = µwriter_class
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
							// line 88: """Set `self.writer` by name."""
							πF.SetLineno(88)
							// line 89: writer_class = writers.get_writer_class(writer_name)
							πF.SetLineno(89)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
								continue
							}
							πTemp001[0] = µwriter_name
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_writer_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µwriter_class = πTemp002
							// line 90: self.writer = writer_class()
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µwriter_class, "writer_class"); πE != nil {
								continue
							}
							if πTemp002, πE = µwriter_class.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwriter, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_writer.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 88: """Set `self.writer` by name."""
					πF.SetLineno(88)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Set `self.writer` by name.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßset_writer); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 92: def set_components(self, reader_name, parser_name, writer_name):
					πF.SetLineno(92)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "reader_name", Def: nil}
					πTemp002[2] = πg.Param{Name: "parser_name", Def: nil}
					πTemp002[3] = πg.Param{Name: "writer_name", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("set_components", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µreader_name *πg.Object = πArgs[1]
						_ = µreader_name
						var µparser_name *πg.Object = πArgs[2]
						_ = µparser_name
						var µwriter_name *πg.Object = πArgs[3]
						_ = µwriter_name
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
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
							// line 93: if self.reader is None:
							πF.SetLineno(93)
						Label1:
							// line 94: self.set_reader(reader_name, self.parser, parser_name)
							πF.SetLineno(94)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
								continue
							}
							πTemp005[0] = µreader_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
								continue
							}
							πTemp005[2] = µparser_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_reader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 95: if self.parser is None:
							πF.SetLineno(95)
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparser, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 96: if self.reader.parser is None:
							πF.SetLineno(96)
						Label5:
							// line 97: self.reader.set_parser(parser_name)
							πF.SetLineno(97)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
								continue
							}
							πTemp005[0] = µparser_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßset_parser, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label6
						Label6:
							// line 98: self.parser = self.reader.parser
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparser, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
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
								goto Label7
							}
							goto Label8
							// line 99: if self.writer is None:
							πF.SetLineno(99)
						Label7:
							// line 100: self.set_writer(writer_name)
							πF.SetLineno(100)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
								continue
							}
							πTemp005[0] = µwriter_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_writer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßset_components.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 102: def setup_option_parser(self, usage=None, description=None,
					πF.SetLineno(102)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "usage", Def: πTemp007}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "description", Def: πTemp007}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "settings_spec", Def: πTemp007}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "config_section", Def: πTemp007}
					πTemp006 = πg.NewFunction(πg.NewCode("setup_option_parser", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µusage *πg.Object = πArgs[1]
						_ = µusage
						var µdescription *πg.Object = πArgs[2]
						_ = µdescription
						var µsettings_spec *πg.Object = πArgs[3]
						_ = µsettings_spec
						var µconfig_section *πg.Object = πArgs[4]
						_ = µconfig_section
						var µdefaults *πg.Object = πArgs[5]
						_ = µdefaults
						var µparts *πg.Object = πg.UnboundLocal
						_ = µparts
						var µoption_parser *πg.Object = πg.UnboundLocal
						_ = µoption_parser
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
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µconfig_section); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 105: if config_section:
							πF.SetLineno(105)
						Label1:
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µsettings_spec); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp001).ToObject()
							if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label3
							}
							goto Label4
							// line 106: if not settings_spec:
							πF.SetLineno(106)
						Label3:
							// line 107: settings_spec = SettingsSpec()
							πF.SetLineno(107)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSettingsSpec); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsettings_spec = πTemp003
							goto Label4
						Label4:
							// line 108: settings_spec.config_section = config_section
							πF.SetLineno(108)
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µconfig_section); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µsettings_spec, ßconfig_section, πTemp002); πE != nil {
								continue
							}
							// line 109: parts = config_section.split()
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µconfig_section, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µparts = πTemp003
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							πTemp004[0] = µparts
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp001 {
								goto Label5
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µparts, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, ßapplication.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label5:
							if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label6
							}
							goto Label7
							// line 110: if len(parts) > 1 and parts[-1] == 'application':
							πF.SetLineno(110)
						Label6:
							// line 111: settings_spec.config_section_dependencies = ['applications']
							πF.SetLineno(111)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = ßapplications.ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µsettings_spec, ßconfig_section_dependencies, πTemp003); πE != nil {
								continue
							}
							goto Label7
						Label7:
							goto Label2
						Label2:
							// line 113: option_parser = OptionParser(
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp003, πTemp005, πTemp006, µsettings_spec).ToObject()
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"components", πTemp002},
								{"defaults", µdefaults},
								{"read_config_files", πTemp003},
								{"usage", µusage},
								{"description", µdescription},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOptionParser); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							µoption_parser = πTemp003
							// line 117: return option_parser
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							πR = µoption_parser
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetup_option_parser.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 119: def get_settings(self, usage=None, description=None,
					πF.SetLineno(119)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "usage", Def: πTemp008}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "description", Def: πTemp008}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "settings_spec", Def: πTemp008}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "config_section", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("get_settings", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µusage *πg.Object = πArgs[1]
						_ = µusage
						var µdescription *πg.Object = πArgs[2]
						_ = µdescription
						var µsettings_spec *πg.Object = πArgs[3]
						_ = µsettings_spec
						var µconfig_section *πg.Object = πArgs[4]
						_ = µconfig_section
						var µdefaults *πg.Object = πArgs[5]
						_ = µdefaults
						var µoption_parser *πg.Object = πg.UnboundLocal
						_ = µoption_parser
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
							// line 121: """
							πF.SetLineno(121)
							// line 128: option_parser = self.setup_option_parser(
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							πTemp001[0] = µusage
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp001[1] = µdescription
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							πTemp001[2] = µsettings_spec
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							πTemp001[3] = µconfig_section
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_option_parser, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, µdefaults); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µoption_parser = πTemp003
							// line 130: self.settings = option_parser.get_default_values()
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption_parser, ßget_default_values, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßsettings, πTemp002); πE != nil {
								continue
							}
							// line 131: return self.settings
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_settings.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 121: """
					πF.SetLineno(121)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Set and return default settings (overrides in `defaults` dict).\n\n        Set components first (`self.set_reader` & `self.set_writer`).\n        Explicitly setting `self.settings` disables command line option\n        processing from `self.publish()`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßget_settings); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 133: def process_programmatic_settings(self, settings_spec,
					πF.SetLineno(133)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "settings_spec", Def: nil}
					πTemp002[2] = πg.Param{Name: "settings_overrides", Def: nil}
					πTemp002[3] = πg.Param{Name: "config_section", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("process_programmatic_settings", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsettings_spec *πg.Object = πArgs[1]
						_ = µsettings_spec
						var µsettings_overrides *πg.Object = πArgs[2]
						_ = µsettings_overrides
						var µconfig_section *πg.Object = πArgs[3]
						_ = µconfig_section
						var µdefaults *πg.Object = πg.UnboundLocal
						_ = µdefaults
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Dict
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
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
							// line 136: if self.settings is None:
							πF.SetLineno(136)
						Label1:
							// line 137: defaults = (settings_overrides or {}).copy()
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
								continue
							}
							πTemp001 = µsettings_overrides
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							πTemp005 = πg.NewDict()
							πTemp002 = πTemp005.ToObject()
							πTemp001 = πTemp002
						Label3:
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdefaults = πTemp001
							// line 139: defaults.setdefault('traceback', True)
							πF.SetLineno(139)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßtraceback.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdefaults, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 140: self.get_settings(settings_spec=settings_spec,
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"settings_spec", µsettings_spec},
								{"config_section", µconfig_section},
							}
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_settings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, nil, πTemp007, µdefaults); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprocess_programmatic_settings.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 144: def process_command_line(self, argv=None, usage=None, description=None,
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "argv", Def: πTemp010}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "usage", Def: πTemp010}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "description", Def: πTemp010}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "settings_spec", Def: πTemp010}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "config_section", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("process_command_line", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargv *πg.Object = πArgs[1]
						_ = µargv
						var µusage *πg.Object = πArgs[2]
						_ = µusage
						var µdescription *πg.Object = πArgs[3]
						_ = µdescription
						var µsettings_spec *πg.Object = πArgs[4]
						_ = µsettings_spec
						var µconfig_section *πg.Object = πArgs[5]
						_ = µconfig_section
						var µdefaults *πg.Object = πArgs[6]
						_ = µdefaults
						var µoption_parser *πg.Object = πg.UnboundLocal
						_ = µoption_parser
						var µargv_encoding *πg.Object = πg.UnboundLocal
						_ = µargv_encoding
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
							// line 147: """
							πF.SetLineno(147)
							// line 153: option_parser = self.setup_option_parser(
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							πTemp001[0] = µusage
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp001[1] = µdescription
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							πTemp001[2] = µsettings_spec
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							πTemp001[3] = µconfig_section
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_option_parser, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, µdefaults); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µoption_parser = πTemp003
							if πE = πg.CheckLocal(πF, µargv, "argv"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µargv == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 155: if argv is None:
							πF.SetLineno(155)
						Label1:
							// line 156: argv = sys.argv[1:]
							πF.SetLineno(156)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßargv, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							µargv = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp002, πE = πg.LT(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 158: if sys.version_info < (3, 0):
							πF.SetLineno(158)
						Label3:
							// line 160: argv_encoding = (frontend.locale_encoding or 'ascii')
							πF.SetLineno(160)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrontend); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßlocale_encoding, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πTemp002 = ßascii.ToObject()
						Label5:
							µargv_encoding = πTemp002
							// line 161: argv = [a.decode(argv_encoding) for a in argv]
							πF.SetLineno(161)
							πTemp007 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πg.UnboundLocal
								_ = µa
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
										if πE = πg.CheckLocal(πF, µargv, "argv"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µargv); πE != nil {
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
											µa = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 161: argv = [a.decode(argv_encoding) for a in argv]
										πF.SetLineno(161)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µargv_encoding, "argv_encoding"); πE != nil {
											continue
										}
										πTemp005[0] = µargv_encoding
										if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µa, ßdecode, nil); πE != nil {
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
							µargv = πTemp002
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 162: self.settings = option_parser.parse_args(argv)
							πF.SetLineno(162)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargv, "argv"); πE != nil {
								continue
							}
							πTemp001[0] = µargv
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption_parser, ßparse_args, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsettings, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprocess_command_line.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 147: """
					πF.SetLineno(147)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n        Pass an empty list to `argv` to avoid reading `sys.argv` (the\n        default).\n\n        Set components first (`self.set_reader` & `self.set_writer`).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßprocess_command_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 164: def set_io(self, source_path=None, destination_path=None):
					πF.SetLineno(164)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "source_path", Def: πTemp011}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "destination_path", Def: πTemp011}
					πTemp010 = πg.NewFunction(πg.NewCode("set_io", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource_path *πg.Object = πArgs[1]
						_ = µsource_path
						var µdestination_path *πg.Object = πArgs[2]
						_ = µdestination_path
						var πTemp001 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
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
							// line 165: if self.source is None:
							πF.SetLineno(165)
						Label1:
							// line 166: self.set_source(source_path=source_path)
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"source_path", µsource_path},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_source, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 167: if self.destination is None:
							πF.SetLineno(167)
						Label3:
							// line 168: self.set_destination(destination_path=destination_path)
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"destination_path", µdestination_path},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_destination, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_io.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 170: def set_source(self, source=None, source_path=None):
					πF.SetLineno(170)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "source", Def: πTemp012}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "source_path", Def: πTemp012}
					πTemp011 = πg.NewFunction(πg.NewCode("set_source", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πArgs[1]
						_ = µsource
						var µsource_path *πg.Object = πArgs[2]
						_ = µsource_path
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µsource_path == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 171: if source_path is None:
							πF.SetLineno(171)
						Label1:
							// line 172: source_path = self.settings._source
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_source, nil); πE != nil {
								continue
							}
							µsource_path = πTemp002
							goto Label3
						Label2:
							// line 174: self.settings._source = source_path
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource_path); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp002, ß_source, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 177: try:
							πF.SetLineno(177)
							πF.PushCheckpoint(5)
							// line 178: self.source = self.source_class(
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinput_encoding, nil); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"source", µsource},
								{"source_path", µsource_path},
								{"encoding", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 181: except TypeError:
							πF.SetLineno(181)
						Label6:
							// line 182: self.source = self.source_class(
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinput_encoding, nil); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"source", µsource},
								{"source_path", µsource_path},
								{"encoding", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp001); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ßset_source.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 186: def set_destination(self, destination=None, destination_path=None):
					πF.SetLineno(186)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "destination", Def: πTemp013}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "destination_path", Def: πTemp013}
					πTemp012 = πg.NewFunction(πg.NewCode("set_destination", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdestination *πg.Object = πArgs[1]
						_ = µdestination
						var µdestination_path *πg.Object = πArgs[2]
						_ = µdestination_path
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdestination_path == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 187: if destination_path is None:
							πF.SetLineno(187)
						Label1:
							// line 188: destination_path = self.settings._destination
							πF.SetLineno(188)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_destination, nil); πE != nil {
								continue
							}
							µdestination_path = πTemp002
							goto Label3
						Label2:
							// line 190: self.settings._destination = destination_path
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdestination_path); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp002, ß_destination, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 191: self.destination = self.destination_class(
							πF.SetLineno(191)
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßoutput_encoding_error_handler, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"destination", µdestination},
								{"destination_path", µdestination_path},
								{"encoding", πTemp002},
								{"error_handler", πTemp004},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_destination.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 196: def apply_transforms(self):
					πF.SetLineno(196)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("apply_transforms", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 197: self.document.transformer.populate_from_components(
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßparser, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(πTemp003, πTemp004, πTemp006, πTemp005, πTemp007).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtransformer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßpopulate_from_components, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 200: self.document.transformer.apply_transforms()
							πF.SetLineno(200)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtransformer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßapply_transforms, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply_transforms.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 202: def publish(self, argv=None, usage=None, description=None,
					πF.SetLineno(202)
					πTemp002 = make([]πg.Param, 8)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "argv", Def: πTemp015}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "usage", Def: πTemp015}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "description", Def: πTemp015}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "settings_spec", Def: πTemp015}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "settings_overrides", Def: πTemp015}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "config_section", Def: πTemp015}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[7] = πg.Param{Name: "enable_exit_status", Def: πTemp015}
					πTemp014 = πg.NewFunction(πg.NewCode("publish", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargv *πg.Object = πArgs[1]
						_ = µargv
						var µusage *πg.Object = πArgs[2]
						_ = µusage
						var µdescription *πg.Object = πArgs[3]
						_ = µdescription
						var µsettings_spec *πg.Object = πArgs[4]
						_ = µsettings_spec
						var µsettings_overrides *πg.Object = πArgs[5]
						_ = µsettings_overrides
						var µconfig_section *πg.Object = πArgs[6]
						_ = µconfig_section
						var µenable_exit_status *πg.Object = πArgs[7]
						_ = µenable_exit_status
						var µexit *πg.Object = πg.UnboundLocal
						_ = µexit
						var µoutput *πg.Object = πg.UnboundLocal
						_ = µoutput
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µexit_status *πg.Object = πg.UnboundLocal
						_ = µexit_status
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
						var πTemp006 *πg.Dict
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
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
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 205: """
							πF.SetLineno(205)
							// line 210: exit = None
							πF.SetLineno(210)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µexit = πTemp001
							// line 211: try:
							πF.SetLineno(211)
							πF.PushCheckpoint(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 212: if self.settings is None:
							πF.SetLineno(212)
						Label3:
							// line 213: self.process_command_line(
							πF.SetLineno(213)
							πTemp005 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µargv, "argv"); πE != nil {
								continue
							}
							πTemp005[0] = µargv
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							πTemp005[1] = µusage
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp005[2] = µdescription
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							πTemp005[3] = µsettings_spec
							if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
								continue
							}
							πTemp005[4] = µconfig_section
							if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
								continue
							}
							πTemp001 = µsettings_overrides
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πTemp006 = πg.NewDict()
							πTemp002 = πTemp006.ToObject()
							πTemp001 = πTemp002
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprocess_command_line, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp005, nil, nil, πTemp001); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label4
						Label4:
							// line 216: self.set_io()
							πF.SetLineno(216)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_io, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 217: self.document = self.reader.read(self.source, self.parser,
							πF.SetLineno(217)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp002); πE != nil {
								continue
							}
							// line 219: self.apply_transforms()
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßapply_transforms, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 220: output = self.writer.write(self.document, self.destination)
							πF.SetLineno(220)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µoutput = πTemp001
							// line 221: self.writer.assemble_parts()
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßassemble_parts, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 222: except SystemExit as error:
							πF.SetLineno(222)
						Label6:
							µerror = πTemp007.ToObject()
							// line 223: exit = 1
							πF.SetLineno(223)
							µexit = πg.NewInt(1).ToObject()
							// line 224: exit_status = error.code
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerror, ßcode, nil); πE != nil {
								continue
							}
							µexit_status = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label1
							// line 225: except Exception as error:
							πF.SetLineno(225)
						Label7:
							µerror = πTemp007.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 226: if not self.settings:       # exception too early to report nicely
							πF.SetLineno(226)
						Label8:
							// line 227: raise
							πF.SetLineno(227)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtraceback, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 228: if self.settings.traceback: # Propagate exceptions?
							πF.SetLineno(228)
						Label10:
							// line 229: self.debugging_dumps()
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebugging_dumps, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 230: raise
							πF.SetLineno(230)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label11
						Label11:
							// line 231: self.report_Exception(error)
							πF.SetLineno(231)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp005[0] = µerror
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreport_Exception, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 232: exit = True
							πF.SetLineno(232)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µexit = πTemp001
							// line 233: exit_status = 1
							πF.SetLineno(233)
							µexit_status = πg.NewInt(1).ToObject()
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 234: self.debugging_dumps()
							πF.SetLineno(234)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebugging_dumps, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
								continue
							}
							πTemp001 = µenable_exit_status
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp009, ßmax_level, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßexit_status_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, πTemp003, πTemp010); πE != nil {
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
							if πE = πg.CheckLocal(πF, µexit, "exit"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µexit); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 235: if (enable_exit_status and self.document
							πF.SetLineno(235)
						Label13:
							// line 238: sys.exit(self.document.reporter.max_level + 10)
							πF.SetLineno(238)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßmax_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label15
							// line 239: elif exit:
							πF.SetLineno(239)
						Label14:
							// line 240: sys.exit(exit_status)
							πF.SetLineno(240)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexit_status, "exit_status"); πE != nil {
								continue
							}
							πTemp005[0] = µexit_status
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label15
						Label15:
							// line 241: return output
							πF.SetLineno(241)
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							πR = µoutput
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpublish.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 205: """
					πF.SetLineno(205)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n        Process command line options and arguments (if `self.settings` not\n        already set), run `self.reader` and then `self.writer`.  Return\n        `self.writer`'s output.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßpublish); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 243: def debugging_dumps(self):
					πF.SetLineno(243)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("debugging_dumps", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []πg.Param
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
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
							// line 244: if not self.document:
							πF.SetLineno(244)
						Label1:
							// line 245: return
							πF.SetLineno(245)
							πR = πg.None
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdump_settings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 246: if self.settings.dump_settings:
							πF.SetLineno(246)
						Label3:
							// line 247: print('\n::: Runtime settings:', file=self._stderr)
							πF.SetLineno(247)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n::: Runtime settings:").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 248: print(pprint.pformat(self.settings.__dict__), file=self._stderr)
							πF.SetLineno(248)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdump_internals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 249: if self.settings.dump_internals:
							πF.SetLineno(249)
						Label5:
							// line 250: print('\n::: Document internals:', file=self._stderr)
							πF.SetLineno(250)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n::: Document internals:").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 251: print(pprint.pformat(self.document.__dict__), file=self._stderr)
							πF.SetLineno(251)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdump_transforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 252: if self.settings.dump_transforms:
							πF.SetLineno(252)
						Label7:
							// line 253: print('\n::: Transforms applied:', file=self._stderr)
							πF.SetLineno(253)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n::: Transforms applied:").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 254: print(' (priority, transform class, pending node details, '
							πF.SetLineno(254)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr(" (priority, transform class, pending node details, keyword args)").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 256: print(pprint.pformat(
							πF.SetLineno(256)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µpriority *πg.Object = πg.UnboundLocal
								_ = µpriority
								var µxclass *πg.Object = πg.UnboundLocal
								_ = µxclass
								var µpending *πg.Object = πg.UnboundLocal
								_ = µpending
								var µkwargs *πg.Object = πg.UnboundLocal
								_ = µkwargs
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
								var πTemp007 *πg.Object
								_ = πTemp007
								var πTemp008 *πg.Object
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
										case 5:
											goto Label5
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtransformer, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßapplied, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
											if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
												continue
											}
											µpriority = πTemp003
											µxclass = πTemp006
											µpending = πTemp007
											µkwargs = πTemp008
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 257: [(priority, '%s.%s' % (xclass.__module__, xclass.__name__),
										πF.SetLineno(257)
										if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µxclass, "xclass"); πE != nil {
											continue
										}
										if πTemp007, πE = πg.GetAttr(πF, µxclass, ß__module__, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µxclass, "xclass"); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, µxclass, ß__name__, nil); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
										if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s.%s").ToObject(), πTemp006); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
											continue
										}
										πTemp006 = µpending
										if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
											continue
										}
										if !πTemp005 {
											goto Label4
										}
										if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
											continue
										}
										if πTemp007, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
											continue
										}
										πTemp006 = πTemp007
									Label4:
										if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
											continue
										}
										πTemp002 = πg.NewTuple4(µpriority, πTemp003, πTemp006, µkwargs).ToObject()
										πF.PushCheckpoint(5)
										return πTemp002, nil
									Label5:
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
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp001, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp001, ßdump_pseudo_xml, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 261: if self.settings.dump_pseudo_xml:
							πF.SetLineno(261)
						Label9:
							// line 262: print('\n::: Pseudo-XML:', file=self._stderr)
							πF.SetLineno(262)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n::: Pseudo-XML:").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 263: print(self.document.pformat().encode(
							πF.SetLineno(263)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßraw_unicode_escape.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp001, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp001, ßencode, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdebugging_dumps.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 266: def report_Exception(self, error):
					πF.SetLineno(266)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "error", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("report_Exception", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µerror *πg.Object = πArgs[1]
						_ = µerror
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
						var πTemp007 πg.KWArgs
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
						var πR *πg.Object
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
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSystemMessage, nil); πE != nil {
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
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
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
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßInputError, nil); πE != nil {
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
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßOutputError, nil); πE != nil {
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
							goto Label5
							// line 267: if isinstance(error, utils.SystemMessage):
							πF.SetLineno(267)
						Label1:
							// line 268: self.report_SystemMessage(error)
							πF.SetLineno(268)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreport_SystemMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
							// line 269: elif isinstance(error, UnicodeEncodeError):
							πF.SetLineno(269)
						Label2:
							// line 270: self.report_UnicodeError(error)
							πF.SetLineno(270)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreport_UnicodeError, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
							// line 271: elif isinstance(error, io.InputError):
							πF.SetLineno(271)
						Label3:
							// line 272: self._stderr.write(u'Unable to open source file for reading:\n'
							πF.SetLineno(272)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp005[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("Unable to open source file for reading:\n  %s\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
							// line 274: elif isinstance(error, io.OutputError):
							πF.SetLineno(274)
						Label4:
							// line 275: self._stderr.write(
							πF.SetLineno(275)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp005[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("Unable to open destination file for writing:\n  %s\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label5:
							// line 279: print(u'%s' % ErrorString(error), file=self._stderr)
							πF.SetLineno(279)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp005[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("%s").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 280: print(("""\
							πF.SetLineno(280)
							πTemp001 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ß__version__); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ß__version_details__, nil); πE != nil {
								continue
							}
							πTemp009 = πTemp012
							if πTemp010, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label8
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ß__version_details__, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Mod(πF, πg.NewStr(" [%s]").ToObject(), πTemp013); πE != nil {
								continue
							}
							πTemp009 = πTemp011
						Label8:
							πTemp008 = πTemp009
							if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							πTemp008 = ß.ToObject()
						Label7:
							πTemp009 = πg.NewInt(0).ToObject()
							if πTemp012, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßversion, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp013, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp009); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(πTemp006, πTemp008, πTemp011).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Exiting due to error.  Use \"--traceback\" to diagnose.\nPlease report errors to <docutils-users@lists.sf.net>.\nInclude \"--traceback\" output, Docutils version (%s%s),\nPython version (%s), your OS type & version, and the\ncommand line used.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßreport_Exception.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 290: def report_SystemMessage(self, error):
					πF.SetLineno(290)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "error", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("report_SystemMessage", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µerror *πg.Object = πArgs[1]
						_ = µerror
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
							// line 291: print('Exiting due to level-%s (%s) system message.' % (
							πF.SetLineno(291)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µerror, ßlevel, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µerror, ßlevel, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πTemp007, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßReporter, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp008, ßlevels, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Exiting due to level-%s (%s) system message.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp009); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreport_SystemMessage.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 295: def report_UnicodeError(self, error):
					πF.SetLineno(295)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "error", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("report_UnicodeError", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µerror *πg.Object = πArgs[1]
						_ = µerror
						var µdata *πg.Object = πg.UnboundLocal
						_ = µdata
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
							default:
								panic("unexpected function state")
							}
							// line 296: data = error.object[error.start:error.end]
							πF.SetLineno(296)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerror, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µerror, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µerror, ßobject, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µdata = πTemp002
							// line 297: self._stderr.write(
							πF.SetLineno(297)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 7)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp006[0] = µerror
							if πTemp003, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßoutput_encoding, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp007
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßascii.ToObject()
							πTemp006[1] = ßxmlcharrefreplace.ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdata, ßencode, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[2] = πTemp007
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßascii.ToObject()
							πTemp006[1] = ßbackslashreplace.ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdata, ßencode, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[3] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßoutput_encoding_error_handler, nil); πE != nil {
								continue
							}
							πTemp005[4] = πTemp007
							if πTemp003, πE = πg.ResolveGlobal(πF, ß__version__); πE != nil {
								continue
							}
							πTemp005[5] = πTemp003
							πTemp003 = πg.NewInt(0).ToObject()
							if πTemp008, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßversion, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							πTemp005[6] = πTemp007
							πTemp002 = πg.NewTuple(πTemp005...).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s\n\nThe specified output encoding (%s) cannot\nhandle all of the output.\nTry setting \"--output-encoding-error-handler\" to\n\n* \"xmlcharrefreplace\" (for HTML & XML output);\n  the output will contain \"%s\" and should be usable.\n* \"backslashreplace\" (for other output formats);\n  look for \"%s\" in the output.\n* \"replace\"; look for \"?\" in the output.\n\n\"--output-encoding-error-handler\" is currently set to \"%s\".\n\nExiting due to error.  Use \"--traceback\" to diagnose.\nIf the advice above doesn't eliminate the error,\nplease report it to <docutils-users@lists.sf.net>.\nInclude \"--traceback\" output, Docutils version (%s),\nPython version (%s), your OS type & version, and the\ncommand line used.\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreport_UnicodeError.ToObject(), πTemp018); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Publisher").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPublisher.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 325: default_usage = '%prog [options] [<source> [<destination>]]'
			πF.SetLineno(325)
			if πE = πF.Globals().SetItem(πF, ßdefault_usage.ToObject(), πg.NewStr("%prog [options] [<source> [<destination>]]").ToObject()); πE != nil {
				continue
			}
			// line 326: default_description = ('Reads from <source> (default is stdin) and writes to '
			πF.SetLineno(326)
			if πE = πF.Globals().SetItem(πF, ßdefault_description.ToObject(), πg.NewStr("Reads from <source> (default is stdin) and writes to <destination> (default is stdout).  See <http://docutils.sf.net/docs/user/config.html> for the full reference.").ToObject()); πE != nil {
				continue
			}
			// line 331: def publish_cmdline(reader=None, reader_name='standalone',
			πF.SetLineno(331)
			πTemp006 = make([]πg.Param, 14)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "reader", Def: πTemp003}
			πTemp006[1] = πg.Param{Name: "reader_name", Def: ßstandalone.ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "parser", Def: πTemp003}
			πTemp006[3] = πg.Param{Name: "parser_name", Def: ßrestructuredtext.ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "writer", Def: πTemp003}
			πTemp006[5] = πg.Param{Name: "writer_name", Def: ßpseudoxml.ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[6] = πg.Param{Name: "settings", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[7] = πg.Param{Name: "settings_spec", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[8] = πg.Param{Name: "settings_overrides", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[9] = πg.Param{Name: "config_section", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp006[10] = πg.Param{Name: "enable_exit_status", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[11] = πg.Param{Name: "argv", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_usage); πE != nil {
				continue
			}
			πTemp006[12] = πg.Param{Name: "usage", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			πTemp006[13] = πg.Param{Name: "description", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("publish_cmdline", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µreader *πg.Object = πArgs[0]
				_ = µreader
				var µreader_name *πg.Object = πArgs[1]
				_ = µreader_name
				var µparser *πg.Object = πArgs[2]
				_ = µparser
				var µparser_name *πg.Object = πArgs[3]
				_ = µparser_name
				var µwriter *πg.Object = πArgs[4]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[5]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[6]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[7]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[8]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[9]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[10]
				_ = µenable_exit_status
				var µargv *πg.Object = πArgs[11]
				_ = µargv
				var µusage *πg.Object = πArgs[12]
				_ = µusage
				var µdescription *πg.Object = πArgs[13]
				_ = µdescription
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
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
					// line 338: """
					πF.SetLineno(338)
					// line 351: pub = Publisher(reader, parser, writer, settings=settings)
					πF.SetLineno(351)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					πTemp001[0] = µreader
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					πTemp001[1] = µparser
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					πTemp001[2] = µwriter
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"settings", µsettings},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPublisher); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpub = πTemp004
					// line 352: pub.set_components(reader_name, parser_name, writer_name)
					πF.SetLineno(352)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp001[0] = µreader_name
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp001[1] = µparser_name
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp001[2] = µwriter_name
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßset_components, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 353: output = pub.publish(
					πF.SetLineno(353)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µargv, "argv"); πE != nil {
						continue
					}
					πTemp001[0] = µargv
					if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
						continue
					}
					πTemp001[1] = µusage
					if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
						continue
					}
					πTemp001[2] = µdescription
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					πTemp001[3] = µsettings_spec
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					πTemp001[4] = µsettings_overrides
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"config_section", µconfig_section},
						{"enable_exit_status", µenable_exit_status},
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßpublish, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µoutput = πTemp004
					// line 356: return output
					πF.SetLineno(356)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πR = µoutput
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpublish_cmdline.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 338: """
			πF.SetLineno(338)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Set up & run a `Publisher` for command-line-based file I/O (input and\n    output file paths taken automatically from the command line).  Return the\n    encoded string output also.\n\n    Parameters: see `publish_programmatically` for the remainder.\n\n    - `argv`: Command-line argument list to use instead of ``sys.argv[1:]``.\n    - `usage`: Usage string, output if there's a problem parsing the command\n      line.\n    - `description`: Program description, output for the \"--help\" option\n      (along with command-line option descriptions).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßpublish_cmdline); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 358: def publish_file(source=None, source_path=None,
			πF.SetLineno(358)
			πTemp006 = make([]πg.Param, 15)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "source", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "source_path", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "destination", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "destination_path", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "reader", Def: πTemp005}
			πTemp006[5] = πg.Param{Name: "reader_name", Def: ßstandalone.ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[6] = πg.Param{Name: "parser", Def: πTemp005}
			πTemp006[7] = πg.Param{Name: "parser_name", Def: ßrestructuredtext.ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[8] = πg.Param{Name: "writer", Def: πTemp005}
			πTemp006[9] = πg.Param{Name: "writer_name", Def: ßpseudoxml.ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[10] = πg.Param{Name: "settings", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[11] = πg.Param{Name: "settings_spec", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[12] = πg.Param{Name: "settings_overrides", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[13] = πg.Param{Name: "config_section", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[14] = πg.Param{Name: "enable_exit_status", Def: πTemp005}
			πTemp003 = πg.NewFunction(πg.NewCode("publish_file", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]
				_ = µsource
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µdestination *πg.Object = πArgs[2]
				_ = µdestination
				var µdestination_path *πg.Object = πArgs[3]
				_ = µdestination_path
				var µreader *πg.Object = πArgs[4]
				_ = µreader
				var µreader_name *πg.Object = πArgs[5]
				_ = µreader_name
				var µparser *πg.Object = πArgs[6]
				_ = µparser
				var µparser_name *πg.Object = πArgs[7]
				_ = µparser_name
				var µwriter *πg.Object = πArgs[8]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[9]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[10]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[11]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[12]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[13]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[14]
				_ = µenable_exit_status
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var πTemp001 *πg.Object
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
					// line 365: """
					πF.SetLineno(365)
					// line 371: output, pub = publish_programmatically(
					πF.SetLineno(371)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFileInput, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFileOutput, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"source_class", πTemp002},
						{"source", µsource},
						{"source_path", µsource_path},
						{"destination_class", πTemp003},
						{"destination", µdestination},
						{"destination_path", µdestination_path},
						{"reader", µreader},
						{"reader_name", µreader_name},
						{"parser", µparser},
						{"parser_name", µparser_name},
						{"writer", µwriter},
						{"writer_name", µwriter_name},
						{"settings", µsettings},
						{"settings_spec", µsettings_spec},
						{"settings_overrides", µsettings_overrides},
						{"config_section", µconfig_section},
						{"enable_exit_status", µenable_exit_status},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpublish_programmatically); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µoutput = πTemp001
					µpub = πTemp003
					// line 382: return output
					πF.SetLineno(382)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πR = µoutput
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpublish_file.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 365: """
			πF.SetLineno(365)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n    Set up & run a `Publisher` for programmatic use with file-like I/O.\n    Return the encoded string output also.\n\n    Parameters: see `publish_programmatically`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßpublish_file); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 384: def publish_string(source, source_path=None, destination_path=None,
			πF.SetLineno(384)
			πTemp006 = make([]πg.Param, 14)
			πTemp006[0] = πg.Param{Name: "source", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "source_path", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "destination_path", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "reader", Def: πTemp007}
			πTemp006[4] = πg.Param{Name: "reader_name", Def: ßstandalone.ToObject()}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[5] = πg.Param{Name: "parser", Def: πTemp007}
			πTemp006[6] = πg.Param{Name: "parser_name", Def: ßrestructuredtext.ToObject()}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[7] = πg.Param{Name: "writer", Def: πTemp007}
			πTemp006[8] = πg.Param{Name: "writer_name", Def: ßpseudoxml.ToObject()}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[9] = πg.Param{Name: "settings", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[10] = πg.Param{Name: "settings_spec", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[11] = πg.Param{Name: "settings_overrides", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[12] = πg.Param{Name: "config_section", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[13] = πg.Param{Name: "enable_exit_status", Def: πTemp007}
			πTemp005 = πg.NewFunction(πg.NewCode("publish_string", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]
				_ = µsource
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µdestination_path *πg.Object = πArgs[2]
				_ = µdestination_path
				var µreader *πg.Object = πArgs[3]
				_ = µreader
				var µreader_name *πg.Object = πArgs[4]
				_ = µreader_name
				var µparser *πg.Object = πArgs[5]
				_ = µparser
				var µparser_name *πg.Object = πArgs[6]
				_ = µparser_name
				var µwriter *πg.Object = πArgs[7]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[8]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[9]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[10]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[11]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[12]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[13]
				_ = µenable_exit_status
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var πTemp001 *πg.Object
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
					// line 391: """
					πF.SetLineno(391)
					// line 407: output, pub = publish_programmatically(
					πF.SetLineno(407)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßStringInput, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringOutput, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"source_class", πTemp002},
						{"source", µsource},
						{"source_path", µsource_path},
						{"destination_class", πTemp003},
						{"destination", πTemp001},
						{"destination_path", µdestination_path},
						{"reader", µreader},
						{"reader_name", µreader_name},
						{"parser", µparser},
						{"parser_name", µparser_name},
						{"writer", µwriter},
						{"writer_name", µwriter_name},
						{"settings", µsettings},
						{"settings_spec", µsettings_spec},
						{"settings_overrides", µsettings_overrides},
						{"config_section", µconfig_section},
						{"enable_exit_status", µenable_exit_status},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpublish_programmatically); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µoutput = πTemp001
					µpub = πTemp003
					// line 418: return output
					πF.SetLineno(418)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πR = µoutput
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpublish_string.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 391: """
			πF.SetLineno(391)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n    Set up & run a `Publisher` for programmatic use with string I/O.  Return\n    the encoded string or Unicode string output.\n\n    For encoded string output, be sure to set the 'output_encoding' setting to\n    the desired encoding.  Set it to 'unicode' for unencoded Unicode string\n    output.  Here's one way::\n\n        publish_string(..., settings_overrides={'output_encoding': 'unicode'})\n\n    Similarly for Unicode string input (`source`)::\n\n        publish_string(..., settings_overrides={'input_encoding': 'unicode'})\n\n    Parameters: see `publish_programmatically`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßpublish_string); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 420: def publish_parts(source, source_path=None, source_class=io.StringInput,
			πF.SetLineno(420)
			πTemp006 = make([]πg.Param, 15)
			πTemp006[0] = πg.Param{Name: "source", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "source_path", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßStringInput, nil); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "source_class", Def: πTemp009}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "destination_path", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "reader", Def: πTemp008}
			πTemp006[5] = πg.Param{Name: "reader_name", Def: ßstandalone.ToObject()}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[6] = πg.Param{Name: "parser", Def: πTemp008}
			πTemp006[7] = πg.Param{Name: "parser_name", Def: ßrestructuredtext.ToObject()}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[8] = πg.Param{Name: "writer", Def: πTemp008}
			πTemp006[9] = πg.Param{Name: "writer_name", Def: ßpseudoxml.ToObject()}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[10] = πg.Param{Name: "settings", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[11] = πg.Param{Name: "settings_spec", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[12] = πg.Param{Name: "settings_overrides", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[13] = πg.Param{Name: "config_section", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[14] = πg.Param{Name: "enable_exit_status", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("publish_parts", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]
				_ = µsource
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µsource_class *πg.Object = πArgs[2]
				_ = µsource_class
				var µdestination_path *πg.Object = πArgs[3]
				_ = µdestination_path
				var µreader *πg.Object = πArgs[4]
				_ = µreader
				var µreader_name *πg.Object = πArgs[5]
				_ = µreader_name
				var µparser *πg.Object = πArgs[6]
				_ = µparser
				var µparser_name *πg.Object = πArgs[7]
				_ = µparser_name
				var µwriter *πg.Object = πArgs[8]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[9]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[10]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[11]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[12]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[13]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[14]
				_ = µenable_exit_status
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var πTemp001 *πg.Object
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
					// line 428: """
					πF.SetLineno(428)
					// line 441: output, pub = publish_programmatically(
					πF.SetLineno(441)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_class, "source_class"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßStringOutput, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"source", µsource},
						{"source_path", µsource_path},
						{"source_class", µsource_class},
						{"destination_class", πTemp002},
						{"destination", πTemp001},
						{"destination_path", µdestination_path},
						{"reader", µreader},
						{"reader_name", µreader_name},
						{"parser", µparser},
						{"parser_name", µparser_name},
						{"writer", µwriter},
						{"writer_name", µwriter_name},
						{"settings", µsettings},
						{"settings_spec", µsettings_spec},
						{"settings_overrides", µsettings_overrides},
						{"config_section", µconfig_section},
						{"enable_exit_status", µenable_exit_status},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßpublish_programmatically); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
						continue
					}
					µoutput = πTemp001
					µpub = πTemp004
					// line 452: return pub.writer.parts
					πF.SetLineno(452)
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßwriter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparts, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpublish_parts.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 428: """
			πF.SetLineno(428)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n    Set up & run a `Publisher`, and return a dictionary of document parts.\n    Dictionary keys are the names of parts, and values are Unicode strings;\n    encoding is up to the client.  For programmatic use with string I/O.\n\n    For encoded string input, be sure to set the 'input_encoding' setting to\n    the desired encoding.  Set it to 'unicode' for unencoded Unicode string\n    input.  Here's how::\n\n        publish_parts(..., settings_overrides={'input_encoding': 'unicode'})\n\n    Parameters: see `publish_programmatically`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßpublish_parts); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 454: def publish_doctree(source, source_path=None,
			πF.SetLineno(454)
			πTemp006 = make([]πg.Param, 12)
			πTemp006[0] = πg.Param{Name: "source", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "source_path", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßStringInput, nil); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "source_class", Def: πTemp010}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "reader", Def: πTemp009}
			πTemp006[4] = πg.Param{Name: "reader_name", Def: ßstandalone.ToObject()}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[5] = πg.Param{Name: "parser", Def: πTemp009}
			πTemp006[6] = πg.Param{Name: "parser_name", Def: ßrestructuredtext.ToObject()}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[7] = πg.Param{Name: "settings", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[8] = πg.Param{Name: "settings_spec", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[9] = πg.Param{Name: "settings_overrides", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[10] = πg.Param{Name: "config_section", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[11] = πg.Param{Name: "enable_exit_status", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("publish_doctree", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]
				_ = µsource
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µsource_class *πg.Object = πArgs[2]
				_ = µsource_class
				var µreader *πg.Object = πArgs[3]
				_ = µreader
				var µreader_name *πg.Object = πArgs[4]
				_ = µreader_name
				var µparser *πg.Object = πArgs[5]
				_ = µparser
				var µparser_name *πg.Object = πArgs[6]
				_ = µparser_name
				var µsettings *πg.Object = πArgs[7]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[8]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[9]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[10]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[11]
				_ = µenable_exit_status
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
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
					// line 461: """
					πF.SetLineno(461)
					// line 473: pub = Publisher(reader=reader, parser=parser, writer=None,
					πF.SetLineno(473)
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_class, "source_class"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßNullOutput, nil); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"reader", µreader},
						{"parser", µparser},
						{"writer", πTemp001},
						{"settings", µsettings},
						{"source_class", µsource_class},
						{"destination_class", πTemp003},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPublisher); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
						continue
					}
					µpub = πTemp002
					// line 477: pub.set_components(reader_name, parser_name, 'null')
					πF.SetLineno(477)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp005[0] = µreader_name
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp005[1] = µparser_name
					πTemp005[2] = ßnull.ToObject()
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßset_components, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 478: pub.process_programmatic_settings(
					πF.SetLineno(478)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					πTemp005[0] = µsettings_spec
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					πTemp005[1] = µsettings_overrides
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp005[2] = µconfig_section
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßprocess_programmatic_settings, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 480: pub.set_source(source, source_path)
					πF.SetLineno(480)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp005[0] = µsource
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp005[1] = µsource_path
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßset_source, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 481: pub.set_destination(None, None)
					πF.SetLineno(481)
					πTemp005 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßset_destination, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 482: output = pub.publish(enable_exit_status=enable_exit_status)
					πF.SetLineno(482)
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"enable_exit_status", µenable_exit_status},
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßpublish, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
						continue
					}
					µoutput = πTemp002
					// line 483: return pub.document
					πF.SetLineno(483)
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpub, ßdocument, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpublish_doctree.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 461: """
			πF.SetLineno(461)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Set up & run a `Publisher` for programmatic use with string I/O.\n    Return the document tree.\n\n    For encoded string input, be sure to set the 'input_encoding' setting to\n    the desired encoding.  Set it to 'unicode' for unencoded Unicode string\n    input.  Here's one way::\n\n        publish_doctree(..., settings_overrides={'input_encoding': 'unicode'})\n\n    Parameters: see `publish_programmatically`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßpublish_doctree); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 485: def publish_from_doctree(document, destination_path=None,
			πF.SetLineno(485)
			πTemp006 = make([]πg.Param, 9)
			πTemp006[0] = πg.Param{Name: "document", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "destination_path", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "writer", Def: πTemp010}
			πTemp006[3] = πg.Param{Name: "writer_name", Def: ßpseudoxml.ToObject()}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "settings", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[5] = πg.Param{Name: "settings_spec", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[6] = πg.Param{Name: "settings_overrides", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[7] = πg.Param{Name: "config_section", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[8] = πg.Param{Name: "enable_exit_status", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("publish_from_doctree", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdocument *πg.Object = πArgs[0]
				_ = µdocument
				var µdestination_path *πg.Object = πArgs[1]
				_ = µdestination_path
				var µwriter *πg.Object = πArgs[2]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[3]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[4]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[5]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[6]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[7]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[8]
				_ = µenable_exit_status
				var µreader *πg.Object = πg.UnboundLocal
				_ = µreader
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var πTemp001 πg.KWArgs
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					default:
						panic("unexpected function state")
					}
					// line 490: """
					πF.SetLineno(490)
					// line 513: reader = docutils.readers.doctree.Reader(parser_name='null')
					πF.SetLineno(513)
					πTemp001 = πg.KWArgs{
						{"parser_name", ßnull.ToObject()},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreaders, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßdoctree, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßReader, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					µreader = πTemp002
					// line 514: pub = Publisher(reader, None, writer,
					πF.SetLineno(514)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					πTemp004[0] = µreader
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004[1] = πTemp002
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					πTemp004[2] = µwriter
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
						continue
					}
					πTemp005[0] = µdocument
					if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßDocTreeInput, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßStringOutput, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp001 = πg.KWArgs{
						{"source", πTemp002},
						{"destination_class", πTemp006},
						{"settings", µsettings},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPublisher); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µpub = πTemp003
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µwriter); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp002 = µwriter_name
				Label1:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label2
					}
					goto Label3
					// line 517: if not writer and writer_name:
					πF.SetLineno(517)
				Label2:
					// line 518: pub.set_writer(writer_name)
					πF.SetLineno(518)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp004[0] = µwriter_name
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpub, ßset_writer, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label3
				Label3:
					// line 519: pub.process_programmatic_settings(
					πF.SetLineno(519)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					πTemp004[0] = µsettings_spec
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					πTemp004[1] = µsettings_overrides
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp004[2] = µconfig_section
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpub, ßprocess_programmatic_settings, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 521: pub.set_destination(None, destination_path)
					πF.SetLineno(521)
					πTemp004 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					πTemp004[1] = µdestination_path
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpub, ßset_destination, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 522: return pub.publish(enable_exit_status=enable_exit_status)
					πF.SetLineno(522)
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp001 = πg.KWArgs{
						{"enable_exit_status", µenable_exit_status},
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpub, ßpublish, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßpublish_from_doctree.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 490: """
			πF.SetLineno(490)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    Set up & run a `Publisher` to render from an existing document\n    tree data structure, for programmatic use with string I/O.  Return\n    the encoded string output.\n\n    Note that document.settings is overridden; if you want to use the settings\n    of the original `document`, pass settings=document.settings.\n\n    Also, new document.transformer and document.reporter objects are\n    generated.\n\n    For encoded string output, be sure to set the 'output_encoding' setting to\n    the desired encoding.  Set it to 'unicode' for unencoded Unicode string\n    output.  Here's one way::\n\n        publish_from_doctree(\n            ..., settings_overrides={'output_encoding': 'unicode'})\n\n    Parameters: `document` is a `docutils.nodes.document` object, an existing\n    document tree.\n\n    Other parameters: see `publish_programmatically`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßpublish_from_doctree); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 524: def publish_cmdline_to_binary(reader=None, reader_name='standalone',
			πF.SetLineno(524)
			πTemp006 = make([]πg.Param, 16)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[0] = πg.Param{Name: "reader", Def: πTemp011}
			πTemp006[1] = πg.Param{Name: "reader_name", Def: ßstandalone.ToObject()}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "parser", Def: πTemp011}
			πTemp006[3] = πg.Param{Name: "parser_name", Def: ßrestructuredtext.ToObject()}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "writer", Def: πTemp011}
			πTemp006[5] = πg.Param{Name: "writer_name", Def: ßpseudoxml.ToObject()}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[6] = πg.Param{Name: "settings", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[7] = πg.Param{Name: "settings_spec", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[8] = πg.Param{Name: "settings_overrides", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[9] = πg.Param{Name: "config_section", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp006[10] = πg.Param{Name: "enable_exit_status", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[11] = πg.Param{Name: "argv", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßdefault_usage); πE != nil {
				continue
			}
			πTemp006[12] = πg.Param{Name: "usage", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			πTemp006[13] = πg.Param{Name: "description", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[14] = πg.Param{Name: "destination", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßBinaryFileOutput, nil); πE != nil {
				continue
			}
			πTemp006[15] = πg.Param{Name: "destination_class", Def: πTemp012}
			πTemp010 = πg.NewFunction(πg.NewCode("publish_cmdline_to_binary", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µreader *πg.Object = πArgs[0]
				_ = µreader
				var µreader_name *πg.Object = πArgs[1]
				_ = µreader_name
				var µparser *πg.Object = πArgs[2]
				_ = µparser
				var µparser_name *πg.Object = πArgs[3]
				_ = µparser_name
				var µwriter *πg.Object = πArgs[4]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[5]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[6]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[7]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[8]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[9]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[10]
				_ = µenable_exit_status
				var µargv *πg.Object = πArgs[11]
				_ = µargv
				var µusage *πg.Object = πArgs[12]
				_ = µusage
				var µdescription *πg.Object = πArgs[13]
				_ = µdescription
				var µdestination *πg.Object = πArgs[14]
				_ = µdestination
				var µdestination_class *πg.Object = πArgs[15]
				_ = µdestination_class
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
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
					// line 533: """
					πF.SetLineno(533)
					// line 549: pub = Publisher(reader, parser, writer, settings=settings,
					πF.SetLineno(549)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					πTemp001[0] = µreader
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					πTemp001[1] = µparser
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					πTemp001[2] = µwriter
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_class, "destination_class"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"settings", µsettings},
						{"destination_class", µdestination_class},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPublisher); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpub = πTemp004
					// line 551: pub.set_components(reader_name, parser_name, writer_name)
					πF.SetLineno(551)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp001[0] = µreader_name
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp001[1] = µparser_name
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp001[2] = µwriter_name
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßset_components, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 552: output = pub.publish(
					πF.SetLineno(552)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µargv, "argv"); πE != nil {
						continue
					}
					πTemp001[0] = µargv
					if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
						continue
					}
					πTemp001[1] = µusage
					if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
						continue
					}
					πTemp001[2] = µdescription
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					πTemp001[3] = µsettings_spec
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					πTemp001[4] = µsettings_overrides
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"config_section", µconfig_section},
						{"enable_exit_status", µenable_exit_status},
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßpublish, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µoutput = πTemp004
					// line 555: return output
					πF.SetLineno(555)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					πR = µoutput
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpublish_cmdline_to_binary.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 533: """
			πF.SetLineno(533)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n    Set up & run a `Publisher` for command-line-based file I/O (input and\n    output file paths taken automatically from the command line).  Return the\n    encoded string output also.\n\n    This is just like publish_cmdline, except that it uses\n    io.BinaryFileOutput instead of io.FileOutput.\n\n    Parameters: see `publish_programmatically` for the remainder.\n\n    - `argv`: Command-line argument list to use instead of ``sys.argv[1:]``.\n    - `usage`: Usage string, output if there's a problem parsing the command\n      line.\n    - `description`: Program description, output for the \"--help\" option\n      (along with command-line option descriptions).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßpublish_cmdline_to_binary); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 557: def publish_programmatically(source_class, source, source_path,
			πF.SetLineno(557)
			πTemp006 = make([]πg.Param, 17)
			πTemp006[0] = πg.Param{Name: "source_class", Def: nil}
			πTemp006[1] = πg.Param{Name: "source", Def: nil}
			πTemp006[2] = πg.Param{Name: "source_path", Def: nil}
			πTemp006[3] = πg.Param{Name: "destination_class", Def: nil}
			πTemp006[4] = πg.Param{Name: "destination", Def: nil}
			πTemp006[5] = πg.Param{Name: "destination_path", Def: nil}
			πTemp006[6] = πg.Param{Name: "reader", Def: nil}
			πTemp006[7] = πg.Param{Name: "reader_name", Def: nil}
			πTemp006[8] = πg.Param{Name: "parser", Def: nil}
			πTemp006[9] = πg.Param{Name: "parser_name", Def: nil}
			πTemp006[10] = πg.Param{Name: "writer", Def: nil}
			πTemp006[11] = πg.Param{Name: "writer_name", Def: nil}
			πTemp006[12] = πg.Param{Name: "settings", Def: nil}
			πTemp006[13] = πg.Param{Name: "settings_spec", Def: nil}
			πTemp006[14] = πg.Param{Name: "settings_overrides", Def: nil}
			πTemp006[15] = πg.Param{Name: "config_section", Def: nil}
			πTemp006[16] = πg.Param{Name: "enable_exit_status", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("publish_programmatically", "/usr/lib/python2.7/site-packages/docutils/core.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource_class *πg.Object = πArgs[0]
				_ = µsource_class
				var µsource *πg.Object = πArgs[1]
				_ = µsource
				var µsource_path *πg.Object = πArgs[2]
				_ = µsource_path
				var µdestination_class *πg.Object = πArgs[3]
				_ = µdestination_class
				var µdestination *πg.Object = πArgs[4]
				_ = µdestination
				var µdestination_path *πg.Object = πArgs[5]
				_ = µdestination_path
				var µreader *πg.Object = πArgs[6]
				_ = µreader
				var µreader_name *πg.Object = πArgs[7]
				_ = µreader_name
				var µparser *πg.Object = πArgs[8]
				_ = µparser
				var µparser_name *πg.Object = πArgs[9]
				_ = µparser_name
				var µwriter *πg.Object = πArgs[10]
				_ = µwriter
				var µwriter_name *πg.Object = πArgs[11]
				_ = µwriter_name
				var µsettings *πg.Object = πArgs[12]
				_ = µsettings
				var µsettings_spec *πg.Object = πArgs[13]
				_ = µsettings_spec
				var µsettings_overrides *πg.Object = πArgs[14]
				_ = µsettings_overrides
				var µconfig_section *πg.Object = πArgs[15]
				_ = µconfig_section
				var µenable_exit_status *πg.Object = πArgs[16]
				_ = µenable_exit_status
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
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
					// line 565: """
					πF.SetLineno(565)
					// line 657: pub = Publisher(reader, parser, writer, settings=settings,
					πF.SetLineno(657)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader, "reader"); πE != nil {
						continue
					}
					πTemp001[0] = µreader
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					πTemp001[1] = µparser
					if πE = πg.CheckLocal(πF, µwriter, "writer"); πE != nil {
						continue
					}
					πTemp001[2] = µwriter
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_class, "source_class"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_class, "destination_class"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"settings", µsettings},
						{"source_class", µsource_class},
						{"destination_class", µdestination_class},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPublisher); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpub = πTemp004
					// line 660: pub.set_components(reader_name, parser_name, writer_name)
					πF.SetLineno(660)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp001[0] = µreader_name
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp001[1] = µparser_name
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp001[2] = µwriter_name
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßset_components, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 661: pub.process_programmatic_settings(
					πF.SetLineno(661)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					πTemp001[0] = µsettings_spec
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					πTemp001[1] = µsettings_overrides
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp001[2] = µconfig_section
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßprocess_programmatic_settings, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 663: pub.set_source(source, source_path)
					πF.SetLineno(663)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					πTemp001[1] = µsource_path
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßset_source, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 664: pub.set_destination(destination, destination_path)
					πF.SetLineno(664)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
						continue
					}
					πTemp001[0] = µdestination
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					πTemp001[1] = µdestination_path
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßset_destination, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 665: output = pub.publish(enable_exit_status=enable_exit_status)
					πF.SetLineno(665)
					if πE = πg.CheckLocal(πF, µenable_exit_status, "enable_exit_status"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"enable_exit_status", µenable_exit_status},
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßpublish, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, πTemp002); πE != nil {
						continue
					}
					µoutput = πTemp004
					// line 666: return output, pub
					πF.SetLineno(666)
					if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µoutput, µpub).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßpublish_programmatically.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 565: """
			πF.SetLineno(565)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n    Set up & run a `Publisher` for custom programmatic use.  Return the\n    encoded string output and the Publisher object.\n\n    Applications should not need to call this function directly.  If it does\n    seem to be necessary to call this function directly, please write to the\n    Docutils-develop mailing list\n    <http://docutils.sf.net/docs/user/mailing-lists.html#docutils-develop>.\n\n    Parameters:\n\n    * `source_class` **required**: The class for dynamically created source\n      objects.  Typically `io.FileInput` or `io.StringInput`.\n\n    * `source`: Type depends on `source_class`:\n\n      - If `source_class` is `io.FileInput`: Either a file-like object\n        (must have 'read' and 'close' methods), or ``None``\n        (`source_path` is opened).  If neither `source` nor\n        `source_path` are supplied, `sys.stdin` is used.\n\n      - If `source_class` is `io.StringInput` **required**: The input\n        string, either an encoded 8-bit string (set the\n        'input_encoding' setting to the correct encoding) or a Unicode\n        string (set the 'input_encoding' setting to 'unicode').\n\n    * `source_path`: Type depends on `source_class`:\n\n      - `io.FileInput`: Path to the input file, opened if no `source`\n        supplied.\n\n      - `io.StringInput`: Optional.  Path to the file or object that produced\n        `source`.  Only used for diagnostic output.\n\n    * `destination_class` **required**: The class for dynamically created\n      destination objects.  Typically `io.FileOutput` or `io.StringOutput`.\n\n    * `destination`: Type depends on `destination_class`:\n\n      - `io.FileOutput`: Either a file-like object (must have 'write' and\n        'close' methods), or ``None`` (`destination_path` is opened).  If\n        neither `destination` nor `destination_path` are supplied,\n        `sys.stdout` is used.\n\n      - `io.StringOutput`: Not used; pass ``None``.\n\n    * `destination_path`: Type depends on `destination_class`:\n\n      - `io.FileOutput`: Path to the output file.  Opened if no `destination`\n        supplied.\n\n      - `io.StringOutput`: Path to the file or object which will receive the\n        output; optional.  Used for determining relative paths (stylesheets,\n        source links, etc.).\n\n    * `reader`: A `docutils.readers.Reader` object.\n\n    * `reader_name`: Name or alias of the Reader class to be instantiated if\n      no `reader` supplied.\n\n    * `parser`: A `docutils.parsers.Parser` object.\n\n    * `parser_name`: Name or alias of the Parser class to be instantiated if\n      no `parser` supplied.\n\n    * `writer`: A `docutils.writers.Writer` object.\n\n    * `writer_name`: Name or alias of the Writer class to be instantiated if\n      no `writer` supplied.\n\n    * `settings`: A runtime settings (`docutils.frontend.Values`) object, for\n      dotted-attribute access to runtime settings.  It's the end result of the\n      `SettingsSpec`, config file, and option processing.  If `settings` is\n      passed, it's assumed to be complete and no further setting/config/option\n      processing is done.\n\n    * `settings_spec`: A `docutils.SettingsSpec` subclass or object.  Provides\n      extra application-specific settings definitions independently of\n      components.  In other words, the application becomes a component, and\n      its settings data is processed along with that of the other components.\n      Used only if no `settings` specified.\n\n    * `settings_overrides`: A dictionary containing application-specific\n      settings defaults that override the defaults of other components.\n      Used only if no `settings` specified.\n\n    * `config_section`: A string, the name of the configuration file section\n      for this application.  Overrides the ``config_section`` attribute\n      defined by `settings_spec`.  Used only if no `settings` specified.\n\n    * `enable_exit_status`: Boolean; enable exit status at end of processing?\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßpublish_programmatically); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("core", Code)
}
