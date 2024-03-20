# UI Elements

Contrary to the regular technical documentation, UI elements follow different guidelines and best practices. Due to the limited space you can use to convey a message, your UI text must be as precise and concise as possible. Use the following general guidelines to design and write content for UI elements:

- Use active voice and action-oriented language, for example:

  ✅ Show active contacts

  ⛔️ This toggle shows active contacts

  ✅ Change the password

  ⛔️ The password must be changed

- Use title case for short one-line texts, such as labels or headings, and sentence case for messages and explanations, unless specified otherwise.

  ✅ Postal Code

  ⛔️ Postal code

  ✅ Use sentence case for messages and explanations.

  ⛔️ Don't Use Title Case For Longer Messages.

- Use plural for the object names, as users usually interact with multiple objects within the program, for example:

  ✅ Service Instances

  ⛔️ Service Instance

- Use plain, natural language. Even though some verbs may be similar in meaning, they may sound more or less natural to the reader. For example, from these pairs of verbs, the first one sounds more natural:

  ✅ Choose

  ⛔️ Select

  ✅ Edit

  ⛔️ Modify

- Use the three-letter currency code instead of the currency symbol. This prevents ambiguity when multiple currencies share the same symbol.

  ✅ USD

  ⛔️ $

- Avoid ampersands (&), technical and mathematical symbols, and symbols that may have multiple meanings. Commonly used symbols are acceptable.

  ✅ % CPU

  ⛔️ Due in < 5 days

  ⛔️ Save & exit

## Labels

Labels are one-line texts that appear in various types of UI elements, such as buttons, titles and headings, forms, and drop-down menus. When designing labels for the UI elements in the Kyma project, make sure to:

- Use title case for titles and headings.
- Use sentence case in drop-down menus and button labels.
- Omit punctuation to avoid cluttering your text.

### Buttons

When designing a button label, you must be as straightforward and concise as possible, since you have very limited space to convey the message. Moreover, buttons are used to invite a user to perform an action. Make them as inviting and intuitive as possible. To do so:

- Use action verbs that explain what will happen when a user presses the button, without requiring them to read any supporting text. Depending on your needs, you can also use plain adverb. Alternatively, add a noun to indicate what a verb or an adverb refers to. Do **not** use articles and punctuation. They only clutter your button label.

  ✅ Save

  ✅ Next

  ✅ Save draft

  ✅ Next step

  ⛔️ Save a draft

- Avoid labels that are too generic, such as `Yes` or `No`. They are not intuitive and require the user to read the supporting text.  
- Use sentence case as it reads more natural and friendly to the users, which makes them more willing to press the button. Title case doesn't read as easily as sentence case and looks more unnatural.

  ✅ Next section

  ⛔️ Next Section

- Be careful when using words of similar meaning, such as `Delete` and `Remove`.

## Tooltips

Tooltips are messages that provide additional information about a certain UI element. Still, the text within a tooltip must contain information that brings value to the user, so avoid adding content that may be irrelevant. Tooltips cannot contain information that is necessary to complete a given task. Such information should be visible at first glimpse. Otherwise, users will have to remember the content of a tooltip and refer back to it after the tooltip is deactivated. Here are some general guidelines on creating tooltip texts:

- Keep your tooltips 1-2 sentences long. If you need to explain a concept in more details, create a separate document and link to it instead of providing the whole explanation in a tooltip.
- Use sentence case inside a tooltip and follow punctuation rules.
- Use tooltips to provide examples:

  ✅ The example of a domain name is `yourdomain.example.com`.

- Omit punctuation in case of one-line label tooltips. For more information, see [Labels](#labels).

## Messages

There are many different types of messages appearing in every UI. The most common ones are pop-ups and regular messages that users can find on your page. To avoid information pollution, make sure that all your messages are informative and useful to the users. Moreover:

- Use punctuation in case of messages that are full sentences. In case of clauses (fragments of a sentence), omit punctuation.
- Use title case for the title of your pop-up message. Keep it as simple as possible and omit punctuation. For more information, see [Labels](#labels).

### Error Messages

An error message informs about unsuccessful outcome of an action or conveys any other negative information, so you must choose your words wisely not to upset the user. Here are some tips that will help you design your error message:  

- Describe what happened, why it happened, and what the user can do to fix the issue. Avoid abstract messages that only inform that something went wrong. This could make the user feel frustrated not only about the fact that the error has occurred, but also for the fact they cannot do anything about it. However, do not describe the whole troubleshooting in an error message. Link to a separate troubleshooting guide instead.
- Avoid technical terms and jargon. Every user should be able to understand the message, regardless their technical knowledge. Use language as plain as possible.  
- Don't blame the user for making an error. Focus on the problem instead, for example:

  ✅ Your password is incorrect.
  
  ⛔️ You have entered an incorrect password.

- Avoid title case. It can give users a feeling they are being looking down on.

## Placeholder Texts

A placeholder is a tricky UI element. It tends to disappear when the user clicks a given form field, so it requires them to use their memory, which increases memory load and hurts usability. For this reason:

- Use placeholders in combination with labels that do not disappear and provide all the necessary information visible at all times. Don't place information essential to complete a task in a placeholder. Also, do not repeat the same information in a label and a placeholder. In such a case, give up on using a placeholder.
- Use placeholders to give an additional hint, description, or example.
- Use sentence case for your placeholder.
- Omit full stop at the end of the placeholder text.

## Reference

Follow these resources for further reference:

- [UI Text Guidelines for SAP Fiori Apps](https://experience.sap.com/internal/fiori-design-web/ui-text-guidelines-for-sap-fiori/)
- [5 Rules for Choosing the Right Words on Button Labels](https://uxmovement.medium.com/5-rules-for-choosing-the-right-words-on-button-labels-dc3f74c2c2a3)
- [Tooltip Guidelines](https://www.nngroup.com/articles/tooltip-guidelines/)
- [Tooltips: How to Craft Effective Guiding Text](https://www.wix.com/wordsmatter/blog/2020/06/tooltips/)
- [Placeholders in Form Fields Are Harmful](https://www.nngroup.com/articles/form-design-placeholders/)
- [Alternatives to Placeholder Text](https://coyleandrew.medium.com/alternatives-to-placeholder-text-13f430abc56f)
- [How to Write and Design User-Friendly Error Messages](https://xd.adobe.com/ideas/process/information-architecture/error-message-design-ux/)
