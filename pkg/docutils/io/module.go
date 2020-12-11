package io

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/codecs"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßBOM_UTF16_BE := πg.InternStr("BOM_UTF16_BE")
		ßBOM_UTF16_LE := πg.InternStr("BOM_UTF16_LE")
		ßBOM_UTF8 := πg.InternStr("BOM_UTF8")
		ßBinaryFileOutput := πg.InternStr("BinaryFileOutput")
		ßDocTreeInput := πg.InternStr("DocTreeInput")
		ßErrorOutput := πg.InternStr("ErrorOutput")
		ßErrorString := πg.InternStr("ErrorString")
		ßFalse := πg.InternStr("False")
		ßFileInput := πg.InternStr("FileInput")
		ßFileOutput := πg.InternStr("FileOutput")
		ßIOError := πg.InternStr("IOError")
		ßInput := πg.InternStr("Input")
		ßInputError := πg.InternStr("InputError")
		ßLookupError := πg.InternStr("LookupError")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßNullInput := πg.InternStr("NullInput")
		ßNullOutput := πg.InternStr("NullOutput")
		ßOutput := πg.InternStr("Output")
		ßOutputError := πg.InternStr("OutputError")
		ßStringInput := πg.InternStr("StringInput")
		ßStringOutput := πg.InternStr("StringOutput")
		ßTransformSpec := πg.InternStr("TransformSpec")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeError := πg.InternStr("UnicodeError")
		ßValueError := πg.InternStr("ValueError")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß_stderr := πg.InternStr("_stderr")
		ßascii := πg.InternStr("ascii")
		ßautoclose := πg.InternStr("autoclose")
		ßb := πg.InternStr("b")
		ßbuffer := πg.InternStr("buffer")
		ßbyte_order_marks := πg.InternStr("byte_order_marks")
		ßbytes := πg.InternStr("bytes")
		ßcheck_encoding := πg.InternStr("check_encoding")
		ßclose := πg.InternStr("close")
		ßcodecs := πg.InternStr("codecs")
		ßcoding_slug := πg.InternStr("coding_slug")
		ßcompile := πg.InternStr("compile")
		ßcomponent_type := πg.InternStr("component_type")
		ßdecode := πg.InternStr("decode")
		ßdefault_destination_path := πg.InternStr("default_destination_path")
		ßdefault_source_path := πg.InternStr("default_source_path")
		ßdestination := πg.InternStr("destination")
		ßdestination_path := πg.InternStr("destination_path")
		ßdetermine_encoding_from_data := πg.InternStr("determine_encoding_from_data")
		ßencode := πg.InternStr("encode")
		ßencoding := πg.InternStr("encoding")
		ßerrno := πg.InternStr("errno")
		ßerror_handler := πg.InternStr("error_handler")
		ßerrors := πg.InternStr("errors")
		ßgroup := πg.InternStr("group")
		ßhasattr := πg.InternStr("hasattr")
		ßinput := πg.InternStr("input")
		ßinsert := πg.InternStr("insert")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlinesep := πg.InternStr("linesep")
		ßlocale_encoding := πg.InternStr("locale_encoding")
		ßlookup := πg.InternStr("lookup")
		ßlower := πg.InternStr("lower")
		ßmode := πg.InternStr("mode")
		ßname := πg.InternStr("name")
		ßopen := πg.InternStr("open")
		ßopened := πg.InternStr("opened")
		ßos := πg.InternStr("os")
		ßoutput := πg.InternStr("output")
		ßprint := πg.InternStr("print")
		ßr := πg.InternStr("r")
		ßrU := πg.InternStr("rU")
		ßrb := πg.InternStr("rb")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßreadlines := πg.InternStr("readlines")
		ßreplace := πg.InternStr("replace")
		ßrepr := πg.InternStr("repr")
		ßsearch := πg.InternStr("search")
		ßsource := πg.InternStr("source")
		ßsource_path := πg.InternStr("source_path")
		ßsplitlines := πg.InternStr("splitlines")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßstr := πg.InternStr("str")
		ßstrerror := πg.InternStr("strerror")
		ßstrict := πg.InternStr("strict")
		ßsuccessful_encoding := πg.InternStr("successful_encoding")
		ßsys := πg.InternStr("sys")
		ßunicode := πg.InternStr("unicode")
		ßversion_info := πg.InternStr("version_info")
		ßw := πg.InternStr("w")
		ßwb := πg.InternStr("wb")
		ßwrite := πg.InternStr("write")
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
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nI/O classes provide a uniform API for low-level input and output.  Subclasses\nexist for a variety of input/output mechanisms.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: from __future__ import print_function
			πF.SetLineno(9)
			// line 11: __docformat__ = 'reStructuredText'
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 13: import sys
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import os
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: import re
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import codecs
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "codecs"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcodecs.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: from docutils import TransformSpec
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformSpec); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformSpec.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from docutils.utils.error_reporting import locale_encoding, ErrorString, ErrorOutput
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßlocale_encoding); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorString.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorOutput); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorOutput.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 20: if sys.version_info >= (3, 0):
			πF.SetLineno(20)
		Label1:
			// line 21: unicode = str  # noqa
			πF.SetLineno(21)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 24: class InputError(IOError): pass
			πF.SetLineno(24)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("InputError", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 24: class InputError(IOError): pass
					πF.SetLineno(24)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("InputError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInputError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 25: class OutputError(IOError): pass
			πF.SetLineno(25)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OutputError", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 25: class OutputError(IOError): pass
					πF.SetLineno(25)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("OutputError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOutputError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 27: def check_encoding(stream, encoding):
			πF.SetLineno(27)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "stream", Def: nil}
			πTemp007[1] = πg.Param{Name: "encoding", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("check_encoding", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstream *πg.Object = πArgs[0]
				_ = µstream
				var µencoding *πg.Object = πArgs[1]
				_ = µencoding
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
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
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 28: """Test, whether the encoding of `stream` matches `encoding`.
					πF.SetLineno(28)
					// line 37: try:
					πF.SetLineno(37)
					πF.PushCheckpoint(2)
					// line 38: return codecs.lookup(stream.encoding) == codecs.lookup(encoding)
					πF.SetLineno(38)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstream, "stream"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstream, ßencoding, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlookup, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp002[0] = µencoding
					if πTemp004, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßlookup, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 39: except (LookupError, AttributeError, TypeError):
					πF.SetLineno(39)
				Label3:
					// line 40: return None
					πF.SetLineno(40)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πR = πTemp001
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
			if πE = πF.Globals().SetItem(πF, ßcheck_encoding.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 28: """Test, whether the encoding of `stream` matches `encoding`.
			πF.SetLineno(28)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Test, whether the encoding of `stream` matches `encoding`.\n\n    Returns\n\n    :None:  if `encoding` or `stream.encoding` are not a valid encoding\n            argument (e.g. ``None``) or `stream.encoding is missing.\n    :True:  if the encoding argument resolves to the same value as `encoding`,\n    :False: if the encodings differ.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßcheck_encoding); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 43: class Input(TransformSpec):
			πF.SetLineno(43)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßTransformSpec); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Input", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
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
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class for input wrappers.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 49: component_type = 'input'
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ßcomponent_type.ToObject(), ßinput.ToObject()); πE != nil {
						continue
					}
					// line 51: default_source_path = None
					πF.SetLineno(51)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdefault_source_path.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 53: def __init__(self, source=None, source_path=None, encoding=None,
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "source", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "source_path", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "encoding", Def: πTemp003}
					πTemp002[4] = πg.Param{Name: "error_handler", Def: ßstrict.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πArgs[1]
						_ = µsource
						var µsource_path *πg.Object = πArgs[2]
						_ = µsource_path
						var µencoding *πg.Object = πArgs[3]
						_ = µencoding
						var µerror_handler *πg.Object = πArgs[4]
						_ = µerror_handler
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 55: self.encoding = encoding
							πF.SetLineno(55)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µencoding); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp001); πE != nil {
								continue
							}
							// line 56: """Text encoding for the input source."""
							πF.SetLineno(56)
							// line 58: self.error_handler = error_handler
							πF.SetLineno(58)
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µerror_handler); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßerror_handler, πTemp001); πE != nil {
								continue
							}
							// line 59: """Text decoding error handler."""
							πF.SetLineno(59)
							// line 61: self.source = source
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp001); πE != nil {
								continue
							}
							// line 62: """The source of input data."""
							πF.SetLineno(62)
							// line 64: self.source_path = source_path
							πF.SetLineno(64)
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µsource_path); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource_path, πTemp001); πE != nil {
								continue
							}
							// line 65: """A text reference to the source."""
							πF.SetLineno(65)
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µsource_path); πE != nil {
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
							// line 67: if not source_path:
							πF.SetLineno(67)
						Label1:
							// line 68: self.source_path = self.default_source_path
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_source_path, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource_path, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 70: self.successful_encoding = None
							πF.SetLineno(70)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsuccessful_encoding, πTemp003); πE != nil {
								continue
							}
							// line 71: """The encoding that successfully decoded the source data."""
							πF.SetLineno(71)
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
					// line 73: def __repr__(self):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 74: return '%s: source=%r, source_path=%r' % (self.__class__, self.source,
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßsource_path, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: source=%r, source_path=%r").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 77: def read(self):
					πF.SetLineno(77)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 78: raise NotImplementedError
							πF.SetLineno(78)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 80: def decode(self, data):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("decode", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µencodings *πg.Object = πg.UnboundLocal
						_ = µencodings
						var µdata_encoding *πg.Object = πg.UnboundLocal
						_ = µdata_encoding
						var µenc *πg.Object = πg.UnboundLocal
						_ = µenc
						var µdecoded *πg.Object = πg.UnboundLocal
						_ = µdecoded
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []πg.Param
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
							case 18:
								goto Label18
							case 14:
								goto Label14
							case 15:
								goto Label15
							default:
								panic("unexpected function state")
							}
							// line 81: """
							πF.SetLineno(81)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßlower, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, ßunicode.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 90: if self.encoding and self.encoding.lower() == 'unicode':
							πF.SetLineno(90)
						Label2:
							// line 91: assert isinstance(data, unicode), (
							πF.SetLineno(91)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Assert(πF, πTemp003, πg.NewStr("input encoding is \"unicode\" but input is not a unicode object").ToObject()); πE != nil {
								continue
							}
							goto Label3
						Label3:
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 94: if isinstance(data, unicode):
							πF.SetLineno(94)
						Label4:
							// line 96: return data
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 97: if self.encoding:
							πF.SetLineno(97)
						Label6:
							// line 100: encodings = [self.encoding]
							πF.SetLineno(100)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µencodings = πTemp001
							goto Label8
						Label7:
							// line 102: data_encoding = self.determine_encoding_from_data(data)
							πF.SetLineno(102)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdetermine_encoding_from_data, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µdata_encoding = πTemp003
							if πE = πg.CheckLocal(πF, µdata_encoding, "data_encoding"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µdata_encoding); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 103: if data_encoding:
							πF.SetLineno(103)
						Label9:
							// line 106: encodings = [data_encoding]
							πF.SetLineno(106)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µdata_encoding, "data_encoding"); πE != nil {
								continue
							}
							πTemp006[0] = µdata_encoding
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µencodings = πTemp001
							goto Label11
						Label10:
							// line 111: encodings = ['utf-8', 'latin-1']
							πF.SetLineno(111)
							πTemp006 = make([]*πg.Object, 2)
							πTemp006[0] = πg.NewStr("utf-8").ToObject()
							πTemp006[1] = πg.NewStr("latin-1").ToObject()
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							µencodings = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label12
							}
							goto Label13
							// line 112: if locale_encoding:
							πF.SetLineno(112)
						Label12:
							// line 113: encodings.insert(1, locale_encoding)
							πF.SetLineno(113)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlocale_encoding); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µencodings, "encodings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µencodings, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label13
						Label13:
							goto Label11
						Label11:
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µencodings, "encodings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µencodings); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp002 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label16
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µenc = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(14)
							// line 115: try:
							πF.SetLineno(115)
							πF.PushCheckpoint(18)
							// line 116: decoded = unicode(data, enc, self.error_handler)
							πF.SetLineno(116)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
								continue
							}
							πTemp006[1] = µenc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßerror_handler, nil); πE != nil {
								continue
							}
							πTemp006[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µdecoded = πTemp004
							// line 117: self.successful_encoding = enc
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µenc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsuccessful_encoding, πTemp003); πE != nil {
								continue
							}
							// line 119: return decoded.replace(u'\ufeff', u'')
							πF.SetLineno(119)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewUnicode("\xef\xbb\xbf").ToObject()
							πTemp006[1] = πg.NewUnicode("").ToObject()
							if πE = πg.CheckLocal(πF, µdecoded, "decoded"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdecoded, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πR = πTemp004
							continue
							πF.PopCheckpoint()
							goto Label17
						Label18:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp007, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label19
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 120: except (UnicodeError, LookupError) as err:
							πF.SetLineno(120)
						Label19:
							µerr = πTemp008.ToObject()
							// line 121: error = err # in Python 3, the <exception instance> is
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							µerror = µerr
							πF.RestoreExc(nil, nil)
							goto Label17
						Label17:
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							πTemp006 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							πTemp011 = make([]πg.Param, 0)
							πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µenc *πg.Object = πg.UnboundLocal
								_ = µenc
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
										if πE = πg.CheckLocal(πF, µencodings, "encodings"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µencodings); πE != nil {
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
											µenc = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 125: '%s.\n(%s)' % (', '.join([repr(enc) for enc in encodings]),
										πF.SetLineno(125)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µenc, "enc"); πE != nil {
											continue
										}
										πTemp005[0] = µenc
										if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
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
							if πTemp012, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ListType.Call(πF, πg.Args{πTemp012}, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp004.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp010[0] = µerror
							if πTemp004, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp004.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp003 = πg.NewTuple2(πTemp012, πTemp013).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unable to decode input data.  Tried the following encodings: %s.\n(%s)").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 123: raise UnicodeError(
							πF.SetLineno(123)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdecode.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 81: """
					πF.SetLineno(81)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Decode a string, `data`, heuristically.\n        Raise UnicodeError if unsuccessful.\n\n        The client application should call ``locale.setlocale`` at the\n        beginning of processing::\n\n            locale.setlocale(locale.LC_ALL, '')\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdecode); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 128: coding_slug = re.compile(br"coding[:=]\s*([-\w.]+)")
					πF.SetLineno(128)
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = πg.NewStr("coding[:=]\\s*([-\\w.]+)").ToObject()
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßcoding_slug.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 129: """Encoding declaration pattern."""
					πF.SetLineno(129)
					// line 131: byte_order_marks = ((codecs.BOM_UTF8, 'utf-8'),
					πF.SetLineno(131)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßcodecs); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßBOM_UTF8, nil); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(πTemp010, πg.NewStr("utf-8").ToObject()).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßcodecs); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßBOM_UTF16_BE, nil); πE != nil {
						continue
					}
					πTemp009 = πg.NewTuple2(πTemp011, πg.NewStr("utf-16-be").ToObject()).ToObject()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßcodecs); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßBOM_UTF16_LE, nil); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(πTemp012, πg.NewStr("utf-16-le").ToObject()).ToObject()
					πTemp006 = πg.NewTuple3(πTemp007, πTemp009, πTemp010).ToObject()
					if πE = πClass.SetItem(πF, ßbyte_order_marks.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 134: """Sequence of (start_bytes, encoding) tuples for encoding detection.
					πF.SetLineno(134)
					// line 138: def determine_encoding_from_data(self, data):
					πF.SetLineno(138)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("determine_encoding_from_data", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µstart_bytes *πg.Object = πg.UnboundLocal
						_ = µstart_bytes
						var µencoding *πg.Object = πg.UnboundLocal
						_ = µencoding
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µmatch *πg.Object = πg.UnboundLocal
						_ = µmatch
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 139: """
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbyte_order_marks, nil); πE != nil {
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
								µstart_bytes = πTemp005
								µencoding = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstart_bytes, "start_bytes"); πE != nil {
								continue
							}
							πTemp007[0] = µstart_bytes
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdata, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 145: if data.startswith(start_bytes):
							πF.SetLineno(145)
						Label4:
							// line 146: return encoding
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πR = µencoding
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µdata, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µline = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 149: match = self.coding_slug.search(line)
							πF.SetLineno(149)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp007[0] = µline
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcoding_slug, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µmatch = πTemp002
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmatch); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 150: if match:
							πF.SetLineno(150)
						Label9:
							// line 151: return match.group(1).decode('ascii')
							πF.SetLineno(151)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = ßascii.ToObject()
							πTemp009 = πF.MakeArgs(1)
							πTemp009[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πR = πTemp005
							continue
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 152: return None
							πF.SetLineno(152)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdetermine_encoding_from_data.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 139: """
					πF.SetLineno(139)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Try to determine the encoding of `data` by looking *in* `data`.\n        Check for a byte order mark (BOM) or an encoding declaration.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdetermine_encoding_from_data); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Input").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßInput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 155: class Output(TransformSpec):
			πF.SetLineno(155)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßTransformSpec); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Output", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 157: """
					πF.SetLineno(157)
					// line 157: """
					πF.SetLineno(157)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class for output wrappers.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 161: component_type = 'output'
					πF.SetLineno(161)
					if πE = πClass.SetItem(πF, ßcomponent_type.ToObject(), ßoutput.ToObject()); πE != nil {
						continue
					}
					// line 163: default_destination_path = None
					πF.SetLineno(163)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdefault_destination_path.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 165: def __init__(self, destination=None, destination_path=None,
					πF.SetLineno(165)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "destination", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "destination_path", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "encoding", Def: πTemp003}
					πTemp002[4] = πg.Param{Name: "error_handler", Def: ßstrict.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdestination *πg.Object = πArgs[1]
						_ = µdestination
						var µdestination_path *πg.Object = πArgs[2]
						_ = µdestination_path
						var µencoding *πg.Object = πArgs[3]
						_ = µencoding
						var µerror_handler *πg.Object = πArgs[4]
						_ = µerror_handler
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 167: self.encoding = encoding
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µencoding); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp001); πE != nil {
								continue
							}
							// line 168: """Text encoding for the output destination."""
							πF.SetLineno(168)
							// line 170: self.error_handler = error_handler or 'strict'
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							πTemp001 = µerror_handler
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp001 = ßstrict.ToObject()
						Label1:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßerror_handler, πTemp003); πE != nil {
								continue
							}
							// line 171: """Text encoding error handler."""
							πF.SetLineno(171)
							// line 173: self.destination = destination
							πF.SetLineno(173)
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdestination); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination, πTemp001); πE != nil {
								continue
							}
							// line 174: """The destination for output data."""
							πF.SetLineno(174)
							// line 176: self.destination_path = destination_path
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdestination_path); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination_path, πTemp001); πE != nil {
								continue
							}
							// line 177: """A text reference to the destination."""
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µdestination_path); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 179: if not destination_path:
							πF.SetLineno(179)
						Label2:
							// line 180: self.destination_path = self.default_destination_path
							πF.SetLineno(180)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_destination_path, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination_path, πTemp003); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 182: def __repr__(self):
					πF.SetLineno(182)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 183: return ('%s: destination=%r, destination_path=%r'
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdestination_path, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp003, πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: destination=%r, destination_path=%r").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 186: def write(self, data):
					πF.SetLineno(186)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 187: """`data` is a Unicode string, to be encoded by `self.encode`."""
							πF.SetLineno(187)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 188: raise NotImplementedError
							πF.SetLineno(188)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 187: """`data` is a Unicode string, to be encoded by `self.encode`."""
					πF.SetLineno(187)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("`data` is a Unicode string, to be encoded by `self.encode`.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßwrite); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 190: def encode(self, data):
					πF.SetLineno(190)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("encode", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßlower, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, ßunicode.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 191: if self.encoding and self.encoding.lower() == 'unicode':
							πF.SetLineno(191)
						Label2:
							// line 192: assert isinstance(data, unicode), (
							πF.SetLineno(192)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.Assert(πF, πTemp003, πg.NewStr("the encoding given is \"unicode\" but the output is not a Unicode string").ToObject()); πE != nil {
								continue
							}
							// line 195: return data
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
							goto Label3
						Label3:
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 196: if not isinstance(data, unicode):
							πF.SetLineno(196)
						Label4:
							// line 198: return data
							πF.SetLineno(198)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
							goto Label6
						Label5:
							// line 200: return data.encode(self.encoding, self.error_handler)
							πF.SetLineno(200)
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror_handler, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πR = πTemp003
							continue
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
					if πE = πClass.SetItem(πF, ßencode.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Output").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOutput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 203: class FileInput(Input):
			πF.SetLineno(203)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßInput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FileInput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				var πTemp007 bool
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 205: """
					πF.SetLineno(205)
					// line 205: """
					πF.SetLineno(205)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Input for single, simple file-like objects.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 208: def __init__(self, source=None, source_path=None,
					πF.SetLineno(208)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "source", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "source_path", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "encoding", Def: πTemp003}
					πTemp002[4] = πg.Param{Name: "error_handler", Def: ßstrict.ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "autoclose", Def: πTemp003}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					if πTemp004, πE = πg.GE(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label1
					}
					πTemp003 = ßr.ToObject()
					goto Label2
				Label1:
					πTemp003 = ßrU.ToObject()
				Label2:
					πTemp002[6] = πg.Param{Name: "mode", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsource *πg.Object = πArgs[1]
						_ = µsource
						var µsource_path *πg.Object = πArgs[2]
						_ = µsource_path
						var µencoding *πg.Object = πArgs[3]
						_ = µencoding
						var µerror_handler *πg.Object = πArgs[4]
						_ = µerror_handler
						var µautoclose *πg.Object = πArgs[5]
						_ = µautoclose
						var µmode *πg.Object = πArgs[6]
						_ = µmode
						var µkwargs *πg.Object = πg.UnboundLocal
						_ = µkwargs
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
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
							case 17:
								goto Label17
							case 12:
								goto Label12
							default:
								panic("unexpected function state")
							}
							// line 212: """
							πF.SetLineno(212)
							// line 225: Input.__init__(self, source, source_path, encoding, error_handler)
							πF.SetLineno(225)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[1] = µsource
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							πTemp001[2] = µsource_path
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001[3] = µencoding
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							πTemp001[4] = µerror_handler
							if πTemp002, πE = πg.ResolveGlobal(πF, ßInput); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 226: self.autoclose = autoclose
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µautoclose, "autoclose"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µautoclose); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautoclose, πTemp002); πE != nil {
								continue
							}
							// line 227: self._stderr = ErrorOutput()
							πF.SetLineno(227)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stderr, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µsource == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp003, πE = πg.GE(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcheck_encoding); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006 == πTemp005).ToObject()
							πTemp002 = πTemp003
						Label2:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 229: if source is None:
							πF.SetLineno(229)
						Label1:
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µsource_path); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 230: if source_path:
							πF.SetLineno(230)
						Label5:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp002, πE = πg.GE(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 232: if sys.version_info >= (3, 0):
							πF.SetLineno(232)
						Label8:
							// line 233: kwargs = {'encoding': self.encoding,
							πF.SetLineno(233)
							πTemp007 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßencoding.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror_handler, nil); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßerrors.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp007.ToObject()
							µkwargs = πTemp002
							goto Label10
						Label9:
							// line 236: kwargs = {}
							πF.SetLineno(236)
							πTemp007 = πg.NewDict()
							πTemp002 = πTemp007.ToObject()
							µkwargs = πTemp002
							goto Label10
						Label10:
							// line 237: try:
							πF.SetLineno(237)
							πF.PushCheckpoint(12)
							// line 238: self.source = open(source_path, mode, **kwargs)
							πF.SetLineno(238)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							πTemp001[0] = µsource_path
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp001[1] = µmode
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label11
						Label12:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 239: except IOError as error:
							πF.SetLineno(239)
						Label13:
							µerror = πTemp008.ToObject()
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerror, ßerrno, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µerror, ßstrerror, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							πTemp001[2] = µsource_path
							if πTemp002, πE = πg.ResolveGlobal(πF, ßInputError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 240: raise InputError(error.errno, error.strerror, source_path)
							πF.SetLineno(240)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label11
						Label11:
							goto Label7
						Label6:
							// line 242: self.source = sys.stdin
							πF.SetLineno(242)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstdin, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							goto Label4
							// line 243: elif (sys.version_info >= (3, 0) and
							πF.SetLineno(243)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp006, ßencoding, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp010).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Encoding clash: encoding given is \"%s\" but source is opened with encoding \"%s\".").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 246: raise UnicodeError('Encoding clash: encoding given is "%s" '
							πF.SetLineno(246)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µsource_path); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 249: if not source_path:
							πF.SetLineno(249)
						Label14:
							// line 250: try:
							πF.SetLineno(250)
							πF.PushCheckpoint(17)
							// line 251: self.source_path = self.source.name
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsource_path, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label16
						Label17:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 252: except AttributeError:
							πF.SetLineno(252)
						Label18:
							// line 253: pass
							πF.SetLineno(253)
							πF.RestoreExc(nil, nil)
							goto Label16
						Label16:
							goto Label15
						Label15:
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
					// line 212: """
					πF.SetLineno(212)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        :Parameters:\n            - `source`: either a file-like object (which is read directly), or\n              `None` (which implies `sys.stdin` if no `source_path` given).\n            - `source_path`: a path to a file, which is opened and then read.\n            - `encoding`: the expected text encoding of the input file.\n            - `error_handler`: the encoding error handler to use.\n            - `autoclose`: close automatically after read (except when\n              `sys.stdin` is the source).\n            - `mode`: how the file is to be opened (see standard function\n              `open`). The default 'rU' provides universal newline support\n              for text files with Python 2.x.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 255: def read(self):
					πF.SetLineno(255)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πg.UnboundLocal
						_ = µdata
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
						var µb_source *πg.Object = πg.UnboundLocal
						_ = µb_source
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
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
							// line 256: """
							πF.SetLineno(256)
							// line 259: try:
							πF.SetLineno(259)
							πF.PushCheckpoint(1)
							πF.PushCheckpoint(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßstdin, nil); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label3
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp003, πE = πg.GE(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label3:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 260: if self.source is sys.stdin and sys.version_info >= (3, 0):
							πF.SetLineno(260)
						Label4:
							// line 262: data = self.source.buffer.read()
							πF.SetLineno(262)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßbuffer, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdata = πTemp003
							// line 264: data = b'\n'.join(data.splitlines()) + b'\n'
							πF.SetLineno(264)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdata, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µdata = πTemp001
							goto Label6
						Label5:
							// line 266: data = self.source.read()
							πF.SetLineno(266)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdata = πTemp001
							goto Label6
						Label6:
							πF.PopCheckpoint()
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πTemp002, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 267: except (UnicodeError, LookupError) as err: # (in Py3k read() decodes)
							πF.SetLineno(267)
						Label7:
							µerr = πTemp008.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp010).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsource_path, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label8:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 268: if not self.encoding and self.source_path:
							πF.SetLineno(268)
						Label9:
							// line 270: b_source = open(self.source_path, 'rb')
							πF.SetLineno(270)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource_path, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							πTemp007[1] = ßrb.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µb_source = πTemp003
							// line 271: data = b_source.read()
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µb_source, "b_source"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µb_source, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdata = πTemp003
							// line 272: b_source.close()
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µb_source, "b_source"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µb_source, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 274: data = b'\n'.join(data.splitlines()) + b'\n'
							πF.SetLineno(274)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdata, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µdata = πTemp001
							goto Label11
						Label10:
							// line 276: raise
							πF.SetLineno(276)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label11
						Label11:
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßautoclose, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label12
							}
							goto Label13
							// line 278: if self.autoclose:
							πF.SetLineno(278)
						Label12:
							// line 279: self.close()
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label13
						Label13:
							if πTemp008 != nil {
								πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							// line 280: return self.decode(data)
							πF.SetLineno(280)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp007[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 256: """
					πF.SetLineno(256)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Read and decode a single file and return the data (Unicode string).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßread); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 282: def readlines(self):
					πF.SetLineno(282)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("readlines", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 283: """
							πF.SetLineno(283)
							// line 286: return self.read().splitlines(True)
							πF.SetLineno(286)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreadlines.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 283: """
					πF.SetLineno(283)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Return lines of a single file as list of Unicode strings.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßreadlines); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 288: def close(self):
					πF.SetLineno(288)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp005 bool
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstdin, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 289: if self.source is not sys.stdin:
							πF.SetLineno(289)
						Label1:
							// line 290: self.source.close()
							πF.SetLineno(290)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FileInput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFileInput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 293: class FileOutput(Output):
			πF.SetLineno(293)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßOutput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FileOutput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 295: """
					πF.SetLineno(295)
					// line 295: """
					πF.SetLineno(295)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Output for single, simple file-like objects.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 299: mode = 'w'
					πF.SetLineno(299)
					if πE = πClass.SetItem(πF, ßmode.ToObject(), ßw.ToObject()); πE != nil {
						continue
					}
					// line 300: """The mode argument for `open()`."""
					πF.SetLineno(300)
					// line 305: def __init__(self, destination=None, destination_path=None,
					πF.SetLineno(305)
					πTemp002 = make([]πg.Param, 8)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "destination", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "destination_path", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "encoding", Def: πTemp003}
					πTemp002[4] = πg.Param{Name: "error_handler", Def: ßstrict.ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "autoclose", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "handle_io_errors", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[7] = πg.Param{Name: "mode", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdestination *πg.Object = πArgs[1]
						_ = µdestination
						var µdestination_path *πg.Object = πArgs[2]
						_ = µdestination_path
						var µencoding *πg.Object = πArgs[3]
						_ = µencoding
						var µerror_handler *πg.Object = πArgs[4]
						_ = µerror_handler
						var µautoclose *πg.Object = πArgs[5]
						_ = µautoclose
						var µhandle_io_errors *πg.Object = πArgs[6]
						_ = µhandle_io_errors
						var µmode *πg.Object = πArgs[7]
						_ = µmode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							case 13:
								goto Label13
							default:
								panic("unexpected function state")
							}
							// line 308: """
							πF.SetLineno(308)
							// line 324: Output.__init__(self, destination, destination_path,
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							πTemp001[1] = µdestination
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							πTemp001[2] = µdestination_path
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001[3] = µencoding
							if πE = πg.CheckLocal(πF, µerror_handler, "error_handler"); πE != nil {
								continue
							}
							πTemp001[4] = µerror_handler
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOutput); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 326: self.opened = True
							πF.SetLineno(326)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopened, πTemp003); πE != nil {
								continue
							}
							// line 327: self.autoclose = autoclose
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µautoclose, "autoclose"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µautoclose); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßautoclose, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmode != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 328: if mode is not None:
							πF.SetLineno(328)
						Label1:
							// line 329: self.mode = mode
							πF.SetLineno(329)
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmode, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 330: self._stderr = ErrorOutput()
							πF.SetLineno(330)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stderr, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µdestination == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp002 = µmode
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßmode.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, µmode, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label4:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 331: if destination is None:
							πF.SetLineno(331)
						Label3:
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µdestination_path); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 332: if destination_path:
							πF.SetLineno(332)
						Label7:
							// line 333: self.opened = False
							πF.SetLineno(333)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopened, πTemp003); πE != nil {
								continue
							}
							goto Label9
						Label8:
							// line 335: self.destination = sys.stdout
							πF.SetLineno(335)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstdout, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination, πTemp002); πE != nil {
								continue
							}
							goto Label9
						Label9:
							goto Label6
							// line 336: elif (# destination is file-type object -> check mode:
							πF.SetLineno(336)
						Label5:
							// line 339: print('Warning: Destination mode "%s" differs from specified '
							πF.SetLineno(339)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmode, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, µmode).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Warning: Destination mode \"%s\" differs from specified mode \"%s\"").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µdestination_path); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 342: if not destination_path:
							πF.SetLineno(342)
						Label10:
							// line 343: try:
							πF.SetLineno(343)
							πF.PushCheckpoint(13)
							// line 344: self.destination_path = self.destination.name
							πF.SetLineno(344)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination_path, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 345: except AttributeError:
							πF.SetLineno(345)
						Label14:
							// line 346: pass
							πF.SetLineno(346)
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 308: """
					πF.SetLineno(308)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        :Parameters:\n            - `destination`: either a file-like object (which is written\n              directly) or `None` (which implies `sys.stdout` if no\n              `destination_path` given).\n            - `destination_path`: a path to a file, which is opened and then\n              written.\n            - `encoding`: the text encoding of the output file.\n            - `error_handler`: the encoding error handler to use.\n            - `autoclose`: close automatically after write (except when\n              `sys.stdout` or `sys.stderr` is the destination).\n            - `handle_io_errors`: ignored, deprecated, will be removed.\n            - `mode`: how the file is to be opened (see standard function\n              `open`). The default is 'w', providing universal newline\n              support for text files.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 348: def open(self):
					πF.SetLineno(348)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("open", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µkwargs *πg.Object = πg.UnboundLocal
						_ = µkwargs
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
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
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp003, πE = πg.GE(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, ßb.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 350: if sys.version_info >= (3, 0) and 'b' not in self.mode:
							πF.SetLineno(350)
						Label2:
							// line 351: kwargs = {'encoding': self.encoding,
							πF.SetLineno(351)
							πTemp007 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßencoding.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror_handler, nil); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßerrors.ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							µkwargs = πTemp001
							goto Label4
						Label3:
							// line 354: kwargs = {}
							πF.SetLineno(354)
							πTemp007 = πg.NewDict()
							πTemp001 = πTemp007.ToObject()
							µkwargs = πTemp001
							goto Label4
						Label4:
							// line 355: try:
							πF.SetLineno(355)
							πF.PushCheckpoint(6)
							// line 356: self.destination = open(self.destination_path, self.mode, **kwargs)
							πF.SetLineno(356)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination_path, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßopen); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp001, πTemp008, nil, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdestination, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 357: except IOError as error:
							πF.SetLineno(357)
						Label7:
							µerror = πTemp009.ToObject()
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerror, ßerrno, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µerror, ßstrerror, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination_path, nil); πE != nil {
								continue
							}
							πTemp008[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOutputError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 358: raise OutputError(error.errno, error.strerror,
							πF.SetLineno(358)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							// line 360: self.opened = True
							πF.SetLineno(360)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopened, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßopen.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 362: def write(self, data):
					πF.SetLineno(362)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var µe *πg.Object = πg.UnboundLocal
						_ = µe
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.BaseException
						_ = πTemp013
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10:
								goto Label10
							case 11:
								goto Label11
							case 18:
								goto Label18
							default:
								panic("unexpected function state")
							}
							// line 363: """Encode `data`, write it to a single file, and return it.
							πF.SetLineno(363)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßopened, nil); πE != nil {
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
								goto Label1
							}
							goto Label2
							// line 368: if not self.opened:
							πF.SetLineno(368)
						Label1:
							// line 369: self.open()
							πF.SetLineno(369)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßopen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmode, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp006, ßb.ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp005, πE = πg.LT(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πTemp005
						Label4:
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcheck_encoding); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006 == πTemp005).ToObject()
							πTemp001 = πTemp002
						Label3:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 370: if ('b' not in self.mode and sys.version_info < (3, 0)
							πF.SetLineno(370)
						Label5:
							// line 373: data = self.encode(data)
							πF.SetLineno(373)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp009[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µdata = πTemp002
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp002, πE = πg.GE(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label7
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßlinesep, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label7:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 374: if sys.version_info >= (3, 0) and os.linesep != '\n':
							πF.SetLineno(374)
						Label8:
							// line 375: data = data.replace(b'\n', bytes(os.linesep, 'ascii')) # fix endings
							πF.SetLineno(375)
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewStr("\n").ToObject()
							πTemp010 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlinesep, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							πTemp010[1] = ßascii.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp009[1] = πTemp002
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µdata = πTemp002
							goto Label9
						Label9:
							goto Label6
						Label6:
							// line 377: try:
							πF.SetLineno(377)
							πF.PushCheckpoint(10)
							πF.PushCheckpoint(11)
							// line 378: self.destination.write(data)
							πF.SetLineno(378)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp009[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πF.PopCheckpoint()
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp005).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label13
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 379: except TypeError as e:
							πF.SetLineno(379)
						Label12:
							µe = πTemp011.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp002, πE = πg.GE(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label14
							}
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp009[0] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
								continue
							}
							πTemp009[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp001 = πTemp005
						Label14:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							goto Label16
							// line 380: if sys.version_info >= (3, 0) and isinstance(data, bytes):
							πF.SetLineno(380)
						Label15:
							// line 381: try:
							πF.SetLineno(381)
							πF.PushCheckpoint(18)
							// line 382: self.destination.buffer.write(data)
							πF.SetLineno(382)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp009[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßbuffer, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πF.PopCheckpoint()
							goto Label17
						Label18:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp013, πTemp012 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label19
							}
							πE = πF.Raise(πTemp013.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 383: except AttributeError:
							πF.SetLineno(383)
						Label19:
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp009[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcheck_encoding); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label20
							}
							goto Label21
							// line 384: if check_encoding(self.destination,
							πF.SetLineno(384)
						Label20:
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdestination_path, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label23
							}
							πTemp005 = ßdestination.ToObject()
						Label23:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßencoding, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp005, πTemp008, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Encoding of %s (%s) differs \n  from specified encoding (%s)").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							// line 386: raise ValueError('Encoding of %s (%s) differs \n'
							πF.SetLineno(386)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label22
						Label21:
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							// line 391: raise e
							πF.SetLineno(391)
							πE = πF.Raise(µe, nil, nil)
							continue
							goto Label22
						Label22:
							πF.RestoreExc(nil, nil)
							goto Label17
						Label17:
							goto Label16
						Label16:
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label10
							// line 392: except (UnicodeError, LookupError) as err:
							πF.SetLineno(392)
						Label13:
							µerr = πTemp011.ToObject()
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßencoding, nil); πE != nil {
								continue
							}
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp010[0] = µerr
							if πTemp006, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp002 = πg.NewTuple2(πTemp005, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unable to encode output data. output-encoding is: %s.\n(%s)").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							// line 393: raise UnicodeError(
							πF.SetLineno(393)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label10
						Label10:
							πTemp011, πTemp012 = πF.RestoreExc(nil, nil)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßautoclose, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label24
							}
							goto Label25
							// line 397: if self.autoclose:
							πF.SetLineno(397)
						Label24:
							// line 398: self.close()
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclose, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label25
						Label25:
							if πTemp011 != nil {
								πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							// line 399: return data
							πF.SetLineno(399)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πR = µdata
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 363: """Encode `data`, write it to a single file, and return it.
					πF.SetLineno(363)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Encode `data`, write it to a single file, and return it.\n\n        With Python 3 or binary output mode, `data` is returned unchanged,\n        except when specified encoding and output encoding differ.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßwrite); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 401: def close(self):
					πF.SetLineno(401)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
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
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 402: if self.destination not in (sys.stdout, sys.stderr):
							πF.SetLineno(402)
						Label1:
							// line 403: self.destination.close()
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 404: self.opened = False
							πF.SetLineno(404)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopened, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FileOutput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFileOutput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 407: class BinaryFileOutput(FileOutput):
			πF.SetLineno(407)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßFileOutput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BinaryFileOutput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 408: """
					πF.SetLineno(408)
					// line 408: """
					πF.SetLineno(408)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A version of docutils.io.FileOutput which writes to a binary file.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 413: mode = 'wb'
					πF.SetLineno(413)
					if πE = πClass.SetItem(πF, ßmode.ToObject(), ßwb.ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BinaryFileOutput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBinaryFileOutput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 416: class StringInput(Input):
			πF.SetLineno(416)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßInput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StringInput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 418: """
					πF.SetLineno(418)
					// line 418: """
					πF.SetLineno(418)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Direct string input.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 422: default_source_path = '<string>'
					πF.SetLineno(422)
					if πE = πClass.SetItem(πF, ßdefault_source_path.ToObject(), πg.NewStr("<string>").ToObject()); πE != nil {
						continue
					}
					// line 424: def read(self):
					πF.SetLineno(424)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 425: """Decode and return the source string."""
							πF.SetLineno(425)
							// line 426: return self.decode(self.source)
							πF.SetLineno(426)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdecode, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 425: """Decode and return the source string."""
					πF.SetLineno(425)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Decode and return the source string.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßread); πE != nil {
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
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("StringInput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringInput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 429: class StringOutput(Output):
			πF.SetLineno(429)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßOutput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StringOutput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 431: """
					πF.SetLineno(431)
					// line 431: """
					πF.SetLineno(431)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Direct string output.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 435: default_destination_path = '<string>'
					πF.SetLineno(435)
					if πE = πClass.SetItem(πF, ßdefault_destination_path.ToObject(), πg.NewStr("<string>").ToObject()); πE != nil {
						continue
					}
					// line 437: def write(self, data):
					πF.SetLineno(437)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
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
							// line 438: """Encode `data`, store it in `self.destination`, and return it."""
							πF.SetLineno(438)
							// line 439: self.destination = self.encode(data)
							πF.SetLineno(439)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßencode, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßdestination, πTemp002); πE != nil {
								continue
							}
							// line 440: return self.destination
							πF.SetLineno(440)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdestination, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 438: """Encode `data`, store it in `self.destination`, and return it."""
					πF.SetLineno(438)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Encode `data`, store it in `self.destination`, and return it.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßwrite); πE != nil {
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
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("StringOutput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringOutput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 443: class NullInput(Input):
			πF.SetLineno(443)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßInput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NullInput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 445: """
					πF.SetLineno(445)
					// line 445: """
					πF.SetLineno(445)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Degenerate input: read nothing.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 449: default_source_path = 'null input'
					πF.SetLineno(449)
					if πE = πClass.SetItem(πF, ßdefault_source_path.ToObject(), πg.NewStr("null input").ToObject()); πE != nil {
						continue
					}
					// line 451: def read(self):
					πF.SetLineno(451)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 452: """Return a null string."""
							πF.SetLineno(452)
							// line 453: return u''
							πF.SetLineno(453)
							πR = πg.NewUnicode("").ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 452: """Return a null string."""
					πF.SetLineno(452)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return a null string.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßread); πE != nil {
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
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("NullInput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNullInput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 456: class NullOutput(Output):
			πF.SetLineno(456)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßOutput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("NullOutput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 458: """
					πF.SetLineno(458)
					// line 458: """
					πF.SetLineno(458)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Degenerate output: write nothing.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 462: default_destination_path = 'null output'
					πF.SetLineno(462)
					if πE = πClass.SetItem(πF, ßdefault_destination_path.ToObject(), πg.NewStr("null output").ToObject()); πE != nil {
						continue
					}
					// line 464: def write(self, data):
					πF.SetLineno(464)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("write", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
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
							// line 465: """Do nothing ([don't even] send data to the bit bucket)."""
							πF.SetLineno(465)
							// line 466: pass
							πF.SetLineno(466)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 465: """Do nothing ([don't even] send data to the bit bucket)."""
					πF.SetLineno(465)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Do nothing ([don't even] send data to the bit bucket).").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßwrite); πE != nil {
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
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("NullOutput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNullOutput.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 469: class DocTreeInput(Input):
			πF.SetLineno(469)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßInput); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DocTreeInput", "/usr/lib/python2.7/site-packages/docutils/io.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 471: """
					πF.SetLineno(471)
					// line 471: """
					πF.SetLineno(471)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Adapter for document tree input.\n\n    The document tree must be passed in the ``source`` parameter.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 477: default_source_path = 'doctree input'
					πF.SetLineno(477)
					if πE = πClass.SetItem(πF, ßdefault_source_path.ToObject(), πg.NewStr("doctree input").ToObject()); πE != nil {
						continue
					}
					// line 479: def read(self):
					πF.SetLineno(479)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/io.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 480: """Return the document tree."""
							πF.SetLineno(480)
							// line 481: return self.source
							πF.SetLineno(481)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsource, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 480: """Return the document tree."""
					πF.SetLineno(480)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return the document tree.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßread); πE != nil {
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
			if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("DocTreeInput").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDocTreeInput.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("io", Code)
}
