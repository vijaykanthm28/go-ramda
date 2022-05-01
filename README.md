# go-ramda

Functional ramda (<https://ramdajs.com>) methods implement in Go

This package supports limited functions of ramda which are used ofter

## Functions

*The following functions are supported ramda func with some changes which is supported by go*

| function      | purpose                                                                               |
| ------------- | ------------------------------------------------------------------------------------- |
| And           | Logical `And` operation on given list of bool args                                    |
| Or            | Logical `Or` operation on given list of bool args                                     |
| Not           | Logical `Not` operation on given bool                                                 |
| IfElse        | IfElse is just act like normal if else if true return 1st value other wise 2nd value  |
| Path          | Path will return given path has value other wise return nil                           |
| PathOr        | Path will return given path has value other wise return default value that on 1st arg |
| PathEq        | PathEq will return func gets value and checks that value equals to given element      |
| HasPath       | HasPath return true if given path avail on data object                                |
| IsEmpty       | IsEmpty return true if given object does't have any value                             |
| IsNil         | IsNil return true if given object is nil                                              |
| PropSatisfies | PropSatisfies return execut given func if given key field avail on data               |
| PathSatisfies | PathSatisfies return execut given func if given path avail on data                    |
| Equals        | Equals return true if given two arguments are equal                                   |
| NotEquals     | NotEquals return true if given two arguments are not equal                            |
| Union         | Merges to array of elements into an array with out duplicate                          |
| Difference    | returns an array that contain elements are in 1st array and not in 2nd array          |
| Head          | returns the first element on given array                                              |
| Find          | returns the matched element on given array                                            |
| Tail          | returns array without first element on given array                                    |
| Last          | returns the last element on given array                                               |
| IndexOf       | returns the first element matched index on given array                                |
| LastIndexOf   | returns the last element matched index on given array                                 |
| Includes      | returns true if the given element available on given array                            |
| Drop          | returns removes sets of elements upto given value index on given array                |
| DropLast      | returns removes sets of elements upto given value index from the last on given array  |
| Append        | appends given elements into give first array argument                                 |
|---------------|---------------------------------------------------------------------------------------|

*Following functions are type assertion functions which assert interface to go types*

| function      | purpose                                                                               |
| ------------- | ------------------------------------------------------------------------------------- |
| String        | assert to string type also converts string convertable values int, float64 and etc types into string     |
| Integer       | assert to int from interface  |
| Integer32     | assert to int32 from interface                             |
| Integer64     | assert to int64 from interface |
| Float64       | assert to float64 from interface |
| IntSlice      | assert to []int from interface |
| StringSlice   | assert to []string from interface |
| Typify        | returns given data if type of first arg and second arg other wise returns 2nd default value of 2nd argument type |

* * *

* * *
# Examples

*And(bool, bool...)*
This function will return bool value true or false
will return true If give all arguments are true.

```
  And(true, false, false) ==> false
  And(true, false, true) ==> false
  And(true, true, true, true, true) ==> true
```

*Or(bool, bool...)*
This function will return bool value true or false
Or will return true if give arguments atleast one argument is true.

```
  Or(true, false, false) ==> false
  Or(true, false, true) ==> false
  Or(true, true, true, true, true) ==> true
```

*Not(bool)*
This function will return bool value true or false
Not will return true if give argument is false.

```
user := User{
  Student: &UserStudent{
    Name: "Ambika",
    Address: Address{
      Street: "123 Car street",
    },
  },
  }
  Not(true) ==> false
  Not(false) ==> true
  Not(HasPath([]string{"Student", "Name"}, )) ==> false
```

*IfElse(bool, interface{}, interface{}) interface{}*
This function will return interface that can be any type based on 2nd and 3rd argument
If the given bool argument is true then will return 2nd argument otherwise will return 3rd argument.
arguments:
  - bool -> decider argument based on this return first or second argument
  - interface{} -> if section value this can be any type of value
  - interface{} -> if section value this can be any type of value

```
user := User{
  Student: &UserStudent{
    Name: "Ambika",
    Address: Address{
      Street: "123 Car street",
    },
  },
  }
  a := 5
  b := 6
  IfElse(a > b, a, b) ==> 6
  IfElse(IsEmpty(user.Student.Name), "Student-Name", user.Student.Name) ==> Ambika
  IfElse(HasPath([]string{"Student", "Name"}, user)) ==> false
```
_Note: response is interface so we need to do type assertion_

*Path([]string, interface{}) interface{}*
  This function will return any value on the given path in the 2rd argument data.
  If the given data is not have that path or traversable then will return nil.
arguments:
  - []string -> array of key string field name, path value to travers into the given object value, To indexing on slice next to the slice key field should be number string so we can index value on array.
  - interface{} -> object value to be search given path. can be Struct/Map/Slice that can be traversable

```
user := User{
  Student: &UserStudent{
    Name: "Ambika",
    Address: Address{
      Street: "123 Car street",
    },
  },
  Subjects: []Subject{
    {
      Name: "tamil",
      Total: 100,
      MarkSplit: MarkSplit{
        Theory: 80,
        Practical: 20,
      },
    },
    {
      Name: "chemistry",
      Total: 200,
      MarkSplit: MarkSplit{
        Theory: 60,
        Practical: 40,
      },
    },
  },
  }

  Path([]string{"Student", "Name"}, user) ==> Ambika
  Path([]string{"Student", "FirstName"}, user) ==> nil
  String(Path([]string{"Student", "Subjects", "2", "Name"}, user)) ==> "chemistry"

  a := "path"
  Path([]string{"A"}, a) ==> nil
```
_Note: response is interface so we need to do type assertion_


*PathOr(interface{}, []string, interface{}) interface{}*
This function will return any value on the given path in the 3rd argument data.
If the given data is not have that path or traversable then will return 1st argument value.
arguments:
  - interface{} -> default value argument this can be any type of value
  - []string -> array of key string field name, path value to travers into the given object value, To indexing on slice next to the slice key field should be number string so we can index value on array.
  - interface{} -> object value to be search given path. can be Struct/Map/Slice that can be traversable

```
// refer the user on example of Path function
PathOr("Student-Name", []string{"Student", "Name"}, user) ==> Ambika
PathOr("-", []string{"Student", "FirstName"}, user) ==> "-"
String(PathOr("unknown", []string{"Student", "Subjects", "2", "Name"}, user)) ==> "chemistry"
Int(PathOr(0, []string{"Student", "Subjects", "1", "MarkSplit", "Theory"}, user)) ==> 80

a := "path"
PathOr("", []string{"A"}, a) ==> ""

```
_Note: response is interface so we need to do type assertion_
