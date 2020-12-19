
export interface IDeviceModel {
    id: number
    name: string
    desc: string
    deviceType: number
    cameraCompany: number
    createdAt: number
    updatedAt: number
}

export interface IDevice {
    id: number
    name: string
    ip: string
    port: number
    accessKey: string
    secretKey: string
    deviceModelID: number
    deviceModel: IDeviceModel
    health?: number
    username: string
    password: string
    createdAt: number
    updatedAt: number
}

export interface IDeviceDashboard {
    id: number
    name: string
    isVisible: boolean
    telemetries: IDeviceDashboardTelemetry[]
    cameras: IDeviceDashboardCamera[]
    deviceType: number
    createdAt: number
    updatedAt: number
}

export interface IDeviceDashboardTelemetry {
    id: number
    deviceID: number
    deviceName: string
    name: string
    value?: number
    factor: number
    unit: string
    scale: number
    telemetryID: number
}

export interface IDeviceDashboardCamera {
    id: number
    deviceID: number
    deviceName: string
}

export interface INewDeviceModel {
    name: string
    deviceType: number
    desc: string
    cameraCompany: number
}

export interface IUpdateDeviceModel {
    id: number
    name: string
    desc: string
}

export interface INewDevice {
    name: string
    deviceModelID: number
    ip: string
    port: number
    username: string
    password: string
}

export interface IUpdateDevice {
    id: number
    name: string
    ip: string
    port: number
    username: string
    password: string
}

export interface INewDeviceDashboard {
    name: string
    isVisible: boolean
    deviceType: number
}

export interface IUpdateDeviceDashboard {
    id: number
    name: string
    isVisible: boolean
}

export interface INewDeviceDashboardTelemetry {
    deviceDashboardID: number
    telemetryIDs: number[]
}

export interface INewDeviceDashboardCamera {
    deviceDashboardID: number
    deviceIDs: number[]
}

export interface INewAttributeModel {
    key: string
    name: string
}

export interface IUpdateAttributeModel {
    id: number
    key: string
    name: string
}

export interface INewTelemetryModel {
    key: string
    name: string
    factor: number
    unit: string
    unitName: string
    scale: number
}

export interface IUpdateTelemetryModel {
    id: number
    key: string
    name: string
    factor: number
    unit: string
    unitName: string
    scale: number
}
