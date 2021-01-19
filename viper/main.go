package main

import (
    "fmt"
    "github.com/spf13/viper"
    "os"
    "strings"
)

const cmdRoot = "core"

func main() {
    fmt.Println("start...")
	config := viper.New()
    config.SetEnvPrefix(cmdRoot)
    
    config.AutomaticEnv()
    replacer := strings.NewReplacer(".", "_")
    config.SetEnvKeyReplacer(replacer)
    config.SetConfigName(cmdRoot)
    config.AddConfigPath("./")
    err := config.ReadInConfig()
    if err != nil {
        fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
        os.Exit(1)
    }

    environment := config.GetBool("security.enabled")
    fmt.Println("security.enabled:", environment)

    fullstate := config.GetString("statetransfer.timeout.fullstate")
    fmt.Println("statetransfer.timeout.fullstate:", fullstate)

    abcdValue := config.GetString("peer.abcd")
    fmt.Println("peer.abcd:", abcdValue)
}