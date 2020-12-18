package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ginner/db/mongodb"
	"ginner/errcode"
)

type userInfoDemo struct {
	Name   string `bson:"name"`
	Age    uint16 `bson:"age"`
	Weight uint32 `bson:"weight"`
}

// HandleMongodbDemo a demo of handle mongodb
func HandleMongodbDemo() error {
	ctx := context.Background()
	// default database that configured in conf/config.toml
	//col := mongodb.GetCollection("user")

	// specific database demo2
	col := mongodb.GetDBCollection("demo2", "user")
	col.EnsureIndexes(ctx, []string{}, []string{"age", "name,weight"})

	// insert one document
	user := userInfoDemo{
		Name:   "xm",
		Age:    7,
		Weight: 40,
	}
	_, err := col.InsertOne(context.Background(), user)
	if err != nil {
		return nil
	}

	// find one document
	one := userInfoDemo{}
	err = col.Find(ctx, bson.M{"name": user.Name}).One(&one)
	if err == mongo.ErrNoDocuments {
		return errcode.ErrRecordNotFound
	}
	if err != nil {
		return err
	}

	// delete one ducument
	//err = col.Remove(ctx, bson.M{"age": 7})
	//if err != nil {
	//return err
	//}

	// multiple insert
	//var userInfos = []userInfo{
	//{Name: "a1", Age: 6, Weight: 20},
	//{Name: "b2", Age: 6, Weight: 25},
	//{Name: "c3", Age: 6, Weight: 30},
	//{Name: "d4", Age: 6, Weight: 35},
	//{Name: "a1", Age: 7, Weight: 40},
	//{Name: "a1", Age: 8, Weight: 45},
	//}
	//result2, err := col.InsertMany(ctx, userInfos)
	//if err != nil {
	//return err
	//}

	// find all, sort and limit
	//batch := []userInfo{}
	//col.Find(ctx, bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)

	// count
	//count, err := col.Find(ctx, bson.M{"age": 6}).Count()

	// UpdateOne one
	//err = col.UpdateOne(ctx, bson.M{"name": "d4"}, bson.M{"$set": bson.M{"age": 7}})

	// UpdateAll
	//result3, err = col.UpdateAll(ctx, bson.M{"age": 6}, bson.M{"$set": bson.M{"age": 10}})

	return nil
}
