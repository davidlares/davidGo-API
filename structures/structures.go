package structures
// JSON Struct
type User struct {
  Id int `json:"id"`
  Username string `json:"username"`
  First_Name string `json:"first_name"`
  Last_Name string `json:"last_name"`
}

type Response struct {
  Status string `json:"status"` // ok and error
  Data User `json:"data"`
  Message string `json:"message"`
}
