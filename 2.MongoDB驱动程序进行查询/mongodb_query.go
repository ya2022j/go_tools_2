以下是一个使用MongoDB驱动程序进行查询的示例代码：

go
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// 设置MongoDB连接选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("连接到MongoDB时发生错误：", err)
		return
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("检查MongoDB连接时发生错误：", err)
		return
	}

	fmt.Println("成功连接到MongoDB！")

	// 获取集合对象
	collection := client.Database("mydatabase").Collection("mycollection")

	// 创建过滤条件
	filter := bson.M{"name": "John"}

	// 执行查询操作
	var result bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("执行查询时发生错误：", err)
		return
	}

	// 输出结果
	fmt.Println("查询结果：", result)
}
