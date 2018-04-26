## testing

go 提供了 testing 包来做 unittest, go 的 test 是以同名文件命名的 `_test.go`文件，放在相同的目录。

go 提供了几个方法来实现一个 case 测试状况

- Log: Similar to fmt.Println; records the text in the error log.
- Logf: Similar to fmt.Printf. It formats its arguments according to the given
format and records the text in the error log.
- Fail: Marks the test function as having failed but allows the execution to
continue.
- FailNow: Marks the test function as having failed and stops its execution.



|         | log   | logf   |
| ------- | ----- | ------ |
| fail    | error | errorf |
| failnow | fatal | fatalf |


执行 go test -v

```
go test -v 
=== RUN   TestDecode
--- PASS: TestDecode (0.00s)
=== RUN   TestEncode
--- SKIP: TestEncode (0.00s)
    decode_test.go:25: skipping encode test now
PASS
ok      go-tutorial/go-web/ch08 0.010s
```

## skip test


执行 `go test -v -short -cover` 告诉 go test 跳过满足条件的 test


## go 可以执行 parallel test

parallel test 能够加速测试

`go test –v –short –parallel 3` 


## Benchmarking

```
go test -v -cover -short –bench .
```


