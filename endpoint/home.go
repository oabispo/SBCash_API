package endpoint

import (
	"database/sql"
	"net/http"

	"fmt"
	router "sbcash_api/BARoutingHelper"
)

func NewRouterHome(db *sql.DB) *router.BAUrlSegmentRoute {

	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Home</h1>"))
	}
	handler_na := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<h1>404</h1>"))
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NA, handler_na))

	segment.AddSubRouter(NewRouterFavIcon(segment))
	segment.AddSubRouter(NewRouterAPI(db, segment))
	segment.AddSubRouter(NewRouterIndex(db, segment))
	segment.AddSubRouter(NewRouterStaticFiles(db, segment))

	return segment
}

func NewRouterIndex(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/static", segmentParent)

	handler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	return segment
}

func NewRouterStaticFiles(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/static/(\\w)(.{1})(\\w)", segmentParent)
	//	handler := http.FileServer( http.Dir("./static") );
	//	h := http.Handle( "/static", handler );

	handler := func(w http.ResponseWriter, r *http.Request) {
		fileName := fmt.Sprintf(".%v", r.URL.String())
		http.ServeFile(w, r, fileName)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	return segment
}

func NewRouterAPI(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/api", segmentParent)

	segment.AddSubRouter(NewRouterV1(db, segment))

	return segment
}

func NewRouterV1(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/v1", segmentParent)

	segment.AddSubRouter(NewRouterUsuario(db, segment))
	segment.AddSubRouter(NewRouterConfig(db, segment))
	segment.AddSubRouter(NewRouterImposto(db, segment))
	segment.AddSubRouter(NewRouterTipoMovimento(db, segment))

	return segment
}

func NewRouterFavIcon(segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {

	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/favicon.ico", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>Favicon.ico</h1>"))
	}
	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>404</h1>"))
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}
