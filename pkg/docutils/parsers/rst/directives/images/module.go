package images

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/Image"
	_ "github.com/pygolin/stdlib/pkg/PIL"
	_ "github.com/pygolin/stdlib/pkg/PIL/Image"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/urllib"
	_ "github.com/pygolin/stdlib/pkg/urllib/request"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßDirective := πg.InternStr("Directive")
		ßElement := πg.InternStr("Element")
		ßFigure := πg.InternStr("Figure")
		ßIOError := πg.InternStr("IOError")
		ßImage := πg.InternStr("Image")
		ßImportError := πg.InternStr("ImportError")
		ßNone := πg.InternStr("None")
		ßPIL := πg.InternStr("PIL")
		ßSubstitutionDef := πg.InternStr("SubstitutionDef")
		ßTrue := πg.InternStr("True")
		ßUnicodeEncodeError := πg.InternStr("UnicodeEncodeError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd := πg.InternStr("add")
		ßadd_name := πg.InternStr("add_name")
		ßalign := πg.InternStr("align")
		ßalign_h_values := πg.InternStr("align_h_values")
		ßalign_v_values := πg.InternStr("align_v_values")
		ßalign_values := πg.InternStr("align_values")
		ßalt := πg.InternStr("alt")
		ßappend := πg.InternStr("append")
		ßarguments := πg.InternStr("arguments")
		ßblock_text := πg.InternStr("block_text")
		ßbottom := πg.InternStr("bottom")
		ßcaption := πg.InternStr("caption")
		ßcenter := πg.InternStr("center")
		ßchildren := πg.InternStr("children")
		ßchoice := πg.InternStr("choice")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcomment := πg.InternStr("comment")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßcopy := πg.InternStr("copy")
		ßdirectives := πg.InternStr("directives")
		ßdocument := πg.InternStr("document")
		ßencode := πg.InternStr("encode")
		ßerror := πg.InternStr("error")
		ßescape2null := πg.InternStr("escape2null")
		ßfigclass := πg.InternStr("figclass")
		ßfigure := πg.InternStr("figure")
		ßfigwidth := πg.InternStr("figwidth")
		ßfigwidth_value := πg.InternStr("figwidth_value")
		ßfile_insertion_enabled := πg.InternStr("file_insertion_enabled")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßfully_normalize_name := πg.InternStr("fully_normalize_name")
		ßgetfilesystemencoding := πg.InternStr("getfilesystemencoding")
		ßhas_content := πg.InternStr("has_content")
		ßheight := πg.InternStr("height")
		ßimage := πg.InternStr("image")
		ßindirect_reference_name := πg.InternStr("indirect_reference_name")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßleft := πg.InternStr("left")
		ßlegend := πg.InternStr("legend")
		ßlen := πg.InternStr("len")
		ßlength_or_percentage_or_unitless := πg.InternStr("length_or_percentage_or_unitless")
		ßlength_or_unitless := πg.InternStr("length_or_unitless")
		ßline := πg.InternStr("line")
		ßlineno := πg.InternStr("lineno")
		ßliteral_block := πg.InternStr("literal_block")
		ßlower := πg.InternStr("lower")
		ßmiddle := πg.InternStr("middle")
		ßname := πg.InternStr("name")
		ßnested_parse := πg.InternStr("nested_parse")
		ßnodes := πg.InternStr("nodes")
		ßnote_refname := πg.InternStr("note_refname")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptional_arguments := πg.InternStr("optional_arguments")
		ßoptions := πg.InternStr("options")
		ßparagraph := πg.InternStr("paragraph")
		ßparse_target := πg.InternStr("parse_target")
		ßpercentage := πg.InternStr("percentage")
		ßpop := πg.InternStr("pop")
		ßpx := πg.InternStr("px")
		ßrawsource := πg.InternStr("rawsource")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßrecord_dependencies := πg.InternStr("record_dependencies")
		ßreference := πg.InternStr("reference")
		ßrefname := πg.InternStr("refname")
		ßrefuri := πg.InternStr("refuri")
		ßreplace := πg.InternStr("replace")
		ßreporter := πg.InternStr("reporter")
		ßrequired_arguments := πg.InternStr("required_arguments")
		ßright := πg.InternStr("right")
		ßrun := πg.InternStr("run")
		ßscale := πg.InternStr("scale")
		ßset_classes := πg.InternStr("set_classes")
		ßsettings := πg.InternStr("settings")
		ßsize := πg.InternStr("size")
		ßsource := πg.InternStr("source")
		ßsplitlines := πg.InternStr("splitlines")
		ßstate := πg.InternStr("state")
		ßstate_machine := πg.InternStr("state_machine")
		ßstates := πg.InternStr("states")
		ßsys := πg.InternStr("sys")
		ßsystem_message := πg.InternStr("system_message")
		ßtarget := πg.InternStr("target")
		ßtop := πg.InternStr("top")
		ßunchanged := πg.InternStr("unchanged")
		ßunchanged_required := πg.InternStr("unchanged_required")
		ßuri := πg.InternStr("uri")
		ßurl2pathname := πg.InternStr("url2pathname")
		ßutils := πg.InternStr("utils")
		ßversion_info := πg.InternStr("version_info")
		ßwhitespace_normalize_name := πg.InternStr("whitespace_normalize_name")
		ßwidth := πg.InternStr("width")
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
		var πTemp007 *πg.Dict
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.BaseException
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2:
				goto Label2
			case 5:
				goto Label5
			default:
				panic("unexpected function state")
			}
			// line 5: """
			πF.SetLineno(5)
			// line 5: """
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nDirectives for figures and simple images.\n").ToObject()); πE != nil {
				continue
			}
			// line 9: __docformat__ = 'reStructuredText'
			πF.SetLineno(9)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 12: import sys
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: from docutils import nodes, utils
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: from docutils.parsers.rst import Directive
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDirective); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDirective.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 16: from docutils.parsers.rst import directives, states
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.states"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßstates.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: from docutils.nodes import fully_normalize_name, whitespace_normalize_name
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßfully_normalize_name); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßfully_normalize_name.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßwhitespace_normalize_name); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwhitespace_normalize_name.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 18: from docutils.parsers.rst.roles import set_classes
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.roles"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßset_classes); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßset_classes.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: try: # check for the Python Imaging Library
			πF.SetLineno(20)
			πF.PushCheckpoint(2)
			// line 21: import PIL.Image
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "PIL.Image"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßPIL.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp004, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label3
			}
			πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 22: except ImportError:
			πF.SetLineno(22)
		Label3:
			// line 23: try:  # sometimes PIL modules are put in PYTHONPATH's root
			πF.SetLineno(23)
			πF.PushCheckpoint(5)
			// line 24: import Image
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "Image"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßImage.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: class PIL(object): pass  # dummy wrapper
			πF.SetLineno(25)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PIL", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 25: class PIL(object): pass  # dummy wrapper
					πF.SetLineno(25)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PIL").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPIL.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 26: PIL.Image = Image
			πF.SetLineno(26)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImage); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßPIL); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ßImage, πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp009, πTemp005 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label6
			}
			πE = πF.Raise(πTemp009.ToObject(), nil, πTemp005.ToObject())
			continue
			// line 27: except ImportError:
			πF.SetLineno(27)
		Label6:
			// line 28: PIL = None
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPIL.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp008, πTemp003); πE != nil {
				continue
			}
			if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp006 {
				goto Label7
			}
			goto Label8
			// line 30: if sys.version_info >= (3, 0):
			πF.SetLineno(30)
		Label7:
			// line 31: from urllib.request import url2pathname
			πF.SetLineno(31)
			if πTemp002, πE = πg.ImportModule(πF, "urllib.request"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßurl2pathname); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßurl2pathname.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label9
		Label8:
			// line 33: from urllib import url2pathname
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "urllib"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßurl2pathname); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßurl2pathname.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label9
		Label9:
			// line 36: class Image(Directive):
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Image", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
				_ = πTemp004
				var πTemp005 *πg.Dict
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 38: align_h_values = ('left', 'center', 'right')
					πF.SetLineno(38)
					πTemp001 = πg.NewTuple3(ßleft.ToObject(), ßcenter.ToObject(), ßright.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßalign_h_values.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 39: align_v_values = ('top', 'middle', 'bottom')
					πF.SetLineno(39)
					πTemp001 = πg.NewTuple3(ßtop.ToObject(), ßmiddle.ToObject(), ßbottom.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßalign_v_values.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 40: align_values = align_v_values + align_h_values
					πF.SetLineno(40)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßalign_v_values); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßalign_h_values); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßalign_values.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 42: def align(argument):
					πF.SetLineno(42)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "argument", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("align", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargument *πg.Object = πArgs[0]
						_ = µargument
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 46: return directives.choice(argument, Image.align_values)
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp001[0] = µargument
							if πTemp002, πE = πg.ResolveGlobal(πF, ßImage); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßalign_values, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchoice, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßalign.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 48: required_arguments = 1
					πF.SetLineno(48)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					// line 49: optional_arguments = 0
					πF.SetLineno(49)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 50: final_argument_whitespace = True
					πF.SetLineno(50)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 51: option_spec = {'alt': directives.unchanged,
					πF.SetLineno(51)
					πTemp005 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßalt.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlength_or_unitless, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßheight.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlength_or_percentage_or_unitless, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßwidth.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpercentage, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßscale.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßalign); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßalign.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunchanged_required, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßtarget.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp005.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp002 = πTemp005.ToObject()
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 60: def run(self):
					πF.SetLineno(60)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessages *πg.Object = πg.UnboundLocal
						_ = µmessages
						var µreference *πg.Object = πg.UnboundLocal
						_ = µreference
						var µreference_node *πg.Object = πg.UnboundLocal
						_ = µreference_node
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var µtarget_type *πg.Object = πg.UnboundLocal
						_ = µtarget_type
						var µdata *πg.Object = πg.UnboundLocal
						_ = µdata
						var µimage_node *πg.Object = πg.UnboundLocal
						_ = µimage_node
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []πg.Param
						_ = πTemp010
						var πTemp011 πg.KWArgs
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßalign.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 61: if 'align' in self.options:
							πF.SetLineno(61)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSubstitutionDef, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							πTemp002 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßalign_h_values, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 62: if isinstance(self.state, states.SubstitutionDef):
							πF.SetLineno(62)
						Label3:
							πTemp002 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßalign_v_values, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 64: if self.options['align'] not in self.align_v_values:
							πF.SetLineno(64)
						Label6:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp006 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßalign_v_values, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πTemp006, πE = πg.GetAttr(πF, πg.NewStr("\", \"").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp002 = πg.NewTuple3(πTemp005, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error in \"%s\" directive: \"%s\" is not a valid value for the \"align\" option within a substitution definition.  Valid values for \"align\" are: \"%s\".").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 65: raise self.error(
							πF.SetLineno(65)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label7
						Label7:
							goto Label5
							// line 71: elif self.options['align'] not in self.align_h_values:
							πF.SetLineno(71)
						Label4:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp006 = ßalign.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßalign_h_values, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πTemp006, πE = πg.GetAttr(πF, πg.NewStr("\", \"").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp002 = πg.NewTuple3(πTemp005, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Error in \"%s\" directive: \"%s\" is not a valid value for the \"align\" option.  Valid values for \"align\" are: \"%s\".").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 72: raise self.error(
							πF.SetLineno(72)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 77: messages = []
							πF.SetLineno(77)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µmessages = πTemp001
							// line 78: reference = directives.uri(self.arguments[0])
							πF.SetLineno(78)
							πTemp004 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßuri, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µreference = πTemp001
							// line 79: self.options['uri'] = reference
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µreference, "reference"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreference); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp005 = ßuri.ToObject()
							if πE = πg.SetItem(πF, πTemp002, πTemp005, πTemp001); πE != nil {
								continue
							}
							// line 80: reference_node = None
							πF.SetLineno(80)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µreference_node = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßtarget.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 81: if 'target' in self.options:
							πF.SetLineno(81)
						Label8:
							// line 82: block = states.escape2null(
							πF.SetLineno(82)
							πTemp004 = πF.MakeArgs(1)
							πTemp001 = ßtarget.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßescape2null, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µblock = πTemp001
							// line 84: block = [line for line in block]
							πF.SetLineno(84)
							πTemp010 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
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
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µblock); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp002 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp002 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp003 = !isStop
										} else {
											πTemp003 = true
											µline = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 84: block = [line for line in block]
										πF.SetLineno(84)
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return µline, nil
									Label4:
										πTemp004 = πSent
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
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µblock = πTemp001
							// line 85: target_type, data = self.state.parse_target(
							πF.SetLineno(85)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp004[0] = µblock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßparse_target, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
								continue
							}
							µtarget_type = πTemp005
							µdata = πTemp006
							if πE = πg.CheckLocal(πF, µtarget_type, "target_type"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µtarget_type, ßrefuri.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µtarget_type, "target_type"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µtarget_type, ßrefname.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 87: if target_type == 'refuri':
							πF.SetLineno(87)
						Label10:
							// line 88: reference_node = nodes.reference(refuri=data)
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"refuri", µdata},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, nil, πTemp011); πE != nil {
								continue
							}
							µreference_node = πTemp001
							goto Label13
							// line 89: elif target_type == 'refname':
							πF.SetLineno(89)
						Label11:
							// line 90: reference_node = nodes.reference(
							πF.SetLineno(90)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßfully_normalize_name); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßwhitespace_normalize_name); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp011 = πg.KWArgs{
								{"refname", πTemp005},
								{"name", πTemp006},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, nil, πTemp011); πE != nil {
								continue
							}
							µreference_node = πTemp001
							// line 93: reference_node.indirect_reference_name = data
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdata); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µreference_node, "reference_node"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µreference_node, ßindirect_reference_name, πTemp001); πE != nil {
								continue
							}
							// line 94: self.state.document.note_refname(reference_node)
							πF.SetLineno(94)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreference_node, "reference_node"); πE != nil {
								continue
							}
							πTemp004[0] = µreference_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßnote_refname, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label12:
							// line 96: messages.append(data)       # data is a system message
							πF.SetLineno(96)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label13:
							// line 97: del self.options['target']
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp005 = ßtarget.ToObject()
							if πE = πg.DelItem(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							goto Label9
						Label9:
							// line 98: set_classes(self.options)
							πF.SetLineno(98)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 99: image_node = nodes.image(self.block_text, **self.options)
							πF.SetLineno(99)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßimage, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, πTemp006, πTemp004, nil, nil, πTemp001); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µimage_node = πTemp005
							// line 100: self.add_name(image_node)
							πF.SetLineno(100)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							πTemp004[0] = µimage_node
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_name, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µreference_node, "reference_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µreference_node); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							goto Label15
							// line 101: if reference_node:
							πF.SetLineno(101)
						Label14:
							// line 102: reference_node += image_node
							πF.SetLineno(102)
							if πE = πg.CheckLocal(πF, µreference_node, "reference_node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µreference_node, µimage_node); πE != nil {
								continue
							}
							µreference_node = πTemp001
							// line 103: return messages + [reference_node]
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µreference_node, "reference_node"); πE != nil {
								continue
							}
							πTemp004[0] = µreference_node
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp005); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label16
						Label15:
							// line 105: return messages + [image_node]
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							πTemp004[0] = µimage_node
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πTemp001, πE = πg.Add(πF, µmessages, πTemp005); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label16
						Label16:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp002); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Image").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßImage.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 108: class Figure(Image):
			πF.SetLineno(108)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp008, πE = πg.ResolveGlobal(πF, ßImage); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp007 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Figure", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
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
					// line 110: def align(argument):
					πF.SetLineno(110)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "argument", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("align", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargument *πg.Object = πArgs[0]
						_ = µargument
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 111: return directives.choice(argument, Figure.align_h_values)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp001[0] = µargument
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFigure); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßalign_h_values, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßchoice, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßalign.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 113: def figwidth_value(argument):
					πF.SetLineno(113)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "argument", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("figwidth_value", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargument *πg.Object = πArgs[0]
						_ = µargument
						var πTemp001 *πg.Object
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
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µargument, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, ßimage.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 114: if argument.lower() == 'image':
							πF.SetLineno(114)
						Label1:
							// line 115: return 'image'
							πF.SetLineno(115)
							πR = ßimage.ToObject()
							continue
							goto Label3
						Label2:
							// line 117: return directives.length_or_percentage_or_unitless(argument, 'px')
							πF.SetLineno(117)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp005[0] = µargument
							πTemp005[1] = ßpx.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlength_or_percentage_or_unitless, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πR = πTemp001
							continue
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfigwidth_value.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 119: option_spec = Image.option_spec.copy()
					πF.SetLineno(119)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßImage); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßoption_spec, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßcopy, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 120: option_spec['figwidth'] = figwidth_value
					πF.SetLineno(120)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßfigwidth_value); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßoption_spec); πE != nil {
						continue
					}
					πTemp007 = ßfigwidth.ToObject()
					if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp005); πE != nil {
						continue
					}
					// line 121: option_spec['figclass'] = directives.class_option
					πF.SetLineno(121)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßdirectives); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßoption_spec); πE != nil {
						continue
					}
					πTemp007 = ßfigclass.ToObject()
					if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp004); πE != nil {
						continue
					}
					// line 122: option_spec['align'] = align
					πF.SetLineno(122)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßalign); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßoption_spec); πE != nil {
						continue
					}
					πTemp007 = ßalign.ToObject()
					if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp005); πE != nil {
						continue
					}
					// line 123: has_content = True
					πF.SetLineno(123)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 125: def run(self):
					πF.SetLineno(125)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/directives/images.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfigwidth *πg.Object = πg.UnboundLocal
						_ = µfigwidth
						var µfigclasses *πg.Object = πg.UnboundLocal
						_ = µfigclasses
						var µalign *πg.Object = πg.UnboundLocal
						_ = µalign
						var µimage_node *πg.Object = πg.UnboundLocal
						_ = µimage_node
						var µfigure_node *πg.Object = πg.UnboundLocal
						_ = µfigure_node
						var µimagepath *πg.Object = πg.UnboundLocal
						_ = µimagepath
						var µimg *πg.Object = πg.UnboundLocal
						_ = µimg
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µfirst_node *πg.Object = πg.UnboundLocal
						_ = µfirst_node
						var µcaption *πg.Object = πg.UnboundLocal
						_ = µcaption
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 πg.KWArgs
						_ = πTemp011
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10:
								goto Label10
							default:
								panic("unexpected function state")
							}
							// line 126: figwidth = self.options.pop('figwidth', None)
							πF.SetLineno(126)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßfigwidth.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfigwidth = πTemp002
							// line 127: figclasses = self.options.pop('figclass', None)
							πF.SetLineno(127)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßfigclass.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfigclasses = πTemp002
							// line 128: align = self.options.pop('align', None)
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßalign.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µalign = πTemp002
							// line 129: (image_node,) = Image.run(self)
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßImage); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µimage_node = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							πTemp001[0] = µimage_node
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsystem_message, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 130: if isinstance(image_node, nodes.system_message):
							πF.SetLineno(130)
						Label1:
							// line 131: return [image_node]
							πF.SetLineno(131)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							πTemp001[0] = µimage_node
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 132: figure_node = nodes.figure('', image_node)
							πF.SetLineno(132)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							πTemp001[1] = µimage_node
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfigure, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfigure_node = πTemp002
							if πE = πg.CheckLocal(πF, µfigwidth, "figwidth"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µfigwidth, ßimage.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µfigwidth, "figwidth"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µfigwidth != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 133: if figwidth == 'image':
							πF.SetLineno(133)
						Label3:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßPIL); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßfile_insertion_enabled, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 134: if PIL and self.state.document.settings.file_insertion_enabled:
							πF.SetLineno(134)
						Label7:
							// line 135: imagepath = url2pathname(image_node['uri'])
							πF.SetLineno(135)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = ßuri.ToObject()
							if πE = πg.CheckLocal(πF, µimage_node, "image_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µimage_node, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßurl2pathname); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µimagepath = πTemp003
							// line 136: try:
							πF.SetLineno(136)
							πF.PushCheckpoint(10)
							// line 137: img = PIL.Image.open(
							πF.SetLineno(137)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetfilesystemencoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µimagepath, "imagepath"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µimagepath, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPIL); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßImage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µimg = πTemp003
							πF.PopCheckpoint()
							// line 142: self.state.document.settings.record_dependencies.add(
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("\\").ToObject()
							πTemp006[1] = πg.NewStr("/").ToObject()
							if πE = πg.CheckLocal(πF, µimagepath, "imagepath"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µimagepath, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 144: figure_node['width'] = '%dpx' % img.size[0]
							πF.SetLineno(144)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µimg, "img"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µimg, ßsize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%dpx").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp005 = ßwidth.ToObject()
							if πE = πg.SetItem(πF, µfigure_node, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 145: del img
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µimg, "img"); πE != nil {
								continue
							}
							µimg = πg.UnboundLocal
							goto Label9
						Label10:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßUnicodeEncodeError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
							if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 139: except (IOError, UnicodeEncodeError):
							πF.SetLineno(139)
						Label11:
							// line 140: pass # TODO: warn?
							πF.SetLineno(140)
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							goto Label8
						Label8:
							goto Label5
							// line 146: elif figwidth is not None:
							πF.SetLineno(146)
						Label4:
							// line 147: figure_node['width'] = figwidth
							πF.SetLineno(147)
							if πE = πg.CheckLocal(πF, µfigwidth, "figwidth"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µfigwidth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp003 = ßwidth.ToObject()
							if πE = πg.SetItem(πF, µfigure_node, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µfigclasses, "figclasses"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µfigclasses); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 148: if figclasses:
							πF.SetLineno(148)
						Label12:
							// line 149: figure_node['classes'] += figclasses
							πF.SetLineno(149)
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µfigure_node, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfigclasses, "figclasses"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp003, µfigclasses); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp005 = ßclasses.ToObject()
							if πE = πg.SetItem(πF, µfigure_node, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µalign, "align"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µalign); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 150: if align:
							πF.SetLineno(150)
						Label14:
							// line 151: figure_node['align'] = align
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µalign, "align"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µalign); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp003 = ßalign.ToObject()
							if πE = πg.SetItem(πF, µfigure_node, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label15
						Label15:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							goto Label17
							// line 152: if self.content:
							πF.SetLineno(152)
						Label16:
							// line 153: node = nodes.Element()          # anonymous container for parsing
							πF.SetLineno(153)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßElement, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnode = πTemp002
							// line 154: self.state.nested_parse(self.content, self.content_offset, node)
							πF.SetLineno(154)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[2] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnested_parse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 155: first_node = node[0]
							πF.SetLineno(155)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							µfirst_node = πTemp003
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst_node
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst_node
							if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßcomment, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label19
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst_node
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Eq(πF, πTemp010, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
						Label19:
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							goto Label21
							// line 156: if isinstance(first_node, nodes.paragraph):
							πF.SetLineno(156)
						Label18:
							// line 157: caption = nodes.caption(first_node.rawsource, '',
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfirst_node, ßrawsource, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfirst_node, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßcaption, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp005, πTemp001, πTemp002, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcaption = πTemp003
							// line 159: caption.source = first_node.source
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfirst_node, ßsource, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcaption, "caption"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcaption, ßsource, πTemp003); πE != nil {
								continue
							}
							// line 160: caption.line = first_node.line
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µfirst_node, "first_node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfirst_node, ßline, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcaption, "caption"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcaption, ßline, πTemp003); πE != nil {
								continue
							}
							// line 161: figure_node += caption
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcaption, "caption"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µfigure_node, µcaption); πE != nil {
								continue
							}
							µfigure_node = πTemp002
							goto Label21
							// line 162: elif not (isinstance(first_node, nodes.comment)
							πF.SetLineno(162)
						Label20:
							// line 164: error = self.state_machine.reporter.error(
							πF.SetLineno(164)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Figure caption must be a paragraph or empty comment.").ToObject()
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"line", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µerror = πTemp003
							// line 168: return [figure_node, error]
							πF.SetLineno(168)
							πTemp001 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp001[0] = µfigure_node
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp001[1] = µerror
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
							goto Label21
						Label21:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label22
							}
							goto Label23
							// line 169: if len(node) > 1:
							πF.SetLineno(169)
						Label22:
							// line 170: figure_node += nodes.legend('', *node[1:])
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßlegend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp005, πTemp001, πTemp003, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, µfigure_node, πTemp002); πE != nil {
								continue
							}
							µfigure_node = πTemp003
							goto Label23
						Label23:
							goto Label17
						Label17:
							// line 171: return [figure_node]
							πF.SetLineno(171)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µfigure_node, "figure_node"); πE != nil {
								continue
							}
							πTemp001[0] = µfigure_node
							πTemp002 = πg.NewList(πTemp001...).ToObject()
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp008, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Figure").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFigure.ToObject(), πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("images", Code)
}
