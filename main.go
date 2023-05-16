package main

import "autoCreate/cmd"

func main() {
	cmd.Execute()
}

// import (
// 	"demo/env"
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	var myClient = env.GetClient()
// 	var myEnv = env.GetEnv()

// 	//命令行工具
// 	var command = flag.String("cmd", "", "Command to run (init, build, review, run)")
// 	flag.Parse()

// 	switch *command {
// 	case "init":
// 		myEnv.InitEnvFile(myClient.Provider)
// 	case "build":
// 		myEnv.WriteInstanceFormatFile()
// 	case "review":
// 		fmt.Println("Running 'review' function...")
// 	case "run":
// 		fmt.Println("Running 'run' function...")
// 	default:
// 		usage()
// 	}
// }

// // 帮助信息
// func usage() {
// 	println("\nUsage:\n  -cmd string\n        Command to run (init, build, review, run)\n")
// }
