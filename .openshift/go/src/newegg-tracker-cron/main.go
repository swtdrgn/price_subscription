package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
	"code.google.com/p/go.net/html"
)

type msg struct {
  Id    bson.ObjectId `bson:"_id"`
  Msg   string        `bson:"msg"`
  Count int           `bson:"count"`
}

type subscription struct {
  Id              string                 `bson:"_id"`
  Current         string                 `bson:"current,omitempty"`
  Subscriptions   []string               `bson:"email"`
  Prices          map[string]float64     `bson:"prices,omitempty"`
}

const (
    REGULAR      = 1
    REBATE       = 2
    OUT_OF_STOCK = 3

    PRICE_OUT_OF_STOCK = "Out of Stock"

	EMAIL_USERNAME = "neweggtracker"
	EMAIL_PASSWORD = "newguyintown"
	EMAIL_SERVER   = "smtp.gmail.com"
	EMAIL_PORT     = 587
)

func findElement2 (token *html.Tokenizer, id string) {
	/*fmt.Println(n.Attr)
	if n.FirstChild != nil {
		findElement(n.FirstChild,id)
	}
	if n.NextSibling != nil {
		findElement(n.NextSibling,id)
	}*/
	for ; token.Err() == nil; token.Next() {
		fmt.Printf("TEXT: %s\n", token.Text())
		for key,val,moreAttr := token.TagAttr(); moreAttr; key,val,moreAttr = token.TagAttr() {
			fmt.Printf("%s: %s\n", key, val)
			if string(key) == "id" && string(val) == id {
				fmt.Println("yes.")
			}
		}
	}
}

func getNeweggPrice (n *html.Node) (price string, status int) {
	if n.Data == "strong" {
		price = strings.Trim(n.FirstChild.Data,"$ \n\t")
		status = REBATE
	}
	for _,attr := range n.Attr {
		if attr.Key == "class" && attr.Val == "zmp" {
			price = strings.Trim(n.FirstChild.Data,"$ \n\t")
			status = REGULAR
		} else if (attr.Key == "alt" && attr.Val == "Auto Notify") || (attr.Key == "title" && attr.Val == "Auto Notify") {
			price = PRICE_OUT_OF_STOCK
			status = OUT_OF_STOCK
		} else if attr.Key == "class" && attr.Val == "soldout" {
			price = PRICE_OUT_OF_STOCK
			status = OUT_OF_STOCK
		}
	}
	if n.FirstChild != nil {
		childPrice,childStatus := getNeweggPrice(n.FirstChild)
		if childStatus > status {
			price = childPrice
			status = childStatus
		}
	}
	if n.NextSibling != nil {
		siblingPrice,siblingStatus := getNeweggPrice(n.NextSibling)
		if siblingStatus > status {
			price = siblingPrice
			status = siblingStatus
		}
	}
	return
}

func SendMail(to []string, subject string, msg string) error {
	fmt.Println("Sending Email: " + msg)
	fmt.Println(to)
	auth := smtp.PlainAuth("",EMAIL_USERNAME,EMAIL_PASSWORD,EMAIL_SERVER)
 
	address := fmt.Sprintf("%v:%v",EMAIL_SERVER,EMAIL_PORT)
 
	//	build our message
	body := []byte("Subject: " + subject + "\r\n\r\n" + msg)
 
	err := smtp.SendMail(
		address,
		auth,
		EMAIL_USERNAME,
		to,
		body,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
 
	return nil
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

	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("neweggtracker").C("subscriptions")
	//collection := sess.DB("subscriptions").C("foo")
	//collectionNames,err := db.CollectionNames()
	//it := db.Find(nil).Iter() //.All(&results)

	//doc := msg{Id: bson.NewObjectId(), Msg: "Hello from go"}
	//collection.Insert(doc)

	//var results interface{}

	var sub subscription;
	for it := collection.Find(bson.M{}).Iter(); it.Next(&sub); {
		t := time.Now()
		year,month,date := t.Date()
		hour,_,_ := t.Clock()
		timestamp := fmt.Sprintf("prices.%4d-%02d-%02d %02d:00", year,month,date,hour)

		neweggResponse,_ := http.Get("http://www.newegg.com/Product/MappingPrice.aspx?Item="+sub.Id)
		defer neweggResponse.Body.Close()

		node,_ := html.Parse(neweggResponse.Body)
		price,status := getNeweggPrice(node)

		if sub.Current == "" {
		} else if status != OUT_OF_STOCK {
			if sub.Current == PRICE_OUT_OF_STOCK {
				SendMail(sub.Subscriptions,"Item Restocked: Newegg Item#"+sub.Id,"Newegg Item #"+sub.Id+" has been restocked for $"+price+".")
			} else {
				p,_ := strconv.ParseFloat(price,64)
				old_price,_ := strconv.ParseFloat(sub.Current,64)
				if p < old_price && sub.Current != price {
					SendMail(sub.Subscriptions,"Price Dropped: Newegg Item#"+sub.Id,"Newegg Item #"+sub.Id+" has dropped from $"+sub.Current+" to $"+price+".")
				} else if p > old_price && sub.Current != price {
					SendMail(sub.Subscriptions,"Price Raised: Newegg Item#"+sub.Id,"Newegg Item #"+sub.Id+" has increased from $"+sub.Current+" to $"+price+".")
				}
			}
		} else if sub.Current != PRICE_OUT_OF_STOCK {
			SendMail(sub.Subscriptions,"Out of Stock: Newegg Item#"+sub.Id,"Newegg Item #"+sub.Id+" is out of stock, previously $"+sub.Current+".")
		}
		
		if price != sub.Current {
			err = collection.Update(bson.M{"_id": sub.Id}, bson.M{"$set": bson.M{timestamp: price, "current": price}})
			if err != nil {
				fmt.Printf("Can't update document %v\n", err)
				return
			}
		}
	}
}