package main

import (
	"bytes"
    "fmt"
    "net/http"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	
	u,_ := url.Parse(r.RequestURI)
	
	if u.RawQuery!="" {
		
		q := u.Query()
		//fmt.Println("%s",q["cmd"][0])
		//fileName := "aaa.go"
		dstFile,err := os.Create("../test/test.go")
		if err!=nil{
			fmt.Println(err.Error())    
			return
		}   

		defer dstFile.Close()
		for _,value:=range q{ 
			//fmt.Println(key)
			//fmt.Println(value)
			dstFile.WriteString(value[0])
		}
		
		dstFile.Close()
		
		cmd := exec.Command("go", "install", "test")
		err = cmd.Run()
		
		if(nil != err) {
			fmt.Println(err)
		}

		run := exec.Command("test.exe")
		var out bytes.Buffer
		run.Stdout = &out
		err = run.Run()
		
		if(nil != err) {
			fmt.Println(err)
		}
		
		w.Write([]byte(out.String()))
		

	} else {
		index, _ := ioutil.ReadFile("html/index.html")
		w.Write([]byte(index))
	}	
	
}

func main() {
    http.HandleFunc("/index", indexHandler)
    http.ListenAndServe(":8080", nil)
}