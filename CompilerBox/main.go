package main

import (
	"context"
	"log"
	"fmt"

	proto "github.com/TechMaster/LearnGo/CompilerBox/proto"
	"github.com/micro/go-micro"
	"github.com/teris-io/shortid"
	"os"
	"bufio"
	"os/exec"
)

type Compiler struct{}
func saveToFile(req *proto.CompileRequest) (string, error) {
	//1. Write to file with unique id.req.
	//2. Call appropriate function to compile file and write result to file. Dùng lệnh Switch kiểm tra ngôn ngữ
	fileid, _ := shortid.Generate()
	path := "tmp/" + fileid + "." + req.Language
	file, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range req.SourceCode {
		fmt.Fprintln(w, line)
	}

	return path, w.Flush()
}

func (g *Compiler) Run(ctx context.Context, req *proto.CompileRequest, rsp *proto.CompileResponse) error {
	path, err := saveToFile(req)
	if err != nil {
		return err
	}

	switch req.Language  {
	case "go":
		compileGo(path)
	case "js":
		compileNode(path)
	case "py":
	}
	return nil
}

func compileGo(path string) error {
	go func(){
		out, err := exec.Command("go", "run", path).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}()
	return nil
}

func compileNode(path string) error {
	go func(){
		out, err := exec.Command("node", path).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}()
	return nil
}


/*func compilePython(fileID string) error {

}

func compileC(fileID string) error {

}*/

func main() {
	service := micro.NewService(
		micro.Name("compiler"),
	)

	service.Init()
	proto.RegisterCompileHandler(service.Server(), new(Compiler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
