# Links

Linking is a great tool to use to incorporate a lot of content into your document with fewer words. That being said, overuse of linking can cause "link rot" when links break, and if a page has more links than content, it is not very pleasant to read. Make sure to use links correctly by adhering to these best practices:

- Use **relative links** to link to documents or files located in the same repository. For example:<br>
  ```Read the [contributing rules](../../contributing/02-contributing.md) before you contribute to the project.```

- Use **absolute links** to link to other repositories and external sources. For example:<br>
  ```To learn more about the project, visit the [kyma-poject.io](https://kyma-project.io/) website.```

- Every link has the potential to go bad over time and the more links you include, the higher the chance that one will break. If something is not central to the subject at hand, is well-known by your audience, or can be found with a simple search, there is no point in linking.
- Choose the link text carefully. Use descriptive text for the search engines to understand your content. Do not link entire phrases which become overemphatic. Avoid certain [vague words](https://web.dev/link-text/#how-the-lighthouse-link-text-audit-fails) like `this`, `that`, or `here`.

    Example:  
     ⛔️ "For more information, see **this** guide." or "Read more **here**."  
     ✅ "For more information, see the **installation guide**."  

## Links to Headings

Within a relative link, it's possible to link to the heading of a document. To do so, add `#{name-of-the-heading}` after the document's filepath. See the example:<br>
```Read the guidelines about [headings formatting](./03-formatting.md#headings).```

## Links to the Assets Folder

To add a reference to a YAML, JSON, SVG, PNG, or JPG file located in the `assets` folder in the same topic, use GitHub relative links. For example, write `[Here](./assets/mf-namespaced.yaml) you can find a sample micro frontend entity.` When you click such a link on the `kyma-project.io` website, it opens the file content in the same tab.

## Links in Documentation Toggles

To link to a document in a documentation toggle, the toggle must start with the `<div tabs name="{toggle-name}">` tag and end with the `</div>` tag, where **name** is a distinctive ID used for linking. For more information, read a separate [Tabs and toggles](./05-tabs-toggles.md) document.
