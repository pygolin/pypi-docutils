package __main__

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/locale"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2html5.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßError := πg.InternStr("Error")
		ßLC_ALL := πg.InternStr("LC_ALL")
		ß__doc__ := πg.InternStr("__doc__")
		ßdefault_description := πg.InternStr("default_description")
		ßdescription := πg.InternStr("description")
		ßhtml5 := πg.InternStr("html5")
		ßlocale := πg.InternStr("locale")
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
		var πTemp006 bool
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
			// line 16: """
			πF.SetLineno(16)
			// line 16: """
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nA minimal front end to the Docutils Publisher, producing HTML 5 documents.\n\nThe output also conforms to XHTML 1.0 transitional\n(except for the doctype declaration).\n").ToObject()); πE != nil {
				continue
			}
			// line 23: try:
			πF.SetLineno(23)
			πF.PushCheckpoint(2)
			// line 24: import locale # module missing in Jython
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "locale"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlocale.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: locale.setlocale(locale.LC_ALL, '')
			πF.SetLineno(25)
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
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßError, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label3
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 26: except locale.Error:
			πF.SetLineno(26)
		Label3:
			// line 27: pass
			πF.SetLineno(27)
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 29: from docutils.core import publish_cmdline, default_description
			πF.SetLineno(29)
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
			// line 31: description = (u'Generates HTML 5 documents from standalone '
			πF.SetLineno(31)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewUnicode("Generates HTML 5 documents from standalone reStructuredText sources ").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdescription.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: publish_cmdline(writer_name='html5', description=description)
			πF.SetLineno(35)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdescription); πE != nil {
				continue
			}
			πTemp007 = πg.KWArgs{
				{"writer_name", ßhtml5.ToObject()},
				{"description", πTemp001},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßpublish_cmdline); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, nil, πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
