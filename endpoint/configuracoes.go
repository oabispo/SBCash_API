package endpoint

import (
	"database/sql"
	"encoding/json"
	"net/http"

	router "sbcash_api/BARoutingHelper"

	config_sql "sbcash_api/model/mySql/richData"

	repo "sbcash_api/repositories/rich"
)

func NewRouterConfig(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/configuracoes", segmentParent)
	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		NOTFOUNDResponse(w, nil)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		item, err := config_sql.Estabelecimento_GetInfo(db)
		if err == nil {
			GETResponse(w, item)
		}
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	segment.AddSubRouter(NewRouterConfig_Update(db, segment))
	segment.AddSubRouter(NewRouterCupom(db, segment))
	segment.AddSubRouter(NewRouterCupom_Update(db, segment))
	segment.AddSubRouter(NewRouterCortesia(db, segment))
	segment.AddSubRouter(NewRouterCortesia_Update(db, segment))

	return segment
}

func NewRouterConfig_Update(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_PUT, "", segmentParent)
	handler_nf := func(w http.ResponseWriter, r *http.Request) {
		NOTFOUNDResponse(w, nil)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		decode := json.NewDecoder(r.Body)
		e := repo.NewEstabelecimento()
		decode.Decode(&e)

		err := config_sql.Estabelecimento_UpdateInfo(db, e.Nome, e.RazaoSocial, e.CNPJ, e.Endereco, e.Cidade, e.Data_Instalacao, e.Hora_Abre, e.Hora_Fecha, e.Tempo_login_Expira, e.NumViasDescricao, e.Impressao_parcial, e.Idioma)

		if err == nil {
			PUTResponse(w, nil)
		} else {
			ERRORResponse(w, nil)
		}
	}
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler))
	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_NF, handler_nf))

	segment.AddSubRouter(NewRouterCupom(db, segment))
	segment.AddSubRouter(NewRouterCortesia(db, segment))

	return segment
}

func NewRouterCupom(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/cupom", segmentParent)

	handler_ok := func(w http.ResponseWriter, r *http.Request) {

		if items, err := config_sql.Cupom_GetInfo(db); err == nil {
			GETResponse(w, items)
		} else {
			ERRORResponse(w, nil)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))
	return segment
}

func NewRouterCupom_Update(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_PUT, "/cupom", segmentParent)

	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		decode := json.NewDecoder(r.Body)
		c := repo.NewCupom()
		decode.Decode(&c)

		if err := config_sql.Cupom_UpdateInfo(db, c.LinhaA, c.LinhaB, c.DiasVigencia, c.Operacao, c.ValorMaxVenda); err == nil {
			PUTResponse(w, nil)
		} else {
			ERRORResponse(w, nil)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))
	return segment
}

func NewRouterCortesia(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_GET, "/cortesia", segmentParent)

	handler_ok := func(w http.ResponseWriter, r *http.Request) {

		if items, err := config_sql.Cortesia_GetInfo(db); err == nil {
			GETResponse(w, items)
		} else {
			ERRORResponse(w, nil)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))

	return segment
}

func NewRouterCortesia_Update(db *sql.DB, segmentParent *router.BAUrlSegmentRoute) *router.BAUrlSegmentRoute {
	segment := router.NewBAUrlSegmentRoute(router.MT_PUT, "/cortesia", segmentParent)

	handler_ok := func(w http.ResponseWriter, r *http.Request) {
		decode := json.NewDecoder(r.Body)
		c := repo.NewCortesia()
		decode.Decode(&c)

		if err := config_sql.Cortesia_UpdateInfo(db, c.ValorMaximo, c.Vigencia, c.RegraCortesia, c.Valor1, c.Valor2, c.Valor3, c.Valor4, c.Valor5); err == nil {
			PUTResponse(w, nil)
		} else {
			ERRORResponse(w, nil)
		}
	}

	segment.AddHandlerFunc(router.NewBASegmentServeHTTP(router.HM_OK, handler_ok))

	return segment
}
