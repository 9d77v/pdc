import gql from 'graphql-tag';

const LIST_DEVICE_MODEL = gql`
query deviceModels($searchParam:SearchParam!) {
  deviceModels(searchParam: $searchParam) {
    totalCount
    edges {
      id
      name
      desc
      deviceType
    }
  }
}
`;

const DEVICE_MODEL_COMBO = gql`
query deviceModels($searchParam:SearchParam!) {
  deviceModels(searchParam: $searchParam) {
       edges{
          value:id 
          text:name 
       }
   }
  }
`;

const GET_DEVICE_MODEL = gql`
query deviceModels($searchParam:SearchParam!) {
  deviceModels(searchParam: $searchParam) {
    edges {
      id
      name
      desc
      deviceType
      cameraCompany
      attributeModels{
        id
        key
        name
        createdAt
        updatedAt
      }
      telemetryModels{
        id
        key
        name
        factor
        unit
        unitName
        scale
        createdAt
        updatedAt
      }
      createdAt
      updatedAt
    }
  }
}
`;

const LIST_DEVICE = gql`
query devices($searchParam:SearchParam!) {
  devices(searchParam: $searchParam) {
    totalCount
    edges {
      id
      name
      deviceModelID
      deviceModel{
        name
      }
      ip
      port
      username
      password
    }
  }
}
`

const LIST_DEVICE_SELECTOR = gql`
query devices($deviceType:Int!, $searchParam:SearchParam!) {
  devices(deviceType:$deviceType,searchParam: $searchParam) {
    edges {
      id
      name
      telemetries{
        id
        key
        name
      }
    }
  }
}
`

const GET_DEVICE = gql`
query devices( $searchParam:SearchParam!) {
  devices(searchParam: $searchParam) {
    edges {
      id
      name
      ip
      port
      accessKey
      secretKey
      username
      password
      deviceModelID
      deviceModel{
        id
        name
        desc
        deviceType
        cameraCompany
      }
      attributes{
        id
        key
        name
        value
        createdAt
        updatedAt
      }
      telemetries{
        id
        key
        name
        value
        unit
        unitName
        factor
        scale
        createdAt
        updatedAt
      }
      createdAt
      updatedAt
    }
  }
}
`

const GET_MOBILE_HOME_DEVICES = gql`
query appDeviceDashboards($deviceType: Int) {
  appDeviceDashboards(deviceType: $deviceType) {
    totalCount
    edges {
      id
      name
      isVisible
      deviceType
      telemetries{
        id
        deviceID
        deviceName
        telemetryID
        name
        value
        factor
        scale
        unit
      }
      cameras{
        id
        deviceID
        deviceName
      }
    }
  }
}
`

const LIST_DEVICE_DASHBOARD = gql`
query deviceDashboards($searchParam:SearchParam!) {
  deviceDashboards(searchParam: $searchParam) {
    totalCount
    edges {
      id
      name
      isVisible
      deviceType
      telemetries{
        id
        deviceID
        deviceName
        telemetryID
        name
        value
        factor
        scale
        unit
      }
      cameras{
        id
        deviceID
        deviceName
      }
    }
  }
}
`

const GET_CAMERA_TIME_LAPSE_VIDEOS = gql`
query cameraTimeLapseVideos($deviceID: Int!) {
  cameraTimeLapseVideos(deviceID: $deviceID) {
    totalCount
    edges {
      id
      deviceID
      date
      videoURL
    }
  }
}
`

export {
  DEVICE_MODEL_COMBO, LIST_DEVICE_MODEL, GET_DEVICE_MODEL,
  LIST_DEVICE, LIST_DEVICE_SELECTOR, GET_DEVICE, GET_MOBILE_HOME_DEVICES,
  LIST_DEVICE_DASHBOARD, GET_CAMERA_TIME_LAPSE_VIDEOS
}
