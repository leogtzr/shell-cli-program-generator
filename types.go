package shellcligen

// CLIProgram ...
type CLIProgram struct {
	Help    string      `json:"message" yaml:"help_message"`
	Options []CLIOption `json:"options" yaml:"options"`
}

// CLIOption ...
type CLIOption struct {
	// Name ...
	Name string `json:"name" yaml:"name"`
	// Required ...
	Required bool `json:"required" yaml:"required"`
	// ArgsNum ...
	ArgsNum int `json:"args_num" yaml:"args_num"`
}
