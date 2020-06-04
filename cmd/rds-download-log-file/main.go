package main

import (
	"log"
	"os"
	"rds_download_log_file"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags := parseFlags()
	err := rds_download_log_file.Donwload(flags.Region, flags.DBInstanceIdentifier, flags.LogFileName, os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
