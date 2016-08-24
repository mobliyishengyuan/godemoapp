package actions

import (
	"github.com/mobliyishengyuan/goframework/library"
	"net/http"
	"fmt"
)

func init() {
	library.AddRoute("/test", testAction)
}

func testAction(w http.ResponseWriter, r *http.Request) {
	var str = "ok"
	
	fmt.Fprintf(w, str)
}
