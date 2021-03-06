package config

import (
	"bufio"
	"fmt"
	"github.com/scinna/CLIent/api"
	"github.com/scinna/CLIent/utils"
	"os"
	"regexp"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var reader *bufio.Scanner
var re = regexp.MustCompile(`\r?\n`)

func Setup(cfg *Config) {
	reader = bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Scinna !")
	fmt.Println("Let's walk together through the initial setup")

	fmt.Println()

	fmt.Print("What's your server URL ? ")
	cfg.ServerURL = readInput()
	cfg.ServerURL = strings.TrimSpace(re.ReplaceAllString(cfg.ServerURL, ""))

	if string(cfg.ServerURL[len(cfg.ServerURL)-1]) != "/" {
		cfg.ServerURL = cfg.ServerURL + "/"
	}

	fmt.Print("Your username ? ")
	cfg.Username = readInput()
	cfg.Username = strings.TrimSpace(re.ReplaceAllString(cfg.Username, ""))

	fmt.Print("Your password ? ")
	password := readPassword()
	password = strings.TrimSpace(re.ReplaceAllString(password, ""))

	token, err := api.Login(cfg.ServerURL, cfg.Username, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg.Token = token

	fmt.Println()
	fmt.Println()

	fmt.Println("Ok! now let's configure defaults")

	accepted := false
	for !accepted {
		fmt.Print("Visibility (Public, Unlisted, Private) ? ")
		value := utils.VisibilityFromString(readInput())

		accepted = value.IsValid()
		if !accepted {
			fmt.Println("This visibility does not exists")
		} else {
			cfg.DefaultVisibility = value
		}
	}

	fmt.Println()
	fmt.Println("For the title and description, you can interpolate string with ${my_command --with-args}")
	fmt.Println()

	fmt.Print("Title ? ")
	cfg.DefaultTitle = StringCommand(readInput())

	fmt.Print("Description ? ")
	cfg.DefaultDescription = StringCommand(readInput())

	err = WriteConfiguration(cfg)
	if err != nil {
		fmt.Println("Could not write config: ", err)
	}
}

func readInput() string {
	reader.Scan()
	return reader.Text()
}

func readPassword() string {
	pwdBytes, _ := terminal.ReadPassword(syscall.Stdin)
	return re.ReplaceAllString(string(pwdBytes), "")
}