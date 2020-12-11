package peps

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/time"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßApplicationError := πg.InternStr("ApplicationError")
		ßCallBack := πg.InternStr("CallBack")
		ßContents := πg.InternStr("Contents")
		ßDataError := πg.InternStr("DataError")
		ßHeaders := πg.InternStr("Headers")
		ßIGNORECASE := πg.InternStr("IGNORECASE")
		ßNone := πg.InternStr("None")
		ßPEPZero := πg.InternStr("PEPZero")
		ßPEPZeroSpecial := πg.InternStr("PEPZeroSpecial")
		ßReferences := πg.InternStr("References")
		ßSkipNode := πg.InternStr("SkipNode")
		ßSparseNodeVisitor := πg.InternStr("SparseNodeVisitor")
		ßTargetNotes := πg.InternStr("TargetNotes")
		ßText := πg.InternStr("Text")
		ßTransform := πg.InternStr("Transform")
		ßTransformError := πg.InternStr("TransformError")
		ßValueError := πg.InternStr("ValueError")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßadd_backref := πg.InternStr("add_backref")
		ßappend := πg.InternStr("append")
		ßapply := πg.InternStr("apply")
		ßastext := πg.InternStr("astext")
		ßauthor := πg.InternStr("author")
		ßcallback := πg.InternStr("callback")
		ßclasses := πg.InternStr("classes")
		ßclean_rcs_keywords := πg.InternStr("clean_rcs_keywords")
		ßcleanup_callback := πg.InternStr("cleanup_callback")
		ßcols := πg.InternStr("cols")
		ßcompile := πg.InternStr("compile")
		ßcontents := πg.InternStr("contents")
		ßcopyright := πg.InternStr("copyright")
		ßdefault_priority := πg.InternStr("default_priority")
		ßdocument := πg.InternStr("document")
		ßentry := πg.InternStr("entry")
		ßfield_list := πg.InternStr("field_list")
		ßfully_normalize_name := πg.InternStr("fully_normalize_name")
		ßget_language := πg.InternStr("get_language")
		ßhas_name := πg.InternStr("has_name")
		ßhasattr := πg.InternStr("hasattr")
		ßhtml := πg.InternStr("html")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßlabels := πg.InternStr("labels")
		ßlanguage_code := πg.InternStr("language_code")
		ßlanguages := πg.InternStr("languages")
		ßlen := πg.InternStr("len")
		ßlocaltime := πg.InternStr("localtime")
		ßlower := πg.InternStr("lower")
		ßmask_email := πg.InternStr("mask_email")
		ßmisc := πg.InternStr("misc")
		ßnames := πg.InternStr("names")
		ßnodes := πg.InternStr("nodes")
		ßnon_masked_addresses := πg.InternStr("non_masked_addresses")
		ßnote_implicit_target := πg.InternStr("note_implicit_target")
		ßnote_pending := πg.InternStr("note_pending")
		ßnum := πg.InternStr("num")
		ßos := πg.InternStr("os")
		ßparagraph := πg.InternStr("paragraph")
		ßparent := πg.InternStr("parent")
		ßparts := πg.InternStr("parts")
		ßpending := πg.InternStr("pending")
		ßpep := πg.InternStr("pep")
		ßpep_base_url := πg.InternStr("pep_base_url")
		ßpep_cvs_url := πg.InternStr("pep_cvs_url")
		ßpep_table := πg.InternStr("pep_table")
		ßpep_url := πg.InternStr("pep_url")
		ßpformat := πg.InternStr("pformat")
		ßproblematic := πg.InternStr("problematic")
		ßraw := πg.InternStr("raw")
		ßrcs_keyword_substitutions := πg.InternStr("rcs_keyword_substitutions")
		ßre := πg.InternStr("re")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßreferences := πg.InternStr("references")
		ßrefuri := πg.InternStr("refuri")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßreplace_self := πg.InternStr("replace_self")
		ßreplaces := πg.InternStr("replaces")
		ßreporter := πg.InternStr("reporter")
		ßrequires := πg.InternStr("requires")
		ßrfc2822 := πg.InternStr("rfc2822")
		ßsection := πg.InternStr("section")
		ßset_id := πg.InternStr("set_id")
		ßsettings := πg.InternStr("settings")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßstartnode := πg.InternStr("startnode")
		ßstartswith := πg.InternStr("startswith")
		ßstat := πg.InternStr("stat")
		ßstrftime := πg.InternStr("strftime")
		ßsys := πg.InternStr("sys")
		ßtime := πg.InternStr("time")
		ßtitle := πg.InternStr("title")
		ßtopic := πg.InternStr("topic")
		ßunknown_visit := πg.InternStr("unknown_visit")
		ßutils := πg.InternStr("utils")
		ßversion := πg.InternStr("version")
		ßvisit_colspec := πg.InternStr("visit_colspec")
		ßvisit_entry := πg.InternStr("visit_entry")
		ßvisit_field_list := πg.InternStr("visit_field_list")
		ßvisit_reference := πg.InternStr("visit_reference")
		ßvisit_row := πg.InternStr("visit_row")
		ßvisit_tgroup := πg.InternStr("visit_tgroup")
		ßwalk := πg.InternStr("walk")
		ßwarning := πg.InternStr("warning")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nTransforms for PEP processing.\n\n- `Headers`: Used to transform a PEP's initial RFC-2822 header.  It remains a\n  field list, but some entries get processed.\n- `Contents`: Auto-inserts a table of contents.\n- `PEPZero`: Special processing for PEP 0.\n").ToObject()); πE != nil {
				continue
			}
			// line 14: __docformat__ = 'reStructuredText'
			πF.SetLineno(14)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 16: import sys
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import os
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import re
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: import time
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: from docutils import nodes, utils, languages
			πF.SetLineno(20)
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
			if πTemp002, πE = πg.ImportModule(πF, "docutils.languages"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlanguages.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: from docutils import ApplicationError, DataError
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßApplicationError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßApplicationError.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßDataError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDataError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 22: from docutils.transforms import Transform, TransformError
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransform); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransform.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßTransformError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransformError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 23: from docutils.transforms import parts, references, misc
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.parts"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßparts.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.references"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßreferences.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.misc"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßmisc.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 26: class Headers(Transform):
			πF.SetLineno(26)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Headers", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []πg.Param
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 28: """
					πF.SetLineno(28)
					// line 28: """
					πF.SetLineno(28)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Process fields in a PEP's initial RFC-2822 header.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 32: default_priority = 360
					πF.SetLineno(32)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(360).ToObject()); πE != nil {
						continue
					}
					// line 34: pep_url = 'pep-%04d'
					πF.SetLineno(34)
					if πE = πClass.SetItem(πF, ßpep_url.ToObject(), πg.NewStr("pep-%04d").ToObject()); πE != nil {
						continue
					}
					// line 35: pep_cvs_url = ('http://hg.python.org'
					πF.SetLineno(35)
					if πE = πClass.SetItem(πF, ßpep_cvs_url.ToObject(), πg.NewStr("http://hg.python.org/peps/file/default/pep-%04d.txt").ToObject()); πE != nil {
						continue
					}
					// line 37: rcs_keyword_substitutions = (
					πF.SetLineno(37)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr("\\$RCSfile: (.+),v \\$$").ToObject()
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßIGNORECASE, nil); πE != nil {
						continue
					}
					πTemp003[1] = πTemp005
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002 = πg.NewTuple2(πTemp004, πg.NewStr("\\1").ToObject()).ToObject()
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("\\$[a-zA-Z]+: (.+) \\$$").ToObject()
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp004 = πg.NewTuple2(πTemp005, πg.NewStr("\\1").ToObject()).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
					if πE = πClass.SetItem(πF, ßrcs_keyword_substitutions.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 41: def apply(self):
					πF.SetLineno(41)
					πTemp007 = make([]πg.Param, 1)
					πTemp007[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µheader *πg.Object = πg.UnboundLocal
						_ = µheader
						var µpep *πg.Object = πg.UnboundLocal
						_ = µpep
						var µfield *πg.Object = πg.UnboundLocal
						_ = µfield
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µcvs_url *πg.Object = πg.UnboundLocal
						_ = µcvs_url
						var µmsg *πg.Object = πg.UnboundLocal
						_ = µmsg
						var µmsgid *πg.Object = πg.UnboundLocal
						_ = µmsgid
						var µprb *πg.Object = πg.UnboundLocal
						_ = µprb
						var µprbid *πg.Object = πg.UnboundLocal
						_ = µprbid
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µbody *πg.Object = πg.UnboundLocal
						_ = µbody
						var µdate *πg.Object = πg.UnboundLocal
						_ = µdate
						var µpara *πg.Object = πg.UnboundLocal
						_ = µpara
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var µnewbody *πg.Object = πg.UnboundLocal
						_ = µnewbody
						var µspace *πg.Object = πg.UnboundLocal
						_ = µspace
						var µrefpep *πg.Object = πg.UnboundLocal
						_ = µrefpep
						var µpepno *πg.Object = πg.UnboundLocal
						_ = µpepno
						var µpep_type *πg.Object = πg.UnboundLocal
						_ = µpep_type
						var µuri *πg.Object = πg.UnboundLocal
						_ = µuri
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 []*πg.Object
						_ = πTemp014
						var πTemp015 []*πg.Object
						_ = πTemp015
						var πTemp016 bool
						_ = πTemp016
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6:
								goto Label6
							case 7:
								goto Label7
							case 12:
								goto Label12
							case 45:
								goto Label45
							case 46:
								goto Label46
							case 50:
								goto Label50
							case 51:
								goto Label51
							case 55:
								goto Label55
							case 56:
								goto Label56
							case 25:
								goto Label25
							case 26:
								goto Label26
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 42: if not len(self.document):
							πF.SetLineno(42)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Document tree is empty.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 44: raise DataError('Document tree is empty.')
							πF.SetLineno(44)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 45: header = self.document[0]
							πF.SetLineno(45)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µheader = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp002[0] = µheader
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßfield_list, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πTemp004 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µheader, πTemp004); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp006, ßrfc2822.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label3:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 46: if not isinstance(header, nodes.field_list) or \
							πF.SetLineno(46)
						Label4:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Document does not begin with an RFC-2822 header; it is not a PEP.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 48: raise DataError('Document does not begin with an RFC-2822 '
							πF.SetLineno(48)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 50: pep = None
							πF.SetLineno(50)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µpep = πTemp001
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µheader); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp005 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µfield = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µfield, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßlower, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, ßpep.ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 52: if field[0].astext().lower() == 'pep': # should be the first field
							πF.SetLineno(52)
						Label9:
							// line 53: value = field[1].astext()
							πF.SetLineno(53)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp004
							// line 54: try:
							πF.SetLineno(54)
							πF.PushCheckpoint(12)
							// line 55: pep = int(value)
							πF.SetLineno(55)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µpep = πTemp004
							// line 56: cvs_url = self.pep_cvs_url % pep
							πF.SetLineno(56)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpep_cvs_url, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpep, "pep"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πTemp004, µpep); πE != nil {
								continue
							}
							µcvs_url = πTemp003
							πF.PopCheckpoint()
							goto Label11
						Label12:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label13
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 57: except ValueError:
							πF.SetLineno(57)
						Label13:
							// line 58: pep = value
							πF.SetLineno(58)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							µpep = µvalue
							// line 59: cvs_url = None
							πF.SetLineno(59)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µcvs_url = πTemp003
							// line 60: msg = self.document.reporter.warning(
							πF.SetLineno(60)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpep, "pep"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("\"PEP\" header must contain an integer; \"%s\" is an invalid value.").ToObject(), µpep); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"base_node", µfield},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwarning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmsg = πTemp004
							// line 63: msgid = self.document.set_id(msg)
							πF.SetLineno(63)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmsgid = πTemp003
							// line 64: prb = nodes.problematic(value, value or '(none)',
							πF.SetLineno(64)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003 = µvalue
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label14
							}
							πTemp003 = πg.NewStr("(none)").ToObject()
						Label14:
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µmsgid, "msgid"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"refid", µmsgid},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßproblematic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µprb = πTemp003
							// line 66: prbid = self.document.set_id(prb)
							πF.SetLineno(66)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp002[0] = µprb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µprbid = πTemp003
							// line 67: msg.add_backref(prbid)
							πF.SetLineno(67)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprbid, "prbid"); πE != nil {
								continue
							}
							πTemp002[0] = µprbid
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmsg, ßadd_backref, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label15
							}
							goto Label16
							// line 68: if len(field[1]):
							πF.SetLineno(68)
						Label15:
							// line 69: field[1][0][:] = [prb]
							πF.SetLineno(69)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp002[0] = µprb
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							πTemp006 = πg.NewInt(0).ToObject()
							πTemp012 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µfield, πTemp012); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp011, πTemp006, πTemp004); πE != nil {
								continue
							}
							goto Label17
						Label16:
							// line 71: field[1] += nodes.paragraph('', '', prb)
							πF.SetLineno(71)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = ß.ToObject()
							πTemp002[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
								continue
							}
							πTemp002[2] = µprb
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IAdd(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp011 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µfield, πTemp011, πTemp006); πE != nil {
								continue
							}
							goto Label17
						Label17:
							πF.RestoreExc(nil, nil)
							goto Label11
						Label11:
							// line 72: break
							πF.SetLineno(72)
							πTemp005 = true
							continue
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							if πE = πg.CheckLocal(πF, µpep, "pep"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µpep == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							goto Label19
							// line 73: if pep is None:
							πF.SetLineno(73)
						Label18:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Document does not contain an RFC-2822 \"PEP\" header.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 74: raise DataError('Document does not contain an RFC-2822 "PEP" '
							πF.SetLineno(74)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label19
						Label19:
							if πE = πg.CheckLocal(πF, µpep, "pep"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µpep, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label20
							}
							goto Label21
							// line 76: if pep == 0:
							πF.SetLineno(76)
						Label20:
							// line 78: pending = nodes.pending(PEPZero)
							πF.SetLineno(78)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPEPZero); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µpending = πTemp001
							// line 79: self.document.insert(1, pending)
							πF.SetLineno(79)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[1] = µpending
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 80: self.document.note_pending(pending)
							πF.SetLineno(80)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label21
						Label21:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp002[0] = µheader
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.LT(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							πTemp004 = πg.NewInt(0).ToObject()
							πTemp011 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µheader, πTemp011); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp012, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßlower, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, πTemp006, ßtitle.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label22:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label23
							}
							goto Label24
							// line 81: if len(header) < 2 or header[1][0].astext().lower() != 'title':
							πF.SetLineno(81)
						Label23:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("No title!").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 82: raise DataError('No title!')
							πF.SetLineno(82)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label24
						Label24:
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µheader); πE != nil {
								continue
							}
							πF.PushCheckpoint(26)
							πTemp005 = false
						Label25:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label27
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µfield = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(25)
							// line 84: name = field[0].astext().lower()
							πF.SetLineno(84)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßlower, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µname = πTemp004
							// line 85: body = field[1]
							πF.SetLineno(85)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
								continue
							}
							µbody = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							πTemp002[0] = µbody
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label28
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							πTemp002[0] = µbody
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label29
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µname, πg.NewStr("last-modified").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label30
							}
							goto Label31
							// line 86: if len(body) > 1:
							πF.SetLineno(86)
						Label28:
							πTemp002 = πF.MakeArgs(1)
							πTemp010 = πg.KWArgs{
								{"level", πg.NewInt(1).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µfield, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("PEP header field body contains multiple elements:\n%s").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 87: raise DataError('PEP header field body contains multiple '
							πF.SetLineno(87)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label32
							// line 89: elif len(body) == 1:
							πF.SetLineno(89)
						Label29:
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µbody, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label33
							}
							goto Label34
							// line 90: if not isinstance(body[0], nodes.paragraph):
							πF.SetLineno(90)
						Label33:
							πTemp002 = πF.MakeArgs(1)
							πTemp010 = πg.KWArgs{
								{"level", πg.NewInt(1).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µfield, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("PEP header field body may only contain a single paragraph:\n%s").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßDataError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 91: raise DataError('PEP header field body may only contain '
							πF.SetLineno(91)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label34
						Label34:
							goto Label32
							// line 94: elif name == 'last-modified':
							πF.SetLineno(94)
						Label30:
							// line 95: date = time.strftime(
							πF.SetLineno(95)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("%d-%b-%Y").ToObject()
							πTemp014 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(8).ToObject()
							πTemp015 = πF.MakeArgs(1)
							πTemp006 = ßsource.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp006); πE != nil {
								continue
							}
							πTemp015[0] = πTemp011
							if πTemp006, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp006, ßstat, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp011.Call(πF, πTemp015, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp015)
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp014[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßlocaltime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstrftime, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdate = πTemp003
							if πE = πg.CheckLocal(πF, µcvs_url, "cvs_url"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µcvs_url); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label35
							}
							goto Label36
							// line 98: if cvs_url:
							πF.SetLineno(98)
						Label35:
							// line 99: body += nodes.paragraph(
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = ß.ToObject()
							πTemp002[1] = ß.ToObject()
							πTemp014 = πF.MakeArgs(2)
							πTemp014[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µdate, "date"); πE != nil {
								continue
							}
							πTemp014[1] = µdate
							if πE = πg.CheckLocal(πF, µcvs_url, "cvs_url"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"refuri", µcvs_url},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp014, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[2] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßparagraph, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IAdd(πF, µbody, πTemp003); πE != nil {
								continue
							}
							µbody = πTemp004
							goto Label36
						Label36:
							goto Label32
						Label31:
							// line 103: continue
							πF.SetLineno(103)
							continue
							goto Label32
						Label32:
							// line 104: para = body[0]
							πF.SetLineno(104)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µbody, πTemp003); πE != nil {
								continue
							}
							µpara = πTemp004
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µname, ßauthor.ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label37
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µname, πg.NewStr("discussions-to").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label38
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(ßreplaces.ToObject(), πg.NewStr("replaced-by").ToObject(), ßrequires.ToObject()).ToObject()
							if πTemp007, πE = πg.Contains(πF, πTemp004, µname); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label39
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µname, πg.NewStr("last-modified").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label40
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µname, πg.NewStr("content-type").ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label41
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µname, ßversion.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label42
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							πTemp002[0] = µbody
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp006
						Label42:
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label43
							}
							goto Label44
							// line 105: if name == 'author':
							πF.SetLineno(105)
						Label37:
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µpara); πE != nil {
								continue
							}
							πF.PushCheckpoint(46)
							πTemp007 = false
						Label45:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label47
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp016 = !isStop
							} else {
								πTemp016 = true
								µnode = πTemp004
							}
							if πE != nil || !πTemp016 {
								continue
							}
							πF.PushCheckpoint(45)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßreference, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp016, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp016 {
								goto Label48
							}
							goto Label49
							// line 107: if isinstance(node, nodes.reference):
							πF.SetLineno(107)
						Label48:
							// line 108: node.replace_self(mask_email(node))
							πF.SetLineno(108)
							πTemp002 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp014[0] = µnode
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmask_email); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[0] = πTemp006
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label49
						Label49:
							continue
						Label46:
							if πE != nil || πR != nil {
								continue
							}
						Label47:
							goto Label44
							// line 109: elif name == 'discussions-to':
							πF.SetLineno(109)
						Label38:
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µpara); πE != nil {
								continue
							}
							πF.PushCheckpoint(51)
							πTemp007 = false
						Label50:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label52
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp016 = !isStop
							} else {
								πTemp016 = true
								µnode = πTemp004
							}
							if πE != nil || !πTemp016 {
								continue
							}
							πF.PushCheckpoint(50)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßreference, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp016, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp016 {
								goto Label53
							}
							goto Label54
							// line 111: if isinstance(node, nodes.reference):
							πF.SetLineno(111)
						Label53:
							// line 112: node.replace_self(mask_email(node, pep))
							πF.SetLineno(112)
							πTemp002 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp014[0] = µnode
							if πE = πg.CheckLocal(πF, µpep, "pep"); πE != nil {
								continue
							}
							πTemp014[1] = µpep
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmask_email); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[0] = πTemp006
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnode, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label54
						Label54:
							continue
						Label51:
							if πE != nil || πR != nil {
								continue
							}
						Label52:
							goto Label44
							// line 113: elif name in ('replaces', 'replaced-by', 'requires'):
							πF.SetLineno(113)
						Label39:
							// line 114: newbody = []
							πF.SetLineno(114)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							µnewbody = πTemp003
							// line 115: space = nodes.Text(' ')
							πF.SetLineno(115)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(" ").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µspace = πTemp003
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr(",?\\s+").ToObject()
							if πE = πg.CheckLocal(πF, µbody, "body"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µbody, ßastext, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp006
							if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(56)
							πTemp007 = false
						Label55:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label57
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp016 = !isStop
							} else {
								πTemp016 = true
								µrefpep = πTemp004
							}
							if πE != nil || !πTemp016 {
								continue
							}
							πF.PushCheckpoint(55)
							// line 117: pepno = int(refpep)
							πF.SetLineno(117)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefpep, "refpep"); πE != nil {
								continue
							}
							πTemp002[0] = µrefpep
							if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µpepno = πTemp006
							// line 118: newbody.append(nodes.reference(
							πF.SetLineno(118)
							πTemp002 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrefpep, "refpep"); πE != nil {
								continue
							}
							πTemp014[0] = µrefpep
							if πE = πg.CheckLocal(πF, µrefpep, "refpep"); πE != nil {
								continue
							}
							πTemp014[1] = µrefpep
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp006, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp011, ßpep_base_url, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßpep_url, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpepno, "pepno"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Mod(πF, πTemp012, µpepno); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp006, πTemp011); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"refuri", πTemp004},
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßreference, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp014, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µnewbody, "newbody"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnewbody, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 122: newbody.append(space)
							πF.SetLineno(122)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µspace, "space"); πE != nil {
								continue
							}
							πTemp002[0] = µspace
							if πE = πg.CheckLocal(πF, µnewbody, "newbody"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnewbody, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label56:
							if πE != nil || πR != nil {
								continue
							}
						Label57:
							// line 123: para[:] = newbody[:-1] # drop trailing space
							πF.SetLineno(123)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewbody, "newbody"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µnewbody, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µpara, πTemp006, πTemp003); πE != nil {
								continue
							}
							goto Label44
							// line 124: elif name == 'last-modified':
							πF.SetLineno(124)
						Label40:
							// line 125: utils.clean_rcs_keywords(para, self.rcs_keyword_substitutions)
							πF.SetLineno(125)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							πTemp002[0] = µpara
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrcs_keyword_substitutions, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclean_rcs_keywords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µcvs_url, "cvs_url"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µcvs_url); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label58
							}
							goto Label59
							// line 126: if cvs_url:
							πF.SetLineno(126)
						Label58:
							// line 127: date = para.astext()
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpara, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdate = πTemp004
							// line 128: para[:] = [nodes.reference('', date, refuri=cvs_url)]
							πF.SetLineno(128)
							πTemp002 = make([]*πg.Object, 1)
							πTemp014 = πF.MakeArgs(2)
							πTemp014[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µdate, "date"); πE != nil {
								continue
							}
							πTemp014[1] = µdate
							if πE = πg.CheckLocal(πF, µcvs_url, "cvs_url"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"refuri", µcvs_url},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp014, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µpara, πTemp006, πTemp004); πE != nil {
								continue
							}
							goto Label59
						Label59:
							goto Label44
							// line 129: elif name == 'content-type':
							πF.SetLineno(129)
						Label41:
							// line 130: pep_type = para.astext()
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpara, ßastext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µpep_type = πTemp004
							// line 131: uri = self.document.settings.pep_base_url + self.pep_url % 12
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßpep_base_url, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßpep_url, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mod(πF, πTemp011, πg.NewInt(12).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							µuri = πTemp003
							// line 132: para[:] = [nodes.reference('', pep_type, refuri=uri)]
							πF.SetLineno(132)
							πTemp002 = make([]*πg.Object, 1)
							πTemp014 = πF.MakeArgs(2)
							πTemp014[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µpep_type, "pep_type"); πE != nil {
								continue
							}
							πTemp014[1] = µpep_type
							if πE = πg.CheckLocal(πF, µuri, "uri"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"refuri", µuri},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp014, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, µpara, πTemp006, πTemp004); πE != nil {
								continue
							}
							goto Label44
							// line 133: elif name == 'version' and len(body):
							πF.SetLineno(133)
						Label43:
							// line 134: utils.clean_rcs_keywords(para, self.rcs_keyword_substitutions)
							πF.SetLineno(134)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpara, "para"); πE != nil {
								continue
							}
							πTemp002[0] = µpara
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrcs_keyword_substitutions, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclean_rcs_keywords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label44
						Label44:
							continue
						Label26:
							if πE != nil || πR != nil {
								continue
							}
						Label27:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Headers").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHeaders.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 137: class Contents(Transform):
			πF.SetLineno(137)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Contents", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 139: """
					πF.SetLineno(139)
					// line 139: """
					πF.SetLineno(139)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Insert an empty table of contents topic and a transform placeholder into\n    the document after the RFC 2822 header.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 144: default_priority = 380
					πF.SetLineno(144)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(380).ToObject()); πE != nil {
						continue
					}
					// line 146: def apply(self):
					πF.SetLineno(146)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlanguage *πg.Object = πg.UnboundLocal
						_ = µlanguage
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µtopic *πg.Object = πg.UnboundLocal
						_ = µtopic
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
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
							// line 147: language = languages.get_language(self.document.settings.language_code,
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlanguage_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreporter, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlanguages); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_language, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlanguage = πTemp002
							// line 149: name = language.labels['contents']
							πF.SetLineno(149)
							πTemp002 = ßcontents.ToObject()
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µlanguage, ßlabels, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µname = πTemp003
							// line 150: title = nodes.title('', name)
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtitle = πTemp002
							// line 151: topic = nodes.topic('', title, classes=['contents'])
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp001[1] = µtitle
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = ßcontents.ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp006 = πg.KWArgs{
								{"classes", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtopic, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtopic = πTemp002
							// line 152: name = nodes.fully_normalize_name(name)
							πF.SetLineno(152)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfully_normalize_name, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µname = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßhas_name, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 153: if not self.document.has_name(name):
							πF.SetLineno(153)
						Label1:
							// line 154: topic['names'].append(name)
							πF.SetLineno(154)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[0] = µname
							πTemp002 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µtopic, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							// line 155: self.document.note_implicit_target(topic)
							πF.SetLineno(155)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							πTemp001[0] = µtopic
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_implicit_target, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 156: pending = nodes.pending(parts.Contents)
							πF.SetLineno(156)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßparts); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßContents, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpending, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µpending = πTemp002
							// line 157: topic += pending
							πF.SetLineno(157)
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µtopic, µpending); πE != nil {
								continue
							}
							µtopic = πTemp002
							// line 158: self.document.insert(1, topic)
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µtopic, "topic"); πE != nil {
								continue
							}
							πTemp001[1] = µtopic
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 159: self.document.note_pending(pending)
							πF.SetLineno(159)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp001[0] = µpending
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Contents").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßContents.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 162: class TargetNotes(Transform):
			πF.SetLineno(162)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TargetNotes", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 164: """
					πF.SetLineno(164)
					// line 164: """
					πF.SetLineno(164)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Locate the \"References\" section, insert a placeholder for an external\n    target footnote insertion transform at the end, and schedule the\n    transform to run immediately.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 170: default_priority = 520
					πF.SetLineno(170)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(520).ToObject()); πE != nil {
						continue
					}
					// line 172: def apply(self):
					πF.SetLineno(172)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdoc *πg.Object = πg.UnboundLocal
						_ = µdoc
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µrefsect *πg.Object = πg.UnboundLocal
						_ = µrefsect
						var µcopyright *πg.Object = πg.UnboundLocal
						_ = µcopyright
						var µtitle_words *πg.Object = πg.UnboundLocal
						_ = µtitle_words
						var µpending *πg.Object = πg.UnboundLocal
						_ = µpending
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Dict
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 173: doc = self.document
							πF.SetLineno(173)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							µdoc = πTemp001
							// line 174: i = len(doc) - 1
							πF.SetLineno(174)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							πTemp002[0] = µdoc
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							// line 175: refsect = copyright = None
							πF.SetLineno(175)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µrefsect = πTemp001
							µcopyright = πTemp001
							// line 176: while i >= 0 and isinstance(doc[i], nodes.section):
							πF.SetLineno(176)
							πF.PushCheckpoint(2)
							πTemp005 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label4
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003 = µi
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µdoc, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsection, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label4:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 177: title_words = doc[i][0].astext().lower().split()
							πF.SetLineno(177)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdoc, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßastext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtitle_words = πTemp003
							if πE = πg.CheckLocal(πF, µtitle_words, "title_words"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µtitle_words, ßreferences.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µtitle_words, "title_words"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µtitle_words, ßcopyright.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 178: if 'references' in title_words:
							πF.SetLineno(178)
						Label5:
							// line 179: refsect = doc[i]
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µdoc, πTemp001); πE != nil {
								continue
							}
							µrefsect = πTemp003
							// line 180: break
							πF.SetLineno(180)
							πTemp005 = true
							continue
							goto Label7
							// line 181: elif 'copyright' in title_words:
							πF.SetLineno(181)
						Label6:
							// line 182: copyright = i
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							µcopyright = µi
							goto Label7
						Label7:
							// line 183: i -= 1
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µrefsect); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							goto Label9
							// line 184: if not refsect:
							πF.SetLineno(184)
						Label8:
							// line 185: refsect = nodes.section()
							πF.SetLineno(185)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsection, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrefsect = πTemp001
							// line 186: refsect += nodes.title('', 'References')
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ß.ToObject()
							πTemp002[1] = ßReferences.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßtitle, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.IAdd(πF, µrefsect, πTemp001); πE != nil {
								continue
							}
							µrefsect = πTemp003
							// line 187: doc.set_id(refsect)
							πF.SetLineno(187)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							πTemp002[0] = µrefsect
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdoc, ßset_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µcopyright, "copyright"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µcopyright); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 188: if copyright:
							πF.SetLineno(188)
						Label10:
							// line 190: doc.insert(copyright, refsect)
							πF.SetLineno(190)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcopyright, "copyright"); πE != nil {
								continue
							}
							πTemp002[0] = µcopyright
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							πTemp002[1] = µrefsect
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdoc, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label12
						Label11:
							// line 193: doc.append(refsect)
							πF.SetLineno(193)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							πTemp002[0] = µrefsect
							if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdoc, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label12
						Label12:
							goto Label9
						Label9:
							// line 194: pending = nodes.pending(references.TargetNotes)
							πF.SetLineno(194)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßreferences); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTargetNotes, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µpending = πTemp001
							// line 195: refsect.append(pending)
							πF.SetLineno(195)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrefsect, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 196: self.document.note_pending(pending, 0)
							πF.SetLineno(196)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 197: pending = nodes.pending(misc.CallBack,
							πF.SetLineno(197)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmisc); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßCallBack, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp009 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcleanup_callback, nil); πE != nil {
								continue
							}
							if πE = πTemp009.SetItem(πF, ßcallback.ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp009.ToObject()
							πTemp010 = πg.KWArgs{
								{"details", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, πTemp010); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µpending = πTemp001
							// line 199: refsect.append(pending)
							πF.SetLineno(199)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							if πE = πg.CheckLocal(πF, µrefsect, "refsect"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrefsect, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 200: self.document.note_pending(pending, 1)
							πF.SetLineno(200)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							πTemp002[0] = µpending
							πTemp002[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnote_pending, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 202: def cleanup_callback(self, pending):
					πF.SetLineno(202)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pending", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("cleanup_callback", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpending *πg.Object = πArgs[1]
						_ = µpending
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
							// line 203: """
							πF.SetLineno(203)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 208: if len(pending.parent) == 2:    # <title> and <pending>
							πF.SetLineno(208)
						Label1:
							// line 209: pending.parent.parent.remove(pending.parent)
							πF.SetLineno(209)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µpending, "pending"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µpending, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßparent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcleanup_callback.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 203: """
					πF.SetLineno(203)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Remove an empty \"References\" section.\n\n        Called after the `references.TargetNotes` transform is complete.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßcleanup_callback); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TargetNotes").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTargetNotes.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 212: class PEPZero(Transform):
			πF.SetLineno(212)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTransform); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PEPZero", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 214: """
					πF.SetLineno(214)
					// line 214: """
					πF.SetLineno(214)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Special processing for PEP 0.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 218: default_priority =760
					πF.SetLineno(218)
					if πE = πClass.SetItem(πF, ßdefault_priority.ToObject(), πg.NewInt(760).ToObject()); πE != nil {
						continue
					}
					// line 220: def apply(self):
					πF.SetLineno(220)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("apply", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvisitor *πg.Object = πg.UnboundLocal
						_ = µvisitor
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
							// line 221: visitor = PEPZeroSpecial(self.document)
							πF.SetLineno(221)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPEPZeroSpecial); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvisitor = πTemp003
							// line 222: self.document.walk(visitor)
							πF.SetLineno(222)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvisitor, "visitor"); πE != nil {
								continue
							}
							πTemp001[0] = µvisitor
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwalk, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 223: self.startnode.parent.remove(self.startnode)
							πF.SetLineno(223)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstartnode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßapply.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PEPZero").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPEPZero.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 226: class PEPZeroSpecial(nodes.SparseNodeVisitor):
			πF.SetLineno(226)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßSparseNodeVisitor, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("PEPZeroSpecial", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
				var πTemp008 *πg.Object
				_ = πTemp008
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 228: """
					πF.SetLineno(228)
					// line 228: """
					πF.SetLineno(228)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Perform the special processing needed by PEP 0:\n\n    - Mask email addresses.\n\n    - Link PEP numbers in the second column of 4-column tables to the PEPs\n      themselves.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 237: pep_url = Headers.pep_url
					πF.SetLineno(237)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßHeaders); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpep_url, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßpep_url.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 239: def unknown_visit(self, node):
					πF.SetLineno(239)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("unknown_visit", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
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
							// line 240: pass
							πF.SetLineno(240)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßunknown_visit.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 242: def visit_reference(self, node):
					πF.SetLineno(242)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("visit_reference", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 243: node.replace_self(mask_email(node))
							πF.SetLineno(243)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp002[0] = µnode
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmask_email); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnode, ßreplace_self, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_reference.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 245: def visit_field_list(self, node):
					πF.SetLineno(245)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("visit_field_list", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							πTemp002 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, ßrfc2822.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 246: if 'rfc2822' in node['classes']:
							πF.SetLineno(246)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSkipNode, nil); πE != nil {
								continue
							}
							// line 247: raise nodes.SkipNode
							πF.SetLineno(247)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_field_list.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 249: def visit_tgroup(self, node):
					πF.SetLineno(249)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("visit_tgroup", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
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
							// line 250: self.pep_table = node['cols'] == 4
							πF.SetLineno(250)
							πTemp002 = ßcols.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µnode, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpep_table, πTemp002); πE != nil {
								continue
							}
							// line 251: self.entry = 0
							πF.SetLineno(251)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßentry, πTemp001); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_tgroup.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 253: def visit_colspec(self, node):
					πF.SetLineno(253)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("visit_colspec", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 254: self.entry += 1
							πF.SetLineno(254)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßentry, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßentry, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpep_table, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßentry, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 255: if self.pep_table and self.entry == 2:
							πF.SetLineno(255)
						Label2:
							// line 256: node['classes'].append('num')
							πF.SetLineno(256)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßnum.ToObject()
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßvisit_colspec.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 258: def visit_row(self, node):
					πF.SetLineno(258)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("visit_row", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 259: self.entry = 0
							πF.SetLineno(259)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßentry, πTemp001); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvisit_row.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 261: def visit_entry(self, node):
					πF.SetLineno(261)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "node", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("visit_entry", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µp *πg.Object = πg.UnboundLocal
						_ = µp
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
						var µpep *πg.Object = πg.UnboundLocal
						_ = µpep
						var µref *πg.Object = πg.UnboundLocal
						_ = µref
						var πTemp001 *πg.Object
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
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							default:
								panic("unexpected function state")
							}
							// line 262: self.entry += 1
							πF.SetLineno(262)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßentry, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßentry, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpep_table, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßentry, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp005[0] = µnode
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 263: if self.pep_table and self.entry == 2 and len(node) == 1:
							πF.SetLineno(263)
						Label2:
							// line 264: node['classes'].append('num')
							πF.SetLineno(264)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßnum.ToObject()
							πTemp001 = ßclasses.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 265: p = node[0]
							πF.SetLineno(265)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							µp = πTemp002
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp005[0] = µp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßparagraph, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp004
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label4
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp005[0] = µp
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label4:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 266: if isinstance(p, nodes.paragraph) and len(p) == 1:
							πF.SetLineno(266)
						Label5:
							// line 267: text = p.astext()
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßastext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtext = πTemp002
							// line 268: try:
							πF.SetLineno(268)
							πF.PushCheckpoint(8)
							// line 269: pep = int(text)
							πF.SetLineno(269)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp005[0] = µtext
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpep = πTemp002
							// line 270: ref = (self.document.settings.pep_base_url
							πF.SetLineno(270)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßpep_base_url, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpep_url, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpep, "pep"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πTemp006, µpep); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							µref = πTemp001
							// line 272: p[0] = nodes.reference(text, text, refuri=ref)
							πF.SetLineno(272)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp005[0] = µtext
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp005[1] = µtext
							if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"refuri", µref},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, µp, πTemp004, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 273: except ValueError:
							πF.SetLineno(273)
						Label9:
							// line 274: pass
							πF.SetLineno(274)
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ßvisit_entry.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PEPZeroSpecial").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPEPZeroSpecial.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 277: non_masked_addresses = ('peps@python.org',
			πF.SetLineno(277)
			πTemp001 = πg.NewTuple3(πg.NewStr("peps@python.org").ToObject(), πg.NewStr("python-list@python.org").ToObject(), πg.NewStr("python-dev@python.org").ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnon_masked_addresses.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 281: def mask_email(ref, pepno=None):
			πF.SetLineno(281)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "ref", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[1] = πg.Param{Name: "pepno", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("mask_email", "/usr/lib/python2.7/site-packages/docutils/transforms/peps.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µref *πg.Object = πArgs[0]
				_ = µref
				var µpepno *πg.Object = πArgs[1]
				_ = µpepno
				var µreplacement *πg.Object = πg.UnboundLocal
				_ = µreplacement
				var µreplacement_text *πg.Object = πg.UnboundLocal
				_ = µreplacement_text
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 πg.KWArgs
				_ = πTemp008
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
					// line 282: """
					πF.SetLineno(282)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßrefuri.ToObject()
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µref, ßhasattr, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("mailto:").ToObject()
					πTemp004 = ßrefuri.ToObject()
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µref, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp001 = πTemp005
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 292: if ref.hasattr('refuri') and ref['refuri'].startswith('mailto:'):
					πF.SetLineno(292)
				Label2:
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(8).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					πTemp006 = ßrefuri.ToObject()
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µref, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßnon_masked_addresses); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 293: if ref['refuri'][8:] in non_masked_addresses:
					πF.SetLineno(293)
				Label5:
					// line 294: replacement = ref[0]
					πF.SetLineno(294)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µref, πTemp001); πE != nil {
						continue
					}
					µreplacement = πTemp004
					goto Label7
				Label6:
					// line 296: replacement_text = ref.astext().replace('@', '&#32;&#97;t&#32;')
					πF.SetLineno(296)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr("@").ToObject()
					πTemp003[1] = πg.NewStr("&#32;&#97;t&#32;").ToObject()
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µref, ßastext, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µreplacement_text = πTemp004
					// line 297: replacement = nodes.raw('', replacement_text, format='html')
					πF.SetLineno(297)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µreplacement_text, "replacement_text"); πE != nil {
						continue
					}
					πTemp003[1] = µreplacement_text
					πTemp008 = πg.KWArgs{
						{"format", ßhtml.ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßraw, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µreplacement = πTemp001
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µpepno, "pepno"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µpepno == πTemp004).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 298: if pepno is None:
					πF.SetLineno(298)
				Label8:
					// line 299: return replacement
					πF.SetLineno(299)
					if πE = πg.CheckLocal(πF, µreplacement, "replacement"); πE != nil {
						continue
					}
					πR = µreplacement
					continue
					goto Label10
				Label9:
					// line 301: ref['refuri'] += '?subject=PEP%%20%s' % pepno
					πF.SetLineno(301)
					πTemp001 = ßrefuri.ToObject()
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µref, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpepno, "pepno"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("?subject=PEP%%20%s").ToObject(), µpepno); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					πTemp006 = ßrefuri.ToObject()
					if πE = πg.SetItem(πF, µref, πTemp006, πTemp005); πE != nil {
						continue
					}
					// line 302: ref[:] = [replacement]
					πF.SetLineno(302)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µreplacement, "replacement"); πE != nil {
						continue
					}
					πTemp003[0] = µreplacement
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, µref, πTemp005, πTemp004); πE != nil {
						continue
					}
					// line 303: return ref
					πF.SetLineno(303)
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					πR = µref
					continue
					goto Label10
				Label10:
					goto Label4
				Label3:
					// line 305: return ref
					πF.SetLineno(305)
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					πR = µref
					continue
					goto Label4
				Label4:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmask_email.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 282: """
			πF.SetLineno(282)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Mask the email address in `ref` and return a replacement node.\n\n    `ref` is returned unchanged if it contains no email address.\n\n    For email addresses such as \"user@host\", mask the address as \"user at\n    host\" (text) to thwart simple email address harvesters (except for those\n    listed in `non_masked_addresses`).  If a PEP number (`pepno`) is given,\n    return a reference including a default email subject.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmask_email); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("peps", Code)
}
