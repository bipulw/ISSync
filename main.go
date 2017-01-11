//Created By Bipul Wagle
//Bipul Wagle. All rights reserved.

package main

import(
  "github.com/bipulw/ISSync/ISSync"
  "fmt"
  "encoding/json"
  "net/http"
)


type User struct{
        Name string      `json:"name,omitempty"`
        Age int          `json:"age,omitempty"`
        Updated_at string`json:"updated_at,omitempty"`
}

func (self User) CreateResponseObj() ISSync.ISSyncModelInterface{
        obj := self
        obj.Name = ""
        return obj
}

func (self User) ReadResponseObj() ISSync.ISSyncModelInterface{
        obj := self
        obj.Age = 0
        return obj
}

func (self User) UpdateResponseObj() ISSync.ISSyncModelInterface{
        obj := self
        obj.Updated_at = ""
        return obj
}

func (self User) DeleteResponseObj() ISSync.ISSyncModelInterface{
        obj := self
        return obj
}

func (self User) Create() bool{
        fmt.Println("create")
        return false
}

func (self User) Update() bool{
        fmt.Println("update")
        return false
}

func (self User) Delete() bool{
        fmt.Println("delete")
        return false
}

func (self User) Read() bool{
        fmt.Println("read")
        return false

}

func (self User) print(){
        fmt.Println("Name       :",self.Name);
        fmt.Println("Age        :",self.Age);
        fmt.Println("Updated at :",self.Updated_at);
}
var _ ISSync.ISSyncModelInterface = (*User)(nil)

// func processRequest(w http.ResponseWriter, r *http.Request){
//   ISSync.processRequest(w,r)
// }
func generateObject(req ISSync.ISSyncRequest) (ISSync.ISSyncModelInterface){
        //example usage
        if req.ModelID == "user"{
                var usrHoldr User
                jsonInBytes := []byte(req.Data)
                if err := json.Unmarshal(jsonInBytes, &usrHoldr); err != nil {
                        return nil
                }
                return usrHoldr
        }
        return nil
}

func main() {
        //http.HandleFunc("/curd", processRequest)
        ISSync.SetObjectGenerator(generateObject)
        http.HandleFunc("/crud", ISSync.ProcessRequest)
}


// func processResquestWorking(jsonString string){
//      var parsed ISSyncRequest
//      var response ISSyncBaseResponse
//``
//      jsonByte := []byte(jsonString)
//      err := json.Unmarshal(jsonByte, &parsed)
//              if err != nil{
//                      response.Error = true
//                      response.Error_msg = "Invalid Requst - Not Json"
//              }else{
//                      obj := generateObject(parsed)
//                      if obj == nil{
//                              response.Error = true
//                              response.Error_msg = "Error No Object Relation Found"
//                      }else{
//                              switch parsed.Crud {
//                              case "create":
//                                      response = create(obj)
//                              case "update":
//                                      response = update(obj)
//                              case "read":
//                                      response = read(obj)
//                              case "delete":
//                                      response = delete(obj)
//                              }
//                      }
//              }
//      fmt.Println(getJson(response))
//
// }


// func olDmain() {
//      fmt.Println("Hello, playground")
//      usr := new(User)
//      usr.Name = "bipul wagle"
//      usr.Age = 21
//      usr.Updated_at = "today"
//      jsonResult, err:= json.Marshal(&usr)
//      fmt.Println("string" , string(jsonResult), "erro ", err)
//      create(usr)
//      update(usr)
//      delete(usr)
//      read(usr)
//      jsonUsr := []byte("{\"age\":21,\"updated_at\":\"today\"}")
//      usrHoldr := new(User);
//      if err := json.Unmarshal(jsonUsr, usrHoldr); err != nil {
//              panic(err)
//      }
//      usrHoldr.print()
//      http.HandleFunc("/curd", processRequest)
//      req := new(ISSyncRequest)
//      req.Crud = "update"
//      req.ModelID = "user"
//      req.Data = "{\"age\":21,\"updated_at\":\"today\"}"
//
//      fmt.Println(getJson(req))
//      processResquestWorking(getJson(req))
//      processResquestWorking("testing")
//      //processResquestWorking()
//      //processRequest();
// }
