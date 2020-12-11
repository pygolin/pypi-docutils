package __main__

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/locale"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rstpep2html.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
		ßpep := πg.InternStr("pep")
		ßpep_html := πg.InternStr("pep_html")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nA minimal front end to the Docutils Publisher, producing HTML from PEP\n(Python Enhancement Proposal) documents.\n").ToObject()); πE != nil {
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
			// line 18: from docutils.core import publish_cmdline, default_description
			πF.SetLineno(18)
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
			// line 21: description = ('Generates (X)HTML from reStructuredText-format PEP files.  '
			πF.SetLineno(21)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdefault_description); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πg.NewStr("Generates (X)HTML from reStructuredText-format PEP files.  ").ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdescription.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 24: publish_cmdline(reader_name='pep', writer_name='pep_html',
			πF.SetLineno(24)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdescription); πE != nil {
				continue
			}
			πTemp006 = πg.KWArgs{
				{"reader_name", ßpep.ToObject()},
				{"writer_name", ßpep_html.ToObject()},
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
