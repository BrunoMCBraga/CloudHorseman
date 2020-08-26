package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const tempPrefix string = "ch-"

type FileSplitterFunction func(string, int) (map[string]bool, error)

var FileSplitterMappings map[string]FileSplitterFunction = map[string]FileSplitterFunction{
	"SPLITINPUT":     SplitFileSlitInput,
	"MULTIPLYINPUT":  SplitFiletMultiplyInput,
	"DUPLICATEINPUT": SplitFiletDuplicateInput}

func FileExists(filePath string) bool {

	_, statError := os.Stat(filePath)

	if statError == nil {
		return true

	} else {
		fmt.Println("|CloudHorseman->util->file_util->FileExists->os.Stat:" + statError.Error() + "|")
		return false
	}
}

func ReadFile(filePath string) ([]byte, error) {

	openResult, openError := os.Open(filePath)
	if openError != nil {
		return nil, errors.New("|CloudHorseman->util->file_util->ReadFile->os.Open:" + openError.Error() + "|")
	}

	defer openResult.Close()

	readAllResult, readAllError := ioutil.ReadAll(openResult)
	if readAllError != nil {
		return nil, errors.New("|CloudHorseman->util->file_util->ReadFile->ioutil.ReadAll:" + readAllError.Error() + "|")
	}

	return readAllResult, nil
}

func WriteStringToFile(fileContent string, filePath string) error {

	openFileResult, openFileError := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if openFileError != nil {
		return errors.New("|CloudHorseman->util->file_util->WriteStringToFile->os.OpenFile:" + openFileError.Error() + "|")
	}

	defer openFileResult.Close()

	_, writeStringError := openFileResult.WriteString(fileContent)
	if writeStringError != nil {
		return errors.New("|CloudHorseman->util->file_util->WriteStringToFile->WriteString:" + writeStringError.Error() + "|")
	}

	return nil
}

func CreateTemporaryFile() (string, error) {

	tempFileResult, tempFileError := ioutil.TempFile(os.TempDir(), tempPrefix+GetRandomTimestampString())
	if tempFileError != nil {
		return "", errors.New("|CloudHorseman->util->file_util->CreateTemporaryFile->ioutil.TempFile:" + tempFileError.Error() + "|")
	} else {
		return tempFileResult.Name(), nil
	}

}

func splitSliceIntoSubslices(slice []string, numberOfSlices int) []interface{} {

	subSlices := make([]interface{}, 0)
	sliceLength := len(slice) / numberOfSlices //3/2 = 1
	sizeOfLastSlice := len(slice) % numberOfSlices
	lastSubsliceIndex := 0

	if numberOfSlices > len(slice) {
		for _, sliceElement := range slice {
			subSlice := make([]string, 0)
			subSlice = append(subSlice, sliceElement)
			subSlices = append(subSlices, subSlice)
		}
	} else {

		for lastSubsliceIndex = 0; lastSubsliceIndex < numberOfSlices; lastSubsliceIndex++ {
			subSlice := make([]string, 0)
			for i := 0; i < sliceLength; i++ {
				subSlice = append(subSlice, slice[lastSubsliceIndex*sliceLength+i])
			}
			subSlices = append(subSlices, subSlice)
		}

		remainingSlice := slice[lastSubsliceIndex:]

		if sizeOfLastSlice > 0 {
			for remainingSliceIndex, remainingSliceElement := range remainingSlice {
				subSlices[remainingSliceIndex] = append(subSlices[remainingSliceIndex].([]string), remainingSliceElement)
			}
		}
	}

	return subSlices
}

//numberOfSubFiles should be the number of Docker containers.
func SplitFileSlitInput(filePath string, numberOfSubFiles int) (map[string]bool, error) {

	readFileResult, readFileError := ioutil.ReadFile(filePath)
	tempFiles := make(map[string]bool, 0)

	if readFileError != nil {
		return tempFiles, errors.New("|CloudHorseman->util->file_util->SplitFileSlitInput->ioutil.ReadFile:" + readFileError.Error() + "|")
	}

	splittedByLines := strings.Split(strings.Trim(string(readFileResult), "\n"), "\n")

	splittedByLinesAndGrouped := splitSliceIntoSubslices(splittedByLines, numberOfSubFiles)

	InitRandomGeneratorSeed()

	for _, listOfStringsToWriteToFile := range splittedByLinesAndGrouped {
		tempFileResult, tempFileError := ioutil.TempFile(os.TempDir(), tempPrefix+GetRandomTimestampString())

		if tempFileError != nil {
			for tempFile, _ := range tempFiles {
				os.Remove(tempFile)
			}

			return tempFiles, errors.New("|CloudHorseman->util->file_util->SplitFileSlitInput->ioutil.TempFile:" + tempFileError.Error() + "|")
		}

		dataToWrite := ""
		for _, lineToWrite := range listOfStringsToWriteToFile.([]string) {
			dataToWrite += (lineToWrite + "\n")
		}

		tempFileResult.Write([]byte(dataToWrite))
		tempFileResult.Close()

		tempFiles[tempFileResult.Name()] = true
	}

	return tempFiles, nil
}

func SplitFiletDuplicateInput(filePath string, numberOfSubFiles int) (map[string]bool, error) {

	readFileResult, readFileError := ioutil.ReadFile(filePath)
	tempFiles := make(map[string]bool, 0)

	if readFileError != nil {
		return tempFiles, errors.New("|CloudHorseman->util->file_util->SplitFiletDuplicateInput->ioutil.ReadFile:" + readFileError.Error() + "|")
	}

	InitRandomGeneratorSeed()

	for i := 1; i <= numberOfSubFiles; i++ {
		tempFileResult, tempFileError := ioutil.TempFile(os.TempDir(), tempPrefix+GetRandomTimestampString())

		if tempFileError != nil {
			for tempFile, _ := range tempFiles {
				os.Remove(tempFile)
			}
			return tempFiles, errors.New("|CloudHorseman->util->file_util->SplitFiletDuplicateInput->ioutil.TempFile:" + tempFileError.Error() + "|")
		}

		tempFileResult.Write(readFileResult)
		tempFileResult.Close()
		tempFiles[tempFileResult.Name()] = true
	}

	return tempFiles, nil
}

func SplitFiletMultiplyInput(filePath string, numberOfSubFiles int) (map[string]bool, error) {

	readFileResult, readFileError := ioutil.ReadFile(filePath)
	tempFiles := make(map[string]bool, 0)

	if readFileError != nil {
		return tempFiles, errors.New("|CloudHorseman->util->file_util->SplitFiletMultiplyInput->ioutil.ReadFile:" + readFileError.Error() + "|")
	}

	splittedByLines := strings.Split(strings.Trim(string(readFileResult), "\n"), "\n")
	InitRandomGeneratorSeed()

	for _, stringToWriteToFile := range splittedByLines {

		for i := 1; i <= numberOfSubFiles; i++ {
			tempFileResult, tempFileError := ioutil.TempFile(os.TempDir(), tempPrefix+GetRandomTimestampString())

			if tempFileError != nil {
				for tempFile, _ := range tempFiles {
					os.Remove(tempFile)
				}

				return tempFiles, errors.New("|CloudHorseman->util->file_util->SplitFiletMultiplyInput->ioutil.TempFile:" + tempFileError.Error() + "|")
			}
			tempFileResult.Write([]byte(stringToWriteToFile))
			tempFileResult.Close()
			tempFiles[tempFileResult.Name()] = true
		}
	}

	return tempFiles, nil
}

func SplitFileBasedOnSplitType(filePath string, replicas int, splitType string) (map[string]bool, error) {

	if splitType == "SPLITINPUT" || splitType == "MULTIPLYINPUT" || splitType == "DUPLICATEINPUT" {

		fileSplitterFunction := FileSplitterMappings[splitType]
		fileSplitterFunctionResult, fileSplitterFunctionError := fileSplitterFunction(filePath, replicas)
		if fileSplitterFunctionError != nil {
			return nil, errors.New("|CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->fileSplitterFunction:" + fileSplitterFunctionError.Error() + "|")
		} else if fileSplitterFunctionResult == nil {
			return nil, errors.New("|CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->fileSplitterFunction: returned null commands structure.|")
		}

		return fileSplitterFunctionResult, nil

	} else {
		return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->Invalid WORKLOADSPLIT argument.|")
	}

	/*
		switch splitType {
		//I think this switch case is not necessary.....i can index directly...
		case "SPLITINPUT":

			splitInputFileSplitterFunction := FileSplitterMappings["SPLITINPUT"]
			splitInputFileSplitterFunctionResult, splitInputFileSplitterFunctionError := splitInputFileSplitterFunction(filePath, replicas)
			if splitInputFileSplitterFunctionError != nil {
				return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->splitInputFileSplitterFunction:" + splitInputFileSplitterFunctionError.Error() + "|")
			} else if splitInputFileSplitterFunctionResult == nil {
				return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->splitInputFileSplitterFunction: returned null commands structure.|")
			}

			return splitInputFileSplitterFunctionResult, nil

		case "MULTIPLYINPUT":
			multiplyInputFileSplitterFunction := FileSplitterMappings["MULTIPLYINPUT"]
			multiplyInputFileSplitterFunctionResult, multiplyInputFileSplitterFunctionError := multiplyInputFileSplitterFunction(filePath, replicas)
			if multiplyInputFileSplitterFunctionError != nil {
				return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->multiplyInputFileSplitterFunction:" + multiplyInputFileSplitterFunctionError.Error() + "|")
			} else if multiplyInputFileSplitterFunctionResult == nil {
				return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->multiplyInputFileSplitterFunction: returned null commands structure.|")
			}

			return multiplyInputFileSplitterFunctionResult, nil

		case "DUPLICATEINPUT":
			duplicateInputFileSplitterFunction := FileSplitterMappings["DUPLICATEINPUT"]
			duplicateInputFileSplitterFunctionResult, duplicateInputFileSplitterFunctionError := duplicateInputFileSplitterFunction(filePath, replicas)
			if duplicateInputFileSplitterFunctionError != nil {
				return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->duplicateInputFileSplitterFunction:" + duplicateInputFileSplitterFunctionError.Error() + "|")
			} else if duplicateInputFileSplitterFunctionResult == nil {
				return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->duplicateInputFileSplitterFunction: returned null commands structure.|")
			}

			return duplicateInputFileSplitterFunctionResult, nil

		default:
			return nil, errors.New("|" + "CloudHorseman->util->heavy_work->SplitFileBasedOnSplitType->Invalid WORKLOADSPLIT argument.|")

		}*/

}
