package collectionfunctions

import (
    "fmt"
    "strings"
    "testing"
)

func fruits() []string {
    return []string{"peach", "apple", "pear", "plum"}
}

func errorMsg(funcName string) string {
    return fmt.Sprintf("%s test failed.", funcName)
}

var strs = fruits()

func TestIndexFindsPeach(t *testing.T) {
    if Index(strs, "peach") != 0 {
        t.Error(errorMsg("Index"))
    }
}

func TestIndexDoesNotFindGrape(t *testing.T) {
    if Index(strs, "grape") != -1 {
        t.Error(errorMsg("Index"))
    }
}

func TestIncludeFindsPeach(t *testing.T) {
    if Include(strs, "peach") == false {
        t.Error(errorMsg("Include"))
    }
}

func TestIncludeDoesNotFindGrape(t *testing.T) {
    if Include(strs, "grape") == true {
        t.Error(errorMsg("Include"))
    }
}

func TestAnyStartsWithP(t *testing.T) {
    if !Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }) {
        t.Error(errorMsg("Any"))
    }
}

func TestAnyDoNotStartWithS(t *testing.T) {
    if Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "s")
    }) {
        t.Error(errorMsg("Any"))
    }
}

func TestAllStartWithP(t *testing.T) {
    if All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }) {
        t.Error(errorMsg("All"))
    }
}

func TestFilterThoseWhichContainE(t *testing.T) {
    expected := []string{"peach", "apple", "pear"}
    filtered := Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    })
    // Can't compare slices
    // Slices can only be compared to nil
    for i := 0; i < 2; i++ {
        if filtered[i] != expected[i] {
            t.Error(errorMsg("Filter"))
        }
    }
}

func TestMap(t *testing.T) {
    expected := []string{"PEACH", "APPLE", "PEAR", "PLUM"}
    upcased := Map(strs, strings.ToUpper)
    // Can't compare slices
    // Slices can only be compared to nil
    for i := 0; i < 3; i++ {
        if upcased[i] != expected[i] {
            t.Error(errorMsg("Map"))
        }
    }
}
