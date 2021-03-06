package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/scinna/CLIent/api"
	"github.com/scinna/CLIent/config"
	"github.com/scinna/CLIent/utils"
	"io"
	"os"
)

func main() {
	fmt.Println("Scinna CLIent")

	cfg, err := config.ReadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if !cfg.IsConfigValid() {
		config.Setup(cfg)
	}

	if !cfg.AreDefaultValid() {
		fmt.Println("The defaults are not correctly set. Please check the config file")
		return
	}

	var titleFlag = flag.String("t", "", `Title of the picture`)
	var descriptionFlag = flag.String("d", "", `Description of the picture`)
	var visibilityFlag = flag.String("v", cfg.DefaultVisibility.ToString(), `Visibility of the picture. Can be either PUBLIC, UNLISTED or PRIVATE. Case-insensitive`)
	var collectionFlag = flag.String("c", cfg.DefaultCollection, `Collection in which to put the picture. Defaults to empty string (Root collection)`)

	var useClipboard = flag.Bool("b", false, `Copy to clipboard after uploading it`)
	flag.Parse()

	if len(*titleFlag) == 0 {
		*titleFlag = string(cfg.DefaultTitle)
	}

	if len(*descriptionFlag) == 0 {
		*descriptionFlag = string(cfg.DefaultDescription)
	}

	titleCommand := config.StringCommand(*titleFlag)
	descriptionCommand := config.StringCommand(*descriptionFlag)
	visibility := utils.VisibilityFromString(*visibilityFlag)
	collection := *collectionFlag

	if !visibility.IsValid() {
		fmt.Println("Invalid visibility")
		return
	}

	var file *os.File

	if len(flag.Args()) != 0 {
		file, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println("Can't read the file!")
			fmt.Println(err)
			return
		}

		if len(*titleFlag) == 0 {
			*titleFlag = flag.Arg(0)
		}
	} else if isInputFromPipe() {
		// UGLY HACK / @TODO: fix this with reading it correctly
		// Theoretically we should stream the response so that when we'll implement video it will work
		// Currently we will just prevent users from uploading them if I was too lazy to fix this
		path := "/tmp/" + utils.RandomFilename()
		file = os.Stdin
		tmpFile, _ := os.Create(path)
		tmp, _ := io.ReadAll(file)
		tmpFile.Write(tmp)
		tmpFile.Close()
		file, _ = os.Open(path)
	} else {
		fmt.Println("No picture given!")
		return
	}

	response, err := api.Upload(cfg.ServerURL, cfg.Token, titleCommand.Process(), descriptionCommand.Process(), visibility, collection, file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Image uploaded at %v !\n", response)

	if *useClipboard {
		err := clipboard.WriteAll(response)
		if err != nil {
			fmt.Println("Could not write to clipboard: ", err)
		}
	}
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}