package pep_html

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/codecs"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/random"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßHTMLTranslator := πg.InternStr("HTMLTranslator")
		ßSUPPRESS_HELP := πg.InternStr("SUPPRESS_HELP")
		ßValueError := πg.InternStr("ValueError")
		ßWriter := πg.InternStr("Writer")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__file__ := πg.InternStr("__file__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßaction := πg.InternStr("action")
		ßappend := πg.InternStr("append")
		ßassemble_parts := πg.InternStr("assemble_parts")
		ßastext := πg.InternStr("astext")
		ßbanner := πg.InternStr("banner")
		ßbody := πg.InternStr("body")
		ßbody_pre_docinfo := πg.InternStr("body_pre_docinfo")
		ßclasses := πg.InternStr("classes")
		ßcodecs := πg.InternStr("codecs")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßdefault := πg.InternStr("default")
		ßdefault_stylesheet := πg.InternStr("default_stylesheet")
		ßdefault_stylesheet_path := πg.InternStr("default_stylesheet_path")
		ßdefault_template := πg.InternStr("default_template")
		ßdefault_template_path := πg.InternStr("default_template_path")
		ßdepart_field_list := πg.InternStr("depart_field_list")
		ßdirname := πg.InternStr("dirname")
		ßdocinfo := πg.InternStr("docinfo")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßdummy := πg.InternStr("dummy")
		ßfield_list := πg.InternStr("field_list")
		ßfirst_child_matching_class := πg.InternStr("first_child_matching_class")
		ßfrontend := πg.InternStr("frontend")
		ßgetcwd := πg.InternStr("getcwd")
		ßhtml4css1 := πg.InternStr("html4css1")
		ßint := πg.InternStr("int")
		ßinterpolation_dict := πg.InternStr("interpolation_dict")
		ßjoin := πg.InternStr("join")
		ßmetavar := πg.InternStr("metavar")
		ßno_random := πg.InternStr("no_random")
		ßnodes := πg.InternStr("nodes")
		ßos := πg.InternStr("os")
		ßparts := πg.InternStr("parts")
		ßpath := πg.InternStr("path")
		ßpep := πg.InternStr("pep")
		ßpep_home := πg.InternStr("pep_home")
		ßpephome := πg.InternStr("pephome")
		ßpepindex := πg.InternStr("pepindex")
		ßpepnum := πg.InternStr("pepnum")
		ßpyhome := πg.InternStr("pyhome")
		ßpython_home := πg.InternStr("python_home")
		ßrandrange := πg.InternStr("randrange")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrelative_path := πg.InternStr("relative_path")
		ßrelative_path_settings := πg.InternStr("relative_path_settings")
		ßrfc2822 := πg.InternStr("rfc2822")
		ßsettings := πg.InternStr("settings")
		ßsettings_default_overrides := πg.InternStr("settings_default_overrides")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßstore_true := πg.InternStr("store_true")
		ßstylesheet_path := πg.InternStr("stylesheet_path")
		ßsys := πg.InternStr("sys")
		ßtemplate := πg.InternStr("template")
		ßtitle := πg.InternStr("title")
		ßtranslator_class := πg.InternStr("translator_class")
		ßutils := πg.InternStr("utils")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidator := πg.InternStr("validator")
		ßwriters := πg.InternStr("writers")
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
			// line 5: """
			πF.SetLineno(5)
			// line 5: """
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nPEP HTML Writer.\n").ToObject()); πE != nil {
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
			// line 15: import codecs
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "codecs"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcodecs.ToObject(), πTemp001); πE != nil {
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
			// line 17: from docutils import frontend, nodes, utils, writers
			πF.SetLineno(17)
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßwriters.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: from docutils.writers import html4css1
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers.html4css1"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßhtml4css1.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: class Writer(html4css1.Writer):
			πF.SetLineno(21)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
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
				var πTemp011 *πg.Dict
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
				var πTemp017 []πg.Param
				_ = πTemp017
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 23: default_stylesheet = 'pep.css'
					πF.SetLineno(23)
					// line 23: default_stylesheet = 'pep.css'
					πF.SetLineno(23)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("pep.css").ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdefault_stylesheet.ToObject(), πg.NewStr("pep.css").ToObject()); πE != nil {
						continue
					}
					// line 25: default_stylesheet_path = utils.relative_path(
					πF.SetLineno(25)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					πTemp002[1] = ßdummy.ToObject()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp002 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp004
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßutils); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrelative_path, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßdefault_stylesheet_path.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 29: default_template = 'template.txt'
					πF.SetLineno(29)
					if πE = πClass.SetItem(πF, ßdefault_template.ToObject(), πg.NewStr("template.txt").ToObject()); πE != nil {
						continue
					}
					// line 31: default_template_path = utils.relative_path(
					πF.SetLineno(31)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßgetcwd, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					πTemp002[1] = ßdummy.ToObject()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp002 = πF.MakeArgs(2)
					πTemp005 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ß__file__); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßdirname, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[1] = πTemp004
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßutils); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrelative_path, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßdefault_template_path.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 35: settings_spec = html4css1.Writer.settings_spec + (
					πF.SetLineno(35)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßhtml4css1); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßWriter, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßsettings_spec, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet_path); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template_path); πE != nil {
						continue
					}
					πTemp008 = πg.NewTuple2(πTemp009, πTemp010).ToObject()
					if πTemp007, πE = πg.Mod(πF, πg.NewStr("For the PEP/HTML writer, the default value for the --stylesheet-path option is \"%s\", and the default value for --template is \"%s\". See HTML-Specific Options above.").ToObject(), πTemp008); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("--python-home").ToObject()
					πTemp010 = πg.NewList(πTemp001...).ToObject()
					πTemp011 = πg.NewDict()
					if πE = πTemp011.SetItem(πF, ßdefault.ToObject(), πg.NewStr("http://www.python.org").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp011.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					πTemp012 = πTemp011.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Python's home URL.  Default is \"http://www.python.org\".").ToObject(), πTemp010, πTemp012).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("--pep-home").ToObject()
					πTemp012 = πg.NewList(πTemp001...).ToObject()
					πTemp011 = πg.NewDict()
					if πE = πTemp011.SetItem(πF, ßdefault.ToObject(), πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp011.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					πTemp013 = πTemp011.ToObject()
					πTemp010 = πg.NewTuple3(πg.NewStr("Home URL prefix for PEPs.  Default is \".\" (current directory).").ToObject(), πTemp012, πTemp013).ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßSUPPRESS_HELP, nil); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("--no-random").ToObject()
					πTemp013 = πg.NewList(πTemp001...).ToObject()
					πTemp011 = πg.NewDict()
					if πE = πTemp011.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp011.SetItem(πF, ßvalidator.ToObject(), πTemp016); πE != nil {
						continue
					}
					πTemp015 = πTemp011.ToObject()
					πTemp012 = πg.NewTuple3(πTemp014, πTemp013, πTemp015).ToObject()
					πTemp008 = πg.NewTuple3(πTemp009, πTemp010, πTemp012).ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("PEP/HTML-Specific Options").ToObject(), πTemp007, πTemp008).ToObject()
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 52: settings_default_overrides = {'stylesheet_path': default_stylesheet_path,
					πF.SetLineno(52)
					πTemp011 = πg.NewDict()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_stylesheet_path); πE != nil {
						continue
					}
					if πE = πTemp011.SetItem(πF, ßstylesheet_path.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template_path); πE != nil {
						continue
					}
					if πE = πTemp011.SetItem(πF, ßtemplate.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp003 = πTemp011.ToObject()
					if πE = πClass.SetItem(πF, ßsettings_default_overrides.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 55: relative_path_settings = ('template',)
					πF.SetLineno(55)
					πTemp003 = πg.NewTuple1(ßtemplate.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßrelative_path_settings.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 57: config_section = 'pep_html writer'
					πF.SetLineno(57)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("pep_html writer").ToObject()); πE != nil {
						continue
					}
					// line 58: config_section_dependencies = ('writers', 'html writers',
					πF.SetLineno(58)
					πTemp003 = πg.NewTuple3(ßwriters.ToObject(), πg.NewStr("html writers").ToObject(), πg.NewStr("html4css1 writer").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 61: def __init__(self):
					πF.SetLineno(61)
					πTemp017 = make([]πg.Param, 1)
					πTemp017[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", πTemp017, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 62: html4css1.Writer.__init__(self)
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
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
							// line 63: self.translator_class = HTMLTranslator
							πF.SetLineno(63)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 65: def interpolation_dict(self):
					πF.SetLineno(65)
					πTemp017 = make([]πg.Param, 1)
					πTemp017[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("interpolation_dict", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", πTemp017, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsubs *πg.Object = πg.UnboundLocal
						_ = µsubs
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µpyhome *πg.Object = πg.UnboundLocal
						_ = µpyhome
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
						var µrandom *πg.Object = πg.UnboundLocal
						_ = µrandom
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
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
							case 8:
								goto Label8
							default:
								panic("unexpected function state")
							}
							// line 66: subs = html4css1.Writer.interpolation_dict(self)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßWriter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßinterpolation_dict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsubs = πTemp003
							// line 67: settings = self.document.settings
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							µsettings = πTemp003
							// line 68: pyhome = settings.python_home
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßpython_home, nil); πE != nil {
								continue
							}
							µpyhome = πTemp002
							// line 69: subs['pyhome'] = pyhome
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µpyhome, "pyhome"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µpyhome); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp003 = ßpyhome.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 70: subs['pephome'] = settings.pep_home
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßpep_home, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßpephome.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpyhome, "pyhome"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µpyhome, πg.NewStr("..").ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 71: if pyhome == '..':
							πF.SetLineno(71)
						Label1:
							// line 72: subs['pepindex'] = '.'
							πF.SetLineno(72)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr(".").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp003 = ßpepindex.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 74: subs['pepindex'] = pyhome + '/dev/peps'
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µpyhome, "pyhome"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µpyhome, πg.NewStr("/dev/peps").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßpepindex.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 75: index = self.document.first_child_matching_class(nodes.field_list)
							πF.SetLineno(75)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfield_list, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfirst_child_matching_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp002
							// line 76: header = self.document[index]
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µheader = πTemp003
							// line 77: self.pepnum = header[0][1].astext()
							πF.SetLineno(77)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µheader, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßastext, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßpepnum, πTemp002); πE != nil {
								continue
							}
							// line 78: subs['pep'] = self.pepnum
							πF.SetLineno(78)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpepnum, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßpep.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßno_random, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 79: if settings.no_random:
							πF.SetLineno(79)
						Label4:
							// line 80: subs['banner'] = 0
							πF.SetLineno(80)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp003 = ßbanner.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label5:
							// line 82: import random
							πF.SetLineno(82)
							if πTemp001, πE = πg.ImportModule(πF, "random"); πE != nil {
								continue
							}
							πTemp002 = πTemp001[0]
							µrandom = πTemp002
							// line 83: subs['banner'] = random.randrange(64)
							πF.SetLineno(83)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(64).ToObject()
							if πE = πg.CheckLocal(πF, µrandom, "random"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µrandom, ßrandrange, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßbanner.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
							// line 84: try:
							πF.SetLineno(84)
							πF.PushCheckpoint(8)
							// line 85: subs['pepnum'] = '%04i' % int(self.pepnum)
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpepnum, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%04i").ToObject(), πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßpepnum.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 86: except ValueError:
							πF.SetLineno(86)
						Label9:
							// line 87: subs['pepnum'] = self.pepnum
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpepnum, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßpepnum.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp003); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 88: self.title = header[1][1].astext()
							πF.SetLineno(88)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µheader, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßastext, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßtitle, πTemp002); πE != nil {
								continue
							}
							// line 89: subs['title'] = self.title
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtitle, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßtitle.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 90: subs['body'] = ''.join(
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody_pre_docinfo, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
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
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πTemp004 = ßbody.ToObject()
							if πE = πg.SetItem(πF, µsubs, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 92: return subs
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µsubs, "subs"); πE != nil {
								continue
							}
							πR = µsubs
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßinterpolation_dict.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 94: def assemble_parts(self):
					πF.SetLineno(94)
					πTemp017 = make([]πg.Param, 1)
					πTemp017[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("assemble_parts", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", πTemp017, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 95: html4css1.Writer.assemble_parts(self)
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßWriter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßassemble_parts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 96: self.parts['title'] = [self.title]
							πF.SetLineno(96)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtitle, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßparts, nil); πE != nil {
								continue
							}
							πTemp005 = ßtitle.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 97: self.parts['pepnum'] = self.pepnum
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpepnum, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßparts, nil); πE != nil {
								continue
							}
							πTemp005 = ßpepnum.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
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
			// line 100: class HTMLTranslator(html4css1.HTMLTranslator):
			πF.SetLineno(100)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßHTMLTranslator, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("HTMLTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 102: def depart_field_list(self, node):
					πF.SetLineno(102)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("depart_field_list", "/usr/lib/python2.7/site-packages/docutils/writers/pep_html/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 103: html4css1.HTMLTranslator.depart_field_list(self, node)
							πF.SetLineno(103)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[1] = µnode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßHTMLTranslator, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßdepart_field_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, ßrfc2822.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 104: if 'rfc2822' in node['classes']:
							πF.SetLineno(104)
						Label1:
							// line 105: self.body.append('<hr />\n')
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<hr />\n").ToObject()
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
					if πE = πClass.SetItem(πF, ßdepart_field_list.ToObject(), πTemp001); πE != nil {
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
	πg.RegisterModule("pep_html", Code)
}
