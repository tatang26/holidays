package holiday

import (
	"net/http"

	"go.leapkit.dev/core/render"
	"go.leapkit.dev/core/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	holidays, err := holidaysFor(r.Context(), params.Get("countryCode"))
	if err != nil {
		server.Error(w, err, http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("holidays", holidays)
	rw.RenderClean("holiday/holidays.html")
}
