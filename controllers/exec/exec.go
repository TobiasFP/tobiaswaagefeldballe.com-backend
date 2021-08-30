package exec

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func RunFortranSpeedtest(ctx *gin.Context) {
	out, err := exec.Command("exec/speedtest_fortran").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	ctx.JSON(http.StatusOK, output)
}

func RunGolangSpeedtest(ctx *gin.Context) {
	out, err := exec.Command("exec/speedtest_golang").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	ctx.JSON(http.StatusOK, output)
}
