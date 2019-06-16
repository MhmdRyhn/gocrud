package schemas


import (
    "fmt"
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

    // Validation for "email" key
    // "email" key is mandatory
    if value, ok := data[updateKey]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "email must be string",
            }, errors.New("type_error")
        } else {
            fmt.Println("****** Email ******** : ", value)
            requiredData[updateKey] = value
        }
    } else {
        return map[string]interface{} {
            "verdict": "email key not found",
        }, errors.New("not_found")
    }

    // Validation for "name" key
    if value, ok := data["name"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "name must be string",
            }, errors.New("type_error")
        } else {
            requiredData["name"] = value
        }
    }
    // Validation for "phone" key
    if value, ok := data["phone"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "phone must be string",
            }, errors.New("type_error")
        } else {
            requiredData["phone"] = value
        }
    }
    // Validation for "age" key
    if value, ok := data["age"]; ok {
        if value, ok = value.(float64); !ok {
            return map[string]interface{} {
                "verdict": "age must be int or folat",
            }, errors.New("type_error")
        } else {
            requiredData["age"] = value
        }
    }
    // Validation for "address" key
    if value, ok := data["address"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "address must be string",
            }, errors.New("type_error")
        } else {
            requiredData["address"] = value
        }
    }

    return requiredData, nil
}
