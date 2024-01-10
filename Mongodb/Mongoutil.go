package Mongodb

import (
	"awesomeProjects/models"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

// Student represents the data structure for a student

var client *mongo.Client
var databaseName = "UTDStudents" // Replace with your actual MongoDB database name
var collectionName = "Students"

// Function to establish a connection to MongoDB
func ConnectMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return nil
}

// Function to disconnect from MongoDB
func DisconnectMongoDB() error {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("Disconnected from MongoDB")
	}
	return nil
}

// Function to add a student to the MongoDB collection
func AddStudent(student models.Student) error {
	collection := client.Database(databaseName).Collection(collectionName)

	_, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		return fmt.Errorf("error inserting student into MongoDB: %v", err)
	}

	fmt.Printf("Student with ID %s added to MongoDB\n", 5)
	return nil
}

func SearchStudentsByName(name string) ([]models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define a filter for the search
	filter := bson.M{"name": name}

	// Perform the search
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	// Decode the results into a slice of students
	var students []models.Student
	err = cursor.All(context.TODO(), &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

// Function to search for students by balance
func SearchStudentsByBalance(minBalance float64) ([]models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define a filter for the search
	filter := bson.M{"balance": bson.M{"$gte": minBalance}}

	// Perform the search
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Decode the results into a slice of students
	var students []models.Student
	err = cursor.All(context.TODO(), &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func DeleteAllStudents() error {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define an empty filter to match all documents
	filter := bson.M{}

	// Perform the delete operation
	_, err := collection.DeleteMany(context.TODO(), filter)
	return err
}

func UpdateStudentBalance(studentID primitive.ObjectID, amount float64) error {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define the filter to match the specific student
	filter := bson.M{"_id": studentID}

	// Define the update to subtract the specified amount from the balance
	update := bson.M{"$inc": bson.M{"balance": -amount}}

	// Perform the update operation
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func GetstudentByID(studentID primitive.ObjectID) (models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define a filter to match the specific student by ObjectID
	filter := bson.M{"_id": studentID}

	// Perform the find operation
	result := collection.FindOne(context.TODO(), filter)

	// Decode the result into a student object
	var student models.Student
	if err := result.Decode(&student); err != nil {
		return models.Student{}, err
	}

	return student, nil
}
func GetAnyStudent() (models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define options to limit the result to one document
	options := options.FindOne()

	// Perform the find operation
	result := collection.FindOne(context.TODO(), bson.D{}, options)

	// Decode the result into a student object
	var student models.Student
	if err := result.Decode(&student); err != nil {
		return models.Student{}, err
	}

	return student, nil
}
func UpdateStudentEmail(studentID primitive.ObjectID, newEmail string) error {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define the filter to find the student by ID
	filter := bson.D{{"_id", studentID}}

	// Define the update to set the new email
	update := bson.D{{"$set", bson.D{{"email", newEmail}}}}

	// Perform the update operation
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	// Check if the update modified any documents
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no document with ID %s found", studentID.Hex())
	}

	return nil
}

func GetStudentByEmail(email string) (models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define the filter to find the student by email
	filter := bson.D{{"email", email}}

	// Create an empty student object to store the result
	var student models.Student

	// Perform the find one operation
	err := collection.FindOne(context.Background(), filter).Decode(&student)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return student, fmt.Errorf("no student found with email: %s", email)
		}
		return student, err
	}

	return student, nil
}

func DeleteStudentByID(studentID int) error {
	collection := client.Database(databaseName).Collection(collectionName)

	// Construct a filter to find the student by ID
	filter := bson.M{"id": studentID}

	// Delete the student from MongoDB
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

func GetAllStudents() ([]models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Define an empty filter to retrieve all documents
	filter := bson.M{}

	// Retrieve all students from MongoDB
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var students []models.Student
	if err := cursor.All(context.Background(), &students); err != nil {
		return nil, err
	}

	return students, nil
}

func GetStudentByID(studentID int) (*models.Student, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	// Construct a filter to find the student by ID
	filter := bson.M{"id": studentID}

	// Retrieve the student from MongoDB
	var student models.Student
	err := collection.FindOne(context.Background(), filter).Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func UpdateStudentByID(studentID int, updatedStudent *models.Student) error {
	collection := client.Database(databaseName).Collection(collectionName)

	// Construct a filter to find the student by ID
	filter := bson.M{"id": studentID}
	updatedStudent.ID = studentID
	// Convert the updatedStudent to BSON format
	update := bson.M{"$set": updatedStudent}

	// Update the student in MongoDB
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func UpdateBalanceByID(id int, amount float64) error {
	collection := client.Database(databaseName).Collection(collectionName)

	// Construct a filter to find the student by ID
	filter := bson.M{"id": id}

	// Construct an update to deduct the amount from the balance
	update := bson.M{"$inc": bson.M{"balance": -amount}}

	// Update the balance in MongoDB
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	// Check if any document was modified (student with the given ID exists)
	if result.ModifiedCount == 0 {
		return errors.New("student not found")
	}

	return nil
}
