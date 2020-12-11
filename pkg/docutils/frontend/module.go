package frontend

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/ConfigParser"
	_ "github.com/pygolin/stdlib/pkg/codecs"
	_ "github.com/pygolin/stdlib/pkg/configparser"
	_ "github.com/pygolin/stdlib/pkg/optparse"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/os/path"
	_ "github.com/pygolin/stdlib/pkg/pwd"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/warnings"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/frontend.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß1 := πg.InternStr("1")
		ßATTRS := πg.InternStr("ATTRS")
		ßAttributeError := πg.InternStr("AttributeError")
		ßConfigDeprecationWarning := πg.InternStr("ConfigDeprecationWarning")
		ßConfigParser := πg.InternStr("ConfigParser")
		ßDOCUTILSCONFIG := πg.InternStr("DOCUTILSCONFIG")
		ßDependencyList := πg.InternStr("DependencyList")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßErrorOutput := πg.InternStr("ErrorOutput")
		ßErrorString := πg.InternStr("ErrorString")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßHOME := πg.InternStr("HOME")
		ßIOError := πg.InternStr("IOError")
		ßImportError := πg.InternStr("ImportError")
		ßKeyError := πg.InternStr("KeyError")
		ßLookupError := πg.InternStr("LookupError")
		ßNone := πg.InternStr("None")
		ßOption := πg.InternStr("Option")
		ßOptionGroup := πg.InternStr("OptionGroup")
		ßOptionParser := πg.InternStr("OptionParser")
		ßOptionValueError := πg.InternStr("OptionValueError")
		ßRawConfigParser := πg.InternStr("RawConfigParser")
		ßSUPPRESS_HELP := πg.InternStr("SUPPRESS_HELP")
		ßSafeString := πg.InternStr("SafeString")
		ßSettingsSpec := πg.InternStr("SettingsSpec")
		ßTitledHelpFormatter := πg.InternStr("TitledHelpFormatter")
		ßTrue := πg.InternStr("True")
		ßUnicodeDecodeError := πg.InternStr("UnicodeDecodeError")
		ßValueError := πg.InternStr("ValueError")
		ßValues := πg.InternStr("Values")
		ß_ := πg.InternStr("_")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__docformat__ := πg.InternStr("__docformat__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß__version_details__ := πg.InternStr("__version_details__")
		ß_config_files := πg.InternStr("_config_files")
		ß_destination := πg.InternStr("_destination")
		ß_disable_config := πg.InternStr("_disable_config")
		ß_error_handler := πg.InternStr("_error_handler")
		ß_files := πg.InternStr("_files")
		ß_source := πg.InternStr("_source")
		ß_stderr := πg.InternStr("_stderr")
		ß_update_loose := πg.InternStr("_update_loose")
		ßabspath := πg.InternStr("abspath")
		ßaction := πg.InternStr("action")
		ßadd_option := πg.InternStr("add_option")
		ßadd_option_group := πg.InternStr("add_option_group")
		ßadd_section := πg.InternStr("add_section")
		ßappend := πg.InternStr("append")
		ßascii := πg.InternStr("ascii")
		ßbackslashreplace := πg.InternStr("backslashreplace")
		ßbool := πg.InternStr("bool")
		ßbooleans := πg.InternStr("booleans")
		ßcallback := πg.InternStr("callback")
		ßcallback_args := πg.InternStr("callback_args")
		ßcheck_args := πg.InternStr("check_args")
		ßcheck_values := πg.InternStr("check_values")
		ßchoices := πg.InternStr("choices")
		ßclose := πg.InternStr("close")
		ßcodecs := πg.InternStr("codecs")
		ßcomponents := πg.InternStr("components")
		ßconfig_files := πg.InternStr("config_files")
		ßconfig_section := πg.InternStr("config_section")
		ßconfig_section_dependencies := πg.InternStr("config_section_dependencies")
		ßconst := πg.InternStr("const")
		ßcopy := πg.InternStr("copy")
		ßdatestamp := πg.InternStr("datestamp")
		ßdebug := πg.InternStr("debug")
		ßdefault := πg.InternStr("default")
		ßdefault_error_encoding := πg.InternStr("default_error_encoding")
		ßdefault_error_encoding_error_handler := πg.InternStr("default_error_encoding_error_handler")
		ßdefaults := πg.InternStr("defaults")
		ßdest := πg.InternStr("dest")
		ßdirname := πg.InternStr("dirname")
		ßdocutils := πg.InternStr("docutils")
		ßen := πg.InternStr("en")
		ßencode := πg.InternStr("encode")
		ßencoding := πg.InternStr("encoding")
		ßendswith := πg.InternStr("endswith")
		ßentry := πg.InternStr("entry")
		ßenviron := πg.InternStr("environ")
		ßerror := πg.InternStr("error")
		ßexit_status_level := πg.InternStr("exit_status_level")
		ßexpanduser := πg.InternStr("expanduser")
		ßexpose_internals := πg.InternStr("expose_internals")
		ßextend := πg.InternStr("extend")
		ßfalse := πg.InternStr("false")
		ßfilter_settings_spec := πg.InternStr("filter_settings_spec")
		ßfootnote_backlinks := πg.InternStr("footnote_backlinks")
		ßgeneral := πg.InternStr("general")
		ßgenerator := πg.InternStr("generator")
		ßget := πg.InternStr("get")
		ßget_config_file_settings := πg.InternStr("get_config_file_settings")
		ßget_default_values := πg.InternStr("get_default_values")
		ßget_option_by_dest := πg.InternStr("get_option_by_dest")
		ßget_section := πg.InternStr("get_section")
		ßget_standard_config_files := πg.InternStr("get_standard_config_files")
		ßget_standard_config_settings := πg.InternStr("get_standard_config_settings")
		ßgetattr := πg.InternStr("getattr")
		ßgetcwd := πg.InternStr("getcwd")
		ßgetcwdu := πg.InternStr("getcwdu")
		ßhalt_level := πg.InternStr("halt_level")
		ßhandle_old_config := πg.InternStr("handle_old_config")
		ßhas_option := πg.InternStr("has_option")
		ßhas_section := πg.InternStr("has_section")
		ßhasattr := πg.InternStr("hasattr")
		ßhelp := πg.InternStr("help")
		ßid := πg.InternStr("id")
		ßinfo := πg.InternStr("info")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlanguage_code := πg.InternStr("language_code")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlists := πg.InternStr("lists")
		ßlocale_encoding := πg.InternStr("locale_encoding")
		ßlookup := πg.InternStr("lookup")
		ßlookup_error := πg.InternStr("lookup_error")
		ßlower := πg.InternStr("lower")
		ßmake_id := πg.InternStr("make_id")
		ßmake_one_path_absolute := πg.InternStr("make_one_path_absolute")
		ßmake_paths_absolute := πg.InternStr("make_paths_absolute")
		ßmetavar := πg.InternStr("metavar")
		ßno := πg.InternStr("no")
		ßnodes := πg.InternStr("nodes")
		ßnone := πg.InternStr("none")
		ßnot_utf8_error := πg.InternStr("not_utf8_error")
		ßoff := πg.InternStr("off")
		ßold_settings := πg.InternStr("old_settings")
		ßold_warning := πg.InternStr("old_warning")
		ßon := πg.InternStr("on")
		ßopen := πg.InternStr("open")
		ßoption_groups := πg.InternStr("option_groups")
		ßoption_list := πg.InternStr("option_list")
		ßoptions := πg.InternStr("options")
		ßoptionxform := πg.InternStr("optionxform")
		ßoptparse := πg.InternStr("optparse")
		ßos := πg.InternStr("os")
		ßoverrides := πg.InternStr("overrides")
		ßpath := πg.InternStr("path")
		ßpathsep := πg.InternStr("pathsep")
		ßpep_stylesheet := πg.InternStr("pep_stylesheet")
		ßpep_stylesheet_path := πg.InternStr("pep_stylesheet_path")
		ßpep_template := πg.InternStr("pep_template")
		ßplatform := πg.InternStr("platform")
		ßpop := πg.InternStr("pop")
		ßpopulate_from_components := πg.InternStr("populate_from_components")
		ßprocess := πg.InternStr("process")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßreStructuredText := πg.InternStr("reStructuredText")
		ßread := πg.InternStr("read")
		ßread_config_file := πg.InternStr("read_config_file")
		ßread_file := πg.InternStr("read_file")
		ßreadfp := πg.InternStr("readfp")
		ßrecord_dependencies := πg.InternStr("record_dependencies")
		ßrelative_path_settings := πg.InternStr("relative_path_settings")
		ßremove_section := πg.InternStr("remove_section")
		ßreplace := πg.InternStr("replace")
		ßreport_level := πg.InternStr("report_level")
		ßsections := πg.InternStr("sections")
		ßsectnum_xform := πg.InternStr("sectnum_xform")
		ßset := πg.InternStr("set")
		ßset_defaults_from_dict := πg.InternStr("set_defaults_from_dict")
		ßsetattr := πg.InternStr("setattr")
		ßsettings_default_overrides := πg.InternStr("settings_default_overrides")
		ßsettings_defaults := πg.InternStr("settings_defaults")
		ßsettings_spec := πg.InternStr("settings_spec")
		ßsevere := πg.InternStr("severe")
		ßsource_link := πg.InternStr("source_link")
		ßsource_url := πg.InternStr("source_url")
		ßsplit := πg.InternStr("split")
		ßstandard_config_files := πg.InternStr("standard_config_files")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstore_const := πg.InternStr("store_const")
		ßstore_false := πg.InternStr("store_false")
		ßstore_multiple := πg.InternStr("store_multiple")
		ßstore_true := πg.InternStr("store_true")
		ßstr := πg.InternStr("str")
		ßstrict := πg.InternStr("strict")
		ßstring := πg.InternStr("string")
		ßstrip := πg.InternStr("strip")
		ßstrip_classes := πg.InternStr("strip_classes")
		ßstrip_comments := πg.InternStr("strip_comments")
		ßstrip_elements_with_classes := πg.InternStr("strip_elements_with_classes")
		ßstylesheet := πg.InternStr("stylesheet")
		ßstylesheet_path := πg.InternStr("stylesheet_path")
		ßsys := πg.InternStr("sys")
		ßtemplate := πg.InternStr("template")
		ßthreshold_choices := πg.InternStr("threshold_choices")
		ßthresholds := πg.InternStr("thresholds")
		ßtoc_backlinks := πg.InternStr("toc_backlinks")
		ßtop := πg.InternStr("top")
		ßtraceback := πg.InternStr("traceback")
		ßtrue := πg.InternStr("true")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßunicode := πg.InternStr("unicode")
		ßupdate := πg.InternStr("update")
		ßutils := πg.InternStr("utils")
		ßvalidate_boolean := πg.InternStr("validate_boolean")
		ßvalidate_colon_separated_string_list := πg.InternStr("validate_colon_separated_string_list")
		ßvalidate_comma_separated_list := πg.InternStr("validate_comma_separated_list")
		ßvalidate_dependency_file := πg.InternStr("validate_dependency_file")
		ßvalidate_encoding := πg.InternStr("validate_encoding")
		ßvalidate_encoding_and_error_handler := πg.InternStr("validate_encoding_and_error_handler")
		ßvalidate_encoding_error_handler := πg.InternStr("validate_encoding_error_handler")
		ßvalidate_nonnegative_int := πg.InternStr("validate_nonnegative_int")
		ßvalidate_settings := πg.InternStr("validate_settings")
		ßvalidate_smartquotes_locales := πg.InternStr("validate_smartquotes_locales")
		ßvalidate_strip_class := πg.InternStr("validate_strip_class")
		ßvalidate_ternary := πg.InternStr("validate_ternary")
		ßvalidate_threshold := πg.InternStr("validate_threshold")
		ßvalidate_url_trailing_slash := πg.InternStr("validate_url_trailing_slash")
		ßvalidator := πg.InternStr("validator")
		ßvalues := πg.InternStr("values")
		ßversion := πg.InternStr("version")
		ßversion_info := πg.InternStr("version_info")
		ßversion_template := πg.InternStr("version_template")
		ßwarn_explicit := πg.InternStr("warn_explicit")
		ßwarning := πg.InternStr("warning")
		ßwarning_stream := πg.InternStr("warning_stream")
		ßwarnings := πg.InternStr("warnings")
		ßwrite := πg.InternStr("write")
		ßyes := πg.InternStr("yes")
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
		var πTemp006 []πg.Param
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
		var πTemp024 *πg.Dict
		_ = πTemp024
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 *πg.Object
		_ = πTemp026
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nCommand-line and common processing for Docutils front-end tools.\n\nExports the following classes:\n\n* `OptionParser`: Standard Docutils command-line processing.\n* `Option`: Customized version of `optparse.Option`; validation support.\n* `Values`: Runtime settings; objects are simple structs\n  (``object.attribute``).  Supports cumulative list settings (attributes).\n* `ConfigParser`: Standard Docutils config file processing.\n\nAlso exports the following functions:\n\n* Option callbacks: `store_multiple`, `read_config_file`.\n* Setting validators: `validate_encoding`,\n  `validate_encoding_error_handler`,\n  `validate_encoding_and_error_handler`,\n  `validate_boolean`, `validate_ternary`, `validate_threshold`,\n  `validate_colon_separated_string_list`,\n  `validate_comma_separated_string_list`,\n  `validate_dependency_file`.\n* `make_paths_absolute`.\n* SettingSpec manipulation: `filter_settings_spec`.\n").ToObject()); πE != nil {
				continue
			}
			// line 30: __docformat__ = 'reStructuredText'
			πF.SetLineno(30)
			if πE = πF.Globals().SetItem(πF, ß__docformat__.ToObject(), ßreStructuredText.ToObject()); πE != nil {
				continue
			}
			// line 32: import os
			πF.SetLineno(32)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: import os.path
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "os.path"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: import sys
			πF.SetLineno(34)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: import warnings
			πF.SetLineno(35)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 36: import codecs
			πF.SetLineno(36)
			if πTemp002, πE = πg.ImportModule(πF, "codecs"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcodecs.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: import optparse
			πF.SetLineno(37)
			if πTemp002, πE = πg.ImportModule(πF, "optparse"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßoptparse.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 38: from optparse import SUPPRESS_HELP
			πF.SetLineno(38)
			if πTemp002, πE = πg.ImportModule(πF, "optparse"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSUPPRESS_HELP); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSUPPRESS_HELP.ToObject(), πTemp003); πE != nil {
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
			// line 39: if sys.version_info >= (3, 0):
			πF.SetLineno(39)
		Label1:
			// line 40: from configparser import RawConfigParser
			πF.SetLineno(40)
			if πTemp002, πE = πg.ImportModule(πF, "configparser"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßRawConfigParser); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRawConfigParser.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 41: from os import getcwd
			πF.SetLineno(41)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßgetcwd); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetcwd.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label3
		Label2:
			// line 43: from ConfigParser import RawConfigParser
			πF.SetLineno(43)
			if πTemp002, πE = πg.ImportModule(πF, "ConfigParser"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßRawConfigParser); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRawConfigParser.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 44: from os import getcwdu as getcwd
			πF.SetLineno(44)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßgetcwdu); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßgetcwd.ToObject(), πTemp003); πE != nil {
				continue
			}
			goto Label3
		Label3:
			// line 46: import docutils
			πF.SetLineno(46)
			if πTemp002, πE = πg.ImportModule(πF, "docutils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 47: import docutils.utils
			πF.SetLineno(47)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 48: import docutils.nodes
			πF.SetLineno(48)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.nodes"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdocutils.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 49: from docutils.utils.error_reporting import (locale_encoding, SafeString,
			πF.SetLineno(49)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.error_reporting"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßlocale_encoding); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocale_encoding.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßSafeString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSafeString.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorOutput); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorOutput.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßErrorString); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßErrorString.ToObject(), πTemp003); πE != nil {
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
				goto Label4
			}
			goto Label5
			// line 52: if sys.version_info >= (3, 0):
			πF.SetLineno(52)
		Label4:
			// line 53: unicode = str  # noqa
			πF.SetLineno(53)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunicode.ToObject(), πTemp001); πE != nil {
				continue
			}
			goto Label5
		Label5:
			// line 56: def store_multiple(option, opt, value, parser, *args, **kwargs):
			πF.SetLineno(56)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "option", Def: nil}
			πTemp006[1] = πg.Param{Name: "opt", Def: nil}
			πTemp006[2] = πg.Param{Name: "value", Def: nil}
			πTemp006[3] = πg.Param{Name: "parser", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("store_multiple", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoption *πg.Object = πArgs[0]
				_ = µoption
				var µopt *πg.Object = πArgs[1]
				_ = µopt
				var µvalue *πg.Object = πArgs[2]
				_ = µvalue
				var µparser *πg.Object = πArgs[3]
				_ = µparser
				var µargs *πg.Object = πArgs[4]
				_ = µargs
				var µkwargs *πg.Object = πArgs[5]
				_ = µkwargs
				var µattribute *πg.Object = πg.UnboundLocal
				_ = µattribute
				var µkey *πg.Object = πg.UnboundLocal
				_ = µkey
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
					case 4:
						goto Label4
					case 5:
						goto Label5
					default:
						panic("unexpected function state")
					}
					// line 57: """
					πF.SetLineno(57)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
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
						µattribute = πTemp004
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 64: setattr(parser.values, attribute, None)
					πF.SetLineno(64)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µparser, ßvalues, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µattribute, "attribute"); πE != nil {
						continue
					}
					πTemp005[1] = µattribute
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp005[2] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
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
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µkwargs, ßitems, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp002 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label6
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
							continue
						}
						µkey = πTemp006
						µvalue = πTemp007
					}
					if πE != nil || !πTemp003 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 66: setattr(parser.values, key, value)
					πF.SetLineno(66)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µparser, ßvalues, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp005[1] = µkey
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp005[2] = µvalue
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstore_multiple.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 57: """
			πF.SetLineno(57)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n    Store multiple values in `parser.values`.  (Option callback.)\n\n    Store `None` for each attribute named in `args`, and store the value for\n    each key (attribute name) in `kwargs`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßstore_multiple); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 68: def read_config_file(option, opt, value, parser):
			πF.SetLineno(68)
			πTemp006 = make([]πg.Param, 4)
			πTemp006[0] = πg.Param{Name: "option", Def: nil}
			πTemp006[1] = πg.Param{Name: "opt", Def: nil}
			πTemp006[2] = πg.Param{Name: "value", Def: nil}
			πTemp006[3] = πg.Param{Name: "parser", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("read_config_file", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoption *πg.Object = πArgs[0]
				_ = µoption
				var µopt *πg.Object = πArgs[1]
				_ = µopt
				var µvalue *πg.Object = πArgs[2]
				_ = µvalue
				var µparser *πg.Object = πArgs[3]
				_ = µparser
				var µnew_settings *πg.Object = πg.UnboundLocal
				_ = µnew_settings
				var µerror *πg.Object = πg.UnboundLocal
				_ = µerror
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
					// line 69: """
					πF.SetLineno(69)
					// line 72: try:
					πF.SetLineno(72)
					πF.PushCheckpoint(2)
					// line 73: new_settings = parser.get_config_file_settings(value)
					πF.SetLineno(73)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparser, ßget_config_file_settings, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnew_settings = πTemp003
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
					// line 74: except ValueError as error:
					πF.SetLineno(74)
				Label3:
					µerror = πTemp004.ToObject()
					// line 75: parser.error(error)
					πF.SetLineno(75)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
						continue
					}
					πTemp001[0] = µerror
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparser, ßerror, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 76: parser.values.update(new_settings, parser)
					πF.SetLineno(76)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µnew_settings, "new_settings"); πE != nil {
						continue
					}
					πTemp001[0] = µnew_settings
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					πTemp001[1] = µparser
					if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µparser, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßread_config_file.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 69: """
			πF.SetLineno(69)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Read a configuration file during option processing.  (Option callback.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßread_config_file); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 78: def validate_encoding(setting, value, option_parser,
			πF.SetLineno(78)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp007}
			πTemp004 = πg.NewFunction(πg.NewCode("validate_encoding", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
					// line 80: try:
					πF.SetLineno(80)
					πF.PushCheckpoint(2)
					// line 81: codecs.lookup(value)
					πF.SetLineno(81)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlookup, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
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
					// line 82: except LookupError:
					πF.SetLineno(82)
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µsetting, µvalue).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("setting \"%s\": unknown encoding: \"%s\"").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 83: raise LookupError('setting "%s": unknown encoding: "%s"'
					πF.SetLineno(83)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 85: return value
					πF.SetLineno(85)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_encoding.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 87: def validate_encoding_error_handler(setting, value, option_parser,
			πF.SetLineno(87)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp008}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("validate_encoding_error_handler", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
					// line 89: try:
					πF.SetLineno(89)
					πF.PushCheckpoint(2)
					// line 90: codecs.lookup_error(value)
					πF.SetLineno(90)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlookup_error, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
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
					// line 91: except LookupError:
					πF.SetLineno(91)
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown encoding error handler: \"%s\" (choices: \"strict\", \"ignore\", \"replace\", \"backslashreplace\", \"xmlcharrefreplace\", and possibly others; see documentation for the Python ``codecs`` module)").ToObject(), µvalue); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 92: raise LookupError(
					πF.SetLineno(92)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 97: return value
					πF.SetLineno(97)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_encoding_error_handler.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 99: def validate_encoding_and_error_handler(
			πF.SetLineno(99)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("validate_encoding_and_error_handler", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
				var µencoding *πg.Object = πg.UnboundLocal
				_ = µencoding
				var µhandler *πg.Object = πg.UnboundLocal
				_ = µhandler
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
					// line 101: """
					πF.SetLineno(101)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µvalue, πg.NewStr(":").ToObject()); πE != nil {
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
					// line 105: if ':' in value:
					πF.SetLineno(105)
				Label1:
					// line 106: encoding, handler = value.split(':')
					πF.SetLineno(106)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
						continue
					}
					µencoding = πTemp001
					µhandler = πTemp005
					// line 107: validate_encoding_error_handler(
					πF.SetLineno(107)
					πTemp003 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µsetting, ß_error_handler.ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
						continue
					}
					πTemp003[1] = µhandler
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					πTemp003[2] = µoption_parser
					if πE = πg.CheckLocal(πF, µconfig_parser, "config_parser"); πE != nil {
						continue
					}
					πTemp003[3] = µconfig_parser
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp003[4] = µconfig_section
					if πTemp001, πE = πg.ResolveGlobal(πF, ßvalidate_encoding_error_handler); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µconfig_parser, "config_parser"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µconfig_parser); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 110: if config_parser:
					πF.SetLineno(110)
				Label4:
					// line 111: config_parser.set(config_section, setting + '_error_handler',
					πF.SetLineno(111)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp003[0] = µconfig_section
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µsetting, ß_error_handler.ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
						continue
					}
					πTemp003[2] = µhandler
					if πE = πg.CheckLocal(πF, µconfig_parser, "config_parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µconfig_parser, ßset, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label6
				Label5:
					// line 114: setattr(option_parser.values, setting + '_error_handler', handler)
					πF.SetLineno(114)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoption_parser, ßvalues, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µsetting, ß_error_handler.ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
						continue
					}
					πTemp003[2] = µhandler
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label6
				Label6:
					goto Label3
				Label2:
					// line 116: encoding = value
					πF.SetLineno(116)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					µencoding = µvalue
					goto Label3
				Label3:
					// line 117: validate_encoding(setting, encoding, option_parser,
					πF.SetLineno(117)
					πTemp003 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					πTemp003[0] = µsetting
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp003[1] = µencoding
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					πTemp003[2] = µoption_parser
					if πE = πg.CheckLocal(πF, µconfig_parser, "config_parser"); πE != nil {
						continue
					}
					πTemp003[3] = µconfig_parser
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp003[4] = µconfig_section
					if πTemp001, πE = πg.ResolveGlobal(πF, ßvalidate_encoding); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 119: return encoding
					πF.SetLineno(119)
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πR = µencoding
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_encoding_and_error_handler.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 101: """
			πF.SetLineno(101)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Side-effect: if an error handler is included in the value, it is inserted\n    into the appropriate place as if it was a separate setting/option.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßvalidate_encoding_and_error_handler); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 121: def validate_boolean(setting, value, option_parser,
			πF.SetLineno(121)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp010}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("validate_boolean", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
					case 4:
						goto Label4
					default:
						panic("unexpected function state")
					}
					// line 123: """Check/normalize boolean settings:
					πF.SetLineno(123)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
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
					// line 127: if isinstance(value, bool):
					πF.SetLineno(127)
				Label1:
					// line 128: return value
					πF.SetLineno(128)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
					goto Label2
				Label2:
					// line 129: try:
					πF.SetLineno(129)
					πF.PushCheckpoint(4)
					// line 130: return option_parser.booleans[value.strip().lower()]
					πF.SetLineno(130)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µvalue, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßlower, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp005
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µoption_parser, ßbooleans, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label3
				Label4:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 131: except KeyError:
					πF.SetLineno(131)
				Label5:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown boolean value: \"%s\"").ToObject(), µvalue); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 132: raise LookupError('unknown boolean value: "%s"' % value)
					πF.SetLineno(132)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßvalidate_boolean.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 123: """Check/normalize boolean settings:
			πF.SetLineno(123)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Check/normalize boolean settings:\n         True:  '1', 'on', 'yes', 'true'\n         False: '0', 'off', 'no','false', ''\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßvalidate_boolean); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 134: def validate_ternary(setting, value, option_parser,
			πF.SetLineno(134)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp011}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp011}
			πTemp010 = πg.NewFunction(πg.NewCode("validate_ternary", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
					// line 136: """Check/normalize three-value settings:
					πF.SetLineno(136)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp003[0] = µvalue
					if πTemp004, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µvalue == πTemp005).ToObject()
					πTemp001 = πTemp004
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 141: if isinstance(value, bool) or value is None:
					πF.SetLineno(141)
				Label2:
					// line 142: return value
					πF.SetLineno(142)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
					goto Label3
				Label3:
					// line 143: try:
					πF.SetLineno(143)
					πF.PushCheckpoint(5)
					// line 144: return option_parser.booleans[value.strip().lower()]
					πF.SetLineno(144)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µvalue, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßlower, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp005
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µoption_parser, ßbooleans, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 145: except KeyError:
					πF.SetLineno(145)
				Label6:
					// line 146: return value
					πF.SetLineno(146)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
					πF.RestoreExc(nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßvalidate_ternary.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 136: """Check/normalize three-value settings:
			πF.SetLineno(136)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Check/normalize three-value settings:\n         True:  '1', 'on', 'yes', 'true'\n         False: '0', 'off', 'no','false', ''\n         any other value: returned as-is.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßvalidate_ternary); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 148: def validate_nonnegative_int(setting, value, option_parser,
			πF.SetLineno(148)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp012}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("validate_nonnegative_int", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
					// line 150: value = int(value)
					πF.SetLineno(150)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp003
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µvalue, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 151: if value < 0:
					πF.SetLineno(151)
				Label1:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("negative value; must be positive or zero").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 152: raise ValueError('negative value; must be positive or zero')
					πF.SetLineno(152)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 153: return value
					πF.SetLineno(153)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_nonnegative_int.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 155: def validate_threshold(setting, value, option_parser,
			πF.SetLineno(155)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp013}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("validate_threshold", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
				var πTemp008 *πg.BaseException
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
					case 5:
						goto Label5
					default:
						panic("unexpected function state")
					}
					// line 157: try:
					πF.SetLineno(157)
					πF.PushCheckpoint(2)
					// line 158: return int(value)
					πF.SetLineno(158)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
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
					// line 159: except ValueError:
					πF.SetLineno(159)
				Label3:
					// line 160: try:
					πF.SetLineno(160)
					πF.PushCheckpoint(5)
					// line 161: return option_parser.thresholds[value.lower()]
					πF.SetLineno(161)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µvalue, ßlower, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp007
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µoption_parser, ßthresholds, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp008, πTemp005 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					if πTemp006, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp005.ToObject())
					continue
					// line 162: except (KeyError, AttributeError):
					πF.SetLineno(162)
				Label6:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown threshold: %r.").ToObject(), µvalue); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 163: raise LookupError('unknown threshold: %r.' % value)
					πF.SetLineno(163)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
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
			if πE = πF.Globals().SetItem(πF, ßvalidate_threshold.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 165: def validate_colon_separated_string_list(
			πF.SetLineno(165)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp014}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp014}
			πTemp013 = πg.NewFunction(πg.NewCode("validate_colon_separated_string_list", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
				var µlast *πg.Object = πg.UnboundLocal
				_ = µlast
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
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp002[0] = µvalue
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
					// line 167: if not isinstance(value, list):
					πF.SetLineno(167)
				Label1:
					// line 168: value = value.split(':')
					πF.SetLineno(168)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µvalue = πTemp003
					goto Label3
				Label2:
					// line 170: last = value.pop()
					πF.SetLineno(170)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlast = πTemp003
					// line 171: value.extend(last.split(':'))
					πF.SetLineno(171)
					πTemp002 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µlast, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßextend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label3
				Label3:
					// line 172: return value
					πF.SetLineno(172)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_colon_separated_string_list.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 174: def validate_comma_separated_list(setting, value, option_parser,
			πF.SetLineno(174)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp015}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("validate_comma_separated_list", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
				var µlast *πg.Object = πg.UnboundLocal
				_ = µlast
				var µitems *πg.Object = πg.UnboundLocal
				_ = µitems
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
				var πTemp006 []πg.Param
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
					// line 176: """Check/normalize list arguments (split at "," and strip whitespace).
					πF.SetLineno(176)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp002[0] = µvalue
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
					// line 180: if not isinstance(value, list):
					πF.SetLineno(180)
				Label1:
					// line 181: value = [value]
					πF.SetLineno(181)
					πTemp002 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp002[0] = µvalue
					πTemp001 = πg.NewList(πTemp002...).ToObject()
					µvalue = πTemp001
					goto Label2
				Label2:
					// line 184: last = value.pop()
					πF.SetLineno(184)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßpop, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µlast = πTemp003
					// line 185: items = [i.strip(u' \t\n') for i in last.split(u',') if i.strip(u' \t\n')]
					πF.SetLineno(185)
					πTemp006 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewUnicode(",").ToObject()
								if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µlast, ßsplit, nil); πE != nil {
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
									µi = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewUnicode(" \t\n").ToObject()
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µi, ßstrip, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 185: items = [i.strip(u' \t\n') for i in last.split(u',') if i.strip(u' \t\n')]
								πF.SetLineno(185)
							Label4:
								// line 185: items = [i.strip(u' \t\n') for i in last.split(u',') if i.strip(u' \t\n')]
								πF.SetLineno(185)
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewUnicode(" \t\n").ToObject()
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µi, ßstrip, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								πF.PushCheckpoint(6)
								return πTemp004, nil
							Label6:
								πTemp003 = πSent
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
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					µitems = πTemp001
					// line 186: value.extend(items)
					πF.SetLineno(186)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp002[0] = µitems
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßextend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 187: return value
					πF.SetLineno(187)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_comma_separated_list.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 176: """Check/normalize list arguments (split at "," and strip whitespace).
			πF.SetLineno(176)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Check/normalize list arguments (split at \",\" and strip whitespace).\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßvalidate_comma_separated_list); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 189: def validate_url_trailing_slash(
			πF.SetLineno(189)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp016}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp016}
			πTemp015 = πg.NewFunction(πg.NewCode("validate_url_trailing_slash", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µvalue); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("/").ToObject()
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalue, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 191: if not value:
					πF.SetLineno(191)
				Label1:
					// line 192: return './'
					πF.SetLineno(192)
					πR = πg.NewStr("./").ToObject()
					continue
					goto Label4
					// line 193: elif value.endswith('/'):
					πF.SetLineno(193)
				Label2:
					// line 194: return value
					πF.SetLineno(194)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
					goto Label4
				Label3:
					// line 196: return value + '/'
					πF.SetLineno(196)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µvalue, πg.NewStr("/").ToObject()); πE != nil {
						continue
					}
					πR = πTemp001
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
			if πE = πF.Globals().SetItem(πF, ßvalidate_url_trailing_slash.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 198: def validate_dependency_file(setting, value, option_parser,
			πF.SetLineno(198)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp017}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp017}
			πTemp016 = πg.NewFunction(πg.NewCode("validate_dependency_file", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
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
					// line 200: try:
					πF.SetLineno(200)
					πF.PushCheckpoint(2)
					// line 201: return docutils.utils.DependencyList(value)
					πF.SetLineno(201)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßutils, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßDependencyList, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
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
					// line 202: except IOError:
					πF.SetLineno(202)
				Label3:
					// line 203: return docutils.utils.DependencyList(None)
					πF.SetLineno(203)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßutils, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßDependencyList, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
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
			if πE = πF.Globals().SetItem(πF, ßvalidate_dependency_file.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 205: def validate_strip_class(setting, value, option_parser,
			πF.SetLineno(205)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp018}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp018}
			πTemp017 = πg.NewFunction(πg.NewCode("validate_strip_class", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
				var µcls *πg.Object = πg.UnboundLocal
				_ = µcls
				var µnormalized *πg.Object = πg.UnboundLocal
				_ = µnormalized
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
					// line 208: value = validate_comma_separated_list(setting, value, option_parser,
					πF.SetLineno(208)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					πTemp001[0] = µsetting
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[1] = µvalue
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					πTemp001[2] = µoption_parser
					if πE = πg.CheckLocal(πF, µconfig_parser, "config_parser"); πE != nil {
						continue
					}
					πTemp001[3] = µconfig_parser
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp001[4] = µconfig_section
					if πTemp002, πE = πg.ResolveGlobal(πF, ßvalidate_comma_separated_list); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp003
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µvalue); πE != nil {
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
						µcls = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 212: normalized = docutils.nodes.make_id(cls)
					πF.SetLineno(212)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					πTemp001[0] = µcls
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßnodes, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßmake_id, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnormalized = πTemp006
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnormalized, "normalized"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µcls, µnormalized); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 213: if cls != normalized:
					πF.SetLineno(213)
				Label4:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnormalized, "normalized"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µcls, µnormalized).ToObject()
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("Invalid class value %r (perhaps %r?)").ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 214: raise ValueError('Invalid class value %r (perhaps %r?)'
					πF.SetLineno(214)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 216: return value
					πF.SetLineno(216)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_strip_class.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 218: def validate_smartquotes_locales(setting, value, option_parser,
			πF.SetLineno(218)
			πTemp006 = make([]πg.Param, 5)
			πTemp006[0] = πg.Param{Name: "setting", Def: nil}
			πTemp006[1] = πg.Param{Name: "value", Def: nil}
			πTemp006[2] = πg.Param{Name: "option_parser", Def: nil}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[3] = πg.Param{Name: "config_parser", Def: πTemp019}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[4] = πg.Param{Name: "config_section", Def: πTemp019}
			πTemp018 = πg.NewFunction(πg.NewCode("validate_smartquotes_locales", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsetting *πg.Object = πArgs[0]
				_ = µsetting
				var µvalue *πg.Object = πArgs[1]
				_ = µvalue
				var µoption_parser *πg.Object = πArgs[2]
				_ = µoption_parser
				var µconfig_parser *πg.Object = πArgs[3]
				_ = µconfig_parser
				var µconfig_section *πg.Object = πArgs[4]
				_ = µconfig_section
				var µlc_quotes *πg.Object = πg.UnboundLocal
				_ = µlc_quotes
				var µitem *πg.Object = πg.UnboundLocal
				_ = µitem
				var µlang *πg.Object = πg.UnboundLocal
				_ = µlang
				var µquotes *πg.Object = πg.UnboundLocal
				_ = µquotes
				var µmultichar_quotes *πg.Object = πg.UnboundLocal
				_ = µmultichar_quotes
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
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 []*πg.Object
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
					case 5:
						goto Label5
					default:
						panic("unexpected function state")
					}
					// line 220: """Check/normalize a comma separated list of smart quote definitions.
					πF.SetLineno(220)
					// line 225: value = validate_comma_separated_list(setting, value, option_parser,
					πF.SetLineno(225)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
						continue
					}
					πTemp001[0] = µsetting
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001[1] = µvalue
					if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
						continue
					}
					πTemp001[2] = µoption_parser
					if πE = πg.CheckLocal(πF, µconfig_parser, "config_parser"); πE != nil {
						continue
					}
					πTemp001[3] = µconfig_parser
					if πE = πg.CheckLocal(πF, µconfig_section, "config_section"); πE != nil {
						continue
					}
					πTemp001[4] = µconfig_section
					if πTemp002, πE = πg.ResolveGlobal(πF, ßvalidate_comma_separated_list); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µvalue = πTemp003
					// line 228: lc_quotes = []
					πF.SetLineno(228)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µlc_quotes = πTemp002
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µvalue); πE != nil {
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
						µitem = πTemp003
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 230: try:
					πF.SetLineno(230)
					πF.PushCheckpoint(5)
					// line 231: lang, quotes = item.split(':', 1)
					πF.SetLineno(231)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(":").ToObject()
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µitem, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp006); πE != nil {
						continue
					}
					µlang = πTemp003
					µquotes = πTemp007
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label7
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 232: except AttributeError:
					πF.SetLineno(232)
				Label6:
					// line 235: lc_quotes.append(item)
					πF.SetLineno(235)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µlc_quotes, "lc_quotes"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µlc_quotes, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 236: continue
					πF.SetLineno(236)
					continue
					πF.RestoreExc(nil, nil)
					goto Label4
					// line 237: except ValueError:
					πF.SetLineno(237)
				Label7:
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(2)
					πTemp010[0] = ßascii.ToObject()
					πTemp010[1] = ßbackslashreplace.ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µitem, ßencode, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp003, πE = πg.Mod(πF, πg.NewUnicode("Invalid value \"%s\". Format is \"<language>:<quotes>\".").ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 238: raise ValueError(u'Invalid value "%s".'
					πF.SetLineno(238)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 242: quotes = quotes.strip()
					πF.SetLineno(242)
					if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µquotes, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					µquotes = πTemp006
					// line 243: multichar_quotes = quotes.split(':')
					πF.SetLineno(243)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µquotes, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmultichar_quotes = πTemp006
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmultichar_quotes, "multichar_quotes"); πE != nil {
						continue
					}
					πTemp001[0] = µmultichar_quotes
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
						continue
					}
					πTemp001[0] = µquotes
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.NE(πF, πTemp007, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label9
					}
					goto Label10
					// line 244: if len(multichar_quotes) == 4:
					πF.SetLineno(244)
				Label8:
					// line 245: quotes = multichar_quotes
					πF.SetLineno(245)
					if πE = πg.CheckLocal(πF, µmultichar_quotes, "multichar_quotes"); πE != nil {
						continue
					}
					µquotes = µmultichar_quotes
					goto Label10
					// line 246: elif len(quotes) != 4:
					πF.SetLineno(246)
				Label9:
					πTemp001 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(2)
					πTemp010[0] = ßascii.ToObject()
					πTemp010[1] = ßbackslashreplace.ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µitem, ßencode, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("Invalid value \"%s\". Please specify 4 quotes\n    (primary open/close; secondary open/close).").ToObject(), πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 247: raise ValueError('Invalid value "%s". Please specify 4 quotes\n'
					πF.SetLineno(247)
					πE = πF.Raise(πTemp006, nil, nil)
					continue
					goto Label10
				Label10:
					// line 250: lc_quotes.append((lang, quotes))
					πF.SetLineno(250)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlang, "lang"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µlang, µquotes).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µlc_quotes, "lc_quotes"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µlc_quotes, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 251: return lc_quotes
					πF.SetLineno(251)
					if πE = πg.CheckLocal(πF, µlc_quotes, "lc_quotes"); πE != nil {
						continue
					}
					πR = µlc_quotes
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßvalidate_smartquotes_locales.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 220: """Check/normalize a comma separated list of smart quote definitions.
			πF.SetLineno(220)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Check/normalize a comma separated list of smart quote definitions.\n\n    Return a list of (language-tag, quotes) string tuples.").ToObject()); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßvalidate_smartquotes_locales); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
				continue
			}
			// line 253: def make_paths_absolute(pathdict, keys, base_path=None):
			πF.SetLineno(253)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "pathdict", Def: nil}
			πTemp006[1] = πg.Param{Name: "keys", Def: nil}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp006[2] = πg.Param{Name: "base_path", Def: πTemp020}
			πTemp019 = πg.NewFunction(πg.NewCode("make_paths_absolute", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpathdict *πg.Object = πArgs[0]
				_ = µpathdict
				var µkeys *πg.Object = πArgs[1]
				_ = µkeys
				var µbase_path *πg.Object = πArgs[2]
				_ = µbase_path
				var µkey *πg.Object = πg.UnboundLocal
				_ = µkey
				var µvalue *πg.Object = πg.UnboundLocal
				_ = µvalue
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []πg.Param
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
					case 3:
						goto Label3
					case 4:
						goto Label4
					default:
						panic("unexpected function state")
					}
					// line 254: """
					πF.SetLineno(254)
					if πE = πg.CheckLocal(πF, µbase_path, "base_path"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µbase_path == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 260: if base_path is None:
					πF.SetLineno(260)
				Label1:
					// line 261: base_path = getcwd() # type(base_path) == unicode
					πF.SetLineno(261)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßgetcwd); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µbase_path = πTemp002
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µkeys, "keys"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µkeys); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp003 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label5
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
						µkey = πTemp002
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(3)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpathdict, "pathdict"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µpathdict, µkey); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 264: if key in pathdict:
					πF.SetLineno(264)
				Label6:
					// line 265: value = pathdict[key]
					πF.SetLineno(265)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp002 = µkey
					if πE = πg.CheckLocal(πF, µpathdict, "pathdict"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µpathdict, πTemp002); πE != nil {
						continue
					}
					µvalue = πTemp005
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp006[0] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					πTemp006[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µvalue); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 266: if isinstance(value, list):
					πF.SetLineno(266)
				Label8:
					// line 267: value = [make_one_path_absolute(base_path, path)
					πF.SetLineno(267)
					πTemp007 = make([]πg.Param, 0)
					πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µpath *πg.Object = πg.UnboundLocal
						_ = µpath
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
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µvalue); πE != nil {
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
									µpath = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 267: value = [make_one_path_absolute(base_path, path)
								πF.SetLineno(267)
								πTemp005 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µbase_path, "base_path"); πE != nil {
									continue
								}
								πTemp005[0] = µbase_path
								if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
									continue
								}
								πTemp005[1] = µpath
								if πTemp004, πE = πg.ResolveGlobal(πF, ßmake_one_path_absolute); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								πF.PushCheckpoint(4)
								return πTemp006, nil
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
					if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
						continue
					}
					µvalue = πTemp002
					goto Label10
					// line 269: elif value:
					πF.SetLineno(269)
				Label9:
					// line 270: value = make_one_path_absolute(base_path, value)
					πF.SetLineno(270)
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µbase_path, "base_path"); πE != nil {
						continue
					}
					πTemp006[0] = µbase_path
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp006[1] = µvalue
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmake_one_path_absolute); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µvalue = πTemp008
					goto Label10
				Label10:
					// line 271: pathdict[key] = value
					πF.SetLineno(271)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpathdict, "pathdict"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp008 = µkey
					if πE = πg.SetItem(πF, µpathdict, πTemp008, πTemp002); πE != nil {
						continue
					}
					goto Label7
				Label7:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmake_paths_absolute.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 254: """
			πF.SetLineno(254)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("\n    Interpret filesystem path settings relative to the `base_path` given.\n\n    Paths are values in `pathdict` whose keys are in `keys`.  Get `keys` from\n    `OptionParser.relative_path_settings`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßmake_paths_absolute); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
				continue
			}
			// line 273: def make_one_path_absolute(base_path, path):
			πF.SetLineno(273)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "base_path", Def: nil}
			πTemp006[1] = πg.Param{Name: "path", Def: nil}
			πTemp020 = πg.NewFunction(πg.NewCode("make_one_path_absolute", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µbase_path *πg.Object = πArgs[0]
				_ = µbase_path
				var µpath *πg.Object = πArgs[1]
				_ = µpath
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
					// line 274: return os.path.abspath(os.path.join(base_path, path))
					πF.SetLineno(274)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µbase_path, "base_path"); πE != nil {
						continue
					}
					πTemp002[0] = µbase_path
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					πTemp002[1] = µpath
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpath, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßabspath, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmake_one_path_absolute.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 276: def filter_settings_spec(settings_spec, *exclude, **replace):
			πF.SetLineno(276)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "settings_spec", Def: nil}
			πTemp021 = πg.NewFunction(πg.NewCode("filter_settings_spec", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp006, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsettings_spec *πg.Object = πArgs[0]
				_ = µsettings_spec
				var µexclude *πg.Object = πArgs[1]
				_ = µexclude
				var µreplace *πg.Object = πArgs[2]
				_ = µreplace
				var µsettings *πg.Object = πg.UnboundLocal
				_ = µsettings
				var µi *πg.Object = πg.UnboundLocal
				_ = µi
				var µnewopts *πg.Object = πg.UnboundLocal
				_ = µnewopts
				var µopt_spec *πg.Object = πg.UnboundLocal
				_ = µopt_spec
				var µopt_name *πg.Object = πg.UnboundLocal
				_ = µopt_name
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
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []πg.Param
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
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
					case 4:
						goto Label4
					case 5:
						goto Label5
					default:
						panic("unexpected function state")
					}
					// line 277: """Return a copy of `settings_spec` excluding/replacing some settings.
					πF.SetLineno(277)
					// line 286: settings = list(settings_spec)
					πF.SetLineno(286)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
						continue
					}
					πTemp001[0] = µsettings_spec
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsettings = πTemp003
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewInt(2).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp004[0] = µsettings
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp001[1] = πTemp005
					πTemp001[2] = πg.NewInt(3).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µi = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)
					// line 289: newopts = []
					πF.SetLineno(289)
					πTemp001 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					µnewopts = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µsettings, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp007 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µopt_spec = πTemp005
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 292: opt_name = [opt_string[2:].replace('-', '_')
					πF.SetLineno(292)
					πTemp005 = πg.NewInt(0).ToObject()
					πTemp012 = make([]πg.Param, 0)
					πTemp011 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µopt_string *πg.Object = πg.UnboundLocal
						_ = µopt_string
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
								πTemp002 = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µopt_spec, "opt_spec"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µopt_spec, πTemp002); πE != nil {
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
									µopt_string = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)
								πTemp006 = πF.MakeArgs(1)
								πTemp006[0] = πg.NewStr("--").ToObject()
								if πE = πg.CheckLocal(πF, µopt_string, "opt_string"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µopt_string, ßstartswith, nil); πE != nil {
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
								// line 292: opt_name = [opt_string[2:].replace('-', '_')
								πF.SetLineno(292)
							Label4:
								// line 292: opt_name = [opt_string[2:].replace('-', '_')
								πF.SetLineno(292)
								πTemp006 = πF.MakeArgs(2)
								πTemp006[0] = πg.NewStr("-").ToObject()
								πTemp006[1] = ß_.ToObject()
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µopt_string, "opt_string"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetItem(πF, µopt_string, πTemp002); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								πF.PushCheckpoint(6)
								return πTemp003, nil
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
					if πTemp013, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ListType.Call(πF, πg.Args{πTemp013}, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
						continue
					}
					µopt_name = πTemp008
					if πE = πg.CheckLocal(πF, µopt_name, "opt_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µexclude, "exclude"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, µexclude, µopt_name); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label7
					}
					goto Label8
					// line 296: if opt_name in exclude:
					πF.SetLineno(296)
				Label7:
					// line 297: continue
					πF.SetLineno(297)
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µopt_name, "opt_name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, µreplace, ßkeys, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp010, µopt_name); πE != nil {
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
					// line 298: if opt_name in replace.keys():
					πF.SetLineno(298)
				Label9:
					// line 299: newopts.append(replace[opt_name])
					πF.SetLineno(299)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µopt_name, "opt_name"); πE != nil {
						continue
					}
					πTemp005 = µopt_name
					if πE = πg.CheckLocal(πF, µreplace, "replace"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µreplace, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πE = πg.CheckLocal(πF, µnewopts, "newopts"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µnewopts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label11
				Label10:
					// line 301: newopts.append(opt_spec)
					πF.SetLineno(301)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µopt_spec, "opt_spec"); πE != nil {
						continue
					}
					πTemp001[0] = µopt_spec
					if πE = πg.CheckLocal(πF, µnewopts, "newopts"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µnewopts, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label11
				Label11:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 302: settings[i] = tuple(newopts)
					πF.SetLineno(302)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnewopts, "newopts"); πE != nil {
						continue
					}
					πTemp001[0] = µnewopts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp008 = µi
					if πE = πg.SetItem(πF, µsettings, πTemp008, πTemp003); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 303: return tuple(settings)
					πF.SetLineno(303)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
						continue
					}
					πTemp001[0] = µsettings
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßfilter_settings_spec.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 277: """Return a copy of `settings_spec` excluding/replacing some settings.
			πF.SetLineno(277)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("Return a copy of `settings_spec` excluding/replacing some settings.\n\n    `settings_spec` is a tuple of configuration settings with a structure\n    described for docutils.SettingsSpec.settings_spec.\n\n    Optional positional arguments are names of to-be-excluded settings.\n    Keyword arguments are option specification replacements.\n    (See the html4strict writer for an example.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßfilter_settings_spec); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
				continue
			}
			// line 306: class Values(optparse.Values):
			πF.SetLineno(306)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp025, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
				continue
			}
			if πTemp026, πE = πg.GetAttr(πF, πTemp025, ßValues, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp026
			πTemp024 = πg.NewDict()
			if πTemp022, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp024.SetItem(πF, ß__module__.ToObject(), πTemp022); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Values", "/usr/lib/python2.7/site-packages/docutils/frontend.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp024
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
					// line 308: """
					πF.SetLineno(308)
					// line 308: """
					πF.SetLineno(308)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Updates list attributes by extension rather than by replacement.\n    Works in conjunction with the `OptionParser.lists` instance attribute.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 313: def __init__(self, *args, **kwargs):
					πF.SetLineno(313)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp002, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 314: optparse.Values.__init__(self, *args, **kwargs)
							πF.SetLineno(314)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßValues, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ßrecord_dependencies.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßrecord_dependencies, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005 == πTemp006).ToObject()
							πTemp002 = πTemp003
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 315: if (not hasattr(self, 'record_dependencies')
							πF.SetLineno(315)
						Label2:
							// line 318: self.record_dependencies = docutils.utils.DependencyList()
							πF.SetLineno(318)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßutils, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßDependencyList, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrecord_dependencies, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 320: def update(self, other_dict, option_parser):
					πF.SetLineno(320)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other_dict", Def: nil}
					πTemp002[2] = πg.Param{Name: "option_parser", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µother_dict *πg.Object = πArgs[1]
						_ = µother_dict
						var µoption_parser *πg.Object = πArgs[2]
						_ = µoption_parser
						var µsetting *πg.Object = πg.UnboundLocal
						_ = µsetting
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
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							πTemp001[0] = µother_dict
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
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
							// line 321: if isinstance(other_dict, Values):
							πF.SetLineno(321)
						Label1:
							// line 322: other_dict = other_dict.__dict__
							πF.SetLineno(322)
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother_dict, ß__dict__, nil); πE != nil {
								continue
							}
							µother_dict = πTemp002
							goto Label2
						Label2:
							// line 323: other_dict = other_dict.copy()
							πF.SetLineno(323)
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother_dict, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µother_dict = πTemp003
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption_parser, ßlists, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp004 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label5
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
								µsetting = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[1] = µsetting
							if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πTemp007
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, µother_dict, µsetting); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp008).ToObject()
							πTemp003 = πTemp005
						Label6:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 325: if (hasattr(self, setting) and setting in other_dict):
							πF.SetLineno(325)
						Label7:
							// line 326: value = getattr(self, setting)
							πF.SetLineno(326)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[1] = µsetting
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp005
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µvalue); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 327: if value:
							πF.SetLineno(327)
						Label9:
							// line 328: value += other_dict[setting]
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp003 = µsetting
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µother_dict, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µvalue, πTemp005); πE != nil {
								continue
							}
							µvalue = πTemp003
							// line 329: del other_dict[setting]
							πF.SetLineno(329)
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp003 = µsetting
							if πE = πg.DelItem(πF, µother_dict, πTemp003); πE != nil {
								continue
							}
							goto Label10
						Label10:
							goto Label8
						Label8:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 330: self._update_loose(other_dict)
							πF.SetLineno(330)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother_dict, "other_dict"); πE != nil {
								continue
							}
							πTemp001[0] = µother_dict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_update_loose, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 332: def copy(self):
					πF.SetLineno(332)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 333: """Return a shallow copy of `self`."""
							πF.SetLineno(333)
							// line 334: return self.__class__(defaults=self.__dict__)
							πF.SetLineno(334)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"defaults", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 333: """Return a shallow copy of `self`."""
					πF.SetLineno(333)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a shallow copy of `self`.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßcopy); πE != nil {
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
			if πTemp023, πE = πTemp024.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp023 == nil {
				πTemp023 = πg.TypeType.ToObject()
			}
			if πTemp025, πE = πTemp023.Call(πF, []*πg.Object{πg.NewStr("Values").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp024.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßValues.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 337: class Option(optparse.Option):
			πF.SetLineno(337)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp025, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
				continue
			}
			if πTemp026, πE = πg.GetAttr(πF, πTemp025, ßOption, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp026
			πTemp024 = πg.NewDict()
			if πTemp022, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp024.SetItem(πF, ß__module__.ToObject(), πTemp022); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Option", "/usr/lib/python2.7/site-packages/docutils/frontend.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp024
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 []πg.Param
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 339: ATTRS = optparse.Option.ATTRS + ['validator', 'overrides']
					πF.SetLineno(339)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßoptparse); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßOption, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßATTRS, nil); πE != nil {
						continue
					}
					πTemp004 = make([]*πg.Object, 2)
					πTemp004[0] = ßvalidator.ToObject()
					πTemp004[1] = ßoverrides.ToObject()
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßATTRS.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 341: def process(self, opt, value, values, parser):
					πF.SetLineno(341)
					πTemp005 = make([]πg.Param, 5)
					πTemp005[0] = πg.Param{Name: "self", Def: nil}
					πTemp005[1] = πg.Param{Name: "opt", Def: nil}
					πTemp005[2] = πg.Param{Name: "value", Def: nil}
					πTemp005[3] = πg.Param{Name: "values", Def: nil}
					πTemp005[4] = πg.Param{Name: "parser", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("process", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µopt *πg.Object = πArgs[1]
						_ = µopt
						var µvalue *πg.Object = πArgs[2]
						_ = µvalue
						var µvalues *πg.Object = πArgs[3]
						_ = µvalues
						var µparser *πg.Object = πArgs[4]
						_ = µparser
						var µresult *πg.Object = πg.UnboundLocal
						_ = µresult
						var µsetting *πg.Object = πg.UnboundLocal
						_ = µsetting
						var µnew_value *πg.Object = πg.UnboundLocal
						_ = µnew_value
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
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 342: """
							πF.SetLineno(342)
							// line 347: result = optparse.Option.process(self, opt, value, values, parser)
							πF.SetLineno(347)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[1] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[2] = µvalue
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[3] = µvalues
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							πTemp001[4] = µparser
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßOption, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßprocess, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp003
							// line 348: setting = self.dest
							πF.SetLineno(348)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdest, nil); πE != nil {
								continue
							}
							µsetting = πTemp002
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µsetting); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 349: if setting:
							πF.SetLineno(349)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvalidator, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 350: if self.validator:
							πF.SetLineno(350)
						Label3:
							// line 351: value = getattr(values, setting)
							πF.SetLineno(351)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[0] = µvalues
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[1] = µsetting
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp003
							// line 352: try:
							πF.SetLineno(352)
							πF.PushCheckpoint(6)
							// line 353: new_value = self.validator(setting, value, parser)
							πF.SetLineno(353)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[0] = µsetting
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[1] = µvalue
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							πTemp001[2] = µparser
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßvalidator, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnew_value = πTemp003
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 354: except Exception as error:
							πF.SetLineno(354)
						Label7:
							µerror = πTemp005.ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp007[0] = µerror
							if πTemp008, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003 = πg.NewTuple2(µopt, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Error in option \"%s\":\n    %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßOptionValueError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 355: raise optparse.OptionValueError(
							πF.SetLineno(355)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							// line 358: setattr(values, setting, new_value)
							πF.SetLineno(358)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[0] = µvalues
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[1] = µsetting
							if πE = πg.CheckLocal(πF, µnew_value, "new_value"); πE != nil {
								continue
							}
							πTemp001[2] = µnew_value
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoverrides, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 359: if self.overrides:
							πF.SetLineno(359)
						Label8:
							// line 360: setattr(values, self.overrides, None)
							πF.SetLineno(360)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[0] = µvalues
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoverrides, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label9
						Label9:
							goto Label2
						Label2:
							// line 361: return result
							πF.SetLineno(361)
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
					if πE = πClass.SetItem(πF, ßprocess.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 342: """
					πF.SetLineno(342)
					// line 342: """
					πF.SetLineno(342)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n        Call the validator function on applicable settings and\n        evaluate the 'overrides' option.\n        Extends `optparse.Option.process`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("\n        Call the validator function on applicable settings and\n        evaluate the 'overrides' option.\n        Extends `optparse.Option.process`.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßprocess); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp003, ß__doc__, πTemp002); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp023, πE = πTemp024.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp023 == nil {
				πTemp023 = πg.TypeType.ToObject()
			}
			if πTemp025, πE = πTemp023.Call(πF, []*πg.Object{πg.NewStr("Option").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp024.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOption.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 364: class OptionParser(optparse.OptionParser, docutils.SettingsSpec):
			πF.SetLineno(364)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp025, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
				continue
			}
			if πTemp026, πE = πg.GetAttr(πF, πTemp025, ßOptionParser, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp026
			if πTemp025, πE = πg.ResolveGlobal(πF, ßdocutils); πE != nil {
				continue
			}
			if πTemp026, πE = πg.GetAttr(πF, πTemp025, ßSettingsSpec, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp026
			πTemp024 = πg.NewDict()
			if πTemp022, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp024.SetItem(πF, ß__module__.ToObject(), πTemp022); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionParser", "/usr/lib/python2.7/site-packages/docutils/frontend.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp024
				_ = πClass
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πTemp014 []πg.Param
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
					// line 366: """
					πF.SetLineno(366)
					// line 366: """
					πF.SetLineno(366)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Parser for command-line and library use.  The `settings_spec`\n    specification here and in other Docutils components are merged to build\n    the set of command-line options and runtime settings for this process.\n\n    Common settings (defined below) and component-specific settings must not\n    conflict.  Short options are reserved for common settings, and components\n    are restrict to using long options.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 376: standard_config_files = [
					πF.SetLineno(376)
					πTemp001 = make([]*πg.Object, 3)
					πTemp001[0] = πg.NewStr("/etc/docutils.conf").ToObject()
					πTemp001[1] = πg.NewStr("./docutils.conf").ToObject()
					πTemp001[2] = πg.NewStr("~/.docutils").ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßstandard_config_files.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 380: """Docutils configuration files, using ConfigParser syntax.  Filenames
					πF.SetLineno(380)
					// line 383: threshold_choices = 'info 1 warning 2 error 3 severe 4 none 5'.split()
					πF.SetLineno(383)
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("info 1 warning 2 error 3 severe 4 none 5").ToObject(), ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßthreshold_choices.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 384: """Possible inputs for for --report and --halt threshold values."""
					πF.SetLineno(384)
					// line 386: thresholds = {'info': 1, 'warning': 2, 'error': 3, 'severe': 4, 'none': 5}
					πF.SetLineno(386)
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßinfo.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßwarning.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßerror.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßsevere.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßnone.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßthresholds.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 387: """Lookup table for --report and --halt threshold values."""
					πF.SetLineno(387)
					// line 389: booleans={'1': True, 'on': True, 'yes': True, 'true': True,
					πF.SetLineno(389)
					πTemp004 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß1.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßon.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßyes.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßtrue.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß0.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßoff.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßno.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßfalse.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßbooleans.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 391: """Lookup table for boolean configuration file settings."""
					πF.SetLineno(391)
					// line 393: default_error_encoding = getattr(sys.stderr, 'encoding',
					πF.SetLineno(393)
					πTemp001 = πF.MakeArgs(3)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßstderr, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					πTemp001[1] = ßencoding.ToObject()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßgetattr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp006
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßlocale_encoding); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					πTemp002 = ßascii.ToObject()
				Label1:
					if πE = πClass.SetItem(πF, ßdefault_error_encoding.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 396: default_error_encoding_error_handler = 'backslashreplace'
					πF.SetLineno(396)
					if πE = πClass.SetItem(πF, ßdefault_error_encoding_error_handler.ToObject(), ßbackslashreplace.ToObject()); πE != nil {
						continue
					}
					// line 398: settings_spec = (
					πF.SetLineno(398)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 50)
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--title").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify the document title as metadata.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[0] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--generator").ToObject()
					πTemp008[1] = πg.NewStr("-g").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Include a \"Generated by Docutils\" credit and link.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[1] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-generator").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßgenerator.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Do not include a generator credit.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[2] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--date").ToObject()
					πTemp008[1] = πg.NewStr("-d").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), πg.NewStr("%Y-%m-%d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßdatestamp.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Include the date at the end of the document (UTC).").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[3] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--time").ToObject()
					πTemp008[1] = πg.NewStr("-t").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), πg.NewStr("%Y-%m-%d %H:%M UTC").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßdatestamp.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Include the time & date (UTC).").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[4] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-datestamp").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßdatestamp.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Do not include a datestamp of any kind.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[5] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--source-link").ToObject()
					πTemp008[1] = πg.NewStr("-s").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Include a \"View document source\" link.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[6] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--source-url").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<URL>").ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Use <URL> for a source link; implies --source-link.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[7] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-source-link").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßcallback.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßstore_multiple); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßcallback.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(ßsource_link.ToObject(), ßsource_url.ToObject()).ToObject()
					if πE = πTemp004.SetItem(πF, ßcallback_args.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Do not include a \"View document source\" link.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[8] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--toc-entry-backlinks").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßtoc_backlinks.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), ßentry.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), ßentry.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Link from section headers to TOC entries.  (default)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[9] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--toc-top-backlinks").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßtoc_backlinks.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), ßtop.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Link from section headers to the top of the TOC.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[10] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-toc-backlinks").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßtoc_backlinks.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable backlinks to the table of contents.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[11] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--footnote-backlinks").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Link from footnotes/citations to references. (default)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[12] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-footnote-backlinks").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßfootnote_backlinks.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable backlinks from footnotes and citations.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[13] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--section-numbering").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßsectnum_xform.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Enable section numbering by Docutils.  (default)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[14] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-section-numbering").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßsectnum_xform.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable section numbering by Docutils.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[15] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--strip-comments").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Remove comment elements from the document tree.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[16] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--leave-comments").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßstrip_comments.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Leave comment elements in the document tree. (default)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[17] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--strip-elements-with-class").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßappend.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßstrip_elements_with_classes.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<class>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_strip_class); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Remove all elements with classes=\"<class>\" from the document tree. Warning: potentially dangerous; use with caution. (Multiple-use option.)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[18] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--strip-class").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßappend.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßstrip_classes.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<class>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_strip_class); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Remove all classes=\"<class>\" attributes from elements in the document tree. Warning: potentially dangerous; use with caution. (Multiple-use option.)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[19] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--report").ToObject()
					πTemp008[1] = πg.NewStr("-r").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßthreshold_choices); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßchoices.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßreport_level.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_threshold); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Report system messages at or higher than <level>: \"info\" or \"1\", \"warning\"/\"2\" (default), \"error\"/\"3\", \"severe\"/\"4\", \"none\"/\"5\"").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[20] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--verbose").ToObject()
					πTemp008[1] = πg.NewStr("-v").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßreport_level.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Report all system messages.  (Same as \"--report=1\".)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[21] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--quiet").ToObject()
					πTemp008[1] = πg.NewStr("-q").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßreport_level.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Report no system messages.  (Same as \"--report=5\".)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[22] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--halt").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßthreshold_choices); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßchoices.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßhalt_level.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_threshold); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Halt execution at system messages at or above <level>.  Levels as in --report.  Default: 4 (severe).").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[23] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--strict").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_const.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßconst.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßhalt_level.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Halt at the slightest problem.  Same as \"--halt=info\".").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[24] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--exit-status").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßthreshold_choices); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßchoices.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßexit_status_level.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<level>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_threshold); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Enable a non-zero exit status for non-halting system messages at or above <level>.  Default: 5 (disabled).").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[25] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--debug").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Enable debug-level system messages and diagnostics.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[26] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-debug").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßdebug.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable debug output.  (default)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[27] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--warnings").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßwarning_stream.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file>").ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Send the output of system messages to <file>.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[28] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--traceback").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_boolean); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Enable Python tracebacks when Docutils is halted.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[29] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--no-traceback").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßtraceback.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_false.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Disable Python tracebacks.  (default)").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[30] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--input-encoding").ToObject()
					πTemp008[1] = πg.NewStr("-i").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<name[:handler]>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_encoding_and_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify the encoding and optionally the error handler of input text.  Default: <locale-dependent>:strict.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[31] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--input-encoding-error-handler").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), ßstrict.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_encoding_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify the error handler for undecodable characters.  Choices: \"strict\" (default), \"ignore\", and \"replace\".").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[32] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--output-encoding").ToObject()
					πTemp008[1] = πg.NewStr("-o").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<name[:handler]>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πg.NewStr("utf-8").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_encoding_and_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify the text encoding and optionally the error handler for output.  Default: UTF-8:strict.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[33] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--output-encoding-error-handler").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), ßstrict.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_encoding_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify error handler for unencodable output characters; \"strict\" (default), \"ignore\", \"replace\", \"xmlcharrefreplace\", \"backslashreplace\".").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[34] = πTemp007
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_error_encoding); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_error_encoding_error_handler); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(πTemp011, πTemp012).ToObject()
					if πTemp009, πE = πg.Mod(πF, πg.NewStr("Specify text encoding and error handler for error output.  Default: %s:%s.").ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--error-encoding").ToObject()
					πTemp008[1] = πg.NewStr("-e").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<name[:handler]>").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_error_encoding); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πTemp011); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_encoding_and_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[35] = πTemp007
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_error_encoding_error_handler); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Mod(πF, πg.NewStr("Specify the error handler for unencodable characters in error output.  Default: %s.").ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--error-encoding-error-handler").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdefault_error_encoding_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πTemp011); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_encoding_error_handler); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[36] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--language").ToObject()
					πTemp008[1] = πg.NewStr("-l").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßlanguage_code.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), ßen.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<name>").ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Specify the language (as BCP 47 language tag).  Default: en.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[37] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--record-dependencies").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file>").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_dependency_file); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp010); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Write output file dependencies to <file>.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[38] = πTemp007
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--config").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßmetavar.ToObject(), πg.NewStr("<file>").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßtype.ToObject(), ßstring.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßcallback.ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßread_config_file); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßcallback.ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Read configuration settings from <file>, if it exists.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[39] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--version").ToObject()
					πTemp008[1] = πg.NewStr("-V").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßversion.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Show this program's version number and exit.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[40] = πTemp007
					πTemp008 = make([]*πg.Object, 2)
					πTemp008[0] = πg.NewStr("--help").ToObject()
					πTemp008[1] = πg.NewStr("-h").ToObject()
					πTemp009 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßhelp.ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πg.NewStr("Show this help message and exit.").ToObject(), πTemp009, πTemp010).ToObject()
					πTemp001[41] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--id-prefix").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[42] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--auto-id-prefix").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßdefault.ToObject(), ßid.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[43] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--dump-settings").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[44] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--dump-internals").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[45] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--dump-transforms").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[46] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--dump-pseudo-xml").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[47] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--expose-internal-attribute").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßappend.ToObject()); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßdest.ToObject(), ßexpose_internals.ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_colon_separated_string_list); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ßvalidator.ToObject(), πTemp011); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[48] = πTemp007
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßSUPPRESS_HELP); πE != nil {
						continue
					}
					πTemp008 = make([]*πg.Object, 1)
					πTemp008[0] = πg.NewStr("--strict-visitor").ToObject()
					πTemp010 = πg.NewList(πTemp008...).ToObject()
					πTemp004 = πg.NewDict()
					if πE = πTemp004.SetItem(πF, ßaction.ToObject(), ßstore_true.ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp004.ToObject()
					πTemp007 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
					πTemp001[49] = πTemp007
					πTemp006 = πg.NewTuple(πTemp001...).ToObject()
					πTemp002 = πg.NewTuple3(πg.NewStr("General Docutils Options").ToObject(), πTemp003, πTemp006).ToObject()
					if πE = πClass.SetItem(πF, ßsettings_spec.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 563: """Runtime settings and command-line options common to all Docutils front
					πF.SetLineno(563)
					// line 567: settings_defaults = {'_disable_config': None,
					πF.SetLineno(567)
					πTemp004 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß_disable_config.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß_source.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß_destination.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πTemp004.SetItem(πF, ß_config_files.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp004.ToObject()
					if πE = πClass.SetItem(πF, ßsettings_defaults.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 571: """Defaults for settings that don't have command-line option equivalents."""
					πF.SetLineno(571)
					// line 573: relative_path_settings = ('warning_stream',)
					πF.SetLineno(573)
					πTemp002 = πg.NewTuple1(ßwarning_stream.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßrelative_path_settings.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 575: config_section = 'general'
					πF.SetLineno(575)
					if πE = πClass.SetItem(πF, ßconfig_section.ToObject(), ßgeneral.ToObject()); πE != nil {
						continue
					}
					// line 577: version_template = ('%%prog (Docutils %s%s, Python %s, on %s)'
					πF.SetLineno(577)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdocutils); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__version__, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdocutils); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ß__version_details__, nil); πE != nil {
						continue
					}
					πTemp009 = πTemp011
					if πTemp013, πE = πg.IsTrue(πF, πTemp009); πE != nil {
						continue
					}
					if !πTemp013 {
						goto Label3
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdocutils); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ß__version_details__, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Mod(πF, πg.NewStr(" [%s]").ToObject(), πTemp012); πE != nil {
						continue
					}
					πTemp009 = πTemp010
				Label3:
					πTemp006 = πTemp009
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label2
					}
					πTemp006 = ß.ToObject()
				Label2:
					πTemp009 = πg.NewInt(0).ToObject()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßversion, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp012, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp012, πTemp009); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßplatform, nil); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple4(πTemp007, πTemp006, πTemp010, πTemp011).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("%%prog (Docutils %s%s, Python %s, on %s)").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßversion_template.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 582: """Default version message."""
					πF.SetLineno(582)
					// line 584: def __init__(self, components=(), defaults=None, read_config_files=None,
					πF.SetLineno(584)
					πTemp014 = make([]πg.Param, 4)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewTuple0().ToObject()
					πTemp014[1] = πg.Param{Name: "components", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp014[2] = πg.Param{Name: "defaults", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp014[3] = πg.Param{Name: "read_config_files", Def: πTemp003}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcomponents *πg.Object = πArgs[1]
						_ = µcomponents
						var µdefaults *πg.Object = πArgs[2]
						_ = µdefaults
						var µread_config_files *πg.Object = πArgs[3]
						_ = µread_config_files
						var µargs *πg.Object = πArgs[4]
						_ = µargs
						var µkwargs *πg.Object = πArgs[5]
						_ = µkwargs
						var µconfig_settings *πg.Object = πg.UnboundLocal
						_ = µconfig_settings
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 []*πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 586: """
							πF.SetLineno(586)
							// line 592: self.lists = {}
							πF.SetLineno(592)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlists, πTemp003); πE != nil {
								continue
							}
							// line 593: """Set of list-type settings."""
							πF.SetLineno(593)
							// line 595: self.config_files = []
							πF.SetLineno(595)
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßconfig_files, πTemp003); πE != nil {
								continue
							}
							// line 596: """List of paths of applied configuration files."""
							πF.SetLineno(596)
							// line 598: optparse.OptionParser.__init__(
							πF.SetLineno(598)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOption); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"width", πg.NewInt(78).ToObject()},
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßTitledHelpFormatter, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"option_class", πTemp002},
								{"add_help_option", πTemp003},
								{"formatter", πTemp006},
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßOptionParser, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp004, µargs, πTemp005, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßversion, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label1
							}
							goto Label2
							// line 602: if not self.version:
							πF.SetLineno(602)
						Label1:
							// line 603: self.version = self.version_template
							πF.SetLineno(603)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßversion_template, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßversion, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 605: self.relative_path_settings = list(self.relative_path_settings)
							πF.SetLineno(605)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrelative_path_settings, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrelative_path_settings, πTemp002); πE != nil {
								continue
							}
							// line 606: self.components = (self,) + tuple(components)
							πF.SetLineno(606)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(µself).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
								continue
							}
							πTemp004[0] = µcomponents
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcomponents, πTemp003); πE != nil {
								continue
							}
							// line 607: self.populate_from_components(self.components)
							πF.SetLineno(607)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcomponents, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpopulate_from_components, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 608: self.set_defaults_from_dict(defaults or {})
							πF.SetLineno(608)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							πTemp002 = µdefaults
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							πTemp001 = πg.NewDict()
							πTemp003 = πTemp001.ToObject()
							πTemp002 = πTemp003
						Label3:
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_defaults_from_dict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µread_config_files, "read_config_files"); πE != nil {
								continue
							}
							πTemp002 = µread_config_files
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label4
							}
							πTemp006 = ß_disable_config.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp010).ToObject()
							πTemp002 = πTemp003
						Label4:
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							goto Label6
							// line 609: if read_config_files and not self.defaults['_disable_config']:
							πF.SetLineno(609)
						Label5:
							// line 610: try:
							πF.SetLineno(610)
							πF.PushCheckpoint(8)
							// line 611: config_settings = self.get_standard_config_settings()
							πF.SetLineno(611)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_standard_config_settings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µconfig_settings = πTemp003
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp011, πTemp012 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
							continue
							// line 612: except ValueError as error:
							πF.SetLineno(612)
						Label9:
							µerror = πTemp011.ToObject()
							// line 613: self.error(SafeString(error))
							πF.SetLineno(613)
							πTemp004 = πF.MakeArgs(1)
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp013[0] = µerror
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSafeString); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							// line 614: self.set_defaults_from_dict(config_settings.__dict__)
							πF.SetLineno(614)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µconfig_settings, "config_settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µconfig_settings, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_defaults_from_dict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 586: """
					πF.SetLineno(586)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        `components` is a list of Docutils components each containing a\n        ``.settings_spec`` attribute.  `defaults` is a mapping of setting\n        default overrides.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 616: def populate_from_components(self, components):
					πF.SetLineno(616)
					πTemp014 = make([]πg.Param, 2)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp014[1] = πg.Param{Name: "components", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("populate_from_components", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcomponents *πg.Object = πArgs[1]
						_ = µcomponents
						var µcomponent *πg.Object = πg.UnboundLocal
						_ = µcomponent
						var µsettings_spec *πg.Object = πg.UnboundLocal
						_ = µsettings_spec
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µtitle *πg.Object = πg.UnboundLocal
						_ = µtitle
						var µdescription *πg.Object = πg.UnboundLocal
						_ = µdescription
						var µoption_spec *πg.Object = πg.UnboundLocal
						_ = µoption_spec
						var µgroup *πg.Object = πg.UnboundLocal
						_ = µgroup
						var µhelp_text *πg.Object = πg.UnboundLocal
						_ = µhelp_text
						var µoption_strings *πg.Object = πg.UnboundLocal
						_ = µoption_strings
						var µkwargs *πg.Object = πg.UnboundLocal
						_ = µkwargs
						var µoption *πg.Object = πg.UnboundLocal
						_ = µoption
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 πg.KWArgs
						_ = πTemp014
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
							case 12:
								goto Label12
							case 13:
								goto Label13
							case 19:
								goto Label19
							case 20:
								goto Label20
							default:
								panic("unexpected function state")
							}
							// line 617: """
							πF.SetLineno(617)
							if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µcomponents); πE != nil {
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
								µcomponent = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µcomponent == πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 624: if component is None:
							πF.SetLineno(624)
						Label4:
							// line 625: continue
							πF.SetLineno(625)
							continue
							goto Label5
						Label5:
							// line 626: settings_spec = component.settings_spec
							πF.SetLineno(626)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcomponent, ßsettings_spec, nil); πE != nil {
								continue
							}
							µsettings_spec = πTemp004
							// line 627: self.relative_path_settings.extend(
							πF.SetLineno(627)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcomponent, ßrelative_path_settings, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrelative_path_settings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßextend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp006 = πF.MakeArgs(3)
							πTemp006[0] = πg.NewInt(0).ToObject()
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							πTemp007[0] = µsettings_spec
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[1] = πTemp008
							πTemp006[2] = πg.NewInt(3).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp003 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µi = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 630: title, description, option_spec = settings_spec[i:i+3]
							πF.SetLineno(630)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, µi, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{µi, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsettings_spec, "settings_spec"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µsettings_spec, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}, πTemp008); πE != nil {
								continue
							}
							µtitle = πTemp005
							µdescription = πTemp010
							µoption_spec = πTemp011
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µtitle); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							goto Label10
							// line 631: if title:
							πF.SetLineno(631)
						Label9:
							// line 632: group = optparse.OptionGroup(self, title, description)
							πF.SetLineno(632)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[0] = µself
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							πTemp006[1] = µtitle
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp006[2] = µdescription
							if πTemp005, πE = πg.ResolveGlobal(πF, ßoptparse); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßOptionGroup, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µgroup = πTemp005
							// line 633: self.add_option_group(group)
							πF.SetLineno(633)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp006[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßadd_option_group, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label11
						Label10:
							// line 635: group = self        # single options
							πF.SetLineno(635)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							µgroup = µself
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µoption_spec, "option_spec"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Iter(πF, µoption_spec); πE != nil {
								continue
							}
							πF.PushCheckpoint(13)
							πTemp009 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label14
							}
							if πTemp008, πE = πg.Next(πF, πTemp005); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp012 = !isStop
							} else {
								πTemp012 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp013}}}, πTemp008); πE != nil {
									continue
								}
								µhelp_text = πTemp010
								µoption_strings = πTemp011
								µkwargs = πTemp013
							}
							if πE != nil || !πTemp012 {
								continue
							}
							πF.PushCheckpoint(12)
							// line 637: option = group.add_option(help=help_text, *option_strings,
							πF.SetLineno(637)
							if πE = πg.CheckLocal(πF, µoption_strings, "option_strings"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhelp_text, "help_text"); πE != nil {
								continue
							}
							πTemp014 = πg.KWArgs{
								{"help", µhelp_text},
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µgroup, ßadd_option, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Invoke(πF, πTemp008, nil, µoption_strings, πTemp014, µkwargs); πE != nil {
								continue
							}
							µoption = πTemp010
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßaction.ToObject()
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µkwargs, ßget, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp008, πE = πg.Eq(πF, πTemp011, ßappend.ToObject()); πE != nil {
								continue
							}
							if πTemp012, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label15
							}
							goto Label16
							// line 639: if kwargs.get('action') == 'append':
							πF.SetLineno(639)
						Label15:
							// line 640: self.lists[option.dest] = 1
							πF.SetLineno(640)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µself, ßlists, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							πTemp011 = πTemp013
							if πE = πg.SetItem(πF, πTemp010, πTemp011, πTemp008); πE != nil {
								continue
							}
							goto Label16
						Label16:
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcomponent, ßsettings_defaults, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label17
							}
							goto Label18
							// line 641: if component.settings_defaults:
							πF.SetLineno(641)
						Label17:
							// line 642: self.defaults.update(component.settings_defaults)
							πF.SetLineno(642)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcomponent, ßsettings_defaults, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp005, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label18
						Label18:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µcomponents, "components"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µcomponents); πE != nil {
								continue
							}
							πF.PushCheckpoint(20)
							πTemp002 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label21
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
								µcomponent = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(19)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							πTemp004 = µcomponent
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label22
							}
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcomponent, ßsettings_default_overrides, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
						Label22:
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label23
							}
							goto Label24
							// line 644: if component and component.settings_default_overrides:
							πF.SetLineno(644)
						Label23:
							// line 645: self.defaults.update(component.settings_default_overrides)
							πF.SetLineno(645)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcomponent, ßsettings_default_overrides, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label24
						Label24:
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpopulate_from_components.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 617: """
					πF.SetLineno(617)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        For each component, first populate from the `SettingsSpec.settings_spec`\n        structure, then from the `SettingsSpec.settings_defaults` dictionary.\n        After all components have been processed, check for and populate from\n        each component's `SettingsSpec.settings_default_overrides` dictionary.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßpopulate_from_components); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 647: def get_standard_config_files(self):
					πF.SetLineno(647)
					πTemp014 = make([]πg.Param, 1)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("get_standard_config_files", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µconfig_files *πg.Object = πg.UnboundLocal
						_ = µconfig_files
						var µexpand *πg.Object = πg.UnboundLocal
						_ = µexpand
						var µpwd *πg.Object = πg.UnboundLocal
						_ = µpwd
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []πg.Param
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2:
								goto Label2
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 648: """Return list of config files, from environment or standard."""
							πF.SetLineno(648)
							// line 649: try:
							πF.SetLineno(649)
							πF.PushCheckpoint(2)
							// line 650: config_files = os.environ['DOCUTILSCONFIG'].split(os.pathsep)
							πF.SetLineno(650)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpathsep, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp002 = ßDOCUTILSCONFIG.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßenviron, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µconfig_files = πTemp003
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 651: except KeyError:
							πF.SetLineno(651)
						Label3:
							// line 652: config_files = self.standard_config_files
							πF.SetLineno(652)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstandard_config_files, nil); πE != nil {
								continue
							}
							µconfig_files = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 659: expand = os.path.expanduser
							πF.SetLineno(659)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpath, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßexpanduser, nil); πE != nil {
								continue
							}
							µexpand = πTemp002
							if πTemp003, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßenviron, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, πTemp004, ßHOME.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label4
							}
							goto Label5
							// line 660: if 'HOME' not in os.environ:
							πF.SetLineno(660)
						Label4:
							// line 661: try:
							πF.SetLineno(661)
							πF.PushCheckpoint(7)
							// line 662: import pwd
							πF.SetLineno(662)
							if πTemp001, πE = πg.ImportModule(πF, "pwd"); πE != nil {
								continue
							}
							πTemp002 = πTemp001[0]
							µpwd = πTemp002
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label8
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 663: except ImportError:
							πF.SetLineno(663)
						Label8:
							// line 664: expand = lambda x: x
							πF.SetLineno(664)
							πTemp009 = make([]πg.Param, 1)
							πTemp009[0] = πg.Param{Name: "x", Def: nil}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]
								_ = µx
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
									// line 664: expand = lambda x: x
									πF.SetLineno(664)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πR = µx
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µexpand = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							goto Label5
						Label5:
							// line 665: return [expand(f) for f in config_files if f.strip()]
							πF.SetLineno(665)
							πTemp009 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µf *πg.Object = πg.UnboundLocal
								_ = µf
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
										if πE = πg.CheckLocal(πF, µconfig_files, "config_files"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µconfig_files); πE != nil {
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
											µf = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)
										if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µf, ßstrip, nil); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 665: return [expand(f) for f in config_files if f.strip()]
										πF.SetLineno(665)
									Label4:
										// line 665: return [expand(f) for f in config_files if f.strip()]
										πF.SetLineno(665)
										πTemp006 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
											continue
										}
										πTemp006[0] = µf
										if πE = πg.CheckLocal(πF, µexpand, "expand"); πE != nil {
											continue
										}
										if πTemp004, πE = µexpand.Call(πF, πTemp006, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp006)
										πF.PushCheckpoint(6)
										return πTemp004, nil
									Label6:
										πTemp005 = πSent
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_standard_config_files.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 648: """Return list of config files, from environment or standard."""
					πF.SetLineno(648)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Return list of config files, from environment or standard.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßget_standard_config_files); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 667: def get_standard_config_settings(self):
					πF.SetLineno(667)
					πTemp014 = make([]πg.Param, 1)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("get_standard_config_settings", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µfilename *πg.Object = πg.UnboundLocal
						_ = µfilename
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
						var πTemp007 []*πg.Object
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
							// line 668: settings = Values()
							πF.SetLineno(668)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsettings = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_standard_config_files, nil); πE != nil {
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
								µfilename = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 670: settings.update(self.get_config_file_settings(filename), self)
							πF.SetLineno(670)
							πTemp006 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp007[0] = µfilename
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_config_file_settings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[1] = µself
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsettings, ßupdate, nil); πE != nil {
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
							// line 671: return settings
							πF.SetLineno(671)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							πR = µsettings
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_standard_config_settings.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 673: def get_config_file_settings(self, config_file):
					πF.SetLineno(673)
					πTemp014 = make([]πg.Param, 2)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp014[1] = πg.Param{Name: "config_file", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("get_config_file_settings", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µconfig_file *πg.Object = πArgs[1]
						_ = µconfig_file
						var µparser *πg.Object = πg.UnboundLocal
						_ = µparser
						var µbase_path *πg.Object = πg.UnboundLocal
						_ = µbase_path
						var µapplied *πg.Object = πg.UnboundLocal
						_ = µapplied
						var µsettings *πg.Object = πg.UnboundLocal
						_ = µsettings
						var µcomponent *πg.Object = πg.UnboundLocal
						_ = µcomponent
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 []*πg.Object
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 674: """Returns a dictionary containing appropriate config file settings."""
							πF.SetLineno(674)
							// line 675: parser = ConfigParser()
							πF.SetLineno(675)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßConfigParser); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µparser = πTemp002
							// line 676: parser.read(config_file, self)
							πF.SetLineno(676)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µconfig_file, "config_file"); πE != nil {
								continue
							}
							πTemp003[0] = µconfig_file
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[1] = µself
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 677: self.config_files.extend(parser._files)
							πF.SetLineno(677)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ß_files, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßconfig_files, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 678: base_path = os.path.dirname(config_file)
							πF.SetLineno(678)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µconfig_file, "config_file"); πE != nil {
								continue
							}
							πTemp003[0] = µconfig_file
							if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpath, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdirname, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µbase_path = πTemp002
							// line 679: applied = {}
							πF.SetLineno(679)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							µapplied = πTemp001
							// line 680: settings = Values()
							πF.SetLineno(680)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsettings = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcomponents, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µcomponent = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µcomponent); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 682: if not component:
							πF.SetLineno(682)
						Label4:
							// line 683: continue
							πF.SetLineno(683)
							continue
							goto Label5
						Label5:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µcomponent, ßconfig_section_dependencies, nil); πE != nil {
								continue
							}
							πTemp008 = πTemp009
							if πTemp006, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							πTemp009 = πg.NewTuple0().ToObject()
							πTemp008 = πTemp009
						Label6:
							πTemp003[0] = πTemp008
							if πTemp008, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µcomponent, "component"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µcomponent, ßconfig_section, nil); πE != nil {
								continue
							}
							πTemp008 = πg.NewTuple1(πTemp010).ToObject()
							if πTemp007, πE = πg.Add(πF, πTemp009, πTemp008); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp006 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp007, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								µsection = πTemp007
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(7)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µapplied, "applied"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Contains(πF, µapplied, µsection); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(πTemp011).ToObject()
							if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label10
							}
							goto Label11
							// line 686: if section in applied:
							πF.SetLineno(686)
						Label10:
							// line 687: continue
							πF.SetLineno(687)
							continue
							goto Label11
						Label11:
							// line 688: applied[section] = 1
							πF.SetLineno(688)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µapplied, "applied"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp008 = µsection
							if πE = πg.SetItem(πF, µapplied, πTemp008, πTemp007); πE != nil {
								continue
							}
							// line 689: settings.update(parser.get_section(section), self)
							πF.SetLineno(689)
							πTemp003 = πF.MakeArgs(2)
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp012[0] = µsection
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µparser, ßget_section, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp003[0] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[1] = µself
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µsettings, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 690: make_paths_absolute(
							πF.SetLineno(690)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrelative_path_settings, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µbase_path, "base_path"); πE != nil {
								continue
							}
							πTemp003[2] = µbase_path
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmake_paths_absolute); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 692: return settings.__dict__
							πF.SetLineno(692)
							if πE = πg.CheckLocal(πF, µsettings, "settings"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsettings, ß__dict__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_config_file_settings.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 674: """Returns a dictionary containing appropriate config file settings."""
					πF.SetLineno(674)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Returns a dictionary containing appropriate config file settings.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßget_config_file_settings); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 694: def check_values(self, values, args):
					πF.SetLineno(694)
					πTemp014 = make([]πg.Param, 3)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp014[1] = πg.Param{Name: "values", Def: nil}
					πTemp014[2] = πg.Param{Name: "args", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("check_values", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvalues *πg.Object = πArgs[1]
						_ = µvalues
						var µargs *πg.Object = πArgs[2]
						_ = µargs
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
							// line 695: """Store positional arguments as runtime settings."""
							πF.SetLineno(695)
							// line 696: values._source, values._destination = self.check_args(args)
							πF.SetLineno(696)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheck_args, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µvalues, ß_source, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µvalues, ß_destination, πTemp004); πE != nil {
								continue
							}
							// line 697: make_paths_absolute(values.__dict__, self.relative_path_settings)
							πF.SetLineno(697)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µvalues, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrelative_path_settings, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmake_paths_absolute); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 698: values._config_files = self.config_files
							πF.SetLineno(698)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßconfig_files, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µvalues, ß_config_files, πTemp003); πE != nil {
								continue
							}
							// line 699: return values
							πF.SetLineno(699)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πR = µvalues
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_values.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 695: """Store positional arguments as runtime settings."""
					πF.SetLineno(695)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Store positional arguments as runtime settings.").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_values); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 701: def check_args(self, args):
					πF.SetLineno(701)
					πTemp014 = make([]πg.Param, 2)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp014[1] = πg.Param{Name: "args", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("check_args", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µsource *πg.Object = πg.UnboundLocal
						_ = µsource
						var µdestination *πg.Object = πg.UnboundLocal
						_ = µdestination
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 702: source = destination = None
							πF.SetLineno(702)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µsource = πTemp001
							µdestination = πTemp001
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 703: if args:
							πF.SetLineno(703)
						Label1:
							// line 704: source = args.pop(0)
							πF.SetLineno(704)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µargs, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µsource = πTemp004
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µsource, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 705: if source == '-':           # means stdin
							πF.SetLineno(705)
						Label3:
							// line 706: source = None
							πF.SetLineno(706)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µsource = πTemp001
							goto Label4
						Label4:
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 707: if args:
							πF.SetLineno(707)
						Label5:
							// line 708: destination = args.pop(0)
							πF.SetLineno(708)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µargs, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdestination = πTemp004
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µdestination, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 709: if destination == '-':      # means stdout
							πF.SetLineno(709)
						Label7:
							// line 710: destination = None
							πF.SetLineno(710)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µdestination = πTemp001
							goto Label8
						Label8:
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 711: if args:
							πF.SetLineno(711)
						Label9:
							// line 712: self.error('Maximum 2 arguments allowed.')
							πF.SetLineno(712)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Maximum 2 arguments allowed.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							πTemp001 = µsource
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µsource, µdestination); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label11:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label12
							}
							goto Label13
							// line 713: if source and source == destination:
							πF.SetLineno(713)
						Label12:
							// line 714: self.error('Do not specify the same file for both source and '
							πF.SetLineno(714)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Do not specify the same file for both source and destination.  It will clobber the source file.").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label13
						Label13:
							// line 716: return source, destination
							πF.SetLineno(716)
							if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdestination, "destination"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µsource, µdestination).ToObject()
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
					if πE = πClass.SetItem(πF, ßcheck_args.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 718: def set_defaults_from_dict(self, defaults):
					πF.SetLineno(718)
					πTemp014 = make([]πg.Param, 2)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp014[1] = πg.Param{Name: "defaults", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("set_defaults_from_dict", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdefaults *πg.Object = πArgs[1]
						_ = µdefaults
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
							// line 719: self.defaults.update(defaults)
							πF.SetLineno(719)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							πTemp001[0] = µdefaults
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_defaults_from_dict.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 721: def get_default_values(self):
					πF.SetLineno(721)
					πTemp014 = make([]πg.Param, 1)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("get_default_values", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdefaults *πg.Object = πg.UnboundLocal
						_ = µdefaults
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
							// line 722: """Needed to get custom `Values` instances."""
							πF.SetLineno(722)
							// line 723: defaults = Values(self.defaults)
							πF.SetLineno(723)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdefaults = πTemp003
							// line 724: defaults._config_files = self.config_files
							πF.SetLineno(724)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßconfig_files, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdefaults, ß_config_files, πTemp003); πE != nil {
								continue
							}
							// line 725: return defaults
							πF.SetLineno(725)
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							πR = µdefaults
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_default_values.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 722: """Needed to get custom `Values` instances."""
					πF.SetLineno(722)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Needed to get custom `Values` instances.").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßget_default_values); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 727: def get_option_by_dest(self, dest):
					πF.SetLineno(727)
					πTemp014 = make([]πg.Param, 2)
					πTemp014[0] = πg.Param{Name: "self", Def: nil}
					πTemp014[1] = πg.Param{Name: "dest", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("get_option_by_dest", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdest *πg.Object = πArgs[1]
						_ = µdest
						var µgroup *πg.Object = πg.UnboundLocal
						_ = µgroup
						var µoption *πg.Object = πg.UnboundLocal
						_ = µoption
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 728: """
							πF.SetLineno(728)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoption_groups, nil); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µgroup = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µgroup, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp007 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µoption = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, µdest); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 739: if option.dest == dest:
							πF.SetLineno(739)
						Label7:
							// line 740: return option
							πF.SetLineno(740)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πR = µoption
							continue
							goto Label8
						Label8:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("No option with dest == %r.").ToObject(), µdest); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 741: raise KeyError('No option with dest == %r.' % dest)
							πF.SetLineno(741)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_option_by_dest.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 728: """
					πF.SetLineno(728)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("\n        Get an option by its dest.\n\n        If you're supplying a dest which is shared by several options,\n        it is undefined which option of those is returned.\n\n        A KeyError is raised if there is no option with the supplied\n        dest.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßget_option_by_dest); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp023, πE = πTemp024.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp023 == nil {
				πTemp023 = πg.TypeType.ToObject()
			}
			if πTemp025, πE = πTemp023.Call(πF, []*πg.Object{πg.NewStr("OptionParser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp024.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionParser.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 744: class ConfigParser(RawConfigParser):
			πF.SetLineno(744)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp025, πE = πg.ResolveGlobal(πF, ßRawConfigParser); πE != nil {
				continue
			}
			πTemp002[0] = πTemp025
			πTemp024 = πg.NewDict()
			if πTemp022, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp024.SetItem(πF, ß__module__.ToObject(), πTemp022); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ConfigParser", "/usr/lib/python2.7/site-packages/docutils/frontend.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp024
				_ = πClass
				var πTemp001 *πg.Dict
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 746: old_settings = {
					πF.SetLineno(746)
					πTemp001 = πg.NewDict()
					πTemp002 = πg.NewTuple2(πg.NewStr("pep_html writer").ToObject(), ßstylesheet.ToObject()).ToObject()
					if πE = πTemp001.SetItem(πF, ßpep_stylesheet.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewStr("pep_html writer").ToObject(), ßstylesheet_path.ToObject()).ToObject()
					if πE = πTemp001.SetItem(πF, ßpep_stylesheet_path.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewStr("pep_html writer").ToObject(), ßtemplate.ToObject()).ToObject()
					if πE = πTemp001.SetItem(πF, ßpep_template.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßold_settings.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 750: """{old setting: (new section, new setting)} mapping, used by
					πF.SetLineno(750)
					// line 750: """{old setting: (new section, new setting)} mapping, used by
					πF.SetLineno(750)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("{old setting: (new section, new setting)} mapping, used by\n    `handle_old_config`, to convert settings from the old [options] section.").ToObject()); πE != nil {
						continue
					}
					// line 753: old_warning = """
					πF.SetLineno(753)
					if πE = πClass.SetItem(πF, ßold_warning.ToObject(), πg.NewStr("\nThe \"[option]\" section is deprecated.  Support for old-format configuration\nfiles may be removed in a future Docutils release.  Please revise your\nconfiguration files.  See <http://docutils.sf.net/docs/user/config.html>,\nsection \"Old-Format Configuration Files\".\n").ToObject()); πE != nil {
						continue
					}
					// line 760: not_utf8_error = """\
					πF.SetLineno(760)
					if πE = πClass.SetItem(πF, ßnot_utf8_error.ToObject(), πg.NewStr("Unable to read configuration file \"%s\": content not encoded as UTF-8.\nSkipping \"%s\" configuration file.\n").ToObject()); πE != nil {
						continue
					}
					// line 765: def __init__(self, *args, **kwargs):
					πF.SetLineno(765)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp003, πg.CodeFlagVarArg|πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πArgs[1]
						_ = µargs
						var µkwargs *πg.Object = πArgs[2]
						_ = µkwargs
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
							// line 766: RawConfigParser.__init__(self, *args, **kwargs)
							πF.SetLineno(766)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRawConfigParser); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, πTemp001, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 768: self._files = []
							πF.SetLineno(768)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_files, πTemp003); πE != nil {
								continue
							}
							// line 769: """List of paths of configuration files read."""
							πF.SetLineno(769)
							// line 771: self._stderr = ErrorOutput()
							πF.SetLineno(771)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßErrorOutput); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stderr, πTemp002); πE != nil {
								continue
							}
							// line 772: """Wrapper around sys.stderr catching en-/decoding errors"""
							πF.SetLineno(772)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 774: def read(self, filenames, option_parser):
					πF.SetLineno(774)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "filenames", Def: nil}
					πTemp003[2] = πg.Param{Name: "option_parser", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfilenames *πg.Object = πArgs[1]
						_ = µfilenames
						var µoption_parser *πg.Object = πArgs[2]
						_ = µoption_parser
						var µfilename *πg.Object = πg.UnboundLocal
						_ = µfilename
						var µfp *πg.Object = πg.UnboundLocal
						_ = µfp
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							case 10:
								goto Label10
							case 3:
								goto Label3
							case 4:
								goto Label4
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilenames, "filenames"); πE != nil {
								continue
							}
							πTemp002[0] = µfilenames
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πTemp007, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 775: if type(filenames) in (str, unicode):
							πF.SetLineno(775)
						Label1:
							// line 776: filenames = [filenames]
							πF.SetLineno(776)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µfilenames, "filenames"); πE != nil {
								continue
							}
							πTemp002[0] = µfilenames
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µfilenames = πTemp001
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µfilenames, "filenames"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µfilenames); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp007 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µfilename = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 778: try:
							πF.SetLineno(778)
							πF.PushCheckpoint(7)
							// line 780: fp = codecs.open(filename, 'r', 'utf-8')
							πF.SetLineno(780)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp002[0] = µfilename
							πTemp002[1] = ßr.ToObject()
							πTemp002[2] = πg.NewStr("utf-8").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcodecs); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfp = πTemp003
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label8
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 781: except IOError:
							πF.SetLineno(781)
						Label8:
							// line 782: continue
							πF.SetLineno(782)
							continue
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							// line 783: try:
							πF.SetLineno(783)
							πF.PushCheckpoint(10)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßversion_info, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							if πTemp003, πE = πg.LT(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label11
							}
							goto Label12
							// line 784: if sys.version_info < (3, 0):
							πF.SetLineno(784)
						Label11:
							// line 785: RawConfigParser.readfp(self, fp, filename)
							πF.SetLineno(785)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							πTemp002[1] = µfp
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp002[2] = µfilename
							if πTemp003, πE = πg.ResolveGlobal(πF, ßRawConfigParser); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreadfp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label13
						Label12:
							// line 787: RawConfigParser.read_file(self, fp, filename)
							πF.SetLineno(787)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							πTemp002[1] = µfp
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp002[2] = µfilename
							if πTemp003, πE = πg.ResolveGlobal(πF, ßRawConfigParser); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread_file, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label13
						Label13:
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label14
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 788: except UnicodeDecodeError:
							πF.SetLineno(788)
						Label14:
							// line 789: self._stderr.write(self.not_utf8_error % (filename, filename))
							πF.SetLineno(789)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßnot_utf8_error, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µfilename, µfilename).ToObject()
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_stderr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 790: fp.close()
							πF.SetLineno(790)
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfp, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 791: continue
							πF.SetLineno(791)
							continue
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							// line 792: fp.close()
							πF.SetLineno(792)
							if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µfp, ßclose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 793: self._files.append(filename)
							πF.SetLineno(793)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp002[0] = µfilename
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_files, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßoptions.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhas_section, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label15
							}
							goto Label16
							// line 794: if self.has_section('options'):
							πF.SetLineno(794)
						Label15:
							// line 795: self.handle_old_config(filename)
							πF.SetLineno(795)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp002[0] = µfilename
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhandle_old_config, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label16
						Label16:
							// line 796: self.validate_settings(filename, option_parser)
							πF.SetLineno(796)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp002[0] = µfilename
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							πTemp002[1] = µoption_parser
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßvalidate_settings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 798: def handle_old_config(self, filename):
					πF.SetLineno(798)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "filename", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("handle_old_config", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfilename *πg.Object = πArgs[1]
						_ = µfilename
						var µoptions *πg.Object = πg.UnboundLocal
						_ = µoptions
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
						var µsetting *πg.Object = πg.UnboundLocal
						_ = µsetting
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 799: warnings.warn_explicit(self.old_warning, ConfigDeprecationWarning,
							πF.SetLineno(799)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßold_warning, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßConfigDeprecationWarning); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp001[2] = µfilename
							πTemp001[3] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwarn_explicit, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 801: options = self.get_section('options')
							πF.SetLineno(801)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßoptions.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_section, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µoptions = πTemp003
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßgeneral.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßhas_section, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 802: if not self.has_section('general'):
							πF.SetLineno(802)
						Label1:
							// line 803: self.add_section('general')
							πF.SetLineno(803)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßgeneral.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßadd_section, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoptions, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßold_settings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, µkey); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 805: if key in self.old_settings:
							πF.SetLineno(805)
						Label6:
							// line 806: section, setting = self.old_settings[key]
							πF.SetLineno(806)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßold_settings, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
								continue
							}
							µsection = πTemp003
							µsetting = πTemp007
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp001[0] = µsection
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhas_section, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 807: if not self.has_section(section):
							πF.SetLineno(807)
						Label9:
							// line 808: self.add_section(section)
							πF.SetLineno(808)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp001[0] = µsection
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßadd_section, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
						Label10:
							goto Label8
						Label7:
							// line 810: section = 'general'
							πF.SetLineno(810)
							µsection = ßgeneral.ToObject()
							// line 811: setting = key
							πF.SetLineno(811)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							µsetting = µkey
							goto Label8
						Label8:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp001[0] = µsection
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[1] = µsetting
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhas_option, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label11
							}
							goto Label12
							// line 812: if not self.has_option(section, setting):
							πF.SetLineno(812)
						Label11:
							// line 813: self.set(section, setting, value)
							πF.SetLineno(813)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp001[0] = µsection
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp001[1] = µsetting
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[2] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßset, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label12
						Label12:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 814: self.remove_section('options')
							πF.SetLineno(814)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßoptions.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßremove_section, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßhandle_old_config.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 816: def validate_settings(self, filename, option_parser):
					πF.SetLineno(816)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "filename", Def: nil}
					πTemp003[2] = πg.Param{Name: "option_parser", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("validate_settings", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfilename *πg.Object = πArgs[1]
						_ = µfilename
						var µoption_parser *πg.Object = πArgs[2]
						_ = µoption_parser
						var µsection *πg.Object = πg.UnboundLocal
						_ = µsection
						var µsetting *πg.Object = πg.UnboundLocal
						_ = µsetting
						var µoption *πg.Object = πg.UnboundLocal
						_ = µoption
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µnew_value *πg.Object = πg.UnboundLocal
						_ = µnew_value
						var µerror *πg.Object = πg.UnboundLocal
						_ = µerror
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 πg.KWArgs
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
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
							case 4:
								goto Label4
							case 5:
								goto Label5
							case 8:
								goto Label8
							case 13:
								goto Label13
							default:
								panic("unexpected function state")
							}
							// line 817: """
							πF.SetLineno(817)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsections, nil); πE != nil {
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
								µsection = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp006[0] = µsection
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µsetting = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 823: try:
							πF.SetLineno(823)
							πF.PushCheckpoint(8)
							// line 824: option = option_parser.get_option_by_dest(setting)
							πF.SetLineno(824)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp006[0] = µsetting
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption_parser, ßget_option_by_dest, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µoption = πTemp007
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 825: except KeyError:
							πF.SetLineno(825)
						Label9:
							// line 826: continue
							πF.SetLineno(826)
							continue
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßvalidator, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label10
							}
							goto Label11
							// line 827: if option.validator:
							πF.SetLineno(827)
						Label10:
							// line 828: value = self.get(section, setting)
							πF.SetLineno(828)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp006[0] = µsection
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp006[1] = µsetting
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µvalue = πTemp007
							// line 829: try:
							πF.SetLineno(829)
							πF.PushCheckpoint(13)
							// line 830: new_value = option.validator(
							πF.SetLineno(830)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp006[0] = µsetting
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp006[1] = µvalue
							if πE = πg.CheckLocal(πF, µoption_parser, "option_parser"); πE != nil {
								continue
							}
							πTemp006[2] = µoption_parser
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp011 = πg.KWArgs{
								{"config_parser", µself},
								{"config_section", µsection},
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßvalidator, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, πTemp011); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µnew_value = πTemp007
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label14
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 833: except Exception as error:
							πF.SetLineno(833)
						Label14:
							µerror = πTemp009.ToObject()
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerror, "error"); πE != nil {
								continue
							}
							πTemp012[0] = µerror
							if πTemp013, πE = πg.ResolveGlobal(πF, ßErrorString); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple5(µfilename, µsection, πTemp014, µsetting, µvalue).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Error in config file \"%s\", section \"[%s]\":\n    %s\n        %s = %s").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 834: raise ValueError(
							πF.SetLineno(834)
							πE = πF.Raise(πTemp007, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							// line 840: self.set(section, setting, new_value)
							πF.SetLineno(840)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp006[0] = µsection
							if πE = πg.CheckLocal(πF, µsetting, "setting"); πE != nil {
								continue
							}
							πTemp006[1] = µsetting
							if πE = πg.CheckLocal(πF, µnew_value, "new_value"); πE != nil {
								continue
							}
							πTemp006[2] = µnew_value
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßset, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßoverrides, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label15
							}
							goto Label16
							// line 841: if option.overrides:
							πF.SetLineno(841)
						Label15:
							// line 842: self.set(section, option.overrides, None)
							πF.SetLineno(842)
							πTemp006 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp006[0] = µsection
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßoverrides, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp006[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßset, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label16
						Label16:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
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
					if πE = πClass.SetItem(πF, ßvalidate_settings.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 817: """
					πF.SetLineno(817)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Call the validator function and implement overrides on all applicable\n        settings.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßvalidate_settings); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 844: def optionxform(self, optionstr):
					πF.SetLineno(844)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "optionstr", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("optionxform", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µoptionstr *πg.Object = πArgs[1]
						_ = µoptionstr
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
							// line 845: """
							πF.SetLineno(845)
							// line 848: return optionstr.lower().replace('-', '_')
							πF.SetLineno(848)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-").ToObject()
							πTemp001[1] = ß_.ToObject()
							if πE = πg.CheckLocal(πF, µoptionstr, "optionstr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoptionstr, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßoptionxform.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 845: """
					πF.SetLineno(845)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        Transform '-' to '_' so the cmdline form of option names can be used.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßoptionxform); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 850: def get_section(self, section):
					πF.SetLineno(850)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "section", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("get_section", "/usr/lib/python2.7/site-packages/docutils/frontend.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsection *πg.Object = πArgs[1]
						_ = µsection
						var µsection_dict *πg.Object = πg.UnboundLocal
						_ = µsection_dict
						var µoption *πg.Object = πg.UnboundLocal
						_ = µoption
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
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
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 851: """
							πF.SetLineno(851)
							// line 855: section_dict = {}
							πF.SetLineno(855)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µsection_dict = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp003[0] = µsection
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhas_section, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 856: if self.has_section(section):
							πF.SetLineno(856)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp003[0] = µsection
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoptions, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp005 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µoption = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 858: section_dict[option] = self.get(section, option)
							πF.SetLineno(858)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsection, "section"); πE != nil {
								continue
							}
							πTemp003[0] = µsection
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp003[1] = µoption
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsection_dict, "section_dict"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp008 = µoption
							if πE = πg.SetItem(πF, µsection_dict, πTemp008, πTemp004); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							goto Label2
						Label2:
							// line 859: return section_dict
							πF.SetLineno(859)
							if πE = πg.CheckLocal(πF, µsection_dict, "section_dict"); πE != nil {
								continue
							}
							πR = µsection_dict
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget_section.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 851: """
					πF.SetLineno(851)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n        Return a given section as a dictionary (empty if the section\n        doesn't exist).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßget_section); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp023, πE = πTemp024.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp023 == nil {
				πTemp023 = πg.TypeType.ToObject()
			}
			if πTemp025, πE = πTemp023.Call(πF, []*πg.Object{πg.NewStr("ConfigParser").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp024.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßConfigParser.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 862: class ConfigDeprecationWarning(DeprecationWarning):
			πF.SetLineno(862)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp025, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
				continue
			}
			πTemp002[0] = πTemp025
			πTemp024 = πg.NewDict()
			if πTemp022, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp024.SetItem(πF, ß__module__.ToObject(), πTemp022); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ConfigDeprecationWarning", "/usr/lib/python2.7/site-packages/docutils/frontend.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp024
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 863: """Warning for deprecated configuration file features."""
					πF.SetLineno(863)
					// line 863: """Warning for deprecated configuration file features."""
					πF.SetLineno(863)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Warning for deprecated configuration file features.").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp023, πE = πTemp024.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp023 == nil {
				πTemp023 = πg.TypeType.ToObject()
			}
			if πTemp025, πE = πTemp023.Call(πF, []*πg.Object{πg.NewStr("ConfigDeprecationWarning").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp024.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßConfigDeprecationWarning.ToObject(), πTemp025); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("frontend", Code)
}
