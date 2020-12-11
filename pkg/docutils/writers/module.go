package writers

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßComponent := πg.InternStr("Component")
		ßFilterMessages := πg.InternStr("FilterMessages")
		ßImportError := πg.InternStr("ImportError")
		ßMessages := πg.InternStr("Messages")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßStripClassesAndElements := πg.InternStr("StripClassesAndElements")
		ßUnfilteredWriter := πg.InternStr("UnfilteredWriter")
		ßWriter := πg.InternStr("Writer")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__import__ := πg.InternStr("__import__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß_writer_aliases := πg.InternStr("_writer_aliases")
		ßassemble_parts := πg.InternStr("assemble_parts")
		ßcomponent_type := πg.InternStr("component_type")
		ßconfig_section := πg.InternStr("config_section")
		ßdestination := πg.InternStr("destination")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßdocutils_xml := πg.InternStr("docutils_xml")
		ßencoding := πg.InternStr("encoding")
		ßget_language := πg.InternStr("get_language")
		ßget_transforms := πg.InternStr("get_transforms")
		ßget_writer_class := πg.InternStr("get_writer_class")
		ßglobals := πg.InternStr("globals")
		ßhtml := πg.InternStr("html")
		ßhtml4 := πg.InternStr("html4")
		ßhtml4css1 := πg.InternStr("html4css1")
		ßhtml5 := πg.InternStr("html5")
		ßhtml5_polyglot := πg.InternStr("html5_polyglot")
		ßlanguage := πg.InternStr("language")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßlatex := πg.InternStr("latex")
		ßlatex2e := πg.InternStr("latex2e")
		ßlocals := πg.InternStr("locals")
		ßlower := πg.InternStr("lower")
		ßos := πg.InternStr("os")
		ßoutput := πg.InternStr("output")
		ßoutput_encoding := πg.InternStr("output_encoding")
		ßparts := πg.InternStr("parts")
		ßpdf := πg.InternStr("pdf")
		ßpformat := πg.InternStr("pformat")
		ßpprint := πg.InternStr("pprint")
		ßpseudoxml := πg.InternStr("pseudoxml")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrlpdf := πg.InternStr("rlpdf")
		ßs5 := πg.InternStr("s5")
		ßs5_html := πg.InternStr("s5_html")
		ßsettings := πg.InternStr("settings")
		ßsys := πg.InternStr("sys")
		ßtranslate := πg.InternStr("translate")
		ßuniversal := πg.InternStr("universal")
		ßversion := πg.InternStr("version")
		ßwhole := πg.InternStr("whole")
		ßwrite := πg.InternStr("write")
		ßwriter := πg.InternStr("writer")
		ßwriters := πg.InternStr("writers")
		ßxelatex := πg.InternStr("xelatex")
		ßxetex := πg.InternStr("xetex")
		ßxhtml := πg.InternStr("xhtml")
		ßxhtml10 := πg.InternStr("xhtml10")
		ßxml := πg.InternStr("xml")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis package contains Docutils Writer modules.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 11: import os.path
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
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
			// line 14: import docutils
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: from docutils import languages, Component
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
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
			// line 16: from docutils.transforms import universal
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.universal"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßuniversal.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: class Writer(Component):
			πF.SetLineno(19)
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
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 21: """
					πF.SetLineno(21)
					// line 21: """
					πF.SetLineno(21)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class for docutils Writers.\n\n    Each writer module or package must export a subclass also called 'Writer'.\n    Each writer must support all standard node types listed in\n    `docutils.nodes.node_class_names`.\n\n    The `write()` method is the main entry point.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 31: component_type = 'writer'
					πF.SetLineno(31)
					if πE = πClass.SetItem(πF, ßcomponent_type.ToObject(), ßwriter.ToObject()); πE != nil {
						continue
					}
					// line 32: config_section = 'writers'
					πF.SetLineno(32)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), ßwriters.ToObject()); πE != nil {
						continue
					}
					// line 34: def get_transforms(self):
					πF.SetLineno(34)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 35: return Component.get_transforms(self) + [
							πF.SetLineno(35)
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
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßMessages, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßuniversal); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßFilterMessages, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßuniversal); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßStripClassesAndElements, nil); πE != nil {
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
					// line 40: document = None
					πF.SetLineno(40)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdocument.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 41: """The document to write (Docutils doctree); set by `write`."""
					πF.SetLineno(41)
					// line 43: output = None
					πF.SetLineno(43)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßoutput.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 44: """Final translated form of `document` (Unicode string for text, binary
					πF.SetLineno(44)
					// line 47: language = None
					πF.SetLineno(47)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßlanguage.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 48: """Language module for the document; set by `write`."""
					πF.SetLineno(48)
					// line 50: destination = None
					πF.SetLineno(50)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdestination.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 51: """`docutils.io` Output object; where to write the document.
					πF.SetLineno(51)
					// line 54: def __init__(self):
					πF.SetLineno(54)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 57: self.parts = {}
							πF.SetLineno(57)
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
							// line 58: """Mapping of document part names to fragments of `self.output`.
							πF.SetLineno(58)
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
					// line 63: def write(self, document, destination):
					πF.SetLineno(63)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp002[2] = πg.Param{Name: "destination", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µdestination *πg.Object = πArgs[2]
						_ = µdestination
						var µoutput *πg.Object = πg.UnboundLocal
						_ = µoutput
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
							// line 64: """
							πF.SetLineno(64)
							// line 73: self.document = document
							πF.SetLineno(73)
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
							// line 74: self.language = languages.get_language(
							πF.SetLineno(74)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlanguage_code, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlanguages); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage, πTemp003); πE != nil {
								continue
							}
							// line 77: self.destination = destination
							πF.SetLineno(77)
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
							// line 78: self.translate()
							πF.SetLineno(78)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtranslate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 79: output = self.destination.write(self.output)
							πF.SetLineno(79)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µoutput = πTemp001
							// line 80: return output
							πF.SetLineno(80)
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
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 64: """
					πF.SetLineno(64)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Process a document into its final form.\n\n        Translate `document` (a Docutils document tree) into the Writer's\n        native format, and write it out to its `destination` (a\n        `docutils.io.Output` subclass object).\n\n        Normally not overridden or extended in subclasses.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßwrite); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 82: def translate(self):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("translate", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 83: """
							πF.SetLineno(83)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("subclass must override this method").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 94: raise NotImplementedError('subclass must override this method')
							πF.SetLineno(94)
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
					if πE = πClass.SetItem(πF, ßtranslate.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 83: """
					πF.SetLineno(83)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Do final translation of `self.document` into `self.output`.  Called\n        from `write`.  Override in subclasses.\n\n        Usually done with a `docutils.nodes.NodeVisitor` subclass, in\n        combination with a call to `docutils.nodes.Node.walk()` or\n        `docutils.nodes.Node.walkabout()`.  The ``NodeVisitor`` subclass must\n        support all standard elements (listed in\n        `docutils.nodes.node_class_names`) and possibly non-standard elements\n        used by the current Reader as well.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtranslate); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 96: def assemble_parts(self):
					πF.SetLineno(96)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("assemble_parts", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 97: """Assemble the `self.parts` dictionary.  Extend in subclasses."""
							πF.SetLineno(97)
							// line 98: self.parts['whole'] = self.output
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparts, nil); πE != nil {
								continue
							}
							πTemp004 = ßwhole.ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 99: self.parts['encoding'] = self.document.settings.output_encoding
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparts, nil); πE != nil {
								continue
							}
							πTemp004 = ßencoding.ToObject()
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 100: self.parts['version'] = docutils.__version__
							πF.SetLineno(100)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__version__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparts, nil); πE != nil {
								continue
							}
							πTemp004 = ßversion.ToObject()
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
					if πE = πClass.SetItem(πF, ßassemble_parts.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 97: """Assemble the `self.parts` dictionary.  Extend in subclasses."""
					πF.SetLineno(97)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Assemble the `self.parts` dictionary.  Extend in subclasses.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßassemble_parts); πE != nil {
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
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 103: class UnfilteredWriter(Writer):
			πF.SetLineno(103)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßWriter); πE != nil {
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
			_, πE = πg.NewCode("UnfilteredWriter", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 105: """
					πF.SetLineno(105)
					// line 105: """
					πF.SetLineno(105)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A writer that passes the document tree on unchanged (e.g. a\n    serializer.)\n\n    Documents written by UnfilteredWriters are typically reused at a\n    later date using a subclass of `readers.ReReader`.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 113: def get_transforms(self):
					πF.SetLineno(113)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 117: return Component.get_transforms(self)
							πF.SetLineno(117)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UnfilteredWriter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnfilteredWriter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 120: _writer_aliases = {
			πF.SetLineno(120)
			πTemp004 = πg.NewDict()
			if πE = πTemp004.SetItem(πF, ßhtml.ToObject(), ßhtml4css1.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßhtml4.ToObject(), ßhtml4css1.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßhtml5.ToObject(), ßhtml5_polyglot.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßlatex.ToObject(), ßlatex2e.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßpprint.ToObject(), ßpseudoxml.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßpformat.ToObject(), ßpseudoxml.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßpdf.ToObject(), ßrlpdf.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßs5.ToObject(), ßs5_html.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßxelatex.ToObject(), ßxetex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßxhtml.ToObject(), ßhtml5_polyglot.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßxhtml10.ToObject(), ßhtml4css1.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßxml.ToObject(), ßdocutils_xml.ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_writer_aliases.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 134: def get_writer_class(writer_name):
			πF.SetLineno(134)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "writer_name", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("get_writer_class", "/usr/lib/python2.7/site-packages/docutils/writers/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µwriter_name *πg.Object = πArgs[0]
				_ = µwriter_name
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
					// line 135: """Return the Writer class from the `writer_name` module."""
					πF.SetLineno(135)
					// line 136: writer_name = writer_name.lower()
					πF.SetLineno(136)
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µwriter_name, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µwriter_name = πTemp002
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_writer_aliases); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µwriter_name); πE != nil {
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
					// line 137: if writer_name in _writer_aliases:
					πF.SetLineno(137)
				Label1:
					// line 138: writer_name = _writer_aliases[writer_name]
					πF.SetLineno(138)
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp001 = µwriter_name
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_writer_aliases); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µwriter_name = πTemp002
					goto Label2
				Label2:
					// line 139: try:
					πF.SetLineno(139)
					πF.PushCheckpoint(4)
					// line 140: module = __import__(writer_name, globals(), locals(), level=1)
					πF.SetLineno(140)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp005[0] = µwriter_name
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
					// line 141: except ImportError:
					πF.SetLineno(141)
				Label5:
					// line 142: module = __import__(writer_name, globals(), locals(), level=0)
					πF.SetLineno(142)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µwriter_name, "writer_name"); πE != nil {
						continue
					}
					πTemp005[0] = µwriter_name
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
					// line 143: return module.Writer
					πF.SetLineno(143)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmodule, ßWriter, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßget_writer_class.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 135: """Return the Writer class from the `writer_name` module."""
			πF.SetLineno(135)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return the Writer class from the `writer_name` module.").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßget_writer_class); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("writers", Code)
}
