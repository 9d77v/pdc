import gql from 'graphql-tag';


const ADD_DEVICE_MODEL = gql`
mutation createDeviceModel($input:NewDeviceModel!){
   createDeviceModel(input:$input){
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
      deviceType
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
export {
  ADD_DEVICE_MODEL, LIST_DEVICE_MODEL, GET_DEVICE_MODEL,
  ADD_ATTRIBUTE_MODEL, UPDATE_ATTRIBUTE_MODEL,
  ADD_TELEMETRY_MODEL, UPDATE_TELEMETRY_MODEL
}
