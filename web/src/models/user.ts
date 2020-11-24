export interface NewUser {
    uid: string
    name: string
    password: string
    avatar: string
    roleID: number
    gender: number
    birthDate: number
    ip: string
}

export interface IUpdateUser {
    id: number
    name: string
    avatar: string
    password: string
    roleID: number
    gender: number
    birthDate: number
    ip: string
}