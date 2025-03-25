package changejson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Beauty() {
	file, err := os.Open("task.json")
	if err != nil {
		fmt.Println("file can not open")
		os.Exit(1)
	}

	data, err1 := io.ReadAll(file)
	if err1 != nil {
		fmt.Println("file can not read")
		os.Exit(1)
	}

	var beautyJson bytes.Buffer
	err = json.Indent(&beautyJson, data, "", "  ")
	if err != nil {
		fmt.Println("JSON can not beauty: ", err)
		os.Exit(1)
	}

	err = os.WriteFile("task.json", beautyJson.Bytes(), 0644)
	if err != nil {
		fmt.Println("json file can not write: ", err)
		os.Exit(1)
	}

	fmt.Println("JSON successfully beauty!")
}
