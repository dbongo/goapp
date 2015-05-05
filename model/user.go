package model

// User ...
// type User struct {
// 	ID        bson.ObjectId `bson:"_id" json:"-"`
// 	Admin     bool          `bson:"admin" json:"admin"`
// 	Name      string        `bson:"name" json:"name"`
// 	Email     string        `bson:"email" json:"email"`
// 	Username  string        `bson:"username" json:"username"`
// 	Password  string        `bson:"password" json:"-"`
// 	Created   int64         `bson:"created" json:"created"`
// 	Updated   int64         `bson:"updated" json:"updated"`
// 	LastLogin string        `bson:"lastlogin" json:"lastlogin"`
// }

type User struct {
	ID       int64  `meddler:"user_id,pk"    json:"-"`
	Admin    bool   `meddler:"user_admin"    json:"admin"`
	Name     string `meddler:"user_name"     json:"name"`
	Email    string `meddler:"user_email"    json:"email,omitempty"`
	Username string `meddler:"user_username" json:"username,omitempty"`
	Password string `meddler:"user_password" json:"-"`
	Token    string `meddler:"user_token"    json:"-"`
	Active   bool   `meddler:"user_active"   json:"active"`
	Created  int64  `meddler:"user_created"  json:"created_at"`
	Updated  int64  `meddler:"user_updated"  json:"updated_at"`
}
