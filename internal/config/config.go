package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jatalocks/opsilon/internal/engine"
	"github.com/jatalocks/opsilon/internal/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Location struct {
	Path      string `mapstructure:"path" validate:"nonzero"`
	Type      string `mapstructure:"type" validate:"nonzero"`
	Subfolder string `mapstructure:"subfolder,omitempty"`
	Branch    string `mapstructure:"branch,omitempty"`
}

type Repo struct {
	Name        string   `mapstructure:"name" validate:"nonzero"`
	Description string   `mapstructure:"description"`
	Location    Location `mapstructure:"location" validate:"nonzero"`
}

type RepoFile struct {
	Repositories []Repo `mapstructure:"repositories" validate:"nonzero"`
}

var C RepoFile

func PrintRepos(repos []Repo) {
	var data [][]string

	for _, r := range repos {
		row := []string{r.Name, r.Description, r.Location.Path, r.Location.Type, r.Location.Branch, r.Location.Subfolder}
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "Path/URL", "Type", "Branch", "Subfolder"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PrintWorkflows(workflows []engine.Workflow) {
	var data [][]string

	for _, r := range workflows {
		out := ""
		for _, v := range r.Input {
			out += fmt.Sprintf("%v,", v.Name)
		}
		images := []string{r.Image}
		for _, v := range r.Stages {
			if !StringInSlice(v.Image, images) {
				images = append(images, v.Image)
			}
		}

		row := []string{r.Repo, r.ID, r.Description, TrimSuffix(strings.Join(images, ","), ","), TrimSuffix(out, ","), strconv.Itoa(len(r.Stages))}
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Repository", "ID", "Description", "Images Used", "Inputs", "Stage Count"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func GetConfig() []Repo {
	err2 := viper.Unmarshal(&C)
	logger.HandleErr(err2)
	return C.Repositories
}
func GetConfigFile() *RepoFile {
	err2 := viper.Unmarshal(&C)
	logger.HandleErr(err2)
	return &C
}

func SaveToConfig(r RepoFile) {
	file, err := os.OpenFile(viper.ConfigFileUsed(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(r)
	if err != nil {
		log.Fatalf("error encoding: %v", err)
	}
}

func GetRepoList() []string {
	temp := []string{}
	err2 := viper.Unmarshal(&C)
	logger.HandleErr(err2)
	for _, r := range C.Repositories {
		temp = append(temp, r.Name)
	}
	return temp
}
