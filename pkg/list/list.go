package list

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/safeguardproperties/gocontentservice/pkg/contentservice"
	ora "gopkg.in/rana/ora.v4"
)

// Info holds all the information required for listing the imagedetails
type Info struct {
	OrderNumber     int64
	Db              *ora.Ses
	MemClient       *memcache.Client
	SecondsToExpiry int32
}

// ImageDetailList contains interfaces for mocking
type ImageDetailList struct {
	DatabaseCaller
	CacheCaller
}

// DatabaseCaller contains methods to be performed on database
type DatabaseCaller interface {
	GetImageDetailsFromDatabase() ([]*contentservice.ImageDetail, error)
}

//CacheCaller calls methods on memcache
type CacheCaller interface {
	GetImageDetailsFromCache() ([]*contentservice.ImageDetail, error)
	SetImageDetailsToCache([]*contentservice.ImageDetail) error
}
