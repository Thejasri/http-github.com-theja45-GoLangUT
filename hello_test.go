package hello
import "testing"
func TestHello(t *testing.T){
expected := "Hello openebs!"
actual := hello("openebs")
if actual != expected {
t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
}
}
