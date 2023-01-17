package todo

// import (
// 	"context"
// 	"encoding/json"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-redis/redis/v8"
// 	"yehuizhang.com/go-webapp-gin/db"
// 	"yehuizhang.com/go-webapp-gin/models/user"
// )

// // "todo:userId"
// const DB_PREFIX string = "todo:"

// type TodoItem struct {
// 	ItemId    string `json:"item_id"`
// 	UserId    string `json:"user_id"`
// 	Content   string `json:"content"`
// 	IsDone    bool   `json:"is_done"`
// 	DueAt     int64  `json:"due_at"`
// 	Severity  int8   `json:"severity"`
// 	CreatedAt int64  `json:"created_at"`
// }

// type TodoList = []TodoItem

// // read userInput and assign uid to it
// func processInputData(c *gin.Context) (*TodoItem, error) {

// 	todoItem := TodoItem{}

// 	err := c.Bind(&todoItem)

// 	if err != nil {
// 		return nil, err
// 	}

// 	uid := c.GetString(user.UID)
// 	todoItem.UserId = uid

// 	return &todoItem, nil
// }

// func getTodoListFromDB(uid string) (*[]TodoItem, error) {

// 	dbClient := db.GetRedisDB()

// 	result, err := dbClient.Get(context.Background(), DB_PREFIX+uid).Result()

// 	switch {
// 	case err == redis.Nil:
// 		return nil, nil
// 	case err != nil:
// 		return nil, err
// 	}

// 	var todoList TodoList
// 	err = json.Unmarshal([]byte(result), &todoList)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &todoList, nil
// }

// func storeTodoLIstToDB(uid string, todoList *[]TodoItem) error {

// 	// encodedList, err := json.Marshal(*todoList)

// 	// if err != nil {
// 	// 	return err
// 	// }

// 	dbClient := db.GetRedisDB()
// 	return dbClient.Set(context.Background(), DB_PREFIX+uid, *todoList, 0).Err()

// }

// // func (t TodoItem) Add(c *gin.Context) error {
// // 	item, err := processInputData(c)
// // 	if err != nil {
// // 		return nil
// // 	}

// // 	item.CreatedAt = time.Now().UnixNano()
// // 	item.ItemId = uuid.NewV4().String()

// // }
