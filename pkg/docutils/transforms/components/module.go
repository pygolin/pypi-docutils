package components

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/time"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/components.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßApplicationError := πg.InternStr("ApplicationError")
		ßDataError := πg.InternStr("DataError")
		ßFilter := πg.InternStr("Filter")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßapply := πg.InternStr("apply")
		ßcomponent := πg.InternStr("component")
		ßcomponents := πg.InternStr("components")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdetails := πg.InternStr("details")
		ßdocument := πg.InternStr("document")
		ßformat := πg.InternStr("format")
		ßnodes := πg.InternStr("nodes")
		ßos := πg.InternStr("os")
		ßparent := πg.InternStr("parent")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßremove := πg.InternStr("remove")
		ßreplace_self := πg.InternStr("replace_self")
		ßstartnode := πg.InternStr("startnode")
		ßsupports := πg.InternStr("supports")
		ßsys := πg.InternStr("sys")
		ßtime := πg.InternStr("time")
		ßtransformer := πg.InternStr("transformer")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDocutils component-related transforms.\n").ToObject()); πE != nil {
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
			// line 12: import os
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: import re
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import time
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: from docutils import nodes, utils
			πF.SetLineno(15)
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
			// line 16: from docutils import ApplicationError, DataError
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßApplicationError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßApplicationError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDataError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDataError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: from docutils.transforms import Transform, TransformError
			πF.SetLineno(17)
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
			// line 20: class Filter(Transform):
			πF.SetLineno(20)
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
			_, πE = πg.NewCode("Filter", "/usr/lib/python2.7/site-packages/docutils/transforms/components.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 22: """
					πF.SetLineno(22)
					// line 22: """
					πF.SetLineno(22)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Include or exclude elements which depend on a specific Docutils component.\n\n    For use with `nodes.pending` elements.  A \"pending\" element's dictionary\n    attribute ``details`` must contain the keys \"component\" and \"format\".  The\n    value of ``details['component']`` must match the type name of the\n    component the elements depend on (e.g. \"writer\").  The value of\n    ``details['format']`` is the name of a specific format or context of that\n    component (e.g. \"html\").  If the matching Docutils component supports that\n    format or context, the \"pending\" element is replaced by the contents of\n    ``details['nodes']`` (a list of nodes); otherwise, the \"pending\" element\n    is removed.\n\n    For example, the reStructuredText \"meta\" directive creates a \"pending\"\n    element containing a \"meta\" element (in ``pending.details['nodes']``).\n    Only writers (``pending.details['component'] == 'writer'``) supporting the\n    \"html\" format (``pending.details['format'] == 'html'``) will include the\n    \"meta\" element; it will be deleted from the output of all other writers.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 42: default_priority = 780
					πF.SetLineno(42)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(780).ToObject()); πE != nil {
						continue
					}
					// line 44: def apply(self):
					πF.SetLineno(44)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/components.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var µcomponent_type *πg.Object = πg.UnboundLocal
						_ = µcomponent_type
						var µformat *πg.Object = πg.UnboundLocal
						_ = µformat
						var µcomponent *πg.Object = πg.UnboundLocal
						_ = µcomponent
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
							// line 45: pending = self.startnode
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							µpending = πTemp001
							// line 46: component_type = pending.details['component'] # 'reader' or 'writer'
							πF.SetLineno(46)
							πTemp001 = ßcomponent.ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µcomponent_type = πTemp002
							// line 47: format = pending.details['format']
							πF.SetLineno(47)
							πTemp001 = ßformat.ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µformat = πTemp002
							// line 48: component = self.document.transformer.components[component_type]
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µcomponent_type, "component_type"); πE != nil {
								continue
							}
							πTemp001 = µcomponent_type
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtransformer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßcomponents, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µcomponent = πTemp002
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							πTemp005[0] = µformat
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcomponent, ßsupports, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 49: if component.supports(format):
							πF.SetLineno(49)
						Label1:
							// line 50: pending.replace_self(pending.details['nodes'])
							πF.SetLineno(50)
							πTemp005 = πF.MakeArgs(1)
							πTemp001 = ßnodes.ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpending, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label3
						Label2:
							// line 52: pending.parent.remove(pending)
							πF.SetLineno(52)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp005[0] = µpending
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Filter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFilter.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("components", Code)
}
