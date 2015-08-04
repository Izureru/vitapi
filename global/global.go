// package global

// import (
// 	"os"
// 	"strconv"
// )

// var (
// 	Config ConfigStruct
// )

// type ConfigStruct struct {
// 	Port int
// }

// func LoadConfig() {

// 	Config = ConfigStruct{}

// 	port, _ := strconv.Atoi(os.Getenv("PORT"))

// 	Config.Port = port
// }
