/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-06-21T02:22:22Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type DeleteScheduledActionRequest struct {

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 스케쥴액션명
ScheduledActionName *string `json:"scheduledActionName"`
}