package schemas


import "errors"


func ValidateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    var resp map[string]interface{}
    resp = make(map[string]interface{})
    // author := models.Author{}

    if _, ok := data["name"]; !ok {
        resp["verdict"] = "name key is missing"
    } else if _, ok := data["email"]; !ok {
        resp["verdict"] = "email key is missing"
    }
    // else if _, ok := data["phone"]; !ok {
    //     resp["verdict"] = "phone key is missing"
    // } else if _, ok := data["age"]; !ok {
    //     resp["verdict"] = "age key is missing"
    // } else if _, ok := data["address"]; !ok {
    //     resp["verdict"] = "address key is missing"
    // }


    // if value, ok := data["phone"]; ok {
    //     resp["Phone"] = value.(string)
    // }
    // if value, ok := data["age"]; ok {
    //     resp["Age"] = value.(string)
    // }
    // if value, ok := data["address"]; ok {
    //     resp["Address"] = value.(string)
    // }

    if _, ok := resp["verdict"]; ok {
        return resp, errors.New("invalid")
    }
    return resp, nil
}