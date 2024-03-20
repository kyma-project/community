# Formatting

Format the content in an attention-grabbing way. In general, content is easier to read when it is in chunks. Consider breaking up endless paragraphs by using a list or a table. Use action verbs and present tense for headings to engage the reader, and also follow the guidelines for the best way to include links and images. When you include lists, tables, code samples, or images, precede them with a brief explanation of what they describe.

## Code Font and Bold Font

It is important to consistently format items such as code or filenames to quickly distinguish them while reading technical documentation. The following tables outline when to use **bold** font and when to use `code` font:

### Items Written in Bold Font

Items       | Examples
----------- | ----------------------------------------------------------------------
Parameters   | The **env** attribute is optional.
HTTP headers | The Authorization Proxy validates the JWT token passed in the **Authorization Bearer** request header.
Events       | The service publishes an **order.created** event.
Roles        | Only the users with the **kyma_admin** role can list Pods in the Kyma system namespaces.
UI elements  | Click **Subscribe**.
Variables and placeholders | Click **Project** > **{YOUR_PROJECT_NAME}**.

### Items Written in Code Font

Items                     | Examples
------------------------- | -----------------------------------------------------------------------------------------------
Code examples             | Get the list of all Pods in a namespace using the `kubectl get pods -n {NAMESPACE}` command.
Values                    | Set the partial attribute to `true` to perform a partial replacement.
Endpoints                 | Send a POST request to the `/{tenant}/categories/{categoryId}/media/{mediaId}/commit` endpoint.
File names                | Open the `deployment.yaml` file.
File extensions           | Modify all `.yaml` and `.json` files.
Path names                | Save the file in the `\services\repository` folder.
Repository names          | The file is located in the `kyma` repository.
Status and error codes    | A successful response includes a status code of `200 OK`.
Parameter and value pairs | The controller adds the `env=true` label to all newly created namespaces.
Metadata names            | When you create a Markdown document, define its `title` and `type`.
Flags                     | Add the `--tls` flag to every Helm command you run.
GraphQL queries and mutations   | The `requestOneTimeTokenForApplication` mutation calls the internal GraphQL API to get a one-time token.
Rights                    | The preset gives multiple resource cleaners the `list` and `delete` rights in GCP.
Custom resources (if you want to refer to the code specifically) | Define the **namespace** parameter in the `APIRule` **metadata**.

>**NOTE:** When you mention specific configuration files in your documents, consider linking to them instead of just mentioning their names. When you link to a file, use its name without the format extension. See the following example:
> `To adjust the number of Pods in your Deployment, edit the [deployment](./deployment.yaml) file.`

### Omission in Code

Often when you quote snippets of code, it is not necessary to include them in full. For example, when you quote an HTTP response, it is enough to quote only the relevant part instead of enclosing the whole response. In such a case, replace the omitted parts with `...` to signalize that something has been removed.

For example, you might have an HTTP response as such:

   ```json
   {
     "csrUrl":"{CSR_URL_WITH TOKEN}",
     "api":{
       "eventsUrl":"https://gateway.name.cluster.extend.cx.cloud.sap/app/v1/events",
       "metadataUrl":"https://gateway.name.cluster.extend.cx.cloud.sap/app/v1/metadata/services",
       "infoUrl":"https://gateway.name.cluster.extend.cx.cloud.sap/v1/applications/management/info",
       "certificatesUrl":"https://connector-service.name.cluster.extend.cx.cloud.sap/v1/applications/certificates"
     },
     "certificate":{
       "subject":"O=Organization,OU=OrgUnit,L=Waldorf,ST=Waldorf,C=DE,CN=CNAME",
       "extensions":"",
       "key-algorithm":"rsa2048"
     }
   }
   ```

Suppose that it is only the part containing the certificate information that is of interest to the user. In that case, instead of quoting the whole response, quote it like this:

   ```json
   {
     ...
     "certificate":{
       "subject":"O=Organization,OU=OrgUnit,L=Waldorf,ST=Waldorf,C=DE,CN=CNAME",
       "extensions":"",
       "key-algorithm":"rsa2048"
     }
   }
   ```

## Panels

>**NOTE** The following panels are not rendered correctly when indented. Use them only in the docs that are displayed on the website.
For the docs not rendered on the website, use **NOTE**, **TIP**, and **CAUTION**.

Panels are colorful containers that call out important or additional information within a topic. The Kyma documentation uses the [Flexible Alerts](https://github.com/fzankl/docsify-plugin-flexible-alerts) docsify plugin for the panels' formatting. To call attention to a specific note, a word of caution or a tip, use the  [!NOTE], [!TIP], and [!WARNING] panels. Use:

- The blue [!NOTE] panel to point to something specific, usually relating to the topic.
- The orange [!WARNING] panel to call attention to something critical that can cause inoperable behavior.
- The green [!TIP] panel to share helpful advice, such as a shortcut to save time.

See an example:

> [!NOTE]
> Provision a Public IP for Ingress and a DNS record before you start the installation.

## Ordered and Unordered Lists

As you write about your topic, use lists to create visual clarity within your content. List items in a category or in a sequence. Use an ordered list for sequential, instructional steps. Unordered lists are appropriate for items that have no sequential order, such as a list of valid file types. Follow these guidelines:

* Make list content consistent in structure. For example, make all the bullet points sentences, questions, or sentence fragments, but do not mix types.
* Punctuate bullet points consistently. If they are sentences, use periods. If they are sentence fragments, do not use periods.
* Avoid ending bullet points with semicolons or commas.
* Capitalize the first letter of each bullet point consistently. Capitalize the first letter unless the list items are always lowercased, as with parameters names.
* Emphasize the beginning of the bullet point to capture the main idea.
* If readers must perform list items in order, as in a step-by-step procedure, use an ordered list and maintain consistency in structure.
* If you explain terms, bold them and provide their definitions either after a hyphen or in the sentence structure:  

    ✅ **ClusterServiceBroker** - an endpoint for ...

    ✅ **ClusterServiceBroker** is an endpoint for ...

    Use one type of formatting throughout the whole document. Do not write:

    ⛔️ **ClusterServiceBroker**, an endpoint for ...

## Tables

Another effective way to chunk content is to use tables. Use tables when content needs comparison, or as a way to provide information mapping. Think of a table as a list with optional columns useful to provide and organize more information than belongs in a list. Make sure tables are not too long or hard to read, causing the reader to scroll a lot. If possible, break up a long table into multiple tables.

When creating a table, centralize the columns that have choice-type values, such as `Yes/No` or `true/false`. If you include the **Default** column in your table, write `None` for parameters that have no default values. See the example:

   ```bash
   |     Parameter      |      Required      |  Description  | Default |
   |--------------------|:------------------:|---------------|---------|
   | {Parameter's name} |      {Yes/No}      | {Description} |   None  |
   ```

## Headings

Ideally, headings fit into one line in the generated output. Be concise, but also make sure to adequately describe the main point of the document or a section. Follow these guidelines when writing headings:

* Write headings in title case. For example, **Expose a Service**.
* Use action verbs and present tense verbs in headings when possible, especially in tutorials. For example, **Add a Document Type**.
* While gerunds are acceptable in body-level content, DO NOT use gerunds in headings. Use **Create a Storefront** instead of **Creating a Storefront**.
* Avoid stacked headings, which are headings without body-level content in between. For example, DO NOT use a Heading 2 (H2) to introduce one or more Heading 3s. Instead, add a paragraph after the H2 that describes the main idea of the content in the headings that follow.
* Do not use small headings, such as Heading 4 (H4) and smaller. Use Heading 1 (H1) for the document title, and Heading 2s (H2) and Heading 3s (H3) to organize the content of the document.
