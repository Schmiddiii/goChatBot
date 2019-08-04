package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strings"

	goweb "github.com/Schmiddiii/goWebGui"
)

type resStruct struct {
	KeyStatement string
	Responses    []string
}

func (r *resStruct) getResponse(req string) string {
	if req == r.KeyStatement {
		return r.Responses[rand.Int()%len(r.Responses)]
	}
	return ""
}

type allResponses struct {
	Responses []resStruct
}

func (res *allResponses) getResponse(req string) string {
	for _, r := range res.Responses {
		possibleRes := r.getResponse(req)
		if possibleRes != "" {
			return possibleRes
		}
	}
	return ""
}

func (res *allResponses) addResponse(r resStruct) {
	if res.getResponse(r.KeyStatement) == "" {
		res.Responses = append(res.Responses, r)
	} else {
		for i := range res.Responses {
			if res.Responses[i].KeyStatement == r.KeyStatement {
				for _, posResp := range r.Responses {
					alreadyIn := false
					for _, alreadyResp := range res.Responses[i].Responses {
						if posResp == alreadyResp {
							alreadyIn = true
						}
					}
					if !alreadyIn {
						res.Responses[i].Responses = append(res.Responses[i].Responses, posResp)
					}
				}
			}
		}
	}
}

var possibleResponses allResponses

func main() {
	possibleResponses = readResponses()

	goweb.SetUpCode("3030", handler)
}

func handler(msg goweb.Message) goweb.Message {
	if strings.HasPrefix(msg.Extras[0], "!a ") {
		statement := msg.Extras[0][3:len(msg.Extras[0])]

		parts := strings.Split(statement, "=>")

		if len(parts) != 2 {
			return goweb.Message{ID: "msg", Extras: []string{"That is not valid"}}
		}

		key := strings.ToLower(strings.Trim(parts[0], " ."))

		values := strings.Split(parts[1], ";")
		for k := range values {
			values[k] = strings.Trim(values[k], " .")
		}

		possibleResponses.addResponse(resStruct{KeyStatement: key, Responses: values})

		saveResponses(possibleResponses)

		return goweb.Message{ID: "msg", Extras: []string{"Thank you for teaching me this knowledge"}}
	}
	response := possibleResponses.getResponse(strings.ToLower(msg.Extras[0]))
	if response == "" {
		response = "I dont know what to say to that"
	}
	return goweb.Message{ID: "msg", Extras: []string{response}}
}

func readResponses() allResponses {
	file, err := ioutil.ReadFile("./response.json")
	checkError(err)

	var resStruct []resStruct
	err = json.Unmarshal(file, &resStruct)
	checkError(err)

	allRes := allResponses{Responses: resStruct}
	return allRes

}

func saveResponses(res allResponses) {

	json, err := json.Marshal(res.Responses)
	checkError(err)

	ioutil.WriteFile("./response.json", json, 0644)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
