# testingExtensions

Simple AssertTrue/AssertFalse functions for testing.
Create the TestUnit and test for some condition. If the test fails, it will fail in the wrapped *testing.T as well.

```go
func TestSomething(t *testing.T) {
  tT := TestUnit(t)
  tT.AssertTrue(someCondition)
}
```
