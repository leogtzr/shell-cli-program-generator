package shellcligen

import (
	"testing"
)

func Test_hasRequiredOptions(t *testing.T) {
	t.Parallel()

	type test struct {
		cliProgram CLIProgram
		has        bool
	}

	tests := []test{
		{
			cliProgram: CLIProgram{
				Help: ``,
				Options: []CLIOption{
					{
						Required: false,
						LongName: "article",
					},
					{
						Required:  false,
						ShortName: "a",
					},
					{
						Required: true,
						LongName: "verbose",
					},
				},
			},
			has: true,
		},
		{
			cliProgram: CLIProgram{
				Help: ``,
				Options: []CLIOption{
					{
						Required: false,
						LongName: "article",
					},
					{
						Required:  false,
						ShortName: "a",
					},
				},
			},
			has: false,
		},
	}

	for _, tt := range tests {
		got := hasRequiredOptions(&tt.cliProgram)
		if got != tt.has {
			t.Errorf("got=%t, expected=%t", got, tt.has)
		}
	}
}

func TestCLIOption_String(t *testing.T) {
	t.Parallel()

	type test struct {
		option CLIOption
		want   string
	}

	tests := []test{
		{
			option: CLIOption{
				ConflictsWith: []string{},
				ArgsRequired:  false,
				LongName:      "verbose",
				ShortName:     "v",
				Required:      false,
			},
			want: "Long name: `verbose`, Short name: `v`, Required: false",
		},
	}

	for _, tt := range tests {
		got := tt.option.String()
		if got != tt.want {
			t.Errorf("got=[%s] wants=[%s]", got, tt.want)
		}
	}
}

func Test_existInArrayName(t *testing.T) {
	t.Parallel()

	type test struct {
		optName string
		names   []Name
		wants   bool
	}

	tests := []test{
		{
			optName: "P",
			names: []Name{
				{
					Short: "P",
					Long:  "",
				},
			},
			wants: true,
		},
	}

	for _, tt := range tests {
		got := existInArrayName(tt.optName, &tt.names)
		if got != tt.wants {
			t.Errorf("got=[%t], wants=[%t]", got, tt.wants)
		}
	}
}

func Test_hasTheOption(t *testing.T) {
	t.Parallel()

	allOptionNames := []Name{
		{
			Short: "P",
			Long:  "",
		},
		{
			Long:  "extended-regexp",
			Short: "E",
		},
	}

	type test struct {
		conflictsWith []string
		hasWant       bool
	}

	tests := []test{
		{
			conflictsWith: []string{"P"},
			hasWant:       true,
		},
		{
			conflictsWith: []string{"X", "extended-regexp"},
			hasWant:       true,
		},
	}

	for _, tt := range tests {
		if got := hasTheOption(&tt.conflictsWith, &allOptionNames); got != tt.hasWant {
			t.Errorf("got=[%t], wants=[%t]", got, tt.hasWant)
		}
	}
}

func Test_validateOptionNames(t *testing.T) {
	t.Parallel()

	type test struct {
		cliProram           CLIProgram
		wantsValidConflicts bool
	}

	tests := []test{
		{
			cliProram: CLIProgram{
				Help: `HelpTxtMessage1`,
				Options: []CLIOption{
					{
						LongName:      "extended-regexp",
						ShortName:     "E",
						Required:      false,
						ConflictsWith: []string{"P"},
					},
					{
						LongName:      "extended-regexp",
						ShortName:     "P",
						Required:      false,
						ConflictsWith: []string{"extended-regexp"},
					},
				},
			},
			wantsValidConflicts: true,
		},
		{
			cliProram: CLIProgram{
				Help: `HelpTxtMessage2`,
				Options: []CLIOption{
					{
						LongName:      "extended-regexp",
						ShortName:     "E",
						Required:      false,
						ConflictsWith: []string{},
					},
					{
						LongName:      "verbose",
						ShortName:     "v",
						Required:      false,
						ConflictsWith: []string{},
					},
				},
			},
			wantsValidConflicts: true,
		},
		{
			cliProram: CLIProgram{
				Help: `HelpTxtMessage3`,
				Options: []CLIOption{
					{
						LongName:      "version",
						ShortName:     "v",
						Required:      false,
						ConflictsWith: []string{},
					},
					{
						LongName:      "description",
						ShortName:     "d",
						Required:      false,
						ConflictsWith: []string{},
					},
				},
			},
			wantsValidConflicts: true,
		},
		// "X" option doesn't exist, validation should fail.
		{
			cliProram: CLIProgram{
				Help: `HelpTxtMessage4`,
				Options: []CLIOption{
					{
						LongName:      "version",
						ShortName:     "v",
						Required:      false,
						ConflictsWith: []string{"X"},
					},
					{
						LongName:      "description",
						ShortName:     "d",
						Required:      false,
						ConflictsWith: []string{},
					},
				},
			},
			wantsValidConflicts: false,
		},
	}

	for _, tt := range tests {
		if got := validateExistingConflictingOptionNames(&tt.cliProram); got != tt.wantsValidConflicts {
			t.Errorf("got=%t, wants=%t for cli option with help message `%s`", got, tt.wantsValidConflicts, tt.cliProram.Help)
		}
	}
}
