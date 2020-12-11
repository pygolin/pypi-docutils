package docutils

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/collections"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/warnings"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßApplicationError := πg.InternStr("ApplicationError")
		ßComponent := πg.InternStr("Component")
		ßDataError := πg.InternStr("DataError")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßException := πg.InternStr("Exception")
		ßNone := πg.InternStr("None")
		ßSettingsSpec := πg.InternStr("SettingsSpec")
		ßTransformSpec := πg.InternStr("TransformSpec")
		ßTrue := πg.InternStr("True")
		ßVersionInfo := πg.InternStr("VersionInfo")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß__version_details__ := πg.InternStr("__version_details__")
		ß__version_info__ := πg.InternStr("__version_info__")
		ßcomponent_type := πg.InternStr("component_type")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßdefault_transforms := πg.InternStr("default_transforms")
		ßfinal := πg.InternStr("final")
		ßget_transforms := πg.InternStr("get_transforms")
		ßlist := πg.InternStr("list")
		ßnamedtuple := πg.InternStr("namedtuple")
		ßobject := πg.InternStr("object")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrelative_path_settings := πg.InternStr("relative_path_settings")
		ßrelease := πg.InternStr("release")
		ßsettings_default_overrides := πg.InternStr("settings_default_overrides")
		ßsettings_defaults := πg.InternStr("settings_defaults")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsupported := πg.InternStr("supported")
		ßsupports := πg.InternStr("supports")
		ßsys := πg.InternStr("sys")
		ßunknown_reference_resolvers := πg.InternStr("unknown_reference_resolvers")
		ßwarn := πg.InternStr("warn")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 πg.KWArgs
		_ = πTemp004
		var πTemp005 *πg.Dict
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis is the Docutils (Python Documentation Utilities) package.\n\nPackage Structure\n=================\n\nModules:\n\n- __init__.py: Contains component base classes, exception classes, and\n  Docutils version information.\n\n- core.py: Contains the ``Publisher`` class and ``publish_*()`` convenience\n  functions.\n\n- frontend.py: Runtime settings (command-line interface, configuration files)\n  processing, for Docutils front-ends.\n\n- io.py: Provides a uniform API for low-level input and output.\n\n- nodes.py: Docutils document tree (doctree) node class library.\n\n- statemachine.py: A finite state machine specialized for\n  regular-expression-based text filters.\n\nSubpackages:\n\n- languages: Language-specific mappings of terms.\n\n- parsers: Syntax-specific input parser modules or packages.\n\n- readers: Context-specific input handlers which understand the data\n  source and manage a parser.\n\n- transforms: Modules used by readers and writers to modify DPS\n  doctrees.\n\n- utils: Contains the ``Reporter`` system warning class and miscellaneous\n  utilities used by readers, writers, and transforms.\n\n  utils/urischemes.py: Contains a complete mapping of known URI addressing\n  scheme names to descriptions.\n\n- utils/math: Contains functions for conversion of mathematical notation\n  between different formats (LaTeX, MathML, text, ...).\n\n- writers: Format-specific output translators.\n").ToObject()); πE != nil {
				continue
			}
			// line 53: import sys
			πF.SetLineno(53)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 54: from collections import namedtuple
			πF.SetLineno(54)
			if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßnamedtuple); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnamedtuple.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 57: __docformat__ = 'reStructuredText'
			πF.SetLineno(57)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 59: __version__ = '0.16'
			πF.SetLineno(59)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("0.16").ToObject()); πE != nil {
				continue
			}
			// line 60: """Docutils version identifier (complies with PEP 440)::
			πF.SetLineno(60)
			// line 70: VersionInfo = namedtuple(
			πF.SetLineno(70)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßVersionInfo.ToObject()
			πTemp002[1] = πg.NewStr("major minor micro releaselevel serial release").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßnamedtuple); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßVersionInfo.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 73: __version_info__ = VersionInfo(
			πF.SetLineno(73)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp004 = πg.KWArgs{
				{"major", πg.NewInt(0).ToObject()},
				{"minor", πg.NewInt(16).ToObject()},
				{"micro", πg.NewInt(0).ToObject()},
				{"releaselevel", ßfinal.ToObject()},
				{"serial", πg.NewInt(0).ToObject()},
				{"release", πTemp001},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßVersionInfo); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__version_info__.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 82: """Comprehensive version information tuple. See 'Version Numbering' in
			πF.SetLineno(82)
			// line 85: __version_details__ = 'release'
			πF.SetLineno(85)
			if πE = πF.Globals().SetItem(πF, ß__version_details__.ToObject(), ßrelease.ToObject()); πE != nil {
				continue
			}
			// line 86: """Optional extra version details (e.g. 'snapshot 2005-05-29, r3410').
			πF.SetLineno(86)
			// line 91: class ApplicationError(Exception): pass
			πF.SetLineno(91)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ApplicationError", "/usr/lib/python2.7/site-packages/docutils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 91: class ApplicationError(Exception): pass
					πF.SetLineno(91)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ApplicationError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßApplicationError.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 93: class DataError(ApplicationError): pass
			πF.SetLineno(93)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßApplicationError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DataError", "/usr/lib/python2.7/site-packages/docutils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 93: class DataError(ApplicationError): pass
					πF.SetLineno(93)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DataError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDataError.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 96: class SettingsSpec(object):
			πF.SetLineno(96)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SettingsSpec", "/usr/lib/python2.7/site-packages/docutils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 98: """
					πF.SetLineno(98)
					// line 98: """
					πF.SetLineno(98)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Runtime setting specification base class.\n\n    SettingsSpec subclass objects used by `docutils.frontend.OptionParser`.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 104: settings_spec = ()
					πF.SetLineno(104)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 105: """Runtime settings specification.  Override in subclasses.
					πF.SetLineno(105)
					// line 143: settings_defaults = None
					πF.SetLineno(143)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsettings_defaults.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 144: """A dictionary of defaults for settings not in `settings_spec` (internal
					πF.SetLineno(144)
					// line 148: settings_default_overrides = None
					πF.SetLineno(148)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßsettings_default_overrides.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 149: """A dictionary of auxiliary defaults, to override defaults for settings
					πF.SetLineno(149)
					// line 152: relative_path_settings = ()
					πF.SetLineno(152)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßrelative_path_settings.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 153: """Settings containing filesystem paths.  Override in subclasses.
					πF.SetLineno(153)
					// line 157: config_section = None
					πF.SetLineno(157)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 158: """The name of the config file section specific to this component
					πF.SetLineno(158)
					// line 161: config_section_dependencies = None
					πF.SetLineno(161)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 162: """A list of names of config file sections that are to be applied before
					πF.SetLineno(162)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SettingsSpec").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSettingsSpec.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 169: class TransformSpec:
			πF.SetLineno(169)
			πTemp002 = make([]*πg.Object, 0)
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TransformSpec", "/usr/lib/python2.7/site-packages/docutils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
					// line 171: """
					πF.SetLineno(171)
					// line 171: """
					πF.SetLineno(171)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Runtime transform specification base class.\n\n    TransformSpec subclass objects used by `docutils.transforms.Transformer`.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 177: def get_transforms(self):
					πF.SetLineno(177)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µwarnings *πg.Object = πg.UnboundLocal
						_ = µwarnings
						var πTemp001 *πg.Object
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
							// line 178: """Transforms required by this class.  Override in subclasses."""
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefault_transforms, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple0().ToObject()
							if πTemp001, πE = πg.NE(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 179: if self.default_transforms != ():
							πF.SetLineno(179)
						Label1:
							// line 180: import warnings
							πF.SetLineno(180)
							if πTemp005, πE = πg.ImportModule(πF, "warnings"); πE != nil {
								continue
							}
							πTemp001 = πTemp005[0]
							µwarnings = πTemp001
							// line 181: warnings.warn('default_transforms attribute deprecated.\n'
							πF.SetLineno(181)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("default_transforms attribute deprecated.\nUse get_transforms() method instead.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µwarnings, "warnings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µwarnings, ßwarn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 184: return list(self.default_transforms)
							πF.SetLineno(184)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_transforms, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 185: return []
							πF.SetLineno(185)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
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
					// line 178: """Transforms required by this class.  Override in subclasses."""
					πF.SetLineno(178)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Transforms required by this class.  Override in subclasses.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßget_transforms); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 188: default_transforms = ()
					πF.SetLineno(188)
					πTemp003 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßdefault_transforms.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 190: unknown_reference_resolvers = ()
					πF.SetLineno(190)
					πTemp003 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßunknown_reference_resolvers.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 191: """List of functions to try to resolve unknown references.  Unknown
					πF.SetLineno(191)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TransformSpec").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformSpec.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 215: class Component(SettingsSpec, TransformSpec):
			πF.SetLineno(215)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßSettingsSpec); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			if πTemp006, πE = πg.ResolveGlobal(πF, ßTransformSpec); πE != nil {
				continue
			}
			πTemp002[1] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Component", "/usr/lib/python2.7/site-packages/docutils/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
					// line 217: """Base class for Docutils components."""
					πF.SetLineno(217)
					// line 217: """Base class for Docutils components."""
					πF.SetLineno(217)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Base class for Docutils components.").ToObject()); πE != nil {
						continue
					}
					// line 219: component_type = None
					πF.SetLineno(219)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßcomponent_type.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 220: """Name of the component type ('reader', 'parser', 'writer').  Override in
					πF.SetLineno(220)
					// line 223: supported = ()
					πF.SetLineno(223)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 224: """Names for this component.  Override in subclasses."""
					πF.SetLineno(224)
					// line 226: def supports(self, format):
					πF.SetLineno(226)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "format", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("supports", "/usr/lib/python2.7/site-packages/docutils/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µformat *πg.Object = πArgs[1]
						_ = µformat
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 227: """
							πF.SetLineno(227)
							// line 233: return format in self.supported
							πF.SetLineno(233)
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsupported, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µformat); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßsupports.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 227: """
					πF.SetLineno(227)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Is `format` supported by this component?\n\n        To be used by transforms to ask the dependent component if it supports\n        a certain input context or output format.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßsupports); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Component").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßComponent.ToObject(), πTemp006); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("docutils", Code)
}
