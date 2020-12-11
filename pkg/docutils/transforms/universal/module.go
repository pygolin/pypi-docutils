package universal

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/time"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßDecorations := πg.InternStr("Decorations")
		ßDocutils := πg.InternStr("Docutils")
		ßElement := πg.InternStr("Element")
		ßExposeInternals := πg.InternStr("ExposeInternals")
		ßFalse := πg.InternStr("False")
		ßFilterMessages := πg.InternStr("FilterMessages")
		ßFixedTextElement := πg.InternStr("FixedTextElement")
		ßMessages := πg.InternStr("Messages")
		ßNone := πg.InternStr("None")
		ßSmartQuotes := πg.InternStr("SmartQuotes")
		ßSpecial := πg.InternStr("Special")
		ßStripClassesAndElements := πg.InternStr("StripClassesAndElements")
		ßStripComments := πg.InternStr("StripComments")
		ßTestMessages := πg.InternStr("TestMessages")
		ßText := πg.InternStr("Text")
		ßTextElement := πg.InternStr("TextElement")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_destination := πg.InternStr("_destination")
		ß_source := πg.InternStr("_source")
		ßadd := πg.InternStr("add")
		ßalt := πg.InternStr("alt")
		ßappend := πg.InternStr("append")
		ßapply := πg.InternStr("apply")
		ßcheck_classes := πg.InternStr("check_classes")
		ßclasses := πg.InternStr("classes")
		ßcomment := πg.InternStr("comment")
		ßdatestamp := πg.InternStr("datestamp")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdict := πg.InternStr("dict")
		ßdocument := πg.InternStr("document")
		ßeducate_tokens := πg.InternStr("educate_tokens")
		ßexpose_internals := πg.InternStr("expose_internals")
		ßextend := πg.InternStr("extend")
		ßgenerate_footer := πg.InternStr("generate_footer")
		ßgenerate_header := πg.InternStr("generate_header")
		ßgenerator := πg.InternStr("generator")
		ßget_decoration := πg.InternStr("get_decoration")
		ßget_footer := πg.InternStr("get_footer")
		ßget_header := πg.InternStr("get_header")
		ßget_language_code := πg.InternStr("get_language_code")
		ßget_tokens := πg.InternStr("get_tokens")
		ßgetattr := πg.InternStr("getattr")
		ßgmtime := πg.InternStr("gmtime")
		ßimage := πg.InternStr("image")
		ßisinstance := πg.InternStr("isinstance")
		ßlanguage_code := πg.InternStr("language_code")
		ßlevel := πg.InternStr("level")
		ßliteral := πg.InternStr("literal")
		ßliteral_nodes := πg.InternStr("literal_nodes")
		ßmath := πg.InternStr("math")
		ßnodes := πg.InternStr("nodes")
		ßnodes_to_skip := πg.InternStr("nodes_to_skip")
		ßnormalize_language_tag := πg.InternStr("normalize_language_tag")
		ßnot_Text := πg.InternStr("not_Text")
		ßoption_string := πg.InternStr("option_string")
		ßparagraph := πg.InternStr("paragraph")
		ßparent := πg.InternStr("parent")
		ßplain := πg.InternStr("plain")
		ßproblematic := πg.InternStr("problematic")
		ßqDe := πg.InternStr("qDe")
		ßquotes := πg.InternStr("quotes")
		ßraw := πg.InternStr("raw")
		ßrawsource := πg.InternStr("rawsource")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßrelative_path := πg.InternStr("relative_path")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßreport_level := πg.InternStr("report_level")
		ßreporter := πg.InternStr("reporter")
		ßsection := πg.InternStr("section")
		ßset := πg.InternStr("set")
		ßsettings := πg.InternStr("settings")
		ßsmart_quotes := πg.InternStr("smart_quotes")
		ßsmartchars := πg.InternStr("smartchars")
		ßsmartquotes := πg.InternStr("smartquotes")
		ßsmartquotes_action := πg.InternStr("smartquotes_action")
		ßsmartquotes_locales := πg.InternStr("smartquotes_locales")
		ßsource_link := πg.InternStr("source_link")
		ßsource_url := πg.InternStr("source_url")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßstrftime := πg.InternStr("strftime")
		ßstrip_classes := πg.InternStr("strip_classes")
		ßstrip_comments := πg.InternStr("strip_comments")
		ßstrip_elements := πg.InternStr("strip_elements")
		ßstrip_elements_with_classes := πg.InternStr("strip_elements_with_classes")
		ßsub := πg.InternStr("sub")
		ßsys := πg.InternStr("sys")
		ßsystem_message := πg.InternStr("system_message")
		ßtime := πg.InternStr("time")
		ßtitle := πg.InternStr("title")
		ßtransform_messages := πg.InternStr("transform_messages")
		ßtraverse := πg.InternStr("traverse")
		ßtuple := πg.InternStr("tuple")
		ßunicode := πg.InternStr("unicode")
		ßunsupported_languages := πg.InternStr("unsupported_languages")
		ßupdate := πg.InternStr("update")
		ßutils := πg.InternStr("utils")
		ßversion_info := πg.InternStr("version_info")
		ßwarning := πg.InternStr("warning")
		ßzip := πg.InternStr("zip")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 7: """
			πF.SetLineno(7)
			// line 7: """
			πF.SetLineno(7)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nTransforms needed by most or all documents:\n\n- `Decorations`: Generate a document's header & footer.\n- `Messages`: Placement of system messages stored in\n  `nodes.document.transform_messages`.\n- `TestMessages`: Like `Messages`, used on test runs.\n- `FinalReferences`: Resolve remaining references.\n").ToObject()); πE != nil {
				continue
			}
			// line 17: __docformat__ = 'reStructuredText'
			πF.SetLineno(17)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 19: import re
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
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
			// line 21: import time
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: from docutils import nodes, utils
			πF.SetLineno(22)
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
			// line 23: from docutils.transforms import TransformError, Transform
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransform); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransform.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: from docutils.utils import smartquotes
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.smartquotes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßsmartquotes.ToObject(), πTemp001); πE != nil {
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
			// line 27: if sys.version_info >= (3, 0):
			πF.SetLineno(27)
		Label1:
			// line 28: unicode = str  # noqa
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 31: class Decorations(Transform):
			πF.SetLineno(31)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("Decorations", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 33: """
					πF.SetLineno(33)
					// line 33: """
					πF.SetLineno(33)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Populate a document's decoration element (header, footer).\n    ").ToObject()); πE != nil {
						continue
					}
					// line 37: default_priority = 820
					πF.SetLineno(37)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(820).ToObject()); πE != nil {
						continue
					}
					// line 39: def apply(self):
					πF.SetLineno(39)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µheader_nodes *πg.Object = πg.UnboundLocal
						_ = µheader_nodes
						var µdecoration *πg.Object = πg.UnboundLocal
						_ = µdecoration
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
						var µfooter_nodes *πg.Object = πg.UnboundLocal
						_ = µfooter_nodes
						var µfooter *πg.Object = πg.UnboundLocal
						_ = µfooter
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
							// line 40: header_nodes = self.generate_header()
							πF.SetLineno(40)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgenerate_header, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µheader_nodes = πTemp002
							if πE = πg.CheckLocal(πF, µheader_nodes, "header_nodes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µheader_nodes); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 41: if header_nodes:
							πF.SetLineno(41)
						Label1:
							// line 42: decoration = self.document.get_decoration()
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_decoration, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdecoration = πTemp001
							// line 43: header = decoration.get_header()
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µdecoration, "decoration"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdecoration, ßget_header, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µheader = πTemp002
							// line 44: header.extend(header_nodes)
							πF.SetLineno(44)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader_nodes, "header_nodes"); πE != nil {
								continue
							}
							πTemp004[0] = µheader_nodes
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µheader, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							// line 45: footer_nodes = self.generate_footer()
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgenerate_footer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfooter_nodes = πTemp002
							if πE = πg.CheckLocal(πF, µfooter_nodes, "footer_nodes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µfooter_nodes); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 46: if footer_nodes:
							πF.SetLineno(46)
						Label3:
							// line 47: decoration = self.document.get_decoration()
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_decoration, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdecoration = πTemp001
							// line 48: footer = decoration.get_footer()
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µdecoration, "decoration"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdecoration, ßget_footer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfooter = πTemp002
							// line 49: footer.extend(footer_nodes)
							πF.SetLineno(49)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfooter_nodes, "footer_nodes"); πE != nil {
								continue
							}
							πTemp004[0] = µfooter_nodes
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfooter, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 51: def generate_header(self):
					πF.SetLineno(51)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("generate_header", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 52: return None
							πF.SetLineno(52)
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
					if πE = πClass.SetItem(πF, ßgenerate_header.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 54: def generate_footer(self):
					πF.SetLineno(54)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("generate_footer", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µdatestamp *πg.Object = πg.UnboundLocal
						_ = µdatestamp
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 61: settings = self.document.settings
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							µsettings = πTemp002
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßgenerator, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßdatestamp, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßsource_link, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßsource_url, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 62: if settings.generator or settings.datestamp or settings.source_link \
							πF.SetLineno(62)
						Label2:
							// line 64: text = []
							πF.SetLineno(64)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µtext = πTemp001
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µsettings, ßsource_link, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µsettings, ß_source, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp006
						Label6:
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßsource_url, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label5:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 65: if settings.source_link and settings._source \
							πF.SetLineno(65)
						Label7:
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßsource_url, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 67: if settings.source_url:
							πF.SetLineno(67)
						Label9:
							// line 68: source = settings.source_url
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßsource_url, nil); πE != nil {
								continue
							}
							µsource = πTemp001
							goto Label11
						Label10:
							// line 70: source = utils.relative_path(settings._destination,
							πF.SetLineno(70)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ß_destination, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ß_source, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelative_path, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µsource = πTemp001
							goto Label11
						Label11:
							// line 72: text.extend([
							πF.SetLineno(72)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 2)
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = ß.ToObject()
							πTemp008[1] = πg.NewStr("View document source").ToObject()
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"refuri", µsource},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr(".\n").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp001
							πTemp001 = πg.NewList(πTemp007...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßdatestamp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 76: if settings.datestamp:
							πF.SetLineno(76)
						Label12:
							// line 77: datestamp = time.strftime(settings.datestamp, time.gmtime())
							πF.SetLineno(77)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßdatestamp, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgmtime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstrftime, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µdatestamp = πTemp001
							// line 78: text.append(nodes.Text('Generated on: ' + datestamp + '.\n'))
							πF.SetLineno(78)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdatestamp, "datestamp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("Generated on: ").ToObject(), µdatestamp); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr(".\n").ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßgenerator, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							goto Label15
							// line 79: if settings.generator:
							πF.SetLineno(79)
						Label14:
							// line 80: text.extend([
							πF.SetLineno(80)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 5)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("Generated by ").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = ß.ToObject()
							πTemp008[1] = ßDocutils.ToObject()
							πTemp009 = πg.KWArgs{
								{"refuri", πg.NewStr("http://docutils.sourceforge.net/").ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp001
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr(" from ").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[2] = πTemp001
							πTemp008 = πF.MakeArgs(2)
							πTemp008[0] = ß.ToObject()
							πTemp008[1] = ßreStructuredText.ToObject()
							πTemp009 = πg.KWArgs{
								{"refuri", πg.NewStr("http://docutils.sourceforge.net/rst.html").ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[3] = πTemp001
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr(" source.\n").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßText, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[4] = πTemp001
							πTemp001 = πg.NewList(πTemp007...).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtext, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label15
						Label15:
							// line 88: return [nodes.paragraph('', '', *text)]
							πF.SetLineno(88)
							πTemp004 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(2)
							πTemp007[0] = ß.ToObject()
							πTemp007[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp007, µtext, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp001
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πR = πTemp001
							continue
							goto Label4
						Label3:
							// line 90: return None
							πF.SetLineno(90)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
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
					if πE = πClass.SetItem(πF, ßgenerate_footer.ToObject(), πTemp004); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Decorations").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDecorations.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 93: class ExposeInternals(Transform):
			πF.SetLineno(93)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("ExposeInternals", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 95: """
					πF.SetLineno(95)
					// line 95: """
					πF.SetLineno(95)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Expose internal attributes if ``expose_internals`` setting is set.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 99: default_priority = 840
					πF.SetLineno(99)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(840).ToObject()); πE != nil {
						continue
					}
					// line 101: def not_Text(self, node):
					πF.SetLineno(101)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("not_Text", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 102: return not isinstance(node, nodes.Text)
							πF.SetLineno(102)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnot_Text.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 104: def apply(self):
					πF.SetLineno(104)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µatt *πg.Object = πg.UnboundLocal
						_ = µatt
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßexpose_internals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 105: if self.document.settings.expose_internals:
							πF.SetLineno(105)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnot_Text, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µnode = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßexpose_internals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								µatt = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 108: value = getattr(node, att, None)
							πF.SetLineno(108)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							πTemp004[1] = µatt
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µvalue = πTemp007
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(µvalue != πTemp007).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 109: if value is not None:
							πF.SetLineno(109)
						Label9:
							// line 110: node['internal:' + att] = value
							πF.SetLineno(110)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µatt, "att"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, πg.NewStr("internal:").ToObject(), µatt); πE != nil {
								continue
							}
							πTemp007 = πTemp009
							if πE = πg.SetItem(πF, µnode, πTemp007, πTemp005); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp003); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ExposeInternals").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExposeInternals.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 113: class Messages(Transform):
			πF.SetLineno(113)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("Messages", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 115: """
					πF.SetLineno(115)
					// line 115: """
					πF.SetLineno(115)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Place any system messages generated after parsing into a dedicated section\n    of the document.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 120: default_priority = 860
					πF.SetLineno(120)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(860).ToObject()); πE != nil {
						continue
					}
					// line 122: def apply(self):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µunfiltered *πg.Object = πg.UnboundLocal
						_ = µunfiltered
						var µthreshold *πg.Object = πg.UnboundLocal
						_ = µthreshold
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 123: unfiltered = self.document.transform_messages
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtransform_messages, nil); πE != nil {
								continue
							}
							µunfiltered = πTemp002
							// line 124: threshold = self.document.reporter.report_level
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreport_level, nil); πE != nil {
								continue
							}
							µthreshold = πTemp001
							// line 125: messages = []
							πF.SetLineno(125)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µmessages = πTemp001
							if πE = πg.CheckLocal(πF, µunfiltered, "unfiltered"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µunfiltered); πE != nil {
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
								µmsg = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp007 = ßlevel.ToObject()
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µmsg, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthreshold, "threshold"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GE(πF, πTemp008, µthreshold); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µmsg, ßparent, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(!πTemp009).ToObject()
							πTemp002 = πTemp006
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 127: if msg['level'] >= threshold and not msg.parent:
							πF.SetLineno(127)
						Label5:
							// line 128: messages.append(msg)
							πF.SetLineno(128)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003[0] = µmsg
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmessages); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 129: if messages:
							πF.SetLineno(129)
						Label7:
							// line 130: section = nodes.section(classes=['system-messages'])
							πF.SetLineno(130)
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewStr("system-messages").ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							πTemp010 = πg.KWArgs{
								{"classes", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsection, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							µsection = πTemp001
							// line 132: section += nodes.title('', 'Docutils System Messages')
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ß.ToObject()
							πTemp003[1] = πg.NewStr("Docutils System Messages").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.IAdd(πF, µsection, πTemp001); πE != nil {
								continue
							}
							µsection = πTemp002
							// line 133: section += messages
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µsection, µmessages); πE != nil {
								continue
							}
							µsection = πTemp001
							// line 134: self.document.transform_messages[:] = []
							πF.SetLineno(134)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßtransform_messages, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp007, πTemp006, πTemp002); πE != nil {
								continue
							}
							// line 135: self.document += section
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, µsection); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Messages").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMessages.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 138: class FilterMessages(Transform):
			πF.SetLineno(138)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("FilterMessages", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 140: """
					πF.SetLineno(140)
					// line 140: """
					πF.SetLineno(140)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Remove system messages below verbosity threshold.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 144: default_priority = 870
					πF.SetLineno(144)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(870).ToObject()); πE != nil {
						continue
					}
					// line 146: def apply(self):
					πF.SetLineno(146)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πg.UnboundLocal
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
						var πTemp007 bool
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
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsystem_message, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
								µnode = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp005 = ßlevel.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp005, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp009, ßreport_level, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 148: if node['level'] < self.document.reporter.report_level:
							πF.SetLineno(148)
						Label4:
							// line 149: node.parent.remove(node)
							πF.SetLineno(149)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßremove, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("FilterMessages").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFilterMessages.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 152: class TestMessages(Transform):
			πF.SetLineno(152)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("TestMessages", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 154: """
					πF.SetLineno(154)
					// line 154: """
					πF.SetLineno(154)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Append all post-parse system messages to the end of the document.\n\n    Used for testing purposes.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 160: default_priority = 880
					πF.SetLineno(160)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(880).ToObject()); πE != nil {
						continue
					}
					// line 162: def apply(self):
					πF.SetLineno(162)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp005 bool
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtransform_messages, nil); πE != nil {
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
								µmsg = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 164: if not msg.parent:
							πF.SetLineno(164)
						Label4:
							// line 165: self.document += msg
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, µmsg); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp003); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestMessages").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestMessages.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 168: class StripComments(Transform):
			πF.SetLineno(168)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("StripComments", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 170: """
					πF.SetLineno(170)
					// line 170: """
					πF.SetLineno(170)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Remove comment elements from the document tree (only if the\n    ``strip_comments`` setting is enabled).\n    ").ToObject()); πE != nil {
						continue
					}
					// line 175: default_priority = 740
					πF.SetLineno(175)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(740).ToObject()); πE != nil {
						continue
					}
					// line 177: def apply(self):
					πF.SetLineno(177)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
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
						var πTemp007 bool
						_ = πTemp007
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstrip_comments, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 178: if self.document.settings.strip_comments:
							πF.SetLineno(178)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßcomment, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µnode = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 180: node.parent.remove(node)
							πF.SetLineno(180)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StripComments").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStripComments.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 183: class StripClassesAndElements(Transform):
			πF.SetLineno(183)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("StripClassesAndElements", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 185: """
					πF.SetLineno(185)
					// line 185: """
					πF.SetLineno(185)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Remove from the document tree all elements with classes in\n    `self.document.settings.strip_elements_with_classes` and all \"classes\"\n    attribute values in `self.document.settings.strip_classes`.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 191: default_priority = 420
					πF.SetLineno(191)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(420).ToObject()); πE != nil {
						continue
					}
					// line 193: def apply(self):
					πF.SetLineno(193)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µstrip_classes *πg.Object = πg.UnboundLocal
						_ = µstrip_classes
						var µclass_value *πg.Object = πg.UnboundLocal
						_ = µclass_value
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
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
							case 8:
								goto Label8
							case 9:
								goto Label9
							case 11:
								goto Label11
							case 12:
								goto Label12
							case 15:
								goto Label15
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstrip_elements_with_classes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 194: if self.document.settings.strip_elements_with_classes:
							πF.SetLineno(194)
						Label1:
							// line 195: self.strip_elements = set(
							πF.SetLineno(195)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstrip_elements_with_classes, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstrip_elements, πTemp001); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheck_classes, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µnode = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 200: node.parent.remove(node)
							πF.SetLineno(200)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßstrip_classes, nil); πE != nil {
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
								goto Label6
							}
							goto Label7
							// line 202: if not self.document.settings.strip_classes:
							πF.SetLineno(202)
						Label6:
							// line 203: return
							πF.SetLineno(203)
							πR = πg.None
							continue
							goto Label7
						Label7:
							// line 204: strip_classes = self.document.settings.strip_classes
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstrip_classes, nil); πE != nil {
								continue
							}
							µstrip_classes = πTemp001
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßElement, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp003 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label10
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
								µnode = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(8)
							if πE = πg.CheckLocal(πF, µstrip_classes, "strip_classes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µstrip_classes); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp007 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µclass_value = πTemp006
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(11)
							// line 207: try:
							πF.SetLineno(207)
							πF.PushCheckpoint(15)
							// line 208: node['classes'].remove(class_value)
							πF.SetLineno(208)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclass_value, "class_value"); πE != nil {
								continue
							}
							πTemp004[0] = µclass_value
							πTemp006 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp009, ßremove, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.PopCheckpoint()
							goto Label14
						Label15:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label16
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 209: except ValueError:
							πF.SetLineno(209)
						Label16:
							// line 210: pass
							πF.SetLineno(210)
							πF.RestoreExc(nil, nil)
							goto Label14
						Label14:
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 212: def check_classes(self, node):
					πF.SetLineno(212)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("check_classes", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µclass_value *πg.Object = πg.UnboundLocal
						_ = µclass_value
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßElement, nil); πE != nil {
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
							// line 213: if not isinstance(node, nodes.Element):
							πF.SetLineno(213)
						Label1:
							// line 214: return False
							πF.SetLineno(214)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp006 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
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
								µclass_value = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µclass_value, "class_value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstrip_elements, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, πTemp004, µclass_value); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 216: if class_value in self.strip_elements:
							πF.SetLineno(216)
						Label6:
							// line 217: return True
							πF.SetLineno(217)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 218: return False
							πF.SetLineno(218)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheck_classes.ToObject(), πTemp003); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StripClassesAndElements").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStripClassesAndElements.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 221: class SmartQuotes(Transform):
			πF.SetLineno(221)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("SmartQuotes", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
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
				var πTemp006 []πg.Param
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 223: """
					πF.SetLineno(223)
					// line 223: """
					πF.SetLineno(223)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Replace ASCII quotation marks with typographic form.\n\n    Also replace multiple dashes with em-dash/en-dash characters.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 229: default_priority = 850
					πF.SetLineno(229)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(850).ToObject()); πE != nil {
						continue
					}
					// line 231: nodes_to_skip = (nodes.FixedTextElement, nodes.Special)
					πF.SetLineno(231)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßFixedTextElement, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßSpecial, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
					if πE = πClass.SetItem(πF, ßnodes_to_skip.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 232: """Do not apply "smartquotes" to instances of these block-level nodes."""
					πF.SetLineno(232)
					// line 234: literal_nodes = (nodes.FixedTextElement, nodes.Special,
					πF.SetLineno(234)
					πTemp005 = make([]*πg.Object, 7)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßFixedTextElement, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSpecial, nil); πE != nil {
						continue
					}
					πTemp005[1] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßimage, nil); πE != nil {
						continue
					}
					πTemp005[2] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral, nil); πE != nil {
						continue
					}
					πTemp005[3] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmath, nil); πE != nil {
						continue
					}
					πTemp005[4] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßraw, nil); πE != nil {
						continue
					}
					πTemp005[5] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßproblematic, nil); πE != nil {
						continue
					}
					πTemp005[6] = πTemp003
					πTemp001 = πg.NewTuple(πTemp005...).ToObject()
					if πE = πClass.SetItem(πF, ßliteral_nodes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 237: """Do apply smartquotes to instances of these inline nodes."""
					πF.SetLineno(237)
					// line 239: smartquotes_action = 'qDe'
					πF.SetLineno(239)
					if πE = πClass.SetItem(πF, ßsmartquotes_action.ToObject(), ßqDe.ToObject()); πE != nil {
						continue
					}
					// line 240: """Setting to select smartquote transformations.
					πF.SetLineno(240)
					// line 246: def __init__(self, document, startnode):
					πF.SetLineno(246)
					πTemp006 = make([]πg.Param, 3)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "document", Def: nil}
					πTemp006[2] = πg.Param{Name: "startnode", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µstartnode *πg.Object = πArgs[2]
						_ = µstartnode
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
							// line 247: Transform.__init__(self, document, startnode=startnode)
							πF.SetLineno(247)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"startnode", µstartnode},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 248: self.unsupported_languages = set()
							πF.SetLineno(248)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßunsupported_languages, πTemp003); πE != nil {
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
					// line 250: def get_tokens(self, txtnodes):
					πF.SetLineno(250)
					πTemp006 = make([]πg.Param, 2)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp006[1] = πg.Param{Name: "txtnodes", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("get_tokens", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtxtnodes *πg.Object = πArgs[1]
						_ = µtxtnodes
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µtxt *πg.Object = πg.UnboundLocal
						_ = µtxt
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
								case 8:
									goto Label8
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 9:
									goto Label9
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µtxtnodes, "txtnodes"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µtxtnodes); πE != nil {
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
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp006
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßliteral_nodes, nil); πE != nil {
									continue
								}
								πTemp005[1] = πTemp006
								if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp004 = πTemp007
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label4
								}
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßparent, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp007
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßliteral_nodes, nil); πE != nil {
									continue
								}
								πTemp005[1] = πTemp006
								if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp004 = πTemp007
							Label4:
								if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label5
								}
								goto Label6
								// line 254: if (isinstance(node.parent, self.literal_nodes)
								πF.SetLineno(254)
							Label5:
								// line 256: yield ('literal', unicode(node))
								πF.SetLineno(256)
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								πTemp005[0] = µnode
								if πTemp006, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πTemp004 = πg.NewTuple2(ßliteral.ToObject(), πTemp007).ToObject()
								πF.PushCheckpoint(8)
								return πTemp004, nil
							Label8:
								πTemp006 = πSent
								goto Label7
							Label6:
								// line 259: txt = re.sub('(?<=\x00)([-\\\'".`])', r'\\\1', unicode(node))
								πF.SetLineno(259)
								πTemp005 = πF.MakeArgs(3)
								πTemp005[0] = πg.NewStr("(?<=\x00)([-\\'\".`])").ToObject()
								πTemp005[1] = πg.NewStr("\\\\\\1").ToObject()
								πTemp008 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
									continue
								}
								πTemp008[0] = µnode
								if πTemp006, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp008)
								πTemp005[2] = πTemp007
								if πTemp006, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßsub, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µtxt = πTemp006
								// line 260: yield ('plain', txt)
								πF.SetLineno(260)
								if πE = πg.CheckLocal(πF, µtxt, "txt"); πE != nil {
									continue
								}
								πTemp006 = πg.NewTuple2(ßplain.ToObject(), µtxt).ToObject()
								πF.PushCheckpoint(9)
								return πTemp006, nil
							Label9:
								πTemp007 = πSent
								goto Label7
							Label7:
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
					if πE = πClass.SetItem(πF, ßget_tokens.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 262: def apply(self):
					πF.SetLineno(262)
					πTemp006 = make([]πg.Param, 1)
					πTemp006[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsmart_quotes *πg.Object = πg.UnboundLocal
						_ = µsmart_quotes
						var µalternative *πg.Object = πg.UnboundLocal
						_ = µalternative
						var µdocument_language *πg.Object = πg.UnboundLocal
						_ = µdocument_language
						var µlc_smartquotes *πg.Object = πg.UnboundLocal
						_ = µlc_smartquotes
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µtxtnodes *πg.Object = πg.UnboundLocal
						_ = µtxtnodes
						var µlang *πg.Object = πg.UnboundLocal
						_ = µlang
						var µtag *πg.Object = πg.UnboundLocal
						_ = µtag
						var µteacher *πg.Object = πg.UnboundLocal
						_ = µteacher
						var µtxtnode *πg.Object = πg.UnboundLocal
						_ = µtxtnode
						var µnewtext *πg.Object = πg.UnboundLocal
						_ = µnewtext
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []πg.Param
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 πg.KWArgs
						_ = πTemp015
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4:
								goto Label4
							case 8:
								goto Label8
							case 9:
								goto Label9
							case 20:
								goto Label20
							case 21:
								goto Label21
							case 27:
								goto Label27
							case 28:
								goto Label28
							default:
								panic("unexpected function state")
							}
							// line 263: smart_quotes = self.document.settings.smart_quotes
							πF.SetLineno(263)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsmart_quotes, nil); πE != nil {
								continue
							}
							µsmart_quotes = πTemp001
							if πE = πg.CheckLocal(πF, µsmart_quotes, "smart_quotes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µsmart_quotes); πE != nil {
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
							// line 264: if not smart_quotes:
							πF.SetLineno(264)
						Label1:
							// line 265: return
							πF.SetLineno(265)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 266: try:
							πF.SetLineno(266)
							πF.PushCheckpoint(4)
							// line 267: alternative = smart_quotes.startswith('alt')
							πF.SetLineno(267)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßalt.ToObject()
							if πE = πg.CheckLocal(πF, µsmart_quotes, "smart_quotes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsmart_quotes, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µalternative = πTemp002
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 268: except AttributeError:
							πF.SetLineno(268)
						Label5:
							// line 269: alternative = False
							πF.SetLineno(269)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µalternative = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 271: document_language = self.document.settings.language_code
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßlanguage_code, nil); πE != nil {
								continue
							}
							µdocument_language = πTemp001
							// line 272: lc_smartquotes = self.document.settings.smartquotes_locales
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsmartquotes_locales, nil); πE != nil {
								continue
							}
							µlc_smartquotes = πTemp001
							if πE = πg.CheckLocal(πF, µlc_smartquotes, "lc_smartquotes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µlc_smartquotes); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 273: if lc_smartquotes:
							πF.SetLineno(273)
						Label6:
							// line 274: smartquotes.smartchars.quotes.update(dict(lc_smartquotes))
							πF.SetLineno(274)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlc_smartquotes, "lc_smartquotes"); πE != nil {
								continue
							}
							πTemp007[0] = µlc_smartquotes
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsmartquotes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsmartchars, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßquotes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label7
						Label7:
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp003 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label10
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µnode = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnodes_to_skip, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label11
							}
							goto Label12
							// line 280: if isinstance(node, self.nodes_to_skip):
							πF.SetLineno(280)
						Label11:
							// line 281: continue
							πF.SetLineno(281)
							continue
							goto Label12
						Label12:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp008
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label13
							}
							goto Label14
							// line 283: if isinstance(node.parent, nodes.TextElement):
							πF.SetLineno(283)
						Label13:
							// line 284: continue
							πF.SetLineno(284)
							continue
							goto Label14
						Label14:
							// line 287: txtnodes = [txtnode for txtnode in node.traverse(nodes.Text)
							πF.SetLineno(287)
							πTemp010 = make([]πg.Param, 0)
							πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/universal.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µtxtnode *πg.Object = πg.UnboundLocal
								_ = µtxtnode
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
										if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
											continue
										}
										πTemp002[0] = πTemp004
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µnode, ßtraverse, nil); πE != nil {
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
											µtxtnode = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)
										πTemp002 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µtxtnode, "txtnode"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µtxtnode, ßparent, nil); πE != nil {
											continue
										}
										πTemp002[0] = πTemp004
										if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßoption_string, nil); πE != nil {
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
										if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
											continue
										}
										πTemp003 = πg.GetBool(!πTemp006).ToObject()
										if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
											continue
										}
										if πTemp006 {
											goto Label4
										}
										goto Label5
										// line 287: txtnodes = [txtnode for txtnode in node.traverse(nodes.Text)
										πF.SetLineno(287)
									Label4:
										// line 287: txtnodes = [txtnode for txtnode in node.traverse(nodes.Text)
										πF.SetLineno(287)
										if πE = πg.CheckLocal(πF, µtxtnode, "txtnode"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µtxtnode, nil
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
							if πTemp011, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
								continue
							}
							µtxtnodes = πTemp002
							// line 292: lang = node.get_language_code(document_language)
							πF.SetLineno(292)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdocument_language, "document_language"); πE != nil {
								continue
							}
							πTemp004[0] = µdocument_language
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßget_language_code, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µlang = πTemp011
							if πE = πg.CheckLocal(πF, µalternative, "alternative"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µalternative); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label15
							}
							goto Label16
							// line 294: if alternative:
							πF.SetLineno(294)
						Label15:
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µlang, πg.NewStr("-x-altquot").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label17
							}
							goto Label18
							// line 295: if '-x-altquot' in lang:
							πF.SetLineno(295)
						Label17:
							// line 296: lang = lang.replace('-x-altquot', '')
							πF.SetLineno(296)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("-x-altquot").ToObject()
							πTemp004[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlang, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µlang = πTemp011
							goto Label19
						Label18:
							// line 298: lang += '-x-altquot'
							πF.SetLineno(298)
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µlang, πg.NewStr("-x-altquot").ToObject()); πE != nil {
								continue
							}
							µlang = πTemp002
							goto Label19
						Label19:
							goto Label16
						Label16:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							πTemp004[0] = µlang
							if πTemp011, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßnormalize_language_tag, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Iter(πF, πTemp011); πE != nil {
								continue
							}
							πF.PushCheckpoint(21)
							πTemp009 = false
						Label20:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label22
							}
							if πTemp011, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µtag = πTemp011
							}
							if πE != nil || !πTemp013 {
								continue
							}
							πF.PushCheckpoint(20)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßsmartquotes); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, πTemp012, ßsmartchars, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp014, ßquotes, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Contains(πF, πTemp012, µtag); πE != nil {
								continue
							}
							πTemp011 = πg.GetBool(πTemp013).ToObject()
							if πTemp013, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if πTemp013 {
								goto Label23
							}
							goto Label24
							// line 301: if tag in smartquotes.smartchars.quotes:
							πF.SetLineno(301)
						Label23:
							// line 302: lang = tag
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							µlang = µtag
							// line 303: break
							πF.SetLineno(303)
							πTemp009 = true
							continue
							goto Label24
						Label24:
							continue
						Label21:
							if πE != nil || πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßunsupported_languages, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Contains(πF, πTemp012, µlang); πE != nil {
								continue
							}
							πTemp011 = πg.GetBool(!πTemp013).ToObject()
							if πTemp013, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if πTemp013 {
								goto Label25
							}
							goto Label26
							// line 305: if lang not in self.unsupported_languages:
							πF.SetLineno(305)
						Label25:
							// line 306: self.document.reporter.warning('No smart quotes '
							πF.SetLineno(306)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Mod(πF, πg.NewStr("No smart quotes defined for language \"%s\".").ToObject(), µlang); πE != nil {
								continue
							}
							πTemp004[0] = πTemp011
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp012, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp004, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label26
						Label26:
							// line 308: self.unsupported_languages.add(lang)
							πF.SetLineno(308)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							πTemp004[0] = µlang
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßunsupported_languages, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßadd, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 309: lang = ''
							πF.SetLineno(309)
							µlang = ß.ToObject()
						Label22:
							// line 313: teacher = smartquotes.educate_tokens(self.get_tokens(txtnodes),
							πF.SetLineno(313)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtxtnodes, "txtnodes"); πE != nil {
								continue
							}
							πTemp007[0] = µtxtnodes
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_tokens, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp011
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsmartquotes_action, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"attr", πTemp002},
								{"language", µlang},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartquotes); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp002, ßeducate_tokens, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp011.Call(πF, πTemp004, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µteacher = πTemp002
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtxtnodes, "txtnodes"); πE != nil {
								continue
							}
							πTemp004[0] = µtxtnodes
							if πE = πg.CheckLocal(πF, µteacher, "teacher"); πE != nil {
								continue
							}
							πTemp004[1] = µteacher
							if πTemp011, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Iter(πF, πTemp012); πE != nil {
								continue
							}
							πF.PushCheckpoint(28)
							πTemp009 = false
						Label27:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label29
							}
							if πTemp011, πE = πg.Next(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp014}}}, πTemp011); πE != nil {
									continue
								}
								µtxtnode = πTemp012
								µnewtext = πTemp014
							}
							if πE != nil || !πTemp013 {
								continue
							}
							πF.PushCheckpoint(27)
							// line 317: txtnode.parent.replace(txtnode, nodes.Text(newtext,
							πF.SetLineno(317)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtxtnode, "txtnode"); πE != nil {
								continue
							}
							πTemp004[0] = µtxtnode
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnewtext, "newtext"); πE != nil {
								continue
							}
							πTemp007[0] = µnewtext
							if πE = πg.CheckLocal(πF, µtxtnode, "txtnode"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µtxtnode, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"rawsource", πTemp011},
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßText, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, πTemp007, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[1] = πTemp011
							if πE = πg.CheckLocal(πF, µtxtnode, "txtnode"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µtxtnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label28:
							if πE != nil || πR != nil {
								continue
							}
						Label29:
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 320: self.unsupported_languages = set() # reset
							πF.SetLineno(320)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßunsupported_languages, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp003); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SmartQuotes").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSmartQuotes.ToObject(), πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("universal", Code)
}
