package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	// 字段需要首字母大写
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		logArr     []interface{}
		// result     *mongo.InsertOneResult
		result   *mongo.InsertManyResult
		insertId interface{}
		// docId    primitive.ObjectID
	)
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
		return
	}
	// client, err = mongo.Connect(
	// 	context.TODO(),
	// 	options.Client().ApplyURI("mongodb://localhost:27017"))

	// 2, 选择数据库my_db
	database = client.Database("cron")

	// 3, 选择表my_collection
	collection = database.Collection("log")

	// 4, 插入记录(写入)(bson)
	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	// **插入单条**
	// if result, err = collection.InsertOne(context.TODO(), record); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// docId = result.InsertedID.(primitive.ObjectID)
	// fmt.Println("自增ID:", docId.Hex())

	// 5, 批量插入多条document
	record2 := &LogRecord{
		JobName:   "job11",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	record3 := &LogRecord{
		JobName:   "job12",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	logArr = []interface{}{record, record2, record3}

	// 批量插入多条document
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}
	for _, insertId = range result.InsertedIDs {
		fmt.Println("自增ID:", insertId.(primitive.ObjectID).Hex())
	}

}
