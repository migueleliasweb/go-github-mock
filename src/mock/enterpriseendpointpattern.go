package mock

// Code generated; DO NOT EDIT.

var GetEnterpriseSlash EndpointPattern = EndpointPattern{
	Pattern: "/",
	Method:  "GET",
}

var GetEnterpriseAdvisories EndpointPattern = EndpointPattern{
	Pattern: "/advisories",
	Method:  "GET",
}

var GetEnterpriseAdvisoriesByGhsaId EndpointPattern = EndpointPattern{
	Pattern: "/advisories/{ghsa_id}",
	Method:  "GET",
}

var GetEnterpriseApp EndpointPattern = EndpointPattern{
	Pattern: "/app",
	Method:  "GET",
}

var PostEnterpriseAppManifestsConversionsByCode EndpointPattern = EndpointPattern{
	Pattern: "/app-manifests/{code}/conversions",
	Method:  "POST",
}

var GetEnterpriseAppHookConfig EndpointPattern = EndpointPattern{
	Pattern: "/app/hook/config",
	Method:  "GET",
}

var PatchEnterpriseAppHookConfig EndpointPattern = EndpointPattern{
	Pattern: "/app/hook/config",
	Method:  "PATCH",
}

var GetEnterpriseAppHookDeliveries EndpointPattern = EndpointPattern{
	Pattern: "/app/hook/deliveries",
	Method:  "GET",
}

var GetEnterpriseAppHookDeliveriesByDeliveryId EndpointPattern = EndpointPattern{
	Pattern: "/app/hook/deliveries/{delivery_id}",
	Method:  "GET",
}

var PostEnterpriseAppHookDeliveriesAttemptsByDeliveryId EndpointPattern = EndpointPattern{
	Pattern: "/app/hook/deliveries/{delivery_id}/attempts",
	Method:  "POST",
}

var GetEnterpriseAppInstallationRequests EndpointPattern = EndpointPattern{
	Pattern: "/app/installation-requests",
	Method:  "GET",
}

var GetEnterpriseAppInstallations EndpointPattern = EndpointPattern{
	Pattern: "/app/installations",
	Method:  "GET",
}

var GetEnterpriseAppInstallationsByInstallationId EndpointPattern = EndpointPattern{
	Pattern: "/app/installations/{installation_id}",
	Method:  "GET",
}

var DeleteEnterpriseAppInstallationsByInstallationId EndpointPattern = EndpointPattern{
	Pattern: "/app/installations/{installation_id}",
	Method:  "DELETE",
}

var PostEnterpriseAppInstallationsAccessTokensByInstallationId EndpointPattern = EndpointPattern{
	Pattern: "/app/installations/{installation_id}/access_tokens",
	Method:  "POST",
}

var PutEnterpriseAppInstallationsSuspendedByInstallationId EndpointPattern = EndpointPattern{
	Pattern: "/app/installations/{installation_id}/suspended",
	Method:  "PUT",
}

var DeleteEnterpriseAppInstallationsSuspendedByInstallationId EndpointPattern = EndpointPattern{
	Pattern: "/app/installations/{installation_id}/suspended",
	Method:  "DELETE",
}

var DeleteEnterpriseApplicationsGrantByClientId EndpointPattern = EndpointPattern{
	Pattern: "/applications/{client_id}/grant",
	Method:  "DELETE",
}

var PostEnterpriseApplicationsTokenByClientId EndpointPattern = EndpointPattern{
	Pattern: "/applications/{client_id}/token",
	Method:  "POST",
}

var PatchEnterpriseApplicationsTokenByClientId EndpointPattern = EndpointPattern{
	Pattern: "/applications/{client_id}/token",
	Method:  "PATCH",
}

var DeleteEnterpriseApplicationsTokenByClientId EndpointPattern = EndpointPattern{
	Pattern: "/applications/{client_id}/token",
	Method:  "DELETE",
}

var PostEnterpriseApplicationsTokenScopedByClientId EndpointPattern = EndpointPattern{
	Pattern: "/applications/{client_id}/token/scoped",
	Method:  "POST",
}

var GetEnterpriseAppsByAppSlug EndpointPattern = EndpointPattern{
	Pattern: "/apps/{app_slug}",
	Method:  "GET",
}

var GetEnterpriseAssignmentsByAssignmentId EndpointPattern = EndpointPattern{
	Pattern: "/assignments/{assignment_id}",
	Method:  "GET",
}

var GetEnterpriseAssignmentsAcceptedAssignmentsByAssignmentId EndpointPattern = EndpointPattern{
	Pattern: "/assignments/{assignment_id}/accepted_assignments",
	Method:  "GET",
}

var GetEnterpriseAssignmentsGradesByAssignmentId EndpointPattern = EndpointPattern{
	Pattern: "/assignments/{assignment_id}/grades",
	Method:  "GET",
}

var GetEnterpriseClassrooms EndpointPattern = EndpointPattern{
	Pattern: "/classrooms",
	Method:  "GET",
}

var GetEnterpriseClassroomsByClassroomId EndpointPattern = EndpointPattern{
	Pattern: "/classrooms/{classroom_id}",
	Method:  "GET",
}

var GetEnterpriseClassroomsAssignmentsByClassroomId EndpointPattern = EndpointPattern{
	Pattern: "/classrooms/{classroom_id}/assignments",
	Method:  "GET",
}

var GetEnterpriseCodesOfConduct EndpointPattern = EndpointPattern{
	Pattern: "/codes_of_conduct",
	Method:  "GET",
}

var GetEnterpriseCodesOfConductByKey EndpointPattern = EndpointPattern{
	Pattern: "/codes_of_conduct/{key}",
	Method:  "GET",
}

var GetEnterpriseEmojis EndpointPattern = EndpointPattern{
	Pattern: "/emojis",
	Method:  "GET",
}

var GetEnterpriseInstallationServerStatisticsByEnterpriseOrOrg EndpointPattern = EndpointPattern{
	Pattern: "/enterprise-installation/{enterprise_or_org}/server-statistics",
	Method:  "GET",
}

var GetEnterpriseActionsCacheUsageByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/cache/usage",
	Method:  "GET",
}

var PutEnterpriseActionsOidcCustomizationIssuerByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/oidc/customization/issuer",
	Method:  "PUT",
}

var GetEnterpriseActionsPermissionsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions",
	Method:  "GET",
}

var PutEnterpriseActionsPermissionsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions",
	Method:  "PUT",
}

var GetEnterpriseActionsPermissionsOrganizationsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/organizations",
	Method:  "GET",
}

var PutEnterpriseActionsPermissionsOrganizationsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/organizations",
	Method:  "PUT",
}

var PutEnterpriseActionsPermissionsOrganizationsByEnterpriseByOrgId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/organizations/{org_id}",
	Method:  "PUT",
}

var DeleteEnterpriseActionsPermissionsOrganizationsByEnterpriseByOrgId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/organizations/{org_id}",
	Method:  "DELETE",
}

var GetEnterpriseActionsPermissionsSelectedActionsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/selected-actions",
	Method:  "GET",
}

var PutEnterpriseActionsPermissionsSelectedActionsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/selected-actions",
	Method:  "PUT",
}

var GetEnterpriseActionsPermissionsWorkflowByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/workflow",
	Method:  "GET",
}

var PutEnterpriseActionsPermissionsWorkflowByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/permissions/workflow",
	Method:  "PUT",
}

var GetEnterpriseActionsRunnerGroupsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups",
	Method:  "GET",
}

var PostEnterpriseActionsRunnerGroupsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups",
	Method:  "POST",
}

var GetEnterpriseActionsRunnerGroupsByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}",
	Method:  "GET",
}

var PatchEnterpriseActionsRunnerGroupsByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseActionsRunnerGroupsByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}",
	Method:  "DELETE",
}

var GetEnterpriseActionsRunnerGroupsOrganizationsByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations",
	Method:  "GET",
}

var PutEnterpriseActionsRunnerGroupsOrganizationsByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations",
	Method:  "PUT",
}

var PutEnterpriseActionsRunnerGroupsOrganizationsByEnterpriseByRunnerGroupIdByOrgId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}",
	Method:  "PUT",
}

var DeleteEnterpriseActionsRunnerGroupsOrganizationsByEnterpriseByRunnerGroupIdByOrgId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}",
	Method:  "DELETE",
}

var GetEnterpriseActionsRunnerGroupsRunnersByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners",
	Method:  "GET",
}

var PutEnterpriseActionsRunnerGroupsRunnersByEnterpriseByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners",
	Method:  "PUT",
}

var PutEnterpriseActionsRunnerGroupsRunnersByEnterpriseByRunnerGroupIdByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}",
	Method:  "PUT",
}

var DeleteEnterpriseActionsRunnerGroupsRunnersByEnterpriseByRunnerGroupIdByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}",
	Method:  "DELETE",
}

var GetEnterpriseActionsRunnersByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners",
	Method:  "GET",
}

var GetEnterpriseActionsRunnersDownloadsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/downloads",
	Method:  "GET",
}

var PostEnterpriseActionsRunnersGenerateJitconfigByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/generate-jitconfig",
	Method:  "POST",
}

var PostEnterpriseActionsRunnersRegistrationTokenByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/registration-token",
	Method:  "POST",
}

var PostEnterpriseActionsRunnersRemoveTokenByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/remove-token",
	Method:  "POST",
}

var GetEnterpriseActionsRunnersByEnterpriseByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}",
	Method:  "GET",
}

var DeleteEnterpriseActionsRunnersByEnterpriseByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}",
	Method:  "DELETE",
}

var GetEnterpriseActionsRunnersLabelsByEnterpriseByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}/labels",
	Method:  "GET",
}

var PostEnterpriseActionsRunnersLabelsByEnterpriseByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}/labels",
	Method:  "POST",
}

var PutEnterpriseActionsRunnersLabelsByEnterpriseByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}/labels",
	Method:  "PUT",
}

var DeleteEnterpriseActionsRunnersLabelsByEnterpriseByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}/labels",
	Method:  "DELETE",
}

var DeleteEnterpriseActionsRunnersLabelsByEnterpriseByRunnerIdByName EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/actions/runners/{runner_id}/labels/{name}",
	Method:  "DELETE",
}

var GetEnterpriseAnnouncementByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/announcement",
	Method:  "GET",
}

var PatchEnterpriseAnnouncementByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/announcement",
	Method:  "PATCH",
}

var DeleteEnterpriseAnnouncementByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/announcement",
	Method:  "DELETE",
}

var GetEnterpriseAuditLogByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/audit-log",
	Method:  "GET",
}

var GetEnterpriseCodeScanningAlertsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/code-scanning/alerts",
	Method:  "GET",
}

var GetEnterpriseCodeSecurityAndAnalysisByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/code_security_and_analysis",
	Method:  "GET",
}

var PatchEnterpriseCodeSecurityAndAnalysisByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/code_security_and_analysis",
	Method:  "PATCH",
}

var GetEnterpriseConsumedLicensesByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/consumed-licenses",
	Method:  "GET",
}

var GetEnterpriseCopilotBillingSeatsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/copilot/billing/seats",
	Method:  "GET",
}

var GetEnterpriseCopilotUsageByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/copilot/usage",
	Method:  "GET",
}

var GetEnterpriseDependabotAlertsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/dependabot/alerts",
	Method:  "GET",
}

var GetEnterpriseLicenseSyncStatusByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/license-sync-status",
	Method:  "GET",
}

var GetEnterpriseSecretScanningAlertsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/secret-scanning/alerts",
	Method:  "GET",
}

var GetEnterpriseSettingsBillingActionsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/actions",
	Method:  "GET",
}

var GetEnterpriseSettingsBillingAdvancedSecurityByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/advanced-security",
	Method:  "GET",
}

var GetEnterpriseSettingsBillingCostCentersByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/cost-centers",
	Method:  "GET",
}

var PostEnterpriseSettingsBillingCostCentersResourceByEnterpriseByCostCenterId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}/resource",
	Method:  "POST",
}

var DeleteEnterpriseSettingsBillingCostCentersResourceByEnterpriseByCostCenterId EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/cost-centers/{cost_center_id}/resource",
	Method:  "DELETE",
}

var GetEnterpriseSettingsBillingPackagesByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/packages",
	Method:  "GET",
}

var GetEnterpriseSettingsBillingSharedStorageByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/shared-storage",
	Method:  "GET",
}

var GetEnterpriseSettingsBillingUsageByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/settings/billing/usage",
	Method:  "GET",
}

var GetEnterpriseTeamCopilotUsageByEnterpriseByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/team/{team_slug}/copilot/usage",
	Method:  "GET",
}

var PostEnterpriseByEnterpriseBySecurityProductByEnablement EndpointPattern = EndpointPattern{
	Pattern: "/enterprises/{enterprise}/{security_product}/{enablement}",
	Method:  "POST",
}

var GetEnterpriseEvents EndpointPattern = EndpointPattern{
	Pattern: "/events",
	Method:  "GET",
}

var GetEnterpriseFeeds EndpointPattern = EndpointPattern{
	Pattern: "/feeds",
	Method:  "GET",
}

var GetEnterpriseGists EndpointPattern = EndpointPattern{
	Pattern: "/gists",
	Method:  "GET",
}

var PostEnterpriseGists EndpointPattern = EndpointPattern{
	Pattern: "/gists",
	Method:  "POST",
}

var GetEnterpriseGistsPublic EndpointPattern = EndpointPattern{
	Pattern: "/gists/public",
	Method:  "GET",
}

var GetEnterpriseGistsStarred EndpointPattern = EndpointPattern{
	Pattern: "/gists/starred",
	Method:  "GET",
}

var GetEnterpriseGistsByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}",
	Method:  "GET",
}

var PatchEnterpriseGistsByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseGistsByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}",
	Method:  "DELETE",
}

var GetEnterpriseGistsCommentsByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/comments",
	Method:  "GET",
}

var PostEnterpriseGistsCommentsByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/comments",
	Method:  "POST",
}

var GetEnterpriseGistsCommentsByGistIdByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/comments/{comment_id}",
	Method:  "GET",
}

var PatchEnterpriseGistsCommentsByGistIdByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/comments/{comment_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseGistsCommentsByGistIdByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/comments/{comment_id}",
	Method:  "DELETE",
}

var GetEnterpriseGistsCommitsByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/commits",
	Method:  "GET",
}

var GetEnterpriseGistsForksByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/forks",
	Method:  "GET",
}

var PostEnterpriseGistsForksByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/forks",
	Method:  "POST",
}

var GetEnterpriseGistsStarByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/star",
	Method:  "GET",
}

var PutEnterpriseGistsStarByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/star",
	Method:  "PUT",
}

var DeleteEnterpriseGistsStarByGistId EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/star",
	Method:  "DELETE",
}

var GetEnterpriseGistsByGistIdBySha EndpointPattern = EndpointPattern{
	Pattern: "/gists/{gist_id}/{sha}",
	Method:  "GET",
}

var GetEnterpriseGitignoreTemplates EndpointPattern = EndpointPattern{
	Pattern: "/gitignore/templates",
	Method:  "GET",
}

var GetEnterpriseGitignoreTemplatesByName EndpointPattern = EndpointPattern{
	Pattern: "/gitignore/templates/{name}",
	Method:  "GET",
}

var GetEnterpriseInstallationRepositories EndpointPattern = EndpointPattern{
	Pattern: "/installation/repositories",
	Method:  "GET",
}

var DeleteEnterpriseInstallationToken EndpointPattern = EndpointPattern{
	Pattern: "/installation/token",
	Method:  "DELETE",
}

var GetEnterpriseIssues EndpointPattern = EndpointPattern{
	Pattern: "/issues",
	Method:  "GET",
}

var GetEnterpriseLicenses EndpointPattern = EndpointPattern{
	Pattern: "/licenses",
	Method:  "GET",
}

var GetEnterpriseLicensesByLicense EndpointPattern = EndpointPattern{
	Pattern: "/licenses/{license}",
	Method:  "GET",
}

var PostEnterpriseMarkdown EndpointPattern = EndpointPattern{
	Pattern: "/markdown",
	Method:  "POST",
}

var PostEnterpriseMarkdownRaw EndpointPattern = EndpointPattern{
	Pattern: "/markdown/raw",
	Method:  "POST",
}

var GetEnterpriseMarketplaceListingAccountsByAccountId EndpointPattern = EndpointPattern{
	Pattern: "/marketplace_listing/accounts/{account_id}",
	Method:  "GET",
}

var GetEnterpriseMarketplaceListingPlans EndpointPattern = EndpointPattern{
	Pattern: "/marketplace_listing/plans",
	Method:  "GET",
}

var GetEnterpriseMarketplaceListingPlansAccountsByPlanId EndpointPattern = EndpointPattern{
	Pattern: "/marketplace_listing/plans/{plan_id}/accounts",
	Method:  "GET",
}

var GetEnterpriseMarketplaceListingStubbedAccountsByAccountId EndpointPattern = EndpointPattern{
	Pattern: "/marketplace_listing/stubbed/accounts/{account_id}",
	Method:  "GET",
}

var GetEnterpriseMarketplaceListingStubbedPlans EndpointPattern = EndpointPattern{
	Pattern: "/marketplace_listing/stubbed/plans",
	Method:  "GET",
}

var GetEnterpriseMarketplaceListingStubbedPlansAccountsByPlanId EndpointPattern = EndpointPattern{
	Pattern: "/marketplace_listing/stubbed/plans/{plan_id}/accounts",
	Method:  "GET",
}

var GetEnterpriseMeta EndpointPattern = EndpointPattern{
	Pattern: "/meta",
	Method:  "GET",
}

var GetEnterpriseNetworksEventsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/networks/{owner}/{repo}/events",
	Method:  "GET",
}

var GetEnterpriseNotifications EndpointPattern = EndpointPattern{
	Pattern: "/notifications",
	Method:  "GET",
}

var PutEnterpriseNotifications EndpointPattern = EndpointPattern{
	Pattern: "/notifications",
	Method:  "PUT",
}

var GetEnterpriseNotificationsThreadsByThreadId EndpointPattern = EndpointPattern{
	Pattern: "/notifications/threads/{thread_id}",
	Method:  "GET",
}

var PatchEnterpriseNotificationsThreadsByThreadId EndpointPattern = EndpointPattern{
	Pattern: "/notifications/threads/{thread_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseNotificationsThreadsByThreadId EndpointPattern = EndpointPattern{
	Pattern: "/notifications/threads/{thread_id}",
	Method:  "DELETE",
}

var GetEnterpriseNotificationsThreadsSubscriptionByThreadId EndpointPattern = EndpointPattern{
	Pattern: "/notifications/threads/{thread_id}/subscription",
	Method:  "GET",
}

var PutEnterpriseNotificationsThreadsSubscriptionByThreadId EndpointPattern = EndpointPattern{
	Pattern: "/notifications/threads/{thread_id}/subscription",
	Method:  "PUT",
}

var DeleteEnterpriseNotificationsThreadsSubscriptionByThreadId EndpointPattern = EndpointPattern{
	Pattern: "/notifications/threads/{thread_id}/subscription",
	Method:  "DELETE",
}

var GetEnterpriseOctocat EndpointPattern = EndpointPattern{
	Pattern: "/octocat",
	Method:  "GET",
}

var GetEnterpriseOrganizations EndpointPattern = EndpointPattern{
	Pattern: "/organizations",
	Method:  "GET",
}

var GetEnterpriseOrganizationsCustomRolesByOrganizationId EndpointPattern = EndpointPattern{
	Pattern: "/organizations/{organization_id}/custom_roles",
	Method:  "GET",
}

var GetEnterpriseOrgsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}",
	Method:  "GET",
}

var PatchEnterpriseOrgsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsCacheUsageByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/cache/usage",
	Method:  "GET",
}

var GetEnterpriseOrgsActionsCacheUsageByRepositoryByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/cache/usage-by-repository",
	Method:  "GET",
}

var GetEnterpriseOrgsActionsOidcCustomizationSubByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/oidc/customization/sub",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsOidcCustomizationSubByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/oidc/customization/sub",
	Method:  "PUT",
}

var GetEnterpriseOrgsActionsPermissionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsPermissionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions",
	Method:  "PUT",
}

var GetEnterpriseOrgsActionsPermissionsRepositoriesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/repositories",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsPermissionsRepositoriesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/repositories",
	Method:  "PUT",
}

var PutEnterpriseOrgsActionsPermissionsRepositoriesByOrgByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsPermissionsRepositoriesByOrgByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsPermissionsSelectedActionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/selected-actions",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsPermissionsSelectedActionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/selected-actions",
	Method:  "PUT",
}

var GetEnterpriseOrgsActionsPermissionsWorkflowByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/workflow",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsPermissionsWorkflowByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/permissions/workflow",
	Method:  "PUT",
}

var GetEnterpriseOrgsActionsRunnerGroupsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups",
	Method:  "GET",
}

var PostEnterpriseOrgsActionsRunnerGroupsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups",
	Method:  "POST",
}

var GetEnterpriseOrgsActionsRunnerGroupsByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}",
	Method:  "GET",
}

var PatchEnterpriseOrgsActionsRunnerGroupsByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsActionsRunnerGroupsByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsRunnerGroupsRepositoriesByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/repositories",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsRunnerGroupsRepositoriesByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/repositories",
	Method:  "PUT",
}

var PutEnterpriseOrgsActionsRunnerGroupsRepositoriesByOrgByRunnerGroupIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsRunnerGroupsRepositoriesByOrgByRunnerGroupIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsRunnerGroupsRunnersByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/runners",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsRunnerGroupsRunnersByOrgByRunnerGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/runners",
	Method:  "PUT",
}

var PutEnterpriseOrgsActionsRunnerGroupsRunnersByOrgByRunnerGroupIdByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsRunnerGroupsRunnersByOrgByRunnerGroupIdByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsRunnersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners",
	Method:  "GET",
}

var GetEnterpriseOrgsActionsRunnersDownloadsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/downloads",
	Method:  "GET",
}

var PostEnterpriseOrgsActionsRunnersGenerateJitconfigByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/generate-jitconfig",
	Method:  "POST",
}

var PostEnterpriseOrgsActionsRunnersRegistrationTokenByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/registration-token",
	Method:  "POST",
}

var PostEnterpriseOrgsActionsRunnersRemoveTokenByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/remove-token",
	Method:  "POST",
}

var GetEnterpriseOrgsActionsRunnersByOrgByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}",
	Method:  "GET",
}

var DeleteEnterpriseOrgsActionsRunnersByOrgByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsRunnersLabelsByOrgByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}/labels",
	Method:  "GET",
}

var PostEnterpriseOrgsActionsRunnersLabelsByOrgByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}/labels",
	Method:  "POST",
}

var PutEnterpriseOrgsActionsRunnersLabelsByOrgByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}/labels",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsRunnersLabelsByOrgByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}/labels",
	Method:  "DELETE",
}

var DeleteEnterpriseOrgsActionsRunnersLabelsByOrgByRunnerIdByName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/runners/{runner_id}/labels/{name:.+}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsSecretsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets",
	Method:  "GET",
}

var GetEnterpriseOrgsActionsSecretsPublicKeyByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseOrgsActionsSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsSecretsRepositoriesByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}/repositories",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsSecretsRepositoriesByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}/repositories",
	Method:  "PUT",
}

var PutEnterpriseOrgsActionsSecretsRepositoriesByOrgBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsSecretsRepositoriesByOrgBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsVariablesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables",
	Method:  "GET",
}

var PostEnterpriseOrgsActionsVariablesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables",
	Method:  "POST",
}

var GetEnterpriseOrgsActionsVariablesByOrgByName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}",
	Method:  "GET",
}

var PatchEnterpriseOrgsActionsVariablesByOrgByName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsActionsVariablesByOrgByName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsActionsVariablesRepositoriesByOrgByName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}/repositories",
	Method:  "GET",
}

var PutEnterpriseOrgsActionsVariablesRepositoriesByOrgByName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}/repositories",
	Method:  "PUT",
}

var PutEnterpriseOrgsActionsVariablesRepositoriesByOrgByNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsActionsVariablesRepositoriesByOrgByNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/actions/variables/{name}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsAnnouncementByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/announcement",
	Method:  "GET",
}

var PatchEnterpriseOrgsAnnouncementByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/announcement",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsAnnouncementByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/announcement",
	Method:  "DELETE",
}

var GetEnterpriseOrgsAttestationsByOrgBySubjectDigest EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/attestations/{subject_digest}",
	Method:  "GET",
}

var GetEnterpriseOrgsAuditLogByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/audit-log",
	Method:  "GET",
}

var GetEnterpriseOrgsBlocksByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/blocks",
	Method:  "GET",
}

var GetEnterpriseOrgsBlocksByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/blocks/{username}",
	Method:  "GET",
}

var PutEnterpriseOrgsBlocksByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/blocks/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsBlocksByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/blocks/{username}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCodeScanningAlertsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-scanning/alerts",
	Method:  "GET",
}

var GetEnterpriseOrgsCodeSecurityConfigurationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations",
	Method:  "GET",
}

var PostEnterpriseOrgsCodeSecurityConfigurationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations",
	Method:  "POST",
}

var GetEnterpriseOrgsCodeSecurityConfigurationsDefaultsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/defaults",
	Method:  "GET",
}

var DeleteEnterpriseOrgsCodeSecurityConfigurationsDetachByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/detach",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCodeSecurityConfigurationsByOrgByConfigurationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/{configuration_id}",
	Method:  "GET",
}

var PatchEnterpriseOrgsCodeSecurityConfigurationsByOrgByConfigurationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/{configuration_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsCodeSecurityConfigurationsByOrgByConfigurationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/{configuration_id}",
	Method:  "DELETE",
}

var PostEnterpriseOrgsCodeSecurityConfigurationsAttachByOrgByConfigurationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/{configuration_id}/attach",
	Method:  "POST",
}

var PutEnterpriseOrgsCodeSecurityConfigurationsDefaultsByOrgByConfigurationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/{configuration_id}/defaults",
	Method:  "PUT",
}

var GetEnterpriseOrgsCodeSecurityConfigurationsRepositoriesByOrgByConfigurationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/code-security/configurations/{configuration_id}/repositories",
	Method:  "GET",
}

var GetEnterpriseOrgsCodespacesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces",
	Method:  "GET",
}

var PutEnterpriseOrgsCodespacesAccessByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/access",
	Method:  "PUT",
}

var PostEnterpriseOrgsCodespacesAccessSelectedUsersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/access/selected_users",
	Method:  "POST",
}

var DeleteEnterpriseOrgsCodespacesAccessSelectedUsersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/access/selected_users",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCodespacesSecretsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets",
	Method:  "GET",
}

var GetEnterpriseOrgsCodespacesSecretsPublicKeyByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseOrgsCodespacesSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseOrgsCodespacesSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsCodespacesSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCodespacesSecretsRepositoriesByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}/repositories",
	Method:  "GET",
}

var PutEnterpriseOrgsCodespacesSecretsRepositoriesByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}/repositories",
	Method:  "PUT",
}

var PutEnterpriseOrgsCodespacesSecretsRepositoriesByOrgBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsCodespacesSecretsRepositoriesByOrgBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/codespaces/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCopilotBillingByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/billing",
	Method:  "GET",
}

var GetEnterpriseOrgsCopilotBillingSeatsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/billing/seats",
	Method:  "GET",
}

var PostEnterpriseOrgsCopilotBillingSelectedTeamsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/billing/selected_teams",
	Method:  "POST",
}

var DeleteEnterpriseOrgsCopilotBillingSelectedTeamsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/billing/selected_teams",
	Method:  "DELETE",
}

var PostEnterpriseOrgsCopilotBillingSelectedUsersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/billing/selected_users",
	Method:  "POST",
}

var DeleteEnterpriseOrgsCopilotBillingSelectedUsersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/billing/selected_users",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCopilotUsageByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/copilot/usage",
	Method:  "GET",
}

var GetEnterpriseOrgsCredentialAuthorizationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/credential-authorizations",
	Method:  "GET",
}

var DeleteEnterpriseOrgsCredentialAuthorizationsByOrgByCredentialId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/credential-authorizations/{credential_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsCustomRepositoryRolesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom-repository-roles",
	Method:  "GET",
}

var PostEnterpriseOrgsCustomRepositoryRolesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom-repository-roles",
	Method:  "POST",
}

var GetEnterpriseOrgsCustomRepositoryRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom-repository-roles/{role_id}",
	Method:  "GET",
}

var PatchEnterpriseOrgsCustomRepositoryRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom-repository-roles/{role_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsCustomRepositoryRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom-repository-roles/{role_id}",
	Method:  "DELETE",
}

var PostEnterpriseOrgsCustomRolesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom_roles",
	Method:  "POST",
}

var GetEnterpriseOrgsCustomRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom_roles/{role_id}",
	Method:  "GET",
}

var PatchEnterpriseOrgsCustomRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom_roles/{role_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsCustomRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/custom_roles/{role_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsDependabotAlertsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/alerts",
	Method:  "GET",
}

var GetEnterpriseOrgsDependabotSecretsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets",
	Method:  "GET",
}

var GetEnterpriseOrgsDependabotSecretsPublicKeyByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseOrgsDependabotSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseOrgsDependabotSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsDependabotSecretsByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsDependabotSecretsRepositoriesByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}/repositories",
	Method:  "GET",
}

var PutEnterpriseOrgsDependabotSecretsRepositoriesByOrgBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}/repositories",
	Method:  "PUT",
}

var PutEnterpriseOrgsDependabotSecretsRepositoriesByOrgBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsDependabotSecretsRepositoriesByOrgBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/dependabot/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsDockerConflictsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/docker/conflicts",
	Method:  "GET",
}

var GetEnterpriseOrgsEventsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/events",
	Method:  "GET",
}

var GetEnterpriseOrgsExternalGroupByOrgByGroupId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/external-group/{group_id}",
	Method:  "GET",
}

var GetEnterpriseOrgsExternalGroupsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/external-groups",
	Method:  "GET",
}

var GetEnterpriseOrgsFailedInvitationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/failed_invitations",
	Method:  "GET",
}

var GetEnterpriseOrgsFineGrainedPermissionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/fine_grained_permissions",
	Method:  "GET",
}

var GetEnterpriseOrgsHooksByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks",
	Method:  "GET",
}

var PostEnterpriseOrgsHooksByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks",
	Method:  "POST",
}

var GetEnterpriseOrgsHooksByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}",
	Method:  "GET",
}

var PatchEnterpriseOrgsHooksByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsHooksByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsHooksConfigByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}/config",
	Method:  "GET",
}

var PatchEnterpriseOrgsHooksConfigByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}/config",
	Method:  "PATCH",
}

var GetEnterpriseOrgsHooksDeliveriesByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}/deliveries",
	Method:  "GET",
}

var GetEnterpriseOrgsHooksDeliveriesByOrgByHookIdByDeliveryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}",
	Method:  "GET",
}

var PostEnterpriseOrgsHooksDeliveriesAttemptsByOrgByHookIdByDeliveryId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}/attempts",
	Method:  "POST",
}

var PostEnterpriseOrgsHooksPingsByOrgByHookId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/hooks/{hook_id}/pings",
	Method:  "POST",
}

var GetEnterpriseOrgsInstallationByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/installation",
	Method:  "GET",
}

var GetEnterpriseOrgsInstallationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/installations",
	Method:  "GET",
}

var GetEnterpriseOrgsInteractionLimitsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/interaction-limits",
	Method:  "GET",
}

var PutEnterpriseOrgsInteractionLimitsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/interaction-limits",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsInteractionLimitsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/interaction-limits",
	Method:  "DELETE",
}

var GetEnterpriseOrgsInvitationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/invitations",
	Method:  "GET",
}

var PostEnterpriseOrgsInvitationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/invitations",
	Method:  "POST",
}

var DeleteEnterpriseOrgsInvitationsByOrgByInvitationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/invitations/{invitation_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsInvitationsTeamsByOrgByInvitationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/invitations/{invitation_id}/teams",
	Method:  "GET",
}

var GetEnterpriseOrgsIssuesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/issues",
	Method:  "GET",
}

var GetEnterpriseOrgsMembersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members",
	Method:  "GET",
}

var GetEnterpriseOrgsMembersByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members/{username}",
	Method:  "GET",
}

var DeleteEnterpriseOrgsMembersByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members/{username}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsMembersCodespacesByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members/{username}/codespaces",
	Method:  "GET",
}

var DeleteEnterpriseOrgsMembersCodespacesByOrgByUsernameByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members/{username}/codespaces/{codespace_name}",
	Method:  "DELETE",
}

var PostEnterpriseOrgsMembersCodespacesStopByOrgByUsernameByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members/{username}/codespaces/{codespace_name}/stop",
	Method:  "POST",
}

var GetEnterpriseOrgsMembersCopilotByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/members/{username}/copilot",
	Method:  "GET",
}

var GetEnterpriseOrgsMembershipsByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/memberships/{username}",
	Method:  "GET",
}

var PutEnterpriseOrgsMembershipsByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/memberships/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsMembershipsByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/memberships/{username}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsMigrationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations",
	Method:  "GET",
}

var PostEnterpriseOrgsMigrationsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations",
	Method:  "POST",
}

var GetEnterpriseOrgsMigrationsByOrgByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations/{migration_id}",
	Method:  "GET",
}

var GetEnterpriseOrgsMigrationsArchiveByOrgByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations/{migration_id}/archive",
	Method:  "GET",
}

var DeleteEnterpriseOrgsMigrationsArchiveByOrgByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations/{migration_id}/archive",
	Method:  "DELETE",
}

var DeleteEnterpriseOrgsMigrationsReposLockByOrgByMigrationIdByRepoName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock",
	Method:  "DELETE",
}

var GetEnterpriseOrgsMigrationsRepositoriesByOrgByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/migrations/{migration_id}/repositories",
	Method:  "GET",
}

var GetEnterpriseOrgsOrganizationFineGrainedPermissionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-fine-grained-permissions",
	Method:  "GET",
}

var GetEnterpriseOrgsOrganizationRolesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles",
	Method:  "GET",
}

var PostEnterpriseOrgsOrganizationRolesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles",
	Method:  "POST",
}

var DeleteEnterpriseOrgsOrganizationRolesTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/teams/{team_slug}",
	Method:  "DELETE",
}

var PutEnterpriseOrgsOrganizationRolesTeamsByOrgByTeamSlugByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/teams/{team_slug}/{role_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsOrganizationRolesTeamsByOrgByTeamSlugByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/teams/{team_slug}/{role_id}",
	Method:  "DELETE",
}

var DeleteEnterpriseOrgsOrganizationRolesUsersByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/users/{username}",
	Method:  "DELETE",
}

var PutEnterpriseOrgsOrganizationRolesUsersByOrgByUsernameByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/users/{username}/{role_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsOrganizationRolesUsersByOrgByUsernameByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/users/{username}/{role_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsOrganizationRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/{role_id}",
	Method:  "GET",
}

var PatchEnterpriseOrgsOrganizationRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/{role_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsOrganizationRolesByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/{role_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsOrganizationRolesTeamsByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/{role_id}/teams",
	Method:  "GET",
}

var GetEnterpriseOrgsOrganizationRolesUsersByOrgByRoleId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/organization-roles/{role_id}/users",
	Method:  "GET",
}

var GetEnterpriseOrgsOutsideCollaboratorsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/outside_collaborators",
	Method:  "GET",
}

var PutEnterpriseOrgsOutsideCollaboratorsByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/outside_collaborators/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsOutsideCollaboratorsByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/outside_collaborators/{username}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsPackagesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages",
	Method:  "GET",
}

var GetEnterpriseOrgsPackagesByOrgByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}",
	Method:  "GET",
}

var DeleteEnterpriseOrgsPackagesByOrgByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}",
	Method:  "DELETE",
}

var PostEnterpriseOrgsPackagesRestoreByOrgByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}/restore",
	Method:  "POST",
}

var GetEnterpriseOrgsPackagesVersionsByOrgByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}/versions",
	Method:  "GET",
}

var GetEnterpriseOrgsPackagesVersionsByOrgByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}",
	Method:  "GET",
}

var DeleteEnterpriseOrgsPackagesVersionsByOrgByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}",
	Method:  "DELETE",
}

var PostEnterpriseOrgsPackagesVersionsRestoreByOrgByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore",
	Method:  "POST",
}

var GetEnterpriseOrgsPersonalAccessTokenRequestsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-token-requests",
	Method:  "GET",
}

var PostEnterpriseOrgsPersonalAccessTokenRequestsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-token-requests",
	Method:  "POST",
}

var PostEnterpriseOrgsPersonalAccessTokenRequestsByOrgByPatRequestId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-token-requests/{pat_request_id}",
	Method:  "POST",
}

var GetEnterpriseOrgsPersonalAccessTokenRequestsRepositoriesByOrgByPatRequestId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-token-requests/{pat_request_id}/repositories",
	Method:  "GET",
}

var GetEnterpriseOrgsPersonalAccessTokensByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-tokens",
	Method:  "GET",
}

var PostEnterpriseOrgsPersonalAccessTokensByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-tokens",
	Method:  "POST",
}

var PostEnterpriseOrgsPersonalAccessTokensByOrgByPatId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-tokens/{pat_id}",
	Method:  "POST",
}

var GetEnterpriseOrgsPersonalAccessTokensRepositoriesByOrgByPatId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/personal-access-tokens/{pat_id}/repositories",
	Method:  "GET",
}

var GetEnterpriseOrgsProjectsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/projects",
	Method:  "GET",
}

var PostEnterpriseOrgsProjectsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/projects",
	Method:  "POST",
}

var GetEnterpriseOrgsPropertiesSchemaByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/schema",
	Method:  "GET",
}

var PatchEnterpriseOrgsPropertiesSchemaByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/schema",
	Method:  "PATCH",
}

var GetEnterpriseOrgsPropertiesSchemaByOrgByCustomPropertyName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/schema/{custom_property_name}",
	Method:  "GET",
}

var PutEnterpriseOrgsPropertiesSchemaByOrgByCustomPropertyName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/schema/{custom_property_name}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsPropertiesSchemaByOrgByCustomPropertyName EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/schema/{custom_property_name}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsPropertiesValuesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/values",
	Method:  "GET",
}

var PatchEnterpriseOrgsPropertiesValuesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/properties/values",
	Method:  "PATCH",
}

var GetEnterpriseOrgsPublicMembersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/public_members",
	Method:  "GET",
}

var GetEnterpriseOrgsPublicMembersByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/public_members/{username}",
	Method:  "GET",
}

var PutEnterpriseOrgsPublicMembersByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/public_members/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsPublicMembersByOrgByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/public_members/{username}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsReposByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/repos",
	Method:  "GET",
}

var PostEnterpriseOrgsReposByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/repos",
	Method:  "POST",
}

var GetEnterpriseOrgsRepositoryFineGrainedPermissionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/repository-fine-grained-permissions",
	Method:  "GET",
}

var GetEnterpriseOrgsRulesetsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets",
	Method:  "GET",
}

var PostEnterpriseOrgsRulesetsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets",
	Method:  "POST",
}

var GetEnterpriseOrgsRulesetsRuleSuitesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets/rule-suites",
	Method:  "GET",
}

var GetEnterpriseOrgsRulesetsRuleSuitesByOrgByRuleSuiteId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets/rule-suites/{rule_suite_id}",
	Method:  "GET",
}

var GetEnterpriseOrgsRulesetsByOrgByRulesetId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets/{ruleset_id}",
	Method:  "GET",
}

var PutEnterpriseOrgsRulesetsByOrgByRulesetId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets/{ruleset_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsRulesetsByOrgByRulesetId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/rulesets/{ruleset_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsSecretScanningAlertsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/secret-scanning/alerts",
	Method:  "GET",
}

var GetEnterpriseOrgsSecurityAdvisoriesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/security-advisories",
	Method:  "GET",
}

var GetEnterpriseOrgsSecurityManagersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/security-managers",
	Method:  "GET",
}

var PutEnterpriseOrgsSecurityManagersTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/security-managers/teams/{team_slug}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsSecurityManagersTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/security-managers/teams/{team_slug}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsSettingsBillingActionsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/settings/billing/actions",
	Method:  "GET",
}

var GetEnterpriseOrgsSettingsBillingAdvancedSecurityByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/settings/billing/advanced-security",
	Method:  "GET",
}

var GetEnterpriseOrgsSettingsBillingPackagesByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/settings/billing/packages",
	Method:  "GET",
}

var GetEnterpriseOrgsSettingsBillingSharedStorageByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/settings/billing/shared-storage",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamSyncGroupsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/team-sync/groups",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamCopilotUsageByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/team/{team_slug}/copilot/usage",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams",
	Method:  "GET",
}

var PostEnterpriseOrgsTeamsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams",
	Method:  "POST",
}

var GetEnterpriseOrgsTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}",
	Method:  "GET",
}

var PatchEnterpriseOrgsTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsDiscussionsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions",
	Method:  "GET",
}

var PostEnterpriseOrgsTeamsDiscussionsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions",
	Method:  "POST",
}

var GetEnterpriseOrgsTeamsDiscussionsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}",
	Method:  "GET",
}

var PatchEnterpriseOrgsTeamsDiscussionsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsTeamsDiscussionsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsDiscussionsCommentsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments",
	Method:  "GET",
}

var PostEnterpriseOrgsTeamsDiscussionsCommentsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments",
	Method:  "POST",
}

var GetEnterpriseOrgsTeamsDiscussionsCommentsByOrgByTeamSlugByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}",
	Method:  "GET",
}

var PatchEnterpriseOrgsTeamsDiscussionsCommentsByOrgByTeamSlugByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsTeamsDiscussionsCommentsByOrgByTeamSlugByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsDiscussionsCommentsReactionsByOrgByTeamSlugByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions",
	Method:  "GET",
}

var PostEnterpriseOrgsTeamsDiscussionsCommentsReactionsByOrgByTeamSlugByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseOrgsTeamsDiscussionsCommentsReactionsByOrgByTeamSlugByDiscussionNumberByCommentNumberByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsDiscussionsReactionsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions",
	Method:  "GET",
}

var PostEnterpriseOrgsTeamsDiscussionsReactionsByOrgByTeamSlugByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseOrgsTeamsDiscussionsReactionsByOrgByTeamSlugByDiscussionNumberByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsExternalGroupsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/external-groups",
	Method:  "GET",
}

var PatchEnterpriseOrgsTeamsExternalGroupsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/external-groups",
	Method:  "PATCH",
}

var DeleteEnterpriseOrgsTeamsExternalGroupsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/external-groups",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsInvitationsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/invitations",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamsMembersByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/members",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamsMembershipsByOrgByTeamSlugByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/memberships/{username}",
	Method:  "GET",
}

var PutEnterpriseOrgsTeamsMembershipsByOrgByTeamSlugByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/memberships/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsTeamsMembershipsByOrgByTeamSlugByUsername EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/memberships/{username}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsProjectsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/projects",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamsProjectsByOrgByTeamSlugByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/projects/{project_id}",
	Method:  "GET",
}

var PutEnterpriseOrgsTeamsProjectsByOrgByTeamSlugByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/projects/{project_id}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsTeamsProjectsByOrgByTeamSlugByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/projects/{project_id}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsReposByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/repos",
	Method:  "GET",
}

var GetEnterpriseOrgsTeamsReposByOrgByTeamSlugByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}",
	Method:  "GET",
}

var PutEnterpriseOrgsTeamsReposByOrgByTeamSlugByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}",
	Method:  "PUT",
}

var DeleteEnterpriseOrgsTeamsReposByOrgByTeamSlugByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}",
	Method:  "DELETE",
}

var GetEnterpriseOrgsTeamsTeamSyncGroupMappingsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/team-sync/group-mappings",
	Method:  "GET",
}

var PatchEnterpriseOrgsTeamsTeamSyncGroupMappingsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/team-sync/group-mappings",
	Method:  "PATCH",
}

var GetEnterpriseOrgsTeamsTeamsByOrgByTeamSlug EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/teams/{team_slug}/teams",
	Method:  "GET",
}

var PostEnterpriseOrgsByOrgBySecurityProductByEnablement EndpointPattern = EndpointPattern{
	Pattern: "/orgs/{org}/{security_product}/{enablement}",
	Method:  "POST",
}

var GetEnterpriseProjectsColumnsCardsByCardId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/cards/{card_id}",
	Method:  "GET",
}

var PatchEnterpriseProjectsColumnsCardsByCardId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/cards/{card_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseProjectsColumnsCardsByCardId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/cards/{card_id}",
	Method:  "DELETE",
}

var PostEnterpriseProjectsColumnsCardsMovesByCardId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/cards/{card_id}/moves",
	Method:  "POST",
}

var GetEnterpriseProjectsColumnsByColumnId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/{column_id}",
	Method:  "GET",
}

var PatchEnterpriseProjectsColumnsByColumnId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/{column_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseProjectsColumnsByColumnId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/{column_id}",
	Method:  "DELETE",
}

var GetEnterpriseProjectsColumnsCardsByColumnId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/{column_id}/cards",
	Method:  "GET",
}

var PostEnterpriseProjectsColumnsCardsByColumnId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/{column_id}/cards",
	Method:  "POST",
}

var PostEnterpriseProjectsColumnsMovesByColumnId EndpointPattern = EndpointPattern{
	Pattern: "/projects/columns/{column_id}/moves",
	Method:  "POST",
}

var GetEnterpriseProjectsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}",
	Method:  "GET",
}

var PatchEnterpriseProjectsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseProjectsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}",
	Method:  "DELETE",
}

var GetEnterpriseProjectsCollaboratorsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}/collaborators",
	Method:  "GET",
}

var PutEnterpriseProjectsCollaboratorsByProjectIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}/collaborators/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseProjectsCollaboratorsByProjectIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}/collaborators/{username}",
	Method:  "DELETE",
}

var GetEnterpriseProjectsCollaboratorsPermissionByProjectIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}/collaborators/{username}/permission",
	Method:  "GET",
}

var GetEnterpriseProjectsColumnsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}/columns",
	Method:  "GET",
}

var PostEnterpriseProjectsColumnsByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/projects/{project_id}/columns",
	Method:  "POST",
}

var GetEnterpriseRateLimit EndpointPattern = EndpointPattern{
	Pattern: "/rate_limit",
	Method:  "GET",
}

var GetEnterpriseReposByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}",
	Method:  "GET",
}

var PatchEnterpriseReposByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsArtifactsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/artifacts",
	Method:  "GET",
}

var GetEnterpriseReposActionsArtifactsByOwnerByRepoByArtifactId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/artifacts/{artifact_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposActionsArtifactsByOwnerByRepoByArtifactId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/artifacts/{artifact_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsArtifactsByOwnerByRepoByArtifactIdByArchiveFormat EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/artifacts/{artifact_id}/{archive_format}",
	Method:  "GET",
}

var GetEnterpriseReposActionsCacheUsageByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/cache/usage",
	Method:  "GET",
}

var GetEnterpriseReposActionsCachesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/caches",
	Method:  "GET",
}

var DeleteEnterpriseReposActionsCachesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/caches",
	Method:  "DELETE",
}

var DeleteEnterpriseReposActionsCachesByOwnerByRepoByCacheId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/caches/{cache_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsJobsByOwnerByRepoByJobId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/jobs/{job_id}",
	Method:  "GET",
}

var GetEnterpriseReposActionsJobsLogsByOwnerByRepoByJobId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/jobs/{job_id}/logs",
	Method:  "GET",
}

var PostEnterpriseReposActionsJobsRerunByOwnerByRepoByJobId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/jobs/{job_id}/rerun",
	Method:  "POST",
}

var GetEnterpriseReposActionsOidcCustomizationSubByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/oidc/customization/sub",
	Method:  "GET",
}

var PutEnterpriseReposActionsOidcCustomizationSubByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/oidc/customization/sub",
	Method:  "PUT",
}

var GetEnterpriseReposActionsOrganizationSecretsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/organization-secrets",
	Method:  "GET",
}

var GetEnterpriseReposActionsOrganizationVariablesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/organization-variables",
	Method:  "GET",
}

var GetEnterpriseReposActionsPermissionsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions",
	Method:  "GET",
}

var PutEnterpriseReposActionsPermissionsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions",
	Method:  "PUT",
}

var GetEnterpriseReposActionsPermissionsAccessByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions/access",
	Method:  "GET",
}

var PutEnterpriseReposActionsPermissionsAccessByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions/access",
	Method:  "PUT",
}

var GetEnterpriseReposActionsPermissionsSelectedActionsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions/selected-actions",
	Method:  "GET",
}

var PutEnterpriseReposActionsPermissionsSelectedActionsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions/selected-actions",
	Method:  "PUT",
}

var GetEnterpriseReposActionsPermissionsWorkflowByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions/workflow",
	Method:  "GET",
}

var PutEnterpriseReposActionsPermissionsWorkflowByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/permissions/workflow",
	Method:  "PUT",
}

var GetEnterpriseReposActionsRunnersByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners",
	Method:  "GET",
}

var GetEnterpriseReposActionsRunnersDownloadsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/downloads",
	Method:  "GET",
}

var PostEnterpriseReposActionsRunnersGenerateJitconfigByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/generate-jitconfig",
	Method:  "POST",
}

var PostEnterpriseReposActionsRunnersRegistrationTokenByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/registration-token",
	Method:  "POST",
}

var PostEnterpriseReposActionsRunnersRemoveTokenByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/remove-token",
	Method:  "POST",
}

var GetEnterpriseReposActionsRunnersByOwnerByRepoByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposActionsRunnersByOwnerByRepoByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsRunnersLabelsByOwnerByRepoByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}/labels",
	Method:  "GET",
}

var PostEnterpriseReposActionsRunnersLabelsByOwnerByRepoByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}/labels",
	Method:  "POST",
}

var PutEnterpriseReposActionsRunnersLabelsByOwnerByRepoByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}/labels",
	Method:  "PUT",
}

var DeleteEnterpriseReposActionsRunnersLabelsByOwnerByRepoByRunnerId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}/labels",
	Method:  "DELETE",
}

var DeleteEnterpriseReposActionsRunnersLabelsByOwnerByRepoByRunnerIdByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runners/{runner_id}/labels/{name}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsRunsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs",
	Method:  "GET",
}

var GetEnterpriseReposActionsRunsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposActionsRunsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsRunsApprovalsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/approvals",
	Method:  "GET",
}

var PostEnterpriseReposActionsRunsApproveByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/approve",
	Method:  "POST",
}

var GetEnterpriseReposActionsRunsArtifactsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/artifacts",
	Method:  "GET",
}

var GetEnterpriseReposActionsRunsAttemptsByOwnerByRepoByRunIdByAttemptNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}",
	Method:  "GET",
}

var GetEnterpriseReposActionsRunsAttemptsJobsByOwnerByRepoByRunIdByAttemptNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/jobs",
	Method:  "GET",
}

var GetEnterpriseReposActionsRunsAttemptsLogsByOwnerByRepoByRunIdByAttemptNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/attempts/{attempt_number}/logs",
	Method:  "GET",
}

var PostEnterpriseReposActionsRunsCancelByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/cancel",
	Method:  "POST",
}

var PostEnterpriseReposActionsRunsDeploymentProtectionRuleByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/deployment_protection_rule",
	Method:  "POST",
}

var PostEnterpriseReposActionsRunsForceCancelByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/force-cancel",
	Method:  "POST",
}

var GetEnterpriseReposActionsRunsJobsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/jobs",
	Method:  "GET",
}

var GetEnterpriseReposActionsRunsLogsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/logs",
	Method:  "GET",
}

var DeleteEnterpriseReposActionsRunsLogsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/logs",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsRunsPendingDeploymentsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments",
	Method:  "GET",
}

var PostEnterpriseReposActionsRunsPendingDeploymentsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments",
	Method:  "POST",
}

var PostEnterpriseReposActionsRunsRerunByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/rerun",
	Method:  "POST",
}

var PostEnterpriseReposActionsRunsRerunFailedJobsByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/rerun-failed-jobs",
	Method:  "POST",
}

var GetEnterpriseReposActionsRunsTimingByOwnerByRepoByRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/timing",
	Method:  "GET",
}

var GetEnterpriseReposActionsSecretsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/secrets",
	Method:  "GET",
}

var GetEnterpriseReposActionsSecretsPublicKeyByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseReposActionsSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseReposActionsSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseReposActionsSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsVariablesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/variables",
	Method:  "GET",
}

var PostEnterpriseReposActionsVariablesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/variables",
	Method:  "POST",
}

var GetEnterpriseReposActionsVariablesByOwnerByRepoByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/variables/{name}",
	Method:  "GET",
}

var PatchEnterpriseReposActionsVariablesByOwnerByRepoByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/variables/{name}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposActionsVariablesByOwnerByRepoByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/variables/{name}",
	Method:  "DELETE",
}

var GetEnterpriseReposActionsWorkflowsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows",
	Method:  "GET",
}

var GetEnterpriseReposActionsWorkflowsByOwnerByRepoByWorkflowId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows/{workflow_id}",
	Method:  "GET",
}

var PutEnterpriseReposActionsWorkflowsDisableByOwnerByRepoByWorkflowId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable",
	Method:  "PUT",
}

var PostEnterpriseReposActionsWorkflowsDispatchesByOwnerByRepoByWorkflowId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches",
	Method:  "POST",
}

var PutEnterpriseReposActionsWorkflowsEnableByOwnerByRepoByWorkflowId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable",
	Method:  "PUT",
}

var GetEnterpriseReposActionsWorkflowsRunsByOwnerByRepoByWorkflowId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs",
	Method:  "GET",
}

var GetEnterpriseReposActionsWorkflowsTimingByOwnerByRepoByWorkflowId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing",
	Method:  "GET",
}

var GetEnterpriseReposActivityByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/activity",
	Method:  "GET",
}

var GetEnterpriseReposAssigneesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/assignees",
	Method:  "GET",
}

var GetEnterpriseReposAssigneesByOwnerByRepoByAssignee EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/assignees/{assignee}",
	Method:  "GET",
}

var PostEnterpriseReposAttestationsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/attestations",
	Method:  "POST",
}

var GetEnterpriseReposAttestationsByOwnerByRepoBySubjectDigest EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/attestations/{subject_digest}",
	Method:  "GET",
}

var GetEnterpriseReposAutolinksByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/autolinks",
	Method:  "GET",
}

var PostEnterpriseReposAutolinksByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/autolinks",
	Method:  "POST",
}

var GetEnterpriseReposAutolinksByOwnerByRepoByAutolinkId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/autolinks/{autolink_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposAutolinksByOwnerByRepoByAutolinkId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/autolinks/{autolink_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposAutomatedSecurityFixesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/automated-security-fixes",
	Method:  "GET",
}

var PutEnterpriseReposAutomatedSecurityFixesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/automated-security-fixes",
	Method:  "PUT",
}

var DeleteEnterpriseReposAutomatedSecurityFixesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/automated-security-fixes",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches",
	Method:  "GET",
}

var GetEnterpriseReposBranchesByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}",
	Method:  "GET",
}

var GetEnterpriseReposBranchesProtectionByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection",
	Method:  "GET",
}

var PutEnterpriseReposBranchesProtectionByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection",
	Method:  "PUT",
}

var DeleteEnterpriseReposBranchesProtectionByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionEnforceAdminsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins",
	Method:  "GET",
}

var PostEnterpriseReposBranchesProtectionEnforceAdminsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins",
	Method:  "POST",
}

var DeleteEnterpriseReposBranchesProtectionEnforceAdminsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRequiredPullRequestReviewsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews",
	Method:  "GET",
}

var PatchEnterpriseReposBranchesProtectionRequiredPullRequestReviewsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews",
	Method:  "PATCH",
}

var DeleteEnterpriseReposBranchesProtectionRequiredPullRequestReviewsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRequiredSignaturesByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_signatures",
	Method:  "GET",
}

var PostEnterpriseReposBranchesProtectionRequiredSignaturesByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_signatures",
	Method:  "POST",
}

var DeleteEnterpriseReposBranchesProtectionRequiredSignaturesByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_signatures",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRequiredStatusChecksByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks",
	Method:  "GET",
}

var PatchEnterpriseReposBranchesProtectionRequiredStatusChecksByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks",
	Method:  "PATCH",
}

var DeleteEnterpriseReposBranchesProtectionRequiredStatusChecksByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRequiredStatusChecksContextsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts",
	Method:  "GET",
}

var PostEnterpriseReposBranchesProtectionRequiredStatusChecksContextsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts",
	Method:  "POST",
}

var PutEnterpriseReposBranchesProtectionRequiredStatusChecksContextsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts",
	Method:  "PUT",
}

var DeleteEnterpriseReposBranchesProtectionRequiredStatusChecksContextsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRestrictionsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions",
	Method:  "GET",
}

var DeleteEnterpriseReposBranchesProtectionRestrictionsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRestrictionsAppsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps",
	Method:  "GET",
}

var PostEnterpriseReposBranchesProtectionRestrictionsAppsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps",
	Method:  "POST",
}

var PutEnterpriseReposBranchesProtectionRestrictionsAppsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps",
	Method:  "PUT",
}

var DeleteEnterpriseReposBranchesProtectionRestrictionsAppsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRestrictionsTeamsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams",
	Method:  "GET",
}

var PostEnterpriseReposBranchesProtectionRestrictionsTeamsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams",
	Method:  "POST",
}

var PutEnterpriseReposBranchesProtectionRestrictionsTeamsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams",
	Method:  "PUT",
}

var DeleteEnterpriseReposBranchesProtectionRestrictionsTeamsByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams",
	Method:  "DELETE",
}

var GetEnterpriseReposBranchesProtectionRestrictionsUsersByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users",
	Method:  "GET",
}

var PostEnterpriseReposBranchesProtectionRestrictionsUsersByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users",
	Method:  "POST",
}

var PutEnterpriseReposBranchesProtectionRestrictionsUsersByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users",
	Method:  "PUT",
}

var DeleteEnterpriseReposBranchesProtectionRestrictionsUsersByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users",
	Method:  "DELETE",
}

var PostEnterpriseReposBranchesRenameByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/branches/{branch}/rename",
	Method:  "POST",
}

var PostEnterpriseReposCheckRunsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-runs",
	Method:  "POST",
}

var GetEnterpriseReposCheckRunsByOwnerByRepoByCheckRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-runs/{check_run_id}",
	Method:  "GET",
}

var PatchEnterpriseReposCheckRunsByOwnerByRepoByCheckRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-runs/{check_run_id}",
	Method:  "PATCH",
}

var GetEnterpriseReposCheckRunsAnnotationsByOwnerByRepoByCheckRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-runs/{check_run_id}/annotations",
	Method:  "GET",
}

var PostEnterpriseReposCheckRunsRerequestByOwnerByRepoByCheckRunId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-runs/{check_run_id}/rerequest",
	Method:  "POST",
}

var PostEnterpriseReposCheckSuitesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-suites",
	Method:  "POST",
}

var PatchEnterpriseReposCheckSuitesPreferencesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-suites/preferences",
	Method:  "PATCH",
}

var GetEnterpriseReposCheckSuitesByOwnerByRepoByCheckSuiteId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-suites/{check_suite_id}",
	Method:  "GET",
}

var GetEnterpriseReposCheckSuitesCheckRunsByOwnerByRepoByCheckSuiteId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-suites/{check_suite_id}/check-runs",
	Method:  "GET",
}

var PostEnterpriseReposCheckSuitesRerequestByOwnerByRepoByCheckSuiteId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/check-suites/{check_suite_id}/rerequest",
	Method:  "POST",
}

var GetEnterpriseReposCodeScanningAlertsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/alerts",
	Method:  "GET",
}

var GetEnterpriseReposCodeScanningAlertsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/alerts/{alert_number}",
	Method:  "GET",
}

var PatchEnterpriseReposCodeScanningAlertsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/alerts/{alert_number}",
	Method:  "PATCH",
}

var GetEnterpriseReposCodeScanningAlertsInstancesByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/alerts/{alert_number}/instances",
	Method:  "GET",
}

var GetEnterpriseReposCodeScanningAnalysesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/analyses",
	Method:  "GET",
}

var GetEnterpriseReposCodeScanningAnalysesByOwnerByRepoByAnalysisId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/analyses/{analysis_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposCodeScanningAnalysesByOwnerByRepoByAnalysisId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/analyses/{analysis_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposCodeScanningCodeqlDatabasesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/codeql/databases",
	Method:  "GET",
}

var GetEnterpriseReposCodeScanningCodeqlDatabasesByOwnerByRepoByLanguage EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/codeql/databases/{language}",
	Method:  "GET",
}

var PostEnterpriseReposCodeScanningCodeqlVariantAnalysesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/codeql/variant-analyses",
	Method:  "POST",
}

var GetEnterpriseReposCodeScanningCodeqlVariantAnalysesByOwnerByRepoByCodeqlVariantAnalysisId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/codeql/variant-analyses/{codeql_variant_analysis_id}",
	Method:  "GET",
}

var GetEnterpriseReposCodeScanningCodeqlVariantAnalysesReposByOwnerByRepoByCodeqlVariantAnalysisIdByRepoOwnerByRepoName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/codeql/variant-analyses/{codeql_variant_analysis_id}/repos/{repo_owner}/{repo_name}",
	Method:  "GET",
}

var GetEnterpriseReposCodeScanningDefaultSetupByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/default-setup",
	Method:  "GET",
}

var PatchEnterpriseReposCodeScanningDefaultSetupByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/default-setup",
	Method:  "PATCH",
}

var PostEnterpriseReposCodeScanningSarifsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/sarifs",
	Method:  "POST",
}

var GetEnterpriseReposCodeScanningSarifsByOwnerByRepoBySarifId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-scanning/sarifs/{sarif_id}",
	Method:  "GET",
}

var GetEnterpriseReposCodeSecurityConfigurationByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/code-security-configuration",
	Method:  "GET",
}

var GetEnterpriseReposCodeownersErrorsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codeowners/errors",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces",
	Method:  "GET",
}

var PostEnterpriseReposCodespacesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces",
	Method:  "POST",
}

var GetEnterpriseReposCodespacesDevcontainersByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/devcontainers",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesMachinesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/machines",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesNewByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/new",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesPermissionsCheckByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/permissions_check",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesSecretsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/secrets",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesSecretsPublicKeyByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseReposCodespacesSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseReposCodespacesSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseReposCodespacesSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/codespaces/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseReposCollaboratorsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/collaborators",
	Method:  "GET",
}

var GetEnterpriseReposCollaboratorsByOwnerByRepoByUsername EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/collaborators/{username}",
	Method:  "GET",
}

var PutEnterpriseReposCollaboratorsByOwnerByRepoByUsername EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/collaborators/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseReposCollaboratorsByOwnerByRepoByUsername EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/collaborators/{username}",
	Method:  "DELETE",
}

var GetEnterpriseReposCollaboratorsPermissionByOwnerByRepoByUsername EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/collaborators/{username}/permission",
	Method:  "GET",
}

var GetEnterpriseReposCommentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments",
	Method:  "GET",
}

var GetEnterpriseReposCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments/{comment_id}",
	Method:  "GET",
}

var PatchEnterpriseReposCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments/{comment_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments/{comment_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposCommentsReactionsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments/{comment_id}/reactions",
	Method:  "GET",
}

var PostEnterpriseReposCommentsReactionsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments/{comment_id}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseReposCommentsReactionsByOwnerByRepoByCommentIdByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/comments/{comment_id}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposCommitsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits",
	Method:  "GET",
}

var GetEnterpriseReposCommitsBranchesWhereHeadByOwnerByRepoByCommitSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{commit_sha}/branches-where-head",
	Method:  "GET",
}

var GetEnterpriseReposCommitsCommentsByOwnerByRepoByCommitSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{commit_sha}/comments",
	Method:  "GET",
}

var PostEnterpriseReposCommitsCommentsByOwnerByRepoByCommitSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{commit_sha}/comments",
	Method:  "POST",
}

var GetEnterpriseReposCommitsPullsByOwnerByRepoByCommitSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{commit_sha}/pulls",
	Method:  "GET",
}

var GetEnterpriseReposCommitsByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{ref:.+}",
	Method:  "GET",
}

var GetEnterpriseReposCommitsCheckRunsByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{ref}/check-runs",
	Method:  "GET",
}

var GetEnterpriseReposCommitsCheckSuitesByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{ref}/check-suites",
	Method:  "GET",
}

var GetEnterpriseReposCommitsStatusByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{ref}/status",
	Method:  "GET",
}

var GetEnterpriseReposCommitsStatusesByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/commits/{ref}/statuses",
	Method:  "GET",
}

var GetEnterpriseReposCommunityProfileByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/community/profile",
	Method:  "GET",
}

var GetEnterpriseReposCompareByOwnerByRepoByBasehead EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/compare/{basehead}",
	Method:  "GET",
}

var GetEnterpriseReposContentsByOwnerByRepoByPath EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/contents/{path:.+}",
	Method:  "GET",
}

var PutEnterpriseReposContentsByOwnerByRepoByPath EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/contents/{path:.+}",
	Method:  "PUT",
}

var DeleteEnterpriseReposContentsByOwnerByRepoByPath EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/contents/{path:.+}",
	Method:  "DELETE",
}

var GetEnterpriseReposContributorsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/contributors",
	Method:  "GET",
}

var GetEnterpriseReposDependabotAlertsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/alerts",
	Method:  "GET",
}

var GetEnterpriseReposDependabotAlertsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/alerts/{alert_number}",
	Method:  "GET",
}

var PatchEnterpriseReposDependabotAlertsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/alerts/{alert_number}",
	Method:  "PATCH",
}

var GetEnterpriseReposDependabotSecretsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/secrets",
	Method:  "GET",
}

var GetEnterpriseReposDependabotSecretsPublicKeyByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseReposDependabotSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseReposDependabotSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseReposDependabotSecretsByOwnerByRepoBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependabot/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseReposDependencyGraphCompareByOwnerByRepoByBasehead EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependency-graph/compare/{basehead}",
	Method:  "GET",
}

var GetEnterpriseReposDependencyGraphSbomByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependency-graph/sbom",
	Method:  "GET",
}

var PostEnterpriseReposDependencyGraphSnapshotsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dependency-graph/snapshots",
	Method:  "POST",
}

var GetEnterpriseReposDeploymentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments",
	Method:  "GET",
}

var PostEnterpriseReposDeploymentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments",
	Method:  "POST",
}

var GetEnterpriseReposDeploymentsByOwnerByRepoByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments/{deployment_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposDeploymentsByOwnerByRepoByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments/{deployment_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposDeploymentsStatusesByOwnerByRepoByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments/{deployment_id}/statuses",
	Method:  "GET",
}

var PostEnterpriseReposDeploymentsStatusesByOwnerByRepoByDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments/{deployment_id}/statuses",
	Method:  "POST",
}

var GetEnterpriseReposDeploymentsStatusesByOwnerByRepoByDeploymentIdByStatusId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/deployments/{deployment_id}/statuses/{status_id}",
	Method:  "GET",
}

var PostEnterpriseReposDispatchesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/dispatches",
	Method:  "POST",
}

var GetEnterpriseReposEnvironmentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments",
	Method:  "GET",
}

var GetEnterpriseReposEnvironmentsByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}",
	Method:  "GET",
}

var PutEnterpriseReposEnvironmentsByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}",
	Method:  "PUT",
}

var DeleteEnterpriseReposEnvironmentsByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}",
	Method:  "DELETE",
}

var GetEnterpriseReposEnvironmentsDeploymentBranchPoliciesByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies",
	Method:  "GET",
}

var PostEnterpriseReposEnvironmentsDeploymentBranchPoliciesByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies",
	Method:  "POST",
}

var GetEnterpriseReposEnvironmentsDeploymentBranchPoliciesByOwnerByRepoByEnvironmentNameByBranchPolicyId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}",
	Method:  "GET",
}

var PutEnterpriseReposEnvironmentsDeploymentBranchPoliciesByOwnerByRepoByEnvironmentNameByBranchPolicyId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}",
	Method:  "PUT",
}

var DeleteEnterpriseReposEnvironmentsDeploymentBranchPoliciesByOwnerByRepoByEnvironmentNameByBranchPolicyId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment-branch-policies/{branch_policy_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposEnvironmentsDeploymentProtectionRulesByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules",
	Method:  "GET",
}

var PostEnterpriseReposEnvironmentsDeploymentProtectionRulesByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules",
	Method:  "POST",
}

var GetEnterpriseReposEnvironmentsDeploymentProtectionRulesAppsByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/apps",
	Method:  "GET",
}

var GetEnterpriseReposEnvironmentsDeploymentProtectionRulesByOwnerByRepoByEnvironmentNameByProtectionRuleId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/{protection_rule_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposEnvironmentsDeploymentProtectionRulesByOwnerByRepoByEnvironmentNameByProtectionRuleId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/{protection_rule_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposEnvironmentsSecretsByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/secrets",
	Method:  "GET",
}

var GetEnterpriseReposEnvironmentsSecretsPublicKeyByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseReposEnvironmentsSecretsByOwnerByRepoByEnvironmentNameBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseReposEnvironmentsSecretsByOwnerByRepoByEnvironmentNameBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseReposEnvironmentsSecretsByOwnerByRepoByEnvironmentNameBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseReposEnvironmentsVariablesByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/variables",
	Method:  "GET",
}

var PostEnterpriseReposEnvironmentsVariablesByOwnerByRepoByEnvironmentName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/variables",
	Method:  "POST",
}

var GetEnterpriseReposEnvironmentsVariablesByOwnerByRepoByEnvironmentNameByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/variables/{name}",
	Method:  "GET",
}

var PatchEnterpriseReposEnvironmentsVariablesByOwnerByRepoByEnvironmentNameByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/variables/{name}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposEnvironmentsVariablesByOwnerByRepoByEnvironmentNameByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/environments/{environment_name}/variables/{name}",
	Method:  "DELETE",
}

var GetEnterpriseReposEventsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/events",
	Method:  "GET",
}

var GetEnterpriseReposForksByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/forks",
	Method:  "GET",
}

var PostEnterpriseReposForksByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/forks",
	Method:  "POST",
}

var PostEnterpriseReposGitBlobsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/blobs",
	Method:  "POST",
}

var GetEnterpriseReposGitBlobsByOwnerByRepoByFileSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/blobs/{file_sha}",
	Method:  "GET",
}

var PostEnterpriseReposGitCommitsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/commits",
	Method:  "POST",
}

var GetEnterpriseReposGitCommitsByOwnerByRepoByCommitSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/commits/{commit_sha}",
	Method:  "GET",
}

var GetEnterpriseReposGitMatchingRefsByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/matching-refs/{ref}",
	Method:  "GET",
}

var GetEnterpriseReposGitRefByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/ref/{ref:.+}",
	Method:  "GET",
}

var PostEnterpriseReposGitRefsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/refs",
	Method:  "POST",
}

var PatchEnterpriseReposGitRefsByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/refs/{ref:.+}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposGitRefsByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/refs/{ref:.+}",
	Method:  "DELETE",
}

var PostEnterpriseReposGitTagsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/tags",
	Method:  "POST",
}

var GetEnterpriseReposGitTagsByOwnerByRepoByTagSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/tags/{tag_sha}",
	Method:  "GET",
}

var PostEnterpriseReposGitTreesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/trees",
	Method:  "POST",
}

var GetEnterpriseReposGitTreesByOwnerByRepoByTreeSha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/git/trees/{tree_sha}",
	Method:  "GET",
}

var GetEnterpriseReposHooksByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks",
	Method:  "GET",
}

var PostEnterpriseReposHooksByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks",
	Method:  "POST",
}

var GetEnterpriseReposHooksByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}",
	Method:  "GET",
}

var PatchEnterpriseReposHooksByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposHooksByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposHooksConfigByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/config",
	Method:  "GET",
}

var PatchEnterpriseReposHooksConfigByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/config",
	Method:  "PATCH",
}

var GetEnterpriseReposHooksDeliveriesByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/deliveries",
	Method:  "GET",
}

var GetEnterpriseReposHooksDeliveriesByOwnerByRepoByHookIdByDeliveryId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}",
	Method:  "GET",
}

var PostEnterpriseReposHooksDeliveriesAttemptsByOwnerByRepoByHookIdByDeliveryId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}/attempts",
	Method:  "POST",
}

var PostEnterpriseReposHooksPingsByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/pings",
	Method:  "POST",
}

var PostEnterpriseReposHooksTestsByOwnerByRepoByHookId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/hooks/{hook_id}/tests",
	Method:  "POST",
}

var GetEnterpriseReposImportByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import",
	Method:  "GET",
}

var PutEnterpriseReposImportByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import",
	Method:  "PUT",
}

var PatchEnterpriseReposImportByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import",
	Method:  "PATCH",
}

var DeleteEnterpriseReposImportByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import",
	Method:  "DELETE",
}

var GetEnterpriseReposImportAuthorsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import/authors",
	Method:  "GET",
}

var PatchEnterpriseReposImportAuthorsByOwnerByRepoByAuthorId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import/authors/{author_id}",
	Method:  "PATCH",
}

var GetEnterpriseReposImportLargeFilesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import/large_files",
	Method:  "GET",
}

var PatchEnterpriseReposImportLfsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/import/lfs",
	Method:  "PATCH",
}

var GetEnterpriseReposInstallationByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/installation",
	Method:  "GET",
}

var GetEnterpriseReposInteractionLimitsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/interaction-limits",
	Method:  "GET",
}

var PutEnterpriseReposInteractionLimitsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/interaction-limits",
	Method:  "PUT",
}

var DeleteEnterpriseReposInteractionLimitsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/interaction-limits",
	Method:  "DELETE",
}

var GetEnterpriseReposInvitationsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/invitations",
	Method:  "GET",
}

var PatchEnterpriseReposInvitationsByOwnerByRepoByInvitationId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/invitations/{invitation_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposInvitationsByOwnerByRepoByInvitationId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/invitations/{invitation_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposIssuesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues",
	Method:  "GET",
}

var PostEnterpriseReposIssuesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues",
	Method:  "POST",
}

var GetEnterpriseReposIssuesCommentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments",
	Method:  "GET",
}

var GetEnterpriseReposIssuesCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments/{comment_id}",
	Method:  "GET",
}

var PatchEnterpriseReposIssuesCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments/{comment_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposIssuesCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments/{comment_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposIssuesCommentsReactionsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments/{comment_id}/reactions",
	Method:  "GET",
}

var PostEnterpriseReposIssuesCommentsReactionsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments/{comment_id}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseReposIssuesCommentsReactionsByOwnerByRepoByCommentIdByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/comments/{comment_id}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposIssuesEventsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/events",
	Method:  "GET",
}

var GetEnterpriseReposIssuesEventsByOwnerByRepoByEventId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/events/{event_id}",
	Method:  "GET",
}

var GetEnterpriseReposIssuesByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}",
	Method:  "GET",
}

var PatchEnterpriseReposIssuesByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}",
	Method:  "PATCH",
}

var PostEnterpriseReposIssuesAssigneesByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/assignees",
	Method:  "POST",
}

var DeleteEnterpriseReposIssuesAssigneesByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/assignees",
	Method:  "DELETE",
}

var GetEnterpriseReposIssuesAssigneesByOwnerByRepoByIssueNumberByAssignee EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/assignees/{assignee}",
	Method:  "GET",
}

var GetEnterpriseReposIssuesCommentsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/comments",
	Method:  "GET",
}

var PostEnterpriseReposIssuesCommentsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/comments",
	Method:  "POST",
}

var GetEnterpriseReposIssuesEventsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/events",
	Method:  "GET",
}

var GetEnterpriseReposIssuesLabelsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/labels",
	Method:  "GET",
}

var PostEnterpriseReposIssuesLabelsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/labels",
	Method:  "POST",
}

var PutEnterpriseReposIssuesLabelsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/labels",
	Method:  "PUT",
}

var DeleteEnterpriseReposIssuesLabelsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/labels",
	Method:  "DELETE",
}

var DeleteEnterpriseReposIssuesLabelsByOwnerByRepoByIssueNumberByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/labels/{name:.+}",
	Method:  "DELETE",
}

var PutEnterpriseReposIssuesLockByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/lock",
	Method:  "PUT",
}

var DeleteEnterpriseReposIssuesLockByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/lock",
	Method:  "DELETE",
}

var GetEnterpriseReposIssuesReactionsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/reactions",
	Method:  "GET",
}

var PostEnterpriseReposIssuesReactionsByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseReposIssuesReactionsByOwnerByRepoByIssueNumberByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposIssuesTimelineByOwnerByRepoByIssueNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/issues/{issue_number}/timeline",
	Method:  "GET",
}

var GetEnterpriseReposKeysByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/keys",
	Method:  "GET",
}

var PostEnterpriseReposKeysByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/keys",
	Method:  "POST",
}

var GetEnterpriseReposKeysByOwnerByRepoByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/keys/{key_id}",
	Method:  "GET",
}

var DeleteEnterpriseReposKeysByOwnerByRepoByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/keys/{key_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposLabelsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/labels",
	Method:  "GET",
}

var PostEnterpriseReposLabelsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/labels",
	Method:  "POST",
}

var GetEnterpriseReposLabelsByOwnerByRepoByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/labels/{name:.+}",
	Method:  "GET",
}

var PatchEnterpriseReposLabelsByOwnerByRepoByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/labels/{name:.+}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposLabelsByOwnerByRepoByName EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/labels/{name:.+}",
	Method:  "DELETE",
}

var GetEnterpriseReposLanguagesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/languages",
	Method:  "GET",
}

var PutEnterpriseReposLfsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/lfs",
	Method:  "PUT",
}

var DeleteEnterpriseReposLfsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/lfs",
	Method:  "DELETE",
}

var GetEnterpriseReposLicenseByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/license",
	Method:  "GET",
}

var PostEnterpriseReposMergeUpstreamByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/merge-upstream",
	Method:  "POST",
}

var PostEnterpriseReposMergesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/merges",
	Method:  "POST",
}

var GetEnterpriseReposMilestonesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/milestones",
	Method:  "GET",
}

var PostEnterpriseReposMilestonesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/milestones",
	Method:  "POST",
}

var GetEnterpriseReposMilestonesByOwnerByRepoByMilestoneNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/milestones/{milestone_number}",
	Method:  "GET",
}

var PatchEnterpriseReposMilestonesByOwnerByRepoByMilestoneNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/milestones/{milestone_number}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposMilestonesByOwnerByRepoByMilestoneNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/milestones/{milestone_number}",
	Method:  "DELETE",
}

var GetEnterpriseReposMilestonesLabelsByOwnerByRepoByMilestoneNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/milestones/{milestone_number}/labels",
	Method:  "GET",
}

var GetEnterpriseReposNotificationsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/notifications",
	Method:  "GET",
}

var PutEnterpriseReposNotificationsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/notifications",
	Method:  "PUT",
}

var GetEnterpriseReposPagesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages",
	Method:  "GET",
}

var PostEnterpriseReposPagesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages",
	Method:  "POST",
}

var PutEnterpriseReposPagesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages",
	Method:  "PUT",
}

var DeleteEnterpriseReposPagesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages",
	Method:  "DELETE",
}

var GetEnterpriseReposPagesBuildsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/builds",
	Method:  "GET",
}

var PostEnterpriseReposPagesBuildsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/builds",
	Method:  "POST",
}

var GetEnterpriseReposPagesBuildsLatestByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/builds/latest",
	Method:  "GET",
}

var GetEnterpriseReposPagesBuildsByOwnerByRepoByBuildId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/builds/{build_id}",
	Method:  "GET",
}

var PostEnterpriseReposPagesDeploymentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/deployments",
	Method:  "POST",
}

var GetEnterpriseReposPagesDeploymentsByOwnerByRepoByPagesDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/deployments/{pages_deployment_id}",
	Method:  "GET",
}

var PostEnterpriseReposPagesDeploymentsCancelByOwnerByRepoByPagesDeploymentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/deployments/{pages_deployment_id}/cancel",
	Method:  "POST",
}

var GetEnterpriseReposPagesHealthByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pages/health",
	Method:  "GET",
}

var GetEnterpriseReposPrivateVulnerabilityReportingByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/private-vulnerability-reporting",
	Method:  "GET",
}

var PutEnterpriseReposPrivateVulnerabilityReportingByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/private-vulnerability-reporting",
	Method:  "PUT",
}

var DeleteEnterpriseReposPrivateVulnerabilityReportingByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/private-vulnerability-reporting",
	Method:  "DELETE",
}

var GetEnterpriseReposProjectsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/projects",
	Method:  "GET",
}

var PostEnterpriseReposProjectsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/projects",
	Method:  "POST",
}

var GetEnterpriseReposPropertiesValuesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/properties/values",
	Method:  "GET",
}

var PatchEnterpriseReposPropertiesValuesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/properties/values",
	Method:  "PATCH",
}

var GetEnterpriseReposPullsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls",
	Method:  "GET",
}

var PostEnterpriseReposPullsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls",
	Method:  "POST",
}

var GetEnterpriseReposPullsCommentsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments",
	Method:  "GET",
}

var GetEnterpriseReposPullsCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments/{comment_id}",
	Method:  "GET",
}

var PatchEnterpriseReposPullsCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments/{comment_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposPullsCommentsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments/{comment_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposPullsCommentsReactionsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions",
	Method:  "GET",
}

var PostEnterpriseReposPullsCommentsReactionsByOwnerByRepoByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseReposPullsCommentsReactionsByOwnerByRepoByCommentIdByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposPullsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}",
	Method:  "GET",
}

var PatchEnterpriseReposPullsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}",
	Method:  "PATCH",
}

var PostEnterpriseReposPullsCodespacesByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/codespaces",
	Method:  "POST",
}

var GetEnterpriseReposPullsCommentsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/comments",
	Method:  "GET",
}

var PostEnterpriseReposPullsCommentsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/comments",
	Method:  "POST",
}

var PostEnterpriseReposPullsCommentsRepliesByOwnerByRepoByPullNumberByCommentId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/comments/{comment_id}/replies",
	Method:  "POST",
}

var GetEnterpriseReposPullsCommitsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/commits",
	Method:  "GET",
}

var GetEnterpriseReposPullsFilesByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/files",
	Method:  "GET",
}

var GetEnterpriseReposPullsMergeByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/merge",
	Method:  "GET",
}

var PutEnterpriseReposPullsMergeByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/merge",
	Method:  "PUT",
}

var GetEnterpriseReposPullsRequestedReviewersByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers",
	Method:  "GET",
}

var PostEnterpriseReposPullsRequestedReviewersByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers",
	Method:  "POST",
}

var DeleteEnterpriseReposPullsRequestedReviewersByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers",
	Method:  "DELETE",
}

var GetEnterpriseReposPullsReviewsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews",
	Method:  "GET",
}

var PostEnterpriseReposPullsReviewsByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews",
	Method:  "POST",
}

var GetEnterpriseReposPullsReviewsByOwnerByRepoByPullNumberByReviewId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}",
	Method:  "GET",
}

var PutEnterpriseReposPullsReviewsByOwnerByRepoByPullNumberByReviewId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}",
	Method:  "PUT",
}

var DeleteEnterpriseReposPullsReviewsByOwnerByRepoByPullNumberByReviewId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposPullsReviewsCommentsByOwnerByRepoByPullNumberByReviewId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/comments",
	Method:  "GET",
}

var PutEnterpriseReposPullsReviewsDismissalsByOwnerByRepoByPullNumberByReviewId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/dismissals",
	Method:  "PUT",
}

var PostEnterpriseReposPullsReviewsEventsByOwnerByRepoByPullNumberByReviewId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/events",
	Method:  "POST",
}

var PutEnterpriseReposPullsUpdateBranchByOwnerByRepoByPullNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/pulls/{pull_number}/update-branch",
	Method:  "PUT",
}

var GetEnterpriseReposReadmeByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/readme",
	Method:  "GET",
}

var GetEnterpriseReposReadmeByOwnerByRepoByDir EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/readme/{dir}",
	Method:  "GET",
}

var GetEnterpriseReposReleasesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases",
	Method:  "GET",
}

var PostEnterpriseReposReleasesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases",
	Method:  "POST",
}

var GetEnterpriseReposReleasesAssetsByOwnerByRepoByAssetId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/assets/{asset_id}",
	Method:  "GET",
}

var PatchEnterpriseReposReleasesAssetsByOwnerByRepoByAssetId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/assets/{asset_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposReleasesAssetsByOwnerByRepoByAssetId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/assets/{asset_id}",
	Method:  "DELETE",
}

var PostEnterpriseReposReleasesGenerateNotesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/generate-notes",
	Method:  "POST",
}

var GetEnterpriseReposReleasesLatestByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/latest",
	Method:  "GET",
}

var GetEnterpriseReposReleasesTagsByOwnerByRepoByTag EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/tags/{tag}",
	Method:  "GET",
}

var GetEnterpriseReposReleasesByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}",
	Method:  "GET",
}

var PatchEnterpriseReposReleasesByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseReposReleasesByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposReleasesAssetsByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}/assets",
	Method:  "GET",
}

var PostEnterpriseReposReleasesAssetsByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}/assets",
	Method:  "POST",
}

var GetEnterpriseReposReleasesReactionsByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}/reactions",
	Method:  "GET",
}

var PostEnterpriseReposReleasesReactionsByOwnerByRepoByReleaseId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}/reactions",
	Method:  "POST",
}

var DeleteEnterpriseReposReleasesReactionsByOwnerByRepoByReleaseIdByReactionId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/releases/{release_id}/reactions/{reaction_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposRulesBranchesByOwnerByRepoByBranch EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rules/branches/{branch}",
	Method:  "GET",
}

var GetEnterpriseReposRulesetsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets",
	Method:  "GET",
}

var PostEnterpriseReposRulesetsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets",
	Method:  "POST",
}

var GetEnterpriseReposRulesetsRuleSuitesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets/rule-suites",
	Method:  "GET",
}

var GetEnterpriseReposRulesetsRuleSuitesByOwnerByRepoByRuleSuiteId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets/rule-suites/{rule_suite_id}",
	Method:  "GET",
}

var GetEnterpriseReposRulesetsByOwnerByRepoByRulesetId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets/{ruleset_id}",
	Method:  "GET",
}

var PutEnterpriseReposRulesetsByOwnerByRepoByRulesetId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets/{ruleset_id}",
	Method:  "PUT",
}

var DeleteEnterpriseReposRulesetsByOwnerByRepoByRulesetId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/rulesets/{ruleset_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposSecretScanningAlertsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/secret-scanning/alerts",
	Method:  "GET",
}

var GetEnterpriseReposSecretScanningAlertsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}",
	Method:  "GET",
}

var PatchEnterpriseReposSecretScanningAlertsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}",
	Method:  "PATCH",
}

var GetEnterpriseReposSecretScanningAlertsLocationsByOwnerByRepoByAlertNumber EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}/locations",
	Method:  "GET",
}

var PostEnterpriseReposSecretScanningPushProtectionBypassesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/secret-scanning/push-protection-bypasses",
	Method:  "POST",
}

var GetEnterpriseReposSecurityAdvisoriesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories",
	Method:  "GET",
}

var PostEnterpriseReposSecurityAdvisoriesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories",
	Method:  "POST",
}

var PostEnterpriseReposSecurityAdvisoriesReportsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories/reports",
	Method:  "POST",
}

var GetEnterpriseReposSecurityAdvisoriesByOwnerByRepoByGhsaId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories/{ghsa_id}",
	Method:  "GET",
}

var PatchEnterpriseReposSecurityAdvisoriesByOwnerByRepoByGhsaId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories/{ghsa_id}",
	Method:  "PATCH",
}

var PostEnterpriseReposSecurityAdvisoriesCveByOwnerByRepoByGhsaId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories/{ghsa_id}/cve",
	Method:  "POST",
}

var PostEnterpriseReposSecurityAdvisoriesForksByOwnerByRepoByGhsaId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/security-advisories/{ghsa_id}/forks",
	Method:  "POST",
}

var GetEnterpriseReposStargazersByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/stargazers",
	Method:  "GET",
}

var GetEnterpriseReposStatsCodeFrequencyByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/stats/code_frequency",
	Method:  "GET",
}

var GetEnterpriseReposStatsCommitActivityByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/stats/commit_activity",
	Method:  "GET",
}

var GetEnterpriseReposStatsContributorsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/stats/contributors",
	Method:  "GET",
}

var GetEnterpriseReposStatsParticipationByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/stats/participation",
	Method:  "GET",
}

var GetEnterpriseReposStatsPunchCardByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/stats/punch_card",
	Method:  "GET",
}

var PostEnterpriseReposStatusesByOwnerByRepoBySha EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/statuses/{sha}",
	Method:  "POST",
}

var GetEnterpriseReposSubscribersByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/subscribers",
	Method:  "GET",
}

var GetEnterpriseReposSubscriptionByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/subscription",
	Method:  "GET",
}

var PutEnterpriseReposSubscriptionByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/subscription",
	Method:  "PUT",
}

var DeleteEnterpriseReposSubscriptionByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/subscription",
	Method:  "DELETE",
}

var GetEnterpriseReposTagsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/tags",
	Method:  "GET",
}

var GetEnterpriseReposTagsProtectionByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/tags/protection",
	Method:  "GET",
}

var PostEnterpriseReposTagsProtectionByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/tags/protection",
	Method:  "POST",
}

var DeleteEnterpriseReposTagsProtectionByOwnerByRepoByTagProtectionId EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/tags/protection/{tag_protection_id}",
	Method:  "DELETE",
}

var GetEnterpriseReposTarballByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/tarball/{ref}",
	Method:  "GET",
}

var GetEnterpriseReposTeamsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/teams",
	Method:  "GET",
}

var GetEnterpriseReposTopicsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/topics",
	Method:  "GET",
}

var PutEnterpriseReposTopicsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/topics",
	Method:  "PUT",
}

var GetEnterpriseReposTrafficClonesByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/traffic/clones",
	Method:  "GET",
}

var GetEnterpriseReposTrafficPopularPathsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/traffic/popular/paths",
	Method:  "GET",
}

var GetEnterpriseReposTrafficPopularReferrersByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/traffic/popular/referrers",
	Method:  "GET",
}

var GetEnterpriseReposTrafficViewsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/traffic/views",
	Method:  "GET",
}

var PostEnterpriseReposTransferByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/transfer",
	Method:  "POST",
}

var GetEnterpriseReposVulnerabilityAlertsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/vulnerability-alerts",
	Method:  "GET",
}

var PutEnterpriseReposVulnerabilityAlertsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/vulnerability-alerts",
	Method:  "PUT",
}

var DeleteEnterpriseReposVulnerabilityAlertsByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/vulnerability-alerts",
	Method:  "DELETE",
}

var GetEnterpriseReposZipballByOwnerByRepoByRef EndpointPattern = EndpointPattern{
	Pattern: "/repos/{owner}/{repo}/zipball/{ref}",
	Method:  "GET",
}

var PostEnterpriseReposGenerateByTemplateOwnerByTemplateRepo EndpointPattern = EndpointPattern{
	Pattern: "/repos/{template_owner}/{template_repo}/generate",
	Method:  "POST",
}

var GetEnterpriseRepositories EndpointPattern = EndpointPattern{
	Pattern: "/repositories",
	Method:  "GET",
}

var GetEnterpriseScimV2GroupsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Groups",
	Method:  "GET",
}

var PostEnterpriseScimV2GroupsByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Groups",
	Method:  "POST",
}

var GetEnterpriseScimV2GroupsByEnterpriseByScimGroupId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}",
	Method:  "GET",
}

var PutEnterpriseScimV2GroupsByEnterpriseByScimGroupId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}",
	Method:  "PUT",
}

var PatchEnterpriseScimV2GroupsByEnterpriseByScimGroupId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseScimV2GroupsByEnterpriseByScimGroupId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}",
	Method:  "DELETE",
}

var GetEnterpriseScimV2UsersByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Users",
	Method:  "GET",
}

var PostEnterpriseScimV2UsersByEnterprise EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Users",
	Method:  "POST",
}

var GetEnterpriseScimV2UsersByEnterpriseByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}",
	Method:  "GET",
}

var PutEnterpriseScimV2UsersByEnterpriseByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}",
	Method:  "PUT",
}

var PatchEnterpriseScimV2UsersByEnterpriseByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseScimV2UsersByEnterpriseByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/enterprises/{enterprise}/Users/{scim_user_id}",
	Method:  "DELETE",
}

var GetEnterpriseScimV2OrganizationsUsersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/organizations/{org}/Users",
	Method:  "GET",
}

var PostEnterpriseScimV2OrganizationsUsersByOrg EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/organizations/{org}/Users",
	Method:  "POST",
}

var GetEnterpriseScimV2OrganizationsUsersByOrgByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/organizations/{org}/Users/{scim_user_id}",
	Method:  "GET",
}

var PutEnterpriseScimV2OrganizationsUsersByOrgByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/organizations/{org}/Users/{scim_user_id}",
	Method:  "PUT",
}

var PatchEnterpriseScimV2OrganizationsUsersByOrgByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/organizations/{org}/Users/{scim_user_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseScimV2OrganizationsUsersByOrgByScimUserId EndpointPattern = EndpointPattern{
	Pattern: "/scim/v2/organizations/{org}/Users/{scim_user_id}",
	Method:  "DELETE",
}

var GetEnterpriseSearchCode EndpointPattern = EndpointPattern{
	Pattern: "/search/code",
	Method:  "GET",
}

var GetEnterpriseSearchCommits EndpointPattern = EndpointPattern{
	Pattern: "/search/commits",
	Method:  "GET",
}

var GetEnterpriseSearchIssues EndpointPattern = EndpointPattern{
	Pattern: "/search/issues",
	Method:  "GET",
}

var GetEnterpriseSearchLabels EndpointPattern = EndpointPattern{
	Pattern: "/search/labels",
	Method:  "GET",
}

var GetEnterpriseSearchRepositories EndpointPattern = EndpointPattern{
	Pattern: "/search/repositories",
	Method:  "GET",
}

var GetEnterpriseSearchTopics EndpointPattern = EndpointPattern{
	Pattern: "/search/topics",
	Method:  "GET",
}

var GetEnterpriseSearchUsers EndpointPattern = EndpointPattern{
	Pattern: "/search/users",
	Method:  "GET",
}

var GetEnterpriseTeamsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}",
	Method:  "GET",
}

var PatchEnterpriseTeamsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseTeamsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsDiscussionsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions",
	Method:  "GET",
}

var PostEnterpriseTeamsDiscussionsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions",
	Method:  "POST",
}

var GetEnterpriseTeamsDiscussionsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}",
	Method:  "GET",
}

var PatchEnterpriseTeamsDiscussionsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}",
	Method:  "PATCH",
}

var DeleteEnterpriseTeamsDiscussionsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsDiscussionsCommentsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments",
	Method:  "GET",
}

var PostEnterpriseTeamsDiscussionsCommentsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments",
	Method:  "POST",
}

var GetEnterpriseTeamsDiscussionsCommentsByTeamIdByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}",
	Method:  "GET",
}

var PatchEnterpriseTeamsDiscussionsCommentsByTeamIdByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}",
	Method:  "PATCH",
}

var DeleteEnterpriseTeamsDiscussionsCommentsByTeamIdByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsDiscussionsCommentsReactionsByTeamIdByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions",
	Method:  "GET",
}

var PostEnterpriseTeamsDiscussionsCommentsReactionsByTeamIdByDiscussionNumberByCommentNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions",
	Method:  "POST",
}

var GetEnterpriseTeamsDiscussionsReactionsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/reactions",
	Method:  "GET",
}

var PostEnterpriseTeamsDiscussionsReactionsByTeamIdByDiscussionNumber EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/discussions/{discussion_number}/reactions",
	Method:  "POST",
}

var GetEnterpriseTeamsInvitationsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/invitations",
	Method:  "GET",
}

var GetEnterpriseTeamsMembersByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/members",
	Method:  "GET",
}

var GetEnterpriseTeamsMembersByTeamIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/members/{username}",
	Method:  "GET",
}

var PutEnterpriseTeamsMembersByTeamIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/members/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseTeamsMembersByTeamIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/members/{username}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsMembershipsByTeamIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/memberships/{username}",
	Method:  "GET",
}

var PutEnterpriseTeamsMembershipsByTeamIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/memberships/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseTeamsMembershipsByTeamIdByUsername EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/memberships/{username}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsProjectsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/projects",
	Method:  "GET",
}

var GetEnterpriseTeamsProjectsByTeamIdByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/projects/{project_id}",
	Method:  "GET",
}

var PutEnterpriseTeamsProjectsByTeamIdByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/projects/{project_id}",
	Method:  "PUT",
}

var DeleteEnterpriseTeamsProjectsByTeamIdByProjectId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/projects/{project_id}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsReposByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/repos",
	Method:  "GET",
}

var GetEnterpriseTeamsReposByTeamIdByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/repos/{owner}/{repo}",
	Method:  "GET",
}

var PutEnterpriseTeamsReposByTeamIdByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/repos/{owner}/{repo}",
	Method:  "PUT",
}

var DeleteEnterpriseTeamsReposByTeamIdByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/repos/{owner}/{repo}",
	Method:  "DELETE",
}

var GetEnterpriseTeamsTeamSyncGroupMappingsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/team-sync/group-mappings",
	Method:  "GET",
}

var PatchEnterpriseTeamsTeamSyncGroupMappingsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/team-sync/group-mappings",
	Method:  "PATCH",
}

var GetEnterpriseTeamsTeamsByTeamId EndpointPattern = EndpointPattern{
	Pattern: "/teams/{team_id}/teams",
	Method:  "GET",
}

var GetEnterpriseUser EndpointPattern = EndpointPattern{
	Pattern: "/user",
	Method:  "GET",
}

var PatchEnterpriseUser EndpointPattern = EndpointPattern{
	Pattern: "/user",
	Method:  "PATCH",
}

var GetEnterpriseUserBlocks EndpointPattern = EndpointPattern{
	Pattern: "/user/blocks",
	Method:  "GET",
}

var GetEnterpriseUserBlocksByUsername EndpointPattern = EndpointPattern{
	Pattern: "/user/blocks/{username}",
	Method:  "GET",
}

var PutEnterpriseUserBlocksByUsername EndpointPattern = EndpointPattern{
	Pattern: "/user/blocks/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseUserBlocksByUsername EndpointPattern = EndpointPattern{
	Pattern: "/user/blocks/{username}",
	Method:  "DELETE",
}

var GetEnterpriseUserCodespaces EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces",
	Method:  "GET",
}

var PostEnterpriseUserCodespaces EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces",
	Method:  "POST",
}

var GetEnterpriseUserCodespacesSecrets EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets",
	Method:  "GET",
}

var GetEnterpriseUserCodespacesSecretsPublicKey EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/public-key",
	Method:  "GET",
}

var GetEnterpriseUserCodespacesSecretsBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}",
	Method:  "GET",
}

var PutEnterpriseUserCodespacesSecretsBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}",
	Method:  "PUT",
}

var DeleteEnterpriseUserCodespacesSecretsBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}",
	Method:  "DELETE",
}

var GetEnterpriseUserCodespacesSecretsRepositoriesBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}/repositories",
	Method:  "GET",
}

var PutEnterpriseUserCodespacesSecretsRepositoriesBySecretName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}/repositories",
	Method:  "PUT",
}

var PutEnterpriseUserCodespacesSecretsRepositoriesBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseUserCodespacesSecretsRepositoriesBySecretNameByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/secrets/{secret_name}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseUserCodespacesByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}",
	Method:  "GET",
}

var PatchEnterpriseUserCodespacesByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}",
	Method:  "PATCH",
}

var DeleteEnterpriseUserCodespacesByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}",
	Method:  "DELETE",
}

var PostEnterpriseUserCodespacesExportsByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}/exports",
	Method:  "POST",
}

var GetEnterpriseUserCodespacesExportsByCodespaceNameByExportId EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}/exports/{export_id}",
	Method:  "GET",
}

var GetEnterpriseUserCodespacesMachinesByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}/machines",
	Method:  "GET",
}

var PostEnterpriseUserCodespacesPublishByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}/publish",
	Method:  "POST",
}

var PostEnterpriseUserCodespacesStartByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}/start",
	Method:  "POST",
}

var PostEnterpriseUserCodespacesStopByCodespaceName EndpointPattern = EndpointPattern{
	Pattern: "/user/codespaces/{codespace_name}/stop",
	Method:  "POST",
}

var GetEnterpriseUserDockerConflicts EndpointPattern = EndpointPattern{
	Pattern: "/user/docker/conflicts",
	Method:  "GET",
}

var PatchEnterpriseUserEmailVisibility EndpointPattern = EndpointPattern{
	Pattern: "/user/email/visibility",
	Method:  "PATCH",
}

var GetEnterpriseUserEmails EndpointPattern = EndpointPattern{
	Pattern: "/user/emails",
	Method:  "GET",
}

var PostEnterpriseUserEmails EndpointPattern = EndpointPattern{
	Pattern: "/user/emails",
	Method:  "POST",
}

var DeleteEnterpriseUserEmails EndpointPattern = EndpointPattern{
	Pattern: "/user/emails",
	Method:  "DELETE",
}

var GetEnterpriseUserFollowers EndpointPattern = EndpointPattern{
	Pattern: "/user/followers",
	Method:  "GET",
}

var GetEnterpriseUserFollowing EndpointPattern = EndpointPattern{
	Pattern: "/user/following",
	Method:  "GET",
}

var GetEnterpriseUserFollowingByUsername EndpointPattern = EndpointPattern{
	Pattern: "/user/following/{username}",
	Method:  "GET",
}

var PutEnterpriseUserFollowingByUsername EndpointPattern = EndpointPattern{
	Pattern: "/user/following/{username}",
	Method:  "PUT",
}

var DeleteEnterpriseUserFollowingByUsername EndpointPattern = EndpointPattern{
	Pattern: "/user/following/{username}",
	Method:  "DELETE",
}

var GetEnterpriseUserGpgKeys EndpointPattern = EndpointPattern{
	Pattern: "/user/gpg_keys",
	Method:  "GET",
}

var PostEnterpriseUserGpgKeys EndpointPattern = EndpointPattern{
	Pattern: "/user/gpg_keys",
	Method:  "POST",
}

var GetEnterpriseUserGpgKeysByGpgKeyId EndpointPattern = EndpointPattern{
	Pattern: "/user/gpg_keys/{gpg_key_id}",
	Method:  "GET",
}

var DeleteEnterpriseUserGpgKeysByGpgKeyId EndpointPattern = EndpointPattern{
	Pattern: "/user/gpg_keys/{gpg_key_id}",
	Method:  "DELETE",
}

var GetEnterpriseUserInstallations EndpointPattern = EndpointPattern{
	Pattern: "/user/installations",
	Method:  "GET",
}

var GetEnterpriseUserInstallationsRepositoriesByInstallationId EndpointPattern = EndpointPattern{
	Pattern: "/user/installations/{installation_id}/repositories",
	Method:  "GET",
}

var PutEnterpriseUserInstallationsRepositoriesByInstallationIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/user/installations/{installation_id}/repositories/{repository_id}",
	Method:  "PUT",
}

var DeleteEnterpriseUserInstallationsRepositoriesByInstallationIdByRepositoryId EndpointPattern = EndpointPattern{
	Pattern: "/user/installations/{installation_id}/repositories/{repository_id}",
	Method:  "DELETE",
}

var GetEnterpriseUserInteractionLimits EndpointPattern = EndpointPattern{
	Pattern: "/user/interaction-limits",
	Method:  "GET",
}

var PutEnterpriseUserInteractionLimits EndpointPattern = EndpointPattern{
	Pattern: "/user/interaction-limits",
	Method:  "PUT",
}

var DeleteEnterpriseUserInteractionLimits EndpointPattern = EndpointPattern{
	Pattern: "/user/interaction-limits",
	Method:  "DELETE",
}

var GetEnterpriseUserIssues EndpointPattern = EndpointPattern{
	Pattern: "/user/issues",
	Method:  "GET",
}

var GetEnterpriseUserKeys EndpointPattern = EndpointPattern{
	Pattern: "/user/keys",
	Method:  "GET",
}

var PostEnterpriseUserKeys EndpointPattern = EndpointPattern{
	Pattern: "/user/keys",
	Method:  "POST",
}

var GetEnterpriseUserKeysByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/user/keys/{key_id}",
	Method:  "GET",
}

var DeleteEnterpriseUserKeysByKeyId EndpointPattern = EndpointPattern{
	Pattern: "/user/keys/{key_id}",
	Method:  "DELETE",
}

var GetEnterpriseUserMarketplacePurchases EndpointPattern = EndpointPattern{
	Pattern: "/user/marketplace_purchases",
	Method:  "GET",
}

var GetEnterpriseUserMarketplacePurchasesStubbed EndpointPattern = EndpointPattern{
	Pattern: "/user/marketplace_purchases/stubbed",
	Method:  "GET",
}

var GetEnterpriseUserMembershipsOrgs EndpointPattern = EndpointPattern{
	Pattern: "/user/memberships/orgs",
	Method:  "GET",
}

var GetEnterpriseUserMembershipsOrgsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/user/memberships/orgs/{org}",
	Method:  "GET",
}

var PatchEnterpriseUserMembershipsOrgsByOrg EndpointPattern = EndpointPattern{
	Pattern: "/user/memberships/orgs/{org}",
	Method:  "PATCH",
}

var GetEnterpriseUserMigrations EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations",
	Method:  "GET",
}

var PostEnterpriseUserMigrations EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations",
	Method:  "POST",
}

var GetEnterpriseUserMigrationsByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations/{migration_id}",
	Method:  "GET",
}

var GetEnterpriseUserMigrationsArchiveByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations/{migration_id}/archive",
	Method:  "GET",
}

var DeleteEnterpriseUserMigrationsArchiveByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations/{migration_id}/archive",
	Method:  "DELETE",
}

var DeleteEnterpriseUserMigrationsReposLockByMigrationIdByRepoName EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations/{migration_id}/repos/{repo_name}/lock",
	Method:  "DELETE",
}

var GetEnterpriseUserMigrationsRepositoriesByMigrationId EndpointPattern = EndpointPattern{
	Pattern: "/user/migrations/{migration_id}/repositories",
	Method:  "GET",
}

var GetEnterpriseUserOrgs EndpointPattern = EndpointPattern{
	Pattern: "/user/orgs",
	Method:  "GET",
}

var GetEnterpriseUserPackages EndpointPattern = EndpointPattern{
	Pattern: "/user/packages",
	Method:  "GET",
}

var GetEnterpriseUserPackagesByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}",
	Method:  "GET",
}

var DeleteEnterpriseUserPackagesByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}",
	Method:  "DELETE",
}

var PostEnterpriseUserPackagesRestoreByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}/restore",
	Method:  "POST",
}

var GetEnterpriseUserPackagesVersionsByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}/versions",
	Method:  "GET",
}

var GetEnterpriseUserPackagesVersionsByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}/versions/{package_version_id}",
	Method:  "GET",
}

var DeleteEnterpriseUserPackagesVersionsByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}/versions/{package_version_id}",
	Method:  "DELETE",
}

var PostEnterpriseUserPackagesVersionsRestoreByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/user/packages/{package_type}/{package_name}/versions/{package_version_id}/restore",
	Method:  "POST",
}

var PostEnterpriseUserProjects EndpointPattern = EndpointPattern{
	Pattern: "/user/projects",
	Method:  "POST",
}

var GetEnterpriseUserPublicEmails EndpointPattern = EndpointPattern{
	Pattern: "/user/public_emails",
	Method:  "GET",
}

var GetEnterpriseUserRepos EndpointPattern = EndpointPattern{
	Pattern: "/user/repos",
	Method:  "GET",
}

var PostEnterpriseUserRepos EndpointPattern = EndpointPattern{
	Pattern: "/user/repos",
	Method:  "POST",
}

var GetEnterpriseUserRepositoryInvitations EndpointPattern = EndpointPattern{
	Pattern: "/user/repository_invitations",
	Method:  "GET",
}

var PatchEnterpriseUserRepositoryInvitationsByInvitationId EndpointPattern = EndpointPattern{
	Pattern: "/user/repository_invitations/{invitation_id}",
	Method:  "PATCH",
}

var DeleteEnterpriseUserRepositoryInvitationsByInvitationId EndpointPattern = EndpointPattern{
	Pattern: "/user/repository_invitations/{invitation_id}",
	Method:  "DELETE",
}

var GetEnterpriseUserSocialAccounts EndpointPattern = EndpointPattern{
	Pattern: "/user/social_accounts",
	Method:  "GET",
}

var PostEnterpriseUserSocialAccounts EndpointPattern = EndpointPattern{
	Pattern: "/user/social_accounts",
	Method:  "POST",
}

var DeleteEnterpriseUserSocialAccounts EndpointPattern = EndpointPattern{
	Pattern: "/user/social_accounts",
	Method:  "DELETE",
}

var GetEnterpriseUserSshSigningKeys EndpointPattern = EndpointPattern{
	Pattern: "/user/ssh_signing_keys",
	Method:  "GET",
}

var PostEnterpriseUserSshSigningKeys EndpointPattern = EndpointPattern{
	Pattern: "/user/ssh_signing_keys",
	Method:  "POST",
}

var GetEnterpriseUserSshSigningKeysBySshSigningKeyId EndpointPattern = EndpointPattern{
	Pattern: "/user/ssh_signing_keys/{ssh_signing_key_id}",
	Method:  "GET",
}

var DeleteEnterpriseUserSshSigningKeysBySshSigningKeyId EndpointPattern = EndpointPattern{
	Pattern: "/user/ssh_signing_keys/{ssh_signing_key_id}",
	Method:  "DELETE",
}

var GetEnterpriseUserStarred EndpointPattern = EndpointPattern{
	Pattern: "/user/starred",
	Method:  "GET",
}

var GetEnterpriseUserStarredByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/user/starred/{owner}/{repo}",
	Method:  "GET",
}

var PutEnterpriseUserStarredByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/user/starred/{owner}/{repo}",
	Method:  "PUT",
}

var DeleteEnterpriseUserStarredByOwnerByRepo EndpointPattern = EndpointPattern{
	Pattern: "/user/starred/{owner}/{repo}",
	Method:  "DELETE",
}

var GetEnterpriseUserSubscriptions EndpointPattern = EndpointPattern{
	Pattern: "/user/subscriptions",
	Method:  "GET",
}

var GetEnterpriseUserTeams EndpointPattern = EndpointPattern{
	Pattern: "/user/teams",
	Method:  "GET",
}

var GetEnterpriseUserByAccountId EndpointPattern = EndpointPattern{
	Pattern: "/user/{account_id}",
	Method:  "GET",
}

var GetEnterpriseUsers EndpointPattern = EndpointPattern{
	Pattern: "/users",
	Method:  "GET",
}

var GetEnterpriseUsersByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}",
	Method:  "GET",
}

var GetEnterpriseUsersAttestationsByUsernameBySubjectDigest EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/attestations/{subject_digest}",
	Method:  "GET",
}

var GetEnterpriseUsersDockerConflictsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/docker/conflicts",
	Method:  "GET",
}

var GetEnterpriseUsersEventsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/events",
	Method:  "GET",
}

var GetEnterpriseUsersEventsOrgsByUsernameByOrg EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/events/orgs/{org}",
	Method:  "GET",
}

var GetEnterpriseUsersEventsPublicByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/events/public",
	Method:  "GET",
}

var GetEnterpriseUsersFollowersByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/followers",
	Method:  "GET",
}

var GetEnterpriseUsersFollowingByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/following",
	Method:  "GET",
}

var GetEnterpriseUsersFollowingByUsernameByTargetUser EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/following/{target_user}",
	Method:  "GET",
}

var GetEnterpriseUsersGistsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/gists",
	Method:  "GET",
}

var GetEnterpriseUsersGpgKeysByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/gpg_keys",
	Method:  "GET",
}

var GetEnterpriseUsersHovercardByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/hovercard",
	Method:  "GET",
}

var GetEnterpriseUsersInstallationByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/installation",
	Method:  "GET",
}

var GetEnterpriseUsersKeysByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/keys",
	Method:  "GET",
}

var GetEnterpriseUsersOrgsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/orgs",
	Method:  "GET",
}

var GetEnterpriseUsersPackagesByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages",
	Method:  "GET",
}

var GetEnterpriseUsersPackagesByUsernameByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}",
	Method:  "GET",
}

var DeleteEnterpriseUsersPackagesByUsernameByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}",
	Method:  "DELETE",
}

var PostEnterpriseUsersPackagesRestoreByUsernameByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}/restore",
	Method:  "POST",
}

var GetEnterpriseUsersPackagesVersionsByUsernameByPackageTypeByPackageName EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}/versions",
	Method:  "GET",
}

var GetEnterpriseUsersPackagesVersionsByUsernameByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}",
	Method:  "GET",
}

var DeleteEnterpriseUsersPackagesVersionsByUsernameByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}",
	Method:  "DELETE",
}

var PostEnterpriseUsersPackagesVersionsRestoreByUsernameByPackageTypeByPackageNameByPackageVersionId EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore",
	Method:  "POST",
}

var GetEnterpriseUsersProjectsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/projects",
	Method:  "GET",
}

var GetEnterpriseUsersReceivedEventsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/received_events",
	Method:  "GET",
}

var GetEnterpriseUsersReceivedEventsPublicByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/received_events/public",
	Method:  "GET",
}

var GetEnterpriseUsersReposByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/repos",
	Method:  "GET",
}

var GetEnterpriseUsersSettingsBillingActionsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/settings/billing/actions",
	Method:  "GET",
}

var GetEnterpriseUsersSettingsBillingPackagesByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/settings/billing/packages",
	Method:  "GET",
}

var GetEnterpriseUsersSettingsBillingSharedStorageByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/settings/billing/shared-storage",
	Method:  "GET",
}

var GetEnterpriseUsersSocialAccountsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/social_accounts",
	Method:  "GET",
}

var GetEnterpriseUsersSshSigningKeysByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/ssh_signing_keys",
	Method:  "GET",
}

var GetEnterpriseUsersStarredByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/starred",
	Method:  "GET",
}

var GetEnterpriseUsersSubscriptionsByUsername EndpointPattern = EndpointPattern{
	Pattern: "/users/{username}/subscriptions",
	Method:  "GET",
}

var GetEnterpriseVersions EndpointPattern = EndpointPattern{
	Pattern: "/versions",
	Method:  "GET",
}

var GetEnterpriseZen EndpointPattern = EndpointPattern{
	Pattern: "/zen",
	Method:  "GET",
}
