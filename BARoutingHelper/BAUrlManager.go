package BARoutingHelper

import "fmt"
import "strings"

type BAUrlManager interface {
	ForwardPathSegment() ( bool )
	BackPathSegment() ( bool )
	CurrentSegment() ( string )
	UrlToCurrentSegment() ( string )
	LastSegment() ( string )
	SegmentCount() ( int )
	HeadToSegment( index int ) ( bool )
	String() ( string )
}

type BAUrlManage struct {
	segments []string
	index int
}

func NewBAUrlManage( url string ) ( *BAUrlManage ) {
	var seg []string = make( []string, 0 );
	for _, item := range strings.Split( url, "/") {
		if ( item != "" ) { seg = append( seg, item ); }
	}

	result := &BAUrlManage{ segments: seg, index: 0 };

	if ( len( result.segments ) == 0 ) { result.index = -1; }
	return result;
}

func (us *BAUrlManage) ForwardPathSegment() ( bool ) {
	result := (us.index < ( us.SegmentCount() - 1 ) );
	if ( result ) {
		us.index++;		
	}
	return result;
}

func (us *BAUrlManage) BackPathSegment() ( bool ) {
	result := (us.index > 0 );
	if ( result ) {
		us.index--;
	}
	return result;
}

func (us *BAUrlManage) CurrentSegment() ( string ) {
	if ( us.index != -1 ) {
		return us.segments[us.index];
	} else {
		return "";
	}
}

func (us *BAUrlManage) UrlToCurrentSegment() ( string ) {
	segmentIndex := us.index;

	var sb strings.Builder;
	if ( ( segmentIndex + 1 ) > 0 ) {
		for i:= 0; i < ( segmentIndex + 1 ); i++ {
			sb.WriteString( fmt.Sprintf("/%v", us.CurrentSegment() ) );
		}
	} else {
		sb.WriteString( "/" );
	}

	return sb.String();
}

func (us *BAUrlManage) LastSegment() ( string ) {
	if ( us.SegmentCount() > 0 ) {
		return us.segments[us.SegmentCount() - 1];
	} else {
		return "";
	}
}

func (us *BAUrlManage) HeadToSegment( index int ) ( bool ) {
	result := ( us.index > -1 ) && ( index >= 0 ) && ( index < us.SegmentCount() );
	if ( result ) {
		us.index = index;
	}

	return result;
}

func (us *BAUrlManage) SegmentCount() ( int ) {
	return len( us.segments );
}

func (us *BAUrlManage) String() ( string ) {
	var sb strings.Builder;
	for _, item := range us.segments {
		sb.WriteString( fmt.Sprintf( "/%v", item ) );
	}

	if ( len( us.segments ) == 0 ) { sb.WriteString( fmt.Sprintf( "/" ) ); }

	return sb.String();
}
