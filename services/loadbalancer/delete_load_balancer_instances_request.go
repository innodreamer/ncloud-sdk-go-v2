/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type DeleteLoadBalancerInstancesRequest struct {

	// 로드밸런서인스턴스번호리스트
LoadBalancerInstanceNoList []*string `json:"loadBalancerInstanceNoList"`
}
