---
title: UI elements
---

Contrary to the regular technical documentation, UI elements follow different guidelines and best practices for content. Due to the limited space you can use to convey a message,

Here are the general guidelines you should follow when designing and writing UI elements:

- Use active voice is action-oriented language, starting with a verb. it is preferred for any UX text, and especially for tooltips since they need to be brief and empowering.
- Use title case for short texts, such as labels or headings, and sentence case for messages and explanations, unless specified otherwise.
- Use plural for the business object, as the user can generally process more than one object within the app.

  ✅ Service Instances
  ⛔️ Service Instance

- Use natural language. Some verbs may be similar in meaning but sound more or less natural to the reader. For example, from the pair "Choose" / "Select" or "Edit" / "Modify", the first verb sound more natural.
- Use the three-letter currency code instead of the currency symbol. This prevents ambiguity when multiple currencies share the same symbol.

  ✅ USD
  ⛔️ $

## Buttons

Be as straightforward and concise when designing the button label. To do so:

- Use action verbs . The button labels should tell the user what will happen without requiring him to read any supporting text. Most buttons contain verbs to indicate what the button will do, e.g. ‘Save’, ‘Publish’, ‘Edit’. While ‘Back’, and ‘Next’ aren’t verbs, in the context of an interface they seem to work in the same way. I like to keep the ‘verb’ + ‘noun’ structure when writing button labels — this makes the action more prescriptive, e.g. ‘Save post’, ‘Next step’, etc, as opposed to ‘Save’, ‘Next’ — but the choice is up to you and your users.
- Avoid labels that are too generic, such as "Yes" or "No". They are not intuitive and require the user to read the supporting text.  
- Use active imperative form to make your label as precise and task-specific as possible. When you use the imperative form in an active voice, you minimize the number of words and make the button label easier to scan.
- Omit articles and punctuation to make the button label as concise as possible. If possible, follow the verb+noun or
This form allows you to drop the subject and unnecessary articles for a more concise button label. All you need to include on your button label is a verb with an adverb or direct object.
- Use sentence case. Title case, e.g. ‘Next Section’ doesn’t read as easily as sentence case. It makes users feel like someone is speaking to them in a stilted voice. Sentence case, e.g. ‘Next section’ gives much more ‘friendly’ or ‘conversational’ tone. A sentence-style capitalization is best because it conveys a friendly tone that invites users to press the button. Users are most familiar with sentence case and it feels like someone is speaking to them in a natural voice.
- Use plain, natural language. Remember to write how you would speak to another human being. Have you ever used the word ‘modify’ with a friend? Maybe you have, but it won’t be in the everyday language of your users.
- Be careful when using words of similar meaning, such as "delete" and "remove".


## Tooltips

Tooltips are user-triggered messages that provide additional information about a page element or feature. In other words, if you have nothing to say that adds value to the user, don’t use a tooltip. It’s not only unnecessary, but annoying. In general, keep your tooltips to 1-2 sentences. If there’s more you need to say, you can add a ‘Learn more’ link leading to a Help article.

1. Don’t use tooltips for information that is vital to task completion.

Users shouldn’t need to find a tooltip in order to complete their task. Tooltips are best when they provide additional explanation for a form field unfamiliar to some users or reasoning for what may seem like an unusual request. Remember that tooltips disappear, so instructions or other directly actionable information, like field requirements, shouldn’t be in a tooltip. (If it is, people will have to commit it to their working memory in order to be able to act upon it.)

2. Provide brief and helpful content inside the tooltip.

Tooltips with obvious or redundant text are not beneficial to users. If you can’t think of particularly helpful content, don’t offer a tooltip. Otherwise, you’ll just add information pollution to your UI and waste the time of any users unlucky enough to activate that tooltip.

Provide tooltips for unlabeled icons.
Most icons have some level of ambiguity, which is why we recommend text labels for all icons. If you’re too stubborn to provide text labels for the icons on your site, the least you can do is provide your users with a descriptive tooltip.

7. Give examples

Nothing beats a good example to illustrate a point. Take advantage of the extra space provided by a tooltip to inspire users with real examples:

### Label tooltips

unline contextual tooltips, no dot
title case

### Contextual tooltips

6. Punctuate sentences

You may be tempted to drop the period for a one-line tooltip. However, a 2-3 sentence tooltip with no punctuation would obviously look strange. So to stay consistent and play it safe, punctuate.  




## Messages

Informational messages (aka passive notifications, something is ready to view)

If there are full sentences, let's use punctuation. Contrary wise, if the message is only a clause (fragment of a sentence), punctuation feels odd and redundant.

### Pop-up messages / hoverbox

### Drop-down messages

Gray out any unavailable options instead of removing them: any items that cannot b­­e selected should remain in view. For extra UX credit, consider showing a short balloon help message if users hover over a grayed-out option for more than a second, explaining why that option is disabled and how to make it active.
If disabled items are removed, the interface loses spatial consistency and becomes harder to learn.

Keep the menu label or description in view when the dropdown is open. Menu titles provide scope and direction by reminding users what they are choosing. Whenever the labels are obscured or removed when the menu is open, users must recall what they need to select before they can take action. Plan for interruptions that can disrupt the user’s task at any time.

### Error messages


Avoid the word “please,” except in situations in which the user is asked to do something inconvenient (such as waiting) or the software is to blame for the situation. ... Error messages need to clearly convey information to the user and if an error is serious, the tone should reflect that.

When it comes to writing error messages, clarity is your top priority. You need to describe what happened, why it happened, and what the user can do about it. The message should be written in plain language so that the target users can easily understand both the problem and the solution.

Avoid abstract error messages
Abstract error messages don’t contain enough information about the problem. In many cases, they simply state the fact that something went wrong and don’t help users understand the root cause of the problem. Don’t just assume people know about the context of a message—be explicit and indicate what exactly has gone wrong.

Get rid of technical terms
If an error message contains technical terms or jargon, the user gets confused. The error message should always describe the problem in terms of target user actions or goals. Even when your users are tech-savvy, it’s still better to use non-technical terms that everyone can easily understand.

Don’t try to explain a complicated troubleshooting process within an error message. Instead, use progressive disclosure to provide this information. The section that contains the steps should be hidden by default, and when the users want to learn more about the problem, they click “How to fix it.”

Avoid phrases like“You did,” “Your action caused.”
Some error messages are phrased in a way that accuses the user of making an error; errors are already frustrating, and there’s no need to add to frustration with judgment. In the end, these messages are an important, albeit, small way that we communicate and build relationships with our users. Always focus on the problem, not the user action that led to the problem.

Here are two ways you can handle a situation when the user enters incorrect login credentials:

Don’t say: You have entered an incorrect login or password.
Do say: Your login and password do not match.

4. Give users a solution
Imagine you wrote a very important email and clicked the “Send” button. Right after that you see the message, “Your email could not be sent,” without any details. As a result, you don’t know what you can do about it. You have to pause your task and invest your time in finding the solution to the problem.

8. Avoid Uppercase Text
Upper case text is difficult to read it gives an impact of shouting on user.
Error message is a place where user is informed about some critical scenario, so using upper case text can give him a feeling of discouragement.

## Labels

2. Without labels, users cannot check their work before submitting a form.
The lack of labels makes it impossible for customers to glance through the form and make sure that their responses are correct. Similarly, browsers that autocomplete form fields may fill in information incorrectly.

title case

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

### Reference

Follow these resources for further reference:
- [UI Text Guidelines for SAP Fiori Apps](https://experience.sap.com/internal/fiori-design-web/ui-text-guidelines-for-sap-fiori/)
- [5 rules for choosing button labels](https://uxmovement.medium.com/5-rules-for-choosing-the-right-words-on-button-labels-dc3f74c2c2a3)
- [Tooltip Guidelines](https://www.nngroup.com/articles/tooltip-guidelines/)
- [Tooltips: How to Craft Effective Guiding Text](https://www.wix.com/wordsmatter/blog/2020/06/tooltips/)
- [Placeholders in Form Fields Are Harmful](https://www.nngroup.com/articles/form-design-placeholders/)
- [Alternatives to Placeholder Text](https://medium.com/nextux/alternatives-to-placeholder-text-13f430abc56f)
- [How to Write and Design User-Friendly Error Messages](https://xd.adobe.com/ideas/process/information-architecture/error-message-design-ux/)
