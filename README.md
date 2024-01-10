# StudentAPI
Go Student API to demonstrate Student CRUD with Mongo DB
Purpose: The StudentAPI is presumably a part of a larger project or system developed in the Go programming language. 
It is designed to demonstrate CRUD (Create, Read, Update, Delete) operations for student records using MongoDB as the underlying database.
Components: The API would likely include endpoints for creating, reading, updating, and deleting student records. 
These endpoints would interact with a MongoDB database to perform the corresponding operations.
# Install Mongo DB
Purpose: MongoDB is a NoSQL database that stores data in a flexible, JSON-like format. It's commonly used in web development to store and manage data.
Installation: The installation process involves downloading and installing the MongoDB server on your machine. Depending on your operating system, the steps might vary. 
You can refer to the MongoDB official documentation for detailed instructions.

Initial start Mongo DB command

docker run --name mongo -d -p 27017:27017 mongodb/mongodb-community-server:latest

Start Mongo DB

docker start mongo

Stop mongo DB

docker stop mongo

//Use following command to check if mongo process is working
docker ps

Configuration
Database Name: UTDStudents
Collection Name: Students

# Git Checkout
Purpose: This step involves using Git, a version control system, to check out a specific version or branch of 
the project from a Git repository.
Command: The exact command may vary based on your repository structure and branch or tag names. Generally, it would be something like:
git clone <url>


# Execute Go build commands
go mod tidy
go mod vendor
go build <name of project>
Purpose: These commands are used to manage dependencies and build the Go project.
Commands:
go mod tidy: Cleans up the go.mod file, removing any dependencies that are no longer needed.
go mod vendor: Copies dependencies into the vendor directory.
go build <name of project>: Builds the Go project. Replace <name of project> with the actual name of your project.
# Import Student API postman json in Postman
Purpose: Postman is a popular API testing tool. This step involves importing a 
JSON file (presumably containing Postman requests and configurations) related to the Student API.
Process: Open Postman, look for the import option, and select the JSON file. 
This will load the predefined requests and configurations into Postman.
# Test CRUD API using Postman
Purpose: After importing the Postman collection, you can use it to test the CRUD operations of the Student API.
Steps: Open the imported collection in Postman, select the appropriate request (e.g., create, read, update, delete), 
and execute it against the Student API. Verify that the API behaves as expected, 
interacting with the MongoDB database to perform the CRUD operations on student records.
