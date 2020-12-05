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
  ADD_ATTRIBUTE_MODEL, UPDATE_ATTRIBUTE_MODEL, DELETE_ATTRIBUTE_MODEL,
  ADD_TELEMETRY_MODEL, UPDATE_TELEMETRY_MODEL, DELETE_TELEMETRY_MODEL,
  ADD_DEVICE, UPDATE_DEVICE,
  ADD_DEVICE_DASHBOARD, UPDATE_DEVICE_DASHBOARD, DELETE_DEVICE_DASHBOARD,
  ADD_DEVICE_DASHBOARD_TELEMETRY, REMOVE_DEVICE_DASHBOARD_TELEMETRY,
  ADD_DEVICE_DASHBOARD_CAMERA, REMOVE_DEVICE_DASHBOARD_CAMERA,
  CAMERA_CAPTURE
}
