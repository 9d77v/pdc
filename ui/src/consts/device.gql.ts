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

const LIST_DEVICE_MODEL = gql`
query deviceModels($keyword: String, $page: Int, $pageSize: Int) {
  deviceModels(keyword: $keyword, page: $page, pageSize: $pageSize) {
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
query devices($keyword: String, $page: Int, $pageSize: Int) {
  devices(keyword: $keyword, page: $page, pageSize: $pageSize) {
    totalCount
    edges {
      id
      name
      deviceModelID
      deviceModelName
      ip
      port
    }
  }
}
`;

const GET_DEVICE = gql`
query devices($ids:[ID!]) {
  devices(ids:$ids) {
    edges {
      id
      name
      ip
      port
      deviceModelID
      deviceModelName
      deviceModelDesc
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
        createdAt
        updatedAt
      }
      createdAt
      updatedAt
    }
  }
}
`;
export {
  ADD_DEVICE_MODEL, UPDATE_DEVICE_MODEL,
  DEVICE_MODEL_COMBO,
  LIST_DEVICE_MODEL, GET_DEVICE_MODEL,
  ADD_ATTRIBUTE_MODEL, UPDATE_ATTRIBUTE_MODEL,
  ADD_TELEMETRY_MODEL, UPDATE_TELEMETRY_MODEL,
  ADD_DEVICE, UPDATE_DEVICE,
  LIST_DEVICE, GET_DEVICE
}
