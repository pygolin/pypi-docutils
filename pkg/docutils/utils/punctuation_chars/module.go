package punctuation_chars

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/unicodedata"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/punctuation_chars.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßFalse := πg.InternStr("False")
		ßValueError := πg.InternStr("ValueError")
		ß__doc__ := πg.InternStr("__doc__")
		ßclosers := πg.InternStr("closers")
		ßclosing_delimiters := πg.InternStr("closing_delimiters")
		ßdelimiters := πg.InternStr("delimiters")
		ßget := πg.InternStr("get")
		ßindex := πg.InternStr("index")
		ßmatch_chars := πg.InternStr("match_chars")
		ßmaxunicode := πg.InternStr("maxunicode")
		ßopeners := πg.InternStr("openers")
		ßquote_pairs := πg.InternStr("quote_pairs")
		ßre := πg.InternStr("re")
		ßsys := πg.InternStr("sys")
		ßunicodedata := πg.InternStr("unicodedata")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 18: import sys, re
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: import unicodedata
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "unicodedata"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunicodedata.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: """Docutils character category patterns.
			πF.SetLineno(21)
			// line 21: """Docutils character category patterns.
			πF.SetLineno(21)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Docutils character category patterns.\n\n   Patterns for the implementation of the `inline markup recognition rules`_\n   in the reStructuredText parser `docutils.parsers.rst.states.py` based\n   on Unicode character categories.\n   The patterns are used inside ``[ ]`` in regular expressions.\n\n   Rule (5) requires determination of matching open/close pairs. However, the\n   pairing of open/close quotes is ambiguous due to  different typographic\n   conventions in different languages. The ``quote_pairs`` function tests\n   whether two characters form an open/close pair.\n\n   The patterns are generated by\n   ``docutils/tools/dev/generate_punctuation_chars.py`` to  prevent dependence\n   on the Python version and avoid the time-consuming generation with every\n   Docutils run. See there for motives and implementation details.\n\n   The category of some characters changed with the development of the\n   Unicode standard. The current lists are generated with the help of the\n   \"unicodedata\" module of Python 2.7.13 (based on Unicode version 5.2.0).\n\n   .. _inline markup recognition rules:\n      http://docutils.sf.net/docs/ref/rst/restructuredtext.html#inline-markup-recognition-rules\n").ToObject()); πE != nil {
				continue
			}
			// line 46: openers = (u'"\'(<\\[{\u0f3a\u0f3c\u169b\u2045\u207d\u208d\u2329\u2768'
			πF.SetLineno(46)
			if πE = πF.Globals().SetItem(πF, ßopeners.ToObject(), πg.NewUnicode("\"'(<\\[{\xe0\xbc\xba\xe0\xbc\xbc\xe1\x9a\x9b\xe2\x81\x85\xe2\x81\xbd\xe2\x82\x8d\xe2\x8c\xa9\xe2\x9d\xa8\xe2\x9d\xaa\xe2\x9d\xac\xe2\x9d\xae\xe2\x9d\xb0\xe2\x9d\xb2\xe2\x9d\xb4\xe2\x9f\x85\xe2\x9f\xa6\xe2\x9f\xa8\xe2\x9f\xaa\xe2\x9f\xac\xe2\x9f\xae\xe2\xa6\x83\xe2\xa6\x85\xe2\xa6\x87\xe2\xa6\x89\xe2\xa6\x8b\xe2\xa6\x8d\xe2\xa6\x8f\xe2\xa6\x91\xe2\xa6\x93\xe2\xa6\x95\xe2\xa6\x97\xe2\xa7\x98\xe2\xa7\x9a\xe2\xa7\xbc\xe2\xb8\xa2\xe2\xb8\xa4\xe2\xb8\xa6\xe2\xb8\xa8\xe3\x80\x88\xe3\x80\x8a\xe3\x80\x8c\xe3\x80\x8e\xe3\x80\x90\xe3\x80\x94\xe3\x80\x96\xe3\x80\x98\xe3\x80\x9a\xe3\x80\x9d\xe3\x80\x9d\xef\xb4\xbe\xef\xb8\x97\xef\xb8\xb5\xef\xb8\xb7\xef\xb8\xb9\xef\xb8\xbb\xef\xb8\xbd\xef\xb8\xbf\xef\xb9\x81\xef\xb9\x83\xef\xb9\x87\xef\xb9\x99\xef\xb9\x9b\xef\xb9\x9d\xef\xbc\x88\xef\xbc\xbb\xef\xbd\x9b\xef\xbd\x9f\xef\xbd\xa2\xc2\xab\xe2\x80\x98\xe2\x80\x9c\xe2\x80\xb9\xe2\xb8\x82\xe2\xb8\x84\xe2\xb8\x89\xe2\xb8\x8c\xe2\xb8\x9c\xe2\xb8\xa0\xe2\x80\x9a\xe2\x80\x9e\xc2\xbb\xe2\x80\x99\xe2\x80\x9d\xe2\x80\xba\xe2\xb8\x83\xe2\xb8\x85\xe2\xb8\x8a\xe2\xb8\x8d\xe2\xb8\x9d\xe2\xb8\xa1\xe2\x80\x9b\xe2\x80\x9f").ToObject()); πE != nil {
				continue
			}
			// line 56: closers = (u'"\')>\\]}\u0f3b\u0f3d\u169c\u2046\u207e\u208e\u232a\u2769'
			πF.SetLineno(56)
			if πE = πF.Globals().SetItem(πF, ßclosers.ToObject(), πg.NewUnicode("\"')>\\]}\xe0\xbc\xbb\xe0\xbc\xbd\xe1\x9a\x9c\xe2\x81\x86\xe2\x81\xbe\xe2\x82\x8e\xe2\x8c\xaa\xe2\x9d\xa9\xe2\x9d\xab\xe2\x9d\xad\xe2\x9d\xaf\xe2\x9d\xb1\xe2\x9d\xb3\xe2\x9d\xb5\xe2\x9f\x86\xe2\x9f\xa7\xe2\x9f\xa9\xe2\x9f\xab\xe2\x9f\xad\xe2\x9f\xaf\xe2\xa6\x84\xe2\xa6\x86\xe2\xa6\x88\xe2\xa6\x8a\xe2\xa6\x8c\xe2\xa6\x8e\xe2\xa6\x90\xe2\xa6\x92\xe2\xa6\x94\xe2\xa6\x96\xe2\xa6\x98\xe2\xa7\x99\xe2\xa7\x9b\xe2\xa7\xbd\xe2\xb8\xa3\xe2\xb8\xa5\xe2\xb8\xa7\xe2\xb8\xa9\xe3\x80\x89\xe3\x80\x8b\xe3\x80\x8d\xe3\x80\x8f\xe3\x80\x91\xe3\x80\x95\xe3\x80\x97\xe3\x80\x99\xe3\x80\x9b\xe3\x80\x9e\xe3\x80\x9f\xef\xb4\xbf\xef\xb8\x98\xef\xb8\xb6\xef\xb8\xb8\xef\xb8\xba\xef\xb8\xbc\xef\xb8\xbe\xef\xb9\x80\xef\xb9\x82\xef\xb9\x84\xef\xb9\x88\xef\xb9\x9a\xef\xb9\x9c\xef\xb9\x9e\xef\xbc\x89\xef\xbc\xbd\xef\xbd\x9d\xef\xbd\xa0\xef\xbd\xa3\xc2\xbb\xe2\x80\x99\xe2\x80\x9d\xe2\x80\xba\xe2\xb8\x83\xe2\xb8\x85\xe2\xb8\x8a\xe2\xb8\x8d\xe2\xb8\x9d\xe2\xb8\xa1\xe2\x80\x9b\xe2\x80\x9f\xc2\xab\xe2\x80\x98\xe2\x80\x9c\xe2\x80\xb9\xe2\xb8\x82\xe2\xb8\x84\xe2\xb8\x89\xe2\xb8\x8c\xe2\xb8\x9c\xe2\xb8\xa0\xe2\x80\x9a\xe2\x80\x9e").ToObject()); πE != nil {
				continue
			}
			// line 66: delimiters = (u'\\-/:\u058a\xa1\xb7\xbf\u037e\u0387\u055a-\u055f\u0589'
			πF.SetLineno(66)
			if πE = πF.Globals().SetItem(πF, ßdelimiters.ToObject(), πg.NewUnicode("\\-/:\xd6\x8a\xc2\xa1\xc2\xb7\xc2\xbf\xcd\xbe\xce\x87\xd5\x9a-\xd5\x9f\xd6\x89\xd6\xbe\xd7\x80\xd7\x83\xd7\x86\xd7\xb3\xd7\xb4\xd8\x89\xd8\x8a\xd8\x8c\xd8\x8d\xd8\x9b\xd8\x9e\xd8\x9f\xd9\xaa-\xd9\xad\xdb\x94\xdc\x80-\xdc\x8d\xdf\xb7-\xdf\xb9\xe0\xa0\xb0-\xe0\xa0\xbe\xe0\xa5\xa4\xe0\xa5\xa5\xe0\xa5\xb0\xe0\xb7\xb4\xe0\xb9\x8f\xe0\xb9\x9a\xe0\xb9\x9b\xe0\xbc\x84-\xe0\xbc\x92\xe0\xbe\x85\xe0\xbf\x90-\xe0\xbf\x94\xe1\x81\x8a-\xe1\x81\x8f\xe1\x83\xbb\xe1\x8d\xa1-\xe1\x8d\xa8\xe1\x90\x80\xe1\x99\xad\xe1\x99\xae\xe1\x9b\xab-\xe1\x9b\xad\xe1\x9c\xb5\xe1\x9c\xb6\xe1\x9f\x94-\xe1\x9f\x96\xe1\x9f\x98-\xe1\x9f\x9a\xe1\xa0\x80-\xe1\xa0\x8a\xe1\xa5\x84\xe1\xa5\x85\xe1\xa7\x9e\xe1\xa7\x9f\xe1\xa8\x9e\xe1\xa8\x9f\xe1\xaa\xa0-\xe1\xaa\xa6\xe1\xaa\xa8-\xe1\xaa\xad\xe1\xad\x9a-\xe1\xad\xa0\xe1\xb0\xbb-\xe1\xb0\xbf\xe1\xb1\xbe\xe1\xb1\xbf\xe1\xb3\x93\xe2\x80\x90-\xe2\x80\x97\xe2\x80\xa0-\xe2\x80\xa7\xe2\x80\xb0-\xe2\x80\xb8\xe2\x80\xbb-\xe2\x80\xbe\xe2\x81\x81-\xe2\x81\x83\xe2\x81\x87-\xe2\x81\x91\xe2\x81\x93\xe2\x81\x95-\xe2\x81\x9e\xe2\xb3\xb9-\xe2\xb3\xbc\xe2\xb3\xbe\xe2\xb3\xbf\xe2\xb8\x80\xe2\xb8\x81\xe2\xb8\x86-\xe2\xb8\x88\xe2\xb8\x8b\xe2\xb8\x8e-\xe2\xb8\x9b\xe2\xb8\x9e\xe2\xb8\x9f\xe2\xb8\xaa-\xe2\xb8\xae\xe2\xb8\xb0\xe2\xb8\xb1\xe3\x80\x81-\xe3\x80\x83\xe3\x80\x9c\xe3\x80\xb0\xe3\x80\xbd\xe3\x82\xa0\xe3\x83\xbb\xea\x93\xbe\xea\x93\xbf\xea\x98\x8d-\xea\x98\x8f\xea\x99\xb3\xea\x99\xbe\xea\x9b\xb2-\xea\x9b\xb7\xea\xa1\xb4-\xea\xa1\xb7\xea\xa3\x8e\xea\xa3\x8f\xea\xa3\xb8-\xea\xa3\xba\xea\xa4\xae\xea\xa4\xaf\xea\xa5\x9f\xea\xa7\x81-\xea\xa7\x8d\xea\xa7\x9e\xea\xa7\x9f\xea\xa9\x9c-\xea\xa9\x9f\xea\xab\x9e\xea\xab\x9f\xea\xaf\xab\xef\xb8\x90-\xef\xb8\x96\xef\xb8\x99\xef\xb8\xb0-\xef\xb8\xb2\xef\xb9\x85\xef\xb9\x86\xef\xb9\x89-\xef\xb9\x8c\xef\xb9\x90-\xef\xb9\x92\xef\xb9\x94-\xef\xb9\x98\xef\xb9\x9f-\xef\xb9\xa1\xef\xb9\xa3\xef\xb9\xa8\xef\xb9\xaa\xef\xb9\xab\xef\xbc\x81-\xef\xbc\x83\xef\xbc\x85-\xef\xbc\x87\xef\xbc\x8a\xef\xbc\x8c-\xef\xbc\x8f\xef\xbc\x9a\xef\xbc\x9b\xef\xbc\x9f\xef\xbc\xa0\xef\xbc\xbc\xef\xbd\xa1\xef\xbd\xa4\xef\xbd\xa5").ToObject()); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaxunicode, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.GE(πF, πTemp004, πg.NewInt(1114111).ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 86: if sys.maxunicode >= 0x10FFFF: # "wide" build
			πF.SetLineno(86)
		Label1:
			// line 87: delimiters += (u'\U00010100\U00010101\U0001039f\U000103d0\U00010857'
			πF.SetLineno(87)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßdelimiters); πE != nil {
				continue
			}
			if πTemp003, πE = πg.IAdd(πF, πTemp001, πg.NewUnicode("\xf0\x90\x84\x80\xf0\x90\x84\x81\xf0\x90\x8e\x9f\xf0\x90\x8f\x90\xf0\x90\xa1\x97\xf0\x90\xa4\x9f\xf0\x90\xa4\xbf\xf0\x90\xa9\x90-\xf0\x90\xa9\x98\xf0\x90\xa9\xbf\xf0\x90\xac\xb9-\xf0\x90\xac\xbf\xf0\x91\x82\xbb\xf0\x91\x82\xbc\xf0\x91\x82\xbe-\xf0\x91\x83\x81\xf0\x92\x91\xb0-\xf0\x92\x91\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdelimiters.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 91: closing_delimiters = u'\\\\.,;!?'
			πF.SetLineno(91)
			if πE = πF.Globals().SetItem(πF, ßclosing_delimiters.ToObject(), πg.NewUnicode("\\\\.,;!?").ToObject()); πE != nil {
				continue
			}
			// line 97: quote_pairs = {# open char: matching closing characters # usage example
			πF.SetLineno(97)
			πTemp006 = πg.NewDict()
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xc2\xbb").ToObject(), πg.NewUnicode("\xc2\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\x98").ToObject(), πg.NewUnicode("\xe2\x80\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\x99").ToObject(), πg.NewUnicode("\xe2\x80\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\x9a").ToObject(), πg.NewUnicode("\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\x9c").ToObject(), πg.NewUnicode("\xe2\x80\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\x9e").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\x9d").ToObject(), πg.NewUnicode("\xe2\x80\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, πg.NewUnicode("\xe2\x80\xba").ToObject(), πg.NewUnicode("\xe2\x80\xba").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ßquote_pairs.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 107: """Additional open/close quote pairs."""
			πF.SetLineno(107)
			// line 109: def match_chars(c1, c2):
			πF.SetLineno(109)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "c1", Def: nil}
			πTemp007[1] = πg.Param{Name: "c2", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("match_chars", "/usr/lib/python2.7/site-packages/docutils/utils/punctuation_chars.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µc1 *πg.Object = πArgs[0]
				_ = µc1
				var µc2 *πg.Object = πArgs[1]
				_ = µc2
				var µi *πg.Object = πg.UnboundLocal
				_ = µi
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
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
					case 2:
						goto Label2
					default:
						panic("unexpected function state")
					}
					// line 110: """Test whether `c1` and `c2` are a matching open/close character pair.
					πF.SetLineno(110)
					// line 118: try:
					πF.SetLineno(118)
					πF.PushCheckpoint(2)
					// line 119: i = openers.index(c1)
					πF.SetLineno(119)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc1, "c1"); πE != nil {
						continue
					}
					πTemp001[0] = µc1
					if πTemp002, πE = πg.ResolveGlobal(πF, ßopeners); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindex, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp002
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 120: except ValueError:  # c1 not in openers
					πF.SetLineno(120)
				Label3:
					// line 121: return False
					πF.SetLineno(121)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 122: return c2 == closers[i] or c2 in quote_pairs.get(c1, u'')
					πF.SetLineno(122)
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πTemp009, πE = πg.ResolveGlobal(πF, ßclosers); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp007); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc2, πTemp008); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µc1, "c1"); πE != nil {
						continue
					}
					πTemp001[0] = µc1
					πTemp001[1] = πg.NewUnicode("").ToObject()
					if πTemp007, πE = πg.ResolveGlobal(πF, ßquote_pairs); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßget, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp010, πE = πg.Contains(πF, πTemp007, µc2); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp010).ToObject()
					πTemp002 = πTemp003
				Label4:
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
			if πE = πF.Globals().SetItem(πF, ßmatch_chars.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 110: """Test whether `c1` and `c2` are a matching open/close character pair.
			πF.SetLineno(110)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Test whether `c1` and `c2` are a matching open/close character pair.\n\n    Matching open/close pairs are at the same position in\n    `punctuation_chars.openers` and `punctuation_chars.closers`.\n    The pairing of open/close quotes is ambiguous due to  different\n    typographic conventions in different languages,\n    so we test for additional matches stored in `quote_pairs`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßmatch_chars); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("punctuation_chars", Code)
}