/*
Copyright 2020 Devtron Labs Pvt Ltd.

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

package language

import "testing"

func Test_updateKubernetesObjectsJson(t *testing.T) {
	to := `
{
  "apiVersion": "v1",
  "kind": "ConfigMap",
  "metadata": {
    "name": "test-cm",
    "namespace": "dev",
    "labels": {
      "app.kubernetes.io/instance": "my-app"
    }
  },
  "data": {
    "name": "abc"
  }
}
`
	from := `
{
  "apiVersion": "v1",
  "kind": "ConfigMap",
  "metadata": {
    "name": "test-cm"
  },
  "update": {
    "data": {
      "name": "def"
    }
  }
}
`
	want := `
{
  "apiVersion": "v1",
  "kind": "ConfigMap",
  "metadata": {
    "name": "test-cm",
    "namespace": "dev",
    "labels": {
      "app.kubernetes.io/instance": "my-app"
    }
  },
  "data": {
    "name": "def"
  }
}
`
	patterns := []string{"apiVersion", "kind", "metadata.name"}
	type args struct {
		from     string
		to       string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "json merge",
			args: args{
				from:     from,
				to:       to,
				patterns: patterns,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateKubernetesObjectsJson(tt.args.from, tt.args.to, tt.args.patterns); got != tt.want {
				t.Errorf("updateKubernetesObjectsJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateKubernetesObjectsYaml(t *testing.T) {
	to := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
data:
  name: abc
`
	from := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
  namespace: dev
  labels:
    app.kubernetes.io/instance: my-app
update:
  data:
    name: def
`
	patterns := []string{"apiVersion", "kind", "metadata.name"}
	type args struct {
		from     string
		to       string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "yaml merge",
			args: args{
				from:     from,
				to:       to,
				patterns: patterns,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateKubernetesObjectsYaml(tt.args.from, tt.args.to, tt.args.patterns); got != tt.want {
				t.Errorf("updateKubernetesObjectsYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateKubernetesObjectsYaml1(t *testing.T) {
	to := `
# Source: guard/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: guard
  labels:
    app: guard
    chart: guard-3.9.1
    release: devtron
    releaseVersion: "1"
    pipelineName: guard-dt-prod
spec:
  selector:
    matchLabels:
      app: guard
      release: devtron
  replicas: 2
  minReadySeconds: 60
  template:
    metadata:
      labels:
        app: guard
        appId: "181"
        envId: "4"
        release: devtron
    spec:
      terminationGracePeriodSeconds: 30
      restartPolicy: Always
      containers:
        - name: guard
          image: "686244538589.dkr.ecr.us-east-2.amazonaws.com/release/guard:4bea84da-113-1767"
          imagePullPolicy: IfNotPresent
          ports:
            - name: app
              containerPort: 8080
              protocol: TCP
          args:
            - -alsologtostderr
            - --log_dir=/
            - -v=10
            - --validator_url=http://devtron-service.devtroncd:80
            - 2>&1
          env:
            - name: CONFIG_HASH
              value: 01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b
            - name: SECRET_HASH
              value: abaee17d930a742e0a3554336348fde8a2b20e23bbdabb29b4acb8ac393b5da9
            - name: DEVTRON_APP_NAME
              value: guard
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
          resources:
            limits:
              cpu: "0.1"
              memory: 100Mi
            requests:
              cpu: "0.1"
              memory: 100Mi

          volumeMounts:
            - mountPath: /tmp
              name: log-volume
            - name: guard-secret-vol
              mountPath: /etc/certs
      volumes:
        - emptyDir: {}
          name: log-volume
        - name: guard-secret-vol
          secret:
            secretName: guard-secret
  revisionHistoryLimit: 3
`
	from := `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: guard
update:
  spec:
    template:
      spec:
        containers:
        - resources:
            limits: 
              cpu: "0.15"
              memory: 110Mi
            requests:
              cpu: "0.1"
              memory: 100Mi
`
	patterns := []string{"apiVersion", "kind", "metadata.name"}
	type args struct {
		from     string
		to       string
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "array merge test",
			args: args{
				from: from,
				to: to,
				patterns: patterns,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateKubernetesObjectsYaml(tt.args.from, tt.args.to, tt.args.patterns); got != tt.want {
				t.Errorf("updateKubernetesObjectsYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}