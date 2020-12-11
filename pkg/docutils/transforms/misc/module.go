package misc

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßCallBack := πg.InternStr("CallBack")
		ßClassAttribute := πg.InternStr("ClassAttribute")
		ßInvisible := πg.InternStr("Invisible")
		ßNone := πg.InternStr("None")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßTransitions := πg.InternStr("Transitions")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßapply := πg.InternStr("apply")
		ßcallback := πg.InternStr("callback")
		ßclass := πg.InternStr("class")
		ßclasses := πg.InternStr("classes")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdetails := πg.InternStr("details")
		ßdirective := πg.InternStr("directive")
		ßdocument := πg.InternStr("document")
		ßerror := πg.InternStr("error")
		ßindex := πg.InternStr("index")
		ßinsert := πg.InternStr("insert")
		ßisinstance := πg.InternStr("isinstance")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßliteral_block := πg.InternStr("literal_block")
		ßnodes := πg.InternStr("nodes")
		ßparent := πg.InternStr("parent")
		ßrange := πg.InternStr("range")
		ßrawsource := πg.InternStr("rawsource")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßremove := πg.InternStr("remove")
		ßreplace_self := πg.InternStr("replace_self")
		ßreporter := πg.InternStr("reporter")
		ßsection := πg.InternStr("section")
		ßsource := πg.InternStr("source")
		ßstartnode := πg.InternStr("startnode")
		ßsubtitle := πg.InternStr("subtitle")
		ßsystem_message := πg.InternStr("system_message")
		ßtitle := πg.InternStr("title")
		ßtransition := πg.InternStr("transition")
		ßtraverse := πg.InternStr("traverse")
		ßvisit_transition := πg.InternStr("visit_transition")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nMiscellaneous transforms.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 11: from docutils import nodes
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: from docutils.transforms import Transform, TransformError
			πF.SetLineno(12)
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
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: class CallBack(Transform):
			πF.SetLineno(15)
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
			_, πE = πg.NewCode("CallBack", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 17: """
					πF.SetLineno(17)
					// line 17: """
					πF.SetLineno(17)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Inserts a callback into a document.  The callback is called when the\n    transform is applied, which is determined by its priority.\n\n    For use with `nodes.pending` elements.  Requires a ``details['callback']``\n    entry, a bound method or function which takes one parameter: the pending\n    node.  Other data can be stored in the ``details`` attribute or in the\n    object hosting the callback method.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 27: default_priority = 990
					πF.SetLineno(27)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(990).ToObject()); πE != nil {
						continue
					}
					// line 29: def apply(self):
					πF.SetLineno(29)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var πTemp001 *πg.Object
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
							// line 30: pending = self.startnode
							πF.SetLineno(30)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							µpending = πTemp001
							// line 31: pending.details['callback'](pending)
							πF.SetLineno(31)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							πTemp001 = ßcallback.ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 32: pending.parent.remove(pending)
							πF.SetLineno(32)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("CallBack").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCallBack.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 35: class ClassAttribute(Transform):
			πF.SetLineno(35)
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
			_, πE = πg.NewCode("ClassAttribute", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 37: """
					πF.SetLineno(37)
					// line 37: """
					πF.SetLineno(37)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Move the \"class\" attribute specified in the \"pending\" node into the\n    immediately following non-comment element.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 42: default_priority = 210
					πF.SetLineno(42)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(210).ToObject()); πE != nil {
						continue
					}
					// line 44: def apply(self):
					πF.SetLineno(44)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var µparent *πg.Object = πg.UnboundLocal
						_ = µparent
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µelement *πg.Object = πg.UnboundLocal
						_ = µelement
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 πg.KWArgs
						_ = πTemp011
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
							// line 45: pending = self.startnode
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							µpending = πTemp001
							// line 46: parent = pending.parent
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							µparent = πTemp001
							// line 47: child = pending
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							µchild = µpending
							// line 48: while parent:
							πF.SetLineno(48)
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
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µparent); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp006[0] = µchild
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µparent, ßindex, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp005, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πTemp006[0] = µparent
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[1] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp003 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp005, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µindex = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 51: element = parent[index]
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp005 = µindex
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µparent, πTemp005); πE != nil {
								continue
							}
							µelement = πTemp007
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp004[0] = µelement
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßInvisible, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp005 = πTemp008
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label7
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp004[0] = µelement
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßsystem_message, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp005 = πTemp008
						Label7:
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label8
							}
							goto Label9
							// line 52: if (isinstance(element, nodes.Invisible) or
							πF.SetLineno(52)
						Label8:
							// line 54: continue
							πF.SetLineno(54)
							continue
							goto Label9
						Label9:
							// line 55: element['classes'] += pending.details['class']
							πF.SetLineno(55)
							πTemp005 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µelement, πTemp005); πE != nil {
								continue
							}
							πTemp005 = ßclass.ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp010 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µelement, πTemp010, πTemp005); πE != nil {
								continue
							}
							// line 56: pending.parent.remove(pending)
							πF.SetLineno(56)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp004[0] = µpending
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßremove, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 57: return
							πF.SetLineno(57)
							πR = πg.None
							continue
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							// line 60: child = parent
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							µchild = µparent
							// line 61: parent = parent.parent
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µparent, ßparent, nil); πE != nil {
								continue
							}
							µparent = πTemp005
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 62: error = self.document.reporter.error(
							πF.SetLineno(62)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = ßdirective.ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("No suitable element following \"%s\" directive").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßline, nil); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"line", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßerror, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µerror = πTemp005
							// line 67: pending.replace_self(error)
							πF.SetLineno(67)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp004[0] = µerror
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ClassAttribute").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßClassAttribute.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 70: class Transitions(Transform):
			πF.SetLineno(70)
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
			_, πE = πg.NewCode("Transitions", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 72: """
					πF.SetLineno(72)
					// line 72: """
					πF.SetLineno(72)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Move transitions at the end of sections up the tree.  Complain\n    on transitions after a title, at the beginning or end of the\n    document, and after another transition.\n\n    For example, transform this::\n\n        <section>\n            ...\n            <transition>\n        <section>\n            ...\n\n    into this::\n\n        <section>\n            ...\n        <transition>\n        <section>\n            ...\n    ").ToObject()); πE != nil {
						continue
					}
					// line 94: default_priority = 830
					πF.SetLineno(94)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(830).ToObject()); πE != nil {
						continue
					}
					// line 96: def apply(self):
					πF.SetLineno(96)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
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
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtransition, nil); πE != nil {
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
								µnode = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 98: self.visit_transition(node)
							πF.SetLineno(98)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvisit_transition, nil); πE != nil {
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
					// line 100: def visit_transition(self, node):
					πF.SetLineno(100)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("visit_transition", "/usr/lib/python2.7/site-packages/docutils/transforms/misc.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µsibling *πg.Object = πg.UnboundLocal
						_ = µsibling
						var πTemp001 []*πg.Object
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 πg.KWArgs
						_ = πTemp013
						var πTemp014 []*πg.Object
						_ = πTemp014
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 13:
								goto Label13
							case 14:
								goto Label14
							default:
								panic("unexpected function state")
							}
							// line 101: index = node.parent.index(node)
							πF.SetLineno(101)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp002
							// line 102: error = None
							πF.SetLineno(102)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µerror = πTemp002
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µindex, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßtitle, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Eq(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label3
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp008 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp008); πE != nil {
								continue
							}
							πTemp001[0] = πTemp011
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp008, ßsubtitle, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp011
							if πTemp008, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp007 = πTemp011
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Eq(πF, µindex, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
						Label4:
							πTemp006 = πTemp007
						Label3:
							πTemp003 = πTemp006
						Label2:
							πTemp002 = πTemp003
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtransition, nil); πE != nil {
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
								goto Label6
							}
							goto Label7
							// line 103: if (index == 0 or
							πF.SetLineno(103)
						Label5:
							// line 108: assert (isinstance(node.parent, nodes.document) or
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp006
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßsection, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp006
						Label8:
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							// line 110: error = self.document.reporter.error(
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Document or section may not begin with a transition.").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßline, nil); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"source", πTemp002},
								{"line", πTemp003},
							}
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
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp013); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							goto Label7
							// line 113: elif isinstance(node.parent[index - 1], nodes.transition):
							πF.SetLineno(113)
						Label6:
							// line 114: error = self.document.reporter.error(
							πF.SetLineno(114)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("At least one body element must separate transitions; adjacent transitions are not allowed.").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßline, nil); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"source", πTemp002},
								{"line", πTemp003},
							}
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
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp013); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µerror); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 118: if error:
							πF.SetLineno(118)
						Label9:
							// line 120: node.parent.insert(index, error)
							πF.SetLineno(120)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[1] = µerror
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 121: index += 1
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µindex = πTemp002
							goto Label10
						Label10:
							// line 122: assert index < len(node.parent)
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, µindex, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µindex, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 123: if index != len(node.parent) - 1:
							πF.SetLineno(123)
						Label11:
							// line 125: return
							πF.SetLineno(125)
							πR = πg.None
							continue
							goto Label12
						Label12:
							// line 127: sibling = node
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							µsibling = µnode
							// line 129: while index == len(sibling.parent) - 1:
							πF.SetLineno(129)
							πF.PushCheckpoint(14)
							πTemp004 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µsibling, ßparent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µindex, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(13)
							// line 130: sibling = sibling.parent
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsibling, ßparent, nil); πE != nil {
								continue
							}
							µsibling = πTemp002
							if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsibling, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp006).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label16
							}
							goto Label17
							// line 132: if sibling.parent is None:
							πF.SetLineno(132)
						Label16:
							// line 135: error = self.document.reporter.error(
							πF.SetLineno(135)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Document may not end with a transition.").ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßline, nil); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"line", πTemp002},
							}
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
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp013); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							// line 138: node.parent.insert(node.parent.index(node) + 1, error)
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(2)
							πTemp014 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp014[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[1] = µerror
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 139: return
							πF.SetLineno(139)
							πR = πg.None
							continue
							goto Label17
						Label17:
							// line 140: index = sibling.parent.index(sibling)
							πF.SetLineno(140)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
								continue
							}
							πTemp001[0] = µsibling
							if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsibling, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp002
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 142: node.parent.remove(node)
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 144: sibling.parent.insert(index + 1, node)
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[1] = µnode
							if πE = πg.CheckLocal(πF, µsibling, "sibling"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsibling, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvisit_transition.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Transitions").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransitions.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("misc", Code)
}
