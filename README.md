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
  And(true, false, false) ==> false
  And(true, false, true) ==> false
  And(true, true, true, true, true) ==> true
```
