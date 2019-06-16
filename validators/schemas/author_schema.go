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
        return resp, errors.New("not_found")
    }
    return resp, nil
}


func ValidateAuthorFilterCondition(data map[string]interface{}) (map[string]interface{}, error) {
    resp := make(map[string]interface{})

    if _, ok := data["email"]; !ok {
        resp["verdict"] = "email key is missing"
    }

    if _, ok := resp["verdict"]; ok {
        return resp, errors.New("not_found")
    }
    return resp, nil
}


func ValidateAuthorUpdateData(data map[string]interface{}) (map[string]interface{}, error) {
    requiredData := make(map[string]interface{})

    updateKey := "email"
    if value, ok := data[updateKey]; ok {
        requiredData[updateKey] = value
    } else {
        return map[string]interface{} {
            "verdict": "email key not found",
        }, errors.New("not_found")
    }


    allowedKeys := []string {
        "email", "name", "phone", "age", "address",
    }

    for _, key := range allowedKeys {
        if value, ok := data[key]; ok {
            requiredData[key] = value
        }
    }

    return requiredData, nil
}
