package geecache

import pb "github.com/SeanChan0901/gee-cache/geecache/geecachepb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implement by a peer
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
