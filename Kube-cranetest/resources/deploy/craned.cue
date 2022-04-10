stateless: {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "apps/v1"
		kind:       "Deployment"
	}
	description: ""
	labels: {}
	type: "component"
}

template: {
	output: {
		spec: {
			selector: matchLabels: "app.oam.dev/component": "crane"
			template: {
				metadata: labels: "app.oam.dev/component": "crane"
				spec: containers: [{
					name:  "crane"
					image: "docker.io/gocrane/craned:v0.2.0"
				}]
			}
		}
		apiVersion: "apps/v1"
		kind:       "Deployment"
	}
	outputs: {}
	parameter: {}
}
