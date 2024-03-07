// // // package main

// // // import (
// // // 	"context"
// // // 	"fmt"
// // // 	"time"

// // // 	"go.mongodb.org/mongo-driver/bson/primitive"
// // // 	"go.mongodb.org/mongo-driver/mongo"
// // // 	"go.mongodb.org/mongo-driver/mongo/options"
// // // )

// // // type Product struct {
// // // 	ID          primitive.ObjectID `json:"_id" bson:"_id"`
// // // 	Name        string             `json:"product_name" bson:"product_name"`
// // // 	Price       int                `json:"price" bson:"price"`
// // // 	Currency    string             `json:"currency" bson:"currency"`
// // // 	Quantity    string             `json:"quantity" bson:"quantity"`
// // // 	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
// // // 	Vendor      string             `json:"vendor" bson:"vendor"`
// // // 	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
// // // 	SkuID       string             `json:"sku_id" bson:"sku_id"`
// // // }

// // // var iphone11 = Product{
// // // 	ID:          primitive.NewObjectID(),
// // // 	Name:        "iphone 11",
// // // 	Price:       800,
// // // 	Currency:    "dollar",
// // // 	Quantity:    "5",
// // // 	Discount:    0,
// // // 	Vendor:      "peter",
// // // 	Accessories: []string{"lens", "silicone cover", "charger cable"},
// // // 	SkuID:       "2345",
// // // }
// // // var trimer = Product{
// // // 	ID:          primitive.NewObjectID(),
// // // 	Name:        "trimmer",
// // // 	Price:       120,
// // // 	Currency:    "dollar",
// // // 	Quantity:    "50",
// // // 	Discount:    20,
// // // 	Vendor:      "clepman agency",
// // // 	Accessories: []string{"screw", "saw", "pipe wilder"},
// // // 	SkuID:       "258",
// // // }

// // // func main() {
// // // 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// // // 	if err != nil {
// // // 		fmt.Println(err.Error())
// // // 	}
// // // 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// // // 	defer cancel()
// // // 	err = client.Connect(ctx)
// // // 	if err != nil {
// // // 		fmt.Println(err.Error())
// // // 	}
// // // 	db := client.Database("tronics")
// // // 	collection := db.Collection("products")
// // // 	res, err := collection.InsertOne(context.Background(), trimer)
// // // 	if err != nil {
// // // 		fmt.Println(err.Error())
// // // 	}
// // // 	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp()) //this is typecasting
// // // 	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())

// // // }
// // package main

// // import (
// // 	"context"
// // 	"fmt"
// // 	"strconv"
// // 	"strings"
// // 	"time"

// // 	_ "github.com/go-sql-driver/mysql"
// // 	"go.mongodb.org/mongo-driver/bson"
// // 	"go.mongodb.org/mongo-driver/bson/primitive"
// // 	"go.mongodb.org/mongo-driver/mongo"
// // 	"go.mongodb.org/mongo-driver/mongo/options"
// // )

// // type Product struct {
// // 	ID           primitive.ObjectID `json:"_id" bson:"_id"`
// // 	Name         string             `json:"product_name" bson:"product_name"`
// // 	Price        int                `json:"price" bson:"price"`
// // 	Currency     string             `json:"currency" bson:"currency"`
// // 	Quantity     string             `json:"quantity" bson:"quantity"`
// // 	Discount     int                `json:"discount,omitempty" bson:"discount,omitempty"`
// // 	Vendor       string             `json:"vendor" bson:"vendor"`
// // 	Accessories  []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
// // 	SkuID        string             `json:"sku_id" bson:"sku_id"`
// // 	AdditionTime int64              `json:"addition_time,omniempty" bson:"addition_time,omniempty`
// // }

// // var iphone11 = Product{
// // 	ID:           primitive.NewObjectID(),
// // 	Name:         "iphone 11",
// // 	Price:        800,
// // 	Currency:     "dollar",
// // 	Quantity:     "5",
// // 	Discount:     0,
// // 	Vendor:       "peter",
// // 	Accessories:  []string{"lens", "silicone cover", "charger cable"},
// // 	SkuID:        "2345",
// // 	AdditionTime: time.Now().UnixMicro(),
// // }
// // var trimer = Product{
// // 	ID:       primitive.NewObjectID(),
// // 	Name:     "hammer",
// // 	Price:    150,
// // 	Currency: "dollar",
// // 	Quantity: "50",
// // 	Discount: 20,
// // 	Vendor:   "clepman agency",
// // 	// Accessories:  []string{"screw", "saw", "pipe wilder"},
// // 	Accessories:  []string{},
// // 	SkuID:        "258",
// // 	AdditionTime: time.Now().UnixMicro(),
// // }

// // func main() {
// // 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) //timeout => to get the output
// // 	defer cancel()
// // 	err = client.Connect(ctx)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	db := client.Database("tronics")
// // 	collection := db.Collection("products")
// // 	// res, err := collection.InsertOne(context.Background(), trimer)
// // 	// if err != nil {
// // 	// 	fmt.Println(err)
// // 	// }

// // 	// res, err = collection.InsertOne(context.Background(), bson.D{{"_id", primitive.NewObjectID()}, {"name", "iphone14"}})
// // 	// res, err = collection.InsertOne(context.Background(), bson.D{{"names", bson.A{"andrew", "peter"}}})
// // 	// res, err = collection.InsertOne(context.Background(), bson.M{
// // 	// 	"name":           "iphone",
// // 	// 	"names":          bson.A{"andrew", "peter"},
// // 	// 	"calssification": bson.D{{"type", "electronic"}},
// // 	// })
// // 	// resMany, err := collection.InsertMany(context.Background(), []interface{}{iphone11, trimer})//interface
// // 	// if err != nil {
// // 	// 	fmt.Println(err)
// // 	// }
// // 	// fmt.Println(resMany.InsertedIDs)
// // 	// fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
// // 	var foundPro Product
// // 	err = collection.FindOne(context.Background(), bson.M{"product_name": "trimmer"}).Decode(&foundPro)
// // 	// fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	fmt.Println("findOne=> ", foundPro.ID)
// // 	var find Product
// // 	// comparission operators
// // 	findCursor, err := collection.Find(context.Background(), bson.M{"price": bson.M{"$lt": 900}})
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	for findCursor.Next(context.Background()) {
// // 		err = findCursor.Decode(&find)
// // 		if err != nil {
// // 			fmt.Println(err)
// // 		}
// // 		fmt.Println(find.Name, find.Price)
// // 	}

// // 	// logical operators or , and
// // 	var logicalOperators = bson.M{
// // 		"$and": bson.A{
// // 			bson.M{"price": bson.M{"$gt": 400}},
// // 			bson.M{"price": bson.M{"$lt": 900}}},
// // 	}
// // 	findCursorE, err := collection.Find(context.Background(), logicalOperators)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	var ele Product
// // 	for findCursorE.Next(context.Background()) {
// // 		err = findCursorE.Decode(&ele)
// // 		if err != nil {
// // 			fmt.Println(err.Error())
// // 		}
// // 		fmt.Println(ele.Name, ele.Price)
// // 	}
// // 	fmt.Println(strings.Repeat("*", 20))
// // 	// element operators exists type and so on
// // 	var findWithAccessories = bson.M{
// // 		"accessories": bson.M{"$not": bson.M{"$exists": true}}, //not inverts the relult
// // 	}
// // 	var findEle Product
// // 	findAccessCursor, err := collection.Find(context.Background(), findWithAccessories)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	for findAccessCursor.Next(context.Background()) {
// // 		err = findAccessCursor.Decode(&findEle)
// // 		if err != nil {
// // 			fmt.Println(err)
// // 		}
// // 		if findEle.Currency == "dollar" {
// // 			fmt.Println(findEle.Name, strconv.Itoa(findEle.Price)+"$")
// // 		}
// // 	}
// // 	var findLensArrayOp = bson.M{
// // 		"accessories": bson.M{"$all": bson.A{"saw"}}, //all from the array
// // 	}
// // 	// element operators
// // 	// $all
// // 	// $exists
// // 	// logical operators
// // 	// $and
// // 	// $or
// // 	// comparison operators
// // 	// $gt
// // 	// $lt
// // 	findCursorEle, err := collection.Find(context.Background(), findLensArrayOp)
// // 	var findELement Product
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	for findCursorEle.Next(context.Background()) {
// // 		err = findCursorEle.Decode(&findELement)
// // 		if err != nil {
// // 			fmt.Println(err)
// // 		}
// // 		fmt.Println(findELement.Name, findELement.Accessories)
// // 	}

// // 	// collection.
// // 	// bson.D document
// // 	// el1 := bson.D{{"ele1", "andrew doghry"}} //document
// // 	// fmt.Println(el1)
// // 	// el2 := bson.M{"ele2": "andrew doghry"} // map unordered
// // 	// fmt.Println(el2)
// // 	// el3 := bson.A{"ele2", "andrew doghry"} // array
// // 	// fmt.Println(el3)
// // 	// el4 := bson.E{"ele2", "andrew doghry"} // usually used inside bson.D
// // 	// fmt.Println(el4)

// // 	// el3  :=  bson.A("name": "andrew doghry")// array
// // 	// el4  :=  bson.E("name": "andrew doghry")// unusual used data type

// // 	// bson map
// // 	// mysdb, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/testdb")
// // 	// if err != nil {
// // 	// 	fmt.Println(err)
// // 	// }
// // 	// defer mysdb.Close()
// // 	// err = mysdb.Ping()
// // 	// if err != nil {
// // 	// 	fmt.Println(err)
// // 	// }
// // 	// response, err := mysdb.Query("select id from testdb.students;  ")
// // 	// if err != nil {
// // 	// 	fmt.Println(err)
// // 	// }
// // 	// var data any
// // 	// for response.Next() {
// // 	// 	fmt.Println(response.Scan(&data))
// // 	// }
// // 	// defer response.Close()

// // 	findcursor, err := collection.Find(context.Background(), bson.M{"sku_id": "2345"})
// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}
// // 	for findcursor.Next(context.Background()) {
// // 		findcursor.Decode(&find)
// // 		fmt.Println(find)
// // 	}
// // }

// // /*
// // objectId => counter process id machine id and the time stamp
// // unique one
// // Bson => binary encoded json
// // has more additional types int long date floating point

// // */

// // // maximaillian jonas
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Product struct {
// 	ID           primitive.ObjectID `json:"_id" bson:"_id"`
// 	Name         string             `json:"product_name" bson:"product_name"`
// 	Price        int                `json:"price" bson:"price"`
// 	Currency     string             `json:"currency" bson:"currency"`
// 	Quantity     string             `json:"quantity" bson:"quantity"`
// 	Discount     int                `json:"discount,omitempty" bson:"discount,omitempty"`
// 	Vendor       string             `json:"vendor" bson:"vendor"`
// 	Accessories  []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
// 	SkuID        string             `json:"sku_id" bson:"sku_id"`
// 	AdditionTime int64              `json:"addition_time,omniempty" bson:"addition_time,omniempty`
// }

// var iphone11 = Product{
// 	ID:           primitive.NewObjectID(),
// 	Name:         "iphone 11",
// 	Price:        800,
// 	Currency:     "dollar",
// 	Quantity:     "5",
// 	Discount:     0,
// 	Vendor:       "peter",
// 	Accessories:  []string{"lens", "silicone cover", "charger cable"},
// 	SkuID:        "2345",
// 	AdditionTime: time.Now().UnixMicro(),
// }
// var trimer = Product{
// 	ID:       primitive.NewObjectID(),
// 	Name:     "hammer",
// 	Price:    150,
// 	Currency: "dollar",
// 	Quantity: "50",
// 	Discount: 20,
// 	Vendor:   "clepman agency",
// 	// Accessories:  []string{"screw", "saw", "pipe wilder"},
// 	Accessories:  []string{},
// 	SkuID:        "258",
// 	AdditionTime: time.Now().UnixMicro(),
// }

// func main() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
// 	defer cancel()
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	db := client.Database("tronics")
// 	collection := db.Collection("products")
// 	// collection.UpdateMany(context.Background(), bson.M{}, bson.M{})
// 	findcursoor, err := collection.Find(context.Background(), bson.M{"accessories": bson.M{"$all": bson.A{"saw"}}})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var find Product
// 	for findcursoor.Next(context.Background()) {
// 		err = findcursoor.Decode(&find)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(find)
// 	}
// 	setElement := bson.M{
// 		"$set": bson.M{
// 			"accessories": bson.A{"ele1", "ele2", "ele3"},
// 		},
// 	}
// 	// findEle := bson.M{
// 	// 	"accessories": bson.M{
// 	// 		"$all": bson.A{"saw"},
// 	// 	},
// 	// }
// 	findEle := bson.M{
// 		"accessories": bson.M{"$not": bson.M{
// 			"$exists": true,
// 		}},
// 	}
// 	// updatedOnes, err := collection.UpdateMany(context.Background(), bson.M{"accessories": bson.M{"$all": bson.A{"saw"}}}, bson.M{"$set": bson.M{"currency": "le"}})
// 	updatedOnes, err := collection.UpdateMany(context.Background(), findEle, setElement)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(updatedOnes.ModifiedCount)
// 	// multiply the prices
// 	fmt.Println("---multiply operator ---")
// 	// using bson.M empty will cause cahanges over all
// 	updatedPrices, err := collection.UpdateMany(context.Background(), bson.M{}, bson.M{
// 		"$mul": bson.M{
// 			"price": 1.2,
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("updatedPrices", updatedPrices.ModifiedCount)
// 	fmt.Println("___delete method___")
// 	deletedElec, err := collection.DeleteMany(context.Background(), bson.M{
// 		"accessories": bson.M{
// 			"$all": bson.A{"ele1"},
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("deletedcount => ", deletedElec.DeletedCount)
// 	// deleted all
// 	/*query operators to get the element*/
// }

package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	FName string ` bson:"firstname"`
	LName string ` bson:"lastname"`
}

// json:"firstname"
// json:"lastname"
func Start() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("tronics")
	collection := db.Collection("products")
	// res, err := collection.InsertMany(context.Background(), []interface{}{User{FName: "amal", LName: "doghry"}, User{FName: "john", LName: "doghry"}, User{FName: "andrew", LName: "peter"}})
	res, err := insertData(*collection, []interface{}{User{FName: "amal", LName: "doghry"}, User{FName: "john", LName: "doghry"}, User{FName: "andrew", LName: "peter"}})
	fmt.Println(res.InsertedIDs...)
	if err != nil {
		fmt.Println(err)
	}
	resCursor, err := findMany(*collection, bson.M{"firstname": "andrew"})
	if err != nil {
		log.Println(err)
	}
	var found User
	for resCursor.Next(context.Background()) {
		err = resCursor.Decode(&found)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("found obj =>", found)
	}

	fmt.Println(resCursor)

}
func insertData(col mongo.Collection, i []interface{}) (*mongo.InsertManyResult, error) {
	res, err := col.InsertMany(context.Background(), i)
	if err != nil {
		return res, err
	}
	return res, nil
}
func findMany(col mongo.Collection, filter bson.M) (mongo.Cursor, error) {
	resCuror, err := col.Find(context.Background(), filter)
	if err != nil {
		return *resCuror, err
	}
	return *resCuror, nil
}
func insertOne(col *mongo.Collection, ele interface{}) (*mongo.InsertOneResult, error) {

	res, err := col.InsertOne(context.Background(), ele)
	if err != nil {
		return res, err
	}
	return res, nil
}
