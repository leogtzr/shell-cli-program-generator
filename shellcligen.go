package shellcligen

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
)

func ParseCLIProgram(filename string) (CLIProgram, error) {
	file, err := os.Open(filename)
	if err != nil {
		return CLIProgram{}, err
	}

	defer file.Close()

	cli := CLIProgram{}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return CLIProgram{}, err
	}

	err = yaml.Unmarshal(fileContent, &cli)
	if err != nil {
		return CLIProgram{}, err
	}

	return cli, nil
}
