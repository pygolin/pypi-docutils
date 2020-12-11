package pep

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/readers/pep.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßContents := πg.InternStr("Contents")
		ßDocInfo := πg.InternStr("DocInfo")
		ßDocTitle := πg.InternStr("DocTitle")
		ßHeaders := πg.InternStr("Headers")
		ßInliner := πg.InternStr("Inliner")
		ßNone := πg.InternStr("None")
		ßParser := πg.InternStr("Parser")
		ßReader := πg.InternStr("Reader")
		ßSectionSubTitle := πg.InternStr("SectionSubTitle")
		ßTargetNotes := πg.InternStr("TargetNotes")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßextend := πg.InternStr("extend")
		ßfrontmatter := πg.InternStr("frontmatter")
		ßget_transforms := πg.InternStr("get_transforms")
		ßinliner_class := πg.InternStr("inliner_class")
		ßmisc := πg.InternStr("misc")
		ßpep := πg.InternStr("pep")
		ßpep_references := πg.InternStr("pep_references")
		ßpeps := πg.InternStr("peps")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreaders := πg.InternStr("readers")
		ßreferences := πg.InternStr("references")
		ßremove := πg.InternStr("remove")
		ßrfc_references := πg.InternStr("rfc_references")
		ßrst := πg.InternStr("rst")
		ßsettings_default_overrides := πg.InternStr("settings_default_overrides")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßstandalone := πg.InternStr("standalone")
		ßstates := πg.InternStr("states")
		ßsupported := πg.InternStr("supported")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nPython Enhancement Proposal (PEP) Reader.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 12: from docutils.readers import standalone
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.readers.standalone"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßstandalone.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from docutils.transforms import peps, references, misc, frontmatter
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.peps"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßpeps.ToObject(), πTemp001); πE != nil {
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.frontmatter"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßfrontmatter.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: from docutils.parsers import rst
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßrst.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: class Reader(standalone.Reader):
			πF.SetLineno(17)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstandalone); πE != nil {
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
			_, πE = πg.NewCode("Reader", "/usr/lib/python2.7/site-packages/docutils/readers/pep.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
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
					// line 19: supported = ('pep',)
					πF.SetLineno(19)
					πTemp001 = πg.NewTuple1(ßpep.ToObject()).ToObject()
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
					// line 22: settings_spec = (
					πF.SetLineno(22)
					πTemp002 = πg.NewTuple0().ToObject()
					πTemp001 = πg.NewTuple3(πg.NewStr("PEP Reader Option Defaults").ToObject(), πg.NewStr("The --pep-references and --rfc-references options (for the reStructuredText parser) are on by default.").ToObject(), πTemp002).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 28: config_section = 'pep reader'
					πF.SetLineno(28)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("pep reader").ToObject()); πE != nil {
						continue
					}
					// line 29: config_section_dependencies = ('readers', 'standalone reader')
					πF.SetLineno(29)
					πTemp001 = πg.NewTuple2(ßreaders.ToObject(), πg.NewStr("standalone reader").ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 31: def get_transforms(self):
					πF.SetLineno(31)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/readers/pep.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtransforms *πg.Object = πg.UnboundLocal
						_ = µtransforms
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 32: transforms = standalone.Reader.get_transforms(self)
							πF.SetLineno(32)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstandalone); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßReader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget_transforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtransforms = πTemp003
							// line 34: transforms.remove(frontmatter.DocTitle)
							πF.SetLineno(34)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfrontmatter); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßDocTitle, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtransforms, "transforms"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtransforms, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 35: transforms.remove(frontmatter.SectionSubTitle)
							πF.SetLineno(35)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfrontmatter); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSectionSubTitle, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtransforms, "transforms"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtransforms, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 36: transforms.remove(frontmatter.DocInfo)
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfrontmatter); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßDocInfo, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtransforms, "transforms"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtransforms, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 37: transforms.extend([peps.Headers, peps.Contents, peps.TargetNotes])
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpeps); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßHeaders, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpeps); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßContents, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpeps); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTargetNotes, nil); πE != nil {
								continue
							}
							πTemp004[2] = πTemp003
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µtransforms, "transforms"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtransforms, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 38: return transforms
							πF.SetLineno(38)
							if πE = πg.CheckLocal(πF, µtransforms, "transforms"); πE != nil {
								continue
							}
							πR = µtransforms
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
					// line 40: settings_default_overrides = {'pep_references': 1, 'rfc_references': 1}
					πF.SetLineno(40)
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßpep_references.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßrfc_references.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßsettings_default_overrides.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 42: inliner_class = rst.states.Inliner
					πF.SetLineno(42)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßrst); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßstates, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßInliner, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßinliner_class.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 44: def __init__(self, parser=None, parser_name=None):
					πF.SetLineno(44)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "parser", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "parser_name", Def: πTemp005}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/readers/pep.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µparser *πg.Object = πArgs[1]
						_ = µparser
						var µparser_name *πg.Object = πArgs[2]
						_ = µparser_name
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 45: """`parser` should be ``None``."""
							πF.SetLineno(45)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µparser == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 46: if parser is None:
							πF.SetLineno(46)
						Label1:
							// line 47: parser = rst.Parser(rfc2822=True, inliner=self.inliner_class())
							πF.SetLineno(47)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinliner_class, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"rfc2822", πTemp001},
								{"inliner", πTemp004},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrst); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßParser, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							µparser = πTemp001
							goto Label2
						Label2:
							// line 48: standalone.Reader.__init__(self, parser, '')
							πF.SetLineno(48)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[0] = µself
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							πTemp006[1] = µparser
							πTemp006[2] = ß.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstandalone); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßReader, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 45: """`parser` should be ``None``."""
					πF.SetLineno(45)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("`parser` should be ``None``.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
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
	πg.RegisterModule("pep", Code)
}
