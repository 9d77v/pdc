import { TagProps } from "antd/lib/tag"

export interface TagStyle {
    color: TagProps["color"]
    text: string
}

export interface Combo {
    value: number
    text: string
}

export interface IVideoPagination {
    keyword: string
    page: number,
    pageSize: number,
    selectedTags: string[]
}

export interface IDeviceModel {
    id: number
    name: string
    desc: string
    deviceType: number
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
    deviceModelName: string,
    deviceModelDesc: string,
    deviceModelID: number
    deviceType: number
    health?: number
    createdAt: number
    updatedAt: number
}

export interface IDeviceDashboard {
    id: number
    name: string
    isVisible: boolean
    telemetries: IDeviceDashboardTelemetry[]
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
}

const RubbishCategoryMap = new Map<number, TagStyle>([
    [0, {
        color: 'black',
        text: "其他垃圾"
    }],
    [1, {
        color: 'blue',
        text: "可回收垃圾"
    }],
    [2, {
        color: 'green',
        text: "厨余垃圾"
    }],
    [3, {
        color: 'red',
        text: "有害垃圾"
    }],
])

const ConsumerExpenditureMap = new Map<string, string>([
    ["01", "食品烟酒"],
    ["02", "衣着"],
    ["03", "居住"],
    ["04", "生活用品及服务"],
    ["05", "交通通信"],
    ["06", "教育文化娱乐"],
    ["07", "医疗保健"],
    ["08", "其他用品及服务"],
])

const ThingStatusMap = new Map<number, TagStyle>([
    [0, {
        color: "#111111",
        text: "待采购"
    }],
    [1, {
        color: "#87d068",
        text: "使用中"
    }],
    [2, {
        color: "#4ada0c",
        text: "已收纳"
    }],
    [3, {
        color: "#da5e0c",
        text: "故障"
    }],
    [4, {
        color: "#eebb14",
        text: "维修中"
    }],
    [5, {
        color: "#928f8f",
        text: "待清理"
    }],
    [6, {
        color: "#d4d2cc",
        text: "已清理"
    }]
])


const GenderMap = new Map<number, string>([
    [0, "男"],
    [1, "女"],
    [2, '未知']
])

const RoleMap = new Map<number, string>([
    [2, "管理员"],
    [3, '普通用户'],
    [4, '访客']
])

const FullRoleMap = new Map<number, string>([
    [0, "未知"],
    [1, "所有者"],
    [2, "管理员"],
    [3, '普通用户'],
    [4, '访客']
])

const ThingStatusArr = ['待采购', '使用中', '已收纳', '故障', '维修中', '待清理', '已清理']

const DeviceTypeMap = new Map<number, string>([
    [0, "默认"]
])

const GesturePasswordKey = "gesture-password"


export {
    RubbishCategoryMap, ConsumerExpenditureMap, ThingStatusMap, ThingStatusArr,
    GenderMap, RoleMap, FullRoleMap, DeviceTypeMap, GesturePasswordKey
}
