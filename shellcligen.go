package shellcligen

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrOpeningInputFile        = errors.New("error opening input file")
	ErrReadingInputFile        = errors.New("error reading input file")
	ErrParsingInputFile        = errors.New("error parsing input file")
	ErrMissingRequiredArgument = errors.New("error missing required argument")
	ErrParsingConflictOptions  = errors.New("error parsing options with conflicts")
	ErrInvalidOptionName       = errors.New("error invalid option name")

	cliOptionRegex = regexp.MustCompile("^[a-zA-Z_]([a-zA-Z0-9_]*)$")
)

func isOptionNameValid(optionName string, rgx *regexp.Regexp) bool {
	return rgx.MatchString(optionName)
}

func existInArrayName(name string, arr *[]Name) bool {
	exist := false

	for _, x := range *arr {
		if strings.TrimSpace(x.Long) == name || strings.TrimSpace(x.Short) == name {
			exist = true

			break
		}
	}

	return exist
}

func hasTheOption(options *[]string, names *[]Name) bool {
	exists := false

	for _, option := range *options {
		if existInArrayName(option, names) {
			exists = true

			break
		}
	}

	return exists
}

func validateExistingConflictingOptionNames(cli *CLIProgram) bool {
	valid := true

	optionNames := make([]Name, 0)

	for _, opt := range cli.Options {
		optionNames = append(optionNames, Name{
			Short: opt.ShortName,
			Long:  opt.LongName,
		})
	}

	for _, opt := range cli.Options {
		if len(opt.ConflictsWith) > 0 {
			if !hasTheOption(&opt.ConflictsWith, &optionNames) {
				valid = false

				break
			}
		}
	}

	return valid
}

func validateCliOptionNames(cli *CLIProgram, regex *regexp.Regexp) bool {
	valid := true

	for _, opt := range cli.Options {
		if !isOptionNameValid(opt.LongName, regex) || !isOptionNameValid(opt.ShortName, regex) {
			valid = false

			break
		}
	}

	return valid
}

// ParseCLIProgram ...
func ParseCLIProgram(filename string) (CLIProgram, error) {
	file, err := os.Open(filename)
	if err != nil {
		return CLIProgram{}, fmt.Errorf("error opening input file: %w", ErrOpeningInputFile)
	}

	defer file.Close()

	cli := CLIProgram{}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return CLIProgram{}, fmt.Errorf("error reading input file: %w", ErrReadingInputFile)
	}

	err = yaml.Unmarshal(fileContent, &cli)
	if err != nil {
		return CLIProgram{}, fmt.Errorf("error parsing input file: %w", err)
	}

	if !validateExistingConflictingOptionNames(&cli) {
		return CLIProgram{}, fmt.Errorf("error with conflicting options: %w", ErrParsingConflictOptions)
	}

	if !validateCliOptionNames(&cli, cliOptionRegex) {
		return CLIProgram{}, fmt.Errorf("error invalid option name: %w", ErrInvalidOptionName)
	}

	return cli, nil
}

func hasRequiredOptions(cliProgram *CLIProgram) bool {
	required := false

	for _, option := range cliProgram.Options {
		if option.Required {
			required = true

			break
		}
	}

	return required
}
