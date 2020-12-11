package fa

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/fa.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßdirectives := πg.InternStr("directives")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßroles := πg.InternStr("roles")
		var πTemp001 *πg.Dict
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 11: """
			πF.SetLineno(11)
			// line 11: """
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nPersian-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
				continue
			}
			// line 16: __docformat__ = 'reStructuredText'
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 19: directives = {
			πF.SetLineno(19)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd9\x88\xd8\xac\xd9\x87").ToObject(), πg.NewUnicode("attention").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa7\xd8\xad\xd8\xaa\xdb\x8c\xd8\xa7\xd8\xb7").ToObject(), πg.NewUnicode("caution").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\xa9\xd8\xaf").ToObject(), πg.NewUnicode("code").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa8\xd9\x84\xd9\x88\xda\xa9-\xda\xa9\xd8\xaf").ToObject(), πg.NewUnicode("code").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\xa9\xd8\xaf-\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9").ToObject(), πg.NewUnicode("code").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xae\xd8\xb7\xd8\xb1").ToObject(), πg.NewUnicode("danger").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xae\xd8\xb7\xd8\xa7").ToObject(), πg.NewUnicode("error").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb1\xd8\xa7\xd9\x87\xd9\x86\xd9\x85\xd8\xa7").ToObject(), πg.NewUnicode("hint").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x87\xd9\x85").ToObject(), πg.NewUnicode("important").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xdb\x8c\xd8\xa7\xd8\xaf\xd8\xaf\xd8\xa7\xd8\xb4\xd8\xaa").ToObject(), πg.NewUnicode("note").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xda\xa9\xd8\xaa\xd9\x87").ToObject(), πg.NewUnicode("tip").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa7\xd8\xae\xd8\xb7\xd8\xa7\xd8\xb1").ToObject(), πg.NewUnicode("warning").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xb0\xda\xa9\xd8\xb1").ToObject(), πg.NewUnicode("admonition").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xd9\x88\xd8\xa7\xd8\xb1-\xda\xa9\xd9\x86\xd8\xa7\xd8\xb1\xdb\x8c").ToObject(), πg.NewUnicode("sidebar").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x88\xd8\xb6\xd9\x88\xd8\xb9").ToObject(), πg.NewUnicode("topic").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa8\xd9\x84\xd9\x88\xda\xa9-\xd8\xae\xd8\xb7").ToObject(), πg.NewUnicode("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd9\x84\xd9\x81\xd8\xb8-\xd9\xbe\xd8\xb1\xd8\xaf\xd8\xa7\xd8\xb2\xd8\xb4-\xd8\xb4\xd8\xaf\xd9\x87").ToObject(), πg.NewUnicode("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb3\xd8\xb1-\xd9\x81\xd8\xb5\xd9\x84").ToObject(), πg.NewUnicode("rubric").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\xa9\xd8\xaa\xdb\x8c\xd8\xa8\xd9\x87").ToObject(), πg.NewUnicode("epigraph").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xda\xa9\xd8\xa7\xd8\xaa-\xd8\xa8\xd8\xb1\xd8\xac\xd8\xb3\xd8\xaa\xd9\x87").ToObject(), πg.NewUnicode("highlights").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xd9\x82\xd9\x84-\xd9\x82\xd9\x88\xd9\x84").ToObject(), πg.NewUnicode("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xb1\xda\xa9\xdb\x8c\xd8\xa8").ToObject(), πg.NewUnicode("compound").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb8\xd8\xb1\xd9\x81").ToObject(), πg.NewUnicode("container").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xac\xd8\xaf\xd9\x88\xd9\x84").ToObject(), πg.NewUnicode("table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xac\xd8\xaf\xd9\x88\xd9\x84-csv").ToObject(), πg.NewUnicode("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xac\xd8\xaf\xd9\x88\xd9\x84-\xd9\x84\xdb\x8c\xd8\xb3\xd8\xaa").ToObject(), πg.NewUnicode("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd8\xaa\xd8\xa7").ToObject(), πg.NewUnicode("meta").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb1\xdb\x8c\xd8\xa7\xd8\xb6\xdb\x8c").ToObject(), πg.NewUnicode("math").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xb5\xd9\x88\xdb\x8c\xd8\xb1").ToObject(), πg.NewUnicode("image").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb4\xda\xa9\xd9\x84").ToObject(), πg.NewUnicode("figure").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb4\xd8\xa7\xd9\x85\xd9\x84").ToObject(), πg.NewUnicode("include").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xae\xd8\xa7\xd9\x85").ToObject(), πg.NewUnicode("raw").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xac\xd8\xa7\xdb\x8c\xda\xaf\xd8\xb2\xdb\x8c\xd9\x86").ToObject(), πg.NewUnicode("replace").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xdb\x8c\xd9\x88\xd9\x86\xdb\x8c\xda\xa9\xd8\xaf").ToObject(), πg.NewUnicode("unicode").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xa7\xd8\xb1\xdb\x8c\xd8\xae").ToObject(), πg.NewUnicode("date").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\xa9\xd9\x84\xd8\xa7\xd8\xb3").ToObject(), πg.NewUnicode("class").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x82\xd8\xa7\xd9\x86\xd9\x88\xd9\x86").ToObject(), πg.NewUnicode("role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x82\xd8\xa7\xd9\x86\xd9\x88\xd9\x86-\xd9\xbe\xdb\x8c\xd8\xb4\xe2\x80\x8c\xd9\x81\xd8\xb1\xd8\xb6").ToObject(), πg.NewUnicode("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb9\xd9\x86\xd9\x88\xd8\xa7\xd9\x86").ToObject(), πg.NewUnicode("title").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd8\xad\xd8\xaa\xd9\x88\xd8\xa7").ToObject(), πg.NewUnicode("contents").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb4\xd9\x85\xd8\xa7\xd8\xb1\xd9\x87-\xd9\x81\xd8\xb5\xd9\x84").ToObject(), πg.NewUnicode("sectnum").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb4\xd9\x85\xd8\xa7\xd8\xb1\xd9\x87\xe2\x80\x8c\xda\xaf\xd8\xb0\xd8\xa7\xd8\xb1\xdb\x8c-\xd9\x81\xd8\xb5\xd9\x84").ToObject(), πg.NewUnicode("sectnum").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb3\xd8\xb1\xd8\xa2\xdb\x8c\xd9\x86\xd8\xaf").ToObject(), πg.NewUnicode("header").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\xbe\xd8\xa7\xd8\xb5\xd9\x81\xd8\xad\xd9\x87").ToObject(), πg.NewUnicode("footer").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xdb\x8c\xd8\xa7\xd8\xaf\xd8\xaf\xd8\xa7\xd8\xb4\xd8\xaa-\xd9\x87\xd8\xaf\xd9\x81").ToObject(), πg.NewUnicode("target-notes").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 73: """Persian name to registered (in directives/__init__.py) directive name
			πF.SetLineno(73)
			// line 76: roles = {
			πF.SetLineno(76)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd8\xae\xd9\x81\xd9\x81").ToObject(), πg.NewUnicode("abbreviation").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb3\xd8\xb1\xd9\x86\xd8\xa7\xd9\x85").ToObject(), πg.NewUnicode("acronym").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xda\xa9\xd8\xaf").ToObject(), πg.NewUnicode("code").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb4\xd8\xa7\xd8\xae\xd8\xb5").ToObject(), πg.NewUnicode("index").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb2\xdb\x8c\xd8\xb1\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3").ToObject(), πg.NewUnicode("subscript").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xa8\xd8\xa7\xd9\x84\xd8\xa7\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3").ToObject(), πg.NewUnicode("superscript").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb9\xd9\x86\xd9\x88\xd8\xa7\xd9\x86").ToObject(), πg.NewUnicode("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x86\xdb\x8c\xd8\xb1\xd9\x88").ToObject(), πg.NewUnicode("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("rfc-reference (translation required)").ToObject(), πg.NewUnicode("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xaa\xd8\xa7\xda\xa9\xdb\x8c\xd8\xaf").ToObject(), πg.NewUnicode("emphasis").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x82\xd9\x88\xdb\x8c").ToObject(), πg.NewUnicode("strong").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x84\xd9\x81\xd8\xb8\xdb\x8c").ToObject(), πg.NewUnicode("literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xb1\xdb\x8c\xd8\xa7\xd8\xb6\xdb\x8c").ToObject(), πg.NewUnicode("math").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9-\xd9\x86\xd8\xa7\xd9\x85\xe2\x80\x8c\xda\xaf\xd8\xb0\xd8\xa7\xd8\xb1\xdb\x8c").ToObject(), πg.NewUnicode("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9-\xd9\x86\xd8\xa7\xd8\xb4\xd9\x86\xd8\xa7\xd8\xb3").ToObject(), πg.NewUnicode("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9-\xd9\xbe\xd8\xa7\xd9\x86\xd9\x88\xdb\x8c\xd8\xb3").ToObject(), πg.NewUnicode("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9-\xd9\x86\xd9\x82\xd9\x84\xe2\x80\x8c\xd9\x81\xd9\x88\xd9\x84").ToObject(), πg.NewUnicode("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9-\xd8\xac\xd8\xa7\xdb\x8c\xda\xaf\xd8\xb2\xdb\x8c\xd9\x86\xdb\x8c").ToObject(), πg.NewUnicode("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x87\xd8\xaf\xd9\x81").ToObject(), πg.NewUnicode("target").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd9\x85\xd9\x86\xd8\xa8\xd8\xb9-uri").ToObject(), πg.NewUnicode("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("uri").ToObject(), πg.NewUnicode("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("url").ToObject(), πg.NewUnicode("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd8\xae\xd8\xa7\xd9\x85").ToObject(), πg.NewUnicode("raw").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 101: """Mapping of Persian role names to canonical role names for interpreted text.
			πF.SetLineno(101)
		}
		return nil, πE
	})
	πg.RegisterModule("fa", Code)
}
