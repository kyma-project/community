# Add monitoring to a Kyma component

## Enable dashboard in Grafana
- Add label `kyma-grafana: enabled` along with label `app: <value>` to the **PodTemplate**. Make sure label `app: <value>` is added to either `Deployment` or `Statefulset` spec as well. Then a pre-packaged dashboard with [RED](https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture) and [USE](http://www.brendangregg.com/usemethod.html) metrics for the application gets enabled. The dashboard is defined [here](link).
- The dashboard is available as `Services` where the application is visible with the value of it's **app** label in its concerned namespace.


## Enable alerts
- Add label `kyma-alerts: enabled` along with label `app: <value>` to your Kyma component. Make sure label `app: <value>` is added to either `Deployment` or `Statefulset` spec as well. Then a pre-packaged sets of alert rules get enabled. The alert rules are defined [here](link).

## Sample application where monitoring is enabled
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: demo-service
  name: demo-service
  namespace: stage
spec:
  replicas: 1
  selector:
    matchLabels:
      app: demo-service
  template:
    metadata:
      labels:
        app: demo-service
        kyma-alerts: enabled
        kyma-grafana: enabled
    spec:
      containers:
      - image: shazra/demo-service
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 5000
          protocol: TCP
          name: http
        name: demo-service
        resources:
          limits:
            cpu: 100m
            memory: 256Mi
          requests:
            cpu: 10m
            memory: 128Mi
      restartPolicy: Always
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: demo-service
  name: demo-service
  namespace: stage
spec:
  ports:
  - name: http
    port: 5000
    protocol: TCP
    targetPort: 5000
  selector:
    app: demo-service
```
- Dashboard for this application will be available in Grafana `Services` dashboard as `demo-service` in namespace `stage`.