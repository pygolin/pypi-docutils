package frontmatter

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßDecorative := πg.InternStr("Decorative")
		ßDocInfo := πg.InternStr("DocInfo")
		ßDocTitle := πg.InternStr("DocTitle")
		ßElement := πg.InternStr("Element")
		ßFalse := πg.InternStr("False")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßNone := πg.InternStr("None")
		ßPreBibliographic := πg.InternStr("PreBibliographic")
		ßSectionSubTitle := πg.InternStr("SectionSubTitle")
		ßText := πg.InternStr("Text")
		ßTextElement := πg.InternStr("TextElement")
		ßTitlePromoter := πg.InternStr("TitlePromoter")
		ßTitular := πg.InternStr("Titular")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_traverse := πg.InternStr("_traverse")
		ßabstract := πg.InternStr("abstract")
		ßaddress := πg.InternStr("address")
		ßappend := πg.InternStr("append")
		ßapply := πg.InternStr("apply")
		ßastext := πg.InternStr("astext")
		ßauthor := πg.InternStr("author")
		ßauthor_separators := πg.InternStr("author_separators")
		ßauthors := πg.InternStr("authors")
		ßauthors_from_bullet_list := πg.InternStr("authors_from_bullet_list")
		ßauthors_from_one_paragraph := πg.InternStr("authors_from_one_paragraph")
		ßauthors_from_paragraphs := πg.InternStr("authors_from_paragraphs")
		ßbiblio_nodes := πg.InternStr("biblio_nodes")
		ßbibliographic_fields := πg.InternStr("bibliographic_fields")
		ßbullet_list := πg.InternStr("bullet_list")
		ßcandidate_index := πg.InternStr("candidate_index")
		ßcheck_compound_biblio_field := πg.InternStr("check_compound_biblio_field")
		ßcheck_empty_biblio_field := πg.InternStr("check_empty_biblio_field")
		ßchildren := πg.InternStr("children")
		ßclasses := πg.InternStr("classes")
		ßclean_rcs_keywords := πg.InternStr("clean_rcs_keywords")
		ßcomment := πg.InternStr("comment")
		ßcompile := πg.InternStr("compile")
		ßcontact := πg.InternStr("contact")
		ßcopyright := πg.InternStr("copyright")
		ßdate := πg.InternStr("date")
		ßdedication := πg.InternStr("dedication")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdocinfo := πg.InternStr("docinfo")
		ßdocinfo_xform := πg.InternStr("docinfo_xform")
		ßdoctitle_xform := πg.InternStr("doctitle_xform")
		ßdocument := πg.InternStr("document")
		ßextract_authors := πg.InternStr("extract_authors")
		ßextract_bibliographic := πg.InternStr("extract_bibliographic")
		ßfield_list := πg.InternStr("field_list")
		ßfirst_child_not_matching_class := πg.InternStr("first_child_not_matching_class")
		ßfully_normalize_name := πg.InternStr("fully_normalize_name")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßjoin := πg.InternStr("join")
		ßlabels := πg.InternStr("labels")
		ßlanguage := πg.InternStr("language")
		ßlen := πg.InternStr("len")
		ßmake_id := πg.InternStr("make_id")
		ßnodes := πg.InternStr("nodes")
		ßorganization := πg.InternStr("organization")
		ßparagraph := πg.InternStr("paragraph")
		ßpromote_subtitle := πg.InternStr("promote_subtitle")
		ßpromote_title := πg.InternStr("promote_title")
		ßrawsource := πg.InternStr("rawsource")
		ßrcs_keyword_substitutions := πg.InternStr("rcs_keyword_substitutions")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrevision := πg.InternStr("revision")
		ßsection := πg.InternStr("section")
		ßsectsubtitle_xform := πg.InternStr("sectsubtitle_xform")
		ßset_metadata := πg.InternStr("set_metadata")
		ßsettings := πg.InternStr("settings")
		ßsplit := πg.InternStr("split")
		ßstatus := πg.InternStr("status")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsubtitle := πg.InternStr("subtitle")
		ßsys := πg.InternStr("sys")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßtraverse := πg.InternStr("traverse")
		ßunescape := πg.InternStr("unescape")
		ßunicode := πg.InternStr("unicode")
		ßupdate_all_atts_concatenating := πg.InternStr("update_all_atts_concatenating")
		ßutils := πg.InternStr("utils")
		ßversion := πg.InternStr("version")
		ßversion_info := πg.InternStr("version_info")
		ßwarning := πg.InternStr("warning")
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
		var πTemp006 *πg.Dict
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nTransforms related to the front matter of a document or a section\n(information found before the main text):\n\n- `DocTitle`: Used to transform a lone top level section's title to\n  the document title, promote a remaining lone top-level section's\n  title to the document subtitle, and determine the document's title\n  metadata (document['title']) based on the document title and/or the\n  \"title\" setting.\n\n- `SectionSubTitle`: Used to transform a lone subsection into a\n  subtitle.\n\n- `DocInfo`: Used to transform a bibliographic field list into docinfo\n  elements.\n").ToObject()); πE != nil {
				continue
			}
			// line 22: __docformat__ = 'reStructuredText'
			πF.SetLineno(22)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 24: import re
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: import sys
			πF.SetLineno(25)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: from docutils import nodes, utils
			πF.SetLineno(27)
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
			// line 28: from docutils.transforms import TransformError, Transform
			πF.SetLineno(28)
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
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 31: if sys.version_info >= (3, 0):
			πF.SetLineno(31)
		Label1:
			// line 32: unicode = str  # noqa
			πF.SetLineno(32)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 35: class TitlePromoter(Transform):
			πF.SetLineno(35)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TitlePromoter", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 37: """
					πF.SetLineno(37)
					// line 37: """
					πF.SetLineno(37)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class for DocTitle and SectionSubTitle transforms.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 41: def promote_title(self, node):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("promote_title", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 42: """
							πF.SetLineno(42)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßElement, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 59: if not isinstance(node, nodes.Element):
							πF.SetLineno(59)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("node must be of Element-derived type.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 60: raise TypeError('node must be of Element-derived type.')
							πF.SetLineno(60)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 63: assert not (len(node) and isinstance(node[0], nodes.title))
							πF.SetLineno(63)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp006
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßtitle, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp006
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 64: section, index = self.candidate_index(node)
							πF.SetLineno(64)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcandidate_index, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µsection = πTemp001
							µindex = πTemp004
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µindex == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 65: if index is None:
							πF.SetLineno(65)
						Label4:
							// line 66: return False
							πF.SetLineno(66)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							// line 73: node.update_all_atts_concatenating(section, replace=True, and_source=True)
							πF.SetLineno(73)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp002[0] = µsection
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"replace", πTemp001},
								{"and_source", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßupdate_all_atts_concatenating, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 76: node[:] = (section[:1]        # section title
							πF.SetLineno(76)
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µsection, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µindex, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µsection, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µnode, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 80: assert isinstance(node[0], nodes.title)
							πF.SetLineno(80)
							πTemp002 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 81: return True
							πF.SetLineno(81)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpromote_title.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 42: """
					πF.SetLineno(42)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Transform the following tree::\n\n            <node>\n                <section>\n                    <title>\n                    ...\n\n        into ::\n\n            <node>\n                <title>\n                ...\n\n        `node` is normally a document.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßpromote_title); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 83: def promote_subtitle(self, node):
					πF.SetLineno(83)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("promote_subtitle", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µsubsection *πg.Object = πg.UnboundLocal
						_ = µsubsection
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µsubtitle *πg.Object = πg.UnboundLocal
						_ = µsubtitle
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
						var πTemp006 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							// line 84: """
							πF.SetLineno(84)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßElement, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 101: if not isinstance(node, nodes.Element):
							πF.SetLineno(101)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("node must be of Element-derived type.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 102: raise TypeError('node must be of Element-derived type.')
							πF.SetLineno(102)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 104: subsection, index = self.candidate_index(node)
							πF.SetLineno(104)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcandidate_index, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µsubsection = πTemp001
							µindex = πTemp004
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µindex == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 105: if index is None:
							πF.SetLineno(105)
						Label3:
							// line 106: return False
							πF.SetLineno(106)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 107: subtitle = nodes.subtitle()
							πF.SetLineno(107)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsubtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsubtitle = πTemp001
							// line 114: subtitle.update_all_atts_concatenating(subsection, replace=True, and_source=True)
							πF.SetLineno(114)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsubsection, "subsection"); πE != nil {
								continue
							}
							πTemp002[0] = µsubsection
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"replace", πTemp001},
								{"and_source", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µsubtitle, "subtitle"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsubtitle, ßupdate_all_atts_concatenating, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 118: subtitle[:] = subsection[0][:]
							πF.SetLineno(118)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µsubsection, "subsection"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µsubsection, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubtitle, "subtitle"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µsubtitle, πTemp004, πTemp001); πE != nil {
								continue
							}
							// line 119: node[:] = (node[:1]       # title
							πF.SetLineno(119)
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp007); πE != nil {
								continue
							}
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µsubtitle, "subtitle"); πE != nil {
								continue
							}
							πTemp002[0] = µsubtitle
							πTemp007 = πg.NewList(πTemp002...).ToObject()
							if πTemp004, πE = πg.Add(πF, πTemp008, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), µindex, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µnode, πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsubsection, "subsection"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µsubsection, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µnode, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 125: return True
							πF.SetLineno(125)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpromote_subtitle.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 84: """
					πF.SetLineno(84)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Transform the following node tree::\n\n            <node>\n                <title>\n                <section>\n                    <title>\n                    ...\n\n        into ::\n\n            <node>\n                <title>\n                <subtitle>\n                ...\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßpromote_subtitle); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 127: def candidate_index(self, node):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("candidate_index", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
						var πTemp007 bool
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
							// line 128: """
							πF.SetLineno(128)
							// line 133: index = node.first_child_not_matching_class(
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßPreBibliographic, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnode, ßfirst_child_not_matching_class, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp003
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µindex == πTemp005).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp006, πTemp005); πE != nil {
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
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp005 = µindex
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µnode, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsection, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 135: if (index is None or len(node) > (index + 1)
							πF.SetLineno(135)
						Label2:
							// line 137: return None, None
							πF.SetLineno(137)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
							πR = πTemp002
							continue
							goto Label4
						Label3:
							// line 139: return node[index], index
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp003 = µindex
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µnode, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, µindex).ToObject()
							πR = πTemp002
							continue
							goto Label4
						Label4:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcandidate_index.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 128: """
					πF.SetLineno(128)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Find and return the promotion candidate and its index.\n\n        Return (None, None) if no valid candidate was found.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßcandidate_index); πE != nil {
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
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TitlePromoter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTitlePromoter.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 142: class DocTitle(TitlePromoter):
			πF.SetLineno(142)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTitlePromoter); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DocTitle", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 144: """
					πF.SetLineno(144)
					// line 144: """
					πF.SetLineno(144)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    In reStructuredText_, there is no way to specify a document title\n    and subtitle explicitly. Instead, we can supply the document title\n    (and possibly the subtitle as well) implicitly, and use this\n    two-step transform to \"raise\" or \"promote\" the title(s) (and their\n    corresponding section contents) to the document level.\n\n    1. If the document contains a single top-level section as its\n       first non-comment element, the top-level section's title\n       becomes the document's title, and the top-level section's\n       contents become the document's immediate contents. The lone\n       top-level section header must be the first non-comment element\n       in the document.\n\n       For example, take this input text::\n\n           =================\n            Top-Level Title\n           =================\n\n           A paragraph.\n\n       Once parsed, it looks like this::\n\n           <document>\n               <section names=\"top-level title\">\n                   <title>\n                       Top-Level Title\n                   <paragraph>\n                       A paragraph.\n\n       After running the DocTitle transform, we have::\n\n           <document names=\"top-level title\">\n               <title>\n                   Top-Level Title\n               <paragraph>\n                   A paragraph.\n\n    2. If step 1 successfully determines the document title, we\n       continue by checking for a subtitle.\n\n       If the lone top-level section itself contains a single\n       second-level section as its first non-comment element, that\n       section's title is promoted to the document's subtitle, and\n       that section's contents become the document's immediate\n       contents. Given this input text::\n\n           =================\n            Top-Level Title\n           =================\n\n           Second-Level Title\n           ~~~~~~~~~~~~~~~~~~\n\n           A paragraph.\n\n       After parsing and running the Section Promotion transform, the\n       result is::\n\n           <document names=\"top-level title\">\n               <title>\n                   Top-Level Title\n               <subtitle names=\"second-level title\">\n                   Second-Level Title\n               <paragraph>\n                   A paragraph.\n\n       (Note that the implicit hyperlink target generated by the\n       \"Second-Level Title\" is preserved on the \"subtitle\" element\n       itself.)\n\n    Any comment elements occurring before the document title or\n    subtitle are accumulated and inserted as the first body elements\n    after the title(s).\n\n    This transform also sets the document's metadata title\n    (document['title']).\n\n    .. _reStructuredText: http://docutils.sf.net/rst.html\n    ").ToObject()); πE != nil {
						continue
					}
					// line 226: default_priority = 320
					πF.SetLineno(226)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(320).ToObject()); πE != nil {
						continue
					}
					// line 228: def set_metadata(self):
					πF.SetLineno(228)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("set_metadata", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 229: """
							πF.SetLineno(229)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßhasattr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 237: if not self.document.hasattr('title'):
							πF.SetLineno(237)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 != πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtitle, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 238: if self.document.settings.title is not None:
							πF.SetLineno(238)
						Label3:
							// line 239: self.document['title'] = self.document.settings.title
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßtitle, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp006 = ßtitle.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp006, πTemp003); πE != nil {
								continue
							}
							goto Label6
							// line 240: elif len(self.document) and isinstance(self.document[0], nodes.title):
							πF.SetLineno(240)
						Label5:
							// line 241: self.document['title'] = self.document[0].astext()
							πF.SetLineno(241)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp006 = ßtitle.ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp006, πTemp001); πE != nil {
								continue
							}
							goto Label6
						Label6:
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßset_metadata.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 229: """
					πF.SetLineno(229)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Set document['title'] metadata title from the following\n        sources, listed in order of priority:\n\n        * Existing document['title'] attribute.\n        * \"title\" setting.\n        * Document title node (as promoted by promote_title).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßset_metadata); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 243: def apply(self):
					πF.SetLineno(243)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßdoctitle_xform.ToObject()
							πTemp001[2] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
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
							// line 244: if getattr(self.document.settings, 'doctitle_xform', 1):
							πF.SetLineno(244)
						Label1:
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpromote_title, nil); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 246: if self.promote_title(self.document):
							πF.SetLineno(246)
						Label3:
							// line 249: self.promote_subtitle(self.document)
							πF.SetLineno(249)
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpromote_subtitle, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 251: self.set_metadata()
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_metadata, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DocTitle").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDocTitle.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 254: class SectionSubTitle(TitlePromoter):
			πF.SetLineno(254)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTitlePromoter); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SectionSubTitle", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 256: """
					πF.SetLineno(256)
					// line 256: """
					πF.SetLineno(256)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    This works like document subtitles, but for sections.  For example, ::\n\n        <section>\n            <title>\n                Title\n            <section>\n                <title>\n                    Subtitle\n                ...\n\n    is transformed into ::\n\n        <section>\n            <title>\n                Title\n            <subtitle>\n                Subtitle\n            ...\n\n    For details refer to the docstring of DocTitle.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 279: default_priority = 350
					πF.SetLineno(279)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(350).ToObject()); πE != nil {
						continue
					}
					// line 281: def apply(self):
					πF.SetLineno(281)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = ßsectsubtitle_xform.ToObject()
							πTemp002[2] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 282: if not getattr(self.document.settings, 'sectsubtitle_xform', 1):
							πF.SetLineno(282)
						Label1:
							// line 283: return
							πF.SetLineno(283)
							πR = πg.None
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsection, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_traverse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								µsection = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 288: self.promote_subtitle(section)
							πF.SetLineno(288)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp002[0] = µsection
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpromote_subtitle, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
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
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SectionSubTitle").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSectionSubTitle.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 291: class DocInfo(Transform):
			πF.SetLineno(291)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DocInfo", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 293: """
					πF.SetLineno(293)
					// line 293: """
					πF.SetLineno(293)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    This transform is specific to the reStructuredText_ markup syntax;\n    see \"Bibliographic Fields\" in the `reStructuredText Markup\n    Specification`_ for a high-level description. This transform\n    should be run *after* the `DocTitle` transform.\n\n    Given a field list as the first non-comment element after the\n    document title and subtitle (if present), registered bibliographic\n    field names are transformed to the corresponding DTD elements,\n    becoming child elements of the \"docinfo\" element (except for a\n    dedication and/or an abstract, which become \"topic\" elements after\n    \"docinfo\").\n\n    For example, given this document fragment after parsing::\n\n        <document>\n            <title>\n                Document Title\n            <field_list>\n                <field>\n                    <field_name>\n                        Author\n                    <field_body>\n                        <paragraph>\n                            A. Name\n                <field>\n                    <field_name>\n                        Status\n                    <field_body>\n                        <paragraph>\n                            $RCSfile$\n            ...\n\n    After running the bibliographic field list transform, the\n    resulting document tree would look like this::\n\n        <document>\n            <title>\n                Document Title\n            <docinfo>\n                <author>\n                    A. Name\n                <status>\n                    frontmatter.py\n            ...\n\n    The \"Status\" field contained an expanded RCS keyword, which is\n    normally (but optionally) cleaned up by the transform. The sole\n    contents of the field body must be a paragraph containing an\n    expanded RCS keyword of the form \"$keyword: expansion text $\". Any\n    RCS keyword can be processed in any bibliographic field. The\n    dollar signs and leading RCS keyword name are removed. Extra\n    processing is done for the following RCS keywords:\n\n    - \"RCSfile\" expands to the name of the file in the RCS or CVS\n      repository, which is the name of the source file with a \",v\"\n      suffix appended. The transform will remove the \",v\" suffix.\n\n    - \"Date\" expands to the format \"YYYY/MM/DD hh:mm:ss\" (in the UTC\n      time zone). The RCS Keywords transform will extract just the\n      date itself and transform it to an ISO 8601 format date, as in\n      \"2000-12-31\".\n\n      (Since the source file for this text is itself stored under CVS,\n      we can't show an example of the \"Date\" RCS keyword because we\n      can't prevent any RCS keywords used in this explanation from\n      being expanded. Only the \"RCSfile\" keyword is stable; its\n      expansion text changes only if the file name changes.)\n\n    .. _reStructuredText: http://docutils.sf.net/rst.html\n    .. _reStructuredText Markup Specification:\n       http://docutils.sf.net/docs/ref/rst/restructuredtext.html\n    ").ToObject()); πE != nil {
						continue
					}
					// line 367: default_priority = 340
					πF.SetLineno(367)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(340).ToObject()); πE != nil {
						continue
					}
					// line 369: biblio_nodes = {
					πF.SetLineno(369)
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßauthor, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßauthor.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßauthors, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßauthors.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßorganization, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßorganization.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßaddress, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßaddress.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontact, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßcontact.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßversion, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßversion.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrevision, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßrevision.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstatus, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßstatus.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdate, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdate.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcopyright, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßcopyright.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtopic, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdedication.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtopic, nil); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßabstract.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßbiblio_nodes.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 382: """Canonical field name (lowcased) to node class name mapping for
					πF.SetLineno(382)
					// line 385: def apply(self):
					πF.SetLineno(385)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdocument *πg.Object = πg.UnboundLocal
						_ = µdocument
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
						var µcandidate *πg.Object = πg.UnboundLocal
						_ = µcandidate
						var µbiblioindex *πg.Object = πg.UnboundLocal
						_ = µbiblioindex
						var µnodelist *πg.Object = πg.UnboundLocal
						_ = µnodelist
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
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = ßdocinfo_xform.ToObject()
							πTemp002[2] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 386: if not getattr(self.document.settings, 'docinfo_xform', 1):
							πF.SetLineno(386)
						Label1:
							// line 387: return
							πF.SetLineno(387)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 388: document = self.document
							πF.SetLineno(388)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							µdocument = πTemp001
							// line 389: index = document.first_child_not_matching_class(
							πF.SetLineno(389)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßPreBibliographic, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßfirst_child_not_matching_class, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µindex = πTemp003
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µindex == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 391: if index is None:
							πF.SetLineno(391)
						Label3:
							// line 392: return
							πF.SetLineno(392)
							πR = πg.None
							continue
							goto Label4
						Label4:
							// line 393: candidate = document[index]
							πF.SetLineno(393)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = µindex
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdocument, πTemp001); πE != nil {
								continue
							}
							µcandidate = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
								continue
							}
							πTemp002[0] = µcandidate
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßfield_list, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 394: if isinstance(candidate, nodes.field_list):
							πF.SetLineno(394)
						Label5:
							// line 395: biblioindex = document.first_child_not_matching_class(
							πF.SetLineno(395)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTitular, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßDecorative, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocument, ßfirst_child_not_matching_class, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µbiblioindex = πTemp003
							// line 397: nodelist = self.extract_bibliographic(candidate)
							πF.SetLineno(397)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcandidate, "candidate"); πE != nil {
								continue
							}
							πTemp002[0] = µcandidate
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßextract_bibliographic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µnodelist = πTemp003
							// line 398: del document[index]         # untransformed field list (candidate)
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = µindex
							if πE = πg.DelItem(πF, µdocument, πTemp001); πE != nil {
								continue
							}
							// line 399: document[biblioindex:biblioindex] = nodelist
							πF.SetLineno(399)
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnodelist); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbiblioindex, "biblioindex"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbiblioindex, "biblioindex"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µbiblioindex, µbiblioindex, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µdocument, πTemp003, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 401: def extract_bibliographic(self, field_list):
					πF.SetLineno(401)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field_list", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("extract_bibliographic", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield_list *πg.Object = πArgs[1]
						_ = µfield_list
						var µdocinfo *πg.Object = πg.UnboundLocal
						_ = µdocinfo
						var µbibliofields *πg.Object = πg.UnboundLocal
						_ = µbibliofields
						var µlabels *πg.Object = πg.UnboundLocal
						_ = µlabels
						var µtopics *πg.Object = πg.UnboundLocal
						_ = µtopics
						var µfield *πg.Object = πg.UnboundLocal
						_ = µfield
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µnormedname *πg.Object = πg.UnboundLocal
						_ = µnormedname
						var µcanonical *πg.Object = πg.UnboundLocal
						_ = µcanonical
						var µbiblioclass *πg.Object = πg.UnboundLocal
						_ = µbiblioclass
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µclassvalue *πg.Object = πg.UnboundLocal
						_ = µclassvalue
						var µnodelist *πg.Object = πg.UnboundLocal
						_ = µnodelist
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
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
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 πg.KWArgs
						_ = πTemp013
						var πTemp014 *πg.BaseException
						_ = πTemp014
						var πTemp015 *πg.Traceback
						_ = πTemp015
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 27:
								goto Label27
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 26:
								goto Label26
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 402: docinfo = nodes.docinfo()
							πF.SetLineno(402)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocinfo, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdocinfo = πTemp001
							// line 403: bibliofields = self.language.bibliographic_fields
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßbibliographic_fields, nil); πE != nil {
								continue
							}
							µbibliofields = πTemp002
							// line 404: labels = self.language.labels
							πF.SetLineno(404)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlabels, nil); πE != nil {
								continue
							}
							µlabels = πTemp002
							// line 405: topics = {'dedication': None, 'abstract': None}
							πF.SetLineno(405)
							πTemp003 = πg.NewDict()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, ßdedication.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, ßabstract.ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp003.ToObject()
							µtopics = πTemp001
							if πE = πg.CheckLocal(πF, µfield_list, "field_list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µfield_list); πE != nil {
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
								µfield = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 407: try:
							πF.SetLineno(407)
							πF.PushCheckpoint(5)
							// line 408: name = field[0][0].astext()
							πF.SetLineno(408)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp007 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfield, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µname = πTemp006
							// line 409: normedname = nodes.fully_normalize_name(name)
							πF.SetLineno(409)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009[0] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßfully_normalize_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µnormedname = πTemp002
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp009[0] = µfield
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp007, πE = πg.Eq(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µnormedname, "normedname"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbibliofields, "bibliofields"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, µbibliofields, µnormedname); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(πTemp011).ToObject()
							πTemp006 = πTemp007
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label6
							}
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp009[0] = µfield
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcheck_empty_biblio_field, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp006 = πTemp008
						Label6:
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 410: if not (len(field) == 2 and normedname in bibliofields
							πF.SetLineno(410)
						Label7:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 412: raise TransformError
							πF.SetLineno(412)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label8
						Label8:
							// line 413: canonical = bibliofields[normedname]
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µnormedname, "normedname"); πE != nil {
								continue
							}
							πTemp002 = µnormedname
							if πE = πg.CheckLocal(πF, µbibliofields, "bibliofields"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µbibliofields, πTemp002); πE != nil {
								continue
							}
							µcanonical = πTemp006
							// line 414: biblioclass = self.biblio_nodes[canonical]
							πF.SetLineno(414)
							if πE = πg.CheckLocal(πF, µcanonical, "canonical"); πE != nil {
								continue
							}
							πTemp002 = µcanonical
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßbiblio_nodes, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							µbiblioclass = πTemp006
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbiblioclass, "biblioclass"); πE != nil {
								continue
							}
							πTemp009[0] = µbiblioclass
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßTextElement, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbiblioclass, "biblioclass"); πE != nil {
								continue
							}
							πTemp009[0] = µbiblioclass
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßauthors, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbiblioclass, "biblioclass"); πE != nil {
								continue
							}
							πTemp009[0] = µbiblioclass
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßtopic, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 415: if issubclass(biblioclass, nodes.TextElement):
							πF.SetLineno(415)
						Label9:
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp009[0] = µfield
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcheck_compound_biblio_field, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							goto Label15
							// line 416: if not self.check_compound_biblio_field(field, name):
							πF.SetLineno(416)
						Label14:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 417: raise TransformError
							πF.SetLineno(417)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label15
						Label15:
							// line 418: utils.clean_rcs_keywords(
							πF.SetLineno(418)
							πTemp009 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp007 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfield, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrcs_keyword_substitutions, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßclean_rcs_keywords, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							// line 420: docinfo.append(biblioclass('', '', *field[1][0]))
							πF.SetLineno(420)
							πTemp009 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = ß.ToObject()
							πTemp012[1] = ß.ToObject()
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp007 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfield, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbiblioclass, "biblioclass"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µbiblioclass, πTemp012, πTemp006, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocinfo, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label13
							// line 421: elif issubclass(biblioclass, nodes.authors):
							πF.SetLineno(421)
						Label10:
							// line 422: self.extract_authors(field, name, docinfo)
							πF.SetLineno(422)
							πTemp009 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp009[0] = µfield
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009[1] = µname
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							πTemp009[2] = µdocinfo
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßextract_authors, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label13
							// line 423: elif issubclass(biblioclass, nodes.topic):
							πF.SetLineno(423)
						Label11:
							if πE = πg.CheckLocal(πF, µcanonical, "canonical"); πE != nil {
								continue
							}
							πTemp002 = µcanonical
							if πE = πg.CheckLocal(πF, µtopics, "topics"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µtopics, πTemp002); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label16
							}
							goto Label17
							// line 424: if topics[canonical]:
							πF.SetLineno(424)
						Label16:
							// line 425: field[-1] += self.document.reporter.warning(
							πF.SetLineno(425)
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp002); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("There can only be one \"%s\" field.").ToObject(), µname); πE != nil {
								continue
							}
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"base_node", µfield},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp007, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp009, πTemp013); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.IAdd(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.SetItem(πF, µfield, πTemp008, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 428: raise TransformError
							πF.SetLineno(428)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label17
						Label17:
							// line 429: title = nodes.title(name, labels[canonical])
							πF.SetLineno(429)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp009[0] = µname
							if πE = πg.CheckLocal(πF, µcanonical, "canonical"); πE != nil {
								continue
							}
							πTemp002 = µcanonical
							if πE = πg.CheckLocal(πF, µlabels, "labels"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µlabels, πTemp002); πE != nil {
								continue
							}
							πTemp009[1] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µtitle = πTemp002
							// line 430: title[0].rawsource =  labels[canonical]
							πF.SetLineno(430)
							if πE = πg.CheckLocal(πF, µcanonical, "canonical"); πE != nil {
								continue
							}
							πTemp002 = µcanonical
							if πE = πg.CheckLocal(πF, µlabels, "labels"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µlabels, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp006); πE != nil {
								continue
							}
							πTemp007 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µtitle, πTemp007); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp008, ßrawsource, πTemp002); πE != nil {
								continue
							}
							// line 431: topics[canonical] = biblioclass(
							πF.SetLineno(431)
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp009[1] = µtitle
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßchildren, nil); πE != nil {
								continue
							}
							πTemp012 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µcanonical, "canonical"); πE != nil {
								continue
							}
							πTemp012[0] = µcanonical
							πTemp006 = πg.NewList(πTemp012...).ToObject()
							πTemp013 = πg.KWArgs{
								{"classes", πTemp006},
							}
							if πE = πg.CheckLocal(πF, µbiblioclass, "biblioclass"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Invoke(πF, µbiblioclass, πTemp009, πTemp002, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtopics, "topics"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcanonical, "canonical"); πE != nil {
								continue
							}
							πTemp007 = µcanonical
							if πE = πg.SetItem(πF, µtopics, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label13
						Label12:
							// line 434: docinfo.append(biblioclass('', *field[1].children))
							πF.SetLineno(434)
							πTemp009 = πF.MakeArgs(1)
							πTemp012 = πF.MakeArgs(1)
							πTemp012[0] = ß.ToObject()
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßchildren, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbiblioclass, "biblioclass"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Invoke(πF, µbiblioclass, πTemp012, πTemp002, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocinfo, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label13
						Label13:
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp014, πTemp015 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp014.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							πE = πF.Raise(πTemp014.ToObject(), nil, πTemp015.ToObject())
							continue
							// line 435: except TransformError:
							πF.SetLineno(435)
						Label18:
							πTemp009 = πF.MakeArgs(1)
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfield, πTemp007); πE != nil {
								continue
							}
							πTemp009[0] = πTemp008
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp006, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label19
							}
							πTemp009 = πF.MakeArgs(2)
							πTemp006 = πg.NewInt(0).ToObject()
							if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µfield, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
								continue
							}
							πTemp009[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp002 = πTemp007
						Label19:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label20
							}
							goto Label21
							// line 436: if len(field[-1]) == 1 \
							πF.SetLineno(436)
						Label20:
							// line 438: utils.clean_rcs_keywords(
							πF.SetLineno(438)
							πTemp009 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfield, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrcs_keyword_substitutions, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßclean_rcs_keywords, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label21
						Label21:
							// line 441: classvalue = nodes.make_id(normedname)
							πF.SetLineno(441)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnormedname, "normedname"); πE != nil {
								continue
							}
							πTemp009[0] = µnormedname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßmake_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µclassvalue = πTemp002
							if πE = πg.CheckLocal(πF, µclassvalue, "classvalue"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µclassvalue); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							goto Label23
							// line 442: if classvalue:
							πF.SetLineno(442)
						Label22:
							// line 443: field['classes'].append(classvalue)
							πF.SetLineno(443)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µclassvalue, "classvalue"); πE != nil {
								continue
							}
							πTemp009[0] = µclassvalue
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp006, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label23
						Label23:
							// line 444: docinfo.append(field)
							πF.SetLineno(444)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp009[0] = µfield
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocinfo, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 445: nodelist = []
							πF.SetLineno(445)
							πTemp009 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp009...).ToObject()
							µnodelist = πTemp001
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							πTemp009[0] = µdocinfo
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp001, πE = πg.NE(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label24
							}
							goto Label25
							// line 446: if len(docinfo) != 0:
							πF.SetLineno(446)
						Label24:
							// line 447: nodelist.append(docinfo)
							πF.SetLineno(447)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							πTemp009[0] = µdocinfo
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnodelist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label25
						Label25:
							πTemp002 = πg.NewTuple2(ßdedication.ToObject(), ßabstract.ToObject()).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(27)
							πTemp004 = false
						Label26:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label28
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
							πF.PushCheckpoint(26)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.CheckLocal(πF, µtopics, "topics"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µtopics, πTemp002); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label29
							}
							goto Label30
							// line 449: if topics[name]:
							πF.SetLineno(449)
						Label29:
							// line 450: nodelist.append(topics[name])
							πF.SetLineno(450)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.CheckLocal(πF, µtopics, "topics"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µtopics, πTemp002); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnodelist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label30
						Label30:
							continue
						Label27:
							if πE != nil || πR != nil {
								continue
							}
						Label28:
							// line 451: return nodelist
							πF.SetLineno(451)
							if πE = πg.CheckLocal(πF, µnodelist, "nodelist"); πE != nil {
								continue
							}
							πR = µnodelist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßextract_bibliographic.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 453: def check_empty_biblio_field(self, field, name):
					πF.SetLineno(453)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field", Def: nil}
					πTemp004[2] = πg.Param{Name: "name", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("check_empty_biblio_field", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield *πg.Object = πArgs[1]
						_ = µfield
						var µname *πg.Object = πArgs[2]
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
						var πTemp006 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
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
							if πTemp001, πE = πg.LT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 454: if len(field[-1]) < 1:
							πF.SetLineno(454)
						Label1:
							// line 455: field[-1] += self.document.reporter.warning(
							πF.SetLineno(455)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µfield, πTemp001); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Cannot extract empty bibliographic field \"%s\".").ToObject(), µname); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"base_node", µfield},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.IAdd(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.SetItem(πF, µfield, πTemp007, πTemp001); πE != nil {
								continue
							}
							// line 458: return None
							πF.SetLineno(458)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 459: return 1
							πF.SetLineno(459)
							πR = πg.NewInt(1).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_empty_biblio_field.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 461: def check_compound_biblio_field(self, field, name):
					πF.SetLineno(461)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field", Def: nil}
					πTemp004[2] = πg.Param{Name: "name", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("check_compound_biblio_field", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield *πg.Object = πArgs[1]
						_ = µfield
						var µname *πg.Object = πArgs[2]
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
						var πTemp006 πg.KWArgs
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
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
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
							if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 462: if len(field[-1]) > 1:
							πF.SetLineno(462)
						Label1:
							// line 463: field[-1] += self.document.reporter.warning(
							πF.SetLineno(463)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µfield, πTemp001); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Cannot extract compound bibliographic field \"%s\".").ToObject(), µname); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"base_node", µfield},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.IAdd(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.SetItem(πF, µfield, πTemp007, πTemp001); πE != nil {
								continue
							}
							// line 466: return None
							πF.SetLineno(466)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(0).ToObject()
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µfield, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 467: if not isinstance(field[-1][0], nodes.paragraph):
							πF.SetLineno(467)
						Label3:
							// line 468: field[-1] += self.document.reporter.warning(
							πF.SetLineno(468)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µfield, πTemp001); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Cannot extract bibliographic field \"%s\" containing anything other than a single paragraph.").ToObject(), µname); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"base_node", µfield},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.IAdd(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.SetItem(πF, µfield, πTemp007, πTemp001); πE != nil {
								continue
							}
							// line 472: return None
							πF.SetLineno(472)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 473: return 1
							πF.SetLineno(473)
							πR = πg.NewInt(1).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_compound_biblio_field.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 475: rcs_keyword_substitutions = [
					πF.SetLineno(475)
					πTemp007 = make([]*πg.Object, 3)
					πTemp009 = πF.MakeArgs(2)
					πTemp009[0] = πg.NewStr("\\$Date: (\\d\\d\\d\\d)[-/](\\d\\d)[-/](\\d\\d)[ T][\\d:]+[^$]* \\$").ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßIGNORECASE, nil); πE != nil {
						continue
					}
					πTemp009[1] = πTemp011
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008 = πg.NewTuple2(πTemp010, πg.NewStr("\\1-\\2-\\3").ToObject()).ToObject()
					πTemp007[0] = πTemp008
					πTemp009 = πF.MakeArgs(2)
					πTemp009[0] = πg.NewStr("\\$RCSfile: (.+),v \\$").ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßIGNORECASE, nil); πE != nil {
						continue
					}
					πTemp009[1] = πTemp011
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008 = πg.NewTuple2(πTemp010, πg.NewStr("\\1").ToObject()).ToObject()
					πTemp007[1] = πTemp008
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr("\\$[a-zA-Z]+: (.+) \\$").ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008 = πg.NewTuple2(πTemp010, πg.NewStr("\\1").ToObject()).ToObject()
					πTemp007[2] = πTemp008
					πTemp008 = πg.NewList(πTemp007...).ToObject()
					if πE = πClass.SetItem(πF, ßrcs_keyword_substitutions.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 481: def extract_authors(self, field, name, docinfo):
					πF.SetLineno(481)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field", Def: nil}
					πTemp004[2] = πg.Param{Name: "name", Def: nil}
					πTemp004[3] = πg.Param{Name: "docinfo", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("extract_authors", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield *πg.Object = πArgs[1]
						_ = µfield
						var µname *πg.Object = πArgs[2]
						_ = µname
						var µdocinfo *πg.Object = πArgs[3]
						_ = µdocinfo
						var µauthors *πg.Object = πg.UnboundLocal
						_ = µauthors
						var µauthornodes *πg.Object = πg.UnboundLocal
						_ = µauthornodes
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 πg.KWArgs
						_ = πTemp013
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
							// line 482: try:
							πF.SetLineno(482)
							πF.PushCheckpoint(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
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
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 483: if len(field[1]) == 1:
							πF.SetLineno(483)
						Label3:
							πTemp002 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßbullet_list, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 484: if isinstance(field[1][0], nodes.paragraph):
							πF.SetLineno(484)
						Label6:
							// line 485: authors = self.authors_from_one_paragraph(field)
							πF.SetLineno(485)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßauthors_from_one_paragraph, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µauthors = πTemp003
							goto Label9
							// line 486: elif isinstance(field[1][0], nodes.bullet_list):
							πF.SetLineno(486)
						Label7:
							// line 487: authors = self.authors_from_bullet_list(field)
							πF.SetLineno(487)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßauthors_from_bullet_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µauthors = πTemp003
							goto Label9
						Label8:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 489: raise TransformError
							πF.SetLineno(489)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label9
						Label9:
							goto Label5
						Label4:
							// line 491: authors = self.authors_from_paragraphs(field)
							πF.SetLineno(491)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßauthors_from_paragraphs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µauthors = πTemp003
							goto Label5
						Label5:
							// line 492: authornodes = [nodes.author('', '', *author)
							πF.SetLineno(492)
							πTemp007 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µauthor *πg.Object = πg.UnboundLocal
								_ = µauthor
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
										if πE = πg.CheckLocal(πF, µauthors, "authors"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µauthors); πE != nil {
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
											µauthor = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										if πE = πg.CheckLocal(πF, µauthor, "author"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.IsTrue(πF, µauthor); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 492: authornodes = [nodes.author('', '', *author)
										πF.SetLineno(492)
									Label4:
										// line 492: authornodes = [nodes.author('', '', *author)
										πF.SetLineno(492)
										πTemp005 = πF.MakeArgs(2)
										πTemp005[0] = ß.ToObject()
										πTemp005[1] = ß.ToObject()
										if πE = πg.CheckLocal(πF, µauthor, "author"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßauthor, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Invoke(πF, πTemp006, πTemp005, µauthor, nil, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(6)
										return πTemp004, nil
									Label6:
										πTemp006 = πSent
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µauthornodes = πTemp001
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µauthornodes, "authornodes"); πE != nil {
								continue
							}
							πTemp002[0] = µauthornodes
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.GE(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 494: if len(authornodes) >= 1:
							πF.SetLineno(494)
						Label10:
							// line 495: docinfo.append(nodes.authors('', *authornodes))
							πF.SetLineno(495)
							πTemp002 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µauthornodes, "authornodes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßauthors, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp004, πTemp008, µauthornodes, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdocinfo, "docinfo"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdocinfo, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label12
						Label11:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 497: raise TransformError
							πF.SetLineno(497)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label12
						Label12:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label13
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 498: except TransformError:
							πF.SetLineno(498)
						Label13:
							// line 499: field[-1] += self.document.reporter.warning(
							πF.SetLineno(499)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp001); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßauthor_separators, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp012
							if πTemp011, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp006 = πg.NewTuple2(µname, πTemp012).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Bibliographic field \"%s\" incompatible with extraction: it must contain either a single paragraph (with authors separated by one of \"%s\"), multiple paragraphs (one per author), or a bullet list with one paragraph (one author) per item.").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp013 = πg.KWArgs{
								{"base_node", µfield},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp006, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp002, πTemp013); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.IAdd(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp011 = πTemp012
							if πE = πg.SetItem(πF, µfield, πTemp011, πTemp001); πE != nil {
								continue
							}
							// line 507: raise
							πF.SetLineno(507)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßextract_authors.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 509: def authors_from_one_paragraph(self, field):
					πF.SetLineno(509)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("authors_from_one_paragraph", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield *πg.Object = πArgs[1]
						_ = µfield
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µauthorsep *πg.Object = πg.UnboundLocal
						_ = µauthorsep
						var µpattern *πg.Object = πg.UnboundLocal
						_ = µpattern
						var µauthornames *πg.Object = πg.UnboundLocal
						_ = µauthornames
						var µauthors *πg.Object = πg.UnboundLocal
						_ = µauthors
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 510: """Return list of Text nodes for authornames.
							πF.SetLineno(510)
							// line 515: text = ''.join(unicode(node)
							πF.SetLineno(515)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										πTemp002 = πF.MakeArgs(1)
										if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
											continue
										}
										πTemp002[0] = πTemp004
										πTemp003 = πg.NewInt(1).ToObject()
										if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßtraverse, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
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
										// line 515: text = ''.join(unicode(node)
										πF.SetLineno(515)
										πTemp002 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
											continue
										}
										πTemp002[0] = µnode
										if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp003 = πSent
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
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext = πTemp005
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µtext); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 517: if not text:
							πF.SetLineno(517)
						Label1:
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 518: raise TransformError
							πF.SetLineno(518)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßauthor_separators, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp006 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µauthorsep = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 521: pattern = '(?<!\x00)%s' % authorsep
							πF.SetLineno(521)
							if πE = πg.CheckLocal(πF, µauthorsep, "authorsep"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mod(πF, πg.NewStr("(?<!\x00)%s").ToObject(), µauthorsep); πE != nil {
								continue
							}
							µpattern = πTemp005
							// line 522: authornames = re.split(pattern, text)
							πF.SetLineno(522)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp001[0] = µpattern
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[1] = µtext
							if πTemp005, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µauthornames = πTemp005
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µauthornames, "authornames"); πE != nil {
								continue
							}
							πTemp001[0] = µauthornames
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.GT(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 523: if len(authornames) > 1:
							πF.SetLineno(523)
						Label6:
							// line 524: break
							πF.SetLineno(524)
							πTemp006 = true
							continue
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 525: authornames = (name.strip() for name in authornames)
							πF.SetLineno(525)
							πTemp003 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µname *πg.Object = πg.UnboundLocal
								_ = µname
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
										if πE = πg.CheckLocal(πF, µauthornames, "authornames"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µauthornames); πE != nil {
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
											µname = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 525: authornames = (name.strip() for name in authornames)
										πF.SetLineno(525)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µname, ßstrip, nil); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
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
							µauthornames = πTemp005
							// line 526: authors = [[nodes.Text(name, utils.unescape(name, True))]
							πF.SetLineno(526)
							πTemp003 = make([]πg.Param, 0)
							πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µname *πg.Object = πg.UnboundLocal
								_ = µname
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
								var πTemp006 []*πg.Object
								_ = πTemp006
								var πTemp007 []*πg.Object
								_ = πTemp007
								var πTemp008 *πg.Object
								_ = πTemp008
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
										if πE = πg.CheckLocal(πF, µauthornames, "authornames"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µauthornames); πE != nil {
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
											µname = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.IsTrue(πF, µname); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 526: authors = [[nodes.Text(name, utils.unescape(name, True))]
										πF.SetLineno(526)
									Label4:
										// line 526: authors = [[nodes.Text(name, utils.unescape(name, True))]
										πF.SetLineno(526)
										πTemp005 = make([]*πg.Object, 1)
										πTemp006 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										πTemp006[0] = µname
										πTemp007 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										πTemp007[0] = µname
										if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
											continue
										}
										πTemp007[1] = πTemp004
										if πTemp004, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßunescape, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										πTemp006[1] = πTemp004
										if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßText, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
										πTemp005[0] = πTemp004
										πTemp004 = πg.NewList(πTemp005...).ToObject()
										πF.PushCheckpoint(6)
										return πTemp004, nil
									Label6:
										πTemp008 = πSent
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
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							µauthors = πTemp005
							// line 528: return authors
							πF.SetLineno(528)
							if πE = πg.CheckLocal(πF, µauthors, "authors"); πE != nil {
								continue
							}
							πR = µauthors
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßauthors_from_one_paragraph.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 510: """Return list of Text nodes for authornames.
					πF.SetLineno(510)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Return list of Text nodes for authornames.\n        \n        The set of separators is locale dependent (default: \";\"- or \",\").\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßauthors_from_one_paragraph); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 530: def authors_from_bullet_list(self, field):
					πF.SetLineno(530)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("authors_from_bullet_list", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield *πg.Object = πArgs[1]
						_ = µfield
						var µauthors *πg.Object = πg.UnboundLocal
						_ = µauthors
						var µitem *πg.Object = πg.UnboundLocal
						_ = µitem
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 531: authors = []
							πF.SetLineno(531)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µauthors = πTemp002
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µitem = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcomment, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label4
							}
							goto Label5
							// line 533: if isinstance(item, nodes.comment):
							πF.SetLineno(533)
						Label4:
							// line 534: continue
							πF.SetLineno(534)
							continue
							goto Label5
						Label5:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.NE(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µitem, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp009).ToObject()
							πTemp003 = πTemp004
						Label6:
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 535: if len(item) != 1 or not isinstance(item[0], nodes.paragraph):
							πF.SetLineno(535)
						Label7:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 536: raise TransformError
							πF.SetLineno(536)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label8
						Label8:
							// line 537: authors.append(item[0].children)
							πF.SetLineno(537)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µitem, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßchildren, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µauthors, "authors"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µauthors, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µauthors, "authors"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µauthors); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 538: if not authors:
							πF.SetLineno(538)
						Label9:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 539: raise TransformError
							πF.SetLineno(539)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label10
						Label10:
							// line 540: return authors
							πF.SetLineno(540)
							if πE = πg.CheckLocal(πF, µauthors, "authors"); πE != nil {
								continue
							}
							πR = µauthors
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßauthors_from_bullet_list.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 542: def authors_from_paragraphs(self, field):
					πF.SetLineno(542)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "field", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("authors_from_paragraphs", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfield *πg.Object = πArgs[1]
						_ = µfield
						var µitem *πg.Object = πg.UnboundLocal
						_ = µitem
						var µauthors *πg.Object = πg.UnboundLocal
						_ = µauthors
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []πg.Param
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
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µfield, πTemp002); πE != nil {
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
								µitem = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp006[0] = µitem
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßcomment, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp008, πTemp009).ToObject()
							πTemp006[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
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
							// line 544: if not isinstance(item, (nodes.paragraph, nodes.comment)):
							πF.SetLineno(544)
						Label4:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTransformError); πE != nil {
								continue
							}
							// line 545: raise TransformError
							πF.SetLineno(545)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 546: authors = [item.children for item in field[1]
							πF.SetLineno(546)
							πTemp010 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/transforms/frontmatter.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µitem *πg.Object = πg.UnboundLocal
								_ = µitem
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
										πTemp002 = πg.NewInt(1).ToObject()
										if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, µfield, πTemp002); πE != nil {
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
											µitem = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										πTemp006 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
											continue
										}
										πTemp006[0] = µitem
										if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
											continue
										}
										if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßcomment, nil); πE != nil {
											continue
										}
										πTemp006[1] = πTemp007
										if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
											continue
										}
										if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
										if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
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
										// line 546: authors = [item.children for item in field[1]
										πF.SetLineno(546)
									Label4:
										// line 546: authors = [item.children for item in field[1]
										πF.SetLineno(546)
										if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µitem, ßchildren, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return πTemp002, nil
									Label6:
										πTemp003 = πSent
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
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							µauthors = πTemp001
							// line 548: return authors
							πF.SetLineno(548)
							if πE = πg.CheckLocal(πF, µauthors, "authors"); πE != nil {
								continue
							}
							πR = µauthors
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßauthors_from_paragraphs.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DocInfo").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDocInfo.ToObject(), πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("frontmatter", Code)
}
