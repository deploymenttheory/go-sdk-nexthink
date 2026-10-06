package alert_hub

const Endpoint = "/apigateway/mnt/alert/hub/graphql"

const operationID = "graphql.alert_hub"

const queryAlertImpactAssessment = `
	query AlertImpactAssessment(
		$alertUuid: String!
		$timeFrame: IssueTimeFrame!
	) {
		alertImpactAssessment(alertUuid: $alertUuid, timeFrame: $timeFrame) {
			impactLevel
			assessment
		}
	}
`

const queryEventsDrillDown = `
	query EventsDrillDown(
		$monitorName: String!
		$startedAt: Int!
		$context: [ComplexDataFieldInput!]!
	) {
		eventsDrillDown(
			monitorName: $monitorName
			startedAt: $startedAt
			context: $context
		)
	}
`

const queryIssues = `
	query Issues(
		$filterInput: [FilterInput!]
		$orderBy: AlertOrderByInput
		$timeFrame: IssueTimeFrame
	) {
		issues(
			filterInput: $filterInput
			orderBy: $orderBy
			timeFrame: $timeFrame
		) {
			count
			openCount
			criticalCount
			list {
				lastAlertUuid
				configUuid
				nqlId
				monitorUuid
				priority
				status
				name
				occurrence
				lastTriggerDateTime
				resolvedAt
				tags
				origin
				associations {
					count
					type
				}
				impactType
				context {
					key
					value
					dataType
					dataPath
					visible
				}
				customActions
				contextHash
				metricType
				detectionType
				benchmarkType
				benchmarkMetricUri
				mainMetricUri
			}
		}
	}
`

const queryIssuesSelectedPeriod = `
	query IssuesSelectedPeriod(
		$filterInput: [FilterInput!]
		$orderBy: AlertOrderByInput
		$timeFrame: IssueTimeFrame
	) {
		issuesSelectedPeriod(
			filterInput: $filterInput
			orderBy: $orderBy
			timeFrame: $timeFrame
		) {
			count
			openCount
			criticalCount
			list {
				lastAlertUuid
				configUuid
				nqlId
				monitorUuid
				priority
				status
				name
				occurrence
				lastTriggerDateTime
				resolvedAt
				tags
				origin
				associations {
					count
					type
				}
				impactType
				context {
					key
					value
					dataType
					dataPath
					visible
				}
				customActions
				contextHash
				metricType
				detectionType
				benchmarkType
				benchmarkMetricUri
				mainMetricUri
			}
		}
	}
`

const queryIssuesTimeline = `
	query IssuesTimeline(
		$filterInput: [FilterInput!]
		$timeFrame: IssueTimeFrame
	) {
		issuesTimeline(filterInput: $filterInput, timeFrame: $timeFrame) {
			issuesDataPoints {
				fromDate
				toDate
				values {
					openLow
					openMedium
					openHigh
					openCritical
					openTotal
				}
			}
			scope
		}
	}
`

const queryMonitorConfigView = `
	query MonitorConfigView($monitorConfigViewInput: MonitorConfigViewInput!) {
		monitorConfigView(monitorConfigViewInput: $monitorConfigViewInput)
	}
`

const queryAlertTriggerInfo = `
	query alertTriggerInfo($alertUuid: String!) {
		alertTriggerInfo(alertUuid: $alertUuid) {
			monitorMetricType
			issueDescription
			metrics {
				label
				value
				valueDataType
				threshold
				thresholdDataType
			}
		}
	}
`

const queryAlertsImpactedAssociationOverTime = `
	query alertsImpactedAssociationOverTime(
		$nqlId: String!
		$contextHash: String!
		$timeFrame: IssueTimeFrame!
	) {
		alertsImpactedAssociationOverTime(
			nqlId: $nqlId
			contextHash: $contextHash
			timeFrame: $timeFrame
		) {
			dataPoints {
				dateTime
				open
				devices
				users
			}
		}
	}
`

const queryGetAlertOnChangeTimeSeries = `
  query getAlertOnChangeTimeSeries(
    $monitorUuid: String!
    $contextHash: String!
    $alertContext: String
    $timeFrame: IssueTimeFrameForWidget!
    $monitorType: String
  ) {
    timeSeriesData(
      monitorUuid: $monitorUuid
      context: $alertContext
      timeFrame: $timeFrame
      monitorType: $monitorType
    ) {
      metric {
        dataPoints {
          time
          value
        }
      }
      baseline {
        dataPoints {
          time
          value
        }
      }
      threshold {
        dataPoints {
          time
          value
        }
      }
    }
    timeSeriesIssuePeriods(
      monitorUuid: $monitorUuid
      contextHash: $contextHash
      timeFrame: $timeFrame
    ) {
      periods {
        fromDate
        toDate
      }
      count
    }
  }
`

const queryGetIssuePeriods = `
	query getIssuePeriods(
		$monitorUuid: String!
		$contextHash: String!
		$timeFrame: IssueTimeFrame!
	) {
		issuePeriods(
			monitorUuid: $monitorUuid
			contextHash: $contextHash
			timeFrame: $timeFrame
		) {
			periods {
				fromDate
				toDate
			}
			count
		}
	}
`

const queryGetStaticAlertTimeSeries = `
  query getStaticAlertTimeSeries(
    $monitorUuid: String!
    $contextHash: String!
    $timeFrame: IssueTimeFrameForWidget!
  ) {
    timeSeriesIssuePeriods(
      monitorUuid: $monitorUuid
      contextHash: $contextHash
      timeFrame: $timeFrame
    ) {
      periods {
        fromDate
        toDate
      }
      count
    }
  }
`

const queryTags = `
	query Tags {
		tags {
			name
			color
		}
	}
`
