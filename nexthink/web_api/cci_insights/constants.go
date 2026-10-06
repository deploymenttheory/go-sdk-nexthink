package cci_insights

const Endpoint = "/apigateway/cci/graphql"

const operationID = "graphql.benchmark"

const queryGetBinaryInsights = `
	query getBinaryInsights($filters: BinaryInsightsInput!) {
		binaryInsights(filters: $filters) {
			resultCode
			binaryInsights {
				attributes {
					dmUri
					value
				}
				tenant_device_count
				issue_type
				insight_type
				root_metric_value
				group_metric_value
				tenant_metric_value
				rec_binary_version
				rec_upgrade
				rec_tenant_metric_improvement
			}
		}
	}
`

const queryGetDiagnosticBinaryInsights = `
  query getBinaryInsights($filters: BinaryInsightsInput!) {
    binaryInsights(filters: $filters) {
      resultCode
      binaryInsights {
        attributes {
          dmUri
          value
        }
        parameter_set_id
        group_id
        region
        tenant_device_count
        group_timeframe_start
        group_timeframe_duration_days
        tenant_timeframe_start
        tenant_timeframe_duration_days
        issue_type
        insight_type
        root_metric_value
        group_metric_value
        tenant_metric_value
        rec_binary_version
        rec_upgrade
        rec_tenant_metric_improvement
      }
    }
  }
`

const queryGetDataset = `
  query getDataset($name: String!) {
    dataset(name: $name) {
      resultCode
      schema {
        name
        type
      }
      rows
    }
  }
`
