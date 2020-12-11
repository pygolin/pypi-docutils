package body

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßBasePseudoSection := πg.InternStr("BasePseudoSection")
		ßBlockQuote := πg.InternStr("BlockQuote")
		ßCodeBlock := πg.InternStr("CodeBlock")
		ßCompound := πg.InternStr("Compound")
		ßContainer := πg.InternStr("Container")
		ßDirective := πg.InternStr("Directive")
		ßEpigraph := πg.InternStr("Epigraph")
		ßHighlights := πg.InternStr("Highlights")
		ßLexer := πg.InternStr("Lexer")
		ßLexerError := πg.InternStr("LexerError")
		ßLineBlock := πg.InternStr("LineBlock")
		ßMathBlock := πg.InternStr("MathBlock")
		ßNone := πg.InternStr("None")
		ßNumberLines := πg.InternStr("NumberLines")
		ßParsedLiteral := πg.InternStr("ParsedLiteral")
		ßPullQuote := πg.InternStr("PullQuote")
		ßRubric := πg.InternStr("Rubric")
		ßSidebar := πg.InternStr("Sidebar")
		ßText := πg.InternStr("Text")
		ßTopic := πg.InternStr("Topic")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd_name := πg.InternStr("add_name")
		ßappend := πg.InternStr("append")
		ßarguments := πg.InternStr("arguments")
		ßassert_has_content := πg.InternStr("assert_has_content")
		ßattributes := πg.InternStr("attributes")
		ßblock_quote := πg.InternStr("block_quote")
		ßblock_text := πg.InternStr("block_text")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcode := πg.InternStr("code")
		ßcompound := πg.InternStr("compound")
		ßcontainer := πg.InternStr("container")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßcopy := πg.InternStr("copy")
		ßdirectives := πg.InternStr("directives")
		ßdocument := πg.InternStr("document")
		ßepigraph := πg.InternStr("epigraph")
		ßerror := πg.InternStr("error")
		ßextend := πg.InternStr("extend")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßget := πg.InternStr("get")
		ßhas_content := πg.InternStr("has_content")
		ßhighlights := πg.InternStr("highlights")
		ßindent := πg.InternStr("indent")
		ßinline := πg.InternStr("inline")
		ßinline_text := πg.InternStr("inline_text")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßline_block := πg.InternStr("line_block")
		ßlineno := πg.InternStr("lineno")
		ßliteral_block := πg.InternStr("literal_block")
		ßlstrip := πg.InternStr("lstrip")
		ßmatch_titles := πg.InternStr("match_titles")
		ßmath_block := πg.InternStr("math_block")
		ßname := πg.InternStr("name")
		ßnest_line_block_lines := πg.InternStr("nest_line_block_lines")
		ßnested_parse := πg.InternStr("nested_parse")
		ßnode := πg.InternStr("node")
		ßnode_class := πg.InternStr("node_class")
		ßnodes := πg.InternStr("nodes")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptional_arguments := πg.InternStr("optional_arguments")
		ßoptions := πg.InternStr("options")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrequired_arguments := πg.InternStr("required_arguments")
		ßrubric := πg.InternStr("rubric")
		ßrun := πg.InternStr("run")
		ßset_classes := πg.InternStr("set_classes")
		ßsettings := πg.InternStr("settings")
		ßsidebar := πg.InternStr("sidebar")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßstate := πg.InternStr("state")
		ßstate_machine := πg.InternStr("state_machine")
		ßstrip := πg.InternStr("strip")
		ßsubtitle := πg.InternStr("subtitle")
		ßsyntax_highlight := πg.InternStr("syntax_highlight")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßunchanged := πg.InternStr("unchanged")
		ßunchanged_required := πg.InternStr("unchanged_required")
		ßwarning := πg.InternStr("warning")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDirectives for additional body elements.\n\nSee `docutils.parsers.rst.directives` for API details.\n").ToObject()); πE != nil {
				continue
			}
			// line 11: __docformat__ = 'reStructuredText'
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 14: from docutils import nodes
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: from docutils.parsers.rst import Directive
			πF.SetLineno(15)
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
			// line 16: from docutils.parsers.rst import directives
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: from docutils.parsers.rst.roles import set_classes
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.roles"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßset_classes); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßset_classes.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from docutils.utils.code_analyzer import Lexer, LexerError, NumberLines
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.code_analyzer"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLexer); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLexer.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLexerError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLexerError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßNumberLines); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumberLines.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: class BasePseudoSection(Directive):
			πF.SetLineno(21)
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
			_, πE = πg.NewCode("BasePseudoSection", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
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
					// line 23: required_arguments = 1
					πF.SetLineno(23)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 24: optional_arguments = 0
					πF.SetLineno(24)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 25: final_argument_whitespace = True
					πF.SetLineno(25)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 26: option_spec = {'class': directives.class_option,
					πF.SetLineno(26)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 28: has_content = True
					πF.SetLineno(28)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 30: node_class = None
					πF.SetLineno(30)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnode_class.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 31: """Node class to be used (must be set in subclasses)."""
					πF.SetLineno(31)
					// line 31: """Node class to be used (must be set in subclasses)."""
					πF.SetLineno(31)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Node class to be used (must be set in subclasses).").ToObject()); πE != nil {
						continue
					}
					// line 33: def run(self):
					πF.SetLineno(33)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtitle_text *πg.Object = πg.UnboundLocal
						_ = µtitle_text
						var µtextnodes *πg.Object = πg.UnboundLocal
						_ = µtextnodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µtitles *πg.Object = πg.UnboundLocal
						_ = µtitles
						var µmore_messages *πg.Object = πg.UnboundLocal
						_ = µmore_messages
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
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
							// line 34: if not (self.state_machine.match_titles
							πF.SetLineno(34)
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
							// line 36: raise self.error('The "%s" directive may not be used within '
							πF.SetLineno(36)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label3
						Label3:
							// line 38: self.assert_has_content()
							πF.SetLineno(38)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 39: title_text = self.arguments[0]
							πF.SetLineno(39)
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
							// line 40: textnodes, messages = self.state.inline_text(title_text, self.lineno)
							πF.SetLineno(40)
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
							µtextnodes = πTemp002
							µmessages = πTemp004
							// line 41: titles = [nodes.title(title_text, '', *textnodes)]
							πF.SetLineno(41)
							πTemp006 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp007[0] = µtitle_text
							πTemp007[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtextnodes, "textnodes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp007, µtextnodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp001
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µtitles = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßsubtitle.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 43: if 'subtitle' in self.options:
							πF.SetLineno(43)
						Label4:
							// line 44: textnodes, more_messages = self.state.inline_text(
							πF.SetLineno(44)
							πTemp006 = πF.MakeArgs(2)
							πTemp001 = ßsubtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
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
							µtextnodes = πTemp002
							µmore_messages = πTemp004
							// line 46: titles.append(nodes.subtitle(self.options['subtitle'], '',
							πF.SetLineno(46)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							πTemp001 = ßsubtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							πTemp007[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtextnodes, "textnodes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsubtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp007, µtextnodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µtitles, "titles"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtitles, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 48: messages.extend(more_messages)
							πF.SetLineno(48)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmore_messages, "more_messages"); πE != nil {
								continue
							}
							πTemp006[0] = µmore_messages
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmessages, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label5
						Label5:
							// line 49: text = '\n'.join(self.content)
							πF.SetLineno(49)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µtext = πTemp002
							// line 50: node = self.node_class(text, *(titles + messages))
							πF.SetLineno(50)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp006[0] = µtext
							if πE = πg.CheckLocal(πF, µtitles, "titles"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µtitles, µmessages); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnode_class, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp002, πTemp006, πTemp001, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µnode = πTemp004
							// line 51: node['classes'] += self.options.get('class', [])
							πF.SetLineno(51)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = ßclass.ToObject()
							πTemp007 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp007...).ToObject()
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 52: self.add_name(node)
							πF.SetLineno(52)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µtext); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 53: if text:
							πF.SetLineno(53)
						Label6:
							// line 54: self.state.nested_parse(self.content, self.content_offset, node)
							πF.SetLineno(54)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[2] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label7
						Label7:
							// line 55: return [node]
							πF.SetLineno(55)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006[0] = µnode
							πTemp001 = πg.NewList(πTemp006...).ToObject()
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BasePseudoSection").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBasePseudoSection.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 58: class Topic(BasePseudoSection):
			πF.SetLineno(58)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBasePseudoSection); πE != nil {
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
			_, πE = πg.NewCode("Topic", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 60: node_class = nodes.topic
					πF.SetLineno(60)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtopic, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnode_class.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Topic").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTopic.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 63: class Sidebar(BasePseudoSection):
			πF.SetLineno(63)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBasePseudoSection); πE != nil {
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
			_, πE = πg.NewCode("Sidebar", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 65: node_class = nodes.sidebar
					πF.SetLineno(65)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsidebar, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnode_class.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 67: option_spec = BasePseudoSection.option_spec.copy()
					πF.SetLineno(67)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßBasePseudoSection); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßoption_spec, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 68: option_spec['subtitle'] = directives.unchanged_required
					πF.SetLineno(68)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßoption_spec); πE != nil {
						continue
					}
					πTemp004 = ßsubtitle.ToObject()
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
						continue
					}
					// line 70: def run(self):
					πF.SetLineno(70)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnode, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsidebar, nil); πE != nil {
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
							// line 71: if isinstance(self.state_machine.node, nodes.sidebar):
							πF.SetLineno(71)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("The \"%s\" directive may not be used within a sidebar element.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 72: raise self.error('The "%s" directive may not be used within a '
							πF.SetLineno(72)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 74: return BasePseudoSection.run(self)
							πF.SetLineno(74)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasePseudoSection); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Sidebar").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSidebar.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 77: class LineBlock(Directive):
			πF.SetLineno(77)
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
			_, πE = πg.NewCode("LineBlock", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 79: option_spec = {'class': directives.class_option,
					πF.SetLineno(79)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 81: has_content = True
					πF.SetLineno(81)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 83: def run(self):
					πF.SetLineno(83)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var µnode_list *πg.Object = πg.UnboundLocal
						_ = µnode_list
						var µline_text *πg.Object = πg.UnboundLocal
						_ = µline_text
						var µtext_nodes *πg.Object = πg.UnboundLocal
						_ = µtext_nodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
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
							// line 84: self.assert_has_content()
							πF.SetLineno(84)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 85: block = nodes.line_block(classes=self.options.get('class', []))
							πF.SetLineno(85)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßclass.ToObject()
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp005 = πg.KWArgs{
								{"classes", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßline_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							µblock = πTemp001
							// line 86: self.add_name(block)
							πF.SetLineno(86)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp003[0] = µblock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 87: node_list = [block]
							πF.SetLineno(87)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp003[0] = µblock
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µnode_list = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µline_text = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 89: text_nodes, messages = self.state.inline_text(
							πF.SetLineno(89)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline_text, "line_text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µline_text, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßinline_text, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
								continue
							}
							µtext_nodes = πTemp008
							µmessages = πTemp009
							// line 91: line = nodes.line(line_text, '', *text_nodes)
							πF.SetLineno(91)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline_text, "line_text"); πE != nil {
								continue
							}
							πTemp003[0] = µline_text
							πTemp003[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtext_nodes, "text_nodes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp008, πTemp003, µtext_nodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µline = πTemp002
							if πE = πg.CheckLocal(πF, µline_text, "line_text"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µline_text, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 92: if line_text.strip():
							πF.SetLineno(92)
						Label4:
							// line 93: line.indent = len(line_text) - len(line_text.lstrip())
							πF.SetLineno(93)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline_text, "line_text"); πE != nil {
								continue
							}
							πTemp003[0] = µline_text
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline_text, "line_text"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µline_text, ßlstrip, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Sub(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µline, ßindent, πTemp008); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 94: block += line
							πF.SetLineno(94)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µblock, µline); πE != nil {
								continue
							}
							µblock = πTemp002
							// line 95: node_list.extend(messages)
							πF.SetLineno(95)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp003[0] = µmessages
							if πE = πg.CheckLocal(πF, µnode_list, "node_list"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode_list, ßextend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 96: self.content_offset += 1
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontent_offset, πTemp008); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 97: self.state.nest_line_block_lines(block)
							πF.SetLineno(97)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp003[0] = µblock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnest_line_block_lines, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 98: return node_list
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µnode_list, "node_list"); πE != nil {
								continue
							}
							πR = µnode_list
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("LineBlock").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLineBlock.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 101: class ParsedLiteral(Directive):
			πF.SetLineno(101)
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
			_, πE = πg.NewCode("ParsedLiteral", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 103: option_spec = {'class': directives.class_option,
					πF.SetLineno(103)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 105: has_content = True
					πF.SetLineno(105)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 107: def run(self):
					πF.SetLineno(107)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µtext_nodes *πg.Object = πg.UnboundLocal
						_ = µtext_nodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var πTemp001 []*πg.Object
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
							// line 108: set_classes(self.options)
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 109: self.assert_has_content()
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 110: text = '\n'.join(self.content)
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp003
							// line 111: text_nodes, messages = self.state.inline_text(text, self.lineno)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinline_text, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µtext_nodes = πTemp003
							µmessages = πTemp004
							// line 112: node = nodes.literal_block(text, '', *text_nodes, **self.options)
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtext_nodes, "text_nodes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, πTemp001, µtext_nodes, nil, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnode = πTemp003
							// line 113: node.line = self.content_offset + 1
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnode, ßline, πTemp003); πE != nil {
								continue
							}
							// line 114: self.add_name(node)
							πF.SetLineno(114)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 115: return [node] + messages
							πF.SetLineno(115)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µmessages); πE != nil {
								continue
							}
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ParsedLiteral").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßParsedLiteral.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 118: class CodeBlock(Directive):
			πF.SetLineno(118)
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
			_, πE = πg.NewCode("CodeBlock", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 119: """Parse and mark up content of a code block.
					πF.SetLineno(119)
					// line 119: """Parse and mark up content of a code block.
					πF.SetLineno(119)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Parse and mark up content of a code block.\n\n    Configuration setting: syntax_highlight\n       Highlight Code content with Pygments?\n       Possible values: ('long', 'short', 'none')\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 126: optional_arguments = 1
					πF.SetLineno(126)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 127: option_spec = {'class': directives.class_option,
					πF.SetLineno(127)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("number-lines").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 131: has_content = True
					πF.SetLineno(131)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 133: def run(self):
					πF.SetLineno(133)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlanguage *πg.Object = πg.UnboundLocal
						_ = µlanguage
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µtokens *πg.Object = πg.UnboundLocal
						_ = µtokens
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µstartline *πg.Object = πg.UnboundLocal
						_ = µstartline
						var µendline *πg.Object = πg.UnboundLocal
						_ = µendline
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9:
								goto Label9
							case 19:
								goto Label19
							case 20:
								goto Label20
							case 14:
								goto Label14
							default:
								panic("unexpected function state")
							}
							// line 134: self.assert_has_content()
							πF.SetLineno(134)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
								goto Label1
							}
							goto Label2
							// line 135: if self.arguments:
							πF.SetLineno(135)
						Label1:
							// line 136: language = self.arguments[0]
							πF.SetLineno(136)
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
							µlanguage = πTemp002
							goto Label3
						Label2:
							// line 138: language = ''
							πF.SetLineno(138)
							µlanguage = ß.ToObject()
							goto Label3
						Label3:
							// line 139: set_classes(self.options)
							πF.SetLineno(139)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 140: classes = ['code']
							πF.SetLineno(140)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = ßcode.ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µclasses = πTemp001
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µlanguage); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 141: if language:
							πF.SetLineno(141)
						Label4:
							// line 142: classes.append(language)
							πF.SetLineno(142)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							πTemp005[0] = µlanguage
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µclasses, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßclasses.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 143: if 'classes' in self.options:
							πF.SetLineno(143)
						Label6:
							// line 144: classes.extend(self.options['classes'])
							πF.SetLineno(144)
							πTemp005 = πF.MakeArgs(1)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µclasses, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label7:
							// line 147: try:
							πF.SetLineno(147)
							πF.PushCheckpoint(9)
							// line 148: tokens = Lexer(u'\n'.join(self.content), language,
							πF.SetLineno(148)
							πTemp005 = πF.MakeArgs(3)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewUnicode("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							πTemp005[1] = µlanguage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsyntax_highlight, nil); πE != nil {
								continue
							}
							πTemp005[2] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßLexer); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtokens = πTemp002
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßLexerError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 150: except LexerError as error:
							πF.SetLineno(150)
						Label10:
							µerror = πTemp007.ToObject()
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp005[0] = µerror
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 151: raise self.warning(error)
							πF.SetLineno(151)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, πg.NewStr("number-lines").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 153: if 'number-lines' in self.options:
							πF.SetLineno(153)
						Label11:
							// line 155: try:
							πF.SetLineno(155)
							πF.PushCheckpoint(14)
							// line 156: startline = int(self.options['number-lines'] or 1)
							πF.SetLineno(156)
							πTemp005 = πF.MakeArgs(1)
							πTemp002 = πg.NewStr("number-lines").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							πTemp001 = πg.NewInt(1).ToObject()
						Label15:
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µstartline = πTemp002
							πF.PopCheckpoint()
							goto Label13
						Label14:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 157: except ValueError:
							πF.SetLineno(157)
						Label16:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr(":number-lines: with non-integer start value").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 158: raise self.error(':number-lines: with non-integer start value')
							πF.SetLineno(158)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label13
						Label13:
							// line 159: endline = startline + len(self.content)
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, µstartline, πTemp004); πE != nil {
								continue
							}
							µendline = πTemp001
							// line 161: tokens = NumberLines(tokens, startline, endline)
							πF.SetLineno(161)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							πTemp005[0] = µtokens
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							πTemp005[1] = µstartline
							if πE = πg.CheckLocal(πF, µendline, "endline"); πE != nil {
								continue
							}
							πTemp005[2] = µendline
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNumberLines); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtokens = πTemp002
							goto Label12
						Label12:
							// line 163: node = nodes.literal_block('\n'.join(self.content), classes=classes)
							πF.SetLineno(163)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"classes", µclasses},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µnode = πTemp001
							// line 164: self.add_name(node)
							πF.SetLineno(164)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßsource.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label17
							}
							goto Label18
							// line 166: if 'source' in self.options:
							πF.SetLineno(166)
						Label17:
							// line 167: node.attributes['source'] = self.options['source']
							πF.SetLineno(167)
							πTemp001 = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							πTemp009 = ßsource.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp009, πTemp001); πE != nil {
								continue
							}
							goto Label18
						Label18:
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtokens); πE != nil {
								continue
							}
							πF.PushCheckpoint(20)
							πTemp003 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label21
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
									continue
								}
								µclasses = πTemp004
								µvalue = πTemp009
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(19)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, µclasses); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label22
							}
							goto Label23
							// line 170: if classes:
							πF.SetLineno(170)
						Label22:
							// line 171: node += nodes.inline(value, value, classes=classes)
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[1] = µvalue
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"classes", µclasses},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßinline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp005, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.IAdd(πF, µnode, πTemp002); πE != nil {
								continue
							}
							µnode = πTemp004
							goto Label24
						Label23:
							// line 174: node += nodes.Text(value)
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßText, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.IAdd(πF, µnode, πTemp002); πE != nil {
								continue
							}
							µnode = πTemp004
							goto Label24
						Label24:
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
							// line 176: return [node]
							πF.SetLineno(176)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							πTemp001 = πg.NewList(πTemp005...).ToObject()
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("CodeBlock").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCodeBlock.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 179: class MathBlock(Directive):
			πF.SetLineno(179)
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
			_, πE = πg.NewCode("MathBlock", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 181: option_spec = {'class': directives.class_option,
					πF.SetLineno(181)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 185: has_content = True
					πF.SetLineno(185)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 187: def run(self):
					πF.SetLineno(187)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcontent *πg.Object = πg.UnboundLocal
						_ = µcontent
						var µ_nodes *πg.Object = πg.UnboundLocal
						_ = µ_nodes
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
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
							// line 188: set_classes(self.options)
							πF.SetLineno(188)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 189: self.assert_has_content()
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 191: content = '\n'.join(self.content).split('\n\n')
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n\n").ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcontent = πTemp003
							// line 192: _nodes = []
							πF.SetLineno(192)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µ_nodes = πTemp002
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µcontent); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µblock = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µblock); πE != nil {
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
							// line 194: if not block:
							πF.SetLineno(194)
						Label4:
							// line 195: continue
							πF.SetLineno(195)
							continue
							goto Label5
						Label5:
							// line 196: node = nodes.math_block(self.block_text, block, **self.options)
							πF.SetLineno(196)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp001[1] = µblock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßmath_block, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp008, πTemp001, nil, nil, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnode = πTemp007
							// line 197: node.line = self.content_offset + 1
							πF.SetLineno(197)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnode, ßline, πTemp007); πE != nil {
								continue
							}
							// line 198: self.add_name(node)
							πF.SetLineno(198)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 199: _nodes.append(node)
							πF.SetLineno(199)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µ_nodes, "_nodes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µ_nodes, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 200: return _nodes
							πF.SetLineno(200)
							if πE = πg.CheckLocal(πF, µ_nodes, "_nodes"); πE != nil {
								continue
							}
							πR = µ_nodes
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MathBlock").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMathBlock.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 203: class Rubric(Directive):
			πF.SetLineno(203)
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
			_, πE = πg.NewCode("Rubric", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
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
					// line 205: required_arguments = 1
					πF.SetLineno(205)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 206: optional_arguments = 0
					πF.SetLineno(206)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 207: final_argument_whitespace = True
					πF.SetLineno(207)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 208: option_spec = {'class': directives.class_option,
					πF.SetLineno(208)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 211: def run(self):
					πF.SetLineno(211)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrubric_text *πg.Object = πg.UnboundLocal
						_ = µrubric_text
						var µtextnodes *πg.Object = πg.UnboundLocal
						_ = µtextnodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µrubric *πg.Object = πg.UnboundLocal
						_ = µrubric
						var πTemp001 []*πg.Object
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
							// line 212: set_classes(self.options)
							πF.SetLineno(212)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 213: rubric_text = self.arguments[0]
							πF.SetLineno(213)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µrubric_text = πTemp003
							// line 214: textnodes, messages = self.state.inline_text(rubric_text, self.lineno)
							πF.SetLineno(214)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrubric_text, "rubric_text"); πE != nil {
								continue
							}
							πTemp001[0] = µrubric_text
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinline_text, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µtextnodes = πTemp003
							µmessages = πTemp004
							// line 215: rubric = nodes.rubric(rubric_text, '', *textnodes, **self.options)
							πF.SetLineno(215)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrubric_text, "rubric_text"); πE != nil {
								continue
							}
							πTemp001[0] = µrubric_text
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtextnodes, "textnodes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrubric, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, πTemp001, µtextnodes, nil, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrubric = πTemp003
							// line 216: self.add_name(rubric)
							πF.SetLineno(216)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrubric, "rubric"); πE != nil {
								continue
							}
							πTemp001[0] = µrubric
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 217: return [rubric] + messages
							πF.SetLineno(217)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µrubric, "rubric"); πE != nil {
								continue
							}
							πTemp001[0] = µrubric
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µmessages); πE != nil {
								continue
							}
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Rubric").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRubric.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 220: class BlockQuote(Directive):
			πF.SetLineno(220)
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
			_, πE = πg.NewCode("BlockQuote", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 222: has_content = True
					πF.SetLineno(222)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 223: classes = []
					πF.SetLineno(223)
					πTemp002 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßclasses.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 225: def run(self):
					πF.SetLineno(225)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µelements *πg.Object = πg.UnboundLocal
						_ = µelements
						var µelement *πg.Object = πg.UnboundLocal
						_ = µelement
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
							// line 226: self.assert_has_content()
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 227: elements = self.state.block_quote(self.content, self.content_offset)
							πF.SetLineno(227)
							πTemp003 = πF.MakeArgs(2)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßblock_quote, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µelements = πTemp001
							if πE = πg.CheckLocal(πF, µelements, "elements"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µelements); πE != nil {
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
								µelement = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp003[0] = µelement
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßblock_quote, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 229: if isinstance(element, nodes.block_quote):
							πF.SetLineno(229)
						Label4:
							// line 230: element['classes'] += self.classes
							πF.SetLineno(230)
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µelement, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclasses, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IAdd(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp008 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µelement, πTemp008, πTemp007); πE != nil {
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
							// line 231: return elements
							πF.SetLineno(231)
							if πE = πg.CheckLocal(πF, µelements, "elements"); πE != nil {
								continue
							}
							πR = µelements
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BlockQuote").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBlockQuote.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 234: class Epigraph(BlockQuote):
			πF.SetLineno(234)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBlockQuote); πE != nil {
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
			_, πE = πg.NewCode("Epigraph", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 236: classes = ['epigraph']
					πF.SetLineno(236)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßepigraph.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßclasses.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Epigraph").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEpigraph.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 239: class Highlights(BlockQuote):
			πF.SetLineno(239)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBlockQuote); πE != nil {
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
			_, πE = πg.NewCode("Highlights", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 241: classes = ['highlights']
					πF.SetLineno(241)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßhighlights.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßclasses.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Highlights").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHighlights.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 244: class PullQuote(BlockQuote):
			πF.SetLineno(244)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBlockQuote); πE != nil {
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
			_, πE = πg.NewCode("PullQuote", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 246: classes = ['pull-quote']
					πF.SetLineno(246)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("pull-quote").ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßclasses.ToObject(), πTemp002); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PullQuote").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPullQuote.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 249: class Compound(Directive):
			πF.SetLineno(249)
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
			_, πE = πg.NewCode("Compound", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 251: option_spec = {'class': directives.class_option,
					πF.SetLineno(251)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 253: has_content = True
					πF.SetLineno(253)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 255: def run(self):
					πF.SetLineno(255)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 256: self.assert_has_content()
							πF.SetLineno(256)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 257: text = '\n'.join(self.content)
							πF.SetLineno(257)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp002
							// line 258: node = nodes.compound(text)
							πF.SetLineno(258)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcompound, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µnode = πTemp001
							// line 259: node['classes'] += self.options.get('class', [])
							πF.SetLineno(259)
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßclass.ToObject()
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp006 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp006, πTemp005); πE != nil {
								continue
							}
							// line 260: self.add_name(node)
							πF.SetLineno(260)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 261: self.state.nested_parse(self.content, self.content_offset, node)
							πF.SetLineno(261)
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[2] = µnode
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
							// line 262: return [node]
							πF.SetLineno(262)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Compound").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCompound.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 265: class Container(Directive):
			πF.SetLineno(265)
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
			_, πE = πg.NewCode("Container", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
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
					// line 267: optional_arguments = 1
					πF.SetLineno(267)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 268: final_argument_whitespace = True
					πF.SetLineno(268)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 269: option_spec = {'name': directives.unchanged}
					πF.SetLineno(269)
					πTemp002 = πg.NewDict()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp002.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 270: has_content = True
					πF.SetLineno(270)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 272: def run(self):
					πF.SetLineno(272)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/body.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 273: self.assert_has_content()
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 274: text = '\n'.join(self.content)
							πF.SetLineno(274)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtext = πTemp002
							// line 275: try:
							πF.SetLineno(275)
							πF.PushCheckpoint(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 276: if self.arguments:
							πF.SetLineno(276)
						Label3:
							// line 277: classes = directives.class_option(self.arguments[0])
							πF.SetLineno(277)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µclasses = πTemp001
							goto Label5
						Label4:
							// line 279: classes = []
							πF.SetLineno(279)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µclasses = πTemp001
							goto Label5
						Label5:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 280: except ValueError:
							πF.SetLineno(280)
						Label6:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid class attribute value for \"%s\" directive: \"%s\".").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 281: raise self.error(
							πF.SetLineno(281)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 284: node = nodes.container(text)
							πF.SetLineno(284)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcontainer, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µnode = πTemp001
							// line 285: node['classes'].extend(classes)
							πF.SetLineno(285)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
								continue
							}
							πTemp003[0] = µclasses
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 286: self.add_name(node)
							πF.SetLineno(286)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 287: self.state.nested_parse(self.content, self.content_offset, node)
							πF.SetLineno(287)
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
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[2] = µnode
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
							// line 288: return [node]
							πF.SetLineno(288)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Container").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßContainer.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("body", Code)
}
