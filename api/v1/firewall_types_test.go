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

package v1

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("C Controller", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	Context("API", func() {
		It("should facilitate API operations", func() {
			key := types.NamespacedName{
				Namespace: "test-ns",
				Name:      "test",
			}

			toCreate := &Firewall{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: key.Namespace,
					Name:      key.Name,
				}}

			By("creating an API object")
			Expect(k8sClient.Create(context.TODO(), toCreate)).To(Succeed())

			By("getting an existing API object back")
			got := &Firewall{} // placeholder
			Expect(k8sClient.Get(context.TODO(), key, got)).To(Succeed())
			Expect(got).To(Equal(toCreate))

			By("deleting an existing API object")
			Expect(k8sClient.Delete(context.TODO(), toCreate)).To(Succeed())
			Expect(k8sClient.Get(context.TODO(), key, toCreate)).ToNot(Succeed())
		})
	})
})
