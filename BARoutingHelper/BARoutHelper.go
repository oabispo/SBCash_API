package BARoutingHelper

import "net/http"
import "strconv"
import "log"

type OnStart func( portNum int );
type OnBeforeHandling func( r *http.Request );
type OnAfterHandling func( path string, method string );

type BARoutConfig struct {
	portNum int
	onStart OnStart
	onBefore OnBeforeHandling
	onAfter OnAfterHandling
}

func NewBARoutConfig(portNum int, onStart OnStart, onBeforeHandling OnBeforeHandling, onAfterHandling OnAfterHandling ) *BARoutConfig {
	return &BARoutConfig{ portNum: portNum, onStart: onStart, onBefore: onBeforeHandling, onAfter: onAfterHandling };
}

type BARoutHelper struct {
	config *BARoutConfig
}

func NewBARoutHelper( config *BARoutConfig ) *BARoutHelper {
	return &BARoutHelper{ config: config };
}

func handlerNotAvaliable( w http.ResponseWriter, r *http.Request ) {
	w.Header().Set( "Content-Type", "text/html" );
	w.WriteHeader( http.StatusNotFound );
	w.Write( []byte("<h1>No 404 defined for / or any other subgroup </h1>") );
}

func (e *BARoutHelper) StartRouting( segmentHandler *BAUrlSegmentRoute ) {
	if ( e.config.onStart != nil ) {
		e.config.onStart( e.config.portNum );
	}

	mux := http.NewServeMux();

	if segmentHandler != nil {
		handle := func( w http.ResponseWriter, r *http.Request ) {
			segmentMng := NewBAUrlManage( r.URL.String() );
			handlerFunc := segmentHandler.Execute( r.Method, segmentMng, 0 );
			if ( handlerFunc != nil ) {
				if ( e.config.onBefore != nil ) {
					e.config.onBefore( r );
				}

				handlerFunc( w, r );

				if ( e.config.onAfter != nil ) {
					e.config.onAfter( r.URL.String(), r.Method );
				}
			} else {
				handlerNotAvaliable( w, r );
			}
		}

		mux.HandleFunc( segmentHandler.GetUrl(), handle );
	}
		
	server := &http.Server{ Addr: "localhost:" + strconv.Itoa( e.config.portNum ), Handler: mux };
	err := server.ListenAndServe();
	log.Fatal( err );
}
