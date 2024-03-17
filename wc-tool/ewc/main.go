package main

import (
	"github.com/ercanyu/wc-tool/internal/command"
	ufcli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	cliApp := ufcli.App{
		Name:    "ewc",
		Version: "1.0.0",
		Usage:   "A simple word count tool",
		Commands: []*ufcli.Command{
			command.NewCCommand(),
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	fmt.Print("Enter commands(type 'exit' to quit): ")
//
//	for {
//		fmt.Print("wc-tool> ")
//		command, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println("Error reading input: ", err)
//			continue
//		}
//		command = command[:len(command)-1]
//
//		if command == "exit" {
//			fmt.Println("Exiting...")
//			break
//		}
//
//		_ = processCommand(command)
//	}
//}
//
//func processCommand(command string) error {
//	commandParts := strings.Split(command, " ")
//	if len(commandParts) < 3 || commandParts[0] != ToolName {
//		printInvalidCommandMessage()
//		return errors.New("invalid command")
//	}
//	commandToExecute := commandParts[1]
//
//	if commandToExecute == "-c" {
//		fileName := commandParts[2]
//		sizeInBytes := findFileSizeInBytes(fileName)
//		fmt.Printf("%d %s\n", sizeInBytes, fileName)
//	} else {
//		printInvalidCommandMessage()
//		return errors.New("invalid command")
//	}
//	return nil
//}
//
//func findFileSizeInBytes(fileName string) int64 {
//	file, err := openFile(fileName)
//	if err != nil {
//		fmt.Println("Error opening file: ", err)
//		return 0
//	}
//	sizeInBytes := calculateSizeInBytes(file)
//	closeFile(file)
//	return sizeInBytes
//}
//
//func calculateSizeInBytes(file *os.File) int64 {
//	var sizeInBytes int64
//	buffer := make([]byte, 1024)
//	for {
//		n, err := file.Read(buffer)
//		if err != nil {
//			if err.Error() != "EOF" {
//				fmt.Println("Error reading file: ", err)
//			}
//			break
//		}
//		sizeInBytes += int64(n)
//	}
//	return sizeInBytes
//}
//
//func openFile(fileName string) (*os.File, error) {
//	fullFilePath := getParentDirectory() + "/data/" + fileName
//	return os.Open(fullFilePath)
//}
//
//func closeFile(file *os.File) {
//	_ = file.Close()
//}
//
//func getParentDirectory() string {
//	currentWorkingDirectory, err := os.Getwd()
//	if err != nil {
//		fmt.Println("Error getting current working directory: ", err)
//		return ""
//	}
//
//	return filepath.Join(currentWorkingDirectory, "..")
//}
//
//func printInvalidCommandMessage() {
//	fmt.Printf("Invalid command, exc. '%s -c <file-name>'\n", ToolName)
//}
