package endpoint

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	router "sbcash_api/BARoutingHelper"

	imposto_sql "sbcash_api/model/mySql/plainData"

	repo "sbcash_api/repositories/plain"
)

func NewRouterImposto(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/imposto", segmentParent)

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		NOTFOUNDResponse(w, nil)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		items, err := imposto_sql.Imposto_GetAll(db)
		if err == nil {
			if len(items) < 1 {
				handler_nf(w, r)
			} else {
				GETResponse(w, items)
			}
		} else {
			ERRORResponse(w, err)
		}
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	segment.AddSubRouter(NewRouterImposto_GetById(db, segment))
	segment.AddSubRouter(NewRouterImposto_Post(db, segment))
	segment.AddSubRouter(NewRouterImposto_Update(db, segment))
	segment.AddSubRouter(NewRouterImposto_Delete(db, segment))

	return segment
}

func NewRouterImposto_GetById(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/([0-9]+)", segmentParent)

	var cod_imposto int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_imposto))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_imposto, _ = strconv.Atoi(basm.LastSegment())

		items, err := imposto_sql.Imposto_GetByID(db, cod_imposto)
		if err == nil {
			if items == nil {
				handler_nf(w, r)
			} else {
				GETResponse(w, items)
			}
		} else {
			ERRORResponse(w, err)
		}
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}

func NewRouterImposto_Post(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_POST, "", segmentParent)

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		ERRORResponse(w, nil)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {

		decode := json.NewDecoder(r.Body)
		imposto := repo.NewImposto()
		decode.Decode(&imposto)

		id, err := imposto_sql.Imposto_Inserir(db, imposto.Descricao, imposto.Aliquota)
		if err == nil {
			strId := strconv.Itoa(int(id))
			POSTResponse(w, map[string]string{"id": strId})
		} else {
			handler_nf(w, r)
		}
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}

func NewRouterImposto_Delete(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_DELETE, "/([0-9]+)", segmentParent)
	var cod_imposto int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_imposto))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_imposto, _ = strconv.Atoi(basm.LastSegment())

		rowsAffected, err := imposto_sql.Imposto_Remover(db, cod_imposto)
		if (err == nil) && (rowsAffected > 0) {
			var response map[string]string = make(map[string]string)
			response["id"] = strconv.Itoa(cod_imposto)
			response["totalRemovido"] = strconv.Itoa(int(rowsAffected))
			DELETEResponse(w, response)
		} else {
			handler_nf(w, r)
		}
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}

func NewRouterImposto_Update(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_PUT, "/([0-9]+)", segmentParent)
	var cod_imposto int

	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		strId := strconv.Itoa(int(cod_imposto))
		NOTFOUNDResponse(w, map[string]string{"id": strId})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		basm := router.NewBAUrlManage(r.URL.String())
		cod_imposto, _ = strconv.Atoi(basm.LastSegment())

		decode := json.NewDecoder(r.Body)
		imposto := repo.NewImposto()
		decode.Decode(&imposto)

		rowsAffected, err := imposto_sql.Imposto_Atualizar(db, imposto.Descricao, imposto.Aliquota, cod_imposto)
		if err == nil {
			if rowsAffected > 0 {
				strId := strconv.Itoa(cod_imposto)
				PUTResponse(w, map[string]string{"id": strId})
			} else {
				handler_nf(w, r)
			}
		} else {
			data := map[string]string{"id": strconv.Itoa(cod_imposto)}
			data["message"] = err.Error()
			ERRORResponse(w, data)
		}

	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	return segment
}
