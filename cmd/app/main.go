package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Qemu_system_rules map[string]interface{}
}

type Parameters struct {
    Soc string
    Cpu string
    Options string
}

func main() {
    filename, _ := filepath.Abs("./example.yml")
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var config Config

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Value: %#v\n", config.Qemu_system_rules)
}