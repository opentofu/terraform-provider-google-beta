// this file is copied from mmv1, any changes made here will be overwritten

package tests

import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot.AgentCleanPolicy.ON_BRANCH_CHANGE
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot.AgentCleanFilesPolicy.ALL_UNTRACKED
import jetbrains.buildServer.configs.kotlin.AbsoluteId
import ClientConfiguration

fun TestConfiguration() : ClientConfiguration {
    return ClientConfiguration("custId", "org", "org2", "billingAccount", "billingAccount2", "masterBillingAccount", "credentials", "project", "orgDomain", "projectNumber", "region", "serviceAccount", "zone", "firestoreProject", "identityUser")
}

fun TestVcsRootId() : AbsoluteId {
    return AbsoluteId("TerraformProviderFoobar")
}