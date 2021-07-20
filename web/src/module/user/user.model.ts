import dayjs from "dayjs"

export interface IUser {
    uid: number
    name: string
    password?: string
    avatar: string
    roleID: number
    gender: number
    color: string
    birthDate?: dayjs.Dayjs
    ip: string
}

export interface IUpdateUser {
    id: number
    name: string
    avatar: string
    password: string
    roleID: number
    gender: number
    color: string
    birthDate?: dayjs.Dayjs
    ip: string
}
