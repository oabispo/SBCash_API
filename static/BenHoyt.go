package endpoint

import (
	"fmt"
	"net/http"

	router "sbcash_api/BARoutingHelper"
)

func NewBHMainRouter() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/", nil)
	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>Home</h1>"))
		fmt.Fprint(w)
	}
	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>404</h1>"))
		fmt.Fprint(w)
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	segment.AddSubRouter(NewBHContact())
	segment.AddSubRouter(NewBHAPIWidgets())
	segment.AddSubRouter(NewBHSlugPorId())

	return segment
}

func NewBHContact() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/contact", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>Contact</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	return segment
}

func NewBHAPIWidgets() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/api/widgets", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>API Widgets - GET</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	segment.AddSubRouter(NewBHAPIWidgetsPorId())

	return segment
}

func NewBHAPIWidgetsPorId() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/api/widgets/([0-9]+)", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>API Widgets por Id - GET</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	segment.AddSubRouter(NewBHAPIWidgetsPorIdParts())

	return segment
}

func NewBHAPIWidgetsPorIdParts() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/api/widgets/([0-9]+)/parts", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>API Widgets por Id Showing all parts - GET</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	segment.AddSubRouter(NewBHAPIWidgetsPorIdPartPorId())

	return segment
}

func NewBHAPIWidgetsPorIdPartPorId() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/api/widgets/([0-9]+)/parts/([0-9]+)", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>API Widgets por Id Showing part por Id - GET</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	return segment
}

func NewBHSlugPorId() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/([0-9]+)", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>SLUG por Id - GET</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	segment.AddSubRouter(NewBHSlugPorIdAdmin())
	segment.AddSubRouter(NewBHSlugPorIdImage())

	return segment
}

func NewBHSlugPorIdAdmin() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/([0-9]+)/admin", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>SLUG por Id | Admin - GET</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	return segment
}

func NewBHSlugPorIdImage() *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/([0-9]+)/image", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte("<h1>SLUG por Id | Image - Post</h1>"))
		fmt.Fprint(w)
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))

	return segment
}
