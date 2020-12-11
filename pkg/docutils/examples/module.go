package examples

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/examples.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßNone := πg.InternStr("None")
		ßNullOutput := πg.InternStr("NullOutput")
		ßStringInput := πg.InternStr("StringInput")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ßcopy := πg.InternStr("copy")
		ßcore := πg.InternStr("core")
		ßdoctitle_xform := πg.InternStr("doctitle_xform")
		ßdocument := πg.InternStr("document")
		ßencode := πg.InternStr("encode")
		ßhtml := πg.InternStr("html")
		ßhtml_body := πg.InternStr("html_body")
		ßhtml_parts := πg.InternStr("html_parts")
		ßinitial_header_level := πg.InternStr("initial_header_level")
		ßinput_encoding := πg.InternStr("input_encoding")
		ßinternals := πg.InternStr("internals")
		ßio := πg.InternStr("io")
		ßnull := πg.InternStr("null")
		ßpublish_parts := πg.InternStr("publish_parts")
		ßpublish_programmatically := πg.InternStr("publish_programmatically")
		ßrestructuredtext := πg.InternStr("restructuredtext")
		ßstandalone := πg.InternStr("standalone")
		ßunicode := πg.InternStr("unicode")
		ßwriter := πg.InternStr("writer")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis module contains practical examples of Docutils client code.\n\nImporting this module from client code is not recommended; its contents are\nsubject to change in future Docutils releases.  Instead, it is recommended\nthat you copy and paste the parts you need into your own code, modifying as\nnecessary.\n").ToObject()); πE != nil {
				continue
			}
			// line 14: from docutils import core, io
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.core"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßcore.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.io"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßio.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: def html_parts(input_string, source_path=None, destination_path=None,
			πF.SetLineno(17)
			πTemp003 = make([]πg.Param, 6)
			πTemp003[0] = πg.Param{Name: "input_string", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "source_path", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "destination_path", Def: πTemp004}
			πTemp003[3] = πg.Param{Name: "input_encoding", Def: ßunicode.ToObject()}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp003[4] = πg.Param{Name: "doctitle", Def: πTemp004}
			πTemp003[5] = πg.Param{Name: "initial_header_level", Def: πg.NewInt(1).ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("html_parts", "/usr/lib/python2.7/site-packages/docutils/examples.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput_string *πg.Object = πArgs[0]
				_ = µinput_string
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µdestination_path *πg.Object = πArgs[2]
				_ = µdestination_path
				var µinput_encoding *πg.Object = πArgs[3]
				_ = µinput_encoding
				var µdoctitle *πg.Object = πArgs[4]
				_ = µdoctitle
				var µinitial_header_level *πg.Object = πArgs[5]
				_ = µinitial_header_level
				var µoverrides *πg.Object = πg.UnboundLocal
				_ = µoverrides
				var µparts *πg.Object = πg.UnboundLocal
				_ = µparts
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
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
					// line 20: """
					πF.SetLineno(20)
					// line 43: overrides = {'input_encoding': input_encoding,
					πF.SetLineno(43)
					πTemp001 = πg.NewDict()
					if πE = πg.CheckLocal(πF, µinput_encoding, "input_encoding"); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßinput_encoding.ToObject(), µinput_encoding); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdoctitle, "doctitle"); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßdoctitle_xform.ToObject(), µdoctitle); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinitial_header_level, "initial_header_level"); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßinitial_header_level.ToObject(), µinitial_header_level); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					µoverrides = πTemp002
					// line 46: parts = core.publish_parts(
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µinput_string, "input_string"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoverrides, "overrides"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"source", µinput_string},
						{"source_path", µsource_path},
						{"destination_path", µdestination_path},
						{"writer_name", ßhtml.ToObject()},
						{"settings_overrides", µoverrides},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcore); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßpublish_parts, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
						continue
					}
					µparts = πTemp002
					// line 50: return parts
					πF.SetLineno(50)
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					πR = µparts
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßhtml_parts.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: """
			πF.SetLineno(20)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Given an input string, returns a dictionary of HTML document parts.\n\n    Dictionary keys are the names of parts, and values are Unicode strings;\n    encoding is up to the client.\n\n    Parameters:\n\n    - `input_string`: A multi-line text string; required.\n    - `source_path`: Path to the source file or object.  Optional, but useful\n      for diagnostic output (system messages).\n    - `destination_path`: Path to the file or object which will receive the\n      output; optional.  Used for determining relative paths (stylesheets,\n      source links, etc.).\n    - `input_encoding`: The encoding of `input_string`.  If it is an encoded\n      8-bit string, provide the correct encoding.  If it is a Unicode string,\n      use \"unicode\", the default.\n    - `doctitle`: Disable the promotion of a lone top-level section title to\n      document title (and subsequent section title to document subtitle\n      promotion); enabled by default.\n    - `initial_header_level`: The initial level for header elements (e.g. 1\n      for \"<h1>\").\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßhtml_parts); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 52: def html_body(input_string, source_path=None, destination_path=None,
			πF.SetLineno(52)
			πTemp003 = make([]πg.Param, 7)
			πTemp003[0] = πg.Param{Name: "input_string", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "source_path", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "destination_path", Def: πTemp005}
			πTemp003[3] = πg.Param{Name: "input_encoding", Def: ßunicode.ToObject()}
			πTemp003[4] = πg.Param{Name: "output_encoding", Def: ßunicode.ToObject()}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp003[5] = πg.Param{Name: "doctitle", Def: πTemp005}
			πTemp003[6] = πg.Param{Name: "initial_header_level", Def: πg.NewInt(1).ToObject()}
			πTemp004 = πg.NewFunction(πg.NewCode("html_body", "/usr/lib/python2.7/site-packages/docutils/examples.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput_string *πg.Object = πArgs[0]
				_ = µinput_string
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µdestination_path *πg.Object = πArgs[2]
				_ = µdestination_path
				var µinput_encoding *πg.Object = πArgs[3]
				_ = µinput_encoding
				var µoutput_encoding *πg.Object = πArgs[4]
				_ = µoutput_encoding
				var µdoctitle *πg.Object = πArgs[5]
				_ = µdoctitle
				var µinitial_header_level *πg.Object = πArgs[6]
				_ = µinitial_header_level
				var µparts *πg.Object = πg.UnboundLocal
				_ = µparts
				var µfragment *πg.Object = πg.UnboundLocal
				_ = µfragment
				var πTemp001 πg.KWArgs
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 []*πg.Object
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
					// line 55: """
					πF.SetLineno(55)
					// line 65: parts = html_parts(
					πF.SetLineno(65)
					if πE = πg.CheckLocal(πF, µinput_string, "input_string"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinput_encoding, "input_encoding"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdoctitle, "doctitle"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinitial_header_level, "initial_header_level"); πE != nil {
						continue
					}
					πTemp001 = πg.KWArgs{
						{"input_string", µinput_string},
						{"source_path", µsource_path},
						{"destination_path", µdestination_path},
						{"input_encoding", µinput_encoding},
						{"doctitle", µdoctitle},
						{"initial_header_level", µinitial_header_level},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhtml_parts); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, πTemp001); πE != nil {
						continue
					}
					µparts = πTemp003
					// line 70: fragment = parts['html_body']
					πF.SetLineno(70)
					πTemp002 = ßhtml_body.ToObject()
					if πE = πg.CheckLocal(πF, µparts, "parts"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µparts, πTemp002); πE != nil {
						continue
					}
					µfragment = πTemp003
					if πE = πg.CheckLocal(πF, µoutput_encoding, "output_encoding"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, µoutput_encoding, ßunicode.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 71: if output_encoding != 'unicode':
					πF.SetLineno(71)
				Label1:
					// line 72: fragment = fragment.encode(output_encoding)
					πF.SetLineno(72)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoutput_encoding, "output_encoding"); πE != nil {
						continue
					}
					πTemp005[0] = µoutput_encoding
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µfragment, ßencode, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µfragment = πTemp003
					goto Label2
				Label2:
					// line 73: return fragment
					πF.SetLineno(73)
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πR = µfragment
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßhtml_body.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 55: """
			πF.SetLineno(55)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n    Given an input string, returns an HTML fragment as a string.\n\n    The return value is the contents of the <body> element.\n\n    Parameters (see `html_parts()` for the remainder):\n\n    - `output_encoding`: The desired encoding of the output.  If a Unicode\n      string is desired, use the default value of \"unicode\" .\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßhtml_body); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 75: def internals(input_string, source_path=None, destination_path=None,
			πF.SetLineno(75)
			πTemp003 = make([]πg.Param, 5)
			πTemp003[0] = πg.Param{Name: "input_string", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[1] = πg.Param{Name: "source_path", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[2] = πg.Param{Name: "destination_path", Def: πTemp006}
			πTemp003[3] = πg.Param{Name: "input_encoding", Def: ßunicode.ToObject()}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp003[4] = πg.Param{Name: "settings_overrides", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("internals", "/usr/lib/python2.7/site-packages/docutils/examples.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µinput_string *πg.Object = πArgs[0]
				_ = µinput_string
				var µsource_path *πg.Object = πArgs[1]
				_ = µsource_path
				var µdestination_path *πg.Object = πArgs[2]
				_ = µdestination_path
				var µinput_encoding *πg.Object = πArgs[3]
				_ = µinput_encoding
				var µsettings_overrides *πg.Object = πArgs[4]
				_ = µsettings_overrides
				var µoverrides *πg.Object = πg.UnboundLocal
				_ = µoverrides
				var µoutput *πg.Object = πg.UnboundLocal
				_ = µoutput
				var µpub *πg.Object = πg.UnboundLocal
				_ = µpub
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 πg.KWArgs
				_ = πTemp013
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
					// line 77: """
					πF.SetLineno(77)
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µsettings_overrides); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 82: if settings_overrides:
					πF.SetLineno(82)
				Label1:
					// line 83: overrides = settings_overrides.copy()
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µsettings_overrides, "settings_overrides"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsettings_overrides, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µoverrides = πTemp003
					goto Label3
				Label2:
					// line 85: overrides = {}
					πF.SetLineno(85)
					πTemp004 = πg.NewDict()
					πTemp002 = πTemp004.ToObject()
					µoverrides = πTemp002
					goto Label3
				Label3:
					// line 86: overrides['input_encoding'] = input_encoding
					πF.SetLineno(86)
					if πE = πg.CheckLocal(πF, µinput_encoding, "input_encoding"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µinput_encoding); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoverrides, "overrides"); πE != nil {
						continue
					}
					πTemp003 = ßinput_encoding.ToObject()
					if πE = πg.SetItem(πF, µoverrides, πTemp003, πTemp002); πE != nil {
						continue
					}
					// line 87: output, pub = core.publish_programmatically(
					πF.SetLineno(87)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßStringInput, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinput_string, "input_string"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource_path, "source_path"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßio); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßNullOutput, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdestination_path, "destination_path"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoverrides, "overrides"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp013 = πg.KWArgs{
						{"source_class", πTemp003},
						{"source", µinput_string},
						{"source_path", µsource_path},
						{"destination_class", πTemp005},
						{"destination", πTemp002},
						{"destination_path", µdestination_path},
						{"reader", πTemp006},
						{"reader_name", ßstandalone.ToObject()},
						{"parser", πTemp007},
						{"parser_name", ßrestructuredtext.ToObject()},
						{"writer", πTemp008},
						{"writer_name", ßnull.ToObject()},
						{"settings", πTemp009},
						{"settings_spec", πTemp010},
						{"settings_overrides", µoverrides},
						{"config_section", πTemp011},
						{"enable_exit_status", πTemp012},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcore); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpublish_programmatically, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, πTemp013); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
						continue
					}
					µoutput = πTemp003
					µpub = πTemp005
					// line 97: return pub.writer.document, pub
					πF.SetLineno(97)
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpub, ßwriter, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßdocument, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpub, "pub"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp005, µpub).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßinternals.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 77: """
			πF.SetLineno(77)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n    Return the document tree and publisher, for exploring Docutils internals.\n\n    Parameters: see `html_parts()`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßinternals); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("examples", Code)
}
