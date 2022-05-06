---
title: UI elements
---

Contrary to the regular technical documentation, UI elements follow different guidelines and best practices. Due to the limited space you can use to convey a message, your UI text must be as precise and concise as possible. Here are some general guidelines you should follow when designing and writing content for UI elements:

- Use active voice and action-oriented language, starting with a verb.
- Use title case for short texts, such as labels or headings, and sentence case for messages and explanations, unless specified otherwise.
- Use plural for the object names as users usually interact with multiple objects within the program:

  ✅ Service Instances <br>
  ⛔️ Service Instance

- Use plain, natural language. Even though some verbs may be similar in meaning, they may sound more or less natural to the reader. For example, from these pairs of verbs, the first one sound more natural:

  ✅ Choose <br>
  ⛔️ Select

  ✅ Edit <br>
  ⛔️ Modify

- Use the three-letter currency code instead of the currency symbol. This prevents ambiguity when multiple currencies share the same symbol.

  ✅ USD <br>
  ⛔️ $

## Buttons

When designing a button label, you must be as straightforward and concise as possible, since you have very limited space to convey the message. Moreover, buttons are used to invite a user to perform an action. Make them as inviting and intuitive as possible. To do so:

- Use action verbs that explain what will happen when a user presses the button, without requiring them to read any supporting text. Depending on your needs, you can also use plain adverb. Alternatively, add a noun to indicate what a verb or an adverb refers to. Do **not** use articles and punctuation. They will only add unnecessary "word pollution" to your button label.

  ✅ Save <br>
  ✅ Next <br>
  ✅ Save draft <br>
  ✅ Next step <br>
  ⛔️ Save a draft

- Avoid labels that are too generic, such as `Yes` or `No`. They are not intuitive and require the user to read the supporting text.  
- Use sentence case as it reads more natural and friendly to the users, which makes them more willing to press the button. Title case doesn't read as easily as sentence case and looks more unnatural.

  ✅ Next section <br>
  ⛔️ Next Section

- Be careful when using words of similar meaning, such as `Delete` and `Remove`.


## Tooltips

Tooltips are messages that provide additional information about a certain UI element. Still, the text within a tooltip should contain information that brings value to the user, so avoid adding content that may be irrelevant. Tooltips shouldn't also contain information that is necessary to complete a given task. Such information should be visible at first glimpse. Otherwise, users will have to remember the content of a tooltip and refer back to it after the tooltip is deactivated. Here are some general guidelines on creating tooltip texts:

- Keep your tooltips 1-2 sentences long. If you need to explain a concept in more details, create a separate document and link to it instead of providing the whole explanation in a tooltip.
- Use sentence case inside a tooltip and follow punctuation rules.
- Use tooltips to provide examples, for example:

  ✅ The example of a domain name is `yourdomain.example.com`.

- Omit punctuation in case of one-line label tooltips (for more information, see [Headings and labels](#headings-and-labels)).


## Messages

There are many different types of messages appearing in every UI. The most common ones are pop-ups and regular messages that  users can find on your page. To avoid information pollution, make sure that all your messages are informative and useful to the users. Moreover:

- Use punctuation in case of messages that are full sentences. In case of clauses (fragments of a sentence), omit punctuation.
- Use title case for the title of your pop-up message. Keep it as simple as possible and omit punctuation. (?)

### Error messages

Error message is a special type of message you must be really careful about. It informs about unsuccessful outcome of an action or conveys any other negative information, so you must choose your words wisely not to upset the user. Here are some tips that will help you design your error message:  

- Describe what happened, why it happened, and what the user can do to fix the occurred issue. Avoid abstract messages that only inform that something went wrong. This could make the user feel frustrated not only about the fact that the error has occurred, but also for the fact they cannot do anything about it. However, do not describe the whole troubleshooting in an error message. Link to a separate troubleshooting guide instead.
- Avoid technical terms and jargon. Every user should be able to understand the message, regardless their technical knowledge. Use language as plain as possible.  
- Don't blame the user for making an error. Focus on the problem instead, for example:

  ✅ Your password is incorrect. <br>
  ⛔️ You have entered an incorrect password.

- Avoid title case. It can give users a feeling they are being looking down on. (what about labels/titles? to clarify)

## Headings and labels

Labels are one-line texts that appear in various types of UI elements, such as titles and headings, forms, and drop-down menus.

- Use title case for your labels. (?)
- Omit punctuation to avoid word pollution. 


## Placeholder texts

A common implementation is by inserting instructions within form fields. Unfortunately, user testing continually shows that placeholders in form fields often hurt usability more than help it.

Labels tell users what information belongs in a given form field and are usually positioned outside the form field. Placeholder text, located inside a form field, is an additional hint, description, or example of the information required for a particular field. These hints typically disappear when the user types in the field.

If the user forgets the hint, which people often do while filling out long forms, he has to delete what he wrote and, in some cases, click away from the field to reveal the placeholder text again.

Using placeholder text in combination with form labels is a step in the right direction. Labels outside the form fields make the essential information visible at all times, while placeholder text inside form fields is reserved for supplementary information. However, even when using labels, placing important hints or instructions within a form field can still cause the 7 issues mentioned above, albeit with less severity. If some of the fields require an extra description that is essential to completing the form correctly, it’s best to place that text outside the field so that it is always visible.

The default light-grey color of placeholder text has poor color contrast against most backgrounds. For users with a visual impairment, poor color contrast makes it difficult to read the text. Because not all browsers allow placeholder text to be styled using CSS, this is a difficult issue to mitigate.
Users with cognitive or motor impairments are more heavily burdened.  As we saw, placeholders can be problematic for all users: disappearing placeholders increase the memory load; persistent dimmable placeholders cause confusion when they look clickable but aren’t, and placeholders that do not disappear require more keyboard or mouse interaction to be deleted. These difficulties are magnified for people with cognitive or motor impairments.
Not all screen readers read placeholder text aloud. Blind or visually impaired users may miss the hint completely if their software does not speak the placeholder content.

never use placeholder as a label!

In an attempt to shorten the length of a form or reduce visual noise, designers use placeholder text as an input label. This practice places a burden on short-term memory. The label disappears as soon as the user clicks and/or types. The entry must be deleted to expose the label again.

 Providing an example of the needed input helps a user understand the request. However, incorporating the example as placeholder text causes issues including disappearance on focus, confusion regarding what has been entered, and reduction of the input acting as an affordance. As an alternative, example text can be placed below the input field.

 -Placeholders should be of a lighter value than input text
- Placeholders should be visible on all screens
- Placeholders should not disappear when a user clicks into the input
-

### dropdown menu placeholders

### Drop-down messages

Gray out any unavailable options instead of removing them: any items that cannot b­­e selected should remain in view. For extra UX credit, consider showing a short balloon help message if users hover over a grayed-out option for more than a second, explaining why that option is disabled and how to make it active.
If disabled items are removed, the interface loses spatial consistency and becomes harder to learn.

Keep the menu label or description in view when the dropdown is open. Menu titles provide scope and direction by reminding users what they are choosing. Whenever the labels are obscured or removed when the menu is open, users must recall what they need to select before they can take action. Plan for interruptions that can disrupt the user’s task at any time.

### Reference

Follow these resources for further reference:
- [UI Text Guidelines for SAP Fiori Apps](https://experience.sap.com/internal/fiori-design-web/ui-text-guidelines-for-sap-fiori/)
- [5 rules for choosing button labels](https://uxmovement.medium.com/5-rules-for-choosing-the-right-words-on-button-labels-dc3f74c2c2a3)
- [Tooltip Guidelines](https://www.nngroup.com/articles/tooltip-guidelines/)
- [Tooltips: How to Craft Effective Guiding Text](https://www.wix.com/wordsmatter/blog/2020/06/tooltips/)
- [Placeholders in Form Fields Are Harmful](https://www.nngroup.com/articles/form-design-placeholders/)
- [Alternatives to Placeholder Text](https://medium.com/nextux/alternatives-to-placeholder-text-13f430abc56f)
- [How to Write and Design User-Friendly Error Messages](https://xd.adobe.com/ideas/process/information-architecture/error-message-design-ux/)
