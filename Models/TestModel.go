package Models
import "local/Grocy/Godeps/_workspace/src/labix.org/v2/mgo/bson"

type TestModel struct {
    Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name string `json:"name"`
}