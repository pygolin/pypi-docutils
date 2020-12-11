package null

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/null.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßParser := πg.InternStr("Parser")
		ß__doc__ := πg.InternStr("__doc__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßnull := πg.InternStr("null")
		ßparse := πg.InternStr("parse")
		ßparsers := πg.InternStr("parsers")
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
			// line 5: """A do-nothing parser."""
			πF.SetLineno(5)
			// line 5: """A do-nothing parser."""
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A do-nothing parser.").ToObject()); πE != nil {
				continue
			}
			// line 7: from docutils import parsers
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßparsers.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: class Parser(parsers.Parser):
			πF.SetLineno(10)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßparsers); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßParser, nil); πE != nil {
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
			_, πE = πg.NewCode("Parser", "/usr/lib/python2.7/site-packages/docutils/parsers/null.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 12: """A do-nothing parser."""
					πF.SetLineno(12)
					// line 12: """A do-nothing parser."""
					πF.SetLineno(12)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A do-nothing parser.").ToObject()); πE != nil {
						continue
					}
					// line 14: supported = ('null',)
					πF.SetLineno(14)
					πTemp001 = πg.NewTuple1(ßnull.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 16: config_section = 'null parser'
					πF.SetLineno(16)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("null parser").ToObject()); πE != nil {
						continue
					}
					// line 17: config_section_dependencies = ('parsers',)
					πF.SetLineno(17)
					πTemp001 = πg.NewTuple1(ßparsers.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 19: def parse(self, inputstring, document):
					πF.SetLineno(19)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "inputstring", Def: nil}
					πTemp002[2] = πg.Param{Name: "document", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/site-packages/docutils/parsers/null.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinputstring *πg.Object = πArgs[1]
						_ = µinputstring
						var µdocument *πg.Object = πArgs[2]
						_ = µdocument
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
							// line 20: pass
							πF.SetLineno(20)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßparse.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Parser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßParser.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("null", Code)
}
