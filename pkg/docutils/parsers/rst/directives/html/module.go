package html

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßDirective := πg.InternStr("Directive")
		ßElement := πg.InternStr("Element")
		ßFilter := πg.InternStr("Filter")
		ßMeta := πg.InternStr("Meta")
		ßMetaBody := πg.InternStr("MetaBody")
		ßNameValueError := πg.InternStr("NameValueError")
		ßPreBibliographic := πg.InternStr("PreBibliographic")
		ßSMkwargs := πg.InternStr("SMkwargs")
		ßSpecial := πg.InternStr("Special")
		ßSpecializedBody := πg.InternStr("SpecializedBody")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßassert_has_content := πg.InternStr("assert_has_content")
		ßblock_text := πg.InternStr("block_text")
		ßchildren := πg.InternStr("children")
		ßcomponent := πg.InternStr("component")
		ßcomponents := πg.InternStr("components")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßdocument := πg.InternStr("document")
		ßend := πg.InternStr("end")
		ßerror := πg.InternStr("error")
		ßescape2null := πg.InternStr("escape2null")
		ßextract_name_value := πg.InternStr("extract_name_value")
		ßfield_marker := πg.InternStr("field_marker")
		ßformat := πg.InternStr("format")
		ßget_first_known_indented := πg.InternStr("get_first_known_indented")
		ßhas_content := πg.InternStr("has_content")
		ßhtml := πg.InternStr("html")
		ßinfo := πg.InternStr("info")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßlineno := πg.InternStr("lineno")
		ßliteral_block := πg.InternStr("literal_block")
		ßlower := πg.InternStr("lower")
		ßmeta := πg.InternStr("meta")
		ßname := πg.InternStr("name")
		ßnested_list_parse := πg.InternStr("nested_list_parse")
		ßnodes := πg.InternStr("nodes")
		ßnote_pending := πg.InternStr("note_pending")
		ßparent := πg.InternStr("parent")
		ßparse_field_marker := πg.InternStr("parse_field_marker")
		ßparsemeta := πg.InternStr("parsemeta")
		ßpending := πg.InternStr("pending")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrun := πg.InternStr("run")
		ßsplit := πg.InternStr("split")
		ßstate := πg.InternStr("state")
		ßstate_classes := πg.InternStr("state_classes")
		ßstate_machine := πg.InternStr("state_machine")
		ßstates := πg.InternStr("states")
		ßsys := πg.InternStr("sys")
		ßunescape := πg.InternStr("unescape")
		ßutils := πg.InternStr("utils")
		ßwriter := πg.InternStr("writer")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDirectives for typically HTML-specific constructs.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 11: import sys
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: from docutils import nodes, utils
			πF.SetLineno(12)
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
			// line 14: from docutils.parsers.rst import states
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.states"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßstates.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: from docutils.transforms import components
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.components"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßcomponents.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: class MetaBody(states.SpecializedBody):
			πF.SetLineno(18)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßSpecializedBody, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("MetaBody", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []πg.Param
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 20: class meta(nodes.Special, nodes.PreBibliographic, nodes.Element):
					πF.SetLineno(20)
					πTemp003 = make([]*πg.Object, 3)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßSpecial, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßPreBibliographic, nil); πE != nil {
						continue
					}
					πTemp003[1] = πTemp006
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßElement, nil); πE != nil {
						continue
					}
					πTemp003[2] = πTemp006
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("meta", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 21: """HTML-specific "meta" element."""
							πF.SetLineno(21)
							// line 21: """HTML-specific "meta" element."""
							πF.SetLineno(21)
							if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("HTML-specific \"meta\" element.").ToObject()); πE != nil {
								continue
							}
							// line 22: pass
							πF.SetLineno(22)
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp004 == nil {
						πTemp004 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("meta").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßmeta.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 24: def field_marker(self, match, context, next_state):
					πF.SetLineno(24)
					πTemp007 = make([]πg.Param, 4)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp007[1] = πg.Param{Name: "match", Def: nil}
					πTemp007[2] = πg.Param{Name: "context", Def: nil}
					πTemp007[3] = πg.Param{Name: "next_state", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("field_marker", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µcontext *πg.Object = πArgs[2]
						_ = µcontext
						var µnext_state *πg.Object = πArgs[3]
						_ = µnext_state
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
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
							// line 25: """Meta element."""
							πF.SetLineno(25)
							// line 26: node, blank_finish = self.parsemeta(match)
							πF.SetLineno(26)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							πTemp001[0] = µmatch
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparsemeta, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µnode = πTemp002
							µblank_finish = πTemp004
							// line 27: self.parent += node
							πF.SetLineno(27)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, µnode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp003); πE != nil {
								continue
							}
							// line 28: return [], next_state, []
							πF.SetLineno(28)
							πTemp001 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp001 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							πTemp002 = πg.NewTuple3(πTemp003, µnext_state, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßfield_marker.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 25: """Meta element."""
					πF.SetLineno(25)
					// line 25: """Meta element."""
					πF.SetLineno(25)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Meta element.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Meta element.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßfield_marker); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 30: def parsemeta(self, match):
					πF.SetLineno(30)
					πTemp007 = make([]πg.Param, 2)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp007[1] = πg.Param{Name: "match", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("parsemeta", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µindent *πg.Object = πg.UnboundLocal
						_ = µindent
						var µline_offset *πg.Object = πg.UnboundLocal
						_ = µline_offset
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µtokens *πg.Object = πg.UnboundLocal
						_ = µtokens
						var µattname *πg.Object = πg.UnboundLocal
						_ = µattname
						var µval *πg.Object = πg.UnboundLocal
						_ = µval
						var µtoken *πg.Object = πg.UnboundLocal
						_ = µtoken
						var µdetail *πg.Object = πg.UnboundLocal
						_ = µdetail
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Dict
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10:
								goto Label10
							case 4:
								goto Label4
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 31: name = self.parse_field_marker(match)
							πF.SetLineno(31)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							πTemp001[0] = µmatch
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_field_marker, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µname = πTemp003
							// line 32: name = utils.unescape(utils.escape2null(name))
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßescape2null, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunescape, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µname = πTemp002
							// line 33: indented, indent, line_offset, blank_finish = \
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_first_known_indented, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
								continue
							}
							µindented = πTemp003
							µindent = πTemp005
							µline_offset = πTemp006
							µblank_finish = πTemp007
							// line 35: node = self.meta()
							πF.SetLineno(35)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmeta, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnode = πTemp003
							// line 36: pending = nodes.pending(components.Filter,
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcomponents); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßFilter, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp008 = πg.NewDict()
							if πE = πTemp008.SetItem(πF, ßcomponent.ToObject(), ßwriter.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp008.SetItem(πF, ßformat.ToObject(), ßhtml.ToObject()); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							if πE = πTemp008.SetItem(πF, ßnodes.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp008.ToObject()
							πTemp001[1] = πTemp002
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
							// line 40: node['content'] = utils.unescape(utils.escape2null(
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp009[0] = µindented
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßescape2null, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunescape, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005 = ßcontent.ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, µindented); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp010).ToObject()
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label1
							}
							goto Label2
							// line 42: if not indented:
							πF.SetLineno(42)
						Label1:
							// line 43: line = self.state_machine.line
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßline, nil); πE != nil {
								continue
							}
							µline = πTemp003
							// line 44: msg = self.reporter.info(
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("No content for meta tag \"%s\".").ToObject(), µname); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[1] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp002
							// line 47: return msg, blank_finish
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µmsg, µblank_finish).ToObject()
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 48: tokens = name.split()
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µname, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtokens = πTemp003
							// line 49: try:
							πF.SetLineno(49)
							πF.PushCheckpoint(4)
							// line 50: attname, val = utils.extract_name_value(tokens[0])[0]
							πF.SetLineno(50)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µtokens, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßextract_name_value, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µattname = πTemp002
							µval = πTemp005
							// line 51: node[attname.lower()] = val
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µattname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, µnode, πTemp003, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßNameValueError, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label5
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 52: except utils.NameValueError:
							πF.SetLineno(52)
						Label5:
							// line 53: node['name'] = tokens[0]
							πF.SetLineno(53)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtokens, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005 = ßname.ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp005, πTemp002); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtokens, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp010 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µtoken = πTemp003
							}
							if πE != nil || !πTemp013 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 55: try:
							πF.SetLineno(55)
							πF.PushCheckpoint(10)
							// line 56: attname, val = utils.extract_name_value(token)[0]
							πF.SetLineno(56)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							πTemp001[0] = µtoken
							if πTemp006, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßextract_name_value, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp005); πE != nil {
								continue
							}
							µattname = πTemp003
							µval = πTemp006
							// line 57: node[attname.lower()] = val
							πF.SetLineno(57)
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µattname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, µnode, πTemp005, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßNameValueError, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp013 {
								goto Label11
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 58: except utils.NameValueError as detail:
							πF.SetLineno(58)
						Label11:
							µdetail = πTemp011.ToObject()
							// line 59: line = self.state_machine.line
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßline, nil); πE != nil {
								continue
							}
							µline = πTemp005
							// line 60: msg = self.reporter.error(
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtoken, "token"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdetail, "detail"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µtoken, µdetail).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Error parsing meta tag attribute \"%s\": %s.").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[1] = µline
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp003
							// line 63: return msg, blank_finish
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µmsg, µblank_finish).ToObject()
							πR = πTemp003
							continue
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 64: self.document.note_pending(pending)
							πF.SetLineno(64)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 65: return pending, blank_finish
							πF.SetLineno(65)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µpending, µblank_finish).ToObject()
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
					if πE = πClass.SetItem(πF, ßparsemeta.ToObject(), πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MetaBody").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMetaBody.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 68: class Meta(Directive):
			πF.SetLineno(68)
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
			_, πE = πg.NewCode("Meta", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 70: has_content = True
					πF.SetLineno(70)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 72: SMkwargs = {'state_classes': (MetaBody,)}
					πF.SetLineno(72)
					πTemp002 = πg.NewDict()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßMetaBody); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple1(πTemp003).ToObject()
					if πE = πTemp002.SetItem(πF, ßstate_classes.ToObject(), πTemp001); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßSMkwargs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 74: def run(self):
					πF.SetLineno(74)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/html.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µnew_line_offset *πg.Object = πg.UnboundLocal
						_ = µnew_line_offset
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
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
							// line 75: self.assert_has_content()
							πF.SetLineno(75)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassert_has_content, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 76: node = nodes.Element()
							πF.SetLineno(76)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßElement, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnode = πTemp001
							// line 77: new_line_offset, blank_finish = self.state.nested_list_parse(
							πF.SetLineno(77)
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßSMkwargs, nil); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"initial_state", ßMetaBody.ToObject()},
								{"blank_finish", πTemp001},
								{"state_machine_kwargs", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnested_list_parse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µnew_line_offset = πTemp002
							µblank_finish = πTemp005
							if πE = πg.CheckLocal(πF, µnew_line_offset, "new_line_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µnew_line_offset, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.NE(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 81: if (new_line_offset - self.content_offset) != len(self.content):
							πF.SetLineno(81)
						Label1:
							// line 83: error = self.state_machine.reporter.error(
							πF.SetLineno(83)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("Invalid meta directive.").ToObject()
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µerror = πTemp002
							// line 87: node += error
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µnode, µerror); πE != nil {
								continue
							}
							µnode = πTemp001
							goto Label2
						Label2:
							// line 88: return node.children
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Meta").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMeta.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("html", Code)
}
