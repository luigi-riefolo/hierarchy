package hierarchy

import (
	"fmt"
	"log"
	"os"
	"testing"
)

type hierarchyTest struct {
	employeeOne string
	employeeTwo string
	manager     string
	hasToPass   bool
}

var hierarchyObj *Hierarchy

func TestMain(m *testing.M) {
	dataFile := "data/complete.json"
	var err error
	hierarchyObj, err = NewHierarchy(dataFile)
	if err != nil {
		log.Fatal(err)
	}
	ret := m.Run()

	os.Exit(ret)
}

func testHierarchy(t *testing.T, test hierarchyTest) {
	t.Logf("Finding lowest common manager for '%s' and '%s'\n",
		test.employeeOne, test.employeeTwo)
	manager, err := hierarchyObj.FindClosestManager(test.employeeOne, test.employeeTwo)

	msg := ""
	if err != nil {
		msg = err.Error()
	} else {
		msg = fmt.Sprintf("Found manager:\t%s\n", manager.ID)
	}

	if test.hasToPass == true && manager.ID != test.manager {
		t.Error(msg)
	} else {
		t.Log(msg)
	}

}

func TestHierarchies(t *testing.T) {
	tests := []hierarchyTest{
		{
			employeeOne: "Q",
			employeeTwo: "T1",
			manager:     "B",
			hasToPass:   true,
		},
		{
			employeeOne: "L",
			employeeTwo: "G",
			manager:     "A",
			hasToPass:   true,
		},
		{
			employeeOne: "F",
			employeeTwo: "P",
			manager:     "B",
			hasToPass:   true,
		},
		{
			employeeOne: "A",
			employeeTwo: "B",
			manager:     "A",
			hasToPass:   true,
		},
		{
			employeeOne: "X",
			employeeTwo: "T1",
			manager:     "H",
			hasToPass:   true,
		},
	}
	for _, test := range tests {
		testHierarchy(t, test)
	}
}

func TestNotExistingEmployees(t *testing.T) {
	tests := []hierarchyTest{
		{
			employeeOne: "Z1",
			employeeTwo: "T1",
			manager:     "B",
			hasToPass:   false,
		},
		{
			employeeOne: "Q",
			employeeTwo: "L1",
			manager:     "B",
			hasToPass:   false,
		},
	}
	for _, test := range tests {
		testHierarchy(t, test)
	}
}
