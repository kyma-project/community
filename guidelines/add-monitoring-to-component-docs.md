# Add monitoring to a Kyma component

## Enable dashboard in Grafana
- Add the `kyma-grafana: enabled` and the `app: <value>` label to the **PodTemplate**. Make sure you add the `app: <value>` label either to `Deployment` or `Statefulset` specification as well. Performing this step enables the pre-packaged dashboard with [RED](https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture) and [USE](http://www.brendangregg.com/usemethod.html) metrics for the application. The dashboard is defined [here](https://github.com/kyma-project/kyma/blob/master/resources/monitoring/charts/grafana/dashboards/rest-service.json).
- To see the dashboard, go to **General > Services** in Grafana. There, you can find your application using the values for the `app` label and the Namespace you specified in your configuration.

## Enable alerts
- Add the `kyma-alerts: enabled` and the `app: <value>` label to your Kyma component. Make sure you add the `app: <value>` label either to `Deployment` or `Statefulset` specification as well. Performing this step enables pre-packaged sets of alert rules. The alert rules are defined [here](https://github.com/kyma-project/kyma/blob/master/resources/monitoring/charts/alert-rules/templates/alert-rules-rest-services.yaml).

## Sample application with enabled monitoring
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
- The dashboard for this application is available in Grafana under **General > Services**. To display it, select `demo-service` in the `stage` Namespace.
