###### Introduction To Backend With Golang

### Welcome To Class 5

---

### TODO

1. Databases - What and Why?
2. Types Of Databases
	1. SQL
	2. NoSQL
3. ORMs
4. **GORM**
	1. Installation
	2. SQLite
	3. An Example
5. In Class Assignment

---

### Databases

---

#### Databases - What

They are programs that help you store data


---

#### Databases - Why

Why not just store in a `.txt` file?
- How are you gonna scale?
- How big can a file be?
- You have to write the code for searching, reading, indexing etc
- Bunch of other reasons - Google It

---

### Types Of Databases

There are two **main** types of databases:
1. SQL (aka relational)
2. NoSQL (aka non-relational)

Other types of databases do exist: vector database, etc.

---

#### SQL

- You know tables? You know SQL
- Has a syntax we can use to create, retrieve, update and delete data - which we aren't gonna use. 
- A primary key: Unique identifier for every row in the table
- Two tables can have a "relationship" (aww)

---

- A foreign key: A primary key from another table
- For example, if you had two tables, called user and car. The user table has UserIDs, and the car table has CarID and UserID of the person who owns the car. 

---

User Table:

| UserID | UserName |
| ------ | -------- |
| 1      | Anurag   |
| 2      | Shreya   |

Car Table:

| CarID | UserID | CarName |
| ----- | ------ | ------- |
| 1     | 2      | Innova  |
| 2     | 1      | Ciaz    |

UserID in the Car table is a foreign key

---

#### NoSQL

- Has no structure
- A bunch of "documents"
- What are "documents"? Just JSON
- A collection of documents make a "collection"
- Example: MongoDB, Firestore, etc. 

---

```json
{
	"UserID": 1,
	"UserName": "Anurag",
},
{
	"UserID": 2,
	"UserName": "Shreya"
}
```

---

### ORMs

- **O**bject **R**elational **M**odels
- Provide a layer of abstraction over traditional SQL
- Map table entries to "objects"
- In Golang's case - Structures
- How?

---

### GORM

- A library that provides an ORM in Golang
- Really useful if you don't want to write tedious SQL

---

#### GORM - Installation

- Create a new empty folder, run the following commands in it:

```bash
go mod init day5
```

```bash
go get -u gorm.io/gorm  
```

```bash
go get -u gorm.io/driver/sqlite
```


---

#### GORM - SQLite

- SQLite is a database engine that implements SQL in a contained `.db` file. It requires near 0 configuration
- I've been blabbering about why databases, and now we are back to a file? 
- This is for small applications where a full SQL server is not required and the amount of data is small. 
- Use an actual SQL server like MariaDB, MySQL, Postgress in production. 

---

#### GORM - By Example

```go
// main.go

package main  
  
import (  
  "gorm.io/gorm"  
  "gorm.io/driver/sqlite"  
)
```

---

##### The 'Object Relational Part'

```go
// continuation of main.go

type Classroom struct {
  gorm.Model
  Name string
  Section string
  Strength int
}
```

---

##### The Connecting to Database Part

```go
// continuation of main.go

func main() {  
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})  
  if err != nil {
    panic("database connection screwed up")  
  }
  
  // Migrate the schema
  db.AutoMigrate(&Classroom{})
  
```

---

##### CRUD

```go
// continuation of main.go

  // create
  newClassroom := Classroom{
    Name: "Introduction To Backend",
	Section: "B",
	Strength: 45,
  }
  db.Create(&newClassroom)
```

---

##### CRUD

```go
// continuation of main.go

  // read
  var retreivedClass Classroom  
  db.First(&retreivedClass, 1) 
  // find class with ID = 1
  
  // or

  db.First(&retreivedClass, "strength = ?", 45)
  // finds the first class with strength = 45
```

---

##### CRUD

```go
// continuation of main.go
  // update
  db.Model(&retreivedClass).Update("Strength", 20)
  // update the retreived class to have strength 20
```

---

##### CRUD

```go
// continuation of main.go
  // delete
  db.Delete(&retreivedClass)
  // delete the retreived class
  
```

---

#### Final Code:

```go
// main.go

package main  
  
import (  
  "gorm.io/gorm"  
  "gorm.io/driver/sqlite"  
)

type Classroom struct {
  gorm.Model
  Name string
  Section string
  Strength int
}


func main() {  
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})  
  if err != nil {
    panic("database connection screwed up")  
  }
  
  // Migrate the schema
  db.AutoMigrate(&Classroom{})

  // create
  newClassroom := Classroom{
    Name: "Introduction To Backend",
	Section: "B",
	Strength: 45,
  }
  db.Create(&newClassroom)

  // read
  var retreivedClass Classroom  
  db.First(&retreivedClass, 1) 
  // find class with ID = 1
  fmt.Println(retreivedClass)
  
  // or

  db.First(&retreivedClass, "strength = ?", 45)
  // finds the first class with strength = 45

  // update
  db.Model(&retreivedClass).Update("Strength", 20)
  // update the retreived class to have strength 20

  db.Delete(&retreivedClass)
  // delete the retreived class
}
```


---

### In Class Assignment

- Create a GORM + GIN application where the following routes are defined:
- `POST /classroom` will accept classroom details in the request and create a classroom with the given details
- `GET /classroom` will return a list of classrooms that are created

---

- GORM has excellent documentation and this is an assignment to practice your docs reading skills. 
- Particularly, focus on the "CRUD Interface" section of the documentation.
- Do not directly use ChatGPT - you're not getting learning anything with that




