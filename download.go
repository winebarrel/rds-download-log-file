package rds_download_log_file

import (
	"fmt"
	"io"
	"strings"

	"github.com/cenkalti/backoff"
	"github.com/winebarrel/goamz/aws"
	"github.com/winebarrel/goamz/rds"
)

func Donwload(region, dbInstanceIdentifier, logFileName string, writer io.Writer) error {
	auth, err := aws.EnvAuth()

	if err != nil {
		return fmt.Errorf("AWS env auth error: %w", err)
	}

	svc, err := rds.New(auth, aws.GetRegion(region))

	if err != nil {
		return fmt.Errorf("RDS client create error: %w", err)
	}

	ticker := backoff.NewTicker(backoff.NewExponentialBackOff())

	for range ticker.C {
		var reader io.ReadCloser
		reader, err = svc.DownloadCompleteDBLogFile(dbInstanceIdentifier, logFileName)

		if err != nil {
			if strings.Contains(err.Error(), "Body: Rate exceeded") {
				continue
			}

			return fmt.Errorf("log download error: %w", err)
		}

		_, err = io.Copy(writer, reader)
		reader.Close()

		if err != nil {
			continue
		}

		ticker.Stop()
		break
	}

	if err != nil {
		return fmt.Errorf("buffer copy error: %w", err)
	}

	return nil
}
