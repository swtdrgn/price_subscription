package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"time"
)

type msg struct {
  Id    bson.ObjectId `bson:"_id"`
  Msg   string        `bson:"msg"`
  Count int           `bson:"count"`
}

type subscriptions struct {
  Id              bson.ObjectId          `bson:"_id"`
  Subscriptions   map[string]string      `bson:"subscriptions"`
  Prices          map[string]float64     `bson:"prices"`
}
 
func main() {
	uri := os.Getenv("OPENSHIFT_MONGODB_DB_URL")
	fmt.Println(uri)
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}
	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	fmt.Println(sess.DB("neweggtracker").CollectionNames())

	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("neweggtracker").C("subscriptions")
	fmt.Println(collection)
	fmt.Println("Here...")
	//collection := sess.DB("subscriptions").C("foo")
	//collectionNames,err := db.CollectionNames()
	//it := db.Find(nil).Iter() //.All(&results)

	var results interface{}
	fmt.Println("Finding...")
	fmt.Println(collection.Find(nil))
	fmt.Println("Found...")
	fmt.Println(results)

	var subscription subscriptions;
	for it := collection.Find(nil).Iter(); it.Next(subscription); {
		fmt.Println("Updating...")
		err = collection.Update(bson.M{"_id": subscription.Id}, bson.M{"$set": bson.M{"prices": bson.M{time.Now().String(): 12.2}}})
		if err != nil {
			fmt.Printf("Can't update document %v\n", err)
			os.Exit(1)
		}
	}

	/*for _,collectionName := range collectionNames {
		collection := db.C(collectionName);
		err = collection.Update(bson.M{"_id": collectionName}, bson.M{"$set": bson.M{"prices": bson.M{time.Now().String(): 12.2}}})
		if err != nil {
			fmt.Printf("Can't update document %v\n", err)
			os.Exit(1)
		}
	}*/

	/*doc := msg{Id: bson.NewObjectId(), Msg: "Hello from go"}
	err = collection.Insert(doc)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}

	err = collection.Update(bson.M{"_id": doc.Id}, bson.M{"$inc": bson.M{"count": 1}})
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
		os.Exit(1)
	}

	var updatedmsg msg
	err = sess.DB("test").C("foo").Find(bson.M{}).One(&updatedmsg)
	if err != nil {
		fmt.Printf("got an error finding a doc %v\n")
		os.Exit(1)
	}

	fmt.Printf("Found document: %+v\n", updatedmsg)*/
}