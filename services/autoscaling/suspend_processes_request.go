/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type SuspendProcessesRequest struct {

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName"`

	// 분류프로세스코드리스트
ScalingProcessCodeList []*string `json:"scalingProcessCodeList,omitempty"`
}
