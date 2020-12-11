package languages

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/languages/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßImportError := πg.InternStr("ImportError")
		ßNone := πg.InternStr("None")
		ß_ := πg.InternStr("_")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__import__ := πg.InternStr("__import__")
		ß_languages := πg.InternStr("_languages")
		ßen := πg.InternStr("en")
		ßget_language := πg.InternStr("get_language")
		ßglobals := πg.InternStr("globals")
		ßlocals := πg.InternStr("locals")
		ßnormalize_language_tag := πg.InternStr("normalize_language_tag")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreplace := πg.InternStr("replace")
		ßsys := πg.InternStr("sys")
		ßwarning := πg.InternStr("warning")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 []πg.Param
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 8: """
			πF.SetLineno(8)
			// line 8: """
			πF.SetLineno(8)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis package contains modules for language-dependent features of Docutils.\n").ToObject()); πE != nil {
				continue
			}
			// line 12: __docformat__ = 'reStructuredText'
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 14: import sys
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: from docutils.utils import normalize_language_tag
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßnormalize_language_tag); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßnormalize_language_tag.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: _languages = {}
			πF.SetLineno(19)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_languages.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: def get_language(language_code, reporter=None):
			πF.SetLineno(21)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "language_code", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp005[1] = πg.Param{Name: "reporter", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("get_language", "/usr/lib/python2.7/site-packages/docutils/languages/__init__.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlanguage_code *πg.Object = πArgs[0]
				_ = µlanguage_code
				var µreporter *πg.Object = πArgs[1]
				_ = µreporter
				var µtag *πg.Object = πg.UnboundLocal
				_ = µtag
				var µmodule *πg.Object = πg.UnboundLocal
				_ = µmodule
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
				var πTemp008 πg.KWArgs
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πTemp011 *πg.BaseException
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
					case 10:
						goto Label10
					case 7:
						goto Label7
					default:
						panic("unexpected function state")
					}
					// line 22: """Return module with language localizations.
					πF.SetLineno(22)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage_code, "language_code"); πE != nil {
						continue
					}
					πTemp002[0] = µlanguage_code
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnormalize_language_tag); πE != nil {
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
						µtag = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 29: tag = tag.replace('-', '_') # '-' not valid in module names
					πF.SetLineno(29)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr("-").ToObject()
					πTemp002[1] = ß_.ToObject()
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µtag, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µtag = πTemp004
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_languages); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, πTemp004, µtag); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					goto Label5
					// line 30: if tag in _languages:
					πF.SetLineno(30)
				Label4:
					// line 31: return _languages[tag]
					πF.SetLineno(31)
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp003 = µtag
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_languages); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					goto Label5
				Label5:
					// line 32: try:
					πF.SetLineno(32)
					πF.PushCheckpoint(7)
					// line 33: module = __import__(tag, globals(), locals(), level=1)
					πF.SetLineno(33)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp002[0] = µtag
					if πTemp003, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp004
					πTemp008 = πg.KWArgs{
						{"level", πg.NewInt(1).ToObject()},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmodule = πTemp004
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 34: except ImportError:
					πF.SetLineno(34)
				Label8:
					// line 35: try:
					πF.SetLineno(35)
					πF.PushCheckpoint(10)
					// line 36: module = __import__(tag, globals(), locals(), level=0)
					πF.SetLineno(36)
					πTemp002 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp002[0] = µtag
					if πTemp003, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp004
					πTemp008 = πg.KWArgs{
						{"level", πg.NewInt(0).ToObject()},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmodule = πTemp004
					πF.PopCheckpoint()
					goto Label9
				Label10:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp011, πTemp010 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 37: except ImportError:
					πF.SetLineno(37)
				Label11:
					// line 38: continue
					πF.SetLineno(38)
					continue
					πF.RestoreExc(nil, nil)
					goto Label9
				Label9:
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					// line 39: _languages[tag] = module
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µmodule); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_languages); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp007 = µtag
					if πE = πg.SetItem(πF, πTemp004, πTemp007, πTemp003); πE != nil {
						continue
					}
					// line 40: return module
					πF.SetLineno(40)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πR = µmodule
					continue
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µreporter != πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					goto Label13
					// line 41: if reporter is not None:
					πF.SetLineno(41)
				Label12:
					// line 42: reporter.warning(
					πF.SetLineno(42)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage_code, "language_code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("language \"%s\" not supported: ").ToObject(), µlanguage_code); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr("Docutils-generated text will be in English.").ToObject()); πE != nil {
						continue
					}
					πTemp002[0] = πTemp001
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µreporter, ßwarning, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label13
				Label13:
					// line 45: module = __import__('en', globals(), locals(), level=1)
					πF.SetLineno(45)
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = ßen.ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlocals); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[2] = πTemp003
					πTemp008 = πg.KWArgs{
						{"level", πg.NewInt(1).ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µmodule = πTemp003
					// line 46: _languages[tag] = module # warn only one time!
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmodule); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_languages); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtag, "tag"); πE != nil {
						continue
					}
					πTemp004 = µtag
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
						continue
					}
					// line 47: return module
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πR = µmodule
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßget_language.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: """Return module with language localizations.
			πF.SetLineno(22)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return module with language localizations.\n\n    `language_code` is a \"BCP 47\" language tag.\n    If there is no matching module, warn and fall back to English.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßget_language); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("languages", Code)
}
