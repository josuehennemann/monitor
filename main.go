package main

import (
	"fmt"
	//"net/http"
	"os"
	"os/exec"
	"bytes"
)

func main() {
	/*fmt.Println(commandLine("cat", "/etc/*-release"))
	http.HandleFunc("/", handler)
		
	auth := &http.Server{Addr: ":80"}
	err := auth.ListenAndServe() 
	if err != nil {
		fmt.Println("Can't start service http: %s\n", err.Error())
		os.Exit(0)
	}*/

}

/*func handler(w http.ResponseWriter, r *http.Request) {
		
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}*/

func commandLine(command string, args ...string) (string, error){
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd :=  exec.Command("ls", "-la")
	//cmd.Stdout = &stdOut
	//cmd.Stderr = &stdErr
	tmp, err := cmd.Output()
	if err != nil {
		return "a", err
	}
	return fmt.Sprintf("%s-%s-%s",stdOut.Bytes(),stdErr.Bytes(), tmp), nil
}

