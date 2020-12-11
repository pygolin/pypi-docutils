package code_analyzer

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/pkg_resources"
	_ "github.com/pygolin/stdlib/pkg/pygments"
	_ "github.com/pygolin/stdlib/pkg/pygments/formatters"
	_ "github.com/pygolin/stdlib/pkg/pygments/formatters/html"
	_ "github.com/pygolin/stdlib/pkg/pygments/lexers"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßApplicationError := πg.InternStr("ApplicationError")
		ßClassNotFound := πg.InternStr("ClassNotFound")
		ßDistributionNotFound := πg.InternStr("DistributionNotFound")
		ßFalse := πg.InternStr("False")
		ßImportError := πg.InternStr("ImportError")
		ßLexer := πg.InternStr("Lexer")
		ßLexerError := πg.InternStr("LexerError")
		ßNone := πg.InternStr("None")
		ßNumberLines := πg.InternStr("NumberLines")
		ßResourceError := πg.InternStr("ResourceError")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_get_ttype_class := πg.InternStr("_get_ttype_class")
		ßcode := πg.InternStr("code")
		ßendswith := πg.InternStr("endswith")
		ßfmt_str := πg.InternStr("fmt_str")
		ßget_lexer_by_name := πg.InternStr("get_lexer_by_name")
		ßiter := πg.InternStr("iter")
		ßlanguage := πg.InternStr("language")
		ßlen := πg.InternStr("len")
		ßlex := πg.InternStr("lex")
		ßlexer := πg.InternStr("lexer")
		ßln := πg.InternStr("ln")
		ßlong := πg.InternStr("long")
		ßlower := πg.InternStr("lower")
		ßmerge := πg.InternStr("merge")
		ßnext := πg.InternStr("next")
		ßnone := πg.InternStr("none")
		ßobject := πg.InternStr("object")
		ßpygments := πg.InternStr("pygments")
		ßshort := πg.InternStr("short")
		ßsplit := πg.InternStr("split")
		ßstartline := πg.InternStr("startline")
		ßstr := πg.InternStr("str")
		ßtext := πg.InternStr("text")
		ßtoken := πg.InternStr("token")
		ßtokennames := πg.InternStr("tokennames")
		ßtokens := πg.InternStr("tokens")
		ßunstyled_tokens := πg.InternStr("unstyled_tokens")
		ßutil := πg.InternStr("util")
		ßwith_pygments := πg.InternStr("with_pygments")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.BaseException
		_ = πTemp004
		var πTemp005 *πg.Traceback
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 bool
		_ = πTemp007
		var πTemp008 *πg.Dict
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2:
				goto Label2
			case 5:
				goto Label5
			default:
				panic("unexpected function state")
			}
			// line 4: """Lexical analysis of formal languages (i.e. code) using Pygments."""
			πF.SetLineno(4)
			// line 4: """Lexical analysis of formal languages (i.e. code) using Pygments."""
			πF.SetLineno(4)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Lexical analysis of formal languages (i.e. code) using Pygments.").ToObject()); πE != nil {
				continue
			}
			// line 10: from docutils import ApplicationError
			πF.SetLineno(10)
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
			// line 11: try:
			πF.SetLineno(11)
			πF.PushCheckpoint(2)
			// line 12: from pkg_resources import DistributionNotFound as ResourceError
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "pkg_resources"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDistributionNotFound); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßResourceError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
			if πTemp007, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label3
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 13: except (ImportError, RuntimeError):
			πF.SetLineno(13)
		Label3:
			// line 14: class ResourceError(ApplicationError):
			πF.SetLineno(14)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßApplicationError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ResourceError", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 15: pass # stub
					πF.SetLineno(15)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ResourceError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßResourceError.ToObject(), πTemp006); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 16: try:
			πF.SetLineno(16)
			πF.PushCheckpoint(5)
			// line 17: import pygments
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "pygments"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßpygments.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: from pygments.lexers import get_lexer_by_name
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "pygments.lexers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßget_lexer_by_name); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßget_lexer_by_name.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from pygments.formatters.html import _get_ttype_class
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "pygments.formatters.html"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_get_ttype_class); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_get_ttype_class.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: with_pygments = True
			πF.SetLineno(20)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwith_pygments.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label6
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 21: except ImportError:
			πF.SetLineno(21)
		Label6:
			// line 22: with_pygments = False
			πF.SetLineno(22)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwith_pygments.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 25: unstyled_tokens = ['token', # Token (base token type)
			πF.SetLineno(25)
			πTemp002 = make([]*πg.Object, 3)
			πTemp002[0] = ßtoken.ToObject()
			πTemp002[1] = ßtext.ToObject()
			πTemp002[2] = ß.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßunstyled_tokens.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 30: class LexerError(ApplicationError):
			πF.SetLineno(30)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßApplicationError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LexerError", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 31: pass
					πF.SetLineno(31)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("LexerError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLexerError.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 33: class Lexer(object):
			πF.SetLineno(33)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Lexer", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
					// line 34: """Parse `code` lines and yield "classified" tokens.
					πF.SetLineno(34)
					// line 34: """Parse `code` lines and yield "classified" tokens.
					πF.SetLineno(34)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Parse `code` lines and yield \"classified\" tokens.\n\n    Arguments\n\n      code       -- string of source code to parse,\n      language   -- formal language the code is written in,\n      tokennames -- either 'long', 'short', or '' (see below).\n\n    Merge subsequent tokens of the same token-type.\n\n    Iterating over an instance yields the tokens as ``(tokentype, value)``\n    tuples. The value of `tokennames` configures the naming of the tokentype:\n\n      'long':  downcased full token type name,\n      'short': short name defined by pygments.token.STANDARD_TYPES\n               (= class argument used in pygments html output),\n      'none':      skip lexical analysis.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 53: def __init__(self, code, language, tokennames='short'):
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "code", Def: nil}
					πTemp002[2] = πg.Param{Name: "language", Def: nil}
					πTemp002[3] = πg.Param{Name: "tokennames", Def: ßshort.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcode *πg.Object = πArgs[1]
						_ = µcode
						var µlanguage *πg.Object = πArgs[2]
						_ = µlanguage
						var µtokennames *πg.Object = πArgs[3]
						_ = µtokennames
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 54: """
							πF.SetLineno(54)
							// line 57: self.code = code
							πF.SetLineno(57)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcode, πTemp001); πE != nil {
								continue
							}
							// line 58: self.language = language
							πF.SetLineno(58)
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlanguage); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage, πTemp001); πE != nil {
								continue
							}
							// line 59: self.tokennames = tokennames
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µtokennames, "tokennames"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtokennames); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtokennames, πTemp001); πE != nil {
								continue
							}
							// line 60: self.lexer = None
							πF.SetLineno(60)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlexer, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(ß.ToObject(), ßtext.ToObject()).ToObject()
							if πTemp005, πE = πg.Contains(πF, πTemp004, µlanguage); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µtokennames, "tokennames"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µtokennames, ßnone.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 62: if language in ('', 'text') or tokennames == 'none':
							πF.SetLineno(62)
						Label2:
							// line 63: return
							πF.SetLineno(63)
							πR = πg.None
							continue
							goto Label3
						Label3:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwith_pygments); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 64: if not with_pygments:
							πF.SetLineno(64)
						Label4:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("Cannot analyze code. Pygments package not found.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßLexerError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 65: raise LexerError('Cannot analyze code. '
							πF.SetLineno(65)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label5
						Label5:
							// line 67: try:
							πF.SetLineno(67)
							πF.PushCheckpoint(7)
							// line 68: self.lexer = get_lexer_by_name(self.language)
							πF.SetLineno(68)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlanguage, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßget_lexer_by_name); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlexer, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpygments); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßutil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßClassNotFound, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßResourceError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 69: except (pygments.util.ClassNotFound, ResourceError):
							πF.SetLineno(69)
						Label8:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Cannot analyze code. No Pygments lexer found for \"%s\".").ToObject(), µlanguage); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßLexerError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 70: raise LexerError('Cannot analyze code. '
							πF.SetLineno(70)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 54: """
					πF.SetLineno(54)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Set up a lexical analyzer for `code` in `language`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 80: def merge(self, tokens):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tokens", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("merge", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtokens *πg.Object = πArgs[1]
						_ = µtokens
						var µlasttype *πg.Object = πg.UnboundLocal
						_ = µlasttype
						var µlastval *πg.Object = πg.UnboundLocal
						_ = µlastval
						var µttype *πg.Object = πg.UnboundLocal
						_ = µttype
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
								case 12:
									goto Label12
								case 7:
									goto Label7
								default:
									panic("unexpected function state")
								}
								// line 81: """Merge subsequent tokens of same token-type.
								πF.SetLineno(81)
								// line 85: tokens = iter(tokens)
								πF.SetLineno(85)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
									continue
								}
								πTemp001[0] = µtokens
								if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µtokens = πTemp003
								// line 86: (lasttype, lastval) = next(tokens)
								πF.SetLineno(86)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
									continue
								}
								πTemp001[0] = µtokens
								if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
									continue
								}
								µlasttype = πTemp002
								µlastval = πTemp004
								if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Iter(πF, µtokens); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
										continue
									}
									µttype = πTemp004
									µvalue = πTemp007
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlasttype, "lasttype"); πE != nil {
									continue
								}
								πTemp003 = πg.GetBool(µttype == µlasttype).ToObject()
								if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 88: if ttype is lasttype:
								πF.SetLineno(88)
							Label4:
								// line 89: lastval += value
								πF.SetLineno(89)
								if πE = πg.CheckLocal(πF, µlastval, "lastval"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µlastval, µvalue); πE != nil {
									continue
								}
								µlastval = πTemp003
								goto Label6
							Label5:
								// line 91: yield(lasttype, lastval)
								πF.SetLineno(91)
								if πE = πg.CheckLocal(πF, µlasttype, "lasttype"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlastval, "lastval"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µlasttype, µlastval).ToObject()
								πF.PushCheckpoint(7)
								return πTemp003, nil
							Label7:
								πTemp004 = πSent
								// line 92: (lasttype, lastval) = (ttype, value)
								πF.SetLineno(92)
								if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µttype, µvalue).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
									continue
								}
								µlasttype = πTemp007
								µlastval = πTemp008
								goto Label6
							Label6:
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewStr("\n").ToObject()
								if πE = πg.CheckLocal(πF, µlastval, "lastval"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µlastval, ßendswith, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label8
								}
								goto Label9
								// line 93: if lastval.endswith('\n'):
								πF.SetLineno(93)
							Label8:
								// line 94: lastval = lastval[:-1]
								πF.SetLineno(94)
								if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlastval, "lastval"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µlastval, πTemp002); πE != nil {
									continue
								}
								µlastval = πTemp004
								goto Label9
							Label9:
								if πE = πg.CheckLocal(πF, µlastval, "lastval"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, µlastval); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label10
								}
								goto Label11
								// line 95: if lastval:
								πF.SetLineno(95)
							Label10:
								// line 96: yield(lasttype, lastval)
								πF.SetLineno(96)
								if πE = πg.CheckLocal(πF, µlasttype, "lasttype"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlastval, "lastval"); πE != nil {
									continue
								}
								πTemp002 = πg.NewTuple2(µlasttype, µlastval).ToObject()
								πF.PushCheckpoint(12)
								return πTemp002, nil
							Label12:
								πTemp004 = πSent
								goto Label11
							Label11:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßmerge.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 81: """Merge subsequent tokens of same token-type.
					πF.SetLineno(81)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Merge subsequent tokens of same token-type.\n\n           Also strip the final newline (added by pygments).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßmerge); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 98: def __iter__(self):
					πF.SetLineno(98)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtokens *πg.Object = πg.UnboundLocal
						_ = µtokens
						var µtokentype *πg.Object = πg.UnboundLocal
						_ = µtokentype
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µclasses *πg.Object = πg.UnboundLocal
						_ = µclasses
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []πg.Param
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 10:
									goto Label10
								case 3:
									goto Label3
								case 4:
									goto Label4
								case 5:
									goto Label5
								default:
									panic("unexpected function state")
								}
								// line 99: """Parse self.code and yield "classified" tokens.
								πF.SetLineno(99)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßlexer, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
								// line 101: if self.lexer is None:
								πF.SetLineno(101)
							Label1:
								// line 102: yield ([], self.code)
								πF.SetLineno(102)
								πTemp005 = make([]*πg.Object, 0)
								πTemp002 = πg.NewList(πTemp005...).ToObject()
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßcode, nil); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
								πF.PushCheckpoint(3)
								return πTemp001, nil
							Label3:
								πTemp002 = πSent
								// line 103: return
								πF.SetLineno(103)
								πR = πg.None
								continue
								goto Label2
							Label2:
								// line 104: tokens = pygments.lex(self.code, self.lexer)
								πF.SetLineno(104)
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßcode, nil); πE != nil {
									continue
								}
								πTemp005[0] = πTemp002
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßlexer, nil); πE != nil {
									continue
								}
								πTemp005[1] = πTemp002
								if πTemp002, πE = πg.ResolveGlobal(πF, ßpygments); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlex, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µtokens = πTemp002
								πTemp005 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
									continue
								}
								πTemp005[0] = µtokens
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßmerge, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								πTemp004 = false
							Label4:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
										continue
									}
									µtokentype = πTemp006
									µvalue = πTemp008
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(4)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßtokennames, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, πTemp006, ßlong.ToObject()); πE != nil {
									continue
								}
								if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp007 {
									goto Label7
								}
								goto Label8
								// line 106: if self.tokennames == 'long': # long CSS class args
								πF.SetLineno(106)
							Label7:
								// line 107: classes = str(tokentype).lower().split('.')
								πF.SetLineno(107)
								πTemp005 = πF.MakeArgs(1)
								πTemp005[0] = πg.NewStr(".").ToObject()
								πTemp009 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtokentype, "tokentype"); πE != nil {
									continue
								}
								πTemp009[0] = µtokentype
								if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp009)
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßlower, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßsplit, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								µclasses = πTemp006
								goto Label9
							Label8:
								// line 109: classes = [_get_ttype_class(tokentype)]
								πF.SetLineno(109)
								πTemp005 = make([]*πg.Object, 1)
								πTemp009 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtokentype, "tokentype"); πE != nil {
									continue
								}
								πTemp009[0] = µtokentype
								if πTemp003, πE = πg.ResolveGlobal(πF, ß_get_ttype_class); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp009)
								πTemp005[0] = πTemp006
								πTemp003 = πg.NewList(πTemp005...).ToObject()
								µclasses = πTemp003
								goto Label9
							Label9:
								// line 110: classes = [cls for cls in classes if cls not in unstyled_tokens]
								πF.SetLineno(110)
								πTemp010 = make([]πg.Param, 0)
								πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
									var µcls *πg.Object = πg.UnboundLocal
									_ = µcls
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
											case 6:
												goto Label6
											default:
												panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Iter(πF, µclasses); πE != nil {
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
												µcls = πTemp004
											}
											if πE != nil || !πTemp003 {
												continue
											}
											πF.PushCheckpoint(1)
											if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
												continue
											}
											if πTemp005, πE = πg.ResolveGlobal(πF, ßunstyled_tokens); πE != nil {
												continue
											}
											if πTemp003, πE = πg.Contains(πF, πTemp005, µcls); πE != nil {
												continue
											}
											πTemp004 = πg.GetBool(!πTemp003).ToObject()
											if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label4
											}
											goto Label5
											// line 110: classes = [cls for cls in classes if cls not in unstyled_tokens]
											πF.SetLineno(110)
										Label4:
											// line 110: classes = [cls for cls in classes if cls not in unstyled_tokens]
											πF.SetLineno(110)
											if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
												continue
											}
											πF.PushCheckpoint(6)
											return µcls, nil
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
								if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
									continue
								}
								µclasses = πTemp003
								// line 111: yield (classes, value)
								πF.SetLineno(111)
								if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple2(µclasses, µvalue).ToObject()
								πF.PushCheckpoint(10)
								return πTemp003, nil
							Label10:
								πTemp008 = πSent
								continue
							Label5:
								if πE != nil || πR != nil {
									continue
								}
							Label6:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 99: """Parse self.code and yield "classified" tokens.
					πF.SetLineno(99)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Parse self.code and yield \"classified\" tokens.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__iter__); πE != nil {
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
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Lexer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLexer.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 114: class NumberLines(object):
			πF.SetLineno(114)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NumberLines", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
					// line 115: """Insert linenumber-tokens at the start of every code line.
					πF.SetLineno(115)
					// line 115: """Insert linenumber-tokens at the start of every code line.
					πF.SetLineno(115)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Insert linenumber-tokens at the start of every code line.\n\n    Arguments\n\n       tokens    -- iterable of ``(classes, value)`` tuples\n       startline -- first line number\n       endline   -- last line number\n\n    Iterating over an instance yields the tokens with a\n    ``(['ln'], '<the line number>')`` token added for every code line.\n    Multi-line tokens are splitted.").ToObject()); πE != nil {
						continue
					}
					// line 127: def __init__(self, tokens, startline, endline):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tokens", Def: nil}
					πTemp002[2] = πg.Param{Name: "startline", Def: nil}
					πTemp002[3] = πg.Param{Name: "endline", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtokens *πg.Object = πArgs[1]
						_ = µtokens
						var µstartline *πg.Object = πArgs[2]
						_ = µstartline
						var µendline *πg.Object = πArgs[3]
						_ = µendline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 128: self.tokens = tokens
							πF.SetLineno(128)
							if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtokens); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtokens, πTemp001); πE != nil {
								continue
							}
							// line 129: self.startline = startline
							πF.SetLineno(129)
							if πE = πg.CheckLocal(πF, µstartline, "startline"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstartline); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstartline, πTemp001); πE != nil {
								continue
							}
							// line 131: self.fmt_str = '%%%dd ' % len(str(endline))
							πF.SetLineno(131)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µendline, "endline"); πE != nil {
								continue
							}
							πTemp003[0] = µendline
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%%%dd ").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfmt_str, πTemp004); πE != nil {
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
					// line 133: def __iter__(self):
					πF.SetLineno(133)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/site-packages/docutils/utils/code_analyzer.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlineno *πg.Object = πg.UnboundLocal
						_ = µlineno
						var µttype *πg.Object = πg.UnboundLocal
						_ = µttype
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µlines *πg.Object = πg.UnboundLocal
						_ = µlines
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
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
						var πTemp009 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1:
									goto Label1
								case 2:
									goto Label2
								case 3:
									goto Label3
								case 5:
									goto Label5
								case 6:
									goto Label6
								case 8:
									goto Label8
								case 9:
									goto Label9
								case 10:
									goto Label10
								default:
									panic("unexpected function state")
								}
								// line 134: lineno = self.startline
								πF.SetLineno(134)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßstartline, nil); πE != nil {
									continue
								}
								µlineno = πTemp001
								// line 135: yield (['ln'], self.fmt_str % lineno)
								πF.SetLineno(135)
								πTemp002 = make([]*πg.Object, 1)
								πTemp002[0] = ßln.ToObject()
								πTemp003 = πg.NewList(πTemp002...).ToObject()
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µself, ßfmt_str, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Mod(πF, πTemp005, µlineno); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
								πF.PushCheckpoint(1)
								return πTemp001, nil
							Label1:
								πTemp003 = πSent
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µself, ßtokens, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								πTemp006 = false
							Label2:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp006 {
									πF.PopCheckpoint()
									goto Label4
								}
								if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
										continue
									}
									µttype = πTemp005
									µvalue = πTemp008
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(2)
								// line 137: lines = value.split('\n')
								πF.SetLineno(137)
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewStr("\n").ToObject()
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µvalue, ßsplit, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								µlines = πTemp005
								if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp008, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetItem(πF, µlines, πTemp005); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Iter(πF, πTemp008); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								πTemp007 = false
							Label5:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp007 {
									πF.PopCheckpoint()
									goto Label7
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
									µline = πTemp005
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(5)
								// line 139: yield (ttype, line + '\n')
								πF.SetLineno(139)
								if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Add(πF, µline, πg.NewStr("\n").ToObject()); πE != nil {
									continue
								}
								πTemp005 = πg.NewTuple2(µttype, πTemp008).ToObject()
								πF.PushCheckpoint(8)
								return πTemp005, nil
							Label8:
								πTemp008 = πSent
								// line 140: lineno += 1
								πF.SetLineno(140)
								if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.IAdd(πF, µlineno, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µlineno = πTemp008
								// line 141: yield (['ln'], self.fmt_str % lineno)
								πF.SetLineno(141)
								πTemp002 = make([]*πg.Object, 1)
								πTemp002[0] = ßln.ToObject()
								πTemp010 = πg.NewList(πTemp002...).ToObject()
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp012, πE = πg.GetAttr(πF, µself, ßfmt_str, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.Mod(πF, πTemp012, µlineno); πE != nil {
									continue
								}
								πTemp008 = πg.NewTuple2(πTemp010, πTemp011).ToObject()
								πF.PushCheckpoint(9)
								return πTemp008, nil
							Label9:
								πTemp010 = πSent
								continue
							Label6:
								if πE != nil || πR != nil {
									continue
								}
							Label7:
								// line 142: yield (ttype, lines[-1])
								πF.SetLineno(142)
								if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp010 = πTemp011
								if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetItem(πF, µlines, πTemp010); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µttype, πTemp011).ToObject()
								πF.PushCheckpoint(10)
								return πTemp004, nil
							Label10:
								πTemp010 = πSent
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("NumberLines").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNumberLines.ToObject(), πTemp006); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("code_analyzer", Code)
}
