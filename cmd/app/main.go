package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "gopkg.in/yaml.v2"
    "os/exec"
    "strconv"
)

type Config struct {
    Qemu_system_rules Parameters `yaml:"qemu_system_rules"`
}

type Parameters struct {
    Scale string `yaml:"scale"`
    Soc []string `yaml:"soc"`
    Cpu []string `yaml:"cpu"`
    Sys_options []Options `yaml:"options"`
}

type Options struct {
    Machine string `yaml:"machine"`
    Gdb string `yaml:"gdb"`
    Nongraphic string `yaml:"nongraphic"`
    Kernel string `yaml:"kernel"`
}

func getYmlFile() *Parameters {
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

    return &config.Qemu_system_rules
}

func execute(Params Parameters) {
    
    fmt.Printf("%#v\n", Params)

    app := "echo"

    scale, err := strconv.Atoi(Params.Scale)
    if err != nil {
        panic(err)
    }

    for i := 0; i < scale; i++ {
        fmt.Println(i)
        cmd := exec.Command(app, Params.Scale)
        stdout, err := cmd.Output()

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        fmt.Println(string(stdout))
    }
}

func main() {
    file := getYmlFile()
    execute(*file)
}