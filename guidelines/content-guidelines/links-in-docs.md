# Links in docs

These are the guidelines for making cross-references between the documents in the [`kyma/docs`](https://github.com/kyma-project/kyma/tree/master/docs) folder.

>**NOTE:** The linking works only on the [Kyma website](https://kyma-project.io/docs). Currently, the cross-references between [GitHub documentation](ttps://github.com/kyma-project/kyma/tree/master/docs) is not available.

## Links between documents in the same topic

If you want to link to another document in the same topic, create a reference using the `#{type}-{title}-{header}` pattern, where:
- `{type}` is a type of the document that you want to reference.
- `{title}` is a title of the document that you want to reference.
- `{header}` is a header located in the document that you want to reference.

>**NOTE:** All variables must consist of lowercase characters separated with dashes (-).

`{type}` and `{title}` are placed in a metadata section of each document. If the `{type}` doesn't exist, the pattern has the form of `#{title}-{title}-{header}`. If you don't want to create a reference to the `{header}`, the pattern has the form of `#{type}-{title}`. If you want to create a reference to the whole `{type}`, the pattern has the form of `#{type}-{type}-{header}`.

For example, there are two documents named `001-overview-service-brokers.md` and `002-details-azure-broker.md`. Their metadata fields look as follows:
- `001-overview-service-brokers.md`
  ```
  ---
  title: Service Brokers
  type: Overview
  ---
  ```
- `002-details-azure-broker.md`
  ```
  ---
  title: Azure Broker
  type: Details
  ---
  ```

In this case, a reference from `001-overview-service-brokers.md` to `002-details-azure-broker.md` is  `#details-azure-broker`.

### Links to the assets folder

To add a reference to a YAML, JSON, SVG, PNG, or JPG file located in the `assets` folder in the same topic, use GitHub relative links. For example, write `[Here](./assets/mf-namespaced.yaml) you can find a sample micro front-end entity.` When you click such a link on the `kyma-project.io` website, it opens the file content in the same tab.

## Links between documents in different topics

If you want to link to a document in the different topic, create a reference by using this pattern `/{type-of-topic}/{id}#{type-of-document}-{title}-{header}`, where:
- `{type-of-topic}` is a type of the topic that you want to reference.
- `{id}` is an ID of the topic that you want to reference.
- `{type-of-document}` is a type of the document that you want to reference.
- `{title}` is a title of the document that you want to reference.
- `{header}` is a header located in the document that you want to reference.

>**NOTE:** All variables must consist of lowercase characters separated with dashes (-).

`{type-of-topic}` and `{id}` are metadata fields of the given topic. They are placed in the `docs.config.json` file, in the root directory of each topic. `{type-of-document}` and `{title}` are placed in a metadata section of each document. If the `{type-of-document}` doesn't exist, the pattern has the form of `/{type-of-topic}/{id}#{title}-{title}-{header}`. If you don't want to create a reference to the `{header}`, the pattern has the form of `/{type-of-topic}/{id}#{type-of-document}-{title}`. If you want to create a reference to the whole `{type-of-document}`, the pattern has the form of `/{type-of-topic}/{id}#{type-of-document}-{type-of-document}-{header}`.

For example, there are two documents with the following metadata sections:
- `service-catalog/docs/001-overview-service-catalog.md`
  ```
  ---
  title: Service Catalog
  type: Overview
  ---
  ```
- `service-brokers/docs/001-overview-service-brokers.md`
  ```
  ---
  title: Service Brokers
  type: Overview
  ---
  ```

The `{id}` variables of the documents are `service-catalog` and `service-brokers` respectively, and their `{type-of-topic}` is `components`.

In this case, a reference from `service-catalog/...` to `service-brokers/...` is `/service-brokers#overview-service-brokers`.
