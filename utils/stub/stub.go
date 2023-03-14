package stub

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"template/utils/modname"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type StubDetails struct {
	Name, FileName, Destination string
	Values                      map[string]string
}

type file struct {
	name string
	path string
}

var excludedPath = []string{
	".git",
	".github",
	"src",
}

func MakeStubs() error {
	err := os.RemoveAll("./stubs/core")
	if err != nil {
		return err
	}

	var files []file
	err = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}

		for _, v := range excludedPath {
			if strings.Contains(path, v) {
				return nil
			}
		}

		if !info.IsDir() {
			f := file{
				name: info.Name(),
				path: path,
			}
			files = append(files, f)
		}
		return nil
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	fmt.Printf("file: %v\n", files)

	for _, f := range files {
		destination := strings.Split(f.path, "/")
		fmt.Println("Debug")
		stubDestination := fmt.Sprintf("./stubs/core/%s", strings.Join(destination[0:len(destination)-1], "/"))
		fmt.Println(destination)

		input, err := os.ReadFile(f.path)
		if err != nil {
			log.Fatal(err.Error())
		}

		output := bytes.Replace(input, []byte("template/"), []byte("{{.Module}}/"), -1)
		output = bytes.Replace(output, []byte("module template"), []byte("module {{.Module}}"), -1)

		if _, err := os.Stat(stubDestination); os.IsNotExist(err) {
			os.MkdirAll(stubDestination, 0755)
		}

		if err := os.WriteFile(fmt.Sprintf("./stubs/core/%s.stub", f.path), output, 0755); err != nil {
			log.Fatal(err.Error())
		}

		// s := StubDetails{
		// 	Name:        f.path,
		// 	FileName:    strings.ReplaceAll(f.name, ".stub", ""),
		// 	Destination: fmt.Sprintf("./%s/", strings.Join(destination[1:len(destination)-1], "/")),
		// 	Values: map[string]string{
		// 		"Model": "User",
		// 	},
		// }

		// contentsBuff, err := os.ReadFile(s.Name)
		// if err != nil {
		// 	log.Fatalf("Unable to read file: %s", s.Name)
		// }

		// if _, err := os.Stat(s.Destination); os.IsNotExist(err) {
		// 	os.MkdirAll(s.Destination, 0755)
		// }

		// f, err := os.OpenFile(s.Destination+s.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		// if err != nil {
		// 	log.Fatalf("Unable to open file: %s", s.FileName)
		// }
		// defer f.Close()

		// template, err := template.New(s.FileName).Parse(string(contentsBuff))
		// if err != nil {
		// 	log.Fatalf("Unable to parse template: %s", s.Name)
		// }
		// template.Execute(f, s.Values)

	}
	return nil
}

func Stubs(module string) {
	var files []file
	err := filepath.Walk("./stubs/core",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size(), info.IsDir())
			if !info.IsDir() {
				f := file{
					name: info.Name(),
					path: path,
				}
				files = append(files, f)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("file: %v\n", files)

	for _, f := range files {
		destination := strings.Split(f.path, "/")

		s := StubDetails{
			Name:        f.path,
			FileName:    strings.ReplaceAll(f.name, ".stub", ""),
			Destination: fmt.Sprintf("./src/%s/%s/", module, strings.Join(destination[2:len(destination)-1], "/")),
			Values: map[string]string{
				"Module": module,
			},
		}

		contentsBuff, err := os.ReadFile(s.Name)
		if err != nil {
			log.Fatalf("Unable to read file: %s", s.Name)
		}

		if _, err := os.Stat(s.Destination); os.IsNotExist(err) {
			os.MkdirAll(s.Destination, 0755)
		}

		f, err := os.OpenFile(s.Destination+s.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalf("Unable to open file: %s", s.FileName)
		}
		defer f.Close()

		template, err := template.New(s.FileName).Parse(string(contentsBuff))
		if err != nil {
			log.Fatalf("Unable to parse template: %s", s.Name)
		}
		template.Execute(f, s.Values)

	}
}

func TemplateStub(templateType string, templateName string, name string) {
	contentsBuff, err := os.ReadFile(fmt.Sprintf("./stubs/template/%s/%s.go.stuba", templateType, templateName))
	if err != nil {
		log.Fatalf("Unable to read file: %s", name+".go")
	}

	f, err := os.OpenFile("./internal/"+templateType+"/"+name+".go", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("Unable to open file: %s", name+".go")
	}
	defer f.Close()

	template, err := template.New("./internal/" + templateType + "/" + name + ".go").Parse(string(contentsBuff))
	if err != nil {
		log.Fatalf("Unable to parse template: %s", name+".go")
	}
	template.Execute(f, map[string]string{
		"Module":    modname.GetModuleName(),
		"Name":      cases.Title(language.English).String(strings.ToLower(name)),
		"NameLower": strings.ToLower(name),
		"NameFirst": strings.ToLower(string([]rune(name)[0:1])),
	})
	log.Println(templateType + " created")
}
