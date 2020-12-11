package parsers

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßComponent := πg.InternStr("Component")
		ßImportError := πg.InternStr("ImportError")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßParser := πg.InternStr("Parser")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__import__ := πg.InternStr("__import__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_parser_aliases := πg.InternStr("_parser_aliases")
		ßattach_observer := πg.InternStr("attach_observer")
		ßcomponent_type := πg.InternStr("component_type")
		ßconfig_section := πg.InternStr("config_section")
		ßdetach_observer := πg.InternStr("detach_observer")
		ßdocument := πg.InternStr("document")
		ßfinish_parse := πg.InternStr("finish_parse")
		ßget_parser_class := πg.InternStr("get_parser_class")
		ßglobals := πg.InternStr("globals")
		ßinputstring := πg.InternStr("inputstring")
		ßlocals := πg.InternStr("locals")
		ßlower := πg.InternStr("lower")
		ßnote_parse_message := πg.InternStr("note_parse_message")
		ßparse := πg.InternStr("parse")
		ßparser := πg.InternStr("parser")
		ßparsers := πg.InternStr("parsers")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrest := πg.InternStr("rest")
		ßrestructuredtext := πg.InternStr("restructuredtext")
		ßrestx := πg.InternStr("restx")
		ßrst := πg.InternStr("rst")
		ßrtxt := πg.InternStr("rtxt")
		ßsetup_parse := πg.InternStr("setup_parse")
		ßsys := πg.InternStr("sys")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis package contains Docutils parser modules.\n").ToObject()); πE != nil {
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
			// line 12: from docutils import Component
			πF.SetLineno(12)
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
			// line 15: class Parser(Component):
			πF.SetLineno(15)
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
			_, πE = πg.NewCode("Parser", "/usr/lib/python2.7/site-packages/docutils/parsers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 17: component_type = 'parser'
					πF.SetLineno(17)
					// line 17: component_type = 'parser'
					πF.SetLineno(17)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ßparser.ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßcomponent_type.ToObject(), ßparser.ToObject()); πE != nil {
						continue
					}
					// line 18: config_section = 'parsers'
					πF.SetLineno(18)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), ßparsers.ToObject()); πE != nil {
						continue
					}
					// line 20: def parse(self, inputstring, document):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "inputstring", Def: nil}
					πTemp002[2] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/site-packages/docutils/parsers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinputstring *πg.Object = πArgs[1]
						_ = µinputstring
						var µdocument *πg.Object = πArgs[2]
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
							// line 21: """Override to parse `inputstring` into document tree `document`."""
							πF.SetLineno(21)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("subclass must override this method").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 22: raise NotImplementedError('subclass must override this method')
							πF.SetLineno(22)
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
					if πE = πClass.SetItem(πF, ßparse.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 21: """Override to parse `inputstring` into document tree `document`."""
					πF.SetLineno(21)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Override to parse `inputstring` into document tree `document`.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßparse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 24: def setup_parse(self, inputstring, document):
					πF.SetLineno(24)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "inputstring", Def: nil}
					πTemp002[2] = πg.Param{Name: "document", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("setup_parse", "/usr/lib/python2.7/site-packages/docutils/parsers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinputstring *πg.Object = πArgs[1]
						_ = µinputstring
						var µdocument *πg.Object = πArgs[2]
						_ = µdocument
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 25: """Initial parse setup.  Call at start of `self.parse()`."""
							πF.SetLineno(25)
							// line 26: self.inputstring = inputstring
							πF.SetLineno(26)
							if πE = πg.CheckLocal(πF, µinputstring, "inputstring"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinputstring); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinputstring, πTemp001); πE != nil {
								continue
							}
							// line 27: self.document = document
							πF.SetLineno(27)
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
							// line 28: document.reporter.attach_observer(document.note_parse_message)
							πF.SetLineno(28)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßnote_parse_message, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßattach_observer, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetup_parse.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 25: """Initial parse setup.  Call at start of `self.parse()`."""
					πF.SetLineno(25)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Initial parse setup.  Call at start of `self.parse()`.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßsetup_parse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 30: def finish_parse(self):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("finish_parse", "/usr/lib/python2.7/site-packages/docutils/parsers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 31: """Finalize parse details.  Call at end of `self.parse()`."""
							πF.SetLineno(31)
							// line 32: self.document.reporter.detach_observer(
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_parse_message, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßdetach_observer, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfinish_parse.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 31: """Finalize parse details.  Call at end of `self.parse()`."""
					πF.SetLineno(31)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Finalize parse details.  Call at end of `self.parse()`.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßfinish_parse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Parser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßParser.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 36: _parser_aliases = {
			πF.SetLineno(36)
			πTemp004 = πg.NewDict()
			if πE = πTemp004.SetItem(πF, ßrestructuredtext.ToObject(), ßrst.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßrest.ToObject(), ßrst.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßrestx.ToObject(), ßrst.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßrtxt.ToObject(), ßrst.ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_parser_aliases.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 42: def get_parser_class(parser_name):
			πF.SetLineno(42)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "parser_name", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("get_parser_class", "/usr/lib/python2.7/site-packages/docutils/parsers/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µparser_name *πg.Object = πArgs[0]
				_ = µparser_name
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
					// line 43: """Return the Parser class from the `parser_name` module."""
					πF.SetLineno(43)
					// line 44: parser_name = parser_name.lower()
					πF.SetLineno(44)
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µparser_name, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µparser_name = πTemp002
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parser_aliases); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µparser_name); πE != nil {
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
					// line 45: if parser_name in _parser_aliases:
					πF.SetLineno(45)
				Label1:
					// line 46: parser_name = _parser_aliases[parser_name]
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp001 = µparser_name
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_parser_aliases); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µparser_name = πTemp002
					goto Label2
				Label2:
					// line 47: try:
					πF.SetLineno(47)
					πF.PushCheckpoint(4)
					// line 48: module = __import__(parser_name, globals(), locals(), level=1)
					πF.SetLineno(48)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp005[0] = µparser_name
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
					// line 49: except ImportError:
					πF.SetLineno(49)
				Label5:
					// line 50: module = __import__(parser_name, globals(), locals(), level=0)
					πF.SetLineno(50)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µparser_name, "parser_name"); πE != nil {
						continue
					}
					πTemp005[0] = µparser_name
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
					// line 51: return module.Parser
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmodule, ßParser, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_parser_class.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 43: """Return the Parser class from the `parser_name` module."""
			πF.SetLineno(43)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return the Parser class from the `parser_name` module.").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßget_parser_class); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("parsers", Code)
}
