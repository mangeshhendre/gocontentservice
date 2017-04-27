package list

import (
	"fmt"

	"github.com/divyag9/gocontentservice/pkg/contentservice"
	"gopkg.in/rana/ora.v4"
)

// OrderNumberInfo holds all the information required for listing the imagedetails for an ordernumber
type OrderNumberInfo struct {
	OrderNumber int64
	Db          *ora.Ses
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
}

//GetImageDetails retrieves image details for an ordernumber. It will first see if the details are in the cache if not gets the details from database
func (i *ImageDetailList) GetImageDetails() ([]*contentservice.ImageDetail, error) {
	// Retrive the imagedetails from cache if it exists
	imageDetails, err := i.CacheCaller.GetImageDetailsFromCache()
	if err != nil {
		return nil, err
	}
	//Making a call to the database cause the results are not in cache
	if len(imageDetails) == 0 {
		imageDetails, err = i.DatabaseCaller.GetImageDetailsFromDatabase()
		if err != nil {
			return nil, err
		}
	}
	return imageDetails, nil
}

// GetImageDetailsFromCache retrieves the imagedetails for an ordernumber from the cache
func (o *OrderNumberInfo) GetImageDetailsFromCache() ([]*contentservice.ImageDetail, error) {
	return []*contentservice.ImageDetail{}, nil
}

// GetImageDetailsFromDatabase retrieves the imagedetails for an ordernumber from the database
func (o *OrderNumberInfo) GetImageDetailsFromDatabase() ([]*contentservice.ImageDetail, error) {
	// prepStatement, err := o.Db.Prep("CALL CONTENTSERVICE.RETRIEVEDEPARTMENTS(:1)")
	// if err != nil {
	// 	return nil, err
	// }
	// defer prepStatement.Close()

	//var resultSet *sql.Rows
	resultSet := &ora.Rset{}
	//orderNumber := o.OrderNumber
	imageDetails := []*contentservice.ImageDetail{}
	//rownum, err := o.Db.PrepAndExe("CALL CONTENTSERVICE.RETRIEVEDEPARTMENTS(:1)", resultSet)
	//rownum, err := prepStatement.Exe(resultSet)
	//rows, err := o.Db.Query("SELECT * FROM TABLE(CONTENTSERVICE.RETRIEVEIMAGEDETAILLIST(:1))", o.OrderNumber)
	//fmt.Println(rownum)
	//var str string
	stmt, err := o.Db.Prep("CALL INTG_PKG.PROC1(:1)")
	stmt.Exe(resultSet)
	if err != nil {
		fmt.Println("in prep error")
		return nil, err
	}
	fmt.Println("rset", resultSet.Row)

	//defer rows.Close()

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
	count := len(resultSet.Columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	values = resultSet.Row
	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	for i := range resultSet.Columns {
		var v interface{}
		val := values[i]
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		fmt.Println("v", v)
	}

	// for _, v := range resultSet.Columns {
	// 	fmt.Println(v.Name)
	// }
	o.Db.Close()
	return imageDetails, nil
}
