# HubSpot 2026-09 API migration assessment

Status: comparison, live contract verification, and package implementation complete in the working tree. The two explicitly accepted live-verification risks remain Communication Preferences batch write and OAuth.

Assessment date: 2026-09-25

Implementation date: 2026-09-29

Target: HubSpot's current documented GA REST API release for each route family. Most families are `2026-09`; CRM Associations and CRM Lists are currently `2026-03`. Beta endpoints are out of scope.

Repository: `github.com/karman-digital/hubspot`

## Purpose

This document inventories the outbound HubSpot API surface in this package, compares each legacy route family with HubSpot's current documented dated equivalent, and records the resulting implementation. The comparison and live canaries were completed before package code was changed.

At assessment start, the repository contained 133 exported functions or methods in `methods.go` files. Some are local orchestration helpers or expose the same HTTP operation, so the tables below group them by remote route contract rather than treating each Go method as a distinct API.

## Classification

| Classification | Meaning |
| --- | --- |
| URL only | The method, parameters, request body, and response fields consumed by this package remain compatible. Additive response fields may be ignored safely. |
| Contract change | The HTTP method, parameter placement, request schema, response schema, or documented behaviour changes. Models and/or method logic must change. |
| Verify before implementation | HubSpot's current guide and current OpenAPI operation pages disagree or the exact operation is missing. Do not choose a route through inference. |
| No dated equivalent found | The Developer MCP returned only legacy or unversioned documentation. Keep the current endpoint until HubSpot publishes an equivalent or an intentional semantic replacement is approved. |
| Existing defect | Current package behaviour is already inconsistent with its method name and must not be hidden inside the version migration. |

## Pre-change baseline

- Working tree was clean before this assessment.
- `go test ./...` passes on commit `7d6170c`.
- The repository contains 136 source/test references to numeric HubSpot API paths.
- Before implementation, only associations, custom objects, invoices, pipelines, and properties had meaningful route assertions. Most migrated routes needed new URL and contract tests.

## Implementation record

The package now uses the current documented dated routes for every inventoried operation that has a dated equivalent. Most are `2026-09`; CRM Associations and CRM Lists are documented at `2026-03`. It deliberately retains Email Analytics v1, Files `stat/{filePath}`, and the unversioned GraphQL collector because no equivalent dated operation was found.

The contract migrations were implemented explicitly:

- Communication Preferences now uses dated definition and status action envelopes, the unified status write operation, and the dated batch route. The existing subscribe/unsubscribe helpers remain EMAIL convenience methods; they do not invent legal-basis values or enforce portal policy. A direct `SetCommunicationPreferenceStatus` method exposes the complete caller-controlled request and response.
- OAuth exchange and refresh now use the dated token route. Introspection is an authenticated form-body `POST`; the token is no longer placed in a URL. `Credentials.GetBearerTokenData` uses the instance's client credentials, and the standalone helper now requires client ID, client secret, and token-type hint.
- Notes batch operations now call Notes routes rather than Email routes.
- The schema-association reader added on `main` during implementation now uses the dated CRM Object Schemas route; its projected response fields remain compatible.
- Dated Settings user response fields, including `roleIds` and `seatNames`, are represented.
- The four unique-property helpers return an error when `idProperty` options are absent instead of panicking, and their path values are escaped.
- Campaign asset response-body debug logging was removed.

The package remains portal-agnostic. CRM property maps and Communication Preferences legal-basis fields are serialized from caller input without portal-specific defaults. HubSpot remains responsible for validating the target portal's configuration.

### Final post-merge documentation audit

On 2026-09-29, every implemented route family was checked again against HubSpot's current Developer documentation. That audit found two accepted-but-undocumented `2026-09` aliases in the initial implementation: generic CRM Associations and CRM Lists. Both are documented at `2026-03`, so the package was corrected to those exact routes. Invoice-specific association operations remain on their separately documented `2026-09` invoice routes. No other route-family mismatch was found.

## Executive findings

1. Most CRM v3 and v4 operations have dated equivalents. The new path is not a global substitution: the date segment moves to a resource-specific position.
2. Standard CRM read/update/delete/search and batch inputs used by this package remain compatible. Live writes exposed a portal-level create-behaviour change: the dated object API honours required-property configuration that the legacy v3 API did not enforce for companies, contacts, and deals in the test portal. This package must remain portal-agnostic: it should continue to accept arbitrary property maps, send them unchanged, and return HubSpot's validation response. The observed fields are live-canary fixtures and a downstream compatibility warning, not library validation rules. The `2026-09` batch response consolidates partial-error fields into the normal batch response. The existing `BatchResponseBase` already contains `errors` and `numErrors`, so this is additive for the package.
3. OAuth token exchange and refresh are URL-only changes, but token introspection is a real contract migration: `GET` with a token in the URL becomes authenticated `POST` form data and returns a new model.
4. Communication Preferences v3 methods are not URL-only. The current API uses a unified status endpoint, different request fields, new status values, and action-response wrappers.
5. HubSpot's Deals guide documents the `deals` slug while some `2026-09` operation specs emit object type ID `0-3`. The live portal accepts both on the exact dated route; use the guide's public `deals` slug.
6. The `2026-09` Campaigns guide documents campaign asset retrieval, but the current operation page named `get-assets` describes asset-type retrieval. Live collection, asset, and write canaries resolve that documentation inconsistency: `/marketing/campaigns/2026-09/{guid}/assets/{assetType}` is the working migration target.
7. No date-versioned equivalent was found for Email Analytics v1, Files `stat/{filePath}`, or the unversioned GraphQL collector endpoint.
8. Before implementation, `crm/engagements/notes/batch` called `/objects/emails/batch/...`. This was an existing defect, not a migration consequence; it is now fixed after the correct Notes route was proven live.

## Live portal verification

Verification date: 2026-09-25

The original live canaries used the dated routes available to the portal at the time. A final documentation audit after merge found that generic CRM Associations and CRM Lists are currently documented at `2026-03`, while the other dated CRM object families are documented at `2026-09`. HubSpot accepted undocumented aliases during the canaries, but accepted runtime aliases are not used as the package contract.

### Authentication and scopes

The supplied temporary private-app token was introspected against the live portal before testing. The following newly changed scopes propagated successfully during the run:

- `communication_preferences.read`
- `communication_preferences.write`
- `hubdb.rows.read`
- `hubdb.rows.write`
- `hubdb.tables.publish`
- `hubdb.tables.read`
- `hubdb.tables.write`
- `files.ui_hidden.write`
- `marketing.campaigns.read`
- `marketing.campaigns.write`
- `settings.users.read`
- `settings.users.write`

The user explicitly accepted the following unverified risks for this migration assessment:

| Operation | Verification status | Accepted risk |
| --- | --- | --- |
| Communication Preferences batch write, v4 and `2026-09` | Not run because `communication_preferences.statuses.batch.write` is unavailable | Treat the URL-only migration as accepted risk based on the matching v4 and dated documented contracts |
| OAuth token exchange, refresh, and dated introspection | Not run with an OAuth client and authorization-code flow | Treat the documented contracts as accepted risk until downstream OAuth credentials are available |

HubSpot's current Communication Preferences guide names the dated batch scope differently (`subscription-status-batch-write`) from the granular scope returned by the live API. Do not resolve that discrepancy by renaming a scope in code; the app configuration must be verified in HubSpot.

### CRM read contracts

Fresh legacy-versus-exact-target requests returned `200` and identical `results` for list operations on companies, contacts, deals, line items, products, calls, CRM emails, meetings, notes, tasks, and invoices. Fresh searches for companies, contacts, deals, and invoices also returned identical results. Batch reads returned identical results for companies, contacts, deals, line items, products, calls, CRM emails, and notes.

Single-object reads were also equal for the representative records tested. Both `/crm/objects/2026-09/deals/{id}` and `/crm/objects/2026-09/0-3/{id}` work on this portal, resolving the route gate in favour of the public `deals` slug used by the guide.

The package's current Notes batch route defect was reproduced live: sending a real note ID to the package's email batch-read family produced a partial-error response with no note result, while the legacy Notes route and exact dated Notes route each returned the note.

### CRM write contracts

The following operations were run against both the legacy family and the exact documented dated family, with property read-back after updates and `404` read-back after cleanup:

| Family | Operations proven live |
| --- | --- |
| Companies, contacts, deals | Single create/update/read/archive; batch create/update/read/archive |
| Custom objects | Single create/update/read/archive, writable unique-property read/update, batch create/update/read/archive, and batch upsert create/update |
| Line items | Single create/update/read and batch create/update/read/archive |
| Products | Read and batch create/update/read/archive |
| Calls and CRM emails | Single read and batch create/update/read/archive |
| Notes | Single create/read/archive and intended Notes batch create/update/read/archive |
| Tasks | Single create and archive |
| Invoices | Single create/update/read/search and archive |

Numeric JSON IDs in batch archive bodies were accepted by the live legacy and exact dated endpoints for every tested object family. The existing integer model therefore works at runtime on this portal, although it still disagrees with the OpenAPI string schema and remains a documented compatibility risk to revalidate if HubSpot changes its runtime acceptance.

The material write difference is required-property enforcement:

| Object | Minimal v3 create | Same dated create | Dated properties required by this portal |
| --- | --- | --- | --- |
| Company | `201` | `400` | `domain` |
| Contact | `201` | `400` | `email`, `lead_source`, `hs_legal_basis` |
| Deal | `201` | `400` | `amount`, `revenue_line`, `closedate`, `description`, `deal_source` |

Retries using values read from the portal's live property metadata succeeded so that the remaining canary operations could be exercised. Those values must not be hard-coded, defaulted, or validated by this package. The existing arbitrary property maps remain the provider-neutral boundary; any downstream caller audit belongs to the later application migration.

Other live write details:

- CRM email creation required the portal's `hs_email_direction` property on both versions. The retry used the live `DRAFT_EMAIL` enumeration value and succeeded.
- A custom object's `hasUniqueValue` flag is insufficient to prove upsert suitability: the `reviews` schema exposed only the read-only `hs_unique_creation_key`. The `relationships` schema exposed the writable unique `relationship_key`; unique reads, unique updates, and both create/update branches of batch upsert passed on both versions.
- Company lookup by `domain` returned `404` on both versions for a newly created canary. The same company helper succeeded on both versions with the portal's writable unique `xero_contact_id`; contact lookup by `email`, product lookup by `hs_sku`, and custom-object lookup by `relationship_key` also succeeded. No writable unique Deal property exists in this portal, so `GetDealByUniqueProperty` remains unproved rather than assumed.

### Associations and metadata

Single default and typed associations, batch default creation, batch typed creation, batch reads, and batch label archive all passed on v4 and dated association aliases. The final documentation audit corrected the implementation to HubSpot's documented `2026-03` association routes. Association result arrays can differ in order; they are equal after sorting their association types by ID.

Invoice-company default and typed association create/read/remove flows passed on both route families. The correct live invoice-to-company association type is `179`; the bidirectional default-association response also included reverse type `180`, which is not valid for an invoice-to-company typed request.

Property group create, property create/get/update, enumeration option replacement, property deletion, and group deletion all passed on legacy and exact dated routes. Both disposable properties and groups returned `404` after deletion. Pipeline, owner, list search/get/membership, and association read contracts were also verified read-only.

### Communication Preferences

Legacy definitions use `subscriptionDefinitions`; dated definitions use the action envelope with `results`. After ID normalization, all seven definition records and their fields matched.

The correct current status routes are `/communication-preferences/v4/statuses/{email}?channel=EMAIL` and `/communication-preferences/2026-09/statuses/{email}?channel=EMAIL`. The package's v3 status route still returned `200` during the canary, but its response contract is the legacy model and cannot be reused for the dated response without an adapter.

Single write behaviour was tested with two unique `example.com` addresses that were first proven not to be CRM contacts:

- Dated subscribe and unsubscribe both returned `200` without legal-basis fields.
- On this GDPR-enabled portal, legacy v3 subscribe and unsubscribe returned `400` without `legalBasis` and `legalBasisExplanation`; both succeeded when those documented fields were supplied.
- Final read-back is `NOT_SUBSCRIBED` for the legacy canary and `UNSUBSCRIBED` for the dated canary. HubSpot exposes no delete/reset operation for these preference histories.

This confirms that preserving the convenience method signatures is possible, but preserving the legacy response and error semantics is not. The implementation now exposes the exact dated response and keeps legal-basis fields caller-controlled; the portal behaviour described above was established live.

### Files, CMS, marketing, OAuth, settings, and GraphQL

- Files search and signed URLs returned compatible responses. Multipart upload and replacement passed on v3 and `2026-09`, including equal post-update sizes. URL imports on both versions returned `202`, reached `COMPLETE`, and produced files with the same MD5, 5,123-byte size, and 288-by-288 dimensions. Both the response-provided v3 task-status link and the documented `2026-09` task-status route returned identical payloads for the same dated task; the dated status path also returned `200` through the package's normal `api.hubapi.com` host. Use the documented dated route because the returned link still embeds `v3`. The two earlier private files and all URL-import canaries were deleted with `204` and returned `404` on read-back. Final v3 and dated searches found zero `codex-api` files.
- Blog posts, blog tags list/batch, and HubDB table-row reads passed. The exact HubDB OpenAPI route `/cms/hubdb/2026-09/...` works; the guide's `/cms/v3/hubdb/2026-09/...` variant also works but should not be selected over the operation spec.
- Marketing email list/single reads passed. Email Analytics collection, event, and individual campaign reads still work only on the legacy family. Campaign v3 and dated collections returned the same ten records after removing the versioned paging link. `MARKETING_EMAIL` asset reads returned the same eight assets and fields after sorting by ID; result order differed and must not be treated as stable. A disposable campaign created on `2026-09` was patched through both v3 and dated routes, and each distinct value was confirmed by read-back. It was then deleted with `204`, after which both versions returned `404` for its UUID.
- OAuth v1 PAT introspection rejected the private-app token format. Dated introspection requires client ID and client secret, so a PAT alone cannot prove that migration contract; the user accepted OAuth as an unverified risk for now.
- `admin@karman.digital` was absent from all 35 live users before the canary. With `sendWelcomeEmail: false`, both v3 and dated create calls returned `201` and identical response payloads, and both versions could read the resulting numeric user ID. Each created state was deleted before the next create; the final dated delete returned `204`, both versions returned `404` for the ID, and the final user list contained zero exact-email matches. Existing-user list payloads differed only because `2026-09` adds `seatNames`.
- GraphQL collector introspection returned `200`; it remains unversioned.

### Cleanup state

Every CRM record, association, custom property, custom property group, campaign, file, and Settings user created by the canaries was removed and verified by read-back. Final marker searches found no canary files, and `admin@karman.digital` is absent from the final user list. Two fake, non-contact Communication Preferences addresses remain in an unsubscribed state because HubSpot exposes no delete/reset operation; no deletable canary state remains.

## Complete package inventory

In the comparison tables below, “Current” means the pre-migration implementation captured during assessment and “Target” is the route now implemented unless the row explicitly says no dated equivalent was found.

### CRM object APIs

Unless a row says otherwise, all listed object methods use the same public request/response models in `hubspot/api/models/crm`.

| Package surface | Go methods covered | Pre-migration route family | Implemented target | Contract result |
| --- | --- | --- | --- | --- |
| Companies | `CreateCompany`, `UpdateCompany`, `GetCompany`, `GetCompanyByUniqueProperty`, `GetCompanies`, `SearchCompanies`, `DeleteCompany`; four `Batch*` methods | `/crm/v3/objects/companies...` | `/crm/objects/2026-09/companies...` | Package contract remains arbitrary-property pass-through. The test portal's dated API required `domain`; this is downstream portal configuration, not a library rule |
| Contacts | `CreateContact`, `UpdateContact`, `GetContact`, `GetContactByUniqueProperty`, `SearchContacts`, `DeleteContact`; four `Batch*` methods | `/crm/v3/objects/contacts...` | `/crm/objects/2026-09/contacts...` | Package contract remains arbitrary-property pass-through. The test portal's dated API required `email`, `lead_source`, and `hs_legal_basis`; this is downstream portal configuration, not a library rule |
| Deals | `CreateDeal`, `UpdateDeal`, `GetDeal`, `GetDealByUniqueProperty`, `GetDeals`, `SearchDeals`, `DeleteDeal`; four `Batch*` methods | `/crm/v3/objects/deals...` | `/crm/objects/2026-09/deals...`; live canary also accepted object type ID `0-3` | Package contract remains arbitrary-property pass-through. The test portal's dated API required five properties; this is downstream portal configuration, not a library rule. Live route gate resolved in favour of `deals` |
| Custom objects | Eight `CustomObjectService` methods and five `BatchCustomObjectService` methods, including upsert | `/crm/v3/objects/{objectType}...` | `/crm/objects/2026-09/{objectType}...` | URL only; additive batch error fields already represented |
| Line items | `CreateLineItem`, `UpdateLineItem`, `GetLineItem`; four `Batch*` methods | `/crm/v3/objects/line_items...` | `/crm/objects/2026-09/line_items...` | URL only |
| Products | `GetProductByUniqueId`; four `Batch*` methods | `/crm/v3/objects/products...` | `/crm/objects/2026-09/products...` | URL only |
| Calls | `GetCall`; four `Batch*` methods | `/crm/v3/objects/calls...` | `/crm/objects/2026-09/calls...` | URL only |
| CRM emails | `GetEmail`; four `Batch*` methods | `/crm/v3/objects/emails...` | `/crm/objects/2026-09/emails...` | URL only |
| Meetings | `GetMeeting` | `/crm/v3/objects/meetings/{id}` | `/crm/objects/2026-09/meetings/{id}` | URL only |
| Notes | `CreateNoteWithAssociations`, `GetNote` | `/crm/v3/objects/notes...` | `/crm/objects/2026-09/notes...` | URL only |
| Notes batch | Four `BatchNotesService` methods | **Called `/crm/v3/objects/emails/batch/...`** | `/crm/objects/2026-09/notes/batch/...` | Existing defect repaired after live reproduction against Notes and Email routes |
| Tasks | `CreateTaskWithAssociations` | `/crm/v3/objects/tasks` | `/crm/objects/2026-09/tasks` | URL only |
| Invoices | `CreateInvoice`, `UpdateInvoice`, `GetInvoice`, `SearchInvoices` | `/crm/v3/objects/invoices...` | `/crm/objects/2026-09/invoices...` | URL only |

The generic CRM route transformations required by those rows are:

| Operation | Current | Target |
| --- | --- | --- |
| Create/list | `/crm/v3/objects/{objectType}` | `/crm/objects/2026-09/{objectType}` |
| Read/update/archive | `/crm/v3/objects/{objectType}/{objectId}` | `/crm/objects/2026-09/{objectType}/{objectId}` |
| Search | `/crm/v3/objects/{objectType}/search` | `/crm/objects/2026-09/{objectType}/search` |
| Batch create/update/read/archive/upsert | `/crm/v3/objects/{objectType}/batch/{action}` | `/crm/objects/2026-09/{objectType}/batch/{action}` |

Contract notes:

- Single-object get, update, archive, list, and search retain the fields the package uses. Create request shapes remain compatible, but the dated API enforces portal-configured required properties that v3 may allow callers to omit.
- `2026-09` permits nullable values in some returned `properties` maps. The package uses `map[string]any`, so decoding remains compatible.
- The dated batch response includes `errors` and `numErrors` on the primary response schema. `sharedmodels.BatchResponseBase` already has both fields.
- Existing package batch-delete IDs are integers, while HubSpot's OpenAPI schema represents object IDs as strings. This mismatch predates the migration and should be corrected or compatibility-tested separately.

### CRM object schemas

| Go methods covered | Current | Target | Contract result |
| --- | --- | --- | --- |
| `GetSchema` | `/crm/v3/schemas/{objectType}` | `/crm-object-schemas/2026-09/schemas/{objectType}` | URL only for the package's projected `objectTypeId` and association fields; the path value is escaped |

### CRM associations

| Go methods covered | Current | Target | Contract result |
| --- | --- | --- | --- |
| `CreateDefaultAssociation` | `PUT /crm/v4/objects/{fromType}/{fromId}/associations/default/{toType}/{toId}` | `PUT /crm/objects/2026-03/{fromType}/{fromId}/associations/default/{toType}/{toId}` | URL only |
| `GetAssociations` | `GET /crm/v4/objects/{fromType}/{fromId}/associations/{toType}` | `GET /crm/objects/2026-03/{fromType}/{fromId}/associations/{toType}` | URL only |
| `BatchCreateDefaultAssociations` | `POST /crm/v4/associations/{fromType}/{toType}/batch/associate/default` | `POST /crm/associations/2026-03/{fromType}/{toType}/batch/associate/default` | URL only |
| `BatchGetAssociations`, `BatchGetAllAssociations` | `POST /crm/v4/associations/{fromType}/{toType}/batch/read` | `POST /crm/associations/2026-03/{fromType}/{toType}/batch/read` | URL only; dated API documents up to 1,000 source IDs per read |
| `BatchCreateAssociations` | `POST /crm/v4/associations/{fromType}/{toType}/batch/create` | `POST /crm/associations/2026-03/{fromType}/{toType}/batch/create` | URL only |
| `CreateAssociation` | `PUT /crm/v4/objects/{fromType}/{fromId}/associations/{toType}/{toId}` | `PUT /crm/objects/2026-03/{fromType}/{fromId}/associations/{toType}/{toId}` | URL only |
| `BatchArchiveAssociationLabels` | `POST /crm/v4/associations/{fromType}/{toType}/batch/labels/archive` | `POST /crm/associations/2026-03/{fromType}/{toType}/batch/labels/archive` | URL only |
| Six invoice association methods | The same `/crm/v4/objects/...` association families | The corresponding `/crm/objects/2026-09/...` families | URL only |

### CRM metadata, owners, and lists

| Package surface | Current | Target | Contract result |
| --- | --- | --- | --- |
| Properties: group create, property create/get/update | `/crm/v3/properties/{objectType}...` | `/crm/properties/2026-09/{objectType}...` | URL only for package fields |
| Pipelines: `GetPipelines` | `/crm/v3/pipelines/{objectType}` | `/crm/pipelines/2026-09/{objectType}` | URL only |
| Owners: `GetOwners`, `GetAllOwners`, `GetOwner` | `/crm/v3/owners...` | `/crm/owners/2026-09...` | URL only |
| Lists: `SearchLists` | `/crm/v3/lists/search` | `/crm/lists/2026-03/search` | URL only |
| Lists: `GetLists` | `/crm/v3/lists/?listIds=...` | `/crm/lists/2026-03?listIds=...` | URL only; remove the package's unnecessary trailing slash |
| Lists: `GetListMemberships` | `/crm/v3/lists/{listId}/memberships` | `/crm/lists/2026-03/{listId}/memberships` | URL only |

### OAuth and credentials

| Go methods covered | Current | Target | Contract result |
| --- | --- | --- | --- |
| `GenerateTokenPair`, `RefreshTokenPair` | `POST /oauth/v1/token` | `POST /oauth/2026-09/token` | URL only for form inputs and existing return fields; response adds `token_use`, scopes, and IDs |
| `ValidateBearerToken`, `GetBearerTokenData` | `GET /oauth/v1/access-tokens/{token}` | `POST /oauth/2026-09/token/introspect` | Contract change |

Introspection migration requirements:

- Move the token out of the URL and into an `application/x-www-form-urlencoded` body.
- Send `client_id`, `client_secret`, `token`, and `token_type_hint`.
- Replace the legacy `BearerTokenBody` mapping with the dated access/refresh token response, including `active`, `client_id`, `token_use`, and `is_private_distribution`.
- `GetBearerTokenData(bearerToken string)` cannot call the new endpoint with its current inputs because it has no client credentials. Its API must change or it must become a method on `Credentials`.
- Error handling should use OAuth `error` and `error_description` while retaining HubSpot fields for diagnostics.

### Communication Preferences

| Go methods covered | Current | Target | Contract result |
| --- | --- | --- | --- |
| `GetCommunicationPreferences` | `GET /communication-preferences/v3/definitions` | `GET /communication-preferences/2026-09/definitions` | Contract change: outer response changes from `subscriptionDefinitions` to an action response containing `results` |
| `SubscribeToCommunicationPreference`, `UnsubscribeFromCommunicationPreference` | Separate `POST /communication-preferences/v3/subscribe` and `/unsubscribe`; email is in the body; `subscriptionId` is a string | Unified `POST /communication-preferences/2026-09/statuses/{subscriberIdString}`; body requires `channel`, `statusState`, and integer `subscriptionId` | Contract and behaviour change |
| `GetCommunicationPreferenceStatus` | `GET /communication-preferences/v3/status/email/{email}` | `GET /communication-preferences/2026-09/statuses/{subscriberIdString}?channel=EMAIL` | Contract change: action response and new public-status model |
| `BatchUpdateCommunicationPreferences` | `POST /communication-preferences/v4/statuses/batch/write` | `POST /communication-preferences/2026-09/statuses/batch/write` | URL only; current batch model already uses the v4/current field layout |

Implementation decisions:

- The existing single-update method signatures remain as EMAIL convenience wrappers that derive `statusState=SUBSCRIBED` or `UNSUBSCRIBED` internally.
- Exact dated response models replace the incompatible legacy response shape rather than fabricating a lossy adapter. Downstream callers must migrate to `results` and the dated status fields.
- `SetCommunicationPreferenceStatus` returns the full action response, including documented successful no-op reasons. The convenience wrappers continue to return only success or failure for source compatibility.

### Files

| Go methods covered | Current | Target | Contract result |
| --- | --- | --- | --- |
| `ImportFileViaUrl` and its status polling | `POST /files/v3/files/import-from-url/async`; `GET /files/v3/files/import-from-url/async/tasks/{taskId}/status` | `POST /files/2026-09/files/import-from-url/async`; `GET /files/2026-09/files/import-from-url/async/tasks/{taskId}/status` | URL only; both dated operations passed live. The dated POST currently returns a status link containing `v3`, but the documented dated status route also returned `200` with an identical payload |
| `UploadFile` | `POST /files/v3/files` | `POST /files/2026-09/files` | URL only according to the current Files guide; multipart upload passed live |
| `UpdateFile` | `PUT /files/v3/files/{fileId}` | `PUT /files/2026-09/files/{fileId}` | URL only according to the current Files guide; multipart replacement passed live |
| `GetSignedUrl` | `GET /files/v3/files/{fileId}/signed-url` | `GET /files/2026-09/files/{fileId}/signed-url` | URL only according to the current Files guide |
| `GetFileByPath` | `GET /files/v3/files/stat/{encodedPath}` | No exact dated equivalent found | No dated equivalent found; do not silently replace it with search because that changes semantics and response shape |

The current Files guide documents upload, replacement, and signed URLs, but individual `2026-09` OpenAPI operation pages were not returned for all three. Live canaries passed for each package operation that has a dated target.

### CMS

| Go methods covered | Current | Target | Contract result |
| --- | --- | --- | --- |
| `GetAllBlogPosts` | `GET /cms/v3/blogs/posts` | `GET /cms/blogs/2026-09/posts` | URL only; filters, sort, state, limit, and cursor remain supported |
| `GetBatchBlogTags` | `POST /cms/v3/blogs/tags/batch/read` | `POST /cms/blogs/2026-09/tags/batch/read` | URL only; request and response schema references are unchanged |
| `GetTableRow` | `GET /cms/v3/hubdb/tables/{table}/rows/{row}` | `GET /cms/hubdb/2026-09/tables/{table}/rows/{row}` | URL only based on the exact operation spec |

HubDB documentation note: the current guide includes examples containing `/cms/v3/hubdb/2026-09/...`, while the exact `2026-09` OpenAPI operation is `/cms/hubdb/2026-09/...`. The exact OpenAPI route passed the live read canary and is implemented.

### Marketing

| Package surface | Current | Target | Contract result |
| --- | --- | --- | --- |
| Campaigns: `GetCampaigns` | `GET /marketing/v3/campaigns` | `GET /marketing/campaigns/2026-09` | URL only; collection and campaign schemas used by the package are structurally compatible |
| Campaigns: `PatchCampaign` | `PATCH /marketing/v3/campaigns/{guid}` | `PATCH /marketing/campaigns/2026-09/{guid}` | URL only; input and response schema remain compatible |
| Campaigns: `GetCampaignAssets` | `GET /marketing/v3/campaigns/{guid}/assets/{assetType}` | `GET /marketing/campaigns/2026-09/{guid}/assets/{assetType}` | URL only, proven live against eight `MARKETING_EMAIL` assets; response ordering differed, but records were identical after sorting by ID |
| Marketing email: `GetMarketingEmail` | `GET /marketing/v3/emails/{emailId}` | `GET /marketing/emails/2026-09/{emailId}` | URL only; the `PublicEmail` response fields consumed by the package remain compatible |
| Email Analytics: `GetEmailEvents`, `GetEmailCampaigns`, `GetEmailCampaign` | `/email/public/v1/events`, `/email/public/v1/campaigns...` | No dated Email Analytics equivalent found | No dated equivalent found; the current Events API explicitly excludes marketing email events |

### Settings and GraphQL

| Package surface | Current | Target | Contract result |
| --- | --- | --- | --- |
| Users: `Create` | `POST /settings/v3/users` | `POST /settings/users/2026-09` | URL only for the request and canary response. Additive dated fields `roleIds` and `seatNames` are represented in the package model |
| GraphQL: `MakeRequest`, `MakeRequestWithFullResponse` | `POST /collector/graphql` | Still `/collector/graphql` | Unversioned; no change |

## Package risks discovered during inventory

These are not reasons to broaden the migration silently, but they affect safe delivery:

1. **Resolved:** Notes batch now uses the dated Notes routes verified during the live comparison.
2. **Retained and live compatibility-tested:** `BatchDeleteBody` uses integer IDs while the dated OpenAPI schema uses string IDs. Both legacy and dated live endpoints accepted the existing integer representation across every tested object family.
3. **Resolved:** unconditional Campaign asset request and response-body debug logging was removed.
4. **Covered for this migration:** several existing clients still build absolute URLs. Their migrated routes were included in the source inventory and exercised by the live canaries where credentials and scopes allowed.
5. **Resolved through evidence rather than mocked endpoint tests:** the implementation follows the documented contracts and completed live canaries. The accepted Communication Preferences batch-write and OAuth live-proof gaps remain explicit.
6. **Resolved:** all four unique-property helpers now check for absent options and return the documented error rather than panicking.

## Migration gates arising from the comparison

1. **Resolved:** both dated Deal identifiers work live. Use the documented public `deals` slug rather than the numeric implementation ID.
2. **Resolved:** the documented `2026-09` campaign-assets route returned the same eight assets as v3 after sorting by ID. Do not rely on result order.
3. **Resolved:** explicit dated Communication Preferences models were introduced. This is intentionally source-breaking where the legacy response cannot represent the dated contract losslessly.
4. **Implemented with accepted live-verification risk:** token introspection is now a method on `Credentials`; the standalone helper requires client credentials and a token-type hint. It awaits a real OAuth client flow.
5. **Resolved:** the Notes batch defect was repaired after being reproduced against the live Notes and Email route families.
6. **Deferred to downstream migration:** audit each consuming application's company, contact, and deal creates against its own target portal. Do not add portal-specific required fields or validation to this package.
7. **Resolved:** hidden-file write propagated; all prior and new Files canaries were deleted and verified absent. URL import and both versions of status polling passed live.
8. **Resolved for package implementation:** the documented comparison, source inventory, and completed live canaries support the migration. Downstream function migration remains a separate phase.

## Authoritative documentation used

- [2026-09 API reference](https://developers.hubspot.com/docs/api-reference/latest/overview)
- [Legacy API migration guide](https://developers.hubspot.com/docs/api-reference/legacy/migration-guide)
- [Developer platform and API versioning](https://developers.hubspot.com/docs/developer-tooling/platform/versioning)
- [CRM Search API](https://developers.hubspot.com/docs/api-reference/latest/crm/search-the-crm)
- [CRM Associations 2026-03 guide](https://developers.hubspot.com/docs/api-reference/2026-03/crm/associations/associate-records/guide)
- [CRM Lists 2026-03 guide](https://developers.hubspot.com/docs/api-reference/2026-03/crm/lists/guide)
- [OAuth v1 migration guide](https://developers.hubspot.com/docs/api-reference/legacy/authentication/oauth-tokens/v1/migration-guide)
- [Communication Preferences guide](https://developers.hubspot.com/docs/api-reference/latest/communication-preferences/guide)
- [Files API guide](https://developers.hubspot.com/docs/api-reference/latest/files/guide)
- [Campaigns API guide](https://developers.hubspot.com/docs/api-reference/latest/marketing/campaigns/guide)
- [Settings user provisioning guide](https://developers.hubspot.com/docs/api-reference/latest/account/settings/user-provisioning/guide)
- [Marketing Email API guide](https://developers.hubspot.com/docs/api-reference/latest/marketing/marketing-emails/guide)
- [Blog Posts API guide](https://developers.hubspot.com/docs/api-reference/latest/cms/blogs/posts/guide)
- [Blog Tags API guide](https://developers.hubspot.com/docs/api-reference/latest/cms/blogs/tags/guide)
- [HubDB API guide](https://developers.hubspot.com/docs/api-reference/latest/cms/hubdb/guide)
- [Legacy Email Events API guide](https://developers.hubspot.com/docs/api-reference/legacy/reporting/email-analytics/guide)
