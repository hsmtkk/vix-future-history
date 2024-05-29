package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("EntryPoint", entryPoint)
}

func entryPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
