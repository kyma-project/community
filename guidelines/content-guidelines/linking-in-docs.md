# Linking in docs

These are the guidelines for making cross-reference between the documents in the [kyma/docs](https://github.com/kyma-project/kyma/tree/master/docs) folder.

>**NOTE:** Remeber that linking between documents in [github.com/kyma-project/kyma/docs](https://github.com/kyma-project/kyma/tree/master/docs) doesn't work. At the moment it only work on [kyma-project.io/docs](https://kyma-project.io/docs).

## Linking between the same topic

If you want to linking to another document in the same topic, you must create reference by this pattern `#{type}-#{title}`, where:
- `{type}` is a type of the document, that you want to reference.
- `{title}` is a title from document that you want to reference.

`{type}` and `{title}` are a metadata of relevant document. They are at the beginning of each document.

>**NOTE:** All of vars in reference must be [tokenized](#tokenization).

>**NOTE 2:** If `{type}` doesn't exist, the pattern has a form `#{title}-#{title}`.

>**NOTE 3:** If you want create reference only to `{type}`, the pattern has a form `#{type}-#{type}`.

### Example

For the existing topic with files and theirs content:
- `001-overview-of-some-content.md`
  ```
  ---
  title: Some content
  type: Overview
  ---

  ...
  ```
- `002-details-of-some-content.md`
  ```
  ---
  title: Some content
  type: Details
  ---
    
  ...
  ```

a reference from `001-overview-of-some-content.md` to `002-details-of-some-content.md` must be: `#details-some-content`.

## Linking between different topics

If you want to linking to another document in the different topic, you must create reference by this pattern `/docs/{type-of-topic}/{id}#{type-of-document}-#{title}`, where:
- `{type-of-topic}` is a type of the topic, that you want to reference.
- `{id}` is a id from topic, that you want to reference.
- `{type-of-document}` is a type of the document, that you want to reference.
- `{title}` is a title from document that you want to reference.

`{type-of-topic}` and `{id}` are a metadata of relevant topic. They are in `docs.config.json` file, which is in the root directory of each topic.

`{type-of-document}` and `{title}` are a metadata of relevant document. They are at the beginning of each document.

>**NOTE:** All of vars in reference must be [tokenized](#tokenization).

>**NOTE 2:** If `{type-of-document}` doesn't exist, the pattern has a form `/docs/{type-of-topic}/{id}#{title}-#{title}`.

>**NOTE 3:** If you want create reference only to `{type-of-document}`, the pattern has a form `/docs/{type-of-topic}/{id}#{type}-#{type}`.

### Example

For the following architecture with files and theirs content:
- `service-catalog/docs/001-overview-of-some-content.md`
  ```
  ---
  title: Some content
  type: Overview
  ---

  ...
  ```
  where `service-catalog` has a:
  - `type`: `service-catalog`.
  - `id`: `components`.
- `service-brokers/docs/001-overview-of-some-content.md`
  ```
  ---
  title: Some content
  type: Overview
  ---
    
  ...
  ```
  where `service-brokers` has a:
  - `type`: `service-brokers`.
  - `id`: `components`.

a reference from `service-catalog/...` to `service-brokers/...` must be: `/docs/service-brokers#overview-some-content`.

## Tokenization

Tokenization is a operation of changing data to lowercasing form with a dash sign (`-`) instead of space.

### Example

- Before tokenization:  `Document with some content`.
- After tokenization:   `document-with-some-content`.
