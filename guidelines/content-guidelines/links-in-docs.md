# Links in docs

These are the guidelines for making cross-references between the documents in the [`kyma/docs`](https://github.com/kyma-project/kyma/tree/master/docs) folder.

>**NOTE:** The linking works only on the [Kyma website](https://kyma-project.io/docs). Currently, the cross-references between [GitHub documentation](ttps://github.com/kyma-project/kyma/tree/master/docs) is not available.

## Links between documents in the same topic

If you want to link to another document in the same topic, create a reference using the `#{type}-{title}-{header}` pattern, where:
- `{type}` is a metadata type of the document that you want to reference.
- `{title}` is a metadata title of the document that you want to reference.
- `{header}` is a header located in the document that you want to reference.

>**NOTE:** All variables must consist of lowercase characters separated with dashes (-).

If the `{type}` doesn't exist, the pattern has the form of `#{title}-{title}-{header}`. If you want to create a reference to the whole `{type}`, use the `#{type}-{type}` pattern.

>**TIP:** You can copy the reference to the documentation directly from the website. See the same topic reference to the Helm Broker details document:

![Same topic reference](../../assets/reference-1.png)

### Links to the assets folder

To add a reference to a YAML, JSON, SVG, PNG, or JPG file located in the `assets` folder in the same topic, use GitHub relative links. For example, write `[Here](./assets/mf-namespaced.yaml) you can find a sample micro front-end entity.` When you click such a link on the `kyma-project.io` website, it opens the file content in the same tab.

## Links between documents in different topics

If you want to link to a document in the different topic, create a reference by using this pattern `/{type-of-topic}/{id}#{type}-{title}-{header}`, where:
- `{type-of-topic}` is a type of the topic that you want to reference. Apart from documents related directly to Kyma, all components have the `components` type of topic. For Kyma, use `root` in place of the type of topic.
- `{id}` is an ID of the topic that you want to reference. It is identical with the name of the component, for example: `helm-broker`, `kyma`.
- `{type}` is a metadata type of the document that you want to reference.
- `{title}` is a metadata title of the document that you want to reference.
- `{header}` is a header located in the document that you want to reference.

>**NOTE:** All variables must consist of lowercase characters separated with dashes (-).

If the `{type}` doesn't exist, the pattern has the form of `/{type-of-topic}/{id}#{title}-{title}-{header}`. If you want to create a reference to the whole `{type}`, the pattern has the form of `/{type-of-topic}/{id}#{type}-{type}`.

>**TIP:** You can copy the reference to the documentation directly from the website. See the cross-topic reference to the Helm Broker details document:

![Different topic reference](../../assets/reference-2.png)
