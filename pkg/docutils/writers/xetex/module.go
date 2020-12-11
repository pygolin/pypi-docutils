package xetex

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/re"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßBabel := πg.InternStr("Babel")
		ßLaTeXTranslator := πg.InternStr("LaTeXTranslator")
		ßNone := πg.InternStr("None")
		ßTrue := πg.InternStr("True")
		ßWriter := πg.InternStr("Writer")
		ßXeLaTeXTranslator := πg.InternStr("XeLaTeXTranslator")
		ß__call__ := πg.InternStr("__call__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_inputenc := πg.InternStr("_inputenc")
		ßaf := πg.InternStr("af")
		ßalbanian := πg.InternStr("albanian")
		ßancientgreek := πg.InternStr("ancientgreek")
		ßappend := πg.InternStr("append")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßcop := πg.InternStr("cop")
		ßcoptic := πg.InternStr("coptic")
		ßcopy := πg.InternStr("copy")
		ßcroatian := πg.InternStr("croatian")
		ßde := πg.InternStr("de")
		ßdefault := πg.InternStr("default")
		ßdefault_preamble := πg.InternStr("default_preamble")
		ßdefault_template := πg.InternStr("default_template")
		ßdict := πg.InternStr("dict")
		ßdivehi := πg.InternStr("divehi")
		ßdocutils := πg.InternStr("docutils")
		ßdsb := πg.InternStr("dsb")
		ßdv := πg.InternStr("dv")
		ßfa := πg.InternStr("fa")
		ßfarsi := πg.InternStr("farsi")
		ßfilter_settings_spec := πg.InternStr("filter_settings_spec")
		ßfont_encoding := πg.InternStr("font_encoding")
		ßfontencoding := πg.InternStr("fontencoding")
		ßfrontend := πg.InternStr("frontend")
		ßgerman := πg.InternStr("german")
		ßgrc := πg.InternStr("grc")
		ßhsb := πg.InternStr("hsb")
		ßis_xetex := πg.InternStr("is_xetex")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkey := πg.InternStr("key")
		ßkeys := πg.InternStr("keys")
		ßlanguage := πg.InternStr("language")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguage_codes := πg.InternStr("language_codes")
		ßlanguage_name := πg.InternStr("language_name")
		ßlanguages := πg.InternStr("languages")
		ßlatex2e := πg.InternStr("latex2e")
		ßlatex_encoding := πg.InternStr("latex_encoding")
		ßliteral_double_quote := πg.InternStr("literal_double_quote")
		ßlower := πg.InternStr("lower")
		ßlsorbian := πg.InternStr("lsorbian")
		ßlualatex := πg.InternStr("lualatex")
		ßluatex := πg.InternStr("luatex")
		ßlxtex := πg.InternStr("lxtex")
		ßmetavar := πg.InternStr("metavar")
		ßnodes := πg.InternStr("nodes")
		ßogerman := πg.InternStr("ogerman")
		ßos := πg.InternStr("os")
		ßotherlanguages := πg.InternStr("otherlanguages")
		ßpolygreek := πg.InternStr("polygreek")
		ßpop := πg.InternStr("pop")
		ßquote_index := πg.InternStr("quote_index")
		ßquotes := πg.InternStr("quotes")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrequirements := πg.InternStr("requirements")
		ßserbian := πg.InternStr("serbian")
		ßsettings_defaults := πg.InternStr("settings_defaults")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsorted := πg.InternStr("sorted")
		ßsq := πg.InternStr("sq")
		ßsr := πg.InternStr("sr")
		ßsupported := πg.InternStr("supported")
		ßth := πg.InternStr("th")
		ßthai := πg.InternStr("thai")
		ßtranslator_class := πg.InternStr("translator_class")
		ßupdate := πg.InternStr("update")
		ßusorbian := πg.InternStr("usorbian")
		ßutf8 := πg.InternStr("utf8")
		ßutils := πg.InternStr("utils")
		ßvi := πg.InternStr("vi")
		ßvietnamese := πg.InternStr("vietnamese")
		ßwarn_msg := πg.InternStr("warn_msg")
		ßwriters := πg.InternStr("writers")
		ßxelatex := πg.InternStr("xelatex")
		ßxetex := πg.InternStr("xetex")
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
			// line 17: """
			πF.SetLineno(17)
			// line 17: """
			πF.SetLineno(17)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nXeLaTeX document tree Writer.\n\nA variant of Docutils' standard 'latex2e' writer producing LaTeX output\nsuited for processing with the Unicode-aware TeX engines\nLuaTeX and XeTeX.\n").ToObject()); πE != nil {
				continue
			}
			// line 25: __docformat__ = 'reStructuredText'
			πF.SetLineno(25)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 27: import os
			πF.SetLineno(27)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: import os.path
			πF.SetLineno(28)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: import re
			πF.SetLineno(29)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
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
			// line 32: from docutils import frontend, nodes, utils, writers, languages
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: from docutils.writers import latex2e
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers.latex2e"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßlatex2e.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: class Writer(latex2e.Writer):
			πF.SetLineno(35)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlatex2e); πE != nil {
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
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
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
				var πTemp006 *πg.Dict
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 πg.KWArgs
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 36: """A writer for Unicode-aware LaTeX variants (XeTeX, LuaTeX)"""
					πF.SetLineno(36)
					// line 36: """A writer for Unicode-aware LaTeX variants (XeTeX, LuaTeX)"""
					πF.SetLineno(36)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A writer for Unicode-aware LaTeX variants (XeTeX, LuaTeX)").ToObject()); πE != nil {
						continue
					}
					// line 38: supported = ('lxtex', 'xetex', 'xelatex', 'luatex', 'lualatex')
					πF.SetLineno(38)
					πTemp001 = πg.NewTuple5(ßlxtex.ToObject(), ßxetex.ToObject(), ßxelatex.ToObject(), ßluatex.ToObject(), ßlualatex.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 39: """Formats this writer supports."""
					πF.SetLineno(39)
					// line 41: default_template = 'xelatex.tex'
					πF.SetLineno(41)
					if πE = πClass.SetItem(πF, ßdefault_template.ToObject(), πg.NewStr("xelatex.tex").ToObject()); πE != nil {
						continue
					}
					// line 42: default_preamble = '\n'.join([
					πF.SetLineno(42)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = make([]*πg.Object, 4)
					πTemp003[0] = πg.NewStr("% Linux Libertine (free, wide coverage, not only for Linux)").ToObject()
					πTemp003[1] = πg.NewStr("\\setmainfont{Linux Libertine O}").ToObject()
					πTemp003[2] = πg.NewStr("\\setsansfont{Linux Biolinum O}").ToObject()
					πTemp003[3] = πg.NewStr("\\setmonofont[HyphenChar=None,Scale=MatchLowercase]{DejaVu Sans Mono}").ToObject()
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					πTemp002[0] = πTemp001
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßdefault_preamble.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 49: config_section = 'xetex writer'
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("xetex writer").ToObject()); πE != nil {
						continue
					}
					// line 50: config_section_dependencies = ('writers', 'latex writers',
					πF.SetLineno(50)
					πTemp001 = πg.NewTuple3(ßwriters.ToObject(), πg.NewStr("latex writers").ToObject(), πg.NewStr("latex2e writer").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 53: settings_spec = frontend.filter_settings_spec(
					πF.SetLineno(53)
					πTemp002 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßlatex2e); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßWriter, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßsettings_spec, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					πTemp002[1] = ßfont_encoding.ToObject()
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("Template file. Default: \"%s\".").ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--template").ToObject()
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					πTemp006 = πg.NewDict()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_template); πE != nil {
						continue
					}
					if πE = πTemp006.SetItem(πF, ßdefault.ToObject(), πTemp007); πE != nil {
						continue
					}
					if πE = πTemp006.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file>").ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp006.ToObject()
					πTemp001 = πg.NewTuple3(πTemp004, πTemp005, πTemp007).ToObject()
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = πg.NewStr("--latex-preamble").ToObject()
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					πTemp006 = πg.NewDict()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_preamble); πE != nil {
						continue
					}
					if πE = πTemp006.SetItem(πF, ßdefault.ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp007 = πTemp006.ToObject()
					πTemp004 = πg.NewTuple3(πg.NewStr("Customization by LaTeX code in the preamble. Default: select \"Linux Libertine\" fonts.").ToObject(), πTemp005, πTemp007).ToObject()
					πTemp008 = πg.KWArgs{
						{"template", πTemp001},
						{"latex_preamble", πTemp004},
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßfilter_settings_spec, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp002, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 64: def __init__(self):
					πF.SetLineno(64)
					πTemp009 = make([]πg.Param, 1)
					πTemp009[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 65: latex2e.Writer.__init__(self)
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlatex2e); πE != nil {
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
							// line 66: self.settings_defaults.update({'fontencoding': ''}) # use default (EU1 or EU2)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßfontencoding.ToObject(), ß.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsettings_defaults, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 67: self.translator_class = XeLaTeXTranslator
							πF.SetLineno(67)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßXeLaTeXTranslator); πE != nil {
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
			// line 70: class Babel(latex2e.Babel):
			πF.SetLineno(70)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlatex2e); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßBabel, nil); πE != nil {
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
			_, πE = πg.NewCode("Babel", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
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
					// line 71: """Language specifics for XeTeX.
					πF.SetLineno(71)
					// line 71: """Language specifics for XeTeX.
					πF.SetLineno(71)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Language specifics for XeTeX.\n\n    Use `polyglossia` instead of `babel` and adapt settings.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 75: language_codes = latex2e.Babel.language_codes.copy()
					πF.SetLineno(75)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßlatex2e); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßBabel, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßlanguage_codes, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßlanguage_codes.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 77: language_codes.update({
					πF.SetLineno(77)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßcop.ToObject(), ßcoptic.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßde.ToObject(), ßgerman.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, πg.NewStr("de-1901").ToObject(), ßogerman.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdv.ToObject(), ßdivehi.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdsb.ToObject(), ßlsorbian.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, πg.NewStr("el-polyton").ToObject(), ßpolygreek.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßfa.ToObject(), ßfarsi.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßgrc.ToObject(), ßancientgreek.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßhsb.ToObject(), ßusorbian.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, πg.NewStr("sh-Cyrl").ToObject(), ßserbian.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, πg.NewStr("sh-Latn").ToObject(), ßcroatian.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßsq.ToObject(), ßalbanian.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßsr.ToObject(), ßserbian.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßth.ToObject(), ßthai.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvi.ToObject(), ßvietnamese.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004.ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßlanguage_codes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 97: language_codes = dict([(k.lower(), v) for (k, v) in language_codes.items()])
					πF.SetLineno(97)
					πTemp003 = πF.MakeArgs(1)
					πTemp005 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
						var µv *πg.Object = πg.UnboundLocal
						_ = µv
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
								if πTemp002, πE = πg.ResolveGlobal(πF, ßlanguage_codes); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
										continue
									}
									µk = πTemp003
									µv = πTemp006
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 97: language_codes = dict([(k.lower(), v) for (k, v) in language_codes.items()])
								πF.SetLineno(97)
								if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µk, ßlower, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple2(πTemp006, µv).ToObject()
								πF.PushCheckpoint(4)
								return πTemp002, nil
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
					if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdict); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßlanguage_codes.ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 10)
					πTemp003[0] = ßaf.ToObject()
					πTemp003[1] = πg.NewStr("de-AT").ToObject()
					πTemp003[2] = πg.NewStr("de-AT-1901").ToObject()
					πTemp003[3] = πg.NewStr("en-CA").ToObject()
					πTemp003[4] = πg.NewStr("en-GB").ToObject()
					πTemp003[5] = πg.NewStr("en-NZ").ToObject()
					πTemp003[6] = πg.NewStr("en-US").ToObject()
					πTemp003[7] = πg.NewStr("fr-CA").ToObject()
					πTemp003[8] = πg.NewStr("grc-ibycus").ToObject()
					πTemp003[9] = πg.NewStr("sr-Latn").ToObject()
					πTemp006 = πg.NewTuple(πTemp003...).ToObject()
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
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
					if πTemp006, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πClass.SetItem(πF, ßkey.ToObject(), πTemp006); πE != nil {
							continue
						}
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 112: del(language_codes[key.lower()])
					πF.SetLineno(112)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßlanguage_codes); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßkey); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßlower, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.DelItem(πF, πTemp006, πTemp009); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 114: def __init__(self, language_code, reporter):
					πF.SetLineno(114)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "language_code", Def: nil}
					πTemp005[2] = πg.Param{Name: "reporter", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlanguage_code *πg.Object = πArgs[1]
						_ = µlanguage_code
						var µreporter *πg.Object = πArgs[2]
						_ = µreporter
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 115: self.language_code = language_code
							πF.SetLineno(115)
							if πE = πg.CheckLocal(πF, µlanguage_code, "language_code"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlanguage_code); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage_code, πTemp001); πE != nil {
								continue
							}
							// line 116: self.reporter = reporter
							πF.SetLineno(116)
							if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreporter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreporter, πTemp001); πE != nil {
								continue
							}
							// line 117: self.language = self.language_name(language_code)
							πF.SetLineno(117)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlanguage_code, "language_code"); πE != nil {
								continue
							}
							πTemp002[0] = µlanguage_code
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlanguage_name, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage, πTemp001); πE != nil {
								continue
							}
							// line 118: self.otherlanguages = {}
							πF.SetLineno(118)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßotherlanguages, πTemp003); πE != nil {
								continue
							}
							// line 119: self.warn_msg = 'Language "%s" not supported by Polyglossia.'
							πF.SetLineno(119)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("Language \"%s\" not supported by Polyglossia.").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwarn_msg, πTemp001); πE != nil {
								continue
							}
							// line 120: self.quote_index = 0
							πF.SetLineno(120)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßquote_index, πTemp001); πE != nil {
								continue
							}
							// line 121: self.quotes = ('"', '"')
							πF.SetLineno(121)
							πTemp001 = πg.NewTuple2(πg.NewStr("\"").ToObject(), πg.NewStr("\"").ToObject()).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßquotes, πTemp003); πE != nil {
								continue
							}
							// line 124: self.literal_double_quote = u'"' # TODO: use \textquotedbl ?
							πF.SetLineno(124)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewUnicode("\"").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßliteral_double_quote, πTemp001); πE != nil {
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
					// line 126: def __call__(self):
					πF.SetLineno(126)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__call__", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsetup *πg.Object = πg.UnboundLocal
						_ = µsetup
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
							// line 127: setup = [r'\usepackage{polyglossia}',
							πF.SetLineno(127)
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = πg.NewStr("\\usepackage{polyglossia}").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\\setdefaultlanguage{%s}").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µsetup = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßotherlanguages, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 129: if self.otherlanguages:
							πF.SetLineno(129)
						Label1:
							// line 130: setup.append(r'\setotherlanguages{%s}' %
							πF.SetLineno(130)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßotherlanguages, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp007
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\\setotherlanguages{%s}").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsetup, "setup"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsetup, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 132: return '\n'.join(setup)
							πF.SetLineno(132)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsetup, "setup"); πE != nil {
								continue
							}
							πTemp001[0] = µsetup
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp006); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Babel").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBabel.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 135: class XeLaTeXTranslator(latex2e.LaTeXTranslator):
			πF.SetLineno(135)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlatex2e); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßLaTeXTranslator, nil); πE != nil {
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
			_, πE = πg.NewCode("XeLaTeXTranslator", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 136: """
					πF.SetLineno(136)
					// line 136: """
					πF.SetLineno(136)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Generate code for LaTeX using Unicode fonts (XeLaTex or LuaLaTeX).\n\n    See the docstring of docutils.writers._html_base.HTMLTranslator for\n    notes on and examples of safe subclassing.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 143: def __init__(self, document):
					πF.SetLineno(143)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/xetex/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 144: self.is_xetex = True  # typeset with XeTeX or LuaTeX engine
							πF.SetLineno(144)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßis_xetex, πTemp002); πE != nil {
								continue
							}
							// line 145: latex2e.LaTeXTranslator.__init__(self, document, Babel)
							πF.SetLineno(145)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp003[1] = µdocument
							if πTemp001, πE = πg.ResolveGlobal(πF, ßBabel); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlatex2e); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßLaTeXTranslator, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlatex_encoding, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßutf8.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 146: if self.latex_encoding == 'utf8':
							πF.SetLineno(146)
						Label1:
							// line 147: self.requirements.pop('_inputenc', None)
							πF.SetLineno(147)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ß_inputenc.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrequirements, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 149: self.requirements['_inputenc'] = (r'\XeTeXinputencoding %s '
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlatex_encoding, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\\XeTeXinputencoding %s ").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßrequirements, nil); πE != nil {
								continue
							}
							πTemp006 = ß_inputenc.ToObject()
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("XeLaTeXTranslator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßXeLaTeXTranslator.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("xetex", Code)
}
