# Eventing team code guidelines

**Document Intention** 

The intention of this document is to provide style, testing, and code guidelines for components of the **eventing** team. Other teams in Kyma may follow different guidelines.
The document is supposed to be a **living** document that can be changed by anyone in the team. However, the team majority (> 50%) must agree on changes for this guideline.

This guide shall be publicly available so that it can be referenced when reviewing pull requests. The agreements of this guide shall not affect **external collaborators**. If an external PR does not align with our guidelines, we should accept the PR as it is and apply the guidelines ourselves in a follow-up PR so we don't block external contributors. Of course, also external PRs shall meet our quality standards but we shouldn't be too nitpicky.

For now, the guidelines shall be applied to new code (PRs). There is no need to rewrite our entire code base now. This will gradually happen over time. As soon as the code reflects our guidelines, the need to look into this guide will decrease more and more.

## Table of contents

<!-- generate me with markdown-toc 
```bash
// source: https://github.com/jonschlinkert/markdown-toc
markdown-toc -i --maxdepth 2 eventing-code-guidelines.md
Do NOT TOUCH anything between the toc comments because this is used as a `marker` where to place the toc for markdown-toc.
```
-->

<!-- toc -->

<!-- tocstop -->

## Recommended libraries

The following section describes the desired choice of libraries for testing, logging, etc.

### Testing

**Testing frameworks**:

Use [t.Testing](https://pkg.go.dev/testing) for unit tests and for controller integration tests.

**Assertion libraries**:

Use [stretchr/testify](https://github.com/stretchr/testify) as assertion library. The controller integration tests are an exception to this. Use [onsi/gomega](https://github.com/onsi/gomega) as an assertion library for them instead.

**Mocking libraries**:

Use [stretchr/testify/mock](https://github.com/stretchr/testify#mock-package) in combination with [vektra/mockery](https://github.com/vektra/mockery) for generating mocks, or create your own mock by implementing the corresponding interface.

### Structured Logging

Both [logrus](https://github.com/sirupsen/logrus) and [zap](https://github.com/uber-go/zap) are widely used. logrus has the **advantage** and **disadvantage** at the same time of being a **drop-in** replacement for the **stdlib**.
logrus supports structured logging but does not enforce using it. 
Citing from the [logrus github page](https://github.com/sirupsen/logrus), it is clear that zap is the modern alternative to logrus:
> Logrus is in maintenance-mode. We will not be introducing new features. It’s simply too hard to do in a way that won’t break many people’s projects, which is the last thing you want from your Logging library (again...).

> Many fantastic alternatives have sprung up. Logrus would look like those, had it been re-designed with what we know about structured logging in Go today. Check out, for example, Zerolog, Zap, and Apex.

For these reasons, `uber-go/zap` is the **preferred** structured logging library.

Furthermore, consider using `github.com/kyma-project/kyma/common/logging/logger`, which provides further **abstraction** over `uber-go/zap`.

#### See also

- [Kyma Logging Proposal](https://github.com/kyma-project/community/blob/main/concepts/observability-consistent-logging/improvement-of-log-messages-usability.md)

## Documentation guidelines

Code is read many times but sometimes only written once. Therefore you should always make sure that you follow these guidelines.

**Guidelines**:
- Add **documentation** to **all exported** functions, variables, types etc. Comments on private elements are also welcome because they help the reader.
- Assume that there is a generated html **version** of our **docs** on <https://pkg.go.dev>, for example [here](https://pkg.go.dev/github.com/kyma-project/kyma/components/eventing-controller/api/v1alpha1) for the eventing-controller.
- Comments on exported functions, variables, or types shall **start** with the name of the element and **end** with a dot (".") Write real sentences.

### Good practice

The following example from [go.dev](https://go.dev/blog/godoc) shows how you should write documentation:

```go
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

If you are in doubt how to write the godoc properly, run `godoc` and check the generated documentation:
```bash
$ go install golang.org/x/tools/cmd/godoc@latest
$ godoc -http=:8080
```

### Bad practice

**Example 1**: Using backticks when referring to arguments in the method signature

<details>
	<summary>Don't</summary>

In the example, **w** is the first argument to the function `Fprint`.
It is not necessary to put w in backticks. Goland supports jumping to the element definition (at least sometimes :-D). Using backticks however breaks the feature.

```go
// Fprint formats using the default formats for its operands and writes to `w`.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```

</details>


### See also

- [Tutorial on writing go doc using godoc itself](https://pkg.go.dev/github.com/natefinch/godocgo#hdr-Formatting)

- [godoc on go bloc](https://go.dev/blog/godoc)

- [Go doc finder - Kyma eventing-controller example](https://pkg.go.dev/github.com/kyma-project/kyma/components/eventing-controller/api/v1alpha1)

## Dockerfile guidelines 

- no shell in container
- do not run as root user if not necessary
- use [multi-stage](https://docs.docker.com/develop/develop-images/multistage-build/) build pattern

## Coding guidelines 

### Single-line arguments vs multiline arguments

<details>
	<summary>Don't: method call with arguments on same line</summary>

```go
// source: https://github.com/kyma-project/kyma/pull/13242/files#diff-2f168bb71c1ca8d2f5781bac737393acd9bc3a1f829bdbe91093b7db54972433L130
subWithGivenWebhookAuth.Spec.ProtocolSettings = reconcilertesting.NewProtocolSettings(eventingtesting.WithBinaryContentMode, eventingtesting.WithExemptHandshake, eventingtesting.WithAtLeastOnceQOS, eventingtesting.WithDefaultWebhookAuth)
```
</details>

<details>
	<summary>Do: method call arguments on multiple lines</summary>

```go
subWithGivenWebhookAuth.Spec.ProtocolSettings = reconcilertesting.NewProtocolSettings(
	eventingtesting.WithBinaryContentMode(),
	eventingtesting.WithExemptHandshake(),
	eventingtesting.WithAtLeastOnceQOS(),
	eventingtesting.WithDefaultWebhookAuth(),
)
```
</details>


### Functional Options pattern

**Summary**: Functional Options is a pattern to simplify configuration of objects and at the same time make this configuration more expressive.

<details>
	<summary>Classic way</summary>
	
```go
package main

import "fmt"

type foo struct {
	name string
	info string
	bar  *bar
}

type bar struct {
	name string
}

func NewFoo(name, info string, bar *bar) *foo {
	return &foo{
		name: name,
		info: info,
		bar:  bar,
	}
}

func main() {

	bar := &bar{"bar"}

	foo := NewFoo("name",
		"info",
		bar,
	)

	fmt.Printf("%v", foo)
}
```
</details>

The pattern changes this code to
<details>
	<summary>With pattern</summary>
	
```go
package main

import "fmt"

type foo struct {
	name string
	info string
	bar  *bar
}

type bar struct {
	name string
}

type Opt func(f *foo)

func WithInfo(info string) Opt {
	return func(f *foo) {
		f.info = info
	}
}

func WithBar(bar *bar) Opt {
	return func(f *foo) {
		f.bar = bar
	}
}

func NewFoo(name string, opts ...Opt) *foo {
	foo := &foo{
		name: name,
	}
	for _, opt := range opts {
		opt(foo)
	}
	return foo
}

func main() {

	bar := &bar{"bar"}

	foo := NewFoo("name",
		WithInfo("info"),
		WithBar(bar),
	)

	fmt.Printf("%v", foo)

}
```

</details>

**Advantage**:
The benefit of this pattern is that with Functional Options it is easy to omit unnecessary arguments. For example, omitting "info" in the classic way would require explicitly creating a new function to do this, or requires the user to explicitly pass "info" as an empty string. In the second implementation, you just omit the call to `WithInfo`. 
Fundamentally, this pattern can increase readability and allows reducing configurations to the necessary information.

The pattern works very well in tests where similar objects must be constructed multiple times. Here, it is essential that the developer quickly sees what the important configuration of these objects for a given test case is.

#### Guidelines for the pattern

* Don't create Option functions for required arguments. It makes no sense to make them optional because required arguments must not be omitted.
* Names for Option functions should start with `With` followed by a short description of their purpose. Keep in mind that these names might be shortened by the developer's IDE. Important information must be easily spotted in those names.
Examples:

```
WithSinkFromSVC // configures a sink. The values are taken from a service
WithProtocolBEB // configures the Protocol to be BEB
```

#### Bad practice

**Example 1**: Return function from WithSomething method
<details>
	<summary>Don't</summary>

```go
// source: https://github.com/kyma-project/kyma/pull/13242/files#diff-2f168bb71c1ca8d2f5781bac737393acd9bc3a1f829bdbe91093b7db54972433L130
func WithExemptHandshakeBefore(p *eventingv1alpha1.ProtocolSettings) {
		p.ExemptHandshake = utils.BoolPtr(true)
}

subscription.Spec.ProtocolSettings = reconcilertesting.NewProtocolSettings(
  eventingtesting.WithExemptHandshakeBefore,
)
```
</details>

<details>
	<summary>Do</summary>

In contrast to `WithExemptHandshakeBefore`, `WithExemptHandshakeAfter` returns a function of type `ProtoOpt`. However, `WithExemptHandshakeBefore` itself is of type `ProtoOpt`. For simplicity and consistency, you should **always** return a `ProtoOpt` from **inside** the `With` function.

```go
func WithExemptHandshakeAfter() ProtoOpt {
	return func(p *eventingv1alpha1.ProtocolSettings) { // instead return a function here
		p.ExemptHandshake = utils.BoolPtr(true)
	}
}

subscription.Spec.ProtocolSettings = reconcilertesting.NewProtocolSettings(
	eventingtesting.WithExemptHandshakeAfter(), // and execute it here
)
```
</details>

**Example 2**: Usage of Functional Options pattern outside of factory
<details>
	<summary>Don't</summary>

`WithServiceBefore` is used outside of `NewAPIRule`. Moreover, it is used to apply a **side-effect** (setting apiRule.Spec.Service). This is **weird** because it misuses the `WithServiceBefore`, which shall only be used as an argument to a function that supports an APIRuleOption.

```go
// source: https://github.com/kyma-project/kyma/pull/13242/files#diff-2f168bb71c1ca8d2f5781bac737393acd9bc3a1f829bdbe91093b7db54972433L130

func WithServiceBefore(host, svcName string, apiRule *apigatewayv1alpha1.APIRule) {
	port := uint32(443)
	isExternal := true
	apiRule.Spec.Service = &apigatewayv1alpha1.Service{
		Name:       &svcName,
		Port:       &port,
		Host:       &host,
		IsExternal: &isExternal,
	}
}

apiRule = NewAPIRule(
	subscriptionWithoutWebhookAuth, 
	WithPath
)
WithServiceBefore(host, svcName, apiRule) // this applies a side effect
```
</details>

<details>
	<summary>Do</summary>

In the following example, `WithServiceAfter` is used as an argument to `NewAPIRule`.

```go
func WithServiceAfter(name, host string) APIRuleOption {
	return func(r *apigatewayv1alpha1.APIRule) {
		port := uint32(443)
		isExternal := true
		r.Spec.Service = &apigatewayv1alpha1.Service{
			Name:       &name,
			Port:       &port,
			Host:       &host,
			IsExternal: &isExternal,
		}
	}
}

apiRule = NewAPIRule(
	subscriptionWithoutWebhookAuth,
	WithPath(),
	WithServiceAfter(svcName, host), // instead pass it as an argument to NewAPIRule
)
```
</details>

**Example 3**: Side effect as new function

<details>
  <summary>Don't</summary>

`WithStatusReady` is not used inside a function that accepts a functional option as input.
As a reader, you would not expect that it performs a **side-effect** on the `apiRuleNew`.

```go
// source: https://github.com/kyma-project/kyma/blob/a84f76f674babfc63e369bbba44de76499cf475d/components/eventing-controller/controllers/subscription/beb/reconciler_test.go#L592
getAPIRule(ctx, apiRuleNew).Should(And(
	HaveNotEmptyHost(),
	HaveNotEmptyAPIRule(),
))
WithStatusReady(apiRuleNew) // <= WithStatusReady not used inside a function that accepts a functional option as input (for example, `APIRuleOption`)
```

</details>

<details>
  <summary>Do</summary>

To make the side-effect more obvious, the following example uses a function named `MarkReady` instead.

```go
getAPIRule(ctx, apiRuleNew).Should(And(
	HaveNotEmptyHost(),
	HaveNotEmptyAPIRule(),
))
MarkReady(apiRuleNew) // instead consider using another function that explicitly states from the name that it applies a side effect
```
</details>

### Logging guidelines

Standardize logs across Eventing components according to the following general and code-specific rules.

#### General rules

- Each log message should be as short and meaningful as possible.
- Logs should be aligned with the unified way of logging inside Kyma (see [Consistent Logging](https://github.com/kyma-project/community/blob/main/concepts/observability-consistent-logging/README.md)).
- Each log message should have enough context to convey what happened.
- Each log message should have the proper log level (see [Unified approach to logging levels](https://github.com/kyma-project/community/blob/main/concepts/observability-consistent-logging/unified-approach-to-logging-levels.md)).

#### Code-specific rules:

- Consider using the `uber-go/zap` logging library. Label the logger by naming it with the component name and add the context whenever it is possible (see [Log structure](https://github.com/kyma-project/community/blob/main/concepts/observability-consistent-logging/improvement-of-log-messages-usability.md#log-structure)):
  <details>
        <summary>Example</summary>

  ```go
  namedLogger := r.logger.WithContext().Named("logger-name").With("backend", "BEB")
  r.namedLogger().Info("Creating Event Publisher deployment")
  ```
  will output:
  ```
  {"level":"INFO","timestamp":"2022-07-04T12:00:36+02:00","logger":"logger-name","caller":"backend/reconciler.go:741","message":"Creating Event Publisher deployment","context":{"backend":"BEB"}}
  ```
  </details>


- If you return the same error as a result for the `Reconcile()` method, don't log the error. That's because Kubebuilder will output it too, so the user gets two very similar logs one after another:
  <details>
      <summary>Example</summary>

    ```go
    namedLogger.Errorw("Failed to sync BEB subscription", "error", err)
    updateErr := r.updateSubscription(ctx, subscription, log); updateErr != nil {
          return ctrl.Result{}, errors.Wrap(err, updateErr.Error())
    }
    ```
  will result in duplication of logs:
    ```
    {"level":"ERROR","timestamp":"2022-07-01T08:20:26Z","logger":"beb-subscription-reconciler","caller":"beb/reconciler.go:275","message":"Failed to sync BEB subscription","context":{"kind":"Subscription","version":2,"namespace":"tunas-testing","name":"test-noapp","error":"prefix not found"}}
    {"level":"ERROR","timestamp":"2022-07-01T08:20:26Z","caller":"controller/controller.go:326","message":"Reconciler error","context":{"controller":"beb-subscription-reconciler","object":{"name":"test-noapp","namespace":"tunas-testing"},"namespace":"tunas-testing","name":"test-noapp","reconcileID":"9994dd3e-0104-4170-82aa-79df9ec41af1","error":"prefix not found"}}
    ```
  </details>


- Capitalize the component names, for example Event Publisher, or EventingBackend:
  ```go
  namedLogger.Debug("Event Publisher deployment not ready...")
  ```
- Use the standardized structure:
  *past tense starting with **Failed to...**, followed by the error wrapped with some meaningful context*:
    ```go
    namedLogger.Errorw("Failed to update Event Publisher secret", "error", err)
    ```
-  Capitalize the first letter of the first word in the logs, including the error logs:
    ```go
    namedLogger.Debug("Creating secret for BEB publisher")
    ```

## Testing pyramid in Kyma

The following section describes the different testing levels for Kyma. These levels are usually seen as a [pyramid](https://en.wikipedia.org/wiki/Test_automation#Testing_at_different_levels):
- At the bottom of the pyramid, you should have the **most** test cases because they are usually very **fast**. However, you are limited in what you can test because external dependencies are **mocked**. Unit tests are an example of this.
- At the top of the pyramid, there are usually **fewer** tests because these tests take **longer**. However, these tests are a more realistic scenario of how the end user will use our product. E2E tests are an example of this.

### Controller Integration Tests (kubebuilder)

**Test Setup**
Controllers which are bootsprapped by the kubebuilder framework use `Ginkgo` (for testing) and `Gomega` (for assertions).
Using the `controller-runtime/pkg/envtest` package, a Kubernetes **control-plane** (**API server** and **etcd**) is started locally once before the first test runs (`BeforeSuite`).
You can also use a custom Kubernetes cluster (see [USE_EXISTING_CLUSTER](https://book.kubebuilder.io/reference/envtest.html) environment variable).

**Goal**
The goal of a controller integration test is to test the Kubernetes controller in a **limited** Kubernetes **environment**. That means that dependencies of the controller - such as other controllers (APIRule controller) or external eventing system - are not present. These systems are only present in higher hierarchies of the testing pyramid.

Controller integration tests should be reserved for cases when we care about the eventual state of the object or need a Kubernetes cluster.

**See Also**
- [kubebuilder documentation](https://book.kubebuilder.io/reference/envtest.html)
- [envtest godoc](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/envtest)

### Unit Tests

[Unit tests](https://en.wikipedia.org/wiki/Unit_testing) are used to test individual **units** of the code. To test the unit in an **isolated** fashion, external **dependencies** are typically **mocked**.
Unit tests provide a very fast **feedback cycle** and help discovering bugs in a very **early stage**.
The advantage of unit testing is that "[Code can be impossible or difficult to unit test if poorly written, thus unit testing can force developers to structure functions and objects in better ways.](https://en.wikipedia.org/wiki/Unit_testing)"
The big disadvantage of unit tests is that they are tightly coupled to the code, so rewriting the code may require new or modified unit tests.

## Testing guidelines 

The following section describes ugly tests or problems that we discovered in our code base. The testing guidelines provide suggestions how the tests can be written instead.
The guides should be used whenever possible but are not set in stone. There might be cases where the proposed solution does not fit and that is ok.

**Best practice**:
- Find a **balance** between [KISS](https://en.wikipedia.org/wiki/KISS_principle) (keep it simple stupid) and [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) (don't repeat yourself). Tests should be easily `understandable`. At the same time, you should **avoid** code **duplication** because it gets **unmaintainable** very fast.
- **Split** heavy test **setup** from actual test.
- Test **one thing** at a time. This aligns with the KISS principle as well. Focus on the **main concern** of a test and don't test more than what you need. Comments on tests are welcome, but might also be an indication for a test that is either (1) too complex or (2) tests too many things simultaneously. 
Concentrating on the main concern makes the individual tests shorter, easier to understand, better to maintain and expresses the intention of the test better. If multiple things are tested at the same time, there is **no guarantee** that someone will remove this secondary thing from the test. However, if there are multiple tests, you can easily see that a test case was removed and argue whether this is ok or not.
- **Don't** test the **unforeseen**. Instead, think about the code and derive useful test scenarios (best and worst case). Think about what is the **happy path** ? What is the **unhappy path** ?
- Especially complex tests need **documentation** so that the reader understands them better.

**Goals**:
- Avoid regression: Whenever you close a bug, make sure that this bug will not get reintroduced. 
- Express intent of code: Code should be self-explanatory and documented. In addition, tests help in understanding the intention of the code by providing test cases where you can easily see the (1) input to the code and (2) the expected outcome.
- Drive development: Yes, tests can help drive development. They enable you to run code that is otherwise hard to run because it needs a complex setup. Imagine that in order to test a Kubernetes controller, you need to build a Docker image, push the image, and deploy the controller. This is a very time-consuming approach. You should always aim for a **short feedback cycle**.
<!-- inspired by https://fossa.com/blog/golang-best-practices-testing-go/#Why -->


### Regression Tests

Whenever you close a bug, add a test that prevents the bug from getting re-introduced (*regression*).
Provide context to help understanding the purpose of the test. For example, if a test is based on an issue, you must **link** the test to the **issue**.

The following example shows how the issue must be linked to the test:

```go
// TestBugIsFixed ensures something.
// issue: https://github.com/kyma-project/kyma/issues/12979
func TestBugIsFixed(t *testing.T) {
  // test goes here ...
}
```

### Separate test setup from actual test 

<details>
  <summary>TestSendCloudEvent example</summary>

```go 
// source: https://github.com/kyma-project/kyma/blob/d6662ab956c18cfc9b3e0c7deebd26da3a56ae77/components/event-publisher-proxy/pkg/sender/nats_test.go#L58 

func TestSendCloudEvent(t *testing.T) {
	//////////////////////////////
	// test setup start
	//////////////////////////////
	logger := logrus.New()
	logger.Info("TestNatsSender started")

	// Start Nats server
	natsServer := testingutils.StartNatsServer()
	assert.NotNil(t, natsServer)
	defer natsServer.Shutdown()

	// connect to nats
	bc := pkgnats.NewBackendConnection(natsServer.ClientURL(), true, 1, time.Second)
	err := bc.Connect()
	assert.Nil(t, err)
	assert.NotNil(t, bc.Connection)

	// create message sender
	ctx := context.Background()
	sender := NewNatsMessageSender(ctx, bc, logger)

	//////////////////////////////
	// test setup end
	//////////////////////////////

	// subscribe to subject
	done := make(chan bool, 1)
	validator := testingutils.ValidateNatsMessageDataOrFail(t, fmt.Sprintf(`"%s"`, testingutils.CloudEventData), done)
	testingutils.SubscribeToEventOrFail(t, bc.Connection, testingutils.CloudEventType, validator)

	// create cloudevent with default data (testing.CloudEventData)
	ce := cloudevents.NewEvent()
	ce.SetType(testingutils.CloudEventType)
	err := json.Unmarshal([]byte(testingutils.StructuredCloudEventPayloadWithCleanEventType), &ce)
	assert.Nil(t, err)

	// send the event to NATS and assert that the expectedStatus is returned from NATS
	status, err := testEnv.natsMessageSender.Send(ctx, &ce)
	assert.Nil(t, err)
	assert.Equal(t, status, http.StatusNoContent)

	// wait for subscriber to receive the messages
	err = testingutils.WaitForChannelOrTimeout(done, time.Second*3)
	assert.NoError(t, err, "Subscriber did not receive the message")
}
```

</details>

The test `TestSendCloudEventWithReconnect` contains a lot of code only to setup the test environment (see `test setup start` and `test setup end` markers). To make the test itself **shorter** and **cleaner**, extract the test setup code.

The idea is to move all parts of the test setup into a struct and create a helper method to start and stop the environment.

<details>
  <summary>TestEnvironment struct</summary>

```go 
// TestEnvironment contains the necessary entities to perform NATS integration tests
type TestEnvironment struct {
	context context.Context
	logger  *logrus.Logger

	// natsServer is a real NATS server for integration testing.
	natsServer *server.Server
	// backendConnection is a connection to the NATS server.
	backendConnection *pkgnats.Connection
	// natsMessageSender is a sender for publishing events to the NATS server.
	natsMessageSender *NatsMessageSender
} 
```
</details>

The struct contains the NATS server, a client to send messages to NATS, and the connection to NATS, as well as a context and a logger.
In addition to the struct, create a helper method to setup the actual environment:

<details>
  <summary>setupTestEnvironment()</summary>

```go
func setupTestEnvironment(t *testing.T, connectionOpts ...pkgnats.BackendConnectionOpt) TestEnvironment {
	// ... some code is left our for readability here
	natsServer := testingutils.StartNatsServer()
	return TestEnvironment{
		context:           ctx,
		natsServer:        natsServer,
		backendConnection: bc,
		natsMessageSender: sender,
		logger:            logger,
	}
}
```

A common problem with this approach is that the caller (the test) must do the cleanup.
However, we can shift the cleanup to the helper method as well using `testing.CleanUp`:

<details>
  <summary>setupTestEnvironment() with t.Cleanup</summary>

```go
func setupTestEnvironment(t *testing.T, connectionOpts ...pkgnats.BackendConnectionOpt) TestEnvironment {
	// ... some code is left out for readability here
	natsServer := testingutils.StartNatsServer()
	t.Cleanup(func() {
		natsServer.Shutdown()
	})
	return TestEnvironment{
		context:           ctx,
		natsServer:        natsServer,
		backendConnection: bc,
		natsMessageSender: sender,
		logger:            logger,
	}
}
```
</details>

<details>
  <summary>t.Cleanup vs defer</summary>

[t.Cleanup](https://cs.opensource.google/go/go/+/go1.17.6:src/testing/testing.go;l=892) is executed in the same order as defer (`LIFO`), but defers are executed before t.Cleanup.

```go
func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("cleanup: 1")
	})
	defer func() {
		fmt.Println("defer: 1")
	}()
	t.Cleanup(func() {
		fmt.Println("cleanup: 2")
	})
	defer func() {
		fmt.Println("defer: 2")
	}()
	t.Cleanup(func() {
		fmt.Println("cleanup: 3")
	})
	defer func() {
		fmt.Println("defer: 3")
	}()
}

// output:
defer: 3
defer: 2
defer: 1
cleanup: 3
cleanup: 2
cleanup: 1

```
</details>

After the refactoring, the test looks like this:

<details>
  <summary>Refactored test</summary>

```go
func TestSendCloudEvent(t *testing.T) {
	//////////////////////////////
	// test setup start
	//////////////////////////////
	testEnv := setupTestEnvironment(t, pkgnats.WithMaxReconnects(tc.givenRetries))

	//////////////////////////////
	// test setup end
	//////////////////////////////

	// subscribe to subject
	done := make(chan bool, 1)
	validator := testingutils.ValidateNatsMessageDataOrFail(t, fmt.Sprintf(`"%s"`, testingutils.CloudEventData), done)
	testingutils.SubscribeToEventOrFail(t, testEnv.BackendConnection, testingutils.CloudEventType, validator)

	// create cloudevent with default data (testing.CloudEventData)
	ce := cloudevents.NewEvent()
	ce.SetType(testingutils.CloudEventType)
	err := json.Unmarshal([]byte(testingutils.StructuredCloudEventPayloadWithCleanEventType), &ce)
	assert.Nil(t, err)

	// send the event to NATS and assert that the expectedStatus is returned from NATS
	status, err := testEnv.backendConnection.Send(testEnv.context, &ce)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, status)

	// wait for subscriber to receive the messages
	err = testingutils.WaitForChannelOrTimeout(done, time.Second*3)
	assert.NoError(t, err, "Subscriber did not receive the message")
}
```
</details>
</details>

### GWT (Given-When-Then) pattern

The [Given-When-Then](https://en.wikipedia.org/wiki/Given-When-Then) pattern is a good way to structure code inside a test by adding comments so that the following intentions are made clear:
1. Given: The **preconditions** which are required for the test.
1. When: This is usually the time when you want to call the `functionUnderTest`.
1. Then: This is the part where you make sure that the `functionUnderTest` is working as expected by using **assertions**.

```go
// given
// create a subscription using the mocked client
sub, err := client.Create(subscription)
assert.NotNil(t, err)

// when
// call the function
functionUnderTest(sub)

// then
assert.Nil(t, sub)
```

Sometimes you need to add more comments to the tests (in addition to the GWT comments). We recommend to add them underneath the GWT comments.

Especially when the test code is very long, the GWT comment pattern improves the readability of the code.

### Table-driven tests

[Table-driven tests](https://go.dev/blog/subtests) are a very common way of expressing **multiple test cases** while **sharing** the same test setup, thus avoiding test duplication. 
They also provide context and meaning to the test cases.

**Conventions**:
- Every test must have a `name`. Input variables to the test are prefixed with `given`, and expected output is prefixed with `want`.
  - Keep the `name` short. For example, use **"event order.created received"** instead of "test that event order.created was received" or "ensure that event order.created was received".
- To improve readability, always set the **field names** when initializing the test case struct.

#### Best practice

<details>
  <summary>Normal table-driven test</summary>

```go
func TestSomething(t *testing.T) {
	testCases := []struct {
		name           string
		givenAttribute string
		wantAttribute  string
	}{
		{
			name:           "meaningful test name",
			givenAttribute: "this attribute is an input to the test",
			wantAttribute:  "this is the expected output",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := functionUnderTest(tc.givenAttribute)
			// check res == tc.wantAttribute
		})
	}
}
```

</details>

<details>
  <summary>2-dimensional table-driven test</summary>

There may be situations when you need a table-driven test with more than one dimension. See the following best practice example for a 2-dimensional table-driven test:

```go
func TestTwoDimensions(t *testing.T) {
	// first dimension
	testCases := []struct {
		name               string
		givenSender        CESender
		wantHTTPStatusCode int
	}{
		{
			name:        "binary cloud event sender",
			givenSender: myCEBinarySender,
		},
	}
	// second dimension
	cloudEvents := []struct {
		name, givenCEType string
	}{
		{
			name:               "proper cloud event",
			givenCEType:        "order.created.v1",
			wantHTTPStatusCode: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for _, ce := range cloudEvents {
				ce := ce
				t.Run(ce.name, func(t *testing.T) {
					t.Parallel()
					res := functionUnderTest(tc.givenSender, ce.givenCEType)
					// check res == ce.wantHTTPStatusCode
				})
			}
		})
	}
}
```

<details>
  <summary>Reason for using nested t.Run</summary>

To understand why we use t.Run in a nested way, look at the output that both tests produce.
When not using t.Run in a nested way, the test could look like this:

```go
// ...
for _, tc := range testCases {
	tc := tc
	for _, ce := range ce {
		ce := ce
		t.Run(tc.name + " - " + ce.name, func(t *testing.T) { // only this line changed
			t.Parallel()
			res := functionUnderTest(tc.givenSender, ce.givenCEType)
			// check res == ce.wantHTTPStatusCode
		})
	}
}

$ go test -v
=== RUN   TestTwoDimensions
=== RUN   TestTwoDimensions/binary_cloud_event_sender_-_proper_cloud_event
--- PASS: TestTwoDimensions (0.00s)
    --- PASS: TestTwoDimensions/binary_cloud_event_sender_-_proper_cloud_event (0.00s)
PASS
ok      test    0.199s
```

The output of the test with nested t.Run looks like this:

```shell
$ go test -v
=== RUN   TestTwoDimensions
=== RUN   TestTwoDimensions/binary_cloud_event_sender
=== RUN   TestTwoDimensions/binary_cloud_event_sender/proper_cloud_event
--- PASS: TestTwoDimensions (0.00s)
    --- PASS: TestTwoDimensions/binary_cloud_event_sender (0.00s)
        --- PASS: TestTwoDimensions/binary_cloud_event_sender/proper_cloud_event
 (0.00s)
PASS
ok      test    0.182s
```

When you look at the output of both examples, you can see that the test name is different: (`binary_cloud_event_sender_-_proper_cloud_event` vs `binary_cloud_event_sender/proper_cloud_event`). Each subtest adds `/<test_name>` to the test name.
The **advantages** of using t.Run in a nested way are:
- The test name is easier to read (`binary_cloud_event_sender/proper_cloud_event`).
- There is no need to use a combined name (`tc.name+" - "+ce.name`).
- The nesting of t.Run is displayed in a nicer way (in IDEs, this is used to group tests and make them collapsable).

</details>

</details>


#### Bad practice

**Example 1**: Test case struct initialized without setting field name

<details>
  <summary>Don't</summary>

```go
func TestSomething(t *testing.T) {
	testCases := []struct {
		name           string
		givenAttribute string
		wantAttribute  string
	}{
		{ // field names are not set
			"meaningful test name",
			"this attribute is an input to the test",
			"this is the expected output",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := functionUnderTest(tc.givenAttribute)
			fmt.Println(res)
		})
	}
}
```
</details>

<details>
  <summary>Do</summary>

```go
func TestSomething(t *testing.T) {
	testCases := []struct {
		name           string
		givenAttribute string
		wantAttribute  string
	}{
		{ // field names are set
			name:           "meaningful test name",
			givenAttribute: "this attribute is an input to the test",
			wantAttribute:  "this is the expected output",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res := functionUnderTest(tc.givenAttribute)
			fmt.Println(res)
		})
	}
}
```

</details>

**Example 2**: Two-dimensional table test using exec function

<details>
  <summary>Don't</summary>

The following example shows a version of a **two-dimensional** **table test**. The test is encapsulated in a function named **exec** and called multiple times with different arguments. This is the **first dimension** of the test.

The second dimension of the test is visible in the `t.Run` command, which creates a subtest for each entry in `handlertest.TestCasesForCloudEvents`.

Technically, you can use it this way. However, it is better to use a **pure two-dimensional table test** instead.

```go
// source: https://github.com/nachtmaar/kyma/blob/13029-retry-fail-publish/components/event-publisher-proxy/pkg/handler/nats/handler_test.go
func TestNatsHandlerForCloudEvents(t *testing.T) {
	exec := func(t *testing.T, applicationName, expectedNatsSubject, eventTypePrefix, eventType string) {
		test.logger.Info("TestNatsHandlerForCloudEvents started")

		// setup test environment
		publishEndpoint := fmt.Sprintf("http://localhost:%d/publish", test.natsConfig.Port)
		subscription := testingutils.NewSubscription(testingutils.SubscriptionWithFilter(testingutils.MessagingNamespace, eventType))
		cancel := test.setupResources(t, subscription, applicationName, eventTypePrefix)
		defer cancel()

		// prepare event type from subscription
		assert.NotNil(t, subscription.Spec.Filter)
		assert.NotEmpty(t, subscription.Spec.Filter.Filters)
		eventTypeToSubscribe := subscription.Spec.Filter.Filters[0].EventType.Value

		// connect to nats
		bc := pkgnats.NewConnection(
			test.natsURL,
			pkgnats.WithMaxReconnects(3),
			pkgnats.WithRetryOnFailedConnect(true),
			pkgnats.WithReconnectWait(time.Second),
		)
		err := bc.Connect()
		assert.Nil(t, err)
		assert.NotNil(t, bc.Connection)

		// publish a message to NATS and validate it
		validator := testingutils.ValidateNatsSubjectOrFail(t, expectedNatsSubject)
		testingutils.SubscribeToEventOrFail(t, bc.Connection, eventTypeToSubscribe, validator)

		// nolint:scopelint
		// run the tests for publishing cloudevents
		for _, testCase := range handlertest.TestCasesForCloudEvents {
			t.Run(testCase.Name, func(t *testing.T) {
				body, headers := testCase.ProvideMessage()
				resp, err := testingutils.SendEvent(publishEndpoint, body, headers)
				if err != nil {
					t.Errorf("Failed to send event with error: %v", err)
				}
				_ = resp.Body.Close()
				if testCase.WantStatusCode != resp.StatusCode {
					t.Errorf("Test failed, want status code:%d but got:%d", testCase.WantStatusCode, resp.StatusCode)
				}
				if testingutils.Is2XX(resp.StatusCode) {
					metricstest.EnsureMetricLatency(t, test.collector)
				}
			})
		}
	}

	// do not to change the cloudevent, even if its event-type contains none-alphanumeric characters or the event-type-prefix is empty
	exec(t, testingutils.ApplicationName, testingutils.CloudEventTypeNotClean, testingutils.MessagingEventTypePrefix, testingutils.CloudEventTypeNotClean)
	exec(t, testingutils.ApplicationName, testingutils.CloudEventTypeNotClean, testingutils.MessagingEventTypePrefixEmpty, testingutils.CloudEventTypeNotCleanPrefixEmpty)
	exec(t, testingutils.ApplicationNameNotClean, testingutils.CloudEventTypeNotClean, testingutils.MessagingEventTypePrefix, testingutils.CloudEventTypeNotClean)
	exec(t, testingutils.ApplicationNameNotClean, testingutils.CloudEventTypeNotClean, testingutils.MessagingEventTypePrefixEmpty, testingutils.CloudEventTypeNotCleanPrefixEmpty)
}

var (
	TestCasesForCloudEvents = []struct {
		Name           string
		ProvideMessage func() (string, http.Header)
		WantStatusCode int
	}{
		{
			Name: "Structured CloudEvent without id",
			ProvideMessage: func() (string, http.Header) {
				return testingutils.StructuredCloudEventPayloadWithoutID, testingutils.GetStructuredMessageHeaders()
			},
			WantStatusCode: http.StatusBadRequest,
		},
	}
)
```
</details>

<details>
  <summary>Do</summary>

This is the same test as before, but each `exec` function was replaced with an entry in the `testCases` list. For each entry in `testCases` and `handlertest.TestCasesForCloudEvents`, a subtest is started using `t.Run`.

```go
func TestNatsHandlerForCloudEvents(t *testing.T) {
	testCases := []struct {
		name                 string
		givenApplicationName string
		givenEventType       string
		givenEventTypePrefix string
		wantNatsSubject      string
	}{
		{
			name:                 "meaningful test name",
			givenApplicationName: testingutils.ApplicationName,
			givenEventType:       testingutils.CloudEventTypeNotClean,
			givenEventTypePrefix: testingutils.MessagingEventTypePrefix,
			wantNatsSubject:      testingutils.CloudEventTypeNotClean,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for _, ceTestCase := range handlertest.TestCasesForCloudEvents {
				ceTestCase := ceTestCase  
				t.Run(ceTestCase.Name, func(t *testing.T) {
					// other code is unchanged
				})
			}
		})
	}
}
```
</details>

**Example 3:** Don't repeat yourself (**DRY**) and consider a table-driven test instead.

<details>
  <summary>Don't</summary>

The following example shows two nearly identical tests. The only difference is that `TestSendCloudEventWithReconnect` additionally **closes** the **connection** before sending the event, in contrast to `TestSendCloudEvent`.
The test `TestSendCloudEventWithReconnect` has another problem that prevents rewriting both tests as a table test: It sends an event twice. Once with an open connection, then with a closed connection. The **main concern** of the test is to ensure that the **connection** is **re-established**. Then it is in closed state.
However, it is enough to close the connection before sending the first event. There is no need to send the event twice. Reducing the test to the **bare minimum** enables us to rewrite it as a table-driven test.

```go 
// source: https://github.com/kyma-project/kyma/blob/d6662ab956c18cfc9b3e0c7deebd26da3a56ae77/components/event-publisher-proxy/pkg/sender/nats_test.go#L58 

func TestSendCloudEvent(t *testing.T) {
	logger := logrus.New()
	logger.Info("TestNatsSender started")

	// Start Nats server
	natsServer := testingutils.StartNatsServer()
	assert.NotNil(t, natsServer)
	defer natsServer.Shutdown()

	// connect to nats
	bc := pkgnats.NewBackendConnection(natsServer.ClientURL(), true, 1, time.Second)
	err := bc.Connect()
	assert.Nil(t, err)
	assert.NotNil(t, bc.Connection)

	// create message sender
	ctx := context.Background()
	sender := NewNatsMessageSender(ctx, bc, logger)

	// subscribe to subject
	done := make(chan bool, 1)
	validator := testingutils.ValidateNatsMessageDataOrFail(t, fmt.Sprintf(`"%s"`, testingutils.CloudEventData), done)
	testingutils.SubscribeToEventOrFail(t, bc.Connection, testingutils.CloudEventType, validator)

	// create cloudevent
	ce := testingutils.StructuredCloudEventPayloadWithCleanEventType
	event := cloudevents.NewEvent()
	event.SetType(testingutils.CloudEventType)
	err = json.Unmarshal([]byte(ce), &event)
	assert.Nil(t, err)

	// send cloudevent
	status, err := sender.Send(ctx, &event)
	assert.Nil(t, err)
	assert.Equal(t, status, http.StatusNoContent)

	// wait for subscriber to receive the messages
	if err := testingutils.WaitForChannelOrTimeout(done, time.Second*3); err != nil {
		t.Fatalf("Subscriber did not receive the message with error: %v", err)
	}
}

func TestSendCloudEventWithReconnect(t *testing.T) { 

	logger := logrus.New() 
	logger.Info("TestNatsSender started") 

	// Start Nats server 
	natsServer := testingutils.StartNatsServer() 
	assert.NotNil(t, natsServer) 
	defer natsServer.Shutdown() 
	// connect to nats 
	bc := pkgnats.NewBackendConnection(natsServer.ClientURL(), true, 10, time.Second) 
	err := bc.Connect() 
	assert.Nil(t, err) 
	assert.NotNil(t, bc.Connection) 

	// create message sender 
	ctx := context.Background() 
	sender := NewNatsMessageSender(ctx, bc, logger) 

	// subscribe to subject 
	done := make(chan bool, 1) 
	validator := testingutils.ValidateNatsMessageDataOrFail(t, fmt.Sprintf(`"%s"`, testingutils.CloudEventData), done) 
	testingutils.SubscribeToEventOrFail(t, bc.Connection, testingutils.CloudEventType, validator) 

	// create cloudevent 
	ce := cloudevents.NewEvent() 
	ce.SetType(testingutils.CloudEventType) 
	err = json.Unmarshal([]byte(testingutils.StructuredCloudEventPayloadWithCleanEventType), &ce) 
	assert.Nil(t, err) 

	sendEventAndAssertStatus(ctx, t, sender, &ce, http.StatusNoContent) 

	// wait for subscriber to receive the messages 
	if err := testingutils.WaitForChannelOrTimeout(done, time.Second*3); err != nil { 
		t.Fatalf("Subscriber did not receive the message with error: %v", err) 
	} 

	// close connection 
	bc.Connection.Close() 
	assert.True(t, bc.Connection.IsClosed()) 
	sendEventAndAssertStatus(ctx, t, sender, &ce, http.StatusNoContent) 
} 
``` 

</details>

<details>
  <summary>Do</summary>

```go 
func TestSendCloudEventsToNats(t *testing.T) {
	testCases := []struct { 
		name                               string
		givenRetries                       int
		wantHTTPStatusCode                 int
		wantReconnectAfterConnectionClosed bool
	}{ 
		{
			name:                              "sending event to NATS works",
			givenRetries:                      1,
			wantHTTPStatusCode:                http.StatusNoContent,
			wantClosedConnectionBeforeSending: false,
		},
		{
			name:               "sending event to NATS works given a closed connection",
			givenRetries:       10,
			wantHTTPStatusCode: http.StatusNoContent,
			// Close connection before sending so we can check the reconnect behaviour of the NATS connection.
			wantClosedConnectionBeforeSending: true,
		},
	} 
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testEnv := setupTestEnvironment(t, pkgnats.WithMaxReconnects(tc.givenRetries))

			// subscribe to subject
			subject := fmt.Sprintf(`"%s"`, testingutils.CloudEventData)
			// NOTE: we are using the testEnv.natsRecvConnection instead of testEnv.natsSendConnection because the latter will get reconnected based on wantClosedConnectionBeforeSending. This will fail the test when trying to receive a message.
			done := subscribeToSubject(t, testEnv.natsRecvConnection, subject)

			// create cloudevent with default data (testing.CloudEventData)
			ce := cloudevents.NewEvent()
			ce.SetType(testingutils.CloudEventType)
			err := json.Unmarshal([]byte(testingutils.StructuredCloudEventPayloadWithCleanEventType), &ce)
			assert.Nil(t, err)

			if tc.wantClosedConnectionBeforeSending {
				// close connection
				testEnv.natsSendConnection.Connection.Close()
				// ensure connection is closed
				// this is important because we want to test that the connection is re-established as soon as we send an event
				assert.True(t, testEnv.natsSendConnection.Connection.IsClosed())
			}

			// send the event to NATS and assert that the expectedStatus is returned from NATS
			status, err := testEnv.natsMessageSender.Send(testEnv.context, &ce)
			assert.Nil(t, err)
			assert.Equal(t, tc.wantHTTPStatusCode, status)

			// wait for subscriber to receive the messages
			err = testingutils.WaitForChannelOrTimeout(done, time.Second*3)
			assert.NoError(t, err, "Subscriber did not receive the message")
		})
	}
}
```
</details>

#### See also
- [Table-driven test basics](https://go.dev/blog/subtests)

### Provide test documentation on package level

It is very useful to have some documentation in the test file, describing how the test works from a bird's-eye view.

Put this documentation directly at the beginning of the file, before the `package` declaration and before the `imports`.
See the following example of the file *components/event-publisher-proxy/pkg/sender/nats_test.go*:

```go
// Tests in this file are integration tests.
// They use a real NATS server using github.com/nats-io/nats-server/v2/server.
// Messages are sent using NatsMessageSender interface.
package sender
```

With this approach, no godoc is generated for the package `sender` because it is a test file (ending with **_test** prefix).

### Consistency

#### Assertion library

The following example is from the event-publisher-proxy. The tests mostly use the `testify/assert` package for writing test assertions, but sometimes `testing.T` is used as well.

<details>
  <summary>Testify assertion and t.Errorf mixed</summary>

```go
// source: https://github.com/nachtmaar/kyma/blob/13029-retry-fail-publish/components/event-publisher-proxy/pkg/sender/nats_test.go
// send cloudevent
status, err := sender.Send(ctx, &event)
assert.Nil(t, err)
assert.Equal(t, status, http.StatusNoContent)

// wait for subscriber to receive the messages
if err := testingutils.WaitForChannelOrTimeout(done, time.Second*3); err != nil {
  t.Fatalf("Subscriber did not receive the message with error: %v", err)
}
```
</details>

The self-written assertion using `t.Fatalf` can be rewritten using `assert.NoError` as follows:

<details>
  <summary>Only testify assertions</summary>

```go
// send cloudevent
status, err := sender.Send(ctx, &event)
assert.Nil(t, err)
assert.Equal(t, status, http.StatusNoContent)

// wait for subscriber to receive the messages
err = testingutils.WaitForChannelOrTimeout(done, time.Second*3)
require.NoError(t, err, "Subscriber did not receive the message") // <= the custom message can be supplied to the testify assertion, there is no need for t.Fatalf anymore
```
</details>


