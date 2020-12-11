package html5_polyglot

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß1 := πg.InternStr("1")
		ßFalse := πg.InternStr("False")
		ßHTMLTranslator := πg.InternStr("HTMLTranslator")
		ßNone := πg.InternStr("None")
		ßTrue := πg.InternStr("True")
		ßWriter := πg.InternStr("Writer")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__file__ := πg.InternStr("__file__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_html_base := πg.InternStr("_html_base")
		ßabbr := πg.InternStr("abbr")
		ßabspath := πg.InternStr("abspath")
		ßaction := πg.InternStr("action")
		ßadd_meta := πg.InternStr("add_meta")
		ßappend := πg.InternStr("append")
		ßastext := πg.InternStr("astext")
		ßattval := πg.InternStr("attval")
		ßauthors := πg.InternStr("authors")
		ßbody := πg.InternStr("body")
		ßbrackets := πg.InternStr("brackets")
		ßchoices := πg.InternStr("choices")
		ßcompact_field_lists := πg.InternStr("compact_field_lists")
		ßcompact_lists := πg.InternStr("compact_lists")
		ßconfig_section := πg.InternStr("config_section")
		ßcopyright := πg.InternStr("copyright")
		ßdash := πg.InternStr("dash")
		ßdate := πg.InternStr("date")
		ßdefault := πg.InternStr("default")
		ßdefault_stylesheet_dirs := πg.InternStr("default_stylesheet_dirs")
		ßdefault_stylesheets := πg.InternStr("default_stylesheets")
		ßdefault_template := πg.InternStr("default_template")
		ßdefault_template_path := πg.InternStr("default_template_path")
		ßdepart_acronym := πg.InternStr("depart_acronym")
		ßdepart_authors := πg.InternStr("depart_authors")
		ßdepart_copyright := πg.InternStr("depart_copyright")
		ßdepart_date := πg.InternStr("depart_date")
		ßdepart_docinfo_item := πg.InternStr("depart_docinfo_item")
		ßdepart_meta := πg.InternStr("depart_meta")
		ßdepart_organization := πg.InternStr("depart_organization")
		ßdest := πg.InternStr("dest")
		ßdirname := πg.InternStr("dirname")
		ßdocutils := πg.InternStr("docutils")
		ßembed_stylesheet := πg.InternStr("embed_stylesheet")
		ßemptytag := πg.InternStr("emptytag")
		ßfrontend := πg.InternStr("frontend")
		ßhasattr := πg.InternStr("hasattr")
		ßhtml := πg.InternStr("html")
		ßhtml4 := πg.InternStr("html4")
		ßhtml5 := πg.InternStr("html5")
		ßio := πg.InternStr("io")
		ßjoin := πg.InternStr("join")
		ßlang := πg.InternStr("lang")
		ßmeta := πg.InternStr("meta")
		ßmetavar := πg.InternStr("metavar")
		ßnodes := πg.InternStr("nodes")
		ßnon_default_attributes := πg.InternStr("non_default_attributes")
		ßnone := πg.InternStr("none")
		ßorganization := πg.InternStr("organization")
		ßos := πg.InternStr("os")
		ßoverrides := πg.InternStr("overrides")
		ßparens := πg.InternStr("parens")
		ßparentheses := πg.InternStr("parentheses")
		ßparts := πg.InternStr("parts")
		ßpath := πg.InternStr("path")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsplit := πg.InternStr("split")
		ßstarttag := πg.InternStr("starttag")
		ßstore_false := πg.InternStr("store_false")
		ßstore_true := πg.InternStr("store_true")
		ßstylesheet := πg.InternStr("stylesheet")
		ßstylesheet_path := πg.InternStr("stylesheet_path")
		ßsuperscript := πg.InternStr("superscript")
		ßsupported := πg.InternStr("supported")
		ßtranslator_class := πg.InternStr("translator_class")
		ßtrim_footnote_reference_space := πg.InternStr("trim_footnote_reference_space")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidate_comma_separated_list := πg.InternStr("validate_comma_separated_list")
		ßvalidator := πg.InternStr("validator")
		ßvisit_acronym := πg.InternStr("visit_acronym")
		ßvisit_authors := πg.InternStr("visit_authors")
		ßvisit_copyright := πg.InternStr("visit_copyright")
		ßvisit_date := πg.InternStr("visit_date")
		ßvisit_docinfo_item := πg.InternStr("visit_docinfo_item")
		ßvisit_meta := πg.InternStr("visit_meta")
		ßvisit_organization := πg.InternStr("visit_organization")
		ßwriter_aux := πg.InternStr("writer_aux")
		ßwriters := πg.InternStr("writers")
		ßxhtml := πg.InternStr("xhtml")
		ßxhtml10 := πg.InternStr("xhtml10")
		ßxml_declaration := πg.InternStr("xml_declaration")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
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
			// line 20: """
			πF.SetLineno(20)
			// line 20: """
			πF.SetLineno(20)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nPlain HyperText Markup Language document tree Writer.\n\nThe output conforms to the `HTML 5` specification.\n\nThe cascading style sheet \"minimal.css\" is required for proper viewing,\nthe style sheet \"plain.css\" improves reading experience.\n").ToObject()); πE != nil {
				continue
			}
			// line 28: __docformat__ = 'reStructuredText'
			πF.SetLineno(28)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 30: import os.path
			πF.SetLineno(30)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: import docutils
			πF.SetLineno(31)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 32: from docutils import frontend, nodes, writers, io
			πF.SetLineno(32)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßfrontend.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßwriters.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: from docutils.transforms import writer_aux
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.writer_aux"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßwriter_aux.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: from docutils.writers import _html_base
			πF.SetLineno(34)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers._html_base"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ß_html_base.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 36: class Writer(writers._html_base.Writer):
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß_html_base, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Dict
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
					// line 38: supported = ('html', 'html5', 'html4', 'xhtml', 'xhtml10')
					πF.SetLineno(38)
					πTemp001 = πg.NewTuple5(ßhtml.ToObject(), ßhtml5.ToObject(), ßhtml4.ToObject(), ßxhtml.ToObject(), ßxhtml10.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 39: """Formats this writer supports."""
					πF.SetLineno(39)
					// line 39: """Formats this writer supports."""
					πF.SetLineno(39)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Formats this writer supports.").ToObject()); πE != nil {
						continue
					}
					// line 41: default_stylesheets = ['minimal.css', 'plain.css']
					πF.SetLineno(41)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr("minimal.css").ToObject()
					πTemp002[1] = πg.NewStr("plain.css").ToObject()
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßdefault_stylesheets.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 42: default_stylesheet_dirs = ['.', os.path.abspath(os.path.dirname(__file__))]
					πF.SetLineno(42)
					πTemp002 = make([]*πg.Object, 2)
					πTemp002[0] = πg.NewStr(".").ToObject()
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[1] = πTemp005
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßdefault_stylesheet_dirs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 44: default_template = 'template.txt'
					πF.SetLineno(44)
					if πE = πClass.SetItem(πF, ßdefault_template.ToObject(), πg.NewStr("template.txt").ToObject()); πE != nil {
						continue
					}
					// line 45: default_template_path = os.path.join(
					πF.SetLineno(45)
					πTemp002 = πF.MakeArgs(2)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[0] = πTemp005
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template); πE != nil {
						continue
					}
					πTemp002[1] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßdefault_template_path.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 48: settings_spec = (
					πF.SetLineno(48)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002 = make([]*πg.Object, 18)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template_path); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Mod(πF, πg.NewStr("Specify the template file (UTF-8 encoded).  Default is \"%s\".").ToObject(), πTemp009); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--template").ToObject()
					πTemp009 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template_path); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πTemp011); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file>").ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πTemp008, πTemp009, πTemp011).ToObject()
					πTemp002[0] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--stylesheet").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL[,URL,...]>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßoverrides.ToObject(), ßstylesheet_path.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßvalidate_comma_separated_list, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Comma separated list of stylesheet URLs. Overrides previous --stylesheet and --stylesheet-path settings.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[1] = πTemp007
					πTemp003 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheets); πE != nil {
						continue
					}
					πTemp003[0] = πTemp009
					if πTemp009, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp008, πE = πg.Mod(πF, πg.NewStr("Comma separated list of stylesheet paths. Relative paths are expanded if a matching file is found in the --stylesheet-dirs. With --link-stylesheet, the path is rewritten relative to the output HTML file. Default: \"%s\"").ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--stylesheet-path").ToObject()
					πTemp009 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file[,file,...]>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßoverrides.ToObject(), ßstylesheet.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßvalidate_comma_separated_list, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheets); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp011 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πTemp008, πTemp009, πTemp011).ToObject()
					πTemp002[2] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--embed-stylesheet").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Embed the stylesheet(s) in the output HTML file.  The stylesheet files must be accessible during processing. This is the default.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[3] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--link-stylesheet").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdest.ToObject(), ßembed_stylesheet.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Link to the stylesheet(s) in the output HTML file. Default: embed stylesheets.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[4] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet_dirs); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Mod(πF, πg.NewStr("Comma-separated list of directories where stylesheets are found. Used by --stylesheet-path when expanding relative path arguments. Default: \"%s\"").ToObject(), πTemp009); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--stylesheet-dirs").ToObject()
					πTemp009 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<dir[,dir,...]>").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßvalidate_comma_separated_list, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet_dirs); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp011 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πTemp008, πTemp009, πTemp011).ToObject()
					πTemp002[5] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--initial-header-level").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πTemp009, πE = πg.GetAttr(πF, πg.NewStr("1 2 3 4 5 6").ToObject(), ßsplit, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp009.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßchoices.ToObject(), πTemp011); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), ß1.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify the initial header level.  Default is 1 for \"<h1>\".  Does not affect document title & subtitle (see --no-doc-title).").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[6] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--footnote-references").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ßsuperscript.ToObject()
					πTemp003[1] = ßbrackets.ToObject()
					πTemp009 = πg.NewList(πTemp003...).ToObject()
					if πE = πTemp010.SetItem(πF, ßchoices.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), ßbrackets.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<format>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßoverrides.ToObject(), ßtrim_footnote_reference_space.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Format for footnote references: one of \"superscript\" or \"brackets\".  Default is \"brackets\".").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[7] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--attribution").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = ßdash.ToObject()
					πTemp003[1] = ßparentheses.ToObject()
					πTemp003[2] = ßparens.ToObject()
					πTemp003[3] = ßnone.ToObject()
					πTemp009 = πg.NewList(πTemp003...).ToObject()
					if πE = πTemp010.SetItem(πF, ßchoices.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), ßdash.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<format>").ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Format for block quote attributions: one of \"dash\" (em-dash prefix), \"parentheses\"/\"parens\", or \"none\".  Default is \"dash\".").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[8] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--compact-lists").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Remove extra vertical whitespace between items of \"simple\" bullet lists and enumerated lists.  Default: enabled.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[9] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--no-compact-lists").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdest.ToObject(), ßcompact_lists.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable compact simple bullet and enumerated lists.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[10] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--compact-field-lists").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Remove extra vertical whitespace between items of simple field lists.  Default: enabled.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[11] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--no-compact-field-lists").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdest.ToObject(), ßcompact_field_lists.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable compact simple field lists.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[12] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--table-style").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Added to standard table classes. Defined styles: borderless, booktabs, align-left, align-center, align-right, colwidths-auto. Default: \"\"").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[13] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--math-output").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πg.NewStr("HTML math.css").ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Math output format (one of \"MathML\", \"HTML\", \"MathJax\", or \"LaTeX\") and option(s). Default: \"HTML math.css\"").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[14] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--xml-declaration").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßdefault.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Prepend an XML declaration. (Thwarts HTML5 conformance.) Default: False").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[15] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--no-xml-declaration").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßdest.ToObject(), ßxml_declaration.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Omit the XML declaration.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[16] = πTemp007
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--cloak-email-addresses").ToObject()
					πTemp008 = πg.NewList(πTemp003...).ToObject()
					πTemp010 = πg.NewDict()
					if πE = πTemp010.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp010.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp009 = πTemp010.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Obfuscate email addresses to confuse harvesters while still keeping email links usable with standards-compliant browsers.").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp002[17] = πTemp007
					πTemp006 = πg.NewTuple(πTemp002...).ToObject()
					πTemp001 = πg.NewTuple3(πg.NewStr("HTML-Specific Options").ToObject(), πTemp005, πTemp006).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 141: config_section = 'html5 writer'
					πF.SetLineno(141)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("html5 writer").ToObject()); πE != nil {
						continue
					}
					// line 143: def __init__(self):
					πF.SetLineno(143)
					πTemp013 = make([]πg.Param, 1)
					πTemp013[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 144: self.parts = {}
							πF.SetLineno(144)
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
							// line 145: self.translator_class = HTMLTranslator
							πF.SetLineno(145)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHTMLTranslator); πE != nil {
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
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 148: class HTMLTranslator(writers._html_base.HTMLTranslator):
			πF.SetLineno(148)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßwriters); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß_html_base, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßHTMLTranslator, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("HTMLTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 149: """
					πF.SetLineno(149)
					// line 149: """
					πF.SetLineno(149)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    This writer generates `polyglot markup`: HTML5 that is also valid XML.\n\n    Safe subclassing: when overriding, treat ``visit_*`` and ``depart_*``\n    methods as a unit to prevent breaks due to internal changes. See the\n    docstring of docutils.writers._html_base.HTMLTranslator for details\n    and examples.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 159: def visit_acronym(self, node):
					πF.SetLineno(159)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("visit_acronym", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 161: self.body.append(self.starttag(node, 'abbr', ''))
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							πTemp002[1] = ßabbr.ToObject()
							πTemp002[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßvisit_acronym.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 162: def depart_acronym(self, node):
					πF.SetLineno(162)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("depart_acronym", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 163: self.body.append('</abbr>')
							πF.SetLineno(163)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</abbr>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_acronym.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 167: def visit_authors(self, node):
					πF.SetLineno(167)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("visit_authors", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µsubnode *πg.Object = πg.UnboundLocal
						_ = µsubnode
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 168: self.visit_docinfo_item(node, 'authors', meta=False)
							πF.SetLineno(168)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßauthors.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"meta", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µnode); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µsubnode = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 170: self.add_meta('<meta name="author" content="%s" />\n' %
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubnode, "subnode"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µsubnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßattval, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("<meta name=\"author\" content=\"%s\" />\n").ToObject(), πTemp009); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßadd_meta, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_authors.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 172: def depart_authors(self, node):
					πF.SetLineno(172)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("depart_authors", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 173: self.depart_docinfo_item()
							πF.SetLineno(173)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_authors.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 177: def visit_copyright(self, node):
					πF.SetLineno(177)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("visit_copyright", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 178: self.visit_docinfo_item(node, 'copyright', meta=False)
							πF.SetLineno(178)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßcopyright.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"meta", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 179: self.add_meta('<meta name="dcterms.rights" content="%s" />\n'
							πF.SetLineno(179)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßattval, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<meta name=\"dcterms.rights\" content=\"%s\" />\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_meta, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_copyright.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 181: def depart_copyright(self, node):
					πF.SetLineno(181)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("depart_copyright", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 182: self.depart_docinfo_item()
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_copyright.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 185: def visit_date(self, node):
					πF.SetLineno(185)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("visit_date", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 186: self.visit_docinfo_item(node, 'date', meta=False)
							πF.SetLineno(186)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßdate.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"meta", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 187: self.add_meta('<meta name="dcterms.date" content="%s" />\n'
							πF.SetLineno(187)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßattval, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<meta name=\"dcterms.date\" content=\"%s\" />\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_meta, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_date.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 189: def depart_date(self, node):
					πF.SetLineno(189)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("depart_date", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 190: self.depart_docinfo_item()
							πF.SetLineno(190)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_date.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 202: def visit_meta(self, node):
					πF.SetLineno(202)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("visit_meta", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µmeta *πg.Object = πg.UnboundLocal
						_ = µmeta
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßlang.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
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
							// line 203: if node.hasattr('lang'):
							πF.SetLineno(203)
						Label1:
							// line 204: node['xml:lang'] = node['lang']
							πF.SetLineno(204)
							πTemp002 = ßlang.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005 = πg.NewStr("xml:lang").ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 206: meta = self.emptytag(node, 'meta', **node.non_default_attributes())
							πF.SetLineno(206)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßmeta.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßnon_default_attributes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßemptytag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmeta = πTemp005
							// line 207: self.add_meta(meta)
							πF.SetLineno(207)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmeta, "meta"); πE != nil {
								continue
							}
							πTemp001[0] = µmeta
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_meta, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_meta.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 208: def depart_meta(self, node):
					πF.SetLineno(208)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("depart_meta", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 209: pass
							πF.SetLineno(209)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdepart_meta.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 212: def visit_organization(self, node):
					πF.SetLineno(212)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("visit_organization", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 213: self.visit_docinfo_item(node, 'organization', meta=False)
							πF.SetLineno(213)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp001[1] = ßorganization.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"meta", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvisit_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_organization.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 214: def depart_organization(self, node):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("depart_organization", "/usr/lib/python2.7/site-packages/docutils/writers/html5_polyglot/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 215: self.depart_docinfo_item()
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdepart_docinfo_item, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_organization.ToObject(), πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("HTMLTranslator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHTMLTranslator.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("html5_polyglot", Code)
}
