package BARoutingHelper

import "net/url"
import "strings"

type BAUrlKeyValue struct {
	operation string
	value string
}

type BAUrlQueryManager interface {
	AddKeyValue( key string, value string )
	GetKeyValue( key string ) ( string )	
}

type BAUrlQueryManage struct {
	keys map[string]*BAUrlKeyValue
}

func NewBAUrlQueryManage( url *url.URL ) ( *BAUrlQueryManage )  {
	uqm := &BAUrlQueryManage{};
	uqm.keys = make( map[string]*BAUrlKeyValue );
	kvs := strings.Split( strings.TrimSpace( url.RawQuery ), "&");
	for _, item := range kvs {
		kv := strings.Split( item, "=" );
		if ( len(kv) > 1 ) {
			uqm.AddKeyValue( kv[0], kv[1] );
		}
	}

	return uqm;
}

func (km *BAUrlQueryManage) AddKeyValue( key string, value string ) {
	km.keys[ key ] = &BAUrlKeyValue{ operation: "", value: value };
}

func (km *BAUrlQueryManage) GetKeyValue( key string ) ( string, bool ) {
	keyValue := km.keys[ key ]; 
	found := ( keyValue != nil );
	if ( found ) {
		return keyValue.value, found;
	}

	return "", false;
}
