package option

import "github.com/senzing/go-common/option/optiontype"

type ContextVariable struct {
	Default any                   `json:"default"`
	Envar   string                `json:"envar"`
	Help    string                `json:"help"`
	Arg     string                `json:"option"`
	Type    optiontype.OptionType `json:"optiontype"`
}
