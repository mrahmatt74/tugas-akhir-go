package main

import (
	"fmt"
)

func main() {
	// var getMinMax = func(numbers []int) (int, int) {
	// 	var min, max int

	// 	for i, n := range numbers {
	// 		if i == 0 {
	// 			max, min = n, n
	// 		} else if n > max {
	// 			max = n
	// 		} else if n < min {
	// 			min = n
	// 		}
	// 	}
	// 	return min, max
	// }

	// var numbers = []int{1, 4, 7, 2}
	// var min, max = getMinMax(numbers)

	// fmt.Println("Nilai min adalah %d dan nilai max adalah %d \n", min, max)
	// pointer

	// var dataCustomer = customer{Id: 2, Name: "RICHARD", Address: "Balaraja", contact: contact{Phone: "085781544293", Email: "richard@gmail.com"}}

	// CustomerInfo := NewCustomer(&dataCustomer)
	// customerName := CustomerInfo.getName()
	// // infoCustomer = dataCustomer
	// // customerName := infoCustomer.getName()
	// // customerContacts := infoCustomer.getContacts()

	// fmt.Println("Customer name is ", customerName)

	// // for key, v := range customerContacts {
	// // 	fmt.Printf("customer %s is %s \n", key, v)
	// // }

	// var myInterface interface{}

	// myInterface = "Budi"
	// fmt.Println(myInterface)

	// myInterface = 123456
	// fmt.Println(myInterface)

	// myInterface = true
	// fmt.Println(myInterface)

	// var data = map[string]any{
	// 	"name":    "RAHMAT",
	// 	"age":     21,
	// 	"hobbies": []string{"coding", "volley", "tahfidz"},
	// }

	// fmt.Println(data)

	var myInterface any

	myInterface = 10
	var hasilPerkalian = myInterface.(int) * 2

	fmt.Println(myInterface, "dikalikan", 2, " = ", hasilPerkalian)

	var myInterface2 any = &customer{Id: 1, Name: "Rahmat"}
	var name = myInterface2.(*customer).Name

	fmt.Println(name)
}

type customer struct {
	Id   uint
	Name string
	// Address string
	// contact
}

type contact struct {
	Phone, Email string
}

func (cust customer) getName() string {
	return cust.Name
}

// func (cust customer) getContacts() map[string]string {
// 	return map[string]string{
// 		"Phone": cust.contact.Phone,
// 		"Email": cust.contact.Email,
// 	}
// }

type Information interface {
	getName() string
}

// func NewCustomer(cust *customer) Information {
// 	return &customer{
// 		Id:      cust.Id,
// 		Name:    cust.Name,
// 		Address: cust.Address,
// 	}
// }
