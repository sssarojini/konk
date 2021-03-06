// Code generated by apiregister-gen. DO NOT EDIT.

package apis

import (
	"github.com/infobloxopen/konk/test/apiserver/pkg/apis/contact"
	_ "github.com/infobloxopen/konk/test/apiserver/pkg/apis/contact/install" // Install the contact group
	contactv1alpha1 "github.com/infobloxopen/konk/test/apiserver/pkg/apis/contact/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		contactv1alpha1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetContactAPIBuilder(),
	}
}

var contactApiGroup = builders.NewApiGroupBuilder(
	"contact.example.infoblox.com",
	"github.com/infobloxopen/konk/test/apiserver/pkg/apis/contact").
	WithUnVersionedApi(contact.ApiVersion).
	WithVersionedApis(
		contactv1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetContactAPIBuilder() *builders.APIGroupBuilder {
	return contactApiGroup
}
