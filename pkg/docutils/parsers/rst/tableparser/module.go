package tableparser

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAssertionError := πg.InternStr("AssertionError")
		ßDataError := πg.InternStr("DataError")
		ßFalse := πg.InternStr("False")
		ßGridTableParser := πg.InternStr("GridTableParser")
		ßIndexError := πg.InternStr("IndexError")
		ßNone := πg.InternStr("None")
		ßSimpleTableParser := πg.InternStr("SimpleTableParser")
		ßTableMarkupError := πg.InternStr("TableMarkupError")
		ßTableParser := πg.InternStr("TableParser")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßblock := πg.InternStr("block")
		ßborder_end := πg.InternStr("border_end")
		ßbottom := πg.InternStr("bottom")
		ßcells := πg.InternStr("cells")
		ßcheck_columns := πg.InternStr("check_columns")
		ßcheck_parse_complete := πg.InternStr("check_parse_complete")
		ßcolseps := πg.InternStr("colseps")
		ßcolumns := πg.InternStr("columns")
		ßcompile := πg.InternStr("compile")
		ßdisconnect := πg.InternStr("disconnect")
		ßdone := πg.InternStr("done")
		ßdouble_width_pad_char := πg.InternStr("double_width_pad_char")
		ßextend := πg.InternStr("extend")
		ßfind := πg.InternStr("find")
		ßfind_head_body_sep := πg.InternStr("find_head_body_sep")
		ßget_2D_block := πg.InternStr("get_2D_block")
		ßhead_body_sep := πg.InternStr("head_body_sep")
		ßhead_body_separator_pat := πg.InternStr("head_body_separator_pat")
		ßinit_row := πg.InternStr("init_row")
		ßitems := πg.InternStr("items")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßmark_done := πg.InternStr("mark_done")
		ßmatch := πg.InternStr("match")
		ßmax := πg.InternStr("max")
		ßmaxsize := πg.InternStr("maxsize")
		ßobject := πg.InternStr("object")
		ßoffset := πg.InternStr("offset")
		ßparse := πg.InternStr("parse")
		ßparse_columns := πg.InternStr("parse_columns")
		ßparse_row := πg.InternStr("parse_row")
		ßparse_table := πg.InternStr("parse_table")
		ßpop := πg.InternStr("pop")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreplace := πg.InternStr("replace")
		ßright := πg.InternStr("right")
		ßrowseps := πg.InternStr("rowseps")
		ßrstrip := πg.InternStr("rstrip")
		ßscan_cell := πg.InternStr("scan_cell")
		ßscan_down := πg.InternStr("scan_down")
		ßscan_left := πg.InternStr("scan_left")
		ßscan_right := πg.InternStr("scan_right")
		ßscan_up := πg.InternStr("scan_up")
		ßsetdefault := πg.InternStr("setdefault")
		ßsetup := πg.InternStr("setup")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßspan_pat := πg.InternStr("span_pat")
		ßstrip := πg.InternStr("strip")
		ßstrip_combining_chars := πg.InternStr("strip_combining_chars")
		ßstructure_from_cells := πg.InternStr("structure_from_cells")
		ßsys := πg.InternStr("sys")
		ßtable := πg.InternStr("table")
		ßupdate_dict_of_lists := πg.InternStr("update_dict_of_lists")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 []πg.Param
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis module defines table parser classes,which parse plaintext-graphic tables\nand produce a well-formed data structure suitable for building a CALS table.\n\n:Classes:\n    - `GridTableParser`: Parse fully-formed tables represented with a grid.\n    - `SimpleTableParser`: Parse simple tables, delimited by top & bottom\n      borders.\n\n:Exception class: `TableMarkupError`\n\n:Function:\n    `update_dict_of_lists()`: Merge two dictionaries containing list values.\n").ToObject()); πE != nil {
				continue
			}
			// line 20: __docformat__ = 'reStructuredText'
			πF.SetLineno(20)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 23: import re
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 24: import sys
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: from docutils import DataError
			πF.SetLineno(25)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDataError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDataError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 26: from docutils.utils import strip_combining_chars
			πF.SetLineno(26)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßstrip_combining_chars); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstrip_combining_chars.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 29: class TableMarkupError(DataError):
			πF.SetLineno(29)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TableMarkupError", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 31: """
					πF.SetLineno(31)
					// line 31: """
					πF.SetLineno(31)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise if there is any problem with table markup.\n\n    The keyword argument `offset` denotes the offset of the problem\n    from the table's start line.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 38: def __init__(self, *args, **kwargs):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 39: self.offset = kwargs.pop('offset', 0)
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßoffset.ToObject()
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µkwargs, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoffset, πTemp002); πE != nil {
								continue
							}
							// line 40: DataError.__init__(self, *args)
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µargs, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TableMarkupError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTableMarkupError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 43: class TableParser(object):
			πF.SetLineno(43)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TableParser", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 45: """
					πF.SetLineno(45)
					// line 45: """
					πF.SetLineno(45)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract superclass for the common parts of the syntax-specific parsers.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 49: head_body_separator_pat = None
					πF.SetLineno(49)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhead_body_separator_pat.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 50: """Matches the row separator between head rows and body rows."""
					πF.SetLineno(50)
					// line 52: double_width_pad_char = '\x00'
					πF.SetLineno(52)
					if πE = πClass.SetItem(πF, ßdouble_width_pad_char.ToObject(), πg.NewStr("\x00").ToObject()); πE != nil {
						continue
					}
					// line 53: """Padding character for East Asian double-width text."""
					πF.SetLineno(53)
					// line 55: def parse(self, block):
					πF.SetLineno(55)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "block", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µblock *πg.Object = πArgs[1]
						_ = µblock
						var µstructure *πg.Object = πg.UnboundLocal
						_ = µstructure
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
							// line 56: """
							πF.SetLineno(56)
							// line 65: self.setup(block)
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp001[0] = µblock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 66: self.find_head_body_sep()
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfind_head_body_sep, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 67: self.parse_table()
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_table, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 68: structure = self.structure_from_cells()
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstructure_from_cells, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstructure = πTemp003
							// line 69: return structure
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µstructure, "structure"); πE != nil {
								continue
							}
							πR = µstructure
							continue
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
					// line 56: """
					πF.SetLineno(56)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Analyze the text `block` and return a table data structure.\n\n        Given a plaintext-graphic table in `block` (list of lines of text; no\n        whitespace padding), parse the table, construct and return the data\n        necessary to construct a CALS table or equivalent.\n\n        Raise `TableMarkupError` if there is any problem with the markup.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßparse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 71: def find_head_body_sep(self):
					πF.SetLineno(71)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("find_head_body_sep", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 72: """Look for a head/body row separator line; store the line index."""
							πF.SetLineno(72)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 74: line = self.block[i]
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							µline = πTemp005
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp002[0] = µline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhead_body_separator_pat, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 75: if self.head_body_separator_pat.match(line):
							πF.SetLineno(75)
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							goto Label7
							// line 76: if self.head_body_sep:
							πF.SetLineno(76)
						Label6:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp008, πTemp009).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("Multiple head/body row separators (table lines %s and %s); only one allowed.").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"offset", µi},
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTableMarkupError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 77: raise TableMarkupError(
							πF.SetLineno(77)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label8
						Label7:
							// line 82: self.head_body_sep = i
							πF.SetLineno(82)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhead_body_sep, πTemp004); πE != nil {
								continue
							}
							// line 83: self.block[i] = line.replace('=', '-')
							πF.SetLineno(83)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("=").ToObject()
							πTemp002[1] = πg.NewStr("-").ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µline, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp004); πE != nil {
								continue
							}
							goto Label8
						Label8:
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp008, πE = πg.Sub(πF, πTemp011, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label9:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label10
							}
							goto Label11
							// line 84: if self.head_body_sep == 0 or self.head_body_sep == (len(self.block)
							πF.SetLineno(84)
						Label10:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("The head/body row separator may not be the first or last line of the table.").ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"offset", µi},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTableMarkupError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 86: raise TableMarkupError('The head/body row separator may not be '
							πF.SetLineno(86)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label11
						Label11:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfind_head_body_sep.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 72: """Look for a head/body row separator line; store the line index."""
					πF.SetLineno(72)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Look for a head/body row separator line; store the line index.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßfind_head_body_sep); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TableParser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTableParser.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 91: class GridTableParser(TableParser):
			πF.SetLineno(91)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTableParser); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GridTableParser", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
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
				var πTemp007 *πg.Object
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
				var πTemp014 *πg.Object
				_ = πTemp014
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 93: """
					πF.SetLineno(93)
					// line 93: """
					πF.SetLineno(93)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Parse a grid table using `parse()`.\n\n    Here's an example of a grid table::\n\n        +------------------------+------------+----------+----------+\n        | Header row, column 1   | Header 2   | Header 3 | Header 4 |\n        +========================+============+==========+==========+\n        | body row 1, column 1   | column 2   | column 3 | column 4 |\n        +------------------------+------------+----------+----------+\n        | body row 2             | Cells may span columns.          |\n        +------------------------+------------+---------------------+\n        | body row 3             | Cells may  | - Table cells       |\n        +------------------------+ span rows. | - contain           |\n        | body row 4             |            | - body elements.    |\n        +------------------------+------------+---------------------+\n\n    Intersections use '+', row separators use '-' (except for one optional\n    head/body row separator, which uses '='), and column separators use '|'.\n\n    Passing the above table to the `parse()` method will result in the\n    following data structure::\n\n        ([24, 12, 10, 10],\n         [[(0, 0, 1, ['Header row, column 1']),\n           (0, 0, 1, ['Header 2']),\n           (0, 0, 1, ['Header 3']),\n           (0, 0, 1, ['Header 4'])]],\n         [[(0, 0, 3, ['body row 1, column 1']),\n           (0, 0, 3, ['column 2']),\n           (0, 0, 3, ['column 3']),\n           (0, 0, 3, ['column 4'])],\n          [(0, 0, 5, ['body row 2']),\n           (0, 2, 5, ['Cells may span columns.']),\n           None,\n           None],\n          [(0, 0, 7, ['body row 3']),\n           (1, 0, 7, ['Cells may', 'span rows.', '']),\n           (1, 1, 7, ['- Table cells', '- contain', '- body elements.']),\n           None],\n          [(0, 0, 9, ['body row 4']), None, None, None]])\n\n    The first item is a list containing column widths (colspecs). The second\n    item is a list of head rows, and the third is a list of body rows. Each\n    row contains a list of cells. Each cell is either None (for a cell unused\n    because of another cell's span), or a tuple. A cell tuple contains four\n    items: the number of extra rows used by the cell in a vertical span\n    (morerows); the number of extra columns used by the cell in a horizontal\n    span (morecols); the line offset of the first line of the cell contents;\n    and the cell contents, a list of lines of text.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 145: head_body_separator_pat = re.compile(r'\+=[=+]+=\+ *$')
					πF.SetLineno(145)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("\\+=[=+]+=\\+ *$").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßhead_body_separator_pat.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 147: def setup(self, block):
					πF.SetLineno(147)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "block", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("setup", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µblock *πg.Object = πArgs[1]
						_ = µblock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Dict
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
							// line 148: self.block = block[:]           # make a copy; it may be modified
							πF.SetLineno(148)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßblock, πTemp001); πE != nil {
								continue
							}
							// line 149: self.block.disconnect()         # don't propagate changes to parent
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdisconnect, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 150: self.bottom = len(block) - 1
							πF.SetLineno(150)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp003[0] = µblock
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbottom, πTemp002); πE != nil {
								continue
							}
							// line 151: self.right = len(block[0]) - 1
							πF.SetLineno(151)
							πTemp003 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µblock, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßright, πTemp002); πE != nil {
								continue
							}
							// line 152: self.head_body_sep = None
							πF.SetLineno(152)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhead_body_sep, πTemp002); πE != nil {
								continue
							}
							// line 153: self.done = [-1] * len(block[0])
							πF.SetLineno(153)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µblock, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdone, πTemp002); πE != nil {
								continue
							}
							// line 154: self.cells = []
							πF.SetLineno(154)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcells, πTemp002); πE != nil {
								continue
							}
							// line 155: self.rowseps = {0: [0]}
							πF.SetLineno(155)
							πTemp006 = πg.NewDict()
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πTemp006.SetItem(πF, πg.NewInt(0).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp006.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrowseps, πTemp002); πE != nil {
								continue
							}
							// line 156: self.colseps = {0: [0]}
							πF.SetLineno(156)
							πTemp006 = πg.NewDict()
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πTemp006.SetItem(πF, πg.NewInt(0).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp006.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcolseps, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetup.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 158: def parse_table(self):
					πF.SetLineno(158)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("parse_table", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcorners *πg.Object = πg.UnboundLocal
						_ = µcorners
						var µtop *πg.Object = πg.UnboundLocal
						_ = µtop
						var µleft *πg.Object = πg.UnboundLocal
						_ = µleft
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
						var µbottom *πg.Object = πg.UnboundLocal
						_ = µbottom
						var µright *πg.Object = πg.UnboundLocal
						_ = µright
						var µrowseps *πg.Object = πg.UnboundLocal
						_ = µrowseps
						var µcolseps *πg.Object = πg.UnboundLocal
						_ = µcolseps
						var µcellblock *πg.Object = πg.UnboundLocal
						_ = µcellblock
						var πTemp001 []*πg.Object
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
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
							// line 159: """
							πF.SetLineno(159)
							// line 170: corners = [(0, 0)]
							πF.SetLineno(170)
							πTemp001 = make([]*πg.Object, 1)
							πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µcorners = πTemp002
							// line 171: while corners:
							πF.SetLineno(171)
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
							if πE = πg.CheckLocal(πF, µcorners, "corners"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µcorners); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 172: top, left = corners.pop(0)
							πF.SetLineno(172)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcorners, "corners"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcorners, ßpop, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp005); πE != nil {
								continue
							}
							µtop = πTemp002
							µleft = πTemp006
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßbottom, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µtop, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µleft, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp006 = µleft
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdone, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LE(πF, µtop, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πTemp005
						Label4:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 173: if top == self.bottom or left == self.right \
							πF.SetLineno(173)
						Label5:
							// line 175: continue
							πF.SetLineno(175)
							continue
							goto Label6
						Label6:
							// line 176: result = self.scan_cell(top, left)
							πF.SetLineno(176)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp001[0] = µtop
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp001[1] = µleft
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßscan_cell, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp005
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µresult); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 177: if not result:
							πF.SetLineno(177)
						Label7:
							// line 178: continue
							πF.SetLineno(178)
							continue
							goto Label8
						Label8:
							// line 179: bottom, right, rowseps, colseps = result
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, µresult); πE != nil {
								continue
							}
							µbottom = πTemp002
							µright = πTemp005
							µrowseps = πTemp006
							µcolseps = πTemp007
							// line 180: update_dict_of_lists(self.rowseps, rowseps)
							πF.SetLineno(180)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrowseps, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							πTemp001[1] = µrowseps
							if πTemp002, πE = πg.ResolveGlobal(πF, ßupdate_dict_of_lists); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 181: update_dict_of_lists(self.colseps, colseps)
							πF.SetLineno(181)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcolseps, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp001[1] = µcolseps
							if πTemp002, πE = πg.ResolveGlobal(πF, ßupdate_dict_of_lists); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 182: self.mark_done(top, left, bottom, right)
							πF.SetLineno(182)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp001[0] = µtop
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp001[1] = µleft
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							πTemp001[2] = µbottom
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp001[3] = µright
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmark_done, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 183: cellblock = self.block.get_2D_block(top + 1, left + 1,
							πF.SetLineno(183)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µtop, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µleft, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							πTemp001[2] = µbottom
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp001[3] = µright
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßget_2D_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcellblock = πTemp002
							// line 185: cellblock.disconnect()      # lines in cell can't sync with parent
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µcellblock, "cellblock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcellblock, ßdisconnect, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 186: cellblock.replace(self.double_width_pad_char, '')
							πF.SetLineno(186)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdouble_width_pad_char, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µcellblock, "cellblock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcellblock, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 187: self.cells.append((top, left, bottom, right, cellblock))
							πF.SetLineno(187)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcellblock, "cellblock"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(µtop, µleft, µbottom, µright, µcellblock).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcells, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 188: corners.extend([(top, right), (bottom, left)])
							πF.SetLineno(188)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µtop, µright).ToObject()
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µbottom, µleft).ToObject()
							πTemp009[1] = πTemp002
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µcorners, "corners"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcorners, ßextend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 189: corners.sort()
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µcorners, "corners"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcorners, ßsort, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcheck_parse_complete, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 190: if not self.check_parse_complete():
							πF.SetLineno(190)
						Label9:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Malformed table; parse incomplete.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTableMarkupError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 191: raise TableMarkupError('Malformed table; parse incomplete.')
							πF.SetLineno(191)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßparse_table.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 159: """
					πF.SetLineno(159)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Start with a queue of upper-left corners, containing the upper-left\n        corner of the table itself. Trace out one rectangular cell, remember\n        it, and add its upper-right and lower-left corners to the queue of\n        potential upper-left corners of further cells. Process the queue in\n        top-to-bottom order, keeping track of how much of each text column has\n        been seen.\n\n        We'll end up knowing all the row and column boundaries, cell positions\n        and their dimensions.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßparse_table); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 193: def mark_done(self, top, left, bottom, right):
					πF.SetLineno(193)
					πTemp004 = make([]πg.Param, 5)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "top", Def: nil}
					πTemp004[2] = πg.Param{Name: "left", Def: nil}
					πTemp004[3] = πg.Param{Name: "bottom", Def: nil}
					πTemp004[4] = πg.Param{Name: "right", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("mark_done", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µbottom *πg.Object = πArgs[3]
						_ = µbottom
						var µright *πg.Object = πArgs[4]
						_ = µright
						var µbefore *πg.Object = πg.UnboundLocal
						_ = µbefore
						var µafter *πg.Object = πg.UnboundLocal
						_ = µafter
						var µcol *πg.Object = πg.UnboundLocal
						_ = µcol
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
						var πTemp008 *πg.Object
						_ = πTemp008
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
							// line 194: """For keeping track of how much of each text column has been seen."""
							πF.SetLineno(194)
							// line 195: before = top - 1
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µtop, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µbefore = πTemp001
							// line 196: after = bottom - 1
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µbottom, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µafter = πTemp001
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp002[0] = µleft
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp002[1] = µright
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
								µcol = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 198: assert self.done[col] == before
							πF.SetLineno(198)
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp004 = µcol
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdone, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbefore, "before"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp007, µbefore); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 199: self.done[col] = after
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µafter, "after"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µafter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdone, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp007 = µcol
							if πE = πg.SetItem(πF, πTemp004, πTemp007, πTemp003); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßmark_done.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 194: """For keeping track of how much of each text column has been seen."""
					πF.SetLineno(194)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("For keeping track of how much of each text column has been seen.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßmark_done); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 201: def check_parse_complete(self):
					πF.SetLineno(201)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("check_parse_complete", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlast *πg.Object = πg.UnboundLocal
						_ = µlast
						var µcol *πg.Object = πg.UnboundLocal
						_ = µcol
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 202: """Each text column should have been completely seen."""
							πF.SetLineno(202)
							// line 203: last = self.bottom - 1
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbottom, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µlast = πTemp001
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
								µcol = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp004 = µcol
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdone, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp007, µlast); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 205: if self.done[col] != last:
							πF.SetLineno(205)
						Label4:
							// line 206: return False
							πF.SetLineno(206)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 207: return True
							πF.SetLineno(207)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheck_parse_complete.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 202: """Each text column should have been completely seen."""
					πF.SetLineno(202)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Each text column should have been completely seen.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_parse_complete); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 209: def scan_cell(self, top, left):
					πF.SetLineno(209)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "top", Def: nil}
					πTemp004[2] = πg.Param{Name: "left", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("scan_cell", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
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
						var πTemp007 []*πg.Object
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
							// line 210: """Starting at the top-left corner, start tracing out a cell."""
							πF.SetLineno(210)
							// line 211: assert self.block[top][left] == '+'
							πF.SetLineno(211)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp002 = µleft
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp004 = µtop
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 212: result = self.scan_right(top, left)
							πF.SetLineno(212)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp007[0] = µtop
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp007[1] = µleft
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßscan_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µresult = πTemp002
							// line 213: return result
							πF.SetLineno(213)
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
					if πE = πClass.SetItem(πF, ßscan_cell.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 210: """Starting at the top-left corner, start tracing out a cell."""
					πF.SetLineno(210)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Starting at the top-left corner, start tracing out a cell.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßscan_cell); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 215: def scan_right(self, top, left):
					πF.SetLineno(215)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "top", Def: nil}
					πTemp004[2] = πg.Param{Name: "left", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("scan_right", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µcolseps *πg.Object = πg.UnboundLocal
						_ = µcolseps
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
						var µbottom *πg.Object = πg.UnboundLocal
						_ = µbottom
						var µrowseps *πg.Object = πg.UnboundLocal
						_ = µrowseps
						var µnewcolseps *πg.Object = πg.UnboundLocal
						_ = µnewcolseps
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
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
							// line 216: """
							πF.SetLineno(216)
							// line 220: colseps = {}
							πF.SetLineno(220)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µcolseps = πTemp002
							// line 221: line = self.block[top]
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp002 = µtop
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µline = πTemp003
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µleft, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp008, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp008, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							goto Label6
							// line 223: if line[i] == '+':
							πF.SetLineno(223)
						Label4:
							// line 224: colseps[i] = [top]
							πF.SetLineno(224)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp005[0] = µtop
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.SetItem(πF, µcolseps, πTemp008, πTemp004); πE != nil {
								continue
							}
							// line 225: result = self.scan_down(top, left, i)
							πF.SetLineno(225)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp005[0] = µtop
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp005[1] = µleft
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005[2] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßscan_down, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µresult = πTemp004
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µresult); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							goto Label8
							// line 226: if result:
							πF.SetLineno(226)
						Label7:
							// line 227: bottom, rowseps, newcolseps = result
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp008}}}, µresult); πE != nil {
								continue
							}
							µbottom = πTemp003
							µrowseps = πTemp004
							µnewcolseps = πTemp008
							// line 228: update_dict_of_lists(colseps, newcolseps)
							πF.SetLineno(228)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp005[0] = µcolseps
							if πE = πg.CheckLocal(πF, µnewcolseps, "newcolseps"); πE != nil {
								continue
							}
							πTemp005[1] = µnewcolseps
							if πTemp003, πE = πg.ResolveGlobal(πF, ßupdate_dict_of_lists); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 229: return bottom, i, rowseps, colseps
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple4(µbottom, µi, µrowseps, µcolseps).ToObject()
							πR = πTemp003
							continue
							goto Label8
						Label8:
							goto Label6
							// line 230: elif line[i] != '-':
							πF.SetLineno(230)
						Label5:
							// line 231: return None
							πF.SetLineno(231)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 232: return None
							πF.SetLineno(232)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßscan_right.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 216: """
					πF.SetLineno(216)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Look for the top-right corner of the cell, and make note of all column\n        boundaries ('+').\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßscan_right); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 234: def scan_down(self, top, left, right):
					πF.SetLineno(234)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "top", Def: nil}
					πTemp004[2] = πg.Param{Name: "left", Def: nil}
					πTemp004[3] = πg.Param{Name: "right", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("scan_down", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µright *πg.Object = πArgs[3]
						_ = µright
						var µrowseps *πg.Object = πg.UnboundLocal
						_ = µrowseps
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
						var µnewrowseps *πg.Object = πg.UnboundLocal
						_ = µnewrowseps
						var µcolseps *πg.Object = πg.UnboundLocal
						_ = µcolseps
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 235: """
							πF.SetLineno(235)
							// line 239: rowseps = {}
							πF.SetLineno(239)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µrowseps = πTemp002
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µtop, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßbottom, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp005 = µright
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp008, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp005 = µright
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.NE(πF, πTemp008, πg.NewStr("|").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							goto Label6
							// line 241: if self.block[i][right] == '+':
							πF.SetLineno(241)
						Label4:
							// line 242: rowseps[i] = [right]
							πF.SetLineno(242)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp003[0] = µright
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.SetItem(πF, µrowseps, πTemp008, πTemp005); πE != nil {
								continue
							}
							// line 243: result = self.scan_left(top, left, i, right)
							πF.SetLineno(243)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp003[0] = µtop
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp003[1] = µleft
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003[2] = µi
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp003[3] = µright
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßscan_left, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µresult = πTemp005
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µresult); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							goto Label8
							// line 244: if result:
							πF.SetLineno(244)
						Label7:
							// line 245: newrowseps, colseps = result
							πF.SetLineno(245)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, µresult); πE != nil {
								continue
							}
							µnewrowseps = πTemp004
							µcolseps = πTemp005
							// line 246: update_dict_of_lists(rowseps, newrowseps)
							πF.SetLineno(246)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							πTemp003[0] = µrowseps
							if πE = πg.CheckLocal(πF, µnewrowseps, "newrowseps"); πE != nil {
								continue
							}
							πTemp003[1] = µnewrowseps
							if πTemp004, πE = πg.ResolveGlobal(πF, ßupdate_dict_of_lists); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 247: return i, rowseps, colseps
							πF.SetLineno(247)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(µi, µrowseps, µcolseps).ToObject()
							πR = πTemp004
							continue
							goto Label8
						Label8:
							goto Label6
							// line 248: elif self.block[i][right] != '|':
							πF.SetLineno(248)
						Label5:
							// line 249: return None
							πF.SetLineno(249)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 250: return None
							πF.SetLineno(250)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßscan_down.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 235: """
					πF.SetLineno(235)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n        Look for the bottom-right corner of the cell, making note of all row\n        boundaries.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßscan_down); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 252: def scan_left(self, top, left, bottom, right):
					πF.SetLineno(252)
					πTemp004 = make([]πg.Param, 5)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "top", Def: nil}
					πTemp004[2] = πg.Param{Name: "left", Def: nil}
					πTemp004[3] = πg.Param{Name: "bottom", Def: nil}
					πTemp004[4] = πg.Param{Name: "right", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("scan_left", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µbottom *πg.Object = πArgs[3]
						_ = µbottom
						var µright *πg.Object = πArgs[4]
						_ = µright
						var µcolseps *πg.Object = πg.UnboundLocal
						_ = µcolseps
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
						var µrowseps *πg.Object = πg.UnboundLocal
						_ = µrowseps
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
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
							// line 253: """
							πF.SetLineno(253)
							// line 257: colseps = {}
							πF.SetLineno(257)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µcolseps = πTemp002
							// line 258: line = self.block[bottom]
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							πTemp002 = µbottom
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µline = πTemp003
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µright, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp005[1] = µleft
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp008, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp008, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							goto Label6
							// line 260: if line[i] == '+':
							πF.SetLineno(260)
						Label4:
							// line 261: colseps[i] = [bottom]
							πF.SetLineno(261)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							πTemp005[0] = µbottom
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.SetItem(πF, µcolseps, πTemp008, πTemp004); πE != nil {
								continue
							}
							goto Label6
							// line 262: elif line[i] != '-':
							πF.SetLineno(262)
						Label5:
							// line 263: return None
							πF.SetLineno(263)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp003 = µleft
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µline, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp004, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 264: if line[left] != '+':
							πF.SetLineno(264)
						Label7:
							// line 265: return None
							πF.SetLineno(265)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label8
						Label8:
							// line 266: result = self.scan_up(top, left, bottom, right)
							πF.SetLineno(266)
							πTemp005 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp005[0] = µtop
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp005[1] = µleft
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							πTemp005[2] = µbottom
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp005[3] = µright
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßscan_up, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µresult = πTemp003
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µresult != πTemp003).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 267: if result is not None:
							πF.SetLineno(267)
						Label9:
							// line 268: rowseps = result
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							µrowseps = µresult
							// line 269: return rowseps, colseps
							πF.SetLineno(269)
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µrowseps, µcolseps).ToObject()
							πR = πTemp002
							continue
							goto Label10
						Label10:
							// line 270: return None
							πF.SetLineno(270)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßscan_left.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 253: """
					πF.SetLineno(253)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n        Noting column boundaries, look for the bottom-left corner of the cell.\n        It must line up with the starting point.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßscan_left); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 272: def scan_up(self, top, left, bottom, right):
					πF.SetLineno(272)
					πTemp004 = make([]πg.Param, 5)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "top", Def: nil}
					πTemp004[2] = πg.Param{Name: "left", Def: nil}
					πTemp004[3] = πg.Param{Name: "bottom", Def: nil}
					πTemp004[4] = πg.Param{Name: "right", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("scan_up", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µbottom *πg.Object = πArgs[3]
						_ = µbottom
						var µright *πg.Object = πArgs[4]
						_ = µright
						var µrowseps *πg.Object = πg.UnboundLocal
						_ = µrowseps
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 273: """
							πF.SetLineno(273)
							// line 276: rowseps = {}
							πF.SetLineno(276)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µrowseps = πTemp002
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µbottom, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp003[1] = µtop
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp005 = µleft
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp008, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp005 = µleft
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.NE(πF, πTemp008, πg.NewStr("|").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							goto Label6
							// line 278: if self.block[i][left] == '+':
							πF.SetLineno(278)
						Label4:
							// line 279: rowseps[i] = [left]
							πF.SetLineno(279)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp003[0] = µleft
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.SetItem(πF, µrowseps, πTemp008, πTemp005); πE != nil {
								continue
							}
							goto Label6
							// line 280: elif self.block[i][left] != '|':
							πF.SetLineno(280)
						Label5:
							// line 281: return None
							πF.SetLineno(281)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 282: return rowseps
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							πR = µrowseps
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßscan_up.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 273: """
					πF.SetLineno(273)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n        Noting row boundaries, see if we can return to the starting point.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßscan_up); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 284: def structure_from_cells(self):
					πF.SetLineno(284)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("structure_from_cells", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrowseps *πg.Object = πg.UnboundLocal
						_ = µrowseps
						var µrowindex *πg.Object = πg.UnboundLocal
						_ = µrowindex
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µcolseps *πg.Object = πg.UnboundLocal
						_ = µcolseps
						var µcolindex *πg.Object = πg.UnboundLocal
						_ = µcolindex
						var µcolspecs *πg.Object = πg.UnboundLocal
						_ = µcolspecs
						var µonerow *πg.Object = πg.UnboundLocal
						_ = µonerow
						var µrows *πg.Object = πg.UnboundLocal
						_ = µrows
						var µremaining *πg.Object = πg.UnboundLocal
						_ = µremaining
						var µtop *πg.Object = πg.UnboundLocal
						_ = µtop
						var µleft *πg.Object = πg.UnboundLocal
						_ = µleft
						var µbottom *πg.Object = πg.UnboundLocal
						_ = µbottom
						var µright *πg.Object = πg.UnboundLocal
						_ = µright
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var µrownum *πg.Object = πg.UnboundLocal
						_ = µrownum
						var µcolnum *πg.Object = πg.UnboundLocal
						_ = µcolnum
						var µmorerows *πg.Object = πg.UnboundLocal
						_ = µmorerows
						var µmorecols *πg.Object = πg.UnboundLocal
						_ = µmorecols
						var µnumheadrows *πg.Object = πg.UnboundLocal
						_ = µnumheadrows
						var µheadrows *πg.Object = πg.UnboundLocal
						_ = µheadrows
						var µbodyrows *πg.Object = πg.UnboundLocal
						_ = µbodyrows
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []πg.Param
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 *πg.Object
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							case 7:
								goto Label7
							case 8:
								goto Label8
							default:
								panic("unexpected function state")
							}
							// line 285: """
							πF.SetLineno(285)
							// line 289: rowseps = sorted(self.rowseps.keys())   # list of row boundaries
							πF.SetLineno(289)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrowseps, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrowseps = πTemp003
							// line 290: rowindex = {}
							πF.SetLineno(290)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							µrowindex = πTemp002
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							πTemp005[0] = µrowseps
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µi = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 292: rowindex[rowseps[i]] = i    # row boundary -> row number mapping
							πF.SetLineno(292)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrowindex, "rowindex"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µrowseps, πTemp009); πE != nil {
								continue
							}
							πTemp006 = πTemp010
							if πE = πg.SetItem(πF, µrowindex, πTemp006, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 293: colseps = sorted(self.colseps.keys())   # list of column boundaries
							πF.SetLineno(293)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcolseps, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcolseps = πTemp003
							// line 294: colindex = {}
							πF.SetLineno(294)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							µcolindex = πTemp002
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp005[0] = µcolseps
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp007 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µi = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 296: colindex[colseps[i]] = i    # column boundary -> col number map
							πF.SetLineno(296)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolindex, "colindex"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µcolseps, πTemp009); πE != nil {
								continue
							}
							πTemp006 = πTemp010
							if πE = πg.SetItem(πF, µcolindex, πTemp006, πTemp003); πE != nil {
								continue
							}
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 297: colspecs = [(colseps[i] - colseps[i - 1] - 1)
							πF.SetLineno(297)
							πTemp011 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
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
								var πTemp006 bool
								_ = πTemp006
								var πTemp007 bool
								_ = πTemp007
								var πTemp008 *πg.Object
								_ = πTemp008
								var πTemp009 *πg.Object
								_ = πTemp009
								var πTemp010 *πg.Object
								_ = πTemp010
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
										πTemp002 = πF.MakeArgs(2)
										πTemp002[0] = πg.NewInt(1).ToObject()
										πTemp003 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
											continue
										}
										πTemp003[0] = µcolseps
										if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp003)
										πTemp002[1] = πTemp005
										if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp006 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp006 {
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
											πTemp007 = !isStop
										} else {
											πTemp007 = true
											µi = πTemp004
										}
										if πE != nil || !πTemp007 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 297: colspecs = [(colseps[i] - colseps[i - 1] - 1)
										πF.SetLineno(297)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp008 = µi
										if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
											continue
										}
										if πTemp009, πE = πg.GetItem(πF, µcolseps, πTemp008); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										if πTemp010, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
											continue
										}
										πTemp008 = πTemp010
										if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
											continue
										}
										if πTemp010, πE = πg.GetItem(πF, µcolseps, πTemp008); πE != nil {
											continue
										}
										if πTemp005, πE = πg.Sub(πF, πTemp009, πTemp010); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp005 = πSent
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
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							µcolspecs = πTemp002
							// line 300: onerow = [None for i in range(len(colseps) - 1)]
							πF.SetLineno(300)
							πTemp011 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 []*πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 bool
								_ = πTemp007
								var πTemp008 bool
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
										πTemp002 = πF.MakeArgs(1)
										πTemp004 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
											continue
										}
										πTemp004[0] = µcolseps
										if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp004)
										if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
											continue
										}
										πTemp002[0] = πTemp003
										if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp007 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp007 {
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
											πTemp008 = !isStop
										} else {
											πTemp008 = true
											µi = πTemp003
										}
										if πE != nil || !πTemp008 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 300: onerow = [None for i in range(len(colseps) - 1)]
										πF.SetLineno(300)
										if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp005 = πSent
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
							if πTemp009, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							µonerow = πTemp002
							// line 301: rows = [onerow[:] for i in range(len(rowseps) - 1)]
							πF.SetLineno(301)
							πTemp011 = make([]πg.Param, 0)
							πTemp009 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 []*πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 bool
								_ = πTemp007
								var πTemp008 bool
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
										πTemp002 = πF.MakeArgs(1)
										πTemp004 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
											continue
										}
										πTemp004[0] = µrowseps
										if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp004)
										if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
											continue
										}
										πTemp002[0] = πTemp003
										if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp007 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp007 {
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
											πTemp008 = !isStop
										} else {
											πTemp008 = true
											µi = πTemp003
										}
										if πE != nil || !πTemp008 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 301: rows = [onerow[:] for i in range(len(rowseps) - 1)]
										πF.SetLineno(301)
										if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µonerow, "onerow"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µonerow, πTemp003); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
									Label4:
										πTemp003 = πSent
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
							if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
								continue
							}
							µrows = πTemp002
							// line 303: remaining = (len(rowseps) - 1) * (len(colseps) - 1)
							πF.SetLineno(303)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrowseps, "rowseps"); πE != nil {
								continue
							}
							πTemp001[0] = µrowseps
							if πTemp012, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp010, πE = πg.Sub(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolseps, "colseps"); πE != nil {
								continue
							}
							πTemp001[0] = µcolseps
							if πTemp013, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp012, πE = πg.Sub(πF, πTemp014, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp010, πTemp012); πE != nil {
								continue
							}
							µremaining = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßcells, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp010); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp007 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp010, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp013}, πg.TieTarget{Target: &πTemp014}, πg.TieTarget{Target: &πTemp015}, πg.TieTarget{Target: &πTemp016}}}, πTemp010); πE != nil {
									continue
								}
								µtop = πTemp012
								µleft = πTemp013
								µbottom = πTemp014
								µright = πTemp015
								µblock = πTemp016
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 305: rownum = rowindex[top]
							πF.SetLineno(305)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							πTemp010 = µtop
							if πE = πg.CheckLocal(πF, µrowindex, "rowindex"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µrowindex, πTemp010); πE != nil {
								continue
							}
							µrownum = πTemp012
							// line 306: colnum = colindex[left]
							πF.SetLineno(306)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp010 = µleft
							if πE = πg.CheckLocal(πF, µcolindex, "colindex"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µcolindex, πTemp010); πE != nil {
								continue
							}
							µcolnum = πTemp012
							// line 307: assert rows[rownum][colnum] is None, (
							πF.SetLineno(307)
							if πE = πg.CheckLocal(πF, µrownum, "rownum"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Add(πF, µrownum, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolnum, "colnum"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.Add(πF, µcolnum, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp012 = πg.NewTuple2(πTemp013, πTemp014).ToObject()
							if πTemp010, πE = πg.Mod(πF, πg.NewStr("Cell (row %s, column %s) already used.").ToObject(), πTemp012); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolnum, "colnum"); πE != nil {
								continue
							}
							πTemp013 = µcolnum
							if πE = πg.CheckLocal(πF, µrownum, "rownum"); πE != nil {
								continue
							}
							πTemp015 = µrownum
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetItem(πF, µrows, πTemp015); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp013); πE != nil {
								continue
							}
							if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp012 = πg.GetBool(πTemp014 == πTemp013).ToObject()
							if πE = πg.Assert(πF, πTemp012, πTemp010); πE != nil {
								continue
							}
							// line 310: morerows = rowindex[bottom] - rownum - 1
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							πTemp013 = µbottom
							if πE = πg.CheckLocal(πF, µrowindex, "rowindex"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µrowindex, πTemp013); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrownum, "rownum"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Sub(πF, πTemp014, µrownum); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Sub(πF, πTemp012, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µmorerows = πTemp010
							// line 311: morecols = colindex[right] - colnum - 1
							πF.SetLineno(311)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp013 = µright
							if πE = πg.CheckLocal(πF, µcolindex, "colindex"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µcolindex, πTemp013); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolnum, "colnum"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Sub(πF, πTemp014, µcolnum); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Sub(πF, πTemp012, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µmorecols = πTemp010
							// line 312: remaining -= (morerows + 1) * (morecols + 1)
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µremaining, "remaining"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmorerows, "morerows"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Add(πF, µmorerows, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmorecols, "morecols"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Add(πF, µmorecols, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Mul(πF, πTemp012, πTemp013); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ISub(πF, µremaining, πTemp010); πE != nil {
								continue
							}
							µremaining = πTemp012
							// line 314: rows[rownum][colnum] = (morerows, morecols, top + 1, block)
							πF.SetLineno(314)
							if πE = πg.CheckLocal(πF, µmorerows, "morerows"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmorecols, "morecols"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Add(πF, µtop, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp010 = πg.NewTuple4(µmorerows, µmorecols, πTemp012, µblock).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πTemp010); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrownum, "rownum"); πE != nil {
								continue
							}
							πTemp013 = µrownum
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µrows, πTemp013); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolnum, "colnum"); πE != nil {
								continue
							}
							πTemp013 = µcolnum
							if πE = πg.SetItem(πF, πTemp014, πTemp013, πTemp012); πE != nil {
								continue
							}
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 315: assert remaining == 0, 'Unused cells remaining.'
							πF.SetLineno(315)
							if πE = πg.CheckLocal(πF, µremaining, "remaining"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µremaining, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("Unused cells remaining.").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 316: if self.head_body_sep:          # separate head rows from body rows
							πF.SetLineno(316)
						Label10:
							// line 317: numheadrows = rowindex[self.head_body_sep]
							πF.SetLineno(317)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp010
							if πE = πg.CheckLocal(πF, µrowindex, "rowindex"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µrowindex, πTemp002); πE != nil {
								continue
							}
							µnumheadrows = πTemp010
							// line 318: headrows = rows[:numheadrows]
							πF.SetLineno(318)
							if πE = πg.CheckLocal(πF, µnumheadrows, "numheadrows"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µnumheadrows, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µrows, πTemp002); πE != nil {
								continue
							}
							µheadrows = πTemp010
							// line 319: bodyrows = rows[numheadrows:]
							πF.SetLineno(319)
							if πE = πg.CheckLocal(πF, µnumheadrows, "numheadrows"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µnumheadrows, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µrows, πTemp002); πE != nil {
								continue
							}
							µbodyrows = πTemp010
							goto Label12
						Label11:
							// line 321: headrows = []
							πF.SetLineno(321)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µheadrows = πTemp002
							// line 322: bodyrows = rows
							πF.SetLineno(322)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							µbodyrows = µrows
							goto Label12
						Label12:
							// line 323: return (colspecs, headrows, bodyrows)
							πF.SetLineno(323)
							if πE = πg.CheckLocal(πF, µcolspecs, "colspecs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µheadrows, "headrows"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbodyrows, "bodyrows"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µcolspecs, µheadrows, µbodyrows).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßstructure_from_cells.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 285: """
					πF.SetLineno(285)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n        From the data collected by `scan_cell()`, convert to the final data\n        structure.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßstructure_from_cells); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("GridTableParser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGridTableParser.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 326: class SimpleTableParser(TableParser):
			πF.SetLineno(326)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTableParser); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SimpleTableParser", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 328: """
					πF.SetLineno(328)
					// line 328: """
					πF.SetLineno(328)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Parse a simple table using `parse()`.\n\n    Here's an example of a simple table::\n\n        =====  =====\n        col 1  col 2\n        =====  =====\n        1      Second column of row 1.\n        2      Second column of row 2.\n               Second line of paragraph.\n        3      - Second column of row 3.\n\n               - Second item in bullet\n                 list (row 3, column 2).\n        4 is a span\n        ------------\n        5\n        =====  =====\n\n    Top and bottom borders use '=', column span underlines use '-', column\n    separation is indicated with spaces.\n\n    Passing the above table to the `parse()` method will result in the\n    following data structure, whose interpretation is the same as for\n    `GridTableParser`::\n\n        ([5, 25],\n         [[(0, 0, 1, ['col 1']),\n           (0, 0, 1, ['col 2'])]],\n         [[(0, 0, 3, ['1']),\n           (0, 0, 3, ['Second column of row 1.'])],\n          [(0, 0, 4, ['2']),\n           (0, 0, 4, ['Second column of row 2.',\n                      'Second line of paragraph.'])],\n          [(0, 0, 6, ['3']),\n           (0, 0, 6, ['- Second column of row 3.',\n                      '',\n                      '- Second item in bullet',\n                      '  list (row 3, column 2).'])],\n          [(0, 1, 10, ['4 is a span'])],\n          [(0, 0, 12, ['5']),\n           (0, 0, 12, [''])]])\n    ").ToObject()); πE != nil {
						continue
					}
					// line 373: head_body_separator_pat = re.compile('=[ =]*$')
					πF.SetLineno(373)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=[ =]*$").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßhead_body_separator_pat.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 374: span_pat = re.compile('-[ -]*$')
					πF.SetLineno(374)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("-[ -]*$").ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßspan_pat.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 376: def setup(self, block):
					πF.SetLineno(376)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "block", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("setup", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µblock *πg.Object = πArgs[1]
						_ = µblock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
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
							// line 377: self.block = block[:]           # make a copy; it will be modified
							πF.SetLineno(377)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßblock, πTemp001); πE != nil {
								continue
							}
							// line 378: self.block.disconnect()         # don't propagate changes to parent
							πF.SetLineno(378)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdisconnect, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 380: self.block[0] = self.block[0].replace('=', '-')
							πF.SetLineno(380)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("=").ToObject()
							πTemp003[1] = πg.NewStr("-").ToObject()
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp001); πE != nil {
								continue
							}
							// line 381: self.block[-1] = self.block[-1].replace('=', '-')
							πF.SetLineno(381)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("=").ToObject()
							πTemp003[1] = πg.NewStr("-").ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp001); πE != nil {
								continue
							}
							// line 382: self.head_body_sep = None
							πF.SetLineno(382)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhead_body_sep, πTemp002); πE != nil {
								continue
							}
							// line 383: self.columns = []
							πF.SetLineno(383)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcolumns, πTemp002); πE != nil {
								continue
							}
							// line 384: self.border_end = None
							πF.SetLineno(384)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßborder_end, πTemp002); πE != nil {
								continue
							}
							// line 385: self.table = []
							πF.SetLineno(385)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtable, πTemp002); πE != nil {
								continue
							}
							// line 386: self.done = [-1] * len(block[0])
							πF.SetLineno(386)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µblock, πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdone, πTemp002); πE != nil {
								continue
							}
							// line 387: self.rowseps = {0: [0]}
							πF.SetLineno(387)
							πTemp007 = πg.NewDict()
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πTemp007.SetItem(πF, πg.NewInt(0).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrowseps, πTemp002); πE != nil {
								continue
							}
							// line 388: self.colseps = {0: [0]}
							πF.SetLineno(388)
							πTemp007 = πg.NewDict()
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πTemp007.SetItem(πF, πg.NewInt(0).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcolseps, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetup.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 390: def parse_table(self):
					πF.SetLineno(390)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("parse_table", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfirststart *πg.Object = πg.UnboundLocal
						_ = µfirststart
						var µfirstend *πg.Object = πg.UnboundLocal
						_ = µfirstend
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var µtext_found *πg.Object = πg.UnboundLocal
						_ = µtext_found
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var πTemp001 []*πg.Object
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
						var πTemp008 bool
						_ = πTemp008
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
							// line 391: """
							πF.SetLineno(391)
							// line 398: self.columns = self.parse_columns(self.block[0], 0)
							πF.SetLineno(398)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_columns, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcolumns, πTemp002); πE != nil {
								continue
							}
							// line 399: self.border_end = self.columns[-1][1]
							πF.SetLineno(399)
							πTemp002 = πg.NewInt(1).ToObject()
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßborder_end, πTemp002); πE != nil {
								continue
							}
							// line 400: firststart, firstend = self.columns[0]
							πF.SetLineno(400)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µfirststart = πTemp002
							µfirstend = πTemp004
							// line 401: offset = 1                      # skip top border
							πF.SetLineno(401)
							µoffset = πg.NewInt(1).ToObject()
							// line 402: start = 1
							πF.SetLineno(402)
							µstart = πg.NewInt(1).ToObject()
							// line 403: text_found = None
							πF.SetLineno(403)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtext_found = πTemp002
							// line 404: while offset < len(self.block):
							πF.SetLineno(404)
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, µoffset, πTemp004); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 405: line = self.block[offset]
							πF.SetLineno(405)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp002 = µoffset
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µline = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[0] = µline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßspan_pat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µfirststart, "firststart"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfirstend, "firstend"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µfirststart, µfirstend, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µtext_found, "text_found"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µtext_found); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label6
							}
							goto Label7
							// line 406: if self.span_pat.match(line):
							πF.SetLineno(406)
						Label4:
							// line 408: self.parse_row(self.block[start:offset], start,
							πF.SetLineno(408)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, µoffset, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001[1] = µstart
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µline, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, µoffset).ToObject()
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_row, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 410: start = offset + 1
							πF.SetLineno(410)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstart = πTemp002
							// line 411: text_found = None
							πF.SetLineno(411)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtext_found = πTemp002
							goto Label7
							// line 412: elif line[firststart:firstend].strip():
							πF.SetLineno(412)
						Label5:
							if πE = πg.CheckLocal(πF, µtext_found, "text_found"); πE != nil {
								continue
							}
							πTemp002 = µtext_found
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, µoffset, µstart); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label8:
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 414: if text_found and offset != start:
							πF.SetLineno(414)
						Label9:
							// line 415: self.parse_row(self.block[start:offset], start)
							πF.SetLineno(415)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, µoffset, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßblock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001[1] = µstart
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparse_row, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
						Label10:
							// line 416: start = offset
							πF.SetLineno(416)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							µstart = µoffset
							// line 417: text_found = 1
							πF.SetLineno(417)
							µtext_found = πg.NewInt(1).ToObject()
							goto Label7
							// line 418: elif not text_found:
							πF.SetLineno(418)
						Label6:
							// line 419: start = offset + 1
							πF.SetLineno(419)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstart = πTemp002
							goto Label7
						Label7:
							// line 420: offset += 1
							πF.SetLineno(420)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µoffset = πTemp002
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
					if πE = πClass.SetItem(πF, ßparse_table.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 391: """
					πF.SetLineno(391)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        First determine the column boundaries from the top border, then\n        process rows.  Each row may consist of multiple lines; accumulate\n        lines until a row is complete.  Call `self.parse_row` to finish the\n        job.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßparse_table); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 422: def parse_columns(self, line, offset):
					πF.SetLineno(422)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "line", Def: nil}
					πTemp004[2] = πg.Param{Name: "offset", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("parse_columns", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µline *πg.Object = πArgs[1]
						_ = µline
						var µoffset *πg.Object = πArgs[2]
						_ = µoffset
						var µcols *πg.Object = πg.UnboundLocal
						_ = µcols
						var µend *πg.Object = πg.UnboundLocal
						_ = µend
						var µbegin *πg.Object = πg.UnboundLocal
						_ = µbegin
						var πTemp001 []*πg.Object
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 423: """
							πF.SetLineno(423)
							// line 426: cols = []
							πF.SetLineno(426)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µcols = πTemp002
							// line 427: end = 0
							πF.SetLineno(427)
							µend = πg.NewInt(0).ToObject()
							// line 428: while True:
							πF.SetLineno(428)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 429: begin = line.find('-', end)
							πF.SetLineno(429)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-").ToObject()
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp001[1] = µend
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µline, ßfind, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbegin = πTemp005
							// line 430: end = line.find(' ', begin)
							πF.SetLineno(430)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
								continue
							}
							πTemp001[1] = µbegin
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µline, ßfind, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µend = πTemp005
							if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µbegin, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 431: if begin < 0:
							πF.SetLineno(431)
						Label4:
							// line 432: break
							πF.SetLineno(432)
							πTemp003 = true
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µend, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 433: if end < 0:
							πF.SetLineno(433)
						Label6:
							// line 434: end = len(line)
							πF.SetLineno(434)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001[0] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µend = πTemp005
							goto Label7
						Label7:
							// line 435: cols.append((begin, end))
							πF.SetLineno(435)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µbegin, µend).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µcols, "cols"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcols, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 436: if self.columns:
							πF.SetLineno(436)
						Label8:
							πTemp005 = πg.NewInt(1).ToObject()
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µcols, "cols"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µcols, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßborder_end, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							goto Label11
							// line 437: if cols[-1][1] != self.border_end:
							πF.SetLineno(437)
						Label10:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Column span incomplete in table line %s.").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"offset", µoffset},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTableMarkupError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, πTemp009); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 438: raise TableMarkupError('Column span incomplete in table '
							πF.SetLineno(438)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label11
						Label11:
							// line 442: cols[-1] = (cols[-1][0], self.columns[-1][1])
							πF.SetLineno(442)
							πTemp005 = πg.NewInt(0).ToObject()
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µcols, "cols"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µcols, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcols, "cols"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.SetItem(πF, µcols, πTemp006, πTemp005); πE != nil {
								continue
							}
							goto Label9
						Label9:
							// line 443: return cols
							πF.SetLineno(443)
							if πE = πg.CheckLocal(πF, µcols, "cols"); πE != nil {
								continue
							}
							πR = µcols
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßparse_columns.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 423: """
					πF.SetLineno(423)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Given a column span underline, return a list of (begin, end) pairs.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßparse_columns); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 445: def init_row(self, colspec, offset):
					πF.SetLineno(445)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "colspec", Def: nil}
					πTemp004[2] = πg.Param{Name: "offset", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("init_row", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcolspec *πg.Object = πArgs[1]
						_ = µcolspec
						var µoffset *πg.Object = πArgs[2]
						_ = µoffset
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µcells *πg.Object = πg.UnboundLocal
						_ = µcells
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var µend *πg.Object = πg.UnboundLocal
						_ = µend
						var µmorecols *πg.Object = πg.UnboundLocal
						_ = µmorecols
						var πTemp001 []*πg.Object
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.BaseException
						_ = πTemp012
						var πTemp013 *πg.Traceback
						_ = πTemp013
						var πTemp014 πg.KWArgs
						_ = πTemp014
						var πTemp015 []*πg.Object
						_ = πTemp015
						var πTemp016 []*πg.Object
						_ = πTemp016
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
							case 5:
								goto Label5
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 446: i = 0
							πF.SetLineno(446)
							µi = πg.NewInt(0).ToObject()
							// line 447: cells = []
							πF.SetLineno(447)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µcells = πTemp002
							if πE = πg.CheckLocal(πF, µcolspec, "colspec"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µcolspec); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
									continue
								}
								µstart = πTemp006
								µend = πTemp007
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 449: morecols = 0
							πF.SetLineno(449)
							µmorecols = πg.NewInt(0).ToObject()
							// line 450: try:
							πF.SetLineno(450)
							πF.PushCheckpoint(5)
							// line 451: assert start == self.columns[i][0]
							πF.SetLineno(451)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µstart, πTemp007); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp005, nil); πE != nil {
								continue
							}
							// line 452: while end != self.columns[i][1]:
							πF.SetLineno(452)
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp006 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.NE(πF, µend, πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 453: i += 1
							πF.SetLineno(453)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp005
							// line 454: morecols += 1
							πF.SetLineno(454)
							if πE = πg.CheckLocal(πF, µmorecols, "morecols"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µmorecols, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µmorecols = πTemp005
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp012, πTemp013 = πF.ExcInfo()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp004, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
							continue
							// line 455: except (AssertionError, IndexError):
							πF.SetLineno(455)
						Label9:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µoffset, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mod(πF, πg.NewStr("Column span alignment problem in table line %s.").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp014 = πg.KWArgs{
								{"offset", πTemp005},
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTableMarkupError); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, πTemp014); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 456: raise TableMarkupError('Column span alignment problem '
							πF.SetLineno(456)
							πE = πF.Raise(πTemp006, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 459: cells.append([0, morecols, offset, []])
							πF.SetLineno(459)
							πTemp001 = πF.MakeArgs(1)
							πTemp015 = make([]*πg.Object, 4)
							πTemp015[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µmorecols, "morecols"); πE != nil {
								continue
							}
							πTemp015[1] = µmorecols
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp015[2] = µoffset
							πTemp016 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp016...).ToObject()
							πTemp015[3] = πTemp005
							πTemp005 = πg.NewList(πTemp015...).ToObject()
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µcells, "cells"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcells, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 460: i += 1
							πF.SetLineno(460)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp005
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 461: return cells
							πF.SetLineno(461)
							if πE = πg.CheckLocal(πF, µcells, "cells"); πE != nil {
								continue
							}
							πR = µcells
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßinit_row.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 463: def parse_row(self, lines, start, spanline=None):
					πF.SetLineno(463)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "lines", Def: nil}
					πTemp004[2] = πg.Param{Name: "start", Def: nil}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004[3] = πg.Param{Name: "spanline", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("parse_row", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlines *πg.Object = πArgs[1]
						_ = µlines
						var µstart *πg.Object = πArgs[2]
						_ = µstart
						var µspanline *πg.Object = πArgs[3]
						_ = µspanline
						var µcolumns *πg.Object = πg.UnboundLocal
						_ = µcolumns
						var µspan_offset *πg.Object = πg.UnboundLocal
						_ = µspan_offset
						var µrow *πg.Object = πg.UnboundLocal
						_ = µrow
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µend *πg.Object = πg.UnboundLocal
						_ = µend
						var µcellblock *πg.Object = πg.UnboundLocal
						_ = µcellblock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 464: """
							πF.SetLineno(464)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp002 = µlines
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µspanline, "spanline"); πE != nil {
								continue
							}
							πTemp002 = µspanline
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
							// line 472: if not (lines or spanline):
							πF.SetLineno(472)
						Label2:
							// line 474: return
							πF.SetLineno(474)
							πR = πg.None
							continue
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µspanline, "spanline"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µspanline); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 475: if spanline:
							πF.SetLineno(475)
						Label4:
							// line 476: columns = self.parse_columns(*spanline)
							πF.SetLineno(476)
							if πE = πg.CheckLocal(πF, µspanline, "spanline"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparse_columns, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µspanline, nil, nil); πE != nil {
								continue
							}
							µcolumns = πTemp002
							// line 477: span_offset = spanline[1]
							πF.SetLineno(477)
							πTemp001 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µspanline, "spanline"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µspanline, πTemp001); πE != nil {
								continue
							}
							µspan_offset = πTemp002
							goto Label6
						Label5:
							// line 479: columns = self.columns[:]
							πF.SetLineno(479)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µcolumns = πTemp002
							// line 480: span_offset = start
							πF.SetLineno(480)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							µspan_offset = µstart
							goto Label6
						Label6:
							// line 481: self.check_columns(lines, start, columns)
							πF.SetLineno(481)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp005[0] = µlines
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp005[1] = µstart
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp005[2] = µcolumns
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_columns, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 482: row = self.init_row(columns, start)
							πF.SetLineno(482)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp005[0] = µcolumns
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp005[1] = µstart
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinit_row, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µrow = πTemp002
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp006[0] = µcolumns
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp003 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 484: start, end = columns[i]
							πF.SetLineno(484)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µcolumns, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
								continue
							}
							µstart = πTemp002
							µend = πTemp008
							// line 485: cellblock = lines.get_2D_block(0, start, len(lines), end)
							πF.SetLineno(485)
							πTemp005 = πF.MakeArgs(4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp005[1] = µstart
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp006[0] = µlines
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[2] = πTemp004
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp005[3] = µend
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlines, ßget_2D_block, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µcellblock = πTemp004
							// line 486: cellblock.disconnect()      # lines in cell can't sync with parent
							πF.SetLineno(486)
							if πE = πg.CheckLocal(πF, µcellblock, "cellblock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcellblock, ßdisconnect, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 487: cellblock.replace(self.double_width_pad_char, '')
							πF.SetLineno(487)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdouble_width_pad_char, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp005[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µcellblock, "cellblock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcellblock, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 488: row[i][3] = cellblock
							πF.SetLineno(488)
							if πE = πg.CheckLocal(πF, µcellblock, "cellblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µcellblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µrow, πTemp004); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(3).ToObject()
							if πE = πg.SetItem(πF, πTemp008, πTemp004, πTemp002); πE != nil {
								continue
							}
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 489: self.table.append(row)
							πF.SetLineno(489)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp005[0] = µrow
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtable, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßparse_row.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 464: """
					πF.SetLineno(464)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Given the text `lines` of a row, parse it and append to `self.table`.\n\n        The row is parsed according to the current column spec (either\n        `spanline` if provided or `self.columns`).  For each column, extract\n        text from each line, and check for text in column margins.  Finally,\n        adjust for insignificant whitespace.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßparse_row); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 491: def check_columns(self, lines, first_line, columns):
					πF.SetLineno(491)
					πTemp004 = make([]πg.Param, 4)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "lines", Def: nil}
					πTemp004[2] = πg.Param{Name: "first_line", Def: nil}
					πTemp004[3] = πg.Param{Name: "columns", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("check_columns", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlines *πg.Object = πArgs[1]
						_ = µlines
						var µfirst_line *πg.Object = πArgs[2]
						_ = µfirst_line
						var µcolumns *πg.Object = πArgs[3]
						_ = µcolumns
						var µlastcol *πg.Object = πg.UnboundLocal
						_ = µlastcol
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µstart *πg.Object = πg.UnboundLocal
						_ = µstart
						var µend *πg.Object = πg.UnboundLocal
						_ = µend
						var µnextstart *πg.Object = πg.UnboundLocal
						_ = µnextstart
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µnew_end *πg.Object = πg.UnboundLocal
						_ = µnew_end
						var µmain_start *πg.Object = πg.UnboundLocal
						_ = µmain_start
						var µmain_end *πg.Object = πg.UnboundLocal
						_ = µmain_end
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 πg.KWArgs
						_ = πTemp015
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 492: """
							πF.SetLineno(492)
							// line 499: columns.append((sys.maxsize, None))
							πF.SetLineno(499)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaxsize, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp003).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcolumns, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 500: lastcol = len(columns) - 2
							πF.SetLineno(500)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp001[0] = µcolumns
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Sub(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							µlastcol = πTemp002
							// line 502: lines = [strip_combining_chars(line) for line in lines]
							πF.SetLineno(502)
							πTemp005 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
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
										if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µlines); πE != nil {
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
											µline = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 502: lines = [strip_combining_chars(line) for line in lines]
										πF.SetLineno(502)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp005[0] = µline
										if πTemp004, πE = πg.ResolveGlobal(πF, ßstrip_combining_chars); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µlines = πTemp002
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp006[0] = µcolumns
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Sub(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp009 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 505: start, end = columns[i]
							πF.SetLineno(505)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µcolumns, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp008}}}, πTemp007); πE != nil {
								continue
							}
							µstart = πTemp004
							µend = πTemp008
							// line 506: nextstart = columns[i+1][0]
							πF.SetLineno(506)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp011
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µcolumns, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
								continue
							}
							µnextstart = πTemp007
							// line 507: offset = 0
							πF.SetLineno(507)
							µoffset = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µlines); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp010 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp012 = !isStop
							} else {
								πTemp012 = true
								µline = πTemp007
							}
							if πE != nil || !πTemp012 {
								continue
							}
							πF.PushCheckpoint(4)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlastcol, "lastcol"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Eq(πF, µi, µlastcol); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{µend, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µline, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp011, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp011
						Label7:
							if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnextstart, "nextstart"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{µend, µnextstart, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp007); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp008, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label9
							}
							goto Label10
							// line 509: if i == lastcol and line[end:].strip():
							πF.SetLineno(509)
						Label8:
							// line 510: text = line[start:].rstrip()
							πF.SetLineno(510)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{µstart, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µline, πTemp007); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp008, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp008
							// line 511: new_end = start + len(text)
							πF.SetLineno(511)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.Add(πF, µstart, πTemp011); πE != nil {
								continue
							}
							µnew_end = πTemp007
							// line 512: main_start, main_end = self.columns[-1]
							πF.SetLineno(512)
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp007); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp011}}}, πTemp008); πE != nil {
								continue
							}
							µmain_start = πTemp007
							µmain_end = πTemp011
							// line 513: columns[i] = (start, max(main_end, new_end))
							πF.SetLineno(513)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmain_end, "main_end"); πE != nil {
								continue
							}
							πTemp001[0] = µmain_end
							if πE = πg.CheckLocal(πF, µnew_end, "new_end"); πE != nil {
								continue
							}
							πTemp001[1] = µnew_end
							if πTemp008, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp007 = πg.NewTuple2(µstart, πTemp011).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp011 = µi
							if πE = πg.SetItem(πF, µcolumns, πTemp011, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew_end, "new_end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmain_end, "main_end"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GT(πF, µnew_end, µmain_end); πE != nil {
								continue
							}
							if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label11
							}
							goto Label12
							// line 514: if new_end > main_end:
							πF.SetLineno(514)
						Label11:
							// line 515: self.columns[-1] = (main_start, new_end)
							πF.SetLineno(515)
							if πE = πg.CheckLocal(πF, µmain_start, "main_start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew_end, "new_end"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(µmain_start, µnew_end).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp013 = πTemp014
							if πE = πg.SetItem(πF, πTemp011, πTemp013, πTemp008); πE != nil {
								continue
							}
							goto Label12
						Label12:
							goto Label10
							// line 516: elif line[end:nextstart].strip():
							πF.SetLineno(516)
						Label9:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst_line, "first_line"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Add(πF, µfirst_line, µoffset); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, πTemp011, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mod(πF, πg.NewStr("Text in column margin in table line %s.").ToObject(), πTemp008); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µfirst_line, "first_line"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µfirst_line, µoffset); πE != nil {
								continue
							}
							πTemp015 = πg.KWArgs{
								{"offset", πTemp007},
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßTableMarkupError); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, πTemp015); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 517: raise TableMarkupError('Text in column margin '
							πF.SetLineno(517)
							πE = πF.Raise(πTemp008, nil, nil)
							continue
							goto Label10
						Label10:
							// line 520: offset += 1
							πF.SetLineno(520)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IAdd(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µoffset = πTemp007
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 521: columns.pop()
							πF.SetLineno(521)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcolumns, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheck_columns.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 492: """
					πF.SetLineno(492)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Check for text in column margins and text overflow in the last column.\n        Raise TableMarkupError if anything but whitespace is in column margins.\n        Adjust the end value for the last column if there is text overflow.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_columns); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 523: def structure_from_cells(self):
					πF.SetLineno(523)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("structure_from_cells", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcolspecs *πg.Object = πg.UnboundLocal
						_ = µcolspecs
						var µfirst_body_row *πg.Object = πg.UnboundLocal
						_ = µfirst_body_row
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 *πg.Object
						_ = πTemp015
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 524: colspecs = [end - start for start, end in self.columns]
							πF.SetLineno(524)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µstart *πg.Object = πg.UnboundLocal
								_ = µstart
								var µend *πg.Object = πg.UnboundLocal
								_ = µend
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßcolumns, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
											µstart = πTemp005
											µend = πTemp006
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 524: colspecs = [end - start for start, end in self.columns]
										πF.SetLineno(524)
										if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.Sub(πF, µend, µstart); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp002, nil
									Label4:
										πTemp005 = πSent
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µcolspecs = πTemp001
							// line 525: first_body_row = 0
							πF.SetLineno(525)
							µfirst_body_row = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 526: if self.head_body_sep:
							πF.SetLineno(526)
						Label1:
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtable, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(3)
							πTemp008 = πg.NewInt(2).ToObject()
							πTemp011 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp013 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, µself, ßtable, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, πTemp015, πTemp013); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp012, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßhead_body_sep, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label6
							}
							goto Label7
							// line 528: if self.table[i][0][2] > self.head_body_sep:
							πF.SetLineno(528)
						Label6:
							// line 529: first_body_row = i
							πF.SetLineno(529)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							µfirst_body_row = µi
							// line 530: break
							πF.SetLineno(530)
							πTemp005 = true
							continue
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							goto Label2
						Label2:
							// line 531: return (colspecs, self.table[:first_body_row],
							πF.SetLineno(531)
							if πE = πg.CheckLocal(πF, µcolspecs, "colspecs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfirst_body_row, "first_body_row"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µfirst_body_row, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßtable, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfirst_body_row, "first_body_row"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{µfirst_body_row, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßtable, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µcolspecs, πTemp008, πTemp010).ToObject()
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
					if πE = πClass.SetItem(πF, ßstructure_from_cells.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SimpleTableParser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSimpleTableParser.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 535: def update_dict_of_lists(master, newdata):
			πF.SetLineno(535)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "master", Def: nil}
			πTemp006[1] = πg.Param{Name: "newdata", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("update_dict_of_lists", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/tableparser.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmaster *πg.Object = πArgs[0]
				_ = µmaster
				var µnewdata *πg.Object = πArgs[1]
				_ = µnewdata
				var µkey *πg.Object = πg.UnboundLocal
				_ = µkey
				var µvalues *πg.Object = πg.UnboundLocal
				_ = µvalues
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
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
					// line 536: """
					πF.SetLineno(536)
					if πE = πg.CheckLocal(πF, µnewdata, "newdata"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnewdata, ßitems, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
							continue
						}
						µkey = πTemp003
						µvalues = πTemp006
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 542: master.setdefault(key, []).extend(values)
					πF.SetLineno(542)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					πTemp007[0] = µvalues
					πTemp008 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp008[0] = µkey
					πTemp009 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp009...).ToObject()
					πTemp008[1] = πTemp002
					if πE = πg.CheckLocal(πF, µmaster, "master"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmaster, ßsetdefault, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßextend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßupdate_dict_of_lists.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 536: """
			πF.SetLineno(536)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Extend the list values of `master` with those from `newdata`.\n\n    Both parameters must be dictionaries containing list values.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßupdate_dict_of_lists); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("tableparser", Code)
}
