//Created By Bipul Wagle
//Bipul Wagle. All rights reserved.

package ISSync

import (
        "fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type ISSyncBaseResponse struct{
	Error bool	 `json:"error"`
	Error_msg string `json:"error_msg,omitempty"`
	Data string	 `json:"data,omitempty"`
}

type ISSyncRequest struct{
	Crud string    `json:"crud"`
	ModelID string `json:"model_id"`
	Data string    `json:"data"`
}

type ISSyncModelInterface interface{
	CreateResponseObj() ISSyncModelInterface
	UpdateResponseObj() ISSyncModelInterface
	DeleteResponseObj() ISSyncModelInterface
	ReadResponseObj()   ISSyncModelInterface
	Create() bool
	Delete() bool
	Update() bool
	Read() bool
}
type ISSyncModelGenerator func (ISSyncRequest) ISSyncModelInterface

//func (self User) print(){
//	fmt.Println("Name       :",self.Name);
//	fmt.Println("Age        :",self.Age);
//	fmt.Println("Updated at :",self.Updated_at);
//}

// add all the model you want to support
var generateObject ISSyncModelGenerator

func SetObjectGenerator(req ISSyncModelGenerator){
	generateObject = req
}

func create(obj ISSyncModelInterface) ISSyncBaseResponse{
	error := obj.Create()
	var response ISSyncBaseResponse
	if !error{
		response.Error = false
		response.Data = getJson(obj.CreateResponseObj())
	}else{
		response.Error = true
		response.Error_msg = "Error Creating"
	}
	return response

}

func update(obj ISSyncModelInterface) ISSyncBaseResponse{
	error := obj.Update()
	var response ISSyncBaseResponse
	if !error{
		response.Error = false
		response.Data = getJson(obj.UpdateResponseObj())
	}else{
		response.Error = true
		response.Error_msg = "Error Updating"
	}
	return response

}

func delete(obj ISSyncModelInterface) ISSyncBaseResponse{
	error := obj.Delete()
	var response ISSyncBaseResponse
	if !error{
		response.Error = false
		response.Data = getJson(obj.DeleteResponseObj())
	}else{
		response.Error = true
		response.Error_msg = "Error Updating"
	}
	return response
}

func read(obj ISSyncModelInterface) ISSyncBaseResponse{
	error := obj.Read()
	var response ISSyncBaseResponse
	if !error{
		response.Error = false
		response.Data = getJson(obj.ReadResponseObj())
	}else{
		response.Error = true
		response.Error_msg = "Error Reading"
	}
	return response
}

func getJsonString(obj ISSyncModelInterface) string{
	jsonResult, _ := json.Marshal(obj)
	return(string(jsonResult))
}

func getJson(obj interface{}) string{
	jsonResult, _ := json.Marshal(obj)
	return(string(jsonResult))
}

func ProcessRequest(w http.ResponseWriter, r *http.Request){
	var parsed ISSyncRequest
	var response ISSyncBaseResponse

  	data, err := ioutil.ReadAll(r.Body)
  	if err == nil && data != nil {
      		err = json.Unmarshal(data, &parsed)

      		if err != nil{
        		response.Error = true
			response.Error_msg = "Invalid Requst - Not Json"
		}else{
			obj := generateObject(parsed)
			if obj == nil{
				response.Error = true
				response.Error_msg = "Error No Object Relation Found"
			}else{
				switch parsed.Crud {
				case "create":
					response = create(obj)
				case "update":
					response = update(obj)
				case "read":
					response = read(obj)
				case "delete":
					response = delete(obj)
				}
			}
		}
	}else{
		response.Error = true
		response.Error_msg = "Invalid Request"
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,getJson(response))
}
