# Objective

The telemetry operator comes with a pipeline configured to push logs to loki. A user of Kyma can define his own pipeline and push the logs to his own logging backend.

The task here is to understand in a scenario when eg. the logging backend defined by  user is not working anymore what side effects it can have.

## Setup

For the setup the configuration consisted of used `filesystem` buffer. So when one of the logpipeline output fail then the logs could be buffered on filesystem. The `in-memory` buffer has a disadvantage that if the `fluent-bit` pod is restarted due to any reason then it could lead to loss of logs.  The following [knowledge article](https://docs.fluentbit.io/manual/administration/buffering-and-storage) was refferd during the tests.

Setup consisted of following items
1. Installed kyma with telemetry operator
2. Two outputs were deployed in the kyma cluster
    - Loki which comes with kyma
    - Mockserver [deployment](./assets/logpipeline-invstigation/mock-server.yaml)
3. Log generator [daemon set](./assets/logpipeline-invstigation/log-generator.yaml) to generate huge amount of logs to saturate the buffer faster.
4. [Function](./assets/logpipeline-invstigation/func.js) to check if the logs are being delivered when one of the output is down.
5. To simulate failures the port of the service was changed so that dns resolution would still work but logs wont be deliverd.

## Test Cases

### Case 1
![a](./assets/logpipeline-invstigation/case-1/case-1.svg)

Setup:
1. one input and 2 outputs (with limiting max number of chunks in filesystem) without rewrite tags.
2. logpipelines: [loki](./assets/logpipeline-invstigation/case-1/loki.yaml), [mock-server](./assets/logpipeline-invstigation/case-1/mockserver.yml)

Result:
1. When one of the output is down then we see the filesystem buffer is being filled up and would only keep 150M of latest data (old data would purged). Eventually the output plugin for loki stops as well as the buffer at tail plugin is full.
2. When both down then tail plugin stopped is stopped after filling up the buffer. Although the buffer was not fully filled (104m/150M). Still the tail plugin was stopped

### Case 2
![a](./assets/logpipeline-invstigation/case-2/case-2.svg)

Setup:
1. one input and 2 outputs (loki (fluentbit-loki plugin) + mockserver) with rewrite-tags
2. logpipelines: [loki](./assets/logpipeline-invstigation/case-2/loki.yaml), [mock-server](./assets/logpipeline-invstigation/case-2/mockserver.yml)


Result:
1. Output chunking keeps with 150M of newest data.
2.  The tail plugin sends the data to next phase of pipeline (filter plugin). If the outputs is not working then only the buffer at the rewrite tag would be filled and tail plugin buffer wont be filled
3. When the buffer is full then we see some error logs stating that and the old logs are discarded
    ```unix
    [2022/04/21 14:38:23] [error] [input:emitter:log_emitter] error registering chunk with tag: log_rewritten
    [2022/04/21 14:38:24] [error] [input chunk] chunk 1-1650551903.999375896.flb would exceed total limit size in plugin http.1
    ```
4. When grafana-loki plugin is stopped it stopped all the pipelines (so mockserver which was fine was also stopped) 
5. When mock server was down the logs were still shipped to loki


### Case 3

Setup:
1. one input with 2 outputs (mockserver + mockserver) with rewrite tags
2. logpipelines: [mockserver-1](./assets/logpipeline-invstigation/case-3/mockserver-1.yml), [mockserver-2](./assets/logpipeline-invstigation/case-3/mockserver-2.yml)

![a](./assets/logpipeline-invstigation/case-3/case-3.svg)
Result
1. When both outputs are down then the buffer was filled and eventually the fluentbit pod got killed because of 500 (most probably because of CPU throttling)
2. Tail plugin kept pushing logs to rewrite buffer and they were eventually lost

### Case 4
Setup:
1. one input with 2 outputs (loki (with official loki plugin) + mockserver) with rewrite tags
2. logpipelines: [loki](./assets/logpipeline-invstigation/case-4/loki.yml), [mockserver](./assets/logpipeline-invstigation/case-4/mock-server.yml)

![a](./assets/logpipeline-invstigation/case-4/case-4.svg)

Result
1. Mockserver was down and the loki output was still working
2. The chunks in the filesystem buffer are rolled  (the old chunks deleted new ones created)
3. When loki output was down the mockserver output was still working

![mockserver-down](/assets/logpipeline-invstigation/case-4/dashboard-mock-down.png)


![loki-down](/assets/logpipeline-invstigation/case-4/dashboard-loki-down.png)

### Case 5
Setup:
1. one input with 2 outputs (loki (with official loki plugin) + mockserver) without rewrite tags
2. logpipelines: [loki](./assets/logpipeline-invstigation/case-5/loki.yml), [mockserver](./assets/logpipeline-invstigation/case-5/mock-server.yml)

![a](./assets/logpipeline-invstigation/case-5/case-5.svg)
Result
1. Mockserver was down, this led to filling of the tail buffer
2. When the tail plugin buffer is filled, it kept still reading.
3. The chunks in tail plugin is rolled (the old chunks deleted new ones created)
4. Loki output stopped working as well
5. Found following [issue](https://github.com/fluent/fluent-bit/issues/4373) in github which describes the same problem.

![mockserver-down](/assets/logpipeline-invstigation/case-5/dashboard-mock-down.png)

### Case 6
Setup
1. 2 Inputs and 2 outputs (loki (fluentbit-loki plugin) + mockserver) and no rewrite tags
![a](./assets/logpipeline-invstigation/case-6/case-6.svg)
Result
1. Each pipeline has its own buffer
2. When both the outputs are stopped both the tail plugins are loosing logs. It would keep the latest logs (although the amount of logs are differemt)


## Summary
We performed various tests and found that a buffer with filesystem (through rewrite_tag filter) is necessary to prevent loss of logs when one of the output is down. However loss of logs cannot be prevented completely, if the buffer is filled up then fluentbit keeps only the latest logs.



