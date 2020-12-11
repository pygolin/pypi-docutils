package transforms

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßApplicationError := πg.InternStr("ApplicationError")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßTransformSpec := πg.InternStr("TransformSpec")
		ßTransformer := πg.InternStr("Transformer")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd_pending := πg.InternStr("add_pending")
		ßadd_transform := πg.InternStr("add_transform")
		ßadd_transforms := πg.InternStr("add_transforms")
		ßappend := πg.InternStr("append")
		ßapplied := πg.InternStr("applied")
		ßapply := πg.InternStr("apply")
		ßapply_transforms := πg.InternStr("apply_transforms")
		ßattach_observer := πg.InternStr("attach_observer")
		ßcomponent_type := πg.InternStr("component_type")
		ßcomponents := πg.InternStr("components")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdocument := πg.InternStr("document")
		ßextend := πg.InternStr("extend")
		ßget_language := πg.InternStr("get_language")
		ßget_priority_string := πg.InternStr("get_priority_string")
		ßget_transforms := πg.InternStr("get_transforms")
		ßlanguage := πg.InternStr("language")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßnote_transform_message := πg.InternStr("note_transform_message")
		ßobject := πg.InternStr("object")
		ßpop := πg.InternStr("pop")
		ßpopulate_from_components := πg.InternStr("populate_from_components")
		ßpriority := πg.InternStr("priority")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßreverse := πg.InternStr("reverse")
		ßserialno := πg.InternStr("serialno")
		ßsettings := πg.InternStr("settings")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßstartnode := πg.InternStr("startnode")
		ßtransform := πg.InternStr("transform")
		ßtransforms := πg.InternStr("transforms")
		ßunknown_reference_resolvers := πg.InternStr("unknown_reference_resolvers")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis package contains modules for standard tree transforms available\nto Docutils components. Tree transforms serve a variety of purposes:\n\n- To tie up certain syntax-specific \"loose ends\" that remain after the\n  initial parsing of the input plaintext. These transforms are used to\n  supplement a limited syntax.\n\n- To automate the internal linking of the document tree (hyperlink\n  references, footnote references, etc.).\n\n- To extract useful information from the document tree. These\n  transforms may be used to construct (for example) indexes and tables\n  of contents.\n\nEach transform is an optional step that a Docutils component may\nchoose to perform on the parsed document.\n").ToObject()); πE != nil {
				continue
			}
			// line 24: __docformat__ = 'reStructuredText'
			πF.SetLineno(24)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 27: from docutils import languages, ApplicationError, TransformSpec
			πF.SetLineno(27)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
				continue
			}
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
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformSpec); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformSpec.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 30: class TransformError(ApplicationError): pass
			πF.SetLineno(30)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßApplicationError); πE != nil {
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
			_, πE = πg.NewCode("TransformError", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 30: class TransformError(ApplicationError): pass
					πF.SetLineno(30)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TransformError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 33: class Transform(object):
			πF.SetLineno(33)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("Transform", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 35: """
					πF.SetLineno(35)
					// line 35: """
					πF.SetLineno(35)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Docutils transform component abstract base class.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 39: default_priority = None
					πF.SetLineno(39)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 40: """Numerical priority of this transform, 0 through 999 (override)."""
					πF.SetLineno(40)
					// line 42: def __init__(self, document, startnode=None):
					πF.SetLineno(42)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "startnode", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µstartnode *πg.Object = πArgs[2]
						_ = µstartnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 43: """
							πF.SetLineno(43)
							// line 47: self.document = document
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdocument); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp001); πE != nil {
								continue
							}
							// line 48: """The document tree to transform."""
							πF.SetLineno(48)
							// line 50: self.startnode = startnode
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstartnode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstartnode, πTemp001); πE != nil {
								continue
							}
							// line 51: """Node from which to begin the transform.  For many transforms which
							πF.SetLineno(51)
							// line 55: self.language = languages.get_language(
							πF.SetLineno(55)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßlanguage_code, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlanguages); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage, πTemp003); πE != nil {
								continue
							}
							// line 57: """Language module local to this document."""
							πF.SetLineno(57)
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
					// line 43: """
					πF.SetLineno(43)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Initial setup for in-place document transforms.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 59: def apply(self, **kwargs):
					πF.SetLineno(59)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkwargs *πg.Object = πArgs[1]
						_ = µkwargs
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
							// line 60: """Override to apply the transform to the document tree."""
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("subclass must override this method").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 61: raise NotImplementedError('subclass must override this method')
							πF.SetLineno(61)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 60: """Override to apply the transform to the document tree."""
					πF.SetLineno(60)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Override to apply the transform to the document tree.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßapply); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Transform").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransform.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 64: class Transformer(TransformSpec):
			πF.SetLineno(64)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransformSpec); πE != nil {
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
			_, πE = πg.NewCode("Transformer", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 66: """
					πF.SetLineno(66)
					// line 66: """
					πF.SetLineno(66)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Stores transforms (`Transform` classes) and applies them to document\n    trees.  Also keeps track of components by component type name.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 71: def __init__(self, document):
					πF.SetLineno(71)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
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
							// line 72: self.transforms = []
							πF.SetLineno(72)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtransforms, πTemp003); πE != nil {
								continue
							}
							// line 73: """List of transforms to apply.  Each item is a 4-tuple:
							πF.SetLineno(73)
							// line 77: self.unknown_reference_resolvers = []
							πF.SetLineno(77)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßunknown_reference_resolvers, πTemp003); πE != nil {
								continue
							}
							// line 78: """List of hook functions which assist in resolving references"""
							πF.SetLineno(78)
							// line 80: self.document = document
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdocument); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdocument, πTemp002); πE != nil {
								continue
							}
							// line 81: """The `nodes.document` object this Transformer is attached to."""
							πF.SetLineno(81)
							// line 83: self.applied = []
							πF.SetLineno(83)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßapplied, πTemp003); πE != nil {
								continue
							}
							// line 84: """Transforms already applied, in order."""
							πF.SetLineno(84)
							// line 86: self.sorted = 0
							πF.SetLineno(86)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsorted, πTemp002); πE != nil {
								continue
							}
							// line 87: """Boolean: is `self.tranforms` sorted?"""
							πF.SetLineno(87)
							// line 89: self.components = {}
							πF.SetLineno(89)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcomponents, πTemp003); πE != nil {
								continue
							}
							// line 90: """Mapping of component type name to component object.  Set by
							πF.SetLineno(90)
							// line 93: self.serialno = 0
							πF.SetLineno(93)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßserialno, πTemp002); πE != nil {
								continue
							}
							// line 94: """Internal serial number to keep track of the add order of
							πF.SetLineno(94)
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
					// line 97: def add_transform(self, transform_class, priority=None, **kwargs):
					πF.SetLineno(97)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "transform_class", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "priority", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("add_transform", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtransform_class *πg.Object = πArgs[1]
						_ = µtransform_class
						var µpriority *πg.Object = πArgs[2]
						_ = µpriority
						var µkwargs *πg.Object = πArgs[3]
						_ = µkwargs
						var µpriority_string *πg.Object = πg.UnboundLocal
						_ = µpriority_string
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
							// line 98: """
							πF.SetLineno(98)
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µpriority == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 104: if priority is None:
							πF.SetLineno(104)
						Label1:
							// line 105: priority = transform_class.default_priority
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtransform_class, ßdefault_priority, nil); πE != nil {
								continue
							}
							µpriority = πTemp001
							goto Label2
						Label2:
							// line 106: priority_string = self.get_priority_string(priority)
							πF.SetLineno(106)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp004[0] = µpriority
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_priority_string, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µpriority_string = πTemp002
							// line 107: self.transforms.append(
							πF.SetLineno(107)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpriority_string, "priority_string"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(µpriority_string, µtransform_class, πTemp002, µkwargs).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 109: self.sorted = 0
							πF.SetLineno(109)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsorted, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_transform.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 98: """
					πF.SetLineno(98)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Store a single transform.  Use `priority` to override the default.\n        `kwargs` is a dictionary whose contents are passed as keyword\n        arguments to the `apply` method of the transform.  This can be used to\n        pass application-specific data to the transform instance.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßadd_transform); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 111: def add_transforms(self, transform_list):
					πF.SetLineno(111)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "transform_list", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("add_transforms", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtransform_list *πg.Object = πArgs[1]
						_ = µtransform_list
						var µtransform_class *πg.Object = πg.UnboundLocal
						_ = µtransform_class
						var µpriority_string *πg.Object = πg.UnboundLocal
						_ = µpriority_string
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
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
							// line 112: """Store multiple transforms, with default priorities."""
							πF.SetLineno(112)
							if πE = πg.CheckLocal(πF, µtransform_list, "transform_list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtransform_list); πE != nil {
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
								µtransform_class = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 114: priority_string = self.get_priority_string(
							πF.SetLineno(114)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtransform_class, ßdefault_priority, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget_priority_string, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpriority_string = πTemp006
							// line 116: self.transforms.append(
							πF.SetLineno(116)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpriority_string, "priority_string"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007 = πg.NewDict()
							πTemp008 = πTemp007.ToObject()
							πTemp004 = πg.NewTuple4(µpriority_string, µtransform_class, πTemp006, πTemp008).ToObject()
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 118: self.sorted = 0
							πF.SetLineno(118)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsorted, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_transforms.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 112: """Store multiple transforms, with default priorities."""
					πF.SetLineno(112)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Store multiple transforms, with default priorities.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßadd_transforms); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 120: def add_pending(self, pending, priority=None):
					πF.SetLineno(120)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pending", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "priority", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("add_pending", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πArgs[1]
						_ = µpending
						var µpriority *πg.Object = πArgs[2]
						_ = µpriority
						var µtransform_class *πg.Object = πg.UnboundLocal
						_ = µtransform_class
						var µpriority_string *πg.Object = πg.UnboundLocal
						_ = µpriority_string
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
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
							// line 121: """Store a transform with an associated `pending` node."""
							πF.SetLineno(121)
							// line 122: transform_class = pending.transform
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßtransform, nil); πE != nil {
								continue
							}
							µtransform_class = πTemp001
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µpriority == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 123: if priority is None:
							πF.SetLineno(123)
						Label1:
							// line 124: priority = transform_class.default_priority
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtransform_class, ßdefault_priority, nil); πE != nil {
								continue
							}
							µpriority = πTemp001
							goto Label2
						Label2:
							// line 125: priority_string = self.get_priority_string(priority)
							πF.SetLineno(125)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							πTemp004[0] = µpriority
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_priority_string, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µpriority_string = πTemp002
							// line 126: self.transforms.append(
							πF.SetLineno(126)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpriority_string, "priority_string"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp005 = πg.NewDict()
							πTemp002 = πTemp005.ToObject()
							πTemp001 = πg.NewTuple4(µpriority_string, µtransform_class, µpending, πTemp002).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 128: self.sorted = 0
							πF.SetLineno(128)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsorted, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_pending.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 121: """Store a transform with an associated `pending` node."""
					πF.SetLineno(121)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Store a transform with an associated `pending` node.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßadd_pending); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 130: def get_priority_string(self, priority):
					πF.SetLineno(130)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "priority", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("get_priority_string", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpriority *πg.Object = πArgs[1]
						_ = µpriority
						var πTemp001 *πg.Object
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
							// line 131: """
							πF.SetLineno(131)
							// line 136: self.serialno += 1
							πF.SetLineno(136)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßserialno, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßserialno, πTemp002); πE != nil {
								continue
							}
							// line 137: return '%03d-%03d' % (priority, self.serialno)
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßserialno, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µpriority, πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%03d-%03d").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_priority_string.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 131: """
					πF.SetLineno(131)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Return a string, `priority` combined with `self.serialno`.\n\n        This ensures FIFO order on transforms with identical priority.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßget_priority_string); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 139: def populate_from_components(self, components):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "components", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("populate_from_components", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcomponents *πg.Object = πArgs[1]
						_ = µcomponents
						var µcomponent *πg.Object = πg.UnboundLocal
						_ = µcomponent
						var µunknown_reference_resolvers *πg.Object = πg.UnboundLocal
						_ = µunknown_reference_resolvers
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µdecorated_list *πg.Object = πg.UnboundLocal
						_ = µdecorated_list
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []πg.Param
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 140: """
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µcomponents); πE != nil {
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
								µcomponent = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µcomponent == πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 145: if component is None:
							πF.SetLineno(145)
						Label4:
							// line 146: continue
							πF.SetLineno(146)
							continue
							goto Label5
						Label5:
							// line 147: self.add_transforms(component.get_transforms())
							πF.SetLineno(147)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcomponent, ßget_transforms, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßadd_transforms, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 148: self.components[component.component_type] = component
							πF.SetLineno(148)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µcomponent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcomponents, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µcomponent, ßcomponent_type, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.SetItem(πF, πTemp005, πTemp007, πTemp004); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 149: self.sorted = 0
							πF.SetLineno(149)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsorted, πTemp001); πE != nil {
								continue
							}
							// line 153: unknown_reference_resolvers = []
							πF.SetLineno(153)
							πTemp006 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µunknown_reference_resolvers = πTemp001
							if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µcomponents); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp002 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label8
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
								µi = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 155: unknown_reference_resolvers.extend(i.unknown_reference_resolvers)
							πF.SetLineno(155)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µi, ßunknown_reference_resolvers, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µunknown_reference_resolvers, "unknown_reference_resolvers"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µunknown_reference_resolvers, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 156: decorated_list = sorted((f.priority, f) for f in unknown_reference_resolvers)
							πF.SetLineno(156)
							πTemp006 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µunknown_reference_resolvers, "unknown_reference_resolvers"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µunknown_reference_resolvers); πE != nil {
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
										// line 156: decorated_list = sorted((f.priority, f) for f in unknown_reference_resolvers)
										πF.SetLineno(156)
										if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetAttr(πF, µf, ßpriority, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
											continue
										}
										πTemp004 = πg.NewTuple2(πTemp005, µf).ToObject()
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
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µdecorated_list = πTemp005
							// line 157: self.unknown_reference_resolvers.extend(f[1] for f in decorated_list)
							πF.SetLineno(157)
							πTemp006 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µdecorated_list, "decorated_list"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µdecorated_list); πE != nil {
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
										// line 157: self.unknown_reference_resolvers.extend(f[1] for f in decorated_list)
										πF.SetLineno(157)
										πTemp004 = πg.NewInt(1).ToObject()
										if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µf, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
									Label4:
										πTemp004 = πSent
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
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßunknown_reference_resolvers, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpopulate_from_components.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 140: """
					πF.SetLineno(140)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Store each component's default transforms, with default priorities.\n        Also, store components by type name in a mapping for later lookup.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßpopulate_from_components); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 159: def apply_transforms(self):
					πF.SetLineno(159)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("apply_transforms", "/usr/lib/python2.7/site-packages/docutils/transforms/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpriority *πg.Object = πg.UnboundLocal
						_ = µpriority
						var µtransform_class *πg.Object = πg.UnboundLocal
						_ = µtransform_class
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var µkwargs *πg.Object = πg.UnboundLocal
						_ = µkwargs
						var µtransform *πg.Object = πg.UnboundLocal
						_ = µtransform
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
						var πTemp009 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 160: """Apply all of the stored transforms, in priority order."""
							πF.SetLineno(160)
							// line 161: self.document.reporter.attach_observer(
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_transform_message, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßattach_observer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 163: while self.transforms:
							πF.SetLineno(163)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsorted, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 164: if not self.sorted:
							πF.SetLineno(164)
						Label4:
							// line 166: self.transforms.sort()
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsort, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 167: self.transforms.reverse()
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 168: self.sorted = 1
							πF.SetLineno(168)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsorted, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 169: priority, transform_class, pending, kwargs = self.transforms.pop()
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µpriority = πTemp003
							µtransform_class = πTemp006
							µpending = πTemp007
							µkwargs = πTemp008
							// line 170: transform = transform_class(self.document, startnode=pending)
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"startnode", µpending},
							}
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πTemp002, πE = µtransform_class.Call(πF, πTemp001, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtransform = πTemp002
							// line 171: transform.apply(**kwargs)
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransform, "transform"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtransform, ßapply, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, nil, nil, µkwargs); πE != nil {
								continue
							}
							// line 172: self.applied.append((priority, transform_class, pending, kwargs))
							πF.SetLineno(172)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpriority, "priority"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransform_class, "transform_class"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(µpriority, µtransform_class, µpending, µkwargs).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßapplied, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßapply_transforms.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 160: """Apply all of the stored transforms, in priority order."""
					πF.SetLineno(160)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Apply all of the stored transforms, in priority order.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßapply_transforms); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Transformer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformer.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("transforms", Code)
}
