package shellcligen

// CLIProgram ...
type CLIProgram struct {
	Help    string      `json:"message"`
	Options []CLIOption `json:"options"`
}

// CLIOption ...
type CLIOption struct {
	// Name ...
	Name string `json:"name"`
	// Required ...
	Required bool `json:"required"`
}
