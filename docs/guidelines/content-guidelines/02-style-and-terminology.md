---
title: Style and terminology
---

When you're writing Kyma documentation, refer to the following guidelines for grammar, capitalization, and preferred word choices. These guidelines help that all contributors write in the same way to ensure a uniform flow throughout the whole Kyma documentation.

## Grammar

These are the generally accepted grammar rules for writing Kyma documentation.

### Active voice

Use active voice whenever possible. Active voice is clear, concise, and it avoids misinterpretation. It is also easier for non-native speakers to understand. Passive voice is indirect, uses more words, and can be misleading because it reverses the logical order of events.

✅ The endpoint path includes your service name.  
⛔️ Your service name is to be included in the endpoint path.

### Voice and tone

There are different tones for different types of technical documentation, which can range from instructional to somewhat conversational. The goal is always to support people using the product and, in blogs and release notes, also to help business users understand changes.

While writing Kyma documentation, use semi-formal style and imperative mood. The imperative mood tells the reader directly to do something. Use the imperative mood to write instructional documentation such as procedures and tutorials. Other moods can imply optional behavior.

Avoid click-level instructions unless the UI is confusing. Readers of Kyma documentation are assumed to be fairly tech-savvy.
Avoid using unnecessary words such as "please" or "remember". If there's just one way to do something, don't use "can".

✅ Add your {XYZ}.  
✅ Click **Add**.  
⛔️ Click the **Add** button.  
⛔️ Please, click **Add**.  
⛔️ Remember to click **Add**.

✅ Click **Upload** and select one or more documents.  
⛔️ If you want to upload a document, you can click **Upload**.

### Tenses

Use present tense. In technical writing, present tense is easier to read than past or future tense. Simple verbs are easier to read and understand than complex verbs, such as verbs in the progressive or perfect tense. The more concise and straightforward you are, the better.

✅ If the information does not match, an error message **appears**.  
⛔️ If the information does not match, an error message **will appear**.

### Pronouns

Use the second person and the pronouns "you," "your," and "yours" to speak directly to the reader. Do not use the first person pronouns "we," "us," or "let's."

### Contractions

It's okay to use contractions in the documentation from time to time. However, do not overuse it. Otherwise, your text becomes messy and hard to read.

✅ It's okay to use contractions.  
✅ It is okay to not use them as well.

### Articles

Always verify the use of the articles "a", "an", and "the" where appropriate. Use "the" when you refer to a specific example. Use "a" when you refer to something non-specific or hypothetical.

Whenever you refer to the name of our product or one of our components, don't use an article:  

✅ Kyma is awesome.  
⛔️ The Kyma is not awesome.
✅ Application Connector  
⛔️ the Application Connector

### Punctuation

Use colons and semicolons sparingly. Use the colon ( : ) to introduce a list of things. The semicolon ( ; ) separates two distinct clauses within one sentence. The phrase after a semicolon is a complete sentence. However, the preferred method is without a colon or semicolon.

Use serial commas. A missing serial comma can create confusion about whether the statement applies to each list item individually, or whether the last two items are connected.

✅ In your request, include the values for the request date, name, and ID.  
⛔️ In your request, include the values for the request date, name and ID.

Avoid using parenthesis. Use lists instead, to make your sentences as simple as possible.

✅ Consider which tasks, such as unit tests, linting, and compilation, are relevant and necessary for your example.  
⛔️ The author of each example should consider which tasks (i.e. unit tests, linting and compilation) are relevant and necessary for their example.

## Capitalization

### Lowercase words

Sentence case is the normal style in English: Lowercase for all words, uppercase for the first letter of the first word in a sentence, and of proper nouns and names.

Kyma documentation uses sentence case for headings (more details about [headings](03-formatting.md#headings)).

### Capitalized words

Capitalize the first letter of a proper noun and name.
Use sentence case for normal, everyday words in body text. Don't capitalize words just because they are "special" or "important".

Capitalize Kyma components such as Application Connector or API Gateway Controller.
Whenever you point to the outside sources, research whether the name of the source starts with a capital letter or not.

>**NOTE:** Kubernetes is capitalized. Do not write it in lowercase or in the abbreviated version.  
  ✅ Kubernetes  
  ⛔️ kubernetes  
  ⛔️ k8s

### Camel case

Camel case means writing words or phrases without a blank space, using a capitalized letter to separate the words. It's very common in Kubernetes terminology.

Use camel case for Kubernetes resources and custom resources.

✅ API Gateway is a Kubernetes controller, which operates on APIRule custom resources.  
⛔️ API Gateway is a Kubernetes controller, which operates on API Rule custom resources.

For plurals, add "s".
For titles and navigation, add blank spaces, so that it's natural language instead of camel case.

✅ ConfigMap, ConfigMaps (resource)
✅ Config Map, Config Maps (resource in titles and navigation)
⛔️ configmap
⛔️ config map
  
If you refer to the code specifically, format it as code (more details about [code font](03-formatting.md#code-font-and-bold-font)).

✅ `APIRule`
  
If the words are not used in relation to Kubernetes resources, do not capitalize them.

See the following examples of Kubernetes resources:

* ConfigMap
* CronJob
* CustomResourceDefinition (CRD) - note that "custom resource" alone isn't Kubernetes-specific, thus it's lowercase.
* Deployment
* Function
* Ingress
* Namespace
* Node
* PodPreset
* Pod
* ProwJob
* Secret
* Service
* ServiceBinding
* ServiceClass
* ServiceInstance

## Terminology

Use the American English spelling, not British English.

✅ The **color** of the message changes from blue to red if there are errors.  
⛔️ The **colour** of the message changes from blue to red if there are errors.

> **NOTE:** Do not use words such as "currently" or "now" to indicate that something is in the transitional phase of development. Avoid promising anything and mention only those components and functionalities that are already in use.

Here is the preferred terminology to use in the Kyma documentation:

| ✅                       | ⛔️                              | comment                           |
| ------------------------ | ------------------------------- | --------------------------------- |
| API Micro Gateway        | API Gateway                     |                                   |
| ID                       | id                              |                                   |
| and                      | + (plus), & (ampersand)         |                                   |
| backend                  | back end, back-end              |                                   |
| connect, connection      | integrate, integration          |                                   |
| custom resource          | Custom Resource, CustomResource |                                   |
| document                 | doc                             |                                   |
| email                    | e-mail                          |                                   |
| fill in                  | complete                        |                                   |
| frontend                 | front end, front-end            |                                   |
| key-value                | key/value, key:value            |                                   |
| micro frontend           | microfrontend, micro front-end  |                                   |
| must                     | have to, need to                |                                   |
| need                     | require                         |                                   |
| or                       | / (slash)                       |                                   |
| repository               | repo                            |                                   |
| run                      | execute                         |                                   |
| single sign on (SSO)     | single sign-on                  |                                   |
| the following, several,… | a specific number               |                                   |
| typically                | usually                         |                                   |
| use                      | utilize                         |                                   |
| using, with              | via                             |                                   |
| YAML (file format)       | yaml                            | If it's a file extension or file name, use `.yaml` (see [formatting](03-formatting.md)) |
| Prow Job (process)       | Prowjob, prowjob                | If it's a resource, use [camel case](#camel-case): "ProwJob" |
| must, can                | should                          | If mandatory, use “must”, if optional, use “can”. |
| (you) can                | it is possible to, allows you to, there is the possibility to, you have the option to, … | If there’s no other option, drop the “can”, simply use imperative. |
| that is                  | i.e.                            | If you must explain a statement with "i.e." or "that is,…", often the first statement adds little value and can be dropped completely, keeping just the part after "i.e.". |
| for example, such as     | e.g.                            | In the middle of a sentence, "such as" is better than "for example". ⛔️ Kyma has many components, for example, Observability, that improve your life. |
| cloud-native (adjective) | cloud native                    | Rule of thumb: Noun compounds usually don't need a hyphen, adjective compounds often do. |
| application              | app                             | Use "Application" to describe an external solution connected to Kyma through Application Connector, "application" to describe a microservice deployed on Kyma or other kind of software. |
| the following         | below, this, the described, ...    | or "as shown in the example"      |
| the previous, earlier | above, this, the described, ...    | or "as shown in the example"      |

### Command line arguments

Use short command line arguments whenever possible.

* `-n`, not `--namespace`

 Short command line arguments can differ between the tools as shown in the following example:

* Kubernetes: `kubectl -n` equals `kubectl --namespace`
* Helm: `helm -n` equals `helm --name`, not `helm --namespace`

In such a case, explain the context in the document.

> **TIP:** In case of any further doubts concerning the style and standards of writing, see [Microsoft Writing Style Guide](https://docs.microsoft.com/en-us/style-guide/welcome/).
