package network_insights

const Endpoint = "/apigateway/dash/graphql"

const operationID = "graphql.dashboards"

const queryGetInsights = `
  query GetInsights($input: GraphInsightsInput!) {
    graphInsights(input: $input) {
      nodes {
        identifier {
          columnName
          name
        }
        items
        edges {
          targetNode {
            columnName
            name
          }
          items
        }
      }
    }
  }
`
