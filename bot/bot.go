package bot

import (
	"io/ioutil"
	"strings"

	"../body"
)

func main() {
	token, _ := ioutil.ReadFile("token")
	Belt.Initialize(strings.TrimSpace(string(token)))
}
