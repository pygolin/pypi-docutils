package error_reporting

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/codecs"
	_ "github.com/pygolin/stdlib/pkg/locale"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßEnvironmentError := πg.InternStr("EnvironmentError")
		ßErrorOutput := πg.InternStr("ErrorOutput")
		ßErrorString := πg.InternStr("ErrorString")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßImportError := πg.InternStr("ImportError")
		ßLookupError := πg.InternStr("LookupError")
		ßNone := πg.InternStr("None")
		ßSafeString := πg.InternStr("SafeString")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeDecodeError := πg.InternStr("UnicodeDecodeError")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ßUnicodeError := πg.InternStr("UnicodeError")
		ßValueError := πg.InternStr("ValueError")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__str__ := πg.InternStr("__str__")
		ß__unicode__ := πg.InternStr("__unicode__")
		ßargs := πg.InternStr("args")
		ßascii := πg.InternStr("ascii")
		ßbackslashreplace := πg.InternStr("backslashreplace")
		ßbuffer := πg.InternStr("buffer")
		ßclose := πg.InternStr("close")
		ßcodecs := πg.InternStr("codecs")
		ßdata := πg.InternStr("data")
		ßdecoding_errors := πg.InternStr("decoding_errors")
		ßencode := πg.InternStr("encode")
		ßencoding := πg.InternStr("encoding")
		ßencoding_errors := πg.InternStr("encoding_errors")
		ßerrno := πg.InternStr("errno")
		ßerror := πg.InternStr("error")
		ßfilename := πg.InternStr("filename")
		ßgetattr := πg.InternStr("getattr")
		ßgetdefaultlocale := πg.InternStr("getdefaultlocale")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßgetlocale := πg.InternStr("getlocale")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlocale := πg.InternStr("locale")
		ßlocale_encoding := πg.InternStr("locale_encoding")
		ßlookup := πg.InternStr("lookup")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßreplace := πg.InternStr("replace")
		ßstderr := πg.InternStr("stderr")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßstream := πg.InternStr("stream")
		ßstrerror := πg.InternStr("strerror")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßunicode := πg.InternStr("unicode")
		ßversion_info := πg.InternStr("version_info")
		ßw := πg.InternStr("w")
		ßwrite := πg.InternStr("write")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 bool
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.BaseException
		_ = πTemp008
		var πTemp009 *πg.Traceback
		_ = πTemp009
		var πTemp010 *πg.Dict
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 12:
				goto Label12
			case 2:
				goto Label2
			case 4:
				goto Label4
			default:
				panic("unexpected function state")
			}
			// line 15: """
			πF.SetLineno(15)
			// line 15: """
			πF.SetLineno(15)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nError reporting should be safe from encoding/decoding errors.\nHowever, implicit conversions of strings and exceptions like\n\n>>> u'%s world: %s' % ('H\xe4llo', Exception(u'H\xe4llo')\n\nfail in some Python versions:\n\n* In Python <= 2.6, ``unicode(<exception instance>)`` uses\n  `__str__` and fails with non-ASCII chars in`unicode` arguments.\n  (work around http://bugs.python.org/issue2517):\n\n* In Python 2, unicode(<exception instance>) fails, with non-ASCII\n  chars in arguments. (Use case: in some locales, the errstr\n  argument of IOError contains non-ASCII chars.)\n\n* In Python 2, str(<exception instance>) fails, with non-ASCII chars\n  in `unicode` arguments.\n\nThe `SafeString`, `ErrorString` and `ErrorOutput` classes handle\ncommon exceptions.\n").ToObject()); πE != nil {
				continue
			}
			// line 38: import codecs
			πF.SetLineno(38)
			if πTemp002, πE = πg.ImportModule(πF, "codecs"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcodecs.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 39: import sys
			πF.SetLineno(39)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 43: try:
			πF.SetLineno(43)
			πF.PushCheckpoint(2)
			// line 44: import locale # module missing in Jython
			πF.SetLineno(44)
			if πTemp002, πE = πg.ImportModule(πF, "locale"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlocale.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			// line 48: try:
			πF.SetLineno(48)
			πF.PushCheckpoint(4)
			// line 49: locale_encoding = locale.getlocale()[1] or locale.getdefaultlocale()[1]
			πF.SetLineno(49)
			πTemp004 = πg.NewInt(1).ToObject()
			if πTemp006, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßgetlocale, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
				continue
			}
			πTemp001 = πTemp005
			if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label5
			}
			πTemp004 = πg.NewInt(1).ToObject()
			if πTemp006, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßgetdefaultlocale, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
				continue
			}
			πTemp001 = πTemp005
		Label5:
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label3
		Label4:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp008, πTemp009 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
				continue
			}
			if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label6
			}
			goto Label7
			// line 53: except ValueError as error: # OS X may set UTF-8 without language code
			πF.SetLineno(53)
		Label6:
			if πE = πF.Globals().SetItem(πF, ßerror.ToObject(), πTemp008.ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßargs, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Contains(πF, πTemp005, πg.NewStr("unknown locale: UTF-8").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp003).ToObject()
			if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label8
			}
			goto Label9
			// line 56: if "unknown locale: UTF-8" in error.args:
			πF.SetLineno(56)
		Label8:
			// line 57: locale_encoding = "UTF-8"
			πF.SetLineno(57)
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πg.NewStr("UTF-8").ToObject()); πE != nil {
				continue
			}
			goto Label10
		Label9:
			// line 59: locale_encoding = None
			πF.SetLineno(59)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label10
		Label10:
			πF.RestoreExc(nil, nil)
			goto Label3
			// line 60: except: # any other problems determining the locale -> use None
			πF.SetLineno(60)
		Label7:
			// line 61: locale_encoding = None
			πF.SetLineno(61)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label3
		Label3:
			// line 62: try:
			πF.SetLineno(62)
			πF.PushCheckpoint(12)
			// line 63: codecs.lookup(locale_encoding or '') # None -> ''
			πF.SetLineno(63)
			πTemp002 = πF.MakeArgs(1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
				continue
			}
			πTemp001 = πTemp004
			if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label13
			}
			πTemp001 = ß.ToObject()
		Label13:
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßlookup, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πF.PopCheckpoint()
			goto Label11
		Label12:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp008, πTemp009 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
				continue
			}
			if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label14
			}
			πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
			continue
			// line 64: except LookupError:
			πF.SetLineno(64)
		Label14:
			// line 65: locale_encoding = None
			πF.SetLineno(65)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label11
		Label11:
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp008, πTemp009 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label15
			}
			πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
			continue
			// line 45: except ImportError:
			πF.SetLineno(45)
		Label15:
			// line 46: locale_encoding = None
			πF.SetLineno(46)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp005, πTemp004); πE != nil {
				continue
			}
			if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp003 {
				goto Label16
			}
			goto Label17
			// line 68: if sys.version_info >= (3, 0):
			πF.SetLineno(68)
		Label16:
			// line 69: unicode = str  # noqa
			πF.SetLineno(69)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label17
		Label17:
			// line 72: class SafeString(object):
			πF.SetLineno(72)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SafeString", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
					// line 73: """
					πF.SetLineno(73)
					// line 73: """
					πF.SetLineno(73)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A wrapper providing robust conversion to `str` and `unicode`.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 77: def __init__(self, data, encoding=None, encoding_errors='backslashreplace',
					πF.SetLineno(77)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "encoding", Def: πTemp003}
					πTemp002[3] = πg.Param{Name: "encoding_errors", Def: ßbackslashreplace.ToObject()}
					πTemp002[4] = πg.Param{Name: "decoding_errors", Def: ßreplace.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µencoding *πg.Object = πArgs[2]
						_ = µencoding
						var µencoding_errors *πg.Object = πArgs[3]
						_ = µencoding_errors
						var µdecoding_errors *πg.Object = πArgs[4]
						_ = µdecoding_errors
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 79: self.data = data
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdata); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp001); πE != nil {
								continue
							}
							// line 80: self.encoding = (encoding or getattr(data, 'encoding', None) or
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001 = µencoding
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003[0] = µdata
							πTemp003[1] = ßencoding.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp005
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp001 = ßascii.ToObject()
						Label1:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp004); πE != nil {
								continue
							}
							// line 82: self.encoding_errors = encoding_errors
							πF.SetLineno(82)
							if πE = πg.CheckLocal(πF, µencoding_errors, "encoding_errors"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µencoding_errors); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding_errors, πTemp001); πE != nil {
								continue
							}
							// line 83: self.decoding_errors = decoding_errors
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µdecoding_errors, "decoding_errors"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdecoding_errors); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdecoding_errors, πTemp001); πE != nil {
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
					// line 86: def __str__(self):
					πF.SetLineno(86)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πTemp008 *πg.Object
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
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 87: try:
							πF.SetLineno(87)
							πF.PushCheckpoint(2)
							// line 88: return str(self.data)
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 89: except UnicodeEncodeError:
							πF.SetLineno(89)
						Label3:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 90: if isinstance(self.data, Exception):
							πF.SetLineno(90)
						Label4:
							// line 91: args = [str(SafeString(arg, self.encoding,
							πF.SetLineno(91)
							πTemp007 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µarg *πg.Object = πg.UnboundLocal
								_ = µarg
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßargs, nil); πE != nil {
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
											µarg = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 91: args = [str(SafeString(arg, self.encoding,
										πF.SetLineno(91)
										πTemp006 = πF.MakeArgs(1)
										πTemp007 = πF.MakeArgs(3)
										if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
											continue
										}
										πTemp007[0] = µarg
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
											continue
										}
										πTemp007[1] = πTemp002
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding_errors, nil); πE != nil {
											continue
										}
										πTemp007[2] = πTemp002
										if πTemp002, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										πTemp006[0] = πTemp003
										if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp002 = πSent
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
							if πTemp008, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							// line 94: return ', '.join(args)
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp008
							continue
							goto Label5
						Label5:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 95: if isinstance(self.data, unicode):
							πF.SetLineno(95)
						Label6:
							if πTemp008, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp008 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp002, πE = πg.GT(πF, πTemp009, πTemp008); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 96: if sys.version_info > (3, 0):
							πF.SetLineno(96)
						Label8:
							// line 97: return self.data
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label10
						Label9:
							// line 99: return self.data.encode(self.encoding,
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding_errors, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp002, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label10
						Label10:
							goto Label7
						Label7:
							// line 101: raise
							πF.SetLineno(101)
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 103: def __unicode__(self):
					πF.SetLineno(103)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__unicode__", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µu *πg.Object = πg.UnboundLocal
						_ = µu
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []πg.Param
						_ = πTemp011
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
							// line 104: """
							πF.SetLineno(104)
							// line 115: try:
							πF.SetLineno(115)
							πF.PushCheckpoint(2)
							// line 116: u = unicode(self.data)
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEnvironmentError); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label3
							}
							goto Label4
							// line 117: if isinstance(self.data, EnvironmentError):
							πF.SetLineno(117)
						Label3:
							// line 118: u = u.replace(": u'", ": '") # normalize filename quoting
							πF.SetLineno(118)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr(": u'").ToObject()
							πTemp001[1] = πg.NewStr(": '").ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µu, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp003
							goto Label4
						Label4:
							// line 119: return u
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πR = µu
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 120: except UnicodeError as error: # catch ..Encode.. and ..Decode.. errors
							πF.SetLineno(120)
						Label5:
							µerror = πTemp005.ToObject()
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEnvironmentError); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 121: if isinstance(self.data, EnvironmentError):
							πF.SetLineno(121)
						Label6:
							// line 122: return  u"[Errno %s] %s: '%s'" % (self.data.errno,
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßerrno, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßstrerror, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdecoding_errors, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp007, ßfilename, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßdecoding_errors, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple3(πTemp008, πTemp009, πTemp010).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("[Errno %s] %s: '%s'").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label7
						Label7:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
								goto Label8
							}
							goto Label9
							// line 127: if isinstance(self.data, Exception):
							πF.SetLineno(127)
						Label8:
							// line 128: args = [unicode(SafeString(arg, self.encoding,
							πF.SetLineno(128)
							πTemp011 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µarg *πg.Object = πg.UnboundLocal
								_ = µarg
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
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßargs, nil); πE != nil {
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
											µarg = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 128: args = [unicode(SafeString(arg, self.encoding,
										πF.SetLineno(128)
										πTemp006 = πF.MakeArgs(1)
										πTemp007 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
											continue
										}
										πTemp007[0] = µarg
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
											continue
										}
										πTemp007[1] = πTemp002
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßdecoding_errors, nil); πE != nil {
											continue
										}
										πTemp008 = πg.KWArgs{
											{"decoding_errors", πTemp002},
										}
										if πTemp002, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp007, πTemp008); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										πTemp006[0] = πTemp003
										if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp002 = πSent
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
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							// line 131: return u', '.join(args)
							πF.SetLineno(131)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πTemp002, πE = πg.GetAttr(πF, πg.NewUnicode(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp007
							continue
							goto Label9
						Label9:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 132: if isinstance(error, UnicodeDecodeError):
							πF.SetLineno(132)
						Label10:
							// line 133: return unicode(self.data, self.encoding, self.decoding_errors)
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdecoding_errors, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp007
							continue
							goto Label11
						Label11:
							// line 134: raise
							πF.SetLineno(134)
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
					if πE = πClass.SetItem(πF, ß__unicode__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 104: """
					πF.SetLineno(104)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Return unicode representation of `self.data`.\n\n        Try ``unicode(self.data)``, catch `UnicodeError` and\n\n        * if `self.data` is an Exception instance, work around\n          http://bugs.python.org/issue2517 with an emulation of\n          Exception.__unicode__,\n\n        * else decode with `self.encoding` and `self.decoding_errors`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__unicode__); πE != nil {
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
			if πTemp004, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("SafeString").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSafeString.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 136: class ErrorString(SafeString):
			πF.SetLineno(136)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ErrorString", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
					// line 137: """
					πF.SetLineno(137)
					// line 137: """
					πF.SetLineno(137)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Safely report exception type and message.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 140: def __str__(self):
					πF.SetLineno(140)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 141: return '%s: %s' % (self.data.__class__.__name__,
							πF.SetLineno(141)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ß__str__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 144: def __unicode__(self):
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__unicode__", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 145: return u'%s: %s' % (self.data.__class__.__name__,
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ß__unicode__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("%s: %s").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__unicode__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("ErrorString").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorString.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 149: class ErrorOutput(object):
			πF.SetLineno(149)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ErrorOutput", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
					// line 150: """
					πF.SetLineno(150)
					// line 150: """
					πF.SetLineno(150)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Wrapper class for file-like error streams with\n    failsave de- and encoding of `str`, `bytes`, `unicode` and\n    `Exception` instances.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 156: def __init__(self, stream=None, encoding=None,
					πF.SetLineno(156)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "stream", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "encoding", Def: πTemp003}
					πTemp002[3] = πg.Param{Name: "encoding_errors", Def: ßbackslashreplace.ToObject()}
					πTemp002[4] = πg.Param{Name: "decoding_errors", Def: ßreplace.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstream *πg.Object = πArgs[1]
						_ = µstream
						var µencoding *πg.Object = πArgs[2]
						_ = µencoding
						var µencoding_errors *πg.Object = πArgs[3]
						_ = µencoding_errors
						var µdecoding_errors *πg.Object = πArgs[4]
						_ = µdecoding_errors
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 159: """
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstream == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µstream); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp004[0] = µstream
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp004[0] = µstream
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 168: if stream is None:
							πF.SetLineno(168)
						Label1:
							// line 169: stream = sys.stderr
							πF.SetLineno(169)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
								continue
							}
							µstream = πTemp002
							goto Label5
							// line 170: elif not(stream):
							πF.SetLineno(170)
						Label2:
							// line 171: stream = False
							πF.SetLineno(171)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µstream = πTemp001
							goto Label5
							// line 173: elif isinstance(stream, str):
							πF.SetLineno(173)
						Label3:
							// line 174: stream = open(stream, 'w')
							πF.SetLineno(174)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp004[0] = µstream
							πTemp004[1] = ßw.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstream = πTemp002
							goto Label5
							// line 175: elif isinstance(stream, unicode):
							πF.SetLineno(175)
						Label4:
							// line 176: stream = open(stream.encode(sys.getfilesystemencoding()), 'w')
							πF.SetLineno(176)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetfilesystemencoding, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstream, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							πTemp004[1] = ßw.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstream = πTemp002
							goto Label5
						Label5:
							// line 178: self.stream = stream
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstream); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstream, πTemp001); πE != nil {
								continue
							}
							// line 179: """Where warning output is sent."""
							πF.SetLineno(179)
							// line 181: self.encoding = (encoding or getattr(stream, 'encoding', None) or
							πF.SetLineno(181)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001 = µencoding
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
								continue
							}
							πTemp004[0] = µstream
							πTemp004[1] = ßencoding.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001 = πTemp006
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πTemp001 = ßascii.ToObject()
						Label6:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp002); πE != nil {
								continue
							}
							// line 183: """The output character encoding."""
							πF.SetLineno(183)
							// line 185: self.encoding_errors = encoding_errors
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µencoding_errors, "encoding_errors"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µencoding_errors); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding_errors, πTemp001); πE != nil {
								continue
							}
							// line 186: """Encoding error handler."""
							πF.SetLineno(186)
							// line 188: self.decoding_errors = decoding_errors
							πF.SetLineno(188)
							if πE = πg.CheckLocal(πF, µdecoding_errors, "decoding_errors"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdecoding_errors); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdecoding_errors, πTemp001); πE != nil {
								continue
							}
							// line 189: """Decoding error handler."""
							πF.SetLineno(189)
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
					// line 159: """
					πF.SetLineno(159)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        :Parameters:\n            - `stream`: a file-like object,\n                        a string (path to a file),\n                        `None` (write to `sys.stderr`, default), or\n                        evaluating to `False` (write() requests are ignored).\n            - `encoding`: `stream` text encoding. Guessed if None.\n            - `encoding_errors`: how to treat encoding errors.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 191: def write(self, data):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 192: """
							πF.SetLineno(192)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 197: if self.stream is False:
							πF.SetLineno(197)
						Label1:
							// line 198: return
							πF.SetLineno(198)
							πR = πg.None
							continue
							goto Label2
						Label2:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp005[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 199: if isinstance(data, Exception):
							πF.SetLineno(199)
						Label3:
							// line 200: data = unicode(SafeString(data, self.encoding,
							πF.SetLineno(200)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding_errors, nil); πE != nil {
								continue
							}
							πTemp006[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdecoding_errors, nil); πE != nil {
								continue
							}
							πTemp006[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µdata = πTemp002
							goto Label4
						Label4:
							// line 202: try:
							πF.SetLineno(202)
							πF.PushCheckpoint(6)
							// line 203: self.stream.write(data)
							πF.SetLineno(203)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp005[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 204: except UnicodeEncodeError:
							πF.SetLineno(204)
						Label7:
							// line 205: self.stream.write(data.encode(self.encoding, self.encoding_errors))
							πF.SetLineno(205)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding_errors, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.RestoreExc(nil, nil)
							goto Label5
							// line 206: except TypeError:
							πF.SetLineno(206)
						Label8:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp005[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 207: if isinstance(data, unicode): # passed stream may expect bytes
							πF.SetLineno(207)
						Label9:
							// line 208: self.stream.write(data.encode(self.encoding,
							πF.SetLineno(208)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding_errors, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 210: return
							πF.SetLineno(210)
							πR = πg.None
							continue
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßstdout, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp010, πTemp011).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 211: if self.stream in (sys.stderr, sys.stdout):
							πF.SetLineno(211)
						Label11:
							// line 212: self.stream.buffer.write(data) # write bytes to raw stream
							πF.SetLineno(212)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp005[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßbuffer, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label13
						Label12:
							// line 214: self.stream.write(unicode(data, self.encoding,
							πF.SetLineno(214)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdecoding_errors, nil); πE != nil {
								continue
							}
							πTemp006[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label13
						Label13:
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 192: """
					πF.SetLineno(192)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Write `data` to self.stream. Ignore, if self.stream is False.\n\n        `data` can be a `string`, `unicode`, or `Exception` instance.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßwrite); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 217: def close(self):
					πF.SetLineno(217)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/site-packages/docutils/utils/error_reporting.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 218: """
							πF.SetLineno(218)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßstderr, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πTemp007, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 224: if self.stream in (sys.stdout, sys.stderr):
							πF.SetLineno(224)
						Label1:
							// line 225: return
							πF.SetLineno(225)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 226: try:
							πF.SetLineno(226)
							πF.PushCheckpoint(4)
							// line 227: self.stream.close()
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstream, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 228: except AttributeError:
							πF.SetLineno(228)
						Label5:
							// line 229: pass
							πF.SetLineno(229)
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 218: """
					πF.SetLineno(218)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Close the error-output stream.\n\n        Ignored if the stream is` sys.stderr` or `sys.stdout` or has no\n        close() method.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßclose); πE != nil {
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
			if πTemp004, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("ErrorOutput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorOutput.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("error_reporting", Code)
}
