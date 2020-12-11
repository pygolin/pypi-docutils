package statemachine

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/unicodedata"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßDuplicateStateError := πg.InternStr("DuplicateStateError")
		ßDuplicateTransitionError := πg.InternStr("DuplicateTransitionError")
		ßEOFError := πg.InternStr("EOFError")
		ßErrorOutput := πg.InternStr("ErrorOutput")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ßSearchStateMachine := πg.InternStr("SearchStateMachine")
		ßSearchStateMachineWS := πg.InternStr("SearchStateMachineWS")
		ßState := πg.InternStr("State")
		ßStateCorrection := πg.InternStr("StateCorrection")
		ßStateMachine := πg.InternStr("StateMachine")
		ßStateMachineError := πg.InternStr("StateMachineError")
		ßStateMachineWS := πg.InternStr("StateMachineWS")
		ßStateWS := πg.InternStr("StateWS")
		ßStringList := πg.InternStr("StringList")
		ßTransitionCorrection := πg.InternStr("TransitionCorrection")
		ßTransitionMethodNotFound := πg.InternStr("TransitionMethodNotFound")
		ßTransitionPatternNotFound := πg.InternStr("TransitionPatternNotFound")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnexpectedIndentationError := πg.InternStr("UnexpectedIndentationError")
		ßUnknownStateError := πg.InternStr("UnknownStateError")
		ßUnknownTransitionError := πg.InternStr("UnknownTransitionError")
		ßViewList := πg.InternStr("ViewList")
		ßWF := πg.InternStr("WF")
		ß_SearchOverride := πg.InternStr("_SearchOverride")
		ß__add__ := πg.InternStr("__add__")
		ß__cast := πg.InternStr("__cast")
		ß__class__ := πg.InternStr("__class__")
		ß__cmp__ := πg.InternStr("__cmp__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__gt__ := πg.InternStr("__gt__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__imul__ := πg.InternStr("__imul__")
		ß__init__ := πg.InternStr("__init__")
		ß__le__ := πg.InternStr("__le__")
		ß__len__ := πg.InternStr("__len__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__mul__ := πg.InternStr("__mul__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__radd__ := πg.InternStr("__radd__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__rmul__ := πg.InternStr("__rmul__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__str__ := πg.InternStr("__str__")
		ß_exception_data := πg.InternStr("_exception_data")
		ß_stderr := πg.InternStr("_stderr")
		ßabs_line_number := πg.InternStr("abs_line_number")
		ßabs_line_offset := πg.InternStr("abs_line_offset")
		ßadd_initial_transitions := πg.InternStr("add_initial_transitions")
		ßadd_state := πg.InternStr("add_state")
		ßadd_states := πg.InternStr("add_states")
		ßadd_transition := πg.InternStr("add_transition")
		ßadd_transitions := πg.InternStr("add_transitions")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßat_bof := πg.InternStr("at_bof")
		ßat_eof := πg.InternStr("at_eof")
		ßattach_observer := πg.InternStr("attach_observer")
		ßblank := πg.InternStr("blank")
		ßbof := πg.InternStr("bof")
		ßcheck_line := πg.InternStr("check_line")
		ßco_filename := πg.InternStr("co_filename")
		ßco_name := πg.InternStr("co_name")
		ßcolumn_indices := πg.InternStr("column_indices")
		ßcompile := πg.InternStr("compile")
		ßcount := πg.InternStr("count")
		ßcurrent_state := πg.InternStr("current_state")
		ßdata := πg.InternStr("data")
		ßdebug := πg.InternStr("debug")
		ßdetach_observer := πg.InternStr("detach_observer")
		ßdisconnect := πg.InternStr("disconnect")
		ßeast_asian_width := πg.InternStr("east_asian_width")
		ßend := πg.InternStr("end")
		ßeof := πg.InternStr("eof")
		ßerror := πg.InternStr("error")
		ßexc_info := πg.InternStr("exc_info")
		ßexpandtabs := πg.InternStr("expandtabs")
		ßextend := πg.InternStr("extend")
		ßf_code := πg.InternStr("f_code")
		ßfirst_known_indent := πg.InternStr("first_known_indent")
		ßget_2D_block := πg.InternStr("get_2D_block")
		ßget_first_known_indented := πg.InternStr("get_first_known_indented")
		ßget_indented := πg.InternStr("get_indented")
		ßget_known_indented := πg.InternStr("get_known_indented")
		ßget_source := πg.InternStr("get_source")
		ßget_source_and_line := πg.InternStr("get_source_and_line")
		ßget_state := πg.InternStr("get_state")
		ßget_text_block := πg.InternStr("get_text_block")
		ßgetattr := πg.InternStr("getattr")
		ßgoto_line := πg.InternStr("goto_line")
		ßhasattr := πg.InternStr("hasattr")
		ßindent := πg.InternStr("indent")
		ßindent_sm := πg.InternStr("indent_sm")
		ßindent_sm_kwargs := πg.InternStr("indent_sm_kwargs")
		ßindex := πg.InternStr("index")
		ßinfo := πg.InternStr("info")
		ßinitial_state := πg.InternStr("initial_state")
		ßinitial_transitions := πg.InternStr("initial_transitions")
		ßinput_lines := πg.InternStr("input_lines")
		ßinput_offset := πg.InternStr("input_offset")
		ßinsert := πg.InternStr("insert")
		ßinsert_input := πg.InternStr("insert_input")
		ßis_next_line_blank := πg.InternStr("is_next_line_blank")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßknown_indent := πg.InternStr("known_indent")
		ßknown_indent_sm := πg.InternStr("known_indent_sm")
		ßknown_indent_sm_kwargs := πg.InternStr("known_indent_sm_kwargs")
		ßlen := πg.InternStr("len")
		ßline := πg.InternStr("line")
		ßline_offset := πg.InternStr("line_offset")
		ßlist := πg.InternStr("list")
		ßlstrip := πg.InternStr("lstrip")
		ßmake_transition := πg.InternStr("make_transition")
		ßmake_transitions := πg.InternStr("make_transitions")
		ßmatch := πg.InternStr("match")
		ßmaxsize := πg.InternStr("maxsize")
		ßmin := πg.InternStr("min")
		ßnested_sm := πg.InternStr("nested_sm")
		ßnested_sm_kwargs := πg.InternStr("nested_sm_kwargs")
		ßnext_line := πg.InternStr("next_line")
		ßno_match := πg.InternStr("no_match")
		ßnop := πg.InternStr("nop")
		ßnotify_observers := πg.InternStr("notify_observers")
		ßobject := πg.InternStr("object")
		ßobservers := πg.InternStr("observers")
		ßoffset := πg.InternStr("offset")
		ßpad_double_width := πg.InternStr("pad_double_width")
		ßparent := πg.InternStr("parent")
		ßparent_offset := πg.InternStr("parent_offset")
		ßpatterns := πg.InternStr("patterns")
		ßpop := πg.InternStr("pop")
		ßpprint := πg.InternStr("pprint")
		ßprevious_line := πg.InternStr("previous_line")
		ßprint := πg.InternStr("print")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßremove := πg.InternStr("remove")
		ßremove_transition := πg.InternStr("remove_transition")
		ßreplace := πg.InternStr("replace")
		ßrestructuredtext := πg.InternStr("restructuredtext")
		ßreverse := πg.InternStr("reverse")
		ßrstrip := πg.InternStr("rstrip")
		ßrun := πg.InternStr("run")
		ßruntime_init := πg.InternStr("runtime_init")
		ßsearch := πg.InternStr("search")
		ßslice := πg.InternStr("slice")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßsource := πg.InternStr("source")
		ßsplitlines := πg.InternStr("splitlines")
		ßstart := πg.InternStr("start")
		ßstate_classes := πg.InternStr("state_classes")
		ßstate_machine := πg.InternStr("state_machine")
		ßstates := πg.InternStr("states")
		ßstep := πg.InternStr("step")
		ßstop := πg.InternStr("stop")
		ßstr := πg.InternStr("str")
		ßstring2lines := πg.InternStr("string2lines")
		ßstrip := πg.InternStr("strip")
		ßsub := πg.InternStr("sub")
		ßsys := πg.InternStr("sys")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtb_lineno := πg.InternStr("tb_lineno")
		ßtb_next := πg.InternStr("tb_next")
		ßtransition_order := πg.InternStr("transition_order")
		ßtransitions := πg.InternStr("transitions")
		ßtrim_end := πg.InternStr("trim_end")
		ßtrim_left := πg.InternStr("trim_left")
		ßtrim_start := πg.InternStr("trim_start")
		ßtype := πg.InternStr("type")
		ßunicode := πg.InternStr("unicode")
		ßunicodedata := πg.InternStr("unicodedata")
		ßunlink := πg.InternStr("unlink")
		ßupdate := πg.InternStr("update")
		ßutils := πg.InternStr("utils")
		ßvalues := πg.InternStr("values")
		ßversion_info := πg.InternStr("version_info")
		ßws_initial_transitions := πg.InternStr("ws_initial_transitions")
		ßws_patterns := πg.InternStr("ws_patterns")
		ßxitems := πg.InternStr("xitems")
		ßzip := πg.InternStr("zip")
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
		var πTemp008 *πg.Object
		_ = πTemp008
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nA finite state machine specialized for regular-expression-based text filters,\nthis module defines the following classes:\n\n- `StateMachine`, a state machine\n- `State`, a state superclass\n- `StateMachineWS`, a whitespace-sensitive version of `StateMachine`\n- `StateWS`, a state superclass for use with `StateMachineWS`\n- `SearchStateMachine`, uses `re.search()` instead of `re.match()`\n- `SearchStateMachineWS`, uses `re.search()` instead of `re.match()`\n- `ViewList`, extends standard Python lists.\n- `StringList`, string-specific ViewList.\n\nException classes:\n\n- `StateMachineError`\n- `UnknownStateError`\n- `DuplicateStateError`\n- `UnknownTransitionError`\n- `DuplicateTransitionError`\n- `TransitionPatternNotFound`\n- `TransitionMethodNotFound`\n- `UnexpectedIndentationError`\n- `TransitionCorrection`: Raised to switch to another transition.\n- `StateCorrection`: Raised to switch to another state & transition.\n\nFunctions:\n\n- `string2lines()`: split a multi-line string into a list of one-line strings\n\n\nHow To Use This Module\n======================\n(See the individual classes, methods, and attributes for details.)\n\n1. Import it: ``import statemachine`` or ``from statemachine import ...``.\n   You will also need to ``import re``.\n\n2. Derive a subclass of `State` (or `StateWS`) for each state in your state\n   machine::\n\n       class MyState(statemachine.State):\n\n   Within the state's class definition:\n\n   a) Include a pattern for each transition, in `State.patterns`::\n\n          patterns = {'atransition': r'pattern', ...}\n\n   b) Include a list of initial transitions to be set up automatically, in\n      `State.initial_transitions`::\n\n          initial_transitions = ['atransition', ...]\n\n   c) Define a method for each transition, with the same name as the\n      transition pattern::\n\n          def atransition(self, match, context, next_state):\n              # do something\n              result = [...]  # a list\n              return context, next_state, result\n              # context, next_state may be altered\n\n      Transition methods may raise an `EOFError` to cut processing short.\n\n   d) You may wish to override the `State.bof()` and/or `State.eof()` implicit\n      transition methods, which handle the beginning- and end-of-file.\n\n   e) In order to handle nested processing, you may wish to override the\n      attributes `State.nested_sm` and/or `State.nested_sm_kwargs`.\n\n      If you are using `StateWS` as a base class, in order to handle nested\n      indented blocks, you may wish to:\n\n      - override the attributes `StateWS.indent_sm`,\n        `StateWS.indent_sm_kwargs`, `StateWS.known_indent_sm`, and/or\n        `StateWS.known_indent_sm_kwargs`;\n      - override the `StateWS.blank()` method; and/or\n      - override or extend the `StateWS.indent()`, `StateWS.known_indent()`,\n        and/or `StateWS.firstknown_indent()` methods.\n\n3. Create a state machine object::\n\n       sm = StateMachine(state_classes=[MyState, ...],\n                         initial_state='MyState')\n\n4. Obtain the input text, which needs to be converted into a tab-free list of\n   one-line strings. For example, to read text from a file called\n   'inputfile'::\n\n       input_string = open('inputfile').read()\n       input_lines = statemachine.string2lines(input_string)\n\n5. Run the state machine on the input text and collect the results, a list::\n\n       results = sm.run(input_lines)\n\n6. Remove any lingering circular references::\n\n       sm.unlink()\n").ToObject()); πE != nil {
				continue
			}
			// line 106: from __future__ import print_function
			πF.SetLineno(106)
			// line 108: __docformat__ = 'restructuredtext'
			πF.SetLineno(108)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßrestructuredtext.ToObject()); πE != nil {
				continue
			}
			// line 110: import sys
			πF.SetLineno(110)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 111: import re
			πF.SetLineno(111)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 112: import unicodedata
			πF.SetLineno(112)
			if πTemp002, πE = πg.ImportModule(πF, "unicodedata"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunicodedata.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 113: from docutils import utils
			πF.SetLineno(113)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 114: from docutils.utils.error_reporting import ErrorOutput
			πF.SetLineno(114)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorOutput); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorOutput.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßversion_info, nil); πE != nil {
				continue
			}
			πTemp003 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
			if πTemp001, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 116: if sys.version_info >= (3, 0):
			πF.SetLineno(116)
		Label1:
			// line 117: unicode = str  # noqa
			πF.SetLineno(117)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 120: class StateMachine(object):
			πF.SetLineno(120)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StateMachine", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				var πTemp019 *πg.Object
				_ = πTemp019
				var πTemp020 *πg.Object
				_ = πTemp020
				var πTemp021 *πg.Object
				_ = πTemp021
				var πTemp022 *πg.Object
				_ = πTemp022
				var πTemp023 *πg.Object
				_ = πTemp023
				var πTemp024 *πg.Object
				_ = πTemp024
				var πTemp025 *πg.Object
				_ = πTemp025
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 122: """
					πF.SetLineno(122)
					// line 122: """
					πF.SetLineno(122)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A finite state machine for text filters using regular expressions.\n\n    The input is provided in the form of a list of one-line strings (no\n    newlines). States are subclasses of the `State` class. Transitions consist\n    of regular expression patterns and transition methods, and are defined in\n    each state.\n\n    The state machine is started with the `run()` method, which returns the\n    results of processing in a list.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 134: def __init__(self, state_classes, initial_state, debug=False):
					πF.SetLineno(134)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state_classes", Def: nil}
					πTemp002[2] = πg.Param{Name: "initial_state", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "debug", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate_classes *πg.Object = πArgs[1]
						_ = µstate_classes
						var µinitial_state *πg.Object = πArgs[2]
						_ = µinitial_state
						var µdebug *πg.Object = πArgs[3]
						_ = µdebug
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
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
							// line 135: """
							πF.SetLineno(135)
							// line 145: self.input_lines = None
							πF.SetLineno(145)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput_lines, πTemp002); πE != nil {
								continue
							}
							// line 146: """`StringList` of input lines (without newlines).
							πF.SetLineno(146)
							// line 149: self.input_offset = 0
							πF.SetLineno(149)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput_offset, πTemp001); πE != nil {
								continue
							}
							// line 150: """Offset of `self.input_lines` from the beginning of the file."""
							πF.SetLineno(150)
							// line 152: self.line = None
							πF.SetLineno(152)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp002); πE != nil {
								continue
							}
							// line 153: """Current input line."""
							πF.SetLineno(153)
							// line 155: self.line_offset = -1
							πF.SetLineno(155)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_offset, πTemp002); πE != nil {
								continue
							}
							// line 156: """Current input line offset from beginning of `self.input_lines`."""
							πF.SetLineno(156)
							// line 158: self.debug = debug
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdebug); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdebug, πTemp001); πE != nil {
								continue
							}
							// line 159: """Debugging mode on/off."""
							πF.SetLineno(159)
							// line 161: self.initial_state = initial_state
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µinitial_state, "initial_state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinitial_state); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinitial_state, πTemp001); πE != nil {
								continue
							}
							// line 162: """The name of the initial state (key to `self.states`)."""
							πF.SetLineno(162)
							// line 164: self.current_state = initial_state
							πF.SetLineno(164)
							if πE = πg.CheckLocal(πF, µinitial_state, "initial_state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinitial_state); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_state, πTemp001); πE != nil {
								continue
							}
							// line 165: """The name of the current state (key to `self.states`)."""
							πF.SetLineno(165)
							// line 167: self.states = {}
							πF.SetLineno(167)
							πTemp003 = πg.NewDict()
							πTemp001 = πTemp003.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstates, πTemp002); πE != nil {
								continue
							}
							// line 168: """Mapping of {state_name: State_object}."""
							πF.SetLineno(168)
							// line 170: self.add_states(state_classes)
							πF.SetLineno(170)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate_classes, "state_classes"); πE != nil {
								continue
							}
							πTemp004[0] = µstate_classes
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_states, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 172: self.observers = []
							πF.SetLineno(172)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßobservers, πTemp002); πE != nil {
								continue
							}
							// line 173: """List of bound methods or functions to call whenever the current
							πF.SetLineno(173)
							// line 177: self._stderr = ErrorOutput()
							πF.SetLineno(177)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stderr, πTemp001); πE != nil {
								continue
							}
							// line 178: """Wrapper around sys.stderr catching en-/decoding errors"""
							πF.SetLineno(178)
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
					// line 135: """
					πF.SetLineno(135)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Initialize a `StateMachine` object; add state objects.\n\n        Parameters:\n\n        - `state_classes`: a list of `State` (sub)classes.\n        - `initial_state`: a string, the class name of the initial state.\n        - `debug`: a boolean; produce verbose output if true (nonzero).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 181: def unlink(self):
					πF.SetLineno(181)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("unlink", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate *πg.Object = πg.UnboundLocal
						_ = µstate
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
							// line 182: """Remove circular references to objects no longer required."""
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstates, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µstate = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 184: state.unlink()
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßunlink, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 185: self.states = None
							πF.SetLineno(185)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstates, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßunlink.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 182: """Remove circular references to objects no longer required."""
					πF.SetLineno(182)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Remove circular references to objects no longer required.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßunlink); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 187: def run(self, input_lines, input_offset=0, context=None,
					πF.SetLineno(187)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "input_lines", Def: nil}
					πTemp002[2] = πg.Param{Name: "input_offset", Def: πg.NewInt(0).ToObject()}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "context", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "input_source", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "initial_state", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinput_lines *πg.Object = πArgs[1]
						_ = µinput_lines
						var µinput_offset *πg.Object = πArgs[2]
						_ = µinput_offset
						var µcontext *πg.Object = πArgs[3]
						_ = µcontext
						var µinput_source *πg.Object = πArgs[4]
						_ = µinput_source
						var µinitial_state *πg.Object = πArgs[5]
						_ = µinitial_state
						var µtransitions *πg.Object = πg.UnboundLocal
						_ = µtransitions
						var µresults *πg.Object = πg.UnboundLocal
						_ = µresults
						var µstate *πg.Object = πg.UnboundLocal
						_ = µstate
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µnext_state *πg.Object = πg.UnboundLocal
						_ = µnext_state
						var µexception *πg.Object = πg.UnboundLocal
						_ = µexception
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 17:
								goto Label17
							case 11:
								goto Label11
							case 12:
								goto Label12
							case 15:
								goto Label15
							default:
								panic("unexpected function state")
							}
							// line 189: """
							πF.SetLineno(189)
							// line 210: self.runtime_init()
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßruntime_init, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinput_lines, "input_lines"); πE != nil {
								continue
							}
							πTemp003[0] = µinput_lines
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringList); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 211: if isinstance(input_lines, StringList):
							πF.SetLineno(211)
						Label1:
							// line 212: self.input_lines = input_lines
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µinput_lines, "input_lines"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinput_lines); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput_lines, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 214: self.input_lines = StringList(input_lines, source=input_source)
							πF.SetLineno(214)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinput_lines, "input_lines"); πE != nil {
								continue
							}
							πTemp003[0] = µinput_lines
							if πE = πg.CheckLocal(πF, µinput_source, "input_source"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"source", µinput_source},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStringList); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput_lines, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 215: self.input_offset = input_offset
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µinput_offset, "input_offset"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinput_offset); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput_offset, πTemp001); πE != nil {
								continue
							}
							// line 216: self.line_offset = -1
							πF.SetLineno(216)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_offset, πTemp002); πE != nil {
								continue
							}
							// line 217: self.current_state = initial_state or self.initial_state
							πF.SetLineno(217)
							if πE = πg.CheckLocal(πF, µinitial_state, "initial_state"); πE != nil {
								continue
							}
							πTemp001 = µinitial_state
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinitial_state, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label4:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_state, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 218: if self.debug:
							πF.SetLineno(218)
						Label5:
							// line 219: print((
							πF.SetLineno(219)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp008
							if πTemp008, πE = πg.GetAttr(πF, πg.NewUnicode("\n| ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002 = πg.NewTuple2(πTemp006, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("\nStateMachine.run: input_lines (line_offset=%s):\n| %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label6
						Label6:
							// line 222: transitions = None
							πF.SetLineno(222)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtransitions = πTemp001
							// line 223: results = []
							πF.SetLineno(223)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresults = πTemp001
							// line 224: state = self.get_state()
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_state, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstate = πTemp002
							// line 225: try:
							πF.SetLineno(225)
							πF.PushCheckpoint(8)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 226: if self.debug:
							πF.SetLineno(226)
						Label9:
							// line 227: print('\nStateMachine.run: bof transition', file=self._stderr)
							πF.SetLineno(227)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\nStateMachine.run: bof transition").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label10
						Label10:
							// line 228: context, result = state.bof(context)
							πF.SetLineno(228)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp003[0] = µcontext
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßbof, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
								continue
							}
							µcontext = πTemp001
							µresult = πTemp006
							// line 229: results.extend(result)
							πF.SetLineno(229)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresults, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 230: while True:
							πF.SetLineno(230)
							πF.PushCheckpoint(12)
							πTemp004 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(11)
							// line 231: try:
							πF.SetLineno(231)
							πF.PushCheckpoint(15)
							// line 232: try:
							πF.SetLineno(232)
							πF.PushCheckpoint(17)
							// line 233: self.next_line()
							πF.SetLineno(233)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnext_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label18
							}
							goto Label19
							// line 234: if self.debug:
							πF.SetLineno(234)
						Label18:
							// line 235: source, offset = self.input_lines.info(
							πF.SetLineno(235)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
								continue
							}
							µsource = πTemp002
							µoffset = πTemp006
							// line 237: print((
							πF.SetLineno(237)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µsource, µoffset, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("\nStateMachine.run: line (source=%r, offset=%r):\n| %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label19
						Label19:
							// line 241: context, next_state, result = self.check_line(
							πF.SetLineno(241)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp003[0] = µcontext
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003[1] = µstate
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp003[2] = µtransitions
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µcontext = πTemp001
							µnext_state = πTemp006
							µresult = πTemp008
							πF.PopCheckpoint()
							// line 252: results.extend(result)
							πF.SetLineno(252)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresults, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label16
						Label17:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label20
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 243: except EOFError:
							πF.SetLineno(243)
						Label20:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label21
							}
							goto Label22
							// line 244: if self.debug:
							πF.SetLineno(244)
						Label21:
							// line 245: print((
							πF.SetLineno(245)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.run: %s.eof transition").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label22
						Label22:
							// line 248: result = state.eof(context)
							πF.SetLineno(248)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp003[0] = µcontext
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßeof, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µresult = πTemp002
							// line 249: results.extend(result)
							πF.SetLineno(249)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp003[0] = µresult
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresults, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 250: break
							πF.SetLineno(250)
							πTemp004 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label16
						Label16:
							πF.PopCheckpoint()
							// line 275: transitions = None
							πF.SetLineno(275)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtransitions = πTemp001
							goto Label14
						Label15:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTransitionCorrection); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label23
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStateCorrection); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label24
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 253: except TransitionCorrection as exception:
							πF.SetLineno(253)
						Label23:
							µexception = πTemp011.ToObject()
							// line 254: self.previous_line() # back up for another try
							πF.SetLineno(254)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßprevious_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 255: transitions = (exception.args[0],)
							πF.SetLineno(255)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µexception, ßargs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(πTemp006).ToObject()
							µtransitions = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label25
							}
							goto Label26
							// line 256: if self.debug:
							πF.SetLineno(256)
						Label25:
							// line 257: print((
							πF.SetLineno(257)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ß__name__, nil); πE != nil {
								continue
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µtransitions, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp008, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.run: TransitionCorrection to state \"%s\", transition %s.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label26
						Label26:
							// line 261: continue
							πF.SetLineno(261)
							continue
							πF.RestoreExc(nil, nil)
							goto Label14
							// line 262: except StateCorrection as exception:
							πF.SetLineno(262)
						Label24:
							µexception = πTemp011.ToObject()
							// line 263: self.previous_line() # back up for another try
							πF.SetLineno(263)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßprevious_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 264: next_state = exception.args[0]
							πF.SetLineno(264)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µexception, ßargs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µnext_state = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µexception, ßargs, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label27
							}
							goto Label28
							// line 265: if len(exception.args) == 1:
							πF.SetLineno(265)
						Label27:
							// line 266: transitions = None
							πF.SetLineno(266)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µtransitions = πTemp001
							goto Label29
						Label28:
							// line 268: transitions = (exception.args[1],)
							πF.SetLineno(268)
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µexception, ßargs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(πTemp006).ToObject()
							µtransitions = πTemp001
							goto Label29
						Label29:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label30
							}
							goto Label31
							// line 269: if self.debug:
							πF.SetLineno(269)
						Label30:
							// line 270: print((
							πF.SetLineno(270)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µtransitions, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µnext_state, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.run: StateCorrection to state \"%s\", transition %s.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label31
						Label31:
							πF.RestoreExc(nil, nil)
							goto Label14
						Label14:
							// line 276: state = self.get_state(next_state)
							πF.SetLineno(276)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp003[0] = µnext_state
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_state, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µstate = πTemp002
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							goto Label32
							// line 277: except:
							πF.SetLineno(277)
						Label32:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label33
							}
							goto Label34
							// line 278: if self.debug:
							πF.SetLineno(278)
						Label33:
							// line 279: self.error()
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label34
						Label34:
							// line 280: raise
							πF.SetLineno(280)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 281: self.observers = []
							πF.SetLineno(281)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßobservers, πTemp002); πE != nil {
								continue
							}
							// line 282: return results
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							πR = µresults
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
					// line 189: """
					πF.SetLineno(189)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Run the state machine on `input_lines`. Return results (a list).\n\n        Reset `self.line_offset` and `self.current_state`. Run the\n        beginning-of-file transition. Input one line at a time and check for a\n        matching transition. If a match is found, call the transition method\n        and possibly change the state. Store the context returned by the\n        transition method to be passed on to the next transition matched.\n        Accumulate the results returned by the transition methods in a list.\n        Run the end-of-file transition. Finally, return the accumulated\n        results.\n\n        Parameters:\n\n        - `input_lines`: a list of strings without newlines, or `StringList`.\n        - `input_offset`: the line offset of `input_lines` from the beginning\n          of the file.\n        - `context`: application-specific storage.\n        - `input_source`: name or path of source of `input_lines`.\n        - `initial_state`: name of initial state.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßrun); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 284: def get_state(self, next_state=None):
					πF.SetLineno(284)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "next_state", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("get_state", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnext_state *πg.Object = πArgs[1]
						_ = µnext_state
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 285: """
							πF.SetLineno(285)
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µnext_state); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 292: if next_state:
							πF.SetLineno(292)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp001 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrent_state, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, µnext_state, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label3:
							if πTemp001, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label4
							}
							goto Label5
							// line 293: if self.debug and next_state != self.current_state:
							πF.SetLineno(293)
						Label4:
							// line 294: print((
							πF.SetLineno(294)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrent_state, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßabs_line_number, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(πTemp004, µnext_state, πTemp007).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.get_state: Changing state from \"%s\" to \"%s\" (input line %s).").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label5
						Label5:
							// line 299: self.current_state = next_state
							πF.SetLineno(299)
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnext_state); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_state, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 300: try:
							πF.SetLineno(300)
							πF.PushCheckpoint(7)
							// line 301: return self.states[self.current_state]
							πF.SetLineno(301)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcurrent_state, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstates, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label8
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 302: except KeyError:
							πF.SetLineno(302)
						Label8:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcurrent_state, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnknownStateError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 303: raise UnknownStateError(self.current_state)
							πF.SetLineno(303)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_state.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 285: """
					πF.SetLineno(285)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Return current state object; set it first if `next_state` given.\n\n        Parameter `next_state`: a string, the name of the next state.\n\n        Exception: `UnknownStateError` raised if `next_state` unknown.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßget_state); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 305: def next_line(self, n=1):
					πF.SetLineno(305)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(1).ToObject()}
					πTemp006 = πg.NewFunction(πg.NewCode("next_line", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
						var πTemp001 *πg.Object
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 3:
								goto Label3
							default:
								panic("unexpected function state")
							}
							// line 306: """Load `self.line` with the `n`'th next line and return it."""
							πF.SetLineno(306)
							// line 307: try:
							πF.SetLineno(307)
							πF.PushCheckpoint(1)
							// line 308: try:
							πF.SetLineno(308)
							πF.PushCheckpoint(3)
							// line 309: self.line_offset += n
							πF.SetLineno(309)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_offset, πTemp002); πE != nil {
								continue
							}
							// line 310: self.line = self.input_lines[self.line_offset]
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label2
						Label3:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 311: except IndexError:
							πF.SetLineno(311)
						Label4:
							// line 312: self.line = None
							πF.SetLineno(312)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
								continue
							}
							// line 313: raise EOFError
							πF.SetLineno(313)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label2
						Label2:
							// line 314: return self.line
							πF.SetLineno(314)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
							// line 316: self.notify_observers()
							πF.SetLineno(316)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnotify_observers, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004 != nil {
								πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßnext_line.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 306: """Load `self.line` with the `n`'th next line and return it."""
					πF.SetLineno(306)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Load `self.line` with the `n`'th next line and return it.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßnext_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 318: def is_next_line_blank(self):
					πF.SetLineno(318)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("is_next_line_blank", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 319: """Return 1 if the next line is blank or non-existant."""
							πF.SetLineno(319)
							// line 320: try:
							πF.SetLineno(320)
							πF.PushCheckpoint(2)
							// line 321: return not self.input_lines[self.line_offset + 1].strip()
							πF.SetLineno(321)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 322: except IndexError:
							πF.SetLineno(322)
						Label3:
							// line 323: return 1
							πF.SetLineno(323)
							πR = πg.NewInt(1).ToObject()
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßis_next_line_blank.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 319: """Return 1 if the next line is blank or non-existant."""
					πF.SetLineno(319)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return 1 if the next line is blank or non-existant.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßis_next_line_blank); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 325: def at_eof(self):
					πF.SetLineno(325)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("at_eof", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp006 *πg.Object
						_ = πTemp006
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
							// line 326: """Return 1 if the input is at or past end-of-file."""
							πF.SetLineno(326)
							// line 327: return self.line_offset >= len(self.input_lines) - 1
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Sub(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_eof.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 326: """Return 1 if the input is at or past end-of-file."""
					πF.SetLineno(326)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return 1 if the input is at or past end-of-file.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßat_eof); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 329: def at_bof(self):
					πF.SetLineno(329)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("at_bof", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 330: """Return 1 if the input is at or before beginning-of-file."""
							πF.SetLineno(330)
							// line 331: return self.line_offset <= 0
							πF.SetLineno(331)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_bof.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 330: """Return 1 if the input is at or before beginning-of-file."""
					πF.SetLineno(330)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Return 1 if the input is at or before beginning-of-file.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßat_bof); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 333: def previous_line(self, n=1):
					πF.SetLineno(333)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(1).ToObject()}
					πTemp010 = πg.NewFunction(πg.NewCode("previous_line", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 334: """Load `self.line` with the `n`'th previous line and return it."""
							πF.SetLineno(334)
							// line 335: self.line_offset -= n
							πF.SetLineno(335)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_offset, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 336: if self.line_offset < 0:
							πF.SetLineno(336)
						Label1:
							// line 337: self.line = None
							πF.SetLineno(337)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp002); πE != nil {
								continue
							}
							goto Label3
						Label2:
							// line 339: self.line = self.input_lines[self.line_offset]
							πF.SetLineno(339)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp001); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 340: self.notify_observers()
							πF.SetLineno(340)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnotify_observers, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 341: return self.line
							πF.SetLineno(341)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprevious_line.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 334: """Load `self.line` with the `n`'th previous line and return it."""
					πF.SetLineno(334)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Load `self.line` with the `n`'th previous line and return it.").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßprevious_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 343: def goto_line(self, line_offset):
					πF.SetLineno(343)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "line_offset", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("goto_line", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µline_offset *πg.Object = πArgs[1]
						_ = µline_offset
						var πTemp001 *πg.Object
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 3:
								goto Label3
							default:
								panic("unexpected function state")
							}
							// line 344: """Jump to absolute line offset `line_offset`, load and return it."""
							πF.SetLineno(344)
							// line 345: try:
							πF.SetLineno(345)
							πF.PushCheckpoint(1)
							// line 346: try:
							πF.SetLineno(346)
							πF.PushCheckpoint(3)
							// line 347: self.line_offset = line_offset - self.input_offset
							πF.SetLineno(347)
							if πE = πg.CheckLocal(πF, µline_offset, "line_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µline_offset, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_offset, πTemp002); πE != nil {
								continue
							}
							// line 348: self.line = self.input_lines[self.line_offset]
							πF.SetLineno(348)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label2
						Label3:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 349: except IndexError:
							πF.SetLineno(349)
						Label4:
							// line 350: self.line = None
							πF.SetLineno(350)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
								continue
							}
							// line 351: raise EOFError
							πF.SetLineno(351)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label2
						Label2:
							// line 352: return self.line
							πF.SetLineno(352)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
							// line 354: self.notify_observers()
							πF.SetLineno(354)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnotify_observers, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004 != nil {
								πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßgoto_line.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 344: """Jump to absolute line offset `line_offset`, load and return it."""
					πF.SetLineno(344)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Jump to absolute line offset `line_offset`, load and return it.").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßgoto_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 356: def get_source(self, line_offset):
					πF.SetLineno(356)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "line_offset", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("get_source", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µline_offset *πg.Object = πArgs[1]
						_ = µline_offset
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
							// line 357: """Return source of line at absolute line offset `line_offset`."""
							πF.SetLineno(357)
							// line 358: return self.input_lines.source(line_offset - self.input_offset)
							πF.SetLineno(358)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline_offset, "line_offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µline_offset, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsource, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_source.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 357: """Return source of line at absolute line offset `line_offset`."""
					πF.SetLineno(357)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Return source of line at absolute line offset `line_offset`.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßget_source); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 360: def abs_line_offset(self):
					πF.SetLineno(360)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("abs_line_offset", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 361: """Return line offset of current line, from beginning of file."""
							πF.SetLineno(361)
							// line 362: return self.line_offset + self.input_offset
							πF.SetLineno(362)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßabs_line_offset.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 361: """Return line offset of current line, from beginning of file."""
					πF.SetLineno(361)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("Return line offset of current line, from beginning of file.").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßabs_line_offset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 364: def abs_line_number(self):
					πF.SetLineno(364)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("abs_line_number", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 365: """Return line number of current line (counting from 1)."""
							πF.SetLineno(365)
							// line 366: return self.line_offset + self.input_offset + 1
							πF.SetLineno(366)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßabs_line_number.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 365: """Return line number of current line (counting from 1)."""
					πF.SetLineno(365)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Return line number of current line (counting from 1).").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßabs_line_number); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 368: def get_source_and_line(self, lineno=None):
					πF.SetLineno(368)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "lineno", Def: πTemp016}
					πTemp015 = πg.NewFunction(πg.NewCode("get_source_and_line", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlineno *πg.Object = πArgs[1]
						_ = µlineno
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µsrc *πg.Object = πg.UnboundLocal
						_ = µsrc
						var µsrcoffset *πg.Object = πg.UnboundLocal
						_ = µsrcoffset
						var µsrcline *πg.Object = πg.UnboundLocal
						_ = µsrcline
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 369: """Return (source, line) tuple for current or given line number.
							πF.SetLineno(369)
							if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µlineno == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 377: if lineno is None:
							πF.SetLineno(377)
						Label1:
							// line 378: offset = self.line_offset
							πF.SetLineno(378)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							µoffset = πTemp001
							goto Label3
						Label2:
							// line 380: offset = lineno - self.input_offset - 1
							πF.SetLineno(380)
							if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µlineno, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µoffset = πTemp001
							goto Label3
						Label3:
							// line 381: try:
							πF.SetLineno(381)
							πF.PushCheckpoint(5)
							// line 382: src, srcoffset = self.input_lines.info(offset)
							πF.SetLineno(382)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp005[0] = µoffset
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µsrc = πTemp002
							µsrcoffset = πTemp004
							// line 383: srcline = srcoffset + 1
							πF.SetLineno(383)
							if πE = πg.CheckLocal(πF, µsrcoffset, "srcoffset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µsrcoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µsrcline = πTemp001
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 384: except (TypeError):
							πF.SetLineno(384)
						Label6:
							// line 386: src, srcline = self.get_source_and_line(offset + self.input_offset)
							πF.SetLineno(386)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µoffset, πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_source_and_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µsrc = πTemp001
							µsrcline = πTemp004
							// line 387: return src, srcline + 1
							πF.SetLineno(387)
							if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsrcline, "srcline"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µsrcline, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µsrc, πTemp002).ToObject()
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
							// line 388: except (IndexError):  # `offset` is off the list
							πF.SetLineno(388)
						Label7:
							// line 389: src, srcline = None, None
							πF.SetLineno(389)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µsrc = πTemp002
							µsrcline = πTemp004
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 393: return (src, srcline)
							πF.SetLineno(393)
							if πE = πg.CheckLocal(πF, µsrc, "src"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsrcline, "srcline"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µsrc, µsrcline).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_source_and_line.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 369: """Return (source, line) tuple for current or given line number.
					πF.SetLineno(369)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Return (source, line) tuple for current or given line number.\n\n        Looks up the source and line number in the `self.input_lines`\n        StringList instance to count for included source files.\n\n        If the optional argument `lineno` is given, convert it from an\n        absolute line number to the corresponding (source, line) pair.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßget_source_and_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 395: def insert_input(self, input_lines, source):
					πF.SetLineno(395)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "input_lines", Def: nil}
					πTemp002[2] = πg.Param{Name: "source", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("insert_input", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinput_lines *πg.Object = πArgs[1]
						_ = µinput_lines
						var µsource *πg.Object = πArgs[2]
						_ = µsource
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
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
							// line 396: self.input_lines.insert(self.line_offset + 1, '',
							πF.SetLineno(396)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("internal padding after ").ToObject(), µsource); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinput_lines, "input_lines"); πE != nil {
								continue
							}
							πTemp004[0] = µinput_lines
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp006 = πg.KWArgs{
								{"source", πTemp002},
								{"offset", πTemp005},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 399: self.input_lines.insert(self.line_offset + 1, '',
							πF.SetLineno(399)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("internal padding before ").ToObject(), µsource); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"source", πTemp002},
								{"offset", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 402: self.input_lines.insert(self.line_offset + 2,
							πF.SetLineno(402)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinput_lines, "input_lines"); πE != nil {
								continue
							}
							πTemp004[0] = µinput_lines
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp004[1] = µsource
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStringList); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßinsert_input.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 405: def get_text_block(self, flush_left=False):
					πF.SetLineno(405)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "flush_left", Def: πTemp018}
					πTemp017 = πg.NewFunction(πg.NewCode("get_text_block", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µflush_left *πg.Object = πArgs[1]
						_ = µflush_left
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
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
							// line 406: """
							πF.SetLineno(406)
							// line 413: try:
							πF.SetLineno(413)
							πF.PushCheckpoint(2)
							// line 414: block = self.input_lines.get_text_block(self.line_offset,
							πF.SetLineno(414)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µflush_left, "flush_left"); πE != nil {
								continue
							}
							πTemp001[1] = µflush_left
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_text_block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µblock = πTemp002
							// line 416: self.next_line(len(block) - 1)
							πF.SetLineno(416)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp004[0] = µblock
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnext_line, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 417: return block
							πF.SetLineno(417)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πR = µblock
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnexpectedIndentationError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 418: except UnexpectedIndentationError as err:
							πF.SetLineno(418)
						Label3:
							µerr = πTemp006.ToObject()
							// line 419: block = err.args[0]
							πF.SetLineno(419)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µerr, ßargs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							µblock = πTemp003
							// line 420: self.next_line(len(block) - 1) # advance to last line of block
							πF.SetLineno(420)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp004[0] = µblock
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnext_line, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 421: raise
							πF.SetLineno(421)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_text_block.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 406: """
					πF.SetLineno(406)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("\n        Return a contiguous block of text.\n\n        If `flush_left` is true, raise `UnexpectedIndentationError` if an\n        indented line is encountered before the text block ends (with a blank\n        line).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßget_text_block); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 423: def check_line(self, context, state, transitions=None):
					πF.SetLineno(423)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp002[2] = πg.Param{Name: "state", Def: nil}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "transitions", Def: πTemp019}
					πTemp018 = πg.NewFunction(πg.NewCode("check_line", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcontext *πg.Object = πArgs[1]
						_ = µcontext
						var µstate *πg.Object = πArgs[2]
						_ = µstate
						var µtransitions *πg.Object = πArgs[3]
						_ = µtransitions
						var µstate_correction *πg.Object = πg.UnboundLocal
						_ = µstate_correction
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µpattern *πg.Object = πg.UnboundLocal
						_ = µpattern
						var µmethod *πg.Object = πg.UnboundLocal
						_ = µmethod
						var µnext_state *πg.Object = πg.UnboundLocal
						_ = µnext_state
						var µmatch *πg.Object = πg.UnboundLocal
						_ = µmatch
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
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5:
								goto Label5
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 424: """
							πF.SetLineno(424)
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtransitions == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 443: if transitions is None:
							πF.SetLineno(443)
						Label1:
							// line 444: transitions =  state.transition_order
							πF.SetLineno(444)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßtransition_order, nil); πE != nil {
								continue
							}
							µtransitions = πTemp001
							goto Label2
						Label2:
							// line 445: state_correction = None
							πF.SetLineno(445)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µstate_correction = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 446: if self.debug:
							πF.SetLineno(446)
						Label3:
							// line 447: print((
							πF.SetLineno(447)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µstate, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, µtransitions).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.check_line: state=\"%s\", transitions=%r.").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtransitions); πE != nil {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp003 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label7
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µname = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(5)
							// line 451: pattern, method, next_state = state.transitions[name]
							πF.SetLineno(451)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ßtransitions, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}}}, πTemp005); πE != nil {
								continue
							}
							µpattern = πTemp002
							µmethod = πTemp006
							µnext_state = πTemp009
							// line 452: match = pattern.match(self.line)
							πF.SetLineno(452)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpattern, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmatch = πTemp005
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µmatch); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label8
							}
							goto Label9
							// line 453: if match:
							πF.SetLineno(453)
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label10
							}
							goto Label11
							// line 454: if self.debug:
							πF.SetLineno(454)
						Label10:
							// line 455: print((
							πF.SetLineno(455)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp006, ß__name__, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µname, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.check_line: Matched transition \"%s\" in state \"%s\".").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label11
						Label11:
							// line 459: return method(match, context, next_state)
							πF.SetLineno(459)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							πTemp004[0] = µmatch
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp004[1] = µcontext
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp004[2] = µnext_state
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp002, πE = µmethod.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label9
						Label9:
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label12
							}
							goto Label13
							// line 461: if self.debug:
							πF.SetLineno(461)
						Label12:
							// line 462: print((
							πF.SetLineno(462)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µstate, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\nStateMachine.check_line: No match in state \"%s\".").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"file", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label13:
							// line 465: return state.no_match(context, transitions)
							πF.SetLineno(465)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp004[0] = µcontext
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp004[1] = µtransitions
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßno_match, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp005
							continue
						Label7:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_line.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 424: """
					πF.SetLineno(424)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("\n        Examine one line of input for a transition match & execute its method.\n\n        Parameters:\n\n        - `context`: application-dependent storage.\n        - `state`: a `State` object, the current state.\n        - `transitions`: an optional ordered list of transition names to try,\n          instead of ``state.transition_order``.\n\n        Return the values returned by the transition method:\n\n        - context: possibly modified from the parameter `context`;\n        - next state name (`State` subclass name);\n        - the result output of the transition, a list.\n\n        When there is no match, ``state.no_match()`` is called and its return\n        value is returned.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_line); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
						continue
					}
					// line 467: def add_state(self, state_class):
					πF.SetLineno(467)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state_class", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("add_state", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate_class *πg.Object = πArgs[1]
						_ = µstate_class
						var µstatename *πg.Object = πg.UnboundLocal
						_ = µstatename
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
							// line 468: """
							πF.SetLineno(468)
							// line 474: statename = state_class.__name__
							πF.SetLineno(474)
							if πE = πg.CheckLocal(πF, µstate_class, "state_class"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate_class, ß__name__, nil); πE != nil {
								continue
							}
							µstatename = πTemp001
							if πE = πg.CheckLocal(πF, µstatename, "statename"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstates, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µstatename); πE != nil {
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
							// line 475: if statename in self.states:
							πF.SetLineno(475)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstatename, "statename"); πE != nil {
								continue
							}
							πTemp004[0] = µstatename
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDuplicateStateError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 476: raise DuplicateStateError(statename)
							πF.SetLineno(476)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 477: self.states[statename] = state_class(self, self.debug)
							πF.SetLineno(477)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µstate_class, "state_class"); πE != nil {
								continue
							}
							if πTemp001, πE = µstate_class.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstates, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstatename, "statename"); πE != nil {
								continue
							}
							πTemp006 = µstatename
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_state.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 468: """
					πF.SetLineno(468)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("\n        Initialize & add a `state_class` (`State` subclass) object.\n\n        Exception: `DuplicateStateError` raised if `state_class` was already\n        added.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßadd_state); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
						continue
					}
					// line 479: def add_states(self, state_classes):
					πF.SetLineno(479)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state_classes", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("add_states", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate_classes *πg.Object = πArgs[1]
						_ = µstate_classes
						var µstate_class *πg.Object = πg.UnboundLocal
						_ = µstate_class
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
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
							// line 480: """
							πF.SetLineno(480)
							if πE = πg.CheckLocal(πF, µstate_classes, "state_classes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µstate_classes); πE != nil {
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
								µstate_class = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 484: self.add_state(state_class)
							πF.SetLineno(484)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate_class, "state_class"); πE != nil {
								continue
							}
							πTemp005[0] = µstate_class
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßadd_state, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßadd_states.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 480: """
					πF.SetLineno(480)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("\n        Add `state_classes` (a list of `State` subclasses).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßadd_states); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
						continue
					}
					// line 486: def runtime_init(self):
					πF.SetLineno(486)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("runtime_init", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate *πg.Object = πg.UnboundLocal
						_ = µstate
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
							// line 487: """
							πF.SetLineno(487)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstates, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µstate = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 491: state.runtime_init()
							πF.SetLineno(491)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßruntime_init, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßruntime_init.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 487: """
					πF.SetLineno(487)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("\n        Initialize `self.states`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßruntime_init); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
						continue
					}
					// line 493: def error(self):
					πF.SetLineno(493)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("error", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtype *πg.Object = πg.UnboundLocal
						_ = µtype
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µmodule *πg.Object = πg.UnboundLocal
						_ = µmodule
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µfunction *πg.Object = πg.UnboundLocal
						_ = µfunction
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 494: """Report error details."""
							πF.SetLineno(494)
							// line 495: type, value, module, line, function = _exception_data()
							πF.SetLineno(495)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_exception_data); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
								continue
							}
							µtype = πTemp001
							µvalue = πTemp003
							µmodule = πTemp004
							µline = πTemp005
							µfunction = πTemp006
							// line 496: print(u'%s: %s' % (type, value), file=self._stderr)
							πF.SetLineno(496)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µtype, µvalue).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("%s: %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 497: print('input line %s' % (self.abs_line_number()), file=self._stderr)
							πF.SetLineno(497)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßabs_line_number, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("input line %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 498: print((u'module %s, line %s, function %s' %
							πF.SetLineno(498)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µmodule, µline, µfunction).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewUnicode("module %s, line %s, function %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"file", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßerror.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 494: """Report error details."""
					πF.SetLineno(494)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("Report error details.").ToObject()); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßerror); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
						continue
					}
					// line 501: def attach_observer(self, observer):
					πF.SetLineno(501)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "observer", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("attach_observer", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobserver *πg.Object = πArgs[1]
						_ = µobserver
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
							// line 502: """
							πF.SetLineno(502)
							// line 506: self.observers.append(observer)
							πF.SetLineno(506)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobserver, "observer"); πE != nil {
								continue
							}
							πTemp001[0] = µobserver
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßobservers, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßattach_observer.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 502: """
					πF.SetLineno(502)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("\n        The `observer` parameter is a function or bound method which takes two\n        arguments, the source and offset of the current line.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßattach_observer); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 508: def detach_observer(self, observer):
					πF.SetLineno(508)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "observer", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("detach_observer", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobserver *πg.Object = πArgs[1]
						_ = µobserver
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
							// line 509: self.observers.remove(observer)
							πF.SetLineno(509)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobserver, "observer"); πE != nil {
								continue
							}
							πTemp001[0] = µobserver
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßobservers, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßremove, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdetach_observer.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 511: def notify_observers(self):
					πF.SetLineno(511)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("notify_observers", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobserver *πg.Object = πg.UnboundLocal
						_ = µobserver
						var µinfo *πg.Object = πg.UnboundLocal
						_ = µinfo
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
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
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßobservers, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp003 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µobserver = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 513: try:
							πF.SetLineno(513)
							πF.PushCheckpoint(5)
							// line 514: info = self.input_lines.info(self.line_offset)
							πF.SetLineno(514)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µinfo = πTemp002
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 515: except IndexError:
							πF.SetLineno(515)
						Label6:
							// line 516: info = (None, None)
							πF.SetLineno(516)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, πTemp009).ToObject()
							µinfo = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 517: observer(*info)
							πF.SetLineno(517)
							if πE = πg.CheckLocal(πF, µinfo, "info"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobserver, "observer"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µobserver, nil, µinfo, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnotify_observers.ToObject(), πTemp025); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StateMachine").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStateMachine.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 520: class State(object):
			πF.SetLineno(520)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("State", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 522: """
					πF.SetLineno(522)
					// line 522: """
					πF.SetLineno(522)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    State superclass. Contains a list of transitions, and transition methods.\n\n    Transition methods all have the same signature. They take 3 parameters:\n\n    - An `re` match object. ``match.string`` contains the matched input line,\n      ``match.start()`` gives the start index of the match, and\n      ``match.end()`` gives the end index.\n    - A context object, whose meaning is application-defined (initial value\n      ``None``). It can be used to store any information required by the state\n      machine, and the retured context is passed on to the next transition\n      method unchanged.\n    - The name of the next state, a string, taken from the transitions list;\n      normally it is returned unchanged, but it may be altered by the\n      transition method if necessary.\n\n    Transition methods all return a 3-tuple:\n\n    - A context object, as (potentially) modified by the transition method.\n    - The next state name (a return value of ``None`` means no state change).\n    - The processing result, a list, which is accumulated by the state\n      machine.\n\n    Transition methods may raise an `EOFError` to cut processing short.\n\n    There are two implicit transitions, and corresponding transition methods\n    are defined: `bof()` handles the beginning-of-file, and `eof()` handles\n    the end-of-file. These methods have non-standard signatures and return\n    values. `bof()` returns the initial context and results, and may be used\n    to return a header string, or do any other processing needed. `eof()`\n    should handle any remaining context and wrap things up; it returns the\n    final processing result.\n\n    Typical applications need only subclass `State` (or a subclass), set the\n    `patterns` and `initial_transitions` class attributes, and provide\n    corresponding transition methods. The default object initialization will\n    take care of constructing the list of transitions.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 561: patterns = None
					πF.SetLineno(561)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßpatterns.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 562: """
					πF.SetLineno(562)
					// line 567: initial_transitions = None
					πF.SetLineno(567)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßinitial_transitions.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 568: """
					πF.SetLineno(568)
					// line 574: nested_sm = None
					πF.SetLineno(574)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnested_sm.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 575: """
					πF.SetLineno(575)
					// line 582: nested_sm_kwargs = None
					πF.SetLineno(582)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnested_sm_kwargs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 583: """
					πF.SetLineno(583)
					// line 597: def __init__(self, state_machine, debug=False):
					πF.SetLineno(597)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state_machine", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "debug", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate_machine *πg.Object = πArgs[1]
						_ = µstate_machine
						var µdebug *πg.Object = πArgs[2]
						_ = µdebug
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
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
							// line 598: """
							πF.SetLineno(598)
							// line 607: self.transition_order = []
							πF.SetLineno(607)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtransition_order, πTemp003); πE != nil {
								continue
							}
							// line 608: """A list of transition names in search order."""
							πF.SetLineno(608)
							// line 610: self.transitions = {}
							πF.SetLineno(610)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtransitions, πTemp003); πE != nil {
								continue
							}
							// line 611: """
							πF.SetLineno(611)
							// line 619: self.add_initial_transitions()
							πF.SetLineno(619)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_initial_transitions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 621: self.state_machine = state_machine
							πF.SetLineno(621)
							if πE = πg.CheckLocal(πF, µstate_machine, "state_machine"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µstate_machine); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate_machine, πTemp002); πE != nil {
								continue
							}
							// line 622: """A reference to the controlling `StateMachine` object."""
							πF.SetLineno(622)
							// line 624: self.debug = debug
							πF.SetLineno(624)
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdebug); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdebug, πTemp002); πE != nil {
								continue
							}
							// line 625: """Debugging mode on/off."""
							πF.SetLineno(625)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnested_sm, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 627: if self.nested_sm is None:
							πF.SetLineno(627)
						Label1:
							// line 628: self.nested_sm = self.state_machine.__class__
							πF.SetLineno(628)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnested_sm, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnested_sm_kwargs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 629: if self.nested_sm_kwargs is None:
							πF.SetLineno(629)
						Label3:
							// line 630: self.nested_sm_kwargs = {'state_classes': [self.__class__],
							πF.SetLineno(630)
							πTemp004 = πg.NewDict()
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πTemp004.SetItem(πF, ßstate_classes.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßinitial_state.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnested_sm_kwargs, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 598: """
					πF.SetLineno(598)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Initialize a `State` object; make & add initial transitions.\n\n        Parameters:\n\n        - `statemachine`: the controlling `StateMachine` object.\n        - `debug`: a boolean; produce verbose output if true.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 633: def runtime_init(self):
					πF.SetLineno(633)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("runtime_init", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 634: """
							πF.SetLineno(634)
							// line 638: pass
							πF.SetLineno(638)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßruntime_init.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 634: """
					πF.SetLineno(634)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Initialize this `State` before running the state machine; called from\n        `self.state_machine.run()`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßruntime_init); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 640: def unlink(self):
					πF.SetLineno(640)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("unlink", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 641: """Remove circular references to objects no longer required."""
							πF.SetLineno(641)
							// line 642: self.state_machine = None
							πF.SetLineno(642)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate_machine, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßunlink.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 641: """Remove circular references to objects no longer required."""
					πF.SetLineno(641)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Remove circular references to objects no longer required.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßunlink); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 644: def add_initial_transitions(self):
					πF.SetLineno(644)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("add_initial_transitions", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnames *πg.Object = πg.UnboundLocal
						_ = µnames
						var µtransitions *πg.Object = πg.UnboundLocal
						_ = µtransitions
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
							// line 645: """Make and add transitions listed in `self.initial_transitions`."""
							πF.SetLineno(645)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinitial_transitions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 646: if self.initial_transitions:
							πF.SetLineno(646)
						Label1:
							// line 647: names, transitions = self.make_transitions(
							πF.SetLineno(647)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinitial_transitions, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmake_transitions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µnames = πTemp001
							µtransitions = πTemp005
							// line 649: self.add_transitions(names, transitions)
							πF.SetLineno(649)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							πTemp003[0] = µnames
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp003[1] = µtransitions
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_transitions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßadd_initial_transitions.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 645: """Make and add transitions listed in `self.initial_transitions`."""
					πF.SetLineno(645)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Make and add transitions listed in `self.initial_transitions`.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßadd_initial_transitions); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 651: def add_transitions(self, names, transitions):
					πF.SetLineno(651)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "names", Def: nil}
					πTemp002[2] = πg.Param{Name: "transitions", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("add_transitions", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnames *πg.Object = πArgs[1]
						_ = µnames
						var µtransitions *πg.Object = πArgs[2]
						_ = µtransitions
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
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
							// line 652: """
							πF.SetLineno(652)
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
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
								µname = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtransitions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp005, µname); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 663: if name in self.transitions:
							πF.SetLineno(663)
						Label4:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							if πTemp004, πE = πg.ResolveGlobal(πF, ßDuplicateTransitionError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 664: raise DuplicateTransitionError(name)
							πF.SetLineno(664)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µtransitions, µname); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 665: if name not in transitions:
							πF.SetLineno(665)
						Label6:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006[0] = µname
							if πTemp004, πE = πg.ResolveGlobal(πF, ßUnknownTransitionError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 666: raise UnknownTransitionError(name)
							πF.SetLineno(666)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label7
						Label7:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 667: self.transition_order[:0] = names
							πF.SetLineno(667)
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnames); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtransition_order, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp001); πE != nil {
								continue
							}
							// line 668: self.transitions.update(transitions)
							πF.SetLineno(668)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp006[0] = µtransitions
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtransitions, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßadd_transitions.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 652: """
					πF.SetLineno(652)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Add a list of transitions to the start of the transition list.\n\n        Parameters:\n\n        - `names`: a list of transition names.\n        - `transitions`: a mapping of names to transition tuples.\n\n        Exceptions: `DuplicateTransitionError`, `UnknownTransitionError`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßadd_transitions); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 670: def add_transition(self, name, transition):
					πF.SetLineno(670)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "transition", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("add_transition", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var µtransition *πg.Object = πArgs[2]
						_ = µtransition
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
							// line 671: """
							πF.SetLineno(671)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransitions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µname); πE != nil {
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
							// line 678: if name in self.transitions:
							πF.SetLineno(678)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßDuplicateTransitionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 679: raise DuplicateTransitionError(name)
							πF.SetLineno(679)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 680: self.transition_order[:0] = [name]
							πF.SetLineno(680)
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtransition_order, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							// line 681: self.transitions[name] = transition
							πF.SetLineno(681)
							if πE = πg.CheckLocal(πF, µtransition, "transition"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtransition); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtransitions, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = µname
							if πE = πg.SetItem(πF, πTemp002, πTemp005, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_transition.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 671: """
					πF.SetLineno(671)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Add a transition to the start of the transition list.\n\n        Parameter `transition`: a ready-made transition 3-tuple.\n\n        Exception: `DuplicateTransitionError`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßadd_transition); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 683: def remove_transition(self, name):
					πF.SetLineno(683)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("remove_transition", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
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
							// line 684: """
							πF.SetLineno(684)
							// line 689: try:
							πF.SetLineno(689)
							πF.PushCheckpoint(2)
							// line 690: del self.transitions[name]
							πF.SetLineno(690)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtransitions, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = µname
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 691: self.transition_order.remove(name)
							πF.SetLineno(691)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[0] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtransition_order, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							goto Label3
							// line 692: except:
							πF.SetLineno(692)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnknownTransitionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 693: raise UnknownTransitionError(name)
							πF.SetLineno(693)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßremove_transition.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 684: """
					πF.SetLineno(684)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Remove a transition by `name`.\n\n        Exception: `UnknownTransitionError`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßremove_transition); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 695: def make_transition(self, name, next_state=None):
					πF.SetLineno(695)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "next_state", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("make_transition", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname *πg.Object = πArgs[1]
						_ = µname
						var µnext_state *πg.Object = πArgs[2]
						_ = µnext_state
						var µpattern *πg.Object = πg.UnboundLocal
						_ = µpattern
						var µmethod *πg.Object = πg.UnboundLocal
						_ = µmethod
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9:
								goto Label9
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 696: """
							πF.SetLineno(696)
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µnext_state == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 712: if next_state is None:
							πF.SetLineno(712)
						Label1:
							// line 713: next_state = self.__class__.__name__
							πF.SetLineno(713)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
								continue
							}
							µnext_state = πTemp002
							goto Label2
						Label2:
							// line 714: try:
							πF.SetLineno(714)
							πF.PushCheckpoint(4)
							// line 715: pattern = self.patterns[name]
							πF.SetLineno(715)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpatterns, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µpattern = πTemp002
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp005[0] = µpattern
							πTemp005[1] = ßmatch.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 716: if not hasattr(pattern, 'match'):
							πF.SetLineno(716)
						Label5:
							// line 717: pattern = self.patterns[name] = re.compile(pattern)
							πF.SetLineno(717)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							πTemp005[0] = µpattern
							if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpattern = πTemp001
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpatterns, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp006 = µname
							if πE = πg.SetItem(πF, πTemp004, πTemp006, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 718: except KeyError:
							πF.SetLineno(718)
						Label7:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, µname).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s.patterns[%r]").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTransitionPatternNotFound); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 719: raise TransitionPatternNotFound(
							πF.SetLineno(719)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 721: try:
							πF.SetLineno(721)
							πF.PushCheckpoint(9)
							// line 722: method = getattr(self, name)
							πF.SetLineno(722)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmethod = πTemp002
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 723: except AttributeError:
							πF.SetLineno(723)
						Label10:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, µname).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s.%s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTransitionMethodNotFound); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 724: raise TransitionMethodNotFound(
							πF.SetLineno(724)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label8
						Label8:
							// line 726: return (pattern, method, next_state)
							πF.SetLineno(726)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µpattern, µmethod, µnext_state).ToObject()
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
					if πE = πClass.SetItem(πF, ßmake_transition.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 696: """
					πF.SetLineno(696)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n        Make & return a transition tuple based on `name`.\n\n        This is a convenience function to simplify transition creation.\n\n        Parameters:\n\n        - `name`: a string, the name of the transition pattern & method. This\n          `State` object must have a method called '`name`', and a dictionary\n          `self.patterns` containing a key '`name`'.\n        - `next_state`: a string, the name of the next `State` object for this\n          transition. A value of ``None`` (or absent) implies no state change\n          (i.e., continue with the same state).\n\n        Exceptions: `TransitionPatternNotFound`, `TransitionMethodNotFound`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßmake_transition); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 728: def make_transitions(self, name_list):
					πF.SetLineno(728)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name_list", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("make_transitions", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µname_list *πg.Object = πArgs[1]
						_ = µname_list
						var µstringtype *πg.Object = πg.UnboundLocal
						_ = µstringtype
						var µnames *πg.Object = πg.UnboundLocal
						_ = µnames
						var µtransitions *πg.Object = πg.UnboundLocal
						_ = µtransitions
						var µnamestate *πg.Object = πg.UnboundLocal
						_ = µnamestate
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 729: """
							πF.SetLineno(729)
							// line 736: stringtype = type('')
							πF.SetLineno(736)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstringtype = πTemp003
							// line 737: names = []
							πF.SetLineno(737)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µnames = πTemp002
							// line 738: transitions = {}
							πF.SetLineno(738)
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							µtransitions = πTemp002
							if πE = πg.CheckLocal(πF, µname_list, "name_list"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µname_list); πE != nil {
								continue
							}
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µnamestate = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							πTemp001[0] = µnamestate
							if πE = πg.CheckLocal(πF, µstringtype, "stringtype"); πE != nil {
								continue
							}
							πTemp001[1] = µstringtype
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 740: if isinstance(namestate, stringtype):
							πF.SetLineno(740)
						Label4:
							// line 741: transitions[namestate] = self.make_transition(namestate)
							πF.SetLineno(741)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							πTemp001[0] = µnamestate
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmake_transition, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							πTemp008 = µnamestate
							if πE = πg.SetItem(πF, µtransitions, πTemp008, πTemp003); πE != nil {
								continue
							}
							// line 742: names.append(namestate)
							πF.SetLineno(742)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							πTemp001[0] = µnamestate
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnames, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label5:
							// line 744: transitions[namestate[0]] = self.make_transition(*namestate)
							πF.SetLineno(744)
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmake_transition, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Invoke(πF, πTemp003, nil, µnamestate, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp009 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µnamestate, πTemp009); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.SetItem(πF, µtransitions, πTemp008, πTemp003); πE != nil {
								continue
							}
							// line 745: names.append(namestate[0])
							πF.SetLineno(745)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnamestate, "namestate"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µnamestate, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µnames, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 746: return names, transitions
							πF.SetLineno(746)
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µnames, µtransitions).ToObject()
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
					if πE = πClass.SetItem(πF, ßmake_transitions.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 729: """
					πF.SetLineno(729)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n        Return a list of transition names and a transition mapping.\n\n        Parameter `name_list`: a list, where each entry is either a transition\n        name string, or a 1- or 2-tuple (transition name, optional next state\n        name).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßmake_transitions); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 748: def no_match(self, context, transitions):
					πF.SetLineno(748)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp002[2] = πg.Param{Name: "transitions", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("no_match", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcontext *πg.Object = πArgs[1]
						_ = µcontext
						var µtransitions *πg.Object = πArgs[2]
						_ = µtransitions
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 749: """
							πF.SetLineno(749)
							// line 760: return context, None, []
							πF.SetLineno(760)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp001 = πg.NewTuple3(µcontext, πTemp002, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßno_match.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 749: """
					πF.SetLineno(749)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n        Called when there is no match from `StateMachine.check_line()`.\n\n        Return the same values returned by transition methods:\n\n        - context: unchanged;\n        - next state name: ``None``;\n        - empty result list.\n\n        Override in subclasses to catch this event.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßno_match); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 762: def bof(self, context):
					πF.SetLineno(762)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("bof", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcontext *πg.Object = πArgs[1]
						_ = µcontext
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 763: """
							πF.SetLineno(763)
							// line 770: return context, []
							πF.SetLineno(770)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001 = πg.NewTuple2(µcontext, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßbof.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 763: """
					πF.SetLineno(763)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n        Handle beginning-of-file. Return unchanged `context`, empty result.\n\n        Override in subclasses.\n\n        Parameter `context`: application-defined storage.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßbof); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 772: def eof(self, context):
					πF.SetLineno(772)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("eof", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcontext *πg.Object = πArgs[1]
						_ = µcontext
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
							// line 773: """
							πF.SetLineno(773)
							// line 780: return []
							πF.SetLineno(780)
							πTemp001 = make([]*πg.Object, 0)
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
					if πE = πClass.SetItem(πF, ßeof.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 773: """
					πF.SetLineno(773)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n        Handle end-of-file. Return empty result.\n\n        Override in subclasses.\n\n        Parameter `context`: application-defined storage.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßeof); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
						continue
					}
					// line 782: def nop(self, match, context, next_state):
					πF.SetLineno(782)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "match", Def: nil}
					πTemp002[2] = πg.Param{Name: "context", Def: nil}
					πTemp002[3] = πg.Param{Name: "next_state", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("nop", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µcontext *πg.Object = πArgs[2]
						_ = µcontext
						var µnext_state *πg.Object = πArgs[3]
						_ = µnext_state
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 783: """
							πF.SetLineno(783)
							// line 789: return context, next_state, []
							πF.SetLineno(789)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001 = πg.NewTuple3(µcontext, µnext_state, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßnop.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 783: """
					πF.SetLineno(783)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n        A \"do nothing\" transition method.\n\n        Return unchanged `context` & `next_state`, empty result. Useful for\n        simple state changes (actionless transitions).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßnop); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("State").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßState.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 792: class StateMachineWS(StateMachine):
			πF.SetLineno(792)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachine); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StateMachineWS", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 794: """
					πF.SetLineno(794)
					// line 794: """
					πF.SetLineno(794)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    `StateMachine` subclass specialized for whitespace recognition.\n\n    There are three methods provided for extracting indented text blocks:\n\n    - `get_indented()`: use when the indent is unknown.\n    - `get_known_indented()`: use when the indent is known for all lines.\n    - `get_first_known_indented()`: use when only the first line's indent is\n      known.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 805: def get_indented(self, until_blank=False, strip_indent=True):
					πF.SetLineno(805)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "until_blank", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "strip_indent", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("get_indented", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µuntil_blank *πg.Object = πArgs[1]
						_ = µuntil_blank
						var µstrip_indent *πg.Object = πArgs[2]
						_ = µstrip_indent
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µindent *πg.Object = πg.UnboundLocal
						_ = µindent
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 806: """
							πF.SetLineno(806)
							// line 821: offset = self.abs_line_offset()
							πF.SetLineno(821)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßabs_line_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoffset = πTemp002
							// line 822: indented, indent, blank_finish = self.input_lines.get_indented(
							πF.SetLineno(822)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µuntil_blank, "until_blank"); πE != nil {
								continue
							}
							πTemp003[1] = µuntil_blank
							if πE = πg.CheckLocal(πF, µstrip_indent, "strip_indent"); πE != nil {
								continue
							}
							πTemp003[2] = µstrip_indent
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_indented, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µindented = πTemp002
							µindent = πTemp004
							µblank_finish = πTemp005
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µindented); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 824: if indented:
							πF.SetLineno(824)
						Label1:
							// line 825: self.next_line(len(indented) - 1) # advance to last indented line
							πF.SetLineno(825)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp007[0] = µindented
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnext_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label2
						Label2:
							// line 826: while indented and not indented[0].strip():
							πF.SetLineno(826)
							πF.PushCheckpoint(4)
							πTemp006 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp001 = µindented
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label6
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µindented, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp010).ToObject()
							πTemp001 = πTemp002
						Label6:
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 827: indented.trim_start()
							πF.SetLineno(827)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µindented, ßtrim_start, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 828: offset += 1
							πF.SetLineno(828)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µoffset = πTemp001
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 829: return indented, indent, offset, blank_finish
							πF.SetLineno(829)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(µindented, µindent, µoffset, µblank_finish).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_indented.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 806: """
					πF.SetLineno(806)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Return a block of indented lines of text, and info.\n\n        Extract an indented block where the indent is unknown for all lines.\n\n        :Parameters:\n            - `until_blank`: Stop collecting at the first blank line if true.\n            - `strip_indent`: Strip common leading indent if true (default).\n\n        :Return:\n            - the indented block (a list of lines of text),\n            - its indent,\n            - its first line offset from BOF, and\n            - whether or not it finished with a blank line.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßget_indented); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 831: def get_known_indented(self, indent, until_blank=False, strip_indent=True):
					πF.SetLineno(831)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "until_blank", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "strip_indent", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("get_known_indented", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindent *πg.Object = πArgs[1]
						_ = µindent
						var µuntil_blank *πg.Object = πArgs[2]
						_ = µuntil_blank
						var µstrip_indent *πg.Object = πArgs[3]
						_ = µstrip_indent
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 832: """
							πF.SetLineno(832)
							// line 851: offset = self.abs_line_offset()
							πF.SetLineno(851)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßabs_line_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoffset = πTemp002
							// line 852: indented, indent, blank_finish = self.input_lines.get_indented(
							πF.SetLineno(852)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µuntil_blank, "until_blank"); πE != nil {
								continue
							}
							πTemp003[1] = µuntil_blank
							if πE = πg.CheckLocal(πF, µstrip_indent, "strip_indent"); πE != nil {
								continue
							}
							πTemp003[2] = µstrip_indent
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"block_indent", µindent},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_indented, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
								continue
							}
							µindented = πTemp002
							µindent = πTemp005
							µblank_finish = πTemp006
							// line 855: self.next_line(len(indented) - 1) # advance to last indented line
							πF.SetLineno(855)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp007[0] = µindented
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnext_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 856: while indented and not indented[0].strip():
							πF.SetLineno(856)
							πF.PushCheckpoint(2)
							πTemp008 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp001 = µindented
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label4
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µindented, πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp011).ToObject()
							πTemp001 = πTemp002
						Label4:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 857: indented.trim_start()
							πF.SetLineno(857)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µindented, ßtrim_start, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 858: offset += 1
							πF.SetLineno(858)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µoffset = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 859: return indented, offset, blank_finish
							πF.SetLineno(859)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µindented, µoffset, µblank_finish).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_known_indented.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 832: """
					πF.SetLineno(832)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Return an indented block and info.\n\n        Extract an indented block where the indent is known for all lines.\n        Starting with the current line, extract the entire text block with at\n        least `indent` indentation (which must be whitespace, except for the\n        first line).\n\n        :Parameters:\n            - `indent`: The number of indent columns/characters.\n            - `until_blank`: Stop collecting at the first blank line if true.\n            - `strip_indent`: Strip `indent` characters of indentation if true\n              (default).\n\n        :Return:\n            - the indented block,\n            - its first line offset from BOF, and\n            - whether or not it finished with a blank line.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßget_known_indented); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 861: def get_first_known_indented(self, indent, until_blank=False,
					πF.SetLineno(861)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "until_blank", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "strip_indent", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "strip_top", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("get_first_known_indented", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µindent *πg.Object = πArgs[1]
						_ = µindent
						var µuntil_blank *πg.Object = πArgs[2]
						_ = µuntil_blank
						var µstrip_indent *πg.Object = πArgs[3]
						_ = µstrip_indent
						var µstrip_top *πg.Object = πArgs[4]
						_ = µstrip_top
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 863: """
							πF.SetLineno(863)
							// line 883: offset = self.abs_line_offset()
							πF.SetLineno(883)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßabs_line_offset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µoffset = πTemp002
							// line 884: indented, indent, blank_finish = self.input_lines.get_indented(
							πF.SetLineno(884)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_offset, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µuntil_blank, "until_blank"); πE != nil {
								continue
							}
							πTemp003[1] = µuntil_blank
							if πE = πg.CheckLocal(πF, µstrip_indent, "strip_indent"); πE != nil {
								continue
							}
							πTemp003[2] = µstrip_indent
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"first_indent", µindent},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput_lines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_indented, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
								continue
							}
							µindented = πTemp002
							µindent = πTemp005
							µblank_finish = πTemp006
							// line 887: self.next_line(len(indented) - 1) # advance to last indented line
							πF.SetLineno(887)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp007[0] = µindented
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnext_line, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µstrip_top, "strip_top"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µstrip_top); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 888: if strip_top:
							πF.SetLineno(888)
						Label1:
							// line 889: while indented and not indented[0].strip():
							πF.SetLineno(889)
							πF.PushCheckpoint(4)
							πTemp008 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp001 = µindented
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label6
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µindented, πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp011).ToObject()
							πTemp001 = πTemp002
						Label6:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 890: indented.trim_start()
							πF.SetLineno(890)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µindented, ßtrim_start, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 891: offset += 1
							πF.SetLineno(891)
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µoffset = πTemp001
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							goto Label2
						Label2:
							// line 892: return indented, indent, offset, blank_finish
							πF.SetLineno(892)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple4(µindented, µindent, µoffset, µblank_finish).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_first_known_indented.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 863: """
					πF.SetLineno(863)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Return an indented block and info.\n\n        Extract an indented block where the indent is known for the first line\n        and unknown for all other lines.\n\n        :Parameters:\n            - `indent`: The first line's indent (# of columns/characters).\n            - `until_blank`: Stop collecting at the first blank line if true\n              (1).\n            - `strip_indent`: Strip `indent` characters of indentation if true\n              (1, default).\n            - `strip_top`: Strip blank lines from the beginning of the block.\n\n        :Return:\n            - the indented block,\n            - its indent,\n            - its first line offset from BOF, and\n            - whether or not it finished with a blank line.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßget_first_known_indented); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StateMachineWS").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStateMachineWS.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 895: class StateWS(State):
			πF.SetLineno(895)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßState); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StateWS", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Dict
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 897: """
					πF.SetLineno(897)
					// line 897: """
					πF.SetLineno(897)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    State superclass specialized for whitespace (blank lines & indents).\n\n    Use this class with `StateMachineWS`.  The transitions 'blank' (for blank\n    lines) and 'indent' (for indented text blocks) are added automatically,\n    before any other transitions.  The transition method `blank()` handles\n    blank lines and `indent()` handles nested indented blocks.  Indented\n    blocks trigger a new state machine to be created by `indent()` and run.\n    The class of the state machine to be created is in `indent_sm`, and the\n    constructor keyword arguments are in the dictionary `indent_sm_kwargs`.\n\n    The methods `known_indent()` and `firstknown_indent()` are provided for\n    indented blocks where the indent (all lines' and first line's only,\n    respectively) is known to the transition method, along with the attributes\n    `known_indent_sm` and `known_indent_sm_kwargs`.  Neither transition method\n    is triggered automatically.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 915: indent_sm = None
					πF.SetLineno(915)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßindent_sm.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 916: """
					πF.SetLineno(916)
					// line 923: indent_sm_kwargs = None
					πF.SetLineno(923)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßindent_sm_kwargs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 924: """
					πF.SetLineno(924)
					// line 931: known_indent_sm = None
					πF.SetLineno(931)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßknown_indent_sm.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 932: """
					πF.SetLineno(932)
					// line 939: known_indent_sm_kwargs = None
					πF.SetLineno(939)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßknown_indent_sm_kwargs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 940: """
					πF.SetLineno(940)
					// line 947: ws_patterns = {'blank': re.compile(' *$'),
					πF.SetLineno(947)
					πTemp002 = πg.NewDict()
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(" *$").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πTemp002.SetItem(πF, ßblank.ToObject(), πTemp001); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(" +").ToObject()
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πTemp002.SetItem(πF, ßindent.ToObject(), πTemp001); πE != nil {
						continue
					}
					πTemp001 = πTemp002.ToObject()
					if πE = πClass.SetItem(πF, ßws_patterns.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 949: """Patterns for default whitespace transitions.  May be overridden in
					πF.SetLineno(949)
					// line 952: ws_initial_transitions = ('blank', 'indent')
					πF.SetLineno(952)
					πTemp001 = πg.NewTuple2(ßblank.ToObject(), ßindent.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßws_initial_transitions.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 953: """Default initial whitespace transitions, added before those listed in
					πF.SetLineno(953)
					// line 956: def __init__(self, state_machine, debug=False):
					πF.SetLineno(956)
					πTemp005 = make([]πg.Param, 3)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "state_machine", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp005[2] = πg.Param{Name: "debug", Def: πTemp004}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstate_machine *πg.Object = πArgs[1]
						_ = µstate_machine
						var µdebug *πg.Object = πArgs[2]
						_ = µdebug
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 957: """
							πF.SetLineno(957)
							// line 962: State.__init__(self, state_machine, debug)
							πF.SetLineno(962)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µstate_machine, "state_machine"); πE != nil {
								continue
							}
							πTemp001[1] = µstate_machine
							if πE = πg.CheckLocal(πF, µdebug, "debug"); πE != nil {
								continue
							}
							πTemp001[2] = µdebug
							if πTemp002, πE = πg.ResolveGlobal(πF, ßState); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßindent_sm, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 963: if self.indent_sm is None:
							πF.SetLineno(963)
						Label1:
							// line 964: self.indent_sm = self.nested_sm
							πF.SetLineno(964)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnested_sm, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindent_sm, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßindent_sm_kwargs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 965: if self.indent_sm_kwargs is None:
							πF.SetLineno(965)
						Label3:
							// line 966: self.indent_sm_kwargs = self.nested_sm_kwargs
							πF.SetLineno(966)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnested_sm_kwargs, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindent_sm_kwargs, πTemp003); πE != nil {
								continue
							}
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßknown_indent_sm, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 967: if self.known_indent_sm is None:
							πF.SetLineno(967)
						Label5:
							// line 968: self.known_indent_sm = self.indent_sm
							πF.SetLineno(968)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent_sm, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßknown_indent_sm, πTemp003); πE != nil {
								continue
							}
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßknown_indent_sm_kwargs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 969: if self.known_indent_sm_kwargs is None:
							πF.SetLineno(969)
						Label7:
							// line 970: self.known_indent_sm_kwargs = self.indent_sm_kwargs
							πF.SetLineno(970)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent_sm_kwargs, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßknown_indent_sm_kwargs, πTemp003); πE != nil {
								continue
							}
							goto Label8
						Label8:
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
					// line 957: """
					πF.SetLineno(957)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Initialize a `StateSM` object; extends `State.__init__()`.\n\n        Check for indent state machine attributes, set defaults if not set.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 972: def add_initial_transitions(self):
					πF.SetLineno(972)
					πTemp005 = make([]πg.Param, 1)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("add_initial_transitions", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µnames *πg.Object = πg.UnboundLocal
						_ = µnames
						var µtransitions *πg.Object = πg.UnboundLocal
						_ = µtransitions
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Dict
						_ = πTemp006
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
							// line 973: """
							πF.SetLineno(973)
							// line 978: State.add_initial_transitions(self)
							πF.SetLineno(978)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßState); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadd_initial_transitions, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpatterns, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 979: if self.patterns is None:
							πF.SetLineno(979)
						Label1:
							// line 980: self.patterns = {}
							πF.SetLineno(980)
							πTemp006 = πg.NewDict()
							πTemp002 = πTemp006.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpatterns, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 981: self.patterns.update(self.ws_patterns)
							πF.SetLineno(981)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßws_patterns, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpatterns, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 982: names, transitions = self.make_transitions(
							πF.SetLineno(982)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßws_initial_transitions, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmake_transitions, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µnames = πTemp002
							µtransitions = πTemp004
							// line 984: self.add_transitions(names, transitions)
							πF.SetLineno(984)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							πTemp001[0] = µnames
							if πE = πg.CheckLocal(πF, µtransitions, "transitions"); πE != nil {
								continue
							}
							πTemp001[1] = µtransitions
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_transitions, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_initial_transitions.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 973: """
					πF.SetLineno(973)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Add whitespace-specific transitions before those defined in subclass.\n\n        Extends `State.add_initial_transitions()`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßadd_initial_transitions); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 986: def blank(self, match, context, next_state):
					πF.SetLineno(986)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "match", Def: nil}
					πTemp005[2] = πg.Param{Name: "context", Def: nil}
					πTemp005[3] = πg.Param{Name: "next_state", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("blank", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µcontext *πg.Object = πArgs[2]
						_ = µcontext
						var µnext_state *πg.Object = πArgs[3]
						_ = µnext_state
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
							// line 987: """Handle blank lines. Does nothing. Override in subclasses."""
							πF.SetLineno(987)
							// line 988: return self.nop(match, context, next_state)
							πF.SetLineno(988)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							πTemp001[0] = µmatch
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp001[1] = µcontext
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							πTemp001[2] = µnext_state
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnop, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßblank.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 987: """Handle blank lines. Does nothing. Override in subclasses."""
					πF.SetLineno(987)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Handle blank lines. Does nothing. Override in subclasses.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßblank); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 990: def indent(self, match, context, next_state):
					πF.SetLineno(990)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "match", Def: nil}
					πTemp005[2] = πg.Param{Name: "context", Def: nil}
					πTemp005[3] = πg.Param{Name: "next_state", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("indent", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µcontext *πg.Object = πArgs[2]
						_ = µcontext
						var µnext_state *πg.Object = πArgs[3]
						_ = µnext_state
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µindent *πg.Object = πg.UnboundLocal
						_ = µindent
						var µline_offset *πg.Object = πg.UnboundLocal
						_ = µline_offset
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var µsm *πg.Object = πg.UnboundLocal
						_ = µsm
						var µresults *πg.Object = πg.UnboundLocal
						_ = µresults
						var πTemp001 *πg.Object
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
							// line 991: """
							πF.SetLineno(991)
							// line 997: indented, indent, line_offset, blank_finish = \
							πF.SetLineno(997)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_indented, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µindented = πTemp002
							µindent = πTemp003
							µline_offset = πTemp004
							µblank_finish = πTemp005
							// line 999: sm = self.indent_sm(debug=self.debug, **self.indent_sm_kwargs)
							πF.SetLineno(999)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"debug", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßindent_sm_kwargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent_sm, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, nil, πTemp006, πTemp001); πE != nil {
								continue
							}
							µsm = πTemp003
							// line 1000: results = sm.run(indented, input_offset=line_offset)
							πF.SetLineno(1000)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp007[0] = µindented
							if πE = πg.CheckLocal(πF, µline_offset, "line_offset"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"input_offset", µline_offset},
							}
							if πE = πg.CheckLocal(πF, µsm, "sm"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsm, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µresults = πTemp002
							// line 1001: return context, next_state, results
							πF.SetLineno(1001)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µcontext, µnext_state, µresults).ToObject()
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
					if πE = πClass.SetItem(πF, ßindent.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 991: """
					πF.SetLineno(991)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Handle an indented text block. Extend or override in subclasses.\n\n        Recursively run the registered state machine for indented blocks\n        (`self.indent_sm`).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßindent); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 1003: def known_indent(self, match, context, next_state):
					πF.SetLineno(1003)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "match", Def: nil}
					πTemp005[2] = πg.Param{Name: "context", Def: nil}
					πTemp005[3] = πg.Param{Name: "next_state", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("known_indent", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µcontext *πg.Object = πArgs[2]
						_ = µcontext
						var µnext_state *πg.Object = πArgs[3]
						_ = µnext_state
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µline_offset *πg.Object = πg.UnboundLocal
						_ = µline_offset
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var µsm *πg.Object = πg.UnboundLocal
						_ = µsm
						var µresults *πg.Object = πg.UnboundLocal
						_ = µresults
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
							// line 1004: """
							πF.SetLineno(1004)
							// line 1011: indented, line_offset, blank_finish = \
							πF.SetLineno(1011)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_known_indented, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µindented = πTemp003
							µline_offset = πTemp004
							µblank_finish = πTemp005
							// line 1013: sm = self.known_indent_sm(debug=self.debug,
							πF.SetLineno(1013)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"debug", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßknown_indent_sm_kwargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßknown_indent_sm, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, nil, nil, πTemp006, πTemp002); πE != nil {
								continue
							}
							µsm = πTemp004
							// line 1015: results = sm.run(indented, input_offset=line_offset)
							πF.SetLineno(1015)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp001[0] = µindented
							if πE = πg.CheckLocal(πF, µline_offset, "line_offset"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"input_offset", µline_offset},
							}
							if πE = πg.CheckLocal(πF, µsm, "sm"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsm, ßrun, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresults = πTemp003
							// line 1016: return context, next_state, results
							πF.SetLineno(1016)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µcontext, µnext_state, µresults).ToObject()
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
					if πE = πClass.SetItem(πF, ßknown_indent.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1004: """
					πF.SetLineno(1004)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Handle a known-indent text block. Extend or override in subclasses.\n\n        Recursively run the registered state machine for known-indent indented\n        blocks (`self.known_indent_sm`). The indent is the length of the\n        match, ``match.end()``.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßknown_indent); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 1018: def first_known_indent(self, match, context, next_state):
					πF.SetLineno(1018)
					πTemp005 = make([]πg.Param, 4)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "match", Def: nil}
					πTemp005[2] = πg.Param{Name: "context", Def: nil}
					πTemp005[3] = πg.Param{Name: "next_state", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("first_known_indent", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmatch *πg.Object = πArgs[1]
						_ = µmatch
						var µcontext *πg.Object = πArgs[2]
						_ = µcontext
						var µnext_state *πg.Object = πArgs[3]
						_ = µnext_state
						var µindented *πg.Object = πg.UnboundLocal
						_ = µindented
						var µline_offset *πg.Object = πg.UnboundLocal
						_ = µline_offset
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var µsm *πg.Object = πg.UnboundLocal
						_ = µsm
						var µresults *πg.Object = πg.UnboundLocal
						_ = µresults
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
							// line 1019: """
							πF.SetLineno(1019)
							// line 1028: indented, line_offset, blank_finish = \
							πF.SetLineno(1028)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate_machine, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget_first_known_indented, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µindented = πTemp003
							µline_offset = πTemp004
							µblank_finish = πTemp005
							// line 1030: sm = self.known_indent_sm(debug=self.debug,
							πF.SetLineno(1030)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdebug, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"debug", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßknown_indent_sm_kwargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßknown_indent_sm, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, nil, nil, πTemp006, πTemp002); πE != nil {
								continue
							}
							µsm = πTemp004
							// line 1032: results = sm.run(indented, input_offset=line_offset)
							πF.SetLineno(1032)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindented, "indented"); πE != nil {
								continue
							}
							πTemp001[0] = µindented
							if πE = πg.CheckLocal(πF, µline_offset, "line_offset"); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"input_offset", µline_offset},
							}
							if πE = πg.CheckLocal(πF, µsm, "sm"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsm, ßrun, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresults = πTemp003
							// line 1033: return context, next_state, results
							πF.SetLineno(1033)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnext_state, "next_state"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µcontext, µnext_state, µresults).ToObject()
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
					if πE = πClass.SetItem(πF, ßfirst_known_indent.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1019: """
					πF.SetLineno(1019)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n        Handle an indented text block (first line's indent known).\n\n        Extend or override in subclasses.\n\n        Recursively run the registered state machine for known-indent indented\n        blocks (`self.known_indent_sm`). The indent is the length of the\n        match, ``match.end()``.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßfirst_known_indent); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StateWS").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStateWS.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1036: class _SearchOverride(object):
			πF.SetLineno(1036)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_SearchOverride", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
					// line 1038: """
					πF.SetLineno(1038)
					// line 1038: """
					πF.SetLineno(1038)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Mix-in class to override `StateMachine` regular expression behavior.\n\n    Changes regular expression matching, from the default `re.match()`\n    (succeeds only if the pattern matches at the start of `self.line`) to\n    `re.search()` (succeeds if the pattern matches anywhere in `self.line`).\n    When subclassing a `StateMachine`, list this class **first** in the\n    inheritance list of the class definition.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1048: def match(self, pattern):
					πF.SetLineno(1048)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("match", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpattern *πg.Object = πArgs[1]
						_ = µpattern
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
							// line 1049: """
							πF.SetLineno(1049)
							// line 1056: return pattern.search(self.line)
							πF.SetLineno(1056)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpattern, ßsearch, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmatch.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1049: """
					πF.SetLineno(1049)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Return the result of a regular expression search.\n\n        Overrides `StateMachine.match()`.\n\n        Parameter `pattern`: `re` compiled regular expression.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßmatch); πE != nil {
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
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_SearchOverride").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SearchOverride.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1059: class SearchStateMachine(_SearchOverride, StateMachine):
			πF.SetLineno(1059)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp004, πE = πg.ResolveGlobal(πF, ß_SearchOverride); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachine); πE != nil {
				continue
			}
			πTemp002[1] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SearchStateMachine", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1060: """`StateMachine` which uses `re.search()` instead of `re.match()`."""
					πF.SetLineno(1060)
					// line 1060: """`StateMachine` which uses `re.search()` instead of `re.match()`."""
					πF.SetLineno(1060)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("`StateMachine` which uses `re.search()` instead of `re.match()`.").ToObject()); πE != nil {
						continue
					}
					// line 1061: pass
					πF.SetLineno(1061)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SearchStateMachine").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSearchStateMachine.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1064: class SearchStateMachineWS(_SearchOverride, StateMachineWS):
			πF.SetLineno(1064)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp004, πE = πg.ResolveGlobal(πF, ß_SearchOverride); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineWS); πE != nil {
				continue
			}
			πTemp002[1] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SearchStateMachineWS", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1065: """`StateMachineWS` which uses `re.search()` instead of `re.match()`."""
					πF.SetLineno(1065)
					// line 1065: """`StateMachineWS` which uses `re.search()` instead of `re.match()`."""
					πF.SetLineno(1065)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("`StateMachineWS` which uses `re.search()` instead of `re.match()`.").ToObject()); πE != nil {
						continue
					}
					// line 1066: pass
					πF.SetLineno(1066)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SearchStateMachineWS").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSearchStateMachineWS.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1069: class ViewList(object):
			πF.SetLineno(1069)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ViewList", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				var πTemp019 *πg.Object
				_ = πTemp019
				var πTemp020 *πg.Object
				_ = πTemp020
				var πTemp021 *πg.Object
				_ = πTemp021
				var πTemp022 *πg.Object
				_ = πTemp022
				var πTemp023 *πg.Object
				_ = πTemp023
				var πTemp024 *πg.Object
				_ = πTemp024
				var πTemp025 *πg.Object
				_ = πTemp025
				var πTemp026 *πg.Object
				_ = πTemp026
				var πTemp027 *πg.Object
				_ = πTemp027
				var πTemp028 *πg.Object
				_ = πTemp028
				var πTemp029 *πg.Object
				_ = πTemp029
				var πTemp030 *πg.Object
				_ = πTemp030
				var πTemp031 *πg.Object
				_ = πTemp031
				var πTemp032 *πg.Object
				_ = πTemp032
				var πTemp033 *πg.Object
				_ = πTemp033
				var πTemp034 *πg.Object
				_ = πTemp034
				var πTemp035 *πg.Object
				_ = πTemp035
				var πTemp036 *πg.Object
				_ = πTemp036
				var πTemp037 *πg.Object
				_ = πTemp037
				var πTemp038 *πg.Object
				_ = πTemp038
				var πTemp039 *πg.Object
				_ = πTemp039
				var πTemp040 *πg.Object
				_ = πTemp040
				var πTemp041 *πg.Object
				_ = πTemp041
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1071: """
					πF.SetLineno(1071)
					// line 1071: """
					πF.SetLineno(1071)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    List with extended functionality: slices of ViewList objects are child\n    lists, linked to their parents. Changes made to a child list also affect\n    the parent list.  A child list is effectively a \"view\" (in the SQL sense)\n    of the parent list.  Changes to parent lists, however, do *not* affect\n    active child lists.  If a parent list is changed, any active child lists\n    should be recreated.\n\n    The start and end of the slice can be trimmed using the `trim_start()` and\n    `trim_end()` methods, without affecting the parent list.  The link between\n    child and parent lists can be broken by calling `disconnect()` on the\n    child list.\n\n    Also, ViewList objects keep track of the source & offset of each item.\n    This information is accessible via the `source()`, `offset()`, and\n    `info()` methods.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1089: def __init__(self, initlist=None, source=None, items=None,
					πF.SetLineno(1089)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "initlist", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "source", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "items", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "parent", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "parent_offset", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µinitlist *πg.Object = πArgs[1]
						_ = µinitlist
						var µsource *πg.Object = πArgs[2]
						_ = µsource
						var µitems *πg.Object = πArgs[3]
						_ = µitems
						var µparent *πg.Object = πArgs[4]
						_ = µparent
						var µparent_offset *πg.Object = πArgs[5]
						_ = µparent_offset
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							// line 1091: self.data = []
							πF.SetLineno(1091)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp003); πE != nil {
								continue
							}
							// line 1092: """The actual list of data, flattened from various sources."""
							πF.SetLineno(1092)
							// line 1094: self.items = []
							πF.SetLineno(1094)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp003); πE != nil {
								continue
							}
							// line 1095: """A list of (source, offset) pairs, same length as `self.data`: the
							πF.SetLineno(1095)
							// line 1099: self.parent = parent
							πF.SetLineno(1099)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µparent); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp002); πE != nil {
								continue
							}
							// line 1100: """The parent list."""
							πF.SetLineno(1100)
							// line 1102: self.parent_offset = parent_offset
							πF.SetLineno(1102)
							if πE = πg.CheckLocal(πF, µparent_offset, "parent_offset"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µparent_offset); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent_offset, πTemp002); πE != nil {
								continue
							}
							// line 1103: """Offset of this list from the beginning of the parent list."""
							πF.SetLineno(1103)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinitlist, "initlist"); πE != nil {
								continue
							}
							πTemp001[0] = µinitlist
							if πTemp002, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µinitlist, "initlist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µinitlist != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 1105: if isinstance(initlist, ViewList):
							πF.SetLineno(1105)
						Label1:
							// line 1106: self.data = initlist.data[:]
							πF.SetLineno(1106)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µinitlist, "initlist"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µinitlist, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp002); πE != nil {
								continue
							}
							// line 1107: self.items = initlist.items[:]
							πF.SetLineno(1107)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µinitlist, "initlist"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µinitlist, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp002); πE != nil {
								continue
							}
							goto Label3
							// line 1108: elif initlist is not None:
							πF.SetLineno(1108)
						Label2:
							// line 1109: self.data = list(initlist)
							πF.SetLineno(1109)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinitlist, "initlist"); πE != nil {
								continue
							}
							πTemp001[0] = µinitlist
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µitems); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 1110: if items:
							πF.SetLineno(1110)
						Label4:
							// line 1111: self.items = items
							πF.SetLineno(1111)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitems); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label5:
							// line 1113: self.items = [(source, i) for i in range(len(initlist))]
							πF.SetLineno(1113)
							πTemp006 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
								var πTemp006 bool
								_ = πTemp006
								var πTemp007 bool
								_ = πTemp007
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
										πTemp002 = πF.MakeArgs(1)
										πTemp003 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µinitlist, "initlist"); πE != nil {
											continue
										}
										πTemp003[0] = µinitlist
										if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp003)
										πTemp002[0] = πTemp005
										if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp006 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp006 {
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
											πTemp007 = !isStop
										} else {
											πTemp007 = true
											µi = πTemp004
										}
										if πE != nil || !πTemp007 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1113: self.items = [(source, i) for i in range(len(initlist))]
										πF.SetLineno(1113)
										if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp004 = πg.NewTuple2(µsource, µi).ToObject()
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp005 = πSent
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp005); πE != nil {
								continue
							}
							goto Label6
						Label6:
							goto Label3
						Label3:
							// line 1114: assert len(self.data) == len(self.items), 'data mismatch'
							πF.SetLineno(1114)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("data mismatch").ToObject()); πE != nil {
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
					// line 1116: def __str__(self):
					πF.SetLineno(1116)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1117: return str(self.data)
							πF.SetLineno(1117)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1119: def __repr__(self):
					πF.SetLineno(1119)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 1120: return '%s(%s, items=%s)' % (self.__class__.__name__,
							πF.SetLineno(1120)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp004, πTemp003, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%s, items=%s)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1123: def __lt__(self, other): return self.data <  self.__cast(other)
					πF.SetLineno(1123)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 1123: def __lt__(self, other): return self.data <  self.__cast(other)
							πF.SetLineno(1123)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.LT(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1124: def __le__(self, other): return self.data <= self.__cast(other)
					πF.SetLineno(1124)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 1124: def __le__(self, other): return self.data <= self.__cast(other)
							πF.SetLineno(1124)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.LE(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1125: def __eq__(self, other): return self.data == self.__cast(other)
					πF.SetLineno(1125)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 1125: def __eq__(self, other): return self.data == self.__cast(other)
							πF.SetLineno(1125)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1126: def __ne__(self, other): return self.data != self.__cast(other)
					πF.SetLineno(1126)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 1126: def __ne__(self, other): return self.data != self.__cast(other)
							πF.SetLineno(1126)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.NE(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1127: def __gt__(self, other): return self.data >  self.__cast(other)
					πF.SetLineno(1127)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 1127: def __gt__(self, other): return self.data >  self.__cast(other)
							πF.SetLineno(1127)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1128: def __ge__(self, other): return self.data >= self.__cast(other)
					πF.SetLineno(1128)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							// line 1128: def __ge__(self, other): return self.data >= self.__cast(other)
							πF.SetLineno(1128)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GE(πF, πTemp002, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1130: def __cmp__(self, other):
					πF.SetLineno(1130)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__cmp__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
						var µmine *πg.Object = πg.UnboundLocal
						_ = µmine
						var µyours *πg.Object = πg.UnboundLocal
						_ = µyours
						var πTemp001 *πg.Object
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
							// line 1132: mine = self.data
							πF.SetLineno(1132)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							µmine = πTemp001
							// line 1133: yours = self.__cast(other)
							πF.SetLineno(1133)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cast, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µyours = πTemp003
							// line 1134: return (mine > yours) - (yours < mine)
							πF.SetLineno(1134)
							if πE = πg.CheckLocal(πF, µmine, "mine"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µyours, "yours"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µmine, µyours); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µyours, "yours"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmine, "mine"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µyours, µmine); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp003, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__cmp__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1136: def __cast(self, other):
					πF.SetLineno(1136)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__cast", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
						var πTemp001 []*πg.Object
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 1137: if isinstance(other, ViewList):
							πF.SetLineno(1137)
						Label1:
							// line 1138: return other.data
							πF.SetLineno(1138)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ßdata, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1140: return other
							πF.SetLineno(1140)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πR = µother
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
					if πE = πClass.SetItem(πF, ß__cast.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1142: def __contains__(self, item): return item in self.data
					πF.SetLineno(1142)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 1142: def __contains__(self, item): return item in self.data
							πF.SetLineno(1142)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp002, µitem); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1143: def __len__(self): return len(self.data)
					πF.SetLineno(1143)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1143: def __len__(self): return len(self.data)
							πF.SetLineno(1143)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1149: def __getitem__(self, i):
					πF.SetLineno(1149)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 1150: if isinstance(i, slice):
							πF.SetLineno(1150)
						Label1:
							// line 1151: assert i.step in (None, 1),  'cannot handle slice with stride'
							πF.SetLineno(1151)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µi, ßstep, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πg.NewInt(1).ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("cannot handle slice with stride").ToObject()); πE != nil {
								continue
							}
							// line 1152: return self.__class__(self.data[i.start:i.stop],
							πF.SetLineno(1152)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							πTemp002 = πg.NewInt(0).ToObject()
						Label4:
							πTemp007 = πg.KWArgs{
								{"items", πTemp003},
								{"parent", µself},
								{"parent_offset", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 1156: return self.data[i]
							πF.SetLineno(1156)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1158: def __setitem__(self, i, item):
					πF.SetLineno(1158)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp002[2] = πg.Param{Name: "item", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var µitem *πg.Object = πArgs[2]
						_ = µitem
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 1159: if isinstance(i, slice):
							πF.SetLineno(1159)
						Label1:
							// line 1160: assert i.step in (None, 1), 'cannot handle slice with stride'
							πF.SetLineno(1160)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µi, ßstep, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πg.NewInt(1).ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("cannot handle slice with stride").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πTemp003, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 1161: if not isinstance(item, ViewList):
							πF.SetLineno(1161)
						Label4:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("assigning non-ViewList to ViewList slice").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1162: raise TypeError('assigning non-ViewList to ViewList slice')
							πF.SetLineno(1162)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 1163: self.data[i.start:i.stop] = item.data
							πF.SetLineno(1163)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitem, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πTemp007, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp003); πE != nil {
								continue
							}
							// line 1164: self.items[i.start:i.stop] = item.items
							πF.SetLineno(1164)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitem, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πTemp007, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp003); πE != nil {
								continue
							}
							// line 1165: assert len(self.data) == len(self.items), 'data mismatch'
							πF.SetLineno(1165)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp002, πg.NewStr("data mismatch").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 1166: if self.parent:
							πF.SetLineno(1166)
						Label6:
							// line 1167: self.parent[(i.start or 0) + self.parent_offset
							πF.SetLineno(1167)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πTemp007 = πg.NewInt(0).ToObject()
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							πTemp008 = πTemp009
							if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp008 = πTemp010
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πTemp007, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							goto Label3
						Label2:
							// line 1170: self.data[i] = item
							πF.SetLineno(1170)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 1171: if self.parent:
							πF.SetLineno(1171)
						Label10:
							// line 1172: self.parent[i + self.parent_offset] = item
							πF.SetLineno(1172)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µitem); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µi, πTemp007); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							goto Label11
						Label11:
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1174: def __delitem__(self, i):
					πF.SetLineno(1174)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
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
							// line 1175: try:
							πF.SetLineno(1175)
							πF.PushCheckpoint(2)
							// line 1176: del self.data[i]
							πF.SetLineno(1176)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 1177: del self.items[i]
							πF.SetLineno(1177)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1178: if self.parent:
							πF.SetLineno(1178)
						Label3:
							// line 1179: del self.parent[i + self.parent_offset]
							πF.SetLineno(1179)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µi, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							goto Label4
						Label4:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 1180: except TypeError:
							πF.SetLineno(1180)
						Label5:
							// line 1181: assert i.step is None, 'cannot handle slice with stride'
							πF.SetLineno(1181)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µi, ßstep, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("cannot handle slice with stride").ToObject()); πE != nil {
								continue
							}
							// line 1182: del self.data[i.start:i.stop]
							πF.SetLineno(1182)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 1183: del self.items[i.start:i.stop]
							πF.SetLineno(1183)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 1184: if self.parent:
							πF.SetLineno(1184)
						Label6:
							// line 1185: del self.parent[(i.start or 0) + self.parent_offset
							πF.SetLineno(1185)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µi, ßstart, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp008
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							πTemp005 = πg.NewInt(0).ToObject()
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µi, ßstop, nil); πE != nil {
								continue
							}
							πTemp008 = πTemp009
							if πTemp003, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp010[0] = µself
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp008 = πTemp011
						Label9:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 1188: def __add__(self, other):
					πF.SetLineno(1188)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 1189: if isinstance(other, ViewList):
							πF.SetLineno(1189)
						Label1:
							// line 1190: return self.__class__(self.data + other.data,
							πF.SetLineno(1190)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"items", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label3
						Label2:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("adding non-ViewList to a ViewList").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1193: raise TypeError('adding non-ViewList to a ViewList')
							πF.SetLineno(1193)
							πE = πF.Raise(πTemp003, nil, nil)
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 1195: def __radd__(self, other):
					πF.SetLineno(1195)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("__radd__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 1196: if isinstance(other, ViewList):
							πF.SetLineno(1196)
						Label1:
							// line 1197: return self.__class__(other.data + self.data,
							πF.SetLineno(1197)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"items", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label3
						Label2:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("adding ViewList to a non-ViewList").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1200: raise TypeError('adding ViewList to a non-ViewList')
							πF.SetLineno(1200)
							πE = πF.Raise(πTemp003, nil, nil)
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
					if πE = πClass.SetItem(πF, ß__radd__.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1202: def __iadd__(self, other):
					πF.SetLineno(1202)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("__iadd__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 1203: if isinstance(other, ViewList):
							πF.SetLineno(1203)
						Label1:
							// line 1204: self.data += other.data
							πF.SetLineno(1204)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßdata, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp005); πE != nil {
								continue
							}
							goto Label3
						Label2:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("argument to += must be a ViewList").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1206: raise TypeError('argument to += must be a ViewList')
							πF.SetLineno(1206)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 1207: return self
							πF.SetLineno(1207)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iadd__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1209: def __mul__(self, n):
					πF.SetLineno(1209)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("__mul__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 1210: return self.__class__(self.data * n, items=(self.items * n))
							πF.SetLineno(1210)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp003, µn); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp003, µn); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"items", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__mul__.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1212: __rmul__ = __mul__
					πF.SetLineno(1212)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ß__mul__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__rmul__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1214: def __imul__(self, n):
					πF.SetLineno(1214)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__imul__", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
						var πTemp001 *πg.Object
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
							// line 1215: self.data *= n
							πF.SetLineno(1215)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IMul(πF, πTemp001, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp002); πE != nil {
								continue
							}
							// line 1216: self.items *= n
							πF.SetLineno(1216)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IMul(πF, πTemp001, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp002); πE != nil {
								continue
							}
							// line 1217: return self
							πF.SetLineno(1217)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__imul__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1219: def extend(self, other):
					πF.SetLineno(1219)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("extend", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother *πg.Object = πArgs[1]
						_ = µother
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
						var πTemp006 []*πg.Object
						_ = πTemp006
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
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 1220: if not isinstance(other, ViewList):
							πF.SetLineno(1220)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("extending a ViewList with a non-ViewList").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1221: raise TypeError('extending a ViewList with a non-ViewList')
							πF.SetLineno(1221)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 1222: if self.parent:
							πF.SetLineno(1222)
						Label3:
							// line 1223: self.parent.insert(len(self.data) + self.parent_offset, other)
							πF.SetLineno(1223)
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[1] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label4
						Label4:
							// line 1224: self.data.extend(other.data)
							πF.SetLineno(1224)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µother, ßdata, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1225: self.items.extend(other.items)
							πF.SetLineno(1225)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßextend.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1227: def append(self, item, source=None, offset=0):
					πF.SetLineno(1227)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "source", Def: πTemp025}
					πTemp002[3] = πg.Param{Name: "offset", Def: πg.NewInt(0).ToObject()}
					πTemp024 = πg.NewFunction(πg.NewCode("append", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
						var µsource *πg.Object = πArgs[2]
						_ = µsource
						var µoffset *πg.Object = πArgs[3]
						_ = µoffset
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
						var πTemp006 *πg.Object
						_ = πTemp006
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
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µsource == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1228: if source is None:
							πF.SetLineno(1228)
						Label1:
							// line 1229: self.extend(item)
							πF.SetLineno(1229)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 1231: if self.parent:
							πF.SetLineno(1231)
						Label4:
							// line 1232: self.parent.insert(len(self.data) + self.parent_offset, item,
							πF.SetLineno(1232)
							πTemp004 = πF.MakeArgs(4)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[1] = µitem
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp004[2] = µsource
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp004[3] = µoffset
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label5:
							// line 1234: self.data.append(item)
							πF.SetLineno(1234)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1235: self.items.append((source, offset))
							πF.SetLineno(1235)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µsource, µoffset).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1237: def insert(self, i, item, source=None, offset=0):
					πF.SetLineno(1237)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp002[2] = πg.Param{Name: "item", Def: nil}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "source", Def: πTemp026}
					πTemp002[4] = πg.Param{Name: "offset", Def: πg.NewInt(0).ToObject()}
					πTemp025 = πg.NewFunction(πg.NewCode("insert", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var µitem *πg.Object = πArgs[2]
						_ = µitem
						var µsource *πg.Object = πArgs[3]
						_ = µsource
						var µoffset *πg.Object = πArgs[4]
						_ = µoffset
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µsource == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1238: if source is None:
							πF.SetLineno(1238)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[0] = µitem
							if πTemp002, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
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
							// line 1239: if not isinstance(item, ViewList):
							πF.SetLineno(1239)
						Label4:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("inserting non-ViewList with no source given").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1240: raise TypeError('inserting non-ViewList with no source given')
							πF.SetLineno(1240)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label5
						Label5:
							// line 1241: self.data[i:i] = item.data
							πF.SetLineno(1241)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µitem, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µi, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							// line 1242: self.items[i:i] = item.items
							πF.SetLineno(1242)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µitem, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µi, µi, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 1243: if self.parent:
							πF.SetLineno(1243)
						Label6:
							// line 1244: index = (len(self.data) + i) % len(self.data)
							πF.SetLineno(1244)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, µi); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							µindex = πTemp001
							// line 1245: self.parent.insert(index + self.parent_offset, item)
							πF.SetLineno(1245)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µindex, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[1] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label7
						Label7:
							goto Label3
						Label2:
							// line 1247: self.data.insert(i, item)
							πF.SetLineno(1247)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[1] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1248: self.items.insert(i, (source, offset))
							πF.SetLineno(1248)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µsource, µoffset).ToObject()
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 1249: if self.parent:
							πF.SetLineno(1249)
						Label8:
							// line 1250: index = (len(self.data) + i) % len(self.data)
							πF.SetLineno(1250)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, µi); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mod(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							µindex = πTemp001
							// line 1251: self.parent.insert(index + self.parent_offset, item,
							πF.SetLineno(1251)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µindex, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp004[1] = µitem
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp004[2] = µsource
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							πTemp004[3] = µoffset
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label9
						Label9:
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
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 1254: def pop(self, i=-1):
					πF.SetLineno(1254)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp027, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "i", Def: πTemp027}
					πTemp026 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1255: if self.parent:
							πF.SetLineno(1255)
						Label1:
							// line 1256: index = (len(self.data) + i) % len(self.data)
							πF.SetLineno(1256)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, µi); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							µindex = πTemp001
							// line 1257: self.parent.pop(index + self.parent_offset)
							πF.SetLineno(1257)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µindex, πTemp003); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							// line 1258: self.items.pop(i)
							πF.SetLineno(1258)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1259: return self.data.pop(i)
							πF.SetLineno(1259)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 1261: def trim_start(self, n=1):
					πF.SetLineno(1261)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(1).ToObject()}
					πTemp027 = πg.NewFunction(πg.NewCode("trim_start", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
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
							// line 1262: """
							πF.SetLineno(1262)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
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
							if πTemp001, πE = πg.GT(πF, µn, πTemp004); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 1265: if n > len(self.data):
							πF.SetLineno(1265)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003 = πg.NewTuple2(µn, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Size of trim too large; can't trim %s items from a list of size %s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1266: raise IndexError("Size of trim too large; can't trim %s items "
							πF.SetLineno(1266)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
							// line 1268: elif n < 0:
							πF.SetLineno(1268)
						Label2:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Trim size must be >= 0.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1269: raise IndexError('Trim size must be >= 0.')
							πF.SetLineno(1269)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 1270: del self.data[:n]
							πF.SetLineno(1270)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							// line 1271: del self.items[:n]
							πF.SetLineno(1271)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 1272: if self.parent:
							πF.SetLineno(1272)
						Label4:
							// line 1273: self.parent_offset += n
							πF.SetLineno(1273)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent_offset, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp001, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent_offset, πTemp003); πE != nil {
								continue
							}
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtrim_start.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 1262: """
					πF.SetLineno(1262)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("\n        Remove items from the start of the list, without touching the parent.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßtrim_start); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
						continue
					}
					// line 1275: def trim_end(self, n=1):
					πF.SetLineno(1275)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(1).ToObject()}
					πTemp028 = πg.NewFunction(πg.NewCode("trim_end", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
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
							// line 1276: """
							πF.SetLineno(1276)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
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
							if πTemp001, πE = πg.GT(πF, µn, πTemp004); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 1279: if n > len(self.data):
							πF.SetLineno(1279)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003 = πg.NewTuple2(µn, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Size of trim too large; can't trim %s items from a list of size %s.").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1280: raise IndexError("Size of trim too large; can't trim %s items "
							πF.SetLineno(1280)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
							// line 1282: elif n < 0:
							πF.SetLineno(1282)
						Label2:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Trim size must be >= 0.").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1283: raise IndexError('Trim size must be >= 0.')
							πF.SetLineno(1283)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 1284: del self.data[-n:]
							πF.SetLineno(1284)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, µn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							// line 1285: del self.items[-n:]
							πF.SetLineno(1285)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, µn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtrim_end.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1276: """
					πF.SetLineno(1276)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πg.NewStr("\n        Remove items from the end of the list, without touching the parent.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßtrim_end); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp030, ß__doc__, πTemp029); πE != nil {
						continue
					}
					// line 1287: def remove(self, item):
					πF.SetLineno(1287)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
						var µindex *πg.Object = πg.UnboundLocal
						_ = µindex
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
							// line 1288: index = self.index(item)
							πF.SetLineno(1288)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µindex = πTemp003
							// line 1289: del self[index]
							πF.SetLineno(1289)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = µindex
							if πE = πg.DelItem(πF, µself, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremove.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1291: def count(self, item): return self.data.count(item)
					πF.SetLineno(1291)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("count", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
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
							// line 1291: def count(self, item): return self.data.count(item)
							πF.SetLineno(1291)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcount, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcount.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 1292: def index(self, item): return self.data.index(item)
					πF.SetLineno(1292)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "item", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("index", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitem *πg.Object = πArgs[1]
						_ = µitem
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
							// line 1292: def index(self, item): return self.data.index(item)
							πF.SetLineno(1292)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindex, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßindex.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 1294: def reverse(self):
					πF.SetLineno(1294)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("reverse", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 1295: self.data.reverse()
							πF.SetLineno(1295)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1296: self.items.reverse()
							πF.SetLineno(1296)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1297: self.parent = None
							πF.SetLineno(1297)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreverse.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 1299: def sort(self, *args):
					πF.SetLineno(1299)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("sort", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µtmp *πg.Object = πg.UnboundLocal
						_ = µtmp
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
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
							// line 1300: tmp = sorted(zip(self.data, self.items), *args)
							πF.SetLineno(1300)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp003, πTemp001, µargs, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtmp = πTemp004
							// line 1301: self.data = [entry[0] for entry in tmp]
							πF.SetLineno(1301)
							πTemp005 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µentry *πg.Object = πg.UnboundLocal
								_ = µentry
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
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
										if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µtmp); πE != nil {
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
											µentry = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1301: self.data = [entry[0] for entry in tmp]
										πF.SetLineno(1301)
										πTemp004 = πg.NewInt(0).ToObject()
										if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µentry, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
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
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp006); πE != nil {
								continue
							}
							// line 1302: self.items = [entry[1] for entry in tmp]
							πF.SetLineno(1302)
							πTemp005 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µentry *πg.Object = πg.UnboundLocal
								_ = µentry
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
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
										if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µtmp); πE != nil {
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
											µentry = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1302: self.items = [entry[1] for entry in tmp]
										πF.SetLineno(1302)
										πTemp004 = πg.NewInt(1).ToObject()
										if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µentry, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
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
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßitems, πTemp007); πE != nil {
								continue
							}
							// line 1303: self.parent = None
							πF.SetLineno(1303)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp007); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsort.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 1305: def info(self, i):
					πF.SetLineno(1305)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("info", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var πTemp001 *πg.Object
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
							// line 1306: """Return source & offset for index `i`."""
							πF.SetLineno(1306)
							// line 1307: try:
							πF.SetLineno(1307)
							πF.PushCheckpoint(2)
							// line 1308: return self.items[i]
							πF.SetLineno(1308)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
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
							// line 1309: except IndexError:
							πF.SetLineno(1309)
						Label3:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Eq(πF, µi, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 1310: if i == len(self.data):     # Just past the end
							πF.SetLineno(1310)
						Label4:
							// line 1311: return self.items[i - 1][0], None
							πF.SetLineno(1311)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp009, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp002).ToObject()
							πR = πTemp001
							continue
							goto Label6
						Label5:
							// line 1313: raise
							πF.SetLineno(1313)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label6
						Label6:
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßinfo.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 1306: """Return source & offset for index `i`."""
					πF.SetLineno(1306)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp035}, πg.NewStr("Return source & offset for index `i`.").ToObject()); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßinfo); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp036, ß__doc__, πTemp035); πE != nil {
						continue
					}
					// line 1315: def source(self, i):
					πF.SetLineno(1315)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("source", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
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
							// line 1316: """Return source for index `i`."""
							πF.SetLineno(1316)
							// line 1317: return self.info(i)[0]
							πF.SetLineno(1317)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßsource.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 1316: """Return source for index `i`."""
					πF.SetLineno(1316)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πg.NewStr("Return source for index `i`.").ToObject()); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßsource); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp037, ß__doc__, πTemp036); πE != nil {
						continue
					}
					// line 1319: def offset(self, i):
					πF.SetLineno(1319)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("offset", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
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
							// line 1320: """Return offset for index `i`."""
							πF.SetLineno(1320)
							// line 1321: return self.info(i)[1]
							πF.SetLineno(1321)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßoffset.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 1320: """Return offset for index `i`."""
					πF.SetLineno(1320)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp037}, πg.NewStr("Return offset for index `i`.").ToObject()); πE != nil {
						continue
					}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßoffset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp038, ß__doc__, πTemp037); πE != nil {
						continue
					}
					// line 1323: def disconnect(self):
					πF.SetLineno(1323)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("disconnect", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
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
							// line 1324: """Break link between this list and parent list."""
							πF.SetLineno(1324)
							// line 1325: self.parent = None
							πF.SetLineno(1325)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparent, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdisconnect.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 1324: """Break link between this list and parent list."""
					πF.SetLineno(1324)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp038}, πg.NewStr("Break link between this list and parent list.").ToObject()); πE != nil {
						continue
					}
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßdisconnect); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp039, ß__doc__, πTemp038); πE != nil {
						continue
					}
					// line 1327: def xitems(self):
					πF.SetLineno(1327)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("xitems", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
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
								// line 1328: """Return iterator yielding (source, offset, value) tuples."""
								πF.SetLineno(1328)
								πTemp002 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								πTemp002[0] = πTemp003
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
									continue
								}
								πTemp002[1] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
									continue
								}
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
								if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp006 = !isStop
								} else {
									πTemp006 = true
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}}}, πTemp003); πE != nil {
										continue
									}
									µvalue = πTemp004
									µsource = πTemp007
									µoffset = πTemp008
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 1330: yield (source, offset, value)
								πF.SetLineno(1330)
								if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp003 = πg.NewTuple3(µsource, µoffset, µvalue).ToObject()
								πF.PushCheckpoint(4)
								return πTemp003, nil
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
					if πE = πClass.SetItem(πF, ßxitems.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 1328: """Return iterator yielding (source, offset, value) tuples."""
					πF.SetLineno(1328)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp039}, πg.NewStr("Return iterator yielding (source, offset, value) tuples.").ToObject()); πE != nil {
						continue
					}
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßxitems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp040, ß__doc__, πTemp039); πE != nil {
						continue
					}
					// line 1332: def pprint(self):
					πF.SetLineno(1332)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("pprint", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1333: """Print the list in `grep` format (`source:offset:value` lines)"""
							πF.SetLineno(1333)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßxitems, nil); πE != nil {
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
							// line 1335: print("%s:%d:%s" % line)
							πF.SetLineno(1335)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s:%d:%s").ToObject(), µline); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpprint.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 1333: """Print the list in `grep` format (`source:offset:value` lines)"""
					πF.SetLineno(1333)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp040}, πg.NewStr("Print the list in `grep` format (`source:offset:value` lines)").ToObject()); πE != nil {
						continue
					}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßpprint); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp041, ß__doc__, πTemp040); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ViewList").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßViewList.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1338: class StringList(ViewList):
			πF.SetLineno(1338)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßViewList); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StringList", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1340: """A `ViewList` with string-specific methods."""
					πF.SetLineno(1340)
					// line 1340: """A `ViewList` with string-specific methods."""
					πF.SetLineno(1340)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A `ViewList` with string-specific methods.").ToObject()); πE != nil {
						continue
					}
					// line 1342: def trim_left(self, length, start=0, end=sys.maxsize):
					πF.SetLineno(1342)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "length", Def: nil}
					πTemp002[2] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaxsize, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "end", Def: πTemp004}
					πTemp001 = πg.NewFunction(πg.NewCode("trim_left", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlength *πg.Object = πArgs[1]
						_ = µlength
						var µstart *πg.Object = πArgs[2]
						_ = µstart
						var µend *πg.Object = πArgs[3]
						_ = µend
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
							// line 1343: """
							πF.SetLineno(1343)
							// line 1348: self.data[start:end] = [line[length:]
							πF.SetLineno(1348)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 bool
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
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
											continue
										}
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
										if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp006 = !isStop
										} else {
											πTemp006 = true
											µline = πTemp002
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1348: self.data[start:end] = [line[length:]
										πF.SetLineno(1348)
										if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µlength, πg.None, πg.None}, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp002 = πSent
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
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
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtrim_left.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1343: """
					πF.SetLineno(1343)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Trim `length` characters off the beginning of each item, in-place,\n        from index `start` to `end`.  No whitespace-checking is done on the\n        trimmed text.  Does not affect slice parent.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtrim_left); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 1351: def get_text_block(self, start, flush_left=False):
					πF.SetLineno(1351)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "start", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "flush_left", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("get_text_block", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstart *πg.Object = πArgs[1]
						_ = µstart
						var µflush_left *πg.Object = πArgs[2]
						_ = µflush_left
						var µend *πg.Object = πg.UnboundLocal
						_ = µend
						var µlast *πg.Object = πg.UnboundLocal
						_ = µlast
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µoffset *πg.Object = πg.UnboundLocal
						_ = µoffset
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
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
							// line 1352: """
							πF.SetLineno(1352)
							// line 1359: end = start
							πF.SetLineno(1359)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							µend = µstart
							// line 1360: last = len(self.data)
							πF.SetLineno(1360)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlast = πTemp003
							// line 1361: while end < last:
							πF.SetLineno(1361)
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
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µend, µlast); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1362: line = self.data[end]
							πF.SetLineno(1362)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp002 = µend
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							µline = πTemp003
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µline, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 1363: if not line.strip():
							πF.SetLineno(1363)
						Label4:
							// line 1364: break
							πF.SetLineno(1364)
							πTemp004 = true
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µflush_left, "flush_left"); πE != nil {
								continue
							}
							πTemp002 = µflush_left
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label6
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µline, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label6:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1365: if flush_left and (line[0] == ' '):
							πF.SetLineno(1365)
						Label7:
							// line 1366: source, offset = self.info(end)
							πF.SetLineno(1366)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp001[0] = µend
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinfo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µsource = πTemp002
							µoffset = πTemp006
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001[1] = µsource
							if πE = πg.CheckLocal(πF, µoffset, "offset"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µoffset, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUnexpectedIndentationError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1367: raise UnexpectedIndentationError(self[start:end], source,
							πF.SetLineno(1367)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label8
						Label8:
							// line 1369: end += 1
							πF.SetLineno(1369)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µend = πTemp002
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1370: return self[start:end]
							πF.SetLineno(1370)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µself, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßget_text_block.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1352: """
					πF.SetLineno(1352)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Return a contiguous block of text.\n\n        If `flush_left` is true, raise `UnexpectedIndentationError` if an\n        indented line is encountered before the text block ends (with a blank\n        line).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßget_text_block); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 1372: def get_indented(self, start=0, until_blank=False, strip_indent=True,
					πF.SetLineno(1372)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "until_blank", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "strip_indent", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "block_indent", Def: πTemp005}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "first_indent", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("get_indented", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstart *πg.Object = πArgs[1]
						_ = µstart
						var µuntil_blank *πg.Object = πArgs[2]
						_ = µuntil_blank
						var µstrip_indent *πg.Object = πArgs[3]
						_ = µstrip_indent
						var µblock_indent *πg.Object = πArgs[4]
						_ = µblock_indent
						var µfirst_indent *πg.Object = πArgs[5]
						_ = µfirst_indent
						var µindent *πg.Object = πg.UnboundLocal
						_ = µindent
						var µend *πg.Object = πg.UnboundLocal
						_ = µend
						var µlast *πg.Object = πg.UnboundLocal
						_ = µlast
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µblank_finish *πg.Object = πg.UnboundLocal
						_ = µblank_finish
						var µstripped *πg.Object = πg.UnboundLocal
						_ = µstripped
						var µline_indent *πg.Object = πg.UnboundLocal
						_ = µline_indent
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
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
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 1374: """
							πF.SetLineno(1374)
							// line 1394: indent = block_indent           # start with None if unknown
							πF.SetLineno(1394)
							if πE = πg.CheckLocal(πF, µblock_indent, "block_indent"); πE != nil {
								continue
							}
							µindent = µblock_indent
							// line 1395: end = start
							πF.SetLineno(1395)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							µend = µstart
							if πE = πg.CheckLocal(πF, µblock_indent, "block_indent"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µblock_indent != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µfirst_indent, "first_indent"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µfirst_indent == πTemp004).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 1396: if block_indent is not None and first_indent is None:
							πF.SetLineno(1396)
						Label2:
							// line 1397: first_indent = block_indent
							πF.SetLineno(1397)
							if πE = πg.CheckLocal(πF, µblock_indent, "block_indent"); πE != nil {
								continue
							}
							µfirst_indent = µblock_indent
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µfirst_indent, "first_indent"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µfirst_indent != πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 1398: if first_indent is not None:
							πF.SetLineno(1398)
						Label4:
							// line 1399: end += 1
							πF.SetLineno(1399)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µend = πTemp001
							goto Label5
						Label5:
							// line 1400: last = len(self.data)
							πF.SetLineno(1400)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µlast = πTemp003
							// line 1401: while end < last:
							πF.SetLineno(1401)
							πF.PushCheckpoint(7)
							πTemp002 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µend, µlast); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 1402: line = self.data[end]
							πF.SetLineno(1402)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp001 = µend
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µline = πTemp003
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp001 = µline
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label9
							}
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µline, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.NE(πF, πTemp009, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µblock_indent, "block_indent"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp008 = πg.GetBool(µblock_indent != πTemp009).ToObject()
							πTemp004 = πTemp008
							if πTemp010, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µblock_indent, "block_indent"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µblock_indent, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µline, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp009
						Label11:
							πTemp003 = πTemp004
						Label10:
							πTemp001 = πTemp003
						Label9:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 1403: if line and (line[0] != ' '
							πF.SetLineno(1403)
						Label12:
							// line 1408: blank_finish = ((end > start)
							πF.SetLineno(1408)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µend, µstart); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp008, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label14:
							µblank_finish = πTemp001
							// line 1410: break
							πF.SetLineno(1410)
							πTemp002 = true
							continue
							goto Label13
						Label13:
							// line 1411: stripped = line.lstrip()
							πF.SetLineno(1411)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µline, ßlstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstripped = πTemp003
							if πE = πg.CheckLocal(πF, µstripped, "stripped"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µstripped); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µblock_indent, "block_indent"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µblock_indent == πTemp003).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label16
							}
							goto Label17
							// line 1412: if not stripped:            # blank line
							πF.SetLineno(1412)
						Label15:
							if πE = πg.CheckLocal(πF, µuntil_blank, "until_blank"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µuntil_blank); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label18
							}
							goto Label19
							// line 1413: if until_blank:
							πF.SetLineno(1413)
						Label18:
							// line 1414: blank_finish = 1
							πF.SetLineno(1414)
							µblank_finish = πg.NewInt(1).ToObject()
							// line 1415: break
							πF.SetLineno(1415)
							πTemp002 = true
							continue
							goto Label19
						Label19:
							goto Label17
							// line 1416: elif block_indent is None:
							πF.SetLineno(1416)
						Label16:
							// line 1417: line_indent = len(line) - len(stripped)
							πF.SetLineno(1417)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp005[0] = µline
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstripped, "stripped"); πE != nil {
								continue
							}
							πTemp005[0] = µstripped
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Sub(πF, πTemp004, πTemp008); πE != nil {
								continue
							}
							µline_indent = πTemp001
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µindent == πTemp003).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label20
							}
							goto Label21
							// line 1418: if indent is None:
							πF.SetLineno(1418)
						Label20:
							// line 1419: indent = line_indent
							πF.SetLineno(1419)
							if πE = πg.CheckLocal(πF, µline_indent, "line_indent"); πE != nil {
								continue
							}
							µindent = µline_indent
							goto Label22
						Label21:
							// line 1421: indent = min(indent, line_indent)
							πF.SetLineno(1421)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp005[0] = µindent
							if πE = πg.CheckLocal(πF, µline_indent, "line_indent"); πE != nil {
								continue
							}
							πTemp005[1] = µline_indent
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µindent = πTemp003
							goto Label22
						Label22:
							goto Label17
						Label17:
							// line 1422: end += 1
							πF.SetLineno(1422)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µend = πTemp001
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							// line 1424: blank_finish = 1            # block ends at end of lines
							πF.SetLineno(1424)
							µblank_finish = πg.NewInt(1).ToObject()
						Label8:
							// line 1425: block = self[start:end]
							πF.SetLineno(1425)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp003
							if πE = πg.CheckLocal(πF, µfirst_indent, "first_indent"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µfirst_indent != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label23
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πTemp001 = µblock
						Label23:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label24
							}
							goto Label25
							// line 1426: if first_indent is not None and block:
							πF.SetLineno(1426)
						Label24:
							// line 1427: block.data[0] = block.data[0][first_indent:]
							πF.SetLineno(1427)
							if πE = πg.CheckLocal(πF, µfirst_indent, "first_indent"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µfirst_indent, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.SetItem(πF, πTemp004, πTemp008, πTemp001); πE != nil {
								continue
							}
							goto Label25
						Label25:
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp001 = µindent
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label26
							}
							if πE = πg.CheckLocal(πF, µstrip_indent, "strip_indent"); πE != nil {
								continue
							}
							πTemp001 = µstrip_indent
						Label26:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label27
							}
							goto Label28
							// line 1428: if indent and strip_indent:
							πF.SetLineno(1428)
						Label27:
							// line 1429: block.trim_left(indent, start=(first_indent is not None))
							πF.SetLineno(1429)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp005[0] = µindent
							if πE = πg.CheckLocal(πF, µfirst_indent, "first_indent"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µfirst_indent != πTemp003).ToObject()
							πTemp011 = πg.KWArgs{
								{"start", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µblock, ßtrim_left, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label28
						Label28:
							// line 1430: return block, indent or 0, blank_finish
							πF.SetLineno(1430)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp003 = µindent
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label29
							}
							πTemp003 = πg.NewInt(0).ToObject()
						Label29:
							if πE = πg.CheckLocal(πF, µblank_finish, "blank_finish"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(µblock, πTemp003, µblank_finish).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_indented.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1374: """
					πF.SetLineno(1374)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Extract and return a StringList of indented lines of text.\n\n        Collect all lines with indentation, determine the minimum indentation,\n        remove the minimum indentation from all indented lines (unless\n        `strip_indent` is false), and return them. All lines up to but not\n        including the first unindented line will be returned.\n\n        :Parameters:\n          - `start`: The index of the first line to examine.\n          - `until_blank`: Stop collecting at the first blank line if true.\n          - `strip_indent`: Strip common leading indent if true (default).\n          - `block_indent`: The indent of the entire block, if known.\n          - `first_indent`: The indent of the first line, if known.\n\n        :Return:\n          - a StringList of indented lines with mininum indent removed;\n          - the amount of the indent;\n          - a boolean: did the indented block finish with a blank line or EOF?\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßget_indented); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 1432: def get_2D_block(self, top, left, bottom, right, strip_indent=True):
					πF.SetLineno(1432)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "top", Def: nil}
					πTemp002[2] = πg.Param{Name: "left", Def: nil}
					πTemp002[3] = πg.Param{Name: "bottom", Def: nil}
					πTemp002[4] = πg.Param{Name: "right", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "strip_indent", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("get_2D_block", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtop *πg.Object = πArgs[1]
						_ = µtop
						var µleft *πg.Object = πArgs[2]
						_ = µleft
						var µbottom *πg.Object = πArgs[3]
						_ = µbottom
						var µright *πg.Object = πArgs[4]
						_ = µright
						var µstrip_indent *πg.Object = πArgs[5]
						_ = µstrip_indent
						var µblock *πg.Object = πg.UnboundLocal
						_ = µblock
						var µindent *πg.Object = πg.UnboundLocal
						_ = µindent
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µci *πg.Object = πg.UnboundLocal
						_ = µci
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 []πg.Param
						_ = πTemp013
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 1433: block = self[top:bottom]
							πF.SetLineno(1433)
							if πE = πg.CheckLocal(πF, µtop, "top"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbottom, "bottom"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µtop, µbottom, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp002
							// line 1434: indent = right
							πF.SetLineno(1434)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							µindent = µright
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1437: ci = utils.column_indices(block.data[i])
							πF.SetLineno(1437)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßutils); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßcolumn_indices, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µci = πTemp002
							// line 1438: try:
							πF.SetLineno(1438)
							πF.PushCheckpoint(5)
							// line 1439: left = ci[left]
							πF.SetLineno(1439)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp002 = µleft
							if πE = πg.CheckLocal(πF, µci, "ci"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µci, πTemp002); πE != nil {
								continue
							}
							µleft = πTemp005
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label6
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 1440: except IndexError:
							πF.SetLineno(1440)
						Label6:
							// line 1441: left += len(block.data[i]) - len(ci)
							πF.SetLineno(1441)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µci, "ci"); πE != nil {
								continue
							}
							πTemp003[0] = µci
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Sub(πF, πTemp008, πTemp011); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µleft, πTemp002); πE != nil {
								continue
							}
							µleft = πTemp005
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 1442: try:
							πF.SetLineno(1442)
							πF.PushCheckpoint(8)
							// line 1443: right = ci[right]
							πF.SetLineno(1443)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp002 = µright
							if πE = πg.CheckLocal(πF, µci, "ci"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µci, πTemp002); πE != nil {
								continue
							}
							µright = πTemp005
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 1444: except IndexError:
							πF.SetLineno(1444)
						Label9:
							// line 1445: right += len(block.data[i]) - len(ci)
							πF.SetLineno(1445)
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µci, "ci"); πE != nil {
								continue
							}
							πTemp003[0] = µci
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Sub(πF, πTemp008, πTemp011); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µright, πTemp002); πE != nil {
								continue
							}
							µright = πTemp005
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 1446: block.data[i] = line = block.data[i][left:right].rstrip()
							πF.SetLineno(1446)
							if πE = πg.CheckLocal(πF, µleft, "left"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µleft, µright, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, πTemp012, πTemp008); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp011, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßrstrip, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp011 = µi
							if πE = πg.SetItem(πF, πTemp008, πTemp011, πTemp002); πE != nil {
								continue
							}
							µline = πTemp005
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µline); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 1447: if line:
							πF.SetLineno(1447)
						Label10:
							// line 1448: indent = min(indent, len(line) - len(line.lstrip()))
							πF.SetLineno(1448)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp003[0] = µindent
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp004[0] = µline
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µline, ßlstrip, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp011
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Sub(πF, πTemp008, πTemp011); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µindent = πTemp005
							goto Label11
						Label11:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µstrip_indent, "strip_indent"); πE != nil {
								continue
							}
							πTemp001 = µstrip_indent
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πg.NewInt(0).ToObject(), µindent); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label13
							}
							if πE = πg.CheckLocal(πF, µright, "right"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µindent, µright); πE != nil {
								continue
							}
						Label13:
							πTemp001 = πTemp002
						Label12:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							goto Label15
							// line 1449: if strip_indent and 0 < indent < right:
							πF.SetLineno(1449)
						Label14:
							// line 1450: block.data = [line[indent:] for line in block.data]
							πF.SetLineno(1450)
							πTemp013 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal
								_ = µline
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 *πg.Object
								_ = πTemp005
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
										if πTemp002, πE = πg.GetAttr(πF, µblock, ßdata, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp003 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp003 {
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
											πTemp004 = !isStop
										} else {
											πTemp004 = true
											µline = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 1450: block.data = [line[indent:] for line in block.data]
										πF.SetLineno(1450)
										if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µindent, πg.None, πg.None}, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µline, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
									Label4:
										πTemp002 = πSent
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
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µblock, ßdata, πTemp005); πE != nil {
								continue
							}
							goto Label15
						Label15:
							// line 1451: return block
							πF.SetLineno(1451)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							πR = µblock
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_2D_block.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1453: def pad_double_width(self, pad_char):
					πF.SetLineno(1453)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pad_char", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("pad_double_width", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpad_char *πg.Object = πArgs[1]
						_ = µpad_char
						var µeast_asian_width *πg.Object = πg.UnboundLocal
						_ = µeast_asian_width
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µline *πg.Object = πg.UnboundLocal
						_ = µline
						var µnew *πg.Object = πg.UnboundLocal
						_ = µnew
						var µchar *πg.Object = πg.UnboundLocal
						_ = µchar
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 1454: """
							πF.SetLineno(1454)
							// line 1458: east_asian_width = unicodedata.east_asian_width
							πF.SetLineno(1458)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicodedata); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßeast_asian_width, nil); πE != nil {
								continue
							}
							µeast_asian_width = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1460: line = self.data[i]
							πF.SetLineno(1460)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							µline = πTemp005
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp003[0] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 1461: if isinstance(line, unicode):
							πF.SetLineno(1461)
						Label4:
							// line 1462: new = []
							πF.SetLineno(1462)
							πTemp003 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							µnew = πTemp002
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µline); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp007 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µchar = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 1464: new.append(char)
							πF.SetLineno(1464)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp003[0] = µchar
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnew, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp003[0] = µchar
							if πE = πg.CheckLocal(πF, µeast_asian_width, "east_asian_width"); πE != nil {
								continue
							}
							if πTemp008, πE = µeast_asian_width.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp009, πE = πg.Contains(πF, ßWF.ToObject(), πTemp008); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 1465: if east_asian_width(char) in 'WF': # 'W'ide & 'F'ull-width
							πF.SetLineno(1465)
						Label9:
							// line 1466: new.append(pad_char)
							πF.SetLineno(1466)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpad_char, "pad_char"); πE != nil {
								continue
							}
							πTemp003[0] = µpad_char
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnew, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 1467: self.data[i] = ''.join(new)
							πF.SetLineno(1467)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp003[0] = µnew
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp010 = µi
							if πE = πg.SetItem(πF, πTemp008, πTemp010, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpad_double_width.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1454: """
					πF.SetLineno(1454)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Pad all double-width characters in self by appending `pad_char` to each.\n        For East Asian language support.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßpad_double_width); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 1469: def replace(self, old, new):
					πF.SetLineno(1469)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "old", Def: nil}
					πTemp002[2] = πg.Param{Name: "new", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("replace", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µold *πg.Object = πArgs[1]
						_ = µold
						var µnew *πg.Object = πArgs[2]
						_ = µnew
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
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
							// line 1470: """Replace all occurrences of substring `old` with `new`."""
							πF.SetLineno(1470)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1472: self.data[i] = self.data[i].replace(old, new)
							πF.SetLineno(1472)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µold, "old"); πE != nil {
								continue
							}
							πTemp002[0] = µold
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp002[1] = µnew
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp004); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßreplace.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1470: """Replace all occurrences of substring `old` with `new`."""
					πF.SetLineno(1470)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Replace all occurrences of substring `old` with `new`.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßreplace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StringList").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringList.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1475: class StateMachineError(Exception): pass
			πF.SetLineno(1475)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StateMachineError", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1475: class StateMachineError(Exception): pass
					πF.SetLineno(1475)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StateMachineError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStateMachineError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1476: class UnknownStateError(StateMachineError): pass
			πF.SetLineno(1476)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("UnknownStateError", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1476: class UnknownStateError(StateMachineError): pass
					πF.SetLineno(1476)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UnknownStateError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnknownStateError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1477: class DuplicateStateError(StateMachineError): pass
			πF.SetLineno(1477)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DuplicateStateError", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1477: class DuplicateStateError(StateMachineError): pass
					πF.SetLineno(1477)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DuplicateStateError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDuplicateStateError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1478: class UnknownTransitionError(StateMachineError): pass
			πF.SetLineno(1478)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("UnknownTransitionError", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1478: class UnknownTransitionError(StateMachineError): pass
					πF.SetLineno(1478)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UnknownTransitionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnknownTransitionError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1479: class DuplicateTransitionError(StateMachineError): pass
			πF.SetLineno(1479)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("DuplicateTransitionError", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1479: class DuplicateTransitionError(StateMachineError): pass
					πF.SetLineno(1479)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DuplicateTransitionError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDuplicateTransitionError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1480: class TransitionPatternNotFound(StateMachineError): pass
			πF.SetLineno(1480)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TransitionPatternNotFound", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1480: class TransitionPatternNotFound(StateMachineError): pass
					πF.SetLineno(1480)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TransitionPatternNotFound").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransitionPatternNotFound.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1481: class TransitionMethodNotFound(StateMachineError): pass
			πF.SetLineno(1481)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TransitionMethodNotFound", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1481: class TransitionMethodNotFound(StateMachineError): pass
					πF.SetLineno(1481)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TransitionMethodNotFound").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransitionMethodNotFound.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1482: class UnexpectedIndentationError(StateMachineError): pass
			πF.SetLineno(1482)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßStateMachineError); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("UnexpectedIndentationError", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1482: class UnexpectedIndentationError(StateMachineError): pass
					πF.SetLineno(1482)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UnexpectedIndentationError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnexpectedIndentationError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1485: class TransitionCorrection(Exception):
			πF.SetLineno(1485)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TransitionCorrection", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1487: """
					πF.SetLineno(1487)
					// line 1487: """
					πF.SetLineno(1487)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise from within a transition method to switch to another transition.\n\n    Raise with one argument, the new transition name.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TransitionCorrection").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTransitionCorrection.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1494: class StateCorrection(Exception):
			πF.SetLineno(1494)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StateCorrection", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1496: """
					πF.SetLineno(1496)
					// line 1496: """
					πF.SetLineno(1496)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise from within a transition method to switch to another state.\n\n    Raise with one or two arguments: new state name, and an optional new\n    transition name.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StateCorrection").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStateCorrection.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 1503: def string2lines(astring, tab_width=8, convert_whitespace=False,
			πF.SetLineno(1503)
			πTemp007 = make([]πg.Param, 4)
			πTemp007[0] = πg.Param{Name: "astring", Def: nil}
			πTemp007[1] = πg.Param{Name: "tab_width", Def: πg.NewInt(8).ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp007[2] = πg.Param{Name: "convert_whitespace", Def: πTemp003}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("[\x0b\x0c]").ToObject()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp007[3] = πg.Param{Name: "whitespace", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("string2lines", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µastring *πg.Object = πArgs[0]
				_ = µastring
				var µtab_width *πg.Object = πArgs[1]
				_ = µtab_width
				var µconvert_whitespace *πg.Object = πArgs[2]
				_ = µconvert_whitespace
				var µwhitespace *πg.Object = πArgs[3]
				_ = µwhitespace
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
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
					// line 1505: """
					πF.SetLineno(1505)
					if πE = πg.CheckLocal(πF, µconvert_whitespace, "convert_whitespace"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µconvert_whitespace); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 1518: if convert_whitespace:
					πF.SetLineno(1518)
				Label1:
					// line 1519: astring = whitespace.sub(' ', astring)
					πF.SetLineno(1519)
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewStr(" ").ToObject()
					if πE = πg.CheckLocal(πF, µastring, "astring"); πE != nil {
						continue
					}
					πTemp002[1] = µastring
					if πE = πg.CheckLocal(πF, µwhitespace, "whitespace"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µwhitespace, ßsub, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µastring = πTemp004
					goto Label2
				Label2:
					// line 1522: return [s.expandtabs(tab_width).rstrip() for s in astring.splitlines()]
					πF.SetLineno(1522)
					πTemp005 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
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
								case 4:
									goto Label4
								default:
									panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µastring, "astring"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µastring, ßsplitlines, nil); πE != nil {
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
									µs = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 1522: return [s.expandtabs(tab_width).rstrip() for s in astring.splitlines()]
								πF.SetLineno(1522)
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtab_width, "tab_width"); πE != nil {
									continue
								}
								πTemp006[0] = µtab_width
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µs, ßexpandtabs, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßrstrip, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp002 = πSent
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
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßstring2lines.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 1505: """
			πF.SetLineno(1505)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Return a list of one-line strings with tabs expanded, no newlines, and\n    trailing whitespace stripped.\n\n    Each tab is expanded with between 1 and `tab_width` spaces, so that the\n    next character's index becomes a multiple of `tab_width` (8 by default).\n\n    Parameters:\n\n    - `astring`: a multi-line string.\n    - `tab_width`: the number of columns between tab stops.\n    - `convert_whitespace`: convert form feeds and vertical tabs to spaces?\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßstring2lines); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 1524: def _exception_data():
			πF.SetLineno(1524)
			πTemp007 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("_exception_data", "/usr/lib/python2.7/site-packages/docutils/statemachine.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtype *πg.Object = πg.UnboundLocal
				_ = µtype
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
				var µtraceback *πg.Object = πg.UnboundLocal
				_ = µtraceback
				var µcode *πg.Object = πg.UnboundLocal
				_ = µcode
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
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
					// line 1525: """
					πF.SetLineno(1525)
					// line 1534: type, value, traceback = sys.exc_info()
					πF.SetLineno(1534)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
						continue
					}
					µtype = πTemp002
					µvalue = πTemp003
					µtraceback = πTemp004
					// line 1535: while traceback.tb_next:
					πF.SetLineno(1535)
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
					if πE = πg.CheckLocal(πF, µtraceback, "traceback"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtraceback, ßtb_next, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 1536: traceback = traceback.tb_next
					πF.SetLineno(1536)
					if πE = πg.CheckLocal(πF, µtraceback, "traceback"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtraceback, ßtb_next, nil); πE != nil {
						continue
					}
					µtraceback = πTemp001
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 1537: code = traceback.tb_frame.f_code
					πF.SetLineno(1537)
					if πE = πg.CheckLocal(πF, µtraceback, "traceback"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µtraceback, ßtb_frame, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßf_code, nil); πE != nil {
						continue
					}
					µcode = πTemp002
					// line 1538: return (type.__name__, value, code.co_filename, traceback.tb_lineno,
					πF.SetLineno(1538)
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtype, ß__name__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µcode, ßco_filename, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtraceback, "traceback"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µtraceback, ßtb_lineno, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µcode, ßco_name, nil); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple5(πTemp002, µvalue, πTemp003, πTemp004, πTemp007).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_exception_data.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 1525: """
			πF.SetLineno(1525)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Return exception information:\n\n    - the exception's class name;\n    - the exception object;\n    - the name of the file containing the offending code;\n    - the line number of the offending code;\n    - the function name of the offending code.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ß_exception_data); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("statemachine", Code)
}
