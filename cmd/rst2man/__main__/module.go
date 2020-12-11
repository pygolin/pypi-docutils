package __main__

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/locale"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2man.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßLC_ALL := πg.InternStr("LC_ALL")
		ßWriter := πg.InternStr("Writer")
		ß__doc__ := πg.InternStr("__doc__")
		ßdefault_description := πg.InternStr("default_description")
		ßdescription := πg.InternStr("description")
		ßlocale := πg.InternStr("locale")
		ßmanpage := πg.InternStr("manpage")
		ßpublish_cmdline := πg.InternStr("publish_cmdline")
		ßsetlocale := πg.InternStr("setlocale")
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
		var πTemp006 πg.KWArgs
		_ = πTemp006
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nman.py\n======\n\nThis module provides a simple command line interface that uses the\nman page writer to output from ReStructuredText source.\n").ToObject()); πE != nil {
				continue
			}
			// line 15: import locale
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "locale"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlocale.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: try:
			πF.SetLineno(16)
			πF.PushCheckpoint(2)
			// line 17: locale.setlocale(locale.LC_ALL, '')
			πF.SetLineno(17)
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
			// line 18: except:
			πF.SetLineno(18)
		Label3:
			// line 19: pass
			πF.SetLineno(19)
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 21: from docutils.core import publish_cmdline, default_description
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.core"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßpublish_cmdline); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßpublish_cmdline.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdefault_description); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefault_description.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: from docutils.writers import manpage
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.writers.manpage"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßmanpage.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 24: description = ("Generates plain unix manual documents.  " + default_description)
			πF.SetLineno(24)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewStr("Generates plain unix manual documents.  ").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdescription.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 26: publish_cmdline(writer=manpage.Writer(), description=description)
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßmanpage); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßWriter, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdescription); πE != nil {
				continue
			}
			πTemp006 = πg.KWArgs{
				{"writer", πTemp001},
				{"description", πTemp003},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßpublish_cmdline); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
