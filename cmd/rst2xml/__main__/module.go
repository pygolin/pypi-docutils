package __main__

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/locale"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2xml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßLC_ALL := πg.InternStr("LC_ALL")
		ß__doc__ := πg.InternStr("__doc__")
		ßdefault_description := πg.InternStr("default_description")
		ßdescription := πg.InternStr("description")
		ßlocale := πg.InternStr("locale")
		ßpublish_cmdline := πg.InternStr("publish_cmdline")
		ßsetlocale := πg.InternStr("setlocale")
		ßxml := πg.InternStr("xml")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nA minimal front end to the Docutils Publisher, producing Docutils XML.\n").ToObject()); πE != nil {
				continue
			}
			// line 11: try:
			πF.SetLineno(11)
			πF.PushCheckpoint(2)
			// line 12: import locale
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "locale"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlocale.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: locale.setlocale(locale.LC_ALL, '')
			πF.SetLineno(13)
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
			// line 14: except:
			πF.SetLineno(14)
		Label3:
			// line 15: pass
			πF.SetLineno(15)
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 17: from docutils.core import publish_cmdline, default_description
			πF.SetLineno(17)
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
			// line 20: description = ('Generates Docutils-native XML from standalone '
			πF.SetLineno(20)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewStr("Generates Docutils-native XML from standalone reStructuredText sources.  ").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdescription.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: publish_cmdline(writer_name='xml', description=description)
			πF.SetLineno(23)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdescription); πE != nil {
				continue
			}
			πTemp006 = πg.KWArgs{
				{"writer_name", ßxml.ToObject()},
				{"description", πTemp001},
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
