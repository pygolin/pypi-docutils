package rst

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßBody := πg.InternStr("Body")
		ßComponent := πg.InternStr("Component")
		ßDirective := πg.InternStr("Directive")
		ßDirectiveError := πg.InternStr("DirectiveError")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßParser := πg.InternStr("Parser")
		ßRFC2822Body := πg.InternStr("RFC2822Body")
		ßRSTStateMachine := πg.InternStr("RSTStateMachine")
		ßSmartQuotes := πg.InternStr("SmartQuotes")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_argument_spec := πg.InternStr("_argument_spec")
		ß_roles := πg.InternStr("_roles")
		ßaction := πg.InternStr("action")
		ßadd_name := πg.InternStr("add_name")
		ßappend := πg.InternStr("append")
		ßarguments := πg.InternStr("arguments")
		ßassert_has_content := πg.InternStr("assert_has_content")
		ßblock_text := πg.InternStr("block_text")
		ßcharacter_level_inline_markup := πg.InternStr("character_level_inline_markup")
		ßchoices := πg.InternStr("choices")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßcontent := πg.InternStr("content")
		ßcontent_offset := πg.InternStr("content_offset")
		ßconvert_directive_function := πg.InternStr("convert_directive_function")
		ßdebug := πg.InternStr("debug")
		ßdebug_flag := πg.InternStr("debug_flag")
		ßdefault := πg.InternStr("default")
		ßdest := πg.InternStr("dest")
		ßdirective_error := πg.InternStr("directive_error")
		ßdirective_fn := πg.InternStr("directive_fn")
		ßdocument := πg.InternStr("document")
		ßdocutils := πg.InternStr("docutils")
		ßerror := πg.InternStr("error")
		ßfile_insertion_enabled := πg.InternStr("file_insertion_enabled")
		ßfinal_argument_whitespace := πg.InternStr("final_argument_whitespace")
		ßfinish_parse := πg.InternStr("finish_parse")
		ßfrontend := πg.InternStr("frontend")
		ßfully_normalize_name := πg.InternStr("fully_normalize_name")
		ßget_transforms := πg.InternStr("get_transforms")
		ßgetattr := πg.InternStr("getattr")
		ßhas_content := πg.InternStr("has_content")
		ßinfo := πg.InternStr("info")
		ßinitial_state := πg.InternStr("initial_state")
		ßinliner := πg.InternStr("inliner")
		ßint := πg.InternStr("int")
		ßlevel := πg.InternStr("level")
		ßlineno := πg.InternStr("lineno")
		ßlong := πg.InternStr("long")
		ßmetavar := πg.InternStr("metavar")
		ßmsg := πg.InternStr("msg")
		ßname := πg.InternStr("name")
		ßnames := πg.InternStr("names")
		ßnodes := πg.InternStr("nodes")
		ßnone := πg.InternStr("none")
		ßnote_explicit_target := πg.InternStr("note_explicit_target")
		ßobject := πg.InternStr("object")
		ßoption_spec := πg.InternStr("option_spec")
		ßoptional_arguments := πg.InternStr("optional_arguments")
		ßoptions := πg.InternStr("options")
		ßparse := πg.InternStr("parse")
		ßparsers := πg.InternStr("parsers")
		ßpop := πg.InternStr("pop")
		ßraw_enabled := πg.InternStr("raw_enabled")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßreporter := πg.InternStr("reporter")
		ßrequired_arguments := πg.InternStr("required_arguments")
		ßrest := πg.InternStr("rest")
		ßrestructuredtext := πg.InternStr("restructuredtext")
		ßrestx := πg.InternStr("restx")
		ßroles := πg.InternStr("roles")
		ßrst := πg.InternStr("rst")
		ßrstx := πg.InternStr("rstx")
		ßrtxt := πg.InternStr("rtxt")
		ßrun := πg.InternStr("run")
		ßsettings := πg.InternStr("settings")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsetup_parse := πg.InternStr("setup_parse")
		ßsevere := πg.InternStr("severe")
		ßshort := πg.InternStr("short")
		ßstate := πg.InternStr("state")
		ßstate_classes := πg.InternStr("state_classes")
		ßstate_machine := πg.InternStr("state_machine")
		ßstatemachine := πg.InternStr("statemachine")
		ßstates := πg.InternStr("states")
		ßstore_false := πg.InternStr("store_false")
		ßstore_true := πg.InternStr("store_true")
		ßstring2lines := πg.InternStr("string2lines")
		ßsupported := πg.InternStr("supported")
		ßtab_width := πg.InternStr("tab_width")
		ßtrim_footnote_reference_space := πg.InternStr("trim_footnote_reference_space")
		ßtype := πg.InternStr("type")
		ßuniversal := πg.InternStr("universal")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidate_nonnegative_int := πg.InternStr("validate_nonnegative_int")
		ßvalidate_smartquotes_locales := πg.InternStr("validate_smartquotes_locales")
		ßvalidate_ternary := πg.InternStr("validate_ternary")
		ßvalidate_url_trailing_slash := πg.InternStr("validate_url_trailing_slash")
		ßvalidator := πg.InternStr("validator")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nThis is ``docutils.parsers.rst`` package. It exports a single class, `Parser`,\nthe reStructuredText parser.\n\n\nUsage\n=====\n\n1. Create a parser::\n\n       parser = docutils.parsers.rst.Parser()\n\n   Several optional arguments may be passed to modify the parser's behavior.\n   Please see `Customizing the Parser`_ below for details.\n\n2. Gather input (a multi-line string), by reading a file or the standard\n   input::\n\n       input = sys.stdin.read()\n\n3. Create a new empty `docutils.nodes.document` tree::\n\n       document = docutils.utils.new_document(source, settings)\n\n   See `docutils.utils.new_document()` for parameter details.\n\n4. Run the parser, populating the document tree::\n\n       parser.parse(input, document)\n\n\nParser Overview\n===============\n\nThe reStructuredText parser is implemented as a state machine, examining its\ninput one line at a time. To understand how the parser works, please first\nbecome familiar with the `docutils.statemachine` module, then see the\n`states` module.\n\n\nCustomizing the Parser\n----------------------\n\nAnything that isn't already customizable is that way simply because that type\nof customizability hasn't been implemented yet.  Patches welcome!\n\nWhen instantiating an object of the `Parser` class, two parameters may be\npassed: ``rfc2822`` and ``inliner``.  Pass ``rfc2822=True`` to enable an\ninitial RFC-2822 style header block, parsed as a \"field_list\" element (with\n\"class\" attribute set to \"rfc2822\").  Currently this is the only body-level\nelement which is customizable without subclassing.  (Tip: subclass `Parser`\nand change its \"state_classes\" and \"initial_state\" attributes to refer to new\nclasses. Contact the author if you need more details.)\n\nThe ``inliner`` parameter takes an instance of `states.Inliner` or a subclass.\nIt handles inline markup recognition.  A common extension is the addition of\nfurther implicit hyperlinks, like \"RFC 2822\".  This can be done by subclassing\n`states.Inliner`, adding a new method for the implicit markup, and adding a\n``(pattern, method)`` pair to the \"implicit_dispatch\" attribute of the\nsubclass.  See `states.Inliner.implicit_inline()` for details.  Explicit\ninline markup can be customized in a `states.Inliner` subclass via the\n``patterns.initial`` and ``dispatch`` attributes (and new methods as\nappropriate).\n").ToObject()); πE != nil {
				continue
			}
			// line 70: __docformat__ = 'reStructuredText'
			πF.SetLineno(70)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 73: import docutils.parsers
			πF.SetLineno(73)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 74: import docutils.statemachine
			πF.SetLineno(74)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.statemachine"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 75: from docutils.parsers.rst import roles, states
			πF.SetLineno(75)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.roles"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßroles.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.parsers.rst.states"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßstates.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 76: from docutils import frontend, nodes, Component
			πF.SetLineno(76)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.frontend"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßfrontend.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßnodes.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßComponent); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßComponent.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 77: from docutils.transforms import universal
			πF.SetLineno(77)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.transforms.universal"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πE = πF.Globals().SetItem(πF, ßuniversal.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 80: class Parser(docutils.parsers.Parser):
			πF.SetLineno(80)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßparsers, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßParser, nil); πE != nil {
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
			_, πE = πg.NewCode("Parser", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Dict
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 82: """The reStructuredText parser."""
					πF.SetLineno(82)
					// line 82: """The reStructuredText parser."""
					πF.SetLineno(82)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("The reStructuredText parser.").ToObject()); πE != nil {
						continue
					}
					// line 84: supported = ('restructuredtext', 'rst', 'rest', 'restx', 'rtxt', 'rstx')
					πF.SetLineno(84)
					πTemp001 = πg.NewTuple6(ßrestructuredtext.ToObject(), ßrst.ToObject(), ßrest.ToObject(), ßrestx.ToObject(), ßrtxt.ToObject(), ßrstx.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßsupported.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 85: """Aliases this parser supports."""
					πF.SetLineno(85)
					// line 87: settings_spec = (
					πF.SetLineno(87)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 17)
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--pep-references").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Recognize and link to standalone PEP references (like \"PEP 258\").").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[0] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--pep-base-url").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πg.NewStr("http://www.python.org/dev/peps/").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_url_trailing_slash, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Base URL for PEP references (default \"http://www.python.org/dev/peps/\").").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[1] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--pep-file-url-template").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πg.NewStr("pep-%04d").ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Template for PEP file part of URL. (default \"pep-%04d\")").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[2] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--rfc-references").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Recognize and link to standalone RFC references (like \"RFC 822\").").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[3] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--rfc-base-url").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πg.NewStr("http://tools.ietf.org/html/").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_url_trailing_slash, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Base URL for RFC references (default \"http://tools.ietf.org/html/\").").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[4] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--tab-width").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<width>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßtype.ToObject(), ßint.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_nonnegative_int, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Set number of spaces for tab expansion (default 8).").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[5] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--trim-footnote-reference-space").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Remove spaces before footnote references.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[6] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--leave-footnote-reference-space").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdest.ToObject(), ßtrim_footnote_reference_space.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Leave spaces before footnote references.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[7] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--no-file-insertion").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdest.ToObject(), ßfile_insertion_enabled.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Disable directives that insert the contents of external file (\"include\" & \"raw\"); replaced with a \"warning\" system message.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[8] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--file-insertion-enabled").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Enable directives that insert the contents of external file (\"include\" & \"raw\").  Enabled by default.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[9] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--no-raw").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdest.ToObject(), ßraw_enabled.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_boolean, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Disable the \"raw\" directives; replaced with a \"warning\" system message.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[10] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--raw-enabled").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Enable the \"raw\" directive.  Enabled by default.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[11] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--syntax-highlight").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					πTemp006 = make([]*πg.Object, 3)
					πTemp006[0] = ßlong.ToObject()
					πTemp006[1] = ßshort.ToObject()
					πTemp006[2] = ßnone.ToObject()
					πTemp009 = πg.NewList(πTemp006...).ToObject()
					if πE = πTemp008.SetItem(πF, ßchoices.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), ßlong.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<format>").ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Token name set for parsing code with Pygments: one of \"long\", \"short\", or \"none (no parsing)\". Default is \"long\".").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[12] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--smart-quotes").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<yes/no/alt>").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_ternary, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Change straight quotation marks to typographic form: one of \"yes\", \"no\", \"alt[ernative]\" (default \"no\").").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[13] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--smartquotes-locales").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<language:quotes[,language:quotes,...]>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßappend.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfrontend); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßvalidate_smartquotes_locales, nil); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Characters to use as \"smart quotes\" for <language>. ").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[14] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--word-level-inline-markup").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdest.ToObject(), ßcharacter_level_inline_markup.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Inline markup recognized at word boundaries only (adjacent to punctuation or whitespace). Force character-level inline markup recognition with \"\\ \" (backslash + space). Default.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[15] = πTemp005
					πTemp006 = make([]*πg.Object, 1)
					πTemp006[0] = πg.NewStr("--character-level-inline-markup").ToObject()
					πTemp007 = πg.NewList(πTemp006...).ToObject()
					πTemp008 = πg.NewDict()
					if πE = πTemp008.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdefault.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πTemp008.SetItem(πF, ßdest.ToObject(), ßcharacter_level_inline_markup.ToObject()); πE != nil {
						continue
					}
					πTemp009 = πTemp008.ToObject()
					πTemp005 = πg.NewTuple3(πg.NewStr("Inline markup recognized anywhere, regardless of surrounding characters. Backslash-escapes must be used to avoid unwanted markup recognition. Useful for East Asian languages. Experimental.").ToObject(), πTemp007, πTemp009).ToObject()
					πTemp004[16] = πTemp005
					πTemp003 = πg.NewTuple(πTemp004...).ToObject()
					πTemp001 = πg.NewTuple3(πg.NewStr("reStructuredText Parser Options").ToObject(), πTemp002, πTemp003).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 166: config_section = 'restructuredtext parser'
					πF.SetLineno(166)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), πg.NewStr("restructuredtext parser").ToObject()); πE != nil {
						continue
					}
					// line 167: config_section_dependencies = ('parsers',)
					πF.SetLineno(167)
					πTemp001 = πg.NewTuple1(ßparsers.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßconfig_section_dependencies.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 169: def __init__(self, rfc2822=False, inliner=None):
					πF.SetLineno(169)
					πTemp011 = make([]πg.Param, 3)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp011[1] = πg.Param{Name: "rfc2822", Def: πTemp002}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp011[2] = πg.Param{Name: "inliner", Def: πTemp002}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrfc2822 *πg.Object = πArgs[1]
						_ = µrfc2822
						var µinliner *πg.Object = πArgs[2]
						_ = µinliner
						var πTemp001 bool
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
							if πE = πg.CheckLocal(πF, µrfc2822, "rfc2822"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µrfc2822); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 170: if rfc2822:
							πF.SetLineno(170)
						Label1:
							// line 171: self.initial_state = 'RFC2822Body'
							πF.SetLineno(171)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßRFC2822Body.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinitial_state, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 173: self.initial_state = 'Body'
							πF.SetLineno(173)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßBody.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinitial_state, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 174: self.state_classes = states.state_classes
							πF.SetLineno(174)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstate_classes, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate_classes, πTemp002); πE != nil {
								continue
							}
							// line 175: self.inliner = inliner
							πF.SetLineno(175)
							if πE = πg.CheckLocal(πF, µinliner, "inliner"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µinliner); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinliner, πTemp002); πE != nil {
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
					// line 177: def get_transforms(self):
					πF.SetLineno(177)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("get_transforms", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 178: return Component.get_transforms(self) + [
							πF.SetLineno(178)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßComponent); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget_transforms, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßuniversal); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßSmartQuotes, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßget_transforms.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 181: def parse(self, inputstring, document):
					πF.SetLineno(181)
					πTemp011 = make([]πg.Param, 3)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "inputstring", Def: nil}
					πTemp011[2] = πg.Param{Name: "document", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinputstring *πg.Object = πArgs[1]
						_ = µinputstring
						var µdocument *πg.Object = πArgs[2]
						_ = µdocument
						var µinputlines *πg.Object = πg.UnboundLocal
						_ = µinputlines
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
							// line 182: """Parse `inputstring` and populate `document`, a document tree."""
							πF.SetLineno(182)
							// line 183: self.setup_parse(inputstring, document)
							πF.SetLineno(183)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinputstring, "inputstring"); πE != nil {
								continue
							}
							πTemp001[0] = µinputstring
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsetup_parse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 184: self.statemachine = states.RSTStateMachine(
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_classes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinitial_state, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µdocument, ßreporter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßdebug_flag, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"state_classes", πTemp002},
								{"initial_state", πTemp003},
								{"debug", πTemp005},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstates); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßRSTStateMachine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstatemachine, πTemp003); πE != nil {
								continue
							}
							// line 188: inputlines = docutils.statemachine.string2lines(
							πF.SetLineno(188)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinputstring, "inputstring"); πE != nil {
								continue
							}
							πTemp001[0] = µinputstring
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdocument, ßsettings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtab_width, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"tab_width", πTemp003},
								{"convert_whitespace", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstatemachine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstring2lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µinputlines = πTemp003
							// line 191: self.statemachine.run(inputlines, document, inliner=self.inliner)
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinputlines, "inputlines"); πE != nil {
								continue
							}
							πTemp001[0] = µinputlines
							if πE = πg.CheckLocal(πF, µdocument, "document"); πE != nil {
								continue
							}
							πTemp001[1] = µdocument
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinliner, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"inliner", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstatemachine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_roles, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, πTemp004, ß.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 193: if '' in roles._roles:
							πF.SetLineno(193)
						Label1:
							// line 194: del roles._roles['']
							πF.SetLineno(194)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßroles); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_roles, nil); πE != nil {
								continue
							}
							πTemp002 = ß.ToObject()
							if πE = πg.DelItem(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 195: self.finish_parse()
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfinish_parse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßparse.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 182: """Parse `inputstring` and populate `document`, a document tree."""
					πF.SetLineno(182)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Parse `inputstring` and populate `document`, a document tree.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßparse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Parser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßParser.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 198: class DirectiveError(Exception):
			πF.SetLineno(198)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("DirectiveError", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 200: """
					πF.SetLineno(200)
					// line 200: """
					πF.SetLineno(200)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Store a message and a system message level.\n\n    To be thrown from inside directive code.\n\n    Do not instantiate directly -- use `Directive.directive_error()`\n    instead!\n    ").ToObject()); πE != nil {
						continue
					}
					// line 209: def __init__(self, level, message):
					πF.SetLineno(209)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "level", Def: nil}
					πTemp002[2] = πg.Param{Name: "message", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlevel *πg.Object = πArgs[1]
						_ = µlevel
						var µmessage *πg.Object = πArgs[2]
						_ = µmessage
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
							// line 210: """Set error `message` and `level`"""
							πF.SetLineno(210)
							// line 211: Exception.__init__(self)
							πF.SetLineno(211)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 212: self.level = level
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µlevel); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
								continue
							}
							// line 213: self.msg = message
							πF.SetLineno(213)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µmessage); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmsg, πTemp002); πE != nil {
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
					// line 210: """Set error `message` and `level`"""
					πF.SetLineno(210)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Set error `message` and `level`").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DirectiveError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDirectiveError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 216: class Directive(object):
			πF.SetLineno(216)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("Directive", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 218: """
					πF.SetLineno(218)
					// line 218: """
					πF.SetLineno(218)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Base class for reStructuredText directives.\n\n    The following attributes may be set by subclasses.  They are\n    interpreted by the directive parser (which runs the directive\n    class):\n\n    - `required_arguments`: The number of required arguments (default:\n      0).\n\n    - `optional_arguments`: The number of optional arguments (default:\n      0).\n\n    - `final_argument_whitespace`: A boolean, indicating if the final\n      argument may contain whitespace (default: False).\n\n    - `option_spec`: A dictionary, mapping known option names to\n      conversion functions such as `int` or `float` (default: {}, no\n      options).  Several conversion functions are defined in the\n      directives/__init__.py module.\n\n      Option conversion functions take a single parameter, the option\n      argument (a string or ``None``), validate it and/or convert it\n      to the appropriate form.  Conversion functions may raise\n      `ValueError` and `TypeError` exceptions.\n\n    - `has_content`: A boolean; True if content is allowed.  Client\n      code must handle the case where content is required but not\n      supplied (an empty content list will be supplied).\n\n    Arguments are normally single whitespace-separated words.  The\n    final argument may contain whitespace and/or newlines if\n    `final_argument_whitespace` is True.\n\n    If the form of the arguments is more complex, specify only one\n    argument (either required or optional) and set\n    `final_argument_whitespace` to True; the client code must do any\n    context-sensitive parsing.\n\n    When a directive implementation is being run, the directive class\n    is instantiated, and the `run()` method is executed.  During\n    instantiation, the following instance variables are set:\n\n    - ``name`` is the directive type or name (string).\n\n    - ``arguments`` is the list of positional arguments (strings).\n\n    - ``options`` is a dictionary mapping option names (strings) to\n      values (type depends on option conversion functions; see\n      `option_spec` above).\n\n    - ``content`` is a list of strings, the directive content line by line.\n\n    - ``lineno`` is the absolute line number of the first line\n      of the directive.\n\n    - ``content_offset`` is the line offset of the first line of the content from\n      the beginning of the current input.  Used when initiating a nested parse.\n\n    - ``block_text`` is a string containing the entire directive.\n\n    - ``state`` is the state which called the directive function.\n\n    - ``state_machine`` is the state machine which controls the state which called\n      the directive function.\n\n    Directive functions return a list of nodes which will be inserted\n    into the document tree at the point where the directive was\n    encountered.  This can be an empty list if there is nothing to\n    insert.\n\n    For ordinary directives, the list must contain body elements or\n    structural elements.  Some directives are intended specifically\n    for substitution definitions, and must return a list of `Text`\n    nodes and/or inline elements (suitable for inline insertion, in\n    place of the substitution reference).  Such directives must verify\n    substitution definition context, typically using code like this::\n\n        if not isinstance(state, states.SubstitutionDef):\n            error = state_machine.reporter.error(\n                'Invalid context: the \"%s\" directive can only be used '\n                'within a substitution definition.' % (name),\n                nodes.literal_block(block_text, block_text), line=lineno)\n            return [error]\n    ").ToObject()); πE != nil {
						continue
					}
					// line 308: required_arguments = 0
					πF.SetLineno(308)
					if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 309: """Number of required directive arguments."""
					πF.SetLineno(309)
					// line 311: optional_arguments = 0
					πF.SetLineno(311)
					if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 312: """Number of optional arguments after the required arguments."""
					πF.SetLineno(312)
					// line 314: final_argument_whitespace = False
					πF.SetLineno(314)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 315: """May the final argument contain whitespace?"""
					πF.SetLineno(315)
					// line 317: option_spec = None
					πF.SetLineno(317)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 318: """Mapping of option names to validator functions."""
					πF.SetLineno(318)
					// line 320: has_content = False
					πF.SetLineno(320)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 321: """May the directive have content?"""
					πF.SetLineno(321)
					// line 323: def __init__(self, name, arguments, options, content, lineno,
					πF.SetLineno(323)
					πTemp002 = make([]πg.Param, 10)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "arguments", Def: nil}
					πTemp002[3] = πg.Param{Name: "options", Def: nil}
					πTemp002[4] = πg.Param{Name: "content", Def: nil}
					πTemp002[5] = πg.Param{Name: "lineno", Def: nil}
					πTemp002[6] = πg.Param{Name: "content_offset", Def: nil}
					πTemp002[7] = πg.Param{Name: "block_text", Def: nil}
					πTemp002[8] = πg.Param{Name: "state", Def: nil}
					πTemp002[9] = πg.Param{Name: "state_machine", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var µarguments *πg.Object = πArgs[2]
						_ = µarguments
						var µoptions *πg.Object = πArgs[3]
						_ = µoptions
						var µcontent *πg.Object = πArgs[4]
						_ = µcontent
						var µlineno *πg.Object = πArgs[5]
						_ = µlineno
						var µcontent_offset *πg.Object = πArgs[6]
						_ = µcontent_offset
						var µblock_text *πg.Object = πArgs[7]
						_ = µblock_text
						var µstate *πg.Object = πArgs[8]
						_ = µstate
						var µstate_machine *πg.Object = πArgs[9]
						_ = µstate_machine
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
							// line 325: self.name = name
							πF.SetLineno(325)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µname); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßname, πTemp001); πE != nil {
								continue
							}
							// line 326: self.arguments = arguments
							πF.SetLineno(326)
							if πE = πg.CheckLocal(πF, µarguments, "arguments"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µarguments); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßarguments, πTemp001); πE != nil {
								continue
							}
							// line 327: self.options = options
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µoptions); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoptions, πTemp001); πE != nil {
								continue
							}
							// line 328: self.content = content
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcontent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontent, πTemp001); πE != nil {
								continue
							}
							// line 329: self.lineno = lineno
							πF.SetLineno(329)
							if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlineno); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlineno, πTemp001); πE != nil {
								continue
							}
							// line 330: self.content_offset = content_offset
							πF.SetLineno(330)
							if πE = πg.CheckLocal(πF, µcontent_offset, "content_offset"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcontent_offset); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontent_offset, πTemp001); πE != nil {
								continue
							}
							// line 331: self.block_text = block_text
							πF.SetLineno(331)
							if πE = πg.CheckLocal(πF, µblock_text, "block_text"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µblock_text); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßblock_text, πTemp001); πE != nil {
								continue
							}
							// line 332: self.state = state
							πF.SetLineno(332)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstate); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp001); πE != nil {
								continue
							}
							// line 333: self.state_machine = state_machine
							πF.SetLineno(333)
							if πE = πg.CheckLocal(πF, µstate_machine, "state_machine"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstate_machine); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate_machine, πTemp001); πE != nil {
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
					// line 335: def run(self):
					πF.SetLineno(335)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Must override run() is subclass.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 336: raise NotImplementedError('Must override run() is subclass.')
							πF.SetLineno(336)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 340: def directive_error(self, level, message):
					πF.SetLineno(340)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "level", Def: nil}
					πTemp002[2] = πg.Param{Name: "message", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("directive_error", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlevel *πg.Object = πArgs[1]
						_ = µlevel
						var µmessage *πg.Object = πArgs[2]
						_ = µmessage
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
							// line 341: """
							πF.SetLineno(341)
							// line 353: return DirectiveError(level, message)
							πF.SetLineno(353)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							πTemp001[0] = µlevel
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πTemp002, πE = πg.ResolveGlobal(πF, ßDirectiveError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßdirective_error.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 341: """
					πF.SetLineno(341)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Return a DirectiveError suitable for being thrown as an exception.\n\n        Call \"raise self.directive_error(level, message)\" from within\n        a directive implementation to return one single system message\n        at level `level`, which automatically gets the directive block\n        and the line number added.\n\n        Preferably use the `debug`, `info`, `warning`, `error`, or `severe`\n        wrapper methods, e.g. ``self.error(message)`` to generate an\n        ERROR-level directive error.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdirective_error); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 355: def debug(self, message):
					πF.SetLineno(355)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("debug", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 356: return self.directive_error(0, message)
							πF.SetLineno(356)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdirective_error, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßdebug.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 358: def info(self, message):
					πF.SetLineno(358)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("info", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 359: return self.directive_error(1, message)
							πF.SetLineno(359)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdirective_error, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßinfo.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 361: def warning(self, message):
					πF.SetLineno(361)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("warning", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 362: return self.directive_error(2, message)
							πF.SetLineno(362)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdirective_error, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßwarning.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 364: def error(self, message):
					πF.SetLineno(364)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("error", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 365: return self.directive_error(3, message)
							πF.SetLineno(365)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdirective_error, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßerror.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 367: def severe(self, message):
					πF.SetLineno(367)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("severe", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmessage *πg.Object = πArgs[1]
						_ = µmessage
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
							// line 368: return self.directive_error(4, message)
							πF.SetLineno(368)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							πTemp001[1] = µmessage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdirective_error, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßsevere.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 372: def assert_has_content(self):
					πF.SetLineno(372)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("assert_has_content", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 373: """
							πF.SetLineno(373)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 377: if not self.content:
							πF.SetLineno(377)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Content block expected for the \"%s\" directive; none found.").ToObject(), πTemp002); πE != nil {
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
							// line 378: raise self.error('Content block expected for the "%s" directive; '
							πF.SetLineno(378)
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
					if πE = πClass.SetItem(πF, ßassert_has_content.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 373: """
					πF.SetLineno(373)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n        Throw an ERROR-level DirectiveError if the directive doesn't\n        have contents.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßassert_has_content); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 381: def add_name(self, node):
					πF.SetLineno(381)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "node", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("add_name", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnode *πg.Object = πArgs[1]
						_ = µnode
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 382: """Append self.options['name'] to node['names'] if it exists.
							πF.SetLineno(382)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, ßname.ToObject()); πE != nil {
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
							// line 386: if 'name' in self.options:
							πF.SetLineno(386)
						Label1:
							// line 387: name = nodes.fully_normalize_name(self.options.pop('name'))
							πF.SetLineno(387)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßname.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnodes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfully_normalize_name, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µname = πTemp001
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µnode, ßname.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 388: if 'name' in node:
							πF.SetLineno(388)
						Label3:
							// line 389: del(node['name'])
							πF.SetLineno(389)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp001 = ßname.ToObject()
							if πE = πg.DelItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 390: node['names'].append(name)
							πF.SetLineno(390)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							πTemp001 = ßnames.ToObject()
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µnode, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 391: self.state.document.note_explicit_target(node, node)
							πF.SetLineno(391)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[0] = µnode
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πTemp004[1] = µnode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdocument, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßnote_explicit_target, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßadd_name.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 382: """Append self.options['name'] to node['names'] if it exists.
					πF.SetLineno(382)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Append self.options['name'] to node['names'] if it exists.\n\n        Also normalize the name string and register it as explicit target.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßadd_name); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Directive").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDirective.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 394: def convert_directive_function(directive_fn):
			πF.SetLineno(394)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "directive_fn", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("convert_directive_function", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdirective_fn *πg.Object = πArgs[0]
				_ = µdirective_fn
				var µFunctionalDirective *πg.Object = πg.UnboundLocal
				_ = µFunctionalDirective
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
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
					// line 395: """
					πF.SetLineno(395)
					// line 401: class FunctionalDirective(Directive):
					πF.SetLineno(401)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßDirective); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp001 = πg.NewDict()
					if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
						continue
					}
					_, πE = πg.NewCode("FunctionalDirective", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
						_ = πClass
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
						var πTemp006 []πg.Param
						_ = πTemp006
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 403: option_spec = getattr(directive_fn, 'options', None)
							πF.SetLineno(403)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveClass(πF, πClass, µdirective_fn, ßdirective_fn); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßoptions.ToObject()
							if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πClass.SetItem(πF, ßoption_spec.ToObject(), πTemp003); πE != nil {
								continue
							}
							// line 404: has_content = getattr(directive_fn, 'content', False)
							πF.SetLineno(404)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveClass(πF, πClass, µdirective_fn, ßdirective_fn); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßcontent.ToObject()
							if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πClass.SetItem(πF, ßhas_content.ToObject(), πTemp003); πE != nil {
								continue
							}
							// line 405: _argument_spec = getattr(directive_fn, 'arguments', (0, 0, False))
							πF.SetLineno(405)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveClass(πF, πClass, µdirective_fn, ßdirective_fn); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßarguments.ToObject()
							if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πTemp003).ToObject()
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πClass.SetItem(πF, ß_argument_spec.ToObject(), πTemp003); πE != nil {
								continue
							}
							// line 406: required_arguments, optional_arguments, final_argument_whitespace \
							πF.SetLineno(406)
							if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ß_argument_spec); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßrequired_arguments.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßoptional_arguments.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πE = πClass.SetItem(πF, ßfinal_argument_whitespace.ToObject(), πTemp005); πE != nil {
								continue
							}
							// line 409: def run(self):
							πF.SetLineno(409)
							πTemp006 = make([]πg.Param, 1)
							πTemp006[0] = πg.Param{Name: "self", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/parsers/rst/__init__.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]
								_ = µself
								var πTemp001 []*πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
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
									// line 410: return directive_fn(
									πF.SetLineno(410)
									πTemp001 = πF.MakeArgs(9)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßarguments, nil); πE != nil {
										continue
									}
									πTemp001[1] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
										continue
									}
									πTemp001[2] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent, nil); πE != nil {
										continue
									}
									πTemp001[3] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßlineno, nil); πE != nil {
										continue
									}
									πTemp001[4] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßcontent_offset, nil); πE != nil {
										continue
									}
									πTemp001[5] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßblock_text, nil); πE != nil {
										continue
									}
									πTemp001[6] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
										continue
									}
									πTemp001[7] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
										continue
									}
									πTemp001[8] = πTemp002
									if πE = πg.CheckLocal(πF, µdirective_fn, "directive_fn"); πE != nil {
										continue
									}
									if πTemp002, πE = µdirective_fn.Call(πF, πTemp001, nil); πE != nil {
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
							if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp002); πE != nil {
								continue
							}
						}
						return nil, nil
					}).Eval(πF, πF.Globals(), nil, nil)
					if πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
						continue
					}
					if πTemp004 == nil {
						πTemp004 = πg.TypeType.ToObject()
					}
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FunctionalDirective").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µFunctionalDirective = πTemp005
					// line 416: return FunctionalDirective
					πF.SetLineno(416)
					if πE = πg.CheckLocal(πF, µFunctionalDirective, "FunctionalDirective"); πE != nil {
						continue
					}
					πR = µFunctionalDirective
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßconvert_directive_function.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 395: """
			πF.SetLineno(395)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Define & return a directive class generated from `directive_fn`.\n\n    `directive_fn` uses the old-style, functional interface.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßconvert_directive_function); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp003); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("rst", Code)
}
