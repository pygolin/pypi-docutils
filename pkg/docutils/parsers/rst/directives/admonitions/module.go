package admonitions

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAdmonition := πg.InternStr("Admonition")
		ßAttention := πg.InternStr("Attention")
		ßBaseAdmonition := πg.InternStr("BaseAdmonition")
		ßCaution := πg.InternStr("Caution")
		ßDanger := πg.InternStr("Danger")
		ßDirective := πg.InternStr("Directive")
		ßError := πg.InternStr("Error")
		ßHint := πg.InternStr("Hint")
		ßImportant := πg.InternStr("Important")
		ßNone := πg.InternStr("None")
		ßNote := πg.InternStr("Note")
		ßTip := πg.InternStr("Tip")
		ßTrue := πg.InternStr("True")
		ßWarning := πg.InternStr("Warning")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd_name := πg.InternStr("add_name")
		ßadmonition := πg.InternStr("admonition")
		ßarguments := πg.InternStr("arguments")
		ßassert_has_content := πg.InternStr("assert_has_content")
		ßattention := πg.InternStr("attention")
		ßcaution := πg.InternStr("caution")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßdanger := πg.InternStr("danger")
		ßdirectives := πg.InternStr("directives")
		ßerror := πg.InternStr("error")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßhas_content := πg.InternStr("has_content")
		ßhint := πg.InternStr("hint")
		ßimportant := πg.InternStr("important")
		ßinline_text := πg.InternStr("inline_text")
		ßjoin := πg.InternStr("join")
		ßline := πg.InternStr("line")
		ßlineno := πg.InternStr("lineno")
		ßmake_id := πg.InternStr("make_id")
		ßname := πg.InternStr("name")
		ßnested_parse := πg.InternStr("nested_parse")
		ßnode_class := πg.InternStr("node_class")
		ßnodes := πg.InternStr("nodes")
		ßnote := πg.InternStr("note")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptions := πg.InternStr("options")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrequired_arguments := πg.InternStr("required_arguments")
		ßrun := πg.InternStr("run")
		ßset_classes := πg.InternStr("set_classes")
		ßsource := πg.InternStr("source")
		ßstate := πg.InternStr("state")
		ßstate_machine := πg.InternStr("state_machine")
		ßstates := πg.InternStr("states")
		ßtip := πg.InternStr("tip")
		ßtitle := πg.InternStr("title")
		ßunchanged := πg.InternStr("unchanged")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nAdmonition directives.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 12: from docutils.parsers.rst import Directive
			πF.SetLineno(12)
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
			// line 13: from docutils.parsers.rst import states, directives
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.states"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßstates.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: from docutils.parsers.rst.roles import set_classes
			πF.SetLineno(14)
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
			// line 15: from docutils import nodes
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: class BaseAdmonition(Directive):
			πF.SetLineno(18)
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
			_, πE = πg.NewCode("BaseAdmonition", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 20: final_argument_whitespace = True
					πF.SetLineno(20)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 21: option_spec = {'class': directives.class_option,
					πF.SetLineno(21)
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
					// line 23: has_content = True
					πF.SetLineno(23)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 25: node_class = None
					πF.SetLineno(25)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnode_class.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 26: """Subclasses must set this to the appropriate admonition node class."""
					πF.SetLineno(26)
					// line 26: """Subclasses must set this to the appropriate admonition node class."""
					πF.SetLineno(26)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Subclasses must set this to the appropriate admonition node class.").ToObject()); πE != nil {
						continue
					}
					// line 28: def run(self):
					πF.SetLineno(28)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µadmonition_node *πg.Object = πg.UnboundLocal
						_ = µadmonition_node
						var µtitle_text *πg.Object = πg.UnboundLocal
						_ = µtitle_text
						var µtextnodes *πg.Object = πg.UnboundLocal
						_ = µtextnodes
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
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
						var πTemp006 bool
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
							// line 29: set_classes(self.options)
							πF.SetLineno(29)
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
							// line 30: self.assert_has_content()
							πF.SetLineno(30)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 31: text = '\n'.join(self.content)
							πF.SetLineno(31)
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
							// line 32: admonition_node = self.node_class(text, **self.options)
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnode_class, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp001, nil, nil, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µadmonition_node = πTemp004
							// line 33: self.add_name(admonition_node)
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							πTemp001[0] = µadmonition_node
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnode_class, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßadmonition, nil); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 34: if self.node_class is nodes.admonition:
							πF.SetLineno(34)
						Label1:
							// line 35: title_text = self.arguments[0]
							πF.SetLineno(35)
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
							µtitle_text = πTemp003
							// line 36: textnodes, messages = self.state.inline_text(title_text,
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp001[0] = µtitle_text
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
							// line 38: title = nodes.title(title_text, '', *textnodes)
							πF.SetLineno(38)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp001[0] = µtitle_text
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtextnodes, "textnodes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µtextnodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtitle = πTemp002
							// line 39: title.source, title.line = (
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_source_and_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtitle, ßsource, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtitle, ßline, πTemp004); πE != nil {
								continue
							}
							// line 41: admonition_node += title
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µadmonition_node, µtitle); πE != nil {
								continue
							}
							µadmonition_node = πTemp002
							// line 42: admonition_node += messages
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µadmonition_node, µmessages); πE != nil {
								continue
							}
							µadmonition_node = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, ßclasses.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 43: if not 'classes' in self.options:
							πF.SetLineno(43)
						Label3:
							// line 44: admonition_node['classes'] += ['admonition-' +
							πF.SetLineno(44)
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µadmonition_node, πTemp002); πE != nil {
								continue
							}
							πTemp001 = make([]*πg.Object, 1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtitle_text, "title_text"); πE != nil {
								continue
							}
							πTemp007[0] = µtitle_text
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmake_id, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Add(πF, πg.NewStr("admonition-").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πTemp004, πE = πg.IAdd(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µadmonition_node, πTemp005, πTemp004); πE != nil {
								continue
							}
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 46: self.state.nested_parse(self.content, self.content_offset,
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							πTemp001[2] = µadmonition_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 48: return [admonition_node]
							πF.SetLineno(48)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µadmonition_node, "admonition_node"); πE != nil {
								continue
							}
							πTemp001[0] = µadmonition_node
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BaseAdmonition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseAdmonition.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 51: class Admonition(BaseAdmonition):
			πF.SetLineno(51)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Admonition", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 53: required_arguments = 1
					πF.SetLineno(53)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 54: node_class = nodes.admonition
					πF.SetLineno(54)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadmonition, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Admonition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAdmonition.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 57: class Attention(BaseAdmonition):
			πF.SetLineno(57)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Attention", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 59: node_class = nodes.attention
					πF.SetLineno(59)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßattention, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Attention").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAttention.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 62: class Caution(BaseAdmonition):
			πF.SetLineno(62)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Caution", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 64: node_class = nodes.caution
					πF.SetLineno(64)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcaution, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Caution").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCaution.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 67: class Danger(BaseAdmonition):
			πF.SetLineno(67)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Danger", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 69: node_class = nodes.danger
					πF.SetLineno(69)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdanger, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Danger").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDanger.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 72: class Error(BaseAdmonition):
			πF.SetLineno(72)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Error", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 74: node_class = nodes.error
					πF.SetLineno(74)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Error").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 77: class Hint(BaseAdmonition):
			πF.SetLineno(77)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Hint", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 79: node_class = nodes.hint
					πF.SetLineno(79)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßhint, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Hint").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHint.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 82: class Important(BaseAdmonition):
			πF.SetLineno(82)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Important", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 84: node_class = nodes.important
					πF.SetLineno(84)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßimportant, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Important").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßImportant.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 87: class Note(BaseAdmonition):
			πF.SetLineno(87)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Note", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 89: node_class = nodes.note
					πF.SetLineno(89)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnote, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Note").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNote.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 92: class Tip(BaseAdmonition):
			πF.SetLineno(92)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Tip", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 94: node_class = nodes.tip
					πF.SetLineno(94)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtip, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Tip").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTip.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 97: class Warning(BaseAdmonition):
			πF.SetLineno(97)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseAdmonition); πE != nil {
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
			_, πE = πg.NewCode("Warning", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/admonitions.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 99: node_class = nodes.warning
					πF.SetLineno(99)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwarning, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Warning").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWarning.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("admonitions", Code)
}
