package main

import (
	"awesomeProjects/Mongodb"
	"awesomeProjects/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Seed the random number generator to get different values each time
	rand.Seed(time.Now().UnixNano())
	err := Mongodb.ConnectMongoDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/students", handlers.AddStudent).Methods(http.MethodPost)
	router.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods(http.MethodPut)
	router.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods(http.MethodDelete)
	router.HandleFunc("/students", handlers.GetAllStudents).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", handlers.GetStudent).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", handlers.PayAmount).Methods(http.MethodPut)

	log.Println("API is running!")
	err = http.ListenAndServe(":4000", router)
	if err != nil {
		log.Println("Rest Server error" + err.Error())
	}
	/*// Define the range
	start := 13
		end := 25

		// Create a list to store 2 students
		var studentList []models.Student
		for i := 0; i < 2; i++ {
			// Randomize values
			name := generateRandomName()
			zipCode := generateRandomZipCode()
			age := generateRandomAge(start, end)
			height := generateRandomHeight()
			balance := generateRandomBalance()
			// Create a student using the CreateStudent function
			student, _ := Services.CreateStudent(name, zipCode, age, height, balance, "veomn@yahoo.com")

			// Append the student to the list
			studentList = append(studentList, student)
		}

		for i := 0; i < len(studentList); i++ {
			fmt.Printf("Student ID: %d\n", studentList[i].ID)
			fmt.Printf("Name: %s\n", studentList[i].Name)
			fmt.Printf("Zip Code: %s\n", studentList[i].ZipCode)
			fmt.Printf("Age: %d\n", studentList[i].Age)
			fmt.Printf("Height: %.2f\n", studentList[i].Height)
			fmt.Printf("Balance: %.2f\n", studentList[i].Balance)
			fmt.Println("---------------------")
		}
		thresholdBalance := 500.0
		belowThresholdStudents := Services.BalanceChecker(studentList, thresholdBalance)

		// Print details of students with a balance below the threshold
		fmt.Println("Students with a balance below $500:")
		for _, student := range belowThresholdStudents {
			fmt.Printf("Student ID: %d\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Balance: %.2f\n", student.Balance)
			fmt.Println("---------------------")
		}
		searchName := "Veom"
		foundStudents, err := Mongodb.SearchStudentsByName(searchName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Students with name '%s':\n", searchName)
		for _, student := range foundStudents {
			fmt.Printf("Student ID: %s\n", student.ID.Hex())
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Zip Code: %s\n", student.ZipCode)
			fmt.Printf("Age: %d\n", student.Age)
			fmt.Printf("Height: %.2f\n", student.Height)
			fmt.Printf("Balance: %.2f\n", student.Balance)
			fmt.Println("---------------------")
		}

		// Search for students with a balance greater than or equal to a certain amount
		minBalance := 500.0
		foundStudents, err = Mongodb.SearchStudentsByBalance(minBalance)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Students with a balance greater than or equal to $%.2f:\n", minBalance)
		for _, student := range foundStudents {
			fmt.Printf("Student ID: %s\n", student.ID.Hex())
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Zip Code: %s\n", student.ZipCode)
			fmt.Printf("Age: %d\n", student.Age)
			fmt.Printf("Height: %.2f\n", student.Height)
			fmt.Printf("Balance: %.2f\n", student.Balance)
			fmt.Println("---------------------")
		}

		// Delete all students from MongoDB
	//err = Services.DeleteAll()
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		fmt.Println("All students deleted from MongoDB.")
		existingStudent, err := Mongodb.GetAnyStudent()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Retrieve the ObjectID of the existing student
		existingStudentID := existingStudent.ID

		// Print details of the existing student
		fmt.Printf("Existing Student Details:\n")
		fmt.Printf("Student ID: %s\n", existingStudentID.Hex())
		fmt.Printf("Name: %s\n", existingStudent.Name)
		fmt.Printf("Zip Code: %s\n", existingStudent.ZipCode)
		fmt.Printf("Age: %d\n", existingStudent.Age)
		fmt.Printf("Height: %.2f\n", existingStudent.Height)
		fmt.Printf("Balance: %.2f\n", existingStudent.Balance)

		// Now you can perform operations on the existing student, such as paying an amount
		amount := 50.0
		err = Services.Payamount(existingStudentID, amount)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Paid $%.2f for existing student with ID %s\n", amount, existingStudentID.Hex())*/

	// Retrieve the student's balance from MongoDB

	/*studentIDToUpdate := 5
		additionalBalance := 100.0
		//updatedStudentList, err := Services.SearchStudentByIDAndUpdateBalance(studentList, studentIDToUpdate, additionalBalance)

		// Handle the error, if any
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print details of the updated student list
		fmt.Println("Updated student list:")
		for _, student := range updatedStudentList {
			fmt.Printf("Student ID: %d\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Balance: %.2f\n", student.Balance)
			fmt.Println("---------------------")
		}
		// Call the function to print odd numbers in the range
		//util.PrintOddNumbers(start, end)
		//util.PrintEvenNumbers(start, end)

		// Define a string
		//inputString := "programming is fun and challenging"

		// Find the maximum occurring character
		//maxChar, maxCount := util.FindMaxOccurrence(inputString)

		// Print the result
		//fmt.Printf("Maximum occurring character: %c\n", maxChar)
		//fmt.Printf("Occurrences: %d\n", maxCount)

		/*student := Services.CreateStudent("Jane Doe", "54321", 22, 160.0, 150.6, 1)
		fmt.Println("Name:", student.Name)
		fmt.Println("Zip Code:", student.ZipCode)
		fmt.Println("Age:", student.Age)
		fmt.Println("Height:", student.Height)

	}

	// Function to generate a random name
	func generateRandomName() string {
		names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack"}
		return names[rand.Intn(len(names))]
	}

	// Function to generate a random zip code
	func generateRandomZipCode() string {
		return fmt.Sprintf("%05d", rand.Intn(100000))
	}

	// Function to generate a random age in the specified range
	func generateRandomAge(start, end int) int {
		return rand.Intn(end-start+1) + start
	}

	// Function to generate a random height
	func generateRandomHeight() float64 {
		return rand.Float64()*50 + 150
	}

	// Function to generate a random balance
	func generateRandomBalance() float64 {
		return rand.Float64() * 1000

	*/
	/*if err != nil {
		fmt.Println(err)
		return
	}
	err = Services.DeleteAll()
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("All students deleted from MongoDB.")
	*/

	Mongodb.DisconnectMongoDB()
}
