package get

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jatalocks/opsilon/internal/config"
	"github.com/jatalocks/opsilon/internal/engine"
	"github.com/jatalocks/opsilon/internal/logger"
	"github.com/jatalocks/opsilon/internal/validate"
	"gopkg.in/yaml.v3"
)

func getWorkflows(location config.Location, repo string) *[]engine.Workflow {
	data := []engine.Workflow{}
	logger.Operation("Getting workflows from repo", repo, "in location", location.Path, "type", location.Type)
	if location.Type == "folder" {
		if location.Path[0:1] == "/" {

			err := filepath.Walk(location.Path,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if !info.IsDir() {
						yfile, err := ioutil.ReadFile(path)
						logger.HandleErr(err)
						temp := engine.Workflow{}
						temp.Repo = repo
						err2 := yaml.Unmarshal(yfile, &temp)
						logger.HandleErr(err2)

						data = append(data, temp)
					}

					return nil
				})
			if err != nil {
				logger.Fatal(err)
			}

		}
	}

	// else {

	// 	yfile, err2 := ioutil.ReadFile(path.Join(path.Dir(viper.ConfigFileUsed()), location.Path))
	// 	logger.HandleErr(err2)
	// 	err3 := yaml.Unmarshal(yfile, &data)
	// 	logger.HandleErr(err3)
	// }

	// } else if location.Type == "url" {
	// 	resp, err := http.Get(location.Path)
	// 	logger.HandleErr(err)
	// 	defer resp.Body.Close()
	// 	buf := new(bytes.Buffer)
	// 	buf.ReadFrom(resp.Body)
	// 	err3 := yaml.Unmarshal(buf.Bytes(), &data)
	// 	logger.HandleErr(err3)
	// }
	return &data
}

func appendToWArray(v config.Repo, workflowArray *[]engine.Workflow) error {
	logger.Info("Repository", v.Name)
	w := *getWorkflows(v.Location, v.Name)
	validate.ValidateWorkflows(&w)
	if len(w) == 0 {
		return errors.New("Cannot fetch workflows from repository " + v.Name + " or it is empty.")
	}
	*workflowArray = append(*workflowArray, w...)
	return nil
}
func GetWorkflowsForRepo(repoList []string) ([]engine.Workflow, error) {
	data := config.GetConfig()
	workflowArray := []engine.Workflow{}
	skipRepoCheck := false
	if len(repoList) == 0 {
		skipRepoCheck = true
	}
	for _, v := range data {
		if skipRepoCheck {
			err := appendToWArray(v, &workflowArray)
			if err != nil {
				return nil, err
			}
		} else {
			if config.StringInSlice(v.Name, repoList) {
				err := appendToWArray(v, &workflowArray)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return workflowArray, nil
}
