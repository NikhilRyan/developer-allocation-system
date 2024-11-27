package models

import (
    "encoding/json"
)

// SerializeMap converts a map to a JSON string.
func SerializeMap(data map[string]int) (string, error) {
    bytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// DeserializeMap converts a JSON string to a map.
func DeserializeMap(data string, result *map[string]int) error {
    return json.Unmarshal([]byte(data), result)
}

// SerializeSlice converts a string slice to a JSON string.
func SerializeSlice(data []string) (string, error) {
    bytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// DeserializeSlice converts a JSON string to a string slice.
func DeserializeSlice(data string, result *[]string) error {
    return json.Unmarshal([]byte(data), result)
}

// SerializeIntSlice converts an int slice to a JSON string.
func SerializeIntSlice(data []int) (string, error) {
    bytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// DeserializeIntSlice converts a JSON string to an int slice.
func DeserializeIntSlice(data string, result *[]int) error {
    return json.Unmarshal([]byte(data), result)
}
