package backup

import (
    "fmt"
    "io/ioutil"
    "os/exec"
    "path/filepath"
    "sort"
    "time"

    "github.com/robfig/cron"

    "github.com/transchain/go-backup/provider"
    "github.com/transchain/go-backup/types"
)

type Backuper struct {
    Providers map[string]provider.Provider
    Projects  map[string]*types.ProjectConfig
}

func NewBackuper(providers map[string]provider.Provider, projects map[string]*types.ProjectConfig) *Backuper {
    return &Backuper{
        Providers: providers,
        Projects:  projects,
    }
}

func (b *Backuper) Run() types.ProjectsCustomErrors {
    projectsCustomErrors := types.ProjectsCustomErrors{}
    now := time.Now()
    currentTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.Local)
    testTime := currentTime.Add(-1 * time.Minute)
    cronParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

    for name, project := range b.Projects {
        schedule, err := cronParser.Parse(project.CronSpec)
        if err != nil {
            projectsCustomErrors.AddErrors(name, types.NewCustomError("Unable to parse cron spec", err.Error()))
            continue
        }
        // Test if we are going to execute this task
        nextTime := schedule.Next(testTime)
        if currentTime.Equal(nextTime) {
            // Do the backup and copy the files to te specified destinations
            filePaths, err := DoBackup(project)
            if err != nil {
                projectsCustomErrors.AddErrors(name, err)
                continue
            }
            for _, providerName := range project.Destinations {
                if prov, ok := b.Providers[providerName]; ok {
                    customErrors := SendFiles(prov, filePaths, name)
                    if customErrors != nil {
                        projectsCustomErrors.AddErrors(name, customErrors...)
                    }
                } else {
                    projectsCustomErrors.AddErrors(name, types.NewCustomError(fmt.Sprintf("Provider %s not found", providerName), ""))
                }
            }
        }
    }
    return projectsCustomErrors
}

func DoBackup(project *types.ProjectConfig) ([]string, *types.CustomError) {
    // Try to execute the task
    cmd := exec.Command(project.Task)
    cmd.Dir = project.Path
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, types.NewCustomError(fmt.Sprintf("Unable to execute task %s", project.Task), string(output))
    }

    // Try to sort the files in the project backup folder
    fullProjectBackupFolder := filepath.Join(project.Path, project.BackupFolder)
    files, err := ioutil.ReadDir(fullProjectBackupFolder)
    if err != nil {
        return nil, types.NewCustomError(fmt.Sprintf("Unable to read directory %s", fullProjectBackupFolder), err.Error())
    }
    sort.Slice(files, func(i, j int) bool {
        return files[i].ModTime().Unix() > files[j].ModTime().Unix()
    })

    // Return the file paths list
    var filePaths []string
    for i := 0; i < project.FileNumber; i++ {
        filePaths = append(filePaths, filepath.Join(fullProjectBackupFolder, files[i].Name()))
    }

    return filePaths, nil
}

func SendFiles(prov provider.Provider, filePaths []string, projectName string) []*types.CustomError {
    var customErrors []*types.CustomError
    for _, filePath := range filePaths {
        err := prov.SendFile(filePath, projectName)
        if err != nil {
            customErrors = append(customErrors, types.NewCustomError(fmt.Sprintf("Unable to send file %s", filePath), err.Error()))
        }
    }
    return customErrors
}
