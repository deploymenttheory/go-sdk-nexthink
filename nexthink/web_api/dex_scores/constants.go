package dex_scores

const Endpoint = "/apigateway/dex-ec/graphql"

const operationID = "graphql.dex_configuration"

const queryGetCampaign = `
  query getCampaign($filter: Filter) {
    campaign(filter: $filter) {
      uuid
      nqlId
      numberOfRespondents
      complaints {
        name
        votes
      }
      complaintsNqlId
    }
  }
`

const queryGetDeviceExperience = `
  query getDeviceExperience($metricId: String!, $filter: Filter) {
    deviceExperience(metricId: $metricId, filter: $filter) {
      devicesWithIssues
    }
  }
`

const queryGetDimensionBreakdowns = `
  query getDimensionBreakdowns($lang: String!) {
    dimensionBreakdowns(lang: $lang) {
      key
      nqlId
      dmUri
      value
    }
  }
`

const queryGetDimensionsV2 = `
  query getDimensionsV2(
    $dmUri: String!
    $filter: Filter
    $pagination: Pagination!
    $orderInput: OrderInput!
  ) {
    dimensionsV2(
      dmUri: $dmUri
      filter: $filter
      pagination: $pagination
      orderInput: $orderInput
    ) {
      totalSize
      dimensions {
        name
        technologyScore
        sentimentScore
      }
    }
  }
`

const queryGetInvestigationUrl = `
  query getInvestigationUrl(
    $investigationKey: InvestigationKey!
    $filter: Filter
    $metricId: String
  ) {
    investigationUrl(
      investigationKey: $investigationKey
      filter: $filter
      metricId: $metricId
    )
  }
`

const queryGetLeaves = `
  query getLeaves($metricId: String!, $filter: Filter) {
    leaves(metricId: $metricId, filter: $filter) {
      id
      name
      score
      improvement
      timeLost
      thresholds {
        experience
        numberOfDevices
      }
      uuid
      appType
      troubleshootUrl
      devicesInvestigationUrl
      usersInvestigationUrl
    }
  }
`

const queryGetMetricThreshold = `
  query getMetricThreshold($metricId: String!) {
    metricThreshold(metricId: $metricId) {
      isInverted
      unit
      average
      frustrating
    }
  }
`

const queryGetScores = `
  query getScores($scoreLevel: ScoreLevel!, $filter: Filter) {
    scores(scoreLevel: $scoreLevel, filter: $filter) {
      id
      name
      score
      nodeType
      improvement
      timeLost
      uuid
      thresholds {
        experience
        numberOfDevices
      }
      level
      appType
      troubleshootUrl
      devicesInvestigationUrl
      usersInvestigationUrl
    }
  }
`

const queryGetTrend = `
  query getTrend($filter: Filter) {
    trend(filter: $filter) {
      dexTrend {
        ...ScoreTrendNodeCommon
      }
      technologyTrend {
        ...ScoreTrendNodeCommon
      }
      sentimentTrend {
        ...ScoreTrendNodeCommon
      }
    }
  }


  fragment ScoreTrendNodeCommon on TrendNodeScore {
    startDate
    endDate
    value
    userCount
  }
`

const queryGetTrendDevices = `
  query getTrendDevices($metricId: String!, $filter: Filter) {
    trendDevices(metricId: $metricId, filter: $filter) {
      startDate
      endDate
      value
    }
  }
`

const queryGetTrendEmployeesWithIssues = `
  query getTrendEmployeesWithIssues($metricId: String!, $filter: Filter) {
    trendEmployeesWithIssues(metricId: $metricId, filter: $filter) {
      startDate
      endDate
      value
      userCount
    }
  }
`

const queryGetTrendImprovement = `
  query getTrendImprovement($metricId: String!, $filter: Filter) {
    trendImprovement(metricId: $metricId, filter: $filter) {
      startDate
      endDate
      value
      userCount
    }
  }
`

const queryGetTrendScore = `
  query getTrendScore($metricId: String!, $filter: Filter) {
    trendScore(metricId: $metricId, filter: $filter) {
      startDate
      endDate
      value
      userCount
    }
  }
`

const queryGetTrendTimeLost = `
  query getTrendTimeLost($metricId: String!, $filter: Filter) {
    trendTimeLost(metricId: $metricId, filter: $filter) {
      startDate
      endDate
      value
      userCount
    }
  }
`

const queryGetTrendWithRange = `
  query getTrendWithRange($filter: Filter, $startTsSec: Long, $endTsSec: Long) {
    trend(filter: $filter, startTsSec: $startTsSec, endTsSec: $endTsSec) {
      dexTrend {
        ...ScoreTrendNodeCommon
      }
      technologyTrend {
        ...ScoreTrendNodeCommon
      }
      sentimentTrend {
        ...ScoreTrendNodeCommon
      }
    }
  }


  fragment ScoreTrendNodeCommon on TrendNodeScore {
    startDate
    endDate
    value
    userCount
  }
`

const queryGetWhatsChanged = `
  query getWhatsChanged($startTsSec: Long!, $endTsSec: Long!, $filter: Filter) {
    whatsChanged(
      startTsSec: $startTsSec
      endTsSec: $endTsSec
      filter: $filter
    ) {
      fromRootScore {
        date
        score
      }
      toRootScore {
        date
        score
      }
      keyVariations {
        parentId
        parentName
        parentUuid
        metricLevel
        id
        name
        impact
        parentAppType
        diagnoseUrl
        devicesInvestigationUrl
        usersInvestigationUrl
        impactInsights
      }
    }
  }
`
