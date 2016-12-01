# hierarcy

Go implementation of a simple hierarchical organization.


## Description

The hierarchical tree is composed of employees, from the root representd by the
CEO down to basic employees.
An employee with at least one reporting is called a Manager and must have a
list of other employees reporting to him/her.

An employee is represented by an employee node containing information such as ID,
name and a map containing their managed employees.

'hierarchy' provides the method FindClosestManager which provides
an interface to find the closest common manager (i.e. farthest from the CEO)
between two employees.


## Testing

    Run all the unit-tests with:

        go test -v


## Examples

	Here's a simple example on how to use the 'hierarchy' package:

		import "github.com/luigi-riefolo/hierarchy"

	Create the hierarchy object:

		dataFile := "organisation-data.json"
        hierarchy, err := hierarchy.NewHierarchy(dataFile)
        if err != nil {
                log.Fatal(err)
        }

	Find the lowest common manager:

        employeeOne, employeeTwo := "0032", "0024"

        manager, err := hierarchy.FindClosestManager(employeeOne, employeeTwo)
        if err != nil {
                log.Fatal(err)
        } else {
                fmt.Println("Found manager: ", manager.ID)
        }


	See main.go in examples/ for additional information.


## License

The code is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
