module github.com/evryfs/github-actions-runner-operator

go 1.15

require (
	github.com/go-logr/logr v1.2.0
	github.com/google/go-github/v33 v33.0.0
	github.com/gophercloud/gophercloud v0.15.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79
	github.com/imdario/mergo v0.3.12
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/palantir/go-githubapp v0.6.0
	github.com/redhat-cop/operator-utils v1.3.4
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.7.0
	k8s.io/api v0.24.0
	k8s.io/apimachinery v0.24.0
	k8s.io/client-go v0.24.0
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
	sigs.k8s.io/controller-runtime v0.12.1
)
