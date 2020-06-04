package rds_download_log_file

import (
	"io"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/rds"
)

func Donwload(region, dbInstanceIdentifier, logFileName string, writer io.Writer) error {
	auth, err := aws.EnvAuth()

	if err != nil {
		return err
	}

	svc, err := rds.New(auth, aws.GetRegion(region))

	if err != nil {
		return err
	}

	reader, err := svc.DownloadCompleteDBLogFile(dbInstanceIdentifier, logFileName)

	if err != nil {
		return err
	}

	defer reader.Close()

	_, err = io.Copy(writer, reader)

	if err != nil {
		return err
	}

	return nil
}
