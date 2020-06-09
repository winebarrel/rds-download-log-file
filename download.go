package rds_download_log_file

import (
	"fmt"
	"io"
	"strings"

	"github.com/cenkalti/backoff"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/rds"
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

	for _ = range ticker.C {
		var reader io.ReadCloser
		reader, err = svc.DownloadCompleteDBLogFile(dbInstanceIdentifier, logFileName)

		if err != nil {
			if strings.Contains(err.Error(), "Body: Rate exceeded") {
				continue
			}

			return fmt.Errorf("log download error: %w", err)
		}

		defer reader.Close()
		_, err = io.Copy(writer, reader)

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
