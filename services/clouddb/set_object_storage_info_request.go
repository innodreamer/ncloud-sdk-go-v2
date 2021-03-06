/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type SetObjectStorageInfoRequest struct {

	// 오브젝트스토리지 AccessKey
ObjectStorageAccessKey *string `json:"objectStorageAccessKey"`

	// 오브젝트스토리지 SecretKey
ObjectStorageSecretKey *string `json:"objectStorageSecretKey"`

	// 엔드포인트
Endpoint *string `json:"endpoint"`

	// 버킷이름
BucketName *string `json:"bucketName"`
}
