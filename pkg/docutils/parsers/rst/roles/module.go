package roles

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßCustomRole := πg.InternStr("CustomRole")
		ßDEFAULT_INTERPRETED_ROLE := πg.InternStr("DEFAULT_INTERPRETED_ROLE")
		ßGenericRole := πg.InternStr("GenericRole")
		ßKeyError := πg.InternStr("KeyError")
		ßLexer := πg.InternStr("Lexer")
		ßLexerError := πg.InternStr("LexerError")
		ßNone := πg.InternStr("None")
		ßText := πg.InternStr("Text")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß__call__ := πg.InternStr("__call__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_fallback_language_module := πg.InternStr("_fallback_language_module")
		ß_role_registry := πg.InternStr("_role_registry")
		ß_roles := πg.InternStr("_roles")
		ßabbreviation := πg.InternStr("abbreviation")
		ßacronym := πg.InternStr("acronym")
		ßappend := πg.InternStr("append")
		ßbase_role := πg.InternStr("base_role")
		ßclass := πg.InternStr("class")
		ßclass_option := πg.InternStr("class_option")
		ßclasses := πg.InternStr("classes")
		ßcode := πg.InternStr("code")
		ßcode_role := πg.InternStr("code_role")
		ßcontent := πg.InternStr("content")
		ßcopy := πg.InternStr("copy")
		ßdirectives := πg.InternStr("directives")
		ßdocument := πg.InternStr("document")
		ßemphasis := πg.InternStr("emphasis")
		ßerror := πg.InternStr("error")
		ßextend := πg.InternStr("extend")
		ßfind := πg.InternStr("find")
		ßformat := πg.InternStr("format")
		ßgeneric_custom_role := πg.InternStr("generic_custom_role")
		ßget := πg.InternStr("get")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßhasattr := πg.InternStr("hasattr")
		ßindex := πg.InternStr("index")
		ßinfo := πg.InternStr("info")
		ßinline := πg.InternStr("inline")
		ßint := πg.InternStr("int")
		ßjoin := πg.InternStr("join")
		ßlanguage := πg.InternStr("language")
		ßline := πg.InternStr("line")
		ßlist := πg.InternStr("list")
		ßliteral := πg.InternStr("literal")
		ßlower := πg.InternStr("lower")
		ßmath := πg.InternStr("math")
		ßmath_role := πg.InternStr("math_role")
		ßname := πg.InternStr("name")
		ßnode_class := πg.InternStr("node_class")
		ßnodes := πg.InternStr("nodes")
		ßobject := πg.InternStr("object")
		ßoptions := πg.InternStr("options")
		ßpep_base_url := πg.InternStr("pep_base_url")
		ßpep_file_url_template := πg.InternStr("pep_file_url_template")
		ßpep_reference_role := πg.InternStr("pep_reference_role")
		ßproblematic := πg.InternStr("problematic")
		ßraw := πg.InternStr("raw")
		ßraw_enabled := πg.InternStr("raw_enabled")
		ßraw_role := πg.InternStr("raw_role")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreference := πg.InternStr("reference")
		ßregister_canonical_role := πg.InternStr("register_canonical_role")
		ßregister_generic_role := πg.InternStr("register_generic_role")
		ßregister_local_role := πg.InternStr("register_local_role")
		ßreporter := πg.InternStr("reporter")
		ßrfc_base_url := πg.InternStr("rfc_base_url")
		ßrfc_reference_role := πg.InternStr("rfc_reference_role")
		ßrfc_url := πg.InternStr("rfc_url")
		ßrole := πg.InternStr("role")
		ßroles := πg.InternStr("roles")
		ßset_classes := πg.InternStr("set_classes")
		ßset_implicit_options := πg.InternStr("set_implicit_options")
		ßsettings := πg.InternStr("settings")
		ßsource := πg.InternStr("source")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßstrong := πg.InternStr("strong")
		ßsubscript := πg.InternStr("subscript")
		ßsuperscript := πg.InternStr("superscript")
		ßsupplied_content := πg.InternStr("supplied_content")
		ßsupplied_options := πg.InternStr("supplied_options")
		ßsyntax_highlight := πg.InternStr("syntax_highlight")
		ßtarget := πg.InternStr("target")
		ßtitle_reference := πg.InternStr("title_reference")
		ßunchanged := πg.InternStr("unchanged")
		ßunescape := πg.InternStr("unescape")
		ßunimplemented_role := πg.InternStr("unimplemented_role")
		ßupdate := πg.InternStr("update")
		ßutils := πg.InternStr("utils")
		ßwarning := πg.InternStr("warning")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 []πg.Param
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
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis module defines standard interpreted text role functions, a registry for\ninterpreted text roles, and an API for adding to and retrieving from the\nregistry.\n\nThe interface for interpreted role functions is as follows::\n\n    def role_fn(name, rawtext, text, lineno, inliner,\n                options={}, content=[]):\n        code...\n\n    # Set function attributes for customization:\n    role_fn.options = ...\n    role_fn.content = ...\n\nParameters:\n\n- ``name`` is the local name of the interpreted text role, the role name\n  actually used in the document.\n\n- ``rawtext`` is a string containing the entire interpreted text construct.\n  Return it as a ``problematic`` node linked to a system message if there is a\n  problem.\n\n- ``text`` is the interpreted text content, with backslash escapes converted\n  to nulls (``\x00``).\n\n- ``lineno`` is the line number where the interpreted text beings.\n\n- ``inliner`` is the Inliner object that called the role function.\n  It defines the following useful attributes: ``reporter``,\n  ``problematic``, ``memo``, ``parent``, ``document``.\n\n- ``options``: A dictionary of directive options for customization, to be\n  interpreted by the role function.  Used for additional attributes for the\n  generated elements and other functionality.\n\n- ``content``: A list of strings, the directive content for customization\n  (\"role\" directive).  To be interpreted by the role function.\n\nFunction attributes for customization, interpreted by the \"role\" directive:\n\n- ``options``: A dictionary, mapping known option names to conversion\n  functions such as `int` or `float`.  ``None`` or an empty dict implies no\n  options to parse.  Several directive option conversion functions are defined\n  in the `directives` module.\n\n  All role functions implicitly support the \"class\" option, unless disabled\n  with an explicit ``{'class': None}``.\n\n- ``content``: A boolean; true if content is allowed.  Client code must handle\n  the case where content is required but not supplied (an empty content list\n  will be supplied).\n\nNote that unlike directives, the \"arguments\" function attribute is not\nsupported for role customization.  Directive arguments are handled by the\n\"role\" directive itself.\n\nInterpreted role functions return a tuple of two values:\n\n- A list of nodes which will be inserted into the document tree at the\n  point where the interpreted role was encountered (can be an empty\n  list).\n\n- A list of system messages, which will be inserted into the document tree\n  immediately after the end of the current inline block (can also be empty).\n").ToObject()); πE != nil {
				continue
			}
			// line 73: __docformat__ = 'reStructuredText'
			πF.SetLineno(73)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 75: from docutils import nodes, utils
			πF.SetLineno(75)
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
			// line 76: from docutils.parsers.rst import directives
			πF.SetLineno(76)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.directives"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßdirectives.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 77: from docutils.parsers.rst.languages import en as _fallback_language_module
			πF.SetLineno(77)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.languages.en"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[4]
			if πE = πF.Globals().SetItem(πF, ß_fallback_language_module.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 78: from docutils.utils.code_analyzer import Lexer, LexerError
			πF.SetLineno(78)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.code_analyzer"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLexer); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLexer.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßLexerError); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLexerError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 80: DEFAULT_INTERPRETED_ROLE = 'title-reference'
			πF.SetLineno(80)
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_INTERPRETED_ROLE.ToObject(), πg.NewStr("title-reference").ToObject()); πE != nil {
				continue
			}
			// line 81: """
			πF.SetLineno(81)
			// line 86: _role_registry = {}
			πF.SetLineno(86)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_role_registry.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 87: """Mapping of canonical role names to role functions.  Language-dependent role
			πF.SetLineno(87)
			// line 90: _roles = {}
			πF.SetLineno(90)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_roles.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 91: """Mapping of local or language-dependent interpreted text role names to role
			πF.SetLineno(91)
			// line 94: def role(role_name, language_module, lineno, reporter):
			πF.SetLineno(94)
			πTemp005 = make([]πg.Param, 4)
			πTemp005[0] = πg.Param{Name: "role_name", Def: nil}
			πTemp005[1] = πg.Param{Name: "language_module", Def: nil}
			πTemp005[2] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[3] = πg.Param{Name: "reporter", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole_name *πg.Object = πArgs[0]
				_ = µrole_name
				var µlanguage_module *πg.Object = πArgs[1]
				_ = µlanguage_module
				var µlineno *πg.Object = πArgs[2]
				_ = µlineno
				var µreporter *πg.Object = πArgs[3]
				_ = µreporter
				var µnormname *πg.Object = πg.UnboundLocal
				_ = µnormname
				var µmessages *πg.Object = πg.UnboundLocal
				_ = µmessages
				var µmsg_text *πg.Object = πg.UnboundLocal
				_ = µmsg_text
				var µcanonicalname *πg.Object = πg.UnboundLocal
				_ = µcanonicalname
				var µerror *πg.Object = πg.UnboundLocal
				_ = µerror
				var µmessage *πg.Object = πg.UnboundLocal
				_ = µmessage
				var µrole_fn *πg.Object = πg.UnboundLocal
				_ = µrole_fn
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 []*πg.Object
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
					case 13:
						goto Label13
					case 7:
						goto Label7
					default:
						panic("unexpected function state")
					}
					// line 95: """
					πF.SetLineno(95)
					// line 101: normname = role_name.lower()
					πF.SetLineno(101)
					if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µrole_name, ßlower, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnormname = πTemp002
					// line 102: messages = []
					πF.SetLineno(102)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µmessages = πTemp001
					// line 103: msg_text = []
					πF.SetLineno(103)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µmsg_text = πTemp001
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_roles); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp002, µnormname); πE != nil {
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
					// line 105: if normname in _roles:
					πF.SetLineno(105)
				Label1:
					// line 106: return _roles[normname], messages
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp002 = µnormname
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_roles); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp005, µmessages).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µrole_name); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 108: if role_name:
					πF.SetLineno(108)
				Label3:
					// line 109: canonicalname = None
					πF.SetLineno(109)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µcanonicalname = πTemp001
					// line 110: try:
					πF.SetLineno(110)
					πF.PushCheckpoint(7)
					// line 111: canonicalname = language_module.roles[normname]
					πF.SetLineno(111)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp001 = µnormname
					if πE = πg.CheckLocal(πF, µlanguage_module, "language_module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlanguage_module, ßroles, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					µcanonicalname = πTemp002
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 112: except AttributeError as error:
					πF.SetLineno(112)
				Label8:
					µerror = πTemp007.ToObject()
					// line 113: msg_text.append('Problem retrieving role entry from language '
					πF.SetLineno(113)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage_module, "language_module"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µlanguage_module, µerror).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Problem retrieving role entry from language module %r: %s.").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.RestoreExc(nil, nil)
					goto Label6
					// line 115: except KeyError:
					πF.SetLineno(115)
				Label9:
					// line 116: msg_text.append('No role entry for "%s" in module "%s".'
					πF.SetLineno(116)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlanguage_module, "language_module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlanguage_module, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µrole_name, πTemp005).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("No role entry for \"%s\" in module \"%s\".").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					goto Label5
				Label4:
					// line 119: canonicalname = DEFAULT_INTERPRETED_ROLE
					πF.SetLineno(119)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßDEFAULT_INTERPRETED_ROLE); πE != nil {
						continue
					}
					µcanonicalname = πTemp001
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µcanonicalname, "canonicalname"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcanonicalname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					goto Label11
					// line 122: if not canonicalname:
					πF.SetLineno(122)
				Label10:
					// line 123: try:
					πF.SetLineno(123)
					πF.PushCheckpoint(13)
					// line 124: canonicalname = _fallback_language_module.roles[normname]
					πF.SetLineno(124)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp001 = µnormname
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_fallback_language_module); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßroles, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
						continue
					}
					µcanonicalname = πTemp002
					// line 125: msg_text.append('Using English fallback for role "%s".'
					πF.SetLineno(125)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Using English fallback for role \"%s\".").ToObject(), µrole_name); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πF.PopCheckpoint()
					goto Label12
				Label13:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label14
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 127: except KeyError:
					πF.SetLineno(127)
				Label14:
					// line 128: msg_text.append('Trying "%s" as canonical role name.'
					πF.SetLineno(128)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Trying \"%s\" as canonical role name.").ToObject(), µrole_name); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmsg_text, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 131: canonicalname = normname
					πF.SetLineno(131)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					µcanonicalname = µnormname
					πF.RestoreExc(nil, nil)
					goto Label12
				Label12:
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µmsg_text); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 134: if msg_text:
					πF.SetLineno(134)
				Label15:
					// line 135: message = reporter.info('\n'.join(msg_text), line=lineno)
					πF.SetLineno(135)
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg_text, "msg_text"); πE != nil {
						continue
					}
					πTemp009[0] = µmsg_text
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"line", µlineno},
					}
					if πE = πg.CheckLocal(πF, µreporter, "reporter"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µreporter, ßinfo, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µmessage = πTemp002
					// line 136: messages.append(message)
					πF.SetLineno(136)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp003[0] = µmessage
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmessages, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label16
				Label16:
					if πE = πg.CheckLocal(πF, µcanonicalname, "canonicalname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_role_registry); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp002, µcanonicalname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label17
					}
					goto Label18
					// line 139: if canonicalname in _role_registry:
					πF.SetLineno(139)
				Label17:
					// line 140: role_fn = _role_registry[canonicalname]
					πF.SetLineno(140)
					if πE = πg.CheckLocal(πF, µcanonicalname, "canonicalname"); πE != nil {
						continue
					}
					πTemp001 = µcanonicalname
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_role_registry); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					µrole_fn = πTemp002
					// line 141: register_local_role(normname, role_fn)
					πF.SetLineno(141)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µnormname, "normname"); πE != nil {
						continue
					}
					πTemp003[0] = µnormname
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					πTemp003[1] = µrole_fn
					if πTemp001, πE = πg.ResolveGlobal(πF, ßregister_local_role); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 142: return role_fn, messages
					πF.SetLineno(142)
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µrole_fn, µmessages).ToObject()
					πR = πTemp001
					continue
					goto Label19
				Label18:
					// line 144: return None, messages # Error message will be generated by caller.
					πF.SetLineno(144)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmessages, "messages"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp002, µmessages).ToObject()
					πR = πTemp001
					continue
					goto Label19
				Label19:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßrole.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 95: """
			πF.SetLineno(95)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Locate and return a role function from its language-dependent name, along\n    with a list of system messages.  If the role is not found in the current\n    language, check English.  Return a 2-tuple: role function (``None`` if the\n    named role cannot be found) and a list of system messages.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßrole); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 146: def register_canonical_role(name, role_fn):
			πF.SetLineno(146)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "name", Def: nil}
			πTemp005[1] = πg.Param{Name: "role_fn", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("register_canonical_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
				var µrole_fn *πg.Object = πArgs[1]
				_ = µrole_fn
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 147: """
					πF.SetLineno(147)
					// line 154: set_implicit_options(role_fn)
					πF.SetLineno(154)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					πTemp001[0] = µrole_fn
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset_implicit_options); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 155: _role_registry[name] = role_fn
					πF.SetLineno(155)
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µrole_fn); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_role_registry); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = µname
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßregister_canonical_role.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 147: """
			πF.SetLineno(147)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n    Register an interpreted text role by its canonical name.\n\n    :Parameters:\n      - `name`: The canonical name of the interpreted role.\n      - `role_fn`: The role function.  See the module docstring.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 157: def register_local_role(name, role_fn):
			πF.SetLineno(157)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "name", Def: nil}
			πTemp005[1] = πg.Param{Name: "role_fn", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("register_local_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
				var µrole_fn *πg.Object = πArgs[1]
				_ = µrole_fn
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 158: """
					πF.SetLineno(158)
					// line 165: set_implicit_options(role_fn)
					πF.SetLineno(165)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					πTemp001[0] = µrole_fn
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset_implicit_options); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 166: _roles[name] = role_fn
					πF.SetLineno(166)
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µrole_fn); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_roles); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = µname
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßregister_local_role.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 158: """
			πF.SetLineno(158)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n    Register an interpreted text role by its local or language-dependent name.\n\n    :Parameters:\n      - `name`: The local or language-dependent name of the interpreted role.\n      - `role_fn`: The role function.  See the module docstring.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßregister_local_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 168: def set_implicit_options(role_fn):
			πF.SetLineno(168)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "role_fn", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("set_implicit_options", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole_fn *πg.Object = πArgs[0]
				_ = µrole_fn
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Dict
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
					// line 169: """
					πF.SetLineno(169)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					πTemp004[0] = µrole_fn
					πTemp004[1] = ßoptions.ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µrole_fn, ßoptions, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005 == πTemp006).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µrole_fn, ßoptions, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, ßclass.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 173: if not hasattr(role_fn, 'options') or role_fn.options is None:
					πF.SetLineno(173)
				Label2:
					// line 174: role_fn.options = {'class': directives.class_option}
					πF.SetLineno(174)
					πTemp008 = πg.NewDict()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßclass.ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp008.ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µrole_fn, ßoptions, πTemp003); πE != nil {
						continue
					}
					goto Label4
					// line 175: elif 'class' not in role_fn.options:
					πF.SetLineno(175)
				Label3:
					// line 176: role_fn.options['class'] = directives.class_option
					πF.SetLineno(176)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclass_option, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrole_fn, "role_fn"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µrole_fn, ßoptions, nil); πE != nil {
						continue
					}
					πTemp006 = ßclass.ToObject()
					if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp001); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßset_implicit_options.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 169: """
			πF.SetLineno(169)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n    Add customization options to role functions, unless explicitly set or\n    disabled.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßset_implicit_options); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 178: def register_generic_role(canonical_name, node_class):
			πF.SetLineno(178)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "canonical_name", Def: nil}
			πTemp005[1] = πg.Param{Name: "node_class", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("register_generic_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcanonical_name *πg.Object = πArgs[0]
				_ = µcanonical_name
				var µnode_class *πg.Object = πArgs[1]
				_ = µnode_class
				var µrole *πg.Object = πg.UnboundLocal
				_ = µrole
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
					// line 179: """For roles which simply wrap a given `node_class` around the text."""
					πF.SetLineno(179)
					// line 180: role = GenericRole(canonical_name, node_class)
					πF.SetLineno(180)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcanonical_name, "canonical_name"); πE != nil {
						continue
					}
					πTemp001[0] = µcanonical_name
					if πE = πg.CheckLocal(πF, µnode_class, "node_class"); πE != nil {
						continue
					}
					πTemp001[1] = µnode_class
					if πTemp002, πE = πg.ResolveGlobal(πF, ßGenericRole); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrole = πTemp003
					// line 181: register_canonical_role(canonical_name, role)
					πF.SetLineno(181)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcanonical_name, "canonical_name"); πE != nil {
						continue
					}
					πTemp001[0] = µcanonical_name
					if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
						continue
					}
					πTemp001[1] = µrole
					if πTemp002, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßregister_generic_role.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 179: """For roles which simply wrap a given `node_class` around the text."""
			πF.SetLineno(179)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("For roles which simply wrap a given `node_class` around the text.").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 184: class GenericRole(object):
			πF.SetLineno(184)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp004 = πg.NewDict()
			if πTemp009, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp009); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GenericRole", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Dict
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 186: """
					πF.SetLineno(186)
					// line 186: """
					πF.SetLineno(186)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Generic interpreted text role, where the interpreted text is simply\n    wrapped with the provided node class.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 191: def __init__(self, role_name, node_class):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "role_name", Def: nil}
					πTemp002[2] = πg.Param{Name: "node_class", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrole_name *πg.Object = πArgs[1]
						_ = µrole_name
						var µnode_class *πg.Object = πArgs[2]
						_ = µnode_class
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
							// line 192: self.name = role_name
							πF.SetLineno(192)
							if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrole_name); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßname, πTemp001); πE != nil {
								continue
							}
							// line 193: self.node_class = node_class
							πF.SetLineno(193)
							if πE = πg.CheckLocal(πF, µnode_class, "node_class"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnode_class); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnode_class, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 195: def __call__(self, role, rawtext, text, lineno, inliner,
					πF.SetLineno(195)
					πTemp002 = make([]πg.Param, 8)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "role", Def: nil}
					πTemp002[2] = πg.Param{Name: "rawtext", Def: nil}
					πTemp002[3] = πg.Param{Name: "text", Def: nil}
					πTemp002[4] = πg.Param{Name: "lineno", Def: nil}
					πTemp002[5] = πg.Param{Name: "inliner", Def: nil}
					πTemp004 = πg.NewDict()
					πTemp005 = πTemp004.ToObject()
					πTemp002[6] = πg.Param{Name: "options", Def: πTemp005}
					πTemp006 = make([]*πg.Object, 0)
					πTemp005 = πg.NewList(πTemp006...).ToObject()
					πTemp002[7] = πg.Param{Name: "content", Def: πTemp005}
					πTemp003 = πg.NewFunction(πg.NewCode("__call__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrole *πg.Object = πArgs[1]
						_ = µrole
						var µrawtext *πg.Object = πArgs[2]
						_ = µrawtext
						var µtext *πg.Object = πArgs[3]
						_ = µtext
						var µlineno *πg.Object = πArgs[4]
						_ = µlineno
						var µinliner *πg.Object = πArgs[5]
						_ = µinliner
						var µoptions *πg.Object = πArgs[6]
						_ = µoptions
						var µcontent *πg.Object = πArgs[7]
						_ = µcontent
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
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
							// line 197: set_classes(options)
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							πTemp001[0] = µoptions
							if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 198: return [self.node_class(rawtext, text, **options)], []
							πF.SetLineno(198)
							πTemp001 = make([]*πg.Object, 1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							πTemp004[0] = µrawtext
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp004[1] = µtext
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnode_class, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, πTemp003, πTemp004, nil, nil, µoptions); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp005
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							πTemp001 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp001...).ToObject()
							πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp010, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp010 == nil {
				πTemp010 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp010.Call(πF, []*πg.Object{πg.NewStr("GenericRole").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGenericRole.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 201: class CustomRole(object):
			πF.SetLineno(201)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp004 = πg.NewDict()
			if πTemp009, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp009); πE != nil {
				continue
			}
			_, πE = πg.NewCode("CustomRole", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Dict
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 203: """
					πF.SetLineno(203)
					// line 203: """
					πF.SetLineno(203)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Wrapper for custom interpreted text roles.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 207: def __init__(self, role_name, base_role, options={}, content=[]):
					πF.SetLineno(207)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "role_name", Def: nil}
					πTemp002[2] = πg.Param{Name: "base_role", Def: nil}
					πTemp003 = πg.NewDict()
					πTemp004 = πTemp003.ToObject()
					πTemp002[3] = πg.Param{Name: "options", Def: πTemp004}
					πTemp005 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					πTemp002[4] = πg.Param{Name: "content", Def: πTemp004}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrole_name *πg.Object = πArgs[1]
						_ = µrole_name
						var µbase_role *πg.Object = πArgs[2]
						_ = µbase_role
						var µoptions *πg.Object = πArgs[3]
						_ = µoptions
						var µcontent *πg.Object = πArgs[4]
						_ = µcontent
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 208: self.name = role_name
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µrole_name, "role_name"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrole_name); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßname, πTemp001); πE != nil {
								continue
							}
							// line 209: self.base_role = base_role
							πF.SetLineno(209)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbase_role); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbase_role, πTemp001); πE != nil {
								continue
							}
							// line 210: self.options = None
							πF.SetLineno(210)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoptions, πTemp002); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							πTemp003[0] = µbase_role
							πTemp003[1] = ßoptions.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 211: if hasattr(base_role, 'options'):
							πF.SetLineno(211)
						Label1:
							// line 212: self.options = base_role.options
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbase_role, ßoptions, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoptions, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 213: self.content = None
							πF.SetLineno(213)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontent, πTemp002); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							πTemp003[0] = µbase_role
							πTemp003[1] = ßcontent.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 214: if hasattr(base_role, 'content'):
							πF.SetLineno(214)
						Label3:
							// line 215: self.content = base_role.content
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µbase_role, "base_role"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µbase_role, ßcontent, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontent, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 216: self.supplied_options = options
							πF.SetLineno(216)
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µoptions); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsupplied_options, πTemp001); πE != nil {
								continue
							}
							// line 217: self.supplied_content = content
							πF.SetLineno(217)
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcontent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsupplied_content, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 219: def __call__(self, role, rawtext, text, lineno, inliner,
					πF.SetLineno(219)
					πTemp002 = make([]πg.Param, 8)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "role", Def: nil}
					πTemp002[2] = πg.Param{Name: "rawtext", Def: nil}
					πTemp002[3] = πg.Param{Name: "text", Def: nil}
					πTemp002[4] = πg.Param{Name: "lineno", Def: nil}
					πTemp002[5] = πg.Param{Name: "inliner", Def: nil}
					πTemp003 = πg.NewDict()
					πTemp006 = πTemp003.ToObject()
					πTemp002[6] = πg.Param{Name: "options", Def: πTemp006}
					πTemp005 = make([]*πg.Object, 0)
					πTemp006 = πg.NewList(πTemp005...).ToObject()
					πTemp002[7] = πg.Param{Name: "content", Def: πTemp006}
					πTemp004 = πg.NewFunction(πg.NewCode("__call__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrole *πg.Object = πArgs[1]
						_ = µrole
						var µrawtext *πg.Object = πArgs[2]
						_ = µrawtext
						var µtext *πg.Object = πArgs[3]
						_ = µtext
						var µlineno *πg.Object = πArgs[4]
						_ = µlineno
						var µinliner *πg.Object = πArgs[5]
						_ = µinliner
						var µoptions *πg.Object = πArgs[6]
						_ = µoptions
						var µcontent *πg.Object = πArgs[7]
						_ = µcontent
						var µopts *πg.Object = πg.UnboundLocal
						_ = µopts
						var µcont *πg.Object = πg.UnboundLocal
						_ = µcont
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 221: opts = self.supplied_options.copy()
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsupplied_options, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µopts = πTemp001
							// line 222: opts.update(options)
							πF.SetLineno(222)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							πTemp003[0] = µoptions
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µopts, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 223: cont = list(self.supplied_content)
							πF.SetLineno(223)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsupplied_content, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µcont = πTemp002
							if πE = πg.CheckLocal(πF, µcont, "cont"); πE != nil {
								continue
							}
							πTemp001 = µcont
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							πTemp001 = µcontent
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 224: if cont and content:
							πF.SetLineno(224)
						Label2:
							// line 225: cont += '\n'
							πF.SetLineno(225)
							if πE = πg.CheckLocal(πF, µcont, "cont"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µcont, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							µcont = πTemp001
							goto Label3
						Label3:
							// line 226: cont.extend(content)
							πF.SetLineno(226)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							πTemp003[0] = µcontent
							if πE = πg.CheckLocal(πF, µcont, "cont"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcont, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 227: return self.base_role(role, rawtext, text, lineno, inliner,
							πF.SetLineno(227)
							πTemp003 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
								continue
							}
							πTemp003[0] = µrole
							if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
								continue
							}
							πTemp003[1] = µrawtext
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003[2] = µtext
							if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
								continue
							}
							πTemp003[3] = µlineno
							if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
								continue
							}
							πTemp003[4] = µinliner
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcont, "cont"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"options", µopts},
								{"content", µcont},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbase_role, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp010, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp010 == nil {
				πTemp010 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp010.Call(πF, []*πg.Object{πg.NewStr("CustomRole").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCustomRole.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 231: def generic_custom_role(role, rawtext, text, lineno, inliner,
			πF.SetLineno(231)
			πTemp005 = make([]πg.Param, 7)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp010 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "options", Def: πTemp010}
			πTemp002 = make([]*πg.Object, 0)
			πTemp010 = πg.NewList(πTemp002...).ToObject()
			πTemp005[6] = πg.Param{Name: "content", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("generic_custom_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µoptions *πg.Object = πArgs[5]
				_ = µoptions
				var µcontent *πg.Object = πArgs[6]
				_ = µcontent
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
					// line 233: """"""
					πF.SetLineno(233)
					// line 236: set_classes(options)
					πF.SetLineno(236)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp001[0] = µoptions
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 237: return [nodes.inline(rawtext, text, **options)], []
					πF.SetLineno(237)
					πTemp001 = make([]*πg.Object, 1)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp004[0] = µrawtext
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp004[1] = µtext
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßinline, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, πTemp004, nil, nil, µoptions); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[0] = πTemp003
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 0)
					πTemp005 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßgeneric_custom_role.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 233: """"""
			πF.SetLineno(233)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, ß.ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßgeneric_custom_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 239: generic_custom_role.options = {'class': directives.class_option}
			πF.SetLineno(239)
			πTemp004 = πg.NewDict()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßclass_option, nil); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßclass.ToObject(), πTemp011); πE != nil {
				continue
			}
			πTemp010 = πTemp004.ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp010); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßgeneric_custom_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ßoptions, πTemp011); πE != nil {
				continue
			}
			// line 246: register_generic_role('abbreviation', nodes.abbreviation)
			πF.SetLineno(246)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßabbreviation.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßabbreviation, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 247: register_generic_role('acronym', nodes.acronym)
			πF.SetLineno(247)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßacronym.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßacronym, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 248: register_generic_role('emphasis', nodes.emphasis)
			πF.SetLineno(248)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßemphasis.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßemphasis, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 249: register_generic_role('literal', nodes.literal)
			πF.SetLineno(249)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßliteral.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßliteral, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 250: register_generic_role('strong', nodes.strong)
			πF.SetLineno(250)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßstrong.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßstrong, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 251: register_generic_role('subscript', nodes.subscript)
			πF.SetLineno(251)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßsubscript.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßsubscript, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 252: register_generic_role('superscript', nodes.superscript)
			πF.SetLineno(252)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßsuperscript.ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßsuperscript, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 253: register_generic_role('title-reference', nodes.title_reference)
			πF.SetLineno(253)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("title-reference").ToObject()
			if πTemp010, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßtitle_reference, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp010, πE = πg.ResolveGlobal(πF, ßregister_generic_role); πE != nil {
				continue
			}
			if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 255: def pep_reference_role(role, rawtext, text, lineno, inliner,
			πF.SetLineno(255)
			πTemp005 = make([]πg.Param, 7)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp011 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "options", Def: πTemp011}
			πTemp002 = make([]*πg.Object, 0)
			πTemp011 = πg.NewList(πTemp002...).ToObject()
			πTemp005[6] = πg.Param{Name: "content", Def: πTemp011}
			πTemp010 = πg.NewFunction(πg.NewCode("pep_reference_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µoptions *πg.Object = πArgs[5]
				_ = µoptions
				var µcontent *πg.Object = πArgs[6]
				_ = µcontent
				var µpepnum *πg.Object = πg.UnboundLocal
				_ = µpepnum
				var µmsg *πg.Object = πg.UnboundLocal
				_ = µmsg
				var µprb *πg.Object = πg.UnboundLocal
				_ = µprb
				var µref *πg.Object = πg.UnboundLocal
				_ = µref
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 πg.KWArgs
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
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
					// line 257: try:
					πF.SetLineno(257)
					πF.PushCheckpoint(2)
					// line 258: pepnum = int(utils.unescape(text))
					πF.SetLineno(258)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp002[0] = µtext
					if πTemp003, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßunescape, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µpepnum = πTemp004
					if πE = πg.CheckLocal(πF, µpepnum, "pepnum"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LT(πF, µpepnum, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µpepnum, "pepnum"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GT(πF, µpepnum, πg.NewInt(9999).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label3:
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 259: if pepnum < 0 or pepnum > 9999:
					πF.SetLineno(259)
				Label4:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 260: raise ValueError
					πF.SetLineno(260)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label5
				Label5:
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 261: except ValueError:
					πF.SetLineno(261)
				Label6:
					// line 262: msg = inliner.reporter.error(
					πF.SetLineno(262)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("PEP number must be a number from 0 to 9999; \"%s\" is invalid.").ToObject(), µtext); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"line", µlineno},
					}
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp008); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmsg = πTemp003
					// line 265: prb = inliner.problematic(rawtext, rawtext, msg)
					πF.SetLineno(265)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[0] = µrawtext
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[1] = µrawtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[2] = µmsg
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µinliner, ßproblematic, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µprb = πTemp004
					// line 266: return [prb], [msg]
					πF.SetLineno(266)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
						continue
					}
					πTemp001[0] = µprb
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					πTemp009 = πg.NewList(πTemp001...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp004, πTemp009).ToObject()
					πR = πTemp003
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 268: ref = (inliner.document.settings.pep_base_url
					πF.SetLineno(268)
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µinliner, ßdocument, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp004, ßsettings, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp009, ßpep_base_url, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, µinliner, ßdocument, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßsettings, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp011, ßpep_file_url_template, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpepnum, "pepnum"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Mod(πF, πTemp010, µpepnum); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					µref = πTemp003
					// line 270: set_classes(options)
					πF.SetLineno(270)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp001[0] = µoptions
					if πTemp003, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 271: return [nodes.reference(rawtext, 'PEP ' + text, refuri=ref,
					πF.SetLineno(271)
					πTemp001 = make([]*πg.Object, 1)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp002[0] = µrawtext
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πg.NewStr("PEP ").ToObject(), µtext); πE != nil {
						continue
					}
					πTemp002[1] = πTemp004
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					πTemp008 = πg.KWArgs{
						{"refuri", µref},
					}
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp004, ßreference, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Invoke(πF, πTemp009, πTemp002, nil, πTemp008, µoptions); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 0)
					πTemp009 = πg.NewList(πTemp001...).ToObject()
					πTemp003 = πg.NewTuple2(πTemp004, πTemp009).ToObject()
					πR = πTemp003
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßpep_reference_role.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 274: register_canonical_role('pep-reference', pep_reference_role)
			πF.SetLineno(274)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("pep-reference").ToObject()
			if πTemp011, πE = πg.ResolveGlobal(πF, ßpep_reference_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp011
			if πTemp011, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 276: def rfc_reference_role(role, rawtext, text, lineno, inliner,
			πF.SetLineno(276)
			πTemp005 = make([]πg.Param, 7)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp012 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "options", Def: πTemp012}
			πTemp002 = make([]*πg.Object, 0)
			πTemp012 = πg.NewList(πTemp002...).ToObject()
			πTemp005[6] = πg.Param{Name: "content", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("rfc_reference_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µoptions *πg.Object = πArgs[5]
				_ = µoptions
				var µcontent *πg.Object = πArgs[6]
				_ = µcontent
				var µrfcnum *πg.Object = πg.UnboundLocal
				_ = µrfcnum
				var µsection *πg.Object = πg.UnboundLocal
				_ = µsection
				var µmsg *πg.Object = πg.UnboundLocal
				_ = µmsg
				var µprb *πg.Object = πg.UnboundLocal
				_ = µprb
				var µref *πg.Object = πg.UnboundLocal
				_ = µref
				var µnode *πg.Object = πg.UnboundLocal
				_ = µnode
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 πg.KWArgs
				_ = πTemp009
				var πTemp010 *πg.Object
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
					// line 278: try:
					πF.SetLineno(278)
					πF.PushCheckpoint(2)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µtext, πg.NewStr("#").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 279: if "#" in text:
					πF.SetLineno(279)
				Label3:
					// line 280: rfcnum, section = utils.unescape(text).split("#", 1)
					πF.SetLineno(280)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr("#").ToObject()
					πTemp003[1] = πg.NewInt(1).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp004[0] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßunescape, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µrfcnum = πTemp005
					µsection = πTemp006
					goto Label5
				Label4:
					// line 282: rfcnum, section  = utils.unescape(text), None
					πF.SetLineno(282)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp003[0] = µtext
					if πTemp005, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßunescape, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µrfcnum = πTemp005
					µsection = πTemp006
					goto Label5
				Label5:
					// line 283: rfcnum = int(rfcnum)
					πF.SetLineno(283)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrfcnum, "rfcnum"); πE != nil {
						continue
					}
					πTemp003[0] = µrfcnum
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µrfcnum = πTemp005
					if πE = πg.CheckLocal(πF, µrfcnum, "rfcnum"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µrfcnum, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 284: if rfcnum < 1:
					πF.SetLineno(284)
				Label6:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 285: raise ValueError
					πF.SetLineno(285)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label7
				Label7:
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 286: except ValueError:
					πF.SetLineno(286)
				Label8:
					// line 287: msg = inliner.reporter.error(
					πF.SetLineno(287)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("RFC number must be a number greater than or equal to 1; \"%s\" is invalid.").ToObject(), µtext); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"line", µlineno},
					}
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp005.Call(πF, πTemp003, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µmsg = πTemp001
					// line 290: prb = inliner.problematic(rawtext, rawtext, msg)
					πF.SetLineno(290)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp003[0] = µrawtext
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp003[1] = µrawtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp003[2] = µmsg
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßproblematic, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µprb = πTemp005
					// line 291: return [prb], [msg]
					πF.SetLineno(291)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
						continue
					}
					πTemp003[0] = µprb
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp003[0] = µmsg
					πTemp006 = πg.NewList(πTemp003...).ToObject()
					πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πR = πTemp001
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 293: ref = inliner.document.settings.rfc_base_url + inliner.rfc_url % rfcnum
					πF.SetLineno(293)
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µinliner, ßdocument, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettings, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßrfc_base_url, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, µinliner, ßrfc_url, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrfcnum, "rfcnum"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Mod(πF, πTemp010, µrfcnum); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					µref = πTemp001
					if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µsection != πTemp005).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 294: if section is not None:
					πF.SetLineno(294)
				Label9:
					// line 295: ref += "#"+section
					πF.SetLineno(295)
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("#").ToObject(), µsection); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IAdd(πF, µref, πTemp001); πE != nil {
						continue
					}
					µref = πTemp005
					goto Label10
				Label10:
					// line 296: set_classes(options)
					πF.SetLineno(296)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp003[0] = µoptions
					if πTemp001, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 297: node = nodes.reference(rawtext, 'RFC ' + str(rfcnum), refuri=ref,
					πF.SetLineno(297)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp003[0] = µrawtext
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrfcnum, "rfcnum"); πE != nil {
						continue
					}
					πTemp004[0] = µrfcnum
					if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Add(πF, πg.NewStr("RFC ").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"refuri", µref},
					}
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßreference, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, πTemp005, πTemp003, nil, πTemp009, µoptions); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µnode = πTemp001
					// line 299: return [node], []
					πF.SetLineno(299)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp003[0] = µnode
					πTemp005 = πg.NewList(πTemp003...).ToObject()
					πTemp003 = make([]*πg.Object, 0)
					πTemp006 = πg.NewList(πTemp003...).ToObject()
					πTemp001 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßrfc_reference_role.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 301: register_canonical_role('rfc-reference', rfc_reference_role)
			πF.SetLineno(301)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("rfc-reference").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßrfc_reference_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp012
			if πTemp012, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 303: def raw_role(role, rawtext, text, lineno, inliner, options={}, content=[]):
			πF.SetLineno(303)
			πTemp005 = make([]πg.Param, 7)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp013 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "options", Def: πTemp013}
			πTemp002 = make([]*πg.Object, 0)
			πTemp013 = πg.NewList(πTemp002...).ToObject()
			πTemp005[6] = πg.Param{Name: "content", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("raw_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µoptions *πg.Object = πArgs[5]
				_ = µoptions
				var µcontent *πg.Object = πArgs[6]
				_ = µcontent
				var µmsg *πg.Object = πg.UnboundLocal
				_ = µmsg
				var µprb *πg.Object = πg.UnboundLocal
				_ = µprb
				var µnode *πg.Object = πg.UnboundLocal
				_ = µnode
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
				var πTemp006 πg.KWArgs
				_ = πTemp006
				var πTemp007 []*πg.Object
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
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinliner, ßdocument, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßraw_enabled, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 304: if not inliner.document.settings.raw_enabled:
					πF.SetLineno(304)
				Label1:
					// line 305: msg = inliner.reporter.warning('raw (and derived) roles disabled')
					πF.SetLineno(305)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("raw (and derived) roles disabled").ToObject()
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwarning, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µmsg = πTemp001
					// line 306: prb = inliner.problematic(rawtext, rawtext, msg)
					πF.SetLineno(306)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp005[0] = µrawtext
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp005[1] = µrawtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp005[2] = µmsg
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßproblematic, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µprb = πTemp002
					// line 307: return [prb], [msg]
					πF.SetLineno(307)
					πTemp005 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
						continue
					}
					πTemp005[0] = µprb
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp005[0] = µmsg
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
					πR = πTemp001
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µoptions, ßformat.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 308: if 'format' not in options:
					πF.SetLineno(308)
				Label3:
					// line 309: msg = inliner.reporter.error(
					πF.SetLineno(309)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("No format (Writer name) is associated with this role: \"%s\".\nThe \"raw\" role cannot be used directly.\nInstead, use the \"role\" directive to create a new role with an associated format.").ToObject(), µrole); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp006 = πg.KWArgs{
						{"line", µlineno},
					}
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µmsg = πTemp001
					// line 314: prb = inliner.problematic(rawtext, rawtext, msg)
					πF.SetLineno(314)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp005[0] = µrawtext
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp005[1] = µrawtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp005[2] = µmsg
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßproblematic, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µprb = πTemp002
					// line 315: return [prb], [msg]
					πF.SetLineno(315)
					πTemp005 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
						continue
					}
					πTemp005[0] = µprb
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp005[0] = µmsg
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
					πR = πTemp001
					continue
					goto Label4
				Label4:
					// line 316: set_classes(options)
					πF.SetLineno(316)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp005[0] = µoptions
					if πTemp001, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 317: node = nodes.raw(rawtext, utils.unescape(text, True), **options)
					πF.SetLineno(317)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp005[0] = µrawtext
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp007[0] = µtext
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp007[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßunescape, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp005[1] = πTemp001
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßraw, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, πTemp002, πTemp005, nil, nil, µoptions); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µnode = πTemp001
					// line 318: node.source, node.line = inliner.reporter.get_source_and_line(lineno)
					πF.SetLineno(318)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp005[0] = µlineno
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_source_and_line, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µnode, ßsource, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µnode, ßline, πTemp003); πE != nil {
						continue
					}
					// line 319: return [node], []
					πF.SetLineno(319)
					πTemp005 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp005[0] = µnode
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp005...).ToObject()
					πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßraw_role.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 321: raw_role.options = {'format': directives.unchanged}
			πF.SetLineno(321)
			πTemp004 = πg.NewDict()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
				continue
			}
			if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßunchanged, nil); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßformat.ToObject(), πTemp014); πE != nil {
				continue
			}
			πTemp013 = πTemp004.ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßraw_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ßoptions, πTemp014); πE != nil {
				continue
			}
			// line 323: register_canonical_role('raw', raw_role)
			πF.SetLineno(323)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßraw.ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßraw_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 325: def code_role(role, rawtext, text, lineno, inliner, options={}, content=[]):
			πF.SetLineno(325)
			πTemp005 = make([]πg.Param, 7)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp014 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "options", Def: πTemp014}
			πTemp002 = make([]*πg.Object, 0)
			πTemp014 = πg.NewList(πTemp002...).ToObject()
			πTemp005[6] = πg.Param{Name: "content", Def: πTemp014}
			πTemp013 = πg.NewFunction(πg.NewCode("code_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µoptions *πg.Object = πArgs[5]
				_ = µoptions
				var µcontent *πg.Object = πArgs[6]
				_ = µcontent
				var µlanguage *πg.Object = πg.UnboundLocal
				_ = µlanguage
				var µclasses *πg.Object = πg.UnboundLocal
				_ = µclasses
				var µtokens *πg.Object = πg.UnboundLocal
				_ = µtokens
				var µerror *πg.Object = πg.UnboundLocal
				_ = µerror
				var µmsg *πg.Object = πg.UnboundLocal
				_ = µmsg
				var µprb *πg.Object = πg.UnboundLocal
				_ = µprb
				var µnode *πg.Object = πg.UnboundLocal
				_ = µnode
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var πTemp001 []*πg.Object
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
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 πg.KWArgs
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 9:
						goto Label9
					case 10:
						goto Label10
					case 7:
						goto Label7
					default:
						panic("unexpected function state")
					}
					// line 326: set_classes(options)
					πF.SetLineno(326)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp001[0] = µoptions
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 327: language = options.get('language', '')
					πF.SetLineno(327)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = ßlanguage.ToObject()
					πTemp001[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µoptions, ßget, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlanguage = πTemp003
					// line 328: classes = ['code']
					πF.SetLineno(328)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = ßcode.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µclasses = πTemp002
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µoptions, ßclasses.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 329: if 'classes' in options:
					πF.SetLineno(329)
				Label1:
					// line 330: classes.extend(options['classes'])
					πF.SetLineno(330)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = ßclasses.ToObject()
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µoptions, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µclasses, ßextend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp002 = µlanguage
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µclasses, µlanguage); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					πTemp002 = πTemp003
				Label3:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 331: if language and language not in classes:
					πF.SetLineno(331)
				Label4:
					// line 332: classes.append(language)
					πF.SetLineno(332)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp001[0] = µlanguage
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µclasses, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label5
				Label5:
					// line 333: try:
					πF.SetLineno(333)
					πF.PushCheckpoint(7)
					// line 334: tokens = Lexer(utils.unescape(text, True), language,
					πF.SetLineno(334)
					πTemp001 = πF.MakeArgs(3)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp006[0] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp006[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunescape, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp001[1] = µlanguage
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinliner, ßdocument, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsettings, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsyntax_highlight, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLexer); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtokens = πTemp003
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLexerError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 336: except LexerError as error:
					πF.SetLineno(336)
				Label8:
					µerror = πTemp007.ToObject()
					// line 337: msg = inliner.reporter.warning(error)
					πF.SetLineno(337)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
						continue
					}
					πTemp001[0] = µerror
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwarning, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmsg = πTemp002
					// line 338: prb = inliner.problematic(rawtext, rawtext, msg)
					πF.SetLineno(338)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[0] = µrawtext
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[1] = µrawtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[2] = µmsg
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinliner, ßproblematic, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µprb = πTemp003
					// line 339: return [prb], [msg]
					πF.SetLineno(339)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
						continue
					}
					πTemp001[0] = µprb
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					πTemp009 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp009).ToObject()
					πR = πTemp002
					continue
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					// line 341: node = nodes.literal(rawtext, '', classes=classes)
					πF.SetLineno(341)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[0] = µrawtext
					πTemp001[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"classes", µclasses},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßliteral, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp002
					if πE = πg.CheckLocal(πF, µtokens, "tokens"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µtokens); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp004 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label11
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp011}}}, πTemp003); πE != nil {
							continue
						}
						µclasses = πTemp009
						µvalue = πTemp011
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(9)
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µclasses); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					goto Label13
					// line 345: if classes:
					πF.SetLineno(345)
				Label12:
					// line 346: node += nodes.inline(value, value, classes=classes)
					πF.SetLineno(346)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[1] = µvalue
					if πE = πg.CheckLocal(πF, µclasses, "classes"); πE != nil {
						continue
					}
					πTemp010 = πg.KWArgs{
						{"classes", µclasses},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp003, ßinline, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp009.Call(πF, πTemp001, πTemp010); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IAdd(πF, µnode, πTemp003); πE != nil {
						continue
					}
					µnode = πTemp009
					goto Label14
				Label13:
					// line 349: node += nodes.Text(value, value)
					πF.SetLineno(349)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[1] = µvalue
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp003, ßText, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IAdd(πF, µnode, πTemp003); πE != nil {
						continue
					}
					µnode = πTemp009
					goto Label14
				Label14:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					// line 351: return [node], []
					πF.SetLineno(351)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001[0] = µnode
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 0)
					πTemp009 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp009).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßcode_role.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 353: code_role.options = {'class': directives.class_option,
			πF.SetLineno(353)
			πTemp004 = πg.NewDict()
			if πTemp014, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßclass_option, nil); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßclass.ToObject(), πTemp015); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßdirectives); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßunchanged, nil); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßlanguage.ToObject(), πTemp015); πE != nil {
				continue
			}
			πTemp014 = πTemp004.ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßcode_role); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ßoptions, πTemp015); πE != nil {
				continue
			}
			// line 356: register_canonical_role('code', code_role)
			πF.SetLineno(356)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßcode.ToObject()
			if πTemp014, πE = πg.ResolveGlobal(πF, ßcode_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp014
			if πTemp014, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp015, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 358: def math_role(role, rawtext, text, lineno, inliner, options={}, content=[]):
			πF.SetLineno(358)
			πTemp005 = make([]πg.Param, 7)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp015 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "options", Def: πTemp015}
			πTemp002 = make([]*πg.Object, 0)
			πTemp015 = πg.NewList(πTemp002...).ToObject()
			πTemp005[6] = πg.Param{Name: "content", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("math_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µoptions *πg.Object = πArgs[5]
				_ = µoptions
				var µcontent *πg.Object = πArgs[6]
				_ = µcontent
				var µi *πg.Object = πg.UnboundLocal
				_ = µi
				var µnode *πg.Object = πg.UnboundLocal
				_ = µnode
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
					// line 359: set_classes(options)
					πF.SetLineno(359)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp001[0] = µoptions
					if πTemp002, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 360: i = rawtext.find('`')
					πF.SetLineno(360)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("`").ToObject()
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µrawtext, ßfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp003
					// line 361: text = rawtext.split('`')[1]
					πF.SetLineno(361)
					πTemp002 = πg.NewInt(1).ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("`").ToObject()
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µrawtext, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					µtext = πTemp003
					// line 362: node = nodes.math(rawtext, text, **options)
					πF.SetLineno(362)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[0] = µrawtext
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[1] = µtext
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmath, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, nil, nil, µoptions); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp002
					// line 363: return [node], []
					πF.SetLineno(363)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001[0] = µnode
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßmath_role.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 365: register_canonical_role('math', math_role)
			πF.SetLineno(365)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßmath.ToObject()
			if πTemp015, πE = πg.ResolveGlobal(πF, ßmath_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp015
			if πTemp015, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 371: def unimplemented_role(role, rawtext, text, lineno, inliner, attributes={}):
			πF.SetLineno(371)
			πTemp005 = make([]πg.Param, 6)
			πTemp005[0] = πg.Param{Name: "role", Def: nil}
			πTemp005[1] = πg.Param{Name: "rawtext", Def: nil}
			πTemp005[2] = πg.Param{Name: "text", Def: nil}
			πTemp005[3] = πg.Param{Name: "lineno", Def: nil}
			πTemp005[4] = πg.Param{Name: "inliner", Def: nil}
			πTemp004 = πg.NewDict()
			πTemp016 = πTemp004.ToObject()
			πTemp005[5] = πg.Param{Name: "attributes", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("unimplemented_role", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µrole *πg.Object = πArgs[0]
				_ = µrole
				var µrawtext *πg.Object = πArgs[1]
				_ = µrawtext
				var µtext *πg.Object = πArgs[2]
				_ = µtext
				var µlineno *πg.Object = πArgs[3]
				_ = µlineno
				var µinliner *πg.Object = πArgs[4]
				_ = µinliner
				var µattributes *πg.Object = πArgs[5]
				_ = µattributes
				var µmsg *πg.Object = πg.UnboundLocal
				_ = µmsg
				var µprb *πg.Object = πg.UnboundLocal
				_ = µprb
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 πg.KWArgs
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
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
					// line 372: msg = inliner.reporter.error(
					πF.SetLineno(372)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrole, "role"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("Interpreted text role \"%s\" not implemented.").ToObject(), µrole); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp003 = πg.KWArgs{
						{"line", µlineno},
					}
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinliner, ßreporter, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßerror, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp004.Call(πF, πTemp001, πTemp003); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmsg = πTemp002
					// line 374: prb = inliner.problematic(rawtext, rawtext, msg)
					πF.SetLineno(374)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[0] = µrawtext
					if πE = πg.CheckLocal(πF, µrawtext, "rawtext"); πE != nil {
						continue
					}
					πTemp001[1] = µrawtext
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[2] = µmsg
					if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µinliner, ßproblematic, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µprb = πTemp004
					// line 375: return [prb], [msg]
					πF.SetLineno(375)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µprb, "prb"); πE != nil {
						continue
					}
					πTemp001[0] = µprb
					πTemp004 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					πTemp005 = πg.NewList(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßunimplemented_role.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 377: register_canonical_role('index', unimplemented_role)
			πF.SetLineno(377)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßindex.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 378: register_canonical_role('named-reference', unimplemented_role)
			πF.SetLineno(378)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("named-reference").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 379: register_canonical_role('anonymous-reference', unimplemented_role)
			πF.SetLineno(379)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("anonymous-reference").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 380: register_canonical_role('uri-reference', unimplemented_role)
			πF.SetLineno(380)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("uri-reference").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 381: register_canonical_role('footnote-reference', unimplemented_role)
			πF.SetLineno(381)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("footnote-reference").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 382: register_canonical_role('citation-reference', unimplemented_role)
			πF.SetLineno(382)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("citation-reference").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 383: register_canonical_role('substitution-reference', unimplemented_role)
			πF.SetLineno(383)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("substitution-reference").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 384: register_canonical_role('target', unimplemented_role)
			πF.SetLineno(384)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßtarget.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 387: register_canonical_role('restructuredtext-unimplemented-role',
			πF.SetLineno(387)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("restructuredtext-unimplemented-role").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunimplemented_role); πE != nil {
				continue
			}
			πTemp002[1] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßregister_canonical_role); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 391: def set_classes(options):
			πF.SetLineno(391)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "options", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("set_classes", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/roles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoptions *πg.Object = πArgs[0]
				_ = µoptions
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 392: """
					πF.SetLineno(392)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µoptions, ßclass.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 396: if 'class' in options:
					πF.SetLineno(396)
				Label1:
					// line 397: assert 'classes' not in options
					πF.SetLineno(397)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µoptions, ßclasses.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 398: options['classes'] = options['class']
					πF.SetLineno(398)
					πTemp001 = ßclass.ToObject()
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µoptions, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp004 = ßclasses.ToObject()
					if πE = πg.SetItem(πF, µoptions, πTemp004, πTemp001); πE != nil {
						continue
					}
					// line 399: del options['class']
					πF.SetLineno(399)
					if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
						continue
					}
					πTemp001 = ßclass.ToObject()
					if πE = πg.DelItem(πF, µoptions, πTemp001); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßset_classes.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 392: """
			πF.SetLineno(392)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("\n    Auxiliary function to set options['classes'] and delete\n    options['class'].\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßset_classes); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("roles", Code)
}
