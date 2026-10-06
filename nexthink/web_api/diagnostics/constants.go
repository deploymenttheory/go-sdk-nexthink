package diagnostics

const Endpoint = "/apigateway/diagnostic/graphql"

const operationID = "graphql.diagnostics"

const queryBinaryInfo = `
  query BinaryInfo($details: IssueDetails!) {
    binaryInfo(details: $details) {
      productName
      company
      platforms
      hasUI
    }
  }
`

const queryGetDiagnosticContexts = `
  query GetDiagnosticContexts {
    diagnosticContextDefinitions {
      id
      scoreIds
      params {
        key
        type
        required
        defaultValue
        validationRules {
          maxItems
          min
          max
          allowedValues
        }
      }
      timeFrameConfig {
        defaultDuration
        maxDurationDays
        maxDurationDays15m
        maxRetentionDays
        maxRetentionDays15m
      }
      layout
    }
    diagnosticContextLayouts {
      name
      widgets {
        name
        labels {
          key
          value
        }
        parameters {
          key
          value
        }
      }
    }
  }
`

const queryGetDiagnosticOverview = `
  query GetDiagnosticOverview($id: String!) {
    getDiagnosticOverview(id: $id) {
      query
      includeBinaryRecommendations
      hasTotalDevices
    }
  }
`

const queryGetImpactedAndTotalObjects = `
  query GetImpactedAndTotalObjects($details: IssueDetails!) {
    impactedAssociatedObjects(details: $details) {
      drilldownQuery
      count
    }
    totalAssociatedObjects: totalAssociatedObjects(details: $details)
  }
`

const queryGetStandaloneDashboard = `
  query GetStandaloneDashboard($issueReference: IssueReference!) {
    getStandaloneDashboard(issueReference: $issueReference) {
      widgets
    }
  }
`

const queryHierarchyBreakdownDimensions = `
  query HierarchyBreakdownDimensions {
    hierarchyBreakdownDimensions {
      hierarchyId
      hierarchyName
      levelDepth
      levelName
      nodeValues
    }
  }
`

const queryIssueTimeseries = `
  query IssueTimeseries($details: IssueDetails!) {
    issueEventsTimeSeries(details: $details) {
      scope
      items {
        label
        value
      }
    }
  }
`

const queryTroubleshootingInsights = `
  query TroubleshootingInsights($details: IssueDetails!) {
    troubleshootingInsights(details: $details) {
      causes {
        id
        name
        subContextValue
        confidence
        confidenceLevel
        breakdown {
          index {
            name
            values
          }
          intCols {
            name
            values
          }
          floatCols {
            name
            values
          }
        }
      }
    }
  }
`

const queryGetConfiguration = `
  query configuration {
    configuration {
      staticNow
      tenantTimezone
      isMtpOnly
      binaryRecommendationMaxRequests
    }
  }
`

const queryIssueEventsAndAssociatedObjectsByHierarchyBreakdown = `
  query issueEventsAndAssociatedObjectsByHierarchyBreakdown(
    $details: IssueDetails!
    $hierarchyLevel: HierarchyLevelInput!
  ) {
    issueEventsAndAssociatedObjectsByHierarchyBreakdown(
      details: $details
      hierarchyLevel: $hierarchyLevel
    ) {
      label
      items {
        label
        value
      }
    }
  }
`

const queryIssueEventsAndAssociatedObjectsByLocationBreakdown = `
  query issueEventsAndAssociatedObjectsByLocationBreakdown(
    $details: IssueDetails!
    $dimension: LocationBreakdownDimension!
  ) {
    issueEventsAndAssociatedObjectsByLocationBreakdown(
      details: $details
      dimension: $dimension
    ) {
      label
      items {
        label
        value
      }
    }
  }
`

const queryIssueEventsAndAssociatedObjectsByTechnicalBreakdown = `
  query issueEventsAndAssociatedObjectsByTechnicalBreakdown(
    $details: IssueDetails!
    $dimension: TechnicalBreakdownDimension!
  ) {
    issueEventsAndAssociatedObjectsByTechnicalBreakdown(
      details: $details
      dimension: $dimension
    ) {
      label
      items {
        label
        value
      }
    }
  }
`

const queryGetDashboard = `
  query GetDashboard($issueReference: IssueReference!) {
    getDashboard(issueReference: $issueReference) {
      widgets
    }
  }
`
