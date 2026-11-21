package tests

import (
	"testing"

	"github.com/fredcamps/chamego/pkg/core"
)

func TestNew(t *testing.T) {
	p := core.New("chamego", "1.0.0")
	if p.Name != "chamego" {
		t.Errorf("Expected Name to be 'chamego', got %s", p.Name)
	}
	if p.Version != "1.0.0" {
		t.Errorf("Expected Version to be '1.0.0', got %s", p.Version)
	}
}

func TestInfo(t *testing.T) {
	p := core.New("chamego", "1.0.0")
	info := p.Info()
	expected := "chamego 1.0.0 - chamego"
	if info != expected {
		t.Errorf("Expected Info to be '%s', got %s", expected, info)
	}
}
