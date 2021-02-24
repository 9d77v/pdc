import { ReactElement } from "react";

export enum NoteType {
    Directory = 0,
    File = 1
}

export enum NoteState {
    Normal = 0,
    InRubbish = 1,
    Deleted = 2
}

export enum SyncStatus {
    Unsync = 'unsync',
    Syncing = 'syncing',
    Synced = 'synced'
}

export interface INote {
    id: string // 笔记id
    parent_id: string// 笔记父节点
    uid: string// 用户id
    note_type: NoteType// 笔记类型
    level: number// 笔记层级，1，2，3
    title?: string// 笔记标题
    navTitle?: string// 导航标题
    content?: string// 笔记内容
    tags?: string[]// 笔记标签
    state?: NoteState// 笔记状态
    version: number// 笔记版本号
    sha1?: string// 笔记sha1值
    sync_status: SyncStatus// 笔记同步状态
    color?: string// 笔记颜色
    created_at?: string// 笔记创建时间
    updated_at?: string// 笔记更新时间
    editable?: boolean
}

export interface INoteTree {
    value: string// 值
    label: ReactElement// 标签
    title: string
    noteType: NoteType
    color: string
    children?: INoteTree[]// 子树
}
