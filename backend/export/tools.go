package export

import (
	"fmt"
	"net/http"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	ptime "github.com/yaa110/go-persian-calendar"
	"tools"
)

func getEsClient() *elastic.Client {
	h := http.Header{
		"Authorization": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwicm9sZSI6IkFETUlOIiwic3RhdGUiOiIiLCJleHAiOjE2MDU2MDc0ODR9.fYKmhw4N_nqYxj9GUxFf20w_8gI5SXwcNVrsrduxJqQ"},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHeaders(h),
		elastic.SetURL(tools.GetEnv("ES_HOST", "http://tavanir.example.com/gw/admin/api/v1/search/res/")),
		// elastic.SetTraceLog(log.New(os.Stdout, "trace", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}

	return client
}

func queryTimeTostring(input string) string {
	t, _ := time.Parse(time.RFC3339, input)
	pt := ptime.New(t)
	ts := fmt.Sprintf("%d-%02d-%02d", pt.Day(), pt.Month(), pt.Year())
	return ts
}

func nowTimeTostring() string {
	pt := ptime.Now(ptime.Iran())
	ts := fmt.Sprintf("%d-%02d-%02d", pt.Year(), pt.Month(), pt.Day())
	return ts
}

var (
	elasticIndexName = tools.GetEnv("ES_INDEX", "request.v1")
)

func incrementString(s string, v int) string {
	i := cast.ToInt(s)
	i = i + v

	return cast.ToString(i)
}

func nextCharString(ch rune) rune {
	if ch++; ch > 'Z' {
		return 'A'
	}
	return ch
}

func nextChar(ch rune) rune {
	if ch++; ch > 'Z' {
		return 'A'
	}
	return ch
}
