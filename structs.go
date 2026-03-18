package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Person struct{
	firstName string
	lastName string
	age int
	createdAt time.Time

}

func main(){
	// timeValue:=time.Now()
	var person1 Person
	person1.firstName="benjamin"
	person1.lastName="Carson"
	person1.age=45
	person1.createdAt=time.Now()
	person,err:=createPerson(&person1)
	if err != nil{
		log.Fatal(err)
		 // errors.New("")
	}
	fmt.Println(person)
}


func createPerson(person *Person)(personConstruct Person,err error){
	// var person Person
	// state:=true
	if person.age == 0 || person.firstName =="" || person.lastName==""{
		// state=false
		return Person{},errors.New("cannot create person construct")
	}else{

		personResult:=Person{
			firstName: person.firstName,
			lastName:person.lastName,
			age:person.age,
		}
		return personResult,nil
	}






}
