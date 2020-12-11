package tex2mathml_extern

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/subprocess"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/math/tex2mathml_extern.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßError := πg.InternStr("Error")
		ßNone := πg.InternStr("None")
		ßPIPE := πg.InternStr("PIPE")
		ßPopen := πg.InternStr("Popen")
		ßSyntaxError := πg.InternStr("SyntaxError")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßblahtexml := πg.InternStr("blahtexml")
		ßclose := πg.InternStr("close")
		ßdecode := πg.InternStr("decode")
		ßdocument_template := πg.InternStr("document_template")
		ßencode := πg.InternStr("encode")
		ßerror := πg.InternStr("error")
		ßexample := πg.InternStr("example")
		ßfind := πg.InternStr("find")
		ßjoin := πg.InternStr("join")
		ßlatexml := πg.InternStr("latexml")
		ßlatexmlpost := πg.InternStr("latexmlpost")
		ßmoderate := πg.InternStr("moderate")
		ßprint := πg.InternStr("print")
		ßraw := πg.InternStr("raw")
		ßread := πg.InternStr("read")
		ßsplitlines := πg.InternStr("splitlines")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstdin := πg.InternStr("stdin")
		ßstdout := πg.InternStr("stdout")
		ßsubprocess := πg.InternStr("subprocess")
		ßttm := πg.InternStr("ttm")
		ßutf8 := πg.InternStr("utf8")
		ßwrite := πg.InternStr("write")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		var πTemp009 []*πg.Object
		_ = πTemp009
		var πTemp010 []*πg.Object
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 18: from __future__ import print_function
			πF.SetLineno(18)
			// line 19: import subprocess
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "subprocess"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsubprocess.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: document_template = r"""\documentclass{article}
			πF.SetLineno(21)
			// line 21: document_template = r"""\documentclass{article}
			πF.SetLineno(21)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\\documentclass{article}\n\\usepackage{amsmath}\n\\begin{document}\n%s\n\\end{document}\n").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdocument_template.ToObject(), πg.NewStr("\\documentclass{article}\n\\usepackage{amsmath}\n\\begin{document}\n%s\n\\end{document}\n").ToObject()); πE != nil {
				continue
			}
			// line 28: def latexml(math_code, reporter=None):
			πF.SetLineno(28)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "math_code", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "reporter", Def: πTemp004}
			πTemp001 = πg.NewFunction(πg.NewCode("latexml", "/usr/lib/python2.7/site-packages/docutils/utils/math/tex2mathml_extern.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmath_code *πg.Object = πArgs[0]
				_ = µmath_code
				var µreporter *πg.Object = πArgs[1]
				_ = µreporter
				var µp *πg.Object = πg.UnboundLocal
				_ = µp
				var µlatexml_code *πg.Object = πg.UnboundLocal
				_ = µlatexml_code
				var µlatexml_err *πg.Object = πg.UnboundLocal
				_ = µlatexml_err
				var µpost_p *πg.Object = πg.UnboundLocal
				_ = µpost_p
				var µresult *πg.Object = πg.UnboundLocal
				_ = µresult
				var µpost_p_err *πg.Object = πg.UnboundLocal
				_ = µpost_p_err
				var µstart *πg.Object = πg.UnboundLocal
				_ = µstart
				var µend *πg.Object = πg.UnboundLocal
				_ = µend
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
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
					// line 29: """Convert LaTeX math code to MathML with LaTeXML_
					πF.SetLineno(29)
					// line 33: p = subprocess.Popen(['latexml',
					πF.SetLineno(33)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = make([]*πg.Object, 3)
					πTemp002[0] = ßlatexml.ToObject()
					πTemp002[1] = πg.NewStr("-").ToObject()
					πTemp002[2] = πg.NewStr("--inputencoding=utf8").ToObject()
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"stdin", πTemp004},
						{"stdout", πTemp005},
						{"stderr", πTemp006},
						{"close_fds", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					// line 42: p.stdin.write((document_template % math_code).encode('utf8'))
					πF.SetLineno(42)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ßutf8.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdocument_template); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πTemp004, µmath_code); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 43: p.stdin.close()
					πF.SetLineno(43)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 44: latexml_code = p.stdout.read()
					πF.SetLineno(44)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlatexml_code = πTemp003
					// line 45: latexml_err = p.stderr.read().decode('utf8')
					πF.SetLineno(45)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlatexml_err = πTemp003
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					πTemp003 = µreporter
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßError.ToObject()
					if πE = πg.CheckLocal(πF, µlatexml_err, "latexml_err"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µlatexml_err, ßfind, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.GE(πF, πTemp010, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µlatexml_code, "latexml_code"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, µlatexml_code); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp011).ToObject()
					πTemp004 = πTemp005
				Label2:
					πTemp003 = πTemp004
				Label1:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					goto Label4
					// line 46: if reporter and (latexml_err.find('Error') >= 0 or not latexml_code):
					πF.SetLineno(46)
				Label3:
					// line 47: reporter.error(latexml_err)
					πF.SetLineno(47)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlatexml_err, "latexml_err"); πE != nil {
						continue
					}
					πTemp001[0] = µlatexml_err
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µreporter, ßerror, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label4
				Label4:
					// line 49: post_p = subprocess.Popen(['latexmlpost',
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = make([]*πg.Object, 5)
					πTemp002[0] = ßlatexmlpost.ToObject()
					πTemp002[1] = πg.NewStr("-").ToObject()
					πTemp002[2] = πg.NewStr("--nonumbersections").ToObject()
					πTemp002[3] = πg.NewStr("--format=xhtml").ToObject()
					πTemp002[4] = πg.NewStr("--").ToObject()
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"stdin", πTemp004},
						{"stdout", πTemp005},
						{"stderr", πTemp006},
						{"close_fds", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpost_p = πTemp003
					// line 60: post_p.stdin.write(latexml_code)
					πF.SetLineno(60)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlatexml_code, "latexml_code"); πE != nil {
						continue
					}
					πTemp001[0] = µlatexml_code
					if πE = πg.CheckLocal(πF, µpost_p, "post_p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpost_p, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 61: post_p.stdin.close()
					πF.SetLineno(61)
					if πE = πg.CheckLocal(πF, µpost_p, "post_p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpost_p, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 62: result = post_p.stdout.read().decode('utf8')
					πF.SetLineno(62)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µpost_p, "post_p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpost_p, ßstdout, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp003
					// line 63: post_p_err = post_p.stderr.read().decode('utf8')
					πF.SetLineno(63)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µpost_p, "post_p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpost_p, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpost_p_err = πTemp003
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					πTemp003 = µreporter
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label5
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßError.ToObject()
					if πE = πg.CheckLocal(πF, µpost_p_err, "post_p_err"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µpost_p_err, ßfind, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.GE(πF, πTemp010, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp011).ToObject()
					πTemp004 = πTemp005
				Label6:
					πTemp003 = πTemp004
				Label5:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 64: if reporter and (post_p_err.find('Error') >= 0 or not result):
					πF.SetLineno(64)
				Label7:
					// line 65: reporter.error(post_p_err)
					πF.SetLineno(65)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpost_p_err, "post_p_err"); πE != nil {
						continue
					}
					πTemp001[0] = µpost_p_err
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µreporter, ßerror, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label8
				Label8:
					// line 68: start, end = result.find('<math'), result.find('</math>')+7
					πF.SetLineno(68)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<math").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("</math>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Add(πF, πTemp010, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µstart = πTemp004
					µend = πTemp005
					// line 69: result = result[start:end]
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µresult = πTemp004
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µresult, πg.NewStr("class=\"ltx_ERROR").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					goto Label10
					// line 70: if 'class="ltx_ERROR' in result:
					πF.SetLineno(70)
				Label9:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					πTemp001[0] = µresult
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 71: raise SyntaxError(result)
					πF.SetLineno(71)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label10
				Label10:
					// line 72: return result
					πF.SetLineno(72)
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
			if πE = πF.Globals().SetItem(πF, ßlatexml.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: """Convert LaTeX math code to MathML with LaTeXML_
			πF.SetLineno(29)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Convert LaTeX math code to MathML with LaTeXML_\n\n    .. _LaTeXML: http://dlmf.nist.gov/LaTeXML/\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlatexml); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 74: def ttm(math_code, reporter=None):
			πF.SetLineno(74)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "math_code", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "reporter", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("ttm", "/usr/lib/python2.7/site-packages/docutils/utils/math/tex2mathml_extern.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmath_code *πg.Object = πArgs[0]
				_ = µmath_code
				var µreporter *πg.Object = πArgs[1]
				_ = µreporter
				var µp *πg.Object = πg.UnboundLocal
				_ = µp
				var µresult *πg.Object = πg.UnboundLocal
				_ = µresult
				var µerr *πg.Object = πg.UnboundLocal
				_ = µerr
				var µmsg *πg.Object = πg.UnboundLocal
				_ = µmsg
				var µstart *πg.Object = πg.UnboundLocal
				_ = µstart
				var µend *πg.Object = πg.UnboundLocal
				_ = µend
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
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
					// line 75: """Convert LaTeX math code to MathML with TtM_
					πF.SetLineno(75)
					// line 79: p = subprocess.Popen(['ttm',
					πF.SetLineno(79)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = make([]*πg.Object, 3)
					πTemp002[0] = ßttm.ToObject()
					πTemp002[1] = πg.NewStr("-u").ToObject()
					πTemp002[2] = πg.NewStr("-r").ToObject()
					πTemp003 = πg.NewList(πTemp002...).ToObject()
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"stdin", πTemp004},
						{"stdout", πTemp005},
						{"stderr", πTemp006},
						{"close_fds", πTemp003},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					// line 88: p.stdin.write((document_template % math_code).encode('utf8'))
					πF.SetLineno(88)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ßutf8.ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßdocument_template); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πTemp004, µmath_code); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 89: p.stdin.close()
					πF.SetLineno(89)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 90: result = p.stdout.read()
					πF.SetLineno(90)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µresult = πTemp003
					// line 91: err = p.stderr.read().decode('utf8')
					πF.SetLineno(91)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("**** Unknown").ToObject()
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µerr, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label1
					}
					goto Label2
					// line 92: if err.find('**** Unknown') >= 0:
					πF.SetLineno(92)
				Label1:
					// line 93: msg = '\n'.join([line for line in err.splitlines()
					πF.SetLineno(93)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/math/tex2mathml_extern.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
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
						var πTemp006 []*πg.Object
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
								case 6:
									goto Label6
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µerr, ßsplitlines, nil); πE != nil {
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
									µline = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewStr("****").ToObject()
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µline, ßstartswith, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label4
								}
								goto Label5
								// line 93: msg = '\n'.join([line for line in err.splitlines()
								πF.SetLineno(93)
							Label4:
								// line 93: msg = '\n'.join([line for line in err.splitlines()
								πF.SetLineno(93)
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µline, nil
							Label6:
								πTemp002 = πSent
								goto Label5
							Label5:
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
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmsg = πTemp005
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πg.NewStr("\nMessage from external converter TtM:\n").ToObject(), µmsg); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 95: raise SyntaxError('\nMessage from external converter TtM:\n'+ msg)
					πF.SetLineno(95)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					πTemp005 = µreporter
					if πTemp010, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if !πTemp010 {
						goto Label4
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("**** Error").ToObject()
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, µerr, ßfind, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.GE(πF, πTemp012, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
				Label4:
					πTemp003 = πTemp005
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp010).ToObject()
					πTemp003 = πTemp005
				Label3:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label5
					}
					goto Label6
					// line 96: if reporter and err.find('**** Error') >= 0 or not result:
					πF.SetLineno(96)
				Label5:
					// line 97: reporter.error(err)
					πF.SetLineno(97)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp001[0] = µerr
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µreporter, ßerror, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label6
				Label6:
					// line 98: start, end = result.find('<math'), result.find('</math>')+7
					πF.SetLineno(98)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<math").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("</math>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.Add(πF, πTemp012, πg.NewInt(7).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp006, πTemp005).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
						continue
					}
					µstart = πTemp005
					µend = πTemp006
					// line 99: result = result[start:end]
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µresult, πTemp003); πE != nil {
						continue
					}
					µresult = πTemp005
					// line 100: return result
					πF.SetLineno(100)
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
			if πE = πF.Globals().SetItem(πF, ßttm.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 75: """Convert LaTeX math code to MathML with TtM_
			πF.SetLineno(75)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Convert LaTeX math code to MathML with TtM_\n\n    .. _TtM: http://hutchinson.belmont.ma.us/tth/mml/\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßttm); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 102: def blahtexml(math_code, inline=True, reporter=None):
			πF.SetLineno(102)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "math_code", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "inline", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "reporter", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("blahtexml", "/usr/lib/python2.7/site-packages/docutils/utils/math/tex2mathml_extern.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmath_code *πg.Object = πArgs[0]
				_ = µmath_code
				var µinline *πg.Object = πArgs[1]
				_ = µinline
				var µreporter *πg.Object = πArgs[2]
				_ = µreporter
				var µoptions *πg.Object = πg.UnboundLocal
				_ = µoptions
				var µmathmode_arg *πg.Object = πg.UnboundLocal
				_ = µmathmode_arg
				var µp *πg.Object = πg.UnboundLocal
				_ = µp
				var µresult *πg.Object = πg.UnboundLocal
				_ = µresult
				var µerr *πg.Object = πg.UnboundLocal
				_ = µerr
				var µstart *πg.Object = πg.UnboundLocal
				_ = µstart
				var µend *πg.Object = πg.UnboundLocal
				_ = µend
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 πg.KWArgs
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
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
					// line 103: """Convert LaTeX math code to MathML with blahtexml_
					πF.SetLineno(103)
					// line 107: options = ['--mathml',
					πF.SetLineno(107)
					πTemp001 = make([]*πg.Object, 10)
					πTemp001[0] = πg.NewStr("--mathml").ToObject()
					πTemp001[1] = πg.NewStr("--indented").ToObject()
					πTemp001[2] = πg.NewStr("--spacing").ToObject()
					πTemp001[3] = ßmoderate.ToObject()
					πTemp001[4] = πg.NewStr("--mathml-encoding").ToObject()
					πTemp001[5] = ßraw.ToObject()
					πTemp001[6] = πg.NewStr("--other-encoding").ToObject()
					πTemp001[7] = ßraw.ToObject()
					πTemp001[8] = πg.NewStr("--doctype-xhtml+mathml").ToObject()
					πTemp001[9] = πg.NewStr("--annotate-TeX").ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µoptions = πTemp002
					if πE = πg.CheckLocal(πF, µinline, "inline"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, µinline); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 115: if inline:
					πF.SetLineno(115)
				Label1:
					// line 116: mathmode_arg = ''
					πF.SetLineno(116)
					µmathmode_arg = ß.ToObject()
					goto Label3
				Label2:
					// line 118: mathmode_arg = 'mode="display"'
					πF.SetLineno(118)
					µmathmode_arg = πg.NewStr("mode=\"display\"").ToObject()
					// line 119: options.append('--displaymath')
					πF.SetLineno(119)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("--displaymath").ToObject()
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µoptions, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label3
				Label3:
					// line 121: p = subprocess.Popen(['blahtexml']+options,
					πF.SetLineno(121)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = ßblahtexml.ToObject()
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp004, µoptions); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßPIPE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"stdin", πTemp004},
						{"stdout", πTemp006},
						{"stderr", πTemp007},
						{"close_fds", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßPopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp002
					// line 126: p.stdin.write(math_code.encode('utf8'))
					πF.SetLineno(126)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µmath_code, "math_code"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µmath_code, ßencode, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp004
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 127: p.stdin.close()
					πF.SetLineno(127)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßclose, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 128: result = p.stdout.read().decode('utf8')
					πF.SetLineno(128)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßread, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µresult = πTemp002
					// line 129: err = p.stderr.read().decode('utf8')
					πF.SetLineno(129)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßutf8.ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßstderr, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßread, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µerr = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<error>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GE(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label4
					}
					goto Label5
					// line 131: if result.find('<error>') >= 0:
					πF.SetLineno(131)
				Label4:
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("<message>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp006, πE = πg.Add(πF, πTemp009, πg.NewInt(9).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("</message>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πTemp009, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µresult, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πg.NewStr("\nMessage from external converter blahtexml:\n").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 132: raise SyntaxError('\nMessage from external converter blahtexml:\n'
					πF.SetLineno(132)
					πE = πF.Raise(πTemp004, nil, nil)
					continue
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					πTemp002 = µreporter
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp003 {
						goto Label6
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("**** Error").ToObject()
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µerr, ßfind, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.GE(πF, πTemp009, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp006
					if πTemp010, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, µresult); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(!πTemp011).ToObject()
					πTemp004 = πTemp006
				Label7:
					πTemp002 = πTemp004
				Label6:
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label8
					}
					goto Label9
					// line 134: if reporter and (err.find('**** Error') >= 0 or not result):
					πF.SetLineno(134)
				Label8:
					// line 135: reporter.error(err)
					πF.SetLineno(135)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
						continue
					}
					πTemp001[0] = µerr
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µreporter, ßerror, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label9
				Label9:
					// line 136: start, end = result.find('<markup>')+9, result.find('</markup>')
					πF.SetLineno(136)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<markup>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.Add(πF, πTemp007, πg.NewInt(9).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("</markup>").ToObject()
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µresult, ßfind, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µstart = πTemp004
					µend = πTemp006
					// line 137: result = ('<math xmlns="http://www.w3.org/1998/Math/MathML"%s>\n'
					πF.SetLineno(137)
					if πE = πg.CheckLocal(πF, µmathmode_arg, "mathmode_arg"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µresult, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µmathmode_arg, πTemp007).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("<math xmlns=\"http://www.w3.org/1998/Math/MathML\"%s>\n%s</math>\n").ToObject(), πTemp004); πE != nil {
						continue
					}
					µresult = πTemp002
					// line 139: return result
					πF.SetLineno(139)
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
			if πE = πF.Globals().SetItem(πF, ßblahtexml.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 103: """Convert LaTeX math code to MathML with blahtexml_
			πF.SetLineno(103)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Convert LaTeX math code to MathML with blahtexml_\n\n    .. _blahtexml: http://gva.noekeon.org/blahtexml/\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßblahtexml); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp006, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 143: if __name__ == "__main__":
			πF.SetLineno(143)
		Label1:
			// line 144: example = (u'\\frac{\\partial \\sin^2(\\alpha)}{\\partial \\vec r}'
			πF.SetLineno(144)
			if πE = πF.Globals().SetItem(πF, ßexample.ToObject(), πg.NewUnicode("\\frac{\\partial \\sin^2(\\alpha)}{\\partial \\vec r}\\varpi \\, \\text{Gr\xc3\xbc\xc3\x9fe}").ToObject()); πE != nil {
				continue
			}
			// line 148: print(blahtexml(example).encode('utf8'))
			πF.SetLineno(148)
			πTemp002 = πF.MakeArgs(1)
			πTemp009 = πF.MakeArgs(1)
			πTemp009[0] = ßutf8.ToObject()
			πTemp010 = πF.MakeArgs(1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßexample); πE != nil {
				continue
			}
			πTemp010[0] = πTemp006
			if πTemp006, πE = πg.ResolveGlobal(πF, ßblahtexml); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp010)
			if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßencode, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp009)
			πTemp002[0] = πTemp007
			if πTemp006, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("tex2mathml_extern", Code)
}
