package pygmentsformatter

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/pygments"
	_ "github.com/pygolin/stdlib/pkg/pygments/formatter"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßAttribute := πg.InternStr("Attribute")
		ßBacktick := πg.InternStr("Backtick")
		ßBuiltin := πg.InternStr("Builtin")
		ßClass := πg.InternStr("Class")
		ßComment := πg.InternStr("Comment")
		ßFloat := πg.InternStr("Float")
		ßFormatter := πg.InternStr("Formatter")
		ßFunction := πg.InternStr("Function")
		ßHex := πg.InternStr("Hex")
		ßInteger := πg.InternStr("Integer")
		ßKeyword := πg.InternStr("Keyword")
		ßLiteral := πg.InternStr("Literal")
		ßLong := πg.InternStr("Long")
		ßName := πg.InternStr("Name")
		ßNumber := πg.InternStr("Number")
		ßOct := πg.InternStr("Oct")
		ßOdtPygmentsFormatter := πg.InternStr("OdtPygmentsFormatter")
		ßOdtPygmentsLaTeXFormatter := πg.InternStr("OdtPygmentsLaTeXFormatter")
		ßOdtPygmentsProgFormatter := πg.InternStr("OdtPygmentsProgFormatter")
		ßOperator := πg.InternStr("Operator")
		ßString := πg.InternStr("String")
		ßToken := πg.InternStr("Token")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßescape_function := πg.InternStr("escape_function")
		ßformat := πg.InternStr("format")
		ßformatter := πg.InternStr("formatter")
		ßpygments := πg.InternStr("pygments")
		ßrststyle := πg.InternStr("rststyle")
		ßrststyle_function := πg.InternStr("rststyle_function")
		ßtoken := πg.InternStr("token")
		ßwrite := πg.InternStr("write")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n\nAdditional support for Pygments formatter.\n\n").ToObject()); πE != nil {
				continue
			}
			// line 12: import pygments
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "pygments"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßpygments.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: import pygments.formatter
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "pygments.formatter"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßpygments.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: class OdtPygmentsFormatter(pygments.formatter.Formatter):
			πF.SetLineno(16)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßpygments); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßformatter, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßFormatter, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OdtPygmentsFormatter", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 17: def __init__(self, rststyle_function, escape_function):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rststyle_function", Def: nil}
					πTemp002[2] = πg.Param{Name: "escape_function", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrststyle_function *πg.Object = πArgs[1]
						_ = µrststyle_function
						var µescape_function *πg.Object = πArgs[2]
						_ = µescape_function
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
							// line 18: pygments.formatter.Formatter.__init__(self)
							πF.SetLineno(18)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpygments); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßformatter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßFormatter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 19: self.rststyle_function = rststyle_function
							πF.SetLineno(19)
							if πE = πg.CheckLocal(πF, µrststyle_function, "rststyle_function"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µrststyle_function); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrststyle_function, πTemp002); πE != nil {
								continue
							}
							// line 20: self.escape_function = escape_function
							πF.SetLineno(20)
							if πE = πg.CheckLocal(πF, µescape_function, "escape_function"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µescape_function); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßescape_function, πTemp002); πE != nil {
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
					// line 22: def rststyle(self, name, parameters=( )):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002[2] = πg.Param{Name: "parameters", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("rststyle", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var µparameters *πg.Object = πArgs[2]
						_ = µparameters
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
							// line 23: return self.rststyle_function(name, parameters)
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µparameters, "parameters"); πE != nil {
								continue
							}
							πTemp001[1] = µparameters
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle_function, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßrststyle.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("OdtPygmentsFormatter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOdtPygmentsFormatter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 26: class OdtPygmentsProgFormatter(OdtPygmentsFormatter):
			πF.SetLineno(26)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßOdtPygmentsFormatter); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OdtPygmentsProgFormatter", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 27: def format(self, tokensource, outfile):
					πF.SetLineno(27)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tokensource", Def: nil}
					πTemp002[2] = πg.Param{Name: "outfile", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("format", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtokensource *πg.Object = πArgs[1]
						_ = µtokensource
						var µoutfile *πg.Object = πArgs[2]
						_ = µoutfile
						var µtokenclass *πg.Object = πg.UnboundLocal
						_ = µtokenclass
						var µttype *πg.Object = πg.UnboundLocal
						_ = µttype
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µs2 *πg.Object = πg.UnboundLocal
						_ = µs2
						var µs1 *πg.Object = πg.UnboundLocal
						_ = µs1
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 28: tokenclass = pygments.token.Token
							πF.SetLineno(28)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßpygments); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtoken, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßToken, nil); πE != nil {
								continue
							}
							µtokenclass = πTemp001
							if πE = πg.CheckLocal(πF, µtokensource, "tokensource"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtokensource); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp003 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µttype = πTemp005
								µvalue = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 30: value = self.escape_function(value)
							πF.SetLineno(30)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßescape_function, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µvalue = πTemp005
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßKeyword, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßString, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßNumber, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp008, ßInteger, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßNumber, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßInteger, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßLong, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßNumber, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp010, ßFloat, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßNumber, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp011, ßHex, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßNumber, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp012, ßOct, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßNumber, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple6(πTemp006, πTemp009, πTemp008, πTemp010, πTemp011, πTemp013).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, µttype); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßOperator, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßComment, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßName, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßClass, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßName, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßFunction, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßName, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 31: if ttype == tokenclass.Keyword:
							πF.SetLineno(31)
						Label4:
							// line 32: s2 = self.rststyle('codeblock-keyword')
							πF.SetLineno(32)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-keyword").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 33: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 35: elif ttype == tokenclass.Literal.String:
							πF.SetLineno(35)
						Label5:
							// line 36: s2 = self.rststyle('codeblock-string')
							πF.SetLineno(36)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-string").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 37: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(37)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 39: elif ttype in (
							πF.SetLineno(39)
						Label6:
							// line 47: s2 = self.rststyle('codeblock-number')
							πF.SetLineno(47)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-number").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 48: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 50: elif ttype == tokenclass.Operator:
							πF.SetLineno(50)
						Label7:
							// line 51: s2 = self.rststyle('codeblock-operator')
							πF.SetLineno(51)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-operator").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 52: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 54: elif ttype == tokenclass.Comment:
							πF.SetLineno(54)
						Label8:
							// line 55: s2 = self.rststyle('codeblock-comment')
							πF.SetLineno(55)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-comment").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 56: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(56)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 58: elif ttype == tokenclass.Name.Class:
							πF.SetLineno(58)
						Label9:
							// line 59: s2 = self.rststyle('codeblock-classname')
							πF.SetLineno(59)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-classname").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 60: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 62: elif ttype == tokenclass.Name.Function:
							πF.SetLineno(62)
						Label10:
							// line 63: s2 = self.rststyle('codeblock-functionname')
							πF.SetLineno(63)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-functionname").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 64: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(64)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
							// line 66: elif ttype == tokenclass.Name:
							πF.SetLineno(66)
						Label11:
							// line 67: s2 = self.rststyle('codeblock-name')
							πF.SetLineno(67)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-name").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 68: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
						Label12:
							// line 71: s1 = value
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							µs1 = µvalue
							goto Label13
						Label13:
							// line 72: outfile.write(s1)
							πF.SetLineno(72)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp007[0] = µs1
							if πE = πg.CheckLocal(πF, µoutfile, "outfile"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoutfile, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßformat.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("OdtPygmentsProgFormatter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOdtPygmentsProgFormatter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 75: class OdtPygmentsLaTeXFormatter(OdtPygmentsFormatter):
			πF.SetLineno(75)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßOdtPygmentsFormatter); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OdtPygmentsLaTeXFormatter", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 76: def format(self, tokensource, outfile):
					πF.SetLineno(76)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tokensource", Def: nil}
					πTemp002[2] = πg.Param{Name: "outfile", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("format", "/usr/lib/python2.7/site-packages/docutils/writers/odf_odt/pygmentsformatter.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtokensource *πg.Object = πArgs[1]
						_ = µtokensource
						var µoutfile *πg.Object = πArgs[2]
						_ = µoutfile
						var µtokenclass *πg.Object = πg.UnboundLocal
						_ = µtokenclass
						var µttype *πg.Object = πg.UnboundLocal
						_ = µttype
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µs2 *πg.Object = πg.UnboundLocal
						_ = µs2
						var µs1 *πg.Object = πg.UnboundLocal
						_ = µs1
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 77: tokenclass = pygments.token.Token
							πF.SetLineno(77)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßpygments); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtoken, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßToken, nil); πE != nil {
								continue
							}
							µtokenclass = πTemp001
							if πE = πg.CheckLocal(πF, µtokensource, "tokensource"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtokensource); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp003 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µttype = πTemp005
								µvalue = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 79: value = self.escape_function(value)
							πF.SetLineno(79)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßescape_function, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µvalue = πTemp005
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßKeyword, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßString, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µtokenclass, ßLiteral, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßString, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp009, ßBacktick, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp008, πTemp006).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, µttype); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßName, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßAttribute, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßComment, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtokenclass, "tokenclass"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µtokenclass, ßName, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßBuiltin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µttype, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 80: if ttype == tokenclass.Keyword:
							πF.SetLineno(80)
						Label4:
							// line 81: s2 = self.rststyle('codeblock-keyword')
							πF.SetLineno(81)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-keyword").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 82: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(82)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label10
							// line 84: elif ttype in (tokenclass.Literal.String,
							πF.SetLineno(84)
						Label5:
							// line 87: s2 = self.rststyle('codeblock-string')
							πF.SetLineno(87)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-string").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 88: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label10
							// line 90: elif ttype == tokenclass.Name.Attribute:
							πF.SetLineno(90)
						Label6:
							// line 91: s2 = self.rststyle('codeblock-operator')
							πF.SetLineno(91)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-operator").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 92: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label10
							// line 94: elif ttype == tokenclass.Comment:
							πF.SetLineno(94)
						Label7:
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µvalue, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 95: if value[-1] == '\n':
							πF.SetLineno(95)
						Label11:
							// line 96: s2 = self.rststyle('codeblock-comment')
							πF.SetLineno(96)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-comment").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 97: s1 = '<text:span text:style-name="%s">%s</text:span>\n' % \
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µvalue, πTemp006); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, πTemp008).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
						Label12:
							// line 100: s2 = self.rststyle('codeblock-comment')
							πF.SetLineno(100)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-comment").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 101: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label13
						Label13:
							goto Label10
							// line 103: elif ttype == tokenclass.Name.Builtin:
							πF.SetLineno(103)
						Label8:
							// line 104: s2 = self.rststyle('codeblock-name')
							πF.SetLineno(104)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("codeblock-name").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrststyle, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µs2 = πTemp005
							// line 105: s1 = '<text:span text:style-name="%s">%s</text:span>' % \
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µs2, µvalue).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<text:span text:style-name=\"%s\">%s</text:span>").ToObject(), πTemp005); πE != nil {
								continue
							}
							µs1 = πTemp002
							goto Label10
						Label9:
							// line 108: s1 = value
							πF.SetLineno(108)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							µs1 = µvalue
							goto Label10
						Label10:
							// line 109: outfile.write(s1)
							πF.SetLineno(109)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp007[0] = µs1
							if πE = πg.CheckLocal(πF, µoutfile, "outfile"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoutfile, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßformat.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("OdtPygmentsLaTeXFormatter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOdtPygmentsLaTeXFormatter.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("pygmentsformatter", Code)
}
