/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type DeletePolicyRequest struct {

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName"`

	// 정책명
PolicyName *string `json:"policyName"`
}
