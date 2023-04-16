package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ncostamagna/middleware-lambda-go/auth"
	//"strings"
)

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	fmt.Println("Entra")
	fmt.Println(event)
	token := event.AuthorizationToken
	fmt.Println("-", token, "-")

	id := "6b9b1014-05be-4d51-a695-59e2cc5eb2bd"
	auth, err := auth.New("8FQNq9vHcpCourse")
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, err
	}

	if err := auth.Access(id, token); err != nil {
		fmt.Println(err)
		return generatePolicy("user", "Deny", event.MethodArn), nil
	}

	return generatePolicy("user", "Allow", event.MethodArn), nil
	/*switch strings.ToLower(token) {
	case "allow":
		return generatePolicy("user", "Allow", event.MethodArn), nil
	case "deny":
		return generatePolicy("user", "Deny", event.MethodArn), nil
	case "unauthorized":
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
	default:
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Error: Invalid token")
	}*/
}

func main() {
	lambda.Start(handleRequest)
}
