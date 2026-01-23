package flagparser_test
import (
	 "testing"
"ascii-art-color/internal/flagparser"
)
func TestParseArgs_NoArguments (t *testing.T) {
args:=[]string{"./ascii-art"}
err:=flagparser.ParseArgs(args)
if err==nil {
	t.Errorf("Error was expected")
}

}
