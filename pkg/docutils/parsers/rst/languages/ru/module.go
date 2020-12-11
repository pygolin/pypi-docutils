package ru

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/languages/ru.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ßabbreviation := πg.InternStr("abbreviation")
		ßacronym := πg.InternStr("acronym")
		ßcode := πg.InternStr("code")
		ßcompound := πg.InternStr("compound")
		ßcontainer := πg.InternStr("container")
		ßdirectives := πg.InternStr("directives")
		ßemphasis := πg.InternStr("emphasis")
		ßfooter := πg.InternStr("footer")
		ßheader := πg.InternStr("header")
		ßindex := πg.InternStr("index")
		ßliteral := πg.InternStr("literal")
		ßmath := πg.InternStr("math")
		ßraw := πg.InternStr("raw")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrole := πg.InternStr("role")
		ßroles := πg.InternStr("roles")
		ßstrong := πg.InternStr("strong")
		ßsubscript := πg.InternStr("subscript")
		ßsuperscript := πg.InternStr("superscript")
		ßtable := πg.InternStr("table")
		ßtarget := πg.InternStr("target")
		ßtitle := πg.InternStr("title")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nRussian-language mappings for language-dependent features of\nreStructuredText.\n").ToObject()); πE != nil {
				continue
			}
			// line 16: __docformat__ = 'reStructuredText'
			πF.SetLineno(16)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 18: directives = {
			πF.SetLineno(18)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb1\xd0\xbb\xd0\xbe\xd0\xba-\xd1\x81\xd1\x82\xd1\x80\xd0\xbe\xd0\xba").ToObject(), πg.NewUnicode("line-block").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("meta").ToObject(), πg.NewUnicode("meta").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbc\xd0\xb0\xd1\x82\xd0\xb5\xd0\xbc\xd0\xb0\xd1\x82\xd0\xb8\xd0\xba\xd0\xb0").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbe\xd0\xb1\xd1\x80\xd0\xb0\xd0\xb1\xd0\xbe\xd1\x82\xd0\xb0\xd0\xbd\xd0\xbd\xd1\x8b\xd0\xb9-\xd0\xbb\xd0\xb8\xd1\x82\xd0\xb5\xd1\x80\xd0\xb0\xd0\xbb").ToObject(), πg.NewUnicode("parsed-literal").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd1\x8b\xd0\xb4\xd0\xb5\xd0\xbb\xd0\xb5\xd0\xbd\xd0\xbd\xd0\xb0\xd1\x8f-\xd1\x86\xd0\xb8\xd1\x82\xd0\xb0\xd1\x82\xd0\xb0").ToObject(), πg.NewUnicode("pull-quote").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xba\xd0\xbe\xd0\xb4").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("compound (translation required)").ToObject(), ßcompound.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xba\xd0\xbe\xd0\xbd\xd1\x82\xd0\xb5\xd0\xb9\xd0\xbd\xd0\xb5\xd1\x80").ToObject(), ßcontainer.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x82\xd0\xb0\xd0\xb1\xd0\xbb\xd0\xb8\xd1\x86\xd0\xb0").ToObject(), ßtable.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("csv-table (translation required)").ToObject(), πg.NewStr("csv-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("list-table (translation required)").ToObject(), πg.NewStr("list-table").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x8b\xd1\x80\xd0\xbe\xd0\xb9").ToObject(), πg.NewUnicode("raw").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb7\xd0\xb0\xd0\xbc\xd0\xb5\xd0\xbd\xd0\xb0").ToObject(), πg.NewUnicode("replace").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x82\xd0\xb5\xd1\x81\xd1\x82\xd0\xbe\xd0\xb2\xd0\xb0\xd1\x8f-\xd0\xb4\xd0\xb8\xd1\x80\xd0\xb5\xd0\xba\xd1\x82\xd0\xb8\xd0\xb2\xd0\xb0-restructuredtext").ToObject(), πg.NewUnicode("restructuredtext-test-directive").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x86\xd0\xb5\xd0\xbb\xd0\xb5\xd0\xb2\xd1\x8b\xd0\xb5-\xd1\x81\xd0\xbd\xd0\xbe\xd1\x81\xd0\xba\xd0\xb8").ToObject(), πg.NewUnicode("target-notes").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("unicode").ToObject(), πg.NewUnicode("unicode").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb4\xd0\xb0\xd1\x82\xd0\xb0").ToObject(), πg.NewUnicode("date").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb1\xd0\xbe\xd0\xba\xd0\xbe\xd0\xb2\xd0\xb0\xd1\x8f-\xd0\xbf\xd0\xbe\xd0\xbb\xd0\xbe\xd1\x81\xd0\xb0").ToObject(), πg.NewUnicode("sidebar").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd0\xb0\xd0\xb6\xd0\xbd\xd0\xbe").ToObject(), πg.NewUnicode("important").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd0\xba\xd0\xbb\xd1\x8e\xd1\x87\xd0\xb0\xd1\x82\xd1\x8c").ToObject(), πg.NewUnicode("include").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd0\xbd\xd0\xb8\xd0\xbc\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("attention").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd1\x8b\xd0\xb4\xd0\xb5\xd0\xbb\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("highlights").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb7\xd0\xb0\xd0\xbc\xd0\xb5\xd1\x87\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("admonition").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb8\xd0\xb7\xd0\xbe\xd0\xb1\xd1\x80\xd0\xb0\xd0\xb6\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("image").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xba\xd0\xbb\xd0\xb0\xd1\x81\xd1\x81").ToObject(), πg.NewUnicode("class").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x80\xd0\xbe\xd0\xbb\xd1\x8c").ToObject(), ßrole.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("default-role (translation required)").ToObject(), πg.NewStr("default-role").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x82\xd0\xb8\xd1\x82\xd1\x83\xd0\xbb").ToObject(), ßtitle.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbd\xd0\xbe\xd0\xbc\xd0\xb5\xd1\x80-\xd1\x80\xd0\xb0\xd0\xb7\xd0\xb4\xd0\xb5\xd0\xbb\xd0\xb0").ToObject(), πg.NewUnicode("sectnum").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbd\xd1\x83\xd0\xbc\xd0\xb5\xd1\x80\xd0\xb0\xd1\x86\xd0\xb8\xd1\x8f-\xd1\x80\xd0\xb0\xd0\xb7\xd0\xb4\xd0\xb5\xd0\xbb\xd0\xbe\xd0\xb2").ToObject(), πg.NewUnicode("sectnum").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbe\xd0\xbf\xd0\xb0\xd1\x81\xd0\xbd\xd0\xbe").ToObject(), πg.NewUnicode("danger").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbe\xd1\x81\xd1\x82\xd0\xbe\xd1\x80\xd0\xbe\xd0\xb6\xd0\xbd\xd0\xbe").ToObject(), πg.NewUnicode("caution").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbe\xd1\x88\xd0\xb8\xd0\xb1\xd0\xba\xd0\xb0").ToObject(), πg.NewUnicode("error").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbf\xd0\xbe\xd0\xb4\xd1\x81\xd0\xba\xd0\xb0\xd0\xb7\xd0\xba\xd0\xb0").ToObject(), πg.NewUnicode("tip").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbf\xd1\x80\xd0\xb5\xd0\xb4\xd1\x83\xd0\xbf\xd1\x80\xd0\xb5\xd0\xb6\xd0\xb4\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("warning").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbf\xd1\x80\xd0\xb8\xd0\xbc\xd0\xb5\xd1\x87\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("note").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x80\xd0\xb8\xd1\x81\xd1\x83\xd0\xbd\xd0\xbe\xd0\xba").ToObject(), πg.NewUnicode("figure").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x80\xd1\x83\xd0\xb1\xd1\x80\xd0\xb8\xd0\xba\xd0\xb0").ToObject(), πg.NewUnicode("rubric").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd0\xbe\xd0\xb2\xd0\xb5\xd1\x82").ToObject(), πg.NewUnicode("hint").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd0\xbe\xd0\xb4\xd0\xb5\xd1\x80\xd0\xb6\xd0\xb0\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), πg.NewUnicode("contents").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x82\xd0\xb5\xd0\xbc\xd0\xb0").ToObject(), πg.NewUnicode("topic").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x8d\xd0\xbf\xd0\xb8\xd0\xb3\xd1\x80\xd0\xb0\xd1\x84").ToObject(), πg.NewUnicode("epigraph").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("header (translation required)").ToObject(), ßheader.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("footer (translation required)").ToObject(), ßfooter.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 63: """Russian name to registered (in directives/__init__.py) directive name
			πF.SetLineno(63)
			// line 66: roles = {
			πF.SetLineno(66)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb0\xd0\xba\xd1\x80\xd0\xbe\xd0\xbd\xd0\xb8\xd0\xbc").ToObject(), ßacronym.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xba\xd0\xbe\xd0\xb4").ToObject(), ßcode.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb0\xd0\xbd\xd0\xbe\xd0\xbd\xd0\xb8\xd0\xbc\xd0\xbd\xd0\xb0\xd1\x8f-\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0").ToObject(), πg.NewStr("anonymous-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb1\xd1\x83\xd0\xba\xd0\xb2\xd0\xb0\xd0\xbb\xd1\x8c\xd0\xbd\xd0\xbe").ToObject(), ßliteral.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbc\xd0\xb0\xd1\x82\xd0\xb5\xd0\xbc\xd0\xb0\xd1\x82\xd0\xb8\xd0\xba\xd0\xb0").ToObject(), ßmath.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd0\xb5\xd1\x80\xd1\x85\xd0\xbd\xd0\xb8\xd0\xb9-\xd0\xb8\xd0\xbd\xd0\xb4\xd0\xb5\xd0\xba\xd1\x81").ToObject(), ßsuperscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb2\xd1\x8b\xd0\xb4\xd0\xb5\xd0\xbb\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), ßemphasis.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb8\xd0\xbc\xd0\xb5\xd0\xbd\xd0\xbe\xd0\xb2\xd0\xb0\xd0\xbd\xd0\xbd\xd0\xb0\xd1\x8f-\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0").ToObject(), πg.NewStr("named-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xb8\xd0\xbd\xd0\xb4\xd0\xb5\xd0\xba\xd1\x81").ToObject(), ßindex.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd0\xbd\xd0\xb8\xd0\xb6\xd0\xbd\xd0\xb8\xd0\xb9-\xd0\xb8\xd0\xbd\xd0\xb4\xd0\xb5\xd0\xba\xd1\x81").ToObject(), ßsubscript.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd0\xb8\xd0\xbb\xd1\x8c\xd0\xbd\xd0\xbe\xd0\xb5-\xd0\xb2\xd1\x8b\xd0\xb4\xd0\xb5\xd0\xbb\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), ßstrong.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd0\xbe\xd0\xba\xd1\x80\xd0\xb0\xd1\x89\xd0\xb5\xd0\xbd\xd0\xb8\xd0\xb5").ToObject(), ßabbreviation.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0-\xd0\xb7\xd0\xb0\xd0\xbc\xd0\xb5\xd0\xbd\xd0\xb0").ToObject(), πg.NewStr("substitution-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0-\xd0\xbd\xd0\xb0-pep").ToObject(), πg.NewStr("pep-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0-\xd0\xbd\xd0\xb0-rfc").ToObject(), πg.NewStr("rfc-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0-\xd0\xbd\xd0\xb0-uri").ToObject(), πg.NewStr("uri-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0-\xd0\xbd\xd0\xb0-\xd0\xb7\xd0\xb0\xd0\xb3\xd0\xbb\xd0\xb0\xd0\xb2\xd0\xb8\xd0\xb5").ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0-\xd0\xbd\xd0\xb0-\xd1\x81\xd0\xbd\xd0\xbe\xd1\x81\xd0\xba\xd1\x83").ToObject(), πg.NewStr("footnote-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x86\xd0\xb8\xd1\x82\xd0\xb0\xd1\x82\xd0\xbd\xd0\xb0\xd1\x8f-\xd1\x81\xd1\x81\xd1\x8b\xd0\xbb\xd0\xba\xd0\xb0").ToObject(), πg.NewStr("citation-reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x86\xd0\xb5\xd0\xbb\xd1\x8c").ToObject(), ßtarget.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewUnicode("\xd1\x81\xd1\x8b\xd1\x80\xd0\xbe\xd0\xb9").ToObject(), ßraw.ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 88: """Mapping of Russian role names to canonical role names for interpreted text.
			πF.SetLineno(88)
		}
		return nil, πE
	})
	πg.RegisterModule("ru", Code)
}
