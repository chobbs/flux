package discord_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/influxdata/flux"
	_ "github.com/influxdata/flux/builtin"
	"github.com/influxdata/flux/dependencies/url"
	"github.com/influxdata/flux/runtime"
)

func TestDiscord(t *testing.T) {
	var req *http.Request
	//var body []byte
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req = r
		w.WriteHeader(204)
	}))
	defer ts.Close()
	script := fmt.Sprintf(`
		import "internal/testutil"
		import "contrib/chobbs/discord"

response = discord.send(url:"%s/path/fake/fakey",username:"chobbs",content:"it's fake baby!")
response == 204 or testutil.fail()
`, ts.URL)

	ctx := flux.NewDefaultDependencies().Inject(context.Background())
	if _, _, err := runtime.Eval(ctx, script); err != nil {
		t.Fatal("evaluation of discord.send failed: ", err)
	}
	if want, got := "/path/fake/fakey", req.URL.Path; want != got {
		t.Errorf("unexpected url want: %q got: %q", want, got)
	}
	if want, got := "POST", req.Method; want != got {
		t.Errorf("unexpected method want: %q got: %q", want, got)
	}
}

func TestPost_ValidationFail(t *testing.T) {
	script := `
	import "internal/testutil"
import "contrib/chobbs/discord"
discord.send(url:"http://127.1.1.1/fakeyfake",username:"chobbs",content:"it's fake baby!")
`
	deps := flux.NewDefaultDependencies()
	deps.Deps.HTTPClient = http.DefaultClient
	deps.Deps.URLValidator = url.PrivateIPValidator{}
	ctx := deps.Inject(context.Background())
	_, _, err := runtime.Eval(ctx, script)
	if err == nil {
		t.Fatal("expected failure")
	}
	if !strings.Contains(err.Error(), "url is not valid") {
		t.Errorf("unexpected cause of failure, got err: %v", err)
	}
}
