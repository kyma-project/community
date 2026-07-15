# Kyma Technical Writing Style Guide

Consolidated ruleset for writing and reviewing SAP/Kyma technical documentation for the SAP Help Portal. Review documentation in **four sequential passes**, then apply the document-structure rules. Each pass is self-contained; focus on one concern at a time.

When documentation presents a situation no rule below covers, apply general technical-writing best practice and note that no specific rule matched.

---

## Pass 1 — Formatting and mechanics

How text is formatted, punctuated, and capitalized on the page — including structural line-rules for headings, lists, tables, and panels.

### Capitalization

- **Body text:** sentence style — lowercase except the first word and proper nouns/names.
- **Titles, headings, captions, column/row headings, product names, UI-element and software-object titles:** title case.
- **Never all caps** (in topics or UI); do not capitalize for emphasis.
- Do not title-case everyday words in body text even when they name a UI object, dropdown, or concept. Rule of thumb: if there can be more than one of it, it is not a proper noun. Do (body): `business process`, `situation handling`. Proper nouns keep caps: `World Wide Web Consortium`.

**Title case mechanics:**
- Capitalize first and last word regardless of part of speech.
- Capitalize nouns, gerunds, verbs (incl. all forms of "be"), participles, adjectives, adverbs (incl. "Not", "Than").
- Capitalize demonstratives (This/That/These/Those), quantifiers (All/Any/Some/Every), interrogatives/relatives (What/Which/Who/Whose/Where), pronouns (You/They/Its/Them). Articles the/a/an stay lowercase.
- Prepositions: capitalize if 5+ letters (About, Between, Through, Without); lowercase if <5 letters — unless the preposition is a verb particle, which is always capitalized. Ex: `Signing In to the System` (In = particle), `Sold-to Party` (to ≠ particle).
- Subordinating conjunctions (Because, If, That, When) capitalized; coordinating (and, but, or) lowercase.
- Hyphenated words: apply part-of-speech rules to each element (`Self-Service`, `Add-In`). Words in parentheses and after a colon: capitalize as if not enclosed / as a first word.

### Kubernetes resource-name formatting (override)

- Kubernetes resource names (built-in and custom) in body text: **plain CamelCase — no bold, no code font.** Ex: `API Gateway operates on APIRule custom resources.`
- Add "s" for plurals (`ConfigMaps`). In titles/navigation add spaces: `Config Map`, `Config Maps`.
- Code font applies only when inside a fenced block or an inline code span you are writing anyway (e.g. a `kubectl` command, or referring to the resource *in the code* specifically: `APIRule`). Code font is not a Kubernetes-specific formatting rule.
- Capitalize a resource word only in the Kubernetes sense. `Node` (K8s) vs. a VM node (lowercase). `CustomResourceDefinition (CRD)` capitalized; "custom resource" alone is lowercase.
- Cross-ref: for *what to call* resources and the CamelCase reference list, see Pass 3.

### Bold and code font (Kyma taxonomy)

Use **bold** for: parameters, HTTP headers, events, roles, UI elements, variables/placeholders. Ex: `Click **Subscribe**.`, `the **kyma_admin** role`, `**{YOUR_PROJECT_NAME}**`.

Use `code` font for: code examples, values, endpoints, file names, file extensions, path names, repository names, status/error codes, parameter/value pairs, metadata names, flags, GraphQL queries/mutations, rights, and custom resources when referring to the code specifically. Ex: `` `deployment.yaml` ``, `` `200 OK` ``, `` `--tls` ``, `` `env=true` ``.

- Set off UI element labels/titles from body text so screen readers detect them. Ex: `In the **Delivery Date** field, enter a future date.`
- Do not enclose trailing punctuation (period, colon, comma) inside formatting — translation tools segment on them. Do: `the following **options**:` Don't: `the following **options:**`. Exception: a stable label segment like `**Note:**`.
- Do not format approved product names or Kubernetes resource names (see Pass 3 / K8s rule above).

### Panels / admonitions (override — Help Portal blockquote syntax)

Correct syntax:
```markdown
> ### Note:
> Content here.
```
Valid types and when to use:
- **Note** — important/unusual info to understand; no action.
- **Tip** — optional action to solve/avoid a minor issue; no risk of harm.
- **Recommendation** — advantageous or proven settings/procedures/methods.
- **Caution** — avoid severe hazards (data loss/inconsistency, system failure); not easily repaired.
- **Remember** — fundamental info the user needs later, or a summary of a complex thread.
- **Restriction** — use with extreme caution and only when certain it does not describe a software limitation affecting revenue recognition.

Flag as invalid: VitePress/GitHub syntax `[!NOTE]`, `[!WARNING]`, `[!TIP]` — not valid for the Help Portal. Do not indent panels.
Cross-ref: advisory-info moderation (max ~2 per topic) is in Document structure.

### Headings and titles (formulation)

- H1 = document title; use H2/H3 to organize. No H4 or smaller.
- Title case (see Capitalization). No periods in titles; use colons, parentheses, question marks sparingly; avoid exclamation points. No extra formatting in titles (they are already set off) except technical names where technically possible.
- **Kyma heading formulation:** use action verbs and present tense; **do not use gerunds in headings** (gerunds are fine in body). Do: `Expose a Service`. Don't: `Exposing a Service`. (This is the Kyma convention and governs Kyma headings/task titles; it overrides the SAP gerund default.)
- No stacked headings — put an introductory paragraph between a heading and the next heading.
- Balance length: short as possible, informative as necessary; avoid opaque abbreviations and transaction/technical minutiae.
- Cross-ref: topic-type title conventions (concept/task/troubleshooting/CR) are in Document structure and Pass 3.

### Lists

- Introduce every list with a complete sentence ending in a colon; do not embed items in the sentence. Do not state the item count ("the following", "as follows").
- Ordered list if sequence matters; unordered if not.
- Parallel formulations; do not mix noun/verb phrases with complete sentences. Keep all items the same type.
- Punctuation: periods if items are complete sentences; none if fragments. Never end items with semicolons or commas.
- Capitalize the first word of each item (unless items are always-lowercase tokens like parameter names).
- At least 2 items (exception: generated/fill-in lists). Restructure lists of 8+ items on one level. Max 2 nesting levels (3rd rare; 4+ → restructure).
- Put the important information first in each item and set it off (e.g. bold), not buried after repeated "You can".
- Term definitions in a list: bold the term, then define after a hyphen **or** within the sentence — pick one style per document. No comma-introduced appositive. Do: `**ClusterServiceBroker** - an endpoint for ...` or `**ClusterServiceBroker** is an endpoint for ...`. Don't: `**ClusterServiceBroker**, an endpoint for ...`.
- Set off optional steps with `Optional:`; conditional steps with an if-clause or condition label.

### Tables

- Introduce with a complete sentence ending in a colon. Don't state row/column counts.
- Max ~5 columns; break up long tables. Header row with meaningful titles; proportional widths.
- Parallel formulations within a column. Header cells/captions: title case, minimal punctuation. Cells: sentence style, standard punctuation.
- Center-align choice-type columns (`Yes/No`, `true/false`).
- If a **Default** column exists, write `None` for parameters with no default.
- Leave cells empty when there is no info (or "Not applicable" if absence could look like a mistake). Do not use X marks — use words like "Yes".
- Captions (if used): use for all tables, above the table, nominal, no word "table", no numbering.

### Numbers, dates, times, units, currency

- Spell out one–nine; numerals for 10+. Use numerals below 10 for series of many numbers, or for figures/ages/times/dates/years/measurements/dimensions/currencies/pages/percentages/phone numbers, or in "N or more" / "up to N" patterns. Spell out in fixed phrases (`third party`, `first name`).
- Number ranges: en dash, no spaces (`0–9`); use "from ... to ..." / "between ... and ..." for parallel forms; use "to" when a range start is ambiguous.
- Decimal separator: English point, with leading zero for <1 (`0.5`). Thousand separator: English comma for 4+ digits (`25,500`); may omit for a lone low number with no unit (`1500 persons`). Prefer words for large numbers (`10 million`).
- Metric units; language-neutral abbreviation; no plural "s" on unit abbreviations. Repeat the unit in ranges/combos (`10 mm × 17 mm`). Spell out a unit used without a quantity. Percent: no space in English (`5%`). Hyphenate numeral + spelled-out unit (`18-mile`), not abbreviated (`2 kg weight`).
- Currency: ISO 3-letter code, code before amount in English (`EUR 100`), never symbols. Omit trailing `.00` when no fractional digits.
- Dates: months as words, 4-digit years, `Month DD, YYYY` in English; never all-numeric (`08/05/2022`). Make times a.m./p.m. explicit; time zones abbreviated in parentheses. ISO 8601 `YYYY-MM-DD hh:mm:ss` is for machine data only, not body text/UI.

### Punctuation and characters

- Serial (Oxford) comma before the final item in a series of 3+. Cross-ref: Pass 2 for clause/conjunction usage.
- Colon introduces lists/tables/info; capitalize after a colon only if a complete sentence follows.
- Semicolons: prefer separate sentences; use only to join closely related main clauses or to separate complex comma-containing series items.
- En dashes (not em dashes), spaced, used sparingly (no spaces in number ranges). Never em dashes.
- One space between words and after punctuation. No spaces in abbreviations with periods (`a.m.`). Do not insert line-break spaces into paths/URLs/technical names.
- Quotation marks: language-specific; enclose ending punctuation only for a complete quoted sentence. Don't use quotes for titles/emphasis/irony; use formatting instead.
- Slash: not shorthand for "or" (see also Pass 3 terminology); acceptable in unit modifiers (`plan/actual`) and approved names (`SAP S/4HANA`). Refer to a special character as `name (symbol)`, e.g. `ampersand (&)`.
- No `(s)` or `/s` plural forms — rephrase (plural only, "one or more", "each", or disconnect the number from the noun).
- Cross-ref: runnable shell-command and code-block formatting conventions are in Document structure.

---

## Pass 2 — Language and grammar

Word choice, grammar, voice, mood, tense, sentence quality, US English.

### Voice, mood, tense

- **Active voice** by default. Do: `The endpoint path includes your service name.` Don't: `Your service name is to be included in the endpoint path.` (Exception: apologies may use passive to avoid blame; and the advisory passive `It is recommended to ...` — see Recommendations.)
- **Imperative mood** for instructional/procedural content; other moods can imply the action is optional. Do not frame required actions as optional ("if you want to ... you can"). Do: `Click **Upload** and select one or more documents.`
- **Present tense** — easier than past/future; prefer simple over progressive/perfect. Do: `an error message appears`. Don't: `will appear`. Watch overuse of "will"/"you'll"/"won't".
- If there is only one way to do something, drop "can" and use the imperative.

### Addressing the reader

- Second person: address the reader as "you"/"your". Use third person ("the user") only for a different role.
- **No first person** ("we", "us", "let's"). Cross-ref: recommendation phrasing below and in Pass 3.
- People-centric: describe what the reader can do, not what the feature is. Do: `Find people by searching for their name or email.` Don't: `The search tool can search by name and email.`
- Positive phrasing; do not blame the user. Do: `Enter alphanumeric characters only.` / `Your password is incorrect.` Don't: `Don't enter special characters.` / `You entered an incorrect password.`

### Recommendations (override — preserve advisory tone, no first person)

When content is advisory (optional/best-practice), keep it advisory — do not convert to a bare imperative (that changes "suggested" to "required"). The should→must/can rule and the imperative-mood rule do **not** apply here.
- Approved patterns: `It is recommended to {action}.` (accepted passive exception), `Consider {action}.`, `As a best practice, {action}.`, or a **Recommendation** panel.
- Flag & rewrite: `We recommend ...` → use an approved pattern. `It is not recommended ...` → state positively what to do instead.

### Grammar and clarity

- Articles: "the" for a specific thing, "a"/"an" for non-specific; choose a/an by pronunciation (`an SAP solution`). Include articles to remove ambiguity.
- Use "that" to introduce object clauses and in restrictive relative clauses even when omittable. Do: `Define the information that you require.`
- Restrictive clauses: "that", no commas. Nonrestrictive: "which"/"who", with commas.
- Keep phrasal verbs together (`Fill out the form.`). Prefer a relative clause over a post-noun participle (`the template that was used`, not `the template used`).
- Don't use causative "have"/"get" ("get a page created") — name the actor. Avoid ambiguous "since" (because vs. from-the-time). Avoid dangling modifiers. Clarify what "and"/"or" join. Split infinitives only for contrast.
- Collective nouns take a singular verb (`The data is saved`, `SAP is launching`).
- Serial comma (see Pass 1). Use a colon to introduce a list, a semicolon between independent clauses — but prefer neither. Avoid parentheses; prefer a list or restructured sentence.
- Contractions: acceptable in moderation; write out in full for serious warnings (`do not`, `cannot` — never `can not`); no unusual/compound contractions; watch its/it's.

### US English

- American spelling: `color`, `behavior`, `organize`, `center`, `catalog`, `defense`, `license`, `fulfillment`, `program`. Not British (`colour`, `organise`, `centre`, `-ogue`, `-our`).
- American terms over British: `fill out` (not fill in), `for more information about` (not on), `parentheses`/`curly brackets`/`square brackets`, `period`, `exclamation point`, `vacation`, `expiration date`.
- Hyphenate American-style (less than British): hyphen mainly to prevent misreading. Hyphenate an adjective **before** a noun only (`well-known actor` vs. `is well known`); no hyphen when an adverb modifies it (`highly configurable`). Prefixes usually closed (`autogenerated`, `nonnegotiable`, `preassembled`) except all-/cross-/self-/high-/well-/ex-/full-.

### Global English — common words and clear structure

Prefer common, everyday words; avoid telegraphic style (keep conjunctions, articles, prepositions). Sentences must sound natural and be unambiguous.

**Plain-word substitutions (use → avoid):**

| Use | Avoid |
| --- | --- |
| agree | concur |
| although | albeit |
| before | prior to |
| start, begin | commence |
| try | endeavor |
| up to now | heretofore |
| use | utilize |
| get | obtain |
| happen | occur |
| harmful | deleterious |
| can | be able to |
| upgrade | perform an upgrade |
| explain | provide an explanation |
| is transparent | is characterized by transparency |
| that is | i.e. |
| for example, such as | e.g. |
| and so on | etc. |
| note | n.b. |
| and others | et al. |
| versus | vs. |
| compare | cf. |

- Do not use Latin abbreviations in body text (space-restricted UI may keep `e.g.`). Cross-ref: preferred Kyma terms are in Pass 3.

---

## Pass 3 — Terminology and naming

What to call things — product names, Kyma terminology, resource naming, abbreviations.

### Product naming (SAP BTP, Kyma runtime)

- Full name: **SAP BTP, Kyma runtime**. Always keep "runtime" after "Kyma"; lowercase "runtime" except where title case applies. Flag `SAP BTP Kyma` without "runtime".
- Do **not** use: `SAP BTP Kyma`, `SAP Kyma`, `SAP BTP, KR`, or any invented abbreviation.
- Space-limited web/menu/tile contexts only: `SAP BTP Kyma runtime` or `Kyma runtime` (comma/`SAP BTP` may drop when context is clear). Cockpit/service-catalog tiles use the short name `Kyma Runtime`.
- Body text may use natural phrasing like `the Kyma runtime for SAP BTP`. Spell out "Business Technology Platform" early in a document to gloss the abbreviation.
- Dashboard: **Kyma dashboard** (lowercase "dashboard" except in titles). Do not precede with `SAP`/`SAP BTP` (`SAP Kyma dashboard` is wrong). First mention may use `Kyma dashboard for SAP BTP, Kyma runtime`; later just `dashboard`.
- General SAP offering names: use the approved name exactly — no plurals (rephrase: `SAP CRM systems`), no possessives (`SAP HANA benefits`), no trademark symbols, no italics/quotes, no invented abbreviations. Title case in titles even if the body form has lowercase words.
- Descriptors/articles for SAP offerings: at first occurrence use the approved descriptor (`the SAP Business One solution`); afterward use the name stand-alone with no "the" and no descriptor, or a descriptor with "this".

### No article before Kyma product/component names

Do not use an article before the product or a component name. Do: `Kyma is awesome.`, `Application Connector`. Don't: `The Kyma`, `the Application Connector`. (Note: this Kyma rule coexists with the SAP first-occurrence-descriptor rule for SAP *offering* names above; for Kyma component names, no article.)

### Kubernetes / resource naming reference

- Always capitalize **Kubernetes**; never `kubernetes` or `k8s`.
- CamelCase for Kubernetes and custom resources; "s" for plurals; spaces in titles/navigation. Format as code (`` `APIRule` ``) only when referring to the code specifically. Cross-ref: body-text formatting rule in Pass 1.
- Do **not** capitalize `namespace`.
- Kubernetes resource names (CamelCase): `ConfigMap`, `CronJob`, `CustomResourceDefinition (CRD)`, `Deployment`, `Function`, `Ingress`, `Node`, `Pod`, `PodPreset`, `ProwJob`, `Secret`, `Service`, `ServiceBinding`, `ServiceClass`, `ServiceInstance`.
- Title case for Kyma component names: `Application Connector`, `API Gateway Controller`.

### Kyma preferred terminology (use → avoid)

| Use | Avoid | Note |
| --- | --- | --- |
| ID | id | |
| and | + , & | |
| backend | back end, back-end | |
| frontend | front end, front-end | |
| micro frontend | microfrontend, micro front-end | |
| connect, connection | integrate, integration | |
| custom resource | Custom Resource, CustomResource | |
| document | doc | |
| email | e-mail | |
| fill in | complete | |
| key-value | key/value, key:value | |
| must | have to, need to | |
| need | require | |
| or | / (slash) | |
| repository | repo | |
| run | execute | |
| single sign on (SSO) | single sign-on | |
| the following, several | a specific number | |
| typically | usually | |
| use | utilize | |
| using, with | via | |
| YAML | yaml | file extension/name → `.yaml` |
| Prow Job (process) | Prowjob, prowjob | resource → `ProwJob` |
| must / can | should | mandatory → must; optional → can |
| (you) can | it is possible to, allows you to, you have the option to | if no other option, drop "can", use imperative |
| that is | i.e. | often the lead-in adds little; keep only the clause after |
| for example, such as | e.g. | mid-sentence, prefer "such as" |
| cloud-native (adj.) | cloud native | adjective compounds usually hyphenated |
| application | app | "Application" = external solution via Application Connector; "application" = a microservice on Kyma |
| the following | below, this, the described | |
| the previous, earlier | above, this, the described | |
| Infrastructure Provider, IaaS Provider | hyperscaler, Cloud Provider | |

- Do not use transitional-phase words (`currently`, `now`) or make forward-looking promises; document only what already exists.
- Short command-line arguments where possible (`-n`, not `--namespace`); explain context if ambiguous. Note: `kubectl -n` = `--namespace`, but `helm -n` = `--name`, not `--namespace`.

### Abbreviations

- Prefer no abbreviation; introduce one only if more common than the full form or an industry standard, and never coin on the fly. At first occurrence give `full form (ABBR)`, then use the abbreviation; don't alternate. No abbreviations in titles.
- Plurals without apostrophe (`IDs`, `BAPIs`). No CamelCase in abbreviations of translatable words (CamelCase is reserved for technical names). Use "and", not `&`/`+`.

### Country/region

Do not use "country" alone — use `country/region` or `country or region`, or rephrase. Do: `Select the relevant countries/regions.`

### Software objects and entities

- Real-life entity or where "this is software" is irrelevant: write like any other word — no title case, no formatting. Do: `A strategic purchaser creates a contract.` Don't: `As a Strategic Purchaser ...`.
- Software object (role, API, function module, etc.) at first occurrence: `*Title in Title Case* (\`technical name\`) descriptor`. Later, use the title alone or the technical name. Omit elements that don't apply.
- Software entity behaving as a proper noun (tool/engine, one exists, no plural): title case, no formatting, use an article, do not prefix with `SAP`. Ex: `The BAdI Builder`.
- AI: make it obvious users interact with AI; use the AI's name, avoid pronouns where possible.
- Cross-ref: topic-type naming (concept/task/troubleshooting/CR titles) is in Document structure.

---

## Pass 4 — Reconciliation sweep

Re-read the changes from Passes 1–3 together and verify the document as a whole:

- Verify no pass introduced a new violation of another pass (e.g. a terminology fix in Pass 3 that broke sentence flow, or a formatting change that de-capitalized a proper noun).
- Check consistency across the full document: one term per meaning; consistent capitalization per text type; consistent list/table punctuation; consistent term-definition style; heading style uniform.
- Verify the document conforms to its **topic type** (concept, task, troubleshooting, custom resource) per Document structure below — correct title style, required sections, and shape.
- Confirm advisory panels are within moderation limits and use valid Help Portal syntax.
- Where a situation matched no rule, confirm the best-practice judgment applied is consistent with the rest of the document.

---

## Document structure and conventions

Whole-file / topic-shape rules. Not a numbered pass. Each rule is cross-referenced from its owning pass.

### Topic types (cross-ref: Pass 3 for titling, Pass 1 for heading formulation)

Kyma is topic-based: one file per topic, self-contained, linking out as needed.

- **Concept** (answers "what is"): nominal title (`Security`, `Security Concept`). Background only — no instructions, no reference tables/lists; may link to task/reference topics.
- **Task** (a task the reader accomplishes): imperative-verb title naming the task, not the feature, not a gerund, not "How to...". Do: `Define resource consumption`. Don't: `Selecting a profile`, `How to select...`. Structure: intro paragraph ("why do this?") → prerequisites (if any) → numbered steps → expected result. Aim for 5–9 steps; split longer.
- **Troubleshooting:** title names the symptom/error the reader sees, not the cause. Do: `Cannot access ...`. Don't: `Incompatible version`. Structure: three headings — **Condition**, **Cause**, **Solution** (numbered list for sequential steps; bullets or sub-headings for several equally valid solutions).
- **Custom resource (CR):** file named `{RESOURCE_NAME}.md`; title is the CR name in CamelCase (`LogPipeline`, `Function`).

### Shell commands and code blocks (cross-ref: Pass 1 formatting)

- **Runnable command** (meant to be copied): show the command alone — no prompt prefix, no `$`/`#`/path/hostname. Use environment variables instead of hardcoded values; omit `export` unless the value is noteworthy.
  ```
  kubectl -n $NAMESPACE get pods -l app=$APP_LABEL
  ```
- **Command-plus-output transcript:** prefix only command (input) lines with `$ `; leave output unprefixed; omit path/hostname.
  ```
  $ kubectl -n kyma-system get pods
  NAME                          READY   STATUS    RESTARTS   AGE
  istiod-7c8b9f4d4f-abcde       1/1     Running   0          12d
  ```
- Mark omitted parts of a snippet or response with `...`.

### Links (cross-ref: Pass 2 for link-text wording)

- Relative links within the same repository; absolute links for other repos and external sources.
- Descriptive link text; never vague ("this", "that", "here") and don't link whole phrases. Prefer the cross-reference pattern `For more information, see [target].` Keep links syntactically separate from descriptive sentences (in a cross-reference formulation, after an introductory colon, or in a closing "Related Information" section).
- Don't over-link: skip content that is peripheral, well-known to the audience, or easily searched. Keep to a manageable number per topic.
- Link to a heading with `#{name-of-the-heading}` appended to the filepath. Reference assets (YAML/JSON/SVG/PNG/JPG) in the topic's `assets` folder via a GitHub relative link.
- When mentioning a configuration file, consider linking to it instead of naming it; use the name without the format extension.

### Advisory information moderation (cross-ref: Pass 1 panel syntax/types)

Use advisory panels in moderation — as a rule no more than ~2 per topic, and not directly one after another. Overuse dilutes them. A plain "Note that ..." sentence in body text is sometimes enough. Provide repeated advisory info centrally.

### Examples and sample data

- Bias-free examples reflecting diversity. Prefer existing product-area sample data.
- Names/data must not match real people/companies or be humorous/branded; never include real users, passwords, or internal info.
- Fake links use `https://www.example.com`.
- Set off examples by size: page+ → a `Example: <Title>` topic; ~half page → an "Example" section; a few sentences → a block with an "Example" signal; word/sentence level → "for example" / "such as".
