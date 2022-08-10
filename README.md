
### 构建镜像

```bash
docker build -t gin-bot .

```


### docker 部署
```bash
docker run -ti --rm -p 8080:8080 -v `pwd`/conf:/go/conf webhook-bot
```

```json
curl -X POST http://127.0.0.1:8080/api/v1/webhook -d '
{
  "receiver": "prometheus", 
  "status": "", 
  "alerts": [
    {
      "status": "firing", 
      "labels": {
        "alertname": "test", 
        "alerttype": "metric", 
        "host_ip": "10.48.175.46", 
        "node": "ip-10-48-175-46.ap-east-1.compute.internal", 
        "rule_id": "c67103f7a217b73b", 
        "severity": "warning", 
        "cluster": "default"
      }, 
      "annotations": {
        "rules": "[{\"_metricType\":\"node:node_memory_utilisation:{$1}\",\"condition_type\":\">\",\"thresholds\":\"1\",\"unit\":\"%\"}]",
        "summary": "节点 ip-10-48-175-46.ap-east-1.compute.internal 内存用量 > 1%", 
        "kind": "Node", 
        "message": "", 
        "resources": "[\"ip-10-48-175-46.ap-east-1.compute.internal\"]", 
        "rule_update_time": "2022-08-04T10:11:43Z"
      }, 
      "startsAt": "2022-08-04T10:13:05.289543453Z", 
      "endsAt": "0001-01-01T00:00:00Z", 
      "generatorURL": "/graph?g0.expr=node%3Anode_memory_utilisation%3A%7Bnode%3D%22ip-10-48-175-46.ap-east-1.compute.internal%22%7D+%3E+0.01&g0.tab=1", 
      "fingerprint": "fa3d526de75e6a99"
    }
  ], 
  "groupLabels": {
    "rule_id": "c67103f7a217b73b", 
    "alertname": "test"
  }, 
  "commonLabels": {
    "node": "ip-10-48-175-46.ap-east-1.compute.internal", 
    "rule_id": "c67103f7a217b73b", 
    "severity": "warning", 
    "alertname": "test", 
    "alerttype": "metric", 
    "host_ip": "10.48.175.46"
  }, 
  "commonAnnotations": "", 
  "externalURL": "http://alertmanager-main-0:9093"
}'
```


### k8s 部署

```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: webhook-bot-v1
  namespace: kubesphere-monitoring-system
  labels:
    app: webhook-bot
    version: v1
  annotations:
    deployment.kubernetes.io/revision: '3'
    kubesphere.io/creator: admin
spec:
  replicas: 1
  selector:
    matchLabels:
      app: webhook-bot
      version: v1
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: webhook-bot
        version: v1
      annotations:
        logging.kubesphere.io/logsidecar-config: '{}'
    spec:
      volumes:
        - name: volume-wnmoj2
          configMap:
            name: webhook-bot-template
            items:
              - key: dingtalk.tmpl
                path: dingtalk.tmpl
              - key: wxchat.tmpl
                path: wxchat.tmpl
            defaultMode: 420
        - name: volume-nruk29
          secret:
            secretName: webhook-bot
            items:
              - key: conf.yaml
                path: conf.yaml
            defaultMode: 420
      containers:
        - name: webhook-bot
          image: 'webhook-bot:v1.0.1'
          ports:
            - name: http-0
              containerPort: 8080
              protocol: TCP
          resources:
            limits:
              cpu: 200m
              memory: 200Mi
            requests:
              cpu: 100m
              memory: 10Mi
          volumeMounts:
            - name: volume-wnmoj2
              readOnly: true
              mountPath: /go/templates
            - name: volume-nruk29
              readOnly: true
              mountPath: /go/conf
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: IfNotPresent
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      nodeSelector:
        nodegroup-role: obs
        nodegroup-zone: dmz
      serviceAccountName: default
      serviceAccount: default
      securityContext: {}
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 25%
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
```