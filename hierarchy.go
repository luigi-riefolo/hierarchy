package hierarchy

import (
	"encoding/json"
	"os"
)

// employeesMap contains the list
// of EmployeeNodes indexed using
// the employee's ID.
var employeesMap map[string]*EmployeeNode

func init() {
	employeesMap = make(map[string]*EmployeeNode)
}

// EmployeeNode contains information about an employee,
// such as ID, name and the list of managed employees.
type EmployeeNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// List of employees, temporarily
	// used to load the JSON data.
	ManagedEmployees []string `json:"employees"`
	Visited          bool
	// Map of employees node being managed.
	EmployeesMap map[string]*EmployeeNode
}

// addReportingEmployee adds a managed employee.
func (e *EmployeeNode) addReportingEmployee(employeeID string) {
	e.EmployeesMap[employeeID] = employeesMap[employeeID]
}

// Hierarchy contains the hierarchy tree, the
// number of employees and the data file used.
type Hierarchy struct {
	Root       *EmployeeNode
	EmployeeNo int
	DataFile   string
}

// GetPaths traverse the hierarchy tree and finds
// all the possible pass from the root to 'employee'.
func (h *Hierarchy) GetPaths(employee string) [][]*EmployeeNode {
	paths := make([][]*EmployeeNode, 0)
	path := make([]*EmployeeNode, 0)

	// Call the worker function
	getPath(h.Root, employee, path, &paths)
	return paths
}

// FindClosestManager find the closest manager for 'employeeOne'
// and 'employeeTwo' that is the farthest from the root.
func (h *Hierarchy) FindClosestManager(employeeOne, employeeTwo string) (manager *EmployeeNode, err error) {

	if _, ok := employeesMap[employeeOne]; !ok {
		return nil, NewErrorMsg(EmployeeNotFoundError, employeeOne)
	}
	if _, ok := employeesMap[employeeTwo]; !ok {
		return nil, NewErrorMsg(EmployeeNotFoundError, employeeTwo)
	}

	// Get the list of available paths from root to employeeOne
	pathsOne := h.GetPaths(employeeOne)
	pathsTwo := h.GetPaths(employeeTwo)

	// Find the shortest path
	var cnt, idx, minIdx, minPathOne int
	minCnt := h.EmployeeNo
	// For each path for employeeOne
	for i := range pathsOne {
		pathOne := pathsOne[i]

		// For each path for employeeTwo
		for j := range pathsTwo {
			pathTwo := pathsTwo[j]

			// Get the closest manager
			cnt, idx = h.findLCA(pathOne, pathTwo)
			if cnt < minCnt {
				minCnt = cnt
				minPathOne = i
				minIdx = idx
			}
		}
	}

	manager = pathsOne[minPathOne][minIdx]
	return manager, err
}

// findLCA finds the Lowest Common Ancestor in two lists
// of EmployeeNodes. It returns the minimum number of
// nodes (employees) between employeeOne and employeeTwo.
func (h *Hierarchy) findLCA(pathOne, pathTwo []*EmployeeNode) (cnt, idx int) {

	cnt = h.EmployeeNo
	// Find LCA candidate
	for x := range pathOne[0 : len(pathOne)-1] {
		for y := range pathTwo[0 : len(pathTwo)-1] {

			// Found LCA candidate
			if pathOne[x] == pathTwo[y] {
				// Calculate the number of nodes
				cnt1 := (len(pathOne) - 1) - x
				cnt2 := (len(pathTwo) - 1) - y

				if (cnt1+cnt2)-1 < cnt {
					cnt = (cnt1 + cnt2) - 1
					idx = x
				}
			}
		}
	}
	return cnt, idx
}

// getPath is a recursive function to find all
// possible paths from 'root' to 'employee'.
func getPath(root *EmployeeNode, employee string, path []*EmployeeNode, paths *[][]*EmployeeNode) {

	if root.ID == employee {
		newPath := make([]*EmployeeNode, len(path))
		copy(newPath, path)
		newPath = append(newPath, root)
		*paths = append(*paths, newPath)
		return
	}
	path = append(path, root)
	// Find the next node to visit
	for i := range root.EmployeesMap {
		if root.EmployeesMap[i].Visited == false {
			getPath(root.EmployeesMap[i], employee, path, paths)
		}
	}
}

// NewHierarchy is the constructor function for the hierarchy object.
func NewHierarchy(dataFile string) (hiearchy *Hierarchy, err error) {
	// Load the list of employees
	employees, err := loadEmployeesList(dataFile)
	if err != nil {
		return nil, err
	}

	// Create the organisation's hierarchy
	hierarchyTree := makeHierarchyTree(employees)

	hierarchy := &Hierarchy{
		hierarchyTree,
		len(employees),
		dataFile}

	return hierarchy, err
}

// loadEmployeesList reads and unmarshals
// the management hiearchy data file.
func loadEmployeesList(dataFile string) (employees []*EmployeeNode, err error) {
	data, err := os.Open(dataFile)
	if err != nil {
		return nil, NewError(DataFileOpenError, err)
	}

	// Create the list of employees
	err = json.NewDecoder(data).Decode(&employees)
	if err != nil {
		return nil, NewError(DataFileDecodeError, err)
	}

	err = data.Close()
	if err != nil {
		return nil, NewError(DataFileCloseError, err)
	}
	return employees, err
}

// makeHierarchyTree builds the hierarchy tree.
func makeHierarchyTree(employees []*EmployeeNode) *EmployeeNode {
	// Add all the employees to the temporary map
	for _, employee := range employees {
		employeesMap[employee.ID] = employee
	}

	popEmployee := func() *EmployeeNode {
		if len(employees) == 0 {
			return nil
		}
		var employee = employees[0]
		employees = employees[1:len(employees)]
		return employee
	}

	var root = popEmployee()
	var curr = root
	// Link all the managers to their reporting employees
	for {
		curr.EmployeesMap = make(map[string]*EmployeeNode)
		for _, reportingEmployee := range curr.ManagedEmployees {
			curr.addReportingEmployee(reportingEmployee)
		}

		// Delete current employee from the map
		//delete(employeesMap, curr.ID)
		// Get next employee
		if len(employees) == 0 {
			break
		}
		curr = popEmployee()
	}

	return root
}
