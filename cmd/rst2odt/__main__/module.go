package __main__

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/locale"
	_ "github.com/pygolin/stdlib/pkg/sys"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2odt.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßLC_ALL := πg.InternStr("LC_ALL")
		ßReader := πg.InternStr("Reader")
		ßWriter := πg.InternStr("Writer")
		ß__doc__ := πg.InternStr("__doc__")
		ßdefault_description := πg.InternStr("default_description")
		ßdescription := πg.InternStr("description")
		ßlocale := πg.InternStr("locale")
		ßoutput := πg.InternStr("output")
		ßpublish_cmdline_to_binary := πg.InternStr("publish_cmdline_to_binary")
		ßreader := πg.InternStr("reader")
		ßsetlocale := πg.InternStr("setlocale")
		ßsys := πg.InternStr("sys")
		ßwriter := πg.InternStr("writer")
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
		var πTemp007 πg.KWArgs
		_ = πTemp007
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2:
				goto Label2
			default:
				panic("unexpected function state")
			}
			// line 7: """
			πF.SetLineno(7)
			// line 7: """
			πF.SetLineno(7)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nA front end to the Docutils Publisher, producing OpenOffice documents.\n").ToObject()); πE != nil {
				continue
			}
			// line 11: import sys
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: try:
			πF.SetLineno(12)
			πF.PushCheckpoint(2)
			// line 13: import locale
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "locale"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlocale.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: locale.setlocale(locale.LC_ALL, '')
			πF.SetLineno(14)
			πTemp002 = πF.MakeArgs(2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßLC_ALL, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			πTemp002[1] = ß.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsetlocale, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			goto Label3
			// line 15: except:
			πF.SetLineno(15)
		Label3:
			// line 16: pass
			πF.SetLineno(16)
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 18: from docutils.core import publish_cmdline_to_binary, default_description
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.core"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpublish_cmdline_to_binary); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpublish_cmdline_to_binary.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdefault_description); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefault_description.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 19: from docutils.writers.odf_odt import Writer, Reader
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers.odf_odt"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßWriter); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßReader); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReader.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: description = ('Generates OpenDocument/OpenOffice/ODF documents from '
			πF.SetLineno(22)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewStr("Generates OpenDocument/OpenOffice/ODF documents from standalone reStructuredText sources.  ").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdescription.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 26: writer = Writer()
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßWriter); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwriter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 27: reader = Reader()
			πF.SetLineno(27)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßReader); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreader.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 28: output = publish_cmdline_to_binary(reader=reader, writer=writer,
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßreader); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßwriter); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßdescription); πE != nil {
				continue
			}
			πTemp007 = πg.KWArgs{
				{"reader", πTemp001},
				{"writer", πTemp003},
				{"description", πTemp006},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßpublish_cmdline_to_binary); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, πTemp007); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßoutput.ToObject(), πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
