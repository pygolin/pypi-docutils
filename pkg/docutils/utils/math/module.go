package math

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/math/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßFalse := πg.InternStr("False")
		ß__doc__ := πg.InternStr("__doc__")
		ßalign := πg.InternStr("align")
		ßequation := πg.InternStr("equation")
		ßfind := πg.InternStr("find")
		ßjoin := πg.InternStr("join")
		ßpick_math_environment := πg.InternStr("pick_math_environment")
		ßsplit := πg.InternStr("split")
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
			// line 12: """
			πF.SetLineno(12)
			// line 12: """
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis is the Docutils (Python Documentation Utilities) \"math\" sub-package.\n\nIt contains various modules for conversion between different math formats\n(LaTeX, MathML, HTML).\n\n:math2html:    LaTeX math -> HTML conversion from eLyXer\n:latex2mathml: LaTeX math -> presentational MathML\n:unichar2tex:  Unicode character to LaTeX math translation table\n:tex2unichar:  LaTeX math to Unicode character translation dictionaries\n:tex2mathml_extern: Wrapper for TeX -> MathML command line converters\n").ToObject()); πE != nil {
				continue
			}
			// line 28: def pick_math_environment(code, numbered=False):
			πF.SetLineno(28)
			πTemp002 = make([]πg.Param, 2)
			πTemp002[0] = πg.Param{Name: "code", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp002[1] = πg.Param{Name: "numbered", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("pick_math_environment", "/usr/lib/python2.7/site-packages/docutils/utils/math/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]
				_ = µcode
				var µnumbered *πg.Object = πArgs[1]
				_ = µnumbered
				var µchunks *πg.Object = πg.UnboundLocal
				_ = µchunks
				var µtoplevel_code *πg.Object = πg.UnboundLocal
				_ = µtoplevel_code
				var µenv *πg.Object = πg.UnboundLocal
				_ = µenv
				var πTemp001 []*πg.Object
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
					// line 29: """Return the right math environment to display `code`.
					πF.SetLineno(29)
					// line 39: chunks = code.split(r'\begin{')
					πF.SetLineno(39)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\\begin{").ToObject()
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcode, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µchunks = πTemp003
					// line 40: toplevel_code = ''.join([chunk.split(r'\end{')[-1]
					πF.SetLineno(40)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/math/__init__.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µchunk *πg.Object = πg.UnboundLocal
						_ = µchunk
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
						var πTemp006 []*πg.Object
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
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µchunks); πE != nil {
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
									µchunk = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 40: toplevel_code = ''.join([chunk.split(r'\end{')[-1]
								πF.SetLineno(40)
								if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp004 = πTemp005
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewStr("\\end{").ToObject()
								if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µchunk, ßsplit, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
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
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtoplevel_code = πTemp005
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\\\\").ToObject()
					if πE = πg.CheckLocal(πF, µtoplevel_code, "toplevel_code"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µtoplevel_code, ßfind, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GE(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label1
					}
					goto Label2
					// line 42: if toplevel_code.find(r'\\') >= 0:
					πF.SetLineno(42)
				Label1:
					// line 43: env = 'align'
					πF.SetLineno(43)
					µenv = ßalign.ToObject()
					goto Label3
				Label2:
					// line 45: env = 'equation'
					πF.SetLineno(45)
					µenv = ßequation.ToObject()
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µnumbered, "numbered"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µnumbered); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 46: if not numbered:
					πF.SetLineno(46)
				Label4:
					// line 47: env += '*'
					πF.SetLineno(47)
					if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µenv, πg.NewStr("*").ToObject()); πE != nil {
						continue
					}
					µenv = πTemp002
					goto Label5
				Label5:
					// line 48: return env
					πF.SetLineno(48)
					if πE = πg.CheckLocal(πF, µenv, "env"); πE != nil {
						continue
					}
					πR = µenv
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpick_math_environment.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: """Return the right math environment to display `code`.
			πF.SetLineno(29)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return the right math environment to display `code`.\n\n    The test simply looks for line-breaks (``\\``) outside environments.\n    Multi-line formulae are set with ``align``, one-liners with\n    ``equation``.\n\n    If `numbered` evaluates to ``False``, the \"starred\" versions are used\n    to suppress numbering.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßpick_math_environment); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("math", Code)
}
