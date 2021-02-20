package main

import ("fmt";
		"encoding/json";
		"os";
		"bufio"
		)

type (
	Action struct {
		Action string `json:"action"`
		ObjName string `json:"object"`
	}
	Teacher struct {
		ID string  `json:"id"`
		Salary float64 `json:"salary"`
		Subject string `json:"subject"`
		Classroom []string `json:"classroom"`
		Person struct {
			Name string `json:"name"`
			Surname string `json:"surname"`
			PersonalCode string `json:"personalCode"`
		} `json:"person"`
	}
	UpdateTeacher struct {
		T Teacher `json:"data"`
	}
	CreateTeacher struct {
		T Teacher `json:"data"`
	}
	DeleteTeacher struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	ReadTeacher struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
)

type (
	DefinedAction interface {
		GetFromJSON([]byte)
		Process()
	}
	GeneralObject interface {
		GetCreateAction() DefinedAction
		GetUpdateAction() DefinedAction
		GetReadAction() DefinedAction
		GetDeleteAction() DefinedAction
		Read(str string)bool
		Print()
	}
)

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

func (action CreateTeacher) Process(){
	fmt.Println("Create Teacher")
	arriPerson = append(arriPerson, &action.T)
	PrintAll(arriPerson)
}

func (action UpdateTeacher) Process() {
	fmt.Println("Update Teachers")
	for i:=0;i<len(arriPerson);i++{	
		if arriPerson[i].Read(action.T.ID) {
			arriPerson[i] = &action.T
		}
	}
	PrintAll(arriPerson)
}

func (action ReadTeacher) Process() {
	fmt.Println("Read teacher", action.Data.ID)
	for i:=0;i<len(arriPerson);i++{
		if arriPerson[i].Read(action.Data.ID) {
			arriPerson[i].Print()
		}
	}
}

func (action DeleteTeacher) Process() {
	fmt.Println("Teacher deleted", action.Data.ID)
	for i:=0;i<len(arriPerson);i++{
		if arriPerson[i].Read(action.Data.ID) {
			copy(arriPerson[i:], arriPerson[i+1:])
			arriPerson[len(arriPerson)-1] = nil
			arriPerson = arriPerson[:len(arriPerson)-1]
		}
	}
	PrintAll(arriPerson)
}

func PrintAll(arriPerson []GeneralObject){
	for i:=0; i<len(arriPerson); i++{
		arriPerson[i].Print()
	}	
}

func (i *Teacher) Print(){
	fmt.Println("Id:", i.ID)
	fmt.Println("Salary:", i.Salary)
	fmt.Println("Subject:", i.Subject)
	for j:=0; j<len(i.Classroom);j++{
		fmt.Println("Classroom",j+1, "=", i.Classroom[j])
	}
	fmt.Println("Name:", i.Person.Name)
	fmt.Println("Surname:", i.Person.Surname)
	fmt.Println("PersonalCode:", i.Person.PersonalCode)
	fmt.Println()
}

func (i *Teacher) Read(str string) bool{
	return i.ID == str
}

func (action *ReadTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *DeleteTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *UpdateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (action *CreateTeacher) GetFromJSON (rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println(err)
		return
	}
}

var arriPerson []GeneralObject
	
func main() {
	file, _ := os.Open("data.json")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var act Action
		var obj GeneralObject
		var toDo DefinedAction
		err := json.Unmarshal([]byte(scanner.Text()), &act)
		if err != nil {
			fmt.Println("error")
		}
		switch act.ObjName {
		case "Teacher":
			obj = &Teacher{}
		}
		switch act.Action {
		case "create":
			toDo = obj.GetCreateAction()
		case "update":
			toDo = obj.GetUpdateAction()
		case "read":
			toDo = obj.GetReadAction()
		case "delete":
			toDo = obj.GetDeleteAction()
		}
		toDo.GetFromJSON([]byte(scanner.Text()))
		toDo.Process()
	}
}
