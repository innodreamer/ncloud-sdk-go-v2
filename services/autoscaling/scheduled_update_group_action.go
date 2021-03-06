/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type ScheduledUpdateGroupAction struct {

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 스케쥴액션명
ScheduledActionName *string `json:"scheduledActionName,omitempty"`

	// 기대능력치
DesiredCapacity *int32 `json:"desiredCapacity,omitempty"`

	// 최소사이즈
MinSize *int32 `json:"minSize,omitempty"`

	// 최대사이즈
MaxSize *int32 `json:"maxSize,omitempty"`

	// 스케줄시작시간
StartTime *string `json:"startTime,omitempty"`

	// 반복스케쥴종료시간
EndTime *string `json:"endTime,omitempty"`

	// 반복스케쥴설정
RecurrenceInKST *string `json:"recurrenceInKST,omitempty"`
}
