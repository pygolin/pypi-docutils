package references

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAnonymousHyperlinks := πg.InternStr("AnonymousHyperlinks")
		ßCircularSubstitutionDefinitionError := πg.InternStr("CircularSubstitutionDefinitionError")
		ßDanglingReferences := πg.InternStr("DanglingReferences")
		ßDanglingReferencesVisitor := πg.InternStr("DanglingReferencesVisitor")
		ßException := πg.InternStr("Exception")
		ßExternalTargets := πg.InternStr("ExternalTargets")
		ßFootnotes := πg.InternStr("Footnotes")
		ßIndexError := πg.InternStr("IndexError")
		ßIndirectHyperlinks := πg.InternStr("IndirectHyperlinks")
		ßInternalTargets := πg.InternStr("InternalTargets")
		ßInvisible := πg.InternStr("Invisible")
		ßNone := πg.InternStr("None")
		ßPropagateTargets := πg.InternStr("PropagateTargets")
		ßReferential := πg.InternStr("Referential")
		ßSparseNodeVisitor := πg.InternStr("SparseNodeVisitor")
		ßSubstitutions := πg.InternStr("Substitutions")
		ßTargetNotes := πg.InternStr("TargetNotes")
		ßTargetable := πg.InternStr("Targetable")
		ßText := πg.InternStr("Text")
		ßTextElement := πg.InternStr("TextElement")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd_backref := πg.InternStr("add_backref")
		ßanonymous := πg.InternStr("anonymous")
		ßappend := πg.InternStr("append")
		ßapply := πg.InternStr("apply")
		ßattributes := πg.InternStr("attributes")
		ßauto := πg.InternStr("auto")
		ßautofootnote_labels := πg.InternStr("autofootnote_labels")
		ßautofootnote_refs := πg.InternStr("autofootnote_refs")
		ßautofootnote_start := πg.InternStr("autofootnote_start")
		ßautofootnotes := πg.InternStr("autofootnotes")
		ßchildren := πg.InternStr("children")
		ßcircular_indirect_reference := πg.InternStr("circular_indirect_reference")
		ßcitation_refs := πg.InternStr("citation_refs")
		ßcitations := πg.InternStr("citations")
		ßclass := πg.InternStr("class")
		ßclasses := πg.InternStr("classes")
		ßdeepcopy := πg.InternStr("deepcopy")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdelattr := πg.InternStr("delattr")
		ßdetails := πg.InternStr("details")
		ßdivmod := πg.InternStr("divmod")
		ßdocument := πg.InternStr("document")
		ßdupnames := πg.InternStr("dupnames")
		ßerror := πg.InternStr("error")
		ßexpect_referenced_by_id := πg.InternStr("expect_referenced_by_id")
		ßexpect_referenced_by_name := πg.InternStr("expect_referenced_by_name")
		ßextend := πg.InternStr("extend")
		ßfootnote := πg.InternStr("footnote")
		ßfootnote_reference := πg.InternStr("footnote_reference")
		ßfootnote_refs := πg.InternStr("footnote_refs")
		ßfootnotes := πg.InternStr("footnotes")
		ßget := πg.InternStr("get")
		ßget_trim_footnote_ref_space := πg.InternStr("get_trim_footnote_ref_space")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßids := πg.InternStr("ids")
		ßindex := πg.InternStr("index")
		ßindirect_target_error := πg.InternStr("indirect_target_error")
		ßindirect_targets := πg.InternStr("indirect_targets")
		ßinfo := πg.InternStr("info")
		ßinline := πg.InternStr("inline")
		ßinsert := πg.InternStr("insert")
		ßisinstance := πg.InternStr("isinstance")
		ßlabel := πg.InternStr("label")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßlist := πg.InternStr("list")
		ßliteral_block := πg.InternStr("literal_block")
		ßlower := πg.InternStr("lower")
		ßlstrip := πg.InternStr("lstrip")
		ßltrim := πg.InternStr("ltrim")
		ßmake_target_footnote := πg.InternStr("make_target_footnote")
		ßmultiply_indirect := πg.InternStr("multiply_indirect")
		ßnameids := πg.InternStr("nameids")
		ßnames := πg.InternStr("names")
		ßnext_node := πg.InternStr("next_node")
		ßnodes := πg.InternStr("nodes")
		ßnonexistent_indirect_target := πg.InternStr("nonexistent_indirect_target")
		ßnote_autofootnote := πg.InternStr("note_autofootnote")
		ßnote_autofootnote_ref := πg.InternStr("note_autofootnote_ref")
		ßnote_explicit_target := πg.InternStr("note_explicit_target")
		ßnote_footnote_ref := πg.InternStr("note_footnote_ref")
		ßnote_referenced_by := πg.InternStr("note_referenced_by")
		ßnote_refid := πg.InternStr("note_refid")
		ßnote_refname := πg.InternStr("note_refname")
		ßnumber_footnote_references := πg.InternStr("number_footnote_references")
		ßnumber_footnotes := πg.InternStr("number_footnotes")
		ßparagraph := πg.InternStr("paragraph")
		ßparent := πg.InternStr("parent")
		ßproblematic := πg.InternStr("problematic")
		ßrawsource := πg.InternStr("rawsource")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßreferenced := πg.InternStr("referenced")
		ßrefid := πg.InternStr("refid")
		ßrefids := πg.InternStr("refids")
		ßrefname := πg.InternStr("refname")
		ßrefnames := πg.InternStr("refnames")
		ßrefuri := πg.InternStr("refuri")
		ßreplace_self := πg.InternStr("replace_self")
		ßreporter := πg.InternStr("reporter")
		ßresolve_footnotes_and_citations := πg.InternStr("resolve_footnotes_and_citations")
		ßresolve_indirect_references := πg.InternStr("resolve_indirect_references")
		ßresolve_indirect_target := πg.InternStr("resolve_indirect_target")
		ßresolve_reference_ids := πg.InternStr("resolve_reference_ids")
		ßresolve_references := πg.InternStr("resolve_references")
		ßresolved := πg.InternStr("resolved")
		ßrstrip := πg.InternStr("rstrip")
		ßrtrim := πg.InternStr("rtrim")
		ßset_id := πg.InternStr("set_id")
		ßsetdefault := πg.InternStr("setdefault")
		ßsettings := πg.InternStr("settings")
		ßstartnode := πg.InternStr("startnode")
		ßstr := πg.InternStr("str")
		ßsubstitution_definition := πg.InternStr("substitution_definition")
		ßsubstitution_defs := πg.InternStr("substitution_defs")
		ßsubstitution_names := πg.InternStr("substitution_names")
		ßsubstitution_reference := πg.InternStr("substitution_reference")
		ßsymbol_footnote_refs := πg.InternStr("symbol_footnote_refs")
		ßsymbol_footnote_start := πg.InternStr("symbol_footnote_start")
		ßsymbol_footnotes := πg.InternStr("symbol_footnotes")
		ßsymbolize_footnotes := πg.InternStr("symbolize_footnotes")
		ßsymbols := πg.InternStr("symbols")
		ßsys := πg.InternStr("sys")
		ßtarget := πg.InternStr("target")
		ßtransformer := πg.InternStr("transformer")
		ßtraverse := πg.InternStr("traverse")
		ßtrim := πg.InternStr("trim")
		ßuniq := πg.InternStr("uniq")
		ßunknown_reference_resolvers := πg.InternStr("unknown_reference_resolvers")
		ßunknown_visit := πg.InternStr("unknown_visit")
		ßupdate := πg.InternStr("update")
		ßutils := πg.InternStr("utils")
		ßvisit_citation_reference := πg.InternStr("visit_citation_reference")
		ßvisit_footnote_reference := πg.InternStr("visit_footnote_reference")
		ßvisit_reference := πg.InternStr("visit_reference")
		ßwalk := πg.InternStr("walk")
		ßzip := πg.InternStr("zip")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nTransforms for resolving references.\n").ToObject()); πE != nil {
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
			// line 12: import re
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from docutils import nodes, utils
			πF.SetLineno(13)
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
			// line 14: from docutils.transforms import TransformError, Transform
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransform); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransform.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: class PropagateTargets(Transform):
			πF.SetLineno(17)
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
			_, πE = πg.NewCode("PropagateTargets", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 19: """
					πF.SetLineno(19)
					// line 19: """
					πF.SetLineno(19)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Propagate empty internal targets to the next element.\n\n    Given the following nodes::\n\n        <target ids=\"internal1\" names=\"internal1\">\n        <target anonymous=\"1\" ids=\"id1\">\n        <target ids=\"internal2\" names=\"internal2\">\n        <paragraph>\n            This is a test.\n\n    PropagateTargets propagates the ids and names of the internal\n    targets preceding the paragraph to the paragraph itself::\n\n        <target refid=\"internal1\">\n        <target anonymous=\"1\" refid=\"id1\">\n        <target refid=\"internal2\">\n        <paragraph ids=\"internal2 id1 internal1\" names=\"internal2 internal1\">\n            This is a test.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 40: default_priority = 260
					πF.SetLineno(40)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(260).ToObject()); πE != nil {
						continue
					}
					// line 42: def apply(self):
					πF.SetLineno(42)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
						var µnext_node *πg.Object = πg.UnboundLocal
						_ = µnext_node
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
						var πTemp015 *πg.Dict
						_ = πTemp015
						var πTemp016 []*πg.Object
						_ = πTemp016
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
							case 17:
								goto Label17
							case 18:
								goto Label18
							case 20:
								goto Label20
							case 21:
								goto Label21
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
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
								µtarget = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtarget, ßparent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp007
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp009
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp009
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp009
						Label5:
							πTemp003 = πTemp004
						Label4:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 45: if (isinstance(target.parent, nodes.TextElement) or
							πF.SetLineno(45)
						Label6:
							// line 48: continue
							πF.SetLineno(48)
							continue
							goto Label7
						Label7:
							// line 49: assert len(target) == 0, 'error: block-level target has children'
							πF.SetLineno(49)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp002[0] = µtarget
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp003, πg.NewStr("error: block-level target has children").ToObject()); πE != nil {
								continue
							}
							// line 50: next_node = target.next_node(ascend=True)
							πF.SetLineno(50)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"ascend", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßnext_node, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							µnext_node = πTemp004
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µnext_node != πTemp007).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label8
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							πTemp002[0] = µnext_node
							if πTemp012, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßInvisible, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp013
							if πTemp012, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp014, πE = πg.IsTrue(πF, πTemp013); πE != nil {
								continue
							}
							πTemp009 = πg.GetBool(!πTemp014).ToObject()
							πTemp007 = πTemp009
							if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp011 {
								goto Label10
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							πTemp002[0] = µnext_node
							if πTemp012, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßTargetable, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp013
							if πTemp012, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp014, πE = πg.IsTrue(πF, πTemp013); πE != nil {
								continue
							}
							πTemp009 = πg.GetBool(!πTemp014).ToObject()
							πTemp007 = πTemp009
						Label10:
							πTemp004 = πTemp007
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							πTemp002[0] = µnext_node
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßtarget, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp009
							if πTemp007, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp009
						Label9:
							πTemp003 = πTemp004
						Label8:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							goto Label12
							// line 53: if (next_node is not None and
							πF.SetLineno(53)
						Label11:
							// line 57: next_node['ids'].extend(target['ids'])
							πF.SetLineno(57)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnext_node, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßextend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 58: next_node['names'].extend(target['names'])
							πF.SetLineno(58)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnext_node, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßextend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							πTemp002[0] = µnext_node
							πTemp002[1] = ßexpect_referenced_by_name.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label13
							}
							goto Label14
							// line 60: if not hasattr(next_node, 'expect_referenced_by_name'):
							πF.SetLineno(60)
						Label13:
							// line 61: next_node.expect_referenced_by_name = {}
							πF.SetLineno(61)
							πTemp015 = πg.NewDict()
							πTemp003 = πTemp015.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnext_node, ßexpect_referenced_by_name, πTemp004); πE != nil {
								continue
							}
							goto Label14
						Label14:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							πTemp002[0] = µnext_node
							πTemp002[1] = ßexpect_referenced_by_id.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label15
							}
							goto Label16
							// line 62: if not hasattr(next_node, 'expect_referenced_by_id'):
							πF.SetLineno(62)
						Label15:
							// line 63: next_node.expect_referenced_by_id = {}
							πF.SetLineno(63)
							πTemp015 = πg.NewDict()
							πTemp003 = πTemp015.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnext_node, ßexpect_referenced_by_id, πTemp004); πE != nil {
								continue
							}
							goto Label16
						Label16:
							πTemp004 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(18)
							πTemp006 = false
						Label17:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label19
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µid = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(17)
							// line 66: self.document.ids[id] = next_node
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µnext_node); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßids, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp007 = µid
							if πE = πg.SetItem(πF, πTemp009, πTemp007, πTemp004); πE != nil {
								continue
							}
							// line 69: next_node.expect_referenced_by_id[id] = target
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µtarget); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnext_node, ßexpect_referenced_by_id, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp009 = µid
							if πE = πg.SetItem(πF, πTemp007, πTemp009, πTemp004); πE != nil {
								continue
							}
							continue
						Label18:
							if πE != nil || πR != nil {
								continue
							}
						Label19:
							πTemp004 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(21)
							πTemp006 = false
						Label20:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label22
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µname = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(20)
							// line 71: next_node.expect_referenced_by_name[name] = target
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µtarget); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µnext_node, ßexpect_referenced_by_name, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009 = µname
							if πE = πg.SetItem(πF, πTemp007, πTemp009, πTemp004); πE != nil {
								continue
							}
							continue
						Label21:
							if πE != nil || πR != nil {
								continue
							}
						Label22:
							// line 74: next_node.expect_referenced_by_name.update(
							πF.SetLineno(74)
							πTemp002 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp016[0] = µtarget
							πTemp016[1] = ßexpect_referenced_by_name.ToObject()
							πTemp015 = πg.NewDict()
							πTemp003 = πTemp015.ToObject()
							πTemp016[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp016, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnext_node, ßexpect_referenced_by_name, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 76: next_node.expect_referenced_by_id.update(
							πF.SetLineno(76)
							πTemp002 = πF.MakeArgs(1)
							πTemp016 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp016[0] = µtarget
							πTemp016[1] = ßexpect_referenced_by_id.ToObject()
							πTemp015 = πg.NewDict()
							πTemp003 = πTemp015.ToObject()
							πTemp016[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp016, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp016)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µnext_node, "next_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnext_node, ßexpect_referenced_by_id, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 80: target['refid'] = target['ids'][0]
							πF.SetLineno(80)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp007 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µtarget, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp007 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µtarget, πTemp007, πTemp003); πE != nil {
								continue
							}
							// line 83: target['ids'] = []
							πF.SetLineno(83)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp007 = ßids.ToObject()
							if πE = πg.SetItem(πF, µtarget, πTemp007, πTemp004); πE != nil {
								continue
							}
							// line 84: target['names'] = []
							πF.SetLineno(84)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp007 = ßnames.ToObject()
							if πE = πg.SetItem(πF, µtarget, πTemp007, πTemp004); πE != nil {
								continue
							}
							// line 85: self.document.note_refid(target)
							πF.SetLineno(85)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp002[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label12
						Label12:
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PropagateTargets").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPropagateTargets.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 88: class AnonymousHyperlinks(Transform):
			πF.SetLineno(88)
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
			_, πE = πg.NewCode("AnonymousHyperlinks", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 90: """
					πF.SetLineno(90)
					// line 90: """
					πF.SetLineno(90)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Link anonymous references to targets.  Given::\n\n        <paragraph>\n            <reference anonymous=\"1\">\n                internal\n            <reference anonymous=\"1\">\n                external\n        <target anonymous=\"1\" ids=\"id1\">\n        <target anonymous=\"1\" ids=\"id2\" refuri=\"http://external\">\n\n    Corresponding references are linked via \"refid\" or resolved via \"refuri\"::\n\n        <paragraph>\n            <reference anonymous=\"1\" refid=\"id1\">\n                text\n            <reference anonymous=\"1\" refuri=\"http://external\">\n                external\n        <target anonymous=\"1\" ids=\"id1\">\n        <target anonymous=\"1\" ids=\"id2\" refuri=\"http://external\">\n    ").ToObject()); πE != nil {
						continue
					}
					// line 112: default_priority = 440
					πF.SetLineno(112)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(440).ToObject()); πE != nil {
						continue
					}
					// line 114: def apply(self):
					πF.SetLineno(114)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µanonymous_refs *πg.Object = πg.UnboundLocal
						_ = µanonymous_refs
						var µanonymous_targets *πg.Object = πg.UnboundLocal
						_ = µanonymous_targets
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
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
						var πTemp008 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 6:
								goto Label6
							case 7:
								goto Label7
							case 13:
								goto Label13
							case 14:
								goto Label14
							case 16:
								goto Label16
							case 17:
								goto Label17
							case 19:
								goto Label19
							case 20:
								goto Label20
							default:
								panic("unexpected function state")
							}
							// line 115: anonymous_refs = []
							πF.SetLineno(115)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µanonymous_refs = πTemp002
							// line 116: anonymous_targets = []
							πF.SetLineno(116)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µanonymous_targets = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreference, nil); πE != nil {
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßanonymous.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 118: if node.get('anonymous'):
							πF.SetLineno(118)
						Label4:
							// line 119: anonymous_refs.append(node)
							πF.SetLineno(119)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µanonymous_refs, "anonymous_refs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µanonymous_refs, ßappend, nil); πE != nil {
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
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
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
							πF.PushCheckpoint(7)
							πTemp005 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µnode = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßanonymous.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 121: if node.get('anonymous'):
							πF.SetLineno(121)
						Label9:
							// line 122: anonymous_targets.append(node)
							πF.SetLineno(122)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πE = πg.CheckLocal(πF, µanonymous_targets, "anonymous_targets"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µanonymous_targets, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µanonymous_refs, "anonymous_refs"); πE != nil {
								continue
							}
							πTemp001[0] = µanonymous_refs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µanonymous_targets, "anonymous_targets"); πE != nil {
								continue
							}
							πTemp001[0] = µanonymous_targets
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 123: if len(anonymous_refs) \
							πF.SetLineno(123)
						Label11:
							// line 125: msg = self.document.reporter.error(
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µanonymous_refs, "anonymous_refs"); πE != nil {
								continue
							}
							πTemp008[0] = µanonymous_refs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µanonymous_targets, "anonymous_targets"); πE != nil {
								continue
							}
							πTemp008[0] = µanonymous_targets
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp003 = πg.NewTuple2(πTemp007, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Anonymous hyperlink mismatch: %s references but %s targets.\nSee \"backrefs\" attribute for IDs.").ToObject(), πTemp003); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp003
							// line 129: msgid = self.document.set_id(msg)
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsgid = πTemp002
							if πE = πg.CheckLocal(πF, µanonymous_refs, "anonymous_refs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µanonymous_refs); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp005 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label15
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
								µref = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(13)
							// line 131: prb = nodes.problematic(
							πF.SetLineno(131)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µprb = πTemp003
							// line 133: prbid = self.document.set_id(prb)
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp001[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µprbid = πTemp003
							// line 134: msg.add_backref(prbid)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp001[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 135: ref.replace_self(prb)
							πF.SetLineno(135)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp001[0] = µprb
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 136: return
							πF.SetLineno(136)
							πR = πg.None
							continue
							goto Label12
						Label12:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µanonymous_refs, "anonymous_refs"); πE != nil {
								continue
							}
							πTemp001[0] = µanonymous_refs
							if πE = πg.CheckLocal(πF, µanonymous_targets, "anonymous_targets"); πE != nil {
								continue
							}
							πTemp001[1] = µanonymous_targets
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(17)
							πTemp005 = false
						Label16:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label18
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µref = πTemp004
								µtarget = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(16)
							// line 138: target.referenced = 1
							πF.SetLineno(138)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtarget, ßreferenced, πTemp003); πE != nil {
								continue
							}
							// line 139: while True:
							πF.SetLineno(139)
							πF.PushCheckpoint(20)
							πTemp006 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label21
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(19)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp011, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label22
							}
							goto Label23
							// line 140: if target.hasattr('refuri'):
							πF.SetLineno(140)
						Label22:
							// line 141: ref['refuri'] = target['refuri']
							πF.SetLineno(141)
							πTemp003 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp007 = ßrefuri.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp007, πTemp003); πE != nil {
								continue
							}
							// line 142: ref.resolved = 1
							πF.SetLineno(142)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp003); πE != nil {
								continue
							}
							// line 143: break
							πF.SetLineno(143)
							πTemp006 = true
							continue
							goto Label24
						Label23:
							πTemp004 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp011).ToObject()
							if πTemp011, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label25
							}
							goto Label26
							// line 145: if not target['ids']:
							πF.SetLineno(145)
						Label25:
							// line 147: target = self.document.ids[target['refid']]
							πF.SetLineno(147)
							πTemp004 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßids, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							µtarget = πTemp004
							// line 148: continue
							πF.SetLineno(148)
							continue
							goto Label26
						Label26:
							// line 149: ref['refid'] = target['ids'][0]
							πF.SetLineno(149)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp007 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µtarget, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp007 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp007, πTemp003); πE != nil {
								continue
							}
							// line 150: self.document.note_refid(ref)
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 151: break
							πF.SetLineno(151)
							πTemp006 = true
							continue
							goto Label24
						Label24:
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
							continue
						Label17:
							if πE != nil || πR != nil {
								continue
							}
						Label18:
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("AnonymousHyperlinks").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAnonymousHyperlinks.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 154: class IndirectHyperlinks(Transform):
			πF.SetLineno(154)
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
			_, πE = πg.NewCode("IndirectHyperlinks", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 156: """
					πF.SetLineno(156)
					// line 156: """
					πF.SetLineno(156)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    a) Indirect external references::\n\n           <paragraph>\n               <reference refname=\"indirect external\">\n                   indirect external\n           <target id=\"id1\" name=\"direct external\"\n               refuri=\"http://indirect\">\n           <target id=\"id2\" name=\"indirect external\"\n               refname=\"direct external\">\n\n       The \"refuri\" attribute is migrated back to all indirect targets\n       from the final direct target (i.e. a target not referring to\n       another indirect target)::\n\n           <paragraph>\n               <reference refname=\"indirect external\">\n                   indirect external\n           <target id=\"id1\" name=\"direct external\"\n               refuri=\"http://indirect\">\n           <target id=\"id2\" name=\"indirect external\"\n               refuri=\"http://indirect\">\n\n       Once the attribute is migrated, the preexisting \"refname\" attribute\n       is dropped.\n\n    b) Indirect internal references::\n\n           <target id=\"id1\" name=\"final target\">\n           <paragraph>\n               <reference refname=\"indirect internal\">\n                   indirect internal\n           <target id=\"id2\" name=\"indirect internal 2\"\n               refname=\"final target\">\n           <target id=\"id3\" name=\"indirect internal\"\n               refname=\"indirect internal 2\">\n\n       Targets which indirectly refer to an internal target become one-hop\n       indirect (their \"refid\" attributes are directly set to the internal\n       target's \"id\"). References which indirectly refer to an internal\n       target become direct internal references::\n\n           <target id=\"id1\" name=\"final target\">\n           <paragraph>\n               <reference refid=\"id1\">\n                   indirect internal\n           <target id=\"id2\" name=\"indirect internal 2\" refid=\"id1\">\n           <target id=\"id3\" name=\"indirect internal\" refid=\"id1\">\n    ").ToObject()); πE != nil {
						continue
					}
					// line 206: default_priority = 460
					πF.SetLineno(206)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(460).ToObject()); πE != nil {
						continue
					}
					// line 208: def apply(self):
					πF.SetLineno(208)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
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
						var πTemp006 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindirect_targets, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µtarget = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßresolved, nil); πE != nil {
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
							// line 210: if not target.resolved:
							πF.SetLineno(210)
						Label4:
							// line 211: self.resolve_indirect_target(target)
							πF.SetLineno(211)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp006[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßresolve_indirect_target, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label5
						Label5:
							// line 212: self.resolve_indirect_references(target)
							πF.SetLineno(212)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp006[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßresolve_indirect_references, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					// line 214: def resolve_indirect_target(self, target):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("resolve_indirect_target", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var µrefname *πg.Object = πg.UnboundLocal
						_ = µrefname
						var µreftarget_id *πg.Object = πg.UnboundLocal
						_ = µreftarget_id
						var µresolver_function *πg.Object = πg.UnboundLocal
						_ = µresolver_function
						var µreftarget *πg.Object = πg.UnboundLocal
						_ = µreftarget
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 215: refname = target.get('refname')
							πF.SetLineno(215)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtarget, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrefname = πTemp003
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µrefname == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 216: if refname is None:
							πF.SetLineno(216)
						Label1:
							// line 217: reftarget_id = target['refid']
							πF.SetLineno(217)
							πTemp002 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							µreftarget_id = πTemp003
							goto Label3
						Label2:
							// line 219: reftarget_id = self.document.nameids.get(refname)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							πTemp001[0] = µrefname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreftarget_id = πTemp003
							if πE = πg.CheckLocal(πF, µreftarget_id, "reftarget_id"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µreftarget_id); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 220: if not reftarget_id:
							πF.SetLineno(220)
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßtransformer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßunknown_reference_resolvers, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µresolver_function = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µresolver_function, "resolver_function"); πE != nil {
								continue
							}
							if πTemp003, πE = µresolver_function.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 224: if resolver_function(target):
							πF.SetLineno(224)
						Label9:
							// line 225: break
							πF.SetLineno(225)
							πTemp004 = true
							continue
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							// line 227: self.nonexistent_indirect_target(target)
							πF.SetLineno(227)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnonexistent_indirect_target, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						Label8:
							// line 228: return
							πF.SetLineno(228)
							πR = πg.None
							continue
							goto Label5
						Label5:
							goto Label3
						Label3:
							// line 229: reftarget = self.document.ids[reftarget_id]
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µreftarget_id, "reftarget_id"); πE != nil {
								continue
							}
							πTemp002 = µreftarget_id
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µreftarget = πTemp003
							// line 230: reftarget.note_referenced_by(id=reftarget_id)
							πF.SetLineno(230)
							if πE = πg.CheckLocal(πF, µreftarget_id, "reftarget_id"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"id", µreftarget_id},
							}
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µreftarget, ßnote_referenced_by, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							πTemp001[0] = µreftarget
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µreftarget, ßresolved, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label11
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µreftarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
						Label11:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 231: if isinstance(reftarget, nodes.target) \
							πF.SetLineno(231)
						Label12:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							πTemp001[1] = ßmultiply_indirect.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
								goto Label14
							}
							goto Label15
							// line 233: if hasattr(target, 'multiply_indirect'):
							πF.SetLineno(233)
						Label14:
							// line 236: self.circular_indirect_reference(target)
							πF.SetLineno(236)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcircular_indirect_reference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 237: return
							πF.SetLineno(237)
							πR = πg.None
							continue
							goto Label15
						Label15:
							// line 238: target.multiply_indirect = 1
							πF.SetLineno(238)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtarget, ßmultiply_indirect, πTemp002); πE != nil {
								continue
							}
							// line 239: self.resolve_indirect_target(reftarget) # multiply indirect
							πF.SetLineno(239)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							πTemp001[0] = µreftarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßresolve_indirect_target, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 240: del target.multiply_indirect
							πF.SetLineno(240)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µtarget, ßmultiply_indirect); πE != nil {
								continue
							}
							goto Label13
						Label13:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µreftarget, ßhasattr, nil); πE != nil {
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
								goto Label16
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µreftarget, ßhasattr, nil); πE != nil {
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
								goto Label17
							}
							πTemp002 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µreftarget, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							goto Label19
							// line 241: if reftarget.hasattr('refuri'):
							πF.SetLineno(241)
						Label16:
							// line 242: target['refuri'] = reftarget['refuri']
							πF.SetLineno(242)
							πTemp002 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µreftarget, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp005 = ßrefuri.ToObject()
							if πE = πg.SetItem(πF, µtarget, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µtarget, ßrefid.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 243: if 'refid' in target:
							πF.SetLineno(243)
						Label21:
							// line 244: del target['refid']
							πF.SetLineno(244)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp002 = ßrefid.ToObject()
							if πE = πg.DelItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							goto Label22
						Label22:
							goto Label20
							// line 245: elif reftarget.hasattr('refid'):
							πF.SetLineno(245)
						Label17:
							// line 246: target['refid'] = reftarget['refid']
							πF.SetLineno(246)
							πTemp002 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µreftarget, "reftarget"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µreftarget, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp005 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µtarget, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 247: self.document.note_refid(target)
							πF.SetLineno(247)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label20
							// line 249: if reftarget['ids']:
							πF.SetLineno(249)
						Label18:
							// line 250: target['refid'] = reftarget_id
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µreftarget_id, "reftarget_id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µreftarget_id); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp003 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µtarget, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 251: self.document.note_refid(target)
							πF.SetLineno(251)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label20
						Label19:
							// line 253: self.nonexistent_indirect_target(target)
							πF.SetLineno(253)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnonexistent_indirect_target, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 254: return
							πF.SetLineno(254)
							πR = πg.None
							continue
							goto Label20
						Label20:
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µrefname != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							goto Label24
							// line 255: if refname is not None:
							πF.SetLineno(255)
						Label23:
							// line 256: del target['refname']
							πF.SetLineno(256)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp002 = ßrefname.ToObject()
							if πE = πg.DelItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							goto Label24
						Label24:
							// line 257: target.resolved = 1
							πF.SetLineno(257)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtarget, ßresolved, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßresolve_indirect_target.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 259: def nonexistent_indirect_target(self, target):
					πF.SetLineno(259)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("nonexistent_indirect_target", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							πTemp002 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 260: if target['refname'] in self.document.nameids:
							πF.SetLineno(260)
						Label1:
							// line 261: self.indirect_target_error(target, 'which is a duplicate, and '
							πF.SetLineno(261)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp006[0] = µtarget
							πTemp006[1] = πg.NewStr("which is a duplicate, and cannot be used as a unique reference").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßindirect_target_error, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label3
						Label2:
							// line 264: self.indirect_target_error(target, 'which does not exist')
							πF.SetLineno(264)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp006[0] = µtarget
							πTemp006[1] = πg.NewStr("which does not exist").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßindirect_target_error, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					if πE = πClass.SetItem(πF, ßnonexistent_indirect_target.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 266: def circular_indirect_reference(self, target):
					πF.SetLineno(266)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("circular_indirect_reference", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
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
							// line 267: self.indirect_target_error(target, 'forming a circular reference')
							πF.SetLineno(267)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp001[0] = µtarget
							πTemp001[1] = πg.NewStr("forming a circular reference").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindirect_target_error, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcircular_indirect_reference.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 269: def indirect_target_error(self, target, explanation):
					πF.SetLineno(269)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp002[2] = πg.Param{Name: "explanation", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("indirect_target_error", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var µexplanation *πg.Object = πArgs[2]
						_ = µexplanation
						var µnaming *πg.Object = πg.UnboundLocal
						_ = µnaming
						var µreflist *πg.Object = πg.UnboundLocal
						_ = µreflist
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							case 6:
								goto Label6
							case 7:
								goto Label7
							case 11:
								goto Label11
							case 12:
								goto Label12
							default:
								panic("unexpected function state")
							}
							// line 270: naming = ''
							πF.SetLineno(270)
							µnaming = ß.ToObject()
							// line 271: reflist = []
							πF.SetLineno(271)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µreflist = πTemp002
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 272: if target['names']:
							πF.SetLineno(272)
						Label1:
							// line 273: naming = '"%s" ' % target['names'][0]
							πF.SetLineno(273)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp006 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\"%s\" ").ToObject(), πTemp005); πE != nil {
								continue
							}
							µnaming = πTemp002
							goto Label2
						Label2:
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp004 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µname = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 275: reflist.extend(self.document.refnames.get(name, []))
							πF.SetLineno(275)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009[0] = µname
							πTemp010 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp010...).ToObject()
							πTemp009[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrefnames, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µreflist, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µid = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 277: reflist.extend(self.document.refids.get(id, []))
							πF.SetLineno(277)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp009[0] = µid
							πTemp010 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp010...).ToObject()
							πTemp009[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrefids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µreflist, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							πTemp002 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 278: if target['ids']:
							πF.SetLineno(278)
						Label9:
							// line 279: naming += '(id="%s")' % target['ids'][0]
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µnaming, "naming"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp006 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("(id=\"%s\")").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µnaming, πTemp002); πE != nil {
								continue
							}
							µnaming = πTemp003
							goto Label10
						Label10:
							// line 280: msg = self.document.reporter.error(
							πF.SetLineno(280)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnaming, "naming"); πE != nil {
								continue
							}
							πTemp005 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µtarget, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexplanation, "explanation"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(µnaming, πTemp006, µexplanation).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Indirect hyperlink target %s refers to target \"%s\", %s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"base_node", µtarget},
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
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp003
							// line 283: msgid = self.document.set_id(msg)
							πF.SetLineno(283)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsgid = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							πTemp001[0] = µreflist
							if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßuniq, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp004 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µref = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(11)
							// line 285: prb = nodes.problematic(
							πF.SetLineno(285)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µprb = πTemp003
							// line 287: prbid = self.document.set_id(prb)
							πF.SetLineno(287)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp001[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µprbid = πTemp003
							// line 288: msg.add_backref(prbid)
							πF.SetLineno(288)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp001[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 289: ref.replace_self(prb)
							πF.SetLineno(289)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp001[0] = µprb
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							// line 290: target.resolved = 1
							πF.SetLineno(290)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtarget, ßresolved, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßindirect_target_error.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 292: def resolve_indirect_references(self, target):
					πF.SetLineno(292)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("resolve_indirect_references", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var µattname *πg.Object = πg.UnboundLocal
						_ = µattname
						var µcall_method *πg.Object = πg.UnboundLocal
						_ = µcall_method
						var µattval *πg.Object = πg.UnboundLocal
						_ = µattval
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µreflist *πg.Object = πg.UnboundLocal
						_ = µreflist
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
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
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πTemp009 bool
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
							case 5:
								goto Label5
							case 6:
								goto Label6
							case 10:
								goto Label10
							case 11:
								goto Label11
							case 19:
								goto Label19
							case 20:
								goto Label20
							case 24:
								goto Label24
							case 25:
								goto Label25
							default:
								panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
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
								goto Label2
							}
							goto Label3
							// line 293: if target.hasattr('refid'):
							πF.SetLineno(293)
						Label1:
							// line 294: attname = 'refid'
							πF.SetLineno(294)
							µattname = ßrefid.ToObject()
							// line 295: call_method = self.document.note_refid
							πF.SetLineno(295)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_refid, nil); πE != nil {
								continue
							}
							µcall_method = πTemp003
							goto Label4
							// line 296: elif target.hasattr('refuri'):
							πF.SetLineno(296)
						Label2:
							// line 297: attname = 'refuri'
							πF.SetLineno(297)
							µattname = ßrefuri.ToObject()
							// line 298: call_method = None
							πF.SetLineno(298)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µcall_method = πTemp002
							goto Label4
						Label3:
							// line 300: return
							πF.SetLineno(300)
							πR = πg.None
							continue
							goto Label4
						Label4:
							// line 301: attval = target[attname]
							πF.SetLineno(301)
							if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
								continue
							}
							πTemp002 = µattname
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							µattval = πTemp003
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp004 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label7
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
								µname = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(5)
							// line 303: reflist = self.document.refnames.get(name, [])
							πF.SetLineno(303)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							πTemp007 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrefnames, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreflist = πTemp005
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µreflist); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 304: if reflist:
							πF.SetLineno(304)
						Label8:
							// line 305: target.note_referenced_by(name=name)
							πF.SetLineno(305)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"name", µname},
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßnote_referenced_by, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µreflist); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp006 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µref = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(10)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label13
							}
							goto Label14
							// line 307: if ref.resolved:
							πF.SetLineno(307)
						Label13:
							// line 308: continue
							πF.SetLineno(308)
							continue
							goto Label14
						Label14:
							// line 309: del ref['refname']
							πF.SetLineno(309)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp005 = ßrefname.ToObject()
							if πE = πg.DelItem(πF, µref, πTemp005); πE != nil {
								continue
							}
							// line 310: ref[attname] = attval
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µattval, "attval"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µattval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
								continue
							}
							πTemp010 = µattname
							if πE = πg.SetItem(πF, µref, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcall_method, "call_method"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µcall_method); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label15
							}
							goto Label16
							// line 311: if call_method:
							πF.SetLineno(311)
						Label15:
							// line 312: call_method(ref)
							πF.SetLineno(312)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µcall_method, "call_method"); πE != nil {
								continue
							}
							if πTemp005, πE = µcall_method.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label16
						Label16:
							// line 313: ref.resolved = 1
							πF.SetLineno(313)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp005, ßtarget, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp010
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp009, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label17
							}
							goto Label18
							// line 314: if isinstance(ref, nodes.target):
							πF.SetLineno(314)
						Label17:
							// line 315: self.resolve_indirect_references(ref)
							πF.SetLineno(315)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßresolve_indirect_references, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label18
						Label18:
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(20)
							πTemp004 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label21
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
								µid = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(19)
							// line 317: reflist = self.document.refids.get(id, [])
							πF.SetLineno(317)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp001[0] = µid
							πTemp007 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrefids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreflist = πTemp005
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µreflist); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label22
							}
							goto Label23
							// line 318: if reflist:
							πF.SetLineno(318)
						Label22:
							// line 319: target.note_referenced_by(id=id)
							πF.SetLineno(319)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"id", µid},
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßnote_referenced_by, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							goto Label23
						Label23:
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µreflist); πE != nil {
								continue
							}
							πF.PushCheckpoint(25)
							πTemp006 = false
						Label24:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label26
							}
							if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µref = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(24)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label27
							}
							goto Label28
							// line 321: if ref.resolved:
							πF.SetLineno(321)
						Label27:
							// line 322: continue
							πF.SetLineno(322)
							continue
							goto Label28
						Label28:
							// line 323: del ref['refid']
							πF.SetLineno(323)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp005 = ßrefid.ToObject()
							if πE = πg.DelItem(πF, µref, πTemp005); πE != nil {
								continue
							}
							// line 324: ref[attname] = attval
							πF.SetLineno(324)
							if πE = πg.CheckLocal(πF, µattval, "attval"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µattval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattname, "attname"); πE != nil {
								continue
							}
							πTemp010 = µattname
							if πE = πg.SetItem(πF, µref, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcall_method, "call_method"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µcall_method); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label29
							}
							goto Label30
							// line 325: if call_method:
							πF.SetLineno(325)
						Label29:
							// line 326: call_method(ref)
							πF.SetLineno(326)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µcall_method, "call_method"); πE != nil {
								continue
							}
							if πTemp005, πE = µcall_method.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label30
						Label30:
							// line 327: ref.resolved = 1
							πF.SetLineno(327)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp005, ßtarget, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp010
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp009, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label31
							}
							goto Label32
							// line 328: if isinstance(ref, nodes.target):
							πF.SetLineno(328)
						Label31:
							// line 329: self.resolve_indirect_references(ref)
							πF.SetLineno(329)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßresolve_indirect_references, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label32
						Label32:
							continue
						Label25:
							if πE != nil || πR != nil {
								continue
							}
						Label26:
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßresolve_indirect_references.ToObject(), πTemp007); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("IndirectHyperlinks").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIndirectHyperlinks.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 332: class ExternalTargets(Transform):
			πF.SetLineno(332)
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
			_, πE = πg.NewCode("ExternalTargets", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 334: """
					πF.SetLineno(334)
					// line 334: """
					πF.SetLineno(334)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Given::\n\n        <paragraph>\n            <reference refname=\"direct external\">\n                direct external\n        <target id=\"id1\" name=\"direct external\" refuri=\"http://direct\">\n\n    The \"refname\" attribute is replaced by the direct \"refuri\" attribute::\n\n        <paragraph>\n            <reference refuri=\"http://direct\">\n                direct external\n        <target id=\"id1\" name=\"direct external\" refuri=\"http://direct\">\n    ").ToObject()); πE != nil {
						continue
					}
					// line 350: default_priority = 640
					πF.SetLineno(350)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(640).ToObject()); πE != nil {
						continue
					}
					// line 352: def apply(self):
					πF.SetLineno(352)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
						var µrefuri *πg.Object = πg.UnboundLocal
						_ = µrefuri
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µreflist *πg.Object = πg.UnboundLocal
						_ = µreflist
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πTemp011 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 6:
								goto Label6
							case 7:
								goto Label7
							case 11:
								goto Label11
							case 12:
								goto Label12
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
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
								µtarget = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 354: if target.hasattr('refuri'):
							πF.SetLineno(354)
						Label4:
							// line 355: refuri = target['refuri']
							πF.SetLineno(355)
							πTemp003 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							µrefuri = πTemp004
							πTemp004 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
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
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µname = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 357: reflist = self.document.refnames.get(name, [])
							πF.SetLineno(357)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[0] = µname
							πTemp009 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp009...).ToObject()
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßrefnames, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp007, ßget, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µreflist = πTemp007
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µreflist); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 358: if reflist:
							πF.SetLineno(358)
						Label9:
							// line 359: target.note_referenced_by(name=name)
							πF.SetLineno(359)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"name", µname},
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtarget, ßnote_referenced_by, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µreflist); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp008 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µref = πTemp007
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(11)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label14
							}
							goto Label15
							// line 361: if ref.resolved:
							πF.SetLineno(361)
						Label14:
							// line 362: continue
							πF.SetLineno(362)
							continue
							goto Label15
						Label15:
							// line 363: del ref['refname']
							πF.SetLineno(363)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp007 = ßrefname.ToObject()
							if πE = πg.DelItem(πF, µref, πTemp007); πE != nil {
								continue
							}
							// line 364: ref['refuri'] = refuri
							πF.SetLineno(364)
							if πE = πg.CheckLocal(πF, µrefuri, "refuri"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µrefuri); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp012 = ßrefuri.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp012, πTemp007); πE != nil {
								continue
							}
							// line 365: ref.resolved = 1
							πF.SetLineno(365)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp007); πE != nil {
								continue
							}
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ExternalTargets").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßExternalTargets.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 368: class InternalTargets(Transform):
			πF.SetLineno(368)
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
			_, πE = πg.NewCode("InternalTargets", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 370: default_priority = 660
					πF.SetLineno(370)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(660).ToObject()); πE != nil {
						continue
					}
					// line 372: def apply(self):
					πF.SetLineno(372)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
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
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
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
								µtarget = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp009).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp009).ToObject()
							πTemp003 = πTemp004
						Label4:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 374: if not target.hasattr('refuri') and not target.hasattr('refid'):
							πF.SetLineno(374)
						Label5:
							// line 375: self.resolve_reference_ids(target)
							πF.SetLineno(375)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp002[0] = µtarget
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßresolve_reference_ids, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label6
						Label6:
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
					// line 377: def resolve_reference_ids(self, target):
					πF.SetLineno(377)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "target", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("resolve_reference_ids", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtarget *πg.Object = πArgs[1]
						_ = µtarget
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µrefid *πg.Object = πg.UnboundLocal
						_ = µrefid
						var µreflist *πg.Object = πg.UnboundLocal
						_ = µreflist
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πTemp009 bool
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 378: """
							πF.SetLineno(378)
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtarget, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µname = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 395: refid = self.document.nameids.get(name)
							πF.SetLineno(395)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µrefid = πTemp003
							// line 396: reflist = self.document.refnames.get(name, [])
							πF.SetLineno(396)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							πTemp007 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrefnames, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µreflist = πTemp003
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µreflist); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 397: if reflist:
							πF.SetLineno(397)
						Label4:
							// line 398: target.note_referenced_by(name=name)
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"name", µname},
							}
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtarget, ßnote_referenced_by, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µreflist); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp005 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
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
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µref = πTemp003
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 400: if ref.resolved:
							πF.SetLineno(400)
						Label9:
							// line 401: continue
							πF.SetLineno(401)
							continue
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µrefid, "refid"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µrefid); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label11
							}
							goto Label12
							// line 402: if refid:
							πF.SetLineno(402)
						Label11:
							// line 403: del ref['refname']
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp003 = ßrefname.ToObject()
							if πE = πg.DelItem(πF, µref, πTemp003); πE != nil {
								continue
							}
							// line 404: ref['refid'] = refid
							πF.SetLineno(404)
							if πE = πg.CheckLocal(πF, µrefid, "refid"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µrefid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp010 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp010, πTemp003); πE != nil {
								continue
							}
							goto Label12
						Label12:
							// line 405: ref.resolved = 1
							πF.SetLineno(405)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp003); πE != nil {
								continue
							}
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
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
					if πE = πClass.SetItem(πF, ßresolve_reference_ids.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 378: """
					πF.SetLineno(378)
					// line 378: """
					πF.SetLineno(378)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n        Given::\n\n            <paragraph>\n                <reference refname=\"direct internal\">\n                    direct internal\n            <target id=\"id1\" name=\"direct internal\">\n\n        The \"refname\" attribute is replaced by \"refid\" linking to the target's\n        \"id\"::\n\n            <paragraph>\n                <reference refid=\"id1\">\n                    direct internal\n            <target id=\"id1\" name=\"direct internal\">\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Given::\n\n            <paragraph>\n                <reference refname=\"direct internal\">\n                    direct internal\n            <target id=\"id1\" name=\"direct internal\">\n\n        The \"refname\" attribute is replaced by \"refid\" linking to the target's\n        \"id\"::\n\n            <paragraph>\n                <reference refid=\"id1\">\n                    direct internal\n            <target id=\"id1\" name=\"direct internal\">\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßresolve_reference_ids); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("InternalTargets").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInternalTargets.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 408: class Footnotes(Transform):
			πF.SetLineno(408)
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
			_, πE = πg.NewCode("Footnotes", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 410: """
					πF.SetLineno(410)
					// line 410: """
					πF.SetLineno(410)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Assign numbers to autonumbered footnotes, and resolve links to footnotes,\n    citations, and their references.\n\n    Given the following ``document`` as input::\n\n        <document>\n            <paragraph>\n                A labeled autonumbered footnote referece:\n                <footnote_reference auto=\"1\" id=\"id1\" refname=\"footnote\">\n            <paragraph>\n                An unlabeled autonumbered footnote referece:\n                <footnote_reference auto=\"1\" id=\"id2\">\n            <footnote auto=\"1\" id=\"id3\">\n                <paragraph>\n                    Unlabeled autonumbered footnote.\n            <footnote auto=\"1\" id=\"footnote\" name=\"footnote\">\n                <paragraph>\n                    Labeled autonumbered footnote.\n\n    Auto-numbered footnotes have attribute ``auto=\"1\"`` and no label.\n    Auto-numbered footnote_references have no reference text (they're\n    empty elements). When resolving the numbering, a ``label`` element\n    is added to the beginning of the ``footnote``, and reference text\n    to the ``footnote_reference``.\n\n    The transformed result will be::\n\n        <document>\n            <paragraph>\n                A labeled autonumbered footnote referece:\n                <footnote_reference auto=\"1\" id=\"id1\" refid=\"footnote\">\n                    2\n            <paragraph>\n                An unlabeled autonumbered footnote referece:\n                <footnote_reference auto=\"1\" id=\"id2\" refid=\"id3\">\n                    1\n            <footnote auto=\"1\" id=\"id3\" backrefs=\"id2\">\n                <label>\n                    1\n                <paragraph>\n                    Unlabeled autonumbered footnote.\n            <footnote auto=\"1\" id=\"footnote\" name=\"footnote\" backrefs=\"id1\">\n                <label>\n                    2\n                <paragraph>\n                    Labeled autonumbered footnote.\n\n    Note that the footnotes are not in the same order as the references.\n\n    The labels and reference text are added to the auto-numbered ``footnote``\n    and ``footnote_reference`` elements.  Footnote elements are backlinked to\n    their references via \"refids\" attributes.  References are assigned \"id\"\n    and \"refid\" attributes.\n\n    After adding labels and reference text, the \"auto\" attributes can be\n    ignored.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 469: default_priority = 620
					πF.SetLineno(469)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(620).ToObject()); πE != nil {
						continue
					}
					// line 471: autofootnote_labels = None
					πF.SetLineno(471)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßautofootnote_labels.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 472: """Keep track of unlabeled autonumbered footnotes."""
					πF.SetLineno(472)
					// line 474: symbols = [
					πF.SetLineno(474)
					πTemp002 = make([]*πg.Object, 10)
					πTemp002[0] = πg.NewStr("*").ToObject()
					πTemp002[1] = πg.NewUnicode("\xe2\x80\xa0").ToObject()
					πTemp002[2] = πg.NewUnicode("\xe2\x80\xa1").ToObject()
					πTemp002[3] = πg.NewUnicode("\xc2\xa7").ToObject()
					πTemp002[4] = πg.NewUnicode("\xc2\xb6").ToObject()
					πTemp002[5] = πg.NewStr("#").ToObject()
					πTemp002[6] = πg.NewUnicode("\xe2\x99\xa0").ToObject()
					πTemp002[7] = πg.NewUnicode("\xe2\x99\xa5").ToObject()
					πTemp002[8] = πg.NewUnicode("\xe2\x99\xa6").ToObject()
					πTemp002[9] = πg.NewUnicode("\xe2\x99\xa3").ToObject()
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					if πE = πClass.SetItem(πF, ßsymbols.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 491: def apply(self):
					πF.SetLineno(491)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstartnum *πg.Object = πg.UnboundLocal
						_ = µstartnum
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
							// line 492: self.autofootnote_labels = []
							πF.SetLineno(492)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautofootnote_labels, πTemp003); πE != nil {
								continue
							}
							// line 493: startnum = self.document.autofootnote_start
							πF.SetLineno(493)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßautofootnote_start, nil); πE != nil {
								continue
							}
							µstartnum = πTemp003
							// line 494: self.document.autofootnote_start = self.number_footnotes(startnum)
							πF.SetLineno(494)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstartnum, "startnum"); πE != nil {
								continue
							}
							πTemp001[0] = µstartnum
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnumber_footnotes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßautofootnote_start, πTemp002); πE != nil {
								continue
							}
							// line 495: self.number_footnote_references(startnum)
							πF.SetLineno(495)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstartnum, "startnum"); πE != nil {
								continue
							}
							πTemp001[0] = µstartnum
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnumber_footnote_references, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 496: self.symbolize_footnotes()
							πF.SetLineno(496)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsymbolize_footnotes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 497: self.resolve_footnotes_and_citations()
							πF.SetLineno(497)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßresolve_footnotes_and_citations, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 499: def number_footnotes(self, startnum):
					πF.SetLineno(499)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "startnum", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("number_footnotes", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstartnum *πg.Object = πArgs[1]
						_ = µstartnum
						var µfootnote *πg.Object = πg.UnboundLocal
						_ = µfootnote
						var µlabel *πg.Object = πg.UnboundLocal
						_ = µlabel
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
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
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 12:
								goto Label12
							case 13:
								goto Label13
							default:
								panic("unexpected function state")
							}
							// line 500: """
							πF.SetLineno(500)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßautofootnotes, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µfootnote = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 507: while True:
							πF.SetLineno(507)
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 508: label = str(startnum)
							πF.SetLineno(508)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstartnum, "startnum"); πE != nil {
								continue
							}
							πTemp007[0] = µstartnum
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µlabel = πTemp003
							// line 509: startnum += 1
							πF.SetLineno(509)
							if πE = πg.CheckLocal(πF, µstartnum, "startnum"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µstartnum, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstartnum = πTemp002
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp008, µlabel); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 510: if label not in self.document.nameids:
							πF.SetLineno(510)
						Label7:
							// line 511: break
							πF.SetLineno(511)
							πTemp005 = true
							continue
							goto Label8
						Label8:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 512: footnote.insert(0, nodes.label('', label))
							πF.SetLineno(512)
							πTemp007 = πF.MakeArgs(2)
							πTemp007[0] = πg.NewInt(0).ToObject()
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp009[1] = µlabel
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlabel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[1] = πTemp002
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfootnote, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfootnote, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp005 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label11
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
								µname = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[0] = µname
							πTemp009 = make([]*πg.Object, 0)
							πTemp008 = πg.NewList(πTemp009...).ToObject()
							πTemp007[1] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßfootnote_refs, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp010, ßget, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.Iter(πF, πTemp010); πE != nil {
								continue
							}
							πF.PushCheckpoint(13)
							πTemp006 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label14
							}
							if πTemp008, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µref = πTemp008
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(12)
							// line 515: ref += nodes.Text(label)
							πF.SetLineno(515)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp007[0] = µlabel
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßText, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp010, πE = πg.IAdd(πF, µref, πTemp008); πE != nil {
								continue
							}
							µref = πTemp010
							// line 516: ref.delattr('refname')
							πF.SetLineno(516)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µref, ßdelattr, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 517: assert len(footnote['ids']) == len(ref['ids']) == 1
							πF.SetLineno(517)
							πTemp007 = πF.MakeArgs(1)
							πTemp010 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µfootnote, πTemp010); πE != nil {
								continue
							}
							πTemp007[0] = πTemp012
							if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp007 = πF.MakeArgs(1)
							πTemp010 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µref, πTemp010); πE != nil {
								continue
							}
							πTemp007[0] = πTemp013
							if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp008, πE = πg.Eq(πF, πTemp012, πTemp013); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if !πTemp011 {
								goto Label15
							}
							if πTemp008, πE = πg.Eq(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
						Label15:
							if πE = πg.Assert(πF, πTemp008, nil); πE != nil {
								continue
							}
							// line 518: ref['refid'] = footnote['ids'][0]
							πF.SetLineno(518)
							πTemp008 = πg.NewInt(0).ToObject()
							πTemp012 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µfootnote, πTemp012); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp013, πTemp008); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp012 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp012, πTemp008); πE != nil {
								continue
							}
							// line 519: footnote.add_backref(ref['ids'][0])
							πF.SetLineno(519)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πg.NewInt(0).ToObject()
							πTemp012 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µref, πTemp012); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp013, πTemp008); πE != nil {
								continue
							}
							πTemp007[0] = πTemp010
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µfootnote, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 520: self.document.note_refid(ref)
							πF.SetLineno(520)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp007[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 521: ref.resolved = 1
							πF.SetLineno(521)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp008); πE != nil {
								continue
							}
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							πTemp008 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µfootnote, πTemp008); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label16
							}
							πTemp008 = ßdupnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µfootnote, πTemp008); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
						Label16:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label17
							}
							goto Label18
							// line 522: if not footnote['names'] and not footnote['dupnames']:
							πF.SetLineno(522)
						Label17:
							// line 523: footnote['names'].append(label)
							πF.SetLineno(523)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp007[0] = µlabel
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µfootnote, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 524: self.document.note_explicit_target(footnote, footnote)
							πF.SetLineno(524)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp007[0] = µfootnote
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp007[1] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_explicit_target, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 525: self.autofootnote_labels.append(label)
							πF.SetLineno(525)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp007[0] = µlabel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßautofootnote_labels, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label18
						Label18:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 526: return startnum
							πF.SetLineno(526)
							if πE = πg.CheckLocal(πF, µstartnum, "startnum"); πE != nil {
								continue
							}
							πR = µstartnum
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnumber_footnotes.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 500: """
					πF.SetLineno(500)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Assign numbers to autonumbered footnotes.\n\n        For labeled autonumbered footnotes, copy the number over to\n        corresponding footnote references.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßnumber_footnotes); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 528: def number_footnote_references(self, startnum):
					πF.SetLineno(528)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "startnum", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("number_footnote_references", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstartnum *πg.Object = πArgs[1]
						_ = µstartnum
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µlabel *πg.Object = πg.UnboundLocal
						_ = µlabel
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µfootnote *πg.Object = πg.UnboundLocal
						_ = µfootnote
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 πg.KWArgs
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 11:
								goto Label11
							case 10:
								goto Label10
							default:
								panic("unexpected function state")
							}
							// line 529: """Assign numbers to autonumbered footnote references."""
							πF.SetLineno(529)
							// line 530: i = 0
							πF.SetLineno(530)
							µi = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßautofootnote_refs, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µref = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πTemp007
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 532: if ref.resolved or ref.hasattr('refid'):
							πF.SetLineno(532)
						Label5:
							// line 533: continue
							πF.SetLineno(533)
							continue
							goto Label6
						Label6:
							// line 534: try:
							πF.SetLineno(534)
							πF.PushCheckpoint(8)
							// line 535: label = self.autofootnote_labels[i]
							πF.SetLineno(535)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßautofootnote_labels, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µlabel = πTemp003
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 536: except IndexError:
							πF.SetLineno(536)
						Label9:
							// line 537: msg = self.document.reporter.error(
							πF.SetLineno(537)
							πTemp006 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßautofootnote_labels, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Too many autonumbered footnote references: only %s corresponding footnotes available.").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"base_node", µref},
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
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µmsg = πTemp003
							// line 541: msgid = self.document.set_id(msg)
							πF.SetLineno(541)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp006[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µmsgid = πTemp002
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßautofootnote_refs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp013, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp005 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp014 = !isStop
							} else {
								πTemp014 = true
								µref = πTemp003
							}
							if πE != nil || !πTemp014 {
								continue
							}
							πF.PushCheckpoint(10)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp007
							if πTemp014, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp014 {
								goto Label13
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µref, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003 = πTemp012
						Label13:
							if πTemp014, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp014 {
								goto Label14
							}
							goto Label15
							// line 543: if ref.resolved or ref.hasattr('refname'):
							πF.SetLineno(543)
						Label14:
							// line 544: continue
							πF.SetLineno(544)
							continue
							goto Label15
						Label15:
							// line 545: prb = nodes.problematic(
							πF.SetLineno(545)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp006, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µprb = πTemp003
							// line 547: prbid = self.document.set_id(prb)
							πF.SetLineno(547)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp006[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µprbid = πTemp003
							// line 548: msg.add_backref(prbid)
							πF.SetLineno(548)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp006[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 549: ref.replace_self(prb)
							πF.SetLineno(549)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp006[0] = µprb
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 550: break
							πF.SetLineno(550)
							πTemp004 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 551: ref += nodes.Text(label)
							πF.SetLineno(551)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp006[0] = µlabel
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßText, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.IAdd(πF, µref, πTemp002); πE != nil {
								continue
							}
							µref = πTemp003
							// line 552: id = self.document.nameids[label]
							πF.SetLineno(552)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp002 = µlabel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp007, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp012, πTemp002); πE != nil {
								continue
							}
							µid = πTemp003
							// line 553: footnote = self.document.ids[id]
							πF.SetLineno(553)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp002 = µid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp007, ßids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp012, πTemp002); πE != nil {
								continue
							}
							µfootnote = πTemp003
							// line 554: ref['refid'] = id
							πF.SetLineno(554)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp003 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 555: self.document.note_refid(ref)
							πF.SetLineno(555)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp006[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 556: assert len(ref['ids']) == 1
							πF.SetLineno(556)
							πTemp006 = πF.MakeArgs(1)
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µref, πTemp003); πE != nil {
								continue
							}
							πTemp006[0] = πTemp007
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, nil); πE != nil {
								continue
							}
							// line 557: footnote.add_backref(ref['ids'][0])
							πF.SetLineno(557)
							πTemp006 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp007 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µref, πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp012, πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfootnote, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 558: ref.resolved = 1
							πF.SetLineno(558)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp002); πE != nil {
								continue
							}
							// line 559: i += 1
							πF.SetLineno(559)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp002
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
					if πE = πClass.SetItem(πF, ßnumber_footnote_references.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 529: """Assign numbers to autonumbered footnote references."""
					πF.SetLineno(529)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Assign numbers to autonumbered footnote references.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßnumber_footnote_references); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 561: def symbolize_footnotes(self):
					πF.SetLineno(561)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("symbolize_footnotes", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlabels *πg.Object = πg.UnboundLocal
						_ = µlabels
						var µfootnote *πg.Object = πg.UnboundLocal
						_ = µfootnote
						var µreps *πg.Object = πg.UnboundLocal
						_ = µreps
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µlabeltext *πg.Object = πg.UnboundLocal
						_ = µlabeltext
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πTemp012 πg.KWArgs
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
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
							case 8:
								goto Label8
							case 10:
								goto Label10
							case 11:
								goto Label11
							default:
								panic("unexpected function state")
							}
							// line 562: """Add symbols indexes to "[*]"-style footnotes and references."""
							πF.SetLineno(562)
							// line 563: labels = []
							πF.SetLineno(563)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlabels = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsymbol_footnotes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								µfootnote = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 565: reps, index = divmod(self.document.symbol_footnote_start,
							πF.SetLineno(565)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsymbol_footnote_start, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsymbols, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdivmod); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
								continue
							}
							µreps = πTemp003
							µindex = πTemp008
							// line 567: labeltext = self.symbols[index] * (reps + 1)
							πF.SetLineno(567)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp004 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßsymbols, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µreps, "reps"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µreps, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							µlabeltext = πTemp003
							// line 568: labels.append(labeltext)
							πF.SetLineno(568)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabeltext, "labeltext"); πE != nil {
								continue
							}
							πTemp001[0] = µlabeltext
							if πE = πg.CheckLocal(πF, µlabels, "labels"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µlabels, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 569: footnote.insert(0, nodes.label('', labeltext))
							πF.SetLineno(569)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp007 = πF.MakeArgs(2)
							πTemp007[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µlabeltext, "labeltext"); πE != nil {
								continue
							}
							πTemp007[1] = µlabeltext
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlabel, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfootnote, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 570: self.document.symbol_footnote_start += 1
							πF.SetLineno(570)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsymbol_footnote_start, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp008, ßsymbol_footnote_start, πTemp003); πE != nil {
								continue
							}
							// line 571: self.document.set_id(footnote)
							πF.SetLineno(571)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp001[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 572: i = 0
							πF.SetLineno(572)
							µi = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsymbol_footnote_refs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label6
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
								µref = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 574: try:
							πF.SetLineno(574)
							πF.PushCheckpoint(8)
							// line 575: ref += nodes.Text(labels[i])
							πF.SetLineno(575)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.CheckLocal(πF, µlabels, "labels"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µlabels, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IAdd(πF, µref, πTemp003); πE != nil {
								continue
							}
							µref = πTemp004
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 576: except IndexError:
							πF.SetLineno(576)
						Label9:
							// line 577: msg = self.document.reporter.error(
							πF.SetLineno(577)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlabels, "labels"); πE != nil {
								continue
							}
							πTemp007[0] = µlabels
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Too many symbol footnote references: only %s corresponding footnotes available.").ToObject(), πTemp008); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp012 = πg.KWArgs{
								{"base_node", µref},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßerror, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp012); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsg = πTemp004
							// line 581: msgid = self.document.set_id(msg)
							πF.SetLineno(581)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmsgid = πTemp003
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp009, ßsymbol_footnote_refs, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp013, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp006 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp014 = !isStop
							} else {
								πTemp014 = true
								µref = πTemp004
							}
							if πE != nil || !πTemp014 {
								continue
							}
							πF.PushCheckpoint(10)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp008
							if πTemp014, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp014 {
								goto Label13
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µref, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp004 = πTemp009
						Label13:
							if πTemp014, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp014 {
								goto Label14
							}
							goto Label15
							// line 583: if ref.resolved or ref.hasattr('refid'):
							πF.SetLineno(583)
						Label14:
							// line 584: continue
							πF.SetLineno(584)
							continue
							goto Label15
						Label15:
							// line 585: prb = nodes.problematic(
							πF.SetLineno(585)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp012 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp001, πTemp012); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µprb = πTemp004
							// line 587: prbid = self.document.set_id(prb)
							πF.SetLineno(587)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp001[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µprbid = πTemp004
							// line 588: msg.add_backref(prbid)
							πF.SetLineno(588)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp001[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 589: ref.replace_self(prb)
							πF.SetLineno(589)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp001[0] = µprb
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 590: break
							πF.SetLineno(590)
							πTemp005 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 591: footnote = self.document.symbol_footnotes[i]
							πF.SetLineno(591)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßsymbol_footnotes, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							µfootnote = πTemp004
							// line 592: assert len(footnote['ids']) == 1
							πF.SetLineno(592)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfootnote, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 593: ref['refid'] = footnote['ids'][0]
							πF.SetLineno(593)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp008 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µfootnote, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp008 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp008, πTemp003); πE != nil {
								continue
							}
							// line 594: self.document.note_refid(ref)
							πF.SetLineno(594)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp001[0] = µref
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßnote_refid, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 595: footnote.add_backref(ref['ids'][0])
							πF.SetLineno(595)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp008 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µref, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfootnote, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 596: i += 1
							πF.SetLineno(596)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp003
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsymbolize_footnotes.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 562: """Add symbols indexes to "[*]"-style footnotes and references."""
					πF.SetLineno(562)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Add symbols indexes to \"[*]\"-style footnotes and references.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsymbolize_footnotes); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 598: def resolve_footnotes_and_citations(self):
					πF.SetLineno(598)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("resolve_footnotes_and_citations", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfootnote *πg.Object = πg.UnboundLocal
						_ = µfootnote
						var µlabel *πg.Object = πg.UnboundLocal
						_ = µlabel
						var µreflist *πg.Object = πg.UnboundLocal
						_ = µreflist
						var µcitation *πg.Object = πg.UnboundLocal
						_ = µcitation
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 12:
								goto Label12
							case 13:
								goto Label13
							default:
								panic("unexpected function state")
							}
							// line 599: """
							πF.SetLineno(599)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfootnotes, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µfootnote = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfootnote, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µlabel = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßfootnote_refs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp008, µlabel); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							goto Label8
							// line 605: if label in self.document.footnote_refs:
							πF.SetLineno(605)
						Label7:
							// line 606: reflist = self.document.footnote_refs[label]
							πF.SetLineno(606)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp003 = µlabel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßfootnote_refs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							µreflist = πTemp006
							// line 607: self.resolve_references(footnote, reflist)
							πF.SetLineno(607)
							πTemp010 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp010[0] = µfootnote
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							πTemp010[1] = µreflist
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßresolve_references, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							goto Label8
						Label8:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcitations, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp004 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
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
								πTemp005 = !isStop
							} else {
								πTemp005 = true
								µcitation = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µcitation, "citation"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µcitation, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(13)
							πTemp005 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label14
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µlabel = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(12)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßcitation_refs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp008, µlabel); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label15
							}
							goto Label16
							// line 610: if label in self.document.citation_refs:
							πF.SetLineno(610)
						Label15:
							// line 611: reflist = self.document.citation_refs[label]
							πF.SetLineno(611)
							if πE = πg.CheckLocal(πF, µlabel, "label"); πE != nil {
								continue
							}
							πTemp003 = µlabel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßcitation_refs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							µreflist = πTemp006
							// line 612: self.resolve_references(citation, reflist)
							πF.SetLineno(612)
							πTemp010 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcitation, "citation"); πE != nil {
								continue
							}
							πTemp010[0] = µcitation
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							πTemp010[1] = µreflist
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßresolve_references, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							goto Label16
						Label16:
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßresolve_footnotes_and_citations.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 599: """
					πF.SetLineno(599)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Link manually-labeled footnotes and citations to/from their\n        references.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßresolve_footnotes_and_citations); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 614: def resolve_references(self, note, reflist):
					πF.SetLineno(614)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "note", Def: nil}
					πTemp003[2] = πg.Param{Name: "reflist", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("resolve_references", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnote *πg.Object = πArgs[1]
						_ = µnote
						var µreflist *πg.Object = πArgs[2]
						_ = µreflist
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
							// line 615: assert len(note['ids']) == 1
							πF.SetLineno(615)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnote, "note"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnote, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 616: id = note['ids'][0]
							πF.SetLineno(616)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnote, "note"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnote, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µid = πTemp003
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µreflist); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µref = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßresolved, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 618: if ref.resolved:
							πF.SetLineno(618)
						Label4:
							// line 619: continue
							πF.SetLineno(619)
							continue
							goto Label5
						Label5:
							// line 620: ref.delattr('refname')
							πF.SetLineno(620)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßdelattr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 621: ref['refid'] = id
							πF.SetLineno(621)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp004 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µref, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 622: assert len(ref['ids']) == 1
							πF.SetLineno(622)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µref, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 623: note.add_backref(ref['ids'][0])
							πF.SetLineno(623)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp005 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µref, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µnote, "note"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnote, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 624: ref.resolved = 1
							πF.SetLineno(624)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µref, ßresolved, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 625: note.resolved = 1
							πF.SetLineno(625)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnote, "note"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnote, ßresolved, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßresolve_references.ToObject(), πTemp008); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Footnotes").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFootnotes.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 628: class CircularSubstitutionDefinitionError(Exception): pass
			πF.SetLineno(628)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("CircularSubstitutionDefinitionError", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 628: class CircularSubstitutionDefinitionError(Exception): pass
					πF.SetLineno(628)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("CircularSubstitutionDefinitionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCircularSubstitutionDefinitionError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 631: class Substitutions(Transform):
			πF.SetLineno(631)
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
			_, πE = πg.NewCode("Substitutions", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 633: """
					πF.SetLineno(633)
					// line 633: """
					πF.SetLineno(633)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Given the following ``document`` as input::\n\n        <document>\n            <paragraph>\n                The\n                <substitution_reference refname=\"biohazard\">\n                    biohazard\n                 symbol is deservedly scary-looking.\n            <substitution_definition name=\"biohazard\">\n                <image alt=\"biohazard\" uri=\"biohazard.png\">\n\n    The ``substitution_reference`` will simply be replaced by the\n    contents of the corresponding ``substitution_definition``.\n\n    The transformed result will be::\n\n        <document>\n            <paragraph>\n                The\n                <image alt=\"biohazard\" uri=\"biohazard.png\">\n                 symbol is deservedly scary-looking.\n            <substitution_definition name=\"biohazard\">\n                <image alt=\"biohazard\" uri=\"biohazard.png\">\n    ").ToObject()); πE != nil {
						continue
					}
					// line 659: default_priority = 220
					πF.SetLineno(659)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(220).ToObject()); πE != nil {
						continue
					}
					// line 660: """The Substitutions transform has to be applied very early, before
					πF.SetLineno(660)
					// line 663: def apply(self):
					πF.SetLineno(663)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdefs *πg.Object = πg.UnboundLocal
						_ = µdefs
						var µnormed *πg.Object = πg.UnboundLocal
						_ = µnormed
						var µsubreflist *πg.Object = πg.UnboundLocal
						_ = µsubreflist
						var µnested *πg.Object = πg.UnboundLocal
						_ = µnested
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µrefname *πg.Object = πg.UnboundLocal
						_ = µrefname
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µnormed_name *πg.Object = πg.UnboundLocal
						_ = µnormed_name
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
						var µsubdef *πg.Object = πg.UnboundLocal
						_ = µsubdef
						var µparent *πg.Object = πg.UnboundLocal
						_ = µparent
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µsubdef_copy *πg.Object = πg.UnboundLocal
						_ = µsubdef_copy
						var µnested_ref *πg.Object = πg.UnboundLocal
						_ = µnested_ref
						var µnested_name *πg.Object = πg.UnboundLocal
						_ = µnested_name
						var µref_origin *πg.Object = πg.UnboundLocal
						_ = µref_origin
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
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 32:
								goto Label32
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 33:
								goto Label33
							case 43:
								goto Label43
							case 44:
								goto Label44
							case 25:
								goto Label25
							case 26:
								goto Label26
							case 27:
								goto Label27
							default:
								panic("unexpected function state")
							}
							// line 664: defs = self.document.substitution_defs
							πF.SetLineno(664)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsubstitution_defs, nil); πE != nil {
								continue
							}
							µdefs = πTemp002
							// line 665: normed = self.document.substitution_names
							πF.SetLineno(665)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsubstitution_names, nil); πE != nil {
								continue
							}
							µnormed = πTemp002
							// line 666: subreflist = list(self.document.traverse(nodes.substitution_reference))
							πF.SetLineno(666)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsubstitution_reference, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µsubreflist = πTemp002
							// line 667: nested = {}
							πF.SetLineno(667)
							πTemp005 = πg.NewDict()
							πTemp001 = πTemp005.ToObject()
							µnested = πTemp001
							if πE = πg.CheckLocal(πF, µsubreflist, "subreflist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µsubreflist); πE != nil {
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
								µref = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 669: refname = ref['refname']
							πF.SetLineno(669)
							πTemp002 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µref, πTemp002); πE != nil {
								continue
							}
							µrefname = πTemp008
							// line 670: key = None
							πF.SetLineno(670)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µkey = πTemp002
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdefs, "defs"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µdefs, µrefname); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 671: if refname in defs:
							πF.SetLineno(671)
						Label4:
							// line 672: key = refname
							πF.SetLineno(672)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							µkey = µrefname
							goto Label6
						Label5:
							// line 674: normed_name = refname.lower()
							πF.SetLineno(674)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µrefname, ßlower, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnormed_name = πTemp008
							if πE = πg.CheckLocal(πF, µnormed_name, "normed_name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnormed, "normed"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnormed, µnormed_name); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							goto Label8
							// line 675: if normed_name in normed:
							πF.SetLineno(675)
						Label7:
							// line 676: key = normed[normed_name]
							πF.SetLineno(676)
							if πE = πg.CheckLocal(πF, µnormed_name, "normed_name"); πE != nil {
								continue
							}
							πTemp002 = µnormed_name
							if πE = πg.CheckLocal(πF, µnormed, "normed"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnormed, πTemp002); πE != nil {
								continue
							}
							µkey = πTemp008
							goto Label8
						Label8:
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µkey == πTemp008).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 677: if key is None:
							πF.SetLineno(677)
						Label9:
							// line 678: msg = self.document.reporter.error(
							πF.SetLineno(678)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Undefined substitution referenced: \"%s\".").ToObject(), µrefname); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"base_node", µref},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp008, ßerror, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmsg = πTemp008
							// line 681: msgid = self.document.set_id(msg)
							πF.SetLineno(681)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmsgid = πTemp002
							// line 682: prb = nodes.problematic(
							πF.SetLineno(682)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µprb = πTemp002
							// line 684: prbid = self.document.set_id(prb)
							πF.SetLineno(684)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp003[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µprbid = πTemp002
							// line 685: msg.add_backref(prbid)
							πF.SetLineno(685)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp003[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 686: ref.replace_self(prb)
							πF.SetLineno(686)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp003[0] = µprb
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label11
						Label10:
							// line 688: subdef = defs[key]
							πF.SetLineno(688)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µdefs, "defs"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdefs, πTemp002); πE != nil {
								continue
							}
							µsubdef = πTemp008
							// line 689: parent = ref.parent
							πF.SetLineno(689)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßparent, nil); πE != nil {
								continue
							}
							µparent = πTemp002
							// line 690: index = parent.index(ref)
							πF.SetLineno(690)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp003[0] = µref
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparent, ßindex, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µindex = πTemp008
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µsubdef, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp010, ßltrim.ToObject()); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp011).ToObject()
							πTemp002 = πTemp008
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µsubdef, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp010, ßtrim.ToObject()); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp011).ToObject()
							πTemp002 = πTemp008
						Label12:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label13
							}
							goto Label14
							// line 691: if  ('ltrim' in subdef.attributes
							πF.SetLineno(691)
						Label13:
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GT(πF, µindex, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp008
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label15
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Sub(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µparent, πTemp008); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßText, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp010
							if πTemp008, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πTemp010
						Label15:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label16
							}
							goto Label17
							// line 693: if index > 0 and isinstance(parent[index - 1],
							πF.SetLineno(693)
						Label16:
							// line 695: parent[index - 1] = parent[index - 1].rstrip()
							πF.SetLineno(695)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp008
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µparent, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp008, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Sub(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp012
							if πE = πg.SetItem(πF, µparent, πTemp010, πTemp002); πE != nil {
								continue
							}
							goto Label17
						Label17:
							goto Label14
						Label14:
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µsubdef, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp010, ßrtrim.ToObject()); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp011).ToObject()
							πTemp002 = πTemp008
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label18
							}
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µsubdef, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, πTemp010, ßtrim.ToObject()); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp011).ToObject()
							πTemp002 = πTemp008
						Label18:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label19
							}
							goto Label20
							// line 696: if  ('rtrim' in subdef.attributes
							πF.SetLineno(696)
						Label19:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πTemp003[0] = µparent
							if πTemp010, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GT(πF, πTemp012, πTemp010); πE != nil {
								continue
							}
							πTemp002 = πTemp008
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label21
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µparent, πTemp008); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßText, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp010
							if πTemp008, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πTemp010
						Label21:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label22
							}
							goto Label23
							// line 698: if  (len(parent) > index + 1
							πF.SetLineno(698)
						Label22:
							// line 700: parent[index + 1] = parent[index + 1].lstrip()
							πF.SetLineno(700)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp008
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µparent, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp008, ßlstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp012
							if πE = πg.SetItem(πF, µparent, πTemp010, πTemp002); πE != nil {
								continue
							}
							goto Label23
						Label23:
							goto Label20
						Label20:
							// line 701: subdef_copy = subdef.deepcopy()
							πF.SetLineno(701)
							if πE = πg.CheckLocal(πF, µsubdef, "subdef"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsubdef, ßdeepcopy, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsubdef_copy = πTemp008
							// line 702: try:
							πF.SetLineno(702)
							πF.PushCheckpoint(25)
							πTemp003 = πF.MakeArgs(1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßsubstitution_reference, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp010
							if πE = πg.CheckLocal(πF, µsubdef_copy, "subdef_copy"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µsubdef_copy, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Iter(πF, πTemp010); πE != nil {
								continue
							}
							πF.PushCheckpoint(27)
							πTemp007 = false
						Label26:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label28
							}
							if πTemp008, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µnested_ref = πTemp008
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(26)
							// line 706: nested_name = normed[nested_ref['refname'].lower()]
							πF.SetLineno(706)
							πTemp010 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnested_ref, "nested_ref"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µnested_ref, πTemp010); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp012, ßlower, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008 = πTemp012
							if πE = πg.CheckLocal(πF, µnormed, "normed"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µnormed, πTemp008); πE != nil {
								continue
							}
							µnested_name = πTemp010
							if πE = πg.CheckLocal(πF, µnested_name, "nested_name"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnested_name, "nested_name"); πE != nil {
								continue
							}
							πTemp003[0] = µnested_name
							πTemp004 = make([]*πg.Object, 0)
							πTemp010 = πg.NewList(πTemp004...).ToObject()
							πTemp003[1] = πTemp010
							if πE = πg.CheckLocal(πF, µnested, "nested"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µnested, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp011, πE = πg.Contains(πF, πTemp012, µnested_name); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp011).ToObject()
							if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label29
							}
							goto Label30
							// line 707: if nested_name in nested.setdefault(nested_name, []):
							πF.SetLineno(707)
						Label29:
							if πTemp008, πE = πg.ResolveGlobal(πF, ßCircularSubstitutionDefinitionError); πE != nil {
								continue
							}
							// line 708: raise CircularSubstitutionDefinitionError
							πF.SetLineno(708)
							πE = πF.Raise(πTemp008, nil, nil)
							continue
							goto Label31
						Label30:
							// line 710: nested[nested_name].append(key)
							πF.SetLineno(710)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003[0] = µkey
							if πE = πg.CheckLocal(πF, µnested_name, "nested_name"); πE != nil {
								continue
							}
							πTemp008 = µnested_name
							if πE = πg.CheckLocal(πF, µnested, "nested"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µnested, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp010, ßappend, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 711: nested_ref['ref-origin'] = ref
							πF.SetLineno(711)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, µref); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnested_ref, "nested_ref"); πE != nil {
								continue
							}
							πTemp010 = πg.NewStr("ref-origin").ToObject()
							if πE = πg.SetItem(πF, µnested_ref, πTemp010, πTemp008); πE != nil {
								continue
							}
							// line 712: subreflist.append(nested_ref)
							πF.SetLineno(712)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnested_ref, "nested_ref"); πE != nil {
								continue
							}
							πTemp003[0] = µnested_ref
							if πE = πg.CheckLocal(πF, µsubreflist, "subreflist"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µsubreflist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label31
						Label31:
							continue
						Label27:
							if πE != nil || πR != nil {
								continue
							}
						Label28:
							πF.PopCheckpoint()
							// line 737: ref.replace_self(subdef_copy.children)
							πF.SetLineno(737)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubdef_copy, "subdef_copy"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsubdef_copy, ßchildren, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µsubdef_copy, "subdef_copy"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µsubdef_copy, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(33)
							πTemp007 = false
						Label32:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label34
							}
							if πTemp008, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µnode = πTemp008
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(32)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πTemp008, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßReferential, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp010
							if πTemp008, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp011, πE = πg.IsTrue(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label35
							}
							goto Label36
							// line 741: if isinstance(node, nodes.Referential):
							πF.SetLineno(741)
						Label35:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, µnode, ßrefname.ToObject()); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp011).ToObject()
							if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label37
							}
							goto Label38
							// line 744: if 'refname' in node:
							πF.SetLineno(744)
						Label37:
							// line 745: self.document.note_refname(node)
							πF.SetLineno(745)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßnote_refname, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label38
						Label38:
							goto Label36
						Label36:
							continue
						Label33:
							if πE != nil || πR != nil {
								continue
							}
						Label34:
							goto Label24
						Label25:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp013, πTemp014 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCircularSubstitutionDefinitionError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label39
							}
							πE = πF.Raise(πTemp013.ToObject(), nil, πTemp014.ToObject())
							continue
							// line 713: except CircularSubstitutionDefinitionError:
							πF.SetLineno(713)
						Label39:
							// line 714: parent = ref.parent
							πF.SetLineno(714)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßparent, nil); πE != nil {
								continue
							}
							µparent = πTemp002
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πTemp003[0] = µparent
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßsubstitution_definition, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp008
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label40
							}
							goto Label41
							// line 715: if isinstance(parent, nodes.substitution_definition):
							πF.SetLineno(715)
						Label40:
							// line 716: msg = self.document.reporter.error(
							πF.SetLineno(716)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("Circular substitution definition detected:").ToObject()
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparent, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparent, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparent, ßline, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"line", πTemp002},
								{"base_node", µparent},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp008, ßerror, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmsg = πTemp008
							// line 721: parent.replace_self(msg)
							πF.SetLineno(721)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003[0] = µmsg
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparent, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label42
						Label41:
							// line 724: ref_origin = ref
							πF.SetLineno(724)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							µref_origin = µref
							// line 725: while ref_origin.hasattr('ref-origin'):
							πF.SetLineno(725)
							πF.PushCheckpoint(44)
							πTemp007 = false
						Label43:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label45
							}
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("ref-origin").ToObject()
							if πE = πg.CheckLocal(πF, µref_origin, "ref_origin"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref_origin, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(43)
							// line 726: ref_origin = ref_origin['ref-origin']
							πF.SetLineno(726)
							πTemp002 = πg.NewStr("ref-origin").ToObject()
							if πE = πg.CheckLocal(πF, µref_origin, "ref_origin"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µref_origin, πTemp002); πE != nil {
								continue
							}
							µref_origin = πTemp008
							continue
						Label44:
							if πE != nil || πR != nil {
								continue
							}
						Label45:
							// line 727: msg = self.document.reporter.error(
							πF.SetLineno(727)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Circular substitution definition referenced: \"%s\".").ToObject(), µrefname); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µref_origin, "ref_origin"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"base_node", µref_origin},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp008, ßerror, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmsg = πTemp008
							// line 730: msgid = self.document.set_id(msg)
							πF.SetLineno(730)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmsgid = πTemp002
							// line 731: prb = nodes.problematic(
							πF.SetLineno(731)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µprb = πTemp002
							// line 733: prbid = self.document.set_id(prb)
							πF.SetLineno(733)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp003[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µprbid = πTemp002
							// line 734: msg.add_backref(prbid)
							πF.SetLineno(734)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp003[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 735: ref.replace_self(prb)
							πF.SetLineno(735)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp003[0] = µprb
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µref, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label42
						Label42:
							πF.RestoreExc(nil, nil)
							goto Label24
						Label24:
							goto Label11
						Label11:
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Substitutions").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSubstitutions.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 748: class TargetNotes(Transform):
			πF.SetLineno(748)
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
			_, πE = πg.NewCode("TargetNotes", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 750: """
					πF.SetLineno(750)
					// line 750: """
					πF.SetLineno(750)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Creates a footnote for each external target in the text, and corresponding\n    footnote references after each reference.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 755: default_priority = 540
					πF.SetLineno(755)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(540).ToObject()); πE != nil {
						continue
					}
					// line 756: """The TargetNotes transform has to be applied after `IndirectHyperlinks`
					πF.SetLineno(756)
					// line 760: def __init__(self, document, startnode):
					πF.SetLineno(760)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp002[2] = πg.Param{Name: "startnode", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µstartnode *πg.Object = πArgs[2]
						_ = µstartnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 761: Transform.__init__(self, document, startnode=startnode)
							πF.SetLineno(761)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"startnode", µstartnode},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 763: self.classes = startnode.details.get('class', [])
							πF.SetLineno(763)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßclass.ToObject()
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstartnode, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßclasses, πTemp004); πE != nil {
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
					// line 765: def apply(self):
					πF.SetLineno(765)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnotes *πg.Object = πg.UnboundLocal
						_ = µnotes
						var µnodelist *πg.Object = πg.UnboundLocal
						_ = µnodelist
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
						var µnames *πg.Object = πg.UnboundLocal
						_ = µnames
						var µrefs *πg.Object = πg.UnboundLocal
						_ = µrefs
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µfootnote *πg.Object = πg.UnboundLocal
						_ = µfootnote
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 6:
								goto Label6
							case 7:
								goto Label7
							case 13:
								goto Label13
							case 14:
								goto Label14
							default:
								panic("unexpected function state")
							}
							// line 766: notes = {}
							πF.SetLineno(766)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µnotes = πTemp002
							// line 767: nodelist = []
							πF.SetLineno(767)
							πTemp003 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							µnodelist = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßtarget, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µtarget = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtarget, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 770: if not target.hasattr('refuri'):
							πF.SetLineno(770)
						Label4:
							// line 771: continue
							πF.SetLineno(771)
							continue
							goto Label5
						Label5:
							// line 772: names = target['names']
							πF.SetLineno(772)
							πTemp004 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							µnames = πTemp005
							// line 773: refs = []
							πF.SetLineno(773)
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							µrefs = πTemp004
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µnames); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp007 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µname = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 775: refs.extend(self.document.refnames.get(name, []))
							πF.SetLineno(775)
							πTemp003 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp010[0] = µname
							πTemp011 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp011...).ToObject()
							πTemp010[1] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßrefnames, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp008, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp003[0] = πTemp008
							if πE = πg.CheckLocal(πF, µrefs, "refs"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µrefs, ßextend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							if πE = πg.CheckLocal(πF, µrefs, "refs"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µrefs); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 776: if not refs:
							πF.SetLineno(776)
						Label9:
							// line 777: continue
							πF.SetLineno(777)
							continue
							goto Label10
						Label10:
							// line 778: footnote = self.make_target_footnote(target['refuri'], refs,
							πF.SetLineno(778)
							πTemp003 = πF.MakeArgs(3)
							πTemp004 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µtarget, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µrefs, "refs"); πE != nil {
								continue
							}
							πTemp003[1] = µrefs
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							πTemp003[2] = µnotes
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmake_target_footnote, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µfootnote = πTemp005
							πTemp005 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µtarget, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnotes, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label11
							}
							goto Label12
							// line 780: if target['refuri'] not in notes:
							πF.SetLineno(780)
						Label11:
							// line 781: notes[target['refuri']] = footnote
							πF.SetLineno(781)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µfootnote); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							πTemp008 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µtarget, πTemp008); πE != nil {
								continue
							}
							πTemp005 = πTemp012
							if πE = πg.SetItem(πF, µnotes, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 782: nodelist.append(footnote)
							πF.SetLineno(782)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp003[0] = µfootnote
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnodelist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
						Label12:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßreference, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßtraverse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp006 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µref = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(13)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßanonymous.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µref, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label16
							}
							goto Label17
							// line 785: if not ref.get('anonymous'):
							πF.SetLineno(785)
						Label16:
							// line 786: continue
							πF.SetLineno(786)
							continue
							goto Label17
						Label17:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µref, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label18
							}
							goto Label19
							// line 787: if ref.hasattr('refuri'):
							πF.SetLineno(787)
						Label18:
							// line 788: footnote = self.make_target_footnote(ref['refuri'], [ref],
							πF.SetLineno(788)
							πTemp003 = πF.MakeArgs(3)
							πTemp004 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µref, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp010 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp010[0] = µref
							πTemp004 = πg.NewList(πTemp010...).ToObject()
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							πTemp003[2] = µnotes
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmake_target_footnote, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µfootnote = πTemp005
							πTemp005 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µref, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnotes, πTemp008); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label20
							}
							goto Label21
							// line 790: if ref['refuri'] not in notes:
							πF.SetLineno(790)
						Label20:
							// line 791: notes[ref['refuri']] = footnote
							πF.SetLineno(791)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µfootnote); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							πTemp008 = ßrefuri.ToObject()
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µref, πTemp008); πE != nil {
								continue
							}
							πTemp005 = πTemp012
							if πE = πg.SetItem(πF, µnotes, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 792: nodelist.append(footnote)
							πF.SetLineno(792)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp003[0] = µfootnote
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnodelist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label21
						Label21:
							goto Label19
						Label19:
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 793: self.startnode.replace_self(nodelist)
							πF.SetLineno(793)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							πTemp003[0] = µnodelist
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 795: def make_target_footnote(self, refuri, refs, notes):
					πF.SetLineno(795)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "refuri", Def: nil}
					πTemp002[2] = πg.Param{Name: "refs", Def: nil}
					πTemp002[3] = πg.Param{Name: "notes", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("make_target_footnote", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrefuri *πg.Object = πArgs[1]
						_ = µrefuri
						var µrefs *πg.Object = πArgs[2]
						_ = µrefs
						var µnotes *πg.Object = πArgs[3]
						_ = µnotes
						var µfootnote *πg.Object = πg.UnboundLocal
						_ = µfootnote
						var µfootnote_name *πg.Object = πg.UnboundLocal
						_ = µfootnote_name
						var µfootnote_id *πg.Object = πg.UnboundLocal
						_ = µfootnote_id
						var µfootnote_paragraph *πg.Object = πg.UnboundLocal
						_ = µfootnote_paragraph
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var µrefnode *πg.Object = πg.UnboundLocal
						_ = µrefnode
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µreflist *πg.Object = πg.UnboundLocal
						_ = µreflist
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µrefuri, "refuri"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnotes, µrefuri); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 796: if refuri in notes:  # duplicate?
							πF.SetLineno(796)
						Label1:
							// line 797: footnote = notes[refuri]
							πF.SetLineno(797)
							if πE = πg.CheckLocal(πF, µrefuri, "refuri"); πE != nil {
								continue
							}
							πTemp001 = µrefuri
							if πE = πg.CheckLocal(πF, µnotes, "notes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnotes, πTemp001); πE != nil {
								continue
							}
							µfootnote = πTemp003
							// line 798: assert len(footnote['names']) == 1
							πF.SetLineno(798)
							πTemp004 = πF.MakeArgs(1)
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µfootnote, πTemp003); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 799: footnote_name = footnote['names'][0]
							πF.SetLineno(799)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp005 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfootnote, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µfootnote_name = πTemp003
							goto Label3
						Label2:
							// line 801: footnote = nodes.footnote()
							πF.SetLineno(801)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßfootnote, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfootnote = πTemp001
							// line 802: footnote_id = self.document.set_id(footnote)
							πF.SetLineno(802)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp004[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µfootnote_id = πTemp001
							// line 805: footnote_name = 'TARGET_NOTE: ' + footnote_id
							πF.SetLineno(805)
							if πE = πg.CheckLocal(πF, µfootnote_id, "footnote_id"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πg.NewStr("TARGET_NOTE: ").ToObject(), µfootnote_id); πE != nil {
								continue
							}
							µfootnote_name = πTemp001
							// line 806: footnote['auto'] = 1
							πF.SetLineno(806)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp003 = ßauto.ToObject()
							if πE = πg.SetItem(πF, µfootnote, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 807: footnote['names'] = [footnote_name]
							πF.SetLineno(807)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µfootnote_name, "footnote_name"); πE != nil {
								continue
							}
							πTemp004[0] = µfootnote_name
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp005 = ßnames.ToObject()
							if πE = πg.SetItem(πF, µfootnote, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 808: footnote_paragraph = nodes.paragraph()
							πF.SetLineno(808)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfootnote_paragraph = πTemp001
							// line 809: footnote_paragraph += nodes.reference('', refuri, refuri=refuri)
							πF.SetLineno(809)
							if πE = πg.CheckLocal(πF, µfootnote_paragraph, "footnote_paragraph"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µrefuri, "refuri"); πE != nil {
								continue
							}
							πTemp004[1] = µrefuri
							if πE = πg.CheckLocal(πF, µrefuri, "refuri"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"refuri", µrefuri},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IAdd(πF, µfootnote_paragraph, πTemp001); πE != nil {
								continue
							}
							µfootnote_paragraph = πTemp003
							// line 810: footnote += footnote_paragraph
							πF.SetLineno(810)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfootnote_paragraph, "footnote_paragraph"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µfootnote, µfootnote_paragraph); πE != nil {
								continue
							}
							µfootnote = πTemp001
							// line 811: self.document.note_autofootnote(footnote)
							πF.SetLineno(811)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp004[0] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnote_autofootnote, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 812: self.document.note_explicit_target(footnote, footnote)
							πF.SetLineno(812)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp004[0] = µfootnote
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πTemp004[1] = µfootnote
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnote_explicit_target, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µrefs, "refs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µrefs); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp002 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µref = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp004[0] = µref
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 814: if isinstance(ref, nodes.target):
							πF.SetLineno(814)
						Label7:
							// line 815: continue
							πF.SetLineno(815)
							continue
							goto Label8
						Label8:
							// line 816: refnode = nodes.footnote_reference(refname=footnote_name, auto=1)
							πF.SetLineno(816)
							if πE = πg.CheckLocal(πF, µfootnote_name, "footnote_name"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"refname", µfootnote_name},
								{"auto", πg.NewInt(1).ToObject()},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßfootnote_reference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							µrefnode = πTemp003
							// line 817: refnode['classes'] += self.classes
							πF.SetLineno(817)
							πTemp003 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µrefnode, "refnode"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µrefnode, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßclasses, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IAdd(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrefnode, "refnode"); πE != nil {
								continue
							}
							πTemp009 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µrefnode, πTemp009, πTemp006); πE != nil {
								continue
							}
							// line 818: self.document.note_autofootnote_ref(refnode)
							πF.SetLineno(818)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefnode, "refnode"); πE != nil {
								continue
							}
							πTemp004[0] = µrefnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßnote_autofootnote_ref, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 819: self.document.note_footnote_ref(refnode)
							πF.SetLineno(819)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefnode, "refnode"); πE != nil {
								continue
							}
							πTemp004[0] = µrefnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßnote_footnote_ref, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 820: index = ref.parent.index(ref) + 1
							πF.SetLineno(820)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp004[0] = µref
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µref, ßparent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßindex, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µindex = πTemp003
							// line 821: reflist = [refnode]
							πF.SetLineno(821)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µrefnode, "refnode"); πE != nil {
								continue
							}
							πTemp004[0] = µrefnode
							πTemp003 = πg.NewList(πTemp004...).ToObject()
							µreflist = πTemp003
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßget_trim_footnote_ref_space, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 822: if not utils.get_trim_footnote_ref_space(self.document.settings):
							πF.SetLineno(822)
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßclasses, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label11
							}
							goto Label12
							// line 823: if self.classes:
							πF.SetLineno(823)
						Label11:
							// line 824: reflist.insert(0, nodes.inline(text=' ', Classes=self.classes))
							πF.SetLineno(824)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßclasses, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"text", πg.NewStr(" ").ToObject()},
								{"Classes", πTemp003},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßinline, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µreflist, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label12:
							// line 826: reflist.insert(0, nodes.Text(' '))
							πF.SetLineno(826)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp010 = πF.MakeArgs(1)
							πTemp010[0] = πg.NewStr(" ").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp004[1] = πTemp003
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µreflist, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label13:
							goto Label10
						Label10:
							// line 827: ref.parent.insert(index, reflist)
							πF.SetLineno(827)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp004[0] = µindex
							if πE = πg.CheckLocal(πF, µreflist, "reflist"); πE != nil {
								continue
							}
							πTemp004[1] = µreflist
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µref, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 828: return footnote
							πF.SetLineno(828)
							if πE = πg.CheckLocal(πF, µfootnote, "footnote"); πE != nil {
								continue
							}
							πR = µfootnote
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßmake_target_footnote.ToObject(), πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TargetNotes").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTargetNotes.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 831: class DanglingReferences(Transform):
			πF.SetLineno(831)
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
			_, πE = πg.NewCode("DanglingReferences", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 833: """
					πF.SetLineno(833)
					// line 833: """
					πF.SetLineno(833)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Check for dangling references (incl. footnote & citation) and for\n    unreferenced targets.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 838: default_priority = 850
					πF.SetLineno(838)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(850).ToObject()); πE != nil {
						continue
					}
					// line 840: def apply(self):
					πF.SetLineno(840)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
						var µtarget *πg.Object = πg.UnboundLocal
						_ = µtarget
						var µnaming *πg.Object = πg.UnboundLocal
						_ = µnaming
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
							// line 841: visitor = DanglingReferencesVisitor(
							πF.SetLineno(841)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtransformer, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßunknown_reference_resolvers, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßDanglingReferencesVisitor); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvisitor = πTemp003
							// line 844: self.document.walk(visitor)
							πF.SetLineno(844)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp001[0] = µvisitor
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwalk, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtarget, nil); πE != nil {
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
								µtarget = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtarget, ßreferenced, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
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
							// line 848: if not target.referenced:
							πF.SetLineno(848)
						Label4:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßanonymous.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µtarget, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 849: if target.get('anonymous'):
							πF.SetLineno(849)
						Label6:
							// line 854: continue
							πF.SetLineno(854)
							continue
							goto Label7
						Label7:
							πTemp003 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							πTemp003 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 855: if target['names']:
							πF.SetLineno(855)
						Label8:
							// line 856: naming = target['names'][0]
							πF.SetLineno(856)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp007 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µtarget, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							µnaming = πTemp004
							goto Label11
							// line 857: elif target['ids']:
							πF.SetLineno(857)
						Label9:
							// line 858: naming = target['ids'][0]
							πF.SetLineno(858)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp007 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µtarget, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							µnaming = πTemp004
							goto Label11
						Label10:
							// line 862: naming = target['refid']
							πF.SetLineno(862)
							πTemp003 = ßrefid.ToObject()
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µtarget, πTemp003); πE != nil {
								continue
							}
							µnaming = πTemp004
							goto Label11
						Label11:
							// line 863: self.document.reporter.info(
							πF.SetLineno(863)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnaming, "naming"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Hyperlink target \"%s\" is not referenced.").ToObject(), µnaming); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"base_node", µtarget},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp009); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DanglingReferences").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDanglingReferences.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 868: class DanglingReferencesVisitor(nodes.SparseNodeVisitor):
			πF.SetLineno(868)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßSparseNodeVisitor, nil); πE != nil {
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
			_, πE = πg.NewCode("DanglingReferencesVisitor", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 870: def __init__(self, document, unknown_reference_resolvers):
					πF.SetLineno(870)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "document", Def: nil}
					πTemp002[2] = πg.Param{Name: "unknown_reference_resolvers", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πArgs[1]
						_ = µdocument
						var µunknown_reference_resolvers *πg.Object = πArgs[2]
						_ = µunknown_reference_resolvers
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
							// line 871: nodes.SparseNodeVisitor.__init__(self, document)
							πF.SetLineno(871)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSparseNodeVisitor, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 872: self.document = document
							πF.SetLineno(872)
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
							// line 873: self.unknown_reference_resolvers = unknown_reference_resolvers
							πF.SetLineno(873)
							if πE = πg.CheckLocal(πF, µunknown_reference_resolvers, "unknown_reference_resolvers"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µunknown_reference_resolvers); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßunknown_reference_resolvers, πTemp002); πE != nil {
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
					// line 875: def unknown_visit(self, node):
					πF.SetLineno(875)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("unknown_visit", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 876: pass
							πF.SetLineno(876)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßunknown_visit.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 878: def visit_reference(self, node):
					πF.SetLineno(878)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("visit_reference", "/usr/lib/python2.7/site-packages/docutils/transforms/references.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µrefname *πg.Object = πg.UnboundLocal
						_ = µrefname
						var µid *πg.Object = πg.UnboundLocal
						_ = µid
						var µresolver_function *πg.Object = πg.UnboundLocal
						_ = µresolver_function
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 16:
								goto Label16
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßresolved, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 879: if node.resolved or not node.hasattr('refname'):
							πF.SetLineno(879)
						Label2:
							// line 880: return
							πF.SetLineno(880)
							πR = πg.None
							continue
							goto Label3
						Label3:
							// line 881: refname = node['refname']
							πF.SetLineno(881)
							πTemp001 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							µrefname = πTemp003
							// line 882: id = self.document.nameids.get(refname)
							πF.SetLineno(882)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							πTemp004[0] = µrefname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µid = πTemp003
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µid == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 883: if id is None:
							πF.SetLineno(883)
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßunknown_reference_resolvers, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp002 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µresolver_function = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µresolver_function, "resolver_function"); πE != nil {
								continue
							}
							if πTemp003, πE = µresolver_function.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 885: if resolver_function(node):
							πF.SetLineno(885)
						Label10:
							// line 886: break
							πF.SetLineno(886)
							πTemp002 = true
							continue
							goto Label11
						Label11:
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrefname, "refname"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßnameids, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp006, µrefname); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							goto Label13
							// line 888: if refname in self.document.nameids:
							πF.SetLineno(888)
						Label12:
							// line 889: msg = self.document.reporter.error(
							πF.SetLineno(889)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Duplicate target name, cannot be used as a unique reference: \"%s\".").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßerror, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmsg = πTemp005
							goto Label14
						Label13:
							// line 893: msg = self.document.reporter.error(
							πF.SetLineno(893)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = ßrefname.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Unknown target name: \"%s\".").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"base_node", µnode},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßerror, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmsg = πTemp005
							goto Label14
						Label14:
							// line 896: msgid = self.document.set_id(msg)
							πF.SetLineno(896)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmsgid = πTemp003
							// line 897: prb = nodes.problematic(
							πF.SetLineno(897)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µprb = πTemp003
							// line 899: try:
							πF.SetLineno(899)
							πF.PushCheckpoint(16)
							// line 900: prbid = node['ids'][0]
							πF.SetLineno(900)
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp006 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µnode, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp003); πE != nil {
								continue
							}
							µprbid = πTemp005
							πF.PopCheckpoint()
							goto Label15
						Label16:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label17
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 901: except IndexError:
							πF.SetLineno(901)
						Label17:
							// line 902: prbid = self.document.set_id(prb)
							πF.SetLineno(902)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp004[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µprbid = πTemp003
							πF.RestoreExc(nil, nil)
							goto Label15
						Label15:
							// line 903: msg.add_backref(prbid)
							πF.SetLineno(903)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp004[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 904: node.replace_self(prb)
							πF.SetLineno(904)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp004[0] = µprb
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
						Label9:
							goto Label6
						Label5:
							// line 906: del node['refname']
							πF.SetLineno(906)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001 = ßrefname.ToObject()
							if πE = πg.DelItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							// line 907: node['refid'] = id
							πF.SetLineno(907)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp003 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µnode, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 908: self.document.ids[id].note_referenced_by(id=id)
							πF.SetLineno(908)
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"id", µid},
							}
							if πE = πg.CheckLocal(πF, µid, "id"); πE != nil {
								continue
							}
							πTemp001 = µid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßids, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßnote_referenced_by, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							// line 909: node.resolved = 1
							πF.SetLineno(909)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µnode, ßresolved, πTemp001); πE != nil {
								continue
							}
							goto Label6
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_reference.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 911: visit_footnote_reference = visit_citation_reference = visit_reference
					πF.SetLineno(911)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßvisit_reference); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_footnote_reference.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_citation_reference.ToObject(), πTemp005); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DanglingReferencesVisitor").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDanglingReferencesVisitor.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("references", Code)
}
