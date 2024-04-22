package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusDone       TaskStatus = "done"
	StatusNotStarted TaskStatus = "not_started"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MODEL
type Task struct {
	gorm.Model
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
	User   User       `json:"user"`
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Task{})
	fmt.Println("migrated")
}

func GetAllTasks(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")

	var tasks []Task

	offset := (page - 1) * limit

	query := db.Limit(limit).Offset(offset)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"tasks": tasks,
	})

}

func CreateTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetSingleTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	if err := db.Where("id =?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&task).Where("id =?", id).Updates(task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	if err := db.Where("id =?", id).Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func main() {
	r := gin.Default()
	r.GET("/", GetAllTasks)
	r.POST("/create-task", CreateTask)
	r.GET("task/:id/", GetSingleTask)

	r.Run()
}

// var secretKey = []byte("secret-key")

// func main() {
//   router := mux.NewRouter()

//   router.HandleFunc("/login", login.LoginHandler).Methods("POST")
//   router.HandleFunc("/protected", login.ProtectedHandler).Methods("GET")

//   fmt.Println("Starting the server")
//   err := http.ListenAndServe("localhost:4000", router)
//   if err != nil {
//     fmt.Println("Could not start the server", err)
//   }
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")

//   var u User
//   json.NewDecoder(r.Body).Decode(&u)
//   fmt.Printf("The user request value %v", u)

//   if u.Username == "Chek" && u.Password == "123456" {
//     tokenString, err := CreateToken(u.Username)
//     if err != nil {
//        w.WriteHeader(http.StatusInternalServerError)
//        fmt.Errorf("No username found")
//      }
//     w.WriteHeader(http.StatusOK)
//     fmt.Fprint(w, tokenString)
//     return
//   } else {
//     w.WriteHeader(http.StatusUnauthorized)
//     fmt.Fprint(w, "Invalid credentials")
//   }
// }

// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
//   tokenString := r.Header.Get("Authorization")
//   if tokenString == "" {
//     w.WriteHeader(http.StatusUnauthorized)
//     fmt.Fprint(w, "Missing authorization header")
//     return
//   }
//   tokenString = tokenString[len("Bearer "):]

//   err := verifyToken(tokenString)
//   if err != nil {
//     w.WriteHeader(http.StatusUnauthorized)
//     fmt.Fprint(w, "Invalid token")
//     return
//   }

//   fmt.Fprint(w, "Welcome to the the protected area")

// }

// func createToken(username string) (string, error) {
//  token := jwt.NewWithClaims(jwt.SigningMethodHS256,
//         jwt.MapClaims{
//         "username": username,
//         "exp": time.Now().Add(time.Hour * 24).Unix(),
//         })

//    tokenString, err := token.SignedString(secretKey)
//    if err != nil {
//    return "", err
//    }

//  return tokenString, nil
// }

// func verifyToken(tokenString string) error {
//    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//       return secretKey, nil
//    })

//    if err != nil {
//       return err
//    }

//    if !token.Valid {
//       return fmt.Errorf("invalid token")
//    }

//    return nil
// }
