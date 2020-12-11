package parts

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßContents := πg.InternStr("Contents")
		ßDirective := πg.InternStr("Directive")
		ßFooter := πg.InternStr("Footer")
		ßHeader := πg.InternStr("Header")
		ßNone := πg.InternStr("None")
		ßSectNum := πg.InternStr("SectNum")
		ßSectnum := πg.InternStr("Sectnum")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßarguments := πg.InternStr("arguments")
		ßassert_has_content := πg.InternStr("assert_has_content")
		ßastext := πg.InternStr("astext")
		ßbacklinks := πg.InternStr("backlinks")
		ßbacklinks_values := πg.InternStr("backlinks_values")
		ßblock_text := πg.InternStr("block_text")
		ßchoice := πg.InternStr("choice")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßcontents := πg.InternStr("contents")
		ßdepth := πg.InternStr("depth")
		ßdetails := πg.InternStr("details")
		ßdirectives := πg.InternStr("directives")
		ßdocument := πg.InternStr("document")
		ßentry := πg.InternStr("entry")
		ßerror := πg.InternStr("error")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßflag := πg.InternStr("flag")
		ßfully_normalize_name := πg.InternStr("fully_normalize_name")
		ßget := πg.InternStr("get")
		ßget_decoration := πg.InternStr("get_decoration")
		ßget_footer := πg.InternStr("get_footer")
		ßget_header := πg.InternStr("get_header")
		ßget_language := πg.InternStr("get_language")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßhas_content := πg.InternStr("has_content")
		ßhas_name := πg.InternStr("has_name")
		ßinline_text := πg.InternStr("inline_text")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßlabels := πg.InternStr("labels")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßline := πg.InternStr("line")
		ßlineno := πg.InternStr("lineno")
		ßlocal := πg.InternStr("local")
		ßmatch_titles := πg.InternStr("match_titles")
		ßname := πg.InternStr("name")
		ßnames := πg.InternStr("names")
		ßnested_parse := πg.InternStr("nested_parse")
		ßnode := πg.InternStr("node")
		ßnodes := πg.InternStr("nodes")
		ßnone := πg.InternStr("none")
		ßnonnegative_int := πg.InternStr("nonnegative_int")
		ßnote_implicit_target := πg.InternStr("note_implicit_target")
		ßnote_pending := πg.InternStr("note_pending")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptional_arguments := πg.InternStr("optional_arguments")
		ßoptions := πg.InternStr("options")
		ßparts := πg.InternStr("parts")
		ßpending := πg.InternStr("pending")
		ßprefix := πg.InternStr("prefix")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrun := πg.InternStr("run")
		ßsettings := πg.InternStr("settings")
		ßsidebar := πg.InternStr("sidebar")
		ßsource := πg.InternStr("source")
		ßstart := πg.InternStr("start")
		ßstate := πg.InternStr("state")
		ßstate_machine := πg.InternStr("state_machine")
		ßsuffix := πg.InternStr("suffix")
		ßtitle := πg.InternStr("title")
		ßtop := πg.InternStr("top")
		ßtopic := πg.InternStr("topic")
		ßunchanged_required := πg.InternStr("unchanged_required")
		ßupdate := πg.InternStr("update")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDirectives for document parts.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 11: from docutils import nodes, languages
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: from docutils.transforms import parts
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.parts"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßparts.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from docutils.parsers.rst import Directive
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDirective); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDirective.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: from docutils.parsers.rst import directives
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: class Contents(Directive):
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
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
			_, πE = πg.NewCode("Contents", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
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
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Table of contents.\n\n    The table of contents is generated in two passes: initial parse and\n    transform.  During the initial parse, a 'pending' element is generated\n    which acts as a placeholder, storing the TOC title and any options\n    internally.  At a later stage in the processing, the 'pending' element is\n    replaced by a 'topic' element, a title and the table of contents proper.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 29: backlinks_values = ('top', 'entry', 'none')
					πF.SetLineno(29)
					πTemp001 = πg.NewTuple3(ßtop.ToObject(), ßentry.ToObject(), ßnone.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßbacklinks_values.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 31: def backlinks(arg):
					πF.SetLineno(31)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "arg", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("backlinks", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µarg *πg.Object = πArgs[0]
						_ = µarg
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							// line 32: value = directives.choice(arg, Contents.backlinks_values)
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp001[0] = µarg
							if πTemp002, πE = πg.ResolveGlobal(πF, ßContents); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßbacklinks_values, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchoice, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp002
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µvalue, ßnone.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 33: if value == 'none':
							πF.SetLineno(33)
						Label1:
							// line 34: return None
							πF.SetLineno(34)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 36: return value
							πF.SetLineno(36)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πR = µvalue
							continue
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
					if πE = πClass.SetItem(πF, ßbacklinks.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 38: optional_arguments = 1
					πF.SetLineno(38)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 39: final_argument_whitespace = True
					πF.SetLineno(39)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 40: option_spec = {'depth': directives.nonnegative_int,
					πF.SetLineno(40)
					πTemp004 = πg.NewDict()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßnonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdepth.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßflag, nil); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßlocal.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßbacklinks); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßbacklinks.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßclass.ToObject(), πTemp005); πE != nil {
						continue
					}
					πTemp003 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 45: def run(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πg.UnboundLocal
						_ = µdocument
						var µlanguage *πg.Object = πg.UnboundLocal
						_ = µlanguage
						var µtitle_text *πg.Object = πg.UnboundLocal
						_ = µtitle_text
						var µtext_nodes *πg.Object = πg.UnboundLocal
						_ = µtext_nodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µtopic *πg.Object = πg.UnboundLocal
						_ = µtopic
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmatch_titles, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßnode, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsidebar, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πTemp005
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 46: if not (self.state_machine.match_titles
							πF.SetLineno(46)
						Label2:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("The \"%s\" directive may not be used within topics or body elements.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 48: raise self.error('The "%s" directive may not be used within '
							πF.SetLineno(48)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label3
						Label3:
							// line 50: document = self.state_machine.document
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							µdocument = πTemp002
							// line 51: language = languages.get_language(document.settings.language_code,
							πF.SetLineno(51)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlanguage_code, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlanguages); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µlanguage = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 53: if self.arguments:
							πF.SetLineno(53)
						Label4:
							// line 54: title_text = self.arguments[0]
							πF.SetLineno(54)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µtitle_text = πTemp002
							// line 55: text_nodes, messages = self.state.inline_text(title_text,
							πF.SetLineno(55)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp006[0] = µtitle_text
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinline_text, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µtext_nodes = πTemp002
							µmessages = πTemp004
							// line 57: title = nodes.title(title_text, '', *text_nodes)
							πF.SetLineno(57)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp006[0] = µtitle_text
							πTemp006[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtext_nodes, "text_nodes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp006, µtext_nodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µtitle = πTemp001
							goto Label6
						Label5:
							// line 59: messages = []
							πF.SetLineno(59)
							πTemp006 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µmessages = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßlocal.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 60: if 'local' in self.options:
							πF.SetLineno(60)
						Label7:
							// line 61: title = None
							πF.SetLineno(61)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtitle = πTemp001
							goto Label9
						Label8:
							// line 63: title = nodes.title('', language.labels['contents'])
							πF.SetLineno(63)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ß.ToObject()
							πTemp001 = ßcontents.ToObject()
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µlanguage, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µtitle = πTemp001
							goto Label9
						Label9:
							goto Label6
						Label6:
							// line 64: topic = nodes.topic(classes=['contents'])
							πF.SetLineno(64)
							πTemp006 = make([]*πg.Object, 1)
							πTemp006[0] = ßcontents.ToObject()
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp007 = πg.KWArgs{
								{"classes", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtopic, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							µtopic = πTemp001
							// line 65: topic['classes'] += self.options.get('class', [])
							πF.SetLineno(65)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtopic, πTemp001); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßclass.ToObject()
							πTemp008 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µtopic, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 67: topic.source, topic.line = self.state_machine.get_source_and_line()
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_source_and_line, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtopic, ßsource, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtopic, ßline, πTemp004); πE != nil {
								continue
							}
							// line 68: topic.line -= 1
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtopic, ßline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtopic, ßline, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßlocal.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							goto Label11
							// line 69: if 'local' in self.options:
							πF.SetLineno(69)
						Label10:
							// line 70: topic['classes'].append('local')
							πF.SetLineno(70)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßlocal.ToObject()
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtopic, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µtitle); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 71: if title:
							πF.SetLineno(71)
						Label12:
							// line 72: name = title.astext()
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtitle, ßastext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µname = πTemp002
							// line 73: topic += title
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µtopic, µtitle); πE != nil {
								continue
							}
							µtopic = πTemp001
							goto Label14
						Label13:
							// line 75: name = language.labels['contents']
							πF.SetLineno(75)
							πTemp001 = ßcontents.ToObject()
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µlanguage, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µname = πTemp002
							goto Label14
						Label14:
							// line 76: name = nodes.fully_normalize_name(name)
							πF.SetLineno(76)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfully_normalize_name, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µname = πTemp001
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocument, ßhas_name, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							goto Label16
							// line 77: if not document.has_name(name):
							πF.SetLineno(77)
						Label15:
							// line 78: topic['names'].append(name)
							πF.SetLineno(78)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							πTemp001 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtopic, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label16
						Label16:
							// line 79: document.note_implicit_target(topic)
							πF.SetLineno(79)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							πTemp006[0] = µtopic
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßnote_implicit_target, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 80: pending = nodes.pending(parts.Contents, rawsource=self.block_text)
							πF.SetLineno(80)
							πTemp006 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßparts); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßContents, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"rawsource", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µpending = πTemp001
							// line 81: pending.details.update(self.options)
							πF.SetLineno(81)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 82: document.note_pending(pending)
							πF.SetLineno(82)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp006[0] = µpending
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 83: topic += pending
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µtopic, µpending); πE != nil {
								continue
							}
							µtopic = πTemp001
							// line 84: return [topic] + messages
							πF.SetLineno(84)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							πTemp006[0] = µtopic
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µmessages); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Contents").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßContents.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 87: class Sectnum(Directive):
			πF.SetLineno(87)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
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
			_, πE = πg.NewCode("Sectnum", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 89: """Automatic section numbering."""
					πF.SetLineno(89)
					// line 89: """Automatic section numbering."""
					πF.SetLineno(89)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Automatic section numbering.").ToObject()); πE != nil {
						continue
					}
					// line 91: option_spec = {'depth': int,
					πF.SetLineno(91)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdepth.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßstart.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßprefix.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsuffix.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 96: def run(self):
					πF.SetLineno(96)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
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
							// line 97: pending = nodes.pending(parts.SectNum)
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßparts); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSectNum, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpending, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µpending = πTemp002
							// line 98: pending.details.update(self.options)
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 99: self.state_machine.document.note_pending(pending)
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 100: return [pending]
							πF.SetLineno(100)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							πTemp002 = πg.NewList(πTemp001...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Sectnum").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSectnum.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 103: class Header(Directive):
			πF.SetLineno(103)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
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
			_, πE = πg.NewCode("Header", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 105: """Contents of document header."""
					πF.SetLineno(105)
					// line 105: """Contents of document header."""
					πF.SetLineno(105)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Contents of document header.").ToObject()); πE != nil {
						continue
					}
					// line 107: has_content = True
					πF.SetLineno(107)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 109: def run(self):
					πF.SetLineno(109)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
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
							// line 110: self.assert_has_content()
							πF.SetLineno(110)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 111: header = self.state_machine.document.get_decoration().get_header()
							πF.SetLineno(111)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßget_decoration, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßget_header, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µheader = πTemp002
							// line 112: self.state.nested_parse(self.content, self.content_offset, header)
							πF.SetLineno(112)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp003[2] = µheader
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 113: return []
							πF.SetLineno(113)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Header").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHeader.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 116: class Footer(Directive):
			πF.SetLineno(116)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
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
			_, πE = πg.NewCode("Footer", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 118: """Contents of document footer."""
					πF.SetLineno(118)
					// line 118: """Contents of document footer."""
					πF.SetLineno(118)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Contents of document footer.").ToObject()); πE != nil {
						continue
					}
					// line 120: has_content = True
					πF.SetLineno(120)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 122: def run(self):
					πF.SetLineno(122)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfooter *πg.Object = πg.UnboundLocal
						_ = µfooter
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
							// line 123: self.assert_has_content()
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 124: footer = self.state_machine.document.get_decoration().get_footer()
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßget_decoration, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßget_footer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfooter = πTemp002
							// line 125: self.state.nested_parse(self.content, self.content_offset, footer)
							πF.SetLineno(125)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µfooter, "footer"); πE != nil {
								continue
							}
							πTemp003[2] = µfooter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 126: return []
							πF.SetLineno(126)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Footer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFooter.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("parts", Code)
}
