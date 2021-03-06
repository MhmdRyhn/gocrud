package schemas


import (
    "github.com/mhmdryhn/gocrud/exceptions"
)


func ValidateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    requiredData := make(map[string]interface{})

    // Validation for "email" key
    // "email" is mandatory
    if value, ok := data["email"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "email must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["email"] = value
        }
    } else {
        return map[string]interface{} {
            "verdict": "email key is required",
        }, exceptions.KEY_ERROR
    }

    // Validation for "name" key
    // "name" is mandatory
    if value, ok := data["name"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "name must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["name"] = value
        }
    } else {
        return map[string]interface{} {
            "verdict": "name key is required",
        }, exceptions.KEY_ERROR
    }
    // Validation for "phone" key
    if value, ok := data["phone"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "phone must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["phone"] = value
        }
    }
    // Validation for "age" key
    if value, ok := data["age"]; ok {
        if value, ok = value.(float64); !ok {
            return map[string]interface{} {
                "verdict": "age must be int or folat",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["age"] = value
        }
    }
    // Validation for "address" key
    if value, ok := data["address"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "address must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["address"] = value
        }
    }

    return requiredData, nil
}


func ValidateAuthorFilterCondition(data map[string]interface{}) (map[string]interface{}, error) {
    requiredData := make(map[string]interface{})
    filterKey := "email"

    // Validation for "email" key
    // "email" key is mandatory
    if value, ok := data[filterKey]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "email must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData[filterKey] = value
        }
    } else {
        return map[string]interface{} {
            "verdict": "email key is required",
        }, exceptions.KEY_ERROR
    }

    return requiredData, nil
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
            }, exceptions.TYPE_ERROR
        } else {
            requiredData[updateKey] = value
        }
    } else {
        return map[string]interface{} {
            "verdict": "email key is required",
        }, exceptions.KEY_ERROR
    }

    // Validation for "name" key
    if value, ok := data["name"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "name must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["name"] = value
        }
    }
    // Validation for "phone" key
    if value, ok := data["phone"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "phone must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["phone"] = value
        }
    }
    // Validation for "age" key
    if value, ok := data["age"]; ok {
        if value, ok = value.(float64); !ok {
            return map[string]interface{} {
                "verdict": "age must be int or folat",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["age"] = value
        }
    }
    // Validation for "address" key
    if value, ok := data["address"]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "address must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData["address"] = value
        }
    }

    return requiredData, nil
}


func ValidateAuthorDeleteCondition(data map[string]interface{}) (map[string]interface{}, error) {
    requiredData := make(map[string]interface{})
    deleteKey := "email"

    // Validation for "email" key
    // "email" key is mandatory
    if value, ok := data[deleteKey]; ok {
        if value, ok := value.(string); !ok {
            return map[string]interface{} {
                "verdict": "email must be string",
            }, exceptions.TYPE_ERROR
        } else {
            requiredData[deleteKey] = value
        }
    } else {
        return map[string]interface{} {
            "verdict": "email key is required",
        }, exceptions.KEY_ERROR
    }

    return requiredData, nil
}
