package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func reloadProcess(command string, process []string, c Config) {
	if len(process) != 0 {
		fmt.Println("Error: reload accepts no arguments")
		fmt.Println("reload 		Restart the remote taskmasterd.")
	} else if !cli {
		fmt.Println("Really restart the remote taskmasterd process y/N?")
		buf := bufio.NewReader(os.Stdin)
		response, err := buf.ReadByte()
		if err != nil {
			return
		}
		switch string(response) {
		case "y":
			request(command, process, c)
		case "n":
			break
		case "N":
			break
		default:
			reloadProcess(command, process, c)
		}
	} else {
		request(command, process, c)
	}
}

func listProgs(command string, c Config) {
	port := strconv.Itoa(int(c.Port))
	client := &http.Client{}
	url := c.Serverurl + ":" + port + command
	req, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	if req.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		list := parserRequest(body)
		log.Println("=> " + list)
	}
}

func getInfo(command string, process []string, c Config, i int) {
	port := strconv.Itoa(int(c.Port))
	client := &http.Client{}
	url := c.Serverurl + ":" + port + "/" + command
	if len(process) != 0 {
		url = c.Serverurl + ":" + port + "/" + command + "/" + process[i]
	}
	req, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()
	if req.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		bodystring := parserRequest(body)
		fmt.Println(bodystring)
	}
}

func request(command string, process []string, c Config) {
	if len(process) > 1 {
		for i := 0; i < len(process); i++ {
			getInfo(command, process, c, i)
		}
	} else {
		getInfo(command, process, c, 0)
	}
}
