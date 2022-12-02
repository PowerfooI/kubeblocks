/*
Copyright ApeCloud Inc.

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

package playground

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/apecloud/kubeblocks/internal/cli/cloudprovider"
	"github.com/apecloud/kubeblocks/internal/cli/types"
	"github.com/apecloud/kubeblocks/internal/cli/util"
	"github.com/apecloud/kubeblocks/internal/cli/util/fake"
)

var _ = Describe("playground", func() {
	var streams genericclioptions.IOStreams

	// use a fake URL to test
	types.KubeBlocksChartName = fake.KubeBlocksChartName
	types.KubeBlocksChartURL = fake.KubeBlocksChartURL

	BeforeEach(func() {
		streams, _, _, _ = genericclioptions.NewTestIOStreams()
	})

	It("new playground command", func() {
		cmd := NewPlaygroundCmd(streams)
		Expect(cmd).ShouldNot(BeNil())
		Expect(cmd.HasSubCommands()).Should(BeTrue())
	})

	It("init at local host", func() {
		cmd := newInitCmd(streams)
		Expect(cmd != nil).Should(BeTrue())

		o := &initOptions{
			Replicas:      0,
			CloudProvider: defaultCloudProvider,
			IOStreams:     streams,
		}
		Expect(o.validate()).To(MatchError("replicas should greater than 0"))

		o.Replicas = 1
		Expect(o.validate()).Should(Succeed())
		Expect(o.run()).To(MatchError(MatchRegexp("failed to set up k3d cluster")))

		Expect(o.installKubeBlocks()).Should(HaveOccurred())
		Expect(o.installCluster()).Should(HaveOccurred())
	})

	It("init at remote cloud", func() {
		o := &initOptions{
			Replicas:      0,
			IOStreams:     streams,
			CloudProvider: cloudprovider.AWS,
		}
		Expect(o.run()).Should(HaveOccurred())
	})

	It("destroy command", func() {
		cmd := newDestroyCmd(streams)
		Expect(cmd).ShouldNot(BeNil())

		o := &destroyOptions{
			IOStreams: streams,
		}
		Expect(o.destroyPlayground()).Should(HaveOccurred())
		_, _ = removePlaygroundDir()
	})

	It("guide", func() {
		cmd := newGuideCmd()
		Expect(cmd).ShouldNot(BeNil())
		Expect(runGuide()).Should(HaveOccurred())

		getter, err := newObjectsGetter(util.ConfigPath("test"))
		Expect(getter).ShouldNot(BeNil())
		Expect(err).Should(Succeed())

		objs, err := getter.Get()
		Expect(objs).Should(BeNil())
		Expect(err).Should(HaveOccurred())
	})
})