package templates

/*
Not including Vsphere(cloudProvider) and Authz templates
Will they change and require Rancher to pass them to RKE
depending on Kubernetes version?
*/

const (
	Calico        = "calico"
	Canal         = "canal"
	Flannel       = "flannel"
	Weave         = "weave"
	CoreDNS       = "coreDNS"
	KubeDNS       = "kubeDNS"
	MetricsServer = "metricsServer"
	NginxIngress  = "nginxIngress"
	TemplateKeys  = "templateKeys"

	calicov18  = "calico-v1.8"
	calicov113 = "calico-v1.13"
	calicov115 = "calico-v1.15"
	calicov116 = "calico-v1.16"

	canalv115 = "canal-v1.15"
	canalv116 = "canal-v1.16"
	canalv113 = "canal-v1.13"
	canalv18  = "canal-v1.8"

	flannelv116 = "flannel-v1.16"
	flannelv115 = "flannel-v1.15"
	flannelv18  = "flannel-v1.8"

	coreDnsv18 = "coredns-v1.8"
	kubeDnsv18 = "kubedns-v1.8"

	metricsServerv18 = "metricsserver-v1.8"

	weavev18        = "weave-v1.8"
	nginxIngressv18 = "nginxingress-v1.8"
)

func LoadK8sVersionedTemplates() map[string]map[string]string {
	return map[string]map[string]string{
		Calico: {
			">=1.16.0-alpha":         calicov116,
			">=1.15.0 <1.16.0-alpha": calicov115,
			">=1.13.0 <1.15.0": calicov113,
			">=1.8.0 <1.13.0":  calicov18,
		},
		Canal: {
			">=1.16.0-alpha":         canalv116,
			">=1.15.0 <1.16.0-alpha": canalv115,
			">=1.13.0 <1.15.0": canalv113,
			">=1.8.0 <1.13.0":  canalv18,
		},
		Flannel: {
			">=1.16.0-alpha":         flannelv116,
			">=1.15.0 <1.16.0-alpha": flannelv115,
			">=1.8.0 <1.15.0":  flannelv18,
		},
		CoreDNS: {
			">=1.8.0 <1.16.0": coreDnsv18,
		},
		KubeDNS: {
			">=1.8.0 <1.16.0": kubeDnsv18,
		},
		MetricsServer: {
			">=1.8.0 <1.16.0": metricsServerv18,
		},
		Weave: {
			">=1.8.0 <1.16.0": weavev18,
		},
		NginxIngress: {
			">=1.8.0 <1.16.0": nginxIngressv18,
		},
		TemplateKeys: getTemplates(),
	}
}

func getTemplates() map[string]string {
	return map[string]string{
		calicov113: CalicoTemplateV113,
		calicov115: CalicoTemplateV115,
		calicov116: CalicoTemplateV116,
		calicov18:  CalicoTemplateV112,

		flannelv115: FlannelTemplateV115,
		flannelv116: FlannelTemplateV116,
		flannelv18:  FlannelTemplate,

		canalv113: CanalTemplateV113,
		canalv18:  CanalTemplateV112,
		canalv115: CanalTemplateV115,
		canalv116: CanalTemplateV116,

		coreDnsv18: CoreDNSTemplate,
		kubeDnsv18: KubeDNSTemplate,

		metricsServerv18: MetricsServerTemplate,

		weavev18: WeaveTemplate,

		nginxIngressv18: NginxIngressTemplate,
	}
}
