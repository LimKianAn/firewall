/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	firewallv1 "github.com/LimKianAn/firewall/api/v1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	// defer close(done)

	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

	cfg, err := testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	Expect(scheme.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(firewallv1.AddToScheme(scheme.Scheme)).To(Succeed())

	// +kubebuilder:scaffold:scheme

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	// Expect(
	// 	(&FirewallReconciler{
	// 		Client:       k8sManager.GetClient(),
	// 		Log:          ctrl.Log.WithName("controllers").WithName("Firewall"),
	// 		Scheme:       k8sManager.GetScheme(),
	// 		ServiceIP:    "172.17.0.1", // FIXME: Avoid hard-coded defaults
	// 		PrivateVrfID: 0,            // FIXME: Avoid hard-coded defaults
	// 		EnableIDS:    false,        // FIXME: Avoid hard-coded defaults
	// 	}).SetupWithManager(k8sManager)).To(Succeed())

	Expect(
		(&ClusterwideNetworkPolicyReconciler{
			Client: k8sManager.GetClient(),
			Log:    ctrl.Log.WithName("controllers").WithName("ClusterwideNetworkPolicy"),
			// Scheme:   k8sManager.GetScheme(),
			Recorder: k8sManager.GetEventRecorderFor("clusterwidenetworkpolicy-controller"),
		}).SetupWithManager(k8sManager)).To(Succeed())

	go func() {
		Expect(k8sManager.Start(ctrl.SetupSignalHandler())).To(Succeed())
	}()

	k8sClient = k8sManager.GetClient()
	Expect(k8sClient).ToNot(BeNil())

	close(done)
}, 30)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	// gexec.KillAndWait(5 * time.Second)
	Expect(testEnv.Stop()).To(Succeed())
})

// func ExpectSuccess(err error) {
// 	Expect(err).ToNot(HaveOccurred())
// }
