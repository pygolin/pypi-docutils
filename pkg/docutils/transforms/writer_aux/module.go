package writer_aux

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/writer_aux.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAdmonition := πg.InternStr("Admonition")
		ßAdmonitions := πg.InternStr("Admonitions")
		ßCompound := πg.InternStr("Compound")
		ßFalse := πg.InternStr("False")
		ßInvisible := πg.InternStr("Invisible")
		ßTransform := πg.InternStr("Transform")
		ßTrue := πg.InternStr("True")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadmonition := πg.InternStr("admonition")
		ßappend := πg.InternStr("append")
		ßapply := πg.InternStr("apply")
		ßattributes := πg.InternStr("attributes")
		ßchildren := πg.InternStr("children")
		ßclasses := πg.InternStr("classes")
		ßcompound := πg.InternStr("compound")
		ßcontinued := πg.InternStr("continued")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdocument := πg.InternStr("document")
		ßget_language := πg.InternStr("get_language")
		ßinsert := πg.InternStr("insert")
		ßisinstance := πg.InternStr("isinstance")
		ßlabels := πg.InternStr("labels")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßnodes := πg.InternStr("nodes")
		ßrawsource := πg.InternStr("rawsource")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreplace_self := πg.InternStr("replace_self")
		ßreporter := πg.InternStr("reporter")
		ßsettings := πg.InternStr("settings")
		ßtitle := πg.InternStr("title")
		ßtraverse := πg.InternStr("traverse")
		ßutils := πg.InternStr("utils")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nAuxiliary transforms mainly to be used by Writer components.\n\nThis module is called \"writer_aux\" because otherwise there would be\nconflicting imports like this one::\n\n    from docutils import writers\n    from docutils.transforms import writers\n").ToObject()); πE != nil {
				continue
			}
			// line 15: __docformat__ = 'reStructuredText'
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 17: from docutils import nodes, utils, languages
			πF.SetLineno(17)
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: from docutils.transforms import Transform
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransform); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransform.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 21: class Compound(Transform):
			πF.SetLineno(21)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("Compound", "/usr/lib/python2.7/site-packages/docutils/transforms/writer_aux.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 23: """
					πF.SetLineno(23)
					// line 23: """
					πF.SetLineno(23)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Flatten all compound paragraphs.  For example, transform ::\n\n        <compound>\n            <paragraph>\n            <literal_block>\n            <paragraph>\n\n    into ::\n\n        <paragraph>\n        <literal_block classes=\"continued\">\n        <paragraph classes=\"continued\">\n    ").ToObject()); πE != nil {
						continue
					}
					// line 38: default_priority = 910
					πF.SetLineno(38)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(910).ToObject()); πE != nil {
						continue
					}
					// line 40: def apply(self):
					πF.SetLineno(40)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/writer_aux.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcompound *πg.Object = πg.UnboundLocal
						_ = µcompound
						var µfirst_child *πg.Object = πg.UnboundLocal
						_ = µfirst_child
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcompound, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µcompound = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 42: first_child = True
							πF.SetLineno(42)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µfirst_child = πTemp003
							if πE = πg.CheckLocal(πF, µcompound, "compound"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µcompound); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µchild = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)
							if πE = πg.CheckLocal(πF, µfirst_child, "first_child"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µfirst_child); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							goto Label8
							// line 44: if first_child:
							πF.SetLineno(44)
						Label7:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp002[0] = µchild
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßInvisible, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 45: if not isinstance(child, nodes.Invisible):
							πF.SetLineno(45)
						Label10:
							// line 46: first_child = False
							πF.SetLineno(46)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µfirst_child = πTemp004
							goto Label11
						Label11:
							goto Label9
						Label8:
							// line 48: child['classes'].append('continued')
							πF.SetLineno(48)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßcontinued.ToObject()
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µchild, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp008, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label9
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 50: compound.replace_self(compound[:])
							πF.SetLineno(50)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcompound, "compound"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µcompound, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µcompound, "compound"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcompound, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
			// line 53: class Admonitions(Transform):
			πF.SetLineno(53)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
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
			_, πE = πg.NewCode("Admonitions", "/usr/lib/python2.7/site-packages/docutils/transforms/writer_aux.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 55: """
					πF.SetLineno(55)
					// line 55: """
					πF.SetLineno(55)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Transform specific admonitions, like this:\n\n        <note>\n            <paragraph>\n                 Note contents ...\n\n    into generic admonitions, like this::\n\n        <admonition classes=\"note\">\n            <title>\n                Note\n            <paragraph>\n                Note contents ...\n\n    The admonition title is localized.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 73: default_priority = 920
					πF.SetLineno(73)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(920).ToObject()); πE != nil {
						continue
					}
					// line 75: def apply(self):
					πF.SetLineno(75)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/writer_aux.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlanguage *πg.Object = πg.UnboundLocal
						_ = µlanguage
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µnode_name *πg.Object = πg.UnboundLocal
						_ = µnode_name
						var µadmonition *πg.Object = πg.UnboundLocal
						_ = µadmonition
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
							// line 76: language = languages.get_language(self.document.settings.language_code,
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlanguage_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlanguages); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlanguage = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßAdmonition, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µnode = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 79: node_name = node.__class__.__name__
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							µnode_name = πTemp004
							// line 81: node['classes'].append(node_name)
							πF.SetLineno(81)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode_name, "node_name"); πE != nil {
								continue
							}
							πTemp001[0] = µnode_name
							πTemp003 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßadmonition, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
							// line 82: if not isinstance(node, nodes.admonition):
							πF.SetLineno(82)
						Label4:
							// line 84: admonition = nodes.admonition(node.rawsource, *node.children,
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßadmonition, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp008, πTemp001, πTemp003, nil, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µadmonition = πTemp007
							// line 86: title = nodes.title('', language.labels[node_name])
							πF.SetLineno(86)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µnode_name, "node_name"); πE != nil {
								continue
							}
							πTemp003 = µnode_name
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µlanguage, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtitle = πTemp003
							// line 87: admonition.insert(0, title)
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp001[1] = µtitle
							if πE = πg.CheckLocal(πF, µadmonition, "admonition"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µadmonition, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 88: node.replace_self(admonition)
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µadmonition, "admonition"); πE != nil {
								continue
							}
							πTemp001[0] = µadmonition
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Admonitions").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAdmonitions.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("writer_aux", Code)
}
