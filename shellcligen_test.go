package shellcligen

import "testing"

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
				ConflictsWith: ``,
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
