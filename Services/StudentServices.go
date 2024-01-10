package Services

import (
	"awesomeProjects/Mongodb"
	"awesomeProjects/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateStudent(name string, zipCode string, age int, height float64, balance float64, email string) (models.Student, error) {

	// Create and return a new Student object
	id := 5
	student := models.Student{
		Name:    name,
		ZipCode: zipCode,
		Age:     age,
		Height:  height,
		Balance: balance,
		ID:      id,
		Email:   email,
	}
	err := Mongodb.AddStudent(models.Student{})
	if err != nil {
		return models.Student{}, err
	}
	return student, nil
}

// Function to check if a student has a balance less than a certain amount
func BalanceChecker(students []models.Student, threshold float64) []models.Student {
	var belowThresholdStudents []models.Student

	for _, student := range students {
		if student.Balance < threshold {
			belowThresholdStudents = append(belowThresholdStudents, student)
		}
	}

	return belowThresholdStudents
}

func DeleteAll() error {
	return Mongodb.DeleteAllStudents()
}
func Payamount(studentID primitive.ObjectID, amount float64) error {
	return Mongodb.UpdateStudentBalance(studentID, amount)
}
