# Sharing of peer/client/service attributes among peers

**Authors**: [Ziv Nevo](mailto:nevo@il.ibm.com)

**Begin Design Discussion**: 2023-11-01

**(optional) Status:** wip

**Checklist**:

- [ ] Defining peer attributes
- [ ] Defining/extracting service attributes
- [ ] Defining/extracting client attributes
- [ ] Sharing peer attributes with other peers
- [ ] Sharing service attributes with other peers
- [ ] Passing client attributes on a connection request
- [ ] Passing all attributes to the policy engine
- [ ] Merging peer and client/service attributes
- [ ] Docs
- [ ] Tests

## Summary/Abstract

Attributes are properties or metadata attached to entities.
In the context of ClusterLink, attributes may be properties such as peer location, service tier or client environment.
Attributes are currently being used in ClusterLink Access Policies as a mean to group endpoints together.
Hence, in order to make decisions on an incoming/outgoing connections,
the policy engine must get all attributes of the requesting client, the target service and the two peers involved.

In a multi-cluster setting, attributes are not natively shared; trusted channels to distribute them must be therefore established.
This design proposal describes how users can define attributes for peers, clients and services, and how these 
attributes will be shared among ClusterLink peers.

## Background

### Motivation and problem space

Attributes are very convenient to use in policies. They allow referring to a large group of entities sharing a
common property, rather than enumerating all members of this group.
This makes policies more concise, more readable and easier to maintain over time.
For example, a policy to deny connections from test workloads into prod services can simply use the `env` attribute
attached to clients and services, rather than listing all entities in each environment.
Hence, ClusterLink access policies (and load-balancing policies in the future) are based on attributes.

### Impact and desired outcome

* Users will be able to attach attributes of their choice to new peers, services and clients.
* Users will be able to define policies that use client/service/peer attributes.
* User-defined policies will be properly enforced by ClusterLink's policy engine.

## Goals

The goal of this design document is provide detailed specification for the following behaviors.
* Attaching attributes to peers - UX, data structures and persistency
* Attaching attributes to imported/exported services - UX, data structures and persistency
* Attaching attributes to clients - UX, data structures and persistency
* Sharing peer attributes with another peer
* Sharing service attributes with another peer
* Passing client attribute to a remote peer upon connection request

## Non-Goals

TBD

## Proposal

### Attaching attributes to clients, services and peers
We note that there is an inherent difference between clients on one side and services and peers on the other side.
While services and even-more-so peers are ClusterLink-controlled entities, clients become known to ClusterLink
only when they attempt to connect to a remote service. In fact, this is a ClusterLink design principal [IS IT?] -
ClusterLink should be transparent to clients.

This means client-workloads should somehow present to ClusterLink their set of attributes on the first time they try to connect.


### Sharing attributes


## Design Details

This section should contain enough information to allow the following to occur:

- potential contributors understand how the feature or change should be implemented
- users or operators understand how the feature of change is expected to function and
 interact with other components of the project
- users or operators can take action to pre-plan any needed changes within their
 architecture that impacted by the upcoming feature or change if it's approved for
 implementation
- decisions or opinions on a specific approach are fully discussed and explained
- users, operators, and contributors can gain a comprehensive understanding of
 compatibility of the feature or change with past releases of the project.

This may include API specs (though not always required), code snippets, data flow
 diagrams, sequence diagrams, etc.

If there's any ambiguity about HOW your proposal will be implemented, this is the place
 to discuss them. This can also be combined with the proposal section above. It should
 also address how the solution is backward compatible and how to deal with these
 incompatibilities, possibly with defaulting or migrations. It may be useful to refer
 back to the goals and non-goals to assist in articulating the "why" behind your approach.

## Impacts / Key Questions

List crucial impacts and key questions, some of which may still be open. They likely
 require discussion and are required to understand the trade-offs of the design. During
 the lifecycle of a design proposal, discussion on design aspects can be moved into this
 section. After reading through this section, it should be possible to understand any
 potentially negative or controversial impact of the design. It should also be possible
 to derive the key design questions: X vs Y.

This will also help people understand the caveats to the proposal, other important
 details that didn't come across above, and alternatives that could be considered. It can
 also be a good place to talk about core concepts and how they relate. It can be helpful
 to explicitly list the pros and cons of each decision. Later, this information can be
 reused to update project documentation, guides, and Frequently Asked Questions (FAQs).

### Pros

Pros are defined as the benefits and positive aspects of the design as described. It
 should further reinforce how and why the design meets its goals and intended outcomes.
 This is a good place to check for any assumptions that have been made in the design.

### Cons

Cons are defined as the negative aspects or disadvantages of the design as described.
 This section has the potential to capture outstanding challenge areas or future
 improvements needed for the project and could be referenced in future PRs and issues.
 This is also a good place to check for any assumptions that have been made in the design.

## Risks and Mitigations

Describe the risks of this proposal and how they can be mitigated. This should be broadly
 scoped and describe how it will impact the larger ecosystem and potentially adopters of
 the project; such as if adopters need to immediately update, or support a new port or
 protocol. It should include drawbacks to the proposed solution.

### Security Considerations

When attempting to identify security implications of the changes, consider the following questions:

- Does the change alter the permissions or access of users, services, components - this
 could be an improvement or downgrade or even just a different way of doing it?
- Does the change alter the flow of information, events, and logs stored, processed, or
 transmitted?
- Does the change increase the 'surface area' exposed - meaning, if an operator of the
 project or user were to go rogue or be uninformed in its operation, do they have more
 areas that could be manipulated unfavorably?
- What existing security features, controls, or boundaries would be affected by this
 change?

This section can also be combined into the one above.

## (Optional) Future Milestones

List things that the design will enable but that are out of scope for now. This can help
 understand the greater impact of a proposal without requiring to extend the scope of a
 proposal unnecessarily.

## (Optional) Implementation Details

Some projects may desire to track the implementation details in the design proposal. Some
 sections may include:

### Testing Plan

An overview on the approaches used to test the implementation.

### Update/Rollback Compatibility

How the design impacts update compatibility and how users can test rollout and rollback.

### Scalability

Describe how the design scales, especially how changes API calls, resource usage, or
 impacts SLI/SLOs.

### Implementation Phases/History

Describe the development and implementation phases planned to break up the work and/or
 record them here as they occur. Provide enough detail so readers may track the major
 milestones in the lifecycle of the design proposal and correlate them with issues, PRs,
 and releases occurring within the project.
