package schemas


import (
    "errors"
)


func ValidateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    var resp map[string]interface{}
    resp = make(map[string]interface{})

    if _, ok := data["name"]; !ok {
        resp["verdict"] = "name key is missing"
    } else if _, ok := data["email"]; !ok {
        resp["verdict"] = "email key is missing"
    }

    if _, ok := resp["verdict"]; ok {
        return resp, errors.New("invalid")
    }
    return resp, nil
}


func ValidateAuthorFilterCondition(data map[string]interface{}) (map[string]interface{}, error) {
    resp := make(map[string]interface{})

    if _, ok := data["email"]; !ok {
        resp["verdict"] = "email key is missing"
    }

    // if _, ok := data["apply"]; !ok {
    //     resp["verdict"] = "operation not defined (apply key is missing)"
    // }

    if _, ok := resp["verdict"]; ok {
        return resp, errors.New("invalid")
    }
    return resp, nil
}


func ValidateAuthorUpdateData(data map[string]interface{}) (map[string]interface{}, map[string]interface{}, error) {
    keyValue := make(map[string]interface{})
    allowedKeys := []string {
        "email", "name", "phone", "age", "address",
    }
    updateKey := "email"
    for _, key := range allowedKeys {
        if value, ok := data[key]; ok {
            keyValue[key] = value
        }
    }
}
