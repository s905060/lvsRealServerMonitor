package main

/*
   Author: Jash Lee
   Date: November 17, 2015
   Purpose: Simple HTTP server that responds the total number of healthy real server in the LVS serving pool
*/

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func getHTTPRunningRealServer() string {
	cmd := "ipvsadm -L -n | grep ':80' | tr -s ' ' '\t'  | awk '{ if ($4 != 0) print $4}' | wc -l"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}

func monitor(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, getHTTPRunningRealServer())
}

func main() {
	http.HandleFunc("/", monitor)
	http.ListenAndServe(":9999", nil)
}
