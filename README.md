# Golang Basics

A Brief example for Golang syntanx and few more, Folder wise Codes are added, you can download and see all the Codes

- Hello World
- Variables
- If Else
- Loop
- Functions
- Structs
- DB Connection
- Array
- Pointers
- Slices
- Date Time and Conversion
- Web Services
- HTTP Handlers
- Methods
- Modules
- Interface
- Testing Methods
- Microservices




## Usage/Examples

```javascript
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello world !")
    }
```


# Golang Testing Sample


```
go test -v

=== RUN   TestCalculate
--- PASS: TestCalculate (0.00s)
=== RUN   TestTblDataToCalculate
--- PASS: TestTblDataToCalculate (0.00s)

PASS
ok      gotest  0.898s


```

# Golang Interface Example (Variable Type, Slice Type, Parameter Type)

```javascript

type Males struct {
	rollNo     int
	name       string
	totalMarks int
}

type Females struct {
	rollNo     int
	name       string
	totalMarks int
}

func (M Males) getInfo() {
	fmt.Println(M.rollNo, M.name, M.totalMarks)
}
func (F Females) getInfo() {
	fmt.Println(F.rollNo, F.name, F.totalMarks)
}

func (F *Females) setInfo() {
	fmt.Scan(&F.rollNo, &F.name, &F.totalMarks)
}
func (M *Males) setInfo() {
	fmt.Scan(&M.rollNo, &M.name, &M.totalMarks)
}

type Students interface {
	setInfo()
	getInfo()
}

func process(S Students) {
	S.setInfo()
	S.getInfo()
}

func main() {

	// Opt-1 Interface setting - Variable type
	// M := Males{}
	// F := Females{}
	// students := []Students{&M, &F}

	// Opt-2 Interface setting - Slice type
	students := []Students{new(Males), new(Females)}

	// Opt-3 Interface setting - Parameter type
	for _, tot := range students {
		process(tot)
	}

}

```



