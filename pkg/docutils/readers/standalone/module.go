package standalone

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/readers/standalone.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßAnonymousHyperlinks := πg.InternStr("AnonymousHyperlinks")
		ßDanglingReferences := πg.InternStr("DanglingReferences")
		ßDocInfo := πg.InternStr("DocInfo")
		ßDocTitle := πg.InternStr("DocTitle")
		ßExternalTargets := πg.InternStr("ExternalTargets")
		ßFootnotes := πg.InternStr("Footnotes")
		ßIndirectHyperlinks := πg.InternStr("IndirectHyperlinks")
		ßInternalTargets := πg.InternStr("InternalTargets")
		ßNone := πg.InternStr("None")
		ßPropagateTargets := πg.InternStr("PropagateTargets")
		ßReader := πg.InternStr("Reader")
		ßSectionSubTitle := πg.InternStr("SectionSubTitle")
		ßSubstitutions := πg.InternStr("Substitutions")
		ßTransitions := πg.InternStr("Transitions")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßaction := πg.InternStr("action")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßdefault := πg.InternStr("default")
		ßdest := πg.InternStr("dest")
		ßdocinfo_xform := πg.InternStr("docinfo_xform")
		ßdoctitle_xform := πg.InternStr("doctitle_xform")
		ßdocument := πg.InternStr("document")
		ßfrontend := πg.InternStr("frontend")
		ßfrontmatter := πg.InternStr("frontmatter")
		ßget_transforms := πg.InternStr("get_transforms")
		ßmisc := πg.InternStr("misc")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreaders := πg.InternStr("readers")
		ßreferences := πg.InternStr("references")
		ßsectsubtitle_xform := πg.InternStr("sectsubtitle_xform")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßstandalone := πg.InternStr("standalone")
		ßstore_false := πg.InternStr("store_false")
		ßstore_true := πg.InternStr("store_true")
		ßsupported := πg.InternStr("supported")
		ßsys := πg.InternStr("sys")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidator := πg.InternStr("validator")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
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
			// line 5: """
			πF.SetLineno(5)
			// line 5: """
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nStandalone file Reader for the reStructuredText markup syntax.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 12: import sys
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from docutils import frontend, readers
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßfrontend.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.readers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßreaders.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: from docutils.transforms import frontmatter, references, misc
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.frontmatter"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßfrontmatter.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.references"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßreferences.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.misc"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßmisc.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: class Reader(readers.Reader):
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßreaders); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßReader, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Reader", "/usr/lib/python2.7/site-packages/docutils/readers/standalone.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Dict
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []πg.Param
				_ = πTemp012
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 19: supported = ('standalone',)
					πF.SetLineno(19)
					πTemp001 = πg.NewTuple1(ßstandalone.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 20: """Contexts this reader supports."""
					πF.SetLineno(20)
					// line 20: """Contexts this reader supports."""
					πF.SetLineno(20)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Contexts this reader supports.").ToObject()); πE != nil {
						continue
					}
					// line 22: document = None
					πF.SetLineno(22)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdocument.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 23: """A single document tree."""
					πF.SetLineno(23)
					// line 25: settings_spec = (
					πF.SetLineno(25)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--no-doc-title").ToObject()
					πTemp006 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßdest.ToObject(), ßdoctitle_xform.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp009); πE != nil {
						continue
					}
					πTemp008 = πTemp007.ToObject()
					πTemp004 = πg.NewTuple3(πg.NewStr("Disable the promotion of a lone top-level section title to document title (and subsequent section title to document subtitle promotion; enabled by default).").ToObject(), πTemp006, πTemp008).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--no-doc-info").ToObject()
					πTemp008 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßdest.ToObject(), ßdocinfo_xform.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp007.ToObject()
					πTemp006 = πg.NewTuple3(πg.NewStr("Disable the bibliographic field list transform (enabled by default).").ToObject(), πTemp008, πTemp009).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--section-subtitles").ToObject()
					πTemp009 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßdest.ToObject(), ßsectsubtitle_xform.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßdefault.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp010 = πTemp007.ToObject()
					πTemp008 = πg.NewTuple3(πg.NewStr("Activate the promotion of lone subsection titles to section subtitles (disabled by default).").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewStr("--no-section-subtitles").ToObject()
					πTemp010 = πg.NewList(πTemp005...).ToObject()
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßdest.ToObject(), ßsectsubtitle_xform.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp007.ToObject()
					πTemp009 = πg.NewTuple3(πg.NewStr("Deactivate the promotion of lone subsection titles.").ToObject(), πTemp010, πTemp011).ToObject()
					πTemp003 = πg.NewTuple4(πTemp004, πTemp006, πTemp008, πTemp009).ToObject()
					πTemp001 = πg.NewTuple3(πg.NewStr("Standalone Reader").ToObject(), πTemp002, πTemp003).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 49: config_section = 'standalone reader'
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("standalone reader").ToObject()); πE != nil {
						continue
					}
					// line 50: config_section_dependencies = ('readers',)
					πF.SetLineno(50)
					πTemp001 = πg.NewTuple1(ßreaders.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 52: def get_transforms(self):
					πF.SetLineno(52)
					πTemp012 = make([]πg.Param, 1)
					πTemp012[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/readers/standalone.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 *πg.Object
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
							// line 53: return readers.Reader.get_transforms(self) + [
							πF.SetLineno(53)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreaders); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßReader, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßget_transforms, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = make([]*πg.Object, 12)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßSubstitutions, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPropagateTargets, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrontmatter); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßDocTitle, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrontmatter); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßSectionSubTitle, nil); πE != nil {
								continue
							}
							πTemp002[3] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfrontmatter); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßDocInfo, nil); πE != nil {
								continue
							}
							πTemp002[4] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßAnonymousHyperlinks, nil); πE != nil {
								continue
							}
							πTemp002[5] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßIndirectHyperlinks, nil); πE != nil {
								continue
							}
							πTemp002[6] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßFootnotes, nil); πE != nil {
								continue
							}
							πTemp002[7] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßExternalTargets, nil); πE != nil {
								continue
							}
							πTemp002[8] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßInternalTargets, nil); πE != nil {
								continue
							}
							πTemp002[9] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßDanglingReferences, nil); πE != nil {
								continue
							}
							πTemp002[10] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmisc); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßTransitions, nil); πE != nil {
								continue
							}
							πTemp002[11] = πTemp005
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_transforms.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Reader").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReader.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("standalone", Code)
}
