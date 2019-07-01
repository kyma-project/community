# Release communication

This document describes a release process communication.

## Communication channels

The communication Kyma team should be conducted simultaneously on the two important channels: 

- `c4core-xf-team` - main channel to communicate about the release progress. 
- `c4core-kyma-scrum-masters` - a channel to escalate actions in case when some team lingers with their job. In that case you can also write directly to the proper Scrum Master.

The Kyma community should be notified about the release on the external channel. Our official channel for external announcements in `#general` on the `kyma-community` slack workspace.

## Communication persons

The Kyma team persons which are responsible for the releasing:

- [Jose Cortina](https://github.com/jose-cortina)

## Release master responsibilities

At the beginning of the release the releasing team should pick a release master which will manage the release process. 

The release master is obliged to do the following things:

- Inform the Kyma team about release deadlines about one week before RC1. The release important dates should be pinned in all channels mentioned above. The Kyma team must know the planned date of the release to don't exceed them. The release master must watch if all Kyma teams respect the deadlines. In case not, he must escalate that to a proper channel or person.

- Prepare the environment to test the release candidates at least 3 days before merging it. All kyma teams should use the prepared cluster to test their components. 

- Create a excel sheet where all the manual tests are defined and he must take care that all of them are executed by the teams responsible for them. If some team doesnt executed their manual tests, that case must be escalated as well.

- Inform the Kyma teams about deadlines for cherry-picking to the release candidates.

- Inform the Kyma teams about all the release related pull requests which requires approvals in order to merge them as quickly as possible.

- Inform the Kyma teams which tests are failing to fix them as quickly as possible. It should be escalated if there is any risk that some test won't be fixed which will affect the release. The Kyma teams which are responsible for failing test should communicate with release master about fixing progress.

## Communication rules

The messages about the release sent to the Kyma slack channels should be written transparently. They should inform about each aspect of the next release step.

It's recommended to write updates about releases using the multiline sentences which won't be missed among other messages in the channel.

When the updates about the release are important, they should be pinned to the channel so they will be more visible.

## Scrum of scrums meetings

Scrum of Scrums is the internal meeting for the Kyma Scrum Masters.

The Scrum Master which team is responsible for the release should give a status about the release progress at the Scrum of Scrums meetings, so all of the Scrum Masters will know if there are some obstacles for the release.

The Scrum Master is also responsible for updating a Scrum of Scrums MagicBox under the Wiki page for the adequate release.