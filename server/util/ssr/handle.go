package ssr

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/parakeet/server/models"
	"github.com/mohemohe/parakeet/server/util"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Handle(c echo.Context, pool *Pool) error {
	defer func() {
		if r := recover(); r != nil {
			_ = c.Render(http.StatusInternalServerError, "ssr.html", Result{
				Error: r.(string),
			})
		}
	}()

	js := pool.get()
	defer pool.release(js)

	if !js.Callable {
		if js.err != nil {
			return c.Render(http.StatusInternalServerError, "ssr.html", Result{
				Error: js.err.Error(),
				Title: "parakeet",
				Meta:  "",
			})
		}
		return c.Render(http.StatusInternalServerError, "ssr.html", Result{
			Error: "invalid js",
			Title: "parakeet",
			Meta:  "",
		})
	}

	title := "parakeet"
	titleKV := models.GetKVS("site_title")
	if titleKV != nil && titleKV.Value != "" {
		title = titleKV.Value.(string)
	}
	initialState, entry := initializeState(c)
	ch := js.Exec(map[string]interface{}{
		"url":     c.Request().URL.String(),
		"title":   title,
		"headers": map[string][]string(c.Request().Header),
		"state":   initialState,
	})

	select {
	case res := <-ch:
		res.Title = title
		if js.err != nil {
			util.Logger().WithField("error", res.Error).Errorln("js eval error")
			return c.Render(http.StatusInternalServerError, "ssr.html", Result{
				Error: js.err.Error(),
				Title: res.Title,
				Meta:  "",
			})
		}
		if len(res.Error) == 0 {
			if entry.Title != "" {
				res.Title = entry.Title + " - " + res.Title
			} else {

			}
			return c.Render(http.StatusOK, "ssr.html", res)
		} else {
			util.Logger().WithField("error", res.Error).Errorln("js result error")
			return c.Render(http.StatusInternalServerError, "ssr.html", res)
		}
	case <-time.After(3 * time.Second):
		util.Logger().Errorln("cant keep up!")
		return c.Render(http.StatusInternalServerError, "ssr.html", Result{
			Error: "timeout",
			Title: title,
			Meta:  "",
		})
	}
}

func initializeState(c echo.Context) (map[string]interface{}, *models.Entry) {
	path := c.Request().URL.Path

	entries := &models.Entries{}
	entry := &models.Entry{
		Tag: make([]string, 0),
	}
	if path == "/" {
		entries = models.GetEntries(10, 1)
	}
	if strings.HasPrefix(path, "/entries/") {
		r := regexp.MustCompile(`^/entries/(\d+)`).FindAllStringSubmatch(path, -1)
		if len(r) == 1 && len(r[0]) == 2 {
			if page, err := strconv.Atoi(r[0][1]); err == nil {
				entries = models.GetEntries(10, page)
			}
		}
	}
	if strings.HasPrefix(path, "/entry/") {
		r := regexp.MustCompile(`^/entry/(.*)`).FindAllStringSubmatch(path, -1)
		if len(r) == 1 && len(r[0]) == 2 {
			entry = models.GetEntryById(r[0][1])
		}
	}

	return map[string]interface{}{
		"entryStore": map[string]interface{}{
			"entries":  toJson(entries.Entries),
			"paginate": toJson(entries.Info),
			"entry":    toJson(entry),
		},
	}, entry
}

func toJson(t interface{}) (result string) {
	defer func() {
		if err := recover(); err != nil {
			util.Logger().WithField("target", t).Warn("toJson failed")
			result = "{}"
		}
	}()
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	result = string(b)
	return
}
