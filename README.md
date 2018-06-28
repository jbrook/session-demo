# Simple app for showing k8s horizontal pod auto-scaling with a custom metric

[HPA with custom metrics on GKE](https://cloud.google.com/kubernetes-engine/docs/tutorials/custom-metrics-autoscaling)

**Note:**

The application is currently instrumented to serve Prometheus metrics - active sessions is represented as a
'guage'. A ' [prometheus-to-sd](https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/prometheus-to-sd)' sidecar is then run next to the application container to collect the metric and forward to Stackdriver.

Another and new approach would be to use [OpenCensus](https://opencensus.io/) and [export directly to Stackdriver](https://medium.com/@DazWilkin/return-to-opencensus-42623f1b55b8).