package s5_html

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßApplicationError := πg.InternStr("ApplicationError")
		ßFalse := πg.InternStr("False")
		ßHTMLTranslator := πg.InternStr("HTMLTranslator")
		ßNone := πg.InternStr("None")
		ßS5HTMLTranslator := πg.InternStr("S5HTMLTranslator")
		ßTrue := πg.InternStr("True")
		ßWriter := πg.InternStr("Writer")
		ß__base__ := πg.InternStr("__base__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__file__ := πg.InternStr("__file__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_destination := πg.InternStr("_destination")
		ßaction := πg.InternStr("action")
		ßadd := πg.InternStr("add")
		ßadd_meta := πg.InternStr("add_meta")
		ßappend := πg.InternStr("append")
		ßbase_theme_file := πg.InternStr("base_theme_file")
		ßbody := πg.InternStr("body")
		ßbody_pre_docinfo := πg.InternStr("body_pre_docinfo")
		ßbody_prefix := πg.InternStr("body_prefix")
		ßbody_suffix := πg.InternStr("body_suffix")
		ßchoices := πg.InternStr("choices")
		ßclasses := πg.InternStr("classes")
		ßclose := πg.InternStr("close")
		ßcompile := πg.InternStr("compile")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßcontent_type := πg.InternStr("content_type")
		ßcontext := πg.InternStr("context")
		ßcontrol_visibility := πg.InternStr("control_visibility")
		ßcopy_file := πg.InternStr("copy_file")
		ßcopy_theme := πg.InternStr("copy_theme")
		ßcurrent_slide := πg.InternStr("current_slide")
		ßdefault := πg.InternStr("default")
		ßdefault_theme := πg.InternStr("default_theme")
		ßdepart_document := πg.InternStr("depart_document")
		ßdepart_footer := πg.InternStr("depart_footer")
		ßdepart_header := πg.InternStr("depart_header")
		ßdest := πg.InternStr("dest")
		ßdirect_theme_files := πg.InternStr("direct_theme_files")
		ßdirectives := πg.InternStr("directives")
		ßdirname := πg.InternStr("dirname")
		ßdisable_current_slide := πg.InternStr("disable_current_slide")
		ßdiv := πg.InternStr("div")
		ßdocinfo := πg.InternStr("docinfo")
		ßdoctype := πg.InternStr("doctype")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßdummy := πg.InternStr("dummy")
		ßencode := πg.InternStr("encode")
		ßexists := πg.InternStr("exists")
		ßextend := πg.InternStr("extend")
		ßfiles_to_skip_pattern := πg.InternStr("files_to_skip_pattern")
		ßfind_theme := πg.InternStr("find_theme")
		ßfooter := πg.InternStr("footer")
		ßfragment := πg.InternStr("fragment")
		ßfrontend := πg.InternStr("frontend")
		ßgetcwd := πg.InternStr("getcwd")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßhead := πg.InternStr("head")
		ßhead_prefix := πg.InternStr("head_prefix")
		ßhead_prefix_template := πg.InternStr("head_prefix_template")
		ßheader := πg.InternStr("header")
		ßhidden := πg.InternStr("hidden")
		ßhidden_controls := πg.InternStr("hidden_controls")
		ßhtml4css1 := πg.InternStr("html4css1")
		ßhtml_body := πg.InternStr("html_body")
		ßhtml_head := πg.InternStr("html_head")
		ßhtml_prolog := πg.InternStr("html_prolog")
		ßhtml_title := πg.InternStr("html_title")
		ßids := πg.InternStr("ids")
		ßindirect_theme_files := πg.InternStr("indirect_theme_files")
		ßinitial_header_level := πg.InternStr("initial_header_level")
		ßinsert := πg.InternStr("insert")
		ßisdir := πg.InternStr("isdir")
		ßisfile := πg.InternStr("isfile")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlang := πg.InternStr("lang")
		ßlanguage_code := πg.InternStr("language_code")
		ßlayout_template := πg.InternStr("layout_template")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlistdir := πg.InternStr("listdir")
		ßmakedirs := πg.InternStr("makedirs")
		ßmath_header := πg.InternStr("math_header")
		ßmath_output := πg.InternStr("math_output")
		ßmathjax := πg.InternStr("mathjax")
		ßmeta := πg.InternStr("meta")
		ßmetavar := πg.InternStr("metavar")
		ßnodes := πg.InternStr("nodes")
		ßopen := πg.InternStr("open")
		ßos := πg.InternStr("os")
		ßoutline := πg.InternStr("outline")
		ßoutput_encoding := πg.InternStr("output_encoding")
		ßoverrides := πg.InternStr("overrides")
		ßoverwrite_theme_files := πg.InternStr("overwrite_theme_files")
		ßparent := πg.InternStr("parent")
		ßpath := πg.InternStr("path")
		ßpop := πg.InternStr("pop")
		ßrb := πg.InternStr("rb")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßreadlines := πg.InternStr("readlines")
		ßrecord_dependencies := πg.InternStr("record_dependencies")
		ßrelative_path := πg.InternStr("relative_path")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßrequired_theme_files := πg.InternStr("required_theme_files")
		ßrfind := πg.InternStr("rfind")
		ßs5_footer := πg.InternStr("s5_footer")
		ßs5_header := πg.InternStr("s5_header")
		ßs5_stylesheet_template := πg.InternStr("s5_stylesheet_template")
		ßsearch := πg.InternStr("search")
		ßsection := πg.InternStr("section")
		ßsection_count := πg.InternStr("section_count")
		ßsection_level := πg.InternStr("section_level")
		ßsep := πg.InternStr("sep")
		ßsettings := πg.InternStr("settings")
		ßsettings_default_overrides := πg.InternStr("settings_default_overrides")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsetup_theme := πg.InternStr("setup_theme")
		ßslide := πg.InternStr("slide")
		ßslide0 := πg.InternStr("slide0")
		ßslideshow := πg.InternStr("slideshow")
		ßstartswith := πg.InternStr("startswith")
		ßstarttag := πg.InternStr("starttag")
		ßstore_false := πg.InternStr("store_false")
		ßstore_true := πg.InternStr("store_true")
		ßstrip := πg.InternStr("strip")
		ßstylesheet := πg.InternStr("stylesheet")
		ßsys := πg.InternStr("sys")
		ßtheme := πg.InternStr("theme")
		ßtheme_file_path := πg.InternStr("theme_file_path")
		ßtheme_files_copied := πg.InternStr("theme_files_copied")
		ßtheme_url := πg.InternStr("theme_url")
		ßthemes := πg.InternStr("themes")
		ßthemes_dir_path := πg.InternStr("themes_dir_path")
		ßtitle := πg.InternStr("title")
		ßtoc_backlinks := πg.InternStr("toc_backlinks")
		ßtranslator_class := πg.InternStr("translator_class")
		ßui := πg.InternStr("ui")
		ßutils := πg.InternStr("utils")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidator := πg.InternStr("validator")
		ßview_mode := πg.InternStr("view_mode")
		ßvisible := πg.InternStr("visible")
		ßvisit_section := πg.InternStr("visit_section")
		ßvisit_subtitle := πg.InternStr("visit_subtitle")
		ßvisit_title := πg.InternStr("visit_title")
		ßwb := πg.InternStr("wb")
		ßwrite := πg.InternStr("write")
		ßwriters := πg.InternStr("writers")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []*πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 []*πg.Object
		_ = πTemp005
		var πTemp006 []πg.Param
		_ = πTemp006
		var πTemp007 *πg.Dict
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
			// line 6: """
			πF.SetLineno(6)
			// line 6: """
			πF.SetLineno(6)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nS5/HTML Slideshow Writer.\n").ToObject()); πE != nil {
				continue
			}
			// line 10: __docformat__ = 'reStructuredText'
			πF.SetLineno(10)
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
			// line 14: import os
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
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
			// line 16: import docutils
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: from docutils import frontend, nodes, utils
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
			// line 18: from docutils.writers import html4css1
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers.html4css1"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßhtml4css1.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: from docutils.parsers.rst import directives
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: themes_dir_path = utils.relative_path(
			πF.SetLineno(21)
			πTemp002 = πF.MakeArgs(2)
			πTemp003 = πF.MakeArgs(2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßgetcwd, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp003[0] = πTemp001
			πTemp003[1] = ßdummy.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp003)
			πTemp002[0] = πTemp004
			πTemp003 = πF.MakeArgs(2)
			πTemp005 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß__file__); πE != nil {
				continue
			}
			πTemp005[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßdirname, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp005)
			πTemp003[0] = πTemp004
			πTemp003[1] = ßthemes.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp003)
			πTemp002[1] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßrelative_path, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßthemes_dir_path.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: def find_theme(name):
			πF.SetLineno(25)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "name", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("find_theme", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
				var µpath *πg.Object = πg.UnboundLocal
				_ = µpath
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
					// line 28: path = os.path.join(themes_dir_path, name)
					πF.SetLineno(28)
					πTemp001 = πF.MakeArgs(2)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßthemes_dir_path); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[1] = µname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpath = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp001[0] = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßisdir, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 29: if not os.path.isdir(path):
					πF.SetLineno(29)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µname, µpath).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Theme directory not found: %r (path: %r)").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßApplicationError, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 30: raise docutils.ApplicationError(
					πF.SetLineno(30)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label2
				Label2:
					// line 32: return path
					πF.SetLineno(32)
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
			if πE = πF.Globals().SetItem(πF, ßfind_theme.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: class Writer(html4css1.Writer):
			πF.SetLineno(35)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßWriter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp007 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Dict
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []πg.Param
				_ = πTemp012
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 37: settings_spec = html4css1.Writer.settings_spec + (
					πF.SetLineno(37)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßhtml4css1); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßWriter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsettings_spec, nil); πE != nil {
						continue
					}
					πTemp005 = make([]*πg.Object, 9)
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--theme").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßdefault.ToObject(), ßdefault.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<name>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßoverrides.ToObject(), ßtheme_url.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Specify an installed S5 theme by name.  Overrides --theme-url.  The default theme name is \"default\".  The theme files will be copied into a \"ui/<theme>\" directory, in the same directory as the destination file (output HTML).  Note that existing theme files will not be overwritten (unless --overwrite-theme-files is used).").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[0] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--theme-url").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßoverrides.ToObject(), ßtheme.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Specify an S5 theme URL.  The destination file (output HTML) will link to this theme; nothing will be copied.  Overrides --theme.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[1] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--overwrite-theme-files").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Allow existing theme files in the ``ui/<theme>`` directory to be overwritten.  The default is not to overwrite theme files.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[2] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--keep-theme-files").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßdest.ToObject(), ßoverwrite_theme_files.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Keep existing theme files in the ``ui/<theme>`` directory; do not overwrite any.  This is the default.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[3] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--view-mode").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					πTemp007 = make([]*πg.Object, 2)
					πTemp007[0] = ßslideshow.ToObject()
					πTemp007[1] = ßoutline.ToObject()
					πTemp010 = πg.NewList(πTemp007...).ToObject()
					if πE = πTemp009.SetItem(πF, ßchoices.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßdefault.ToObject(), ßslideshow.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<mode>").ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Set the initial view mode to \"slideshow\" [default] or \"outline\".").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[4] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--hidden-controls").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßdefault.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Normally hide the presentation controls in slideshow mode. This is the default.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[5] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--visible-controls").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßdest.ToObject(), ßhidden_controls.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Always show the presentation controls in slideshow mode.  The default is to hide the controls.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[6] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--current-slide").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Enable the current slide indicator (\"1 / 15\").  The default is to disable it.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[7] = πTemp006
					πTemp007 = make([]*πg.Object, 1)
					πTemp007[0] = πg.NewStr("--no-current-slide").ToObject()
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßdest.ToObject(), ßcurrent_slide.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp009.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp009.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Disable the current slide indicator.  This is the default.").ToObject(), πTemp008, πTemp010).ToObject()
					πTemp005[8] = πTemp006
					πTemp004 = πg.NewTuple(πTemp005...).ToObject()
					πTemp003 = πg.NewTuple3(πg.NewStr("S5 Slideshow Specific Options").ToObject(), πg.NewStr("For the S5/HTML writer, the --no-toc-backlinks option (defined in General Docutils Options above) is the default, and should not be changed.").ToObject(), πTemp004).ToObject()
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 83: settings_default_overrides = {'toc_backlinks': 0}
					πF.SetLineno(83)
					πTemp009 = πg.NewDict()
					if πE = πTemp009.SetItem(πF, ßtoc_backlinks.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp009.ToObject()
					if πE = πClass.SetItem(πF, ßsettings_default_overrides.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 85: config_section = 's5_html writer'
					πF.SetLineno(85)
					// line 85: config_section = 's5_html writer'
					πF.SetLineno(85)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("s5_html writer").ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("s5_html writer").ToObject()); πE != nil {
						continue
					}
					// line 86: config_section_dependencies = ('writers', 'html writers',
					πF.SetLineno(86)
					πTemp001 = πg.NewTuple3(ßwriters.ToObject(), πg.NewStr("html writers").ToObject(), πg.NewStr("html4css1 writer").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 89: def __init__(self):
					πF.SetLineno(89)
					πTemp012 = make([]πg.Param, 1)
					πTemp012[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 90: html4css1.Writer.__init__(self)
							πF.SetLineno(90)
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
							// line 91: self.translator_class = S5HTMLTranslator
							πF.SetLineno(91)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßS5HTMLTranslator); πE != nil {
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
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 94: class S5HTMLTranslator(html4css1.HTMLTranslator):
			πF.SetLineno(94)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßHTMLTranslator, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp007 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("S5HTMLTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				var πTemp001 *πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 96: s5_stylesheet_template = """\
					πF.SetLineno(96)
					// line 96: s5_stylesheet_template = """\
					πF.SetLineno(96)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("<!-- configuration parameters -->\n<meta name=\"defaultView\" content=\"%(view_mode)s\" />\n<meta name=\"controlVis\" content=\"%(control_visibility)s\" />\n<!-- style sheet links -->\n<script src=\"%(path)s/slides.js\" type=\"text/javascript\"></script>\n<link rel=\"stylesheet\" href=\"%(path)s/slides.css\"\n      type=\"text/css\" media=\"projection\" id=\"slideProj\" />\n<link rel=\"stylesheet\" href=\"%(path)s/outline.css\"\n      type=\"text/css\" media=\"screen\" id=\"outlineStyle\" />\n<link rel=\"stylesheet\" href=\"%(path)s/print.css\"\n      type=\"text/css\" media=\"print\" id=\"slidePrint\" />\n<link rel=\"stylesheet\" href=\"%(path)s/opera.css\"\n      type=\"text/css\" media=\"projection\" id=\"operaFix\" />\n").ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßs5_stylesheet_template.ToObject(), πg.NewStr("<!-- configuration parameters -->\n<meta name=\"defaultView\" content=\"%(view_mode)s\" />\n<meta name=\"controlVis\" content=\"%(control_visibility)s\" />\n<!-- style sheet links -->\n<script src=\"%(path)s/slides.js\" type=\"text/javascript\"></script>\n<link rel=\"stylesheet\" href=\"%(path)s/slides.css\"\n      type=\"text/css\" media=\"projection\" id=\"slideProj\" />\n<link rel=\"stylesheet\" href=\"%(path)s/outline.css\"\n      type=\"text/css\" media=\"screen\" id=\"outlineStyle\" />\n<link rel=\"stylesheet\" href=\"%(path)s/print.css\"\n      type=\"text/css\" media=\"print\" id=\"slidePrint\" />\n<link rel=\"stylesheet\" href=\"%(path)s/opera.css\"\n      type=\"text/css\" media=\"projection\" id=\"operaFix\" />\n").ToObject()); πE != nil {
						continue
					}
					// line 114: disable_current_slide = """
					πF.SetLineno(114)
					if πE = πClass.SetItem(πF, ßdisable_current_slide.ToObject(), πg.NewStr("\n<style type=\"text/css\">\n#currentSlide {display: none;}\n</style>\n").ToObject()); πE != nil {
						continue
					}
					// line 119: layout_template = """\
					πF.SetLineno(119)
					if πE = πClass.SetItem(πF, ßlayout_template.ToObject(), πg.NewStr("<div class=\"layout\">\n<div id=\"controls\"></div>\n<div id=\"currentSlide\"></div>\n<div id=\"header\">\n%(header)s\n</div>\n<div id=\"footer\">\n%(title)s%(footer)s\n</div>\n</div>\n").ToObject()); πE != nil {
						continue
					}
					// line 135: default_theme = 'default'
					πF.SetLineno(135)
					if πE = πClass.SetItem(πF, ßdefault_theme.ToObject(), ßdefault.ToObject()); πE != nil {
						continue
					}
					// line 136: """Name of the default theme."""
					πF.SetLineno(136)
					// line 138: base_theme_file = '__base__'
					πF.SetLineno(138)
					if πE = πClass.SetItem(πF, ßbase_theme_file.ToObject(), ß__base__.ToObject()); πE != nil {
						continue
					}
					// line 139: """Name of the file containing the name of the base theme."""
					πF.SetLineno(139)
					// line 141: direct_theme_files = (
					πF.SetLineno(141)
					πTemp001 = πg.NewTuple5(πg.NewStr("slides.css").ToObject(), πg.NewStr("outline.css").ToObject(), πg.NewStr("print.css").ToObject(), πg.NewStr("opera.css").ToObject(), πg.NewStr("slides.js").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßdirect_theme_files.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 143: """Names of theme files directly linked to in the output HTML"""
					πF.SetLineno(143)
					// line 145: indirect_theme_files = (
					πF.SetLineno(145)
					πTemp001 = πg.NewTuple5(πg.NewStr("s5-core.css").ToObject(), πg.NewStr("framing.css").ToObject(), πg.NewStr("pretty.css").ToObject(), πg.NewStr("blank.gif").ToObject(), πg.NewStr("iepngfix.htc").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßindirect_theme_files.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 147: """Names of files used indirectly; imported or used by files in
					πF.SetLineno(147)
					// line 150: required_theme_files = indirect_theme_files + direct_theme_files
					πF.SetLineno(150)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßindirect_theme_files); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdirect_theme_files); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßrequired_theme_files.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 151: """Names of mandatory theme files."""
					πF.SetLineno(151)
					// line 153: def __init__(self, *args):
					πF.SetLineno(153)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µview_mode *πg.Object = πg.UnboundLocal
						_ = µview_mode
						var µcontrol_visibility *πg.Object = πg.UnboundLocal
						_ = µcontrol_visibility
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
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
							// line 154: html4css1.HTMLTranslator.__init__(self, *args)
							πF.SetLineno(154)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhtml4css1); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßHTMLTranslator, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 156: self.theme_file_path = None
							πF.SetLineno(156)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtheme_file_path, πTemp003); πE != nil {
								continue
							}
							// line 157: self.setup_theme()
							πF.SetLineno(157)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_theme, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 158: view_mode = self.document.settings.view_mode
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßview_mode, nil); πE != nil {
								continue
							}
							µview_mode = πTemp002
							// line 159: control_visibility = ('visible', 'hidden')[self.document.settings
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßhidden_controls, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							πTemp004 = πg.NewTuple2(ßvisible.ToObject(), ßhidden.ToObject()).ToObject()
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µcontrol_visibility = πTemp003
							// line 161: self.stylesheet.append(self.s5_stylesheet_template
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßs5_stylesheet_template, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtheme_file_path, nil); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßpath.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µview_mode, "view_mode"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßview_mode.ToObject(), µview_mode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontrol_visibility, "control_visibility"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßcontrol_visibility.ToObject(), µcontrol_visibility); πE != nil {
								continue
							}
							πTemp004 = πTemp005.ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstylesheet, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßcurrent_slide, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
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
							// line 165: if not self.document.settings.current_slide:
							πF.SetLineno(165)
						Label1:
							// line 166: self.stylesheet.append(self.disable_current_slide)
							πF.SetLineno(166)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdisable_current_slide, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstylesheet, nil); πE != nil {
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
							// line 167: self.add_meta('<meta name="version" content="S5 1.1" />\n')
							πF.SetLineno(167)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<meta name=\"version\" content=\"S5 1.1\" />\n").ToObject()
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
							// line 168: self.s5_footer = []
							πF.SetLineno(168)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßs5_footer, πTemp003); πE != nil {
								continue
							}
							// line 169: self.s5_header = []
							πF.SetLineno(169)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßs5_header, πTemp003); πE != nil {
								continue
							}
							// line 170: self.section_count = 0
							πF.SetLineno(170)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_count, πTemp002); πE != nil {
								continue
							}
							// line 171: self.theme_files_copied = None
							πF.SetLineno(171)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtheme_files_copied, πTemp003); πE != nil {
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
					// line 173: def setup_theme(self):
					πF.SetLineno(173)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("setup_theme", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtheme, nil); πE != nil {
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtheme_url, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 174: if self.document.settings.theme:
							πF.SetLineno(174)
						Label1:
							// line 175: self.copy_theme()
							πF.SetLineno(175)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcopy_theme, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label4
							// line 176: elif self.document.settings.theme_url:
							πF.SetLineno(176)
						Label2:
							// line 177: self.theme_file_path = self.document.settings.theme_url
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtheme_url, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtheme_file_path, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label3:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("No theme specified for S5/HTML writer.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßApplicationError, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 179: raise docutils.ApplicationError(
							πF.SetLineno(179)
							πE = πF.Raise(πTemp001, nil, nil)
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
					if πE = πClass.SetItem(πF, ßsetup_theme.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 182: def copy_theme(self):
					πF.SetLineno(182)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("copy_theme", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µpath *πg.Object = πg.UnboundLocal
						_ = µpath
						var µtheme_paths *πg.Object = πg.UnboundLocal
						_ = µtheme_paths
						var µrequired_files_copied *πg.Object = πg.UnboundLocal
						_ = µrequired_files_copied
						var µdest *πg.Object = πg.UnboundLocal
						_ = µdest
						var µdefault *πg.Object = πg.UnboundLocal
						_ = µdefault
						var µf *πg.Object = πg.UnboundLocal
						_ = µf
						var µbase_theme_file *πg.Object = πg.UnboundLocal
						_ = µbase_theme_file
						var µlines *πg.Object = πg.UnboundLocal
						_ = µlines
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µrequired *πg.Object = πg.UnboundLocal
						_ = µrequired
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
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
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.BaseException
						_ = πTemp013
						var πTemp014 *πg.Traceback
						_ = πTemp014
						var πTemp015 *πg.Type
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
						var πTemp017 *πg.Object
						_ = πTemp017
						var πTemp018 *πg.Object
						_ = πTemp018
						var πTemp019 []πg.Param
						_ = πTemp019
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 36:
								goto Label36
							case 37:
								goto Label37
							case 6:
								goto Label6
							case 7:
								goto Label7
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 22:
								goto Label22
							case 23:
								goto Label23
							case 24:
								goto Label24
							default:
								panic("unexpected function state")
							}
							// line 183: """
							πF.SetLineno(183)
							// line 190: settings = self.document.settings
							πF.SetLineno(190)
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
							// line 191: path = find_theme(settings.theme)
							πF.SetLineno(191)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßtheme, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßfind_theme); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µpath = πTemp002
							// line 192: theme_paths = [path]
							πF.SetLineno(192)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[0] = µpath
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µtheme_paths = πTemp001
							// line 193: self.theme_files_copied = {}
							πF.SetLineno(193)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtheme_files_copied, πTemp002); πE != nil {
								continue
							}
							// line 194: required_files_copied = {}
							πF.SetLineno(194)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							µrequired_files_copied = πTemp001
							// line 196: self.theme_file_path = '%s/%s' % ('ui', settings.theme)
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µsettings, ßtheme, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßui.ToObject(), πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s/%s").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtheme_file_path, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ß_destination, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 197: if settings._destination:
							πF.SetLineno(197)
						Label1:
							// line 198: dest = os.path.join(
							πF.SetLineno(198)
							πTemp003 = πF.MakeArgs(3)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ß_destination, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßui.ToObject()
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ßtheme, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdest = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[0] = µdest
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßisdir, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 200: if not os.path.isdir(dest):
							πF.SetLineno(200)
						Label4:
							// line 201: os.makedirs(dest)
							πF.SetLineno(201)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[0] = µdest
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmakedirs, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label5
						Label5:
							goto Label3
						Label2:
							// line 204: return
							πF.SetLineno(204)
							πR = πg.None
							continue
							goto Label3
						Label3:
							// line 205: default = False
							πF.SetLineno(205)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µdefault = πTemp001
							// line 206: while path:
							πF.SetLineno(206)
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
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µpath); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[0] = µpath
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßlistdir, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp008 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label11
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
								µf = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(9)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbase_theme_file, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µf, πTemp005); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label12
							}
							goto Label13
							// line 208: if f == self.base_theme_file:
							πF.SetLineno(208)
						Label12:
							// line 209: continue            # ... except the "__base__" file
							πF.SetLineno(209)
							continue
							goto Label13
						Label13:
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[1] = µpath
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[2] = µdest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcopy_file, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πTemp010
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßrequired_theme_files, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp010, µf); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp011).ToObject()
							πTemp002 = πTemp005
						Label14:
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label15
							}
							goto Label16
							// line 210: if ( self.copy_file(f, path, dest)
							πF.SetLineno(210)
						Label15:
							// line 212: required_files_copied[f] = 1
							πF.SetLineno(212)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrequired_files_copied, "required_files_copied"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp005 = µf
							if πE = πg.SetItem(πF, µrequired_files_copied, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label16
						Label16:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µdefault); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label17
							}
							goto Label18
							// line 213: if default:
							πF.SetLineno(213)
						Label17:
							// line 214: break                   # "default" theme has no base theme
							πF.SetLineno(214)
							πTemp006 = true
							continue
							goto Label18
						Label18:
							// line 216: base_theme_file = os.path.join(path, self.base_theme_file)
							πF.SetLineno(216)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[0] = µpath
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbase_theme_file, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µbase_theme_file = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbase_theme_file, "base_theme_file"); πE != nil {
								continue
							}
							πTemp003[0] = µbase_theme_file
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßisfile, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label19
							}
							goto Label20
							// line 218: if os.path.isfile(base_theme_file):
							πF.SetLineno(218)
						Label19:
							// line 219: with open(base_theme_file) as f:
							πF.SetLineno(219)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbase_theme_file, "base_theme_file"); πE != nil {
								continue
							}
							πTemp003[0] = µbase_theme_file
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(22)
							µf = πTemp005
							// line 220: lines = f.readlines()
							πF.SetLineno(220)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µf, ßreadlines, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlines = πTemp012
							πF.PopCheckpoint()
						Label22:
							πTemp013, πTemp014 = nil, nil
							if πE != nil {
								πTemp013, πTemp014 = πF.ExcInfo()
							}
							if πTemp013 != nil {
								πTemp015 = πTemp013.Type()
								if πTemp010, πE = πTemp001.Call(πF, πg.Args{πTemp002, πTemp015.ToObject(), πTemp013.ToObject(), πTemp014.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp010, πE = πTemp001.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp013 != nil && πTemp008 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Iter(πF, µlines); πE != nil {
								continue
							}
							πF.PushCheckpoint(24)
							πTemp008 = false
						Label23:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label25
							}
							if πTemp012, πE = πg.Next(πF, πTemp010); πE != nil {
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
								µline = πTemp012
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(23)
							// line 222: line = line.strip()
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							µline = πTemp016
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp012 = µline
							if πTemp009, πE = πg.IsTrue(πF, πTemp012); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label26
							}
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("#").ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetAttr(πF, µline, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp018, πE = πTemp017.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp011, πE = πg.IsTrue(πF, πTemp018); πE != nil {
								continue
							}
							πTemp016 = πg.GetBool(!πTemp011).ToObject()
							πTemp012 = πTemp016
						Label26:
							if πTemp009, πE = πg.IsTrue(πF, πTemp012); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label27
							}
							goto Label28
							// line 223: if line and not line.startswith('#'):
							πF.SetLineno(223)
						Label27:
							// line 224: path = find_theme(line)
							πF.SetLineno(224)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp003[0] = µline
							if πTemp012, πE = πg.ResolveGlobal(πF, ßfind_theme); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp012.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µpath = πTemp016
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtheme_paths, "theme_paths"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µtheme_paths, µpath); πE != nil {
								continue
							}
							πTemp012 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp012); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label29
							}
							goto Label30
							// line 225: if path in theme_paths: # check for duplicates (cycles)
							πF.SetLineno(225)
						Label29:
							// line 226: path = None         # if found, use default base
							πF.SetLineno(226)
							if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µpath = πTemp012
							goto Label31
						Label30:
							// line 228: theme_paths.append(path)
							πF.SetLineno(228)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[0] = µpath
							if πE = πg.CheckLocal(πF, µtheme_paths, "theme_paths"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µtheme_paths, ßappend, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp012.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label31
						Label31:
							// line 229: break
							πF.SetLineno(229)
							πTemp008 = true
							continue
							goto Label28
						Label28:
							continue
						Label24:
							if πE != nil || πR != nil {
								continue
							}
							// line 231: path = None         # use default base
							πF.SetLineno(231)
							if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µpath = πTemp012
						Label25:
							goto Label21
						Label20:
							// line 233: path = None             # use default base
							πF.SetLineno(233)
							if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µpath = πTemp010
							goto Label21
						Label21:
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µpath); πE != nil {
								continue
							}
							πTemp010 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label32
							}
							goto Label33
							// line 234: if not path:
							πF.SetLineno(234)
						Label32:
							// line 235: path = find_theme(self.default_theme)
							πF.SetLineno(235)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßdefault_theme, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πTemp010, πE = πg.ResolveGlobal(πF, ßfind_theme); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µpath = πTemp012
							// line 236: theme_paths.append(path)
							πF.SetLineno(236)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							πTemp003[0] = µpath
							if πE = πg.CheckLocal(πF, µtheme_paths, "theme_paths"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µtheme_paths, ßappend, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 237: default = True
							πF.SetLineno(237)
							if πTemp010, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µdefault = πTemp010
							goto Label33
						Label33:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrequired_files_copied, "required_files_copied"); πE != nil {
								continue
							}
							πTemp003[0] = µrequired_files_copied
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp012.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßrequired_theme_files, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp012
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp017, πE = πTemp012.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp010, πE = πg.NE(πF, πTemp016, πTemp017); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label34
							}
							goto Label35
							// line 238: if len(required_files_copied) != len(self.required_theme_files):
							πF.SetLineno(238)
						Label34:
							// line 240: required = list(self.required_theme_files)
							πF.SetLineno(240)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßrequired_theme_files, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πTemp010, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µrequired = πTemp012
							if πE = πg.CheckLocal(πF, µrequired_files_copied, "required_files_copied"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µrequired_files_copied, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Iter(πF, πTemp016); πE != nil {
								continue
							}
							πF.PushCheckpoint(37)
							πTemp006 = false
						Label36:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label38
							}
							if πTemp012, πE = πg.Next(πF, πTemp010); πE != nil {
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
								µf = πTemp012
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(36)
							// line 242: required.remove(f)
							πF.SetLineno(242)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							if πE = πg.CheckLocal(πF, µrequired, "required"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µrequired, ßremove, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp012.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label37:
							if πE != nil || πR != nil {
								continue
							}
						Label38:
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp019 = make([]πg.Param, 0)
							πTemp016 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp019, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µf *πg.Object = πg.UnboundLocal
								_ = µf
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
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µrequired, "required"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µrequired); πE != nil {
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
											µf = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 245: % ', '.join(['%r' % f for f in required]))
										πF.SetLineno(245)
										if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Mod(πF, πg.NewStr("%r").ToObject(), µf); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp005 = πSent
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
							if πTemp017, πE = πTemp016.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ListType.Call(πF, πg.Args{πTemp017}, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp012
							if πTemp012, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp017, πE = πTemp012.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp010, πE = πg.Mod(πF, πg.NewStr("Theme files not found: %s").ToObject(), πTemp017); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πTemp010, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp010, ßApplicationError, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp012.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 243: raise docutils.ApplicationError(
							πF.SetLineno(243)
							πE = πF.Raise(πTemp010, nil, nil)
							continue
							goto Label35
						Label35:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy_theme.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 183: """
					πF.SetLineno(183)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Locate & copy theme files.\n\n        A theme may be explicitly based on another theme via a '__base__'\n        file.  The default base theme is 'default'.  Files are accumulated\n        from the specified theme, any base themes, and 'default'.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_theme); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 247: files_to_skip_pattern = re.compile(r'~$|\.bak$|#$|\.cvsignore$')
					πF.SetLineno(247)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr("~$|\\.bak$|#$|\\.cvsignore$").ToObject()
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πClass.SetItem(πF, ßfiles_to_skip_pattern.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 249: def copy_file(self, name, source_dir, dest_dir):
					πF.SetLineno(249)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "name", Def: nil}
					πTemp004[2] = πg.Param{Name: "source_dir", Def: nil}
					πTemp004[3] = πg.Param{Name: "dest_dir", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("copy_file", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var µsource_dir *πg.Object = πArgs[2]
						_ = µsource_dir
						var µdest_dir *πg.Object = πArgs[3]
						_ = µdest_dir
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µdest *πg.Object = πg.UnboundLocal
						_ = µdest
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µsrc_file *πg.Object = πg.UnboundLocal
						_ = µsrc_file
						var µsrc_data *πg.Object = πg.UnboundLocal
						_ = µsrc_data
						var µdest_file *πg.Object = πg.UnboundLocal
						_ = µdest_file
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
						var πTemp008 []*πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 250: """
							πF.SetLineno(250)
							// line 254: source = os.path.join(source_dir, name)
							πF.SetLineno(254)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsource_dir, "source_dir"); πE != nil {
								continue
							}
							πTemp001[0] = µsource_dir
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsource = πTemp003
							// line 255: dest = os.path.join(dest_dir, name)
							πF.SetLineno(255)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdest_dir, "dest_dir"); πE != nil {
								continue
							}
							πTemp001[0] = µdest_dir
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdest = πTemp003
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtheme_files_copied, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, µdest); πE != nil {
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
							// line 256: if dest in self.theme_files_copied:
							πF.SetLineno(256)
						Label1:
							// line 257: return 1
							πF.SetLineno(257)
							πR = πg.NewInt(1).ToObject()
							continue
							goto Label3
						Label2:
							// line 259: self.theme_files_copied[dest] = 1
							πF.SetLineno(259)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtheme_files_copied, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp005 = µdest
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[0] = µsource
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisfile, nil); πE != nil {
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
							// line 260: if os.path.isfile(source):
							πF.SetLineno(260)
						Label4:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[0] = µsource
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfiles_to_skip_pattern, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 261: if self.files_to_skip_pattern.search(source):
							πF.SetLineno(261)
						Label6:
							// line 262: return None
							πF.SetLineno(262)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label7
						Label7:
							// line 263: settings = self.document.settings
							πF.SetLineno(263)
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp001[0] = µdest
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßexists, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µsettings, ßoverwrite_theme_files, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
						Label8:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 264: if os.path.exists(dest) and not settings.overwrite_theme_files:
							πF.SetLineno(264)
						Label9:
							// line 265: settings.record_dependencies.add(dest)
							πF.SetLineno(265)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp001[0] = µdest
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label11
						Label10:
							// line 267: src_file = open(source, 'rb')
							πF.SetLineno(267)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[0] = µsource
							πTemp001[1] = ßrb.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsrc_file = πTemp003
							// line 268: src_data = src_file.read()
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µsrc_file, "src_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsrc_file, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsrc_data = πTemp003
							// line 269: src_file.close()
							πF.SetLineno(269)
							if πE = πg.CheckLocal(πF, µsrc_file, "src_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsrc_file, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 270: dest_file = open(dest, 'wb')
							πF.SetLineno(270)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp001[0] = µdest
							πTemp001[1] = ßwb.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdest_file = πTemp003
							// line 271: dest_dir = dest_dir.replace(os.sep, '/')
							πF.SetLineno(271)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsep, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("/").ToObject()
							if πE = πg.CheckLocal(πF, µdest_dir, "dest_dir"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdest_dir, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdest_dir = πTemp003
							// line 272: dest_file.write(src_data.replace(b'ui/default',
							πF.SetLineno(272)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							πTemp007[0] = πg.NewStr("ui/default").ToObject()
							πTemp008 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetfilesystemencoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp002
							πTemp009 = πF.MakeArgs(1)
							πTemp009[0] = πg.NewStr("ui/").ToObject()
							if πE = πg.CheckLocal(πF, µdest_dir, "dest_dir"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdest_dir, ßrfind, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdest_dir, "dest_dir"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdest_dir, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[1] = πTemp003
							if πE = πg.CheckLocal(πF, µsrc_data, "src_data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsrc_data, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µdest_file, "dest_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdest_file, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 275: dest_file.close()
							πF.SetLineno(275)
							if πE = πg.CheckLocal(πF, µdest_file, "dest_file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdest_file, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 276: settings.record_dependencies.add(source)
							πF.SetLineno(276)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[0] = µsource
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label11
						Label11:
							// line 277: return 1
							πF.SetLineno(277)
							πR = πg.NewInt(1).ToObject()
							continue
							goto Label5
						Label5:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp001[0] = µdest
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisfile, nil); πE != nil {
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
								goto Label12
							}
							goto Label13
							// line 278: if os.path.isfile(dest):
							πF.SetLineno(278)
						Label12:
							// line 279: return 1
							πF.SetLineno(279)
							πR = πg.NewInt(1).ToObject()
							continue
							goto Label13
						Label13:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy_file.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 250: """
					πF.SetLineno(250)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Copy file `name` from `source_dir` to `dest_dir`.\n        Return 1 if the file exists in either `source_dir` or `dest_dir`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_file); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 281: def depart_document(self, node):
					πF.SetLineno(281)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("depart_document", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
						var µfooter *πg.Object = πg.UnboundLocal
						_ = µfooter
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µlayout *πg.Object = πg.UnboundLocal
						_ = µlayout
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 282: self.head_prefix.extend([self.doctype,
							πF.SetLineno(282)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdoctype, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhead_prefix_template, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßlanguage_code, nil); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßlang.ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp006 = πTemp005.ToObject()
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 285: self.html_prolog.append(self.doctype)
							πF.SetLineno(285)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdoctype, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_prolog, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 286: self.meta.insert(0, self.content_type % self.settings.output_encoding)
							πF.SetLineno(286)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent_type, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmeta, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 287: self.head.insert(0, self.content_type % self.settings.output_encoding)
							πF.SetLineno(287)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontent_type, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßoutput_encoding, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 288: if self.math_header:
							πF.SetLineno(288)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmath_output, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, ßmathjax.ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							goto Label4
							// line 289: if self.math_output == 'mathjax':
							πF.SetLineno(289)
						Label3:
							// line 290: self.head.extend(self.math_header)
							πF.SetLineno(290)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label4:
							// line 292: self.stylesheet.extend(self.math_header)
							πF.SetLineno(292)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmath_header, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstylesheet, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 294: self.html_head.extend(self.head[1:])
							πF.SetLineno(294)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßhead, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_head, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 295: self.fragment.extend(self.body)
							πF.SetLineno(295)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfragment, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 297: header = ''.join(self.s5_header)
							πF.SetLineno(297)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßs5_header, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µheader = πTemp004
							// line 298: footer = ''.join(self.s5_footer)
							πF.SetLineno(298)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßs5_footer, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfooter = πTemp004
							// line 299: title = ''.join(self.html_title).replace('<h1 class="title">', '<h1>')
							πF.SetLineno(299)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("<h1 class=\"title\">").ToObject()
							πTemp001[1] = πg.NewStr("<h1>").ToObject()
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_title, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtitle = πTemp004
							// line 300: layout = self.layout_template % {'header': header,
							πF.SetLineno(300)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßlayout_template, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßheader.ToObject(), µheader); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßtitle.ToObject(), µtitle); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßfooter.ToObject(), µfooter); πE != nil {
								continue
							}
							πTemp006 = πTemp005.ToObject()
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							µlayout = πTemp003
							// line 303: self.body_prefix.extend(layout)
							πF.SetLineno(303)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlayout, "layout"); πE != nil {
								continue
							}
							πTemp001[0] = µlayout
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 304: self.body_prefix.append('<div class="presentation">\n')
							πF.SetLineno(304)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("<div class=\"presentation\">\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 305: self.body_prefix.append(
							πF.SetLineno(305)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πg.NewDict()
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = ßslide.ToObject()
							πTemp003 = πg.NewList(πTemp009...).ToObject()
							if πE = πTemp005.SetItem(πF, ßclasses.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = ßslide0.ToObject()
							πTemp003 = πg.NewList(πTemp009...).ToObject()
							if πE = πTemp005.SetItem(πF, ßids.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp003 = πTemp005.ToObject()
							πTemp002[0] = πTemp003
							πTemp002[1] = ßdiv.ToObject()
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsection_count, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 307: if not self.section_count:
							πF.SetLineno(307)
						Label6:
							// line 308: self.body.append('</div>\n')
							πF.SetLineno(308)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("</div>\n").ToObject()
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
							goto Label7
						Label7:
							// line 310: self.body_suffix.insert(0, '</div>\n')
							πF.SetLineno(310)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = πg.NewStr("</div>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbody_suffix, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 311: self.html_body.extend(self.body_prefix[1:] + self.body_pre_docinfo
							πF.SetLineno(311)
							πTemp001 = πF.MakeArgs(1)
							if πTemp010, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßbody_prefix, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßbody_pre_docinfo, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp011, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßdocinfo, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßbody_suffix, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhtml_body, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_document.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 315: def depart_footer(self, node):
					πF.SetLineno(315)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("depart_footer", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 316: start = self.context.pop()
							πF.SetLineno(316)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstart = πTemp001
							// line 317: self.s5_footer.append('<h2>')
							πF.SetLineno(317)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("<h2>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßs5_footer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 318: self.s5_footer.extend(self.body[start:])
							πF.SetLineno(318)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßs5_footer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 319: self.s5_footer.append('</h2>')
							πF.SetLineno(319)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("</h2>").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßs5_footer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 320: del self.body[start:]
							πF.SetLineno(320)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_footer.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 322: def depart_header(self, node):
					πF.SetLineno(322)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("depart_header", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 323: start = self.context.pop()
							πF.SetLineno(323)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstart = πTemp001
							// line 324: header = ['<div id="header">\n']
							πF.SetLineno(324)
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewStr("<div id=\"header\">\n").ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µheader = πTemp001
							// line 325: header.extend(self.body[start:])
							πF.SetLineno(325)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µheader, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 326: header.append('\n</div>\n')
							πF.SetLineno(326)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\n</div>\n").ToObject()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µheader, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 327: del self.body[start:]
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 328: self.s5_header.extend(header)
							πF.SetLineno(328)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp003[0] = µheader
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßs5_header, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdepart_header.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 330: def visit_section(self, node):
					πF.SetLineno(330)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("visit_section", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsection_count, nil); πE != nil {
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
							// line 331: if not self.section_count:
							πF.SetLineno(331)
						Label1:
							// line 332: self.body.append('\n</div>\n')
							πF.SetLineno(332)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n</div>\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
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
							// line 333: self.section_count += 1
							πF.SetLineno(333)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsection_count, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_count, πTemp002); πE != nil {
								continue
							}
							// line 334: self.section_level += 1
							πF.SetLineno(334)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsection_level, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 335: if self.section_level > 1:
							πF.SetLineno(335)
						Label3:
							// line 337: self.body.append(self.starttag(node, 'div', CLASS='section'))
							πF.SetLineno(337)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßdiv.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßsection.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label4:
							// line 339: self.body.append(self.starttag(node, 'div', CLASS='slide'))
							πF.SetLineno(339)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp005[1] = ßdiv.ToObject()
							πTemp006 = πg.KWArgs{
								{"CLASS", ßslide.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbody, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_section.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 341: def visit_subtitle(self, node):
					πF.SetLineno(341)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("visit_subtitle", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µlevel *πg.Object = πg.UnboundLocal
						_ = µlevel
						var µtag *πg.Object = πg.UnboundLocal
						_ = µtag
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
						var πTemp007 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsection, nil); πE != nil {
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
							goto Label2
							// line 342: if isinstance(node.parent, nodes.section):
							πF.SetLineno(342)
						Label1:
							// line 343: level = self.section_level + self.initial_header_level - 1
							πF.SetLineno(343)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßsection_level, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßinitial_header_level, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µlevel = πTemp002
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 344: if level == 1:
							πF.SetLineno(344)
						Label4:
							// line 345: level = 2
							πF.SetLineno(345)
							µlevel = πg.NewInt(2).ToObject()
							goto Label5
						Label5:
							// line 346: tag = 'h%s' % level
							πF.SetLineno(346)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("h%s").ToObject(), µlevel); πE != nil {
								continue
							}
							µtag = πTemp002
							// line 347: self.body.append(self.starttag(node, tag, ''))
							πF.SetLineno(347)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp007[0] = µnode
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							πTemp007[1] = µtag
							πTemp007[2] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstarttag, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
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
							// line 348: self.context.append('</%s>\n' % tag)
							πF.SetLineno(348)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("</%s>\n").ToObject(), µtag); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 350: html4css1.HTMLTranslator.visit_subtitle(self, node)
							πF.SetLineno(350)
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßvisit_subtitle, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßvisit_subtitle.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 352: def visit_title(self, node):
					πF.SetLineno(352)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "node", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("visit_title", "/usr/lib/python2.7/site-packages/docutils/writers/s5_html/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 353: html4css1.HTMLTranslator.visit_title(self, node)
							πF.SetLineno(353)
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßvisit_title, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_title.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("S5HTMLTranslator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßS5HTMLTranslator.ToObject(), πTemp009); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("s5_html", Code)
}
