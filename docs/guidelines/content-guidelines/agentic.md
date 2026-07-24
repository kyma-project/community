# Kyma technical writer review — style ruleset

Rules for reviewing SAP/Kyma documentation. Assembled from the SAP Technical Writing Style Guide (primary) and the Kyma community content guidelines (secondary), with SAP/Kyma over-the-source requirements marked `(override — …)`. Pure rules: how documentation must be written.

Review documentation in four sequential passes. The Document structure and conventions section defines the topic-type shapes your content must conform to — consult it from any pass when topic classification, required sections, or heading conventions are relevant.

- **Pass 1 — Formatting and mechanics:** how text is formatted, punctuated, capitalized on the page.
- **Pass 2 — Language and grammar:** word choice, grammar, voice, mood, tense, sentence quality.
- **Pass 3 — Terminology and naming:** what to call things.
- **Pass 4 — Reconciliation sweep:** re-read Passes 1–3, resolve cross-pass conflicts, verify topic-shape conformance.

## Pass 1 — Formatting and mechanics

### Headings and titles
- Write titles/headings in **title case** (uppercase first and last word; nouns, verbs, adjectives, adverbs, gerunds, participles, "Not", forms of "Be"; determiners/pronouns like This/Your; subordinating conjunctions; prepositions of 5+ letters like "About"). Lowercase: articles (the/a/an), coordinating conjunctions (and/but/or), prepositions under 5 letters (to/on/in/from). Capitalize hyphenated words as separate words (`Sold-to Party`); capitalize words in parentheses and after a colon as first words.
- No periods in titles. Use colons, parentheses, question marks sparingly; avoid exclamation points. No extra formatting in titles (they are already set off) — exception: technical-name tags (for example, parameter names).
- **Heading formulation depends on topic type** (see **Document structure and conventions**): task titles use the **imperative verb** (`Create a namespace`), not a gerund and not "How to…"; concept titles are **nominal** (`Authorization Concept`). Gerunds are fine in body content, not in Kyma task-topic headings. Use action verbs and present tense in headings.
- Product/approved names in titles take title case even when their body-text form is lowercase (`SAP HANA Cloud Connector` in a title). No abbreviations in titles (rephrase; last resort: abbreviation in title + full form in body).
- Heading levels: H1 = document title; use H2/H3 to organize. Do **not** use H4 or smaller (if you need one, restructure — see Pass 4). Keep headings concise (ideally one line) but descriptive.
- **Stacked headings** (override — relaxed per techwriter review): prefer an introductory paragraph between a heading and the next heading. Do **not** flag a stacked heading when an intervening sentence would add no information beyond what the heading structure already conveys (self-evident from context).

### Lists
- Introduce a list with a complete sentence ending in a colon; do not embed list items in the introductory sentence. Do **not** state the number of items (it may change) — use "the following" or "as follows".
  - (override — relaxed per techwriter review): do **not** flag a missing introductory sentence when the list's purpose is self-evident from the preceding heading or context. The bar on stating a numeric count is **not** relaxed; "the following"/"as follows" remain endorsed.
- Choose list type by sequence: **numbered/ordered** for sequential steps; **bulleted/unordered** for unordered items.
- Use **parallel formulations** across items (do not mix noun/verb phrases with complete sentences). Make items easy to scan; put the important information first (a bold lead phrase is fine — bold the term, not "You can …").
- Use at least 2 items (exception: generated/fill-in lists). 8+ items on one level → consider restructuring. Prefer ≤2 levels.
- Set off optional/conditional steps: `Optional:` for optional; an if-clause or stated condition for conditional.
- **List-item punctuation:** complete sentences take a period; fragments take none (abbreviated questions may take `?`). Apply **one consistent punctuation approach across the whole list** — do not mix periods, commas, none; do not end items with `;` or `,`. If a mix of complete/incomplete sentences is unavoidable, still apply punctuation consistently (a fragment also gets a full stop when the chosen approach uses full stops).
- **List-item capitalization (English):** capitalize the first word of each item, whether complete sentence or fragment (exception: items that are always lowercase, such as parameter names).
- Defining a term in a list: bold the term, then a **colon**, then the definition (override — O11, pending source update). Apply one format consistently throughout the document. Colon guidance is separate from the semicolon guidance below.
  - **Do:** `**Namespace**: an isolation boundary for Kubernetes resources.`
  - **Don't:** `**Namespace** - an isolation boundary for Kubernetes resources.`

### Tables
- Introduce a table with a complete sentence ending in a colon; do not state the number of rows/columns. Use a header row with meaningful titles; proportional column widths; parallel formulations down each column. Max ~5 columns — restructure if overloaded.
- Keep cell text short and specific. Leave a cell empty when there is no information (or write "Not applicable" if absence would look like a mistake). Do **not** use X marks — use words ("Yes"/"Correct"). Keep punctuation consistent within a mixed column.
- Capitalization: table caption and header cell → title case; table cell → sentence style. Captions (if used): use for all tables, nominal phrasing, above the table, unnumbered, without the word "table".
- Center-align choice-type columns (`Yes/No`, `true/false`). A **Default** column uses `None` where a parameter has no default value.
- Break a long or hard-to-read table into multiple tables. Prefer a table over a list when the key information is otherwise buried.

### Panels / admonitions (override — Help Portal blockquote syntax, O2)
- The correct syntax is the blockquote panel:
  ```markdown
  > ### Note:
  > Content here.
  ```
- Valid panel types (SAP): **Note, Tip, Recommendation, Caution, Remember**. **Restriction** only with extreme caution and only when certain it does not describe a software limitation affecting revenue recognition.
- The community `[!NOTE]` / `[!WARNING]` / `[!TIP]` GitHub/VitePress alert syntax is **not valid** for the Help Portal — flag it.
- Choose type by purpose: Note = important/non-obvious info, no action; Tip = optional advice to avoid minor issues; Recommendation = advantageous settings/methods; Caution = severe hazards (data loss, system failure); Remember = fundamental info needed later.
- Use advisory info in moderation (≤2 per topic as a general rule; do not stack them). For recommendation phrasing, see Pass 2 (O4).

### Code and bold formatting
- Use **bold** for: parameters/variables, HTTP headers, events, roles, and UI elements.
  - **Do:** `Choose **Subscribe**.` (use the device-neutral verb `Choose`/`Select`, never `Click`/`Tap` — see Pass 2)
- Use `code` font for: code examples, values, endpoints, file names, file extensions, path names, repository names, status/error codes, parameter-value pairs, metadata names, flags, GraphQL queries/mutations, rights, and custom resources **when referring to the code specifically**.
- Placeholders: SCREAMING_SNAKE_CASE in curly brackets with a descriptive name — `{NAMESPACE}`, `{YOUR_PROJECT_NAME}`.
- Do **not** enclose trailing punctuation (colon, comma, sentence-ending period) inside formatting — translation software segments on `.`/`:`. **Exception:** a panel-type label that forms one stable translation segment, for example `**Note:**` (limited to the panel-type labels — Note, Tip, Caution, …; see Pass 4).

### Kubernetes resource names (override — O1; primary statement, cross-ref Pass 3)
Kubernetes resource names (built-in and custom) are written in **plain CamelCase in body text** — no bold, no code font. In an inline code span (a `kubectl` command, or referring to the resource *in the code* specifically), code font comes from the code span itself, not as a Kubernetes rule. The code-vs-conceptual boundary is subjective, so:
- **Code (format as code):** "When you apply the `APIRule` manifest, set `spec.host`."
- **Conceptual (plain CamelCase):** "API Gateway operates on APIRule custom resources."
- Capitalize a resource word only in the Kubernetes sense: `Node` (K8s resource) vs. a VM node (lowercase); `CustomResourceDefinition (CRD)` capitalized, but "custom resource" alone is lowercase.
- Genuinely ambiguous code-vs-conceptual cases: flag for human review (see Pass 4, Tier 2).

### Punctuation and characters
- **Colon:** use to introduce information (lists, tables, definitions). Capitalize after a colon before a complete sentence; lowercase before a fragment (exception: titles/proper nouns). The colon is a normal, encouraged introducer.
- **Semicolon:** use sparingly with main clauses — **prefer separate sentences** (shorter, easier to read). Use a semicolon only to separate long/complex series items that contain internal commas, or rephrase into a list. Do not end list items with a semicolon.
- **Serial (Oxford) comma:** required before the conjunction in every series of 3+ items, **with no exception** (override — O12) — including series joined entirely by conjunctions.
  - **Do:** `Paris, Rome, or Brussels`
  - **Don't:** `Paris or Rome or Brussels`
- Comma before a coordinating conjunction (and/but/or/so/yet) that introduces an independent clause; after an introductory adverbial clause, a 5+-word introductory prepositional phrase, or a sentence-referring adverb (therefore/however); with "that is" and "for example"; with "such as" only when nonrestrictive. No comma between subject and verb.
- **Dashes (override — O9): use the plain hyphen (`-`) only; do not use en dashes (`–`) or em dashes (`—`)** in Kyma docs.
  - **Number ranges:** plain hyphen, no spaces — `0-9`, `50-100`.
  - Do not use en/em dashes for asides — rephrase, or use commas, parentheses, or a colon.
  - When a range start is ambiguous, use the word `to` (`50 to 150,000`, not `50-150,000`).
  - Flag any `–` or `—` in reviewed content and replace with a hyphen or a rephrase.
- **Slash:** do **not** use a forward slash as shorthand for "or" — write "or", or "…, …, or both"; `and/or` only if that reads clumsy or under space limits. Exceptions: part of an approved name/proper noun (`SAP S/4HANA`); an established unit modifier (`plan/actual comparison`); and `countries/regions` (see Pass 3). No other slash-as-"or" is permitted. Use slashes as URL/path separators.
- **Special-character reference pattern:** refer to a special character as *name* (*symbol*) — for example, `ampersand (&)`, `asterisk (*)`.
- No apostrophe for plurals of nouns/acronyms (`BAdIs`, not `BAdI's`). Use the typographically correct apostrophe; never an accent as a surrogate.
- Quotation marks: do not use them to refer to titles/headings/captions (use formatting), for irony, or as surrogate formatting except as a last resort. Enclose ending punctuation inside quotation marks only for a complete quoted sentence. Do not format approved names.
- Avoid ampersands (&), technical/mathematical symbols, and symbols with multiple meanings in UI text; commonly used symbols (`%`) are acceptable.
- One space between words and after punctuation; no space in period-abbreviations (`a.m.`); no space breaks inside URLs/paths/technical names; use nonbreaking spaces to keep closely related words together (`5 kg`). No Latin abbreviations in body text — use the full English form (see the Pass 2 plain-word table for `e.g.`/`i.e.`/`etc.`).

### Numbers, dates, units, currency
- Spell out one–nine; numerals for 10+ (including ordinals). Use numerals to draw attention to a number or for figures/ages/times/dates/units/currencies/percentages/page numbers. Spelled-out number in fixed phrases (`third party`).
- Thousand separator (English comma) for 4+ digits before the decimal (`7,654,321`), grouped in 3s; exception for an odd stand-alone low number with no unit (`1500 persons`). Decimal point in English; leading zero for fractions < 1 (`0.5`). Prefer words for millions/billions.
- Metric units with standard symbols; no plural "s" on symbols; repeat the unit in ranges (`10 mm × 17 mm`). In English, no space before `%`. Hyphenate a numeral plus a spelled-out unit (`18-mile`), not an abbreviated unit (`2 kg`) — see Pass 2 hyphenation.
- Dates: month as a word, year as 4 digits (`August 5, 2024`) — not `08/05/2024`. Times: mark a.m./p.m.; time zone abbreviated in parentheses. ISO 8601 (`YYYY-MM-DD`) only in technical texts (log files), with hyphen separators.
- Currency: ISO 3-letter code before the amount in English (`EUR 100`). Omit trailing `.00` when a sum has no fractional digits — **exception:** keep aligned decimals when another sum in the same sentence/paragraph has fractional digits (`USD 70.00 net or USD 74.45 gross`). Separators for 4+ digits; minus sign immediately before a negative amount.

## Pass 2 — Language and grammar

### US English
- Use American English spelling: `-ense` (defense, license), `-er` (center), `-ize/-yze` (organize, analyze), `-og` (catalog, dialog), `-or` (color); `canceled`/`modeling` (single `l`); `fulfill`, `program`, `toward`, `color` (not `colour`).
- Use the American word/phrase: `fill out` (not `fill in`; see Pass 3, O8), `as of` (not `as at`), `for more information about` (not `on`), `parentheses`/`curly brackets`/`square brackets`, `period` (not `full stop`), `exclamation point`, `quotation marks`. Exception: keep country-specific terms in legal/financial/HR contexts (for example, `ZIP code` for the US).

### Voice and mood
- **Active voice is the strong default** — it names who does what, uses fewer words, and is clearer for non-native readers. Address the user directly with "you" or an imperative. Name the agent in the third person when not "you" (the user, the system, the service).
- Passive voice is acceptable only when: the agent is unknown, unimportant, or omitted for politeness; for apologies/blame removal; or when the context already makes the agent clear and passive improves flow.
- Imperative mood for instructions (procedures, steps) — it tells the reader directly to do something. Do **not** convert a recommendation to a bare imperative (see O4 below).
- **Do:** "The endpoint path includes your service name." **Don't:** "Your service name is to be included in the endpoint path."

### Recommendation language (override — first-person ban preserves advisory tone, O4)
Do not use first-person pronouns (`we`, `us`, `let's`) — but preserve recommendation semantics. Do not turn advisory content into a mandatory imperative (that changes "suggested" to "required").
- Rewrite `We recommend {X}` → `It is recommended to {X}.` (accepted passive), `Consider {X}.`, or `As a best practice, {X}.`; or use the **Recommendation** panel.
- `It is not recommended…` → state positively what to do instead.
- The imperative rule and the `should → must/can` rule do **not** apply in recommendation contexts.

### Sentence quality
- **Keep sentences at or under about 20 words** (override — O15). If a sentence runs longer, split it into shorter sentences, or into a list or a table. Split any sentence with two or more subordinate clauses.
- Use simple syntax; keep syntactic cues (articles, "that" to introduce object clauses and restrictive relative clauses) — no telegraphic style. Keep phrasal verbs together (`Fill out this form`, not `Fill this form out`). Prefer a relative clause over a participle construction; do not use causative "have"/"get".
- Use **positive formulations**; rephrase double negatives.
- **What before how:** state the goal/consequence before the action. Describe chronological actions in sequence. In conditional sentences, put the condition ("if") before the statement; no "then".
- Split a long series of nouns/prepositional phrases into smaller units to avoid ambiguity. Do not generalize ("generally", "usually", "sometimes") — list the actual cases. Make every word count; use intensifiers sparingly. Prefer precise verbs; avoid SAP jargon like `maintain` (use `edit`, `create`, `change`, `manage`).
- Do **not** use "Note that…" (or "Please note that…") lead-ins (override — O16) — delete the lead-in and state the fact directly, or use a **Note** panel for a genuine advisory.
  - **Do:** `The operator reconciles the resource every 30 seconds.`
  - **Don't:** `Note that the operator reconciles the resource every 30 seconds.`

### Tense (override — O13, pending source update)
- Use **present tense** for documentation content. Reserve the future tense ("will") strictly for actions that genuinely occur in the future relative to the described step — not for what a system does, what a result is, or what happens as a consequence of a step (those are present tense). Flag future-tense constructions that describe present behavior and rewrite them in the present. The one acceptable forward reference is within a procedure ("Note down the token. You need it in a later step.").
  - **Do:** `When you apply the manifest, the operator creates the resource.`
  - **Don't:** `When you apply the manifest, the operator will create the resource.`

### Contractions
- **Use common contractions — they make writing sound natural and human; the spelled-out form sounds robotic by comparison.** Test by reading aloud: if you'd use a contraction when speaking, use it.
  - **Do:** "you **don't** need to provide your account information." **Don't (robotic):** "you **do not** need to provide your account information."
- **Exception (emphasis / serious warnings):** spell out for emphasis — a warning reads more seriously spelled out. Use `do not` / `cannot` (never `can not`), `is not`: "**Do not** manually edit the INI file."
- Do **not** use unusual or compound contractions (`couldn't've`, `the user'll`). Watch `it's` (it is) vs `its` (possessive).

### Politeness (override — O14)
- Use "please" only when appropriate (an inconvenient request, or a troubling situation such as an error) — not for plain instructions or cross-references. Be polite only in rare cases.
- Apologize without blaming the user or the software; use the passive voice to remove blame.

### Grammar
- Singular verb with collective nouns (`data`, `team`, `SAP`, `company`): "The data **is** saved."
- Percentage verb agrees with the referent noun (singular/plural).
- Relative clauses: restrictive = `that`, no commas; nonrestrictive = `which`, commas; `who` for persons. Avoid dangling modifiers. Keep `that` (don't omit) in restrictive clauses; a stranded preposition is fine when it sounds natural.
- Be careful with singular `they`/`their` — rewrite if it could be unclear for translation or non-native readers.
- Articles: verify `a`/`an`/`the`. Use `the` for a specific referent, `a` for something non-specific; do **not** use an article before a product name or one of its components (`Application Connector`, not `the Application Connector`). Use articles to avoid ambiguity.

### Hyphenation and compound words
- **Hyphenate per US conventions** — the purpose is to prevent misreading, and US hyphenates less than British. US prefixes are mostly closed (no hyphen): `autogenerated`, `preassembled`, `nonnegotiable`, `cooperate`, `multimedia`. Hyphenate only to prevent confusion or for fixed forms: `cross-application`, `self-service`, `well-known`, `high-level`.
- Hyphenate a **phrasal (compound) adjective before a noun** (`well-known actor`, `copy-protected folders`), but **not** when it stands alone (`the actor is well known`) or is modified by an adverb (`highly configurable applications`, not `highly-configurable`).
- **Suspended hyphen** only when the **second** element is omitted (`short- and long-term memory`) — never when the first is omitted.
- Hyphenate a **numeral plus a spelled-out unit** (`18-mile highway`); no hyphen with an abbreviated unit (`2 kg weight`, `200 GB hard drive`).
- Do not over-hyphenate noun compounds — hyphenate only to prevent misreading; a compound of 4+ nouns usually signals you should rephrase.

### Tone and inclusive language
- Positive, conversational-but-concise voice; not chatty, no exaggeration. Use excitement and exclamation points sparingly. No capitalization for emphasis. No emoticons/emojis. No jargon, colloquialisms, or dialect.
- Present functionality from the target group's perspective — describe what they can do, not the feature. Address the user with "you".
- Replace non-inclusive terms (prefer `blocklist`/`allowlist` over `blacklist`/`whitelist`; avoid `master`/`slave`) — prefer a clearer rephrase over a one-to-one swap. Gender-inclusive nouns (`chairperson`, `workforce`, not `chairman`, `manpower`). Prefer imperative/"you"/third-person-plural over gendered singular; never `he/she`.
- Address users by their need, not their disability (`if you use a screen reader`, not `wheelchair-bound users`). Do not use `accessible` to state a system's accessibility status or to mean `available` — say `available`.
- Describe bias-free examples; do not use real people's/companies' names, addresses, or confidential data in examples (use `example.com` for fake links).

### Global English and clause usage
- Use common, everyday words; make sentence structure explicit (don't drop conjunctions, articles, prepositions for brevity). Sentences must sound natural, even if that means rephrasing.
- Use "that" to introduce an object clause ("verifies **that** the name matches") and in restrictive relative clauses. Clarify what "and"/"or" join. Avoid transitional-phase words (`currently`, `now`) that promise future functionality.

### Plain, common words — use → avoid
Prefer simple, common words. No Latin abbreviations in body text — use the English form; this covers the `e.g.`/`i.e.`/`etc.` rows below.

| Use | Avoid |
| --- | --- |
| use | utilize |
| use | leverage *(override — O10)* |
| get | obtain |
| happen | occur |
| agree | concur |
| although | albeit |
| before | prior to |
| harmful | deleterious |
| start, begin | commence |
| try | endeavor |
| up to now | heretofore |
| decide | take a decision |
| upgrade | perform an upgrade |
| explain | provide an explanation |
| can | be able to, it is possible to |
| is transparent | is characterized by transparency |
| outcome | final outcome |
| for example, such as | e.g. |
| that is | i.e. |
| and so on | etc. |
| and others | et al. |
| note | n.b. |
| opposed to, versus | vs. |
| compare | cf. |

(Kyma/domain-specific term pairs — `backend`, `frontend`, `YAML`, `repository`, `via`, and so on — are in the Pass 3 terminology table, not here.)

## Pass 3 — Terminology and naming

> No internal-SAP-systems lookups (override — O3): this ruleset is self-contained. Rules that in the source depend on SAPterm or the SAP Approved Names repository are inlined here with the relevant data; a reviewer must not be told to "look it up" in an internal system.

### Product naming (override — SAP Brand approved names, O6)
- The full name is **`SAP BTP, Kyma runtime`**. Always include `runtime` after Kyma — **flag `SAP BTP Kyma` without `runtime`**. Lowercase `runtime` except where title case is required.
- **First mention → short form.** Use the approved full name at first mention, then an approved short form afterward. Do not contort a sentence around the full name when a short form exists (backs Pass 4 Tier 2).
  - Body-text natural language is fine: "the Kyma runtime for SAP BTP".
  - Space-limited contexts (menus, tiles, web): `SAP BTP Kyma runtime` or `Kyma runtime` (comma and/or `SAP BTP` may drop). Cockpit/service-catalog tile short name: `Kyma Runtime`.
- Do **not**: write `SAP BTP Kyma`; abbreviate (`SAP BTP, KR`); alter the name (`SAP Kyma`).
- **`Kyma dashboard`**: refer to it as `Kyma dashboard`; on subsequent mention, `dashboard`. Lowercase `dashboard` except where title case is required. Do **not** precede with `SAP` or `SAP BTP` (`SAP Kyma dashboard` is wrong).
- Spell out "Business Technology Platform" at an early point so the `BTP` abbreviation is clear.

### General SAP naming
- Use approved names exactly — correct spelling/wording, no plurals or possessives, no invented abbreviations. Do **not** add trademark symbols in user assistance. Do **not** set off approved names with italics/quotes, and do **not** translate them.
- No article before a stand-alone approved name ("SAP Business One helps…", not "The SAP Business One helps…"). At first occurrence use the approved descriptor (solution, service, application, platform…); afterward use the name stand-alone (or a generic "software"/"offering").
- Use the prefix `SAP` with care — do not create phrases misreadable as offerings. Avoid compounds with approved names (rephrase with a preposition). Use approved names for AI and specify the type using approved qualifiers; avoid pronouns where possible.

### Abbreviations and acronyms
- Prefer no abbreviation, especially in UI text. Introduce an unavoidable abbreviation with the full form + abbreviation in parentheses at first occurrence, then use the abbreviation; **exception:** use straightaway if more common than the full form (`ID`, `PC`, `ABAP` for developers). Do not coin new abbreviations or use CamelCase in abbreviations (CamelCase is fine only in non-translated technical names). No abbreviations in titles. Plurals of acronyms take no apostrophe. Use an indefinite article as in speech (`a BAPI`, `an SAP solution`).

### Software objects and entities
- Determine what a term refers to before formatting it: an approved name, a real-life thing, a software object, or a software entity — each is handled differently.
- **Real-life entity** (or when "this is software" is irrelevant): write like any other word — no formatting, no title case (`a strategic purchaser creates a purchasing contract`).
- **Software object** (role, API, function module): at first occurrence set off *Title in Title Case* (`technical name`) *descriptor*; later use the title alone. Do not overuse title case.
- **Software entity** (tool, engine): treat like a proper noun — title case, no formatting, use an article, do not prefix with `SAP` (`The BAdI Builder…`).

### Kubernetes and Kyma capitalization
- `Kubernetes` is always capitalized — never `kubernetes` or `k8s`.
- Kyma component names and headings: title case (`Application Connector`, `API Gateway Controller`).
- Kubernetes/custom resources: CamelCase (`ConfigMap`, `ConfigMaps`; in titles/navigation add spaces: `Config Map`). Format as code only when referring to the code specifically (see Pass 1 / O1). Capitalize these **only** when referring to the Kubernetes resource; otherwise lowercase:
  `ConfigMap, CronJob, CustomResourceDefinition (CRD), Deployment, Function, Ingress, Node, PodPreset, Pod, ProwJob, Secret, Service, ServiceBinding, ServiceClass, ServiceInstance`.
  - `Node` = K8s resource only; a VM or billing node is lowercase. "custom resource" alone is lowercase (not Kubernetes-specific).
- Do **not** capitalize `namespace`. Use sentence case for standard sentences; do not over-capitalize words just because they seem special or important.

### Country/region
- Do **not** use `country` alone — use `country/region` or `country or region`, or rephrase (`Local Version`). Field labels use title case (`Country/Region`). This `/`-as-"or" is an **approved exception**; no other slash-as-"or" is allowed (see Pass 1).

### Topic-type naming
Titles follow the topic type (see **Document structure and conventions**): Task = imperative verb (`Create a namespace`); Concept = nominal (`Security Concept`); Troubleshooting = the symptom/error message (not the cause); Custom Resource = the CR name in CamelCase, file `{RESOURCE_NAME}.md`.

### Kyma preferred terminology — use → avoid

| Use | Avoid | Note |
| --- | --- | --- |
| ID | id | |
| and | + , & | |
| backend | back end, back-end | |
| frontend | front end, front-end | |
| micro frontend | microfrontend, micro front-end | |
| connect, connection | integrate, integration | |
| custom resource | Custom Resource, CustomResource | lowercase unless a K8s resource name |
| document | doc | |
| email | e-mail | |
| **fill out** | complete, fill in | *(override — O8; agrees with Pass 2 US-English)* |
| key-value | key/value, key:value | |
| must | have to, need to | |
| need | require | |
| or | / (slash) | except approved names, unit modifiers, `countries/regions` |
| repository | repo | |
| run | execute | |
| single sign on (SSO) | single sign-on | |
| typically | usually | |
| using, with | via | |
| YAML | yaml | file extension/name: `.yaml` |
| Prow Job (process) | Prowjob, prowjob | resource: CamelCase `ProwJob` |
| must, can | should | mandatory → must; optional → can |
| (you) can | it is possible to, allows you to, you have the option to | else drop "can" and use the imperative |
| that is | i.e. | often the clause adds little — drop it |
| for example, such as | e.g. | mid-sentence, "such as" beats "for example" |
| cloud-native (adj.) | cloud native | noun compounds usually no hyphen; adjective compounds often do |
| Application (external solution via Application Connector) | app | "application" = a microservice or software |
| Infrastructure Provider, IaaS Provider | hyperscaler, Cloud Provider | |

- The count-avoidance rule ("do not state the item/row/column count — use 'the following'/'as follows'") lives in the Pass 1 list/table-introduction rule.
- Use terms consistently (no synonyms); prefer the standard verb (`create`, not make/build/generate as synonyms).

## Pass 4 — Reconciliation sweep

Re-read the changes from Passes 1–3, confirm no pass introduced a new violation, check consistency across the whole document, and verify the document conforms to its topic-type shape (see **Document structure and conventions**). Signals to flag: a task over ~9 steps (consider splitting — a signal, not an automatic failure); any H4 (the document may need restructuring).

### Cross-pass conflict-resolution framework
When fixing a rule from one pass would break a rule from another, resolve by these tiers — **the higher tier wins on the specific token; lower tiers rephrase around it:**

1. **Topic-type constraints (binding).** Topic types are structural contracts, not suggestions — content that does not fit a type is reconceived, not squeezed in (see the non-negotiable topic-type framing in **Document structure and conventions**). Signals to flag: a task over ~9 steps (consider splitting); any H4 usage (the document may need restructuring).
2. **Terminology and naming correctness.** Wrong names are factual errors; sentence flow is always fixable, so naming wins over flow. Use brand-approved names and approved **short forms** on subsequent mention (O6); no article before an approved product name; product names in headings may take title case; Kubernetes resources are plain CamelCase in body / code font in code context — **flag genuinely ambiguous code-vs-conceptual cases for human review**.
3. **Semantic formatting intent over Markdown rendering.** Formatting follows the semantic role of a span, not typography; when a Pass 1 formatting rule conflicts with a span's semantic role, intent prevails. *(Light note: the target format is typically DITA, where code font maps to `<sap-technical-name>` and a UI label to `<ui-element>`; see the SAP `cms-tags-guide`. This ruleset stays Markdown-focused.)*
4. **Natural-language flow.** Rephrase to accommodate tiers 1–3, but do **not** generate filler to satisfy a structural rule — if a list's purpose is obvious, omit the introductory sentence rather than emit a "the following items are:" placeholder (consistent with the O5 list-intro relaxation). Use common abbreviations only when introduced at first mention; do not coin new ones.
5. **Author intent for instructions (needs human judgment).** The instruction's purpose sets the verb pattern: instruction → plain imperative; mandatory → "you must"; recommended → approved advisory pattern (O4); optional → "can" / "Optionally, …". When intent is genuinely ambiguous from the text, **flag for author decision** rather than guessing.

### Specific tie-break resolutions
- Full product name makes a sentence unwieldy → use the approved short form on subsequent mention; do not rephrase around the full name when a short form exists (O6).
- List-item lead word is a product name → bold it for pattern consistency within the list.
- Gerund in a heading → imperative prevails for task topics; if it describes an ongoing state, reclassify as a concept and use an unambiguous nominal phrase.
- `countries/regions` slash-as-"or" → approved exception; **no other** slash-as-"or" is permitted.
- Colon inside formatting (`**Note:**`) → exception limited to panel-type labels (Note, Tip, Caution, …); confirms the Pass 1 formatting exception.
- `fill in` (Kyma table) vs `fill out` (SAP/US English) → **`fill out`** wins (confirms O8; see Pass 3 — do not emit a competing rule).

### Edge-case fallback (override — O7)
When the documentation under review presents a situation not covered by any rule here, apply general best-practice judgment and note that no specific rule matched.

## Document structure and conventions

Reference material for all passes — not a step performed after them. It defines the topic-type shapes and structural conventions the passes conform content to. Consult it from Pass 1 (heading/title formulation), Pass 3 (topic-type naming), and Pass 4 (topic-shape conformance).

### Topic types (non-negotiable)
Kyma documentation is topic-based (one file per topic). **Topic types are non-negotiable, mandatory shapes** — hardcoded, DITA-bound patterns. Classify the content and make it **conform** to its topic-type shape; content that does not fit a type is reconceived, not squeezed in.

- **Concept** — answers "what is it?", gives background. Nominal title (`Security`, `Security Concept`). Must **not** give instructions or hold reference tables/lists (link to task/reference topics instead).
- **Task** — "how to" for one procedure. **Imperative-verb** title (`Define resource consumption`), not a gerund, not "How to…". Structure: intro paragraph ("why?"), prerequisites if needed, numbered steps, expected result. Roughly **5–9 steps**; longer probably splits.
- **Troubleshooting** — a condition to correct, its cause(s), and remedies. Title names the **symptom or error message**, not the cause. Standard headings **Condition / Cause / Solution**; numbered list for sequential fixes, bullets/sub-headings for equally valid alternatives.
- **Custom Resource** — file `{RESOURCE_NAME}.md`; title = the CR name in CamelCase (`LogPipeline`). Provides a sample CR, describes its fields, and points to the CRD; keep the structure whether written manually or autogenerated.

### Shell commands (cross-ref Pass 1 formatting)
- Use environment variables instead of hardcoded values so readers can copy-paste. Omit `export` statements unless the variable value itself is noteworthy.
- **Runnable command** (meant to be copied): show the command alone — no prompt prefix, no `$`/`#`/path/hostname.
- **Command-plus-output transcript:** prefix only the command lines with `$ ` to mark input; leave output unprefixed; no other `PS1` elements.
- **Omission in code:** quote only the relevant part; replace omitted parts with `...`.
- **Prefer long-form CLI flags** — `--namespace`, not `-n`. Long flags are self-explanatory and unambiguous, and readers copy-paste, so typing length is no concern.
- Use `in` or `to` when referring to a cluster — never `on` (`the Pod runs in your cluster`; `deploy a Function to the cluster`).

### Links
- **Relative** links within the same repository; **absolute** links to other repositories and external sources. Link to a heading with `#{name-of-the-heading}` after the file path. Reference assets (YAML/JSON/SVG/PNG/JPG) in the topic's `assets` folder with GitHub relative links.
- When mentioning a config file, consider linking to it; when linking to a file, use its name without the format extension.
- Link only what adds value — every link can rot. Do not link what is well-known or trivially searchable. Keep no more than ~7 cross-references per topic; prefer a "More Information"/"Related Information" section at the end over inline links.
- Descriptive link text; do not link whole phrases; avoid vague `this`/`that`/`here`.
  - **Do:** "see the **installation guide**." **Don't:** "see **this** guide." / "Read more **here**."

### Release notes
- Write from the user's perspective. Headlines in **sentence case**, short and summarizing. New features → an enticing paragraph (marketing the benefit), not a bare bullet list. Call issues **known issue** / **resolved issues**, never "bugs". Consistent sentence structure across a bulleted list. Provide a **migration guide** whenever a release requires manual upgrade steps (steps only — no new-feature descriptions).

### Screenshots
- Complement text, don't replace instructions; use sparingly (they age fast, are inaccessible to screen readers, and can't be translated). No directional indicators ("above"/"below") — introduce each screenshot with a brief purpose. Exclude the mouse pointer (unless functionally relevant) and irrelevant chrome (browser toolbar).
- **Alt text is required and descriptive:** `![Create a bucket](./assets/create-bucket.png)`, not `![](…)`.
- Format SVG preferred (PNG/JPG acceptable), ≤1MB, sized at 100/75/50/25%, saved under `assets`. Grey (`#D2D5D9`) 1pt border. Mark steps with blue (`#0A6ED1`) round stamps + white numbers, explained in an ordered list below. Highlight with red (`#EF2727`) 10pt arrows/boxes, at most one indicator per screenshot. Prefer Simplified UIs (blur non-essential elements); do not cover logo, menu/search icons, expand/close buttons.

### Diagrams
- Give a diagram context (purpose clear from surrounding content); consistent look, minimal noise, simple but descriptive; left-to-right workflow direction.
- **Alt text required:** `![Authentication flow between client and server](./assets/authentication-flow.svg)`, not `![](…)`. Brief if the body text explains it; add key details for screen readers otherwise.
- Export SVG under `assets`; keep it legible but not dominating (max website width 860px). White background (no transparency). Secondary backgrounds: mild blue (`#F0F6FF`) main environment, mint green (`#DEF2DD`) subsidiary. Rounded rectangles; white fill for boxes, blue (`#0A6EC7`) for actors. Grey (`#666666`) 1pt outlines and connectors (no outlines on actors/steps/backgrounds). Black Helvetica text: 15pt bold headings, 13pt primary, 12pt secondary; horizontal. Blue (`#0A6EC7`) round step stamps + ordered list. Add a reference key when introducing a distinct element.
