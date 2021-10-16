package BARoutingHelper

import "net/http"
import "regexp"

type BAHandlerMode int
type BAMethodType int

const (
	HM_OK BAHandlerMode = iota // handler for correct segment
	HM_NA // handler for not available
	HM_NF // handler for not found
)

const (
	MT_GET BAMethodType = iota
	MT_POST
	MT_PUT
	MT_DELETE
	MT_OPTIONS
	MT_PATCH
	MT_HEAD
	// ...
)

type BAUrlSegmentRouterFunc interface {
	GetMode() BAHandlerMode
	ServeHTTP( w http.ResponseWriter, r *http.Request )
}

type BAUrlSegmentServeHTTP struct {
	mode BAHandlerMode
	serveHTTP func(w http.ResponseWriter, r *http.Request )
}

func NewBASegmentServeHTTP( mode BAHandlerMode, handler func(w http.ResponseWriter, r *http.Request ) ) ( *BAUrlSegmentServeHTTP ) {	
	return &BAUrlSegmentServeHTTP{ mode: mode, serveHTTP: handler }
}

func ( ssh *BAUrlSegmentServeHTTP ) GetMode() ( BAHandlerMode ) {
	return ssh.mode;
}

func ( ssh *BAUrlSegmentServeHTTP ) ServeHTTP( w http.ResponseWriter, r *http.Request ) {
	ssh.serveHTTP( w, r );
}

type BAUrlSegmentRouter interface {
	GetMethod() ( string )
	GetUrl() ( string )
	IsRelativePath() ( bool )
	GetSubRouter( method BAMethodType, segmentPath string ) ( BAUrlSegmentRouter )
	AddSubRouter( handler BAUrlSegmentRouter )
	GetSubRouterCount() ( int )
	AddHandlerFunc( handler BAUrlSegmentRouterFunc )
	GetHandlerFunc( mode BAHandlerMode ) ( BAUrlSegmentRouterFunc )
	Execute( method string, url BAUrlManager, segmentIndex int ) ( http.HandlerFunc )
}

type BAUrlSegmentRoute struct {
	method BAMethodType
	isRelativePath bool
	segmentPath string
	subSegmentHandler []BAUrlSegmentRouter
	handleFunc []BAUrlSegmentRouterFunc
}

func NewBAUrlSegmentRoute( method BAMethodType, segmentPath string, segmentParent *BAUrlSegmentRoute ) ( *BAUrlSegmentRoute ) {
	var ppath string;
	isRelative := ( segmentParent != nil );
	if ( isRelative ) {
		if ( segmentParent.GetUrl() != "/" ) {
			ppath = segmentParent.GetUrl();
		}
	}

	return &BAUrlSegmentRoute{ method: method, segmentPath: ppath + segmentPath, isRelativePath: isRelative, subSegmentHandler: make( []BAUrlSegmentRouter, 0 ), handleFunc: make( []BAUrlSegmentRouterFunc, 0 ) };
}

func ( sh *BAUrlSegmentRoute ) GetMethod() ( string ) {
	return getMethod( sh.method );
}

func getMethod( method BAMethodType ) ( string ) {
	switch method {
		case MT_POST: return "POST"
		case MT_PUT: return "PUT"
		case MT_DELETE: return "DELETE"
		case MT_OPTIONS: return "OPTIONS"	
		case MT_HEAD: return "HEAD"
		case MT_PATCH: return "PATCH"
		default : return "GET"
	}
}

func ( sh *BAUrlSegmentRoute ) GetUrl() ( string ) {
	return sh.segmentPath;
}

func ( sh *BAUrlSegmentRoute )	IsRelativePath() ( bool ) {
	return sh.isRelativePath;
}

func ( sh *BAUrlSegmentRoute ) GetSubRouter( method BAMethodType, segmentPath string ) ( BAUrlSegmentRouter ) {
	var handler BAUrlSegmentRouter = nil;
	for _, item := range sh.subSegmentHandler {
		if ( ( item.GetMethod() == getMethod( method ) ) && ( item.GetUrl() == segmentPath ) ) {
			handler = item;
			break;
		}
	}

	return handler;
}

func ( sh *BAUrlSegmentRoute ) AddSubRouter( handler BAUrlSegmentRouter ) {
	sh.subSegmentHandler = append( sh.subSegmentHandler, handler );
}

func ( sh *BAUrlSegmentRoute ) GetSubRouterCount() (int) {
	return len(sh.subSegmentHandler);
}

func ( sh *BAUrlSegmentRoute ) AddHandlerFunc( handler BAUrlSegmentRouterFunc ) {
	sh.handleFunc = append( sh.handleFunc, handler );
}

func ( sh *BAUrlSegmentRoute ) GetHandlerFunc( mode BAHandlerMode ) ( BAUrlSegmentRouterFunc ) {
	var result BAUrlSegmentRouterFunc;
	for _, item := range sh.handleFunc {
		if ( item.GetMode() == mode ) {
			result = item;
			break;
		}
	}
	return result
}

func ( sh *BAUrlSegmentRoute ) isHandlerCandidate( method string, url BAUrlManager, segmentIndex int ) ( bool ) { 	
	if isThatMethod := ( sh.GetMethod() == method ); ( !isThatMethod ) {
		return false;
	}

	var candidate bool = false;

	url.HeadToSegment( segmentIndex );

	urlBaseSegmentPath := NewBAUrlManage( sh.GetUrl() );
	urlBaseSegmentPath.HeadToSegment( segmentIndex );
	if lhe := ( urlBaseSegmentPath.SegmentCount() == url.SegmentCount() ); ( lhe ) {
		var isOK bool = true;
		for i:= segmentIndex; i < urlBaseSegmentPath.SegmentCount(); i++ {
			regex := urlBaseSegmentPath.CurrentSegment();
			check := url.CurrentSegment();
			isOK = ( check == regex );
			if ( !isOK ) {
				hasRegex, err := regexp.MatchString( regex, check );
				isOK = hasRegex;
				if ( ( err != nil ) || ( !hasRegex ) ) {
					break;
				}
			}
			url.ForwardPathSegment();
			urlBaseSegmentPath.ForwardPathSegment();
		}

		candidate = isOK;
	} 

	return candidate;
}

func isHandlerRelative( sh BAUrlSegmentRouter, method string, url BAUrlManager ) ( bool, int ) {
	shPath := NewBAUrlManage( sh.GetUrl() );

	url.HeadToSegment( 0 );
	var isOK bool = true;
	var position int = 0;
	for i := 0; i < ( shPath.SegmentCount() - 1 ); i++ {
		regex := shPath.CurrentSegment();
		check := url.CurrentSegment();
		isOK = ( check == regex );
		if ( !isOK ) {
			hasRegex, err := regexp.MatchString( regex, check );
			isOK = hasRegex;
			if ( ( err != nil ) || ( !hasRegex ) ) {
				break;
			}
		}
		position++;
		url.ForwardPathSegment();
		shPath.ForwardPathSegment();
	}

	return isOK, position;
}

func ( sh *BAUrlSegmentRoute ) Execute( method string, url BAUrlManager, segmentIndex int ) ( http.HandlerFunc ) {
	isCandidate := sh.isHandlerCandidate( method, url, segmentIndex );

	if ( isCandidate ) {
		var handleFunc http.HandlerFunc = nil;
		if ok_hf := sh.GetHandlerFunc( HM_OK );( ok_hf != nil ) {
			handleFunc = ok_hf.ServeHTTP;
		}

		return handleFunc;
	} else {
		if ( sh.GetSubRouterCount() > 0 ) {
			var handleFunc http.HandlerFunc = nil;
			var isRelative bool;
			var index int;
			for _, item := range sh.subSegmentHandler {
				if ( item.IsRelativePath() ) {
					if isRelative, index = isHandlerRelative( item, method, url ); ( !isRelative ) {
						continue;
					}
				}

				handleFunc = item.Execute( method, url, index );
				if ( handleFunc != nil ) { 
					return handleFunc; 
				}
			}

			if ( handleFunc == nil ) {
				if isRelative, _ = isHandlerRelative( sh, method, url ); ( isRelative ) {
					if nf_hf := sh.GetHandlerFunc( HM_NA ); ( nf_hf != nil ) { 
						return nf_hf.ServeHTTP; 
					}
				}
			}
		}

		return nil;
	} 
}

