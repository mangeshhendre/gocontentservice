package list

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/divyag9/gocontentservice/pkg/contentservice"
	"gopkg.in/rana/ora.v4"
)

//GetImageDetails retrieves image details for an ordernumber. It will first see if the details are in the cache if not gets the details from database
func (i *ImageDetailList) GetImageDetails() ([]*contentservice.ImageDetail, error) {
	// Retrive the imagedetails from cache if it exists
	imageDetails, err := i.CacheCaller.GetImageDetailsFromCache()
	if err != nil {
		//Making a call to the database cause the results are not in cache
		if reflect.DeepEqual(err, fmt.Errorf("memcache: cache miss")) {
			imageDetails, err = i.DatabaseCaller.GetImageDetailsFromDatabase()
			if err != nil {
				return nil, err
			}

			err = i.CacheCaller.SetImageDetailsToCache(imageDetails)
			if err != nil {
				return nil, err
			}
			return imageDetails, nil
		}
		return nil, err
	}

	return imageDetails, nil
}

// GetImageDetailsFromCache retrieves the imagedetails for an ordernumber from the cache
func (o *Info) GetImageDetailsFromCache() ([]*contentservice.ImageDetail, error) {
	item, err := o.MemClient.Get(fmt.Sprintf("OrderNumber:%q", o.OrderNumber))
	if err != nil {
		return nil, err
	}

	imageDetailBytes := item.Value
	imageDetails := []*contentservice.ImageDetail{}
	decBuf := bytes.NewBuffer(imageDetailBytes)
	err = gob.NewDecoder(decBuf).Decode(&imageDetails)
	if err != nil {
		return nil, err
	}
	return imageDetails, nil
}

// GetImageDetailsFromDatabase retrieves the imagedetails for an ordernumber from the database
func (o *Info) GetImageDetailsFromDatabase() ([]*contentservice.ImageDetail, error) {
	prepStatement, err := o.Db.Prep("CALL CONTENTSERVICE.RETRIEVEDEPARTMENTS(:1)")
	if err != nil {
		return nil, err
	}
	defer prepStatement.Close()

	resultSet := &ora.Rset{}
	imageDetails := []*contentservice.ImageDetail{}
	//rownum, err := o.Db.PrepAndExe("CALL CONTENTSERVICE.RETRIEVEDEPARTMENTS(:1)", resultSet)
	rownum, err := prepStatement.Exe(resultSet)
	//rows, err := o.Db.Query("SELECT * FROM TABLE(CONTENTSERVICE.RETRIEVEIMAGEDETAILLIST(:1))", o.OrderNumber)
	if err != nil {
		return nil, err
	}
	fmt.Println(rownum)

	// for resultSet.Next() {
	// 	var imageDetail *contentservice.ImageDetail
	// 	// err = resultSet.Scan(imageDetail.ImageId,
	// 	// 	imageDetail.ImageFileName,
	// 	// 	imageDetail.ScanDate,
	// 	// 	imageDetail.ImageUTCDate,
	// 	// 	imageDetail.ImageTakenDate,
	// 	// 	imageDetail.DateCreated,
	// 	// 	imageDetail.OrderNumber,
	// 	// 	imageDetail.Archived,
	// 	// 	imageDetail.Category,
	// 	// 	imageDetail.ContractorId,
	// 	// 	imageDetail.DateModified,
	// 	// 	imageDetail.DeptCode,
	// 	// 	imageDetail.DescPrefix,
	// 	// 	imageDetail.DescText,
	// 	// 	imageDetail.FileSize,
	// 	// 	imageDetail.ImageHeight,
	// 	// 	imageDetail.ImageWidth,
	// 	// 	imageDetail.ImageType,
	// 	// 	imageDetail.ImageRotated,
	// 	// 	imageDetail.ReleaseDate,
	// 	// 	imageDetail.ThumbnailSize,
	// 	// 	imageDetail.MimeType,
	// 	// 	imageDetail.WebFileName,
	// 	// 	imageDetail.GeneratedImageFilePath,
	// 	// 	imageDetail.Guid)
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }
	// 	imageDetails = append(imageDetails, imageDetail)
	// }
	// err = resultSet.Err()
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println(resultSet.Row)

	for _, v := range resultSet.Columns {
		fmt.Println(v.Name)
	}
	o.Db.Close()
	return imageDetails, nil
}

//SetImageDetailsToCache adds the imagedetails to cahce
func (o *Info) SetImageDetailsToCache(imageDetails []*contentservice.ImageDetail) error {
	key := fmt.Sprintf("OrderNumber:%q", o.OrderNumber)
	encBuf := new(bytes.Buffer)
	err := gob.NewEncoder(encBuf).Encode(imageDetails)
	if err != nil {
		return err
	}
	err = o.MemClient.Set(&memcache.Item{Key: key, Value: encBuf.Bytes(), Expiration: o.SecondsToExpiry})
	if err != nil {
		return err
	}

	return nil
}
