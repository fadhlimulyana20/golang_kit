package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// type StubDetails struct {
// 	Name, FileName, Destination string
// 	Values                      map[string]string
// }

// type file struct {
// 	name string
// 	path string
// }

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)

	// var files []file
	// err := filepath.Walk("./stubs",
	// 	func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// fmt.Println(path, info.Size(), info.IsDir())
	// 		if !info.IsDir() {
	// 			f := file{
	// 				name: info.Name(),
	// 				path: path,
	// 			}
	// 			files = append(files, f)
	// 		}
	// 		return nil
	// 	})
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Printf("file: %v\n", files)

	// for _, f := range files {
	// 	destination := strings.Split(f.path, "/")

	// 	s := StubDetails{
	// 		Name:        f.path,
	// 		FileName:    strings.ReplaceAll(f.name, ".stub", ""),
	// 		Destination: fmt.Sprintf("./%s/", strings.Join(destination[1:len(destination)-1], "/")),
	// 		Values: map[string]string{
	// 			"Model": "User",
	// 		},
	// 	}

	// 	contentsBuff, err := os.ReadFile(s.Name)
	// 	if err != nil {
	// 		log.Fatalf("Unable to read file: %s", s.Name)
	// 	}

	// 	if _, err := os.Stat(s.Destination); os.IsNotExist(err) {
	// 		os.MkdirAll(s.Destination, 0755)
	// 	}

	// 	f, err := os.OpenFile(s.Destination+s.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	// 	if err != nil {
	// 		log.Fatalf("Unable to open file: %s", s.FileName)
	// 	}
	// 	defer f.Close()

	// 	template, err := template.New(s.FileName).Parse(string(contentsBuff))
	// 	if err != nil {
	// 		log.Fatalf("Unable to parse template: %s", s.Name)
	// 	}
	// 	template.Execute(f, s.Values)

	// }
}
