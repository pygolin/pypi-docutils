package parts

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßContents := πg.InternStr("Contents")
		ßContentsFilter := πg.InternStr("ContentsFilter")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßSectNum := πg.InternStr("SectNum")
		ßSkipDeparture := πg.InternStr("SkipDeparture")
		ßSkipNode := πg.InternStr("SkipNode")
		ßText := πg.InternStr("Text")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßTreeCopyVisitor := πg.InternStr("TreeCopyVisitor")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßalt := πg.InternStr("alt")
		ßappend := πg.InternStr("append")
		ßapply := πg.InternStr("apply")
		ßattributes := πg.InternStr("attributes")
		ßauto := πg.InternStr("auto")
		ßbacklinks := πg.InternStr("backlinks")
		ßbuild_contents := πg.InternStr("build_contents")
		ßbullet_list := πg.InternStr("bullet_list")
		ßchildren := πg.InternStr("children")
		ßclasses := πg.InternStr("classes")
		ßcopy_and_filter := πg.InternStr("copy_and_filter")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdepth := πg.InternStr("depth")
		ßdetails := πg.InternStr("details")
		ßdocument := πg.InternStr("document")
		ßentry := πg.InternStr("entry")
		ßgenerated := πg.InternStr("generated")
		ßget := πg.InternStr("get")
		ßget_entry_text := πg.InternStr("get_entry_text")
		ßget_tree_copy := πg.InternStr("get_tree_copy")
		ßhasattr := πg.InternStr("hasattr")
		ßids := πg.InternStr("ids")
		ßignore_node_but_process_children := πg.InternStr("ignore_node_but_process_children")
		ßinsert := πg.InternStr("insert")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlist_item := πg.InternStr("list_item")
		ßlocal := πg.InternStr("local")
		ßmaxdepth := πg.InternStr("maxdepth")
		ßmaxsize := πg.InternStr("maxsize")
		ßnext_node := πg.InternStr("next_node")
		ßnodes := πg.InternStr("nodes")
		ßparagraph := πg.InternStr("paragraph")
		ßparent := πg.InternStr("parent")
		ßprefix := πg.InternStr("prefix")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßrefid := πg.InternStr("refid")
		ßremove := πg.InternStr("remove")
		ßreplace_self := πg.InternStr("replace_self")
		ßsection := πg.InternStr("section")
		ßsectnum := πg.InternStr("sectnum")
		ßsectnum_depth := πg.InternStr("sectnum_depth")
		ßsectnum_prefix := πg.InternStr("sectnum_prefix")
		ßsectnum_start := πg.InternStr("sectnum_start")
		ßsectnum_suffix := πg.InternStr("sectnum_suffix")
		ßsectnum_xform := πg.InternStr("sectnum_xform")
		ßset_id := πg.InternStr("set_id")
		ßsettings := πg.InternStr("settings")
		ßstart := πg.InternStr("start")
		ßstartnode := πg.InternStr("startnode")
		ßstartvalue := πg.InternStr("startvalue")
		ßstr := πg.InternStr("str")
		ßsuffix := πg.InternStr("suffix")
		ßsys := πg.InternStr("sys")
		ßtoc_backlinks := πg.InternStr("toc_backlinks")
		ßtoc_id := πg.InternStr("toc_id")
		ßtop := πg.InternStr("top")
		ßupdate := πg.InternStr("update")
		ßupdate_section_numbers := πg.InternStr("update_section_numbers")
		ßuse_latex_toc := πg.InternStr("use_latex_toc")
		ßutils := πg.InternStr("utils")
		ßvisit_citation_reference := πg.InternStr("visit_citation_reference")
		ßvisit_footnote_reference := πg.InternStr("visit_footnote_reference")
		ßvisit_image := πg.InternStr("visit_image")
		ßvisit_problematic := πg.InternStr("visit_problematic")
		ßvisit_reference := πg.InternStr("visit_reference")
		ßvisit_target := πg.InternStr("visit_target")
		ßwalkabout := πg.InternStr("walkabout")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nTransforms related to document parts.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
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
			// line 13: import sys
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: from docutils import nodes, utils
			πF.SetLineno(14)
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
			// line 15: from docutils.transforms import TransformError, Transform
			πF.SetLineno(15)
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
			// line 18: class SectNum(Transform):
			πF.SetLineno(18)
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
			_, πE = πg.NewCode("SectNum", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 20: """
					πF.SetLineno(20)
					// line 20: """
					πF.SetLineno(20)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Automatically assigns numbers to the titles of document sections.\n\n    It is possible to limit the maximum section level for which the numbers\n    are added.  For those sections that are auto-numbered, the \"autonum\"\n    attribute is set, informing the contents table generator that a different\n    form of the TOC should be used.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 29: default_priority = 710
					πF.SetLineno(29)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(710).ToObject()); πE != nil {
						continue
					}
					// line 30: """Should be applied before `Contents`."""
					πF.SetLineno(30)
					// line 32: def apply(self):
					πF.SetLineno(32)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 33: self.maxdepth = self.startnode.details.get('depth', None)
							πF.SetLineno(33)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßdepth.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßmaxdepth, πTemp002); πE != nil {
								continue
							}
							// line 34: self.startvalue = self.startnode.details.get('start', 1)
							πF.SetLineno(34)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßstart.ToObject()
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßstartvalue, πTemp002); πE != nil {
								continue
							}
							// line 35: self.prefix = self.startnode.details.get('prefix', '')
							πF.SetLineno(35)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßprefix.ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßprefix, πTemp002); πE != nil {
								continue
							}
							// line 36: self.suffix = self.startnode.details.get('suffix', '')
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßsuffix.ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßsuffix, πTemp002); πE != nil {
								continue
							}
							// line 37: self.startnode.parent.remove(self.startnode)
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsectnum_xform, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 38: if self.document.settings.sectnum_xform:
							πF.SetLineno(38)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxdepth, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp005).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 39: if self.maxdepth is None:
							πF.SetLineno(39)
						Label4:
							// line 40: self.maxdepth = sys.maxsize
							πF.SetLineno(40)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmaxsize, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmaxdepth, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 41: self.update_section_numbers(self.document)
							πF.SetLineno(41)
							πTemp001 = πF.MakeArgs(1)
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßupdate_section_numbers, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 43: self.document.settings.sectnum_depth = self.maxdepth
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxdepth, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßsectnum_depth, πTemp003); πE != nil {
								continue
							}
							// line 44: self.document.settings.sectnum_start = self.startvalue
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartvalue, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßsectnum_start, πTemp003); πE != nil {
								continue
							}
							// line 45: self.document.settings.sectnum_prefix = self.prefix
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprefix, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßsectnum_prefix, πTemp003); πE != nil {
								continue
							}
							// line 46: self.document.settings.sectnum_suffix = self.suffix
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsuffix, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp006, ßsectnum_suffix, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 48: def update_section_numbers(self, node, prefix=(), depth=0):
					πF.SetLineno(48)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002[2] = πg.Param{Name: "prefix", Def: πTemp004}
					πTemp002[3] = πg.Param{Name: "depth", Def: πg.NewInt(0).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("update_section_numbers", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µprefix *πg.Object = πArgs[2]
						_ = µprefix
						var µdepth *πg.Object = πArgs[3]
						_ = µdepth
						var µsectnum *πg.Object = πg.UnboundLocal
						_ = µsectnum
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var µnumbers *πg.Object = πg.UnboundLocal
						_ = µnumbers
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µgenerated *πg.Object = πg.UnboundLocal
						_ = µgenerated
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 πg.KWArgs
						_ = πTemp012
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
							// line 49: depth += 1
							πF.SetLineno(49)
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µdepth, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µdepth = πTemp001
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µprefix); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 50: if prefix:
							πF.SetLineno(50)
						Label1:
							// line 51: sectnum = 1
							πF.SetLineno(51)
							µsectnum = πg.NewInt(1).ToObject()
							goto Label3
						Label2:
							// line 53: sectnum = self.startvalue
							πF.SetLineno(53)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartvalue, nil); πE != nil {
								continue
							}
							µsectnum = πTemp001
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µnode); πE != nil {
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
								µchild = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp005[0] = µchild
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßsection, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 55: if isinstance(child, nodes.section):
							πF.SetLineno(55)
						Label7:
							// line 56: numbers = prefix + (str(sectnum),)
							πF.SetLineno(56)
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsectnum, "sectnum"); πE != nil {
								continue
							}
							πTemp005[0] = µsectnum
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp006 = πg.NewTuple1(πTemp008).ToObject()
							if πTemp004, πE = πg.Add(πF, µprefix, πTemp006); πE != nil {
								continue
							}
							µnumbers = πTemp004
							// line 57: title = child[0]
							πF.SetLineno(57)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µchild, πTemp004); πE != nil {
								continue
							}
							µtitle = πTemp006
							// line 59: generated = nodes.generated(
							πF.SetLineno(59)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßprefix, nil); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnumbers, "numbers"); πE != nil {
								continue
							}
							πTemp009[0] = µnumbers
							if πTemp010, πE = πg.GetAttr(πF, πg.NewStr(".").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp007, πE = πg.Add(πF, πTemp008, πTemp011); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßsuffix, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mul(πF, πg.NewUnicode("\xc2\xa0").ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = ßsectnum.ToObject()
							πTemp004 = πg.NewList(πTemp009...).ToObject()
							πTemp012 = πg.KWArgs{
								{"classes", πTemp004},
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßgenerated, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp005, πTemp012); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µgenerated = πTemp004
							// line 63: title.insert(0, generated)
							πF.SetLineno(63)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µgenerated, "generated"); πE != nil {
								continue
							}
							πTemp005[1] = µgenerated
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtitle, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 64: title['auto'] = 1
							πF.SetLineno(64)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp006 = ßauto.ToObject()
							if πE = πg.SetItem(πF, µtitle, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmaxdepth, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µdepth, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 65: if depth < self.maxdepth:
							πF.SetLineno(65)
						Label9:
							// line 66: self.update_section_numbers(child, numbers, depth)
							πF.SetLineno(66)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp005[0] = µchild
							if πE = πg.CheckLocal(πF, µnumbers, "numbers"); πE != nil {
								continue
							}
							πTemp005[1] = µnumbers
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							πTemp005[2] = µdepth
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßupdate_section_numbers, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label10
						Label10:
							// line 67: sectnum += 1
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µsectnum, "sectnum"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µsectnum, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µsectnum = πTemp004
							goto Label8
						Label8:
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
					if πE = πClass.SetItem(πF, ßupdate_section_numbers.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SectNum").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSectNum.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 70: class Contents(Transform):
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
			_, πE = πg.NewCode("Contents", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    This transform generates a table of contents from the entire document tree\n    or from a single branch.  It locates \"section\" elements and builds them\n    into a nested bullet list, which is placed within a \"topic\" created by the\n    contents directive.  A title is either explicitly specified, taken from\n    the appropriate language module, or omitted (local table of contents).\n    The depth may be specified.  Two-way references between the table of\n    contents and section titles are generated (requires Writer support).\n\n    This transform requires a startnode, which contains generation\n    options and provides the location for the generated table of contents (the\n    startnode is replaced by the table of contents \"topic\").\n    ").ToObject()); πE != nil {
						continue
					}
					// line 86: default_priority = 720
					πF.SetLineno(86)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(720).ToObject()); πE != nil {
						continue
					}
					// line 88: def apply(self):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtoc_by_writer *πg.Object = πg.UnboundLocal
						_ = µtoc_by_writer
						var µdetails *πg.Object = πg.UnboundLocal
						_ = µdetails
						var µstartnode *πg.Object = πg.UnboundLocal
						_ = µstartnode
						var µcontents *πg.Object = πg.UnboundLocal
						_ = µcontents
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
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
							case 8:
								goto Label8
							case 2:
								goto Label2
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 89: try: # let the writer (or output software) build the contents list?
							πF.SetLineno(89)
							πF.PushCheckpoint(2)
							// line 90: toc_by_writer = self.document.settings.use_latex_toc
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßuse_latex_toc, nil); πE != nil {
								continue
							}
							µtoc_by_writer = πTemp001
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
							continue
							// line 91: except AttributeError:
							πF.SetLineno(91)
						Label3:
							// line 92: toc_by_writer = False
							πF.SetLineno(92)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µtoc_by_writer = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 93: details = self.startnode.details
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdetails, nil); πE != nil {
								continue
							}
							µdetails = πTemp002
							if πE = πg.CheckLocal(πF, µdetails, "details"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µdetails, ßlocal.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 94: if 'local' in details:
							πF.SetLineno(94)
						Label4:
							// line 95: startnode = self.startnode.parent.parent
							πF.SetLineno(95)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßparent, nil); πE != nil {
								continue
							}
							µstartnode = πTemp001
							// line 96: while not (isinstance(startnode, nodes.section)
							πF.SetLineno(96)
							πF.PushCheckpoint(8)
							πTemp005 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label9
							}
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							πTemp008[0] = µstartnode
							if πTemp009, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßsection, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp010
							if πTemp009, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πTemp010
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							πTemp008[0] = µstartnode
							if πTemp009, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßdocument, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp010
							if πTemp009, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πTemp010
						Label10:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 99: startnode = startnode.parent
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstartnode, ßparent, nil); πE != nil {
								continue
							}
							µstartnode = πTemp001
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							goto Label6
						Label5:
							// line 101: startnode = self.document
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							µstartnode = πTemp001
							goto Label6
						Label6:
							// line 102: self.toc_id = self.startnode.parent['ids'][0]
							πF.SetLineno(102)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp009 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßparent, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp012, πTemp009); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp010, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtoc_id, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdetails, "details"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µdetails, ßbacklinks.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 103: if 'backlinks' in details:
							πF.SetLineno(103)
						Label11:
							// line 104: self.backlinks = details['backlinks']
							πF.SetLineno(104)
							πTemp001 = ßbacklinks.ToObject()
							if πE = πg.CheckLocal(πF, µdetails, "details"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µdetails, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbacklinks, πTemp001); πE != nil {
								continue
							}
							goto Label13
						Label12:
							// line 106: self.backlinks = self.document.settings.toc_backlinks
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßtoc_backlinks, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbacklinks, πTemp002); πE != nil {
								continue
							}
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µtoc_by_writer, "toc_by_writer"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µtoc_by_writer); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							goto Label15
							// line 107: if toc_by_writer:
							πF.SetLineno(107)
						Label14:
							// line 109: self.startnode.parent.attributes.update(details)
							πF.SetLineno(109)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdetails, "details"); πE != nil {
								continue
							}
							πTemp008[0] = µdetails
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßattributes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 110: self.startnode.parent.remove(self.startnode)
							πF.SetLineno(110)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label16
						Label15:
							// line 112: contents = self.build_contents(startnode)
							πF.SetLineno(112)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstartnode, "startnode"); πE != nil {
								continue
							}
							πTemp008[0] = µstartnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbuild_contents, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µcontents = πTemp002
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
								continue
							}
							πTemp008[0] = µcontents
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label17
							}
							goto Label18
							// line 113: if len(contents):
							πF.SetLineno(113)
						Label17:
							// line 114: self.startnode.replace_self(contents)
							πF.SetLineno(114)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
								continue
							}
							πTemp008[0] = µcontents
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label19
						Label18:
							// line 116: self.startnode.parent.parent.remove(self.startnode.parent)
							πF.SetLineno(116)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparent, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßparent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label19
						Label19:
							goto Label16
						Label16:
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
					// line 118: def build_contents(self, node, level=0):
					πF.SetLineno(118)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp002[2] = πg.Param{Name: "level", Def: πg.NewInt(0).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("build_contents", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µlevel *πg.Object = πArgs[2]
						_ = µlevel
						var µsections *πg.Object = πg.UnboundLocal
						_ = µsections
						var µentries *πg.Object = πg.UnboundLocal
						_ = µentries
						var µautonum *πg.Object = πg.UnboundLocal
						_ = µautonum
						var µdepth *πg.Object = πg.UnboundLocal
						_ = µdepth
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µauto *πg.Object = πg.UnboundLocal
						_ = µauto
						var µentrytext *πg.Object = πg.UnboundLocal
						_ = µentrytext
						var µreference *πg.Object = πg.UnboundLocal
						_ = µreference
						var µref_id *πg.Object = πg.UnboundLocal
						_ = µref_id
						var µentry *πg.Object = πg.UnboundLocal
						_ = µentry
						var µitem *πg.Object = πg.UnboundLocal
						_ = µitem
						var µsubsects *πg.Object = πg.UnboundLocal
						_ = µsubsects
						var µcontents *πg.Object = πg.UnboundLocal
						_ = µcontents
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
						var πTemp011 πg.KWArgs
						_ = πTemp011
						var πTemp012 bool
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
							default:
								panic("unexpected function state")
							}
							// line 119: level += 1
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µlevel = πTemp001
							// line 120: sections = [sect for sect in node if isinstance(sect, nodes.section)]
							πF.SetLineno(120)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µsect *πg.Object = πg.UnboundLocal
								_ = µsect
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
										case 6:
											goto Label6
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µnode); πE != nil {
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
											µsect = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µsect, "sect"); πE != nil {
											continue
										}
										πTemp005[0] = µsect
										if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßsection, nil); πE != nil {
											continue
										}
										πTemp005[1] = πTemp006
										if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 120: sections = [sect for sect in node if isinstance(sect, nodes.section)]
										πF.SetLineno(120)
									Label4:
										// line 120: sections = [sect for sect in node if isinstance(sect, nodes.section)]
										πF.SetLineno(120)
										if πE = πg.CheckLocal(πF, µsect, "sect"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µsect, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µsections = πTemp001
							// line 121: entries = []
							πF.SetLineno(121)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µentries = πTemp001
							// line 122: autonum = 0
							πF.SetLineno(122)
							µautonum = πg.NewInt(0).ToObject()
							// line 123: depth = self.startnode.details.get('depth', sys.maxsize)
							πF.SetLineno(123)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ßdepth.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmaxsize, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßdetails, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µdepth = πTemp004
							if πE = πg.CheckLocal(πF, µsections, "sections"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µsections); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µsection = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 125: title = section[0]
							πF.SetLineno(125)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µsection, πTemp004); πE != nil {
								continue
							}
							µtitle = πTemp008
							// line 126: auto = title.get('auto')    # May be set by SectNum.
							πF.SetLineno(126)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßauto.ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µtitle, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µauto = πTemp008
							// line 127: entrytext = self.copy_and_filter(title)
							πF.SetLineno(127)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp005[0] = µtitle
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcopy_and_filter, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µentrytext = πTemp008
							// line 128: reference = nodes.reference('', '', refid=section['ids'][0],
							πF.SetLineno(128)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ß.ToObject()
							πTemp005[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µentrytext, "entrytext"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							πTemp009 = ßids.ToObject()
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µsection, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp004); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"refid", πTemp008},
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßreference, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp008, πTemp005, µentrytext, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µreference = πTemp004
							// line 130: ref_id = self.document.set_id(reference,
							πF.SetLineno(130)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreference, "reference"); πE != nil {
								continue
							}
							πTemp005[0] = µreference
							πTemp011 = πg.KWArgs{
								{"suggested_prefix", πg.NewStr("toc-entry").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp005, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µref_id = πTemp004
							// line 132: entry = nodes.paragraph('', '', reference)
							πF.SetLineno(132)
							πTemp005 = πF.MakeArgs(3)
							πTemp005[0] = ß.ToObject()
							πTemp005[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µreference, "reference"); πE != nil {
								continue
							}
							πTemp005[2] = µreference
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µentry = πTemp004
							// line 133: item = nodes.list_item('', entry)
							πF.SetLineno(133)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
								continue
							}
							πTemp005[1] = µentry
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßlist_item, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µitem = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßbacklinks, nil); πE != nil {
								continue
							}
							πTemp010 = πg.NewTuple2(ßentry.ToObject(), ßtop.ToObject()).ToObject()
							if πTemp012, πE = πg.Contains(πF, πTemp010, πTemp009); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp012).ToObject()
							πTemp004 = πTemp008
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label4
							}
							πTemp005 = πF.MakeArgs(1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßreference, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp010
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µtitle, ßnext_node, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(πTemp010 == πTemp009).ToObject()
							πTemp004 = πTemp008
						Label4:
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							goto Label6
							// line 134: if ( self.backlinks in ('entry', 'top')
							πF.SetLineno(134)
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßbacklinks, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp008, ßentry.ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßbacklinks, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp008, ßtop.ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 136: if self.backlinks == 'entry':
							πF.SetLineno(136)
						Label7:
							// line 137: title['refid'] = ref_id
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µref_id, "ref_id"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µref_id); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp008 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µtitle, πTemp008, πTemp004); πE != nil {
								continue
							}
							goto Label9
							// line 138: elif self.backlinks == 'top':
							πF.SetLineno(138)
						Label8:
							// line 139: title['refid'] = self.toc_id
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtoc_id, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp009 = ßrefid.ToObject()
							if πE = πg.SetItem(πF, µtitle, πTemp009, πTemp008); πE != nil {
								continue
							}
							goto Label9
						Label9:
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µlevel, µdepth); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 140: if level < depth:
							πF.SetLineno(140)
						Label10:
							// line 141: subsects = self.build_contents(section, level)
							πF.SetLineno(141)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp005[0] = µsection
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp005[1] = µlevel
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßbuild_contents, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µsubsects = πTemp008
							// line 142: item += subsects
							πF.SetLineno(142)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubsects, "subsects"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µitem, µsubsects); πE != nil {
								continue
							}
							µitem = πTemp004
							goto Label11
						Label11:
							// line 143: entries.append(item)
							πF.SetLineno(143)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp005[0] = µitem
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µentries, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µentries); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 144: if entries:
							πF.SetLineno(144)
						Label12:
							// line 145: contents = nodes.bullet_list('', *entries)
							πF.SetLineno(145)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µentries, "entries"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßbullet_list, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp004, πTemp005, µentries, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µcontents = πTemp001
							if πE = πg.CheckLocal(πF, µauto, "auto"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µauto); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label15
							}
							goto Label16
							// line 146: if auto:
							πF.SetLineno(146)
						Label15:
							// line 147: contents['classes'].append('auto-toc')
							πF.SetLineno(147)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("auto-toc").ToObject()
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µcontents, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label16
						Label16:
							// line 148: return contents
							πF.SetLineno(148)
							if πE = πg.CheckLocal(πF, µcontents, "contents"); πE != nil {
								continue
							}
							πR = µcontents
							continue
							goto Label14
						Label13:
							// line 150: return []
							πF.SetLineno(150)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πR = πTemp001
							continue
							goto Label14
						Label14:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßbuild_contents.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 152: def copy_and_filter(self, node):
					πF.SetLineno(152)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("copy_and_filter", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
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
							// line 153: """Return a copy of a title, with references, images, etc. removed."""
							πF.SetLineno(153)
							// line 154: visitor = ContentsFilter(self.document)
							πF.SetLineno(154)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßContentsFilter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvisitor = πTemp003
							// line 155: node.walkabout(visitor)
							πF.SetLineno(155)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp001[0] = µvisitor
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßwalkabout, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 156: return visitor.get_entry_text()
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µvisitor, ßget_entry_text, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßcopy_and_filter.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 153: """Return a copy of a title, with references, images, etc. removed."""
					πF.SetLineno(153)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a copy of a title, with references, images, etc. removed.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßcopy_and_filter); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
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
			// line 159: class ContentsFilter(nodes.TreeCopyVisitor):
			πF.SetLineno(159)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTreeCopyVisitor, nil); πE != nil {
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
			_, πE = πg.NewCode("ContentsFilter", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 161: def get_entry_text(self):
					πF.SetLineno(161)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_entry_text", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							// line 162: return self.get_tree_copy().children
							πF.SetLineno(162)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_tree_copy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßchildren, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_entry_text.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 164: def visit_citation_reference(self, node):
					πF.SetLineno(164)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("visit_citation_reference", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 165: raise nodes.SkipNode
							πF.SetLineno(165)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_citation_reference.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 167: def visit_footnote_reference(self, node):
					πF.SetLineno(167)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("visit_footnote_reference", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 168: raise nodes.SkipNode
							πF.SetLineno(168)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_footnote_reference.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 170: def visit_image(self, node):
					πF.SetLineno(170)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("visit_image", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßalt.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßhasattr, nil); πE != nil {
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
							// line 171: if node.hasattr('alt'):
							πF.SetLineno(171)
						Label1:
							// line 172: self.parent.append(nodes.Text(node['alt']))
							πF.SetLineno(172)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp002 = ßalt.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßText, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 173: raise nodes.SkipNode
							πF.SetLineno(173)
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
					if πE = πClass.SetItem(πF, ßvisit_image.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 175: def ignore_node_but_process_children(self, node):
					πF.SetLineno(175)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("ignore_node_but_process_children", "/usr/lib/python2.7/site-packages/docutils/transforms/parts.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipDeparture, nil); πE != nil {
								continue
							}
							// line 176: raise nodes.SkipDeparture
							πF.SetLineno(176)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßignore_node_but_process_children.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 178: visit_problematic = ignore_node_but_process_children
					πF.SetLineno(178)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node_but_process_children); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_problematic.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 179: visit_reference = ignore_node_but_process_children
					πF.SetLineno(179)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node_but_process_children); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_reference.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 180: visit_target = ignore_node_but_process_children
					πF.SetLineno(180)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßignore_node_but_process_children); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßvisit_target.ToObject(), πTemp007); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ContentsFilter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßContentsFilter.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("parts", Code)
}
