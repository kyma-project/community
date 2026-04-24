# SAP/Kyma Technical Writing Style Rules

Review documentation in sequential passes. Each pass focuses on one concern — complete it fully before moving to the next. After all passes, run a reconciliation sweep.

---

## Pass 1 — Formatting and Mechanics

### Capitalization

**Title case** for: document titles, topic titles, section headings, table captions, column headings, product names, software object titles, UI element labels/titles.

Title case rules:
- Always capitalize: first and last word, nouns, gerunds, verbs (including "be" forms), participles, adjectives, adverbs (including "than", "not"), determiners/pronouns ("This", "That", "All", "Any", "Which", "You"), prepositions of 5+ letters ("About", "Between", "Through"), all subordinating conjunctions ("Because", "If", "When"), prepositions used as adverbs in verb phrases ("In" in "Signing In", "Up" in "Setting Up").
- Keep lowercase: articles ("the", "a", "an"), coordinating conjunctions ("and", "but", "or"), prepositions under 5 letters when not part of a verb phrase ("to", "for", "in", "on", "at", "from", "with").
- With punctuation: capitalize hyphenated words as separate words ("Self-Service Material"), capitalize after colons as if first word, capitalize inside parentheses as if no parentheses.

**Sentence style** for: body text, messages, tooltips that read as clauses/sentences, non-header table cells.

**Never** use all-caps text. Do not use capitalization for emphasis.

### Headings

- **Procedural topics**: gerund formulation in title case. Example: "Setting Up Authorizations".
- **Conceptual topics**: nominal formulation in title case. Example: "Authorization Concept".
- Do not use "How to..." formulations — they are redundant.
- Avoid abbreviations in headings; use full forms.
- No additional formatting (bold, italic) in headings. No ending periods.
- No stacked headings — always add body content between consecutive headings.
- Use H1 for document title, H2 and H3 for content organization. Avoid H4 and smaller.
- Keep titles short but informative with relevant keywords.

### Lists

- Introduce every list with a complete sentence ending in a colon.
- Ordered lists for sequential steps; unordered for non-sequential items.
- All items must use parallel grammatical structure.
- Capitalize the first word of every item.
- No trailing punctuation on incomplete-sentence items. Complete sentences end with a period.
- If items mix complete and incomplete sentences, revise for parallelism.
- Minimum 2 items; restructure lists with 8+ items by grouping.
- Maximum 2 nesting levels; restructure if deeper.
- Do not state the number of items in the introductory sentence.
- Label optional steps with "Optional:". Conditional steps: use an if-clause or state the condition explicitly.
- When items pair a key concept with a description, consider converting to a table.
- Avoid embedding tables within lists (simple 2-column tables are acceptable).
- Bold the term and use a hyphen or sentence structure for term definitions. Be consistent.

### Tables

- Introduce every table with a complete sentence ending in a colon.
- Maximum 5 columns. Always include a header row with meaningful titles.
- Use proportional column widths.
- Parallel grammatical structure within each column.
- Title case for captions and header cells; sentence style for data cells.
- Do not state the number of rows or columns.
- Captions: if used, apply consistently across all tables. No "table" in the caption. Nominal formulation. Position above. Do not number.
- Keep cell content concise. Move common words to header cells. Combine long technical strings with descriptions in the same cell.
- Column headings and cell text must be independently readable — do not form a sentence across them.
- Leave cells empty when no information applies (or write "Not applicable").
- Do not use "X" marks — use "Yes", "No", or descriptive words.
- If a table includes a **Default** column, write `None` for parameters with no default.
- Center-align columns containing choice-type values (Yes/No, true/false).
- Avoid overly long tables; break into multiple shorter tables when possible.

### Panels and Admonitions

**Valid panel types:** Note, Tip, Recommendation, Caution, Remember. The Restriction type may be used only when absolutely certain it does not describe a software limitation affecting revenue recognition.

**Correct syntax** (Help Portal blockquote format):

```markdown
> ### Note:
> Content here.
```

**Invalid syntax**: `[!WARNING]`, `[!NOTE]`, `[!TIP]` (GitHub/VitePress format) — flag these if encountered.

**Usage rules:**
- Maximum 2 advisory items per topic. Do not place them consecutively.
- Consider inline alternatives ("Note that...") for minor advisories.
- If the same advisory appears in many topics, centralize it and link.

**Panel type semantics:**
- **Note**: important or non-obvious information; no action involved.
- **Tip**: optional advice to avoid minor issues; no risk of harm.
- **Recommendation**: advantageous settings, procedures, or methods. See Pass 2 for recommendation language rules.
- **Caution**: severe hazards — data loss, system failure, file damage.
- **Remember**: fundamental information needed later, or summary of complex content.

### Code and Bold Formatting

**Bold** for: UI element labels, parameters, HTTP headers, events, roles, variables/placeholders.

**Italic** (or `uicontrol` tag): UI element labels when referring to them in procedural text. Reproduce the label exactly as displayed.

**Code font** for:

| Item | Example |
|------|---------|
| Code examples | `kubectl get pods -n {NAMESPACE}` |
| Values | `true`, `false` |
| Endpoints | `/{tenant}/categories/...` |
| File names | `deployment.yaml` |
| File extensions | `.yaml`, `.json` |
| Path names | `\services\repository` |
| Repository names | `kyma` |
| Status/error codes | `200 OK` |
| Parameter-value pairs | `env=true` |
| Metadata names | `title`, `type` |
| CLI flags | `--tls` |
| GraphQL queries/mutations | `requestOneTimeTokenForApplication` |
| Rights | `list`, `delete` |

**Code omission**: when quoting partial code snippets, replace omitted parts with `...`.

### Kubernetes Resource Names

In body text: plain CamelCase — no bold, no code font. Example: "Create a ConfigMap with the required data."

In inline code spans (e.g., showing a `kubectl` command): code font applies from the span itself.

In titles and navigation: add spaces for natural language. Example: "Config Map" in a heading.

For plurals: add lowercase "s". Example: ConfigMaps, Deployments.

"custom resource" (generic concept) is lowercase. CamelCase only for specific resource type names.

**Kubernetes resource reference list**: ConfigMap, CronJob, CustomResourceDefinition (CRD), Deployment, Function, Ingress, Node (capitalize only for the K8s resource — lowercase for VMs or billing units), PodPreset, Pod, ProwJob, Secret, Service, ServiceBinding, ServiceClass, ServiceInstance.

Always capitalize "Kubernetes". Never use "k8s".

"namespace" is always lowercase.

### Punctuation

**Serial comma**: always use before the final "and"/"or" in a series of three or more items.

**Periods**: end complete sentences with a period. Do not add a period after standalone incomplete sentences (e.g., success confirmations: "Purchase order created"). Do not enclose periods inside formatting of the preceding word. No double periods after abbreviations.

**Colons**: use to introduce lists, tables, graphics, or explanatory information. Capitalize after a colon if followed by a complete sentence; lowercase if followed by a phrase. Do not enclose colons in the formatting of the preceding word.

**Semicolons**: use sparingly. Prefer separate sentences. Acceptable in complex series with internal commas.

**Commas**: serial comma required. Comma before coordinating conjunctions introducing independent clauses. Set off nonrestrictive relative clauses. Comma after introductory adverbial clauses, long prepositional phrases (5+ words), and sentence adverbs. Never place a comma between subject and verb.

**En dashes**: use instead of em dashes for sentence-level interruptions, with spaces on both sides. No spaces in number ranges (e.g., 1–3). Do not use em dashes.

**Hyphens**: follow U.S. conventions.
- Always hyphenated prefixes: all-, cross-, ex-, full-, high-, self-, well- (before noun only).
- Not hyphenated: anti, auto, bio, co, counter, extra, multi, non, post, pre, over, re, semi, sub, super, under.
- Hyphenate phrasal adjectives before a noun ("time-consuming exercise") but not in predicative position ("the exercise is time consuming") or when modified by an adverb ("highly configurable").

**Quotation marks**: punctuation goes inside only for complete quoted sentences. For individual words/phrases, punctuation goes outside.

**Apostrophes**: never use for plurals. "BAPIs" not "BAPI's". "cannot" as one word, never "can not".

**Slashes**: do not use as shorthand for "or" or "and". Spell out the conjunction. Exception: established compounds ("plan/actual"), approved names ("SAP S/4HANA"), URLs/paths. For slash-separated labels with spaces: "ZIP Code / City".

**Ellipsis**: no space before. Used for busy dialog text ("Loading...") and menu items leading to further options. Do not use in input hints.

**Parentheses**: use for explanatory/supplementary information. Period inside if parenthetical is a standalone sentence after another sentence. Period outside if parenthetical is a phrase within a sentence.

### Numbers, Dates, Units

**Numbers**: spell out one through nine in running text; use numerals for 10+. Use numerals regardless of size for: series of numbers in a sentence, figures/ages/times/dates/years, units, dimensions, currencies, percentages, "N or more"/"up to N" formulations. Spell out numbers in fixed phrases ("third party", "first name").

**Number ranges**: en dash without spaces (0–9). Use "from...to" not "from...–". For ambiguous ranges, use "to" instead of en dash.

**Decimal separator**: period in English. Include leading zero for fractions less than 1 (0.5).

**Thousands separator**: comma for 4+ digit numbers (7,654,321). Use words for large round numbers when clearer ("10 million"). For space-restricted: K/M/B per SAP Fiori guidelines.

**Negative numbers**: minus sign directly before the number, no space (USD -100).

**Dates**: Month DD, YYYY format. Write months as words, years as 4 digits. No all-numeric formats. Comma after the day and after the year mid-sentence.

**Times**: H:MM a.m./p.m. with periods. Include time zone abbreviation in parentheses.

**Percentages**: no space between numeral and % sign (5%).

**Units**: metric (SI). No plural "s" on abbreviations (25 km not 25 kms). Spell out unit names when no quantity is given. Repeat units in ranges (10 mm x 17 mm). Add nonmetric equivalents in parentheses for English when relevant.

**Currencies**: ISO 3-letter codes (EUR, USD) before the amount. Spell out currency names without an amount ("pay in euros"). Omit trailing .00 for whole amounts unless another amount in the same context has decimals.

### File References

- Code font for specific file names (`NewsletterTemplate.docx`). Uppercase extension with descriptor for file types ("ZIP file", "PDF").
- Follow OS conventions for case and path separators (/ for Unix, \ for Windows).
- For long paths: complete introductory sentence ending in colon, then path on next line.
- When mentioning configuration files, link to them. Use the file name without extension in link text.

### Topic Structure

**Task topics**: structure with an introductory paragraph (why), prerequisites (if needed), numbered steps, expected result. Aim for 5–9 steps; split longer tasks.

**Concept topics**: answer "what-is" questions, provide background. Must not contain instructions or reference tables/lists. Link to related task or reference topics.

**Troubleshooting topics**: title is the symptom or error message, not the cause. Three sections: **Condition**, **Cause**, **Solution**. Numbered list for multi-step solutions; bullet list or sub-headings for alternatives.

**Custom resource topics**: CamelCase resource name as title. Use `{RESOURCE_NAME}.md` filename convention.

**General structure**: one paragraph per idea. Big picture first — title and opening paragraph must convey relevance. Balance plain text with block elements. Use links to avoid redundancy.

### Cross-References and Hyperlinks

- Make link text match or closely match the target title. Never use generic "click here" or "read more".
- Maximum 7 cross-references per topic.
- Keep links out of descriptive sentences — use standard formulations: "For more information, see *Target*."
- Do not adapt link text to fit sentence grammar.
- Prefer end-of-topic "Related Information" section for cross-references.
- Relative links for same-repository documents; absolute links for external sources.
- Link to headings with `#{heading-name}` appended to the file path.
- Link sparingly — do not link well-known or easily searchable items.

### Screenshots and Diagrams

**Screenshots**:
- Use only when text alone cannot convey the information. Screenshots are costly to maintain.
- Complement text, do not replace instructions.
- No directional references ("above", "below") — use introductory text.
- No mouse pointer unless demonstrating a function.
- Always provide alt text.
- Preferred format: SVG, then PNG/JPG. Compress; max 1 MB per image.
- Size: 100%, 75%, 50%, or 25% of screen width. Max 860px.
- Border: grey (#D2D5D9) 1pt.
- Step markers: blue (#0A6ED1) round stamps with white numbers; explain with ordered list.
- Indicators: red (#EF2727) 10pt arrows/boxes, max one per screenshot.
- Simplified User Interfaces (SUI): blur/cover non-essential elements. Light gray (#F2F2F2) for text, dark gray (#D9D9D9) for headlines. Do not cover: logo, sandwich icon, search icon, expand/close buttons.

**Diagrams**:
- Use to visualize complex concepts, workflows, architectures. Not decorative.
- Always precede with context text. Always provide alt text.
- Format: SVG via draw.io. Max width 860px.
- Background: white (not transparent). Rounded secondary backgrounds: mild blue (#F0F6FF) for main environment, mint green (#DEF2DD) for subsidiary.
- Shapes: rounded rectangles, white fill. Actors: blue (#0A6EC7) fill.
- Outlines: grey (#666666) 1pt for shapes. No outlines for actors, steps, or backgrounds.
- Text: black, Helvetica. 15pt bold headings, 13pt primary, 12pt secondary. Always horizontal. Titles centered in shapes.
- Step markers: blue (#0A6EC7) stamps with white numbers.
- Connectors: 1pt rounded grey (#666666) lines.
- Left-to-right workflow direction.
- Include a reference key when introducing elements that differ from others in the diagram.

### Release Notes

- Write from user perspective: what changed, how behavior differed before, UI/functionality impact.
- Headlines: short, interesting, sentence case.
- New features: enticing paragraph, not just a bulleted list.
- Use "known issue" and "resolved issue" — never "bug".
- Consistent sentence structure within bulleted lists.
- Provide a migration guide when manual upgrade steps are required; do not describe new features in it.

---

## Pass 2 — Language and Grammar

### Voice and Mood

**Active voice** at least 80% of the time. Address the user with "you" or use imperatives. Name the agent in active-voice sentences involving third parties ("the system", "the vendor").

Passive voice acceptable when:
- The agent is unknown, unimportant, or omitted for politeness.
- Context makes the agent clear and passive improves flow.

**Imperative mood** for mandatory instructions ("Save your changes."). Do not use imperatives for advisory/recommendation content — see Recommendation Language below.

### Recommendation Language

When content is intentionally advisory (optional, best-practice, suggested), preserve the advisory tone. Do not convert recommendations to bare imperatives — this changes meaning from "suggested" to "required".

Approved recommendation patterns (no first person):
- "It is recommended to {action}." (acceptable passive — recognized exception to active-voice preference)
- "Consider {action}." / "As a best practice, {action}."
- For formal advisories, use the **Recommendation** panel type.

Flag and rewrite:
- "We recommend..." → use an approved pattern above.
- "It is not recommended..." → rephrase positively: state what to do instead.

The imperative mood rule and the "should → must/can" terminology rule do not apply in recommendation contexts.

### Tense, Length, Ordering

**Present tense** for generic/habitual statements. Reserve future tense for genuinely future actions.

**Sentence length**: ~20 words average. Split longer sentences or convert to lists/tables.

**Ordering**:
- State what before how (purpose before mechanism).
- Describe actions in chronological order.
- Place conditions ("if" clause) before the result.

### Positive and Concise Formulations

- Rephrase double negatives into positive statements.
- Use positive formulations in error messages: tell users what to do, not what is invalid.
- Eliminate wordy formulations: "perform an upgrade" → "upgrade", "be able to" → "can".
- Remove intensifiers (very, basically, really, absolutely) unless essential.
- Use precise verbs; do not nominalize ("performs the upgrade" → "upgrades").
- Split noun stacks with prepositions.

### Global English

- Use common, everyday words (see word substitution table in Pass 3).
- Do not omit articles, conjunctions, or prepositions for conciseness.
- Include "that" after verbs like "make sure", "indicate", "verify" to introduce object clauses.
- Include "that" in restrictive relative clauses (do not omit even when grammatically optional).
- Disambiguate coordinating conjunctions — split into separate sentences if ambiguous.
- Keep phrasal verbs together; do not separate with long objects.
- Prefer relative clauses over post-nominal participle constructions for clarity.
- Avoid causative "have" and "get".

### Relative Clauses

- "that" for restrictive (defining) clauses — no commas.
- "which" for nonrestrictive (nondefining) clauses — with commas.
- "who" for persons.
- Commas signal the difference in meaning. Verify comma usage matches intent.

### People-Centric Language

- Use "you" to address the reader directly. "Your" instead of "the" where natural.
- Present features from the user's perspective (what the user can do, not what the feature does).
- No first person ("we", "us", "let's") in documentation. Exception: recommendation patterns are handled via approved alternatives above.
- Third person acceptable when explaining different roles.
- Sound natural — read content aloud. No exaggeration, no chatty padding.
- Use conversational transitions where appropriate ("Here's how", "What's next?").
- Exclamation points: use sparingly. They can be perceived as shouting.

### Tone

- Positive, encouraging, helpful, natural-sounding voice. Follow the CIAO principle: Clear, Insightful, Approachable, Optimistic.
- Match tone to target group: conversational for end users, direct/expert for IT consultants.
- No cliches, buzzwords, colloquialisms, jargon.
- Error messages must match severity — no cheerful tone for serious failures.
- Be polite but direct. No unnecessary "please" in routine instructions. Reserve "please" for inconvenient requests or error recovery.
- "Sorry" only when software failed through no user fault. Do not overuse.
- Do not blame the user. In error messages, use passive voice to remove blame.
- Do not blame the software excessively.

### Contractions

- Common contractions (don't, can't, it's, you're) are acceptable for natural tone.
- Expand in warnings or serious messages ("Do not" instead of "Don't").
- Never use unusual or compound contractions ("couldn't've", "the user'll").
- "cannot" as one word, always.

### Bias-Free and Inclusive Language

- No stereotypes or discrimination. Be sensitive to culture, gender, age, ethnicity.
- Avoid culture-dependent metaphors and idioms.
- Use culturally neutral terms ("holiday bonus" not "Christmas bonus").
- Replace non-inclusive terminology ("blocklist" not "blacklist"). "scrum master", "master data", "white paper" are acceptable.
- Gender-inclusive: "chairperson", "workforce", "person days".
- Singular "they" when gender is unknown. Never "he/she". Prefer "you" or imperatives to avoid the issue entirely.
- No emoticons or emojis.

### Accessibility

- Do not refer to people by disabilities. Use people-first language ("people with disabilities").
- Focus on needs, not disabilities ("if you use a screen reader" not "if you are blind").
- Never state software "is accessible" or "is not accessible". Use "usability improvements".
- Use "choose" as device-neutral verb instead of "click" for multidevice applications.
- Provide text alternatives for non-text content (icons, graphics, video, audio).
- Describe what before how in step-by-step instructions (helps screen reader users).
- Structure content with proper headings. Introduce lists/tables/graphics with introductory sentences.

### UI Text Rules

These apply when writing or reviewing UI microcopy (labels, buttons, tooltips, messages, placeholders).

**General**: embed task-critical information on the UI, not in separate guides. Do not state the obvious. No all-caps.

**Labels and buttons**: title case for short labels. Use action verbs on buttons; no articles or punctuation. Avoid generic "Yes"/"No" — use specific actions ("Delete", "Leave Page"). Distinguish "Delete" vs. "Remove" accurately.

**Tooltips**: 5–15 words. Only when label is insufficient. Do not copy label text. Mandatory for icons. Include keyboard shortcuts in parentheses.

**Messages**: friendly, supportive tone. No generic messages ("An error occurred"). No exclamation marks. Positive formulations. Provide actionable error messages (what happened + how to fix). Complete sentences end with a period; incomplete do not.
- Success: `<Object> <past participle>` — no "successfully".
- Error: provide diagnosis + solution.
- Confirm: "Do you want to...?" with self-descriptive action buttons.

**Input hints**: only when no self-descriptive label exists. Verbal formulation. No ellipsis. No ending period. Do not duplicate label text.

**Referring to UI elements**: prefer indirect device-neutral formulations. When click-level detail is needed, use "choose" (not "click"/"tap"). Format labels distinctly (italic or `uicontrol`). Reproduce labels exactly. Tell what before how.

**UI element prepositions**: "in the *X* field", "on the *X* screen/tab", "in the *X* dialog box", "select the *X* checkbox", "choose *X*" (buttons), "choose *X* > *Y*" (navigation path with > separator).

Avoid click-level detail in Kyma documentation — readers are typically tech-savvy. Reserve step-by-step UI instructions for genuinely complex interactions.

### Keyboard Keys

- Do not mention keyboard keys unless necessary (accessibility, time-saving shortcuts).
- Describe the action first, then the key combination.
- Simultaneous keys: plus sign with spaces (Ctrl + Alt + Del).
- Sequential keys: comma separator.

---

## Pass 3 — Terminology and Naming

### Product Naming: SAP BTP, Kyma Runtime

**Full name**: "SAP BTP, Kyma runtime". Always include "runtime" after "Kyma". Lowercase "runtime" except in title-case contexts.

Flag these as incorrect:
- "SAP BTP Kyma" (missing "runtime")
- "SAP Kyma" (missing "BTP")
- "SAP BTP, KR" (no abbreviation)

Permitted short forms (space-limited contexts only, when context is clear):
- "SAP BTP Kyma runtime" (without comma)
- "Kyma runtime"

On first use, expand "BTP" to "Business Technology Platform" to clarify the abbreviation.

In body text, natural phrasing is acceptable: "the Kyma runtime for SAP BTP."

For cockpit/catalog tiles: "Kyma Runtime" (title case, short form).

### Product Naming: Kyma Dashboard

**Name**: "Kyma dashboard". Lowercase "dashboard" except in title-case contexts.

For clarity when needed: "Kyma dashboard for SAP BTP, Kyma runtime." Subsequent mentions: "the dashboard."

Flag as incorrect:
- "SAP Kyma dashboard"
- "SAP BTP Kyma dashboard"

### SAP Product Names (General)

- Use exact approved names. No abbreviations, plurals, or possessives of product names.
- No trademark symbols in documentation.
- No italics or quotation marks for product names.
- At first mention, include the approved descriptor ("the SAP Business One solution"). After first mention, use the name alone without article.
- Use only approved abbreviations. Introduce at first occurrence: full name (abbreviation) descriptor.
- Do not create false product-name impressions with the "SAP" prefix.
- Avoid compound nouns with approved names — rephrase with prepositions ("settings in SAP Extended Warehouse Management" not "SAP Extended Warehouse Management settings").
- In title-case contexts, capitalize approved names in title case even if they use lowercase in body text.

### SAP Terminology Conventions

- Use consistent terminology — one word, one meaning. Do not use synonyms for defined terms.
- Use product-neutral verbs consistently in procedural titles (e.g., pick one of "create"/"configure"/"set up" and use it consistently for the same action).
- Replace "country" with "country/region" or "country or region" (geopolitical compliance).
- Replace the SAP-internal verb "maintain" with the precise action: "edit", "specify", "create", "change", or "manage".
- Use precise terms instead of generic "system" for software.

### U.S. English

U.S. English is the corporate standard. Key suffix conventions:

| U.S. | Not (British) |
|------|---------------|
| -ense (defense) | -ence |
| -er (center) | -re |
| -ize (organize) | -ise |
| -og (catalog) | -ogue |
| -or (color) | -our |

"canceled" not "cancelled"; "modeling" not "modelling"; "fulfill" not "fulfil"; "toward" not "towards".

U.S. vocabulary: "check" (not "cheque"), "parentheses" (not "brackets"), "exclamation point" (not "exclamation mark"), "period" (not "full stop"), "vacation" (not "holiday/leave").

U.S. prepositions: "fill out" (not "fill in"), "on the weekend" (not "at the weekend"), "because" (not "as" for causal meaning).

Reference works: *Merriam-Webster's Collegiate Dictionary* and *The Chicago Manual of Style*. SAP style guide takes priority on conflicts.

### Kyma-Specific Terminology

Use these preferred terms in Kyma documentation:

| Use | Don't use |
|-----|-----------|
| ID | id |
| backend | back end, back-end |
| frontend | front end, front-end |
| key-value | key/value, key:value |
| micro frontend | microfrontend, micro front-end |
| email | e-mail |
| repository | repo |
| document | doc |
| must | have to, need to, should (when mandatory) |
| can | should (when optional), it is possible to, allows you to |
| need | require |
| run | execute |
| use | utilize |
| using, with | via |
| typically | usually |
| fill in | complete |
| connect, connection | integrate, integration |
| YAML | yaml (use `.yaml` for file extensions) |
| for example, such as | e.g. |
| that is | i.e. |
| the following, several | a specific number |
| the previous, earlier | above, this |
| cloud-native (adjective) | cloud native |
| Infrastructure Provider, IaaS Provider | hyperscaler, Cloud Provider |
| application (lowercase for microservices; capitalized "Application" for Application Connector entities) | app |
| Prow Job (process) | Prowjob, prowjob (use CamelCase "ProwJob" for the K8s resource) |

### Preferred Simple Words

| Don't | Do |
|---|---|
| concur | agree |
| albeit | although |
| prior to | before |
| commence | start, begin |
| endeavor | try |
| heretofore | up to now |
| utilize | use |
| obtain | get |
| occur | happen |
| perform an upgrade | upgrade |
| provide an explanation | explain |
| be able to | can |
| is characterized by transparency | is transparent |

Remove redundant modifiers: "final outcome" → "outcome"; "absolutely perfect" → "perfect".

### Abbreviations

- Minimize abbreviations. Use only if more common than the full form (ID, PC, API) or industry-standard.
- At first occurrence: full form (abbreviation). Thereafter: abbreviation only. Do not alternate.
- No Latin abbreviations in body text: "for example" not "e.g.", "that is" not "i.e.", "and so on" not "etc." Exception: space-restricted UI labels.
- "and" not "&" or "+" unless part of an approved name.
- Plural of acronyms: add lowercase "s", no apostrophe (APIs, IDs).
- Indefinite article by pronunciation: "an SAP solution" (ess-ay-pee), "a BAPI".
- Do not create ad-hoc CamelCase abbreviations of normal words.
- No spaces in abbreviated forms with periods (a.m. not a. m.).
- Use short CLI arguments when possible (`-n` not `--namespace`). Explain when short arguments differ between tools.

### Foreign-Language and Country/Region Terms

- Use established English equivalents when they exist. Use the foreign term if no equivalent.
- Do not italicize foreign terms.
- Explain foreign terms at first use, including the full original form if an acronym.
- Use language-specific geographic names in English when available ("Vienna" not "Wien").

### Software Objects and Entities

- Distinguish real-life entities from software objects. Real-life terms use lowercase; software objects get title case + formatting only at first mention.
- First occurrence: *Title in Title Case* (`technical_name`) descriptor.
- Subsequent mentions: title (formatted) or descriptor with demonstrative pronoun.
- Do not capitalize everyday words just because they correspond to software concepts ("business process" not "Business Process" in body text).
- Software entity tools/engines: title case without special formatting ("the BAdI Builder").

### Versioning

- Avoid unnecessary version numbers. Include only when distinctive.
- Pattern: `<Approved Name> <version> <SPS/FPS number>`. No leading zero in body text.
- "and higher" or "up to" for version ranges.
- Acronym with number (SPS 9); full form standalone ("support packages").

### Edge-Case Fallback

When reviewing documentation that presents a situation not explicitly covered by any rule above, apply general best-practice judgment and note that no specific rule was matched.

---

## Pass 4 — Reconciliation Sweep

After completing Passes 1–3, re-read all changes and verify:

1. **Internal consistency**: changes from one pass have not introduced violations caught by another pass. Example: a terminology fix in Pass 3 may have broken a formatting rule from Pass 1.
2. **Cross-reference integrity**: links, headings, and cross-references still resolve correctly after edits.
3. **Tone coherence**: the document maintains a consistent voice throughout. No section sounds markedly different in formality or style.
4. **Recommendation language**: verify that advisory content retains its advisory tone after edits. No recommendations were converted to bare imperatives.
5. **Product naming**: verify all mentions of SAP BTP, Kyma runtime and Kyma dashboard follow the naming rules. Check for common errors: "SAP BTP Kyma" without "runtime", "SAP Kyma dashboard".
6. **Panel syntax**: verify all admonitions use the blockquote format, not GitHub/VitePress syntax.
7. **Kubernetes resource formatting**: verify CamelCase in body text (no bold, no code font) and proper capitalization.
