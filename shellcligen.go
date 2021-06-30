package shellcligen

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
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
	ErrCreatingOutputProgram   = errors.New("error creating output script")
	ErrRepeatedOptionNames     = errors.New("error repeated option names")

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

func validateCLIOptionNames(cli *CLIProgram, regex *regexp.Regexp) bool {
	valid := true

	for _, opt := range cli.Options {
		if !isOptionNameValid(opt.LongName, regex) || !isOptionNameValid(opt.ShortName, regex) {
			valid = false

			break
		}
	}

	return valid
}

func haveRepeatedElements(cliOptionNamesCount *map[string]int) bool {
	for _, count := range *cliOptionNamesCount {
		if count > 1 {
			return true
		}
	}

	return false
}

func validateUniqueCLIOptionNamesCount(cliOptions *[]CLIOption) bool {
	shortCLIOptionNamesCount := make(map[string]int)
	longCLIOptionNamesCount := make(map[string]int)

	for _, cliOption := range *cliOptions {
		if len(cliOption.LongName) != 0 {
			longCLIOptionNamesCount[cliOption.LongName]++
		}

		if len(cliOption.ShortName) != 0 {
			shortCLIOptionNamesCount[cliOption.ShortName]++
		}
	}

	return !haveRepeatedElements(&shortCLIOptionNamesCount) && !haveRepeatedElements(&longCLIOptionNamesCount)
}

func createCliProgramScript(cli *CLIProgram, outputDirectory string) error {
	outputScriptFile, err := os.Create(path.Join(outputDirectory, scriptFileName))
	if err != nil {
		return err
	}
	defer outputScriptFile.Close()

	outputScriptConfFile, err := os.Create(path.Join(outputDirectory, scriptConfigFileName))
	if err != nil {
		return err
	}
	defer outputScriptConfFile.Close()

	template := templateWithConflictChecking
	if cli.SafeFlags {
		template = strings.ReplaceAll(template, safeFlagsTemplateTag, safeFlagsTemplate)
	}

	_, _ = outputScriptFile.WriteString(template)

	return nil
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

func optionName(cliOption *CLIOption) string {
	shortOptionName := strings.TrimSpace(cliOption.ShortName)
	longOptionName := strings.TrimSpace(cliOption.LongName)

	if len(shortOptionName) > 0 {
		return shortOptionName
	}

	return longOptionName
}

// ParseCLIProgram ...
func ParseCLIProgram(configFile, outputDirectory string) (CLIProgram, error) {
	file, err := os.Open(configFile)
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

	if !validateCLIOptionNames(&cli, cliOptionRegex) {
		return CLIProgram{}, fmt.Errorf("error invalid option name: %w", ErrInvalidOptionName)
	}

	if !validateUniqueCLIOptionNamesCount(&cli.Options) {
		return CLIProgram{}, fmt.Errorf("error repeated option names: %w", ErrRepeatedOptionNames)
	}

	if err = createCliProgramScript(&cli, outputDirectory); err != nil {
		return CLIProgram{}, fmt.Errorf("error creating output script: %w", ErrCreatingOutputProgram)
	}

	return cli, nil
}
