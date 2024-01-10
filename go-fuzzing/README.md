# Go Fuzzing
A look at fuzzing in Go. This project has an intentional bug that can be found out with fuzzing.

## Bug
The bug is that if the highest value is 50 then it will return an internal server error.

## Testing
In order to run the test you need to run the following command. (Uses `Makefile`)
```shell
make fuzz-test
```