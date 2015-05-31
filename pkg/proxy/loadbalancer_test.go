package proxy

import (
	"testing"
)

func TestSimple(t *testing.T) {
	endpoints := []string{"1:3000", "1:3001", "1:3002"}
	sb := &serviceBalancer{
		services: make(map[string]*balancerState),
	}
	sb.services["myservice"] = &balancerState{endpoints: endpoints}
	expectEndpoint(t, "myservice", sb, "1:3000")
	expectEndpoint(t, "myservice", sb, "1:3001")
	expectEndpoint(t, "myservice", sb, "1:3002")
}

func expectEndpoint(t *testing.T, svcName string, balancer *serviceBalancer, expected string) {
	endpoint, err := balancer.NextEndpoint(svcName)
	if err != nil {
		t.Fatal(err)
	}
	if expected != endpoint {
		t.Fatalf("expected %s got %s", expected, endpoint)
	}
}
