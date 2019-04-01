# Faster Code Review Guideline

Kyma is quite a large project with many areas, each has a group of core developers
who are working hard on developing that area of Kyma. While that means they are
working hard on new features and enhancements, they also work on doing code
reviews for PRs coming from different contributors from the community and it is
crucial that the code review process of new PRs becomes quick and efficient for
both parties. This guide defines a set of principles and practices that can help
achieving such a goal.

## 0. Familiarize yourself with project conventions

- [CONTRIBUTING.md](../CONTRIBUTING.md)
- [Naming conventions](naming.md)
- [Content strategy](content-guidelines/content-strategy.md)
- [Content guidelines](content-guidelines)
- [Coding standards](coding-standards)

## 1. Before coding communication

We recommend these steps before you jump into coding that nice feature you want

1. Make sure you have a GitHub issue created for the work you want to do that
   - Follows one of the issues templates and make sure to describe well the
    the changes you intend to introduce.
   - If possible identify in the issue the relevant area label and components
    you might touch and change.
2. Communicate that issue upfront on the slack channel of the relevant area and
start a discussion there to have it pre-reviewed.
3. Ask the area group to assign you an issue buddy from the CODEOWNERS.
    >Since reviewing a PR always needs context, it is best if the whole pre-review,
assistance, and code review stays with a single area member who is also a
CODEOWNER. We call that person an `issue buddy`.

## 2. Don't build a cathedral in one PR

Are you sure you will have the CODEOWNERS buy-in and support in a single PR?
Are you sure they won't have doubts about your technique and whether this is the
best way to tackle that issue? Are you willing to bet a few days or weeks of
work on it?  If in doubt, consider one of the following:

1. Start with a design proposal and have it reviewed.
2. Create WIP PR where you get early feedback on initial implementation.
3. Whenever possible, break up your PRs into multiple commits.
   >Making a series of discrete commits is a powerful way to express the
evolution of an idea or the different ideas that make up a single feature.
There's a balance to be struck, obviously. If your commits are too small they
become more cumbersome to deal with. Strive to group logically distinct ideas
into commits.
4. Multiple small PRs are often better than multiple commits.
    >If you can extract whole ideas from your PR and send those as PRs of their
own, you can avoid the painful problem of continually rebasing. Obviously,
we want every PR to be useful on its own, so you'll have to use common sense in 
deciding what can be a PR vs what should be a commit in a larger PR.

## 3. Efficient code review

1. When PR is ready, schedule a short 15 min walk through video call with your
    issue buddy.
    > To avoid a PR going ping-pong, always favor direct communication as a way
to explain what is already in the PR. Whether to give your buddy a walkthrough
or to discuss comments that were already provided on the PR.
2. Agree with your issue buddy on a time to provide the final code review.
3. Try to make the expectations clear. If PR is not expected to be reviewed for
now, consider blocking the issue until feedback is available.
4. Try to include all stakeholders in the discussions as early as possible.

## 4. Fix feedback in a new commit

Your reviewer has finally sent you some feedback. You make a bunch of changes
and ... what? You could patch those into your commits with git "squash" or
"fixup" logic. But that makes your changes hard to verify. Unless your whole PR
is pretty trivial, you should instead put your fixups into a new commit and
re-push. Your reviewer can then look at that commit on its own - so much faster
to review than starting over.

We might still ask you to clean up your commits at the very end, for the sake of
a more readable history, but don't do this until asked, typically at the point
where the PR would otherwise be approved.

General squashing guidelines:

Sausage => squash

When there are several commits to fix bugs in the original commit(s), address
reviewer feedback, etc. Really we only want to see the end state and commit
message for the whole PR.

Layers => don't squash

When there are independent changes layered upon each other to achieve a single
goal.

> A commit, as much as possible, should be a single logical change. Each commit
should always have a good title line (<70 characters) and include an additional
description paragraph describing in more detail the change intended. Do not link
pull requests by # in a commit description, because GitHub creates lots of spam.
Instead, reference other PRs via the PR your commit is in.

## 5. KISS, YAGNI, MVP, etc
Sometimes we need to remind each other of core tenets of software design - Keep
It Simple, You Aren't Gonna Need It, Minimum Viable Product, and so on. Adding
features "because we might need it later" is antithetical to software that
ships. Add the things you need NOW and (ideally) leave room for things you might
need later - but don't implement them now.

## 6. Push back
We understand that it is hard to imagine, but sometimes we make mistakes. It's
OK to push back on changes requested during a review. If you have a good reason
for doing something a certain way, you are absolutely allowed to debate the
merits of a requested change. You might be overruled, but you might also
prevail. We're mostly pretty reasonable people. Mostly.

## Final: Use common sense
Obviously, none of these points are hard rules. There is no document that can take the place of common sense and good taste. Use your best judgment, but put a bit of thought into how your work can be made easier to review. If you do these things your PRs will flow much more easily.