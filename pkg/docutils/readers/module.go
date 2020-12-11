package readers

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßComponent := πg.InternStr("Component")
		ßDecorations := πg.InternStr("Decorations")
		ßExposeInternals := πg.InternStr("ExposeInternals")
		ßImportError := πg.InternStr("ImportError")
		ßNone := πg.InternStr("None")
		ßReReader := πg.InternStr("ReReader")
		ßReader := πg.InternStr("Reader")
		ßStripComments := πg.InternStr("StripComments")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__import__ := πg.InternStr("__import__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_reader_aliases := πg.InternStr("_reader_aliases")
		ßcomponent_type := πg.InternStr("component_type")
		ßconfig_section := πg.InternStr("config_section")
		ßcurrent_line := πg.InternStr("current_line")
		ßcurrent_source := πg.InternStr("current_source")
		ßdocument := πg.InternStr("document")
		ßget_parser_class := πg.InternStr("get_parser_class")
		ßget_reader_class := πg.InternStr("get_reader_class")
		ßget_transforms := πg.InternStr("get_transforms")
		ßglobals := πg.InternStr("globals")
		ßinput := πg.InternStr("input")
		ßlocals := πg.InternStr("locals")
		ßlower := πg.InternStr("lower")
		ßnew_document := πg.InternStr("new_document")
		ßparse := πg.InternStr("parse")
		ßparser := πg.InternStr("parser")
		ßparsers := πg.InternStr("parsers")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßreader := πg.InternStr("reader")
		ßreaders := πg.InternStr("readers")
		ßset_parser := πg.InternStr("set_parser")
		ßsettings := πg.InternStr("settings")
		ßsource := πg.InternStr("source")
		ßsource_path := πg.InternStr("source_path")
		ßsys := πg.InternStr("sys")
		ßuniversal := πg.InternStr("universal")
		ßutils := πg.InternStr("utils")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis package contains Docutils Reader modules.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 11: import sys
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from docutils import utils, parsers, Component
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßparsers.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßComponent); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßComponent.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: from docutils.transforms import universal
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.universal"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßuniversal.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: class Reader(Component):
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßComponent); πE != nil {
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
			_, πE = πg.NewCode("Reader", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 19: """
					πF.SetLineno(19)
					// line 19: """
					πF.SetLineno(19)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class for docutils Readers.\n\n    Each reader module or package must export a subclass also called 'Reader'.\n\n    The two steps of a Reader's responsibility are to read data from the\n    source Input object and parse the data with the Parser object.\n    Call `read()` to process a document.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 29: component_type = 'reader'
					πF.SetLineno(29)
					if πE = πClass.SetItem(πF, ßcomponent_type.ToObject(), ßreader.ToObject()); πE != nil {
						continue
					}
					// line 30: config_section = 'readers'
					πF.SetLineno(30)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), ßreaders.ToObject()); πE != nil {
						continue
					}
					// line 32: def get_transforms(self):
					πF.SetLineno(32)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 33: return Component.get_transforms(self) + [
							πF.SetLineno(33)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßComponent); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget_transforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = make([]*πg.Object, 3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßuniversal); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßDecorations, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßuniversal); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßExposeInternals, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßuniversal); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßStripComments, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp005
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
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
					// line 38: def __init__(self, parser=None, parser_name=None):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "parser", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "parser_name", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µparser *πg.Object = πArgs[1]
						_ = µparser
						var µparser_name *πg.Object = πArgs[2]
						_ = µparser_name
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
							// line 39: """
							πF.SetLineno(39)
							// line 46: self.parser = parser
							πF.SetLineno(46)
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
							// line 47: """A `parsers.Parser` instance shared by all doctrees.  May be left
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µparser == πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
								continue
							}
							πTemp001 = µparser_name
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 50: if parser is None and parser_name:
							πF.SetLineno(50)
						Label2:
							// line 51: self.set_parser(parser_name)
							πF.SetLineno(51)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
								continue
							}
							πTemp005[0] = µparser_name
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_parser, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label3
						Label3:
							// line 53: self.source = None
							πF.SetLineno(53)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp003); πE != nil {
								continue
							}
							// line 54: """`docutils.io` IO object, source of input data."""
							πF.SetLineno(54)
							// line 56: self.input = None
							πF.SetLineno(56)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp003); πE != nil {
								continue
							}
							// line 57: """Raw text input; either a single string or, for more complex cases,
							πF.SetLineno(57)
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
					// line 39: """
					πF.SetLineno(39)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Initialize the Reader instance.\n\n        Several instance attributes are defined with dummy initial values.\n        Subclasses may use these attributes as they wish.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 60: def set_parser(self, parser_name):
					πF.SetLineno(60)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "parser_name", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("set_parser", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µparser_name *πg.Object = πArgs[1]
						_ = µparser_name
						var µparser_class *πg.Object = πg.UnboundLocal
						_ = µparser_class
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
							// line 61: """Set `self.parser` by name."""
							πF.SetLineno(61)
							// line 62: parser_class = parsers.get_parser_class(parser_name)
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
								continue
							}
							πTemp001[0] = µparser_name
							if πTemp002, πE = πg.ResolveGlobal(πF, ßparsers); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_parser_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µparser_class = πTemp002
							// line 63: self.parser = parser_class()
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µparser_class, "parser_class"); πE != nil {
								continue
							}
							if πTemp002, πE = µparser_class.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_parser.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 61: """Set `self.parser` by name."""
					πF.SetLineno(61)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Set `self.parser` by name.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßset_parser); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 65: def read(self, source, parser, settings):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "source", Def: nil}
					πTemp002[2] = πg.Param{Name: "parser", Def: nil}
					πTemp002[3] = πg.Param{Name: "settings", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πArgs[1]
						_ = µsource
						var µparser *πg.Object = πArgs[2]
						_ = µparser
						var µsettings *πg.Object = πArgs[3]
						_ = µsettings
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
							// line 66: self.source = source
							πF.SetLineno(66)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
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
							// line 67: if not self.parser:
							πF.SetLineno(67)
						Label1:
							// line 68: self.parser = parser
							πF.SetLineno(68)
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
							goto Label2
						Label2:
							// line 69: self.settings = settings
							πF.SetLineno(69)
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
							// line 70: self.input = self.source.read()
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßread, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßinput, πTemp002); πE != nil {
								continue
							}
							// line 71: self.parse()
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 72: return self.document
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 74: def parse(self):
					πF.SetLineno(74)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πg.UnboundLocal
						_ = µdocument
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
							// line 75: """Parse `self.input` into a document tree."""
							πF.SetLineno(75)
							// line 76: self.document = document = self.new_document()
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnew_document, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp001); πE != nil {
								continue
							}
							µdocument = πTemp002
							// line 77: self.parser.parse(self.input, document)
							πF.SetLineno(77)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp003[1] = µdocument
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 78: document.current_source = document.current_line = None
							πF.SetLineno(78)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdocument, ßcurrent_source, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdocument, ßcurrent_line, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßparse.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 75: """Parse `self.input` into a document tree."""
					πF.SetLineno(75)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Parse `self.input` into a document tree.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßparse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 80: def new_document(self):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("new_document", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πg.UnboundLocal
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
							// line 81: """Create and return a new empty document tree (root node)."""
							πF.SetLineno(81)
							// line 82: document = utils.new_document(self.source.source_path, self.settings)
							πF.SetLineno(82)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsource_path, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnew_document, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdocument = πTemp002
							// line 83: return document
							πF.SetLineno(83)
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
					if πE = πClass.SetItem(πF, ßnew_document.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 81: """Create and return a new empty document tree (root node)."""
					πF.SetLineno(81)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Create and return a new empty document tree (root node).").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßnew_document); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Reader").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReader.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 86: class ReReader(Reader):
			πF.SetLineno(86)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßReader); πE != nil {
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
			_, πE = πg.NewCode("ReReader", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 88: """
					πF.SetLineno(88)
					// line 88: """
					πF.SetLineno(88)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A reader which rereads an existing document tree (e.g. a\n    deserializer).\n\n    Often used in conjunction with `writers.UnfilteredWriter`.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 95: def get_transforms(self):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 98: return Component.get_transforms(self)
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßComponent); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_transforms, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_transforms.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ReReader").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReReader.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 101: _reader_aliases = {}
			πF.SetLineno(101)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_reader_aliases.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 103: def get_reader_class(reader_name):
			πF.SetLineno(103)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "reader_name", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("get_reader_class", "/usr/lib/python2.7/site-packages/docutils/readers/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µreader_name *πg.Object = πArgs[0]
				_ = µreader_name
				var µmodule *πg.Object = πg.UnboundLocal
				_ = µmodule
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
				var πTemp006 πg.KWArgs
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
					case 4:
						goto Label4
					default:
						panic("unexpected function state")
					}
					// line 104: """Return the Reader class from the `reader_name` module."""
					πF.SetLineno(104)
					// line 105: reader_name = reader_name.lower()
					πF.SetLineno(105)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µreader_name, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µreader_name = πTemp002
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_reader_aliases); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µreader_name); πE != nil {
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
					// line 106: if reader_name in _reader_aliases:
					πF.SetLineno(106)
				Label1:
					// line 107: reader_name = _reader_aliases[reader_name]
					πF.SetLineno(107)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp001 = µreader_name
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_reader_aliases); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µreader_name = πTemp002
					goto Label2
				Label2:
					// line 108: try:
					πF.SetLineno(108)
					πF.PushCheckpoint(4)
					// line 109: module = __import__(reader_name, globals(), locals(), level=1)
					πF.SetLineno(109)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp005[0] = µreader_name
					if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[1] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[2] = πTemp002
					πTemp006 = πg.KWArgs{
						{"level", πg.NewInt(1).ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µmodule = πTemp002
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label5
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 110: except ImportError:
					πF.SetLineno(110)
				Label5:
					// line 111: module = __import__(reader_name, globals(), locals(), level=0)
					πF.SetLineno(111)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µreader_name, "reader_name"); πE != nil {
						continue
					}
					πTemp005[0] = µreader_name
					if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[1] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp005[2] = πTemp002
					πTemp006 = πg.KWArgs{
						{"level", πg.NewInt(0).ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µmodule = πTemp002
					πF.RestoreExc(nil, nil)
					goto Label3
				Label3:
					// line 112: return module.Reader
					πF.SetLineno(112)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmodule, ßReader, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_reader_class.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 104: """Return the Reader class from the `reader_name` module."""
			πF.SetLineno(104)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return the Reader class from the `reader_name` module.").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßget_reader_class); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("readers", Code)
}
