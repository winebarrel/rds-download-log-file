package main

import (
	"flag"
	"fmt"
	"os"
)

var version string

type Flags struct {
	DBInstanceIdentifier string
	LogFileName          string
	Region               string
}

func parseFlags() (flags *Flags) {
	flags = &Flags{}

	flag.StringVar(&flags.DBInstanceIdentifier, "db-instance-identifier", "", "DB instance identifier")
	flag.StringVar(&flags.LogFileName, "log-file-name", "", "log file name")
	flag.StringVar(&flags.Region, "region", getDefaultRegion(), "AWS region")
	argVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if flag.NFlag() == 0 {
		printUsageAndExit()
	}

	if *argVersion {
		printVersionAndEixt()
	}

	if flags.DBInstanceIdentifier == "" {
		printErrorAndExit("'-db-instance-identifier' is required")
	}

	if flags.LogFileName == "" {
		printErrorAndExit("'-log-file-name' is required")
	}

	if flags.Region == "" {
		printErrorAndExit("'-region' is required")
	}

	return
}

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func printVersionAndEixt() {
	fmt.Fprintln(os.Stderr, version)
	os.Exit(0)
}

func printErrorAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func getDefaultRegion() string {
	region := os.Getenv("AWS_DEFAULT_REGION")

	if region == "" {
		region = os.Getenv("AWS_REGION")
	}

	return region
}
