package commands

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/pivotalservices/tile-config-generator/metadata"
)

func getProvider(pathToPivotalFile string, pivnet *PivnetConfiguration) metadata.Provider {
	if pivnet != nil {
		return metadata.NewPivnetProvider(pivnet.Token, pivnet.Slug, pivnet.Version, pivnet.Glob)
	} else {
		return metadata.NewFileProvider(pathToPivotalFile)
	}
}
func extractMetadataBytes(pathToPivotalFile string) ([]byte, error) {
	zipReader, err := zip.OpenReader(pathToPivotalFile)
	if err != nil {
		return nil, err
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {
		metadataRegexp := regexp.MustCompile("metadata/.*\\.yml")
		matched := metadataRegexp.MatchString(file.Name)

		if matched {
			metadataFile, err := file.Open()
			contents, err := ioutil.ReadAll(metadataFile)
			if err != nil {
				return nil, err
			}
			return contents, nil
		}
	}
	return nil, errors.New("no metadata file was found in provided .pivotal")
}