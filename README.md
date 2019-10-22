# Go Backup

[![Build Status](https://travis-ci.org/transchain/go-backup.svg)](https://travis-ci.org/transchain/go-backup)

This repository contains a Go application to execute the backup command
on multiple projects and copy the generated backup files into global
backup folders or send it to a sftp server. If an error occurs, an email is sent.

## Content

- [Requirements](#requirements)
  - [Tested versions](#tested-versions)
- [Build the project](#build-the-project)
- [Usage](#usage)

## Requirements

In development, you'll need to install Go on your machine.

### Tested versions

- [bash](https://www.gnu.org/software/bash/) (Tested: v4.4.19)

Development:

- [go](https://golang.org/doc/install) (Tested: v1.12.*)

## Build the project

In order to copy the default configuration, please run the following
script:

```bash
./bin/install_all.sh
```

In development, you can build the project with the following command:

```bash
GO111MODULE=off go get github.com/go-bindata/go-bindata/...
go generate ./...
go build -o build/go-backup ./cmd
```

## Usage

You can directly run the binary (no arguments defaults the config to ./config.yml).
```bash
./build/go-backup run
```
Or specify the config file path :
```bash
./build/go-backup run -c $(pwd)/config.yml
```

And you can run this project with a real crontab command as well:

```bash
crontab -e
```

And add this lines at the end:

```bash
* * * * * {{absolute_path_to_go_backup_dir}}/build/go-backup run -c {{absolute_path_to_go_backup_dir}}/config.yml 2>> {{absolute_path_to_go_backup_dir}}/error.log
```

You can redirect the stdout to a log file to spot the configuration errors.
