package tools

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

// log info from client connection
// log when server starts listening
func LogPortListening(port string) {
	LogInfoIntoStdout(fmt.Sprintf("%s %s %s", "port", port, "is listening"))
}

func LogOKInfo(method string, operation string) {
	LogInfoIntoStdout(fmt.Sprintf("%s | %d | %s", method, 200, operation))
}

func LogNoFound(method string) {
	LogInfoIntoStdout(fmt.Sprintf("%s | %d", method, 400))
}

func LogInfoIntoStdout(message string) {
	log.SetOutput(os.Stdout)
	log.Println("[Todos]  " + message)
}

// to encrypt password by MD5
func MD5Encryption(text string) string {
	hash := md5.New()
	io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
