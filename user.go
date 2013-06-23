package authygo

import(
    "encoding/json"
    "bytes"
)

type messageT string
type userIdT int

type User struct {
    Id userIdT `json:"user"`
    Errors map[string]string `json:"errors"`
    Message messageT `string:"message"`
    Valid bool
}

func (s *messageT) UnmarshalJSON(b []byte) error {
    if bytes.Compare([]byte("null"), b) == 0 {
        return nil
    }

    return json.Unmarshal(b, (*string)(s))
}

func (i *userIdT) UnmarshalJSON(b []byte) error {
    var userInfo map[string]int
    err := json.Unmarshal(b, (&userInfo))
    *i = userIdT(userInfo["id"])

    return err
}

