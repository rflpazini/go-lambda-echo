package healthcheck

import (
	"net/http"
	"os"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/rflpazini/articles/lambda/pkg/utils"
)

// Handler api handler
func Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func InfoHandler(c echo.Context) error {
	var commitHash string
	var goVersion string

	if bi, ok := debug.ReadBuildInfo(); ok {
		goVersion = bi.GoVersion

		for _, kv := range bi.Settings {
			if kv.Key == "vcs.revision" {
				commitHash = kv.Value
				break
			}
		}
	}

	rsp := Response{
		App: App{
			Name:      "test-lambda-echo-framework",
			Version:   "v1.0.0",
			GoVersion: goVersion,
			Codebase: &Codebase{
				CommitHash: commitHash,
				Branch:     os.Getenv("BRANCH_NAME"),
			},
			Environment: &Environment{
				Name:       utils.GetHostName(),
				Region:     os.Getenv("EC2_REGION"),
				InstanceId: os.Getenv("INSTANCE_ID"),
			},
		},
	}

	return c.JSON(http.StatusOK, rsp)
}
