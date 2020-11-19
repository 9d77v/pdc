import gql from 'graphql-tag';


const ADD_DEVICE_MODEL = gql`
mutation createDeviceModel($input:NewDeviceModel!){
   createDeviceModel(input:$input){
     id
   }
}
`

const UPDATE_DEVICE_MODEL = gql`
mutation updateDeviceModel($input:NewUpdateDeviceModel!){
   updateDeviceModel(input:$input){
     id
   }
}
`

const ADD_ATTRIBUTE_MODEL = gql`
mutation createAttributeModel($input:NewAttributeModel!){
   createAttributeModel(input:$input){
     id
   }
}
`

const UPDATE_ATTRIBUTE_MODEL = gql`
mutation updateAttributeModel($input:NewUpdateAttributeModel!){
   updateAttributeModel(input:$input){
     id
   }
}
`

const DELETE_ATTRIBUTE_MODEL = gql`
mutation deleteAttributeModel($id:Int!){
   deleteAttributeModel(id:$id){
     id
   }
}
`

const ADD_TELEMETRY_MODEL = gql`
mutation createTelemetryModel($input:NewTelemetryModel!){
   createTelemetryModel(input:$input){
     id
   }
}
`

const UPDATE_TELEMETRY_MODEL = gql`
mutation updateTelemetryModel($input:NewUpdateTelemetryModel!){
   updateTelemetryModel(input:$input){
     id
   }
}
`

const DELETE_TELEMETRY_MODEL = gql`
mutation deleteTelemetryModel($id:Int!){
   deleteTelemetryModel(id:$id){
     id
   }
}
`

const LIST_DEVICE_MODEL = gql`
query deviceModels($keyword: String, $page: Int, $pageSize: Int, $sorts: [Sort!]) {
  deviceModels(keyword: $keyword, page: $page, pageSize: $pageSize, sorts: $sorts) {
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
query deviceModels($keyword: String, $page: Int, $pageSize: Int) {
  deviceModels(keyword: $keyword, page: $page, pageSize: $pageSize) {
       edges{
          value:id 
          text:name 
       }
   }
  }
`;


const GET_DEVICE_MODEL = gql`
query deviceModels($ids:[ID!]) {
  deviceModels(ids:$ids) {
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


const ADD_DEVICE = gql`
mutation createDevice($input:NewDevice!){
   createDevice(input:$input){
     id
   }
}
`

const UPDATE_DEVICE = gql`
mutation updateDevice($input:NewUpdateDevice!){
   updateDevice(input:$input){
     id
   }
}
`


const LIST_DEVICE = gql`
query devices($keyword: String, $page: Int, $pageSize: Int, $sorts: [Sort!]) {
  devices(keyword: $keyword, page: $page, pageSize: $pageSize, sorts: $sorts) {
    totalCount
    edges {
      id
      name
      deviceModelID
      deviceModelName
      ip
      port
      username
      password
    }
  }
}
`

const LIST_DEVICE_SELECTOR = gql`
query devices($deviceType:Int!, $sorts: [Sort!]) {
  devices(deviceType:$deviceType,sorts: $sorts) {
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
query devices($ids:[ID!]) {
  devices(ids:$ids) {
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
      deviceModelName
      deviceModelDesc
      deviceModelDeviceType
      deviceModelCameraCompany
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

const ADD_DEVICE_DASHBOARD = gql`
mutation createDeviceDashboard($input:NewDeviceDashboard!){
   createDeviceDashboard(input:$input){
     id
   }
}
`

const UPDATE_DEVICE_DASHBOARD = gql`
mutation updateDeviceDashboard($input:NewUpdateDeviceDashboard!){
   updateDeviceDashboard(input:$input){
     id
   }
}
`

const DELETE_DEVICE_DASHBOARD = gql`
mutation deleteDeviceDashboard($id:Int!){
   deleteDeviceDashboard(id:$id){
     id
   }
}
`
const LIST_DEVICE_DASHBOARD = gql`
query deviceDashboards($keyword: String, $page: Int, $pageSize: Int, $sorts: [Sort!]) {
  deviceDashboards(keyword: $keyword, page: $page, pageSize: $pageSize, sorts: $sorts) {
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

const ADD_DEVICE_DASHBOARD_TELEMETRY = gql`
mutation addDeviceDashboardTelemetry($input:NewDeviceDashboardTelemetry!){
  addDeviceDashboardTelemetry(input:$input){
     id
   }
}
`

const REMOVE_DEVICE_DASHBOARD_TELEMETRY = gql`
mutation removeDeviceDashboardTelemetry($ids:[Int!]!){
   removeDeviceDashboardTelemetry(ids:$ids){
     id
   }
}
`

const ADD_DEVICE_DASHBOARD_CAMERA = gql`
mutation addDeviceDashboardCamera($input:NewDeviceDashboardCamera!){
  addDeviceDashboardCamera(input:$input){
     id
   }
}
`

const REMOVE_DEVICE_DASHBOARD_CAMERA = gql`
mutation removeDeviceDashboardCamera($ids:[Int!]!){
   removeDeviceDashboardCamera(ids:$ids){
     id
   }
}
`

const CAMERA_CAPTURE = gql`
mutation cameraCapture($deviceID:Int!){
   cameraCapture(deviceID:$deviceID) 
}
`

export {
  ADD_DEVICE_MODEL, UPDATE_DEVICE_MODEL,
  DEVICE_MODEL_COMBO,
  LIST_DEVICE_MODEL, GET_DEVICE_MODEL,
  ADD_ATTRIBUTE_MODEL, UPDATE_ATTRIBUTE_MODEL, DELETE_ATTRIBUTE_MODEL,
  ADD_TELEMETRY_MODEL, UPDATE_TELEMETRY_MODEL, DELETE_TELEMETRY_MODEL,
  ADD_DEVICE, UPDATE_DEVICE,
  LIST_DEVICE, LIST_DEVICE_SELECTOR, GET_DEVICE, GET_MOBILE_HOME_DEVICES,
  ADD_DEVICE_DASHBOARD, UPDATE_DEVICE_DASHBOARD, DELETE_DEVICE_DASHBOARD,
  LIST_DEVICE_DASHBOARD,
  ADD_DEVICE_DASHBOARD_TELEMETRY, REMOVE_DEVICE_DASHBOARD_TELEMETRY,
  ADD_DEVICE_DASHBOARD_CAMERA, REMOVE_DEVICE_DASHBOARD_CAMERA,
  CAMERA_CAPTURE
}
