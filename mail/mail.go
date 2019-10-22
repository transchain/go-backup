package mail

import (
    "bytes"
    "errors"
    "os"
    "text/template"
    "time"

    "gopkg.in/gomail.v2"

    "github.com/transchain/go-backup/assets"
    "github.com/transchain/go-backup/types"
)

func Notify(cfg *types.MailConfig, customProjectsErrors types.ProjectsCustomErrors) error {
    if len(cfg.DestEmails) == 0 && cfg.DestEmails[0] == "" {
        return errors.New("no valid dest emails provided in config")
    }

    tmplContent, err := assets.Asset("../tmpl/email-error.tmpl.html")
    if err != nil {
        return err
    }
    hostname, _ := os.Hostname()
    // Generate error mail
    now := time.Now().Format("2006-01-02T15:04:05")
    additionalVars := template.FuncMap{
        "now":      func() string {return now},
        "hostname": func() string {return hostname},
    }
    tmpl, err := template.New("email-error").Funcs(additionalVars).Parse(string(tmplContent))
    if err != nil {
        return err
    }
    var body bytes.Buffer
    err = tmpl.Execute(&body, customProjectsErrors)
    if err != nil {
        return err
    }

    email := gomail.NewMessage()
    email.SetHeader("From", "go backup <"+cfg.Username+">")
    email.SetHeader("To", cfg.DestEmails...)
    email.SetHeader("Subject", "Go Backup Error - "+hostname+" - "+now)
    email.SetBody("text/html", body.String())

    dialer := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
    return dialer.DialAndSend(email)
}
