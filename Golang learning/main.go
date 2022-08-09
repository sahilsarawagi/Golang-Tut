package main

import "fmt"

func main(){
	// intializing variable and Constants
	var confrenceName = "Python Confrence"  // you can change the var datatype  
	const confrenceTicket = 55    // You cannot change the constants
	// we can also intialize var as 
	eventName:= "Circus";
	fmt.Println(eventName)
	// Defining Array
	var Booking [50]string    // the problem with array is that we have to define how many value array going to have
	// Defining Array and intiallizing
	var Bookings = [5]string{"Samh","bkh"} 
	Bookings[2] = "5454"
	// Intializing Empty Slice
	var emptySlice = []string{}
	emptySlice = append(emptySlice, "24")
	// Defining Slice
	var vookingInSlice []string  // slice is a abstraction of array but here you don't need to provide value that how many value it has going to have
	// Giving value to arrays's first index
	Booking[0]="Samuell" 
	// Giving value to slice's first index
	vookingInSlice=append(vookingInSlice,"Rammau")    // as we appending more and more value to Slice it will be set to the +1 indices

	// getting the value from both array and slice is same 
	fmt.Printf("First value of the array is %v and first vlaue of the slice is %v\n",Booking[0],vookingInSlice[0])
	// Geeting the length, type and printing whole array
	fmt.Printf("Type of Array is %T\nLenght of the array is %v\nPrinting whole array %v\n",Booking,len(Booking),Booking)
	// Geeting the length, type and printing whole Slice
	fmt.Printf("Type of slice is %T\nLenght of the Slice is %v\nPrinting whole Slice %v\n",vookingInSlice,len(vookingInSlice),vookingInSlice)


	// intializing variable and constant expliticitely
	var remainingTicket uint= confrenceTicket

	// Printing on the terminal by using print format and print in new line
	fmt.Printf("Remaining Ticket is of type %T confrencenceName is of Type %T\n",remainingTicket,confrenceName)// %T is the type holder
	fmt.Println("Welcome to our",confrenceName,"to book your ticket")  
	fmt.Printf("We have Total of %v tickets and %v are still remaining\n",confrenceTicket,remainingTicket) // Printing somthing in formated form

	// Defining Variable 
	var userName string       // if we donnt explicityly define the variable it will give us error becos here we have not given any value to it
	var userTicket uint

	// providing the name to the user
	fmt.Println("Enter your name")

	// Input Data through console
	fmt.Scan(&userName)  // Instead of passing the value, we should pass the reference value

	// Printing the value and its address
	fmt.Println(remainingTicket)
	fmt.Println(&remainingTicket)   // address of the remainingTicket stores at the &remainingTicket

	// intializing pre defined  variable and constants
	userTicket= 2
	remainingTicket=remainingTicket-userTicket
	fmt.Printf("%v booked %v tickets",userName,userTicket)

}