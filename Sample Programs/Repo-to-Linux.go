package linux_transformation

include {

	"fmt"

 	"log"

	"os/exec"

	"runtime"
}

func start() {

	cmd := exec.Command("go get -u github.com/gorilla/mux")
	cmd := exec.Command("curl -LJO https://github.com/joyent/node/tarball/v0.7.1")
}

func main(){

      go start()
}

