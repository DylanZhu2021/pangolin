package AhoCorasickDoubleArrayTrie

import "github.com/RoaringBitmap/roaring"

type InvertedIndex struct {
	Key   string
	Value *roaring.Bitmap
}
