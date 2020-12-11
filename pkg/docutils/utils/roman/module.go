package roman

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßC := πg.InternStr("C")
		ßCD := πg.InternStr("CD")
		ßCM := πg.InternStr("CM")
		ßD := πg.InternStr("D")
		ßException := πg.InternStr("Exception")
		ßI := πg.InternStr("I")
		ßIV := πg.InternStr("IV")
		ßIX := πg.InternStr("IX")
		ßInvalidRomanNumeralError := πg.InternStr("InvalidRomanNumeralError")
		ßL := πg.InternStr("L")
		ßM := πg.InternStr("M")
		ßNotIntegerError := πg.InternStr("NotIntegerError")
		ßOutOfRangeError := πg.InternStr("OutOfRangeError")
		ßRomanError := πg.InternStr("RomanError")
		ßV := πg.InternStr("V")
		ßVERBOSE := πg.InternStr("VERBOSE")
		ßX := πg.InternStr("X")
		ßXC := πg.InternStr("XC")
		ßXL := πg.InternStr("XL")
		ß__author__ := πg.InternStr("__author__")
		ß__copyright__ := πg.InternStr("__copyright__")
		ß__date__ := πg.InternStr("__date__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ßcompile := πg.InternStr("compile")
		ßfromRoman := πg.InternStr("fromRoman")
		ßint := πg.InternStr("int")
		ßlen := πg.InternStr("len")
		ßre := πg.InternStr("re")
		ßromanNumeralMap := πg.InternStr("romanNumeralMap")
		ßromanNumeralPattern := πg.InternStr("romanNumeralPattern")
		ßsearch := πg.InternStr("search")
		ßtoRoman := πg.InternStr("toRoman")
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
		var πTemp006 []πg.Param
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 1: """Convert to and from Roman numerals"""
			πF.SetLineno(1)
			// line 1: """Convert to and from Roman numerals"""
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Convert to and from Roman numerals").ToObject()); πE != nil {
				continue
			}
			// line 3: __author__ = "Mark Pilgrim (f8dy@diveintopython.org)"
			πF.SetLineno(3)
			if πE = πF.Globals().SetItem(πF, ß__author__.ToObject(), πg.NewStr("Mark Pilgrim (f8dy@diveintopython.org)").ToObject()); πE != nil {
				continue
			}
			// line 4: __version__ = "1.4"
			πF.SetLineno(4)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("1.4").ToObject()); πE != nil {
				continue
			}
			// line 5: __date__ = "8 August 2001"
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__date__.ToObject(), πg.NewStr("8 August 2001").ToObject()); πE != nil {
				continue
			}
			// line 6: __copyright__ = """Copyright (c) 2001 Mark Pilgrim
			πF.SetLineno(6)
			if πE = πF.Globals().SetItem(πF, ß__copyright__.ToObject(), πg.NewStr("Copyright (c) 2001 Mark Pilgrim\n\nThis program is part of \"Dive Into Python\", a free Python tutorial for\nexperienced programmers.  Visit http://diveintopython.org/ for the\nlatest version.\n\nThis program is free software; you can redistribute it and/or modify\nit under the terms of the Python 2.1.1 license, available at\nhttp://www.python.org/2.1.1/license.html\n").ToObject()); πE != nil {
				continue
			}
			// line 17: import re
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: class RomanError(Exception): pass
			πF.SetLineno(20)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("RomanError", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 20: class RomanError(Exception): pass
					πF.SetLineno(20)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("RomanError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRomanError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 21: class OutOfRangeError(RomanError): pass
			πF.SetLineno(21)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßRomanError); πE != nil {
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
			_, πE = πg.NewCode("OutOfRangeError", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 21: class OutOfRangeError(RomanError): pass
					πF.SetLineno(21)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("OutOfRangeError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOutOfRangeError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 22: class NotIntegerError(RomanError): pass
			πF.SetLineno(22)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßRomanError); πE != nil {
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
			_, πE = πg.NewCode("NotIntegerError", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 22: class NotIntegerError(RomanError): pass
					πF.SetLineno(22)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("NotIntegerError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNotIntegerError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 23: class InvalidRomanNumeralError(RomanError): pass
			πF.SetLineno(23)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßRomanError); πE != nil {
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
			_, πE = πg.NewCode("InvalidRomanNumeralError", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 23: class InvalidRomanNumeralError(RomanError): pass
					πF.SetLineno(23)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("InvalidRomanNumeralError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInvalidRomanNumeralError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 26: romanNumeralMap = (('M',  1000),
			πF.SetLineno(26)
			πTemp002 = make([]*πg.Object, 13)
			πTemp004 = πg.NewTuple2(ßM.ToObject(), πg.NewInt(1000).ToObject()).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewTuple2(ßCM.ToObject(), πg.NewInt(900).ToObject()).ToObject()
			πTemp002[1] = πTemp004
			πTemp004 = πg.NewTuple2(ßD.ToObject(), πg.NewInt(500).ToObject()).ToObject()
			πTemp002[2] = πTemp004
			πTemp004 = πg.NewTuple2(ßCD.ToObject(), πg.NewInt(400).ToObject()).ToObject()
			πTemp002[3] = πTemp004
			πTemp004 = πg.NewTuple2(ßC.ToObject(), πg.NewInt(100).ToObject()).ToObject()
			πTemp002[4] = πTemp004
			πTemp004 = πg.NewTuple2(ßXC.ToObject(), πg.NewInt(90).ToObject()).ToObject()
			πTemp002[5] = πTemp004
			πTemp004 = πg.NewTuple2(ßL.ToObject(), πg.NewInt(50).ToObject()).ToObject()
			πTemp002[6] = πTemp004
			πTemp004 = πg.NewTuple2(ßXL.ToObject(), πg.NewInt(40).ToObject()).ToObject()
			πTemp002[7] = πTemp004
			πTemp004 = πg.NewTuple2(ßX.ToObject(), πg.NewInt(10).ToObject()).ToObject()
			πTemp002[8] = πTemp004
			πTemp004 = πg.NewTuple2(ßIX.ToObject(), πg.NewInt(9).ToObject()).ToObject()
			πTemp002[9] = πTemp004
			πTemp004 = πg.NewTuple2(ßV.ToObject(), πg.NewInt(5).ToObject()).ToObject()
			πTemp002[10] = πTemp004
			πTemp004 = πg.NewTuple2(ßIV.ToObject(), πg.NewInt(4).ToObject()).ToObject()
			πTemp002[11] = πTemp004
			πTemp004 = πg.NewTuple2(ßI.ToObject(), πg.NewInt(1).ToObject()).ToObject()
			πTemp002[12] = πTemp004
			πTemp001 = πg.NewTuple(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßromanNumeralMap.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: def toRoman(n):
			πF.SetLineno(40)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "n", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("toRoman", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]
				_ = µn
				var µresult *πg.Object = πg.UnboundLocal
				_ = µresult
				var µnumeral *πg.Object = πg.UnboundLocal
				_ = µnumeral
				var µinteger *πg.Object = πg.UnboundLocal
				_ = µinteger
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 9:
						goto Label9
					case 10:
						goto Label10
					case 6:
						goto Label6
					case 7:
						goto Label7
					default:
						panic("unexpected function state")
					}
					// line 41: """convert integer to Roman numeral"""
					πF.SetLineno(41)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, πg.NewInt(0).ToObject(), µn); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label1
					}
					if πTemp002, πE = πg.LT(πF, µn, πg.NewInt(5000).ToObject()); πE != nil {
						continue
					}
				Label1:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label2
					}
					goto Label3
					// line 42: if not (0 < n < 5000):
					πF.SetLineno(42)
				Label2:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("number out of range (must be 1..4999)").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOutOfRangeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 43: raise OutOfRangeError("number out of range (must be 1..4999)")
					πF.SetLineno(43)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label3
				Label3:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πTemp004[0] = µn
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp005, µn); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 44: if int(n) != n:
					πF.SetLineno(44)
				Label4:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("decimals can not be converted").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNotIntegerError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 45: raise NotIntegerError("decimals can not be converted")
					πF.SetLineno(45)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label5
				Label5:
					// line 47: result = ""
					πF.SetLineno(47)
					µresult = ß.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßromanNumeralMap); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
						continue
					}
					πF.PushCheckpoint(7)
					πTemp003 = false
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label8
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
							continue
						}
						µnumeral = πTemp005
						µinteger = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(6)
					// line 49: while n >= integer:
					πF.SetLineno(49)
					πF.PushCheckpoint(10)
					πTemp006 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µn, µinteger); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(9)
					// line 50: result += numeral
					πF.SetLineno(50)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnumeral, "numeral"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µresult, µnumeral); πE != nil {
						continue
					}
					µresult = πTemp002
					// line 51: n -= integer
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ISub(πF, µn, µinteger); πE != nil {
						continue
					}
					µn = πTemp002
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 52: return result
					πF.SetLineno(52)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtoRoman.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 41: """convert integer to Roman numeral"""
			πF.SetLineno(41)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("convert integer to Roman numeral").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßtoRoman); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 55: romanNumeralPattern = re.compile("""
			πF.SetLineno(55)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("\n    ^                   # beginning of string\n    M{0,4}              # thousands - 0 to 4 M's\n    (CM|CD|D?C{0,3})    # hundreds - 900 (CM), 400 (CD), 0-300 (0 to 3 C's),\n                        #            or 500-800 (D, followed by 0 to 3 C's)\n    (XC|XL|L?X{0,3})    # tens - 90 (XC), 40 (XL), 0-30 (0 to 3 X's),\n                        #        or 50-80 (L, followed by 0 to 3 X's)\n    (IX|IV|V?I{0,3})    # ones - 9 (IX), 4 (IV), 0-3 (0 to 3 I's),\n                        #        or 5-8 (V, followed by 0 to 3 I's)\n    $                   # end of string\n    ").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßVERBOSE, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp005
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßromanNumeralPattern.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 67: def fromRoman(s):
			πF.SetLineno(67)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "s", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("fromRoman", "/usr/lib/python2.7/site-packages/docutils/utils/roman.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]
				_ = µs
				var µresult *πg.Object = πg.UnboundLocal
				_ = µresult
				var µindex *πg.Object = πg.UnboundLocal
				_ = µindex
				var µnumeral *πg.Object = πg.UnboundLocal
				_ = µnumeral
				var µinteger *πg.Object = πg.UnboundLocal
				_ = µinteger
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8:
						goto Label8
					case 9:
						goto Label9
					case 5:
						goto Label5
					case 6:
						goto Label6
					default:
						panic("unexpected function state")
					}
					// line 68: """convert Roman numeral to integer"""
					πF.SetLineno(68)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µs); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 69: if not s:
					πF.SetLineno(69)
				Label1:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("Input can not be blank").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßInvalidRomanNumeralError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 70: raise InvalidRomanNumeralError('Input can not be blank')
					πF.SetLineno(70)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label2
				Label2:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πTemp004, πE = πg.ResolveGlobal(πF, ßromanNumeralPattern); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsearch, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 72: if not romanNumeralPattern.search(s):
					πF.SetLineno(72)
				Label3:
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid Roman numeral: %s").ToObject(), µs); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßInvalidRomanNumeralError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 73: raise InvalidRomanNumeralError('Invalid Roman numeral: %s' % s)
					πF.SetLineno(73)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label4
				Label4:
					// line 75: result = 0
					πF.SetLineno(75)
					µresult = πg.NewInt(0).ToObject()
					// line 76: index = 0
					πF.SetLineno(76)
					µindex = πg.NewInt(0).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßromanNumeralMap); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label7
					}
					if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
							continue
						}
						µnumeral = πTemp005
						µinteger = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)
					// line 78: while s[index:index+len(numeral)] == numeral:
					πF.SetLineno(78)
					πF.PushCheckpoint(9)
					πTemp006 = false
				Label8:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnumeral, "numeral"); πE != nil {
						continue
					}
					πTemp003[0] = µnumeral
					if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp007, πE = πg.Add(πF, µindex, πTemp010); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{µindex, πTemp007, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µs, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnumeral, "numeral"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp007, µnumeral); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(8)
					// line 79: result += integer
					πF.SetLineno(79)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinteger, "integer"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µresult, µinteger); πE != nil {
						continue
					}
					µresult = πTemp004
					// line 80: index += len(numeral)
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnumeral, "numeral"); πE != nil {
						continue
					}
					πTemp003[0] = µnumeral
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.IAdd(πF, µindex, πTemp005); πE != nil {
						continue
					}
					µindex = πTemp004
					continue
				Label9:
					if πE != nil || πR != nil {
						continue
					}
				Label10:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 81: return result
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πR = µresult
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfromRoman.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 68: """convert Roman numeral to integer"""
			πF.SetLineno(68)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("convert Roman numeral to integer").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßfromRoman); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("roman", Code)
}
