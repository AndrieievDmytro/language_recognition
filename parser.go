package main

import (
	"io/ioutil"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func GetLetterProportion(text string) []float64 {
	var fin []float64 = make([]float64, 26)
	var count float64 = 0.0
	for _, ll := range text {
		var l rune = ll
		// Convert to lower case (Upper + 'a' - 'A') = U - 65 + 97
		if ll >= 'A' && ll <= 'Z' {
			l = ll + ' ' // 97-65=32
		}
		if l >= 'a' && l <= 'z' {
			fin[l-'a'] += 1
			count++
		}
	}
	for idx, val := range fin {
		fin[idx] = val / count
	}
	return fin
}

func ReadFile(file string) string {
	data, err := ioutil.ReadFile(file)
	errCheck(err)

	return string(data)
}

func GetFilesList(folder string) []string {

	folders, err := ioutil.ReadDir(folder)
	errCheck(err)
	var folderList []string
	for _, f := range folders {
		if !f.IsDir() {
			folderList = append(folderList, f.Name())
		}
	}
	return folderList
}

func GetFoldersList(folder string) map[string][]string {

	folders, err := ioutil.ReadDir(folder)
	errCheck(err)
	// var folderList [string][]string
	folderList := make(map[string][]string)
	for _, f := range folders {
		// folderList = append(folderList, f.Name())
		folderList[f.Name()] = GetFilesList(folder + "/" + f.Name())
	}
	return folderList
}

func AssignPercName(folder string) []string {
	folders, err := ioutil.ReadDir(folder)
	errCheck(err)

	var percNames []string
	for _, f := range folders {
		if f.IsDir() {
			percNames = append(percNames, f.Name())
		}
	}
	return percNames
}

func FileLetterPropirtion(folder string) map[string]map[string][]float64 {
	folders := GetFoldersList(folder)
	m := make(map[string]map[string][]float64)
	for n, p := range folders {
		fp := make(map[string][]float64)
		for _, f := range p {
			fp[f] = GetLetterProportion(ReadFile(folder + "/" + n + "/" + f))
		}
		m[n] = fp
	}
	return m
}

// func FileLetterPropirtion(folder string) map[string][]float64 {
// 	folders := GetFoldersList(folder)
// 	m := make(map[string][]float64)
// 	for n, p := range folders {
// 		s := ""
// 		for _, f := range p {
// 			s += ReadFile(folder + "/" + n + "/" + f)
// 		}
// 		m[n] = GetLetterProportion(s)
// 	}
// 	return m
// }
