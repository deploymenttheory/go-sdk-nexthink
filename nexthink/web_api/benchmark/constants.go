package benchmark

const Endpoint = "/apigateway/benchmark/graphql"

const operationID = "graphql.benchmarking"

const queryGetBinaryProductMaps = `
  query getBinaryProductMaps($params: ProfileType!) {
    binaryProductMaps(params: $params) {
      resultCode
      binaryToProductsMap {
        key
        values
      }
      productToBinariesMap {
        key
        values
      }
    }
  }
`

const queryGetProfile = `
  query getProfile($params: ProfileParams!) {
    profile(params: $params) {
      resultCode
      tenants
      devices
      metrics {
        name
        mean
        p25
        p95
      }
    }
  }
`

const queryGetProfileSearchItems = `
  query getProfileSearchItems($params: ProfileType!) {
    profileSearchItems(params: $params) {
      resultCode
      items {
        name
        extra
      }
    }
  }
`

const queryGetProfileVersions = `
  query getProfileVersions($params: ProfileVersionsParams!) {
    profileVersions(params: $params) {
      resultCode
      versions {
        version
        versionInfo {
          resultCode
          tenants
          devices
          devicesRatio
          metrics {
            name
            mean
          }
        }
      }
    }
  }
`

const queryLookupBenchmark = `
  query lookupBenchmark($params: BenchmarkParams!) {
    lookup(params: $params) {
      resultCode
      lookupMatches {
        key {
          key
          value
        }
        value
        weight
        origins
        precedence
      }
    }
  }
`

const queryProductProperties = `
  query productProperties($params: ProductPropertiesParams!) {
    productProperties(params: $params) {
      resultCode
      productName
      properties {
        name
        value
      }
    }
  }
`

const queryProfileProperties = `
  query profileProperties($params: ProfilePropertiesParams!) {
    profileProperties(params: $params) {
      resultCode
      profileKey {
        binaryName
        productName
      }
      properties {
        name
        value
      }
    }
  }
`
