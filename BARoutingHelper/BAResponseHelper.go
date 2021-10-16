package BARoutingHelper

import "net/http"
import "encoding/json"
import "io"

func DecodeBodyToStruct( body io.ReadCloser, newStructCallback func( params ...interface{} ) ( interface{} ), params ...interface{} ) ( interface {} )  {
	// chamar panic( "" ) para casos em que não tenha o callback
	if ( newStructCallback != nil ) {
		if obj := newStructCallback( params ); ( obj != nil ) {
			d := json.NewDecoder( body );
			d.Decode( obj );
			return obj;
		}
	}
		
	return nil;
}

func responseTextAs( w http.ResponseWriter, contentType string, httpCode int, bodyData interface{} ) {
	response, _ := json.Marshal( bodyData )

	w.Header().Set( "Content-Type", contentType );
	w.WriteHeader( httpCode );
	w.Write( response );
}

func ResponseTextAsJSON( w http.ResponseWriter, httpCode int, bodyData interface{} )  {
	responseTextAs( w, "application/json", httpCode, bodyData );
}

func ResponseTextAsHTML( w http.ResponseWriter, httpCode int, bodyData interface{} )  {
	responseTextAs( w, "text/html", httpCode, bodyData );
}

func responseByteAs( w http.ResponseWriter, contentType string, httpCode int, bodyData interface{} ) {
	var e *json.Encoder = json.NewEncoder( w );

	w.Header().Set( "Content-Type", contentType );
	w.WriteHeader( httpCode );
	e.Encode( bodyData )
}

func ResponseByteAsJSON( w http.ResponseWriter, httpCode int, bodyData interface{} )  {
	responseByteAs( w, "application/json", httpCode, bodyData );
}