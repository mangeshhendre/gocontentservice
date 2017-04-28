package list

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/divyag9/gocontentservice/pkg/contentservice"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"gopkg.in/rana/ora.v4"
	_ "gopkg.in/rana/ora.v4"
)

type FakeInfo struct {
	orderNumber  int
	imageDetails []*contentservice.ImageDetail
}

func (f *FakeInfo) GetImageDetailsFromDatabase() ([]*contentservice.ImageDetail, error) {
	if f.orderNumber == 0 {
		return nil, errors.New("Invalid OrderNumber")
	}
	return f.imageDetails, nil
}

func (f *FakeInfo) GetImageDetailsFromCache() ([]*contentservice.ImageDetail, error) {
	return f.imageDetails, nil
}

var cases = []struct {
	fakeInfo             *FakeInfo
	info                 *Info
	expectedImageDetails []*contentservice.ImageDetail
	expectedErr          error
}{
	{
		fakeInfo: &FakeInfo{
			orderNumber: 600016555,
			imageDetails: []*contentservice.ImageDetail{&contentservice.ImageDetail{
				Archived:               "",
				Category:               "foo",
				ContractorId:           72494,
				DateCreated:            testDate,
				DateModified:           testDate,
				ImageUTCDate:           nil,
				ImageTakenDate:         nil,
				DeptCode:               "01",
				DescPrefix:             "foo",
				DescText:               "foo bar",
				FileSize:               180,
				ImageId:                3001240405,
				ImageFileName:          "C:\\Temp\\images600\\016\\555\\76215592-b810-48f0-a9e2-ac681ab0ea38.png",
				ImageHeight:            100,
				ImageType:              1,
				ImageRotated:           false,
				ImageWidth:             100,
				OrderNumber:            600016555,
				ReleaseDate:            testDate,
				ScanDate:               testDate,
				ThumbnailSize:          0,
				WebFileName:            "images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png",
				MimeType:               "image/png",
				GeneratedImageFilePath: "https://sbimage.sgpdev.com/images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png?d794b9a5-02ac-4b86-be81-cbbf0d22abf7",
				Guid: "",
			}},
		},
		info: &Info{
			OrderNumber: 600016554,
			MemClient:   getMemcacheClient(),
		},
		expectedImageDetails: []*contentservice.ImageDetail{&contentservice.ImageDetail{
			Archived:               "",
			Category:               "foo",
			ContractorId:           72494,
			DateCreated:            testDate,
			DateModified:           testDate,
			ImageUTCDate:           nil,
			ImageTakenDate:         nil,
			DeptCode:               "01",
			DescPrefix:             "foo",
			DescText:               "foo bar",
			FileSize:               180,
			ImageId:                3001240405,
			ImageFileName:          "C:\\Temp\\images600\\016\\555\\76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			ImageHeight:            100,
			ImageType:              1,
			ImageRotated:           false,
			ImageWidth:             100,
			OrderNumber:            600016555,
			ReleaseDate:            testDate,
			ScanDate:               testDate,
			ThumbnailSize:          0,
			WebFileName:            "images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			MimeType:               "image/png",
			GeneratedImageFilePath: "https://sbimage.sgpdev.com/images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png?d794b9a5-02ac-4b86-be81-cbbf0d22abf7",
			Guid: "",
		}},
		expectedErr: nil,
	},
	{
		fakeInfo: &FakeInfo{
			orderNumber:  0,
			imageDetails: []*contentservice.ImageDetail{},
		},
		info: &Info{
			OrderNumber: 600016554,
			MemClient:   getMemcacheClient(),
		},
		expectedImageDetails: []*contentservice.ImageDetail{&contentservice.ImageDetail{
			Archived:               "",
			Category:               "foo",
			ContractorId:           72494,
			DateCreated:            testDate,
			DateModified:           testDate,
			ImageUTCDate:           nil,
			ImageTakenDate:         nil,
			DeptCode:               "01",
			DescPrefix:             "foo",
			DescText:               "foo bar",
			FileSize:               180,
			ImageId:                3001240405,
			ImageFileName:          "C:\\Temp\\images600\\016\\555\\76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			ImageHeight:            100,
			ImageType:              1,
			ImageRotated:           false,
			ImageWidth:             100,
			OrderNumber:            600016555,
			ReleaseDate:            testDate,
			ScanDate:               testDate,
			ThumbnailSize:          0,
			WebFileName:            "images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			MimeType:               "image/png",
			GeneratedImageFilePath: "https://sbimage.sgpdev.com/images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png?d794b9a5-02ac-4b86-be81-cbbf0d22abf7",
			Guid: "",
		}},
		expectedErr: nil,
	},
	{
		fakeInfo: &FakeInfo{
			orderNumber:  0,
			imageDetails: []*contentservice.ImageDetail{},
		},
		info: &Info{
			OrderNumber: 600016556,
			MemClient:   getMemcacheClient(),
		},
		expectedErr: errors.New("Invalid OrderNumber"),
	},
}

var imageDetailDatabaseCases = []struct {
	Info                 *Info
	expectedImageDetails []*contentservice.ImageDetail
	expectedError        error
}{
	{
		Info: &Info{
			OrderNumber: int64(600016555),
			Db:          getDb(),
		},
		expectedImageDetails: []*contentservice.ImageDetail{{
			Archived:               "",
			Category:               "foo",
			ContractorId:           72494,
			DateCreated:            testDate,
			DateModified:           testDate,
			ImageUTCDate:           nil,
			ImageTakenDate:         nil,
			DeptCode:               "01",
			DescPrefix:             "foo",
			DescText:               "foo bar",
			FileSize:               180,
			ImageId:                3001240405,
			ImageFileName:          "C:\\Temp\\images600\\016\\555\\76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			ImageHeight:            100,
			ImageType:              1,
			ImageRotated:           false,
			ImageWidth:             100,
			OrderNumber:            600016555,
			ReleaseDate:            testDate,
			ScanDate:               testDate,
			ThumbnailSize:          0,
			WebFileName:            "images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			MimeType:               "image/png",
			GeneratedImageFilePath: "https://sbimage.sgpdev.com/images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png?d794b9a5-02ac-4b86-be81-cbbf0d22abf7",
			Guid: "",
		}},
		expectedError: nil,
	},
}

var imageDetailCacheCases = []struct {
	Info                 *Info
	expectedImageDetails []*contentservice.ImageDetail
	expectedError        error
}{
	{
		Info: &Info{
			OrderNumber: 600016554,
			MemClient:   getMemcacheClient(),
		},
		expectedImageDetails: []*contentservice.ImageDetail{{
			Archived:               "",
			Category:               "foo",
			ContractorId:           72494,
			DateCreated:            testDate,
			DateModified:           testDate,
			ImageUTCDate:           nil,
			ImageTakenDate:         nil,
			DeptCode:               "01",
			DescPrefix:             "foo",
			DescText:               "foo bar",
			FileSize:               180,
			ImageId:                3001240405,
			ImageFileName:          "C:\\Temp\\images600\\016\\555\\76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			ImageHeight:            100,
			ImageType:              1,
			ImageRotated:           false,
			ImageWidth:             100,
			OrderNumber:            600016555,
			ReleaseDate:            testDate,
			ScanDate:               testDate,
			ThumbnailSize:          0,
			WebFileName:            "images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png",
			MimeType:               "image/png",
			GeneratedImageFilePath: "https://sbimage.sgpdev.com/images/600/016/555/76215592-b810-48f0-a9e2-ac681ab0ea38.png?d794b9a5-02ac-4b86-be81-cbbf0d22abf7",
			Guid: "",
		}},
		expectedError: nil,
	},
}

var testDate *timestamp.Timestamp

func getDate() *timestamp.Timestamp {
	testDate, _ = ptypes.TimestampProto(time.Date(2017, 03, 14, 20, 52, 45, 0, time.UTC))
	return testDate
}

func getDb() *ora.Ses {
	dsn := os.Getenv("GO_OCI8_CONNECT_STRING")
	// db, err := sql.Open("ora", dsn)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return db

	_, _, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		fmt.Println(err)
	}
	// defer env.Close()
	// defer srv.Close()
	// defer ses.Close()

	return ses
}

func getMemcacheClient() *memcache.Client {
	servers := os.Getenv("MEMCACHE_SERVERS")
	mc := memcache.New(servers)

	return mc
}

func TestGetImageDetails(t *testing.T) {
	imageDetailList := &ImageDetailList{}
	for _, c := range cases {
		imageDetailList.DatabaseCaller = c.fakeInfo
		imageDetailList.CacheCaller = c.info
		//imageDetailList.CacheCaller = c.fakeInfo
		imageDetails, err := imageDetailList.GetImageDetails()

		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
		if !reflect.DeepEqual(c.expectedImageDetails, imageDetails) {
			t.Errorf("Expected %q but got %q", c.expectedImageDetails, imageDetails)
		}
	}
}

func TestGetImageDetailsFromDatabase(t *testing.T) {
	for _, c := range imageDetailDatabaseCases {
		imageDetails, err := c.Info.GetImageDetailsFromDatabase()

		if !reflect.DeepEqual(err, c.expectedError) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedError, err)
		}
		if !reflect.DeepEqual(c.expectedImageDetails, imageDetails) {
			t.Errorf("Expected %q but got %q", c.expectedImageDetails, imageDetails)
		}
		//Look into testing numFreeConnections
	}
}

func TestGetImageDetailsFromCache(t *testing.T) {
	for _, c := range imageDetailCacheCases {
		imageDetails, err := c.Info.GetImageDetailsFromCache()

		if !reflect.DeepEqual(err, c.expectedError) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedError, err)
		}
		if !reflect.DeepEqual(c.expectedImageDetails, imageDetails) {
			t.Errorf("Expected %q but got %q", c.expectedImageDetails, imageDetails)
		}
	}
}
