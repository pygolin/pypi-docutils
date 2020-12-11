package docutils_xml

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/StringIO"
	_ "github.com/pygolin/stdlib/pkg/io"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/xml"
	_ "github.com/pygolin/stdlib/pkg/xml/sax"
	_ "github.com/pygolin/stdlib/pkg/xml/sax/saxutils"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßApplicationError := πg.InternStr("ApplicationError")
		ßContentHandler := πg.InternStr("ContentHandler")
		ßFixedTextElement := πg.InternStr("FixedTextElement")
		ßGenericNodeVisitor := πg.InternStr("GenericNodeVisitor")
		ßNodeVisitor := πg.InternStr("NodeVisitor")
		ßNone := πg.InternStr("None")
		ßRawXmlError := πg.InternStr("RawXmlError")
		ßSAXParseException := πg.InternStr("SAXParseException")
		ßSkipNode := πg.InternStr("SkipNode")
		ßStringIO := πg.InternStr("StringIO")
		ßTestXml := πg.InternStr("TestXml")
		ßTextElement := πg.InternStr("TextElement")
		ßTrue := πg.InternStr("True")
		ßWriter := πg.InternStr("Writer")
		ßXMLTranslator := πg.InternStr("XMLTranslator")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß_exceptions := πg.InternStr("_exceptions")
		ßaction := πg.InternStr("action")
		ßappend := πg.InternStr("append")
		ßastext := πg.InternStr("astext")
		ßcolspec := πg.InternStr("colspec")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßdefault := πg.InternStr("default")
		ßdefault_departure := πg.InternStr("default_departure")
		ßdefault_visit := πg.InternStr("default_visit")
		ßdepart_Text := πg.InternStr("depart_Text")
		ßdest := πg.InternStr("dest")
		ßdoctype := πg.InternStr("doctype")
		ßdoctype_declaration := πg.InternStr("doctype_declaration")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßencode := πg.InternStr("encode")
		ßendtag := πg.InternStr("endtag")
		ßerror := πg.InternStr("error")
		ßescape := πg.InternStr("escape")
		ßfixed_text := πg.InternStr("fixed_text")
		ßformat := πg.InternStr("format")
		ßfrontend := πg.InternStr("frontend")
		ßgenerator := πg.InternStr("generator")
		ßget := πg.InternStr("get")
		ßgetColumnNumber := πg.InternStr("getColumnNumber")
		ßgetLineNumber := πg.InternStr("getLineNumber")
		ßhandler := πg.InternStr("handler")
		ßimage := πg.InternStr("image")
		ßin_simple := πg.InternStr("in_simple")
		ßindent := πg.InternStr("indent")
		ßindents := πg.InternStr("indents")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlevel := πg.InternStr("level")
		ßline := πg.InternStr("line")
		ßliteral := πg.InternStr("literal")
		ßlocator := πg.InternStr("locator")
		ßmake_parser := πg.InternStr("make_parser")
		ßnewline := πg.InternStr("newline")
		ßnewlines := πg.InternStr("newlines")
		ßnodes := πg.InternStr("nodes")
		ßoutput := πg.InternStr("output")
		ßoutput_encoding := πg.InternStr("output_encoding")
		ßoutput_encoding_error_handler := πg.InternStr("output_encoding_error_handler")
		ßparent := πg.InternStr("parent")
		ßparse := πg.InternStr("parse")
		ßquoteattr := πg.InternStr("quoteattr")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreplace := πg.InternStr("replace")
		ßreporter := πg.InternStr("reporter")
		ßsax := πg.InternStr("sax")
		ßsaxutils := πg.InternStr("saxutils")
		ßsetContentHandler := πg.InternStr("setContentHandler")
		ßsetDocumentLocator := πg.InternStr("setDocumentLocator")
		ßsetFeature := πg.InternStr("setFeature")
		ßsettings := πg.InternStr("settings")
		ßsettings_defaults := πg.InternStr("settings_defaults")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsimple_nodes := πg.InternStr("simple_nodes")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßstarttag := πg.InternStr("starttag")
		ßstore_false := πg.InternStr("store_false")
		ßstore_true := πg.InternStr("store_true")
		ßstr := πg.InternStr("str")
		ßsupported := πg.InternStr("supported")
		ßsys := πg.InternStr("sys")
		ßthe_handle := πg.InternStr("the_handle")
		ßtransition := πg.InternStr("transition")
		ßtranslate := πg.InternStr("translate")
		ßtranslator_class := πg.InternStr("translator_class")
		ßunicode := πg.InternStr("unicode")
		ßutf8 := πg.InternStr("utf8")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidator := πg.InternStr("validator")
		ßversion_info := πg.InternStr("version_info")
		ßvisit_Text := πg.InternStr("visit_Text")
		ßvisit_raw := πg.InternStr("visit_raw")
		ßvisitor := πg.InternStr("visitor")
		ßwalkabout := πg.InternStr("walkabout")
		ßwarn := πg.InternStr("warn")
		ßwarning := πg.InternStr("warning")
		ßwriters := πg.InternStr("writers")
		ßxml := πg.InternStr("xml")
		ßxml_declaration := πg.InternStr("xml_declaration")
		ßxmlcharrefreplace := πg.InternStr("xmlcharrefreplace")
		ßxmlparser := πg.InternStr("xmlparser")
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
		var πTemp007 *πg.Object
		_ = πTemp007
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nSimple document tree Writer, writes Docutils XML according to\nhttp://docutils.sourceforge.net/docs/ref/docutils.dtd.\n").ToObject()); πE != nil {
				continue
			}
			// line 11: __docformat__ = 'reStructuredText'
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 13: import sys
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import xml.sax.saxutils
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "xml.sax.saxutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßxml.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import docutils
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: from docutils import frontend, writers, nodes
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßfrontend.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßwriters.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
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
			// line 19: if sys.version_info >= (3, 0):
			πF.SetLineno(19)
		Label1:
			// line 20: from io import StringIO  # noqa
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label3
		Label2:
			// line 22: from StringIO import StringIO  # noqa
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label3
		Label3:
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
				goto Label4
			}
			goto Label5
			// line 25: if sys.version_info >= (3, 0):
			πF.SetLineno(25)
		Label4:
			// line 26: unicode = str  # noqa
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label5
		Label5:
			// line 29: class RawXmlError(docutils.ApplicationError): pass
			πF.SetLineno(29)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßApplicationError, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("RawXmlError", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 29: class RawXmlError(docutils.ApplicationError): pass
					πF.SetLineno(29)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("RawXmlError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRawXmlError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 32: class Writer(writers.Writer):
			πF.SetLineno(32)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Dict
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
				var πTemp013 []πg.Param
				_ = πTemp013
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 34: supported = ('xml',)
					πF.SetLineno(34)
					πTemp001 = πg.NewTuple1(ßxml.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 35: """Formats this writer supports."""
					πF.SetLineno(35)
					// line 35: """Formats this writer supports."""
					πF.SetLineno(35)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Formats this writer supports.").ToObject()); πE != nil {
						continue
					}
					// line 37: settings_spec = (
					πF.SetLineno(37)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--newlines").ToObject()
					πTemp006 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp009); πE != nil {
						continue
					}
					πTemp008 = πTemp007.ToObject()
					πTemp004 = πg.NewTuple3(πg.NewStr("Generate XML with newlines before and after tags.").ToObject(), πTemp006, πTemp008).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--indents").ToObject()
					πTemp008 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp007.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Generate XML with indents and newlines.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--no-xml-declaration").ToObject()
					πTemp009 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßdest.ToObject(), ßxml_declaration.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp010 = πTemp007.ToObject()
					πTemp008 = πg.NewTuple3(πg.NewStr("Omit the XML declaration.  Use with caution.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--no-doctype").ToObject()
					πTemp010 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßdest.ToObject(), ßdoctype_declaration.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp012); πE != nil {
						continue
					}
					πTemp011 = πTemp007.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Omit the DOCTYPE declaration.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp003 = πg.NewTuple4(πTemp004, πTemp006, πTemp008, πTemp009).ToObject()
					πTemp001 = πg.NewTuple3(πg.NewStr("\"Docutils XML\" Writer Options").ToObject(), πTemp002, πTemp003).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 55: settings_defaults = {'output_encoding_error_handler': 'xmlcharrefreplace'}
					πF.SetLineno(55)
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßoutput_encoding_error_handler.ToObject(), ßxmlcharrefreplace.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp007.ToObject()
					if πE = πClass.SetItem(πF, ßsettings_defaults.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 57: config_section = 'docutils_xml writer'
					πF.SetLineno(57)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("docutils_xml writer").ToObject()); πE != nil {
						continue
					}
					// line 58: config_section_dependencies = ('writers',)
					πF.SetLineno(58)
					πTemp001 = πg.NewTuple1(ßwriters.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 60: output = None
					πF.SetLineno(60)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßoutput.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 61: """Final translated form of `document`."""
					πF.SetLineno(61)
					// line 63: def __init__(self):
					πF.SetLineno(63)
					πTemp013 = make([]πg.Param, 1)
					πTemp013[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 64: writers.Writer.__init__(self)
							πF.SetLineno(64)
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 65: self.translator_class = XMLTranslator
							πF.SetLineno(65)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßXMLTranslator); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtranslator_class, πTemp003); πE != nil {
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
					// line 67: def translate(self):
					πF.SetLineno(67)
					πTemp013 = make([]πg.Param, 1)
					πTemp013[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("translate", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
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
							// line 68: self.visitor = visitor = self.translator_class(self.document)
							πF.SetLineno(68)
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
							// line 69: self.document.walkabout(visitor)
							πF.SetLineno(69)
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
							// line 70: self.output = ''.join(visitor.output)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µvisitor, ßoutput, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtranslate.ToObject(), πTemp002); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 73: class XMLTranslator(nodes.GenericNodeVisitor):
			πF.SetLineno(73)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßGenericNodeVisitor, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("XMLTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
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
					// line 75: xml_declaration = '<?xml version="1.0" encoding="%s"?>\n'
					πF.SetLineno(75)
					// line 75: xml_declaration = '<?xml version="1.0" encoding="%s"?>\n'
					πF.SetLineno(75)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("<?xml version=\"1.0\" encoding=\"%s\"?>\n").ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßxml_declaration.ToObject(), πg.NewStr("<?xml version=\"1.0\" encoding=\"%s\"?>\n").ToObject()); πE != nil {
						continue
					}
					// line 78: doctype = (
					πF.SetLineno(78)
					if πE = πClass.SetItem(πF, ßdoctype.ToObject(), πg.NewStr("<!DOCTYPE document PUBLIC \"+//IDN docutils.sourceforge.net//DTD Docutils Generic//EN//XML\" \"http://docutils.sourceforge.net/docs/ref/docutils.dtd\">\n").ToObject()); πE != nil {
						continue
					}
					// line 82: generator = '<!-- Generated by Docutils %s -->\n'
					πF.SetLineno(82)
					if πE = πClass.SetItem(πF, ßgenerator.ToObject(), πg.NewStr("<!-- Generated by Docutils %s -->\n").ToObject()); πE != nil {
						continue
					}
					// line 84: xmlparser = xml.sax.make_parser()
					πF.SetLineno(84)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßxml); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsax, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßmake_parser, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßxmlparser.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 85: """SAX parser instance to check/exctract raw XML."""
					πF.SetLineno(85)
					// line 86: xmlparser.setFeature(
					πF.SetLineno(86)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr("http://xml.org/sax/features/external-general-entities").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßxmlparser); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsetFeature, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 89: def __init__(self, document):
					πF.SetLineno(89)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
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
							// line 90: nodes.NodeVisitor.__init__(self, document)
							πF.SetLineno(90)
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
							// line 93: self.warn = self.document.reporter.warning
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßwarning, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwarn, πTemp003); πE != nil {
								continue
							}
							// line 94: self.error = self.document.reporter.error
							πF.SetLineno(94)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßerror, πTemp003); πE != nil {
								continue
							}
							// line 97: self.settings = settings = document.settings
							πF.SetLineno(97)
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
							// line 98: self.indent = self.newline = ''
							πF.SetLineno(98)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindent, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnewline, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßnewlines, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 99: if settings.newlines:
							πF.SetLineno(99)
						Label1:
							// line 100: self.newline = '\n'
							πF.SetLineno(100)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnewline, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßindents, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 101: if settings.indents:
							πF.SetLineno(101)
						Label3:
							// line 102: self.newline = '\n'
							πF.SetLineno(102)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnewline, πTemp002); πE != nil {
								continue
							}
							// line 103: self.indent = '    ' #@ TODO make this configurable?
							πF.SetLineno(103)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("    ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindent, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 104: self.level = 0  # indentation level
							πF.SetLineno(104)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
								continue
							}
							// line 105: self.in_simple = 0 # level of nesting inside mixed-content elements
							πF.SetLineno(105)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_simple, πTemp002); πE != nil {
								continue
							}
							// line 106: self.fixed_text = 0 # level of nesting inside FixedText elements
							πF.SetLineno(106)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfixed_text, πTemp002); πE != nil {
								continue
							}
							// line 109: self.output = []
							πF.SetLineno(109)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoutput, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßxml_declaration, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 110: if settings.xml_declaration:
							πF.SetLineno(110)
						Label5:
							// line 111: self.output.append(
							πF.SetLineno(111)
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
							if πTemp005, πE = πg.GetAttr(πF, µsettings, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßdoctype_declaration, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 113: if settings.doctype_declaration:
							πF.SetLineno(113)
						Label7:
							// line 114: self.output.append(self.doctype)
							πF.SetLineno(114)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdoctype, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label8
						Label8:
							// line 115: self.output.append(self.generator % docutils.__version__)
							πF.SetLineno(115)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßgenerator, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__version__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 118: self.the_handle=TestXml()
							πF.SetLineno(118)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTestXml); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßthe_handle, πTemp002); πE != nil {
								continue
							}
							// line 119: self.xmlparser.setContentHandler(self.the_handle)
							πF.SetLineno(119)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßthe_handle, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßxmlparser, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsetContentHandler, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 124: simple_nodes = (nodes.TextElement,
					πF.SetLineno(124)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTextElement, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßimage, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßcolspec, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp005, ßtransition, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πTemp006, πTemp007, πTemp008, πTemp009).ToObject()
					if πE = πClass.SetItem(πF, ßsimple_nodes.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 127: def default_visit(self, node):
					πF.SetLineno(127)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("default_visit", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 *πg.Object
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
							// line 128: """Default node visit method."""
							πF.SetLineno(128)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_simple, nil); πE != nil {
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
							// line 129: if not self.in_simple:
							πF.SetLineno(129)
						Label1:
							// line 130: self.output.append(self.indent*self.level)
							πF.SetLineno(130)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							// line 131: self.output.append(node.starttag(xml.sax.saxutils.quoteattr))
							πF.SetLineno(131)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßxml); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsax, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsaxutils, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßquoteattr, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 132: self.level += 1
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßFixedTextElement, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßliteral, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp005, πTemp007).ToObject()
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
								goto Label3
							}
							goto Label4
							// line 134: if isinstance(node, (nodes.FixedTextElement, nodes.literal)):
							πF.SetLineno(134)
						Label3:
							// line 135: self.fixed_text += 1
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfixed_text, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfixed_text, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsimple_nodes, nil); πE != nil {
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
								goto Label5
							}
							goto Label6
							// line 136: if isinstance(node, self.simple_nodes):
							πF.SetLineno(136)
						Label5:
							// line 137: self.in_simple += 1
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßin_simple, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_simple, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_simple, nil); πE != nil {
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
								goto Label7
							}
							goto Label8
							// line 138: if not self.in_simple:
							πF.SetLineno(138)
						Label7:
							// line 139: self.output.append(self.newline)
							πF.SetLineno(139)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnewline, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßdefault_visit.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 128: """Default node visit method."""
					πF.SetLineno(128)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Default node visit method.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_visit); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 141: def default_departure(self, node):
					πF.SetLineno(141)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("default_departure", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 142: """Default node depart method."""
							πF.SetLineno(142)
							// line 143: self.level -= 1
							πF.SetLineno(143)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_simple, nil); πE != nil {
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
							// line 144: if not self.in_simple:
							πF.SetLineno(144)
						Label1:
							// line 145: self.output.append(self.indent*self.level)
							πF.SetLineno(145)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							// line 146: self.output.append(node.endtag())
							πF.SetLineno(146)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßendtag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßFixedTextElement, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßliteral, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
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
								goto Label3
							}
							goto Label4
							// line 147: if isinstance(node, (nodes.FixedTextElement, nodes.literal)):
							πF.SetLineno(147)
						Label3:
							// line 148: self.fixed_text -= 1
							πF.SetLineno(148)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfixed_text, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfixed_text, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsimple_nodes, nil); πE != nil {
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
								goto Label5
							}
							goto Label6
							// line 149: if isinstance(node, self.simple_nodes):
							πF.SetLineno(149)
						Label5:
							// line 150: self.in_simple -= 1
							πF.SetLineno(150)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßin_simple, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßin_simple, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßin_simple, nil); πE != nil {
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
								goto Label7
							}
							goto Label8
							// line 151: if not self.in_simple:
							πF.SetLineno(151)
						Label7:
							// line 152: self.output.append(self.newline)
							πF.SetLineno(152)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnewline, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßdefault_departure.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 142: """Default node depart method."""
					πF.SetLineno(142)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Default node depart method.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_departure); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 158: def visit_Text(self, node):
					πF.SetLineno(158)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("visit_Text", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
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
							// line 159: text = xml.sax.saxutils.escape(node.astext())
							πF.SetLineno(159)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßxml); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsax, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsaxutils, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßescape, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfixed_text, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 161: if not self.fixed_text:
							πF.SetLineno(161)
						Label1:
							// line 162: text = text.replace('\n', '\n'+self.indent*self.level)
							πF.SetLineno(162)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp003
							goto Label2
						Label2:
							// line 163: self.output.append(text)
							πF.SetLineno(163)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_Text.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 165: def depart_Text(self, node):
					πF.SetLineno(165)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("depart_Text", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 166: pass
							πF.SetLineno(166)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_Text.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 168: def visit_raw(self, node):
					πF.SetLineno(168)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("visit_raw", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µxml_string *πg.Object = πg.UnboundLocal
						_ = µxml_string
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µcol_num *πg.Object = πg.UnboundLocal
						_ = µcol_num
						var µline_num *πg.Object = πg.UnboundLocal
						_ = µline_num
						var µsrcline *πg.Object = πg.UnboundLocal
						_ = µsrcline
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
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
							case 7:
								goto Label7
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
							if πTemp005, πE = πg.Contains(πF, πTemp004, ßxml.ToObject()); πE != nil {
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
							// line 169: if 'xml' not in node.get('format', '').split():
							πF.SetLineno(169)
						Label1:
							// line 172: self.default_visit(node)
							πF.SetLineno(172)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_visit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 173: return
							πF.SetLineno(173)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 175: self.default_visit(node)      # or not?
							πF.SetLineno(175)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_visit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 176: xml_string = node.astext()
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µxml_string = πTemp003
							// line 177: self.output.append(xml_string)
							πF.SetLineno(177)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µxml_string, "xml_string"); πE != nil {
								continue
							}
							πTemp002[0] = µxml_string
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoutput, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 178: self.default_departure(node)  # or not?
							πF.SetLineno(178)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_departure, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µxml_string, "xml_string"); πE != nil {
								continue
							}
							πTemp002[0] = µxml_string
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
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
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp003, πE = πg.LT(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 180: if isinstance(xml_string, unicode) and sys.version_info < (3, 0):
							πF.SetLineno(180)
						Label4:
							// line 181: xml_string = xml_string.encode('utf8')
							πF.SetLineno(181)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßutf8.ToObject()
							if πE = πg.CheckLocal(πF, µxml_string, "xml_string"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µxml_string, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µxml_string = πTemp003
							goto Label5
						Label5:
							// line 182: try:
							πF.SetLineno(182)
							πF.PushCheckpoint(7)
							// line 183: self.xmlparser.parse(StringIO(xml_string))
							πF.SetLineno(183)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µxml_string, "xml_string"); πE != nil {
								continue
							}
							πTemp007[0] = µxml_string
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßxmlparser, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßparse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßxml); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsax, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ß_exceptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSAXParseException, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 184: except xml.sax._exceptions.SAXParseException as error:
							πF.SetLineno(184)
						Label8:
							µerror = πTemp008.ToObject()
							// line 185: col_num = self.the_handle.locator.getColumnNumber()
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßthe_handle, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlocator, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßgetColumnNumber, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcol_num = πTemp003
							// line 186: line_num =  self.the_handle.locator.getLineNumber()
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßthe_handle, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlocator, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßgetLineNumber, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline_num = πTemp003
							// line 187: srcline = node.line
							πF.SetLineno(187)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßline, nil); πE != nil {
								continue
							}
							µsrcline = πTemp001
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
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTextElement, nil); πE != nil {
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
								goto Label9
							}
							goto Label10
							// line 188: if not isinstance(node.parent, nodes.TextElement):
							πF.SetLineno(188)
						Label9:
							// line 189: srcline += 2 # directive content start line
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µsrcline, "srcline"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µsrcline, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							µsrcline = πTemp001
							goto Label10
						Label10:
							// line 190: msg = 'Invalid raw XML in column %d, line offset %d:\n%s' % (
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µcol_num, "col_num"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline_num, "line_num"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(µcol_num, µline_num, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid raw XML in column %d, line offset %d:\n%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µmsg = πTemp001
							// line 192: self.warn(msg, source=node.source, line=srcline+line_num-1)
							πF.SetLineno(192)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsrcline, "srcline"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline_num, "line_num"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µsrcline, µline_num); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"source", πTemp001},
								{"line", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwarn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 193: raise nodes.SkipNode # content already processed
							πF.SetLineno(193)
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
					if πE = πClass.SetItem(πF, ßvisit_raw.ToObject(), πTemp008); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("XMLTranslator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßXMLTranslator.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 196: class TestXml(xml.sax.handler.ContentHandler):
			πF.SetLineno(196)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßxml); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßsax, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp007, ßhandler, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßContentHandler, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestXml", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 198: def setDocumentLocator(self, locator):
					πF.SetLineno(198)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "locator", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setDocumentLocator", "/usr/lib/python2.7/site-packages/docutils/writers/docutils_xml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlocator *πg.Object = πArgs[1]
						_ = µlocator
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
							// line 199: self.locator = locator
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µlocator, "locator"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlocator); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocator, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetDocumentLocator.ToObject(), πTemp001); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestXml").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestXml.ToObject(), πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("docutils_xml", Code)
}
